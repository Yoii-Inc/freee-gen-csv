# QuotationIndexResponseQuotationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyAddress1** | Pointer to **NullableString** | 市区町村・番地 | [optional] 
**CompanyAddress2** | Pointer to **NullableString** | 建物名・部屋番号など | [optional] 
**CompanyContactInfo** | Pointer to **NullableString** | 事業所担当者名 | [optional] 
**CompanyId** | **int64** | 事業所ID | 
**CompanyName** | **string** | 事業所名 | 
**CompanyPrefectureCode** | Pointer to **NullableInt64** | 都道府県コード（-1: 設定しない、0:北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄 | [optional] 
**CompanyPrefectureName** | Pointer to **NullableString** | 都道府県 | [optional] 
**CompanyZipcode** | Pointer to **NullableString** | 郵便番号 | [optional] 
**Description** | Pointer to **NullableString** | 概要 | [optional] 
**Id** | **int64** | 見積書ID | 
**IssueDate** | **string** | 見積日 (yyyy-mm-dd) | 
**MailSentAt** | Pointer to **NullableString** | メール送信日時(最新) | [optional] 
**Message** | Pointer to **NullableString** | メッセージ | [optional] 
**Notes** | Pointer to **NullableString** | 備考 | [optional] 
**PartnerAddress1** | Pointer to **NullableString** | 市区町村・番地 | [optional] 
**PartnerAddress2** | Pointer to **NullableString** | 建物名・部屋番号など | [optional] 
**PartnerCode** | Pointer to **NullableString** | 取引先コード | [optional] 
**PartnerContactInfo** | Pointer to **NullableString** | 取引先担当者名 | [optional] 
**PartnerDisplayName** | Pointer to **NullableString** | 見積書に表示する取引先名 | [optional] 
**PartnerId** | **NullableInt64** | 取引先ID | 
**PartnerName** | Pointer to **NullableString** | 取引先名 | [optional] 
**PartnerPrefectureCode** | Pointer to **NullableInt64** | 都道府県コード（-1: 設定しない、0:北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄 | [optional] 
**PartnerPrefectureName** | Pointer to **NullableString** | 都道府県 | [optional] 
**PartnerTitle** | **NullableString** | 敬称（御中、様、(空白)の3つから選択） | 
**PartnerZipcode** | Pointer to **NullableString** | 郵便番号 | [optional] 
**QuotationContents** | Pointer to [**[]QuotationIndexResponseQuotationsInnerQuotationContentsInner**](QuotationIndexResponseQuotationsInnerQuotationContentsInner.md) | 見積内容 | [optional] 
**QuotationLayout** | **string** | 見積書レイアウト * &#x60;default_classic&#x60; - レイアウト１/クラシック (デフォルト)  * &#x60;standard_classic&#x60; - レイアウト２/クラシック  * &#x60;envelope_classic&#x60; - 封筒１/クラシック  * &#x60;default_modern&#x60; - レイアウト１/モダン  * &#x60;standard_modern&#x60; - レイアウト２/モダン  * &#x60;envelope_modern&#x60; - 封筒/モダン | 
**QuotationNumber** | **string** | 見積書番号 | 
**QuotationStatus** | **string** | 見積書ステータス  (unsubmitted: 送付待ち, submitted: 送付済み, all: 全て) | 
**SubTotal** | Pointer to **int64** | 小計 | [optional] 
**TaxEntryMethod** | **string** | 見積書の消費税計算方法(inclusive: 内税, exclusive: 外税) | 
**Title** | Pointer to **NullableString** | タイトル | [optional] 
**TotalAmount** | **int64** | 合計金額 | 
**TotalAmountPerVatRate** | [**InvoiceIndexResponseInvoicesInnerTotalAmountPerVatRate**](InvoiceIndexResponseInvoicesInnerTotalAmountPerVatRate.md) |  | 
**TotalVat** | Pointer to **int64** | 消費税 | [optional] 
**WebConfirmedAt** | Pointer to **NullableString** | Web共有取引先確認日時(最新) | [optional] 
**WebDownloadedAt** | Pointer to **NullableString** | Web共有ダウンロード日時(最新) | [optional] 
**WebPublishedAt** | Pointer to **NullableString** | Web共有日時(最新) | [optional] 

## Methods

### NewQuotationIndexResponseQuotationsInner

`func NewQuotationIndexResponseQuotationsInner(companyId int64, companyName string, id int64, issueDate string, partnerId NullableInt64, partnerTitle NullableString, quotationLayout string, quotationNumber string, quotationStatus string, taxEntryMethod string, totalAmount int64, totalAmountPerVatRate InvoiceIndexResponseInvoicesInnerTotalAmountPerVatRate, ) *QuotationIndexResponseQuotationsInner`

NewQuotationIndexResponseQuotationsInner instantiates a new QuotationIndexResponseQuotationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotationIndexResponseQuotationsInnerWithDefaults

