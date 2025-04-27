# SkyBlockProfileMemberNetherIslandPlayerQuests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlchemistQuest** | [**SkyBlockProfileMemberNetherIslandPlayerQuestsAlchemistQuest**](SkyBlockProfileMemberNetherIslandPlayerQuestsAlchemistQuest.md) |  | 
**AranyaQuest** | **map[string]interface{}** |  | 
**CavityRarity** | Pointer to [**SkyBlockProfileMemberNetherIslandPlayerQuestsCavityRarity**](SkyBlockProfileMemberNetherIslandPlayerQuestsCavityRarity.md) |  | [optional] 
**ChickenQuest** | [**SkyBlockProfileMemberNetherIslandPlayerQuestsChickenQuest**](SkyBlockProfileMemberNetherIslandPlayerQuestsChickenQuest.md) |  | 
**ChickenQuestHandedIn** | **int64** |  | 
**DuelTrainingQuest** | [**SkyBlockProfileMemberNetherIslandPlayerQuestsDuelTrainingQuest**](SkyBlockProfileMemberNetherIslandPlayerQuestsDuelTrainingQuest.md) |  | 
**EdelisQuest** | [**SkyBlockProfileMemberNetherIslandPlayerQuestsEdelisQuest**](SkyBlockProfileMemberNetherIslandPlayerQuestsEdelisQuest.md) |  | 
**FishedWetNapkin** | Pointer to **bool** |  | [optional] 
**FoundKuudraBook** | Pointer to **bool** |  | [optional] 
**FoundKuudraBoots** | Pointer to **bool** |  | [optional] 
**FoundKuudraChestplate** | Pointer to **bool** |  | [optional] 
**FoundKuudraHelmet** | Pointer to **bool** |  | [optional] 
**FoundKuudraLeggings** | Pointer to **bool** |  | [optional] 
**KuudaBossDaily** | **map[string]interface{}** |  | 
**KuudraLoremaster** | Pointer to **bool** |  | [optional] 
**LastBelieverBlessing** | Pointer to **int64** |  | [optional] 
**LastKuudraRelic** | Pointer to **int64** |  | [optional] 
**LastReset** | **int64** |  | 
**LastVampireBlood** | Pointer to **int64** |  | [optional] 
**MinibossDaily** | **map[string]interface{}** |  | 
**MinibossData** | [**SkyBlockProfileMemberNetherIslandPlayerQuestsMinibossData**](SkyBlockProfileMemberNetherIslandPlayerQuestsMinibossData.md) |  | 
**MollimQuest** | [**SkyBlockProfileMemberNetherIslandPlayerQuestsMollimQuest**](SkyBlockProfileMemberNetherIslandPlayerQuestsMollimQuest.md) |  | 
**PabloQuest** | [**SkyBlockProfileMemberNetherIslandPlayerQuestsPabloQuest**](SkyBlockProfileMemberNetherIslandPlayerQuestsPabloQuest.md) |  | 
**PaidBruuh** | **bool** |  | 
**PomtairQuest** | [**SkyBlockProfileMemberNetherIslandPlayerQuestsPomtairQuest**](SkyBlockProfileMemberNetherIslandPlayerQuestsPomtairQuest.md) |  | 
**QuestData** | [**SkyBlockProfileMemberNetherIslandPlayerQuestsQuestData**](SkyBlockProfileMemberNetherIslandPlayerQuestsQuestData.md) |  | 
**QuestRewards** | [**SkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewards**](SkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewards.md) |  | 
**Rulenor** | [**SkyBlockProfileMemberNetherIslandPlayerQuestsRulenor**](SkyBlockProfileMemberNetherIslandPlayerQuestsRulenor.md) |  | 
**SirihQuest** | [**SkyBlockProfileMemberNetherIslandPlayerQuestsSirihQuest**](SkyBlockProfileMemberNetherIslandPlayerQuestsSirihQuest.md) |  | 
**SuusQuest** | [**SkyBlockProfileMemberNetherIslandPlayerQuestsSuusQuest**](SkyBlockProfileMemberNetherIslandPlayerQuestsSuusQuest.md) |  | 
**UnlockedCavityNpcs** | Pointer to [**[]SkyBlockProfileMemberNetherIslandPlayerQuestsUnlockedCavityNpcsInner**](SkyBlockProfileMemberNetherIslandPlayerQuestsUnlockedCavityNpcsInner.md) |  | [optional] 
**WeirdSailor** | Pointer to **bool** |  | [optional] 

