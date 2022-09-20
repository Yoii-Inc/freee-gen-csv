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

// ApprovalRequestCreateParams struct for ApprovalRequestCreateParams
type ApprovalRequestCreateParams struct {
	// 申請日 (yyyy-mm-dd)<br> 指定しない場合は当日の日付が登録されます。 
	ApplicationDate *string `json:"application_date,omitempty"`
	// 申請経路ID
	ApprovalFlowRouteId int64 `json:"approval_flow_route_id"`
	// 承認者のユーザーID
	ApproverId *int64 `json:"approver_id,omitempty"`
	// 事業所ID
	CompanyId int64 `json:"company_id"`
	// 各種申請のステータス<br> falseを指定した時は申請中（in_progress）で各種申請を作成します。<br> trueを指定した時は下書き（draft）で各種申請を作成します。 
	Draft bool `json:"draft"`
	// 申請フォームID
	FormId int64 `json:"form_id"`
	// 親申請ID(既存各種申請IDのみ指定可能です。)
	ParentId *int64 `json:"parent_id,omitempty"`
	RequestItems []ApprovalRequestCreateParamsRequestItemsInner `json:"request_items"`
}

// NewApprovalRequestCreateParams instantiates a new ApprovalRequestCreateParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApprovalRequestCreateParams(approvalFlowRouteId int64, companyId int64, draft bool, formId int64, requestItems []ApprovalRequestCreateParamsRequestItemsInner) *ApprovalRequestCreateParams {
	this := ApprovalRequestCreateParams{}
	this.ApprovalFlowRouteId = approvalFlowRouteId
	this.CompanyId = companyId
	this.Draft = draft
	this.FormId = formId
	this.RequestItems = requestItems
	return &this
}

// NewApprovalRequestCreateParamsWithDefaults instantiates a new ApprovalRequestCreateParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApprovalRequestCreateParamsWithDefaults() *ApprovalRequestCreateParams {
	this := ApprovalRequestCreateParams{}
	return &this
}

// GetApplicationDate returns the ApplicationDate field value if set, zero value otherwise.
func (o *ApprovalRequestCreateParams) GetApplicationDate() string {
	if o == nil || o.ApplicationDate == nil {
		var ret string
		return ret
	}
	return *o.ApplicationDate
}

// GetApplicationDateOk returns a tuple with the ApplicationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalRequestCreateParams) GetApplicationDateOk() (*string, bool) {
	if o == nil || o.ApplicationDate == nil {
		return nil, false
	}
	return o.ApplicationDate, true
}

// HasApplicationDate returns a boolean if a field has been set.
func (o *ApprovalRequestCreateParams) HasApplicationDate() bool {
	if o != nil && o.ApplicationDate != nil {
		return true
	}

	return false
}

// SetApplicationDate gets a reference to the given string and assigns it to the ApplicationDate field.
func (o *ApprovalRequestCreateParams) SetApplicationDate(v string) {
	o.ApplicationDate = &v
}

// GetApprovalFlowRouteId returns the ApprovalFlowRouteId field value
func (o *ApprovalRequestCreateParams) GetApprovalFlowRouteId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ApprovalFlowRouteId
}

// GetApprovalFlowRouteIdOk returns a tuple with the ApprovalFlowRouteId field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestCreateParams) GetApprovalFlowRouteIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApprovalFlowRouteId, true
}

// SetApprovalFlowRouteId sets field value
func (o *ApprovalRequestCreateParams) SetApprovalFlowRouteId(v int64) {
	o.ApprovalFlowRouteId = v
}

// GetApproverId returns the ApproverId field value if set, zero value otherwise.
func (o *ApprovalRequestCreateParams) GetApproverId() int64 {
	if o == nil || o.ApproverId == nil {
		var ret int64
		return ret
	}
	return *o.ApproverId
}

// GetApproverIdOk returns a tuple with the ApproverId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalRequestCreateParams) GetApproverIdOk() (*int64, bool) {
	if o == nil || o.ApproverId == nil {
		return nil, false
	}
	return o.ApproverId, true
}

