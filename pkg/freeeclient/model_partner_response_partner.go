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

// PartnerResponsePartner struct for PartnerResponsePartner
type PartnerResponsePartner struct {
	AddressAttributes *PartnerResponsePartnerAddressAttributes `json:"address_attributes,omitempty"`
	// 取引先の使用設定（true: 使用する、false: 使用しない） <br> <ul>   <li>     本APIでpartnerを作成した場合はtrueになります。   </li>   <li>     falseにする場合はWeb画面から変更できます。   </li>   <li>     trueの場合、Web画面での取引登録時などに入力候補として表示されます。   </li>   <li>     falseの場合、取引先自体は削除せず、Web画面での取引登録時などに入力候補として表示されません。ただし取引（収入／支出）の作成APIなどでfalseの取引先をパラメータに指定すれば、取引などにfalseの取引先を設定できます。   </li> </ul>
	Available bool `json:"available"`
	// 取引先コード
	Code NullableString `json:"code"`
	// 事業所ID
	CompanyId int64 `json:"company_id"`
	// 担当者 氏名
	ContactName NullableString `json:"contact_name,omitempty"`
	// 地域（JP: 国内、ZZ:国外）
	CountryCode *string `json:"country_code,omitempty"`
	// 敬称（御中、様、(空白)の3つから選択）
	DefaultTitle NullableString `json:"default_title,omitempty"`
	// 担当者 メールアドレス
	Email NullableString `json:"email,omitempty"`
	// 取引先ID
	Id int64 `json:"id"`
	InvoicePaymentTermAttributes *PartnerResponsePartnerInvoicePaymentTermAttributes `json:"invoice_payment_term_attributes,omitempty"`
	// 正式名称（255文字以内）
	LongName NullableString `json:"long_name,omitempty"`
	// 取引先名
	Name string `json:"name"`
	// カナ名称（255文字以内）
	NameKana NullableString `json:"name_kana,omitempty"`
	// 事業所種別（null: 未設定、1: 法人、2: 個人）
	OrgCode NullableInt64 `json:"org_code,omitempty"`
	PartnerBankAccountAttributes *PartnerResponsePartnerPartnerBankAccountAttributes `json:"partner_bank_account_attributes,omitempty"`
	PartnerDocSettingAttributes *PartnerCreateParamsPartnerDocSettingAttributes `json:"partner_doc_setting_attributes,omitempty"`
	// 振込元口座ID（一括振込ファイル用）:（未設定の場合は、nullです。）
	PayerWalletableId NullableInt64 `json:"payer_walletable_id,omitempty"`
	PaymentTermAttributes *PartnerResponsePartnerPaymentTermAttributes `json:"payment_term_attributes,omitempty"`
	// 電話番号
	Phone NullableString `json:"phone,omitempty"`
	// ショートカット1 (255文字以内)
	Shortcut1 NullableString `json:"shortcut1,omitempty"`
	// ショートカット2 (255文字以内)
	Shortcut2 NullableString `json:"shortcut2,omitempty"`
	// 振込手数料負担（一括振込ファイル用）: (振込元(当方): payer, 振込先(先方): payee)
	TransferFeeHandlingSide *string `json:"transfer_fee_handling_side,omitempty"`
	// 更新日 (yyyy-mm-dd)
	UpdateDate string `json:"update_date"`
}

// NewPartnerResponsePartner instantiates a new PartnerResponsePartner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerResponsePartner(available bool, code NullableString, companyId int64, id int64, name string, updateDate string) *PartnerResponsePartner {
	this := PartnerResponsePartner{}
	this.Available = available
	this.Code = code
	this.CompanyId = companyId
	this.Id = id
	this.Name = name
	this.UpdateDate = updateDate
	return &this
}

// NewPartnerResponsePartnerWithDefaults instantiates a new PartnerResponsePartner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerResponsePartnerWithDefaults() *PartnerResponsePartner {
	this := PartnerResponsePartner{}
	return &this
}

