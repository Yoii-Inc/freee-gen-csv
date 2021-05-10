# MeResponseUserCompanies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvisorId** | Pointer to **NullableInt32** | アドバイザープロファイルID（アドバイザー事業所で無い場合にnullになります） | [optional] 
**DisplayName** | **string** | 表示名 | 
**Id** | **int32** | 事業所ID | 
**Role** | **string** | ユーザーの権限 | 
**UseCustomRole** | **bool** | カスタム権限（true: 使用する、false: 使用しない） | 

## Methods

### NewMeResponseUserCompanies

`func NewMeResponseUserCompanies(displayName string, id int32, role string, useCustomRole bool, ) *MeResponseUserCompanies`

NewMeResponseUserCompanies instantiates a new MeResponseUserCompanies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeResponseUserCompaniesWithDefaults

`func NewMeResponseUserCompaniesWithDefaults() *MeResponseUserCompanies`

NewMeResponseUserCompaniesWithDefaults instantiates a new MeResponseUserCompanies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvisorId

`func (o *MeResponseUserCompanies) GetAdvisorId() int32`

GetAdvisorId returns the AdvisorId field if non-nil, zero value otherwise.

### GetAdvisorIdOk

`func (o *MeResponseUserCompanies) GetAdvisorIdOk() (*int32, bool)`

GetAdvisorIdOk returns a tuple with the AdvisorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisorId

`func (o *MeResponseUserCompanies) SetAdvisorId(v int32)`

SetAdvisorId sets AdvisorId field to given value.

### HasAdvisorId

`func (o *MeResponseUserCompanies) HasAdvisorId() bool`

HasAdvisorId returns a boolean if a field has been set.

### SetAdvisorIdNil

`func (o *MeResponseUserCompanies) SetAdvisorIdNil(b bool)`

 SetAdvisorIdNil sets the value for AdvisorId to be an explicit nil

### UnsetAdvisorId
`func (o *MeResponseUserCompanies) UnsetAdvisorId()`

UnsetAdvisorId ensures that no value is present for AdvisorId, not even an explicit nil
### GetDisplayName

`func (o *MeResponseUserCompanies) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MeResponseUserCompanies) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MeResponseUserCompanies) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetId

`func (o *MeResponseUserCompanies) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MeResponseUserCompanies) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MeResponseUserCompanies) SetId(v int32)`

SetId sets Id field to given value.


### GetRole

`func (o *MeResponseUserCompanies) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MeResponseUserCompanies) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MeResponseUserCompanies) SetRole(v string)`

SetRole sets Role field to given value.


### GetUseCustomRole

`func (o *MeResponseUserCompanies) GetUseCustomRole() bool`

GetUseCustomRole returns the UseCustomRole field if non-nil, zero value otherwise.

### GetUseCustomRoleOk

`func (o *MeResponseUserCompanies) GetUseCustomRoleOk() (*bool, bool)`

GetUseCustomRoleOk returns a tuple with the UseCustomRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCustomRole

`func (o *MeResponseUserCompanies) SetUseCustomRole(v bool)`

SetUseCustomRole sets UseCustomRole field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


