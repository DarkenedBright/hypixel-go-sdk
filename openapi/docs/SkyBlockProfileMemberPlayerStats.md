# SkyBlockProfileMemberPlayerStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auctions** | Pointer to [**SkyBlockProfileMemberPlayerStatsAuctions**](SkyBlockProfileMemberPlayerStatsAuctions.md) |  | [optional] 
**CandyCollected** | Pointer to [**SkyBlockProfileMemberPlayerStatsCandyCollected**](SkyBlockProfileMemberPlayerStatsCandyCollected.md) |  | [optional] 
**CropsMined** | Pointer to **float64** |  | [optional] 
**Deaths** | Pointer to **map[string]float64** |  | [optional] 
**EndIsland** | Pointer to [**SkyBlockProfileMemberPlayerStatsEndIsland**](SkyBlockProfileMemberPlayerStatsEndIsland.md) |  | [optional] 
**Gifts** | Pointer to [**SkyBlockProfileMemberPlayerStatsGifts**](SkyBlockProfileMemberPlayerStatsGifts.md) |  | [optional] 
**GlowingMushroomsBroken** | Pointer to **float64** |  | [optional] 
**HighestCriticalDamage** | Pointer to **float64** |  | [optional] 
**HighestDamage** | Pointer to **float64** |  | [optional] 
**ItemsFished** | Pointer to [**SkyBlockProfileMemberPlayerStatsItemsFished**](SkyBlockProfileMemberPlayerStatsItemsFished.md) |  | [optional] 
**Kills** | Pointer to **map[string]float64** |  | [optional] 
**Mythos** | Pointer to [**SkyBlockProfileMemberPlayerStatsMythos**](SkyBlockProfileMemberPlayerStatsMythos.md) |  | [optional] 
**Pets** | Pointer to [**SkyBlockProfileMemberPlayerStatsPets**](SkyBlockProfileMemberPlayerStatsPets.md) |  | [optional] 
**Races** | Pointer to [**SkyBlockProfileMemberPlayerStatsRaces**](SkyBlockProfileMemberPlayerStatsRaces.md) |  | [optional] 
**Rift** | Pointer to [**SkyBlockProfileMemberPlayerStatsRift**](SkyBlockProfileMemberPlayerStatsRift.md) |  | [optional] 
**SeaCreatureKills** | Pointer to **float64** |  | [optional] 
**ShredderRod** | Pointer to [**SkyBlockProfileMemberPlayerStatsShredderRod**](SkyBlockProfileMemberPlayerStatsShredderRod.md) |  | [optional] 
**Spooky** | Pointer to [**SkyBlockProfileMemberPlayerStatsSpooky**](SkyBlockProfileMemberPlayerStatsSpooky.md) |  | [optional] 
**Winter** | Pointer to [**SkyBlockProfileMemberPlayerStatsWinter**](SkyBlockProfileMemberPlayerStatsWinter.md) |  | [optional] 

## Methods

### NewSkyBlockProfileMemberPlayerStats

`func NewSkyBlockProfileMemberPlayerStats() *SkyBlockProfileMemberPlayerStats`

NewSkyBlockProfileMemberPlayerStats instantiates a new SkyBlockProfileMemberPlayerStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkyBlockProfileMemberPlayerStatsWithDefaults

`func NewSkyBlockProfileMemberPlayerStatsWithDefaults() *SkyBlockProfileMemberPlayerStats`

NewSkyBlockProfileMemberPlayerStatsWithDefaults instantiates a new SkyBlockProfileMemberPlayerStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctions

`func (o *SkyBlockProfileMemberPlayerStats) GetAuctions() SkyBlockProfileMemberPlayerStatsAuctions`

GetAuctions returns the Auctions field if non-nil, zero value otherwise.

### GetAuctionsOk

`func (o *SkyBlockProfileMemberPlayerStats) GetAuctionsOk() (*SkyBlockProfileMemberPlayerStatsAuctions, bool)`

GetAuctionsOk returns a tuple with the Auctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctions

`func (o *SkyBlockProfileMemberPlayerStats) SetAuctions(v SkyBlockProfileMemberPlayerStatsAuctions)`

SetAuctions sets Auctions field to given value.

### HasAuctions

`func (o *SkyBlockProfileMemberPlayerStats) HasAuctions() bool`

HasAuctions returns a boolean if a field has been set.

### GetCandyCollected

`func (o *SkyBlockProfileMemberPlayerStats) GetCandyCollected() SkyBlockProfileMemberPlayerStatsCandyCollected`

GetCandyCollected returns the CandyCollected field if non-nil, zero value otherwise.

### GetCandyCollectedOk

`func (o *SkyBlockProfileMemberPlayerStats) GetCandyCollectedOk() (*SkyBlockProfileMemberPlayerStatsCandyCollected, bool)`

