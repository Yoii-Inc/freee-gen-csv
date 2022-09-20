# DealDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountItemId** | **int32** | 勘定科目ID | 
**Amount** | **int64** | 取引金額 | 
**Description** | Pointer to **string** | 備考 | [optional] 
**EntrySide** | **string** | 貸借（貸方: credit, 借方: debit） | 
**Id** | **int64** | 取引行ID | 
**ItemId** | Pointer to **NullableInt32** | 品目ID | [optional] 
**SectionId** | Pointer to **NullableInt32** | 部門ID | [optional] 
**Segment1TagId** | Pointer to **NullableInt64** | セグメント１ID | [optional] 
**Segment2TagId** | Pointer to **NullableInt64** | セグメント２ID | [optional] 
**Segment3TagId** | Pointer to **NullableInt64** | セグメント３ID | [optional] 
**TagIds** | Pointer to **[]int32** | メモタグID | [optional] 
**TaxCode** | **int32** | 税区分コード | 
**Vat** | **int32** | 消費税額 | 

## Methods

### NewDealDetailsInner

`func NewDealDetailsInner(accountItemId int32, amount int64, entrySide string, id int64, taxCode int32, vat int32, ) *DealDetailsInner`

NewDealDetailsInner instantiates a new DealDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealDetailsInnerWithDefaults

`func NewDealDetailsInnerWithDefaults() *DealDetailsInner`

NewDealDetailsInnerWithDefaults instantiates a new DealDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountItemId

`func (o *DealDetailsInner) GetAccountItemId() int32`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *DealDetailsInner) GetAccountItemIdOk() (*int32, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *DealDetailsInner) SetAccountItemId(v int32)`

SetAccountItemId sets AccountItemId field to given value.


### GetAmount

`func (o *DealDetailsInner) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DealDetailsInner) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DealDetailsInner) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *DealDetailsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DealDetailsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DealDetailsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DealDetailsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntrySide

`func (o *DealDetailsInner) GetEntrySide() string`

GetEntrySide returns the EntrySide field if non-nil, zero value otherwise.

### GetEntrySideOk

`func (o *DealDetailsInner) GetEntrySideOk() (*string, bool)`

GetEntrySideOk returns a tuple with the EntrySide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrySide

`func (o *DealDetailsInner) SetEntrySide(v string)`

SetEntrySide sets EntrySide field to given value.


### GetId

`func (o *DealDetailsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DealDetailsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DealDetailsInner) SetId(v int64)`

SetId sets Id field to given value.


### GetItemId

`func (o *DealDetailsInner) GetItemId() int32`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *DealDetailsInner) GetItemIdOk() (*int32, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *DealDetailsInner) SetItemId(v int32)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *DealDetailsInner) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *DealDetailsInner) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *DealDetailsInner) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetSectionId

`func (o *DealDetailsInner) GetSectionId() int32`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *DealDetailsInner) GetSectionIdOk() (*int32, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *DealDetailsInner) SetSectionId(v int32)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *DealDetailsInner) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### SetSectionIdNil

`func (o *DealDetailsInner) SetSectionIdNil(b bool)`

 SetSectionIdNil sets the value for SectionId to be an explicit nil

### UnsetSectionId
`func (o *DealDetailsInner) UnsetSectionId()`

UnsetSectionId ensures that no value is present for SectionId, not even an explicit nil
### GetSegment1TagId

`func (o *DealDetailsInner) GetSegment1TagId() int64`

GetSegment1TagId returns the Segment1TagId field if non-nil, zero value otherwise.

### GetSegment1TagIdOk

`func (o *DealDetailsInner) GetSegment1TagIdOk() (*int64, bool)`

GetSegment1TagIdOk returns a tuple with the Segment1TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagId

`func (o *DealDetailsInner) SetSegment1TagId(v int64)`

SetSegment1TagId sets Segment1TagId field to given value.

### HasSegment1TagId

`func (o *DealDetailsInner) HasSegment1TagId() bool`

HasSegment1TagId returns a boolean if a field has been set.

### SetSegment1TagIdNil

`func (o *DealDetailsInner) SetSegment1TagIdNil(b bool)`

 SetSegment1TagIdNil sets the value for Segment1TagId to be an explicit nil

### UnsetSegment1TagId
`func (o *DealDetailsInner) UnsetSegment1TagId()`

UnsetSegment1TagId ensures that no value is present for Segment1TagId, not even an explicit nil
### GetSegment2TagId

`func (o *DealDetailsInner) GetSegment2TagId() int64`

GetSegment2TagId returns the Segment2TagId field if non-nil, zero value otherwise.

### GetSegment2TagIdOk

`func (o *DealDetailsInner) GetSegment2TagIdOk() (*int64, bool)`

GetSegment2TagIdOk returns a tuple with the Segment2TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagId

`func (o *DealDetailsInner) SetSegment2TagId(v int64)`

SetSegment2TagId sets Segment2TagId field to given value.

### HasSegment2TagId

`func (o *DealDetailsInner) HasSegment2TagId() bool`

HasSegment2TagId returns a boolean if a field has been set.

### SetSegment2TagIdNil

`func (o *DealDetailsInner) SetSegment2TagIdNil(b bool)`

 SetSegment2TagIdNil sets the value for Segment2TagId to be an explicit nil

### UnsetSegment2TagId
`func (o *DealDetailsInner) UnsetSegment2TagId()`

UnsetSegment2TagId ensures that no value is present for Segment2TagId, not even an explicit nil
### GetSegment3TagId

`func (o *DealDetailsInner) GetSegment3TagId() int64`

GetSegment3TagId returns the Segment3TagId field if non-nil, zero value otherwise.

### GetSegment3TagIdOk

`func (o *DealDetailsInner) GetSegment3TagIdOk() (*int64, bool)`

GetSegment3TagIdOk returns a tuple with the Segment3TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagId

`func (o *DealDetailsInner) SetSegment3TagId(v int64)`

SetSegment3TagId sets Segment3TagId field to given value.

### HasSegment3TagId

`func (o *DealDetailsInner) HasSegment3TagId() bool`

HasSegment3TagId returns a boolean if a field has been set.

### SetSegment3TagIdNil

`func (o *DealDetailsInner) SetSegment3TagIdNil(b bool)`

 SetSegment3TagIdNil sets the value for Segment3TagId to be an explicit nil

### UnsetSegment3TagId
`func (o *DealDetailsInner) UnsetSegment3TagId()`

UnsetSegment3TagId ensures that no value is present for Segment3TagId, not even an explicit nil
### GetTagIds

`func (o *DealDetailsInner) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *DealDetailsInner) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *DealDetailsInner) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *DealDetailsInner) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTaxCode

`func (o *DealDetailsInner) GetTaxCode() int32`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *DealDetailsInner) GetTaxCodeOk() (*int32, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *DealDetailsInner) SetTaxCode(v int32)`

SetTaxCode sets TaxCode field to given value.


### GetVat

`func (o *DealDetailsInner) GetVat() int32`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *DealDetailsInner) GetVatOk() (*int32, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *DealDetailsInner) SetVat(v int32)`

SetVat sets Vat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


