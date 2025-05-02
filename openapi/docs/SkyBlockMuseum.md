# SkyBlockMuseum

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | [**map[string]SkyBlockMuseumMember**](SkyBlockMuseumMember.md) |  | 
**Success** | **bool** |  | 

## Methods

### NewSkyBlockMuseum

`func NewSkyBlockMuseum(members map[string]SkyBlockMuseumMember, success bool, ) *SkyBlockMuseum`

NewSkyBlockMuseum instantiates a new SkyBlockMuseum object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkyBlockMuseumWithDefaults

`func NewSkyBlockMuseumWithDefaults() *SkyBlockMuseum`

NewSkyBlockMuseumWithDefaults instantiates a new SkyBlockMuseum object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *SkyBlockMuseum) GetMembers() map[string]SkyBlockMuseumMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *SkyBlockMuseum) GetMembersOk() (*map[string]SkyBlockMuseumMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *SkyBlockMuseum) SetMembers(v map[string]SkyBlockMuseumMember)`

SetMembers sets Members field to given value.


### GetSuccess

`func (o *SkyBlockMuseum) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SkyBlockMuseum) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SkyBlockMuseum) SetSuccess(v bool)`

SetSuccess sets Success field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


