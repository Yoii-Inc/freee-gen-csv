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

// ApprovalRequestFormIndexResponseApprovalRequestFormsInner struct for ApprovalRequestFormIndexResponseApprovalRequestFormsInner
type ApprovalRequestFormIndexResponseApprovalRequestFormsInner struct {
	// 事業所ID
	CompanyId int64 `json:"company_id"`
	// 作成日時
	CreatedDate string `json:"created_date"`
	// 申請フォームの説明
	Description string `json:"description"`
	// 表示順（申請者が選択する申請フォームの表示順を設定できます。小さい数ほど上位に表示されます。（0を除く整数のみ。マイナス不可）未入力の場合、表示順が後ろになります。同じ数字が入力された場合、登録順で表示されます。）
	FormOrder NullableInt64 `json:"form_order"`
	// 申請フォームID
	Id int64 `json:"id"`
	// 申請フォームの名前
	Name string `json:"name"`
	// 適用された経路数
	RouteSettingCount int64 `json:"route_setting_count"`
	// ステータス(draft: 申請で使用しない、active: 申請で使用する、deleted: 削除済み)
	Status string `json:"status"`
}

// NewApprovalRequestFormIndexResponseApprovalRequestFormsInner instantiates a new ApprovalRequestFormIndexResponseApprovalRequestFormsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApprovalRequestFormIndexResponseApprovalRequestFormsInner(companyId int64, createdDate string, description string, formOrder NullableInt64, id int64, name string, routeSettingCount int64, status string) *ApprovalRequestFormIndexResponseApprovalRequestFormsInner {
	this := ApprovalRequestFormIndexResponseApprovalRequestFormsInner{}
	this.CompanyId = companyId
	this.CreatedDate = createdDate
	this.Description = description
	this.FormOrder = formOrder
	this.Id = id
	this.Name = name
	this.RouteSettingCount = routeSettingCount
	this.Status = status
	return &this
}

// NewApprovalRequestFormIndexResponseApprovalRequestFormsInnerWithDefaults instantiates a new ApprovalRequestFormIndexResponseApprovalRequestFormsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApprovalRequestFormIndexResponseApprovalRequestFormsInnerWithDefaults() *ApprovalRequestFormIndexResponseApprovalRequestFormsInner {
	this := ApprovalRequestFormIndexResponseApprovalRequestFormsInner{}
	return &this
}

// GetCompanyId returns the CompanyId field value
func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetCompanyId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetCompanyIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyId, true
}

// SetCompanyId sets field value
func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) SetCompanyId(v int64) {
	o.CompanyId = v
}

// GetCreatedDate returns the CreatedDate field value
func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetCreatedDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetCreatedDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedDate, true
}

// SetCreatedDate sets field value
func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) SetCreatedDate(v string) {
	o.CreatedDate = v
}

// GetDescription returns the Description field value
func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) SetDescription(v string) {
	o.Description = v
}

// GetFormOrder returns the FormOrder field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetFormOrder() int64 {
	if o == nil || o.FormOrder.Get() == nil {
		var ret int64
		return ret
	}

	return *o.FormOrder.Get()
}

// GetFormOrderOk returns a tuple with the FormOrder field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetFormOrderOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FormOrder.Get(), o.FormOrder.IsSet()
}

// SetFormOrder sets field value
func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) SetFormOrder(v int64) {
	o.FormOrder.Set(&v)
}

// GetId returns the Id field value
func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) SetName(v string) {
	o.Name = v
}

// GetRouteSettingCount returns the RouteSettingCount field value
func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetRouteSettingCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RouteSettingCount
}

// GetRouteSettingCountOk returns a tuple with the RouteSettingCount field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetRouteSettingCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RouteSettingCount, true
}

// SetRouteSettingCount sets field value
func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) SetRouteSettingCount(v int64) {
	o.RouteSettingCount = v
}

// GetStatus returns the Status field value
func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) SetStatus(v string) {
	o.Status = v
}

func (o ApprovalRequestFormIndexResponseApprovalRequestFormsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["company_id"] = o.CompanyId
	}
	if true {
		toSerialize["created_date"] = o.CreatedDate
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["form_order"] = o.FormOrder.Get()
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["route_setting_count"] = o.RouteSettingCount
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableApprovalRequestFormIndexResponseApprovalRequestFormsInner struct {
	value *ApprovalRequestFormIndexResponseApprovalRequestFormsInner
	isSet bool
}

func (v NullableApprovalRequestFormIndexResponseApprovalRequestFormsInner) Get() *ApprovalRequestFormIndexResponseApprovalRequestFormsInner {
	return v.value
}

func (v *NullableApprovalRequestFormIndexResponseApprovalRequestFormsInner) Set(val *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalRequestFormIndexResponseApprovalRequestFormsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalRequestFormIndexResponseApprovalRequestFormsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalRequestFormIndexResponseApprovalRequestFormsInner(val *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) *NullableApprovalRequestFormIndexResponseApprovalRequestFormsInner {
	return &NullableApprovalRequestFormIndexResponseApprovalRequestFormsInner{value: val, isSet: true}
}

func (v NullableApprovalRequestFormIndexResponseApprovalRequestFormsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalRequestFormIndexResponseApprovalRequestFormsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