`func NewQuotationIndexResponseQuotationsInnerWithDefaults() *QuotationIndexResponseQuotationsInner`

NewQuotationIndexResponseQuotationsInnerWithDefaults instantiates a new QuotationIndexResponseQuotationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyAddress1

`func (o *QuotationIndexResponseQuotationsInner) GetCompanyAddress1() string`

GetCompanyAddress1 returns the CompanyAddress1 field if non-nil, zero value otherwise.

### GetCompanyAddress1Ok

`func (o *QuotationIndexResponseQuotationsInner) GetCompanyAddress1Ok() (*string, bool)`

GetCompanyAddress1Ok returns a tuple with the CompanyAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAddress1

`func (o *QuotationIndexResponseQuotationsInner) SetCompanyAddress1(v string)`

SetCompanyAddress1 sets CompanyAddress1 field to given value.

### HasCompanyAddress1

`func (o *QuotationIndexResponseQuotationsInner) HasCompanyAddress1() bool`

HasCompanyAddress1 returns a boolean if a field has been set.

### SetCompanyAddress1Nil

`func (o *QuotationIndexResponseQuotationsInner) SetCompanyAddress1Nil(b bool)`

 SetCompanyAddress1Nil sets the value for CompanyAddress1 to be an explicit nil

### UnsetCompanyAddress1
`func (o *QuotationIndexResponseQuotationsInner) UnsetCompanyAddress1()`

UnsetCompanyAddress1 ensures that no value is present for CompanyAddress1, not even an explicit nil
### GetCompanyAddress2

`func (o *QuotationIndexResponseQuotationsInner) GetCompanyAddress2() string`

GetCompanyAddress2 returns the CompanyAddress2 field if non-nil, zero value otherwise.

### GetCompanyAddress2Ok

`func (o *QuotationIndexResponseQuotationsInner) GetCompanyAddress2Ok() (*string, bool)`

GetCompanyAddress2Ok returns a tuple with the CompanyAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAddress2

`func (o *QuotationIndexResponseQuotationsInner) SetCompanyAddress2(v string)`

SetCompanyAddress2 sets CompanyAddress2 field to given value.

### HasCompanyAddress2

`func (o *QuotationIndexResponseQuotationsInner) HasCompanyAddress2() bool`

HasCompanyAddress2 returns a boolean if a field has been set.

### SetCompanyAddress2Nil

`func (o *QuotationIndexResponseQuotationsInner) SetCompanyAddress2Nil(b bool)`

 SetCompanyAddress2Nil sets the value for CompanyAddress2 to be an explicit nil

### UnsetCompanyAddress2
`func (o *QuotationIndexResponseQuotationsInner) UnsetCompanyAddress2()`

UnsetCompanyAddress2 ensures that no value is present for CompanyAddress2, not even an explicit nil
### GetCompanyContactInfo

`func (o *QuotationIndexResponseQuotationsInner) GetCompanyContactInfo() string`

GetCompanyContactInfo returns the CompanyContactInfo field if non-nil, zero value otherwise.

### GetCompanyContactInfoOk

`func (o *QuotationIndexResponseQuotationsInner) GetCompanyContactInfoOk() (*string, bool)`

GetCompanyContactInfoOk returns a tuple with the CompanyContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyContactInfo

`func (o *QuotationIndexResponseQuotationsInner) SetCompanyContactInfo(v string)`

SetCompanyContactInfo sets CompanyContactInfo field to given value.

### HasCompanyContactInfo

`func (o *QuotationIndexResponseQuotationsInner) HasCompanyContactInfo() bool`

HasCompanyContactInfo returns a boolean if a field has been set.

### SetCompanyContactInfoNil

`func (o *QuotationIndexResponseQuotationsInner) SetCompanyContactInfoNil(b bool)`

 SetCompanyContactInfoNil sets the value for CompanyContactInfo to be an explicit nil

### UnsetCompanyContactInfo
`func (o *QuotationIndexResponseQuotationsInner) UnsetCompanyContactInfo()`

UnsetCompanyContactInfo ensures that no value is present for CompanyContactInfo, not even an explicit nil
### GetCompanyId

`func (o *QuotationIndexResponseQuotationsInner) GetCompanyId() int64`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *QuotationIndexResponseQuotationsInner) GetCompanyIdOk() (*int64, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *QuotationIndexResponseQuotationsInner) SetCompanyId(v int64)`

SetCompanyId sets CompanyId field to given value.


### GetCompanyName

`func (o *QuotationIndexResponseQuotationsInner) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *QuotationIndexResponseQuotationsInner) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *QuotationIndexResponseQuotationsInner) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### GetCompanyPrefectureCode

`func (o *QuotationIndexResponseQuotationsInner) GetCompanyPrefectureCode() int64`

GetCompanyPrefectureCode returns the CompanyPrefectureCode field if non-nil, zero value otherwise.

### GetCompanyPrefectureCodeOk

