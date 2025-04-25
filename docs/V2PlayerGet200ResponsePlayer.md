# V2PlayerGet200ResponsePlayer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**Displayname** | Pointer to **NullableString** |  | [optional] 
**Rank** | Pointer to **NullableString** |  | [optional] 
**PackageRank** | Pointer to **NullableString** |  | [optional] 
**NewPackageRank** | Pointer to **NullableString** |  | [optional] 
**MonthlyPackageRank** | Pointer to **NullableString** |  | [optional] 
**FirstLogin** | Pointer to **NullableFloat32** |  | [optional] 
**LastLogin** | Pointer to **NullableFloat32** |  | [optional] 
**LastLogout** | Pointer to **NullableFloat32** |  | [optional] 
**Stats** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewV2PlayerGet200ResponsePlayer

`func NewV2PlayerGet200ResponsePlayer() *V2PlayerGet200ResponsePlayer`

NewV2PlayerGet200ResponsePlayer instantiates a new V2PlayerGet200ResponsePlayer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2PlayerGet200ResponsePlayerWithDefaults

`func NewV2PlayerGet200ResponsePlayerWithDefaults() *V2PlayerGet200ResponsePlayer`

NewV2PlayerGet200ResponsePlayerWithDefaults instantiates a new V2PlayerGet200ResponsePlayer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *V2PlayerGet200ResponsePlayer) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *V2PlayerGet200ResponsePlayer) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *V2PlayerGet200ResponsePlayer) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *V2PlayerGet200ResponsePlayer) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetDisplayname

`func (o *V2PlayerGet200ResponsePlayer) GetDisplayname() string`

GetDisplayname returns the Displayname field if non-nil, zero value otherwise.

### GetDisplaynameOk

`func (o *V2PlayerGet200ResponsePlayer) GetDisplaynameOk() (*string, bool)`

GetDisplaynameOk returns a tuple with the Displayname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayname

`func (o *V2PlayerGet200ResponsePlayer) SetDisplayname(v string)`

SetDisplayname sets Displayname field to given value.

### HasDisplayname

`func (o *V2PlayerGet200ResponsePlayer) HasDisplayname() bool`

HasDisplayname returns a boolean if a field has been set.

### SetDisplaynameNil

`func (o *V2PlayerGet200ResponsePlayer) SetDisplaynameNil(b bool)`

 SetDisplaynameNil sets the value for Displayname to be an explicit nil

### UnsetDisplayname
`func (o *V2PlayerGet200ResponsePlayer) UnsetDisplayname()`

UnsetDisplayname ensures that no value is present for Displayname, not even an explicit nil
### GetRank

