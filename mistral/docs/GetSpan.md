# GetSpan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** |  | 
**OrganizationId** | **string** |  | 
**WorkspaceId** | **string** |  | 
**UserId** | **string** |  | 
**TraceId** | **string** |  | 
**SpanId** | **string** |  | 
**ParentSpanId** | **string** |  | 
**TraceState** | **string** |  | 
**StartTime** | **time.Time** |  | 
**EndTime** | **time.Time** |  | 
**DurationNs** | **int32** |  | 
**SpanName** | **string** |  | 
**SpanKind** | **string** |  | 
**ServiceName** | **string** |  | 
**StatusCode** | **string** |  | 
**StatusMessage** | **string** |  | 
**ErrorType** | **string** |  | 
**OperationName** | **string** |  | 
**ProviderName** | **string** |  | 
**RequestModel** | **string** |  | 
**ResponseModel** | **string** |  | 
**ResponseId** | **string** |  | 
**OutputType** | **string** |  | 
**ConversationId** | **string** |  | 
**DataSourceId** | **string** |  | 
**AgentId** | **string** |  | 
**AgentName** | **string** |  | 
**AgentVersion** | **string** |  | 
**AgentDescription** | **string** |  | 
**WorkflowName** | **string** |  | 
**PromptName** | **string** |  | 
**ToolName** | **string** |  | 
**ToolType** | **string** |  | 
**ToolCallId** | **string** |  | 
**InputMessages** | **string** |  | 
**OutputMessages** | **string** |  | 
**SystemInstructions** | **string** |  | 
**ToolDefinitions** | **string** |  | 
**ToolCallArguments** | **string** |  | 
**ToolCallResult** | **string** |  | 
**RequestChoiceCount** | **int32** |  | 
**RequestMaxTokens** | **int32** |  | 
**RequestTemperature** | **NullableFloat32** |  | 
**RequestTopP** | **NullableFloat32** |  | 
**RequestTopK** | **NullableFloat32** |  | 
**RequestPresencePenalty** | **NullableFloat32** |  | 
**RequestFrequencyPenalty** | **NullableFloat32** |  | 
**RequestSeed** | **int32** |  | 
**RequestStopSequences** | **[]string** |  | 
**RequestEncodingFormats** | **[]string** |  | 
**ResponseFinishReasons** | **[]string** |  | 
**UsageInputTokens** | **int32** |  | 
**UsageOutputTokens** | **int32** |  | 
**UsageCacheReadInputTokens** | **int32** |  | 
**UsageCacheCreationInputTokens** | **int32** |  | 
**ResourceAttributes** | **map[string]string** |  | 
**SpanAttributes** | **map[string]string** |  | 
**ScopeName** | **string** |  | 
**ScopeVersion** | **string** |  | 

## Methods

### NewGetSpan

`func NewGetSpan(customerId string, organizationId string, workspaceId string, userId string, traceId string, spanId string, parentSpanId string, traceState string, startTime time.Time, endTime time.Time, durationNs int32, spanName string, spanKind string, serviceName string, statusCode string, statusMessage string, errorType string, operationName string, providerName string, requestModel string, responseModel string, responseId string, outputType string, conversationId string, dataSourceId string, agentId string, agentName string, agentVersion string, agentDescription string, workflowName string, promptName string, toolName string, toolType string, toolCallId string, inputMessages string, outputMessages string, systemInstructions string, toolDefinitions string, toolCallArguments string, toolCallResult string, requestChoiceCount int32, requestMaxTokens int32, requestTemperature NullableFloat32, requestTopP NullableFloat32, requestTopK NullableFloat32, requestPresencePenalty NullableFloat32, requestFrequencyPenalty NullableFloat32, requestSeed int32, requestStopSequences []string, requestEncodingFormats []string, responseFinishReasons []string, usageInputTokens int32, usageOutputTokens int32, usageCacheReadInputTokens int32, usageCacheCreationInputTokens int32, resourceAttributes map[string]string, spanAttributes map[string]string, scopeName string, scopeVersion string, ) *GetSpan`

