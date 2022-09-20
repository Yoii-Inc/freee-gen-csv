# CompanyResponseCompany

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to [**[]CompanyResponseCompanyTagsInner**](CompanyResponseCompanyTagsInner.md) |  | [optional] 
**AccountItems** | Pointer to [**[]CompanyResponseCompanyAccountItemsInner**](CompanyResponseCompanyAccountItemsInner.md) |  | [optional] 
**AmountFraction** | **int64** | 金額端数処理方法（0: 切り捨て、1: 切り上げ、2: 四捨五入） | 
**ContactName** | **NullableString** | 担当者名 (50文字以内) | 
**CorporateNumber** | **string** | 法人番号 (半角数字13桁、法人のみ) | 
**DefaultWalletAccountId** | Pointer to **int64** | 決済口座のデフォルト | [optional] 
**DisplayName** | **string** | 事業所名 | 
**Fax** | **NullableString** | FAX | 
**FiscalYears** | [**[]FiscalYears**](FiscalYears.md) |  | 
**HeadCount** | **NullableInt64** | 従業員数（0: 経営者のみ、1: 2~5人、2: 6~10人、3: 11~20人、4: 21~30人、5: 31~40人、6: 41~100人、7: 100人以上 | 
**Id** | **int64** | 事業所ID | 
**IndustryClass** | **string** | 種別（agriculture_forestry_fisheries_ore: 農林水産業/鉱業、construction: 建設、manufacturing_processing: 製造/加工、it: IT、transportation_logistics: 運輸/物流、retail_wholesale: 小売/卸売、finance_insurance: 金融/保険、real_estate_rental: 不動産/レンタル、profession: 士業/学術/専門技術サービス、design_production: デザイン/制作、food: 飲食、leisure_entertainment: レジャー/娯楽、lifestyle: 生活関連サービス、education: 教育/学習支援、medical_welfare: 医療/福祉、other_services: その他サービス、other: その他） | 
**IndustryCode** | **string** | 業種（agriculture: 農業, forestry: 林業, fishing_industry: 漁業、水産養殖業, mining: 鉱業、採石業、砂利採取業, civil_contractors: 土木工事業, pavement: 舗装工事業, carpenter: とび、大工、左官等の建設工事業, renovation: リフォーム工事業, electrical_plumbing: 電気、管工事等の設備工事業, grocery: 食料品の製造加工業, machinery_manufacturing: 機械器具の製造加工業, printing: 印刷業, other_manufacturing: その他の製造加工業, software_development: 受託：ソフトウェア、アプリ開発業, system_development: 受託：システム開発業, survey_analysis: 受託：調査、分析等の情報処理業, server_management: 受託：サーバー運営管理, website_production: 受託：ウェブサイト制作, online_service_management: オンラインサービス運営業, online_advertising_agency: オンライン広告代理店業, online_advertising_planning_production: オンライン広告企画・制作業, online_media_management: オンラインメディア運営業, portal_site_management: ポータルサイト運営業, other_it_services: その他、IT サービス業, transport_delivery: 輸送業、配送業, delivery: バイク便等の配達業, other_transportation_logistics: その他の運輸業、物流業, other_wholesale: 卸売業：その他, clothing_wholesale_fiber: 卸売業：衣類卸売／繊維, food_wholesale: 卸売業：飲食料品, entrusted_development_wholesale: 卸売業：機械器具, online_shop: 小売業：無店舗　オンラインショップ, fashion_grocery_store: 小売業：店舗あり　ファッション、雑貨, food_store: 小売業：店舗あり　生鮮食品、飲食料品, entrusted_store: 小売業：店舗あり　機械、器具, other_store: 小売業：店舗あり　その他, financial_instruments_exchange: 金融業：金融商品取引, commodity_futures_investment_advisor: 金融業：商品先物取引、商品投資顧問, other_financial: 金融業：その他, brokerage_insurance: 保険業：仲介、代理, other_insurance: 保険業：その他, real_estate_developer: 不動産業：ディベロッパー, real_estate_brokerage: 不動産業：売買、仲介, rent_coin_parking_management: 不動産業：賃貸、コインパーキング、管理, rental_office_co_working_space: 不動産業：レンタルオフィス、コワーキングスペース, rental_lease: レンタル業、リース業, cpa_tax_accountant: 士業：公認会計士事務所、税理士事務所, law_office: 士業：法律事務所, judicial_and_administrative_scrivener: 士業：司法書士事務所／行政書士事務所, labor_consultant: 士業：社会保険労務士事務所, other_profession: 士業：その他, business_consultant: 経営コンサルタント, academic_research_development: 学術・開発研究機関, advertising_agency: 広告代理店, advertising_planning_production: 広告企画／制作, design_development: ソフトウェア、アプリ開発業（受託）, apparel_industry_design: 服飾デザイン業、工業デザイン業, website_design: ウェブサイト制作（受託）, advertising_planning_design: 広告企画／制作業, other_design: その他、デザイン／制作, restaurants_coffee_shops: レストラン、喫茶店等の飲食店業, sale_of_lunch: 弁当の販売業, bread_confectionery_manufacture_sale: パン、菓子等の製造販売業, delivery_catering_mobile_catering: デリバリー業、ケータリング業、移動販売業, hotel_inn: 宿泊業：ホテル、旅館, homestay: 宿泊業：民泊, travel_agency: 旅行代理店業, leisure_sports_facility_management: レジャー、スポーツ等の施設運営業, show_event_management: ショー、イベント等の興行、イベント運営業, barber: ビューティ、ヘルスケア業：床屋、理容室, beauty_salon: ビューティ、ヘルスケア業：美容室, spa_sand_bath_sauna: ビューティ、ヘルスケア業：スパ、砂風呂、サウナ等, este_ail_salon: ビューティ、ヘルスケア業：その他、エステサロン、ネイルサロン等, bridal_planning_introduce_wedding: 冠婚葬祭業：ブライダルプランニング、結婚式場紹介等, memorial_ceremony_funeral: 冠婚葬祭業：メモリアルセレモニー、葬儀等, moving: 引っ越し業, courier_industry: 宅配業, house_maid_cleaning_agency: 家事代行サービス業：無店舗　ハウスメイド、掃除代行等, re_tailoring_clothes: 家事代行サービス業：店舗あり　衣類修理、衣類仕立て直し等, training_institute_management: 研修所等の施設運営業, tutoring_school: 学習塾、進学塾等の教育・学習支援業, music_calligraphy_abacus_classroom: 音楽教室、書道教室、そろばん教室等のの教育・学習支援業, english_school: 英会話スクール等の語学学習支援業, tennis_yoga_judo_school: テニススクール、ヨガ教室、柔道場等のスポーツ指導、支援業, culture_school: その他、カルチャースクール等の教育・学習支援業, seminar_planning_management: セミナー等の企画、運営業, hospital_clinic: 医療業：病院、一般診療所、クリニック等, dental_clinic: 医療業：歯科診療所, other_medical_services: 医療業：その他、医療サービス等, nursery: 福祉業：保育所等、児童向け施設型サービス, nursing_home: 福祉業：老人ホーム等、老人向け施設型サービス, rehabilitation_support_services: 福祉業：療育支援サービス等、障害者等向け施設型サービス, other_welfare: 福祉業：その他、施設型福祉サービス, visit_welfare_service: 福祉業：訪問型福祉サービス, recruitment_temporary_staffing: 人材紹介業、人材派遣業, life_related_recruitment_temporary_staffing: 生活関連サービスの人材紹介業、人材派遣業, car_maintenance_car_repair: 自動車整備業、自動車修理業, machinery_equipment_maintenance_repair: 機械機器類の整備業、修理業, cleaning_maintenance_building_management: 清掃業、メンテナンス業、建物管理業, security: 警備業, other_services: その他のサービス業, npo: NPO, general_incorporated_association: 一般社団法人, general_incorporated_foundation: 一般財団法人, other_association: その他組織) | 
**InvoiceLayout** | **string** | 請求書レイアウト * &#x60;default_classic&#x60; - レイアウト１/クラシック (デフォルト)  * &#x60;standard_classic&#x60; - レイアウト２/クラシック  * &#x60;envelope_classic&#x60; - 封筒１/クラシック  * &#x60;carried_forward_standard_classic&#x60; - レイアウト３（繰越金額欄あり）/クラシック  * &#x60;carried_forward_envelope_classic&#x60; - 封筒２（繰越金額欄あり）/クラシック  * &#x60;default_modern&#x60; - レイアウト１/モダン  * &#x60;standard_modern&#x60; - レイアウト２/モダン  * &#x60;envelope_modern&#x60; - 封筒/モダン | 
**Items** | Pointer to [**[]CompanyResponseCompanyItemsInner**](CompanyResponseCompanyItemsInner.md) |  | [optional] 
**MinusFormat** | **int64** | マイナスの表示方法（0: -、 1: △） | 
**Name** | **NullableString** | 事業所の正式名称 (100文字以内) | 
**NameKana** | **NullableString** | 正式名称フリガナ (100文字以内) | 
**Partners** | Pointer to [**[]CompanyResponseCompanyPartnersInner**](CompanyResponseCompanyPartnersInner.md) |  | [optional] 
**Phone1** | **NullableString** | 電話番号１ | 
**Phone2** | **NullableString** | 電話番号２ | 
**PrefectureCode** | **NullableInt64** | 都道府県コード（-1: 設定しない、0: 北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄 | 
**PrivateSettlement** | **bool** | プライベート資金/役員資金（false: 使用しない、true: 使用する） | 
**Role** | **string** | ユーザーの権限 | 
**Sections** | Pointer to [**[]CompanyResponseCompanySectionsInner**](CompanyResponseCompanySectionsInner.md) |  | [optional] 
**StreetName1** | **string** | 市区町村・番地 | 
**StreetName2** | **string** | 建物名・部屋番号など | 
**TaxAtSourceCalcType** | **int64** | 源泉徴収税計算（0: 消費税を含める、1: 消費税を含めない） | 
**TaxCodes** | Pointer to [**[]CompanyResponseCompanyTaxCodesInner**](CompanyResponseCompanyTaxCodesInner.md) |  | [optional] 
**Taxes** | Pointer to [**[]CompanyResponseCompanyTaxesInner**](CompanyResponseCompanyTaxesInner.md) |  | [optional] 
**TxnNumberFormat** | **string** | 仕訳番号形式（not_used: 使用しない、digits: 数字（例：5091824）、alnum: 英数字（例：59J0P）） | 
**UsePartnerCode** | **bool** | 取引先コードの利用設定（true: 有効、 false: 無効） | 
**Walletables** | Pointer to [**[]CompanyResponseCompanyWalletablesInner**](CompanyResponseCompanyWalletablesInner.md) |  | [optional] 
**WorkflowSetting** | **string** | 仕訳承認フロー（enable: 有効、 disable: 無効） | 
**Zipcode** | **string** | 郵便番号 | 

