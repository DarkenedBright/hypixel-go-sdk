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

// checks if the SkyBlockProfileMemberPlayerStatsRiftShenItemBought type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkyBlockProfileMemberPlayerStatsRiftShenItemBought{}

// SkyBlockProfileMemberPlayerStatsRiftShenItemBought struct for SkyBlockProfileMemberPlayerStatsRiftShenItemBought
type SkyBlockProfileMemberPlayerStatsRiftShenItemBought struct {
	ANCIENT_ELEVATOR          *float64 `json:"ANCIENT_ELEVATOR,omitempty"`
	DEAD_CAT_FOOD             *float64 `json:"DEAD_CAT_FOOD,omitempty"`
	FLEX_HELMET               *float64 `json:"FLEX_HELMET,omitempty"`
	PLACEABLE_FAIRY_SOUL_RIFT *float64 `json:"PLACEABLE_FAIRY_SOUL_RIFT,omitempty"`
	PUNCHCARD_ARTIFACT        *float64 `json:"PUNCHCARD_ARTIFACT,omitempty"`
	RIFT_PRISM                *float64 `json:"RIFT_PRISM,omitempty"`
	Total                     float64  `json:"total"`
}

type _SkyBlockProfileMemberPlayerStatsRiftShenItemBought SkyBlockProfileMemberPlayerStatsRiftShenItemBought

// NewSkyBlockProfileMemberPlayerStatsRiftShenItemBought instantiates a new SkyBlockProfileMemberPlayerStatsRiftShenItemBought object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkyBlockProfileMemberPlayerStatsRiftShenItemBought(total float64) *SkyBlockProfileMemberPlayerStatsRiftShenItemBought {
	this := SkyBlockProfileMemberPlayerStatsRiftShenItemBought{}
	this.Total = total
	return &this
}

// NewSkyBlockProfileMemberPlayerStatsRiftShenItemBoughtWithDefaults instantiates a new SkyBlockProfileMemberPlayerStatsRiftShenItemBought object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkyBlockProfileMemberPlayerStatsRiftShenItemBoughtWithDefaults() *SkyBlockProfileMemberPlayerStatsRiftShenItemBought {
	this := SkyBlockProfileMemberPlayerStatsRiftShenItemBought{}
	return &this
}

// GetANCIENT_ELEVATOR returns the ANCIENT_ELEVATOR field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStatsRiftShenItemBought) GetANCIENT_ELEVATOR() float64 {
	if o == nil || IsNil(o.ANCIENT_ELEVATOR) {
		var ret float64
		return ret
	}
	return *o.ANCIENT_ELEVATOR
}

// GetANCIENT_ELEVATOROk returns a tuple with the ANCIENT_ELEVATOR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStatsRiftShenItemBought) GetANCIENT_ELEVATOROk() (*float64, bool) {
	if o == nil || IsNil(o.ANCIENT_ELEVATOR) {
		return nil, false
	}
	return o.ANCIENT_ELEVATOR, true
}

// HasANCIENT_ELEVATOR returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStatsRiftShenItemBought) HasANCIENT_ELEVATOR() bool {
	if o != nil && !IsNil(o.ANCIENT_ELEVATOR) {
		return true
	}

	return false
}

// SetANCIENT_ELEVATOR gets a reference to the given float64 and assigns it to the ANCIENT_ELEVATOR field.
func (o *SkyBlockProfileMemberPlayerStatsRiftShenItemBought) SetANCIENT_ELEVATOR(v float64) {
	o.ANCIENT_ELEVATOR = &v
}

// GetDEAD_CAT_FOOD returns the DEAD_CAT_FOOD field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStatsRiftShenItemBought) GetDEAD_CAT_FOOD() float64 {
	if o == nil || IsNil(o.DEAD_CAT_FOOD) {
		var ret float64
		return ret
	}
	return *o.DEAD_CAT_FOOD
}

// GetDEAD_CAT_FOODOk returns a tuple with the DEAD_CAT_FOOD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStatsRiftShenItemBought) GetDEAD_CAT_FOODOk() (*float64, bool) {
	if o == nil || IsNil(o.DEAD_CAT_FOOD) {
		return nil, false
	}
	return o.DEAD_CAT_FOOD, true
}

// HasDEAD_CAT_FOOD returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStatsRiftShenItemBought) HasDEAD_CAT_FOOD() bool {
	if o != nil && !IsNil(o.DEAD_CAT_FOOD) {
		return true
	}

	return false
}

