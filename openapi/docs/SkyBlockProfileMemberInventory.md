# SkyBlockProfileMemberInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackpackContents** | Pointer to [**map[string]SkyBlockProfileMemberInventoryBase64GzipData**](SkyBlockProfileMemberInventoryBase64GzipData.md) |  | [optional] 
**BackpackIcons** | Pointer to [**map[string]SkyBlockProfileMemberInventoryBase64GzipData**](SkyBlockProfileMemberInventoryBase64GzipData.md) |  | [optional] 
**BagContents** | Pointer to [**SkyBlockProfileMemberInventoryBagContents**](SkyBlockProfileMemberInventoryBagContents.md) |  | [optional] 
**EnderChestContents** | Pointer to [**SkyBlockProfileMemberInventoryBase64GzipData**](SkyBlockProfileMemberInventoryBase64GzipData.md) |  | [optional] 
**EquipmentContents** | Pointer to [**SkyBlockProfileMemberInventoryBase64GzipData**](SkyBlockProfileMemberInventoryBase64GzipData.md) |  | [optional] 
**InvArmor** | Pointer to [**SkyBlockProfileMemberInventoryBase64GzipData**](SkyBlockProfileMemberInventoryBase64GzipData.md) |  | [optional] 
**InvContents** | Pointer to [**SkyBlockProfileMemberInventoryBase64GzipData**](SkyBlockProfileMemberInventoryBase64GzipData.md) |  | [optional] 
**PersonalVaultContents** | Pointer to [**SkyBlockProfileMemberInventoryBase64GzipData**](SkyBlockProfileMemberInventoryBase64GzipData.md) |  | [optional] 
**SacksCounts** | Pointer to **map[string]int64** |  | [optional] 
**WardrobeContents** | Pointer to [**SkyBlockProfileMemberInventoryBase64GzipData**](SkyBlockProfileMemberInventoryBase64GzipData.md) |  | [optional] 
**WardrobeEquippedSlot** | Pointer to **int64** |  | [optional] 

## Methods

### NewSkyBlockProfileMemberInventory

`func NewSkyBlockProfileMemberInventory() *SkyBlockProfileMemberInventory`

NewSkyBlockProfileMemberInventory instantiates a new SkyBlockProfileMemberInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkyBlockProfileMemberInventoryWithDefaults

`func NewSkyBlockProfileMemberInventoryWithDefaults() *SkyBlockProfileMemberInventory`

NewSkyBlockProfileMemberInventoryWithDefaults instantiates a new SkyBlockProfileMemberInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackpackContents

`func (o *SkyBlockProfileMemberInventory) GetBackpackContents() map[string]SkyBlockProfileMemberInventoryBase64GzipData`

GetBackpackContents returns the BackpackContents field if non-nil, zero value otherwise.

### GetBackpackContentsOk

`func (o *SkyBlockProfileMemberInventory) GetBackpackContentsOk() (*map[string]SkyBlockProfileMemberInventoryBase64GzipData, bool)`

GetBackpackContentsOk returns a tuple with the BackpackContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackpackContents

`func (o *SkyBlockProfileMemberInventory) SetBackpackContents(v map[string]SkyBlockProfileMemberInventoryBase64GzipData)`

SetBackpackContents sets BackpackContents field to given value.

### HasBackpackContents

`func (o *SkyBlockProfileMemberInventory) HasBackpackContents() bool`

HasBackpackContents returns a boolean if a field has been set.

### GetBackpackIcons

`func (o *SkyBlockProfileMemberInventory) GetBackpackIcons() map[string]SkyBlockProfileMemberInventoryBase64GzipData`

GetBackpackIcons returns the BackpackIcons field if non-nil, zero value otherwise.

### GetBackpackIconsOk

`func (o *SkyBlockProfileMemberInventory) GetBackpackIconsOk() (*map[string]SkyBlockProfileMemberInventoryBase64GzipData, bool)`

GetBackpackIconsOk returns a tuple with the BackpackIcons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackpackIcons

`func (o *SkyBlockProfileMemberInventory) SetBackpackIcons(v map[string]SkyBlockProfileMemberInventoryBase64GzipData)`

