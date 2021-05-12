# ApprovalRequestsIndexResponseApprovalRequests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicantId** | **int32** | 申請者のユーザーID | 
**ApplicationDate** | **string** | 申請日 (yyyy-mm-dd) | 
**ApplicationNumber** | **string** | 申請No. | 
**CompanyId** | **int32** | 事業所ID | 
**CurrentRound** | **int32** | 現在のround。差し戻し等により申請がstepの最初からやり直しになるとroundの値が増えます。 | 
**CurrentStepId** | **NullableInt32** | 現在承認ステップID | 
**DealId** | **NullableInt32** | 取引ID (申請ステータス:statusがapprovedで、取引が存在する時のみdeal_idが表示されます) | 
**DealStatus** | **NullableString** | 取引ステータス (申請ステータス:statusがapprovedで、取引が存在する時のみdeal_statusが表示されます settled:精算済み, unsettled:清算待ち) | 
**FormId** | **int32** | 申請フォームID | 
**Id** | **int32** | 各種申請ID | 
**RequestItems** | [**[]ApprovalRequestResponseApprovalRequestRequestItems**](ApprovalRequestResponseApprovalRequestRequestItems.md) | 各種申請の項目一覧（配列） | 
**Status** | **string** | 申請ステータス(draft:下書き, in_progress:申請中, approved:承認済, rejected:却下, feedback:差戻し) | 
**Title** | **string** | 申請タイトル | 

## Methods

### NewApprovalRequestsIndexResponseApprovalRequests

`func NewApprovalRequestsIndexResponseApprovalRequests(applicantId int32, applicationDate string, applicationNumber string, companyId int32, currentRound int32, currentStepId NullableInt32, dealId NullableInt32, dealStatus NullableString, formId int32, id int32, requestItems []ApprovalRequestResponseApprovalRequestRequestItems, status string, title string, ) *ApprovalRequestsIndexResponseApprovalRequests`

NewApprovalRequestsIndexResponseApprovalRequests instantiates a new ApprovalRequestsIndexResponseApprovalRequests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalRequestsIndexResponseApprovalRequestsWithDefaults

`func NewApprovalRequestsIndexResponseApprovalRequestsWithDefaults() *ApprovalRequestsIndexResponseApprovalRequests`

NewApprovalRequestsIndexResponseApprovalRequestsWithDefaults instantiates a new ApprovalRequestsIndexResponseApprovalRequests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicantId

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetApplicantId() int32`

GetApplicantId returns the ApplicantId field if non-nil, zero value otherwise.

### GetApplicantIdOk

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetApplicantIdOk() (*int32, bool)`

GetApplicantIdOk returns a tuple with the ApplicantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantId

`func (o *ApprovalRequestsIndexResponseApprovalRequests) SetApplicantId(v int32)`

SetApplicantId sets ApplicantId field to given value.


### GetApplicationDate

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetApplicationDate() string`

GetApplicationDate returns the ApplicationDate field if non-nil, zero value otherwise.

### GetApplicationDateOk

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetApplicationDateOk() (*string, bool)`

GetApplicationDateOk returns a tuple with the ApplicationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationDate

`func (o *ApprovalRequestsIndexResponseApprovalRequests) SetApplicationDate(v string)`

SetApplicationDate sets ApplicationDate field to given value.


### GetApplicationNumber

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetApplicationNumber() string`

GetApplicationNumber returns the ApplicationNumber field if non-nil, zero value otherwise.

### GetApplicationNumberOk

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetApplicationNumberOk() (*string, bool)`

GetApplicationNumberOk returns a tuple with the ApplicationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationNumber

`func (o *ApprovalRequestsIndexResponseApprovalRequests) SetApplicationNumber(v string)`

SetApplicationNumber sets ApplicationNumber field to given value.


### GetCompanyId

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ApprovalRequestsIndexResponseApprovalRequests) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetCurrentRound

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetCurrentRound() int32`

GetCurrentRound returns the CurrentRound field if non-nil, zero value otherwise.

### GetCurrentRoundOk

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetCurrentRoundOk() (*int32, bool)`