`func (o *QuotationIndexResponseQuotationsInner) GetCompanyPrefectureCodeOk() (*int64, bool)`

GetCompanyPrefectureCodeOk returns a tuple with the CompanyPrefectureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyPrefectureCode

`func (o *QuotationIndexResponseQuotationsInner) SetCompanyPrefectureCode(v int64)`

SetCompanyPrefectureCode sets CompanyPrefectureCode field to given value.

### HasCompanyPrefectureCode

`func (o *QuotationIndexResponseQuotationsInner) HasCompanyPrefectureCode() bool`

HasCompanyPrefectureCode returns a boolean if a field has been set.

### SetCompanyPrefectureCodeNil

`func (o *QuotationIndexResponseQuotationsInner) SetCompanyPrefectureCodeNil(b bool)`

 SetCompanyPrefectureCodeNil sets the value for CompanyPrefectureCode to be an explicit nil

### UnsetCompanyPrefectureCode
`func (o *QuotationIndexResponseQuotationsInner) UnsetCompanyPrefectureCode()`

UnsetCompanyPrefectureCode ensures that no value is present for CompanyPrefectureCode, not even an explicit nil
### GetCompanyPrefectureName

`func (o *QuotationIndexResponseQuotationsInner) GetCompanyPrefectureName() string`

GetCompanyPrefectureName returns the CompanyPrefectureName field if non-nil, zero value otherwise.

### GetCompanyPrefectureNameOk

`func (o *QuotationIndexResponseQuotationsInner) GetCompanyPrefectureNameOk() (*string, bool)`

GetCompanyPrefectureNameOk returns a tuple with the CompanyPrefectureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyPrefectureName

`func (o *QuotationIndexResponseQuotationsInner) SetCompanyPrefectureName(v string)`

SetCompanyPrefectureName sets CompanyPrefectureName field to given value.

### HasCompanyPrefectureName

`func (o *QuotationIndexResponseQuotationsInner) HasCompanyPrefectureName() bool`

HasCompanyPrefectureName returns a boolean if a field has been set.

### SetCompanyPrefectureNameNil

`func (o *QuotationIndexResponseQuotationsInner) SetCompanyPrefectureNameNil(b bool)`

 SetCompanyPrefectureNameNil sets the value for CompanyPrefectureName to be an explicit nil

### UnsetCompanyPrefectureName
`func (o *QuotationIndexResponseQuotationsInner) UnsetCompanyPrefectureName()`

UnsetCompanyPrefectureName ensures that no value is present for CompanyPrefectureName, not even an explicit nil
### GetCompanyZipcode

`func (o *QuotationIndexResponseQuotationsInner) GetCompanyZipcode() string`

GetCompanyZipcode returns the CompanyZipcode field if non-nil, zero value otherwise.

### GetCompanyZipcodeOk

`func (o *QuotationIndexResponseQuotationsInner) GetCompanyZipcodeOk() (*string, bool)`

GetCompanyZipcodeOk returns a tuple with the CompanyZipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyZipcode

`func (o *QuotationIndexResponseQuotationsInner) SetCompanyZipcode(v string)`

SetCompanyZipcode sets CompanyZipcode field to given value.

### HasCompanyZipcode

`func (o *QuotationIndexResponseQuotationsInner) HasCompanyZipcode() bool`

HasCompanyZipcode returns a boolean if a field has been set.

### SetCompanyZipcodeNil

`func (o *QuotationIndexResponseQuotationsInner) SetCompanyZipcodeNil(b bool)`

 SetCompanyZipcodeNil sets the value for CompanyZipcode to be an explicit nil

### UnsetCompanyZipcode
`func (o *QuotationIndexResponseQuotationsInner) UnsetCompanyZipcode()`

UnsetCompanyZipcode ensures that no value is present for CompanyZipcode, not even an explicit nil
### GetDescription

`func (o *QuotationIndexResponseQuotationsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuotationIndexResponseQuotationsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuotationIndexResponseQuotationsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuotationIndexResponseQuotationsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *QuotationIndexResponseQuotationsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *QuotationIndexResponseQuotationsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *QuotationIndexResponseQuotationsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuotationIndexResponseQuotationsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuotationIndexResponseQuotationsInner) SetId(v int64)`

SetId sets Id field to given value.


### GetIssueDate

