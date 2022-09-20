# DealRenewsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | [**[]DealRenewsInnerDetailsInner**](DealRenewsInnerDetailsInner.md) | +更新の明細行一覧（配列） | 
**Id** | **int64** | +更新行ID | 
**RenewTargetId** | **int64** | +更新の対象行ID | 
**RenewTargetType** | **string** | +更新の対象行タイプ | 
**UpdateDate** | **string** | 更新日 (yyyy-mm-dd) | 

## Methods

### NewDealRenewsInner

`func NewDealRenewsInner(details []DealRenewsInnerDetailsInner, id int64, renewTargetId int64, renewTargetType string, updateDate string, ) *DealRenewsInner`

NewDealRenewsInner instantiates a new DealRenewsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealRenewsInnerWithDefaults

`func NewDealRenewsInnerWithDefaults() *DealRenewsInner`

NewDealRenewsInnerWithDefaults instantiates a new DealRenewsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *DealRenewsInner) GetDetails() []DealRenewsInnerDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *DealRenewsInner) GetDetailsOk() (*[]DealRenewsInnerDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *DealRenewsInner) SetDetails(v []DealRenewsInnerDetailsInner)`

SetDetails sets Details field to given value.


### GetId

`func (o *DealRenewsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DealRenewsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DealRenewsInner) SetId(v int64)`

SetId sets Id field to given value.


### GetRenewTargetId

`func (o *DealRenewsInner) GetRenewTargetId() int64`

GetRenewTargetId returns the RenewTargetId field if non-nil, zero value otherwise.

### GetRenewTargetIdOk

`func (o *DealRenewsInner) GetRenewTargetIdOk() (*int64, bool)`

GetRenewTargetIdOk returns a tuple with the RenewTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewTargetId

`func (o *DealRenewsInner) SetRenewTargetId(v int64)`

SetRenewTargetId sets RenewTargetId field to given value.


### GetRenewTargetType

`func (o *DealRenewsInner) GetRenewTargetType() string`

GetRenewTargetType returns the RenewTargetType field if non-nil, zero value otherwise.

### GetRenewTargetTypeOk

`func (o *DealRenewsInner) GetRenewTargetTypeOk() (*string, bool)`

GetRenewTargetTypeOk returns a tuple with the RenewTargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewTargetType

`func (o *DealRenewsInner) SetRenewTargetType(v string)`

SetRenewTargetType sets RenewTargetType field to given value.


### GetUpdateDate

`func (o *DealRenewsInner) GetUpdateDate() string`

GetUpdateDate returns the UpdateDate field if non-nil, zero value otherwise.

### GetUpdateDateOk

`func (o *DealRenewsInner) GetUpdateDateOk() (*string, bool)`

GetUpdateDateOk returns a tuple with the UpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDate

`func (o *DealRenewsInner) SetUpdateDate(v string)`

SetUpdateDate sets UpdateDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


