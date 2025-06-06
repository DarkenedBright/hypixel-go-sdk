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

// checks if the SkyBlockProfileMemberPetsDataAutopetRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkyBlockProfileMemberPetsDataAutopetRule{}

// SkyBlockProfileMemberPetsDataAutopetRule struct for SkyBlockProfileMemberPetsDataAutopetRule
type SkyBlockProfileMemberPetsDataAutopetRule struct {
	Data       SkyBlockProfileMemberPetsDataAutopetRuleData        `json:"data"`
	Disabled   bool                                                `json:"disabled"`
	Exceptions []SkyBlockProfileMemberPetsDataAutopetRuleException `json:"exceptions"`
	Id         SkyBlockProfileMemberPetsDataAutopetRuleId          `json:"id"`
	Name       *string                                             `json:"name,omitempty"`
	Pet        *SkyBlockProfileMemberPetsDataPetType               `json:"pet,omitempty"`
	UniqueId   *string                                             `json:"uniqueId,omitempty"`
	Uuid       *string                                             `json:"uuid,omitempty"`
}

type _SkyBlockProfileMemberPetsDataAutopetRule SkyBlockProfileMemberPetsDataAutopetRule

// NewSkyBlockProfileMemberPetsDataAutopetRule instantiates a new SkyBlockProfileMemberPetsDataAutopetRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkyBlockProfileMemberPetsDataAutopetRule(data SkyBlockProfileMemberPetsDataAutopetRuleData, disabled bool, exceptions []SkyBlockProfileMemberPetsDataAutopetRuleException, id SkyBlockProfileMemberPetsDataAutopetRuleId) *SkyBlockProfileMemberPetsDataAutopetRule {
	this := SkyBlockProfileMemberPetsDataAutopetRule{}
	this.Data = data
	this.Disabled = disabled
	this.Exceptions = exceptions
	this.Id = id
	return &this
}

// NewSkyBlockProfileMemberPetsDataAutopetRuleWithDefaults instantiates a new SkyBlockProfileMemberPetsDataAutopetRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkyBlockProfileMemberPetsDataAutopetRuleWithDefaults() *SkyBlockProfileMemberPetsDataAutopetRule {
	this := SkyBlockProfileMemberPetsDataAutopetRule{}
	return &this
}

// GetData returns the Data field value
func (o *SkyBlockProfileMemberPetsDataAutopetRule) GetData() SkyBlockProfileMemberPetsDataAutopetRuleData {
	if o == nil {
		var ret SkyBlockProfileMemberPetsDataAutopetRuleData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRule) GetDataOk() (*SkyBlockProfileMemberPetsDataAutopetRuleData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SkyBlockProfileMemberPetsDataAutopetRule) SetData(v SkyBlockProfileMemberPetsDataAutopetRuleData) {
	o.Data = v
}

// GetDisabled returns the Disabled field value
func (o *SkyBlockProfileMemberPetsDataAutopetRule) GetDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRule) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disabled, true
}

// SetDisabled sets field value
func (o *SkyBlockProfileMemberPetsDataAutopetRule) SetDisabled(v bool) {
	o.Disabled = v
}

// GetExceptions returns the Exceptions field value
func (o *SkyBlockProfileMemberPetsDataAutopetRule) GetExceptions() []SkyBlockProfileMemberPetsDataAutopetRuleException {
	if o == nil {
		var ret []SkyBlockProfileMemberPetsDataAutopetRuleException
		return ret
	}

	return o.Exceptions
}

// GetExceptionsOk returns a tuple with the Exceptions field value
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRule) GetExceptionsOk() ([]SkyBlockProfileMemberPetsDataAutopetRuleException, bool) {
	if o == nil {
		return nil, false
	}
	return o.Exceptions, true
}

// SetExceptions sets field value
func (o *SkyBlockProfileMemberPetsDataAutopetRule) SetExceptions(v []SkyBlockProfileMemberPetsDataAutopetRuleException) {
	o.Exceptions = v
}

// GetId returns the Id field value
func (o *SkyBlockProfileMemberPetsDataAutopetRule) GetId() SkyBlockProfileMemberPetsDataAutopetRuleId {
	if o == nil {
		var ret SkyBlockProfileMemberPetsDataAutopetRuleId
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRule) GetIdOk() (*SkyBlockProfileMemberPetsDataAutopetRuleId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SkyBlockProfileMemberPetsDataAutopetRule) SetId(v SkyBlockProfileMemberPetsDataAutopetRuleId) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPetsDataAutopetRule) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRule) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRule) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SkyBlockProfileMemberPetsDataAutopetRule) SetName(v string) {
	o.Name = &v
}

