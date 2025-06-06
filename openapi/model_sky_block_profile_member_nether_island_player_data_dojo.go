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

// checks if the SkyBlockProfileMemberNetherIslandPlayerDataDojo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkyBlockProfileMemberNetherIslandPlayerDataDojo{}

// SkyBlockProfileMemberNetherIslandPlayerDataDojo struct for SkyBlockProfileMemberNetherIslandPlayerDataDojo
type SkyBlockProfileMemberNetherIslandPlayerDataDojo struct {
	DojoPointsArcher    *int64 `json:"dojo_points_archer,omitempty"`
	DojoPointsFireball  *int64 `json:"dojo_points_fireball,omitempty"`
	DojoPointsLockHead  *int64 `json:"dojo_points_lock_head,omitempty"`
	DojoPointsMobKb     *int64 `json:"dojo_points_mob_kb,omitempty"`
	DojoPointsSnake     *int64 `json:"dojo_points_snake,omitempty"`
	DojoPointsSwordSwap *int64 `json:"dojo_points_sword_swap,omitempty"`
	DojoPointsWallJump  *int64 `json:"dojo_points_wall_jump,omitempty"`
	DojoTimeArcher      *int64 `json:"dojo_time_archer,omitempty"`
	DojoTimeFireball    *int64 `json:"dojo_time_fireball,omitempty"`
	DojoTimeLockHead    *int64 `json:"dojo_time_lock_head,omitempty"`
	DojoTimeMobKb       *int64 `json:"dojo_time_mob_kb,omitempty"`
	DojoTimeSnake       *int64 `json:"dojo_time_snake,omitempty"`
	DojoTimeSwordSwap   *int64 `json:"dojo_time_sword_swap,omitempty"`
	DojoTimeWallJump    *int64 `json:"dojo_time_wall_jump,omitempty"`
}

// NewSkyBlockProfileMemberNetherIslandPlayerDataDojo instantiates a new SkyBlockProfileMemberNetherIslandPlayerDataDojo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkyBlockProfileMemberNetherIslandPlayerDataDojo() *SkyBlockProfileMemberNetherIslandPlayerDataDojo {
	this := SkyBlockProfileMemberNetherIslandPlayerDataDojo{}
	return &this
}

// NewSkyBlockProfileMemberNetherIslandPlayerDataDojoWithDefaults instantiates a new SkyBlockProfileMemberNetherIslandPlayerDataDojo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkyBlockProfileMemberNetherIslandPlayerDataDojoWithDefaults() *SkyBlockProfileMemberNetherIslandPlayerDataDojo {
	this := SkyBlockProfileMemberNetherIslandPlayerDataDojo{}
	return &this
}

// GetDojoPointsArcher returns the DojoPointsArcher field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) GetDojoPointsArcher() int64 {
	if o == nil || IsNil(o.DojoPointsArcher) {
		var ret int64
		return ret
	}
	return *o.DojoPointsArcher
}

// GetDojoPointsArcherOk returns a tuple with the DojoPointsArcher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) GetDojoPointsArcherOk() (*int64, bool) {
	if o == nil || IsNil(o.DojoPointsArcher) {
		return nil, false
	}
	return o.DojoPointsArcher, true
}

// HasDojoPointsArcher returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) HasDojoPointsArcher() bool {
	if o != nil && !IsNil(o.DojoPointsArcher) {
		return true
	}

	return false
}

// SetDojoPointsArcher gets a reference to the given int64 and assigns it to the DojoPointsArcher field.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) SetDojoPointsArcher(v int64) {
	o.DojoPointsArcher = &v
}

// GetDojoPointsFireball returns the DojoPointsFireball field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) GetDojoPointsFireball() int64 {
	if o == nil || IsNil(o.DojoPointsFireball) {
		var ret int64
		return ret
	}
	return *o.DojoPointsFireball
}