## Methods

### NewCompanyResponseCompany

`func NewCompanyResponseCompany(amountFraction int64, contactName NullableString, corporateNumber string, displayName string, fax NullableString, fiscalYears []FiscalYears, headCount NullableInt64, id int64, industryClass string, industryCode string, invoiceLayout string, minusFormat int64, name NullableString, nameKana NullableString, phone1 NullableString, phone2 NullableString, prefectureCode NullableInt64, privateSettlement bool, role string, streetName1 string, streetName2 string, taxAtSourceCalcType int64, txnNumberFormat string, usePartnerCode bool, workflowSetting string, zipcode string, ) *CompanyResponseCompany`

NewCompanyResponseCompany instantiates a new CompanyResponseCompany object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyResponseCompanyWithDefaults

`func NewCompanyResponseCompanyWithDefaults() *CompanyResponseCompany`

NewCompanyResponseCompanyWithDefaults instantiates a new CompanyResponseCompany object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *CompanyResponseCompany) GetTags() []CompanyResponseCompanyTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CompanyResponseCompany) GetTagsOk() (*[]CompanyResponseCompanyTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CompanyResponseCompany) SetTags(v []CompanyResponseCompanyTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CompanyResponseCompany) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAccountItems

`func (o *CompanyResponseCompany) GetAccountItems() []CompanyResponseCompanyAccountItemsInner`

