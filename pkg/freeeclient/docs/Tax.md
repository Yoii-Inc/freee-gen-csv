# Tax

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int64** | 税区分コード | 
**Name** | **string** | 税区分名 | 
**NameJa** | **string** | 税区分名（日本語表示用） | 

## Methods

### NewTax

`func NewTax(code int64, name string, nameJa string, ) *Tax`

NewTax instantiates a new Tax object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxWithDefaults

`func NewTaxWithDefaults() *Tax`

NewTaxWithDefaults instantiates a new Tax object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Tax) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Tax) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Tax) SetCode(v int64)`

SetCode sets Code field to given value.


### GetName

`func (o *Tax) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tax) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tax) SetName(v string)`

SetName sets Name field to given value.


### GetNameJa

`func (o *Tax) GetNameJa() string`

GetNameJa returns the NameJa field if non-nil, zero value otherwise.

### GetNameJaOk

`func (o *Tax) GetNameJaOk() (*string, bool)`

GetNameJaOk returns a tuple with the NameJa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameJa

`func (o *Tax) SetNameJa(v string)`

SetNameJa sets NameJa field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


