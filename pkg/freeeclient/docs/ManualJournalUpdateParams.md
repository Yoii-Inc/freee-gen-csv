# ManualJournalUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adjustment** | Pointer to **bool** | 決算整理仕訳フラグ（falseまたは未指定の場合: 日常仕訳） | [optional] 
**CompanyId** | **int64** | 事業所ID | 
**Details** | [**[]ManualJournalUpdateParamsDetailsInner**](ManualJournalUpdateParamsDetailsInner.md) |  | 
**IssueDate** | **string** | 発生日 (yyyy-mm-dd) | 
**ReceiptIds** | Pointer to **[]int64** | 証憑ファイルID（ファイルボックスのファイルID）（配列） | [optional] 

## Methods

### NewManualJournalUpdateParams

`func NewManualJournalUpdateParams(companyId int64, details []ManualJournalUpdateParamsDetailsInner, issueDate string, ) *ManualJournalUpdateParams`

NewManualJournalUpdateParams instantiates a new ManualJournalUpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualJournalUpdateParamsWithDefaults

`func NewManualJournalUpdateParamsWithDefaults() *ManualJournalUpdateParams`

NewManualJournalUpdateParamsWithDefaults instantiates a new ManualJournalUpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdjustment

`func (o *ManualJournalUpdateParams) GetAdjustment() bool`

GetAdjustment returns the Adjustment field if non-nil, zero value otherwise.

### GetAdjustmentOk

`func (o *ManualJournalUpdateParams) GetAdjustmentOk() (*bool, bool)`

GetAdjustmentOk returns a tuple with the Adjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustment

`func (o *ManualJournalUpdateParams) SetAdjustment(v bool)`

SetAdjustment sets Adjustment field to given value.

### HasAdjustment

`func (o *ManualJournalUpdateParams) HasAdjustment() bool`

HasAdjustment returns a boolean if a field has been set.

### GetCompanyId

`func (o *ManualJournalUpdateParams) GetCompanyId() int64`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ManualJournalUpdateParams) GetCompanyIdOk() (*int64, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ManualJournalUpdateParams) SetCompanyId(v int64)`

SetCompanyId sets CompanyId field to given value.


### GetDetails

`func (o *ManualJournalUpdateParams) GetDetails() []ManualJournalUpdateParamsDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ManualJournalUpdateParams) GetDetailsOk() (*[]ManualJournalUpdateParamsDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ManualJournalUpdateParams) SetDetails(v []ManualJournalUpdateParamsDetailsInner)`

SetDetails sets Details field to given value.


### GetIssueDate

`func (o *ManualJournalUpdateParams) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *ManualJournalUpdateParams) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *ManualJournalUpdateParams) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetReceiptIds

`func (o *ManualJournalUpdateParams) GetReceiptIds() []int64`

GetReceiptIds returns the ReceiptIds field if non-nil, zero value otherwise.

### GetReceiptIdsOk

`func (o *ManualJournalUpdateParams) GetReceiptIdsOk() (*[]int64, bool)`

GetReceiptIdsOk returns a tuple with the ReceiptIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptIds

`func (o *ManualJournalUpdateParams) SetReceiptIds(v []int64)`

SetReceiptIds sets ReceiptIds field to given value.

### HasReceiptIds

`func (o *ManualJournalUpdateParams) HasReceiptIds() bool`

HasReceiptIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


