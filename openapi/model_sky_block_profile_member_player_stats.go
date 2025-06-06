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

// checks if the SkyBlockProfileMemberPlayerStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkyBlockProfileMemberPlayerStats{}

// SkyBlockProfileMemberPlayerStats struct for SkyBlockProfileMemberPlayerStats
type SkyBlockProfileMemberPlayerStats struct {
	Auctions               *SkyBlockProfileMemberPlayerStatsAuctions       `json:"auctions,omitempty"`
	CandyCollected         *SkyBlockProfileMemberPlayerStatsCandyCollected `json:"candy_collected,omitempty"`
	CropsMined             *float64                                        `json:"crops_mined,omitempty"`
	Deaths                 *map[string]float64                             `json:"deaths,omitempty"`
	EndIsland              *SkyBlockProfileMemberPlayerStatsEndIsland      `json:"end_island,omitempty"`
	Gifts                  *SkyBlockProfileMemberPlayerStatsGifts          `json:"gifts,omitempty"`
	GlowingMushroomsBroken *float64                                        `json:"glowing_mushrooms_broken,omitempty"`
	HighestCriticalDamage  *float64                                        `json:"highest_critical_damage,omitempty"`
	HighestDamage          *float64                                        `json:"highest_damage,omitempty"`
	ItemsFished            *SkyBlockProfileMemberPlayerStatsItemsFished    `json:"items_fished,omitempty"`
	Kills                  *map[string]float64                             `json:"kills,omitempty"`
	Mythos                 *SkyBlockProfileMemberPlayerStatsMythos         `json:"mythos,omitempty"`
	Pets                   *SkyBlockProfileMemberPlayerStatsPets           `json:"pets,omitempty"`
	Races                  *SkyBlockProfileMemberPlayerStatsRaces          `json:"races,omitempty"`
	Rift                   *SkyBlockProfileMemberPlayerStatsRift           `json:"rift,omitempty"`
	SeaCreatureKills       *float64                                        `json:"sea_creature_kills,omitempty"`
	ShredderRod            *SkyBlockProfileMemberPlayerStatsShredderRod    `json:"shredder_rod,omitempty"`
	Spooky                 *SkyBlockProfileMemberPlayerStatsSpooky         `json:"spooky,omitempty"`
	Winter                 *SkyBlockProfileMemberPlayerStatsWinter         `json:"winter,omitempty"`
}

// NewSkyBlockProfileMemberPlayerStats instantiates a new SkyBlockProfileMemberPlayerStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkyBlockProfileMemberPlayerStats() *SkyBlockProfileMemberPlayerStats {
	this := SkyBlockProfileMemberPlayerStats{}
	return &this
}

// NewSkyBlockProfileMemberPlayerStatsWithDefaults instantiates a new SkyBlockProfileMemberPlayerStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkyBlockProfileMemberPlayerStatsWithDefaults() *SkyBlockProfileMemberPlayerStats {
	this := SkyBlockProfileMemberPlayerStats{}
	return &this
}

// GetAuctions returns the Auctions field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStats) GetAuctions() SkyBlockProfileMemberPlayerStatsAuctions {
	if o == nil || IsNil(o.Auctions) {
		var ret SkyBlockProfileMemberPlayerStatsAuctions
		return ret
	}
	return *o.Auctions
}

// GetAuctionsOk returns a tuple with the Auctions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStats) GetAuctionsOk() (*SkyBlockProfileMemberPlayerStatsAuctions, bool) {
	if o == nil || IsNil(o.Auctions) {
		return nil, false
	}
	return o.Auctions, true
}

// HasAuctions returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStats) HasAuctions() bool {
	if o != nil && !IsNil(o.Auctions) {
		return true
	}

	return false
}

// SetAuctions gets a reference to the given SkyBlockProfileMemberPlayerStatsAuctions and assigns it to the Auctions field.
func (o *SkyBlockProfileMemberPlayerStats) SetAuctions(v SkyBlockProfileMemberPlayerStatsAuctions) {
	o.Auctions = &v
}

// GetCandyCollected returns the CandyCollected field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStats) GetCandyCollected() SkyBlockProfileMemberPlayerStatsCandyCollected {
	if o == nil || IsNil(o.CandyCollected) {
		var ret SkyBlockProfileMemberPlayerStatsCandyCollected
		return ret
	}
	return *o.CandyCollected
}

