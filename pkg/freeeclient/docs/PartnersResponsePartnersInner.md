# PartnersResponsePartnersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressAttributes** | Pointer to [**PartnersResponsePartnersInnerAddressAttributes**](PartnersResponsePartnersInnerAddressAttributes.md) |  | [optional] 
**Available** | **bool** | 取引先の使用設定（true: 使用する、false: 使用しない） &lt;br&gt; &lt;ul&gt;   &lt;li&gt;     本APIでpartnerを作成した場合はtrueになります。   &lt;/li&gt;   &lt;li&gt;     falseにする場合はWeb画面から変更できます。   &lt;/li&gt;   &lt;li&gt;     trueの場合、Web画面での取引登録時などに入力候補として表示されます。   &lt;/li&gt;   &lt;li&gt;     falseの場合、取引先自体は削除せず、Web画面での取引登録時などに入力候補として表示されません。ただし取引（収入／支出）の作成APIなどでfalseの取引先をパラメータに指定すれば、取引などにfalseの取引先を設定できます。   &lt;/li&gt; &lt;/ul&gt; | 
**Code** | **NullableString** | 取引先コード | 
**CompanyId** | **int64** | 事業所ID | 
**ContactName** | Pointer to **NullableString** | 担当者 氏名 | [optional] 
**CountryCode** | Pointer to **string** | 地域（JP: 国内、ZZ:国外） | [optional] 
**DefaultTitle** | Pointer to **NullableString** | 敬称（御中、様、(空白)の3つから選択） | [optional] 
**Email** | Pointer to **NullableString** | 担当者 メールアドレス | [optional] 
**Id** | **int64** | 取引先ID | 
**LongName** | Pointer to **NullableString** | 正式名称（255文字以内） | [optional] 
**Name** | **string** | 取引先名 | 
**NameKana** | Pointer to **NullableString** | カナ名称（255文字以内） | [optional] 
**OrgCode** | Pointer to **NullableInt64** | 事業所種別（null: 未設定、1: 法人、2: 個人） | [optional] 
**PartnerBankAccountAttributes** | Pointer to [**PartnerResponsePartnerPartnerBankAccountAttributes**](PartnerResponsePartnerPartnerBankAccountAttributes.md) |  | [optional] 
**PartnerDocSettingAttributes** | Pointer to [**PartnerCreateParamsPartnerDocSettingAttributes**](PartnerCreateParamsPartnerDocSettingAttributes.md) |  | [optional] 
**PayerWalletableId** | Pointer to **NullableInt64** | 振込元口座ID（一括振込ファイル用）:（未設定の場合は、nullです。） | [optional] 
**Phone** | Pointer to **NullableString** | 電話番号 | [optional] 
**Shortcut1** | Pointer to **NullableString** | ショートカット1 (255文字以内) | [optional] 
**Shortcut2** | Pointer to **NullableString** | ショートカット2 (255文字以内) | [optional] 
**TransferFeeHandlingSide** | Pointer to **string** | 振込手数料負担（一括振込ファイル用）: (振込元(当方): payer, 振込先(先方): payee) | [optional] 
**UpdateDate** | **string** | 更新日 (yyyy-mm-dd) | 

## Methods

### NewPartnersResponsePartnersInner

`func NewPartnersResponsePartnersInner(available bool, code NullableString, companyId int64, id int64, name string, updateDate string, ) *PartnersResponsePartnersInner`

NewPartnersResponsePartnersInner instantiates a new PartnersResponsePartnersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnersResponsePartnersInnerWithDefaults

`func NewPartnersResponsePartnersInnerWithDefaults() *PartnersResponsePartnersInner`

NewPartnersResponsePartnersInnerWithDefaults instantiates a new PartnersResponsePartnersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressAttributes

`func (o *PartnersResponsePartnersInner) GetAddressAttributes() PartnersResponsePartnersInnerAddressAttributes`

GetAddressAttributes returns the AddressAttributes field if non-nil, zero value otherwise.

### GetAddressAttributesOk

`func (o *PartnersResponsePartnersInner) GetAddressAttributesOk() (*PartnersResponsePartnersInnerAddressAttributes, bool)`

GetAddressAttributesOk returns a tuple with the AddressAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressAttributes

`func (o *PartnersResponsePartnersInner) SetAddressAttributes(v PartnersResponsePartnersInnerAddressAttributes)`

SetAddressAttributes sets AddressAttributes field to given value.

### HasAddressAttributes

`func (o *PartnersResponsePartnersInner) HasAddressAttributes() bool`

