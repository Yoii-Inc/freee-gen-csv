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

// PaymentRequestsIndexResponsePaymentRequestsInner struct for PaymentRequestsIndexResponsePaymentRequestsInner
type PaymentRequestsIndexResponsePaymentRequestsInner struct {
	// 申請者のユーザーID
	ApplicantId int64 `json:"applicant_id"`
	// 申請日 (yyyy-mm-dd)
	ApplicationDate string `json:"application_date"`
	// 申請No.
	ApplicationNumber string `json:"application_number"`
	// 承認者（配列）   承認ステップのresource_typeがunspecified (指定なし)の場合はapproversはレスポンスに含まれません。   しかし、resource_typeがunspecifiedの承認ステップにおいて誰かが承認・却下・差し戻しのいずれかのアクションを取った後は、   approversはレスポンスに含まれるようになります。   その場合approversにはアクションを行ったステップのIDとアクションを行ったユーザーのIDが含まれます。
	Approvers []ApprovalRequestResponseApprovalRequestApproversInner `json:"approvers"`
	// 事業所ID
	CompanyId int64 `json:"company_id"`
	// 現在のround。差し戻し等により申請がstepの最初からやり直しになるとroundの値が増えます。
	CurrentRound int64 `json:"current_round"`
	// 現在承認ステップID
	CurrentStepId NullableInt64 `json:"current_step_id"`
	// 取引ID (申請ステータス:statusがapprovedで、取引が存在する時のみdeal_idが表示されます)
	DealId NullableInt64 `json:"deal_id,omitempty"`
	// 取引ステータス (申請ステータス:statusがapprovedで、取引が存在する時のみdeal_statusが表示されます settled:支払済み, unsettled:支払待ち)
	DealStatus NullableString `json:"deal_status,omitempty"`
	// 請求書番号
	DocumentCode string `json:"document_code"`
	// 支払依頼ID
	Id int64 `json:"id"`
	// 発生日 (yyyy-mm-dd)
	IssueDate string `json:"issue_date"`
	// 取引先コード
	PartnerCode NullableString `json:"partner_code"`
	// 取引先ID
	PartnerId NullableInt64 `json:"partner_id"`
	// 取引先名
	PartnerName NullableString `json:"partner_name"`
	// 支払期限 (yyyy-mm-dd)
	PaymentDate NullableString `json:"payment_date"`
	// 支払方法(none: 指定なし, domestic_bank_transfer: 国内振込, abroad_bank_transfer: 国外振込, account_transfer: 口座振替, credit_card: クレジットカード)
	PaymentMethod string `json:"payment_method"`
	// 申請ステータス(draft:下書き, in_progress:申請中, approved:承認済, rejected:却下, feedback:差戻し)
	Status string `json:"status"`
	// 申請タイトル
	Title string `json:"title"`
	// 合計金額
	TotalAmount int64 `json:"total_amount"`
}

// NewPaymentRequestsIndexResponsePaymentRequestsInner instantiates a new PaymentRequestsIndexResponsePaymentRequestsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentRequestsIndexResponsePaymentRequestsInner(applicantId int64, applicationDate string, applicationNumber string, approvers []ApprovalRequestResponseApprovalRequestApproversInner, companyId int64, currentRound int64, currentStepId NullableInt64, documentCode string, id int64, issueDate string, partnerCode NullableString, partnerId NullableInt64, partnerName NullableString, paymentDate NullableString, paymentMethod string, status string, title string, totalAmount int64) *PaymentRequestsIndexResponsePaymentRequestsInner {
	this := PaymentRequestsIndexResponsePaymentRequestsInner{}
	this.ApplicantId = applicantId
	this.ApplicationDate = applicationDate
	this.ApplicationNumber = applicationNumber
	this.Approvers = approvers
	this.CompanyId = companyId
	this.CurrentRound = currentRound
	this.CurrentStepId = currentStepId
	this.DocumentCode = documentCode
	this.Id = id
	this.IssueDate = issueDate
	this.PartnerCode = partnerCode
	this.PartnerId = partnerId
	this.PartnerName = partnerName
	this.PaymentDate = paymentDate
	this.PaymentMethod = paymentMethod
	this.Status = status
	this.Title = title
	this.TotalAmount = totalAmount
	return &this
}

// NewPaymentRequestsIndexResponsePaymentRequestsInnerWithDefaults instantiates a new PaymentRequestsIndexResponsePaymentRequestsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentRequestsIndexResponsePaymentRequestsInnerWithDefaults() *PaymentRequestsIndexResponsePaymentRequestsInner {
	this := PaymentRequestsIndexResponsePaymentRequestsInner{}
	return &this
}

// GetApplicantId returns the ApplicantId field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetApplicantId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ApplicantId
}

