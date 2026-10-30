# CreateBatchJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputFiles** | Pointer to **[]string** | A list of &#x60;.jsonl&#x60; files for batch inference. Each line must be a JSON object with a &#x60;body&#x60; field containing the request payload: &#x60;&#x60;&#x60;json {\&quot;custom_id\&quot;: \&quot;0\&quot;, \&quot;body\&quot;: {\&quot;max_tokens\&quot;: 100, \&quot;messages\&quot;: [{\&quot;role\&quot;: \&quot;user\&quot;, \&quot;content\&quot;: \&quot;What is the best French cheese?\&quot;}]}} {\&quot;custom_id\&quot;: \&quot;1\&quot;, \&quot;body\&quot;: {\&quot;max_tokens\&quot;: 100, \&quot;messages\&quot;: [{\&quot;role\&quot;: \&quot;user\&quot;, \&quot;content\&quot;: \&quot;What is the best French wine?\&quot;}]}} &#x60;&#x60;&#x60; | [optional] 
**Requests** | Pointer to [**[]BatchRequest**](BatchRequest.md) |  | [optional] 
**Endpoint** | [**ApiEndpoint**](ApiEndpoint.md) | The endpoint to be used for batch inference. | 
**Model** | Pointer to **NullableString** | The model to be used for batch inference. | [optional] 
**AgentId** | Pointer to **NullableString** | In case you want to use a specific agent from the **deprecated** agents api for batch inference, you can specify the agent ID here. | [optional] 
**Metadata** | Pointer to **map[string]string** | The metadata of your choice to be associated with the batch inference job. | [optional] 
**TimeoutHours** | Pointer to **int32** | The timeout in hours for the batch inference job. | [optional] [default to 24]

## Methods

### NewCreateBatchJobRequest

`func NewCreateBatchJobRequest(endpoint ApiEndpoint, ) *CreateBatchJobRequest`

NewCreateBatchJobRequest instantiates a new CreateBatchJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBatchJobRequestWithDefaults

`func NewCreateBatchJobRequestWithDefaults() *CreateBatchJobRequest`

NewCreateBatchJobRequestWithDefaults instantiates a new CreateBatchJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputFiles

`func (o *CreateBatchJobRequest) GetInputFiles() []string`

GetInputFiles returns the InputFiles field if non-nil, zero value otherwise.

### GetInputFilesOk

`func (o *CreateBatchJobRequest) GetInputFilesOk() (*[]string, bool)`

GetInputFilesOk returns a tuple with the InputFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFiles

`func (o *CreateBatchJobRequest) SetInputFiles(v []string)`

SetInputFiles sets InputFiles field to given value.

### HasInputFiles

`func (o *CreateBatchJobRequest) HasInputFiles() bool`

HasInputFiles returns a boolean if a field has been set.

### SetInputFilesNil

`func (o *CreateBatchJobRequest) SetInputFilesNil(b bool)`

 SetInputFilesNil sets the value for InputFiles to be an explicit nil

### UnsetInputFiles
`func (o *CreateBatchJobRequest) UnsetInputFiles()`

UnsetInputFiles ensures that no value is present for InputFiles, not even an explicit nil
### GetRequests

`func (o *CreateBatchJobRequest) GetRequests() []BatchRequest`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *CreateBatchJobRequest) GetRequestsOk() (*[]BatchRequest, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *CreateBatchJobRequest) SetRequests(v []BatchRequest)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *CreateBatchJobRequest) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### SetRequestsNil

`func (o *CreateBatchJobRequest) SetRequestsNil(b bool)`

 SetRequestsNil sets the value for Requests to be an explicit nil

### UnsetRequests
`func (o *CreateBatchJobRequest) UnsetRequests()`

UnsetRequests ensures that no value is present for Requests, not even an explicit nil
### GetEndpoint

`func (o *CreateBatchJobRequest) GetEndpoint() ApiEndpoint`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *CreateBatchJobRequest) GetEndpointOk() (*ApiEndpoint, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *CreateBatchJobRequest) SetEndpoint(v ApiEndpoint)`

SetEndpoint sets Endpoint field to given value.


### GetModel

`func (o *CreateBatchJobRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CreateBatchJobRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CreateBatchJobRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CreateBatchJobRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *CreateBatchJobRequest) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *CreateBatchJobRequest) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetAgentId

`func (o *CreateBatchJobRequest) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *CreateBatchJobRequest) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *CreateBatchJobRequest) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *CreateBatchJobRequest) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### SetAgentIdNil

`func (o *CreateBatchJobRequest) SetAgentIdNil(b bool)`

 SetAgentIdNil sets the value for AgentId to be an explicit nil

### UnsetAgentId
`func (o *CreateBatchJobRequest) UnsetAgentId()`

UnsetAgentId ensures that no value is present for AgentId, not even an explicit nil
### GetMetadata

`func (o *CreateBatchJobRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateBatchJobRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateBatchJobRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateBatchJobRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CreateBatchJobRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CreateBatchJobRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetTimeoutHours

`func (o *CreateBatchJobRequest) GetTimeoutHours() int32`

GetTimeoutHours returns the TimeoutHours field if non-nil, zero value otherwise.

### GetTimeoutHoursOk

`func (o *CreateBatchJobRequest) GetTimeoutHoursOk() (*int32, bool)`

GetTimeoutHoursOk returns a tuple with the TimeoutHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutHours

`func (o *CreateBatchJobRequest) SetTimeoutHours(v int32)`

SetTimeoutHours sets TimeoutHours field to given value.

### HasTimeoutHours

`func (o *CreateBatchJobRequest) HasTimeoutHours() bool`

HasTimeoutHours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


