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

// checks if the SkyBlockProfileMemberRiftWestVillage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkyBlockProfileMemberRiftWestVillage{}

// SkyBlockProfileMemberRiftWestVillage struct for SkyBlockProfileMemberRiftWestVillage
type SkyBlockProfileMemberRiftWestVillage struct {
	CrazyKloon  *SkyBlockProfileMemberRiftWestVillageCrazyKloon  `json:"crazy_kloon,omitempty"`
	Glyphs      *SkyBlockProfileMemberRiftWestVillageGlyphs      `json:"glyphs,omitempty"`
	KatHouse    *SkyBlockProfileMemberRiftWestVillageKatHouse    `json:"kat_house,omitempty"`
	Mirrorverse *SkyBlockProfileMemberRiftWestVillageMirrorverse `json:"mirrorverse,omitempty"`
}

// NewSkyBlockProfileMemberRiftWestVillage instantiates a new SkyBlockProfileMemberRiftWestVillage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkyBlockProfileMemberRiftWestVillage() *SkyBlockProfileMemberRiftWestVillage {
	this := SkyBlockProfileMemberRiftWestVillage{}
	return &this
}

// NewSkyBlockProfileMemberRiftWestVillageWithDefaults instantiates a new SkyBlockProfileMemberRiftWestVillage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkyBlockProfileMemberRiftWestVillageWithDefaults() *SkyBlockProfileMemberRiftWestVillage {
	this := SkyBlockProfileMemberRiftWestVillage{}
	return &this
}

// GetCrazyKloon returns the CrazyKloon field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberRiftWestVillage) GetCrazyKloon() SkyBlockProfileMemberRiftWestVillageCrazyKloon {
	if o == nil || IsNil(o.CrazyKloon) {
		var ret SkyBlockProfileMemberRiftWestVillageCrazyKloon
		return ret
	}
	return *o.CrazyKloon
}

// GetCrazyKloonOk returns a tuple with the CrazyKloon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberRiftWestVillage) GetCrazyKloonOk() (*SkyBlockProfileMemberRiftWestVillageCrazyKloon, bool) {
	if o == nil || IsNil(o.CrazyKloon) {
		return nil, false
	}
	return o.CrazyKloon, true
}

// HasCrazyKloon returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberRiftWestVillage) HasCrazyKloon() bool {
	if o != nil && !IsNil(o.CrazyKloon) {
		return true
	}

	return false
}

// SetCrazyKloon gets a reference to the given SkyBlockProfileMemberRiftWestVillageCrazyKloon and assigns it to the CrazyKloon field.
func (o *SkyBlockProfileMemberRiftWestVillage) SetCrazyKloon(v SkyBlockProfileMemberRiftWestVillageCrazyKloon) {
	o.CrazyKloon = &v
}

// GetGlyphs returns the Glyphs field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberRiftWestVillage) GetGlyphs() SkyBlockProfileMemberRiftWestVillageGlyphs {
	if o == nil || IsNil(o.Glyphs) {
		var ret SkyBlockProfileMemberRiftWestVillageGlyphs
		return ret
	}
	return *o.Glyphs
}

// GetGlyphsOk returns a tuple with the Glyphs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberRiftWestVillage) GetGlyphsOk() (*SkyBlockProfileMemberRiftWestVillageGlyphs, bool) {
	if o == nil || IsNil(o.Glyphs) {
		return nil, false
	}
	return o.Glyphs, true
}

// HasGlyphs returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberRiftWestVillage) HasGlyphs() bool {
	if o != nil && !IsNil(o.Glyphs) {
		return true
	}

	return false
}

// SetGlyphs gets a reference to the given SkyBlockProfileMemberRiftWestVillageGlyphs and assigns it to the Glyphs field.
func (o *SkyBlockProfileMemberRiftWestVillage) SetGlyphs(v SkyBlockProfileMemberRiftWestVillageGlyphs) {
	o.Glyphs = &v
}