`func (o *QuotationIndexResponseQuotationsInner) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *QuotationIndexResponseQuotationsInner) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *QuotationIndexResponseQuotationsInner) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetMailSentAt

`func (o *QuotationIndexResponseQuotationsInner) GetMailSentAt() string`

GetMailSentAt returns the MailSentAt field if non-nil, zero value otherwise.

### GetMailSentAtOk

`func (o *QuotationIndexResponseQuotationsInner) GetMailSentAtOk() (*string, bool)`

GetMailSentAtOk returns a tuple with the MailSentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailSentAt

`func (o *QuotationIndexResponseQuotationsInner) SetMailSentAt(v string)`

SetMailSentAt sets MailSentAt field to given value.

### HasMailSentAt

`func (o *QuotationIndexResponseQuotationsInner) HasMailSentAt() bool`

HasMailSentAt returns a boolean if a field has been set.

### SetMailSentAtNil

`func (o *QuotationIndexResponseQuotationsInner) SetMailSentAtNil(b bool)`

 SetMailSentAtNil sets the value for MailSentAt to be an explicit nil

### UnsetMailSentAt
`func (o *QuotationIndexResponseQuotationsInner) UnsetMailSentAt()`

UnsetMailSentAt ensures that no value is present for MailSentAt, not even an explicit nil
### GetMessage

`func (o *QuotationIndexResponseQuotationsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *QuotationIndexResponseQuotationsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *QuotationIndexResponseQuotationsInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *QuotationIndexResponseQuotationsInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *QuotationIndexResponseQuotationsInner) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *QuotationIndexResponseQuotationsInner) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetNotes

`func (o *QuotationIndexResponseQuotationsInner) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *QuotationIndexResponseQuotationsInner) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *QuotationIndexResponseQuotationsInner) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *QuotationIndexResponseQuotationsInner) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *QuotationIndexResponseQuotationsInner) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *QuotationIndexResponseQuotationsInner) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetPartnerAddress1

`func (o *QuotationIndexResponseQuotationsInner) GetPartnerAddress1() string`

GetPartnerAddress1 returns the PartnerAddress1 field if non-nil, zero value otherwise.

### GetPartnerAddress1Ok

`func (o *QuotationIndexResponseQuotationsInner) GetPartnerAddress1Ok() (*string, bool)`

GetPartnerAddress1Ok returns a tuple with the PartnerAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerAddress1

`func (o *QuotationIndexResponseQuotationsInner) SetPartnerAddress1(v string)`

SetPartnerAddress1 sets PartnerAddress1 field to given value.

### HasPartnerAddress1

`func (o *QuotationIndexResponseQuotationsInner) HasPartnerAddress1() bool`

HasPartnerAddress1 returns a boolean if a field has been set.

### SetPartnerAddress1Nil

`func (o *QuotationIndexResponseQuotationsInner) SetPartnerAddress1Nil(b bool)`

 SetPartnerAddress1Nil sets the value for PartnerAddress1 to be an explicit nil

### UnsetPartnerAddress1
`func (o *QuotationIndexResponseQuotationsInner) UnsetPartnerAddress1()`

UnsetPartnerAddress1 ensures that no value is present for PartnerAddress1, not even an explicit nil
### GetPartnerAddress2

`func (o *QuotationIndexResponseQuotationsInner) GetPartnerAddress2() string`

GetPartnerAddress2 returns the PartnerAddress2 field if non-nil, zero value otherwise.

### GetPartnerAddress2Ok

`func (o *QuotationIndexResponseQuotationsInner) GetPartnerAddress2Ok() (*string, bool)`

GetPartnerAddress2Ok returns a tuple with the PartnerAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerAddress2

`func (o *QuotationIndexResponseQuotationsInner) SetPartnerAddress2(v string)`

SetPartnerAddress2 sets PartnerAddress2 field to given value.

### HasPartnerAddress2

`func (o *QuotationIndexResponseQuotationsInner) HasPartnerAddress2() bool`

HasPartnerAddress2 returns a boolean if a field has been set.

### SetPartnerAddress2Nil

`func (o *QuotationIndexResponseQuotationsInner) SetPartnerAddress2Nil(b bool)`

 SetPartnerAddress2Nil sets the value for PartnerAddress2 to be an explicit nil

### UnsetPartnerAddress2
`func (o *QuotationIndexResponseQuotationsInner) UnsetPartnerAddress2()`

UnsetPartnerAddress2 ensures that no value is present for PartnerAddress2, not even an explicit nil
### GetPartnerCode

`func (o *QuotationIndexResponseQuotationsInner) GetPartnerCode() string`

GetPartnerCode returns the PartnerCode field if non-nil, zero value otherwise.

### GetPartnerCodeOk

`func (o *QuotationIndexResponseQuotationsInner) GetPartnerCodeOk() (*string, bool)`

GetPartnerCodeOk returns a tuple with the PartnerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCode

`func (o *QuotationIndexResponseQuotationsInner) SetPartnerCode(v string)`

SetPartnerCode sets PartnerCode field to given value.

### HasPartnerCode

`func (o *QuotationIndexResponseQuotationsInner) HasPartnerCode() bool`

HasPartnerCode returns a boolean if a field has been set.

### SetPartnerCodeNil

`func (o *QuotationIndexResponseQuotationsInner) SetPartnerCodeNil(b bool)`

 SetPartnerCodeNil sets the value for PartnerCode to be an explicit nil

### UnsetPartnerCode
`func (o *QuotationIndexResponseQuotationsInner) UnsetPartnerCode()`

UnsetPartnerCode ensures that no value is present for PartnerCode, not even an explicit nil
### GetPartnerContactInfo

`func (o *QuotationIndexResponseQuotationsInner) GetPartnerContactInfo() string`

GetPartnerContactInfo returns the PartnerContactInfo field if non-nil, zero value otherwise.

### GetPartnerContactInfoOk

`func (o *QuotationIndexResponseQuotationsInner) GetPartnerContactInfoOk() (*string, bool)`

GetPartnerContactInfoOk returns a tuple with the PartnerContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerContactInfo

`func (o *QuotationIndexResponseQuotationsInner) SetPartnerContactInfo(v string)`

SetPartnerContactInfo sets PartnerContactInfo field to given value.

### HasPartnerContactInfo

`func (o *QuotationIndexResponseQuotationsInner) HasPartnerContactInfo() bool`

HasPartnerContactInfo returns a boolean if a field has been set.

### SetPartnerContactInfoNil

`func (o *QuotationIndexResponseQuotationsInner) SetPartnerContactInfoNil(b bool)`

 SetPartnerContactInfoNil sets the value for PartnerContactInfo to be an explicit nil

### UnsetPartnerContactInfo
`func (o *QuotationIndexResponseQuotationsInner) UnsetPartnerContactInfo()`

UnsetPartnerContactInfo ensures that no value is present for PartnerContactInfo, not even an explicit nil
### GetPartnerDisplayName

`func (o *QuotationIndexResponseQuotationsInner) GetPartnerDisplayName() string`

GetPartnerDisplayName returns the PartnerDisplayName field if non-nil, zero value otherwise.

### GetPartnerDisplayNameOk

`func (o *QuotationIndexResponseQuotationsInner) GetPartnerDisplayNameOk() (*string, bool)`

GetPartnerDisplayNameOk returns a tuple with the PartnerDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerDisplayName

`func (o *QuotationIndexResponseQuotationsInner) SetPartnerDisplayName(v string)`

SetPartnerDisplayName sets PartnerDisplayName field to given value.

### HasPartnerDisplayName

`func (o *QuotationIndexResponseQuotationsInner) HasPartnerDisplayName() bool`

HasPartnerDisplayName returns a boolean if a field has been set.

### SetPartnerDisplayNameNil

`func (o *QuotationIndexResponseQuotationsInner) SetPartnerDisplayNameNil(b bool)`

 SetPartnerDisplayNameNil sets the value for PartnerDisplayName to be an explicit nil

### UnsetPartnerDisplayName
`func (o *QuotationIndexResponseQuotationsInner) UnsetPartnerDisplayName()`

UnsetPartnerDisplayName ensures that no value is present for PartnerDisplayName, not even an explicit nil
### GetPartnerId

`func (o *QuotationIndexResponseQuotationsInner) GetPartnerId() int64`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *QuotationIndexResponseQuotationsInner) GetPartnerIdOk() (*int64, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *QuotationIndexResponseQuotationsInner) SetPartnerId(v int64)`

