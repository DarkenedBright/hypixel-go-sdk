# V2ResourcesSkyblockBingoGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**LastUpdated** | **float64** | The unix milliseconds timestamp of the last time this data was updated | 
**Id** | **float64** | The current bingo event ID, increments by 1 for each bingo hosted | 
**Name** | **string** | The display name for the current bingo event | 
**Start** | **float64** | The start time of the current bingo event in unix milliseconds | 
**End** | **float64** | The end time of the current bingo event in unix milliseconds | 
**Modifier** | **string** | The modifier for the current bingo event | 
**Goals** | [**[]V2ResourcesSkyblockBingoGet200ResponseGoalsInner**](V2ResourcesSkyblockBingoGet200ResponseGoalsInner.md) | The goals for the current bingo event, as well as their progress | 

## Methods

### NewV2ResourcesSkyblockBingoGet200Response

`func NewV2ResourcesSkyblockBingoGet200Response(success bool, lastUpdated float64, id float64, name string, start float64, end float64, modifier string, goals []V2ResourcesSkyblockBingoGet200ResponseGoalsInner, ) *V2ResourcesSkyblockBingoGet200Response`

NewV2ResourcesSkyblockBingoGet200Response instantiates a new V2ResourcesSkyblockBingoGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ResourcesSkyblockBingoGet200ResponseWithDefaults

`func NewV2ResourcesSkyblockBingoGet200ResponseWithDefaults() *V2ResourcesSkyblockBingoGet200Response`

NewV2ResourcesSkyblockBingoGet200ResponseWithDefaults instantiates a new V2ResourcesSkyblockBingoGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *V2ResourcesSkyblockBingoGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *V2ResourcesSkyblockBingoGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *V2ResourcesSkyblockBingoGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetLastUpdated

`func (o *V2ResourcesSkyblockBingoGet200Response) GetLastUpdated() float64`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *V2ResourcesSkyblockBingoGet200Response) GetLastUpdatedOk() (*float64, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *V2ResourcesSkyblockBingoGet200Response) SetLastUpdated(v float64)`

SetLastUpdated sets LastUpdated field to given value.


### GetId

`func (o *V2ResourcesSkyblockBingoGet200Response) GetId() float64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2ResourcesSkyblockBingoGet200Response) GetIdOk() (*float64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2ResourcesSkyblockBingoGet200Response) SetId(v float64)`

SetId sets Id field to given value.


### GetName

`func (o *V2ResourcesSkyblockBingoGet200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2ResourcesSkyblockBingoGet200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2ResourcesSkyblockBingoGet200Response) SetName(v string)`

SetName sets Name field to given value.


### GetStart

`func (o *V2ResourcesSkyblockBingoGet200Response) GetStart() float64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *V2ResourcesSkyblockBingoGet200Response) GetStartOk() (*float64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *V2ResourcesSkyblockBingoGet200Response) SetStart(v float64)`

SetStart sets Start field to given value.


### GetEnd

`func (o *V2ResourcesSkyblockBingoGet200Response) GetEnd() float64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *V2ResourcesSkyblockBingoGet200Response) GetEndOk() (*float64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *V2ResourcesSkyblockBingoGet200Response) SetEnd(v float64)`

SetEnd sets End field to given value.


### GetModifier

`func (o *V2ResourcesSkyblockBingoGet200Response) GetModifier() string`

GetModifier returns the Modifier field if non-nil, zero value otherwise.

### GetModifierOk

`func (o *V2ResourcesSkyblockBingoGet200Response) GetModifierOk() (*string, bool)`

GetModifierOk returns a tuple with the Modifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifier

`func (o *V2ResourcesSkyblockBingoGet200Response) SetModifier(v string)`

SetModifier sets Modifier field to given value.


### GetGoals

`func (o *V2ResourcesSkyblockBingoGet200Response) GetGoals() []V2ResourcesSkyblockBingoGet200ResponseGoalsInner`

GetGoals returns the Goals field if non-nil, zero value otherwise.

### GetGoalsOk

`func (o *V2ResourcesSkyblockBingoGet200Response) GetGoalsOk() (*[]V2ResourcesSkyblockBingoGet200ResponseGoalsInner, bool)`

GetGoalsOk returns a tuple with the Goals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoals

`func (o *V2ResourcesSkyblockBingoGet200Response) SetGoals(v []V2ResourcesSkyblockBingoGet200ResponseGoalsInner)`

SetGoals sets Goals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


