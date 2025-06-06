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

// checks if the SkyBlockProfileMemberDungeonsTreasuresRun type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkyBlockProfileMemberDungeonsTreasuresRun{}

// SkyBlockProfileMemberDungeonsTreasuresRun struct for SkyBlockProfileMemberDungeonsTreasuresRun
type SkyBlockProfileMemberDungeonsTreasuresRun struct {
	CompletionTs int64                                                  `json:"completion_ts"`
	DungeonTier  int64                                                  `json:"dungeon_tier"`
	DungeonType  SkyBlockProfileMemberDungeonsTreasuresRunDungeonType   `json:"dungeon_type"`
	Participants []SkyBlockProfileMemberDungeonsTreasuresRunParticipant `json:"participants"`
	RunId        string                                                 `json:"run_id"`
}

type _SkyBlockProfileMemberDungeonsTreasuresRun SkyBlockProfileMemberDungeonsTreasuresRun

// NewSkyBlockProfileMemberDungeonsTreasuresRun instantiates a new SkyBlockProfileMemberDungeonsTreasuresRun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkyBlockProfileMemberDungeonsTreasuresRun(completionTs int64, dungeonTier int64, dungeonType SkyBlockProfileMemberDungeonsTreasuresRunDungeonType, participants []SkyBlockProfileMemberDungeonsTreasuresRunParticipant, runId string) *SkyBlockProfileMemberDungeonsTreasuresRun {
	this := SkyBlockProfileMemberDungeonsTreasuresRun{}
	this.CompletionTs = completionTs
	this.DungeonTier = dungeonTier
	this.DungeonType = dungeonType
	this.Participants = participants
	this.RunId = runId
	return &this
}

// NewSkyBlockProfileMemberDungeonsTreasuresRunWithDefaults instantiates a new SkyBlockProfileMemberDungeonsTreasuresRun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkyBlockProfileMemberDungeonsTreasuresRunWithDefaults() *SkyBlockProfileMemberDungeonsTreasuresRun {
	this := SkyBlockProfileMemberDungeonsTreasuresRun{}
	return &this
}

// GetCompletionTs returns the CompletionTs field value
func (o *SkyBlockProfileMemberDungeonsTreasuresRun) GetCompletionTs() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CompletionTs
}

// GetCompletionTsOk returns a tuple with the CompletionTs field value
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberDungeonsTreasuresRun) GetCompletionTsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletionTs, true
}

// SetCompletionTs sets field value
func (o *SkyBlockProfileMemberDungeonsTreasuresRun) SetCompletionTs(v int64) {
	o.CompletionTs = v
}

// GetDungeonTier returns the DungeonTier field value
func (o *SkyBlockProfileMemberDungeonsTreasuresRun) GetDungeonTier() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DungeonTier
}

// GetDungeonTierOk returns a tuple with the DungeonTier field value
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberDungeonsTreasuresRun) GetDungeonTierOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DungeonTier, true
}

// SetDungeonTier sets field value
func (o *SkyBlockProfileMemberDungeonsTreasuresRun) SetDungeonTier(v int64) {
	o.DungeonTier = v
}

// GetDungeonType returns the DungeonType field value
func (o *SkyBlockProfileMemberDungeonsTreasuresRun) GetDungeonType() SkyBlockProfileMemberDungeonsTreasuresRunDungeonType {
	if o == nil {
		var ret SkyBlockProfileMemberDungeonsTreasuresRunDungeonType
		return ret
	}

	return o.DungeonType
}

// GetDungeonTypeOk returns a tuple with the DungeonType field value
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberDungeonsTreasuresRun) GetDungeonTypeOk() (*SkyBlockProfileMemberDungeonsTreasuresRunDungeonType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DungeonType, true
}

// SetDungeonType sets field value
func (o *SkyBlockProfileMemberDungeonsTreasuresRun) SetDungeonType(v SkyBlockProfileMemberDungeonsTreasuresRunDungeonType) {
	o.DungeonType = v
}

// GetParticipants returns the Participants field value
func (o *SkyBlockProfileMemberDungeonsTreasuresRun) GetParticipants() []SkyBlockProfileMemberDungeonsTreasuresRunParticipant {
	if o == nil {
		var ret []SkyBlockProfileMemberDungeonsTreasuresRunParticipant
		return ret
	}

	return o.Participants
}

// GetParticipantsOk returns a tuple with the Participants field value
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberDungeonsTreasuresRun) GetParticipantsOk() ([]SkyBlockProfileMemberDungeonsTreasuresRunParticipant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Participants, true
}

// SetParticipants sets field value
func (o *SkyBlockProfileMemberDungeonsTreasuresRun) SetParticipants(v []SkyBlockProfileMemberDungeonsTreasuresRunParticipant) {
	o.Participants = v
}

// GetRunId returns the RunId field value
func (o *SkyBlockProfileMemberDungeonsTreasuresRun) GetRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberDungeonsTreasuresRun) GetRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value
func (o *SkyBlockProfileMemberDungeonsTreasuresRun) SetRunId(v string) {
	o.RunId = v
}

func (o SkyBlockProfileMemberDungeonsTreasuresRun) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkyBlockProfileMemberDungeonsTreasuresRun) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["completion_ts"] = o.CompletionTs
	toSerialize["dungeon_tier"] = o.DungeonTier
	toSerialize["dungeon_type"] = o.DungeonType
	toSerialize["participants"] = o.Participants
	toSerialize["run_id"] = o.RunId
	return toSerialize, nil
}

func (o *SkyBlockProfileMemberDungeonsTreasuresRun) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"completion_ts",
		"dungeon_tier",
		"dungeon_type",
		"participants",
		"run_id",
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

	varSkyBlockProfileMemberDungeonsTreasuresRun := _SkyBlockProfileMemberDungeonsTreasuresRun{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSkyBlockProfileMemberDungeonsTreasuresRun)

	if err != nil {
		return err
	}

	*o = SkyBlockProfileMemberDungeonsTreasuresRun(varSkyBlockProfileMemberDungeonsTreasuresRun)

	return err
}

type NullableSkyBlockProfileMemberDungeonsTreasuresRun struct {
	value *SkyBlockProfileMemberDungeonsTreasuresRun
	isSet bool
}

func (v NullableSkyBlockProfileMemberDungeonsTreasuresRun) Get() *SkyBlockProfileMemberDungeonsTreasuresRun {
	return v.value
}

func (v *NullableSkyBlockProfileMemberDungeonsTreasuresRun) Set(val *SkyBlockProfileMemberDungeonsTreasuresRun) {
	v.value = val
	v.isSet = true
}

func (v NullableSkyBlockProfileMemberDungeonsTreasuresRun) IsSet() bool {
	return v.isSet
}

func (v *NullableSkyBlockProfileMemberDungeonsTreasuresRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkyBlockProfileMemberDungeonsTreasuresRun(val *SkyBlockProfileMemberDungeonsTreasuresRun) *NullableSkyBlockProfileMemberDungeonsTreasuresRun {
	return &NullableSkyBlockProfileMemberDungeonsTreasuresRun{value: val, isSet: true}
}

func (v NullableSkyBlockProfileMemberDungeonsTreasuresRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkyBlockProfileMemberDungeonsTreasuresRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
