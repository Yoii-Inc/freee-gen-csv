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

// DealUpdateParamsDetailsInner struct for DealUpdateParamsDetailsInner
type DealUpdateParamsDetailsInner struct {
	// 勘定科目ID
	AccountItemId int64 `json:"account_item_id"`
	// 取引金額（税込で指定してください）
	Amount int64 `json:"amount"`
	// 備考
	Description *string `json:"description,omitempty"`
	// 取引行ID: 既存取引行を更新する場合に指定します。IDを指定しない取引行は、新規行として扱われ追加されます。また、detailsに含まれない既存の取引行は削除されます。更新後も残したい行は、必ず取引行IDを指定してdetailsに含めてください。
	Id *int64 `json:"id,omitempty"`
	// 品目ID
	ItemId *int64 `json:"item_id,omitempty"`
	// 部門ID
	SectionId *int64 `json:"section_id,omitempty"`
	// セグメント１ID
	Segment1TagId *int64 `json:"segment_1_tag_id,omitempty"`
	// セグメント２ID
	Segment2TagId *int64 `json:"segment_2_tag_id,omitempty"`
	// セグメント３ID
	Segment3TagId *int64 `json:"segment_3_tag_id,omitempty"`
	// メモタグID
	TagIds []int64 `json:"tag_ids,omitempty"`
	// 税区分コード
	TaxCode int64 `json:"tax_code"`
	// 消費税額（指定しない場合は自動で計算されます）
	Vat *int64 `json:"vat,omitempty"`
}

// NewDealUpdateParamsDetailsInner instantiates a new DealUpdateParamsDetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDealUpdateParamsDetailsInner(accountItemId int64, amount int64, taxCode int64) *DealUpdateParamsDetailsInner {
	this := DealUpdateParamsDetailsInner{}
	this.AccountItemId = accountItemId
	this.Amount = amount
	this.TaxCode = taxCode
	return &this
}

// NewDealUpdateParamsDetailsInnerWithDefaults instantiates a new DealUpdateParamsDetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDealUpdateParamsDetailsInnerWithDefaults() *DealUpdateParamsDetailsInner {
	this := DealUpdateParamsDetailsInner{}
	return &this
}

// GetAccountItemId returns the AccountItemId field value
func (o *DealUpdateParamsDetailsInner) GetAccountItemId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AccountItemId
}

// GetAccountItemIdOk returns a tuple with the AccountItemId field value
// and a boolean to check if the value has been set.
func (o *DealUpdateParamsDetailsInner) GetAccountItemIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountItemId, true
}

// SetAccountItemId sets field value
func (o *DealUpdateParamsDetailsInner) SetAccountItemId(v int64) {
	o.AccountItemId = v
}

// GetAmount returns the Amount field value
func (o *DealUpdateParamsDetailsInner) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *DealUpdateParamsDetailsInner) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *DealUpdateParamsDetailsInner) SetAmount(v int64) {
	o.Amount = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DealUpdateParamsDetailsInner) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealUpdateParamsDetailsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DealUpdateParamsDetailsInner) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DealUpdateParamsDetailsInner) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DealUpdateParamsDetailsInner) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealUpdateParamsDetailsInner) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DealUpdateParamsDetailsInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DealUpdateParamsDetailsInner) SetId(v int64) {
	o.Id = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *DealUpdateParamsDetailsInner) GetItemId() int64 {
	if o == nil || o.ItemId == nil {
		var ret int64
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealUpdateParamsDetailsInner) GetItemIdOk() (*int64, bool) {
	if o == nil || o.ItemId == nil {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *DealUpdateParamsDetailsInner) HasItemId() bool {
	if o != nil && o.ItemId != nil {
		return true
	}

	return false
}

// SetItemId gets a reference to the given int64 and assigns it to the ItemId field.
func (o *DealUpdateParamsDetailsInner) SetItemId(v int64) {
	o.ItemId = &v
}

// GetSectionId returns the SectionId field value if set, zero value otherwise.
func (o *DealUpdateParamsDetailsInner) GetSectionId() int64 {
	if o == nil || o.SectionId == nil {
		var ret int64
		return ret
	}
	return *o.SectionId
}

// GetSectionIdOk returns a tuple with the SectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealUpdateParamsDetailsInner) GetSectionIdOk() (*int64, bool) {
	if o == nil || o.SectionId == nil {
		return nil, false
	}
	return o.SectionId, true
}

// HasSectionId returns a boolean if a field has been set.
func (o *DealUpdateParamsDetailsInner) HasSectionId() bool {
	if o != nil && o.SectionId != nil {
		return true
	}

	return false
}

// SetSectionId gets a reference to the given int64 and assigns it to the SectionId field.
func (o *DealUpdateParamsDetailsInner) SetSectionId(v int64) {
	o.SectionId = &v
}

// GetSegment1TagId returns the Segment1TagId field value if set, zero value otherwise.
func (o *DealUpdateParamsDetailsInner) GetSegment1TagId() int64 {
	if o == nil || o.Segment1TagId == nil {
		var ret int64
		return ret
	}
	return *o.Segment1TagId
}

// GetSegment1TagIdOk returns a tuple with the Segment1TagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealUpdateParamsDetailsInner) GetSegment1TagIdOk() (*int64, bool) {
	if o == nil || o.Segment1TagId == nil {
		return nil, false
	}
	return o.Segment1TagId, true
}

// HasSegment1TagId returns a boolean if a field has been set.
func (o *DealUpdateParamsDetailsInner) HasSegment1TagId() bool {
	if o != nil && o.Segment1TagId != nil {
		return true
	}

	return false
}

