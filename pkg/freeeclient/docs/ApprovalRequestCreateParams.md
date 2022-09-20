# ApprovalRequestCreateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationDate** | Pointer to **string** | 申請日 (yyyy-mm-dd)&lt;br&gt; 指定しない場合は当日の日付が登録されます。  | [optional] 
**ApprovalFlowRouteId** | **int64** | 申請経路ID | 
**ApproverId** | Pointer to **int64** | 承認者のユーザーID | [optional] 
**CompanyId** | **int64** | 事業所ID | 
**Draft** | **bool** | 各種申請のステータス&lt;br&gt; falseを指定した時は申請中（in_progress）で各種申請を作成します。&lt;br&gt; trueを指定した時は下書き（draft）で各種申請を作成します。  | 
**FormId** | **int64** | 申請フォームID | 
**ParentId** | Pointer to **int64** | 親申請ID(既存各種申請IDのみ指定可能です。) | [optional] 
**RequestItems** | [**[]ApprovalRequestCreateParamsRequestItemsInner**](ApprovalRequestCreateParamsRequestItemsInner.md) |  | 

## Methods

### NewApprovalRequestCreateParams

`func NewApprovalRequestCreateParams(approvalFlowRouteId int64, companyId int64, draft bool, formId int64, requestItems []ApprovalRequestCreateParamsRequestItemsInner, ) *ApprovalRequestCreateParams`

NewApprovalRequestCreateParams instantiates a new ApprovalRequestCreateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalRequestCreateParamsWithDefaults

`func NewApprovalRequestCreateParamsWithDefaults() *ApprovalRequestCreateParams`

NewApprovalRequestCreateParamsWithDefaults instantiates a new ApprovalRequestCreateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationDate

`func (o *ApprovalRequestCreateParams) GetApplicationDate() string`

GetApplicationDate returns the ApplicationDate field if non-nil, zero value otherwise.

### GetApplicationDateOk

`func (o *ApprovalRequestCreateParams) GetApplicationDateOk() (*string, bool)`

GetApplicationDateOk returns a tuple with the ApplicationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationDate

`func (o *ApprovalRequestCreateParams) SetApplicationDate(v string)`

SetApplicationDate sets ApplicationDate field to given value.

### HasApplicationDate

`func (o *ApprovalRequestCreateParams) HasApplicationDate() bool`

HasApplicationDate returns a boolean if a field has been set.

### GetApprovalFlowRouteId

`func (o *ApprovalRequestCreateParams) GetApprovalFlowRouteId() int64`

GetApprovalFlowRouteId returns the ApprovalFlowRouteId field if non-nil, zero value otherwise.

### GetApprovalFlowRouteIdOk

`func (o *ApprovalRequestCreateParams) GetApprovalFlowRouteIdOk() (*int64, bool)`

GetApprovalFlowRouteIdOk returns a tuple with the ApprovalFlowRouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalFlowRouteId

`func (o *ApprovalRequestCreateParams) SetApprovalFlowRouteId(v int64)`

SetApprovalFlowRouteId sets ApprovalFlowRouteId field to given value.


### GetApproverId

`func (o *ApprovalRequestCreateParams) GetApproverId() int64`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *ApprovalRequestCreateParams) GetApproverIdOk() (*int64, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *ApprovalRequestCreateParams) SetApproverId(v int64)`

SetApproverId sets ApproverId field to given value.

### HasApproverId

`func (o *ApprovalRequestCreateParams) HasApproverId() bool`

HasApproverId returns a boolean if a field has been set.

### GetCompanyId

`func (o *ApprovalRequestCreateParams) GetCompanyId() int64`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ApprovalRequestCreateParams) GetCompanyIdOk() (*int64, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ApprovalRequestCreateParams) SetCompanyId(v int64)`

SetCompanyId sets CompanyId field to given value.


### GetDraft

`func (o *ApprovalRequestCreateParams) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *ApprovalRequestCreateParams) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *ApprovalRequestCreateParams) SetDraft(v bool)`

SetDraft sets Draft field to given value.


### GetFormId

`func (o *ApprovalRequestCreateParams) GetFormId() int64`

GetFormId returns the FormId field if non-nil, zero value otherwise.

### GetFormIdOk

`func (o *ApprovalRequestCreateParams) GetFormIdOk() (*int64, bool)`

GetFormIdOk returns a tuple with the FormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormId

`func (o *ApprovalRequestCreateParams) SetFormId(v int64)`

SetFormId sets FormId field to given value.


### GetParentId

`func (o *ApprovalRequestCreateParams) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ApprovalRequestCreateParams) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ApprovalRequestCreateParams) SetParentId(v int64)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ApprovalRequestCreateParams) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetRequestItems

`func (o *ApprovalRequestCreateParams) GetRequestItems() []ApprovalRequestCreateParamsRequestItemsInner`

GetRequestItems returns the RequestItems field if non-nil, zero value otherwise.

### GetRequestItemsOk

`func (o *ApprovalRequestCreateParams) GetRequestItemsOk() (*[]ApprovalRequestCreateParamsRequestItemsInner, bool)`

GetRequestItemsOk returns a tuple with the RequestItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestItems

`func (o *ApprovalRequestCreateParams) SetRequestItems(v []ApprovalRequestCreateParamsRequestItemsInner)`

SetRequestItems sets RequestItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


