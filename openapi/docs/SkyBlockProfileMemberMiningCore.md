# SkyBlockProfileMemberMiningCore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Biomes** | Pointer to [**SkyBlockProfileMemberMiningCoreBiomes**](SkyBlockProfileMemberMiningCoreBiomes.md) |  | [optional] 
**Crystals** | Pointer to [**SkyBlockProfileMemberMiningCoreCrystals**](SkyBlockProfileMemberMiningCoreCrystals.md) |  | [optional] 
**CurrentDailyEffect** | Pointer to **string** |  | [optional] 
**CurrentDailyEffectLastChanged** | Pointer to **int64** |  | [optional] 
**DailyOresMined** | Pointer to **int64** |  | [optional] 
**DailyOresMinedDay** | Pointer to **int64** |  | [optional] 
**DailyOresMinedDayGemstone** | Pointer to **int64** |  | [optional] 
**DailyOresMinedDayGlacite** | Pointer to **int64** |  | [optional] 
**DailyOresMinedDayMithrilOre** | Pointer to **int64** |  | [optional] 
**DailyOresMinedGemstone** | Pointer to **int64** |  | [optional] 
**DailyOresMinedGlacite** | Pointer to **int64** |  | [optional] 
**DailyOresMinedMithrilOre** | Pointer to **int64** |  | [optional] 
**Experience** | Pointer to **float64** |  | [optional] 
**GreaterMinesLastAccess** | Pointer to **int64** |  | [optional] 
**HotmMigratorTreeResetSendMessage** | Pointer to **bool** |  | [optional] 
**LastReset** | Pointer to **int64** |  | [optional] 
**Nodes** | [**SkyBlockProfileMemberMiningCoreNodes**](SkyBlockProfileMemberMiningCoreNodes.md) |  | 
**PowderGemstone** | Pointer to **int64** |  | [optional] 
**PowderGemstoneTotal** | Pointer to **int64** |  | [optional] 
**PowderGlacite** | Pointer to **int64** |  | [optional] 
**PowderGlaciteTotal** | Pointer to **int64** |  | [optional] 
**PowderMithril** | Pointer to **int64** |  | [optional] 
**PowderMithrilTotal** | Pointer to **int64** |  | [optional] 
**PowderSpentGemstone** | Pointer to **int64** |  | [optional] 
**PowderSpentGlacite** | Pointer to **int64** |  | [optional] 
**PowderSpentMithril** | Pointer to **int64** |  | [optional] 
**ReceivedFreeTier** | Pointer to **bool** |  | [optional] 
**RetroactiveTier2Token** | Pointer to **bool** |  | [optional] 
**SelectedPickaxeAbility** | Pointer to [**SkyBlockProfileMemberMiningCoreSelectedPickaxeAbility**](SkyBlockProfileMemberMiningCoreSelectedPickaxeAbility.md) |  | [optional] 
**StashIfFullNotification** | Pointer to **bool** |  | [optional] 
**Tokens** | Pointer to **int64** |  | [optional] 
**TokensSpent** | Pointer to **int64** |  | [optional] 

## Methods

### NewSkyBlockProfileMemberMiningCore

`func NewSkyBlockProfileMemberMiningCore(nodes SkyBlockProfileMemberMiningCoreNodes, ) *SkyBlockProfileMemberMiningCore`

NewSkyBlockProfileMemberMiningCore instantiates a new SkyBlockProfileMemberMiningCore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkyBlockProfileMemberMiningCoreWithDefaults

`func NewSkyBlockProfileMemberMiningCoreWithDefaults() *SkyBlockProfileMemberMiningCore`

NewSkyBlockProfileMemberMiningCoreWithDefaults instantiates a new SkyBlockProfileMemberMiningCore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBiomes

`func (o *SkyBlockProfileMemberMiningCore) GetBiomes() SkyBlockProfileMemberMiningCoreBiomes`

GetBiomes returns the Biomes field if non-nil, zero value otherwise.

