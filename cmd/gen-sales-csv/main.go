package main

import (
	"context"
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Yoii-Inc/freee-gen-csv/internal/config"
	tokensource "github.com/Yoii-Inc/freee-gen-csv/internal/token"
	"github.com/Yoii-Inc/freee-gen-csv/pkg/freeeclient"

	"github.com/gocarina/gocsv"
	"golang.org/x/oauth2"
)

type CSVOut struct {
	PartnerId   int32  `csv:"customer_id"`
	StartDate   string `csv:"period_start_date_at"`
	EndDate     string `csv:"period_end_date_at"`
	PaymentDate string `csv:"payment_date_at"`
	Amount      int64  `csv:"value"`
}

var tokenFilename = "token.b"

func main() {
	env, err := config.LoadConfigForYaml()
	if err != nil {
		log.Fatal(err)
	}

	ac := flag.String("access-token", "", "Access Token")
	rf := flag.String("refresh-token", "", "Refresh Token")
	et := flag.String("expiry", "", "Expiry Time (ISO8601)")
	flag.Parse()

	ett, err := time.Parse(time.RFC3339, *et)
	if err != nil {
		fmt.Print(err)
	}

	token := &oauth2.Token{
		AccessToken:  *ac,
		RefreshToken: *rf,
		TokenType:    "bearer",
		Expiry:       ett,
	}

	if token.AccessToken == "" || token.RefreshToken == "" {
		func() {
			f, err := os.Open(tokenFilename)
			if err != nil {
				log.Fatal("access token and refresh token is not set")
			}
			defer f.Close()

			dec := gob.NewDecoder(f)
			if err := dec.Decode(&token); err != nil {
				log.Fatal("decode error of token.txt:", err)
			}
		}()
	} else {
		func() {
			f, err := os.Create(tokenFilename)
			if err != nil {
				log.Fatalf("cannot create %s", tokenFilename)
			}
			defer f.Close()

			enc := gob.NewEncoder(f)
			if err := enc.Encode(token); err != nil {
				log.Fatal(err)
			}
			fmt.Println("new token saved!")
		}()
	}
	fmt.Printf("token: %v\n", token)
	fmt.Printf("access token: %s\n", token.AccessToken)

	datefmt := "2006-01-02"
	tz, _ := time.LoadLocation("Asia/Tokyo")
	oauthcfg := &oauth2.Config{
		ClientID:     env.ClientId,
		ClientSecret: env.ClientSecret,
		Scopes:       []string{"read", "default_read"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.secure.freee.co.jp/public_api/authorize",
			TokenURL: "https://accounts.secure.freee.co.jp/public_api/token",
		},
		RedirectURL: env.RedirectURL,
	}
	cfg := freeeclient.NewConfiguration()
	ctx := context.Background()

	src := &tokensource.MyTokenSource{
		Src: oauthcfg.TokenSource(ctx, token),
		F:   Do,
	}
	reuseSrc := oauth2.ReuseTokenSource(token, src)
	ctx = context.WithValue(ctx, freeeclient.ContextOAuth2, reuseSrc)
	client := freeeclient.NewAPIClient(cfg)
	fmt.Printf("%v\n", ctx)

	accountItemMap := make(map[int32]freeeclient.AccountItemsResponseAccountItems)

	{
		resp, r, err := client.AccountItemsApi.GetAccountItems(ctx).CompanyId(env.CompanyId).Execute()
		if err != nil {
			fmt.Printf("Error when calling : %v\n", err)
			fmt.Printf("Full HTTP response: %v\n", r)
		}
		for _, data := range resp.AccountItems {
			accountItemMap[data.Id] = data
		}

	}

	var dealsAll []freeeclient.Deal

	for offset := int64(0); ; offset += 100 {
		resp, r, err := client.DealsApi.GetDeals(ctx).CompanyId(env.CompanyId).Offset(offset).Limit(int32(100)).Execute()
		if err != nil {
			fmt.Printf("Error when calling `InvoicesApi.GetInvoices``: %v\n", err)
			fmt.Printf("Full HTTP response: %v\n", r)
		}

		if resp.Meta.TotalCount > int32(len(dealsAll)) {
			dealsAll = append(dealsAll, resp.Deals...)
		} else {
			break
		}
		fmt.Printf("Got %d deals\n", len(dealsAll))
	}

	dealDetails := []CSVOut{}

	pushDetails := func(details interface{}, partner int32, date string) {
		var amount int64
		var tags []int32
		var item *int32
		var section *int32

		pushItem := func() {
			now, _ := time.ParseInLocation(datefmt, date, tz)
			// start := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, tz)
			// end := start.AddDate(0, 1, 0).Add(-1)

			dealDetails = append(dealDetails, CSVOut{
				Amount:      amount,
				PartnerId:   partner,
				StartDate:   now.Format(time.RFC3339),
				EndDate:     now.Format(time.RFC3339),
				PaymentDate: now.Format(time.RFC3339),
			})
		}
		switch v := details.(type) {
		case []freeeclient.DealResponseDealDetails:
			for j := range v {
				detail := v[j]
				amount = detail.Amount
				tags = detail.TagIds
				item = detail.ItemId.Get()
				section = detail.SectionId.Get()
				if accountItemMap[detail.AccountItemId].Name == "売上高" && filterDeal(item, section, &tags, &env.Filter.Items, &env.Filter.Sections, &env.Filter.Tags) {
					pushItem()
				}
			}
		case []freeeclient.DealCreateResponseDealDetails:
			for j := range v {
				detail := v[j]
				amount = detail.Amount
				if accountItemMap[detail.AccountItemId].Name == "売上高" {
					pushItem()
				}
			}
		}
	}

	for i := range dealsAll {
		partner := dealsAll[i].PartnerId
		if dealsAll[i].Details != nil {
			details := dealsAll[i].Details
			issueDate := dealsAll[i].IssueDate
			pushDetails(&details, partner, issueDate)
		}
		if dealsAll[i].Renews != nil {
			renews := *dealsAll[i].Renews
			for j := range renews {
				details := renews[j].Details
				updateDate := renews[j].UpdateDate
				pushDetails(details, partner, updateDate)
			}
		}
	}

	if f, err := os.Stat("out"); os.IsNotExist(err) || !f.IsDir() {
		if err := os.Mkdir("out", 0777); err != nil {
			log.Fatal(err)
		}
	}
	file, err := os.Create(config.Csvfname)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	gocsv.MarshalFile(&dealDetails, file)

}

