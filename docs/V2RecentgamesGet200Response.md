# V2RecentgamesGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Games** | Pointer to [**[]V2RecentgamesGet200ResponseGamesInner**](V2RecentgamesGet200ResponseGamesInner.md) |  | [optional] 

## Methods

### NewV2RecentgamesGet200Response

`func NewV2RecentgamesGet200Response() *V2RecentgamesGet200Response`

NewV2RecentgamesGet200Response instantiates a new V2RecentgamesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2RecentgamesGet200ResponseWithDefaults

`func NewV2RecentgamesGet200ResponseWithDefaults() *V2RecentgamesGet200Response`

NewV2RecentgamesGet200ResponseWithDefaults instantiates a new V2RecentgamesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *V2RecentgamesGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *V2RecentgamesGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *V2RecentgamesGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *V2RecentgamesGet200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetUuid

`func (o *V2RecentgamesGet200Response) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *V2RecentgamesGet200Response) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *V2RecentgamesGet200Response) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *V2RecentgamesGet200Response) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetGames

`func (o *V2RecentgamesGet200Response) GetGames() []V2RecentgamesGet200ResponseGamesInner`

GetGames returns the Games field if non-nil, zero value otherwise.

### GetGamesOk

`func (o *V2RecentgamesGet200Response) GetGamesOk() (*[]V2RecentgamesGet200ResponseGamesInner, bool)`

GetGamesOk returns a tuple with the Games field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGames

`func (o *V2RecentgamesGet200Response) SetGames(v []V2RecentgamesGet200ResponseGamesInner)`

SetGames sets Games field to given value.

### HasGames

`func (o *V2RecentgamesGet200Response) HasGames() bool`

HasGames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


