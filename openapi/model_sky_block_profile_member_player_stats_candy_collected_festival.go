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

// checks if the SkyBlockProfileMemberPlayerStatsCandyCollectedFestival type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkyBlockProfileMemberPlayerStatsCandyCollectedFestival{}

// SkyBlockProfileMemberPlayerStatsCandyCollectedFestival struct for SkyBlockProfileMemberPlayerStatsCandyCollectedFestival
type SkyBlockProfileMemberPlayerStatsCandyCollectedFestival struct {
	GreenCandy  *int64 `json:"green_candy,omitempty"`
	PurpleCandy *int64 `json:"purple_candy,omitempty"`
	Total       int64  `json:"total"`
}

type _SkyBlockProfileMemberPlayerStatsCandyCollectedFestival SkyBlockProfileMemberPlayerStatsCandyCollectedFestival

// NewSkyBlockProfileMemberPlayerStatsCandyCollectedFestival instantiates a new SkyBlockProfileMemberPlayerStatsCandyCollectedFestival object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkyBlockProfileMemberPlayerStatsCandyCollectedFestival(total int64) *SkyBlockProfileMemberPlayerStatsCandyCollectedFestival {
	this := SkyBlockProfileMemberPlayerStatsCandyCollectedFestival{}
	this.Total = total
	return &this
}

// NewSkyBlockProfileMemberPlayerStatsCandyCollectedFestivalWithDefaults instantiates a new SkyBlockProfileMemberPlayerStatsCandyCollectedFestival object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkyBlockProfileMemberPlayerStatsCandyCollectedFestivalWithDefaults() *SkyBlockProfileMemberPlayerStatsCandyCollectedFestival {
	this := SkyBlockProfileMemberPlayerStatsCandyCollectedFestival{}
	return &this
}

// GetGreenCandy returns the GreenCandy field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStatsCandyCollectedFestival) GetGreenCandy() int64 {
	if o == nil || IsNil(o.GreenCandy) {
		var ret int64
		return ret
	}
	return *o.GreenCandy
}

// GetGreenCandyOk returns a tuple with the GreenCandy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStatsCandyCollectedFestival) GetGreenCandyOk() (*int64, bool) {
	if o == nil || IsNil(o.GreenCandy) {
		return nil, false
	}
	return o.GreenCandy, true
}

// HasGreenCandy returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStatsCandyCollectedFestival) HasGreenCandy() bool {
	if o != nil && !IsNil(o.GreenCandy) {
		return true
	}

	return false
}

// SetGreenCandy gets a reference to the given int64 and assigns it to the GreenCandy field.
func (o *SkyBlockProfileMemberPlayerStatsCandyCollectedFestival) SetGreenCandy(v int64) {
	o.GreenCandy = &v
}

// GetPurpleCandy returns the PurpleCandy field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStatsCandyCollectedFestival) GetPurpleCandy() int64 {
	if o == nil || IsNil(o.PurpleCandy) {
		var ret int64
		return ret
	}
	return *o.PurpleCandy
}

// GetPurpleCandyOk returns a tuple with the PurpleCandy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStatsCandyCollectedFestival) GetPurpleCandyOk() (*int64, bool) {
	if o == nil || IsNil(o.PurpleCandy) {
		return nil, false
	}
	return o.PurpleCandy, true
}

// HasPurpleCandy returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStatsCandyCollectedFestival) HasPurpleCandy() bool {
	if o != nil && !IsNil(o.PurpleCandy) {
		return true
	}

	return false
}

// SetPurpleCandy gets a reference to the given int64 and assigns it to the PurpleCandy field.
func (o *SkyBlockProfileMemberPlayerStatsCandyCollectedFestival) SetPurpleCandy(v int64) {
	o.PurpleCandy = &v
}

// GetTotal returns the Total field value
func (o *SkyBlockProfileMemberPlayerStatsCandyCollectedFestival) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStatsCandyCollectedFestival) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *SkyBlockProfileMemberPlayerStatsCandyCollectedFestival) SetTotal(v int64) {
	o.Total = v
}

func (o SkyBlockProfileMemberPlayerStatsCandyCollectedFestival) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkyBlockProfileMemberPlayerStatsCandyCollectedFestival) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GreenCandy) {
		toSerialize["green_candy"] = o.GreenCandy
	}
	if !IsNil(o.PurpleCandy) {
		toSerialize["purple_candy"] = o.PurpleCandy
	}
	toSerialize["total"] = o.Total
	return toSerialize, nil
}

func (o *SkyBlockProfileMemberPlayerStatsCandyCollectedFestival) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"total",
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

	varSkyBlockProfileMemberPlayerStatsCandyCollectedFestival := _SkyBlockProfileMemberPlayerStatsCandyCollectedFestival{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSkyBlockProfileMemberPlayerStatsCandyCollectedFestival)

	if err != nil {
		return err
	}

	*o = SkyBlockProfileMemberPlayerStatsCandyCollectedFestival(varSkyBlockProfileMemberPlayerStatsCandyCollectedFestival)

	return err
}

type NullableSkyBlockProfileMemberPlayerStatsCandyCollectedFestival struct {
	value *SkyBlockProfileMemberPlayerStatsCandyCollectedFestival
	isSet bool
}

func (v NullableSkyBlockProfileMemberPlayerStatsCandyCollectedFestival) Get() *SkyBlockProfileMemberPlayerStatsCandyCollectedFestival {
	return v.value
}

func (v *NullableSkyBlockProfileMemberPlayerStatsCandyCollectedFestival) Set(val *SkyBlockProfileMemberPlayerStatsCandyCollectedFestival) {
	v.value = val
	v.isSet = true
}

func (v NullableSkyBlockProfileMemberPlayerStatsCandyCollectedFestival) IsSet() bool {
	return v.isSet
}

func (v *NullableSkyBlockProfileMemberPlayerStatsCandyCollectedFestival) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkyBlockProfileMemberPlayerStatsCandyCollectedFestival(val *SkyBlockProfileMemberPlayerStatsCandyCollectedFestival) *NullableSkyBlockProfileMemberPlayerStatsCandyCollectedFestival {
	return &NullableSkyBlockProfileMemberPlayerStatsCandyCollectedFestival{value: val, isSet: true}
}

func (v NullableSkyBlockProfileMemberPlayerStatsCandyCollectedFestival) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkyBlockProfileMemberPlayerStatsCandyCollectedFestival) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
