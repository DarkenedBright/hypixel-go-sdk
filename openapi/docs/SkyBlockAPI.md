# \SkyBlockAPI

All URIs are relative to *https://api.hypixel.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2SkyblockProfileGet**](SkyBlockAPI.md#V2SkyblockProfileGet) | **Get** /v2/skyblock/profile | Profile by UUID
[**V2SkyblockProfilesGet**](SkyBlockAPI.md#V2SkyblockProfilesGet) | **Get** /v2/skyblock/profiles | Profiles by player



## V2SkyblockProfileGet

> V2SkyblockProfileGet200Response V2SkyblockProfileGet(ctx).Profile(profile).Execute()

Profile by UUID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/DarkenedBright/hypixel-go-sdk"
)

func main() {
	profile := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkyBlockAPI.V2SkyblockProfileGet(context.Background()).Profile(profile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkyBlockAPI.V2SkyblockProfileGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2SkyblockProfileGet`: V2SkyblockProfileGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SkyBlockAPI.V2SkyblockProfileGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2SkyblockProfileGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profile** | **string** |  | 

### Return type

[**V2SkyblockProfileGet200Response**](V2SkyblockProfileGet200Response.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2SkyblockProfilesGet

> V2SkyblockProfilesGet200Response V2SkyblockProfilesGet(ctx).Uuid(uuid).Execute()

Profiles by player

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/DarkenedBright/hypixel-go-sdk"
)

func main() {
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkyBlockAPI.V2SkyblockProfilesGet(context.Background()).Uuid(uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkyBlockAPI.V2SkyblockProfilesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2SkyblockProfilesGet`: V2SkyblockProfilesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SkyBlockAPI.V2SkyblockProfilesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2SkyblockProfilesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **string** |  | 

### Return type

[**V2SkyblockProfilesGet200Response**](V2SkyblockProfilesGet200Response.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

