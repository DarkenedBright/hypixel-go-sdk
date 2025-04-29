# StatusGetSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GameType** | Pointer to **string** |  | [optional] 
**Map** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**Online** | **bool** |  | 

## Methods

### NewStatusGetSession

`func NewStatusGetSession(online bool, ) *StatusGetSession`

NewStatusGetSession instantiates a new StatusGetSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusGetSessionWithDefaults

`func NewStatusGetSessionWithDefaults() *StatusGetSession`

NewStatusGetSessionWithDefaults instantiates a new StatusGetSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGameType

`func (o *StatusGetSession) GetGameType() string`

GetGameType returns the GameType field if non-nil, zero value otherwise.

### GetGameTypeOk

`func (o *StatusGetSession) GetGameTypeOk() (*string, bool)`

GetGameTypeOk returns a tuple with the GameType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameType

`func (o *StatusGetSession) SetGameType(v string)`

SetGameType sets GameType field to given value.

### HasGameType

`func (o *StatusGetSession) HasGameType() bool`

HasGameType returns a boolean if a field has been set.

### GetMap

`func (o *StatusGetSession) GetMap() string`

GetMap returns the Map field if non-nil, zero value otherwise.

### GetMapOk

`func (o *StatusGetSession) GetMapOk() (*string, bool)`

GetMapOk returns a tuple with the Map field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMap

`func (o *StatusGetSession) SetMap(v string)`

SetMap sets Map field to given value.

### HasMap

`func (o *StatusGetSession) HasMap() bool`

HasMap returns a boolean if a field has been set.

### GetMode

`func (o *StatusGetSession) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *StatusGetSession) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *StatusGetSession) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *StatusGetSession) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetOnline

`func (o *StatusGetSession) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *StatusGetSession) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *StatusGetSession) SetOnline(v bool)`

SetOnline sets Online field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


