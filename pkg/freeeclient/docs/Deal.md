# Deal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | 金額 | 
**CompanyId** | **int64** | 事業所ID | 
**Details** | Pointer to [**[]DealDetailsInner**](DealDetailsInner.md) | 取引の明細行 | [optional] 
**DueAmount** | Pointer to **int64** | 支払残額 | [optional] 
**DueDate** | Pointer to **string** | 支払期日 (yyyy-mm-dd) | [optional] 
**Id** | **int64** | 取引ID | 
**IssueDate** | **string** | 発生日 (yyyy-mm-dd) | 
**PartnerCode** | Pointer to **NullableString** | 取引先コード | [optional] 
**PartnerId** | **int64** | 取引先ID | 
**Payments** | Pointer to [**[]DealPaymentsInner**](DealPaymentsInner.md) | 取引の支払行 | [optional] 
**Receipts** | Pointer to [**[]DealReceiptsInner**](DealReceiptsInner.md) | 証憑ファイル（ファイルボックスのファイル） | [optional] 
**RefNumber** | Pointer to **string** | 管理番号 | [optional] 
**Renews** | Pointer to [**[]DealRenewsInner**](DealRenewsInner.md) | 取引の+更新行 | [optional] 
**Status** | **string** | 決済状況 (未決済: unsettled, 完了: settled) | 
**Type** | Pointer to **string** | 収支区分 (収入: income, 支出: expense) | [optional] 

## Methods

### NewDeal

`func NewDeal(amount int64, companyId int64, id int64, issueDate string, partnerId int64, status string, ) *Deal`

NewDeal instantiates a new Deal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealWithDefaults

`func NewDealWithDefaults() *Deal`

NewDealWithDefaults instantiates a new Deal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Deal) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Deal) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Deal) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCompanyId

`func (o *Deal) GetCompanyId() int64`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *Deal) GetCompanyIdOk() (*int64, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *Deal) SetCompanyId(v int64)`

SetCompanyId sets CompanyId field to given value.


### GetDetails

`func (o *Deal) GetDetails() []DealDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Deal) GetDetailsOk() (*[]DealDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Deal) SetDetails(v []DealDetailsInner)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Deal) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetDueAmount

`func (o *Deal) GetDueAmount() int64`

GetDueAmount returns the DueAmount field if non-nil, zero value otherwise.

### GetDueAmountOk

`func (o *Deal) GetDueAmountOk() (*int64, bool)`

GetDueAmountOk returns a tuple with the DueAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueAmount

`func (o *Deal) SetDueAmount(v int64)`

SetDueAmount sets DueAmount field to given value.

### HasDueAmount

`func (o *Deal) HasDueAmount() bool`

HasDueAmount returns a boolean if a field has been set.

### GetDueDate

`func (o *Deal) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *Deal) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *Deal) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *Deal) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetId

`func (o *Deal) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Deal) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Deal) SetId(v int64)`

SetId sets Id field to given value.


### GetIssueDate

`func (o *Deal) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *Deal) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *Deal) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetPartnerCode

`func (o *Deal) GetPartnerCode() string`

GetPartnerCode returns the PartnerCode field if non-nil, zero value otherwise.

### GetPartnerCodeOk

`func (o *Deal) GetPartnerCodeOk() (*string, bool)`

GetPartnerCodeOk returns a tuple with the PartnerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCode

`func (o *Deal) SetPartnerCode(v string)`

SetPartnerCode sets PartnerCode field to given value.

### HasPartnerCode

`func (o *Deal) HasPartnerCode() bool`

HasPartnerCode returns a boolean if a field has been set.

### SetPartnerCodeNil

`func (o *Deal) SetPartnerCodeNil(b bool)`

 SetPartnerCodeNil sets the value for PartnerCode to be an explicit nil

### UnsetPartnerCode
`func (o *Deal) UnsetPartnerCode()`

UnsetPartnerCode ensures that no value is present for PartnerCode, not even an explicit nil
### GetPartnerId

`func (o *Deal) GetPartnerId() int64`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *Deal) GetPartnerIdOk() (*int64, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *Deal) SetPartnerId(v int64)`

SetPartnerId sets PartnerId field to given value.


### GetPayments

`func (o *Deal) GetPayments() []DealPaymentsInner`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *Deal) GetPaymentsOk() (*[]DealPaymentsInner, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *Deal) SetPayments(v []DealPaymentsInner)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *Deal) HasPayments() bool`

HasPayments returns a boolean if a field has been set.

### GetReceipts

`func (o *Deal) GetReceipts() []DealReceiptsInner`

GetReceipts returns the Receipts field if non-nil, zero value otherwise.

### GetReceiptsOk

`func (o *Deal) GetReceiptsOk() (*[]DealReceiptsInner, bool)`

GetReceiptsOk returns a tuple with the Receipts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceipts

`func (o *Deal) SetReceipts(v []DealReceiptsInner)`

SetReceipts sets Receipts field to given value.

### HasReceipts

`func (o *Deal) HasReceipts() bool`

HasReceipts returns a boolean if a field has been set.

### GetRefNumber

`func (o *Deal) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *Deal) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *Deal) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *Deal) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetRenews

`func (o *Deal) GetRenews() []DealRenewsInner`

GetRenews returns the Renews field if non-nil, zero value otherwise.

### GetRenewsOk

`func (o *Deal) GetRenewsOk() (*[]DealRenewsInner, bool)`

GetRenewsOk returns a tuple with the Renews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenews

`func (o *Deal) SetRenews(v []DealRenewsInner)`

SetRenews sets Renews field to given value.

### HasRenews

`func (o *Deal) HasRenews() bool`

HasRenews returns a boolean if a field has been set.

### GetStatus

`func (o *Deal) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Deal) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Deal) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *Deal) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Deal) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Deal) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Deal) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