GetCandyCollectedOk returns a tuple with the CandyCollected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandyCollected

`func (o *SkyBlockProfileMemberPlayerStats) SetCandyCollected(v SkyBlockProfileMemberPlayerStatsCandyCollected)`

SetCandyCollected sets CandyCollected field to given value.

### HasCandyCollected

`func (o *SkyBlockProfileMemberPlayerStats) HasCandyCollected() bool`

HasCandyCollected returns a boolean if a field has been set.

### GetCropsMined

`func (o *SkyBlockProfileMemberPlayerStats) GetCropsMined() float64`

GetCropsMined returns the CropsMined field if non-nil, zero value otherwise.

### GetCropsMinedOk

`func (o *SkyBlockProfileMemberPlayerStats) GetCropsMinedOk() (*float64, bool)`

GetCropsMinedOk returns a tuple with the CropsMined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCropsMined

`func (o *SkyBlockProfileMemberPlayerStats) SetCropsMined(v float64)`

SetCropsMined sets CropsMined field to given value.

### HasCropsMined

`func (o *SkyBlockProfileMemberPlayerStats) HasCropsMined() bool`

HasCropsMined returns a boolean if a field has been set.

### GetDeaths

`func (o *SkyBlockProfileMemberPlayerStats) GetDeaths() map[string]float64`

GetDeaths returns the Deaths field if non-nil, zero value otherwise.

### GetDeathsOk

`func (o *SkyBlockProfileMemberPlayerStats) GetDeathsOk() (*map[string]float64, bool)`

GetDeathsOk returns a tuple with the Deaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeaths

`func (o *SkyBlockProfileMemberPlayerStats) SetDeaths(v map[string]float64)`

SetDeaths sets Deaths field to given value.

### HasDeaths

`func (o *SkyBlockProfileMemberPlayerStats) HasDeaths() bool`

HasDeaths returns a boolean if a field has been set.

### GetEndIsland

`func (o *SkyBlockProfileMemberPlayerStats) GetEndIsland() SkyBlockProfileMemberPlayerStatsEndIsland`

GetEndIsland returns the EndIsland field if non-nil, zero value otherwise.

### GetEndIslandOk

`func (o *SkyBlockProfileMemberPlayerStats) GetEndIslandOk() (*SkyBlockProfileMemberPlayerStatsEndIsland, bool)`

GetEndIslandOk returns a tuple with the EndIsland field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndIsland

`func (o *SkyBlockProfileMemberPlayerStats) SetEndIsland(v SkyBlockProfileMemberPlayerStatsEndIsland)`

SetEndIsland sets EndIsland field to given value.

### HasEndIsland

`func (o *SkyBlockProfileMemberPlayerStats) HasEndIsland() bool`

HasEndIsland returns a boolean if a field has been set.

### GetGifts

`func (o *SkyBlockProfileMemberPlayerStats) GetGifts() SkyBlockProfileMemberPlayerStatsGifts`

GetGifts returns the Gifts field if non-nil, zero value otherwise.

### GetGiftsOk

`func (o *SkyBlockProfileMemberPlayerStats) GetGiftsOk() (*SkyBlockProfileMemberPlayerStatsGifts, bool)`

GetGiftsOk returns a tuple with the Gifts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGifts

`func (o *SkyBlockProfileMemberPlayerStats) SetGifts(v SkyBlockProfileMemberPlayerStatsGifts)`

SetGifts sets Gifts field to given value.

### HasGifts

`func (o *SkyBlockProfileMemberPlayerStats) HasGifts() bool`

HasGifts returns a boolean if a field has been set.

### GetGlowingMushroomsBroken

`func (o *SkyBlockProfileMemberPlayerStats) GetGlowingMushroomsBroken() float64`

GetGlowingMushroomsBroken returns the GlowingMushroomsBroken field if non-nil, zero value otherwise.

### GetGlowingMushroomsBrokenOk

`func (o *SkyBlockProfileMemberPlayerStats) GetGlowingMushroomsBrokenOk() (*float64, bool)`

GetGlowingMushroomsBrokenOk returns a tuple with the GlowingMushroomsBroken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlowingMushroomsBroken

`func (o *SkyBlockProfileMemberPlayerStats) SetGlowingMushroomsBroken(v float64)`

SetGlowingMushroomsBroken sets GlowingMushroomsBroken field to given value.

### HasGlowingMushroomsBroken

`func (o *SkyBlockProfileMemberPlayerStats) HasGlowingMushroomsBroken() bool`

HasGlowingMushroomsBroken returns a boolean if a field has been set.

### GetHighestCriticalDamage

`func (o *SkyBlockProfileMemberPlayerStats) GetHighestCriticalDamage() float64`

