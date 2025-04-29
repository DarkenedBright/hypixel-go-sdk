# SkyBlockProfileMemberDungeonsDungeonTypesBestRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllyHealing** | Pointer to **float64** |  | [optional] 
**DamageDealt** | **float64** |  | 
**DamageMitigated** | Pointer to **float64** |  | [optional] 
**Deaths** | **int64** |  | 
**DungeonClass** | [**SkyBlockProfileMemberDungeonsSelectedDungeonClass**](SkyBlockProfileMemberDungeonsSelectedDungeonClass.md) |  | 
**ElapsedTime** | **int64** |  | 
**MobsKilled** | **int64** |  | 
**ScoreBonus** | **int64** |  | 
**ScoreExploration** | **int64** |  | 
**ScoreSkill** | **int64** |  | 
**ScoreSpeed** | **int64** |  | 
**SecretsFound** | **int64** |  | 
**Teammates** | **[]string** |  | 
**Timestamp** | **int64** |  | 

## Methods

### NewSkyBlockProfileMemberDungeonsDungeonTypesBestRun

`func NewSkyBlockProfileMemberDungeonsDungeonTypesBestRun(damageDealt float64, deaths int64, dungeonClass SkyBlockProfileMemberDungeonsSelectedDungeonClass, elapsedTime int64, mobsKilled int64, scoreBonus int64, scoreExploration int64, scoreSkill int64, scoreSpeed int64, secretsFound int64, teammates []string, timestamp int64, ) *SkyBlockProfileMemberDungeonsDungeonTypesBestRun`

NewSkyBlockProfileMemberDungeonsDungeonTypesBestRun instantiates a new SkyBlockProfileMemberDungeonsDungeonTypesBestRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkyBlockProfileMemberDungeonsDungeonTypesBestRunWithDefaults

`func NewSkyBlockProfileMemberDungeonsDungeonTypesBestRunWithDefaults() *SkyBlockProfileMemberDungeonsDungeonTypesBestRun`

NewSkyBlockProfileMemberDungeonsDungeonTypesBestRunWithDefaults instantiates a new SkyBlockProfileMemberDungeonsDungeonTypesBestRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllyHealing

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) GetAllyHealing() float64`

GetAllyHealing returns the AllyHealing field if non-nil, zero value otherwise.

### GetAllyHealingOk

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) GetAllyHealingOk() (*float64, bool)`

GetAllyHealingOk returns a tuple with the AllyHealing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllyHealing

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) SetAllyHealing(v float64)`

SetAllyHealing sets AllyHealing field to given value.

### HasAllyHealing

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) HasAllyHealing() bool`

HasAllyHealing returns a boolean if a field has been set.

### GetDamageDealt

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) GetDamageDealt() float64`

GetDamageDealt returns the DamageDealt field if non-nil, zero value otherwise.

### GetDamageDealtOk

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) GetDamageDealtOk() (*float64, bool)`

GetDamageDealtOk returns a tuple with the DamageDealt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamageDealt

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) SetDamageDealt(v float64)`

SetDamageDealt sets DamageDealt field to given value.


### GetDamageMitigated

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) GetDamageMitigated() float64`

GetDamageMitigated returns the DamageMitigated field if non-nil, zero value otherwise.

### GetDamageMitigatedOk

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) GetDamageMitigatedOk() (*float64, bool)`

GetDamageMitigatedOk returns a tuple with the DamageMitigated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamageMitigated

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) SetDamageMitigated(v float64)`

SetDamageMitigated sets DamageMitigated field to given value.

### HasDamageMitigated

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) HasDamageMitigated() bool`

HasDamageMitigated returns a boolean if a field has been set.

### GetDeaths

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) GetDeaths() int64`

GetDeaths returns the Deaths field if non-nil, zero value otherwise.

### GetDeathsOk

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) GetDeathsOk() (*int64, bool)`

GetDeathsOk returns a tuple with the Deaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeaths

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) SetDeaths(v int64)`

