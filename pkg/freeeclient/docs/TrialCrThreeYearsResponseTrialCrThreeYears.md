# TrialCrThreeYearsResponseTrialCrThreeYears

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountItemDisplayType** | Pointer to **string** | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）(条件に指定した時のみ含まれる） | [optional] 
**Adjustment** | Pointer to **string** | 決算整理仕訳のみ: only, 決算整理仕訳以外: without(条件に指定した時のみ含まれる） | [optional] 
**ApprovalFlowStatus** | Pointer to **string** | 未承認を除く: without_in_progress (デフォルト), 全てのステータス: all(条件に指定した時のみ含まれる） | [optional] 
**Balances** | [**[]TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner**](TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner.md) |  | 
**BreakdownDisplayType** | Pointer to **string** | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item, セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン): segment_1_tag, セグメント2(法人向け エンタープライズプラン):segment_2_tag, セグメント3(法人向け エンタープライズプラン): segment_3_tag）(条件に指定した時のみ含まれる） | [optional] 
**CompanyId** | **int32** | 事業所ID | 
**CostAllocation** | Pointer to **string** | 配賦仕訳のみ：only,配賦仕訳以外：without(条件に指定した時のみ含まれる） | [optional] 
**CreatedAt** | Pointer to **string** | 作成日時 | [optional] 
**EndDate** | Pointer to **string** | 発生日で絞込：終了日(yyyy-mm-dd)(条件に指定した時のみ含まれる） | [optional] 
**EndMonth** | Pointer to **int32** | 発生月で絞込：終了会計月(1-12)(条件に指定した時のみ含まれる） | [optional] 
**FiscalYear** | Pointer to **int32** | 会計年度(条件に指定した時、または条件に月、日条件がない時のみ含まれる） | [optional] 
**ItemId** | Pointer to **int32** | 品目ID(条件に指定した時のみ含まれる） | [optional] 
**PartnerCode** | Pointer to **string** | 取引先コード(条件に指定した時のみ含まれる） | [optional] 
**PartnerId** | Pointer to **int32** | 取引先ID(条件に指定した時のみ含まれる） | [optional] 
**SectionId** | Pointer to **int32** | 部門ID(条件に指定した時のみ含まれる） | [optional] 
**StartDate** | Pointer to **string** | 発生日で絞込：開始日(yyyy-mm-dd)(条件に指定した時のみ含まれる） | [optional] 
**StartMonth** | Pointer to **int32** | 発生月で絞込：開始会計月(1-12)(条件に指定した時のみ含まれる） | [optional] 

## Methods

### NewTrialCrThreeYearsResponseTrialCrThreeYears

`func NewTrialCrThreeYearsResponseTrialCrThreeYears(balances []TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner, companyId int32, ) *TrialCrThreeYearsResponseTrialCrThreeYears`

NewTrialCrThreeYearsResponseTrialCrThreeYears instantiates a new TrialCrThreeYearsResponseTrialCrThreeYears object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialCrThreeYearsResponseTrialCrThreeYearsWithDefaults

`func NewTrialCrThreeYearsResponseTrialCrThreeYearsWithDefaults() *TrialCrThreeYearsResponseTrialCrThreeYears`

NewTrialCrThreeYearsResponseTrialCrThreeYearsWithDefaults instantiates a new TrialCrThreeYearsResponseTrialCrThreeYears object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountItemDisplayType

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetAccountItemDisplayType() string`

GetAccountItemDisplayType returns the AccountItemDisplayType field if non-nil, zero value otherwise.

### GetAccountItemDisplayTypeOk

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetAccountItemDisplayTypeOk() (*string, bool)`

GetAccountItemDisplayTypeOk returns a tuple with the AccountItemDisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemDisplayType

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) SetAccountItemDisplayType(v string)`

SetAccountItemDisplayType sets AccountItemDisplayType field to given value.

### HasAccountItemDisplayType

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) HasAccountItemDisplayType() bool`

HasAccountItemDisplayType returns a boolean if a field has been set.

### GetAdjustment

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetAdjustment() string`

GetAdjustment returns the Adjustment field if non-nil, zero value otherwise.

### GetAdjustmentOk

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetAdjustmentOk() (*string, bool)`

GetAdjustmentOk returns a tuple with the Adjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustment

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) SetAdjustment(v string)`

SetAdjustment sets Adjustment field to given value.

### HasAdjustment

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) HasAdjustment() bool`

HasAdjustment returns a boolean if a field has been set.

### GetApprovalFlowStatus

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetApprovalFlowStatus() string`

GetApprovalFlowStatus returns the ApprovalFlowStatus field if non-nil, zero value otherwise.

### GetApprovalFlowStatusOk

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetApprovalFlowStatusOk() (*string, bool)`

GetApprovalFlowStatusOk returns a tuple with the ApprovalFlowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalFlowStatus

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) SetApprovalFlowStatus(v string)`

SetApprovalFlowStatus sets ApprovalFlowStatus field to given value.

### HasApprovalFlowStatus

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) HasApprovalFlowStatus() bool`

HasApprovalFlowStatus returns a boolean if a field has been set.

### GetBalances

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetBalances() []TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetBalancesOk() (*[]TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) SetBalances(v []TrialBsThreeYearsResponseTrialBsThreeYearsBalancesInner)`

SetBalances sets Balances field to given value.


### GetBreakdownDisplayType

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetBreakdownDisplayType() string`

GetBreakdownDisplayType returns the BreakdownDisplayType field if non-nil, zero value otherwise.

### GetBreakdownDisplayTypeOk

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetBreakdownDisplayTypeOk() (*string, bool)`

GetBreakdownDisplayTypeOk returns a tuple with the BreakdownDisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownDisplayType

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) SetBreakdownDisplayType(v string)`

SetBreakdownDisplayType sets BreakdownDisplayType field to given value.

### HasBreakdownDisplayType

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) HasBreakdownDisplayType() bool`

HasBreakdownDisplayType returns a boolean if a field has been set.

### GetCompanyId

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetCostAllocation

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetCostAllocation() string`

GetCostAllocation returns the CostAllocation field if non-nil, zero value otherwise.

### GetCostAllocationOk

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetCostAllocationOk() (*string, bool)`

GetCostAllocationOk returns a tuple with the CostAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostAllocation

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) SetCostAllocation(v string)`

SetCostAllocation sets CostAllocation field to given value.

### HasCostAllocation

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) HasCostAllocation() bool`

HasCostAllocation returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEndDate

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetEndMonth

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetEndMonth() int32`

GetEndMonth returns the EndMonth field if non-nil, zero value otherwise.

### GetEndMonthOk

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetEndMonthOk() (*int32, bool)`

GetEndMonthOk returns a tuple with the EndMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndMonth

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) SetEndMonth(v int32)`

SetEndMonth sets EndMonth field to given value.

### HasEndMonth

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) HasEndMonth() bool`

HasEndMonth returns a boolean if a field has been set.

### GetFiscalYear

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetFiscalYear() int32`

GetFiscalYear returns the FiscalYear field if non-nil, zero value otherwise.

### GetFiscalYearOk

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetFiscalYearOk() (*int32, bool)`

GetFiscalYearOk returns a tuple with the FiscalYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYear

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) SetFiscalYear(v int32)`

SetFiscalYear sets FiscalYear field to given value.

### HasFiscalYear

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) HasFiscalYear() bool`

HasFiscalYear returns a boolean if a field has been set.

### GetItemId

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetItemId() int32`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetItemIdOk() (*int32, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) SetItemId(v int32)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetPartnerCode

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetPartnerCode() string`

GetPartnerCode returns the PartnerCode field if non-nil, zero value otherwise.

### GetPartnerCodeOk

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetPartnerCodeOk() (*string, bool)`

GetPartnerCodeOk returns a tuple with the PartnerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCode

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) SetPartnerCode(v string)`

SetPartnerCode sets PartnerCode field to given value.

### HasPartnerCode

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) HasPartnerCode() bool`

HasPartnerCode returns a boolean if a field has been set.

### GetPartnerId

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetSectionId

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetSectionId() int32`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetSectionIdOk() (*int32, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) SetSectionId(v int32)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### GetStartDate

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStartMonth

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetStartMonth() int32`

GetStartMonth returns the StartMonth field if non-nil, zero value otherwise.

### GetStartMonthOk

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) GetStartMonthOk() (*int32, bool)`

GetStartMonthOk returns a tuple with the StartMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartMonth

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) SetStartMonth(v int32)`

SetStartMonth sets StartMonth field to given value.

### HasStartMonth

`func (o *TrialCrThreeYearsResponseTrialCrThreeYears) HasStartMonth() bool`

HasStartMonth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


