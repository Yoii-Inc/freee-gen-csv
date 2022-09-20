# ExpenseApplicationUpdateParamsExpenseApplicationLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | 金額 | [optional] 
**Description** | Pointer to **string** | 内容 (250文字以内) | [optional] 
**ExpenseApplicationLineTemplateId** | Pointer to **int64** | 経費科目ID | [optional] 
**Id** | Pointer to **int64** | 経費申請の項目行ID: 既存項目行を更新する場合に指定します。IDを指定しない項目行は、新規行として扱われ追加されます。また、expense_application_linesに含まれない既存の項目行は削除されます。更新後も残したい行は、必ず経費申請の項目行IDを指定してexpense_application_linesに含めてください。 | [optional] 
**ReceiptId** | Pointer to **int64** | 証憑ファイルID（ファイルボックスのファイルID） | [optional] 
**TransactionDate** | Pointer to **string** | 日付 (yyyy-mm-dd) | [optional] 

## Methods

### NewExpenseApplicationUpdateParamsExpenseApplicationLinesInner

`func NewExpenseApplicationUpdateParamsExpenseApplicationLinesInner() *ExpenseApplicationUpdateParamsExpenseApplicationLinesInner`

NewExpenseApplicationUpdateParamsExpenseApplicationLinesInner instantiates a new ExpenseApplicationUpdateParamsExpenseApplicationLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseApplicationUpdateParamsExpenseApplicationLinesInnerWithDefaults

`func NewExpenseApplicationUpdateParamsExpenseApplicationLinesInnerWithDefaults() *ExpenseApplicationUpdateParamsExpenseApplicationLinesInner`

NewExpenseApplicationUpdateParamsExpenseApplicationLinesInnerWithDefaults instantiates a new ExpenseApplicationUpdateParamsExpenseApplicationLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLinesInner) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLinesInner) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLinesInner) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLinesInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDescription

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLinesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLinesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLinesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLinesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpenseApplicationLineTemplateId

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLinesInner) GetExpenseApplicationLineTemplateId() int64`

GetExpenseApplicationLineTemplateId returns the ExpenseApplicationLineTemplateId field if non-nil, zero value otherwise.

### GetExpenseApplicationLineTemplateIdOk

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLinesInner) GetExpenseApplicationLineTemplateIdOk() (*int64, bool)`

GetExpenseApplicationLineTemplateIdOk returns a tuple with the ExpenseApplicationLineTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseApplicationLineTemplateId

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLinesInner) SetExpenseApplicationLineTemplateId(v int64)`

SetExpenseApplicationLineTemplateId sets ExpenseApplicationLineTemplateId field to given value.

### HasExpenseApplicationLineTemplateId

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLinesInner) HasExpenseApplicationLineTemplateId() bool`

HasExpenseApplicationLineTemplateId returns a boolean if a field has been set.

### GetId

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLinesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLinesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLinesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLinesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReceiptId

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLinesInner) GetReceiptId() int64`

GetReceiptId returns the ReceiptId field if non-nil, zero value otherwise.

### GetReceiptIdOk

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLinesInner) GetReceiptIdOk() (*int64, bool)`

GetReceiptIdOk returns a tuple with the ReceiptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptId

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLinesInner) SetReceiptId(v int64)`

SetReceiptId sets ReceiptId field to given value.

### HasReceiptId

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLinesInner) HasReceiptId() bool`

HasReceiptId returns a boolean if a field has been set.

### GetTransactionDate

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLinesInner) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLinesInner) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLinesInner) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLinesInner) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


