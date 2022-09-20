# ManualJournalUpdateParamsDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountItemId** | **int32** | 勘定科目ID | 
**Amount** | **int64** | 取引金額（税込で指定してください） | 
**Description** | Pointer to **string** | 備考 | [optional] 
**EntrySide** | **string** | 貸借（貸方: credit, 借方: debit） | 
**Id** | Pointer to **int64** | 貸借行ID: 既存貸借行を更新または削除する場合に指定します。IDを指定しない貸借行は、新規行として扱われ追加されます。 | [optional] 
**ItemId** | Pointer to **int32** | 品目ID | [optional] 
**PartnerCode** | Pointer to **string** | 取引先コード | [optional] 
**PartnerId** | Pointer to **int32** | 取引先ID | [optional] 
**SectionId** | Pointer to **int32** | 部門ID | [optional] 
**Segment1TagId** | Pointer to **int64** | セグメント１ID | [optional] 
**Segment2TagId** | Pointer to **int64** | セグメント２ID | [optional] 
**Segment3TagId** | Pointer to **int64** | セグメント３ID | [optional] 
**TagIds** | Pointer to **[]int32** | メモタグID | [optional] 
**TaxCode** | **int32** | 税区分コード | 
**Vat** | Pointer to **int32** | 消費税額（指定しない場合は自動で計算されます） | [optional] 

## Methods

### NewManualJournalUpdateParamsDetailsInner

`func NewManualJournalUpdateParamsDetailsInner(accountItemId int32, amount int64, entrySide string, taxCode int32, ) *ManualJournalUpdateParamsDetailsInner`

NewManualJournalUpdateParamsDetailsInner instantiates a new ManualJournalUpdateParamsDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualJournalUpdateParamsDetailsInnerWithDefaults

`func NewManualJournalUpdateParamsDetailsInnerWithDefaults() *ManualJournalUpdateParamsDetailsInner`

NewManualJournalUpdateParamsDetailsInnerWithDefaults instantiates a new ManualJournalUpdateParamsDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountItemId

`func (o *ManualJournalUpdateParamsDetailsInner) GetAccountItemId() int32`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *ManualJournalUpdateParamsDetailsInner) GetAccountItemIdOk() (*int32, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *ManualJournalUpdateParamsDetailsInner) SetAccountItemId(v int32)`

SetAccountItemId sets AccountItemId field to given value.


### GetAmount

`func (o *ManualJournalUpdateParamsDetailsInner) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ManualJournalUpdateParamsDetailsInner) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ManualJournalUpdateParamsDetailsInner) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *ManualJournalUpdateParamsDetailsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManualJournalUpdateParamsDetailsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManualJournalUpdateParamsDetailsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManualJournalUpdateParamsDetailsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntrySide

`func (o *ManualJournalUpdateParamsDetailsInner) GetEntrySide() string`

GetEntrySide returns the EntrySide field if non-nil, zero value otherwise.

### GetEntrySideOk

`func (o *ManualJournalUpdateParamsDetailsInner) GetEntrySideOk() (*string, bool)`

GetEntrySideOk returns a tuple with the EntrySide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrySide

`func (o *ManualJournalUpdateParamsDetailsInner) SetEntrySide(v string)`

SetEntrySide sets EntrySide field to given value.


### GetId

`func (o *ManualJournalUpdateParamsDetailsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManualJournalUpdateParamsDetailsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManualJournalUpdateParamsDetailsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ManualJournalUpdateParamsDetailsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItemId

`func (o *ManualJournalUpdateParamsDetailsInner) GetItemId() int32`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ManualJournalUpdateParamsDetailsInner) GetItemIdOk() (*int32, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ManualJournalUpdateParamsDetailsInner) SetItemId(v int32)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ManualJournalUpdateParamsDetailsInner) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetPartnerCode

`func (o *ManualJournalUpdateParamsDetailsInner) GetPartnerCode() string`

GetPartnerCode returns the PartnerCode field if non-nil, zero value otherwise.

### GetPartnerCodeOk

`func (o *ManualJournalUpdateParamsDetailsInner) GetPartnerCodeOk() (*string, bool)`

GetPartnerCodeOk returns a tuple with the PartnerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCode

`func (o *ManualJournalUpdateParamsDetailsInner) SetPartnerCode(v string)`

SetPartnerCode sets PartnerCode field to given value.

### HasPartnerCode

`func (o *ManualJournalUpdateParamsDetailsInner) HasPartnerCode() bool`

HasPartnerCode returns a boolean if a field has been set.

### GetPartnerId

`func (o *ManualJournalUpdateParamsDetailsInner) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *ManualJournalUpdateParamsDetailsInner) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *ManualJournalUpdateParamsDetailsInner) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *ManualJournalUpdateParamsDetailsInner) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetSectionId

`func (o *ManualJournalUpdateParamsDetailsInner) GetSectionId() int32`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *ManualJournalUpdateParamsDetailsInner) GetSectionIdOk() (*int32, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *ManualJournalUpdateParamsDetailsInner) SetSectionId(v int32)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *ManualJournalUpdateParamsDetailsInner) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### GetSegment1TagId

`func (o *ManualJournalUpdateParamsDetailsInner) GetSegment1TagId() int64`

GetSegment1TagId returns the Segment1TagId field if non-nil, zero value otherwise.

### GetSegment1TagIdOk

`func (o *ManualJournalUpdateParamsDetailsInner) GetSegment1TagIdOk() (*int64, bool)`

GetSegment1TagIdOk returns a tuple with the Segment1TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagId

`func (o *ManualJournalUpdateParamsDetailsInner) SetSegment1TagId(v int64)`

SetSegment1TagId sets Segment1TagId field to given value.

### HasSegment1TagId

`func (o *ManualJournalUpdateParamsDetailsInner) HasSegment1TagId() bool`

HasSegment1TagId returns a boolean if a field has been set.

### GetSegment2TagId

`func (o *ManualJournalUpdateParamsDetailsInner) GetSegment2TagId() int64`

GetSegment2TagId returns the Segment2TagId field if non-nil, zero value otherwise.

### GetSegment2TagIdOk

`func (o *ManualJournalUpdateParamsDetailsInner) GetSegment2TagIdOk() (*int64, bool)`

GetSegment2TagIdOk returns a tuple with the Segment2TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagId

`func (o *ManualJournalUpdateParamsDetailsInner) SetSegment2TagId(v int64)`

SetSegment2TagId sets Segment2TagId field to given value.

### HasSegment2TagId

`func (o *ManualJournalUpdateParamsDetailsInner) HasSegment2TagId() bool`

HasSegment2TagId returns a boolean if a field has been set.

### GetSegment3TagId

`func (o *ManualJournalUpdateParamsDetailsInner) GetSegment3TagId() int64`

GetSegment3TagId returns the Segment3TagId field if non-nil, zero value otherwise.

### GetSegment3TagIdOk

`func (o *ManualJournalUpdateParamsDetailsInner) GetSegment3TagIdOk() (*int64, bool)`

GetSegment3TagIdOk returns a tuple with the Segment3TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagId

`func (o *ManualJournalUpdateParamsDetailsInner) SetSegment3TagId(v int64)`

SetSegment3TagId sets Segment3TagId field to given value.

### HasSegment3TagId

`func (o *ManualJournalUpdateParamsDetailsInner) HasSegment3TagId() bool`

HasSegment3TagId returns a boolean if a field has been set.

### GetTagIds

`func (o *ManualJournalUpdateParamsDetailsInner) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *ManualJournalUpdateParamsDetailsInner) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *ManualJournalUpdateParamsDetailsInner) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *ManualJournalUpdateParamsDetailsInner) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTaxCode

`func (o *ManualJournalUpdateParamsDetailsInner) GetTaxCode() int32`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *ManualJournalUpdateParamsDetailsInner) GetTaxCodeOk() (*int32, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *ManualJournalUpdateParamsDetailsInner) SetTaxCode(v int32)`

SetTaxCode sets TaxCode field to given value.


### GetVat

`func (o *ManualJournalUpdateParamsDetailsInner) GetVat() int32`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *ManualJournalUpdateParamsDetailsInner) GetVatOk() (*int32, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *ManualJournalUpdateParamsDetailsInner) SetVat(v int32)`

SetVat sets Vat field to given value.

### HasVat

`func (o *ManualJournalUpdateParamsDetailsInner) HasVat() bool`

HasVat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


