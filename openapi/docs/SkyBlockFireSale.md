# SkyBlockFireSale

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | The amount of items available for this sale | [optional] 
**End** | Pointer to **int64** | The end time in unix milliseconds for the sale | [optional] 
**ItemId** | Pointer to **string** | The SkyBlock item ID for this sale | [optional] 
**Price** | Pointer to **int64** | The price in Gems for this sale | [optional] 
**Start** | Pointer to **int64** | The start time in unix milliseconds for the sale | [optional] 

## Methods

### NewSkyBlockFireSale

`func NewSkyBlockFireSale() *SkyBlockFireSale`

NewSkyBlockFireSale instantiates a new SkyBlockFireSale object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkyBlockFireSaleWithDefaults

`func NewSkyBlockFireSaleWithDefaults() *SkyBlockFireSale`

NewSkyBlockFireSaleWithDefaults instantiates a new SkyBlockFireSale object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *SkyBlockFireSale) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SkyBlockFireSale) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SkyBlockFireSale) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *SkyBlockFireSale) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetEnd

`func (o *SkyBlockFireSale) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *SkyBlockFireSale) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *SkyBlockFireSale) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *SkyBlockFireSale) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetItemId

`func (o *SkyBlockFireSale) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *SkyBlockFireSale) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *SkyBlockFireSale) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *SkyBlockFireSale) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetPrice

`func (o *SkyBlockFireSale) GetPrice() int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SkyBlockFireSale) GetPriceOk() (*int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SkyBlockFireSale) SetPrice(v int64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *SkyBlockFireSale) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetStart

`func (o *SkyBlockFireSale) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *SkyBlockFireSale) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *SkyBlockFireSale) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *SkyBlockFireSale) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