GetAccountItems returns the AccountItems field if non-nil, zero value otherwise.

### GetAccountItemsOk

`func (o *CompanyResponseCompany) GetAccountItemsOk() (*[]CompanyResponseCompanyAccountItemsInner, bool)`

GetAccountItemsOk returns a tuple with the AccountItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItems

`func (o *CompanyResponseCompany) SetAccountItems(v []CompanyResponseCompanyAccountItemsInner)`

SetAccountItems sets AccountItems field to given value.

### HasAccountItems

`func (o *CompanyResponseCompany) HasAccountItems() bool`

HasAccountItems returns a boolean if a field has been set.

### GetAmountFraction

`func (o *CompanyResponseCompany) GetAmountFraction() int64`

GetAmountFraction returns the AmountFraction field if non-nil, zero value otherwise.

### GetAmountFractionOk

`func (o *CompanyResponseCompany) GetAmountFractionOk() (*int64, bool)`

GetAmountFractionOk returns a tuple with the AmountFraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountFraction

`func (o *CompanyResponseCompany) SetAmountFraction(v int64)`

SetAmountFraction sets AmountFraction field to given value.


### GetContactName

`func (o *CompanyResponseCompany) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *CompanyResponseCompany) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *CompanyResponseCompany) SetContactName(v string)`