// GetAddressAttributes returns the AddressAttributes field value if set, zero value otherwise.
func (o *PartnerResponsePartner) GetAddressAttributes() PartnerResponsePartnerAddressAttributes {
	if o == nil || o.AddressAttributes == nil {
		var ret PartnerResponsePartnerAddressAttributes
		return ret
	}
	return *o.AddressAttributes
}

// GetAddressAttributesOk returns a tuple with the AddressAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerResponsePartner) GetAddressAttributesOk() (*PartnerResponsePartnerAddressAttributes, bool) {
	if o == nil || o.AddressAttributes == nil {
		return nil, false
	}
	return o.AddressAttributes, true
}

// HasAddressAttributes returns a boolean if a field has been set.
func (o *PartnerResponsePartner) HasAddressAttributes() bool {
	if o != nil && o.AddressAttributes != nil {
		return true
	}

	return false
}

// SetAddressAttributes gets a reference to the given PartnerResponsePartnerAddressAttributes and assigns it to the AddressAttributes field.
func (o *PartnerResponsePartner) SetAddressAttributes(v PartnerResponsePartnerAddressAttributes) {
	o.AddressAttributes = &v
}

// GetAvailable returns the Available field value
func (o *PartnerResponsePartner) GetAvailable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Available
}

// GetAvailableOk returns a tuple with the Available field value
// and a boolean to check if the value has been set.
func (o *PartnerResponsePartner) GetAvailableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Available, true
}

// SetAvailable sets field value
func (o *PartnerResponsePartner) SetAvailable(v bool) {
	o.Available = v
}

// GetCode returns the Code field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PartnerResponsePartner) GetCode() string {
	if o == nil || o.Code.Get() == nil {
		var ret string
		return ret
	}

	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerResponsePartner) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// SetCode sets field value
func (o *PartnerResponsePartner) SetCode(v string) {
	o.Code.Set(&v)
}

// GetCompanyId returns the CompanyId field value
func (o *PartnerResponsePartner) GetCompanyId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value
// and a boolean to check if the value has been set.
func (o *PartnerResponsePartner) GetCompanyIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyId, true
}

// SetCompanyId sets field value
func (o *PartnerResponsePartner) SetCompanyId(v int64) {
	o.CompanyId = v
}

// GetContactName returns the ContactName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerResponsePartner) GetContactName() string {
	if o == nil || o.ContactName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContactName.Get()
}

// GetContactNameOk returns a tuple with the ContactName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerResponsePartner) GetContactNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContactName.Get(), o.ContactName.IsSet()
}

// HasContactName returns a boolean if a field has been set.
func (o *PartnerResponsePartner) HasContactName() bool {
	if o != nil && o.ContactName.IsSet() {
		return true
	}

	return false
}

// SetContactName gets a reference to the given NullableString and assigns it to the ContactName field.
func (o *PartnerResponsePartner) SetContactName(v string) {
	o.ContactName.Set(&v)
}
// SetContactNameNil sets the value for ContactName to be an explicit nil
func (o *PartnerResponsePartner) SetContactNameNil() {
	o.ContactName.Set(nil)
}

// UnsetContactName ensures that no value is present for ContactName, not even an explicit nil
func (o *PartnerResponsePartner) UnsetContactName() {
	o.ContactName.Unset()
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *PartnerResponsePartner) GetCountryCode() string {
	if o == nil || o.CountryCode == nil {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerResponsePartner) GetCountryCodeOk() (*string, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *PartnerResponsePartner) HasCountryCode() bool {
	if o != nil && o.CountryCode != nil {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *PartnerResponsePartner) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetDefaultTitle returns the DefaultTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerResponsePartner) GetDefaultTitle() string {
	if o == nil || o.DefaultTitle.Get() == nil {
		var ret string
		return ret
	}
	return *o.DefaultTitle.Get()
}

// GetDefaultTitleOk returns a tuple with the DefaultTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerResponsePartner) GetDefaultTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultTitle.Get(), o.DefaultTitle.IsSet()
}

