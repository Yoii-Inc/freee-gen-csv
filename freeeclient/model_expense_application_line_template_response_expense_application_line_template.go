/*
 * freee API
 *
 *  <h1 id=\"freee_api\">freee API</h1> <hr /> <h2 id=\"start_guide\">スタートガイド</h2>  <p>freee API開発がはじめての方は<a href=\"https://developer.freee.co.jp/getting-started\">freee API スタートガイド</a>を参照してください。</p>  <hr /> <h2 id=\"specification\">仕様</h2>  <h3 id=\"api_endpoint\">APIエンドポイント</h3>  <p>https://api.freee.co.jp/ (httpsのみ)</p>  <h3 id=\"about_authorize\">認証について</h3> <p>OAuth2.0を利用します。詳細は<a href=\"https://developer.freee.co.jp/docs\" target=\"_blank\">ドキュメントの認証</a>パートを参照してください。</p>  <h3 id=\"data_format\">データフォーマット</h3>  <p>リクエスト、レスポンスともにJSON形式をサポートしていますが、詳細は、API毎の説明欄（application/jsonなど）を確認してください。</p>  <h3 id=\"compatibility\">後方互換性ありの変更</h3>  <p>freeeでは、APIを改善していくために以下のような変更は後方互換性ありとして通知なく変更を入れることがあります。アプリケーション実装者は以下を踏まえて開発を行ってください。</p>  <ul> <li>新しいAPIリソース・エンドポイントの追加</li> <li>既存のAPIに対して必須ではない新しいリクエストパラメータの追加</li> <li>既存のAPIレスポンスに対する新しいプロパティの追加</li> <li>既存のAPIレスポンスに対するプロパティの順番の入れ変え</li> <li>keyとなっているidやcodeの長さの変更（長くする）</li> </ul>  <h3 id=\"common_response_header\">共通レスポンスヘッダー</h3>  <p>すべてのAPIのレスポンスには以下のHTTPヘッダーが含まれます。</p>  <ul> <li> <p>X-Freee-Request-ID</p> <ul> <li>各リクエスト毎に発行されるID</li> </ul> </li> </ul>  <h3 id=\"common_error_response\">共通エラーレスポンス</h3>  <ul> <li> <p>ステータスコードはレスポンス内のJSONに含まれる他、HTTPヘッダにも含まれる</p> </li> <li> <p>一部のエラーレスポンスにはエラーコードが含まれます。<br>詳細は、<a href=\"https://developer.freee.co.jp/tips/faq/40x-checkpoint\">HTTPステータスコード400台エラー時のチェックポイント</a>を参照してください</p> </li> <p>type</p>  <ul> <li>status : HTTPステータスコードの説明</li>  <li>validation : エラーの詳細の説明（開発者向け）</li> </ul> </li> </ul>  <p>レスポンスの例</p>  <pre><code>  {     &quot;status_code&quot; : 400,     &quot;errors&quot; : [       {         &quot;type&quot; : &quot;status&quot;,         &quot;messages&quot; : [&quot;不正なリクエストです。&quot;]       },       {         &quot;type&quot; : &quot;validation&quot;,         &quot;messages&quot; : [&quot;Date は不正な日付フォーマットです。入力例：2019-12-17&quot;]       }     ]   }</code></pre>  </br>  <h3 id=\"api_rate_limit\">API使用制限</h3>    <p>freeeは一定期間に過度のアクセスを検知した場合、APIアクセスをコントロールする場合があります。</p>   <p>その際のhttp status codeは403となります。制限がかかってから10分程度が過ぎると再度使用することができるようになります。</p>  <h4 id=\"reports_api_endpoint\">/reportsと/receipts/{id}/downloadエンドポイント</h4>  <p>freeeはエンドポイント毎に一定頻度以上のアクセスを検知した場合、APIアクセスをコントロールする場合があります。その際のhttp status codeは429（too many requests）となります。</p>  <ul>   <li>/reports:1秒に10回まで</li>   <li>/receipts/{id}/download:1秒に3回まで</li> </ul>  <p>レスポンスボディのmetaプロパティに以下を含めます。</p>  <ul>   <li>設定されている上限値</li>   <li>上限に達するまでの使用可能回数</li>   <li>（上限値に達した場合）使用回数がリセットされる時刻</li> </ul>  <h3 id=\"plan_api_rate_limit\">プラン別のAPI Rate Limit</h3>   <table border=\"1\">     <tbody>       <tr>         <th style=\"padding: 10px\"><strong>会計freeeプラン名</strong></th>         <th style=\"padding: 10px\"><strong>事業所とアプリケーション毎に1日でのAPIコール数</strong></th>       </tr>       <tr>         <td style=\"padding: 10px\">エンタープライズ</td>         <td style=\"padding: 10px\">10,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">プロフェッショナル</td>         <td style=\"padding: 10px\">5,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">ベーシック</td>         <td style=\"padding: 10px\">3,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">ミニマム</td>         <td style=\"padding: 10px\">3,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">上記以外</td>         <td style=\"padding: 10px\">3,000</td>       </tr>     </tbody>   </table>  <h3 id=\"webhook\">Webhookについて</h3>  <p>詳細は<a href=\"https://developer.freee.co.jp/docs/accounting/webhook\" target=\"_blank\">会計Webhook概要</a>を参照してください。</p>  <hr /> <h2 id=\"contact\">連絡先</h2>  <p>ご不明点、ご要望等は <a href=\"https://support.freee.co.jp/hc/ja/requests/new\">freee サポートデスクへのお問い合わせフォーム</a> からご連絡ください。</p> <hr />&copy; Since 2013 freee K.K.
 *
 * API version: v1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package freeeclient

import (
	"encoding/json"
)

// ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate struct for ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate
type ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate struct {
	// 勘定科目ID
	AccountItemId *int32 `json:"account_item_id,omitempty"`
	// 勘定科目名
	AccountItemName string `json:"account_item_name"`
	// 経費科目の説明
	Description *string `json:"description,omitempty"`
	// 経費科目ID
	Id int32 `json:"id"`
	// 内容の補足
	LineDescription *string `json:"line_description,omitempty"`
	// 経費科目名
	Name string `json:"name"`
	// 添付ファイルの必須/任意
	RequiredReceipt *bool `json:"required_receipt,omitempty"`
	// 税区分コード
	TaxCode *int32 `json:"tax_code,omitempty"`
	// 税区分名
	TaxName string `json:"tax_name"`
}

// NewExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate instantiates a new ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate(accountItemName string, id int32, name string, taxName string) *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate {
	this := ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate{}
	this.AccountItemName = accountItemName
	this.Id = id
	this.Name = name
	this.TaxName = taxName
	return &this
}

// NewExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplateWithDefaults instantiates a new ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplateWithDefaults() *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate {
	this := ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate{}
	return &this
}

// GetAccountItemId returns the AccountItemId field value if set, zero value otherwise.
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetAccountItemId() int32 {
	if o == nil || o.AccountItemId == nil {
		var ret int32
		return ret
	}
	return *o.AccountItemId
}

// GetAccountItemIdOk returns a tuple with the AccountItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetAccountItemIdOk() (*int32, bool) {
	if o == nil || o.AccountItemId == nil {
		return nil, false
	}
	return o.AccountItemId, true
}

// HasAccountItemId returns a boolean if a field has been set.
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) HasAccountItemId() bool {
	if o != nil && o.AccountItemId != nil {
		return true
	}

	return false
}

// SetAccountItemId gets a reference to the given int32 and assigns it to the AccountItemId field.
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) SetAccountItemId(v int32) {
	o.AccountItemId = &v
}

// GetAccountItemName returns the AccountItemName field value
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetAccountItemName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountItemName
}

// GetAccountItemNameOk returns a tuple with the AccountItemName field value
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetAccountItemNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountItemName, true
}

// SetAccountItemName sets field value
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) SetAccountItemName(v string) {
	o.AccountItemName = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) SetId(v int32) {
	o.Id = v
}

// GetLineDescription returns the LineDescription field value if set, zero value otherwise.
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetLineDescription() string {
	if o == nil || o.LineDescription == nil {
		var ret string
		return ret
	}
	return *o.LineDescription
}

// GetLineDescriptionOk returns a tuple with the LineDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetLineDescriptionOk() (*string, bool) {
	if o == nil || o.LineDescription == nil {
		return nil, false
	}
	return o.LineDescription, true
}

// HasLineDescription returns a boolean if a field has been set.
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) HasLineDescription() bool {
	if o != nil && o.LineDescription != nil {
		return true
	}

	return false
}

// SetLineDescription gets a reference to the given string and assigns it to the LineDescription field.
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) SetLineDescription(v string) {
	o.LineDescription = &v
}

// GetName returns the Name field value
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) SetName(v string) {
	o.Name = v
}

// GetRequiredReceipt returns the RequiredReceipt field value if set, zero value otherwise.
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetRequiredReceipt() bool {
	if o == nil || o.RequiredReceipt == nil {
		var ret bool
		return ret
	}
	return *o.RequiredReceipt
}

// GetRequiredReceiptOk returns a tuple with the RequiredReceipt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetRequiredReceiptOk() (*bool, bool) {
	if o == nil || o.RequiredReceipt == nil {
		return nil, false
	}
	return o.RequiredReceipt, true
}

// HasRequiredReceipt returns a boolean if a field has been set.
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) HasRequiredReceipt() bool {
	if o != nil && o.RequiredReceipt != nil {
		return true
	}

	return false
}

// SetRequiredReceipt gets a reference to the given bool and assigns it to the RequiredReceipt field.
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) SetRequiredReceipt(v bool) {
	o.RequiredReceipt = &v
}

// GetTaxCode returns the TaxCode field value if set, zero value otherwise.
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetTaxCode() int32 {
	if o == nil || o.TaxCode == nil {
		var ret int32
		return ret
	}
	return *o.TaxCode
}

// GetTaxCodeOk returns a tuple with the TaxCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetTaxCodeOk() (*int32, bool) {
	if o == nil || o.TaxCode == nil {
		return nil, false
	}
	return o.TaxCode, true
}

// HasTaxCode returns a boolean if a field has been set.
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) HasTaxCode() bool {
	if o != nil && o.TaxCode != nil {
		return true
	}

	return false
}

// SetTaxCode gets a reference to the given int32 and assigns it to the TaxCode field.
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) SetTaxCode(v int32) {
	o.TaxCode = &v
}

// GetTaxName returns the TaxName field value
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetTaxName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaxName
}

// GetTaxNameOk returns a tuple with the TaxName field value
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetTaxNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TaxName, true
}

// SetTaxName sets field value
func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) SetTaxName(v string) {
	o.TaxName = v
}

func (o ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountItemId != nil {
		toSerialize["account_item_id"] = o.AccountItemId
	}
	if true {
		toSerialize["account_item_name"] = o.AccountItemName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.LineDescription != nil {
		toSerialize["line_description"] = o.LineDescription
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.RequiredReceipt != nil {
		toSerialize["required_receipt"] = o.RequiredReceipt
	}
	if o.TaxCode != nil {
		toSerialize["tax_code"] = o.TaxCode
	}
	if true {
		toSerialize["tax_name"] = o.TaxName
	}
	return json.Marshal(toSerialize)
}

type NullableExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate struct {
	value *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate
	isSet bool
}

func (v NullableExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) Get() *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate {
	return v.value
}

func (v *NullableExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) Set(val *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate(val *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) *NullableExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate {
	return &NullableExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate{value: val, isSet: true}
}

func (v NullableExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