SetDeaths sets Deaths field to given value.


### GetDungeonClass

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) GetDungeonClass() SkyBlockProfileMemberDungeonsSelectedDungeonClass`

GetDungeonClass returns the DungeonClass field if non-nil, zero value otherwise.

### GetDungeonClassOk

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) GetDungeonClassOk() (*SkyBlockProfileMemberDungeonsSelectedDungeonClass, bool)`

GetDungeonClassOk returns a tuple with the DungeonClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDungeonClass

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) SetDungeonClass(v SkyBlockProfileMemberDungeonsSelectedDungeonClass)`

SetDungeonClass sets DungeonClass field to given value.


### GetElapsedTime

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) GetElapsedTime() int64`

GetElapsedTime returns the ElapsedTime field if non-nil, zero value otherwise.

### GetElapsedTimeOk

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) GetElapsedTimeOk() (*int64, bool)`

GetElapsedTimeOk returns a tuple with the ElapsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedTime

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) SetElapsedTime(v int64)`

SetElapsedTime sets ElapsedTime field to given value.


### GetMobsKilled

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) GetMobsKilled() int64`

GetMobsKilled returns the MobsKilled field if non-nil, zero value otherwise.

### GetMobsKilledOk

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) GetMobsKilledOk() (*int64, bool)`

GetMobsKilledOk returns a tuple with the MobsKilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobsKilled

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) SetMobsKilled(v int64)`

SetMobsKilled sets MobsKilled field to given value.


### GetScoreBonus

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) GetScoreBonus() int64`

GetScoreBonus returns the ScoreBonus field if non-nil, zero value otherwise.

### GetScoreBonusOk

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) GetScoreBonusOk() (*int64, bool)`

GetScoreBonusOk returns a tuple with the ScoreBonus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreBonus

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) SetScoreBonus(v int64)`

SetScoreBonus sets ScoreBonus field to given value.


### GetScoreExploration

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) GetScoreExploration() int64`

GetScoreExploration returns the ScoreExploration field if non-nil, zero value otherwise.

### GetScoreExplorationOk

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) GetScoreExplorationOk() (*int64, bool)`

GetScoreExplorationOk returns a tuple with the ScoreExploration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreExploration

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) SetScoreExploration(v int64)`

SetScoreExploration sets ScoreExploration field to given value.


### GetScoreSkill

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) GetScoreSkill() int64`

GetScoreSkill returns the ScoreSkill field if non-nil, zero value otherwise.

### GetScoreSkillOk

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) GetScoreSkillOk() (*int64, bool)`

GetScoreSkillOk returns a tuple with the ScoreSkill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreSkill

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) SetScoreSkill(v int64)`

SetScoreSkill sets ScoreSkill field to given value.


### GetScoreSpeed

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) GetScoreSpeed() int64`

GetScoreSpeed returns the ScoreSpeed field if non-nil, zero value otherwise.

### GetScoreSpeedOk

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) GetScoreSpeedOk() (*int64, bool)`

GetScoreSpeedOk returns a tuple with the ScoreSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreSpeed

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) SetScoreSpeed(v int64)`

SetScoreSpeed sets ScoreSpeed field to given value.


### GetSecretsFound

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) GetSecretsFound() int64`

GetSecretsFound returns the SecretsFound field if non-nil, zero value otherwise.

### GetSecretsFoundOk

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) GetSecretsFoundOk() (*int64, bool)`

GetSecretsFoundOk returns a tuple with the SecretsFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsFound

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) SetSecretsFound(v int64)`

SetSecretsFound sets SecretsFound field to given value.


### GetTeammates

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) GetTeammates() []string`

GetTeammates returns the Teammates field if non-nil, zero value otherwise.

### GetTeammatesOk

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) GetTeammatesOk() (*[]string, bool)`

GetTeammatesOk returns a tuple with the Teammates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeammates

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) SetTeammates(v []string)`

SetTeammates sets Teammates field to given value.


### GetTimestamp

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRun) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


