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

// checks if the SkyBlockProfileMemberRiftWestVillageCrazyKloon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkyBlockProfileMemberRiftWestVillageCrazyKloon{}

// SkyBlockProfileMemberRiftWestVillageCrazyKloon struct for SkyBlockProfileMemberRiftWestVillageCrazyKloon
type SkyBlockProfileMemberRiftWestVillageCrazyKloon struct {
	HackedTerminals []SkyBlockProfileMemberRiftWestVillageCrazyKloonHackedTerminalsInner `json:"hacked_terminals,omitempty"`
	QuestComplete   *bool                                                                `json:"quest_complete,omitempty"`
	SelectedColors  SkyBlockProfileMemberRiftWestVillageCrazyKloonSelectedColors         `json:"selected_colors"`
	Talked          *bool                                                                `json:"talked,omitempty"`
}

type _SkyBlockProfileMemberRiftWestVillageCrazyKloon SkyBlockProfileMemberRiftWestVillageCrazyKloon

// NewSkyBlockProfileMemberRiftWestVillageCrazyKloon instantiates a new SkyBlockProfileMemberRiftWestVillageCrazyKloon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkyBlockProfileMemberRiftWestVillageCrazyKloon(selectedColors SkyBlockProfileMemberRiftWestVillageCrazyKloonSelectedColors) *SkyBlockProfileMemberRiftWestVillageCrazyKloon {
	this := SkyBlockProfileMemberRiftWestVillageCrazyKloon{}
	this.SelectedColors = selectedColors
	return &this
}

// NewSkyBlockProfileMemberRiftWestVillageCrazyKloonWithDefaults instantiates a new SkyBlockProfileMemberRiftWestVillageCrazyKloon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkyBlockProfileMemberRiftWestVillageCrazyKloonWithDefaults() *SkyBlockProfileMemberRiftWestVillageCrazyKloon {
	this := SkyBlockProfileMemberRiftWestVillageCrazyKloon{}
	return &this
}

// GetHackedTerminals returns the HackedTerminals field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberRiftWestVillageCrazyKloon) GetHackedTerminals() []SkyBlockProfileMemberRiftWestVillageCrazyKloonHackedTerminalsInner {
	if o == nil || IsNil(o.HackedTerminals) {
		var ret []SkyBlockProfileMemberRiftWestVillageCrazyKloonHackedTerminalsInner
		return ret
	}
	return o.HackedTerminals
}

// GetHackedTerminalsOk returns a tuple with the HackedTerminals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberRiftWestVillageCrazyKloon) GetHackedTerminalsOk() ([]SkyBlockProfileMemberRiftWestVillageCrazyKloonHackedTerminalsInner, bool) {
	if o == nil || IsNil(o.HackedTerminals) {
		return nil, false
	}
	return o.HackedTerminals, true
}

// HasHackedTerminals returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberRiftWestVillageCrazyKloon) HasHackedTerminals() bool {
	if o != nil && !IsNil(o.HackedTerminals) {
		return true
	}

	return false
}

// SetHackedTerminals gets a reference to the given []SkyBlockProfileMemberRiftWestVillageCrazyKloonHackedTerminalsInner and assigns it to the HackedTerminals field.
func (o *SkyBlockProfileMemberRiftWestVillageCrazyKloon) SetHackedTerminals(v []SkyBlockProfileMemberRiftWestVillageCrazyKloonHackedTerminalsInner) {
	o.HackedTerminals = v
}

// GetQuestComplete returns the QuestComplete field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberRiftWestVillageCrazyKloon) GetQuestComplete() bool {
	if o == nil || IsNil(o.QuestComplete) {
		var ret bool
		return ret
	}
	return *o.QuestComplete
}

// GetQuestCompleteOk returns a tuple with the QuestComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberRiftWestVillageCrazyKloon) GetQuestCompleteOk() (*bool, bool) {
	if o == nil || IsNil(o.QuestComplete) {
		return nil, false
	}
	return o.QuestComplete, true
}

// HasQuestComplete returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberRiftWestVillageCrazyKloon) HasQuestComplete() bool {
	if o != nil && !IsNil(o.QuestComplete) {
		return true
	}

	return false
}

// SetQuestComplete gets a reference to the given bool and assigns it to the QuestComplete field.
func (o *SkyBlockProfileMemberRiftWestVillageCrazyKloon) SetQuestComplete(v bool) {
	o.QuestComplete = &v
}

