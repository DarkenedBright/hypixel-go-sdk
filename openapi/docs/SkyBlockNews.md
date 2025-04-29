# SkyBlockNews

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]SkyBlockNewsElement**](SkyBlockNewsElement.md) |  | 
**Success** | **bool** |  | 

## Methods

### NewSkyBlockNews

`func NewSkyBlockNews(items []SkyBlockNewsElement, success bool, ) *SkyBlockNews`

NewSkyBlockNews instantiates a new SkyBlockNews object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkyBlockNewsWithDefaults

`func NewSkyBlockNewsWithDefaults() *SkyBlockNews`

NewSkyBlockNewsWithDefaults instantiates a new SkyBlockNews object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *SkyBlockNews) GetItems() []SkyBlockNewsElement`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SkyBlockNews) GetItemsOk() (*[]SkyBlockNewsElement, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SkyBlockNews) SetItems(v []SkyBlockNewsElement)`

SetItems sets Items field to given value.


### GetSuccess

`func (o *SkyBlockNews) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SkyBlockNews) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SkyBlockNews) SetSuccess(v bool)`

SetSuccess sets Success field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


