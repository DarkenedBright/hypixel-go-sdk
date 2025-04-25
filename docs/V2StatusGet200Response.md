# V2StatusGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Session** | Pointer to [**V2StatusGet200ResponseSession**](V2StatusGet200ResponseSession.md) |  | [optional] 

## Methods

### NewV2StatusGet200Response

`func NewV2StatusGet200Response() *V2StatusGet200Response`

NewV2StatusGet200Response instantiates a new V2StatusGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2StatusGet200ResponseWithDefaults

`func NewV2StatusGet200ResponseWithDefaults() *V2StatusGet200Response`

NewV2StatusGet200ResponseWithDefaults instantiates a new V2StatusGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *V2StatusGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *V2StatusGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *V2StatusGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *V2StatusGet200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetUuid

`func (o *V2StatusGet200Response) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *V2StatusGet200Response) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *V2StatusGet200Response) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *V2StatusGet200Response) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetSession

`func (o *V2StatusGet200Response) GetSession() V2StatusGet200ResponseSession`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *V2StatusGet200Response) GetSessionOk() (*V2StatusGet200ResponseSession, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *V2StatusGet200Response) SetSession(v V2StatusGet200ResponseSession)`

SetSession sets Session field to given value.

### HasSession

`func (o *V2StatusGet200Response) HasSession() bool`

HasSession returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


