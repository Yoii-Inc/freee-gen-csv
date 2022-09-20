# CompanyResponseCompanyWalletablesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 口座ID | 
**Name** | **string** | 口座名 (255文字以内) | 
**Type** | **string** | 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet) | 

## Methods

### NewCompanyResponseCompanyWalletablesInner

`func NewCompanyResponseCompanyWalletablesInner(id int32, name string, type_ string, ) *CompanyResponseCompanyWalletablesInner`

NewCompanyResponseCompanyWalletablesInner instantiates a new CompanyResponseCompanyWalletablesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyResponseCompanyWalletablesInnerWithDefaults

`func NewCompanyResponseCompanyWalletablesInnerWithDefaults() *CompanyResponseCompanyWalletablesInner`

NewCompanyResponseCompanyWalletablesInnerWithDefaults instantiates a new CompanyResponseCompanyWalletablesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompanyResponseCompanyWalletablesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyResponseCompanyWalletablesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyResponseCompanyWalletablesInner) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *CompanyResponseCompanyWalletablesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyResponseCompanyWalletablesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyResponseCompanyWalletablesInner) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CompanyResponseCompanyWalletablesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CompanyResponseCompanyWalletablesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CompanyResponseCompanyWalletablesInner) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


