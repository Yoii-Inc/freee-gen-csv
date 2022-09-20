# ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | 金額 | [optional] 
**Description** | Pointer to **string** | 内容 | [optional] 
**ExpenseApplicationLineTemplateId** | Pointer to **int32** | 経費科目ID | [optional] 
**Id** | **int64** | 経費申請の項目行ID | 
**ReceiptId** | Pointer to **int32** | 証憑ファイルID（ファイルボックスのファイルID） | [optional] 
**TransactionDate** | Pointer to **string** | 日付 (yyyy-mm-dd) | [optional] 

## Methods

### NewExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner

`func NewExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner(id int64, ) *ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner`

NewExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner instantiates a new ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInnerWithDefaults

`func NewExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInnerWithDefaults() *ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner`

NewExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInnerWithDefaults instantiates a new ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDescription

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpenseApplicationLineTemplateId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner) GetExpenseApplicationLineTemplateId() int32`

GetExpenseApplicationLineTemplateId returns the ExpenseApplicationLineTemplateId field if non-nil, zero value otherwise.

### GetExpenseApplicationLineTemplateIdOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner) GetExpenseApplicationLineTemplateIdOk() (*int32, bool)`

GetExpenseApplicationLineTemplateIdOk returns a tuple with the ExpenseApplicationLineTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseApplicationLineTemplateId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner) SetExpenseApplicationLineTemplateId(v int32)`

SetExpenseApplicationLineTemplateId sets ExpenseApplicationLineTemplateId field to given value.

### HasExpenseApplicationLineTemplateId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner) HasExpenseApplicationLineTemplateId() bool`

HasExpenseApplicationLineTemplateId returns a boolean if a field has been set.

### GetId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner) SetId(v int64)`

SetId sets Id field to given value.


### GetReceiptId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner) GetReceiptId() int32`

GetReceiptId returns the ReceiptId field if non-nil, zero value otherwise.

### GetReceiptIdOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner) GetReceiptIdOk() (*int32, bool)`

GetReceiptIdOk returns a tuple with the ReceiptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner) SetReceiptId(v int32)`

SetReceiptId sets ReceiptId field to given value.

### HasReceiptId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner) HasReceiptId() bool`

HasReceiptId returns a boolean if a field has been set.

### GetTransactionDate

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


