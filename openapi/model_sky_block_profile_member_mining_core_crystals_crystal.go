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

// checks if the SkyBlockProfileMemberMiningCoreCrystalsCrystal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkyBlockProfileMemberMiningCoreCrystalsCrystal{}

// SkyBlockProfileMemberMiningCoreCrystalsCrystal struct for SkyBlockProfileMemberMiningCoreCrystalsCrystal
type SkyBlockProfileMemberMiningCoreCrystalsCrystal struct {
	State       *SkyBlockProfileMemberMiningCoreCrystalsCrystalState `json:"state,omitempty"`
	TotalFound  *int64                                               `json:"total_found,omitempty"`
	TotalPlaced *int64                                               `json:"total_placed,omitempty"`
}

// NewSkyBlockProfileMemberMiningCoreCrystalsCrystal instantiates a new SkyBlockProfileMemberMiningCoreCrystalsCrystal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkyBlockProfileMemberMiningCoreCrystalsCrystal() *SkyBlockProfileMemberMiningCoreCrystalsCrystal {
	this := SkyBlockProfileMemberMiningCoreCrystalsCrystal{}
	return &this
}

// NewSkyBlockProfileMemberMiningCoreCrystalsCrystalWithDefaults instantiates a new SkyBlockProfileMemberMiningCoreCrystalsCrystal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkyBlockProfileMemberMiningCoreCrystalsCrystalWithDefaults() *SkyBlockProfileMemberMiningCoreCrystalsCrystal {
	this := SkyBlockProfileMemberMiningCoreCrystalsCrystal{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberMiningCoreCrystalsCrystal) GetState() SkyBlockProfileMemberMiningCoreCrystalsCrystalState {
	if o == nil || IsNil(o.State) {
		var ret SkyBlockProfileMemberMiningCoreCrystalsCrystalState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberMiningCoreCrystalsCrystal) GetStateOk() (*SkyBlockProfileMemberMiningCoreCrystalsCrystalState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberMiningCoreCrystalsCrystal) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given SkyBlockProfileMemberMiningCoreCrystalsCrystalState and assigns it to the State field.
func (o *SkyBlockProfileMemberMiningCoreCrystalsCrystal) SetState(v SkyBlockProfileMemberMiningCoreCrystalsCrystalState) {
	o.State = &v
}

// GetTotalFound returns the TotalFound field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberMiningCoreCrystalsCrystal) GetTotalFound() int64 {
	if o == nil || IsNil(o.TotalFound) {
		var ret int64
		return ret
	}
	return *o.TotalFound
}

// GetTotalFoundOk returns a tuple with the TotalFound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberMiningCoreCrystalsCrystal) GetTotalFoundOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalFound) {
		return nil, false
	}
	return o.TotalFound, true
}

// HasTotalFound returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberMiningCoreCrystalsCrystal) HasTotalFound() bool {
	if o != nil && !IsNil(o.TotalFound) {
		return true
	}

	return false
}

// SetTotalFound gets a reference to the given int64 and assigns it to the TotalFound field.
func (o *SkyBlockProfileMemberMiningCoreCrystalsCrystal) SetTotalFound(v int64) {
	o.TotalFound = &v
}

// GetTotalPlaced returns the TotalPlaced field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberMiningCoreCrystalsCrystal) GetTotalPlaced() int64 {
	if o == nil || IsNil(o.TotalPlaced) {
		var ret int64
		return ret
	}
	return *o.TotalPlaced
}

// GetTotalPlacedOk returns a tuple with the TotalPlaced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberMiningCoreCrystalsCrystal) GetTotalPlacedOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalPlaced) {
		return nil, false
	}
	return o.TotalPlaced, true
}

// HasTotalPlaced returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberMiningCoreCrystalsCrystal) HasTotalPlaced() bool {
	if o != nil && !IsNil(o.TotalPlaced) {
		return true
	}

	return false
}

// SetTotalPlaced gets a reference to the given int64 and assigns it to the TotalPlaced field.
func (o *SkyBlockProfileMemberMiningCoreCrystalsCrystal) SetTotalPlaced(v int64) {
	o.TotalPlaced = &v
}

func (o SkyBlockProfileMemberMiningCoreCrystalsCrystal) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkyBlockProfileMemberMiningCoreCrystalsCrystal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.TotalFound) {
		toSerialize["total_found"] = o.TotalFound
	}
	if !IsNil(o.TotalPlaced) {
		toSerialize["total_placed"] = o.TotalPlaced
	}
	return toSerialize, nil
}

type NullableSkyBlockProfileMemberMiningCoreCrystalsCrystal struct {
	value *SkyBlockProfileMemberMiningCoreCrystalsCrystal
	isSet bool
}

func (v NullableSkyBlockProfileMemberMiningCoreCrystalsCrystal) Get() *SkyBlockProfileMemberMiningCoreCrystalsCrystal {
	return v.value
}

func (v *NullableSkyBlockProfileMemberMiningCoreCrystalsCrystal) Set(val *SkyBlockProfileMemberMiningCoreCrystalsCrystal) {
	v.value = val
	v.isSet = true
}

func (v NullableSkyBlockProfileMemberMiningCoreCrystalsCrystal) IsSet() bool {
	return v.isSet
}

func (v *NullableSkyBlockProfileMemberMiningCoreCrystalsCrystal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkyBlockProfileMemberMiningCoreCrystalsCrystal(val *SkyBlockProfileMemberMiningCoreCrystalsCrystal) *NullableSkyBlockProfileMemberMiningCoreCrystalsCrystal {
	return &NullableSkyBlockProfileMemberMiningCoreCrystalsCrystal{value: val, isSet: true}
}

func (v NullableSkyBlockProfileMemberMiningCoreCrystalsCrystal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkyBlockProfileMemberMiningCoreCrystalsCrystal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
