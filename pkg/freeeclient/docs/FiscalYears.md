# FiscalYears

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepreciationRecordMethod** | **int64** | 月次償却（0: しない、1: する） | 
**EndDate** | Pointer to **string** | 期末日 | [optional] 
**IndirectWriteOffMethod** | **bool** | 固定資産の控除法（true: 間接控除法、false: 直接控除法） | 
**ReturnCode** | **int64** | 不動産所得使用区分（0: 一般、3: 一般/不動産） ※個人事業主のみ設定可能 | 
**SalesTaxBusinessCode** | **int64** | 簡易課税用事業区分（0: 第一種：卸売業、1: 第二種：小売業、2: 第三種：農林水産業、工業、建設業、製造業など、3: 第四種：飲食店業など、4: 第五種：金融・保険業、運輸通信業、サービス業など、5: 第六種：不動産業など | 
**StartDate** | Pointer to **string** | 期首日 | [optional] 
**TaxAccountMethod** | **int64** | 消費税経理処理方法（0: 税込経理、1: 旧税抜経理、2: 税抜経理） | 
**TaxFraction** | **int64** | 消費税端数処理方法（0: 切り捨て、1: 切り上げ、2: 四捨五入） | 
**TaxMethod** | **int64** | 課税区分（0: 免税、1: 簡易課税、2: 本則課税（個別対応方式）、3: 本則課税（一括比例配分方式）、4: 本則課税（全額控除）） | 
**UseIndustryTemplate** | **bool** | 製造業向け機能（true: 使用する、false: 使用しない） | 

## Methods

### NewFiscalYears

`func NewFiscalYears(depreciationRecordMethod int64, indirectWriteOffMethod bool, returnCode int64, salesTaxBusinessCode int64, taxAccountMethod int64, taxFraction int64, taxMethod int64, useIndustryTemplate bool, ) *FiscalYears`

NewFiscalYears instantiates a new FiscalYears object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiscalYearsWithDefaults

`func NewFiscalYearsWithDefaults() *FiscalYears`

NewFiscalYearsWithDefaults instantiates a new FiscalYears object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepreciationRecordMethod

`func (o *FiscalYears) GetDepreciationRecordMethod() int64`

GetDepreciationRecordMethod returns the DepreciationRecordMethod field if non-nil, zero value otherwise.

### GetDepreciationRecordMethodOk

`func (o *FiscalYears) GetDepreciationRecordMethodOk() (*int64, bool)`

GetDepreciationRecordMethodOk returns a tuple with the DepreciationRecordMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepreciationRecordMethod

`func (o *FiscalYears) SetDepreciationRecordMethod(v int64)`

SetDepreciationRecordMethod sets DepreciationRecordMethod field to given value.


### GetEndDate

`func (o *FiscalYears) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *FiscalYears) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *FiscalYears) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *FiscalYears) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetIndirectWriteOffMethod

`func (o *FiscalYears) GetIndirectWriteOffMethod() bool`

GetIndirectWriteOffMethod returns the IndirectWriteOffMethod field if non-nil, zero value otherwise.

### GetIndirectWriteOffMethodOk

`func (o *FiscalYears) GetIndirectWriteOffMethodOk() (*bool, bool)`

GetIndirectWriteOffMethodOk returns a tuple with the IndirectWriteOffMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndirectWriteOffMethod

`func (o *FiscalYears) SetIndirectWriteOffMethod(v bool)`

SetIndirectWriteOffMethod sets IndirectWriteOffMethod field to given value.


### GetReturnCode

`func (o *FiscalYears) GetReturnCode() int64`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *FiscalYears) GetReturnCodeOk() (*int64, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *FiscalYears) SetReturnCode(v int64)`

SetReturnCode sets ReturnCode field to given value.


### GetSalesTaxBusinessCode

`func (o *FiscalYears) GetSalesTaxBusinessCode() int64`

GetSalesTaxBusinessCode returns the SalesTaxBusinessCode field if non-nil, zero value otherwise.

### GetSalesTaxBusinessCodeOk

`func (o *FiscalYears) GetSalesTaxBusinessCodeOk() (*int64, bool)`

GetSalesTaxBusinessCodeOk returns a tuple with the SalesTaxBusinessCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxBusinessCode

`func (o *FiscalYears) SetSalesTaxBusinessCode(v int64)`

SetSalesTaxBusinessCode sets SalesTaxBusinessCode field to given value.


### GetStartDate

`func (o *FiscalYears) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *FiscalYears) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *FiscalYears) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *FiscalYears) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTaxAccountMethod

`func (o *FiscalYears) GetTaxAccountMethod() int64`

GetTaxAccountMethod returns the TaxAccountMethod field if non-nil, zero value otherwise.

### GetTaxAccountMethodOk

`func (o *FiscalYears) GetTaxAccountMethodOk() (*int64, bool)`

GetTaxAccountMethodOk returns a tuple with the TaxAccountMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAccountMethod

`func (o *FiscalYears) SetTaxAccountMethod(v int64)`

SetTaxAccountMethod sets TaxAccountMethod field to given value.


### GetTaxFraction

`func (o *FiscalYears) GetTaxFraction() int64`

GetTaxFraction returns the TaxFraction field if non-nil, zero value otherwise.

### GetTaxFractionOk

`func (o *FiscalYears) GetTaxFractionOk() (*int64, bool)`

GetTaxFractionOk returns a tuple with the TaxFraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxFraction

`func (o *FiscalYears) SetTaxFraction(v int64)`

SetTaxFraction sets TaxFraction field to given value.


### GetTaxMethod

`func (o *FiscalYears) GetTaxMethod() int64`

GetTaxMethod returns the TaxMethod field if non-nil, zero value otherwise.

### GetTaxMethodOk

`func (o *FiscalYears) GetTaxMethodOk() (*int64, bool)`

GetTaxMethodOk returns a tuple with the TaxMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxMethod

`func (o *FiscalYears) SetTaxMethod(v int64)`

SetTaxMethod sets TaxMethod field to given value.


### GetUseIndustryTemplate

`func (o *FiscalYears) GetUseIndustryTemplate() bool`

GetUseIndustryTemplate returns the UseIndustryTemplate field if non-nil, zero value otherwise.

### GetUseIndustryTemplateOk

`func (o *FiscalYears) GetUseIndustryTemplateOk() (*bool, bool)`

GetUseIndustryTemplateOk returns a tuple with the UseIndustryTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIndustryTemplate

`func (o *FiscalYears) SetUseIndustryTemplate(v bool)`

SetUseIndustryTemplate sets UseIndustryTemplate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


