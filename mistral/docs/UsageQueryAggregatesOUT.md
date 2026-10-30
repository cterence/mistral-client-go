# UsageQueryAggregatesOUT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalValue** | **float32** | Total usage value across all matching rows. | 
**TotalCount** | **int32** | Total number of matching usage records. | 
**TotalApiCalls** | **int32** | Total number of distinct API calls. | 
**CustomerCount** | **int32** | Number of matching customers. | 
**UsageTypeCount** | **int32** | Number of matching usage types. | 

## Methods

### NewUsageQueryAggregatesOUT

`func NewUsageQueryAggregatesOUT(totalValue float32, totalCount int32, totalApiCalls int32, customerCount int32, usageTypeCount int32, ) *UsageQueryAggregatesOUT`

NewUsageQueryAggregatesOUT instantiates a new UsageQueryAggregatesOUT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageQueryAggregatesOUTWithDefaults

`func NewUsageQueryAggregatesOUTWithDefaults() *UsageQueryAggregatesOUT`

NewUsageQueryAggregatesOUTWithDefaults instantiates a new UsageQueryAggregatesOUT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalValue

`func (o *UsageQueryAggregatesOUT) GetTotalValue() float32`

GetTotalValue returns the TotalValue field if non-nil, zero value otherwise.

### GetTotalValueOk

`func (o *UsageQueryAggregatesOUT) GetTotalValueOk() (*float32, bool)`

GetTotalValueOk returns a tuple with the TotalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalValue

`func (o *UsageQueryAggregatesOUT) SetTotalValue(v float32)`

SetTotalValue sets TotalValue field to given value.


### GetTotalCount

`func (o *UsageQueryAggregatesOUT) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *UsageQueryAggregatesOUT) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *UsageQueryAggregatesOUT) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetTotalApiCalls

`func (o *UsageQueryAggregatesOUT) GetTotalApiCalls() int32`

GetTotalApiCalls returns the TotalApiCalls field if non-nil, zero value otherwise.

### GetTotalApiCallsOk

`func (o *UsageQueryAggregatesOUT) GetTotalApiCallsOk() (*int32, bool)`

GetTotalApiCallsOk returns a tuple with the TotalApiCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalApiCalls

`func (o *UsageQueryAggregatesOUT) SetTotalApiCalls(v int32)`

SetTotalApiCalls sets TotalApiCalls field to given value.


### GetCustomerCount

`func (o *UsageQueryAggregatesOUT) GetCustomerCount() int32`

GetCustomerCount returns the CustomerCount field if non-nil, zero value otherwise.

### GetCustomerCountOk

`func (o *UsageQueryAggregatesOUT) GetCustomerCountOk() (*int32, bool)`

GetCustomerCountOk returns a tuple with the CustomerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCount

`func (o *UsageQueryAggregatesOUT) SetCustomerCount(v int32)`

SetCustomerCount sets CustomerCount field to given value.


### GetUsageTypeCount

`func (o *UsageQueryAggregatesOUT) GetUsageTypeCount() int32`

GetUsageTypeCount returns the UsageTypeCount field if non-nil, zero value otherwise.

### GetUsageTypeCountOk

`func (o *UsageQueryAggregatesOUT) GetUsageTypeCountOk() (*int32, bool)`

GetUsageTypeCountOk returns a tuple with the UsageTypeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageTypeCount

`func (o *UsageQueryAggregatesOUT) SetUsageTypeCount(v int32)`

SetUsageTypeCount sets UsageTypeCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


