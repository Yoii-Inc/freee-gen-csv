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

// ApprovalRequestFormResponseApprovalRequestFormPartsInner struct for ApprovalRequestFormResponseApprovalRequestFormPartsInner
type ApprovalRequestFormResponseApprovalRequestFormPartsInner struct {
	// 追加説明
	Annotation NullableString `json:"annotation,omitempty"`
	// 項目ID
	Id int32 `json:"id"`
	// 項目名
	Label *string `json:"label,omitempty"`
	// 上限金額
	MaxAmount NullableInt32 `json:"max_amount,omitempty"`
	// 下限金額
	MinAmount NullableInt32 `json:"min_amount,omitempty"`
	// 順序
	Order *int32 `json:"order,omitempty"`
	// 必須かどうか
	Required NullableBool `json:"required,omitempty"`
	// 項目種別 (title: 申請タイトル, single_line: 自由記述形式 1行, multi_line: 自由記述形式 複数行, select: プルダウン, date: 日付, amount: 金額, receipt: 添付ファイル, section: 部門ID, partner: 取引先ID, ninja_sign_document: 契約書（freeeサイン連携）)
	Type *string `json:"type,omitempty"`
	// 選択項目
	Values []ApprovalRequestFormResponseApprovalRequestFormPartsInnerValuesInner `json:"values,omitempty"`
}

// NewApprovalRequestFormResponseApprovalRequestFormPartsInner instantiates a new ApprovalRequestFormResponseApprovalRequestFormPartsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApprovalRequestFormResponseApprovalRequestFormPartsInner(id int32) *ApprovalRequestFormResponseApprovalRequestFormPartsInner {
	this := ApprovalRequestFormResponseApprovalRequestFormPartsInner{}
	this.Id = id
	return &this
}

// NewApprovalRequestFormResponseApprovalRequestFormPartsInnerWithDefaults instantiates a new ApprovalRequestFormResponseApprovalRequestFormPartsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApprovalRequestFormResponseApprovalRequestFormPartsInnerWithDefaults() *ApprovalRequestFormResponseApprovalRequestFormPartsInner {
	this := ApprovalRequestFormResponseApprovalRequestFormPartsInner{}
	return &this
}

// GetAnnotation returns the Annotation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetAnnotation() string {
	if o == nil || o.Annotation.Get() == nil {
		var ret string
		return ret
	}
	return *o.Annotation.Get()
}

// GetAnnotationOk returns a tuple with the Annotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetAnnotationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Annotation.Get(), o.Annotation.IsSet()
}

// HasAnnotation returns a boolean if a field has been set.
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) HasAnnotation() bool {
	if o != nil && o.Annotation.IsSet() {
		return true
	}

	return false
}

// SetAnnotation gets a reference to the given NullableString and assigns it to the Annotation field.
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) SetAnnotation(v string) {
	o.Annotation.Set(&v)
}
// SetAnnotationNil sets the value for Annotation to be an explicit nil
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) SetAnnotationNil() {
	o.Annotation.Set(nil)
}

// UnsetAnnotation ensures that no value is present for Annotation, not even an explicit nil
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) UnsetAnnotation() {
	o.Annotation.Unset()
}

// GetId returns the Id field value
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) SetId(v int32) {
	o.Id = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) SetLabel(v string) {
	o.Label = &v
}

// GetMaxAmount returns the MaxAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetMaxAmount() int32 {
	if o == nil || o.MaxAmount.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MaxAmount.Get()
}

// GetMaxAmountOk returns a tuple with the MaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetMaxAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxAmount.Get(), o.MaxAmount.IsSet()
}

// HasMaxAmount returns a boolean if a field has been set.
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) HasMaxAmount() bool {
	if o != nil && o.MaxAmount.IsSet() {
		return true
	}

	return false
}

// SetMaxAmount gets a reference to the given NullableInt32 and assigns it to the MaxAmount field.
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) SetMaxAmount(v int32) {
	o.MaxAmount.Set(&v)
}
// SetMaxAmountNil sets the value for MaxAmount to be an explicit nil
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) SetMaxAmountNil() {
	o.MaxAmount.Set(nil)
}

