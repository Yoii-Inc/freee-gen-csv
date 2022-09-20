# RenewUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int64** | 事業所ID | 
**Details** | [**[]RenewUpdateParamsDetailsInner**](RenewUpdateParamsDetailsInner.md) | +更新の明細行 | 
**UpdateDate** | **string** | 更新日 (yyyy-mm-dd) | 

## Methods

### NewRenewUpdateParams

`func NewRenewUpdateParams(companyId int64, details []RenewUpdateParamsDetailsInner, updateDate string, ) *RenewUpdateParams`

NewRenewUpdateParams instantiates a new RenewUpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenewUpdateParamsWithDefaults

`func NewRenewUpdateParamsWithDefaults() *RenewUpdateParams`

NewRenewUpdateParamsWithDefaults instantiates a new RenewUpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *RenewUpdateParams) GetCompanyId() int64`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *RenewUpdateParams) GetCompanyIdOk() (*int64, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *RenewUpdateParams) SetCompanyId(v int64)`

SetCompanyId sets CompanyId field to given value.


### GetDetails

`func (o *RenewUpdateParams) GetDetails() []RenewUpdateParamsDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *RenewUpdateParams) GetDetailsOk() (*[]RenewUpdateParamsDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *RenewUpdateParams) SetDetails(v []RenewUpdateParamsDetailsInner)`

SetDetails sets Details field to given value.


### GetUpdateDate

`func (o *RenewUpdateParams) GetUpdateDate() string`

GetUpdateDate returns the UpdateDate field if non-nil, zero value otherwise.

### GetUpdateDateOk

`func (o *RenewUpdateParams) GetUpdateDateOk() (*string, bool)`

GetUpdateDateOk returns a tuple with the UpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDate

`func (o *RenewUpdateParams) SetUpdateDate(v string)`

SetUpdateDate sets UpdateDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


