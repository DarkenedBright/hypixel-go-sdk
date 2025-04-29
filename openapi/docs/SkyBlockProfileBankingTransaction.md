# SkyBlockProfileBankingTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**SkyBlockProfileBankingTransactionAction**](SkyBlockProfileBankingTransactionAction.md) |  | 
**Amount** | **float64** |  | 
**InitiatorName** | **string** |  | 
**Timestamp** | **int64** |  | 

## Methods

### NewSkyBlockProfileBankingTransaction

`func NewSkyBlockProfileBankingTransaction(action SkyBlockProfileBankingTransactionAction, amount float64, initiatorName string, timestamp int64, ) *SkyBlockProfileBankingTransaction`

NewSkyBlockProfileBankingTransaction instantiates a new SkyBlockProfileBankingTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkyBlockProfileBankingTransactionWithDefaults

`func NewSkyBlockProfileBankingTransactionWithDefaults() *SkyBlockProfileBankingTransaction`

NewSkyBlockProfileBankingTransactionWithDefaults instantiates a new SkyBlockProfileBankingTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *SkyBlockProfileBankingTransaction) GetAction() SkyBlockProfileBankingTransactionAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SkyBlockProfileBankingTransaction) GetActionOk() (*SkyBlockProfileBankingTransactionAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SkyBlockProfileBankingTransaction) SetAction(v SkyBlockProfileBankingTransactionAction)`

SetAction sets Action field to given value.


### GetAmount

`func (o *SkyBlockProfileBankingTransaction) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SkyBlockProfileBankingTransaction) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SkyBlockProfileBankingTransaction) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetInitiatorName

`func (o *SkyBlockProfileBankingTransaction) GetInitiatorName() string`

GetInitiatorName returns the InitiatorName field if non-nil, zero value otherwise.

### GetInitiatorNameOk

`func (o *SkyBlockProfileBankingTransaction) GetInitiatorNameOk() (*string, bool)`

GetInitiatorNameOk returns a tuple with the InitiatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorName

`func (o *SkyBlockProfileBankingTransaction) SetInitiatorName(v string)`

SetInitiatorName sets InitiatorName field to given value.


### GetTimestamp

`func (o *SkyBlockProfileBankingTransaction) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SkyBlockProfileBankingTransaction) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SkyBlockProfileBankingTransaction) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


