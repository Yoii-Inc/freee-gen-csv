# ExpenseApplicationLineTemplateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountItemId** | **int64** | 勘定科目ID | 
**CompanyId** | **int64** | 事業所ID | 
**Description** | Pointer to **string** | 経費科目の説明 (1000文字以内) | [optional] 
**ItemId** | Pointer to **int64** | 品目ID | [optional] 
**LineDescription** | Pointer to **string** | 内容の補足 (1000文字以内) | [optional] 
**Name** | **string** | 経費科目名 (100文字以内) | 
**RequiredReceipt** | Pointer to **bool** | 添付ファイルの必須/任意 | [optional] 
**TaxCode** | **int64** | 税区分コード（税区分のdisplay_categoryがtax_5: 5%表示の税区分, tax_r8: 軽減税率8%表示の税区分に該当するtax_codeのみ利用可能です。税区分のdisplay_categoryは /taxes/companies/{:company_id}のAPIから取得可能です。） | 

## Methods

### NewExpenseApplicationLineTemplateParams

`func NewExpenseApplicationLineTemplateParams(accountItemId int64, companyId int64, name string, taxCode int64, ) *ExpenseApplicationLineTemplateParams`

NewExpenseApplicationLineTemplateParams instantiates a new ExpenseApplicationLineTemplateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseApplicationLineTemplateParamsWithDefaults

`func NewExpenseApplicationLineTemplateParamsWithDefaults() *ExpenseApplicationLineTemplateParams`

NewExpenseApplicationLineTemplateParamsWithDefaults instantiates a new ExpenseApplicationLineTemplateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountItemId

`func (o *ExpenseApplicationLineTemplateParams) GetAccountItemId() int64`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *ExpenseApplicationLineTemplateParams) GetAccountItemIdOk() (*int64, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *ExpenseApplicationLineTemplateParams) SetAccountItemId(v int64)`

SetAccountItemId sets AccountItemId field to given value.


### GetCompanyId

`func (o *ExpenseApplicationLineTemplateParams) GetCompanyId() int64`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ExpenseApplicationLineTemplateParams) GetCompanyIdOk() (*int64, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ExpenseApplicationLineTemplateParams) SetCompanyId(v int64)`

SetCompanyId sets CompanyId field to given value.


### GetDescription

`func (o *ExpenseApplicationLineTemplateParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExpenseApplicationLineTemplateParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExpenseApplicationLineTemplateParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExpenseApplicationLineTemplateParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetItemId

`func (o *ExpenseApplicationLineTemplateParams) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ExpenseApplicationLineTemplateParams) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ExpenseApplicationLineTemplateParams) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ExpenseApplicationLineTemplateParams) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetLineDescription

`func (o *ExpenseApplicationLineTemplateParams) GetLineDescription() string`

GetLineDescription returns the LineDescription field if non-nil, zero value otherwise.

### GetLineDescriptionOk

`func (o *ExpenseApplicationLineTemplateParams) GetLineDescriptionOk() (*string, bool)`

GetLineDescriptionOk returns a tuple with the LineDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineDescription

`func (o *ExpenseApplicationLineTemplateParams) SetLineDescription(v string)`

SetLineDescription sets LineDescription field to given value.

### HasLineDescription

`func (o *ExpenseApplicationLineTemplateParams) HasLineDescription() bool`

HasLineDescription returns a boolean if a field has been set.

### GetName

`func (o *ExpenseApplicationLineTemplateParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExpenseApplicationLineTemplateParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExpenseApplicationLineTemplateParams) SetName(v string)`

SetName sets Name field to given value.


### GetRequiredReceipt

`func (o *ExpenseApplicationLineTemplateParams) GetRequiredReceipt() bool`

GetRequiredReceipt returns the RequiredReceipt field if non-nil, zero value otherwise.

### GetRequiredReceiptOk

`func (o *ExpenseApplicationLineTemplateParams) GetRequiredReceiptOk() (*bool, bool)`

GetRequiredReceiptOk returns a tuple with the RequiredReceipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredReceipt

`func (o *ExpenseApplicationLineTemplateParams) SetRequiredReceipt(v bool)`

SetRequiredReceipt sets RequiredReceipt field to given value.

### HasRequiredReceipt

`func (o *ExpenseApplicationLineTemplateParams) HasRequiredReceipt() bool`

HasRequiredReceipt returns a boolean if a field has been set.

### GetTaxCode

`func (o *ExpenseApplicationLineTemplateParams) GetTaxCode() int64`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *ExpenseApplicationLineTemplateParams) GetTaxCodeOk() (*int64, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *ExpenseApplicationLineTemplateParams) SetTaxCode(v int64)`

SetTaxCode sets TaxCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