// GetDojoPointsFireballOk returns a tuple with the DojoPointsFireball field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) GetDojoPointsFireballOk() (*int64, bool) {
	if o == nil || IsNil(o.DojoPointsFireball) {
		return nil, false
	}
	return o.DojoPointsFireball, true
}

// HasDojoPointsFireball returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) HasDojoPointsFireball() bool {
	if o != nil && !IsNil(o.DojoPointsFireball) {
		return true
	}

	return false
}

// SetDojoPointsFireball gets a reference to the given int64 and assigns it to the DojoPointsFireball field.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) SetDojoPointsFireball(v int64) {
	o.DojoPointsFireball = &v
}

// GetDojoPointsLockHead returns the DojoPointsLockHead field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) GetDojoPointsLockHead() int64 {
	if o == nil || IsNil(o.DojoPointsLockHead) {
		var ret int64
		return ret
	}
	return *o.DojoPointsLockHead
}

// GetDojoPointsLockHeadOk returns a tuple with the DojoPointsLockHead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) GetDojoPointsLockHeadOk() (*int64, bool) {
	if o == nil || IsNil(o.DojoPointsLockHead) {
		return nil, false
	}
	return o.DojoPointsLockHead, true
}

// HasDojoPointsLockHead returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) HasDojoPointsLockHead() bool {
	if o != nil && !IsNil(o.DojoPointsLockHead) {
		return true
	}

	return false
}

// SetDojoPointsLockHead gets a reference to the given int64 and assigns it to the DojoPointsLockHead field.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) SetDojoPointsLockHead(v int64) {
	o.DojoPointsLockHead = &v
}

// GetDojoPointsMobKb returns the DojoPointsMobKb field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) GetDojoPointsMobKb() int64 {
	if o == nil || IsNil(o.DojoPointsMobKb) {
		var ret int64
		return ret
	}
	return *o.DojoPointsMobKb
}

// GetDojoPointsMobKbOk returns a tuple with the DojoPointsMobKb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) GetDojoPointsMobKbOk() (*int64, bool) {
	if o == nil || IsNil(o.DojoPointsMobKb) {
		return nil, false
	}
	return o.DojoPointsMobKb, true
}

// HasDojoPointsMobKb returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) HasDojoPointsMobKb() bool {
	if o != nil && !IsNil(o.DojoPointsMobKb) {
		return true
	}

	return false
}

// SetDojoPointsMobKb gets a reference to the given int64 and assigns it to the DojoPointsMobKb field.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) SetDojoPointsMobKb(v int64) {
	o.DojoPointsMobKb = &v
}

// GetDojoPointsSnake returns the DojoPointsSnake field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) GetDojoPointsSnake() int64 {
	if o == nil || IsNil(o.DojoPointsSnake) {
		var ret int64
		return ret
	}
	return *o.DojoPointsSnake
}

// GetDojoPointsSnakeOk returns a tuple with the DojoPointsSnake field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) GetDojoPointsSnakeOk() (*int64, bool) {
	if o == nil || IsNil(o.DojoPointsSnake) {
		return nil, false
	}
	return o.DojoPointsSnake, true
}

// HasDojoPointsSnake returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) HasDojoPointsSnake() bool {
	if o != nil && !IsNil(o.DojoPointsSnake) {
		return true
	}

	return false
}

// SetDojoPointsSnake gets a reference to the given int64 and assigns it to the DojoPointsSnake field.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) SetDojoPointsSnake(v int64) {
	o.DojoPointsSnake = &v
}

// GetDojoPointsSwordSwap returns the DojoPointsSwordSwap field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) GetDojoPointsSwordSwap() int64 {
	if o == nil || IsNil(o.DojoPointsSwordSwap) {
		var ret int64
		return ret
	}
	return *o.DojoPointsSwordSwap
}

// GetDojoPointsSwordSwapOk returns a tuple with the DojoPointsSwordSwap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) GetDojoPointsSwordSwapOk() (*int64, bool) {
	if o == nil || IsNil(o.DojoPointsSwordSwap) {
		return nil, false
	}
	return o.DojoPointsSwordSwap, true
}

