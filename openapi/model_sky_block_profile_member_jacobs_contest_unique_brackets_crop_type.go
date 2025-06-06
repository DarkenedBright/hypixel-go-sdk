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

// SkyBlockProfileMemberJacobsContestUniqueBracketsCropType the model 'SkyBlockProfileMemberJacobsContestUniqueBracketsCropType'
type SkyBlockProfileMemberJacobsContestUniqueBracketsCropType string

// List of SkyBlockProfileMemberJacobsContestUniqueBracketsCropType
const (
	SKYBLOCKPROFILEMEMBERJACOBSCONTESTUNIQUEBRACKETSCROPTYPE_CACTUS              SkyBlockProfileMemberJacobsContestUniqueBracketsCropType = "CACTUS"
	SKYBLOCKPROFILEMEMBERJACOBSCONTESTUNIQUEBRACKETSCROPTYPE_CARROT_ITEM         SkyBlockProfileMemberJacobsContestUniqueBracketsCropType = "CARROT_ITEM"
	SKYBLOCKPROFILEMEMBERJACOBSCONTESTUNIQUEBRACKETSCROPTYPE_INK_SACK_3          SkyBlockProfileMemberJacobsContestUniqueBracketsCropType = "INK_SACK:3"
	SKYBLOCKPROFILEMEMBERJACOBSCONTESTUNIQUEBRACKETSCROPTYPE_MELON               SkyBlockProfileMemberJacobsContestUniqueBracketsCropType = "MELON"
	SKYBLOCKPROFILEMEMBERJACOBSCONTESTUNIQUEBRACKETSCROPTYPE_MUSHROOM_COLLECTION SkyBlockProfileMemberJacobsContestUniqueBracketsCropType = "MUSHROOM_COLLECTION"
	SKYBLOCKPROFILEMEMBERJACOBSCONTESTUNIQUEBRACKETSCROPTYPE_NETHER_STALK        SkyBlockProfileMemberJacobsContestUniqueBracketsCropType = "NETHER_STALK"
	SKYBLOCKPROFILEMEMBERJACOBSCONTESTUNIQUEBRACKETSCROPTYPE_POTATO_ITEM         SkyBlockProfileMemberJacobsContestUniqueBracketsCropType = "POTATO_ITEM"
	SKYBLOCKPROFILEMEMBERJACOBSCONTESTUNIQUEBRACKETSCROPTYPE_PUMPKIN             SkyBlockProfileMemberJacobsContestUniqueBracketsCropType = "PUMPKIN"
	SKYBLOCKPROFILEMEMBERJACOBSCONTESTUNIQUEBRACKETSCROPTYPE_SUGAR_CANE          SkyBlockProfileMemberJacobsContestUniqueBracketsCropType = "SUGAR_CANE"
	SKYBLOCKPROFILEMEMBERJACOBSCONTESTUNIQUEBRACKETSCROPTYPE_WHEAT               SkyBlockProfileMemberJacobsContestUniqueBracketsCropType = "WHEAT"
)

// All allowed values of SkyBlockProfileMemberJacobsContestUniqueBracketsCropType enum
var AllowedSkyBlockProfileMemberJacobsContestUniqueBracketsCropTypeEnumValues = []SkyBlockProfileMemberJacobsContestUniqueBracketsCropType{
	"CACTUS",
	"CARROT_ITEM",
	"INK_SACK:3",
	"MELON",
	"MUSHROOM_COLLECTION",
	"NETHER_STALK",
	"POTATO_ITEM",
	"PUMPKIN",
	"SUGAR_CANE",
	"WHEAT",
}

func (v *SkyBlockProfileMemberJacobsContestUniqueBracketsCropType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SkyBlockProfileMemberJacobsContestUniqueBracketsCropType(value)
	for _, existing := range AllowedSkyBlockProfileMemberJacobsContestUniqueBracketsCropTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SkyBlockProfileMemberJacobsContestUniqueBracketsCropType", value)
}

// NewSkyBlockProfileMemberJacobsContestUniqueBracketsCropTypeFromValue returns a pointer to a valid SkyBlockProfileMemberJacobsContestUniqueBracketsCropType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSkyBlockProfileMemberJacobsContestUniqueBracketsCropTypeFromValue(v string) (*SkyBlockProfileMemberJacobsContestUniqueBracketsCropType, error) {
	ev := SkyBlockProfileMemberJacobsContestUniqueBracketsCropType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SkyBlockProfileMemberJacobsContestUniqueBracketsCropType: valid values are %v", v, AllowedSkyBlockProfileMemberJacobsContestUniqueBracketsCropTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SkyBlockProfileMemberJacobsContestUniqueBracketsCropType) IsValid() bool {
	for _, existing := range AllowedSkyBlockProfileMemberJacobsContestUniqueBracketsCropTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SkyBlockProfileMemberJacobsContestUniqueBracketsCropType value
func (v SkyBlockProfileMemberJacobsContestUniqueBracketsCropType) Ptr() *SkyBlockProfileMemberJacobsContestUniqueBracketsCropType {
	return &v
}

type NullableSkyBlockProfileMemberJacobsContestUniqueBracketsCropType struct {
	value *SkyBlockProfileMemberJacobsContestUniqueBracketsCropType
	isSet bool
}

func (v NullableSkyBlockProfileMemberJacobsContestUniqueBracketsCropType) Get() *SkyBlockProfileMemberJacobsContestUniqueBracketsCropType {
	return v.value
}

func (v *NullableSkyBlockProfileMemberJacobsContestUniqueBracketsCropType) Set(val *SkyBlockProfileMemberJacobsContestUniqueBracketsCropType) {
	v.value = val
	v.isSet = true
}

func (v NullableSkyBlockProfileMemberJacobsContestUniqueBracketsCropType) IsSet() bool {
	return v.isSet
}

func (v *NullableSkyBlockProfileMemberJacobsContestUniqueBracketsCropType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkyBlockProfileMemberJacobsContestUniqueBracketsCropType(val *SkyBlockProfileMemberJacobsContestUniqueBracketsCropType) *NullableSkyBlockProfileMemberJacobsContestUniqueBracketsCropType {
	return &NullableSkyBlockProfileMemberJacobsContestUniqueBracketsCropType{value: val, isSet: true}
}

func (v NullableSkyBlockProfileMemberJacobsContestUniqueBracketsCropType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkyBlockProfileMemberJacobsContestUniqueBracketsCropType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