// HasDefaultTitle returns a boolean if a field has been set.
func (o *PartnerResponsePartner) HasDefaultTitle() bool {
	if o != nil && o.DefaultTitle.IsSet() {
		return true
	}

	return false
}

// SetDefaultTitle gets a reference to the given NullableString and assigns it to the DefaultTitle field.
func (o *PartnerResponsePartner) SetDefaultTitle(v string) {
	o.DefaultTitle.Set(&v)
}
// SetDefaultTitleNil sets the value for DefaultTitle to be an explicit nil
func (o *PartnerResponsePartner) SetDefaultTitleNil() {
	o.DefaultTitle.Set(nil)
}

// UnsetDefaultTitle ensures that no value is present for DefaultTitle, not even an explicit nil
func (o *PartnerResponsePartner) UnsetDefaultTitle() {
	o.DefaultTitle.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerResponsePartner) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerResponsePartner) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *PartnerResponsePartner) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *PartnerResponsePartner) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *PartnerResponsePartner) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *PartnerResponsePartner) UnsetEmail() {
	o.Email.Unset()
}

// GetId returns the Id field value
func (o *PartnerResponsePartner) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PartnerResponsePartner) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PartnerResponsePartner) SetId(v int64) {
	o.Id = v
}

// GetInvoicePaymentTermAttributes returns the InvoicePaymentTermAttributes field value if set, zero value otherwise.
func (o *PartnerResponsePartner) GetInvoicePaymentTermAttributes() PartnerResponsePartnerInvoicePaymentTermAttributes {
	if o == nil || o.InvoicePaymentTermAttributes == nil {
		var ret PartnerResponsePartnerInvoicePaymentTermAttributes
		return ret
	}
	return *o.InvoicePaymentTermAttributes
}

// GetInvoicePaymentTermAttributesOk returns a tuple with the InvoicePaymentTermAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerResponsePartner) GetInvoicePaymentTermAttributesOk() (*PartnerResponsePartnerInvoicePaymentTermAttributes, bool) {
	if o == nil || o.InvoicePaymentTermAttributes == nil {
		return nil, false
	}
	return o.InvoicePaymentTermAttributes, true
}

// HasInvoicePaymentTermAttributes returns a boolean if a field has been set.
func (o *PartnerResponsePartner) HasInvoicePaymentTermAttributes() bool {
	if o != nil && o.InvoicePaymentTermAttributes != nil {
		return true
	}

	return false
}

// SetInvoicePaymentTermAttributes gets a reference to the given PartnerResponsePartnerInvoicePaymentTermAttributes and assigns it to the InvoicePaymentTermAttributes field.
func (o *PartnerResponsePartner) SetInvoicePaymentTermAttributes(v PartnerResponsePartnerInvoicePaymentTermAttributes) {
	o.InvoicePaymentTermAttributes = &v
}

// GetLongName returns the LongName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerResponsePartner) GetLongName() string {
	if o == nil || o.LongName.Get() == nil {
		var ret string
		return ret
	}
	return *o.LongName.Get()
}

// GetLongNameOk returns a tuple with the LongName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerResponsePartner) GetLongNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LongName.Get(), o.LongName.IsSet()
}

// HasLongName returns a boolean if a field has been set.
func (o *PartnerResponsePartner) HasLongName() bool {
	if o != nil && o.LongName.IsSet() {
		return true
	}

	return false
}

// SetLongName gets a reference to the given NullableString and assigns it to the LongName field.
func (o *PartnerResponsePartner) SetLongName(v string) {
	o.LongName.Set(&v)
}
// SetLongNameNil sets the value for LongName to be an explicit nil
func (o *PartnerResponsePartner) SetLongNameNil() {
	o.LongName.Set(nil)
}

