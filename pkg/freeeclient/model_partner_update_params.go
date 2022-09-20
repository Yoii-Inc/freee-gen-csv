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

// PartnerUpdateParams struct for PartnerUpdateParams
type PartnerUpdateParams struct {
	AddressAttributes *PartnerCreateParamsAddressAttributes `json:"address_attributes,omitempty"`
	// 事業所ID
	CompanyId int64 `json:"company_id"`
	// 担当者 氏名 (255文字以内)
	ContactName *string `json:"contact_name,omitempty"`
	// 地域（JP: 国内、ZZ:国外）、指定しない場合JPになります。
	CountryCode *string `json:"country_code,omitempty"`
	// 敬称（御中、様、(空白)の3つから選択）
	DefaultTitle *string `json:"default_title,omitempty"`
	// 担当者 メールアドレス (255文字以内)
	Email *string `json:"email,omitempty"`
	InvoicePaymentTermAttributes NullablePartnerUpdateParamsInvoicePaymentTermAttributes `json:"invoice_payment_term_attributes,omitempty"`
	// 正式名称（255文字以内）
	LongName *string `json:"long_name,omitempty"`
	// 取引先名 (255文字以内)
	Name string `json:"name"`
	// カナ名称（255文字以内）
	NameKana *string `json:"name_kana,omitempty"`
	// 事業所種別（null: 未設定、1: 法人、2: 個人）
	OrgCode NullableInt64 `json:"org_code,omitempty"`
	PartnerBankAccountAttributes *PartnerCreateParamsPartnerBankAccountAttributes `json:"partner_bank_account_attributes,omitempty"`
	PartnerDocSettingAttributes *PartnerCreateParamsPartnerDocSettingAttributes `json:"partner_doc_setting_attributes,omitempty"`
	// 振込元口座ID（一括振込ファイル用）:（walletableのtypeが'bank_account'のidのみ指定できます。また、未設定にする場合は、nullを指定してください。）
	PayerWalletableId NullableInt64 `json:"payer_walletable_id,omitempty"`
	PaymentTermAttributes NullablePartnerUpdateParamsPaymentTermAttributes `json:"payment_term_attributes,omitempty"`
	// 電話番号
	Phone *string `json:"phone,omitempty"`
	// ショートカット１ (255文字以内)
	Shortcut1 *string `json:"shortcut1,omitempty"`
	// ショートカット２ (255文字以内)
	Shortcut2 *string `json:"shortcut2,omitempty"`
	// 振込手数料負担（一括振込ファイル用）: (振込元(当方): payer, 振込先(先方): payee)、指定しない場合payerになります。
	TransferFeeHandlingSide *string `json:"transfer_fee_handling_side,omitempty"`
}

// NewPartnerUpdateParams instantiates a new PartnerUpdateParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerUpdateParams(companyId int64, name string) *PartnerUpdateParams {
	this := PartnerUpdateParams{}
	this.CompanyId = companyId
	this.Name = name
	return &this
}

// NewPartnerUpdateParamsWithDefaults instantiates a new PartnerUpdateParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerUpdateParamsWithDefaults() *PartnerUpdateParams {
	this := PartnerUpdateParams{}
	return &this
}

// GetAddressAttributes returns the AddressAttributes field value if set, zero value otherwise.
func (o *PartnerUpdateParams) GetAddressAttributes() PartnerCreateParamsAddressAttributes {
	if o == nil || o.AddressAttributes == nil {
		var ret PartnerCreateParamsAddressAttributes
		return ret
	}
	return *o.AddressAttributes
}

// GetAddressAttributesOk returns a tuple with the AddressAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerUpdateParams) GetAddressAttributesOk() (*PartnerCreateParamsAddressAttributes, bool) {
	if o == nil || o.AddressAttributes == nil {
		return nil, false
	}
	return o.AddressAttributes, true
}

// HasAddressAttributes returns a boolean if a field has been set.
func (o *PartnerUpdateParams) HasAddressAttributes() bool {
	if o != nil && o.AddressAttributes != nil {
		return true
	}

	return false
}

// SetAddressAttributes gets a reference to the given PartnerCreateParamsAddressAttributes and assigns it to the AddressAttributes field.
func (o *PartnerUpdateParams) SetAddressAttributes(v PartnerCreateParamsAddressAttributes) {
	o.AddressAttributes = &v
}

// GetCompanyId returns the CompanyId field value
func (o *PartnerUpdateParams) GetCompanyId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value
// and a boolean to check if the value has been set.
func (o *PartnerUpdateParams) GetCompanyIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyId, true
}

