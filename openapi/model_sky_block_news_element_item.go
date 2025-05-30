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

// checks if the SkyBlockNewsElementItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkyBlockNewsElementItem{}

// SkyBlockNewsElementItem struct for SkyBlockNewsElementItem
type SkyBlockNewsElementItem struct {
	Material string `json:"material"`
}

type _SkyBlockNewsElementItem SkyBlockNewsElementItem

// NewSkyBlockNewsElementItem instantiates a new SkyBlockNewsElementItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkyBlockNewsElementItem(material string) *SkyBlockNewsElementItem {
	this := SkyBlockNewsElementItem{}
	this.Material = material
	return &this
}

// NewSkyBlockNewsElementItemWithDefaults instantiates a new SkyBlockNewsElementItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkyBlockNewsElementItemWithDefaults() *SkyBlockNewsElementItem {
	this := SkyBlockNewsElementItem{}
	return &this
}

// GetMaterial returns the Material field value
func (o *SkyBlockNewsElementItem) GetMaterial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Material
}

// GetMaterialOk returns a tuple with the Material field value
// and a boolean to check if the value has been set.
func (o *SkyBlockNewsElementItem) GetMaterialOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Material, true
}

// SetMaterial sets field value
func (o *SkyBlockNewsElementItem) SetMaterial(v string) {
	o.Material = v
}

func (o SkyBlockNewsElementItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkyBlockNewsElementItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["material"] = o.Material
	return toSerialize, nil
}

func (o *SkyBlockNewsElementItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"material",
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

	varSkyBlockNewsElementItem := _SkyBlockNewsElementItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSkyBlockNewsElementItem)

	if err != nil {
		return err
	}

	*o = SkyBlockNewsElementItem(varSkyBlockNewsElementItem)

	return err
}

type NullableSkyBlockNewsElementItem struct {
	value *SkyBlockNewsElementItem
	isSet bool
}

func (v NullableSkyBlockNewsElementItem) Get() *SkyBlockNewsElementItem {
	return v.value
}

func (v *NullableSkyBlockNewsElementItem) Set(val *SkyBlockNewsElementItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSkyBlockNewsElementItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSkyBlockNewsElementItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkyBlockNewsElementItem(val *SkyBlockNewsElementItem) *NullableSkyBlockNewsElementItem {
	return &NullableSkyBlockNewsElementItem{value: val, isSet: true}
}

func (v NullableSkyBlockNewsElementItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkyBlockNewsElementItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
