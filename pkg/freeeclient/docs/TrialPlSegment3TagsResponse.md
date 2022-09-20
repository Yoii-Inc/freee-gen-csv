# TrialPlSegment3TagsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialPlSegment3Tags** | [**TrialCrSegment3TagsResponseTrialCrSegment3Tags**](TrialCrSegment3TagsResponseTrialCrSegment3Tags.md) |  | 
**UpToDate** | **bool** | 集計結果が最新かどうか | 
**UpToDateReasons** | Pointer to [**[]JournalsResponseJournalsUpToDateReasonsInner**](JournalsResponseJournalsUpToDateReasonsInner.md) | 集計が最新でない場合の要因情報 | [optional] 

## Methods

### NewTrialPlSegment3TagsResponse

`func NewTrialPlSegment3TagsResponse(trialPlSegment3Tags TrialCrSegment3TagsResponseTrialCrSegment3Tags, upToDate bool, ) *TrialPlSegment3TagsResponse`

NewTrialPlSegment3TagsResponse instantiates a new TrialPlSegment3TagsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialPlSegment3TagsResponseWithDefaults

`func NewTrialPlSegment3TagsResponseWithDefaults() *TrialPlSegment3TagsResponse`

NewTrialPlSegment3TagsResponseWithDefaults instantiates a new TrialPlSegment3TagsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialPlSegment3Tags

`func (o *TrialPlSegment3TagsResponse) GetTrialPlSegment3Tags() TrialCrSegment3TagsResponseTrialCrSegment3Tags`

GetTrialPlSegment3Tags returns the TrialPlSegment3Tags field if non-nil, zero value otherwise.

### GetTrialPlSegment3TagsOk

`func (o *TrialPlSegment3TagsResponse) GetTrialPlSegment3TagsOk() (*TrialCrSegment3TagsResponseTrialCrSegment3Tags, bool)`

GetTrialPlSegment3TagsOk returns a tuple with the TrialPlSegment3Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPlSegment3Tags

`func (o *TrialPlSegment3TagsResponse) SetTrialPlSegment3Tags(v TrialCrSegment3TagsResponseTrialCrSegment3Tags)`

SetTrialPlSegment3Tags sets TrialPlSegment3Tags field to given value.


### GetUpToDate

`func (o *TrialPlSegment3TagsResponse) GetUpToDate() bool`

GetUpToDate returns the UpToDate field if non-nil, zero value otherwise.

### GetUpToDateOk

`func (o *TrialPlSegment3TagsResponse) GetUpToDateOk() (*bool, bool)`

GetUpToDateOk returns a tuple with the UpToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDate

`func (o *TrialPlSegment3TagsResponse) SetUpToDate(v bool)`

SetUpToDate sets UpToDate field to given value.


### GetUpToDateReasons

`func (o *TrialPlSegment3TagsResponse) GetUpToDateReasons() []JournalsResponseJournalsUpToDateReasonsInner`

GetUpToDateReasons returns the UpToDateReasons field if non-nil, zero value otherwise.

### GetUpToDateReasonsOk

`func (o *TrialPlSegment3TagsResponse) GetUpToDateReasonsOk() (*[]JournalsResponseJournalsUpToDateReasonsInner, bool)`

GetUpToDateReasonsOk returns a tuple with the UpToDateReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDateReasons

`func (o *TrialPlSegment3TagsResponse) SetUpToDateReasons(v []JournalsResponseJournalsUpToDateReasonsInner)`

SetUpToDateReasons sets UpToDateReasons field to given value.

### HasUpToDateReasons

`func (o *TrialPlSegment3TagsResponse) HasUpToDateReasons() bool`

HasUpToDateReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


