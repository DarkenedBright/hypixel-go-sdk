# SkyBlockBazaar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastUpdated** | **int64** |  | 
**Products** | [**SkyBlockBazaarProducts**](SkyBlockBazaarProducts.md) |  | 
**Success** | **bool** |  | 

## Methods

### NewSkyBlockBazaar

`func NewSkyBlockBazaar(lastUpdated int64, products SkyBlockBazaarProducts, success bool, ) *SkyBlockBazaar`

NewSkyBlockBazaar instantiates a new SkyBlockBazaar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkyBlockBazaarWithDefaults

`func NewSkyBlockBazaarWithDefaults() *SkyBlockBazaar`

NewSkyBlockBazaarWithDefaults instantiates a new SkyBlockBazaar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastUpdated

`func (o *SkyBlockBazaar) GetLastUpdated() int64`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SkyBlockBazaar) GetLastUpdatedOk() (*int64, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SkyBlockBazaar) SetLastUpdated(v int64)`

SetLastUpdated sets LastUpdated field to given value.


### GetProducts

`func (o *SkyBlockBazaar) GetProducts() SkyBlockBazaarProducts`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *SkyBlockBazaar) GetProductsOk() (*SkyBlockBazaarProducts, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *SkyBlockBazaar) SetProducts(v SkyBlockBazaarProducts)`

SetProducts sets Products field to given value.


### GetSuccess

`func (o *SkyBlockBazaar) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SkyBlockBazaar) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SkyBlockBazaar) SetSuccess(v bool)`

SetSuccess sets Success field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


