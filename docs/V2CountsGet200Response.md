# V2CountsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**PlayerCount** | Pointer to **int32** |  | [optional] 
**Games** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewV2CountsGet200Response

`func NewV2CountsGet200Response() *V2CountsGet200Response`

NewV2CountsGet200Response instantiates a new V2CountsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2CountsGet200ResponseWithDefaults

`func NewV2CountsGet200ResponseWithDefaults() *V2CountsGet200Response`

NewV2CountsGet200ResponseWithDefaults instantiates a new V2CountsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *V2CountsGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *V2CountsGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *V2CountsGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *V2CountsGet200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetPlayerCount

`func (o *V2CountsGet200Response) GetPlayerCount() int32`

GetPlayerCount returns the PlayerCount field if non-nil, zero value otherwise.

### GetPlayerCountOk

`func (o *V2CountsGet200Response) GetPlayerCountOk() (*int32, bool)`

GetPlayerCountOk returns a tuple with the PlayerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerCount

`func (o *V2CountsGet200Response) SetPlayerCount(v int32)`

SetPlayerCount sets PlayerCount field to given value.

### HasPlayerCount

`func (o *V2CountsGet200Response) HasPlayerCount() bool`

HasPlayerCount returns a boolean if a field has been set.

### GetGames

`func (o *V2CountsGet200Response) GetGames() map[string]interface{}`

GetGames returns the Games field if non-nil, zero value otherwise.

### GetGamesOk

`func (o *V2CountsGet200Response) GetGamesOk() (*map[string]interface{}, bool)`

GetGamesOk returns a tuple with the Games field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGames

`func (o *V2CountsGet200Response) SetGames(v map[string]interface{})`

SetGames sets Games field to given value.

### HasGames

`func (o *V2CountsGet200Response) HasGames() bool`

HasGames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


