# V2BoostersGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Boosters** | Pointer to [**[]V2BoostersGet200ResponseBoostersInner**](V2BoostersGet200ResponseBoostersInner.md) |  | [optional] 
**BoosterState** | Pointer to [**V2BoostersGet200ResponseBoosterState**](V2BoostersGet200ResponseBoosterState.md) |  | [optional] 

## Methods

### NewV2BoostersGet200Response

`func NewV2BoostersGet200Response() *V2BoostersGet200Response`

NewV2BoostersGet200Response instantiates a new V2BoostersGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2BoostersGet200ResponseWithDefaults

`func NewV2BoostersGet200ResponseWithDefaults() *V2BoostersGet200Response`

NewV2BoostersGet200ResponseWithDefaults instantiates a new V2BoostersGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *V2BoostersGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *V2BoostersGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *V2BoostersGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *V2BoostersGet200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetBoosters

`func (o *V2BoostersGet200Response) GetBoosters() []V2BoostersGet200ResponseBoostersInner`

GetBoosters returns the Boosters field if non-nil, zero value otherwise.

### GetBoostersOk

`func (o *V2BoostersGet200Response) GetBoostersOk() (*[]V2BoostersGet200ResponseBoostersInner, bool)`

GetBoostersOk returns a tuple with the Boosters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoosters

`func (o *V2BoostersGet200Response) SetBoosters(v []V2BoostersGet200ResponseBoostersInner)`

SetBoosters sets Boosters field to given value.

### HasBoosters

`func (o *V2BoostersGet200Response) HasBoosters() bool`

HasBoosters returns a boolean if a field has been set.

### GetBoosterState

`func (o *V2BoostersGet200Response) GetBoosterState() V2BoostersGet200ResponseBoosterState`

GetBoosterState returns the BoosterState field if non-nil, zero value otherwise.

### GetBoosterStateOk

`func (o *V2BoostersGet200Response) GetBoosterStateOk() (*V2BoostersGet200ResponseBoosterState, bool)`

GetBoosterStateOk returns a tuple with the BoosterState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoosterState

`func (o *V2BoostersGet200Response) SetBoosterState(v V2BoostersGet200ResponseBoosterState)`

SetBoosterState sets BoosterState field to given value.

### HasBoosterState

`func (o *V2BoostersGet200Response) HasBoosterState() bool`

HasBoosterState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