## Methods

### NewSkyBlockProfileMemberNetherIslandPlayerQuests

`func NewSkyBlockProfileMemberNetherIslandPlayerQuests(alchemistQuest SkyBlockProfileMemberNetherIslandPlayerQuestsAlchemistQuest, aranyaQuest map[string]interface{}, chickenQuest SkyBlockProfileMemberNetherIslandPlayerQuestsChickenQuest, chickenQuestHandedIn int64, duelTrainingQuest SkyBlockProfileMemberNetherIslandPlayerQuestsDuelTrainingQuest, edelisQuest SkyBlockProfileMemberNetherIslandPlayerQuestsEdelisQuest, kuudaBossDaily map[string]interface{}, lastReset int64, minibossDaily map[string]interface{}, minibossData SkyBlockProfileMemberNetherIslandPlayerQuestsMinibossData, mollimQuest SkyBlockProfileMemberNetherIslandPlayerQuestsMollimQuest, pabloQuest SkyBlockProfileMemberNetherIslandPlayerQuestsPabloQuest, paidBruuh bool, pomtairQuest SkyBlockProfileMemberNetherIslandPlayerQuestsPomtairQuest, questData SkyBlockProfileMemberNetherIslandPlayerQuestsQuestData, questRewards SkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewards, rulenor SkyBlockProfileMemberNetherIslandPlayerQuestsRulenor, sirihQuest SkyBlockProfileMemberNetherIslandPlayerQuestsSirihQuest, suusQuest SkyBlockProfileMemberNetherIslandPlayerQuestsSuusQuest, ) *SkyBlockProfileMemberNetherIslandPlayerQuests`

NewSkyBlockProfileMemberNetherIslandPlayerQuests instantiates a new SkyBlockProfileMemberNetherIslandPlayerQuests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkyBlockProfileMemberNetherIslandPlayerQuestsWithDefaults

`func NewSkyBlockProfileMemberNetherIslandPlayerQuestsWithDefaults() *SkyBlockProfileMemberNetherIslandPlayerQuests`

NewSkyBlockProfileMemberNetherIslandPlayerQuestsWithDefaults instantiates a new SkyBlockProfileMemberNetherIslandPlayerQuests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlchemistQuest

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetAlchemistQuest() SkyBlockProfileMemberNetherIslandPlayerQuestsAlchemistQuest`

GetAlchemistQuest returns the AlchemistQuest field if non-nil, zero value otherwise.

### GetAlchemistQuestOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetAlchemistQuestOk() (*SkyBlockProfileMemberNetherIslandPlayerQuestsAlchemistQuest, bool)`

GetAlchemistQuestOk returns a tuple with the AlchemistQuest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlchemistQuest

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetAlchemistQuest(v SkyBlockProfileMemberNetherIslandPlayerQuestsAlchemistQuest)`

SetAlchemistQuest sets AlchemistQuest field to given value.


### GetAranyaQuest

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetAranyaQuest() map[string]interface{}`

GetAranyaQuest returns the AranyaQuest field if non-nil, zero value otherwise.

### GetAranyaQuestOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetAranyaQuestOk() (*map[string]interface{}, bool)`

GetAranyaQuestOk returns a tuple with the AranyaQuest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAranyaQuest

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetAranyaQuest(v map[string]interface{})`

SetAranyaQuest sets AranyaQuest field to given value.


### GetCavityRarity

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetCavityRarity() SkyBlockProfileMemberNetherIslandPlayerQuestsCavityRarity`

GetCavityRarity returns the CavityRarity field if non-nil, zero value otherwise.

### GetCavityRarityOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetCavityRarityOk() (*SkyBlockProfileMemberNetherIslandPlayerQuestsCavityRarity, bool)`

GetCavityRarityOk returns a tuple with the CavityRarity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCavityRarity

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetCavityRarity(v SkyBlockProfileMemberNetherIslandPlayerQuestsCavityRarity)`

SetCavityRarity sets CavityRarity field to given value.

### HasCavityRarity

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) HasCavityRarity() bool`

HasCavityRarity returns a boolean if a field has been set.