SetPartnerId sets PartnerId field to given value.


### SetPartnerIdNil

`func (o *QuotationIndexResponseQuotationsInner) SetPartnerIdNil(b bool)`

 SetPartnerIdNil sets the value for PartnerId to be an explicit nil

### UnsetPartnerId
`func (o *QuotationIndexResponseQuotationsInner) UnsetPartnerId()`

UnsetPartnerId ensures that no value is present for PartnerId, not even an explicit nil
### GetPartnerName

`func (o *QuotationIndexResponseQuotationsInner) GetPartnerName() string`

GetPartnerName returns the PartnerName field if non-nil, zero value otherwise.

### GetPartnerNameOk

`func (o *QuotationIndexResponseQuotationsInner) GetPartnerNameOk() (*string, bool)`

GetPartnerNameOk returns a tuple with the PartnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerName

`func (o *QuotationIndexResponseQuotationsInner) SetPartnerName(v string)`

SetPartnerName sets PartnerName field to given value.

### HasPartnerName

`func (o *QuotationIndexResponseQuotationsInner) HasPartnerName() bool`

HasPartnerName returns a boolean if a field has been set.

### SetPartnerNameNil

`func (o *QuotationIndexResponseQuotationsInner) SetPartnerNameNil(b bool)`

 SetPartnerNameNil sets the value for PartnerName to be an explicit nil

### UnsetPartnerName
`func (o *QuotationIndexResponseQuotationsInner) UnsetPartnerName()`

UnsetPartnerName ensures that no value is present for PartnerName, not even an explicit nil
### GetPartnerPrefectureCode

`func (o *QuotationIndexResponseQuotationsInner) GetPartnerPrefectureCode() int64`

GetPartnerPrefectureCode returns the PartnerPrefectureCode field if non-nil, zero value otherwise.

### GetPartnerPrefectureCodeOk

`func (o *QuotationIndexResponseQuotationsInner) GetPartnerPrefectureCodeOk() (*int64, bool)`

