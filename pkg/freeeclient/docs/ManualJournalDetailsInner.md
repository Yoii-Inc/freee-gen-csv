# ManualJournalDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountItemId** | **int32** | 勘定科目ID | 
**Amount** | **int32** | 金額（税込で指定してください） | 
**Description** | **string** | 備考 | 
**EntrySide** | **string** | 貸借(貸方: credit, 借方: debit) | 
**Id** | **int64** | 貸借行ID | 
**ItemId** | **NullableInt32** | 品目ID | 
**ItemName** | **NullableString** | 品目 | 
**PartnerCode** | Pointer to **NullableString** | 取引先コード | [optional] 
**PartnerId** | **NullableInt32** | 取引先ID | 
**PartnerLongName** | **NullableString** | 正式名称（255文字以内） | 
**PartnerName** | **NullableString** | 取引先名 | 
**SectionId** | **NullableInt32** | 部門ID | 
**SectionName** | **NullableString** | 部門 | 
**Segment1TagId** | Pointer to **int64** | セグメント１ID | [optional] 
**Segment1TagName** | Pointer to **int32** | セグメント１ID | [optional] 
**Segment2TagId** | Pointer to **int64** | セグメント２ID | [optional] 
**Segment2TagName** | Pointer to **int32** | セグメント２ | [optional] 
**Segment3TagId** | Pointer to **int64** | セグメント３ID | [optional] 
**Segment3TagName** | Pointer to **int32** | セグメント３ | [optional] 
**TagIds** | **[]int32** |  | 
**TagNames** | **[]string** |  | 
**TaxCode** | **int32** | 税区分コード | 
**Vat** | **int32** | 消費税額（指定しない場合は自動で計算されます） | 

## Methods

### NewManualJournalDetailsInner

`func NewManualJournalDetailsInner(accountItemId int32, amount int32, description string, entrySide string, id int64, itemId NullableInt32, itemName NullableString, partnerId NullableInt32, partnerLongName NullableString, partnerName NullableString, sectionId NullableInt32, sectionName NullableString, tagIds []int32, tagNames []string, taxCode int32, vat int32, ) *ManualJournalDetailsInner`

NewManualJournalDetailsInner instantiates a new ManualJournalDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualJournalDetailsInnerWithDefaults

`func NewManualJournalDetailsInnerWithDefaults() *ManualJournalDetailsInner`

NewManualJournalDetailsInnerWithDefaults instantiates a new ManualJournalDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountItemId

`func (o *ManualJournalDetailsInner) GetAccountItemId() int32`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *ManualJournalDetailsInner) GetAccountItemIdOk() (*int32, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *ManualJournalDetailsInner) SetAccountItemId(v int32)`

SetAccountItemId sets AccountItemId field to given value.


### GetAmount

`func (o *ManualJournalDetailsInner) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ManualJournalDetailsInner) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ManualJournalDetailsInner) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *ManualJournalDetailsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManualJournalDetailsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManualJournalDetailsInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEntrySide

`func (o *ManualJournalDetailsInner) GetEntrySide() string`

GetEntrySide returns the EntrySide field if non-nil, zero value otherwise.

### GetEntrySideOk

`func (o *ManualJournalDetailsInner) GetEntrySideOk() (*string, bool)`

GetEntrySideOk returns a tuple with the EntrySide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrySide

`func (o *ManualJournalDetailsInner) SetEntrySide(v string)`

SetEntrySide sets EntrySide field to given value.


### GetId

`func (o *ManualJournalDetailsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManualJournalDetailsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManualJournalDetailsInner) SetId(v int64)`

SetId sets Id field to given value.


### GetItemId

`func (o *ManualJournalDetailsInner) GetItemId() int32`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ManualJournalDetailsInner) GetItemIdOk() (*int32, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ManualJournalDetailsInner) SetItemId(v int32)`

SetItemId sets ItemId field to given value.


### SetItemIdNil

`func (o *ManualJournalDetailsInner) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *ManualJournalDetailsInner) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetItemName

`func (o *ManualJournalDetailsInner) GetItemName() string`

GetItemName returns the ItemName field if non-nil, zero value otherwise.

### GetItemNameOk

`func (o *ManualJournalDetailsInner) GetItemNameOk() (*string, bool)`

GetItemNameOk returns a tuple with the ItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemName

`func (o *ManualJournalDetailsInner) SetItemName(v string)`

SetItemName sets ItemName field to given value.


### SetItemNameNil

`func (o *ManualJournalDetailsInner) SetItemNameNil(b bool)`

 SetItemNameNil sets the value for ItemName to be an explicit nil

### UnsetItemName
`func (o *ManualJournalDetailsInner) UnsetItemName()`

UnsetItemName ensures that no value is present for ItemName, not even an explicit nil
### GetPartnerCode

`func (o *ManualJournalDetailsInner) GetPartnerCode() string`

GetPartnerCode returns the PartnerCode field if non-nil, zero value otherwise.

