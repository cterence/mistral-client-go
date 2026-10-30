# PublicConnectorExecutionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Integrations** | [**[]PublicExecutionConnector**](PublicExecutionConnector.md) |  | 
**Tools** | [**[]ExecutionTool**](ExecutionTool.md) |  | 

## Methods

### NewPublicConnectorExecutionData

`func NewPublicConnectorExecutionData(integrations []PublicExecutionConnector, tools []ExecutionTool, ) *PublicConnectorExecutionData`

NewPublicConnectorExecutionData instantiates a new PublicConnectorExecutionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicConnectorExecutionDataWithDefaults

`func NewPublicConnectorExecutionDataWithDefaults() *PublicConnectorExecutionData`

NewPublicConnectorExecutionDataWithDefaults instantiates a new PublicConnectorExecutionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrations

`func (o *PublicConnectorExecutionData) GetIntegrations() []PublicExecutionConnector`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *PublicConnectorExecutionData) GetIntegrationsOk() (*[]PublicExecutionConnector, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *PublicConnectorExecutionData) SetIntegrations(v []PublicExecutionConnector)`

SetIntegrations sets Integrations field to given value.


### GetTools

`func (o *PublicConnectorExecutionData) GetTools() []ExecutionTool`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *PublicConnectorExecutionData) GetToolsOk() (*[]ExecutionTool, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *PublicConnectorExecutionData) SetTools(v []ExecutionTool)`

SetTools sets Tools field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