// UnsetLongName ensures that no value is present for LongName, not even an explicit nil
func (o *PartnerResponsePartner) UnsetLongName() {
	o.LongName.Unset()
}

// GetName returns the Name field value
func (o *PartnerResponsePartner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PartnerResponsePartner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PartnerResponsePartner) SetName(v string) {
	o.Name = v
}

// GetNameKana returns the NameKana field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerResponsePartner) GetNameKana() string {
	if o == nil || o.NameKana.Get() == nil {
		var ret string
		return ret
	}
	return *o.NameKana.Get()
}

// GetNameKanaOk returns a tuple with the NameKana field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerResponsePartner) GetNameKanaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NameKana.Get(), o.NameKana.IsSet()
}

// HasNameKana returns a boolean if a field has been set.
func (o *PartnerResponsePartner) HasNameKana() bool {
	if o != nil && o.NameKana.IsSet() {
		return true
	}

	return false
}

// SetNameKana gets a reference to the given NullableString and assigns it to the NameKana field.
func (o *PartnerResponsePartner) SetNameKana(v string) {
	o.NameKana.Set(&v)
}
// SetNameKanaNil sets the value for NameKana to be an explicit nil
func (o *PartnerResponsePartner) SetNameKanaNil() {
	o.NameKana.Set(nil)
}

// UnsetNameKana ensures that no value is present for NameKana, not even an explicit nil
func (o *PartnerResponsePartner) UnsetNameKana() {
	o.NameKana.Unset()
}

// GetOrgCode returns the OrgCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerResponsePartner) GetOrgCode() int64 {
	if o == nil || o.OrgCode.Get() == nil {
		var ret int64
		return ret
	}
	return *o.OrgCode.Get()
}

// GetOrgCodeOk returns a tuple with the OrgCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerResponsePartner) GetOrgCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrgCode.Get(), o.OrgCode.IsSet()
}

// HasOrgCode returns a boolean if a field has been set.
func (o *PartnerResponsePartner) HasOrgCode() bool {
	if o != nil && o.OrgCode.IsSet() {
		return true
	}

	return false
}

// SetOrgCode gets a reference to the given NullableInt64 and assigns it to the OrgCode field.
func (o *PartnerResponsePartner) SetOrgCode(v int64) {
	o.OrgCode.Set(&v)
}
// SetOrgCodeNil sets the value for OrgCode to be an explicit nil
func (o *PartnerResponsePartner) SetOrgCodeNil() {
	o.OrgCode.Set(nil)
}

// UnsetOrgCode ensures that no value is present for OrgCode, not even an explicit nil
func (o *PartnerResponsePartner) UnsetOrgCode() {
	o.OrgCode.Unset()
}

// GetPartnerBankAccountAttributes returns the PartnerBankAccountAttributes field value if set, zero value otherwise.
func (o *PartnerResponsePartner) GetPartnerBankAccountAttributes() PartnerResponsePartnerPartnerBankAccountAttributes {
	if o == nil || o.PartnerBankAccountAttributes == nil {
		var ret PartnerResponsePartnerPartnerBankAccountAttributes
		return ret
	}
	return *o.PartnerBankAccountAttributes
}

// GetPartnerBankAccountAttributesOk returns a tuple with the PartnerBankAccountAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerResponsePartner) GetPartnerBankAccountAttributesOk() (*PartnerResponsePartnerPartnerBankAccountAttributes, bool) {
	if o == nil || o.PartnerBankAccountAttributes == nil {
		return nil, false
	}
	return o.PartnerBankAccountAttributes, true
}

// HasPartnerBankAccountAttributes returns a boolean if a field has been set.
func (o *PartnerResponsePartner) HasPartnerBankAccountAttributes() bool {
	if o != nil && o.PartnerBankAccountAttributes != nil {
		return true
	}

	return false
}

