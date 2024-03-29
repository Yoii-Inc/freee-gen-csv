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

// InvoiceIndexResponseInvoicesInnerInvoiceContentsInner struct for InvoiceIndexResponseInvoicesInnerInvoiceContentsInner
type InvoiceIndexResponseInvoicesInnerInvoiceContentsInner struct {
	// 勘定科目ID
	AccountItemId NullableInt64 `json:"account_item_id"`
	// 勘定科目名
	AccountItemName NullableString `json:"account_item_name"`
	// 内税/外税の判別とamountの税込み、税抜きについて <ul> <li>tax_entry_methodがexclusive (外税)の場合</li> <ul> <li>amount: 消費税抜きの金額</li> <li>vat: 消費税の金額</li> </ul> <li>tax_entry_methodがinclusive (内税)の場合</li> <ul> <li>amount: 消費税込みの金額</li> <li>vat: 消費税の金額</li> </ul> </ul> 
	Amount int64 `json:"amount"`
	// 備考
	Description NullableString `json:"description"`
	// 請求内容ID
	Id int64 `json:"id"`
	// 品目ID
	ItemId NullableInt64 `json:"item_id"`
	// 品目
	ItemName NullableString `json:"item_name"`
	// 順序
	Order NullableInt64 `json:"order"`
	// 数量
	Qty float32 `json:"qty"`
	// 軽減税率税区分（true: 対象、false: 対象外）
	ReducedVat bool `json:"reduced_vat"`
	// 部門ID
	SectionId NullableInt64 `json:"section_id"`
	// 部門
	SectionName NullableString `json:"section_name"`
	// セグメント１ID
	Segment1TagId NullableInt64 `json:"segment_1_tag_id,omitempty"`
	// セグメント１
	Segment1TagName NullableString `json:"segment_1_tag_name,omitempty"`
	// セグメント２ID
	Segment2TagId NullableInt64 `json:"segment_2_tag_id,omitempty"`
	// セグメント２
	Segment2TagName NullableString `json:"segment_2_tag_name,omitempty"`
	// セグメント３ID
	Segment3TagId NullableInt64 `json:"segment_3_tag_id,omitempty"`
	// セグメント３
	Segment3TagName NullableString `json:"segment_3_tag_name,omitempty"`
	TagIds []int64 `json:"tag_ids"`
	TagNames []string `json:"tag_names"`
	// 税区分コード
	TaxCode NullableInt64 `json:"tax_code"`
	// 行の種類
	Type string `json:"type"`
	// 単位
	Unit NullableString `json:"unit"`
	// 単価
	UnitPrice float32 `json:"unit_price"`
	// 消費税額
	Vat int64 `json:"vat"`
}

// NewInvoiceIndexResponseInvoicesInnerInvoiceContentsInner instantiates a new InvoiceIndexResponseInvoicesInnerInvoiceContentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceIndexResponseInvoicesInnerInvoiceContentsInner(accountItemId NullableInt64, accountItemName NullableString, amount int64, description NullableString, id int64, itemId NullableInt64, itemName NullableString, order NullableInt64, qty float32, reducedVat bool, sectionId NullableInt64, sectionName NullableString, tagIds []int64, tagNames []string, taxCode NullableInt64, type_ string, unit NullableString, unitPrice float32, vat int64) *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner {
	this := InvoiceIndexResponseInvoicesInnerInvoiceContentsInner{}
	this.AccountItemId = accountItemId
	this.AccountItemName = accountItemName
	this.Amount = amount
	this.Description = description
	this.Id = id
	this.ItemId = itemId
	this.ItemName = itemName
	this.Order = order
	this.Qty = qty
	this.ReducedVat = reducedVat
	this.SectionId = sectionId
	this.SectionName = sectionName
	this.TagIds = tagIds
	this.TagNames = tagNames
	this.TaxCode = taxCode
	this.Type = type_
	this.Unit = unit
	this.UnitPrice = unitPrice
	this.Vat = vat
	return &this
}

