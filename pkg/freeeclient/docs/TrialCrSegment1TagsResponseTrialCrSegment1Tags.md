# TrialCrSegment1TagsResponseTrialCrSegment1Tags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountItemDisplayType** | Pointer to **string** | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）(条件に指定した時のみ含まれる） | [optional] 
**Adjustment** | Pointer to **string** | 決算整理仕訳のみ: only, 決算整理仕訳以外: without(条件に指定した時のみ含まれる） | [optional] 
**ApprovalFlowStatus** | Pointer to **string** | 未承認を除く: without_in_progress (デフォルト), 全てのステータス: all(条件に指定した時のみ含まれる） | [optional] 
**Balances** | [**[]TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner**](TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner.md) |  | 
**BreakdownDisplayType** | Pointer to **string** | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item）(条件に指定した時のみ含まれる） | [optional] 
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
**Segment1TagIds** | **string** | 出力するセグメント1の指定 | 
**StartDate** | Pointer to **string** | 発生日で絞込：開始日(yyyy-mm-dd)(条件に指定した時のみ含まれる） | [optional] 
**StartMonth** | Pointer to **int32** | 発生月で絞込：開始会計月(1-12)(条件に指定した時のみ含まれる） | [optional] 

## Methods

### NewTrialCrSegment1TagsResponseTrialCrSegment1Tags

`func NewTrialCrSegment1TagsResponseTrialCrSegment1Tags(balances []TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner, companyId int32, segment1TagIds string, ) *TrialCrSegment1TagsResponseTrialCrSegment1Tags`

NewTrialCrSegment1TagsResponseTrialCrSegment1Tags instantiates a new TrialCrSegment1TagsResponseTrialCrSegment1Tags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialCrSegment1TagsResponseTrialCrSegment1TagsWithDefaults

`func NewTrialCrSegment1TagsResponseTrialCrSegment1TagsWithDefaults() *TrialCrSegment1TagsResponseTrialCrSegment1Tags`

NewTrialCrSegment1TagsResponseTrialCrSegment1TagsWithDefaults instantiates a new TrialCrSegment1TagsResponseTrialCrSegment1Tags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountItemDisplayType

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetAccountItemDisplayType() string`

GetAccountItemDisplayType returns the AccountItemDisplayType field if non-nil, zero value otherwise.

### GetAccountItemDisplayTypeOk

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetAccountItemDisplayTypeOk() (*string, bool)`

GetAccountItemDisplayTypeOk returns a tuple with the AccountItemDisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemDisplayType

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) SetAccountItemDisplayType(v string)`

SetAccountItemDisplayType sets AccountItemDisplayType field to given value.

### HasAccountItemDisplayType

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) HasAccountItemDisplayType() bool`

HasAccountItemDisplayType returns a boolean if a field has been set.

### GetAdjustment

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetAdjustment() string`

GetAdjustment returns the Adjustment field if non-nil, zero value otherwise.

### GetAdjustmentOk

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetAdjustmentOk() (*string, bool)`

GetAdjustmentOk returns a tuple with the Adjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustment

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) SetAdjustment(v string)`

SetAdjustment sets Adjustment field to given value.

### HasAdjustment

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) HasAdjustment() bool`

HasAdjustment returns a boolean if a field has been set.

### GetApprovalFlowStatus

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetApprovalFlowStatus() string`

GetApprovalFlowStatus returns the ApprovalFlowStatus field if non-nil, zero value otherwise.

### GetApprovalFlowStatusOk

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetApprovalFlowStatusOk() (*string, bool)`

GetApprovalFlowStatusOk returns a tuple with the ApprovalFlowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalFlowStatus

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) SetApprovalFlowStatus(v string)`

SetApprovalFlowStatus sets ApprovalFlowStatus field to given value.

### HasApprovalFlowStatus

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) HasApprovalFlowStatus() bool`

HasApprovalFlowStatus returns a boolean if a field has been set.

### GetBalances

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetBalances() []TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetBalancesOk() (*[]TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) SetBalances(v []TrialCrSegment1TagsResponseTrialCrSegment1TagsBalancesInner)`

SetBalances sets Balances field to given value.


### GetBreakdownDisplayType

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetBreakdownDisplayType() string`

GetBreakdownDisplayType returns the BreakdownDisplayType field if non-nil, zero value otherwise.

### GetBreakdownDisplayTypeOk

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetBreakdownDisplayTypeOk() (*string, bool)`

GetBreakdownDisplayTypeOk returns a tuple with the BreakdownDisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownDisplayType

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) SetBreakdownDisplayType(v string)`

SetBreakdownDisplayType sets BreakdownDisplayType field to given value.

### HasBreakdownDisplayType

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) HasBreakdownDisplayType() bool`

HasBreakdownDisplayType returns a boolean if a field has been set.

### GetCompanyId

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetCostAllocation

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetCostAllocation() string`

GetCostAllocation returns the CostAllocation field if non-nil, zero value otherwise.

### GetCostAllocationOk

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetCostAllocationOk() (*string, bool)`

GetCostAllocationOk returns a tuple with the CostAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostAllocation

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) SetCostAllocation(v string)`

SetCostAllocation sets CostAllocation field to given value.

### HasCostAllocation

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) HasCostAllocation() bool`

HasCostAllocation returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEndDate

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetEndMonth

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetEndMonth() int32`

