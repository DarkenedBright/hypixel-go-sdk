# SkyBlockProfileMemberPetsDataPet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** |  | 
**CandyUsed** | **int64** |  | 
**Exp** | **float64** |  | 
**Extra** | Pointer to **map[string]int64** |  | [optional] 
**HeldItem** | [**NullableSkyBlockProfileMemberPetsDataPetHeldItem**](SkyBlockProfileMemberPetsDataPetHeldItem.md) |  | 
**Skin** | **string** |  | 
**Tier** | [**SkyBlockProfileMemberPetsDataPetTier**](SkyBlockProfileMemberPetsDataPetTier.md) |  | 
**Type** | [**SkyBlockProfileMemberPetsDataPetType**](SkyBlockProfileMemberPetsDataPetType.md) |  | 
**UniqueId** | **string** |  | 
**Uuid** | **string** |  | 

## Methods

### NewSkyBlockProfileMemberPetsDataPet

`func NewSkyBlockProfileMemberPetsDataPet(active bool, candyUsed int64, exp float64, heldItem NullableSkyBlockProfileMemberPetsDataPetHeldItem, skin string, tier SkyBlockProfileMemberPetsDataPetTier, type_ SkyBlockProfileMemberPetsDataPetType, uniqueId string, uuid string, ) *SkyBlockProfileMemberPetsDataPet`

NewSkyBlockProfileMemberPetsDataPet instantiates a new SkyBlockProfileMemberPetsDataPet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkyBlockProfileMemberPetsDataPetWithDefaults

`func NewSkyBlockProfileMemberPetsDataPetWithDefaults() *SkyBlockProfileMemberPetsDataPet`

NewSkyBlockProfileMemberPetsDataPetWithDefaults instantiates a new SkyBlockProfileMemberPetsDataPet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *SkyBlockProfileMemberPetsDataPet) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SkyBlockProfileMemberPetsDataPet) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SkyBlockProfileMemberPetsDataPet) SetActive(v bool)`

SetActive sets Active field to given value.


### GetCandyUsed

`func (o *SkyBlockProfileMemberPetsDataPet) GetCandyUsed() int64`

GetCandyUsed returns the CandyUsed field if non-nil, zero value otherwise.

### GetCandyUsedOk

`func (o *SkyBlockProfileMemberPetsDataPet) GetCandyUsedOk() (*int64, bool)`

GetCandyUsedOk returns a tuple with the CandyUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandyUsed

`func (o *SkyBlockProfileMemberPetsDataPet) SetCandyUsed(v int64)`

SetCandyUsed sets CandyUsed field to given value.


### GetExp

`func (o *SkyBlockProfileMemberPetsDataPet) GetExp() float64`

GetExp returns the Exp field if non-nil, zero value otherwise.

### GetExpOk

`func (o *SkyBlockProfileMemberPetsDataPet) GetExpOk() (*float64, bool)`

GetExpOk returns a tuple with the Exp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExp

`func (o *SkyBlockProfileMemberPetsDataPet) SetExp(v float64)`

SetExp sets Exp field to given value.


### GetExtra

`func (o *SkyBlockProfileMemberPetsDataPet) GetExtra() map[string]int64`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *SkyBlockProfileMemberPetsDataPet) GetExtraOk() (*map[string]int64, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *SkyBlockProfileMemberPetsDataPet) SetExtra(v map[string]int64)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *SkyBlockProfileMemberPetsDataPet) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetHeldItem

`func (o *SkyBlockProfileMemberPetsDataPet) GetHeldItem() SkyBlockProfileMemberPetsDataPetHeldItem`

GetHeldItem returns the HeldItem field if non-nil, zero value otherwise.

### GetHeldItemOk

`func (o *SkyBlockProfileMemberPetsDataPet) GetHeldItemOk() (*SkyBlockProfileMemberPetsDataPetHeldItem, bool)`

GetHeldItemOk returns a tuple with the HeldItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeldItem

`func (o *SkyBlockProfileMemberPetsDataPet) SetHeldItem(v SkyBlockProfileMemberPetsDataPetHeldItem)`

SetHeldItem sets HeldItem field to given value.


### SetHeldItemNil

`func (o *SkyBlockProfileMemberPetsDataPet) SetHeldItemNil(b bool)`

 SetHeldItemNil sets the value for HeldItem to be an explicit nil

### UnsetHeldItem
`func (o *SkyBlockProfileMemberPetsDataPet) UnsetHeldItem()`

UnsetHeldItem ensures that no value is present for HeldItem, not even an explicit nil
### GetSkin

`func (o *SkyBlockProfileMemberPetsDataPet) GetSkin() string`

GetSkin returns the Skin field if non-nil, zero value otherwise.

### GetSkinOk

`func (o *SkyBlockProfileMemberPetsDataPet) GetSkinOk() (*string, bool)`

GetSkinOk returns a tuple with the Skin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkin

`func (o *SkyBlockProfileMemberPetsDataPet) SetSkin(v string)`

SetSkin sets Skin field to given value.


### GetTier

`func (o *SkyBlockProfileMemberPetsDataPet) GetTier() SkyBlockProfileMemberPetsDataPetTier`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *SkyBlockProfileMemberPetsDataPet) GetTierOk() (*SkyBlockProfileMemberPetsDataPetTier, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *SkyBlockProfileMemberPetsDataPet) SetTier(v SkyBlockProfileMemberPetsDataPetTier)`

SetTier sets Tier field to given value.


### GetType

`func (o *SkyBlockProfileMemberPetsDataPet) GetType() SkyBlockProfileMemberPetsDataPetType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SkyBlockProfileMemberPetsDataPet) GetTypeOk() (*SkyBlockProfileMemberPetsDataPetType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SkyBlockProfileMemberPetsDataPet) SetType(v SkyBlockProfileMemberPetsDataPetType)`

SetType sets Type field to given value.


### GetUniqueId

`func (o *SkyBlockProfileMemberPetsDataPet) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *SkyBlockProfileMemberPetsDataPet) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *SkyBlockProfileMemberPetsDataPet) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.


### GetUuid

`func (o *SkyBlockProfileMemberPetsDataPet) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SkyBlockProfileMemberPetsDataPet) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SkyBlockProfileMemberPetsDataPet) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


