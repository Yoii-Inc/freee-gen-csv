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

// PartnerResponsePartnerPartnerBankAccountAttributes struct for PartnerResponsePartnerPartnerBankAccountAttributes
type PartnerResponsePartnerPartnerBankAccountAttributes struct {
	// 受取人名（カナ）
	AccountName NullableString `json:"account_name,omitempty"`
	// 口座番号
	AccountNumber NullableString `json:"account_number,omitempty"`
	// 口座種別(ordinary:普通、checking:当座、earmarked:納税準備預金、savings:貯蓄、other:その他)
	AccountType NullableString `json:"account_type,omitempty"`
	// 銀行コード
	BankCode NullableString `json:"bank_code,omitempty"`
	// 銀行名
	BankName NullableString `json:"bank_name,omitempty"`
	// 銀行名（カナ）
	BankNameKana NullableString `json:"bank_name_kana,omitempty"`
	// 支店番号
	BranchCode NullableString `json:"branch_code,omitempty"`
	// 支店名（カナ）
	BranchKana NullableString `json:"branch_kana,omitempty"`
	// 支店名
	BranchName NullableString `json:"branch_name,omitempty"`
	// 受取人名
	LongAccountName NullableString `json:"long_account_name,omitempty"`
}

// NewPartnerResponsePartnerPartnerBankAccountAttributes instantiates a new PartnerResponsePartnerPartnerBankAccountAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerResponsePartnerPartnerBankAccountAttributes() *PartnerResponsePartnerPartnerBankAccountAttributes {
	this := PartnerResponsePartnerPartnerBankAccountAttributes{}
	return &this
}

// NewPartnerResponsePartnerPartnerBankAccountAttributesWithDefaults instantiates a new PartnerResponsePartnerPartnerBankAccountAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerResponsePartnerPartnerBankAccountAttributesWithDefaults() *PartnerResponsePartnerPartnerBankAccountAttributes {
	this := PartnerResponsePartnerPartnerBankAccountAttributes{}
	return &this
}

// GetAccountName returns the AccountName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) GetAccountName() string {
	if o == nil || o.AccountName.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountName.Get()
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) GetAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountName.Get(), o.AccountName.IsSet()
}

// HasAccountName returns a boolean if a field has been set.
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) HasAccountName() bool {
	if o != nil && o.AccountName.IsSet() {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given NullableString and assigns it to the AccountName field.
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) SetAccountName(v string) {
	o.AccountName.Set(&v)
}
// SetAccountNameNil sets the value for AccountName to be an explicit nil
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) SetAccountNameNil() {
	o.AccountName.Set(nil)
}

// UnsetAccountName ensures that no value is present for AccountName, not even an explicit nil
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) UnsetAccountName() {
	o.AccountName.Unset()
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) GetAccountNumber() string {
	if o == nil || o.AccountNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountNumber.Get()
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountNumber.Get(), o.AccountNumber.IsSet()
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) HasAccountNumber() bool {
	if o != nil && o.AccountNumber.IsSet() {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given NullableString and assigns it to the AccountNumber field.
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) SetAccountNumber(v string) {
	o.AccountNumber.Set(&v)
}
// SetAccountNumberNil sets the value for AccountNumber to be an explicit nil
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) SetAccountNumberNil() {
	o.AccountNumber.Set(nil)
}

// UnsetAccountNumber ensures that no value is present for AccountNumber, not even an explicit nil
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) UnsetAccountNumber() {
	o.AccountNumber.Unset()
}

// GetAccountType returns the AccountType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) GetAccountType() string {
	if o == nil || o.AccountType.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountType.Get()
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) GetAccountTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountType.Get(), o.AccountType.IsSet()
}

// HasAccountType returns a boolean if a field has been set.
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) HasAccountType() bool {
	if o != nil && o.AccountType.IsSet() {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given NullableString and assigns it to the AccountType field.
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) SetAccountType(v string) {
	o.AccountType.Set(&v)
}
// SetAccountTypeNil sets the value for AccountType to be an explicit nil
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) SetAccountTypeNil() {
	o.AccountType.Set(nil)
}

// UnsetAccountType ensures that no value is present for AccountType, not even an explicit nil
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) UnsetAccountType() {
	o.AccountType.Unset()
}

// GetBankCode returns the BankCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) GetBankCode() string {
	if o == nil || o.BankCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.BankCode.Get()
}

// GetBankCodeOk returns a tuple with the BankCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) GetBankCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BankCode.Get(), o.BankCode.IsSet()
}

