# DealCreateResponseDealReceiptsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | 作成日時（ISO8601形式） | 
**Description** | Pointer to **string** | メモ | [optional] 
**Id** | **int64** | 証憑ファイルID（ファイルボックスのファイルID） | 
**IssueDate** | Pointer to **string** | 発生日 | [optional] 
**MimeType** | **string** | MIMEタイプ | 
**Origin** | **string** | アップロード元種別 | 
**ReceiptMetadatum** | Pointer to [**DealReceiptsInnerReceiptMetadatum**](DealReceiptsInnerReceiptMetadatum.md) |  | [optional] 
**Status** | **string** | ステータス(confirmed:確認済み、deleted:削除済み、ignored:無視) | 
**User** | [**DealReceiptsInnerUser**](DealReceiptsInnerUser.md) |  | 

## Methods

### NewDealCreateResponseDealReceiptsInner

`func NewDealCreateResponseDealReceiptsInner(createdAt string, id int64, mimeType string, origin string, status string, user DealReceiptsInnerUser, ) *DealCreateResponseDealReceiptsInner`

NewDealCreateResponseDealReceiptsInner instantiates a new DealCreateResponseDealReceiptsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealCreateResponseDealReceiptsInnerWithDefaults

`func NewDealCreateResponseDealReceiptsInnerWithDefaults() *DealCreateResponseDealReceiptsInner`

NewDealCreateResponseDealReceiptsInnerWithDefaults instantiates a new DealCreateResponseDealReceiptsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DealCreateResponseDealReceiptsInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DealCreateResponseDealReceiptsInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DealCreateResponseDealReceiptsInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *DealCreateResponseDealReceiptsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DealCreateResponseDealReceiptsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DealCreateResponseDealReceiptsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DealCreateResponseDealReceiptsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *DealCreateResponseDealReceiptsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DealCreateResponseDealReceiptsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DealCreateResponseDealReceiptsInner) SetId(v int64)`

SetId sets Id field to given value.


### GetIssueDate

`func (o *DealCreateResponseDealReceiptsInner) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *DealCreateResponseDealReceiptsInner) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *DealCreateResponseDealReceiptsInner) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *DealCreateResponseDealReceiptsInner) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetMimeType

`func (o *DealCreateResponseDealReceiptsInner) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *DealCreateResponseDealReceiptsInner) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *DealCreateResponseDealReceiptsInner) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetOrigin

`func (o *DealCreateResponseDealReceiptsInner) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *DealCreateResponseDealReceiptsInner) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *DealCreateResponseDealReceiptsInner) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetReceiptMetadatum

`func (o *DealCreateResponseDealReceiptsInner) GetReceiptMetadatum() DealReceiptsInnerReceiptMetadatum`

GetReceiptMetadatum returns the ReceiptMetadatum field if non-nil, zero value otherwise.

### GetReceiptMetadatumOk

`func (o *DealCreateResponseDealReceiptsInner) GetReceiptMetadatumOk() (*DealReceiptsInnerReceiptMetadatum, bool)`

GetReceiptMetadatumOk returns a tuple with the ReceiptMetadatum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptMetadatum

`func (o *DealCreateResponseDealReceiptsInner) SetReceiptMetadatum(v DealReceiptsInnerReceiptMetadatum)`

SetReceiptMetadatum sets ReceiptMetadatum field to given value.

### HasReceiptMetadatum

`func (o *DealCreateResponseDealReceiptsInner) HasReceiptMetadatum() bool`

HasReceiptMetadatum returns a boolean if a field has been set.

### GetStatus

`func (o *DealCreateResponseDealReceiptsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DealCreateResponseDealReceiptsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DealCreateResponseDealReceiptsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUser

`func (o *DealCreateResponseDealReceiptsInner) GetUser() DealReceiptsInnerUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DealCreateResponseDealReceiptsInner) GetUserOk() (*DealReceiptsInnerUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DealCreateResponseDealReceiptsInner) SetUser(v DealReceiptsInnerUser)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


