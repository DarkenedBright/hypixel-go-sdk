# \SkyBlockAPI

All URIs are relative to *https://api.hypixel.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2SkyblockAuctionsEndedGet**](SkyBlockAPI.md#V2SkyblockAuctionsEndedGet) | **Get** /v2/skyblock/auctions_ended | Recently ended auctions
[**V2SkyblockBazaarGet**](SkyBlockAPI.md#V2SkyblockBazaarGet) | **Get** /v2/skyblock/bazaar | Bazaar
[**V2SkyblockBingoGet**](SkyBlockAPI.md#V2SkyblockBingoGet) | **Get** /v2/skyblock/bingo | Bingo data by player
[**V2SkyblockFiresalesGet**](SkyBlockAPI.md#V2SkyblockFiresalesGet) | **Get** /v2/skyblock/firesales | Active/Upcoming Fire Sales
[**V2SkyblockGardenGet**](SkyBlockAPI.md#V2SkyblockGardenGet) | **Get** /v2/skyblock/garden | Garden data by profile ID
[**V2SkyblockMuseumGet**](SkyBlockAPI.md#V2SkyblockMuseumGet) | **Get** /v2/skyblock/museum | Museum data by profile ID
[**V2SkyblockNewsGet**](SkyBlockAPI.md#V2SkyblockNewsGet) | **Get** /v2/skyblock/news | 
[**V2SkyblockProfileGet**](SkyBlockAPI.md#V2SkyblockProfileGet) | **Get** /v2/skyblock/profile | 
[**V2SkyblockProfilesGet**](SkyBlockAPI.md#V2SkyblockProfilesGet) | **Get** /v2/skyblock/profiles | 



## V2SkyblockAuctionsEndedGet

> SkyBlockAuctionsEnded V2SkyblockAuctionsEndedGet(ctx).Execute()

Recently ended auctions



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
	resp, r, err := apiClient.SkyBlockAPI.V2SkyblockAuctionsEndedGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkyBlockAPI.V2SkyblockAuctionsEndedGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2SkyblockAuctionsEndedGet`: SkyBlockAuctionsEnded
	fmt.Fprintf(os.Stdout, "Response from `SkyBlockAPI.V2SkyblockAuctionsEndedGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2SkyblockAuctionsEndedGetRequest struct via the builder pattern


### Return type

[**SkyBlockAuctionsEnded**](SkyBlockAuctionsEnded.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2SkyblockBazaarGet

> SkyBlockBazaar V2SkyblockBazaarGet(ctx).Execute()

Bazaar

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
	resp, r, err := apiClient.SkyBlockAPI.V2SkyblockBazaarGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkyBlockAPI.V2SkyblockBazaarGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2SkyblockBazaarGet`: SkyBlockBazaar
	fmt.Fprintf(os.Stdout, "Response from `SkyBlockAPI.V2SkyblockBazaarGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2SkyblockBazaarGetRequest struct via the builder pattern


### Return type

[**SkyBlockBazaar**](SkyBlockBazaar.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2SkyblockBingoGet

> SkyBlockBingo V2SkyblockBingoGet(ctx).Uuid(uuid).Execute()

Bingo data by player



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
	resp, r, err := apiClient.SkyBlockAPI.V2SkyblockBingoGet(context.Background()).Uuid(uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkyBlockAPI.V2SkyblockBingoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2SkyblockBingoGet`: SkyBlockBingo
	fmt.Fprintf(os.Stdout, "Response from `SkyBlockAPI.V2SkyblockBingoGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2SkyblockBingoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **string** |  | 

### Return type

[**SkyBlockBingo**](SkyBlockBingo.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2SkyblockFiresalesGet

> SkyBlockFireSales V2SkyblockFiresalesGet(ctx).Execute()

Active/Upcoming Fire Sales



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
	resp, r, err := apiClient.SkyBlockAPI.V2SkyblockFiresalesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkyBlockAPI.V2SkyblockFiresalesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2SkyblockFiresalesGet`: SkyBlockFireSales
	fmt.Fprintf(os.Stdout, "Response from `SkyBlockAPI.V2SkyblockFiresalesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2SkyblockFiresalesGetRequest struct via the builder pattern


### Return type

[**SkyBlockFireSales**](SkyBlockFireSales.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2SkyblockGardenGet

> SkyBlockGardenGet V2SkyblockGardenGet(ctx).Profile(profile).Execute()

Garden data by profile ID



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
	resp, r, err := apiClient.SkyBlockAPI.V2SkyblockGardenGet(context.Background()).Profile(profile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkyBlockAPI.V2SkyblockGardenGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2SkyblockGardenGet`: SkyBlockGardenGet
	fmt.Fprintf(os.Stdout, "Response from `SkyBlockAPI.V2SkyblockGardenGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2SkyblockGardenGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profile** | **string** |  | 

### Return type

[**SkyBlockGardenGet**](SkyBlockGardenGet.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2SkyblockMuseumGet

> SkyBlockMuseum V2SkyblockMuseumGet(ctx).Profile(profile).Execute()

Museum data by profile ID



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
	resp, r, err := apiClient.SkyBlockAPI.V2SkyblockMuseumGet(context.Background()).Profile(profile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkyBlockAPI.V2SkyblockMuseumGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2SkyblockMuseumGet`: SkyBlockMuseum
	fmt.Fprintf(os.Stdout, "Response from `SkyBlockAPI.V2SkyblockMuseumGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2SkyblockMuseumGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profile** | **string** |  | 

### Return type

[**SkyBlockMuseum**](SkyBlockMuseum.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2SkyblockNewsGet

> SkyBlockNews V2SkyblockNewsGet(ctx).Execute()



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
	resp, r, err := apiClient.SkyBlockAPI.V2SkyblockNewsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkyBlockAPI.V2SkyblockNewsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2SkyblockNewsGet`: SkyBlockNews
	fmt.Fprintf(os.Stdout, "Response from `SkyBlockAPI.V2SkyblockNewsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2SkyblockNewsGetRequest struct via the builder pattern


### Return type

[**SkyBlockNews**](SkyBlockNews.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2SkyblockProfileGet

> SkyBlockProfileGet V2SkyblockProfileGet(ctx).Profile(profile).Execute()



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
	// response from `V2SkyblockProfileGet`: SkyBlockProfileGet
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

[**SkyBlockProfileGet**](SkyBlockProfileGet.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2SkyblockProfilesGet

> SkyBlockProfilesGet V2SkyblockProfilesGet(ctx).Uuid(uuid).Execute()



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
	// response from `V2SkyblockProfilesGet`: SkyBlockProfilesGet
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

[**SkyBlockProfilesGet**](SkyBlockProfilesGet.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

