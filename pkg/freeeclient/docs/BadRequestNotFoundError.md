# BadRequestNotFoundError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]BadRequestNotFoundErrorErrorsInner**](BadRequestNotFoundErrorErrorsInner.md) |  | [optional] 
**StatusCode** | Pointer to **int64** |  | [optional] 

## Methods

### NewBadRequestNotFoundError

`func NewBadRequestNotFoundError() *BadRequestNotFoundError`

NewBadRequestNotFoundError instantiates a new BadRequestNotFoundError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBadRequestNotFoundErrorWithDefaults

`func NewBadRequestNotFoundErrorWithDefaults() *BadRequestNotFoundError`

NewBadRequestNotFoundErrorWithDefaults instantiates a new BadRequestNotFoundError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *BadRequestNotFoundError) GetErrors() []BadRequestNotFoundErrorErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BadRequestNotFoundError) GetErrorsOk() (*[]BadRequestNotFoundErrorErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BadRequestNotFoundError) SetErrors(v []BadRequestNotFoundErrorErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BadRequestNotFoundError) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetStatusCode

`func (o *BadRequestNotFoundError) GetStatusCode() int64`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *BadRequestNotFoundError) GetStatusCodeOk() (*int64, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *BadRequestNotFoundError) SetStatusCode(v int64)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *BadRequestNotFoundError) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


