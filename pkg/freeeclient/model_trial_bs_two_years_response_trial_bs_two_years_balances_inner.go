/*
freee API

 <h1 id=\"freee_api\">freee API</h1> <hr /> <h2 id=\"start_guide\">スタートガイド</h2>  <p>freee API開発がはじめての方は<a href=\"https://developer.freee.co.jp/getting-started\">freee API スタートガイド</a>を参照してください。</p>  <hr /> <h2 id=\"specification\">仕様</h2>  <h3 id=\"api_endpoint\">APIエンドポイント</h3>  <p>https://api.freee.co.jp/ (httpsのみ)</p>  <h3 id=\"about_authorize\">認証について</h3> <p>OAuth2.0を利用します。詳細は<a href=\"https://developer.freee.co.jp/docs\" target=\"_blank\">ドキュメントの認証</a>パートを参照してください。</p>  <h3 id=\"data_format\">データフォーマット</h3>  <p>リクエスト、レスポンスともにJSON形式をサポートしていますが、詳細は、API毎の説明欄（application/jsonなど）を確認してください。</p>  <h3 id=\"compatibility\">後方互換性ありの変更</h3>  <p>freeeでは、APIを改善していくために以下のような変更は後方互換性ありとして通知なく変更を入れることがあります。アプリケーション実装者は以下を踏まえて開発を行ってください。</p>  <ul> <li>新しいAPIリソース・エンドポイントの追加</li> <li>既存のAPIに対して必須ではない新しいリクエストパラメータの追加</li> <li>既存のAPIレスポンスに対する新しいプロパティの追加</li> <li>既存のAPIレスポンスに対するプロパティの順番の入れ変え</li> <li>keyとなっているidやcodeの長さの変更（長くする）</li> </ul>  <h3 id=\"common_response_header\">共通レスポンスヘッダー</h3>  <p>すべてのAPIのレスポンスには以下のHTTPヘッダーが含まれます。</p>  <ul> <li> <p>X-Freee-Request-ID</p> <ul> <li>各リクエスト毎に発行されるID</li> </ul> </li> </ul>  <h3 id=\"common_error_response\">共通エラーレスポンス</h3>  <ul> <li> <p>ステータスコードはレスポンス内のJSONに含まれる他、HTTPヘッダにも含まれる</p> </li> <li> <p>一部のエラーレスポンスにはエラーコードが含まれます。<br>詳細は、<a href=\"https://developer.freee.co.jp/tips/faq/40x-checkpoint\">HTTPステータスコード400台エラー時のチェックポイント</a>を参照してください</p> </li> <p>type</p>  <ul> <li>status : HTTPステータスコードの説明</li>  <li>validation : エラーの詳細の説明（開発者向け）</li> </ul> </li> </ul>  <p>レスポンスの例</p>  <pre><code>  {     &quot;status_code&quot; : 400,     &quot;errors&quot; : [       {         &quot;type&quot; : &quot;status&quot;,         &quot;messages&quot; : [&quot;不正なリクエストです。&quot;]       },       {         &quot;type&quot; : &quot;validation&quot;,         &quot;messages&quot; : [&quot;Date は不正な日付フォーマットです。入力例：2019-12-17&quot;]       }     ]   }</code></pre>  </br>  <h3 id=\"api_rate_limit\">API使用制限</h3>    <p>freeeは一定期間に過度のアクセスを検知した場合、APIアクセスをコントロールする場合があります。</p>   <p>その際のhttp status codeは403となります。制限がかかってから10分程度が過ぎると再度使用することができるようになります。</p>  <h4 id=\"reports_api_endpoint\">/reportsと/receipts/{id}/downloadエンドポイント</h4>  <p>freeeはエンドポイント毎に一定頻度以上のアクセスを検知した場合、APIアクセスをコントロールする場合があります。その際のhttp status codeは429（too many requests）となります。</p>  <ul>   <li>/reports:1秒に10回まで</li>   <li>/receipts/{id}/download:1秒に3回まで</li> </ul>  <p>レスポンスボディのmetaプロパティに以下を含めます。</p>  <ul>   <li>設定されている上限値</li>   <li>上限に達するまでの使用可能回数</li>   <li>（上限値に達した場合）使用回数がリセットされる時刻</li> </ul>  <h3 id=\"plan_api_rate_limit\">プラン別のAPI Rate Limit</h3>   <table border=\"1\">     <tbody>       <tr>         <th style=\"padding: 10px\"><strong>freee会計プラン名</strong></th>         <th style=\"padding: 10px\"><strong>事業所とアプリケーション毎に1日でのAPIコール数</strong></th>       </tr>       <tr>         <td style=\"padding: 10px\">エンタープライズ</td>         <td style=\"padding: 10px\">10,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">プロフェッショナル</td>         <td style=\"padding: 10px\">5,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">ベーシック</td>         <td style=\"padding: 10px\">3,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">ミニマム</td>         <td style=\"padding: 10px\">3,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">上記以外</td>         <td style=\"padding: 10px\">3,000</td>       </tr>     </tbody>   </table>  <h3 id=\"webhook\">Webhookについて</h3>  <p>詳細は<a href=\"https://developer.freee.co.jp/docs/accounting/webhook\" target=\"_blank\">会計Webhook概要</a>を参照してください。</p>  <hr /> <h2 id=\"contact\">連絡先</h2>  <p>ご不明点、ご要望等は <a href=\"https://support.freee.co.jp/hc/ja/requests/new\">freee サポートデスクへのお問い合わせフォーム</a> からご連絡ください。</p> <hr />&copy; Since 2013 freee K.K.

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner struct for TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner
type TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner struct {
	// 勘定科目カテゴリー名
	AccountCategoryName *string `json:"account_category_name,omitempty"`
	// 決算書表示名(account_item_display_type:group指定時に決算書表示名の時のみ含まれる)
	AccountGroupName *string `json:"account_group_name,omitempty"`
	// 勘定科目ID(勘定科目の時のみ含まれる)
	AccountItemId *int32 `json:"account_item_id,omitempty"`
	// 勘定科目名(勘定科目の時のみ含まれる)
	AccountItemName *string `json:"account_item_name,omitempty"`
	// 期末残高
	ClosingBalance *int32 `json:"closing_balance,omitempty"`
	// 階層レベル
	HierarchyLevel *int32 `json:"hierarchy_level,omitempty"`
	// breakdown_display_type:item, account_item_display_type:account_item指定時のみ含まれる
	Items []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerItemsInner `json:"items,omitempty"`
	// 前年度期末残高
	LastYearClosingBalance *int32 `json:"last_year_closing_balance,omitempty"`
	// 上位勘定科目カテゴリー名(勘定科目カテゴリーの時のみ、上層が存在する場合含まれる)
	ParentAccountCategoryName *string `json:"parent_account_category_name,omitempty"`
	// breakdown_display_type:partner, account_item_display_type:account_item指定時のみ含まれる
	Partners []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerPartnersInner `json:"partners,omitempty"`
	// breakdown_display_type:section, account_item_display_type:account_item指定時のみ含まれる
	Sections []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSectionsInner `json:"sections,omitempty"`
	// breakdown_display_type:segment_1_tag, account_item_display_type:account_item指定時のみ含まれる
	Segment1Tags []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment1TagsInner `json:"segment_1_tags,omitempty"`
	// breakdown_display_type:segment_2_tag, account_item_display_type:account_item指定時のみ含まれる
	Segment2Tags []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment2TagsInner `json:"segment_2_tags,omitempty"`
	// breakdown_display_type:segment_3_tag, account_item_display_type:account_item指定時のみ含まれる
	Segment3Tags []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment3TagsInner `json:"segment_3_tags,omitempty"`
	// 合計行(勘定科目カテゴリーの時のみ含まれる)
	TotalLine *bool `json:"total_line,omitempty"`
	// 前年比
	YearOnYear *float32 `json:"year_on_year,omitempty"`
}

// NewTrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner instantiates a new TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner() *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner {
	this := TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner{}
	return &this
}

// NewTrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerWithDefaults instantiates a new TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerWithDefaults() *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner {
	this := TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner{}
	return &this
}

// GetAccountCategoryName returns the AccountCategoryName field value if set, zero value otherwise.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetAccountCategoryName() string {
	if o == nil || o.AccountCategoryName == nil {
		var ret string
		return ret
	}
	return *o.AccountCategoryName
}

// GetAccountCategoryNameOk returns a tuple with the AccountCategoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetAccountCategoryNameOk() (*string, bool) {
	if o == nil || o.AccountCategoryName == nil {
		return nil, false
	}
	return o.AccountCategoryName, true
}

// HasAccountCategoryName returns a boolean if a field has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasAccountCategoryName() bool {
	if o != nil && o.AccountCategoryName != nil {
		return true
	}

	return false
}

// SetAccountCategoryName gets a reference to the given string and assigns it to the AccountCategoryName field.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetAccountCategoryName(v string) {
	o.AccountCategoryName = &v
}

// GetAccountGroupName returns the AccountGroupName field value if set, zero value otherwise.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetAccountGroupName() string {
	if o == nil || o.AccountGroupName == nil {
		var ret string
		return ret
	}
	return *o.AccountGroupName
}

// GetAccountGroupNameOk returns a tuple with the AccountGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetAccountGroupNameOk() (*string, bool) {
	if o == nil || o.AccountGroupName == nil {
		return nil, false
	}
	return o.AccountGroupName, true
}

// HasAccountGroupName returns a boolean if a field has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasAccountGroupName() bool {
	if o != nil && o.AccountGroupName != nil {
		return true
	}

	return false
}

// SetAccountGroupName gets a reference to the given string and assigns it to the AccountGroupName field.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetAccountGroupName(v string) {
	o.AccountGroupName = &v
}

// GetAccountItemId returns the AccountItemId field value if set, zero value otherwise.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetAccountItemId() int32 {
	if o == nil || o.AccountItemId == nil {
		var ret int32
		return ret
	}
	return *o.AccountItemId
}

// GetAccountItemIdOk returns a tuple with the AccountItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetAccountItemIdOk() (*int32, bool) {
	if o == nil || o.AccountItemId == nil {
		return nil, false
	}
	return o.AccountItemId, true
}

// HasAccountItemId returns a boolean if a field has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasAccountItemId() bool {
	if o != nil && o.AccountItemId != nil {
		return true
	}

	return false
}

// SetAccountItemId gets a reference to the given int32 and assigns it to the AccountItemId field.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetAccountItemId(v int32) {
	o.AccountItemId = &v
}

// GetAccountItemName returns the AccountItemName field value if set, zero value otherwise.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetAccountItemName() string {
	if o == nil || o.AccountItemName == nil {
		var ret string
		return ret
	}
	return *o.AccountItemName
}

// GetAccountItemNameOk returns a tuple with the AccountItemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetAccountItemNameOk() (*string, bool) {
	if o == nil || o.AccountItemName == nil {
		return nil, false
	}
	return o.AccountItemName, true
}

// HasAccountItemName returns a boolean if a field has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasAccountItemName() bool {
	if o != nil && o.AccountItemName != nil {
		return true
	}

	return false
}

// SetAccountItemName gets a reference to the given string and assigns it to the AccountItemName field.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetAccountItemName(v string) {
	o.AccountItemName = &v
}

// GetClosingBalance returns the ClosingBalance field value if set, zero value otherwise.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetClosingBalance() int32 {
	if o == nil || o.ClosingBalance == nil {
		var ret int32
		return ret
	}
	return *o.ClosingBalance
}

// GetClosingBalanceOk returns a tuple with the ClosingBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetClosingBalanceOk() (*int32, bool) {
	if o == nil || o.ClosingBalance == nil {
		return nil, false
	}
	return o.ClosingBalance, true
}

// HasClosingBalance returns a boolean if a field has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasClosingBalance() bool {
	if o != nil && o.ClosingBalance != nil {
		return true
	}

	return false
}

// SetClosingBalance gets a reference to the given int32 and assigns it to the ClosingBalance field.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetClosingBalance(v int32) {
	o.ClosingBalance = &v
}

// GetHierarchyLevel returns the HierarchyLevel field value if set, zero value otherwise.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetHierarchyLevel() int32 {
	if o == nil || o.HierarchyLevel == nil {
		var ret int32
		return ret
	}
	return *o.HierarchyLevel
}

// GetHierarchyLevelOk returns a tuple with the HierarchyLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetHierarchyLevelOk() (*int32, bool) {
	if o == nil || o.HierarchyLevel == nil {
		return nil, false
	}
	return o.HierarchyLevel, true
}

// HasHierarchyLevel returns a boolean if a field has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasHierarchyLevel() bool {
	if o != nil && o.HierarchyLevel != nil {
		return true
	}

	return false
}

// SetHierarchyLevel gets a reference to the given int32 and assigns it to the HierarchyLevel field.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetHierarchyLevel(v int32) {
	o.HierarchyLevel = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetItems() []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerItemsInner {
	if o == nil || o.Items == nil {
		var ret []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerItemsInner
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetItemsOk() ([]TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerItemsInner, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerItemsInner and assigns it to the Items field.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetItems(v []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerItemsInner) {
	o.Items = v
}

// GetLastYearClosingBalance returns the LastYearClosingBalance field value if set, zero value otherwise.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetLastYearClosingBalance() int32 {
	if o == nil || o.LastYearClosingBalance == nil {
		var ret int32
		return ret
	}
	return *o.LastYearClosingBalance
}

// GetLastYearClosingBalanceOk returns a tuple with the LastYearClosingBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetLastYearClosingBalanceOk() (*int32, bool) {
	if o == nil || o.LastYearClosingBalance == nil {
		return nil, false
	}
	return o.LastYearClosingBalance, true
}

// HasLastYearClosingBalance returns a boolean if a field has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasLastYearClosingBalance() bool {
	if o != nil && o.LastYearClosingBalance != nil {
		return true
	}

	return false
}

// SetLastYearClosingBalance gets a reference to the given int32 and assigns it to the LastYearClosingBalance field.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetLastYearClosingBalance(v int32) {
	o.LastYearClosingBalance = &v
}

// GetParentAccountCategoryName returns the ParentAccountCategoryName field value if set, zero value otherwise.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetParentAccountCategoryName() string {
	if o == nil || o.ParentAccountCategoryName == nil {
		var ret string
		return ret
	}
	return *o.ParentAccountCategoryName
}

// GetParentAccountCategoryNameOk returns a tuple with the ParentAccountCategoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetParentAccountCategoryNameOk() (*string, bool) {
	if o == nil || o.ParentAccountCategoryName == nil {
		return nil, false
	}
	return o.ParentAccountCategoryName, true
}

// HasParentAccountCategoryName returns a boolean if a field has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasParentAccountCategoryName() bool {
	if o != nil && o.ParentAccountCategoryName != nil {
		return true
	}

	return false
}

// SetParentAccountCategoryName gets a reference to the given string and assigns it to the ParentAccountCategoryName field.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetParentAccountCategoryName(v string) {
	o.ParentAccountCategoryName = &v
}

// GetPartners returns the Partners field value if set, zero value otherwise.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetPartners() []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerPartnersInner {
	if o == nil || o.Partners == nil {
		var ret []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerPartnersInner
		return ret
	}
	return o.Partners
}

// GetPartnersOk returns a tuple with the Partners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetPartnersOk() ([]TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerPartnersInner, bool) {
	if o == nil || o.Partners == nil {
		return nil, false
	}
	return o.Partners, true
}

// HasPartners returns a boolean if a field has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasPartners() bool {
	if o != nil && o.Partners != nil {
		return true
	}

	return false
}

// SetPartners gets a reference to the given []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerPartnersInner and assigns it to the Partners field.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetPartners(v []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerPartnersInner) {
	o.Partners = v
}

// GetSections returns the Sections field value if set, zero value otherwise.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetSections() []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSectionsInner {
	if o == nil || o.Sections == nil {
		var ret []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSectionsInner
		return ret
	}
	return o.Sections
}

// GetSectionsOk returns a tuple with the Sections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetSectionsOk() ([]TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSectionsInner, bool) {
	if o == nil || o.Sections == nil {
		return nil, false
	}
	return o.Sections, true
}

// HasSections returns a boolean if a field has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasSections() bool {
	if o != nil && o.Sections != nil {
		return true
	}

	return false
}

// SetSections gets a reference to the given []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSectionsInner and assigns it to the Sections field.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetSections(v []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSectionsInner) {
	o.Sections = v
}

// GetSegment1Tags returns the Segment1Tags field value if set, zero value otherwise.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetSegment1Tags() []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment1TagsInner {
	if o == nil || o.Segment1Tags == nil {
		var ret []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment1TagsInner
		return ret
	}
	return o.Segment1Tags
}

// GetSegment1TagsOk returns a tuple with the Segment1Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetSegment1TagsOk() ([]TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment1TagsInner, bool) {
	if o == nil || o.Segment1Tags == nil {
		return nil, false
	}
	return o.Segment1Tags, true
}

// HasSegment1Tags returns a boolean if a field has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasSegment1Tags() bool {
	if o != nil && o.Segment1Tags != nil {
		return true
	}

	return false
}

// SetSegment1Tags gets a reference to the given []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment1TagsInner and assigns it to the Segment1Tags field.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetSegment1Tags(v []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment1TagsInner) {
	o.Segment1Tags = v
}

// GetSegment2Tags returns the Segment2Tags field value if set, zero value otherwise.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetSegment2Tags() []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment2TagsInner {
	if o == nil || o.Segment2Tags == nil {
		var ret []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment2TagsInner
		return ret
	}
	return o.Segment2Tags
}

// GetSegment2TagsOk returns a tuple with the Segment2Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetSegment2TagsOk() ([]TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment2TagsInner, bool) {
	if o == nil || o.Segment2Tags == nil {
		return nil, false
	}
	return o.Segment2Tags, true
}

// HasSegment2Tags returns a boolean if a field has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasSegment2Tags() bool {
	if o != nil && o.Segment2Tags != nil {
		return true
	}

	return false
}

// SetSegment2Tags gets a reference to the given []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment2TagsInner and assigns it to the Segment2Tags field.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetSegment2Tags(v []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment2TagsInner) {
	o.Segment2Tags = v
}

// GetSegment3Tags returns the Segment3Tags field value if set, zero value otherwise.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetSegment3Tags() []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment3TagsInner {
	if o == nil || o.Segment3Tags == nil {
		var ret []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment3TagsInner
		return ret
	}
	return o.Segment3Tags
}

// GetSegment3TagsOk returns a tuple with the Segment3Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetSegment3TagsOk() ([]TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment3TagsInner, bool) {
	if o == nil || o.Segment3Tags == nil {
		return nil, false
	}
	return o.Segment3Tags, true
}

// HasSegment3Tags returns a boolean if a field has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasSegment3Tags() bool {
	if o != nil && o.Segment3Tags != nil {
		return true
	}

	return false
}

// SetSegment3Tags gets a reference to the given []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment3TagsInner and assigns it to the Segment3Tags field.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetSegment3Tags(v []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment3TagsInner) {
	o.Segment3Tags = v
}

// GetTotalLine returns the TotalLine field value if set, zero value otherwise.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetTotalLine() bool {
	if o == nil || o.TotalLine == nil {
		var ret bool
		return ret
	}
	return *o.TotalLine
}

// GetTotalLineOk returns a tuple with the TotalLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetTotalLineOk() (*bool, bool) {
	if o == nil || o.TotalLine == nil {
		return nil, false
	}
	return o.TotalLine, true
}

// HasTotalLine returns a boolean if a field has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasTotalLine() bool {
	if o != nil && o.TotalLine != nil {
		return true
	}

	return false
}

// SetTotalLine gets a reference to the given bool and assigns it to the TotalLine field.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetTotalLine(v bool) {
	o.TotalLine = &v
}

// GetYearOnYear returns the YearOnYear field value if set, zero value otherwise.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetYearOnYear() float32 {
	if o == nil || o.YearOnYear == nil {
		var ret float32
		return ret
	}
	return *o.YearOnYear
}

// GetYearOnYearOk returns a tuple with the YearOnYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetYearOnYearOk() (*float32, bool) {
	if o == nil || o.YearOnYear == nil {
		return nil, false
	}
	return o.YearOnYear, true
}

// HasYearOnYear returns a boolean if a field has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasYearOnYear() bool {
	if o != nil && o.YearOnYear != nil {
		return true
	}

	return false
}

// SetYearOnYear gets a reference to the given float32 and assigns it to the YearOnYear field.
func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetYearOnYear(v float32) {
	o.YearOnYear = &v
}

func (o TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountCategoryName != nil {
		toSerialize["account_category_name"] = o.AccountCategoryName
	}
	if o.AccountGroupName != nil {
		toSerialize["account_group_name"] = o.AccountGroupName
	}
	if o.AccountItemId != nil {
		toSerialize["account_item_id"] = o.AccountItemId
	}
	if o.AccountItemName != nil {
		toSerialize["account_item_name"] = o.AccountItemName
	}
	if o.ClosingBalance != nil {
		toSerialize["closing_balance"] = o.ClosingBalance
	}
	if o.HierarchyLevel != nil {
		toSerialize["hierarchy_level"] = o.HierarchyLevel
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.LastYearClosingBalance != nil {
		toSerialize["last_year_closing_balance"] = o.LastYearClosingBalance
	}
	if o.ParentAccountCategoryName != nil {
		toSerialize["parent_account_category_name"] = o.ParentAccountCategoryName
	}
	if o.Partners != nil {
		toSerialize["partners"] = o.Partners
	}
	if o.Sections != nil {
		toSerialize["sections"] = o.Sections
	}
	if o.Segment1Tags != nil {
		toSerialize["segment_1_tags"] = o.Segment1Tags
	}
	if o.Segment2Tags != nil {
		toSerialize["segment_2_tags"] = o.Segment2Tags
	}
	if o.Segment3Tags != nil {
		toSerialize["segment_3_tags"] = o.Segment3Tags
	}
	if o.TotalLine != nil {
		toSerialize["total_line"] = o.TotalLine
	}
	if o.YearOnYear != nil {
		toSerialize["year_on_year"] = o.YearOnYear
	}
	return json.Marshal(toSerialize)
}

type NullableTrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner struct {
	value *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner
	isSet bool
}

func (v NullableTrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) Get() *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner {
	return v.value
}

func (v *NullableTrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) Set(val *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner(val *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) *NullableTrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner {
	return &NullableTrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner{value: val, isSet: true}
}

func (v NullableTrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


