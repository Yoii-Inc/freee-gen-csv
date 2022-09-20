# PaymentRequestUpdateParamsPaymentRequestLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountItemId** | Pointer to **NullableInt32** | 勘定科目ID | [optional] 
**Amount** | **int64** | 金額 | 
**Description** | Pointer to **string** | 内容 | [optional] 
**Id** | Pointer to **int64** | 支払依頼の項目行ID: 既存項目行を更新する場合に指定します。IDを指定しない項目行は、新規行として扱われ追加されます。また、payment_request_linesに含まれない既存の項目行は削除されます。更新後も残したい行は、必ず支払依頼の項目行IDを指定してpayment_request_linesに含めてください。 | [optional] 
**ItemId** | Pointer to **int32** | 品目ID | [optional] 
**LineType** | Pointer to **string** | &#39;行の種類 (deal_line: 支払依頼, withholding_tax: 源泉徴収税)&#39;&lt;br&gt; &#39;デフォルトは deal_line: 支払依頼 です&#39;  | [optional] 
**SectionId** | Pointer to **int32** | 部門ID | [optional] 
**Segment1TagId** | Pointer to **int64** | セグメント１ID&lt;br&gt; セグメントタグ一覧APIを利用して取得してください。&lt;br&gt; &lt;a href&#x3D;\&quot;https://support.freee.co.jp/hc/ja/articles/360020679611\&quot; target&#x3D;\&quot;_blank\&quot;&gt;セグメント（分析用タグ）の設定&lt;/a&gt;&lt;br&gt;  | [optional] 
**Segment2TagId** | Pointer to **int64** | セグメント２ID(法人向けエンタープライズプラン)&lt;br&gt; セグメントタグ一覧APIを利用して取得してください。&lt;br&gt; &lt;a href&#x3D;\&quot;https://support.freee.co.jp/hc/ja/articles/360020679611\&quot; target&#x3D;\&quot;_blank\&quot;&gt;セグメント（分析用タグ）の設定&lt;/a&gt;&lt;br&gt;  | [optional] 
**Segment3TagId** | Pointer to **int64** | セグメント３ID(法人向けエンタープライズプラン)&lt;br&gt; セグメントタグ一覧APIを利用して取得してください。&lt;br&gt; &lt;a href&#x3D;\&quot;https://support.freee.co.jp/hc/ja/articles/360020679611\&quot; target&#x3D;\&quot;_blank\&quot;&gt;セグメント（分析用タグ）の設定&lt;/a&gt;&lt;br&gt;  | [optional] 
**TagIds** | Pointer to **[]int32** | メモタグID | [optional] 
**TaxCode** | Pointer to **int32** | 税区分コード&lt;br&gt; 勘定科目IDを指定する場合は必須です。  | [optional] 

## Methods

### NewPaymentRequestUpdateParamsPaymentRequestLinesInner

`func NewPaymentRequestUpdateParamsPaymentRequestLinesInner(amount int64, ) *PaymentRequestUpdateParamsPaymentRequestLinesInner`

NewPaymentRequestUpdateParamsPaymentRequestLinesInner instantiates a new PaymentRequestUpdateParamsPaymentRequestLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRequestUpdateParamsPaymentRequestLinesInnerWithDefaults

`func NewPaymentRequestUpdateParamsPaymentRequestLinesInnerWithDefaults() *PaymentRequestUpdateParamsPaymentRequestLinesInner`

NewPaymentRequestUpdateParamsPaymentRequestLinesInnerWithDefaults instantiates a new PaymentRequestUpdateParamsPaymentRequestLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountItemId

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetAccountItemId() int32`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetAccountItemIdOk() (*int32, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) SetAccountItemId(v int32)`

SetAccountItemId sets AccountItemId field to given value.

### HasAccountItemId

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) HasAccountItemId() bool`

HasAccountItemId returns a boolean if a field has been set.

### SetAccountItemIdNil

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) SetAccountItemIdNil(b bool)`

 SetAccountItemIdNil sets the value for AccountItemId to be an explicit nil

### UnsetAccountItemId
`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) UnsetAccountItemId()`

UnsetAccountItemId ensures that no value is present for AccountItemId, not even an explicit nil
### GetAmount

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItemId

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetItemId() int32`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetItemIdOk() (*int32, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) SetItemId(v int32)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetLineType

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetLineType() string`

GetLineType returns the LineType field if non-nil, zero value otherwise.

### GetLineTypeOk

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetLineTypeOk() (*string, bool)`

GetLineTypeOk returns a tuple with the LineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineType

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) SetLineType(v string)`

SetLineType sets LineType field to given value.

### HasLineType

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) HasLineType() bool`

HasLineType returns a boolean if a field has been set.

### GetSectionId

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetSectionId() int32`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetSectionIdOk() (*int32, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) SetSectionId(v int32)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### GetSegment1TagId

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetSegment1TagId() int64`

GetSegment1TagId returns the Segment1TagId field if non-nil, zero value otherwise.

### GetSegment1TagIdOk

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetSegment1TagIdOk() (*int64, bool)`

GetSegment1TagIdOk returns a tuple with the Segment1TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagId

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) SetSegment1TagId(v int64)`

SetSegment1TagId sets Segment1TagId field to given value.

### HasSegment1TagId

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) HasSegment1TagId() bool`

HasSegment1TagId returns a boolean if a field has been set.

### GetSegment2TagId

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetSegment2TagId() int64`

GetSegment2TagId returns the Segment2TagId field if non-nil, zero value otherwise.

### GetSegment2TagIdOk

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetSegment2TagIdOk() (*int64, bool)`

GetSegment2TagIdOk returns a tuple with the Segment2TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagId

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) SetSegment2TagId(v int64)`

SetSegment2TagId sets Segment2TagId field to given value.

### HasSegment2TagId

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) HasSegment2TagId() bool`

HasSegment2TagId returns a boolean if a field has been set.

### GetSegment3TagId

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetSegment3TagId() int64`

GetSegment3TagId returns the Segment3TagId field if non-nil, zero value otherwise.

### GetSegment3TagIdOk

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetSegment3TagIdOk() (*int64, bool)`

GetSegment3TagIdOk returns a tuple with the Segment3TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagId

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) SetSegment3TagId(v int64)`

SetSegment3TagId sets Segment3TagId field to given value.

### HasSegment3TagId

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) HasSegment3TagId() bool`

HasSegment3TagId returns a boolean if a field has been set.

### GetTagIds

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTaxCode

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetTaxCode() int32`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) GetTaxCodeOk() (*int32, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) SetTaxCode(v int32)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *PaymentRequestUpdateParamsPaymentRequestLinesInner) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


