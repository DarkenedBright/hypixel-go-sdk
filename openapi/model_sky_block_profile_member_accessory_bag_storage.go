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

// checks if the SkyBlockProfileMemberAccessoryBagStorage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkyBlockProfileMemberAccessoryBagStorage{}

// SkyBlockProfileMemberAccessoryBagStorage struct for SkyBlockProfileMemberAccessoryBagStorage
type SkyBlockProfileMemberAccessoryBagStorage struct {
	BagUpgradesPurchased *int64                                          `json:"bag_upgrades_purchased,omitempty"`
	HighestMagicalPower  *int64                                          `json:"highest_magical_power,omitempty"`
	SelectedPower        *SkyBlockProfileMemberAccessoryBagStoragePower  `json:"selected_power,omitempty"`
	Tuning               SkyBlockProfileMemberAccessoryBagStorageTuning  `json:"tuning"`
	UnlockedPowers       []SkyBlockProfileMemberAccessoryBagStoragePower `json:"unlocked_powers,omitempty"`
}

type _SkyBlockProfileMemberAccessoryBagStorage SkyBlockProfileMemberAccessoryBagStorage

// NewSkyBlockProfileMemberAccessoryBagStorage instantiates a new SkyBlockProfileMemberAccessoryBagStorage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkyBlockProfileMemberAccessoryBagStorage(tuning SkyBlockProfileMemberAccessoryBagStorageTuning) *SkyBlockProfileMemberAccessoryBagStorage {
	this := SkyBlockProfileMemberAccessoryBagStorage{}
	this.Tuning = tuning
	return &this
}

// NewSkyBlockProfileMemberAccessoryBagStorageWithDefaults instantiates a new SkyBlockProfileMemberAccessoryBagStorage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkyBlockProfileMemberAccessoryBagStorageWithDefaults() *SkyBlockProfileMemberAccessoryBagStorage {
	this := SkyBlockProfileMemberAccessoryBagStorage{}
	return &this
}

// GetBagUpgradesPurchased returns the BagUpgradesPurchased field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberAccessoryBagStorage) GetBagUpgradesPurchased() int64 {
	if o == nil || IsNil(o.BagUpgradesPurchased) {
		var ret int64
		return ret
	}
	return *o.BagUpgradesPurchased
}

// GetBagUpgradesPurchasedOk returns a tuple with the BagUpgradesPurchased field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberAccessoryBagStorage) GetBagUpgradesPurchasedOk() (*int64, bool) {
	if o == nil || IsNil(o.BagUpgradesPurchased) {
		return nil, false
	}
	return o.BagUpgradesPurchased, true
}

// HasBagUpgradesPurchased returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberAccessoryBagStorage) HasBagUpgradesPurchased() bool {
	if o != nil && !IsNil(o.BagUpgradesPurchased) {
		return true
	}

	return false
}

// SetBagUpgradesPurchased gets a reference to the given int64 and assigns it to the BagUpgradesPurchased field.
func (o *SkyBlockProfileMemberAccessoryBagStorage) SetBagUpgradesPurchased(v int64) {
	o.BagUpgradesPurchased = &v
}

// GetHighestMagicalPower returns the HighestMagicalPower field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberAccessoryBagStorage) GetHighestMagicalPower() int64 {
	if o == nil || IsNil(o.HighestMagicalPower) {
		var ret int64
		return ret
	}
	return *o.HighestMagicalPower
}

// GetHighestMagicalPowerOk returns a tuple with the HighestMagicalPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberAccessoryBagStorage) GetHighestMagicalPowerOk() (*int64, bool) {
	if o == nil || IsNil(o.HighestMagicalPower) {
		return nil, false
	}
	return o.HighestMagicalPower, true
}

// HasHighestMagicalPower returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberAccessoryBagStorage) HasHighestMagicalPower() bool {
	if o != nil && !IsNil(o.HighestMagicalPower) {
		return true
	}

	return false
}

// SetHighestMagicalPower gets a reference to the given int64 and assigns it to the HighestMagicalPower field.
func (o *SkyBlockProfileMemberAccessoryBagStorage) SetHighestMagicalPower(v int64) {
	o.HighestMagicalPower = &v
}

// GetSelectedPower returns the SelectedPower field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberAccessoryBagStorage) GetSelectedPower() SkyBlockProfileMemberAccessoryBagStoragePower {
	if o == nil || IsNil(o.SelectedPower) {
		var ret SkyBlockProfileMemberAccessoryBagStoragePower
		return ret
	}
	return *o.SelectedPower
}

// GetSelectedPowerOk returns a tuple with the SelectedPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberAccessoryBagStorage) GetSelectedPowerOk() (*SkyBlockProfileMemberAccessoryBagStoragePower, bool) {
	if o == nil || IsNil(o.SelectedPower) {
		return nil, false
	}
	return o.SelectedPower, true
}

