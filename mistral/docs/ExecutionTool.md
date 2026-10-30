# ExecutionTool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**IntegrationId** | **string** |  | 
**ExecutionConfig** | [**NullableExecutionConfig**](ExecutionConfig.md) |  | 

## Methods

### NewExecutionTool

`func NewExecutionTool(name string, integrationId string, executionConfig NullableExecutionConfig, ) *ExecutionTool`

NewExecutionTool instantiates a new ExecutionTool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionToolWithDefaults

`func NewExecutionToolWithDefaults() *ExecutionTool`

NewExecutionToolWithDefaults instantiates a new ExecutionTool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExecutionTool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExecutionTool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExecutionTool) SetName(v string)`

SetName sets Name field to given value.


### GetIntegrationId

`func (o *ExecutionTool) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *ExecutionTool) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *ExecutionTool) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetExecutionConfig

`func (o *ExecutionTool) GetExecutionConfig() ExecutionConfig`

GetExecutionConfig returns the ExecutionConfig field if non-nil, zero value otherwise.

### GetExecutionConfigOk

`func (o *ExecutionTool) GetExecutionConfigOk() (*ExecutionConfig, bool)`

GetExecutionConfigOk returns a tuple with the ExecutionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionConfig

`func (o *ExecutionTool) SetExecutionConfig(v ExecutionConfig)`

SetExecutionConfig sets ExecutionConfig field to given value.


### SetExecutionConfigNil

`func (o *ExecutionTool) SetExecutionConfigNil(b bool)`

 SetExecutionConfigNil sets the value for ExecutionConfig to be an explicit nil

### UnsetExecutionConfig
`func (o *ExecutionTool) UnsetExecutionConfig()`

UnsetExecutionConfig ensures that no value is present for ExecutionConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