SetContactName sets ContactName field to given value.


### SetContactNameNil

`func (o *CompanyResponseCompany) SetContactNameNil(b bool)`

 SetContactNameNil sets the value for ContactName to be an explicit nil

### UnsetContactName
`func (o *CompanyResponseCompany) UnsetContactName()`

UnsetContactName ensures that no value is present for ContactName, not even an explicit nil
### GetCorporateNumber

`func (o *CompanyResponseCompany) GetCorporateNumber() string`

GetCorporateNumber returns the CorporateNumber field if non-nil, zero value otherwise.

### GetCorporateNumberOk

`func (o *CompanyResponseCompany) GetCorporateNumberOk() (*string, bool)`

GetCorporateNumberOk returns a tuple with the CorporateNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporateNumber

`func (o *CompanyResponseCompany) SetCorporateNumber(v string)`

SetCorporateNumber sets CorporateNumber field to given value.


### GetDefaultWalletAccountId

`func (o *CompanyResponseCompany) GetDefaultWalletAccountId() int64`

GetDefaultWalletAccountId returns the DefaultWalletAccountId field if non-nil, zero value otherwise.

### GetDefaultWalletAccountIdOk

`func (o *CompanyResponseCompany) GetDefaultWalletAccountIdOk() (*int64, bool)`

GetDefaultWalletAccountIdOk returns a tuple with the DefaultWalletAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultWalletAccountId

`func (o *CompanyResponseCompany) SetDefaultWalletAccountId(v int64)`

SetDefaultWalletAccountId sets DefaultWalletAccountId field to given value.

### HasDefaultWalletAccountId

`func (o *CompanyResponseCompany) HasDefaultWalletAccountId() bool`

HasDefaultWalletAccountId returns a boolean if a field has been set.

### GetDisplayName