// SetCompanyId sets field value
func (o *PartnerUpdateParams) SetCompanyId(v int64) {
	o.CompanyId = v
}

// GetContactName returns the ContactName field value if set, zero value otherwise.
func (o *PartnerUpdateParams) GetContactName() string {
	if o == nil || o.ContactName == nil {
		var ret string
		return ret
	}
	return *o.ContactName
}

// GetContactNameOk returns a tuple with the ContactName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerUpdateParams) GetContactNameOk() (*string, bool) {
	if o == nil || o.ContactName == nil {
		return nil, false
	}
	return o.ContactName, true
}

// HasContactName returns a boolean if a field has been set.
func (o *PartnerUpdateParams) HasContactName() bool {
	if o != nil && o.ContactName != nil {
		return true
	}

	return false
}

// SetContactName gets a reference to the given string and assigns it to the ContactName field.
func (o *PartnerUpdateParams) SetContactName(v string) {
	o.ContactName = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *PartnerUpdateParams) GetCountryCode() string {
	if o == nil || o.CountryCode == nil {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerUpdateParams) GetCountryCodeOk() (*string, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *PartnerUpdateParams) HasCountryCode() bool {
	if o != nil && o.CountryCode != nil {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *PartnerUpdateParams) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetDefaultTitle returns the DefaultTitle field value if set, zero value otherwise.
func (o *PartnerUpdateParams) GetDefaultTitle() string {
	if o == nil || o.DefaultTitle == nil {
		var ret string
		return ret
	}
	return *o.DefaultTitle
}

// GetDefaultTitleOk returns a tuple with the DefaultTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerUpdateParams) GetDefaultTitleOk() (*string, bool) {
	if o == nil || o.DefaultTitle == nil {
		return nil, false
	}
	return o.DefaultTitle, true
}

// HasDefaultTitle returns a boolean if a field has been set.
func (o *PartnerUpdateParams) HasDefaultTitle() bool {
	if o != nil && o.DefaultTitle != nil {
		return true
	}

	return false
}

// SetDefaultTitle gets a reference to the given string and assigns it to the DefaultTitle field.
func (o *PartnerUpdateParams) SetDefaultTitle(v string) {
	o.DefaultTitle = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PartnerUpdateParams) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerUpdateParams) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PartnerUpdateParams) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PartnerUpdateParams) SetEmail(v string) {
	o.Email = &v
}

// GetInvoicePaymentTermAttributes returns the InvoicePaymentTermAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerUpdateParams) GetInvoicePaymentTermAttributes() PartnerUpdateParamsInvoicePaymentTermAttributes {
	if o == nil || o.InvoicePaymentTermAttributes.Get() == nil {
		var ret PartnerUpdateParamsInvoicePaymentTermAttributes
		return ret
	}
	return *o.InvoicePaymentTermAttributes.Get()
}

// GetInvoicePaymentTermAttributesOk returns a tuple with the InvoicePaymentTermAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerUpdateParams) GetInvoicePaymentTermAttributesOk() (*PartnerUpdateParamsInvoicePaymentTermAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoicePaymentTermAttributes.Get(), o.InvoicePaymentTermAttributes.IsSet()
}

// HasInvoicePaymentTermAttributes returns a boolean if a field has been set.
func (o *PartnerUpdateParams) HasInvoicePaymentTermAttributes() bool {
	if o != nil && o.InvoicePaymentTermAttributes.IsSet() {
		return true
	}

	return false
}

// SetInvoicePaymentTermAttributes gets a reference to the given NullablePartnerUpdateParamsInvoicePaymentTermAttributes and assigns it to the InvoicePaymentTermAttributes field.
func (o *PartnerUpdateParams) SetInvoicePaymentTermAttributes(v PartnerUpdateParamsInvoicePaymentTermAttributes) {
	o.InvoicePaymentTermAttributes.Set(&v)
}
// SetInvoicePaymentTermAttributesNil sets the value for InvoicePaymentTermAttributes to be an explicit nil
func (o *PartnerUpdateParams) SetInvoicePaymentTermAttributesNil() {
	o.InvoicePaymentTermAttributes.Set(nil)
}

// UnsetInvoicePaymentTermAttributes ensures that no value is present for InvoicePaymentTermAttributes, not even an explicit nil
func (o *PartnerUpdateParams) UnsetInvoicePaymentTermAttributes() {
	o.InvoicePaymentTermAttributes.Unset()
}

