# ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | 操作(apply: 申請, approve: 承認, force_approve: 代理承認, cancel: 取消, reject: 却下, feedback: 差戻し) | 
**UpdatedAt** | **string** | 更新日時(ISO8601形式) | 
**UserId** | **int64** | ユーザーID | 

## Methods

### NewApprovalRequestResponseApprovalRequestApprovalFlowLogsInner

`func NewApprovalRequestResponseApprovalRequestApprovalFlowLogsInner(action string, updatedAt string, userId int64, ) *ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner`

NewApprovalRequestResponseApprovalRequestApprovalFlowLogsInner instantiates a new ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalRequestResponseApprovalRequestApprovalFlowLogsInnerWithDefaults

`func NewApprovalRequestResponseApprovalRequestApprovalFlowLogsInnerWithDefaults() *ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner`

NewApprovalRequestResponseApprovalRequestApprovalFlowLogsInnerWithDefaults instantiates a new ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner) SetAction(v string)`

SetAction sets Action field to given value.


### GetUpdatedAt

`func (o *ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserId

`func (o *ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