### GetBiomesOk

`func (o *SkyBlockProfileMemberMiningCore) GetBiomesOk() (*SkyBlockProfileMemberMiningCoreBiomes, bool)`

GetBiomesOk returns a tuple with the Biomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomes

`func (o *SkyBlockProfileMemberMiningCore) SetBiomes(v SkyBlockProfileMemberMiningCoreBiomes)`

SetBiomes sets Biomes field to given value.

### HasBiomes

`func (o *SkyBlockProfileMemberMiningCore) HasBiomes() bool`

HasBiomes returns a boolean if a field has been set.

### GetCrystals

`func (o *SkyBlockProfileMemberMiningCore) GetCrystals() SkyBlockProfileMemberMiningCoreCrystals`

GetCrystals returns the Crystals field if non-nil, zero value otherwise.

### GetCrystalsOk

`func (o *SkyBlockProfileMemberMiningCore) GetCrystalsOk() (*SkyBlockProfileMemberMiningCoreCrystals, bool)`

GetCrystalsOk returns a tuple with the Crystals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrystals

`func (o *SkyBlockProfileMemberMiningCore) SetCrystals(v SkyBlockProfileMemberMiningCoreCrystals)`

SetCrystals sets Crystals field to given value.

### HasCrystals

`func (o *SkyBlockProfileMemberMiningCore) HasCrystals() bool`

HasCrystals returns a boolean if a field has been set.

### GetCurrentDailyEffect

`func (o *SkyBlockProfileMemberMiningCore) GetCurrentDailyEffect() string`

GetCurrentDailyEffect returns the CurrentDailyEffect field if non-nil, zero value otherwise.

### GetCurrentDailyEffectOk

`func (o *SkyBlockProfileMemberMiningCore) GetCurrentDailyEffectOk() (*string, bool)`

GetCurrentDailyEffectOk returns a tuple with the CurrentDailyEffect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDailyEffect

`func (o *SkyBlockProfileMemberMiningCore) SetCurrentDailyEffect(v string)`

SetCurrentDailyEffect sets CurrentDailyEffect field to given value.

### HasCurrentDailyEffect

`func (o *SkyBlockProfileMemberMiningCore) HasCurrentDailyEffect() bool`

HasCurrentDailyEffect returns a boolean if a field has been set.

### GetCurrentDailyEffectLastChanged

`func (o *SkyBlockProfileMemberMiningCore) GetCurrentDailyEffectLastChanged() int64`

GetCurrentDailyEffectLastChanged returns the CurrentDailyEffectLastChanged field if non-nil, zero value otherwise.

### GetCurrentDailyEffectLastChangedOk

`func (o *SkyBlockProfileMemberMiningCore) GetCurrentDailyEffectLastChangedOk() (*int64, bool)`

GetCurrentDailyEffectLastChangedOk returns a tuple with the CurrentDailyEffectLastChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDailyEffectLastChanged

`func (o *SkyBlockProfileMemberMiningCore) SetCurrentDailyEffectLastChanged(v int64)`

SetCurrentDailyEffectLastChanged sets CurrentDailyEffectLastChanged field to given value.

### HasCurrentDailyEffectLastChanged

`func (o *SkyBlockProfileMemberMiningCore) HasCurrentDailyEffectLastChanged() bool`

HasCurrentDailyEffectLastChanged returns a boolean if a field has been set.

### GetDailyOresMined

`func (o *SkyBlockProfileMemberMiningCore) GetDailyOresMined() int64`

GetDailyOresMined returns the DailyOresMined field if non-nil, zero value otherwise.

### GetDailyOresMinedOk

`func (o *SkyBlockProfileMemberMiningCore) GetDailyOresMinedOk() (*int64, bool)`

GetDailyOresMinedOk returns a tuple with the DailyOresMined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyOresMined

`func (o *SkyBlockProfileMemberMiningCore) SetDailyOresMined(v int64)`

