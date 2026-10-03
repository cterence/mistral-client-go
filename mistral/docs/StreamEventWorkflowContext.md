# StreamEventWorkflowContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | **string** |  | 
**WorkflowName** | **string** |  | 
**WorkflowExecId** | **string** |  | 
**ParentWorkflowExecId** | Pointer to **NullableString** |  | [optional] 
**RootWorkflowExecId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewStreamEventWorkflowContext

`func NewStreamEventWorkflowContext(namespace string, workflowName string, workflowExecId string, ) *StreamEventWorkflowContext`

NewStreamEventWorkflowContext instantiates a new StreamEventWorkflowContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamEventWorkflowContextWithDefaults

`func NewStreamEventWorkflowContextWithDefaults() *StreamEventWorkflowContext`

NewStreamEventWorkflowContextWithDefaults instantiates a new StreamEventWorkflowContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *StreamEventWorkflowContext) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *StreamEventWorkflowContext) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *StreamEventWorkflowContext) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetWorkflowName

`func (o *StreamEventWorkflowContext) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *StreamEventWorkflowContext) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *StreamEventWorkflowContext) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.


### GetWorkflowExecId

`func (o *StreamEventWorkflowContext) GetWorkflowExecId() string`

GetWorkflowExecId returns the WorkflowExecId field if non-nil, zero value otherwise.

### GetWorkflowExecIdOk

`func (o *StreamEventWorkflowContext) GetWorkflowExecIdOk() (*string, bool)`

GetWorkflowExecIdOk returns a tuple with the WorkflowExecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowExecId

`func (o *StreamEventWorkflowContext) SetWorkflowExecId(v string)`

SetWorkflowExecId sets WorkflowExecId field to given value.


### GetParentWorkflowExecId

`func (o *StreamEventWorkflowContext) GetParentWorkflowExecId() string`

GetParentWorkflowExecId returns the ParentWorkflowExecId field if non-nil, zero value otherwise.

### GetParentWorkflowExecIdOk

`func (o *StreamEventWorkflowContext) GetParentWorkflowExecIdOk() (*string, bool)`

GetParentWorkflowExecIdOk returns a tuple with the ParentWorkflowExecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentWorkflowExecId

`func (o *StreamEventWorkflowContext) SetParentWorkflowExecId(v string)`

SetParentWorkflowExecId sets ParentWorkflowExecId field to given value.

### HasParentWorkflowExecId

`func (o *StreamEventWorkflowContext) HasParentWorkflowExecId() bool`

HasParentWorkflowExecId returns a boolean if a field has been set.

### SetParentWorkflowExecIdNil

`func (o *StreamEventWorkflowContext) SetParentWorkflowExecIdNil(b bool)`

 SetParentWorkflowExecIdNil sets the value for ParentWorkflowExecId to be an explicit nil

### UnsetParentWorkflowExecId
`func (o *StreamEventWorkflowContext) UnsetParentWorkflowExecId()`

UnsetParentWorkflowExecId ensures that no value is present for ParentWorkflowExecId, not even an explicit nil
### GetRootWorkflowExecId

`func (o *StreamEventWorkflowContext) GetRootWorkflowExecId() string`

GetRootWorkflowExecId returns the RootWorkflowExecId field if non-nil, zero value otherwise.

### GetRootWorkflowExecIdOk

`func (o *StreamEventWorkflowContext) GetRootWorkflowExecIdOk() (*string, bool)`

GetRootWorkflowExecIdOk returns a tuple with the RootWorkflowExecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootWorkflowExecId

`func (o *StreamEventWorkflowContext) SetRootWorkflowExecId(v string)`

SetRootWorkflowExecId sets RootWorkflowExecId field to given value.

### HasRootWorkflowExecId

`func (o *StreamEventWorkflowContext) HasRootWorkflowExecId() bool`

HasRootWorkflowExecId returns a boolean if a field has been set.

### SetRootWorkflowExecIdNil

`func (o *StreamEventWorkflowContext) SetRootWorkflowExecIdNil(b bool)`

 SetRootWorkflowExecIdNil sets the value for RootWorkflowExecId to be an explicit nil

### UnsetRootWorkflowExecId
`func (o *StreamEventWorkflowContext) UnsetRootWorkflowExecId()`

UnsetRootWorkflowExecId ensures that no value is present for RootWorkflowExecId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


