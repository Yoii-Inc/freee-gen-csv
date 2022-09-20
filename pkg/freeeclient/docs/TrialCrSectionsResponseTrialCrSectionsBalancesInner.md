# TrialCrSectionsResponseTrialCrSectionsBalancesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountCategoryName** | Pointer to **string** | 勘定科目カテゴリー名 | [optional] 
**AccountGroupName** | Pointer to **string** | 決算書表示名(account_item_display_type:group指定時に決算書表示名の時のみ含まれる) | [optional] 
**AccountItemId** | Pointer to **int64** | 勘定科目ID(勘定科目の時のみ含まれる) | [optional] 
**AccountItemName** | Pointer to **string** | 勘定科目名(勘定科目の時のみ含まれる) | [optional] 
**ClosingBalance** | Pointer to **int64** | 期末残高 | [optional] 
**HierarchyLevel** | Pointer to **int64** | 階層レベル | [optional] 
**ParentAccountCategoryName** | Pointer to **string** | 上位勘定科目カテゴリー名(勘定科目カテゴリーの時のみ、上層が存在する場合含まれる) | [optional] 
**Sections** | Pointer to [**[]TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner**](TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner.md) | 部門 | [optional] 
**TotalLine** | Pointer to **bool** | 合計行(勘定科目カテゴリーの時のみ含まれる) | [optional] 

## Methods

### NewTrialCrSectionsResponseTrialCrSectionsBalancesInner

`func NewTrialCrSectionsResponseTrialCrSectionsBalancesInner() *TrialCrSectionsResponseTrialCrSectionsBalancesInner`

NewTrialCrSectionsResponseTrialCrSectionsBalancesInner instantiates a new TrialCrSectionsResponseTrialCrSectionsBalancesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialCrSectionsResponseTrialCrSectionsBalancesInnerWithDefaults

`func NewTrialCrSectionsResponseTrialCrSectionsBalancesInnerWithDefaults() *TrialCrSectionsResponseTrialCrSectionsBalancesInner`

NewTrialCrSectionsResponseTrialCrSectionsBalancesInnerWithDefaults instantiates a new TrialCrSectionsResponseTrialCrSectionsBalancesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountCategoryName

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) GetAccountCategoryName() string`

GetAccountCategoryName returns the AccountCategoryName field if non-nil, zero value otherwise.

### GetAccountCategoryNameOk

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) GetAccountCategoryNameOk() (*string, bool)`

GetAccountCategoryNameOk returns a tuple with the AccountCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCategoryName

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) SetAccountCategoryName(v string)`

SetAccountCategoryName sets AccountCategoryName field to given value.

### HasAccountCategoryName

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) HasAccountCategoryName() bool`

HasAccountCategoryName returns a boolean if a field has been set.

### GetAccountGroupName

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) GetAccountGroupName() string`

GetAccountGroupName returns the AccountGroupName field if non-nil, zero value otherwise.

### GetAccountGroupNameOk

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) GetAccountGroupNameOk() (*string, bool)`

GetAccountGroupNameOk returns a tuple with the AccountGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroupName

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) SetAccountGroupName(v string)`

SetAccountGroupName sets AccountGroupName field to given value.

### HasAccountGroupName

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) HasAccountGroupName() bool`

HasAccountGroupName returns a boolean if a field has been set.

### GetAccountItemId

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) GetAccountItemId() int64`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) GetAccountItemIdOk() (*int64, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) SetAccountItemId(v int64)`

SetAccountItemId sets AccountItemId field to given value.

### HasAccountItemId

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) HasAccountItemId() bool`

HasAccountItemId returns a boolean if a field has been set.

### GetAccountItemName

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) GetAccountItemName() string`

GetAccountItemName returns the AccountItemName field if non-nil, zero value otherwise.

### GetAccountItemNameOk

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) GetAccountItemNameOk() (*string, bool)`

GetAccountItemNameOk returns a tuple with the AccountItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemName

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) SetAccountItemName(v string)`

SetAccountItemName sets AccountItemName field to given value.

### HasAccountItemName

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) HasAccountItemName() bool`

HasAccountItemName returns a boolean if a field has been set.

### GetClosingBalance

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) GetClosingBalance() int64`

GetClosingBalance returns the ClosingBalance field if non-nil, zero value otherwise.

### GetClosingBalanceOk

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) GetClosingBalanceOk() (*int64, bool)`

GetClosingBalanceOk returns a tuple with the ClosingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingBalance

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) SetClosingBalance(v int64)`

SetClosingBalance sets ClosingBalance field to given value.

### HasClosingBalance

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) HasClosingBalance() bool`

HasClosingBalance returns a boolean if a field has been set.

### GetHierarchyLevel

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) GetHierarchyLevel() int64`

GetHierarchyLevel returns the HierarchyLevel field if non-nil, zero value otherwise.

### GetHierarchyLevelOk

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) GetHierarchyLevelOk() (*int64, bool)`

GetHierarchyLevelOk returns a tuple with the HierarchyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchyLevel

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) SetHierarchyLevel(v int64)`

SetHierarchyLevel sets HierarchyLevel field to given value.

### HasHierarchyLevel

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) HasHierarchyLevel() bool`

HasHierarchyLevel returns a boolean if a field has been set.

### GetParentAccountCategoryName

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) GetParentAccountCategoryName() string`

GetParentAccountCategoryName returns the ParentAccountCategoryName field if non-nil, zero value otherwise.

### GetParentAccountCategoryNameOk

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) GetParentAccountCategoryNameOk() (*string, bool)`

GetParentAccountCategoryNameOk returns a tuple with the ParentAccountCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountCategoryName

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) SetParentAccountCategoryName(v string)`

SetParentAccountCategoryName sets ParentAccountCategoryName field to given value.

### HasParentAccountCategoryName

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) HasParentAccountCategoryName() bool`

HasParentAccountCategoryName returns a boolean if a field has been set.

### GetSections

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) GetSections() []TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) GetSectionsOk() (*[]TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) SetSections(v []TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner)`

SetSections sets Sections field to given value.

### HasSections

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) HasSections() bool`

HasSections returns a boolean if a field has been set.

### GetTotalLine

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) GetTotalLine() bool`

GetTotalLine returns the TotalLine field if non-nil, zero value otherwise.

### GetTotalLineOk

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) GetTotalLineOk() (*bool, bool)`

GetTotalLineOk returns a tuple with the TotalLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLine

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) SetTotalLine(v bool)`

SetTotalLine sets TotalLine field to given value.

### HasTotalLine

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInner) HasTotalLine() bool`

HasTotalLine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