GetCurrentRoundOk returns a tuple with the CurrentRound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRound

`func (o *ApprovalRequestsIndexResponseApprovalRequests) SetCurrentRound(v int32)`

SetCurrentRound sets CurrentRound field to given value.


### GetCurrentStepId

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetCurrentStepId() int32`

GetCurrentStepId returns the CurrentStepId field if non-nil, zero value otherwise.

### GetCurrentStepIdOk

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetCurrentStepIdOk() (*int32, bool)`

GetCurrentStepIdOk returns a tuple with the CurrentStepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStepId

`func (o *ApprovalRequestsIndexResponseApprovalRequests) SetCurrentStepId(v int32)`

SetCurrentStepId sets CurrentStepId field to given value.


### SetCurrentStepIdNil

`func (o *ApprovalRequestsIndexResponseApprovalRequests) SetCurrentStepIdNil(b bool)`

 SetCurrentStepIdNil sets the value for CurrentStepId to be an explicit nil

### UnsetCurrentStepId
`func (o *ApprovalRequestsIndexResponseApprovalRequests) UnsetCurrentStepId()`

UnsetCurrentStepId ensures that no value is present for CurrentStepId, not even an explicit nil
### GetDealId

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetDealId() int32`

GetDealId returns the DealId field if non-nil, zero value otherwise.

### GetDealIdOk

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetDealIdOk() (*int32, bool)`

GetDealIdOk returns a tuple with the DealId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealId

`func (o *ApprovalRequestsIndexResponseApprovalRequests) SetDealId(v int32)`

SetDealId sets DealId field to given value.


### SetDealIdNil

`func (o *ApprovalRequestsIndexResponseApprovalRequests) SetDealIdNil(b bool)`

 SetDealIdNil sets the value for DealId to be an explicit nil

### UnsetDealId
`func (o *ApprovalRequestsIndexResponseApprovalRequests) UnsetDealId()`

UnsetDealId ensures that no value is present for DealId, not even an explicit nil
### GetDealStatus

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetDealStatus() string`

GetDealStatus returns the DealStatus field if non-nil, zero value otherwise.

### GetDealStatusOk

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetDealStatusOk() (*string, bool)`

GetDealStatusOk returns a tuple with the DealStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealStatus

`func (o *ApprovalRequestsIndexResponseApprovalRequests) SetDealStatus(v string)`

SetDealStatus sets DealStatus field to given value.


### SetDealStatusNil

`func (o *ApprovalRequestsIndexResponseApprovalRequests) SetDealStatusNil(b bool)`

 SetDealStatusNil sets the value for DealStatus to be an explicit nil

### UnsetDealStatus
`func (o *ApprovalRequestsIndexResponseApprovalRequests) UnsetDealStatus()`

UnsetDealStatus ensures that no value is present for DealStatus, not even an explicit nil
### GetFormId

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetFormId() int32`

GetFormId returns the FormId field if non-nil, zero value otherwise.

### GetFormIdOk

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetFormIdOk() (*int32, bool)`

GetFormIdOk returns a tuple with the FormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormId

`func (o *ApprovalRequestsIndexResponseApprovalRequests) SetFormId(v int32)`

SetFormId sets FormId field to given value.


### GetId

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalRequestsIndexResponseApprovalRequests) SetId(v int32)`

SetId sets Id field to given value.


### GetRequestItems

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetRequestItems() []ApprovalRequestResponseApprovalRequestRequestItems`

GetRequestItems returns the RequestItems field if non-nil, zero value otherwise.

### GetRequestItemsOk

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetRequestItemsOk() (*[]ApprovalRequestResponseApprovalRequestRequestItems, bool)`

GetRequestItemsOk returns a tuple with the RequestItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestItems

`func (o *ApprovalRequestsIndexResponseApprovalRequests) SetRequestItems(v []ApprovalRequestResponseApprovalRequestRequestItems)`

SetRequestItems sets RequestItems field to given value.


### GetStatus

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApprovalRequestsIndexResponseApprovalRequests) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApprovalRequestsIndexResponseApprovalRequests) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


