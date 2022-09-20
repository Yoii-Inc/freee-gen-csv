# ApprovalRequestResponseApprovalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicantId** | **int64** | 申請者のユーザーID | 
**ApplicationDate** | **string** | 申請日 (yyyy-mm-dd) | 
**ApplicationNumber** | **string** | 申請No. | 
**ApprovalFlowLogs** | [**[]ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner**](ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner.md) | 各種申請の承認履歴（配列） | 
**ApprovalFlowRouteId** | **int64** | 申請経路ID | 
**ApprovalRequestForm** | [**ApprovalRequestResponseApprovalRequestApprovalRequestForm**](ApprovalRequestResponseApprovalRequestApprovalRequestForm.md) |  | 
**Approvers** | [**[]ApprovalRequestResponseApprovalRequestApproversInner**](ApprovalRequestResponseApprovalRequestApproversInner.md) | 承認者（配列）   承認ステップのresource_typeがunspecified (指定なし)の場合はapproversはレスポンスに含まれません。   しかし、resource_typeがunspecifiedの承認ステップにおいて誰かが承認・却下・差し戻しのいずれかのアクションを取った後は、   approversはレスポンスに含まれるようになります。   その場合approversにはアクションを行ったステップのIDとアクションを行ったユーザーのIDが含まれます。 | 
**Comments** | [**[]ApprovalRequestResponseApprovalRequestCommentsInner**](ApprovalRequestResponseApprovalRequestCommentsInner.md) | 各種申請のコメント一覧（配列） | 
**CompanyId** | **int64** | 事業所ID | 
**CurrentRound** | **int64** | 現在のround。差し戻し等により申請がstepの最初からやり直しになるとroundの値が増えます。 | 
**CurrentStepId** | **NullableInt64** | 現在承認ステップID | 
**DealId** | **NullableInt64** | 取引ID (申請ステータス:statusがapprovedで、取引が存在する時のみdeal_idが表示されます) | 
**DealStatus** | **NullableString** | 取引ステータス (申請ステータス:statusがapprovedで、取引が存在する時のみdeal_statusが表示されます settled:決済済み, unsettled:未決済) | 
**FormId** | **int64** | 申請フォームID | 
**Id** | **int64** | 各種申請ID | 
**ManualJournalId** | **NullableInt64** | 振替伝票のID (申請ステータス:statusがapprovedで、関連する振替伝票が存在する時のみmanual_journal_idが表示されます)  &lt;a href&#x3D;\&quot;https://support.freee.co.jp/hc/ja/articles/115003827683-#5\&quot; target&#x3D;\&quot;_blank\&quot;&gt;承認された各種申請から支払依頼等を作成する&lt;/a&gt;  | 
**RequestItems** | [**[]ApprovalRequestResponseApprovalRequestRequestItemsInner**](ApprovalRequestResponseApprovalRequestRequestItemsInner.md) | 各種申請の項目一覧（配列） | 
**Status** | **string** | 申請ステータス(draft:下書き, in_progress:申請中, approved:承認済, rejected:却下, feedback:差戻し) | 
**Title** | **string** | 申請タイトル | 

## Methods

### NewApprovalRequestResponseApprovalRequest

`func NewApprovalRequestResponseApprovalRequest(applicantId int64, applicationDate string, applicationNumber string, approvalFlowLogs []ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner, approvalFlowRouteId int64, approvalRequestForm ApprovalRequestResponseApprovalRequestApprovalRequestForm, approvers []ApprovalRequestResponseApprovalRequestApproversInner, comments []ApprovalRequestResponseApprovalRequestCommentsInner, companyId int64, currentRound int64, currentStepId NullableInt64, dealId NullableInt64, dealStatus NullableString, formId int64, id int64, manualJournalId NullableInt64, requestItems []ApprovalRequestResponseApprovalRequestRequestItemsInner, status string, title string, ) *ApprovalRequestResponseApprovalRequest`

NewApprovalRequestResponseApprovalRequest instantiates a new ApprovalRequestResponseApprovalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalRequestResponseApprovalRequestWithDefaults

