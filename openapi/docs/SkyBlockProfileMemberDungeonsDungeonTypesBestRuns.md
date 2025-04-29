# SkyBlockProfileMemberDungeonsDungeonTypesBestRuns

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

### NewSkyBlockProfileMemberDungeonsDungeonTypesBestRuns

`func NewSkyBlockProfileMemberDungeonsDungeonTypesBestRuns(damageDealt float64, deaths int64, dungeonClass SkyBlockProfileMemberDungeonsSelectedDungeonClass, elapsedTime int64, mobsKilled int64, scoreBonus int64, scoreExploration int64, scoreSkill int64, scoreSpeed int64, secretsFound int64, teammates []string, timestamp int64, ) *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns`

NewSkyBlockProfileMemberDungeonsDungeonTypesBestRuns instantiates a new SkyBlockProfileMemberDungeonsDungeonTypesBestRuns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkyBlockProfileMemberDungeonsDungeonTypesBestRunsWithDefaults

`func NewSkyBlockProfileMemberDungeonsDungeonTypesBestRunsWithDefaults() *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns`

NewSkyBlockProfileMemberDungeonsDungeonTypesBestRunsWithDefaults instantiates a new SkyBlockProfileMemberDungeonsDungeonTypesBestRuns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllyHealing

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) GetAllyHealing() float64`

GetAllyHealing returns the AllyHealing field if non-nil, zero value otherwise.

### GetAllyHealingOk

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) GetAllyHealingOk() (*float64, bool)`

GetAllyHealingOk returns a tuple with the AllyHealing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllyHealing

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) SetAllyHealing(v float64)`

SetAllyHealing sets AllyHealing field to given value.

### HasAllyHealing

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) HasAllyHealing() bool`

HasAllyHealing returns a boolean if a field has been set.

### GetDamageDealt

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) GetDamageDealt() float64`

GetDamageDealt returns the DamageDealt field if non-nil, zero value otherwise.

### GetDamageDealtOk

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) GetDamageDealtOk() (*float64, bool)`

GetDamageDealtOk returns a tuple with the DamageDealt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamageDealt

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) SetDamageDealt(v float64)`

SetDamageDealt sets DamageDealt field to given value.


### GetDamageMitigated

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) GetDamageMitigated() float64`

GetDamageMitigated returns the DamageMitigated field if non-nil, zero value otherwise.

### GetDamageMitigatedOk

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) GetDamageMitigatedOk() (*float64, bool)`

GetDamageMitigatedOk returns a tuple with the DamageMitigated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamageMitigated

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) SetDamageMitigated(v float64)`

SetDamageMitigated sets DamageMitigated field to given value.

### HasDamageMitigated

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) HasDamageMitigated() bool`

HasDamageMitigated returns a boolean if a field has been set.

### GetDeaths

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) GetDeaths() int64`

GetDeaths returns the Deaths field if non-nil, zero value otherwise.

### GetDeathsOk

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) GetDeathsOk() (*int64, bool)`

GetDeathsOk returns a tuple with the Deaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeaths

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) SetDeaths(v int64)`

SetDeaths sets Deaths field to given value.


### GetDungeonClass

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) GetDungeonClass() SkyBlockProfileMemberDungeonsSelectedDungeonClass`

GetDungeonClass returns the DungeonClass field if non-nil, zero value otherwise.

### GetDungeonClassOk

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) GetDungeonClassOk() (*SkyBlockProfileMemberDungeonsSelectedDungeonClass, bool)`

GetDungeonClassOk returns a tuple with the DungeonClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDungeonClass

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) SetDungeonClass(v SkyBlockProfileMemberDungeonsSelectedDungeonClass)`

SetDungeonClass sets DungeonClass field to given value.


### GetElapsedTime

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) GetElapsedTime() int64`

GetElapsedTime returns the ElapsedTime field if non-nil, zero value otherwise.

### GetElapsedTimeOk

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) GetElapsedTimeOk() (*int64, bool)`

GetElapsedTimeOk returns a tuple with the ElapsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedTime

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) SetElapsedTime(v int64)`

SetElapsedTime sets ElapsedTime field to given value.


### GetMobsKilled

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) GetMobsKilled() int64`

GetMobsKilled returns the MobsKilled field if non-nil, zero value otherwise.

### GetMobsKilledOk

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) GetMobsKilledOk() (*int64, bool)`

GetMobsKilledOk returns a tuple with the MobsKilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobsKilled

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) SetMobsKilled(v int64)`

SetMobsKilled sets MobsKilled field to given value.


### GetScoreBonus

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) GetScoreBonus() int64`

GetScoreBonus returns the ScoreBonus field if non-nil, zero value otherwise.

### GetScoreBonusOk

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) GetScoreBonusOk() (*int64, bool)`

GetScoreBonusOk returns a tuple with the ScoreBonus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreBonus

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) SetScoreBonus(v int64)`

SetScoreBonus sets ScoreBonus field to given value.


### GetScoreExploration

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) GetScoreExploration() int64`

GetScoreExploration returns the ScoreExploration field if non-nil, zero value otherwise.

### GetScoreExplorationOk

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) GetScoreExplorationOk() (*int64, bool)`

GetScoreExplorationOk returns a tuple with the ScoreExploration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreExploration

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) SetScoreExploration(v int64)`

SetScoreExploration sets ScoreExploration field to given value.


### GetScoreSkill

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) GetScoreSkill() int64`

GetScoreSkill returns the ScoreSkill field if non-nil, zero value otherwise.

### GetScoreSkillOk

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) GetScoreSkillOk() (*int64, bool)`

GetScoreSkillOk returns a tuple with the ScoreSkill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreSkill

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) SetScoreSkill(v int64)`

SetScoreSkill sets ScoreSkill field to given value.


### GetScoreSpeed

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) GetScoreSpeed() int64`

GetScoreSpeed returns the ScoreSpeed field if non-nil, zero value otherwise.

### GetScoreSpeedOk

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) GetScoreSpeedOk() (*int64, bool)`

GetScoreSpeedOk returns a tuple with the ScoreSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreSpeed

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) SetScoreSpeed(v int64)`

SetScoreSpeed sets ScoreSpeed field to given value.


### GetSecretsFound

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) GetSecretsFound() int64`

GetSecretsFound returns the SecretsFound field if non-nil, zero value otherwise.

### GetSecretsFoundOk

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) GetSecretsFoundOk() (*int64, bool)`

GetSecretsFoundOk returns a tuple with the SecretsFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsFound

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) SetSecretsFound(v int64)`

SetSecretsFound sets SecretsFound field to given value.


### GetTeammates

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) GetTeammates() []string`

GetTeammates returns the Teammates field if non-nil, zero value otherwise.

### GetTeammatesOk

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) GetTeammatesOk() (*[]string, bool)`

GetTeammatesOk returns a tuple with the Teammates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeammates

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) SetTeammates(v []string)`

SetTeammates sets Teammates field to given value.


### GetTimestamp

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SkyBlockProfileMemberDungeonsDungeonTypesBestRuns) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


