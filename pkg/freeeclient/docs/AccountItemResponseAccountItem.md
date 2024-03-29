# AccountItemResponseAccountItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountCategory** | **string** | 勘定科目カテゴリー | 
**AccountCategoryId** | **int64** | 勘定科目のカテゴリーID | 
**AccumulatedDepAccountItemId** | Pointer to **NullableInt64** | 減価償却累計額勘定科目ID（法人のみ利用可能） | [optional] 
**AccumulatedDepAccountItemName** | Pointer to **string** | 減価償却累計額勘定科目（法人のみ利用可能） | [optional] 
**Available** | **bool** | 勘定科目の使用設定（true: 使用する、false: 使用しない） | 
**CompanyId** | **int64** | 事業所ID | 
**CorrespondingExpenseId** | Pointer to **NullableInt64** | 支出取引相手勘定科目ID | [optional] 
**CorrespondingExpenseName** | Pointer to **NullableString** | 支出取引相手勘定科目名 | [optional] 
**CorrespondingIncomeId** | Pointer to **NullableInt64** | 収入取引相手勘定科目ID | [optional] 
**CorrespondingIncomeName** | Pointer to **NullableString** | 収入取引相手勘定科目名 | [optional] 
**GroupName** | Pointer to **NullableString** | 決算書表示名（小カテゴリー） | [optional] 
**Id** | **int64** | 勘定科目ID | 
**Items** | Pointer to [**[]AccountItemResponseAccountItemItemsInner**](AccountItemResponseAccountItemItemsInner.md) |  | [optional] 
**Name** | **string** | 勘定科目名 (30文字以内) | 
**Partners** | Pointer to [**[]AccountItemResponseAccountItemPartnersInner**](AccountItemResponseAccountItemPartnersInner.md) |  | [optional] 
**Searchable** | **int64** | 検索可能:2, 検索不可：3 | 
**Shortcut** | Pointer to **string** | ショートカット1 (20文字以内) | [optional] 
**ShortcutNum** | Pointer to **string** | ショートカット2(勘定科目コード) (20文字以内) | [optional] 
**TaxCode** | **int64** | 税区分コード | 
**WalletableId** | **NullableInt64** | 口座ID | 

## Methods

### NewAccountItemResponseAccountItem

`func NewAccountItemResponseAccountItem(accountCategory string, accountCategoryId int64, available bool, companyId int64, id int64, name string, searchable int64, taxCode int64, walletableId NullableInt64, ) *AccountItemResponseAccountItem`

NewAccountItemResponseAccountItem instantiates a new AccountItemResponseAccountItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountItemResponseAccountItemWithDefaults

`func NewAccountItemResponseAccountItemWithDefaults() *AccountItemResponseAccountItem`

NewAccountItemResponseAccountItemWithDefaults instantiates a new AccountItemResponseAccountItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountCategory

`func (o *AccountItemResponseAccountItem) GetAccountCategory() string`

GetAccountCategory returns the AccountCategory field if non-nil, zero value otherwise.

### GetAccountCategoryOk

`func (o *AccountItemResponseAccountItem) GetAccountCategoryOk() (*string, bool)`

GetAccountCategoryOk returns a tuple with the AccountCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCategory

`func (o *AccountItemResponseAccountItem) SetAccountCategory(v string)`

SetAccountCategory sets AccountCategory field to given value.


### GetAccountCategoryId

`func (o *AccountItemResponseAccountItem) GetAccountCategoryId() int64`

GetAccountCategoryId returns the AccountCategoryId field if non-nil, zero value otherwise.

### GetAccountCategoryIdOk

`func (o *AccountItemResponseAccountItem) GetAccountCategoryIdOk() (*int64, bool)`

GetAccountCategoryIdOk returns a tuple with the AccountCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCategoryId

`func (o *AccountItemResponseAccountItem) SetAccountCategoryId(v int64)`

SetAccountCategoryId sets AccountCategoryId field to given value.


### GetAccumulatedDepAccountItemId

`func (o *AccountItemResponseAccountItem) GetAccumulatedDepAccountItemId() int64`

GetAccumulatedDepAccountItemId returns the AccumulatedDepAccountItemId field if non-nil, zero value otherwise.

### GetAccumulatedDepAccountItemIdOk

