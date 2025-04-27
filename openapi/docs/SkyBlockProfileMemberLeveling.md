# SkyBlockProfileMemberLeveling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BopBonus** | Pointer to **string** |  | [optional] 
**CategoryExpanded** | Pointer to **bool** |  | [optional] 
**ClaimedTalisman** | Pointer to **bool** |  | [optional] 
**CompletedTasks** | Pointer to [**[]SkyBlockXPTask**](SkyBlockXPTask.md) |  | [optional] 
**Completions** | Pointer to [**SkyBlockProfileMemberLevelingCompletions**](SkyBlockProfileMemberLevelingCompletions.md) |  | [optional] 
**EmblemUnlocks** | Pointer to [**[]SkyBlockEmblem**](SkyBlockEmblem.md) |  | [optional] 
**Experience** | Pointer to **int64** |  | [optional] 
**FishingFestivalSharksKilled** | Pointer to **int64** |  | [optional] 
**HighestPetScore** | Pointer to **int64** |  | [optional] 
**LastViewedTasks** | Pointer to **[]string** |  | [optional] 
**Migrated** | Pointer to **bool** |  | [optional] 
**MigratedCompletions** | Pointer to **bool** |  | [optional] 
**MigratedCompletions2** | Pointer to **bool** |  | [optional] 
**MiningFiestaOresMined** | Pointer to **int64** |  | [optional] 
**SelectedSymbol** | Pointer to [**SkyBlockEmblem**](SkyBlockEmblem.md) |  | [optional] 

## Methods

### NewSkyBlockProfileMemberLeveling

`func NewSkyBlockProfileMemberLeveling() *SkyBlockProfileMemberLeveling`

NewSkyBlockProfileMemberLeveling instantiates a new SkyBlockProfileMemberLeveling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkyBlockProfileMemberLevelingWithDefaults

`func NewSkyBlockProfileMemberLevelingWithDefaults() *SkyBlockProfileMemberLeveling`

NewSkyBlockProfileMemberLevelingWithDefaults instantiates a new SkyBlockProfileMemberLeveling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBopBonus

`func (o *SkyBlockProfileMemberLeveling) GetBopBonus() string`

GetBopBonus returns the BopBonus field if non-nil, zero value otherwise.

### GetBopBonusOk

`func (o *SkyBlockProfileMemberLeveling) GetBopBonusOk() (*string, bool)`

GetBopBonusOk returns a tuple with the BopBonus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBopBonus

`func (o *SkyBlockProfileMemberLeveling) SetBopBonus(v string)`

SetBopBonus sets BopBonus field to given value.

### HasBopBonus

`func (o *SkyBlockProfileMemberLeveling) HasBopBonus() bool`

HasBopBonus returns a boolean if a field has been set.

### GetCategoryExpanded

`func (o *SkyBlockProfileMemberLeveling) GetCategoryExpanded() bool`

GetCategoryExpanded returns the CategoryExpanded field if non-nil, zero value otherwise.

### GetCategoryExpandedOk

`func (o *SkyBlockProfileMemberLeveling) GetCategoryExpandedOk() (*bool, bool)`

GetCategoryExpandedOk returns a tuple with the CategoryExpanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryExpanded

`func (o *SkyBlockProfileMemberLeveling) SetCategoryExpanded(v bool)`

SetCategoryExpanded sets CategoryExpanded field to given value.

### HasCategoryExpanded

`func (o *SkyBlockProfileMemberLeveling) HasCategoryExpanded() bool`

HasCategoryExpanded returns a boolean if a field has been set.

### GetClaimedTalisman

`func (o *SkyBlockProfileMemberLeveling) GetClaimedTalisman() bool`

GetClaimedTalisman returns the ClaimedTalisman field if non-nil, zero value otherwise.

### GetClaimedTalismanOk

`func (o *SkyBlockProfileMemberLeveling) GetClaimedTalismanOk() (*bool, bool)`

GetClaimedTalismanOk returns a tuple with the ClaimedTalisman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedTalisman

`func (o *SkyBlockProfileMemberLeveling) SetClaimedTalisman(v bool)`

SetClaimedTalisman sets ClaimedTalisman field to given value.

### HasClaimedTalisman

`func (o *SkyBlockProfileMemberLeveling) HasClaimedTalisman() bool`

HasClaimedTalisman returns a boolean if a field has been set.

### GetCompletedTasks

`func (o *SkyBlockProfileMemberLeveling) GetCompletedTasks() []SkyBlockXPTask`

GetCompletedTasks returns the CompletedTasks field if non-nil, zero value otherwise.

### GetCompletedTasksOk