// GetLongName returns the LongName field value if set, zero value otherwise.
func (o *PartnerUpdateParams) GetLongName() string {
	if o == nil || o.LongName == nil {
		var ret string
		return ret
	}
	return *o.LongName
}

// GetLongNameOk returns a tuple with the LongName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerUpdateParams) GetLongNameOk() (*string, bool) {
	if o == nil || o.LongName == nil {
		return nil, false
	}
	return o.LongName, true
}

// HasLongName returns a boolean if a field has been set.
func (o *PartnerUpdateParams) HasLongName() bool {
	if o != nil && o.LongName != nil {
		return true
	}

	return false
}

// SetLongName gets a reference to the given string and assigns it to the LongName field.
func (o *PartnerUpdateParams) SetLongName(v string) {
	o.LongName = &v
}

// GetName returns the Name field value
func (o *PartnerUpdateParams) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PartnerUpdateParams) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PartnerUpdateParams) SetName(v string) {
	o.Name = v
}

// GetNameKana returns the NameKana field value if set, zero value otherwise.
func (o *PartnerUpdateParams) GetNameKana() string {
	if o == nil || o.NameKana == nil {
		var ret string
		return ret
	}
	return *o.NameKana
}

// GetNameKanaOk returns a tuple with the NameKana field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerUpdateParams) GetNameKanaOk() (*string, bool) {
	if o == nil || o.NameKana == nil {
		return nil, false
	}
	return o.NameKana, true
}

// HasNameKana returns a boolean if a field has been set.
func (o *PartnerUpdateParams) HasNameKana() bool {
	if o != nil && o.NameKana != nil {
		return true
	}

	return false
}

// SetNameKana gets a reference to the given string and assigns it to the NameKana field.
func (o *PartnerUpdateParams) SetNameKana(v string) {
	o.NameKana = &v
}

// GetOrgCode returns the OrgCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerUpdateParams) GetOrgCode() int64 {
	if o == nil || o.OrgCode.Get() == nil {
		var ret int64
		return ret
	}
	return *o.OrgCode.Get()
}

// GetOrgCodeOk returns a tuple with the OrgCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerUpdateParams) GetOrgCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrgCode.Get(), o.OrgCode.IsSet()
}

// HasOrgCode returns a boolean if a field has been set.
func (o *PartnerUpdateParams) HasOrgCode() bool {
	if o != nil && o.OrgCode.IsSet() {
		return true
	}

	return false
}

// SetOrgCode gets a reference to the given NullableInt64 and assigns it to the OrgCode field.
func (o *PartnerUpdateParams) SetOrgCode(v int64) {
	o.OrgCode.Set(&v)
}
// SetOrgCodeNil sets the value for OrgCode to be an explicit nil
func (o *PartnerUpdateParams) SetOrgCodeNil() {
	o.OrgCode.Set(nil)
}

// UnsetOrgCode ensures that no value is present for OrgCode, not even an explicit nil
func (o *PartnerUpdateParams) UnsetOrgCode() {
	o.OrgCode.Unset()
}

// GetPartnerBankAccountAttributes returns the PartnerBankAccountAttributes field value if set, zero value otherwise.
func (o *PartnerUpdateParams) GetPartnerBankAccountAttributes() PartnerCreateParamsPartnerBankAccountAttributes {
	if o == nil || o.PartnerBankAccountAttributes == nil {
		var ret PartnerCreateParamsPartnerBankAccountAttributes
		return ret
	}
	return *o.PartnerBankAccountAttributes
}

// GetPartnerBankAccountAttributesOk returns a tuple with the PartnerBankAccountAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerUpdateParams) GetPartnerBankAccountAttributesOk() (*PartnerCreateParamsPartnerBankAccountAttributes, bool) {
	if o == nil || o.PartnerBankAccountAttributes == nil {
		return nil, false
	}
	return o.PartnerBankAccountAttributes, true
}

// HasPartnerBankAccountAttributes returns a boolean if a field has been set.
func (o *PartnerUpdateParams) HasPartnerBankAccountAttributes() bool {
	if o != nil && o.PartnerBankAccountAttributes != nil {
		return true
	}

	return false
}

// SetPartnerBankAccountAttributes gets a reference to the given PartnerCreateParamsPartnerBankAccountAttributes and assigns it to the PartnerBankAccountAttributes field.
func (o *PartnerUpdateParams) SetPartnerBankAccountAttributes(v PartnerCreateParamsPartnerBankAccountAttributes) {
	o.PartnerBankAccountAttributes = &v
}

