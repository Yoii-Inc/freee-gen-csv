# SegmentTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **NullableString** | 備考 | 
**Id** | **int64** | セグメントタグID | 
**Name** | **string** | セグメントタグ名 | 
**Shortcut1** | **NullableString** | ショートカット１ (20文字以内) | 
**Shortcut2** | **NullableString** | ショートカット２ (20文字以内) | 

## Methods

### NewSegmentTag

`func NewSegmentTag(description NullableString, id int64, name string, shortcut1 NullableString, shortcut2 NullableString, ) *SegmentTag`

NewSegmentTag instantiates a new SegmentTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentTagWithDefaults

`func NewSegmentTagWithDefaults() *SegmentTag`

NewSegmentTagWithDefaults instantiates a new SegmentTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SegmentTag) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SegmentTag) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SegmentTag) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *SegmentTag) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SegmentTag) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *SegmentTag) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SegmentTag) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SegmentTag) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *SegmentTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SegmentTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SegmentTag) SetName(v string)`

SetName sets Name field to given value.


### GetShortcut1

`func (o *SegmentTag) GetShortcut1() string`

GetShortcut1 returns the Shortcut1 field if non-nil, zero value otherwise.

### GetShortcut1Ok

`func (o *SegmentTag) GetShortcut1Ok() (*string, bool)`

GetShortcut1Ok returns a tuple with the Shortcut1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut1

`func (o *SegmentTag) SetShortcut1(v string)`

SetShortcut1 sets Shortcut1 field to given value.


### SetShortcut1Nil

`func (o *SegmentTag) SetShortcut1Nil(b bool)`

 SetShortcut1Nil sets the value for Shortcut1 to be an explicit nil

### UnsetShortcut1
`func (o *SegmentTag) UnsetShortcut1()`

UnsetShortcut1 ensures that no value is present for Shortcut1, not even an explicit nil
### GetShortcut2

`func (o *SegmentTag) GetShortcut2() string`

GetShortcut2 returns the Shortcut2 field if non-nil, zero value otherwise.

### GetShortcut2Ok

`func (o *SegmentTag) GetShortcut2Ok() (*string, bool)`

GetShortcut2Ok returns a tuple with the Shortcut2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut2

`func (o *SegmentTag) SetShortcut2(v string)`

SetShortcut2 sets Shortcut2 field to given value.


### SetShortcut2Nil

`func (o *SegmentTag) SetShortcut2Nil(b bool)`

 SetShortcut2Nil sets the value for Shortcut2 to be an explicit nil

### UnsetShortcut2
`func (o *SegmentTag) UnsetShortcut2()`

UnsetShortcut2 ensures that no value is present for Shortcut2, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


