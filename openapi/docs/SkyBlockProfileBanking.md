# SkyBlockProfileBanking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | **float64** |  | 
**Transactions** | [**[]SkyBlockProfileBankingTransaction**](SkyBlockProfileBankingTransaction.md) |  | 

## Methods

### NewSkyBlockProfileBanking

`func NewSkyBlockProfileBanking(balance float64, transactions []SkyBlockProfileBankingTransaction, ) *SkyBlockProfileBanking`

NewSkyBlockProfileBanking instantiates a new SkyBlockProfileBanking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkyBlockProfileBankingWithDefaults

`func NewSkyBlockProfileBankingWithDefaults() *SkyBlockProfileBanking`

NewSkyBlockProfileBankingWithDefaults instantiates a new SkyBlockProfileBanking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *SkyBlockProfileBanking) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *SkyBlockProfileBanking) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *SkyBlockProfileBanking) SetBalance(v float64)`

SetBalance sets Balance field to given value.


### GetTransactions

`func (o *SkyBlockProfileBanking) GetTransactions() []SkyBlockProfileBankingTransaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *SkyBlockProfileBanking) GetTransactionsOk() (*[]SkyBlockProfileBankingTransaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *SkyBlockProfileBanking) SetTransactions(v []SkyBlockProfileBankingTransaction)`

SetTransactions sets Transactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


