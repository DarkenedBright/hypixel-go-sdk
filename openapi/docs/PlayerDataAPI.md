# \PlayerDataAPI

All URIs are relative to *https://api.hypixel.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2StatusGet**](PlayerDataAPI.md#V2StatusGet) | **Get** /v2/status | The current online status of a specific player



## V2StatusGet

> StatusGet V2StatusGet(ctx).Uuid(uuid).Execute()

The current online status of a specific player

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
	uuid := "uuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayerDataAPI.V2StatusGet(context.Background()).Uuid(uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayerDataAPI.V2StatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2StatusGet`: StatusGet
	fmt.Fprintf(os.Stdout, "Response from `PlayerDataAPI.V2StatusGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2StatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **string** |  | 

### Return type

[**StatusGet**](StatusGet.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

