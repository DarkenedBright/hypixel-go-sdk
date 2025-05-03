# SkyBlockAuctionsEnded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auctions** | [**[]SkyBlockAuctionsEndedAuction**](SkyBlockAuctionsEndedAuction.md) |  | 
**LastUpdated** | **int64** |  | 
**Success** | **bool** |  | 

## Methods

### NewSkyBlockAuctionsEnded

`func NewSkyBlockAuctionsEnded(auctions []SkyBlockAuctionsEndedAuction, lastUpdated int64, success bool, ) *SkyBlockAuctionsEnded`

NewSkyBlockAuctionsEnded instantiates a new SkyBlockAuctionsEnded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkyBlockAuctionsEndedWithDefaults

`func NewSkyBlockAuctionsEndedWithDefaults() *SkyBlockAuctionsEnded`

NewSkyBlockAuctionsEndedWithDefaults instantiates a new SkyBlockAuctionsEnded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctions

`func (o *SkyBlockAuctionsEnded) GetAuctions() []SkyBlockAuctionsEndedAuction`

GetAuctions returns the Auctions field if non-nil, zero value otherwise.

### GetAuctionsOk

`func (o *SkyBlockAuctionsEnded) GetAuctionsOk() (*[]SkyBlockAuctionsEndedAuction, bool)`

GetAuctionsOk returns a tuple with the Auctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctions

`func (o *SkyBlockAuctionsEnded) SetAuctions(v []SkyBlockAuctionsEndedAuction)`

SetAuctions sets Auctions field to given value.


### GetLastUpdated

`func (o *SkyBlockAuctionsEnded) GetLastUpdated() int64`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SkyBlockAuctionsEnded) GetLastUpdatedOk() (*int64, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SkyBlockAuctionsEnded) SetLastUpdated(v int64)`

SetLastUpdated sets LastUpdated field to given value.


### GetSuccess

`func (o *SkyBlockAuctionsEnded) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SkyBlockAuctionsEnded) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SkyBlockAuctionsEnded) SetSuccess(v bool)`

SetSuccess sets Success field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


