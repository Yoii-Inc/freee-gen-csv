# \TrialBalanceApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTrialBs**](TrialBalanceApi.md#GetTrialBs) | **Get** /api/1/reports/trial_bs | 貸借対照表の取得
[**GetTrialBsThreeYears**](TrialBalanceApi.md#GetTrialBsThreeYears) | **Get** /api/1/reports/trial_bs_three_years | 貸借対照表(３期間比較)の取得
[**GetTrialBsTwoYears**](TrialBalanceApi.md#GetTrialBsTwoYears) | **Get** /api/1/reports/trial_bs_two_years | 貸借対照表(前年比較)の取得
[**GetTrialCr**](TrialBalanceApi.md#GetTrialCr) | **Get** /api/1/reports/trial_cr | 製造原価報告書の取得
[**GetTrialCrSections**](TrialBalanceApi.md#GetTrialCrSections) | **Get** /api/1/reports/trial_cr_sections | 製造原価報告書(部門比較)の取得
[**GetTrialCrSegment1Tags**](TrialBalanceApi.md#GetTrialCrSegment1Tags) | **Get** /api/1/reports/trial_cr_segment_1_tags | 製造原価報告書(セグメント1比較)の取得
[**GetTrialCrSegment2Tags**](TrialBalanceApi.md#GetTrialCrSegment2Tags) | **Get** /api/1/reports/trial_cr_segment_2_tags | 製造原価報告書(セグメント2比較)の取得
[**GetTrialCrSegment3Tags**](TrialBalanceApi.md#GetTrialCrSegment3Tags) | **Get** /api/1/reports/trial_cr_segment_3_tags | 製造原価報告書(セグメント3比較)の取得
[**GetTrialCrThreeYears**](TrialBalanceApi.md#GetTrialCrThreeYears) | **Get** /api/1/reports/trial_cr_three_years | 製造原価報告書(３期間比較)の取得
[**GetTrialCrTwoYears**](TrialBalanceApi.md#GetTrialCrTwoYears) | **Get** /api/1/reports/trial_cr_two_years | 製造原価報告書(前年比較)の取得
[**GetTrialPl**](TrialBalanceApi.md#GetTrialPl) | **Get** /api/1/reports/trial_pl | 損益計算書の取得
[**GetTrialPlSections**](TrialBalanceApi.md#GetTrialPlSections) | **Get** /api/1/reports/trial_pl_sections | 損益計算書(部門比較)の取得
[**GetTrialPlSegment1Tags**](TrialBalanceApi.md#GetTrialPlSegment1Tags) | **Get** /api/1/reports/trial_pl_segment_1_tags | 損益計算書(セグメント1比較)の取得
[**GetTrialPlSegment2Tags**](TrialBalanceApi.md#GetTrialPlSegment2Tags) | **Get** /api/1/reports/trial_pl_segment_2_tags | 損益計算書(セグメント2比較)の取得
[**GetTrialPlSegment3Tags**](TrialBalanceApi.md#GetTrialPlSegment3Tags) | **Get** /api/1/reports/trial_pl_segment_3_tags | 損益計算書(セグメント3比較)の取得
[**GetTrialPlThreeYears**](TrialBalanceApi.md#GetTrialPlThreeYears) | **Get** /api/1/reports/trial_pl_three_years | 損益計算書(３期間比較)の取得
[**GetTrialPlTwoYears**](TrialBalanceApi.md#GetTrialPlTwoYears) | **Get** /api/1/reports/trial_pl_two_years | 損益計算書(前年比較)の取得



## GetTrialBs

> TrialBsResponse GetTrialBs(ctx).CompanyId(companyId).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).ApprovalFlowStatus(approvalFlowStatus).Execute()

貸借対照表の取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyId := int32(56) // int32 | 事業所ID
    fiscalYear := int32(56) // int32 | 会計年度 (optional)
    startMonth := int32(56) // int32 | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 (optional)
    endMonth := int32(56) // int32 | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 (optional)
    startDate := "startDate_example" // string | 発生日で絞込：開始日(yyyy-mm-dd) (optional)
    endDate := "endDate_example" // string | 発生日で絞込：終了日(yyyy-mm-dd) (optional)
    accountItemDisplayType := "accountItemDisplayType_example" // string | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 (optional)
    breakdownDisplayType := "breakdownDisplayType_example" // string | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item, セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン): segment_1_tag, セグメント2(法人向け エンタープライズプラン):segment_2_tag, セグメント3(法人向け エンタープライズプラン): segment_3_tag） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます (optional)
    partnerId := int32(56) // int32 | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） (optional)
    partnerCode := "partnerCode_example" // string | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） (optional)
    itemId := int32(56) // int32 | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） (optional)
    sectionId := int32(56) // int32 | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） (optional)
    adjustment := "adjustment_example" // string | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 (optional)
    approvalFlowStatus := "approvalFlowStatus_example" // string | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)<br> 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。<br> 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrialBalanceApi.GetTrialBs(context.Background()).CompanyId(companyId).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).ApprovalFlowStatus(approvalFlowStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrialBalanceApi.GetTrialBs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrialBs`: TrialBsResponse
    fmt.Fprintf(os.Stdout, "Response from `TrialBalanceApi.GetTrialBs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrialBsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **fiscalYear** | **int32** | 会計年度 | 
 **startMonth** | **int32** | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 | 
 **endMonth** | **int32** | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 | 
 **startDate** | **string** | 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **string** | 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **string** | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 | 
 **breakdownDisplayType** | **string** | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item, セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン): segment_1_tag, セグメント2(法人向け エンタープライズプラン):segment_2_tag, セグメント3(法人向け エンタープライズプラン): segment_3_tag） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **int32** | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **string** | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **int32** | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **sectionId** | **int32** | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） | 
 **adjustment** | **string** | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 | 
 **approvalFlowStatus** | **string** | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)&lt;br&gt; 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。&lt;br&gt; 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  | 

### Return type

[**TrialBsResponse**](TrialBsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrialBsThreeYears

> TrialBsThreeYearsResponse GetTrialBsThreeYears(ctx).CompanyId(companyId).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).ApprovalFlowStatus(approvalFlowStatus).Execute()

貸借対照表(３期間比較)の取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyId := int32(56) // int32 | 事業所ID
    fiscalYear := int32(56) // int32 | 会計年度 (optional)
    startMonth := int32(56) // int32 | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 (optional)
    endMonth := int32(56) // int32 | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 (optional)
    startDate := "startDate_example" // string | 発生日で絞込：開始日(yyyy-mm-dd) (optional)
    endDate := "endDate_example" // string | 発生日で絞込：終了日(yyyy-mm-dd) (optional)
    accountItemDisplayType := "accountItemDisplayType_example" // string | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 (optional)
    breakdownDisplayType := "breakdownDisplayType_example" // string | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item, セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン): segment_1_tag, セグメント2(法人向け エンタープライズプラン):segment_2_tag, セグメント3(法人向け エンタープライズプラン): segment_3_tag） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます (optional)
    partnerId := int32(56) // int32 | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） (optional)
    partnerCode := "partnerCode_example" // string | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） (optional)
    itemId := int32(56) // int32 | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） (optional)
    sectionId := int32(56) // int32 | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） (optional)
    adjustment := "adjustment_example" // string | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 (optional)
    approvalFlowStatus := "approvalFlowStatus_example" // string | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)<br> 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。<br> 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrialBalanceApi.GetTrialBsThreeYears(context.Background()).CompanyId(companyId).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).ApprovalFlowStatus(approvalFlowStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrialBalanceApi.GetTrialBsThreeYears``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrialBsThreeYears`: TrialBsThreeYearsResponse
    fmt.Fprintf(os.Stdout, "Response from `TrialBalanceApi.GetTrialBsThreeYears`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrialBsThreeYearsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **fiscalYear** | **int32** | 会計年度 | 
 **startMonth** | **int32** | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 | 
 **endMonth** | **int32** | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 | 
 **startDate** | **string** | 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **string** | 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **string** | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 | 
 **breakdownDisplayType** | **string** | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item, セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン): segment_1_tag, セグメント2(法人向け エンタープライズプラン):segment_2_tag, セグメント3(法人向け エンタープライズプラン): segment_3_tag） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **int32** | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **string** | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **int32** | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **sectionId** | **int32** | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） | 
 **adjustment** | **string** | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 | 
 **approvalFlowStatus** | **string** | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)&lt;br&gt; 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。&lt;br&gt; 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  | 

### Return type

[**TrialBsThreeYearsResponse**](TrialBsThreeYearsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrialBsTwoYears

> TrialBsTwoYearsResponse GetTrialBsTwoYears(ctx).CompanyId(companyId).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).ApprovalFlowStatus(approvalFlowStatus).Execute()

貸借対照表(前年比較)の取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyId := int32(56) // int32 | 事業所ID
    fiscalYear := int32(56) // int32 | 会計年度 (optional)
    startMonth := int32(56) // int32 | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 (optional)
    endMonth := int32(56) // int32 | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 (optional)
    startDate := "startDate_example" // string | 発生日で絞込：開始日(yyyy-mm-dd) (optional)
    endDate := "endDate_example" // string | 発生日で絞込：終了日(yyyy-mm-dd) (optional)
    accountItemDisplayType := "accountItemDisplayType_example" // string | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 (optional)
    breakdownDisplayType := "breakdownDisplayType_example" // string | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item, セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン): segment_1_tag, セグメント2(法人向け エンタープライズプラン):segment_2_tag, セグメント3(法人向け エンタープライズプラン): segment_3_tag） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます (optional)
    partnerId := int32(56) // int32 | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） (optional)
    partnerCode := "partnerCode_example" // string | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） (optional)
    itemId := int32(56) // int32 | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） (optional)
    sectionId := int32(56) // int32 | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） (optional)
    adjustment := "adjustment_example" // string | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 (optional)
    approvalFlowStatus := "approvalFlowStatus_example" // string | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)<br> 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。<br> 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrialBalanceApi.GetTrialBsTwoYears(context.Background()).CompanyId(companyId).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).ApprovalFlowStatus(approvalFlowStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrialBalanceApi.GetTrialBsTwoYears``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrialBsTwoYears`: TrialBsTwoYearsResponse
    fmt.Fprintf(os.Stdout, "Response from `TrialBalanceApi.GetTrialBsTwoYears`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrialBsTwoYearsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **fiscalYear** | **int32** | 会計年度 | 
 **startMonth** | **int32** | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 | 
 **endMonth** | **int32** | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 | 
 **startDate** | **string** | 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **string** | 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **string** | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 | 
 **breakdownDisplayType** | **string** | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item, セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン): segment_1_tag, セグメント2(法人向け エンタープライズプラン):segment_2_tag, セグメント3(法人向け エンタープライズプラン): segment_3_tag） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **int32** | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **string** | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **int32** | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **sectionId** | **int32** | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） | 
 **adjustment** | **string** | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 | 
 **approvalFlowStatus** | **string** | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)&lt;br&gt; 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。&lt;br&gt; 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  | 

### Return type

[**TrialBsTwoYearsResponse**](TrialBsTwoYearsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrialCr

> TrialCrResponse GetTrialCr(ctx).CompanyId(companyId).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).CostAllocation(costAllocation).ApprovalFlowStatus(approvalFlowStatus).Execute()

製造原価報告書の取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyId := int32(56) // int32 | 事業所ID
    fiscalYear := int32(56) // int32 | 会計年度 (optional)
    startMonth := int32(56) // int32 | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 (optional)
    endMonth := int32(56) // int32 | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 (optional)
    startDate := "startDate_example" // string | 発生日で絞込：開始日(yyyy-mm-dd) (optional)
    endDate := "endDate_example" // string | 発生日で絞込：終了日(yyyy-mm-dd) (optional)
    accountItemDisplayType := "accountItemDisplayType_example" // string | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 (optional)
    breakdownDisplayType := "breakdownDisplayType_example" // string | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item, セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン): segment_1_tag, セグメント2(法人向け エンタープライズプラン):segment_2_tag, セグメント3(法人向け エンタープライズプラン): segment_3_tag） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます (optional)
    partnerId := int32(56) // int32 | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） (optional)
    partnerCode := "partnerCode_example" // string | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） (optional)
    itemId := int32(56) // int32 | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） (optional)
    sectionId := int32(56) // int32 | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） (optional)
    adjustment := "adjustment_example" // string | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 (optional)
    costAllocation := "costAllocation_example" // string | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without）。指定されない場合、配賦仕訳を含む金額が返却されます。 (optional)
    approvalFlowStatus := "approvalFlowStatus_example" // string | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト), 全てのステータス: all)<br> 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。<br> 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrialBalanceApi.GetTrialCr(context.Background()).CompanyId(companyId).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).CostAllocation(costAllocation).ApprovalFlowStatus(approvalFlowStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrialBalanceApi.GetTrialCr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrialCr`: TrialCrResponse
    fmt.Fprintf(os.Stdout, "Response from `TrialBalanceApi.GetTrialCr`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrialCrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **fiscalYear** | **int32** | 会計年度 | 
 **startMonth** | **int32** | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 | 
 **endMonth** | **int32** | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 | 
 **startDate** | **string** | 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **string** | 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **string** | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 | 
 **breakdownDisplayType** | **string** | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item, セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン): segment_1_tag, セグメント2(法人向け エンタープライズプラン):segment_2_tag, セグメント3(法人向け エンタープライズプラン): segment_3_tag） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **int32** | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **string** | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **int32** | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **sectionId** | **int32** | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） | 
 **adjustment** | **string** | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 | 
 **costAllocation** | **string** | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without）。指定されない場合、配賦仕訳を含む金額が返却されます。 | 
 **approvalFlowStatus** | **string** | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト), 全てのステータス: all)&lt;br&gt; 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。&lt;br&gt; 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  | 

### Return type

[**TrialCrResponse**](TrialCrResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrialCrSections

> TrialCrSectionsResponse GetTrialCrSections(ctx).CompanyId(companyId).SectionIds(sectionIds).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).Adjustment(adjustment).CostAllocation(costAllocation).ApprovalFlowStatus(approvalFlowStatus).Execute()

製造原価報告書(部門比較)の取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyId := int32(56) // int32 | 事業所ID
    sectionIds := "sectionIds_example" // string | 出力する部門の指定（半角数字のidを半角カンマ区切りスペースなしで指定してください。0を指定すると、未選択の部門で比較できます）
    fiscalYear := int32(56) // int32 | 会計年度 (optional)
    startMonth := int32(56) // int32 | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 (optional)
    endMonth := int32(56) // int32 | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 (optional)
    startDate := "startDate_example" // string | 発生日で絞込：開始日(yyyy-mm-dd) (optional)
    endDate := "endDate_example" // string | 発生日で絞込：終了日(yyyy-mm-dd) (optional)
    accountItemDisplayType := "accountItemDisplayType_example" // string | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 (optional)
    breakdownDisplayType := "breakdownDisplayType_example" // string | 内訳の表示（取引先: partner, 品目: item, 勘定科目: account_item, セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン): segment_1_tag, セグメント2(法人向け エンタープライズプラン):segment_2_tag, セグメント3(法人向け エンタープライズプラン): segment_3_tag） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます (optional)
    partnerId := int32(56) // int32 | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） (optional)
    partnerCode := "partnerCode_example" // string | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） (optional)
    itemId := int32(56) // int32 | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） (optional)
    adjustment := "adjustment_example" // string | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 (optional)
    costAllocation := "costAllocation_example" // string | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without）。指定されない場合、配賦仕訳を含む金額が返却されます。 (optional)
    approvalFlowStatus := "approvalFlowStatus_example" // string | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)<br> 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。<br> 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrialBalanceApi.GetTrialCrSections(context.Background()).CompanyId(companyId).SectionIds(sectionIds).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).Adjustment(adjustment).CostAllocation(costAllocation).ApprovalFlowStatus(approvalFlowStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrialBalanceApi.GetTrialCrSections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrialCrSections`: TrialCrSectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `TrialBalanceApi.GetTrialCrSections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrialCrSectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **sectionIds** | **string** | 出力する部門の指定（半角数字のidを半角カンマ区切りスペースなしで指定してください。0を指定すると、未選択の部門で比較できます） | 
 **fiscalYear** | **int32** | 会計年度 | 
 **startMonth** | **int32** | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 | 
 **endMonth** | **int32** | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 | 
 **startDate** | **string** | 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **string** | 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **string** | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 | 
 **breakdownDisplayType** | **string** | 内訳の表示（取引先: partner, 品目: item, 勘定科目: account_item, セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン): segment_1_tag, セグメント2(法人向け エンタープライズプラン):segment_2_tag, セグメント3(法人向け エンタープライズプラン): segment_3_tag） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **int32** | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **string** | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **int32** | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **adjustment** | **string** | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 | 
 **costAllocation** | **string** | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without）。指定されない場合、配賦仕訳を含む金額が返却されます。 | 
 **approvalFlowStatus** | **string** | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)&lt;br&gt; 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。&lt;br&gt; 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  | 

### Return type

[**TrialCrSectionsResponse**](TrialCrSectionsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrialCrSegment1Tags

> TrialCrSegment1TagsResponse GetTrialCrSegment1Tags(ctx).CompanyId(companyId).Segment1TagIds(segment1TagIds).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).CostAllocation(costAllocation).ApprovalFlowStatus(approvalFlowStatus).Execute()

製造原価報告書(セグメント1比較)の取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyId := int32(56) // int32 | 事業所ID
    segment1TagIds := "segment1TagIds_example" // string | 出力するセグメント1タグの指定（半角数字のidを半角カンマ区切りスペースなしで指定してください。0を指定すると、未選択のセグメントで比較できます）
    fiscalYear := int32(56) // int32 | 会計年度 (optional)
    startMonth := int32(56) // int32 | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 (optional)
    endMonth := int32(56) // int32 | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 (optional)
    startDate := "startDate_example" // string | 発生日で絞込：開始日(yyyy-mm-dd) (optional)
    endDate := "endDate_example" // string | 発生日で絞込：終了日(yyyy-mm-dd) (optional)
    accountItemDisplayType := "accountItemDisplayType_example" // string | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 (optional)
    breakdownDisplayType := "breakdownDisplayType_example" // string | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます (optional)
    partnerId := int32(56) // int32 | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） (optional)
    partnerCode := "partnerCode_example" // string | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） (optional)
    itemId := int32(56) // int32 | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） (optional)
    sectionId := int32(56) // int32 | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） (optional)
    adjustment := "adjustment_example" // string | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 (optional)
    costAllocation := "costAllocation_example" // string | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without）。指定されない場合、配賦仕訳を含む金額が返却されます。 (optional)
    approvalFlowStatus := "approvalFlowStatus_example" // string | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)<br> 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。<br> 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrialBalanceApi.GetTrialCrSegment1Tags(context.Background()).CompanyId(companyId).Segment1TagIds(segment1TagIds).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).CostAllocation(costAllocation).ApprovalFlowStatus(approvalFlowStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrialBalanceApi.GetTrialCrSegment1Tags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrialCrSegment1Tags`: TrialCrSegment1TagsResponse
    fmt.Fprintf(os.Stdout, "Response from `TrialBalanceApi.GetTrialCrSegment1Tags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrialCrSegment1TagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **segment1TagIds** | **string** | 出力するセグメント1タグの指定（半角数字のidを半角カンマ区切りスペースなしで指定してください。0を指定すると、未選択のセグメントで比較できます） | 
 **fiscalYear** | **int32** | 会計年度 | 
 **startMonth** | **int32** | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 | 
 **endMonth** | **int32** | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 | 
 **startDate** | **string** | 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **string** | 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **string** | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 | 
 **breakdownDisplayType** | **string** | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **int32** | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **string** | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **int32** | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **sectionId** | **int32** | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） | 
 **adjustment** | **string** | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 | 
 **costAllocation** | **string** | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without）。指定されない場合、配賦仕訳を含む金額が返却されます。 | 
 **approvalFlowStatus** | **string** | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)&lt;br&gt; 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。&lt;br&gt; 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  | 

### Return type

[**TrialCrSegment1TagsResponse**](TrialCrSegment1TagsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrialCrSegment2Tags

> TrialCrSegment2TagsResponse GetTrialCrSegment2Tags(ctx).CompanyId(companyId).Segment2TagIds(segment2TagIds).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).CostAllocation(costAllocation).ApprovalFlowStatus(approvalFlowStatus).Execute()

製造原価報告書(セグメント2比較)の取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyId := int32(56) // int32 | 事業所ID
    segment2TagIds := "segment2TagIds_example" // string | 出力するセグメント2タグの指定（半角数字のidを半角カンマ区切りスペースなしで指定してください。0を指定すると、未選択のセグメントで比較できます）
    fiscalYear := int32(56) // int32 | 会計年度 (optional)
    startMonth := int32(56) // int32 | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 (optional)
    endMonth := int32(56) // int32 | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 (optional)
    startDate := "startDate_example" // string | 発生日で絞込：開始日(yyyy-mm-dd) (optional)
    endDate := "endDate_example" // string | 発生日で絞込：終了日(yyyy-mm-dd) (optional)
    accountItemDisplayType := "accountItemDisplayType_example" // string | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 (optional)
    breakdownDisplayType := "breakdownDisplayType_example" // string | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます (optional)
    partnerId := int32(56) // int32 | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） (optional)
    partnerCode := "partnerCode_example" // string | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） (optional)
    itemId := int32(56) // int32 | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） (optional)
    sectionId := int32(56) // int32 | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） (optional)
    adjustment := "adjustment_example" // string | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 (optional)
    costAllocation := "costAllocation_example" // string | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without）。指定されない場合、配賦仕訳を含む金額が返却されます。 (optional)
    approvalFlowStatus := "approvalFlowStatus_example" // string | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)<br> 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。<br> 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrialBalanceApi.GetTrialCrSegment2Tags(context.Background()).CompanyId(companyId).Segment2TagIds(segment2TagIds).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).CostAllocation(costAllocation).ApprovalFlowStatus(approvalFlowStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrialBalanceApi.GetTrialCrSegment2Tags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrialCrSegment2Tags`: TrialCrSegment2TagsResponse
    fmt.Fprintf(os.Stdout, "Response from `TrialBalanceApi.GetTrialCrSegment2Tags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrialCrSegment2TagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **segment2TagIds** | **string** | 出力するセグメント2タグの指定（半角数字のidを半角カンマ区切りスペースなしで指定してください。0を指定すると、未選択のセグメントで比較できます） | 
 **fiscalYear** | **int32** | 会計年度 | 
 **startMonth** | **int32** | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 | 
 **endMonth** | **int32** | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 | 
 **startDate** | **string** | 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **string** | 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **string** | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 | 
 **breakdownDisplayType** | **string** | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **int32** | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **string** | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **int32** | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **sectionId** | **int32** | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） | 
 **adjustment** | **string** | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 | 
 **costAllocation** | **string** | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without）。指定されない場合、配賦仕訳を含む金額が返却されます。 | 
 **approvalFlowStatus** | **string** | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)&lt;br&gt; 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。&lt;br&gt; 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  | 

### Return type

[**TrialCrSegment2TagsResponse**](TrialCrSegment2TagsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrialCrSegment3Tags

> TrialCrSegment3TagsResponse GetTrialCrSegment3Tags(ctx).CompanyId(companyId).Segment3TagIds(segment3TagIds).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).CostAllocation(costAllocation).ApprovalFlowStatus(approvalFlowStatus).Execute()

製造原価報告書(セグメント3比較)の取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyId := int32(56) // int32 | 事業所ID
    segment3TagIds := "segment3TagIds_example" // string | 出力するセグメント3タグの指定（半角数字のidを半角カンマ区切りスペースなしで指定してください。0を指定すると、未選択のセグメントで比較できます）
    fiscalYear := int32(56) // int32 | 会計年度 (optional)
    startMonth := int32(56) // int32 | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 (optional)
    endMonth := int32(56) // int32 | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 (optional)
    startDate := "startDate_example" // string | 発生日で絞込：開始日(yyyy-mm-dd) (optional)
    endDate := "endDate_example" // string | 発生日で絞込：終了日(yyyy-mm-dd) (optional)
    accountItemDisplayType := "accountItemDisplayType_example" // string | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 (optional)
    breakdownDisplayType := "breakdownDisplayType_example" // string | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます (optional)
    partnerId := int32(56) // int32 | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） (optional)
    partnerCode := "partnerCode_example" // string | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） (optional)
    itemId := int32(56) // int32 | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） (optional)
    sectionId := int32(56) // int32 | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） (optional)
    adjustment := "adjustment_example" // string | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 (optional)
    costAllocation := "costAllocation_example" // string | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without）。指定されない場合、配賦仕訳を含む金額が返却されます。 (optional)
    approvalFlowStatus := "approvalFlowStatus_example" // string | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)<br> 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。<br> 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrialBalanceApi.GetTrialCrSegment3Tags(context.Background()).CompanyId(companyId).Segment3TagIds(segment3TagIds).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).CostAllocation(costAllocation).ApprovalFlowStatus(approvalFlowStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrialBalanceApi.GetTrialCrSegment3Tags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrialCrSegment3Tags`: TrialCrSegment3TagsResponse
    fmt.Fprintf(os.Stdout, "Response from `TrialBalanceApi.GetTrialCrSegment3Tags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrialCrSegment3TagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **segment3TagIds** | **string** | 出力するセグメント3タグの指定（半角数字のidを半角カンマ区切りスペースなしで指定してください。0を指定すると、未選択のセグメントで比較できます） | 
 **fiscalYear** | **int32** | 会計年度 | 
 **startMonth** | **int32** | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 | 
 **endMonth** | **int32** | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 | 
 **startDate** | **string** | 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **string** | 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **string** | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 | 
 **breakdownDisplayType** | **string** | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **int32** | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **string** | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **int32** | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **sectionId** | **int32** | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） | 
 **adjustment** | **string** | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 | 
 **costAllocation** | **string** | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without）。指定されない場合、配賦仕訳を含む金額が返却されます。 | 
 **approvalFlowStatus** | **string** | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)&lt;br&gt; 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。&lt;br&gt; 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  | 

### Return type

[**TrialCrSegment3TagsResponse**](TrialCrSegment3TagsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrialCrThreeYears

> TrialCrThreeYearsResponse GetTrialCrThreeYears(ctx).CompanyId(companyId).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).CostAllocation(costAllocation).ApprovalFlowStatus(approvalFlowStatus).Execute()

製造原価報告書(３期間比較)の取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyId := int32(56) // int32 | 事業所ID
    fiscalYear := int32(56) // int32 | 会計年度 (optional)
    startMonth := int32(56) // int32 | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 (optional)
    endMonth := int32(56) // int32 | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 (optional)
    startDate := "startDate_example" // string | 発生日で絞込：開始日(yyyy-mm-dd) (optional)
    endDate := "endDate_example" // string | 発生日で絞込：終了日(yyyy-mm-dd) (optional)
    accountItemDisplayType := "accountItemDisplayType_example" // string | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 (optional)
    breakdownDisplayType := "breakdownDisplayType_example" // string | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item, セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン): segment_1_tag, セグメント2(法人向け エンタープライズプラン):segment_2_tag, セグメント3(法人向け エンタープライズプラン): segment_3_tag） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます (optional)
    partnerId := int32(56) // int32 | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） (optional)
    partnerCode := "partnerCode_example" // string | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） (optional)
    itemId := int32(56) // int32 | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） (optional)
    sectionId := int32(56) // int32 | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） (optional)
    adjustment := "adjustment_example" // string | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 (optional)
    costAllocation := "costAllocation_example" // string | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without）。指定されない場合、配賦仕訳を含む金額が返却されます。 (optional)
    approvalFlowStatus := "approvalFlowStatus_example" // string | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト), 全てのステータス: all)<br> 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。<br> 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrialBalanceApi.GetTrialCrThreeYears(context.Background()).CompanyId(companyId).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).CostAllocation(costAllocation).ApprovalFlowStatus(approvalFlowStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrialBalanceApi.GetTrialCrThreeYears``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrialCrThreeYears`: TrialCrThreeYearsResponse
    fmt.Fprintf(os.Stdout, "Response from `TrialBalanceApi.GetTrialCrThreeYears`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrialCrThreeYearsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **fiscalYear** | **int32** | 会計年度 | 
 **startMonth** | **int32** | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 | 
 **endMonth** | **int32** | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 | 
 **startDate** | **string** | 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **string** | 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **string** | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 | 
 **breakdownDisplayType** | **string** | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item, セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン): segment_1_tag, セグメント2(法人向け エンタープライズプラン):segment_2_tag, セグメント3(法人向け エンタープライズプラン): segment_3_tag） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **int32** | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **string** | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **int32** | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **sectionId** | **int32** | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） | 
 **adjustment** | **string** | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 | 
 **costAllocation** | **string** | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without）。指定されない場合、配賦仕訳を含む金額が返却されます。 | 
 **approvalFlowStatus** | **string** | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト), 全てのステータス: all)&lt;br&gt; 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。&lt;br&gt; 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  | 

### Return type

[**TrialCrThreeYearsResponse**](TrialCrThreeYearsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrialCrTwoYears

> TrialCrTwoYearsResponse GetTrialCrTwoYears(ctx).CompanyId(companyId).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).CostAllocation(costAllocation).ApprovalFlowStatus(approvalFlowStatus).Execute()

製造原価報告書(前年比較)の取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyId := int32(56) // int32 | 事業所ID
    fiscalYear := int32(56) // int32 | 会計年度 (optional)
    startMonth := int32(56) // int32 | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 (optional)
    endMonth := int32(56) // int32 | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 (optional)
    startDate := "startDate_example" // string | 発生日で絞込：開始日(yyyy-mm-dd) (optional)
    endDate := "endDate_example" // string | 発生日で絞込：終了日(yyyy-mm-dd) (optional)
    accountItemDisplayType := "accountItemDisplayType_example" // string | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 (optional)
    breakdownDisplayType := "breakdownDisplayType_example" // string | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item, セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン): segment_1_tag, セグメント2(法人向け エンタープライズプラン):segment_2_tag, セグメント3(法人向け エンタープライズプラン): segment_3_tag） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます (optional)
    partnerId := int32(56) // int32 | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） (optional)
    partnerCode := "partnerCode_example" // string | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） (optional)
    itemId := int32(56) // int32 | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） (optional)
    sectionId := int32(56) // int32 | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） (optional)
    adjustment := "adjustment_example" // string | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 (optional)
    costAllocation := "costAllocation_example" // string | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without）。指定されない場合、配賦仕訳を含む金額が返却されます。 (optional)
    approvalFlowStatus := "approvalFlowStatus_example" // string | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト), 全てのステータス: all)<br> 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。<br> 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrialBalanceApi.GetTrialCrTwoYears(context.Background()).CompanyId(companyId).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).CostAllocation(costAllocation).ApprovalFlowStatus(approvalFlowStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrialBalanceApi.GetTrialCrTwoYears``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrialCrTwoYears`: TrialCrTwoYearsResponse
    fmt.Fprintf(os.Stdout, "Response from `TrialBalanceApi.GetTrialCrTwoYears`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrialCrTwoYearsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **fiscalYear** | **int32** | 会計年度 | 
 **startMonth** | **int32** | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 | 
 **endMonth** | **int32** | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 | 
 **startDate** | **string** | 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **string** | 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **string** | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 | 
 **breakdownDisplayType** | **string** | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item, セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン): segment_1_tag, セグメント2(法人向け エンタープライズプラン):segment_2_tag, セグメント3(法人向け エンタープライズプラン): segment_3_tag） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **int32** | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **string** | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **int32** | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **sectionId** | **int32** | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） | 
 **adjustment** | **string** | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 | 
 **costAllocation** | **string** | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without）。指定されない場合、配賦仕訳を含む金額が返却されます。 | 
 **approvalFlowStatus** | **string** | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト), 全てのステータス: all)&lt;br&gt; 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。&lt;br&gt; 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  | 

### Return type

[**TrialCrTwoYearsResponse**](TrialCrTwoYearsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrialPl

> TrialPlResponse GetTrialPl(ctx).CompanyId(companyId).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).CostAllocation(costAllocation).ApprovalFlowStatus(approvalFlowStatus).Execute()

損益計算書の取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyId := int32(56) // int32 | 事業所ID
    fiscalYear := int32(56) // int32 | 会計年度 (optional)
    startMonth := int32(56) // int32 | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 (optional)
    endMonth := int32(56) // int32 | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 (optional)
    startDate := "startDate_example" // string | 発生日で絞込：開始日(yyyy-mm-dd) (optional)
    endDate := "endDate_example" // string | 発生日で絞込：終了日(yyyy-mm-dd) (optional)
    accountItemDisplayType := "accountItemDisplayType_example" // string | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 (optional)
    breakdownDisplayType := "breakdownDisplayType_example" // string | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item, セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン): segment_1_tag, セグメント2(法人向け エンタープライズプラン):segment_2_tag, セグメント3(法人向け エンタープライズプラン): segment_3_tag） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます (optional)
    partnerId := int32(56) // int32 | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） (optional)
    partnerCode := "partnerCode_example" // string | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） (optional)
    itemId := int32(56) // int32 | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） (optional)
    sectionId := int32(56) // int32 | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） (optional)
    adjustment := "adjustment_example" // string | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 (optional)
    costAllocation := "costAllocation_example" // string | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without）。指定されない場合、配賦仕訳を含む金額が返却されます。 (optional)
    approvalFlowStatus := "approvalFlowStatus_example" // string | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)<br> 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。<br> 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrialBalanceApi.GetTrialPl(context.Background()).CompanyId(companyId).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).CostAllocation(costAllocation).ApprovalFlowStatus(approvalFlowStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrialBalanceApi.GetTrialPl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrialPl`: TrialPlResponse
    fmt.Fprintf(os.Stdout, "Response from `TrialBalanceApi.GetTrialPl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrialPlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **fiscalYear** | **int32** | 会計年度 | 
 **startMonth** | **int32** | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 | 
 **endMonth** | **int32** | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 | 
 **startDate** | **string** | 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **string** | 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **string** | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 | 
 **breakdownDisplayType** | **string** | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item, セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン): segment_1_tag, セグメント2(法人向け エンタープライズプラン):segment_2_tag, セグメント3(法人向け エンタープライズプラン): segment_3_tag） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **int32** | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **string** | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **int32** | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **sectionId** | **int32** | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） | 
 **adjustment** | **string** | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 | 
 **costAllocation** | **string** | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without）。指定されない場合、配賦仕訳を含む金額が返却されます。 | 
 **approvalFlowStatus** | **string** | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)&lt;br&gt; 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。&lt;br&gt; 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  | 

### Return type

[**TrialPlResponse**](TrialPlResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrialPlSections

> TrialPlSectionsResponse GetTrialPlSections(ctx).CompanyId(companyId).SectionIds(sectionIds).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).Adjustment(adjustment).CostAllocation(costAllocation).ApprovalFlowStatus(approvalFlowStatus).Execute()

損益計算書(部門比較)の取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyId := int32(56) // int32 | 事業所ID
    sectionIds := "sectionIds_example" // string | 出力する部門の指定（半角数字のidを半角カンマ区切りスペースなしで指定してください。0を指定すると、未選択の部門で比較できます。）
    fiscalYear := int32(56) // int32 | 会計年度 (optional)
    startMonth := int32(56) // int32 | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 (optional)
    endMonth := int32(56) // int32 | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 (optional)
    startDate := "startDate_example" // string | 発生日で絞込：開始日(yyyy-mm-dd) (optional)
    endDate := "endDate_example" // string | 発生日で絞込：終了日(yyyy-mm-dd) (optional)
    accountItemDisplayType := "accountItemDisplayType_example" // string | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 (optional)
    breakdownDisplayType := "breakdownDisplayType_example" // string | 内訳の表示（取引先: partner, 品目: item, 勘定科目: account_item, セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン): segment_1_tag, セグメント2(法人向け エンタープライズプラン):segment_2_tag, セグメント3(法人向け エンタープライズプラン): segment_3_tag） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます (optional)
    partnerId := int32(56) // int32 | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） (optional)
    partnerCode := "partnerCode_example" // string | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） (optional)
    itemId := int32(56) // int32 | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） (optional)
    adjustment := "adjustment_example" // string | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 (optional)
    costAllocation := "costAllocation_example" // string | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without）。指定されない場合、配賦仕訳を含む金額が返却されます。 (optional)
    approvalFlowStatus := "approvalFlowStatus_example" // string | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)<br> 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。<br> 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrialBalanceApi.GetTrialPlSections(context.Background()).CompanyId(companyId).SectionIds(sectionIds).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).Adjustment(adjustment).CostAllocation(costAllocation).ApprovalFlowStatus(approvalFlowStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrialBalanceApi.GetTrialPlSections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrialPlSections`: TrialPlSectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `TrialBalanceApi.GetTrialPlSections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrialPlSectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **sectionIds** | **string** | 出力する部門の指定（半角数字のidを半角カンマ区切りスペースなしで指定してください。0を指定すると、未選択の部門で比較できます。） | 
 **fiscalYear** | **int32** | 会計年度 | 
 **startMonth** | **int32** | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 | 
 **endMonth** | **int32** | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 | 
 **startDate** | **string** | 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **string** | 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **string** | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 | 
 **breakdownDisplayType** | **string** | 内訳の表示（取引先: partner, 品目: item, 勘定科目: account_item, セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン): segment_1_tag, セグメント2(法人向け エンタープライズプラン):segment_2_tag, セグメント3(法人向け エンタープライズプラン): segment_3_tag） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **int32** | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **string** | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **int32** | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **adjustment** | **string** | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 | 
 **costAllocation** | **string** | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without）。指定されない場合、配賦仕訳を含む金額が返却されます。 | 
 **approvalFlowStatus** | **string** | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)&lt;br&gt; 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。&lt;br&gt; 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  | 

### Return type

[**TrialPlSectionsResponse**](TrialPlSectionsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrialPlSegment1Tags

> TrialPlSegment1TagsResponse GetTrialPlSegment1Tags(ctx).CompanyId(companyId).Segment1TagIds(segment1TagIds).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).CostAllocation(costAllocation).ApprovalFlowStatus(approvalFlowStatus).Execute()

損益計算書(セグメント1比較)の取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyId := int32(56) // int32 | 事業所ID
    segment1TagIds := "segment1TagIds_example" // string | 出力するセグメント1タグの指定（半角数字のidを半角カンマ区切りスペースなしで指定してください。0を指定すると、未選択のセグメントで比較できます）
    fiscalYear := int32(56) // int32 | 会計年度 (optional)
    startMonth := int32(56) // int32 | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 (optional)
    endMonth := int32(56) // int32 | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 (optional)
    startDate := "startDate_example" // string | 発生日で絞込：開始日(yyyy-mm-dd) (optional)
    endDate := "endDate_example" // string | 発生日で絞込：終了日(yyyy-mm-dd) (optional)
    accountItemDisplayType := "accountItemDisplayType_example" // string | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 (optional)
    breakdownDisplayType := "breakdownDisplayType_example" // string | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます (optional)
    partnerId := int32(56) // int32 | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） (optional)
    partnerCode := "partnerCode_example" // string | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） (optional)
    itemId := int32(56) // int32 | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） (optional)
    sectionId := int32(56) // int32 | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） (optional)
    adjustment := "adjustment_example" // string | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 (optional)
    costAllocation := "costAllocation_example" // string | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without）。指定されない場合、配賦仕訳を含む金額が返却されます。 (optional)
    approvalFlowStatus := "approvalFlowStatus_example" // string | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)<br> 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。<br> 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrialBalanceApi.GetTrialPlSegment1Tags(context.Background()).CompanyId(companyId).Segment1TagIds(segment1TagIds).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).CostAllocation(costAllocation).ApprovalFlowStatus(approvalFlowStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrialBalanceApi.GetTrialPlSegment1Tags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrialPlSegment1Tags`: TrialPlSegment1TagsResponse
    fmt.Fprintf(os.Stdout, "Response from `TrialBalanceApi.GetTrialPlSegment1Tags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrialPlSegment1TagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **segment1TagIds** | **string** | 出力するセグメント1タグの指定（半角数字のidを半角カンマ区切りスペースなしで指定してください。0を指定すると、未選択のセグメントで比較できます） | 
 **fiscalYear** | **int32** | 会計年度 | 
 **startMonth** | **int32** | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 | 
 **endMonth** | **int32** | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 | 
 **startDate** | **string** | 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **string** | 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **string** | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 | 
 **breakdownDisplayType** | **string** | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **int32** | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **string** | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **int32** | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **sectionId** | **int32** | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） | 
 **adjustment** | **string** | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 | 
 **costAllocation** | **string** | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without）。指定されない場合、配賦仕訳を含む金額が返却されます。 | 
 **approvalFlowStatus** | **string** | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)&lt;br&gt; 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。&lt;br&gt; 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  | 

### Return type

[**TrialPlSegment1TagsResponse**](TrialPlSegment1TagsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrialPlSegment2Tags

> TrialPlSegment2TagsResponse GetTrialPlSegment2Tags(ctx).CompanyId(companyId).Segment2TagIds(segment2TagIds).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).CostAllocation(costAllocation).ApprovalFlowStatus(approvalFlowStatus).Execute()

損益計算書(セグメント2比較)の取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyId := int32(56) // int32 | 事業所ID
    segment2TagIds := "segment2TagIds_example" // string | 出力するセグメント2タグの指定（半角数字のidを半角カンマ区切りスペースなしで指定してください。0を指定すると、未選択のセグメントで比較できます）
    fiscalYear := int32(56) // int32 | 会計年度 (optional)
    startMonth := int32(56) // int32 | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 (optional)
    endMonth := int32(56) // int32 | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 (optional)
    startDate := "startDate_example" // string | 発生日で絞込：開始日(yyyy-mm-dd) (optional)
    endDate := "endDate_example" // string | 発生日で絞込：終了日(yyyy-mm-dd) (optional)
    accountItemDisplayType := "accountItemDisplayType_example" // string | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 (optional)
    breakdownDisplayType := "breakdownDisplayType_example" // string | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます (optional)
    partnerId := int32(56) // int32 | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） (optional)
    partnerCode := "partnerCode_example" // string | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） (optional)
    itemId := int32(56) // int32 | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） (optional)
    sectionId := int32(56) // int32 | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） (optional)
    adjustment := "adjustment_example" // string | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 (optional)
    costAllocation := "costAllocation_example" // string | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without）。指定されない場合、配賦仕訳を含む金額が返却されます。 (optional)
    approvalFlowStatus := "approvalFlowStatus_example" // string | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)<br> 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。<br> 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrialBalanceApi.GetTrialPlSegment2Tags(context.Background()).CompanyId(companyId).Segment2TagIds(segment2TagIds).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).CostAllocation(costAllocation).ApprovalFlowStatus(approvalFlowStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrialBalanceApi.GetTrialPlSegment2Tags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrialPlSegment2Tags`: TrialPlSegment2TagsResponse
    fmt.Fprintf(os.Stdout, "Response from `TrialBalanceApi.GetTrialPlSegment2Tags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrialPlSegment2TagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **segment2TagIds** | **string** | 出力するセグメント2タグの指定（半角数字のidを半角カンマ区切りスペースなしで指定してください。0を指定すると、未選択のセグメントで比較できます） | 
 **fiscalYear** | **int32** | 会計年度 | 
 **startMonth** | **int32** | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 | 
 **endMonth** | **int32** | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 | 
 **startDate** | **string** | 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **string** | 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **string** | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 | 
 **breakdownDisplayType** | **string** | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **int32** | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **string** | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **int32** | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **sectionId** | **int32** | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） | 
 **adjustment** | **string** | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 | 
 **costAllocation** | **string** | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without）。指定されない場合、配賦仕訳を含む金額が返却されます。 | 
 **approvalFlowStatus** | **string** | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)&lt;br&gt; 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。&lt;br&gt; 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  | 

### Return type

[**TrialPlSegment2TagsResponse**](TrialPlSegment2TagsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrialPlSegment3Tags

> TrialPlSegment3TagsResponse GetTrialPlSegment3Tags(ctx).CompanyId(companyId).Segment3TagIds(segment3TagIds).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).CostAllocation(costAllocation).ApprovalFlowStatus(approvalFlowStatus).Execute()

損益計算書(セグメント3比較)の取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyId := int32(56) // int32 | 事業所ID
    segment3TagIds := "segment3TagIds_example" // string | 出力するセグメント3タグの指定（半角数字のidを半角カンマ区切りスペースなしで指定してください。0を指定すると、未選択のセグメントで比較できます）
    fiscalYear := int32(56) // int32 | 会計年度 (optional)
    startMonth := int32(56) // int32 | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 (optional)
    endMonth := int32(56) // int32 | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 (optional)
    startDate := "startDate_example" // string | 発生日で絞込：開始日(yyyy-mm-dd) (optional)
    endDate := "endDate_example" // string | 発生日で絞込：終了日(yyyy-mm-dd) (optional)
    accountItemDisplayType := "accountItemDisplayType_example" // string | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 (optional)
    breakdownDisplayType := "breakdownDisplayType_example" // string | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます (optional)
    partnerId := int32(56) // int32 | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） (optional)
    partnerCode := "partnerCode_example" // string | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） (optional)
    itemId := int32(56) // int32 | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） (optional)
    sectionId := int32(56) // int32 | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） (optional)
    adjustment := "adjustment_example" // string | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 (optional)
    costAllocation := "costAllocation_example" // string | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without）。指定されない場合、配賦仕訳を含む金額が返却されます。 (optional)
    approvalFlowStatus := "approvalFlowStatus_example" // string | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)<br> 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。<br> 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrialBalanceApi.GetTrialPlSegment3Tags(context.Background()).CompanyId(companyId).Segment3TagIds(segment3TagIds).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).CostAllocation(costAllocation).ApprovalFlowStatus(approvalFlowStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrialBalanceApi.GetTrialPlSegment3Tags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrialPlSegment3Tags`: TrialPlSegment3TagsResponse
    fmt.Fprintf(os.Stdout, "Response from `TrialBalanceApi.GetTrialPlSegment3Tags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrialPlSegment3TagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **segment3TagIds** | **string** | 出力するセグメント3タグの指定（半角数字のidを半角カンマ区切りスペースなしで指定してください。0を指定すると、未選択のセグメントで比較できます） | 
 **fiscalYear** | **int32** | 会計年度 | 
 **startMonth** | **int32** | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 | 
 **endMonth** | **int32** | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 | 
 **startDate** | **string** | 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **string** | 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **string** | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 | 
 **breakdownDisplayType** | **string** | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **int32** | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **string** | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **int32** | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **sectionId** | **int32** | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） | 
 **adjustment** | **string** | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 | 
 **costAllocation** | **string** | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without）。指定されない場合、配賦仕訳を含む金額が返却されます。 | 
 **approvalFlowStatus** | **string** | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)&lt;br&gt; 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。&lt;br&gt; 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  | 

### Return type

[**TrialPlSegment3TagsResponse**](TrialPlSegment3TagsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrialPlThreeYears

> TrialPlThreeYearsResponse GetTrialPlThreeYears(ctx).CompanyId(companyId).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).CostAllocation(costAllocation).ApprovalFlowStatus(approvalFlowStatus).Execute()

損益計算書(３期間比較)の取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyId := int32(56) // int32 | 事業所ID
    fiscalYear := int32(56) // int32 | 会計年度 (optional)
    startMonth := int32(56) // int32 | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 (optional)
    endMonth := int32(56) // int32 | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 (optional)
    startDate := "startDate_example" // string | 発生日で絞込：開始日(yyyy-mm-dd) (optional)
    endDate := "endDate_example" // string | 発生日で絞込：終了日(yyyy-mm-dd) (optional)
    accountItemDisplayType := "accountItemDisplayType_example" // string | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 (optional)
    breakdownDisplayType := "breakdownDisplayType_example" // string | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item, セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン): segment_1_tag, セグメント2(法人向け エンタープライズプラン):segment_2_tag, セグメント3(法人向け エンタープライズプラン): segment_3_tag） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます (optional)
    partnerId := int32(56) // int32 | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） (optional)
    partnerCode := "partnerCode_example" // string | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） (optional)
    itemId := int32(56) // int32 | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） (optional)
    sectionId := int32(56) // int32 | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） (optional)
    adjustment := "adjustment_example" // string | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 (optional)
    costAllocation := "costAllocation_example" // string | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without）。指定されない場合、配賦仕訳を含む金額が返却されます。 (optional)
    approvalFlowStatus := "approvalFlowStatus_example" // string | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)<br> 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。<br> 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrialBalanceApi.GetTrialPlThreeYears(context.Background()).CompanyId(companyId).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).CostAllocation(costAllocation).ApprovalFlowStatus(approvalFlowStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrialBalanceApi.GetTrialPlThreeYears``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrialPlThreeYears`: TrialPlThreeYearsResponse
    fmt.Fprintf(os.Stdout, "Response from `TrialBalanceApi.GetTrialPlThreeYears`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrialPlThreeYearsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **fiscalYear** | **int32** | 会計年度 | 
 **startMonth** | **int32** | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 | 
 **endMonth** | **int32** | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 | 
 **startDate** | **string** | 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **string** | 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **string** | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 | 
 **breakdownDisplayType** | **string** | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item, セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン): segment_1_tag, セグメント2(法人向け エンタープライズプラン):segment_2_tag, セグメント3(法人向け エンタープライズプラン): segment_3_tag） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **int32** | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **string** | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **int32** | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **sectionId** | **int32** | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） | 
 **adjustment** | **string** | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 | 
 **costAllocation** | **string** | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without）。指定されない場合、配賦仕訳を含む金額が返却されます。 | 
 **approvalFlowStatus** | **string** | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)&lt;br&gt; 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。&lt;br&gt; 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  | 

