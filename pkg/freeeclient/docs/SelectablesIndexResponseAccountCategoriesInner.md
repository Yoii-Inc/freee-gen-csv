# SelectablesIndexResponseAccountCategoriesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountItems** | [**[]SelectablesIndexResponseAccountCategoriesInnerAccountItemsInner**](SelectablesIndexResponseAccountCategoriesInnerAccountItemsInner.md) | 勘定科目の一覧 | 
**Balance** | **string** | 収支 | 
**Desc** | Pointer to **string** | カテゴリーの説明 | [optional] 
**OrgCode** | **string** | 事業形態（個人事業主: personal、法人: corporate） | 
**Role** | **string** | カテゴリーコード | 
**Title** | **string** | カテゴリー名 | 

## Methods

### NewSelectablesIndexResponseAccountCategoriesInner

`func NewSelectablesIndexResponseAccountCategoriesInner(accountItems []SelectablesIndexResponseAccountCategoriesInnerAccountItemsInner, balance string, orgCode string, role string, title string, ) *SelectablesIndexResponseAccountCategoriesInner`

NewSelectablesIndexResponseAccountCategoriesInner instantiates a new SelectablesIndexResponseAccountCategoriesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectablesIndexResponseAccountCategoriesInnerWithDefaults

`func NewSelectablesIndexResponseAccountCategoriesInnerWithDefaults() *SelectablesIndexResponseAccountCategoriesInner`

NewSelectablesIndexResponseAccountCategoriesInnerWithDefaults instantiates a new SelectablesIndexResponseAccountCategoriesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountItems

`func (o *SelectablesIndexResponseAccountCategoriesInner) GetAccountItems() []SelectablesIndexResponseAccountCategoriesInnerAccountItemsInner`

GetAccountItems returns the AccountItems field if non-nil, zero value otherwise.

### GetAccountItemsOk

`func (o *SelectablesIndexResponseAccountCategoriesInner) GetAccountItemsOk() (*[]SelectablesIndexResponseAccountCategoriesInnerAccountItemsInner, bool)`

GetAccountItemsOk returns a tuple with the AccountItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItems

`func (o *SelectablesIndexResponseAccountCategoriesInner) SetAccountItems(v []SelectablesIndexResponseAccountCategoriesInnerAccountItemsInner)`

SetAccountItems sets AccountItems field to given value.


### GetBalance

`func (o *SelectablesIndexResponseAccountCategoriesInner) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *SelectablesIndexResponseAccountCategoriesInner) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *SelectablesIndexResponseAccountCategoriesInner) SetBalance(v string)`

SetBalance sets Balance field to given value.


### GetDesc

`func (o *SelectablesIndexResponseAccountCategoriesInner) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *SelectablesIndexResponseAccountCategoriesInner) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *SelectablesIndexResponseAccountCategoriesInner) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *SelectablesIndexResponseAccountCategoriesInner) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetOrgCode

`func (o *SelectablesIndexResponseAccountCategoriesInner) GetOrgCode() string`

GetOrgCode returns the OrgCode field if non-nil, zero value otherwise.

### GetOrgCodeOk

`func (o *SelectablesIndexResponseAccountCategoriesInner) GetOrgCodeOk() (*string, bool)`

GetOrgCodeOk returns a tuple with the OrgCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgCode

`func (o *SelectablesIndexResponseAccountCategoriesInner) SetOrgCode(v string)`

SetOrgCode sets OrgCode field to given value.


### GetRole

`func (o *SelectablesIndexResponseAccountCategoriesInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SelectablesIndexResponseAccountCategoriesInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SelectablesIndexResponseAccountCategoriesInner) SetRole(v string)`

SetRole sets Role field to given value.


### GetTitle

`func (o *SelectablesIndexResponseAccountCategoriesInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SelectablesIndexResponseAccountCategoriesInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SelectablesIndexResponseAccountCategoriesInner) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


