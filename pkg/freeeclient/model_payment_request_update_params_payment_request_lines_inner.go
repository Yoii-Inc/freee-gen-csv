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

// PaymentRequestUpdateParamsPaymentRequestLinesInner struct for PaymentRequestUpdateParamsPaymentRequestLinesInner
type PaymentRequestUpdateParamsPaymentRequestLinesInner struct {
	// 勘定科目ID
	AccountItemId NullableInt32 `json:"account_item_id,omitempty"`
	// 金額
	Amount int64 `json:"amount"`
	// 内容
	Description *string `json:"description,omitempty"`
	// 支払依頼の項目行ID: 既存項目行を更新する場合に指定します。IDを指定しない項目行は、新規行として扱われ追加されます。また、payment_request_linesに含まれない既存の項目行は削除されます。更新後も残したい行は、必ず支払依頼の項目行IDを指定してpayment_request_linesに含めてください。
	Id *int64 `json:"id,omitempty"`
	// 品目ID
	ItemId *int32 `json:"item_id,omitempty"`
	// '行の種類 (deal_line: 支払依頼, withholding_tax: 源泉徴収税)'<br> 'デフォルトは deal_line: 支払依頼 です' 
	LineType *string `json:"line_type,omitempty"`
	// 部門ID
	SectionId *int32 `json:"section_id,omitempty"`
	// セグメント１ID<br> セグメントタグ一覧APIを利用して取得してください。<br> <a href=\"https://support.freee.co.jp/hc/ja/articles/360020679611\" target=\"_blank\">セグメント（分析用タグ）の設定</a><br> 
	Segment1TagId *int64 `json:"segment_1_tag_id,omitempty"`
	// セグメント２ID(法人向けエンタープライズプラン)<br> セグメントタグ一覧APIを利用して取得してください。<br> <a href=\"https://support.freee.co.jp/hc/ja/articles/360020679611\" target=\"_blank\">セグメント（分析用タグ）の設定</a><br> 
	Segment2TagId *int64 `json:"segment_2_tag_id,omitempty"`
	// セグメント３ID(法人向けエンタープライズプラン)<br> セグメントタグ一覧APIを利用して取得してください。<br> <a href=\"https://support.freee.co.jp/hc/ja/articles/360020679611\" target=\"_blank\">セグメント（分析用タグ）の設定</a><br> 
	Segment3TagId *int64 `json:"segment_3_tag_id,omitempty"`
	// メモタグID
	TagIds []int32 `json:"tag_ids,omitempty"`
	// 税区分コード<br> 勘定科目IDを指定する場合は必須です。 
	TaxCode *int32 `json:"tax_code,omitempty"`
}

// NewPaymentRequestUpdateParamsPaymentRequestLinesInner instantiates a new PaymentRequestUpdateParamsPaymentRequestLinesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentRequestUpdateParamsPaymentRequestLinesInner(amount int64) *PaymentRequestUpdateParamsPaymentRequestLinesInner {
	this := PaymentRequestUpdateParamsPaymentRequestLinesInner{}
	this.Amount = amount
	return &this
}

// NewPaymentRequestUpdateParamsPaymentRequestLinesInnerWithDefaults instantiates a new PaymentRequestUpdateParamsPaymentRequestLinesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentRequestUpdateParamsPaymentRequestLinesInnerWithDefaults() *PaymentRequestUpdateParamsPaymentRequestLinesInner {
	this := PaymentRequestUpdateParamsPaymentRequestLinesInner{}
	return &this
}

// GetAccountItemId returns the AccountItemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetAccountItemId() int32 {
	if o == nil || o.AccountItemId.Get() == nil {
		var ret int32
		return ret
	}
	return *o.AccountItemId.Get()
}

// GetAccountItemIdOk returns a tuple with the AccountItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetAccountItemIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountItemId.Get(), o.AccountItemId.IsSet()
}

// HasAccountItemId returns a boolean if a field has been set.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) HasAccountItemId() bool {
	if o != nil && o.AccountItemId.IsSet() {
		return true
	}

	return false
}

// SetAccountItemId gets a reference to the given NullableInt32 and assigns it to the AccountItemId field.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) SetAccountItemId(v int32) {
	o.AccountItemId.Set(&v)
}
// SetAccountItemIdNil sets the value for AccountItemId to be an explicit nil
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) SetAccountItemIdNil() {
	o.AccountItemId.Set(nil)
}

