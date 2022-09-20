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

// PartnerCreateParams struct for PartnerCreateParams
type PartnerCreateParams struct {
	AddressAttributes *PartnerCreateParamsAddressAttributes `json:"address_attributes,omitempty"`
	// 取引先コード（取引先コードの利用を有効にしている場合は、codeの指定は必須です。）
	Code *string `json:"code,omitempty"`
	// 事業所ID
	CompanyId int32 `json:"company_id"`
	// 担当者 氏名 (255文字以内)
	ContactName *string `json:"contact_name,omitempty"`
	// 地域（JP: 国内、ZZ:国外）、指定しない場合JPになります。
	CountryCode *string `json:"country_code,omitempty"`
	// 敬称（御中、様、(空白)の3つから選択）
	DefaultTitle *string `json:"default_title,omitempty"`
	// 担当者 メールアドレス (255文字以内)
	Email *string `json:"email,omitempty"`
	InvoicePaymentTermAttributes *PartnerCreateParamsInvoicePaymentTermAttributes `json:"invoice_payment_term_attributes,omitempty"`
	// 正式名称（255文字以内）
	LongName *string `json:"long_name,omitempty"`
	// 取引先名 (255文字以内)
	Name string `json:"name"`
	// カナ名称（255文字以内）
	NameKana *string `json:"name_kana,omitempty"`
	// 事業所種別（null: 未設定、1: 法人、2: 個人）
	OrgCode NullableInt32 `json:"org_code,omitempty"`
	PartnerBankAccountAttributes *PartnerCreateParamsPartnerBankAccountAttributes `json:"partner_bank_account_attributes,omitempty"`
	PartnerDocSettingAttributes *PartnerCreateParamsPartnerDocSettingAttributes `json:"partner_doc_setting_attributes,omitempty"`
	// 振込元口座ID（一括振込ファイル用）:（walletableのtypeが'bank_account'のidのみ指定できます。また、未設定にする場合は、nullを指定してください。）
	PayerWalletableId NullableInt32 `json:"payer_walletable_id,omitempty"`
	PaymentTermAttributes *PartnerCreateParamsPaymentTermAttributes `json:"payment_term_attributes,omitempty"`
	// 電話番号
	Phone *string `json:"phone,omitempty"`
	// ショートカット１ (255文字以内)
	Shortcut1 *string `json:"shortcut1,omitempty"`
	// ショートカット２ (255文字以内)
	Shortcut2 *string `json:"shortcut2,omitempty"`
	// 振込手数料負担（一括振込ファイル用）: (振込元(当方): payer, 振込先(先方): payee)、指定しない場合payerになります。
	TransferFeeHandlingSide *string `json:"transfer_fee_handling_side,omitempty"`
}

// NewPartnerCreateParams instantiates a new PartnerCreateParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerCreateParams(companyId int32, name string) *PartnerCreateParams {
	this := PartnerCreateParams{}
	this.CompanyId = companyId
	this.Name = name
	return &this
}

// NewPartnerCreateParamsWithDefaults instantiates a new PartnerCreateParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerCreateParamsWithDefaults() *PartnerCreateParams {
	this := PartnerCreateParams{}
	return &this
}

// GetAddressAttributes returns the AddressAttributes field value if set, zero value otherwise.
func (o *PartnerCreateParams) GetAddressAttributes() PartnerCreateParamsAddressAttributes {
	if o == nil || o.AddressAttributes == nil {
		var ret PartnerCreateParamsAddressAttributes
		return ret
	}
	return *o.AddressAttributes
}

// GetAddressAttributesOk returns a tuple with the AddressAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCreateParams) GetAddressAttributesOk() (*PartnerCreateParamsAddressAttributes, bool) {
	if o == nil || o.AddressAttributes == nil {
		return nil, false
	}
	return o.AddressAttributes, true
}

// HasAddressAttributes returns a boolean if a field has been set.
func (o *PartnerCreateParams) HasAddressAttributes() bool {
	if o != nil && o.AddressAttributes != nil {
		return true
	}

	return false
}

// SetAddressAttributes gets a reference to the given PartnerCreateParamsAddressAttributes and assigns it to the AddressAttributes field.
func (o *PartnerCreateParams) SetAddressAttributes(v PartnerCreateParamsAddressAttributes) {
	o.AddressAttributes = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *PartnerCreateParams) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCreateParams) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *PartnerCreateParams) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *PartnerCreateParams) SetCode(v string) {
	o.Code = &v
}

// GetCompanyId returns the CompanyId field value
func (o *PartnerCreateParams) GetCompanyId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value
// and a boolean to check if the value has been set.
func (o *PartnerCreateParams) GetCompanyIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyId, true
}

