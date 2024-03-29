# RenewCreateParamsDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountItemId** | **int64** | 勘定科目ID | 
**Amount** | **int64** | 取引金額（税込で指定してください） | 
**Description** | Pointer to **string** | 備考 | [optional] 
**ItemId** | Pointer to **int64** | 品目ID | [optional] 
**SectionId** | Pointer to **int64** | 部門ID | [optional] 
**Segment1TagId** | Pointer to **int64** | セグメント１ID | [optional] 
**Segment2TagId** | Pointer to **int64** | セグメント２ID | [optional] 
**Segment3TagId** | Pointer to **int64** | セグメント３ID | [optional] 
**TagIds** | Pointer to **[]int64** | メモタグID | [optional] 
**TaxCode** | **int64** | 税区分コード | 
**Vat** | Pointer to **int64** | 消費税額（指定しない場合は自動で計算されます） | [optional] 

## Methods

### NewRenewCreateParamsDetailsInner

`func NewRenewCreateParamsDetailsInner(accountItemId int64, amount int64, taxCode int64, ) *RenewCreateParamsDetailsInner`

NewRenewCreateParamsDetailsInner instantiates a new RenewCreateParamsDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenewCreateParamsDetailsInnerWithDefaults

`func NewRenewCreateParamsDetailsInnerWithDefaults() *RenewCreateParamsDetailsInner`

NewRenewCreateParamsDetailsInnerWithDefaults instantiates a new RenewCreateParamsDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountItemId

`func (o *RenewCreateParamsDetailsInner) GetAccountItemId() int64`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *RenewCreateParamsDetailsInner) GetAccountItemIdOk() (*int64, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *RenewCreateParamsDetailsInner) SetAccountItemId(v int64)`

SetAccountItemId sets AccountItemId field to given value.


### GetAmount

`func (o *RenewCreateParamsDetailsInner) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RenewCreateParamsDetailsInner) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RenewCreateParamsDetailsInner) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *RenewCreateParamsDetailsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RenewCreateParamsDetailsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RenewCreateParamsDetailsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RenewCreateParamsDetailsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetItemId

`func (o *RenewCreateParamsDetailsInner) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *RenewCreateParamsDetailsInner) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *RenewCreateParamsDetailsInner) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *RenewCreateParamsDetailsInner) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetSectionId

`func (o *RenewCreateParamsDetailsInner) GetSectionId() int64`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *RenewCreateParamsDetailsInner) GetSectionIdOk() (*int64, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *RenewCreateParamsDetailsInner) SetSectionId(v int64)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *RenewCreateParamsDetailsInner) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### GetSegment1TagId

`func (o *RenewCreateParamsDetailsInner) GetSegment1TagId() int64`

GetSegment1TagId returns the Segment1TagId field if non-nil, zero value otherwise.

### GetSegment1TagIdOk

`func (o *RenewCreateParamsDetailsInner) GetSegment1TagIdOk() (*int64, bool)`

GetSegment1TagIdOk returns a tuple with the Segment1TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagId

`func (o *RenewCreateParamsDetailsInner) SetSegment1TagId(v int64)`

SetSegment1TagId sets Segment1TagId field to given value.

### HasSegment1TagId

`func (o *RenewCreateParamsDetailsInner) HasSegment1TagId() bool`

HasSegment1TagId returns a boolean if a field has been set.

### GetSegment2TagId

`func (o *RenewCreateParamsDetailsInner) GetSegment2TagId() int64`

GetSegment2TagId returns the Segment2TagId field if non-nil, zero value otherwise.

### GetSegment2TagIdOk

`func (o *RenewCreateParamsDetailsInner) GetSegment2TagIdOk() (*int64, bool)`

GetSegment2TagIdOk returns a tuple with the Segment2TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagId

`func (o *RenewCreateParamsDetailsInner) SetSegment2TagId(v int64)`

SetSegment2TagId sets Segment2TagId field to given value.

### HasSegment2TagId

`func (o *RenewCreateParamsDetailsInner) HasSegment2TagId() bool`

HasSegment2TagId returns a boolean if a field has been set.

### GetSegment3TagId

`func (o *RenewCreateParamsDetailsInner) GetSegment3TagId() int64`

GetSegment3TagId returns the Segment3TagId field if non-nil, zero value otherwise.

### GetSegment3TagIdOk

`func (o *RenewCreateParamsDetailsInner) GetSegment3TagIdOk() (*int64, bool)`

GetSegment3TagIdOk returns a tuple with the Segment3TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagId

`func (o *RenewCreateParamsDetailsInner) SetSegment3TagId(v int64)`

SetSegment3TagId sets Segment3TagId field to given value.

### HasSegment3TagId

`func (o *RenewCreateParamsDetailsInner) HasSegment3TagId() bool`

HasSegment3TagId returns a boolean if a field has been set.

### GetTagIds

`func (o *RenewCreateParamsDetailsInner) GetTagIds() []int64`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *RenewCreateParamsDetailsInner) GetTagIdsOk() (*[]int64, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *RenewCreateParamsDetailsInner) SetTagIds(v []int64)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *RenewCreateParamsDetailsInner) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTaxCode

`func (o *RenewCreateParamsDetailsInner) GetTaxCode() int64`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *RenewCreateParamsDetailsInner) GetTaxCodeOk() (*int64, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *RenewCreateParamsDetailsInner) SetTaxCode(v int64)`

SetTaxCode sets TaxCode field to given value.


### GetVat

`func (o *RenewCreateParamsDetailsInner) GetVat() int64`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *RenewCreateParamsDetailsInner) GetVatOk() (*int64, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *RenewCreateParamsDetailsInner) SetVat(v int64)`

SetVat sets Vat field to given value.

### HasVat

`func (o *RenewCreateParamsDetailsInner) HasVat() bool`

HasVat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


