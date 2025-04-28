# SkyBlockProfileMemberSlayerSlayerQuest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CombatXp** | Pointer to **int64** |  | [optional] 
**CompletionState** | **int64** |  | 
**KillTimestamp** | Pointer to **int64** |  | [optional] 
**LastKilledMobIsland** | Pointer to [**SkyBlockProfileMemberSlayerSlayerQuestLastKilledMobIsland**](SkyBlockProfileMemberSlayerSlayerQuestLastKilledMobIsland.md) |  | [optional] 
**RecentMobKills** | Pointer to [**[]SkyBlockProfileMemberSlayerSlayerQuestRecentMobKill**](SkyBlockProfileMemberSlayerSlayerQuestRecentMobKill.md) |  | [optional] 
**Solo** | Pointer to **bool** |  | [optional] 
**SpawnTimestamp** | Pointer to **int64** |  | [optional] 
**StartTimestamp** | **int64** |  | 
**Tier** | **int64** |  | 
**Type** | [**SkyBlockProfileMemberSlayerSlayerQuestType**](SkyBlockProfileMemberSlayerSlayerQuestType.md) |  | 
**UsedArmor** | Pointer to **bool** |  | [optional] 
**XpOnLastFollowerSpawn** | Pointer to **int64** |  | [optional] 

## Methods

### NewSkyBlockProfileMemberSlayerSlayerQuest

`func NewSkyBlockProfileMemberSlayerSlayerQuest(completionState int64, startTimestamp int64, tier int64, type_ SkyBlockProfileMemberSlayerSlayerQuestType, ) *SkyBlockProfileMemberSlayerSlayerQuest`

NewSkyBlockProfileMemberSlayerSlayerQuest instantiates a new SkyBlockProfileMemberSlayerSlayerQuest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkyBlockProfileMemberSlayerSlayerQuestWithDefaults

`func NewSkyBlockProfileMemberSlayerSlayerQuestWithDefaults() *SkyBlockProfileMemberSlayerSlayerQuest`

NewSkyBlockProfileMemberSlayerSlayerQuestWithDefaults instantiates a new SkyBlockProfileMemberSlayerSlayerQuest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCombatXp

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) GetCombatXp() int64`

GetCombatXp returns the CombatXp field if non-nil, zero value otherwise.

### GetCombatXpOk

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) GetCombatXpOk() (*int64, bool)`

GetCombatXpOk returns a tuple with the CombatXp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombatXp

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) SetCombatXp(v int64)`

SetCombatXp sets CombatXp field to given value.

### HasCombatXp

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) HasCombatXp() bool`

HasCombatXp returns a boolean if a field has been set.

### GetCompletionState

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) GetCompletionState() int64`

GetCompletionState returns the CompletionState field if non-nil, zero value otherwise.

### GetCompletionStateOk

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) GetCompletionStateOk() (*int64, bool)`

GetCompletionStateOk returns a tuple with the CompletionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionState

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) SetCompletionState(v int64)`

SetCompletionState sets CompletionState field to given value.


### GetKillTimestamp

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) GetKillTimestamp() int64`

GetKillTimestamp returns the KillTimestamp field if non-nil, zero value otherwise.

### GetKillTimestampOk

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) GetKillTimestampOk() (*int64, bool)`

GetKillTimestampOk returns a tuple with the KillTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillTimestamp

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) SetKillTimestamp(v int64)`

SetKillTimestamp sets KillTimestamp field to given value.

### HasKillTimestamp

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) HasKillTimestamp() bool`

HasKillTimestamp returns a boolean if a field has been set.

### GetLastKilledMobIsland

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) GetLastKilledMobIsland() SkyBlockProfileMemberSlayerSlayerQuestLastKilledMobIsland`

GetLastKilledMobIsland returns the LastKilledMobIsland field if non-nil, zero value otherwise.

### GetLastKilledMobIslandOk

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) GetLastKilledMobIslandOk() (*SkyBlockProfileMemberSlayerSlayerQuestLastKilledMobIsland, bool)`

GetLastKilledMobIslandOk returns a tuple with the LastKilledMobIsland field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastKilledMobIsland

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) SetLastKilledMobIsland(v SkyBlockProfileMemberSlayerSlayerQuestLastKilledMobIsland)`