// SetCompanyId sets field value
func (o *PartnerCreateParams) SetCompanyId(v int32) {
	o.CompanyId = v
}

// GetContactName returns the ContactName field value if set, zero value otherwise.
func (o *PartnerCreateParams) GetContactName() string {
	if o == nil || o.ContactName == nil {
		var ret string
		return ret
	}
	return *o.ContactName
}

// GetContactNameOk returns a tuple with the ContactName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCreateParams) GetContactNameOk() (*string, bool) {
	if o == nil || o.ContactName == nil {
		return nil, false
	}
	return o.ContactName, true
}

// HasContactName returns a boolean if a field has been set.
func (o *PartnerCreateParams) HasContactName() bool {
	if o != nil && o.ContactName != nil {
		return true
	}

	return false
}

// SetContactName gets a reference to the given string and assigns it to the ContactName field.
func (o *PartnerCreateParams) SetContactName(v string) {
	o.ContactName = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *PartnerCreateParams) GetCountryCode() string {
	if o == nil || o.CountryCode == nil {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCreateParams) GetCountryCodeOk() (*string, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *PartnerCreateParams) HasCountryCode() bool {
	if o != nil && o.CountryCode != nil {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *PartnerCreateParams) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetDefaultTitle returns the DefaultTitle field value if set, zero value otherwise.
func (o *PartnerCreateParams) GetDefaultTitle() string {
	if o == nil || o.DefaultTitle == nil {
		var ret string
		return ret
	}
	return *o.DefaultTitle
}

// GetDefaultTitleOk returns a tuple with the DefaultTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCreateParams) GetDefaultTitleOk() (*string, bool) {
	if o == nil || o.DefaultTitle == nil {
		return nil, false
	}
	return o.DefaultTitle, true
}

// HasDefaultTitle returns a boolean if a field has been set.
func (o *PartnerCreateParams) HasDefaultTitle() bool {
	if o != nil && o.DefaultTitle != nil {
		return true
	}

	return false
}

// SetDefaultTitle gets a reference to the given string and assigns it to the DefaultTitle field.
func (o *PartnerCreateParams) SetDefaultTitle(v string) {
	o.DefaultTitle = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PartnerCreateParams) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCreateParams) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PartnerCreateParams) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PartnerCreateParams) SetEmail(v string) {
	o.Email = &v
}

// GetInvoicePaymentTermAttributes returns the InvoicePaymentTermAttributes field value if set, zero value otherwise.
func (o *PartnerCreateParams) GetInvoicePaymentTermAttributes() PartnerCreateParamsInvoicePaymentTermAttributes {
	if o == nil || o.InvoicePaymentTermAttributes == nil {
		var ret PartnerCreateParamsInvoicePaymentTermAttributes
		return ret
	}
	return *o.InvoicePaymentTermAttributes
}

// GetInvoicePaymentTermAttributesOk returns a tuple with the InvoicePaymentTermAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCreateParams) GetInvoicePaymentTermAttributesOk() (*PartnerCreateParamsInvoicePaymentTermAttributes, bool) {
	if o == nil || o.InvoicePaymentTermAttributes == nil {
		return nil, false
	}
	return o.InvoicePaymentTermAttributes, true
}

// HasInvoicePaymentTermAttributes returns a boolean if a field has been set.
func (o *PartnerCreateParams) HasInvoicePaymentTermAttributes() bool {
	if o != nil && o.InvoicePaymentTermAttributes != nil {
		return true
	}

	return false
}

// SetInvoicePaymentTermAttributes gets a reference to the given PartnerCreateParamsInvoicePaymentTermAttributes and assigns it to the InvoicePaymentTermAttributes field.
func (o *PartnerCreateParams) SetInvoicePaymentTermAttributes(v PartnerCreateParamsInvoicePaymentTermAttributes) {
	o.InvoicePaymentTermAttributes = &v
}

// GetLongName returns the LongName field value if set, zero value otherwise.
func (o *PartnerCreateParams) GetLongName() string {
	if o == nil || o.LongName == nil {
		var ret string
		return ret
	}
	return *o.LongName
}

// GetLongNameOk returns a tuple with the LongName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCreateParams) GetLongNameOk() (*string, bool) {
	if o == nil || o.LongName == nil {
		return nil, false
	}
	return o.LongName, true
}

// HasLongName returns a boolean if a field has been set.
func (o *PartnerCreateParams) HasLongName() bool {
	if o != nil && o.LongName != nil {
		return true
	}

	return false
}

