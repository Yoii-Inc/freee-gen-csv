# TrialPlTwoYearsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialPlTwoYears** | [**TrialCrTwoYearsResponseTrialCrTwoYears**](TrialCrTwoYearsResponseTrialCrTwoYears.md) |  | 
**UpToDate** | **bool** | 集計結果が最新かどうか | 
**UpToDateReasons** | Pointer to [**[]JournalsResponseJournalsUpToDateReasonsInner**](JournalsResponseJournalsUpToDateReasonsInner.md) | 集計が最新でない場合の要因情報 | [optional] 

## Methods

### NewTrialPlTwoYearsResponse

`func NewTrialPlTwoYearsResponse(trialPlTwoYears TrialCrTwoYearsResponseTrialCrTwoYears, upToDate bool, ) *TrialPlTwoYearsResponse`

NewTrialPlTwoYearsResponse instantiates a new TrialPlTwoYearsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialPlTwoYearsResponseWithDefaults

`func NewTrialPlTwoYearsResponseWithDefaults() *TrialPlTwoYearsResponse`

NewTrialPlTwoYearsResponseWithDefaults instantiates a new TrialPlTwoYearsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialPlTwoYears

`func (o *TrialPlTwoYearsResponse) GetTrialPlTwoYears() TrialCrTwoYearsResponseTrialCrTwoYears`

GetTrialPlTwoYears returns the TrialPlTwoYears field if non-nil, zero value otherwise.

### GetTrialPlTwoYearsOk

`func (o *TrialPlTwoYearsResponse) GetTrialPlTwoYearsOk() (*TrialCrTwoYearsResponseTrialCrTwoYears, bool)`

GetTrialPlTwoYearsOk returns a tuple with the TrialPlTwoYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPlTwoYears

`func (o *TrialPlTwoYearsResponse) SetTrialPlTwoYears(v TrialCrTwoYearsResponseTrialCrTwoYears)`

SetTrialPlTwoYears sets TrialPlTwoYears field to given value.


### GetUpToDate

`func (o *TrialPlTwoYearsResponse) GetUpToDate() bool`

GetUpToDate returns the UpToDate field if non-nil, zero value otherwise.

### GetUpToDateOk

`func (o *TrialPlTwoYearsResponse) GetUpToDateOk() (*bool, bool)`

GetUpToDateOk returns a tuple with the UpToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDate

`func (o *TrialPlTwoYearsResponse) SetUpToDate(v bool)`

SetUpToDate sets UpToDate field to given value.


### GetUpToDateReasons

`func (o *TrialPlTwoYearsResponse) GetUpToDateReasons() []JournalsResponseJournalsUpToDateReasonsInner`

GetUpToDateReasons returns the UpToDateReasons field if non-nil, zero value otherwise.

### GetUpToDateReasonsOk

`func (o *TrialPlTwoYearsResponse) GetUpToDateReasonsOk() (*[]JournalsResponseJournalsUpToDateReasonsInner, bool)`

GetUpToDateReasonsOk returns a tuple with the UpToDateReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDateReasons

`func (o *TrialPlTwoYearsResponse) SetUpToDateReasons(v []JournalsResponseJournalsUpToDateReasonsInner)`

SetUpToDateReasons sets UpToDateReasons field to given value.

### HasUpToDateReasons

`func (o *TrialPlTwoYearsResponse) HasUpToDateReasons() bool`

HasUpToDateReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