// GetCandyCollectedOk returns a tuple with the CandyCollected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStats) GetCandyCollectedOk() (*SkyBlockProfileMemberPlayerStatsCandyCollected, bool) {
	if o == nil || IsNil(o.CandyCollected) {
		return nil, false
	}
	return o.CandyCollected, true
}

// HasCandyCollected returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStats) HasCandyCollected() bool {
	if o != nil && !IsNil(o.CandyCollected) {
		return true
	}

	return false
}

// SetCandyCollected gets a reference to the given SkyBlockProfileMemberPlayerStatsCandyCollected and assigns it to the CandyCollected field.
func (o *SkyBlockProfileMemberPlayerStats) SetCandyCollected(v SkyBlockProfileMemberPlayerStatsCandyCollected) {
	o.CandyCollected = &v
}

// GetCropsMined returns the CropsMined field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStats) GetCropsMined() float64 {
	if o == nil || IsNil(o.CropsMined) {
		var ret float64
		return ret
	}
	return *o.CropsMined
}

// GetCropsMinedOk returns a tuple with the CropsMined field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStats) GetCropsMinedOk() (*float64, bool) {
	if o == nil || IsNil(o.CropsMined) {
		return nil, false
	}
	return o.CropsMined, true
}

// HasCropsMined returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStats) HasCropsMined() bool {
	if o != nil && !IsNil(o.CropsMined) {
		return true
	}

	return false
}

// SetCropsMined gets a reference to the given float64 and assigns it to the CropsMined field.
func (o *SkyBlockProfileMemberPlayerStats) SetCropsMined(v float64) {
	o.CropsMined = &v
}

// GetDeaths returns the Deaths field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStats) GetDeaths() map[string]float64 {
	if o == nil || IsNil(o.Deaths) {
		var ret map[string]float64
		return ret
	}
	return *o.Deaths
}

// GetDeathsOk returns a tuple with the Deaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStats) GetDeathsOk() (*map[string]float64, bool) {
	if o == nil || IsNil(o.Deaths) {
		return nil, false
	}
	return o.Deaths, true
}

// HasDeaths returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStats) HasDeaths() bool {
	if o != nil && !IsNil(o.Deaths) {
		return true
	}

	return false
}

// SetDeaths gets a reference to the given map[string]float64 and assigns it to the Deaths field.
func (o *SkyBlockProfileMemberPlayerStats) SetDeaths(v map[string]float64) {
	o.Deaths = &v
}

// GetEndIsland returns the EndIsland field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStats) GetEndIsland() SkyBlockProfileMemberPlayerStatsEndIsland {
	if o == nil || IsNil(o.EndIsland) {
		var ret SkyBlockProfileMemberPlayerStatsEndIsland
		return ret
	}
	return *o.EndIsland
}

// GetEndIslandOk returns a tuple with the EndIsland field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStats) GetEndIslandOk() (*SkyBlockProfileMemberPlayerStatsEndIsland, bool) {
	if o == nil || IsNil(o.EndIsland) {
		return nil, false
	}
	return o.EndIsland, true
}

// HasEndIsland returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStats) HasEndIsland() bool {
	if o != nil && !IsNil(o.EndIsland) {
		return true
	}

	return false
}

// SetEndIsland gets a reference to the given SkyBlockProfileMemberPlayerStatsEndIsland and assigns it to the EndIsland field.
func (o *SkyBlockProfileMemberPlayerStats) SetEndIsland(v SkyBlockProfileMemberPlayerStatsEndIsland) {
	o.EndIsland = &v
}

// GetGifts returns the Gifts field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStats) GetGifts() SkyBlockProfileMemberPlayerStatsGifts {
	if o == nil || IsNil(o.Gifts) {
		var ret SkyBlockProfileMemberPlayerStatsGifts
		return ret
	}
	return *o.Gifts
}

// GetGiftsOk returns a tuple with the Gifts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStats) GetGiftsOk() (*SkyBlockProfileMemberPlayerStatsGifts, bool) {
	if o == nil || IsNil(o.Gifts) {
		return nil, false
	}
	return o.Gifts, true
}