### GetPartnerCodeOk

`func (o *ManualJournalDetailsInner) GetPartnerCodeOk() (*string, bool)`

GetPartnerCodeOk returns a tuple with the PartnerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCode

`func (o *ManualJournalDetailsInner) SetPartnerCode(v string)`

SetPartnerCode sets PartnerCode field to given value.

### HasPartnerCode

`func (o *ManualJournalDetailsInner) HasPartnerCode() bool`

HasPartnerCode returns a boolean if a field has been set.

### SetPartnerCodeNil

`func (o *ManualJournalDetailsInner) SetPartnerCodeNil(b bool)`

 SetPartnerCodeNil sets the value for PartnerCode to be an explicit nil

### UnsetPartnerCode
`func (o *ManualJournalDetailsInner) UnsetPartnerCode()`

UnsetPartnerCode ensures that no value is present for PartnerCode, not even an explicit nil
### GetPartnerId

`func (o *ManualJournalDetailsInner) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *ManualJournalDetailsInner) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *ManualJournalDetailsInner) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.


### SetPartnerIdNil

`func (o *ManualJournalDetailsInner) SetPartnerIdNil(b bool)`

 SetPartnerIdNil sets the value for PartnerId to be an explicit nil

### UnsetPartnerId
`func (o *ManualJournalDetailsInner) UnsetPartnerId()`

UnsetPartnerId ensures that no value is present for PartnerId, not even an explicit nil
### GetPartnerLongName

`func (o *ManualJournalDetailsInner) GetPartnerLongName() string`

GetPartnerLongName returns the PartnerLongName field if non-nil, zero value otherwise.

### GetPartnerLongNameOk

`func (o *ManualJournalDetailsInner) GetPartnerLongNameOk() (*string, bool)`

GetPartnerLongNameOk returns a tuple with the PartnerLongName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerLongName

`func (o *ManualJournalDetailsInner) SetPartnerLongName(v string)`

SetPartnerLongName sets PartnerLongName field to given value.


### SetPartnerLongNameNil

`func (o *ManualJournalDetailsInner) SetPartnerLongNameNil(b bool)`

 SetPartnerLongNameNil sets the value for PartnerLongName to be an explicit nil

### UnsetPartnerLongName
`func (o *ManualJournalDetailsInner) UnsetPartnerLongName()`

UnsetPartnerLongName ensures that no value is present for PartnerLongName, not even an explicit nil
### GetPartnerName

`func (o *ManualJournalDetailsInner) GetPartnerName() string`

GetPartnerName returns the PartnerName field if non-nil, zero value otherwise.

### GetPartnerNameOk

`func (o *ManualJournalDetailsInner) GetPartnerNameOk() (*string, bool)`

GetPartnerNameOk returns a tuple with the PartnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerName

`func (o *ManualJournalDetailsInner) SetPartnerName(v string)`

SetPartnerName sets PartnerName field to given value.


### SetPartnerNameNil

`func (o *ManualJournalDetailsInner) SetPartnerNameNil(b bool)`

 SetPartnerNameNil sets the value for PartnerName to be an explicit nil

### UnsetPartnerName
`func (o *ManualJournalDetailsInner) UnsetPartnerName()`

UnsetPartnerName ensures that no value is present for PartnerName, not even an explicit nil
### GetSectionId

`func (o *ManualJournalDetailsInner) GetSectionId() int32`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *ManualJournalDetailsInner) GetSectionIdOk() (*int32, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *ManualJournalDetailsInner) SetSectionId(v int32)`

SetSectionId sets SectionId field to given value.


### SetSectionIdNil

`func (o *ManualJournalDetailsInner) SetSectionIdNil(b bool)`

 SetSectionIdNil sets the value for SectionId to be an explicit nil

### UnsetSectionId
`func (o *ManualJournalDetailsInner) UnsetSectionId()`

UnsetSectionId ensures that no value is present for SectionId, not even an explicit nil
### GetSectionName

`func (o *ManualJournalDetailsInner) GetSectionName() string`

GetSectionName returns the SectionName field if non-nil, zero value otherwise.

### GetSectionNameOk

`func (o *ManualJournalDetailsInner) GetSectionNameOk() (*string, bool)`

GetSectionNameOk returns a tuple with the SectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionName

`func (o *ManualJournalDetailsInner) SetSectionName(v string)`

SetSectionName sets SectionName field to given value.


### SetSectionNameNil

`func (o *ManualJournalDetailsInner) SetSectionNameNil(b bool)`

 SetSectionNameNil sets the value for SectionName to be an explicit nil

### UnsetSectionName
`func (o *ManualJournalDetailsInner) UnsetSectionName()`

UnsetSectionName ensures that no value is present for SectionName, not even an explicit nil
### GetSegment1TagId

`func (o *ManualJournalDetailsInner) GetSegment1TagId() int64`

GetSegment1TagId returns the Segment1TagId field if non-nil, zero value otherwise.

