# DealRenewsInnerDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountItemId** | **int64** | 勘定科目ID | 
**Amount** | **int64** | 金額（税込で指定してください） | 
**Description** | Pointer to **NullableString** | 備考 | [optional] 
**EntrySide** | **string** | 貸借(貸方: credit, 借方: debit) | 
**Id** | **int64** | +更新の明細行ID | 
**ItemId** | Pointer to **NullableInt64** | 品目ID | [optional] 
**SectionId** | Pointer to **NullableInt64** | 部門ID | [optional] 
**Segment1TagId** | Pointer to **NullableInt64** | セグメント１ID | [optional] 
**Segment2TagId** | Pointer to **NullableInt64** | セグメント２ID | [optional] 
**Segment3TagId** | Pointer to **NullableInt64** | セグメント３ID | [optional] 
**TagIds** | **[]int64** |  | 
**TaxCode** | **int64** | 税区分コード | 
**Vat** | **int64** | 消費税額（指定しない場合は自動で計算されます） | 

## Methods

### NewDealRenewsInnerDetailsInner

`func NewDealRenewsInnerDetailsInner(accountItemId int64, amount int64, entrySide string, id int64, tagIds []int64, taxCode int64, vat int64, ) *DealRenewsInnerDetailsInner`

NewDealRenewsInnerDetailsInner instantiates a new DealRenewsInnerDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealRenewsInnerDetailsInnerWithDefaults

`func NewDealRenewsInnerDetailsInnerWithDefaults() *DealRenewsInnerDetailsInner`

NewDealRenewsInnerDetailsInnerWithDefaults instantiates a new DealRenewsInnerDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountItemId

`func (o *DealRenewsInnerDetailsInner) GetAccountItemId() int64`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *DealRenewsInnerDetailsInner) GetAccountItemIdOk() (*int64, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *DealRenewsInnerDetailsInner) SetAccountItemId(v int64)`

SetAccountItemId sets AccountItemId field to given value.


### GetAmount

`func (o *DealRenewsInnerDetailsInner) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DealRenewsInnerDetailsInner) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DealRenewsInnerDetailsInner) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *DealRenewsInnerDetailsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DealRenewsInnerDetailsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DealRenewsInnerDetailsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DealRenewsInnerDetailsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DealRenewsInnerDetailsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DealRenewsInnerDetailsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEntrySide

`func (o *DealRenewsInnerDetailsInner) GetEntrySide() string`

GetEntrySide returns the EntrySide field if non-nil, zero value otherwise.

### GetEntrySideOk

`func (o *DealRenewsInnerDetailsInner) GetEntrySideOk() (*string, bool)`

GetEntrySideOk returns a tuple with the EntrySide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrySide

`func (o *DealRenewsInnerDetailsInner) SetEntrySide(v string)`

SetEntrySide sets EntrySide field to given value.


### GetId

`func (o *DealRenewsInnerDetailsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DealRenewsInnerDetailsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DealRenewsInnerDetailsInner) SetId(v int64)`

SetId sets Id field to given value.


### GetItemId