func Do(t *oauth2.Token) error {
	fmt.Println("refreshed!")
	fmt.Printf("struct: %v\n", t)
	fmt.Printf("new access token: %s\n", t.AccessToken)
	fmt.Printf("new refresh token: %s\n", t.RefreshToken)

	f, err := os.Create(tokenFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if t.Valid() {
		enc := gob.NewEncoder(f)
		if err := enc.Encode(t); err != nil {
			log.Fatal(err)
		}
		fmt.Println("new token saved!")
	}
	return nil
}

func contains(s *[]int32, e int32) bool {
	for _, v := range *s {
		if e == v {
			return true
		}
	}
	return false
}

func filterDeal(dealItemId *int32, dealSectionId *int32, dealTagIds *[]int32, condItemIds *[]int32, condSectionIds *[]int32, condTagIds *[]int32) bool {
	tagFlag := true
	if condTagIds != nil && len(*condTagIds) > 0 {
		tagFlag = false
		for _, v := range *condTagIds {
			if contains(dealTagIds, v) {
				tagFlag = true
				break
			}
		}
	}

	itemFlag := true
	if condItemIds != nil && len(*condItemIds) > 0 {
		itemFlag = dealItemId != nil && contains(condItemIds, *dealItemId)
	}

	sectionFlag := true
	if condSectionIds != nil && len(*condSectionIds) > 0 {
		sectionFlag = dealSectionId != nil && contains(condSectionIds, *dealSectionId)
	}

	return tagFlag && itemFlag && sectionFlag
}
