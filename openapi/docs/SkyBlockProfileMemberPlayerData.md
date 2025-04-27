# SkyBlockProfileMemberPlayerData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AchievementSpawnedIslandTypes** | Pointer to [**[]SkyBlockProfileMemberPlayerDataAchievementSpawnedIslandTypesInner**](SkyBlockProfileMemberPlayerDataAchievementSpawnedIslandTypesInner.md) |  | [optional] 
**ActiveEffects** | Pointer to [**[]SkyBlockProfileMemberPlayerDataActiveEffect**](SkyBlockProfileMemberPlayerDataActiveEffect.md) |  | [optional] 
**CraftedGenerators** | Pointer to [**[]SkyBlockProfileMemberPlayerDataCraftedGeneratorsInner**](SkyBlockProfileMemberPlayerDataCraftedGeneratorsInner.md) |  | [optional] 
**DeathCount** | Pointer to **int64** |  | [optional] 
**DisabledPotionEffects** | Pointer to [**[]SkyBlockProfileMemberPlayerDataActiveEffectType**](SkyBlockProfileMemberPlayerDataActiveEffectType.md) |  | [optional] 
**Experience** | Pointer to [**SkyBlockProfileMemberPlayerDataExperience**](SkyBlockProfileMemberPlayerDataExperience.md) |  | [optional] 
**FastestTargetPractice** | Pointer to **float64** |  | [optional] 
**FishingTreasureCaught** | Pointer to **int64** |  | [optional] 
**LastDeath** | Pointer to **int64** |  | [optional] 
**PausedEffects** | Pointer to [**[]SkyBlockProfileMemberPlayerDataActiveEffect**](SkyBlockProfileMemberPlayerDataActiveEffect.md) |  | [optional] 
**Perks** | Pointer to [**SkyBlockProfileMemberPlayerDataPerks**](SkyBlockProfileMemberPlayerDataPerks.md) |  | [optional] 
**ReaperPeppersEaten** | Pointer to **int64** |  | [optional] 
**TempStatBuffs** | Pointer to [**[]SkyBlockProfileMemberPlayerDataTempStatBuff**](SkyBlockProfileMemberPlayerDataTempStatBuff.md) |  | [optional] 
**UnlockedCollTiers** | Pointer to [**[]SkyBlockProfileMemberPlayerDataUnlockedCollTiersInner**](SkyBlockProfileMemberPlayerDataUnlockedCollTiersInner.md) |  | [optional] 
**VisitedModes** | Pointer to [**[]SkyBlockProfileMemberPlayerDataVisitedModesInner**](SkyBlockProfileMemberPlayerDataVisitedModesInner.md) |  | [optional] 
**VisitedZones** | Pointer to [**[]SkyBlockProfileMemberPlayerDataVisitedZonesInner**](SkyBlockProfileMemberPlayerDataVisitedZonesInner.md) |  | [optional] 

## Methods

### NewSkyBlockProfileMemberPlayerData

`func NewSkyBlockProfileMemberPlayerData() *SkyBlockProfileMemberPlayerData`

NewSkyBlockProfileMemberPlayerData instantiates a new SkyBlockProfileMemberPlayerData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkyBlockProfileMemberPlayerDataWithDefaults

`func NewSkyBlockProfileMemberPlayerDataWithDefaults() *SkyBlockProfileMemberPlayerData`

NewSkyBlockProfileMemberPlayerDataWithDefaults instantiates a new SkyBlockProfileMemberPlayerData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAchievementSpawnedIslandTypes

`func (o *SkyBlockProfileMemberPlayerData) GetAchievementSpawnedIslandTypes() []SkyBlockProfileMemberPlayerDataAchievementSpawnedIslandTypesInner`

GetAchievementSpawnedIslandTypes returns the AchievementSpawnedIslandTypes field if non-nil, zero value otherwise.

### GetAchievementSpawnedIslandTypesOk

`func (o *SkyBlockProfileMemberPlayerData) GetAchievementSpawnedIslandTypesOk() (*[]SkyBlockProfileMemberPlayerDataAchievementSpawnedIslandTypesInner, bool)`

GetAchievementSpawnedIslandTypesOk returns a tuple with the AchievementSpawnedIslandTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementSpawnedIslandTypes

`func (o *SkyBlockProfileMemberPlayerData) SetAchievementSpawnedIslandTypes(v []SkyBlockProfileMemberPlayerDataAchievementSpawnedIslandTypesInner)`

SetAchievementSpawnedIslandTypes sets AchievementSpawnedIslandTypes field to given value.

