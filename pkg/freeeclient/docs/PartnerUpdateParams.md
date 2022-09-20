# PartnerUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressAttributes** | Pointer to [**PartnerCreateParamsAddressAttributes**](PartnerCreateParamsAddressAttributes.md) |  | [optional] 
**CompanyId** | **int32** | 事業所ID | 
**ContactName** | Pointer to **string** | 担当者 氏名 (255文字以内) | [optional] 
**CountryCode** | Pointer to **string** | 地域（JP: 国内、ZZ:国外）、指定しない場合JPになります。 | [optional] 
**DefaultTitle** | Pointer to **string** | 敬称（御中、様、(空白)の3つから選択） | [optional] 
**Email** | Pointer to **string** | 担当者 メールアドレス (255文字以内) | [optional] 
**InvoicePaymentTermAttributes** | Pointer to [**NullablePartnerUpdateParamsInvoicePaymentTermAttributes**](PartnerUpdateParamsInvoicePaymentTermAttributes.md) |  | [optional] 
**LongName** | Pointer to **string** | 正式名称（255文字以内） | [optional] 
**Name** | **string** | 取引先名 (255文字以内) | 
**NameKana** | Pointer to **string** | カナ名称（255文字以内） | [optional] 
**OrgCode** | Pointer to **NullableInt32** | 事業所種別（null: 未設定、1: 法人、2: 個人） | [optional] 
**PartnerBankAccountAttributes** | Pointer to [**PartnerCreateParamsPartnerBankAccountAttributes**](PartnerCreateParamsPartnerBankAccountAttributes.md) |  | [optional] 
**PartnerDocSettingAttributes** | Pointer to [**PartnerCreateParamsPartnerDocSettingAttributes**](PartnerCreateParamsPartnerDocSettingAttributes.md) |  | [optional] 
**PayerWalletableId** | Pointer to **NullableInt32** | 振込元口座ID（一括振込ファイル用）:（walletableのtypeが&#39;bank_account&#39;のidのみ指定できます。また、未設定にする場合は、nullを指定してください。） | [optional] 
**PaymentTermAttributes** | Pointer to [**NullablePartnerUpdateParamsPaymentTermAttributes**](PartnerUpdateParamsPaymentTermAttributes.md) |  | [optional] 
**Phone** | Pointer to **string** | 電話番号 | [optional] 
**Shortcut1** | Pointer to **string** | ショートカット１ (255文字以内) | [optional] 
**Shortcut2** | Pointer to **string** | ショートカット２ (255文字以内) | [optional] 
**TransferFeeHandlingSide** | Pointer to **string** | 振込手数料負担（一括振込ファイル用）: (振込元(当方): payer, 振込先(先方): payee)、指定しない場合payerになります。 | [optional] 

## Methods

### NewPartnerUpdateParams

`func NewPartnerUpdateParams(companyId int32, name string, ) *PartnerUpdateParams`

NewPartnerUpdateParams instantiates a new PartnerUpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerUpdateParamsWithDefaults

`func NewPartnerUpdateParamsWithDefaults() *PartnerUpdateParams`

NewPartnerUpdateParamsWithDefaults instantiates a new PartnerUpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressAttributes

`func (o *PartnerUpdateParams) GetAddressAttributes() PartnerCreateParamsAddressAttributes`

GetAddressAttributes returns the AddressAttributes field if non-nil, zero value otherwise.

### GetAddressAttributesOk

`func (o *PartnerUpdateParams) GetAddressAttributesOk() (*PartnerCreateParamsAddressAttributes, bool)`

GetAddressAttributesOk returns a tuple with the AddressAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressAttributes

`func (o *PartnerUpdateParams) SetAddressAttributes(v PartnerCreateParamsAddressAttributes)`

SetAddressAttributes sets AddressAttributes field to given value.

### HasAddressAttributes

`func (o *PartnerUpdateParams) HasAddressAttributes() bool`

HasAddressAttributes returns a boolean if a field has been set.

### GetCompanyId

