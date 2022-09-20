# JournalsResponseJournals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int64** | 事業所ID | 
**DownloadType** | Pointer to **string** | ダウンロード形式 | [optional] 
**EndDate** | Pointer to **string** | 取得終了日 (yyyy-mm-dd) | [optional] 
**Id** | **int64** | 受け付けID | 
**Messages** | Pointer to **[]string** |  | [optional] 
**StartDate** | Pointer to **string** | 取得開始日 (yyyy-mm-dd) | [optional] 
**StatusUrl** | Pointer to **string** | ステータス確認用URL | [optional] 
**UpToDate** | Pointer to **bool** | 集計結果が最新かどうか | [optional] 
**UpToDateReasons** | Pointer to [**[]JournalsResponseJournalsUpToDateReasonsInner**](JournalsResponseJournalsUpToDateReasonsInner.md) | 集計が最新でない場合の要因情報 | [optional] 
**VisibleIds** | Pointer to **[]string** |  | [optional] 
**VisibleTags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewJournalsResponseJournals

`func NewJournalsResponseJournals(companyId int64, id int64, ) *JournalsResponseJournals`

NewJournalsResponseJournals instantiates a new JournalsResponseJournals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJournalsResponseJournalsWithDefaults

`func NewJournalsResponseJournalsWithDefaults() *JournalsResponseJournals`

NewJournalsResponseJournalsWithDefaults instantiates a new JournalsResponseJournals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *JournalsResponseJournals) GetCompanyId() int64`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *JournalsResponseJournals) GetCompanyIdOk() (*int64, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *JournalsResponseJournals) SetCompanyId(v int64)`

SetCompanyId sets CompanyId field to given value.


### GetDownloadType

`func (o *JournalsResponseJournals) GetDownloadType() string`

GetDownloadType returns the DownloadType field if non-nil, zero value otherwise.

### GetDownloadTypeOk

`func (o *JournalsResponseJournals) GetDownloadTypeOk() (*string, bool)`

GetDownloadTypeOk returns a tuple with the DownloadType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadType

`func (o *JournalsResponseJournals) SetDownloadType(v string)`

SetDownloadType sets DownloadType field to given value.

### HasDownloadType

`func (o *JournalsResponseJournals) HasDownloadType() bool`

HasDownloadType returns a boolean if a field has been set.

### GetEndDate

`func (o *JournalsResponseJournals) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *JournalsResponseJournals) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *JournalsResponseJournals) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *JournalsResponseJournals) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetId

`func (o *JournalsResponseJournals) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JournalsResponseJournals) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JournalsResponseJournals) SetId(v int64)`

SetId sets Id field to given value.


### GetMessages

`func (o *JournalsResponseJournals) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *JournalsResponseJournals) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *JournalsResponseJournals) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *JournalsResponseJournals) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetStartDate

`func (o *JournalsResponseJournals) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *JournalsResponseJournals) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *JournalsResponseJournals) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *JournalsResponseJournals) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatusUrl

`func (o *JournalsResponseJournals) GetStatusUrl() string`

GetStatusUrl returns the StatusUrl field if non-nil, zero value otherwise.

### GetStatusUrlOk

`func (o *JournalsResponseJournals) GetStatusUrlOk() (*string, bool)`

GetStatusUrlOk returns a tuple with the StatusUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUrl

`func (o *JournalsResponseJournals) SetStatusUrl(v string)`

SetStatusUrl sets StatusUrl field to given value.

### HasStatusUrl

`func (o *JournalsResponseJournals) HasStatusUrl() bool`

HasStatusUrl returns a boolean if a field has been set.

### GetUpToDate

`func (o *JournalsResponseJournals) GetUpToDate() bool`

GetUpToDate returns the UpToDate field if non-nil, zero value otherwise.

### GetUpToDateOk

`func (o *JournalsResponseJournals) GetUpToDateOk() (*bool, bool)`

GetUpToDateOk returns a tuple with the UpToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDate

`func (o *JournalsResponseJournals) SetUpToDate(v bool)`

SetUpToDate sets UpToDate field to given value.

### HasUpToDate

`func (o *JournalsResponseJournals) HasUpToDate() bool`

HasUpToDate returns a boolean if a field has been set.

### GetUpToDateReasons

`func (o *JournalsResponseJournals) GetUpToDateReasons() []JournalsResponseJournalsUpToDateReasonsInner`

GetUpToDateReasons returns the UpToDateReasons field if non-nil, zero value otherwise.

### GetUpToDateReasonsOk

`func (o *JournalsResponseJournals) GetUpToDateReasonsOk() (*[]JournalsResponseJournalsUpToDateReasonsInner, bool)`

GetUpToDateReasonsOk returns a tuple with the UpToDateReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDateReasons

`func (o *JournalsResponseJournals) SetUpToDateReasons(v []JournalsResponseJournalsUpToDateReasonsInner)`

SetUpToDateReasons sets UpToDateReasons field to given value.

### HasUpToDateReasons

`func (o *JournalsResponseJournals) HasUpToDateReasons() bool`

HasUpToDateReasons returns a boolean if a field has been set.

### GetVisibleIds

`func (o *JournalsResponseJournals) GetVisibleIds() []string`

GetVisibleIds returns the VisibleIds field if non-nil, zero value otherwise.

### GetVisibleIdsOk

`func (o *JournalsResponseJournals) GetVisibleIdsOk() (*[]string, bool)`

GetVisibleIdsOk returns a tuple with the VisibleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleIds

`func (o *JournalsResponseJournals) SetVisibleIds(v []string)`

SetVisibleIds sets VisibleIds field to given value.

### HasVisibleIds

`func (o *JournalsResponseJournals) HasVisibleIds() bool`

HasVisibleIds returns a boolean if a field has been set.

### GetVisibleTags

`func (o *JournalsResponseJournals) GetVisibleTags() []string`

GetVisibleTags returns the VisibleTags field if non-nil, zero value otherwise.

### GetVisibleTagsOk

`func (o *JournalsResponseJournals) GetVisibleTagsOk() (*[]string, bool)`

GetVisibleTagsOk returns a tuple with the VisibleTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleTags

`func (o *JournalsResponseJournals) SetVisibleTags(v []string)`

SetVisibleTags sets VisibleTags field to given value.

### HasVisibleTags

`func (o *JournalsResponseJournals) HasVisibleTags() bool`

HasVisibleTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


