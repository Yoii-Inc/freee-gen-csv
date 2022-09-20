# TrialCrSegment2TagsResponseTrialCrSegment2Tags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountItemDisplayType** | Pointer to **string** | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）(条件に指定した時のみ含まれる） | [optional] 
**Adjustment** | Pointer to **string** | 決算整理仕訳のみ: only, 決算整理仕訳以外: without(条件に指定した時のみ含まれる） | [optional] 
**ApprovalFlowStatus** | Pointer to **string** | 未承認を除く: without_in_progress (デフォルト), 全てのステータス: all(条件に指定した時のみ含まれる） | [optional] 
**Balances** | [**[]TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner**](TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner.md) |  | 
**BreakdownDisplayType** | Pointer to **string** | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item）(条件に指定した時のみ含まれる） | [optional] 
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
**Segment2TagIds** | **string** | 出力するセグメント2の指定 | 
**StartDate** | Pointer to **string** | 発生日で絞込：開始日(yyyy-mm-dd)(条件に指定した時のみ含まれる） | [optional] 
**StartMonth** | Pointer to **int64** | 発生月で絞込：開始会計月(1-12)(条件に指定した時のみ含まれる） | [optional] 

## Methods

### NewTrialCrSegment2TagsResponseTrialCrSegment2Tags

`func NewTrialCrSegment2TagsResponseTrialCrSegment2Tags(balances []TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner, companyId int64, segment2TagIds string, ) *TrialCrSegment2TagsResponseTrialCrSegment2Tags`

NewTrialCrSegment2TagsResponseTrialCrSegment2Tags instantiates a new TrialCrSegment2TagsResponseTrialCrSegment2Tags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialCrSegment2TagsResponseTrialCrSegment2TagsWithDefaults

`func NewTrialCrSegment2TagsResponseTrialCrSegment2TagsWithDefaults() *TrialCrSegment2TagsResponseTrialCrSegment2Tags`

NewTrialCrSegment2TagsResponseTrialCrSegment2TagsWithDefaults instantiates a new TrialCrSegment2TagsResponseTrialCrSegment2Tags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountItemDisplayType

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetAccountItemDisplayType() string`

GetAccountItemDisplayType returns the AccountItemDisplayType field if non-nil, zero value otherwise.

### GetAccountItemDisplayTypeOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetAccountItemDisplayTypeOk() (*string, bool)`

GetAccountItemDisplayTypeOk returns a tuple with the AccountItemDisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemDisplayType

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) SetAccountItemDisplayType(v string)`

SetAccountItemDisplayType sets AccountItemDisplayType field to given value.

### HasAccountItemDisplayType

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) HasAccountItemDisplayType() bool`

HasAccountItemDisplayType returns a boolean if a field has been set.

### GetAdjustment

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetAdjustment() string`

GetAdjustment returns the Adjustment field if non-nil, zero value otherwise.

### GetAdjustmentOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetAdjustmentOk() (*string, bool)`

GetAdjustmentOk returns a tuple with the Adjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustment

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) SetAdjustment(v string)`

SetAdjustment sets Adjustment field to given value.

### HasAdjustment

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) HasAdjustment() bool`

HasAdjustment returns a boolean if a field has been set.

### GetApprovalFlowStatus

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetApprovalFlowStatus() string`

GetApprovalFlowStatus returns the ApprovalFlowStatus field if non-nil, zero value otherwise.

### GetApprovalFlowStatusOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetApprovalFlowStatusOk() (*string, bool)`

GetApprovalFlowStatusOk returns a tuple with the ApprovalFlowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalFlowStatus

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) SetApprovalFlowStatus(v string)`

SetApprovalFlowStatus sets ApprovalFlowStatus field to given value.

### HasApprovalFlowStatus

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) HasApprovalFlowStatus() bool`

HasApprovalFlowStatus returns a boolean if a field has been set.

### GetBalances

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetBalances() []TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetBalancesOk() (*[]TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) SetBalances(v []TrialCrSegment2TagsResponseTrialCrSegment2TagsBalancesInner)`

SetBalances sets Balances field to given value.


### GetBreakdownDisplayType

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetBreakdownDisplayType() string`

GetBreakdownDisplayType returns the BreakdownDisplayType field if non-nil, zero value otherwise.

### GetBreakdownDisplayTypeOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetBreakdownDisplayTypeOk() (*string, bool)`

GetBreakdownDisplayTypeOk returns a tuple with the BreakdownDisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownDisplayType

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) SetBreakdownDisplayType(v string)`

SetBreakdownDisplayType sets BreakdownDisplayType field to given value.

### HasBreakdownDisplayType

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) HasBreakdownDisplayType() bool`

HasBreakdownDisplayType returns a boolean if a field has been set.

### GetCompanyId

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetCompanyId() int64`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetCompanyIdOk() (*int64, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) SetCompanyId(v int64)`

SetCompanyId sets CompanyId field to given value.


### GetCostAllocation

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetCostAllocation() string`

GetCostAllocation returns the CostAllocation field if non-nil, zero value otherwise.

### GetCostAllocationOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetCostAllocationOk() (*string, bool)`

GetCostAllocationOk returns a tuple with the CostAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostAllocation

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) SetCostAllocation(v string)`

SetCostAllocation sets CostAllocation field to given value.

### HasCostAllocation

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) HasCostAllocation() bool`

