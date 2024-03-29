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

// CompanyResponseCompanySectionsInner struct for CompanyResponseCompanySectionsInner
type CompanyResponseCompanySectionsInner struct {
	// 部門ID
	Id int64 `json:"id"`
	// 部門名 (30文字以内)
	Name string `json:"name"`
	// ショートカット１ (20文字以内)
	Shortcut1 NullableString `json:"shortcut1,omitempty"`
	// ショートカット２ (20文字以内)
	Shortcut2 NullableString `json:"shortcut2,omitempty"`
}

// NewCompanyResponseCompanySectionsInner instantiates a new CompanyResponseCompanySectionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyResponseCompanySectionsInner(id int64, name string) *CompanyResponseCompanySectionsInner {
	this := CompanyResponseCompanySectionsInner{}
	this.Id = id
	this.Name = name
	return &this
}

// NewCompanyResponseCompanySectionsInnerWithDefaults instantiates a new CompanyResponseCompanySectionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyResponseCompanySectionsInnerWithDefaults() *CompanyResponseCompanySectionsInner {
	this := CompanyResponseCompanySectionsInner{}
	return &this
}

// GetId returns the Id field value
func (o *CompanyResponseCompanySectionsInner) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CompanyResponseCompanySectionsInner) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CompanyResponseCompanySectionsInner) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *CompanyResponseCompanySectionsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CompanyResponseCompanySectionsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CompanyResponseCompanySectionsInner) SetName(v string) {
	o.Name = v
}

// GetShortcut1 returns the Shortcut1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyResponseCompanySectionsInner) GetShortcut1() string {
	if o == nil || o.Shortcut1.Get() == nil {
		var ret string
		return ret
	}
	return *o.Shortcut1.Get()
}

// GetShortcut1Ok returns a tuple with the Shortcut1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyResponseCompanySectionsInner) GetShortcut1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Shortcut1.Get(), o.Shortcut1.IsSet()
}

// HasShortcut1 returns a boolean if a field has been set.
func (o *CompanyResponseCompanySectionsInner) HasShortcut1() bool {
	if o != nil && o.Shortcut1.IsSet() {
		return true
	}

	return false
}

// SetShortcut1 gets a reference to the given NullableString and assigns it to the Shortcut1 field.
func (o *CompanyResponseCompanySectionsInner) SetShortcut1(v string) {
	o.Shortcut1.Set(&v)
}
// SetShortcut1Nil sets the value for Shortcut1 to be an explicit nil
func (o *CompanyResponseCompanySectionsInner) SetShortcut1Nil() {
	o.Shortcut1.Set(nil)
}

// UnsetShortcut1 ensures that no value is present for Shortcut1, not even an explicit nil
func (o *CompanyResponseCompanySectionsInner) UnsetShortcut1() {
	o.Shortcut1.Unset()
}

// GetShortcut2 returns the Shortcut2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyResponseCompanySectionsInner) GetShortcut2() string {
	if o == nil || o.Shortcut2.Get() == nil {
		var ret string
		return ret
	}
	return *o.Shortcut2.Get()
}

// GetShortcut2Ok returns a tuple with the Shortcut2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyResponseCompanySectionsInner) GetShortcut2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Shortcut2.Get(), o.Shortcut2.IsSet()
}

// HasShortcut2 returns a boolean if a field has been set.
func (o *CompanyResponseCompanySectionsInner) HasShortcut2() bool {
	if o != nil && o.Shortcut2.IsSet() {
		return true
	}

	return false
}

// SetShortcut2 gets a reference to the given NullableString and assigns it to the Shortcut2 field.
func (o *CompanyResponseCompanySectionsInner) SetShortcut2(v string) {
	o.Shortcut2.Set(&v)
}
// SetShortcut2Nil sets the value for Shortcut2 to be an explicit nil
func (o *CompanyResponseCompanySectionsInner) SetShortcut2Nil() {
	o.Shortcut2.Set(nil)
}

// UnsetShortcut2 ensures that no value is present for Shortcut2, not even an explicit nil
func (o *CompanyResponseCompanySectionsInner) UnsetShortcut2() {
	o.Shortcut2.Unset()
}

func (o CompanyResponseCompanySectionsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Shortcut1.IsSet() {
		toSerialize["shortcut1"] = o.Shortcut1.Get()
	}
	if o.Shortcut2.IsSet() {
		toSerialize["shortcut2"] = o.Shortcut2.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCompanyResponseCompanySectionsInner struct {
	value *CompanyResponseCompanySectionsInner
	isSet bool
}

func (v NullableCompanyResponseCompanySectionsInner) Get() *CompanyResponseCompanySectionsInner {
	return v.value
}

func (v *NullableCompanyResponseCompanySectionsInner) Set(val *CompanyResponseCompanySectionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyResponseCompanySectionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyResponseCompanySectionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyResponseCompanySectionsInner(val *CompanyResponseCompanySectionsInner) *NullableCompanyResponseCompanySectionsInner {
	return &NullableCompanyResponseCompanySectionsInner{value: val, isSet: true}
}

func (v NullableCompanyResponseCompanySectionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyResponseCompanySectionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


