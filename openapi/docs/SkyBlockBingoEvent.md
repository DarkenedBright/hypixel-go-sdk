# SkyBlockBingoEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedGoals** | [**[]SkyBlockBingoEventCompletedGoalsInner**](SkyBlockBingoEventCompletedGoalsInner.md) |  | 
**Key** | **int64** |  | 
**Points** | **int64** |  | 

## Methods

### NewSkyBlockBingoEvent

`func NewSkyBlockBingoEvent(completedGoals []SkyBlockBingoEventCompletedGoalsInner, key int64, points int64, ) *SkyBlockBingoEvent`

NewSkyBlockBingoEvent instantiates a new SkyBlockBingoEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkyBlockBingoEventWithDefaults

`func NewSkyBlockBingoEventWithDefaults() *SkyBlockBingoEvent`

NewSkyBlockBingoEventWithDefaults instantiates a new SkyBlockBingoEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedGoals

`func (o *SkyBlockBingoEvent) GetCompletedGoals() []SkyBlockBingoEventCompletedGoalsInner`

GetCompletedGoals returns the CompletedGoals field if non-nil, zero value otherwise.

### GetCompletedGoalsOk

`func (o *SkyBlockBingoEvent) GetCompletedGoalsOk() (*[]SkyBlockBingoEventCompletedGoalsInner, bool)`

GetCompletedGoalsOk returns a tuple with the CompletedGoals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedGoals

`func (o *SkyBlockBingoEvent) SetCompletedGoals(v []SkyBlockBingoEventCompletedGoalsInner)`

SetCompletedGoals sets CompletedGoals field to given value.


### GetKey

`func (o *SkyBlockBingoEvent) GetKey() int64`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SkyBlockBingoEvent) GetKeyOk() (*int64, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SkyBlockBingoEvent) SetKey(v int64)`

SetKey sets Key field to given value.


### GetPoints

`func (o *SkyBlockBingoEvent) GetPoints() int64`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *SkyBlockBingoEvent) GetPointsOk() (*int64, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *SkyBlockBingoEvent) SetPoints(v int64)`

SetPoints sets Points field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


