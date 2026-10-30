# AdminScimSyncActiveRunOUT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunId** | **string** | Unique identifier of the already active SCIM synchronization run. | 
**Status** | [**ScimSyncRunStatus**](ScimSyncRunStatus.md) | Current lifecycle status of the active synchronization run. | 

## Methods

### NewAdminScimSyncActiveRunOUT

`func NewAdminScimSyncActiveRunOUT(runId string, status ScimSyncRunStatus, ) *AdminScimSyncActiveRunOUT`

NewAdminScimSyncActiveRunOUT instantiates a new AdminScimSyncActiveRunOUT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminScimSyncActiveRunOUTWithDefaults

`func NewAdminScimSyncActiveRunOUTWithDefaults() *AdminScimSyncActiveRunOUT`

NewAdminScimSyncActiveRunOUTWithDefaults instantiates a new AdminScimSyncActiveRunOUT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunId

`func (o *AdminScimSyncActiveRunOUT) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *AdminScimSyncActiveRunOUT) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *AdminScimSyncActiveRunOUT) SetRunId(v string)`

SetRunId sets RunId field to given value.


### GetStatus

`func (o *AdminScimSyncActiveRunOUT) GetStatus() ScimSyncRunStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminScimSyncActiveRunOUT) GetStatusOk() (*ScimSyncRunStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminScimSyncActiveRunOUT) SetStatus(v ScimSyncRunStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


