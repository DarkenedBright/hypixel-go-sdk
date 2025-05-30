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

// checks if the SkyBlockProfileMemberAccessoryBagStorageTuningSlot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkyBlockProfileMemberAccessoryBagStorageTuningSlot{}

// SkyBlockProfileMemberAccessoryBagStorageTuningSlot struct for SkyBlockProfileMemberAccessoryBagStorageTuningSlot
type SkyBlockProfileMemberAccessoryBagStorageTuningSlot struct {
	AttackSpeed    *int64 `json:"attack_speed,omitempty"`
	CriticalChance *int64 `json:"critical_chance,omitempty"`
	CriticalDamage *int64 `json:"critical_damage,omitempty"`
	Defense        *int64 `json:"defense,omitempty"`
	Health         *int64 `json:"health,omitempty"`
	Intelligence   *int64 `json:"intelligence,omitempty"`
	PurchaseTs     *int64 `json:"purchase_ts,omitempty"`
	Strength       *int64 `json:"strength,omitempty"`
	WalkSpeed      *int64 `json:"walk_speed,omitempty"`
}

// NewSkyBlockProfileMemberAccessoryBagStorageTuningSlot instantiates a new SkyBlockProfileMemberAccessoryBagStorageTuningSlot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkyBlockProfileMemberAccessoryBagStorageTuningSlot() *SkyBlockProfileMemberAccessoryBagStorageTuningSlot {
	this := SkyBlockProfileMemberAccessoryBagStorageTuningSlot{}
	return &this
}

// NewSkyBlockProfileMemberAccessoryBagStorageTuningSlotWithDefaults instantiates a new SkyBlockProfileMemberAccessoryBagStorageTuningSlot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkyBlockProfileMemberAccessoryBagStorageTuningSlotWithDefaults() *SkyBlockProfileMemberAccessoryBagStorageTuningSlot {
	this := SkyBlockProfileMemberAccessoryBagStorageTuningSlot{}
	return &this
}

// GetAttackSpeed returns the AttackSpeed field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) GetAttackSpeed() int64 {
	if o == nil || IsNil(o.AttackSpeed) {
		var ret int64
		return ret
	}
	return *o.AttackSpeed
}

// GetAttackSpeedOk returns a tuple with the AttackSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) GetAttackSpeedOk() (*int64, bool) {
	if o == nil || IsNil(o.AttackSpeed) {
		return nil, false
	}
	return o.AttackSpeed, true
}

// HasAttackSpeed returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) HasAttackSpeed() bool {
	if o != nil && !IsNil(o.AttackSpeed) {
		return true
	}

	return false
}

// SetAttackSpeed gets a reference to the given int64 and assigns it to the AttackSpeed field.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) SetAttackSpeed(v int64) {
	o.AttackSpeed = &v
}

// GetCriticalChance returns the CriticalChance field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) GetCriticalChance() int64 {
	if o == nil || IsNil(o.CriticalChance) {
		var ret int64
		return ret
	}
	return *o.CriticalChance
}

// GetCriticalChanceOk returns a tuple with the CriticalChance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) GetCriticalChanceOk() (*int64, bool) {
	if o == nil || IsNil(o.CriticalChance) {
		return nil, false
	}
	return o.CriticalChance, true
}

// HasCriticalChance returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) HasCriticalChance() bool {
	if o != nil && !IsNil(o.CriticalChance) {
		return true
	}

	return false
}

// SetCriticalChance gets a reference to the given int64 and assigns it to the CriticalChance field.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) SetCriticalChance(v int64) {
	o.CriticalChance = &v
}

// GetCriticalDamage returns the CriticalDamage field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) GetCriticalDamage() int64 {
	if o == nil || IsNil(o.CriticalDamage) {
		var ret int64
		return ret
	}
	return *o.CriticalDamage
}

// GetCriticalDamageOk returns a tuple with the CriticalDamage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) GetCriticalDamageOk() (*int64, bool) {
	if o == nil || IsNil(o.CriticalDamage) {
		return nil, false
	}
	return o.CriticalDamage, true
}

// HasCriticalDamage returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) HasCriticalDamage() bool {
	if o != nil && !IsNil(o.CriticalDamage) {
		return true
	}

	return false
}

// SetCriticalDamage gets a reference to the given int64 and assigns it to the CriticalDamage field.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) SetCriticalDamage(v int64) {
	o.CriticalDamage = &v
}

// GetDefense returns the Defense field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) GetDefense() int64 {
	if o == nil || IsNil(o.Defense) {
		var ret int64
		return ret
	}
	return *o.Defense
}

// GetDefenseOk returns a tuple with the Defense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) GetDefenseOk() (*int64, bool) {
	if o == nil || IsNil(o.Defense) {
		return nil, false
	}
	return o.Defense, true
}

// HasDefense returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) HasDefense() bool {
	if o != nil && !IsNil(o.Defense) {
		return true
	}

	return false
}

// SetDefense gets a reference to the given int64 and assigns it to the Defense field.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) SetDefense(v int64) {
	o.Defense = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) GetHealth() int64 {
	if o == nil || IsNil(o.Health) {
		var ret int64
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) GetHealthOk() (*int64, bool) {
	if o == nil || IsNil(o.Health) {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) HasHealth() bool {
	if o != nil && !IsNil(o.Health) {
		return true
	}

	return false
}

// SetHealth gets a reference to the given int64 and assigns it to the Health field.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) SetHealth(v int64) {
	o.Health = &v
}

// GetIntelligence returns the Intelligence field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) GetIntelligence() int64 {
	if o == nil || IsNil(o.Intelligence) {
		var ret int64
		return ret
	}
	return *o.Intelligence
}

