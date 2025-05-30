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

// checks if the SkyBlockProfileMemberRiftWestVillageGlyphs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkyBlockProfileMemberRiftWestVillageGlyphs{}

// SkyBlockProfileMemberRiftWestVillageGlyphs struct for SkyBlockProfileMemberRiftWestVillageGlyphs
type SkyBlockProfileMemberRiftWestVillageGlyphs struct {
	ClaimedBracelet       *bool  `json:"claimed_bracelet,omitempty"`
	ClaimedWand           *bool  `json:"claimed_wand,omitempty"`
	Completed             *bool  `json:"completed,omitempty"`
	CurrentGlyph          *int64 `json:"current_glyph,omitempty"`
	CurrentGlyphCompleted *bool  `json:"current_glyph_completed,omitempty"`
	CurrentGlyphDelivered *bool  `json:"current_glyph_delivered,omitempty"`
}

// NewSkyBlockProfileMemberRiftWestVillageGlyphs instantiates a new SkyBlockProfileMemberRiftWestVillageGlyphs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkyBlockProfileMemberRiftWestVillageGlyphs() *SkyBlockProfileMemberRiftWestVillageGlyphs {
	this := SkyBlockProfileMemberRiftWestVillageGlyphs{}
	return &this
}

// NewSkyBlockProfileMemberRiftWestVillageGlyphsWithDefaults instantiates a new SkyBlockProfileMemberRiftWestVillageGlyphs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkyBlockProfileMemberRiftWestVillageGlyphsWithDefaults() *SkyBlockProfileMemberRiftWestVillageGlyphs {
	this := SkyBlockProfileMemberRiftWestVillageGlyphs{}
	return &this
}

// GetClaimedBracelet returns the ClaimedBracelet field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberRiftWestVillageGlyphs) GetClaimedBracelet() bool {
	if o == nil || IsNil(o.ClaimedBracelet) {
		var ret bool
		return ret
	}
	return *o.ClaimedBracelet
}

// GetClaimedBraceletOk returns a tuple with the ClaimedBracelet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberRiftWestVillageGlyphs) GetClaimedBraceletOk() (*bool, bool) {
	if o == nil || IsNil(o.ClaimedBracelet) {
		return nil, false
	}
	return o.ClaimedBracelet, true
}

// HasClaimedBracelet returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberRiftWestVillageGlyphs) HasClaimedBracelet() bool {
	if o != nil && !IsNil(o.ClaimedBracelet) {
		return true
	}

	return false
}

// SetClaimedBracelet gets a reference to the given bool and assigns it to the ClaimedBracelet field.
func (o *SkyBlockProfileMemberRiftWestVillageGlyphs) SetClaimedBracelet(v bool) {
	o.ClaimedBracelet = &v
}

// GetClaimedWand returns the ClaimedWand field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberRiftWestVillageGlyphs) GetClaimedWand() bool {
	if o == nil || IsNil(o.ClaimedWand) {
		var ret bool
		return ret
	}
	return *o.ClaimedWand
}

// GetClaimedWandOk returns a tuple with the ClaimedWand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberRiftWestVillageGlyphs) GetClaimedWandOk() (*bool, bool) {
	if o == nil || IsNil(o.ClaimedWand) {
		return nil, false
	}
	return o.ClaimedWand, true
}

// HasClaimedWand returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberRiftWestVillageGlyphs) HasClaimedWand() bool {
	if o != nil && !IsNil(o.ClaimedWand) {
		return true
	}

	return false
}

// SetClaimedWand gets a reference to the given bool and assigns it to the ClaimedWand field.
func (o *SkyBlockProfileMemberRiftWestVillageGlyphs) SetClaimedWand(v bool) {
	o.ClaimedWand = &v
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberRiftWestVillageGlyphs) GetCompleted() bool {
	if o == nil || IsNil(o.Completed) {
		var ret bool
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberRiftWestVillageGlyphs) GetCompletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Completed) {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberRiftWestVillageGlyphs) HasCompleted() bool {
	if o != nil && !IsNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given bool and assigns it to the Completed field.
func (o *SkyBlockProfileMemberRiftWestVillageGlyphs) SetCompleted(v bool) {
	o.Completed = &v
}

// GetCurrentGlyph returns the CurrentGlyph field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberRiftWestVillageGlyphs) GetCurrentGlyph() int64 {
	if o == nil || IsNil(o.CurrentGlyph) {
		var ret int64
		return ret
	}
	return *o.CurrentGlyph
}

// GetCurrentGlyphOk returns a tuple with the CurrentGlyph field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberRiftWestVillageGlyphs) GetCurrentGlyphOk() (*int64, bool) {
	if o == nil || IsNil(o.CurrentGlyph) {
		return nil, false
	}
	return o.CurrentGlyph, true
}

// HasCurrentGlyph returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberRiftWestVillageGlyphs) HasCurrentGlyph() bool {
	if o != nil && !IsNil(o.CurrentGlyph) {
		return true
	}

	return false
}

