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

// ExpenseApplicationResponseExpenseApplication struct for ExpenseApplicationResponseExpenseApplication
type ExpenseApplicationResponseExpenseApplication struct {
	// 申請者のユーザーID
	ApplicantId int64 `json:"applicant_id"`
	// 申請No.
	ApplicationNumber string `json:"application_number"`
	// 経費申請の承認履歴（配列）
	ApprovalFlowLogs []ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner `json:"approval_flow_logs"`
	// 申請経路ID
	ApprovalFlowRouteId int64 `json:"approval_flow_route_id"`
	// 承認者（配列）   承認ステップのresource_typeがunspecified (指定なし)の場合はapproversはレスポンスに含まれません。   しかし、resource_typeがunspecifiedの承認ステップにおいて誰かが承認・却下・差し戻しのいずれかのアクションを取った後は、   approversはレスポンスに含まれるようになります。   その場合approversにはアクションを行ったステップのIDとアクションを行ったユーザーのIDが含まれます。
	Approvers []ApprovalRequestResponseApprovalRequestApproversInner `json:"approvers"`
	// 経費申請のコメント一覧（配列）
	Comments []ApprovalRequestResponseApprovalRequestCommentsInner `json:"comments"`
	// 事業所ID
	CompanyId int64 `json:"company_id"`
	// 現在のround。差し戻し等により申請がstepの最初からやり直しになるとroundの値が増えます。
	CurrentRound int64 `json:"current_round"`
	// 現在承認ステップID
	CurrentStepId NullableInt64 `json:"current_step_id"`
	// 取引ID (申請ステータス:statusがapprovedで、取引が存在する時のみdeal_idが表示されます)
	DealId NullableInt64 `json:"deal_id"`
	// 取引ステータス (申請ステータス:statusがapprovedで、取引が存在する時のみdeal_statusが表示されます settled:精算済み, unsettled:清算待ち)
	DealStatus NullableString `json:"deal_status"`
	// 備考
	Description NullableString `json:"description,omitempty"`
	// 経費申請の項目行一覧（配列）
	ExpenseApplicationLines []ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner `json:"expense_application_lines"`
	// 経費申請ID
	Id int64 `json:"id"`
	// 申請日 (yyyy-mm-dd)
	IssueDate string `json:"issue_date"`
	// 部門ID
	SectionId NullableInt64 `json:"section_id,omitempty"`
	// セグメント１ID。セグメント１が使用可能なプランの時のみレスポンスに含まれます。
	Segment1TagId NullableInt64 `json:"segment_1_tag_id,omitempty"`
	// セグメント２ID。セグメント２が使用可能なプランの時のみレスポンスに含まれます。
	Segment2TagId NullableInt64 `json:"segment_2_tag_id,omitempty"`
	// セグメント３ID。セグメント３が使用可能なプランの時のみレスポンスに含まれます。
	Segment3TagId NullableInt64 `json:"segment_3_tag_id,omitempty"`
	// 申請ステータス(draft:下書き, in_progress:申請中, approved:承認済, rejected:却下, feedback:差戻し)
	Status string `json:"status"`
	// メモタグID
	TagIds []int64 `json:"tag_ids,omitempty"`
	// 申請タイトル
	Title string `json:"title"`
	// 合計金額
	TotalAmount *int64 `json:"total_amount,omitempty"`
}

// NewExpenseApplicationResponseExpenseApplication instantiates a new ExpenseApplicationResponseExpenseApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpenseApplicationResponseExpenseApplication(applicantId int64, applicationNumber string, approvalFlowLogs []ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner, approvalFlowRouteId int64, approvers []ApprovalRequestResponseApprovalRequestApproversInner, comments []ApprovalRequestResponseApprovalRequestCommentsInner, companyId int64, currentRound int64, currentStepId NullableInt64, dealId NullableInt64, dealStatus NullableString, expenseApplicationLines []ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner, id int64, issueDate string, status string, title string) *ExpenseApplicationResponseExpenseApplication {
	this := ExpenseApplicationResponseExpenseApplication{}
	this.ApplicantId = applicantId
	this.ApplicationNumber = applicationNumber
	this.ApprovalFlowLogs = approvalFlowLogs
	this.ApprovalFlowRouteId = approvalFlowRouteId
	this.Approvers = approvers
	this.Comments = comments
	this.CompanyId = companyId
	this.CurrentRound = currentRound
	this.CurrentStepId = currentStepId
	this.DealId = dealId
	this.DealStatus = dealStatus
	this.ExpenseApplicationLines = expenseApplicationLines
	this.Id = id
	this.IssueDate = issueDate
	this.Status = status
	this.Title = title
	return &this
}

