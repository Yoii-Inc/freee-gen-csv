# TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClosingBalance** | Pointer to **int64** | 期末残高 | [optional] 
**Id** | **int64** | 部門ID | 
**Items** | Pointer to [**[]TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerItemsInner**](TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerItemsInner.md) | breakdown_display_type:item, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**Name** | Pointer to **string** | 部門名 | [optional] 
**Partners** | Pointer to [**[]TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerPartnersInner**](TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerPartnersInner.md) | breakdown_display_type:partner, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**Segment1Tags** | Pointer to [**[]TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerSegment1TagsInner**](TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerSegment1TagsInner.md) | breakdown_display_type:segment_1_tag, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**Segment2Tags** | Pointer to [**[]TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerSegment2TagsInner**](TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerSegment2TagsInner.md) | breakdown_display_type:segment_2_tag, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**Segment3Tags** | Pointer to [**[]TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerSegment3TagsInner**](TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerSegment3TagsInner.md) | breakdown_display_type:segment_3_tag, account_item_display_type:account_item指定時のみ含まれる | [optional] 

## Methods

### NewTrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner

`func NewTrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner(id int64, ) *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner`

NewTrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner instantiates a new TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerWithDefaults

`func NewTrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerWithDefaults() *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner`

NewTrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerWithDefaults instantiates a new TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClosingBalance

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) GetClosingBalance() int64`

GetClosingBalance returns the ClosingBalance field if non-nil, zero value otherwise.

### GetClosingBalanceOk

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) GetClosingBalanceOk() (*int64, bool)`

GetClosingBalanceOk returns a tuple with the ClosingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingBalance

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) SetClosingBalance(v int64)`

SetClosingBalance sets ClosingBalance field to given value.

### HasClosingBalance

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) HasClosingBalance() bool`

HasClosingBalance returns a boolean if a field has been set.

### GetId

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) SetId(v int64)`

SetId sets Id field to given value.


### GetItems

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) GetItems() []TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) GetItemsOk() (*[]TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) SetItems(v []TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetName

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartners

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) GetPartners() []TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerPartnersInner`

GetPartners returns the Partners field if non-nil, zero value otherwise.

### GetPartnersOk

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) GetPartnersOk() (*[]TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerPartnersInner, bool)`

GetPartnersOk returns a tuple with the Partners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartners

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) SetPartners(v []TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerPartnersInner)`

SetPartners sets Partners field to given value.

### HasPartners

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) HasPartners() bool`

HasPartners returns a boolean if a field has been set.

### GetSegment1Tags

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) GetSegment1Tags() []TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerSegment1TagsInner`

GetSegment1Tags returns the Segment1Tags field if non-nil, zero value otherwise.

### GetSegment1TagsOk

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) GetSegment1TagsOk() (*[]TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerSegment1TagsInner, bool)`

GetSegment1TagsOk returns a tuple with the Segment1Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1Tags

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) SetSegment1Tags(v []TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerSegment1TagsInner)`

SetSegment1Tags sets Segment1Tags field to given value.

### HasSegment1Tags

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) HasSegment1Tags() bool`

HasSegment1Tags returns a boolean if a field has been set.

### GetSegment2Tags

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) GetSegment2Tags() []TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerSegment2TagsInner`

GetSegment2Tags returns the Segment2Tags field if non-nil, zero value otherwise.

### GetSegment2TagsOk

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) GetSegment2TagsOk() (*[]TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerSegment2TagsInner, bool)`

GetSegment2TagsOk returns a tuple with the Segment2Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2Tags

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) SetSegment2Tags(v []TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerSegment2TagsInner)`

SetSegment2Tags sets Segment2Tags field to given value.

### HasSegment2Tags

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) HasSegment2Tags() bool`

HasSegment2Tags returns a boolean if a field has been set.

### GetSegment3Tags

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) GetSegment3Tags() []TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerSegment3TagsInner`

GetSegment3Tags returns the Segment3Tags field if non-nil, zero value otherwise.

### GetSegment3TagsOk

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) GetSegment3TagsOk() (*[]TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerSegment3TagsInner, bool)`

GetSegment3TagsOk returns a tuple with the Segment3Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3Tags

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) SetSegment3Tags(v []TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerSegment3TagsInner)`

SetSegment3Tags sets Segment3Tags field to given value.

### HasSegment3Tags

`func (o *TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInner) HasSegment3Tags() bool`

HasSegment3Tags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