// HasGifts returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStats) HasGifts() bool {
	if o != nil && !IsNil(o.Gifts) {
		return true
	}

	return false
}

// SetGifts gets a reference to the given SkyBlockProfileMemberPlayerStatsGifts and assigns it to the Gifts field.
func (o *SkyBlockProfileMemberPlayerStats) SetGifts(v SkyBlockProfileMemberPlayerStatsGifts) {
	o.Gifts = &v
}

// GetGlowingMushroomsBroken returns the GlowingMushroomsBroken field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStats) GetGlowingMushroomsBroken() float64 {
	if o == nil || IsNil(o.GlowingMushroomsBroken) {
		var ret float64
		return ret
	}
	return *o.GlowingMushroomsBroken
}

// GetGlowingMushroomsBrokenOk returns a tuple with the GlowingMushroomsBroken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStats) GetGlowingMushroomsBrokenOk() (*float64, bool) {
	if o == nil || IsNil(o.GlowingMushroomsBroken) {
		return nil, false
	}
	return o.GlowingMushroomsBroken, true
}

// HasGlowingMushroomsBroken returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStats) HasGlowingMushroomsBroken() bool {
	if o != nil && !IsNil(o.GlowingMushroomsBroken) {
		return true
	}

	return false
}

// SetGlowingMushroomsBroken gets a reference to the given float64 and assigns it to the GlowingMushroomsBroken field.
func (o *SkyBlockProfileMemberPlayerStats) SetGlowingMushroomsBroken(v float64) {
	o.GlowingMushroomsBroken = &v
}

// GetHighestCriticalDamage returns the HighestCriticalDamage field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStats) GetHighestCriticalDamage() float64 {
	if o == nil || IsNil(o.HighestCriticalDamage) {
		var ret float64
		return ret
	}
	return *o.HighestCriticalDamage
}

// GetHighestCriticalDamageOk returns a tuple with the HighestCriticalDamage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStats) GetHighestCriticalDamageOk() (*float64, bool) {
	if o == nil || IsNil(o.HighestCriticalDamage) {
		return nil, false
	}
	return o.HighestCriticalDamage, true
}

// HasHighestCriticalDamage returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStats) HasHighestCriticalDamage() bool {
	if o != nil && !IsNil(o.HighestCriticalDamage) {
		return true
	}

	return false
}

// SetHighestCriticalDamage gets a reference to the given float64 and assigns it to the HighestCriticalDamage field.
func (o *SkyBlockProfileMemberPlayerStats) SetHighestCriticalDamage(v float64) {
	o.HighestCriticalDamage = &v
}

// GetHighestDamage returns the HighestDamage field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStats) GetHighestDamage() float64 {
	if o == nil || IsNil(o.HighestDamage) {
		var ret float64
		return ret
	}
	return *o.HighestDamage
}

// GetHighestDamageOk returns a tuple with the HighestDamage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStats) GetHighestDamageOk() (*float64, bool) {
	if o == nil || IsNil(o.HighestDamage) {
		return nil, false
	}
	return o.HighestDamage, true
}

// HasHighestDamage returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStats) HasHighestDamage() bool {
	if o != nil && !IsNil(o.HighestDamage) {
		return true
	}

	return false
}

// SetHighestDamage gets a reference to the given float64 and assigns it to the HighestDamage field.
func (o *SkyBlockProfileMemberPlayerStats) SetHighestDamage(v float64) {
	o.HighestDamage = &v
}

// GetItemsFished returns the ItemsFished field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStats) GetItemsFished() SkyBlockProfileMemberPlayerStatsItemsFished {
	if o == nil || IsNil(o.ItemsFished) {
		var ret SkyBlockProfileMemberPlayerStatsItemsFished
		return ret
	}
	return *o.ItemsFished
}

// GetItemsFishedOk returns a tuple with the ItemsFished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStats) GetItemsFishedOk() (*SkyBlockProfileMemberPlayerStatsItemsFished, bool) {
	if o == nil || IsNil(o.ItemsFished) {
		return nil, false
	}
	return o.ItemsFished, true
}

