# DealResponseDealRenews

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | [**[]DealResponseDealDetails**](DealResponseDealDetails.md) | +更新の明細行一覧（配列） | 
**Id** | **int64** | +更新行ID | 
**RenewTargetId** | **int64** | +更新の対象行ID | 
**RenewTargetType** | **string** | +更新の対象行タイプ | 
**UpdateDate** | **string** | 更新日 (yyyy-mm-dd) | 

## Methods

### NewDealResponseDealRenews

`func NewDealResponseDealRenews(details []DealResponseDealDetails, id int64, renewTargetId int64, renewTargetType string, updateDate string, ) *DealResponseDealRenews`

NewDealResponseDealRenews instantiates a new DealResponseDealRenews object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealResponseDealRenewsWithDefaults

`func NewDealResponseDealRenewsWithDefaults() *DealResponseDealRenews`

NewDealResponseDealRenewsWithDefaults instantiates a new DealResponseDealRenews object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *DealResponseDealRenews) GetDetails() []DealResponseDealDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *DealResponseDealRenews) GetDetailsOk() (*[]DealResponseDealDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *DealResponseDealRenews) SetDetails(v []DealResponseDealDetails)`

SetDetails sets Details field to given value.


### GetId

`func (o *DealResponseDealRenews) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DealResponseDealRenews) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DealResponseDealRenews) SetId(v int64)`

SetId sets Id field to given value.


### GetRenewTargetId

`func (o *DealResponseDealRenews) GetRenewTargetId() int64`

GetRenewTargetId returns the RenewTargetId field if non-nil, zero value otherwise.

### GetRenewTargetIdOk

`func (o *DealResponseDealRenews) GetRenewTargetIdOk() (*int64, bool)`

GetRenewTargetIdOk returns a tuple with the RenewTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewTargetId

`func (o *DealResponseDealRenews) SetRenewTargetId(v int64)`

SetRenewTargetId sets RenewTargetId field to given value.


### GetRenewTargetType

`func (o *DealResponseDealRenews) GetRenewTargetType() string`

GetRenewTargetType returns the RenewTargetType field if non-nil, zero value otherwise.

### GetRenewTargetTypeOk

`func (o *DealResponseDealRenews) GetRenewTargetTypeOk() (*string, bool)`

GetRenewTargetTypeOk returns a tuple with the RenewTargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewTargetType

`func (o *DealResponseDealRenews) SetRenewTargetType(v string)`

SetRenewTargetType sets RenewTargetType field to given value.


### GetUpdateDate

`func (o *DealResponseDealRenews) GetUpdateDate() string`

GetUpdateDate returns the UpdateDate field if non-nil, zero value otherwise.

### GetUpdateDateOk

`func (o *DealResponseDealRenews) GetUpdateDateOk() (*string, bool)`

GetUpdateDateOk returns a tuple with the UpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDate

`func (o *DealResponseDealRenews) SetUpdateDate(v string)`

SetUpdateDate sets UpdateDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


