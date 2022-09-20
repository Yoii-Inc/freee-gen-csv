# InvoiceIndexResponseInvoicesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BookingDate** | Pointer to **NullableString** | 売上計上日 | [optional] 
**CompanyAddress1** | Pointer to **NullableString** | 市区町村・番地 | [optional] 
**CompanyAddress2** | Pointer to **NullableString** | 建物名・部屋番号など | [optional] 
**CompanyContactInfo** | Pointer to **NullableString** | 事業所担当者名 | [optional] 
**CompanyId** | **int64** | 事業所ID | 
**CompanyName** | **string** | 事業所名 | 
**CompanyPrefectureCode** | Pointer to **NullableInt64** | 都道府県コード（-1: 設定しない、0:北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄 | [optional] 
**CompanyPrefectureName** | Pointer to **NullableString** | 都道府県 | [optional] 
**CompanyZipcode** | Pointer to **NullableString** | 郵便番号 | [optional] 
**DealId** | Pointer to **NullableInt64** | 取引ID (invoice_statusがsubmitted, unsubmittedの時IDが表示されます) | [optional] 
**Description** | Pointer to **NullableString** | 概要 | [optional] 
**DueDate** | Pointer to **NullableString** | 期日 (yyyy-mm-dd) | [optional] 
**Id** | **int64** | 請求書ID | 
**InvoiceContents** | Pointer to [**[]InvoiceIndexResponseInvoicesInnerInvoiceContentsInner**](InvoiceIndexResponseInvoicesInnerInvoiceContentsInner.md) | 請求内容 | [optional] 
**InvoiceLayout** | **string** | 請求書レイアウト * &#x60;default_classic&#x60; - レイアウト１/クラシック (デフォルト)  * &#x60;standard_classic&#x60; - レイアウト２/クラシック  * &#x60;envelope_classic&#x60; - 封筒１/クラシック  * &#x60;carried_forward_standard_classic&#x60; - レイアウト３（繰越金額欄あり）/クラシック  * &#x60;carried_forward_envelope_classic&#x60; - 封筒２（繰越金額欄あり）/クラシック  * &#x60;default_modern&#x60; - レイアウト１/モダン  * &#x60;standard_modern&#x60; - レイアウト２/モダン  * &#x60;envelope_modern&#x60; - 封筒/モダン | 
**InvoiceNumber** | **string** | 請求書番号 | 
**InvoiceStatus** | **string** | 請求書ステータス  (draft: 下書き, applying: 申請中, remanded: 差し戻し, rejected: 却下, approved: 承認済み, submitted: 送付済み, unsubmitted: 請求書の承認フローが無効の場合のみ、unsubmitted（送付待ち）の値をとります) | 
**IssueDate** | **string** | 請求日 (yyyy-mm-dd) | 
**MailSentAt** | Pointer to **NullableString** | メール送信日時(最新) | [optional] 
**Message** | Pointer to **NullableString** | メッセージ | [optional] 
**Notes** | Pointer to **NullableString** | 備考 | [optional] 
**PartnerAddress1** | Pointer to **NullableString** | 市区町村・番地 | [optional] 
**PartnerAddress2** | Pointer to **NullableString** | 建物名・部屋番号など | [optional] 
**PartnerCode** | Pointer to **NullableString** | 取引先コード | [optional] 
**PartnerContactInfo** | Pointer to **NullableString** | 取引先担当者名 | [optional] 
**PartnerDisplayName** | Pointer to **NullableString** | 請求書に表示する取引先名 | [optional] 
**PartnerId** | **NullableInt64** | 取引先ID | 
**PartnerName** | Pointer to **NullableString** | 取引先名 | [optional] 
**PartnerPrefectureCode** | Pointer to **NullableInt64** | 都道府県コード（-1: 設定しない、0:北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄 | [optional] 
**PartnerPrefectureName** | Pointer to **NullableString** | 都道府県 | [optional] 
**PartnerTitle** | Pointer to **NullableString** | 敬称（御中、様、(空白)の3つから選択） | [optional] 
**PartnerZipcode** | Pointer to **NullableString** | 郵便番号 | [optional] 
**PaymentBankInfo** | Pointer to **NullableString** | 支払口座 | [optional] 
**PaymentDate** | Pointer to **NullableString** | 入金日 | [optional] 
**PaymentStatus** | Pointer to **string** | 入金ステータス  (unsettled: 入金待ち, settled: 入金済み) | [optional] 
**PaymentType** | **string** | 支払方法 (振込: transfer, 引き落とし: direct_debit) | 
**PostingStatus** | **string** | 郵送ステータス(unrequested: リクエスト前, preview_registered: プレビュー登録, preview_failed: プレビュー登録失敗, ordered: 注文中, order_failed: 注文失敗, printing: 印刷中, canceled: キャンセル, posted: 投函済み) | 
**SubTotal** | Pointer to **int64** | 小計 | [optional] 
**TaxEntryMethod** | **string** | 請求書の消費税計算方法(inclusive: 内税, exclusive: 外税) | 
**Title** | Pointer to **NullableString** | タイトル | [optional] 
**TotalAmount** | **int64** | 合計金額 | 
**TotalAmountPerVatRate** | [**InvoiceIndexResponseInvoicesInnerTotalAmountPerVatRate**](InvoiceIndexResponseInvoicesInnerTotalAmountPerVatRate.md) |  | 
**TotalVat** | Pointer to **int64** | 合計金額 | [optional] 
**WebConfirmedAt** | Pointer to **NullableString** | Web共有取引先確認日時(最新) | [optional] 
**WebDownloadedAt** | Pointer to **NullableString** | Web共有ダウンロード日時(最新) | [optional] 
**WebPublishedAt** | Pointer to **NullableString** | Web共有日時(最新) | [optional] 

## Methods

### NewInvoiceIndexResponseInvoicesInner