SetDailyOresMined sets DailyOresMined field to given value.

### HasDailyOresMined

`func (o *SkyBlockProfileMemberMiningCore) HasDailyOresMined() bool`

HasDailyOresMined returns a boolean if a field has been set.

### GetDailyOresMinedDay

`func (o *SkyBlockProfileMemberMiningCore) GetDailyOresMinedDay() int64`

GetDailyOresMinedDay returns the DailyOresMinedDay field if non-nil, zero value otherwise.

### GetDailyOresMinedDayOk

`func (o *SkyBlockProfileMemberMiningCore) GetDailyOresMinedDayOk() (*int64, bool)`

GetDailyOresMinedDayOk returns a tuple with the DailyOresMinedDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyOresMinedDay

`func (o *SkyBlockProfileMemberMiningCore) SetDailyOresMinedDay(v int64)`

SetDailyOresMinedDay sets DailyOresMinedDay field to given value.

### HasDailyOresMinedDay

`func (o *SkyBlockProfileMemberMiningCore) HasDailyOresMinedDay() bool`

HasDailyOresMinedDay returns a boolean if a field has been set.

### GetDailyOresMinedDayGemstone

`func (o *SkyBlockProfileMemberMiningCore) GetDailyOresMinedDayGemstone() int64`

GetDailyOresMinedDayGemstone returns the DailyOresMinedDayGemstone field if non-nil, zero value otherwise.

### GetDailyOresMinedDayGemstoneOk

`func (o *SkyBlockProfileMemberMiningCore) GetDailyOresMinedDayGemstoneOk() (*int64, bool)`

GetDailyOresMinedDayGemstoneOk returns a tuple with the DailyOresMinedDayGemstone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyOresMinedDayGemstone

`func (o *SkyBlockProfileMemberMiningCore) SetDailyOresMinedDayGemstone(v int64)`

SetDailyOresMinedDayGemstone sets DailyOresMinedDayGemstone field to given value.

### HasDailyOresMinedDayGemstone

`func (o *SkyBlockProfileMemberMiningCore) HasDailyOresMinedDayGemstone() bool`

HasDailyOresMinedDayGemstone returns a boolean if a field has been set.

### GetDailyOresMinedDayGlacite

`func (o *SkyBlockProfileMemberMiningCore) GetDailyOresMinedDayGlacite() int64`

GetDailyOresMinedDayGlacite returns the DailyOresMinedDayGlacite field if non-nil, zero value otherwise.

### GetDailyOresMinedDayGlaciteOk

`func (o *SkyBlockProfileMemberMiningCore) GetDailyOresMinedDayGlaciteOk() (*int64, bool)`

GetDailyOresMinedDayGlaciteOk returns a tuple with the DailyOresMinedDayGlacite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyOresMinedDayGlacite

`func (o *SkyBlockProfileMemberMiningCore) SetDailyOresMinedDayGlacite(v int64)`

SetDailyOresMinedDayGlacite sets DailyOresMinedDayGlacite field to given value.

### HasDailyOresMinedDayGlacite

`func (o *SkyBlockProfileMemberMiningCore) HasDailyOresMinedDayGlacite() bool`

HasDailyOresMinedDayGlacite returns a boolean if a field has been set.

### GetDailyOresMinedDayMithrilOre

`func (o *SkyBlockProfileMemberMiningCore) GetDailyOresMinedDayMithrilOre() int64`

GetDailyOresMinedDayMithrilOre returns the DailyOresMinedDayMithrilOre field if non-nil, zero value otherwise.

### GetDailyOresMinedDayMithrilOreOk

`func (o *SkyBlockProfileMemberMiningCore) GetDailyOresMinedDayMithrilOreOk() (*int64, bool)`

GetDailyOresMinedDayMithrilOreOk returns a tuple with the DailyOresMinedDayMithrilOre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyOresMinedDayMithrilOre

`func (o *SkyBlockProfileMemberMiningCore) SetDailyOresMinedDayMithrilOre(v int64)`

