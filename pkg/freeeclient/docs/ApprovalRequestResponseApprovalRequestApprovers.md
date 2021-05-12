# ApprovalRequestResponseApprovalRequestApprovers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsForceAction** | **bool** | 代理承認済みかどうか | 
**ResourceType** | **string** | 承認ステップの承認方法 * &#x60; predefined_user&#x60; - メンバー指定 (1人), * &#x60; selected_user&#x60; - 申請時にメンバー指定 * &#x60; unspecified&#x60; - 指定なし * &#x60; and_resource&#x60; - メンバー指定 (複数、全員の承認), * &#x60; or_resource&#x60; - メンバー指定 (複数、1人の承認) * &#x60; and_position&#x60; - 役職指定 (複数、全員の承認) * &#x60; or_position&#x60; - 役職指定 (複数、1人の承認) | 
**Status** | **string** | 承認者の承認状態 * &#x60;initial&#x60; - 初期状態 * &#x60;approved&#x60; - 承認済 * &#x60;rejected&#x60; - 却下 * &#x60;feedback&#x60; - 差戻し | 
**StepId** | **int32** | 承認ステップID | 
**UserId** | **NullableInt32** | 承認者のユーザーID 下記の場合はnullになります。 &lt;ul&gt;   &lt;li&gt;resource_type:selected_userの場合で承認者未指定時&lt;/li&gt;   &lt;li&gt;resource_type:or_positionで前stepで部門未指定の場合&lt;/li&gt; &lt;/ul&gt; | 

## Methods

### NewApprovalRequestResponseApprovalRequestApprovers

`func NewApprovalRequestResponseApprovalRequestApprovers(isForceAction bool, resourceType string, status string, stepId int32, userId NullableInt32, ) *ApprovalRequestResponseApprovalRequestApprovers`

NewApprovalRequestResponseApprovalRequestApprovers instantiates a new ApprovalRequestResponseApprovalRequestApprovers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalRequestResponseApprovalRequestApproversWithDefaults

`func NewApprovalRequestResponseApprovalRequestApproversWithDefaults() *ApprovalRequestResponseApprovalRequestApprovers`

NewApprovalRequestResponseApprovalRequestApproversWithDefaults instantiates a new ApprovalRequestResponseApprovalRequestApprovers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsForceAction

`func (o *ApprovalRequestResponseApprovalRequestApprovers) GetIsForceAction() bool`

GetIsForceAction returns the IsForceAction field if non-nil, zero value otherwise.

### GetIsForceActionOk

`func (o *ApprovalRequestResponseApprovalRequestApprovers) GetIsForceActionOk() (*bool, bool)`

GetIsForceActionOk returns a tuple with the IsForceAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForceAction

`func (o *ApprovalRequestResponseApprovalRequestApprovers) SetIsForceAction(v bool)`

SetIsForceAction sets IsForceAction field to given value.


### GetResourceType

`func (o *ApprovalRequestResponseApprovalRequestApprovers) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ApprovalRequestResponseApprovalRequestApprovers) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ApprovalRequestResponseApprovalRequestApprovers) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetStatus

`func (o *ApprovalRequestResponseApprovalRequestApprovers) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApprovalRequestResponseApprovalRequestApprovers) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApprovalRequestResponseApprovalRequestApprovers) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStepId

`func (o *ApprovalRequestResponseApprovalRequestApprovers) GetStepId() int32`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *ApprovalRequestResponseApprovalRequestApprovers) GetStepIdOk() (*int32, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *ApprovalRequestResponseApprovalRequestApprovers) SetStepId(v int32)`

SetStepId sets StepId field to given value.


### GetUserId

`func (o *ApprovalRequestResponseApprovalRequestApprovers) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ApprovalRequestResponseApprovalRequestApprovers) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ApprovalRequestResponseApprovalRequestApprovers) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### SetUserIdNil

`func (o *ApprovalRequestResponseApprovalRequestApprovers) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ApprovalRequestResponseApprovalRequestApprovers) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


