# InvoiceUpdateParamsInvoiceContentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountItemId** | Pointer to **int64** | 勘定科目ID | [optional] 
**Description** | Pointer to **string** | 備考 | [optional] 
**Id** | Pointer to **int64** | 請求内容ID | [optional] 
**ItemId** | Pointer to **int64** | 品目ID | [optional] 
**Order** | **int64** | 順序 | 
**Qty** | Pointer to **float32** | 数量 | [optional] 
**SectionId** | Pointer to **int64** | 部門ID | [optional] 
**Segment1TagId** | Pointer to **int64** | セグメント１ID | [optional] 
**Segment2TagId** | Pointer to **int64** | セグメント２ID | [optional] 
**Segment3TagId** | Pointer to **int64** | セグメント３ID | [optional] 
**TagIds** | Pointer to **[]int64** |  | [optional] 
**TaxCode** | Pointer to **int64** | 税区分コード | [optional] 
**Type** | **string** | 行の種類 &lt;ul&gt; &lt;li&gt;normal、discountを指定する場合、account_item_id,tax_codeとunit_priceが必須となります。&lt;/li&gt; &lt;li&gt;normalを指定した場合、qtyが必須となります。&lt;/li&gt; &lt;/ul&gt; | 
**Unit** | Pointer to **string** | 単位 | [optional] 
**UnitPrice** | Pointer to **float32** | 単価 (tax_entry_method: inclusiveの場合は税込価格、tax_entry_method: exclusiveの場合は税抜価格となります) | [optional] 
**Vat** | Pointer to **NullableInt64** | 消費税額 | [optional] 

## Methods

### NewInvoiceUpdateParamsInvoiceContentsInner

`func NewInvoiceUpdateParamsInvoiceContentsInner(order int64, type_ string, ) *InvoiceUpdateParamsInvoiceContentsInner`

NewInvoiceUpdateParamsInvoiceContentsInner instantiates a new InvoiceUpdateParamsInvoiceContentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceUpdateParamsInvoiceContentsInnerWithDefaults

`func NewInvoiceUpdateParamsInvoiceContentsInnerWithDefaults() *InvoiceUpdateParamsInvoiceContentsInner`

NewInvoiceUpdateParamsInvoiceContentsInnerWithDefaults instantiates a new InvoiceUpdateParamsInvoiceContentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountItemId

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetAccountItemId() int64`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetAccountItemIdOk() (*int64, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *InvoiceUpdateParamsInvoiceContentsInner) SetAccountItemId(v int64)`

SetAccountItemId sets AccountItemId field to given value.

### HasAccountItemId

`func (o *InvoiceUpdateParamsInvoiceContentsInner) HasAccountItemId() bool`

HasAccountItemId returns a boolean if a field has been set.

### GetDescription

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceUpdateParamsInvoiceContentsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvoiceUpdateParamsInvoiceContentsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceUpdateParamsInvoiceContentsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceUpdateParamsInvoiceContentsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItemId

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *InvoiceUpdateParamsInvoiceContentsInner) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *InvoiceUpdateParamsInvoiceContentsInner) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetOrder

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *InvoiceUpdateParamsInvoiceContentsInner) SetOrder(v int64)`

SetOrder sets Order field to given value.


### GetQty

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetQty() float32`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetQtyOk() (*float32, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *InvoiceUpdateParamsInvoiceContentsInner) SetQty(v float32)`

SetQty sets Qty field to given value.

### HasQty

`func (o *InvoiceUpdateParamsInvoiceContentsInner) HasQty() bool`

HasQty returns a boolean if a field has been set.

### GetSectionId

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetSectionId() int64`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetSectionIdOk() (*int64, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *InvoiceUpdateParamsInvoiceContentsInner) SetSectionId(v int64)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *InvoiceUpdateParamsInvoiceContentsInner) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### GetSegment1TagId

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetSegment1TagId() int64`

GetSegment1TagId returns the Segment1TagId field if non-nil, zero value otherwise.

### GetSegment1TagIdOk

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetSegment1TagIdOk() (*int64, bool)`

GetSegment1TagIdOk returns a tuple with the Segment1TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagId

`func (o *InvoiceUpdateParamsInvoiceContentsInner) SetSegment1TagId(v int64)`

SetSegment1TagId sets Segment1TagId field to given value.

### HasSegment1TagId

`func (o *InvoiceUpdateParamsInvoiceContentsInner) HasSegment1TagId() bool`

HasSegment1TagId returns a boolean if a field has been set.

### GetSegment2TagId

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetSegment2TagId() int64`

GetSegment2TagId returns the Segment2TagId field if non-nil, zero value otherwise.

### GetSegment2TagIdOk

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetSegment2TagIdOk() (*int64, bool)`

GetSegment2TagIdOk returns a tuple with the Segment2TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagId

`func (o *InvoiceUpdateParamsInvoiceContentsInner) SetSegment2TagId(v int64)`

SetSegment2TagId sets Segment2TagId field to given value.

### HasSegment2TagId

`func (o *InvoiceUpdateParamsInvoiceContentsInner) HasSegment2TagId() bool`

HasSegment2TagId returns a boolean if a field has been set.

### GetSegment3TagId

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetSegment3TagId() int64`

GetSegment3TagId returns the Segment3TagId field if non-nil, zero value otherwise.

### GetSegment3TagIdOk

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetSegment3TagIdOk() (*int64, bool)`

GetSegment3TagIdOk returns a tuple with the Segment3TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagId

`func (o *InvoiceUpdateParamsInvoiceContentsInner) SetSegment3TagId(v int64)`

SetSegment3TagId sets Segment3TagId field to given value.

### HasSegment3TagId

`func (o *InvoiceUpdateParamsInvoiceContentsInner) HasSegment3TagId() bool`

HasSegment3TagId returns a boolean if a field has been set.

### GetTagIds

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetTagIds() []int64`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetTagIdsOk() (*[]int64, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *InvoiceUpdateParamsInvoiceContentsInner) SetTagIds(v []int64)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *InvoiceUpdateParamsInvoiceContentsInner) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTaxCode

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetTaxCode() int64`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetTaxCodeOk() (*int64, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *InvoiceUpdateParamsInvoiceContentsInner) SetTaxCode(v int64)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *InvoiceUpdateParamsInvoiceContentsInner) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetType

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvoiceUpdateParamsInvoiceContentsInner) SetType(v string)`

SetType sets Type field to given value.


### GetUnit

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *InvoiceUpdateParamsInvoiceContentsInner) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *InvoiceUpdateParamsInvoiceContentsInner) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUnitPrice

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *InvoiceUpdateParamsInvoiceContentsInner) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *InvoiceUpdateParamsInvoiceContentsInner) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetVat

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetVat() int64`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *InvoiceUpdateParamsInvoiceContentsInner) GetVatOk() (*int64, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *InvoiceUpdateParamsInvoiceContentsInner) SetVat(v int64)`

SetVat sets Vat field to given value.

### HasVat

`func (o *InvoiceUpdateParamsInvoiceContentsInner) HasVat() bool`

HasVat returns a boolean if a field has been set.

### SetVatNil

`func (o *InvoiceUpdateParamsInvoiceContentsInner) SetVatNil(b bool)`

 SetVatNil sets the value for Vat to be an explicit nil

### UnsetVat
`func (o *InvoiceUpdateParamsInvoiceContentsInner) UnsetVat()`

UnsetVat ensures that no value is present for Vat, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


