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

// TrialBsResponseTrialBsBalancesInnerSegment1TagsInner struct for TrialBsResponseTrialBsBalancesInnerSegment1TagsInner
type TrialBsResponseTrialBsBalancesInnerSegment1TagsInner struct {
	// 期末残高
	ClosingBalance *int64 `json:"closing_balance,omitempty"`
	// 構成比
	CompositionRatio *float32 `json:"composition_ratio,omitempty"`
	// 貸方金額
	CreditAmount *int64 `json:"credit_amount,omitempty"`
	// 借方金額
	DebitAmount *int64 `json:"debit_amount,omitempty"`
	// セグメント1タグID
	Id int64 `json:"id"`
	// セグメント1タグ名
	Name *string `json:"name,omitempty"`
	// 期首残高
	OpeningBalance *int64 `json:"opening_balance,omitempty"`
}

// NewTrialBsResponseTrialBsBalancesInnerSegment1TagsInner instantiates a new TrialBsResponseTrialBsBalancesInnerSegment1TagsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrialBsResponseTrialBsBalancesInnerSegment1TagsInner(id int64) *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner {
	this := TrialBsResponseTrialBsBalancesInnerSegment1TagsInner{}
	this.Id = id
	return &this
}

// NewTrialBsResponseTrialBsBalancesInnerSegment1TagsInnerWithDefaults instantiates a new TrialBsResponseTrialBsBalancesInnerSegment1TagsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrialBsResponseTrialBsBalancesInnerSegment1TagsInnerWithDefaults() *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner {
	this := TrialBsResponseTrialBsBalancesInnerSegment1TagsInner{}
	return &this
}

// GetClosingBalance returns the ClosingBalance field value if set, zero value otherwise.
func (o *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner) GetClosingBalance() int64 {
	if o == nil || o.ClosingBalance == nil {
		var ret int64
		return ret
	}
	return *o.ClosingBalance
}

// GetClosingBalanceOk returns a tuple with the ClosingBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner) GetClosingBalanceOk() (*int64, bool) {
	if o == nil || o.ClosingBalance == nil {
		return nil, false
	}
	return o.ClosingBalance, true
}

// HasClosingBalance returns a boolean if a field has been set.
func (o *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner) HasClosingBalance() bool {
	if o != nil && o.ClosingBalance != nil {
		return true
	}

	return false
}

// SetClosingBalance gets a reference to the given int64 and assigns it to the ClosingBalance field.
func (o *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner) SetClosingBalance(v int64) {
	o.ClosingBalance = &v
}

// GetCompositionRatio returns the CompositionRatio field value if set, zero value otherwise.
func (o *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner) GetCompositionRatio() float32 {
	if o == nil || o.CompositionRatio == nil {
		var ret float32
		return ret
	}
	return *o.CompositionRatio
}

// GetCompositionRatioOk returns a tuple with the CompositionRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner) GetCompositionRatioOk() (*float32, bool) {
	if o == nil || o.CompositionRatio == nil {
		return nil, false
	}
	return o.CompositionRatio, true
}

// HasCompositionRatio returns a boolean if a field has been set.
func (o *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner) HasCompositionRatio() bool {
	if o != nil && o.CompositionRatio != nil {
		return true
	}

	return false
}

// SetCompositionRatio gets a reference to the given float32 and assigns it to the CompositionRatio field.
func (o *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner) SetCompositionRatio(v float32) {
	o.CompositionRatio = &v
}

// GetCreditAmount returns the CreditAmount field value if set, zero value otherwise.
func (o *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner) GetCreditAmount() int64 {
	if o == nil || o.CreditAmount == nil {
		var ret int64
		return ret
	}
	return *o.CreditAmount
}

// GetCreditAmountOk returns a tuple with the CreditAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner) GetCreditAmountOk() (*int64, bool) {
	if o == nil || o.CreditAmount == nil {
		return nil, false
	}
	return o.CreditAmount, true
}