// SetSegment1TagId gets a reference to the given int64 and assigns it to the Segment1TagId field.
func (o *DealUpdateParamsDetailsInner) SetSegment1TagId(v int64) {
	o.Segment1TagId = &v
}

// GetSegment2TagId returns the Segment2TagId field value if set, zero value otherwise.
func (o *DealUpdateParamsDetailsInner) GetSegment2TagId() int64 {
	if o == nil || o.Segment2TagId == nil {
		var ret int64
		return ret
	}
	return *o.Segment2TagId
}

// GetSegment2TagIdOk returns a tuple with the Segment2TagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealUpdateParamsDetailsInner) GetSegment2TagIdOk() (*int64, bool) {
	if o == nil || o.Segment2TagId == nil {
		return nil, false
	}
	return o.Segment2TagId, true
}

// HasSegment2TagId returns a boolean if a field has been set.
func (o *DealUpdateParamsDetailsInner) HasSegment2TagId() bool {
	if o != nil && o.Segment2TagId != nil {
		return true
	}

	return false
}

// SetSegment2TagId gets a reference to the given int64 and assigns it to the Segment2TagId field.
func (o *DealUpdateParamsDetailsInner) SetSegment2TagId(v int64) {
	o.Segment2TagId = &v
}

// GetSegment3TagId returns the Segment3TagId field value if set, zero value otherwise.
func (o *DealUpdateParamsDetailsInner) GetSegment3TagId() int64 {
	if o == nil || o.Segment3TagId == nil {
		var ret int64
		return ret
	}
	return *o.Segment3TagId
}

// GetSegment3TagIdOk returns a tuple with the Segment3TagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealUpdateParamsDetailsInner) GetSegment3TagIdOk() (*int64, bool) {
	if o == nil || o.Segment3TagId == nil {
		return nil, false
	}
	return o.Segment3TagId, true
}

// HasSegment3TagId returns a boolean if a field has been set.
func (o *DealUpdateParamsDetailsInner) HasSegment3TagId() bool {
	if o != nil && o.Segment3TagId != nil {
		return true
	}

	return false
}

// SetSegment3TagId gets a reference to the given int64 and assigns it to the Segment3TagId field.
func (o *DealUpdateParamsDetailsInner) SetSegment3TagId(v int64) {
	o.Segment3TagId = &v
}

// GetTagIds returns the TagIds field value if set, zero value otherwise.
func (o *DealUpdateParamsDetailsInner) GetTagIds() []int64 {
	if o == nil || o.TagIds == nil {
		var ret []int64
		return ret
	}
	return o.TagIds
}

// GetTagIdsOk returns a tuple with the TagIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealUpdateParamsDetailsInner) GetTagIdsOk() ([]int64, bool) {
	if o == nil || o.TagIds == nil {
		return nil, false
	}
	return o.TagIds, true
}

// HasTagIds returns a boolean if a field has been set.
func (o *DealUpdateParamsDetailsInner) HasTagIds() bool {
	if o != nil && o.TagIds != nil {
		return true
	}

	return false
}

// SetTagIds gets a reference to the given []int64 and assigns it to the TagIds field.
func (o *DealUpdateParamsDetailsInner) SetTagIds(v []int64) {
	o.TagIds = v
}

// GetTaxCode returns the TaxCode field value
func (o *DealUpdateParamsDetailsInner) GetTaxCode() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TaxCode
}

// GetTaxCodeOk returns a tuple with the TaxCode field value
// and a boolean to check if the value has been set.
func (o *DealUpdateParamsDetailsInner) GetTaxCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxCode, true
}

// SetTaxCode sets field value
func (o *DealUpdateParamsDetailsInner) SetTaxCode(v int64) {
	o.TaxCode = v
}

// GetVat returns the Vat field value if set, zero value otherwise.
func (o *DealUpdateParamsDetailsInner) GetVat() int64 {
	if o == nil || o.Vat == nil {
		var ret int64
		return ret
	}
	return *o.Vat
}

// GetVatOk returns a tuple with the Vat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealUpdateParamsDetailsInner) GetVatOk() (*int64, bool) {
	if o == nil || o.Vat == nil {
		return nil, false
	}
	return o.Vat, true
}

// HasVat returns a boolean if a field has been set.
func (o *DealUpdateParamsDetailsInner) HasVat() bool {
	if o != nil && o.Vat != nil {
		return true
	}

	return false
}

// SetVat gets a reference to the given int64 and assigns it to the Vat field.
func (o *DealUpdateParamsDetailsInner) SetVat(v int64) {
	o.Vat = &v
}

func (o DealUpdateParamsDetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_item_id"] = o.AccountItemId
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
	if true {
		toSerialize["tax_code"] = o.TaxCode
	}
	if o.Vat != nil {
		toSerialize["vat"] = o.Vat
	}
	return json.Marshal(toSerialize)
}

type NullableDealUpdateParamsDetailsInner struct {
	value *DealUpdateParamsDetailsInner
	isSet bool
}

func (v NullableDealUpdateParamsDetailsInner) Get() *DealUpdateParamsDetailsInner {
	return v.value
}

func (v *NullableDealUpdateParamsDetailsInner) Set(val *DealUpdateParamsDetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDealUpdateParamsDetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDealUpdateParamsDetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDealUpdateParamsDetailsInner(val *DealUpdateParamsDetailsInner) *NullableDealUpdateParamsDetailsInner {
	return &NullableDealUpdateParamsDetailsInner{value: val, isSet: true}
}

func (v NullableDealUpdateParamsDetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDealUpdateParamsDetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