// GetSelectedColors returns the SelectedColors field value
func (o *SkyBlockProfileMemberRiftWestVillageCrazyKloon) GetSelectedColors() SkyBlockProfileMemberRiftWestVillageCrazyKloonSelectedColors {
	if o == nil {
		var ret SkyBlockProfileMemberRiftWestVillageCrazyKloonSelectedColors
		return ret
	}

	return o.SelectedColors
}

// GetSelectedColorsOk returns a tuple with the SelectedColors field value
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberRiftWestVillageCrazyKloon) GetSelectedColorsOk() (*SkyBlockProfileMemberRiftWestVillageCrazyKloonSelectedColors, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelectedColors, true
}

// SetSelectedColors sets field value
func (o *SkyBlockProfileMemberRiftWestVillageCrazyKloon) SetSelectedColors(v SkyBlockProfileMemberRiftWestVillageCrazyKloonSelectedColors) {
	o.SelectedColors = v
}

// GetTalked returns the Talked field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberRiftWestVillageCrazyKloon) GetTalked() bool {
	if o == nil || IsNil(o.Talked) {
		var ret bool
		return ret
	}
	return *o.Talked
}

// GetTalkedOk returns a tuple with the Talked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberRiftWestVillageCrazyKloon) GetTalkedOk() (*bool, bool) {
	if o == nil || IsNil(o.Talked) {
		return nil, false
	}
	return o.Talked, true
}

// HasTalked returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberRiftWestVillageCrazyKloon) HasTalked() bool {
	if o != nil && !IsNil(o.Talked) {
		return true
	}

	return false
}

// SetTalked gets a reference to the given bool and assigns it to the Talked field.
func (o *SkyBlockProfileMemberRiftWestVillageCrazyKloon) SetTalked(v bool) {
	o.Talked = &v
}

func (o SkyBlockProfileMemberRiftWestVillageCrazyKloon) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkyBlockProfileMemberRiftWestVillageCrazyKloon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HackedTerminals) {
		toSerialize["hacked_terminals"] = o.HackedTerminals
	}
	if !IsNil(o.QuestComplete) {
		toSerialize["quest_complete"] = o.QuestComplete
	}
	toSerialize["selected_colors"] = o.SelectedColors
	if !IsNil(o.Talked) {
		toSerialize["talked"] = o.Talked
	}
	return toSerialize, nil
}

func (o *SkyBlockProfileMemberRiftWestVillageCrazyKloon) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"selected_colors",
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

	varSkyBlockProfileMemberRiftWestVillageCrazyKloon := _SkyBlockProfileMemberRiftWestVillageCrazyKloon{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSkyBlockProfileMemberRiftWestVillageCrazyKloon)

	if err != nil {
		return err
	}

	*o = SkyBlockProfileMemberRiftWestVillageCrazyKloon(varSkyBlockProfileMemberRiftWestVillageCrazyKloon)

	return err
}

type NullableSkyBlockProfileMemberRiftWestVillageCrazyKloon struct {
	value *SkyBlockProfileMemberRiftWestVillageCrazyKloon
	isSet bool
}

func (v NullableSkyBlockProfileMemberRiftWestVillageCrazyKloon) Get() *SkyBlockProfileMemberRiftWestVillageCrazyKloon {
	return v.value
}

func (v *NullableSkyBlockProfileMemberRiftWestVillageCrazyKloon) Set(val *SkyBlockProfileMemberRiftWestVillageCrazyKloon) {
	v.value = val
	v.isSet = true
}

func (v NullableSkyBlockProfileMemberRiftWestVillageCrazyKloon) IsSet() bool {
	return v.isSet
}

func (v *NullableSkyBlockProfileMemberRiftWestVillageCrazyKloon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkyBlockProfileMemberRiftWestVillageCrazyKloon(val *SkyBlockProfileMemberRiftWestVillageCrazyKloon) *NullableSkyBlockProfileMemberRiftWestVillageCrazyKloon {
	return &NullableSkyBlockProfileMemberRiftWestVillageCrazyKloon{value: val, isSet: true}
}

func (v NullableSkyBlockProfileMemberRiftWestVillageCrazyKloon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkyBlockProfileMemberRiftWestVillageCrazyKloon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
