# InvoiceIndexResponseInvoicesInnerInvoiceContentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountItemId** | **NullableInt64** | 勘定科目ID | 
**AccountItemName** | **NullableString** | 勘定科目名 | 
**Amount** | **int64** | 内税/外税の判別とamountの税込み、税抜きについて &lt;ul&gt; &lt;li&gt;tax_entry_methodがexclusive (外税)の場合&lt;/li&gt; &lt;ul&gt; &lt;li&gt;amount: 消費税抜きの金額&lt;/li&gt; &lt;li&gt;vat: 消費税の金額&lt;/li&gt; &lt;/ul&gt; &lt;li&gt;tax_entry_methodがinclusive (内税)の場合&lt;/li&gt; &lt;ul&gt; &lt;li&gt;amount: 消費税込みの金額&lt;/li&gt; &lt;li&gt;vat: 消費税の金額&lt;/li&gt; &lt;/ul&gt; &lt;/ul&gt;  | 
**Description** | **NullableString** | 備考 | 
**Id** | **int64** | 請求内容ID | 
**ItemId** | **NullableInt64** | 品目ID | 
**ItemName** | **NullableString** | 品目 | 
**Order** | **NullableInt64** | 順序 | 
**Qty** | **float32** | 数量 | 
**ReducedVat** | **bool** | 軽減税率税区分（true: 対象、false: 対象外） | 
**SectionId** | **NullableInt64** | 部門ID | 
**SectionName** | **NullableString** | 部門 | 
**Segment1TagId** | Pointer to **NullableInt64** | セグメント１ID | [optional] 
**Segment1TagName** | Pointer to **NullableString** | セグメント１ID | [optional] 
**Segment2TagId** | Pointer to **NullableInt64** | セグメント２ID | [optional] 
**Segment2TagName** | Pointer to **NullableString** | セグメント２ | [optional] 
**Segment3TagId** | Pointer to **NullableInt64** | セグメント３ID | [optional] 
**Segment3TagName** | Pointer to **NullableString** | セグメント３ | [optional] 
**TagIds** | **[]int64** |  | 
**TagNames** | **[]string** |  | 
**TaxCode** | **NullableInt64** | 税区分コード | 
**Type** | **string** | 行の種類 | 
**Unit** | **NullableString** | 単位 | 
**UnitPrice** | **float32** | 単価 | 
**Vat** | **int64** | 消費税額 | 

## Methods

### NewInvoiceIndexResponseInvoicesInnerInvoiceContentsInner

`func NewInvoiceIndexResponseInvoicesInnerInvoiceContentsInner(accountItemId NullableInt64, accountItemName NullableString, amount int64, description NullableString, id int64, itemId NullableInt64, itemName NullableString, order NullableInt64, qty float32, reducedVat bool, sectionId NullableInt64, sectionName NullableString, tagIds []int64, tagNames []string, taxCode NullableInt64, type_ string, unit NullableString, unitPrice float32, vat int64, ) *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner`

NewInvoiceIndexResponseInvoicesInnerInvoiceContentsInner instantiates a new InvoiceIndexResponseInvoicesInnerInvoiceContentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceIndexResponseInvoicesInnerInvoiceContentsInnerWithDefaults

`func NewInvoiceIndexResponseInvoicesInnerInvoiceContentsInnerWithDefaults() *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner`

NewInvoiceIndexResponseInvoicesInnerInvoiceContentsInnerWithDefaults instantiates a new InvoiceIndexResponseInvoicesInnerInvoiceContentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountItemId

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetAccountItemId() int64`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetAccountItemIdOk() (*int64, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetAccountItemId(v int64)`

SetAccountItemId sets AccountItemId field to given value.


### SetAccountItemIdNil

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetAccountItemIdNil(b bool)`

 SetAccountItemIdNil sets the value for AccountItemId to be an explicit nil

### UnsetAccountItemId
`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) UnsetAccountItemId()`

UnsetAccountItemId ensures that no value is present for AccountItemId, not even an explicit nil
### GetAccountItemName

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetAccountItemName() string`

GetAccountItemName returns the AccountItemName field if non-nil, zero value otherwise.

### GetAccountItemNameOk

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetAccountItemNameOk() (*string, bool)`

GetAccountItemNameOk returns a tuple with the AccountItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemName

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetAccountItemName(v string)`

SetAccountItemName sets AccountItemName field to given value.


