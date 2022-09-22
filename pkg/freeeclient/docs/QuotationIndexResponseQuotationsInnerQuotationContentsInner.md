# QuotationIndexResponseQuotationsInnerQuotationContentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountItemId** | **NullableInt64** | 勘定科目ID | 
**AccountItemName** | **NullableString** | 勘定科目名 | 
**Amount** | **int64** | 内税/外税の判別とamountの税込み、税抜きについて &lt;ul&gt; &lt;li&gt;tax_entry_methodがexclusive (外税)の場合&lt;/li&gt; &lt;ul&gt; &lt;li&gt;amount: 消費税抜きの金額&lt;/li&gt; &lt;li&gt;vat: 消費税の金額&lt;/li&gt; &lt;/ul&gt; &lt;li&gt;tax_entry_methodがinclusive (内税)の場合&lt;/li&gt; &lt;ul&gt; &lt;li&gt;amount: 消費税込みの金額&lt;/li&gt; &lt;li&gt;vat: 消費税の金額&lt;/li&gt; &lt;/ul&gt; &lt;/ul&gt;  | 
**Description** | **NullableString** | 備考 | 
**Id** | **int64** | 見積内容ID | 
**ItemId** | **NullableInt64** | 品目ID | 
**ItemName** | **NullableString** | 品目 | 
**Order** | **int64** | 順序 | 
**Qty** | **float32** | 数量 | 
**ReducedVat** | **bool** | 軽減税率税区分（true: 対象、false: 対象外） | 
**SectionId** | **NullableInt64** | 部門ID | 
**SectionName** | **NullableString** | 部門 | 
**Segment1TagId** | Pointer to **NullableInt64** | セグメント１ID | [optional] 
**Segment1TagName** | Pointer to **NullableString** | セグメント１ | [optional] 
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

### NewQuotationIndexResponseQuotationsInnerQuotationContentsInner

`func NewQuotationIndexResponseQuotationsInnerQuotationContentsInner(accountItemId NullableInt64, accountItemName NullableString, amount int64, description NullableString, id int64, itemId NullableInt64, itemName NullableString, order int64, qty float32, reducedVat bool, sectionId NullableInt64, sectionName NullableString, tagIds []int64, tagNames []string, taxCode NullableInt64, type_ string, unit NullableString, unitPrice float32, vat int64, ) *QuotationIndexResponseQuotationsInnerQuotationContentsInner`

NewQuotationIndexResponseQuotationsInnerQuotationContentsInner instantiates a new QuotationIndexResponseQuotationsInnerQuotationContentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotationIndexResponseQuotationsInnerQuotationContentsInnerWithDefaults

`func NewQuotationIndexResponseQuotationsInnerQuotationContentsInnerWithDefaults() *QuotationIndexResponseQuotationsInnerQuotationContentsInner`

NewQuotationIndexResponseQuotationsInnerQuotationContentsInnerWithDefaults instantiates a new QuotationIndexResponseQuotationsInnerQuotationContentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountItemId

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetAccountItemId() int64`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetAccountItemIdOk() (*int64, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetAccountItemId(v int64)`

SetAccountItemId sets AccountItemId field to given value.


### SetAccountItemIdNil

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetAccountItemIdNil(b bool)`

 SetAccountItemIdNil sets the value for AccountItemId to be an explicit nil

### UnsetAccountItemId
`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) UnsetAccountItemId()`

UnsetAccountItemId ensures that no value is present for AccountItemId, not even an explicit nil
### GetAccountItemName

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetAccountItemName() string`

GetAccountItemName returns the AccountItemName field if non-nil, zero value otherwise.

### GetAccountItemNameOk

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetAccountItemNameOk() (*string, bool)`

GetAccountItemNameOk returns a tuple with the AccountItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemName

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetAccountItemName(v string)`

SetAccountItemName sets AccountItemName field to given value.


### SetAccountItemNameNil

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetAccountItemNameNil(b bool)`

 SetAccountItemNameNil sets the value for AccountItemName to be an explicit nil

### UnsetAccountItemName
`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) UnsetAccountItemName()`

UnsetAccountItemName ensures that no value is present for AccountItemName, not even an explicit nil
### GetAmount

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetId(v int64)`

SetId sets Id field to given value.


