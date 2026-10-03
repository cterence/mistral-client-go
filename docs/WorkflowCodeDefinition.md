# WorkflowCodeDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputSchema** | **map[string]interface{}** | Input schema of the workflow&#39;s run method | 
**OutputSchema** | Pointer to **map[string]interface{}** | Output JSON schema of the query&#39;s model | [optional] 
**Signals** | Pointer to [**[]SignalDefinition**](SignalDefinition.md) | Signal handlers defined by the workflow | [optional] 
**Queries** | Pointer to [**[]QueryDefinition**](QueryDefinition.md) | Query handlers defined by the workflow | [optional] 
**Updates** | Pointer to [**[]UpdateDefinition**](UpdateDefinition.md) | Update handlers defined by the workflow | [optional] 
**EnforceDeterminism** | Pointer to **bool** | Whether the workflow enforces deterministic execution | [optional] [default to false]
**ExecutionTimeout** | Pointer to **float32** | Maximum total execution time including retries and continue-as-new | [optional] 

## Methods

### NewWorkflowCodeDefinition

`func NewWorkflowCodeDefinition(inputSchema map[string]interface{}, ) *WorkflowCodeDefinition`

NewWorkflowCodeDefinition instantiates a new WorkflowCodeDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowCodeDefinitionWithDefaults

`func NewWorkflowCodeDefinitionWithDefaults() *WorkflowCodeDefinition`

NewWorkflowCodeDefinitionWithDefaults instantiates a new WorkflowCodeDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputSchema

`func (o *WorkflowCodeDefinition) GetInputSchema() map[string]interface{}`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *WorkflowCodeDefinition) GetInputSchemaOk() (*map[string]interface{}, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *WorkflowCodeDefinition) SetInputSchema(v map[string]interface{})`

SetInputSchema sets InputSchema field to given value.


### GetOutputSchema

`func (o *WorkflowCodeDefinition) GetOutputSchema() map[string]interface{}`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *WorkflowCodeDefinition) GetOutputSchemaOk() (*map[string]interface{}, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *WorkflowCodeDefinition) SetOutputSchema(v map[string]interface{})`

SetOutputSchema sets OutputSchema field to given value.

### HasOutputSchema

`func (o *WorkflowCodeDefinition) HasOutputSchema() bool`

HasOutputSchema returns a boolean if a field has been set.

### SetOutputSchemaNil

`func (o *WorkflowCodeDefinition) SetOutputSchemaNil(b bool)`

 SetOutputSchemaNil sets the value for OutputSchema to be an explicit nil

### UnsetOutputSchema
`func (o *WorkflowCodeDefinition) UnsetOutputSchema()`

UnsetOutputSchema ensures that no value is present for OutputSchema, not even an explicit nil
### GetSignals

`func (o *WorkflowCodeDefinition) GetSignals() []SignalDefinition`

GetSignals returns the Signals field if non-nil, zero value otherwise.

### GetSignalsOk

`func (o *WorkflowCodeDefinition) GetSignalsOk() (*[]SignalDefinition, bool)`

GetSignalsOk returns a tuple with the Signals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignals

`func (o *WorkflowCodeDefinition) SetSignals(v []SignalDefinition)`

SetSignals sets Signals field to given value.

### HasSignals

`func (o *WorkflowCodeDefinition) HasSignals() bool`

HasSignals returns a boolean if a field has been set.

### GetQueries

`func (o *WorkflowCodeDefinition) GetQueries() []QueryDefinition`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *WorkflowCodeDefinition) GetQueriesOk() (*[]QueryDefinition, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *WorkflowCodeDefinition) SetQueries(v []QueryDefinition)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *WorkflowCodeDefinition) HasQueries() bool`

HasQueries returns a boolean if a field has been set.

### GetUpdates

`func (o *WorkflowCodeDefinition) GetUpdates() []UpdateDefinition`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *WorkflowCodeDefinition) GetUpdatesOk() (*[]UpdateDefinition, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *WorkflowCodeDefinition) SetUpdates(v []UpdateDefinition)`

SetUpdates sets Updates field to given value.

### HasUpdates

`func (o *WorkflowCodeDefinition) HasUpdates() bool`

HasUpdates returns a boolean if a field has been set.

### GetEnforceDeterminism

`func (o *WorkflowCodeDefinition) GetEnforceDeterminism() bool`

GetEnforceDeterminism returns the EnforceDeterminism field if non-nil, zero value otherwise.

### GetEnforceDeterminismOk

`func (o *WorkflowCodeDefinition) GetEnforceDeterminismOk() (*bool, bool)`

GetEnforceDeterminismOk returns a tuple with the EnforceDeterminism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceDeterminism

`func (o *WorkflowCodeDefinition) SetEnforceDeterminism(v bool)`

SetEnforceDeterminism sets EnforceDeterminism field to given value.

### HasEnforceDeterminism

`func (o *WorkflowCodeDefinition) HasEnforceDeterminism() bool`

HasEnforceDeterminism returns a boolean if a field has been set.

### GetExecutionTimeout

`func (o *WorkflowCodeDefinition) GetExecutionTimeout() float32`

GetExecutionTimeout returns the ExecutionTimeout field if non-nil, zero value otherwise.

### GetExecutionTimeoutOk

`func (o *WorkflowCodeDefinition) GetExecutionTimeoutOk() (*float32, bool)`

GetExecutionTimeoutOk returns a tuple with the ExecutionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionTimeout

`func (o *WorkflowCodeDefinition) SetExecutionTimeout(v float32)`

SetExecutionTimeout sets ExecutionTimeout field to given value.

### HasExecutionTimeout

`func (o *WorkflowCodeDefinition) HasExecutionTimeout() bool`

HasExecutionTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


