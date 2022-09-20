# TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountCategoryName** | Pointer to **string** | 勘定科目カテゴリー名 | [optional] 
**AccountGroupName** | Pointer to **string** | 決算書表示名(account_item_display_type:group指定時に決算書表示名の時のみ含まれる) | [optional] 
**AccountItemId** | Pointer to **int32** | 勘定科目ID(勘定科目の時のみ含まれる) | [optional] 
**AccountItemName** | Pointer to **string** | 勘定科目名(勘定科目の時のみ含まれる) | [optional] 
**ClosingBalance** | Pointer to **int32** | 期末残高 | [optional] 
**HierarchyLevel** | Pointer to **int32** | 階層レベル | [optional] 
**Items** | Pointer to [**[]TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerItemsInner**](TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerItemsInner.md) | breakdown_display_type:item, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**LastYearClosingBalance** | Pointer to **int32** | 前年度期末残高 | [optional] 
**ParentAccountCategoryName** | Pointer to **string** | 上位勘定科目カテゴリー名(勘定科目カテゴリーの時のみ、上層が存在する場合含まれる) | [optional] 
**Partners** | Pointer to [**[]TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerPartnersInner**](TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerPartnersInner.md) | breakdown_display_type:partner, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**Sections** | Pointer to [**[]TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSectionsInner**](TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSectionsInner.md) | breakdown_display_type:section, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**Segment1Tags** | Pointer to [**[]TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment1TagsInner**](TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment1TagsInner.md) | breakdown_display_type:segment_1_tag, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**Segment2Tags** | Pointer to [**[]TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment2TagsInner**](TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment2TagsInner.md) | breakdown_display_type:segment_2_tag, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**Segment3Tags** | Pointer to [**[]TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment3TagsInner**](TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment3TagsInner.md) | breakdown_display_type:segment_3_tag, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**TotalLine** | Pointer to **bool** | 合計行(勘定科目カテゴリーの時のみ含まれる) | [optional] 
**YearOnYear** | Pointer to **float32** | 前年比 | [optional] 

## Methods

### NewTrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner

`func NewTrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner() *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner`

NewTrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner instantiates a new TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerWithDefaults

`func NewTrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerWithDefaults() *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner`

NewTrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerWithDefaults instantiates a new TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountCategoryName

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetAccountCategoryName() string`

GetAccountCategoryName returns the AccountCategoryName field if non-nil, zero value otherwise.

### GetAccountCategoryNameOk

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetAccountCategoryNameOk() (*string, bool)`

GetAccountCategoryNameOk returns a tuple with the AccountCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCategoryName

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetAccountCategoryName(v string)`

SetAccountCategoryName sets AccountCategoryName field to given value.

### HasAccountCategoryName

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasAccountCategoryName() bool`

HasAccountCategoryName returns a boolean if a field has been set.

### GetAccountGroupName

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetAccountGroupName() string`

GetAccountGroupName returns the AccountGroupName field if non-nil, zero value otherwise.

### GetAccountGroupNameOk

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetAccountGroupNameOk() (*string, bool)`

GetAccountGroupNameOk returns a tuple with the AccountGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroupName

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetAccountGroupName(v string)`

SetAccountGroupName sets AccountGroupName field to given value.

### HasAccountGroupName

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasAccountGroupName() bool`

HasAccountGroupName returns a boolean if a field has been set.

### GetAccountItemId

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetAccountItemId() int32`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetAccountItemIdOk() (*int32, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetAccountItemId(v int32)`

SetAccountItemId sets AccountItemId field to given value.

### HasAccountItemId

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasAccountItemId() bool`

HasAccountItemId returns a boolean if a field has been set.

### GetAccountItemName

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetAccountItemName() string`

GetAccountItemName returns the AccountItemName field if non-nil, zero value otherwise.

### GetAccountItemNameOk

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetAccountItemNameOk() (*string, bool)`

GetAccountItemNameOk returns a tuple with the AccountItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemName

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetAccountItemName(v string)`

SetAccountItemName sets AccountItemName field to given value.

### HasAccountItemName

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasAccountItemName() bool`

HasAccountItemName returns a boolean if a field has been set.

### GetClosingBalance

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetClosingBalance() int32`

GetClosingBalance returns the ClosingBalance field if non-nil, zero value otherwise.

### GetClosingBalanceOk

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetClosingBalanceOk() (*int32, bool)`

GetClosingBalanceOk returns a tuple with the ClosingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingBalance

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetClosingBalance(v int32)`

SetClosingBalance sets ClosingBalance field to given value.

### HasClosingBalance

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasClosingBalance() bool`

HasClosingBalance returns a boolean if a field has been set.

### GetHierarchyLevel

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetHierarchyLevel() int32`

GetHierarchyLevel returns the HierarchyLevel field if non-nil, zero value otherwise.

### GetHierarchyLevelOk

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetHierarchyLevelOk() (*int32, bool)`

GetHierarchyLevelOk returns a tuple with the HierarchyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchyLevel

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetHierarchyLevel(v int32)`