// SetDEAD_CAT_FOOD gets a reference to the given float64 and assigns it to the DEAD_CAT_FOOD field.
func (o *SkyBlockProfileMemberPlayerStatsRiftShenItemBought) SetDEAD_CAT_FOOD(v float64) {
	o.DEAD_CAT_FOOD = &v
}

// GetFLEX_HELMET returns the FLEX_HELMET field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStatsRiftShenItemBought) GetFLEX_HELMET() float64 {
	if o == nil || IsNil(o.FLEX_HELMET) {
		var ret float64
		return ret
	}
	return *o.FLEX_HELMET
}

// GetFLEX_HELMETOk returns a tuple with the FLEX_HELMET field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStatsRiftShenItemBought) GetFLEX_HELMETOk() (*float64, bool) {
	if o == nil || IsNil(o.FLEX_HELMET) {
		return nil, false
	}
	return o.FLEX_HELMET, true
}

// HasFLEX_HELMET returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStatsRiftShenItemBought) HasFLEX_HELMET() bool {
	if o != nil && !IsNil(o.FLEX_HELMET) {
		return true
	}

	return false
}

// SetFLEX_HELMET gets a reference to the given float64 and assigns it to the FLEX_HELMET field.
func (o *SkyBlockProfileMemberPlayerStatsRiftShenItemBought) SetFLEX_HELMET(v float64) {
	o.FLEX_HELMET = &v
}

// GetPLACEABLE_FAIRY_SOUL_RIFT returns the PLACEABLE_FAIRY_SOUL_RIFT field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStatsRiftShenItemBought) GetPLACEABLE_FAIRY_SOUL_RIFT() float64 {
	if o == nil || IsNil(o.PLACEABLE_FAIRY_SOUL_RIFT) {
		var ret float64
		return ret
	}
	return *o.PLACEABLE_FAIRY_SOUL_RIFT
}

// GetPLACEABLE_FAIRY_SOUL_RIFTOk returns a tuple with the PLACEABLE_FAIRY_SOUL_RIFT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStatsRiftShenItemBought) GetPLACEABLE_FAIRY_SOUL_RIFTOk() (*float64, bool) {
	if o == nil || IsNil(o.PLACEABLE_FAIRY_SOUL_RIFT) {
		return nil, false
	}
	return o.PLACEABLE_FAIRY_SOUL_RIFT, true
}

// HasPLACEABLE_FAIRY_SOUL_RIFT returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStatsRiftShenItemBought) HasPLACEABLE_FAIRY_SOUL_RIFT() bool {
	if o != nil && !IsNil(o.PLACEABLE_FAIRY_SOUL_RIFT) {
		return true
	}

	return false
}

// SetPLACEABLE_FAIRY_SOUL_RIFT gets a reference to the given float64 and assigns it to the PLACEABLE_FAIRY_SOUL_RIFT field.
func (o *SkyBlockProfileMemberPlayerStatsRiftShenItemBought) SetPLACEABLE_FAIRY_SOUL_RIFT(v float64) {
	o.PLACEABLE_FAIRY_SOUL_RIFT = &v
}

// GetPUNCHCARD_ARTIFACT returns the PUNCHCARD_ARTIFACT field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStatsRiftShenItemBought) GetPUNCHCARD_ARTIFACT() float64 {
	if o == nil || IsNil(o.PUNCHCARD_ARTIFACT) {
		var ret float64
		return ret
	}
	return *o.PUNCHCARD_ARTIFACT
}

// GetPUNCHCARD_ARTIFACTOk returns a tuple with the PUNCHCARD_ARTIFACT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStatsRiftShenItemBought) GetPUNCHCARD_ARTIFACTOk() (*float64, bool) {
	if o == nil || IsNil(o.PUNCHCARD_ARTIFACT) {
		return nil, false
	}
	return o.PUNCHCARD_ARTIFACT, true
}

// HasPUNCHCARD_ARTIFACT returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStatsRiftShenItemBought) HasPUNCHCARD_ARTIFACT() bool {
	if o != nil && !IsNil(o.PUNCHCARD_ARTIFACT) {
		return true
	}

	return false
}

// SetPUNCHCARD_ARTIFACT gets a reference to the given float64 and assigns it to the PUNCHCARD_ARTIFACT field.
func (o *SkyBlockProfileMemberPlayerStatsRiftShenItemBought) SetPUNCHCARD_ARTIFACT(v float64) {
	o.PUNCHCARD_ARTIFACT = &v
}

