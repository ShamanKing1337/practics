# \PurchaseApiServiceApi

All URIs are relative to *https://localhost:2662*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PurchaseApiServiceCreateItem**](PurchaseApiServiceApi.md#PurchaseApiServiceCreateItem) | **Post** /v1/items | 
[**PurchaseApiServiceCreatePurchase**](PurchaseApiServiceApi.md#PurchaseApiServiceCreatePurchase) | **Post** /v1/purchase | 
[**PurchaseApiServiceCreateUser**](PurchaseApiServiceApi.md#PurchaseApiServiceCreateUser) | **Post** /v1/users | 
[**PurchaseApiServiceDeletePurchase**](PurchaseApiServiceApi.md#PurchaseApiServiceDeletePurchase) | **Delete** /v1/purchase/{id} | 
[**PurchaseApiServiceGetPurchases**](PurchaseApiServiceApi.md#PurchaseApiServiceGetPurchases) | **Get** /v1/purchase | 
[**PurchaseApiServiceUpdateUser**](PurchaseApiServiceApi.md#PurchaseApiServiceUpdateUser) | **Put** /v1/users/{id} | 


# **PurchaseApiServiceCreateItem**
> PracticsCreateItemResponse PurchaseApiServiceCreateItem(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PracticsCreateItemRequest**](PracticsCreateItemRequest.md)|  | 

### Return type

[**PracticsCreateItemResponse**](practicsCreateItemResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PurchaseApiServiceCreatePurchase**
> PracticsCreatePurchaseResponse PurchaseApiServiceCreatePurchase(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PracticsCreatePurchaseRequest**](PracticsCreatePurchaseRequest.md)|  | 

### Return type

[**PracticsCreatePurchaseResponse**](practicsCreatePurchaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PurchaseApiServiceCreateUser**
> PracticsCreateUserResponse PurchaseApiServiceCreateUser(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PracticsCreateUserRequest**](PracticsCreateUserRequest.md)|  | 

### Return type

[**PracticsCreateUserResponse**](practicsCreateUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PurchaseApiServiceDeletePurchase**
> interface{} PurchaseApiServiceDeletePurchase(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PurchaseApiServiceGetPurchases**
> PracticsGetPurchasesResponse PurchaseApiServiceGetPurchases(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PurchaseApiServiceApiPurchaseApiServiceGetPurchasesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PurchaseApiServiceApiPurchaseApiServiceGetPurchasesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **optional.String**|  | 
 **itemId** | **optional.String**|  | 

### Return type

[**PracticsGetPurchasesResponse**](practicsGetPurchasesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PurchaseApiServiceUpdateUser**
> interface{} PurchaseApiServiceUpdateUser(ctx, id, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **body** | [**PracticsUpdateUserRequest**](PracticsUpdateUserRequest.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