NewGetSpan instantiates a new GetSpan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSpanWithDefaults

`func NewGetSpanWithDefaults() *GetSpan`

NewGetSpanWithDefaults instantiates a new GetSpan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *GetSpan) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *GetSpan) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *GetSpan) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetOrganizationId

`func (o *GetSpan) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GetSpan) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GetSpan) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetWorkspaceId

`func (o *GetSpan) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *GetSpan) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *GetSpan) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetUserId

`func (o *GetSpan) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetSpan) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetSpan) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetTraceId

`func (o *GetSpan) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *GetSpan) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *GetSpan) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetSpanId

`func (o *GetSpan) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *GetSpan) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *GetSpan) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.


### GetParentSpanId

`func (o *GetSpan) GetParentSpanId() string`

GetParentSpanId returns the ParentSpanId field if non-nil, zero value otherwise.

### GetParentSpanIdOk

`func (o *GetSpan) GetParentSpanIdOk() (*string, bool)`

GetParentSpanIdOk returns a tuple with the ParentSpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSpanId

`func (o *GetSpan) SetParentSpanId(v string)`

SetParentSpanId sets ParentSpanId field to given value.


### GetTraceState

`func (o *GetSpan) GetTraceState() string`

GetTraceState returns the TraceState field if non-nil, zero value otherwise.

### GetTraceStateOk

`func (o *GetSpan) GetTraceStateOk() (*string, bool)`

GetTraceStateOk returns a tuple with the TraceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceState

`func (o *GetSpan) SetTraceState(v string)`

SetTraceState sets TraceState field to given value.


### GetStartTime

`func (o *GetSpan) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *GetSpan) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *GetSpan) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *GetSpan) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *GetSpan) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *GetSpan) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetDurationNs

`func (o *GetSpan) GetDurationNs() int32`

GetDurationNs returns the DurationNs field if non-nil, zero value otherwise.

### GetDurationNsOk

`func (o *GetSpan) GetDurationNsOk() (*int32, bool)`

GetDurationNsOk returns a tuple with the DurationNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationNs

`func (o *GetSpan) SetDurationNs(v int32)`

SetDurationNs sets DurationNs field to given value.


### GetSpanName

`func (o *GetSpan) GetSpanName() string`

GetSpanName returns the SpanName field if non-nil, zero value otherwise.

### GetSpanNameOk

`func (o *GetSpan) GetSpanNameOk() (*string, bool)`

GetSpanNameOk returns a tuple with the SpanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanName

`func (o *GetSpan) SetSpanName(v string)`

SetSpanName sets SpanName field to given value.


### GetSpanKind

`func (o *GetSpan) GetSpanKind() string`

GetSpanKind returns the SpanKind field if non-nil, zero value otherwise.

### GetSpanKindOk

`func (o *GetSpan) GetSpanKindOk() (*string, bool)`

GetSpanKindOk returns a tuple with the SpanKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanKind

`func (o *GetSpan) SetSpanKind(v string)`

SetSpanKind sets SpanKind field to given value.


### GetServiceName

`func (o *GetSpan) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *GetSpan) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *GetSpan) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetStatusCode

`func (o *GetSpan) GetStatusCode() string`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *GetSpan) GetStatusCodeOk() (*string, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *GetSpan) SetStatusCode(v string)`

SetStatusCode sets StatusCode field to given value.


### GetStatusMessage

`func (o *GetSpan) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetSpan) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetSpan) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.


### GetErrorType

`func (o *GetSpan) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *GetSpan) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *GetSpan) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.


### GetOperationName

`func (o *GetSpan) GetOperationName() string`

GetOperationName returns the OperationName field if non-nil, zero value otherwise.

### GetOperationNameOk

`func (o *GetSpan) GetOperationNameOk() (*string, bool)`

GetOperationNameOk returns a tuple with the OperationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationName

`func (o *GetSpan) SetOperationName(v string)`

SetOperationName sets OperationName field to given value.


### GetProviderName