HasAddressAttributes returns a boolean if a field has been set.

### GetAvailable

`func (o *PartnersResponsePartnersInner) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *PartnersResponsePartnersInner) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *PartnersResponsePartnersInner) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetCode

`func (o *PartnersResponsePartnersInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PartnersResponsePartnersInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PartnersResponsePartnersInner) SetCode(v string)`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *PartnersResponsePartnersInner) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *PartnersResponsePartnersInner) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCompanyId

`func (o *PartnersResponsePartnersInner) GetCompanyId() int64`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *PartnersResponsePartnersInner) GetCompanyIdOk() (*int64, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *PartnersResponsePartnersInner) SetCompanyId(v int64)`

SetCompanyId sets CompanyId field to given value.


### GetContactName

`func (o *PartnersResponsePartnersInner) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *PartnersResponsePartnersInner) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *PartnersResponsePartnersInner) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *PartnersResponsePartnersInner) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### SetContactNameNil

`func (o *PartnersResponsePartnersInner) SetContactNameNil(b bool)`

 SetContactNameNil sets the value for ContactName to be an explicit nil

### UnsetContactName
`func (o *PartnersResponsePartnersInner) UnsetContactName()`

UnsetContactName ensures that no value is present for ContactName, not even an explicit nil
### GetCountryCode

`func (o *PartnersResponsePartnersInner) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *PartnersResponsePartnersInner) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *PartnersResponsePartnersInner) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *PartnersResponsePartnersInner) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetDefaultTitle

`func (o *PartnersResponsePartnersInner) GetDefaultTitle() string`

GetDefaultTitle returns the DefaultTitle field if non-nil, zero value otherwise.

### GetDefaultTitleOk

`func (o *PartnersResponsePartnersInner) GetDefaultTitleOk() (*string, bool)`

GetDefaultTitleOk returns a tuple with the DefaultTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTitle

`func (o *PartnersResponsePartnersInner) SetDefaultTitle(v string)`

SetDefaultTitle sets DefaultTitle field to given value.

### HasDefaultTitle

`func (o *PartnersResponsePartnersInner) HasDefaultTitle() bool`

HasDefaultTitle returns a boolean if a field has been set.

### SetDefaultTitleNil

`func (o *PartnersResponsePartnersInner) SetDefaultTitleNil(b bool)`

 SetDefaultTitleNil sets the value for DefaultTitle to be an explicit nil

### UnsetDefaultTitle
`func (o *PartnersResponsePartnersInner) UnsetDefaultTitle()`

UnsetDefaultTitle ensures that no value is present for DefaultTitle, not even an explicit nil
### GetEmail

