# TrialPlSectionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialPlSections** | [**TrialCrSectionsResponseTrialCrSections**](TrialCrSectionsResponseTrialCrSections.md) |  | 
**UpToDate** | **bool** | 集計結果が最新かどうか | 
**UpToDateReasons** | Pointer to [**[]JournalsResponseJournalsUpToDateReasonsInner**](JournalsResponseJournalsUpToDateReasonsInner.md) | 集計が最新でない場合の要因情報 | [optional] 

## Methods

### NewTrialPlSectionsResponse

`func NewTrialPlSectionsResponse(trialPlSections TrialCrSectionsResponseTrialCrSections, upToDate bool, ) *TrialPlSectionsResponse`

NewTrialPlSectionsResponse instantiates a new TrialPlSectionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialPlSectionsResponseWithDefaults

`func NewTrialPlSectionsResponseWithDefaults() *TrialPlSectionsResponse`

NewTrialPlSectionsResponseWithDefaults instantiates a new TrialPlSectionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialPlSections

`func (o *TrialPlSectionsResponse) GetTrialPlSections() TrialCrSectionsResponseTrialCrSections`

GetTrialPlSections returns the TrialPlSections field if non-nil, zero value otherwise.

### GetTrialPlSectionsOk

`func (o *TrialPlSectionsResponse) GetTrialPlSectionsOk() (*TrialCrSectionsResponseTrialCrSections, bool)`

GetTrialPlSectionsOk returns a tuple with the TrialPlSections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPlSections

`func (o *TrialPlSectionsResponse) SetTrialPlSections(v TrialCrSectionsResponseTrialCrSections)`

SetTrialPlSections sets TrialPlSections field to given value.


### GetUpToDate

`func (o *TrialPlSectionsResponse) GetUpToDate() bool`

GetUpToDate returns the UpToDate field if non-nil, zero value otherwise.

### GetUpToDateOk

`func (o *TrialPlSectionsResponse) GetUpToDateOk() (*bool, bool)`

GetUpToDateOk returns a tuple with the UpToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDate

`func (o *TrialPlSectionsResponse) SetUpToDate(v bool)`

SetUpToDate sets UpToDate field to given value.


### GetUpToDateReasons

`func (o *TrialPlSectionsResponse) GetUpToDateReasons() []JournalsResponseJournalsUpToDateReasonsInner`

GetUpToDateReasons returns the UpToDateReasons field if non-nil, zero value otherwise.

### GetUpToDateReasonsOk

`func (o *TrialPlSectionsResponse) GetUpToDateReasonsOk() (*[]JournalsResponseJournalsUpToDateReasonsInner, bool)`

GetUpToDateReasonsOk returns a tuple with the UpToDateReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDateReasons

`func (o *TrialPlSectionsResponse) SetUpToDateReasons(v []JournalsResponseJournalsUpToDateReasonsInner)`

SetUpToDateReasons sets UpToDateReasons field to given value.

### HasUpToDateReasons

`func (o *TrialPlSectionsResponse) HasUpToDateReasons() bool`

HasUpToDateReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


