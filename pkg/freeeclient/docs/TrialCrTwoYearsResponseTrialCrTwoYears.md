# TrialCrTwoYearsResponseTrialCrTwoYears

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountItemDisplayType** | Pointer to **string** | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）(条件に指定した時のみ含まれる） | [optional] 
**Adjustment** | Pointer to **string** | 決算整理仕訳のみ: only, 決算整理仕訳以外: without(条件に指定した時のみ含まれる） | [optional] 
**ApprovalFlowStatus** | Pointer to **string** | 未承認を除く: without_in_progress (デフォルト), 全てのステータス: all(条件に指定した時のみ含まれる） | [optional] 
**Balances** | [**[]TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner**](TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner.md) |  | 
**BreakdownDisplayType** | Pointer to **string** | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item, セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン): segment_1_tag, セグメント2(法人向け エンタープライズプラン):segment_2_tag, セグメント3(法人向け エンタープライズプラン): segment_3_tag）(条件に指定した時のみ含まれる） | [optional] 
**CompanyId** | **int64** | 事業所ID | 
**CostAllocation** | Pointer to **string** | 配賦仕訳のみ：only,配賦仕訳以外：without(条件に指定した時のみ含まれる） | [optional] 
**CreatedAt** | Pointer to **string** | 作成日時 | [optional] 
**EndDate** | Pointer to **string** | 発生日で絞込：終了日(yyyy-mm-dd)(条件に指定した時のみ含まれる） | [optional] 
**EndMonth** | Pointer to **int64** | 発生月で絞込：終了会計月(1-12)(条件に指定した時のみ含まれる） | [optional] 
**FiscalYear** | Pointer to **int64** | 会計年度(条件に指定した時、または条件に月、日条件がない時のみ含まれる） | [optional] 
**ItemId** | Pointer to **int64** | 品目ID(条件に指定した時のみ含まれる） | [optional] 
**PartnerCode** | Pointer to **string** | 取引先コード(条件に指定した時のみ含まれる） | [optional] 
**PartnerId** | Pointer to **int64** | 取引先ID(条件に指定した時のみ含まれる） | [optional] 
**SectionId** | Pointer to **int64** | 部門ID(条件に指定した時のみ含まれる） | [optional] 
**StartDate** | Pointer to **string** | 発生日で絞込：開始日(yyyy-mm-dd)(条件に指定した時のみ含まれる） | [optional] 
**StartMonth** | Pointer to **int64** | 発生月で絞込：開始会計月(1-12)(条件に指定した時のみ含まれる） | [optional] 

## Methods

### NewTrialCrTwoYearsResponseTrialCrTwoYears

`func NewTrialCrTwoYearsResponseTrialCrTwoYears(balances []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner, companyId int64, ) *TrialCrTwoYearsResponseTrialCrTwoYears`

NewTrialCrTwoYearsResponseTrialCrTwoYears instantiates a new TrialCrTwoYearsResponseTrialCrTwoYears object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialCrTwoYearsResponseTrialCrTwoYearsWithDefaults

`func NewTrialCrTwoYearsResponseTrialCrTwoYearsWithDefaults() *TrialCrTwoYearsResponseTrialCrTwoYears`

NewTrialCrTwoYearsResponseTrialCrTwoYearsWithDefaults instantiates a new TrialCrTwoYearsResponseTrialCrTwoYears object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountItemDisplayType

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetAccountItemDisplayType() string`

GetAccountItemDisplayType returns the AccountItemDisplayType field if non-nil, zero value otherwise.

### GetAccountItemDisplayTypeOk

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetAccountItemDisplayTypeOk() (*string, bool)`

GetAccountItemDisplayTypeOk returns a tuple with the AccountItemDisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemDisplayType

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) SetAccountItemDisplayType(v string)`

SetAccountItemDisplayType sets AccountItemDisplayType field to given value.

### HasAccountItemDisplayType

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) HasAccountItemDisplayType() bool`

HasAccountItemDisplayType returns a boolean if a field has been set.

### GetAdjustment

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetAdjustment() string`

GetAdjustment returns the Adjustment field if non-nil, zero value otherwise.

### GetAdjustmentOk

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetAdjustmentOk() (*string, bool)`

GetAdjustmentOk returns a tuple with the Adjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustment

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) SetAdjustment(v string)`

SetAdjustment sets Adjustment field to given value.

### HasAdjustment

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) HasAdjustment() bool`

HasAdjustment returns a boolean if a field has been set.

### GetApprovalFlowStatus

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetApprovalFlowStatus() string`

GetApprovalFlowStatus returns the ApprovalFlowStatus field if non-nil, zero value otherwise.

### GetApprovalFlowStatusOk

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetApprovalFlowStatusOk() (*string, bool)`

GetApprovalFlowStatusOk returns a tuple with the ApprovalFlowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalFlowStatus

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) SetApprovalFlowStatus(v string)`

SetApprovalFlowStatus sets ApprovalFlowStatus field to given value.

### HasApprovalFlowStatus

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) HasApprovalFlowStatus() bool`

HasApprovalFlowStatus returns a boolean if a field has been set.

### GetBalances

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetBalances() []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetBalancesOk() (*[]TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) SetBalances(v []TrialBsTwoYearsResponseTrialBsTwoYearsBalancesInner)`

SetBalances sets Balances field to given value.


### GetBreakdownDisplayType

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetBreakdownDisplayType() string`

GetBreakdownDisplayType returns the BreakdownDisplayType field if non-nil, zero value otherwise.

### GetBreakdownDisplayTypeOk

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetBreakdownDisplayTypeOk() (*string, bool)`

GetBreakdownDisplayTypeOk returns a tuple with the BreakdownDisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownDisplayType

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) SetBreakdownDisplayType(v string)`

SetBreakdownDisplayType sets BreakdownDisplayType field to given value.

### HasBreakdownDisplayType

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) HasBreakdownDisplayType() bool`

HasBreakdownDisplayType returns a boolean if a field has been set.

### GetCompanyId

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetCompanyId() int64`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetCompanyIdOk() (*int64, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) SetCompanyId(v int64)`

SetCompanyId sets CompanyId field to given value.


### GetCostAllocation

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetCostAllocation() string`

GetCostAllocation returns the CostAllocation field if non-nil, zero value otherwise.

### GetCostAllocationOk

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetCostAllocationOk() (*string, bool)`

GetCostAllocationOk returns a tuple with the CostAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostAllocation

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) SetCostAllocation(v string)`

SetCostAllocation sets CostAllocation field to given value.

### HasCostAllocation

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) HasCostAllocation() bool`

HasCostAllocation returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEndDate

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetEndMonth

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetEndMonth() int64`

GetEndMonth returns the EndMonth field if non-nil, zero value otherwise.

### GetEndMonthOk

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetEndMonthOk() (*int64, bool)`

GetEndMonthOk returns a tuple with the EndMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndMonth

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) SetEndMonth(v int64)`

SetEndMonth sets EndMonth field to given value.

### HasEndMonth

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) HasEndMonth() bool`

HasEndMonth returns a boolean if a field has been set.

### GetFiscalYear

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetFiscalYear() int64`

GetFiscalYear returns the FiscalYear field if non-nil, zero value otherwise.

### GetFiscalYearOk

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetFiscalYearOk() (*int64, bool)`

GetFiscalYearOk returns a tuple with the FiscalYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYear

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) SetFiscalYear(v int64)`

SetFiscalYear sets FiscalYear field to given value.

### HasFiscalYear

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) HasFiscalYear() bool`

HasFiscalYear returns a boolean if a field has been set.

### GetItemId

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetPartnerCode

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetPartnerCode() string`

GetPartnerCode returns the PartnerCode field if non-nil, zero value otherwise.

### GetPartnerCodeOk

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetPartnerCodeOk() (*string, bool)`

GetPartnerCodeOk returns a tuple with the PartnerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCode

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) SetPartnerCode(v string)`

SetPartnerCode sets PartnerCode field to given value.

### HasPartnerCode

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) HasPartnerCode() bool`

HasPartnerCode returns a boolean if a field has been set.

### GetPartnerId

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetPartnerId() int64`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetPartnerIdOk() (*int64, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) SetPartnerId(v int64)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetSectionId

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetSectionId() int64`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetSectionIdOk() (*int64, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) SetSectionId(v int64)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### GetStartDate

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStartMonth

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetStartMonth() int64`

GetStartMonth returns the StartMonth field if non-nil, zero value otherwise.

### GetStartMonthOk

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) GetStartMonthOk() (*int64, bool)`

GetStartMonthOk returns a tuple with the StartMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartMonth

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) SetStartMonth(v int64)`

SetStartMonth sets StartMonth field to given value.

### HasStartMonth

`func (o *TrialCrTwoYearsResponseTrialCrTwoYears) HasStartMonth() bool`

HasStartMonth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


