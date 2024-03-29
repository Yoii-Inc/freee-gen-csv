# PaymentRequestResponsePaymentRequestPaymentRequestLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountItemId** | **NullableInt64** | 勘定科目ID | 
**Amount** | **int64** | 金額 | 
**Description** | **string** | 内容 | 
**Id** | **int64** | 支払依頼の項目行ID | 
**ItemId** | **NullableInt64** | 品目ID | 
**LineType** | **string** | 行の種類 (deal_line: 支払依頼, withholding_tax: 源泉徴収税) | 
**SectionId** | **NullableInt64** | 部門ID | 
**Segment1TagId** | Pointer to **NullableInt64** | セグメント１ID。セグメント１が使用可能なプランの時のみレスポンスに含まれます。 | [optional] 
**Segment2TagId** | Pointer to **NullableInt64** | セグメント２ID。セグメント２が使用可能なプランの時のみレスポンスに含まれます。 | [optional] 
**Segment3TagId** | Pointer to **NullableInt64** | セグメント３ID。セグメント３が使用可能なプランの時のみレスポンスに含まれます。 | [optional] 
**TagIds** | **[]int64** | メモタグID | 
**TaxCode** | **NullableInt64** | 税区分コード | 

## Methods

### NewPaymentRequestResponsePaymentRequestPaymentRequestLinesInner

`func NewPaymentRequestResponsePaymentRequestPaymentRequestLinesInner(accountItemId NullableInt64, amount int64, description string, id int64, itemId NullableInt64, lineType string, sectionId NullableInt64, tagIds []int64, taxCode NullableInt64, ) *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner`

NewPaymentRequestResponsePaymentRequestPaymentRequestLinesInner instantiates a new PaymentRequestResponsePaymentRequestPaymentRequestLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRequestResponsePaymentRequestPaymentRequestLinesInnerWithDefaults

`func NewPaymentRequestResponsePaymentRequestPaymentRequestLinesInnerWithDefaults() *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner`

NewPaymentRequestResponsePaymentRequestPaymentRequestLinesInnerWithDefaults instantiates a new PaymentRequestResponsePaymentRequestPaymentRequestLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountItemId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) GetAccountItemId() int64`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) GetAccountItemIdOk() (*int64, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) SetAccountItemId(v int64)`

SetAccountItemId sets AccountItemId field to given value.


### SetAccountItemIdNil

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) SetAccountItemIdNil(b bool)`

 SetAccountItemIdNil sets the value for AccountItemId to be an explicit nil

### UnsetAccountItemId
`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) UnsetAccountItemId()`

UnsetAccountItemId ensures that no value is present for AccountItemId, not even an explicit nil
### GetAmount

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) SetId(v int64)`

SetId sets Id field to given value.


### GetItemId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) SetItemId(v int64)`

SetItemId sets ItemId field to given value.


### SetItemIdNil

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetLineType

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) GetLineType() string`

GetLineType returns the LineType field if non-nil, zero value otherwise.

### GetLineTypeOk

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) GetLineTypeOk() (*string, bool)`

GetLineTypeOk returns a tuple with the LineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineType

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) SetLineType(v string)`

SetLineType sets LineType field to given value.


### GetSectionId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) GetSectionId() int64`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) GetSectionIdOk() (*int64, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) SetSectionId(v int64)`

SetSectionId sets SectionId field to given value.


### SetSectionIdNil

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) SetSectionIdNil(b bool)`

 SetSectionIdNil sets the value for SectionId to be an explicit nil

### UnsetSectionId
`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) UnsetSectionId()`

UnsetSectionId ensures that no value is present for SectionId, not even an explicit nil
### GetSegment1TagId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) GetSegment1TagId() int64`

GetSegment1TagId returns the Segment1TagId field if non-nil, zero value otherwise.

### GetSegment1TagIdOk

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) GetSegment1TagIdOk() (*int64, bool)`

GetSegment1TagIdOk returns a tuple with the Segment1TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) SetSegment1TagId(v int64)`

SetSegment1TagId sets Segment1TagId field to given value.

### HasSegment1TagId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) HasSegment1TagId() bool`

HasSegment1TagId returns a boolean if a field has been set.

### SetSegment1TagIdNil

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) SetSegment1TagIdNil(b bool)`

 SetSegment1TagIdNil sets the value for Segment1TagId to be an explicit nil

### UnsetSegment1TagId
`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) UnsetSegment1TagId()`

UnsetSegment1TagId ensures that no value is present for Segment1TagId, not even an explicit nil
### GetSegment2TagId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) GetSegment2TagId() int64`

GetSegment2TagId returns the Segment2TagId field if non-nil, zero value otherwise.

### GetSegment2TagIdOk

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) GetSegment2TagIdOk() (*int64, bool)`

GetSegment2TagIdOk returns a tuple with the Segment2TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) SetSegment2TagId(v int64)`

SetSegment2TagId sets Segment2TagId field to given value.

### HasSegment2TagId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) HasSegment2TagId() bool`

HasSegment2TagId returns a boolean if a field has been set.

### SetSegment2TagIdNil

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) SetSegment2TagIdNil(b bool)`

 SetSegment2TagIdNil sets the value for Segment2TagId to be an explicit nil

### UnsetSegment2TagId
`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) UnsetSegment2TagId()`

UnsetSegment2TagId ensures that no value is present for Segment2TagId, not even an explicit nil
### GetSegment3TagId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) GetSegment3TagId() int64`

GetSegment3TagId returns the Segment3TagId field if non-nil, zero value otherwise.

### GetSegment3TagIdOk

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) GetSegment3TagIdOk() (*int64, bool)`

GetSegment3TagIdOk returns a tuple with the Segment3TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) SetSegment3TagId(v int64)`

SetSegment3TagId sets Segment3TagId field to given value.

### HasSegment3TagId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) HasSegment3TagId() bool`

HasSegment3TagId returns a boolean if a field has been set.

### SetSegment3TagIdNil

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) SetSegment3TagIdNil(b bool)`

 SetSegment3TagIdNil sets the value for Segment3TagId to be an explicit nil

### UnsetSegment3TagId
`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) UnsetSegment3TagId()`

UnsetSegment3TagId ensures that no value is present for Segment3TagId, not even an explicit nil
### GetTagIds

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) GetTagIds() []int64`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) GetTagIdsOk() (*[]int64, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) SetTagIds(v []int64)`

SetTagIds sets TagIds field to given value.


### GetTaxCode

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) GetTaxCode() int64`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) GetTaxCodeOk() (*int64, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) SetTaxCode(v int64)`

SetTaxCode sets TaxCode field to given value.


### SetTaxCodeNil

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) SetTaxCodeNil(b bool)`

 SetTaxCodeNil sets the value for TaxCode to be an explicit nil

### UnsetTaxCode
`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLinesInner) UnsetTaxCode()`

UnsetTaxCode ensures that no value is present for TaxCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


