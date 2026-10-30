# GetTrace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** |  | 
**OrganizationId** | **string** |  | 
**WorkspaceId** | **string** |  | 
**UserId** | **string** |  | 
**TraceId** | **string** |  | 
**RootSpanId** | **string** |  | 
**RootSpanName** | **string** |  | 
**StartTime** | **time.Time** |  | 
**EndTime** | **time.Time** |  | 
**DurationNs** | **int32** |  | 
**ServiceName** | **string** |  | 
**Environment** | **string** |  | 
**ConversationId** | **string** |  | 
**WorkflowName** | **string** |  | 
**AgentId** | **string** |  | 
**AgentName** | **string** |  | 
**StatusCode** | **string** |  | 
**ErrorCount** | **int32** |  | 
**SpanCount** | **int32** |  | 
**GenAiSpanCount** | **int32** |  | 
**LlmCallCount** | **int32** |  | 
**ToolCallCount** | **int32** |  | 
**RetrievalCount** | **int32** |  | 
**EvaluationCount** | **int32** |  | 
**InputTokens** | **int32** |  | 
**OutputTokens** | **int32** |  | 
**CacheReadInputTokens** | **int32** |  | 
**CacheCreationInputTokens** | **int32** |  | 
**ModelsUsed** | **[]string** |  | 
**ToolsUsed** | **[]string** |  | 
**FirstTurnLastInputMessage** | **string** |  | 
**FirstTurnLastOutputMessage** | **string** |  | 
**LastTurnLastInputMessage** | **string** |  | 
**LastTurnLastOutputMessage** | **string** |  | 

## Methods

### NewGetTrace

`func NewGetTrace(customerId string, organizationId string, workspaceId string, userId string, traceId string, rootSpanId string, rootSpanName string, startTime time.Time, endTime time.Time, durationNs int32, serviceName string, environment string, conversationId string, workflowName string, agentId string, agentName string, statusCode string, errorCount int32, spanCount int32, genAiSpanCount int32, llmCallCount int32, toolCallCount int32, retrievalCount int32, evaluationCount int32, inputTokens int32, outputTokens int32, cacheReadInputTokens int32, cacheCreationInputTokens int32, modelsUsed []string, toolsUsed []string, firstTurnLastInputMessage string, firstTurnLastOutputMessage string, lastTurnLastInputMessage string, lastTurnLastOutputMessage string, ) *GetTrace`

NewGetTrace instantiates a new GetTrace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTraceWithDefaults

`func NewGetTraceWithDefaults() *GetTrace`

NewGetTraceWithDefaults instantiates a new GetTrace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *GetTrace) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *GetTrace) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *GetTrace) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetOrganizationId

`func (o *GetTrace) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GetTrace) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GetTrace) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetWorkspaceId

`func (o *GetTrace) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *GetTrace) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *GetTrace) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetUserId

`func (o *GetTrace) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetTrace) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetTrace) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetTraceId

`func (o *GetTrace) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *GetTrace) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *GetTrace) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetRootSpanId

`func (o *GetTrace) GetRootSpanId() string`

GetRootSpanId returns the RootSpanId field if non-nil, zero value otherwise.

### GetRootSpanIdOk

`func (o *GetTrace) GetRootSpanIdOk() (*string, bool)`

GetRootSpanIdOk returns a tuple with the RootSpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootSpanId

`func (o *GetTrace) SetRootSpanId(v string)`

SetRootSpanId sets RootSpanId field to given value.


### GetRootSpanName

`func (o *GetTrace) GetRootSpanName() string`

GetRootSpanName returns the RootSpanName field if non-nil, zero value otherwise.

### GetRootSpanNameOk

`func (o *GetTrace) GetRootSpanNameOk() (*string, bool)`

GetRootSpanNameOk returns a tuple with the RootSpanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootSpanName

`func (o *GetTrace) SetRootSpanName(v string)`

SetRootSpanName sets RootSpanName field to given value.


### GetStartTime

`func (o *GetTrace) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *GetTrace) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *GetTrace) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *GetTrace) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *GetTrace) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *GetTrace) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetDurationNs

`func (o *GetTrace) GetDurationNs() int32`

GetDurationNs returns the DurationNs field if non-nil, zero value otherwise.

### GetDurationNsOk

`func (o *GetTrace) GetDurationNsOk() (*int32, bool)`

GetDurationNsOk returns a tuple with the DurationNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationNs

`func (o *GetTrace) SetDurationNs(v int32)`

SetDurationNs sets DurationNs field to given value.


### GetServiceName

`func (o *GetTrace) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *GetTrace) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *GetTrace) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetEnvironment

