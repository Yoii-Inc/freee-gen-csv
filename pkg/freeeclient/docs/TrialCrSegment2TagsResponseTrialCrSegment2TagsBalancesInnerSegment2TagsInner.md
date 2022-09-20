# TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClosingBalance** | Pointer to **int64** | 期末残高 | [optional] 
**Id** | **int64** | セグメント2タグID | 
**Items** | Pointer to [**[]TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerItemsInner**](TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerItemsInner.md) | breakdown_display_type:item, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**Name** | Pointer to **string** | セグメント2タグ名 | [optional] 
**Partners** | Pointer to [**[]TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerPartnersInner**](TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerPartnersInner.md) | breakdown_display_type:partner, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**Sections** | Pointer to [**[]TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInnerSegment1TagsInnerSectionsInner**](TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInnerSegment1TagsInnerSectionsInner.md) | breakdown_display_type:section, account_item_display_type:account_item指定時のみ含まれる | [optional] 

## Methods

### NewTrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner

`func NewTrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner(id int64, ) *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner`

NewTrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner instantiates a new TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInnerWithDefaults

`func NewTrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInnerWithDefaults() *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner`

NewTrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInnerWithDefaults instantiates a new TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClosingBalance

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner) GetClosingBalance() int64`

GetClosingBalance returns the ClosingBalance field if non-nil, zero value otherwise.

### GetClosingBalanceOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner) GetClosingBalanceOk() (*int64, bool)`

GetClosingBalanceOk returns a tuple with the ClosingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingBalance

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner) SetClosingBalance(v int64)`

SetClosingBalance sets ClosingBalance field to given value.

### HasClosingBalance

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner) HasClosingBalance() bool`

HasClosingBalance returns a boolean if a field has been set.

### GetId

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner) SetId(v int64)`

SetId sets Id field to given value.


### GetItems

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner) GetItems() []TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner) GetItemsOk() (*[]TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner) SetItems(v []TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetName

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartners

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner) GetPartners() []TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerPartnersInner`

GetPartners returns the Partners field if non-nil, zero value otherwise.

### GetPartnersOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner) GetPartnersOk() (*[]TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerPartnersInner, bool)`

GetPartnersOk returns a tuple with the Partners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartners

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner) SetPartners(v []TrialCrSectionsResponseTrialCrSectionsBalancesInnerSectionsInnerPartnersInner)`

SetPartners sets Partners field to given value.

### HasPartners

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner) HasPartners() bool`

HasPartners returns a boolean if a field has been set.

### GetSections

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner) GetSections() []TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInnerSegment1TagsInnerSectionsInner`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner) GetSectionsOk() (*[]TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInnerSegment1TagsInnerSectionsInner, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner) SetSections(v []TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInnerSegment1TagsInnerSectionsInner)`

SetSections sets Sections field to given value.

### HasSections

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInnerSegment2TagsInner) HasSections() bool`

HasSections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


