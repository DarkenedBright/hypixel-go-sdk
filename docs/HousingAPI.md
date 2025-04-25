# \HousingAPI

All URIs are relative to *https://api.hypixel.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2HousingActiveGet**](HousingAPI.md#V2HousingActiveGet) | **Get** /v2/housing/active | The currently active public houses.
[**V2HousingHouseGet**](HousingAPI.md#V2HousingHouseGet) | **Get** /v2/housing/house | Information about a specific house.
[**V2HousingHousesGet**](HousingAPI.md#V2HousingHousesGet) | **Get** /v2/housing/houses | The public houses for a specific player.



## V2HousingActiveGet

> []HousingHouse V2HousingActiveGet(ctx).Execute()

The currently active public houses.



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
	resp, r, err := apiClient.HousingAPI.V2HousingActiveGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HousingAPI.V2HousingActiveGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2HousingActiveGet`: []HousingHouse
	fmt.Fprintf(os.Stdout, "Response from `HousingAPI.V2HousingActiveGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2HousingActiveGetRequest struct via the builder pattern


### Return type

[**[]HousingHouse**](HousingHouse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2HousingHouseGet

> HousingHouse V2HousingHouseGet(ctx).House(house).Execute()

Information about a specific house.



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
	house := "house_example" // string | The UUID of the house to get information about.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HousingAPI.V2HousingHouseGet(context.Background()).House(house).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HousingAPI.V2HousingHouseGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2HousingHouseGet`: HousingHouse
	fmt.Fprintf(os.Stdout, "Response from `HousingAPI.V2HousingHouseGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2HousingHouseGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **house** | **string** | The UUID of the house to get information about. | 

### Return type

[**HousingHouse**](HousingHouse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2HousingHousesGet

> []HousingHouse V2HousingHousesGet(ctx).Player(player).Execute()

The public houses for a specific player.



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
	player := "player_example" // string | The UUID of the player to get houses for. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HousingAPI.V2HousingHousesGet(context.Background()).Player(player).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HousingAPI.V2HousingHousesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2HousingHousesGet`: []HousingHouse
	fmt.Fprintf(os.Stdout, "Response from `HousingAPI.V2HousingHousesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2HousingHousesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **player** | **string** | The UUID of the player to get houses for. | 

### Return type

[**[]HousingHouse**](HousingHouse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

