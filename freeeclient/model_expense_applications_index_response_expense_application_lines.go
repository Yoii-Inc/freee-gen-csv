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

// ExpenseApplicationsIndexResponseExpenseApplicationLines struct for ExpenseApplicationsIndexResponseExpenseApplicationLines
type ExpenseApplicationsIndexResponseExpenseApplicationLines struct {
	// 金額
	Amount *int32 `json:"amount,omitempty"`
	// 内容
	Description *string `json:"description,omitempty"`
	// 経費科目ID
	ExpenseApplicationLineTemplateId *int32 `json:"expense_application_line_template_id,omitempty"`
	// 経費申請の項目行ID
	Id int64 `json:"id"`
	// 証憑ファイルID（ファイルボックスのファイルID）
	ReceiptId *int32 `json:"receipt_id,omitempty"`
	// 日付 (yyyy-mm-dd)
	TransactionDate *string `json:"transaction_date,omitempty"`
}

// NewExpenseApplicationsIndexResponseExpenseApplicationLines instantiates a new ExpenseApplicationsIndexResponseExpenseApplicationLines object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpenseApplicationsIndexResponseExpenseApplicationLines(id int64) *ExpenseApplicationsIndexResponseExpenseApplicationLines {
	this := ExpenseApplicationsIndexResponseExpenseApplicationLines{}
	this.Id = id
	return &this
}

// NewExpenseApplicationsIndexResponseExpenseApplicationLinesWithDefaults instantiates a new ExpenseApplicationsIndexResponseExpenseApplicationLines object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpenseApplicationsIndexResponseExpenseApplicationLinesWithDefaults() *ExpenseApplicationsIndexResponseExpenseApplicationLines {
	this := ExpenseApplicationsIndexResponseExpenseApplicationLines{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) GetAmount() int32 {
	if o == nil || o.Amount == nil {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) GetAmountOk() (*int32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) SetAmount(v int32) {
	o.Amount = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) SetDescription(v string) {
	o.Description = &v
}

// GetExpenseApplicationLineTemplateId returns the ExpenseApplicationLineTemplateId field value if set, zero value otherwise.
func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) GetExpenseApplicationLineTemplateId() int32 {
	if o == nil || o.ExpenseApplicationLineTemplateId == nil {
		var ret int32
		return ret
	}
	return *o.ExpenseApplicationLineTemplateId
}

// GetExpenseApplicationLineTemplateIdOk returns a tuple with the ExpenseApplicationLineTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) GetExpenseApplicationLineTemplateIdOk() (*int32, bool) {
	if o == nil || o.ExpenseApplicationLineTemplateId == nil {
		return nil, false
	}
	return o.ExpenseApplicationLineTemplateId, true
}

// HasExpenseApplicationLineTemplateId returns a boolean if a field has been set.
func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) HasExpenseApplicationLineTemplateId() bool {
	if o != nil && o.ExpenseApplicationLineTemplateId != nil {
		return true
	}

	return false
}

// SetExpenseApplicationLineTemplateId gets a reference to the given int32 and assigns it to the ExpenseApplicationLineTemplateId field.
func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) SetExpenseApplicationLineTemplateId(v int32) {
	o.ExpenseApplicationLineTemplateId = &v
}

// GetId returns the Id field value
func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) SetId(v int64) {
	o.Id = v
}

// GetReceiptId returns the ReceiptId field value if set, zero value otherwise.
func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) GetReceiptId() int32 {
	if o == nil || o.ReceiptId == nil {
		var ret int32
		return ret
	}
	return *o.ReceiptId
}

// GetReceiptIdOk returns a tuple with the ReceiptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) GetReceiptIdOk() (*int32, bool) {
	if o == nil || o.ReceiptId == nil {
		return nil, false
	}
	return o.ReceiptId, true
}

// HasReceiptId returns a boolean if a field has been set.
func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) HasReceiptId() bool {
	if o != nil && o.ReceiptId != nil {
		return true
	}

	return false
}

// SetReceiptId gets a reference to the given int32 and assigns it to the ReceiptId field.
func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) SetReceiptId(v int32) {
	o.ReceiptId = &v
}

// GetTransactionDate returns the TransactionDate field value if set, zero value otherwise.
func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) GetTransactionDate() string {
	if o == nil || o.TransactionDate == nil {
		var ret string
		return ret
	}
	return *o.TransactionDate
}

// GetTransactionDateOk returns a tuple with the TransactionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) GetTransactionDateOk() (*string, bool) {
	if o == nil || o.TransactionDate == nil {
		return nil, false
	}
	return o.TransactionDate, true
}

// HasTransactionDate returns a boolean if a field has been set.
func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) HasTransactionDate() bool {
	if o != nil && o.TransactionDate != nil {
		return true
	}

	return false
}

// SetTransactionDate gets a reference to the given string and assigns it to the TransactionDate field.
func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) SetTransactionDate(v string) {
	o.TransactionDate = &v
}

func (o ExpenseApplicationsIndexResponseExpenseApplicationLines) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ExpenseApplicationLineTemplateId != nil {
		toSerialize["expense_application_line_template_id"] = o.ExpenseApplicationLineTemplateId
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.ReceiptId != nil {
		toSerialize["receipt_id"] = o.ReceiptId
	}
	if o.TransactionDate != nil {
		toSerialize["transaction_date"] = o.TransactionDate
	}
	return json.Marshal(toSerialize)
}

type NullableExpenseApplicationsIndexResponseExpenseApplicationLines struct {
	value *ExpenseApplicationsIndexResponseExpenseApplicationLines
	isSet bool
}

func (v NullableExpenseApplicationsIndexResponseExpenseApplicationLines) Get() *ExpenseApplicationsIndexResponseExpenseApplicationLines {
	return v.value
}

func (v *NullableExpenseApplicationsIndexResponseExpenseApplicationLines) Set(val *ExpenseApplicationsIndexResponseExpenseApplicationLines) {
	v.value = val
	v.isSet = true
}

func (v NullableExpenseApplicationsIndexResponseExpenseApplicationLines) IsSet() bool {
	return v.isSet
}

func (v *NullableExpenseApplicationsIndexResponseExpenseApplicationLines) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpenseApplicationsIndexResponseExpenseApplicationLines(val *ExpenseApplicationsIndexResponseExpenseApplicationLines) *NullableExpenseApplicationsIndexResponseExpenseApplicationLines {
	return &NullableExpenseApplicationsIndexResponseExpenseApplicationLines{value: val, isSet: true}
}

func (v NullableExpenseApplicationsIndexResponseExpenseApplicationLines) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpenseApplicationsIndexResponseExpenseApplicationLines) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