SetDailyOresMinedDayMithrilOre sets DailyOresMinedDayMithrilOre field to given value.

### HasDailyOresMinedDayMithrilOre

`func (o *SkyBlockProfileMemberMiningCore) HasDailyOresMinedDayMithrilOre() bool`

HasDailyOresMinedDayMithrilOre returns a boolean if a field has been set.

### GetDailyOresMinedGemstone

`func (o *SkyBlockProfileMemberMiningCore) GetDailyOresMinedGemstone() int64`

GetDailyOresMinedGemstone returns the DailyOresMinedGemstone field if non-nil, zero value otherwise.

### GetDailyOresMinedGemstoneOk

`func (o *SkyBlockProfileMemberMiningCore) GetDailyOresMinedGemstoneOk() (*int64, bool)`

GetDailyOresMinedGemstoneOk returns a tuple with the DailyOresMinedGemstone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyOresMinedGemstone

`func (o *SkyBlockProfileMemberMiningCore) SetDailyOresMinedGemstone(v int64)`

SetDailyOresMinedGemstone sets DailyOresMinedGemstone field to given value.

### HasDailyOresMinedGemstone

`func (o *SkyBlockProfileMemberMiningCore) HasDailyOresMinedGemstone() bool`

HasDailyOresMinedGemstone returns a boolean if a field has been set.

### GetDailyOresMinedGlacite

`func (o *SkyBlockProfileMemberMiningCore) GetDailyOresMinedGlacite() int64`

GetDailyOresMinedGlacite returns the DailyOresMinedGlacite field if non-nil, zero value otherwise.

### GetDailyOresMinedGlaciteOk

`func (o *SkyBlockProfileMemberMiningCore) GetDailyOresMinedGlaciteOk() (*int64, bool)`

GetDailyOresMinedGlaciteOk returns a tuple with the DailyOresMinedGlacite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyOresMinedGlacite

`func (o *SkyBlockProfileMemberMiningCore) SetDailyOresMinedGlacite(v int64)`

SetDailyOresMinedGlacite sets DailyOresMinedGlacite field to given value.

### HasDailyOresMinedGlacite

`func (o *SkyBlockProfileMemberMiningCore) HasDailyOresMinedGlacite() bool`

HasDailyOresMinedGlacite returns a boolean if a field has been set.

### GetDailyOresMinedMithrilOre

`func (o *SkyBlockProfileMemberMiningCore) GetDailyOresMinedMithrilOre() int64`

GetDailyOresMinedMithrilOre returns the DailyOresMinedMithrilOre field if non-nil, zero value otherwise.

### GetDailyOresMinedMithrilOreOk

`func (o *SkyBlockProfileMemberMiningCore) GetDailyOresMinedMithrilOreOk() (*int64, bool)`

GetDailyOresMinedMithrilOreOk returns a tuple with the DailyOresMinedMithrilOre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyOresMinedMithrilOre

`func (o *SkyBlockProfileMemberMiningCore) SetDailyOresMinedMithrilOre(v int64)`

SetDailyOresMinedMithrilOre sets DailyOresMinedMithrilOre field to given value.

### HasDailyOresMinedMithrilOre

`func (o *SkyBlockProfileMemberMiningCore) HasDailyOresMinedMithrilOre() bool`

HasDailyOresMinedMithrilOre returns a boolean if a field has been set.

### GetExperience