`func (o *GetSpan) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *GetSpan) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *GetSpan) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetRequestModel

`func (o *GetSpan) GetRequestModel() string`

GetRequestModel returns the RequestModel field if non-nil, zero value otherwise.

### GetRequestModelOk

`func (o *GetSpan) GetRequestModelOk() (*string, bool)`

GetRequestModelOk returns a tuple with the RequestModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestModel

`func (o *GetSpan) SetRequestModel(v string)`

SetRequestModel sets RequestModel field to given value.


### GetResponseModel

`func (o *GetSpan) GetResponseModel() string`

GetResponseModel returns the ResponseModel field if non-nil, zero value otherwise.

### GetResponseModelOk

`func (o *GetSpan) GetResponseModelOk() (*string, bool)`

GetResponseModelOk returns a tuple with the ResponseModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseModel

`func (o *GetSpan) SetResponseModel(v string)`

SetResponseModel sets ResponseModel field to given value.


### GetResponseId

`func (o *GetSpan) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetSpan) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetSpan) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.


### GetOutputType

`func (o *GetSpan) GetOutputType() string`

GetOutputType returns the OutputType field if non-nil, zero value otherwise.

### GetOutputTypeOk

`func (o *GetSpan) GetOutputTypeOk() (*string, bool)`

GetOutputTypeOk returns a tuple with the OutputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputType

`func (o *GetSpan) SetOutputType(v string)`

SetOutputType sets OutputType field to given value.


### GetConversationId

`func (o *GetSpan) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *GetSpan) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *GetSpan) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.


### GetDataSourceId

`func (o *GetSpan) GetDataSourceId() string`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *GetSpan) GetDataSourceIdOk() (*string, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *GetSpan) SetDataSourceId(v string)`

SetDataSourceId sets DataSourceId field to given value.


### GetAgentId

`func (o *GetSpan) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *GetSpan) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *GetSpan) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.


### GetAgentName

`func (o *GetSpan) GetAgentName() string`

GetAgentName returns the AgentName field if non-nil, zero value otherwise.

### GetAgentNameOk

`func (o *GetSpan) GetAgentNameOk() (*string, bool)`

GetAgentNameOk returns a tuple with the AgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentName

`func (o *GetSpan) SetAgentName(v string)`

SetAgentName sets AgentName field to given value.


### GetAgentVersion

`func (o *GetSpan) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *GetSpan) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *GetSpan) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.


### GetAgentDescription

`func (o *GetSpan) GetAgentDescription() string`

GetAgentDescription returns the AgentDescription field if non-nil, zero value otherwise.

### GetAgentDescriptionOk

`func (o *GetSpan) GetAgentDescriptionOk() (*string, bool)`

GetAgentDescriptionOk returns a tuple with the AgentDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentDescription

`func (o *GetSpan) SetAgentDescription(v string)`

SetAgentDescription sets AgentDescription field to given value.


### GetWorkflowName

