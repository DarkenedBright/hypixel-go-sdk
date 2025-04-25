# SkyBlockProfileMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessoryBagStorage** | Pointer to [**SkyBlockProfileMemberAccessoryBagStorage**](SkyBlockProfileMemberAccessoryBagStorage.md) |  | [optional] 
**Bestiary** | Pointer to [**SkyBlockProfileMemberBestiary**](SkyBlockProfileMemberBestiary.md) |  | [optional] 
**Collection** | Pointer to **map[string]int64** |  | [optional] 
**Currencies** | Pointer to [**SkyBlockProfileMemberCurrencies**](SkyBlockProfileMemberCurrencies.md) |  | [optional] 
**Dungeons** | Pointer to [**SkyBlockProfileMemberDungeons**](SkyBlockProfileMemberDungeons.md) |  | [optional] 
**Events** | Pointer to [**SkyBlockProfileMemberEvents**](SkyBlockProfileMemberEvents.md) |  | [optional] 
**Experimentation** | Pointer to [**SkyBlockProfileMemberExperimentation**](SkyBlockProfileMemberExperimentation.md) |  | [optional] 
**FairySoul** | Pointer to [**SkyBlockProfileMemberFairySoul**](SkyBlockProfileMemberFairySoul.md) |  | [optional] 
**Forge** | Pointer to [**SkyBlockProfileMemberForge**](SkyBlockProfileMemberForge.md) |  | [optional] 
**GardenPlayerData** | Pointer to [**SkyBlockProfileMemberGardenPlayerData**](SkyBlockProfileMemberGardenPlayerData.md) |  | [optional] 
**GlacitePlayerData** | Pointer to [**SkyBlockProfileMemberGlacitePlayerData**](SkyBlockProfileMemberGlacitePlayerData.md) |  | [optional] 
**Inventory** | Pointer to [**SkyBlockProfileMemberInventory**](SkyBlockProfileMemberInventory.md) |  | [optional] 
**PlayerId** | Pointer to **string** |  | [optional] 
**Rift** | Pointer to [**SkyBlockProfileMemberRift**](SkyBlockProfileMemberRift.md) |  | [optional] 

## Methods

### NewSkyBlockProfileMember

`func NewSkyBlockProfileMember() *SkyBlockProfileMember`

NewSkyBlockProfileMember instantiates a new SkyBlockProfileMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkyBlockProfileMemberWithDefaults

`func NewSkyBlockProfileMemberWithDefaults() *SkyBlockProfileMember`

NewSkyBlockProfileMemberWithDefaults instantiates a new SkyBlockProfileMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessoryBagStorage

`func (o *SkyBlockProfileMember) GetAccessoryBagStorage() SkyBlockProfileMemberAccessoryBagStorage`

GetAccessoryBagStorage returns the AccessoryBagStorage field if non-nil, zero value otherwise.

### GetAccessoryBagStorageOk

`func (o *SkyBlockProfileMember) GetAccessoryBagStorageOk() (*SkyBlockProfileMemberAccessoryBagStorage, bool)`

GetAccessoryBagStorageOk returns a tuple with the AccessoryBagStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessoryBagStorage

`func (o *SkyBlockProfileMember) SetAccessoryBagStorage(v SkyBlockProfileMemberAccessoryBagStorage)`

SetAccessoryBagStorage sets AccessoryBagStorage field to given value.

### HasAccessoryBagStorage

`func (o *SkyBlockProfileMember) HasAccessoryBagStorage() bool`

HasAccessoryBagStorage returns a boolean if a field has been set.

### GetBestiary

`func (o *SkyBlockProfileMember) GetBestiary() SkyBlockProfileMemberBestiary`

GetBestiary returns the Bestiary field if non-nil, zero value otherwise.

### GetBestiaryOk

`func (o *SkyBlockProfileMember) GetBestiaryOk() (*SkyBlockProfileMemberBestiary, bool)`

GetBestiaryOk returns a tuple with the Bestiary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestiary

`func (o *SkyBlockProfileMember) SetBestiary(v SkyBlockProfileMemberBestiary)`

SetBestiary sets Bestiary field to given value.

### HasBestiary

`func (o *SkyBlockProfileMember) HasBestiary() bool`

HasBestiary returns a boolean if a field has been set.

### GetCollection

`func (o *SkyBlockProfileMember) GetCollection() map[string]int64`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *SkyBlockProfileMember) GetCollectionOk() (*map[string]int64, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *SkyBlockProfileMember) SetCollection(v map[string]int64)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *SkyBlockProfileMember) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetCurrencies

`func (o *SkyBlockProfileMember) GetCurrencies() SkyBlockProfileMemberCurrencies`

GetCurrencies returns the Currencies field if non-nil, zero value otherwise.

### GetCurrenciesOk

`func (o *SkyBlockProfileMember) GetCurrenciesOk() (*SkyBlockProfileMemberCurrencies, bool)`

GetCurrenciesOk returns a tuple with the Currencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencies

`func (o *SkyBlockProfileMember) SetCurrencies(v SkyBlockProfileMemberCurrencies)`

SetCurrencies sets Currencies field to given value.

### HasCurrencies

`func (o *SkyBlockProfileMember) HasCurrencies() bool`

HasCurrencies returns a boolean if a field has been set.

### GetDungeons

`func (o *SkyBlockProfileMember) GetDungeons() SkyBlockProfileMemberDungeons`

GetDungeons returns the Dungeons field if non-nil, zero value otherwise.

### GetDungeonsOk

`func (o *SkyBlockProfileMember) GetDungeonsOk() (*SkyBlockProfileMemberDungeons, bool)`

GetDungeonsOk returns a tuple with the Dungeons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDungeons

`func (o *SkyBlockProfileMember) SetDungeons(v SkyBlockProfileMemberDungeons)`

SetDungeons sets Dungeons field to given value.

### HasDungeons

`func (o *SkyBlockProfileMember) HasDungeons() bool`

HasDungeons returns a boolean if a field has been set.

### GetEvents