`func NewInvoiceIndexResponseInvoicesInner(companyId int64, companyName string, id int64, invoiceLayout string, invoiceNumber string, invoiceStatus string, issueDate string, partnerId NullableInt64, paymentType string, postingStatus string, taxEntryMethod string, totalAmount int64, totalAmountPerVatRate InvoiceIndexResponseInvoicesInnerTotalAmountPerVatRate, ) *InvoiceIndexResponseInvoicesInner`

NewInvoiceIndexResponseInvoicesInner instantiates a new InvoiceIndexResponseInvoicesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceIndexResponseInvoicesInnerWithDefaults

`func NewInvoiceIndexResponseInvoicesInnerWithDefaults() *InvoiceIndexResponseInvoicesInner`

NewInvoiceIndexResponseInvoicesInnerWithDefaults instantiates a new InvoiceIndexResponseInvoicesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBookingDate

`func (o *InvoiceIndexResponseInvoicesInner) GetBookingDate() string`

GetBookingDate returns the BookingDate field if non-nil, zero value otherwise.

### GetBookingDateOk

`func (o *InvoiceIndexResponseInvoicesInner) GetBookingDateOk() (*string, bool)`

GetBookingDateOk returns a tuple with the BookingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookingDate

`func (o *InvoiceIndexResponseInvoicesInner) SetBookingDate(v string)`

SetBookingDate sets BookingDate field to given value.

### HasBookingDate

`func (o *InvoiceIndexResponseInvoicesInner) HasBookingDate() bool`

HasBookingDate returns a boolean if a field has been set.

### SetBookingDateNil

`func (o *InvoiceIndexResponseInvoicesInner) SetBookingDateNil(b bool)`

 SetBookingDateNil sets the value for BookingDate to be an explicit nil

### UnsetBookingDate
`func (o *InvoiceIndexResponseInvoicesInner) UnsetBookingDate()`

UnsetBookingDate ensures that no value is present for BookingDate, not even an explicit nil
### GetCompanyAddress1

`func (o *InvoiceIndexResponseInvoicesInner) GetCompanyAddress1() string`

GetCompanyAddress1 returns the CompanyAddress1 field if non-nil, zero value otherwise.

### GetCompanyAddress1Ok

`func (o *InvoiceIndexResponseInvoicesInner) GetCompanyAddress1Ok() (*string, bool)`

GetCompanyAddress1Ok returns a tuple with the CompanyAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAddress1

`func (o *InvoiceIndexResponseInvoicesInner) SetCompanyAddress1(v string)`

SetCompanyAddress1 sets CompanyAddress1 field to given value.

### HasCompanyAddress1

`func (o *InvoiceIndexResponseInvoicesInner) HasCompanyAddress1() bool`

HasCompanyAddress1 returns a boolean if a field has been set.

### SetCompanyAddress1Nil

`func (o *InvoiceIndexResponseInvoicesInner) SetCompanyAddress1Nil(b bool)`

 SetCompanyAddress1Nil sets the value for CompanyAddress1 to be an explicit nil

### UnsetCompanyAddress1
`func (o *InvoiceIndexResponseInvoicesInner) UnsetCompanyAddress1()`

UnsetCompanyAddress1 ensures that no value is present for CompanyAddress1, not even an explicit nil
### GetCompanyAddress2

`func (o *InvoiceIndexResponseInvoicesInner) GetCompanyAddress2() string`

GetCompanyAddress2 returns the CompanyAddress2 field if non-nil, zero value otherwise.

### GetCompanyAddress2Ok

`func (o *InvoiceIndexResponseInvoicesInner) GetCompanyAddress2Ok() (*string, bool)`

GetCompanyAddress2Ok returns a tuple with the CompanyAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAddress2

`func (o *InvoiceIndexResponseInvoicesInner) SetCompanyAddress2(v string)`

SetCompanyAddress2 sets CompanyAddress2 field to given value.

### HasCompanyAddress2

`func (o *InvoiceIndexResponseInvoicesInner) HasCompanyAddress2() bool`

HasCompanyAddress2 returns a boolean if a field has been set.

### SetCompanyAddress2Nil

`func (o *InvoiceIndexResponseInvoicesInner) SetCompanyAddress2Nil(b bool)`

 SetCompanyAddress2Nil sets the value for CompanyAddress2 to be an explicit nil

### UnsetCompanyAddress2
`func (o *InvoiceIndexResponseInvoicesInner) UnsetCompanyAddress2()`

UnsetCompanyAddress2 ensures that no value is present for CompanyAddress2, not even an explicit nil
### GetCompanyContactInfo

`func (o *InvoiceIndexResponseInvoicesInner) GetCompanyContactInfo() string`

GetCompanyContactInfo returns the CompanyContactInfo field if non-nil, zero value otherwise.

### GetCompanyContactInfoOk

`func (o *InvoiceIndexResponseInvoicesInner) GetCompanyContactInfoOk() (*string, bool)`

GetCompanyContactInfoOk returns a tuple with the CompanyContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyContactInfo

`func (o *InvoiceIndexResponseInvoicesInner) SetCompanyContactInfo(v string)`

SetCompanyContactInfo sets CompanyContactInfo field to given value.

### HasCompanyContactInfo

`func (o *InvoiceIndexResponseInvoicesInner) HasCompanyContactInfo() bool`

HasCompanyContactInfo returns a boolean if a field has been set.

### SetCompanyContactInfoNil

`func (o *InvoiceIndexResponseInvoicesInner) SetCompanyContactInfoNil(b bool)`

 SetCompanyContactInfoNil sets the value for CompanyContactInfo to be an explicit nil

### UnsetCompanyContactInfo
`func (o *InvoiceIndexResponseInvoicesInner) UnsetCompanyContactInfo()`

UnsetCompanyContactInfo ensures that no value is present for CompanyContactInfo, not even an explicit nil
### GetCompanyId