`func (o *PartnerUpdateParams) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *PartnerUpdateParams) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *PartnerUpdateParams) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetContactName

`func (o *PartnerUpdateParams) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *PartnerUpdateParams) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *PartnerUpdateParams) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *PartnerUpdateParams) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### GetCountryCode

`func (o *PartnerUpdateParams) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *PartnerUpdateParams) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *PartnerUpdateParams) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *PartnerUpdateParams) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetDefaultTitle

`func (o *PartnerUpdateParams) GetDefaultTitle() string`

GetDefaultTitle returns the DefaultTitle field if non-nil, zero value otherwise.

### GetDefaultTitleOk

`func (o *PartnerUpdateParams) GetDefaultTitleOk() (*string, bool)`

GetDefaultTitleOk returns a tuple with the DefaultTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTitle

`func (o *PartnerUpdateParams) SetDefaultTitle(v string)`

SetDefaultTitle sets DefaultTitle field to given value.

### HasDefaultTitle

`func (o *PartnerUpdateParams) HasDefaultTitle() bool`

HasDefaultTitle returns a boolean if a field has been set.

### GetEmail

`func (o *PartnerUpdateParams) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PartnerUpdateParams) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PartnerUpdateParams) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PartnerUpdateParams) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetInvoicePaymentTermAttributes

`func (o *PartnerUpdateParams) GetInvoicePaymentTermAttributes() PartnerUpdateParamsInvoicePaymentTermAttributes`

GetInvoicePaymentTermAttributes returns the InvoicePaymentTermAttributes field if non-nil, zero value otherwise.

### GetInvoicePaymentTermAttributesOk

`func (o *PartnerUpdateParams) GetInvoicePaymentTermAttributesOk() (*PartnerUpdateParamsInvoicePaymentTermAttributes, bool)`

GetInvoicePaymentTermAttributesOk returns a tuple with the InvoicePaymentTermAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePaymentTermAttributes

`func (o *PartnerUpdateParams) SetInvoicePaymentTermAttributes(v PartnerUpdateParamsInvoicePaymentTermAttributes)`

SetInvoicePaymentTermAttributes sets InvoicePaymentTermAttributes field to given value.

### HasInvoicePaymentTermAttributes

`func (o *PartnerUpdateParams) HasInvoicePaymentTermAttributes() bool`

HasInvoicePaymentTermAttributes returns a boolean if a field has been set.

### SetInvoicePaymentTermAttributesNil

`func (o *PartnerUpdateParams) SetInvoicePaymentTermAttributesNil(b bool)`

 SetInvoicePaymentTermAttributesNil sets the value for InvoicePaymentTermAttributes to be an explicit nil

### UnsetInvoicePaymentTermAttributes
`func (o *PartnerUpdateParams) UnsetInvoicePaymentTermAttributes()`

UnsetInvoicePaymentTermAttributes ensures that no value is present for InvoicePaymentTermAttributes, not even an explicit nil
### GetLongName

`func (o *PartnerUpdateParams) GetLongName() string`

GetLongName returns the LongName field if non-nil, zero value otherwise.

### GetLongNameOk

`func (o *PartnerUpdateParams) GetLongNameOk() (*string, bool)`

GetLongNameOk returns a tuple with the LongName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongName

`func (o *PartnerUpdateParams) SetLongName(v string)`

SetLongName sets LongName field to given value.

### HasLongName

`func (o *PartnerUpdateParams) HasLongName() bool`

HasLongName returns a boolean if a field has been set.

### GetName

`func (o *PartnerUpdateParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartnerUpdateParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartnerUpdateParams) SetName(v string)`

SetName sets Name field to given value.


### GetNameKana

`func (o *PartnerUpdateParams) GetNameKana() string`

GetNameKana returns the NameKana field if non-nil, zero value otherwise.

### GetNameKanaOk

`func (o *PartnerUpdateParams) GetNameKanaOk() (*string, bool)`

GetNameKanaOk returns a tuple with the NameKana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameKana

`func (o *PartnerUpdateParams) SetNameKana(v string)`

SetNameKana sets NameKana field to given value.

### HasNameKana

`func (o *PartnerUpdateParams) HasNameKana() bool`

