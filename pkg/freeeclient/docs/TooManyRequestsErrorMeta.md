# TooManyRequestsErrorMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int64** | 設定されている上限値 | 
**Period** | Pointer to **int64** | 使用回数をカウントする期間 (秒) | [optional] 
**Remaining** | **int64** | 上限に達するまでの使用可能回数 | 
**Reset** | **string** | （上限値に達した場合）使用回数がリセットされる時刻 | 

## Methods

### NewTooManyRequestsErrorMeta

`func NewTooManyRequestsErrorMeta(limit int64, remaining int64, reset string, ) *TooManyRequestsErrorMeta`

NewTooManyRequestsErrorMeta instantiates a new TooManyRequestsErrorMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTooManyRequestsErrorMetaWithDefaults

`func NewTooManyRequestsErrorMetaWithDefaults() *TooManyRequestsErrorMeta`

NewTooManyRequestsErrorMetaWithDefaults instantiates a new TooManyRequestsErrorMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *TooManyRequestsErrorMeta) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TooManyRequestsErrorMeta) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TooManyRequestsErrorMeta) SetLimit(v int64)`

SetLimit sets Limit field to given value.


### GetPeriod

`func (o *TooManyRequestsErrorMeta) GetPeriod() int64`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *TooManyRequestsErrorMeta) GetPeriodOk() (*int64, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *TooManyRequestsErrorMeta) SetPeriod(v int64)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *TooManyRequestsErrorMeta) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetRemaining

`func (o *TooManyRequestsErrorMeta) GetRemaining() int64`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *TooManyRequestsErrorMeta) GetRemainingOk() (*int64, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *TooManyRequestsErrorMeta) SetRemaining(v int64)`

SetRemaining sets Remaining field to given value.


### GetReset

`func (o *TooManyRequestsErrorMeta) GetReset() string`

GetReset returns the Reset field if non-nil, zero value otherwise.

### GetResetOk

`func (o *TooManyRequestsErrorMeta) GetResetOk() (*string, bool)`

GetResetOk returns a tuple with the Reset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReset

`func (o *TooManyRequestsErrorMeta) SetReset(v string)`

SetReset sets Reset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


