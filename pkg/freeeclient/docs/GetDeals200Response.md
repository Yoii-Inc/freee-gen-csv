# GetDeals200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deals** | [**[]Deal**](Deal.md) |  | 
**Meta** | [**GetDeals200ResponseMeta**](GetDeals200ResponseMeta.md) |  | 

## Methods

### NewGetDeals200Response

`func NewGetDeals200Response(deals []Deal, meta GetDeals200ResponseMeta, ) *GetDeals200Response`

NewGetDeals200Response instantiates a new GetDeals200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeals200ResponseWithDefaults

`func NewGetDeals200ResponseWithDefaults() *GetDeals200Response`

NewGetDeals200ResponseWithDefaults instantiates a new GetDeals200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeals

`func (o *GetDeals200Response) GetDeals() []Deal`

GetDeals returns the Deals field if non-nil, zero value otherwise.

### GetDealsOk

`func (o *GetDeals200Response) GetDealsOk() (*[]Deal, bool)`

GetDealsOk returns a tuple with the Deals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeals

`func (o *GetDeals200Response) SetDeals(v []Deal)`

SetDeals sets Deals field to given value.


### GetMeta

`func (o *GetDeals200Response) GetMeta() GetDeals200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDeals200Response) GetMetaOk() (*GetDeals200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDeals200Response) SetMeta(v GetDeals200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


