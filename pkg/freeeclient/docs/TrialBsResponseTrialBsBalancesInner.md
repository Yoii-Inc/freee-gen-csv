# TrialBsResponseTrialBsBalancesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountCategoryName** | Pointer to **string** | 勘定科目カテゴリー名 | [optional] 
**AccountGroupName** | Pointer to **string** | 決算書表示名(account_item_display_type:group指定時に決算書表示名の時のみ含まれる) | [optional] 
**AccountItemId** | Pointer to **int32** | 勘定科目ID(勘定科目の時のみ含まれる) | [optional] 
**AccountItemName** | Pointer to **string** | 勘定科目名(勘定科目の時のみ含まれる) | [optional] 
**ClosingBalance** | Pointer to **int32** | 期末残高 | [optional] 
**CompositionRatio** | Pointer to **float32** | 構成比 | [optional] 
**CreditAmount** | Pointer to **int32** | 貸方金額 | [optional] 
**DebitAmount** | Pointer to **int32** | 借方金額 | [optional] 
**HierarchyLevel** | Pointer to **int32** | 階層レベル | [optional] 
**Items** | Pointer to [**[]TrialBsResponseTrialBsBalancesInnerItemsInner**](TrialBsResponseTrialBsBalancesInnerItemsInner.md) | breakdown_display_type:item, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**OpeningBalance** | Pointer to **int32** | 期首残高 | [optional] 
**ParentAccountCategoryName** | Pointer to **string** | 上位勘定科目カテゴリー名(勘定科目カテゴリーの時のみ、上層が存在する場合含まれる) | [optional] 
**Partners** | Pointer to [**[]TrialBsResponseTrialBsBalancesInnerPartnersInner**](TrialBsResponseTrialBsBalancesInnerPartnersInner.md) | breakdown_display_type:partner, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**Sections** | Pointer to [**[]TrialBsResponseTrialBsBalancesInnerSectionsInner**](TrialBsResponseTrialBsBalancesInnerSectionsInner.md) | breakdown_display_type:section, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**Segment1Tags** | Pointer to [**[]TrialBsResponseTrialBsBalancesInnerSegment1TagsInner**](TrialBsResponseTrialBsBalancesInnerSegment1TagsInner.md) | breakdown_display_type:segment_1_tag, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**Segment2Tags** | Pointer to [**[]TrialBsResponseTrialBsBalancesInnerSegment2TagsInner**](TrialBsResponseTrialBsBalancesInnerSegment2TagsInner.md) | breakdown_display_type:segment_2_tag, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**Segment3Tags** | Pointer to [**[]TrialBsResponseTrialBsBalancesInnerSegment3TagsInner**](TrialBsResponseTrialBsBalancesInnerSegment3TagsInner.md) | breakdown_display_type:segment_3_tag, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**TotalLine** | Pointer to **bool** | 合計行(勘定科目カテゴリーの時のみ含まれる) | [optional] 

## Methods

### NewTrialBsResponseTrialBsBalancesInner

`func NewTrialBsResponseTrialBsBalancesInner() *TrialBsResponseTrialBsBalancesInner`

NewTrialBsResponseTrialBsBalancesInner instantiates a new TrialBsResponseTrialBsBalancesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialBsResponseTrialBsBalancesInnerWithDefaults

`func NewTrialBsResponseTrialBsBalancesInnerWithDefaults() *TrialBsResponseTrialBsBalancesInner`

NewTrialBsResponseTrialBsBalancesInnerWithDefaults instantiates a new TrialBsResponseTrialBsBalancesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountCategoryName

`func (o *TrialBsResponseTrialBsBalancesInner) GetAccountCategoryName() string`

GetAccountCategoryName returns the AccountCategoryName field if non-nil, zero value otherwise.

### GetAccountCategoryNameOk

`func (o *TrialBsResponseTrialBsBalancesInner) GetAccountCategoryNameOk() (*string, bool)`

GetAccountCategoryNameOk returns a tuple with the AccountCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCategoryName

`func (o *TrialBsResponseTrialBsBalancesInner) SetAccountCategoryName(v string)`

SetAccountCategoryName sets AccountCategoryName field to given value.

### HasAccountCategoryName

`func (o *TrialBsResponseTrialBsBalancesInner) HasAccountCategoryName() bool`

HasAccountCategoryName returns a boolean if a field has been set.

### GetAccountGroupName

