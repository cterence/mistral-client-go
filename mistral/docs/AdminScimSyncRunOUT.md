# AdminScimSyncRunOUT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunId** | **string** | Unique identifier of the SCIM synchronization run. | 
**Status** | [**ScimSyncRunStatus**](ScimSyncRunStatus.md) | Current lifecycle status of the synchronization run. | 
**DryRun** | **bool** | Whether the run only previewed changes. | 
**SyncConfig** | [**NullableAdminScimSyncConfig**](AdminScimSyncConfig.md) | Categories selected for this synchronization run. | 
**CreatedAt** | **time.Time** | Time at which the synchronization run was created. | 
**StartedAt** | **NullableTime** | Only return logs at or after this timestamp | 
**FinishedAt** | **NullableTime** | Only return logs at or after this timestamp | 
**Summary** | [**NullableScimSyncSummaryOut**](ScimSyncSummaryOut.md) | Preview or result summary, available after the synchronization plan is built. | 
**ErrorMessage** | **NullableString** | Filter by agent name | 

## Methods

### NewAdminScimSyncRunOUT

`func NewAdminScimSyncRunOUT(runId string, status ScimSyncRunStatus, dryRun bool, syncConfig NullableAdminScimSyncConfig, createdAt time.Time, startedAt NullableTime, finishedAt NullableTime, summary NullableScimSyncSummaryOut, errorMessage NullableString, ) *AdminScimSyncRunOUT`

NewAdminScimSyncRunOUT instantiates a new AdminScimSyncRunOUT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminScimSyncRunOUTWithDefaults

`func NewAdminScimSyncRunOUTWithDefaults() *AdminScimSyncRunOUT`

NewAdminScimSyncRunOUTWithDefaults instantiates a new AdminScimSyncRunOUT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunId

`func (o *AdminScimSyncRunOUT) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *AdminScimSyncRunOUT) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *AdminScimSyncRunOUT) SetRunId(v string)`

SetRunId sets RunId field to given value.


### GetStatus

`func (o *AdminScimSyncRunOUT) GetStatus() ScimSyncRunStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminScimSyncRunOUT) GetStatusOk() (*ScimSyncRunStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminScimSyncRunOUT) SetStatus(v ScimSyncRunStatus)`

SetStatus sets Status field to given value.


### GetDryRun

`func (o *AdminScimSyncRunOUT) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *AdminScimSyncRunOUT) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *AdminScimSyncRunOUT) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.


### GetSyncConfig

`func (o *AdminScimSyncRunOUT) GetSyncConfig() AdminScimSyncConfig`

GetSyncConfig returns the SyncConfig field if non-nil, zero value otherwise.

### GetSyncConfigOk

`func (o *AdminScimSyncRunOUT) GetSyncConfigOk() (*AdminScimSyncConfig, bool)`

GetSyncConfigOk returns a tuple with the SyncConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncConfig

`func (o *AdminScimSyncRunOUT) SetSyncConfig(v AdminScimSyncConfig)`

SetSyncConfig sets SyncConfig field to given value.


### SetSyncConfigNil

`func (o *AdminScimSyncRunOUT) SetSyncConfigNil(b bool)`

 SetSyncConfigNil sets the value for SyncConfig to be an explicit nil

### UnsetSyncConfig
`func (o *AdminScimSyncRunOUT) UnsetSyncConfig()`

UnsetSyncConfig ensures that no value is present for SyncConfig, not even an explicit nil
### GetCreatedAt

`func (o *AdminScimSyncRunOUT) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdminScimSyncRunOUT) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdminScimSyncRunOUT) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetStartedAt

`func (o *AdminScimSyncRunOUT) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *AdminScimSyncRunOUT) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *AdminScimSyncRunOUT) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### SetStartedAtNil

`func (o *AdminScimSyncRunOUT) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *AdminScimSyncRunOUT) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetFinishedAt

`func (o *AdminScimSyncRunOUT) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *AdminScimSyncRunOUT) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *AdminScimSyncRunOUT) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.


### SetFinishedAtNil

`func (o *AdminScimSyncRunOUT) SetFinishedAtNil(b bool)`

 SetFinishedAtNil sets the value for FinishedAt to be an explicit nil

### UnsetFinishedAt
`func (o *AdminScimSyncRunOUT) UnsetFinishedAt()`

UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil
### GetSummary

`func (o *AdminScimSyncRunOUT) GetSummary() ScimSyncSummaryOut`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdminScimSyncRunOUT) GetSummaryOk() (*ScimSyncSummaryOut, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdminScimSyncRunOUT) SetSummary(v ScimSyncSummaryOut)`

SetSummary sets Summary field to given value.


### SetSummaryNil

`func (o *AdminScimSyncRunOUT) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *AdminScimSyncRunOUT) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetErrorMessage

`func (o *AdminScimSyncRunOUT) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AdminScimSyncRunOUT) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AdminScimSyncRunOUT) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### SetErrorMessageNil

`func (o *AdminScimSyncRunOUT) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *AdminScimSyncRunOUT) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


