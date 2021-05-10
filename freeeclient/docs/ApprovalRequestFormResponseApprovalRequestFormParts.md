# ApprovalRequestFormResponseApprovalRequestFormParts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotation** | Pointer to **NullableString** | 追加説明 | [optional] 
**Id** | **int32** | 項目ID | 
**Label** | Pointer to **string** | 項目名 | [optional] 
**MaxAmount** | Pointer to **NullableInt32** | 上限金額 | [optional] 
**MinAmount** | Pointer to **NullableInt32** | 下限金額 | [optional] 
**Order** | Pointer to **int32** | 順序 | [optional] 
**Required** | Pointer to **NullableBool** | 必須かどうか | [optional] 
**Type** | Pointer to **string** | 項目種別 (title: 申請タイトル, single_line: 自由記述形式 1行, multi_line: 自由記述形式 複数行, select: プルダウン, date: 日付, amount: 金額, receipt: 添付ファイル, section: 部門ID, partner: 取引先ID) | [optional] 
**Values** | Pointer to [**[]ApprovalRequestFormResponseApprovalRequestFormValues**](ApprovalRequestFormResponseApprovalRequestFormValues.md) | 選択項目 | [optional] 

## Methods

### NewApprovalRequestFormResponseApprovalRequestFormParts

`func NewApprovalRequestFormResponseApprovalRequestFormParts(id int32, ) *ApprovalRequestFormResponseApprovalRequestFormParts`

NewApprovalRequestFormResponseApprovalRequestFormParts instantiates a new ApprovalRequestFormResponseApprovalRequestFormParts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalRequestFormResponseApprovalRequestFormPartsWithDefaults

`func NewApprovalRequestFormResponseApprovalRequestFormPartsWithDefaults() *ApprovalRequestFormResponseApprovalRequestFormParts`

NewApprovalRequestFormResponseApprovalRequestFormPartsWithDefaults instantiates a new ApprovalRequestFormResponseApprovalRequestFormParts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotation

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) GetAnnotation() string`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) GetAnnotationOk() (*string, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) SetAnnotation(v string)`

SetAnnotation sets Annotation field to given value.

### HasAnnotation

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.

### SetAnnotationNil

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) SetAnnotationNil(b bool)`

 SetAnnotationNil sets the value for Annotation to be an explicit nil

### UnsetAnnotation
`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) UnsetAnnotation()`

UnsetAnnotation ensures that no value is present for Annotation, not even an explicit nil
### GetId

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) SetId(v int32)`

SetId sets Id field to given value.


### GetLabel

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMaxAmount

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) GetMaxAmount() int32`

GetMaxAmount returns the MaxAmount field if non-nil, zero value otherwise.

### GetMaxAmountOk

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) GetMaxAmountOk() (*int32, bool)`

GetMaxAmountOk returns a tuple with the MaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmount

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) SetMaxAmount(v int32)`

SetMaxAmount sets MaxAmount field to given value.

### HasMaxAmount

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) HasMaxAmount() bool`

HasMaxAmount returns a boolean if a field has been set.

### SetMaxAmountNil

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) SetMaxAmountNil(b bool)`

 SetMaxAmountNil sets the value for MaxAmount to be an explicit nil

### UnsetMaxAmount
`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) UnsetMaxAmount()`

UnsetMaxAmount ensures that no value is present for MaxAmount, not even an explicit nil
### GetMinAmount

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) GetMinAmount() int32`

GetMinAmount returns the MinAmount field if non-nil, zero value otherwise.

### GetMinAmountOk

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) GetMinAmountOk() (*int32, bool)`

GetMinAmountOk returns a tuple with the MinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmount

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) SetMinAmount(v int32)`

SetMinAmount sets MinAmount field to given value.

### HasMinAmount

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) HasMinAmount() bool`

HasMinAmount returns a boolean if a field has been set.

### SetMinAmountNil

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) SetMinAmountNil(b bool)`

 SetMinAmountNil sets the value for MinAmount to be an explicit nil

### UnsetMinAmount
`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) UnsetMinAmount()`

UnsetMinAmount ensures that no value is present for MinAmount, not even an explicit nil
### GetOrder

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetRequired

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil
### GetType

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValues

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) GetValues() []ApprovalRequestFormResponseApprovalRequestFormValues`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) GetValuesOk() (*[]ApprovalRequestFormResponseApprovalRequestFormValues, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) SetValues(v []ApprovalRequestFormResponseApprovalRequestFormValues)`

SetValues sets Values field to given value.

### HasValues

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *ApprovalRequestFormResponseApprovalRequestFormParts) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


