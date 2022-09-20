# UserCapabilityWithSelfOnly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedTarget** | Pointer to **string** | 「自分のみ」がonになっている場合はself_only、offになっている場合はallになります。 | [optional] 
**Create** | Pointer to **bool** | 作成 | [optional] 
**Destroy** | Pointer to **bool** | 削除 | [optional] 
**Read** | Pointer to **bool** | 閲覧 | [optional] 
**Update** | Pointer to **bool** | 更新 | [optional] 

## Methods

### NewUserCapabilityWithSelfOnly

`func NewUserCapabilityWithSelfOnly() *UserCapabilityWithSelfOnly`

NewUserCapabilityWithSelfOnly instantiates a new UserCapabilityWithSelfOnly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCapabilityWithSelfOnlyWithDefaults

`func NewUserCapabilityWithSelfOnlyWithDefaults() *UserCapabilityWithSelfOnly`

NewUserCapabilityWithSelfOnlyWithDefaults instantiates a new UserCapabilityWithSelfOnly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedTarget

`func (o *UserCapabilityWithSelfOnly) GetAllowedTarget() string`

GetAllowedTarget returns the AllowedTarget field if non-nil, zero value otherwise.

### GetAllowedTargetOk

`func (o *UserCapabilityWithSelfOnly) GetAllowedTargetOk() (*string, bool)`

GetAllowedTargetOk returns a tuple with the AllowedTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTarget

`func (o *UserCapabilityWithSelfOnly) SetAllowedTarget(v string)`

SetAllowedTarget sets AllowedTarget field to given value.

### HasAllowedTarget

`func (o *UserCapabilityWithSelfOnly) HasAllowedTarget() bool`

HasAllowedTarget returns a boolean if a field has been set.

### GetCreate

`func (o *UserCapabilityWithSelfOnly) GetCreate() bool`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *UserCapabilityWithSelfOnly) GetCreateOk() (*bool, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *UserCapabilityWithSelfOnly) SetCreate(v bool)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *UserCapabilityWithSelfOnly) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDestroy

`func (o *UserCapabilityWithSelfOnly) GetDestroy() bool`

GetDestroy returns the Destroy field if non-nil, zero value otherwise.

### GetDestroyOk

`func (o *UserCapabilityWithSelfOnly) GetDestroyOk() (*bool, bool)`

GetDestroyOk returns a tuple with the Destroy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroy

`func (o *UserCapabilityWithSelfOnly) SetDestroy(v bool)`

SetDestroy sets Destroy field to given value.

### HasDestroy

`func (o *UserCapabilityWithSelfOnly) HasDestroy() bool`

HasDestroy returns a boolean if a field has been set.

### GetRead

`func (o *UserCapabilityWithSelfOnly) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *UserCapabilityWithSelfOnly) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *UserCapabilityWithSelfOnly) SetRead(v bool)`

SetRead sets Read field to given value.

### HasRead

`func (o *UserCapabilityWithSelfOnly) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetUpdate

`func (o *UserCapabilityWithSelfOnly) GetUpdate() bool`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *UserCapabilityWithSelfOnly) GetUpdateOk() (*bool, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *UserCapabilityWithSelfOnly) SetUpdate(v bool)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *UserCapabilityWithSelfOnly) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