`func (o *CompanyResponseCompany) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CompanyResponseCompany) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CompanyResponseCompany) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetFax

`func (o *CompanyResponseCompany) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *CompanyResponseCompany) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *CompanyResponseCompany) SetFax(v string)`

SetFax sets Fax field to given value.


### SetFaxNil

`func (o *CompanyResponseCompany) SetFaxNil(b bool)`

 SetFaxNil sets the value for Fax to be an explicit nil

### UnsetFax
`func (o *CompanyResponseCompany) UnsetFax()`

UnsetFax ensures that no value is present for Fax, not even an explicit nil
### GetFiscalYears

`func (o *CompanyResponseCompany) GetFiscalYears() []FiscalYears`

GetFiscalYears returns the FiscalYears field if non-nil, zero value otherwise.

### GetFiscalYearsOk

`func (o *CompanyResponseCompany) GetFiscalYearsOk() (*[]FiscalYears, bool)`

GetFiscalYearsOk returns a tuple with the FiscalYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYears

`func (o *CompanyResponseCompany) SetFiscalYears(v []FiscalYears)`

SetFiscalYears sets FiscalYears field to given value.


### GetHeadCount

`func (o *CompanyResponseCompany) GetHeadCount() int64`

GetHeadCount returns the HeadCount field if non-nil, zero value otherwise.

### GetHeadCountOk

`func (o *CompanyResponseCompany) GetHeadCountOk() (*int64, bool)`

GetHeadCountOk returns a tuple with the HeadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadCount

`func (o *CompanyResponseCompany) SetHeadCount(v int64)`

SetHeadCount sets HeadCount field to given value.


### SetHeadCountNil

`func (o *CompanyResponseCompany) SetHeadCountNil(b bool)`

 SetHeadCountNil sets the value for HeadCount to be an explicit nil

### UnsetHeadCount
`func (o *CompanyResponseCompany) UnsetHeadCount()`

UnsetHeadCount ensures that no value is present for HeadCount, not even an explicit nil
### GetId

`func (o *CompanyResponseCompany) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyResponseCompany) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyResponseCompany) SetId(v int64)`

SetId sets Id field to given value.


### GetIndustryClass

`func (o *CompanyResponseCompany) GetIndustryClass() string`

GetIndustryClass returns the IndustryClass field if non-nil, zero value otherwise.

### GetIndustryClassOk

`func (o *CompanyResponseCompany) GetIndustryClassOk() (*string, bool)`

GetIndustryClassOk returns a tuple with the IndustryClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryClass

`func (o *CompanyResponseCompany) SetIndustryClass(v string)`

SetIndustryClass sets IndustryClass field to given value.


### GetIndustryCode

`func (o *CompanyResponseCompany) GetIndustryCode() string`

GetIndustryCode returns the IndustryCode field if non-nil, zero value otherwise.

### GetIndustryCodeOk

`func (o *CompanyResponseCompany) GetIndustryCodeOk() (*string, bool)`

GetIndustryCodeOk returns a tuple with the IndustryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryCode

`func (o *CompanyResponseCompany) SetIndustryCode(v string)`

SetIndustryCode sets IndustryCode field to given value.


### GetInvoiceLayout

`func (o *CompanyResponseCompany) GetInvoiceLayout() string`

GetInvoiceLayout returns the InvoiceLayout field if non-nil, zero value otherwise.

### GetInvoiceLayoutOk

`func (o *CompanyResponseCompany) GetInvoiceLayoutOk() (*string, bool)`

GetInvoiceLayoutOk returns a tuple with the InvoiceLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceLayout

`func (o *CompanyResponseCompany) SetInvoiceLayout(v string)`

SetInvoiceLayout sets InvoiceLayout field to given value.


### GetItems

`func (o *CompanyResponseCompany) GetItems() []CompanyResponseCompanyItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CompanyResponseCompany) GetItemsOk() (*[]CompanyResponseCompanyItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CompanyResponseCompany) SetItems(v []CompanyResponseCompanyItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *CompanyResponseCompany) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMinusFormat

`func (o *CompanyResponseCompany) GetMinusFormat() int64`

