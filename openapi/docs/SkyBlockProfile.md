# SkyBlockProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Banking** | Pointer to [**SkyBlockProfileBanking**](SkyBlockProfileBanking.md) |  | [optional] 
**CommunityUpgrades** | Pointer to [**SkyBlockProfileCommunityUpgrades**](SkyBlockProfileCommunityUpgrades.md) |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**CuteName** | [**SkyBlockProfileCuteName**](SkyBlockProfileCuteName.md) |  | 
**GameMode** | Pointer to [**SkyBlockProfileGameMode**](SkyBlockProfileGameMode.md) |  | [optional] 
**Members** | [**map[string]SkyBlockProfileMember**](SkyBlockProfileMember.md) |  | 
**ProfileId** | **string** |  | 
**Selected** | **bool** |  | 

## Methods

### NewSkyBlockProfile

`func NewSkyBlockProfile(cuteName SkyBlockProfileCuteName, members map[string]SkyBlockProfileMember, profileId string, selected bool, ) *SkyBlockProfile`

NewSkyBlockProfile instantiates a new SkyBlockProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkyBlockProfileWithDefaults

`func NewSkyBlockProfileWithDefaults() *SkyBlockProfile`

NewSkyBlockProfileWithDefaults instantiates a new SkyBlockProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### GetCommunityUpgrades

`func (o *SkyBlockProfile) GetCommunityUpgrades() SkyBlockProfileCommunityUpgrades`

GetCommunityUpgrades returns the CommunityUpgrades field if non-nil, zero value otherwise.

### GetCommunityUpgradesOk

`func (o *SkyBlockProfile) GetCommunityUpgradesOk() (*SkyBlockProfileCommunityUpgrades, bool)`

GetCommunityUpgradesOk returns a tuple with the CommunityUpgrades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityUpgrades

`func (o *SkyBlockProfile) SetCommunityUpgrades(v SkyBlockProfileCommunityUpgrades)`

SetCommunityUpgrades sets CommunityUpgrades field to given value.

### HasCommunityUpgrades

`func (o *SkyBlockProfile) HasCommunityUpgrades() bool`

HasCommunityUpgrades returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SkyBlockProfile) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SkyBlockProfile) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SkyBlockProfile) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SkyBlockProfile) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCuteName

`func (o *SkyBlockProfile) GetCuteName() SkyBlockProfileCuteName`

GetCuteName returns the CuteName field if non-nil, zero value otherwise.

### GetCuteNameOk

`func (o *SkyBlockProfile) GetCuteNameOk() (*SkyBlockProfileCuteName, bool)`

GetCuteNameOk returns a tuple with the CuteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCuteName

`func (o *SkyBlockProfile) SetCuteName(v SkyBlockProfileCuteName)`

SetCuteName sets CuteName field to given value.


### GetGameMode

`func (o *SkyBlockProfile) GetGameMode() SkyBlockProfileGameMode`

GetGameMode returns the GameMode field if non-nil, zero value otherwise.

### GetGameModeOk

`func (o *SkyBlockProfile) GetGameModeOk() (*SkyBlockProfileGameMode, bool)`

GetGameModeOk returns a tuple with the GameMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameMode

`func (o *SkyBlockProfile) SetGameMode(v SkyBlockProfileGameMode)`

SetGameMode sets GameMode field to given value.

### HasGameMode

`func (o *SkyBlockProfile) HasGameMode() bool`

HasGameMode returns a boolean if a field has been set.

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


