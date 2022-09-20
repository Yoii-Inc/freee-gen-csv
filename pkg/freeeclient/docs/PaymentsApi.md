# \PaymentsApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDealPayment**](PaymentsApi.md#CreateDealPayment) | **Post** /api/1/deals/{id}/payments | 取引（収入／支出）の支払行作成
[**DestroyDealPayment**](PaymentsApi.md#DestroyDealPayment) | **Delete** /api/1/deals/{id}/payments/{payment_id} | 取引（収入／支出）の支払行削除
[**UpdateDealPayment**](PaymentsApi.md#UpdateDealPayment) | **Put** /api/1/deals/{id}/payments/{payment_id} | 取引（収入／支出）の支払行更新



## CreateDealPayment

> DealResponse CreateDealPayment(ctx, id).PaymentParams(paymentParams).Execute()

取引（収入／支出）の支払行作成



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
    id := int64(56) // int64 | 取引ID
    paymentParams := *openapiclient.NewPaymentParams(int64(10000), int64(1), "2019-12-17", int64(1), "bank_account") // PaymentParams | 取引（収入／支出）の支払行作成

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsApi.CreateDealPayment(context.Background(), id).PaymentParams(paymentParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.CreateDealPayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDealPayment`: DealResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.CreateDealPayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 取引ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDealPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentParams** | [**PaymentParams**](PaymentParams.md) | 取引（収入／支出）の支払行作成 | 

### Return type

[**DealResponse**](DealResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyDealPayment

> DestroyDealPayment(ctx, id, paymentId).CompanyId(companyId).Execute()

取引（収入／支出）の支払行削除



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
    id := int64(56) // int64 | 取引ID
    paymentId := int64(789) // int64 | 決済ID
    companyId := int64(56) // int64 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsApi.DestroyDealPayment(context.Background(), id, paymentId).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.DestroyDealPayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 取引ID | 
**paymentId** | **int64** | 決済ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyDealPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **companyId** | **int64** | 事業所ID | 

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


## UpdateDealPayment

> DealResponse UpdateDealPayment(ctx, id, paymentId).PaymentParams(paymentParams).Execute()

取引（収入／支出）の支払行更新



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
    id := int64(56) // int64 | 取引ID
    paymentId := int64(789) // int64 | 決済ID
    paymentParams := *openapiclient.NewPaymentParams(int64(10000), int64(1), "2019-12-17", int64(1), "bank_account") // PaymentParams | 取引（収入／支出）の支払行更新

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsApi.UpdateDealPayment(context.Background(), id, paymentId).PaymentParams(paymentParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.UpdateDealPayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDealPayment`: DealResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.UpdateDealPayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 取引ID | 
**paymentId** | **int64** | 決済ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDealPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **paymentParams** | [**PaymentParams**](PaymentParams.md) | 取引（収入／支出）の支払行更新 | 

### Return type

[**DealResponse**](DealResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

