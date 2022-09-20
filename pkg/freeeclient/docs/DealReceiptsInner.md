# DealReceiptsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | 作成日時（ISO8601形式） | 
**Description** | Pointer to **string** | メモ | [optional] 
**FileSrc** | **string** | ファイルのダウンロードURL（freeeにログインした状態でのみ閲覧可能です。） &lt;br&gt; &lt;br&gt; file_srcは廃止予定の属性になります。&lt;br&gt; file_srcに替わり、証憑ファイル（ファイルボックスのファイル）のダウンロード APIをご利用ください。&lt;br&gt; 証憑ファイル（ファイルボックスのファイル）のダウンロードAPIを利用することで、以下のようになります。 &lt;ul&gt;   &lt;li&gt;アプリケーション利用者はfreee APIアプリケーションにログインしていれば、証憑ダウンロード毎にfreeeに改めてログインすることなくファイルが参照できるようになります。&lt;/li&gt; &lt;/ul&gt; | 
**Id** | **int64** | 証憑ファイルID（ファイルボックスのファイルID） | 
**IssueDate** | Pointer to **string** | 発生日 | [optional] 
**MimeType** | **string** | MIMEタイプ | 
**Origin** | **string** | アップロード元種別 | 
**ReceiptMetadatum** | Pointer to [**DealReceiptsInnerReceiptMetadatum**](DealReceiptsInnerReceiptMetadatum.md) |  | [optional] 
**Status** | **string** | ステータス(confirmed:確認済み、deleted:削除済み、ignored:無視) | 
**User** | [**DealReceiptsInnerUser**](DealReceiptsInnerUser.md) |  | 

## Methods

### NewDealReceiptsInner

`func NewDealReceiptsInner(createdAt string, fileSrc string, id int64, mimeType string, origin string, status string, user DealReceiptsInnerUser, ) *DealReceiptsInner`

NewDealReceiptsInner instantiates a new DealReceiptsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealReceiptsInnerWithDefaults

`func NewDealReceiptsInnerWithDefaults() *DealReceiptsInner`

NewDealReceiptsInnerWithDefaults instantiates a new DealReceiptsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DealReceiptsInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DealReceiptsInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DealReceiptsInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *DealReceiptsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DealReceiptsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DealReceiptsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DealReceiptsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFileSrc

`func (o *DealReceiptsInner) GetFileSrc() string`

GetFileSrc returns the FileSrc field if non-nil, zero value otherwise.

### GetFileSrcOk

`func (o *DealReceiptsInner) GetFileSrcOk() (*string, bool)`

GetFileSrcOk returns a tuple with the FileSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSrc

`func (o *DealReceiptsInner) SetFileSrc(v string)`

SetFileSrc sets FileSrc field to given value.


### GetId

`func (o *DealReceiptsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DealReceiptsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DealReceiptsInner) SetId(v int64)`

SetId sets Id field to given value.


### GetIssueDate

`func (o *DealReceiptsInner) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *DealReceiptsInner) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *DealReceiptsInner) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *DealReceiptsInner) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetMimeType

`func (o *DealReceiptsInner) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *DealReceiptsInner) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *DealReceiptsInner) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetOrigin

`func (o *DealReceiptsInner) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *DealReceiptsInner) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *DealReceiptsInner) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetReceiptMetadatum

`func (o *DealReceiptsInner) GetReceiptMetadatum() DealReceiptsInnerReceiptMetadatum`

GetReceiptMetadatum returns the ReceiptMetadatum field if non-nil, zero value otherwise.

### GetReceiptMetadatumOk

`func (o *DealReceiptsInner) GetReceiptMetadatumOk() (*DealReceiptsInnerReceiptMetadatum, bool)`

GetReceiptMetadatumOk returns a tuple with the ReceiptMetadatum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptMetadatum

`func (o *DealReceiptsInner) SetReceiptMetadatum(v DealReceiptsInnerReceiptMetadatum)`

SetReceiptMetadatum sets ReceiptMetadatum field to given value.

### HasReceiptMetadatum

`func (o *DealReceiptsInner) HasReceiptMetadatum() bool`

HasReceiptMetadatum returns a boolean if a field has been set.

### GetStatus

`func (o *DealReceiptsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DealReceiptsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DealReceiptsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUser

`func (o *DealReceiptsInner) GetUser() DealReceiptsInnerUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DealReceiptsInner) GetUserOk() (*DealReceiptsInnerUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DealReceiptsInner) SetUser(v DealReceiptsInnerUser)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