// GetRIFT_PRISM returns the RIFT_PRISM field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStatsRiftShenItemBought) GetRIFT_PRISM() float64 {
	if o == nil || IsNil(o.RIFT_PRISM) {
		var ret float64
		return ret
	}
	return *o.RIFT_PRISM
}

// GetRIFT_PRISMOk returns a tuple with the RIFT_PRISM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStatsRiftShenItemBought) GetRIFT_PRISMOk() (*float64, bool) {
	if o == nil || IsNil(o.RIFT_PRISM) {
		return nil, false
	}
	return o.RIFT_PRISM, true
}

// HasRIFT_PRISM returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStatsRiftShenItemBought) HasRIFT_PRISM() bool {
	if o != nil && !IsNil(o.RIFT_PRISM) {
		return true
	}

	return false
}

// SetRIFT_PRISM gets a reference to the given float64 and assigns it to the RIFT_PRISM field.
func (o *SkyBlockProfileMemberPlayerStatsRiftShenItemBought) SetRIFT_PRISM(v float64) {
	o.RIFT_PRISM = &v
}

// GetTotal returns the Total field value
func (o *SkyBlockProfileMemberPlayerStatsRiftShenItemBought) GetTotal() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStatsRiftShenItemBought) GetTotalOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *SkyBlockProfileMemberPlayerStatsRiftShenItemBought) SetTotal(v float64) {
	o.Total = v
}

func (o SkyBlockProfileMemberPlayerStatsRiftShenItemBought) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkyBlockProfileMemberPlayerStatsRiftShenItemBought) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ANCIENT_ELEVATOR) {
		toSerialize["ANCIENT_ELEVATOR"] = o.ANCIENT_ELEVATOR
	}
	if !IsNil(o.DEAD_CAT_FOOD) {
		toSerialize["DEAD_CAT_FOOD"] = o.DEAD_CAT_FOOD
	}
	if !IsNil(o.FLEX_HELMET) {
		toSerialize["FLEX_HELMET"] = o.FLEX_HELMET
	}
	if !IsNil(o.PLACEABLE_FAIRY_SOUL_RIFT) {
		toSerialize["PLACEABLE_FAIRY_SOUL_RIFT"] = o.PLACEABLE_FAIRY_SOUL_RIFT
	}
	if !IsNil(o.PUNCHCARD_ARTIFACT) {
		toSerialize["PUNCHCARD_ARTIFACT"] = o.PUNCHCARD_ARTIFACT
	}
	if !IsNil(o.RIFT_PRISM) {
		toSerialize["RIFT_PRISM"] = o.RIFT_PRISM
	}
	toSerialize["total"] = o.Total
	return toSerialize, nil
}

func (o *SkyBlockProfileMemberPlayerStatsRiftShenItemBought) UnmarshalJSON(data []byte) (err error) {
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

	varSkyBlockProfileMemberPlayerStatsRiftShenItemBought := _SkyBlockProfileMemberPlayerStatsRiftShenItemBought{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSkyBlockProfileMemberPlayerStatsRiftShenItemBought)

	if err != nil {
		return err
	}

	*o = SkyBlockProfileMemberPlayerStatsRiftShenItemBought(varSkyBlockProfileMemberPlayerStatsRiftShenItemBought)

	return err
}

type NullableSkyBlockProfileMemberPlayerStatsRiftShenItemBought struct {
	value *SkyBlockProfileMemberPlayerStatsRiftShenItemBought
	isSet bool
}

func (v NullableSkyBlockProfileMemberPlayerStatsRiftShenItemBought) Get() *SkyBlockProfileMemberPlayerStatsRiftShenItemBought {
	return v.value
}

func (v *NullableSkyBlockProfileMemberPlayerStatsRiftShenItemBought) Set(val *SkyBlockProfileMemberPlayerStatsRiftShenItemBought) {
	v.value = val
	v.isSet = true
}

func (v NullableSkyBlockProfileMemberPlayerStatsRiftShenItemBought) IsSet() bool {
	return v.isSet
}

func (v *NullableSkyBlockProfileMemberPlayerStatsRiftShenItemBought) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkyBlockProfileMemberPlayerStatsRiftShenItemBought(val *SkyBlockProfileMemberPlayerStatsRiftShenItemBought) *NullableSkyBlockProfileMemberPlayerStatsRiftShenItemBought {
	return &NullableSkyBlockProfileMemberPlayerStatsRiftShenItemBought{value: val, isSet: true}
}

func (v NullableSkyBlockProfileMemberPlayerStatsRiftShenItemBought) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkyBlockProfileMemberPlayerStatsRiftShenItemBought) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