// UnsetAccountItemId ensures that no value is present for AccountItemId, not even an explicit nil
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) UnsetAccountItemId() {
	o.AccountItemId.Unset()
}

// GetAmount returns the Amount field value
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) SetAmount(v int64) {
	o.Amount = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) SetId(v int64) {
	o.Id = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetItemId() int32 {
	if o == nil || o.ItemId == nil {
		var ret int32
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetItemIdOk() (*int32, bool) {
	if o == nil || o.ItemId == nil {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) HasItemId() bool {
	if o != nil && o.ItemId != nil {
		return true
	}

	return false
}

// SetItemId gets a reference to the given int32 and assigns it to the ItemId field.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) SetItemId(v int32) {
	o.ItemId = &v
}

// GetLineType returns the LineType field value if set, zero value otherwise.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetLineType() string {
	if o == nil || o.LineType == nil {
		var ret string
		return ret
	}
	return *o.LineType
}

// GetLineTypeOk returns a tuple with the LineType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetLineTypeOk() (*string, bool) {
	if o == nil || o.LineType == nil {
		return nil, false
	}
	return o.LineType, true
}

// HasLineType returns a boolean if a field has been set.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) HasLineType() bool {
	if o != nil && o.LineType != nil {
		return true
	}

	return false
}

// SetLineType gets a reference to the given string and assigns it to the LineType field.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) SetLineType(v string) {
	o.LineType = &v
}

// GetSectionId returns the SectionId field value if set, zero value otherwise.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetSectionId() int32 {
	if o == nil || o.SectionId == nil {
		var ret int32
		return ret
	}
	return *o.SectionId
}

// GetSectionIdOk returns a tuple with the SectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetSectionIdOk() (*int32, bool) {
	if o == nil || o.SectionId == nil {
		return nil, false
	}
	return o.SectionId, true
}

// HasSectionId returns a boolean if a field has been set.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) HasSectionId() bool {
	if o != nil && o.SectionId != nil {
		return true
	}

	return false
}

// SetSectionId gets a reference to the given int32 and assigns it to the SectionId field.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) SetSectionId(v int32) {
	o.SectionId = &v
}

// GetSegment1TagId returns the Segment1TagId field value if set, zero value otherwise.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetSegment1TagId() int64 {
	if o == nil || o.Segment1TagId == nil {
		var ret int64
		return ret
	}
	return *o.Segment1TagId
}

// GetSegment1TagIdOk returns a tuple with the Segment1TagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetSegment1TagIdOk() (*int64, bool) {
	if o == nil || o.Segment1TagId == nil {
		return nil, false
	}
	return o.Segment1TagId, true
}

// HasSegment1TagId returns a boolean if a field has been set.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) HasSegment1TagId() bool {
	if o != nil && o.Segment1TagId != nil {
		return true
	}

	return false
}

// SetSegment1TagId gets a reference to the given int64 and assigns it to the Segment1TagId field.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) SetSegment1TagId(v int64) {
	o.Segment1TagId = &v
}

// GetSegment2TagId returns the Segment2TagId field value if set, zero value otherwise.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetSegment2TagId() int64 {
	if o == nil || o.Segment2TagId == nil {
		var ret int64
		return ret
	}
	return *o.Segment2TagId
}

// GetSegment2TagIdOk returns a tuple with the Segment2TagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetSegment2TagIdOk() (*int64, bool) {
	if o == nil || o.Segment2TagId == nil {
		return nil, false
	}
	return o.Segment2TagId, true
}

// HasSegment2TagId returns a boolean if a field has been set.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) HasSegment2TagId() bool {
	if o != nil && o.Segment2TagId != nil {
		return true
	}

	return false
}

// SetSegment2TagId gets a reference to the given int64 and assigns it to the Segment2TagId field.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) SetSegment2TagId(v int64) {
	o.Segment2TagId = &v
}

// GetSegment3TagId returns the Segment3TagId field value if set, zero value otherwise.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetSegment3TagId() int64 {
	if o == nil || o.Segment3TagId == nil {
		var ret int64
		return ret
	}
	return *o.Segment3TagId
}