SetLastKilledMobIsland sets LastKilledMobIsland field to given value.

### HasLastKilledMobIsland

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) HasLastKilledMobIsland() bool`

HasLastKilledMobIsland returns a boolean if a field has been set.

### GetRecentMobKills

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) GetRecentMobKills() []SkyBlockProfileMemberSlayerSlayerQuestRecentMobKill`

GetRecentMobKills returns the RecentMobKills field if non-nil, zero value otherwise.

### GetRecentMobKillsOk

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) GetRecentMobKillsOk() (*[]SkyBlockProfileMemberSlayerSlayerQuestRecentMobKill, bool)`

GetRecentMobKillsOk returns a tuple with the RecentMobKills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentMobKills

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) SetRecentMobKills(v []SkyBlockProfileMemberSlayerSlayerQuestRecentMobKill)`

SetRecentMobKills sets RecentMobKills field to given value.

### HasRecentMobKills

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) HasRecentMobKills() bool`

HasRecentMobKills returns a boolean if a field has been set.

### GetSolo

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) GetSolo() bool`

GetSolo returns the Solo field if non-nil, zero value otherwise.

### GetSoloOk

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) GetSoloOk() (*bool, bool)`

GetSoloOk returns a tuple with the Solo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolo

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) SetSolo(v bool)`

SetSolo sets Solo field to given value.

### HasSolo

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) HasSolo() bool`

HasSolo returns a boolean if a field has been set.

### GetSpawnTimestamp

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) GetSpawnTimestamp() int64`

GetSpawnTimestamp returns the SpawnTimestamp field if non-nil, zero value otherwise.

### GetSpawnTimestampOk

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) GetSpawnTimestampOk() (*int64, bool)`

GetSpawnTimestampOk returns a tuple with the SpawnTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpawnTimestamp

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) SetSpawnTimestamp(v int64)`

SetSpawnTimestamp sets SpawnTimestamp field to given value.

### HasSpawnTimestamp

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) HasSpawnTimestamp() bool`

HasSpawnTimestamp returns a boolean if a field has been set.

### GetStartTimestamp

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) GetStartTimestamp() int64`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) GetStartTimestampOk() (*int64, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) SetStartTimestamp(v int64)`

SetStartTimestamp sets StartTimestamp field to given value.


### GetTier

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) GetTier() int64`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) GetTierOk() (*int64, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) SetTier(v int64)`

SetTier sets Tier field to given value.


### GetType

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) GetType() SkyBlockProfileMemberSlayerSlayerQuestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) GetTypeOk() (*SkyBlockProfileMemberSlayerSlayerQuestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) SetType(v SkyBlockProfileMemberSlayerSlayerQuestType)`

SetType sets Type field to given value.


### GetUsedArmor

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) GetUsedArmor() bool`

GetUsedArmor returns the UsedArmor field if non-nil, zero value otherwise.

### GetUsedArmorOk

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) GetUsedArmorOk() (*bool, bool)`

GetUsedArmorOk returns a tuple with the UsedArmor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedArmor

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) SetUsedArmor(v bool)`

SetUsedArmor sets UsedArmor field to given value.

### HasUsedArmor

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) HasUsedArmor() bool`

HasUsedArmor returns a boolean if a field has been set.

### GetXpOnLastFollowerSpawn

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) GetXpOnLastFollowerSpawn() int64`

GetXpOnLastFollowerSpawn returns the XpOnLastFollowerSpawn field if non-nil, zero value otherwise.

### GetXpOnLastFollowerSpawnOk

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) GetXpOnLastFollowerSpawnOk() (*int64, bool)`

GetXpOnLastFollowerSpawnOk returns a tuple with the XpOnLastFollowerSpawn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXpOnLastFollowerSpawn

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) SetXpOnLastFollowerSpawn(v int64)`

SetXpOnLastFollowerSpawn sets XpOnLastFollowerSpawn field to given value.

### HasXpOnLastFollowerSpawn

`func (o *SkyBlockProfileMemberSlayerSlayerQuest) HasXpOnLastFollowerSpawn() bool`

HasXpOnLastFollowerSpawn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


