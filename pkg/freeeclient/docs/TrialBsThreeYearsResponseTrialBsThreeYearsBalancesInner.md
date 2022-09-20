# TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountCategoryName** | Pointer to **string** | 勘定科目カテゴリー名 | [optional] 
**AccountGroupName** | Pointer to **string** | 決算書表示名(account_item_display_type:group指定時に決算書表示名の時のみ含まれる) | [optional] 
**AccountItemId** | Pointer to **int32** | 勘定科目ID(勘定科目の時のみ含まれる) | [optional] 
**AccountItemName** | Pointer to **string** | 勘定科目名(勘定科目の時のみ含まれる) | [optional] 
**ClosingBalance** | Pointer to **int32** | 期末残高 | [optional] 
**HierarchyLevel** | Pointer to **int32** | 階層レベル | [optional] 
**Items** | Pointer to [**[]TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerItemsInner**](TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerItemsInner.md) | breakdown_display_type:item, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**LastYearClosingBalance** | Pointer to **int32** | 前年度期末残高 | [optional] 
**ParentAccountCategoryName** | Pointer to **string** | 上位勘定科目カテゴリー名(勘定科目カテゴリーの時のみ、上層が存在する場合含まれる) | [optional] 
**Partners** | Pointer to [**[]TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerPartnersInner**](TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerPartnersInner.md) | breakdown_display_type:partner, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**Sections** | Pointer to [**[]TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerSectionsInner**](TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerSectionsInner.md) | breakdown_display_type:section, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**Segment1Tags** | Pointer to [**[]TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerSegment1TagsInner**](TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerSegment1TagsInner.md) | breakdown_display_type:segment_1_tag, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**Segment2Tags** | Pointer to [**[]TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerSegment2TagsInner**](TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerSegment2TagsInner.md) | breakdown_display_type:segment_2_tag, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**Segment3Tags** | Pointer to [**[]TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerSegment3TagsInner**](TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerSegment3TagsInner.md) | breakdown_display_type:segment_3_tag, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**TotalLine** | Pointer to **bool** | 合計行(勘定科目カテゴリーの時のみ含まれる) | [optional] 
**TwoYearsBeforeClosingBalance** | Pointer to **int32** | 前々年度期末残高 | [optional] 
**YearOnYear** | Pointer to **float32** | 前年比 | [optional] 

## Methods

### NewTrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner

`func NewTrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner() *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner`

NewTrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner instantiates a new TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerWithDefaults

`func NewTrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerWithDefaults() *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner`

NewTrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerWithDefaults instantiates a new TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountCategoryName

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetAccountCategoryName() string`

GetAccountCategoryName returns the AccountCategoryName field if non-nil, zero value otherwise.

### GetAccountCategoryNameOk

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetAccountCategoryNameOk() (*string, bool)`

GetAccountCategoryNameOk returns a tuple with the AccountCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCategoryName

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) SetAccountCategoryName(v string)`

SetAccountCategoryName sets AccountCategoryName field to given value.

### HasAccountCategoryName

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) HasAccountCategoryName() bool`

HasAccountCategoryName returns a boolean if a field has been set.

### GetAccountGroupName

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetAccountGroupName() string`

GetAccountGroupName returns the AccountGroupName field if non-nil, zero value otherwise.

### GetAccountGroupNameOk

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetAccountGroupNameOk() (*string, bool)`

GetAccountGroupNameOk returns a tuple with the AccountGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroupName

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) SetAccountGroupName(v string)`

SetAccountGroupName sets AccountGroupName field to given value.

### HasAccountGroupName

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) HasAccountGroupName() bool`

HasAccountGroupName returns a boolean if a field has been set.

### GetAccountItemId

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetAccountItemId() int32`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetAccountItemIdOk() (*int32, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) SetAccountItemId(v int32)`

SetAccountItemId sets AccountItemId field to given value.

### HasAccountItemId

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) HasAccountItemId() bool`

HasAccountItemId returns a boolean if a field has been set.

### GetAccountItemName

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetAccountItemName() string`

GetAccountItemName returns the AccountItemName field if non-nil, zero value otherwise.

### GetAccountItemNameOk

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetAccountItemNameOk() (*string, bool)`

GetAccountItemNameOk returns a tuple with the AccountItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemName

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) SetAccountItemName(v string)`

SetAccountItemName sets AccountItemName field to given value.

### HasAccountItemName

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) HasAccountItemName() bool`

HasAccountItemName returns a boolean if a field has been set.

### GetClosingBalance

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetClosingBalance() int32`

GetClosingBalance returns the ClosingBalance field if non-nil, zero value otherwise.

### GetClosingBalanceOk

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetClosingBalanceOk() (*int32, bool)`

GetClosingBalanceOk returns a tuple with the ClosingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingBalance

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) SetClosingBalance(v int32)`

SetClosingBalance sets ClosingBalance field to given value.

### HasClosingBalance

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) HasClosingBalance() bool`

HasClosingBalance returns a boolean if a field has been set.

### GetHierarchyLevel

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetHierarchyLevel() int32`

GetHierarchyLevel returns the HierarchyLevel field if non-nil, zero value otherwise.

### GetHierarchyLevelOk

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetHierarchyLevelOk() (*int32, bool)`

GetHierarchyLevelOk returns a tuple with the HierarchyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchyLevel

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) SetHierarchyLevel(v int32)`

SetHierarchyLevel sets HierarchyLevel field to given value.

### HasHierarchyLevel

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) HasHierarchyLevel() bool`

HasHierarchyLevel returns a boolean if a field has been set.

### GetItems

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetItems() []TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetItemsOk() (*[]TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) SetItems(v []TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLastYearClosingBalance

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetLastYearClosingBalance() int32`

