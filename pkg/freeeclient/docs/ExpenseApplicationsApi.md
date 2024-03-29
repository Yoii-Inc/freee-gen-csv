# \ExpenseApplicationsApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExpenseApplication**](ExpenseApplicationsApi.md#CreateExpenseApplication) | **Post** /api/1/expense_applications | 経費申請の作成
[**DestroyExpenseApplication**](ExpenseApplicationsApi.md#DestroyExpenseApplication) | **Delete** /api/1/expense_applications/{id} | 経費申請の削除
[**GetExpenseApplication**](ExpenseApplicationsApi.md#GetExpenseApplication) | **Get** /api/1/expense_applications/{id} | 経費申請詳細の取得
[**GetExpenseApplications**](ExpenseApplicationsApi.md#GetExpenseApplications) | **Get** /api/1/expense_applications | 経費申請一覧の取得
[**UpdateExpenseApplication**](ExpenseApplicationsApi.md#UpdateExpenseApplication) | **Put** /api/1/expense_applications/{id} | 経費申請の更新
[**UpdateExpenseApplicationAction**](ExpenseApplicationsApi.md#UpdateExpenseApplicationAction) | **Post** /api/1/expense_applications/{id}/actions | 経費申請の承認操作



## CreateExpenseApplication

> ExpenseApplicationResponse CreateExpenseApplication(ctx).ExpenseApplicationCreateParams(expenseApplicationCreateParams).Execute()

経費申請の作成



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
    expenseApplicationCreateParams := *openapiclient.NewExpenseApplicationCreateParams(int64(1), []openapiclient.ExpenseApplicationCreateParamsExpenseApplicationLinesInner{*openapiclient.NewExpenseApplicationCreateParamsExpenseApplicationLinesInner()}, "大阪出張") // ExpenseApplicationCreateParams | 経費申請の作成 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExpenseApplicationsApi.CreateExpenseApplication(context.Background()).ExpenseApplicationCreateParams(expenseApplicationCreateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpenseApplicationsApi.CreateExpenseApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateExpenseApplication`: ExpenseApplicationResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpenseApplicationsApi.CreateExpenseApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExpenseApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expenseApplicationCreateParams** | [**ExpenseApplicationCreateParams**](ExpenseApplicationCreateParams.md) | 経費申請の作成 | 

### Return type

[**ExpenseApplicationResponse**](ExpenseApplicationResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyExpenseApplication

> DestroyExpenseApplication(ctx, id).CompanyId(companyId).Execute()

経費申請の削除



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
    id := int64(56) // int64 | 経費申請ID
    companyId := int64(56) // int64 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExpenseApplicationsApi.DestroyExpenseApplication(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpenseApplicationsApi.DestroyExpenseApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 経費申請ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyExpenseApplicationRequest struct via the builder pattern


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


## GetExpenseApplication

> ExpenseApplicationResponse GetExpenseApplication(ctx, id).CompanyId(companyId).Execute()

経費申請詳細の取得



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
    id := int64(56) // int64 | 経費申請ID
    companyId := int64(56) // int64 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExpenseApplicationsApi.GetExpenseApplication(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpenseApplicationsApi.GetExpenseApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExpenseApplication`: ExpenseApplicationResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpenseApplicationsApi.GetExpenseApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 経費申請ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExpenseApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyId** | **int64** | 事業所ID | 

### Return type

[**ExpenseApplicationResponse**](ExpenseApplicationResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExpenseApplications

> ExpenseApplicationsIndexResponse GetExpenseApplications(ctx).CompanyId(companyId).Status(status).PayrollAttached(payrollAttached).StartTransactionDate(startTransactionDate).EndTransactionDate(endTransactionDate).ApplicationNumber(applicationNumber).Title(title).StartIssueDate(startIssueDate).EndIssueDate(endIssueDate).ApplicantId(applicantId).ApproverId(approverId).MinAmount(minAmount).MaxAmount(maxAmount).Offset(offset).Limit(limit).Execute()

経費申請一覧の取得



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
    status := "status_example" // string | 申請ステータス(draft:下書き, in_progress:申請中, approved:承認済, rejected:却下, feedback:差戻し)、 取引ステータス(unsettled:清算待ち, settled:精算済み) (optional)
    payrollAttached := true // bool | true:給与連携あり、false:給与連携なし、未指定時:絞り込みなし (optional)
    startTransactionDate := "2019-12-17" // string | 発生日(経費申請項目の日付)で絞込：開始日(yyyy-mm-dd) (optional)
    endTransactionDate := "2019-12-17" // string | 発生日(経費申請項目の日付)で絞込：終了日(yyyy-mm-dd) (optional)
    applicationNumber := int64(2) // int64 | 申請No. (optional)
    title := "大阪出張" // string | 申請タイトル (optional)
    startIssueDate := "2019-12-17" // string | 申請日で絞込：開始日(yyyy-mm-dd) (optional)
    endIssueDate := "2019-12-17" // string | 申請日で絞込：終了日(yyyy-mm-dd) (optional)
    applicantId := int64(1) // int64 | 申請者のユーザーID (optional)
    approverId := int64(1) // int64 | 承認者のユーザーID (optional)
    minAmount := int64(5000) // int64 | 金額で絞込 (下限金額) (optional)
    maxAmount := int64(10000) // int64 | 金額で絞込 (上限金額) (optional)
    offset := int64(789) // int64 | 取得レコードのオフセット (デフォルト: 0) (optional)
    limit := int64(56) // int64 | 取得レコードの件数 (デフォルト: 50, 最小: 1, 最大: 500) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExpenseApplicationsApi.GetExpenseApplications(context.Background()).CompanyId(companyId).Status(status).PayrollAttached(payrollAttached).StartTransactionDate(startTransactionDate).EndTransactionDate(endTransactionDate).ApplicationNumber(applicationNumber).Title(title).StartIssueDate(startIssueDate).EndIssueDate(endIssueDate).ApplicantId(applicantId).ApproverId(approverId).MinAmount(minAmount).MaxAmount(maxAmount).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpenseApplicationsApi.GetExpenseApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExpenseApplications`: ExpenseApplicationsIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpenseApplicationsApi.GetExpenseApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExpenseApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int64** | 事業所ID | 
 **status** | **string** | 申請ステータス(draft:下書き, in_progress:申請中, approved:承認済, rejected:却下, feedback:差戻し)、 取引ステータス(unsettled:清算待ち, settled:精算済み) | 
 **payrollAttached** | **bool** | true:給与連携あり、false:給与連携なし、未指定時:絞り込みなし | 
 **startTransactionDate** | **string** | 発生日(経費申請項目の日付)で絞込：開始日(yyyy-mm-dd) | 
 **endTransactionDate** | **string** | 発生日(経費申請項目の日付)で絞込：終了日(yyyy-mm-dd) | 
 **applicationNumber** | **int64** | 申請No. | 
 **title** | **string** | 申請タイトル | 
 **startIssueDate** | **string** | 申請日で絞込：開始日(yyyy-mm-dd) | 
 **endIssueDate** | **string** | 申請日で絞込：終了日(yyyy-mm-dd) | 
 **applicantId** | **int64** | 申請者のユーザーID | 
 **approverId** | **int64** | 承認者のユーザーID | 
 **minAmount** | **int64** | 金額で絞込 (下限金額) | 
 **maxAmount** | **int64** | 金額で絞込 (上限金額) | 
 **offset** | **int64** | 取得レコードのオフセット (デフォルト: 0) | 
 **limit** | **int64** | 取得レコードの件数 (デフォルト: 50, 最小: 1, 最大: 500) | 

### Return type

[**ExpenseApplicationsIndexResponse**](ExpenseApplicationsIndexResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExpenseApplication

> ExpenseApplicationResponse UpdateExpenseApplication(ctx, id).ExpenseApplicationUpdateParams(expenseApplicationUpdateParams).Execute()

経費申請の更新



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
    id := int64(56) // int64 | 経費申請ID
    expenseApplicationUpdateParams := *openapiclient.NewExpenseApplicationUpdateParams(int64(1), []openapiclient.ExpenseApplicationUpdateParamsExpenseApplicationLinesInner{*openapiclient.NewExpenseApplicationUpdateParamsExpenseApplicationLinesInner()}, "大阪出張") // ExpenseApplicationUpdateParams | 経費申請の更新 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExpenseApplicationsApi.UpdateExpenseApplication(context.Background(), id).ExpenseApplicationUpdateParams(expenseApplicationUpdateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpenseApplicationsApi.UpdateExpenseApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateExpenseApplication`: ExpenseApplicationResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpenseApplicationsApi.UpdateExpenseApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 経費申請ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExpenseApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expenseApplicationUpdateParams** | [**ExpenseApplicationUpdateParams**](ExpenseApplicationUpdateParams.md) | 経費申請の更新 | 

### Return type

[**ExpenseApplicationResponse**](ExpenseApplicationResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExpenseApplicationAction

> ExpenseApplicationResponse UpdateExpenseApplicationAction(ctx, id).ExpenseApplicationActionCreateParams(expenseApplicationActionCreateParams).Execute()

経費申請の承認操作



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
    id := int64(56) // int64 | 経費申請ID
    expenseApplicationActionCreateParams := *openapiclient.NewExpenseApplicationActionCreateParams("approve", int64(1), int64(1), int64(1)) // ExpenseApplicationActionCreateParams | 経費申請の承認操作

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExpenseApplicationsApi.UpdateExpenseApplicationAction(context.Background(), id).ExpenseApplicationActionCreateParams(expenseApplicationActionCreateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpenseApplicationsApi.UpdateExpenseApplicationAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateExpenseApplicationAction`: ExpenseApplicationResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpenseApplicationsApi.UpdateExpenseApplicationAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 経費申請ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExpenseApplicationActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expenseApplicationActionCreateParams** | [**ExpenseApplicationActionCreateParams**](ExpenseApplicationActionCreateParams.md) | 経費申請の承認操作 | 

### Return type

[**ExpenseApplicationResponse**](ExpenseApplicationResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