`func (o *GetTrace) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *GetTrace) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *GetTrace) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetConversationId

`func (o *GetTrace) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *GetTrace) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *GetTrace) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.


### GetWorkflowName

`func (o *GetTrace) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *GetTrace) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *GetTrace) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.


### GetAgentId

`func (o *GetTrace) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *GetTrace) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *GetTrace) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.


### GetAgentName

`func (o *GetTrace) GetAgentName() string`

GetAgentName returns the AgentName field if non-nil, zero value otherwise.

### GetAgentNameOk

`func (o *GetTrace) GetAgentNameOk() (*string, bool)`

GetAgentNameOk returns a tuple with the AgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentName

`func (o *GetTrace) SetAgentName(v string)`

SetAgentName sets AgentName field to given value.


### GetStatusCode

`func (o *GetTrace) GetStatusCode() string`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *GetTrace) GetStatusCodeOk() (*string, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *GetTrace) SetStatusCode(v string)`

SetStatusCode sets StatusCode field to given value.


### GetErrorCount

`func (o *GetTrace) GetErrorCount() int32`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *GetTrace) GetErrorCountOk() (*int32, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount

`func (o *GetTrace) SetErrorCount(v int32)`

SetErrorCount sets ErrorCount field to given value.


### GetSpanCount

`func (o *GetTrace) GetSpanCount() int32`

GetSpanCount returns the SpanCount field if non-nil, zero value otherwise.

### GetSpanCountOk

`func (o *GetTrace) GetSpanCountOk() (*int32, bool)`

GetSpanCountOk returns a tuple with the SpanCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanCount

`func (o *GetTrace) SetSpanCount(v int32)`

SetSpanCount sets SpanCount field to given value.


### GetGenAiSpanCount

`func (o *GetTrace) GetGenAiSpanCount() int32`

GetGenAiSpanCount returns the GenAiSpanCount field if non-nil, zero value otherwise.

### GetGenAiSpanCountOk

`func (o *GetTrace) GetGenAiSpanCountOk() (*int32, bool)`

GetGenAiSpanCountOk returns a tuple with the GenAiSpanCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenAiSpanCount

`func (o *GetTrace) SetGenAiSpanCount(v int32)`

SetGenAiSpanCount sets GenAiSpanCount field to given value.


### GetLlmCallCount

`func (o *GetTrace) GetLlmCallCount() int32`

GetLlmCallCount returns the LlmCallCount field if non-nil, zero value otherwise.

### GetLlmCallCountOk

`func (o *GetTrace) GetLlmCallCountOk() (*int32, bool)`

GetLlmCallCountOk returns a tuple with the LlmCallCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlmCallCount

`func (o *GetTrace) SetLlmCallCount(v int32)`

SetLlmCallCount sets LlmCallCount field to given value.


### GetToolCallCount

`func (o *GetTrace) GetToolCallCount() int32`

GetToolCallCount returns the ToolCallCount field if non-nil, zero value otherwise.

### GetToolCallCountOk

`func (o *GetTrace) GetToolCallCountOk() (*int32, bool)`

GetToolCallCountOk returns a tuple with the ToolCallCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCallCount

`func (o *GetTrace) SetToolCallCount(v int32)`

SetToolCallCount sets ToolCallCount field to given value.


### GetRetrievalCount

`func (o *GetTrace) GetRetrievalCount() int32`

GetRetrievalCount returns the RetrievalCount field if non-nil, zero value otherwise.

### GetRetrievalCountOk

`func (o *GetTrace) GetRetrievalCountOk() (*int32, bool)`

GetRetrievalCountOk returns a tuple with the RetrievalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrievalCount

`func (o *GetTrace) SetRetrievalCount(v int32)`

SetRetrievalCount sets RetrievalCount field to given value.


### GetEvaluationCount

`func (o *GetTrace) GetEvaluationCount() int32`

GetEvaluationCount returns the EvaluationCount field if non-nil, zero value otherwise.

### GetEvaluationCountOk

`func (o *GetTrace) GetEvaluationCountOk() (*int32, bool)`

GetEvaluationCountOk returns a tuple with the EvaluationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationCount

`func (o *GetTrace) SetEvaluationCount(v int32)`

SetEvaluationCount sets EvaluationCount field to given value.


### GetInputTokens

`func (o *GetTrace) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *GetTrace) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *GetTrace) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.


### GetOutputTokens

