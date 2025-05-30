/*
Hypixel Public API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SkyBlockProfileMemberPlayerStatsEndIslandDragonFight type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkyBlockProfileMemberPlayerStatsEndIslandDragonFight{}

// SkyBlockProfileMemberPlayerStatsEndIslandDragonFight struct for SkyBlockProfileMemberPlayerStatsEndIslandDragonFight
type SkyBlockProfileMemberPlayerStatsEndIslandDragonFight struct {
	AmountSummoned           *SkyBlockProfileMemberPlayerStatsEndIslandDragonFightAmountSummoned           `json:"amount_summoned,omitempty"`
	EnderCrystalsDestroyed   *float64                                                                      `json:"ender_crystals_destroyed,omitempty"`
	FastestKill              *SkyBlockProfileMemberPlayerStatsEndIslandDragonFightFastestKill              `json:"fastest_kill,omitempty"`
	HighestRank              *SkyBlockProfileMemberPlayerStatsEndIslandDragonFightHighestRank              `json:"highest_rank,omitempty"`
	MostDamage               *SkyBlockProfileMemberPlayerStatsEndIslandDragonFightMostDamage               `json:"most_damage,omitempty"`
	SummoningEyesContributed *SkyBlockProfileMemberPlayerStatsEndIslandDragonFightSummoningEyesContributed `json:"summoning_eyes_contributed,omitempty"`
}

// NewSkyBlockProfileMemberPlayerStatsEndIslandDragonFight instantiates a new SkyBlockProfileMemberPlayerStatsEndIslandDragonFight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkyBlockProfileMemberPlayerStatsEndIslandDragonFight() *SkyBlockProfileMemberPlayerStatsEndIslandDragonFight {
	this := SkyBlockProfileMemberPlayerStatsEndIslandDragonFight{}
	return &this
}

// NewSkyBlockProfileMemberPlayerStatsEndIslandDragonFightWithDefaults instantiates a new SkyBlockProfileMemberPlayerStatsEndIslandDragonFight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkyBlockProfileMemberPlayerStatsEndIslandDragonFightWithDefaults() *SkyBlockProfileMemberPlayerStatsEndIslandDragonFight {
	this := SkyBlockProfileMemberPlayerStatsEndIslandDragonFight{}
	return &this
}

// GetAmountSummoned returns the AmountSummoned field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStatsEndIslandDragonFight) GetAmountSummoned() SkyBlockProfileMemberPlayerStatsEndIslandDragonFightAmountSummoned {
	if o == nil || IsNil(o.AmountSummoned) {
		var ret SkyBlockProfileMemberPlayerStatsEndIslandDragonFightAmountSummoned
		return ret
	}
	return *o.AmountSummoned
}

// GetAmountSummonedOk returns a tuple with the AmountSummoned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStatsEndIslandDragonFight) GetAmountSummonedOk() (*SkyBlockProfileMemberPlayerStatsEndIslandDragonFightAmountSummoned, bool) {
	if o == nil || IsNil(o.AmountSummoned) {
		return nil, false
	}
	return o.AmountSummoned, true
}

// HasAmountSummoned returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStatsEndIslandDragonFight) HasAmountSummoned() bool {
	if o != nil && !IsNil(o.AmountSummoned) {
		return true
	}

	return false
}

// SetAmountSummoned gets a reference to the given SkyBlockProfileMemberPlayerStatsEndIslandDragonFightAmountSummoned and assigns it to the AmountSummoned field.
func (o *SkyBlockProfileMemberPlayerStatsEndIslandDragonFight) SetAmountSummoned(v SkyBlockProfileMemberPlayerStatsEndIslandDragonFightAmountSummoned) {
	o.AmountSummoned = &v
}

// GetEnderCrystalsDestroyed returns the EnderCrystalsDestroyed field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStatsEndIslandDragonFight) GetEnderCrystalsDestroyed() float64 {
	if o == nil || IsNil(o.EnderCrystalsDestroyed) {
		var ret float64
		return ret
	}
	return *o.EnderCrystalsDestroyed
}

// GetEnderCrystalsDestroyedOk returns a tuple with the EnderCrystalsDestroyed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStatsEndIslandDragonFight) GetEnderCrystalsDestroyedOk() (*float64, bool) {
	if o == nil || IsNil(o.EnderCrystalsDestroyed) {
		return nil, false
	}
	return o.EnderCrystalsDestroyed, true
}

// HasEnderCrystalsDestroyed returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStatsEndIslandDragonFight) HasEnderCrystalsDestroyed() bool {
	if o != nil && !IsNil(o.EnderCrystalsDestroyed) {
		return true
	}

	return false
}

// SetEnderCrystalsDestroyed gets a reference to the given float64 and assigns it to the EnderCrystalsDestroyed field.
func (o *SkyBlockProfileMemberPlayerStatsEndIslandDragonFight) SetEnderCrystalsDestroyed(v float64) {
	o.EnderCrystalsDestroyed = &v
}

// GetFastestKill returns the FastestKill field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStatsEndIslandDragonFight) GetFastestKill() SkyBlockProfileMemberPlayerStatsEndIslandDragonFightFastestKill {
	if o == nil || IsNil(o.FastestKill) {
		var ret SkyBlockProfileMemberPlayerStatsEndIslandDragonFightFastestKill
		return ret
	}
	return *o.FastestKill
}

// GetFastestKillOk returns a tuple with the FastestKill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStatsEndIslandDragonFight) GetFastestKillOk() (*SkyBlockProfileMemberPlayerStatsEndIslandDragonFightFastestKill, bool) {
	if o == nil || IsNil(o.FastestKill) {
		return nil, false
	}
	return o.FastestKill, true
}

// HasFastestKill returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStatsEndIslandDragonFight) HasFastestKill() bool {
	if o != nil && !IsNil(o.FastestKill) {
		return true
	}

	return false
}

// SetFastestKill gets a reference to the given SkyBlockProfileMemberPlayerStatsEndIslandDragonFightFastestKill and assigns it to the FastestKill field.
func (o *SkyBlockProfileMemberPlayerStatsEndIslandDragonFight) SetFastestKill(v SkyBlockProfileMemberPlayerStatsEndIslandDragonFightFastestKill) {
	o.FastestKill = &v
}

// GetHighestRank returns the HighestRank field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStatsEndIslandDragonFight) GetHighestRank() SkyBlockProfileMemberPlayerStatsEndIslandDragonFightHighestRank {
	if o == nil || IsNil(o.HighestRank) {
		var ret SkyBlockProfileMemberPlayerStatsEndIslandDragonFightHighestRank
		return ret
	}
	return *o.HighestRank
}

// GetHighestRankOk returns a tuple with the HighestRank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStatsEndIslandDragonFight) GetHighestRankOk() (*SkyBlockProfileMemberPlayerStatsEndIslandDragonFightHighestRank, bool) {
	if o == nil || IsNil(o.HighestRank) {
		return nil, false
	}
	return o.HighestRank, true
}

// HasHighestRank returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStatsEndIslandDragonFight) HasHighestRank() bool {
	if o != nil && !IsNil(o.HighestRank) {
		return true
	}

	return false
}

// SetHighestRank gets a reference to the given SkyBlockProfileMemberPlayerStatsEndIslandDragonFightHighestRank and assigns it to the HighestRank field.
func (o *SkyBlockProfileMemberPlayerStatsEndIslandDragonFight) SetHighestRank(v SkyBlockProfileMemberPlayerStatsEndIslandDragonFightHighestRank) {
	o.HighestRank = &v
}

// GetMostDamage returns the MostDamage field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStatsEndIslandDragonFight) GetMostDamage() SkyBlockProfileMemberPlayerStatsEndIslandDragonFightMostDamage {
	if o == nil || IsNil(o.MostDamage) {
		var ret SkyBlockProfileMemberPlayerStatsEndIslandDragonFightMostDamage
		return ret
	}
	return *o.MostDamage
}

// GetMostDamageOk returns a tuple with the MostDamage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStatsEndIslandDragonFight) GetMostDamageOk() (*SkyBlockProfileMemberPlayerStatsEndIslandDragonFightMostDamage, bool) {
	if o == nil || IsNil(o.MostDamage) {
		return nil, false
	}
	return o.MostDamage, true
}

// HasMostDamage returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStatsEndIslandDragonFight) HasMostDamage() bool {
	if o != nil && !IsNil(o.MostDamage) {
		return true
	}

	return false
}

// SetMostDamage gets a reference to the given SkyBlockProfileMemberPlayerStatsEndIslandDragonFightMostDamage and assigns it to the MostDamage field.
func (o *SkyBlockProfileMemberPlayerStatsEndIslandDragonFight) SetMostDamage(v SkyBlockProfileMemberPlayerStatsEndIslandDragonFightMostDamage) {
	o.MostDamage = &v
}

// GetSummoningEyesContributed returns the SummoningEyesContributed field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStatsEndIslandDragonFight) GetSummoningEyesContributed() SkyBlockProfileMemberPlayerStatsEndIslandDragonFightSummoningEyesContributed {
	if o == nil || IsNil(o.SummoningEyesContributed) {
		var ret SkyBlockProfileMemberPlayerStatsEndIslandDragonFightSummoningEyesContributed
		return ret
	}
	return *o.SummoningEyesContributed
}

// GetSummoningEyesContributedOk returns a tuple with the SummoningEyesContributed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStatsEndIslandDragonFight) GetSummoningEyesContributedOk() (*SkyBlockProfileMemberPlayerStatsEndIslandDragonFightSummoningEyesContributed, bool) {
	if o == nil || IsNil(o.SummoningEyesContributed) {
		return nil, false
	}
	return o.SummoningEyesContributed, true
}

// HasSummoningEyesContributed returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStatsEndIslandDragonFight) HasSummoningEyesContributed() bool {
	if o != nil && !IsNil(o.SummoningEyesContributed) {
		return true
	}

	return false
}

// SetSummoningEyesContributed gets a reference to the given SkyBlockProfileMemberPlayerStatsEndIslandDragonFightSummoningEyesContributed and assigns it to the SummoningEyesContributed field.
func (o *SkyBlockProfileMemberPlayerStatsEndIslandDragonFight) SetSummoningEyesContributed(v SkyBlockProfileMemberPlayerStatsEndIslandDragonFightSummoningEyesContributed) {
	o.SummoningEyesContributed = &v
}

func (o SkyBlockProfileMemberPlayerStatsEndIslandDragonFight) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkyBlockProfileMemberPlayerStatsEndIslandDragonFight) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AmountSummoned) {
		toSerialize["amount_summoned"] = o.AmountSummoned
	}
	if !IsNil(o.EnderCrystalsDestroyed) {
		toSerialize["ender_crystals_destroyed"] = o.EnderCrystalsDestroyed
	}
	if !IsNil(o.FastestKill) {
		toSerialize["fastest_kill"] = o.FastestKill
	}
	if !IsNil(o.HighestRank) {
		toSerialize["highest_rank"] = o.HighestRank
	}
	if !IsNil(o.MostDamage) {
		toSerialize["most_damage"] = o.MostDamage
	}
	if !IsNil(o.SummoningEyesContributed) {
		toSerialize["summoning_eyes_contributed"] = o.SummoningEyesContributed
	}
	return toSerialize, nil
}

type NullableSkyBlockProfileMemberPlayerStatsEndIslandDragonFight struct {
	value *SkyBlockProfileMemberPlayerStatsEndIslandDragonFight
	isSet bool
}

func (v NullableSkyBlockProfileMemberPlayerStatsEndIslandDragonFight) Get() *SkyBlockProfileMemberPlayerStatsEndIslandDragonFight {
	return v.value
}

func (v *NullableSkyBlockProfileMemberPlayerStatsEndIslandDragonFight) Set(val *SkyBlockProfileMemberPlayerStatsEndIslandDragonFight) {
	v.value = val
	v.isSet = true
}

func (v NullableSkyBlockProfileMemberPlayerStatsEndIslandDragonFight) IsSet() bool {
	return v.isSet
}

func (v *NullableSkyBlockProfileMemberPlayerStatsEndIslandDragonFight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkyBlockProfileMemberPlayerStatsEndIslandDragonFight(val *SkyBlockProfileMemberPlayerStatsEndIslandDragonFight) *NullableSkyBlockProfileMemberPlayerStatsEndIslandDragonFight {
	return &NullableSkyBlockProfileMemberPlayerStatsEndIslandDragonFight{value: val, isSet: true}
}

func (v NullableSkyBlockProfileMemberPlayerStatsEndIslandDragonFight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkyBlockProfileMemberPlayerStatsEndIslandDragonFight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