`func (o *DealRenewsInnerDetailsInner) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *DealRenewsInnerDetailsInner) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *DealRenewsInnerDetailsInner) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *DealRenewsInnerDetailsInner) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *DealRenewsInnerDetailsInner) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *DealRenewsInnerDetailsInner) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetSectionId

`func (o *DealRenewsInnerDetailsInner) GetSectionId() int64`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *DealRenewsInnerDetailsInner) GetSectionIdOk() (*int64, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *DealRenewsInnerDetailsInner) SetSectionId(v int64)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *DealRenewsInnerDetailsInner) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### SetSectionIdNil

`func (o *DealRenewsInnerDetailsInner) SetSectionIdNil(b bool)`

 SetSectionIdNil sets the value for SectionId to be an explicit nil

### UnsetSectionId
`func (o *DealRenewsInnerDetailsInner) UnsetSectionId()`

UnsetSectionId ensures that no value is present for SectionId, not even an explicit nil
### GetSegment1TagId

`func (o *DealRenewsInnerDetailsInner) GetSegment1TagId() int64`

GetSegment1TagId returns the Segment1TagId field if non-nil, zero value otherwise.

### GetSegment1TagIdOk

`func (o *DealRenewsInnerDetailsInner) GetSegment1TagIdOk() (*int64, bool)`

GetSegment1TagIdOk returns a tuple with the Segment1TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagId

`func (o *DealRenewsInnerDetailsInner) SetSegment1TagId(v int64)`

SetSegment1TagId sets Segment1TagId field to given value.

### HasSegment1TagId

`func (o *DealRenewsInnerDetailsInner) HasSegment1TagId() bool`

HasSegment1TagId returns a boolean if a field has been set.

### SetSegment1TagIdNil

`func (o *DealRenewsInnerDetailsInner) SetSegment1TagIdNil(b bool)`

 SetSegment1TagIdNil sets the value for Segment1TagId to be an explicit nil

### UnsetSegment1TagId
`func (o *DealRenewsInnerDetailsInner) UnsetSegment1TagId()`

UnsetSegment1TagId ensures that no value is present for Segment1TagId, not even an explicit nil
### GetSegment2TagId

`func (o *DealRenewsInnerDetailsInner) GetSegment2TagId() int64`

GetSegment2TagId returns the Segment2TagId field if non-nil, zero value otherwise.

### GetSegment2TagIdOk

`func (o *DealRenewsInnerDetailsInner) GetSegment2TagIdOk() (*int64, bool)`

GetSegment2TagIdOk returns a tuple with the Segment2TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagId

`func (o *DealRenewsInnerDetailsInner) SetSegment2TagId(v int64)`

SetSegment2TagId sets Segment2TagId field to given value.

### HasSegment2TagId

`func (o *DealRenewsInnerDetailsInner) HasSegment2TagId() bool`

HasSegment2TagId returns a boolean if a field has been set.

### SetSegment2TagIdNil

`func (o *DealRenewsInnerDetailsInner) SetSegment2TagIdNil(b bool)`

 SetSegment2TagIdNil sets the value for Segment2TagId to be an explicit nil

### UnsetSegment2TagId
`func (o *DealRenewsInnerDetailsInner) UnsetSegment2TagId()`

UnsetSegment2TagId ensures that no value is present for Segment2TagId, not even an explicit nil
### GetSegment3TagId

`func (o *DealRenewsInnerDetailsInner) GetSegment3TagId() int64`

GetSegment3TagId returns the Segment3TagId field if non-nil, zero value otherwise.

### GetSegment3TagIdOk

`func (o *DealRenewsInnerDetailsInner) GetSegment3TagIdOk() (*int64, bool)`

GetSegment3TagIdOk returns a tuple with the Segment3TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagId

`func (o *DealRenewsInnerDetailsInner) SetSegment3TagId(v int64)`

SetSegment3TagId sets Segment3TagId field to given value.

### HasSegment3TagId

`func (o *DealRenewsInnerDetailsInner) HasSegment3TagId() bool`

HasSegment3TagId returns a boolean if a field has been set.

### SetSegment3TagIdNil

`func (o *DealRenewsInnerDetailsInner) SetSegment3TagIdNil(b bool)`

 SetSegment3TagIdNil sets the value for Segment3TagId to be an explicit nil

### UnsetSegment3TagId
`func (o *DealRenewsInnerDetailsInner) UnsetSegment3TagId()`

UnsetSegment3TagId ensures that no value is present for Segment3TagId, not even an explicit nil
### GetTagIds

`func (o *DealRenewsInnerDetailsInner) GetTagIds() []int64`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *DealRenewsInnerDetailsInner) GetTagIdsOk() (*[]int64, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *DealRenewsInnerDetailsInner) SetTagIds(v []int64)`

SetTagIds sets TagIds field to given value.


### GetTaxCode

`func (o *DealRenewsInnerDetailsInner) GetTaxCode() int64`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *DealRenewsInnerDetailsInner) GetTaxCodeOk() (*int64, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *DealRenewsInnerDetailsInner) SetTaxCode(v int64)`

SetTaxCode sets TaxCode field to given value.


### GetVat

`func (o *DealRenewsInnerDetailsInner) GetVat() int64`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *DealRenewsInnerDetailsInner) GetVatOk() (*int64, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *DealRenewsInnerDetailsInner) SetVat(v int64)`

SetVat sets Vat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


