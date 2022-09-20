# PartnersResponsePartnersInnerAddressAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrefectureCode** | Pointer to **NullableInt64** | 都道府県コード（-1: 設定しない、0:北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄 | [optional] 
**StreetName1** | Pointer to **NullableString** | 市区町村・番地 | [optional] 
**StreetName2** | Pointer to **NullableString** | 建物名・部屋番号など | [optional] 
**Zipcode** | Pointer to **NullableString** | 郵便番号 | [optional] 

## Methods

### NewPartnersResponsePartnersInnerAddressAttributes

`func NewPartnersResponsePartnersInnerAddressAttributes() *PartnersResponsePartnersInnerAddressAttributes`

NewPartnersResponsePartnersInnerAddressAttributes instantiates a new PartnersResponsePartnersInnerAddressAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnersResponsePartnersInnerAddressAttributesWithDefaults

`func NewPartnersResponsePartnersInnerAddressAttributesWithDefaults() *PartnersResponsePartnersInnerAddressAttributes`

NewPartnersResponsePartnersInnerAddressAttributesWithDefaults instantiates a new PartnersResponsePartnersInnerAddressAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefectureCode

`func (o *PartnersResponsePartnersInnerAddressAttributes) GetPrefectureCode() int64`

GetPrefectureCode returns the PrefectureCode field if non-nil, zero value otherwise.

### GetPrefectureCodeOk

`func (o *PartnersResponsePartnersInnerAddressAttributes) GetPrefectureCodeOk() (*int64, bool)`

GetPrefectureCodeOk returns a tuple with the PrefectureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefectureCode

`func (o *PartnersResponsePartnersInnerAddressAttributes) SetPrefectureCode(v int64)`

SetPrefectureCode sets PrefectureCode field to given value.

### HasPrefectureCode

`func (o *PartnersResponsePartnersInnerAddressAttributes) HasPrefectureCode() bool`

HasPrefectureCode returns a boolean if a field has been set.

### SetPrefectureCodeNil

`func (o *PartnersResponsePartnersInnerAddressAttributes) SetPrefectureCodeNil(b bool)`

 SetPrefectureCodeNil sets the value for PrefectureCode to be an explicit nil

### UnsetPrefectureCode
`func (o *PartnersResponsePartnersInnerAddressAttributes) UnsetPrefectureCode()`

UnsetPrefectureCode ensures that no value is present for PrefectureCode, not even an explicit nil
### GetStreetName1

`func (o *PartnersResponsePartnersInnerAddressAttributes) GetStreetName1() string`

GetStreetName1 returns the StreetName1 field if non-nil, zero value otherwise.

### GetStreetName1Ok

`func (o *PartnersResponsePartnersInnerAddressAttributes) GetStreetName1Ok() (*string, bool)`

GetStreetName1Ok returns a tuple with the StreetName1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetName1

`func (o *PartnersResponsePartnersInnerAddressAttributes) SetStreetName1(v string)`

SetStreetName1 sets StreetName1 field to given value.

### HasStreetName1

`func (o *PartnersResponsePartnersInnerAddressAttributes) HasStreetName1() bool`

HasStreetName1 returns a boolean if a field has been set.

### SetStreetName1Nil

`func (o *PartnersResponsePartnersInnerAddressAttributes) SetStreetName1Nil(b bool)`

 SetStreetName1Nil sets the value for StreetName1 to be an explicit nil

### UnsetStreetName1
`func (o *PartnersResponsePartnersInnerAddressAttributes) UnsetStreetName1()`

UnsetStreetName1 ensures that no value is present for StreetName1, not even an explicit nil
### GetStreetName2

`func (o *PartnersResponsePartnersInnerAddressAttributes) GetStreetName2() string`

GetStreetName2 returns the StreetName2 field if non-nil, zero value otherwise.

### GetStreetName2Ok

`func (o *PartnersResponsePartnersInnerAddressAttributes) GetStreetName2Ok() (*string, bool)`

GetStreetName2Ok returns a tuple with the StreetName2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetName2

`func (o *PartnersResponsePartnersInnerAddressAttributes) SetStreetName2(v string)`

SetStreetName2 sets StreetName2 field to given value.

### HasStreetName2

`func (o *PartnersResponsePartnersInnerAddressAttributes) HasStreetName2() bool`

HasStreetName2 returns a boolean if a field has been set.

### SetStreetName2Nil

`func (o *PartnersResponsePartnersInnerAddressAttributes) SetStreetName2Nil(b bool)`

 SetStreetName2Nil sets the value for StreetName2 to be an explicit nil

### UnsetStreetName2
`func (o *PartnersResponsePartnersInnerAddressAttributes) UnsetStreetName2()`

UnsetStreetName2 ensures that no value is present for StreetName2, not even an explicit nil
### GetZipcode

`func (o *PartnersResponsePartnersInnerAddressAttributes) GetZipcode() string`

GetZipcode returns the Zipcode field if non-nil, zero value otherwise.

### GetZipcodeOk

`func (o *PartnersResponsePartnersInnerAddressAttributes) GetZipcodeOk() (*string, bool)`

GetZipcodeOk returns a tuple with the Zipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipcode

`func (o *PartnersResponsePartnersInnerAddressAttributes) SetZipcode(v string)`

SetZipcode sets Zipcode field to given value.

### HasZipcode

`func (o *PartnersResponsePartnersInnerAddressAttributes) HasZipcode() bool`

HasZipcode returns a boolean if a field has been set.

### SetZipcodeNil

`func (o *PartnersResponsePartnersInnerAddressAttributes) SetZipcodeNil(b bool)`

 SetZipcodeNil sets the value for Zipcode to be an explicit nil

### UnsetZipcode
`func (o *PartnersResponsePartnersInnerAddressAttributes) UnsetZipcode()`

UnsetZipcode ensures that no value is present for Zipcode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


