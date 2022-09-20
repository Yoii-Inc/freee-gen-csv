# ExpenseApplicationLineTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountItemId** | Pointer to **int64** | 勘定科目ID | [optional] 
**AccountItemName** | **string** | 勘定科目名 | 
**Description** | Pointer to **string** | 経費科目の説明 | [optional] 
**Id** | **int64** | 経費科目ID | 
**LineDescription** | Pointer to **string** | 内容の補足 | [optional] 
**Name** | **string** | 経費科目名 | 
**RequiredReceipt** | Pointer to **bool** | 添付ファイルの必須/任意 | [optional] 
**TaxCode** | Pointer to **int64** | 税区分コード | [optional] 
**TaxName** | **string** | 税区分名 | 

## Methods

### NewExpenseApplicationLineTemplate

`func NewExpenseApplicationLineTemplate(accountItemName string, id int64, name string, taxName string, ) *ExpenseApplicationLineTemplate`

NewExpenseApplicationLineTemplate instantiates a new ExpenseApplicationLineTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseApplicationLineTemplateWithDefaults

`func NewExpenseApplicationLineTemplateWithDefaults() *ExpenseApplicationLineTemplate`

NewExpenseApplicationLineTemplateWithDefaults instantiates a new ExpenseApplicationLineTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountItemId

`func (o *ExpenseApplicationLineTemplate) GetAccountItemId() int64`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *ExpenseApplicationLineTemplate) GetAccountItemIdOk() (*int64, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *ExpenseApplicationLineTemplate) SetAccountItemId(v int64)`

SetAccountItemId sets AccountItemId field to given value.

### HasAccountItemId

`func (o *ExpenseApplicationLineTemplate) HasAccountItemId() bool`

HasAccountItemId returns a boolean if a field has been set.

### GetAccountItemName

`func (o *ExpenseApplicationLineTemplate) GetAccountItemName() string`

GetAccountItemName returns the AccountItemName field if non-nil, zero value otherwise.

### GetAccountItemNameOk

`func (o *ExpenseApplicationLineTemplate) GetAccountItemNameOk() (*string, bool)`

GetAccountItemNameOk returns a tuple with the AccountItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemName

`func (o *ExpenseApplicationLineTemplate) SetAccountItemName(v string)`

SetAccountItemName sets AccountItemName field to given value.


### GetDescription

`func (o *ExpenseApplicationLineTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExpenseApplicationLineTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExpenseApplicationLineTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExpenseApplicationLineTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ExpenseApplicationLineTemplate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExpenseApplicationLineTemplate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExpenseApplicationLineTemplate) SetId(v int64)`

SetId sets Id field to given value.


### GetLineDescription

`func (o *ExpenseApplicationLineTemplate) GetLineDescription() string`

GetLineDescription returns the LineDescription field if non-nil, zero value otherwise.

### GetLineDescriptionOk

`func (o *ExpenseApplicationLineTemplate) GetLineDescriptionOk() (*string, bool)`

GetLineDescriptionOk returns a tuple with the LineDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineDescription

`func (o *ExpenseApplicationLineTemplate) SetLineDescription(v string)`

SetLineDescription sets LineDescription field to given value.

### HasLineDescription

`func (o *ExpenseApplicationLineTemplate) HasLineDescription() bool`

HasLineDescription returns a boolean if a field has been set.

### GetName

`func (o *ExpenseApplicationLineTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExpenseApplicationLineTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExpenseApplicationLineTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetRequiredReceipt

`func (o *ExpenseApplicationLineTemplate) GetRequiredReceipt() bool`

GetRequiredReceipt returns the RequiredReceipt field if non-nil, zero value otherwise.

### GetRequiredReceiptOk

`func (o *ExpenseApplicationLineTemplate) GetRequiredReceiptOk() (*bool, bool)`

GetRequiredReceiptOk returns a tuple with the RequiredReceipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredReceipt

`func (o *ExpenseApplicationLineTemplate) SetRequiredReceipt(v bool)`

SetRequiredReceipt sets RequiredReceipt field to given value.

### HasRequiredReceipt

`func (o *ExpenseApplicationLineTemplate) HasRequiredReceipt() bool`

HasRequiredReceipt returns a boolean if a field has been set.

### GetTaxCode

`func (o *ExpenseApplicationLineTemplate) GetTaxCode() int64`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *ExpenseApplicationLineTemplate) GetTaxCodeOk() (*int64, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *ExpenseApplicationLineTemplate) SetTaxCode(v int64)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *ExpenseApplicationLineTemplate) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetTaxName

`func (o *ExpenseApplicationLineTemplate) GetTaxName() string`

GetTaxName returns the TaxName field if non-nil, zero value otherwise.

### GetTaxNameOk

`func (o *ExpenseApplicationLineTemplate) GetTaxNameOk() (*string, bool)`

GetTaxNameOk returns a tuple with the TaxName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName

`func (o *ExpenseApplicationLineTemplate) SetTaxName(v string)`

SetTaxName sets TaxName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


