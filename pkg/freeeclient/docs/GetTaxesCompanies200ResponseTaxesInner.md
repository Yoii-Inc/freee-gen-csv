# GetTaxesCompanies200ResponseTaxesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **bool** | true: 使用する、false: 使用しない | 
**Code** | **int32** | 税区分コード | 
**DisplayCategory** | **NullableString** | 税区分の表示カテゴリ（tax_5: 5%表示の税区分、tax_8: 8%表示の税区分、tax_r8: 軽減税率8%表示の税区分、tax_10: 10%表示の税区分、null: 税率未設定税区分） | 
**Name** | **string** | 税区分名 | 
**NameJa** | **string** | 税区分名（日本語表示用） | 

## Methods

### NewGetTaxesCompanies200ResponseTaxesInner

`func NewGetTaxesCompanies200ResponseTaxesInner(available bool, code int32, displayCategory NullableString, name string, nameJa string, ) *GetTaxesCompanies200ResponseTaxesInner`

NewGetTaxesCompanies200ResponseTaxesInner instantiates a new GetTaxesCompanies200ResponseTaxesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTaxesCompanies200ResponseTaxesInnerWithDefaults

`func NewGetTaxesCompanies200ResponseTaxesInnerWithDefaults() *GetTaxesCompanies200ResponseTaxesInner`

NewGetTaxesCompanies200ResponseTaxesInnerWithDefaults instantiates a new GetTaxesCompanies200ResponseTaxesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *GetTaxesCompanies200ResponseTaxesInner) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *GetTaxesCompanies200ResponseTaxesInner) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *GetTaxesCompanies200ResponseTaxesInner) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetCode

`func (o *GetTaxesCompanies200ResponseTaxesInner) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetTaxesCompanies200ResponseTaxesInner) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetTaxesCompanies200ResponseTaxesInner) SetCode(v int32)`

SetCode sets Code field to given value.


### GetDisplayCategory

`func (o *GetTaxesCompanies200ResponseTaxesInner) GetDisplayCategory() string`

GetDisplayCategory returns the DisplayCategory field if non-nil, zero value otherwise.

### GetDisplayCategoryOk

`func (o *GetTaxesCompanies200ResponseTaxesInner) GetDisplayCategoryOk() (*string, bool)`

GetDisplayCategoryOk returns a tuple with the DisplayCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayCategory

`func (o *GetTaxesCompanies200ResponseTaxesInner) SetDisplayCategory(v string)`

SetDisplayCategory sets DisplayCategory field to given value.


### SetDisplayCategoryNil

`func (o *GetTaxesCompanies200ResponseTaxesInner) SetDisplayCategoryNil(b bool)`

 SetDisplayCategoryNil sets the value for DisplayCategory to be an explicit nil

### UnsetDisplayCategory
`func (o *GetTaxesCompanies200ResponseTaxesInner) UnsetDisplayCategory()`

UnsetDisplayCategory ensures that no value is present for DisplayCategory, not even an explicit nil
### GetName

`func (o *GetTaxesCompanies200ResponseTaxesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTaxesCompanies200ResponseTaxesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTaxesCompanies200ResponseTaxesInner) SetName(v string)`

SetName sets Name field to given value.


### GetNameJa

`func (o *GetTaxesCompanies200ResponseTaxesInner) GetNameJa() string`

GetNameJa returns the NameJa field if non-nil, zero value otherwise.

### GetNameJaOk

`func (o *GetTaxesCompanies200ResponseTaxesInner) GetNameJaOk() (*string, bool)`

GetNameJaOk returns a tuple with the NameJa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameJa

`func (o *GetTaxesCompanies200ResponseTaxesInner) SetNameJa(v string)`

SetNameJa sets NameJa field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