// HasDojoPointsSwordSwap returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) HasDojoPointsSwordSwap() bool {
	if o != nil && !IsNil(o.DojoPointsSwordSwap) {
		return true
	}

	return false
}

// SetDojoPointsSwordSwap gets a reference to the given int64 and assigns it to the DojoPointsSwordSwap field.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) SetDojoPointsSwordSwap(v int64) {
	o.DojoPointsSwordSwap = &v
}

// GetDojoPointsWallJump returns the DojoPointsWallJump field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) GetDojoPointsWallJump() int64 {
	if o == nil || IsNil(o.DojoPointsWallJump) {
		var ret int64
		return ret
	}
	return *o.DojoPointsWallJump
}

// GetDojoPointsWallJumpOk returns a tuple with the DojoPointsWallJump field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) GetDojoPointsWallJumpOk() (*int64, bool) {
	if o == nil || IsNil(o.DojoPointsWallJump) {
		return nil, false
	}
	return o.DojoPointsWallJump, true
}

// HasDojoPointsWallJump returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) HasDojoPointsWallJump() bool {
	if o != nil && !IsNil(o.DojoPointsWallJump) {
		return true
	}

	return false
}

// SetDojoPointsWallJump gets a reference to the given int64 and assigns it to the DojoPointsWallJump field.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) SetDojoPointsWallJump(v int64) {
	o.DojoPointsWallJump = &v
}

// GetDojoTimeArcher returns the DojoTimeArcher field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) GetDojoTimeArcher() int64 {
	if o == nil || IsNil(o.DojoTimeArcher) {
		var ret int64
		return ret
	}
	return *o.DojoTimeArcher
}

// GetDojoTimeArcherOk returns a tuple with the DojoTimeArcher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) GetDojoTimeArcherOk() (*int64, bool) {
	if o == nil || IsNil(o.DojoTimeArcher) {
		return nil, false
	}
	return o.DojoTimeArcher, true
}

// HasDojoTimeArcher returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) HasDojoTimeArcher() bool {
	if o != nil && !IsNil(o.DojoTimeArcher) {
		return true
	}

	return false
}

// SetDojoTimeArcher gets a reference to the given int64 and assigns it to the DojoTimeArcher field.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) SetDojoTimeArcher(v int64) {
	o.DojoTimeArcher = &v
}

// GetDojoTimeFireball returns the DojoTimeFireball field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) GetDojoTimeFireball() int64 {
	if o == nil || IsNil(o.DojoTimeFireball) {
		var ret int64
		return ret
	}
	return *o.DojoTimeFireball
}

// GetDojoTimeFireballOk returns a tuple with the DojoTimeFireball field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) GetDojoTimeFireballOk() (*int64, bool) {
	if o == nil || IsNil(o.DojoTimeFireball) {
		return nil, false
	}
	return o.DojoTimeFireball, true
}

// HasDojoTimeFireball returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) HasDojoTimeFireball() bool {
	if o != nil && !IsNil(o.DojoTimeFireball) {
		return true
	}

	return false
}

// SetDojoTimeFireball gets a reference to the given int64 and assigns it to the DojoTimeFireball field.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) SetDojoTimeFireball(v int64) {
	o.DojoTimeFireball = &v
}

// GetDojoTimeLockHead returns the DojoTimeLockHead field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) GetDojoTimeLockHead() int64 {
	if o == nil || IsNil(o.DojoTimeLockHead) {
		var ret int64
		return ret
	}
	return *o.DojoTimeLockHead
}

// GetDojoTimeLockHeadOk returns a tuple with the DojoTimeLockHead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) GetDojoTimeLockHeadOk() (*int64, bool) {
	if o == nil || IsNil(o.DojoTimeLockHead) {
		return nil, false
	}
	return o.DojoTimeLockHead, true
}

