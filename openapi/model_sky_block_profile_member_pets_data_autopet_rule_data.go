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

// checks if the SkyBlockProfileMemberPetsDataAutopetRuleData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkyBlockProfileMemberPetsDataAutopetRuleData{}

// SkyBlockProfileMemberPetsDataAutopetRuleData struct for SkyBlockProfileMemberPetsDataAutopetRuleData
type SkyBlockProfileMemberPetsDataAutopetRuleData struct {
	Boss       *string `json:"boss,omitempty"`
	Category   *string `json:"category,omitempty"`
	Class      *string `json:"class,omitempty"`
	Collection *string `json:"collection,omitempty"`
	EntityType *string `json:"entity_type,omitempty"`
	Floor      *string `json:"floor,omitempty"`
	Island     *string `json:"island,omitempty"`
	Master     *string `json:"master,omitempty"`
	Skill      *string `json:"skill,omitempty"`
	Slot       *string `json:"slot,omitempty"`
}

// NewSkyBlockProfileMemberPetsDataAutopetRuleData instantiates a new SkyBlockProfileMemberPetsDataAutopetRuleData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkyBlockProfileMemberPetsDataAutopetRuleData() *SkyBlockProfileMemberPetsDataAutopetRuleData {
	this := SkyBlockProfileMemberPetsDataAutopetRuleData{}
	return &this
}

// NewSkyBlockProfileMemberPetsDataAutopetRuleDataWithDefaults instantiates a new SkyBlockProfileMemberPetsDataAutopetRuleData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkyBlockProfileMemberPetsDataAutopetRuleDataWithDefaults() *SkyBlockProfileMemberPetsDataAutopetRuleData {
	this := SkyBlockProfileMemberPetsDataAutopetRuleData{}
	return &this
}

// GetBoss returns the Boss field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) GetBoss() string {
	if o == nil || IsNil(o.Boss) {
		var ret string
		return ret
	}
	return *o.Boss
}

// GetBossOk returns a tuple with the Boss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) GetBossOk() (*string, bool) {
	if o == nil || IsNil(o.Boss) {
		return nil, false
	}
	return o.Boss, true
}

// HasBoss returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) HasBoss() bool {
	if o != nil && !IsNil(o.Boss) {
		return true
	}

	return false
}

// SetBoss gets a reference to the given string and assigns it to the Boss field.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) SetBoss(v string) {
	o.Boss = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) SetCategory(v string) {
	o.Category = &v
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) GetClass() string {
	if o == nil || IsNil(o.Class) {
		var ret string
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) GetClassOk() (*string, bool) {
	if o == nil || IsNil(o.Class) {
		return nil, false
	}
	return o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) HasClass() bool {
	if o != nil && !IsNil(o.Class) {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) SetClass(v string) {
	o.Class = &v
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) GetCollection() string {
	if o == nil || IsNil(o.Collection) {
		var ret string
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) GetCollectionOk() (*string, bool) {
	if o == nil || IsNil(o.Collection) {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) HasCollection() bool {
	if o != nil && !IsNil(o.Collection) {
		return true
	}

	return false
}

// SetCollection gets a reference to the given string and assigns it to the Collection field.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) SetCollection(v string) {
	o.Collection = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) GetEntityType() string {
	if o == nil || IsNil(o.EntityType) {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) GetEntityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EntityType) {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) HasEntityType() bool {
	if o != nil && !IsNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) SetEntityType(v string) {
	o.EntityType = &v
}

// GetFloor returns the Floor field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) GetFloor() string {
	if o == nil || IsNil(o.Floor) {
		var ret string
		return ret
	}
	return *o.Floor
}

// GetFloorOk returns a tuple with the Floor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) GetFloorOk() (*string, bool) {
	if o == nil || IsNil(o.Floor) {
		return nil, false
	}
	return o.Floor, true
}

// HasFloor returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) HasFloor() bool {
	if o != nil && !IsNil(o.Floor) {
		return true
	}

	return false
}

// SetFloor gets a reference to the given string and assigns it to the Floor field.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) SetFloor(v string) {
	o.Floor = &v
}

// GetIsland returns the Island field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) GetIsland() string {
	if o == nil || IsNil(o.Island) {
		var ret string
		return ret
	}
	return *o.Island
}

// GetIslandOk returns a tuple with the Island field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) GetIslandOk() (*string, bool) {
	if o == nil || IsNil(o.Island) {
		return nil, false
	}
	return o.Island, true
}