// GetKatHouse returns the KatHouse field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberRiftWestVillage) GetKatHouse() SkyBlockProfileMemberRiftWestVillageKatHouse {
	if o == nil || IsNil(o.KatHouse) {
		var ret SkyBlockProfileMemberRiftWestVillageKatHouse
		return ret
	}
	return *o.KatHouse
}

// GetKatHouseOk returns a tuple with the KatHouse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberRiftWestVillage) GetKatHouseOk() (*SkyBlockProfileMemberRiftWestVillageKatHouse, bool) {
	if o == nil || IsNil(o.KatHouse) {
		return nil, false
	}
	return o.KatHouse, true
}

// HasKatHouse returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberRiftWestVillage) HasKatHouse() bool {
	if o != nil && !IsNil(o.KatHouse) {
		return true
	}

	return false
}

// SetKatHouse gets a reference to the given SkyBlockProfileMemberRiftWestVillageKatHouse and assigns it to the KatHouse field.
func (o *SkyBlockProfileMemberRiftWestVillage) SetKatHouse(v SkyBlockProfileMemberRiftWestVillageKatHouse) {
	o.KatHouse = &v
}

// GetMirrorverse returns the Mirrorverse field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberRiftWestVillage) GetMirrorverse() SkyBlockProfileMemberRiftWestVillageMirrorverse {
	if o == nil || IsNil(o.Mirrorverse) {
		var ret SkyBlockProfileMemberRiftWestVillageMirrorverse
		return ret
	}
	return *o.Mirrorverse
}

// GetMirrorverseOk returns a tuple with the Mirrorverse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberRiftWestVillage) GetMirrorverseOk() (*SkyBlockProfileMemberRiftWestVillageMirrorverse, bool) {
	if o == nil || IsNil(o.Mirrorverse) {
		return nil, false
	}
	return o.Mirrorverse, true
}

// HasMirrorverse returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberRiftWestVillage) HasMirrorverse() bool {
	if o != nil && !IsNil(o.Mirrorverse) {
		return true
	}

	return false
}

// SetMirrorverse gets a reference to the given SkyBlockProfileMemberRiftWestVillageMirrorverse and assigns it to the Mirrorverse field.
func (o *SkyBlockProfileMemberRiftWestVillage) SetMirrorverse(v SkyBlockProfileMemberRiftWestVillageMirrorverse) {
	o.Mirrorverse = &v
}

func (o SkyBlockProfileMemberRiftWestVillage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkyBlockProfileMemberRiftWestVillage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CrazyKloon) {
		toSerialize["crazy_kloon"] = o.CrazyKloon
	}
	if !IsNil(o.Glyphs) {
		toSerialize["glyphs"] = o.Glyphs
	}
	if !IsNil(o.KatHouse) {
		toSerialize["kat_house"] = o.KatHouse
	}
	if !IsNil(o.Mirrorverse) {
		toSerialize["mirrorverse"] = o.Mirrorverse
	}
	return toSerialize, nil
}

type NullableSkyBlockProfileMemberRiftWestVillage struct {
	value *SkyBlockProfileMemberRiftWestVillage
	isSet bool
}

func (v NullableSkyBlockProfileMemberRiftWestVillage) Get() *SkyBlockProfileMemberRiftWestVillage {
	return v.value
}

func (v *NullableSkyBlockProfileMemberRiftWestVillage) Set(val *SkyBlockProfileMemberRiftWestVillage) {
	v.value = val
	v.isSet = true
}

func (v NullableSkyBlockProfileMemberRiftWestVillage) IsSet() bool {
	return v.isSet
}

func (v *NullableSkyBlockProfileMemberRiftWestVillage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkyBlockProfileMemberRiftWestVillage(val *SkyBlockProfileMemberRiftWestVillage) *NullableSkyBlockProfileMemberRiftWestVillage {
	return &NullableSkyBlockProfileMemberRiftWestVillage{value: val, isSet: true}
}

func (v NullableSkyBlockProfileMemberRiftWestVillage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkyBlockProfileMemberRiftWestVillage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