// HasBankCode returns a boolean if a field has been set.
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) HasBankCode() bool {
	if o != nil && o.BankCode.IsSet() {
		return true
	}

	return false
}

// SetBankCode gets a reference to the given NullableString and assigns it to the BankCode field.
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) SetBankCode(v string) {
	o.BankCode.Set(&v)
}
// SetBankCodeNil sets the value for BankCode to be an explicit nil
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) SetBankCodeNil() {
	o.BankCode.Set(nil)
}

// UnsetBankCode ensures that no value is present for BankCode, not even an explicit nil
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) UnsetBankCode() {
	o.BankCode.Unset()
}

// GetBankName returns the BankName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) GetBankName() string {
	if o == nil || o.BankName.Get() == nil {
		var ret string
		return ret
	}
	return *o.BankName.Get()
}

// GetBankNameOk returns a tuple with the BankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) GetBankNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BankName.Get(), o.BankName.IsSet()
}

// HasBankName returns a boolean if a field has been set.
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) HasBankName() bool {
	if o != nil && o.BankName.IsSet() {
		return true
	}

	return false
}

// SetBankName gets a reference to the given NullableString and assigns it to the BankName field.
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) SetBankName(v string) {
	o.BankName.Set(&v)
}
// SetBankNameNil sets the value for BankName to be an explicit nil
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) SetBankNameNil() {
	o.BankName.Set(nil)
}

// UnsetBankName ensures that no value is present for BankName, not even an explicit nil
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) UnsetBankName() {
	o.BankName.Unset()
}

// GetBankNameKana returns the BankNameKana field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) GetBankNameKana() string {
	if o == nil || o.BankNameKana.Get() == nil {
		var ret string
		return ret
	}
	return *o.BankNameKana.Get()
}

// GetBankNameKanaOk returns a tuple with the BankNameKana field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) GetBankNameKanaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BankNameKana.Get(), o.BankNameKana.IsSet()
}

// HasBankNameKana returns a boolean if a field has been set.
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) HasBankNameKana() bool {
	if o != nil && o.BankNameKana.IsSet() {
		return true
	}

	return false
}

// SetBankNameKana gets a reference to the given NullableString and assigns it to the BankNameKana field.
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) SetBankNameKana(v string) {
	o.BankNameKana.Set(&v)
}
// SetBankNameKanaNil sets the value for BankNameKana to be an explicit nil
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) SetBankNameKanaNil() {
	o.BankNameKana.Set(nil)
}

// UnsetBankNameKana ensures that no value is present for BankNameKana, not even an explicit nil
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) UnsetBankNameKana() {
	o.BankNameKana.Unset()
}

// GetBranchCode returns the BranchCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) GetBranchCode() string {
	if o == nil || o.BranchCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.BranchCode.Get()
}

// GetBranchCodeOk returns a tuple with the BranchCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) GetBranchCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BranchCode.Get(), o.BranchCode.IsSet()
}

// HasBranchCode returns a boolean if a field has been set.
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) HasBranchCode() bool {
	if o != nil && o.BranchCode.IsSet() {
		return true
	}

	return false
}

// SetBranchCode gets a reference to the given NullableString and assigns it to the BranchCode field.
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) SetBranchCode(v string) {
	o.BranchCode.Set(&v)
}
// SetBranchCodeNil sets the value for BranchCode to be an explicit nil
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) SetBranchCodeNil() {
	o.BranchCode.Set(nil)
}

// UnsetBranchCode ensures that no value is present for BranchCode, not even an explicit nil
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) UnsetBranchCode() {
	o.BranchCode.Unset()
}

// GetBranchKana returns the BranchKana field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) GetBranchKana() string {
	if o == nil || o.BranchKana.Get() == nil {
		var ret string
		return ret
	}
	return *o.BranchKana.Get()
}

// GetBranchKanaOk returns a tuple with the BranchKana field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) GetBranchKanaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BranchKana.Get(), o.BranchKana.IsSet()
}

// HasBranchKana returns a boolean if a field has been set.
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) HasBranchKana() bool {
	if o != nil && o.BranchKana.IsSet() {
		return true
	}

	return false
}

// SetBranchKana gets a reference to the given NullableString and assigns it to the BranchKana field.
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) SetBranchKana(v string) {
	o.BranchKana.Set(&v)
}
// SetBranchKanaNil sets the value for BranchKana to be an explicit nil
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) SetBranchKanaNil() {
	o.BranchKana.Set(nil)
}

