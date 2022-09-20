# GetWalletables200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**GetWalletables200ResponseMeta**](GetWalletables200ResponseMeta.md) |  | [optional] 
**Walletables** | [**[]Walletable**](Walletable.md) |  | 

## Methods

### NewGetWalletables200Response

`func NewGetWalletables200Response(walletables []Walletable, ) *GetWalletables200Response`

NewGetWalletables200Response instantiates a new GetWalletables200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWalletables200ResponseWithDefaults

`func NewGetWalletables200ResponseWithDefaults() *GetWalletables200Response`

NewGetWalletables200ResponseWithDefaults instantiates a new GetWalletables200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetWalletables200Response) GetMeta() GetWalletables200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetWalletables200Response) GetMetaOk() (*GetWalletables200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetWalletables200Response) SetMeta(v GetWalletables200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetWalletables200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetWalletables

`func (o *GetWalletables200Response) GetWalletables() []Walletable`

GetWalletables returns the Walletables field if non-nil, zero value otherwise.

### GetWalletablesOk

`func (o *GetWalletables200Response) GetWalletablesOk() (*[]Walletable, bool)`

GetWalletablesOk returns a tuple with the Walletables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletables

`func (o *GetWalletables200Response) SetWalletables(v []Walletable)`

SetWalletables sets Walletables field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