HasNameKana returns a boolean if a field has been set.

### GetOrgCode

`func (o *PartnerUpdateParams) GetOrgCode() int32`

GetOrgCode returns the OrgCode field if non-nil, zero value otherwise.

### GetOrgCodeOk

`func (o *PartnerUpdateParams) GetOrgCodeOk() (*int32, bool)`

GetOrgCodeOk returns a tuple with the OrgCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgCode

`func (o *PartnerUpdateParams) SetOrgCode(v int32)`

SetOrgCode sets OrgCode field to given value.

### HasOrgCode

`func (o *PartnerUpdateParams) HasOrgCode() bool`

HasOrgCode returns a boolean if a field has been set.

### SetOrgCodeNil

`func (o *PartnerUpdateParams) SetOrgCodeNil(b bool)`

 SetOrgCodeNil sets the value for OrgCode to be an explicit nil

### UnsetOrgCode
`func (o *PartnerUpdateParams) UnsetOrgCode()`

UnsetOrgCode ensures that no value is present for OrgCode, not even an explicit nil
### GetPartnerBankAccountAttributes

`func (o *PartnerUpdateParams) GetPartnerBankAccountAttributes() PartnerCreateParamsPartnerBankAccountAttributes`

GetPartnerBankAccountAttributes returns the PartnerBankAccountAttributes field if non-nil, zero value otherwise.

### GetPartnerBankAccountAttributesOk

`func (o *PartnerUpdateParams) GetPartnerBankAccountAttributesOk() (*PartnerCreateParamsPartnerBankAccountAttributes, bool)`

GetPartnerBankAccountAttributesOk returns a tuple with the PartnerBankAccountAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerBankAccountAttributes

`func (o *PartnerUpdateParams) SetPartnerBankAccountAttributes(v PartnerCreateParamsPartnerBankAccountAttributes)`

SetPartnerBankAccountAttributes sets PartnerBankAccountAttributes field to given value.

### HasPartnerBankAccountAttributes

`func (o *PartnerUpdateParams) HasPartnerBankAccountAttributes() bool`

HasPartnerBankAccountAttributes returns a boolean if a field has been set.

### GetPartnerDocSettingAttributes

`func (o *PartnerUpdateParams) GetPartnerDocSettingAttributes() PartnerCreateParamsPartnerDocSettingAttributes`

GetPartnerDocSettingAttributes returns the PartnerDocSettingAttributes field if non-nil, zero value otherwise.

### GetPartnerDocSettingAttributesOk

`func (o *PartnerUpdateParams) GetPartnerDocSettingAttributesOk() (*PartnerCreateParamsPartnerDocSettingAttributes, bool)`

GetPartnerDocSettingAttributesOk returns a tuple with the PartnerDocSettingAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerDocSettingAttributes

`func (o *PartnerUpdateParams) SetPartnerDocSettingAttributes(v PartnerCreateParamsPartnerDocSettingAttributes)`

SetPartnerDocSettingAttributes sets PartnerDocSettingAttributes field to given value.

### HasPartnerDocSettingAttributes

`func (o *PartnerUpdateParams) HasPartnerDocSettingAttributes() bool`

HasPartnerDocSettingAttributes returns a boolean if a field has been set.

### GetPayerWalletableId

`func (o *PartnerUpdateParams) GetPayerWalletableId() int32`

GetPayerWalletableId returns the PayerWalletableId field if non-nil, zero value otherwise.

### GetPayerWalletableIdOk

`func (o *PartnerUpdateParams) GetPayerWalletableIdOk() (*int32, bool)`

GetPayerWalletableIdOk returns a tuple with the PayerWalletableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerWalletableId

`func (o *PartnerUpdateParams) SetPayerWalletableId(v int32)`

SetPayerWalletableId sets PayerWalletableId field to given value.

### HasPayerWalletableId

`func (o *PartnerUpdateParams) HasPayerWalletableId() bool`

HasPayerWalletableId returns a boolean if a field has been set.

### SetPayerWalletableIdNil