// NewExpenseApplicationResponseExpenseApplicationWithDefaults instantiates a new ExpenseApplicationResponseExpenseApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpenseApplicationResponseExpenseApplicationWithDefaults() *ExpenseApplicationResponseExpenseApplication {
	this := ExpenseApplicationResponseExpenseApplication{}
	return &this
}

// GetApplicantId returns the ApplicantId field value
func (o *ExpenseApplicationResponseExpenseApplication) GetApplicantId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ApplicantId
}

// GetApplicantIdOk returns a tuple with the ApplicantId field value
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationResponseExpenseApplication) GetApplicantIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicantId, true
}

// SetApplicantId sets field value
func (o *ExpenseApplicationResponseExpenseApplication) SetApplicantId(v int64) {
	o.ApplicantId = v
}

// GetApplicationNumber returns the ApplicationNumber field value
func (o *ExpenseApplicationResponseExpenseApplication) GetApplicationNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicationNumber
}

// GetApplicationNumberOk returns a tuple with the ApplicationNumber field value
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationResponseExpenseApplication) GetApplicationNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationNumber, true
}

// SetApplicationNumber sets field value
func (o *ExpenseApplicationResponseExpenseApplication) SetApplicationNumber(v string) {
	o.ApplicationNumber = v
}

// GetApprovalFlowLogs returns the ApprovalFlowLogs field value
func (o *ExpenseApplicationResponseExpenseApplication) GetApprovalFlowLogs() []ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner {
	if o == nil {
		var ret []ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner
		return ret
	}

	return o.ApprovalFlowLogs
}

// GetApprovalFlowLogsOk returns a tuple with the ApprovalFlowLogs field value
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationResponseExpenseApplication) GetApprovalFlowLogsOk() ([]ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApprovalFlowLogs, true
}

// SetApprovalFlowLogs sets field value
func (o *ExpenseApplicationResponseExpenseApplication) SetApprovalFlowLogs(v []ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner) {
	o.ApprovalFlowLogs = v
}

// GetApprovalFlowRouteId returns the ApprovalFlowRouteId field value
func (o *ExpenseApplicationResponseExpenseApplication) GetApprovalFlowRouteId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ApprovalFlowRouteId
}

// GetApprovalFlowRouteIdOk returns a tuple with the ApprovalFlowRouteId field value
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationResponseExpenseApplication) GetApprovalFlowRouteIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApprovalFlowRouteId, true
}

// SetApprovalFlowRouteId sets field value
func (o *ExpenseApplicationResponseExpenseApplication) SetApprovalFlowRouteId(v int64) {
	o.ApprovalFlowRouteId = v
}

// GetApprovers returns the Approvers field value
func (o *ExpenseApplicationResponseExpenseApplication) GetApprovers() []ApprovalRequestResponseApprovalRequestApproversInner {
	if o == nil {
		var ret []ApprovalRequestResponseApprovalRequestApproversInner
		return ret
	}

	return o.Approvers
}

// GetApproversOk returns a tuple with the Approvers field value
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationResponseExpenseApplication) GetApproversOk() ([]ApprovalRequestResponseApprovalRequestApproversInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Approvers, true
}

// SetApprovers sets field value
func (o *ExpenseApplicationResponseExpenseApplication) SetApprovers(v []ApprovalRequestResponseApprovalRequestApproversInner) {
	o.Approvers = v
}

// GetComments returns the Comments field value
func (o *ExpenseApplicationResponseExpenseApplication) GetComments() []ApprovalRequestResponseApprovalRequestCommentsInner {
	if o == nil {
		var ret []ApprovalRequestResponseApprovalRequestCommentsInner
		return ret
	}

	return o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationResponseExpenseApplication) GetCommentsOk() ([]ApprovalRequestResponseApprovalRequestCommentsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comments, true
}

// SetComments sets field value
func (o *ExpenseApplicationResponseExpenseApplication) SetComments(v []ApprovalRequestResponseApprovalRequestCommentsInner) {
	o.Comments = v
}

// GetCompanyId returns the CompanyId field value
func (o *ExpenseApplicationResponseExpenseApplication) GetCompanyId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationResponseExpenseApplication) GetCompanyIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyId, true
}

// SetCompanyId sets field value
func (o *ExpenseApplicationResponseExpenseApplication) SetCompanyId(v int64) {
	o.CompanyId = v
}

// GetCurrentRound returns the CurrentRound field value
func (o *ExpenseApplicationResponseExpenseApplication) GetCurrentRound() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CurrentRound
}

