# Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Session** | [**StatusSession**](StatusSession.md) |  | 
**Success** | **bool** |  | 
**Uuid** | **string** |  | 

## Methods

### NewStatus

`func NewStatus(session StatusSession, success bool, uuid string, ) *Status`

NewStatus instantiates a new Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusWithDefaults

`func NewStatusWithDefaults() *Status`

NewStatusWithDefaults instantiates a new Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSession

`func (o *Status) GetSession() StatusSession`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *Status) GetSessionOk() (*StatusSession, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *Status) SetSession(v StatusSession)`

SetSession sets Session field to given value.


### GetSuccess

`func (o *Status) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *Status) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *Status) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetUuid

`func (o *Status) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Status) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Status) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