// NewInvoiceIndexResponseInvoicesInnerInvoiceContentsInnerWithDefaults instantiates a new InvoiceIndexResponseInvoicesInnerInvoiceContentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceIndexResponseInvoicesInnerInvoiceContentsInnerWithDefaults() *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner {
	this := InvoiceIndexResponseInvoicesInnerInvoiceContentsInner{}
	return &this
}

// GetAccountItemId returns the AccountItemId field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetAccountItemId() int64 {
	if o == nil || o.AccountItemId.Get() == nil {
		var ret int64
		return ret
	}

	return *o.AccountItemId.Get()
}

// GetAccountItemIdOk returns a tuple with the AccountItemId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetAccountItemIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountItemId.Get(), o.AccountItemId.IsSet()
}

// SetAccountItemId sets field value
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetAccountItemId(v int64) {
	o.AccountItemId.Set(&v)
}

// GetAccountItemName returns the AccountItemName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetAccountItemName() string {
	if o == nil || o.AccountItemName.Get() == nil {
		var ret string
		return ret
	}

	return *o.AccountItemName.Get()
}

// GetAccountItemNameOk returns a tuple with the AccountItemName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetAccountItemNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountItemName.Get(), o.AccountItemName.IsSet()
}

// SetAccountItemName sets field value
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetAccountItemName(v string) {
	o.AccountItemName.Set(&v)
}

// GetAmount returns the Amount field value
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetAmount(v int64) {
	o.Amount = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetId returns the Id field value
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetId(v int64) {
	o.Id = v
}

// GetItemId returns the ItemId field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetItemId() int64 {
	if o == nil || o.ItemId.Get() == nil {
		var ret int64
		return ret
	}

	return *o.ItemId.Get()
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetItemIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemId.Get(), o.ItemId.IsSet()
}

// SetItemId sets field value
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetItemId(v int64) {
	o.ItemId.Set(&v)
}

// GetItemName returns the ItemName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetItemName() string {
	if o == nil || o.ItemName.Get() == nil {
		var ret string
		return ret
	}

	return *o.ItemName.Get()
}

// GetItemNameOk returns a tuple with the ItemName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetItemNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemName.Get(), o.ItemName.IsSet()
}

// SetItemName sets field value
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetItemName(v string) {
	o.ItemName.Set(&v)
}

// GetOrder returns the Order field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetOrder() int64 {
	if o == nil || o.Order.Get() == nil {
		var ret int64
		return ret
	}

	return *o.Order.Get()
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetOrderOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Order.Get(), o.Order.IsSet()
}

// SetOrder sets field value
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetOrder(v int64) {
	o.Order.Set(&v)
}

// GetQty returns the Qty field value
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetQty() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Qty
}

// GetQtyOk returns a tuple with the Qty field value
// and a boolean to check if the value has been set.
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetQtyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Qty, true
}

// SetQty sets field value
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetQty(v float32) {
	o.Qty = v
}

// GetReducedVat returns the ReducedVat field value
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetReducedVat() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ReducedVat
}

// GetReducedVatOk returns a tuple with the ReducedVat field value
// and a boolean to check if the value has been set.
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetReducedVatOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReducedVat, true
}

// SetReducedVat sets field value
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetReducedVat(v bool) {
	o.ReducedVat = v
}

// GetSectionId returns the SectionId field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSectionId() int64 {
	if o == nil || o.SectionId.Get() == nil {
		var ret int64
		return ret
	}

	return *o.SectionId.Get()
}

// GetSectionIdOk returns a tuple with the SectionId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSectionIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SectionId.Get(), o.SectionId.IsSet()
}

// SetSectionId sets field value
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetSectionId(v int64) {
	o.SectionId.Set(&v)
}

// GetSectionName returns the SectionName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSectionName() string {
	if o == nil || o.SectionName.Get() == nil {
		var ret string
		return ret
	}

	return *o.SectionName.Get()
}