// GetPartnerDocSettingAttributes returns the PartnerDocSettingAttributes field value if set, zero value otherwise.
func (o *PartnerUpdateParams) GetPartnerDocSettingAttributes() PartnerCreateParamsPartnerDocSettingAttributes {
	if o == nil || o.PartnerDocSettingAttributes == nil {
		var ret PartnerCreateParamsPartnerDocSettingAttributes
		return ret
	}
	return *o.PartnerDocSettingAttributes
}

// GetPartnerDocSettingAttributesOk returns a tuple with the PartnerDocSettingAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerUpdateParams) GetPartnerDocSettingAttributesOk() (*PartnerCreateParamsPartnerDocSettingAttributes, bool) {
	if o == nil || o.PartnerDocSettingAttributes == nil {
		return nil, false
	}
	return o.PartnerDocSettingAttributes, true
}

// HasPartnerDocSettingAttributes returns a boolean if a field has been set.
func (o *PartnerUpdateParams) HasPartnerDocSettingAttributes() bool {
	if o != nil && o.PartnerDocSettingAttributes != nil {
		return true
	}

	return false
}

// SetPartnerDocSettingAttributes gets a reference to the given PartnerCreateParamsPartnerDocSettingAttributes and assigns it to the PartnerDocSettingAttributes field.
func (o *PartnerUpdateParams) SetPartnerDocSettingAttributes(v PartnerCreateParamsPartnerDocSettingAttributes) {
	o.PartnerDocSettingAttributes = &v
}

// GetPayerWalletableId returns the PayerWalletableId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerUpdateParams) GetPayerWalletableId() int64 {
	if o == nil || o.PayerWalletableId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.PayerWalletableId.Get()
}

// GetPayerWalletableIdOk returns a tuple with the PayerWalletableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerUpdateParams) GetPayerWalletableIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PayerWalletableId.Get(), o.PayerWalletableId.IsSet()
}

// HasPayerWalletableId returns a boolean if a field has been set.
func (o *PartnerUpdateParams) HasPayerWalletableId() bool {
	if o != nil && o.PayerWalletableId.IsSet() {
		return true
	}

	return false
}

// SetPayerWalletableId gets a reference to the given NullableInt64 and assigns it to the PayerWalletableId field.
func (o *PartnerUpdateParams) SetPayerWalletableId(v int64) {
	o.PayerWalletableId.Set(&v)
}
// SetPayerWalletableIdNil sets the value for PayerWalletableId to be an explicit nil
func (o *PartnerUpdateParams) SetPayerWalletableIdNil() {
	o.PayerWalletableId.Set(nil)
}

// UnsetPayerWalletableId ensures that no value is present for PayerWalletableId, not even an explicit nil
func (o *PartnerUpdateParams) UnsetPayerWalletableId() {
	o.PayerWalletableId.Unset()
}

// GetPaymentTermAttributes returns the PaymentTermAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerUpdateParams) GetPaymentTermAttributes() PartnerUpdateParamsPaymentTermAttributes {
	if o == nil || o.PaymentTermAttributes.Get() == nil {
		var ret PartnerUpdateParamsPaymentTermAttributes
		return ret
	}
	return *o.PaymentTermAttributes.Get()
}

// GetPaymentTermAttributesOk returns a tuple with the PaymentTermAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerUpdateParams) GetPaymentTermAttributesOk() (*PartnerUpdateParamsPaymentTermAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentTermAttributes.Get(), o.PaymentTermAttributes.IsSet()
}

// HasPaymentTermAttributes returns a boolean if a field has been set.
func (o *PartnerUpdateParams) HasPaymentTermAttributes() bool {
	if o != nil && o.PaymentTermAttributes.IsSet() {
		return true
	}

	return false
}

// SetPaymentTermAttributes gets a reference to the given NullablePartnerUpdateParamsPaymentTermAttributes and assigns it to the PaymentTermAttributes field.
func (o *PartnerUpdateParams) SetPaymentTermAttributes(v PartnerUpdateParamsPaymentTermAttributes) {
	o.PaymentTermAttributes.Set(&v)
}
// SetPaymentTermAttributesNil sets the value for PaymentTermAttributes to be an explicit nil
func (o *PartnerUpdateParams) SetPaymentTermAttributesNil() {
	o.PaymentTermAttributes.Set(nil)
}

