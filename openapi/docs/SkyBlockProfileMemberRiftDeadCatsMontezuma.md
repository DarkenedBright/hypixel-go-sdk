# SkyBlockProfileMemberRiftDeadCatsMontezuma

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** |  | 
**CandyUsed** | **int64** |  | 
**Exp** | **float64** |  | 
**HeldItem** | [**NullableSkyBlockProfileMemberRiftDeadCatsMontezumaHeldItem**](SkyBlockProfileMemberRiftDeadCatsMontezumaHeldItem.md) |  | 
**Skin** | **NullableString** |  | 
**Tier** | [**SkyBlockProfileMemberRiftDeadCatsMontezumaTier**](SkyBlockProfileMemberRiftDeadCatsMontezumaTier.md) |  | 
**Type** | [**SkyBlockProfileMemberRiftDeadCatsMontezumaType**](SkyBlockProfileMemberRiftDeadCatsMontezumaType.md) |  | 
**UniqueId** | Pointer to **string** |  | [optional] 
**Uuid** | **NullableString** |  | 

## Methods

### NewSkyBlockProfileMemberRiftDeadCatsMontezuma

`func NewSkyBlockProfileMemberRiftDeadCatsMontezuma(active bool, candyUsed int64, exp float64, heldItem NullableSkyBlockProfileMemberRiftDeadCatsMontezumaHeldItem, skin NullableString, tier SkyBlockProfileMemberRiftDeadCatsMontezumaTier, type_ SkyBlockProfileMemberRiftDeadCatsMontezumaType, uuid NullableString, ) *SkyBlockProfileMemberRiftDeadCatsMontezuma`

NewSkyBlockProfileMemberRiftDeadCatsMontezuma instantiates a new SkyBlockProfileMemberRiftDeadCatsMontezuma object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkyBlockProfileMemberRiftDeadCatsMontezumaWithDefaults

`func NewSkyBlockProfileMemberRiftDeadCatsMontezumaWithDefaults() *SkyBlockProfileMemberRiftDeadCatsMontezuma`

NewSkyBlockProfileMemberRiftDeadCatsMontezumaWithDefaults instantiates a new SkyBlockProfileMemberRiftDeadCatsMontezuma object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) SetActive(v bool)`

SetActive sets Active field to given value.


### GetCandyUsed

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) GetCandyUsed() int64`

GetCandyUsed returns the CandyUsed field if non-nil, zero value otherwise.

### GetCandyUsedOk

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) GetCandyUsedOk() (*int64, bool)`

GetCandyUsedOk returns a tuple with the CandyUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandyUsed

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) SetCandyUsed(v int64)`

SetCandyUsed sets CandyUsed field to given value.


### GetExp

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) GetExp() float64`

GetExp returns the Exp field if non-nil, zero value otherwise.

### GetExpOk

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) GetExpOk() (*float64, bool)`

GetExpOk returns a tuple with the Exp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExp

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) SetExp(v float64)`

SetExp sets Exp field to given value.


### GetHeldItem

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) GetHeldItem() SkyBlockProfileMemberRiftDeadCatsMontezumaHeldItem`

GetHeldItem returns the HeldItem field if non-nil, zero value otherwise.

### GetHeldItemOk

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) GetHeldItemOk() (*SkyBlockProfileMemberRiftDeadCatsMontezumaHeldItem, bool)`

GetHeldItemOk returns a tuple with the HeldItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeldItem

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) SetHeldItem(v SkyBlockProfileMemberRiftDeadCatsMontezumaHeldItem)`

SetHeldItem sets HeldItem field to given value.


### SetHeldItemNil

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) SetHeldItemNil(b bool)`

 SetHeldItemNil sets the value for HeldItem to be an explicit nil

### UnsetHeldItem
`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) UnsetHeldItem()`

UnsetHeldItem ensures that no value is present for HeldItem, not even an explicit nil
### GetSkin

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) GetSkin() string`

GetSkin returns the Skin field if non-nil, zero value otherwise.

### GetSkinOk

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) GetSkinOk() (*string, bool)`

GetSkinOk returns a tuple with the Skin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkin

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) SetSkin(v string)`

SetSkin sets Skin field to given value.


### SetSkinNil

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) SetSkinNil(b bool)`

 SetSkinNil sets the value for Skin to be an explicit nil

### UnsetSkin
`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) UnsetSkin()`

UnsetSkin ensures that no value is present for Skin, not even an explicit nil
### GetTier

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) GetTier() SkyBlockProfileMemberRiftDeadCatsMontezumaTier`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) GetTierOk() (*SkyBlockProfileMemberRiftDeadCatsMontezumaTier, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) SetTier(v SkyBlockProfileMemberRiftDeadCatsMontezumaTier)`

SetTier sets Tier field to given value.


### GetType

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) GetType() SkyBlockProfileMemberRiftDeadCatsMontezumaType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) GetTypeOk() (*SkyBlockProfileMemberRiftDeadCatsMontezumaType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) SetType(v SkyBlockProfileMemberRiftDeadCatsMontezumaType)`

SetType sets Type field to given value.


### GetUniqueId

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetUuid

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### SetUuidNil

`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *SkyBlockProfileMemberRiftDeadCatsMontezuma) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