### GetItemId

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetItemId(v int64)`

SetItemId sets ItemId field to given value.


### SetItemIdNil

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetItemName

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetItemName() string`

GetItemName returns the ItemName field if non-nil, zero value otherwise.

### GetItemNameOk

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetItemNameOk() (*string, bool)`

GetItemNameOk returns a tuple with the ItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemName

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetItemName(v string)`

SetItemName sets ItemName field to given value.


### SetItemNameNil

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetItemNameNil(b bool)`

 SetItemNameNil sets the value for ItemName to be an explicit nil

### UnsetItemName
`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) UnsetItemName()`

UnsetItemName ensures that no value is present for ItemName, not even an explicit nil
### GetOrder

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetOrder(v int64)`

SetOrder sets Order field to given value.


### GetQty

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetQty() float32`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetQtyOk() (*float32, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetQty(v float32)`

SetQty sets Qty field to given value.


### GetReducedVat

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetReducedVat() bool`

GetReducedVat returns the ReducedVat field if non-nil, zero value otherwise.

### GetReducedVatOk

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetReducedVatOk() (*bool, bool)`

GetReducedVatOk returns a tuple with the ReducedVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReducedVat

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetReducedVat(v bool)`

SetReducedVat sets ReducedVat field to given value.


### GetSectionId

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetSectionId() int64`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetSectionIdOk() (*int64, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetSectionId(v int64)`

SetSectionId sets SectionId field to given value.


### SetSectionIdNil

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetSectionIdNil(b bool)`

 SetSectionIdNil sets the value for SectionId to be an explicit nil

### UnsetSectionId
`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) UnsetSectionId()`

UnsetSectionId ensures that no value is present for SectionId, not even an explicit nil
### GetSectionName

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetSectionName() string`

GetSectionName returns the SectionName field if non-nil, zero value otherwise.

### GetSectionNameOk

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetSectionNameOk() (*string, bool)`

GetSectionNameOk returns a tuple with the SectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionName

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetSectionName(v string)`

SetSectionName sets SectionName field to given value.


### SetSectionNameNil

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetSectionNameNil(b bool)`

 SetSectionNameNil sets the value for SectionName to be an explicit nil

### UnsetSectionName
`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) UnsetSectionName()`

UnsetSectionName ensures that no value is present for SectionName, not even an explicit nil
### GetSegment1TagId

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetSegment1TagId() int64`

GetSegment1TagId returns the Segment1TagId field if non-nil, zero value otherwise.

### GetSegment1TagIdOk

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetSegment1TagIdOk() (*int64, bool)`

GetSegment1TagIdOk returns a tuple with the Segment1TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagId

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetSegment1TagId(v int64)`

SetSegment1TagId sets Segment1TagId field to given value.

### HasSegment1TagId

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) HasSegment1TagId() bool`

HasSegment1TagId returns a boolean if a field has been set.

### SetSegment1TagIdNil

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetSegment1TagIdNil(b bool)`

 SetSegment1TagIdNil sets the value for Segment1TagId to be an explicit nil

### UnsetSegment1TagId
`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) UnsetSegment1TagId()`

UnsetSegment1TagId ensures that no value is present for Segment1TagId, not even an explicit nil
### GetSegment1TagName

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetSegment1TagName() string`

GetSegment1TagName returns the Segment1TagName field if non-nil, zero value otherwise.

### GetSegment1TagNameOk

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetSegment1TagNameOk() (*string, bool)`

GetSegment1TagNameOk returns a tuple with the Segment1TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagName

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetSegment1TagName(v string)`

SetSegment1TagName sets Segment1TagName field to given value.

### HasSegment1TagName

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) HasSegment1TagName() bool`

HasSegment1TagName returns a boolean if a field has been set.

### SetSegment1TagNameNil

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetSegment1TagNameNil(b bool)`

 SetSegment1TagNameNil sets the value for Segment1TagName to be an explicit nil

### UnsetSegment1TagName
`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) UnsetSegment1TagName()`

UnsetSegment1TagName ensures that no value is present for Segment1TagName, not even an explicit nil
### GetSegment2TagId

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetSegment2TagId() int64`

GetSegment2TagId returns the Segment2TagId field if non-nil, zero value otherwise.

### GetSegment2TagIdOk

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetSegment2TagIdOk() (*int64, bool)`

