# ApprovalRequestResponseApprovalRequestComments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | **string** | コメント内容 | 
**PostedAt** | **string** | コメント日時(ISO8601形式) | 
**UserId** | **int32** | ユーザーID | 

## Methods

### NewApprovalRequestResponseApprovalRequestComments

`func NewApprovalRequestResponseApprovalRequestComments(comment string, postedAt string, userId int32, ) *ApprovalRequestResponseApprovalRequestComments`

NewApprovalRequestResponseApprovalRequestComments instantiates a new ApprovalRequestResponseApprovalRequestComments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalRequestResponseApprovalRequestCommentsWithDefaults

`func NewApprovalRequestResponseApprovalRequestCommentsWithDefaults() *ApprovalRequestResponseApprovalRequestComments`

NewApprovalRequestResponseApprovalRequestCommentsWithDefaults instantiates a new ApprovalRequestResponseApprovalRequestComments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ApprovalRequestResponseApprovalRequestComments) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ApprovalRequestResponseApprovalRequestComments) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ApprovalRequestResponseApprovalRequestComments) SetComment(v string)`

SetComment sets Comment field to given value.


### GetPostedAt

`func (o *ApprovalRequestResponseApprovalRequestComments) GetPostedAt() string`

GetPostedAt returns the PostedAt field if non-nil, zero value otherwise.

### GetPostedAtOk

`func (o *ApprovalRequestResponseApprovalRequestComments) GetPostedAtOk() (*string, bool)`

GetPostedAtOk returns a tuple with the PostedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedAt

`func (o *ApprovalRequestResponseApprovalRequestComments) SetPostedAt(v string)`

SetPostedAt sets PostedAt field to given value.


### GetUserId

`func (o *ApprovalRequestResponseApprovalRequestComments) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ApprovalRequestResponseApprovalRequestComments) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ApprovalRequestResponseApprovalRequestComments) SetUserId(v int32)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