// GetApplicantIdOk returns a tuple with the ApplicantId field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetApplicantIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicantId, true
}

// SetApplicantId sets field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetApplicantId(v int64) {
	o.ApplicantId = v
}

// GetApplicationDate returns the ApplicationDate field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetApplicationDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicationDate
}

// GetApplicationDateOk returns a tuple with the ApplicationDate field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetApplicationDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationDate, true
}

// SetApplicationDate sets field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetApplicationDate(v string) {
	o.ApplicationDate = v
}

// GetApplicationNumber returns the ApplicationNumber field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetApplicationNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicationNumber
}

// GetApplicationNumberOk returns a tuple with the ApplicationNumber field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetApplicationNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationNumber, true
}

// SetApplicationNumber sets field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetApplicationNumber(v string) {
	o.ApplicationNumber = v
}

// GetApprovers returns the Approvers field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetApprovers() []ApprovalRequestResponseApprovalRequestApproversInner {
	if o == nil {
		var ret []ApprovalRequestResponseApprovalRequestApproversInner
		return ret
	}

	return o.Approvers
}

// GetApproversOk returns a tuple with the Approvers field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetApproversOk() ([]ApprovalRequestResponseApprovalRequestApproversInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Approvers, true
}

// SetApprovers sets field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetApprovers(v []ApprovalRequestResponseApprovalRequestApproversInner) {
	o.Approvers = v
}

// GetCompanyId returns the CompanyId field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetCompanyId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetCompanyIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyId, true
}

// SetCompanyId sets field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetCompanyId(v int64) {
	o.CompanyId = v
}

// GetCurrentRound returns the CurrentRound field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetCurrentRound() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CurrentRound
}

// GetCurrentRoundOk returns a tuple with the CurrentRound field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetCurrentRoundOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentRound, true
}

// SetCurrentRound sets field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetCurrentRound(v int64) {
	o.CurrentRound = v
}

// GetCurrentStepId returns the CurrentStepId field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetCurrentStepId() int64 {
	if o == nil || o.CurrentStepId.Get() == nil {
		var ret int64
		return ret
	}

	return *o.CurrentStepId.Get()
}

// GetCurrentStepIdOk returns a tuple with the CurrentStepId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetCurrentStepIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentStepId.Get(), o.CurrentStepId.IsSet()
}

// SetCurrentStepId sets field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetCurrentStepId(v int64) {
	o.CurrentStepId.Set(&v)
}

// GetDealId returns the DealId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetDealId() int64 {
	if o == nil || o.DealId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DealId.Get()
}

// GetDealIdOk returns a tuple with the DealId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetDealIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DealId.Get(), o.DealId.IsSet()
}

// HasDealId returns a boolean if a field has been set.
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) HasDealId() bool {
	if o != nil && o.DealId.IsSet() {
		return true
	}

	return false
}

// SetDealId gets a reference to the given NullableInt64 and assigns it to the DealId field.
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetDealId(v int64) {
	o.DealId.Set(&v)
}
// SetDealIdNil sets the value for DealId to be an explicit nil
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetDealIdNil() {
	o.DealId.Set(nil)
}

// UnsetDealId ensures that no value is present for DealId, not even an explicit nil
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) UnsetDealId() {
	o.DealId.Unset()
}

// GetDealStatus returns the DealStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetDealStatus() string {
	if o == nil || o.DealStatus.Get() == nil {
		var ret string
		return ret
	}
	return *o.DealStatus.Get()
}

// GetDealStatusOk returns a tuple with the DealStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetDealStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DealStatus.Get(), o.DealStatus.IsSet()
}

// HasDealStatus returns a boolean if a field has been set.
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) HasDealStatus() bool {
	if o != nil && o.DealStatus.IsSet() {
		return true
	}

	return false
}

// SetDealStatus gets a reference to the given NullableString and assigns it to the DealStatus field.
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetDealStatus(v string) {
	o.DealStatus.Set(&v)
}
// SetDealStatusNil sets the value for DealStatus to be an explicit nil
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetDealStatusNil() {
	o.DealStatus.Set(nil)
}

// UnsetDealStatus ensures that no value is present for DealStatus, not even an explicit nil
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) UnsetDealStatus() {
	o.DealStatus.Unset()
}

// GetDocumentCode returns the DocumentCode field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetDocumentCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentCode
}

// GetDocumentCodeOk returns a tuple with the DocumentCode field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetDocumentCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentCode, true
}

// SetDocumentCode sets field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetDocumentCode(v string) {
	o.DocumentCode = v
}

// GetId returns the Id field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetId(v int64) {
	o.Id = v
}

// GetIssueDate returns the IssueDate field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetIssueDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssueDate
}

// GetIssueDateOk returns a tuple with the IssueDate field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetIssueDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueDate, true
}

