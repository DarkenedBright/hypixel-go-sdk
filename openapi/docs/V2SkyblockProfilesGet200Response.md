# V2SkyblockProfilesGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profiles** | [**[]SkyBlockProfile**](SkyBlockProfile.md) |  | 
**Success** | **bool** |  | 

## Methods

### NewV2SkyblockProfilesGet200Response

`func NewV2SkyblockProfilesGet200Response(profiles []SkyBlockProfile, success bool, ) *V2SkyblockProfilesGet200Response`

NewV2SkyblockProfilesGet200Response instantiates a new V2SkyblockProfilesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2SkyblockProfilesGet200ResponseWithDefaults

`func NewV2SkyblockProfilesGet200ResponseWithDefaults() *V2SkyblockProfilesGet200Response`

NewV2SkyblockProfilesGet200ResponseWithDefaults instantiates a new V2SkyblockProfilesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfiles

`func (o *V2SkyblockProfilesGet200Response) GetProfiles() []SkyBlockProfile`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *V2SkyblockProfilesGet200Response) GetProfilesOk() (*[]SkyBlockProfile, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *V2SkyblockProfilesGet200Response) SetProfiles(v []SkyBlockProfile)`

SetProfiles sets Profiles field to given value.


### GetSuccess

`func (o *V2SkyblockProfilesGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *V2SkyblockProfilesGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *V2SkyblockProfilesGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


