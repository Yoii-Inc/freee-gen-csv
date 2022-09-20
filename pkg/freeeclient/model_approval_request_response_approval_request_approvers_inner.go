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

// ApprovalRequestResponseApprovalRequestApproversInner struct for ApprovalRequestResponseApprovalRequestApproversInner
type ApprovalRequestResponseApprovalRequestApproversInner struct {
	// 代理承認済みかどうか
	IsForceAction bool `json:"is_force_action"`
	// 承認ステップの承認方法 * ` predefined_user` - メンバー指定 (1人), * ` selected_user` - 申請時にメンバー指定 * ` unspecified` - 指定なし * ` and_resource` - メンバー指定 (複数、全員の承認), * ` or_resource` - メンバー指定 (複数、1人の承認) * ` and_position` - 役職指定 (複数、全員の承認) * ` or_position` - 役職指定 (複数、1人の承認)
	ResourceType string `json:"resource_type"`
	// 承認者の承認状態 * `initial` - 初期状態 * `approved` - 承認済 * `rejected` - 却下 * `feedback` - 差戻し
	Status string `json:"status"`
	// 承認ステップID
	StepId int32 `json:"step_id"`
	// 承認者のユーザーID 下記の場合はnullになります。 <ul>   <li>resource_type:selected_userの場合で承認者未指定時</li>   <li>resource_type:or_positionで前stepで部門未指定の場合</li> </ul>
	UserId NullableInt32 `json:"user_id"`
}

// NewApprovalRequestResponseApprovalRequestApproversInner instantiates a new ApprovalRequestResponseApprovalRequestApproversInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApprovalRequestResponseApprovalRequestApproversInner(isForceAction bool, resourceType string, status string, stepId int32, userId NullableInt32) *ApprovalRequestResponseApprovalRequestApproversInner {
	this := ApprovalRequestResponseApprovalRequestApproversInner{}
	this.IsForceAction = isForceAction
	this.ResourceType = resourceType
	this.Status = status
	this.StepId = stepId
	this.UserId = userId
	return &this
}

// NewApprovalRequestResponseApprovalRequestApproversInnerWithDefaults instantiates a new ApprovalRequestResponseApprovalRequestApproversInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApprovalRequestResponseApprovalRequestApproversInnerWithDefaults() *ApprovalRequestResponseApprovalRequestApproversInner {
	this := ApprovalRequestResponseApprovalRequestApproversInner{}
	return &this
}

// GetIsForceAction returns the IsForceAction field value
func (o *ApprovalRequestResponseApprovalRequestApproversInner) GetIsForceAction() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsForceAction
}

// GetIsForceActionOk returns a tuple with the IsForceAction field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestResponseApprovalRequestApproversInner) GetIsForceActionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsForceAction, true
}

// SetIsForceAction sets field value
func (o *ApprovalRequestResponseApprovalRequestApproversInner) SetIsForceAction(v bool) {
	o.IsForceAction = v
}

// GetResourceType returns the ResourceType field value
func (o *ApprovalRequestResponseApprovalRequestApproversInner) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestResponseApprovalRequestApproversInner) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *ApprovalRequestResponseApprovalRequestApproversInner) SetResourceType(v string) {
	o.ResourceType = v
}

// GetStatus returns the Status field value
func (o *ApprovalRequestResponseApprovalRequestApproversInner) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestResponseApprovalRequestApproversInner) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ApprovalRequestResponseApprovalRequestApproversInner) SetStatus(v string) {
	o.Status = v
}

// GetStepId returns the StepId field value
func (o *ApprovalRequestResponseApprovalRequestApproversInner) GetStepId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StepId
}

// GetStepIdOk returns a tuple with the StepId field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestResponseApprovalRequestApproversInner) GetStepIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StepId, true
}

// SetStepId sets field value
func (o *ApprovalRequestResponseApprovalRequestApproversInner) SetStepId(v int32) {
	o.StepId = v
}

// GetUserId returns the UserId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ApprovalRequestResponseApprovalRequestApproversInner) GetUserId() int32 {
	if o == nil || o.UserId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApprovalRequestResponseApprovalRequestApproversInner) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// SetUserId sets field value
func (o *ApprovalRequestResponseApprovalRequestApproversInner) SetUserId(v int32) {
	o.UserId.Set(&v)
}

func (o ApprovalRequestResponseApprovalRequestApproversInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["is_force_action"] = o.IsForceAction
	}
	if true {
		toSerialize["resource_type"] = o.ResourceType
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["step_id"] = o.StepId
	}
	if true {
		toSerialize["user_id"] = o.UserId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableApprovalRequestResponseApprovalRequestApproversInner struct {
	value *ApprovalRequestResponseApprovalRequestApproversInner
	isSet bool
}

func (v NullableApprovalRequestResponseApprovalRequestApproversInner) Get() *ApprovalRequestResponseApprovalRequestApproversInner {
	return v.value
}

func (v *NullableApprovalRequestResponseApprovalRequestApproversInner) Set(val *ApprovalRequestResponseApprovalRequestApproversInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalRequestResponseApprovalRequestApproversInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalRequestResponseApprovalRequestApproversInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalRequestResponseApprovalRequestApproversInner(val *ApprovalRequestResponseApprovalRequestApproversInner) *NullableApprovalRequestResponseApprovalRequestApproversInner {
	return &NullableApprovalRequestResponseApprovalRequestApproversInner{value: val, isSet: true}
}

func (v NullableApprovalRequestResponseApprovalRequestApproversInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalRequestResponseApprovalRequestApproversInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


