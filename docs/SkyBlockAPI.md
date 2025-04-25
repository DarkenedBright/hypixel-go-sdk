# \SkyBlockAPI

All URIs are relative to *https://api.hypixel.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2ResourcesSkyblockBingoGet**](SkyBlockAPI.md#V2ResourcesSkyblockBingoGet) | **Get** /v2/resources/skyblock/bingo | Current Bingo Event
[**V2ResourcesSkyblockCollectionsGet**](SkyBlockAPI.md#V2ResourcesSkyblockCollectionsGet) | **Get** /v2/resources/skyblock/collections | Collections
[**V2ResourcesSkyblockElectionGet**](SkyBlockAPI.md#V2ResourcesSkyblockElectionGet) | **Get** /v2/resources/skyblock/election | Election and Mayor
[**V2ResourcesSkyblockItemsGet**](SkyBlockAPI.md#V2ResourcesSkyblockItemsGet) | **Get** /v2/resources/skyblock/items | Items
[**V2ResourcesSkyblockSkillsGet**](SkyBlockAPI.md#V2ResourcesSkyblockSkillsGet) | **Get** /v2/resources/skyblock/skills | Skills
[**V2SkyblockBingoGet**](SkyBlockAPI.md#V2SkyblockBingoGet) | **Get** /v2/skyblock/bingo | Bingo data by player
[**V2SkyblockFiresalesGet**](SkyBlockAPI.md#V2SkyblockFiresalesGet) | **Get** /v2/skyblock/firesales | Active/Upcoming Fire Sales
[**V2SkyblockGardenGet**](SkyBlockAPI.md#V2SkyblockGardenGet) | **Get** /v2/skyblock/garden | Garden data by profile ID
[**V2SkyblockMuseumGet**](SkyBlockAPI.md#V2SkyblockMuseumGet) | **Get** /v2/skyblock/museum | Museum data by profile ID
[**V2SkyblockNewsGet**](SkyBlockAPI.md#V2SkyblockNewsGet) | **Get** /v2/skyblock/news | News
[**V2SkyblockProfileGet**](SkyBlockAPI.md#V2SkyblockProfileGet) | **Get** /v2/skyblock/profile | Profile by UUID
[**V2SkyblockProfilesGet**](SkyBlockAPI.md#V2SkyblockProfilesGet) | **Get** /v2/skyblock/profiles | Profiles by player



## V2ResourcesSkyblockBingoGet

> V2ResourcesSkyblockBingoGet200Response V2ResourcesSkyblockBingoGet(ctx).Execute()

Current Bingo Event



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkyBlockAPI.V2ResourcesSkyblockBingoGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkyBlockAPI.V2ResourcesSkyblockBingoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ResourcesSkyblockBingoGet`: V2ResourcesSkyblockBingoGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SkyBlockAPI.V2ResourcesSkyblockBingoGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2ResourcesSkyblockBingoGetRequest struct via the builder pattern


### Return type

[**V2ResourcesSkyblockBingoGet200Response**](V2ResourcesSkyblockBingoGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ResourcesSkyblockCollectionsGet

> V2ResourcesSkyblockCollectionsGet200Response V2ResourcesSkyblockCollectionsGet(ctx).Execute()

Collections



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkyBlockAPI.V2ResourcesSkyblockCollectionsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkyBlockAPI.V2ResourcesSkyblockCollectionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ResourcesSkyblockCollectionsGet`: V2ResourcesSkyblockCollectionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SkyBlockAPI.V2ResourcesSkyblockCollectionsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2ResourcesSkyblockCollectionsGetRequest struct via the builder pattern


### Return type

[**V2ResourcesSkyblockCollectionsGet200Response**](V2ResourcesSkyblockCollectionsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ResourcesSkyblockElectionGet

> V2ResourcesSkyblockElectionGet200Response V2ResourcesSkyblockElectionGet(ctx).Execute()

Election and Mayor



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkyBlockAPI.V2ResourcesSkyblockElectionGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkyBlockAPI.V2ResourcesSkyblockElectionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ResourcesSkyblockElectionGet`: V2ResourcesSkyblockElectionGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SkyBlockAPI.V2ResourcesSkyblockElectionGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2ResourcesSkyblockElectionGetRequest struct via the builder pattern


### Return type

[**V2ResourcesSkyblockElectionGet200Response**](V2ResourcesSkyblockElectionGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ResourcesSkyblockItemsGet

> V2ResourcesSkyblockItemsGet200Response V2ResourcesSkyblockItemsGet(ctx).Execute()

Items



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkyBlockAPI.V2ResourcesSkyblockItemsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkyBlockAPI.V2ResourcesSkyblockItemsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ResourcesSkyblockItemsGet`: V2ResourcesSkyblockItemsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SkyBlockAPI.V2ResourcesSkyblockItemsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2ResourcesSkyblockItemsGetRequest struct via the builder pattern


### Return type

[**V2ResourcesSkyblockItemsGet200Response**](V2ResourcesSkyblockItemsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ResourcesSkyblockSkillsGet

> V2ResourcesSkyblockSkillsGet200Response V2ResourcesSkyblockSkillsGet(ctx).Execute()

Skills



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkyBlockAPI.V2ResourcesSkyblockSkillsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkyBlockAPI.V2ResourcesSkyblockSkillsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ResourcesSkyblockSkillsGet`: V2ResourcesSkyblockSkillsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SkyBlockAPI.V2ResourcesSkyblockSkillsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2ResourcesSkyblockSkillsGetRequest struct via the builder pattern


### Return type

[**V2ResourcesSkyblockSkillsGet200Response**](V2ResourcesSkyblockSkillsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2SkyblockBingoGet

> V2SkyblockBingoGet200Response V2SkyblockBingoGet(ctx).Uuid(uuid).Execute()

Bingo data by player



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
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
	// response from `V2SkyblockBingoGet`: V2SkyblockBingoGet200Response
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

[**V2SkyblockBingoGet200Response**](V2SkyblockBingoGet200Response.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2SkyblockFiresalesGet

> V2SkyblockFiresalesGet200Response V2SkyblockFiresalesGet(ctx).Execute()

Active/Upcoming Fire Sales



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkyBlockAPI.V2SkyblockFiresalesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkyBlockAPI.V2SkyblockFiresalesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2SkyblockFiresalesGet`: V2SkyblockFiresalesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SkyBlockAPI.V2SkyblockFiresalesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2SkyblockFiresalesGetRequest struct via the builder pattern


### Return type

[**V2SkyblockFiresalesGet200Response**](V2SkyblockFiresalesGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2SkyblockGardenGet

> V2SkyblockGardenGet200Response V2SkyblockGardenGet(ctx).Profile(profile).Execute()

Garden data by profile ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
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
	// response from `V2SkyblockGardenGet`: V2SkyblockGardenGet200Response
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

[**V2SkyblockGardenGet200Response**](V2SkyblockGardenGet200Response.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2SkyblockMuseumGet

> V2SkyblockMuseumGet200Response V2SkyblockMuseumGet(ctx).Profile(profile).Execute()

Museum data by profile ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
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
	// response from `V2SkyblockMuseumGet`: V2SkyblockMuseumGet200Response
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

[**V2SkyblockMuseumGet200Response**](V2SkyblockMuseumGet200Response.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2SkyblockNewsGet

> V2SkyblockNewsGet200Response V2SkyblockNewsGet(ctx).Execute()

News

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkyBlockAPI.V2SkyblockNewsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkyBlockAPI.V2SkyblockNewsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2SkyblockNewsGet`: V2SkyblockNewsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SkyBlockAPI.V2SkyblockNewsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2SkyblockNewsGetRequest struct via the builder pattern


### Return type

[**V2SkyblockNewsGet200Response**](V2SkyblockNewsGet200Response.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
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