`func (o *SkyBlockProfileMemberMiningCore) GetExperience() float64`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *SkyBlockProfileMemberMiningCore) GetExperienceOk() (*float64, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *SkyBlockProfileMemberMiningCore) SetExperience(v float64)`

SetExperience sets Experience field to given value.

### HasExperience

`func (o *SkyBlockProfileMemberMiningCore) HasExperience() bool`

HasExperience returns a boolean if a field has been set.

### GetGreaterMinesLastAccess

`func (o *SkyBlockProfileMemberMiningCore) GetGreaterMinesLastAccess() int64`

GetGreaterMinesLastAccess returns the GreaterMinesLastAccess field if non-nil, zero value otherwise.

### GetGreaterMinesLastAccessOk

`func (o *SkyBlockProfileMemberMiningCore) GetGreaterMinesLastAccessOk() (*int64, bool)`

GetGreaterMinesLastAccessOk returns a tuple with the GreaterMinesLastAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreaterMinesLastAccess

`func (o *SkyBlockProfileMemberMiningCore) SetGreaterMinesLastAccess(v int64)`

SetGreaterMinesLastAccess sets GreaterMinesLastAccess field to given value.

### HasGreaterMinesLastAccess

`func (o *SkyBlockProfileMemberMiningCore) HasGreaterMinesLastAccess() bool`

HasGreaterMinesLastAccess returns a boolean if a field has been set.

### GetHotmMigratorTreeResetSendMessage

`func (o *SkyBlockProfileMemberMiningCore) GetHotmMigratorTreeResetSendMessage() bool`

GetHotmMigratorTreeResetSendMessage returns the HotmMigratorTreeResetSendMessage field if non-nil, zero value otherwise.

### GetHotmMigratorTreeResetSendMessageOk

`func (o *SkyBlockProfileMemberMiningCore) GetHotmMigratorTreeResetSendMessageOk() (*bool, bool)`

GetHotmMigratorTreeResetSendMessageOk returns a tuple with the HotmMigratorTreeResetSendMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotmMigratorTreeResetSendMessage

`func (o *SkyBlockProfileMemberMiningCore) SetHotmMigratorTreeResetSendMessage(v bool)`

SetHotmMigratorTreeResetSendMessage sets HotmMigratorTreeResetSendMessage field to given value.

### HasHotmMigratorTreeResetSendMessage

`func (o *SkyBlockProfileMemberMiningCore) HasHotmMigratorTreeResetSendMessage() bool`

HasHotmMigratorTreeResetSendMessage returns a boolean if a field has been set.

### GetLastReset

`func (o *SkyBlockProfileMemberMiningCore) GetLastReset() int64`

GetLastReset returns the LastReset field if non-nil, zero value otherwise.

### GetLastResetOk

`func (o *SkyBlockProfileMemberMiningCore) GetLastResetOk() (*int64, bool)`

GetLastResetOk returns a tuple with the LastReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReset

`func (o *SkyBlockProfileMemberMiningCore) SetLastReset(v int64)`

SetLastReset sets LastReset field to given value.

### HasLastReset

`func (o *SkyBlockProfileMemberMiningCore) HasLastReset() bool`

HasLastReset returns a boolean if a field has been set.

### GetNodes

`func (o *SkyBlockProfileMemberMiningCore) GetNodes() SkyBlockProfileMemberMiningCoreNodes`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *SkyBlockProfileMemberMiningCore) GetNodesOk() (*SkyBlockProfileMemberMiningCoreNodes, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *SkyBlockProfileMemberMiningCore) SetNodes(v SkyBlockProfileMemberMiningCoreNodes)`

SetNodes sets Nodes field to given value.


### GetPowderGemstone

`func (o *SkyBlockProfileMemberMiningCore) GetPowderGemstone() int64`

GetPowderGemstone returns the PowderGemstone field if non-nil, zero value otherwise.

### GetPowderGemstoneOk

`func (o *SkyBlockProfileMemberMiningCore) GetPowderGemstoneOk() (*int64, bool)`

GetPowderGemstoneOk returns a tuple with the PowderGemstone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowderGemstone

`func (o *SkyBlockProfileMemberMiningCore) SetPowderGemstone(v int64)`

SetPowderGemstone sets PowderGemstone field to given value.

### HasPowderGemstone

`func (o *SkyBlockProfileMemberMiningCore) HasPowderGemstone() bool`

HasPowderGemstone returns a boolean if a field has been set.

### GetPowderGemstoneTotal

