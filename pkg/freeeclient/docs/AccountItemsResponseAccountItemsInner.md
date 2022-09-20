# AccountItemsResponseAccountItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountCategory** | **string** | 勘定科目カテゴリー | 
**AccountCategoryId** | **int64** | 勘定科目のカテゴリーID | 
**Available** | **bool** | 勘定科目の使用設定（true: 使用する、false: 使用しない） | 
**Categories** | **[]string** |  | 
**CorrespondingExpenseId** | Pointer to **NullableInt64** | 支出取引相手勘定科目ID | [optional] 
**CorrespondingExpenseName** | Pointer to **NullableString** | 支出取引相手勘定科目名 | [optional] 
**CorrespondingIncomeId** | Pointer to **NullableInt64** | 収入取引相手勘定科目ID | [optional] 
**CorrespondingIncomeName** | Pointer to **NullableString** | 収入取引相手勘定科目名 | [optional] 
**DefaultTaxCode** | **int64** | デフォルト設定がされている税区分コード | 
**DefaultTaxId** | Pointer to **int64** | デフォルト設定がされている税区分ID | [optional] 
**GroupName** | Pointer to **NullableString** | 決算書表示名（小カテゴリー） | [optional] 
**Id** | **int64** | 勘定科目ID | 
**Name** | **string** | 勘定科目名 (30文字以内) | 
**Shortcut** | Pointer to **NullableString** | ショートカット1 (20文字以内) | [optional] 
**ShortcutNum** | Pointer to **NullableString** | ショートカット2(勘定科目コード) (20文字以内) | [optional] 
**TaxCode** | **int64** | 税区分コード | 
**WalletableId** | **NullableInt64** | 口座ID | 

## Methods

### NewAccountItemsResponseAccountItemsInner

`func NewAccountItemsResponseAccountItemsInner(accountCategory string, accountCategoryId int64, available bool, categories []string, defaultTaxCode int64, id int64, name string, taxCode int64, walletableId NullableInt64, ) *AccountItemsResponseAccountItemsInner`

NewAccountItemsResponseAccountItemsInner instantiates a new AccountItemsResponseAccountItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountItemsResponseAccountItemsInnerWithDefaults

`func NewAccountItemsResponseAccountItemsInnerWithDefaults() *AccountItemsResponseAccountItemsInner`

NewAccountItemsResponseAccountItemsInnerWithDefaults instantiates a new AccountItemsResponseAccountItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountCategory

`func (o *AccountItemsResponseAccountItemsInner) GetAccountCategory() string`

GetAccountCategory returns the AccountCategory field if non-nil, zero value otherwise.

### GetAccountCategoryOk

`func (o *AccountItemsResponseAccountItemsInner) GetAccountCategoryOk() (*string, bool)`

GetAccountCategoryOk returns a tuple with the AccountCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCategory

`func (o *AccountItemsResponseAccountItemsInner) SetAccountCategory(v string)`

SetAccountCategory sets AccountCategory field to given value.


### GetAccountCategoryId

`func (o *AccountItemsResponseAccountItemsInner) GetAccountCategoryId() int64`

GetAccountCategoryId returns the AccountCategoryId field if non-nil, zero value otherwise.

### GetAccountCategoryIdOk

`func (o *AccountItemsResponseAccountItemsInner) GetAccountCategoryIdOk() (*int64, bool)`

GetAccountCategoryIdOk returns a tuple with the AccountCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCategoryId

`func (o *AccountItemsResponseAccountItemsInner) SetAccountCategoryId(v int64)`

SetAccountCategoryId sets AccountCategoryId field to given value.


### GetAvailable

