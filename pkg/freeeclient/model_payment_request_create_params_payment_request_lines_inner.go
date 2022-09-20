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

// PaymentRequestCreateParamsPaymentRequestLinesInner struct for PaymentRequestCreateParamsPaymentRequestLinesInner
type PaymentRequestCreateParamsPaymentRequestLinesInner struct {
	// 勘定科目ID
	AccountItemId NullableInt64 `json:"account_item_id,omitempty"`
	// 金額
	Amount int64 `json:"amount"`
	// 内容
	Description *string `json:"description,omitempty"`
	// 品目ID
	ItemId *int64 `json:"item_id,omitempty"`
	// '行の種類 (deal_line: 支払依頼, withholding_tax: 源泉徴収税)'<br> 'デフォルトは deal_line: 支払依頼 です' 
	LineType *string `json:"line_type,omitempty"`
	// 部門ID
	SectionId *int64 `json:"section_id,omitempty"`
	// セグメント１ID<br> セグメントタグ一覧APIを利用して取得してください。<br> <a href=\"https://support.freee.co.jp/hc/ja/articles/360020679611\" target=\"_blank\">セグメント（分析用タグ）の設定</a><br> 
	Segment1TagId *int64 `json:"segment_1_tag_id,omitempty"`
	// セグメント２ID(法人向けエンタープライズプラン)<br> セグメントタグ一覧APIを利用して取得してください。<br> <a href=\"https://support.freee.co.jp/hc/ja/articles/360020679611\" target=\"_blank\">セグメント（分析用タグ）の設定</a><br> 
	Segment2TagId *int64 `json:"segment_2_tag_id,omitempty"`
	// セグメント３ID(法人向けエンタープライズプラン)<br> セグメントタグ一覧APIを利用して取得してください。<br> <a href=\"https://support.freee.co.jp/hc/ja/articles/360020679611\" target=\"_blank\">セグメント（分析用タグ）の設定</a><br> 
	Segment3TagId *int64 `json:"segment_3_tag_id,omitempty"`
	// メモタグID
	TagIds []int64 `json:"tag_ids,omitempty"`
	// 税区分コード<br> 勘定科目IDを指定する場合は必須です。 
	TaxCode *int64 `json:"tax_code,omitempty"`
}

// NewPaymentRequestCreateParamsPaymentRequestLinesInner instantiates a new PaymentRequestCreateParamsPaymentRequestLinesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentRequestCreateParamsPaymentRequestLinesInner(amount int64) *PaymentRequestCreateParamsPaymentRequestLinesInner {
	this := PaymentRequestCreateParamsPaymentRequestLinesInner{}
	this.Amount = amount
	return &this
}

// NewPaymentRequestCreateParamsPaymentRequestLinesInnerWithDefaults instantiates a new PaymentRequestCreateParamsPaymentRequestLinesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentRequestCreateParamsPaymentRequestLinesInnerWithDefaults() *PaymentRequestCreateParamsPaymentRequestLinesInner {
	this := PaymentRequestCreateParamsPaymentRequestLinesInner{}
	return &this
}

// GetAccountItemId returns the AccountItemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetAccountItemId() int64 {
	if o == nil || o.AccountItemId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AccountItemId.Get()
}

// GetAccountItemIdOk returns a tuple with the AccountItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetAccountItemIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountItemId.Get(), o.AccountItemId.IsSet()
}

// HasAccountItemId returns a boolean if a field has been set.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) HasAccountItemId() bool {
	if o != nil && o.AccountItemId.IsSet() {
		return true
	}

	return false
}

// SetAccountItemId gets a reference to the given NullableInt64 and assigns it to the AccountItemId field.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) SetAccountItemId(v int64) {
	o.AccountItemId.Set(&v)
}
// SetAccountItemIdNil sets the value for AccountItemId to be an explicit nil
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) SetAccountItemIdNil() {
	o.AccountItemId.Set(nil)
}

// UnsetAccountItemId ensures that no value is present for AccountItemId, not even an explicit nil
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) UnsetAccountItemId() {
	o.AccountItemId.Unset()
}

// GetAmount returns the Amount field value
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) SetAmount(v int64) {
	o.Amount = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) SetDescription(v string) {
	o.Description = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetItemId() int64 {
	if o == nil || o.ItemId == nil {
		var ret int64
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetItemIdOk() (*int64, bool) {
	if o == nil || o.ItemId == nil {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) HasItemId() bool {
	if o != nil && o.ItemId != nil {
		return true
	}

	return false
}

// SetItemId gets a reference to the given int64 and assigns it to the ItemId field.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) SetItemId(v int64) {
	o.ItemId = &v
}

// GetLineType returns the LineType field value if set, zero value otherwise.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetLineType() string {
	if o == nil || o.LineType == nil {
		var ret string
		return ret
	}
	return *o.LineType
}

// GetLineTypeOk returns a tuple with the LineType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetLineTypeOk() (*string, bool) {
	if o == nil || o.LineType == nil {
		return nil, false
	}
	return o.LineType, true
}

// HasLineType returns a boolean if a field has been set.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) HasLineType() bool {
	if o != nil && o.LineType != nil {
		return true
	}

	return false
}

// SetLineType gets a reference to the given string and assigns it to the LineType field.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) SetLineType(v string) {
	o.LineType = &v
}

// GetSectionId returns the SectionId field value if set, zero value otherwise.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetSectionId() int64 {
	if o == nil || o.SectionId == nil {
		var ret int64
		return ret
	}
	return *o.SectionId
}

// GetSectionIdOk returns a tuple with the SectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetSectionIdOk() (*int64, bool) {
	if o == nil || o.SectionId == nil {
		return nil, false
	}
	return o.SectionId, true
}

