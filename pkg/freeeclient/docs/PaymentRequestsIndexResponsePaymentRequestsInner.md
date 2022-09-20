# PaymentRequestsIndexResponsePaymentRequestsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicantId** | **int32** | 申請者のユーザーID | 
**ApplicationDate** | **string** | 申請日 (yyyy-mm-dd) | 
**ApplicationNumber** | **string** | 申請No. | 
**Approvers** | [**[]ApprovalRequestResponseApprovalRequestApproversInner**](ApprovalRequestResponseApprovalRequestApproversInner.md) | 承認者（配列）   承認ステップのresource_typeがunspecified (指定なし)の場合はapproversはレスポンスに含まれません。   しかし、resource_typeがunspecifiedの承認ステップにおいて誰かが承認・却下・差し戻しのいずれかのアクションを取った後は、   approversはレスポンスに含まれるようになります。   その場合approversにはアクションを行ったステップのIDとアクションを行ったユーザーのIDが含まれます。 | 
**CompanyId** | **int32** | 事業所ID | 
**CurrentRound** | **int32** | 現在のround。差し戻し等により申請がstepの最初からやり直しになるとroundの値が増えます。 | 
**CurrentStepId** | **NullableInt32** | 現在承認ステップID | 
**DealId** | Pointer to **NullableInt32** | 取引ID (申請ステータス:statusがapprovedで、取引が存在する時のみdeal_idが表示されます) | [optional] 
**DealStatus** | Pointer to **NullableString** | 取引ステータス (申請ステータス:statusがapprovedで、取引が存在する時のみdeal_statusが表示されます settled:支払済み, unsettled:支払待ち) | [optional] 
**DocumentCode** | **string** | 請求書番号 | 
**Id** | **int32** | 支払依頼ID | 
**IssueDate** | **string** | 発生日 (yyyy-mm-dd) | 
**PartnerCode** | **NullableString** | 取引先コード | 
**PartnerId** | **NullableInt32** | 取引先ID | 
**PartnerName** | **NullableString** | 取引先名 | 
**PaymentDate** | **NullableString** | 支払期限 (yyyy-mm-dd) | 
**PaymentMethod** | **string** | 支払方法(none: 指定なし, domestic_bank_transfer: 国内振込, abroad_bank_transfer: 国外振込, account_transfer: 口座振替, credit_card: クレジットカード) | 
**Status** | **string** | 申請ステータス(draft:下書き, in_progress:申請中, approved:承認済, rejected:却下, feedback:差戻し) | 
**Title** | **string** | 申請タイトル | 
**TotalAmount** | **int64** | 合計金額 | 

## Methods

### NewPaymentRequestsIndexResponsePaymentRequestsInner

`func NewPaymentRequestsIndexResponsePaymentRequestsInner(applicantId int32, applicationDate string, applicationNumber string, approvers []ApprovalRequestResponseApprovalRequestApproversInner, companyId int32, currentRound int32, currentStepId NullableInt32, documentCode string, id int32, issueDate string, partnerCode NullableString, partnerId NullableInt32, partnerName NullableString, paymentDate NullableString, paymentMethod string, status string, title string, totalAmount int64, ) *PaymentRequestsIndexResponsePaymentRequestsInner`

NewPaymentRequestsIndexResponsePaymentRequestsInner instantiates a new PaymentRequestsIndexResponsePaymentRequestsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRequestsIndexResponsePaymentRequestsInnerWithDefaults

`func NewPaymentRequestsIndexResponsePaymentRequestsInnerWithDefaults() *PaymentRequestsIndexResponsePaymentRequestsInner`

NewPaymentRequestsIndexResponsePaymentRequestsInnerWithDefaults instantiates a new PaymentRequestsIndexResponsePaymentRequestsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicantId

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetApplicantId() int32`

GetApplicantId returns the ApplicantId field if non-nil, zero value otherwise.

### GetApplicantIdOk

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetApplicantIdOk() (*int32, bool)`

GetApplicantIdOk returns a tuple with the ApplicantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantId

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetApplicantId(v int32)`

SetApplicantId sets ApplicantId field to given value.


### GetApplicationDate

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetApplicationDate() string`

GetApplicationDate returns the ApplicationDate field if non-nil, zero value otherwise.

### GetApplicationDateOk

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetApplicationDateOk() (*string, bool)`

GetApplicationDateOk returns a tuple with the ApplicationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationDate

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetApplicationDate(v string)`

SetApplicationDate sets ApplicationDate field to given value.


### GetApplicationNumber

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetApplicationNumber() string`

GetApplicationNumber returns the ApplicationNumber field if non-nil, zero value otherwise.

### GetApplicationNumberOk

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetApplicationNumberOk() (*string, bool)`

GetApplicationNumberOk returns a tuple with the ApplicationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationNumber

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetApplicationNumber(v string)`

SetApplicationNumber sets ApplicationNumber field to given value.


### GetApprovers

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetApprovers() []ApprovalRequestResponseApprovalRequestApproversInner`

GetApprovers returns the Approvers field if non-nil, zero value otherwise.

### GetApproversOk

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetApproversOk() (*[]ApprovalRequestResponseApprovalRequestApproversInner, bool)`

GetApproversOk returns a tuple with the Approvers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovers

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetApprovers(v []ApprovalRequestResponseApprovalRequestApproversInner)`

SetApprovers sets Approvers field to given value.


