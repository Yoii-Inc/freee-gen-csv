# TrialCrSectionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialCrSections** | [**TrialCrSectionsResponseTrialCrSections**](TrialCrSectionsResponseTrialCrSections.md) |  | 
**UpToDate** | **bool** | 集計結果が最新かどうか | 
**UpToDateReasons** | Pointer to [**[]JournalsResponseJournalsUpToDateReasonsInner**](JournalsResponseJournalsUpToDateReasonsInner.md) | 集計が最新でない場合の要因情報 | [optional] 

## Methods

### NewTrialCrSectionsResponse

`func NewTrialCrSectionsResponse(trialCrSections TrialCrSectionsResponseTrialCrSections, upToDate bool, ) *TrialCrSectionsResponse`

NewTrialCrSectionsResponse instantiates a new TrialCrSectionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialCrSectionsResponseWithDefaults

`func NewTrialCrSectionsResponseWithDefaults() *TrialCrSectionsResponse`

NewTrialCrSectionsResponseWithDefaults instantiates a new TrialCrSectionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialCrSections

`func (o *TrialCrSectionsResponse) GetTrialCrSections() TrialCrSectionsResponseTrialCrSections`

GetTrialCrSections returns the TrialCrSections field if non-nil, zero value otherwise.

### GetTrialCrSectionsOk

`func (o *TrialCrSectionsResponse) GetTrialCrSectionsOk() (*TrialCrSectionsResponseTrialCrSections, bool)`

GetTrialCrSectionsOk returns a tuple with the TrialCrSections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialCrSections

`func (o *TrialCrSectionsResponse) SetTrialCrSections(v TrialCrSectionsResponseTrialCrSections)`

SetTrialCrSections sets TrialCrSections field to given value.


### GetUpToDate

`func (o *TrialCrSectionsResponse) GetUpToDate() bool`

GetUpToDate returns the UpToDate field if non-nil, zero value otherwise.

### GetUpToDateOk

`func (o *TrialCrSectionsResponse) GetUpToDateOk() (*bool, bool)`

GetUpToDateOk returns a tuple with the UpToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDate

`func (o *TrialCrSectionsResponse) SetUpToDate(v bool)`

SetUpToDate sets UpToDate field to given value.


### GetUpToDateReasons

`func (o *TrialCrSectionsResponse) GetUpToDateReasons() []JournalsResponseJournalsUpToDateReasonsInner`

GetUpToDateReasons returns the UpToDateReasons field if non-nil, zero value otherwise.

### GetUpToDateReasonsOk

`func (o *TrialCrSectionsResponse) GetUpToDateReasonsOk() (*[]JournalsResponseJournalsUpToDateReasonsInner, bool)`

GetUpToDateReasonsOk returns a tuple with the UpToDateReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDateReasons

`func (o *TrialCrSectionsResponse) SetUpToDateReasons(v []JournalsResponseJournalsUpToDateReasonsInner)`

SetUpToDateReasons sets UpToDateReasons field to given value.

### HasUpToDateReasons

`func (o *TrialCrSectionsResponse) HasUpToDateReasons() bool`

HasUpToDateReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