### GetChickenQuest

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetChickenQuest() SkyBlockProfileMemberNetherIslandPlayerQuestsChickenQuest`

GetChickenQuest returns the ChickenQuest field if non-nil, zero value otherwise.

### GetChickenQuestOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetChickenQuestOk() (*SkyBlockProfileMemberNetherIslandPlayerQuestsChickenQuest, bool)`

GetChickenQuestOk returns a tuple with the ChickenQuest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChickenQuest

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetChickenQuest(v SkyBlockProfileMemberNetherIslandPlayerQuestsChickenQuest)`

SetChickenQuest sets ChickenQuest field to given value.


### GetChickenQuestHandedIn

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetChickenQuestHandedIn() int64`

GetChickenQuestHandedIn returns the ChickenQuestHandedIn field if non-nil, zero value otherwise.

### GetChickenQuestHandedInOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetChickenQuestHandedInOk() (*int64, bool)`

GetChickenQuestHandedInOk returns a tuple with the ChickenQuestHandedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChickenQuestHandedIn

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetChickenQuestHandedIn(v int64)`

SetChickenQuestHandedIn sets ChickenQuestHandedIn field to given value.


### GetDuelTrainingQuest

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetDuelTrainingQuest() SkyBlockProfileMemberNetherIslandPlayerQuestsDuelTrainingQuest`

GetDuelTrainingQuest returns the DuelTrainingQuest field if non-nil, zero value otherwise.

### GetDuelTrainingQuestOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetDuelTrainingQuestOk() (*SkyBlockProfileMemberNetherIslandPlayerQuestsDuelTrainingQuest, bool)`

GetDuelTrainingQuestOk returns a tuple with the DuelTrainingQuest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuelTrainingQuest

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetDuelTrainingQuest(v SkyBlockProfileMemberNetherIslandPlayerQuestsDuelTrainingQuest)`

SetDuelTrainingQuest sets DuelTrainingQuest field to given value.


### GetEdelisQuest

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetEdelisQuest() SkyBlockProfileMemberNetherIslandPlayerQuestsEdelisQuest`

GetEdelisQuest returns the EdelisQuest field if non-nil, zero value otherwise.

### GetEdelisQuestOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetEdelisQuestOk() (*SkyBlockProfileMemberNetherIslandPlayerQuestsEdelisQuest, bool)`

GetEdelisQuestOk returns a tuple with the EdelisQuest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdelisQuest

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetEdelisQuest(v SkyBlockProfileMemberNetherIslandPlayerQuestsEdelisQuest)`

SetEdelisQuest sets EdelisQuest field to given value.


### GetFishedWetNapkin

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetFishedWetNapkin() bool`

GetFishedWetNapkin returns the FishedWetNapkin field if non-nil, zero value otherwise.

### GetFishedWetNapkinOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetFishedWetNapkinOk() (*bool, bool)`

GetFishedWetNapkinOk returns a tuple with the FishedWetNapkin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFishedWetNapkin

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetFishedWetNapkin(v bool)`

SetFishedWetNapkin sets FishedWetNapkin field to given value.

### HasFishedWetNapkin

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) HasFishedWetNapkin() bool`

HasFishedWetNapkin returns a boolean if a field has been set.

### GetFoundKuudraBook

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetFoundKuudraBook() bool`

GetFoundKuudraBook returns the FoundKuudraBook field if non-nil, zero value otherwise.

### GetFoundKuudraBookOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetFoundKuudraBookOk() (*bool, bool)`

GetFoundKuudraBookOk returns a tuple with the FoundKuudraBook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundKuudraBook

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetFoundKuudraBook(v bool)`

SetFoundKuudraBook sets FoundKuudraBook field to given value.

### HasFoundKuudraBook

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) HasFoundKuudraBook() bool`

HasFoundKuudraBook returns a boolean if a field has been set.

### GetFoundKuudraBoots

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetFoundKuudraBoots() bool`

GetFoundKuudraBoots returns the FoundKuudraBoots field if non-nil, zero value otherwise.

### GetFoundKuudraBootsOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetFoundKuudraBootsOk() (*bool, bool)`

GetFoundKuudraBootsOk returns a tuple with the FoundKuudraBoots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundKuudraBoots

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetFoundKuudraBoots(v bool)`

SetFoundKuudraBoots sets FoundKuudraBoots field to given value.

### HasFoundKuudraBoots

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) HasFoundKuudraBoots() bool`

HasFoundKuudraBoots returns a boolean if a field has been set.

