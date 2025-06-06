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

// checks if the SkyBlockProfileMemberDungeonsDailyRuns type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkyBlockProfileMemberDungeonsDailyRuns{}

// SkyBlockProfileMemberDungeonsDailyRuns struct for SkyBlockProfileMemberDungeonsDailyRuns
type SkyBlockProfileMemberDungeonsDailyRuns struct {
	CompletedRunsCount int64 `json:"completed_runs_count"`
	CurrentDayStamp    int64 `json:"current_day_stamp"`
}

type _SkyBlockProfileMemberDungeonsDailyRuns SkyBlockProfileMemberDungeonsDailyRuns

// NewSkyBlockProfileMemberDungeonsDailyRuns instantiates a new SkyBlockProfileMemberDungeonsDailyRuns object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkyBlockProfileMemberDungeonsDailyRuns(completedRunsCount int64, currentDayStamp int64) *SkyBlockProfileMemberDungeonsDailyRuns {
	this := SkyBlockProfileMemberDungeonsDailyRuns{}
	this.CompletedRunsCount = completedRunsCount
	this.CurrentDayStamp = currentDayStamp
	return &this
}

// NewSkyBlockProfileMemberDungeonsDailyRunsWithDefaults instantiates a new SkyBlockProfileMemberDungeonsDailyRuns object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkyBlockProfileMemberDungeonsDailyRunsWithDefaults() *SkyBlockProfileMemberDungeonsDailyRuns {
	this := SkyBlockProfileMemberDungeonsDailyRuns{}
	return &this
}

// GetCompletedRunsCount returns the CompletedRunsCount field value
func (o *SkyBlockProfileMemberDungeonsDailyRuns) GetCompletedRunsCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CompletedRunsCount
}

// GetCompletedRunsCountOk returns a tuple with the CompletedRunsCount field value
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberDungeonsDailyRuns) GetCompletedRunsCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletedRunsCount, true
}

// SetCompletedRunsCount sets field value
func (o *SkyBlockProfileMemberDungeonsDailyRuns) SetCompletedRunsCount(v int64) {
	o.CompletedRunsCount = v
}

// GetCurrentDayStamp returns the CurrentDayStamp field value
func (o *SkyBlockProfileMemberDungeonsDailyRuns) GetCurrentDayStamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CurrentDayStamp
}

// GetCurrentDayStampOk returns a tuple with the CurrentDayStamp field value
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberDungeonsDailyRuns) GetCurrentDayStampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentDayStamp, true
}

// SetCurrentDayStamp sets field value
func (o *SkyBlockProfileMemberDungeonsDailyRuns) SetCurrentDayStamp(v int64) {
	o.CurrentDayStamp = v
}

func (o SkyBlockProfileMemberDungeonsDailyRuns) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkyBlockProfileMemberDungeonsDailyRuns) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["completed_runs_count"] = o.CompletedRunsCount
	toSerialize["current_day_stamp"] = o.CurrentDayStamp
	return toSerialize, nil
}

func (o *SkyBlockProfileMemberDungeonsDailyRuns) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"completed_runs_count",
		"current_day_stamp",
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

	varSkyBlockProfileMemberDungeonsDailyRuns := _SkyBlockProfileMemberDungeonsDailyRuns{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSkyBlockProfileMemberDungeonsDailyRuns)

	if err != nil {
		return err
	}

	*o = SkyBlockProfileMemberDungeonsDailyRuns(varSkyBlockProfileMemberDungeonsDailyRuns)

	return err
}

type NullableSkyBlockProfileMemberDungeonsDailyRuns struct {
	value *SkyBlockProfileMemberDungeonsDailyRuns
	isSet bool
}

func (v NullableSkyBlockProfileMemberDungeonsDailyRuns) Get() *SkyBlockProfileMemberDungeonsDailyRuns {
	return v.value
}

func (v *NullableSkyBlockProfileMemberDungeonsDailyRuns) Set(val *SkyBlockProfileMemberDungeonsDailyRuns) {
	v.value = val
	v.isSet = true
}

func (v NullableSkyBlockProfileMemberDungeonsDailyRuns) IsSet() bool {
	return v.isSet
}

func (v *NullableSkyBlockProfileMemberDungeonsDailyRuns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkyBlockProfileMemberDungeonsDailyRuns(val *SkyBlockProfileMemberDungeonsDailyRuns) *NullableSkyBlockProfileMemberDungeonsDailyRuns {
	return &NullableSkyBlockProfileMemberDungeonsDailyRuns{value: val, isSet: true}
}

func (v NullableSkyBlockProfileMemberDungeonsDailyRuns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkyBlockProfileMemberDungeonsDailyRuns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