`func (o *V2PlayerGet200ResponsePlayer) GetRank() string`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *V2PlayerGet200ResponsePlayer) GetRankOk() (*string, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *V2PlayerGet200ResponsePlayer) SetRank(v string)`

SetRank sets Rank field to given value.

### HasRank

`func (o *V2PlayerGet200ResponsePlayer) HasRank() bool`

HasRank returns a boolean if a field has been set.

### SetRankNil

`func (o *V2PlayerGet200ResponsePlayer) SetRankNil(b bool)`

 SetRankNil sets the value for Rank to be an explicit nil

### UnsetRank
`func (o *V2PlayerGet200ResponsePlayer) UnsetRank()`

UnsetRank ensures that no value is present for Rank, not even an explicit nil
### GetPackageRank

`func (o *V2PlayerGet200ResponsePlayer) GetPackageRank() string`

GetPackageRank returns the PackageRank field if non-nil, zero value otherwise.

### GetPackageRankOk

`func (o *V2PlayerGet200ResponsePlayer) GetPackageRankOk() (*string, bool)`

GetPackageRankOk returns a tuple with the PackageRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageRank

`func (o *V2PlayerGet200ResponsePlayer) SetPackageRank(v string)`

SetPackageRank sets PackageRank field to given value.

### HasPackageRank

`func (o *V2PlayerGet200ResponsePlayer) HasPackageRank() bool`

HasPackageRank returns a boolean if a field has been set.

### SetPackageRankNil

`func (o *V2PlayerGet200ResponsePlayer) SetPackageRankNil(b bool)`

 SetPackageRankNil sets the value for PackageRank to be an explicit nil

### UnsetPackageRank
`func (o *V2PlayerGet200ResponsePlayer) UnsetPackageRank()`

UnsetPackageRank ensures that no value is present for PackageRank, not even an explicit nil
### GetNewPackageRank

`func (o *V2PlayerGet200ResponsePlayer) GetNewPackageRank() string`

GetNewPackageRank returns the NewPackageRank field if non-nil, zero value otherwise.

### GetNewPackageRankOk

`func (o *V2PlayerGet200ResponsePlayer) GetNewPackageRankOk() (*string, bool)`

GetNewPackageRankOk returns a tuple with the NewPackageRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPackageRank

`func (o *V2PlayerGet200ResponsePlayer) SetNewPackageRank(v string)`

SetNewPackageRank sets NewPackageRank field to given value.

### HasNewPackageRank

`func (o *V2PlayerGet200ResponsePlayer) HasNewPackageRank() bool`

HasNewPackageRank returns a boolean if a field has been set.

### SetNewPackageRankNil

`func (o *V2PlayerGet200ResponsePlayer) SetNewPackageRankNil(b bool)`

 SetNewPackageRankNil sets the value for NewPackageRank to be an explicit nil

### UnsetNewPackageRank
`func (o *V2PlayerGet200ResponsePlayer) UnsetNewPackageRank()`

UnsetNewPackageRank ensures that no value is present for NewPackageRank, not even an explicit nil
### GetMonthlyPackageRank

`func (o *V2PlayerGet200ResponsePlayer) GetMonthlyPackageRank() string`

GetMonthlyPackageRank returns the MonthlyPackageRank field if non-nil, zero value otherwise.

### GetMonthlyPackageRankOk

`func (o *V2PlayerGet200ResponsePlayer) GetMonthlyPackageRankOk() (*string, bool)`

GetMonthlyPackageRankOk returns a tuple with the MonthlyPackageRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyPackageRank

`func (o *V2PlayerGet200ResponsePlayer) SetMonthlyPackageRank(v string)`

SetMonthlyPackageRank sets MonthlyPackageRank field to given value.

### HasMonthlyPackageRank

`func (o *V2PlayerGet200ResponsePlayer) HasMonthlyPackageRank() bool`

HasMonthlyPackageRank returns a boolean if a field has been set.

### SetMonthlyPackageRankNil

`func (o *V2PlayerGet200ResponsePlayer) SetMonthlyPackageRankNil(b bool)`

 SetMonthlyPackageRankNil sets the value for MonthlyPackageRank to be an explicit nil

### UnsetMonthlyPackageRank
`func (o *V2PlayerGet200ResponsePlayer) UnsetMonthlyPackageRank()`

UnsetMonthlyPackageRank ensures that no value is present for MonthlyPackageRank, not even an explicit nil
### GetFirstLogin

`func (o *V2PlayerGet200ResponsePlayer) GetFirstLogin() float32`

GetFirstLogin returns the FirstLogin field if non-nil, zero value otherwise.

### GetFirstLoginOk

`func (o *V2PlayerGet200ResponsePlayer) GetFirstLoginOk() (*float32, bool)`

GetFirstLoginOk returns a tuple with the FirstLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstLogin

`func (o *V2PlayerGet200ResponsePlayer) SetFirstLogin(v float32)`

SetFirstLogin sets FirstLogin field to given value.

### HasFirstLogin

`func (o *V2PlayerGet200ResponsePlayer) HasFirstLogin() bool`

HasFirstLogin returns a boolean if a field has been set.

### SetFirstLoginNil

`func (o *V2PlayerGet200ResponsePlayer) SetFirstLoginNil(b bool)`

 SetFirstLoginNil sets the value for FirstLogin to be an explicit nil

### UnsetFirstLogin
`func (o *V2PlayerGet200ResponsePlayer) UnsetFirstLogin()`

UnsetFirstLogin ensures that no value is present for FirstLogin, not even an explicit nil
### GetLastLogin

`func (o *V2PlayerGet200ResponsePlayer) GetLastLogin() float32`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *V2PlayerGet200ResponsePlayer) GetLastLoginOk() (*float32, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *V2PlayerGet200ResponsePlayer) SetLastLogin(v float32)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *V2PlayerGet200ResponsePlayer) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### SetLastLoginNil

`func (o *V2PlayerGet200ResponsePlayer) SetLastLoginNil(b bool)`

 SetLastLoginNil sets the value for LastLogin to be an explicit nil

### UnsetLastLogin
`func (o *V2PlayerGet200ResponsePlayer) UnsetLastLogin()`

UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
### GetLastLogout

`func (o *V2PlayerGet200ResponsePlayer) GetLastLogout() float32`

GetLastLogout returns the LastLogout field if non-nil, zero value otherwise.

### GetLastLogoutOk

`func (o *V2PlayerGet200ResponsePlayer) GetLastLogoutOk() (*float32, bool)`

GetLastLogoutOk returns a tuple with the LastLogout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogout

`func (o *V2PlayerGet200ResponsePlayer) SetLastLogout(v float32)`

SetLastLogout sets LastLogout field to given value.

### HasLastLogout

`func (o *V2PlayerGet200ResponsePlayer) HasLastLogout() bool`

HasLastLogout returns a boolean if a field has been set.

### SetLastLogoutNil

`func (o *V2PlayerGet200ResponsePlayer) SetLastLogoutNil(b bool)`

 SetLastLogoutNil sets the value for LastLogout to be an explicit nil

### UnsetLastLogout
`func (o *V2PlayerGet200ResponsePlayer) UnsetLastLogout()`

UnsetLastLogout ensures that no value is present for LastLogout, not even an explicit nil
### GetStats

`func (o *V2PlayerGet200ResponsePlayer) GetStats() map[string]interface{}`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *V2PlayerGet200ResponsePlayer) GetStatsOk() (*map[string]interface{}, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *V2PlayerGet200ResponsePlayer) SetStats(v map[string]interface{})`

SetStats sets Stats field to given value.

### HasStats

`func (o *V2PlayerGet200ResponsePlayer) HasStats() bool`

HasStats returns a boolean if a field has been set.

### SetStatsNil

`func (o *V2PlayerGet200ResponsePlayer) SetStatsNil(b bool)`

 SetStatsNil sets the value for Stats to be an explicit nil

### UnsetStats
`func (o *V2PlayerGet200ResponsePlayer) UnsetStats()`

UnsetStats ensures that no value is present for Stats, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