`func (o *SkyBlockProfileMemberLeveling) GetCompletedTasksOk() (*[]SkyBlockXPTask, bool)`

GetCompletedTasksOk returns a tuple with the CompletedTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedTasks

`func (o *SkyBlockProfileMemberLeveling) SetCompletedTasks(v []SkyBlockXPTask)`

SetCompletedTasks sets CompletedTasks field to given value.

### HasCompletedTasks

`func (o *SkyBlockProfileMemberLeveling) HasCompletedTasks() bool`

HasCompletedTasks returns a boolean if a field has been set.

### GetCompletions

`func (o *SkyBlockProfileMemberLeveling) GetCompletions() SkyBlockProfileMemberLevelingCompletions`

GetCompletions returns the Completions field if non-nil, zero value otherwise.

### GetCompletionsOk

`func (o *SkyBlockProfileMemberLeveling) GetCompletionsOk() (*SkyBlockProfileMemberLevelingCompletions, bool)`

GetCompletionsOk returns a tuple with the Completions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletions

`func (o *SkyBlockProfileMemberLeveling) SetCompletions(v SkyBlockProfileMemberLevelingCompletions)`

SetCompletions sets Completions field to given value.

### HasCompletions

`func (o *SkyBlockProfileMemberLeveling) HasCompletions() bool`

HasCompletions returns a boolean if a field has been set.

### GetEmblemUnlocks

`func (o *SkyBlockProfileMemberLeveling) GetEmblemUnlocks() []SkyBlockEmblem`

GetEmblemUnlocks returns the EmblemUnlocks field if non-nil, zero value otherwise.

### GetEmblemUnlocksOk

`func (o *SkyBlockProfileMemberLeveling) GetEmblemUnlocksOk() (*[]SkyBlockEmblem, bool)`

GetEmblemUnlocksOk returns a tuple with the EmblemUnlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmblemUnlocks

`func (o *SkyBlockProfileMemberLeveling) SetEmblemUnlocks(v []SkyBlockEmblem)`

SetEmblemUnlocks sets EmblemUnlocks field to given value.

### HasEmblemUnlocks

`func (o *SkyBlockProfileMemberLeveling) HasEmblemUnlocks() bool`

HasEmblemUnlocks returns a boolean if a field has been set.

### GetExperience