`func (o *InvoiceIndexResponseInvoicesInner) GetCompanyId() int64`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *InvoiceIndexResponseInvoicesInner) GetCompanyIdOk() (*int64, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *InvoiceIndexResponseInvoicesInner) SetCompanyId(v int64)`

SetCompanyId sets CompanyId field to given value.


### GetCompanyName

`func (o *InvoiceIndexResponseInvoicesInner) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *InvoiceIndexResponseInvoicesInner) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *InvoiceIndexResponseInvoicesInner) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### GetCompanyPrefectureCode

`func (o *InvoiceIndexResponseInvoicesInner) GetCompanyPrefectureCode() int64`

GetCompanyPrefectureCode returns the CompanyPrefectureCode field if non-nil, zero value otherwise.

### GetCompanyPrefectureCodeOk

`func (o *InvoiceIndexResponseInvoicesInner) GetCompanyPrefectureCodeOk() (*int64, bool)`

GetCompanyPrefectureCodeOk returns a tuple with the CompanyPrefectureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyPrefectureCode

`func (o *InvoiceIndexResponseInvoicesInner) SetCompanyPrefectureCode(v int64)`

SetCompanyPrefectureCode sets CompanyPrefectureCode field to given value.

### HasCompanyPrefectureCode

`func (o *InvoiceIndexResponseInvoicesInner) HasCompanyPrefectureCode() bool`

HasCompanyPrefectureCode returns a boolean if a field has been set.

### SetCompanyPrefectureCodeNil

`func (o *InvoiceIndexResponseInvoicesInner) SetCompanyPrefectureCodeNil(b bool)`

 SetCompanyPrefectureCodeNil sets the value for CompanyPrefectureCode to be an explicit nil

### UnsetCompanyPrefectureCode
`func (o *InvoiceIndexResponseInvoicesInner) UnsetCompanyPrefectureCode()`

UnsetCompanyPrefectureCode ensures that no value is present for CompanyPrefectureCode, not even an explicit nil
### GetCompanyPrefectureName

`func (o *InvoiceIndexResponseInvoicesInner) GetCompanyPrefectureName() string`

GetCompanyPrefectureName returns the CompanyPrefectureName field if non-nil, zero value otherwise.

### GetCompanyPrefectureNameOk

`func (o *InvoiceIndexResponseInvoicesInner) GetCompanyPrefectureNameOk() (*string, bool)`

GetCompanyPrefectureNameOk returns a tuple with the CompanyPrefectureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyPrefectureName

`func (o *InvoiceIndexResponseInvoicesInner) SetCompanyPrefectureName(v string)`

SetCompanyPrefectureName sets CompanyPrefectureName field to given value.

### HasCompanyPrefectureName

`func (o *InvoiceIndexResponseInvoicesInner) HasCompanyPrefectureName() bool`

HasCompanyPrefectureName returns a boolean if a field has been set.

### SetCompanyPrefectureNameNil

`func (o *InvoiceIndexResponseInvoicesInner) SetCompanyPrefectureNameNil(b bool)`

 SetCompanyPrefectureNameNil sets the value for CompanyPrefectureName to be an explicit nil

### UnsetCompanyPrefectureName
`func (o *InvoiceIndexResponseInvoicesInner) UnsetCompanyPrefectureName()`

UnsetCompanyPrefectureName ensures that no value is present for CompanyPrefectureName, not even an explicit nil
### GetCompanyZipcode

`func (o *InvoiceIndexResponseInvoicesInner) GetCompanyZipcode() string`

GetCompanyZipcode returns the CompanyZipcode field if non-nil, zero value otherwise.

### GetCompanyZipcodeOk

`func (o *InvoiceIndexResponseInvoicesInner) GetCompanyZipcodeOk() (*string, bool)`

GetCompanyZipcodeOk returns a tuple with the CompanyZipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyZipcode

`func (o *InvoiceIndexResponseInvoicesInner) SetCompanyZipcode(v string)`

SetCompanyZipcode sets CompanyZipcode field to given value.

### HasCompanyZipcode

`func (o *InvoiceIndexResponseInvoicesInner) HasCompanyZipcode() bool`

HasCompanyZipcode returns a boolean if a field has been set.

### SetCompanyZipcodeNil

`func (o *InvoiceIndexResponseInvoicesInner) SetCompanyZipcodeNil(b bool)`

 SetCompanyZipcodeNil sets the value for CompanyZipcode to be an explicit nil

### UnsetCompanyZipcode
`func (o *InvoiceIndexResponseInvoicesInner) UnsetCompanyZipcode()`

UnsetCompanyZipcode ensures that no value is present for CompanyZipcode, not even an explicit nil
### GetDealId

`func (o *InvoiceIndexResponseInvoicesInner) GetDealId() int64`

GetDealId returns the DealId field if non-nil, zero value otherwise.

### GetDealIdOk

`func (o *InvoiceIndexResponseInvoicesInner) GetDealIdOk() (*int64, bool)`

GetDealIdOk returns a tuple with the DealId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealId

`func (o *InvoiceIndexResponseInvoicesInner) SetDealId(v int64)`

SetDealId sets DealId field to given value.

### HasDealId

`func (o *InvoiceIndexResponseInvoicesInner) HasDealId() bool`

HasDealId returns a boolean if a field has been set.

### SetDealIdNil

`func (o *InvoiceIndexResponseInvoicesInner) SetDealIdNil(b bool)`

 SetDealIdNil sets the value for DealId to be an explicit nil

### UnsetDealId
`func (o *InvoiceIndexResponseInvoicesInner) UnsetDealId()`

UnsetDealId ensures that no value is present for DealId, not even an explicit nil
### GetDescription

`func (o *InvoiceIndexResponseInvoicesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceIndexResponseInvoicesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceIndexResponseInvoicesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvoiceIndexResponseInvoicesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InvoiceIndexResponseInvoicesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InvoiceIndexResponseInvoicesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDueDate

`func (o *InvoiceIndexResponseInvoicesInner) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InvoiceIndexResponseInvoicesInner) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InvoiceIndexResponseInvoicesInner) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *InvoiceIndexResponseInvoicesInner) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *InvoiceIndexResponseInvoicesInner) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *InvoiceIndexResponseInvoicesInner) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetId

`func (o *InvoiceIndexResponseInvoicesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceIndexResponseInvoicesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceIndexResponseInvoicesInner) SetId(v int64)`

SetId sets Id field to given value.


### GetInvoiceContents

`func (o *InvoiceIndexResponseInvoicesInner) GetInvoiceContents() []InvoiceIndexResponseInvoicesInnerInvoiceContentsInner`

GetInvoiceContents returns the InvoiceContents field if non-nil, zero value otherwise.

### GetInvoiceContentsOk

`func (o *InvoiceIndexResponseInvoicesInner) GetInvoiceContentsOk() (*[]InvoiceIndexResponseInvoicesInnerInvoiceContentsInner, bool)`

GetInvoiceContentsOk returns a tuple with the InvoiceContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceContents

`func (o *InvoiceIndexResponseInvoicesInner) SetInvoiceContents(v []InvoiceIndexResponseInvoicesInnerInvoiceContentsInner)`

SetInvoiceContents sets InvoiceContents field to given value.

### HasInvoiceContents

`func (o *InvoiceIndexResponseInvoicesInner) HasInvoiceContents() bool`

HasInvoiceContents returns a boolean if a field has been set.

### GetInvoiceLayout

`func (o *InvoiceIndexResponseInvoicesInner) GetInvoiceLayout() string`

GetInvoiceLayout returns the InvoiceLayout field if non-nil, zero value otherwise.

### GetInvoiceLayoutOk

`func (o *InvoiceIndexResponseInvoicesInner) GetInvoiceLayoutOk() (*string, bool)`

GetInvoiceLayoutOk returns a tuple with the InvoiceLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceLayout

`func (o *InvoiceIndexResponseInvoicesInner) SetInvoiceLayout(v string)`

SetInvoiceLayout sets InvoiceLayout field to given value.


### GetInvoiceNumber

`func (o *InvoiceIndexResponseInvoicesInner) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *InvoiceIndexResponseInvoicesInner) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *InvoiceIndexResponseInvoicesInner) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.


### GetInvoiceStatus

`func (o *InvoiceIndexResponseInvoicesInner) GetInvoiceStatus() string`

GetInvoiceStatus returns the InvoiceStatus field if non-nil, zero value otherwise.

### GetInvoiceStatusOk

`func (o *InvoiceIndexResponseInvoicesInner) GetInvoiceStatusOk() (*string, bool)`

GetInvoiceStatusOk returns a tuple with the InvoiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceStatus

`func (o *InvoiceIndexResponseInvoicesInner) SetInvoiceStatus(v string)`

SetInvoiceStatus sets InvoiceStatus field to given value.


### GetIssueDate

`func (o *InvoiceIndexResponseInvoicesInner) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *InvoiceIndexResponseInvoicesInner) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *InvoiceIndexResponseInvoicesInner) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetMailSentAt

`func (o *InvoiceIndexResponseInvoicesInner) GetMailSentAt() string`

GetMailSentAt returns the MailSentAt field if non-nil, zero value otherwise.

### GetMailSentAtOk

`func (o *InvoiceIndexResponseInvoicesInner) GetMailSentAtOk() (*string, bool)`

GetMailSentAtOk returns a tuple with the MailSentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailSentAt

`func (o *InvoiceIndexResponseInvoicesInner) SetMailSentAt(v string)`

SetMailSentAt sets MailSentAt field to given value.

### HasMailSentAt

`func (o *InvoiceIndexResponseInvoicesInner) HasMailSentAt() bool`

HasMailSentAt returns a boolean if a field has been set.

### SetMailSentAtNil

`func (o *InvoiceIndexResponseInvoicesInner) SetMailSentAtNil(b bool)`

 SetMailSentAtNil sets the value for MailSentAt to be an explicit nil

### UnsetMailSentAt
`func (o *InvoiceIndexResponseInvoicesInner) UnsetMailSentAt()`

UnsetMailSentAt ensures that no value is present for MailSentAt, not even an explicit nil
### GetMessage

`func (o *InvoiceIndexResponseInvoicesInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InvoiceIndexResponseInvoicesInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InvoiceIndexResponseInvoicesInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InvoiceIndexResponseInvoicesInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *InvoiceIndexResponseInvoicesInner) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *InvoiceIndexResponseInvoicesInner) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetNotes

`func (o *InvoiceIndexResponseInvoicesInner) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InvoiceIndexResponseInvoicesInner) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InvoiceIndexResponseInvoicesInner) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InvoiceIndexResponseInvoicesInner) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *InvoiceIndexResponseInvoicesInner) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *InvoiceIndexResponseInvoicesInner) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetPartnerAddress1

`func (o *InvoiceIndexResponseInvoicesInner) GetPartnerAddress1() string`

GetPartnerAddress1 returns the PartnerAddress1 field if non-nil, zero value otherwise.

### GetPartnerAddress1Ok

`func (o *InvoiceIndexResponseInvoicesInner) GetPartnerAddress1Ok() (*string, bool)`

GetPartnerAddress1Ok returns a tuple with the PartnerAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerAddress1

`func (o *InvoiceIndexResponseInvoicesInner) SetPartnerAddress1(v string)`

SetPartnerAddress1 sets PartnerAddress1 field to given value.

### HasPartnerAddress1

`func (o *InvoiceIndexResponseInvoicesInner) HasPartnerAddress1() bool`

HasPartnerAddress1 returns a boolean if a field has been set.

### SetPartnerAddress1Nil