// HasDojoTimeLockHead returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) HasDojoTimeLockHead() bool {
	if o != nil && !IsNil(o.DojoTimeLockHead) {
		return true
	}

	return false
}

// SetDojoTimeLockHead gets a reference to the given int64 and assigns it to the DojoTimeLockHead field.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) SetDojoTimeLockHead(v int64) {
	o.DojoTimeLockHead = &v
}

// GetDojoTimeMobKb returns the DojoTimeMobKb field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) GetDojoTimeMobKb() int64 {
	if o == nil || IsNil(o.DojoTimeMobKb) {
		var ret int64
		return ret
	}
	return *o.DojoTimeMobKb
}

// GetDojoTimeMobKbOk returns a tuple with the DojoTimeMobKb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) GetDojoTimeMobKbOk() (*int64, bool) {
	if o == nil || IsNil(o.DojoTimeMobKb) {
		return nil, false
	}
	return o.DojoTimeMobKb, true
}

// HasDojoTimeMobKb returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) HasDojoTimeMobKb() bool {
	if o != nil && !IsNil(o.DojoTimeMobKb) {
		return true
	}

	return false
}

// SetDojoTimeMobKb gets a reference to the given int64 and assigns it to the DojoTimeMobKb field.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) SetDojoTimeMobKb(v int64) {
	o.DojoTimeMobKb = &v
}

// GetDojoTimeSnake returns the DojoTimeSnake field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) GetDojoTimeSnake() int64 {
	if o == nil || IsNil(o.DojoTimeSnake) {
		var ret int64
		return ret
	}
	return *o.DojoTimeSnake
}

// GetDojoTimeSnakeOk returns a tuple with the DojoTimeSnake field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) GetDojoTimeSnakeOk() (*int64, bool) {
	if o == nil || IsNil(o.DojoTimeSnake) {
		return nil, false
	}
	return o.DojoTimeSnake, true
}

// HasDojoTimeSnake returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) HasDojoTimeSnake() bool {
	if o != nil && !IsNil(o.DojoTimeSnake) {
		return true
	}

	return false
}

// SetDojoTimeSnake gets a reference to the given int64 and assigns it to the DojoTimeSnake field.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) SetDojoTimeSnake(v int64) {
	o.DojoTimeSnake = &v
}

// GetDojoTimeSwordSwap returns the DojoTimeSwordSwap field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) GetDojoTimeSwordSwap() int64 {
	if o == nil || IsNil(o.DojoTimeSwordSwap) {
		var ret int64
		return ret
	}
	return *o.DojoTimeSwordSwap
}

// GetDojoTimeSwordSwapOk returns a tuple with the DojoTimeSwordSwap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) GetDojoTimeSwordSwapOk() (*int64, bool) {
	if o == nil || IsNil(o.DojoTimeSwordSwap) {
		return nil, false
	}
	return o.DojoTimeSwordSwap, true
}

// HasDojoTimeSwordSwap returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) HasDojoTimeSwordSwap() bool {
	if o != nil && !IsNil(o.DojoTimeSwordSwap) {
		return true
	}

	return false
}

// SetDojoTimeSwordSwap gets a reference to the given int64 and assigns it to the DojoTimeSwordSwap field.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) SetDojoTimeSwordSwap(v int64) {
	o.DojoTimeSwordSwap = &v
}

// GetDojoTimeWallJump returns the DojoTimeWallJump field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) GetDojoTimeWallJump() int64 {
	if o == nil || IsNil(o.DojoTimeWallJump) {
		var ret int64
		return ret
	}
	return *o.DojoTimeWallJump
}

// GetDojoTimeWallJumpOk returns a tuple with the DojoTimeWallJump field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) GetDojoTimeWallJumpOk() (*int64, bool) {
	if o == nil || IsNil(o.DojoTimeWallJump) {
		return nil, false
	}
	return o.DojoTimeWallJump, true
}

