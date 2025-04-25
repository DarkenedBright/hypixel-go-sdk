# V2GuildGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Guild** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewV2GuildGet200Response

`func NewV2GuildGet200Response() *V2GuildGet200Response`

NewV2GuildGet200Response instantiates a new V2GuildGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2GuildGet200ResponseWithDefaults

`func NewV2GuildGet200ResponseWithDefaults() *V2GuildGet200Response`

NewV2GuildGet200ResponseWithDefaults instantiates a new V2GuildGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *V2GuildGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *V2GuildGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *V2GuildGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *V2GuildGet200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetGuild

`func (o *V2GuildGet200Response) GetGuild() map[string]interface{}`

GetGuild returns the Guild field if non-nil, zero value otherwise.

### GetGuildOk

`func (o *V2GuildGet200Response) GetGuildOk() (*map[string]interface{}, bool)`

GetGuildOk returns a tuple with the Guild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuild

`func (o *V2GuildGet200Response) SetGuild(v map[string]interface{})`

SetGuild sets Guild field to given value.

### HasGuild

`func (o *V2GuildGet200Response) HasGuild() bool`

HasGuild returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