`func (o *TrialBsResponseTrialBsBalancesInner) GetAccountGroupName() string`

GetAccountGroupName returns the AccountGroupName field if non-nil, zero value otherwise.

### GetAccountGroupNameOk

`func (o *TrialBsResponseTrialBsBalancesInner) GetAccountGroupNameOk() (*string, bool)`

GetAccountGroupNameOk returns a tuple with the AccountGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroupName

`func (o *TrialBsResponseTrialBsBalancesInner) SetAccountGroupName(v string)`

SetAccountGroupName sets AccountGroupName field to given value.

### HasAccountGroupName

`func (o *TrialBsResponseTrialBsBalancesInner) HasAccountGroupName() bool`

HasAccountGroupName returns a boolean if a field has been set.

### GetAccountItemId

`func (o *TrialBsResponseTrialBsBalancesInner) GetAccountItemId() int32`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *TrialBsResponseTrialBsBalancesInner) GetAccountItemIdOk() (*int32, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *TrialBsResponseTrialBsBalancesInner) SetAccountItemId(v int32)`

SetAccountItemId sets AccountItemId field to given value.

### HasAccountItemId

`func (o *TrialBsResponseTrialBsBalancesInner) HasAccountItemId() bool`

HasAccountItemId returns a boolean if a field has been set.

### GetAccountItemName

`func (o *TrialBsResponseTrialBsBalancesInner) GetAccountItemName() string`

GetAccountItemName returns the AccountItemName field if non-nil, zero value otherwise.

### GetAccountItemNameOk

`func (o *TrialBsResponseTrialBsBalancesInner) GetAccountItemNameOk() (*string, bool)`

GetAccountItemNameOk returns a tuple with the AccountItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemName

`func (o *TrialBsResponseTrialBsBalancesInner) SetAccountItemName(v string)`

SetAccountItemName sets AccountItemName field to given value.

### HasAccountItemName

`func (o *TrialBsResponseTrialBsBalancesInner) HasAccountItemName() bool`

HasAccountItemName returns a boolean if a field has been set.

### GetClosingBalance

`func (o *TrialBsResponseTrialBsBalancesInner) GetClosingBalance() int32`

GetClosingBalance returns the ClosingBalance field if non-nil, zero value otherwise.

### GetClosingBalanceOk

`func (o *TrialBsResponseTrialBsBalancesInner) GetClosingBalanceOk() (*int32, bool)`

GetClosingBalanceOk returns a tuple with the ClosingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingBalance

`func (o *TrialBsResponseTrialBsBalancesInner) SetClosingBalance(v int32)`

SetClosingBalance sets ClosingBalance field to given value.

### HasClosingBalance

`func (o *TrialBsResponseTrialBsBalancesInner) HasClosingBalance() bool`

HasClosingBalance returns a boolean if a field has been set.

### GetCompositionRatio

`func (o *TrialBsResponseTrialBsBalancesInner) GetCompositionRatio() float32`

GetCompositionRatio returns the CompositionRatio field if non-nil, zero value otherwise.

### GetCompositionRatioOk

`func (o *TrialBsResponseTrialBsBalancesInner) GetCompositionRatioOk() (*float32, bool)`

GetCompositionRatioOk returns a tuple with the CompositionRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositionRatio

`func (o *TrialBsResponseTrialBsBalancesInner) SetCompositionRatio(v float32)`

SetCompositionRatio sets CompositionRatio field to given value.

### HasCompositionRatio

`func (o *TrialBsResponseTrialBsBalancesInner) HasCompositionRatio() bool`

HasCompositionRatio returns a boolean if a field has been set.

### GetCreditAmount

`func (o *TrialBsResponseTrialBsBalancesInner) GetCreditAmount() int32`

GetCreditAmount returns the CreditAmount field if non-nil, zero value otherwise.

### GetCreditAmountOk

`func (o *TrialBsResponseTrialBsBalancesInner) GetCreditAmountOk() (*int32, bool)`

GetCreditAmountOk returns a tuple with the CreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmount

`func (o *TrialBsResponseTrialBsBalancesInner) SetCreditAmount(v int32)`

SetCreditAmount sets CreditAmount field to given value.

### HasCreditAmount

`func (o *TrialBsResponseTrialBsBalancesInner) HasCreditAmount() bool`

HasCreditAmount returns a boolean if a field has been set.

### GetDebitAmount

`func (o *TrialBsResponseTrialBsBalancesInner) GetDebitAmount() int32`

GetDebitAmount returns the DebitAmount field if non-nil, zero value otherwise.

### GetDebitAmountOk

`func (o *TrialBsResponseTrialBsBalancesInner) GetDebitAmountOk() (*int32, bool)`

GetDebitAmountOk returns a tuple with the DebitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitAmount

`func (o *TrialBsResponseTrialBsBalancesInner) SetDebitAmount(v int32)`

SetDebitAmount sets DebitAmount field to given value.

### HasDebitAmount

`func (o *TrialBsResponseTrialBsBalancesInner) HasDebitAmount() bool`

HasDebitAmount returns a boolean if a field has been set.

### GetHierarchyLevel

`func (o *TrialBsResponseTrialBsBalancesInner) GetHierarchyLevel() int32`

GetHierarchyLevel returns the HierarchyLevel field if non-nil, zero value otherwise.

### GetHierarchyLevelOk

`func (o *TrialBsResponseTrialBsBalancesInner) GetHierarchyLevelOk() (*int32, bool)`

GetHierarchyLevelOk returns a tuple with the HierarchyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchyLevel

`func (o *TrialBsResponseTrialBsBalancesInner) SetHierarchyLevel(v int32)`

SetHierarchyLevel sets HierarchyLevel field to given value.

### HasHierarchyLevel

`func (o *TrialBsResponseTrialBsBalancesInner) HasHierarchyLevel() bool`

HasHierarchyLevel returns a boolean if a field has been set.

### GetItems

`func (o *TrialBsResponseTrialBsBalancesInner) GetItems() []TrialBsResponseTrialBsBalancesInnerItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TrialBsResponseTrialBsBalancesInner) GetItemsOk() (*[]TrialBsResponseTrialBsBalancesInnerItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TrialBsResponseTrialBsBalancesInner) SetItems(v []TrialBsResponseTrialBsBalancesInnerItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *TrialBsResponseTrialBsBalancesInner) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetOpeningBalance

`func (o *TrialBsResponseTrialBsBalancesInner) GetOpeningBalance() int32`

GetOpeningBalance returns the OpeningBalance field if non-nil, zero value otherwise.

### GetOpeningBalanceOk

`func (o *TrialBsResponseTrialBsBalancesInner) GetOpeningBalanceOk() (*int32, bool)`

GetOpeningBalanceOk returns a tuple with the OpeningBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningBalance

`func (o *TrialBsResponseTrialBsBalancesInner) SetOpeningBalance(v int32)`

SetOpeningBalance sets OpeningBalance field to given value.

### HasOpeningBalance

`func (o *TrialBsResponseTrialBsBalancesInner) HasOpeningBalance() bool`

HasOpeningBalance returns a boolean if a field has been set.

### GetParentAccountCategoryName

`func (o *TrialBsResponseTrialBsBalancesInner) GetParentAccountCategoryName() string`

GetParentAccountCategoryName returns the ParentAccountCategoryName field if non-nil, zero value otherwise.

### GetParentAccountCategoryNameOk

`func (o *TrialBsResponseTrialBsBalancesInner) GetParentAccountCategoryNameOk() (*string, bool)`

GetParentAccountCategoryNameOk returns a tuple with the ParentAccountCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountCategoryName

`func (o *TrialBsResponseTrialBsBalancesInner) SetParentAccountCategoryName(v string)`

SetParentAccountCategoryName sets ParentAccountCategoryName field to given value.

### HasParentAccountCategoryName

`func (o *TrialBsResponseTrialBsBalancesInner) HasParentAccountCategoryName() bool`

HasParentAccountCategoryName returns a boolean if a field has been set.

### GetPartners

`func (o *TrialBsResponseTrialBsBalancesInner) GetPartners() []TrialBsResponseTrialBsBalancesInnerPartnersInner`

GetPartners returns the Partners field if non-nil, zero value otherwise.

### GetPartnersOk

`func (o *TrialBsResponseTrialBsBalancesInner) GetPartnersOk() (*[]TrialBsResponseTrialBsBalancesInnerPartnersInner, bool)`

GetPartnersOk returns a tuple with the Partners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartners

`func (o *TrialBsResponseTrialBsBalancesInner) SetPartners(v []TrialBsResponseTrialBsBalancesInnerPartnersInner)`

SetPartners sets Partners field to given value.

### HasPartners

`func (o *TrialBsResponseTrialBsBalancesInner) HasPartners() bool`

