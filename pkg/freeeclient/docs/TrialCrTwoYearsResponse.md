# TrialCrTwoYearsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialCrTwoYears** | [**TrialCrTwoYearsResponseTrialCrTwoYears**](TrialCrTwoYearsResponseTrialCrTwoYears.md) |  | 
**UpToDate** | **bool** | 集計結果が最新かどうか | 
**UpToDateReasons** | Pointer to [**[]JournalsResponseJournalsUpToDateReasonsInner**](JournalsResponseJournalsUpToDateReasonsInner.md) | 集計が最新でない場合の要因情報 | [optional] 

## Methods

### NewTrialCrTwoYearsResponse

`func NewTrialCrTwoYearsResponse(trialCrTwoYears TrialCrTwoYearsResponseTrialCrTwoYears, upToDate bool, ) *TrialCrTwoYearsResponse`

NewTrialCrTwoYearsResponse instantiates a new TrialCrTwoYearsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialCrTwoYearsResponseWithDefaults

`func NewTrialCrTwoYearsResponseWithDefaults() *TrialCrTwoYearsResponse`

NewTrialCrTwoYearsResponseWithDefaults instantiates a new TrialCrTwoYearsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialCrTwoYears

`func (o *TrialCrTwoYearsResponse) GetTrialCrTwoYears() TrialCrTwoYearsResponseTrialCrTwoYears`

GetTrialCrTwoYears returns the TrialCrTwoYears field if non-nil, zero value otherwise.

### GetTrialCrTwoYearsOk

`func (o *TrialCrTwoYearsResponse) GetTrialCrTwoYearsOk() (*TrialCrTwoYearsResponseTrialCrTwoYears, bool)`

GetTrialCrTwoYearsOk returns a tuple with the TrialCrTwoYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialCrTwoYears

`func (o *TrialCrTwoYearsResponse) SetTrialCrTwoYears(v TrialCrTwoYearsResponseTrialCrTwoYears)`

SetTrialCrTwoYears sets TrialCrTwoYears field to given value.


### GetUpToDate

`func (o *TrialCrTwoYearsResponse) GetUpToDate() bool`

GetUpToDate returns the UpToDate field if non-nil, zero value otherwise.

### GetUpToDateOk

`func (o *TrialCrTwoYearsResponse) GetUpToDateOk() (*bool, bool)`

GetUpToDateOk returns a tuple with the UpToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDate

`func (o *TrialCrTwoYearsResponse) SetUpToDate(v bool)`

SetUpToDate sets UpToDate field to given value.


### GetUpToDateReasons

`func (o *TrialCrTwoYearsResponse) GetUpToDateReasons() []JournalsResponseJournalsUpToDateReasonsInner`

GetUpToDateReasons returns the UpToDateReasons field if non-nil, zero value otherwise.

### GetUpToDateReasonsOk

`func (o *TrialCrTwoYearsResponse) GetUpToDateReasonsOk() (*[]JournalsResponseJournalsUpToDateReasonsInner, bool)`

GetUpToDateReasonsOk returns a tuple with the UpToDateReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDateReasons

`func (o *TrialCrTwoYearsResponse) SetUpToDateReasons(v []JournalsResponseJournalsUpToDateReasonsInner)`

SetUpToDateReasons sets UpToDateReasons field to given value.

### HasUpToDateReasons

`func (o *TrialCrTwoYearsResponse) HasUpToDateReasons() bool`

HasUpToDateReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


