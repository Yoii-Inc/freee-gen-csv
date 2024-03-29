# Section

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **bool** | 部門の使用設定（true: 使用する、false: 使用しない） &lt;br&gt; &lt;ul&gt;   &lt;li&gt;     本APIでsectionを作成した場合はtrueになります。   &lt;/li&gt;   &lt;li&gt;     falseにする場合はWeb画面から変更できます。   &lt;/li&gt;   &lt;li&gt;     trueの場合、Web画面での取引登録時などに入力候補として表示されます。   &lt;/li&gt;   &lt;li&gt;     falseの場合、部門自体は削除せず、Web画面での取引登録時などに入力候補として表示されません。ただし取引（収入／支出）の作成APIなどでfalseの部門をパラメータに指定すれば、取引などにfalseの部門を設定できます。   &lt;/li&gt; &lt;/ul&gt; | 
**CompanyId** | **int64** | 事業所ID | 
**Id** | **int64** | 部門ID | 
**IndentCount** | Pointer to **int64** | &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;https://support.freee.co.jp/hc/ja/articles/209093566\&quot;&gt;部門階層&lt;/a&gt; &lt;br&gt; ※ indent_count が 0 のときは第一階層の親部門です。  | [optional] 
**LongName** | Pointer to **NullableString** | 正式名称（255文字以内） | [optional] 
**Name** | **string** | 部門名 (30文字以内) | 
**ParentId** | Pointer to **NullableInt64** | &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;https://support.freee.co.jp/hc/ja/articles/209093566\&quot;&gt;親部門ID&lt;/a&gt; &lt;br&gt; ※ parent_id が null のときは第一階層の親部門です。  | [optional] 
**Shortcut1** | Pointer to **NullableString** | ショートカット１ (20文字以内) | [optional] 
**Shortcut2** | Pointer to **NullableString** | ショートカット２ (20文字以内) | [optional] 

## Methods

### NewSection

`func NewSection(available bool, companyId int64, id int64, name string, ) *Section`

NewSection instantiates a new Section object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectionWithDefaults

`func NewSectionWithDefaults() *Section`

NewSectionWithDefaults instantiates a new Section object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *Section) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *Section) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *Section) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetCompanyId

`func (o *Section) GetCompanyId() int64`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *Section) GetCompanyIdOk() (*int64, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *Section) SetCompanyId(v int64)`

SetCompanyId sets CompanyId field to given value.


### GetId

`func (o *Section) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Section) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Section) SetId(v int64)`

SetId sets Id field to given value.


### GetIndentCount

`func (o *Section) GetIndentCount() int64`

GetIndentCount returns the IndentCount field if non-nil, zero value otherwise.

### GetIndentCountOk

`func (o *Section) GetIndentCountOk() (*int64, bool)`

GetIndentCountOk returns a tuple with the IndentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndentCount

`func (o *Section) SetIndentCount(v int64)`

SetIndentCount sets IndentCount field to given value.

### HasIndentCount

`func (o *Section) HasIndentCount() bool`

HasIndentCount returns a boolean if a field has been set.

### GetLongName

`func (o *Section) GetLongName() string`

GetLongName returns the LongName field if non-nil, zero value otherwise.

### GetLongNameOk

`func (o *Section) GetLongNameOk() (*string, bool)`

GetLongNameOk returns a tuple with the LongName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongName

`func (o *Section) SetLongName(v string)`

SetLongName sets LongName field to given value.

### HasLongName

`func (o *Section) HasLongName() bool`

HasLongName returns a boolean if a field has been set.

### SetLongNameNil

`func (o *Section) SetLongNameNil(b bool)`

 SetLongNameNil sets the value for LongName to be an explicit nil

### UnsetLongName
`func (o *Section) UnsetLongName()`

UnsetLongName ensures that no value is present for LongName, not even an explicit nil
### GetName

`func (o *Section) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Section) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Section) SetName(v string)`

SetName sets Name field to given value.


### GetParentId

`func (o *Section) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Section) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Section) SetParentId(v int64)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Section) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *Section) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *Section) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetShortcut1

`func (o *Section) GetShortcut1() string`

GetShortcut1 returns the Shortcut1 field if non-nil, zero value otherwise.

### GetShortcut1Ok

`func (o *Section) GetShortcut1Ok() (*string, bool)`

GetShortcut1Ok returns a tuple with the Shortcut1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut1

`func (o *Section) SetShortcut1(v string)`

SetShortcut1 sets Shortcut1 field to given value.

### HasShortcut1

`func (o *Section) HasShortcut1() bool`

HasShortcut1 returns a boolean if a field has been set.

### SetShortcut1Nil

`func (o *Section) SetShortcut1Nil(b bool)`

 SetShortcut1Nil sets the value for Shortcut1 to be an explicit nil

### UnsetShortcut1
`func (o *Section) UnsetShortcut1()`

UnsetShortcut1 ensures that no value is present for Shortcut1, not even an explicit nil
### GetShortcut2

`func (o *Section) GetShortcut2() string`

GetShortcut2 returns the Shortcut2 field if non-nil, zero value otherwise.

### GetShortcut2Ok

`func (o *Section) GetShortcut2Ok() (*string, bool)`

GetShortcut2Ok returns a tuple with the Shortcut2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut2

`func (o *Section) SetShortcut2(v string)`

SetShortcut2 sets Shortcut2 field to given value.

### HasShortcut2

`func (o *Section) HasShortcut2() bool`

HasShortcut2 returns a boolean if a field has been set.

### SetShortcut2Nil

`func (o *Section) SetShortcut2Nil(b bool)`

 SetShortcut2Nil sets the value for Shortcut2 to be an explicit nil

### UnsetShortcut2
`func (o *Section) UnsetShortcut2()`

UnsetShortcut2 ensures that no value is present for Shortcut2, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


