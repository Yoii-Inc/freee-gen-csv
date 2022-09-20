# ApprovalRequestResponseApprovalRequestRequestItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 項目ID | 
**Type** | **string** | 項目タイプ(title: 申請タイトル, single_line: 自由記述形式 1行, multi_line: 自由記述形式 複数行, select: プルダウン, date: 日付, amount: 金額, receipt: 添付ファイル, section: 部門ID, partner: 取引先ID, ninja_sign_document: 契約書（freeeサイン連携）) | 
**Value** | **string** | 項目の値 | 

## Methods

### NewApprovalRequestResponseApprovalRequestRequestItemsInner

`func NewApprovalRequestResponseApprovalRequestRequestItemsInner(id int32, type_ string, value string, ) *ApprovalRequestResponseApprovalRequestRequestItemsInner`

NewApprovalRequestResponseApprovalRequestRequestItemsInner instantiates a new ApprovalRequestResponseApprovalRequestRequestItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalRequestResponseApprovalRequestRequestItemsInnerWithDefaults

`func NewApprovalRequestResponseApprovalRequestRequestItemsInnerWithDefaults() *ApprovalRequestResponseApprovalRequestRequestItemsInner`

NewApprovalRequestResponseApprovalRequestRequestItemsInnerWithDefaults instantiates a new ApprovalRequestResponseApprovalRequestRequestItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApprovalRequestResponseApprovalRequestRequestItemsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalRequestResponseApprovalRequestRequestItemsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalRequestResponseApprovalRequestRequestItemsInner) SetId(v int32)`

SetId sets Id field to given value.


### GetType

`func (o *ApprovalRequestResponseApprovalRequestRequestItemsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApprovalRequestResponseApprovalRequestRequestItemsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApprovalRequestResponseApprovalRequestRequestItemsInner) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ApprovalRequestResponseApprovalRequestRequestItemsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ApprovalRequestResponseApprovalRequestRequestItemsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ApprovalRequestResponseApprovalRequestRequestItemsInner) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


