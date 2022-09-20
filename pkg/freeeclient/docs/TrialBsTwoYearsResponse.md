# TrialBsTwoYearsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialBsTwoYears** | [**TrialBsTwoYearsResponseTrialBsTwoYears**](TrialBsTwoYearsResponseTrialBsTwoYears.md) |  | 
**UpToDate** | **bool** | 集計結果が最新かどうか | 
**UpToDateReasons** | Pointer to [**[]JournalsResponseJournalsUpToDateReasonsInner**](JournalsResponseJournalsUpToDateReasonsInner.md) | 集計が最新でない場合の要因情報 | [optional] 

## Methods

### NewTrialBsTwoYearsResponse

`func NewTrialBsTwoYearsResponse(trialBsTwoYears TrialBsTwoYearsResponseTrialBsTwoYears, upToDate bool, ) *TrialBsTwoYearsResponse`

NewTrialBsTwoYearsResponse instantiates a new TrialBsTwoYearsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialBsTwoYearsResponseWithDefaults

`func NewTrialBsTwoYearsResponseWithDefaults() *TrialBsTwoYearsResponse`

NewTrialBsTwoYearsResponseWithDefaults instantiates a new TrialBsTwoYearsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialBsTwoYears

`func (o *TrialBsTwoYearsResponse) GetTrialBsTwoYears() TrialBsTwoYearsResponseTrialBsTwoYears`

GetTrialBsTwoYears returns the TrialBsTwoYears field if non-nil, zero value otherwise.

### GetTrialBsTwoYearsOk

`func (o *TrialBsTwoYearsResponse) GetTrialBsTwoYearsOk() (*TrialBsTwoYearsResponseTrialBsTwoYears, bool)`

GetTrialBsTwoYearsOk returns a tuple with the TrialBsTwoYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialBsTwoYears

`func (o *TrialBsTwoYearsResponse) SetTrialBsTwoYears(v TrialBsTwoYearsResponseTrialBsTwoYears)`

SetTrialBsTwoYears sets TrialBsTwoYears field to given value.


### GetUpToDate

`func (o *TrialBsTwoYearsResponse) GetUpToDate() bool`

GetUpToDate returns the UpToDate field if non-nil, zero value otherwise.

### GetUpToDateOk

`func (o *TrialBsTwoYearsResponse) GetUpToDateOk() (*bool, bool)`

GetUpToDateOk returns a tuple with the UpToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDate

`func (o *TrialBsTwoYearsResponse) SetUpToDate(v bool)`

SetUpToDate sets UpToDate field to given value.


### GetUpToDateReasons

`func (o *TrialBsTwoYearsResponse) GetUpToDateReasons() []JournalsResponseJournalsUpToDateReasonsInner`

GetUpToDateReasons returns the UpToDateReasons field if non-nil, zero value otherwise.

### GetUpToDateReasonsOk

`func (o *TrialBsTwoYearsResponse) GetUpToDateReasonsOk() (*[]JournalsResponseJournalsUpToDateReasonsInner, bool)`

GetUpToDateReasonsOk returns a tuple with the UpToDateReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDateReasons

`func (o *TrialBsTwoYearsResponse) SetUpToDateReasons(v []JournalsResponseJournalsUpToDateReasonsInner)`

SetUpToDateReasons sets UpToDateReasons field to given value.

### HasUpToDateReasons

`func (o *TrialBsTwoYearsResponse) HasUpToDateReasons() bool`

HasUpToDateReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


