# TrialPlSegment1TagsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialPlSegment1Tags** | [**TrialCrSegment1TagsResponseTrialCrSegment1Tags**](TrialCrSegment1TagsResponseTrialCrSegment1Tags.md) |  | 
**UpToDate** | **bool** | 集計結果が最新かどうか | 
**UpToDateReasons** | Pointer to [**[]JournalsResponseJournalsUpToDateReasonsInner**](JournalsResponseJournalsUpToDateReasonsInner.md) | 集計が最新でない場合の要因情報 | [optional] 

## Methods

### NewTrialPlSegment1TagsResponse

`func NewTrialPlSegment1TagsResponse(trialPlSegment1Tags TrialCrSegment1TagsResponseTrialCrSegment1Tags, upToDate bool, ) *TrialPlSegment1TagsResponse`

NewTrialPlSegment1TagsResponse instantiates a new TrialPlSegment1TagsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialPlSegment1TagsResponseWithDefaults

`func NewTrialPlSegment1TagsResponseWithDefaults() *TrialPlSegment1TagsResponse`

NewTrialPlSegment1TagsResponseWithDefaults instantiates a new TrialPlSegment1TagsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialPlSegment1Tags

`func (o *TrialPlSegment1TagsResponse) GetTrialPlSegment1Tags() TrialCrSegment1TagsResponseTrialCrSegment1Tags`

GetTrialPlSegment1Tags returns the TrialPlSegment1Tags field if non-nil, zero value otherwise.

### GetTrialPlSegment1TagsOk

`func (o *TrialPlSegment1TagsResponse) GetTrialPlSegment1TagsOk() (*TrialCrSegment1TagsResponseTrialCrSegment1Tags, bool)`

GetTrialPlSegment1TagsOk returns a tuple with the TrialPlSegment1Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPlSegment1Tags

`func (o *TrialPlSegment1TagsResponse) SetTrialPlSegment1Tags(v TrialCrSegment1TagsResponseTrialCrSegment1Tags)`

SetTrialPlSegment1Tags sets TrialPlSegment1Tags field to given value.


### GetUpToDate

`func (o *TrialPlSegment1TagsResponse) GetUpToDate() bool`

GetUpToDate returns the UpToDate field if non-nil, zero value otherwise.

### GetUpToDateOk

`func (o *TrialPlSegment1TagsResponse) GetUpToDateOk() (*bool, bool)`

GetUpToDateOk returns a tuple with the UpToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDate

`func (o *TrialPlSegment1TagsResponse) SetUpToDate(v bool)`

SetUpToDate sets UpToDate field to given value.


### GetUpToDateReasons

`func (o *TrialPlSegment1TagsResponse) GetUpToDateReasons() []JournalsResponseJournalsUpToDateReasonsInner`

GetUpToDateReasons returns the UpToDateReasons field if non-nil, zero value otherwise.

### GetUpToDateReasonsOk

`func (o *TrialPlSegment1TagsResponse) GetUpToDateReasonsOk() (*[]JournalsResponseJournalsUpToDateReasonsInner, bool)`

GetUpToDateReasonsOk returns a tuple with the UpToDateReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDateReasons

`func (o *TrialPlSegment1TagsResponse) SetUpToDateReasons(v []JournalsResponseJournalsUpToDateReasonsInner)`

SetUpToDateReasons sets UpToDateReasons field to given value.

### HasUpToDateReasons

`func (o *TrialPlSegment1TagsResponse) HasUpToDateReasons() bool`

HasUpToDateReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