SetBackpackIcons sets BackpackIcons field to given value.

### HasBackpackIcons

`func (o *SkyBlockProfileMemberInventory) HasBackpackIcons() bool`

HasBackpackIcons returns a boolean if a field has been set.

### GetBagContents

`func (o *SkyBlockProfileMemberInventory) GetBagContents() SkyBlockProfileMemberInventoryBagContents`

GetBagContents returns the BagContents field if non-nil, zero value otherwise.

### GetBagContentsOk

`func (o *SkyBlockProfileMemberInventory) GetBagContentsOk() (*SkyBlockProfileMemberInventoryBagContents, bool)`

GetBagContentsOk returns a tuple with the BagContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBagContents

`func (o *SkyBlockProfileMemberInventory) SetBagContents(v SkyBlockProfileMemberInventoryBagContents)`

SetBagContents sets BagContents field to given value.

### HasBagContents

`func (o *SkyBlockProfileMemberInventory) HasBagContents() bool`

HasBagContents returns a boolean if a field has been set.

### GetEnderChestContents

`func (o *SkyBlockProfileMemberInventory) GetEnderChestContents() SkyBlockProfileMemberInventoryBase64GzipData`

GetEnderChestContents returns the EnderChestContents field if non-nil, zero value otherwise.

### GetEnderChestContentsOk

`func (o *SkyBlockProfileMemberInventory) GetEnderChestContentsOk() (*SkyBlockProfileMemberInventoryBase64GzipData, bool)`

GetEnderChestContentsOk returns a tuple with the EnderChestContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnderChestContents

`func (o *SkyBlockProfileMemberInventory) SetEnderChestContents(v SkyBlockProfileMemberInventoryBase64GzipData)`

SetEnderChestContents sets EnderChestContents field to given value.

### HasEnderChestContents

`func (o *SkyBlockProfileMemberInventory) HasEnderChestContents() bool`

HasEnderChestContents returns a boolean if a field has been set.

### GetEquipmentContents

`func (o *SkyBlockProfileMemberInventory) GetEquipmentContents() SkyBlockProfileMemberInventoryBase64GzipData`

GetEquipmentContents returns the EquipmentContents field if non-nil, zero value otherwise.

### GetEquipmentContentsOk

`func (o *SkyBlockProfileMemberInventory) GetEquipmentContentsOk() (*SkyBlockProfileMemberInventoryBase64GzipData, bool)`

GetEquipmentContentsOk returns a tuple with the EquipmentContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentContents

`func (o *SkyBlockProfileMemberInventory) SetEquipmentContents(v SkyBlockProfileMemberInventoryBase64GzipData)`

SetEquipmentContents sets EquipmentContents field to given value.

### HasEquipmentContents

`func (o *SkyBlockProfileMemberInventory) HasEquipmentContents() bool`

HasEquipmentContents returns a boolean if a field has been set.

### GetInvArmor

`func (o *SkyBlockProfileMemberInventory) GetInvArmor() SkyBlockProfileMemberInventoryBase64GzipData`

GetInvArmor returns the InvArmor field if non-nil, zero value otherwise.

### GetInvArmorOk

`func (o *SkyBlockProfileMemberInventory) GetInvArmorOk() (*SkyBlockProfileMemberInventoryBase64GzipData, bool)`

GetInvArmorOk returns a tuple with the InvArmor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvArmor

`func (o *SkyBlockProfileMemberInventory) SetInvArmor(v SkyBlockProfileMemberInventoryBase64GzipData)`

SetInvArmor sets InvArmor field to given value.

### HasInvArmor

`func (o *SkyBlockProfileMemberInventory) HasInvArmor() bool`

HasInvArmor returns a boolean if a field has been set.

### GetInvContents

`func (o *SkyBlockProfileMemberInventory) GetInvContents() SkyBlockProfileMemberInventoryBase64GzipData`

GetInvContents returns the InvContents field if non-nil, zero value otherwise.

### GetInvContentsOk

