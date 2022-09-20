# TrialPlResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialPl** | [**TrialCrResponseTrialCr**](TrialCrResponseTrialCr.md) |  | 
**UpToDate** | **bool** | 集計結果が最新かどうか | 
**UpToDateReasons** | Pointer to [**[]JournalsResponseJournalsUpToDateReasonsInner**](JournalsResponseJournalsUpToDateReasonsInner.md) | 集計が最新でない場合の要因情報 | [optional] 

## Methods

### NewTrialPlResponse

`func NewTrialPlResponse(trialPl TrialCrResponseTrialCr, upToDate bool, ) *TrialPlResponse`

NewTrialPlResponse instantiates a new TrialPlResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialPlResponseWithDefaults

`func NewTrialPlResponseWithDefaults() *TrialPlResponse`

NewTrialPlResponseWithDefaults instantiates a new TrialPlResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialPl

`func (o *TrialPlResponse) GetTrialPl() TrialCrResponseTrialCr`

GetTrialPl returns the TrialPl field if non-nil, zero value otherwise.

### GetTrialPlOk

`func (o *TrialPlResponse) GetTrialPlOk() (*TrialCrResponseTrialCr, bool)`

GetTrialPlOk returns a tuple with the TrialPl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPl

`func (o *TrialPlResponse) SetTrialPl(v TrialCrResponseTrialCr)`

SetTrialPl sets TrialPl field to given value.


### GetUpToDate

`func (o *TrialPlResponse) GetUpToDate() bool`

GetUpToDate returns the UpToDate field if non-nil, zero value otherwise.

### GetUpToDateOk

`func (o *TrialPlResponse) GetUpToDateOk() (*bool, bool)`

GetUpToDateOk returns a tuple with the UpToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDate

`func (o *TrialPlResponse) SetUpToDate(v bool)`

SetUpToDate sets UpToDate field to given value.


### GetUpToDateReasons

`func (o *TrialPlResponse) GetUpToDateReasons() []JournalsResponseJournalsUpToDateReasonsInner`

GetUpToDateReasons returns the UpToDateReasons field if non-nil, zero value otherwise.

### GetUpToDateReasonsOk

`func (o *TrialPlResponse) GetUpToDateReasonsOk() (*[]JournalsResponseJournalsUpToDateReasonsInner, bool)`

GetUpToDateReasonsOk returns a tuple with the UpToDateReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDateReasons

`func (o *TrialPlResponse) SetUpToDateReasons(v []JournalsResponseJournalsUpToDateReasonsInner)`

SetUpToDateReasons sets UpToDateReasons field to given value.

### HasUpToDateReasons

`func (o *TrialPlResponse) HasUpToDateReasons() bool`

HasUpToDateReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