GetMinusFormat returns the MinusFormat field if non-nil, zero value otherwise.

### GetMinusFormatOk

`func (o *CompanyResponseCompany) GetMinusFormatOk() (*int64, bool)`

GetMinusFormatOk returns a tuple with the MinusFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinusFormat

`func (o *CompanyResponseCompany) SetMinusFormat(v int64)`

SetMinusFormat sets MinusFormat field to given value.


### GetName

`func (o *CompanyResponseCompany) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyResponseCompany) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyResponseCompany) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CompanyResponseCompany) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CompanyResponseCompany) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNameKana

`func (o *CompanyResponseCompany) GetNameKana() string`

GetNameKana returns the NameKana field if non-nil, zero value otherwise.

### GetNameKanaOk

`func (o *CompanyResponseCompany) GetNameKanaOk() (*string, bool)`

GetNameKanaOk returns a tuple with the NameKana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameKana

`func (o *CompanyResponseCompany) SetNameKana(v string)`

SetNameKana sets NameKana field to given value.


### SetNameKanaNil

`func (o *CompanyResponseCompany) SetNameKanaNil(b bool)`

 SetNameKanaNil sets the value for NameKana to be an explicit nil

### UnsetNameKana
`func (o *CompanyResponseCompany) UnsetNameKana()`

UnsetNameKana ensures that no value is present for NameKana, not even an explicit nil
### GetPartners

`func (o *CompanyResponseCompany) GetPartners() []CompanyResponseCompanyPartnersInner`

GetPartners returns the Partners field if non-nil, zero value otherwise.

### GetPartnersOk

`func (o *CompanyResponseCompany) GetPartnersOk() (*[]CompanyResponseCompanyPartnersInner, bool)`

GetPartnersOk returns a tuple with the Partners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartners

`func (o *CompanyResponseCompany) SetPartners(v []CompanyResponseCompanyPartnersInner)`

SetPartners sets Partners field to given value.

### HasPartners

`func (o *CompanyResponseCompany) HasPartners() bool`

HasPartners returns a boolean if a field has been set.

### GetPhone1

`func (o *CompanyResponseCompany) GetPhone1() string`

GetPhone1 returns the Phone1 field if non-nil, zero value otherwise.

### GetPhone1Ok

`func (o *CompanyResponseCompany) GetPhone1Ok() (*string, bool)`

GetPhone1Ok returns a tuple with the Phone1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone1

`func (o *CompanyResponseCompany) SetPhone1(v string)`

SetPhone1 sets Phone1 field to given value.


### SetPhone1Nil

`func (o *CompanyResponseCompany) SetPhone1Nil(b bool)`

 SetPhone1Nil sets the value for Phone1 to be an explicit nil

### UnsetPhone1
`func (o *CompanyResponseCompany) UnsetPhone1()`

UnsetPhone1 ensures that no value is present for Phone1, not even an explicit nil
### GetPhone2

`func (o *CompanyResponseCompany) GetPhone2() string`

GetPhone2 returns the Phone2 field if non-nil, zero value otherwise.

### GetPhone2Ok

`func (o *CompanyResponseCompany) GetPhone2Ok() (*string, bool)`

GetPhone2Ok returns a tuple with the Phone2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone2

`func (o *CompanyResponseCompany) SetPhone2(v string)`

SetPhone2 sets Phone2 field to given value.


### SetPhone2Nil

`func (o *CompanyResponseCompany) SetPhone2Nil(b bool)`

 SetPhone2Nil sets the value for Phone2 to be an explicit nil

### UnsetPhone2
`func (o *CompanyResponseCompany) UnsetPhone2()`

UnsetPhone2 ensures that no value is present for Phone2, not even an explicit nil
### GetPrefectureCode

`func (o *CompanyResponseCompany) GetPrefectureCode() int64`

GetPrefectureCode returns the PrefectureCode field if non-nil, zero value otherwise.

### GetPrefectureCodeOk

`func (o *CompanyResponseCompany) GetPrefectureCodeOk() (*int64, bool)`

GetPrefectureCodeOk returns a tuple with the PrefectureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefectureCode

`func (o *CompanyResponseCompany) SetPrefectureCode(v int64)`

SetPrefectureCode sets PrefectureCode field to given value.


### SetPrefectureCodeNil

`func (o *CompanyResponseCompany) SetPrefectureCodeNil(b bool)`

 SetPrefectureCodeNil sets the value for PrefectureCode to be an explicit nil

### UnsetPrefectureCode
`func (o *CompanyResponseCompany) UnsetPrefectureCode()`

UnsetPrefectureCode ensures that no value is present for PrefectureCode, not even an explicit nil
### GetPrivateSettlement

`func (o *CompanyResponseCompany) GetPrivateSettlement() bool`

GetPrivateSettlement returns the PrivateSettlement field if non-nil, zero value otherwise.

### GetPrivateSettlementOk

`func (o *CompanyResponseCompany) GetPrivateSettlementOk() (*bool, bool)`

GetPrivateSettlementOk returns a tuple with the PrivateSettlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateSettlement

`func (o *CompanyResponseCompany) SetPrivateSettlement(v bool)`

SetPrivateSettlement sets PrivateSettlement field to given value.


### GetRole

`func (o *CompanyResponseCompany) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CompanyResponseCompany) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CompanyResponseCompany) SetRole(v string)`

SetRole sets Role field to given value.


### GetSections

`func (o *CompanyResponseCompany) GetSections() []CompanyResponseCompanySectionsInner`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *CompanyResponseCompany) GetSectionsOk() (*[]CompanyResponseCompanySectionsInner, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *CompanyResponseCompany) SetSections(v []CompanyResponseCompanySectionsInner)`