`func (o *AccountItemResponseAccountItem) GetAccumulatedDepAccountItemIdOk() (*int64, bool)`

GetAccumulatedDepAccountItemIdOk returns a tuple with the AccumulatedDepAccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccumulatedDepAccountItemId

`func (o *AccountItemResponseAccountItem) SetAccumulatedDepAccountItemId(v int64)`

SetAccumulatedDepAccountItemId sets AccumulatedDepAccountItemId field to given value.

### HasAccumulatedDepAccountItemId

`func (o *AccountItemResponseAccountItem) HasAccumulatedDepAccountItemId() bool`

HasAccumulatedDepAccountItemId returns a boolean if a field has been set.

### SetAccumulatedDepAccountItemIdNil

`func (o *AccountItemResponseAccountItem) SetAccumulatedDepAccountItemIdNil(b bool)`

 SetAccumulatedDepAccountItemIdNil sets the value for AccumulatedDepAccountItemId to be an explicit nil

### UnsetAccumulatedDepAccountItemId
`func (o *AccountItemResponseAccountItem) UnsetAccumulatedDepAccountItemId()`

UnsetAccumulatedDepAccountItemId ensures that no value is present for AccumulatedDepAccountItemId, not even an explicit nil
### GetAccumulatedDepAccountItemName

`func (o *AccountItemResponseAccountItem) GetAccumulatedDepAccountItemName() string`

GetAccumulatedDepAccountItemName returns the AccumulatedDepAccountItemName field if non-nil, zero value otherwise.

### GetAccumulatedDepAccountItemNameOk

`func (o *AccountItemResponseAccountItem) GetAccumulatedDepAccountItemNameOk() (*string, bool)`

GetAccumulatedDepAccountItemNameOk returns a tuple with the AccumulatedDepAccountItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccumulatedDepAccountItemName

`func (o *AccountItemResponseAccountItem) SetAccumulatedDepAccountItemName(v string)`

SetAccumulatedDepAccountItemName sets AccumulatedDepAccountItemName field to given value.

### HasAccumulatedDepAccountItemName

`func (o *AccountItemResponseAccountItem) HasAccumulatedDepAccountItemName() bool`

HasAccumulatedDepAccountItemName returns a boolean if a field has been set.

### GetAvailable

`func (o *AccountItemResponseAccountItem) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *AccountItemResponseAccountItem) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *AccountItemResponseAccountItem) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetCompanyId

`func (o *AccountItemResponseAccountItem) GetCompanyId() int64`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *AccountItemResponseAccountItem) GetCompanyIdOk() (*int64, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *AccountItemResponseAccountItem) SetCompanyId(v int64)`

SetCompanyId sets CompanyId field to given value.


### GetCorrespondingExpenseId

`func (o *AccountItemResponseAccountItem) GetCorrespondingExpenseId() int64`

GetCorrespondingExpenseId returns the CorrespondingExpenseId field if non-nil, zero value otherwise.

### GetCorrespondingExpenseIdOk

`func (o *AccountItemResponseAccountItem) GetCorrespondingExpenseIdOk() (*int64, bool)`

GetCorrespondingExpenseIdOk returns a tuple with the CorrespondingExpenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrespondingExpenseId

`func (o *AccountItemResponseAccountItem) SetCorrespondingExpenseId(v int64)`

SetCorrespondingExpenseId sets CorrespondingExpenseId field to given value.

### HasCorrespondingExpenseId

`func (o *AccountItemResponseAccountItem) HasCorrespondingExpenseId() bool`

HasCorrespondingExpenseId returns a boolean if a field has been set.

### SetCorrespondingExpenseIdNil

`func (o *AccountItemResponseAccountItem) SetCorrespondingExpenseIdNil(b bool)`

 SetCorrespondingExpenseIdNil sets the value for CorrespondingExpenseId to be an explicit nil

### UnsetCorrespondingExpenseId
`func (o *AccountItemResponseAccountItem) UnsetCorrespondingExpenseId()`

UnsetCorrespondingExpenseId ensures that no value is present for CorrespondingExpenseId, not even an explicit nil
### GetCorrespondingExpenseName

`func (o *AccountItemResponseAccountItem) GetCorrespondingExpenseName() string`

GetCorrespondingExpenseName returns the CorrespondingExpenseName field if non-nil, zero value otherwise.

