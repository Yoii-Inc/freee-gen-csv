# CompanyResponseCompanyItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 品目ID | 
**Name** | **string** | 品目名 (30文字以内) | 
**Shortcut1** | Pointer to **NullableString** | ショートカット１ (20文字以内) | [optional] 
**Shortcut2** | Pointer to **NullableString** | ショートカット２ (20文字以内) | [optional] 

## Methods

### NewCompanyResponseCompanyItemsInner

`func NewCompanyResponseCompanyItemsInner(id int32, name string, ) *CompanyResponseCompanyItemsInner`

NewCompanyResponseCompanyItemsInner instantiates a new CompanyResponseCompanyItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyResponseCompanyItemsInnerWithDefaults

`func NewCompanyResponseCompanyItemsInnerWithDefaults() *CompanyResponseCompanyItemsInner`

NewCompanyResponseCompanyItemsInnerWithDefaults instantiates a new CompanyResponseCompanyItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompanyResponseCompanyItemsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyResponseCompanyItemsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyResponseCompanyItemsInner) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *CompanyResponseCompanyItemsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyResponseCompanyItemsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyResponseCompanyItemsInner) SetName(v string)`

SetName sets Name field to given value.


### GetShortcut1

`func (o *CompanyResponseCompanyItemsInner) GetShortcut1() string`

GetShortcut1 returns the Shortcut1 field if non-nil, zero value otherwise.

### GetShortcut1Ok

`func (o *CompanyResponseCompanyItemsInner) GetShortcut1Ok() (*string, bool)`

GetShortcut1Ok returns a tuple with the Shortcut1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut1

`func (o *CompanyResponseCompanyItemsInner) SetShortcut1(v string)`

SetShortcut1 sets Shortcut1 field to given value.

### HasShortcut1

`func (o *CompanyResponseCompanyItemsInner) HasShortcut1() bool`

HasShortcut1 returns a boolean if a field has been set.

### SetShortcut1Nil

`func (o *CompanyResponseCompanyItemsInner) SetShortcut1Nil(b bool)`

 SetShortcut1Nil sets the value for Shortcut1 to be an explicit nil

### UnsetShortcut1
`func (o *CompanyResponseCompanyItemsInner) UnsetShortcut1()`

UnsetShortcut1 ensures that no value is present for Shortcut1, not even an explicit nil
### GetShortcut2

`func (o *CompanyResponseCompanyItemsInner) GetShortcut2() string`

GetShortcut2 returns the Shortcut2 field if non-nil, zero value otherwise.

### GetShortcut2Ok

`func (o *CompanyResponseCompanyItemsInner) GetShortcut2Ok() (*string, bool)`

GetShortcut2Ok returns a tuple with the Shortcut2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut2

`func (o *CompanyResponseCompanyItemsInner) SetShortcut2(v string)`

SetShortcut2 sets Shortcut2 field to given value.

### HasShortcut2

`func (o *CompanyResponseCompanyItemsInner) HasShortcut2() bool`

HasShortcut2 returns a boolean if a field has been set.

### SetShortcut2Nil

`func (o *CompanyResponseCompanyItemsInner) SetShortcut2Nil(b bool)`

 SetShortcut2Nil sets the value for Shortcut2 to be an explicit nil

### UnsetShortcut2
`func (o *CompanyResponseCompanyItemsInner) UnsetShortcut2()`

UnsetShortcut2 ensures that no value is present for Shortcut2, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


