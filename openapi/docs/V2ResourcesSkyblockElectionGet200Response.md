# V2ResourcesSkyblockElectionGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Current** | Pointer to **map[string]interface{}** | Data regarding the current election, will not be provided if there is no open election ongoing | [optional] 
**LastUpdated** | Pointer to **int64** |  | [optional] 
**Mayor** | Pointer to **map[string]interface{}** | Data regarding the current mayor | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewV2ResourcesSkyblockElectionGet200Response

`func NewV2ResourcesSkyblockElectionGet200Response() *V2ResourcesSkyblockElectionGet200Response`

NewV2ResourcesSkyblockElectionGet200Response instantiates a new V2ResourcesSkyblockElectionGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ResourcesSkyblockElectionGet200ResponseWithDefaults

`func NewV2ResourcesSkyblockElectionGet200ResponseWithDefaults() *V2ResourcesSkyblockElectionGet200Response`

NewV2ResourcesSkyblockElectionGet200ResponseWithDefaults instantiates a new V2ResourcesSkyblockElectionGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrent

`func (o *V2ResourcesSkyblockElectionGet200Response) GetCurrent() map[string]interface{}`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *V2ResourcesSkyblockElectionGet200Response) GetCurrentOk() (*map[string]interface{}, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *V2ResourcesSkyblockElectionGet200Response) SetCurrent(v map[string]interface{})`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *V2ResourcesSkyblockElectionGet200Response) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetLastUpdated

`func (o *V2ResourcesSkyblockElectionGet200Response) GetLastUpdated() int64`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *V2ResourcesSkyblockElectionGet200Response) GetLastUpdatedOk() (*int64, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *V2ResourcesSkyblockElectionGet200Response) SetLastUpdated(v int64)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *V2ResourcesSkyblockElectionGet200Response) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetMayor

`func (o *V2ResourcesSkyblockElectionGet200Response) GetMayor() map[string]interface{}`

GetMayor returns the Mayor field if non-nil, zero value otherwise.

### GetMayorOk

`func (o *V2ResourcesSkyblockElectionGet200Response) GetMayorOk() (*map[string]interface{}, bool)`

GetMayorOk returns a tuple with the Mayor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMayor

`func (o *V2ResourcesSkyblockElectionGet200Response) SetMayor(v map[string]interface{})`

SetMayor sets Mayor field to given value.

### HasMayor

`func (o *V2ResourcesSkyblockElectionGet200Response) HasMayor() bool`

HasMayor returns a boolean if a field has been set.

### GetSuccess

`func (o *V2ResourcesSkyblockElectionGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *V2ResourcesSkyblockElectionGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *V2ResourcesSkyblockElectionGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *V2ResourcesSkyblockElectionGet200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