// GetSectionNameOk returns a tuple with the SectionName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSectionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SectionName.Get(), o.SectionName.IsSet()
}

// SetSectionName sets field value
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetSectionName(v string) {
	o.SectionName.Set(&v)
}

// GetSegment1TagId returns the Segment1TagId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSegment1TagId() int64 {
	if o == nil || o.Segment1TagId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Segment1TagId.Get()
}

// GetSegment1TagIdOk returns a tuple with the Segment1TagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSegment1TagIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Segment1TagId.Get(), o.Segment1TagId.IsSet()
}

// HasSegment1TagId returns a boolean if a field has been set.
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) HasSegment1TagId() bool {
	if o != nil && o.Segment1TagId.IsSet() {
		return true
	}

	return false
}

// SetSegment1TagId gets a reference to the given NullableInt64 and assigns it to the Segment1TagId field.
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetSegment1TagId(v int64) {
	o.Segment1TagId.Set(&v)
}
// SetSegment1TagIdNil sets the value for Segment1TagId to be an explicit nil
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetSegment1TagIdNil() {
	o.Segment1TagId.Set(nil)
}

// UnsetSegment1TagId ensures that no value is present for Segment1TagId, not even an explicit nil
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) UnsetSegment1TagId() {
	o.Segment1TagId.Unset()
}

// GetSegment1TagName returns the Segment1TagName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSegment1TagName() string {
	if o == nil || o.Segment1TagName.Get() == nil {
		var ret string
		return ret
	}
	return *o.Segment1TagName.Get()
}

// GetSegment1TagNameOk returns a tuple with the Segment1TagName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSegment1TagNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Segment1TagName.Get(), o.Segment1TagName.IsSet()
}

// HasSegment1TagName returns a boolean if a field has been set.
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) HasSegment1TagName() bool {
	if o != nil && o.Segment1TagName.IsSet() {
		return true
	}

	return false
}

// SetSegment1TagName gets a reference to the given NullableString and assigns it to the Segment1TagName field.
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetSegment1TagName(v string) {
	o.Segment1TagName.Set(&v)
}
// SetSegment1TagNameNil sets the value for Segment1TagName to be an explicit nil
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetSegment1TagNameNil() {
	o.Segment1TagName.Set(nil)
}

// UnsetSegment1TagName ensures that no value is present for Segment1TagName, not even an explicit nil
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) UnsetSegment1TagName() {
	o.Segment1TagName.Unset()
}

// GetSegment2TagId returns the Segment2TagId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSegment2TagId() int64 {
	if o == nil || o.Segment2TagId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Segment2TagId.Get()
}

// GetSegment2TagIdOk returns a tuple with the Segment2TagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSegment2TagIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Segment2TagId.Get(), o.Segment2TagId.IsSet()
}

// HasSegment2TagId returns a boolean if a field has been set.
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) HasSegment2TagId() bool {
	if o != nil && o.Segment2TagId.IsSet() {
		return true
	}

	return false
}

// SetSegment2TagId gets a reference to the given NullableInt64 and assigns it to the Segment2TagId field.
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetSegment2TagId(v int64) {
	o.Segment2TagId.Set(&v)
}
// SetSegment2TagIdNil sets the value for Segment2TagId to be an explicit nil
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetSegment2TagIdNil() {
	o.Segment2TagId.Set(nil)
}

// UnsetSegment2TagId ensures that no value is present for Segment2TagId, not even an explicit nil
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) UnsetSegment2TagId() {
	o.Segment2TagId.Unset()
}

// GetSegment2TagName returns the Segment2TagName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSegment2TagName() string {
	if o == nil || o.Segment2TagName.Get() == nil {
		var ret string
		return ret
	}
	return *o.Segment2TagName.Get()
}

// GetSegment2TagNameOk returns a tuple with the Segment2TagName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSegment2TagNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Segment2TagName.Get(), o.Segment2TagName.IsSet()
}