### Return type

[**TrialPlThreeYearsResponse**](TrialPlThreeYearsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrialPlTwoYears

> TrialPlTwoYearsResponse GetTrialPlTwoYears(ctx).CompanyId(companyId).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).CostAllocation(costAllocation).ApprovalFlowStatus(approvalFlowStatus).Execute()

損益計算書(前年比較)の取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyId := int32(56) // int32 | 事業所ID
    fiscalYear := int32(56) // int32 | 会計年度 (optional)
    startMonth := int32(56) // int32 | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 (optional)
    endMonth := int32(56) // int32 | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 (optional)
    startDate := "startDate_example" // string | 発生日で絞込：開始日(yyyy-mm-dd) (optional)
    endDate := "endDate_example" // string | 発生日で絞込：終了日(yyyy-mm-dd) (optional)
    accountItemDisplayType := "accountItemDisplayType_example" // string | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 (optional)
    breakdownDisplayType := "breakdownDisplayType_example" // string | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item, セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン): segment_1_tag, セグメント2(法人向け エンタープライズプラン):segment_2_tag, セグメント3(法人向け エンタープライズプラン): segment_3_tag） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます (optional)
    partnerId := int32(56) // int32 | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） (optional)
    partnerCode := "partnerCode_example" // string | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） (optional)
    itemId := int32(56) // int32 | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） (optional)
    sectionId := int32(56) // int32 | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） (optional)
    adjustment := "adjustment_example" // string | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 (optional)
    costAllocation := "costAllocation_example" // string | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without）。指定されない場合、配賦仕訳を含む金額が返却されます。 (optional)
    approvalFlowStatus := "approvalFlowStatus_example" // string | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)<br> 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。<br> 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrialBalanceApi.GetTrialPlTwoYears(context.Background()).CompanyId(companyId).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).CostAllocation(costAllocation).ApprovalFlowStatus(approvalFlowStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrialBalanceApi.GetTrialPlTwoYears``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrialPlTwoYears`: TrialPlTwoYearsResponse
    fmt.Fprintf(os.Stdout, "Response from `TrialBalanceApi.GetTrialPlTwoYears`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrialPlTwoYearsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **fiscalYear** | **int32** | 会計年度 | 
 **startMonth** | **int32** | 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。 | 
 **endMonth** | **int32** | 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。 | 
 **startDate** | **string** | 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **string** | 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **string** | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。 | 
 **breakdownDisplayType** | **string** | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item, セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン): segment_1_tag, セグメント2(法人向け エンタープライズプラン):segment_2_tag, セグメント3(法人向け エンタープライズプラン): segment_3_tag） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **int32** | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **string** | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **int32** | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **sectionId** | **int32** | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） | 
 **adjustment** | **string** | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。 | 
 **costAllocation** | **string** | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without）。指定されない場合、配賦仕訳を含む金額が返却されます。 | 
 **approvalFlowStatus** | **string** | 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)&lt;br&gt; 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。&lt;br&gt; 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。  | 

### Return type

[**TrialPlTwoYearsResponse**](TrialPlTwoYearsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

