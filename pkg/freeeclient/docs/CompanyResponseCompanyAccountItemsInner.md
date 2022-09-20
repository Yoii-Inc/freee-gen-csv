# CompanyResponseCompanyAccountItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | **[]string** |  | 
**DefaultTaxId** | Pointer to **int32** | デフォルト設定がされている税区分ID | [optional] 
**Id** | **int32** | 勘定科目ID | 
**Name** | **string** | 勘定科目名 (30文字以内) | 
**Shortcut** | Pointer to **NullableString** | ショートカット1 (20文字以内) | [optional] 

## Methods

### NewCompanyResponseCompanyAccountItemsInner

`func NewCompanyResponseCompanyAccountItemsInner(categories []string, id int32, name string, ) *CompanyResponseCompanyAccountItemsInner`

NewCompanyResponseCompanyAccountItemsInner instantiates a new CompanyResponseCompanyAccountItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyResponseCompanyAccountItemsInnerWithDefaults

`func NewCompanyResponseCompanyAccountItemsInnerWithDefaults() *CompanyResponseCompanyAccountItemsInner`

NewCompanyResponseCompanyAccountItemsInnerWithDefaults instantiates a new CompanyResponseCompanyAccountItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *CompanyResponseCompanyAccountItemsInner) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *CompanyResponseCompanyAccountItemsInner) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *CompanyResponseCompanyAccountItemsInner) SetCategories(v []string)`

SetCategories sets Categories field to given value.


### GetDefaultTaxId

`func (o *CompanyResponseCompanyAccountItemsInner) GetDefaultTaxId() int32`

GetDefaultTaxId returns the DefaultTaxId field if non-nil, zero value otherwise.

### GetDefaultTaxIdOk

`func (o *CompanyResponseCompanyAccountItemsInner) GetDefaultTaxIdOk() (*int32, bool)`

GetDefaultTaxIdOk returns a tuple with the DefaultTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTaxId

`func (o *CompanyResponseCompanyAccountItemsInner) SetDefaultTaxId(v int32)`

SetDefaultTaxId sets DefaultTaxId field to given value.

### HasDefaultTaxId

`func (o *CompanyResponseCompanyAccountItemsInner) HasDefaultTaxId() bool`

HasDefaultTaxId returns a boolean if a field has been set.

### GetId

`func (o *CompanyResponseCompanyAccountItemsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyResponseCompanyAccountItemsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyResponseCompanyAccountItemsInner) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *CompanyResponseCompanyAccountItemsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyResponseCompanyAccountItemsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyResponseCompanyAccountItemsInner) SetName(v string)`

SetName sets Name field to given value.


### GetShortcut

`func (o *CompanyResponseCompanyAccountItemsInner) GetShortcut() string`

GetShortcut returns the Shortcut field if non-nil, zero value otherwise.

### GetShortcutOk

`func (o *CompanyResponseCompanyAccountItemsInner) GetShortcutOk() (*string, bool)`

GetShortcutOk returns a tuple with the Shortcut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut

`func (o *CompanyResponseCompanyAccountItemsInner) SetShortcut(v string)`

SetShortcut sets Shortcut field to given value.

### HasShortcut

`func (o *CompanyResponseCompanyAccountItemsInner) HasShortcut() bool`

HasShortcut returns a boolean if a field has been set.

### SetShortcutNil

`func (o *CompanyResponseCompanyAccountItemsInner) SetShortcutNil(b bool)`

 SetShortcutNil sets the value for Shortcut to be an explicit nil

### UnsetShortcut
`func (o *CompanyResponseCompanyAccountItemsInner) UnsetShortcut()`

UnsetShortcut ensures that no value is present for Shortcut, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


