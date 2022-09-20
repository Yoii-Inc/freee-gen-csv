# ApprovalRequestResponseApprovalRequestCommentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | **string** | コメント内容 | 
**PostedAt** | **string** | コメント日時(ISO8601形式) | 
**UserId** | **int64** | ユーザーID | 

## Methods

### NewApprovalRequestResponseApprovalRequestCommentsInner

`func NewApprovalRequestResponseApprovalRequestCommentsInner(comment string, postedAt string, userId int64, ) *ApprovalRequestResponseApprovalRequestCommentsInner`

NewApprovalRequestResponseApprovalRequestCommentsInner instantiates a new ApprovalRequestResponseApprovalRequestCommentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalRequestResponseApprovalRequestCommentsInnerWithDefaults

`func NewApprovalRequestResponseApprovalRequestCommentsInnerWithDefaults() *ApprovalRequestResponseApprovalRequestCommentsInner`

NewApprovalRequestResponseApprovalRequestCommentsInnerWithDefaults instantiates a new ApprovalRequestResponseApprovalRequestCommentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ApprovalRequestResponseApprovalRequestCommentsInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ApprovalRequestResponseApprovalRequestCommentsInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ApprovalRequestResponseApprovalRequestCommentsInner) SetComment(v string)`

SetComment sets Comment field to given value.


### GetPostedAt

`func (o *ApprovalRequestResponseApprovalRequestCommentsInner) GetPostedAt() string`

GetPostedAt returns the PostedAt field if non-nil, zero value otherwise.

### GetPostedAtOk

`func (o *ApprovalRequestResponseApprovalRequestCommentsInner) GetPostedAtOk() (*string, bool)`

GetPostedAtOk returns a tuple with the PostedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedAt

`func (o *ApprovalRequestResponseApprovalRequestCommentsInner) SetPostedAt(v string)`

SetPostedAt sets PostedAt field to given value.


### GetUserId

`func (o *ApprovalRequestResponseApprovalRequestCommentsInner) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ApprovalRequestResponseApprovalRequestCommentsInner) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ApprovalRequestResponseApprovalRequestCommentsInner) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