### GetFoundKuudraChestplate

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetFoundKuudraChestplate() bool`

GetFoundKuudraChestplate returns the FoundKuudraChestplate field if non-nil, zero value otherwise.

### GetFoundKuudraChestplateOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetFoundKuudraChestplateOk() (*bool, bool)`

GetFoundKuudraChestplateOk returns a tuple with the FoundKuudraChestplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundKuudraChestplate

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetFoundKuudraChestplate(v bool)`

SetFoundKuudraChestplate sets FoundKuudraChestplate field to given value.

### HasFoundKuudraChestplate

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) HasFoundKuudraChestplate() bool`

HasFoundKuudraChestplate returns a boolean if a field has been set.

### GetFoundKuudraHelmet

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetFoundKuudraHelmet() bool`

GetFoundKuudraHelmet returns the FoundKuudraHelmet field if non-nil, zero value otherwise.

### GetFoundKuudraHelmetOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetFoundKuudraHelmetOk() (*bool, bool)`

GetFoundKuudraHelmetOk returns a tuple with the FoundKuudraHelmet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundKuudraHelmet

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetFoundKuudraHelmet(v bool)`

SetFoundKuudraHelmet sets FoundKuudraHelmet field to given value.

### HasFoundKuudraHelmet

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) HasFoundKuudraHelmet() bool`

HasFoundKuudraHelmet returns a boolean if a field has been set.

### GetFoundKuudraLeggings

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetFoundKuudraLeggings() bool`

GetFoundKuudraLeggings returns the FoundKuudraLeggings field if non-nil, zero value otherwise.

### GetFoundKuudraLeggingsOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetFoundKuudraLeggingsOk() (*bool, bool)`

GetFoundKuudraLeggingsOk returns a tuple with the FoundKuudraLeggings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundKuudraLeggings

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetFoundKuudraLeggings(v bool)`

SetFoundKuudraLeggings sets FoundKuudraLeggings field to given value.

### HasFoundKuudraLeggings

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) HasFoundKuudraLeggings() bool`

HasFoundKuudraLeggings returns a boolean if a field has been set.

### GetKuudaBossDaily

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetKuudaBossDaily() map[string]interface{}`

GetKuudaBossDaily returns the KuudaBossDaily field if non-nil, zero value otherwise.

### GetKuudaBossDailyOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetKuudaBossDailyOk() (*map[string]interface{}, bool)`

GetKuudaBossDailyOk returns a tuple with the KuudaBossDaily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKuudaBossDaily

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetKuudaBossDaily(v map[string]interface{})`

SetKuudaBossDaily sets KuudaBossDaily field to given value.


### GetKuudraLoremaster

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetKuudraLoremaster() bool`

GetKuudraLoremaster returns the KuudraLoremaster field if non-nil, zero value otherwise.

### GetKuudraLoremasterOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetKuudraLoremasterOk() (*bool, bool)`

GetKuudraLoremasterOk returns a tuple with the KuudraLoremaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKuudraLoremaster

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetKuudraLoremaster(v bool)`

SetKuudraLoremaster sets KuudraLoremaster field to given value.

### HasKuudraLoremaster

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) HasKuudraLoremaster() bool`

HasKuudraLoremaster returns a boolean if a field has been set.

### GetLastBelieverBlessing

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetLastBelieverBlessing() int64`

GetLastBelieverBlessing returns the LastBelieverBlessing field if non-nil, zero value otherwise.

### GetLastBelieverBlessingOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetLastBelieverBlessingOk() (*int64, bool)`

GetLastBelieverBlessingOk returns a tuple with the LastBelieverBlessing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBelieverBlessing

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetLastBelieverBlessing(v int64)`

SetLastBelieverBlessing sets LastBelieverBlessing field to given value.

### HasLastBelieverBlessing

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) HasLastBelieverBlessing() bool`

HasLastBelieverBlessing returns a boolean if a field has been set.

### GetLastKuudraRelic

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetLastKuudraRelic() int64`

GetLastKuudraRelic returns the LastKuudraRelic field if non-nil, zero value otherwise.

### GetLastKuudraRelicOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetLastKuudraRelicOk() (*int64, bool)`

GetLastKuudraRelicOk returns a tuple with the LastKuudraRelic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastKuudraRelic

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetLastKuudraRelic(v int64)`

SetLastKuudraRelic sets LastKuudraRelic field to given value.

### HasLastKuudraRelic

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) HasLastKuudraRelic() bool`