SetHierarchyLevel sets HierarchyLevel field to given value.

### HasHierarchyLevel

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasHierarchyLevel() bool`

HasHierarchyLevel returns a boolean if a field has been set.

### GetItems

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetItems() []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetItemsOk() (*[]TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetItems(v []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLastYearClosingBalance

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetLastYearClosingBalance() int32`

GetLastYearClosingBalance returns the LastYearClosingBalance field if non-nil, zero value otherwise.

### GetLastYearClosingBalanceOk

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetLastYearClosingBalanceOk() (*int32, bool)`

GetLastYearClosingBalanceOk returns a tuple with the LastYearClosingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastYearClosingBalance

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetLastYearClosingBalance(v int32)`

SetLastYearClosingBalance sets LastYearClosingBalance field to given value.

### HasLastYearClosingBalance

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasLastYearClosingBalance() bool`

HasLastYearClosingBalance returns a boolean if a field has been set.

### GetParentAccountCategoryName

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetParentAccountCategoryName() string`

GetParentAccountCategoryName returns the ParentAccountCategoryName field if non-nil, zero value otherwise.

### GetParentAccountCategoryNameOk

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetParentAccountCategoryNameOk() (*string, bool)`

GetParentAccountCategoryNameOk returns a tuple with the ParentAccountCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountCategoryName

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetParentAccountCategoryName(v string)`

SetParentAccountCategoryName sets ParentAccountCategoryName field to given value.

### HasParentAccountCategoryName

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasParentAccountCategoryName() bool`

HasParentAccountCategoryName returns a boolean if a field has been set.

### GetPartners

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetPartners() []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerPartnersInner`

GetPartners returns the Partners field if non-nil, zero value otherwise.

### GetPartnersOk

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetPartnersOk() (*[]TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerPartnersInner, bool)`

GetPartnersOk returns a tuple with the Partners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartners

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetPartners(v []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerPartnersInner)`

SetPartners sets Partners field to given value.

### HasPartners

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasPartners() bool`

HasPartners returns a boolean if a field has been set.

### GetSections

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetSections() []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSectionsInner`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetSectionsOk() (*[]TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSectionsInner, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetSections(v []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSectionsInner)`

SetSections sets Sections field to given value.

### HasSections

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasSections() bool`

HasSections returns a boolean if a field has been set.

### GetSegment1Tags

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetSegment1Tags() []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment1TagsInner`

GetSegment1Tags returns the Segment1Tags field if non-nil, zero value otherwise.

### GetSegment1TagsOk

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetSegment1TagsOk() (*[]TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment1TagsInner, bool)`

GetSegment1TagsOk returns a tuple with the Segment1Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1Tags

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetSegment1Tags(v []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment1TagsInner)`

SetSegment1Tags sets Segment1Tags field to given value.

### HasSegment1Tags

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasSegment1Tags() bool`

HasSegment1Tags returns a boolean if a field has been set.

### GetSegment2Tags

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetSegment2Tags() []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment2TagsInner`

GetSegment2Tags returns the Segment2Tags field if non-nil, zero value otherwise.

### GetSegment2TagsOk

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetSegment2TagsOk() (*[]TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment2TagsInner, bool)`

GetSegment2TagsOk returns a tuple with the Segment2Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2Tags

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetSegment2Tags(v []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment2TagsInner)`

SetSegment2Tags sets Segment2Tags field to given value.

### HasSegment2Tags

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasSegment2Tags() bool`

HasSegment2Tags returns a boolean if a field has been set.

### GetSegment3Tags

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetSegment3Tags() []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment3TagsInner`

GetSegment3Tags returns the Segment3Tags field if non-nil, zero value otherwise.

### GetSegment3TagsOk

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetSegment3TagsOk() (*[]TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment3TagsInner, bool)`

GetSegment3TagsOk returns a tuple with the Segment3Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3Tags

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetSegment3Tags(v []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInnerSegment3TagsInner)`

SetSegment3Tags sets Segment3Tags field to given value.

### HasSegment3Tags

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasSegment3Tags() bool`

HasSegment3Tags returns a boolean if a field has been set.

### GetTotalLine

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetTotalLine() bool`

GetTotalLine returns the TotalLine field if non-nil, zero value otherwise.

### GetTotalLineOk

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetTotalLineOk() (*bool, bool)`

GetTotalLineOk returns a tuple with the TotalLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLine

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetTotalLine(v bool)`

SetTotalLine sets TotalLine field to given value.

### HasTotalLine

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasTotalLine() bool`

HasTotalLine returns a boolean if a field has been set.

### GetYearOnYear

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetYearOnYear() float32`

GetYearOnYear returns the YearOnYear field if non-nil, zero value otherwise.

### GetYearOnYearOk

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) GetYearOnYearOk() (*float32, bool)`

GetYearOnYearOk returns a tuple with the YearOnYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearOnYear

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) SetYearOnYear(v float32)`

SetYearOnYear sets YearOnYear field to given value.

### HasYearOnYear

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner) HasYearOnYear() bool`

HasYearOnYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


