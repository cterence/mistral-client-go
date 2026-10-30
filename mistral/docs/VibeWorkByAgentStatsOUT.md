# VibeWorkByAgentStatsOUT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | **int32** | Start of the queried window, as a Unix timestamp in seconds. | 
**EndTime** | **int32** | End of the queried window, as a Unix timestamp in seconds. | 
**Data** | [**[]VibeWorkByAgentStat**](VibeWorkByAgentStat.md) |  | 

## Methods

### NewVibeWorkByAgentStatsOUT

`func NewVibeWorkByAgentStatsOUT(startTime int32, endTime int32, data []VibeWorkByAgentStat, ) *VibeWorkByAgentStatsOUT`

NewVibeWorkByAgentStatsOUT instantiates a new VibeWorkByAgentStatsOUT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVibeWorkByAgentStatsOUTWithDefaults

`func NewVibeWorkByAgentStatsOUTWithDefaults() *VibeWorkByAgentStatsOUT`

NewVibeWorkByAgentStatsOUTWithDefaults instantiates a new VibeWorkByAgentStatsOUT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *VibeWorkByAgentStatsOUT) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *VibeWorkByAgentStatsOUT) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *VibeWorkByAgentStatsOUT) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *VibeWorkByAgentStatsOUT) GetEndTime() int32`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *VibeWorkByAgentStatsOUT) GetEndTimeOk() (*int32, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *VibeWorkByAgentStatsOUT) SetEndTime(v int32)`

SetEndTime sets EndTime field to given value.


### GetData

`func (o *VibeWorkByAgentStatsOUT) GetData() []VibeWorkByAgentStat`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VibeWorkByAgentStatsOUT) GetDataOk() (*[]VibeWorkByAgentStat, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VibeWorkByAgentStatsOUT) SetData(v []VibeWorkByAgentStat)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


