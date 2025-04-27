# SkyBlockProfileMemberPlayerDataActiveEffect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effect** | [**SkyBlockProfileMemberPlayerDataActiveEffectType**](SkyBlockProfileMemberPlayerDataActiveEffectType.md) |  | 
**Flags** | Pointer to **int64** |  | [optional] 
**Infinite** | **bool** |  | 
**Level** | **int64** |  | 
**Modifiers** | [**[]SkyBlockProfileMemberPlayerDataActiveEffectModifier**](SkyBlockProfileMemberPlayerDataActiveEffectModifier.md) |  | 
**TicksRemaining** | **int64** |  | 

## Methods

### NewSkyBlockProfileMemberPlayerDataActiveEffect

`func NewSkyBlockProfileMemberPlayerDataActiveEffect(effect SkyBlockProfileMemberPlayerDataActiveEffectType, infinite bool, level int64, modifiers []SkyBlockProfileMemberPlayerDataActiveEffectModifier, ticksRemaining int64, ) *SkyBlockProfileMemberPlayerDataActiveEffect`

NewSkyBlockProfileMemberPlayerDataActiveEffect instantiates a new SkyBlockProfileMemberPlayerDataActiveEffect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkyBlockProfileMemberPlayerDataActiveEffectWithDefaults

`func NewSkyBlockProfileMemberPlayerDataActiveEffectWithDefaults() *SkyBlockProfileMemberPlayerDataActiveEffect`

NewSkyBlockProfileMemberPlayerDataActiveEffectWithDefaults instantiates a new SkyBlockProfileMemberPlayerDataActiveEffect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffect

`func (o *SkyBlockProfileMemberPlayerDataActiveEffect) GetEffect() SkyBlockProfileMemberPlayerDataActiveEffectType`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *SkyBlockProfileMemberPlayerDataActiveEffect) GetEffectOk() (*SkyBlockProfileMemberPlayerDataActiveEffectType, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *SkyBlockProfileMemberPlayerDataActiveEffect) SetEffect(v SkyBlockProfileMemberPlayerDataActiveEffectType)`

SetEffect sets Effect field to given value.


### GetFlags

`func (o *SkyBlockProfileMemberPlayerDataActiveEffect) GetFlags() int64`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *SkyBlockProfileMemberPlayerDataActiveEffect) GetFlagsOk() (*int64, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *SkyBlockProfileMemberPlayerDataActiveEffect) SetFlags(v int64)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *SkyBlockProfileMemberPlayerDataActiveEffect) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetInfinite

`func (o *SkyBlockProfileMemberPlayerDataActiveEffect) GetInfinite() bool`

GetInfinite returns the Infinite field if non-nil, zero value otherwise.

### GetInfiniteOk

`func (o *SkyBlockProfileMemberPlayerDataActiveEffect) GetInfiniteOk() (*bool, bool)`

GetInfiniteOk returns a tuple with the Infinite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfinite

`func (o *SkyBlockProfileMemberPlayerDataActiveEffect) SetInfinite(v bool)`

SetInfinite sets Infinite field to given value.


### GetLevel

`func (o *SkyBlockProfileMemberPlayerDataActiveEffect) GetLevel() int64`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *SkyBlockProfileMemberPlayerDataActiveEffect) GetLevelOk() (*int64, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *SkyBlockProfileMemberPlayerDataActiveEffect) SetLevel(v int64)`

SetLevel sets Level field to given value.


### GetModifiers

`func (o *SkyBlockProfileMemberPlayerDataActiveEffect) GetModifiers() []SkyBlockProfileMemberPlayerDataActiveEffectModifier`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *SkyBlockProfileMemberPlayerDataActiveEffect) GetModifiersOk() (*[]SkyBlockProfileMemberPlayerDataActiveEffectModifier, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *SkyBlockProfileMemberPlayerDataActiveEffect) SetModifiers(v []SkyBlockProfileMemberPlayerDataActiveEffectModifier)`

SetModifiers sets Modifiers field to given value.


### GetTicksRemaining

`func (o *SkyBlockProfileMemberPlayerDataActiveEffect) GetTicksRemaining() int64`

GetTicksRemaining returns the TicksRemaining field if non-nil, zero value otherwise.

### GetTicksRemainingOk

`func (o *SkyBlockProfileMemberPlayerDataActiveEffect) GetTicksRemainingOk() (*int64, bool)`

GetTicksRemainingOk returns a tuple with the TicksRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicksRemaining

`func (o *SkyBlockProfileMemberPlayerDataActiveEffect) SetTicksRemaining(v int64)`

SetTicksRemaining sets TicksRemaining field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