// HasSectionId returns a boolean if a field has been set.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) HasSectionId() bool {
	if o != nil && o.SectionId != nil {
		return true
	}

	return false
}

// SetSectionId gets a reference to the given int64 and assigns it to the SectionId field.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) SetSectionId(v int64) {
	o.SectionId = &v
}

// GetSegment1TagId returns the Segment1TagId field value if set, zero value otherwise.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetSegment1TagId() int64 {
	if o == nil || o.Segment1TagId == nil {
		var ret int64
		return ret
	}
	return *o.Segment1TagId
}

// GetSegment1TagIdOk returns a tuple with the Segment1TagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetSegment1TagIdOk() (*int64, bool) {
	if o == nil || o.Segment1TagId == nil {
		return nil, false
	}
	return o.Segment1TagId, true
}

// HasSegment1TagId returns a boolean if a field has been set.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) HasSegment1TagId() bool {
	if o != nil && o.Segment1TagId != nil {
		return true
	}

	return false
}

// SetSegment1TagId gets a reference to the given int64 and assigns it to the Segment1TagId field.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) SetSegment1TagId(v int64) {
	o.Segment1TagId = &v
}

// GetSegment2TagId returns the Segment2TagId field value if set, zero value otherwise.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetSegment2TagId() int64 {
	if o == nil || o.Segment2TagId == nil {
		var ret int64
		return ret
	}
	return *o.Segment2TagId
}

// GetSegment2TagIdOk returns a tuple with the Segment2TagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetSegment2TagIdOk() (*int64, bool) {
	if o == nil || o.Segment2TagId == nil {
		return nil, false
	}
	return o.Segment2TagId, true
}

// HasSegment2TagId returns a boolean if a field has been set.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) HasSegment2TagId() bool {
	if o != nil && o.Segment2TagId != nil {
		return true
	}

	return false
}

// SetSegment2TagId gets a reference to the given int64 and assigns it to the Segment2TagId field.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) SetSegment2TagId(v int64) {
	o.Segment2TagId = &v
}

// GetSegment3TagId returns the Segment3TagId field value if set, zero value otherwise.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetSegment3TagId() int64 {
	if o == nil || o.Segment3TagId == nil {
		var ret int64
		return ret
	}
	return *o.Segment3TagId
}

// GetSegment3TagIdOk returns a tuple with the Segment3TagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetSegment3TagIdOk() (*int64, bool) {
	if o == nil || o.Segment3TagId == nil {
		return nil, false
	}
	return o.Segment3TagId, true
}

// HasSegment3TagId returns a boolean if a field has been set.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) HasSegment3TagId() bool {
	if o != nil && o.Segment3TagId != nil {
		return true
	}

	return false
}

// SetSegment3TagId gets a reference to the given int64 and assigns it to the Segment3TagId field.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) SetSegment3TagId(v int64) {
	o.Segment3TagId = &v
}

// GetTagIds returns the TagIds field value if set, zero value otherwise.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetTagIds() []int64 {
	if o == nil || o.TagIds == nil {
		var ret []int64
		return ret
	}
	return o.TagIds
}

// GetTagIdsOk returns a tuple with the TagIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetTagIdsOk() ([]int64, bool) {
	if o == nil || o.TagIds == nil {
		return nil, false
	}
	return o.TagIds, true
}

// HasTagIds returns a boolean if a field has been set.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) HasTagIds() bool {
	if o != nil && o.TagIds != nil {
		return true
	}

	return false
}

// SetTagIds gets a reference to the given []int64 and assigns it to the TagIds field.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) SetTagIds(v []int64) {
	o.TagIds = v
}

// GetTaxCode returns the TaxCode field value if set, zero value otherwise.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetTaxCode() int64 {
	if o == nil || o.TaxCode == nil {
		var ret int64
		return ret
	}
	return *o.TaxCode
}

// GetTaxCodeOk returns a tuple with the TaxCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetTaxCodeOk() (*int64, bool) {
	if o == nil || o.TaxCode == nil {
		return nil, false
	}
	return o.TaxCode, true
}

// HasTaxCode returns a boolean if a field has been set.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) HasTaxCode() bool {
	if o != nil && o.TaxCode != nil {
		return true
	}

	return false
}

// SetTaxCode gets a reference to the given int64 and assigns it to the TaxCode field.
func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) SetTaxCode(v int64) {
	o.TaxCode = &v
}

func (o PaymentRequestCreateParamsPaymentRequestLinesInner) MarshalJSON() ([]byte, error) {
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

type NullablePaymentRequestCreateParamsPaymentRequestLinesInner struct {
	value *PaymentRequestCreateParamsPaymentRequestLinesInner
	isSet bool
}

func (v NullablePaymentRequestCreateParamsPaymentRequestLinesInner) Get() *PaymentRequestCreateParamsPaymentRequestLinesInner {
	return v.value
}

func (v *NullablePaymentRequestCreateParamsPaymentRequestLinesInner) Set(val *PaymentRequestCreateParamsPaymentRequestLinesInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentRequestCreateParamsPaymentRequestLinesInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentRequestCreateParamsPaymentRequestLinesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentRequestCreateParamsPaymentRequestLinesInner(val *PaymentRequestCreateParamsPaymentRequestLinesInner) *NullablePaymentRequestCreateParamsPaymentRequestLinesInner {
	return &NullablePaymentRequestCreateParamsPaymentRequestLinesInner{value: val, isSet: true}
}

func (v NullablePaymentRequestCreateParamsPaymentRequestLinesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentRequestCreateParamsPaymentRequestLinesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


