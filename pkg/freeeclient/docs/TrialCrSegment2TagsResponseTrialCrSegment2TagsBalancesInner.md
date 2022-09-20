# TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner

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
**Segment2Tags** | Pointer to [**[]TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner**](TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner.md) | セグメント2タグ | [optional] 
**TotalLine** | Pointer to **bool** | 合計行(勘定科目カテゴリーの時のみ含まれる) | [optional] 

## Methods

### NewTrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner

`func NewTrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner() *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner`

NewTrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner instantiates a new TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerWithDefaults

`func NewTrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerWithDefaults() *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner`

NewTrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerWithDefaults instantiates a new TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountCategoryName

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) GetAccountCategoryName() string`

GetAccountCategoryName returns the AccountCategoryName field if non-nil, zero value otherwise.

### GetAccountCategoryNameOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) GetAccountCategoryNameOk() (*string, bool)`

GetAccountCategoryNameOk returns a tuple with the AccountCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCategoryName

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) SetAccountCategoryName(v string)`

SetAccountCategoryName sets AccountCategoryName field to given value.

### HasAccountCategoryName

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) HasAccountCategoryName() bool`

HasAccountCategoryName returns a boolean if a field has been set.

### GetAccountGroupName

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) GetAccountGroupName() string`

GetAccountGroupName returns the AccountGroupName field if non-nil, zero value otherwise.

### GetAccountGroupNameOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) GetAccountGroupNameOk() (*string, bool)`

GetAccountGroupNameOk returns a tuple with the AccountGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroupName

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) SetAccountGroupName(v string)`

SetAccountGroupName sets AccountGroupName field to given value.

### HasAccountGroupName

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) HasAccountGroupName() bool`

HasAccountGroupName returns a boolean if a field has been set.

### GetAccountItemId

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) GetAccountItemId() int64`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) GetAccountItemIdOk() (*int64, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) SetAccountItemId(v int64)`

SetAccountItemId sets AccountItemId field to given value.

### HasAccountItemId

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) HasAccountItemId() bool`

HasAccountItemId returns a boolean if a field has been set.

### GetAccountItemName

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) GetAccountItemName() string`

GetAccountItemName returns the AccountItemName field if non-nil, zero value otherwise.

### GetAccountItemNameOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) GetAccountItemNameOk() (*string, bool)`

GetAccountItemNameOk returns a tuple with the AccountItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemName

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) SetAccountItemName(v string)`

SetAccountItemName sets AccountItemName field to given value.

### HasAccountItemName

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) HasAccountItemName() bool`

HasAccountItemName returns a boolean if a field has been set.

### GetClosingBalance

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) GetClosingBalance() int64`

GetClosingBalance returns the ClosingBalance field if non-nil, zero value otherwise.

### GetClosingBalanceOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) GetClosingBalanceOk() (*int64, bool)`

GetClosingBalanceOk returns a tuple with the ClosingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingBalance

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) SetClosingBalance(v int64)`

SetClosingBalance sets ClosingBalance field to given value.

### HasClosingBalance

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) HasClosingBalance() bool`

HasClosingBalance returns a boolean if a field has been set.

### GetHierarchyLevel

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) GetHierarchyLevel() int64`

GetHierarchyLevel returns the HierarchyLevel field if non-nil, zero value otherwise.

### GetHierarchyLevelOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) GetHierarchyLevelOk() (*int64, bool)`

GetHierarchyLevelOk returns a tuple with the HierarchyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchyLevel

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) SetHierarchyLevel(v int64)`

SetHierarchyLevel sets HierarchyLevel field to given value.

### HasHierarchyLevel

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) HasHierarchyLevel() bool`

HasHierarchyLevel returns a boolean if a field has been set.

### GetParentAccountCategoryName

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) GetParentAccountCategoryName() string`

GetParentAccountCategoryName returns the ParentAccountCategoryName field if non-nil, zero value otherwise.

### GetParentAccountCategoryNameOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) GetParentAccountCategoryNameOk() (*string, bool)`

GetParentAccountCategoryNameOk returns a tuple with the ParentAccountCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountCategoryName

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) SetParentAccountCategoryName(v string)`

SetParentAccountCategoryName sets ParentAccountCategoryName field to given value.

### HasParentAccountCategoryName

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) HasParentAccountCategoryName() bool`

HasParentAccountCategoryName returns a boolean if a field has been set.

### GetSegment2Tags

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) GetSegment2Tags() []TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner`

GetSegment2Tags returns the Segment2Tags field if non-nil, zero value otherwise.

### GetSegment2TagsOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) GetSegment2TagsOk() (*[]TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner, bool)`

GetSegment2TagsOk returns a tuple with the Segment2Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2Tags

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) SetSegment2Tags(v []TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner)`

SetSegment2Tags sets Segment2Tags field to given value.

### HasSegment2Tags

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) HasSegment2Tags() bool`

HasSegment2Tags returns a boolean if a field has been set.

### GetTotalLine

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) GetTotalLine() bool`

GetTotalLine returns the TotalLine field if non-nil, zero value otherwise.

### GetTotalLineOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) GetTotalLineOk() (*bool, bool)`

GetTotalLineOk returns a tuple with the TotalLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLine

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) SetTotalLine(v bool)`

SetTotalLine sets TotalLine field to given value.

### HasTotalLine

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner) HasTotalLine() bool`

HasTotalLine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


