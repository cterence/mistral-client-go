# GetLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** |  | 
**OrganizationId** | **string** |  | 
**WorkspaceId** | **string** |  | 
**UserId** | **string** |  | 
**Timestamp** | **time.Time** |  | 
**TraceId** | **string** |  | 
**SpanId** | **string** |  | 
**TraceFlags** | **int32** |  | 
**SeverityText** | **string** |  | 
**SeverityNumber** | **int32** |  | 
**ServiceName** | **string** |  | 
**Body** | **string** |  | 
**EventName** | **string** |  | 
**ResourceSchemaUrl** | **string** |  | 
**ResourceAttributes** | **map[string]string** |  | 
**ScopeSchemaUrl** | **string** |  | 
**ScopeName** | **string** |  | 
**ScopeVersion** | **string** |  | 
**ScopeAttributes** | **map[string]string** |  | 
**LogAttributes** | **map[string]string** |  | 

## Methods

### NewGetLog

`func NewGetLog(customerId string, organizationId string, workspaceId string, userId string, timestamp time.Time, traceId string, spanId string, traceFlags int32, severityText string, severityNumber int32, serviceName string, body string, eventName string, resourceSchemaUrl string, resourceAttributes map[string]string, scopeSchemaUrl string, scopeName string, scopeVersion string, scopeAttributes map[string]string, logAttributes map[string]string, ) *GetLog`

NewGetLog instantiates a new GetLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLogWithDefaults

`func NewGetLogWithDefaults() *GetLog`

NewGetLogWithDefaults instantiates a new GetLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *GetLog) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *GetLog) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *GetLog) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetOrganizationId

`func (o *GetLog) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GetLog) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GetLog) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetWorkspaceId

`func (o *GetLog) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *GetLog) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *GetLog) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetUserId

`func (o *GetLog) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetLog) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetLog) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetTimestamp

`func (o *GetLog) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetLog) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetLog) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetTraceId

`func (o *GetLog) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *GetLog) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *GetLog) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetSpanId

`func (o *GetLog) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *GetLog) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *GetLog) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.


### GetTraceFlags

`func (o *GetLog) GetTraceFlags() int32`

GetTraceFlags returns the TraceFlags field if non-nil, zero value otherwise.

### GetTraceFlagsOk

`func (o *GetLog) GetTraceFlagsOk() (*int32, bool)`

GetTraceFlagsOk returns a tuple with the TraceFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceFlags

`func (o *GetLog) SetTraceFlags(v int32)`

SetTraceFlags sets TraceFlags field to given value.


### GetSeverityText

`func (o *GetLog) GetSeverityText() string`

GetSeverityText returns the SeverityText field if non-nil, zero value otherwise.

### GetSeverityTextOk

`func (o *GetLog) GetSeverityTextOk() (*string, bool)`

GetSeverityTextOk returns a tuple with the SeverityText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityText

`func (o *GetLog) SetSeverityText(v string)`

SetSeverityText sets SeverityText field to given value.


### GetSeverityNumber

`func (o *GetLog) GetSeverityNumber() int32`

GetSeverityNumber returns the SeverityNumber field if non-nil, zero value otherwise.

### GetSeverityNumberOk

`func (o *GetLog) GetSeverityNumberOk() (*int32, bool)`

GetSeverityNumberOk returns a tuple with the SeverityNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityNumber

`func (o *GetLog) SetSeverityNumber(v int32)`

SetSeverityNumber sets SeverityNumber field to given value.


### GetServiceName

`func (o *GetLog) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *GetLog) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *GetLog) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetBody

`func (o *GetLog) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *GetLog) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *GetLog) SetBody(v string)`

SetBody sets Body field to given value.


### GetEventName

`func (o *GetLog) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *GetLog) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *GetLog) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetResourceSchemaUrl

`func (o *GetLog) GetResourceSchemaUrl() string`

GetResourceSchemaUrl returns the ResourceSchemaUrl field if non-nil, zero value otherwise.

### GetResourceSchemaUrlOk

`func (o *GetLog) GetResourceSchemaUrlOk() (*string, bool)`

GetResourceSchemaUrlOk returns a tuple with the ResourceSchemaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSchemaUrl

`func (o *GetLog) SetResourceSchemaUrl(v string)`

SetResourceSchemaUrl sets ResourceSchemaUrl field to given value.


### GetResourceAttributes

`func (o *GetLog) GetResourceAttributes() map[string]string`

GetResourceAttributes returns the ResourceAttributes field if non-nil, zero value otherwise.

### GetResourceAttributesOk

`func (o *GetLog) GetResourceAttributesOk() (*map[string]string, bool)`

GetResourceAttributesOk returns a tuple with the ResourceAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAttributes

`func (o *GetLog) SetResourceAttributes(v map[string]string)`

SetResourceAttributes sets ResourceAttributes field to given value.


### GetScopeSchemaUrl

`func (o *GetLog) GetScopeSchemaUrl() string`

GetScopeSchemaUrl returns the ScopeSchemaUrl field if non-nil, zero value otherwise.

### GetScopeSchemaUrlOk

`func (o *GetLog) GetScopeSchemaUrlOk() (*string, bool)`

GetScopeSchemaUrlOk returns a tuple with the ScopeSchemaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeSchemaUrl

`func (o *GetLog) SetScopeSchemaUrl(v string)`

SetScopeSchemaUrl sets ScopeSchemaUrl field to given value.


### GetScopeName

`func (o *GetLog) GetScopeName() string`

GetScopeName returns the ScopeName field if non-nil, zero value otherwise.

### GetScopeNameOk

`func (o *GetLog) GetScopeNameOk() (*string, bool)`

GetScopeNameOk returns a tuple with the ScopeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeName

`func (o *GetLog) SetScopeName(v string)`

SetScopeName sets ScopeName field to given value.


### GetScopeVersion

`func (o *GetLog) GetScopeVersion() string`

GetScopeVersion returns the ScopeVersion field if non-nil, zero value otherwise.

### GetScopeVersionOk

`func (o *GetLog) GetScopeVersionOk() (*string, bool)`

GetScopeVersionOk returns a tuple with the ScopeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeVersion

`func (o *GetLog) SetScopeVersion(v string)`

SetScopeVersion sets ScopeVersion field to given value.


### GetScopeAttributes

`func (o *GetLog) GetScopeAttributes() map[string]string`

GetScopeAttributes returns the ScopeAttributes field if non-nil, zero value otherwise.

### GetScopeAttributesOk

`func (o *GetLog) GetScopeAttributesOk() (*map[string]string, bool)`

GetScopeAttributesOk returns a tuple with the ScopeAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeAttributes

`func (o *GetLog) SetScopeAttributes(v map[string]string)`

SetScopeAttributes sets ScopeAttributes field to given value.


### GetLogAttributes

`func (o *GetLog) GetLogAttributes() map[string]string`

GetLogAttributes returns the LogAttributes field if non-nil, zero value otherwise.

### GetLogAttributesOk

`func (o *GetLog) GetLogAttributesOk() (*map[string]string, bool)`

GetLogAttributesOk returns a tuple with the LogAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogAttributes

`func (o *GetLog) SetLogAttributes(v map[string]string)`

SetLogAttributes sets LogAttributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


