# \ResourcesAPI

All URIs are relative to *https://api.hypixel.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2ResourcesAchievementsGet**](ResourcesAPI.md#V2ResourcesAchievementsGet) | **Get** /v2/resources/achievements | Achievements
[**V2ResourcesChallengesGet**](ResourcesAPI.md#V2ResourcesChallengesGet) | **Get** /v2/resources/challenges | Challenges
[**V2ResourcesGamesGet**](ResourcesAPI.md#V2ResourcesGamesGet) | **Get** /v2/resources/games | Game Information
[**V2ResourcesGuildsAchievementsGet**](ResourcesAPI.md#V2ResourcesGuildsAchievementsGet) | **Get** /v2/resources/guilds/achievements | Guild Achievements
[**V2ResourcesQuestsGet**](ResourcesAPI.md#V2ResourcesQuestsGet) | **Get** /v2/resources/quests | Quests
[**V2ResourcesVanityCompanionsGet**](ResourcesAPI.md#V2ResourcesVanityCompanionsGet) | **Get** /v2/resources/vanity/companions | Vanity Companions
[**V2ResourcesVanityPetsGet**](ResourcesAPI.md#V2ResourcesVanityPetsGet) | **Get** /v2/resources/vanity/pets | Vanity Pets



## V2ResourcesAchievementsGet

> V2ResourcesAchievementsGet200Response V2ResourcesAchievementsGet(ctx).Execute()

Achievements

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
	resp, r, err := apiClient.ResourcesAPI.V2ResourcesAchievementsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.V2ResourcesAchievementsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ResourcesAchievementsGet`: V2ResourcesAchievementsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.V2ResourcesAchievementsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2ResourcesAchievementsGetRequest struct via the builder pattern


### Return type

[**V2ResourcesAchievementsGet200Response**](V2ResourcesAchievementsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ResourcesChallengesGet

> V2ResourcesChallengesGet200Response V2ResourcesChallengesGet(ctx).Execute()

Challenges

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
	resp, r, err := apiClient.ResourcesAPI.V2ResourcesChallengesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.V2ResourcesChallengesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ResourcesChallengesGet`: V2ResourcesChallengesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.V2ResourcesChallengesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2ResourcesChallengesGetRequest struct via the builder pattern


### Return type

[**V2ResourcesChallengesGet200Response**](V2ResourcesChallengesGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ResourcesGamesGet

> V2ResourcesGamesGet200Response V2ResourcesGamesGet(ctx).Execute()

Game Information



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
	resp, r, err := apiClient.ResourcesAPI.V2ResourcesGamesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.V2ResourcesGamesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ResourcesGamesGet`: V2ResourcesGamesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.V2ResourcesGamesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2ResourcesGamesGetRequest struct via the builder pattern


### Return type

[**V2ResourcesGamesGet200Response**](V2ResourcesGamesGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ResourcesGuildsAchievementsGet

> V2ResourcesGuildsAchievementsGet200Response V2ResourcesGuildsAchievementsGet(ctx).Execute()

Guild Achievements

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
	resp, r, err := apiClient.ResourcesAPI.V2ResourcesGuildsAchievementsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.V2ResourcesGuildsAchievementsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ResourcesGuildsAchievementsGet`: V2ResourcesGuildsAchievementsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.V2ResourcesGuildsAchievementsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2ResourcesGuildsAchievementsGetRequest struct via the builder pattern


### Return type

[**V2ResourcesGuildsAchievementsGet200Response**](V2ResourcesGuildsAchievementsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ResourcesQuestsGet

> V2ResourcesQuestsGet200Response V2ResourcesQuestsGet(ctx).Execute()

Quests

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
	resp, r, err := apiClient.ResourcesAPI.V2ResourcesQuestsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.V2ResourcesQuestsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ResourcesQuestsGet`: V2ResourcesQuestsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.V2ResourcesQuestsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2ResourcesQuestsGetRequest struct via the builder pattern


### Return type

[**V2ResourcesQuestsGet200Response**](V2ResourcesQuestsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ResourcesVanityCompanionsGet

> V2ResourcesVanityPetsGet200Response V2ResourcesVanityCompanionsGet(ctx).Execute()

Vanity Companions

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
	resp, r, err := apiClient.ResourcesAPI.V2ResourcesVanityCompanionsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.V2ResourcesVanityCompanionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ResourcesVanityCompanionsGet`: V2ResourcesVanityPetsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.V2ResourcesVanityCompanionsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2ResourcesVanityCompanionsGetRequest struct via the builder pattern


### Return type

[**V2ResourcesVanityPetsGet200Response**](V2ResourcesVanityPetsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ResourcesVanityPetsGet

> V2ResourcesVanityPetsGet200Response V2ResourcesVanityPetsGet(ctx).Execute()

Vanity Pets

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
	resp, r, err := apiClient.ResourcesAPI.V2ResourcesVanityPetsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.V2ResourcesVanityPetsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ResourcesVanityPetsGet`: V2ResourcesVanityPetsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.V2ResourcesVanityPetsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2ResourcesVanityPetsGetRequest struct via the builder pattern


### Return type

[**V2ResourcesVanityPetsGet200Response**](V2ResourcesVanityPetsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

