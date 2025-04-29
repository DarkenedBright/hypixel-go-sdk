# StatusGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Session** | [**StatusGetSession**](StatusGetSession.md) |  | 
**Success** | **bool** |  | 
**Uuid** | **string** |  | 

## Methods

### NewStatusGet

`func NewStatusGet(session StatusGetSession, success bool, uuid string, ) *StatusGet`

NewStatusGet instantiates a new StatusGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusGetWithDefaults

`func NewStatusGetWithDefaults() *StatusGet`

NewStatusGetWithDefaults instantiates a new StatusGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSession

`func (o *StatusGet) GetSession() StatusGetSession`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *StatusGet) GetSessionOk() (*StatusGetSession, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *StatusGet) SetSession(v StatusGetSession)`

SetSession sets Session field to given value.


### GetSuccess

`func (o *StatusGet) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *StatusGet) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *StatusGet) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetUuid

`func (o *StatusGet) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StatusGet) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StatusGet) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