### GetCorrespondingExpenseNameOk

`func (o *AccountItemResponseAccountItem) GetCorrespondingExpenseNameOk() (*string, bool)`

GetCorrespondingExpenseNameOk returns a tuple with the CorrespondingExpenseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrespondingExpenseName

`func (o *AccountItemResponseAccountItem) SetCorrespondingExpenseName(v string)`

SetCorrespondingExpenseName sets CorrespondingExpenseName field to given value.

### HasCorrespondingExpenseName

`func (o *AccountItemResponseAccountItem) HasCorrespondingExpenseName() bool`

HasCorrespondingExpenseName returns a boolean if a field has been set.

### SetCorrespondingExpenseNameNil

`func (o *AccountItemResponseAccountItem) SetCorrespondingExpenseNameNil(b bool)`

 SetCorrespondingExpenseNameNil sets the value for CorrespondingExpenseName to be an explicit nil

### UnsetCorrespondingExpenseName
`func (o *AccountItemResponseAccountItem) UnsetCorrespondingExpenseName()`

UnsetCorrespondingExpenseName ensures that no value is present for CorrespondingExpenseName, not even an explicit nil
### GetCorrespondingIncomeId

`func (o *AccountItemResponseAccountItem) GetCorrespondingIncomeId() int64`

GetCorrespondingIncomeId returns the CorrespondingIncomeId field if non-nil, zero value otherwise.

### GetCorrespondingIncomeIdOk

`func (o *AccountItemResponseAccountItem) GetCorrespondingIncomeIdOk() (*int64, bool)`

GetCorrespondingIncomeIdOk returns a tuple with the CorrespondingIncomeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrespondingIncomeId

`func (o *AccountItemResponseAccountItem) SetCorrespondingIncomeId(v int64)`

SetCorrespondingIncomeId sets CorrespondingIncomeId field to given value.

### HasCorrespondingIncomeId

`func (o *AccountItemResponseAccountItem) HasCorrespondingIncomeId() bool`

HasCorrespondingIncomeId returns a boolean if a field has been set.

### SetCorrespondingIncomeIdNil

`func (o *AccountItemResponseAccountItem) SetCorrespondingIncomeIdNil(b bool)`

 SetCorrespondingIncomeIdNil sets the value for CorrespondingIncomeId to be an explicit nil

### UnsetCorrespondingIncomeId
`func (o *AccountItemResponseAccountItem) UnsetCorrespondingIncomeId()`

UnsetCorrespondingIncomeId ensures that no value is present for CorrespondingIncomeId, not even an explicit nil
### GetCorrespondingIncomeName

`func (o *AccountItemResponseAccountItem) GetCorrespondingIncomeName() string`

GetCorrespondingIncomeName returns the CorrespondingIncomeName field if non-nil, zero value otherwise.

### GetCorrespondingIncomeNameOk

`func (o *AccountItemResponseAccountItem) GetCorrespondingIncomeNameOk() (*string, bool)`

GetCorrespondingIncomeNameOk returns a tuple with the CorrespondingIncomeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrespondingIncomeName

`func (o *AccountItemResponseAccountItem) SetCorrespondingIncomeName(v string)`

SetCorrespondingIncomeName sets CorrespondingIncomeName field to given value.

### HasCorrespondingIncomeName

`func (o *AccountItemResponseAccountItem) HasCorrespondingIncomeName() bool`

HasCorrespondingIncomeName returns a boolean if a field has been set.

### SetCorrespondingIncomeNameNil

`func (o *AccountItemResponseAccountItem) SetCorrespondingIncomeNameNil(b bool)`

 SetCorrespondingIncomeNameNil sets the value for CorrespondingIncomeName to be an explicit nil

### UnsetCorrespondingIncomeName
`func (o *AccountItemResponseAccountItem) UnsetCorrespondingIncomeName()`

UnsetCorrespondingIncomeName ensures that no value is present for CorrespondingIncomeName, not even an explicit nil
### GetGroupName