GetPartnerPrefectureCodeOk returns a tuple with the PartnerPrefectureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPrefectureCode

`func (o *QuotationIndexResponseQuotationsInner) SetPartnerPrefectureCode(v int64)`

SetPartnerPrefectureCode sets PartnerPrefectureCode field to given value.

### HasPartnerPrefectureCode

`func (o *QuotationIndexResponseQuotationsInner) HasPartnerPrefectureCode() bool`

HasPartnerPrefectureCode returns a boolean if a field has been set.

### SetPartnerPrefectureCodeNil

`func (o *QuotationIndexResponseQuotationsInner) SetPartnerPrefectureCodeNil(b bool)`

 SetPartnerPrefectureCodeNil sets the value for PartnerPrefectureCode to be an explicit nil

### UnsetPartnerPrefectureCode
`func (o *QuotationIndexResponseQuotationsInner) UnsetPartnerPrefectureCode()`

UnsetPartnerPrefectureCode ensures that no value is present for PartnerPrefectureCode, not even an explicit nil
### GetPartnerPrefectureName

`func (o *QuotationIndexResponseQuotationsInner) GetPartnerPrefectureName() string`

GetPartnerPrefectureName returns the PartnerPrefectureName field if non-nil, zero value otherwise.

### GetPartnerPrefectureNameOk

`func (o *QuotationIndexResponseQuotationsInner) GetPartnerPrefectureNameOk() (*string, bool)`

GetPartnerPrefectureNameOk returns a tuple with the PartnerPrefectureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPrefectureName

`func (o *QuotationIndexResponseQuotationsInner) SetPartnerPrefectureName(v string)`

SetPartnerPrefectureName sets PartnerPrefectureName field to given value.

### HasPartnerPrefectureName

`func (o *QuotationIndexResponseQuotationsInner) HasPartnerPrefectureName() bool`

HasPartnerPrefectureName returns a boolean if a field has been set.

### SetPartnerPrefectureNameNil

`func (o *QuotationIndexResponseQuotationsInner) SetPartnerPrefectureNameNil(b bool)`

 SetPartnerPrefectureNameNil sets the value for PartnerPrefectureName to be an explicit nil

### UnsetPartnerPrefectureName
`func (o *QuotationIndexResponseQuotationsInner) UnsetPartnerPrefectureName()`

UnsetPartnerPrefectureName ensures that no value is present for PartnerPrefectureName, not even an explicit nil
### GetPartnerTitle

`func (o *QuotationIndexResponseQuotationsInner) GetPartnerTitle() string`

GetPartnerTitle returns the PartnerTitle field if non-nil, zero value otherwise.

### GetPartnerTitleOk

`func (o *QuotationIndexResponseQuotationsInner) GetPartnerTitleOk() (*string, bool)`

GetPartnerTitleOk returns a tuple with the PartnerTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerTitle

`func (o *QuotationIndexResponseQuotationsInner) SetPartnerTitle(v string)`

SetPartnerTitle sets PartnerTitle field to given value.


### SetPartnerTitleNil

`func (o *QuotationIndexResponseQuotationsInner) SetPartnerTitleNil(b bool)`

 SetPartnerTitleNil sets the value for PartnerTitle to be an explicit nil

### UnsetPartnerTitle
`func (o *QuotationIndexResponseQuotationsInner) UnsetPartnerTitle()`

UnsetPartnerTitle ensures that no value is present for PartnerTitle, not even an explicit nil
### GetPartnerZipcode

`func (o *QuotationIndexResponseQuotationsInner) GetPartnerZipcode() string`

GetPartnerZipcode returns the PartnerZipcode field if non-nil, zero value otherwise.

### GetPartnerZipcodeOk

`func (o *QuotationIndexResponseQuotationsInner) GetPartnerZipcodeOk() (*string, bool)`

GetPartnerZipcodeOk returns a tuple with the PartnerZipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerZipcode

`func (o *QuotationIndexResponseQuotationsInner) SetPartnerZipcode(v string)`

SetPartnerZipcode sets PartnerZipcode field to given value.

### HasPartnerZipcode

`func (o *QuotationIndexResponseQuotationsInner) HasPartnerZipcode() bool`

HasPartnerZipcode returns a boolean if a field has been set.

### SetPartnerZipcodeNil

`func (o *QuotationIndexResponseQuotationsInner) SetPartnerZipcodeNil(b bool)`

 SetPartnerZipcodeNil sets the value for PartnerZipcode to be an explicit nil

### UnsetPartnerZipcode
`func (o *QuotationIndexResponseQuotationsInner) UnsetPartnerZipcode()`

UnsetPartnerZipcode ensures that no value is present for PartnerZipcode, not even an explicit nil
### GetQuotationContents

`func (o *QuotationIndexResponseQuotationsInner) GetQuotationContents() []QuotationIndexResponseQuotationsInnerQuotationContentsInner`