// SetIssueDate sets field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetIssueDate(v string) {
	o.IssueDate = v
}

// GetPartnerCode returns the PartnerCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetPartnerCode() string {
	if o == nil || o.PartnerCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.PartnerCode.Get()
}

// GetPartnerCodeOk returns a tuple with the PartnerCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetPartnerCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PartnerCode.Get(), o.PartnerCode.IsSet()
}

// SetPartnerCode sets field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetPartnerCode(v string) {
	o.PartnerCode.Set(&v)
}

// GetPartnerId returns the PartnerId field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetPartnerId() int64 {
	if o == nil || o.PartnerId.Get() == nil {
		var ret int64
		return ret
	}

	return *o.PartnerId.Get()
}

// GetPartnerIdOk returns a tuple with the PartnerId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetPartnerIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PartnerId.Get(), o.PartnerId.IsSet()
}

// SetPartnerId sets field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetPartnerId(v int64) {
	o.PartnerId.Set(&v)
}

// GetPartnerName returns the PartnerName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetPartnerName() string {
	if o == nil || o.PartnerName.Get() == nil {
		var ret string
		return ret
	}

	return *o.PartnerName.Get()
}

// GetPartnerNameOk returns a tuple with the PartnerName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetPartnerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PartnerName.Get(), o.PartnerName.IsSet()
}

// SetPartnerName sets field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetPartnerName(v string) {
	o.PartnerName.Set(&v)
}

// GetPaymentDate returns the PaymentDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetPaymentDate() string {
	if o == nil || o.PaymentDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.PaymentDate.Get()
}

// GetPaymentDateOk returns a tuple with the PaymentDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetPaymentDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentDate.Get(), o.PaymentDate.IsSet()
}

// SetPaymentDate sets field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetPaymentDate(v string) {
	o.PaymentDate.Set(&v)
}

// GetPaymentMethod returns the PaymentMethod field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetPaymentMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetPaymentMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentMethod, true
}

// SetPaymentMethod sets field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetPaymentMethod(v string) {
	o.PaymentMethod = v
}

// GetStatus returns the Status field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetStatus(v string) {
	o.Status = v
}

// GetTitle returns the Title field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetTitle(v string) {
	o.Title = v
}

// GetTotalAmount returns the TotalAmount field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetTotalAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetTotalAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalAmount, true
}

// SetTotalAmount sets field value
func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetTotalAmount(v int64) {
	o.TotalAmount = v
}

func (o PaymentRequestsIndexResponsePaymentRequestsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["applicant_id"] = o.ApplicantId
	}
	if true {
		toSerialize["application_date"] = o.ApplicationDate
	}
	if true {
		toSerialize["application_number"] = o.ApplicationNumber
	}
	if true {
		toSerialize["approvers"] = o.Approvers
	}
	if true {
		toSerialize["company_id"] = o.CompanyId
	}
	if true {
		toSerialize["current_round"] = o.CurrentRound
	}
	if true {
		toSerialize["current_step_id"] = o.CurrentStepId.Get()
	}
	if o.DealId.IsSet() {
		toSerialize["deal_id"] = o.DealId.Get()
	}
	if o.DealStatus.IsSet() {
		toSerialize["deal_status"] = o.DealStatus.Get()
	}
	if true {
		toSerialize["document_code"] = o.DocumentCode
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["issue_date"] = o.IssueDate
	}
	if true {
		toSerialize["partner_code"] = o.PartnerCode.Get()
	}
	if true {
		toSerialize["partner_id"] = o.PartnerId.Get()
	}
	if true {
		toSerialize["partner_name"] = o.PartnerName.Get()
	}
	if true {
		toSerialize["payment_date"] = o.PaymentDate.Get()
	}
	if true {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["total_amount"] = o.TotalAmount
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentRequestsIndexResponsePaymentRequestsInner struct {
	value *PaymentRequestsIndexResponsePaymentRequestsInner
	isSet bool
}

func (v NullablePaymentRequestsIndexResponsePaymentRequestsInner) Get() *PaymentRequestsIndexResponsePaymentRequestsInner {
	return v.value
}

func (v *NullablePaymentRequestsIndexResponsePaymentRequestsInner) Set(val *PaymentRequestsIndexResponsePaymentRequestsInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentRequestsIndexResponsePaymentRequestsInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentRequestsIndexResponsePaymentRequestsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentRequestsIndexResponsePaymentRequestsInner(val *PaymentRequestsIndexResponsePaymentRequestsInner) *NullablePaymentRequestsIndexResponsePaymentRequestsInner {
	return &NullablePaymentRequestsIndexResponsePaymentRequestsInner{value: val, isSet: true}
}

func (v NullablePaymentRequestsIndexResponsePaymentRequestsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentRequestsIndexResponsePaymentRequestsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