`func (o *InvoiceIndexResponseInvoicesInner) SetPartnerAddress1Nil(b bool)`

 SetPartnerAddress1Nil sets the value for PartnerAddress1 to be an explicit nil

### UnsetPartnerAddress1
`func (o *InvoiceIndexResponseInvoicesInner) UnsetPartnerAddress1()`

UnsetPartnerAddress1 ensures that no value is present for PartnerAddress1, not even an explicit nil
### GetPartnerAddress2

`func (o *InvoiceIndexResponseInvoicesInner) GetPartnerAddress2() string`

GetPartnerAddress2 returns the PartnerAddress2 field if non-nil, zero value otherwise.

### GetPartnerAddress2Ok

`func (o *InvoiceIndexResponseInvoicesInner) GetPartnerAddress2Ok() (*string, bool)`

GetPartnerAddress2Ok returns a tuple with the PartnerAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerAddress2

`func (o *InvoiceIndexResponseInvoicesInner) SetPartnerAddress2(v string)`

SetPartnerAddress2 sets PartnerAddress2 field to given value.

### HasPartnerAddress2

`func (o *InvoiceIndexResponseInvoicesInner) HasPartnerAddress2() bool`

HasPartnerAddress2 returns a boolean if a field has been set.

### SetPartnerAddress2Nil

`func (o *InvoiceIndexResponseInvoicesInner) SetPartnerAddress2Nil(b bool)`

 SetPartnerAddress2Nil sets the value for PartnerAddress2 to be an explicit nil

### UnsetPartnerAddress2
`func (o *InvoiceIndexResponseInvoicesInner) UnsetPartnerAddress2()`

UnsetPartnerAddress2 ensures that no value is present for PartnerAddress2, not even an explicit nil
### GetPartnerCode

`func (o *InvoiceIndexResponseInvoicesInner) GetPartnerCode() string`

GetPartnerCode returns the PartnerCode field if non-nil, zero value otherwise.

### GetPartnerCodeOk

`func (o *InvoiceIndexResponseInvoicesInner) GetPartnerCodeOk() (*string, bool)`

GetPartnerCodeOk returns a tuple with the PartnerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCode

`func (o *InvoiceIndexResponseInvoicesInner) SetPartnerCode(v string)`

SetPartnerCode sets PartnerCode field to given value.

### HasPartnerCode

`func (o *InvoiceIndexResponseInvoicesInner) HasPartnerCode() bool`

HasPartnerCode returns a boolean if a field has been set.

### SetPartnerCodeNil

`func (o *InvoiceIndexResponseInvoicesInner) SetPartnerCodeNil(b bool)`

 SetPartnerCodeNil sets the value for PartnerCode to be an explicit nil

### UnsetPartnerCode
`func (o *InvoiceIndexResponseInvoicesInner) UnsetPartnerCode()`

UnsetPartnerCode ensures that no value is present for PartnerCode, not even an explicit nil
### GetPartnerContactInfo

`func (o *InvoiceIndexResponseInvoicesInner) GetPartnerContactInfo() string`

GetPartnerContactInfo returns the PartnerContactInfo field if non-nil, zero value otherwise.

### GetPartnerContactInfoOk

`func (o *InvoiceIndexResponseInvoicesInner) GetPartnerContactInfoOk() (*string, bool)`

GetPartnerContactInfoOk returns a tuple with the PartnerContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerContactInfo

`func (o *InvoiceIndexResponseInvoicesInner) SetPartnerContactInfo(v string)`

SetPartnerContactInfo sets PartnerContactInfo field to given value.

### HasPartnerContactInfo

`func (o *InvoiceIndexResponseInvoicesInner) HasPartnerContactInfo() bool`

HasPartnerContactInfo returns a boolean if a field has been set.

### SetPartnerContactInfoNil

`func (o *InvoiceIndexResponseInvoicesInner) SetPartnerContactInfoNil(b bool)`

 SetPartnerContactInfoNil sets the value for PartnerContactInfo to be an explicit nil

### UnsetPartnerContactInfo
`func (o *InvoiceIndexResponseInvoicesInner) UnsetPartnerContactInfo()`

UnsetPartnerContactInfo ensures that no value is present for PartnerContactInfo, not even an explicit nil
### GetPartnerDisplayName

`func (o *InvoiceIndexResponseInvoicesInner) GetPartnerDisplayName() string`

GetPartnerDisplayName returns the PartnerDisplayName field if non-nil, zero value otherwise.

### GetPartnerDisplayNameOk

`func (o *InvoiceIndexResponseInvoicesInner) GetPartnerDisplayNameOk() (*string, bool)`

GetPartnerDisplayNameOk returns a tuple with the PartnerDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerDisplayName

`func (o *InvoiceIndexResponseInvoicesInner) SetPartnerDisplayName(v string)`

SetPartnerDisplayName sets PartnerDisplayName field to given value.

### HasPartnerDisplayName

`func (o *InvoiceIndexResponseInvoicesInner) HasPartnerDisplayName() bool`

HasPartnerDisplayName returns a boolean if a field has been set.

### SetPartnerDisplayNameNil

`func (o *InvoiceIndexResponseInvoicesInner) SetPartnerDisplayNameNil(b bool)`

 SetPartnerDisplayNameNil sets the value for PartnerDisplayName to be an explicit nil

### UnsetPartnerDisplayName
`func (o *InvoiceIndexResponseInvoicesInner) UnsetPartnerDisplayName()`

UnsetPartnerDisplayName ensures that no value is present for PartnerDisplayName, not even an explicit nil
### GetPartnerId

