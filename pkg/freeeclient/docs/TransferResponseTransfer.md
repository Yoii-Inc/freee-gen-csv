# TransferResponseTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | 金額 | 
**CompanyId** | **int32** | 事業所ID | 
**Date** | **string** | 振替日 (yyyy-mm-dd) | 
**Description** | **string** | 備考 | 
**FromWalletableId** | **int32** | 振替元口座ID | 
**FromWalletableType** | **NullableString** | 振替元口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet) | 
**Id** | **int32** | 取引(振替)ID | 
**ToWalletableId** | **int32** | 振替先口座ID | 
**ToWalletableType** | **NullableString** | 振替先口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet) | 

## Methods

### NewTransferResponseTransfer

`func NewTransferResponseTransfer(amount int64, companyId int32, date string, description string, fromWalletableId int32, fromWalletableType NullableString, id int32, toWalletableId int32, toWalletableType NullableString, ) *TransferResponseTransfer`

NewTransferResponseTransfer instantiates a new TransferResponseTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferResponseTransferWithDefaults

`func NewTransferResponseTransferWithDefaults() *TransferResponseTransfer`

NewTransferResponseTransferWithDefaults instantiates a new TransferResponseTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TransferResponseTransfer) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferResponseTransfer) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferResponseTransfer) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCompanyId

`func (o *TransferResponseTransfer) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *TransferResponseTransfer) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *TransferResponseTransfer) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetDate

`func (o *TransferResponseTransfer) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TransferResponseTransfer) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TransferResponseTransfer) SetDate(v string)`

SetDate sets Date field to given value.


### GetDescription

`func (o *TransferResponseTransfer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransferResponseTransfer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransferResponseTransfer) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFromWalletableId

`func (o *TransferResponseTransfer) GetFromWalletableId() int32`

GetFromWalletableId returns the FromWalletableId field if non-nil, zero value otherwise.

### GetFromWalletableIdOk

`func (o *TransferResponseTransfer) GetFromWalletableIdOk() (*int32, bool)`

GetFromWalletableIdOk returns a tuple with the FromWalletableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromWalletableId

`func (o *TransferResponseTransfer) SetFromWalletableId(v int32)`

SetFromWalletableId sets FromWalletableId field to given value.


### GetFromWalletableType

`func (o *TransferResponseTransfer) GetFromWalletableType() string`

GetFromWalletableType returns the FromWalletableType field if non-nil, zero value otherwise.

### GetFromWalletableTypeOk

`func (o *TransferResponseTransfer) GetFromWalletableTypeOk() (*string, bool)`

GetFromWalletableTypeOk returns a tuple with the FromWalletableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromWalletableType

`func (o *TransferResponseTransfer) SetFromWalletableType(v string)`

SetFromWalletableType sets FromWalletableType field to given value.


### SetFromWalletableTypeNil

`func (o *TransferResponseTransfer) SetFromWalletableTypeNil(b bool)`

 SetFromWalletableTypeNil sets the value for FromWalletableType to be an explicit nil

### UnsetFromWalletableType
`func (o *TransferResponseTransfer) UnsetFromWalletableType()`

UnsetFromWalletableType ensures that no value is present for FromWalletableType, not even an explicit nil
### GetId

`func (o *TransferResponseTransfer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransferResponseTransfer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransferResponseTransfer) SetId(v int32)`

SetId sets Id field to given value.


### GetToWalletableId

`func (o *TransferResponseTransfer) GetToWalletableId() int32`

GetToWalletableId returns the ToWalletableId field if non-nil, zero value otherwise.

### GetToWalletableIdOk

`func (o *TransferResponseTransfer) GetToWalletableIdOk() (*int32, bool)`

GetToWalletableIdOk returns a tuple with the ToWalletableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToWalletableId

`func (o *TransferResponseTransfer) SetToWalletableId(v int32)`

SetToWalletableId sets ToWalletableId field to given value.


### GetToWalletableType

`func (o *TransferResponseTransfer) GetToWalletableType() string`

GetToWalletableType returns the ToWalletableType field if non-nil, zero value otherwise.

### GetToWalletableTypeOk

`func (o *TransferResponseTransfer) GetToWalletableTypeOk() (*string, bool)`

GetToWalletableTypeOk returns a tuple with the ToWalletableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToWalletableType

`func (o *TransferResponseTransfer) SetToWalletableType(v string)`

SetToWalletableType sets ToWalletableType field to given value.


### SetToWalletableTypeNil

`func (o *TransferResponseTransfer) SetToWalletableTypeNil(b bool)`

 SetToWalletableTypeNil sets the value for ToWalletableType to be an explicit nil

### UnsetToWalletableType
`func (o *TransferResponseTransfer) UnsetToWalletableType()`

UnsetToWalletableType ensures that no value is present for ToWalletableType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