// GetIntelligenceOk returns a tuple with the Intelligence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) GetIntelligenceOk() (*int64, bool) {
	if o == nil || IsNil(o.Intelligence) {
		return nil, false
	}
	return o.Intelligence, true
}

// HasIntelligence returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) HasIntelligence() bool {
	if o != nil && !IsNil(o.Intelligence) {
		return true
	}

	return false
}

// SetIntelligence gets a reference to the given int64 and assigns it to the Intelligence field.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) SetIntelligence(v int64) {
	o.Intelligence = &v
}

// GetPurchaseTs returns the PurchaseTs field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) GetPurchaseTs() int64 {
	if o == nil || IsNil(o.PurchaseTs) {
		var ret int64
		return ret
	}
	return *o.PurchaseTs
}

// GetPurchaseTsOk returns a tuple with the PurchaseTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) GetPurchaseTsOk() (*int64, bool) {
	if o == nil || IsNil(o.PurchaseTs) {
		return nil, false
	}
	return o.PurchaseTs, true
}

// HasPurchaseTs returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) HasPurchaseTs() bool {
	if o != nil && !IsNil(o.PurchaseTs) {
		return true
	}

	return false
}

// SetPurchaseTs gets a reference to the given int64 and assigns it to the PurchaseTs field.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) SetPurchaseTs(v int64) {
	o.PurchaseTs = &v
}

// GetStrength returns the Strength field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) GetStrength() int64 {
	if o == nil || IsNil(o.Strength) {
		var ret int64
		return ret
	}
	return *o.Strength
}

// GetStrengthOk returns a tuple with the Strength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) GetStrengthOk() (*int64, bool) {
	if o == nil || IsNil(o.Strength) {
		return nil, false
	}
	return o.Strength, true
}

// HasStrength returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) HasStrength() bool {
	if o != nil && !IsNil(o.Strength) {
		return true
	}

	return false
}

// SetStrength gets a reference to the given int64 and assigns it to the Strength field.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) SetStrength(v int64) {
	o.Strength = &v
}

// GetWalkSpeed returns the WalkSpeed field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) GetWalkSpeed() int64 {
	if o == nil || IsNil(o.WalkSpeed) {
		var ret int64
		return ret
	}
	return *o.WalkSpeed
}

// GetWalkSpeedOk returns a tuple with the WalkSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) GetWalkSpeedOk() (*int64, bool) {
	if o == nil || IsNil(o.WalkSpeed) {
		return nil, false
	}
	return o.WalkSpeed, true
}

// HasWalkSpeed returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) HasWalkSpeed() bool {
	if o != nil && !IsNil(o.WalkSpeed) {
		return true
	}

	return false
}

// SetWalkSpeed gets a reference to the given int64 and assigns it to the WalkSpeed field.
func (o *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) SetWalkSpeed(v int64) {
	o.WalkSpeed = &v
}

func (o SkyBlockProfileMemberAccessoryBagStorageTuningSlot) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkyBlockProfileMemberAccessoryBagStorageTuningSlot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttackSpeed) {
		toSerialize["attack_speed"] = o.AttackSpeed
	}
	if !IsNil(o.CriticalChance) {
		toSerialize["critical_chance"] = o.CriticalChance
	}
	if !IsNil(o.CriticalDamage) {
		toSerialize["critical_damage"] = o.CriticalDamage
	}
	if !IsNil(o.Defense) {
		toSerialize["defense"] = o.Defense
	}
	if !IsNil(o.Health) {
		toSerialize["health"] = o.Health
	}
	if !IsNil(o.Intelligence) {
		toSerialize["intelligence"] = o.Intelligence
	}
	if !IsNil(o.PurchaseTs) {
		toSerialize["purchase_ts"] = o.PurchaseTs
	}
	if !IsNil(o.Strength) {
		toSerialize["strength"] = o.Strength
	}
	if !IsNil(o.WalkSpeed) {
		toSerialize["walk_speed"] = o.WalkSpeed
	}
	return toSerialize, nil
}

type NullableSkyBlockProfileMemberAccessoryBagStorageTuningSlot struct {
	value *SkyBlockProfileMemberAccessoryBagStorageTuningSlot
	isSet bool
}

func (v NullableSkyBlockProfileMemberAccessoryBagStorageTuningSlot) Get() *SkyBlockProfileMemberAccessoryBagStorageTuningSlot {
	return v.value
}

func (v *NullableSkyBlockProfileMemberAccessoryBagStorageTuningSlot) Set(val *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) {
	v.value = val
	v.isSet = true
}

func (v NullableSkyBlockProfileMemberAccessoryBagStorageTuningSlot) IsSet() bool {
	return v.isSet
}

func (v *NullableSkyBlockProfileMemberAccessoryBagStorageTuningSlot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkyBlockProfileMemberAccessoryBagStorageTuningSlot(val *SkyBlockProfileMemberAccessoryBagStorageTuningSlot) *NullableSkyBlockProfileMemberAccessoryBagStorageTuningSlot {
	return &NullableSkyBlockProfileMemberAccessoryBagStorageTuningSlot{value: val, isSet: true}
}

func (v NullableSkyBlockProfileMemberAccessoryBagStorageTuningSlot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkyBlockProfileMemberAccessoryBagStorageTuningSlot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
