# ExpenseApplicationResponseExpenseApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicantId** | **int64** | 申請者のユーザーID | 
**ApplicationNumber** | **string** | 申請No. | 
**ApprovalFlowLogs** | [**[]ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner**](ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner.md) | 経費申請の承認履歴（配列） | 
**ApprovalFlowRouteId** | **int64** | 申請経路ID | 
**Approvers** | [**[]ApprovalRequestResponseApprovalRequestApproversInner**](ApprovalRequestResponseApprovalRequestApproversInner.md) | 承認者（配列）   承認ステップのresource_typeがunspecified (指定なし)の場合はapproversはレスポンスに含まれません。   しかし、resource_typeがunspecifiedの承認ステップにおいて誰かが承認・却下・差し戻しのいずれかのアクションを取った後は、   approversはレスポンスに含まれるようになります。   その場合approversにはアクションを行ったステップのIDとアクションを行ったユーザーのIDが含まれます。 | 
**Comments** | [**[]ApprovalRequestResponseApprovalRequestCommentsInner**](ApprovalRequestResponseApprovalRequestCommentsInner.md) | 経費申請のコメント一覧（配列） | 
**CompanyId** | **int64** | 事業所ID | 
**CurrentRound** | **int64** | 現在のround。差し戻し等により申請がstepの最初からやり直しになるとroundの値が増えます。 | 
**CurrentStepId** | **NullableInt64** | 現在承認ステップID | 
**DealId** | **NullableInt64** | 取引ID (申請ステータス:statusがapprovedで、取引が存在する時のみdeal_idが表示されます) | 
**DealStatus** | **NullableString** | 取引ステータス (申請ステータス:statusがapprovedで、取引が存在する時のみdeal_statusが表示されます settled:精算済み, unsettled:清算待ち) | 
**Description** | Pointer to **NullableString** | 備考 | [optional] 
**ExpenseApplicationLines** | [**[]ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner**](ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner.md) | 経費申請の項目行一覧（配列） | 
**Id** | **int64** | 経費申請ID | 
**IssueDate** | **string** | 申請日 (yyyy-mm-dd) | 
**SectionId** | Pointer to **NullableInt64** | 部門ID | [optional] 
**Segment1TagId** | Pointer to **NullableInt64** | セグメント１ID。セグメント１が使用可能なプランの時のみレスポンスに含まれます。 | [optional] 
**Segment2TagId** | Pointer to **NullableInt64** | セグメント２ID。セグメント２が使用可能なプランの時のみレスポンスに含まれます。 | [optional] 
**Segment3TagId** | Pointer to **NullableInt64** | セグメント３ID。セグメント３が使用可能なプランの時のみレスポンスに含まれます。 | [optional] 
**Status** | **string** | 申請ステータス(draft:下書き, in_progress:申請中, approved:承認済, rejected:却下, feedback:差戻し) | 
**TagIds** | Pointer to **[]int64** | メモタグID | [optional] 
**Title** | **string** | 申請タイトル | 
**TotalAmount** | Pointer to **int64** | 合計金額 | [optional] 

## Methods

### NewExpenseApplicationResponseExpenseApplication

`func NewExpenseApplicationResponseExpenseApplication(applicantId int64, applicationNumber string, approvalFlowLogs []ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner, approvalFlowRouteId int64, approvers []ApprovalRequestResponseApprovalRequestApproversInner, comments []ApprovalRequestResponseApprovalRequestCommentsInner, companyId int64, currentRound int64, currentStepId NullableInt64, dealId NullableInt64, dealStatus NullableString, expenseApplicationLines []ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner, id int64, issueDate string, status string, title string, ) *ExpenseApplicationResponseExpenseApplication`

NewExpenseApplicationResponseExpenseApplication instantiates a new ExpenseApplicationResponseExpenseApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseApplicationResponseExpenseApplicationWithDefaults

`func NewExpenseApplicationResponseExpenseApplicationWithDefaults() *ExpenseApplicationResponseExpenseApplication`

NewExpenseApplicationResponseExpenseApplicationWithDefaults instantiates a new ExpenseApplicationResponseExpenseApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicantId

`func (o *ExpenseApplicationResponseExpenseApplication) GetApplicantId() int64`

GetApplicantId returns the ApplicantId field if non-nil, zero value otherwise.

### GetApplicantIdOk

`func (o *ExpenseApplicationResponseExpenseApplication) GetApplicantIdOk() (*int64, bool)`

GetApplicantIdOk returns a tuple with the ApplicantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantId

`func (o *ExpenseApplicationResponseExpenseApplication) SetApplicantId(v int64)`

SetApplicantId sets ApplicantId field to given value.


### GetApplicationNumber

