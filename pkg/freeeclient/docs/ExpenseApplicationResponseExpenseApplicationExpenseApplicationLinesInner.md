# ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | 金額 | [optional] 
**Description** | Pointer to **NullableString** | 内容 | [optional] 
**ExpenseApplicationLineTemplateId** | Pointer to **NullableInt32** | 経費科目ID | [optional] 
**Id** | **int64** | 経費申請の項目行ID | 
**ReceiptId** | Pointer to **NullableInt32** | 証憑ファイルID（ファイルボックスのファイルID） | [optional] 
**TransactionDate** | Pointer to **NullableString** | 日付 (yyyy-mm-dd) | [optional] 

## Methods

### NewExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner

`func NewExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner(id int64, ) *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner`

NewExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner instantiates a new ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInnerWithDefaults

`func NewExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInnerWithDefaults() *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner`

NewExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInnerWithDefaults instantiates a new ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDescription

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExpenseApplicationLineTemplateId

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) GetExpenseApplicationLineTemplateId() int32`

GetExpenseApplicationLineTemplateId returns the ExpenseApplicationLineTemplateId field if non-nil, zero value otherwise.

### GetExpenseApplicationLineTemplateIdOk

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) GetExpenseApplicationLineTemplateIdOk() (*int32, bool)`

GetExpenseApplicationLineTemplateIdOk returns a tuple with the ExpenseApplicationLineTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseApplicationLineTemplateId

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) SetExpenseApplicationLineTemplateId(v int32)`

SetExpenseApplicationLineTemplateId sets ExpenseApplicationLineTemplateId field to given value.

### HasExpenseApplicationLineTemplateId

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) HasExpenseApplicationLineTemplateId() bool`

HasExpenseApplicationLineTemplateId returns a boolean if a field has been set.

### SetExpenseApplicationLineTemplateIdNil

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) SetExpenseApplicationLineTemplateIdNil(b bool)`

 SetExpenseApplicationLineTemplateIdNil sets the value for ExpenseApplicationLineTemplateId to be an explicit nil

### UnsetExpenseApplicationLineTemplateId
`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) UnsetExpenseApplicationLineTemplateId()`

UnsetExpenseApplicationLineTemplateId ensures that no value is present for ExpenseApplicationLineTemplateId, not even an explicit nil
### GetId

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) SetId(v int64)`

SetId sets Id field to given value.


### GetReceiptId

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) GetReceiptId() int32`

GetReceiptId returns the ReceiptId field if non-nil, zero value otherwise.

### GetReceiptIdOk

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) GetReceiptIdOk() (*int32, bool)`

GetReceiptIdOk returns a tuple with the ReceiptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptId

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) SetReceiptId(v int32)`

SetReceiptId sets ReceiptId field to given value.

### HasReceiptId

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) HasReceiptId() bool`

HasReceiptId returns a boolean if a field has been set.

### SetReceiptIdNil

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) SetReceiptIdNil(b bool)`

 SetReceiptIdNil sets the value for ReceiptId to be an explicit nil

### UnsetReceiptId
`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) UnsetReceiptId()`

UnsetReceiptId ensures that no value is present for ReceiptId, not even an explicit nil
### GetTransactionDate

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### SetTransactionDateNil

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) SetTransactionDateNil(b bool)`

 SetTransactionDateNil sets the value for TransactionDate to be an explicit nil

### UnsetTransactionDate
`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner) UnsetTransactionDate()`

UnsetTransactionDate ensures that no value is present for TransactionDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


