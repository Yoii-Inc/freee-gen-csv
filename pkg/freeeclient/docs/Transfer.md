# Transfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | 金額 | 
**CompanyId** | **int64** | 事業所ID | 
**Date** | **string** | 振替日 (yyyy-mm-dd) | 
**Description** | **string** | 備考 | 
**FromWalletableId** | **int64** | 振替元口座ID | 
**FromWalletableType** | **NullableString** | 振替元口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet) | 
**Id** | **int64** | 取引(振替)ID | 
**ToWalletableId** | **int64** | 振替先口座ID | 
**ToWalletableType** | **NullableString** | 振替先口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet) | 

## Methods

### NewTransfer

`func NewTransfer(amount int64, companyId int64, date string, description string, fromWalletableId int64, fromWalletableType NullableString, id int64, toWalletableId int64, toWalletableType NullableString, ) *Transfer`

NewTransfer instantiates a new Transfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferWithDefaults

`func NewTransferWithDefaults() *Transfer`

NewTransferWithDefaults instantiates a new Transfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Transfer) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Transfer) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Transfer) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCompanyId

`func (o *Transfer) GetCompanyId() int64`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *Transfer) GetCompanyIdOk() (*int64, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *Transfer) SetCompanyId(v int64)`

SetCompanyId sets CompanyId field to given value.


### GetDate

`func (o *Transfer) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Transfer) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Transfer) SetDate(v string)`

SetDate sets Date field to given value.


### GetDescription

`func (o *Transfer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Transfer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Transfer) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFromWalletableId

`func (o *Transfer) GetFromWalletableId() int64`

GetFromWalletableId returns the FromWalletableId field if non-nil, zero value otherwise.

### GetFromWalletableIdOk

`func (o *Transfer) GetFromWalletableIdOk() (*int64, bool)`

GetFromWalletableIdOk returns a tuple with the FromWalletableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromWalletableId

`func (o *Transfer) SetFromWalletableId(v int64)`

SetFromWalletableId sets FromWalletableId field to given value.


### GetFromWalletableType

`func (o *Transfer) GetFromWalletableType() string`

GetFromWalletableType returns the FromWalletableType field if non-nil, zero value otherwise.

### GetFromWalletableTypeOk

`func (o *Transfer) GetFromWalletableTypeOk() (*string, bool)`

GetFromWalletableTypeOk returns a tuple with the FromWalletableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromWalletableType

`func (o *Transfer) SetFromWalletableType(v string)`

SetFromWalletableType sets FromWalletableType field to given value.


### SetFromWalletableTypeNil

`func (o *Transfer) SetFromWalletableTypeNil(b bool)`

 SetFromWalletableTypeNil sets the value for FromWalletableType to be an explicit nil

### UnsetFromWalletableType
`func (o *Transfer) UnsetFromWalletableType()`

UnsetFromWalletableType ensures that no value is present for FromWalletableType, not even an explicit nil
### GetId

`func (o *Transfer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Transfer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Transfer) SetId(v int64)`

SetId sets Id field to given value.


### GetToWalletableId

`func (o *Transfer) GetToWalletableId() int64`

GetToWalletableId returns the ToWalletableId field if non-nil, zero value otherwise.

### GetToWalletableIdOk

`func (o *Transfer) GetToWalletableIdOk() (*int64, bool)`

GetToWalletableIdOk returns a tuple with the ToWalletableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToWalletableId

`func (o *Transfer) SetToWalletableId(v int64)`

SetToWalletableId sets ToWalletableId field to given value.


### GetToWalletableType

`func (o *Transfer) GetToWalletableType() string`

GetToWalletableType returns the ToWalletableType field if non-nil, zero value otherwise.

### GetToWalletableTypeOk

`func (o *Transfer) GetToWalletableTypeOk() (*string, bool)`

GetToWalletableTypeOk returns a tuple with the ToWalletableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToWalletableType

`func (o *Transfer) SetToWalletableType(v string)`

SetToWalletableType sets ToWalletableType field to given value.


### SetToWalletableTypeNil

`func (o *Transfer) SetToWalletableTypeNil(b bool)`

 SetToWalletableTypeNil sets the value for ToWalletableType to be an explicit nil

### UnsetToWalletableType
`func (o *Transfer) UnsetToWalletableType()`

UnsetToWalletableType ensures that no value is present for ToWalletableType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