HasLastKuudraRelic returns a boolean if a field has been set.

### GetLastReset

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetLastReset() int64`

GetLastReset returns the LastReset field if non-nil, zero value otherwise.

### GetLastResetOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetLastResetOk() (*int64, bool)`

GetLastResetOk returns a tuple with the LastReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReset

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetLastReset(v int64)`

SetLastReset sets LastReset field to given value.


### GetLastVampireBlood

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetLastVampireBlood() int64`

GetLastVampireBlood returns the LastVampireBlood field if non-nil, zero value otherwise.

### GetLastVampireBloodOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetLastVampireBloodOk() (*int64, bool)`

GetLastVampireBloodOk returns a tuple with the LastVampireBlood field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVampireBlood

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetLastVampireBlood(v int64)`

SetLastVampireBlood sets LastVampireBlood field to given value.

### HasLastVampireBlood

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) HasLastVampireBlood() bool`

HasLastVampireBlood returns a boolean if a field has been set.

### GetMinibossDaily

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetMinibossDaily() map[string]interface{}`

GetMinibossDaily returns the MinibossDaily field if non-nil, zero value otherwise.

### GetMinibossDailyOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetMinibossDailyOk() (*map[string]interface{}, bool)`

GetMinibossDailyOk returns a tuple with the MinibossDaily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinibossDaily

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetMinibossDaily(v map[string]interface{})`

SetMinibossDaily sets MinibossDaily field to given value.


### GetMinibossData

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetMinibossData() SkyBlockProfileMemberNetherIslandPlayerQuestsMinibossData`

GetMinibossData returns the MinibossData field if non-nil, zero value otherwise.

### GetMinibossDataOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetMinibossDataOk() (*SkyBlockProfileMemberNetherIslandPlayerQuestsMinibossData, bool)`

GetMinibossDataOk returns a tuple with the MinibossData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinibossData

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetMinibossData(v SkyBlockProfileMemberNetherIslandPlayerQuestsMinibossData)`

SetMinibossData sets MinibossData field to given value.


### GetMollimQuest

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetMollimQuest() SkyBlockProfileMemberNetherIslandPlayerQuestsMollimQuest`

GetMollimQuest returns the MollimQuest field if non-nil, zero value otherwise.

### GetMollimQuestOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetMollimQuestOk() (*SkyBlockProfileMemberNetherIslandPlayerQuestsMollimQuest, bool)`

GetMollimQuestOk returns a tuple with the MollimQuest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMollimQuest

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetMollimQuest(v SkyBlockProfileMemberNetherIslandPlayerQuestsMollimQuest)`

SetMollimQuest sets MollimQuest field to given value.


### GetPabloQuest

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetPabloQuest() SkyBlockProfileMemberNetherIslandPlayerQuestsPabloQuest`

GetPabloQuest returns the PabloQuest field if non-nil, zero value otherwise.

### GetPabloQuestOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetPabloQuestOk() (*SkyBlockProfileMemberNetherIslandPlayerQuestsPabloQuest, bool)`

GetPabloQuestOk returns a tuple with the PabloQuest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPabloQuest

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetPabloQuest(v SkyBlockProfileMemberNetherIslandPlayerQuestsPabloQuest)`

SetPabloQuest sets PabloQuest field to given value.


### GetPaidBruuh

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetPaidBruuh() bool`

GetPaidBruuh returns the PaidBruuh field if non-nil, zero value otherwise.

### GetPaidBruuhOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetPaidBruuhOk() (*bool, bool)`

GetPaidBruuhOk returns a tuple with the PaidBruuh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidBruuh

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetPaidBruuh(v bool)`

SetPaidBruuh sets PaidBruuh field to given value.


### GetPomtairQuest

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetPomtairQuest() SkyBlockProfileMemberNetherIslandPlayerQuestsPomtairQuest`

GetPomtairQuest returns the PomtairQuest field if non-nil, zero value otherwise.

### GetPomtairQuestOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetPomtairQuestOk() (*SkyBlockProfileMemberNetherIslandPlayerQuestsPomtairQuest, bool)`

GetPomtairQuestOk returns a tuple with the PomtairQuest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPomtairQuest

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetPomtairQuest(v SkyBlockProfileMemberNetherIslandPlayerQuestsPomtairQuest)`

SetPomtairQuest sets PomtairQuest field to given value.


### GetQuestData

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetQuestData() SkyBlockProfileMemberNetherIslandPlayerQuestsQuestData`

