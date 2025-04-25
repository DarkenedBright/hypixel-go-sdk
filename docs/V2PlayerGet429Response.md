# V2PlayerGet429Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Cause** | Pointer to **string** |  | [optional] 
**Throttle** | Pointer to **bool** |  | [optional] 
**Global** | Pointer to **bool** | When this boolean exists and is true, the throttle occurring is a global throttle applied to all users | [optional] 

## Methods

### NewV2PlayerGet429Response

`func NewV2PlayerGet429Response() *V2PlayerGet429Response`

NewV2PlayerGet429Response instantiates a new V2PlayerGet429Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2PlayerGet429ResponseWithDefaults

`func NewV2PlayerGet429ResponseWithDefaults() *V2PlayerGet429Response`

NewV2PlayerGet429ResponseWithDefaults instantiates a new V2PlayerGet429Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *V2PlayerGet429Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *V2PlayerGet429Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *V2PlayerGet429Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *V2PlayerGet429Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetCause

`func (o *V2PlayerGet429Response) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *V2PlayerGet429Response) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *V2PlayerGet429Response) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *V2PlayerGet429Response) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetThrottle

`func (o *V2PlayerGet429Response) GetThrottle() bool`

GetThrottle returns the Throttle field if non-nil, zero value otherwise.

### GetThrottleOk

`func (o *V2PlayerGet429Response) GetThrottleOk() (*bool, bool)`

GetThrottleOk returns a tuple with the Throttle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottle

`func (o *V2PlayerGet429Response) SetThrottle(v bool)`

SetThrottle sets Throttle field to given value.

### HasThrottle

`func (o *V2PlayerGet429Response) HasThrottle() bool`

HasThrottle returns a boolean if a field has been set.

### GetGlobal

`func (o *V2PlayerGet429Response) GetGlobal() bool`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *V2PlayerGet429Response) GetGlobalOk() (*bool, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *V2PlayerGet429Response) SetGlobal(v bool)`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *V2PlayerGet429Response) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


