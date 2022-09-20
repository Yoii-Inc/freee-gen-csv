# Walletable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankId** | **NullableInt32** | サービスID | 
**Id** | **int32** | 口座ID | 
**LastBalance** | Pointer to **int32** | 同期残高 | [optional] 
**Name** | **string** | 口座名 (255文字以内) | 
**Type** | **string** | 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet) | 
**WalletableBalance** | Pointer to **int32** | 登録残高 | [optional] 

## Methods

### NewWalletable

`func NewWalletable(bankId NullableInt32, id int32, name string, type_ string, ) *Walletable`

NewWalletable instantiates a new Walletable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletableWithDefaults

`func NewWalletableWithDefaults() *Walletable`

NewWalletableWithDefaults instantiates a new Walletable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankId

`func (o *Walletable) GetBankId() int32`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *Walletable) GetBankIdOk() (*int32, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *Walletable) SetBankId(v int32)`

SetBankId sets BankId field to given value.


### SetBankIdNil

`func (o *Walletable) SetBankIdNil(b bool)`

 SetBankIdNil sets the value for BankId to be an explicit nil

### UnsetBankId
`func (o *Walletable) UnsetBankId()`

UnsetBankId ensures that no value is present for BankId, not even an explicit nil
### GetId

`func (o *Walletable) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Walletable) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Walletable) SetId(v int32)`

SetId sets Id field to given value.


### GetLastBalance

`func (o *Walletable) GetLastBalance() int32`

GetLastBalance returns the LastBalance field if non-nil, zero value otherwise.

### GetLastBalanceOk

`func (o *Walletable) GetLastBalanceOk() (*int32, bool)`

GetLastBalanceOk returns a tuple with the LastBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBalance

`func (o *Walletable) SetLastBalance(v int32)`

SetLastBalance sets LastBalance field to given value.

### HasLastBalance

`func (o *Walletable) HasLastBalance() bool`

HasLastBalance returns a boolean if a field has been set.

### GetName

`func (o *Walletable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Walletable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Walletable) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Walletable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Walletable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Walletable) SetType(v string)`

SetType sets Type field to given value.


### GetWalletableBalance

`func (o *Walletable) GetWalletableBalance() int32`

GetWalletableBalance returns the WalletableBalance field if non-nil, zero value otherwise.

### GetWalletableBalanceOk

`func (o *Walletable) GetWalletableBalanceOk() (*int32, bool)`

GetWalletableBalanceOk returns a tuple with the WalletableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletableBalance

`func (o *Walletable) SetWalletableBalance(v int32)`

SetWalletableBalance sets WalletableBalance field to given value.

### HasWalletableBalance

`func (o *Walletable) HasWalletableBalance() bool`

HasWalletableBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