// HasItemsFished returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStats) HasItemsFished() bool {
	if o != nil && !IsNil(o.ItemsFished) {
		return true
	}

	return false
}

// SetItemsFished gets a reference to the given SkyBlockProfileMemberPlayerStatsItemsFished and assigns it to the ItemsFished field.
func (o *SkyBlockProfileMemberPlayerStats) SetItemsFished(v SkyBlockProfileMemberPlayerStatsItemsFished) {
	o.ItemsFished = &v
}

// GetKills returns the Kills field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStats) GetKills() map[string]float64 {
	if o == nil || IsNil(o.Kills) {
		var ret map[string]float64
		return ret
	}
	return *o.Kills
}

// GetKillsOk returns a tuple with the Kills field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStats) GetKillsOk() (*map[string]float64, bool) {
	if o == nil || IsNil(o.Kills) {
		return nil, false
	}
	return o.Kills, true
}

// HasKills returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStats) HasKills() bool {
	if o != nil && !IsNil(o.Kills) {
		return true
	}

	return false
}

// SetKills gets a reference to the given map[string]float64 and assigns it to the Kills field.
func (o *SkyBlockProfileMemberPlayerStats) SetKills(v map[string]float64) {
	o.Kills = &v
}

// GetMythos returns the Mythos field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStats) GetMythos() SkyBlockProfileMemberPlayerStatsMythos {
	if o == nil || IsNil(o.Mythos) {
		var ret SkyBlockProfileMemberPlayerStatsMythos
		return ret
	}
	return *o.Mythos
}

// GetMythosOk returns a tuple with the Mythos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStats) GetMythosOk() (*SkyBlockProfileMemberPlayerStatsMythos, bool) {
	if o == nil || IsNil(o.Mythos) {
		return nil, false
	}
	return o.Mythos, true
}

// HasMythos returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStats) HasMythos() bool {
	if o != nil && !IsNil(o.Mythos) {
		return true
	}

	return false
}

// SetMythos gets a reference to the given SkyBlockProfileMemberPlayerStatsMythos and assigns it to the Mythos field.
func (o *SkyBlockProfileMemberPlayerStats) SetMythos(v SkyBlockProfileMemberPlayerStatsMythos) {
	o.Mythos = &v
}

// GetPets returns the Pets field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStats) GetPets() SkyBlockProfileMemberPlayerStatsPets {
	if o == nil || IsNil(o.Pets) {
		var ret SkyBlockProfileMemberPlayerStatsPets
		return ret
	}
	return *o.Pets
}

// GetPetsOk returns a tuple with the Pets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStats) GetPetsOk() (*SkyBlockProfileMemberPlayerStatsPets, bool) {
	if o == nil || IsNil(o.Pets) {
		return nil, false
	}
	return o.Pets, true
}

// HasPets returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStats) HasPets() bool {
	if o != nil && !IsNil(o.Pets) {
		return true
	}

	return false
}

// SetPets gets a reference to the given SkyBlockProfileMemberPlayerStatsPets and assigns it to the Pets field.
func (o *SkyBlockProfileMemberPlayerStats) SetPets(v SkyBlockProfileMemberPlayerStatsPets) {
	o.Pets = &v
}

// GetRaces returns the Races field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStats) GetRaces() SkyBlockProfileMemberPlayerStatsRaces {
	if o == nil || IsNil(o.Races) {
		var ret SkyBlockProfileMemberPlayerStatsRaces
		return ret
	}
	return *o.Races
}

// GetRacesOk returns a tuple with the Races field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStats) GetRacesOk() (*SkyBlockProfileMemberPlayerStatsRaces, bool) {
	if o == nil || IsNil(o.Races) {
		return nil, false
	}
	return o.Races, true
}

// HasRaces returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStats) HasRaces() bool {
	if o != nil && !IsNil(o.Races) {
		return true
	}

	return false
}

// SetRaces gets a reference to the given SkyBlockProfileMemberPlayerStatsRaces and assigns it to the Races field.
func (o *SkyBlockProfileMemberPlayerStats) SetRaces(v SkyBlockProfileMemberPlayerStatsRaces) {
	o.Races = &v
}