// UnsetBranchKana ensures that no value is present for BranchKana, not even an explicit nil
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) UnsetBranchKana() {
	o.BranchKana.Unset()
}

// GetBranchName returns the BranchName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) GetBranchName() string {
	if o == nil || o.BranchName.Get() == nil {
		var ret string
		return ret
	}
	return *o.BranchName.Get()
}

// GetBranchNameOk returns a tuple with the BranchName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) GetBranchNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BranchName.Get(), o.BranchName.IsSet()
}

// HasBranchName returns a boolean if a field has been set.
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) HasBranchName() bool {
	if o != nil && o.BranchName.IsSet() {
		return true
	}

	return false
}

// SetBranchName gets a reference to the given NullableString and assigns it to the BranchName field.
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) SetBranchName(v string) {
	o.BranchName.Set(&v)
}
// SetBranchNameNil sets the value for BranchName to be an explicit nil
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) SetBranchNameNil() {
	o.BranchName.Set(nil)
}

// UnsetBranchName ensures that no value is present for BranchName, not even an explicit nil
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) UnsetBranchName() {
	o.BranchName.Unset()
}

// GetLongAccountName returns the LongAccountName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) GetLongAccountName() string {
	if o == nil || o.LongAccountName.Get() == nil {
		var ret string
		return ret
	}
	return *o.LongAccountName.Get()
}

// GetLongAccountNameOk returns a tuple with the LongAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) GetLongAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LongAccountName.Get(), o.LongAccountName.IsSet()
}

// HasLongAccountName returns a boolean if a field has been set.
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) HasLongAccountName() bool {
	if o != nil && o.LongAccountName.IsSet() {
		return true
	}

	return false
}

// SetLongAccountName gets a reference to the given NullableString and assigns it to the LongAccountName field.
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) SetLongAccountName(v string) {
	o.LongAccountName.Set(&v)
}
// SetLongAccountNameNil sets the value for LongAccountName to be an explicit nil
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) SetLongAccountNameNil() {
	o.LongAccountName.Set(nil)
}

// UnsetLongAccountName ensures that no value is present for LongAccountName, not even an explicit nil
func (o *PartnerResponsePartnerPartnerBankAccountAttributes) UnsetLongAccountName() {
	o.LongAccountName.Unset()
}

func (o PartnerResponsePartnerPartnerBankAccountAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountName.IsSet() {
		toSerialize["account_name"] = o.AccountName.Get()
	}
	if o.AccountNumber.IsSet() {
		toSerialize["account_number"] = o.AccountNumber.Get()
	}
	if o.AccountType.IsSet() {
		toSerialize["account_type"] = o.AccountType.Get()
	}
	if o.BankCode.IsSet() {
		toSerialize["bank_code"] = o.BankCode.Get()
	}
	if o.BankName.IsSet() {
		toSerialize["bank_name"] = o.BankName.Get()
	}
	if o.BankNameKana.IsSet() {
		toSerialize["bank_name_kana"] = o.BankNameKana.Get()
	}
	if o.BranchCode.IsSet() {
		toSerialize["branch_code"] = o.BranchCode.Get()
	}
	if o.BranchKana.IsSet() {
		toSerialize["branch_kana"] = o.BranchKana.Get()
	}
	if o.BranchName.IsSet() {
		toSerialize["branch_name"] = o.BranchName.Get()
	}
	if o.LongAccountName.IsSet() {
		toSerialize["long_account_name"] = o.LongAccountName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePartnerResponsePartnerPartnerBankAccountAttributes struct {
	value *PartnerResponsePartnerPartnerBankAccountAttributes
	isSet bool
}

func (v NullablePartnerResponsePartnerPartnerBankAccountAttributes) Get() *PartnerResponsePartnerPartnerBankAccountAttributes {
	return v.value
}

func (v *NullablePartnerResponsePartnerPartnerBankAccountAttributes) Set(val *PartnerResponsePartnerPartnerBankAccountAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerResponsePartnerPartnerBankAccountAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerResponsePartnerPartnerBankAccountAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerResponsePartnerPartnerBankAccountAttributes(val *PartnerResponsePartnerPartnerBankAccountAttributes) *NullablePartnerResponsePartnerPartnerBankAccountAttributes {
	return &NullablePartnerResponsePartnerPartnerBankAccountAttributes{value: val, isSet: true}
}

func (v NullablePartnerResponsePartnerPartnerBankAccountAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerResponsePartnerPartnerBankAccountAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


