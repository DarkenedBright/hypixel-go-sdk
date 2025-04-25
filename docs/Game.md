# Game

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The backend ID of the game. | 
**Name** | **string** | The display name of the game. | 
**DatabaseName** | **string** | The key used for database storage, such as for stats. | 
**ModeNames** | Pointer to **map[string]interface{}** | A map of mode key to display name | [optional] 
**Retired** | Pointer to **bool** | True if the game is retired and no longer playable. | [optional] [default to false]
**Legacy** | Pointer to **bool** | True if the game is legacy and part of the Classic Lobby. | [optional] [default to false]

## Methods

### NewGame

`func NewGame(id int32, name string, databaseName string, ) *Game`

NewGame instantiates a new Game object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameWithDefaults

`func NewGameWithDefaults() *Game`

NewGameWithDefaults instantiates a new Game object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Game) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Game) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Game) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Game) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Game) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Game) SetName(v string)`

SetName sets Name field to given value.


### GetDatabaseName

`func (o *Game) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *Game) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *Game) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### GetModeNames

`func (o *Game) GetModeNames() map[string]interface{}`

GetModeNames returns the ModeNames field if non-nil, zero value otherwise.

### GetModeNamesOk

`func (o *Game) GetModeNamesOk() (*map[string]interface{}, bool)`

GetModeNamesOk returns a tuple with the ModeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeNames

`func (o *Game) SetModeNames(v map[string]interface{})`

SetModeNames sets ModeNames field to given value.

### HasModeNames

`func (o *Game) HasModeNames() bool`

HasModeNames returns a boolean if a field has been set.

### GetRetired

`func (o *Game) GetRetired() bool`

GetRetired returns the Retired field if non-nil, zero value otherwise.

### GetRetiredOk

`func (o *Game) GetRetiredOk() (*bool, bool)`

GetRetiredOk returns a tuple with the Retired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetired

`func (o *Game) SetRetired(v bool)`

SetRetired sets Retired field to given value.

### HasRetired

`func (o *Game) HasRetired() bool`

HasRetired returns a boolean if a field has been set.

### GetLegacy

`func (o *Game) GetLegacy() bool`

GetLegacy returns the Legacy field if non-nil, zero value otherwise.

### GetLegacyOk

`func (o *Game) GetLegacyOk() (*bool, bool)`

GetLegacyOk returns a tuple with the Legacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacy

`func (o *Game) SetLegacy(v bool)`

SetLegacy sets Legacy field to given value.

### HasLegacy

`func (o *Game) HasLegacy() bool`

HasLegacy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