// GetRift returns the Rift field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStats) GetRift() SkyBlockProfileMemberPlayerStatsRift {
	if o == nil || IsNil(o.Rift) {
		var ret SkyBlockProfileMemberPlayerStatsRift
		return ret
	}
	return *o.Rift
}

// GetRiftOk returns a tuple with the Rift field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStats) GetRiftOk() (*SkyBlockProfileMemberPlayerStatsRift, bool) {
	if o == nil || IsNil(o.Rift) {
		return nil, false
	}
	return o.Rift, true
}

// HasRift returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStats) HasRift() bool {
	if o != nil && !IsNil(o.Rift) {
		return true
	}

	return false
}

// SetRift gets a reference to the given SkyBlockProfileMemberPlayerStatsRift and assigns it to the Rift field.
func (o *SkyBlockProfileMemberPlayerStats) SetRift(v SkyBlockProfileMemberPlayerStatsRift) {
	o.Rift = &v
}

// GetSeaCreatureKills returns the SeaCreatureKills field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStats) GetSeaCreatureKills() float64 {
	if o == nil || IsNil(o.SeaCreatureKills) {
		var ret float64
		return ret
	}
	return *o.SeaCreatureKills
}

// GetSeaCreatureKillsOk returns a tuple with the SeaCreatureKills field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStats) GetSeaCreatureKillsOk() (*float64, bool) {
	if o == nil || IsNil(o.SeaCreatureKills) {
		return nil, false
	}
	return o.SeaCreatureKills, true
}

// HasSeaCreatureKills returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStats) HasSeaCreatureKills() bool {
	if o != nil && !IsNil(o.SeaCreatureKills) {
		return true
	}

	return false
}

// SetSeaCreatureKills gets a reference to the given float64 and assigns it to the SeaCreatureKills field.
func (o *SkyBlockProfileMemberPlayerStats) SetSeaCreatureKills(v float64) {
	o.SeaCreatureKills = &v
}

// GetShredderRod returns the ShredderRod field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStats) GetShredderRod() SkyBlockProfileMemberPlayerStatsShredderRod {
	if o == nil || IsNil(o.ShredderRod) {
		var ret SkyBlockProfileMemberPlayerStatsShredderRod
		return ret
	}
	return *o.ShredderRod
}

// GetShredderRodOk returns a tuple with the ShredderRod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStats) GetShredderRodOk() (*SkyBlockProfileMemberPlayerStatsShredderRod, bool) {
	if o == nil || IsNil(o.ShredderRod) {
		return nil, false
	}
	return o.ShredderRod, true
}

// HasShredderRod returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStats) HasShredderRod() bool {
	if o != nil && !IsNil(o.ShredderRod) {
		return true
	}

	return false
}

// SetShredderRod gets a reference to the given SkyBlockProfileMemberPlayerStatsShredderRod and assigns it to the ShredderRod field.
func (o *SkyBlockProfileMemberPlayerStats) SetShredderRod(v SkyBlockProfileMemberPlayerStatsShredderRod) {
	o.ShredderRod = &v
}

// GetSpooky returns the Spooky field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStats) GetSpooky() SkyBlockProfileMemberPlayerStatsSpooky {
	if o == nil || IsNil(o.Spooky) {
		var ret SkyBlockProfileMemberPlayerStatsSpooky
		return ret
	}
	return *o.Spooky
}

// GetSpookyOk returns a tuple with the Spooky field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStats) GetSpookyOk() (*SkyBlockProfileMemberPlayerStatsSpooky, bool) {
	if o == nil || IsNil(o.Spooky) {
		return nil, false
	}
	return o.Spooky, true
}

// HasSpooky returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStats) HasSpooky() bool {
	if o != nil && !IsNil(o.Spooky) {
		return true
	}

	return false
}

// SetSpooky gets a reference to the given SkyBlockProfileMemberPlayerStatsSpooky and assigns it to the Spooky field.
func (o *SkyBlockProfileMemberPlayerStats) SetSpooky(v SkyBlockProfileMemberPlayerStatsSpooky) {
	o.Spooky = &v
}