`func (o *ExpenseApplicationResponseExpenseApplication) GetApplicationNumber() string`

GetApplicationNumber returns the ApplicationNumber field if non-nil, zero value otherwise.

### GetApplicationNumberOk

`func (o *ExpenseApplicationResponseExpenseApplication) GetApplicationNumberOk() (*string, bool)`

GetApplicationNumberOk returns a tuple with the ApplicationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationNumber

`func (o *ExpenseApplicationResponseExpenseApplication) SetApplicationNumber(v string)`

SetApplicationNumber sets ApplicationNumber field to given value.


### GetApprovalFlowLogs

`func (o *ExpenseApplicationResponseExpenseApplication) GetApprovalFlowLogs() []ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner`

GetApprovalFlowLogs returns the ApprovalFlowLogs field if non-nil, zero value otherwise.

### GetApprovalFlowLogsOk

`func (o *ExpenseApplicationResponseExpenseApplication) GetApprovalFlowLogsOk() (*[]ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner, bool)`

GetApprovalFlowLogsOk returns a tuple with the ApprovalFlowLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalFlowLogs

`func (o *ExpenseApplicationResponseExpenseApplication) SetApprovalFlowLogs(v []ApprovalRequestResponseApprovalRequestApprovalFlowLogsInner)`

SetApprovalFlowLogs sets ApprovalFlowLogs field to given value.


### GetApprovalFlowRouteId

`func (o *ExpenseApplicationResponseExpenseApplication) GetApprovalFlowRouteId() int64`

GetApprovalFlowRouteId returns the ApprovalFlowRouteId field if non-nil, zero value otherwise.

### GetApprovalFlowRouteIdOk

`func (o *ExpenseApplicationResponseExpenseApplication) GetApprovalFlowRouteIdOk() (*int64, bool)`

GetApprovalFlowRouteIdOk returns a tuple with the ApprovalFlowRouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalFlowRouteId

`func (o *ExpenseApplicationResponseExpenseApplication) SetApprovalFlowRouteId(v int64)`

SetApprovalFlowRouteId sets ApprovalFlowRouteId field to given value.


### GetApprovers

`func (o *ExpenseApplicationResponseExpenseApplication) GetApprovers() []ApprovalRequestResponseApprovalRequestApproversInner`

GetApprovers returns the Approvers field if non-nil, zero value otherwise.

### GetApproversOk

`func (o *ExpenseApplicationResponseExpenseApplication) GetApproversOk() (*[]ApprovalRequestResponseApprovalRequestApproversInner, bool)`

GetApproversOk returns a tuple with the Approvers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovers

`func (o *ExpenseApplicationResponseExpenseApplication) SetApprovers(v []ApprovalRequestResponseApprovalRequestApproversInner)`

SetApprovers sets Approvers field to given value.


### GetComments

`func (o *ExpenseApplicationResponseExpenseApplication) GetComments() []ApprovalRequestResponseApprovalRequestCommentsInner`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ExpenseApplicationResponseExpenseApplication) GetCommentsOk() (*[]ApprovalRequestResponseApprovalRequestCommentsInner, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ExpenseApplicationResponseExpenseApplication) SetComments(v []ApprovalRequestResponseApprovalRequestCommentsInner)`

SetComments sets Comments field to given value.


### GetCompanyId

`func (o *ExpenseApplicationResponseExpenseApplication) GetCompanyId() int64`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ExpenseApplicationResponseExpenseApplication) GetCompanyIdOk() (*int64, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ExpenseApplicationResponseExpenseApplication) SetCompanyId(v int64)`

SetCompanyId sets CompanyId field to given value.


### GetCurrentRound

`func (o *ExpenseApplicationResponseExpenseApplication) GetCurrentRound() int64`

GetCurrentRound returns the CurrentRound field if non-nil, zero value otherwise.

### GetCurrentRoundOk

`func (o *ExpenseApplicationResponseExpenseApplication) GetCurrentRoundOk() (*int64, bool)`

GetCurrentRoundOk returns a tuple with the CurrentRound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRound

`func (o *ExpenseApplicationResponseExpenseApplication) SetCurrentRound(v int64)`

SetCurrentRound sets CurrentRound field to given value.


### GetCurrentStepId

`func (o *ExpenseApplicationResponseExpenseApplication) GetCurrentStepId() int64`

GetCurrentStepId returns the CurrentStepId field if non-nil, zero value otherwise.

### GetCurrentStepIdOk

`func (o *ExpenseApplicationResponseExpenseApplication) GetCurrentStepIdOk() (*int64, bool)`

GetCurrentStepIdOk returns a tuple with the CurrentStepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStepId

`func (o *ExpenseApplicationResponseExpenseApplication) SetCurrentStepId(v int64)`

SetCurrentStepId sets CurrentStepId field to given value.


### SetCurrentStepIdNil

`func (o *ExpenseApplicationResponseExpenseApplication) SetCurrentStepIdNil(b bool)`

 SetCurrentStepIdNil sets the value for CurrentStepId to be an explicit nil

### UnsetCurrentStepId
`func (o *ExpenseApplicationResponseExpenseApplication) UnsetCurrentStepId()`

UnsetCurrentStepId ensures that no value is present for CurrentStepId, not even an explicit nil
### GetDealId

`func (o *ExpenseApplicationResponseExpenseApplication) GetDealId() int64`

GetDealId returns the DealId field if non-nil, zero value otherwise.

### GetDealIdOk

`func (o *ExpenseApplicationResponseExpenseApplication) GetDealIdOk() (*int64, bool)`

GetDealIdOk returns a tuple with the DealId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealId

`func (o *ExpenseApplicationResponseExpenseApplication) SetDealId(v int64)`

SetDealId sets DealId field to given value.


### SetDealIdNil

`func (o *ExpenseApplicationResponseExpenseApplication) SetDealIdNil(b bool)`

 SetDealIdNil sets the value for DealId to be an explicit nil

### UnsetDealId
`func (o *ExpenseApplicationResponseExpenseApplication) UnsetDealId()`

UnsetDealId ensures that no value is present for DealId, not even an explicit nil
### GetDealStatus

`func (o *ExpenseApplicationResponseExpenseApplication) GetDealStatus() string`

GetDealStatus returns the DealStatus field if non-nil, zero value otherwise.

### GetDealStatusOk

`func (o *ExpenseApplicationResponseExpenseApplication) GetDealStatusOk() (*string, bool)`

GetDealStatusOk returns a tuple with the DealStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealStatus

`func (o *ExpenseApplicationResponseExpenseApplication) SetDealStatus(v string)`

SetDealStatus sets DealStatus field to given value.


### SetDealStatusNil

`func (o *ExpenseApplicationResponseExpenseApplication) SetDealStatusNil(b bool)`

 SetDealStatusNil sets the value for DealStatus to be an explicit nil

### UnsetDealStatus
`func (o *ExpenseApplicationResponseExpenseApplication) UnsetDealStatus()`

UnsetDealStatus ensures that no value is present for DealStatus, not even an explicit nil
### GetDescription

`func (o *ExpenseApplicationResponseExpenseApplication) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExpenseApplicationResponseExpenseApplication) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExpenseApplicationResponseExpenseApplication) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExpenseApplicationResponseExpenseApplication) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ExpenseApplicationResponseExpenseApplication) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ExpenseApplicationResponseExpenseApplication) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExpenseApplicationLines

`func (o *ExpenseApplicationResponseExpenseApplication) GetExpenseApplicationLines() []ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner`

GetExpenseApplicationLines returns the ExpenseApplicationLines field if non-nil, zero value otherwise.

### GetExpenseApplicationLinesOk

`func (o *ExpenseApplicationResponseExpenseApplication) GetExpenseApplicationLinesOk() (*[]ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner, bool)`

GetExpenseApplicationLinesOk returns a tuple with the ExpenseApplicationLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseApplicationLines

`func (o *ExpenseApplicationResponseExpenseApplication) SetExpenseApplicationLines(v []ExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesInner)`

SetExpenseApplicationLines sets ExpenseApplicationLines field to given value.


### GetId

`func (o *ExpenseApplicationResponseExpenseApplication) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExpenseApplicationResponseExpenseApplication) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExpenseApplicationResponseExpenseApplication) SetId(v int64)`

SetId sets Id field to given value.


### GetIssueDate

`func (o *ExpenseApplicationResponseExpenseApplication) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *ExpenseApplicationResponseExpenseApplication) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *ExpenseApplicationResponseExpenseApplication) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetSectionId

`func (o *ExpenseApplicationResponseExpenseApplication) GetSectionId() int64`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *ExpenseApplicationResponseExpenseApplication) GetSectionIdOk() (*int64, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *ExpenseApplicationResponseExpenseApplication) SetSectionId(v int64)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *ExpenseApplicationResponseExpenseApplication) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### SetSectionIdNil

`func (o *ExpenseApplicationResponseExpenseApplication) SetSectionIdNil(b bool)`

 SetSectionIdNil sets the value for SectionId to be an explicit nil

### UnsetSectionId
`func (o *ExpenseApplicationResponseExpenseApplication) UnsetSectionId()`

UnsetSectionId ensures that no value is present for SectionId, not even an explicit nil
### GetSegment1TagId