// UnsetPaymentTermAttributes ensures that no value is present for PaymentTermAttributes, not even an explicit nil
func (o *PartnerUpdateParams) UnsetPaymentTermAttributes() {
	o.PaymentTermAttributes.Unset()
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *PartnerUpdateParams) GetPhone() string {
	if o == nil || o.Phone == nil {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerUpdateParams) GetPhoneOk() (*string, bool) {
	if o == nil || o.Phone == nil {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *PartnerUpdateParams) HasPhone() bool {
	if o != nil && o.Phone != nil {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *PartnerUpdateParams) SetPhone(v string) {
	o.Phone = &v
}

// GetShortcut1 returns the Shortcut1 field value if set, zero value otherwise.
func (o *PartnerUpdateParams) GetShortcut1() string {
	if o == nil || o.Shortcut1 == nil {
		var ret string
		return ret
	}
	return *o.Shortcut1
}

// GetShortcut1Ok returns a tuple with the Shortcut1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerUpdateParams) GetShortcut1Ok() (*string, bool) {
	if o == nil || o.Shortcut1 == nil {
		return nil, false
	}
	return o.Shortcut1, true
}

// HasShortcut1 returns a boolean if a field has been set.
func (o *PartnerUpdateParams) HasShortcut1() bool {
	if o != nil && o.Shortcut1 != nil {
		return true
	}

	return false
}

// SetShortcut1 gets a reference to the given string and assigns it to the Shortcut1 field.
func (o *PartnerUpdateParams) SetShortcut1(v string) {
	o.Shortcut1 = &v
}

// GetShortcut2 returns the Shortcut2 field value if set, zero value otherwise.
func (o *PartnerUpdateParams) GetShortcut2() string {
	if o == nil || o.Shortcut2 == nil {
		var ret string
		return ret
	}
	return *o.Shortcut2
}

// GetShortcut2Ok returns a tuple with the Shortcut2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerUpdateParams) GetShortcut2Ok() (*string, bool) {
	if o == nil || o.Shortcut2 == nil {
		return nil, false
	}
	return o.Shortcut2, true
}

// HasShortcut2 returns a boolean if a field has been set.
func (o *PartnerUpdateParams) HasShortcut2() bool {
	if o != nil && o.Shortcut2 != nil {
		return true
	}

	return false
}

// SetShortcut2 gets a reference to the given string and assigns it to the Shortcut2 field.
func (o *PartnerUpdateParams) SetShortcut2(v string) {
	o.Shortcut2 = &v
}

// GetTransferFeeHandlingSide returns the TransferFeeHandlingSide field value if set, zero value otherwise.
func (o *PartnerUpdateParams) GetTransferFeeHandlingSide() string {
	if o == nil || o.TransferFeeHandlingSide == nil {
		var ret string
		return ret
	}
	return *o.TransferFeeHandlingSide
}

// GetTransferFeeHandlingSideOk returns a tuple with the TransferFeeHandlingSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerUpdateParams) GetTransferFeeHandlingSideOk() (*string, bool) {
	if o == nil || o.TransferFeeHandlingSide == nil {
		return nil, false
	}
	return o.TransferFeeHandlingSide, true
}

// HasTransferFeeHandlingSide returns a boolean if a field has been set.
func (o *PartnerUpdateParams) HasTransferFeeHandlingSide() bool {
	if o != nil && o.TransferFeeHandlingSide != nil {
		return true
	}

	return false
}

// SetTransferFeeHandlingSide gets a reference to the given string and assigns it to the TransferFeeHandlingSide field.
func (o *PartnerUpdateParams) SetTransferFeeHandlingSide(v string) {
	o.TransferFeeHandlingSide = &v
}

func (o PartnerUpdateParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddressAttributes != nil {
		toSerialize["address_attributes"] = o.AddressAttributes
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
	if o.InvoicePaymentTermAttributes.IsSet() {
		toSerialize["invoice_payment_term_attributes"] = o.InvoicePaymentTermAttributes.Get()
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
	if o.PaymentTermAttributes.IsSet() {
		toSerialize["payment_term_attributes"] = o.PaymentTermAttributes.Get()
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

type NullablePartnerUpdateParams struct {
	value *PartnerUpdateParams
	isSet bool
}

func (v NullablePartnerUpdateParams) Get() *PartnerUpdateParams {
	return v.value
}

func (v *NullablePartnerUpdateParams) Set(val *PartnerUpdateParams) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerUpdateParams) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerUpdateParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerUpdateParams(val *PartnerUpdateParams) *NullablePartnerUpdateParams {
	return &NullablePartnerUpdateParams{value: val, isSet: true}
}

func (v NullablePartnerUpdateParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerUpdateParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


