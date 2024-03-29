# ApprovalFlowRouteResponseApprovalFlowRouteStepsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | 承認ステップID | 
**NextStepId** | **NullableInt64** | 次の承認ステップID | 
**ResourceType** | **string** | 承認方法( predefined_user: メンバー指定 (1人), selected_user: 申請時にメンバー指定,unspecified: 指定なし, and_resource: メンバー指定 (複数、全員の承認), or_resource: メンバー指定 (複数、1人の承認), and_position: 役職指定 (複数、全員の承認), or_position: 役職指定 (複数、1人の承認) )  | 
**UserIds** | Pointer to **[]int64** | 承認者のユーザーID (配列) | [optional] 

## Methods

### NewApprovalFlowRouteResponseApprovalFlowRouteStepsInner

`func NewApprovalFlowRouteResponseApprovalFlowRouteStepsInner(id int64, nextStepId NullableInt64, resourceType string, ) *ApprovalFlowRouteResponseApprovalFlowRouteStepsInner`

NewApprovalFlowRouteResponseApprovalFlowRouteStepsInner instantiates a new ApprovalFlowRouteResponseApprovalFlowRouteStepsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalFlowRouteResponseApprovalFlowRouteStepsInnerWithDefaults

`func NewApprovalFlowRouteResponseApprovalFlowRouteStepsInnerWithDefaults() *ApprovalFlowRouteResponseApprovalFlowRouteStepsInner`

NewApprovalFlowRouteResponseApprovalFlowRouteStepsInnerWithDefaults instantiates a new ApprovalFlowRouteResponseApprovalFlowRouteStepsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApprovalFlowRouteResponseApprovalFlowRouteStepsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalFlowRouteResponseApprovalFlowRouteStepsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalFlowRouteResponseApprovalFlowRouteStepsInner) SetId(v int64)`

SetId sets Id field to given value.


### GetNextStepId

`func (o *ApprovalFlowRouteResponseApprovalFlowRouteStepsInner) GetNextStepId() int64`

GetNextStepId returns the NextStepId field if non-nil, zero value otherwise.

### GetNextStepIdOk

`func (o *ApprovalFlowRouteResponseApprovalFlowRouteStepsInner) GetNextStepIdOk() (*int64, bool)`

GetNextStepIdOk returns a tuple with the NextStepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextStepId

`func (o *ApprovalFlowRouteResponseApprovalFlowRouteStepsInner) SetNextStepId(v int64)`

SetNextStepId sets NextStepId field to given value.


### SetNextStepIdNil

`func (o *ApprovalFlowRouteResponseApprovalFlowRouteStepsInner) SetNextStepIdNil(b bool)`

 SetNextStepIdNil sets the value for NextStepId to be an explicit nil

### UnsetNextStepId
`func (o *ApprovalFlowRouteResponseApprovalFlowRouteStepsInner) UnsetNextStepId()`

UnsetNextStepId ensures that no value is present for NextStepId, not even an explicit nil
### GetResourceType

`func (o *ApprovalFlowRouteResponseApprovalFlowRouteStepsInner) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ApprovalFlowRouteResponseApprovalFlowRouteStepsInner) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ApprovalFlowRouteResponseApprovalFlowRouteStepsInner) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetUserIds

`func (o *ApprovalFlowRouteResponseApprovalFlowRouteStepsInner) GetUserIds() []int64`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *ApprovalFlowRouteResponseApprovalFlowRouteStepsInner) GetUserIdsOk() (*[]int64, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *ApprovalFlowRouteResponseApprovalFlowRouteStepsInner) SetUserIds(v []int64)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *ApprovalFlowRouteResponseApprovalFlowRouteStepsInner) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