`func (o *GetSpan) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *GetSpan) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *GetSpan) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.


### GetPromptName

`func (o *GetSpan) GetPromptName() string`

GetPromptName returns the PromptName field if non-nil, zero value otherwise.

### GetPromptNameOk

`func (o *GetSpan) GetPromptNameOk() (*string, bool)`

GetPromptNameOk returns a tuple with the PromptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptName

`func (o *GetSpan) SetPromptName(v string)`

SetPromptName sets PromptName field to given value.


### GetToolName

`func (o *GetSpan) GetToolName() string`

GetToolName returns the ToolName field if non-nil, zero value otherwise.

### GetToolNameOk

`func (o *GetSpan) GetToolNameOk() (*string, bool)`

GetToolNameOk returns a tuple with the ToolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolName

`func (o *GetSpan) SetToolName(v string)`

SetToolName sets ToolName field to given value.


### GetToolType

`func (o *GetSpan) GetToolType() string`

GetToolType returns the ToolType field if non-nil, zero value otherwise.

### GetToolTypeOk

`func (o *GetSpan) GetToolTypeOk() (*string, bool)`

GetToolTypeOk returns a tuple with the ToolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolType

`func (o *GetSpan) SetToolType(v string)`

SetToolType sets ToolType field to given value.


### GetToolCallId

`func (o *GetSpan) GetToolCallId() string`

GetToolCallId returns the ToolCallId field if non-nil, zero value otherwise.

### GetToolCallIdOk

`func (o *GetSpan) GetToolCallIdOk() (*string, bool)`

GetToolCallIdOk returns a tuple with the ToolCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCallId

`func (o *GetSpan) SetToolCallId(v string)`

SetToolCallId sets ToolCallId field to given value.


### GetInputMessages

`func (o *GetSpan) GetInputMessages() string`

GetInputMessages returns the InputMessages field if non-nil, zero value otherwise.

### GetInputMessagesOk

`func (o *GetSpan) GetInputMessagesOk() (*string, bool)`

GetInputMessagesOk returns a tuple with the InputMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputMessages

`func (o *GetSpan) SetInputMessages(v string)`

SetInputMessages sets InputMessages field to given value.


### GetOutputMessages

`func (o *GetSpan) GetOutputMessages() string`

GetOutputMessages returns the OutputMessages field if non-nil, zero value otherwise.

### GetOutputMessagesOk

`func (o *GetSpan) GetOutputMessagesOk() (*string, bool)`

GetOutputMessagesOk returns a tuple with the OutputMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputMessages

`func (o *GetSpan) SetOutputMessages(v string)`

SetOutputMessages sets OutputMessages field to given value.


### GetSystemInstructions

`func (o *GetSpan) GetSystemInstructions() string`

GetSystemInstructions returns the SystemInstructions field if non-nil, zero value otherwise.

### GetSystemInstructionsOk

`func (o *GetSpan) GetSystemInstructionsOk() (*string, bool)`

GetSystemInstructionsOk returns a tuple with the SystemInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemInstructions

`func (o *GetSpan) SetSystemInstructions(v string)`

SetSystemInstructions sets SystemInstructions field to given value.


### GetToolDefinitions

`func (o *GetSpan) GetToolDefinitions() string`

GetToolDefinitions returns the ToolDefinitions field if non-nil, zero value otherwise.

### GetToolDefinitionsOk

`func (o *GetSpan) GetToolDefinitionsOk() (*string, bool)`

GetToolDefinitionsOk returns a tuple with the ToolDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolDefinitions

`func (o *GetSpan) SetToolDefinitions(v string)`

SetToolDefinitions sets ToolDefinitions field to given value.


### GetToolCallArguments

`func (o *GetSpan) GetToolCallArguments() string`

GetToolCallArguments returns the ToolCallArguments field if non-nil, zero value otherwise.

### GetToolCallArgumentsOk

`func (o *GetSpan) GetToolCallArgumentsOk() (*string, bool)`

GetToolCallArgumentsOk returns a tuple with the ToolCallArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCallArguments

`func (o *GetSpan) SetToolCallArguments(v string)`

SetToolCallArguments sets ToolCallArguments field to given value.


### GetToolCallResult

`func (o *GetSpan) GetToolCallResult() string`

GetToolCallResult returns the ToolCallResult field if non-nil, zero value otherwise.

### GetToolCallResultOk

`func (o *GetSpan) GetToolCallResultOk() (*string, bool)`

GetToolCallResultOk returns a tuple with the ToolCallResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCallResult

`func (o *GetSpan) SetToolCallResult(v string)`

SetToolCallResult sets ToolCallResult field to given value.


### GetRequestChoiceCount

`func (o *GetSpan) GetRequestChoiceCount() int32`

GetRequestChoiceCount returns the RequestChoiceCount field if non-nil, zero value otherwise.

### GetRequestChoiceCountOk

`func (o *GetSpan) GetRequestChoiceCountOk() (*int32, bool)`

GetRequestChoiceCountOk returns a tuple with the RequestChoiceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestChoiceCount

`func (o *GetSpan) SetRequestChoiceCount(v int32)`

SetRequestChoiceCount sets RequestChoiceCount field to given value.


### GetRequestMaxTokens

`func (o *GetSpan) GetRequestMaxTokens() int32`

GetRequestMaxTokens returns the RequestMaxTokens field if non-nil, zero value otherwise.

### GetRequestMaxTokensOk

`func (o *GetSpan) GetRequestMaxTokensOk() (*int32, bool)`

GetRequestMaxTokensOk returns a tuple with the RequestMaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMaxTokens

`func (o *GetSpan) SetRequestMaxTokens(v int32)`

SetRequestMaxTokens sets RequestMaxTokens field to given value.


### GetRequestTemperature

`func (o *GetSpan) GetRequestTemperature() float32`

GetRequestTemperature returns the RequestTemperature field if non-nil, zero value otherwise.

### GetRequestTemperatureOk

`func (o *GetSpan) GetRequestTemperatureOk() (*float32, bool)`

GetRequestTemperatureOk returns a tuple with the RequestTemperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTemperature

`func (o *GetSpan) SetRequestTemperature(v float32)`

SetRequestTemperature sets RequestTemperature field to given value.


### SetRequestTemperatureNil

`func (o *GetSpan) SetRequestTemperatureNil(b bool)`

 SetRequestTemperatureNil sets the value for RequestTemperature to be an explicit nil

### UnsetRequestTemperature
`func (o *GetSpan) UnsetRequestTemperature()`

UnsetRequestTemperature ensures that no value is present for RequestTemperature, not even an explicit nil
### GetRequestTopP

`func (o *GetSpan) GetRequestTopP() float32`

GetRequestTopP returns the RequestTopP field if non-nil, zero value otherwise.

### GetRequestTopPOk

`func (o *GetSpan) GetRequestTopPOk() (*float32, bool)`

GetRequestTopPOk returns a tuple with the RequestTopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTopP

`func (o *GetSpan) SetRequestTopP(v float32)`

SetRequestTopP sets RequestTopP field to given value.


### SetRequestTopPNil

`func (o *GetSpan) SetRequestTopPNil(b bool)`

 SetRequestTopPNil sets the value for RequestTopP to be an explicit nil

### UnsetRequestTopP
`func (o *GetSpan) UnsetRequestTopP()`

UnsetRequestTopP ensures that no value is present for RequestTopP, not even an explicit nil
### GetRequestTopK

`func (o *GetSpan) GetRequestTopK() float32`

GetRequestTopK returns the RequestTopK field if non-nil, zero value otherwise.

### GetRequestTopKOk

`func (o *GetSpan) GetRequestTopKOk() (*float32, bool)`

GetRequestTopKOk returns a tuple with the RequestTopK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTopK

`func (o *GetSpan) SetRequestTopK(v float32)`

SetRequestTopK sets RequestTopK field to given value.


### SetRequestTopKNil

`func (o *GetSpan) SetRequestTopKNil(b bool)`

 SetRequestTopKNil sets the value for RequestTopK to be an explicit nil

### UnsetRequestTopK
`func (o *GetSpan) UnsetRequestTopK()`

UnsetRequestTopK ensures that no value is present for RequestTopK, not even an explicit nil
### GetRequestPresencePenalty

`func (o *GetSpan) GetRequestPresencePenalty() float32`

GetRequestPresencePenalty returns the RequestPresencePenalty field if non-nil, zero value otherwise.

### GetRequestPresencePenaltyOk

`func (o *GetSpan) GetRequestPresencePenaltyOk() (*float32, bool)`

GetRequestPresencePenaltyOk returns a tuple with the RequestPresencePenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestPresencePenalty

`func (o *GetSpan) SetRequestPresencePenalty(v float32)`

SetRequestPresencePenalty sets RequestPresencePenalty field to given value.


### SetRequestPresencePenaltyNil

`func (o *GetSpan) SetRequestPresencePenaltyNil(b bool)`

 SetRequestPresencePenaltyNil sets the value for RequestPresencePenalty to be an explicit nil

### UnsetRequestPresencePenalty
`func (o *GetSpan) UnsetRequestPresencePenalty()`

UnsetRequestPresencePenalty ensures that no value is present for RequestPresencePenalty, not even an explicit nil
### GetRequestFrequencyPenalty

`func (o *GetSpan) GetRequestFrequencyPenalty() float32`

GetRequestFrequencyPenalty returns the RequestFrequencyPenalty field if non-nil, zero value otherwise.

### GetRequestFrequencyPenaltyOk

`func (o *GetSpan) GetRequestFrequencyPenaltyOk() (*float32, bool)`

GetRequestFrequencyPenaltyOk returns a tuple with the RequestFrequencyPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestFrequencyPenalty

`func (o *GetSpan) SetRequestFrequencyPenalty(v float32)`

SetRequestFrequencyPenalty sets RequestFrequencyPenalty field to given value.


### SetRequestFrequencyPenaltyNil

`func (o *GetSpan) SetRequestFrequencyPenaltyNil(b bool)`

 SetRequestFrequencyPenaltyNil sets the value for RequestFrequencyPenalty to be an explicit nil

### UnsetRequestFrequencyPenalty
`func (o *GetSpan) UnsetRequestFrequencyPenalty()`

UnsetRequestFrequencyPenalty ensures that no value is present for RequestFrequencyPenalty, not even an explicit nil
### GetRequestSeed

`func (o *GetSpan) GetRequestSeed() int32`

GetRequestSeed returns the RequestSeed field if non-nil, zero value otherwise.

### GetRequestSeedOk

`func (o *GetSpan) GetRequestSeedOk() (*int32, bool)`

GetRequestSeedOk returns a tuple with the RequestSeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSeed

`func (o *GetSpan) SetRequestSeed(v int32)`

SetRequestSeed sets RequestSeed field to given value.


### GetRequestStopSequences

`func (o *GetSpan) GetRequestStopSequences() []string`

GetRequestStopSequences returns the RequestStopSequences field if non-nil, zero value otherwise.

### GetRequestStopSequencesOk

`func (o *GetSpan) GetRequestStopSequencesOk() (*[]string, bool)`

GetRequestStopSequencesOk returns a tuple with the RequestStopSequences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestStopSequences

`func (o *GetSpan) SetRequestStopSequences(v []string)`

SetRequestStopSequences sets RequestStopSequences field to given value.


### GetRequestEncodingFormats

`func (o *GetSpan) GetRequestEncodingFormats() []string`

GetRequestEncodingFormats returns the RequestEncodingFormats field if non-nil, zero value otherwise.

### GetRequestEncodingFormatsOk

`func (o *GetSpan) GetRequestEncodingFormatsOk() (*[]string, bool)`

GetRequestEncodingFormatsOk returns a tuple with the RequestEncodingFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestEncodingFormats

`func (o *GetSpan) SetRequestEncodingFormats(v []string)`

SetRequestEncodingFormats sets RequestEncodingFormats field to given value.


### GetResponseFinishReasons

`func (o *GetSpan) GetResponseFinishReasons() []string`

GetResponseFinishReasons returns the ResponseFinishReasons field if non-nil, zero value otherwise.

### GetResponseFinishReasonsOk

`func (o *GetSpan) GetResponseFinishReasonsOk() (*[]string, bool)`

GetResponseFinishReasonsOk returns a tuple with the ResponseFinishReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseFinishReasons

`func (o *GetSpan) SetResponseFinishReasons(v []string)`

SetResponseFinishReasons sets ResponseFinishReasons field to given value.


### GetUsageInputTokens

`func (o *GetSpan) GetUsageInputTokens() int32`

GetUsageInputTokens returns the UsageInputTokens field if non-nil, zero value otherwise.

### GetUsageInputTokensOk

`func (o *GetSpan) GetUsageInputTokensOk() (*int32, bool)`

GetUsageInputTokensOk returns a tuple with the UsageInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageInputTokens

`func (o *GetSpan) SetUsageInputTokens(v int32)`

SetUsageInputTokens sets UsageInputTokens field to given value.


### GetUsageOutputTokens

`func (o *GetSpan) GetUsageOutputTokens() int32`

GetUsageOutputTokens returns the UsageOutputTokens field if non-nil, zero value otherwise.

### GetUsageOutputTokensOk

`func (o *GetSpan) GetUsageOutputTokensOk() (*int32, bool)`

GetUsageOutputTokensOk returns a tuple with the UsageOutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageOutputTokens

`func (o *GetSpan) SetUsageOutputTokens(v int32)`

SetUsageOutputTokens sets UsageOutputTokens field to given value.


### GetUsageCacheReadInputTokens

`func (o *GetSpan) GetUsageCacheReadInputTokens() int32`

GetUsageCacheReadInputTokens returns the UsageCacheReadInputTokens field if non-nil, zero value otherwise.

### GetUsageCacheReadInputTokensOk

`func (o *GetSpan) GetUsageCacheReadInputTokensOk() (*int32, bool)`

GetUsageCacheReadInputTokensOk returns a tuple with the UsageCacheReadInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCacheReadInputTokens

`func (o *GetSpan) SetUsageCacheReadInputTokens(v int32)`

SetUsageCacheReadInputTokens sets UsageCacheReadInputTokens field to given value.


### GetUsageCacheCreationInputTokens

`func (o *GetSpan) GetUsageCacheCreationInputTokens() int32`

GetUsageCacheCreationInputTokens returns the UsageCacheCreationInputTokens field if non-nil, zero value otherwise.

### GetUsageCacheCreationInputTokensOk

`func (o *GetSpan) GetUsageCacheCreationInputTokensOk() (*int32, bool)`

GetUsageCacheCreationInputTokensOk returns a tuple with the UsageCacheCreationInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCacheCreationInputTokens

`func (o *GetSpan) SetUsageCacheCreationInputTokens(v int32)`

SetUsageCacheCreationInputTokens sets UsageCacheCreationInputTokens field to given value.


### GetResourceAttributes

`func (o *GetSpan) GetResourceAttributes() map[string]string`

GetResourceAttributes returns the ResourceAttributes field if non-nil, zero value otherwise.

### GetResourceAttributesOk

`func (o *GetSpan) GetResourceAttributesOk() (*map[string]string, bool)`

GetResourceAttributesOk returns a tuple with the ResourceAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAttributes

`func (o *GetSpan) SetResourceAttributes(v map[string]string)`

SetResourceAttributes sets ResourceAttributes field to given value.


### GetSpanAttributes

`func (o *GetSpan) GetSpanAttributes() map[string]string`

GetSpanAttributes returns the SpanAttributes field if non-nil, zero value otherwise.

### GetSpanAttributesOk

`func (o *GetSpan) GetSpanAttributesOk() (*map[string]string, bool)`

GetSpanAttributesOk returns a tuple with the SpanAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanAttributes

`func (o *GetSpan) SetSpanAttributes(v map[string]string)`

SetSpanAttributes sets SpanAttributes field to given value.


### GetScopeName

`func (o *GetSpan) GetScopeName() string`

GetScopeName returns the ScopeName field if non-nil, zero value otherwise.

### GetScopeNameOk

`func (o *GetSpan) GetScopeNameOk() (*string, bool)`

GetScopeNameOk returns a tuple with the ScopeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeName

`func (o *GetSpan) SetScopeName(v string)`

SetScopeName sets ScopeName field to given value.


### GetScopeVersion

`func (o *GetSpan) GetScopeVersion() string`

GetScopeVersion returns the ScopeVersion field if non-nil, zero value otherwise.

### GetScopeVersionOk

`func (o *GetSpan) GetScopeVersionOk() (*string, bool)`

GetScopeVersionOk returns a tuple with the ScopeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeVersion

`func (o *GetSpan) SetScopeVersion(v string)`

SetScopeVersion sets ScopeVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


