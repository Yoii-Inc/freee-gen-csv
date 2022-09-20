# ExpenseApplicationsIndexResponseExpenseApplicationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicantId** | **int64** | 申請者のユーザーID | 
**ApplicationNumber** | **string** | 申請No. | 
**CompanyId** | **int64** | 事業所ID | 
**CurrentRound** | Pointer to **int64** | 現在のround。差し戻し等により申請がstepの最初からやり直しになるとroundの値が増えます。 | [optional] 
**CurrentStepId** | Pointer to **NullableInt64** | 現在承認ステップID | [optional] 
**DealId** | **NullableInt64** | 取引ID (申請ステータス:statusがapprovedで、取引が存在する時のみdeal_idが表示されます) | 
**DealStatus** | **NullableString** | 取引ステータス (申請ステータス:statusがapprovedで、取引が存在する時のみdeal_statusが表示されます settled:精算済み, unsettled:清算待ち) | 
**Description** | Pointer to **string** | 備考 | [optional] 
**ExpenseApplicationLines** | [**[]ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner**](ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner.md) | 経費申請の項目行一覧（配列） | 
**Id** | **int64** | 経費申請ID | 
**IssueDate** | **string** | 申請日 (yyyy-mm-dd) | 
**SectionId** | Pointer to **NullableInt64** | 部門ID | [optional] 
**Segment1TagId** | Pointer to **NullableInt64** | セグメント１ID | [optional] 
**Segment2TagId** | Pointer to **NullableInt64** | セグメント２ID | [optional] 
**Segment3TagId** | Pointer to **NullableInt64** | セグメント３ID | [optional] 
**Status** | **string** | 申請ステータス(draft:下書き, in_progress:申請中, approved:承認済, rejected:却下, feedback:差戻し) | 
**TagIds** | Pointer to **[]int64** | メモタグID | [optional] 
**Title** | **string** | 申請タイトル | 
**TotalAmount** | Pointer to **int64** | 合計金額 | [optional] 

## Methods

### NewExpenseApplicationsIndexResponseExpenseApplicationsInner

`func NewExpenseApplicationsIndexResponseExpenseApplicationsInner(applicantId int64, applicationNumber string, companyId int64, dealId NullableInt64, dealStatus NullableString, expenseApplicationLines []ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner, id int64, issueDate string, status string, title string, ) *ExpenseApplicationsIndexResponseExpenseApplicationsInner`

NewExpenseApplicationsIndexResponseExpenseApplicationsInner instantiates a new ExpenseApplicationsIndexResponseExpenseApplicationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseApplicationsIndexResponseExpenseApplicationsInnerWithDefaults

`func NewExpenseApplicationsIndexResponseExpenseApplicationsInnerWithDefaults() *ExpenseApplicationsIndexResponseExpenseApplicationsInner`

NewExpenseApplicationsIndexResponseExpenseApplicationsInnerWithDefaults instantiates a new ExpenseApplicationsIndexResponseExpenseApplicationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicantId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetApplicantId() int64`

GetApplicantId returns the ApplicantId field if non-nil, zero value otherwise.

### GetApplicantIdOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetApplicantIdOk() (*int64, bool)`

GetApplicantIdOk returns a tuple with the ApplicantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) SetApplicantId(v int64)`

SetApplicantId sets ApplicantId field to given value.


### GetApplicationNumber

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetApplicationNumber() string`

GetApplicationNumber returns the ApplicationNumber field if non-nil, zero value otherwise.

### GetApplicationNumberOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetApplicationNumberOk() (*string, bool)`

GetApplicationNumberOk returns a tuple with the ApplicationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationNumber

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) SetApplicationNumber(v string)`

SetApplicationNumber sets ApplicationNumber field to given value.


### GetCompanyId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetCompanyId() int64`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetCompanyIdOk() (*int64, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) SetCompanyId(v int64)`

SetCompanyId sets CompanyId field to given value.


### GetCurrentRound

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetCurrentRound() int64`

GetCurrentRound returns the CurrentRound field if non-nil, zero value otherwise.

### GetCurrentRoundOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetCurrentRoundOk() (*int64, bool)`

GetCurrentRoundOk returns a tuple with the CurrentRound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRound

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) SetCurrentRound(v int64)`

SetCurrentRound sets CurrentRound field to given value.

### HasCurrentRound

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) HasCurrentRound() bool`

HasCurrentRound returns a boolean if a field has been set.

### GetCurrentStepId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetCurrentStepId() int64`

GetCurrentStepId returns the CurrentStepId field if non-nil, zero value otherwise.

### GetCurrentStepIdOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetCurrentStepIdOk() (*int64, bool)`

GetCurrentStepIdOk returns a tuple with the CurrentStepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStepId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) SetCurrentStepId(v int64)`

SetCurrentStepId sets CurrentStepId field to given value.

### HasCurrentStepId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) HasCurrentStepId() bool`

HasCurrentStepId returns a boolean if a field has been set.

### SetCurrentStepIdNil

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) SetCurrentStepIdNil(b bool)`

 SetCurrentStepIdNil sets the value for CurrentStepId to be an explicit nil

### UnsetCurrentStepId
`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) UnsetCurrentStepId()`

UnsetCurrentStepId ensures that no value is present for CurrentStepId, not even an explicit nil
### GetDealId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetDealId() int64`