`func NewApprovalRequestResponseApprovalRequestWithDefaults() *ApprovalRequestResponseApprovalRequest`

NewApprovalRequestResponseApprovalRequestWithDefaults instantiates a new ApprovalRequestResponseApprovalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicantId

`func (o *ApprovalRequestResponseApprovalRequest) GetApplicantId() int64`

GetApplicantId returns the ApplicantId field if non-nil, zero value otherwise.

### GetApplicantIdOk

`func (o *ApprovalRequestResponseApprovalRequest) GetApplicantIdOk() (*int64, bool)`

GetApplicantIdOk returns a tuple with the ApplicantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantId

`func (o *ApprovalRequestResponseApprovalRequest) SetApplicantId(v int64)`

SetApplicantId sets ApplicantId field to given value.


### GetApplicationDate

`func (o *ApprovalRequestResponseApprovalRequest) GetApplicationDate() string`

GetApplicationDate returns the ApplicationDate field if non-nil, zero value otherwise.

### GetApplicationDateOk

`func (o *ApprovalRequestResponseApprovalRequest) GetApplicationDateOk() (*string, bool)`

GetApplicationDateOk returns a tuple with the ApplicationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationDate

`func (o *ApprovalRequestResponseApprovalRequest) SetApplicationDate(v string)`

SetApplicationDate sets ApplicationDate field to given value.


### GetApplicationNumber

`func (o *ApprovalRequestResponseApprovalRequest) GetApplicationNumber() string`

GetApplicationNumber returns the ApplicationNumber field if non-nil, zero value otherwise.

### GetApplicationNumberOk

`func (o *ApprovalRequestResponseApprovalRequest) GetApplicationNumberOk() (*string, bool)`

GetApplicationNumberOk returns a tuple with the ApplicationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationNumber

`func (o *ApprovalRequestResponseApprovalRequest) SetApplicationNumber(v string)`

SetApplicationNumber sets ApplicationNumber field to given value.


### GetApprovalFlowLogs

`func (o *ApprovalRequestResponseApprovalRequest) GetApprovalFlowLogs() []ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner`

GetApprovalFlowLogs returns the ApprovalFlowLogs field if non-nil, zero value otherwise.

### GetApprovalFlowLogsOk

`func (o *ApprovalRequestResponseApprovalRequest) GetApprovalFlowLogsOk() (*[]ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner, bool)`

GetApprovalFlowLogsOk returns a tuple with the ApprovalFlowLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalFlowLogs

`func (o *ApprovalRequestResponseApprovalRequest) SetApprovalFlowLogs(v []ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner)`

SetApprovalFlowLogs sets ApprovalFlowLogs field to given value.


### GetApprovalFlowRouteId

`func (o *ApprovalRequestResponseApprovalRequest) GetApprovalFlowRouteId() int64`

GetApprovalFlowRouteId returns the ApprovalFlowRouteId field if non-nil, zero value otherwise.

### GetApprovalFlowRouteIdOk

`func (o *ApprovalRequestResponseApprovalRequest) GetApprovalFlowRouteIdOk() (*int64, bool)`

GetApprovalFlowRouteIdOk returns a tuple with the ApprovalFlowRouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalFlowRouteId

`func (o *ApprovalRequestResponseApprovalRequest) SetApprovalFlowRouteId(v int64)`

SetApprovalFlowRouteId sets ApprovalFlowRouteId field to given value.


### GetApprovalRequestForm

`func (o *ApprovalRequestResponseApprovalRequest) GetApprovalRequestForm() ApprovalRequestResponseApprovalRequestApprovalRequestForm`

GetApprovalRequestForm returns the ApprovalRequestForm field if non-nil, zero value otherwise.

### GetApprovalRequestFormOk

`func (o *ApprovalRequestResponseApprovalRequest) GetApprovalRequestFormOk() (*ApprovalRequestResponseApprovalRequestApprovalRequestForm, bool)`

