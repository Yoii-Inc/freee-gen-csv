# MeResponseUserCompaniesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvisorId** | Pointer to **NullableInt64** | アドバイザープロファイルID（アドバイザー事業所で無い場合にnullになります） | [optional] 
**DisplayName** | **string** | 表示名 | 
**Id** | **int64** | 事業所ID | 
**Role** | **string** | ユーザーの権限 &lt;ul&gt; &lt;li&gt;admin: 管理者&lt;/li&gt; &lt;li&gt;simple_accounting: 一般&lt;/li&gt; &lt;li&gt;self_only: 取引登録のみ&lt;/li&gt; &lt;li&gt;read_only: 閲覧のみ&lt;/li&gt; &lt;li&gt;workflow: 申請・承認&lt;/li&gt; &lt;/ul&gt; | 
**UseCustomRole** | **bool** | カスタム権限（true: 使用する、false: 使用しない） | 

## Methods

### NewMeResponseUserCompaniesInner

`func NewMeResponseUserCompaniesInner(displayName string, id int64, role string, useCustomRole bool, ) *MeResponseUserCompaniesInner`

NewMeResponseUserCompaniesInner instantiates a new MeResponseUserCompaniesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeResponseUserCompaniesInnerWithDefaults

`func NewMeResponseUserCompaniesInnerWithDefaults() *MeResponseUserCompaniesInner`

NewMeResponseUserCompaniesInnerWithDefaults instantiates a new MeResponseUserCompaniesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvisorId

`func (o *MeResponseUserCompaniesInner) GetAdvisorId() int64`

GetAdvisorId returns the AdvisorId field if non-nil, zero value otherwise.

### GetAdvisorIdOk

`func (o *MeResponseUserCompaniesInner) GetAdvisorIdOk() (*int64, bool)`

GetAdvisorIdOk returns a tuple with the AdvisorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisorId

`func (o *MeResponseUserCompaniesInner) SetAdvisorId(v int64)`

SetAdvisorId sets AdvisorId field to given value.

### HasAdvisorId

`func (o *MeResponseUserCompaniesInner) HasAdvisorId() bool`

HasAdvisorId returns a boolean if a field has been set.

### SetAdvisorIdNil

`func (o *MeResponseUserCompaniesInner) SetAdvisorIdNil(b bool)`

 SetAdvisorIdNil sets the value for AdvisorId to be an explicit nil

### UnsetAdvisorId
`func (o *MeResponseUserCompaniesInner) UnsetAdvisorId()`

UnsetAdvisorId ensures that no value is present for AdvisorId, not even an explicit nil
### GetDisplayName

`func (o *MeResponseUserCompaniesInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MeResponseUserCompaniesInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MeResponseUserCompaniesInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetId

`func (o *MeResponseUserCompaniesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MeResponseUserCompaniesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MeResponseUserCompaniesInner) SetId(v int64)`

SetId sets Id field to given value.


### GetRole

`func (o *MeResponseUserCompaniesInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MeResponseUserCompaniesInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MeResponseUserCompaniesInner) SetRole(v string)`

SetRole sets Role field to given value.


### GetUseCustomRole

`func (o *MeResponseUserCompaniesInner) GetUseCustomRole() bool`

GetUseCustomRole returns the UseCustomRole field if non-nil, zero value otherwise.

### GetUseCustomRoleOk

`func (o *MeResponseUserCompaniesInner) GetUseCustomRoleOk() (*bool, bool)`

GetUseCustomRoleOk returns a tuple with the UseCustomRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCustomRole

`func (o *MeResponseUserCompaniesInner) SetUseCustomRole(v bool)`

SetUseCustomRole sets UseCustomRole field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


