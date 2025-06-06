/*
Hypixel Public API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SkyBlockProfileMemberDungeonsTreasuresChestRewards type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkyBlockProfileMemberDungeonsTreasuresChestRewards{}

// SkyBlockProfileMemberDungeonsTreasuresChestRewards struct for SkyBlockProfileMemberDungeonsTreasuresChestRewards
type SkyBlockProfileMemberDungeonsTreasuresChestRewards struct {
	Rewards                []string `json:"rewards"`
	RngMeterReward         *string  `json:"rng_meter_reward,omitempty"`
	RolledRngMeterRandomly bool     `json:"rolled_rng_meter_randomly"`
}

type _SkyBlockProfileMemberDungeonsTreasuresChestRewards SkyBlockProfileMemberDungeonsTreasuresChestRewards

// NewSkyBlockProfileMemberDungeonsTreasuresChestRewards instantiates a new SkyBlockProfileMemberDungeonsTreasuresChestRewards object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkyBlockProfileMemberDungeonsTreasuresChestRewards(rewards []string, rolledRngMeterRandomly bool) *SkyBlockProfileMemberDungeonsTreasuresChestRewards {
	this := SkyBlockProfileMemberDungeonsTreasuresChestRewards{}
	this.Rewards = rewards
	this.RolledRngMeterRandomly = rolledRngMeterRandomly
	return &this
}

// NewSkyBlockProfileMemberDungeonsTreasuresChestRewardsWithDefaults instantiates a new SkyBlockProfileMemberDungeonsTreasuresChestRewards object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkyBlockProfileMemberDungeonsTreasuresChestRewardsWithDefaults() *SkyBlockProfileMemberDungeonsTreasuresChestRewards {
	this := SkyBlockProfileMemberDungeonsTreasuresChestRewards{}
	return &this
}

// GetRewards returns the Rewards field value
func (o *SkyBlockProfileMemberDungeonsTreasuresChestRewards) GetRewards() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Rewards
}

// GetRewardsOk returns a tuple with the Rewards field value
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberDungeonsTreasuresChestRewards) GetRewardsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rewards, true
}

// SetRewards sets field value
func (o *SkyBlockProfileMemberDungeonsTreasuresChestRewards) SetRewards(v []string) {
	o.Rewards = v
}

// GetRngMeterReward returns the RngMeterReward field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberDungeonsTreasuresChestRewards) GetRngMeterReward() string {
	if o == nil || IsNil(o.RngMeterReward) {
		var ret string
		return ret
	}
	return *o.RngMeterReward
}

// GetRngMeterRewardOk returns a tuple with the RngMeterReward field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberDungeonsTreasuresChestRewards) GetRngMeterRewardOk() (*string, bool) {
	if o == nil || IsNil(o.RngMeterReward) {
		return nil, false
	}
	return o.RngMeterReward, true
}

// HasRngMeterReward returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberDungeonsTreasuresChestRewards) HasRngMeterReward() bool {
	if o != nil && !IsNil(o.RngMeterReward) {
		return true
	}

	return false
}

// SetRngMeterReward gets a reference to the given string and assigns it to the RngMeterReward field.
func (o *SkyBlockProfileMemberDungeonsTreasuresChestRewards) SetRngMeterReward(v string) {
	o.RngMeterReward = &v
}

// GetRolledRngMeterRandomly returns the RolledRngMeterRandomly field value
func (o *SkyBlockProfileMemberDungeonsTreasuresChestRewards) GetRolledRngMeterRandomly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RolledRngMeterRandomly
}

// GetRolledRngMeterRandomlyOk returns a tuple with the RolledRngMeterRandomly field value
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberDungeonsTreasuresChestRewards) GetRolledRngMeterRandomlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RolledRngMeterRandomly, true
}

// SetRolledRngMeterRandomly sets field value
func (o *SkyBlockProfileMemberDungeonsTreasuresChestRewards) SetRolledRngMeterRandomly(v bool) {
	o.RolledRngMeterRandomly = v
}

func (o SkyBlockProfileMemberDungeonsTreasuresChestRewards) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkyBlockProfileMemberDungeonsTreasuresChestRewards) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rewards"] = o.Rewards
	if !IsNil(o.RngMeterReward) {
		toSerialize["rng_meter_reward"] = o.RngMeterReward
	}
	toSerialize["rolled_rng_meter_randomly"] = o.RolledRngMeterRandomly
	return toSerialize, nil
}

func (o *SkyBlockProfileMemberDungeonsTreasuresChestRewards) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rewards",
		"rolled_rng_meter_randomly",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSkyBlockProfileMemberDungeonsTreasuresChestRewards := _SkyBlockProfileMemberDungeonsTreasuresChestRewards{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSkyBlockProfileMemberDungeonsTreasuresChestRewards)

	if err != nil {
		return err
	}

	*o = SkyBlockProfileMemberDungeonsTreasuresChestRewards(varSkyBlockProfileMemberDungeonsTreasuresChestRewards)

	return err
}

type NullableSkyBlockProfileMemberDungeonsTreasuresChestRewards struct {
	value *SkyBlockProfileMemberDungeonsTreasuresChestRewards
	isSet bool
}

func (v NullableSkyBlockProfileMemberDungeonsTreasuresChestRewards) Get() *SkyBlockProfileMemberDungeonsTreasuresChestRewards {
	return v.value
}

func (v *NullableSkyBlockProfileMemberDungeonsTreasuresChestRewards) Set(val *SkyBlockProfileMemberDungeonsTreasuresChestRewards) {
	v.value = val
	v.isSet = true
}

func (v NullableSkyBlockProfileMemberDungeonsTreasuresChestRewards) IsSet() bool {
	return v.isSet
}

func (v *NullableSkyBlockProfileMemberDungeonsTreasuresChestRewards) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkyBlockProfileMemberDungeonsTreasuresChestRewards(val *SkyBlockProfileMemberDungeonsTreasuresChestRewards) *NullableSkyBlockProfileMemberDungeonsTreasuresChestRewards {
	return &NullableSkyBlockProfileMemberDungeonsTreasuresChestRewards{value: val, isSet: true}
}

func (v NullableSkyBlockProfileMemberDungeonsTreasuresChestRewards) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkyBlockProfileMemberDungeonsTreasuresChestRewards) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
