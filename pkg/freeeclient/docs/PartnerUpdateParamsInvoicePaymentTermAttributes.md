# PartnerUpdateParamsInvoicePaymentTermAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalMonths** | Pointer to **int64** | 入金月 | [optional] 
**CutoffDay** | Pointer to **int64** | 締め日（29, 30, 31日の末日を指定する場合は、32を指定してください。） | [optional] 
**FixedDay** | Pointer to **int64** | 入金日（29, 30, 31日の末日を指定する場合は、32を指定してください。） | [optional] 

## Methods

### NewPartnerUpdateParamsInvoicePaymentTermAttributes

`func NewPartnerUpdateParamsInvoicePaymentTermAttributes() *PartnerUpdateParamsInvoicePaymentTermAttributes`

NewPartnerUpdateParamsInvoicePaymentTermAttributes instantiates a new PartnerUpdateParamsInvoicePaymentTermAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerUpdateParamsInvoicePaymentTermAttributesWithDefaults

`func NewPartnerUpdateParamsInvoicePaymentTermAttributesWithDefaults() *PartnerUpdateParamsInvoicePaymentTermAttributes`

NewPartnerUpdateParamsInvoicePaymentTermAttributesWithDefaults instantiates a new PartnerUpdateParamsInvoicePaymentTermAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalMonths

`func (o *PartnerUpdateParamsInvoicePaymentTermAttributes) GetAdditionalMonths() int64`

GetAdditionalMonths returns the AdditionalMonths field if non-nil, zero value otherwise.

### GetAdditionalMonthsOk

`func (o *PartnerUpdateParamsInvoicePaymentTermAttributes) GetAdditionalMonthsOk() (*int64, bool)`

GetAdditionalMonthsOk returns a tuple with the AdditionalMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalMonths

`func (o *PartnerUpdateParamsInvoicePaymentTermAttributes) SetAdditionalMonths(v int64)`

SetAdditionalMonths sets AdditionalMonths field to given value.

### HasAdditionalMonths

`func (o *PartnerUpdateParamsInvoicePaymentTermAttributes) HasAdditionalMonths() bool`

HasAdditionalMonths returns a boolean if a field has been set.

### GetCutoffDay

`func (o *PartnerUpdateParamsInvoicePaymentTermAttributes) GetCutoffDay() int64`

GetCutoffDay returns the CutoffDay field if non-nil, zero value otherwise.

### GetCutoffDayOk

`func (o *PartnerUpdateParamsInvoicePaymentTermAttributes) GetCutoffDayOk() (*int64, bool)`

GetCutoffDayOk returns a tuple with the CutoffDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCutoffDay

`func (o *PartnerUpdateParamsInvoicePaymentTermAttributes) SetCutoffDay(v int64)`

SetCutoffDay sets CutoffDay field to given value.

### HasCutoffDay

`func (o *PartnerUpdateParamsInvoicePaymentTermAttributes) HasCutoffDay() bool`

HasCutoffDay returns a boolean if a field has been set.

### GetFixedDay

`func (o *PartnerUpdateParamsInvoicePaymentTermAttributes) GetFixedDay() int64`

GetFixedDay returns the FixedDay field if non-nil, zero value otherwise.

### GetFixedDayOk

`func (o *PartnerUpdateParamsInvoicePaymentTermAttributes) GetFixedDayOk() (*int64, bool)`

GetFixedDayOk returns a tuple with the FixedDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedDay

`func (o *PartnerUpdateParamsInvoicePaymentTermAttributes) SetFixedDay(v int64)`

SetFixedDay sets FixedDay field to given value.

### HasFixedDay

`func (o *PartnerUpdateParamsInvoicePaymentTermAttributes) HasFixedDay() bool`

HasFixedDay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