`func (o *AccountItemsResponseAccountItemsInner) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *AccountItemsResponseAccountItemsInner) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *AccountItemsResponseAccountItemsInner) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetCategories

`func (o *AccountItemsResponseAccountItemsInner) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *AccountItemsResponseAccountItemsInner) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *AccountItemsResponseAccountItemsInner) SetCategories(v []string)`

SetCategories sets Categories field to given value.


### GetCorrespondingExpenseId

`func (o *AccountItemsResponseAccountItemsInner) GetCorrespondingExpenseId() int64`

GetCorrespondingExpenseId returns the CorrespondingExpenseId field if non-nil, zero value otherwise.

### GetCorrespondingExpenseIdOk

`func (o *AccountItemsResponseAccountItemsInner) GetCorrespondingExpenseIdOk() (*int64, bool)`

GetCorrespondingExpenseIdOk returns a tuple with the CorrespondingExpenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrespondingExpenseId

`func (o *AccountItemsResponseAccountItemsInner) SetCorrespondingExpenseId(v int64)`

SetCorrespondingExpenseId sets CorrespondingExpenseId field to given value.

### HasCorrespondingExpenseId

`func (o *AccountItemsResponseAccountItemsInner) HasCorrespondingExpenseId() bool`

HasCorrespondingExpenseId returns a boolean if a field has been set.

### SetCorrespondingExpenseIdNil

`func (o *AccountItemsResponseAccountItemsInner) SetCorrespondingExpenseIdNil(b bool)`

 SetCorrespondingExpenseIdNil sets the value for CorrespondingExpenseId to be an explicit nil

### UnsetCorrespondingExpenseId
`func (o *AccountItemsResponseAccountItemsInner) UnsetCorrespondingExpenseId()`

UnsetCorrespondingExpenseId ensures that no value is present for CorrespondingExpenseId, not even an explicit nil
### GetCorrespondingExpenseName

`func (o *AccountItemsResponseAccountItemsInner) GetCorrespondingExpenseName() string`

GetCorrespondingExpenseName returns the CorrespondingExpenseName field if non-nil, zero value otherwise.

### GetCorrespondingExpenseNameOk

`func (o *AccountItemsResponseAccountItemsInner) GetCorrespondingExpenseNameOk() (*string, bool)`

GetCorrespondingExpenseNameOk returns a tuple with the CorrespondingExpenseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrespondingExpenseName

`func (o *AccountItemsResponseAccountItemsInner) SetCorrespondingExpenseName(v string)`

SetCorrespondingExpenseName sets CorrespondingExpenseName field to given value.

### HasCorrespondingExpenseName

`func (o *AccountItemsResponseAccountItemsInner) HasCorrespondingExpenseName() bool`

HasCorrespondingExpenseName returns a boolean if a field has been set.

### SetCorrespondingExpenseNameNil

`func (o *AccountItemsResponseAccountItemsInner) SetCorrespondingExpenseNameNil(b bool)`

 SetCorrespondingExpenseNameNil sets the value for CorrespondingExpenseName to be an explicit nil

### UnsetCorrespondingExpenseName
`func (o *AccountItemsResponseAccountItemsInner) UnsetCorrespondingExpenseName()`

UnsetCorrespondingExpenseName ensures that no value is present for CorrespondingExpenseName, not even an explicit nil
### GetCorrespondingIncomeId

`func (o *AccountItemsResponseAccountItemsInner) GetCorrespondingIncomeId() int64`

GetCorrespondingIncomeId returns the CorrespondingIncomeId field if non-nil, zero value otherwise.

### GetCorrespondingIncomeIdOk

`func (o *AccountItemsResponseAccountItemsInner) GetCorrespondingIncomeIdOk() (*int64, bool)`

GetCorrespondingIncomeIdOk returns a tuple with the CorrespondingIncomeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrespondingIncomeId

`func (o *AccountItemsResponseAccountItemsInner) SetCorrespondingIncomeId(v int64)`

SetCorrespondingIncomeId sets CorrespondingIncomeId field to given value.

### HasCorrespondingIncomeId

`func (o *AccountItemsResponseAccountItemsInner) HasCorrespondingIncomeId() bool`

HasCorrespondingIncomeId returns a boolean if a field has been set.

### SetCorrespondingIncomeIdNil

`func (o *AccountItemsResponseAccountItemsInner) SetCorrespondingIncomeIdNil(b bool)`

 SetCorrespondingIncomeIdNil sets the value for CorrespondingIncomeId to be an explicit nil

### UnsetCorrespondingIncomeId
`func (o *AccountItemsResponseAccountItemsInner) UnsetCorrespondingIncomeId()`

UnsetCorrespondingIncomeId ensures that no value is present for CorrespondingIncomeId, not even an explicit nil
### GetCorrespondingIncomeName

`func (o *AccountItemsResponseAccountItemsInner) GetCorrespondingIncomeName() string`

GetCorrespondingIncomeName returns the CorrespondingIncomeName field if non-nil, zero value otherwise.

### GetCorrespondingIncomeNameOk

`func (o *AccountItemsResponseAccountItemsInner) GetCorrespondingIncomeNameOk() (*string, bool)`

GetCorrespondingIncomeNameOk returns a tuple with the CorrespondingIncomeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrespondingIncomeName

`func (o *AccountItemsResponseAccountItemsInner) SetCorrespondingIncomeName(v string)`

SetCorrespondingIncomeName sets CorrespondingIncomeName field to given value.

### HasCorrespondingIncomeName

`func (o *AccountItemsResponseAccountItemsInner) HasCorrespondingIncomeName() bool`

HasCorrespondingIncomeName returns a boolean if a field has been set.

### SetCorrespondingIncomeNameNil

`func (o *AccountItemsResponseAccountItemsInner) SetCorrespondingIncomeNameNil(b bool)`

 SetCorrespondingIncomeNameNil sets the value for CorrespondingIncomeName to be an explicit nil

### UnsetCorrespondingIncomeName
`func (o *AccountItemsResponseAccountItemsInner) UnsetCorrespondingIncomeName()`

UnsetCorrespondingIncomeName ensures that no value is present for CorrespondingIncomeName, not even an explicit nil
### GetDefaultTaxCode

`func (o *AccountItemsResponseAccountItemsInner) GetDefaultTaxCode() int64`

GetDefaultTaxCode returns the DefaultTaxCode field if non-nil, zero value otherwise.

### GetDefaultTaxCodeOk

`func (o *AccountItemsResponseAccountItemsInner) GetDefaultTaxCodeOk() (*int64, bool)`

GetDefaultTaxCodeOk returns a tuple with the DefaultTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTaxCode

`func (o *AccountItemsResponseAccountItemsInner) SetDefaultTaxCode(v int64)`

SetDefaultTaxCode sets DefaultTaxCode field to given value.


### GetDefaultTaxId

`func (o *AccountItemsResponseAccountItemsInner) GetDefaultTaxId() int64`

GetDefaultTaxId returns the DefaultTaxId field if non-nil, zero value otherwise.

### GetDefaultTaxIdOk

`func (o *AccountItemsResponseAccountItemsInner) GetDefaultTaxIdOk() (*int64, bool)`

GetDefaultTaxIdOk returns a tuple with the DefaultTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTaxId

`func (o *AccountItemsResponseAccountItemsInner) SetDefaultTaxId(v int64)`

SetDefaultTaxId sets DefaultTaxId field to given value.

### HasDefaultTaxId

`func (o *AccountItemsResponseAccountItemsInner) HasDefaultTaxId() bool`

HasDefaultTaxId returns a boolean if a field has been set.

### GetGroupName

`func (o *AccountItemsResponseAccountItemsInner) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *AccountItemsResponseAccountItemsInner) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *AccountItemsResponseAccountItemsInner) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *AccountItemsResponseAccountItemsInner) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *AccountItemsResponseAccountItemsInner) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *AccountItemsResponseAccountItemsInner) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetId

`func (o *AccountItemsResponseAccountItemsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountItemsResponseAccountItemsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountItemsResponseAccountItemsInner) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *AccountItemsResponseAccountItemsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountItemsResponseAccountItemsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountItemsResponseAccountItemsInner) SetName(v string)`

SetName sets Name field to given value.


### GetShortcut

`func (o *AccountItemsResponseAccountItemsInner) GetShortcut() string`

GetShortcut returns the Shortcut field if non-nil, zero value otherwise.

### GetShortcutOk

`func (o *AccountItemsResponseAccountItemsInner) GetShortcutOk() (*string, bool)`

GetShortcutOk returns a tuple with the Shortcut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut

`func (o *AccountItemsResponseAccountItemsInner) SetShortcut(v string)`

SetShortcut sets Shortcut field to given value.

### HasShortcut

`func (o *AccountItemsResponseAccountItemsInner) HasShortcut() bool`

HasShortcut returns a boolean if a field has been set.

### SetShortcutNil

`func (o *AccountItemsResponseAccountItemsInner) SetShortcutNil(b bool)`

 SetShortcutNil sets the value for Shortcut to be an explicit nil

### UnsetShortcut
`func (o *AccountItemsResponseAccountItemsInner) UnsetShortcut()`

UnsetShortcut ensures that no value is present for Shortcut, not even an explicit nil
### GetShortcutNum

`func (o *AccountItemsResponseAccountItemsInner) GetShortcutNum() string`

GetShortcutNum returns the ShortcutNum field if non-nil, zero value otherwise.

### GetShortcutNumOk

`func (o *AccountItemsResponseAccountItemsInner) GetShortcutNumOk() (*string, bool)`

GetShortcutNumOk returns a tuple with the ShortcutNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcutNum

`func (o *AccountItemsResponseAccountItemsInner) SetShortcutNum(v string)`

SetShortcutNum sets ShortcutNum field to given value.

### HasShortcutNum

`func (o *AccountItemsResponseAccountItemsInner) HasShortcutNum() bool`

HasShortcutNum returns a boolean if a field has been set.

### SetShortcutNumNil

`func (o *AccountItemsResponseAccountItemsInner) SetShortcutNumNil(b bool)`

 SetShortcutNumNil sets the value for ShortcutNum to be an explicit nil

### UnsetShortcutNum
`func (o *AccountItemsResponseAccountItemsInner) UnsetShortcutNum()`

UnsetShortcutNum ensures that no value is present for ShortcutNum, not even an explicit nil
### GetTaxCode

`func (o *AccountItemsResponseAccountItemsInner) GetTaxCode() int64`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *AccountItemsResponseAccountItemsInner) GetTaxCodeOk() (*int64, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *AccountItemsResponseAccountItemsInner) SetTaxCode(v int64)`

SetTaxCode sets TaxCode field to given value.


### GetWalletableId

`func (o *AccountItemsResponseAccountItemsInner) GetWalletableId() int64`

GetWalletableId returns the WalletableId field if non-nil, zero value otherwise.

### GetWalletableIdOk

`func (o *AccountItemsResponseAccountItemsInner) GetWalletableIdOk() (*int64, bool)`

GetWalletableIdOk returns a tuple with the WalletableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletableId

`func (o *AccountItemsResponseAccountItemsInner) SetWalletableId(v int64)`

SetWalletableId sets WalletableId field to given value.


### SetWalletableIdNil

`func (o *AccountItemsResponseAccountItemsInner) SetWalletableIdNil(b bool)`

 SetWalletableIdNil sets the value for WalletableId to be an explicit nil

### UnsetWalletableId
`func (o *AccountItemsResponseAccountItemsInner) UnsetWalletableId()`

UnsetWalletableId ensures that no value is present for WalletableId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