// HasDojoTimeWallJump returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) HasDojoTimeWallJump() bool {
	if o != nil && !IsNil(o.DojoTimeWallJump) {
		return true
	}

	return false
}

// SetDojoTimeWallJump gets a reference to the given int64 and assigns it to the DojoTimeWallJump field.
func (o *SkyBlockProfileMemberNetherIslandPlayerDataDojo) SetDojoTimeWallJump(v int64) {
	o.DojoTimeWallJump = &v
}

func (o SkyBlockProfileMemberNetherIslandPlayerDataDojo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkyBlockProfileMemberNetherIslandPlayerDataDojo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DojoPointsArcher) {
		toSerialize["dojo_points_archer"] = o.DojoPointsArcher
	}
	if !IsNil(o.DojoPointsFireball) {
		toSerialize["dojo_points_fireball"] = o.DojoPointsFireball
	}
	if !IsNil(o.DojoPointsLockHead) {
		toSerialize["dojo_points_lock_head"] = o.DojoPointsLockHead
	}
	if !IsNil(o.DojoPointsMobKb) {
		toSerialize["dojo_points_mob_kb"] = o.DojoPointsMobKb
	}
	if !IsNil(o.DojoPointsSnake) {
		toSerialize["dojo_points_snake"] = o.DojoPointsSnake
	}
	if !IsNil(o.DojoPointsSwordSwap) {
		toSerialize["dojo_points_sword_swap"] = o.DojoPointsSwordSwap
	}
	if !IsNil(o.DojoPointsWallJump) {
		toSerialize["dojo_points_wall_jump"] = o.DojoPointsWallJump
	}
	if !IsNil(o.DojoTimeArcher) {
		toSerialize["dojo_time_archer"] = o.DojoTimeArcher
	}
	if !IsNil(o.DojoTimeFireball) {
		toSerialize["dojo_time_fireball"] = o.DojoTimeFireball
	}
	if !IsNil(o.DojoTimeLockHead) {
		toSerialize["dojo_time_lock_head"] = o.DojoTimeLockHead
	}
	if !IsNil(o.DojoTimeMobKb) {
		toSerialize["dojo_time_mob_kb"] = o.DojoTimeMobKb
	}
	if !IsNil(o.DojoTimeSnake) {
		toSerialize["dojo_time_snake"] = o.DojoTimeSnake
	}
	if !IsNil(o.DojoTimeSwordSwap) {
		toSerialize["dojo_time_sword_swap"] = o.DojoTimeSwordSwap
	}
	if !IsNil(o.DojoTimeWallJump) {
		toSerialize["dojo_time_wall_jump"] = o.DojoTimeWallJump
	}
	return toSerialize, nil
}

type NullableSkyBlockProfileMemberNetherIslandPlayerDataDojo struct {
	value *SkyBlockProfileMemberNetherIslandPlayerDataDojo
	isSet bool
}

func (v NullableSkyBlockProfileMemberNetherIslandPlayerDataDojo) Get() *SkyBlockProfileMemberNetherIslandPlayerDataDojo {
	return v.value
}

func (v *NullableSkyBlockProfileMemberNetherIslandPlayerDataDojo) Set(val *SkyBlockProfileMemberNetherIslandPlayerDataDojo) {
	v.value = val
	v.isSet = true
}

func (v NullableSkyBlockProfileMemberNetherIslandPlayerDataDojo) IsSet() bool {
	return v.isSet
}

func (v *NullableSkyBlockProfileMemberNetherIslandPlayerDataDojo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkyBlockProfileMemberNetherIslandPlayerDataDojo(val *SkyBlockProfileMemberNetherIslandPlayerDataDojo) *NullableSkyBlockProfileMemberNetherIslandPlayerDataDojo {
	return &NullableSkyBlockProfileMemberNetherIslandPlayerDataDojo{value: val, isSet: true}
}

func (v NullableSkyBlockProfileMemberNetherIslandPlayerDataDojo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkyBlockProfileMemberNetherIslandPlayerDataDojo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
