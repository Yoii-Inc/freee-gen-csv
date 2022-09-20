# WalletTxn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | 取引金額 | 
**Balance** | **int64** | 残高(銀行口座等) | 
**CompanyId** | **int64** | 事業所ID | 
**Date** | **string** | 取引日(yyyy-mm-dd) | 
**Description** | **string** | 取引内容 | 
**DueAmount** | **int64** | 未決済金額 | 
**EntrySide** | **string** | 入金／出金 (入金: income, 出金: expense) | 
**Id** | **int64** | 明細ID | 
**RuleMatched** | **bool** | 登録時に&lt;a href&#x3D;\&quot;https://support.freee.co.jp/hc/ja/articles/202848350-明細の自動登録ルールを設定する\&quot; target&#x3D;\&quot;_blank\&quot;&gt;自動登録ルールの設定&lt;/a&gt;が適用され、登録処理が実行された場合、 trueになります。〜を推測する、〜の消込をするの条件の場合は一致してもfalseになります。  | 
**Status** | **int64** | 明細のステータス（消込待ち: 1, 消込済み: 2, 無視: 3, 消込中: 4, 対象外: 6） | 
**WalletableId** | **int64** | 口座ID | 
**WalletableType** | **string** | 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet) | 

## Methods

### NewWalletTxn

`func NewWalletTxn(amount int64, balance int64, companyId int64, date string, description string, dueAmount int64, entrySide string, id int64, ruleMatched bool, status int64, walletableId int64, walletableType string, ) *WalletTxn`

NewWalletTxn instantiates a new WalletTxn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletTxnWithDefaults

`func NewWalletTxnWithDefaults() *WalletTxn`

NewWalletTxnWithDefaults instantiates a new WalletTxn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *WalletTxn) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WalletTxn) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WalletTxn) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetBalance

`func (o *WalletTxn) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *WalletTxn) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *WalletTxn) SetBalance(v int64)`

SetBalance sets Balance field to given value.


### GetCompanyId

`func (o *WalletTxn) GetCompanyId() int64`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *WalletTxn) GetCompanyIdOk() (*int64, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *WalletTxn) SetCompanyId(v int64)`

SetCompanyId sets CompanyId field to given value.


### GetDate

`func (o *WalletTxn) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *WalletTxn) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *WalletTxn) SetDate(v string)`

SetDate sets Date field to given value.


### GetDescription

`func (o *WalletTxn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WalletTxn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WalletTxn) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDueAmount

`func (o *WalletTxn) GetDueAmount() int64`

GetDueAmount returns the DueAmount field if non-nil, zero value otherwise.

### GetDueAmountOk

`func (o *WalletTxn) GetDueAmountOk() (*int64, bool)`

GetDueAmountOk returns a tuple with the DueAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueAmount

`func (o *WalletTxn) SetDueAmount(v int64)`

SetDueAmount sets DueAmount field to given value.


### GetEntrySide

`func (o *WalletTxn) GetEntrySide() string`

GetEntrySide returns the EntrySide field if non-nil, zero value otherwise.

### GetEntrySideOk

`func (o *WalletTxn) GetEntrySideOk() (*string, bool)`

GetEntrySideOk returns a tuple with the EntrySide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrySide

`func (o *WalletTxn) SetEntrySide(v string)`

SetEntrySide sets EntrySide field to given value.


### GetId

`func (o *WalletTxn) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WalletTxn) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WalletTxn) SetId(v int64)`

SetId sets Id field to given value.


### GetRuleMatched

`func (o *WalletTxn) GetRuleMatched() bool`

GetRuleMatched returns the RuleMatched field if non-nil, zero value otherwise.

### GetRuleMatchedOk

`func (o *WalletTxn) GetRuleMatchedOk() (*bool, bool)`

GetRuleMatchedOk returns a tuple with the RuleMatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleMatched

`func (o *WalletTxn) SetRuleMatched(v bool)`

SetRuleMatched sets RuleMatched field to given value.


### GetStatus

`func (o *WalletTxn) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WalletTxn) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WalletTxn) SetStatus(v int64)`

SetStatus sets Status field to given value.


### GetWalletableId

`func (o *WalletTxn) GetWalletableId() int64`

GetWalletableId returns the WalletableId field if non-nil, zero value otherwise.

### GetWalletableIdOk

`func (o *WalletTxn) GetWalletableIdOk() (*int64, bool)`

GetWalletableIdOk returns a tuple with the WalletableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletableId

`func (o *WalletTxn) SetWalletableId(v int64)`

SetWalletableId sets WalletableId field to given value.


### GetWalletableType

`func (o *WalletTxn) GetWalletableType() string`

GetWalletableType returns the WalletableType field if non-nil, zero value otherwise.

### GetWalletableTypeOk

`func (o *WalletTxn) GetWalletableTypeOk() (*string, bool)`

GetWalletableTypeOk returns a tuple with the WalletableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletableType

`func (o *WalletTxn) SetWalletableType(v string)`

SetWalletableType sets WalletableType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


