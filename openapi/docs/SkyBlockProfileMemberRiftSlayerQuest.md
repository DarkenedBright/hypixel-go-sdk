# SkyBlockProfileMemberRiftSlayerQuest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CombatXp** | Pointer to **int64** |  | [optional] 
**CompletionState** | **int64** |  | 
**LastKilledMobIsland** | Pointer to [**SkyBlockProfileMemberRiftSlayerQuestLastKilledMobIsland**](SkyBlockProfileMemberRiftSlayerQuestLastKilledMobIsland.md) |  | [optional] 
**RecentMobKills** | Pointer to [**[]SkyBlockProfileMemberRiftSlayerQuestRecentMobKill**](SkyBlockProfileMemberRiftSlayerQuestRecentMobKill.md) |  | [optional] 
**Solo** | **bool** |  | 
**SpawnTimestamp** | Pointer to **int64** |  | [optional] 
**StartTimestamp** | **int64** |  | 
**Tier** | **int64** |  | 
**Type** | [**SkyBlockProfileMemberRiftSlayerQuestType**](SkyBlockProfileMemberRiftSlayerQuestType.md) |  | 
**UsedArmor** | **bool** |  | 

## Methods

### NewSkyBlockProfileMemberRiftSlayerQuest

`func NewSkyBlockProfileMemberRiftSlayerQuest(completionState int64, solo bool, startTimestamp int64, tier int64, type_ SkyBlockProfileMemberRiftSlayerQuestType, usedArmor bool, ) *SkyBlockProfileMemberRiftSlayerQuest`

NewSkyBlockProfileMemberRiftSlayerQuest instantiates a new SkyBlockProfileMemberRiftSlayerQuest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkyBlockProfileMemberRiftSlayerQuestWithDefaults

`func NewSkyBlockProfileMemberRiftSlayerQuestWithDefaults() *SkyBlockProfileMemberRiftSlayerQuest`

NewSkyBlockProfileMemberRiftSlayerQuestWithDefaults instantiates a new SkyBlockProfileMemberRiftSlayerQuest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCombatXp

`func (o *SkyBlockProfileMemberRiftSlayerQuest) GetCombatXp() int64`

GetCombatXp returns the CombatXp field if non-nil, zero value otherwise.

### GetCombatXpOk

`func (o *SkyBlockProfileMemberRiftSlayerQuest) GetCombatXpOk() (*int64, bool)`

GetCombatXpOk returns a tuple with the CombatXp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombatXp

`func (o *SkyBlockProfileMemberRiftSlayerQuest) SetCombatXp(v int64)`

SetCombatXp sets CombatXp field to given value.

### HasCombatXp

`func (o *SkyBlockProfileMemberRiftSlayerQuest) HasCombatXp() bool`

HasCombatXp returns a boolean if a field has been set.

### GetCompletionState

`func (o *SkyBlockProfileMemberRiftSlayerQuest) GetCompletionState() int64`

GetCompletionState returns the CompletionState field if non-nil, zero value otherwise.

### GetCompletionStateOk

`func (o *SkyBlockProfileMemberRiftSlayerQuest) GetCompletionStateOk() (*int64, bool)`

GetCompletionStateOk returns a tuple with the CompletionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionState

`func (o *SkyBlockProfileMemberRiftSlayerQuest) SetCompletionState(v int64)`

SetCompletionState sets CompletionState field to given value.


### GetLastKilledMobIsland

`func (o *SkyBlockProfileMemberRiftSlayerQuest) GetLastKilledMobIsland() SkyBlockProfileMemberRiftSlayerQuestLastKilledMobIsland`

GetLastKilledMobIsland returns the LastKilledMobIsland field if non-nil, zero value otherwise.

### GetLastKilledMobIslandOk

`func (o *SkyBlockProfileMemberRiftSlayerQuest) GetLastKilledMobIslandOk() (*SkyBlockProfileMemberRiftSlayerQuestLastKilledMobIsland, bool)`

GetLastKilledMobIslandOk returns a tuple with the LastKilledMobIsland field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastKilledMobIsland

`func (o *SkyBlockProfileMemberRiftSlayerQuest) SetLastKilledMobIsland(v SkyBlockProfileMemberRiftSlayerQuestLastKilledMobIsland)`

SetLastKilledMobIsland sets LastKilledMobIsland field to given value.

### HasLastKilledMobIsland

