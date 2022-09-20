# ManualJournal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adjustment** | **bool** | 決算整理仕訳フラグ（falseまたは未指定の場合: 日常仕訳） | 
**CompanyId** | **int32** | 事業所ID | 
**Details** | [**[]ManualJournalDetailsInner**](ManualJournalDetailsInner.md) | 貸借行一覧（配列）: 貸借合わせて100行まで登録できます。 | 
**Id** | **int32** | 振替伝票ID | 
**IssueDate** | **string** | 発生日 (yyyy-mm-dd) | 
**ReceiptIds** | Pointer to **[]int32** | 証憑ファイルID（ファイルボックスのファイルID） | [optional] 
**TxnNumber** | **NullableString** | 仕訳番号 | 

## Methods

### NewManualJournal

`func NewManualJournal(adjustment bool, companyId int32, details []ManualJournalDetailsInner, id int32, issueDate string, txnNumber NullableString, ) *ManualJournal`

NewManualJournal instantiates a new ManualJournal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualJournalWithDefaults

`func NewManualJournalWithDefaults() *ManualJournal`

NewManualJournalWithDefaults instantiates a new ManualJournal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdjustment

`func (o *ManualJournal) GetAdjustment() bool`

GetAdjustment returns the Adjustment field if non-nil, zero value otherwise.

### GetAdjustmentOk

`func (o *ManualJournal) GetAdjustmentOk() (*bool, bool)`

GetAdjustmentOk returns a tuple with the Adjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustment

`func (o *ManualJournal) SetAdjustment(v bool)`

SetAdjustment sets Adjustment field to given value.


### GetCompanyId

`func (o *ManualJournal) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ManualJournal) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ManualJournal) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetDetails

`func (o *ManualJournal) GetDetails() []ManualJournalDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ManualJournal) GetDetailsOk() (*[]ManualJournalDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ManualJournal) SetDetails(v []ManualJournalDetailsInner)`

SetDetails sets Details field to given value.


### GetId

`func (o *ManualJournal) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManualJournal) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManualJournal) SetId(v int32)`

SetId sets Id field to given value.


### GetIssueDate

`func (o *ManualJournal) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *ManualJournal) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *ManualJournal) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetReceiptIds

`func (o *ManualJournal) GetReceiptIds() []int32`

GetReceiptIds returns the ReceiptIds field if non-nil, zero value otherwise.

### GetReceiptIdsOk

`func (o *ManualJournal) GetReceiptIdsOk() (*[]int32, bool)`

GetReceiptIdsOk returns a tuple with the ReceiptIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptIds

`func (o *ManualJournal) SetReceiptIds(v []int32)`

SetReceiptIds sets ReceiptIds field to given value.

### HasReceiptIds

`func (o *ManualJournal) HasReceiptIds() bool`

HasReceiptIds returns a boolean if a field has been set.

### GetTxnNumber

`func (o *ManualJournal) GetTxnNumber() string`

GetTxnNumber returns the TxnNumber field if non-nil, zero value otherwise.

### GetTxnNumberOk

`func (o *ManualJournal) GetTxnNumberOk() (*string, bool)`

GetTxnNumberOk returns a tuple with the TxnNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnNumber

`func (o *ManualJournal) SetTxnNumber(v string)`

SetTxnNumber sets TxnNumber field to given value.


### SetTxnNumberNil

`func (o *ManualJournal) SetTxnNumberNil(b bool)`

 SetTxnNumberNil sets the value for TxnNumber to be an explicit nil

### UnsetTxnNumber
`func (o *ManualJournal) UnsetTxnNumber()`

UnsetTxnNumber ensures that no value is present for TxnNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


