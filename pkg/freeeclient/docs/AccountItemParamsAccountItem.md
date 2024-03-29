# AccountItemParamsAccountItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountCategoryId** | **int64** | 勘定科目カテゴリーID Selectablesフォーム用選択項目情報エンドポイント(account_groups.account_category_id)で取得可能です | 
**AccumulatedDepAccountItemId** | Pointer to **int64** | 減価償却累計額勘定科目ID（法人のみ利用可能） | [optional] 
**CorrespondingExpenseId** | **int64** | 支出取引相手勘定科目ID | 
**CorrespondingIncomeId** | **int64** | 収入取引相手勘定科目ID | 
**GroupName** | **string** | 決算書表示名（小カテゴリー） Selectablesフォーム用選択項目情報エンドポイント(account_groups.name)で取得可能です | 
**Items** | Pointer to [**[]AccountItemParamsAccountItemItemsInner**](AccountItemParamsAccountItemItemsInner.md) | 品目 | [optional] 
**Name** | **string** | 勘定科目名 (30文字以内) | 
**Partners** | Pointer to [**[]AccountItemParamsAccountItemItemsInner**](AccountItemParamsAccountItemItemsInner.md) | 取引先 | [optional] 
**Searchable** | Pointer to **int64** | 検索可能:2, 検索不可：3(登録時未指定の場合は2で登録されます。更新時未指定の場合はsearchableは変更されません。) | [optional] 
**Shortcut** | Pointer to **string** | ショートカット1 (20文字以内) | [optional] 
**ShortcutNum** | Pointer to **string** | ショートカット2(勘定科目コード)(20文字以内) | [optional] 
**TaxCode** | **int64** | 税区分コード | 

## Methods

### NewAccountItemParamsAccountItem

`func NewAccountItemParamsAccountItem(accountCategoryId int64, correspondingExpenseId int64, correspondingIncomeId int64, groupName string, name string, taxCode int64, ) *AccountItemParamsAccountItem`

NewAccountItemParamsAccountItem instantiates a new AccountItemParamsAccountItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountItemParamsAccountItemWithDefaults

`func NewAccountItemParamsAccountItemWithDefaults() *AccountItemParamsAccountItem`

NewAccountItemParamsAccountItemWithDefaults instantiates a new AccountItemParamsAccountItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountCategoryId

`func (o *AccountItemParamsAccountItem) GetAccountCategoryId() int64`

GetAccountCategoryId returns the AccountCategoryId field if non-nil, zero value otherwise.

### GetAccountCategoryIdOk

`func (o *AccountItemParamsAccountItem) GetAccountCategoryIdOk() (*int64, bool)`

GetAccountCategoryIdOk returns a tuple with the AccountCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCategoryId

`func (o *AccountItemParamsAccountItem) SetAccountCategoryId(v int64)`

SetAccountCategoryId sets AccountCategoryId field to given value.


### GetAccumulatedDepAccountItemId

`func (o *AccountItemParamsAccountItem) GetAccumulatedDepAccountItemId() int64`

GetAccumulatedDepAccountItemId returns the AccumulatedDepAccountItemId field if non-nil, zero value otherwise.

### GetAccumulatedDepAccountItemIdOk

`func (o *AccountItemParamsAccountItem) GetAccumulatedDepAccountItemIdOk() (*int64, bool)`

GetAccumulatedDepAccountItemIdOk returns a tuple with the AccumulatedDepAccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccumulatedDepAccountItemId

`func (o *AccountItemParamsAccountItem) SetAccumulatedDepAccountItemId(v int64)`

SetAccumulatedDepAccountItemId sets AccumulatedDepAccountItemId field to given value.

### HasAccumulatedDepAccountItemId

`func (o *AccountItemParamsAccountItem) HasAccumulatedDepAccountItemId() bool`

HasAccumulatedDepAccountItemId returns a boolean if a field has been set.

### GetCorrespondingExpenseId

`func (o *AccountItemParamsAccountItem) GetCorrespondingExpenseId() int64`

GetCorrespondingExpenseId returns the CorrespondingExpenseId field if non-nil, zero value otherwise.

### GetCorrespondingExpenseIdOk

`func (o *AccountItemParamsAccountItem) GetCorrespondingExpenseIdOk() (*int64, bool)`

GetCorrespondingExpenseIdOk returns a tuple with the CorrespondingExpenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrespondingExpenseId

`func (o *AccountItemParamsAccountItem) SetCorrespondingExpenseId(v int64)`

