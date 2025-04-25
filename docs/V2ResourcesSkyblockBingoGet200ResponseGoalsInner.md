# V2ResourcesSkyblockBingoGet200ResponseGoalsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The backend ID for this goal | 
**Name** | **string** | The user friendly display name for this goal | 
**Lore** | Pointer to **string** | Description of this goal | [optional] 
**FullLore** | Pointer to **[]string** | The full description of this goal | [optional] 
**Tiers** | Pointer to **[]float32** | The tiers of this goal, if a global goal | [optional] 
**Progress** | Pointer to **float32** | The global progress of this goal | [optional] 
**RequiredAmount** | Pointer to **float32** | The required amount for this specific goal | [optional] 

## Methods

### NewV2ResourcesSkyblockBingoGet200ResponseGoalsInner

`func NewV2ResourcesSkyblockBingoGet200ResponseGoalsInner(id string, name string, ) *V2ResourcesSkyblockBingoGet200ResponseGoalsInner`

NewV2ResourcesSkyblockBingoGet200ResponseGoalsInner instantiates a new V2ResourcesSkyblockBingoGet200ResponseGoalsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ResourcesSkyblockBingoGet200ResponseGoalsInnerWithDefaults

`func NewV2ResourcesSkyblockBingoGet200ResponseGoalsInnerWithDefaults() *V2ResourcesSkyblockBingoGet200ResponseGoalsInner`

NewV2ResourcesSkyblockBingoGet200ResponseGoalsInnerWithDefaults instantiates a new V2ResourcesSkyblockBingoGet200ResponseGoalsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V2ResourcesSkyblockBingoGet200ResponseGoalsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2ResourcesSkyblockBingoGet200ResponseGoalsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2ResourcesSkyblockBingoGet200ResponseGoalsInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *V2ResourcesSkyblockBingoGet200ResponseGoalsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2ResourcesSkyblockBingoGet200ResponseGoalsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2ResourcesSkyblockBingoGet200ResponseGoalsInner) SetName(v string)`

SetName sets Name field to given value.


### GetLore

`func (o *V2ResourcesSkyblockBingoGet200ResponseGoalsInner) GetLore() string`

GetLore returns the Lore field if non-nil, zero value otherwise.

### GetLoreOk

`func (o *V2ResourcesSkyblockBingoGet200ResponseGoalsInner) GetLoreOk() (*string, bool)`

GetLoreOk returns a tuple with the Lore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLore

`func (o *V2ResourcesSkyblockBingoGet200ResponseGoalsInner) SetLore(v string)`

SetLore sets Lore field to given value.

### HasLore

`func (o *V2ResourcesSkyblockBingoGet200ResponseGoalsInner) HasLore() bool`

HasLore returns a boolean if a field has been set.

### GetFullLore

`func (o *V2ResourcesSkyblockBingoGet200ResponseGoalsInner) GetFullLore() []string`

GetFullLore returns the FullLore field if non-nil, zero value otherwise.

### GetFullLoreOk

`func (o *V2ResourcesSkyblockBingoGet200ResponseGoalsInner) GetFullLoreOk() (*[]string, bool)`

GetFullLoreOk returns a tuple with the FullLore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullLore

`func (o *V2ResourcesSkyblockBingoGet200ResponseGoalsInner) SetFullLore(v []string)`

SetFullLore sets FullLore field to given value.

### HasFullLore

`func (o *V2ResourcesSkyblockBingoGet200ResponseGoalsInner) HasFullLore() bool`

HasFullLore returns a boolean if a field has been set.

### GetTiers

`func (o *V2ResourcesSkyblockBingoGet200ResponseGoalsInner) GetTiers() []float32`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *V2ResourcesSkyblockBingoGet200ResponseGoalsInner) GetTiersOk() (*[]float32, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *V2ResourcesSkyblockBingoGet200ResponseGoalsInner) SetTiers(v []float32)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *V2ResourcesSkyblockBingoGet200ResponseGoalsInner) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### GetProgress

`func (o *V2ResourcesSkyblockBingoGet200ResponseGoalsInner) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *V2ResourcesSkyblockBingoGet200ResponseGoalsInner) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *V2ResourcesSkyblockBingoGet200ResponseGoalsInner) SetProgress(v float32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *V2ResourcesSkyblockBingoGet200ResponseGoalsInner) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetRequiredAmount

`func (o *V2ResourcesSkyblockBingoGet200ResponseGoalsInner) GetRequiredAmount() float32`

GetRequiredAmount returns the RequiredAmount field if non-nil, zero value otherwise.

### GetRequiredAmountOk

`func (o *V2ResourcesSkyblockBingoGet200ResponseGoalsInner) GetRequiredAmountOk() (*float32, bool)`

GetRequiredAmountOk returns a tuple with the RequiredAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredAmount

`func (o *V2ResourcesSkyblockBingoGet200ResponseGoalsInner) SetRequiredAmount(v float32)`

SetRequiredAmount sets RequiredAmount field to given value.

### HasRequiredAmount

`func (o *V2ResourcesSkyblockBingoGet200ResponseGoalsInner) HasRequiredAmount() bool`

HasRequiredAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


