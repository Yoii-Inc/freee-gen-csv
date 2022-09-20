# DealPaymentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | 支払金額 | 
**Date** | **string** | 支払日 | 
**FromWalletableId** | Pointer to **int32** | 口座ID（from_walletable_typeがprivate_account_itemの場合は勘定科目ID） | [optional] 
**FromWalletableType** | Pointer to **string** | 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet, プライベート資金（法人の場合は役員借入金もしくは役員借入金、個人の場合は事業主貸もしくは事業主借）: private_account_item) | [optional] 
**Id** | **int64** | 取引行ID | 

## Methods

### NewDealPaymentsInner

`func NewDealPaymentsInner(amount int64, date string, id int64, ) *DealPaymentsInner`

NewDealPaymentsInner instantiates a new DealPaymentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealPaymentsInnerWithDefaults

`func NewDealPaymentsInnerWithDefaults() *DealPaymentsInner`

NewDealPaymentsInnerWithDefaults instantiates a new DealPaymentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DealPaymentsInner) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DealPaymentsInner) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DealPaymentsInner) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetDate

`func (o *DealPaymentsInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DealPaymentsInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DealPaymentsInner) SetDate(v string)`

SetDate sets Date field to given value.


### GetFromWalletableId

`func (o *DealPaymentsInner) GetFromWalletableId() int32`

GetFromWalletableId returns the FromWalletableId field if non-nil, zero value otherwise.

### GetFromWalletableIdOk

`func (o *DealPaymentsInner) GetFromWalletableIdOk() (*int32, bool)`

GetFromWalletableIdOk returns a tuple with the FromWalletableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromWalletableId

`func (o *DealPaymentsInner) SetFromWalletableId(v int32)`

SetFromWalletableId sets FromWalletableId field to given value.

### HasFromWalletableId

`func (o *DealPaymentsInner) HasFromWalletableId() bool`

HasFromWalletableId returns a boolean if a field has been set.

### GetFromWalletableType

`func (o *DealPaymentsInner) GetFromWalletableType() string`

GetFromWalletableType returns the FromWalletableType field if non-nil, zero value otherwise.

### GetFromWalletableTypeOk

`func (o *DealPaymentsInner) GetFromWalletableTypeOk() (*string, bool)`

GetFromWalletableTypeOk returns a tuple with the FromWalletableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromWalletableType

`func (o *DealPaymentsInner) SetFromWalletableType(v string)`

SetFromWalletableType sets FromWalletableType field to given value.

### HasFromWalletableType

`func (o *DealPaymentsInner) HasFromWalletableType() bool`

HasFromWalletableType returns a boolean if a field has been set.

### GetId

`func (o *DealPaymentsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DealPaymentsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DealPaymentsInner) SetId(v int64)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


