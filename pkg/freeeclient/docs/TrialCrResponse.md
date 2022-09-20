# TrialCrResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialCr** | [**TrialCrResponseTrialCr**](TrialCrResponseTrialCr.md) |  | 
**UpToDate** | **bool** | 集計結果が最新かどうか | 
**UpToDateReasons** | Pointer to [**[]JournalsResponseJournalsUpToDateReasonsInner**](JournalsResponseJournalsUpToDateReasonsInner.md) | 集計が最新でない場合の要因情報 | [optional] 

## Methods

### NewTrialCrResponse

`func NewTrialCrResponse(trialCr TrialCrResponseTrialCr, upToDate bool, ) *TrialCrResponse`

NewTrialCrResponse instantiates a new TrialCrResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialCrResponseWithDefaults

`func NewTrialCrResponseWithDefaults() *TrialCrResponse`

NewTrialCrResponseWithDefaults instantiates a new TrialCrResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialCr

`func (o *TrialCrResponse) GetTrialCr() TrialCrResponseTrialCr`

GetTrialCr returns the TrialCr field if non-nil, zero value otherwise.

### GetTrialCrOk

`func (o *TrialCrResponse) GetTrialCrOk() (*TrialCrResponseTrialCr, bool)`

GetTrialCrOk returns a tuple with the TrialCr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialCr

`func (o *TrialCrResponse) SetTrialCr(v TrialCrResponseTrialCr)`

SetTrialCr sets TrialCr field to given value.


### GetUpToDate

`func (o *TrialCrResponse) GetUpToDate() bool`

GetUpToDate returns the UpToDate field if non-nil, zero value otherwise.

### GetUpToDateOk

`func (o *TrialCrResponse) GetUpToDateOk() (*bool, bool)`

GetUpToDateOk returns a tuple with the UpToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDate

`func (o *TrialCrResponse) SetUpToDate(v bool)`

SetUpToDate sets UpToDate field to given value.


### GetUpToDateReasons

`func (o *TrialCrResponse) GetUpToDateReasons() []JournalsResponseJournalsUpToDateReasonsInner`

GetUpToDateReasons returns the UpToDateReasons field if non-nil, zero value otherwise.

### GetUpToDateReasonsOk

`func (o *TrialCrResponse) GetUpToDateReasonsOk() (*[]JournalsResponseJournalsUpToDateReasonsInner, bool)`

GetUpToDateReasonsOk returns a tuple with the UpToDateReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDateReasons

`func (o *TrialCrResponse) SetUpToDateReasons(v []JournalsResponseJournalsUpToDateReasonsInner)`

SetUpToDateReasons sets UpToDateReasons field to given value.

### HasUpToDateReasons

`func (o *TrialCrResponse) HasUpToDateReasons() bool`

HasUpToDateReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


