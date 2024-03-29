# DealUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int64** | 事業所ID | 
**Details** | [**[]DealUpdateParamsDetailsInner**](DealUpdateParamsDetailsInner.md) |  | 
**DueDate** | Pointer to **string** | 支払期日(yyyy-mm-dd) | [optional] 
**IssueDate** | **string** | 発生日 (yyyy-mm-dd) | 
**PartnerCode** | Pointer to **string** | 取引先コード | [optional] 
**PartnerId** | Pointer to **int64** | 取引先ID | [optional] 
**ReceiptIds** | Pointer to **[]int64** | 証憑ファイルID（ファイルボックスのファイルID）（配列） | [optional] 
**RefNumber** | Pointer to **string** | 管理番号 | [optional] 
**Type** | **string** | 収支区分 (収入: income, 支出: expense) | 

## Methods

### NewDealUpdateParams

`func NewDealUpdateParams(companyId int64, details []DealUpdateParamsDetailsInner, issueDate string, type_ string, ) *DealUpdateParams`

NewDealUpdateParams instantiates a new DealUpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealUpdateParamsWithDefaults

`func NewDealUpdateParamsWithDefaults() *DealUpdateParams`

NewDealUpdateParamsWithDefaults instantiates a new DealUpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *DealUpdateParams) GetCompanyId() int64`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *DealUpdateParams) GetCompanyIdOk() (*int64, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *DealUpdateParams) SetCompanyId(v int64)`

SetCompanyId sets CompanyId field to given value.


### GetDetails

`func (o *DealUpdateParams) GetDetails() []DealUpdateParamsDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *DealUpdateParams) GetDetailsOk() (*[]DealUpdateParamsDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *DealUpdateParams) SetDetails(v []DealUpdateParamsDetailsInner)`

SetDetails sets Details field to given value.


### GetDueDate

`func (o *DealUpdateParams) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *DealUpdateParams) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *DealUpdateParams) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *DealUpdateParams) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetIssueDate

`func (o *DealUpdateParams) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *DealUpdateParams) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *DealUpdateParams) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetPartnerCode

`func (o *DealUpdateParams) GetPartnerCode() string`

GetPartnerCode returns the PartnerCode field if non-nil, zero value otherwise.

### GetPartnerCodeOk

`func (o *DealUpdateParams) GetPartnerCodeOk() (*string, bool)`

GetPartnerCodeOk returns a tuple with the PartnerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCode

`func (o *DealUpdateParams) SetPartnerCode(v string)`

SetPartnerCode sets PartnerCode field to given value.

### HasPartnerCode

`func (o *DealUpdateParams) HasPartnerCode() bool`

HasPartnerCode returns a boolean if a field has been set.

### GetPartnerId

`func (o *DealUpdateParams) GetPartnerId() int64`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *DealUpdateParams) GetPartnerIdOk() (*int64, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *DealUpdateParams) SetPartnerId(v int64)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *DealUpdateParams) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetReceiptIds

`func (o *DealUpdateParams) GetReceiptIds() []int64`

GetReceiptIds returns the ReceiptIds field if non-nil, zero value otherwise.

### GetReceiptIdsOk

`func (o *DealUpdateParams) GetReceiptIdsOk() (*[]int64, bool)`

GetReceiptIdsOk returns a tuple with the ReceiptIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptIds

`func (o *DealUpdateParams) SetReceiptIds(v []int64)`

SetReceiptIds sets ReceiptIds field to given value.

### HasReceiptIds

`func (o *DealUpdateParams) HasReceiptIds() bool`

HasReceiptIds returns a boolean if a field has been set.

### SetReceiptIdsNil

`func (o *DealUpdateParams) SetReceiptIdsNil(b bool)`

 SetReceiptIdsNil sets the value for ReceiptIds to be an explicit nil

### UnsetReceiptIds
`func (o *DealUpdateParams) UnsetReceiptIds()`

UnsetReceiptIds ensures that no value is present for ReceiptIds, not even an explicit nil
### GetRefNumber

`func (o *DealUpdateParams) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *DealUpdateParams) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *DealUpdateParams) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *DealUpdateParams) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetType

`func (o *DealUpdateParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DealUpdateParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DealUpdateParams) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


