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

// checks if the SkyBlockProfileMemberSharedInventory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkyBlockProfileMemberSharedInventory{}

// SkyBlockProfileMemberSharedInventory struct for SkyBlockProfileMemberSharedInventory
type SkyBlockProfileMemberSharedInventory struct {
	CandyInventoryContents        *SkyBlockProfileMemberInventoryBase64GzipData `json:"candy_inventory_contents,omitempty"`
	CarnivalMaskInventoryContents *SkyBlockProfileMemberInventoryBase64GzipData `json:"carnival_mask_inventory_contents,omitempty"`
}

// NewSkyBlockProfileMemberSharedInventory instantiates a new SkyBlockProfileMemberSharedInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkyBlockProfileMemberSharedInventory() *SkyBlockProfileMemberSharedInventory {
	this := SkyBlockProfileMemberSharedInventory{}
	return &this
}

// NewSkyBlockProfileMemberSharedInventoryWithDefaults instantiates a new SkyBlockProfileMemberSharedInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkyBlockProfileMemberSharedInventoryWithDefaults() *SkyBlockProfileMemberSharedInventory {
	this := SkyBlockProfileMemberSharedInventory{}
	return &this
}

// GetCandyInventoryContents returns the CandyInventoryContents field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberSharedInventory) GetCandyInventoryContents() SkyBlockProfileMemberInventoryBase64GzipData {
	if o == nil || IsNil(o.CandyInventoryContents) {
		var ret SkyBlockProfileMemberInventoryBase64GzipData
		return ret
	}
	return *o.CandyInventoryContents
}

// GetCandyInventoryContentsOk returns a tuple with the CandyInventoryContents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberSharedInventory) GetCandyInventoryContentsOk() (*SkyBlockProfileMemberInventoryBase64GzipData, bool) {
	if o == nil || IsNil(o.CandyInventoryContents) {
		return nil, false
	}
	return o.CandyInventoryContents, true
}

// HasCandyInventoryContents returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberSharedInventory) HasCandyInventoryContents() bool {
	if o != nil && !IsNil(o.CandyInventoryContents) {
		return true
	}

	return false
}

// SetCandyInventoryContents gets a reference to the given SkyBlockProfileMemberInventoryBase64GzipData and assigns it to the CandyInventoryContents field.
func (o *SkyBlockProfileMemberSharedInventory) SetCandyInventoryContents(v SkyBlockProfileMemberInventoryBase64GzipData) {
	o.CandyInventoryContents = &v
}

// GetCarnivalMaskInventoryContents returns the CarnivalMaskInventoryContents field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberSharedInventory) GetCarnivalMaskInventoryContents() SkyBlockProfileMemberInventoryBase64GzipData {
	if o == nil || IsNil(o.CarnivalMaskInventoryContents) {
		var ret SkyBlockProfileMemberInventoryBase64GzipData
		return ret
	}
	return *o.CarnivalMaskInventoryContents
}

// GetCarnivalMaskInventoryContentsOk returns a tuple with the CarnivalMaskInventoryContents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberSharedInventory) GetCarnivalMaskInventoryContentsOk() (*SkyBlockProfileMemberInventoryBase64GzipData, bool) {
	if o == nil || IsNil(o.CarnivalMaskInventoryContents) {
		return nil, false
	}
	return o.CarnivalMaskInventoryContents, true
}

// HasCarnivalMaskInventoryContents returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberSharedInventory) HasCarnivalMaskInventoryContents() bool {
	if o != nil && !IsNil(o.CarnivalMaskInventoryContents) {
		return true
	}

	return false
}

// SetCarnivalMaskInventoryContents gets a reference to the given SkyBlockProfileMemberInventoryBase64GzipData and assigns it to the CarnivalMaskInventoryContents field.
func (o *SkyBlockProfileMemberSharedInventory) SetCarnivalMaskInventoryContents(v SkyBlockProfileMemberInventoryBase64GzipData) {
	o.CarnivalMaskInventoryContents = &v
}

func (o SkyBlockProfileMemberSharedInventory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkyBlockProfileMemberSharedInventory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CandyInventoryContents) {
		toSerialize["candy_inventory_contents"] = o.CandyInventoryContents
	}
	if !IsNil(o.CarnivalMaskInventoryContents) {
		toSerialize["carnival_mask_inventory_contents"] = o.CarnivalMaskInventoryContents
	}
	return toSerialize, nil
}

type NullableSkyBlockProfileMemberSharedInventory struct {
	value *SkyBlockProfileMemberSharedInventory
	isSet bool
}

func (v NullableSkyBlockProfileMemberSharedInventory) Get() *SkyBlockProfileMemberSharedInventory {
	return v.value
}

func (v *NullableSkyBlockProfileMemberSharedInventory) Set(val *SkyBlockProfileMemberSharedInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableSkyBlockProfileMemberSharedInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableSkyBlockProfileMemberSharedInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkyBlockProfileMemberSharedInventory(val *SkyBlockProfileMemberSharedInventory) *NullableSkyBlockProfileMemberSharedInventory {
	return &NullableSkyBlockProfileMemberSharedInventory{value: val, isSet: true}
}

func (v NullableSkyBlockProfileMemberSharedInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkyBlockProfileMemberSharedInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