### GetSegment1TagIdOk

`func (o *ManualJournalDetailsInner) GetSegment1TagIdOk() (*int64, bool)`

GetSegment1TagIdOk returns a tuple with the Segment1TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagId

`func (o *ManualJournalDetailsInner) SetSegment1TagId(v int64)`

SetSegment1TagId sets Segment1TagId field to given value.

### HasSegment1TagId

`func (o *ManualJournalDetailsInner) HasSegment1TagId() bool`

HasSegment1TagId returns a boolean if a field has been set.

### GetSegment1TagName

`func (o *ManualJournalDetailsInner) GetSegment1TagName() int32`

GetSegment1TagName returns the Segment1TagName field if non-nil, zero value otherwise.

### GetSegment1TagNameOk

`func (o *ManualJournalDetailsInner) GetSegment1TagNameOk() (*int32, bool)`

GetSegment1TagNameOk returns a tuple with the Segment1TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagName

`func (o *ManualJournalDetailsInner) SetSegment1TagName(v int32)`

SetSegment1TagName sets Segment1TagName field to given value.

### HasSegment1TagName

`func (o *ManualJournalDetailsInner) HasSegment1TagName() bool`

HasSegment1TagName returns a boolean if a field has been set.

### GetSegment2TagId

`func (o *ManualJournalDetailsInner) GetSegment2TagId() int64`

GetSegment2TagId returns the Segment2TagId field if non-nil, zero value otherwise.

### GetSegment2TagIdOk

`func (o *ManualJournalDetailsInner) GetSegment2TagIdOk() (*int64, bool)`

GetSegment2TagIdOk returns a tuple with the Segment2TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagId

`func (o *ManualJournalDetailsInner) SetSegment2TagId(v int64)`

SetSegment2TagId sets Segment2TagId field to given value.

### HasSegment2TagId

`func (o *ManualJournalDetailsInner) HasSegment2TagId() bool`

HasSegment2TagId returns a boolean if a field has been set.

### GetSegment2TagName

`func (o *ManualJournalDetailsInner) GetSegment2TagName() int32`

GetSegment2TagName returns the Segment2TagName field if non-nil, zero value otherwise.

### GetSegment2TagNameOk

`func (o *ManualJournalDetailsInner) GetSegment2TagNameOk() (*int32, bool)`

GetSegment2TagNameOk returns a tuple with the Segment2TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagName

`func (o *ManualJournalDetailsInner) SetSegment2TagName(v int32)`

SetSegment2TagName sets Segment2TagName field to given value.

### HasSegment2TagName

`func (o *ManualJournalDetailsInner) HasSegment2TagName() bool`

HasSegment2TagName returns a boolean if a field has been set.

### GetSegment3TagId

`func (o *ManualJournalDetailsInner) GetSegment3TagId() int64`

GetSegment3TagId returns the Segment3TagId field if non-nil, zero value otherwise.

### GetSegment3TagIdOk

`func (o *ManualJournalDetailsInner) GetSegment3TagIdOk() (*int64, bool)`

GetSegment3TagIdOk returns a tuple with the Segment3TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagId

`func (o *ManualJournalDetailsInner) SetSegment3TagId(v int64)`

SetSegment3TagId sets Segment3TagId field to given value.

### HasSegment3TagId

`func (o *ManualJournalDetailsInner) HasSegment3TagId() bool`

HasSegment3TagId returns a boolean if a field has been set.

### GetSegment3TagName

`func (o *ManualJournalDetailsInner) GetSegment3TagName() int32`

GetSegment3TagName returns the Segment3TagName field if non-nil, zero value otherwise.

### GetSegment3TagNameOk

`func (o *ManualJournalDetailsInner) GetSegment3TagNameOk() (*int32, bool)`

GetSegment3TagNameOk returns a tuple with the Segment3TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagName

`func (o *ManualJournalDetailsInner) SetSegment3TagName(v int32)`

SetSegment3TagName sets Segment3TagName field to given value.

### HasSegment3TagName

`func (o *ManualJournalDetailsInner) HasSegment3TagName() bool`

HasSegment3TagName returns a boolean if a field has been set.

### GetTagIds

`func (o *ManualJournalDetailsInner) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *ManualJournalDetailsInner) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *ManualJournalDetailsInner) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.


### GetTagNames

`func (o *ManualJournalDetailsInner) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *ManualJournalDetailsInner) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *ManualJournalDetailsInner) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.


### GetTaxCode

`func (o *ManualJournalDetailsInner) GetTaxCode() int32`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *ManualJournalDetailsInner) GetTaxCodeOk() (*int32, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *ManualJournalDetailsInner) SetTaxCode(v int32)`

SetTaxCode sets TaxCode field to given value.


### GetVat

`func (o *ManualJournalDetailsInner) GetVat() int32`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *ManualJournalDetailsInner) GetVatOk() (*int32, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *ManualJournalDetailsInner) SetVat(v int32)`

SetVat sets Vat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


