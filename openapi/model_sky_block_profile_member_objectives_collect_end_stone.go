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

// checks if the SkyBlockProfileMemberObjectivesCollectEndStone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkyBlockProfileMemberObjectivesCollectEndStone{}

// SkyBlockProfileMemberObjectivesCollectEndStone struct for SkyBlockProfileMemberObjectivesCollectEndStone
type SkyBlockProfileMemberObjectivesCollectEndStone struct {
	ENDER_STONE *bool                                                             `json:"ENDER_STONE,omitempty"`
	CompletedAt int64                                                             `json:"completed_at"`
	Progress    int64                                                             `json:"progress"`
	Status      SkyBlockProfileMemberNetherIslandPlayerQuestsQuestDataQuestStatus `json:"status"`
}

type _SkyBlockProfileMemberObjectivesCollectEndStone SkyBlockProfileMemberObjectivesCollectEndStone

// NewSkyBlockProfileMemberObjectivesCollectEndStone instantiates a new SkyBlockProfileMemberObjectivesCollectEndStone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkyBlockProfileMemberObjectivesCollectEndStone(completedAt int64, progress int64, status SkyBlockProfileMemberNetherIslandPlayerQuestsQuestDataQuestStatus) *SkyBlockProfileMemberObjectivesCollectEndStone {
	this := SkyBlockProfileMemberObjectivesCollectEndStone{}
	this.CompletedAt = completedAt
	this.Progress = progress
	this.Status = status
	return &this
}

// NewSkyBlockProfileMemberObjectivesCollectEndStoneWithDefaults instantiates a new SkyBlockProfileMemberObjectivesCollectEndStone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkyBlockProfileMemberObjectivesCollectEndStoneWithDefaults() *SkyBlockProfileMemberObjectivesCollectEndStone {
	this := SkyBlockProfileMemberObjectivesCollectEndStone{}
	return &this
}

// GetENDER_STONE returns the ENDER_STONE field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberObjectivesCollectEndStone) GetENDER_STONE() bool {
	if o == nil || IsNil(o.ENDER_STONE) {
		var ret bool
		return ret
	}
	return *o.ENDER_STONE
}

// GetENDER_STONEOk returns a tuple with the ENDER_STONE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberObjectivesCollectEndStone) GetENDER_STONEOk() (*bool, bool) {
	if o == nil || IsNil(o.ENDER_STONE) {
		return nil, false
	}
	return o.ENDER_STONE, true
}

// HasENDER_STONE returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberObjectivesCollectEndStone) HasENDER_STONE() bool {
	if o != nil && !IsNil(o.ENDER_STONE) {
		return true
	}

	return false
}

// SetENDER_STONE gets a reference to the given bool and assigns it to the ENDER_STONE field.
func (o *SkyBlockProfileMemberObjectivesCollectEndStone) SetENDER_STONE(v bool) {
	o.ENDER_STONE = &v
}

// GetCompletedAt returns the CompletedAt field value
func (o *SkyBlockProfileMemberObjectivesCollectEndStone) GetCompletedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberObjectivesCollectEndStone) GetCompletedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletedAt, true
}

// SetCompletedAt sets field value
func (o *SkyBlockProfileMemberObjectivesCollectEndStone) SetCompletedAt(v int64) {
	o.CompletedAt = v
}

// GetProgress returns the Progress field value
func (o *SkyBlockProfileMemberObjectivesCollectEndStone) GetProgress() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Progress
}

// GetProgressOk returns a tuple with the Progress field value
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberObjectivesCollectEndStone) GetProgressOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Progress, true
}

// SetProgress sets field value
func (o *SkyBlockProfileMemberObjectivesCollectEndStone) SetProgress(v int64) {
	o.Progress = v
}

// GetStatus returns the Status field value
func (o *SkyBlockProfileMemberObjectivesCollectEndStone) GetStatus() SkyBlockProfileMemberNetherIslandPlayerQuestsQuestDataQuestStatus {
	if o == nil {
		var ret SkyBlockProfileMemberNetherIslandPlayerQuestsQuestDataQuestStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberObjectivesCollectEndStone) GetStatusOk() (*SkyBlockProfileMemberNetherIslandPlayerQuestsQuestDataQuestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SkyBlockProfileMemberObjectivesCollectEndStone) SetStatus(v SkyBlockProfileMemberNetherIslandPlayerQuestsQuestDataQuestStatus) {
	o.Status = v
}

func (o SkyBlockProfileMemberObjectivesCollectEndStone) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkyBlockProfileMemberObjectivesCollectEndStone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ENDER_STONE) {
		toSerialize["ENDER_STONE"] = o.ENDER_STONE
	}
	toSerialize["completed_at"] = o.CompletedAt
	toSerialize["progress"] = o.Progress
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *SkyBlockProfileMemberObjectivesCollectEndStone) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"completed_at",
		"progress",
		"status",
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

	varSkyBlockProfileMemberObjectivesCollectEndStone := _SkyBlockProfileMemberObjectivesCollectEndStone{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSkyBlockProfileMemberObjectivesCollectEndStone)

	if err != nil {
		return err
	}

	*o = SkyBlockProfileMemberObjectivesCollectEndStone(varSkyBlockProfileMemberObjectivesCollectEndStone)

	return err
}

type NullableSkyBlockProfileMemberObjectivesCollectEndStone struct {
	value *SkyBlockProfileMemberObjectivesCollectEndStone
	isSet bool
}

func (v NullableSkyBlockProfileMemberObjectivesCollectEndStone) Get() *SkyBlockProfileMemberObjectivesCollectEndStone {
	return v.value
}

func (v *NullableSkyBlockProfileMemberObjectivesCollectEndStone) Set(val *SkyBlockProfileMemberObjectivesCollectEndStone) {
	v.value = val
	v.isSet = true
}

func (v NullableSkyBlockProfileMemberObjectivesCollectEndStone) IsSet() bool {
	return v.isSet
}

func (v *NullableSkyBlockProfileMemberObjectivesCollectEndStone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkyBlockProfileMemberObjectivesCollectEndStone(val *SkyBlockProfileMemberObjectivesCollectEndStone) *NullableSkyBlockProfileMemberObjectivesCollectEndStone {
	return &NullableSkyBlockProfileMemberObjectivesCollectEndStone{value: val, isSet: true}
}

func (v NullableSkyBlockProfileMemberObjectivesCollectEndStone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkyBlockProfileMemberObjectivesCollectEndStone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