`func (o *PartnersResponsePartnersInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PartnersResponsePartnersInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PartnersResponsePartnersInner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PartnersResponsePartnersInner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *PartnersResponsePartnersInner) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *PartnersResponsePartnersInner) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetId

`func (o *PartnersResponsePartnersInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartnersResponsePartnersInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartnersResponsePartnersInner) SetId(v int64)`

SetId sets Id field to given value.


### GetLongName

`func (o *PartnersResponsePartnersInner) GetLongName() string`

GetLongName returns the LongName field if non-nil, zero value otherwise.

### GetLongNameOk

`func (o *PartnersResponsePartnersInner) GetLongNameOk() (*string, bool)`

GetLongNameOk returns a tuple with the LongName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongName

`func (o *PartnersResponsePartnersInner) SetLongName(v string)`

SetLongName sets LongName field to given value.

### HasLongName

`func (o *PartnersResponsePartnersInner) HasLongName() bool`

HasLongName returns a boolean if a field has been set.

### SetLongNameNil

`func (o *PartnersResponsePartnersInner) SetLongNameNil(b bool)`

 SetLongNameNil sets the value for LongName to be an explicit nil

### UnsetLongName
`func (o *PartnersResponsePartnersInner) UnsetLongName()`

UnsetLongName ensures that no value is present for LongName, not even an explicit nil
### GetName

`func (o *PartnersResponsePartnersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartnersResponsePartnersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartnersResponsePartnersInner) SetName(v string)`

SetName sets Name field to given value.


### GetNameKana

`func (o *PartnersResponsePartnersInner) GetNameKana() string`

GetNameKana returns the NameKana field if non-nil, zero value otherwise.

### GetNameKanaOk

`func (o *PartnersResponsePartnersInner) GetNameKanaOk() (*string, bool)`

GetNameKanaOk returns a tuple with the NameKana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameKana

`func (o *PartnersResponsePartnersInner) SetNameKana(v string)`

SetNameKana sets NameKana field to given value.

### HasNameKana

`func (o *PartnersResponsePartnersInner) HasNameKana() bool`

HasNameKana returns a boolean if a field has been set.

### SetNameKanaNil

`func (o *PartnersResponsePartnersInner) SetNameKanaNil(b bool)`

 SetNameKanaNil sets the value for NameKana to be an explicit nil

### UnsetNameKana
`func (o *PartnersResponsePartnersInner) UnsetNameKana()`

UnsetNameKana ensures that no value is present for NameKana, not even an explicit nil
### GetOrgCode

`func (o *PartnersResponsePartnersInner) GetOrgCode() int64`

GetOrgCode returns the OrgCode field if non-nil, zero value otherwise.

### GetOrgCodeOk

`func (o *PartnersResponsePartnersInner) GetOrgCodeOk() (*int64, bool)`

GetOrgCodeOk returns a tuple with the OrgCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgCode

`func (o *PartnersResponsePartnersInner) SetOrgCode(v int64)`

SetOrgCode sets OrgCode field to given value.

### HasOrgCode

`func (o *PartnersResponsePartnersInner) HasOrgCode() bool`

HasOrgCode returns a boolean if a field has been set.

### SetOrgCodeNil

`func (o *PartnersResponsePartnersInner) SetOrgCodeNil(b bool)`

 SetOrgCodeNil sets the value for OrgCode to be an explicit nil

### UnsetOrgCode
`func (o *PartnersResponsePartnersInner) UnsetOrgCode()`

UnsetOrgCode ensures that no value is present for OrgCode, not even an explicit nil
### GetPartnerBankAccountAttributes

`func (o *PartnersResponsePartnersInner) GetPartnerBankAccountAttributes() PartnerResponsePartnerPartnerBankAccountAttributes`

GetPartnerBankAccountAttributes returns the PartnerBankAccountAttributes field if non-nil, zero value otherwise.

### GetPartnerBankAccountAttributesOk

`func (o *PartnersResponsePartnersInner) GetPartnerBankAccountAttributesOk() (*PartnerResponsePartnerPartnerBankAccountAttributes, bool)`

GetPartnerBankAccountAttributesOk returns a tuple with the PartnerBankAccountAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerBankAccountAttributes

`func (o *PartnersResponsePartnersInner) SetPartnerBankAccountAttributes(v PartnerResponsePartnerPartnerBankAccountAttributes)`

SetPartnerBankAccountAttributes sets PartnerBankAccountAttributes field to given value.

### HasPartnerBankAccountAttributes

`func (o *PartnersResponsePartnersInner) HasPartnerBankAccountAttributes() bool`

HasPartnerBankAccountAttributes returns a boolean if a field has been set.

### GetPartnerDocSettingAttributes

`func (o *PartnersResponsePartnersInner) GetPartnerDocSettingAttributes() PartnerCreateParamsPartnerDocSettingAttributes`

GetPartnerDocSettingAttributes returns the PartnerDocSettingAttributes field if non-nil, zero value otherwise.

### GetPartnerDocSettingAttributesOk

`func (o *PartnersResponsePartnersInner) GetPartnerDocSettingAttributesOk() (*PartnerCreateParamsPartnerDocSettingAttributes, bool)`

GetPartnerDocSettingAttributesOk returns a tuple with the PartnerDocSettingAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerDocSettingAttributes

`func (o *PartnersResponsePartnersInner) SetPartnerDocSettingAttributes(v PartnerCreateParamsPartnerDocSettingAttributes)`

SetPartnerDocSettingAttributes sets PartnerDocSettingAttributes field to given value.

### HasPartnerDocSettingAttributes

`func (o *PartnersResponsePartnersInner) HasPartnerDocSettingAttributes() bool`

HasPartnerDocSettingAttributes returns a boolean if a field has been set.

### GetPayerWalletableId

`func (o *PartnersResponsePartnersInner) GetPayerWalletableId() int64`

GetPayerWalletableId returns the PayerWalletableId field if non-nil, zero value otherwise.

### GetPayerWalletableIdOk

`func (o *PartnersResponsePartnersInner) GetPayerWalletableIdOk() (*int64, bool)`

GetPayerWalletableIdOk returns a tuple with the PayerWalletableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerWalletableId

`func (o *PartnersResponsePartnersInner) SetPayerWalletableId(v int64)`

SetPayerWalletableId sets PayerWalletableId field to given value.

### HasPayerWalletableId

`func (o *PartnersResponsePartnersInner) HasPayerWalletableId() bool`

HasPayerWalletableId returns a boolean if a field has been set.

### SetPayerWalletableIdNil

`func (o *PartnersResponsePartnersInner) SetPayerWalletableIdNil(b bool)`

 SetPayerWalletableIdNil sets the value for PayerWalletableId to be an explicit nil

### UnsetPayerWalletableId
`func (o *PartnersResponsePartnersInner) UnsetPayerWalletableId()`

UnsetPayerWalletableId ensures that no value is present for PayerWalletableId, not even an explicit nil
### GetPhone

`func (o *PartnersResponsePartnersInner) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *PartnersResponsePartnersInner) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *PartnersResponsePartnersInner) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *PartnersResponsePartnersInner) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *PartnersResponsePartnersInner) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *PartnersResponsePartnersInner) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetShortcut1

