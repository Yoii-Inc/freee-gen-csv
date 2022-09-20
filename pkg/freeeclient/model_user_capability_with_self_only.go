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

// UserCapabilityWithSelfOnly struct for UserCapabilityWithSelfOnly
type UserCapabilityWithSelfOnly struct {
	// 「自分のみ」がonになっている場合はself_only、offになっている場合はallになります。
	AllowedTarget *string `json:"allowed_target,omitempty"`
	// 作成
	Create *bool `json:"create,omitempty"`
	// 削除
	Destroy *bool `json:"destroy,omitempty"`
	// 閲覧
	Read *bool `json:"read,omitempty"`
	// 更新
	Update *bool `json:"update,omitempty"`
}

// NewUserCapabilityWithSelfOnly instantiates a new UserCapabilityWithSelfOnly object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCapabilityWithSelfOnly() *UserCapabilityWithSelfOnly {
	this := UserCapabilityWithSelfOnly{}
	return &this
}

// NewUserCapabilityWithSelfOnlyWithDefaults instantiates a new UserCapabilityWithSelfOnly object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCapabilityWithSelfOnlyWithDefaults() *UserCapabilityWithSelfOnly {
	this := UserCapabilityWithSelfOnly{}
	return &this
}

// GetAllowedTarget returns the AllowedTarget field value if set, zero value otherwise.
func (o *UserCapabilityWithSelfOnly) GetAllowedTarget() string {
	if o == nil || o.AllowedTarget == nil {
		var ret string
		return ret
	}
	return *o.AllowedTarget
}

// GetAllowedTargetOk returns a tuple with the AllowedTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCapabilityWithSelfOnly) GetAllowedTargetOk() (*string, bool) {
	if o == nil || o.AllowedTarget == nil {
		return nil, false
	}
	return o.AllowedTarget, true
}

// HasAllowedTarget returns a boolean if a field has been set.
func (o *UserCapabilityWithSelfOnly) HasAllowedTarget() bool {
	if o != nil && o.AllowedTarget != nil {
		return true
	}

	return false
}

// SetAllowedTarget gets a reference to the given string and assigns it to the AllowedTarget field.
func (o *UserCapabilityWithSelfOnly) SetAllowedTarget(v string) {
	o.AllowedTarget = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *UserCapabilityWithSelfOnly) GetCreate() bool {
	if o == nil || o.Create == nil {
		var ret bool
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCapabilityWithSelfOnly) GetCreateOk() (*bool, bool) {
	if o == nil || o.Create == nil {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *UserCapabilityWithSelfOnly) HasCreate() bool {
	if o != nil && o.Create != nil {
		return true
	}

	return false
}

// SetCreate gets a reference to the given bool and assigns it to the Create field.
func (o *UserCapabilityWithSelfOnly) SetCreate(v bool) {
	o.Create = &v
}

// GetDestroy returns the Destroy field value if set, zero value otherwise.
func (o *UserCapabilityWithSelfOnly) GetDestroy() bool {
	if o == nil || o.Destroy == nil {
		var ret bool
		return ret
	}
	return *o.Destroy
}

// GetDestroyOk returns a tuple with the Destroy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCapabilityWithSelfOnly) GetDestroyOk() (*bool, bool) {
	if o == nil || o.Destroy == nil {
		return nil, false
	}
	return o.Destroy, true
}

// HasDestroy returns a boolean if a field has been set.
func (o *UserCapabilityWithSelfOnly) HasDestroy() bool {
	if o != nil && o.Destroy != nil {
		return true
	}

	return false
}

// SetDestroy gets a reference to the given bool and assigns it to the Destroy field.
func (o *UserCapabilityWithSelfOnly) SetDestroy(v bool) {
	o.Destroy = &v
}

// GetRead returns the Read field value if set, zero value otherwise.
func (o *UserCapabilityWithSelfOnly) GetRead() bool {
	if o == nil || o.Read == nil {
		var ret bool
		return ret
	}
	return *o.Read
}

// GetReadOk returns a tuple with the Read field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCapabilityWithSelfOnly) GetReadOk() (*bool, bool) {
	if o == nil || o.Read == nil {
		return nil, false
	}
	return o.Read, true
}

// HasRead returns a boolean if a field has been set.
func (o *UserCapabilityWithSelfOnly) HasRead() bool {
	if o != nil && o.Read != nil {
		return true
	}

	return false
}

// SetRead gets a reference to the given bool and assigns it to the Read field.
func (o *UserCapabilityWithSelfOnly) SetRead(v bool) {
	o.Read = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *UserCapabilityWithSelfOnly) GetUpdate() bool {
	if o == nil || o.Update == nil {
		var ret bool
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCapabilityWithSelfOnly) GetUpdateOk() (*bool, bool) {
	if o == nil || o.Update == nil {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *UserCapabilityWithSelfOnly) HasUpdate() bool {
	if o != nil && o.Update != nil {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given bool and assigns it to the Update field.
func (o *UserCapabilityWithSelfOnly) SetUpdate(v bool) {
	o.Update = &v
}

func (o UserCapabilityWithSelfOnly) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowedTarget != nil {
		toSerialize["allowed_target"] = o.AllowedTarget
	}
	if o.Create != nil {
		toSerialize["create"] = o.Create
	}
	if o.Destroy != nil {
		toSerialize["destroy"] = o.Destroy
	}
	if o.Read != nil {
		toSerialize["read"] = o.Read
	}
	if o.Update != nil {
		toSerialize["update"] = o.Update
	}
	return json.Marshal(toSerialize)
}

type NullableUserCapabilityWithSelfOnly struct {
	value *UserCapabilityWithSelfOnly
	isSet bool
}

func (v NullableUserCapabilityWithSelfOnly) Get() *UserCapabilityWithSelfOnly {
	return v.value
}

func (v *NullableUserCapabilityWithSelfOnly) Set(val *UserCapabilityWithSelfOnly) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCapabilityWithSelfOnly) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCapabilityWithSelfOnly) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCapabilityWithSelfOnly(val *UserCapabilityWithSelfOnly) *NullableUserCapabilityWithSelfOnly {
	return &NullableUserCapabilityWithSelfOnly{value: val, isSet: true}
}

func (v NullableUserCapabilityWithSelfOnly) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCapabilityWithSelfOnly) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


