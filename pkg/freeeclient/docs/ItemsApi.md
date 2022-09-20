# \ItemsApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateItem**](ItemsApi.md#CreateItem) | **Post** /api/1/items | 品目の作成
[**DestroyItem**](ItemsApi.md#DestroyItem) | **Delete** /api/1/items/{id} | 品目の削除
[**GetItem**](ItemsApi.md#GetItem) | **Get** /api/1/items/{id} | 品目の取得
[**GetItems**](ItemsApi.md#GetItems) | **Get** /api/1/items | 品目一覧の取得
[**UpdateItem**](ItemsApi.md#UpdateItem) | **Put** /api/1/items/{id} | 品目の更新



## CreateItem

> ItemResponse CreateItem(ctx).ItemParams(itemParams).Execute()

品目の作成



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
    itemParams := *openapiclient.NewItemParams(int32(1), "新しい品目") // ItemParams | 品目の作成 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemsApi.CreateItem(context.Background()).ItemParams(itemParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemsApi.CreateItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateItem`: ItemResponse
    fmt.Fprintf(os.Stdout, "Response from `ItemsApi.CreateItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemParams** | [**ItemParams**](ItemParams.md) | 品目の作成 | 

### Return type

[**ItemResponse**](ItemResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyItem

> DestroyItem(ctx, id).CompanyId(companyId).Execute()

品目の削除



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
    id := int32(56) // int32 | 品目ID
    companyId := int32(56) // int32 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemsApi.DestroyItem(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemsApi.DestroyItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 品目ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyItemRequest struct via the builder pattern


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


## GetItem

> ItemResponse GetItem(ctx, id).CompanyId(companyId).Execute()

品目の取得



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
    id := int32(56) // int32 | 品目ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemsApi.GetItem(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemsApi.GetItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItem`: ItemResponse
    fmt.Fprintf(os.Stdout, "Response from `ItemsApi.GetItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 品目ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 


### Return type

[**ItemResponse**](ItemResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItems

> GetItems200Response GetItems(ctx).CompanyId(companyId).StartUpdateDate(startUpdateDate).EndUpdateDate(endUpdateDate).Offset(offset).Limit(limit).Execute()

品目一覧の取得



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
    startUpdateDate := "startUpdateDate_example" // string | 更新日で絞り込み：開始日(yyyy-mm-dd) (optional)
    endUpdateDate := "endUpdateDate_example" // string | 更新日で絞り込み：終了日(yyyy-mm-dd) (optional)
    offset := int32(56) // int32 | 取得レコードのオフセット (デフォルト: 0) (optional)
    limit := int32(56) // int32 | 取得レコードの件数 (デフォルト: 50, 最小: 1, 最大: 3000) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemsApi.GetItems(context.Background()).CompanyId(companyId).StartUpdateDate(startUpdateDate).EndUpdateDate(endUpdateDate).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemsApi.GetItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItems`: GetItems200Response
    fmt.Fprintf(os.Stdout, "Response from `ItemsApi.GetItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **startUpdateDate** | **string** | 更新日で絞り込み：開始日(yyyy-mm-dd) | 
 **endUpdateDate** | **string** | 更新日で絞り込み：終了日(yyyy-mm-dd) | 
 **offset** | **int32** | 取得レコードのオフセット (デフォルト: 0) | 
 **limit** | **int32** | 取得レコードの件数 (デフォルト: 50, 最小: 1, 最大: 3000) | 

### Return type

[**GetItems200Response**](GetItems200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateItem

> ItemResponse UpdateItem(ctx, id).ItemParams(itemParams).Execute()

品目の更新



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
    id := int32(56) // int32 | 品目ID
    itemParams := *openapiclient.NewItemParams(int32(1), "新しい品目") // ItemParams | 品目の更新 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemsApi.UpdateItem(context.Background(), id).ItemParams(itemParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemsApi.UpdateItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateItem`: ItemResponse
    fmt.Fprintf(os.Stdout, "Response from `ItemsApi.UpdateItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 品目ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **itemParams** | [**ItemParams**](ItemParams.md) | 品目の更新 | 

### Return type

[**ItemResponse**](ItemResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