`func (o *GetTrace) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *GetTrace) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *GetTrace) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.


### GetCacheReadInputTokens

`func (o *GetTrace) GetCacheReadInputTokens() int32`

GetCacheReadInputTokens returns the CacheReadInputTokens field if non-nil, zero value otherwise.

### GetCacheReadInputTokensOk

`func (o *GetTrace) GetCacheReadInputTokensOk() (*int32, bool)`

GetCacheReadInputTokensOk returns a tuple with the CacheReadInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadInputTokens

`func (o *GetTrace) SetCacheReadInputTokens(v int32)`

SetCacheReadInputTokens sets CacheReadInputTokens field to given value.


### GetCacheCreationInputTokens

`func (o *GetTrace) GetCacheCreationInputTokens() int32`

GetCacheCreationInputTokens returns the CacheCreationInputTokens field if non-nil, zero value otherwise.

### GetCacheCreationInputTokensOk

`func (o *GetTrace) GetCacheCreationInputTokensOk() (*int32, bool)`

GetCacheCreationInputTokensOk returns a tuple with the CacheCreationInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCreationInputTokens

`func (o *GetTrace) SetCacheCreationInputTokens(v int32)`

SetCacheCreationInputTokens sets CacheCreationInputTokens field to given value.


### GetModelsUsed

`func (o *GetTrace) GetModelsUsed() []string`

GetModelsUsed returns the ModelsUsed field if non-nil, zero value otherwise.

### GetModelsUsedOk

`func (o *GetTrace) GetModelsUsedOk() (*[]string, bool)`

GetModelsUsedOk returns a tuple with the ModelsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelsUsed

`func (o *GetTrace) SetModelsUsed(v []string)`

SetModelsUsed sets ModelsUsed field to given value.


### GetToolsUsed

`func (o *GetTrace) GetToolsUsed() []string`

GetToolsUsed returns the ToolsUsed field if non-nil, zero value otherwise.

### GetToolsUsedOk

`func (o *GetTrace) GetToolsUsedOk() (*[]string, bool)`

GetToolsUsedOk returns a tuple with the ToolsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolsUsed

`func (o *GetTrace) SetToolsUsed(v []string)`

SetToolsUsed sets ToolsUsed field to given value.


### GetFirstTurnLastInputMessage

`func (o *GetTrace) GetFirstTurnLastInputMessage() string`

GetFirstTurnLastInputMessage returns the FirstTurnLastInputMessage field if non-nil, zero value otherwise.

### GetFirstTurnLastInputMessageOk

`func (o *GetTrace) GetFirstTurnLastInputMessageOk() (*string, bool)`

GetFirstTurnLastInputMessageOk returns a tuple with the FirstTurnLastInputMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTurnLastInputMessage

`func (o *GetTrace) SetFirstTurnLastInputMessage(v string)`

SetFirstTurnLastInputMessage sets FirstTurnLastInputMessage field to given value.


### GetFirstTurnLastOutputMessage

`func (o *GetTrace) GetFirstTurnLastOutputMessage() string`

GetFirstTurnLastOutputMessage returns the FirstTurnLastOutputMessage field if non-nil, zero value otherwise.

### GetFirstTurnLastOutputMessageOk

`func (o *GetTrace) GetFirstTurnLastOutputMessageOk() (*string, bool)`

GetFirstTurnLastOutputMessageOk returns a tuple with the FirstTurnLastOutputMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTurnLastOutputMessage

`func (o *GetTrace) SetFirstTurnLastOutputMessage(v string)`

SetFirstTurnLastOutputMessage sets FirstTurnLastOutputMessage field to given value.


### GetLastTurnLastInputMessage

`func (o *GetTrace) GetLastTurnLastInputMessage() string`

GetLastTurnLastInputMessage returns the LastTurnLastInputMessage field if non-nil, zero value otherwise.

### GetLastTurnLastInputMessageOk

`func (o *GetTrace) GetLastTurnLastInputMessageOk() (*string, bool)`

GetLastTurnLastInputMessageOk returns a tuple with the LastTurnLastInputMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTurnLastInputMessage

`func (o *GetTrace) SetLastTurnLastInputMessage(v string)`

SetLastTurnLastInputMessage sets LastTurnLastInputMessage field to given value.


### GetLastTurnLastOutputMessage

`func (o *GetTrace) GetLastTurnLastOutputMessage() string`

GetLastTurnLastOutputMessage returns the LastTurnLastOutputMessage field if non-nil, zero value otherwise.

### GetLastTurnLastOutputMessageOk

`func (o *GetTrace) GetLastTurnLastOutputMessageOk() (*string, bool)`

GetLastTurnLastOutputMessageOk returns a tuple with the LastTurnLastOutputMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTurnLastOutputMessage

`func (o *GetTrace) SetLastTurnLastOutputMessage(v string)`

SetLastTurnLastOutputMessage sets LastTurnLastOutputMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


