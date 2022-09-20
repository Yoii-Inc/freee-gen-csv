# ExpenseApplicationCreateParamsExpenseApplicationLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | 金額 | [optional] 
**Description** | Pointer to **string** | 内容 (250文字以内) | [optional] 
**ExpenseApplicationLineTemplateId** | Pointer to **int32** | 経費科目ID | [optional] 
**ReceiptId** | Pointer to **int32** | 証憑ファイルID（ファイルボックスのファイルID） | [optional] 
**TransactionDate** | Pointer to **string** | 日付 (yyyy-mm-dd) | [optional] 

## Methods

### NewExpenseApplicationCreateParamsExpenseApplicationLinesInner

`func NewExpenseApplicationCreateParamsExpenseApplicationLinesInner() *ExpenseApplicationCreateParamsExpenseApplicationLinesInner`

NewExpenseApplicationCreateParamsExpenseApplicationLinesInner instantiates a new ExpenseApplicationCreateParamsExpenseApplicationLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseApplicationCreateParamsExpenseApplicationLinesInnerWithDefaults

`func NewExpenseApplicationCreateParamsExpenseApplicationLinesInnerWithDefaults() *ExpenseApplicationCreateParamsExpenseApplicationLinesInner`

NewExpenseApplicationCreateParamsExpenseApplicationLinesInnerWithDefaults instantiates a new ExpenseApplicationCreateParamsExpenseApplicationLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLinesInner) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLinesInner) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLinesInner) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLinesInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDescription

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLinesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLinesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLinesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLinesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpenseApplicationLineTemplateId

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLinesInner) GetExpenseApplicationLineTemplateId() int32`

GetExpenseApplicationLineTemplateId returns the ExpenseApplicationLineTemplateId field if non-nil, zero value otherwise.

### GetExpenseApplicationLineTemplateIdOk

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLinesInner) GetExpenseApplicationLineTemplateIdOk() (*int32, bool)`

GetExpenseApplicationLineTemplateIdOk returns a tuple with the ExpenseApplicationLineTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseApplicationLineTemplateId

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLinesInner) SetExpenseApplicationLineTemplateId(v int32)`

SetExpenseApplicationLineTemplateId sets ExpenseApplicationLineTemplateId field to given value.

### HasExpenseApplicationLineTemplateId

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLinesInner) HasExpenseApplicationLineTemplateId() bool`

HasExpenseApplicationLineTemplateId returns a boolean if a field has been set.

### GetReceiptId

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLinesInner) GetReceiptId() int32`

GetReceiptId returns the ReceiptId field if non-nil, zero value otherwise.

### GetReceiptIdOk

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLinesInner) GetReceiptIdOk() (*int32, bool)`

GetReceiptIdOk returns a tuple with the ReceiptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptId

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLinesInner) SetReceiptId(v int32)`

SetReceiptId sets ReceiptId field to given value.

### HasReceiptId

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLinesInner) HasReceiptId() bool`

HasReceiptId returns a boolean if a field has been set.

### GetTransactionDate

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLinesInner) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLinesInner) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLinesInner) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLinesInner) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