### HasAchievementSpawnedIslandTypes

`func (o *SkyBlockProfileMemberPlayerData) HasAchievementSpawnedIslandTypes() bool`

HasAchievementSpawnedIslandTypes returns a boolean if a field has been set.

### GetActiveEffects

`func (o *SkyBlockProfileMemberPlayerData) GetActiveEffects() []SkyBlockProfileMemberPlayerDataActiveEffect`

GetActiveEffects returns the ActiveEffects field if non-nil, zero value otherwise.

### GetActiveEffectsOk

`func (o *SkyBlockProfileMemberPlayerData) GetActiveEffectsOk() (*[]SkyBlockProfileMemberPlayerDataActiveEffect, bool)`

GetActiveEffectsOk returns a tuple with the ActiveEffects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveEffects

`func (o *SkyBlockProfileMemberPlayerData) SetActiveEffects(v []SkyBlockProfileMemberPlayerDataActiveEffect)`

SetActiveEffects sets ActiveEffects field to given value.

### HasActiveEffects

`func (o *SkyBlockProfileMemberPlayerData) HasActiveEffects() bool`

HasActiveEffects returns a boolean if a field has been set.

### GetCraftedGenerators

`func (o *SkyBlockProfileMemberPlayerData) GetCraftedGenerators() []SkyBlockProfileMemberPlayerDataCraftedGeneratorsInner`

GetCraftedGenerators returns the CraftedGenerators field if non-nil, zero value otherwise.

### GetCraftedGeneratorsOk

`func (o *SkyBlockProfileMemberPlayerData) GetCraftedGeneratorsOk() (*[]SkyBlockProfileMemberPlayerDataCraftedGeneratorsInner, bool)`

GetCraftedGeneratorsOk returns a tuple with the CraftedGenerators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCraftedGenerators

`func (o *SkyBlockProfileMemberPlayerData) SetCraftedGenerators(v []SkyBlockProfileMemberPlayerDataCraftedGeneratorsInner)`

SetCraftedGenerators sets CraftedGenerators field to given value.

### HasCraftedGenerators

`func (o *SkyBlockProfileMemberPlayerData) HasCraftedGenerators() bool`

HasCraftedGenerators returns a boolean if a field has been set.

### GetDeathCount

`func (o *SkyBlockProfileMemberPlayerData) GetDeathCount() int64`

GetDeathCount returns the DeathCount field if non-nil, zero value otherwise.

### GetDeathCountOk

`func (o *SkyBlockProfileMemberPlayerData) GetDeathCountOk() (*int64, bool)`

GetDeathCountOk returns a tuple with the DeathCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeathCount

`func (o *SkyBlockProfileMemberPlayerData) SetDeathCount(v int64)`

SetDeathCount sets DeathCount field to given value.

### HasDeathCount

`func (o *SkyBlockProfileMemberPlayerData) HasDeathCount() bool`

HasDeathCount returns a boolean if a field has been set.

### GetDisabledPotionEffects

`func (o *SkyBlockProfileMemberPlayerData) GetDisabledPotionEffects() []SkyBlockProfileMemberPlayerDataActiveEffectType`

GetDisabledPotionEffects returns the DisabledPotionEffects field if non-nil, zero value otherwise.

### GetDisabledPotionEffectsOk

`func (o *SkyBlockProfileMemberPlayerData) GetDisabledPotionEffectsOk() (*[]SkyBlockProfileMemberPlayerDataActiveEffectType, bool)`

GetDisabledPotionEffectsOk returns a tuple with the DisabledPotionEffects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledPotionEffects

`func (o *SkyBlockProfileMemberPlayerData) SetDisabledPotionEffects(v []SkyBlockProfileMemberPlayerDataActiveEffectType)`

SetDisabledPotionEffects sets DisabledPotionEffects field to given value.

### HasDisabledPotionEffects

`func (o *SkyBlockProfileMemberPlayerData) HasDisabledPotionEffects() bool`

HasDisabledPotionEffects returns a boolean if a field has been set.

### GetExperience