// SetLongName gets a reference to the given string and assigns it to the LongName field.
func (o *PartnerCreateParams) SetLongName(v string) {
	o.LongName = &v
}

// GetName returns the Name field value
func (o *PartnerCreateParams) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PartnerCreateParams) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PartnerCreateParams) SetName(v string) {
	o.Name = v
}

// GetNameKana returns the NameKana field value if set, zero value otherwise.
func (o *PartnerCreateParams) GetNameKana() string {
	if o == nil || o.NameKana == nil {
		var ret string
		return ret
	}
	return *o.NameKana
}

// GetNameKanaOk returns a tuple with the NameKana field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCreateParams) GetNameKanaOk() (*string, bool) {
	if o == nil || o.NameKana == nil {
		return nil, false
	}
	return o.NameKana, true
}

// HasNameKana returns a boolean if a field has been set.
func (o *PartnerCreateParams) HasNameKana() bool {
	if o != nil && o.NameKana != nil {
		return true
	}

	return false
}

// SetNameKana gets a reference to the given string and assigns it to the NameKana field.
func (o *PartnerCreateParams) SetNameKana(v string) {
	o.NameKana = &v
}

// GetOrgCode returns the OrgCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerCreateParams) GetOrgCode() int32 {
	if o == nil || o.OrgCode.Get() == nil {
		var ret int32
		return ret
	}
	return *o.OrgCode.Get()
}

// GetOrgCodeOk returns a tuple with the OrgCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerCreateParams) GetOrgCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrgCode.Get(), o.OrgCode.IsSet()
}

// HasOrgCode returns a boolean if a field has been set.
func (o *PartnerCreateParams) HasOrgCode() bool {
	if o != nil && o.OrgCode.IsSet() {
		return true
	}

	return false
}

// SetOrgCode gets a reference to the given NullableInt32 and assigns it to the OrgCode field.
func (o *PartnerCreateParams) SetOrgCode(v int32) {
	o.OrgCode.Set(&v)
}
// SetOrgCodeNil sets the value for OrgCode to be an explicit nil
func (o *PartnerCreateParams) SetOrgCodeNil() {
	o.OrgCode.Set(nil)
}

// UnsetOrgCode ensures that no value is present for OrgCode, not even an explicit nil
func (o *PartnerCreateParams) UnsetOrgCode() {
	o.OrgCode.Unset()
}

// GetPartnerBankAccountAttributes returns the PartnerBankAccountAttributes field value if set, zero value otherwise.
func (o *PartnerCreateParams) GetPartnerBankAccountAttributes() PartnerCreateParamsPartnerBankAccountAttributes {
	if o == nil || o.PartnerBankAccountAttributes == nil {
		var ret PartnerCreateParamsPartnerBankAccountAttributes
		return ret
	}
	return *o.PartnerBankAccountAttributes
}

// GetPartnerBankAccountAttributesOk returns a tuple with the PartnerBankAccountAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCreateParams) GetPartnerBankAccountAttributesOk() (*PartnerCreateParamsPartnerBankAccountAttributes, bool) {
	if o == nil || o.PartnerBankAccountAttributes == nil {
		return nil, false
	}
	return o.PartnerBankAccountAttributes, true
}

// HasPartnerBankAccountAttributes returns a boolean if a field has been set.
func (o *PartnerCreateParams) HasPartnerBankAccountAttributes() bool {
	if o != nil && o.PartnerBankAccountAttributes != nil {
		return true
	}

	return false
}

// SetPartnerBankAccountAttributes gets a reference to the given PartnerCreateParamsPartnerBankAccountAttributes and assigns it to the PartnerBankAccountAttributes field.
func (o *PartnerCreateParams) SetPartnerBankAccountAttributes(v PartnerCreateParamsPartnerBankAccountAttributes) {
	o.PartnerBankAccountAttributes = &v
}

// GetPartnerDocSettingAttributes returns the PartnerDocSettingAttributes field value if set, zero value otherwise.
func (o *PartnerCreateParams) GetPartnerDocSettingAttributes() PartnerCreateParamsPartnerDocSettingAttributes {
	if o == nil || o.PartnerDocSettingAttributes == nil {
		var ret PartnerCreateParamsPartnerDocSettingAttributes
		return ret
	}
	return *o.PartnerDocSettingAttributes
}

// GetPartnerDocSettingAttributesOk returns a tuple with the PartnerDocSettingAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCreateParams) GetPartnerDocSettingAttributesOk() (*PartnerCreateParamsPartnerDocSettingAttributes, bool) {
	if o == nil || o.PartnerDocSettingAttributes == nil {
		return nil, false
	}
	return o.PartnerDocSettingAttributes, true
}