SetSections sets Sections field to given value.

### HasSections

`func (o *CompanyResponseCompany) HasSections() bool`

HasSections returns a boolean if a field has been set.

### GetStreetName1

`func (o *CompanyResponseCompany) GetStreetName1() string`

GetStreetName1 returns the StreetName1 field if non-nil, zero value otherwise.

### GetStreetName1Ok

`func (o *CompanyResponseCompany) GetStreetName1Ok() (*string, bool)`

GetStreetName1Ok returns a tuple with the StreetName1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetName1

`func (o *CompanyResponseCompany) SetStreetName1(v string)`

SetStreetName1 sets StreetName1 field to given value.


### GetStreetName2

`func (o *CompanyResponseCompany) GetStreetName2() string`

GetStreetName2 returns the StreetName2 field if non-nil, zero value otherwise.

### GetStreetName2Ok

`func (o *CompanyResponseCompany) GetStreetName2Ok() (*string, bool)`

GetStreetName2Ok returns a tuple with the StreetName2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetName2

`func (o *CompanyResponseCompany) SetStreetName2(v string)`

SetStreetName2 sets StreetName2 field to given value.


### GetTaxAtSourceCalcType

`func (o *CompanyResponseCompany) GetTaxAtSourceCalcType() int64`

GetTaxAtSourceCalcType returns the TaxAtSourceCalcType field if non-nil, zero value otherwise.

### GetTaxAtSourceCalcTypeOk

`func (o *CompanyResponseCompany) GetTaxAtSourceCalcTypeOk() (*int64, bool)`

GetTaxAtSourceCalcTypeOk returns a tuple with the TaxAtSourceCalcType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAtSourceCalcType

`func (o *CompanyResponseCompany) SetTaxAtSourceCalcType(v int64)`

SetTaxAtSourceCalcType sets TaxAtSourceCalcType field to given value.


### GetTaxCodes

`func (o *CompanyResponseCompany) GetTaxCodes() []CompanyResponseCompanyTaxCodesInner`

GetTaxCodes returns the TaxCodes field if non-nil, zero value otherwise.

### GetTaxCodesOk

