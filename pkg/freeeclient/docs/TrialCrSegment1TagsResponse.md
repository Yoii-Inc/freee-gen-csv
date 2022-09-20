# TrialCrSegment1TagsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialCrSegment1Tags** | [**TrialCrSegment1TagsResponseTrialCrSegment1Tags**](TrialCrSegment1TagsResponseTrialCrSegment1Tags.md) |  | 
**UpToDate** | **bool** | 集計結果が最新かどうか | 
**UpToDateReasons** | Pointer to [**[]JournalsResponseJournalsUpToDateReasonsInner**](JournalsResponseJournalsUpToDateReasonsInner.md) | 集計が最新でない場合の要因情報 | [optional] 

## Methods

### NewTrialCrSegment1TagsResponse

`func NewTrialCrSegment1TagsResponse(trialCrSegment1Tags TrialCrSegment1TagsResponseTrialCrSegment1Tags, upToDate bool, ) *TrialCrSegment1TagsResponse`

NewTrialCrSegment1TagsResponse instantiates a new TrialCrSegment1TagsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialCrSegment1TagsResponseWithDefaults

`func NewTrialCrSegment1TagsResponseWithDefaults() *TrialCrSegment1TagsResponse`

NewTrialCrSegment1TagsResponseWithDefaults instantiates a new TrialCrSegment1TagsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialCrSegment1Tags

`func (o *TrialCrSegment1TagsResponse) GetTrialCrSegment1Tags() TrialCrSegment1TagsResponseTrialCrSegment1Tags`

GetTrialCrSegment1Tags returns the TrialCrSegment1Tags field if non-nil, zero value otherwise.

### GetTrialCrSegment1TagsOk

`func (o *TrialCrSegment1TagsResponse) GetTrialCrSegment1TagsOk() (*TrialCrSegment1TagsResponseTrialCrSegment1Tags, bool)`

GetTrialCrSegment1TagsOk returns a tuple with the TrialCrSegment1Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialCrSegment1Tags

`func (o *TrialCrSegment1TagsResponse) SetTrialCrSegment1Tags(v TrialCrSegment1TagsResponseTrialCrSegment1Tags)`

SetTrialCrSegment1Tags sets TrialCrSegment1Tags field to given value.


### GetUpToDate

`func (o *TrialCrSegment1TagsResponse) GetUpToDate() bool`

GetUpToDate returns the UpToDate field if non-nil, zero value otherwise.

### GetUpToDateOk

`func (o *TrialCrSegment1TagsResponse) GetUpToDateOk() (*bool, bool)`

GetUpToDateOk returns a tuple with the UpToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDate

`func (o *TrialCrSegment1TagsResponse) SetUpToDate(v bool)`

SetUpToDate sets UpToDate field to given value.


### GetUpToDateReasons

`func (o *TrialCrSegment1TagsResponse) GetUpToDateReasons() []JournalsResponseJournalsUpToDateReasonsInner`

GetUpToDateReasons returns the UpToDateReasons field if non-nil, zero value otherwise.

### GetUpToDateReasonsOk

`func (o *TrialCrSegment1TagsResponse) GetUpToDateReasonsOk() (*[]JournalsResponseJournalsUpToDateReasonsInner, bool)`

GetUpToDateReasonsOk returns a tuple with the UpToDateReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDateReasons

`func (o *TrialCrSegment1TagsResponse) SetUpToDateReasons(v []JournalsResponseJournalsUpToDateReasonsInner)`

SetUpToDateReasons sets UpToDateReasons field to given value.

### HasUpToDateReasons

`func (o *TrialCrSegment1TagsResponse) HasUpToDateReasons() bool`

HasUpToDateReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


