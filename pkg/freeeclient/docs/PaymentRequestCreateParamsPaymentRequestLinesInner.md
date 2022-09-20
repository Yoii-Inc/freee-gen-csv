# PaymentRequestCreateParamsPaymentRequestLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountItemId** | Pointer to **NullableInt32** | 勘定科目ID | [optional] 
**Amount** | **int64** | 金額 | 
**Description** | Pointer to **string** | 内容 | [optional] 
**ItemId** | Pointer to **int32** | 品目ID | [optional] 
**LineType** | Pointer to **string** | &#39;行の種類 (deal_line: 支払依頼, withholding_tax: 源泉徴収税)&#39;&lt;br&gt; &#39;デフォルトは deal_line: 支払依頼 です&#39;  | [optional] 
**SectionId** | Pointer to **int32** | 部門ID | [optional] 
**Segment1TagId** | Pointer to **int64** | セグメント１ID&lt;br&gt; セグメントタグ一覧APIを利用して取得してください。&lt;br&gt; &lt;a href&#x3D;\&quot;https://support.freee.co.jp/hc/ja/articles/360020679611\&quot; target&#x3D;\&quot;_blank\&quot;&gt;セグメント（分析用タグ）の設定&lt;/a&gt;&lt;br&gt;  | [optional] 
**Segment2TagId** | Pointer to **int64** | セグメント２ID(法人向けエンタープライズプラン)&lt;br&gt; セグメントタグ一覧APIを利用して取得してください。&lt;br&gt; &lt;a href&#x3D;\&quot;https://support.freee.co.jp/hc/ja/articles/360020679611\&quot; target&#x3D;\&quot;_blank\&quot;&gt;セグメント（分析用タグ）の設定&lt;/a&gt;&lt;br&gt;  | [optional] 
**Segment3TagId** | Pointer to **int64** | セグメント３ID(法人向けエンタープライズプラン)&lt;br&gt; セグメントタグ一覧APIを利用して取得してください。&lt;br&gt; &lt;a href&#x3D;\&quot;https://support.freee.co.jp/hc/ja/articles/360020679611\&quot; target&#x3D;\&quot;_blank\&quot;&gt;セグメント（分析用タグ）の設定&lt;/a&gt;&lt;br&gt;  | [optional] 
**TagIds** | Pointer to **[]int32** | メモタグID | [optional] 
**TaxCode** | Pointer to **int32** | 税区分コード&lt;br&gt; 勘定科目IDを指定する場合は必須です。  | [optional] 

## Methods

### NewPaymentRequestCreateParamsPaymentRequestLinesInner

`func NewPaymentRequestCreateParamsPaymentRequestLinesInner(amount int64, ) *PaymentRequestCreateParamsPaymentRequestLinesInner`

NewPaymentRequestCreateParamsPaymentRequestLinesInner instantiates a new PaymentRequestCreateParamsPaymentRequestLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRequestCreateParamsPaymentRequestLinesInnerWithDefaults

`func NewPaymentRequestCreateParamsPaymentRequestLinesInnerWithDefaults() *PaymentRequestCreateParamsPaymentRequestLinesInner`

NewPaymentRequestCreateParamsPaymentRequestLinesInnerWithDefaults instantiates a new PaymentRequestCreateParamsPaymentRequestLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountItemId

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetAccountItemId() int32`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetAccountItemIdOk() (*int32, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) SetAccountItemId(v int32)`

SetAccountItemId sets AccountItemId field to given value.

### HasAccountItemId

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) HasAccountItemId() bool`

HasAccountItemId returns a boolean if a field has been set.

### SetAccountItemIdNil

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) SetAccountItemIdNil(b bool)`

 SetAccountItemIdNil sets the value for AccountItemId to be an explicit nil

### UnsetAccountItemId
`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) UnsetAccountItemId()`

UnsetAccountItemId ensures that no value is present for AccountItemId, not even an explicit nil
### GetAmount

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetItemId

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetItemId() int32`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetItemIdOk() (*int32, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) SetItemId(v int32)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetLineType

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetLineType() string`

GetLineType returns the LineType field if non-nil, zero value otherwise.

### GetLineTypeOk

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetLineTypeOk() (*string, bool)`

GetLineTypeOk returns a tuple with the LineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineType

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) SetLineType(v string)`

SetLineType sets LineType field to given value.

### HasLineType

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) HasLineType() bool`

HasLineType returns a boolean if a field has been set.

### GetSectionId

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetSectionId() int32`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetSectionIdOk() (*int32, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) SetSectionId(v int32)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### GetSegment1TagId

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetSegment1TagId() int64`

GetSegment1TagId returns the Segment1TagId field if non-nil, zero value otherwise.

### GetSegment1TagIdOk

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetSegment1TagIdOk() (*int64, bool)`

GetSegment1TagIdOk returns a tuple with the Segment1TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagId

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) SetSegment1TagId(v int64)`

SetSegment1TagId sets Segment1TagId field to given value.

### HasSegment1TagId

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) HasSegment1TagId() bool`

HasSegment1TagId returns a boolean if a field has been set.

### GetSegment2TagId

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetSegment2TagId() int64`

GetSegment2TagId returns the Segment2TagId field if non-nil, zero value otherwise.

### GetSegment2TagIdOk

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetSegment2TagIdOk() (*int64, bool)`

GetSegment2TagIdOk returns a tuple with the Segment2TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagId

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) SetSegment2TagId(v int64)`

SetSegment2TagId sets Segment2TagId field to given value.

### HasSegment2TagId

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) HasSegment2TagId() bool`

HasSegment2TagId returns a boolean if a field has been set.

### GetSegment3TagId

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetSegment3TagId() int64`

GetSegment3TagId returns the Segment3TagId field if non-nil, zero value otherwise.

### GetSegment3TagIdOk

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetSegment3TagIdOk() (*int64, bool)`

GetSegment3TagIdOk returns a tuple with the Segment3TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagId

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) SetSegment3TagId(v int64)`

SetSegment3TagId sets Segment3TagId field to given value.

### HasSegment3TagId

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) HasSegment3TagId() bool`

HasSegment3TagId returns a boolean if a field has been set.

### GetTagIds

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTaxCode

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetTaxCode() int32`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) GetTaxCodeOk() (*int32, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) SetTaxCode(v int32)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *PaymentRequestCreateParamsPaymentRequestLinesInner) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