`func (o *SkyBlockProfileMemberInventory) GetInvContentsOk() (*SkyBlockProfileMemberInventoryBase64GzipData, bool)`

GetInvContentsOk returns a tuple with the InvContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvContents

`func (o *SkyBlockProfileMemberInventory) SetInvContents(v SkyBlockProfileMemberInventoryBase64GzipData)`

SetInvContents sets InvContents field to given value.

### HasInvContents

`func (o *SkyBlockProfileMemberInventory) HasInvContents() bool`

HasInvContents returns a boolean if a field has been set.

### GetPersonalVaultContents

`func (o *SkyBlockProfileMemberInventory) GetPersonalVaultContents() SkyBlockProfileMemberInventoryBase64GzipData`

GetPersonalVaultContents returns the PersonalVaultContents field if non-nil, zero value otherwise.

### GetPersonalVaultContentsOk

`func (o *SkyBlockProfileMemberInventory) GetPersonalVaultContentsOk() (*SkyBlockProfileMemberInventoryBase64GzipData, bool)`

GetPersonalVaultContentsOk returns a tuple with the PersonalVaultContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalVaultContents

`func (o *SkyBlockProfileMemberInventory) SetPersonalVaultContents(v SkyBlockProfileMemberInventoryBase64GzipData)`

SetPersonalVaultContents sets PersonalVaultContents field to given value.

### HasPersonalVaultContents

`func (o *SkyBlockProfileMemberInventory) HasPersonalVaultContents() bool`

HasPersonalVaultContents returns a boolean if a field has been set.

### GetSacksCounts

`func (o *SkyBlockProfileMemberInventory) GetSacksCounts() map[string]int64`

GetSacksCounts returns the SacksCounts field if non-nil, zero value otherwise.

### GetSacksCountsOk

`func (o *SkyBlockProfileMemberInventory) GetSacksCountsOk() (*map[string]int64, bool)`

GetSacksCountsOk returns a tuple with the SacksCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSacksCounts

`func (o *SkyBlockProfileMemberInventory) SetSacksCounts(v map[string]int64)`

SetSacksCounts sets SacksCounts field to given value.

### HasSacksCounts

`func (o *SkyBlockProfileMemberInventory) HasSacksCounts() bool`

HasSacksCounts returns a boolean if a field has been set.

### GetWardrobeContents

`func (o *SkyBlockProfileMemberInventory) GetWardrobeContents() SkyBlockProfileMemberInventoryBase64GzipData`

GetWardrobeContents returns the WardrobeContents field if non-nil, zero value otherwise.

### GetWardrobeContentsOk

`func (o *SkyBlockProfileMemberInventory) GetWardrobeContentsOk() (*SkyBlockProfileMemberInventoryBase64GzipData, bool)`

GetWardrobeContentsOk returns a tuple with the WardrobeContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWardrobeContents

`func (o *SkyBlockProfileMemberInventory) SetWardrobeContents(v SkyBlockProfileMemberInventoryBase64GzipData)`

SetWardrobeContents sets WardrobeContents field to given value.

### HasWardrobeContents

`func (o *SkyBlockProfileMemberInventory) HasWardrobeContents() bool`

HasWardrobeContents returns a boolean if a field has been set.

### GetWardrobeEquippedSlot

`func (o *SkyBlockProfileMemberInventory) GetWardrobeEquippedSlot() int64`

GetWardrobeEquippedSlot returns the WardrobeEquippedSlot field if non-nil, zero value otherwise.

### GetWardrobeEquippedSlotOk

`func (o *SkyBlockProfileMemberInventory) GetWardrobeEquippedSlotOk() (*int64, bool)`

GetWardrobeEquippedSlotOk returns a tuple with the WardrobeEquippedSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWardrobeEquippedSlot

`func (o *SkyBlockProfileMemberInventory) SetWardrobeEquippedSlot(v int64)`

SetWardrobeEquippedSlot sets WardrobeEquippedSlot field to given value.

### HasWardrobeEquippedSlot

`func (o *SkyBlockProfileMemberInventory) HasWardrobeEquippedSlot() bool`

HasWardrobeEquippedSlot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


