# TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner

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
**Segment3Tags** | Pointer to [**[]TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInnerSegment3TagsInner**](TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInnerSegment3TagsInner.md) | セグメント3タグ | [optional] 
**TotalLine** | Pointer to **bool** | 合計行(勘定科目カテゴリーの時のみ含まれる) | [optional] 

## Methods

### NewTrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner

`func NewTrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner() *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner`

NewTrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner instantiates a new TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInnerWithDefaults

`func NewTrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInnerWithDefaults() *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner`

NewTrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInnerWithDefaults instantiates a new TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountCategoryName

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) GetAccountCategoryName() string`

GetAccountCategoryName returns the AccountCategoryName field if non-nil, zero value otherwise.

### GetAccountCategoryNameOk

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) GetAccountCategoryNameOk() (*string, bool)`

GetAccountCategoryNameOk returns a tuple with the AccountCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCategoryName

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) SetAccountCategoryName(v string)`

SetAccountCategoryName sets AccountCategoryName field to given value.

### HasAccountCategoryName

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) HasAccountCategoryName() bool`

HasAccountCategoryName returns a boolean if a field has been set.

### GetAccountGroupName

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) GetAccountGroupName() string`

GetAccountGroupName returns the AccountGroupName field if non-nil, zero value otherwise.

### GetAccountGroupNameOk

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) GetAccountGroupNameOk() (*string, bool)`

GetAccountGroupNameOk returns a tuple with the AccountGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroupName

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) SetAccountGroupName(v string)`

SetAccountGroupName sets AccountGroupName field to given value.

### HasAccountGroupName

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) HasAccountGroupName() bool`

HasAccountGroupName returns a boolean if a field has been set.

### GetAccountItemId

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) GetAccountItemId() int32`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) GetAccountItemIdOk() (*int32, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) SetAccountItemId(v int32)`

SetAccountItemId sets AccountItemId field to given value.

### HasAccountItemId

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) HasAccountItemId() bool`

HasAccountItemId returns a boolean if a field has been set.

### GetAccountItemName

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) GetAccountItemName() string`

GetAccountItemName returns the AccountItemName field if non-nil, zero value otherwise.

### GetAccountItemNameOk

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) GetAccountItemNameOk() (*string, bool)`

GetAccountItemNameOk returns a tuple with the AccountItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemName

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) SetAccountItemName(v string)`

SetAccountItemName sets AccountItemName field to given value.

### HasAccountItemName

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) HasAccountItemName() bool`

HasAccountItemName returns a boolean if a field has been set.

### GetClosingBalance

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) GetClosingBalance() int32`

GetClosingBalance returns the ClosingBalance field if non-nil, zero value otherwise.

### GetClosingBalanceOk

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) GetClosingBalanceOk() (*int32, bool)`

GetClosingBalanceOk returns a tuple with the ClosingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingBalance

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) SetClosingBalance(v int32)`

SetClosingBalance sets ClosingBalance field to given value.

### HasClosingBalance

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) HasClosingBalance() bool`

HasClosingBalance returns a boolean if a field has been set.

### GetHierarchyLevel

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) GetHierarchyLevel() int32`

GetHierarchyLevel returns the HierarchyLevel field if non-nil, zero value otherwise.

### GetHierarchyLevelOk

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) GetHierarchyLevelOk() (*int32, bool)`

GetHierarchyLevelOk returns a tuple with the HierarchyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchyLevel

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) SetHierarchyLevel(v int32)`

SetHierarchyLevel sets HierarchyLevel field to given value.

### HasHierarchyLevel

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) HasHierarchyLevel() bool`

HasHierarchyLevel returns a boolean if a field has been set.

### GetParentAccountCategoryName

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) GetParentAccountCategoryName() string`

GetParentAccountCategoryName returns the ParentAccountCategoryName field if non-nil, zero value otherwise.

### GetParentAccountCategoryNameOk

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) GetParentAccountCategoryNameOk() (*string, bool)`

GetParentAccountCategoryNameOk returns a tuple with the ParentAccountCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountCategoryName

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) SetParentAccountCategoryName(v string)`

SetParentAccountCategoryName sets ParentAccountCategoryName field to given value.

### HasParentAccountCategoryName

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) HasParentAccountCategoryName() bool`

HasParentAccountCategoryName returns a boolean if a field has been set.

### GetSegment3Tags

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) GetSegment3Tags() []TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInnerSegment3TagsInner`

GetSegment3Tags returns the Segment3Tags field if non-nil, zero value otherwise.

### GetSegment3TagsOk

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) GetSegment3TagsOk() (*[]TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInnerSegment3TagsInner, bool)`

GetSegment3TagsOk returns a tuple with the Segment3Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3Tags

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) SetSegment3Tags(v []TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInnerSegment3TagsInner)`

SetSegment3Tags sets Segment3Tags field to given value.

### HasSegment3Tags

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) HasSegment3Tags() bool`

HasSegment3Tags returns a boolean if a field has been set.

### GetTotalLine

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) GetTotalLine() bool`

GetTotalLine returns the TotalLine field if non-nil, zero value otherwise.

### GetTotalLineOk

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) GetTotalLineOk() (*bool, bool)`

GetTotalLineOk returns a tuple with the TotalLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLine

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) SetTotalLine(v bool)`

SetTotalLine sets TotalLine field to given value.

### HasTotalLine

`func (o *TrialCrSegment3TagsResponseTrialCrSegment3TagsBalancesInner) HasTotalLine() bool`

HasTotalLine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


