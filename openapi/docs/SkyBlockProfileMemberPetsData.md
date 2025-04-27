# SkyBlockProfileMemberPetsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autopet** | Pointer to [**SkyBlockProfileMemberPetsDataAutopet**](SkyBlockProfileMemberPetsDataAutopet.md) |  | [optional] 
**PetCare** | Pointer to [**SkyBlockProfileMemberPetsDataPetCare**](SkyBlockProfileMemberPetsDataPetCare.md) |  | [optional] 
**Pets** | [**[]SkyBlockProfileMemberPetsDataPet**](SkyBlockProfileMemberPetsDataPet.md) |  | 

## Methods

### NewSkyBlockProfileMemberPetsData

`func NewSkyBlockProfileMemberPetsData(pets []SkyBlockProfileMemberPetsDataPet, ) *SkyBlockProfileMemberPetsData`

NewSkyBlockProfileMemberPetsData instantiates a new SkyBlockProfileMemberPetsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkyBlockProfileMemberPetsDataWithDefaults

`func NewSkyBlockProfileMemberPetsDataWithDefaults() *SkyBlockProfileMemberPetsData`

NewSkyBlockProfileMemberPetsDataWithDefaults instantiates a new SkyBlockProfileMemberPetsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutopet

`func (o *SkyBlockProfileMemberPetsData) GetAutopet() SkyBlockProfileMemberPetsDataAutopet`

GetAutopet returns the Autopet field if non-nil, zero value otherwise.

### GetAutopetOk

`func (o *SkyBlockProfileMemberPetsData) GetAutopetOk() (*SkyBlockProfileMemberPetsDataAutopet, bool)`

GetAutopetOk returns a tuple with the Autopet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutopet

`func (o *SkyBlockProfileMemberPetsData) SetAutopet(v SkyBlockProfileMemberPetsDataAutopet)`

SetAutopet sets Autopet field to given value.

### HasAutopet

`func (o *SkyBlockProfileMemberPetsData) HasAutopet() bool`

HasAutopet returns a boolean if a field has been set.

### GetPetCare

`func (o *SkyBlockProfileMemberPetsData) GetPetCare() SkyBlockProfileMemberPetsDataPetCare`

GetPetCare returns the PetCare field if non-nil, zero value otherwise.

### GetPetCareOk

`func (o *SkyBlockProfileMemberPetsData) GetPetCareOk() (*SkyBlockProfileMemberPetsDataPetCare, bool)`

GetPetCareOk returns a tuple with the PetCare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPetCare

`func (o *SkyBlockProfileMemberPetsData) SetPetCare(v SkyBlockProfileMemberPetsDataPetCare)`

SetPetCare sets PetCare field to given value.

### HasPetCare

`func (o *SkyBlockProfileMemberPetsData) HasPetCare() bool`

HasPetCare returns a boolean if a field has been set.

### GetPets

`func (o *SkyBlockProfileMemberPetsData) GetPets() []SkyBlockProfileMemberPetsDataPet`

GetPets returns the Pets field if non-nil, zero value otherwise.

### GetPetsOk

`func (o *SkyBlockProfileMemberPetsData) GetPetsOk() (*[]SkyBlockProfileMemberPetsDataPet, bool)`

GetPetsOk returns a tuple with the Pets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPets

`func (o *SkyBlockProfileMemberPetsData) SetPets(v []SkyBlockProfileMemberPetsDataPet)`

SetPets sets Pets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


