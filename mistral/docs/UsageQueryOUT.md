# UsageQueryOUT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | [**[]UsageQueryRowOUT**](UsageQueryRowOUT.md) | Grouped usage rows. | 
**Aggregates** | Pointer to [**NullableUsageQueryAggregatesOUT**](UsageQueryAggregatesOUT.md) | Aggregates over the complete filtered result set when requested. | [optional] 
**Metadata** | [**UsageQueryMetadataOUT**](UsageQueryMetadataOUT.md) | Query parameters and pagination metadata. | 

## Methods

### NewUsageQueryOUT

`func NewUsageQueryOUT(groups []UsageQueryRowOUT, metadata UsageQueryMetadataOUT, ) *UsageQueryOUT`

NewUsageQueryOUT instantiates a new UsageQueryOUT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageQueryOUTWithDefaults

`func NewUsageQueryOUTWithDefaults() *UsageQueryOUT`

NewUsageQueryOUTWithDefaults instantiates a new UsageQueryOUT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *UsageQueryOUT) GetGroups() []UsageQueryRowOUT`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UsageQueryOUT) GetGroupsOk() (*[]UsageQueryRowOUT, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UsageQueryOUT) SetGroups(v []UsageQueryRowOUT)`

SetGroups sets Groups field to given value.


### GetAggregates

`func (o *UsageQueryOUT) GetAggregates() UsageQueryAggregatesOUT`

GetAggregates returns the Aggregates field if non-nil, zero value otherwise.

### GetAggregatesOk

`func (o *UsageQueryOUT) GetAggregatesOk() (*UsageQueryAggregatesOUT, bool)`

GetAggregatesOk returns a tuple with the Aggregates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregates

`func (o *UsageQueryOUT) SetAggregates(v UsageQueryAggregatesOUT)`

SetAggregates sets Aggregates field to given value.

### HasAggregates

`func (o *UsageQueryOUT) HasAggregates() bool`

HasAggregates returns a boolean if a field has been set.

### SetAggregatesNil

`func (o *UsageQueryOUT) SetAggregatesNil(b bool)`

 SetAggregatesNil sets the value for Aggregates to be an explicit nil

### UnsetAggregates
`func (o *UsageQueryOUT) UnsetAggregates()`

UnsetAggregates ensures that no value is present for Aggregates, not even an explicit nil
### GetMetadata

`func (o *UsageQueryOUT) GetMetadata() UsageQueryMetadataOUT`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UsageQueryOUT) GetMetadataOk() (*UsageQueryMetadataOUT, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UsageQueryOUT) SetMetadata(v UsageQueryMetadataOUT)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


