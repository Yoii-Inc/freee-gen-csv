# TrialCrSegment2TagsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialCrSegment2Tags** | [**TrialCrSegment2TagsResponseTrialCrSegment2Tags**](TrialCrSegment2TagsResponseTrialCrSegment2Tags.md) |  | 
**UpToDate** | **bool** | 集計結果が最新かどうか | 
**UpToDateReasons** | Pointer to [**[]JournalsResponseJournalsUpToDateReasonsInner**](JournalsResponseJournalsUpToDateReasonsInner.md) | 集計が最新でない場合の要因情報 | [optional] 

## Methods

### NewTrialCrSegment2TagsResponse

`func NewTrialCrSegment2TagsResponse(trialCrSegment2Tags TrialCrSegment2TagsResponseTrialCrSegment2Tags, upToDate bool, ) *TrialCrSegment2TagsResponse`

NewTrialCrSegment2TagsResponse instantiates a new TrialCrSegment2TagsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialCrSegment2TagsResponseWithDefaults

`func NewTrialCrSegment2TagsResponseWithDefaults() *TrialCrSegment2TagsResponse`

NewTrialCrSegment2TagsResponseWithDefaults instantiates a new TrialCrSegment2TagsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialCrSegment2Tags

`func (o *TrialCrSegment2TagsResponse) GetTrialCrSegment2Tags() TrialCrSegment2TagsResponseTrialCrSegment2Tags`

GetTrialCrSegment2Tags returns the TrialCrSegment2Tags field if non-nil, zero value otherwise.

### GetTrialCrSegment2TagsOk

`func (o *TrialCrSegment2TagsResponse) GetTrialCrSegment2TagsOk() (*TrialCrSegment2TagsResponseTrialCrSegment2Tags, bool)`

GetTrialCrSegment2TagsOk returns a tuple with the TrialCrSegment2Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialCrSegment2Tags

`func (o *TrialCrSegment2TagsResponse) SetTrialCrSegment2Tags(v TrialCrSegment2TagsResponseTrialCrSegment2Tags)`

SetTrialCrSegment2Tags sets TrialCrSegment2Tags field to given value.


### GetUpToDate

`func (o *TrialCrSegment2TagsResponse) GetUpToDate() bool`

GetUpToDate returns the UpToDate field if non-nil, zero value otherwise.

### GetUpToDateOk

`func (o *TrialCrSegment2TagsResponse) GetUpToDateOk() (*bool, bool)`

GetUpToDateOk returns a tuple with the UpToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDate

`func (o *TrialCrSegment2TagsResponse) SetUpToDate(v bool)`

SetUpToDate sets UpToDate field to given value.


### GetUpToDateReasons

`func (o *TrialCrSegment2TagsResponse) GetUpToDateReasons() []JournalsResponseJournalsUpToDateReasonsInner`

GetUpToDateReasons returns the UpToDateReasons field if non-nil, zero value otherwise.

### GetUpToDateReasonsOk

`func (o *TrialCrSegment2TagsResponse) GetUpToDateReasonsOk() (*[]JournalsResponseJournalsUpToDateReasonsInner, bool)`

GetUpToDateReasonsOk returns a tuple with the UpToDateReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDateReasons

`func (o *TrialCrSegment2TagsResponse) SetUpToDateReasons(v []JournalsResponseJournalsUpToDateReasonsInner)`

SetUpToDateReasons sets UpToDateReasons field to given value.

### HasUpToDateReasons

`func (o *TrialCrSegment2TagsResponse) HasUpToDateReasons() bool`

HasUpToDateReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