// GetPet returns the Pet field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPetsDataAutopetRule) GetPet() SkyBlockProfileMemberPetsDataPetType {
	if o == nil || IsNil(o.Pet) {
		var ret SkyBlockProfileMemberPetsDataPetType
		return ret
	}
	return *o.Pet
}

// GetPetOk returns a tuple with the Pet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRule) GetPetOk() (*SkyBlockProfileMemberPetsDataPetType, bool) {
	if o == nil || IsNil(o.Pet) {
		return nil, false
	}
	return o.Pet, true
}

// HasPet returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRule) HasPet() bool {
	if o != nil && !IsNil(o.Pet) {
		return true
	}

	return false
}

// SetPet gets a reference to the given SkyBlockProfileMemberPetsDataPetType and assigns it to the Pet field.
func (o *SkyBlockProfileMemberPetsDataAutopetRule) SetPet(v SkyBlockProfileMemberPetsDataPetType) {
	o.Pet = &v
}

// GetUniqueId returns the UniqueId field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPetsDataAutopetRule) GetUniqueId() string {
	if o == nil || IsNil(o.UniqueId) {
		var ret string
		return ret
	}
	return *o.UniqueId
}

// GetUniqueIdOk returns a tuple with the UniqueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRule) GetUniqueIdOk() (*string, bool) {
	if o == nil || IsNil(o.UniqueId) {
		return nil, false
	}
	return o.UniqueId, true
}

// HasUniqueId returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRule) HasUniqueId() bool {
	if o != nil && !IsNil(o.UniqueId) {
		return true
	}

	return false
}

// SetUniqueId gets a reference to the given string and assigns it to the UniqueId field.
func (o *SkyBlockProfileMemberPetsDataAutopetRule) SetUniqueId(v string) {
	o.UniqueId = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPetsDataAutopetRule) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRule) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRule) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *SkyBlockProfileMemberPetsDataAutopetRule) SetUuid(v string) {
	o.Uuid = &v
}

func (o SkyBlockProfileMemberPetsDataAutopetRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkyBlockProfileMemberPetsDataAutopetRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["disabled"] = o.Disabled
	toSerialize["exceptions"] = o.Exceptions
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Pet) {
		toSerialize["pet"] = o.Pet
	}
	if !IsNil(o.UniqueId) {
		toSerialize["uniqueId"] = o.UniqueId
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	return toSerialize, nil
}

func (o *SkyBlockProfileMemberPetsDataAutopetRule) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"disabled",
		"exceptions",
		"id",
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

	varSkyBlockProfileMemberPetsDataAutopetRule := _SkyBlockProfileMemberPetsDataAutopetRule{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSkyBlockProfileMemberPetsDataAutopetRule)

	if err != nil {
		return err
	}

	*o = SkyBlockProfileMemberPetsDataAutopetRule(varSkyBlockProfileMemberPetsDataAutopetRule)

	return err
}

type NullableSkyBlockProfileMemberPetsDataAutopetRule struct {
	value *SkyBlockProfileMemberPetsDataAutopetRule
	isSet bool
}

func (v NullableSkyBlockProfileMemberPetsDataAutopetRule) Get() *SkyBlockProfileMemberPetsDataAutopetRule {
	return v.value
}

func (v *NullableSkyBlockProfileMemberPetsDataAutopetRule) Set(val *SkyBlockProfileMemberPetsDataAutopetRule) {
	v.value = val
	v.isSet = true
}

func (v NullableSkyBlockProfileMemberPetsDataAutopetRule) IsSet() bool {
	return v.isSet
}

func (v *NullableSkyBlockProfileMemberPetsDataAutopetRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkyBlockProfileMemberPetsDataAutopetRule(val *SkyBlockProfileMemberPetsDataAutopetRule) *NullableSkyBlockProfileMemberPetsDataAutopetRule {
	return &NullableSkyBlockProfileMemberPetsDataAutopetRule{value: val, isSet: true}
}

func (v NullableSkyBlockProfileMemberPetsDataAutopetRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkyBlockProfileMemberPetsDataAutopetRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