// HasSegment2TagName returns a boolean if a field has been set.
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) HasSegment2TagName() bool {
	if o != nil && o.Segment2TagName.IsSet() {
		return true
	}

	return false
}

// SetSegment2TagName gets a reference to the given NullableString and assigns it to the Segment2TagName field.
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetSegment2TagName(v string) {
	o.Segment2TagName.Set(&v)
}
// SetSegment2TagNameNil sets the value for Segment2TagName to be an explicit nil
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetSegment2TagNameNil() {
	o.Segment2TagName.Set(nil)
}

// UnsetSegment2TagName ensures that no value is present for Segment2TagName, not even an explicit nil
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) UnsetSegment2TagName() {
	o.Segment2TagName.Unset()
}

// GetSegment3TagId returns the Segment3TagId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSegment3TagId() int64 {
	if o == nil || o.Segment3TagId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Segment3TagId.Get()
}

// GetSegment3TagIdOk returns a tuple with the Segment3TagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSegment3TagIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Segment3TagId.Get(), o.Segment3TagId.IsSet()
}

// HasSegment3TagId returns a boolean if a field has been set.
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) HasSegment3TagId() bool {
	if o != nil && o.Segment3TagId.IsSet() {
		return true
	}

	return false
}

// SetSegment3TagId gets a reference to the given NullableInt64 and assigns it to the Segment3TagId field.
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetSegment3TagId(v int64) {
	o.Segment3TagId.Set(&v)
}
// SetSegment3TagIdNil sets the value for Segment3TagId to be an explicit nil
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetSegment3TagIdNil() {
	o.Segment3TagId.Set(nil)
}

// UnsetSegment3TagId ensures that no value is present for Segment3TagId, not even an explicit nil
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) UnsetSegment3TagId() {
	o.Segment3TagId.Unset()
}

// GetSegment3TagName returns the Segment3TagName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSegment3TagName() string {
	if o == nil || o.Segment3TagName.Get() == nil {
		var ret string
		return ret
	}
	return *o.Segment3TagName.Get()
}

// GetSegment3TagNameOk returns a tuple with the Segment3TagName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSegment3TagNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Segment3TagName.Get(), o.Segment3TagName.IsSet()
}

// HasSegment3TagName returns a boolean if a field has been set.
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) HasSegment3TagName() bool {
	if o != nil && o.Segment3TagName.IsSet() {
		return true
	}

	return false
}

// SetSegment3TagName gets a reference to the given NullableString and assigns it to the Segment3TagName field.
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetSegment3TagName(v string) {
	o.Segment3TagName.Set(&v)
}
// SetSegment3TagNameNil sets the value for Segment3TagName to be an explicit nil
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetSegment3TagNameNil() {
	o.Segment3TagName.Set(nil)
}

// UnsetSegment3TagName ensures that no value is present for Segment3TagName, not even an explicit nil
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) UnsetSegment3TagName() {
	o.Segment3TagName.Unset()
}

// GetTagIds returns the TagIds field value
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetTagIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.TagIds
}

// GetTagIdsOk returns a tuple with the TagIds field value
// and a boolean to check if the value has been set.
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetTagIdsOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TagIds, true
}

// SetTagIds sets field value
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetTagIds(v []int64) {
	o.TagIds = v
}

// GetTagNames returns the TagNames field value
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetTagNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TagNames
}

// GetTagNamesOk returns a tuple with the TagNames field value
// and a boolean to check if the value has been set.
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetTagNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TagNames, true
}

// SetTagNames sets field value
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetTagNames(v []string) {
	o.TagNames = v
}

// GetTaxCode returns the TaxCode field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetTaxCode() int64 {
	if o == nil || o.TaxCode.Get() == nil {
		var ret int64
		return ret
	}

	return *o.TaxCode.Get()
}

// GetTaxCodeOk returns a tuple with the TaxCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetTaxCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxCode.Get(), o.TaxCode.IsSet()
}

// SetTaxCode sets field value
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetTaxCode(v int64) {
	o.TaxCode.Set(&v)
}