// UnsetMaxAmount ensures that no value is present for MaxAmount, not even an explicit nil
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) UnsetMaxAmount() {
	o.MaxAmount.Unset()
}

// GetMinAmount returns the MinAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetMinAmount() int32 {
	if o == nil || o.MinAmount.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MinAmount.Get()
}

// GetMinAmountOk returns a tuple with the MinAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetMinAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinAmount.Get(), o.MinAmount.IsSet()
}

// HasMinAmount returns a boolean if a field has been set.
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) HasMinAmount() bool {
	if o != nil && o.MinAmount.IsSet() {
		return true
	}

	return false
}

// SetMinAmount gets a reference to the given NullableInt32 and assigns it to the MinAmount field.
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) SetMinAmount(v int32) {
	o.MinAmount.Set(&v)
}
// SetMinAmountNil sets the value for MinAmount to be an explicit nil
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) SetMinAmountNil() {
	o.MinAmount.Set(nil)
}

// UnsetMinAmount ensures that no value is present for MinAmount, not even an explicit nil
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) UnsetMinAmount() {
	o.MinAmount.Unset()
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetOrder() int32 {
	if o == nil || o.Order == nil {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetOrderOk() (*int32, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) SetOrder(v int32) {
	o.Order = &v
}

// GetRequired returns the Required field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetRequired() bool {
	if o == nil || o.Required.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Required.Get()
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Required.Get(), o.Required.IsSet()
}

// HasRequired returns a boolean if a field has been set.
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) HasRequired() bool {
	if o != nil && o.Required.IsSet() {
		return true
	}

	return false
}

// SetRequired gets a reference to the given NullableBool and assigns it to the Required field.
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) SetRequired(v bool) {
	o.Required.Set(&v)
}
// SetRequiredNil sets the value for Required to be an explicit nil
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) SetRequiredNil() {
	o.Required.Set(nil)
}

// UnsetRequired ensures that no value is present for Required, not even an explicit nil
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) UnsetRequired() {
	o.Required.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) SetType(v string) {
	o.Type = &v
}

// GetValues returns the Values field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetValues() []ApprovalRequestFormResponseApprovalRequestFormPartsInnerValuesInner {
	if o == nil {
		var ret []ApprovalRequestFormResponseApprovalRequestFormPartsInnerValuesInner
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetValuesOk() ([]ApprovalRequestFormResponseApprovalRequestFormPartsInnerValuesInner, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []ApprovalRequestFormResponseApprovalRequestFormPartsInnerValuesInner and assigns it to the Values field.
func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) SetValues(v []ApprovalRequestFormResponseApprovalRequestFormPartsInnerValuesInner) {
	o.Values = v
}

func (o ApprovalRequestFormResponseApprovalRequestFormPartsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Annotation.IsSet() {
		toSerialize["annotation"] = o.Annotation.Get()
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.MaxAmount.IsSet() {
		toSerialize["max_amount"] = o.MaxAmount.Get()
	}
	if o.MinAmount.IsSet() {
		toSerialize["min_amount"] = o.MinAmount.Get()
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.Required.IsSet() {
		toSerialize["required"] = o.Required.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableApprovalRequestFormResponseApprovalRequestFormPartsInner struct {
	value *ApprovalRequestFormResponseApprovalRequestFormPartsInner
	isSet bool
}

func (v NullableApprovalRequestFormResponseApprovalRequestFormPartsInner) Get() *ApprovalRequestFormResponseApprovalRequestFormPartsInner {
	return v.value
}

func (v *NullableApprovalRequestFormResponseApprovalRequestFormPartsInner) Set(val *ApprovalRequestFormResponseApprovalRequestFormPartsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalRequestFormResponseApprovalRequestFormPartsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalRequestFormResponseApprovalRequestFormPartsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalRequestFormResponseApprovalRequestFormPartsInner(val *ApprovalRequestFormResponseApprovalRequestFormPartsInner) *NullableApprovalRequestFormResponseApprovalRequestFormPartsInner {
	return &NullableApprovalRequestFormResponseApprovalRequestFormPartsInner{value: val, isSet: true}
}

func (v NullableApprovalRequestFormResponseApprovalRequestFormPartsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalRequestFormResponseApprovalRequestFormPartsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