`func (o *SkyBlockProfileMemberPlayerData) GetExperience() SkyBlockProfileMemberPlayerDataExperience`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *SkyBlockProfileMemberPlayerData) GetExperienceOk() (*SkyBlockProfileMemberPlayerDataExperience, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *SkyBlockProfileMemberPlayerData) SetExperience(v SkyBlockProfileMemberPlayerDataExperience)`

SetExperience sets Experience field to given value.

### HasExperience

`func (o *SkyBlockProfileMemberPlayerData) HasExperience() bool`

HasExperience returns a boolean if a field has been set.

### GetFastestTargetPractice

`func (o *SkyBlockProfileMemberPlayerData) GetFastestTargetPractice() float64`

GetFastestTargetPractice returns the FastestTargetPractice field if non-nil, zero value otherwise.

### GetFastestTargetPracticeOk

`func (o *SkyBlockProfileMemberPlayerData) GetFastestTargetPracticeOk() (*float64, bool)`

GetFastestTargetPracticeOk returns a tuple with the FastestTargetPractice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastestTargetPractice

`func (o *SkyBlockProfileMemberPlayerData) SetFastestTargetPractice(v float64)`

SetFastestTargetPractice sets FastestTargetPractice field to given value.

### HasFastestTargetPractice

`func (o *SkyBlockProfileMemberPlayerData) HasFastestTargetPractice() bool`

HasFastestTargetPractice returns a boolean if a field has been set.

### GetFishingTreasureCaught

`func (o *SkyBlockProfileMemberPlayerData) GetFishingTreasureCaught() int64`

GetFishingTreasureCaught returns the FishingTreasureCaught field if non-nil, zero value otherwise.

### GetFishingTreasureCaughtOk

`func (o *SkyBlockProfileMemberPlayerData) GetFishingTreasureCaughtOk() (*int64, bool)`

GetFishingTreasureCaughtOk returns a tuple with the FishingTreasureCaught field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFishingTreasureCaught

`func (o *SkyBlockProfileMemberPlayerData) SetFishingTreasureCaught(v int64)`

SetFishingTreasureCaught sets FishingTreasureCaught field to given value.

### HasFishingTreasureCaught

`func (o *SkyBlockProfileMemberPlayerData) HasFishingTreasureCaught() bool`

HasFishingTreasureCaught returns a boolean if a field has been set.

### GetLastDeath

`func (o *SkyBlockProfileMemberPlayerData) GetLastDeath() int64`

GetLastDeath returns the LastDeath field if non-nil, zero value otherwise.

### GetLastDeathOk

`func (o *SkyBlockProfileMemberPlayerData) GetLastDeathOk() (*int64, bool)`

GetLastDeathOk returns a tuple with the LastDeath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeath

`func (o *SkyBlockProfileMemberPlayerData) SetLastDeath(v int64)`

SetLastDeath sets LastDeath field to given value.

### HasLastDeath

`func (o *SkyBlockProfileMemberPlayerData) HasLastDeath() bool`

HasLastDeath returns a boolean if a field has been set.

### GetPausedEffects

`func (o *SkyBlockProfileMemberPlayerData) GetPausedEffects() []SkyBlockProfileMemberPlayerDataActiveEffect`

GetPausedEffects returns the PausedEffects field if non-nil, zero value otherwise.

### GetPausedEffectsOk

`func (o *SkyBlockProfileMemberPlayerData) GetPausedEffectsOk() (*[]SkyBlockProfileMemberPlayerDataActiveEffect, bool)`

GetPausedEffectsOk returns a tuple with the PausedEffects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPausedEffects

`func (o *SkyBlockProfileMemberPlayerData) SetPausedEffects(v []SkyBlockProfileMemberPlayerDataActiveEffect)`

SetPausedEffects sets PausedEffects field to given value.

### HasPausedEffects

`func (o *SkyBlockProfileMemberPlayerData) HasPausedEffects() bool`

HasPausedEffects returns a boolean if a field has been set.

### GetPerks

`func (o *SkyBlockProfileMemberPlayerData) GetPerks() SkyBlockProfileMemberPlayerDataPerks`

GetPerks returns the Perks field if non-nil, zero value otherwise.

### GetPerksOk

`func (o *SkyBlockProfileMemberPlayerData) GetPerksOk() (*SkyBlockProfileMemberPlayerDataPerks, bool)`

GetPerksOk returns a tuple with the Perks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerks

`func (o *SkyBlockProfileMemberPlayerData) SetPerks(v SkyBlockProfileMemberPlayerDataPerks)`

SetPerks sets Perks field to given value.

### HasPerks

`func (o *SkyBlockProfileMemberPlayerData) HasPerks() bool`

HasPerks returns a boolean if a field has been set.

### GetReaperPeppersEaten

`func (o *SkyBlockProfileMemberPlayerData) GetReaperPeppersEaten() int64`

GetReaperPeppersEaten returns the ReaperPeppersEaten field if non-nil, zero value otherwise.

### GetReaperPeppersEatenOk

`func (o *SkyBlockProfileMemberPlayerData) GetReaperPeppersEatenOk() (*int64, bool)`

GetReaperPeppersEatenOk returns a tuple with the ReaperPeppersEaten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReaperPeppersEaten

`func (o *SkyBlockProfileMemberPlayerData) SetReaperPeppersEaten(v int64)`

SetReaperPeppersEaten sets ReaperPeppersEaten field to given value.

### HasReaperPeppersEaten

`func (o *SkyBlockProfileMemberPlayerData) HasReaperPeppersEaten() bool`

HasReaperPeppersEaten returns a boolean if a field has been set.

### GetTempStatBuffs

`func (o *SkyBlockProfileMemberPlayerData) GetTempStatBuffs() []SkyBlockProfileMemberPlayerDataTempStatBuff`

GetTempStatBuffs returns the TempStatBuffs field if non-nil, zero value otherwise.

### GetTempStatBuffsOk

`func (o *SkyBlockProfileMemberPlayerData) GetTempStatBuffsOk() (*[]SkyBlockProfileMemberPlayerDataTempStatBuff, bool)`

GetTempStatBuffsOk returns a tuple with the TempStatBuffs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempStatBuffs

`func (o *SkyBlockProfileMemberPlayerData) SetTempStatBuffs(v []SkyBlockProfileMemberPlayerDataTempStatBuff)`

SetTempStatBuffs sets TempStatBuffs field to given value.

### HasTempStatBuffs

`func (o *SkyBlockProfileMemberPlayerData) HasTempStatBuffs() bool`

HasTempStatBuffs returns a boolean if a field has been set.

### GetUnlockedCollTiers

`func (o *SkyBlockProfileMemberPlayerData) GetUnlockedCollTiers() []SkyBlockProfileMemberPlayerDataUnlockedCollTiersInner`

GetUnlockedCollTiers returns the UnlockedCollTiers field if non-nil, zero value otherwise.

### GetUnlockedCollTiersOk

`func (o *SkyBlockProfileMemberPlayerData) GetUnlockedCollTiersOk() (*[]SkyBlockProfileMemberPlayerDataUnlockedCollTiersInner, bool)`

GetUnlockedCollTiersOk returns a tuple with the UnlockedCollTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockedCollTiers

`func (o *SkyBlockProfileMemberPlayerData) SetUnlockedCollTiers(v []SkyBlockProfileMemberPlayerDataUnlockedCollTiersInner)`

SetUnlockedCollTiers sets UnlockedCollTiers field to given value.

### HasUnlockedCollTiers

`func (o *SkyBlockProfileMemberPlayerData) HasUnlockedCollTiers() bool`

HasUnlockedCollTiers returns a boolean if a field has been set.

### GetVisitedModes

`func (o *SkyBlockProfileMemberPlayerData) GetVisitedModes() []SkyBlockProfileMemberPlayerDataVisitedModesInner`

GetVisitedModes returns the VisitedModes field if non-nil, zero value otherwise.

### GetVisitedModesOk

`func (o *SkyBlockProfileMemberPlayerData) GetVisitedModesOk() (*[]SkyBlockProfileMemberPlayerDataVisitedModesInner, bool)`

GetVisitedModesOk returns a tuple with the VisitedModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitedModes

`func (o *SkyBlockProfileMemberPlayerData) SetVisitedModes(v []SkyBlockProfileMemberPlayerDataVisitedModesInner)`

SetVisitedModes sets VisitedModes field to given value.

### HasVisitedModes

`func (o *SkyBlockProfileMemberPlayerData) HasVisitedModes() bool`

HasVisitedModes returns a boolean if a field has been set.

### GetVisitedZones

`func (o *SkyBlockProfileMemberPlayerData) GetVisitedZones() []SkyBlockProfileMemberPlayerDataVisitedZonesInner`

GetVisitedZones returns the VisitedZones field if non-nil, zero value otherwise.

### GetVisitedZonesOk

`func (o *SkyBlockProfileMemberPlayerData) GetVisitedZonesOk() (*[]SkyBlockProfileMemberPlayerDataVisitedZonesInner, bool)`

GetVisitedZonesOk returns a tuple with the VisitedZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitedZones

`func (o *SkyBlockProfileMemberPlayerData) SetVisitedZones(v []SkyBlockProfileMemberPlayerDataVisitedZonesInner)`

SetVisitedZones sets VisitedZones field to given value.

### HasVisitedZones

`func (o *SkyBlockProfileMemberPlayerData) HasVisitedZones() bool`

HasVisitedZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


