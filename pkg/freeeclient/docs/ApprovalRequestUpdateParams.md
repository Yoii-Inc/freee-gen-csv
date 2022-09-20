# ApprovalRequestUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationDate** | Pointer to **string** | 申請日 (yyyy-mm-dd)&lt;br&gt; 指定しない場合は当日の日付が登録されます。  | [optional] 
**ApprovalFlowRouteId** | **int64** | 申請経路ID | 
**ApproverId** | Pointer to **int64** | 承認者のユーザーID | [optional] 
**CompanyId** | **int64** | 事業所ID | 
**Draft** | **bool** | 各種申請のステータス&lt;br&gt; falseを指定した時は申請中（in_progress）で各種申請を更新します。&lt;br&gt; trueを指定した時は下書き（draft）で各種申請を更新します。  | 
**RequestItems** | [**[]ApprovalRequestCreateParamsRequestItemsInner**](ApprovalRequestCreateParamsRequestItemsInner.md) |  | 

## Methods

### NewApprovalRequestUpdateParams

`func NewApprovalRequestUpdateParams(approvalFlowRouteId int64, companyId int64, draft bool, requestItems []ApprovalRequestCreateParamsRequestItemsInner, ) *ApprovalRequestUpdateParams`

NewApprovalRequestUpdateParams instantiates a new ApprovalRequestUpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalRequestUpdateParamsWithDefaults

`func NewApprovalRequestUpdateParamsWithDefaults() *ApprovalRequestUpdateParams`

NewApprovalRequestUpdateParamsWithDefaults instantiates a new ApprovalRequestUpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationDate

`func (o *ApprovalRequestUpdateParams) GetApplicationDate() string`

GetApplicationDate returns the ApplicationDate field if non-nil, zero value otherwise.

### GetApplicationDateOk

`func (o *ApprovalRequestUpdateParams) GetApplicationDateOk() (*string, bool)`

GetApplicationDateOk returns a tuple with the ApplicationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationDate

`func (o *ApprovalRequestUpdateParams) SetApplicationDate(v string)`

SetApplicationDate sets ApplicationDate field to given value.

### HasApplicationDate

`func (o *ApprovalRequestUpdateParams) HasApplicationDate() bool`

HasApplicationDate returns a boolean if a field has been set.

### GetApprovalFlowRouteId

`func (o *ApprovalRequestUpdateParams) GetApprovalFlowRouteId() int64`

GetApprovalFlowRouteId returns the ApprovalFlowRouteId field if non-nil, zero value otherwise.

### GetApprovalFlowRouteIdOk

`func (o *ApprovalRequestUpdateParams) GetApprovalFlowRouteIdOk() (*int64, bool)`

GetApprovalFlowRouteIdOk returns a tuple with the ApprovalFlowRouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalFlowRouteId

`func (o *ApprovalRequestUpdateParams) SetApprovalFlowRouteId(v int64)`

SetApprovalFlowRouteId sets ApprovalFlowRouteId field to given value.


### GetApproverId

`func (o *ApprovalRequestUpdateParams) GetApproverId() int64`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *ApprovalRequestUpdateParams) GetApproverIdOk() (*int64, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *ApprovalRequestUpdateParams) SetApproverId(v int64)`

SetApproverId sets ApproverId field to given value.

### HasApproverId

`func (o *ApprovalRequestUpdateParams) HasApproverId() bool`

HasApproverId returns a boolean if a field has been set.

### GetCompanyId

`func (o *ApprovalRequestUpdateParams) GetCompanyId() int64`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ApprovalRequestUpdateParams) GetCompanyIdOk() (*int64, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ApprovalRequestUpdateParams) SetCompanyId(v int64)`

SetCompanyId sets CompanyId field to given value.


### GetDraft

`func (o *ApprovalRequestUpdateParams) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *ApprovalRequestUpdateParams) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *ApprovalRequestUpdateParams) SetDraft(v bool)`

SetDraft sets Draft field to given value.


### GetRequestItems

`func (o *ApprovalRequestUpdateParams) GetRequestItems() []ApprovalRequestCreateParamsRequestItemsInner`

GetRequestItems returns the RequestItems field if non-nil, zero value otherwise.

### GetRequestItemsOk

`func (o *ApprovalRequestUpdateParams) GetRequestItemsOk() (*[]ApprovalRequestCreateParamsRequestItemsInner, bool)`

GetRequestItemsOk returns a tuple with the RequestItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestItems

`func (o *ApprovalRequestUpdateParams) SetRequestItems(v []ApprovalRequestCreateParamsRequestItemsInner)`

SetRequestItems sets RequestItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