// GetType returns the Type field value
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetType(v string) {
	o.Type = v
}

// GetUnit returns the Unit field value
// If the value is explicit nil, the zero value for string will be returned
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetUnit() string {
	if o == nil || o.Unit.Get() == nil {
		var ret string
		return ret
	}

	return *o.Unit.Get()
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Unit.Get(), o.Unit.IsSet()
}

// SetUnit sets field value
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetUnit(v string) {
	o.Unit.Set(&v)
}

// GetUnitPrice returns the UnitPrice field value
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetUnitPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UnitPrice
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value
// and a boolean to check if the value has been set.
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetUnitPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitPrice, true
}

// SetUnitPrice sets field value
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetUnitPrice(v float32) {
	o.UnitPrice = v
}

// GetVat returns the Vat field value
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetVat() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Vat
}

// GetVatOk returns a tuple with the Vat field value
// and a boolean to check if the value has been set.
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetVatOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vat, true
}

// SetVat sets field value
func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetVat(v int64) {
	o.Vat = v
}

func (o InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_item_id"] = o.AccountItemId.Get()
	}
	if true {
		toSerialize["account_item_name"] = o.AccountItemName.Get()
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["description"] = o.Description.Get()
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["item_id"] = o.ItemId.Get()
	}
	if true {
		toSerialize["item_name"] = o.ItemName.Get()
	}
	if true {
		toSerialize["order"] = o.Order.Get()
	}
	if true {
		toSerialize["qty"] = o.Qty
	}
	if true {
		toSerialize["reduced_vat"] = o.ReducedVat
	}
	if true {
		toSerialize["section_id"] = o.SectionId.Get()
	}
	if true {
		toSerialize["section_name"] = o.SectionName.Get()
	}
	if o.Segment1TagId.IsSet() {
		toSerialize["segment_1_tag_id"] = o.Segment1TagId.Get()
	}
	if o.Segment1TagName.IsSet() {
		toSerialize["segment_1_tag_name"] = o.Segment1TagName.Get()
	}
	if o.Segment2TagId.IsSet() {
		toSerialize["segment_2_tag_id"] = o.Segment2TagId.Get()
	}
	if o.Segment2TagName.IsSet() {
		toSerialize["segment_2_tag_name"] = o.Segment2TagName.Get()
	}
	if o.Segment3TagId.IsSet() {
		toSerialize["segment_3_tag_id"] = o.Segment3TagId.Get()
	}
	if o.Segment3TagName.IsSet() {
		toSerialize["segment_3_tag_name"] = o.Segment3TagName.Get()
	}
	if true {
		toSerialize["tag_ids"] = o.TagIds
	}
	if true {
		toSerialize["tag_names"] = o.TagNames
	}
	if true {
		toSerialize["tax_code"] = o.TaxCode.Get()
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["unit"] = o.Unit.Get()
	}
	if true {
		toSerialize["unit_price"] = o.UnitPrice
	}
	if true {
		toSerialize["vat"] = o.Vat
	}
	return json.Marshal(toSerialize)
}

type NullableInvoiceIndexResponseInvoicesInnerInvoiceContentsInner struct {
	value *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner
	isSet bool
}

func (v NullableInvoiceIndexResponseInvoicesInnerInvoiceContentsInner) Get() *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner {
	return v.value
}

func (v *NullableInvoiceIndexResponseInvoicesInnerInvoiceContentsInner) Set(val *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceIndexResponseInvoicesInnerInvoiceContentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceIndexResponseInvoicesInnerInvoiceContentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceIndexResponseInvoicesInnerInvoiceContentsInner(val *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) *NullableInvoiceIndexResponseInvoicesInnerInvoiceContentsInner {
	return &NullableInvoiceIndexResponseInvoicesInnerInvoiceContentsInner{value: val, isSet: true}
}

func (v NullableInvoiceIndexResponseInvoicesInnerInvoiceContentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceIndexResponseInvoicesInnerInvoiceContentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