`func (o *SkyBlockProfileMemberMiningCore) GetPowderGemstoneTotal() int64`

GetPowderGemstoneTotal returns the PowderGemstoneTotal field if non-nil, zero value otherwise.

### GetPowderGemstoneTotalOk

`func (o *SkyBlockProfileMemberMiningCore) GetPowderGemstoneTotalOk() (*int64, bool)`

GetPowderGemstoneTotalOk returns a tuple with the PowderGemstoneTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowderGemstoneTotal

`func (o *SkyBlockProfileMemberMiningCore) SetPowderGemstoneTotal(v int64)`

SetPowderGemstoneTotal sets PowderGemstoneTotal field to given value.

### HasPowderGemstoneTotal

`func (o *SkyBlockProfileMemberMiningCore) HasPowderGemstoneTotal() bool`

HasPowderGemstoneTotal returns a boolean if a field has been set.

### GetPowderGlacite

`func (o *SkyBlockProfileMemberMiningCore) GetPowderGlacite() int64`

GetPowderGlacite returns the PowderGlacite field if non-nil, zero value otherwise.

### GetPowderGlaciteOk

`func (o *SkyBlockProfileMemberMiningCore) GetPowderGlaciteOk() (*int64, bool)`

GetPowderGlaciteOk returns a tuple with the PowderGlacite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowderGlacite

`func (o *SkyBlockProfileMemberMiningCore) SetPowderGlacite(v int64)`

SetPowderGlacite sets PowderGlacite field to given value.

### HasPowderGlacite

`func (o *SkyBlockProfileMemberMiningCore) HasPowderGlacite() bool`

HasPowderGlacite returns a boolean if a field has been set.

### GetPowderGlaciteTotal

`func (o *SkyBlockProfileMemberMiningCore) GetPowderGlaciteTotal() int64`

GetPowderGlaciteTotal returns the PowderGlaciteTotal field if non-nil, zero value otherwise.

### GetPowderGlaciteTotalOk

`func (o *SkyBlockProfileMemberMiningCore) GetPowderGlaciteTotalOk() (*int64, bool)`

GetPowderGlaciteTotalOk returns a tuple with the PowderGlaciteTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowderGlaciteTotal

`func (o *SkyBlockProfileMemberMiningCore) SetPowderGlaciteTotal(v int64)`

SetPowderGlaciteTotal sets PowderGlaciteTotal field to given value.

### HasPowderGlaciteTotal

`func (o *SkyBlockProfileMemberMiningCore) HasPowderGlaciteTotal() bool`

HasPowderGlaciteTotal returns a boolean if a field has been set.

### GetPowderMithril

`func (o *SkyBlockProfileMemberMiningCore) GetPowderMithril() int64`

GetPowderMithril returns the PowderMithril field if non-nil, zero value otherwise.

### GetPowderMithrilOk

`func (o *SkyBlockProfileMemberMiningCore) GetPowderMithrilOk() (*int64, bool)`

GetPowderMithrilOk returns a tuple with the PowderMithril field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowderMithril

`func (o *SkyBlockProfileMemberMiningCore) SetPowderMithril(v int64)`

SetPowderMithril sets PowderMithril field to given value.

### HasPowderMithril

`func (o *SkyBlockProfileMemberMiningCore) HasPowderMithril() bool`

HasPowderMithril returns a boolean if a field has been set.

### GetPowderMithrilTotal

`func (o *SkyBlockProfileMemberMiningCore) GetPowderMithrilTotal() int64`

GetPowderMithrilTotal returns the PowderMithrilTotal field if non-nil, zero value otherwise.

### GetPowderMithrilTotalOk

`func (o *SkyBlockProfileMemberMiningCore) GetPowderMithrilTotalOk() (*int64, bool)`

GetPowderMithrilTotalOk returns a tuple with the PowderMithrilTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowderMithrilTotal

`func (o *SkyBlockProfileMemberMiningCore) SetPowderMithrilTotal(v int64)`