GetEndMonth returns the EndMonth field if non-nil, zero value otherwise.

### GetEndMonthOk

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetEndMonthOk() (*int32, bool)`

GetEndMonthOk returns a tuple with the EndMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndMonth

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) SetEndMonth(v int32)`

SetEndMonth sets EndMonth field to given value.

### HasEndMonth

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) HasEndMonth() bool`

HasEndMonth returns a boolean if a field has been set.

### GetFiscalYear

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetFiscalYear() int32`

GetFiscalYear returns the FiscalYear field if non-nil, zero value otherwise.

### GetFiscalYearOk

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetFiscalYearOk() (*int32, bool)`

GetFiscalYearOk returns a tuple with the FiscalYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYear

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) SetFiscalYear(v int32)`

SetFiscalYear sets FiscalYear field to given value.

### HasFiscalYear

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) HasFiscalYear() bool`

HasFiscalYear returns a boolean if a field has been set.

### GetItemId

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetItemId() int32`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetItemIdOk() (*int32, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) SetItemId(v int32)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetPartnerCode

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetPartnerCode() string`

GetPartnerCode returns the PartnerCode field if non-nil, zero value otherwise.

### GetPartnerCodeOk

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetPartnerCodeOk() (*string, bool)`

GetPartnerCodeOk returns a tuple with the PartnerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCode

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) SetPartnerCode(v string)`

SetPartnerCode sets PartnerCode field to given value.

### HasPartnerCode

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) HasPartnerCode() bool`

HasPartnerCode returns a boolean if a field has been set.

### GetPartnerId

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetSectionId

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetSectionId() int32`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetSectionIdOk() (*int32, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) SetSectionId(v int32)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### GetSegment1TagIds

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetSegment1TagIds() string`

GetSegment1TagIds returns the Segment1TagIds field if non-nil, zero value otherwise.

### GetSegment1TagIdsOk

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetSegment1TagIdsOk() (*string, bool)`

GetSegment1TagIdsOk returns a tuple with the Segment1TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagIds

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) SetSegment1TagIds(v string)`

SetSegment1TagIds sets Segment1TagIds field to given value.


### GetStartDate

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStartMonth

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetStartMonth() int32`

GetStartMonth returns the StartMonth field if non-nil, zero value otherwise.

### GetStartMonthOk

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) GetStartMonthOk() (*int32, bool)`

GetStartMonthOk returns a tuple with the StartMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartMonth

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) SetStartMonth(v int32)`

SetStartMonth sets StartMonth field to given value.

### HasStartMonth

`func (o *TrialCrSegment1TagsResponseTrialCrSegment1Tags) HasStartMonth() bool`

HasStartMonth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