// HasApproverId returns a boolean if a field has been set.
func (o *ApprovalRequestCreateParams) HasApproverId() bool {
	if o != nil && o.ApproverId != nil {
		return true
	}

	return false
}

// SetApproverId gets a reference to the given int64 and assigns it to the ApproverId field.
func (o *ApprovalRequestCreateParams) SetApproverId(v int64) {
	o.ApproverId = &v
}

// GetCompanyId returns the CompanyId field value
func (o *ApprovalRequestCreateParams) GetCompanyId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestCreateParams) GetCompanyIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyId, true
}

// SetCompanyId sets field value
func (o *ApprovalRequestCreateParams) SetCompanyId(v int64) {
	o.CompanyId = v
}

// GetDraft returns the Draft field value
func (o *ApprovalRequestCreateParams) GetDraft() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Draft
}

// GetDraftOk returns a tuple with the Draft field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestCreateParams) GetDraftOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Draft, true
}

// SetDraft sets field value
func (o *ApprovalRequestCreateParams) SetDraft(v bool) {
	o.Draft = v
}

// GetFormId returns the FormId field value
func (o *ApprovalRequestCreateParams) GetFormId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FormId
}

// GetFormIdOk returns a tuple with the FormId field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestCreateParams) GetFormIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormId, true
}

// SetFormId sets field value
func (o *ApprovalRequestCreateParams) SetFormId(v int64) {
	o.FormId = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *ApprovalRequestCreateParams) GetParentId() int64 {
	if o == nil || o.ParentId == nil {
		var ret int64
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalRequestCreateParams) GetParentIdOk() (*int64, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *ApprovalRequestCreateParams) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given int64 and assigns it to the ParentId field.
func (o *ApprovalRequestCreateParams) SetParentId(v int64) {
	o.ParentId = &v
}

// GetRequestItems returns the RequestItems field value
func (o *ApprovalRequestCreateParams) GetRequestItems() []ApprovalRequestCreateParamsRequestItemsInner {
	if o == nil {
		var ret []ApprovalRequestCreateParamsRequestItemsInner
		return ret
	}

	return o.RequestItems
}

// GetRequestItemsOk returns a tuple with the RequestItems field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestCreateParams) GetRequestItemsOk() ([]ApprovalRequestCreateParamsRequestItemsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestItems, true
}

// SetRequestItems sets field value
func (o *ApprovalRequestCreateParams) SetRequestItems(v []ApprovalRequestCreateParamsRequestItemsInner) {
	o.RequestItems = v
}

func (o ApprovalRequestCreateParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplicationDate != nil {
		toSerialize["application_date"] = o.ApplicationDate
	}
	if true {
		toSerialize["approval_flow_route_id"] = o.ApprovalFlowRouteId
	}
	if o.ApproverId != nil {
		toSerialize["approver_id"] = o.ApproverId
	}
	if true {
		toSerialize["company_id"] = o.CompanyId
	}
	if true {
		toSerialize["draft"] = o.Draft
	}
	if true {
		toSerialize["form_id"] = o.FormId
	}
	if o.ParentId != nil {
		toSerialize["parent_id"] = o.ParentId
	}
	if true {
		toSerialize["request_items"] = o.RequestItems
	}
	return json.Marshal(toSerialize)
}

type NullableApprovalRequestCreateParams struct {
	value *ApprovalRequestCreateParams
	isSet bool
}

func (v NullableApprovalRequestCreateParams) Get() *ApprovalRequestCreateParams {
	return v.value
}

func (v *NullableApprovalRequestCreateParams) Set(val *ApprovalRequestCreateParams) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalRequestCreateParams) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalRequestCreateParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalRequestCreateParams(val *ApprovalRequestCreateParams) *NullableApprovalRequestCreateParams {
	return &NullableApprovalRequestCreateParams{value: val, isSet: true}
}

func (v NullableApprovalRequestCreateParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalRequestCreateParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