SetCorrespondingExpenseId sets CorrespondingExpenseId field to given value.


### GetCorrespondingIncomeId

`func (o *AccountItemParamsAccountItem) GetCorrespondingIncomeId() int64`

GetCorrespondingIncomeId returns the CorrespondingIncomeId field if non-nil, zero value otherwise.

### GetCorrespondingIncomeIdOk

`func (o *AccountItemParamsAccountItem) GetCorrespondingIncomeIdOk() (*int64, bool)`

GetCorrespondingIncomeIdOk returns a tuple with the CorrespondingIncomeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrespondingIncomeId

`func (o *AccountItemParamsAccountItem) SetCorrespondingIncomeId(v int64)`

SetCorrespondingIncomeId sets CorrespondingIncomeId field to given value.


### GetGroupName

`func (o *AccountItemParamsAccountItem) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *AccountItemParamsAccountItem) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *AccountItemParamsAccountItem) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetItems

`func (o *AccountItemParamsAccountItem) GetItems() []AccountItemParamsAccountItemItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AccountItemParamsAccountItem) GetItemsOk() (*[]AccountItemParamsAccountItemItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AccountItemParamsAccountItem) SetItems(v []AccountItemParamsAccountItemItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *AccountItemParamsAccountItem) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetName

`func (o *AccountItemParamsAccountItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountItemParamsAccountItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountItemParamsAccountItem) SetName(v string)`

SetName sets Name field to given value.


### GetPartners

`func (o *AccountItemParamsAccountItem) GetPartners() []AccountItemParamsAccountItemItemsInner`

GetPartners returns the Partners field if non-nil, zero value otherwise.

### GetPartnersOk

`func (o *AccountItemParamsAccountItem) GetPartnersOk() (*[]AccountItemParamsAccountItemItemsInner, bool)`

GetPartnersOk returns a tuple with the Partners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartners

`func (o *AccountItemParamsAccountItem) SetPartners(v []AccountItemParamsAccountItemItemsInner)`

SetPartners sets Partners field to given value.

### HasPartners

`func (o *AccountItemParamsAccountItem) HasPartners() bool`

HasPartners returns a boolean if a field has been set.

### GetSearchable

`func (o *AccountItemParamsAccountItem) GetSearchable() int64`

GetSearchable returns the Searchable field if non-nil, zero value otherwise.

### GetSearchableOk

`func (o *AccountItemParamsAccountItem) GetSearchableOk() (*int64, bool)`

GetSearchableOk returns a tuple with the Searchable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchable

`func (o *AccountItemParamsAccountItem) SetSearchable(v int64)`

SetSearchable sets Searchable field to given value.

### HasSearchable

`func (o *AccountItemParamsAccountItem) HasSearchable() bool`

HasSearchable returns a boolean if a field has been set.

### GetShortcut

`func (o *AccountItemParamsAccountItem) GetShortcut() string`

GetShortcut returns the Shortcut field if non-nil, zero value otherwise.

### GetShortcutOk

`func (o *AccountItemParamsAccountItem) GetShortcutOk() (*string, bool)`

GetShortcutOk returns a tuple with the Shortcut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut

`func (o *AccountItemParamsAccountItem) SetShortcut(v string)`

SetShortcut sets Shortcut field to given value.

### HasShortcut

`func (o *AccountItemParamsAccountItem) HasShortcut() bool`

HasShortcut returns a boolean if a field has been set.

### GetShortcutNum

`func (o *AccountItemParamsAccountItem) GetShortcutNum() string`

GetShortcutNum returns the ShortcutNum field if non-nil, zero value otherwise.

### GetShortcutNumOk

`func (o *AccountItemParamsAccountItem) GetShortcutNumOk() (*string, bool)`

GetShortcutNumOk returns a tuple with the ShortcutNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcutNum

`func (o *AccountItemParamsAccountItem) SetShortcutNum(v string)`

SetShortcutNum sets ShortcutNum field to given value.

### HasShortcutNum

`func (o *AccountItemParamsAccountItem) HasShortcutNum() bool`

HasShortcutNum returns a boolean if a field has been set.

### GetTaxCode

`func (o *AccountItemParamsAccountItem) GetTaxCode() int64`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *AccountItemParamsAccountItem) GetTaxCodeOk() (*int64, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *AccountItemParamsAccountItem) SetTaxCode(v int64)`

SetTaxCode sets TaxCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


