# \PartnersApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePartner**](PartnersApi.md#CreatePartner) | **Post** /api/1/partners | 取引先の作成
[**DestroyPartner**](PartnersApi.md#DestroyPartner) | **Delete** /api/1/partners/{id} | 取引先の削除
[**GetPartner**](PartnersApi.md#GetPartner) | **Get** /api/1/partners/{id} | 取引先の取得
[**GetPartners**](PartnersApi.md#GetPartners) | **Get** /api/1/partners | 取引先一覧の取得
[**UpdatePartner**](PartnersApi.md#UpdatePartner) | **Put** /api/1/partners/{id} | 取引先の更新
[**UpdatePartnerByCode**](PartnersApi.md#UpdatePartnerByCode) | **Put** /api/1/partners/code/{code} | 取引先の更新



## CreatePartner

> PartnerResponse CreatePartner(ctx).PartnerCreateParams(partnerCreateParams).Execute()

取引先の作成



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
    partnerCreateParams := *openapiclient.NewPartnerCreateParams(int64(1), "新しい取引先") // PartnerCreateParams | 取引先の作成

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnersApi.CreatePartner(context.Background()).PartnerCreateParams(partnerCreateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnersApi.CreatePartner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePartner`: PartnerResponse
    fmt.Fprintf(os.Stdout, "Response from `PartnersApi.CreatePartner`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePartnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **partnerCreateParams** | [**PartnerCreateParams**](PartnerCreateParams.md) | 取引先の作成 | 

### Return type

[**PartnerResponse**](PartnerResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyPartner

> DestroyPartner(ctx, id).CompanyId(companyId).Execute()

取引先の削除



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
    id := int64(56) // int64 | 取引先ID
    companyId := int64(56) // int64 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnersApi.DestroyPartner(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnersApi.DestroyPartner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 取引先ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyPartnerRequest struct via the builder pattern


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


## GetPartner

> PartnerResponse GetPartner(ctx, id).CompanyId(companyId).Execute()

取引先の取得



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
    id := int64(56) // int64 | 取引先ID
    companyId := int64(56) // int64 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnersApi.GetPartner(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnersApi.GetPartner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartner`: PartnerResponse
    fmt.Fprintf(os.Stdout, "Response from `PartnersApi.GetPartner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 取引先ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyId** | **int64** | 事業所ID | 

### Return type

[**PartnerResponse**](PartnerResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartners

> PartnersResponse GetPartners(ctx).CompanyId(companyId).StartUpdateDate(startUpdateDate).EndUpdateDate(endUpdateDate).Offset(offset).Limit(limit).Keyword(keyword).Execute()

取引先一覧の取得



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
    startUpdateDate := "startUpdateDate_example" // string | 更新日で絞り込み：開始日(yyyy-mm-dd) (optional)
    endUpdateDate := "endUpdateDate_example" // string | 更新日で絞り込み：終了日(yyyy-mm-dd) (optional)
    offset := int64(789) // int64 | 取得レコードのオフセット (デフォルト: 0) (optional)
    limit := int64(56) // int64 | 取得レコードの件数 (デフォルト: 50, 最小: 1, 最大: 3000) (optional)
    keyword := "keyword_example" // string | 検索キーワード：取引先名・正式名称・カナ名称に対するあいまい検索で一致、またはショートカットキー1・2のいずれかに完全一致 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnersApi.GetPartners(context.Background()).CompanyId(companyId).StartUpdateDate(startUpdateDate).EndUpdateDate(endUpdateDate).Offset(offset).Limit(limit).Keyword(keyword).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnersApi.GetPartners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartners`: PartnersResponse
    fmt.Fprintf(os.Stdout, "Response from `PartnersApi.GetPartners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPartnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int64** | 事業所ID | 
 **startUpdateDate** | **string** | 更新日で絞り込み：開始日(yyyy-mm-dd) | 
 **endUpdateDate** | **string** | 更新日で絞り込み：終了日(yyyy-mm-dd) | 
 **offset** | **int64** | 取得レコードのオフセット (デフォルト: 0) | 
 **limit** | **int64** | 取得レコードの件数 (デフォルト: 50, 最小: 1, 最大: 3000) | 
 **keyword** | **string** | 検索キーワード：取引先名・正式名称・カナ名称に対するあいまい検索で一致、またはショートカットキー1・2のいずれかに完全一致 | 

### Return type

[**PartnersResponse**](PartnersResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePartner

> PartnerResponse UpdatePartner(ctx, id).PartnerUpdateParams(partnerUpdateParams).Execute()

取引先の更新



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
    id := int64(56) // int64 | 取引先ID
    partnerUpdateParams := *openapiclient.NewPartnerUpdateParams(int64(1), "新しい取引先") // PartnerUpdateParams | 取引先の更新

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnersApi.UpdatePartner(context.Background(), id).PartnerUpdateParams(partnerUpdateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnersApi.UpdatePartner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePartner`: PartnerResponse
    fmt.Fprintf(os.Stdout, "Response from `PartnersApi.UpdatePartner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 取引先ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePartnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **partnerUpdateParams** | [**PartnerUpdateParams**](PartnerUpdateParams.md) | 取引先の更新 | 

### Return type

[**PartnerResponse**](PartnerResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePartnerByCode

> PartnerResponse UpdatePartnerByCode(ctx, code).PartnerUpdateParams(partnerUpdateParams).Execute()

取引先の更新



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
    code := "code_example" // string | 取引先コード
    partnerUpdateParams := *openapiclient.NewPartnerUpdateParams(int64(1), "新しい取引先") // PartnerUpdateParams | 取引先の更新

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnersApi.UpdatePartnerByCode(context.Background(), code).PartnerUpdateParams(partnerUpdateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnersApi.UpdatePartnerByCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePartnerByCode`: PartnerResponse
    fmt.Fprintf(os.Stdout, "Response from `PartnersApi.UpdatePartnerByCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | 取引先コード | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePartnerByCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **partnerUpdateParams** | [**PartnerUpdateParams**](PartnerUpdateParams.md) | 取引先の更新 | 

### Return type

[**PartnerResponse**](PartnerResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