`func (o *AccountItemResponseAccountItem) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *AccountItemResponseAccountItem) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *AccountItemResponseAccountItem) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *AccountItemResponseAccountItem) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *AccountItemResponseAccountItem) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *AccountItemResponseAccountItem) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetId

`func (o *AccountItemResponseAccountItem) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountItemResponseAccountItem) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountItemResponseAccountItem) SetId(v int64)`

SetId sets Id field to given value.


### GetItems

`func (o *AccountItemResponseAccountItem) GetItems() []AccountItemResponseAccountItemItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AccountItemResponseAccountItem) GetItemsOk() (*[]AccountItemResponseAccountItemItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AccountItemResponseAccountItem) SetItems(v []AccountItemResponseAccountItemItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *AccountItemResponseAccountItem) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetName

`func (o *AccountItemResponseAccountItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountItemResponseAccountItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountItemResponseAccountItem) SetName(v string)`

SetName sets Name field to given value.


### GetPartners

`func (o *AccountItemResponseAccountItem) GetPartners() []AccountItemResponseAccountItemPartnersInner`

GetPartners returns the Partners field if non-nil, zero value otherwise.

### GetPartnersOk

`func (o *AccountItemResponseAccountItem) GetPartnersOk() (*[]AccountItemResponseAccountItemPartnersInner, bool)`

GetPartnersOk returns a tuple with the Partners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartners

`func (o *AccountItemResponseAccountItem) SetPartners(v []AccountItemResponseAccountItemPartnersInner)`

SetPartners sets Partners field to given value.

### HasPartners

`func (o *AccountItemResponseAccountItem) HasPartners() bool`

HasPartners returns a boolean if a field has been set.

### GetSearchable

`func (o *AccountItemResponseAccountItem) GetSearchable() int64`

GetSearchable returns the Searchable field if non-nil, zero value otherwise.

### GetSearchableOk

`func (o *AccountItemResponseAccountItem) GetSearchableOk() (*int64, bool)`

GetSearchableOk returns a tuple with the Searchable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchable

`func (o *AccountItemResponseAccountItem) SetSearchable(v int64)`

SetSearchable sets Searchable field to given value.


### GetShortcut

`func (o *AccountItemResponseAccountItem) GetShortcut() string`

GetShortcut returns the Shortcut field if non-nil, zero value otherwise.

### GetShortcutOk

`func (o *AccountItemResponseAccountItem) GetShortcutOk() (*string, bool)`

GetShortcutOk returns a tuple with the Shortcut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut

`func (o *AccountItemResponseAccountItem) SetShortcut(v string)`

SetShortcut sets Shortcut field to given value.

### HasShortcut

`func (o *AccountItemResponseAccountItem) HasShortcut() bool`

HasShortcut returns a boolean if a field has been set.

### GetShortcutNum

`func (o *AccountItemResponseAccountItem) GetShortcutNum() string`

GetShortcutNum returns the ShortcutNum field if non-nil, zero value otherwise.

### GetShortcutNumOk

`func (o *AccountItemResponseAccountItem) GetShortcutNumOk() (*string, bool)`

GetShortcutNumOk returns a tuple with the ShortcutNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcutNum

`func (o *AccountItemResponseAccountItem) SetShortcutNum(v string)`

SetShortcutNum sets ShortcutNum field to given value.

### HasShortcutNum

`func (o *AccountItemResponseAccountItem) HasShortcutNum() bool`

HasShortcutNum returns a boolean if a field has been set.

### GetTaxCode

`func (o *AccountItemResponseAccountItem) GetTaxCode() int64`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *AccountItemResponseAccountItem) GetTaxCodeOk() (*int64, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *AccountItemResponseAccountItem) SetTaxCode(v int64)`

SetTaxCode sets TaxCode field to given value.


### GetWalletableId

`func (o *AccountItemResponseAccountItem) GetWalletableId() int64`

GetWalletableId returns the WalletableId field if non-nil, zero value otherwise.

### GetWalletableIdOk

`func (o *AccountItemResponseAccountItem) GetWalletableIdOk() (*int64, bool)`

GetWalletableIdOk returns a tuple with the WalletableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletableId

`func (o *AccountItemResponseAccountItem) SetWalletableId(v int64)`

SetWalletableId sets WalletableId field to given value.


### SetWalletableIdNil

`func (o *AccountItemResponseAccountItem) SetWalletableIdNil(b bool)`

 SetWalletableIdNil sets the value for WalletableId to be an explicit nil

### UnsetWalletableId
`func (o *AccountItemResponseAccountItem) UnsetWalletableId()`

UnsetWalletableId ensures that no value is present for WalletableId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


