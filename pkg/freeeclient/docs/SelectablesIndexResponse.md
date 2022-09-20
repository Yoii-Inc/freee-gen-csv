# SelectablesIndexResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountCategories** | Pointer to [**[]SelectablesIndexResponseAccountCategoriesInner**](SelectablesIndexResponseAccountCategoriesInner.md) |  | [optional] 
**AccountGroups** | Pointer to [**[]SelectablesIndexResponseAccountGroupsInner**](SelectablesIndexResponseAccountGroupsInner.md) | 決算書表示名（小カテゴリー） | [optional] 

## Methods

### NewSelectablesIndexResponse

`func NewSelectablesIndexResponse() *SelectablesIndexResponse`

NewSelectablesIndexResponse instantiates a new SelectablesIndexResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectablesIndexResponseWithDefaults

`func NewSelectablesIndexResponseWithDefaults() *SelectablesIndexResponse`

NewSelectablesIndexResponseWithDefaults instantiates a new SelectablesIndexResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountCategories

`func (o *SelectablesIndexResponse) GetAccountCategories() []SelectablesIndexResponseAccountCategoriesInner`

GetAccountCategories returns the AccountCategories field if non-nil, zero value otherwise.

### GetAccountCategoriesOk

`func (o *SelectablesIndexResponse) GetAccountCategoriesOk() (*[]SelectablesIndexResponseAccountCategoriesInner, bool)`

GetAccountCategoriesOk returns a tuple with the AccountCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCategories

`func (o *SelectablesIndexResponse) SetAccountCategories(v []SelectablesIndexResponseAccountCategoriesInner)`

SetAccountCategories sets AccountCategories field to given value.

### HasAccountCategories

`func (o *SelectablesIndexResponse) HasAccountCategories() bool`

HasAccountCategories returns a boolean if a field has been set.

### GetAccountGroups

`func (o *SelectablesIndexResponse) GetAccountGroups() []SelectablesIndexResponseAccountGroupsInner`

GetAccountGroups returns the AccountGroups field if non-nil, zero value otherwise.

### GetAccountGroupsOk

`func (o *SelectablesIndexResponse) GetAccountGroupsOk() (*[]SelectablesIndexResponseAccountGroupsInner, bool)`

GetAccountGroupsOk returns a tuple with the AccountGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroups

`func (o *SelectablesIndexResponse) SetAccountGroups(v []SelectablesIndexResponseAccountGroupsInner)`

SetAccountGroups sets AccountGroups field to given value.

### HasAccountGroups

`func (o *SelectablesIndexResponse) HasAccountGroups() bool`

HasAccountGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