// HasCreditAmount returns a boolean if a field has been set.
func (o *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner) HasCreditAmount() bool {
	if o != nil && o.CreditAmount != nil {
		return true
	}

	return false
}

// SetCreditAmount gets a reference to the given int64 and assigns it to the CreditAmount field.
func (o *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner) SetCreditAmount(v int64) {
	o.CreditAmount = &v
}

// GetDebitAmount returns the DebitAmount field value if set, zero value otherwise.
func (o *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner) GetDebitAmount() int64 {
	if o == nil || o.DebitAmount == nil {
		var ret int64
		return ret
	}
	return *o.DebitAmount
}

// GetDebitAmountOk returns a tuple with the DebitAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner) GetDebitAmountOk() (*int64, bool) {
	if o == nil || o.DebitAmount == nil {
		return nil, false
	}
	return o.DebitAmount, true
}

// HasDebitAmount returns a boolean if a field has been set.
func (o *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner) HasDebitAmount() bool {
	if o != nil && o.DebitAmount != nil {
		return true
	}

	return false
}

// SetDebitAmount gets a reference to the given int64 and assigns it to the DebitAmount field.
func (o *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner) SetDebitAmount(v int64) {
	o.DebitAmount = &v
}

// GetId returns the Id field value
func (o *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner) SetName(v string) {
	o.Name = &v
}

// GetOpeningBalance returns the OpeningBalance field value if set, zero value otherwise.
func (o *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner) GetOpeningBalance() int64 {
	if o == nil || o.OpeningBalance == nil {
		var ret int64
		return ret
	}
	return *o.OpeningBalance
}

// GetOpeningBalanceOk returns a tuple with the OpeningBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner) GetOpeningBalanceOk() (*int64, bool) {
	if o == nil || o.OpeningBalance == nil {
		return nil, false
	}
	return o.OpeningBalance, true
}

// HasOpeningBalance returns a boolean if a field has been set.
func (o *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner) HasOpeningBalance() bool {
	if o != nil && o.OpeningBalance != nil {
		return true
	}

	return false
}

// SetOpeningBalance gets a reference to the given int64 and assigns it to the OpeningBalance field.
func (o *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner) SetOpeningBalance(v int64) {
	o.OpeningBalance = &v
}

func (o TrialBsResponseTrialBsBalancesInnerSegment1TagsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClosingBalance != nil {
		toSerialize["closing_balance"] = o.ClosingBalance
	}
	if o.CompositionRatio != nil {
		toSerialize["composition_ratio"] = o.CompositionRatio
	}
	if o.CreditAmount != nil {
		toSerialize["credit_amount"] = o.CreditAmount
	}
	if o.DebitAmount != nil {
		toSerialize["debit_amount"] = o.DebitAmount
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OpeningBalance != nil {
		toSerialize["opening_balance"] = o.OpeningBalance
	}
	return json.Marshal(toSerialize)
}

type NullableTrialBsResponseTrialBsBalancesInnerSegment1TagsInner struct {
	value *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner
	isSet bool
}

func (v NullableTrialBsResponseTrialBsBalancesInnerSegment1TagsInner) Get() *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner {
	return v.value
}

func (v *NullableTrialBsResponseTrialBsBalancesInnerSegment1TagsInner) Set(val *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTrialBsResponseTrialBsBalancesInnerSegment1TagsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTrialBsResponseTrialBsBalancesInnerSegment1TagsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrialBsResponseTrialBsBalancesInnerSegment1TagsInner(val *TrialBsResponseTrialBsBalancesInnerSegment1TagsInner) *NullableTrialBsResponseTrialBsBalancesInnerSegment1TagsInner {
	return &NullableTrialBsResponseTrialBsBalancesInnerSegment1TagsInner{value: val, isSet: true}
}

func (v NullableTrialBsResponseTrialBsBalancesInnerSegment1TagsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrialBsResponseTrialBsBalancesInnerSegment1TagsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


