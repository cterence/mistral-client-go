# BatchJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | Pointer to **string** |  | [optional] [default to "batch"]
**InputFiles** | **[]string** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Endpoint** | **string** |  | 
**Model** | Pointer to **NullableString** |  | [optional] 
**AgentId** | Pointer to **NullableString** |  | [optional] 
**OutputFile** | Pointer to **NullableString** |  | [optional] 
**ErrorFile** | Pointer to **NullableString** |  | [optional] 
**Errors** | [**[]BatchError**](BatchError.md) |  | 
**Outputs** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Status** | [**BatchJobStatus**](BatchJobStatus.md) |  | 
**CreatedAt** | **int32** |  | 
**TotalRequests** | **int32** |  | 
**CompletedRequests** | **int32** |  | 
**SucceededRequests** | **int32** |  | 
**FailedRequests** | **int32** |  | 
**StartedAt** | Pointer to **NullableInt32** |  | [optional] 
**CompletedAt** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewBatchJob

`func NewBatchJob(id string, inputFiles []string, endpoint string, errors []BatchError, status BatchJobStatus, createdAt int32, totalRequests int32, completedRequests int32, succeededRequests int32, failedRequests int32, ) *BatchJob`

NewBatchJob instantiates a new BatchJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchJobWithDefaults

`func NewBatchJobWithDefaults() *BatchJob`

NewBatchJobWithDefaults instantiates a new BatchJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BatchJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BatchJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BatchJob) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *BatchJob) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BatchJob) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BatchJob) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *BatchJob) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetInputFiles

`func (o *BatchJob) GetInputFiles() []string`

GetInputFiles returns the InputFiles field if non-nil, zero value otherwise.

### GetInputFilesOk

`func (o *BatchJob) GetInputFilesOk() (*[]string, bool)`

GetInputFilesOk returns a tuple with the InputFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFiles

`func (o *BatchJob) SetInputFiles(v []string)`

SetInputFiles sets InputFiles field to given value.


### GetMetadata

`func (o *BatchJob) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BatchJob) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BatchJob) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BatchJob) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *BatchJob) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *BatchJob) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetEndpoint

`func (o *BatchJob) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *BatchJob) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *BatchJob) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetModel

`func (o *BatchJob) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *BatchJob) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *BatchJob) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *BatchJob) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *BatchJob) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *BatchJob) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetAgentId

`func (o *BatchJob) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *BatchJob) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *BatchJob) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *BatchJob) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### SetAgentIdNil

`func (o *BatchJob) SetAgentIdNil(b bool)`

 SetAgentIdNil sets the value for AgentId to be an explicit nil

### UnsetAgentId
`func (o *BatchJob) UnsetAgentId()`

UnsetAgentId ensures that no value is present for AgentId, not even an explicit nil
### GetOutputFile

`func (o *BatchJob) GetOutputFile() string`

GetOutputFile returns the OutputFile field if non-nil, zero value otherwise.

### GetOutputFileOk

`func (o *BatchJob) GetOutputFileOk() (*string, bool)`

GetOutputFileOk returns a tuple with the OutputFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFile

`func (o *BatchJob) SetOutputFile(v string)`

SetOutputFile sets OutputFile field to given value.

### HasOutputFile

`func (o *BatchJob) HasOutputFile() bool`

HasOutputFile returns a boolean if a field has been set.

### SetOutputFileNil

`func (o *BatchJob) SetOutputFileNil(b bool)`

 SetOutputFileNil sets the value for OutputFile to be an explicit nil

### UnsetOutputFile
`func (o *BatchJob) UnsetOutputFile()`

UnsetOutputFile ensures that no value is present for OutputFile, not even an explicit nil
### GetErrorFile

`func (o *BatchJob) GetErrorFile() string`

GetErrorFile returns the ErrorFile field if non-nil, zero value otherwise.

### GetErrorFileOk

`func (o *BatchJob) GetErrorFileOk() (*string, bool)`

GetErrorFileOk returns a tuple with the ErrorFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorFile

