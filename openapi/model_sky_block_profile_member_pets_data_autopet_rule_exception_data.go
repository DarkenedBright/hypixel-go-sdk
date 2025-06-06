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

// checks if the SkyBlockProfileMemberPetsDataAutopetRuleExceptionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkyBlockProfileMemberPetsDataAutopetRuleExceptionData{}

// SkyBlockProfileMemberPetsDataAutopetRuleExceptionData struct for SkyBlockProfileMemberPetsDataAutopetRuleExceptionData
type SkyBlockProfileMemberPetsDataAutopetRuleExceptionData struct {
	Event  *string `json:"event,omitempty"`
	Island *string `json:"island,omitempty"`
	Pet    *string `json:"pet,omitempty"`
}

// NewSkyBlockProfileMemberPetsDataAutopetRuleExceptionData instantiates a new SkyBlockProfileMemberPetsDataAutopetRuleExceptionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkyBlockProfileMemberPetsDataAutopetRuleExceptionData() *SkyBlockProfileMemberPetsDataAutopetRuleExceptionData {
	this := SkyBlockProfileMemberPetsDataAutopetRuleExceptionData{}
	return &this
}

// NewSkyBlockProfileMemberPetsDataAutopetRuleExceptionDataWithDefaults instantiates a new SkyBlockProfileMemberPetsDataAutopetRuleExceptionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkyBlockProfileMemberPetsDataAutopetRuleExceptionDataWithDefaults() *SkyBlockProfileMemberPetsDataAutopetRuleExceptionData {
	this := SkyBlockProfileMemberPetsDataAutopetRuleExceptionData{}
	return &this
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleExceptionData) GetEvent() string {
	if o == nil || IsNil(o.Event) {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleExceptionData) GetEventOk() (*string, bool) {
	if o == nil || IsNil(o.Event) {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleExceptionData) HasEvent() bool {
	if o != nil && !IsNil(o.Event) {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleExceptionData) SetEvent(v string) {
	o.Event = &v
}

// GetIsland returns the Island field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleExceptionData) GetIsland() string {
	if o == nil || IsNil(o.Island) {
		var ret string
		return ret
	}
	return *o.Island
}

// GetIslandOk returns a tuple with the Island field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleExceptionData) GetIslandOk() (*string, bool) {
	if o == nil || IsNil(o.Island) {
		return nil, false
	}
	return o.Island, true
}

// HasIsland returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleExceptionData) HasIsland() bool {
	if o != nil && !IsNil(o.Island) {
		return true
	}

	return false
}

// SetIsland gets a reference to the given string and assigns it to the Island field.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleExceptionData) SetIsland(v string) {
	o.Island = &v
}

// GetPet returns the Pet field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleExceptionData) GetPet() string {
	if o == nil || IsNil(o.Pet) {
		var ret string
		return ret
	}
	return *o.Pet
}

// GetPetOk returns a tuple with the Pet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleExceptionData) GetPetOk() (*string, bool) {
	if o == nil || IsNil(o.Pet) {
		return nil, false
	}
	return o.Pet, true
}

// HasPet returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleExceptionData) HasPet() bool {
	if o != nil && !IsNil(o.Pet) {
		return true
	}

	return false
}

// SetPet gets a reference to the given string and assigns it to the Pet field.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleExceptionData) SetPet(v string) {
	o.Pet = &v
}

func (o SkyBlockProfileMemberPetsDataAutopetRuleExceptionData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkyBlockProfileMemberPetsDataAutopetRuleExceptionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Event) {
		toSerialize["event"] = o.Event
	}
	if !IsNil(o.Island) {
		toSerialize["island"] = o.Island
	}
	if !IsNil(o.Pet) {
		toSerialize["pet"] = o.Pet
	}
	return toSerialize, nil
}

type NullableSkyBlockProfileMemberPetsDataAutopetRuleExceptionData struct {
	value *SkyBlockProfileMemberPetsDataAutopetRuleExceptionData
	isSet bool
}

func (v NullableSkyBlockProfileMemberPetsDataAutopetRuleExceptionData) Get() *SkyBlockProfileMemberPetsDataAutopetRuleExceptionData {
	return v.value
}

func (v *NullableSkyBlockProfileMemberPetsDataAutopetRuleExceptionData) Set(val *SkyBlockProfileMemberPetsDataAutopetRuleExceptionData) {
	v.value = val
	v.isSet = true
}

func (v NullableSkyBlockProfileMemberPetsDataAutopetRuleExceptionData) IsSet() bool {
	return v.isSet
}

func (v *NullableSkyBlockProfileMemberPetsDataAutopetRuleExceptionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkyBlockProfileMemberPetsDataAutopetRuleExceptionData(val *SkyBlockProfileMemberPetsDataAutopetRuleExceptionData) *NullableSkyBlockProfileMemberPetsDataAutopetRuleExceptionData {
	return &NullableSkyBlockProfileMemberPetsDataAutopetRuleExceptionData{value: val, isSet: true}
}

func (v NullableSkyBlockProfileMemberPetsDataAutopetRuleExceptionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkyBlockProfileMemberPetsDataAutopetRuleExceptionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