HasPartners returns a boolean if a field has been set.

### GetSections

`func (o *TrialBsResponseTrialBsBalancesInner) GetSections() []TrialBsResponseTrialBsBalancesInnerSectionsInner`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *TrialBsResponseTrialBsBalancesInner) GetSectionsOk() (*[]TrialBsResponseTrialBsBalancesInnerSectionsInner, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *TrialBsResponseTrialBsBalancesInner) SetSections(v []TrialBsResponseTrialBsBalancesInnerSectionsInner)`

SetSections sets Sections field to given value.

### HasSections

`func (o *TrialBsResponseTrialBsBalancesInner) HasSections() bool`

HasSections returns a boolean if a field has been set.

### GetSegment1Tags

`func (o *TrialBsResponseTrialBsBalancesInner) GetSegment1Tags() []TrialBsResponseTrialBsBalancesInnerSegment1TagsInner`

GetSegment1Tags returns the Segment1Tags field if non-nil, zero value otherwise.

### GetSegment1TagsOk

`func (o *TrialBsResponseTrialBsBalancesInner) GetSegment1TagsOk() (*[]TrialBsResponseTrialBsBalancesInnerSegment1TagsInner, bool)`

GetSegment1TagsOk returns a tuple with the Segment1Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1Tags

`func (o *TrialBsResponseTrialBsBalancesInner) SetSegment1Tags(v []TrialBsResponseTrialBsBalancesInnerSegment1TagsInner)`

SetSegment1Tags sets Segment1Tags field to given value.

### HasSegment1Tags

`func (o *TrialBsResponseTrialBsBalancesInner) HasSegment1Tags() bool`

HasSegment1Tags returns a boolean if a field has been set.

### GetSegment2Tags

`func (o *TrialBsResponseTrialBsBalancesInner) GetSegment2Tags() []TrialBsResponseTrialBsBalancesInnerSegment2TagsInner`

GetSegment2Tags returns the Segment2Tags field if non-nil, zero value otherwise.

### GetSegment2TagsOk

`func (o *TrialBsResponseTrialBsBalancesInner) GetSegment2TagsOk() (*[]TrialBsResponseTrialBsBalancesInnerSegment2TagsInner, bool)`

GetSegment2TagsOk returns a tuple with the Segment2Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2Tags

`func (o *TrialBsResponseTrialBsBalancesInner) SetSegment2Tags(v []TrialBsResponseTrialBsBalancesInnerSegment2TagsInner)`

SetSegment2Tags sets Segment2Tags field to given value.

### HasSegment2Tags

`func (o *TrialBsResponseTrialBsBalancesInner) HasSegment2Tags() bool`

HasSegment2Tags returns a boolean if a field has been set.

### GetSegment3Tags

`func (o *TrialBsResponseTrialBsBalancesInner) GetSegment3Tags() []TrialBsResponseTrialBsBalancesInnerSegment3TagsInner`

GetSegment3Tags returns the Segment3Tags field if non-nil, zero value otherwise.

### GetSegment3TagsOk

`func (o *TrialBsResponseTrialBsBalancesInner) GetSegment3TagsOk() (*[]TrialBsResponseTrialBsBalancesInnerSegment3TagsInner, bool)`

GetSegment3TagsOk returns a tuple with the Segment3Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3Tags

`func (o *TrialBsResponseTrialBsBalancesInner) SetSegment3Tags(v []TrialBsResponseTrialBsBalancesInnerSegment3TagsInner)`

SetSegment3Tags sets Segment3Tags field to given value.

### HasSegment3Tags

`func (o *TrialBsResponseTrialBsBalancesInner) HasSegment3Tags() bool`

HasSegment3Tags returns a boolean if a field has been set.

### GetTotalLine

`func (o *TrialBsResponseTrialBsBalancesInner) GetTotalLine() bool`

GetTotalLine returns the TotalLine field if non-nil, zero value otherwise.

### GetTotalLineOk

`func (o *TrialBsResponseTrialBsBalancesInner) GetTotalLineOk() (*bool, bool)`

GetTotalLineOk returns a tuple with the TotalLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLine

`func (o *TrialBsResponseTrialBsBalancesInner) SetTotalLine(v bool)`

SetTotalLine sets TotalLine field to given value.

### HasTotalLine

`func (o *TrialBsResponseTrialBsBalancesInner) HasTotalLine() bool`

HasTotalLine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


