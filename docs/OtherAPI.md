# \OtherAPI

All URIs are relative to *https://api.hypixel.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2BoostersGet**](OtherAPI.md#V2BoostersGet) | **Get** /v2/boosters | Active Network Boosters
[**V2CountsGet**](OtherAPI.md#V2CountsGet) | **Get** /v2/counts | Current Player Counts
[**V2LeaderboardsGet**](OtherAPI.md#V2LeaderboardsGet) | **Get** /v2/leaderboards | Current Leaderboards
[**V2PunishmentstatsGet**](OtherAPI.md#V2PunishmentstatsGet) | **Get** /v2/punishmentstats | Punishment Statistics



## V2BoostersGet

> V2BoostersGet200Response V2BoostersGet(ctx).Execute()

Active Network Boosters

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OtherAPI.V2BoostersGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OtherAPI.V2BoostersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2BoostersGet`: V2BoostersGet200Response
	fmt.Fprintf(os.Stdout, "Response from `OtherAPI.V2BoostersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2BoostersGetRequest struct via the builder pattern


### Return type

[**V2BoostersGet200Response**](V2BoostersGet200Response.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2CountsGet

> V2CountsGet200Response V2CountsGet(ctx).Execute()

Current Player Counts

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OtherAPI.V2CountsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OtherAPI.V2CountsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2CountsGet`: V2CountsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `OtherAPI.V2CountsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2CountsGetRequest struct via the builder pattern


### Return type

[**V2CountsGet200Response**](V2CountsGet200Response.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2LeaderboardsGet

> V2LeaderboardsGet200Response V2LeaderboardsGet(ctx).Execute()

Current Leaderboards

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OtherAPI.V2LeaderboardsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OtherAPI.V2LeaderboardsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2LeaderboardsGet`: V2LeaderboardsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `OtherAPI.V2LeaderboardsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2LeaderboardsGetRequest struct via the builder pattern


### Return type

[**V2LeaderboardsGet200Response**](V2LeaderboardsGet200Response.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2PunishmentstatsGet

> V2PunishmentstatsGet200Response V2PunishmentstatsGet(ctx).Execute()

Punishment Statistics

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OtherAPI.V2PunishmentstatsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OtherAPI.V2PunishmentstatsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2PunishmentstatsGet`: V2PunishmentstatsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `OtherAPI.V2PunishmentstatsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2PunishmentstatsGetRequest struct via the builder pattern


### Return type

[**V2PunishmentstatsGet200Response**](V2PunishmentstatsGet200Response.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