### GetCompanyId

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetCurrentRound

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetCurrentRound() int32`

GetCurrentRound returns the CurrentRound field if non-nil, zero value otherwise.

### GetCurrentRoundOk

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetCurrentRoundOk() (*int32, bool)`

GetCurrentRoundOk returns a tuple with the CurrentRound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRound

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetCurrentRound(v int32)`

SetCurrentRound sets CurrentRound field to given value.


### GetCurrentStepId

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetCurrentStepId() int32`

GetCurrentStepId returns the CurrentStepId field if non-nil, zero value otherwise.

### GetCurrentStepIdOk

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetCurrentStepIdOk() (*int32, bool)`

GetCurrentStepIdOk returns a tuple with the CurrentStepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStepId

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetCurrentStepId(v int32)`

SetCurrentStepId sets CurrentStepId field to given value.


### SetCurrentStepIdNil

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetCurrentStepIdNil(b bool)`

 SetCurrentStepIdNil sets the value for CurrentStepId to be an explicit nil

### UnsetCurrentStepId
`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) UnsetCurrentStepId()`

UnsetCurrentStepId ensures that no value is present for CurrentStepId, not even an explicit nil
### GetDealId

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetDealId() int32`

GetDealId returns the DealId field if non-nil, zero value otherwise.

### GetDealIdOk

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetDealIdOk() (*int32, bool)`

GetDealIdOk returns a tuple with the DealId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealId

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetDealId(v int32)`

SetDealId sets DealId field to given value.

### HasDealId

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) HasDealId() bool`

HasDealId returns a boolean if a field has been set.

### SetDealIdNil

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetDealIdNil(b bool)`

 SetDealIdNil sets the value for DealId to be an explicit nil

### UnsetDealId
`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) UnsetDealId()`

UnsetDealId ensures that no value is present for DealId, not even an explicit nil
### GetDealStatus

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetDealStatus() string`

GetDealStatus returns the DealStatus field if non-nil, zero value otherwise.

### GetDealStatusOk

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetDealStatusOk() (*string, bool)`

GetDealStatusOk returns a tuple with the DealStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealStatus

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetDealStatus(v string)`

SetDealStatus sets DealStatus field to given value.

### HasDealStatus

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) HasDealStatus() bool`

HasDealStatus returns a boolean if a field has been set.

### SetDealStatusNil

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetDealStatusNil(b bool)`

 SetDealStatusNil sets the value for DealStatus to be an explicit nil

### UnsetDealStatus
`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) UnsetDealStatus()`

UnsetDealStatus ensures that no value is present for DealStatus, not even an explicit nil
### GetDocumentCode

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetDocumentCode() string`

GetDocumentCode returns the DocumentCode field if non-nil, zero value otherwise.

### GetDocumentCodeOk

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetDocumentCodeOk() (*string, bool)`

GetDocumentCodeOk returns a tuple with the DocumentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentCode

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetDocumentCode(v string)`

SetDocumentCode sets DocumentCode field to given value.


### GetId

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetId(v int32)`

SetId sets Id field to given value.


### GetIssueDate

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetPartnerCode

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetPartnerCode() string`

GetPartnerCode returns the PartnerCode field if non-nil, zero value otherwise.

### GetPartnerCodeOk

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetPartnerCodeOk() (*string, bool)`

GetPartnerCodeOk returns a tuple with the PartnerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCode

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetPartnerCode(v string)`

SetPartnerCode sets PartnerCode field to given value.


### SetPartnerCodeNil

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetPartnerCodeNil(b bool)`

 SetPartnerCodeNil sets the value for PartnerCode to be an explicit nil

### UnsetPartnerCode
`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) UnsetPartnerCode()`

UnsetPartnerCode ensures that no value is present for PartnerCode, not even an explicit nil
### GetPartnerId

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.


### SetPartnerIdNil

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetPartnerIdNil(b bool)`

 SetPartnerIdNil sets the value for PartnerId to be an explicit nil

### UnsetPartnerId
`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) UnsetPartnerId()`

UnsetPartnerId ensures that no value is present for PartnerId, not even an explicit nil
### GetPartnerName

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetPartnerName() string`

GetPartnerName returns the PartnerName field if non-nil, zero value otherwise.

### GetPartnerNameOk

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetPartnerNameOk() (*string, bool)`

GetPartnerNameOk returns a tuple with the PartnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerName

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetPartnerName(v string)`

SetPartnerName sets PartnerName field to given value.


### SetPartnerNameNil

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetPartnerNameNil(b bool)`

 SetPartnerNameNil sets the value for PartnerName to be an explicit nil

### UnsetPartnerName
`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) UnsetPartnerName()`

UnsetPartnerName ensures that no value is present for PartnerName, not even an explicit nil
### GetPaymentDate

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetPaymentDate() string`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetPaymentDateOk() (*string, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetPaymentDate(v string)`

SetPaymentDate sets PaymentDate field to given value.


### SetPaymentDateNil

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetPaymentDateNil(b bool)`

 SetPaymentDateNil sets the value for PaymentDate to be an explicit nil

### UnsetPaymentDate
`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) UnsetPaymentDate()`

UnsetPaymentDate ensures that no value is present for PaymentDate, not even an explicit nil
### GetPaymentMethod

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetStatus

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetTotalAmount

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *PaymentRequestsIndexResponsePaymentRequestsInner) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