GetApprovalRequestFormOk returns a tuple with the ApprovalRequestForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalRequestForm

`func (o *ApprovalRequestResponseApprovalRequest) SetApprovalRequestForm(v ApprovalRequestResponseApprovalRequestApprovalRequestForm)`

SetApprovalRequestForm sets ApprovalRequestForm field to given value.


### GetApprovers

`func (o *ApprovalRequestResponseApprovalRequest) GetApprovers() []ApprovalRequestResponseApprovalRequestApproversInner`

GetApprovers returns the Approvers field if non-nil, zero value otherwise.

### GetApproversOk

`func (o *ApprovalRequestResponseApprovalRequest) GetApproversOk() (*[]ApprovalRequestResponseApprovalRequestApproversInner, bool)`

GetApproversOk returns a tuple with the Approvers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovers

`func (o *ApprovalRequestResponseApprovalRequest) SetApprovers(v []ApprovalRequestResponseApprovalRequestApproversInner)`

SetApprovers sets Approvers field to given value.


### GetComments

`func (o *ApprovalRequestResponseApprovalRequest) GetComments() []ApprovalRequestResponseApprovalRequestCommentsInner`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ApprovalRequestResponseApprovalRequest) GetCommentsOk() (*[]ApprovalRequestResponseApprovalRequestCommentsInner, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ApprovalRequestResponseApprovalRequest) SetComments(v []ApprovalRequestResponseApprovalRequestCommentsInner)`

SetComments sets Comments field to given value.


### GetCompanyId

`func (o *ApprovalRequestResponseApprovalRequest) GetCompanyId() int64`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ApprovalRequestResponseApprovalRequest) GetCompanyIdOk() (*int64, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ApprovalRequestResponseApprovalRequest) SetCompanyId(v int64)`

SetCompanyId sets CompanyId field to given value.


### GetCurrentRound

`func (o *ApprovalRequestResponseApprovalRequest) GetCurrentRound() int64`

GetCurrentRound returns the CurrentRound field if non-nil, zero value otherwise.

### GetCurrentRoundOk

`func (o *ApprovalRequestResponseApprovalRequest) GetCurrentRoundOk() (*int64, bool)`

GetCurrentRoundOk returns a tuple with the CurrentRound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRound

`func (o *ApprovalRequestResponseApprovalRequest) SetCurrentRound(v int64)`

SetCurrentRound sets CurrentRound field to given value.


### GetCurrentStepId

`func (o *ApprovalRequestResponseApprovalRequest) GetCurrentStepId() int64`

GetCurrentStepId returns the CurrentStepId field if non-nil, zero value otherwise.

### GetCurrentStepIdOk

`func (o *ApprovalRequestResponseApprovalRequest) GetCurrentStepIdOk() (*int64, bool)`

GetCurrentStepIdOk returns a tuple with the CurrentStepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStepId

`func (o *ApprovalRequestResponseApprovalRequest) SetCurrentStepId(v int64)`

SetCurrentStepId sets CurrentStepId field to given value.


### SetCurrentStepIdNil

`func (o *ApprovalRequestResponseApprovalRequest) SetCurrentStepIdNil(b bool)`

 SetCurrentStepIdNil sets the value for CurrentStepId to be an explicit nil

### UnsetCurrentStepId
`func (o *ApprovalRequestResponseApprovalRequest) UnsetCurrentStepId()`

UnsetCurrentStepId ensures that no value is present for CurrentStepId, not even an explicit nil
### GetDealId

`func (o *ApprovalRequestResponseApprovalRequest) GetDealId() int64`

GetDealId returns the DealId field if non-nil, zero value otherwise.

### GetDealIdOk

`func (o *ApprovalRequestResponseApprovalRequest) GetDealIdOk() (*int64, bool)`

GetDealIdOk returns a tuple with the DealId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealId

`func (o *ApprovalRequestResponseApprovalRequest) SetDealId(v int64)`

SetDealId sets DealId field to given value.


### SetDealIdNil

