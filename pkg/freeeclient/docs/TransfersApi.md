# \TransfersApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTransfer**](TransfersApi.md#CreateTransfer) | **Post** /api/1/transfers | 取引（振替）の作成
[**DestroyTransfer**](TransfersApi.md#DestroyTransfer) | **Delete** /api/1/transfers/{id} | 取引（振替）の削除する
[**GetTransfer**](TransfersApi.md#GetTransfer) | **Get** /api/1/transfers/{id} | 取引（振替）の取得
[**GetTransfers**](TransfersApi.md#GetTransfers) | **Get** /api/1/transfers | 取引（振替）一覧の取得
[**UpdateTransfer**](TransfersApi.md#UpdateTransfer) | **Put** /api/1/transfers/{id} | 取引（振替）の更新



## CreateTransfer

> TransferResponse CreateTransfer(ctx).TransferParams(transferParams).Execute()

取引（振替）の作成



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
    transferParams := *openapiclient.NewTransferParams(int64(5000), int64(1), "2019-12-17", int64(1), "credit_card", int64(1), "bank_account") // TransferParams | 取引（振替）の作成 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransfersApi.CreateTransfer(context.Background()).TransferParams(transferParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransfersApi.CreateTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTransfer`: TransferResponse
    fmt.Fprintf(os.Stdout, "Response from `TransfersApi.CreateTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferParams** | [**TransferParams**](TransferParams.md) | 取引（振替）の作成 | 

### Return type

[**TransferResponse**](TransferResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyTransfer

> DestroyTransfer(ctx, id).CompanyId(companyId).Execute()

取引（振替）の削除する



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
    id := int64(56) // int64 | 取引(振替)ID
    companyId := int64(56) // int64 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransfersApi.DestroyTransfer(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransfersApi.DestroyTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 取引(振替)ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyTransferRequest struct via the builder pattern


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


## GetTransfer

> TransferResponse GetTransfer(ctx, id).CompanyId(companyId).Execute()

取引（振替）の取得



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
    id := int64(56) // int64 | 取引(振替)ID
    companyId := int64(56) // int64 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransfersApi.GetTransfer(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransfersApi.GetTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransfer`: TransferResponse
    fmt.Fprintf(os.Stdout, "Response from `TransfersApi.GetTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 取引(振替)ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyId** | **int64** | 事業所ID | 

### Return type

[**TransferResponse**](TransferResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransfers

> GetTransfers200Response GetTransfers(ctx).CompanyId(companyId).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()

取引（振替）一覧の取得



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
    companyId := int64(56) // int64 | 事業所ID
    startDate := "startDate_example" // string | 振替日で絞込：開始日 (yyyy-mm-dd) (optional)
    endDate := "endDate_example" // string | 振替日で絞込：終了日 (yyyy-mm-dd) (optional)
    offset := int64(789) // int64 | 取得レコードのオフセット (デフォルト: 0) (optional)
    limit := int64(56) // int64 | 取得レコードの件数 (デフォルト: 20, 最小: 1, 最大: 100)  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransfersApi.GetTransfers(context.Background()).CompanyId(companyId).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransfersApi.GetTransfers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransfers`: GetTransfers200Response
    fmt.Fprintf(os.Stdout, "Response from `TransfersApi.GetTransfers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransfersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int64** | 事業所ID | 
 **startDate** | **string** | 振替日で絞込：開始日 (yyyy-mm-dd) | 
 **endDate** | **string** | 振替日で絞込：終了日 (yyyy-mm-dd) | 
 **offset** | **int64** | 取得レコードのオフセット (デフォルト: 0) | 
 **limit** | **int64** | 取得レコードの件数 (デフォルト: 20, 最小: 1, 最大: 100)  | 

### Return type

[**GetTransfers200Response**](GetTransfers200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTransfer

> TransferResponse UpdateTransfer(ctx, id).TransferParams(transferParams).Execute()

取引（振替）の更新



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
    id := int64(56) // int64 | 取引(振替)ID
    transferParams := *openapiclient.NewTransferParams(int64(5000), int64(1), "2019-12-17", int64(1), "credit_card", int64(1), "bank_account") // TransferParams | 取引（振替）の更新

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransfersApi.UpdateTransfer(context.Background(), id).TransferParams(transferParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransfersApi.UpdateTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTransfer`: TransferResponse
    fmt.Fprintf(os.Stdout, "Response from `TransfersApi.UpdateTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 取引(振替)ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transferParams** | [**TransferParams**](TransferParams.md) | 取引（振替）の更新 | 

### Return type

[**TransferResponse**](TransferResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