// HasIsland returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) HasIsland() bool {
	if o != nil && !IsNil(o.Island) {
		return true
	}

	return false
}

// SetIsland gets a reference to the given string and assigns it to the Island field.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) SetIsland(v string) {
	o.Island = &v
}

// GetMaster returns the Master field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) GetMaster() string {
	if o == nil || IsNil(o.Master) {
		var ret string
		return ret
	}
	return *o.Master
}

// GetMasterOk returns a tuple with the Master field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) GetMasterOk() (*string, bool) {
	if o == nil || IsNil(o.Master) {
		return nil, false
	}
	return o.Master, true
}

// HasMaster returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) HasMaster() bool {
	if o != nil && !IsNil(o.Master) {
		return true
	}

	return false
}

// SetMaster gets a reference to the given string and assigns it to the Master field.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) SetMaster(v string) {
	o.Master = &v
}

// GetSkill returns the Skill field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) GetSkill() string {
	if o == nil || IsNil(o.Skill) {
		var ret string
		return ret
	}
	return *o.Skill
}

// GetSkillOk returns a tuple with the Skill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) GetSkillOk() (*string, bool) {
	if o == nil || IsNil(o.Skill) {
		return nil, false
	}
	return o.Skill, true
}

// HasSkill returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) HasSkill() bool {
	if o != nil && !IsNil(o.Skill) {
		return true
	}

	return false
}

// SetSkill gets a reference to the given string and assigns it to the Skill field.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) SetSkill(v string) {
	o.Skill = &v
}

// GetSlot returns the Slot field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) GetSlot() string {
	if o == nil || IsNil(o.Slot) {
		var ret string
		return ret
	}
	return *o.Slot
}

// GetSlotOk returns a tuple with the Slot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) GetSlotOk() (*string, bool) {
	if o == nil || IsNil(o.Slot) {
		return nil, false
	}
	return o.Slot, true
}

// HasSlot returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) HasSlot() bool {
	if o != nil && !IsNil(o.Slot) {
		return true
	}

	return false
}

// SetSlot gets a reference to the given string and assigns it to the Slot field.
func (o *SkyBlockProfileMemberPetsDataAutopetRuleData) SetSlot(v string) {
	o.Slot = &v
}

func (o SkyBlockProfileMemberPetsDataAutopetRuleData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkyBlockProfileMemberPetsDataAutopetRuleData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Boss) {
		toSerialize["boss"] = o.Boss
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Class) {
		toSerialize["class"] = o.Class
	}
	if !IsNil(o.Collection) {
		toSerialize["collection"] = o.Collection
	}
	if !IsNil(o.EntityType) {
		toSerialize["entity_type"] = o.EntityType
	}
	if !IsNil(o.Floor) {
		toSerialize["floor"] = o.Floor
	}
	if !IsNil(o.Island) {
		toSerialize["island"] = o.Island
	}
	if !IsNil(o.Master) {
		toSerialize["master"] = o.Master
	}
	if !IsNil(o.Skill) {
		toSerialize["skill"] = o.Skill
	}
	if !IsNil(o.Slot) {
		toSerialize["slot"] = o.Slot
	}
	return toSerialize, nil
}

type NullableSkyBlockProfileMemberPetsDataAutopetRuleData struct {
	value *SkyBlockProfileMemberPetsDataAutopetRuleData
	isSet bool
}

func (v NullableSkyBlockProfileMemberPetsDataAutopetRuleData) Get() *SkyBlockProfileMemberPetsDataAutopetRuleData {
	return v.value
}

func (v *NullableSkyBlockProfileMemberPetsDataAutopetRuleData) Set(val *SkyBlockProfileMemberPetsDataAutopetRuleData) {
	v.value = val
	v.isSet = true
}

func (v NullableSkyBlockProfileMemberPetsDataAutopetRuleData) IsSet() bool {
	return v.isSet
}

func (v *NullableSkyBlockProfileMemberPetsDataAutopetRuleData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkyBlockProfileMemberPetsDataAutopetRuleData(val *SkyBlockProfileMemberPetsDataAutopetRuleData) *NullableSkyBlockProfileMemberPetsDataAutopetRuleData {
	return &NullableSkyBlockProfileMemberPetsDataAutopetRuleData{value: val, isSet: true}
}

func (v NullableSkyBlockProfileMemberPetsDataAutopetRuleData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkyBlockProfileMemberPetsDataAutopetRuleData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
