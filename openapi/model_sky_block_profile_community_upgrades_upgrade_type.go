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

// SkyBlockProfileCommunityUpgradesUpgradeType the model 'SkyBlockProfileCommunityUpgradesUpgradeType'
type SkyBlockProfileCommunityUpgradesUpgradeType string

// List of SkyBlockProfileCommunityUpgradesUpgradeType
const (
	SKYBLOCKPROFILECOMMUNITYUPGRADESUPGRADETYPE_COINS_ALLOWANCE SkyBlockProfileCommunityUpgradesUpgradeType = "coins_allowance"
	SKYBLOCKPROFILECOMMUNITYUPGRADESUPGRADETYPE_COOP_SLOTS      SkyBlockProfileCommunityUpgradesUpgradeType = "coop_slots"
	SKYBLOCKPROFILECOMMUNITYUPGRADESUPGRADETYPE_GUESTS_COUNT    SkyBlockProfileCommunityUpgradesUpgradeType = "guests_count"
	SKYBLOCKPROFILECOMMUNITYUPGRADESUPGRADETYPE_ISLAND_SIZE     SkyBlockProfileCommunityUpgradesUpgradeType = "island_size"
	SKYBLOCKPROFILECOMMUNITYUPGRADESUPGRADETYPE_MINION_SLOTS    SkyBlockProfileCommunityUpgradesUpgradeType = "minion_slots"
)

// All allowed values of SkyBlockProfileCommunityUpgradesUpgradeType enum
var AllowedSkyBlockProfileCommunityUpgradesUpgradeTypeEnumValues = []SkyBlockProfileCommunityUpgradesUpgradeType{
	"coins_allowance",
	"coop_slots",
	"guests_count",
	"island_size",
	"minion_slots",
}

func (v *SkyBlockProfileCommunityUpgradesUpgradeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SkyBlockProfileCommunityUpgradesUpgradeType(value)
	for _, existing := range AllowedSkyBlockProfileCommunityUpgradesUpgradeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SkyBlockProfileCommunityUpgradesUpgradeType", value)
}

// NewSkyBlockProfileCommunityUpgradesUpgradeTypeFromValue returns a pointer to a valid SkyBlockProfileCommunityUpgradesUpgradeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSkyBlockProfileCommunityUpgradesUpgradeTypeFromValue(v string) (*SkyBlockProfileCommunityUpgradesUpgradeType, error) {
	ev := SkyBlockProfileCommunityUpgradesUpgradeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SkyBlockProfileCommunityUpgradesUpgradeType: valid values are %v", v, AllowedSkyBlockProfileCommunityUpgradesUpgradeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SkyBlockProfileCommunityUpgradesUpgradeType) IsValid() bool {
	for _, existing := range AllowedSkyBlockProfileCommunityUpgradesUpgradeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SkyBlockProfileCommunityUpgradesUpgradeType value
func (v SkyBlockProfileCommunityUpgradesUpgradeType) Ptr() *SkyBlockProfileCommunityUpgradesUpgradeType {
	return &v
}

type NullableSkyBlockProfileCommunityUpgradesUpgradeType struct {
	value *SkyBlockProfileCommunityUpgradesUpgradeType
	isSet bool
}

func (v NullableSkyBlockProfileCommunityUpgradesUpgradeType) Get() *SkyBlockProfileCommunityUpgradesUpgradeType {
	return v.value
}

func (v *NullableSkyBlockProfileCommunityUpgradesUpgradeType) Set(val *SkyBlockProfileCommunityUpgradesUpgradeType) {
	v.value = val
	v.isSet = true
}

func (v NullableSkyBlockProfileCommunityUpgradesUpgradeType) IsSet() bool {
	return v.isSet
}

func (v *NullableSkyBlockProfileCommunityUpgradesUpgradeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkyBlockProfileCommunityUpgradesUpgradeType(val *SkyBlockProfileCommunityUpgradesUpgradeType) *NullableSkyBlockProfileCommunityUpgradesUpgradeType {
	return &NullableSkyBlockProfileCommunityUpgradesUpgradeType{value: val, isSet: true}
}

func (v NullableSkyBlockProfileCommunityUpgradesUpgradeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkyBlockProfileCommunityUpgradesUpgradeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
