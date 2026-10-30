# PublicExecutionEnv

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tools** | [**[]Tool**](Tool.md) |  | 
**ToolExecutionData** | [**PublicConnectorExecutionData**](PublicConnectorExecutionData.md) |  | 
**Errors** | **[]string** |  | 

## Methods

### NewPublicExecutionEnv

`func NewPublicExecutionEnv(tools []Tool, toolExecutionData PublicConnectorExecutionData, errors []string, ) *PublicExecutionEnv`

NewPublicExecutionEnv instantiates a new PublicExecutionEnv object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicExecutionEnvWithDefaults

`func NewPublicExecutionEnvWithDefaults() *PublicExecutionEnv`

NewPublicExecutionEnvWithDefaults instantiates a new PublicExecutionEnv object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTools

`func (o *PublicExecutionEnv) GetTools() []Tool`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *PublicExecutionEnv) GetToolsOk() (*[]Tool, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *PublicExecutionEnv) SetTools(v []Tool)`

SetTools sets Tools field to given value.


### GetToolExecutionData

`func (o *PublicExecutionEnv) GetToolExecutionData() PublicConnectorExecutionData`

GetToolExecutionData returns the ToolExecutionData field if non-nil, zero value otherwise.

### GetToolExecutionDataOk

`func (o *PublicExecutionEnv) GetToolExecutionDataOk() (*PublicConnectorExecutionData, bool)`

GetToolExecutionDataOk returns a tuple with the ToolExecutionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolExecutionData

`func (o *PublicExecutionEnv) SetToolExecutionData(v PublicConnectorExecutionData)`

SetToolExecutionData sets ToolExecutionData field to given value.


### GetErrors

`func (o *PublicExecutionEnv) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *PublicExecutionEnv) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *PublicExecutionEnv) SetErrors(v []string)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


