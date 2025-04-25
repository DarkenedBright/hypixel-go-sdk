# HousingHouse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | The UUID of this house. | 
**Owner** | **string** | The UUID of the owner of this house. | 
**Name** | **NullableString** | The name of this house, may contain Minecraft color symbols. | 
**CreatedAt** | **float32** | The time this house was created. | 
**Players** | **float32** | The number of players in this house. | 
**Cookies** | [**HousingHouseCookies**](HousingHouseCookies.md) |  | 

## Methods

### NewHousingHouse

`func NewHousingHouse(uuid string, owner string, name NullableString, createdAt float32, players float32, cookies HousingHouseCookies, ) *HousingHouse`

NewHousingHouse instantiates a new HousingHouse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHousingHouseWithDefaults

`func NewHousingHouseWithDefaults() *HousingHouse`

NewHousingHouseWithDefaults instantiates a new HousingHouse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *HousingHouse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HousingHouse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HousingHouse) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetOwner

`func (o *HousingHouse) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *HousingHouse) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *HousingHouse) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetName

`func (o *HousingHouse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HousingHouse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HousingHouse) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *HousingHouse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HousingHouse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCreatedAt

`func (o *HousingHouse) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HousingHouse) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HousingHouse) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.


### GetPlayers

`func (o *HousingHouse) GetPlayers() float32`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *HousingHouse) GetPlayersOk() (*float32, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *HousingHouse) SetPlayers(v float32)`

SetPlayers sets Players field to given value.


### GetCookies

`func (o *HousingHouse) GetCookies() HousingHouseCookies`

GetCookies returns the Cookies field if non-nil, zero value otherwise.

### GetCookiesOk

`func (o *HousingHouse) GetCookiesOk() (*HousingHouseCookies, bool)`

GetCookiesOk returns a tuple with the Cookies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookies

`func (o *HousingHouse) SetCookies(v HousingHouseCookies)`

SetCookies sets Cookies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


