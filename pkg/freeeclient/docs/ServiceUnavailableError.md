# ServiceUnavailableError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]ServiceUnavailableErrorErrorsInner**](ServiceUnavailableErrorErrorsInner.md) |  | 
**StatusCode** | **int32** |  | 

## Methods

### NewServiceUnavailableError

`func NewServiceUnavailableError(errors []ServiceUnavailableErrorErrorsInner, statusCode int32, ) *ServiceUnavailableError`

NewServiceUnavailableError instantiates a new ServiceUnavailableError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceUnavailableErrorWithDefaults

`func NewServiceUnavailableErrorWithDefaults() *ServiceUnavailableError`

NewServiceUnavailableErrorWithDefaults instantiates a new ServiceUnavailableError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *ServiceUnavailableError) GetErrors() []ServiceUnavailableErrorErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ServiceUnavailableError) GetErrorsOk() (*[]ServiceUnavailableErrorErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ServiceUnavailableError) SetErrors(v []ServiceUnavailableErrorErrorsInner)`

SetErrors sets Errors field to given value.


### GetStatusCode

`func (o *ServiceUnavailableError) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ServiceUnavailableError) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ServiceUnavailableError) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