// GetCurrentRoundOk returns a tuple with the CurrentRound field value
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationResponseExpenseApplication) GetCurrentRoundOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentRound, true
}

// SetCurrentRound sets field value
func (o *ExpenseApplicationResponseExpenseApplication) SetCurrentRound(v int64) {
	o.CurrentRound = v
}

// GetCurrentStepId returns the CurrentStepId field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *ExpenseApplicationResponseExpenseApplication) GetCurrentStepId() int64 {
	if o == nil || o.CurrentStepId.Get() == nil {
		var ret int64
		return ret
	}

	return *o.CurrentStepId.Get()
}

// GetCurrentStepIdOk returns a tuple with the CurrentStepId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpenseApplicationResponseExpenseApplication) GetCurrentStepIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentStepId.Get(), o.CurrentStepId.IsSet()
}

// SetCurrentStepId sets field value
func (o *ExpenseApplicationResponseExpenseApplication) SetCurrentStepId(v int64) {
	o.CurrentStepId.Set(&v)
}

// GetDealId returns the DealId field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *ExpenseApplicationResponseExpenseApplication) GetDealId() int64 {
	if o == nil || o.DealId.Get() == nil {
		var ret int64
		return ret
	}

	return *o.DealId.Get()
}

// GetDealIdOk returns a tuple with the DealId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpenseApplicationResponseExpenseApplication) GetDealIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DealId.Get(), o.DealId.IsSet()
}

// SetDealId sets field value
func (o *ExpenseApplicationResponseExpenseApplication) SetDealId(v int64) {
	o.DealId.Set(&v)
}

// GetDealStatus returns the DealStatus field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ExpenseApplicationResponseExpenseApplication) GetDealStatus() string {
	if o == nil || o.DealStatus.Get() == nil {
		var ret string
		return ret
	}

	return *o.DealStatus.Get()
}

// GetDealStatusOk returns a tuple with the DealStatus field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpenseApplicationResponseExpenseApplication) GetDealStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DealStatus.Get(), o.DealStatus.IsSet()
}

// SetDealStatus sets field value
func (o *ExpenseApplicationResponseExpenseApplication) SetDealStatus(v string) {
	o.DealStatus.Set(&v)
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpenseApplicationResponseExpenseApplication) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpenseApplicationResponseExpenseApplication) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ExpenseApplicationResponseExpenseApplication) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ExpenseApplicationResponseExpenseApplication) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ExpenseApplicationResponseExpenseApplication) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ExpenseApplicationResponseExpenseApplication) UnsetDescription() {
	o.Description.Unset()
}

// GetExpenseApplicationLines returns the ExpenseApplicationLines field value
func (o *ExpenseApplicationResponseExpenseApplication) GetExpenseApplicationLines() []ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner {
	if o == nil {
		var ret []ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner
		return ret
	}

	return o.ExpenseApplicationLines
}

// GetExpenseApplicationLinesOk returns a tuple with the ExpenseApplicationLines field value
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationResponseExpenseApplication) GetExpenseApplicationLinesOk() ([]ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpenseApplicationLines, true
}

// SetExpenseApplicationLines sets field value
func (o *ExpenseApplicationResponseExpenseApplication) SetExpenseApplicationLines(v []ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) {
	o.ExpenseApplicationLines = v
}

// GetId returns the Id field value
func (o *ExpenseApplicationResponseExpenseApplication) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationResponseExpenseApplication) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExpenseApplicationResponseExpenseApplication) SetId(v int64) {
	o.Id = v
}

// GetIssueDate returns the IssueDate field value
func (o *ExpenseApplicationResponseExpenseApplication) GetIssueDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssueDate
}

// GetIssueDateOk returns a tuple with the IssueDate field value
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationResponseExpenseApplication) GetIssueDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueDate, true
}

// SetIssueDate sets field value
func (o *ExpenseApplicationResponseExpenseApplication) SetIssueDate(v string) {
	o.IssueDate = v
}

// GetSectionId returns the SectionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpenseApplicationResponseExpenseApplication) GetSectionId() int64 {
	if o == nil || o.SectionId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SectionId.Get()
}

// GetSectionIdOk returns a tuple with the SectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpenseApplicationResponseExpenseApplication) GetSectionIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SectionId.Get(), o.SectionId.IsSet()
}

// HasSectionId returns a boolean if a field has been set.
func (o *ExpenseApplicationResponseExpenseApplication) HasSectionId() bool {
	if o != nil && o.SectionId.IsSet() {
		return true
	}

	return false
}