GetHighestCriticalDamage returns the HighestCriticalDamage field if non-nil, zero value otherwise.

### GetHighestCriticalDamageOk

`func (o *SkyBlockProfileMemberPlayerStats) GetHighestCriticalDamageOk() (*float64, bool)`

GetHighestCriticalDamageOk returns a tuple with the HighestCriticalDamage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestCriticalDamage

`func (o *SkyBlockProfileMemberPlayerStats) SetHighestCriticalDamage(v float64)`

SetHighestCriticalDamage sets HighestCriticalDamage field to given value.

### HasHighestCriticalDamage

`func (o *SkyBlockProfileMemberPlayerStats) HasHighestCriticalDamage() bool`

HasHighestCriticalDamage returns a boolean if a field has been set.

### GetHighestDamage

`func (o *SkyBlockProfileMemberPlayerStats) GetHighestDamage() float64`

GetHighestDamage returns the HighestDamage field if non-nil, zero value otherwise.

### GetHighestDamageOk

`func (o *SkyBlockProfileMemberPlayerStats) GetHighestDamageOk() (*float64, bool)`

GetHighestDamageOk returns a tuple with the HighestDamage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestDamage

`func (o *SkyBlockProfileMemberPlayerStats) SetHighestDamage(v float64)`

SetHighestDamage sets HighestDamage field to given value.

### HasHighestDamage

`func (o *SkyBlockProfileMemberPlayerStats) HasHighestDamage() bool`

HasHighestDamage returns a boolean if a field has been set.

### GetItemsFished

`func (o *SkyBlockProfileMemberPlayerStats) GetItemsFished() SkyBlockProfileMemberPlayerStatsItemsFished`

GetItemsFished returns the ItemsFished field if non-nil, zero value otherwise.

### GetItemsFishedOk

`func (o *SkyBlockProfileMemberPlayerStats) GetItemsFishedOk() (*SkyBlockProfileMemberPlayerStatsItemsFished, bool)`

GetItemsFishedOk returns a tuple with the ItemsFished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsFished

`func (o *SkyBlockProfileMemberPlayerStats) SetItemsFished(v SkyBlockProfileMemberPlayerStatsItemsFished)`

SetItemsFished sets ItemsFished field to given value.

### HasItemsFished

`func (o *SkyBlockProfileMemberPlayerStats) HasItemsFished() bool`

HasItemsFished returns a boolean if a field has been set.

### GetKills

`func (o *SkyBlockProfileMemberPlayerStats) GetKills() map[string]float64`

GetKills returns the Kills field if non-nil, zero value otherwise.

### GetKillsOk

`func (o *SkyBlockProfileMemberPlayerStats) GetKillsOk() (*map[string]float64, bool)`

GetKillsOk returns a tuple with the Kills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKills

`func (o *SkyBlockProfileMemberPlayerStats) SetKills(v map[string]float64)`

SetKills sets Kills field to given value.

### HasKills

`func (o *SkyBlockProfileMemberPlayerStats) HasKills() bool`

HasKills returns a boolean if a field has been set.

### GetMythos

`func (o *SkyBlockProfileMemberPlayerStats) GetMythos() SkyBlockProfileMemberPlayerStatsMythos`

GetMythos returns the Mythos field if non-nil, zero value otherwise.

### GetMythosOk

`func (o *SkyBlockProfileMemberPlayerStats) GetMythosOk() (*SkyBlockProfileMemberPlayerStatsMythos, bool)`

GetMythosOk returns a tuple with the Mythos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMythos

`func (o *SkyBlockProfileMemberPlayerStats) SetMythos(v SkyBlockProfileMemberPlayerStatsMythos)`

SetMythos sets Mythos field to given value.

### HasMythos

`func (o *SkyBlockProfileMemberPlayerStats) HasMythos() bool`

HasMythos returns a boolean if a field has been set.

### GetPets

`func (o *SkyBlockProfileMemberPlayerStats) GetPets() SkyBlockProfileMemberPlayerStatsPets`

GetPets returns the Pets field if non-nil, zero value otherwise.

### GetPetsOk

`func (o *SkyBlockProfileMemberPlayerStats) GetPetsOk() (*SkyBlockProfileMemberPlayerStatsPets, bool)`

GetPetsOk returns a tuple with the Pets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPets

`func (o *SkyBlockProfileMemberPlayerStats) SetPets(v SkyBlockProfileMemberPlayerStatsPets)`

SetPets sets Pets field to given value.

### HasPets

`func (o *SkyBlockProfileMemberPlayerStats) HasPets() bool`

HasPets returns a boolean if a field has been set.

### GetRaces

`func (o *SkyBlockProfileMemberPlayerStats) GetRaces() SkyBlockProfileMemberPlayerStatsRaces`

