# TrialBsThreeYearsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialBsThreeYears** | [**TrialBsThreeYearsResponseTrialBsThreeYears**](TrialBsThreeYearsResponseTrialBsThreeYears.md) |  | 
**UpToDate** | **bool** | 集計結果が最新かどうか | 
**UpToDateReasons** | Pointer to [**[]JournalsResponseJournalsUpToDateReasonsInner**](JournalsResponseJournalsUpToDateReasonsInner.md) | 集計が最新でない場合の要因情報 | [optional] 

## Methods

### NewTrialBsThreeYearsResponse

`func NewTrialBsThreeYearsResponse(trialBsThreeYears TrialBsThreeYearsResponseTrialBsThreeYears, upToDate bool, ) *TrialBsThreeYearsResponse`

NewTrialBsThreeYearsResponse instantiates a new TrialBsThreeYearsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialBsThreeYearsResponseWithDefaults

`func NewTrialBsThreeYearsResponseWithDefaults() *TrialBsThreeYearsResponse`

NewTrialBsThreeYearsResponseWithDefaults instantiates a new TrialBsThreeYearsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialBsThreeYears

`func (o *TrialBsThreeYearsResponse) GetTrialBsThreeYears() TrialBsThreeYearsResponseTrialBsThreeYears`

GetTrialBsThreeYears returns the TrialBsThreeYears field if non-nil, zero value otherwise.

### GetTrialBsThreeYearsOk

`func (o *TrialBsThreeYearsResponse) GetTrialBsThreeYearsOk() (*TrialBsThreeYearsResponseTrialBsThreeYears, bool)`

GetTrialBsThreeYearsOk returns a tuple with the TrialBsThreeYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialBsThreeYears

`func (o *TrialBsThreeYearsResponse) SetTrialBsThreeYears(v TrialBsThreeYearsResponseTrialBsThreeYears)`

SetTrialBsThreeYears sets TrialBsThreeYears field to given value.


### GetUpToDate

`func (o *TrialBsThreeYearsResponse) GetUpToDate() bool`

GetUpToDate returns the UpToDate field if non-nil, zero value otherwise.

### GetUpToDateOk

`func (o *TrialBsThreeYearsResponse) GetUpToDateOk() (*bool, bool)`

GetUpToDateOk returns a tuple with the UpToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDate

`func (o *TrialBsThreeYearsResponse) SetUpToDate(v bool)`

SetUpToDate sets UpToDate field to given value.


### GetUpToDateReasons

`func (o *TrialBsThreeYearsResponse) GetUpToDateReasons() []JournalsResponseJournalsUpToDateReasonsInner`

GetUpToDateReasons returns the UpToDateReasons field if non-nil, zero value otherwise.

### GetUpToDateReasonsOk

`func (o *TrialBsThreeYearsResponse) GetUpToDateReasonsOk() (*[]JournalsResponseJournalsUpToDateReasonsInner, bool)`

GetUpToDateReasonsOk returns a tuple with the UpToDateReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDateReasons

`func (o *TrialBsThreeYearsResponse) SetUpToDateReasons(v []JournalsResponseJournalsUpToDateReasonsInner)`

SetUpToDateReasons sets UpToDateReasons field to given value.

### HasUpToDateReasons

`func (o *TrialBsThreeYearsResponse) HasUpToDateReasons() bool`

HasUpToDateReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