HasCostAllocation returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEndDate

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetEndMonth

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetEndMonth() int64`

GetEndMonth returns the EndMonth field if non-nil, zero value otherwise.

### GetEndMonthOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetEndMonthOk() (*int64, bool)`

GetEndMonthOk returns a tuple with the EndMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndMonth

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) SetEndMonth(v int64)`

SetEndMonth sets EndMonth field to given value.

### HasEndMonth

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) HasEndMonth() bool`

HasEndMonth returns a boolean if a field has been set.

### GetFiscalYear

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetFiscalYear() int64`

GetFiscalYear returns the FiscalYear field if non-nil, zero value otherwise.

### GetFiscalYearOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetFiscalYearOk() (*int64, bool)`

GetFiscalYearOk returns a tuple with the FiscalYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYear

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) SetFiscalYear(v int64)`

SetFiscalYear sets FiscalYear field to given value.

### HasFiscalYear

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) HasFiscalYear() bool`

HasFiscalYear returns a boolean if a field has been set.

### GetItemId

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetPartnerCode

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetPartnerCode() string`

GetPartnerCode returns the PartnerCode field if non-nil, zero value otherwise.

### GetPartnerCodeOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetPartnerCodeOk() (*string, bool)`

GetPartnerCodeOk returns a tuple with the PartnerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCode

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) SetPartnerCode(v string)`

SetPartnerCode sets PartnerCode field to given value.

### HasPartnerCode

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) HasPartnerCode() bool`

HasPartnerCode returns a boolean if a field has been set.

### GetPartnerId

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetPartnerId() int64`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetPartnerIdOk() (*int64, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) SetPartnerId(v int64)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetSectionId

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetSectionId() int64`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetSectionIdOk() (*int64, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) SetSectionId(v int64)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### GetSegment2TagIds

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetSegment2TagIds() string`

GetSegment2TagIds returns the Segment2TagIds field if non-nil, zero value otherwise.

### GetSegment2TagIdsOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetSegment2TagIdsOk() (*string, bool)`

GetSegment2TagIdsOk returns a tuple with the Segment2TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagIds

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) SetSegment2TagIds(v string)`

SetSegment2TagIds sets Segment2TagIds field to given value.


### GetStartDate

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStartMonth

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetStartMonth() int64`

GetStartMonth returns the StartMonth field if non-nil, zero value otherwise.

### GetStartMonthOk

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) GetStartMonthOk() (*int64, bool)`

GetStartMonthOk returns a tuple with the StartMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartMonth

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) SetStartMonth(v int64)`

SetStartMonth sets StartMonth field to given value.

### HasStartMonth

`func (o *TrialCrSegment2TagsResponseTrialCrSegment2Tags) HasStartMonth() bool`

HasStartMonth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


