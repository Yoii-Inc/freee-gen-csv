# TrialCrSegment3TagsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialCrSegment3Tags** | [**TrialCrSegment3TagsResponseTrialCrSegment3Tags**](TrialCrSegment3TagsResponseTrialCrSegment3Tags.md) |  | 
**UpToDate** | **bool** | 集計結果が最新かどうか | 
**UpToDateReasons** | Pointer to [**[]JournalsResponseJournalsUpToDateReasonsInner**](JournalsResponseJournalsUpToDateReasonsInner.md) | 集計が最新でない場合の要因情報 | [optional] 

## Methods

### NewTrialCrSegment3TagsResponse

`func NewTrialCrSegment3TagsResponse(trialCrSegment3Tags TrialCrSegment3TagsResponseTrialCrSegment3Tags, upToDate bool, ) *TrialCrSegment3TagsResponse`

NewTrialCrSegment3TagsResponse instantiates a new TrialCrSegment3TagsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialCrSegment3TagsResponseWithDefaults

`func NewTrialCrSegment3TagsResponseWithDefaults() *TrialCrSegment3TagsResponse`

NewTrialCrSegment3TagsResponseWithDefaults instantiates a new TrialCrSegment3TagsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialCrSegment3Tags

`func (o *TrialCrSegment3TagsResponse) GetTrialCrSegment3Tags() TrialCrSegment3TagsResponseTrialCrSegment3Tags`

GetTrialCrSegment3Tags returns the TrialCrSegment3Tags field if non-nil, zero value otherwise.

### GetTrialCrSegment3TagsOk

`func (o *TrialCrSegment3TagsResponse) GetTrialCrSegment3TagsOk() (*TrialCrSegment3TagsResponseTrialCrSegment3Tags, bool)`

GetTrialCrSegment3TagsOk returns a tuple with the TrialCrSegment3Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialCrSegment3Tags

`func (o *TrialCrSegment3TagsResponse) SetTrialCrSegment3Tags(v TrialCrSegment3TagsResponseTrialCrSegment3Tags)`

SetTrialCrSegment3Tags sets TrialCrSegment3Tags field to given value.


### GetUpToDate

`func (o *TrialCrSegment3TagsResponse) GetUpToDate() bool`

GetUpToDate returns the UpToDate field if non-nil, zero value otherwise.

### GetUpToDateOk

`func (o *TrialCrSegment3TagsResponse) GetUpToDateOk() (*bool, bool)`

GetUpToDateOk returns a tuple with the UpToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDate

`func (o *TrialCrSegment3TagsResponse) SetUpToDate(v bool)`

SetUpToDate sets UpToDate field to given value.


### GetUpToDateReasons

`func (o *TrialCrSegment3TagsResponse) GetUpToDateReasons() []JournalsResponseJournalsUpToDateReasonsInner`

GetUpToDateReasons returns the UpToDateReasons field if non-nil, zero value otherwise.

### GetUpToDateReasonsOk

`func (o *TrialCrSegment3TagsResponse) GetUpToDateReasonsOk() (*[]JournalsResponseJournalsUpToDateReasonsInner, bool)`

GetUpToDateReasonsOk returns a tuple with the UpToDateReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDateReasons

`func (o *TrialCrSegment3TagsResponse) SetUpToDateReasons(v []JournalsResponseJournalsUpToDateReasonsInner)`

SetUpToDateReasons sets UpToDateReasons field to given value.

### HasUpToDateReasons

`func (o *TrialCrSegment3TagsResponse) HasUpToDateReasons() bool`

HasUpToDateReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


