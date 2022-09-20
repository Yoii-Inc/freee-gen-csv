/*
freee API

 <h1 id=\"freee_api\">freee API</h1> <hr /> <h2 id=\"start_guide\">スタートガイド</h2>  <p>freee API開発がはじめての方は<a href=\"https://developer.freee.co.jp/getting-started\">freee API スタートガイド</a>を参照してください。</p>  <hr /> <h2 id=\"specification\">仕様</h2>  <h3 id=\"api_endpoint\">APIエンドポイント</h3>  <p>https://api.freee.co.jp/ (httpsのみ)</p>  <h3 id=\"about_authorize\">認証について</h3> <p>OAuth2.0を利用します。詳細は<a href=\"https://developer.freee.co.jp/docs\" target=\"_blank\">ドキュメントの認証</a>パートを参照してください。</p>  <h3 id=\"data_format\">データフォーマット</h3>  <p>リクエスト、レスポンスともにJSON形式をサポートしていますが、詳細は、API毎の説明欄（application/jsonなど）を確認してください。</p>  <h3 id=\"compatibility\">後方互換性ありの変更</h3>  <p>freeeでは、APIを改善していくために以下のような変更は後方互換性ありとして通知なく変更を入れることがあります。アプリケーション実装者は以下を踏まえて開発を行ってください。</p>  <ul> <li>新しいAPIリソース・エンドポイントの追加</li> <li>既存のAPIに対して必須ではない新しいリクエストパラメータの追加</li> <li>既存のAPIレスポンスに対する新しいプロパティの追加</li> <li>既存のAPIレスポンスに対するプロパティの順番の入れ変え</li> <li>keyとなっているidやcodeの長さの変更（長くする）</li> </ul>  <h3 id=\"common_response_header\">共通レスポンスヘッダー</h3>  <p>すべてのAPIのレスポンスには以下のHTTPヘッダーが含まれます。</p>  <ul> <li> <p>X-Freee-Request-ID</p> <ul> <li>各リクエスト毎に発行されるID</li> </ul> </li> </ul>  <h3 id=\"common_error_response\">共通エラーレスポンス</h3>  <ul> <li> <p>ステータスコードはレスポンス内のJSONに含まれる他、HTTPヘッダにも含まれる</p> </li> <li> <p>一部のエラーレスポンスにはエラーコードが含まれます。<br>詳細は、<a href=\"https://developer.freee.co.jp/tips/faq/40x-checkpoint\">HTTPステータスコード400台エラー時のチェックポイント</a>を参照してください</p> </li> <p>type</p>  <ul> <li>status : HTTPステータスコードの説明</li>  <li>validation : エラーの詳細の説明（開発者向け）</li> </ul> </li> </ul>  <p>レスポンスの例</p>  <pre><code>  {     &quot;status_code&quot; : 400,     &quot;errors&quot; : [       {         &quot;type&quot; : &quot;status&quot;,         &quot;messages&quot; : [&quot;不正なリクエストです。&quot;]       },       {         &quot;type&quot; : &quot;validation&quot;,         &quot;messages&quot; : [&quot;Date は不正な日付フォーマットです。入力例：2019-12-17&quot;]       }     ]   }</code></pre>  </br>  <h3 id=\"api_rate_limit\">API使用制限</h3>    <p>freeeは一定期間に過度のアクセスを検知した場合、APIアクセスをコントロールする場合があります。</p>   <p>その際のhttp status codeは403となります。制限がかかってから10分程度が過ぎると再度使用することができるようになります。</p>  <h4 id=\"reports_api_endpoint\">/reportsと/receipts/{id}/downloadエンドポイント</h4>  <p>freeeはエンドポイント毎に一定頻度以上のアクセスを検知した場合、APIアクセスをコントロールする場合があります。その際のhttp status codeは429（too many requests）となります。</p>  <ul>   <li>/reports:1秒に10回まで</li>   <li>/receipts/{id}/download:1秒に3回まで</li> </ul>  <p>レスポンスボディのmetaプロパティに以下を含めます。</p>  <ul>   <li>設定されている上限値</li>   <li>上限に達するまでの使用可能回数</li>   <li>（上限値に達した場合）使用回数がリセットされる時刻</li> </ul>  <h3 id=\"plan_api_rate_limit\">プラン別のAPI Rate Limit</h3>   <table border=\"1\">     <tbody>       <tr>         <th style=\"padding: 10px\"><strong>freee会計プラン名</strong></th>         <th style=\"padding: 10px\"><strong>事業所とアプリケーション毎に1日でのAPIコール数</strong></th>       </tr>       <tr>         <td style=\"padding: 10px\">エンタープライズ</td>         <td style=\"padding: 10px\">10,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">プロフェッショナル</td>         <td style=\"padding: 10px\">5,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">ベーシック</td>         <td style=\"padding: 10px\">3,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">ミニマム</td>         <td style=\"padding: 10px\">3,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">上記以外</td>         <td style=\"padding: 10px\">3,000</td>       </tr>     </tbody>   </table>  <h3 id=\"webhook\">Webhookについて</h3>  <p>詳細は<a href=\"https://developer.freee.co.jp/docs/accounting/webhook\" target=\"_blank\">会計Webhook概要</a>を参照してください。</p>  <hr /> <h2 id=\"contact\">連絡先</h2>  <p>ご不明点、ご要望等は <a href=\"https://support.freee.co.jp/hc/ja/requests/new\">freee サポートデスクへのお問い合わせフォーム</a> からご連絡ください。</p> <hr />&copy; Since 2013 freee K.K.

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github.com/Yoii-Inc/freee-gen-csv

import (
	"encoding/json"
)

// TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner struct for TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner
type TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner struct {
	// 勘定科目カテゴリー名
	AccountCategoryName *string `json:"account_category_name,omitempty"`
	// 決算書表示名(account_item_display_type:group指定時に決算書表示名の時のみ含まれる)
	AccountGroupName *string `json:"account_group_name,omitempty"`
	// 勘定科目ID(勘定科目の時のみ含まれる)
	AccountItemId *int64 `json:"account_item_id,omitempty"`
	// 勘定科目名(勘定科目の時のみ含まれる)
	AccountItemName *string `json:"account_item_name,omitempty"`
	// 期末残高
	ClosingBalance *int64 `json:"closing_balance,omitempty"`
	// 階層レベル
	HierarchyLevel *int64 `json:"hierarchy_level,omitempty"`
	// 上位勘定科目カテゴリー名(勘定科目カテゴリーの時のみ、上層が存在する場合含まれる)
	ParentAccountCategoryName *string `json:"parent_account_category_name,omitempty"`
	// セグメント1タグ
	Segment1Tags []TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInnerSegment1TagsInner `json:"segment_1_tags,omitempty"`
	// 合計行(勘定科目カテゴリーの時のみ含まれる)
	TotalLine *bool `json:"total_line,omitempty"`
}

// NewTrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner instantiates a new TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner() *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner {
	this := TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner{}
	return &this
}

// NewTrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInnerWithDefaults instantiates a new TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInnerWithDefaults() *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner {
	this := TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner{}
	return &this
}

// GetAccountCategoryName returns the AccountCategoryName field value if set, zero value otherwise.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetAccountCategoryName() string {
	if o == nil || o.AccountCategoryName == nil {
		var ret string
		return ret
	}
	return *o.AccountCategoryName
}

// GetAccountCategoryNameOk returns a tuple with the AccountCategoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetAccountCategoryNameOk() (*string, bool) {
	if o == nil || o.AccountCategoryName == nil {
		return nil, false
	}
	return o.AccountCategoryName, true
}

// HasAccountCategoryName returns a boolean if a field has been set.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) HasAccountCategoryName() bool {
	if o != nil && o.AccountCategoryName != nil {
		return true
	}

	return false
}

// SetAccountCategoryName gets a reference to the given string and assigns it to the AccountCategoryName field.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) SetAccountCategoryName(v string) {
	o.AccountCategoryName = &v
}

// GetAccountGroupName returns the AccountGroupName field value if set, zero value otherwise.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetAccountGroupName() string {
	if o == nil || o.AccountGroupName == nil {
		var ret string
		return ret
	}
	return *o.AccountGroupName
}

// GetAccountGroupNameOk returns a tuple with the AccountGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetAccountGroupNameOk() (*string, bool) {
	if o == nil || o.AccountGroupName == nil {
		return nil, false
	}
	return o.AccountGroupName, true
}

// HasAccountGroupName returns a boolean if a field has been set.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) HasAccountGroupName() bool {
	if o != nil && o.AccountGroupName != nil {
		return true
	}

	return false
}

// SetAccountGroupName gets a reference to the given string and assigns it to the AccountGroupName field.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) SetAccountGroupName(v string) {
	o.AccountGroupName = &v
}

// GetAccountItemId returns the AccountItemId field value if set, zero value otherwise.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetAccountItemId() int64 {
	if o == nil || o.AccountItemId == nil {
		var ret int64
		return ret
	}
	return *o.AccountItemId
}

// GetAccountItemIdOk returns a tuple with the AccountItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetAccountItemIdOk() (*int64, bool) {
	if o == nil || o.AccountItemId == nil {
		return nil, false
	}
	return o.AccountItemId, true
}

// HasAccountItemId returns a boolean if a field has been set.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) HasAccountItemId() bool {
	if o != nil && o.AccountItemId != nil {
		return true
	}

	return false
}

// SetAccountItemId gets a reference to the given int64 and assigns it to the AccountItemId field.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) SetAccountItemId(v int64) {
	o.AccountItemId = &v
}

// GetAccountItemName returns the AccountItemName field value if set, zero value otherwise.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetAccountItemName() string {
	if o == nil || o.AccountItemName == nil {
		var ret string
		return ret
	}
	return *o.AccountItemName
}

// GetAccountItemNameOk returns a tuple with the AccountItemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetAccountItemNameOk() (*string, bool) {
	if o == nil || o.AccountItemName == nil {
		return nil, false
	}
	return o.AccountItemName, true
}

// HasAccountItemName returns a boolean if a field has been set.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) HasAccountItemName() bool {
	if o != nil && o.AccountItemName != nil {
		return true
	}

	return false
}

// SetAccountItemName gets a reference to the given string and assigns it to the AccountItemName field.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) SetAccountItemName(v string) {
	o.AccountItemName = &v
}

// GetClosingBalance returns the ClosingBalance field value if set, zero value otherwise.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetClosingBalance() int64 {
	if o == nil || o.ClosingBalance == nil {
		var ret int64
		return ret
	}
	return *o.ClosingBalance
}

// GetClosingBalanceOk returns a tuple with the ClosingBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetClosingBalanceOk() (*int64, bool) {
	if o == nil || o.ClosingBalance == nil {
		return nil, false
	}
	return o.ClosingBalance, true
}

// HasClosingBalance returns a boolean if a field has been set.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) HasClosingBalance() bool {
	if o != nil && o.ClosingBalance != nil {
		return true
	}

	return false
}

// SetClosingBalance gets a reference to the given int64 and assigns it to the ClosingBalance field.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) SetClosingBalance(v int64) {
	o.ClosingBalance = &v
}

// GetHierarchyLevel returns the HierarchyLevel field value if set, zero value otherwise.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetHierarchyLevel() int64 {
	if o == nil || o.HierarchyLevel == nil {
		var ret int64
		return ret
	}
	return *o.HierarchyLevel
}

// GetHierarchyLevelOk returns a tuple with the HierarchyLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetHierarchyLevelOk() (*int64, bool) {
	if o == nil || o.HierarchyLevel == nil {
		return nil, false
	}
	return o.HierarchyLevel, true
}

// HasHierarchyLevel returns a boolean if a field has been set.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) HasHierarchyLevel() bool {
	if o != nil && o.HierarchyLevel != nil {
		return true
	}

	return false
}

// SetHierarchyLevel gets a reference to the given int64 and assigns it to the HierarchyLevel field.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) SetHierarchyLevel(v int64) {
	o.HierarchyLevel = &v
}

// GetParentAccountCategoryName returns the ParentAccountCategoryName field value if set, zero value otherwise.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetParentAccountCategoryName() string {
	if o == nil || o.ParentAccountCategoryName == nil {
		var ret string
		return ret
	}
	return *o.ParentAccountCategoryName
}

// GetParentAccountCategoryNameOk returns a tuple with the ParentAccountCategoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetParentAccountCategoryNameOk() (*string, bool) {
	if o == nil || o.ParentAccountCategoryName == nil {
		return nil, false
	}
	return o.ParentAccountCategoryName, true
}

// HasParentAccountCategoryName returns a boolean if a field has been set.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) HasParentAccountCategoryName() bool {
	if o != nil && o.ParentAccountCategoryName != nil {
		return true
	}

	return false
}

// SetParentAccountCategoryName gets a reference to the given string and assigns it to the ParentAccountCategoryName field.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) SetParentAccountCategoryName(v string) {
	o.ParentAccountCategoryName = &v
}

// GetSegment1Tags returns the Segment1Tags field value if set, zero value otherwise.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetSegment1Tags() []TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInnerSegment1TagsInner {
	if o == nil || o.Segment1Tags == nil {
		var ret []TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInnerSegment1TagsInner
		return ret
	}
	return o.Segment1Tags
}

// GetSegment1TagsOk returns a tuple with the Segment1Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetSegment1TagsOk() ([]TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInnerSegment1TagsInner, bool) {
	if o == nil || o.Segment1Tags == nil {
		return nil, false
	}
	return o.Segment1Tags, true
}

// HasSegment1Tags returns a boolean if a field has been set.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) HasSegment1Tags() bool {
	if o != nil && o.Segment1Tags != nil {
		return true
	}

	return false
}

// SetSegment1Tags gets a reference to the given []TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInnerSegment1TagsInner and assigns it to the Segment1Tags field.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) SetSegment1Tags(v []TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInnerSegment1TagsInner) {
	o.Segment1Tags = v
}

// GetTotalLine returns the TotalLine field value if set, zero value otherwise.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetTotalLine() bool {
	if o == nil || o.TotalLine == nil {
		var ret bool
		return ret
	}
	return *o.TotalLine
}

// GetTotalLineOk returns a tuple with the TotalLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetTotalLineOk() (*bool, bool) {
	if o == nil || o.TotalLine == nil {
		return nil, false
	}
	return o.TotalLine, true
}

// HasTotalLine returns a boolean if a field has been set.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) HasTotalLine() bool {
	if o != nil && o.TotalLine != nil {
		return true
	}

	return false
}

// SetTotalLine gets a reference to the given bool and assigns it to the TotalLine field.
func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) SetTotalLine(v bool) {
	o.TotalLine = &v
}

func (o TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) MarshalJSON() ([]byte, error) {
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
	if o.ParentAccountCategoryName != nil {
		toSerialize["parent_account_category_name"] = o.ParentAccountCategoryName
	}
	if o.Segment1Tags != nil {
		toSerialize["segment_1_tags"] = o.Segment1Tags
	}
	if o.TotalLine != nil {
		toSerialize["total_line"] = o.TotalLine
	}
	return json.Marshal(toSerialize)
}

type NullableTrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner struct {
	value *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner
	isSet bool
}

func (v NullableTrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) Get() *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner {
	return v.value
}

func (v *NullableTrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) Set(val *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner(val *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) *NullableTrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner {
	return &NullableTrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner{value: val, isSet: true}
}

func (v NullableTrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


