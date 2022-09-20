# \TaxesApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTaxCode**](TaxesApi.md#GetTaxCode) | **Get** /api/1/taxes/codes/{code} | 税区分コードの取得
[**GetTaxCodes**](TaxesApi.md#GetTaxCodes) | **Get** /api/1/taxes/codes | 税区分コード一覧の取得
[**GetTaxesCompanies**](TaxesApi.md#GetTaxesCompanies) | **Get** /api/1/taxes/companies/{company_id} | 税区分コード詳細一覧の取得



## GetTaxCode

> TaxResponse GetTaxCode(ctx, code).Execute()

税区分コードの取得



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
    code := int64(56) // int64 | 税区分コード

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxesApi.GetTaxCode(context.Background(), code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxesApi.GetTaxCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaxCode`: TaxResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxesApi.GetTaxCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **int64** | 税区分コード | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaxResponse**](TaxResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxCodes

> GetTaxCodes200Response GetTaxCodes(ctx).Execute()

税区分コード一覧の取得



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxesApi.GetTaxCodes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxesApi.GetTaxCodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaxCodes`: GetTaxCodes200Response
    fmt.Fprintf(os.Stdout, "Response from `TaxesApi.GetTaxCodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxCodesRequest struct via the builder pattern


### Return type

[**GetTaxCodes200Response**](GetTaxCodes200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxesCompanies

> GetTaxesCompanies200Response GetTaxesCompanies(ctx, companyId).Execute()

税区分コード詳細一覧の取得

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxesApi.GetTaxesCompanies(context.Background(), companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxesApi.GetTaxesCompanies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaxesCompanies`: GetTaxesCompanies200Response
    fmt.Fprintf(os.Stdout, "Response from `TaxesApi.GetTaxesCompanies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int64** | 事業所ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxesCompaniesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetTaxesCompanies200Response**](GetTaxesCompanies200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

