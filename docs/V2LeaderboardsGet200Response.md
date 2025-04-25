# V2LeaderboardsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Leaderboards** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewV2LeaderboardsGet200Response

`func NewV2LeaderboardsGet200Response() *V2LeaderboardsGet200Response`

NewV2LeaderboardsGet200Response instantiates a new V2LeaderboardsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2LeaderboardsGet200ResponseWithDefaults

`func NewV2LeaderboardsGet200ResponseWithDefaults() *V2LeaderboardsGet200Response`

NewV2LeaderboardsGet200ResponseWithDefaults instantiates a new V2LeaderboardsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *V2LeaderboardsGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *V2LeaderboardsGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *V2LeaderboardsGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *V2LeaderboardsGet200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetLeaderboards

`func (o *V2LeaderboardsGet200Response) GetLeaderboards() map[string]interface{}`

GetLeaderboards returns the Leaderboards field if non-nil, zero value otherwise.

### GetLeaderboardsOk

`func (o *V2LeaderboardsGet200Response) GetLeaderboardsOk() (*map[string]interface{}, bool)`

GetLeaderboardsOk returns a tuple with the Leaderboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaderboards

`func (o *V2LeaderboardsGet200Response) SetLeaderboards(v map[string]interface{})`

SetLeaderboards sets Leaderboards field to given value.

### HasLeaderboards

`func (o *V2LeaderboardsGet200Response) HasLeaderboards() bool`

HasLeaderboards returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


