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

// PartnerUpdateParamsPaymentTermAttributes struct for PartnerUpdateParamsPaymentTermAttributes
type PartnerUpdateParamsPaymentTermAttributes struct {
	// 支払月
	AdditionalMonths *int32 `json:"additional_months,omitempty"`
	// 締め日（29, 30, 31日の末日を指定する場合は、32を指定してください。）
	CutoffDay *int32 `json:"cutoff_day,omitempty"`
	// 支払日（29, 30, 31日の末日を指定する場合は、32を指定してください。）
	FixedDay *int32 `json:"fixed_day,omitempty"`
}

// NewPartnerUpdateParamsPaymentTermAttributes instantiates a new PartnerUpdateParamsPaymentTermAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerUpdateParamsPaymentTermAttributes() *PartnerUpdateParamsPaymentTermAttributes {
	this := PartnerUpdateParamsPaymentTermAttributes{}
	return &this
}

// NewPartnerUpdateParamsPaymentTermAttributesWithDefaults instantiates a new PartnerUpdateParamsPaymentTermAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerUpdateParamsPaymentTermAttributesWithDefaults() *PartnerUpdateParamsPaymentTermAttributes {
	this := PartnerUpdateParamsPaymentTermAttributes{}
	return &this
}

// GetAdditionalMonths returns the AdditionalMonths field value if set, zero value otherwise.
func (o *PartnerUpdateParamsPaymentTermAttributes) GetAdditionalMonths() int32 {
	if o == nil || o.AdditionalMonths == nil {
		var ret int32
		return ret
	}
	return *o.AdditionalMonths
}

// GetAdditionalMonthsOk returns a tuple with the AdditionalMonths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerUpdateParamsPaymentTermAttributes) GetAdditionalMonthsOk() (*int32, bool) {
	if o == nil || o.AdditionalMonths == nil {
		return nil, false
	}
	return o.AdditionalMonths, true
}

// HasAdditionalMonths returns a boolean if a field has been set.
func (o *PartnerUpdateParamsPaymentTermAttributes) HasAdditionalMonths() bool {
	if o != nil && o.AdditionalMonths != nil {
		return true
	}

	return false
}

// SetAdditionalMonths gets a reference to the given int32 and assigns it to the AdditionalMonths field.
func (o *PartnerUpdateParamsPaymentTermAttributes) SetAdditionalMonths(v int32) {
	o.AdditionalMonths = &v
}

// GetCutoffDay returns the CutoffDay field value if set, zero value otherwise.
func (o *PartnerUpdateParamsPaymentTermAttributes) GetCutoffDay() int32 {
	if o == nil || o.CutoffDay == nil {
		var ret int32
		return ret
	}
	return *o.CutoffDay
}

// GetCutoffDayOk returns a tuple with the CutoffDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerUpdateParamsPaymentTermAttributes) GetCutoffDayOk() (*int32, bool) {
	if o == nil || o.CutoffDay == nil {
		return nil, false
	}
	return o.CutoffDay, true
}

// HasCutoffDay returns a boolean if a field has been set.
func (o *PartnerUpdateParamsPaymentTermAttributes) HasCutoffDay() bool {
	if o != nil && o.CutoffDay != nil {
		return true
	}

	return false
}

// SetCutoffDay gets a reference to the given int32 and assigns it to the CutoffDay field.
func (o *PartnerUpdateParamsPaymentTermAttributes) SetCutoffDay(v int32) {
	o.CutoffDay = &v
}

// GetFixedDay returns the FixedDay field value if set, zero value otherwise.
func (o *PartnerUpdateParamsPaymentTermAttributes) GetFixedDay() int32 {
	if o == nil || o.FixedDay == nil {
		var ret int32
		return ret
	}
	return *o.FixedDay
}

// GetFixedDayOk returns a tuple with the FixedDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerUpdateParamsPaymentTermAttributes) GetFixedDayOk() (*int32, bool) {
	if o == nil || o.FixedDay == nil {
		return nil, false
	}
	return o.FixedDay, true
}

// HasFixedDay returns a boolean if a field has been set.
func (o *PartnerUpdateParamsPaymentTermAttributes) HasFixedDay() bool {
	if o != nil && o.FixedDay != nil {
		return true
	}

	return false
}

// SetFixedDay gets a reference to the given int32 and assigns it to the FixedDay field.
func (o *PartnerUpdateParamsPaymentTermAttributes) SetFixedDay(v int32) {
	o.FixedDay = &v
}

func (o PartnerUpdateParamsPaymentTermAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalMonths != nil {
		toSerialize["additional_months"] = o.AdditionalMonths
	}
	if o.CutoffDay != nil {
		toSerialize["cutoff_day"] = o.CutoffDay
	}
	if o.FixedDay != nil {
		toSerialize["fixed_day"] = o.FixedDay
	}
	return json.Marshal(toSerialize)
}

type NullablePartnerUpdateParamsPaymentTermAttributes struct {
	value *PartnerUpdateParamsPaymentTermAttributes
	isSet bool
}

func (v NullablePartnerUpdateParamsPaymentTermAttributes) Get() *PartnerUpdateParamsPaymentTermAttributes {
	return v.value
}

func (v *NullablePartnerUpdateParamsPaymentTermAttributes) Set(val *PartnerUpdateParamsPaymentTermAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerUpdateParamsPaymentTermAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerUpdateParamsPaymentTermAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerUpdateParamsPaymentTermAttributes(val *PartnerUpdateParamsPaymentTermAttributes) *NullablePartnerUpdateParamsPaymentTermAttributes {
	return &NullablePartnerUpdateParamsPaymentTermAttributes{value: val, isSet: true}
}

func (v NullablePartnerUpdateParamsPaymentTermAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerUpdateParamsPaymentTermAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


