# ClientTasksRequestsCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sampling** | Pointer to [**NullableTasksSamplingCapability**](TasksSamplingCapability.md) |  | [optional] 
**Elicitation** | Pointer to [**NullableTasksElicitationCapability**](TasksElicitationCapability.md) |  | [optional] 

## Methods

### NewClientTasksRequestsCapability

`func NewClientTasksRequestsCapability() *ClientTasksRequestsCapability`

NewClientTasksRequestsCapability instantiates a new ClientTasksRequestsCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientTasksRequestsCapabilityWithDefaults

`func NewClientTasksRequestsCapabilityWithDefaults() *ClientTasksRequestsCapability`

NewClientTasksRequestsCapabilityWithDefaults instantiates a new ClientTasksRequestsCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSampling

`func (o *ClientTasksRequestsCapability) GetSampling() TasksSamplingCapability`

GetSampling returns the Sampling field if non-nil, zero value otherwise.

### GetSamplingOk

`func (o *ClientTasksRequestsCapability) GetSamplingOk() (*TasksSamplingCapability, bool)`

GetSamplingOk returns a tuple with the Sampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampling

`func (o *ClientTasksRequestsCapability) SetSampling(v TasksSamplingCapability)`

SetSampling sets Sampling field to given value.

### HasSampling

`func (o *ClientTasksRequestsCapability) HasSampling() bool`

HasSampling returns a boolean if a field has been set.

### SetSamplingNil

`func (o *ClientTasksRequestsCapability) SetSamplingNil(b bool)`

 SetSamplingNil sets the value for Sampling to be an explicit nil

### UnsetSampling
`func (o *ClientTasksRequestsCapability) UnsetSampling()`

UnsetSampling ensures that no value is present for Sampling, not even an explicit nil
### GetElicitation

`func (o *ClientTasksRequestsCapability) GetElicitation() TasksElicitationCapability`

GetElicitation returns the Elicitation field if non-nil, zero value otherwise.

### GetElicitationOk

`func (o *ClientTasksRequestsCapability) GetElicitationOk() (*TasksElicitationCapability, bool)`

GetElicitationOk returns a tuple with the Elicitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElicitation

`func (o *ClientTasksRequestsCapability) SetElicitation(v TasksElicitationCapability)`

SetElicitation sets Elicitation field to given value.

### HasElicitation

`func (o *ClientTasksRequestsCapability) HasElicitation() bool`

HasElicitation returns a boolean if a field has been set.

### SetElicitationNil

`func (o *ClientTasksRequestsCapability) SetElicitationNil(b bool)`

 SetElicitationNil sets the value for Elicitation to be an explicit nil

### UnsetElicitation
`func (o *ClientTasksRequestsCapability) UnsetElicitation()`

UnsetElicitation ensures that no value is present for Elicitation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


