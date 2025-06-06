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

// checks if the SkyBlockProfileMemberRiftWestVillageMirrorverse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkyBlockProfileMemberRiftWestVillageMirrorverse{}

// SkyBlockProfileMemberRiftWestVillageMirrorverse struct for SkyBlockProfileMemberRiftWestVillageMirrorverse
type SkyBlockProfileMemberRiftWestVillageMirrorverse struct {
	ClaimedChestItems []SkyBlockProfileMemberRiftWestVillageMirrorverseClaimedChestItemsInner `json:"claimed_chest_items,omitempty"`
	ClaimedReward     *bool                                                                   `json:"claimed_reward,omitempty"`
	UpsideDownHard    *bool                                                                   `json:"upside_down_hard,omitempty"`
	VisitedRooms      []SkyBlockProfileMemberRiftWestVillageMirrorverseVisitedRoomsInner      `json:"visited_rooms,omitempty"`
}

// NewSkyBlockProfileMemberRiftWestVillageMirrorverse instantiates a new SkyBlockProfileMemberRiftWestVillageMirrorverse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkyBlockProfileMemberRiftWestVillageMirrorverse() *SkyBlockProfileMemberRiftWestVillageMirrorverse {
	this := SkyBlockProfileMemberRiftWestVillageMirrorverse{}
	return &this
}

// NewSkyBlockProfileMemberRiftWestVillageMirrorverseWithDefaults instantiates a new SkyBlockProfileMemberRiftWestVillageMirrorverse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkyBlockProfileMemberRiftWestVillageMirrorverseWithDefaults() *SkyBlockProfileMemberRiftWestVillageMirrorverse {
	this := SkyBlockProfileMemberRiftWestVillageMirrorverse{}
	return &this
}

// GetClaimedChestItems returns the ClaimedChestItems field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberRiftWestVillageMirrorverse) GetClaimedChestItems() []SkyBlockProfileMemberRiftWestVillageMirrorverseClaimedChestItemsInner {
	if o == nil || IsNil(o.ClaimedChestItems) {
		var ret []SkyBlockProfileMemberRiftWestVillageMirrorverseClaimedChestItemsInner
		return ret
	}
	return o.ClaimedChestItems
}

// GetClaimedChestItemsOk returns a tuple with the ClaimedChestItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberRiftWestVillageMirrorverse) GetClaimedChestItemsOk() ([]SkyBlockProfileMemberRiftWestVillageMirrorverseClaimedChestItemsInner, bool) {
	if o == nil || IsNil(o.ClaimedChestItems) {
		return nil, false
	}
	return o.ClaimedChestItems, true
}

// HasClaimedChestItems returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberRiftWestVillageMirrorverse) HasClaimedChestItems() bool {
	if o != nil && !IsNil(o.ClaimedChestItems) {
		return true
	}

	return false
}

// SetClaimedChestItems gets a reference to the given []SkyBlockProfileMemberRiftWestVillageMirrorverseClaimedChestItemsInner and assigns it to the ClaimedChestItems field.
func (o *SkyBlockProfileMemberRiftWestVillageMirrorverse) SetClaimedChestItems(v []SkyBlockProfileMemberRiftWestVillageMirrorverseClaimedChestItemsInner) {
	o.ClaimedChestItems = v
}

// GetClaimedReward returns the ClaimedReward field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberRiftWestVillageMirrorverse) GetClaimedReward() bool {
	if o == nil || IsNil(o.ClaimedReward) {
		var ret bool
		return ret
	}
	return *o.ClaimedReward
}

// GetClaimedRewardOk returns a tuple with the ClaimedReward field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberRiftWestVillageMirrorverse) GetClaimedRewardOk() (*bool, bool) {
	if o == nil || IsNil(o.ClaimedReward) {
		return nil, false
	}
	return o.ClaimedReward, true
}

// HasClaimedReward returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberRiftWestVillageMirrorverse) HasClaimedReward() bool {
	if o != nil && !IsNil(o.ClaimedReward) {
		return true
	}

	return false
}

// SetClaimedReward gets a reference to the given bool and assigns it to the ClaimedReward field.
func (o *SkyBlockProfileMemberRiftWestVillageMirrorverse) SetClaimedReward(v bool) {
	o.ClaimedReward = &v
}