### SetAccountItemNameNil

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetAccountItemNameNil(b bool)`

 SetAccountItemNameNil sets the value for AccountItemName to be an explicit nil

### UnsetAccountItemName
`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) UnsetAccountItemName()`

UnsetAccountItemName ensures that no value is present for AccountItemName, not even an explicit nil
### GetAmount

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetId(v int64)`

SetId sets Id field to given value.


### GetItemId

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetItemId(v int64)`

SetItemId sets ItemId field to given value.


### SetItemIdNil

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetItemName

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetItemName() string`

GetItemName returns the ItemName field if non-nil, zero value otherwise.

### GetItemNameOk

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetItemNameOk() (*string, bool)`

GetItemNameOk returns a tuple with the ItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemName

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetItemName(v string)`

SetItemName sets ItemName field to given value.


### SetItemNameNil

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetItemNameNil(b bool)`

 SetItemNameNil sets the value for ItemName to be an explicit nil

### UnsetItemName
`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) UnsetItemName()`

UnsetItemName ensures that no value is present for ItemName, not even an explicit nil
### GetOrder

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetOrder(v int64)`

SetOrder sets Order field to given value.


### SetOrderNil

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil
### GetQty

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetQty() float32`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetQtyOk() (*float32, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetQty(v float32)`

SetQty sets Qty field to given value.


### GetReducedVat

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetReducedVat() bool`

GetReducedVat returns the ReducedVat field if non-nil, zero value otherwise.

### GetReducedVatOk

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetReducedVatOk() (*bool, bool)`

GetReducedVatOk returns a tuple with the ReducedVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReducedVat

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetReducedVat(v bool)`

SetReducedVat sets ReducedVat field to given value.


### GetSectionId

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSectionId() int64`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSectionIdOk() (*int64, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetSectionId(v int64)`

SetSectionId sets SectionId field to given value.


### SetSectionIdNil

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetSectionIdNil(b bool)`

 SetSectionIdNil sets the value for SectionId to be an explicit nil

### UnsetSectionId
`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) UnsetSectionId()`

UnsetSectionId ensures that no value is present for SectionId, not even an explicit nil
### GetSectionName

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSectionName() string`

GetSectionName returns the SectionName field if non-nil, zero value otherwise.

### GetSectionNameOk

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSectionNameOk() (*string, bool)`

GetSectionNameOk returns a tuple with the SectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionName

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetSectionName(v string)`

SetSectionName sets SectionName field to given value.


### SetSectionNameNil

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetSectionNameNil(b bool)`

 SetSectionNameNil sets the value for SectionName to be an explicit nil

### UnsetSectionName
`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) UnsetSectionName()`

UnsetSectionName ensures that no value is present for SectionName, not even an explicit nil
### GetSegment1TagId

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSegment1TagId() int64`

GetSegment1TagId returns the Segment1TagId field if non-nil, zero value otherwise.

### GetSegment1TagIdOk

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSegment1TagIdOk() (*int64, bool)`

GetSegment1TagIdOk returns a tuple with the Segment1TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagId

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetSegment1TagId(v int64)`

SetSegment1TagId sets Segment1TagId field to given value.

### HasSegment1TagId

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) HasSegment1TagId() bool`

HasSegment1TagId returns a boolean if a field has been set.

### SetSegment1TagIdNil

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetSegment1TagIdNil(b bool)`

 SetSegment1TagIdNil sets the value for Segment1TagId to be an explicit nil

### UnsetSegment1TagId
`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) UnsetSegment1TagId()`

UnsetSegment1TagId ensures that no value is present for Segment1TagId, not even an explicit nil
### GetSegment1TagName

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSegment1TagName() string`

GetSegment1TagName returns the Segment1TagName field if non-nil, zero value otherwise.

### GetSegment1TagNameOk

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSegment1TagNameOk() (*string, bool)`

GetSegment1TagNameOk returns a tuple with the Segment1TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagName

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetSegment1TagName(v string)`

SetSegment1TagName sets Segment1TagName field to given value.

### HasSegment1TagName

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) HasSegment1TagName() bool`

HasSegment1TagName returns a boolean if a field has been set.

### SetSegment1TagNameNil

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetSegment1TagNameNil(b bool)`

 SetSegment1TagNameNil sets the value for Segment1TagName to be an explicit nil

### UnsetSegment1TagName
`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) UnsetSegment1TagName()`

UnsetSegment1TagName ensures that no value is present for Segment1TagName, not even an explicit nil
### GetSegment2TagId

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSegment2TagId() int64`

GetSegment2TagId returns the Segment2TagId field if non-nil, zero value otherwise.

### GetSegment2TagIdOk

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSegment2TagIdOk() (*int64, bool)`