GetLastYearClosingBalance returns the LastYearClosingBalance field if non-nil, zero value otherwise.

### GetLastYearClosingBalanceOk

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetLastYearClosingBalanceOk() (*int32, bool)`

GetLastYearClosingBalanceOk returns a tuple with the LastYearClosingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastYearClosingBalance

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) SetLastYearClosingBalance(v int32)`

SetLastYearClosingBalance sets LastYearClosingBalance field to given value.

### HasLastYearClosingBalance

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) HasLastYearClosingBalance() bool`

HasLastYearClosingBalance returns a boolean if a field has been set.

### GetParentAccountCategoryName

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetParentAccountCategoryName() string`

GetParentAccountCategoryName returns the ParentAccountCategoryName field if non-nil, zero value otherwise.

### GetParentAccountCategoryNameOk

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetParentAccountCategoryNameOk() (*string, bool)`

GetParentAccountCategoryNameOk returns a tuple with the ParentAccountCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountCategoryName

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) SetParentAccountCategoryName(v string)`

SetParentAccountCategoryName sets ParentAccountCategoryName field to given value.

### HasParentAccountCategoryName

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) HasParentAccountCategoryName() bool`

HasParentAccountCategoryName returns a boolean if a field has been set.

### GetPartners

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetPartners() []TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerPartnersInner`

GetPartners returns the Partners field if non-nil, zero value otherwise.

### GetPartnersOk

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetPartnersOk() (*[]TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerPartnersInner, bool)`

GetPartnersOk returns a tuple with the Partners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartners

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) SetPartners(v []TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerPartnersInner)`

SetPartners sets Partners field to given value.

### HasPartners

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) HasPartners() bool`

HasPartners returns a boolean if a field has been set.

### GetSections

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetSections() []TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerSectionsInner`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetSectionsOk() (*[]TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerSectionsInner, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) SetSections(v []TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerSectionsInner)`

SetSections sets Sections field to given value.

### HasSections

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) HasSections() bool`

HasSections returns a boolean if a field has been set.

### GetSegment1Tags

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetSegment1Tags() []TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerSegment1TagsInner`

GetSegment1Tags returns the Segment1Tags field if non-nil, zero value otherwise.

### GetSegment1TagsOk

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetSegment1TagsOk() (*[]TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerSegment1TagsInner, bool)`

GetSegment1TagsOk returns a tuple with the Segment1Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1Tags

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) SetSegment1Tags(v []TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerSegment1TagsInner)`

SetSegment1Tags sets Segment1Tags field to given value.

### HasSegment1Tags

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) HasSegment1Tags() bool`

HasSegment1Tags returns a boolean if a field has been set.

### GetSegment2Tags

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetSegment2Tags() []TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerSegment2TagsInner`

GetSegment2Tags returns the Segment2Tags field if non-nil, zero value otherwise.

### GetSegment2TagsOk

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetSegment2TagsOk() (*[]TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerSegment2TagsInner, bool)`

GetSegment2TagsOk returns a tuple with the Segment2Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2Tags

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) SetSegment2Tags(v []TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerSegment2TagsInner)`

SetSegment2Tags sets Segment2Tags field to given value.

### HasSegment2Tags

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) HasSegment2Tags() bool`

HasSegment2Tags returns a boolean if a field has been set.

### GetSegment3Tags

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetSegment3Tags() []TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerSegment3TagsInner`

GetSegment3Tags returns the Segment3Tags field if non-nil, zero value otherwise.

### GetSegment3TagsOk

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetSegment3TagsOk() (*[]TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerSegment3TagsInner, bool)`

GetSegment3TagsOk returns a tuple with the Segment3Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3Tags

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) SetSegment3Tags(v []TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInnerSegment3TagsInner)`

SetSegment3Tags sets Segment3Tags field to given value.

### HasSegment3Tags

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) HasSegment3Tags() bool`

HasSegment3Tags returns a boolean if a field has been set.

### GetTotalLine

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetTotalLine() bool`

GetTotalLine returns the TotalLine field if non-nil, zero value otherwise.

### GetTotalLineOk

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetTotalLineOk() (*bool, bool)`

GetTotalLineOk returns a tuple with the TotalLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLine

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) SetTotalLine(v bool)`

SetTotalLine sets TotalLine field to given value.

### HasTotalLine

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) HasTotalLine() bool`

HasTotalLine returns a boolean if a field has been set.

### GetTwoYearsBeforeClosingBalance

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetTwoYearsBeforeClosingBalance() int32`

GetTwoYearsBeforeClosingBalance returns the TwoYearsBeforeClosingBalance field if non-nil, zero value otherwise.

### GetTwoYearsBeforeClosingBalanceOk

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetTwoYearsBeforeClosingBalanceOk() (*int32, bool)`

GetTwoYearsBeforeClosingBalanceOk returns a tuple with the TwoYearsBeforeClosingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoYearsBeforeClosingBalance

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) SetTwoYearsBeforeClosingBalance(v int32)`

SetTwoYearsBeforeClosingBalance sets TwoYearsBeforeClosingBalance field to given value.

### HasTwoYearsBeforeClosingBalance

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) HasTwoYearsBeforeClosingBalance() bool`

HasTwoYearsBeforeClosingBalance returns a boolean if a field has been set.

### GetYearOnYear

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetYearOnYear() float32`

GetYearOnYear returns the YearOnYear field if non-nil, zero value otherwise.

### GetYearOnYearOk

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) GetYearOnYearOk() (*float32, bool)`

GetYearOnYearOk returns a tuple with the YearOnYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearOnYear

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) SetYearOnYear(v float32)`

SetYearOnYear sets YearOnYear field to given value.

### HasYearOnYear

`func (o *TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner) HasYearOnYear() bool`

HasYearOnYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