// GetUpsideDownHard returns the UpsideDownHard field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberRiftWestVillageMirrorverse) GetUpsideDownHard() bool {
	if o == nil || IsNil(o.UpsideDownHard) {
		var ret bool
		return ret
	}
	return *o.UpsideDownHard
}

// GetUpsideDownHardOk returns a tuple with the UpsideDownHard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberRiftWestVillageMirrorverse) GetUpsideDownHardOk() (*bool, bool) {
	if o == nil || IsNil(o.UpsideDownHard) {
		return nil, false
	}
	return o.UpsideDownHard, true
}

// HasUpsideDownHard returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberRiftWestVillageMirrorverse) HasUpsideDownHard() bool {
	if o != nil && !IsNil(o.UpsideDownHard) {
		return true
	}

	return false
}

// SetUpsideDownHard gets a reference to the given bool and assigns it to the UpsideDownHard field.
func (o *SkyBlockProfileMemberRiftWestVillageMirrorverse) SetUpsideDownHard(v bool) {
	o.UpsideDownHard = &v
}

// GetVisitedRooms returns the VisitedRooms field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberRiftWestVillageMirrorverse) GetVisitedRooms() []SkyBlockProfileMemberRiftWestVillageMirrorverseVisitedRoomsInner {
	if o == nil || IsNil(o.VisitedRooms) {
		var ret []SkyBlockProfileMemberRiftWestVillageMirrorverseVisitedRoomsInner
		return ret
	}
	return o.VisitedRooms
}

// GetVisitedRoomsOk returns a tuple with the VisitedRooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberRiftWestVillageMirrorverse) GetVisitedRoomsOk() ([]SkyBlockProfileMemberRiftWestVillageMirrorverseVisitedRoomsInner, bool) {
	if o == nil || IsNil(o.VisitedRooms) {
		return nil, false
	}
	return o.VisitedRooms, true
}

// HasVisitedRooms returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberRiftWestVillageMirrorverse) HasVisitedRooms() bool {
	if o != nil && !IsNil(o.VisitedRooms) {
		return true
	}

	return false
}

// SetVisitedRooms gets a reference to the given []SkyBlockProfileMemberRiftWestVillageMirrorverseVisitedRoomsInner and assigns it to the VisitedRooms field.
func (o *SkyBlockProfileMemberRiftWestVillageMirrorverse) SetVisitedRooms(v []SkyBlockProfileMemberRiftWestVillageMirrorverseVisitedRoomsInner) {
	o.VisitedRooms = v
}

func (o SkyBlockProfileMemberRiftWestVillageMirrorverse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkyBlockProfileMemberRiftWestVillageMirrorverse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClaimedChestItems) {
		toSerialize["claimed_chest_items"] = o.ClaimedChestItems
	}
	if !IsNil(o.ClaimedReward) {
		toSerialize["claimed_reward"] = o.ClaimedReward
	}
	if !IsNil(o.UpsideDownHard) {
		toSerialize["upside_down_hard"] = o.UpsideDownHard
	}
	if !IsNil(o.VisitedRooms) {
		toSerialize["visited_rooms"] = o.VisitedRooms
	}
	return toSerialize, nil
}

type NullableSkyBlockProfileMemberRiftWestVillageMirrorverse struct {
	value *SkyBlockProfileMemberRiftWestVillageMirrorverse
	isSet bool
}

func (v NullableSkyBlockProfileMemberRiftWestVillageMirrorverse) Get() *SkyBlockProfileMemberRiftWestVillageMirrorverse {
	return v.value
}

func (v *NullableSkyBlockProfileMemberRiftWestVillageMirrorverse) Set(val *SkyBlockProfileMemberRiftWestVillageMirrorverse) {
	v.value = val
	v.isSet = true
}

func (v NullableSkyBlockProfileMemberRiftWestVillageMirrorverse) IsSet() bool {
	return v.isSet
}

func (v *NullableSkyBlockProfileMemberRiftWestVillageMirrorverse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkyBlockProfileMemberRiftWestVillageMirrorverse(val *SkyBlockProfileMemberRiftWestVillageMirrorverse) *NullableSkyBlockProfileMemberRiftWestVillageMirrorverse {
	return &NullableSkyBlockProfileMemberRiftWestVillageMirrorverse{value: val, isSet: true}
}

func (v NullableSkyBlockProfileMemberRiftWestVillageMirrorverse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkyBlockProfileMemberRiftWestVillageMirrorverse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