`func (o *ExpenseApplicationResponseExpenseApplication) GetSegment1TagId() int64`

GetSegment1TagId returns the Segment1TagId field if non-nil, zero value otherwise.

### GetSegment1TagIdOk

`func (o *ExpenseApplicationResponseExpenseApplication) GetSegment1TagIdOk() (*int64, bool)`

GetSegment1TagIdOk returns a tuple with the Segment1TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagId

`func (o *ExpenseApplicationResponseExpenseApplication) SetSegment1TagId(v int64)`

SetSegment1TagId sets Segment1TagId field to given value.

### HasSegment1TagId

`func (o *ExpenseApplicationResponseExpenseApplication) HasSegment1TagId() bool`

HasSegment1TagId returns a boolean if a field has been set.

### SetSegment1TagIdNil

`func (o *ExpenseApplicationResponseExpenseApplication) SetSegment1TagIdNil(b bool)`

 SetSegment1TagIdNil sets the value for Segment1TagId to be an explicit nil

### UnsetSegment1TagId
`func (o *ExpenseApplicationResponseExpenseApplication) UnsetSegment1TagId()`

UnsetSegment1TagId ensures that no value is present for Segment1TagId, not even an explicit nil
### GetSegment2TagId

`func (o *ExpenseApplicationResponseExpenseApplication) GetSegment2TagId() int64`

GetSegment2TagId returns the Segment2TagId field if non-nil, zero value otherwise.

### GetSegment2TagIdOk

`func (o *ExpenseApplicationResponseExpenseApplication) GetSegment2TagIdOk() (*int64, bool)`

GetSegment2TagIdOk returns a tuple with the Segment2TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagId

`func (o *ExpenseApplicationResponseExpenseApplication) SetSegment2TagId(v int64)`

SetSegment2TagId sets Segment2TagId field to given value.

### HasSegment2TagId

`func (o *ExpenseApplicationResponseExpenseApplication) HasSegment2TagId() bool`

HasSegment2TagId returns a boolean if a field has been set.

### SetSegment2TagIdNil

`func (o *ExpenseApplicationResponseExpenseApplication) SetSegment2TagIdNil(b bool)`

 SetSegment2TagIdNil sets the value for Segment2TagId to be an explicit nil

### UnsetSegment2TagId
`func (o *ExpenseApplicationResponseExpenseApplication) UnsetSegment2TagId()`

UnsetSegment2TagId ensures that no value is present for Segment2TagId, not even an explicit nil
### GetSegment3TagId

`func (o *ExpenseApplicationResponseExpenseApplication) GetSegment3TagId() int64`

GetSegment3TagId returns the Segment3TagId field if non-nil, zero value otherwise.

### GetSegment3TagIdOk

`func (o *ExpenseApplicationResponseExpenseApplication) GetSegment3TagIdOk() (*int64, bool)`

GetSegment3TagIdOk returns a tuple with the Segment3TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagId

`func (o *ExpenseApplicationResponseExpenseApplication) SetSegment3TagId(v int64)`

SetSegment3TagId sets Segment3TagId field to given value.

### HasSegment3TagId

`func (o *ExpenseApplicationResponseExpenseApplication) HasSegment3TagId() bool`

HasSegment3TagId returns a boolean if a field has been set.

### SetSegment3TagIdNil

`func (o *ExpenseApplicationResponseExpenseApplication) SetSegment3TagIdNil(b bool)`

 SetSegment3TagIdNil sets the value for Segment3TagId to be an explicit nil

### UnsetSegment3TagId
`func (o *ExpenseApplicationResponseExpenseApplication) UnsetSegment3TagId()`

UnsetSegment3TagId ensures that no value is present for Segment3TagId, not even an explicit nil
### GetStatus

`func (o *ExpenseApplicationResponseExpenseApplication) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExpenseApplicationResponseExpenseApplication) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExpenseApplicationResponseExpenseApplication) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTagIds

`func (o *ExpenseApplicationResponseExpenseApplication) GetTagIds() []int64`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *ExpenseApplicationResponseExpenseApplication) GetTagIdsOk() (*[]int64, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *ExpenseApplicationResponseExpenseApplication) SetTagIds(v []int64)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *ExpenseApplicationResponseExpenseApplication) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTitle

`func (o *ExpenseApplicationResponseExpenseApplication) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ExpenseApplicationResponseExpenseApplication) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ExpenseApplicationResponseExpenseApplication) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetTotalAmount

`func (o *ExpenseApplicationResponseExpenseApplication) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *ExpenseApplicationResponseExpenseApplication) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *ExpenseApplicationResponseExpenseApplication) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *ExpenseApplicationResponseExpenseApplication) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


