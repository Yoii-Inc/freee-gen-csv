# ApprovalRequestResponseApprovalRequestRequestItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 項目ID | 
**Type** | **string** | 項目タイプ(title: 申請タイトル, single_line: 自由記述形式 1行, multi_line: 自由記述形式 複数行, select: プルダウン, date: 日付, amount: 金額, receipt: 添付ファイル, section: 部門ID, partner: 取引先ID) | 
**Value** | **string** | 項目の値 | 

## Methods

### NewApprovalRequestResponseApprovalRequestRequestItems

`func NewApprovalRequestResponseApprovalRequestRequestItems(id int32, type_ string, value string, ) *ApprovalRequestResponseApprovalRequestRequestItems`

NewApprovalRequestResponseApprovalRequestRequestItems instantiates a new ApprovalRequestResponseApprovalRequestRequestItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalRequestResponseApprovalRequestRequestItemsWithDefaults

`func NewApprovalRequestResponseApprovalRequestRequestItemsWithDefaults() *ApprovalRequestResponseApprovalRequestRequestItems`

NewApprovalRequestResponseApprovalRequestRequestItemsWithDefaults instantiates a new ApprovalRequestResponseApprovalRequestRequestItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApprovalRequestResponseApprovalRequestRequestItems) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalRequestResponseApprovalRequestRequestItems) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalRequestResponseApprovalRequestRequestItems) SetId(v int32)`

SetId sets Id field to given value.


### GetType

`func (o *ApprovalRequestResponseApprovalRequestRequestItems) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApprovalRequestResponseApprovalRequestRequestItems) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApprovalRequestResponseApprovalRequestRequestItems) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ApprovalRequestResponseApprovalRequestRequestItems) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ApprovalRequestResponseApprovalRequestRequestItems) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ApprovalRequestResponseApprovalRequestRequestItems) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