`func (o *PartnersResponsePartnersInner) GetShortcut1() string`

GetShortcut1 returns the Shortcut1 field if non-nil, zero value otherwise.

### GetShortcut1Ok

`func (o *PartnersResponsePartnersInner) GetShortcut1Ok() (*string, bool)`

GetShortcut1Ok returns a tuple with the Shortcut1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut1

`func (o *PartnersResponsePartnersInner) SetShortcut1(v string)`

SetShortcut1 sets Shortcut1 field to given value.

### HasShortcut1

`func (o *PartnersResponsePartnersInner) HasShortcut1() bool`

HasShortcut1 returns a boolean if a field has been set.

### SetShortcut1Nil

`func (o *PartnersResponsePartnersInner) SetShortcut1Nil(b bool)`

 SetShortcut1Nil sets the value for Shortcut1 to be an explicit nil

### UnsetShortcut1
`func (o *PartnersResponsePartnersInner) UnsetShortcut1()`

UnsetShortcut1 ensures that no value is present for Shortcut1, not even an explicit nil
### GetShortcut2

`func (o *PartnersResponsePartnersInner) GetShortcut2() string`

GetShortcut2 returns the Shortcut2 field if non-nil, zero value otherwise.

### GetShortcut2Ok

`func (o *PartnersResponsePartnersInner) GetShortcut2Ok() (*string, bool)`

GetShortcut2Ok returns a tuple with the Shortcut2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut2

`func (o *PartnersResponsePartnersInner) SetShortcut2(v string)`

SetShortcut2 sets Shortcut2 field to given value.

### HasShortcut2

`func (o *PartnersResponsePartnersInner) HasShortcut2() bool`

HasShortcut2 returns a boolean if a field has been set.

### SetShortcut2Nil

`func (o *PartnersResponsePartnersInner) SetShortcut2Nil(b bool)`

 SetShortcut2Nil sets the value for Shortcut2 to be an explicit nil

### UnsetShortcut2
`func (o *PartnersResponsePartnersInner) UnsetShortcut2()`

UnsetShortcut2 ensures that no value is present for Shortcut2, not even an explicit nil
### GetTransferFeeHandlingSide

`func (o *PartnersResponsePartnersInner) GetTransferFeeHandlingSide() string`

GetTransferFeeHandlingSide returns the TransferFeeHandlingSide field if non-nil, zero value otherwise.

### GetTransferFeeHandlingSideOk

`func (o *PartnersResponsePartnersInner) GetTransferFeeHandlingSideOk() (*string, bool)`

GetTransferFeeHandlingSideOk returns a tuple with the TransferFeeHandlingSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferFeeHandlingSide

`func (o *PartnersResponsePartnersInner) SetTransferFeeHandlingSide(v string)`

SetTransferFeeHandlingSide sets TransferFeeHandlingSide field to given value.

### HasTransferFeeHandlingSide

`func (o *PartnersResponsePartnersInner) HasTransferFeeHandlingSide() bool`

HasTransferFeeHandlingSide returns a boolean if a field has been set.

### GetUpdateDate

`func (o *PartnersResponsePartnersInner) GetUpdateDate() string`

GetUpdateDate returns the UpdateDate field if non-nil, zero value otherwise.

### GetUpdateDateOk

`func (o *PartnersResponsePartnersInner) GetUpdateDateOk() (*string, bool)`

GetUpdateDateOk returns a tuple with the UpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDate

`func (o *PartnersResponsePartnersInner) SetUpdateDate(v string)`

SetUpdateDate sets UpdateDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