// SetPartnerBankAccountAttributes gets a reference to the given PartnerResponsePartnerPartnerBankAccountAttributes and assigns it to the PartnerBankAccountAttributes field.
func (o *PartnerResponsePartner) SetPartnerBankAccountAttributes(v PartnerResponsePartnerPartnerBankAccountAttributes) {
	o.PartnerBankAccountAttributes = &v
}

// GetPartnerDocSettingAttributes returns the PartnerDocSettingAttributes field value if set, zero value otherwise.
func (o *PartnerResponsePartner) GetPartnerDocSettingAttributes() PartnerCreateParamsPartnerDocSettingAttributes {
	if o == nil || o.PartnerDocSettingAttributes == nil {
		var ret PartnerCreateParamsPartnerDocSettingAttributes
		return ret
	}
	return *o.PartnerDocSettingAttributes
}

// GetPartnerDocSettingAttributesOk returns a tuple with the PartnerDocSettingAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerResponsePartner) GetPartnerDocSettingAttributesOk() (*PartnerCreateParamsPartnerDocSettingAttributes, bool) {
	if o == nil || o.PartnerDocSettingAttributes == nil {
		return nil, false
	}
	return o.PartnerDocSettingAttributes, true
}

// HasPartnerDocSettingAttributes returns a boolean if a field has been set.
func (o *PartnerResponsePartner) HasPartnerDocSettingAttributes() bool {
	if o != nil && o.PartnerDocSettingAttributes != nil {
		return true
	}

	return false
}

// SetPartnerDocSettingAttributes gets a reference to the given PartnerCreateParamsPartnerDocSettingAttributes and assigns it to the PartnerDocSettingAttributes field.
func (o *PartnerResponsePartner) SetPartnerDocSettingAttributes(v PartnerCreateParamsPartnerDocSettingAttributes) {
	o.PartnerDocSettingAttributes = &v
}

// GetPayerWalletableId returns the PayerWalletableId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerResponsePartner) GetPayerWalletableId() int64 {
	if o == nil || o.PayerWalletableId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.PayerWalletableId.Get()
}

// GetPayerWalletableIdOk returns a tuple with the PayerWalletableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerResponsePartner) GetPayerWalletableIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PayerWalletableId.Get(), o.PayerWalletableId.IsSet()
}

// HasPayerWalletableId returns a boolean if a field has been set.
func (o *PartnerResponsePartner) HasPayerWalletableId() bool {
	if o != nil && o.PayerWalletableId.IsSet() {
		return true
	}

	return false
}

// SetPayerWalletableId gets a reference to the given NullableInt64 and assigns it to the PayerWalletableId field.
func (o *PartnerResponsePartner) SetPayerWalletableId(v int64) {
	o.PayerWalletableId.Set(&v)
}
// SetPayerWalletableIdNil sets the value for PayerWalletableId to be an explicit nil
func (o *PartnerResponsePartner) SetPayerWalletableIdNil() {
	o.PayerWalletableId.Set(nil)
}

// UnsetPayerWalletableId ensures that no value is present for PayerWalletableId, not even an explicit nil
func (o *PartnerResponsePartner) UnsetPayerWalletableId() {
	o.PayerWalletableId.Unset()
}

// GetPaymentTermAttributes returns the PaymentTermAttributes field value if set, zero value otherwise.
func (o *PartnerResponsePartner) GetPaymentTermAttributes() PartnerResponsePartnerPaymentTermAttributes {
	if o == nil || o.PaymentTermAttributes == nil {
		var ret PartnerResponsePartnerPaymentTermAttributes
		return ret
	}
	return *o.PaymentTermAttributes
}

// GetPaymentTermAttributesOk returns a tuple with the PaymentTermAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerResponsePartner) GetPaymentTermAttributesOk() (*PartnerResponsePartnerPaymentTermAttributes, bool) {
	if o == nil || o.PaymentTermAttributes == nil {
		return nil, false
	}
	return o.PaymentTermAttributes, true
}

