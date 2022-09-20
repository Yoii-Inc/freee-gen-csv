# QuotationUpdateParamsQuotationContentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountItemId** | Pointer to **int32** | 勘定科目ID | [optional] 
**Description** | Pointer to **string** | 備考 | [optional] 
**Id** | Pointer to **int32** | 見積内容ID | [optional] 
**ItemId** | Pointer to **int32** | 品目ID | [optional] 
**Order** | **int32** | 順序 | 
**Qty** | Pointer to **float32** | 数量 | [optional] 
**SectionId** | Pointer to **int32** | 部門ID | [optional] 
**Segment1TagId** | Pointer to **int64** | セグメント１ID | [optional] 
**Segment2TagId** | Pointer to **int64** | セグメント２ID | [optional] 
**Segment3TagId** | Pointer to **int64** | セグメント３ID | [optional] 
**TagIds** | Pointer to **[]int32** |  | [optional] 
**TaxCode** | Pointer to **int32** | 税区分コード | [optional] 
**Type** | **string** | 行の種類 &lt;ul&gt; &lt;li&gt;normal、discountを指定する場合、account_item_id,tax_codeとunit_priceが必須となります。&lt;/li&gt; &lt;li&gt;normalを指定した場合、qtyが必須となります。&lt;/li&gt; &lt;/ul&gt; | 
**Unit** | Pointer to **string** | 単位 | [optional] 
**UnitPrice** | Pointer to **float32** | 単価 (tax_entry_method: inclusiveの場合は税込価格、tax_entry_method: exclusiveの場合は税抜価格となります) | [optional] 
**Vat** | Pointer to **NullableInt32** | 消費税額 | [optional] 

## Methods

### NewQuotationUpdateParamsQuotationContentsInner

`func NewQuotationUpdateParamsQuotationContentsInner(order int32, type_ string, ) *QuotationUpdateParamsQuotationContentsInner`

NewQuotationUpdateParamsQuotationContentsInner instantiates a new QuotationUpdateParamsQuotationContentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotationUpdateParamsQuotationContentsInnerWithDefaults

`func NewQuotationUpdateParamsQuotationContentsInnerWithDefaults() *QuotationUpdateParamsQuotationContentsInner`

NewQuotationUpdateParamsQuotationContentsInnerWithDefaults instantiates a new QuotationUpdateParamsQuotationContentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountItemId

`func (o *QuotationUpdateParamsQuotationContentsInner) GetAccountItemId() int32`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *QuotationUpdateParamsQuotationContentsInner) GetAccountItemIdOk() (*int32, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *QuotationUpdateParamsQuotationContentsInner) SetAccountItemId(v int32)`

SetAccountItemId sets AccountItemId field to given value.

### HasAccountItemId

`func (o *QuotationUpdateParamsQuotationContentsInner) HasAccountItemId() bool`

HasAccountItemId returns a boolean if a field has been set.

### GetDescription

`func (o *QuotationUpdateParamsQuotationContentsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuotationUpdateParamsQuotationContentsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuotationUpdateParamsQuotationContentsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuotationUpdateParamsQuotationContentsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *QuotationUpdateParamsQuotationContentsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuotationUpdateParamsQuotationContentsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuotationUpdateParamsQuotationContentsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *QuotationUpdateParamsQuotationContentsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItemId

`func (o *QuotationUpdateParamsQuotationContentsInner) GetItemId() int32`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *QuotationUpdateParamsQuotationContentsInner) GetItemIdOk() (*int32, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *QuotationUpdateParamsQuotationContentsInner) SetItemId(v int32)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *QuotationUpdateParamsQuotationContentsInner) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetOrder

`func (o *QuotationUpdateParamsQuotationContentsInner) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *QuotationUpdateParamsQuotationContentsInner) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *QuotationUpdateParamsQuotationContentsInner) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetQty

`func (o *QuotationUpdateParamsQuotationContentsInner) GetQty() float32`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *QuotationUpdateParamsQuotationContentsInner) GetQtyOk() (*float32, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *QuotationUpdateParamsQuotationContentsInner) SetQty(v float32)`

SetQty sets Qty field to given value.

### HasQty

`func (o *QuotationUpdateParamsQuotationContentsInner) HasQty() bool`

HasQty returns a boolean if a field has been set.

### GetSectionId

`func (o *QuotationUpdateParamsQuotationContentsInner) GetSectionId() int32`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *QuotationUpdateParamsQuotationContentsInner) GetSectionIdOk() (*int32, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *QuotationUpdateParamsQuotationContentsInner) SetSectionId(v int32)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *QuotationUpdateParamsQuotationContentsInner) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### GetSegment1TagId

