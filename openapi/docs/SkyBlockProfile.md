# SkyBlockProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **string** |  | [optional] 
**Members** | Pointer to [**map[string]SkyBlockProfileMember**](SkyBlockProfileMember.md) | A map of profile member UUIDs to profile member objects | [optional] 
**CuteName** | Pointer to **NullableString** | The cute name of the profile, only provided on the profiles endpoint | [optional] 
**Selected** | Pointer to **NullableBool** | Whether or not this is the currently selected profile, only provided on the profiles endpoint | [optional] 
**CommunityUpgrades** | Pointer to **map[string]interface{}** |  | [optional] 
**Banking** | Pointer to [**NullableSkyBlockProfileBanking**](SkyBlockProfileBanking.md) |  | [optional] 
**GameMode** | Pointer to **NullableString** | The SkyBlock game mode of the profile, not present if normal mode | [optional] 

## Methods

### NewSkyBlockProfile

`func NewSkyBlockProfile() *SkyBlockProfile`

NewSkyBlockProfile instantiates a new SkyBlockProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkyBlockProfileWithDefaults

`func NewSkyBlockProfileWithDefaults() *SkyBlockProfile`

NewSkyBlockProfileWithDefaults instantiates a new SkyBlockProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *SkyBlockProfile) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *SkyBlockProfile) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *SkyBlockProfile) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *SkyBlockProfile) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetMembers

`func (o *SkyBlockProfile) GetMembers() map[string]SkyBlockProfileMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *SkyBlockProfile) GetMembersOk() (*map[string]SkyBlockProfileMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *SkyBlockProfile) SetMembers(v map[string]SkyBlockProfileMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *SkyBlockProfile) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetCuteName

`func (o *SkyBlockProfile) GetCuteName() string`

GetCuteName returns the CuteName field if non-nil, zero value otherwise.

### GetCuteNameOk

`func (o *SkyBlockProfile) GetCuteNameOk() (*string, bool)`

GetCuteNameOk returns a tuple with the CuteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCuteName

`func (o *SkyBlockProfile) SetCuteName(v string)`

SetCuteName sets CuteName field to given value.

### HasCuteName

`func (o *SkyBlockProfile) HasCuteName() bool`

HasCuteName returns a boolean if a field has been set.

### SetCuteNameNil

`func (o *SkyBlockProfile) SetCuteNameNil(b bool)`

 SetCuteNameNil sets the value for CuteName to be an explicit nil

### UnsetCuteName
`func (o *SkyBlockProfile) UnsetCuteName()`

UnsetCuteName ensures that no value is present for CuteName, not even an explicit nil
### GetSelected

`func (o *SkyBlockProfile) GetSelected() bool`

GetSelected returns the Selected field if non-nil, zero value otherwise.

### GetSelectedOk

`func (o *SkyBlockProfile) GetSelectedOk() (*bool, bool)`

GetSelectedOk returns a tuple with the Selected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelected

`func (o *SkyBlockProfile) SetSelected(v bool)`

SetSelected sets Selected field to given value.

### HasSelected

`func (o *SkyBlockProfile) HasSelected() bool`

HasSelected returns a boolean if a field has been set.

### SetSelectedNil

`func (o *SkyBlockProfile) SetSelectedNil(b bool)`

 SetSelectedNil sets the value for Selected to be an explicit nil

### UnsetSelected
`func (o *SkyBlockProfile) UnsetSelected()`

UnsetSelected ensures that no value is present for Selected, not even an explicit nil
### GetCommunityUpgrades

`func (o *SkyBlockProfile) GetCommunityUpgrades() map[string]interface{}`

GetCommunityUpgrades returns the CommunityUpgrades field if non-nil, zero value otherwise.

### GetCommunityUpgradesOk

`func (o *SkyBlockProfile) GetCommunityUpgradesOk() (*map[string]interface{}, bool)`

GetCommunityUpgradesOk returns a tuple with the CommunityUpgrades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityUpgrades

`func (o *SkyBlockProfile) SetCommunityUpgrades(v map[string]interface{})`

SetCommunityUpgrades sets CommunityUpgrades field to given value.

### HasCommunityUpgrades

`func (o *SkyBlockProfile) HasCommunityUpgrades() bool`

HasCommunityUpgrades returns a boolean if a field has been set.

### SetCommunityUpgradesNil

`func (o *SkyBlockProfile) SetCommunityUpgradesNil(b bool)`

 SetCommunityUpgradesNil sets the value for CommunityUpgrades to be an explicit nil

### UnsetCommunityUpgrades
`func (o *SkyBlockProfile) UnsetCommunityUpgrades()`

UnsetCommunityUpgrades ensures that no value is present for CommunityUpgrades, not even an explicit nil
### GetBanking

`func (o *SkyBlockProfile) GetBanking() SkyBlockProfileBanking`

GetBanking returns the Banking field if non-nil, zero value otherwise.

### GetBankingOk

`func (o *SkyBlockProfile) GetBankingOk() (*SkyBlockProfileBanking, bool)`

GetBankingOk returns a tuple with the Banking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanking

`func (o *SkyBlockProfile) SetBanking(v SkyBlockProfileBanking)`

SetBanking sets Banking field to given value.

### HasBanking

`func (o *SkyBlockProfile) HasBanking() bool`

HasBanking returns a boolean if a field has been set.

### SetBankingNil

`func (o *SkyBlockProfile) SetBankingNil(b bool)`

 SetBankingNil sets the value for Banking to be an explicit nil

### UnsetBanking
`func (o *SkyBlockProfile) UnsetBanking()`

UnsetBanking ensures that no value is present for Banking, not even an explicit nil
### GetGameMode

`func (o *SkyBlockProfile) GetGameMode() string`

GetGameMode returns the GameMode field if non-nil, zero value otherwise.

### GetGameModeOk

`func (o *SkyBlockProfile) GetGameModeOk() (*string, bool)`

GetGameModeOk returns a tuple with the GameMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameMode

`func (o *SkyBlockProfile) SetGameMode(v string)`

SetGameMode sets GameMode field to given value.

### HasGameMode

`func (o *SkyBlockProfile) HasGameMode() bool`

HasGameMode returns a boolean if a field has been set.

### SetGameModeNil

`func (o *SkyBlockProfile) SetGameModeNil(b bool)`

 SetGameModeNil sets the value for GameMode to be an explicit nil

### UnsetGameMode
`func (o *SkyBlockProfile) UnsetGameMode()`

UnsetGameMode ensures that no value is present for GameMode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


