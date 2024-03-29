# Receipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | 作成日時（ISO8601形式） | 
**Description** | Pointer to **string** | メモ | [optional] 
**FileSrc** | **string** | ファイルのダウンロードURL（freeeにログインした状態でのみ閲覧可能です。） &lt;br&gt; &lt;br&gt; file_srcは廃止予定の属性になります。&lt;br&gt; file_srcに替わり、証憑ファイルのダウンロード APIをご利用ください。&lt;br&gt; 証憑ファイルのダウンロードAPIを利用することで、以下のようになります。 &lt;ul&gt;   &lt;li&gt;アプリケーション利用者はfreee APIアプリケーションにログインしていれば、証憑ダウンロード毎にfreeeに改めてログインすることなくファイルが参照できるようになります。&lt;/li&gt; &lt;/ul&gt; | 
**Id** | **int64** | 証憑ファイルID | 
**IssueDate** | Pointer to **string** | 発生日 | [optional] 
**MimeType** | **string** | MIMEタイプ | 
**Origin** | **string** | アップロード元種別 | 
**ReceiptMetadatum** | Pointer to [**DealReceiptsInnerReceiptMetadatum**](DealReceiptsInnerReceiptMetadatum.md) |  | [optional] 
**Status** | **string** | ステータス(confirmed:確認済み、deleted:削除済み、ignored:無視) | 
**User** | [**DealReceiptsInnerUser**](DealReceiptsInnerUser.md) |  | 

## Methods

### NewReceipt

`func NewReceipt(createdAt string, fileSrc string, id int64, mimeType string, origin string, status string, user DealReceiptsInnerUser, ) *Receipt`

NewReceipt instantiates a new Receipt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiptWithDefaults

`func NewReceiptWithDefaults() *Receipt`

NewReceiptWithDefaults instantiates a new Receipt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Receipt) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Receipt) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Receipt) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *Receipt) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Receipt) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Receipt) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Receipt) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFileSrc

`func (o *Receipt) GetFileSrc() string`

GetFileSrc returns the FileSrc field if non-nil, zero value otherwise.

### GetFileSrcOk

`func (o *Receipt) GetFileSrcOk() (*string, bool)`

GetFileSrcOk returns a tuple with the FileSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSrc

`func (o *Receipt) SetFileSrc(v string)`

SetFileSrc sets FileSrc field to given value.


### GetId

`func (o *Receipt) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Receipt) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Receipt) SetId(v int64)`

SetId sets Id field to given value.


### GetIssueDate

`func (o *Receipt) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *Receipt) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *Receipt) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *Receipt) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetMimeType

`func (o *Receipt) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *Receipt) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *Receipt) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetOrigin

`func (o *Receipt) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *Receipt) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *Receipt) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetReceiptMetadatum

`func (o *Receipt) GetReceiptMetadatum() DealReceiptsInnerReceiptMetadatum`

GetReceiptMetadatum returns the ReceiptMetadatum field if non-nil, zero value otherwise.

### GetReceiptMetadatumOk

`func (o *Receipt) GetReceiptMetadatumOk() (*DealReceiptsInnerReceiptMetadatum, bool)`

GetReceiptMetadatumOk returns a tuple with the ReceiptMetadatum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptMetadatum

`func (o *Receipt) SetReceiptMetadatum(v DealReceiptsInnerReceiptMetadatum)`

SetReceiptMetadatum sets ReceiptMetadatum field to given value.

### HasReceiptMetadatum

`func (o *Receipt) HasReceiptMetadatum() bool`

HasReceiptMetadatum returns a boolean if a field has been set.

### GetStatus

`func (o *Receipt) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Receipt) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Receipt) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUser

`func (o *Receipt) GetUser() DealReceiptsInnerUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Receipt) GetUserOk() (*DealReceiptsInnerUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Receipt) SetUser(v DealReceiptsInnerUser)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