GetSegment2TagIdOk returns a tuple with the Segment2TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagId

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetSegment2TagId(v int64)`

SetSegment2TagId sets Segment2TagId field to given value.

### HasSegment2TagId

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) HasSegment2TagId() bool`

HasSegment2TagId returns a boolean if a field has been set.

### SetSegment2TagIdNil

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetSegment2TagIdNil(b bool)`

 SetSegment2TagIdNil sets the value for Segment2TagId to be an explicit nil

### UnsetSegment2TagId
`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) UnsetSegment2TagId()`

UnsetSegment2TagId ensures that no value is present for Segment2TagId, not even an explicit nil
### GetSegment2TagName

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetSegment2TagName() string`

GetSegment2TagName returns the Segment2TagName field if non-nil, zero value otherwise.

### GetSegment2TagNameOk

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetSegment2TagNameOk() (*string, bool)`

GetSegment2TagNameOk returns a tuple with the Segment2TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagName

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetSegment2TagName(v string)`

SetSegment2TagName sets Segment2TagName field to given value.

### HasSegment2TagName

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) HasSegment2TagName() bool`

HasSegment2TagName returns a boolean if a field has been set.

### SetSegment2TagNameNil

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetSegment2TagNameNil(b bool)`

 SetSegment2TagNameNil sets the value for Segment2TagName to be an explicit nil

### UnsetSegment2TagName
`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) UnsetSegment2TagName()`

UnsetSegment2TagName ensures that no value is present for Segment2TagName, not even an explicit nil
### GetSegment3TagId

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetSegment3TagId() int64`

GetSegment3TagId returns the Segment3TagId field if non-nil, zero value otherwise.

### GetSegment3TagIdOk

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetSegment3TagIdOk() (*int64, bool)`

GetSegment3TagIdOk returns a tuple with the Segment3TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagId

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetSegment3TagId(v int64)`

SetSegment3TagId sets Segment3TagId field to given value.

### HasSegment3TagId

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) HasSegment3TagId() bool`

HasSegment3TagId returns a boolean if a field has been set.

### SetSegment3TagIdNil

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetSegment3TagIdNil(b bool)`

 SetSegment3TagIdNil sets the value for Segment3TagId to be an explicit nil

### UnsetSegment3TagId
`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) UnsetSegment3TagId()`

UnsetSegment3TagId ensures that no value is present for Segment3TagId, not even an explicit nil
### GetSegment3TagName

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetSegment3TagName() string`

GetSegment3TagName returns the Segment3TagName field if non-nil, zero value otherwise.

### GetSegment3TagNameOk

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetSegment3TagNameOk() (*string, bool)`

GetSegment3TagNameOk returns a tuple with the Segment3TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagName

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetSegment3TagName(v string)`

SetSegment3TagName sets Segment3TagName field to given value.

### HasSegment3TagName

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) HasSegment3TagName() bool`

HasSegment3TagName returns a boolean if a field has been set.

### SetSegment3TagNameNil

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetSegment3TagNameNil(b bool)`

 SetSegment3TagNameNil sets the value for Segment3TagName to be an explicit nil

### UnsetSegment3TagName
`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) UnsetSegment3TagName()`

UnsetSegment3TagName ensures that no value is present for Segment3TagName, not even an explicit nil
### GetTagIds

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetTagIds() []int64`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetTagIdsOk() (*[]int64, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetTagIds(v []int64)`

SetTagIds sets TagIds field to given value.


### GetTagNames

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.


### GetTaxCode

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetTaxCode() int64`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetTaxCodeOk() (*int64, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetTaxCode(v int64)`

SetTaxCode sets TaxCode field to given value.


### SetTaxCodeNil

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetTaxCodeNil(b bool)`

 SetTaxCodeNil sets the value for TaxCode to be an explicit nil

### UnsetTaxCode
`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) UnsetTaxCode()`

UnsetTaxCode ensures that no value is present for TaxCode, not even an explicit nil
### GetType

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetType(v string)`

SetType sets Type field to given value.


### GetUnit

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetUnit(v string)`

SetUnit sets Unit field to given value.


### SetUnitNil

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetUnitPrice

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.


### GetVat

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetVat() int64`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) GetVatOk() (*int64, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *QuotationIndexResponseQuotationsInnerQuotationContentsInner) SetVat(v int64)`

SetVat sets Vat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


