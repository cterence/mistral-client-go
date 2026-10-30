# GetSpanEvaluation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** |  | 
**OrganizationId** | **string** |  | 
**WorkspaceId** | **string** |  | 
**UserId** | **string** |  | 
**TraceId** | **string** |  | 
**SpanId** | **string** |  | 
**ResponseId** | **string** |  | 
**ConversationId** | **string** |  | 
**Timestamp** | **time.Time** |  | 
**EvaluationName** | **string** |  | 
**ScoreValue** | **float32** |  | 
**ScoreLabel** | **string** |  | 
**Explanation** | **string** |  | 
**Metadata** | **map[string]string** |  | 

## Methods

### NewGetSpanEvaluation

`func NewGetSpanEvaluation(customerId string, organizationId string, workspaceId string, userId string, traceId string, spanId string, responseId string, conversationId string, timestamp time.Time, evaluationName string, scoreValue float32, scoreLabel string, explanation string, metadata map[string]string, ) *GetSpanEvaluation`

NewGetSpanEvaluation instantiates a new GetSpanEvaluation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSpanEvaluationWithDefaults

`func NewGetSpanEvaluationWithDefaults() *GetSpanEvaluation`

NewGetSpanEvaluationWithDefaults instantiates a new GetSpanEvaluation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *GetSpanEvaluation) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *GetSpanEvaluation) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *GetSpanEvaluation) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetOrganizationId

`func (o *GetSpanEvaluation) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GetSpanEvaluation) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GetSpanEvaluation) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetWorkspaceId

`func (o *GetSpanEvaluation) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *GetSpanEvaluation) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *GetSpanEvaluation) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetUserId

`func (o *GetSpanEvaluation) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetSpanEvaluation) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetSpanEvaluation) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetTraceId

`func (o *GetSpanEvaluation) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *GetSpanEvaluation) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *GetSpanEvaluation) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetSpanId

`func (o *GetSpanEvaluation) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *GetSpanEvaluation) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *GetSpanEvaluation) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.


### GetResponseId

`func (o *GetSpanEvaluation) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetSpanEvaluation) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetSpanEvaluation) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.


### GetConversationId

`func (o *GetSpanEvaluation) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *GetSpanEvaluation) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *GetSpanEvaluation) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.


### GetTimestamp

`func (o *GetSpanEvaluation) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetSpanEvaluation) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetSpanEvaluation) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetEvaluationName

`func (o *GetSpanEvaluation) GetEvaluationName() string`

GetEvaluationName returns the EvaluationName field if non-nil, zero value otherwise.

### GetEvaluationNameOk

`func (o *GetSpanEvaluation) GetEvaluationNameOk() (*string, bool)`

GetEvaluationNameOk returns a tuple with the EvaluationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationName

`func (o *GetSpanEvaluation) SetEvaluationName(v string)`

SetEvaluationName sets EvaluationName field to given value.


### GetScoreValue

`func (o *GetSpanEvaluation) GetScoreValue() float32`

GetScoreValue returns the ScoreValue field if non-nil, zero value otherwise.

### GetScoreValueOk

`func (o *GetSpanEvaluation) GetScoreValueOk() (*float32, bool)`

GetScoreValueOk returns a tuple with the ScoreValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreValue

`func (o *GetSpanEvaluation) SetScoreValue(v float32)`

SetScoreValue sets ScoreValue field to given value.


### GetScoreLabel

`func (o *GetSpanEvaluation) GetScoreLabel() string`

GetScoreLabel returns the ScoreLabel field if non-nil, zero value otherwise.

### GetScoreLabelOk

`func (o *GetSpanEvaluation) GetScoreLabelOk() (*string, bool)`

GetScoreLabelOk returns a tuple with the ScoreLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreLabel

`func (o *GetSpanEvaluation) SetScoreLabel(v string)`

SetScoreLabel sets ScoreLabel field to given value.


### GetExplanation

`func (o *GetSpanEvaluation) GetExplanation() string`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *GetSpanEvaluation) GetExplanationOk() (*string, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *GetSpanEvaluation) SetExplanation(v string)`

SetExplanation sets Explanation field to given value.


### GetMetadata

`func (o *GetSpanEvaluation) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetSpanEvaluation) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetSpanEvaluation) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