GetSegment2TagIdOk returns a tuple with the Segment2TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagId

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetSegment2TagId(v int64)`

SetSegment2TagId sets Segment2TagId field to given value.

### HasSegment2TagId

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) HasSegment2TagId() bool`

HasSegment2TagId returns a boolean if a field has been set.

### SetSegment2TagIdNil

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetSegment2TagIdNil(b bool)`

 SetSegment2TagIdNil sets the value for Segment2TagId to be an explicit nil

### UnsetSegment2TagId
`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) UnsetSegment2TagId()`

UnsetSegment2TagId ensures that no value is present for Segment2TagId, not even an explicit nil
### GetSegment2TagName

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSegment2TagName() string`

GetSegment2TagName returns the Segment2TagName field if non-nil, zero value otherwise.

### GetSegment2TagNameOk

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSegment2TagNameOk() (*string, bool)`

GetSegment2TagNameOk returns a tuple with the Segment2TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagName

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetSegment2TagName(v string)`

SetSegment2TagName sets Segment2TagName field to given value.

### HasSegment2TagName

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) HasSegment2TagName() bool`

HasSegment2TagName returns a boolean if a field has been set.

### SetSegment2TagNameNil

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetSegment2TagNameNil(b bool)`

 SetSegment2TagNameNil sets the value for Segment2TagName to be an explicit nil

### UnsetSegment2TagName
`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) UnsetSegment2TagName()`

UnsetSegment2TagName ensures that no value is present for Segment2TagName, not even an explicit nil
### GetSegment3TagId

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSegment3TagId() int64`

GetSegment3TagId returns the Segment3TagId field if non-nil, zero value otherwise.

### GetSegment3TagIdOk

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSegment3TagIdOk() (*int64, bool)`

GetSegment3TagIdOk returns a tuple with the Segment3TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagId

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetSegment3TagId(v int64)`

SetSegment3TagId sets Segment3TagId field to given value.

### HasSegment3TagId

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) HasSegment3TagId() bool`

HasSegment3TagId returns a boolean if a field has been set.

### SetSegment3TagIdNil

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetSegment3TagIdNil(b bool)`

 SetSegment3TagIdNil sets the value for Segment3TagId to be an explicit nil

### UnsetSegment3TagId
`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) UnsetSegment3TagId()`

UnsetSegment3TagId ensures that no value is present for Segment3TagId, not even an explicit nil
### GetSegment3TagName

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSegment3TagName() string`

GetSegment3TagName returns the Segment3TagName field if non-nil, zero value otherwise.

### GetSegment3TagNameOk

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetSegment3TagNameOk() (*string, bool)`

GetSegment3TagNameOk returns a tuple with the Segment3TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagName

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetSegment3TagName(v string)`

SetSegment3TagName sets Segment3TagName field to given value.

### HasSegment3TagName

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) HasSegment3TagName() bool`

HasSegment3TagName returns a boolean if a field has been set.

### SetSegment3TagNameNil

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetSegment3TagNameNil(b bool)`

 SetSegment3TagNameNil sets the value for Segment3TagName to be an explicit nil

### UnsetSegment3TagName
`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) UnsetSegment3TagName()`

UnsetSegment3TagName ensures that no value is present for Segment3TagName, not even an explicit nil
### GetTagIds

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetTagIds() []int64`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetTagIdsOk() (*[]int64, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetTagIds(v []int64)`

SetTagIds sets TagIds field to given value.


### GetTagNames

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.


### GetTaxCode

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetTaxCode() int64`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetTaxCodeOk() (*int64, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetTaxCode(v int64)`

SetTaxCode sets TaxCode field to given value.


### SetTaxCodeNil

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetTaxCodeNil(b bool)`

 SetTaxCodeNil sets the value for TaxCode to be an explicit nil

### UnsetTaxCode
`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) UnsetTaxCode()`

UnsetTaxCode ensures that no value is present for TaxCode, not even an explicit nil
### GetType

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetType(v string)`

SetType sets Type field to given value.


### GetUnit

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetUnit(v string)`

SetUnit sets Unit field to given value.


### SetUnitNil

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetUnitPrice

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.


### GetVat

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetVat() int64`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) GetVatOk() (*int64, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *InvoiceIndexResponseInvoicesInnerInvoiceContentsInner) SetVat(v int64)`

SetVat sets Vat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


