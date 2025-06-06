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

// checks if the SkyBlockProfileMemberProfileCoopInvitation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkyBlockProfileMemberProfileCoopInvitation{}

// SkyBlockProfileMemberProfileCoopInvitation struct for SkyBlockProfileMemberProfileCoopInvitation
type SkyBlockProfileMemberProfileCoopInvitation struct {
	Confirmed          bool   `json:"confirmed"`
	ConfirmedTimestamp *int64 `json:"confirmed_timestamp,omitempty"`
	InvitedBy          string `json:"invited_by"`
	Timestamp          int64  `json:"timestamp"`
}

type _SkyBlockProfileMemberProfileCoopInvitation SkyBlockProfileMemberProfileCoopInvitation

// NewSkyBlockProfileMemberProfileCoopInvitation instantiates a new SkyBlockProfileMemberProfileCoopInvitation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkyBlockProfileMemberProfileCoopInvitation(confirmed bool, invitedBy string, timestamp int64) *SkyBlockProfileMemberProfileCoopInvitation {
	this := SkyBlockProfileMemberProfileCoopInvitation{}
	this.Confirmed = confirmed
	this.InvitedBy = invitedBy
	this.Timestamp = timestamp
	return &this
}

// NewSkyBlockProfileMemberProfileCoopInvitationWithDefaults instantiates a new SkyBlockProfileMemberProfileCoopInvitation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkyBlockProfileMemberProfileCoopInvitationWithDefaults() *SkyBlockProfileMemberProfileCoopInvitation {
	this := SkyBlockProfileMemberProfileCoopInvitation{}
	return &this
}

// GetConfirmed returns the Confirmed field value
func (o *SkyBlockProfileMemberProfileCoopInvitation) GetConfirmed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Confirmed
}

// GetConfirmedOk returns a tuple with the Confirmed field value
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberProfileCoopInvitation) GetConfirmedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Confirmed, true
}

// SetConfirmed sets field value
func (o *SkyBlockProfileMemberProfileCoopInvitation) SetConfirmed(v bool) {
	o.Confirmed = v
}

// GetConfirmedTimestamp returns the ConfirmedTimestamp field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberProfileCoopInvitation) GetConfirmedTimestamp() int64 {
	if o == nil || IsNil(o.ConfirmedTimestamp) {
		var ret int64
		return ret
	}
	return *o.ConfirmedTimestamp
}

// GetConfirmedTimestampOk returns a tuple with the ConfirmedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberProfileCoopInvitation) GetConfirmedTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.ConfirmedTimestamp) {
		return nil, false
	}
	return o.ConfirmedTimestamp, true
}

// HasConfirmedTimestamp returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberProfileCoopInvitation) HasConfirmedTimestamp() bool {
	if o != nil && !IsNil(o.ConfirmedTimestamp) {
		return true
	}

	return false
}

// SetConfirmedTimestamp gets a reference to the given int64 and assigns it to the ConfirmedTimestamp field.
func (o *SkyBlockProfileMemberProfileCoopInvitation) SetConfirmedTimestamp(v int64) {
	o.ConfirmedTimestamp = &v
}

// GetInvitedBy returns the InvitedBy field value
func (o *SkyBlockProfileMemberProfileCoopInvitation) GetInvitedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvitedBy
}

// GetInvitedByOk returns a tuple with the InvitedBy field value
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberProfileCoopInvitation) GetInvitedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvitedBy, true
}

// SetInvitedBy sets field value
func (o *SkyBlockProfileMemberProfileCoopInvitation) SetInvitedBy(v string) {
	o.InvitedBy = v
}

// GetTimestamp returns the Timestamp field value
func (o *SkyBlockProfileMemberProfileCoopInvitation) GetTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberProfileCoopInvitation) GetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *SkyBlockProfileMemberProfileCoopInvitation) SetTimestamp(v int64) {
	o.Timestamp = v
}

func (o SkyBlockProfileMemberProfileCoopInvitation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkyBlockProfileMemberProfileCoopInvitation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["confirmed"] = o.Confirmed
	if !IsNil(o.ConfirmedTimestamp) {
		toSerialize["confirmed_timestamp"] = o.ConfirmedTimestamp
	}
	toSerialize["invited_by"] = o.InvitedBy
	toSerialize["timestamp"] = o.Timestamp
	return toSerialize, nil
}

func (o *SkyBlockProfileMemberProfileCoopInvitation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"confirmed",
		"invited_by",
		"timestamp",
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

	varSkyBlockProfileMemberProfileCoopInvitation := _SkyBlockProfileMemberProfileCoopInvitation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSkyBlockProfileMemberProfileCoopInvitation)

	if err != nil {
		return err
	}

	*o = SkyBlockProfileMemberProfileCoopInvitation(varSkyBlockProfileMemberProfileCoopInvitation)

	return err
}

type NullableSkyBlockProfileMemberProfileCoopInvitation struct {
	value *SkyBlockProfileMemberProfileCoopInvitation
	isSet bool
}

func (v NullableSkyBlockProfileMemberProfileCoopInvitation) Get() *SkyBlockProfileMemberProfileCoopInvitation {
	return v.value
}

func (v *NullableSkyBlockProfileMemberProfileCoopInvitation) Set(val *SkyBlockProfileMemberProfileCoopInvitation) {
	v.value = val
	v.isSet = true
}

func (v NullableSkyBlockProfileMemberProfileCoopInvitation) IsSet() bool {
	return v.isSet
}

func (v *NullableSkyBlockProfileMemberProfileCoopInvitation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkyBlockProfileMemberProfileCoopInvitation(val *SkyBlockProfileMemberProfileCoopInvitation) *NullableSkyBlockProfileMemberProfileCoopInvitation {
	return &NullableSkyBlockProfileMemberProfileCoopInvitation{value: val, isSet: true}
}

func (v NullableSkyBlockProfileMemberProfileCoopInvitation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkyBlockProfileMemberProfileCoopInvitation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