// HasPaymentTermAttributes returns a boolean if a field has been set.
func (o *PartnerResponsePartner) HasPaymentTermAttributes() bool {
	if o != nil && o.PaymentTermAttributes != nil {
		return true
	}

	return false
}

// SetPaymentTermAttributes gets a reference to the given PartnerResponsePartnerPaymentTermAttributes and assigns it to the PaymentTermAttributes field.
func (o *PartnerResponsePartner) SetPaymentTermAttributes(v PartnerResponsePartnerPaymentTermAttributes) {
	o.PaymentTermAttributes = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerResponsePartner) GetPhone() string {
	if o == nil || o.Phone.Get() == nil {
		var ret string
		return ret
	}
	return *o.Phone.Get()
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerResponsePartner) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phone.Get(), o.Phone.IsSet()
}

// HasPhone returns a boolean if a field has been set.
func (o *PartnerResponsePartner) HasPhone() bool {
	if o != nil && o.Phone.IsSet() {
		return true
	}

	return false
}

// SetPhone gets a reference to the given NullableString and assigns it to the Phone field.
func (o *PartnerResponsePartner) SetPhone(v string) {
	o.Phone.Set(&v)
}
// SetPhoneNil sets the value for Phone to be an explicit nil
func (o *PartnerResponsePartner) SetPhoneNil() {
	o.Phone.Set(nil)
}

// UnsetPhone ensures that no value is present for Phone, not even an explicit nil
func (o *PartnerResponsePartner) UnsetPhone() {
	o.Phone.Unset()
}

// GetShortcut1 returns the Shortcut1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerResponsePartner) GetShortcut1() string {
	if o == nil || o.Shortcut1.Get() == nil {
		var ret string
		return ret
	}
	return *o.Shortcut1.Get()
}

// GetShortcut1Ok returns a tuple with the Shortcut1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerResponsePartner) GetShortcut1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Shortcut1.Get(), o.Shortcut1.IsSet()
}

// HasShortcut1 returns a boolean if a field has been set.
func (o *PartnerResponsePartner) HasShortcut1() bool {
	if o != nil && o.Shortcut1.IsSet() {
		return true
	}

	return false
}

// SetShortcut1 gets a reference to the given NullableString and assigns it to the Shortcut1 field.
func (o *PartnerResponsePartner) SetShortcut1(v string) {
	o.Shortcut1.Set(&v)
}
// SetShortcut1Nil sets the value for Shortcut1 to be an explicit nil
func (o *PartnerResponsePartner) SetShortcut1Nil() {
	o.Shortcut1.Set(nil)
}

// UnsetShortcut1 ensures that no value is present for Shortcut1, not even an explicit nil
func (o *PartnerResponsePartner) UnsetShortcut1() {
	o.Shortcut1.Unset()
}

// GetShortcut2 returns the Shortcut2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerResponsePartner) GetShortcut2() string {
	if o == nil || o.Shortcut2.Get() == nil {
		var ret string
		return ret
	}
	return *o.Shortcut2.Get()
}

// GetShortcut2Ok returns a tuple with the Shortcut2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerResponsePartner) GetShortcut2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Shortcut2.Get(), o.Shortcut2.IsSet()
}

// HasShortcut2 returns a boolean if a field has been set.
func (o *PartnerResponsePartner) HasShortcut2() bool {
	if o != nil && o.Shortcut2.IsSet() {
		return true
	}

	return false
}

// SetShortcut2 gets a reference to the given NullableString and assigns it to the Shortcut2 field.
func (o *PartnerResponsePartner) SetShortcut2(v string) {
	o.Shortcut2.Set(&v)
}
// SetShortcut2Nil sets the value for Shortcut2 to be an explicit nil
func (o *PartnerResponsePartner) SetShortcut2Nil() {
	o.Shortcut2.Set(nil)
}

// UnsetShortcut2 ensures that no value is present for Shortcut2, not even an explicit nil
func (o *PartnerResponsePartner) UnsetShortcut2() {
	o.Shortcut2.Unset()
}

