# SkyBlockMuseumMemberItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Borrowing** | Pointer to **bool** |  | [optional] 
**DonatedTime** | **int64** |  | 
**FeaturedSlot** | Pointer to [**SkyBlockMuseumMemberItemFeaturedSlot**](SkyBlockMuseumMemberItemFeaturedSlot.md) |  | [optional] 
**Items** | [**SkyBlockProfileMemberInventoryBase64GzipData**](SkyBlockProfileMemberInventoryBase64GzipData.md) |  | 

## Methods

### NewSkyBlockMuseumMemberItem

`func NewSkyBlockMuseumMemberItem(donatedTime int64, items SkyBlockProfileMemberInventoryBase64GzipData, ) *SkyBlockMuseumMemberItem`

NewSkyBlockMuseumMemberItem instantiates a new SkyBlockMuseumMemberItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkyBlockMuseumMemberItemWithDefaults

`func NewSkyBlockMuseumMemberItemWithDefaults() *SkyBlockMuseumMemberItem`

NewSkyBlockMuseumMemberItemWithDefaults instantiates a new SkyBlockMuseumMemberItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBorrowing

`func (o *SkyBlockMuseumMemberItem) GetBorrowing() bool`

GetBorrowing returns the Borrowing field if non-nil, zero value otherwise.

### GetBorrowingOk

`func (o *SkyBlockMuseumMemberItem) GetBorrowingOk() (*bool, bool)`

GetBorrowingOk returns a tuple with the Borrowing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorrowing

`func (o *SkyBlockMuseumMemberItem) SetBorrowing(v bool)`

SetBorrowing sets Borrowing field to given value.

### HasBorrowing

`func (o *SkyBlockMuseumMemberItem) HasBorrowing() bool`

HasBorrowing returns a boolean if a field has been set.

### GetDonatedTime

`func (o *SkyBlockMuseumMemberItem) GetDonatedTime() int64`

GetDonatedTime returns the DonatedTime field if non-nil, zero value otherwise.

### GetDonatedTimeOk

`func (o *SkyBlockMuseumMemberItem) GetDonatedTimeOk() (*int64, bool)`

GetDonatedTimeOk returns a tuple with the DonatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDonatedTime

`func (o *SkyBlockMuseumMemberItem) SetDonatedTime(v int64)`

SetDonatedTime sets DonatedTime field to given value.


### GetFeaturedSlot

`func (o *SkyBlockMuseumMemberItem) GetFeaturedSlot() SkyBlockMuseumMemberItemFeaturedSlot`

GetFeaturedSlot returns the FeaturedSlot field if non-nil, zero value otherwise.

### GetFeaturedSlotOk

`func (o *SkyBlockMuseumMemberItem) GetFeaturedSlotOk() (*SkyBlockMuseumMemberItemFeaturedSlot, bool)`

GetFeaturedSlotOk returns a tuple with the FeaturedSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturedSlot

`func (o *SkyBlockMuseumMemberItem) SetFeaturedSlot(v SkyBlockMuseumMemberItemFeaturedSlot)`

SetFeaturedSlot sets FeaturedSlot field to given value.

### HasFeaturedSlot

`func (o *SkyBlockMuseumMemberItem) HasFeaturedSlot() bool`

HasFeaturedSlot returns a boolean if a field has been set.

### GetItems

`func (o *SkyBlockMuseumMemberItem) GetItems() SkyBlockProfileMemberInventoryBase64GzipData`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SkyBlockMuseumMemberItem) GetItemsOk() (*SkyBlockProfileMemberInventoryBase64GzipData, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SkyBlockMuseumMemberItem) SetItems(v SkyBlockProfileMemberInventoryBase64GzipData)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