GetQuotationContents returns the QuotationContents field if non-nil, zero value otherwise.

### GetQuotationContentsOk

`func (o *QuotationIndexResponseQuotationsInner) GetQuotationContentsOk() (*[]QuotationIndexResponseQuotationsInnerQuotationContentsInner, bool)`

GetQuotationContentsOk returns a tuple with the QuotationContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotationContents

`func (o *QuotationIndexResponseQuotationsInner) SetQuotationContents(v []QuotationIndexResponseQuotationsInnerQuotationContentsInner)`

SetQuotationContents sets QuotationContents field to given value.

### HasQuotationContents

`func (o *QuotationIndexResponseQuotationsInner) HasQuotationContents() bool`

HasQuotationContents returns a boolean if a field has been set.

### GetQuotationLayout

`func (o *QuotationIndexResponseQuotationsInner) GetQuotationLayout() string`

GetQuotationLayout returns the QuotationLayout field if non-nil, zero value otherwise.

### GetQuotationLayoutOk

`func (o *QuotationIndexResponseQuotationsInner) GetQuotationLayoutOk() (*string, bool)`

GetQuotationLayoutOk returns a tuple with the QuotationLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotationLayout

`func (o *QuotationIndexResponseQuotationsInner) SetQuotationLayout(v string)`

SetQuotationLayout sets QuotationLayout field to given value.


### GetQuotationNumber

`func (o *QuotationIndexResponseQuotationsInner) GetQuotationNumber() string`

GetQuotationNumber returns the QuotationNumber field if non-nil, zero value otherwise.

### GetQuotationNumberOk

`func (o *QuotationIndexResponseQuotationsInner) GetQuotationNumberOk() (*string, bool)`

GetQuotationNumberOk returns a tuple with the QuotationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotationNumber

`func (o *QuotationIndexResponseQuotationsInner) SetQuotationNumber(v string)`

SetQuotationNumber sets QuotationNumber field to given value.


### GetQuotationStatus

`func (o *QuotationIndexResponseQuotationsInner) GetQuotationStatus() string`

GetQuotationStatus returns the QuotationStatus field if non-nil, zero value otherwise.

### GetQuotationStatusOk

`func (o *QuotationIndexResponseQuotationsInner) GetQuotationStatusOk() (*string, bool)`

GetQuotationStatusOk returns a tuple with the QuotationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotationStatus

`func (o *QuotationIndexResponseQuotationsInner) SetQuotationStatus(v string)`

SetQuotationStatus sets QuotationStatus field to given value.


### GetSubTotal

`func (o *QuotationIndexResponseQuotationsInner) GetSubTotal() int64`

GetSubTotal returns the SubTotal field if non-nil, zero value otherwise.

### GetSubTotalOk

`func (o *QuotationIndexResponseQuotationsInner) GetSubTotalOk() (*int64, bool)`

GetSubTotalOk returns a tuple with the SubTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotal

`func (o *QuotationIndexResponseQuotationsInner) SetSubTotal(v int64)`

SetSubTotal sets SubTotal field to given value.

### HasSubTotal

`func (o *QuotationIndexResponseQuotationsInner) HasSubTotal() bool`

HasSubTotal returns a boolean if a field has been set.

### GetTaxEntryMethod

`func (o *QuotationIndexResponseQuotationsInner) GetTaxEntryMethod() string`

GetTaxEntryMethod returns the TaxEntryMethod field if non-nil, zero value otherwise.

### GetTaxEntryMethodOk

`func (o *QuotationIndexResponseQuotationsInner) GetTaxEntryMethodOk() (*string, bool)`

GetTaxEntryMethodOk returns a tuple with the TaxEntryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxEntryMethod

`func (o *QuotationIndexResponseQuotationsInner) SetTaxEntryMethod(v string)`

SetTaxEntryMethod sets TaxEntryMethod field to given value.


### GetTitle

`func (o *QuotationIndexResponseQuotationsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *QuotationIndexResponseQuotationsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *QuotationIndexResponseQuotationsInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *QuotationIndexResponseQuotationsInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *QuotationIndexResponseQuotationsInner) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *QuotationIndexResponseQuotationsInner) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetTotalAmount

`func (o *QuotationIndexResponseQuotationsInner) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *QuotationIndexResponseQuotationsInner) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *QuotationIndexResponseQuotationsInner) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.


### GetTotalAmountPerVatRate

`func (o *QuotationIndexResponseQuotationsInner) GetTotalAmountPerVatRate() InvoiceIndexResponseInvoicesInnerTotalAmountPerVatRate`

GetTotalAmountPerVatRate returns the TotalAmountPerVatRate field if non-nil, zero value otherwise.

### GetTotalAmountPerVatRateOk

`func (o *QuotationIndexResponseQuotationsInner) GetTotalAmountPerVatRateOk() (*InvoiceIndexResponseInvoicesInnerTotalAmountPerVatRate, bool)`