// SetSectionId gets a reference to the given NullableInt64 and assigns it to the SectionId field.
func (o *ExpenseApplicationResponseExpenseApplication) SetSectionId(v int64) {
	o.SectionId.Set(&v)
}
// SetSectionIdNil sets the value for SectionId to be an explicit nil
func (o *ExpenseApplicationResponseExpenseApplication) SetSectionIdNil() {
	o.SectionId.Set(nil)
}

// UnsetSectionId ensures that no value is present for SectionId, not even an explicit nil
func (o *ExpenseApplicationResponseExpenseApplication) UnsetSectionId() {
	o.SectionId.Unset()
}

// GetSegment1TagId returns the Segment1TagId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpenseApplicationResponseExpenseApplication) GetSegment1TagId() int64 {
	if o == nil || o.Segment1TagId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Segment1TagId.Get()
}

// GetSegment1TagIdOk returns a tuple with the Segment1TagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpenseApplicationResponseExpenseApplication) GetSegment1TagIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Segment1TagId.Get(), o.Segment1TagId.IsSet()
}

// HasSegment1TagId returns a boolean if a field has been set.
func (o *ExpenseApplicationResponseExpenseApplication) HasSegment1TagId() bool {
	if o != nil && o.Segment1TagId.IsSet() {
		return true
	}

	return false
}

// SetSegment1TagId gets a reference to the given NullableInt64 and assigns it to the Segment1TagId field.
func (o *ExpenseApplicationResponseExpenseApplication) SetSegment1TagId(v int64) {
	o.Segment1TagId.Set(&v)
}
// SetSegment1TagIdNil sets the value for Segment1TagId to be an explicit nil
func (o *ExpenseApplicationResponseExpenseApplication) SetSegment1TagIdNil() {
	o.Segment1TagId.Set(nil)
}

// UnsetSegment1TagId ensures that no value is present for Segment1TagId, not even an explicit nil
func (o *ExpenseApplicationResponseExpenseApplication) UnsetSegment1TagId() {
	o.Segment1TagId.Unset()
}

// GetSegment2TagId returns the Segment2TagId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpenseApplicationResponseExpenseApplication) GetSegment2TagId() int64 {
	if o == nil || o.Segment2TagId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Segment2TagId.Get()
}

// GetSegment2TagIdOk returns a tuple with the Segment2TagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpenseApplicationResponseExpenseApplication) GetSegment2TagIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Segment2TagId.Get(), o.Segment2TagId.IsSet()
}

// HasSegment2TagId returns a boolean if a field has been set.
func (o *ExpenseApplicationResponseExpenseApplication) HasSegment2TagId() bool {
	if o != nil && o.Segment2TagId.IsSet() {
		return true
	}

	return false
}

// SetSegment2TagId gets a reference to the given NullableInt64 and assigns it to the Segment2TagId field.
func (o *ExpenseApplicationResponseExpenseApplication) SetSegment2TagId(v int64) {
	o.Segment2TagId.Set(&v)
}
// SetSegment2TagIdNil sets the value for Segment2TagId to be an explicit nil
func (o *ExpenseApplicationResponseExpenseApplication) SetSegment2TagIdNil() {
	o.Segment2TagId.Set(nil)
}

// UnsetSegment2TagId ensures that no value is present for Segment2TagId, not even an explicit nil
func (o *ExpenseApplicationResponseExpenseApplication) UnsetSegment2TagId() {
	o.Segment2TagId.Unset()
}

// GetSegment3TagId returns the Segment3TagId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpenseApplicationResponseExpenseApplication) GetSegment3TagId() int64 {
	if o == nil || o.Segment3TagId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Segment3TagId.Get()
}

// GetSegment3TagIdOk returns a tuple with the Segment3TagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpenseApplicationResponseExpenseApplication) GetSegment3TagIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Segment3TagId.Get(), o.Segment3TagId.IsSet()
}

// HasSegment3TagId returns a boolean if a field has been set.
func (o *ExpenseApplicationResponseExpenseApplication) HasSegment3TagId() bool {
	if o != nil && o.Segment3TagId.IsSet() {
		return true
	}

	return false
}

// SetSegment3TagId gets a reference to the given NullableInt64 and assigns it to the Segment3TagId field.
func (o *ExpenseApplicationResponseExpenseApplication) SetSegment3TagId(v int64) {
	o.Segment3TagId.Set(&v)
}
// SetSegment3TagIdNil sets the value for Segment3TagId to be an explicit nil
func (o *ExpenseApplicationResponseExpenseApplication) SetSegment3TagIdNil() {
	o.Segment3TagId.Set(nil)
}

