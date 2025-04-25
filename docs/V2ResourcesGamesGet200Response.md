# V2ResourcesGamesGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**LastUpdated** | Pointer to **float32** |  | [optional] 
**Games** | Pointer to [**map[string]Game**](Game.md) | A map where the key is the backend name of the game | [optional] 

## Methods

### NewV2ResourcesGamesGet200Response

`func NewV2ResourcesGamesGet200Response() *V2ResourcesGamesGet200Response`

NewV2ResourcesGamesGet200Response instantiates a new V2ResourcesGamesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ResourcesGamesGet200ResponseWithDefaults

`func NewV2ResourcesGamesGet200ResponseWithDefaults() *V2ResourcesGamesGet200Response`

NewV2ResourcesGamesGet200ResponseWithDefaults instantiates a new V2ResourcesGamesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *V2ResourcesGamesGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *V2ResourcesGamesGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *V2ResourcesGamesGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *V2ResourcesGamesGet200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetLastUpdated

`func (o *V2ResourcesGamesGet200Response) GetLastUpdated() float32`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *V2ResourcesGamesGet200Response) GetLastUpdatedOk() (*float32, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *V2ResourcesGamesGet200Response) SetLastUpdated(v float32)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *V2ResourcesGamesGet200Response) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetGames

`func (o *V2ResourcesGamesGet200Response) GetGames() map[string]Game`

GetGames returns the Games field if non-nil, zero value otherwise.

### GetGamesOk

`func (o *V2ResourcesGamesGet200Response) GetGamesOk() (*map[string]Game, bool)`

GetGamesOk returns a tuple with the Games field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGames

`func (o *V2ResourcesGamesGet200Response) SetGames(v map[string]Game)`

SetGames sets Games field to given value.

### HasGames

`func (o *V2ResourcesGamesGet200Response) HasGames() bool`

HasGames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


