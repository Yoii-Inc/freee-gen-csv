# WalletableCreateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankId** | Pointer to **int64** | 連携サービスID（typeにbank_account、credit_cardを指定する場合は必須） | [optional] 
**CompanyId** | **int64** | 事業所ID | 
**IsAsset** | Pointer to **bool** | 口座を資産口座とするか負債口座とするか（true: 資産口座 (デフォルト), false: 負債口座）&lt;br&gt; bank_idを指定しない場合にのみ使われます。&lt;br&gt; bank_idを指定する場合には資産口座か負債口座かはbank_idに指定したサービスに応じて決定され、is_assetに指定した値は無視されます。  | [optional] 
**Name** | **string** | 口座名 (255文字以内) | 
**Type** | **string** | 口座種別（bank_account : 銀行口座, credit_card : クレジットカード, wallet : その他の決済口座） | 

## Methods

### NewWalletableCreateParams

`func NewWalletableCreateParams(companyId int64, name string, type_ string, ) *WalletableCreateParams`

NewWalletableCreateParams instantiates a new WalletableCreateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletableCreateParamsWithDefaults

`func NewWalletableCreateParamsWithDefaults() *WalletableCreateParams`

NewWalletableCreateParamsWithDefaults instantiates a new WalletableCreateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankId

`func (o *WalletableCreateParams) GetBankId() int64`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *WalletableCreateParams) GetBankIdOk() (*int64, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *WalletableCreateParams) SetBankId(v int64)`

SetBankId sets BankId field to given value.

### HasBankId

`func (o *WalletableCreateParams) HasBankId() bool`

HasBankId returns a boolean if a field has been set.

### GetCompanyId

`func (o *WalletableCreateParams) GetCompanyId() int64`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *WalletableCreateParams) GetCompanyIdOk() (*int64, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *WalletableCreateParams) SetCompanyId(v int64)`

SetCompanyId sets CompanyId field to given value.


### GetIsAsset

`func (o *WalletableCreateParams) GetIsAsset() bool`

GetIsAsset returns the IsAsset field if non-nil, zero value otherwise.

### GetIsAssetOk

`func (o *WalletableCreateParams) GetIsAssetOk() (*bool, bool)`

GetIsAssetOk returns a tuple with the IsAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAsset

`func (o *WalletableCreateParams) SetIsAsset(v bool)`

SetIsAsset sets IsAsset field to given value.

### HasIsAsset

`func (o *WalletableCreateParams) HasIsAsset() bool`

HasIsAsset returns a boolean if a field has been set.

### GetName

`func (o *WalletableCreateParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WalletableCreateParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WalletableCreateParams) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *WalletableCreateParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WalletableCreateParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WalletableCreateParams) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


