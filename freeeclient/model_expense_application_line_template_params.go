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

// ExpenseApplicationLineTemplateParams struct for ExpenseApplicationLineTemplateParams
type ExpenseApplicationLineTemplateParams struct {
	// 勘定科目ID
	AccountItemId int32 `json:"account_item_id"`
	// 事業所ID
	CompanyId int32 `json:"company_id"`
	// 経費科目の説明 (1000文字以内)
	Description *string `json:"description,omitempty"`
	// 品目ID
	ItemId *int32 `json:"item_id,omitempty"`
	// 内容の補足 (1000文字以内)
	LineDescription *string `json:"line_description,omitempty"`
	// 経費科目名 (100文字以内)
	Name string `json:"name"`
	// 添付ファイルの必須/任意
	RequiredReceipt *bool `json:"required_receipt,omitempty"`
	// 税区分コード（税区分のdisplay_categoryがtax_5: 5%表示の税区分, tax_r8: 軽減税率8%表示の税区分に該当するtax_codeのみ利用可能です。税区分のdisplay_categoryは /taxes/companies/{:company_id}のAPIから取得可能です。）
	TaxCode int32 `json:"tax_code"`
}

// NewExpenseApplicationLineTemplateParams instantiates a new ExpenseApplicationLineTemplateParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpenseApplicationLineTemplateParams(accountItemId int32, companyId int32, name string, taxCode int32) *ExpenseApplicationLineTemplateParams {
	this := ExpenseApplicationLineTemplateParams{}
	this.AccountItemId = accountItemId
	this.CompanyId = companyId
	this.Name = name
	this.TaxCode = taxCode
	return &this
}

// NewExpenseApplicationLineTemplateParamsWithDefaults instantiates a new ExpenseApplicationLineTemplateParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpenseApplicationLineTemplateParamsWithDefaults() *ExpenseApplicationLineTemplateParams {
	this := ExpenseApplicationLineTemplateParams{}
	return &this
}

// GetAccountItemId returns the AccountItemId field value
func (o *ExpenseApplicationLineTemplateParams) GetAccountItemId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountItemId
}

// GetAccountItemIdOk returns a tuple with the AccountItemId field value
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationLineTemplateParams) GetAccountItemIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountItemId, true
}

// SetAccountItemId sets field value
func (o *ExpenseApplicationLineTemplateParams) SetAccountItemId(v int32) {
	o.AccountItemId = v
}

// GetCompanyId returns the CompanyId field value
func (o *ExpenseApplicationLineTemplateParams) GetCompanyId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationLineTemplateParams) GetCompanyIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CompanyId, true
}

// SetCompanyId sets field value
func (o *ExpenseApplicationLineTemplateParams) SetCompanyId(v int32) {
	o.CompanyId = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExpenseApplicationLineTemplateParams) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationLineTemplateParams) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExpenseApplicationLineTemplateParams) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExpenseApplicationLineTemplateParams) SetDescription(v string) {
	o.Description = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *ExpenseApplicationLineTemplateParams) GetItemId() int32 {
	if o == nil || o.ItemId == nil {
		var ret int32
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationLineTemplateParams) GetItemIdOk() (*int32, bool) {
	if o == nil || o.ItemId == nil {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *ExpenseApplicationLineTemplateParams) HasItemId() bool {
	if o != nil && o.ItemId != nil {
		return true
	}

	return false
}

// SetItemId gets a reference to the given int32 and assigns it to the ItemId field.
func (o *ExpenseApplicationLineTemplateParams) SetItemId(v int32) {
	o.ItemId = &v
}

// GetLineDescription returns the LineDescription field value if set, zero value otherwise.
func (o *ExpenseApplicationLineTemplateParams) GetLineDescription() string {
	if o == nil || o.LineDescription == nil {
		var ret string
		return ret
	}
	return *o.LineDescription
}

// GetLineDescriptionOk returns a tuple with the LineDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationLineTemplateParams) GetLineDescriptionOk() (*string, bool) {
	if o == nil || o.LineDescription == nil {
		return nil, false
	}
	return o.LineDescription, true
}

// HasLineDescription returns a boolean if a field has been set.
func (o *ExpenseApplicationLineTemplateParams) HasLineDescription() bool {
	if o != nil && o.LineDescription != nil {
		return true
	}

	return false
}

// SetLineDescription gets a reference to the given string and assigns it to the LineDescription field.
func (o *ExpenseApplicationLineTemplateParams) SetLineDescription(v string) {
	o.LineDescription = &v
}

// GetName returns the Name field value
func (o *ExpenseApplicationLineTemplateParams) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationLineTemplateParams) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExpenseApplicationLineTemplateParams) SetName(v string) {
	o.Name = v
}

// GetRequiredReceipt returns the RequiredReceipt field value if set, zero value otherwise.
func (o *ExpenseApplicationLineTemplateParams) GetRequiredReceipt() bool {
	if o == nil || o.RequiredReceipt == nil {
		var ret bool
		return ret
	}
	return *o.RequiredReceipt
}

// GetRequiredReceiptOk returns a tuple with the RequiredReceipt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationLineTemplateParams) GetRequiredReceiptOk() (*bool, bool) {
	if o == nil || o.RequiredReceipt == nil {
		return nil, false
	}
	return o.RequiredReceipt, true
}

// HasRequiredReceipt returns a boolean if a field has been set.
func (o *ExpenseApplicationLineTemplateParams) HasRequiredReceipt() bool {
	if o != nil && o.RequiredReceipt != nil {
		return true
	}

	return false
}

// SetRequiredReceipt gets a reference to the given bool and assigns it to the RequiredReceipt field.
func (o *ExpenseApplicationLineTemplateParams) SetRequiredReceipt(v bool) {
	o.RequiredReceipt = &v
}

// GetTaxCode returns the TaxCode field value
func (o *ExpenseApplicationLineTemplateParams) GetTaxCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TaxCode
}

// GetTaxCodeOk returns a tuple with the TaxCode field value
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationLineTemplateParams) GetTaxCodeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TaxCode, true
}

// SetTaxCode sets field value
func (o *ExpenseApplicationLineTemplateParams) SetTaxCode(v int32) {
	o.TaxCode = v
}

func (o ExpenseApplicationLineTemplateParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_item_id"] = o.AccountItemId
	}
	if true {
		toSerialize["company_id"] = o.CompanyId
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ItemId != nil {
		toSerialize["item_id"] = o.ItemId
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
	if true {
		toSerialize["tax_code"] = o.TaxCode
	}
	return json.Marshal(toSerialize)
}

type NullableExpenseApplicationLineTemplateParams struct {
	value *ExpenseApplicationLineTemplateParams
	isSet bool
}

func (v NullableExpenseApplicationLineTemplateParams) Get() *ExpenseApplicationLineTemplateParams {
	return v.value
}

func (v *NullableExpenseApplicationLineTemplateParams) Set(val *ExpenseApplicationLineTemplateParams) {
	v.value = val
	v.isSet = true
}

func (v NullableExpenseApplicationLineTemplateParams) IsSet() bool {
	return v.isSet
}

func (v *NullableExpenseApplicationLineTemplateParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpenseApplicationLineTemplateParams(val *ExpenseApplicationLineTemplateParams) *NullableExpenseApplicationLineTemplateParams {
	return &NullableExpenseApplicationLineTemplateParams{value: val, isSet: true}
}

func (v NullableExpenseApplicationLineTemplateParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpenseApplicationLineTemplateParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


