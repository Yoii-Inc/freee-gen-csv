# TrialPlSegment2TagsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialPlSegment2Tags** | [**TrialCrSegment2TagsResponseTrialCrSegment2Tags**](TrialCrSegment2TagsResponseTrialCrSegment2Tags.md) |  | 
**UpToDate** | **bool** | 集計結果が最新かどうか | 
**UpToDateReasons** | Pointer to [**[]JournalsResponseJournalsUpToDateReasonsInner**](JournalsResponseJournalsUpToDateReasonsInner.md) | 集計が最新でない場合の要因情報 | [optional] 

## Methods

### NewTrialPlSegment2TagsResponse

`func NewTrialPlSegment2TagsResponse(trialPlSegment2Tags TrialCrSegment2TagsResponseTrialCrSegment2Tags, upToDate bool, ) *TrialPlSegment2TagsResponse`

NewTrialPlSegment2TagsResponse instantiates a new TrialPlSegment2TagsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialPlSegment2TagsResponseWithDefaults

`func NewTrialPlSegment2TagsResponseWithDefaults() *TrialPlSegment2TagsResponse`

NewTrialPlSegment2TagsResponseWithDefaults instantiates a new TrialPlSegment2TagsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialPlSegment2Tags

`func (o *TrialPlSegment2TagsResponse) GetTrialPlSegment2Tags() TrialCrSegment2TagsResponseTrialCrSegment2Tags`

GetTrialPlSegment2Tags returns the TrialPlSegment2Tags field if non-nil, zero value otherwise.

### GetTrialPlSegment2TagsOk

`func (o *TrialPlSegment2TagsResponse) GetTrialPlSegment2TagsOk() (*TrialCrSegment2TagsResponseTrialCrSegment2Tags, bool)`

GetTrialPlSegment2TagsOk returns a tuple with the TrialPlSegment2Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPlSegment2Tags

`func (o *TrialPlSegment2TagsResponse) SetTrialPlSegment2Tags(v TrialCrSegment2TagsResponseTrialCrSegment2Tags)`

SetTrialPlSegment2Tags sets TrialPlSegment2Tags field to given value.


### GetUpToDate

`func (o *TrialPlSegment2TagsResponse) GetUpToDate() bool`

GetUpToDate returns the UpToDate field if non-nil, zero value otherwise.

### GetUpToDateOk

`func (o *TrialPlSegment2TagsResponse) GetUpToDateOk() (*bool, bool)`

GetUpToDateOk returns a tuple with the UpToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDate

`func (o *TrialPlSegment2TagsResponse) SetUpToDate(v bool)`

SetUpToDate sets UpToDate field to given value.


### GetUpToDateReasons

`func (o *TrialPlSegment2TagsResponse) GetUpToDateReasons() []JournalsResponseJournalsUpToDateReasonsInner`

GetUpToDateReasons returns the UpToDateReasons field if non-nil, zero value otherwise.

### GetUpToDateReasonsOk

`func (o *TrialPlSegment2TagsResponse) GetUpToDateReasonsOk() (*[]JournalsResponseJournalsUpToDateReasonsInner, bool)`

GetUpToDateReasonsOk returns a tuple with the UpToDateReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDateReasons

`func (o *TrialPlSegment2TagsResponse) SetUpToDateReasons(v []JournalsResponseJournalsUpToDateReasonsInner)`

SetUpToDateReasons sets UpToDateReasons field to given value.

### HasUpToDateReasons

`func (o *TrialPlSegment2TagsResponse) HasUpToDateReasons() bool`

HasUpToDateReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