// SetCurrentGlyph gets a reference to the given int64 and assigns it to the CurrentGlyph field.
func (o *SkyBlockProfileMemberRiftWestVillageGlyphs) SetCurrentGlyph(v int64) {
	o.CurrentGlyph = &v
}

// GetCurrentGlyphCompleted returns the CurrentGlyphCompleted field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberRiftWestVillageGlyphs) GetCurrentGlyphCompleted() bool {
	if o == nil || IsNil(o.CurrentGlyphCompleted) {
		var ret bool
		return ret
	}
	return *o.CurrentGlyphCompleted
}

// GetCurrentGlyphCompletedOk returns a tuple with the CurrentGlyphCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberRiftWestVillageGlyphs) GetCurrentGlyphCompletedOk() (*bool, bool) {
	if o == nil || IsNil(o.CurrentGlyphCompleted) {
		return nil, false
	}
	return o.CurrentGlyphCompleted, true
}

// HasCurrentGlyphCompleted returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberRiftWestVillageGlyphs) HasCurrentGlyphCompleted() bool {
	if o != nil && !IsNil(o.CurrentGlyphCompleted) {
		return true
	}

	return false
}

// SetCurrentGlyphCompleted gets a reference to the given bool and assigns it to the CurrentGlyphCompleted field.
func (o *SkyBlockProfileMemberRiftWestVillageGlyphs) SetCurrentGlyphCompleted(v bool) {
	o.CurrentGlyphCompleted = &v
}

// GetCurrentGlyphDelivered returns the CurrentGlyphDelivered field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberRiftWestVillageGlyphs) GetCurrentGlyphDelivered() bool {
	if o == nil || IsNil(o.CurrentGlyphDelivered) {
		var ret bool
		return ret
	}
	return *o.CurrentGlyphDelivered
}

// GetCurrentGlyphDeliveredOk returns a tuple with the CurrentGlyphDelivered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberRiftWestVillageGlyphs) GetCurrentGlyphDeliveredOk() (*bool, bool) {
	if o == nil || IsNil(o.CurrentGlyphDelivered) {
		return nil, false
	}
	return o.CurrentGlyphDelivered, true
}

// HasCurrentGlyphDelivered returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberRiftWestVillageGlyphs) HasCurrentGlyphDelivered() bool {
	if o != nil && !IsNil(o.CurrentGlyphDelivered) {
		return true
	}

	return false
}

// SetCurrentGlyphDelivered gets a reference to the given bool and assigns it to the CurrentGlyphDelivered field.
func (o *SkyBlockProfileMemberRiftWestVillageGlyphs) SetCurrentGlyphDelivered(v bool) {
	o.CurrentGlyphDelivered = &v
}

func (o SkyBlockProfileMemberRiftWestVillageGlyphs) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkyBlockProfileMemberRiftWestVillageGlyphs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClaimedBracelet) {
		toSerialize["claimed_bracelet"] = o.ClaimedBracelet
	}
	if !IsNil(o.ClaimedWand) {
		toSerialize["claimed_wand"] = o.ClaimedWand
	}
	if !IsNil(o.Completed) {
		toSerialize["completed"] = o.Completed
	}
	if !IsNil(o.CurrentGlyph) {
		toSerialize["current_glyph"] = o.CurrentGlyph
	}
	if !IsNil(o.CurrentGlyphCompleted) {
		toSerialize["current_glyph_completed"] = o.CurrentGlyphCompleted
	}
	if !IsNil(o.CurrentGlyphDelivered) {
		toSerialize["current_glyph_delivered"] = o.CurrentGlyphDelivered
	}
	return toSerialize, nil
}

type NullableSkyBlockProfileMemberRiftWestVillageGlyphs struct {
	value *SkyBlockProfileMemberRiftWestVillageGlyphs
	isSet bool
}

func (v NullableSkyBlockProfileMemberRiftWestVillageGlyphs) Get() *SkyBlockProfileMemberRiftWestVillageGlyphs {
	return v.value
}

func (v *NullableSkyBlockProfileMemberRiftWestVillageGlyphs) Set(val *SkyBlockProfileMemberRiftWestVillageGlyphs) {
	v.value = val
	v.isSet = true
}

func (v NullableSkyBlockProfileMemberRiftWestVillageGlyphs) IsSet() bool {
	return v.isSet
}

func (v *NullableSkyBlockProfileMemberRiftWestVillageGlyphs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkyBlockProfileMemberRiftWestVillageGlyphs(val *SkyBlockProfileMemberRiftWestVillageGlyphs) *NullableSkyBlockProfileMemberRiftWestVillageGlyphs {
	return &NullableSkyBlockProfileMemberRiftWestVillageGlyphs{value: val, isSet: true}
}

func (v NullableSkyBlockProfileMemberRiftWestVillageGlyphs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkyBlockProfileMemberRiftWestVillageGlyphs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
