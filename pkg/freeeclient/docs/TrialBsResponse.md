# TrialBsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialBs** | [**TrialBsResponseTrialBs**](TrialBsResponseTrialBs.md) |  | 
**UpToDate** | **bool** | 集計結果が最新かどうか | 
**UpToDateReasons** | Pointer to [**[]JournalsResponseJournalsUpToDateReasonsInner**](JournalsResponseJournalsUpToDateReasonsInner.md) | 集計が最新でない場合の要因情報 | [optional] 

## Methods

### NewTrialBsResponse

`func NewTrialBsResponse(trialBs TrialBsResponseTrialBs, upToDate bool, ) *TrialBsResponse`

NewTrialBsResponse instantiates a new TrialBsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialBsResponseWithDefaults

`func NewTrialBsResponseWithDefaults() *TrialBsResponse`

NewTrialBsResponseWithDefaults instantiates a new TrialBsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialBs

`func (o *TrialBsResponse) GetTrialBs() TrialBsResponseTrialBs`

GetTrialBs returns the TrialBs field if non-nil, zero value otherwise.

### GetTrialBsOk

`func (o *TrialBsResponse) GetTrialBsOk() (*TrialBsResponseTrialBs, bool)`

GetTrialBsOk returns a tuple with the TrialBs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialBs

`func (o *TrialBsResponse) SetTrialBs(v TrialBsResponseTrialBs)`

SetTrialBs sets TrialBs field to given value.


### GetUpToDate

`func (o *TrialBsResponse) GetUpToDate() bool`

GetUpToDate returns the UpToDate field if non-nil, zero value otherwise.

### GetUpToDateOk

`func (o *TrialBsResponse) GetUpToDateOk() (*bool, bool)`

GetUpToDateOk returns a tuple with the UpToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDate

`func (o *TrialBsResponse) SetUpToDate(v bool)`

SetUpToDate sets UpToDate field to given value.


### GetUpToDateReasons

`func (o *TrialBsResponse) GetUpToDateReasons() []JournalsResponseJournalsUpToDateReasonsInner`

GetUpToDateReasons returns the UpToDateReasons field if non-nil, zero value otherwise.

### GetUpToDateReasonsOk

`func (o *TrialBsResponse) GetUpToDateReasonsOk() (*[]JournalsResponseJournalsUpToDateReasonsInner, bool)`

GetUpToDateReasonsOk returns a tuple with the UpToDateReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDateReasons

`func (o *TrialBsResponse) SetUpToDateReasons(v []JournalsResponseJournalsUpToDateReasonsInner)`

SetUpToDateReasons sets UpToDateReasons field to given value.

### HasUpToDateReasons

`func (o *TrialBsResponse) HasUpToDateReasons() bool`

HasUpToDateReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


