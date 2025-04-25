# SkyBlockItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier for this item | [optional] 
**Material** | Pointer to **string** | The Bukkit material enum value for the item | [optional] 
**Name** | Pointer to **string** | The name of the item | [optional] 
**Tier** | Pointer to **string** | The rarity tier of the item | [optional] 
**Color** | Pointer to **string** | The color metadata to be applied to an item, usually leather armor pieces | [optional] 
**Skin** | Pointer to **string** | The skin value for a skull based item | [optional] 

## Methods

### NewSkyBlockItem

`func NewSkyBlockItem() *SkyBlockItem`

NewSkyBlockItem instantiates a new SkyBlockItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkyBlockItemWithDefaults

`func NewSkyBlockItemWithDefaults() *SkyBlockItem`

NewSkyBlockItemWithDefaults instantiates a new SkyBlockItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SkyBlockItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SkyBlockItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SkyBlockItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SkyBlockItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaterial

`func (o *SkyBlockItem) GetMaterial() string`

GetMaterial returns the Material field if non-nil, zero value otherwise.

### GetMaterialOk

`func (o *SkyBlockItem) GetMaterialOk() (*string, bool)`

GetMaterialOk returns a tuple with the Material field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterial

`func (o *SkyBlockItem) SetMaterial(v string)`

SetMaterial sets Material field to given value.

### HasMaterial

`func (o *SkyBlockItem) HasMaterial() bool`

HasMaterial returns a boolean if a field has been set.

### GetName

`func (o *SkyBlockItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SkyBlockItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SkyBlockItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SkyBlockItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTier

`func (o *SkyBlockItem) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *SkyBlockItem) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *SkyBlockItem) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *SkyBlockItem) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetColor

`func (o *SkyBlockItem) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *SkyBlockItem) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *SkyBlockItem) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *SkyBlockItem) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetSkin

`func (o *SkyBlockItem) GetSkin() string`

GetSkin returns the Skin field if non-nil, zero value otherwise.

### GetSkinOk

`func (o *SkyBlockItem) GetSkinOk() (*string, bool)`

GetSkinOk returns a tuple with the Skin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkin

`func (o *SkyBlockItem) SetSkin(v string)`

SetSkin sets Skin field to given value.

### HasSkin

`func (o *SkyBlockItem) HasSkin() bool`

HasSkin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