`func (o *CompanyResponseCompany) GetTaxCodesOk() (*[]CompanyResponseCompanyTaxCodesInner, bool)`

GetTaxCodesOk returns a tuple with the TaxCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCodes

`func (o *CompanyResponseCompany) SetTaxCodes(v []CompanyResponseCompanyTaxCodesInner)`

SetTaxCodes sets TaxCodes field to given value.

### HasTaxCodes

`func (o *CompanyResponseCompany) HasTaxCodes() bool`

HasTaxCodes returns a boolean if a field has been set.

### GetTaxes

`func (o *CompanyResponseCompany) GetTaxes() []CompanyResponseCompanyTaxesInner`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *CompanyResponseCompany) GetTaxesOk() (*[]CompanyResponseCompanyTaxesInner, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *CompanyResponseCompany) SetTaxes(v []CompanyResponseCompanyTaxesInner)`

SetTaxes sets Taxes field to given value.

### HasTaxes

`func (o *CompanyResponseCompany) HasTaxes() bool`

HasTaxes returns a boolean if a field has been set.

### GetTxnNumberFormat

`func (o *CompanyResponseCompany) GetTxnNumberFormat() string`

GetTxnNumberFormat returns the TxnNumberFormat field if non-nil, zero value otherwise.

### GetTxnNumberFormatOk

`func (o *CompanyResponseCompany) GetTxnNumberFormatOk() (*string, bool)`

GetTxnNumberFormatOk returns a tuple with the TxnNumberFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnNumberFormat

`func (o *CompanyResponseCompany) SetTxnNumberFormat(v string)`

SetTxnNumberFormat sets TxnNumberFormat field to given value.


### GetUsePartnerCode

`func (o *CompanyResponseCompany) GetUsePartnerCode() bool`

GetUsePartnerCode returns the UsePartnerCode field if non-nil, zero value otherwise.

### GetUsePartnerCodeOk

`func (o *CompanyResponseCompany) GetUsePartnerCodeOk() (*bool, bool)`

GetUsePartnerCodeOk returns a tuple with the UsePartnerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePartnerCode

`func (o *CompanyResponseCompany) SetUsePartnerCode(v bool)`

SetUsePartnerCode sets UsePartnerCode field to given value.


### GetWalletables

`func (o *CompanyResponseCompany) GetWalletables() []CompanyResponseCompanyWalletablesInner`

GetWalletables returns the Walletables field if non-nil, zero value otherwise.

### GetWalletablesOk

`func (o *CompanyResponseCompany) GetWalletablesOk() (*[]CompanyResponseCompanyWalletablesInner, bool)`

GetWalletablesOk returns a tuple with the Walletables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletables

`func (o *CompanyResponseCompany) SetWalletables(v []CompanyResponseCompanyWalletablesInner)`

SetWalletables sets Walletables field to given value.

### HasWalletables

`func (o *CompanyResponseCompany) HasWalletables() bool`

HasWalletables returns a boolean if a field has been set.

### GetWorkflowSetting

`func (o *CompanyResponseCompany) GetWorkflowSetting() string`

GetWorkflowSetting returns the WorkflowSetting field if non-nil, zero value otherwise.

### GetWorkflowSettingOk

`func (o *CompanyResponseCompany) GetWorkflowSettingOk() (*string, bool)`

GetWorkflowSettingOk returns a tuple with the WorkflowSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowSetting

`func (o *CompanyResponseCompany) SetWorkflowSetting(v string)`

SetWorkflowSetting sets WorkflowSetting field to given value.


### GetZipcode

`func (o *CompanyResponseCompany) GetZipcode() string`

GetZipcode returns the Zipcode field if non-nil, zero value otherwise.

### GetZipcodeOk

`func (o *CompanyResponseCompany) GetZipcodeOk() (*string, bool)`

GetZipcodeOk returns a tuple with the Zipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipcode

`func (o *CompanyResponseCompany) SetZipcode(v string)`

SetZipcode sets Zipcode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