`func (o *BatchJob) SetErrorFile(v string)`

SetErrorFile sets ErrorFile field to given value.

### HasErrorFile

`func (o *BatchJob) HasErrorFile() bool`

HasErrorFile returns a boolean if a field has been set.

### SetErrorFileNil

`func (o *BatchJob) SetErrorFileNil(b bool)`

 SetErrorFileNil sets the value for ErrorFile to be an explicit nil

### UnsetErrorFile
`func (o *BatchJob) UnsetErrorFile()`

UnsetErrorFile ensures that no value is present for ErrorFile, not even an explicit nil
### GetErrors

`func (o *BatchJob) GetErrors() []BatchError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BatchJob) GetErrorsOk() (*[]BatchError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BatchJob) SetErrors(v []BatchError)`

SetErrors sets Errors field to given value.


### GetOutputs

`func (o *BatchJob) GetOutputs() []map[string]interface{}`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *BatchJob) GetOutputsOk() (*[]map[string]interface{}, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *BatchJob) SetOutputs(v []map[string]interface{})`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *BatchJob) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### SetOutputsNil

`func (o *BatchJob) SetOutputsNil(b bool)`

 SetOutputsNil sets the value for Outputs to be an explicit nil

### UnsetOutputs
`func (o *BatchJob) UnsetOutputs()`

UnsetOutputs ensures that no value is present for Outputs, not even an explicit nil
### GetStatus

`func (o *BatchJob) GetStatus() BatchJobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchJob) GetStatusOk() (*BatchJobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchJob) SetStatus(v BatchJobStatus)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *BatchJob) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BatchJob) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BatchJob) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetTotalRequests

`func (o *BatchJob) GetTotalRequests() int32`

GetTotalRequests returns the TotalRequests field if non-nil, zero value otherwise.

### GetTotalRequestsOk

`func (o *BatchJob) GetTotalRequestsOk() (*int32, bool)`

GetTotalRequestsOk returns a tuple with the TotalRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRequests

`func (o *BatchJob) SetTotalRequests(v int32)`

SetTotalRequests sets TotalRequests field to given value.


### GetCompletedRequests

`func (o *BatchJob) GetCompletedRequests() int32`

GetCompletedRequests returns the CompletedRequests field if non-nil, zero value otherwise.

### GetCompletedRequestsOk

`func (o *BatchJob) GetCompletedRequestsOk() (*int32, bool)`

GetCompletedRequestsOk returns a tuple with the CompletedRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedRequests

`func (o *BatchJob) SetCompletedRequests(v int32)`

SetCompletedRequests sets CompletedRequests field to given value.


### GetSucceededRequests

`func (o *BatchJob) GetSucceededRequests() int32`

GetSucceededRequests returns the SucceededRequests field if non-nil, zero value otherwise.

### GetSucceededRequestsOk

`func (o *BatchJob) GetSucceededRequestsOk() (*int32, bool)`

GetSucceededRequestsOk returns a tuple with the SucceededRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceededRequests

`func (o *BatchJob) SetSucceededRequests(v int32)`

SetSucceededRequests sets SucceededRequests field to given value.


### GetFailedRequests

`func (o *BatchJob) GetFailedRequests() int32`

GetFailedRequests returns the FailedRequests field if non-nil, zero value otherwise.

### GetFailedRequestsOk

`func (o *BatchJob) GetFailedRequestsOk() (*int32, bool)`

GetFailedRequestsOk returns a tuple with the FailedRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRequests

`func (o *BatchJob) SetFailedRequests(v int32)`

SetFailedRequests sets FailedRequests field to given value.


### GetStartedAt

`func (o *BatchJob) GetStartedAt() int32`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BatchJob) GetStartedAtOk() (*int32, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BatchJob) SetStartedAt(v int32)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *BatchJob) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *BatchJob) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *BatchJob) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *BatchJob) GetCompletedAt() int32`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *BatchJob) GetCompletedAtOk() (*int32, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *BatchJob) SetCompletedAt(v int32)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *BatchJob) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *BatchJob) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *BatchJob) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


