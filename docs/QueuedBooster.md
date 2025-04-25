# QueuedBooster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**PurchaserUuid** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **float32** |  | [optional] 
**OriginalLength** | Pointer to **int32** |  | [optional] 
**Length** | Pointer to **int32** |  | [optional] 
**GameType** | Pointer to **int32** |  | [optional] 
**DateActivated** | Pointer to **int64** |  | [optional] 
**Stacked** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewQueuedBooster

`func NewQueuedBooster() *QueuedBooster`

NewQueuedBooster instantiates a new QueuedBooster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueuedBoosterWithDefaults

`func NewQueuedBoosterWithDefaults() *QueuedBooster`

NewQueuedBoosterWithDefaults instantiates a new QueuedBooster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QueuedBooster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QueuedBooster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QueuedBooster) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *QueuedBooster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPurchaserUuid

`func (o *QueuedBooster) GetPurchaserUuid() string`

GetPurchaserUuid returns the PurchaserUuid field if non-nil, zero value otherwise.

### GetPurchaserUuidOk

`func (o *QueuedBooster) GetPurchaserUuidOk() (*string, bool)`

GetPurchaserUuidOk returns a tuple with the PurchaserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaserUuid

`func (o *QueuedBooster) SetPurchaserUuid(v string)`

SetPurchaserUuid sets PurchaserUuid field to given value.

### HasPurchaserUuid

`func (o *QueuedBooster) HasPurchaserUuid() bool`

HasPurchaserUuid returns a boolean if a field has been set.

### GetAmount

`func (o *QueuedBooster) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QueuedBooster) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QueuedBooster) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *QueuedBooster) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetOriginalLength

`func (o *QueuedBooster) GetOriginalLength() int32`

GetOriginalLength returns the OriginalLength field if non-nil, zero value otherwise.

### GetOriginalLengthOk

`func (o *QueuedBooster) GetOriginalLengthOk() (*int32, bool)`

GetOriginalLengthOk returns a tuple with the OriginalLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalLength

`func (o *QueuedBooster) SetOriginalLength(v int32)`

SetOriginalLength sets OriginalLength field to given value.

### HasOriginalLength

`func (o *QueuedBooster) HasOriginalLength() bool`

HasOriginalLength returns a boolean if a field has been set.

### GetLength

`func (o *QueuedBooster) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *QueuedBooster) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *QueuedBooster) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *QueuedBooster) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetGameType

`func (o *QueuedBooster) GetGameType() int32`

GetGameType returns the GameType field if non-nil, zero value otherwise.

### GetGameTypeOk

`func (o *QueuedBooster) GetGameTypeOk() (*int32, bool)`

GetGameTypeOk returns a tuple with the GameType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameType

`func (o *QueuedBooster) SetGameType(v int32)`

SetGameType sets GameType field to given value.

### HasGameType

`func (o *QueuedBooster) HasGameType() bool`

HasGameType returns a boolean if a field has been set.

### GetDateActivated

`func (o *QueuedBooster) GetDateActivated() int64`

GetDateActivated returns the DateActivated field if non-nil, zero value otherwise.

### GetDateActivatedOk

`func (o *QueuedBooster) GetDateActivatedOk() (*int64, bool)`

GetDateActivatedOk returns a tuple with the DateActivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateActivated

`func (o *QueuedBooster) SetDateActivated(v int64)`

SetDateActivated sets DateActivated field to given value.

### HasDateActivated

`func (o *QueuedBooster) HasDateActivated() bool`

HasDateActivated returns a boolean if a field has been set.

### GetStacked

`func (o *QueuedBooster) GetStacked() bool`

GetStacked returns the Stacked field if non-nil, zero value otherwise.

### GetStackedOk

`func (o *QueuedBooster) GetStackedOk() (*bool, bool)`

GetStackedOk returns a tuple with the Stacked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacked

`func (o *QueuedBooster) SetStacked(v bool)`

SetStacked sets Stacked field to given value.

### HasStacked

`func (o *QueuedBooster) HasStacked() bool`

HasStacked returns a boolean if a field has been set.

### SetStackedNil

`func (o *QueuedBooster) SetStackedNil(b bool)`

 SetStackedNil sets the value for Stacked to be an explicit nil

### UnsetStacked
`func (o *QueuedBooster) UnsetStacked()`

UnsetStacked ensures that no value is present for Stacked, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


