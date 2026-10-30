# VibeOrganizationStatsOUT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | **int32** | Start of the queried window, as a Unix timestamp in seconds. | 
**EndTime** | **int32** | End of the queried window, as a Unix timestamp in seconds. | 
**NextEditSuggestions** | [**[]VibeNextEditSuggestionStat**](VibeNextEditSuggestionStat.md) |  | 
**NextEditActiveUsers** | [**[]VibeNextEditActiveUsersStat**](VibeNextEditActiveUsersStat.md) |  | 
**NextEditModifiedLoc** | [**[]VibeNextEditModifiedLOCStat**](VibeNextEditModifiedLOCStat.md) |  | 

## Methods

### NewVibeOrganizationStatsOUT

`func NewVibeOrganizationStatsOUT(startTime int32, endTime int32, nextEditSuggestions []VibeNextEditSuggestionStat, nextEditActiveUsers []VibeNextEditActiveUsersStat, nextEditModifiedLoc []VibeNextEditModifiedLOCStat, ) *VibeOrganizationStatsOUT`

NewVibeOrganizationStatsOUT instantiates a new VibeOrganizationStatsOUT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVibeOrganizationStatsOUTWithDefaults

`func NewVibeOrganizationStatsOUTWithDefaults() *VibeOrganizationStatsOUT`

NewVibeOrganizationStatsOUTWithDefaults instantiates a new VibeOrganizationStatsOUT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *VibeOrganizationStatsOUT) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *VibeOrganizationStatsOUT) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *VibeOrganizationStatsOUT) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *VibeOrganizationStatsOUT) GetEndTime() int32`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *VibeOrganizationStatsOUT) GetEndTimeOk() (*int32, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *VibeOrganizationStatsOUT) SetEndTime(v int32)`

SetEndTime sets EndTime field to given value.


### GetNextEditSuggestions

`func (o *VibeOrganizationStatsOUT) GetNextEditSuggestions() []VibeNextEditSuggestionStat`

GetNextEditSuggestions returns the NextEditSuggestions field if non-nil, zero value otherwise.

### GetNextEditSuggestionsOk

`func (o *VibeOrganizationStatsOUT) GetNextEditSuggestionsOk() (*[]VibeNextEditSuggestionStat, bool)`

GetNextEditSuggestionsOk returns a tuple with the NextEditSuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextEditSuggestions

`func (o *VibeOrganizationStatsOUT) SetNextEditSuggestions(v []VibeNextEditSuggestionStat)`

SetNextEditSuggestions sets NextEditSuggestions field to given value.


### GetNextEditActiveUsers

`func (o *VibeOrganizationStatsOUT) GetNextEditActiveUsers() []VibeNextEditActiveUsersStat`

GetNextEditActiveUsers returns the NextEditActiveUsers field if non-nil, zero value otherwise.

### GetNextEditActiveUsersOk

`func (o *VibeOrganizationStatsOUT) GetNextEditActiveUsersOk() (*[]VibeNextEditActiveUsersStat, bool)`

GetNextEditActiveUsersOk returns a tuple with the NextEditActiveUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextEditActiveUsers

`func (o *VibeOrganizationStatsOUT) SetNextEditActiveUsers(v []VibeNextEditActiveUsersStat)`

SetNextEditActiveUsers sets NextEditActiveUsers field to given value.


### GetNextEditModifiedLoc

`func (o *VibeOrganizationStatsOUT) GetNextEditModifiedLoc() []VibeNextEditModifiedLOCStat`

GetNextEditModifiedLoc returns the NextEditModifiedLoc field if non-nil, zero value otherwise.

### GetNextEditModifiedLocOk

`func (o *VibeOrganizationStatsOUT) GetNextEditModifiedLocOk() (*[]VibeNextEditModifiedLOCStat, bool)`

GetNextEditModifiedLocOk returns a tuple with the NextEditModifiedLoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextEditModifiedLoc

`func (o *VibeOrganizationStatsOUT) SetNextEditModifiedLoc(v []VibeNextEditModifiedLOCStat)`

SetNextEditModifiedLoc sets NextEditModifiedLoc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


