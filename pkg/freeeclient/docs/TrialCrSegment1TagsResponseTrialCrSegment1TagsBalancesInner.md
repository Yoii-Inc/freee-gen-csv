# TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountCategoryName** | Pointer to **string** | 勘定科目カテゴリー名 | [optional] 
**AccountGroupName** | Pointer to **string** | 決算書表示名(account_item_display_type:group指定時に決算書表示名の時のみ含まれる) | [optional] 
**AccountItemId** | Pointer to **int32** | 勘定科目ID(勘定科目の時のみ含まれる) | [optional] 
**AccountItemName** | Pointer to **string** | 勘定科目名(勘定科目の時のみ含まれる) | [optional] 
**ClosingBalance** | Pointer to **int32** | 期末残高 | [optional] 
**HierarchyLevel** | Pointer to **int32** | 階層レベル | [optional] 
**ParentAccountCategoryName** | Pointer to **string** | 上位勘定科目カテゴリー名(勘定科目カテゴリーの時のみ、上層が存在する場合含まれる) | [optional] 
**Segment1Tags** | Pointer to [**[]TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInnerSegment1TagsInner**](TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInnerSegment1TagsInner.md) | セグメント1タグ | [optional] 
**TotalLine** | Pointer to **bool** | 合計行(勘定科目カテゴリーの時のみ含まれる) | [optional] 

## Methods

### NewTrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner

`func NewTrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner() *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner`

NewTrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner instantiates a new TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInnerWithDefaults

`func NewTrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInnerWithDefaults() *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner`

NewTrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInnerWithDefaults instantiates a new TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountCategoryName

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetAccountCategoryName() string`

GetAccountCategoryName returns the AccountCategoryName field if non-nil, zero value otherwise.

### GetAccountCategoryNameOk

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetAccountCategoryNameOk() (*string, bool)`

GetAccountCategoryNameOk returns a tuple with the AccountCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCategoryName

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) SetAccountCategoryName(v string)`

SetAccountCategoryName sets AccountCategoryName field to given value.

### HasAccountCategoryName

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) HasAccountCategoryName() bool`

HasAccountCategoryName returns a boolean if a field has been set.

### GetAccountGroupName

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetAccountGroupName() string`

GetAccountGroupName returns the AccountGroupName field if non-nil, zero value otherwise.

### GetAccountGroupNameOk

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetAccountGroupNameOk() (*string, bool)`

GetAccountGroupNameOk returns a tuple with the AccountGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroupName

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) SetAccountGroupName(v string)`

SetAccountGroupName sets AccountGroupName field to given value.

### HasAccountGroupName

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) HasAccountGroupName() bool`

HasAccountGroupName returns a boolean if a field has been set.

### GetAccountItemId

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetAccountItemId() int32`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetAccountItemIdOk() (*int32, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) SetAccountItemId(v int32)`

SetAccountItemId sets AccountItemId field to given value.

### HasAccountItemId

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) HasAccountItemId() bool`

HasAccountItemId returns a boolean if a field has been set.

### GetAccountItemName

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetAccountItemName() string`

GetAccountItemName returns the AccountItemName field if non-nil, zero value otherwise.

### GetAccountItemNameOk

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetAccountItemNameOk() (*string, bool)`

GetAccountItemNameOk returns a tuple with the AccountItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemName

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) SetAccountItemName(v string)`

SetAccountItemName sets AccountItemName field to given value.

### HasAccountItemName

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) HasAccountItemName() bool`

HasAccountItemName returns a boolean if a field has been set.

### GetClosingBalance

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetClosingBalance() int32`

GetClosingBalance returns the ClosingBalance field if non-nil, zero value otherwise.

### GetClosingBalanceOk

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetClosingBalanceOk() (*int32, bool)`

GetClosingBalanceOk returns a tuple with the ClosingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingBalance

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) SetClosingBalance(v int32)`

SetClosingBalance sets ClosingBalance field to given value.

### HasClosingBalance

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) HasClosingBalance() bool`

HasClosingBalance returns a boolean if a field has been set.

### GetHierarchyLevel

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetHierarchyLevel() int32`

GetHierarchyLevel returns the HierarchyLevel field if non-nil, zero value otherwise.

### GetHierarchyLevelOk

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetHierarchyLevelOk() (*int32, bool)`

GetHierarchyLevelOk returns a tuple with the HierarchyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchyLevel

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) SetHierarchyLevel(v int32)`

SetHierarchyLevel sets HierarchyLevel field to given value.

### HasHierarchyLevel

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) HasHierarchyLevel() bool`

HasHierarchyLevel returns a boolean if a field has been set.

### GetParentAccountCategoryName

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetParentAccountCategoryName() string`

GetParentAccountCategoryName returns the ParentAccountCategoryName field if non-nil, zero value otherwise.

### GetParentAccountCategoryNameOk

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetParentAccountCategoryNameOk() (*string, bool)`

GetParentAccountCategoryNameOk returns a tuple with the ParentAccountCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountCategoryName

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) SetParentAccountCategoryName(v string)`

SetParentAccountCategoryName sets ParentAccountCategoryName field to given value.

### HasParentAccountCategoryName

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) HasParentAccountCategoryName() bool`

HasParentAccountCategoryName returns a boolean if a field has been set.

### GetSegment1Tags

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetSegment1Tags() []TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInnerSegment1TagsInner`

GetSegment1Tags returns the Segment1Tags field if non-nil, zero value otherwise.

### GetSegment1TagsOk

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetSegment1TagsOk() (*[]TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInnerSegment1TagsInner, bool)`

GetSegment1TagsOk returns a tuple with the Segment1Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1Tags

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) SetSegment1Tags(v []TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInnerSegment1TagsInner)`

SetSegment1Tags sets Segment1Tags field to given value.

### HasSegment1Tags

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) HasSegment1Tags() bool`

HasSegment1Tags returns a boolean if a field has been set.

### GetTotalLine

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetTotalLine() bool`

GetTotalLine returns the TotalLine field if non-nil, zero value otherwise.

### GetTotalLineOk

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) GetTotalLineOk() (*bool, bool)`

GetTotalLineOk returns a tuple with the TotalLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLine

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) SetTotalLine(v bool)`

SetTotalLine sets TotalLine field to given value.

### HasTotalLine

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner) HasTotalLine() bool`

HasTotalLine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