GetRaces returns the Races field if non-nil, zero value otherwise.

### GetRacesOk

`func (o *SkyBlockProfileMemberPlayerStats) GetRacesOk() (*SkyBlockProfileMemberPlayerStatsRaces, bool)`

GetRacesOk returns a tuple with the Races field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaces

`func (o *SkyBlockProfileMemberPlayerStats) SetRaces(v SkyBlockProfileMemberPlayerStatsRaces)`

SetRaces sets Races field to given value.

### HasRaces

`func (o *SkyBlockProfileMemberPlayerStats) HasRaces() bool`

HasRaces returns a boolean if a field has been set.

### GetRift

`func (o *SkyBlockProfileMemberPlayerStats) GetRift() SkyBlockProfileMemberPlayerStatsRift`

GetRift returns the Rift field if non-nil, zero value otherwise.

### GetRiftOk

`func (o *SkyBlockProfileMemberPlayerStats) GetRiftOk() (*SkyBlockProfileMemberPlayerStatsRift, bool)`

GetRiftOk returns a tuple with the Rift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRift

`func (o *SkyBlockProfileMemberPlayerStats) SetRift(v SkyBlockProfileMemberPlayerStatsRift)`

SetRift sets Rift field to given value.

### HasRift

`func (o *SkyBlockProfileMemberPlayerStats) HasRift() bool`

HasRift returns a boolean if a field has been set.

### GetSeaCreatureKills

`func (o *SkyBlockProfileMemberPlayerStats) GetSeaCreatureKills() float64`

GetSeaCreatureKills returns the SeaCreatureKills field if non-nil, zero value otherwise.

### GetSeaCreatureKillsOk

`func (o *SkyBlockProfileMemberPlayerStats) GetSeaCreatureKillsOk() (*float64, bool)`

GetSeaCreatureKillsOk returns a tuple with the SeaCreatureKills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeaCreatureKills

`func (o *SkyBlockProfileMemberPlayerStats) SetSeaCreatureKills(v float64)`

SetSeaCreatureKills sets SeaCreatureKills field to given value.

### HasSeaCreatureKills

`func (o *SkyBlockProfileMemberPlayerStats) HasSeaCreatureKills() bool`

HasSeaCreatureKills returns a boolean if a field has been set.

### GetShredderRod

`func (o *SkyBlockProfileMemberPlayerStats) GetShredderRod() SkyBlockProfileMemberPlayerStatsShredderRod`

GetShredderRod returns the ShredderRod field if non-nil, zero value otherwise.

### GetShredderRodOk

`func (o *SkyBlockProfileMemberPlayerStats) GetShredderRodOk() (*SkyBlockProfileMemberPlayerStatsShredderRod, bool)`

GetShredderRodOk returns a tuple with the ShredderRod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShredderRod

`func (o *SkyBlockProfileMemberPlayerStats) SetShredderRod(v SkyBlockProfileMemberPlayerStatsShredderRod)`

SetShredderRod sets ShredderRod field to given value.

### HasShredderRod

`func (o *SkyBlockProfileMemberPlayerStats) HasShredderRod() bool`

HasShredderRod returns a boolean if a field has been set.

### GetSpooky

`func (o *SkyBlockProfileMemberPlayerStats) GetSpooky() SkyBlockProfileMemberPlayerStatsSpooky`

GetSpooky returns the Spooky field if non-nil, zero value otherwise.

### GetSpookyOk

`func (o *SkyBlockProfileMemberPlayerStats) GetSpookyOk() (*SkyBlockProfileMemberPlayerStatsSpooky, bool)`

GetSpookyOk returns a tuple with the Spooky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpooky

`func (o *SkyBlockProfileMemberPlayerStats) SetSpooky(v SkyBlockProfileMemberPlayerStatsSpooky)`

SetSpooky sets Spooky field to given value.

### HasSpooky

`func (o *SkyBlockProfileMemberPlayerStats) HasSpooky() bool`

HasSpooky returns a boolean if a field has been set.

### GetWinter

`func (o *SkyBlockProfileMemberPlayerStats) GetWinter() SkyBlockProfileMemberPlayerStatsWinter`

GetWinter returns the Winter field if non-nil, zero value otherwise.

### GetWinterOk

`func (o *SkyBlockProfileMemberPlayerStats) GetWinterOk() (*SkyBlockProfileMemberPlayerStatsWinter, bool)`

GetWinterOk returns a tuple with the Winter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinter

`func (o *SkyBlockProfileMemberPlayerStats) SetWinter(v SkyBlockProfileMemberPlayerStatsWinter)`

SetWinter sets Winter field to given value.

### HasWinter

`func (o *SkyBlockProfileMemberPlayerStats) HasWinter() bool`

HasWinter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


