# ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultRoute** | **bool** | 基本経路として設定されているかどうか&lt;br&gt;&lt;br&gt; リクエストパラメータusageに下記のいずれかが指定され、かつ、基本経路の場合はtrueになります。 * &#x60;TxnApproval&#x60; - 仕訳承認 * &#x60;ExpenseApplication&#x60; - 経費精算 * &#x60;PaymentRequest&#x60; - 支払依頼 * &#x60;ApprovalRequest&#x60;(リクエストパラメータrequest_form_idを同時に指定) - 各種申請 * &#x60;DocApproval&#x60; - 請求書等 (見積書・納品書・請求書・発注書)  &lt;a href&#x3D;\&quot;https://support.freee.co.jp/hc/ja/articles/900000507963\&quot; target&#x3D;\&quot;_blank\&quot;&gt;申請フォームの基本経路設定&lt;/a&gt;  | 
**DefinitionSystem** | Pointer to **bool** | システム作成の申請経路かどうか | [optional] 
**Description** | Pointer to **string** | 申請経路の説明 | [optional] 
**FirstStepId** | Pointer to **int32** | 最初の承認ステップのID | [optional] 
**Id** | **int32** | 申請経路ID | 
**Name** | Pointer to **string** | 申請経路名 | [optional] 
**RequestFormIds** | Pointer to **[]int32** | 申請経路で利用できる申請フォームID配列 | [optional] 
**Usages** | Pointer to **[]string** | 申請種別（申請経路を使用できる申請種別を示します。例えば、ApprovalRequest の場合は、各種申請で使用できる申請経路です。） * &#x60;TxnApproval&#x60; - 仕訳承認 * &#x60;ExpenseApplication&#x60; - 経費精算 * &#x60;PaymentRequest&#x60; - 支払依頼 * &#x60;ApprovalRequest&#x60; - 各種申請 * &#x60;DocApproval&#x60; - 請求書等 (見積書・納品書・請求書・発注書) | [optional] 
**UserId** | Pointer to **NullableInt32** | 更新したユーザーのユーザーID | [optional] 

## Methods

### NewApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner

`func NewApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner(defaultRoute bool, id int32, ) *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner`

NewApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner instantiates a new ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalFlowRoutesIndexResponseApprovalFlowRoutesInnerWithDefaults

`func NewApprovalFlowRoutesIndexResponseApprovalFlowRoutesInnerWithDefaults() *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner`

NewApprovalFlowRoutesIndexResponseApprovalFlowRoutesInnerWithDefaults instantiates a new ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultRoute

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) GetDefaultRoute() bool`

GetDefaultRoute returns the DefaultRoute field if non-nil, zero value otherwise.

### GetDefaultRouteOk

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) GetDefaultRouteOk() (*bool, bool)`

GetDefaultRouteOk returns a tuple with the DefaultRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoute

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) SetDefaultRoute(v bool)`

SetDefaultRoute sets DefaultRoute field to given value.


### GetDefinitionSystem

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) GetDefinitionSystem() bool`

GetDefinitionSystem returns the DefinitionSystem field if non-nil, zero value otherwise.

### GetDefinitionSystemOk

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) GetDefinitionSystemOk() (*bool, bool)`

GetDefinitionSystemOk returns a tuple with the DefinitionSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionSystem

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) SetDefinitionSystem(v bool)`

SetDefinitionSystem sets DefinitionSystem field to given value.

### HasDefinitionSystem

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) HasDefinitionSystem() bool`

HasDefinitionSystem returns a boolean if a field has been set.

### GetDescription

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFirstStepId

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) GetFirstStepId() int32`

GetFirstStepId returns the FirstStepId field if non-nil, zero value otherwise.

### GetFirstStepIdOk

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) GetFirstStepIdOk() (*int32, bool)`

GetFirstStepIdOk returns a tuple with the FirstStepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstStepId

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) SetFirstStepId(v int32)`

SetFirstStepId sets FirstStepId field to given value.

### HasFirstStepId

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) HasFirstStepId() bool`

HasFirstStepId returns a boolean if a field has been set.

### GetId

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequestFormIds

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) GetRequestFormIds() []int32`

GetRequestFormIds returns the RequestFormIds field if non-nil, zero value otherwise.

### GetRequestFormIdsOk

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) GetRequestFormIdsOk() (*[]int32, bool)`

GetRequestFormIdsOk returns a tuple with the RequestFormIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestFormIds

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) SetRequestFormIds(v []int32)`

SetRequestFormIds sets RequestFormIds field to given value.

### HasRequestFormIds

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) HasRequestFormIds() bool`

HasRequestFormIds returns a boolean if a field has been set.

### GetUsages

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) GetUsages() []string`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) GetUsagesOk() (*[]string, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) SetUsages(v []string)`

SetUsages sets Usages field to given value.

### HasUsages

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) HasUsages() bool`

HasUsages returns a boolean if a field has been set.

### GetUserId

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutesInner) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


