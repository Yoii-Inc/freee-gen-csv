# DealReceiptsInnerUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | 表示名 | [optional] 
**Email** | **string** | メールアドレス | 
**Id** | **int32** | ユーザーID | 

## Methods

### NewDealReceiptsInnerUser

`func NewDealReceiptsInnerUser(email string, id int32, ) *DealReceiptsInnerUser`

NewDealReceiptsInnerUser instantiates a new DealReceiptsInnerUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealReceiptsInnerUserWithDefaults

`func NewDealReceiptsInnerUserWithDefaults() *DealReceiptsInnerUser`

NewDealReceiptsInnerUserWithDefaults instantiates a new DealReceiptsInnerUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *DealReceiptsInnerUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DealReceiptsInnerUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DealReceiptsInnerUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DealReceiptsInnerUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *DealReceiptsInnerUser) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *DealReceiptsInnerUser) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEmail

`func (o *DealReceiptsInnerUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DealReceiptsInnerUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DealReceiptsInnerUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetId

`func (o *DealReceiptsInnerUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DealReceiptsInnerUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DealReceiptsInnerUser) SetId(v int32)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


