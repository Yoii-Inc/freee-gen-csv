# TrialPlThreeYearsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialPlThreeYears** | [**TrialCrThreeYearsResponseTrialCrThreeYears**](TrialCrThreeYearsResponseTrialCrThreeYears.md) |  | 
**UpToDate** | **bool** | 集計結果が最新かどうか | 
**UpToDateReasons** | Pointer to [**[]JournalsResponseJournalsUpToDateReasonsInner**](JournalsResponseJournalsUpToDateReasonsInner.md) | 集計が最新でない場合の要因情報 | [optional] 

## Methods

### NewTrialPlThreeYearsResponse

`func NewTrialPlThreeYearsResponse(trialPlThreeYears TrialCrThreeYearsResponseTrialCrThreeYears, upToDate bool, ) *TrialPlThreeYearsResponse`

NewTrialPlThreeYearsResponse instantiates a new TrialPlThreeYearsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialPlThreeYearsResponseWithDefaults

`func NewTrialPlThreeYearsResponseWithDefaults() *TrialPlThreeYearsResponse`

NewTrialPlThreeYearsResponseWithDefaults instantiates a new TrialPlThreeYearsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialPlThreeYears

`func (o *TrialPlThreeYearsResponse) GetTrialPlThreeYears() TrialCrThreeYearsResponseTrialCrThreeYears`

GetTrialPlThreeYears returns the TrialPlThreeYears field if non-nil, zero value otherwise.

### GetTrialPlThreeYearsOk

`func (o *TrialPlThreeYearsResponse) GetTrialPlThreeYearsOk() (*TrialCrThreeYearsResponseTrialCrThreeYears, bool)`

GetTrialPlThreeYearsOk returns a tuple with the TrialPlThreeYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPlThreeYears

`func (o *TrialPlThreeYearsResponse) SetTrialPlThreeYears(v TrialCrThreeYearsResponseTrialCrThreeYears)`

SetTrialPlThreeYears sets TrialPlThreeYears field to given value.


### GetUpToDate

`func (o *TrialPlThreeYearsResponse) GetUpToDate() bool`

GetUpToDate returns the UpToDate field if non-nil, zero value otherwise.

### GetUpToDateOk

`func (o *TrialPlThreeYearsResponse) GetUpToDateOk() (*bool, bool)`

GetUpToDateOk returns a tuple with the UpToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDate

`func (o *TrialPlThreeYearsResponse) SetUpToDate(v bool)`

SetUpToDate sets UpToDate field to given value.


### GetUpToDateReasons

`func (o *TrialPlThreeYearsResponse) GetUpToDateReasons() []JournalsResponseJournalsUpToDateReasonsInner`

GetUpToDateReasons returns the UpToDateReasons field if non-nil, zero value otherwise.

### GetUpToDateReasonsOk

`func (o *TrialPlThreeYearsResponse) GetUpToDateReasonsOk() (*[]JournalsResponseJournalsUpToDateReasonsInner, bool)`

GetUpToDateReasonsOk returns a tuple with the UpToDateReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDateReasons

`func (o *TrialPlThreeYearsResponse) SetUpToDateReasons(v []JournalsResponseJournalsUpToDateReasonsInner)`

SetUpToDateReasons sets UpToDateReasons field to given value.

### HasUpToDateReasons

`func (o *TrialPlThreeYearsResponse) HasUpToDateReasons() bool`

HasUpToDateReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


