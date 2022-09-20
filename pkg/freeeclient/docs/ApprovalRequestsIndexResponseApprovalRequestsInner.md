# ApprovalRequestsIndexResponseApprovalRequestsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicantId** | **int64** | 申請者のユーザーID | 
**ApplicationDate** | **string** | 申請日 (yyyy-mm-dd) | 
**ApplicationNumber** | **string** | 申請No. | 
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

### NewApprovalRequestsIndexResponseApprovalRequestsInner

`func NewApprovalRequestsIndexResponseApprovalRequestsInner(applicantId int64, applicationDate string, applicationNumber string, companyId int64, currentRound int64, currentStepId NullableInt64, dealId NullableInt64, dealStatus NullableString, formId int64, id int64, manualJournalId NullableInt64, requestItems []ApprovalRequestResponseApprovalRequestRequestItemsInner, status string, title string, ) *ApprovalRequestsIndexResponseApprovalRequestsInner`

NewApprovalRequestsIndexResponseApprovalRequestsInner instantiates a new ApprovalRequestsIndexResponseApprovalRequestsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalRequestsIndexResponseApprovalRequestsInnerWithDefaults

`func NewApprovalRequestsIndexResponseApprovalRequestsInnerWithDefaults() *ApprovalRequestsIndexResponseApprovalRequestsInner`

NewApprovalRequestsIndexResponseApprovalRequestsInnerWithDefaults instantiates a new ApprovalRequestsIndexResponseApprovalRequestsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicantId

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) GetApplicantId() int64`

GetApplicantId returns the ApplicantId field if non-nil, zero value otherwise.

### GetApplicantIdOk

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) GetApplicantIdOk() (*int64, bool)`

GetApplicantIdOk returns a tuple with the ApplicantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantId

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) SetApplicantId(v int64)`

SetApplicantId sets ApplicantId field to given value.


### GetApplicationDate

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) GetApplicationDate() string`

GetApplicationDate returns the ApplicationDate field if non-nil, zero value otherwise.

### GetApplicationDateOk

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) GetApplicationDateOk() (*string, bool)`

GetApplicationDateOk returns a tuple with the ApplicationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationDate

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) SetApplicationDate(v string)`

SetApplicationDate sets ApplicationDate field to given value.


### GetApplicationNumber

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) GetApplicationNumber() string`

GetApplicationNumber returns the ApplicationNumber field if non-nil, zero value otherwise.

### GetApplicationNumberOk

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) GetApplicationNumberOk() (*string, bool)`

GetApplicationNumberOk returns a tuple with the ApplicationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationNumber

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) SetApplicationNumber(v string)`

SetApplicationNumber sets ApplicationNumber field to given value.


### GetCompanyId

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) GetCompanyId() int64`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) GetCompanyIdOk() (*int64, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) SetCompanyId(v int64)`

SetCompanyId sets CompanyId field to given value.


### GetCurrentRound

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) GetCurrentRound() int64`

GetCurrentRound returns the CurrentRound field if non-nil, zero value otherwise.

### GetCurrentRoundOk

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) GetCurrentRoundOk() (*int64, bool)`

GetCurrentRoundOk returns a tuple with the CurrentRound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRound

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) SetCurrentRound(v int64)`

SetCurrentRound sets CurrentRound field to given value.


### GetCurrentStepId

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) GetCurrentStepId() int64`

GetCurrentStepId returns the CurrentStepId field if non-nil, zero value otherwise.

### GetCurrentStepIdOk

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) GetCurrentStepIdOk() (*int64, bool)`

GetCurrentStepIdOk returns a tuple with the CurrentStepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStepId

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) SetCurrentStepId(v int64)`

SetCurrentStepId sets CurrentStepId field to given value.


### SetCurrentStepIdNil

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) SetCurrentStepIdNil(b bool)`

 SetCurrentStepIdNil sets the value for CurrentStepId to be an explicit nil

### UnsetCurrentStepId
`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) UnsetCurrentStepId()`

UnsetCurrentStepId ensures that no value is present for CurrentStepId, not even an explicit nil
### GetDealId

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) GetDealId() int64`

GetDealId returns the DealId field if non-nil, zero value otherwise.

### GetDealIdOk

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) GetDealIdOk() (*int64, bool)`

GetDealIdOk returns a tuple with the DealId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealId

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) SetDealId(v int64)`

SetDealId sets DealId field to given value.


### SetDealIdNil

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) SetDealIdNil(b bool)`

 SetDealIdNil sets the value for DealId to be an explicit nil

### UnsetDealId
`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) UnsetDealId()`

UnsetDealId ensures that no value is present for DealId, not even an explicit nil
### GetDealStatus

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) GetDealStatus() string`

GetDealStatus returns the DealStatus field if non-nil, zero value otherwise.

### GetDealStatusOk

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) GetDealStatusOk() (*string, bool)`

GetDealStatusOk returns a tuple with the DealStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealStatus

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) SetDealStatus(v string)`

SetDealStatus sets DealStatus field to given value.


### SetDealStatusNil

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) SetDealStatusNil(b bool)`

 SetDealStatusNil sets the value for DealStatus to be an explicit nil

### UnsetDealStatus
`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) UnsetDealStatus()`

UnsetDealStatus ensures that no value is present for DealStatus, not even an explicit nil
### GetFormId

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) GetFormId() int64`

GetFormId returns the FormId field if non-nil, zero value otherwise.

### GetFormIdOk

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) GetFormIdOk() (*int64, bool)`

GetFormIdOk returns a tuple with the FormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormId

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) SetFormId(v int64)`

SetFormId sets FormId field to given value.


### GetId

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) SetId(v int64)`

SetId sets Id field to given value.


### GetManualJournalId

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) GetManualJournalId() int64`

GetManualJournalId returns the ManualJournalId field if non-nil, zero value otherwise.

### GetManualJournalIdOk

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) GetManualJournalIdOk() (*int64, bool)`

GetManualJournalIdOk returns a tuple with the ManualJournalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualJournalId

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) SetManualJournalId(v int64)`

SetManualJournalId sets ManualJournalId field to given value.


### SetManualJournalIdNil

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) SetManualJournalIdNil(b bool)`

 SetManualJournalIdNil sets the value for ManualJournalId to be an explicit nil

### UnsetManualJournalId
`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) UnsetManualJournalId()`

UnsetManualJournalId ensures that no value is present for ManualJournalId, not even an explicit nil
### GetRequestItems

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) GetRequestItems() []ApprovalRequestResponseApprovalRequestRequestItemsInner`

GetRequestItems returns the RequestItems field if non-nil, zero value otherwise.

### GetRequestItemsOk

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) GetRequestItemsOk() (*[]ApprovalRequestResponseApprovalRequestRequestItemsInner, bool)`

GetRequestItemsOk returns a tuple with the RequestItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestItems

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) SetRequestItems(v []ApprovalRequestResponseApprovalRequestRequestItemsInner)`

SetRequestItems sets RequestItems field to given value.


### GetStatus

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApprovalRequestsIndexResponseApprovalRequestsInner) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