`func (o *ApprovalRequestResponseApprovalRequest) SetDealIdNil(b bool)`

 SetDealIdNil sets the value for DealId to be an explicit nil

### UnsetDealId
`func (o *ApprovalRequestResponseApprovalRequest) UnsetDealId()`

UnsetDealId ensures that no value is present for DealId, not even an explicit nil
### GetDealStatus

`func (o *ApprovalRequestResponseApprovalRequest) GetDealStatus() string`

GetDealStatus returns the DealStatus field if non-nil, zero value otherwise.

### GetDealStatusOk

`func (o *ApprovalRequestResponseApprovalRequest) GetDealStatusOk() (*string, bool)`

GetDealStatusOk returns a tuple with the DealStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealStatus

`func (o *ApprovalRequestResponseApprovalRequest) SetDealStatus(v string)`

SetDealStatus sets DealStatus field to given value.


### SetDealStatusNil

`func (o *ApprovalRequestResponseApprovalRequest) SetDealStatusNil(b bool)`

 SetDealStatusNil sets the value for DealStatus to be an explicit nil

### UnsetDealStatus
`func (o *ApprovalRequestResponseApprovalRequest) UnsetDealStatus()`

UnsetDealStatus ensures that no value is present for DealStatus, not even an explicit nil
### GetFormId

`func (o *ApprovalRequestResponseApprovalRequest) GetFormId() int64`

GetFormId returns the FormId field if non-nil, zero value otherwise.

### GetFormIdOk

`func (o *ApprovalRequestResponseApprovalRequest) GetFormIdOk() (*int64, bool)`

GetFormIdOk returns a tuple with the FormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormId

`func (o *ApprovalRequestResponseApprovalRequest) SetFormId(v int64)`

SetFormId sets FormId field to given value.


### GetId

`func (o *ApprovalRequestResponseApprovalRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalRequestResponseApprovalRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalRequestResponseApprovalRequest) SetId(v int64)`

SetId sets Id field to given value.


### GetManualJournalId

`func (o *ApprovalRequestResponseApprovalRequest) GetManualJournalId() int64`

GetManualJournalId returns the ManualJournalId field if non-nil, zero value otherwise.

### GetManualJournalIdOk

`func (o *ApprovalRequestResponseApprovalRequest) GetManualJournalIdOk() (*int64, bool)`

GetManualJournalIdOk returns a tuple with the ManualJournalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualJournalId

`func (o *ApprovalRequestResponseApprovalRequest) SetManualJournalId(v int64)`

SetManualJournalId sets ManualJournalId field to given value.


### SetManualJournalIdNil

`func (o *ApprovalRequestResponseApprovalRequest) SetManualJournalIdNil(b bool)`

 SetManualJournalIdNil sets the value for ManualJournalId to be an explicit nil

### UnsetManualJournalId
`func (o *ApprovalRequestResponseApprovalRequest) UnsetManualJournalId()`

UnsetManualJournalId ensures that no value is present for ManualJournalId, not even an explicit nil
### GetRequestItems

`func (o *ApprovalRequestResponseApprovalRequest) GetRequestItems() []ApprovalRequestResponseApprovalRequestRequestItemsInner`

GetRequestItems returns the RequestItems field if non-nil, zero value otherwise.

### GetRequestItemsOk

`func (o *ApprovalRequestResponseApprovalRequest) GetRequestItemsOk() (*[]ApprovalRequestResponseApprovalRequestRequestItemsInner, bool)`

GetRequestItemsOk returns a tuple with the RequestItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestItems

`func (o *ApprovalRequestResponseApprovalRequest) SetRequestItems(v []ApprovalRequestResponseApprovalRequestRequestItemsInner)`

SetRequestItems sets RequestItems field to given value.


### GetStatus

`func (o *ApprovalRequestResponseApprovalRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApprovalRequestResponseApprovalRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApprovalRequestResponseApprovalRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *ApprovalRequestResponseApprovalRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApprovalRequestResponseApprovalRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApprovalRequestResponseApprovalRequest) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