`func (o *QuotationUpdateParamsQuotationContentsInner) GetSegment1TagId() int64`

GetSegment1TagId returns the Segment1TagId field if non-nil, zero value otherwise.

### GetSegment1TagIdOk

`func (o *QuotationUpdateParamsQuotationContentsInner) GetSegment1TagIdOk() (*int64, bool)`

GetSegment1TagIdOk returns a tuple with the Segment1TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagId

`func (o *QuotationUpdateParamsQuotationContentsInner) SetSegment1TagId(v int64)`

SetSegment1TagId sets Segment1TagId field to given value.

### HasSegment1TagId

`func (o *QuotationUpdateParamsQuotationContentsInner) HasSegment1TagId() bool`

HasSegment1TagId returns a boolean if a field has been set.

### GetSegment2TagId

`func (o *QuotationUpdateParamsQuotationContentsInner) GetSegment2TagId() int64`

GetSegment2TagId returns the Segment2TagId field if non-nil, zero value otherwise.

### GetSegment2TagIdOk

`func (o *QuotationUpdateParamsQuotationContentsInner) GetSegment2TagIdOk() (*int64, bool)`

GetSegment2TagIdOk returns a tuple with the Segment2TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagId

`func (o *QuotationUpdateParamsQuotationContentsInner) SetSegment2TagId(v int64)`

SetSegment2TagId sets Segment2TagId field to given value.

### HasSegment2TagId

`func (o *QuotationUpdateParamsQuotationContentsInner) HasSegment2TagId() bool`

HasSegment2TagId returns a boolean if a field has been set.

### GetSegment3TagId

`func (o *QuotationUpdateParamsQuotationContentsInner) GetSegment3TagId() int64`

GetSegment3TagId returns the Segment3TagId field if non-nil, zero value otherwise.

### GetSegment3TagIdOk

`func (o *QuotationUpdateParamsQuotationContentsInner) GetSegment3TagIdOk() (*int64, bool)`

GetSegment3TagIdOk returns a tuple with the Segment3TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagId

`func (o *QuotationUpdateParamsQuotationContentsInner) SetSegment3TagId(v int64)`

SetSegment3TagId sets Segment3TagId field to given value.

### HasSegment3TagId

`func (o *QuotationUpdateParamsQuotationContentsInner) HasSegment3TagId() bool`

HasSegment3TagId returns a boolean if a field has been set.

### GetTagIds

`func (o *QuotationUpdateParamsQuotationContentsInner) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *QuotationUpdateParamsQuotationContentsInner) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *QuotationUpdateParamsQuotationContentsInner) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *QuotationUpdateParamsQuotationContentsInner) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTaxCode

`func (o *QuotationUpdateParamsQuotationContentsInner) GetTaxCode() int32`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *QuotationUpdateParamsQuotationContentsInner) GetTaxCodeOk() (*int32, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *QuotationUpdateParamsQuotationContentsInner) SetTaxCode(v int32)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *QuotationUpdateParamsQuotationContentsInner) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetType

`func (o *QuotationUpdateParamsQuotationContentsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QuotationUpdateParamsQuotationContentsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QuotationUpdateParamsQuotationContentsInner) SetType(v string)`

SetType sets Type field to given value.


### GetUnit

`func (o *QuotationUpdateParamsQuotationContentsInner) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *QuotationUpdateParamsQuotationContentsInner) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *QuotationUpdateParamsQuotationContentsInner) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *QuotationUpdateParamsQuotationContentsInner) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUnitPrice

`func (o *QuotationUpdateParamsQuotationContentsInner) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *QuotationUpdateParamsQuotationContentsInner) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *QuotationUpdateParamsQuotationContentsInner) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *QuotationUpdateParamsQuotationContentsInner) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetVat

`func (o *QuotationUpdateParamsQuotationContentsInner) GetVat() int32`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *QuotationUpdateParamsQuotationContentsInner) GetVatOk() (*int32, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *QuotationUpdateParamsQuotationContentsInner) SetVat(v int32)`

SetVat sets Vat field to given value.

### HasVat

`func (o *QuotationUpdateParamsQuotationContentsInner) HasVat() bool`

HasVat returns a boolean if a field has been set.

### SetVatNil

`func (o *QuotationUpdateParamsQuotationContentsInner) SetVatNil(b bool)`

 SetVatNil sets the value for Vat to be an explicit nil

### UnsetVat
`func (o *QuotationUpdateParamsQuotationContentsInner) UnsetVat()`

UnsetVat ensures that no value is present for Vat, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


