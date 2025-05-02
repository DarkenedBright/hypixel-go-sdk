# SkyBlockMuseumMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Appraisal** | **bool** |  | 
**Items** | [**SkyBlockMuseumMemberItems**](SkyBlockMuseumMemberItems.md) |  | 
**Special** | Pointer to [**[]SkyBlockMuseumMemberItem**](SkyBlockMuseumMemberItem.md) |  | [optional] 
**Value** | **int64** |  | 

## Methods

### NewSkyBlockMuseumMember

`func NewSkyBlockMuseumMember(appraisal bool, items SkyBlockMuseumMemberItems, value int64, ) *SkyBlockMuseumMember`

NewSkyBlockMuseumMember instantiates a new SkyBlockMuseumMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkyBlockMuseumMemberWithDefaults

`func NewSkyBlockMuseumMemberWithDefaults() *SkyBlockMuseumMember`

NewSkyBlockMuseumMemberWithDefaults instantiates a new SkyBlockMuseumMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppraisal

`func (o *SkyBlockMuseumMember) GetAppraisal() bool`

GetAppraisal returns the Appraisal field if non-nil, zero value otherwise.

### GetAppraisalOk

`func (o *SkyBlockMuseumMember) GetAppraisalOk() (*bool, bool)`

GetAppraisalOk returns a tuple with the Appraisal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppraisal

`func (o *SkyBlockMuseumMember) SetAppraisal(v bool)`

SetAppraisal sets Appraisal field to given value.


### GetItems

`func (o *SkyBlockMuseumMember) GetItems() SkyBlockMuseumMemberItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SkyBlockMuseumMember) GetItemsOk() (*SkyBlockMuseumMemberItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SkyBlockMuseumMember) SetItems(v SkyBlockMuseumMemberItems)`

SetItems sets Items field to given value.


### GetSpecial

`func (o *SkyBlockMuseumMember) GetSpecial() []SkyBlockMuseumMemberItem`

GetSpecial returns the Special field if non-nil, zero value otherwise.

### GetSpecialOk

`func (o *SkyBlockMuseumMember) GetSpecialOk() (*[]SkyBlockMuseumMemberItem, bool)`

GetSpecialOk returns a tuple with the Special field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecial

`func (o *SkyBlockMuseumMember) SetSpecial(v []SkyBlockMuseumMemberItem)`

SetSpecial sets Special field to given value.

### HasSpecial

`func (o *SkyBlockMuseumMember) HasSpecial() bool`

HasSpecial returns a boolean if a field has been set.

### GetValue

`func (o *SkyBlockMuseumMember) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SkyBlockMuseumMember) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SkyBlockMuseumMember) SetValue(v int64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