`func (o *InvoiceIndexResponseInvoicesInner) GetPartnerId() int64`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *InvoiceIndexResponseInvoicesInner) GetPartnerIdOk() (*int64, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *InvoiceIndexResponseInvoicesInner) SetPartnerId(v int64)`

SetPartnerId sets PartnerId field to given value.


### SetPartnerIdNil

`func (o *InvoiceIndexResponseInvoicesInner) SetPartnerIdNil(b bool)`

 SetPartnerIdNil sets the value for PartnerId to be an explicit nil

### UnsetPartnerId
`func (o *InvoiceIndexResponseInvoicesInner) UnsetPartnerId()`

UnsetPartnerId ensures that no value is present for PartnerId, not even an explicit nil
### GetPartnerName

`func (o *InvoiceIndexResponseInvoicesInner) GetPartnerName() string`

GetPartnerName returns the PartnerName field if non-nil, zero value otherwise.

### GetPartnerNameOk

`func (o *InvoiceIndexResponseInvoicesInner) GetPartnerNameOk() (*string, bool)`

GetPartnerNameOk returns a tuple with the PartnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerName

`func (o *InvoiceIndexResponseInvoicesInner) SetPartnerName(v string)`

SetPartnerName sets PartnerName field to given value.

### HasPartnerName

`func (o *InvoiceIndexResponseInvoicesInner) HasPartnerName() bool`

HasPartnerName returns a boolean if a field has been set.

### SetPartnerNameNil

`func (o *InvoiceIndexResponseInvoicesInner) SetPartnerNameNil(b bool)`

 SetPartnerNameNil sets the value for PartnerName to be an explicit nil

### UnsetPartnerName
`func (o *InvoiceIndexResponseInvoicesInner) UnsetPartnerName()`

UnsetPartnerName ensures that no value is present for PartnerName, not even an explicit nil
### GetPartnerPrefectureCode

`func (o *InvoiceIndexResponseInvoicesInner) GetPartnerPrefectureCode() int64`

GetPartnerPrefectureCode returns the PartnerPrefectureCode field if non-nil, zero value otherwise.

### GetPartnerPrefectureCodeOk

`func (o *InvoiceIndexResponseInvoicesInner) GetPartnerPrefectureCodeOk() (*int64, bool)`

GetPartnerPrefectureCodeOk returns a tuple with the PartnerPrefectureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPrefectureCode

`func (o *InvoiceIndexResponseInvoicesInner) SetPartnerPrefectureCode(v int64)`

SetPartnerPrefectureCode sets PartnerPrefectureCode field to given value.

### HasPartnerPrefectureCode

`func (o *InvoiceIndexResponseInvoicesInner) HasPartnerPrefectureCode() bool`

HasPartnerPrefectureCode returns a boolean if a field has been set.

### SetPartnerPrefectureCodeNil

`func (o *InvoiceIndexResponseInvoicesInner) SetPartnerPrefectureCodeNil(b bool)`

 SetPartnerPrefectureCodeNil sets the value for PartnerPrefectureCode to be an explicit nil

### UnsetPartnerPrefectureCode
`func (o *InvoiceIndexResponseInvoicesInner) UnsetPartnerPrefectureCode()`

UnsetPartnerPrefectureCode ensures that no value is present for PartnerPrefectureCode, not even an explicit nil
### GetPartnerPrefectureName

`func (o *InvoiceIndexResponseInvoicesInner) GetPartnerPrefectureName() string`

GetPartnerPrefectureName returns the PartnerPrefectureName field if non-nil, zero value otherwise.

### GetPartnerPrefectureNameOk

`func (o *InvoiceIndexResponseInvoicesInner) GetPartnerPrefectureNameOk() (*string, bool)`

GetPartnerPrefectureNameOk returns a tuple with the PartnerPrefectureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPrefectureName

`func (o *InvoiceIndexResponseInvoicesInner) SetPartnerPrefectureName(v string)`

SetPartnerPrefectureName sets PartnerPrefectureName field to given value.

### HasPartnerPrefectureName

`func (o *InvoiceIndexResponseInvoicesInner) HasPartnerPrefectureName() bool`

HasPartnerPrefectureName returns a boolean if a field has been set.

### SetPartnerPrefectureNameNil

`func (o *InvoiceIndexResponseInvoicesInner) SetPartnerPrefectureNameNil(b bool)`

 SetPartnerPrefectureNameNil sets the value for PartnerPrefectureName to be an explicit nil

### UnsetPartnerPrefectureName
`func (o *InvoiceIndexResponseInvoicesInner) UnsetPartnerPrefectureName()`

UnsetPartnerPrefectureName ensures that no value is present for PartnerPrefectureName, not even an explicit nil
### GetPartnerTitle

`func (o *InvoiceIndexResponseInvoicesInner) GetPartnerTitle() string`

GetPartnerTitle returns the PartnerTitle field if non-nil, zero value otherwise.

### GetPartnerTitleOk

`func (o *InvoiceIndexResponseInvoicesInner) GetPartnerTitleOk() (*string, bool)`

GetPartnerTitleOk returns a tuple with the PartnerTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerTitle

`func (o *InvoiceIndexResponseInvoicesInner) SetPartnerTitle(v string)`

SetPartnerTitle sets PartnerTitle field to given value.

### HasPartnerTitle

`func (o *InvoiceIndexResponseInvoicesInner) HasPartnerTitle() bool`

HasPartnerTitle returns a boolean if a field has been set.

### SetPartnerTitleNil

`func (o *InvoiceIndexResponseInvoicesInner) SetPartnerTitleNil(b bool)`

 SetPartnerTitleNil sets the value for PartnerTitle to be an explicit nil

### UnsetPartnerTitle
`func (o *InvoiceIndexResponseInvoicesInner) UnsetPartnerTitle()`

UnsetPartnerTitle ensures that no value is present for PartnerTitle, not even an explicit nil
### GetPartnerZipcode

`func (o *InvoiceIndexResponseInvoicesInner) GetPartnerZipcode() string`

GetPartnerZipcode returns the PartnerZipcode field if non-nil, zero value otherwise.

### GetPartnerZipcodeOk

`func (o *InvoiceIndexResponseInvoicesInner) GetPartnerZipcodeOk() (*string, bool)`

GetPartnerZipcodeOk returns a tuple with the PartnerZipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerZipcode

`func (o *InvoiceIndexResponseInvoicesInner) SetPartnerZipcode(v string)`

SetPartnerZipcode sets PartnerZipcode field to given value.

### HasPartnerZipcode

`func (o *InvoiceIndexResponseInvoicesInner) HasPartnerZipcode() bool`

HasPartnerZipcode returns a boolean if a field has been set.

### SetPartnerZipcodeNil

`func (o *InvoiceIndexResponseInvoicesInner) SetPartnerZipcodeNil(b bool)`

 SetPartnerZipcodeNil sets the value for PartnerZipcode to be an explicit nil

### UnsetPartnerZipcode
`func (o *InvoiceIndexResponseInvoicesInner) UnsetPartnerZipcode()`

UnsetPartnerZipcode ensures that no value is present for PartnerZipcode, not even an explicit nil
### GetPaymentBankInfo

`func (o *InvoiceIndexResponseInvoicesInner) GetPaymentBankInfo() string`

GetPaymentBankInfo returns the PaymentBankInfo field if non-nil, zero value otherwise.

### GetPaymentBankInfoOk

`func (o *InvoiceIndexResponseInvoicesInner) GetPaymentBankInfoOk() (*string, bool)`

GetPaymentBankInfoOk returns a tuple with the PaymentBankInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentBankInfo

`func (o *InvoiceIndexResponseInvoicesInner) SetPaymentBankInfo(v string)`

SetPaymentBankInfo sets PaymentBankInfo field to given value.

### HasPaymentBankInfo

`func (o *InvoiceIndexResponseInvoicesInner) HasPaymentBankInfo() bool`

HasPaymentBankInfo returns a boolean if a field has been set.

### SetPaymentBankInfoNil

`func (o *InvoiceIndexResponseInvoicesInner) SetPaymentBankInfoNil(b bool)`

 SetPaymentBankInfoNil sets the value for PaymentBankInfo to be an explicit nil

### UnsetPaymentBankInfo
`func (o *InvoiceIndexResponseInvoicesInner) UnsetPaymentBankInfo()`

UnsetPaymentBankInfo ensures that no value is present for PaymentBankInfo, not even an explicit nil
### GetPaymentDate

`func (o *InvoiceIndexResponseInvoicesInner) GetPaymentDate() string`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *InvoiceIndexResponseInvoicesInner) GetPaymentDateOk() (*string, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *InvoiceIndexResponseInvoicesInner) SetPaymentDate(v string)`

