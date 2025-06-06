/*
Hypixel Public API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// SkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue - struct for SkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue
type SkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue struct {
	Int64  *int64
	String *string
}

// int64AsSkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue is a convenience function that returns int64 wrapped in SkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue
func Int64AsSkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue(v *int64) SkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue {
	return SkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue{
		Int64: v,
	}
}

// stringAsSkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue is a convenience function that returns string wrapped in SkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue
func StringAsSkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue(v *string) SkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue {
	return SkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Int64
	err = json.Unmarshal(data, &dst.Int64)
	if err == nil {
		jsonInt64, _ := json.Marshal(dst.Int64)
		if string(jsonInt64) == "{}" { // empty struct
			dst.Int64 = nil
		} else {
			match++
		}
	} else {
		dst.Int64 = nil
	}

	// try to unmarshal data into String
	err = json.Unmarshal(data, &dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Int64 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue) MarshalJSON() ([]byte, error) {
	if src.Int64 != nil {
		return json.Marshal(&src.Int64)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Int64 != nil {
		return obj.Int64
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj SkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue) GetActualInstanceValue() interface{} {
	if obj.Int64 != nil {
		return *obj.Int64
	}

	if obj.String != nil {
		return *obj.String
	}

	// all schemas are nil
	return nil
}

type NullableSkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue struct {
	value *SkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue
	isSet bool
}

func (v NullableSkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue) Get() *SkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue {
	return v.value
}

func (v *NullableSkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue) Set(val *SkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue) {
	v.value = val
	v.isSet = true
}

func (v NullableSkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue) IsSet() bool {
	return v.isSet
}

func (v *NullableSkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue(val *SkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue) *NullableSkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue {
	return &NullableSkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue{value: val, isSet: true}
}

func (v NullableSkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkyBlockProfileMemberNetherIslandPlayerQuestsQuestRewardsValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