// UnsetSegment3TagId ensures that no value is present for Segment3TagId, not even an explicit nil
func (o *ExpenseApplicationResponseExpenseApplication) UnsetSegment3TagId() {
	o.Segment3TagId.Unset()
}

// GetStatus returns the Status field value
func (o *ExpenseApplicationResponseExpenseApplication) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationResponseExpenseApplication) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ExpenseApplicationResponseExpenseApplication) SetStatus(v string) {
	o.Status = v
}

// GetTagIds returns the TagIds field value if set, zero value otherwise.
func (o *ExpenseApplicationResponseExpenseApplication) GetTagIds() []int64 {
	if o == nil || o.TagIds == nil {
		var ret []int64
		return ret
	}
	return o.TagIds
}

// GetTagIdsOk returns a tuple with the TagIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationResponseExpenseApplication) GetTagIdsOk() ([]int64, bool) {
	if o == nil || o.TagIds == nil {
		return nil, false
	}
	return o.TagIds, true
}

// HasTagIds returns a boolean if a field has been set.
func (o *ExpenseApplicationResponseExpenseApplication) HasTagIds() bool {
	if o != nil && o.TagIds != nil {
		return true
	}

	return false
}

// SetTagIds gets a reference to the given []int64 and assigns it to the TagIds field.
func (o *ExpenseApplicationResponseExpenseApplication) SetTagIds(v []int64) {
	o.TagIds = v
}

// GetTitle returns the Title field value
func (o *ExpenseApplicationResponseExpenseApplication) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationResponseExpenseApplication) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ExpenseApplicationResponseExpenseApplication) SetTitle(v string) {
	o.Title = v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *ExpenseApplicationResponseExpenseApplication) GetTotalAmount() int64 {
	if o == nil || o.TotalAmount == nil {
		var ret int64
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationResponseExpenseApplication) GetTotalAmountOk() (*int64, bool) {
	if o == nil || o.TotalAmount == nil {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *ExpenseApplicationResponseExpenseApplication) HasTotalAmount() bool {
	if o != nil && o.TotalAmount != nil {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given int64 and assigns it to the TotalAmount field.
func (o *ExpenseApplicationResponseExpenseApplication) SetTotalAmount(v int64) {
	o.TotalAmount = &v
}

func (o ExpenseApplicationResponseExpenseApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["applicant_id"] = o.ApplicantId
	}
	if true {
		toSerialize["application_number"] = o.ApplicationNumber
	}
	if true {
		toSerialize["approval_flow_logs"] = o.ApprovalFlowLogs
	}
	if true {
		toSerialize["approval_flow_route_id"] = o.ApprovalFlowRouteId
	}
	if true {
		toSerialize["approvers"] = o.Approvers
	}
	if true {
		toSerialize["comments"] = o.Comments
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
	if true {
		toSerialize["deal_id"] = o.DealId.Get()
	}
	if true {
		toSerialize["deal_status"] = o.DealStatus.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if true {
		toSerialize["expense_application_lines"] = o.ExpenseApplicationLines
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["issue_date"] = o.IssueDate
	}
	if o.SectionId.IsSet() {
		toSerialize["section_id"] = o.SectionId.Get()
	}
	if o.Segment1TagId.IsSet() {
		toSerialize["segment_1_tag_id"] = o.Segment1TagId.Get()
	}
	if o.Segment2TagId.IsSet() {
		toSerialize["segment_2_tag_id"] = o.Segment2TagId.Get()
	}
	if o.Segment3TagId.IsSet() {
		toSerialize["segment_3_tag_id"] = o.Segment3TagId.Get()
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.TagIds != nil {
		toSerialize["tag_ids"] = o.TagIds
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if o.TotalAmount != nil {
		toSerialize["total_amount"] = o.TotalAmount
	}
	return json.Marshal(toSerialize)
}

type NullableExpenseApplicationResponseExpenseApplication struct {
	value *ExpenseApplicationResponseExpenseApplication
	isSet bool
}

func (v NullableExpenseApplicationResponseExpenseApplication) Get() *ExpenseApplicationResponseExpenseApplication {
	return v.value
}

func (v *NullableExpenseApplicationResponseExpenseApplication) Set(val *ExpenseApplicationResponseExpenseApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableExpenseApplicationResponseExpenseApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableExpenseApplicationResponseExpenseApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpenseApplicationResponseExpenseApplication(val *ExpenseApplicationResponseExpenseApplication) *NullableExpenseApplicationResponseExpenseApplication {
	return &NullableExpenseApplicationResponseExpenseApplication{value: val, isSet: true}
}

func (v NullableExpenseApplicationResponseExpenseApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpenseApplicationResponseExpenseApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


