# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | 表示名 | [optional] 
**Email** | **string** | メールアドレス | 
**FirstName** | Pointer to **NullableString** | 氏名（名） | [optional] 
**FirstNameKana** | Pointer to **NullableString** | 氏名（カナ・名） | [optional] 
**Id** | **int64** | ユーザーID | 
**LastName** | Pointer to **NullableString** | 氏名（姓） | [optional] 
**LastNameKana** | Pointer to **NullableString** | 氏名（カナ・姓） | [optional] 

## Methods

### NewUser

`func NewUser(email string, id int64, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *User) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *User) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *User) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *User) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *User) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *User) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *User) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *User) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *User) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *User) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *User) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *User) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetFirstNameKana

`func (o *User) GetFirstNameKana() string`

GetFirstNameKana returns the FirstNameKana field if non-nil, zero value otherwise.

### GetFirstNameKanaOk

`func (o *User) GetFirstNameKanaOk() (*string, bool)`

GetFirstNameKanaOk returns a tuple with the FirstNameKana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstNameKana

`func (o *User) SetFirstNameKana(v string)`

SetFirstNameKana sets FirstNameKana field to given value.

### HasFirstNameKana

`func (o *User) HasFirstNameKana() bool`

HasFirstNameKana returns a boolean if a field has been set.

### SetFirstNameKanaNil

`func (o *User) SetFirstNameKanaNil(b bool)`

 SetFirstNameKanaNil sets the value for FirstNameKana to be an explicit nil

### UnsetFirstNameKana
`func (o *User) UnsetFirstNameKana()`

UnsetFirstNameKana ensures that no value is present for FirstNameKana, not even an explicit nil
### GetId

`func (o *User) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v int64)`

SetId sets Id field to given value.


### GetLastName

`func (o *User) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *User) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *User) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *User) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *User) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *User) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetLastNameKana

`func (o *User) GetLastNameKana() string`

GetLastNameKana returns the LastNameKana field if non-nil, zero value otherwise.

### GetLastNameKanaOk

`func (o *User) GetLastNameKanaOk() (*string, bool)`

GetLastNameKanaOk returns a tuple with the LastNameKana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNameKana

`func (o *User) SetLastNameKana(v string)`

SetLastNameKana sets LastNameKana field to given value.

### HasLastNameKana

`func (o *User) HasLastNameKana() bool`

HasLastNameKana returns a boolean if a field has been set.

### SetLastNameKanaNil

`func (o *User) SetLastNameKanaNil(b bool)`

 SetLastNameKanaNil sets the value for LastNameKana to be an explicit nil

### UnsetLastNameKana
`func (o *User) UnsetLastNameKana()`

UnsetLastNameKana ensures that no value is present for LastNameKana, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