// HasPartnerDocSettingAttributes returns a boolean if a field has been set.
func (o *PartnerCreateParams) HasPartnerDocSettingAttributes() bool {
	if o != nil && o.PartnerDocSettingAttributes != nil {
		return true
	}

	return false
}

// SetPartnerDocSettingAttributes gets a reference to the given PartnerCreateParamsPartnerDocSettingAttributes and assigns it to the PartnerDocSettingAttributes field.
func (o *PartnerCreateParams) SetPartnerDocSettingAttributes(v PartnerCreateParamsPartnerDocSettingAttributes) {
	o.PartnerDocSettingAttributes = &v
}

// GetPayerWalletableId returns the PayerWalletableId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerCreateParams) GetPayerWalletableId() int32 {
	if o == nil || o.PayerWalletableId.Get() == nil {
		var ret int32
		return ret
	}
	return *o.PayerWalletableId.Get()
}

// GetPayerWalletableIdOk returns a tuple with the PayerWalletableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerCreateParams) GetPayerWalletableIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PayerWalletableId.Get(), o.PayerWalletableId.IsSet()
}

// HasPayerWalletableId returns a boolean if a field has been set.
func (o *PartnerCreateParams) HasPayerWalletableId() bool {
	if o != nil && o.PayerWalletableId.IsSet() {
		return true
	}

	return false
}

// SetPayerWalletableId gets a reference to the given NullableInt32 and assigns it to the PayerWalletableId field.
func (o *PartnerCreateParams) SetPayerWalletableId(v int32) {
	o.PayerWalletableId.Set(&v)
}
// SetPayerWalletableIdNil sets the value for PayerWalletableId to be an explicit nil
func (o *PartnerCreateParams) SetPayerWalletableIdNil() {
	o.PayerWalletableId.Set(nil)
}

// UnsetPayerWalletableId ensures that no value is present for PayerWalletableId, not even an explicit nil
func (o *PartnerCreateParams) UnsetPayerWalletableId() {
	o.PayerWalletableId.Unset()
}

// GetPaymentTermAttributes returns the PaymentTermAttributes field value if set, zero value otherwise.
func (o *PartnerCreateParams) GetPaymentTermAttributes() PartnerCreateParamsPaymentTermAttributes {
	if o == nil || o.PaymentTermAttributes == nil {
		var ret PartnerCreateParamsPaymentTermAttributes
		return ret
	}
	return *o.PaymentTermAttributes
}

// GetPaymentTermAttributesOk returns a tuple with the PaymentTermAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCreateParams) GetPaymentTermAttributesOk() (*PartnerCreateParamsPaymentTermAttributes, bool) {
	if o == nil || o.PaymentTermAttributes == nil {
		return nil, false
	}
	return o.PaymentTermAttributes, true
}

// HasPaymentTermAttributes returns a boolean if a field has been set.
func (o *PartnerCreateParams) HasPaymentTermAttributes() bool {
	if o != nil && o.PaymentTermAttributes != nil {
		return true
	}

	return false
}

// SetPaymentTermAttributes gets a reference to the given PartnerCreateParamsPaymentTermAttributes and assigns it to the PaymentTermAttributes field.
func (o *PartnerCreateParams) SetPaymentTermAttributes(v PartnerCreateParamsPaymentTermAttributes) {
	o.PaymentTermAttributes = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *PartnerCreateParams) GetPhone() string {
	if o == nil || o.Phone == nil {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCreateParams) GetPhoneOk() (*string, bool) {
	if o == nil || o.Phone == nil {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *PartnerCreateParams) HasPhone() bool {
	if o != nil && o.Phone != nil {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *PartnerCreateParams) SetPhone(v string) {
	o.Phone = &v
}

// GetShortcut1 returns the Shortcut1 field value if set, zero value otherwise.
func (o *PartnerCreateParams) GetShortcut1() string {
	if o == nil || o.Shortcut1 == nil {
		var ret string
		return ret
	}
	return *o.Shortcut1
}

// GetShortcut1Ok returns a tuple with the Shortcut1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCreateParams) GetShortcut1Ok() (*string, bool) {
	if o == nil || o.Shortcut1 == nil {
		return nil, false
	}
	return o.Shortcut1, true
}

// HasShortcut1 returns a boolean if a field has been set.
func (o *PartnerCreateParams) HasShortcut1() bool {
	if o != nil && o.Shortcut1 != nil {
		return true
	}

	return false
}

// SetShortcut1 gets a reference to the given string and assigns it to the Shortcut1 field.
func (o *PartnerCreateParams) SetShortcut1(v string) {
	o.Shortcut1 = &v
}

// GetShortcut2 returns the Shortcut2 field value if set, zero value otherwise.
func (o *PartnerCreateParams) GetShortcut2() string {
	if o == nil || o.Shortcut2 == nil {
		var ret string
		return ret
	}
	return *o.Shortcut2
}

// GetShortcut2Ok returns a tuple with the Shortcut2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCreateParams) GetShortcut2Ok() (*string, bool) {
	if o == nil || o.Shortcut2 == nil {
		return nil, false
	}
	return o.Shortcut2, true
}

// HasShortcut2 returns a boolean if a field has been set.
func (o *PartnerCreateParams) HasShortcut2() bool {
	if o != nil && o.Shortcut2 != nil {
		return true
	}

	return false
}

// SetShortcut2 gets a reference to the given string and assigns it to the Shortcut2 field.
func (o *PartnerCreateParams) SetShortcut2(v string) {
	o.Shortcut2 = &v
}

// GetTransferFeeHandlingSide returns the TransferFeeHandlingSide field value if set, zero value otherwise.
func (o *PartnerCreateParams) GetTransferFeeHandlingSide() string {
	if o == nil || o.TransferFeeHandlingSide == nil {
		var ret string
		return ret
	}
	return *o.TransferFeeHandlingSide
}

// GetTransferFeeHandlingSideOk returns a tuple with the TransferFeeHandlingSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCreateParams) GetTransferFeeHandlingSideOk() (*string, bool) {
	if o == nil || o.TransferFeeHandlingSide == nil {
		return nil, false
	}
	return o.TransferFeeHandlingSide, true
}

