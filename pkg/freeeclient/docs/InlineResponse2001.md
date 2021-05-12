# InlineResponse2001

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deals** | [**[]Deal**](Deal.md) |  | 
**Meta** | [**InlineResponse2001Meta**](InlineResponse2001Meta.md) |  | 

## Methods

### NewInlineResponse2001

`func NewInlineResponse2001(deals []Deal, meta InlineResponse2001Meta, ) *InlineResponse2001`

NewInlineResponse2001 instantiates a new InlineResponse2001 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2001WithDefaults

`func NewInlineResponse2001WithDefaults() *InlineResponse2001`

NewInlineResponse2001WithDefaults instantiates a new InlineResponse2001 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeals

`func (o *InlineResponse2001) GetDeals() []Deal`

GetDeals returns the Deals field if non-nil, zero value otherwise.

### GetDealsOk

`func (o *InlineResponse2001) GetDealsOk() (*[]Deal, bool)`

GetDealsOk returns a tuple with the Deals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeals

`func (o *InlineResponse2001) SetDeals(v []Deal)`

SetDeals sets Deals field to given value.


### GetMeta

`func (o *InlineResponse2001) GetMeta() InlineResponse2001Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse2001) GetMetaOk() (*InlineResponse2001Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse2001) SetMeta(v InlineResponse2001Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


