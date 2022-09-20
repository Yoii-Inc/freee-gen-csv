# \PaymentRequestsApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePaymentRequest**](PaymentRequestsApi.md#CreatePaymentRequest) | **Post** /api/1/payment_requests | 支払依頼の作成
[**DestroyPaymentRequest**](PaymentRequestsApi.md#DestroyPaymentRequest) | **Delete** /api/1/payment_requests/{id} | 支払依頼の削除
[**GetPaymentRequest**](PaymentRequestsApi.md#GetPaymentRequest) | **Get** /api/1/payment_requests/{id} | 支払依頼詳細の取得
[**GetPaymentRequests**](PaymentRequestsApi.md#GetPaymentRequests) | **Get** /api/1/payment_requests | 支払依頼一覧の取得
[**UpdatePaymentRequest**](PaymentRequestsApi.md#UpdatePaymentRequest) | **Put** /api/1/payment_requests/{id} | 支払依頼の更新
[**UpdatePaymentRequestAction**](PaymentRequestsApi.md#UpdatePaymentRequestAction) | **Post** /api/1/payment_requests/{id}/actions | 支払依頼の承認操作



## CreatePaymentRequest

> PaymentRequestResponse CreatePaymentRequest(ctx).PaymentRequestCreateParams(paymentRequestCreateParams).Execute()

支払依頼の作成



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    paymentRequestCreateParams := *openapiclient.NewPaymentRequestCreateParams(int32(1), int32(1), true, "2019-12-17", []openapiclient.PaymentRequestCreateParamsPaymentRequestLinesInner{*openapiclient.NewPaymentRequestCreateParamsPaymentRequestLinesInner(int64(30000))}, "仕入代金支払い") // PaymentRequestCreateParams | 支払依頼の作成 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentRequestsApi.CreatePaymentRequest(context.Background()).PaymentRequestCreateParams(paymentRequestCreateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentRequestsApi.CreatePaymentRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePaymentRequest`: PaymentRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentRequestsApi.CreatePaymentRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentRequestCreateParams** | [**PaymentRequestCreateParams**](PaymentRequestCreateParams.md) | 支払依頼の作成 | 

### Return type

[**PaymentRequestResponse**](PaymentRequestResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyPaymentRequest

> DestroyPaymentRequest(ctx, id).CompanyId(companyId).Execute()

支払依頼の削除



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | 支払依頼ID
    companyId := int32(56) // int32 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentRequestsApi.DestroyPaymentRequest(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentRequestsApi.DestroyPaymentRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 支払依頼ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyPaymentRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyId** | **int32** | 事業所ID | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentRequest

> PaymentRequestResponse GetPaymentRequest(ctx, id).CompanyId(companyId).Execute()

支払依頼詳細の取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | 支払依頼ID
    companyId := int32(56) // int32 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentRequestsApi.GetPaymentRequest(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentRequestsApi.GetPaymentRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentRequest`: PaymentRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentRequestsApi.GetPaymentRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 支払依頼ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyId** | **int32** | 事業所ID | 

### Return type

[**PaymentRequestResponse**](PaymentRequestResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentRequests

> PaymentRequestsIndexResponse GetPaymentRequests(ctx).CompanyId(companyId).Status(status).StartApplicationDate(startApplicationDate).EndApplicationDate(endApplicationDate).StartIssueDate(startIssueDate).EndIssueDate(endIssueDate).ApplicationNumber(applicationNumber).Title(title).ApplicantId(applicantId).ApproverId(approverId).MinAmount(minAmount).MaxAmount(maxAmount).PartnerId(partnerId).PartnerCode(partnerCode).PaymentMethod(paymentMethod).StartPaymentDate(startPaymentDate).EndPaymentDate(endPaymentDate).DocumentCode(documentCode).Offset(offset).Limit(limit).Execute()

支払依頼一覧の取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyId := int32(56) // int32 | 事業所ID
    status := "status_example" // string | '申請ステータス(draft:下書き, in_progress:申請中, approved:承認済, rejected:却下, feedback:差戻し)、 取引ステータス(unsettled:支払待ち, settled:支払済み)'<br> approver_id を指定した場合は無効です。  (optional)
    startApplicationDate := "2019-12-17" // string | 申請日で絞込：開始日(yyyy-mm-dd) (optional)
    endApplicationDate := "2019-12-17" // string | 申請日で絞込：終了日(yyyy-mm-dd) (optional)
    startIssueDate := "2019-12-17" // string | 発生日で絞込：開始日(yyyy-mm-dd) (optional)
    endIssueDate := "2019-12-17" // string | 発生日で絞込：終了日(yyyy-mm-dd) (optional)
    applicationNumber := int32(2) // int32 | 申請No. (optional)
    title := "大阪出張" // string | 申請タイトル (optional)
    applicantId := int32(1) // int32 | 申請者のユーザーID (optional)
    approverId := int32(1) // int32 | 承認者のユーザーID<br> 'approver_id を指定した場合は、 in_progress: 申請中 の支払依頼のみを返します。'  (optional)
    minAmount := int32(5000) // int32 | 金額で絞込 (下限金額) (optional)
    maxAmount := int32(10000) // int32 | 金額で絞込 (上限金額) (optional)
    partnerId := int32(56) // int32 | 取引先IDで絞込 (optional)
    partnerCode := "partnerCode_example" // string | 取引先コードで絞込 (optional)
    paymentMethod := "credit_card" // string | 支払方法で絞込 (none: 指定なし, domestic_bank_transfer: 国内振込, abroad_bank_transfer: 国外振込, account_transfer: 口座振替, credit_card: クレジットカード) (optional)
    startPaymentDate := "2019-12-17" // string | 支払期限で絞込：開始日(yyyy-mm-dd) (optional)
    endPaymentDate := "2019-12-17" // string | 支払期限で絞込：終了日(yyyy-mm-dd) (optional)
    documentCode := "2" // string | 請求書番号で絞込 (optional)
    offset := int64(789) // int64 | 取得レコードのオフセット (デフォルト: 0) (optional)
    limit := int32(56) // int32 | 取得レコードの件数 (デフォルト: 50, 最小: 1, 最大: 500) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentRequestsApi.GetPaymentRequests(context.Background()).CompanyId(companyId).Status(status).StartApplicationDate(startApplicationDate).EndApplicationDate(endApplicationDate).StartIssueDate(startIssueDate).EndIssueDate(endIssueDate).ApplicationNumber(applicationNumber).Title(title).ApplicantId(applicantId).ApproverId(approverId).MinAmount(minAmount).MaxAmount(maxAmount).PartnerId(partnerId).PartnerCode(partnerCode).PaymentMethod(paymentMethod).StartPaymentDate(startPaymentDate).EndPaymentDate(endPaymentDate).DocumentCode(documentCode).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentRequestsApi.GetPaymentRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentRequests`: PaymentRequestsIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentRequestsApi.GetPaymentRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **status** | **string** | &#39;申請ステータス(draft:下書き, in_progress:申請中, approved:承認済, rejected:却下, feedback:差戻し)、 取引ステータス(unsettled:支払待ち, settled:支払済み)&#39;&lt;br&gt; approver_id を指定した場合は無効です。  | 
 **startApplicationDate** | **string** | 申請日で絞込：開始日(yyyy-mm-dd) | 
 **endApplicationDate** | **string** | 申請日で絞込：終了日(yyyy-mm-dd) | 
 **startIssueDate** | **string** | 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endIssueDate** | **string** | 発生日で絞込：終了日(yyyy-mm-dd) | 
 **applicationNumber** | **int32** | 申請No. | 
 **title** | **string** | 申請タイトル | 
 **applicantId** | **int32** | 申請者のユーザーID | 
 **approverId** | **int32** | 承認者のユーザーID&lt;br&gt; &#39;approver_id を指定した場合は、 in_progress: 申請中 の支払依頼のみを返します。&#39;  | 
 **minAmount** | **int32** | 金額で絞込 (下限金額) | 
 **maxAmount** | **int32** | 金額で絞込 (上限金額) | 
 **partnerId** | **int32** | 取引先IDで絞込 | 
 **partnerCode** | **string** | 取引先コードで絞込 | 
 **paymentMethod** | **string** | 支払方法で絞込 (none: 指定なし, domestic_bank_transfer: 国内振込, abroad_bank_transfer: 国外振込, account_transfer: 口座振替, credit_card: クレジットカード) | 
 **startPaymentDate** | **string** | 支払期限で絞込：開始日(yyyy-mm-dd) | 
 **endPaymentDate** | **string** | 支払期限で絞込：終了日(yyyy-mm-dd) | 
 **documentCode** | **string** | 請求書番号で絞込 | 
 **offset** | **int64** | 取得レコードのオフセット (デフォルト: 0) | 
 **limit** | **int32** | 取得レコードの件数 (デフォルト: 50, 最小: 1, 最大: 500) | 

### Return type

[**PaymentRequestsIndexResponse**](PaymentRequestsIndexResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePaymentRequest

> PaymentRequestResponse UpdatePaymentRequest(ctx, id).PaymentRequestUpdateParams(paymentRequestUpdateParams).Execute()

支払依頼の更新



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | 支払依頼ID
    paymentRequestUpdateParams := *openapiclient.NewPaymentRequestUpdateParams(int32(1), int32(1), true, "2019-12-17", []openapiclient.PaymentRequestUpdateParamsPaymentRequestLinesInner{*openapiclient.NewPaymentRequestUpdateParamsPaymentRequestLinesInner(int64(30000))}, "仕入代金支払い") // PaymentRequestUpdateParams | 支払依頼の更新 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentRequestsApi.UpdatePaymentRequest(context.Background(), id).PaymentRequestUpdateParams(paymentRequestUpdateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentRequestsApi.UpdatePaymentRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePaymentRequest`: PaymentRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentRequestsApi.UpdatePaymentRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 支払依頼ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePaymentRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentRequestUpdateParams** | [**PaymentRequestUpdateParams**](PaymentRequestUpdateParams.md) | 支払依頼の更新 | 

### Return type

[**PaymentRequestResponse**](PaymentRequestResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePaymentRequestAction

> PaymentRequestResponse UpdatePaymentRequestAction(ctx, id).PaymentRequestActionCreateParams(paymentRequestActionCreateParams).Execute()

支払依頼の承認操作



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | 支払依頼ID
    paymentRequestActionCreateParams := *openapiclient.NewPaymentRequestActionCreateParams("approve", int32(1), int32(1), int32(1)) // PaymentRequestActionCreateParams | 支払依頼の承認操作

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentRequestsApi.UpdatePaymentRequestAction(context.Background(), id).PaymentRequestActionCreateParams(paymentRequestActionCreateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentRequestsApi.UpdatePaymentRequestAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePaymentRequestAction`: PaymentRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentRequestsApi.UpdatePaymentRequestAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 支払依頼ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePaymentRequestActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentRequestActionCreateParams** | [**PaymentRequestActionCreateParams**](PaymentRequestActionCreateParams.md) | 支払依頼の承認操作 | 

### Return type

[**PaymentRequestResponse**](PaymentRequestResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

