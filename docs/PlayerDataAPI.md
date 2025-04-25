# \PlayerDataAPI

All URIs are relative to *https://api.hypixel.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2GuildGet**](PlayerDataAPI.md#V2GuildGet) | **Get** /v2/guild | Retrieve a Guild by a player, id, or name
[**V2PlayerGet**](PlayerDataAPI.md#V2PlayerGet) | **Get** /v2/player | Data of a specific player, including game stats
[**V2RecentgamesGet**](PlayerDataAPI.md#V2RecentgamesGet) | **Get** /v2/recentgames | The recently played games of a specific player
[**V2StatusGet**](PlayerDataAPI.md#V2StatusGet) | **Get** /v2/status | The current online status of a specific player



## V2GuildGet

> V2GuildGet200Response V2GuildGet(ctx).Id(id).Player(player).Name(name).Execute()

Retrieve a Guild by a player, id, or name

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
	id := "id_example" // string |  (optional)
	player := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	name := "name_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayerDataAPI.V2GuildGet(context.Background()).Id(id).Player(player).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayerDataAPI.V2GuildGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2GuildGet`: V2GuildGet200Response
	fmt.Fprintf(os.Stdout, "Response from `PlayerDataAPI.V2GuildGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2GuildGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **player** | **string** |  | 
 **name** | **string** |  | 

### Return type

[**V2GuildGet200Response**](V2GuildGet200Response.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2PlayerGet

> V2PlayerGet200Response V2PlayerGet(ctx).Uuid(uuid).Execute()

Data of a specific player, including game stats

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
	resp, r, err := apiClient.PlayerDataAPI.V2PlayerGet(context.Background()).Uuid(uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayerDataAPI.V2PlayerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2PlayerGet`: V2PlayerGet200Response
	fmt.Fprintf(os.Stdout, "Response from `PlayerDataAPI.V2PlayerGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2PlayerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **string** |  | 

### Return type

[**V2PlayerGet200Response**](V2PlayerGet200Response.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2RecentgamesGet

> V2RecentgamesGet200Response V2RecentgamesGet(ctx).Uuid(uuid).Execute()

The recently played games of a specific player

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
	resp, r, err := apiClient.PlayerDataAPI.V2RecentgamesGet(context.Background()).Uuid(uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayerDataAPI.V2RecentgamesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2RecentgamesGet`: V2RecentgamesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `PlayerDataAPI.V2RecentgamesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2RecentgamesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **string** |  | 

### Return type

[**V2RecentgamesGet200Response**](V2RecentgamesGet200Response.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2StatusGet

> V2StatusGet200Response V2StatusGet(ctx).Uuid(uuid).Execute()

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
	// response from `V2StatusGet`: V2StatusGet200Response
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

[**V2StatusGet200Response**](V2StatusGet200Response.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