// GetTransferFeeHandlingSide returns the TransferFeeHandlingSide field value if set, zero value otherwise.
func (o *PartnerResponsePartner) GetTransferFeeHandlingSide() string {
	if o == nil || o.TransferFeeHandlingSide == nil {
		var ret string
		return ret
	}
	return *o.TransferFeeHandlingSide
}

// GetTransferFeeHandlingSideOk returns a tuple with the TransferFeeHandlingSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerResponsePartner) GetTransferFeeHandlingSideOk() (*string, bool) {
	if o == nil || o.TransferFeeHandlingSide == nil {
		return nil, false
	}
	return o.TransferFeeHandlingSide, true
}

// HasTransferFeeHandlingSide returns a boolean if a field has been set.
func (o *PartnerResponsePartner) HasTransferFeeHandlingSide() bool {
	if o != nil && o.TransferFeeHandlingSide != nil {
		return true
	}

	return false
}

// SetTransferFeeHandlingSide gets a reference to the given string and assigns it to the TransferFeeHandlingSide field.
func (o *PartnerResponsePartner) SetTransferFeeHandlingSide(v string) {
	o.TransferFeeHandlingSide = &v
}

// GetUpdateDate returns the UpdateDate field value
func (o *PartnerResponsePartner) GetUpdateDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdateDate
}

// GetUpdateDateOk returns a tuple with the UpdateDate field value
// and a boolean to check if the value has been set.
func (o *PartnerResponsePartner) GetUpdateDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateDate, true
}

// SetUpdateDate sets field value
func (o *PartnerResponsePartner) SetUpdateDate(v string) {
	o.UpdateDate = v
}

func (o PartnerResponsePartner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddressAttributes != nil {
		toSerialize["address_attributes"] = o.AddressAttributes
	}
	if true {
		toSerialize["available"] = o.Available
	}
	if true {
		toSerialize["code"] = o.Code.Get()
	}
	if true {
		toSerialize["company_id"] = o.CompanyId
	}
	if o.ContactName.IsSet() {
		toSerialize["contact_name"] = o.ContactName.Get()
	}
	if o.CountryCode != nil {
		toSerialize["country_code"] = o.CountryCode
	}
	if o.DefaultTitle.IsSet() {
		toSerialize["default_title"] = o.DefaultTitle.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.InvoicePaymentTermAttributes != nil {
		toSerialize["invoice_payment_term_attributes"] = o.InvoicePaymentTermAttributes
	}
	if o.LongName.IsSet() {
		toSerialize["long_name"] = o.LongName.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NameKana.IsSet() {
		toSerialize["name_kana"] = o.NameKana.Get()
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
	if o.Phone.IsSet() {
		toSerialize["phone"] = o.Phone.Get()
	}
	if o.Shortcut1.IsSet() {
		toSerialize["shortcut1"] = o.Shortcut1.Get()
	}
	if o.Shortcut2.IsSet() {
		toSerialize["shortcut2"] = o.Shortcut2.Get()
	}
	if o.TransferFeeHandlingSide != nil {
		toSerialize["transfer_fee_handling_side"] = o.TransferFeeHandlingSide
	}
	if true {
		toSerialize["update_date"] = o.UpdateDate
	}
	return json.Marshal(toSerialize)
}

type NullablePartnerResponsePartner struct {
	value *PartnerResponsePartner
	isSet bool
}

func (v NullablePartnerResponsePartner) Get() *PartnerResponsePartner {
	return v.value
}

func (v *NullablePartnerResponsePartner) Set(val *PartnerResponsePartner) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerResponsePartner) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerResponsePartner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerResponsePartner(val *PartnerResponsePartner) *NullablePartnerResponsePartner {
	return &NullablePartnerResponsePartner{value: val, isSet: true}
}

func (v NullablePartnerResponsePartner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerResponsePartner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