// HasTransferFeeHandlingSide returns a boolean if a field has been set.
func (o *PartnerCreateParams) HasTransferFeeHandlingSide() bool {
	if o != nil && o.TransferFeeHandlingSide != nil {
		return true
	}

	return false
}

// SetTransferFeeHandlingSide gets a reference to the given string and assigns it to the TransferFeeHandlingSide field.
func (o *PartnerCreateParams) SetTransferFeeHandlingSide(v string) {
	o.TransferFeeHandlingSide = &v
}

func (o PartnerCreateParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddressAttributes != nil {
		toSerialize["address_attributes"] = o.AddressAttributes
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["company_id"] = o.CompanyId
	}
	if o.ContactName != nil {
		toSerialize["contact_name"] = o.ContactName
	}
	if o.CountryCode != nil {
		toSerialize["country_code"] = o.CountryCode
	}
	if o.DefaultTitle != nil {
		toSerialize["default_title"] = o.DefaultTitle
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.InvoicePaymentTermAttributes != nil {
		toSerialize["invoice_payment_term_attributes"] = o.InvoicePaymentTermAttributes
	}
	if o.LongName != nil {
		toSerialize["long_name"] = o.LongName
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NameKana != nil {
		toSerialize["name_kana"] = o.NameKana
	}
	if o.OrgCode.IsSet() {
		toSerialize["org_code"] = o.OrgCode.Get()
	}
	if o.PartnerBankAccountAttributes != nil {
		toSerialize["partner_bank_account_attributes"] = o.PartnerBankAccountAttributes
	}
	if o.PartnerDocSettingAttributes != nil {
		toSerialize["partner_doc_setting_attributes"] = o.PartnerDocSettingAttributes
	}
	if o.PayerWalletableId.IsSet() {
		toSerialize["payer_walletable_id"] = o.PayerWalletableId.Get()
	}
	if o.PaymentTermAttributes != nil {
		toSerialize["payment_term_attributes"] = o.PaymentTermAttributes
	}
	if o.Phone != nil {
		toSerialize["phone"] = o.Phone
	}
	if o.Shortcut1 != nil {
		toSerialize["shortcut1"] = o.Shortcut1
	}
	if o.Shortcut2 != nil {
		toSerialize["shortcut2"] = o.Shortcut2
	}
	if o.TransferFeeHandlingSide != nil {
		toSerialize["transfer_fee_handling_side"] = o.TransferFeeHandlingSide
	}
	return json.Marshal(toSerialize)
}

type NullablePartnerCreateParams struct {
	value *PartnerCreateParams
	isSet bool
}

func (v NullablePartnerCreateParams) Get() *PartnerCreateParams {
	return v.value
}

func (v *NullablePartnerCreateParams) Set(val *PartnerCreateParams) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerCreateParams) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerCreateParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerCreateParams(val *PartnerCreateParams) *NullablePartnerCreateParams {
	return &NullablePartnerCreateParams{value: val, isSet: true}
}

func (v NullablePartnerCreateParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerCreateParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


