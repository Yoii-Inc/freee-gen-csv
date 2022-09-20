# DealReceiptsInnerReceiptMetadatum

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **NullableInt64** | 金額 | [optional] 
**IssueDate** | Pointer to **NullableString** | 発行日 (yyyy-mm-dd) | [optional] 
**PartnerName** | Pointer to **NullableString** | 発行元 | [optional] 

## Methods

### NewDealReceiptsInnerReceiptMetadatum

`func NewDealReceiptsInnerReceiptMetadatum() *DealReceiptsInnerReceiptMetadatum`

NewDealReceiptsInnerReceiptMetadatum instantiates a new DealReceiptsInnerReceiptMetadatum object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealReceiptsInnerReceiptMetadatumWithDefaults

`func NewDealReceiptsInnerReceiptMetadatumWithDefaults() *DealReceiptsInnerReceiptMetadatum`

NewDealReceiptsInnerReceiptMetadatumWithDefaults instantiates a new DealReceiptsInnerReceiptMetadatum object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DealReceiptsInnerReceiptMetadatum) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DealReceiptsInnerReceiptMetadatum) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DealReceiptsInnerReceiptMetadatum) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *DealReceiptsInnerReceiptMetadatum) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *DealReceiptsInnerReceiptMetadatum) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *DealReceiptsInnerReceiptMetadatum) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetIssueDate

`func (o *DealReceiptsInnerReceiptMetadatum) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *DealReceiptsInnerReceiptMetadatum) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *DealReceiptsInnerReceiptMetadatum) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *DealReceiptsInnerReceiptMetadatum) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### SetIssueDateNil

`func (o *DealReceiptsInnerReceiptMetadatum) SetIssueDateNil(b bool)`

 SetIssueDateNil sets the value for IssueDate to be an explicit nil

### UnsetIssueDate
`func (o *DealReceiptsInnerReceiptMetadatum) UnsetIssueDate()`

UnsetIssueDate ensures that no value is present for IssueDate, not even an explicit nil
### GetPartnerName

`func (o *DealReceiptsInnerReceiptMetadatum) GetPartnerName() string`

GetPartnerName returns the PartnerName field if non-nil, zero value otherwise.

### GetPartnerNameOk

`func (o *DealReceiptsInnerReceiptMetadatum) GetPartnerNameOk() (*string, bool)`

GetPartnerNameOk returns a tuple with the PartnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerName

`func (o *DealReceiptsInnerReceiptMetadatum) SetPartnerName(v string)`

SetPartnerName sets PartnerName field to given value.

### HasPartnerName

`func (o *DealReceiptsInnerReceiptMetadatum) HasPartnerName() bool`

HasPartnerName returns a boolean if a field has been set.

### SetPartnerNameNil

`func (o *DealReceiptsInnerReceiptMetadatum) SetPartnerNameNil(b bool)`

 SetPartnerNameNil sets the value for PartnerName to be an explicit nil

### UnsetPartnerName
`func (o *DealReceiptsInnerReceiptMetadatum) UnsetPartnerName()`

UnsetPartnerName ensures that no value is present for PartnerName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


