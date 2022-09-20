# GetTaxCodes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Taxes** | [**[]Tax**](Tax.md) |  | 

## Methods

### NewGetTaxCodes200Response

`func NewGetTaxCodes200Response(taxes []Tax, ) *GetTaxCodes200Response`

NewGetTaxCodes200Response instantiates a new GetTaxCodes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTaxCodes200ResponseWithDefaults

`func NewGetTaxCodes200ResponseWithDefaults() *GetTaxCodes200Response`

NewGetTaxCodes200ResponseWithDefaults instantiates a new GetTaxCodes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxes

`func (o *GetTaxCodes200Response) GetTaxes() []Tax`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *GetTaxCodes200Response) GetTaxesOk() (*[]Tax, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *GetTaxCodes200Response) SetTaxes(v []Tax)`

SetTaxes sets Taxes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


