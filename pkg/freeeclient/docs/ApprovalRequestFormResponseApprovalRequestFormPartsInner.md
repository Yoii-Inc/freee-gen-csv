# ApprovalRequestFormResponseApprovalRequestFormPartsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotation** | Pointer to **NullableString** | 追加説明 | [optional] 
**Id** | **int64** | 項目ID | 
**Label** | Pointer to **string** | 項目名 | [optional] 
**MaxAmount** | Pointer to **NullableInt64** | 上限金額 | [optional] 
**MinAmount** | Pointer to **NullableInt64** | 下限金額 | [optional] 
**Order** | Pointer to **int64** | 順序 | [optional] 
**Required** | Pointer to **NullableBool** | 必須かどうか | [optional] 
**Type** | Pointer to **string** | 項目種別 (title: 申請タイトル, single_line: 自由記述形式 1行, multi_line: 自由記述形式 複数行, select: プルダウン, date: 日付, amount: 金額, receipt: 添付ファイル, section: 部門ID, partner: 取引先ID, ninja_sign_document: 契約書（freeeサイン連携）) | [optional] 
**Values** | Pointer to [**[]ApprovalRequestFormResponseApprovalRequestFormPartsInnerValuesInner**](ApprovalRequestFormResponseApprovalRequestFormPartsInnerValuesInner.md) | 選択項目 | [optional] 

## Methods

### NewApprovalRequestFormResponseApprovalRequestFormPartsInner

`func NewApprovalRequestFormResponseApprovalRequestFormPartsInner(id int64, ) *ApprovalRequestFormResponseApprovalRequestFormPartsInner`

NewApprovalRequestFormResponseApprovalRequestFormPartsInner instantiates a new ApprovalRequestFormResponseApprovalRequestFormPartsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalRequestFormResponseApprovalRequestFormPartsInnerWithDefaults

`func NewApprovalRequestFormResponseApprovalRequestFormPartsInnerWithDefaults() *ApprovalRequestFormResponseApprovalRequestFormPartsInner`

NewApprovalRequestFormResponseApprovalRequestFormPartsInnerWithDefaults instantiates a new ApprovalRequestFormResponseApprovalRequestFormPartsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotation

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetAnnotation() string`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetAnnotationOk() (*string, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) SetAnnotation(v string)`

SetAnnotation sets Annotation field to given value.

### HasAnnotation

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.

### SetAnnotationNil

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) SetAnnotationNil(b bool)`

 SetAnnotationNil sets the value for Annotation to be an explicit nil

### UnsetAnnotation
`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) UnsetAnnotation()`

UnsetAnnotation ensures that no value is present for Annotation, not even an explicit nil
### GetId

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) SetId(v int64)`

SetId sets Id field to given value.


### GetLabel

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMaxAmount

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetMaxAmount() int64`

GetMaxAmount returns the MaxAmount field if non-nil, zero value otherwise.

### GetMaxAmountOk

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetMaxAmountOk() (*int64, bool)`

GetMaxAmountOk returns a tuple with the MaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmount

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) SetMaxAmount(v int64)`

SetMaxAmount sets MaxAmount field to given value.

### HasMaxAmount

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) HasMaxAmount() bool`

HasMaxAmount returns a boolean if a field has been set.

### SetMaxAmountNil

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) SetMaxAmountNil(b bool)`

 SetMaxAmountNil sets the value for MaxAmount to be an explicit nil

### UnsetMaxAmount
`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) UnsetMaxAmount()`

UnsetMaxAmount ensures that no value is present for MaxAmount, not even an explicit nil
### GetMinAmount

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetMinAmount() int64`

GetMinAmount returns the MinAmount field if non-nil, zero value otherwise.

### GetMinAmountOk

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetMinAmountOk() (*int64, bool)`

GetMinAmountOk returns a tuple with the MinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmount

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) SetMinAmount(v int64)`

SetMinAmount sets MinAmount field to given value.

### HasMinAmount

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) HasMinAmount() bool`

HasMinAmount returns a boolean if a field has been set.

### SetMinAmountNil

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) SetMinAmountNil(b bool)`

 SetMinAmountNil sets the value for MinAmount to be an explicit nil

### UnsetMinAmount
`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) UnsetMinAmount()`

UnsetMinAmount ensures that no value is present for MinAmount, not even an explicit nil
### GetOrder

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetRequired

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil
### GetType

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValues

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetValues() []ApprovalRequestFormResponseApprovalRequestFormPartsInnerValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) GetValuesOk() (*[]ApprovalRequestFormResponseApprovalRequestFormPartsInnerValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) SetValues(v []ApprovalRequestFormResponseApprovalRequestFormPartsInnerValuesInner)`

SetValues sets Values field to given value.

### HasValues

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *ApprovalRequestFormResponseApprovalRequestFormPartsInner) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