// GetWinter returns the Winter field value if set, zero value otherwise.
func (o *SkyBlockProfileMemberPlayerStats) GetWinter() SkyBlockProfileMemberPlayerStatsWinter {
	if o == nil || IsNil(o.Winter) {
		var ret SkyBlockProfileMemberPlayerStatsWinter
		return ret
	}
	return *o.Winter
}

// GetWinterOk returns a tuple with the Winter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkyBlockProfileMemberPlayerStats) GetWinterOk() (*SkyBlockProfileMemberPlayerStatsWinter, bool) {
	if o == nil || IsNil(o.Winter) {
		return nil, false
	}
	return o.Winter, true
}

// HasWinter returns a boolean if a field has been set.
func (o *SkyBlockProfileMemberPlayerStats) HasWinter() bool {
	if o != nil && !IsNil(o.Winter) {
		return true
	}

	return false
}

// SetWinter gets a reference to the given SkyBlockProfileMemberPlayerStatsWinter and assigns it to the Winter field.
func (o *SkyBlockProfileMemberPlayerStats) SetWinter(v SkyBlockProfileMemberPlayerStatsWinter) {
	o.Winter = &v
}

func (o SkyBlockProfileMemberPlayerStats) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkyBlockProfileMemberPlayerStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Auctions) {
		toSerialize["auctions"] = o.Auctions
	}
	if !IsNil(o.CandyCollected) {
		toSerialize["candy_collected"] = o.CandyCollected
	}
	if !IsNil(o.CropsMined) {
		toSerialize["crops_mined"] = o.CropsMined
	}
	if !IsNil(o.Deaths) {
		toSerialize["deaths"] = o.Deaths
	}
	if !IsNil(o.EndIsland) {
		toSerialize["end_island"] = o.EndIsland
	}
	if !IsNil(o.Gifts) {
		toSerialize["gifts"] = o.Gifts
	}
	if !IsNil(o.GlowingMushroomsBroken) {
		toSerialize["glowing_mushrooms_broken"] = o.GlowingMushroomsBroken
	}
	if !IsNil(o.HighestCriticalDamage) {
		toSerialize["highest_critical_damage"] = o.HighestCriticalDamage
	}
	if !IsNil(o.HighestDamage) {
		toSerialize["highest_damage"] = o.HighestDamage
	}
	if !IsNil(o.ItemsFished) {
		toSerialize["items_fished"] = o.ItemsFished
	}
	if !IsNil(o.Kills) {
		toSerialize["kills"] = o.Kills
	}
	if !IsNil(o.Mythos) {
		toSerialize["mythos"] = o.Mythos
	}
	if !IsNil(o.Pets) {
		toSerialize["pets"] = o.Pets
	}
	if !IsNil(o.Races) {
		toSerialize["races"] = o.Races
	}
	if !IsNil(o.Rift) {
		toSerialize["rift"] = o.Rift
	}
	if !IsNil(o.SeaCreatureKills) {
		toSerialize["sea_creature_kills"] = o.SeaCreatureKills
	}
	if !IsNil(o.ShredderRod) {
		toSerialize["shredder_rod"] = o.ShredderRod
	}
	if !IsNil(o.Spooky) {
		toSerialize["spooky"] = o.Spooky
	}
	if !IsNil(o.Winter) {
		toSerialize["winter"] = o.Winter
	}
	return toSerialize, nil
}

type NullableSkyBlockProfileMemberPlayerStats struct {
	value *SkyBlockProfileMemberPlayerStats
	isSet bool
}

func (v NullableSkyBlockProfileMemberPlayerStats) Get() *SkyBlockProfileMemberPlayerStats {
	return v.value
}

func (v *NullableSkyBlockProfileMemberPlayerStats) Set(val *SkyBlockProfileMemberPlayerStats) {
	v.value = val
	v.isSet = true
}

func (v NullableSkyBlockProfileMemberPlayerStats) IsSet() bool {
	return v.isSet
}

func (v *NullableSkyBlockProfileMemberPlayerStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkyBlockProfileMemberPlayerStats(val *SkyBlockProfileMemberPlayerStats) *NullableSkyBlockProfileMemberPlayerStats {
	return &NullableSkyBlockProfileMemberPlayerStats{value: val, isSet: true}
}

func (v NullableSkyBlockProfileMemberPlayerStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkyBlockProfileMemberPlayerStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
