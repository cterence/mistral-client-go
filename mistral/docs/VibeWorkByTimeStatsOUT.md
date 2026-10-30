# VibeWorkByTimeStatsOUT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | **int32** | Start of the queried window, as a Unix timestamp in seconds. | 
**EndTime** | **int32** | End of the queried window, as a Unix timestamp in seconds. | 
**Data** | [**[]VibeWorkByTimeStat**](VibeWorkByTimeStat.md) |  | 

## Methods

### NewVibeWorkByTimeStatsOUT

`func NewVibeWorkByTimeStatsOUT(startTime int32, endTime int32, data []VibeWorkByTimeStat, ) *VibeWorkByTimeStatsOUT`

NewVibeWorkByTimeStatsOUT instantiates a new VibeWorkByTimeStatsOUT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVibeWorkByTimeStatsOUTWithDefaults

`func NewVibeWorkByTimeStatsOUTWithDefaults() *VibeWorkByTimeStatsOUT`

NewVibeWorkByTimeStatsOUTWithDefaults instantiates a new VibeWorkByTimeStatsOUT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *VibeWorkByTimeStatsOUT) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *VibeWorkByTimeStatsOUT) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *VibeWorkByTimeStatsOUT) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *VibeWorkByTimeStatsOUT) GetEndTime() int32`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *VibeWorkByTimeStatsOUT) GetEndTimeOk() (*int32, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *VibeWorkByTimeStatsOUT) SetEndTime(v int32)`

SetEndTime sets EndTime field to given value.


### GetData

`func (o *VibeWorkByTimeStatsOUT) GetData() []VibeWorkByTimeStat`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VibeWorkByTimeStatsOUT) GetDataOk() (*[]VibeWorkByTimeStat, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VibeWorkByTimeStatsOUT) SetData(v []VibeWorkByTimeStat)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