`func (o *PartnerUpdateParams) SetPayerWalletableIdNil(b bool)`

 SetPayerWalletableIdNil sets the value for PayerWalletableId to be an explicit nil

### UnsetPayerWalletableId
`func (o *PartnerUpdateParams) UnsetPayerWalletableId()`

UnsetPayerWalletableId ensures that no value is present for PayerWalletableId, not even an explicit nil
### GetPaymentTermAttributes

`func (o *PartnerUpdateParams) GetPaymentTermAttributes() PartnerUpdateParamsPaymentTermAttributes`

GetPaymentTermAttributes returns the PaymentTermAttributes field if non-nil, zero value otherwise.

### GetPaymentTermAttributesOk

`func (o *PartnerUpdateParams) GetPaymentTermAttributesOk() (*PartnerUpdateParamsPaymentTermAttributes, bool)`

GetPaymentTermAttributesOk returns a tuple with the PaymentTermAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTermAttributes

`func (o *PartnerUpdateParams) SetPaymentTermAttributes(v PartnerUpdateParamsPaymentTermAttributes)`

SetPaymentTermAttributes sets PaymentTermAttributes field to given value.

### HasPaymentTermAttributes

`func (o *PartnerUpdateParams) HasPaymentTermAttributes() bool`

HasPaymentTermAttributes returns a boolean if a field has been set.

### SetPaymentTermAttributesNil

`func (o *PartnerUpdateParams) SetPaymentTermAttributesNil(b bool)`

 SetPaymentTermAttributesNil sets the value for PaymentTermAttributes to be an explicit nil

### UnsetPaymentTermAttributes
`func (o *PartnerUpdateParams) UnsetPaymentTermAttributes()`

UnsetPaymentTermAttributes ensures that no value is present for PaymentTermAttributes, not even an explicit nil
### GetPhone

`func (o *PartnerUpdateParams) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *PartnerUpdateParams) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *PartnerUpdateParams) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *PartnerUpdateParams) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetShortcut1

`func (o *PartnerUpdateParams) GetShortcut1() string`

GetShortcut1 returns the Shortcut1 field if non-nil, zero value otherwise.

### GetShortcut1Ok

`func (o *PartnerUpdateParams) GetShortcut1Ok() (*string, bool)`

GetShortcut1Ok returns a tuple with the Shortcut1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut1

`func (o *PartnerUpdateParams) SetShortcut1(v string)`

SetShortcut1 sets Shortcut1 field to given value.

### HasShortcut1

`func (o *PartnerUpdateParams) HasShortcut1() bool`

HasShortcut1 returns a boolean if a field has been set.

### GetShortcut2

`func (o *PartnerUpdateParams) GetShortcut2() string`

GetShortcut2 returns the Shortcut2 field if non-nil, zero value otherwise.

### GetShortcut2Ok

`func (o *PartnerUpdateParams) GetShortcut2Ok() (*string, bool)`

GetShortcut2Ok returns a tuple with the Shortcut2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut2

`func (o *PartnerUpdateParams) SetShortcut2(v string)`

SetShortcut2 sets Shortcut2 field to given value.

### HasShortcut2

`func (o *PartnerUpdateParams) HasShortcut2() bool`

HasShortcut2 returns a boolean if a field has been set.

### GetTransferFeeHandlingSide

`func (o *PartnerUpdateParams) GetTransferFeeHandlingSide() string`

GetTransferFeeHandlingSide returns the TransferFeeHandlingSide field if non-nil, zero value otherwise.

### GetTransferFeeHandlingSideOk

`func (o *PartnerUpdateParams) GetTransferFeeHandlingSideOk() (*string, bool)`

GetTransferFeeHandlingSideOk returns a tuple with the TransferFeeHandlingSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferFeeHandlingSide

`func (o *PartnerUpdateParams) SetTransferFeeHandlingSide(v string)`

SetTransferFeeHandlingSide sets TransferFeeHandlingSide field to given value.

### HasTransferFeeHandlingSide

`func (o *PartnerUpdateParams) HasTransferFeeHandlingSide() bool`

HasTransferFeeHandlingSide returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