SetPowderMithrilTotal sets PowderMithrilTotal field to given value.

### HasPowderMithrilTotal

`func (o *SkyBlockProfileMemberMiningCore) HasPowderMithrilTotal() bool`

HasPowderMithrilTotal returns a boolean if a field has been set.

### GetPowderSpentGemstone

`func (o *SkyBlockProfileMemberMiningCore) GetPowderSpentGemstone() int64`

GetPowderSpentGemstone returns the PowderSpentGemstone field if non-nil, zero value otherwise.

### GetPowderSpentGemstoneOk

`func (o *SkyBlockProfileMemberMiningCore) GetPowderSpentGemstoneOk() (*int64, bool)`

GetPowderSpentGemstoneOk returns a tuple with the PowderSpentGemstone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowderSpentGemstone

`func (o *SkyBlockProfileMemberMiningCore) SetPowderSpentGemstone(v int64)`

SetPowderSpentGemstone sets PowderSpentGemstone field to given value.

### HasPowderSpentGemstone

`func (o *SkyBlockProfileMemberMiningCore) HasPowderSpentGemstone() bool`

HasPowderSpentGemstone returns a boolean if a field has been set.

### GetPowderSpentGlacite

`func (o *SkyBlockProfileMemberMiningCore) GetPowderSpentGlacite() int64`

GetPowderSpentGlacite returns the PowderSpentGlacite field if non-nil, zero value otherwise.

### GetPowderSpentGlaciteOk

`func (o *SkyBlockProfileMemberMiningCore) GetPowderSpentGlaciteOk() (*int64, bool)`

GetPowderSpentGlaciteOk returns a tuple with the PowderSpentGlacite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowderSpentGlacite

`func (o *SkyBlockProfileMemberMiningCore) SetPowderSpentGlacite(v int64)`

SetPowderSpentGlacite sets PowderSpentGlacite field to given value.

### HasPowderSpentGlacite

`func (o *SkyBlockProfileMemberMiningCore) HasPowderSpentGlacite() bool`

HasPowderSpentGlacite returns a boolean if a field has been set.

### GetPowderSpentMithril

`func (o *SkyBlockProfileMemberMiningCore) GetPowderSpentMithril() int64`

GetPowderSpentMithril returns the PowderSpentMithril field if non-nil, zero value otherwise.

### GetPowderSpentMithrilOk

`func (o *SkyBlockProfileMemberMiningCore) GetPowderSpentMithrilOk() (*int64, bool)`

GetPowderSpentMithrilOk returns a tuple with the PowderSpentMithril field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowderSpentMithril

`func (o *SkyBlockProfileMemberMiningCore) SetPowderSpentMithril(v int64)`

SetPowderSpentMithril sets PowderSpentMithril field to given value.

### HasPowderSpentMithril

`func (o *SkyBlockProfileMemberMiningCore) HasPowderSpentMithril() bool`

HasPowderSpentMithril returns a boolean if a field has been set.

### GetReceivedFreeTier

`func (o *SkyBlockProfileMemberMiningCore) GetReceivedFreeTier() bool`

GetReceivedFreeTier returns the ReceivedFreeTier field if non-nil, zero value otherwise.

### GetReceivedFreeTierOk

`func (o *SkyBlockProfileMemberMiningCore) GetReceivedFreeTierOk() (*bool, bool)`

GetReceivedFreeTierOk returns a tuple with the ReceivedFreeTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedFreeTier

`func (o *SkyBlockProfileMemberMiningCore) SetReceivedFreeTier(v bool)`

SetReceivedFreeTier sets ReceivedFreeTier field to given value.

### HasReceivedFreeTier

`func (o *SkyBlockProfileMemberMiningCore) HasReceivedFreeTier() bool`

HasReceivedFreeTier returns a boolean if a field has been set.

### GetRetroactiveTier2Token

`func (o *SkyBlockProfileMemberMiningCore) GetRetroactiveTier2Token() bool`