GetDealId returns the DealId field if non-nil, zero value otherwise.

### GetDealIdOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetDealIdOk() (*int64, bool)`

GetDealIdOk returns a tuple with the DealId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) SetDealId(v int64)`

SetDealId sets DealId field to given value.


### SetDealIdNil

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) SetDealIdNil(b bool)`

 SetDealIdNil sets the value for DealId to be an explicit nil

### UnsetDealId
`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) UnsetDealId()`

UnsetDealId ensures that no value is present for DealId, not even an explicit nil
### GetDealStatus

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetDealStatus() string`

GetDealStatus returns the DealStatus field if non-nil, zero value otherwise.

### GetDealStatusOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetDealStatusOk() (*string, bool)`

GetDealStatusOk returns a tuple with the DealStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealStatus

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) SetDealStatus(v string)`

SetDealStatus sets DealStatus field to given value.


### SetDealStatusNil

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) SetDealStatusNil(b bool)`

 SetDealStatusNil sets the value for DealStatus to be an explicit nil

### UnsetDealStatus
`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) UnsetDealStatus()`

UnsetDealStatus ensures that no value is present for DealStatus, not even an explicit nil
### GetDescription

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpenseApplicationLines

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetExpenseApplicationLines() []ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner`

GetExpenseApplicationLines returns the ExpenseApplicationLines field if non-nil, zero value otherwise.

### GetExpenseApplicationLinesOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetExpenseApplicationLinesOk() (*[]ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner, bool)`

GetExpenseApplicationLinesOk returns a tuple with the ExpenseApplicationLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseApplicationLines

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) SetExpenseApplicationLines(v []ExpenseApplicationsIndexResponseExpenseApplicationsInnerExpenseApplicationLinesInner)`

SetExpenseApplicationLines sets ExpenseApplicationLines field to given value.


### GetId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) SetId(v int64)`

SetId sets Id field to given value.


### GetIssueDate

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetSectionId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetSectionId() int64`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetSectionIdOk() (*int64, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) SetSectionId(v int64)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### SetSectionIdNil

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) SetSectionIdNil(b bool)`

 SetSectionIdNil sets the value for SectionId to be an explicit nil

### UnsetSectionId
`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) UnsetSectionId()`

UnsetSectionId ensures that no value is present for SectionId, not even an explicit nil
### GetSegment1TagId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetSegment1TagId() int64`

GetSegment1TagId returns the Segment1TagId field if non-nil, zero value otherwise.

### GetSegment1TagIdOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetSegment1TagIdOk() (*int64, bool)`

GetSegment1TagIdOk returns a tuple with the Segment1TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) SetSegment1TagId(v int64)`

SetSegment1TagId sets Segment1TagId field to given value.

### HasSegment1TagId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) HasSegment1TagId() bool`

HasSegment1TagId returns a boolean if a field has been set.

### SetSegment1TagIdNil

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) SetSegment1TagIdNil(b bool)`

 SetSegment1TagIdNil sets the value for Segment1TagId to be an explicit nil

### UnsetSegment1TagId
`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) UnsetSegment1TagId()`

UnsetSegment1TagId ensures that no value is present for Segment1TagId, not even an explicit nil
### GetSegment2TagId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetSegment2TagId() int64`

GetSegment2TagId returns the Segment2TagId field if non-nil, zero value otherwise.

### GetSegment2TagIdOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetSegment2TagIdOk() (*int64, bool)`

GetSegment2TagIdOk returns a tuple with the Segment2TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) SetSegment2TagId(v int64)`

SetSegment2TagId sets Segment2TagId field to given value.

### HasSegment2TagId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) HasSegment2TagId() bool`

HasSegment2TagId returns a boolean if a field has been set.

### SetSegment2TagIdNil

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) SetSegment2TagIdNil(b bool)`

 SetSegment2TagIdNil sets the value for Segment2TagId to be an explicit nil

### UnsetSegment2TagId
`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) UnsetSegment2TagId()`

UnsetSegment2TagId ensures that no value is present for Segment2TagId, not even an explicit nil
### GetSegment3TagId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetSegment3TagId() int64`

GetSegment3TagId returns the Segment3TagId field if non-nil, zero value otherwise.

### GetSegment3TagIdOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetSegment3TagIdOk() (*int64, bool)`

GetSegment3TagIdOk returns a tuple with the Segment3TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) SetSegment3TagId(v int64)`

SetSegment3TagId sets Segment3TagId field to given value.

### HasSegment3TagId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) HasSegment3TagId() bool`

HasSegment3TagId returns a boolean if a field has been set.

### SetSegment3TagIdNil

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) SetSegment3TagIdNil(b bool)`

 SetSegment3TagIdNil sets the value for Segment3TagId to be an explicit nil

### UnsetSegment3TagId
`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) UnsetSegment3TagId()`

UnsetSegment3TagId ensures that no value is present for Segment3TagId, not even an explicit nil
### GetStatus

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTagIds

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetTagIds() []int64`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetTagIdsOk() (*[]int64, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) SetTagIds(v []int64)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTitle

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetTotalAmount

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationsInner) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