`func (o *SkyBlockProfileMember) GetEvents() SkyBlockProfileMemberEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SkyBlockProfileMember) GetEventsOk() (*SkyBlockProfileMemberEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SkyBlockProfileMember) SetEvents(v SkyBlockProfileMemberEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *SkyBlockProfileMember) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetExperimentation

`func (o *SkyBlockProfileMember) GetExperimentation() SkyBlockProfileMemberExperimentation`

GetExperimentation returns the Experimentation field if non-nil, zero value otherwise.

### GetExperimentationOk

`func (o *SkyBlockProfileMember) GetExperimentationOk() (*SkyBlockProfileMemberExperimentation, bool)`

GetExperimentationOk returns a tuple with the Experimentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentation

`func (o *SkyBlockProfileMember) SetExperimentation(v SkyBlockProfileMemberExperimentation)`

SetExperimentation sets Experimentation field to given value.

### HasExperimentation

`func (o *SkyBlockProfileMember) HasExperimentation() bool`

HasExperimentation returns a boolean if a field has been set.

### GetFairySoul

`func (o *SkyBlockProfileMember) GetFairySoul() SkyBlockProfileMemberFairySoul`

GetFairySoul returns the FairySoul field if non-nil, zero value otherwise.

### GetFairySoulOk

`func (o *SkyBlockProfileMember) GetFairySoulOk() (*SkyBlockProfileMemberFairySoul, bool)`

GetFairySoulOk returns a tuple with the FairySoul field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFairySoul

`func (o *SkyBlockProfileMember) SetFairySoul(v SkyBlockProfileMemberFairySoul)`

SetFairySoul sets FairySoul field to given value.

### HasFairySoul

`func (o *SkyBlockProfileMember) HasFairySoul() bool`

HasFairySoul returns a boolean if a field has been set.

### GetForge

`func (o *SkyBlockProfileMember) GetForge() SkyBlockProfileMemberForge`

GetForge returns the Forge field if non-nil, zero value otherwise.

### GetForgeOk

`func (o *SkyBlockProfileMember) GetForgeOk() (*SkyBlockProfileMemberForge, bool)`

GetForgeOk returns a tuple with the Forge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForge

`func (o *SkyBlockProfileMember) SetForge(v SkyBlockProfileMemberForge)`

SetForge sets Forge field to given value.

### HasForge

`func (o *SkyBlockProfileMember) HasForge() bool`

HasForge returns a boolean if a field has been set.

### GetGardenPlayerData

`func (o *SkyBlockProfileMember) GetGardenPlayerData() SkyBlockProfileMemberGardenPlayerData`

GetGardenPlayerData returns the GardenPlayerData field if non-nil, zero value otherwise.

### GetGardenPlayerDataOk

`func (o *SkyBlockProfileMember) GetGardenPlayerDataOk() (*SkyBlockProfileMemberGardenPlayerData, bool)`

GetGardenPlayerDataOk returns a tuple with the GardenPlayerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGardenPlayerData

`func (o *SkyBlockProfileMember) SetGardenPlayerData(v SkyBlockProfileMemberGardenPlayerData)`

SetGardenPlayerData sets GardenPlayerData field to given value.

### HasGardenPlayerData

`func (o *SkyBlockProfileMember) HasGardenPlayerData() bool`

HasGardenPlayerData returns a boolean if a field has been set.

### GetGlacitePlayerData

`func (o *SkyBlockProfileMember) GetGlacitePlayerData() SkyBlockProfileMemberGlacitePlayerData`

GetGlacitePlayerData returns the GlacitePlayerData field if non-nil, zero value otherwise.

### GetGlacitePlayerDataOk

`func (o *SkyBlockProfileMember) GetGlacitePlayerDataOk() (*SkyBlockProfileMemberGlacitePlayerData, bool)`

GetGlacitePlayerDataOk returns a tuple with the GlacitePlayerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlacitePlayerData

`func (o *SkyBlockProfileMember) SetGlacitePlayerData(v SkyBlockProfileMemberGlacitePlayerData)`

SetGlacitePlayerData sets GlacitePlayerData field to given value.

### HasGlacitePlayerData

`func (o *SkyBlockProfileMember) HasGlacitePlayerData() bool`

HasGlacitePlayerData returns a boolean if a field has been set.

### GetInventory

`func (o *SkyBlockProfileMember) GetInventory() SkyBlockProfileMemberInventory`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *SkyBlockProfileMember) GetInventoryOk() (*SkyBlockProfileMemberInventory, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *SkyBlockProfileMember) SetInventory(v SkyBlockProfileMemberInventory)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *SkyBlockProfileMember) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetPlayerId

`func (o *SkyBlockProfileMember) GetPlayerId() string`

GetPlayerId returns the PlayerId field if non-nil, zero value otherwise.

### GetPlayerIdOk

`func (o *SkyBlockProfileMember) GetPlayerIdOk() (*string, bool)`

GetPlayerIdOk returns a tuple with the PlayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerId

`func (o *SkyBlockProfileMember) SetPlayerId(v string)`

SetPlayerId sets PlayerId field to given value.

### HasPlayerId

`func (o *SkyBlockProfileMember) HasPlayerId() bool`

HasPlayerId returns a boolean if a field has been set.

### GetRift

`func (o *SkyBlockProfileMember) GetRift() SkyBlockProfileMemberRift`

GetRift returns the Rift field if non-nil, zero value otherwise.

### GetRiftOk

`func (o *SkyBlockProfileMember) GetRiftOk() (*SkyBlockProfileMemberRift, bool)`

GetRiftOk returns a tuple with the Rift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRift

`func (o *SkyBlockProfileMember) SetRift(v SkyBlockProfileMemberRift)`

SetRift sets Rift field to given value.

### HasRift

`func (o *SkyBlockProfileMember) HasRift() bool`

HasRift returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


