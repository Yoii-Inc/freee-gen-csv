# DealCreateParamsPaymentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | 支払金額：payments指定時は必須 | 
**Date** | **string** | 支払日：payments指定時は必須 | 
**FromWalletableId** | **int32** | 口座ID（from_walletable_typeがprivate_account_itemの場合は勘定科目ID）：payments指定時は必須 | 
**FromWalletableType** | **string** | 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet, プライベート資金（法人の場合は役員借入金もしくは役員借入金、個人の場合は事業主貸もしくは事業主借）: private_account_item)：payments指定時は必須 | 

## Methods

### NewDealCreateParamsPaymentsInner

`func NewDealCreateParamsPaymentsInner(amount int64, date string, fromWalletableId int32, fromWalletableType string, ) *DealCreateParamsPaymentsInner`

NewDealCreateParamsPaymentsInner instantiates a new DealCreateParamsPaymentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealCreateParamsPaymentsInnerWithDefaults

`func NewDealCreateParamsPaymentsInnerWithDefaults() *DealCreateParamsPaymentsInner`

NewDealCreateParamsPaymentsInnerWithDefaults instantiates a new DealCreateParamsPaymentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DealCreateParamsPaymentsInner) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DealCreateParamsPaymentsInner) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DealCreateParamsPaymentsInner) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetDate

`func (o *DealCreateParamsPaymentsInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DealCreateParamsPaymentsInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DealCreateParamsPaymentsInner) SetDate(v string)`

SetDate sets Date field to given value.


### GetFromWalletableId

`func (o *DealCreateParamsPaymentsInner) GetFromWalletableId() int32`

GetFromWalletableId returns the FromWalletableId field if non-nil, zero value otherwise.

### GetFromWalletableIdOk

`func (o *DealCreateParamsPaymentsInner) GetFromWalletableIdOk() (*int32, bool)`

GetFromWalletableIdOk returns a tuple with the FromWalletableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromWalletableId

`func (o *DealCreateParamsPaymentsInner) SetFromWalletableId(v int32)`

SetFromWalletableId sets FromWalletableId field to given value.


### GetFromWalletableType

`func (o *DealCreateParamsPaymentsInner) GetFromWalletableType() string`

GetFromWalletableType returns the FromWalletableType field if non-nil, zero value otherwise.

### GetFromWalletableTypeOk

`func (o *DealCreateParamsPaymentsInner) GetFromWalletableTypeOk() (*string, bool)`

GetFromWalletableTypeOk returns a tuple with the FromWalletableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromWalletableType

`func (o *DealCreateParamsPaymentsInner) SetFromWalletableType(v string)`

SetFromWalletableType sets FromWalletableType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


