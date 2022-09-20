# GetWalletable200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**GetWalletables200ResponseMeta**](GetWalletables200ResponseMeta.md) |  | [optional] 
**Walletable** | [**Walletable**](Walletable.md) |  | 

## Methods

### NewGetWalletable200Response

`func NewGetWalletable200Response(walletable Walletable, ) *GetWalletable200Response`

NewGetWalletable200Response instantiates a new GetWalletable200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWalletable200ResponseWithDefaults

`func NewGetWalletable200ResponseWithDefaults() *GetWalletable200Response`

NewGetWalletable200ResponseWithDefaults instantiates a new GetWalletable200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetWalletable200Response) GetMeta() GetWalletables200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetWalletable200Response) GetMetaOk() (*GetWalletables200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetWalletable200Response) SetMeta(v GetWalletables200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetWalletable200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetWalletable

`func (o *GetWalletable200Response) GetWalletable() Walletable`

GetWalletable returns the Walletable field if non-nil, zero value otherwise.

### GetWalletableOk

`func (o *GetWalletable200Response) GetWalletableOk() (*Walletable, bool)`

GetWalletableOk returns a tuple with the Walletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletable

`func (o *GetWalletable200Response) SetWalletable(v Walletable)`

SetWalletable sets Walletable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