// HasSelectedPower returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberAccessoryBagStorage) HasSelectedPower() bool {
	if o != nil && !IsNil(o.SelectedPower) {
		return true
	}

	return false
}

// SetSelectedPower gets a reference to the given SkyBlockProfileMemberAccessoryBagStoragePower and assigns it to the SelectedPower field.
func (o *SkyBlockProfileMemberAccessoryBagStorage) SetSelectedPower(v SkyBlockProfileMemberAccessoryBagStoragePower) {
	o.SelectedPower = &v
}

// GetTuning returns the Tuning field value
func (o *SkyBlockProfileMemberAccessoryBagStorage) GetTuning() SkyBlockProfileMemberAccessoryBagStorageTuning {
	if o == nil {
		var ret SkyBlockProfileMemberAccessoryBagStorageTuning
		return ret
	}

	return o.Tuning
}

// GetTuningOk returns a tuple with the Tuning field value
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberAccessoryBagStorage) GetTuningOk() (*SkyBlockProfileMemberAccessoryBagStorageTuning, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tuning, true
}

// SetTuning sets field value
func (o *SkyBlockProfileMemberAccessoryBagStorage) SetTuning(v SkyBlockProfileMemberAccessoryBagStorageTuning) {
	o.Tuning = v
}

// GetUnlockedPowers returns the UnlockedPowers field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberAccessoryBagStorage) GetUnlockedPowers() []SkyBlockProfileMemberAccessoryBagStoragePower {
	if o == nil || IsNil(o.UnlockedPowers) {
		var ret []SkyBlockProfileMemberAccessoryBagStoragePower
		return ret
	}
	return o.UnlockedPowers
}

// GetUnlockedPowersOk returns a tuple with the UnlockedPowers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberAccessoryBagStorage) GetUnlockedPowersOk() ([]SkyBlockProfileMemberAccessoryBagStoragePower, bool) {
	if o == nil || IsNil(o.UnlockedPowers) {
		return nil, false
	}
	return o.UnlockedPowers, true
}

// HasUnlockedPowers returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberAccessoryBagStorage) HasUnlockedPowers() bool {
	if o != nil && !IsNil(o.UnlockedPowers) {
		return true
	}

	return false
}

// SetUnlockedPowers gets a reference to the given []SkyBlockProfileMemberAccessoryBagStoragePower and assigns it to the UnlockedPowers field.
func (o *SkyBlockProfileMemberAccessoryBagStorage) SetUnlockedPowers(v []SkyBlockProfileMemberAccessoryBagStoragePower) {
	o.UnlockedPowers = v
}

func (o SkyBlockProfileMemberAccessoryBagStorage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkyBlockProfileMemberAccessoryBagStorage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BagUpgradesPurchased) {
		toSerialize["bag_upgrades_purchased"] = o.BagUpgradesPurchased
	}
	if !IsNil(o.HighestMagicalPower) {
		toSerialize["highest_magical_power"] = o.HighestMagicalPower
	}
	if !IsNil(o.SelectedPower) {
		toSerialize["selected_power"] = o.SelectedPower
	}
	toSerialize["tuning"] = o.Tuning
	if !IsNil(o.UnlockedPowers) {
		toSerialize["unlocked_powers"] = o.UnlockedPowers
	}
	return toSerialize, nil
}

func (o *SkyBlockProfileMemberAccessoryBagStorage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tuning",
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

	varSkyBlockProfileMemberAccessoryBagStorage := _SkyBlockProfileMemberAccessoryBagStorage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSkyBlockProfileMemberAccessoryBagStorage)

	if err != nil {
		return err
	}

	*o = SkyBlockProfileMemberAccessoryBagStorage(varSkyBlockProfileMemberAccessoryBagStorage)

	return err
}

type NullableSkyBlockProfileMemberAccessoryBagStorage struct {
	value *SkyBlockProfileMemberAccessoryBagStorage
	isSet bool
}

func (v NullableSkyBlockProfileMemberAccessoryBagStorage) Get() *SkyBlockProfileMemberAccessoryBagStorage {
	return v.value
}

func (v *NullableSkyBlockProfileMemberAccessoryBagStorage) Set(val *SkyBlockProfileMemberAccessoryBagStorage) {
	v.value = val
	v.isSet = true
}

func (v NullableSkyBlockProfileMemberAccessoryBagStorage) IsSet() bool {
	return v.isSet
}

func (v *NullableSkyBlockProfileMemberAccessoryBagStorage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkyBlockProfileMemberAccessoryBagStorage(val *SkyBlockProfileMemberAccessoryBagStorage) *NullableSkyBlockProfileMemberAccessoryBagStorage {
	return &NullableSkyBlockProfileMemberAccessoryBagStorage{value: val, isSet: true}
}

func (v NullableSkyBlockProfileMemberAccessoryBagStorage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkyBlockProfileMemberAccessoryBagStorage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
