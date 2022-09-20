/*
freee API

 <h1 id=\"freee_api\">freee API</h1> <hr /> <h2 id=\"start_guide\">スタートガイド</h2>  <p>freee API開発がはじめての方は<a href=\"https://developer.freee.co.jp/getting-started\">freee API スタートガイド</a>を参照してください。</p>  <hr /> <h2 id=\"specification\">仕様</h2>  <h3 id=\"api_endpoint\">APIエンドポイント</h3>  <p>https://api.freee.co.jp/ (httpsのみ)</p>  <h3 id=\"about_authorize\">認証について</h3> <p>OAuth2.0を利用します。詳細は<a href=\"https://developer.freee.co.jp/docs\" target=\"_blank\">ドキュメントの認証</a>パートを参照してください。</p>  <h3 id=\"data_format\">データフォーマット</h3>  <p>リクエスト、レスポンスともにJSON形式をサポートしていますが、詳細は、API毎の説明欄（application/jsonなど）を確認してください。</p>  <h3 id=\"compatibility\">後方互換性ありの変更</h3>  <p>freeeでは、APIを改善していくために以下のような変更は後方互換性ありとして通知なく変更を入れることがあります。アプリケーション実装者は以下を踏まえて開発を行ってください。</p>  <ul> <li>新しいAPIリソース・エンドポイントの追加</li> <li>既存のAPIに対して必須ではない新しいリクエストパラメータの追加</li> <li>既存のAPIレスポンスに対する新しいプロパティの追加</li> <li>既存のAPIレスポンスに対するプロパティの順番の入れ変え</li> <li>keyとなっているidやcodeの長さの変更（長くする）</li> </ul>  <h3 id=\"common_response_header\">共通レスポンスヘッダー</h3>  <p>すべてのAPIのレスポンスには以下のHTTPヘッダーが含まれます。</p>  <ul> <li> <p>X-Freee-Request-ID</p> <ul> <li>各リクエスト毎に発行されるID</li> </ul> </li> </ul>  <h3 id=\"common_error_response\">共通エラーレスポンス</h3>  <ul> <li> <p>ステータスコードはレスポンス内のJSONに含まれる他、HTTPヘッダにも含まれる</p> </li> <li> <p>一部のエラーレスポンスにはエラーコードが含まれます。<br>詳細は、<a href=\"https://developer.freee.co.jp/tips/faq/40x-checkpoint\">HTTPステータスコード400台エラー時のチェックポイント</a>を参照してください</p> </li> <p>type</p>  <ul> <li>status : HTTPステータスコードの説明</li>  <li>validation : エラーの詳細の説明（開発者向け）</li> </ul> </li> </ul>  <p>レスポンスの例</p>  <pre><code>  {     &quot;status_code&quot; : 400,     &quot;errors&quot; : [       {         &quot;type&quot; : &quot;status&quot;,         &quot;messages&quot; : [&quot;不正なリクエストです。&quot;]       },       {         &quot;type&quot; : &quot;validation&quot;,         &quot;messages&quot; : [&quot;Date は不正な日付フォーマットです。入力例：2019-12-17&quot;]       }     ]   }</code></pre>  </br>  <h3 id=\"api_rate_limit\">API使用制限</h3>    <p>freeeは一定期間に過度のアクセスを検知した場合、APIアクセスをコントロールする場合があります。</p>   <p>その際のhttp status codeは403となります。制限がかかってから10分程度が過ぎると再度使用することができるようになります。</p>  <h4 id=\"reports_api_endpoint\">/reportsと/receipts/{id}/downloadエンドポイント</h4>  <p>freeeはエンドポイント毎に一定頻度以上のアクセスを検知した場合、APIアクセスをコントロールする場合があります。その際のhttp status codeは429（too many requests）となります。</p>  <ul>   <li>/reports:1秒に10回まで</li>   <li>/receipts/{id}/download:1秒に3回まで</li> </ul>  <p>レスポンスボディのmetaプロパティに以下を含めます。</p>  <ul>   <li>設定されている上限値</li>   <li>上限に達するまでの使用可能回数</li>   <li>（上限値に達した場合）使用回数がリセットされる時刻</li> </ul>  <h3 id=\"plan_api_rate_limit\">プラン別のAPI Rate Limit</h3>   <table border=\"1\">     <tbody>       <tr>         <th style=\"padding: 10px\"><strong>freee会計プラン名</strong></th>         <th style=\"padding: 10px\"><strong>事業所とアプリケーション毎に1日でのAPIコール数</strong></th>       </tr>       <tr>         <td style=\"padding: 10px\">エンタープライズ</td>         <td style=\"padding: 10px\">10,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">プロフェッショナル</td>         <td style=\"padding: 10px\">5,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">ベーシック</td>         <td style=\"padding: 10px\">3,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">ミニマム</td>         <td style=\"padding: 10px\">3,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">上記以外</td>         <td style=\"padding: 10px\">3,000</td>       </tr>     </tbody>   </table>  <h3 id=\"webhook\">Webhookについて</h3>  <p>詳細は<a href=\"https://developer.freee.co.jp/docs/accounting/webhook\" target=\"_blank\">会計Webhook概要</a>を参照してください。</p>  <hr /> <h2 id=\"contact\">連絡先</h2>  <p>ご不明点、ご要望等は <a href=\"https://support.freee.co.jp/hc/ja/requests/new\">freee サポートデスクへのお問い合わせフォーム</a> からご連絡ください。</p> <hr />&copy; Since 2013 freee K.K.

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package freeeclient

import (
	"encoding/json"
)

// Deal struct for Deal
type Deal struct {
	// 金額
	Amount int64 `json:"amount"`
	// 事業所ID
	CompanyId int64 `json:"company_id"`
	// 取引の明細行
	Details []DealDetailsInner `json:"details,omitempty"`
	// 支払残額
	DueAmount *int64 `json:"due_amount,omitempty"`
	// 支払期日 (yyyy-mm-dd)
	DueDate *string `json:"due_date,omitempty"`
	// 取引ID
	Id int64 `json:"id"`
	// 発生日 (yyyy-mm-dd)
	IssueDate string `json:"issue_date"`
	// 取引先コード
	PartnerCode NullableString `json:"partner_code,omitempty"`
	// 取引先ID
	PartnerId int64 `json:"partner_id"`
	// 取引の支払行
	Payments []DealPaymentsInner `json:"payments,omitempty"`
	// 証憑ファイル（ファイルボックスのファイル）
	Receipts []DealReceiptsInner `json:"receipts,omitempty"`
	// 管理番号
	RefNumber *string `json:"ref_number,omitempty"`
	// 取引の+更新行
	Renews []DealRenewsInner `json:"renews,omitempty"`
	// 決済状況 (未決済: unsettled, 完了: settled)
	Status string `json:"status"`
	// 収支区分 (収入: income, 支出: expense)
	Type *string `json:"type,omitempty"`
}

// NewDeal instantiates a new Deal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeal(amount int64, companyId int64, id int64, issueDate string, partnerId int64, status string) *Deal {
	this := Deal{}
	this.Amount = amount
	this.CompanyId = companyId
	this.Id = id
	this.IssueDate = issueDate
	this.PartnerId = partnerId
	this.Status = status
	return &this
}

// NewDealWithDefaults instantiates a new Deal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDealWithDefaults() *Deal {
	this := Deal{}
	return &this
}

// GetAmount returns the Amount field value
func (o *Deal) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Deal) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Deal) SetAmount(v int64) {
	o.Amount = v
}

// GetCompanyId returns the CompanyId field value
func (o *Deal) GetCompanyId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value
// and a boolean to check if the value has been set.
func (o *Deal) GetCompanyIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyId, true
}

// SetCompanyId sets field value
func (o *Deal) SetCompanyId(v int64) {
	o.CompanyId = v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *Deal) GetDetails() []DealDetailsInner {
	if o == nil || o.Details == nil {
		var ret []DealDetailsInner
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deal) GetDetailsOk() ([]DealDetailsInner, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *Deal) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []DealDetailsInner and assigns it to the Details field.
func (o *Deal) SetDetails(v []DealDetailsInner) {
	o.Details = v
}

// GetDueAmount returns the DueAmount field value if set, zero value otherwise.
func (o *Deal) GetDueAmount() int64 {
	if o == nil || o.DueAmount == nil {
		var ret int64
		return ret
	}
	return *o.DueAmount
}

// GetDueAmountOk returns a tuple with the DueAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deal) GetDueAmountOk() (*int64, bool) {
	if o == nil || o.DueAmount == nil {
		return nil, false
	}
	return o.DueAmount, true
}

// HasDueAmount returns a boolean if a field has been set.
func (o *Deal) HasDueAmount() bool {
	if o != nil && o.DueAmount != nil {
		return true
	}

	return false
}

// SetDueAmount gets a reference to the given int64 and assigns it to the DueAmount field.
func (o *Deal) SetDueAmount(v int64) {
	o.DueAmount = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *Deal) GetDueDate() string {
	if o == nil || o.DueDate == nil {
		var ret string
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deal) GetDueDateOk() (*string, bool) {
	if o == nil || o.DueDate == nil {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *Deal) HasDueDate() bool {
	if o != nil && o.DueDate != nil {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given string and assigns it to the DueDate field.
func (o *Deal) SetDueDate(v string) {
	o.DueDate = &v
}

// GetId returns the Id field value
func (o *Deal) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Deal) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Deal) SetId(v int64) {
	o.Id = v
}

// GetIssueDate returns the IssueDate field value
func (o *Deal) GetIssueDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssueDate
}

// GetIssueDateOk returns a tuple with the IssueDate field value
// and a boolean to check if the value has been set.
func (o *Deal) GetIssueDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueDate, true
}

// SetIssueDate sets field value
func (o *Deal) SetIssueDate(v string) {
	o.IssueDate = v
}

// GetPartnerCode returns the PartnerCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Deal) GetPartnerCode() string {
	if o == nil || o.PartnerCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.PartnerCode.Get()
}

// GetPartnerCodeOk returns a tuple with the PartnerCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Deal) GetPartnerCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PartnerCode.Get(), o.PartnerCode.IsSet()
}

// HasPartnerCode returns a boolean if a field has been set.
func (o *Deal) HasPartnerCode() bool {
	if o != nil && o.PartnerCode.IsSet() {
		return true
	}

	return false
}

// SetPartnerCode gets a reference to the given NullableString and assigns it to the PartnerCode field.
func (o *Deal) SetPartnerCode(v string) {
	o.PartnerCode.Set(&v)
}
// SetPartnerCodeNil sets the value for PartnerCode to be an explicit nil
func (o *Deal) SetPartnerCodeNil() {
	o.PartnerCode.Set(nil)
}

// UnsetPartnerCode ensures that no value is present for PartnerCode, not even an explicit nil
func (o *Deal) UnsetPartnerCode() {
	o.PartnerCode.Unset()
}

// GetPartnerId returns the PartnerId field value
func (o *Deal) GetPartnerId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value
// and a boolean to check if the value has been set.
func (o *Deal) GetPartnerIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartnerId, true
}

// SetPartnerId sets field value
func (o *Deal) SetPartnerId(v int64) {
	o.PartnerId = v
}

// GetPayments returns the Payments field value if set, zero value otherwise.
func (o *Deal) GetPayments() []DealPaymentsInner {
	if o == nil || o.Payments == nil {
		var ret []DealPaymentsInner
		return ret
	}
	return o.Payments
}

// GetPaymentsOk returns a tuple with the Payments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deal) GetPaymentsOk() ([]DealPaymentsInner, bool) {
	if o == nil || o.Payments == nil {
		return nil, false
	}
	return o.Payments, true
}

// HasPayments returns a boolean if a field has been set.
func (o *Deal) HasPayments() bool {
	if o != nil && o.Payments != nil {
		return true
	}

	return false
}

// SetPayments gets a reference to the given []DealPaymentsInner and assigns it to the Payments field.
func (o *Deal) SetPayments(v []DealPaymentsInner) {
	o.Payments = v
}

// GetReceipts returns the Receipts field value if set, zero value otherwise.
func (o *Deal) GetReceipts() []DealReceiptsInner {
	if o == nil || o.Receipts == nil {
		var ret []DealReceiptsInner
		return ret
	}
	return o.Receipts
}

// GetReceiptsOk returns a tuple with the Receipts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deal) GetReceiptsOk() ([]DealReceiptsInner, bool) {
	if o == nil || o.Receipts == nil {
		return nil, false
	}
	return o.Receipts, true
}

// HasReceipts returns a boolean if a field has been set.
func (o *Deal) HasReceipts() bool {
	if o != nil && o.Receipts != nil {
		return true
	}

	return false
}

// SetReceipts gets a reference to the given []DealReceiptsInner and assigns it to the Receipts field.
func (o *Deal) SetReceipts(v []DealReceiptsInner) {
	o.Receipts = v
}

// GetRefNumber returns the RefNumber field value if set, zero value otherwise.
func (o *Deal) GetRefNumber() string {
	if o == nil || o.RefNumber == nil {
		var ret string
		return ret
	}
	return *o.RefNumber
}

// GetRefNumberOk returns a tuple with the RefNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deal) GetRefNumberOk() (*string, bool) {
	if o == nil || o.RefNumber == nil {
		return nil, false
	}
	return o.RefNumber, true
}

// HasRefNumber returns a boolean if a field has been set.
func (o *Deal) HasRefNumber() bool {
	if o != nil && o.RefNumber != nil {
		return true
	}

	return false
}

// SetRefNumber gets a reference to the given string and assigns it to the RefNumber field.
func (o *Deal) SetRefNumber(v string) {
	o.RefNumber = &v
}

// GetRenews returns the Renews field value if set, zero value otherwise.
func (o *Deal) GetRenews() []DealRenewsInner {
	if o == nil || o.Renews == nil {
		var ret []DealRenewsInner
		return ret
	}
	return o.Renews
}

// GetRenewsOk returns a tuple with the Renews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deal) GetRenewsOk() ([]DealRenewsInner, bool) {
	if o == nil || o.Renews == nil {
		return nil, false
	}
	return o.Renews, true
}

// HasRenews returns a boolean if a field has been set.
func (o *Deal) HasRenews() bool {
	if o != nil && o.Renews != nil {
		return true
	}

	return false
}

// SetRenews gets a reference to the given []DealRenewsInner and assigns it to the Renews field.
func (o *Deal) SetRenews(v []DealRenewsInner) {
	o.Renews = v
}

// GetStatus returns the Status field value
func (o *Deal) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Deal) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Deal) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Deal) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deal) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Deal) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Deal) SetType(v string) {
	o.Type = &v
}

func (o Deal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["company_id"] = o.CompanyId
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.DueAmount != nil {
		toSerialize["due_amount"] = o.DueAmount
	}
	if o.DueDate != nil {
		toSerialize["due_date"] = o.DueDate
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["issue_date"] = o.IssueDate
	}
	if o.PartnerCode.IsSet() {
		toSerialize["partner_code"] = o.PartnerCode.Get()
	}
	if true {
		toSerialize["partner_id"] = o.PartnerId
	}
	if o.Payments != nil {
		toSerialize["payments"] = o.Payments
	}
	if o.Receipts != nil {
		toSerialize["receipts"] = o.Receipts
	}
	if o.RefNumber != nil {
		toSerialize["ref_number"] = o.RefNumber
	}
	if o.Renews != nil {
		toSerialize["renews"] = o.Renews
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableDeal struct {
	value *Deal
	isSet bool
}

func (v NullableDeal) Get() *Deal {
	return v.value
}

func (v *NullableDeal) Set(val *Deal) {
	v.value = val
	v.isSet = true
}

func (v NullableDeal) IsSet() bool {
	return v.isSet
}

func (v *NullableDeal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeal(val *Deal) *NullableDeal {
	return &NullableDeal{value: val, isSet: true}
}

func (v NullableDeal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


