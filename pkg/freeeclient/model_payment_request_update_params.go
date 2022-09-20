/*
freee API

 <h1 id=\"freee_api\">freee API</h1> <hr /> <h2 id=\"start_guide\">スタートガイド</h2>  <p>freee API開発がはじめての方は<a href=\"https://developer.freee.co.jp/getting-started\">freee API スタートガイド</a>を参照してください。</p>  <hr /> <h2 id=\"specification\">仕様</h2>  <h3 id=\"api_endpoint\">APIエンドポイント</h3>  <p>https://api.freee.co.jp/ (httpsのみ)</p>  <h3 id=\"about_authorize\">認証について</h3> <p>OAuth2.0を利用します。詳細は<a href=\"https://developer.freee.co.jp/docs\" target=\"_blank\">ドキュメントの認証</a>パートを参照してください。</p>  <h3 id=\"data_format\">データフォーマット</h3>  <p>リクエスト、レスポンスともにJSON形式をサポートしていますが、詳細は、API毎の説明欄（application/jsonなど）を確認してください。</p>  <h3 id=\"compatibility\">後方互換性ありの変更</h3>  <p>freeeでは、APIを改善していくために以下のような変更は後方互換性ありとして通知なく変更を入れることがあります。アプリケーション実装者は以下を踏まえて開発を行ってください。</p>  <ul> <li>新しいAPIリソース・エンドポイントの追加</li> <li>既存のAPIに対して必須ではない新しいリクエストパラメータの追加</li> <li>既存のAPIレスポンスに対する新しいプロパティの追加</li> <li>既存のAPIレスポンスに対するプロパティの順番の入れ変え</li> <li>keyとなっているidやcodeの長さの変更（長くする）</li> </ul>  <h3 id=\"common_response_header\">共通レスポンスヘッダー</h3>  <p>すべてのAPIのレスポンスには以下のHTTPヘッダーが含まれます。</p>  <ul> <li> <p>X-Freee-Request-ID</p> <ul> <li>各リクエスト毎に発行されるID</li> </ul> </li> </ul>  <h3 id=\"common_error_response\">共通エラーレスポンス</h3>  <ul> <li> <p>ステータスコードはレスポンス内のJSONに含まれる他、HTTPヘッダにも含まれる</p> </li> <li> <p>一部のエラーレスポンスにはエラーコードが含まれます。<br>詳細は、<a href=\"https://developer.freee.co.jp/tips/faq/40x-checkpoint\">HTTPステータスコード400台エラー時のチェックポイント</a>を参照してください</p> </li> <p>type</p>  <ul> <li>status : HTTPステータスコードの説明</li>  <li>validation : エラーの詳細の説明（開発者向け）</li> </ul> </li> </ul>  <p>レスポンスの例</p>  <pre><code>  {     &quot;status_code&quot; : 400,     &quot;errors&quot; : [       {         &quot;type&quot; : &quot;status&quot;,         &quot;messages&quot; : [&quot;不正なリクエストです。&quot;]       },       {         &quot;type&quot; : &quot;validation&quot;,         &quot;messages&quot; : [&quot;Date は不正な日付フォーマットです。入力例：2019-12-17&quot;]       }     ]   }</code></pre>  </br>  <h3 id=\"api_rate_limit\">API使用制限</h3>    <p>freeeは一定期間に過度のアクセスを検知した場合、APIアクセスをコントロールする場合があります。</p>   <p>その際のhttp status codeは403となります。制限がかかってから10分程度が過ぎると再度使用することができるようになります。</p>  <h4 id=\"reports_api_endpoint\">/reportsと/receipts/{id}/downloadエンドポイント</h4>  <p>freeeはエンドポイント毎に一定頻度以上のアクセスを検知した場合、APIアクセスをコントロールする場合があります。その際のhttp status codeは429（too many requests）となります。</p>  <ul>   <li>/reports:1秒に10回まで</li>   <li>/receipts/{id}/download:1秒に3回まで</li> </ul>  <p>レスポンスボディのmetaプロパティに以下を含めます。</p>  <ul>   <li>設定されている上限値</li>   <li>上限に達するまでの使用可能回数</li>   <li>（上限値に達した場合）使用回数がリセットされる時刻</li> </ul>  <h3 id=\"plan_api_rate_limit\">プラン別のAPI Rate Limit</h3>   <table border=\"1\">     <tbody>       <tr>         <th style=\"padding: 10px\"><strong>freee会計プラン名</strong></th>         <th style=\"padding: 10px\"><strong>事業所とアプリケーション毎に1日でのAPIコール数</strong></th>       </tr>       <tr>         <td style=\"padding: 10px\">エンタープライズ</td>         <td style=\"padding: 10px\">10,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">プロフェッショナル</td>         <td style=\"padding: 10px\">5,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">ベーシック</td>         <td style=\"padding: 10px\">3,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">ミニマム</td>         <td style=\"padding: 10px\">3,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">上記以外</td>         <td style=\"padding: 10px\">3,000</td>       </tr>     </tbody>   </table>  <h3 id=\"webhook\">Webhookについて</h3>  <p>詳細は<a href=\"https://developer.freee.co.jp/docs/accounting/webhook\" target=\"_blank\">会計Webhook概要</a>を参照してください。</p>  <hr /> <h2 id=\"contact\">連絡先</h2>  <p>ご不明点、ご要望等は <a href=\"https://support.freee.co.jp/hc/ja/requests/new\">freee サポートデスクへのお問い合わせフォーム</a> からご連絡ください。</p> <hr />&copy; Since 2013 freee K.K.

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github.com/Yoii-Inc/freee-gen-csv

import (
	"encoding/json"
)

// PaymentRequestUpdateParams struct for PaymentRequestUpdateParams
type PaymentRequestUpdateParams struct {
	// 受取人名（カナ）（48文字以内）<br> 支払先指定時には無効 
	AccountName *string `json:"account_name,omitempty"`
	// 口座番号（半角数字1桁〜7桁）<br> 支払先指定時には無効 
	AccountNumber *string `json:"account_number,omitempty"`
	// '口座種別(ordinary: 普通、checking: 当座、earmarked: 納税準備預金、savings: 貯蓄、other: その他)'<br> '支払先指定時には無効'<br> 'デフォルトは ordinary: 普通 です' 
	AccountType *string `json:"account_type,omitempty"`
	// 申請日 (yyyy-mm-dd)<br> 指定しない場合は当日の日付が登録されます。<br> 申請者が、下書き状態もしくは差戻し状態の支払依頼に対して指定する場合のみ有効 
	ApplicationDate *string `json:"application_date,omitempty"`
	// 申請経路ID<br> 指定する申請経路IDは、申請経路APIを利用して取得してください。 
	ApprovalFlowRouteId int64 `json:"approval_flow_route_id"`
	// 承認者のユーザーID<br> 「承認者を指定」の経路を申請経路として使用する場合に指定してください。<br> 指定する承認者のユーザーIDは、申請経路APIを利用して取得してください。 
	ApproverId *int64 `json:"approver_id,omitempty"`
	// 銀行コード（半角数字1桁〜4桁）<br> 支払先指定時には無効 
	BankCode *string `json:"bank_code,omitempty"`
	// 銀行名（255文字以内）<br> 支払先指定時には無効 
	BankName *string `json:"bank_name,omitempty"`
	// 銀行名（カナ）（15文字以内）<br> 支払先指定時には無効 
	BankNameKana *string `json:"bank_name_kana,omitempty"`
	// 支店番号（半角数字1桁〜3桁）<br> 支払先指定時には無効 
	BranchCode *string `json:"branch_code,omitempty"`
	// 支店名（カナ）（15文字以内）<br> 指定可能な文字は、英数・カナ・丸括弧・ハイフン・スペースのみです。<br> 支払先指定時には無効 
	BranchKana *string `json:"branch_kana,omitempty"`
	// 支店名（255文字以内）<br> 支払先指定時には無効 
	BranchName *string `json:"branch_name,omitempty"`
	// 事業所ID
	CompanyId int64 `json:"company_id"`
	// 備考
	Description *string `json:"description,omitempty"`
	// 請求書番号（255文字以内）
	DocumentCode *string `json:"document_code,omitempty"`
	// 支払依頼のステータス<br> falseを指定した時は申請中（in_progress）で支払依頼を更新します。<br> trueを指定した時は下書き（draft）で支払依頼を更新します。 
	Draft bool `json:"draft"`
	// 発生日 (yyyy-mm-dd)
	IssueDate string `json:"issue_date"`
	// 支払先の取引先コード<br> 支払先の取引先ID指定時には無効 
	PartnerCode NullableString `json:"partner_code,omitempty"`
	// 支払先の取引先ID
	PartnerId NullableInt64 `json:"partner_id,omitempty"`
	// 支払期限 (yyyy-mm-dd)
	PaymentDate NullableString `json:"payment_date,omitempty"`
	// '支払方法(none: 指定なし, domestic_bank_transfer: 国内振込, abroad_bank_transfer: 国外振込, account_transfer: 口座振替, credit_card: クレジットカード)'<br> 'デフォルトは none: 指定なし です。' 
	PaymentMethod *string `json:"payment_method,omitempty"`
	// 支払依頼の項目行一覧（配列）
	PaymentRequestLines []PaymentRequestUpdateParamsPaymentRequestLinesInner `json:"payment_request_lines"`
	// 証憑ファイルID（ファイルボックスのファイルID）（配列）
	ReceiptIds []int64 `json:"receipt_ids,omitempty"`
	// 申請タイトル<br> 申請者が、下書き状態もしくは差戻し状態の支払依頼に対して指定する場合のみ有効 
	Title string `json:"title"`
}

// NewPaymentRequestUpdateParams instantiates a new PaymentRequestUpdateParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentRequestUpdateParams(approvalFlowRouteId int64, companyId int64, draft bool, issueDate string, paymentRequestLines []PaymentRequestUpdateParamsPaymentRequestLinesInner, title string) *PaymentRequestUpdateParams {
	this := PaymentRequestUpdateParams{}
	this.ApprovalFlowRouteId = approvalFlowRouteId
	this.CompanyId = companyId
	this.Draft = draft
	this.IssueDate = issueDate
	this.PaymentRequestLines = paymentRequestLines
	this.Title = title
	return &this
}

// NewPaymentRequestUpdateParamsWithDefaults instantiates a new PaymentRequestUpdateParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentRequestUpdateParamsWithDefaults() *PaymentRequestUpdateParams {
	this := PaymentRequestUpdateParams{}
	return &this
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *PaymentRequestUpdateParams) GetAccountName() string {
	if o == nil || o.AccountName == nil {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParams) GetAccountNameOk() (*string, bool) {
	if o == nil || o.AccountName == nil {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *PaymentRequestUpdateParams) HasAccountName() bool {
	if o != nil && o.AccountName != nil {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *PaymentRequestUpdateParams) SetAccountName(v string) {
	o.AccountName = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *PaymentRequestUpdateParams) GetAccountNumber() string {
	if o == nil || o.AccountNumber == nil {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParams) GetAccountNumberOk() (*string, bool) {
	if o == nil || o.AccountNumber == nil {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *PaymentRequestUpdateParams) HasAccountNumber() bool {
	if o != nil && o.AccountNumber != nil {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *PaymentRequestUpdateParams) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *PaymentRequestUpdateParams) GetAccountType() string {
	if o == nil || o.AccountType == nil {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParams) GetAccountTypeOk() (*string, bool) {
	if o == nil || o.AccountType == nil {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *PaymentRequestUpdateParams) HasAccountType() bool {
	if o != nil && o.AccountType != nil {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *PaymentRequestUpdateParams) SetAccountType(v string) {
	o.AccountType = &v
}

// GetApplicationDate returns the ApplicationDate field value if set, zero value otherwise.
func (o *PaymentRequestUpdateParams) GetApplicationDate() string {
	if o == nil || o.ApplicationDate == nil {
		var ret string
		return ret
	}
	return *o.ApplicationDate
}

// GetApplicationDateOk returns a tuple with the ApplicationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParams) GetApplicationDateOk() (*string, bool) {
	if o == nil || o.ApplicationDate == nil {
		return nil, false
	}
	return o.ApplicationDate, true
}

// HasApplicationDate returns a boolean if a field has been set.
func (o *PaymentRequestUpdateParams) HasApplicationDate() bool {
	if o != nil && o.ApplicationDate != nil {
		return true
	}

	return false
}

// SetApplicationDate gets a reference to the given string and assigns it to the ApplicationDate field.
func (o *PaymentRequestUpdateParams) SetApplicationDate(v string) {
	o.ApplicationDate = &v
}

// GetApprovalFlowRouteId returns the ApprovalFlowRouteId field value
func (o *PaymentRequestUpdateParams) GetApprovalFlowRouteId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ApprovalFlowRouteId
}

// GetApprovalFlowRouteIdOk returns a tuple with the ApprovalFlowRouteId field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParams) GetApprovalFlowRouteIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApprovalFlowRouteId, true
}

// SetApprovalFlowRouteId sets field value
func (o *PaymentRequestUpdateParams) SetApprovalFlowRouteId(v int64) {
	o.ApprovalFlowRouteId = v
}

// GetApproverId returns the ApproverId field value if set, zero value otherwise.
func (o *PaymentRequestUpdateParams) GetApproverId() int64 {
	if o == nil || o.ApproverId == nil {
		var ret int64
		return ret
	}
	return *o.ApproverId
}

// GetApproverIdOk returns a tuple with the ApproverId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParams) GetApproverIdOk() (*int64, bool) {
	if o == nil || o.ApproverId == nil {
		return nil, false
	}
	return o.ApproverId, true
}

// HasApproverId returns a boolean if a field has been set.
func (o *PaymentRequestUpdateParams) HasApproverId() bool {
	if o != nil && o.ApproverId != nil {
		return true
	}

	return false
}

// SetApproverId gets a reference to the given int64 and assigns it to the ApproverId field.
func (o *PaymentRequestUpdateParams) SetApproverId(v int64) {
	o.ApproverId = &v
}

// GetBankCode returns the BankCode field value if set, zero value otherwise.
func (o *PaymentRequestUpdateParams) GetBankCode() string {
	if o == nil || o.BankCode == nil {
		var ret string
		return ret
	}
	return *o.BankCode
}

// GetBankCodeOk returns a tuple with the BankCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParams) GetBankCodeOk() (*string, bool) {
	if o == nil || o.BankCode == nil {
		return nil, false
	}
	return o.BankCode, true
}

// HasBankCode returns a boolean if a field has been set.
func (o *PaymentRequestUpdateParams) HasBankCode() bool {
	if o != nil && o.BankCode != nil {
		return true
	}

	return false
}

// SetBankCode gets a reference to the given string and assigns it to the BankCode field.
func (o *PaymentRequestUpdateParams) SetBankCode(v string) {
	o.BankCode = &v
}

// GetBankName returns the BankName field value if set, zero value otherwise.
func (o *PaymentRequestUpdateParams) GetBankName() string {
	if o == nil || o.BankName == nil {
		var ret string
		return ret
	}
	return *o.BankName
}

// GetBankNameOk returns a tuple with the BankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParams) GetBankNameOk() (*string, bool) {
	if o == nil || o.BankName == nil {
		return nil, false
	}
	return o.BankName, true
}

// HasBankName returns a boolean if a field has been set.
func (o *PaymentRequestUpdateParams) HasBankName() bool {
	if o != nil && o.BankName != nil {
		return true
	}

	return false
}

// SetBankName gets a reference to the given string and assigns it to the BankName field.
func (o *PaymentRequestUpdateParams) SetBankName(v string) {
	o.BankName = &v
}

// GetBankNameKana returns the BankNameKana field value if set, zero value otherwise.
func (o *PaymentRequestUpdateParams) GetBankNameKana() string {
	if o == nil || o.BankNameKana == nil {
		var ret string
		return ret
	}
	return *o.BankNameKana
}

// GetBankNameKanaOk returns a tuple with the BankNameKana field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParams) GetBankNameKanaOk() (*string, bool) {
	if o == nil || o.BankNameKana == nil {
		return nil, false
	}
	return o.BankNameKana, true
}

// HasBankNameKana returns a boolean if a field has been set.
func (o *PaymentRequestUpdateParams) HasBankNameKana() bool {
	if o != nil && o.BankNameKana != nil {
		return true
	}

	return false
}

// SetBankNameKana gets a reference to the given string and assigns it to the BankNameKana field.
func (o *PaymentRequestUpdateParams) SetBankNameKana(v string) {
	o.BankNameKana = &v
}

// GetBranchCode returns the BranchCode field value if set, zero value otherwise.
func (o *PaymentRequestUpdateParams) GetBranchCode() string {
	if o == nil || o.BranchCode == nil {
		var ret string
		return ret
	}
	return *o.BranchCode
}

// GetBranchCodeOk returns a tuple with the BranchCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParams) GetBranchCodeOk() (*string, bool) {
	if o == nil || o.BranchCode == nil {
		return nil, false
	}
	return o.BranchCode, true
}

// HasBranchCode returns a boolean if a field has been set.
func (o *PaymentRequestUpdateParams) HasBranchCode() bool {
	if o != nil && o.BranchCode != nil {
		return true
	}

	return false
}

// SetBranchCode gets a reference to the given string and assigns it to the BranchCode field.
func (o *PaymentRequestUpdateParams) SetBranchCode(v string) {
	o.BranchCode = &v
}

// GetBranchKana returns the BranchKana field value if set, zero value otherwise.
func (o *PaymentRequestUpdateParams) GetBranchKana() string {
	if o == nil || o.BranchKana == nil {
		var ret string
		return ret
	}
	return *o.BranchKana
}

// GetBranchKanaOk returns a tuple with the BranchKana field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParams) GetBranchKanaOk() (*string, bool) {
	if o == nil || o.BranchKana == nil {
		return nil, false
	}
	return o.BranchKana, true
}

// HasBranchKana returns a boolean if a field has been set.
func (o *PaymentRequestUpdateParams) HasBranchKana() bool {
	if o != nil && o.BranchKana != nil {
		return true
	}

	return false
}

// SetBranchKana gets a reference to the given string and assigns it to the BranchKana field.
func (o *PaymentRequestUpdateParams) SetBranchKana(v string) {
	o.BranchKana = &v
}

// GetBranchName returns the BranchName field value if set, zero value otherwise.
func (o *PaymentRequestUpdateParams) GetBranchName() string {
	if o == nil || o.BranchName == nil {
		var ret string
		return ret
	}
	return *o.BranchName
}

// GetBranchNameOk returns a tuple with the BranchName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParams) GetBranchNameOk() (*string, bool) {
	if o == nil || o.BranchName == nil {
		return nil, false
	}
	return o.BranchName, true
}

// HasBranchName returns a boolean if a field has been set.
func (o *PaymentRequestUpdateParams) HasBranchName() bool {
	if o != nil && o.BranchName != nil {
		return true
	}

	return false
}

// SetBranchName gets a reference to the given string and assigns it to the BranchName field.
func (o *PaymentRequestUpdateParams) SetBranchName(v string) {
	o.BranchName = &v
}

// GetCompanyId returns the CompanyId field value
func (o *PaymentRequestUpdateParams) GetCompanyId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParams) GetCompanyIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyId, true
}

// SetCompanyId sets field value
func (o *PaymentRequestUpdateParams) SetCompanyId(v int64) {
	o.CompanyId = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PaymentRequestUpdateParams) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParams) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PaymentRequestUpdateParams) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PaymentRequestUpdateParams) SetDescription(v string) {
	o.Description = &v
}

// GetDocumentCode returns the DocumentCode field value if set, zero value otherwise.
func (o *PaymentRequestUpdateParams) GetDocumentCode() string {
	if o == nil || o.DocumentCode == nil {
		var ret string
		return ret
	}
	return *o.DocumentCode
}

// GetDocumentCodeOk returns a tuple with the DocumentCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParams) GetDocumentCodeOk() (*string, bool) {
	if o == nil || o.DocumentCode == nil {
		return nil, false
	}
	return o.DocumentCode, true
}

// HasDocumentCode returns a boolean if a field has been set.
func (o *PaymentRequestUpdateParams) HasDocumentCode() bool {
	if o != nil && o.DocumentCode != nil {
		return true
	}

	return false
}

// SetDocumentCode gets a reference to the given string and assigns it to the DocumentCode field.
func (o *PaymentRequestUpdateParams) SetDocumentCode(v string) {
	o.DocumentCode = &v
}

// GetDraft returns the Draft field value
func (o *PaymentRequestUpdateParams) GetDraft() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Draft
}

// GetDraftOk returns a tuple with the Draft field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParams) GetDraftOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Draft, true
}

// SetDraft sets field value
func (o *PaymentRequestUpdateParams) SetDraft(v bool) {
	o.Draft = v
}

// GetIssueDate returns the IssueDate field value
func (o *PaymentRequestUpdateParams) GetIssueDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssueDate
}

// GetIssueDateOk returns a tuple with the IssueDate field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParams) GetIssueDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueDate, true
}

// SetIssueDate sets field value
func (o *PaymentRequestUpdateParams) SetIssueDate(v string) {
	o.IssueDate = v
}

// GetPartnerCode returns the PartnerCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentRequestUpdateParams) GetPartnerCode() string {
	if o == nil || o.PartnerCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.PartnerCode.Get()
}

// GetPartnerCodeOk returns a tuple with the PartnerCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestUpdateParams) GetPartnerCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PartnerCode.Get(), o.PartnerCode.IsSet()
}

// HasPartnerCode returns a boolean if a field has been set.
func (o *PaymentRequestUpdateParams) HasPartnerCode() bool {
	if o != nil && o.PartnerCode.IsSet() {
		return true
	}

	return false
}

// SetPartnerCode gets a reference to the given NullableString and assigns it to the PartnerCode field.
func (o *PaymentRequestUpdateParams) SetPartnerCode(v string) {
	o.PartnerCode.Set(&v)
}
// SetPartnerCodeNil sets the value for PartnerCode to be an explicit nil
func (o *PaymentRequestUpdateParams) SetPartnerCodeNil() {
	o.PartnerCode.Set(nil)
}

// UnsetPartnerCode ensures that no value is present for PartnerCode, not even an explicit nil
func (o *PaymentRequestUpdateParams) UnsetPartnerCode() {
	o.PartnerCode.Unset()
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentRequestUpdateParams) GetPartnerId() int64 {
	if o == nil || o.PartnerId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.PartnerId.Get()
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestUpdateParams) GetPartnerIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PartnerId.Get(), o.PartnerId.IsSet()
}

// HasPartnerId returns a boolean if a field has been set.
func (o *PaymentRequestUpdateParams) HasPartnerId() bool {
	if o != nil && o.PartnerId.IsSet() {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given NullableInt64 and assigns it to the PartnerId field.
func (o *PaymentRequestUpdateParams) SetPartnerId(v int64) {
	o.PartnerId.Set(&v)
}
// SetPartnerIdNil sets the value for PartnerId to be an explicit nil
func (o *PaymentRequestUpdateParams) SetPartnerIdNil() {
	o.PartnerId.Set(nil)
}

// UnsetPartnerId ensures that no value is present for PartnerId, not even an explicit nil
func (o *PaymentRequestUpdateParams) UnsetPartnerId() {
	o.PartnerId.Unset()
}

// GetPaymentDate returns the PaymentDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentRequestUpdateParams) GetPaymentDate() string {
	if o == nil || o.PaymentDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.PaymentDate.Get()
}

// GetPaymentDateOk returns a tuple with the PaymentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestUpdateParams) GetPaymentDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentDate.Get(), o.PaymentDate.IsSet()
}

// HasPaymentDate returns a boolean if a field has been set.
func (o *PaymentRequestUpdateParams) HasPaymentDate() bool {
	if o != nil && o.PaymentDate.IsSet() {
		return true
	}

	return false
}

// SetPaymentDate gets a reference to the given NullableString and assigns it to the PaymentDate field.
func (o *PaymentRequestUpdateParams) SetPaymentDate(v string) {
	o.PaymentDate.Set(&v)
}
// SetPaymentDateNil sets the value for PaymentDate to be an explicit nil
func (o *PaymentRequestUpdateParams) SetPaymentDateNil() {
	o.PaymentDate.Set(nil)
}

// UnsetPaymentDate ensures that no value is present for PaymentDate, not even an explicit nil
func (o *PaymentRequestUpdateParams) UnsetPaymentDate() {
	o.PaymentDate.Unset()
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *PaymentRequestUpdateParams) GetPaymentMethod() string {
	if o == nil || o.PaymentMethod == nil {
		var ret string
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParams) GetPaymentMethodOk() (*string, bool) {
	if o == nil || o.PaymentMethod == nil {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *PaymentRequestUpdateParams) HasPaymentMethod() bool {
	if o != nil && o.PaymentMethod != nil {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given string and assigns it to the PaymentMethod field.
func (o *PaymentRequestUpdateParams) SetPaymentMethod(v string) {
	o.PaymentMethod = &v
}

// GetPaymentRequestLines returns the PaymentRequestLines field value
func (o *PaymentRequestUpdateParams) GetPaymentRequestLines() []PaymentRequestUpdateParamsPaymentRequestLinesInner {
	if o == nil {
		var ret []PaymentRequestUpdateParamsPaymentRequestLinesInner
		return ret
	}

	return o.PaymentRequestLines
}

// GetPaymentRequestLinesOk returns a tuple with the PaymentRequestLines field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParams) GetPaymentRequestLinesOk() ([]PaymentRequestUpdateParamsPaymentRequestLinesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentRequestLines, true
}

// SetPaymentRequestLines sets field value
func (o *PaymentRequestUpdateParams) SetPaymentRequestLines(v []PaymentRequestUpdateParamsPaymentRequestLinesInner) {
	o.PaymentRequestLines = v
}

// GetReceiptIds returns the ReceiptIds field value if set, zero value otherwise.
func (o *PaymentRequestUpdateParams) GetReceiptIds() []int64 {
	if o == nil || o.ReceiptIds == nil {
		var ret []int64
		return ret
	}
	return o.ReceiptIds
}

// GetReceiptIdsOk returns a tuple with the ReceiptIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParams) GetReceiptIdsOk() ([]int64, bool) {
	if o == nil || o.ReceiptIds == nil {
		return nil, false
	}
	return o.ReceiptIds, true
}

// HasReceiptIds returns a boolean if a field has been set.
func (o *PaymentRequestUpdateParams) HasReceiptIds() bool {
	if o != nil && o.ReceiptIds != nil {
		return true
	}

	return false
}

// SetReceiptIds gets a reference to the given []int64 and assigns it to the ReceiptIds field.
func (o *PaymentRequestUpdateParams) SetReceiptIds(v []int64) {
	o.ReceiptIds = v
}

// GetTitle returns the Title field value
func (o *PaymentRequestUpdateParams) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestUpdateParams) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *PaymentRequestUpdateParams) SetTitle(v string) {
	o.Title = v
}

func (o PaymentRequestUpdateParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountName != nil {
		toSerialize["account_name"] = o.AccountName
	}
	if o.AccountNumber != nil {
		toSerialize["account_number"] = o.AccountNumber
	}
	if o.AccountType != nil {
		toSerialize["account_type"] = o.AccountType
	}
	if o.ApplicationDate != nil {
		toSerialize["application_date"] = o.ApplicationDate
	}
	if true {
		toSerialize["approval_flow_route_id"] = o.ApprovalFlowRouteId
	}
	if o.ApproverId != nil {
		toSerialize["approver_id"] = o.ApproverId
	}
	if o.BankCode != nil {
		toSerialize["bank_code"] = o.BankCode
	}
	if o.BankName != nil {
		toSerialize["bank_name"] = o.BankName
	}
	if o.BankNameKana != nil {
		toSerialize["bank_name_kana"] = o.BankNameKana
	}
	if o.BranchCode != nil {
		toSerialize["branch_code"] = o.BranchCode
	}
	if o.BranchKana != nil {
		toSerialize["branch_kana"] = o.BranchKana
	}
	if o.BranchName != nil {
		toSerialize["branch_name"] = o.BranchName
	}
	if true {
		toSerialize["company_id"] = o.CompanyId
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DocumentCode != nil {
		toSerialize["document_code"] = o.DocumentCode
	}
	if true {
		toSerialize["draft"] = o.Draft
	}
	if true {
		toSerialize["issue_date"] = o.IssueDate
	}
	if o.PartnerCode.IsSet() {
		toSerialize["partner_code"] = o.PartnerCode.Get()
	}
	if o.PartnerId.IsSet() {
		toSerialize["partner_id"] = o.PartnerId.Get()
	}
	if o.PaymentDate.IsSet() {
		toSerialize["payment_date"] = o.PaymentDate.Get()
	}
	if o.PaymentMethod != nil {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	if true {
		toSerialize["payment_request_lines"] = o.PaymentRequestLines
	}
	if o.ReceiptIds != nil {
		toSerialize["receipt_ids"] = o.ReceiptIds
	}
	if true {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentRequestUpdateParams struct {
	value *PaymentRequestUpdateParams
	isSet bool
}

func (v NullablePaymentRequestUpdateParams) Get() *PaymentRequestUpdateParams {
	return v.value
}

func (v *NullablePaymentRequestUpdateParams) Set(val *PaymentRequestUpdateParams) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentRequestUpdateParams) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentRequestUpdateParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentRequestUpdateParams(val *PaymentRequestUpdateParams) *NullablePaymentRequestUpdateParams {
	return &NullablePaymentRequestUpdateParams{value: val, isSet: true}
}

func (v NullablePaymentRequestUpdateParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentRequestUpdateParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