GetTotalAmountPerVatRateOk returns a tuple with the TotalAmountPerVatRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountPerVatRate

`func (o *QuotationIndexResponseQuotationsInner) SetTotalAmountPerVatRate(v InvoiceIndexResponseInvoicesInnerTotalAmountPerVatRate)`

SetTotalAmountPerVatRate sets TotalAmountPerVatRate field to given value.


### GetTotalVat

`func (o *QuotationIndexResponseQuotationsInner) GetTotalVat() int64`

GetTotalVat returns the TotalVat field if non-nil, zero value otherwise.

### GetTotalVatOk

`func (o *QuotationIndexResponseQuotationsInner) GetTotalVatOk() (*int64, bool)`

GetTotalVatOk returns a tuple with the TotalVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVat

`func (o *QuotationIndexResponseQuotationsInner) SetTotalVat(v int64)`

SetTotalVat sets TotalVat field to given value.

### HasTotalVat

`func (o *QuotationIndexResponseQuotationsInner) HasTotalVat() bool`

HasTotalVat returns a boolean if a field has been set.

### GetWebConfirmedAt

`func (o *QuotationIndexResponseQuotationsInner) GetWebConfirmedAt() string`

GetWebConfirmedAt returns the WebConfirmedAt field if non-nil, zero value otherwise.

### GetWebConfirmedAtOk

`func (o *QuotationIndexResponseQuotationsInner) GetWebConfirmedAtOk() (*string, bool)`

GetWebConfirmedAtOk returns a tuple with the WebConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebConfirmedAt

`func (o *QuotationIndexResponseQuotationsInner) SetWebConfirmedAt(v string)`

SetWebConfirmedAt sets WebConfirmedAt field to given value.

### HasWebConfirmedAt

`func (o *QuotationIndexResponseQuotationsInner) HasWebConfirmedAt() bool`

HasWebConfirmedAt returns a boolean if a field has been set.

### SetWebConfirmedAtNil

`func (o *QuotationIndexResponseQuotationsInner) SetWebConfirmedAtNil(b bool)`

 SetWebConfirmedAtNil sets the value for WebConfirmedAt to be an explicit nil

### UnsetWebConfirmedAt
`func (o *QuotationIndexResponseQuotationsInner) UnsetWebConfirmedAt()`

UnsetWebConfirmedAt ensures that no value is present for WebConfirmedAt, not even an explicit nil
### GetWebDownloadedAt

`func (o *QuotationIndexResponseQuotationsInner) GetWebDownloadedAt() string`

GetWebDownloadedAt returns the WebDownloadedAt field if non-nil, zero value otherwise.

### GetWebDownloadedAtOk

`func (o *QuotationIndexResponseQuotationsInner) GetWebDownloadedAtOk() (*string, bool)`

GetWebDownloadedAtOk returns a tuple with the WebDownloadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebDownloadedAt

`func (o *QuotationIndexResponseQuotationsInner) SetWebDownloadedAt(v string)`

SetWebDownloadedAt sets WebDownloadedAt field to given value.

### HasWebDownloadedAt

`func (o *QuotationIndexResponseQuotationsInner) HasWebDownloadedAt() bool`

HasWebDownloadedAt returns a boolean if a field has been set.

### SetWebDownloadedAtNil

`func (o *QuotationIndexResponseQuotationsInner) SetWebDownloadedAtNil(b bool)`

 SetWebDownloadedAtNil sets the value for WebDownloadedAt to be an explicit nil

### UnsetWebDownloadedAt
`func (o *QuotationIndexResponseQuotationsInner) UnsetWebDownloadedAt()`

UnsetWebDownloadedAt ensures that no value is present for WebDownloadedAt, not even an explicit nil
### GetWebPublishedAt

`func (o *QuotationIndexResponseQuotationsInner) GetWebPublishedAt() string`

GetWebPublishedAt returns the WebPublishedAt field if non-nil, zero value otherwise.

### GetWebPublishedAtOk

`func (o *QuotationIndexResponseQuotationsInner) GetWebPublishedAtOk() (*string, bool)`

GetWebPublishedAtOk returns a tuple with the WebPublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPublishedAt

`func (o *QuotationIndexResponseQuotationsInner) SetWebPublishedAt(v string)`

SetWebPublishedAt sets WebPublishedAt field to given value.

### HasWebPublishedAt

`func (o *QuotationIndexResponseQuotationsInner) HasWebPublishedAt() bool`

HasWebPublishedAt returns a boolean if a field has been set.

### SetWebPublishedAtNil

`func (o *QuotationIndexResponseQuotationsInner) SetWebPublishedAtNil(b bool)`

 SetWebPublishedAtNil sets the value for WebPublishedAt to be an explicit nil

### UnsetWebPublishedAt
`func (o *QuotationIndexResponseQuotationsInner) UnsetWebPublishedAt()`

UnsetWebPublishedAt ensures that no value is present for WebPublishedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