`func (o *SkyBlockProfileMemberLeveling) GetExperience() int64`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *SkyBlockProfileMemberLeveling) GetExperienceOk() (*int64, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *SkyBlockProfileMemberLeveling) SetExperience(v int64)`

SetExperience sets Experience field to given value.

### HasExperience

`func (o *SkyBlockProfileMemberLeveling) HasExperience() bool`

HasExperience returns a boolean if a field has been set.

### GetFishingFestivalSharksKilled

`func (o *SkyBlockProfileMemberLeveling) GetFishingFestivalSharksKilled() int64`

GetFishingFestivalSharksKilled returns the FishingFestivalSharksKilled field if non-nil, zero value otherwise.

### GetFishingFestivalSharksKilledOk

`func (o *SkyBlockProfileMemberLeveling) GetFishingFestivalSharksKilledOk() (*int64, bool)`

GetFishingFestivalSharksKilledOk returns a tuple with the FishingFestivalSharksKilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFishingFestivalSharksKilled

`func (o *SkyBlockProfileMemberLeveling) SetFishingFestivalSharksKilled(v int64)`

SetFishingFestivalSharksKilled sets FishingFestivalSharksKilled field to given value.

### HasFishingFestivalSharksKilled

`func (o *SkyBlockProfileMemberLeveling) HasFishingFestivalSharksKilled() bool`

HasFishingFestivalSharksKilled returns a boolean if a field has been set.

### GetHighestPetScore

`func (o *SkyBlockProfileMemberLeveling) GetHighestPetScore() int64`

GetHighestPetScore returns the HighestPetScore field if non-nil, zero value otherwise.

### GetHighestPetScoreOk

`func (o *SkyBlockProfileMemberLeveling) GetHighestPetScoreOk() (*int64, bool)`

GetHighestPetScoreOk returns a tuple with the HighestPetScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestPetScore

`func (o *SkyBlockProfileMemberLeveling) SetHighestPetScore(v int64)`

SetHighestPetScore sets HighestPetScore field to given value.

### HasHighestPetScore

`func (o *SkyBlockProfileMemberLeveling) HasHighestPetScore() bool`

HasHighestPetScore returns a boolean if a field has been set.

### GetLastViewedTasks

`func (o *SkyBlockProfileMemberLeveling) GetLastViewedTasks() []string`

GetLastViewedTasks returns the LastViewedTasks field if non-nil, zero value otherwise.

### GetLastViewedTasksOk

`func (o *SkyBlockProfileMemberLeveling) GetLastViewedTasksOk() (*[]string, bool)`

GetLastViewedTasksOk returns a tuple with the LastViewedTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastViewedTasks

`func (o *SkyBlockProfileMemberLeveling) SetLastViewedTasks(v []string)`

SetLastViewedTasks sets LastViewedTasks field to given value.

### HasLastViewedTasks

`func (o *SkyBlockProfileMemberLeveling) HasLastViewedTasks() bool`

HasLastViewedTasks returns a boolean if a field has been set.

### GetMigrated

`func (o *SkyBlockProfileMemberLeveling) GetMigrated() bool`

GetMigrated returns the Migrated field if non-nil, zero value otherwise.

### GetMigratedOk

`func (o *SkyBlockProfileMemberLeveling) GetMigratedOk() (*bool, bool)`

GetMigratedOk returns a tuple with the Migrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrated

`func (o *SkyBlockProfileMemberLeveling) SetMigrated(v bool)`

SetMigrated sets Migrated field to given value.

### HasMigrated

`func (o *SkyBlockProfileMemberLeveling) HasMigrated() bool`

HasMigrated returns a boolean if a field has been set.

### GetMigratedCompletions

`func (o *SkyBlockProfileMemberLeveling) GetMigratedCompletions() bool`

GetMigratedCompletions returns the MigratedCompletions field if non-nil, zero value otherwise.

### GetMigratedCompletionsOk

`func (o *SkyBlockProfileMemberLeveling) GetMigratedCompletionsOk() (*bool, bool)`

GetMigratedCompletionsOk returns a tuple with the MigratedCompletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratedCompletions

`func (o *SkyBlockProfileMemberLeveling) SetMigratedCompletions(v bool)`

SetMigratedCompletions sets MigratedCompletions field to given value.

### HasMigratedCompletions

`func (o *SkyBlockProfileMemberLeveling) HasMigratedCompletions() bool`

HasMigratedCompletions returns a boolean if a field has been set.

### GetMigratedCompletions2

`func (o *SkyBlockProfileMemberLeveling) GetMigratedCompletions2() bool`

GetMigratedCompletions2 returns the MigratedCompletions2 field if non-nil, zero value otherwise.

### GetMigratedCompletions2Ok

`func (o *SkyBlockProfileMemberLeveling) GetMigratedCompletions2Ok() (*bool, bool)`

GetMigratedCompletions2Ok returns a tuple with the MigratedCompletions2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratedCompletions2

`func (o *SkyBlockProfileMemberLeveling) SetMigratedCompletions2(v bool)`

SetMigratedCompletions2 sets MigratedCompletions2 field to given value.

### HasMigratedCompletions2

`func (o *SkyBlockProfileMemberLeveling) HasMigratedCompletions2() bool`

HasMigratedCompletions2 returns a boolean if a field has been set.

### GetMiningFiestaOresMined

`func (o *SkyBlockProfileMemberLeveling) GetMiningFiestaOresMined() int64`

GetMiningFiestaOresMined returns the MiningFiestaOresMined field if non-nil, zero value otherwise.

### GetMiningFiestaOresMinedOk

`func (o *SkyBlockProfileMemberLeveling) GetMiningFiestaOresMinedOk() (*int64, bool)`

GetMiningFiestaOresMinedOk returns a tuple with the MiningFiestaOresMined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiningFiestaOresMined

`func (o *SkyBlockProfileMemberLeveling) SetMiningFiestaOresMined(v int64)`

SetMiningFiestaOresMined sets MiningFiestaOresMined field to given value.

### HasMiningFiestaOresMined

`func (o *SkyBlockProfileMemberLeveling) HasMiningFiestaOresMined() bool`

HasMiningFiestaOresMined returns a boolean if a field has been set.

### GetSelectedSymbol

`func (o *SkyBlockProfileMemberLeveling) GetSelectedSymbol() SkyBlockEmblem`

GetSelectedSymbol returns the SelectedSymbol field if non-nil, zero value otherwise.

### GetSelectedSymbolOk

`func (o *SkyBlockProfileMemberLeveling) GetSelectedSymbolOk() (*SkyBlockEmblem, bool)`

GetSelectedSymbolOk returns a tuple with the SelectedSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedSymbol

`func (o *SkyBlockProfileMemberLeveling) SetSelectedSymbol(v SkyBlockEmblem)`

SetSelectedSymbol sets SelectedSymbol field to given value.

### HasSelectedSymbol

`func (o *SkyBlockProfileMemberLeveling) HasSelectedSymbol() bool`

HasSelectedSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


