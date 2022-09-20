# DealCreateResponseDeal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | 金額 | 
**CompanyId** | **int32** | 事業所ID | 
**Details** | Pointer to [**[]DealDetailsInner**](DealDetailsInner.md) | 取引の明細行 | [optional] 
**DueAmount** | Pointer to **int32** | 支払残額 | [optional] 
**DueDate** | Pointer to **string** | 支払期日 (yyyy-mm-dd) | [optional] 
**Id** | **int32** | 取引ID | 
**IssueDate** | **string** | 発生日 (yyyy-mm-dd) | 
**PartnerCode** | Pointer to **NullableString** | 取引先コード | [optional] 
**PartnerId** | **int32** | 取引先ID | 
**Payments** | Pointer to [**[]DealPaymentsInner**](DealPaymentsInner.md) | 取引の支払行 | [optional] 
**Receipts** | Pointer to [**[]DealCreateResponseDealReceiptsInner**](DealCreateResponseDealReceiptsInner.md) | 証憑ファイル（ファイルボックスのファイル） | [optional] 
**RefNumber** | Pointer to **string** | 管理番号 | [optional] 
**Status** | **string** | 決済状況 (未決済: unsettled, 完了: settled) | 
**Type** | Pointer to **string** | 収支区分 (収入: income, 支出: expense) | [optional] 

## Methods

### NewDealCreateResponseDeal

`func NewDealCreateResponseDeal(amount int64, companyId int32, id int32, issueDate string, partnerId int32, status string, ) *DealCreateResponseDeal`

NewDealCreateResponseDeal instantiates a new DealCreateResponseDeal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealCreateResponseDealWithDefaults

`func NewDealCreateResponseDealWithDefaults() *DealCreateResponseDeal`

NewDealCreateResponseDealWithDefaults instantiates a new DealCreateResponseDeal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DealCreateResponseDeal) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DealCreateResponseDeal) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DealCreateResponseDeal) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCompanyId

`func (o *DealCreateResponseDeal) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *DealCreateResponseDeal) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *DealCreateResponseDeal) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetDetails

`func (o *DealCreateResponseDeal) GetDetails() []DealDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *DealCreateResponseDeal) GetDetailsOk() (*[]DealDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *DealCreateResponseDeal) SetDetails(v []DealDetailsInner)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *DealCreateResponseDeal) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetDueAmount

`func (o *DealCreateResponseDeal) GetDueAmount() int32`

GetDueAmount returns the DueAmount field if non-nil, zero value otherwise.

### GetDueAmountOk

`func (o *DealCreateResponseDeal) GetDueAmountOk() (*int32, bool)`

GetDueAmountOk returns a tuple with the DueAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueAmount

`func (o *DealCreateResponseDeal) SetDueAmount(v int32)`

SetDueAmount sets DueAmount field to given value.

### HasDueAmount

`func (o *DealCreateResponseDeal) HasDueAmount() bool`

HasDueAmount returns a boolean if a field has been set.

### GetDueDate

`func (o *DealCreateResponseDeal) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *DealCreateResponseDeal) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *DealCreateResponseDeal) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *DealCreateResponseDeal) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetId

`func (o *DealCreateResponseDeal) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DealCreateResponseDeal) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DealCreateResponseDeal) SetId(v int32)`

SetId sets Id field to given value.


### GetIssueDate

`func (o *DealCreateResponseDeal) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *DealCreateResponseDeal) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *DealCreateResponseDeal) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetPartnerCode

`func (o *DealCreateResponseDeal) GetPartnerCode() string`

GetPartnerCode returns the PartnerCode field if non-nil, zero value otherwise.

### GetPartnerCodeOk

`func (o *DealCreateResponseDeal) GetPartnerCodeOk() (*string, bool)`

GetPartnerCodeOk returns a tuple with the PartnerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCode

`func (o *DealCreateResponseDeal) SetPartnerCode(v string)`

SetPartnerCode sets PartnerCode field to given value.

### HasPartnerCode

`func (o *DealCreateResponseDeal) HasPartnerCode() bool`

HasPartnerCode returns a boolean if a field has been set.

### SetPartnerCodeNil

`func (o *DealCreateResponseDeal) SetPartnerCodeNil(b bool)`

 SetPartnerCodeNil sets the value for PartnerCode to be an explicit nil

### UnsetPartnerCode
`func (o *DealCreateResponseDeal) UnsetPartnerCode()`

UnsetPartnerCode ensures that no value is present for PartnerCode, not even an explicit nil
### GetPartnerId

`func (o *DealCreateResponseDeal) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *DealCreateResponseDeal) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *DealCreateResponseDeal) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.


### GetPayments

`func (o *DealCreateResponseDeal) GetPayments() []DealPaymentsInner`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *DealCreateResponseDeal) GetPaymentsOk() (*[]DealPaymentsInner, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *DealCreateResponseDeal) SetPayments(v []DealPaymentsInner)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *DealCreateResponseDeal) HasPayments() bool`

HasPayments returns a boolean if a field has been set.

### GetReceipts

`func (o *DealCreateResponseDeal) GetReceipts() []DealCreateResponseDealReceiptsInner`

GetReceipts returns the Receipts field if non-nil, zero value otherwise.

### GetReceiptsOk

`func (o *DealCreateResponseDeal) GetReceiptsOk() (*[]DealCreateResponseDealReceiptsInner, bool)`

GetReceiptsOk returns a tuple with the Receipts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceipts

`func (o *DealCreateResponseDeal) SetReceipts(v []DealCreateResponseDealReceiptsInner)`

SetReceipts sets Receipts field to given value.

### HasReceipts

`func (o *DealCreateResponseDeal) HasReceipts() bool`

HasReceipts returns a boolean if a field has been set.

### GetRefNumber

`func (o *DealCreateResponseDeal) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *DealCreateResponseDeal) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *DealCreateResponseDeal) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *DealCreateResponseDeal) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetStatus

`func (o *DealCreateResponseDeal) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DealCreateResponseDeal) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DealCreateResponseDeal) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *DealCreateResponseDeal) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DealCreateResponseDeal) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DealCreateResponseDeal) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DealCreateResponseDeal) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


