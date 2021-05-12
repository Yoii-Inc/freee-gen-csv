# InlineResponse20010Taxes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **bool** | true: 使用する、false: 使用しない | 
**Code** | **int32** | 税区分コード | 
**DisplayCategory** | **NullableString** | 税区分の表示カテゴリ（tax_5: 5%表示の税区分、tax_8: 8%表示の税区分、tax_r8: 軽減税率8%表示の税区分、tax_10: 10%表示の税区分、null: 税率未設定税区分） | 
**Name** | **string** | 税区分名 | 
**NameJa** | **string** | 税区分名（日本語表示用） | 

## Methods

### NewInlineResponse20010Taxes

`func NewInlineResponse20010Taxes(available bool, code int32, displayCategory NullableString, name string, nameJa string, ) *InlineResponse20010Taxes`

NewInlineResponse20010Taxes instantiates a new InlineResponse20010Taxes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20010TaxesWithDefaults

`func NewInlineResponse20010TaxesWithDefaults() *InlineResponse20010Taxes`

NewInlineResponse20010TaxesWithDefaults instantiates a new InlineResponse20010Taxes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *InlineResponse20010Taxes) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *InlineResponse20010Taxes) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *InlineResponse20010Taxes) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetCode

`func (o *InlineResponse20010Taxes) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InlineResponse20010Taxes) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InlineResponse20010Taxes) SetCode(v int32)`

SetCode sets Code field to given value.


### GetDisplayCategory

`func (o *InlineResponse20010Taxes) GetDisplayCategory() string`

GetDisplayCategory returns the DisplayCategory field if non-nil, zero value otherwise.

### GetDisplayCategoryOk

`func (o *InlineResponse20010Taxes) GetDisplayCategoryOk() (*string, bool)`

GetDisplayCategoryOk returns a tuple with the DisplayCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayCategory

`func (o *InlineResponse20010Taxes) SetDisplayCategory(v string)`

SetDisplayCategory sets DisplayCategory field to given value.


### SetDisplayCategoryNil

`func (o *InlineResponse20010Taxes) SetDisplayCategoryNil(b bool)`

 SetDisplayCategoryNil sets the value for DisplayCategory to be an explicit nil

### UnsetDisplayCategory
`func (o *InlineResponse20010Taxes) UnsetDisplayCategory()`

UnsetDisplayCategory ensures that no value is present for DisplayCategory, not even an explicit nil
### GetName

`func (o *InlineResponse20010Taxes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20010Taxes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20010Taxes) SetName(v string)`

SetName sets Name field to given value.


### GetNameJa

`func (o *InlineResponse20010Taxes) GetNameJa() string`

GetNameJa returns the NameJa field if non-nil, zero value otherwise.

### GetNameJaOk

`func (o *InlineResponse20010Taxes) GetNameJaOk() (*string, bool)`

GetNameJaOk returns a tuple with the NameJa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameJa

`func (o *InlineResponse20010Taxes) SetNameJa(v string)`

SetNameJa sets NameJa field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