SetPaymentDate sets PaymentDate field to given value.

### HasPaymentDate

`func (o *InvoiceIndexResponseInvoicesInner) HasPaymentDate() bool`

HasPaymentDate returns a boolean if a field has been set.

### SetPaymentDateNil

`func (o *InvoiceIndexResponseInvoicesInner) SetPaymentDateNil(b bool)`

 SetPaymentDateNil sets the value for PaymentDate to be an explicit nil

### UnsetPaymentDate
`func (o *InvoiceIndexResponseInvoicesInner) UnsetPaymentDate()`

UnsetPaymentDate ensures that no value is present for PaymentDate, not even an explicit nil
### GetPaymentStatus

`func (o *InvoiceIndexResponseInvoicesInner) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *InvoiceIndexResponseInvoicesInner) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *InvoiceIndexResponseInvoicesInner) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *InvoiceIndexResponseInvoicesInner) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetPaymentType

`func (o *InvoiceIndexResponseInvoicesInner) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *InvoiceIndexResponseInvoicesInner) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *InvoiceIndexResponseInvoicesInner) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.


### GetPostingStatus

`func (o *InvoiceIndexResponseInvoicesInner) GetPostingStatus() string`

GetPostingStatus returns the PostingStatus field if non-nil, zero value otherwise.

### GetPostingStatusOk

`func (o *InvoiceIndexResponseInvoicesInner) GetPostingStatusOk() (*string, bool)`

GetPostingStatusOk returns a tuple with the PostingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostingStatus

`func (o *InvoiceIndexResponseInvoicesInner) SetPostingStatus(v string)`

SetPostingStatus sets PostingStatus field to given value.


### GetSubTotal

`func (o *InvoiceIndexResponseInvoicesInner) GetSubTotal() int64`

GetSubTotal returns the SubTotal field if non-nil, zero value otherwise.

### GetSubTotalOk

`func (o *InvoiceIndexResponseInvoicesInner) GetSubTotalOk() (*int64, bool)`

GetSubTotalOk returns a tuple with the SubTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotal

`func (o *InvoiceIndexResponseInvoicesInner) SetSubTotal(v int64)`

SetSubTotal sets SubTotal field to given value.

### HasSubTotal

`func (o *InvoiceIndexResponseInvoicesInner) HasSubTotal() bool`

HasSubTotal returns a boolean if a field has been set.

### GetTaxEntryMethod

`func (o *InvoiceIndexResponseInvoicesInner) GetTaxEntryMethod() string`

GetTaxEntryMethod returns the TaxEntryMethod field if non-nil, zero value otherwise.

### GetTaxEntryMethodOk

`func (o *InvoiceIndexResponseInvoicesInner) GetTaxEntryMethodOk() (*string, bool)`

GetTaxEntryMethodOk returns a tuple with the TaxEntryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxEntryMethod

`func (o *InvoiceIndexResponseInvoicesInner) SetTaxEntryMethod(v string)`

SetTaxEntryMethod sets TaxEntryMethod field to given value.


### GetTitle

`func (o *InvoiceIndexResponseInvoicesInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InvoiceIndexResponseInvoicesInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InvoiceIndexResponseInvoicesInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InvoiceIndexResponseInvoicesInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *InvoiceIndexResponseInvoicesInner) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *InvoiceIndexResponseInvoicesInner) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetTotalAmount

`func (o *InvoiceIndexResponseInvoicesInner) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *InvoiceIndexResponseInvoicesInner) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *InvoiceIndexResponseInvoicesInner) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.


### GetTotalAmountPerVatRate

`func (o *InvoiceIndexResponseInvoicesInner) GetTotalAmountPerVatRate() InvoiceIndexResponseInvoicesInnerTotalAmountPerVatRate`

GetTotalAmountPerVatRate returns the TotalAmountPerVatRate field if non-nil, zero value otherwise.

### GetTotalAmountPerVatRateOk

`func (o *InvoiceIndexResponseInvoicesInner) GetTotalAmountPerVatRateOk() (*InvoiceIndexResponseInvoicesInnerTotalAmountPerVatRate, bool)`

GetTotalAmountPerVatRateOk returns a tuple with the TotalAmountPerVatRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountPerVatRate

`func (o *InvoiceIndexResponseInvoicesInner) SetTotalAmountPerVatRate(v InvoiceIndexResponseInvoicesInnerTotalAmountPerVatRate)`

SetTotalAmountPerVatRate sets TotalAmountPerVatRate field to given value.


### GetTotalVat

`func (o *InvoiceIndexResponseInvoicesInner) GetTotalVat() int64`

GetTotalVat returns the TotalVat field if non-nil, zero value otherwise.

### GetTotalVatOk

`func (o *InvoiceIndexResponseInvoicesInner) GetTotalVatOk() (*int64, bool)`

GetTotalVatOk returns a tuple with the TotalVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVat

`func (o *InvoiceIndexResponseInvoicesInner) SetTotalVat(v int64)`

SetTotalVat sets TotalVat field to given value.

### HasTotalVat

`func (o *InvoiceIndexResponseInvoicesInner) HasTotalVat() bool`

HasTotalVat returns a boolean if a field has been set.

### GetWebConfirmedAt

`func (o *InvoiceIndexResponseInvoicesInner) GetWebConfirmedAt() string`

GetWebConfirmedAt returns the WebConfirmedAt field if non-nil, zero value otherwise.

### GetWebConfirmedAtOk

`func (o *InvoiceIndexResponseInvoicesInner) GetWebConfirmedAtOk() (*string, bool)`

GetWebConfirmedAtOk returns a tuple with the WebConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebConfirmedAt

`func (o *InvoiceIndexResponseInvoicesInner) SetWebConfirmedAt(v string)`

SetWebConfirmedAt sets WebConfirmedAt field to given value.

### HasWebConfirmedAt

`func (o *InvoiceIndexResponseInvoicesInner) HasWebConfirmedAt() bool`

HasWebConfirmedAt returns a boolean if a field has been set.

### SetWebConfirmedAtNil

`func (o *InvoiceIndexResponseInvoicesInner) SetWebConfirmedAtNil(b bool)`

 SetWebConfirmedAtNil sets the value for WebConfirmedAt to be an explicit nil

### UnsetWebConfirmedAt
`func (o *InvoiceIndexResponseInvoicesInner) UnsetWebConfirmedAt()`

UnsetWebConfirmedAt ensures that no value is present for WebConfirmedAt, not even an explicit nil
### GetWebDownloadedAt

`func (o *InvoiceIndexResponseInvoicesInner) GetWebDownloadedAt() string`

GetWebDownloadedAt returns the WebDownloadedAt field if non-nil, zero value otherwise.

### GetWebDownloadedAtOk

`func (o *InvoiceIndexResponseInvoicesInner) GetWebDownloadedAtOk() (*string, bool)`

GetWebDownloadedAtOk returns a tuple with the WebDownloadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebDownloadedAt

`func (o *InvoiceIndexResponseInvoicesInner) SetWebDownloadedAt(v string)`

SetWebDownloadedAt sets WebDownloadedAt field to given value.

### HasWebDownloadedAt

`func (o *InvoiceIndexResponseInvoicesInner) HasWebDownloadedAt() bool`

HasWebDownloadedAt returns a boolean if a field has been set.

### SetWebDownloadedAtNil

`func (o *InvoiceIndexResponseInvoicesInner) SetWebDownloadedAtNil(b bool)`

 SetWebDownloadedAtNil sets the value for WebDownloadedAt to be an explicit nil

### UnsetWebDownloadedAt
`func (o *InvoiceIndexResponseInvoicesInner) UnsetWebDownloadedAt()`

UnsetWebDownloadedAt ensures that no value is present for WebDownloadedAt, not even an explicit nil
### GetWebPublishedAt

`func (o *InvoiceIndexResponseInvoicesInner) GetWebPublishedAt() string`

GetWebPublishedAt returns the WebPublishedAt field if non-nil, zero value otherwise.

### GetWebPublishedAtOk

`func (o *InvoiceIndexResponseInvoicesInner) GetWebPublishedAtOk() (*string, bool)`

GetWebPublishedAtOk returns a tuple with the WebPublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPublishedAt

`func (o *InvoiceIndexResponseInvoicesInner) SetWebPublishedAt(v string)`

SetWebPublishedAt sets WebPublishedAt field to given value.

### HasWebPublishedAt

`func (o *InvoiceIndexResponseInvoicesInner) HasWebPublishedAt() bool`

HasWebPublishedAt returns a boolean if a field has been set.

### SetWebPublishedAtNil

`func (o *InvoiceIndexResponseInvoicesInner) SetWebPublishedAtNil(b bool)`

 SetWebPublishedAtNil sets the value for WebPublishedAt to be an explicit nil

### UnsetWebPublishedAt
`func (o *InvoiceIndexResponseInvoicesInner) UnsetWebPublishedAt()`

UnsetWebPublishedAt ensures that no value is present for WebPublishedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


