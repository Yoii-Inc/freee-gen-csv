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

// Receipt struct for Receipt
type Receipt struct {
	// 作成日時（ISO8601形式）
	CreatedAt string `json:"created_at"`
	// メモ
	Description *string `json:"description,omitempty"`
	// ファイルのダウンロードURL（freeeにログインした状態でのみ閲覧可能です。） <br> <br> file_srcは廃止予定の属性になります。<br> file_srcに替わり、証憑ファイルのダウンロード APIをご利用ください。<br> 証憑ファイルのダウンロードAPIを利用することで、以下のようになります。 <ul>   <li>アプリケーション利用者はfreee APIアプリケーションにログインしていれば、証憑ダウンロード毎にfreeeに改めてログインすることなくファイルが参照できるようになります。</li> </ul>
	// Deprecated
	FileSrc string `json:"file_src"`
	// 証憑ファイルID
	Id int64 `json:"id"`
	// 発生日
	IssueDate *string `json:"issue_date,omitempty"`
	// MIMEタイプ
	MimeType string `json:"mime_type"`
	// アップロード元種別
	Origin string `json:"origin"`
	ReceiptMetadatum *DealReceiptsInnerReceiptMetadatum `json:"receipt_metadatum,omitempty"`
	// ステータス(confirmed:確認済み、deleted:削除済み、ignored:無視)
	Status string `json:"status"`
	User DealReceiptsInnerUser `json:"user"`
}

// NewReceipt instantiates a new Receipt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceipt(createdAt string, fileSrc string, id int64, mimeType string, origin string, status string, user DealReceiptsInnerUser) *Receipt {
	this := Receipt{}
	this.CreatedAt = createdAt
	this.FileSrc = fileSrc
	this.Id = id
	this.MimeType = mimeType
	this.Origin = origin
	this.Status = status
	this.User = user
	return &this
}

// NewReceiptWithDefaults instantiates a new Receipt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceiptWithDefaults() *Receipt {
	this := Receipt{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *Receipt) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Receipt) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Receipt) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Receipt) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Receipt) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Receipt) SetDescription(v string) {
	o.Description = &v
}

// GetFileSrc returns the FileSrc field value
// Deprecated
func (o *Receipt) GetFileSrc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileSrc
}

// GetFileSrcOk returns a tuple with the FileSrc field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *Receipt) GetFileSrcOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileSrc, true
}

// SetFileSrc sets field value
// Deprecated
func (o *Receipt) SetFileSrc(v string) {
	o.FileSrc = v
}

// GetId returns the Id field value
func (o *Receipt) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Receipt) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Receipt) SetId(v int64) {
	o.Id = v
}

// GetIssueDate returns the IssueDate field value if set, zero value otherwise.
func (o *Receipt) GetIssueDate() string {
	if o == nil || o.IssueDate == nil {
		var ret string
		return ret
	}
	return *o.IssueDate
}

// GetIssueDateOk returns a tuple with the IssueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetIssueDateOk() (*string, bool) {
	if o == nil || o.IssueDate == nil {
		return nil, false
	}
	return o.IssueDate, true
}

// HasIssueDate returns a boolean if a field has been set.
func (o *Receipt) HasIssueDate() bool {
	if o != nil && o.IssueDate != nil {
		return true
	}

	return false
}

// SetIssueDate gets a reference to the given string and assigns it to the IssueDate field.
func (o *Receipt) SetIssueDate(v string) {
	o.IssueDate = &v
}

// GetMimeType returns the MimeType field value
func (o *Receipt) GetMimeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value
// and a boolean to check if the value has been set.
func (o *Receipt) GetMimeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MimeType, true
}

// SetMimeType sets field value
func (o *Receipt) SetMimeType(v string) {
	o.MimeType = v
}

// GetOrigin returns the Origin field value
func (o *Receipt) GetOrigin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Origin
}

// GetOriginOk returns a tuple with the Origin field value
// and a boolean to check if the value has been set.
func (o *Receipt) GetOriginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Origin, true
}

// SetOrigin sets field value
func (o *Receipt) SetOrigin(v string) {
	o.Origin = v
}

// GetReceiptMetadatum returns the ReceiptMetadatum field value if set, zero value otherwise.
func (o *Receipt) GetReceiptMetadatum() DealReceiptsInnerReceiptMetadatum {
	if o == nil || o.ReceiptMetadatum == nil {
		var ret DealReceiptsInnerReceiptMetadatum
		return ret
	}
	return *o.ReceiptMetadatum
}

// GetReceiptMetadatumOk returns a tuple with the ReceiptMetadatum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetReceiptMetadatumOk() (*DealReceiptsInnerReceiptMetadatum, bool) {
	if o == nil || o.ReceiptMetadatum == nil {
		return nil, false
	}
	return o.ReceiptMetadatum, true
}

// HasReceiptMetadatum returns a boolean if a field has been set.
func (o *Receipt) HasReceiptMetadatum() bool {
	if o != nil && o.ReceiptMetadatum != nil {
		return true
	}

	return false
}

// SetReceiptMetadatum gets a reference to the given DealReceiptsInnerReceiptMetadatum and assigns it to the ReceiptMetadatum field.
func (o *Receipt) SetReceiptMetadatum(v DealReceiptsInnerReceiptMetadatum) {
	o.ReceiptMetadatum = &v
}

// GetStatus returns the Status field value
func (o *Receipt) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Receipt) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Receipt) SetStatus(v string) {
	o.Status = v
}

// GetUser returns the User field value
func (o *Receipt) GetUser() DealReceiptsInnerUser {
	if o == nil {
		var ret DealReceiptsInnerUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *Receipt) GetUserOk() (*DealReceiptsInnerUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *Receipt) SetUser(v DealReceiptsInnerUser) {
	o.User = v
}

func (o Receipt) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["file_src"] = o.FileSrc
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.IssueDate != nil {
		toSerialize["issue_date"] = o.IssueDate
	}
	if true {
		toSerialize["mime_type"] = o.MimeType
	}
	if true {
		toSerialize["origin"] = o.Origin
	}
	if o.ReceiptMetadatum != nil {
		toSerialize["receipt_metadatum"] = o.ReceiptMetadatum
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableReceipt struct {
	value *Receipt
	isSet bool
}

func (v NullableReceipt) Get() *Receipt {
	return v.value
}

func (v *NullableReceipt) Set(val *Receipt) {
	v.value = val
	v.isSet = true
}

func (v NullableReceipt) IsSet() bool {
	return v.isSet
}

func (v *NullableReceipt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceipt(val *Receipt) *NullableReceipt {
	return &NullableReceipt{value: val, isSet: true}
}

func (v NullableReceipt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceipt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


