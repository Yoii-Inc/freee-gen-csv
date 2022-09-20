# ApprovalRequestFormIndexResponseApprovalRequestFormsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int32** | 事業所ID | 
**CreatedDate** | **string** | 作成日時 | 
**Description** | **string** | 申請フォームの説明 | 
**FormOrder** | **NullableInt32** | 表示順（申請者が選択する申請フォームの表示順を設定できます。小さい数ほど上位に表示されます。（0を除く整数のみ。マイナス不可）未入力の場合、表示順が後ろになります。同じ数字が入力された場合、登録順で表示されます。） | 
**Id** | **int32** | 申請フォームID | 
**Name** | **string** | 申請フォームの名前 | 
**RouteSettingCount** | **int32** | 適用された経路数 | 
**Status** | **string** | ステータス(draft: 申請で使用しない、active: 申請で使用する、deleted: 削除済み) | 

## Methods

### NewApprovalRequestFormIndexResponseApprovalRequestFormsInner

`func NewApprovalRequestFormIndexResponseApprovalRequestFormsInner(companyId int32, createdDate string, description string, formOrder NullableInt32, id int32, name string, routeSettingCount int32, status string, ) *ApprovalRequestFormIndexResponseApprovalRequestFormsInner`

NewApprovalRequestFormIndexResponseApprovalRequestFormsInner instantiates a new ApprovalRequestFormIndexResponseApprovalRequestFormsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalRequestFormIndexResponseApprovalRequestFormsInnerWithDefaults

`func NewApprovalRequestFormIndexResponseApprovalRequestFormsInnerWithDefaults() *ApprovalRequestFormIndexResponseApprovalRequestFormsInner`

NewApprovalRequestFormIndexResponseApprovalRequestFormsInnerWithDefaults instantiates a new ApprovalRequestFormIndexResponseApprovalRequestFormsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetCreatedDate

`func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.


### GetDescription

`func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFormOrder

`func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetFormOrder() int32`

GetFormOrder returns the FormOrder field if non-nil, zero value otherwise.

### GetFormOrderOk

`func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetFormOrderOk() (*int32, bool)`

GetFormOrderOk returns a tuple with the FormOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormOrder

`func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) SetFormOrder(v int32)`

SetFormOrder sets FormOrder field to given value.


### SetFormOrderNil

`func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) SetFormOrderNil(b bool)`

 SetFormOrderNil sets the value for FormOrder to be an explicit nil

### UnsetFormOrder
`func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) UnsetFormOrder()`

UnsetFormOrder ensures that no value is present for FormOrder, not even an explicit nil
### GetId

`func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) SetName(v string)`

SetName sets Name field to given value.


### GetRouteSettingCount

`func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetRouteSettingCount() int32`

GetRouteSettingCount returns the RouteSettingCount field if non-nil, zero value otherwise.

### GetRouteSettingCountOk

`func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetRouteSettingCountOk() (*int32, bool)`

GetRouteSettingCountOk returns a tuple with the RouteSettingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteSettingCount

`func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) SetRouteSettingCount(v int32)`

SetRouteSettingCount sets RouteSettingCount field to given value.


### GetStatus

`func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApprovalRequestFormIndexResponseApprovalRequestFormsInner) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