GetRetroactiveTier2Token returns the RetroactiveTier2Token field if non-nil, zero value otherwise.

### GetRetroactiveTier2TokenOk

`func (o *SkyBlockProfileMemberMiningCore) GetRetroactiveTier2TokenOk() (*bool, bool)`

GetRetroactiveTier2TokenOk returns a tuple with the RetroactiveTier2Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetroactiveTier2Token

`func (o *SkyBlockProfileMemberMiningCore) SetRetroactiveTier2Token(v bool)`

SetRetroactiveTier2Token sets RetroactiveTier2Token field to given value.

### HasRetroactiveTier2Token

`func (o *SkyBlockProfileMemberMiningCore) HasRetroactiveTier2Token() bool`

HasRetroactiveTier2Token returns a boolean if a field has been set.

### GetSelectedPickaxeAbility

`func (o *SkyBlockProfileMemberMiningCore) GetSelectedPickaxeAbility() SkyBlockProfileMemberMiningCoreSelectedPickaxeAbility`

GetSelectedPickaxeAbility returns the SelectedPickaxeAbility field if non-nil, zero value otherwise.

### GetSelectedPickaxeAbilityOk

`func (o *SkyBlockProfileMemberMiningCore) GetSelectedPickaxeAbilityOk() (*SkyBlockProfileMemberMiningCoreSelectedPickaxeAbility, bool)`

GetSelectedPickaxeAbilityOk returns a tuple with the SelectedPickaxeAbility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedPickaxeAbility

`func (o *SkyBlockProfileMemberMiningCore) SetSelectedPickaxeAbility(v SkyBlockProfileMemberMiningCoreSelectedPickaxeAbility)`

SetSelectedPickaxeAbility sets SelectedPickaxeAbility field to given value.

### HasSelectedPickaxeAbility

`func (o *SkyBlockProfileMemberMiningCore) HasSelectedPickaxeAbility() bool`

HasSelectedPickaxeAbility returns a boolean if a field has been set.

### GetStashIfFullNotification

`func (o *SkyBlockProfileMemberMiningCore) GetStashIfFullNotification() bool`

GetStashIfFullNotification returns the StashIfFullNotification field if non-nil, zero value otherwise.

### GetStashIfFullNotificationOk

`func (o *SkyBlockProfileMemberMiningCore) GetStashIfFullNotificationOk() (*bool, bool)`

GetStashIfFullNotificationOk returns a tuple with the StashIfFullNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStashIfFullNotification

`func (o *SkyBlockProfileMemberMiningCore) SetStashIfFullNotification(v bool)`

SetStashIfFullNotification sets StashIfFullNotification field to given value.

### HasStashIfFullNotification

`func (o *SkyBlockProfileMemberMiningCore) HasStashIfFullNotification() bool`

HasStashIfFullNotification returns a boolean if a field has been set.

### GetTokens

`func (o *SkyBlockProfileMemberMiningCore) GetTokens() int64`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *SkyBlockProfileMemberMiningCore) GetTokensOk() (*int64, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *SkyBlockProfileMemberMiningCore) SetTokens(v int64)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *SkyBlockProfileMemberMiningCore) HasTokens() bool`

HasTokens returns a boolean if a field has been set.

### GetTokensSpent

`func (o *SkyBlockProfileMemberMiningCore) GetTokensSpent() int64`

GetTokensSpent returns the TokensSpent field if non-nil, zero value otherwise.

### GetTokensSpentOk

`func (o *SkyBlockProfileMemberMiningCore) GetTokensSpentOk() (*int64, bool)`

GetTokensSpentOk returns a tuple with the TokensSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokensSpent

`func (o *SkyBlockProfileMemberMiningCore) SetTokensSpent(v int64)`

SetTokensSpent sets TokensSpent field to given value.

### HasTokensSpent

`func (o *SkyBlockProfileMemberMiningCore) HasTokensSpent() bool`

HasTokensSpent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