GetQuestData returns the QuestData field if non-nil, zero value otherwise.

### GetQuestDataOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetQuestDataOk() (*SkyBlockProfileMemberNetherIslandPlayerQuestsQuestData, bool)`

GetQuestDataOk returns a tuple with the QuestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestData

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetQuestData(v SkyBlockProfileMemberNetherIslandPlayerQuestsQuestData)`

SetQuestData sets QuestData field to given value.


### GetQuestRewards

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetQuestRewards() SkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewards`

GetQuestRewards returns the QuestRewards field if non-nil, zero value otherwise.

### GetQuestRewardsOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetQuestRewardsOk() (*SkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewards, bool)`

GetQuestRewardsOk returns a tuple with the QuestRewards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestRewards

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetQuestRewards(v SkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewards)`

SetQuestRewards sets QuestRewards field to given value.


### GetRulenor

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetRulenor() SkyBlockProfileMemberNetherIslandPlayerQuestsRulenor`

GetRulenor returns the Rulenor field if non-nil, zero value otherwise.

### GetRulenorOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetRulenorOk() (*SkyBlockProfileMemberNetherIslandPlayerQuestsRulenor, bool)`

GetRulenorOk returns a tuple with the Rulenor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulenor

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetRulenor(v SkyBlockProfileMemberNetherIslandPlayerQuestsRulenor)`

SetRulenor sets Rulenor field to given value.


### GetSirihQuest

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetSirihQuest() SkyBlockProfileMemberNetherIslandPlayerQuestsSirihQuest`

GetSirihQuest returns the SirihQuest field if non-nil, zero value otherwise.

### GetSirihQuestOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetSirihQuestOk() (*SkyBlockProfileMemberNetherIslandPlayerQuestsSirihQuest, bool)`

GetSirihQuestOk returns a tuple with the SirihQuest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSirihQuest

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetSirihQuest(v SkyBlockProfileMemberNetherIslandPlayerQuestsSirihQuest)`

SetSirihQuest sets SirihQuest field to given value.


### GetSuusQuest

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetSuusQuest() SkyBlockProfileMemberNetherIslandPlayerQuestsSuusQuest`

GetSuusQuest returns the SuusQuest field if non-nil, zero value otherwise.

### GetSuusQuestOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetSuusQuestOk() (*SkyBlockProfileMemberNetherIslandPlayerQuestsSuusQuest, bool)`

GetSuusQuestOk returns a tuple with the SuusQuest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuusQuest

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetSuusQuest(v SkyBlockProfileMemberNetherIslandPlayerQuestsSuusQuest)`

SetSuusQuest sets SuusQuest field to given value.


### GetUnlockedCavityNpcs

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetUnlockedCavityNpcs() []SkyBlockProfileMemberNetherIslandPlayerQuestsUnlockedCavityNpcsInner`

GetUnlockedCavityNpcs returns the UnlockedCavityNpcs field if non-nil, zero value otherwise.

### GetUnlockedCavityNpcsOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetUnlockedCavityNpcsOk() (*[]SkyBlockProfileMemberNetherIslandPlayerQuestsUnlockedCavityNpcsInner, bool)`

GetUnlockedCavityNpcsOk returns a tuple with the UnlockedCavityNpcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockedCavityNpcs

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetUnlockedCavityNpcs(v []SkyBlockProfileMemberNetherIslandPlayerQuestsUnlockedCavityNpcsInner)`

SetUnlockedCavityNpcs sets UnlockedCavityNpcs field to given value.

### HasUnlockedCavityNpcs

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) HasUnlockedCavityNpcs() bool`

HasUnlockedCavityNpcs returns a boolean if a field has been set.

### GetWeirdSailor

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetWeirdSailor() bool`

GetWeirdSailor returns the WeirdSailor field if non-nil, zero value otherwise.

### GetWeirdSailorOk

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) GetWeirdSailorOk() (*bool, bool)`

GetWeirdSailorOk returns a tuple with the WeirdSailor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeirdSailor

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) SetWeirdSailor(v bool)`

SetWeirdSailor sets WeirdSailor field to given value.

### HasWeirdSailor

`func (o *SkyBlockProfileMemberNetherIslandPlayerQuests) HasWeirdSailor() bool`

HasWeirdSailor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


