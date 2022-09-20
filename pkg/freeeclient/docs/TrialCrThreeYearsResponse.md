# TrialCrThreeYearsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialCrThreeYears** | [**TrialCrThreeYearsResponseTrialCrThreeYears**](TrialCrThreeYearsResponseTrialCrThreeYears.md) |  | 
**UpToDate** | **bool** | 集計結果が最新かどうか | 
**UpToDateReasons** | Pointer to [**[]JournalsResponseJournalsUpToDateReasonsInner**](JournalsResponseJournalsUpToDateReasonsInner.md) | 集計が最新でない場合の要因情報 | [optional] 

## Methods

### NewTrialCrThreeYearsResponse

`func NewTrialCrThreeYearsResponse(trialCrThreeYears TrialCrThreeYearsResponseTrialCrThreeYears, upToDate bool, ) *TrialCrThreeYearsResponse`

NewTrialCrThreeYearsResponse instantiates a new TrialCrThreeYearsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialCrThreeYearsResponseWithDefaults

`func NewTrialCrThreeYearsResponseWithDefaults() *TrialCrThreeYearsResponse`

NewTrialCrThreeYearsResponseWithDefaults instantiates a new TrialCrThreeYearsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialCrThreeYears

`func (o *TrialCrThreeYearsResponse) GetTrialCrThreeYears() TrialCrThreeYearsResponseTrialCrThreeYears`

GetTrialCrThreeYears returns the TrialCrThreeYears field if non-nil, zero value otherwise.

### GetTrialCrThreeYearsOk

`func (o *TrialCrThreeYearsResponse) GetTrialCrThreeYearsOk() (*TrialCrThreeYearsResponseTrialCrThreeYears, bool)`

GetTrialCrThreeYearsOk returns a tuple with the TrialCrThreeYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialCrThreeYears

`func (o *TrialCrThreeYearsResponse) SetTrialCrThreeYears(v TrialCrThreeYearsResponseTrialCrThreeYears)`

SetTrialCrThreeYears sets TrialCrThreeYears field to given value.


### GetUpToDate

`func (o *TrialCrThreeYearsResponse) GetUpToDate() bool`

GetUpToDate returns the UpToDate field if non-nil, zero value otherwise.

### GetUpToDateOk

`func (o *TrialCrThreeYearsResponse) GetUpToDateOk() (*bool, bool)`

GetUpToDateOk returns a tuple with the UpToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDate

`func (o *TrialCrThreeYearsResponse) SetUpToDate(v bool)`

SetUpToDate sets UpToDate field to given value.


### GetUpToDateReasons

`func (o *TrialCrThreeYearsResponse) GetUpToDateReasons() []JournalsResponseJournalsUpToDateReasonsInner`

GetUpToDateReasons returns the UpToDateReasons field if non-nil, zero value otherwise.

### GetUpToDateReasonsOk

`func (o *TrialCrThreeYearsResponse) GetUpToDateReasonsOk() (*[]JournalsResponseJournalsUpToDateReasonsInner, bool)`

GetUpToDateReasonsOk returns a tuple with the UpToDateReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDateReasons

`func (o *TrialCrThreeYearsResponse) SetUpToDateReasons(v []JournalsResponseJournalsUpToDateReasonsInner)`

SetUpToDateReasons sets UpToDateReasons field to given value.

### HasUpToDateReasons

`func (o *TrialCrThreeYearsResponse) HasUpToDateReasons() bool`

HasUpToDateReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


