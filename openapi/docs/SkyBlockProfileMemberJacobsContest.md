# SkyBlockProfileMemberJacobsContest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contests** | Pointer to [**map[string]SkyBlockProfileMemberJacobsContestContestsValue**](SkyBlockProfileMemberJacobsContestContestsValue.md) |  | [optional] 
**MedalsInv** | Pointer to [**SkyBlockProfileMemberJacobsContestMedalsInv**](SkyBlockProfileMemberJacobsContestMedalsInv.md) |  | [optional] 
**Migration** | Pointer to **bool** |  | [optional] 
**Perks** | Pointer to [**SkyBlockProfileMemberJacobsContestPerks**](SkyBlockProfileMemberJacobsContestPerks.md) |  | [optional] 
**PersonalBests** | Pointer to [**SkyBlockProfileMemberJacobsContestPersonalBests**](SkyBlockProfileMemberJacobsContestPersonalBests.md) |  | [optional] 
**Talked** | Pointer to **bool** |  | [optional] 
**UniqueBrackets** | Pointer to [**SkyBlockProfileMemberJacobsContestUniqueBrackets**](SkyBlockProfileMemberJacobsContestUniqueBrackets.md) |  | [optional] 

## Methods

### NewSkyBlockProfileMemberJacobsContest

`func NewSkyBlockProfileMemberJacobsContest() *SkyBlockProfileMemberJacobsContest`

NewSkyBlockProfileMemberJacobsContest instantiates a new SkyBlockProfileMemberJacobsContest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkyBlockProfileMemberJacobsContestWithDefaults

`func NewSkyBlockProfileMemberJacobsContestWithDefaults() *SkyBlockProfileMemberJacobsContest`

NewSkyBlockProfileMemberJacobsContestWithDefaults instantiates a new SkyBlockProfileMemberJacobsContest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContests

`func (o *SkyBlockProfileMemberJacobsContest) GetContests() map[string]SkyBlockProfileMemberJacobsContestContestsValue`

GetContests returns the Contests field if non-nil, zero value otherwise.

### GetContestsOk

`func (o *SkyBlockProfileMemberJacobsContest) GetContestsOk() (*map[string]SkyBlockProfileMemberJacobsContestContestsValue, bool)`

GetContestsOk returns a tuple with the Contests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContests

`func (o *SkyBlockProfileMemberJacobsContest) SetContests(v map[string]SkyBlockProfileMemberJacobsContestContestsValue)`

SetContests sets Contests field to given value.

### HasContests

`func (o *SkyBlockProfileMemberJacobsContest) HasContests() bool`

HasContests returns a boolean if a field has been set.

### GetMedalsInv

`func (o *SkyBlockProfileMemberJacobsContest) GetMedalsInv() SkyBlockProfileMemberJacobsContestMedalsInv`

GetMedalsInv returns the MedalsInv field if non-nil, zero value otherwise.

### GetMedalsInvOk

`func (o *SkyBlockProfileMemberJacobsContest) GetMedalsInvOk() (*SkyBlockProfileMemberJacobsContestMedalsInv, bool)`

GetMedalsInvOk returns a tuple with the MedalsInv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedalsInv

`func (o *SkyBlockProfileMemberJacobsContest) SetMedalsInv(v SkyBlockProfileMemberJacobsContestMedalsInv)`

SetMedalsInv sets MedalsInv field to given value.

### HasMedalsInv

`func (o *SkyBlockProfileMemberJacobsContest) HasMedalsInv() bool`

HasMedalsInv returns a boolean if a field has been set.

### GetMigration

`func (o *SkyBlockProfileMemberJacobsContest) GetMigration() bool`

GetMigration returns the Migration field if non-nil, zero value otherwise.

### GetMigrationOk

`func (o *SkyBlockProfileMemberJacobsContest) GetMigrationOk() (*bool, bool)`

GetMigrationOk returns a tuple with the Migration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigration

`func (o *SkyBlockProfileMemberJacobsContest) SetMigration(v bool)`

SetMigration sets Migration field to given value.

### HasMigration

`func (o *SkyBlockProfileMemberJacobsContest) HasMigration() bool`

HasMigration returns a boolean if a field has been set.

### GetPerks

`func (o *SkyBlockProfileMemberJacobsContest) GetPerks() SkyBlockProfileMemberJacobsContestPerks`

GetPerks returns the Perks field if non-nil, zero value otherwise.

### GetPerksOk

`func (o *SkyBlockProfileMemberJacobsContest) GetPerksOk() (*SkyBlockProfileMemberJacobsContestPerks, bool)`

GetPerksOk returns a tuple with the Perks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerks

`func (o *SkyBlockProfileMemberJacobsContest) SetPerks(v SkyBlockProfileMemberJacobsContestPerks)`

SetPerks sets Perks field to given value.

### HasPerks

`func (o *SkyBlockProfileMemberJacobsContest) HasPerks() bool`

HasPerks returns a boolean if a field has been set.

### GetPersonalBests

`func (o *SkyBlockProfileMemberJacobsContest) GetPersonalBests() SkyBlockProfileMemberJacobsContestPersonalBests`

GetPersonalBests returns the PersonalBests field if non-nil, zero value otherwise.

### GetPersonalBestsOk

`func (o *SkyBlockProfileMemberJacobsContest) GetPersonalBestsOk() (*SkyBlockProfileMemberJacobsContestPersonalBests, bool)`

GetPersonalBestsOk returns a tuple with the PersonalBests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalBests

`func (o *SkyBlockProfileMemberJacobsContest) SetPersonalBests(v SkyBlockProfileMemberJacobsContestPersonalBests)`

SetPersonalBests sets PersonalBests field to given value.

### HasPersonalBests

`func (o *SkyBlockProfileMemberJacobsContest) HasPersonalBests() bool`

HasPersonalBests returns a boolean if a field has been set.

### GetTalked

`func (o *SkyBlockProfileMemberJacobsContest) GetTalked() bool`

GetTalked returns the Talked field if non-nil, zero value otherwise.

### GetTalkedOk

`func (o *SkyBlockProfileMemberJacobsContest) GetTalkedOk() (*bool, bool)`

GetTalkedOk returns a tuple with the Talked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTalked

`func (o *SkyBlockProfileMemberJacobsContest) SetTalked(v bool)`

SetTalked sets Talked field to given value.

### HasTalked

`func (o *SkyBlockProfileMemberJacobsContest) HasTalked() bool`

HasTalked returns a boolean if a field has been set.

### GetUniqueBrackets

`func (o *SkyBlockProfileMemberJacobsContest) GetUniqueBrackets() SkyBlockProfileMemberJacobsContestUniqueBrackets`

GetUniqueBrackets returns the UniqueBrackets field if non-nil, zero value otherwise.

### GetUniqueBracketsOk

`func (o *SkyBlockProfileMemberJacobsContest) GetUniqueBracketsOk() (*SkyBlockProfileMemberJacobsContestUniqueBrackets, bool)`

GetUniqueBracketsOk returns a tuple with the UniqueBrackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueBrackets

`func (o *SkyBlockProfileMemberJacobsContest) SetUniqueBrackets(v SkyBlockProfileMemberJacobsContestUniqueBrackets)`

SetUniqueBrackets sets UniqueBrackets field to given value.

### HasUniqueBrackets

`func (o *SkyBlockProfileMemberJacobsContest) HasUniqueBrackets() bool`

HasUniqueBrackets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


