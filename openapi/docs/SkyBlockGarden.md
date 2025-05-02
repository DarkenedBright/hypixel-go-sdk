# SkyBlockGarden

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveCommissions** | Pointer to [**map[string]SkyBlockGardenActiveCommissionsCommission**](SkyBlockGardenActiveCommissionsCommission.md) |  | [optional] 
**CommissionData** | [**SkyBlockGardenCommissionData**](SkyBlockGardenCommissionData.md) |  | 
**ComposterData** | [**SkyBlockGardenComposterData**](SkyBlockGardenComposterData.md) |  | 
**CropUpgradeLevels** | Pointer to [**SkyBlockGardenCropUpgradeLevels**](SkyBlockGardenCropUpgradeLevels.md) |  | [optional] 
**GardenExperience** | Pointer to **float64** |  | [optional] 
**ResourcesCollected** | Pointer to [**SkyBlockGardenResourcesCollected**](SkyBlockGardenResourcesCollected.md) |  | [optional] 
**SelectedBarnSkin** | Pointer to [**SkyBlockGardenBarnSkin**](SkyBlockGardenBarnSkin.md) |  | [optional] 
**UnlockedBarnSkins** | Pointer to [**[]SkyBlockGardenBarnSkin**](SkyBlockGardenBarnSkin.md) |  | [optional] 
**UnlockedPlotsIds** | [**[]SkyBlockGardenUnlockedPlotsIdsInner**](SkyBlockGardenUnlockedPlotsIdsInner.md) |  | 
**Uuid** | **string** |  | 

## Methods

### NewSkyBlockGarden

`func NewSkyBlockGarden(commissionData SkyBlockGardenCommissionData, composterData SkyBlockGardenComposterData, unlockedPlotsIds []SkyBlockGardenUnlockedPlotsIdsInner, uuid string, ) *SkyBlockGarden`

NewSkyBlockGarden instantiates a new SkyBlockGarden object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkyBlockGardenWithDefaults

`func NewSkyBlockGardenWithDefaults() *SkyBlockGarden`

NewSkyBlockGardenWithDefaults instantiates a new SkyBlockGarden object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveCommissions

`func (o *SkyBlockGarden) GetActiveCommissions() map[string]SkyBlockGardenActiveCommissionsCommission`

GetActiveCommissions returns the ActiveCommissions field if non-nil, zero value otherwise.

### GetActiveCommissionsOk

`func (o *SkyBlockGarden) GetActiveCommissionsOk() (*map[string]SkyBlockGardenActiveCommissionsCommission, bool)`

GetActiveCommissionsOk returns a tuple with the ActiveCommissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveCommissions

`func (o *SkyBlockGarden) SetActiveCommissions(v map[string]SkyBlockGardenActiveCommissionsCommission)`

SetActiveCommissions sets ActiveCommissions field to given value.

### HasActiveCommissions

`func (o *SkyBlockGarden) HasActiveCommissions() bool`

HasActiveCommissions returns a boolean if a field has been set.

### GetCommissionData

`func (o *SkyBlockGarden) GetCommissionData() SkyBlockGardenCommissionData`

GetCommissionData returns the CommissionData field if non-nil, zero value otherwise.

### GetCommissionDataOk

`func (o *SkyBlockGarden) GetCommissionDataOk() (*SkyBlockGardenCommissionData, bool)`

GetCommissionDataOk returns a tuple with the CommissionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionData

`func (o *SkyBlockGarden) SetCommissionData(v SkyBlockGardenCommissionData)`

SetCommissionData sets CommissionData field to given value.


### GetComposterData

`func (o *SkyBlockGarden) GetComposterData() SkyBlockGardenComposterData`

GetComposterData returns the ComposterData field if non-nil, zero value otherwise.

### GetComposterDataOk

`func (o *SkyBlockGarden) GetComposterDataOk() (*SkyBlockGardenComposterData, bool)`

GetComposterDataOk returns a tuple with the ComposterData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComposterData

`func (o *SkyBlockGarden) SetComposterData(v SkyBlockGardenComposterData)`

SetComposterData sets ComposterData field to given value.


### GetCropUpgradeLevels

`func (o *SkyBlockGarden) GetCropUpgradeLevels() SkyBlockGardenCropUpgradeLevels`

GetCropUpgradeLevels returns the CropUpgradeLevels field if non-nil, zero value otherwise.

### GetCropUpgradeLevelsOk

`func (o *SkyBlockGarden) GetCropUpgradeLevelsOk() (*SkyBlockGardenCropUpgradeLevels, bool)`

GetCropUpgradeLevelsOk returns a tuple with the CropUpgradeLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCropUpgradeLevels

`func (o *SkyBlockGarden) SetCropUpgradeLevels(v SkyBlockGardenCropUpgradeLevels)`

SetCropUpgradeLevels sets CropUpgradeLevels field to given value.

### HasCropUpgradeLevels

`func (o *SkyBlockGarden) HasCropUpgradeLevels() bool`

HasCropUpgradeLevels returns a boolean if a field has been set.

### GetGardenExperience

`func (o *SkyBlockGarden) GetGardenExperience() float64`

GetGardenExperience returns the GardenExperience field if non-nil, zero value otherwise.

### GetGardenExperienceOk

`func (o *SkyBlockGarden) GetGardenExperienceOk() (*float64, bool)`

GetGardenExperienceOk returns a tuple with the GardenExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGardenExperience

`func (o *SkyBlockGarden) SetGardenExperience(v float64)`

SetGardenExperience sets GardenExperience field to given value.

### HasGardenExperience

`func (o *SkyBlockGarden) HasGardenExperience() bool`

HasGardenExperience returns a boolean if a field has been set.

### GetResourcesCollected

`func (o *SkyBlockGarden) GetResourcesCollected() SkyBlockGardenResourcesCollected`

GetResourcesCollected returns the ResourcesCollected field if non-nil, zero value otherwise.

### GetResourcesCollectedOk

`func (o *SkyBlockGarden) GetResourcesCollectedOk() (*SkyBlockGardenResourcesCollected, bool)`

GetResourcesCollectedOk returns a tuple with the ResourcesCollected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesCollected

`func (o *SkyBlockGarden) SetResourcesCollected(v SkyBlockGardenResourcesCollected)`

SetResourcesCollected sets ResourcesCollected field to given value.

### HasResourcesCollected

`func (o *SkyBlockGarden) HasResourcesCollected() bool`

HasResourcesCollected returns a boolean if a field has been set.

### GetSelectedBarnSkin

`func (o *SkyBlockGarden) GetSelectedBarnSkin() SkyBlockGardenBarnSkin`

GetSelectedBarnSkin returns the SelectedBarnSkin field if non-nil, zero value otherwise.

### GetSelectedBarnSkinOk

`func (o *SkyBlockGarden) GetSelectedBarnSkinOk() (*SkyBlockGardenBarnSkin, bool)`

GetSelectedBarnSkinOk returns a tuple with the SelectedBarnSkin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedBarnSkin

`func (o *SkyBlockGarden) SetSelectedBarnSkin(v SkyBlockGardenBarnSkin)`

SetSelectedBarnSkin sets SelectedBarnSkin field to given value.

### HasSelectedBarnSkin

`func (o *SkyBlockGarden) HasSelectedBarnSkin() bool`

HasSelectedBarnSkin returns a boolean if a field has been set.

### GetUnlockedBarnSkins

`func (o *SkyBlockGarden) GetUnlockedBarnSkins() []SkyBlockGardenBarnSkin`

GetUnlockedBarnSkins returns the UnlockedBarnSkins field if non-nil, zero value otherwise.

### GetUnlockedBarnSkinsOk

`func (o *SkyBlockGarden) GetUnlockedBarnSkinsOk() (*[]SkyBlockGardenBarnSkin, bool)`

GetUnlockedBarnSkinsOk returns a tuple with the UnlockedBarnSkins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockedBarnSkins

`func (o *SkyBlockGarden) SetUnlockedBarnSkins(v []SkyBlockGardenBarnSkin)`

SetUnlockedBarnSkins sets UnlockedBarnSkins field to given value.

### HasUnlockedBarnSkins

`func (o *SkyBlockGarden) HasUnlockedBarnSkins() bool`

HasUnlockedBarnSkins returns a boolean if a field has been set.

### GetUnlockedPlotsIds

`func (o *SkyBlockGarden) GetUnlockedPlotsIds() []SkyBlockGardenUnlockedPlotsIdsInner`

GetUnlockedPlotsIds returns the UnlockedPlotsIds field if non-nil, zero value otherwise.

### GetUnlockedPlotsIdsOk

`func (o *SkyBlockGarden) GetUnlockedPlotsIdsOk() (*[]SkyBlockGardenUnlockedPlotsIdsInner, bool)`

GetUnlockedPlotsIdsOk returns a tuple with the UnlockedPlotsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockedPlotsIds

`func (o *SkyBlockGarden) SetUnlockedPlotsIds(v []SkyBlockGardenUnlockedPlotsIdsInner)`

SetUnlockedPlotsIds sets UnlockedPlotsIds field to given value.


### GetUuid

`func (o *SkyBlockGarden) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SkyBlockGarden) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SkyBlockGarden) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