`func (o *SkyBlockProfileMemberRiftSlayerQuest) HasLastKilledMobIsland() bool`

HasLastKilledMobIsland returns a boolean if a field has been set.

### GetRecentMobKills

`func (o *SkyBlockProfileMemberRiftSlayerQuest) GetRecentMobKills() []SkyBlockProfileMemberRiftSlayerQuestRecentMobKill`

GetRecentMobKills returns the RecentMobKills field if non-nil, zero value otherwise.

### GetRecentMobKillsOk

`func (o *SkyBlockProfileMemberRiftSlayerQuest) GetRecentMobKillsOk() (*[]SkyBlockProfileMemberRiftSlayerQuestRecentMobKill, bool)`

GetRecentMobKillsOk returns a tuple with the RecentMobKills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentMobKills

`func (o *SkyBlockProfileMemberRiftSlayerQuest) SetRecentMobKills(v []SkyBlockProfileMemberRiftSlayerQuestRecentMobKill)`

SetRecentMobKills sets RecentMobKills field to given value.

### HasRecentMobKills

`func (o *SkyBlockProfileMemberRiftSlayerQuest) HasRecentMobKills() bool`

HasRecentMobKills returns a boolean if a field has been set.

### GetSolo

`func (o *SkyBlockProfileMemberRiftSlayerQuest) GetSolo() bool`

GetSolo returns the Solo field if non-nil, zero value otherwise.

### GetSoloOk

`func (o *SkyBlockProfileMemberRiftSlayerQuest) GetSoloOk() (*bool, bool)`

GetSoloOk returns a tuple with the Solo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolo

`func (o *SkyBlockProfileMemberRiftSlayerQuest) SetSolo(v bool)`

SetSolo sets Solo field to given value.


### GetSpawnTimestamp

`func (o *SkyBlockProfileMemberRiftSlayerQuest) GetSpawnTimestamp() int64`

GetSpawnTimestamp returns the SpawnTimestamp field if non-nil, zero value otherwise.

### GetSpawnTimestampOk

`func (o *SkyBlockProfileMemberRiftSlayerQuest) GetSpawnTimestampOk() (*int64, bool)`

GetSpawnTimestampOk returns a tuple with the SpawnTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpawnTimestamp

`func (o *SkyBlockProfileMemberRiftSlayerQuest) SetSpawnTimestamp(v int64)`

SetSpawnTimestamp sets SpawnTimestamp field to given value.

### HasSpawnTimestamp

`func (o *SkyBlockProfileMemberRiftSlayerQuest) HasSpawnTimestamp() bool`

HasSpawnTimestamp returns a boolean if a field has been set.

### GetStartTimestamp

`func (o *SkyBlockProfileMemberRiftSlayerQuest) GetStartTimestamp() int64`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *SkyBlockProfileMemberRiftSlayerQuest) GetStartTimestampOk() (*int64, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *SkyBlockProfileMemberRiftSlayerQuest) SetStartTimestamp(v int64)`

SetStartTimestamp sets StartTimestamp field to given value.


### GetTier

`func (o *SkyBlockProfileMemberRiftSlayerQuest) GetTier() int64`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *SkyBlockProfileMemberRiftSlayerQuest) GetTierOk() (*int64, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *SkyBlockProfileMemberRiftSlayerQuest) SetTier(v int64)`

SetTier sets Tier field to given value.


### GetType

`func (o *SkyBlockProfileMemberRiftSlayerQuest) GetType() SkyBlockProfileMemberRiftSlayerQuestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SkyBlockProfileMemberRiftSlayerQuest) GetTypeOk() (*SkyBlockProfileMemberRiftSlayerQuestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SkyBlockProfileMemberRiftSlayerQuest) SetType(v SkyBlockProfileMemberRiftSlayerQuestType)`

SetType sets Type field to given value.


### GetUsedArmor

`func (o *SkyBlockProfileMemberRiftSlayerQuest) GetUsedArmor() bool`

GetUsedArmor returns the UsedArmor field if non-nil, zero value otherwise.

### GetUsedArmorOk

`func (o *SkyBlockProfileMemberRiftSlayerQuest) GetUsedArmorOk() (*bool, bool)`

GetUsedArmorOk returns a tuple with the UsedArmor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedArmor

`func (o *SkyBlockProfileMemberRiftSlayerQuest) SetUsedArmor(v bool)`

SetUsedArmor sets UsedArmor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