// GetSegment3TagIdOk returns a tuple with the Segment3TagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetSegment3TagIdOk() (*int64, bool) {
	if o == nil || o.Segment3TagId == nil {
		return nil, false
	}
	return o.Segment3TagId, true
}

// HasSegment3TagId returns a boolean if a field has been set.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) HasSegment3TagId() bool {
	if o != nil && o.Segment3TagId != nil {
		return true
	}

	return false
}

// SetSegment3TagId gets a reference to the given int64 and assigns it to the Segment3TagId field.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) SetSegment3TagId(v int64) {
	o.Segment3TagId = &v
}

// GetTagIds returns the TagIds field value if set, zero value otherwise.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetTagIds() []int32 {
	if o == nil || o.TagIds == nil {
		var ret []int32
		return ret
	}
	return o.TagIds
}

// GetTagIdsOk returns a tuple with the TagIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetTagIdsOk() ([]int32, bool) {
	if o == nil || o.TagIds == nil {
		return nil, false
	}
	return o.TagIds, true
}

// HasTagIds returns a boolean if a field has been set.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) HasTagIds() bool {
	if o != nil && o.TagIds != nil {
		return true
	}

	return false
}

// SetTagIds gets a reference to the given []int32 and assigns it to the TagIds field.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) SetTagIds(v []int32) {
	o.TagIds = v
}

// GetTaxCode returns the TaxCode field value if set, zero value otherwise.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetTaxCode() int32 {
	if o == nil || o.TaxCode == nil {
		var ret int32
		return ret
	}
	return *o.TaxCode
}

// GetTaxCodeOk returns a tuple with the TaxCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetTaxCodeOk() (*int32, bool) {
	if o == nil || o.TaxCode == nil {
		return nil, false
	}
	return o.TaxCode, true
}

// HasTaxCode returns a boolean if a field has been set.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) HasTaxCode() bool {
	if o != nil && o.TaxCode != nil {
		return true
	}

	return false
}

// SetTaxCode gets a reference to the given int32 and assigns it to the TaxCode field.
func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) SetTaxCode(v int32) {
	o.TaxCode = &v
}

func (o PaymentRequestUpdateParamsPaymentRequestLinesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountItemId.IsSet() {
		toSerialize["account_item_id"] = o.AccountItemId.Get()
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ItemId != nil {
		toSerialize["item_id"] = o.ItemId
	}
	if o.LineType != nil {
		toSerialize["line_type"] = o.LineType
	}
	if o.SectionId != nil {
		toSerialize["section_id"] = o.SectionId
	}
	if o.Segment1TagId != nil {
		toSerialize["segment_1_tag_id"] = o.Segment1TagId
	}
	if o.Segment2TagId != nil {
		toSerialize["segment_2_tag_id"] = o.Segment2TagId
	}
	if o.Segment3TagId != nil {
		toSerialize["segment_3_tag_id"] = o.Segment3TagId
	}
	if o.TagIds != nil {
		toSerialize["tag_ids"] = o.TagIds
	}
	if o.TaxCode != nil {
		toSerialize["tax_code"] = o.TaxCode
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentRequestUpdateParamsPaymentRequestLinesInner struct {
	value *PaymentRequestUpdateParamsPaymentRequestLinesInner
	isSet bool
}

func (v NullablePaymentRequestUpdateParamsPaymentRequestLinesInner) Get() *PaymentRequestUpdateParamsPaymentRequestLinesInner {
	return v.value
}

func (v *NullablePaymentRequestUpdateParamsPaymentRequestLinesInner) Set(val *PaymentRequestUpdateParamsPaymentRequestLinesInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentRequestUpdateParamsPaymentRequestLinesInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentRequestUpdateParamsPaymentRequestLinesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentRequestUpdateParamsPaymentRequestLinesInner(val *PaymentRequestUpdateParamsPaymentRequestLinesInner) *NullablePaymentRequestUpdateParamsPaymentRequestLinesInner {
	return &NullablePaymentRequestUpdateParamsPaymentRequestLinesInner{value: val, isSet: true}
}

func (v NullablePaymentRequestUpdateParamsPaymentRequestLinesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentRequestUpdateParamsPaymentRequestLinesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


