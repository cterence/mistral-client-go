# WorkflowMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedNamespace** | Pointer to **NullableString** | Namespace for shared workflows, None if user-owned | [optional] 

## Methods

### NewWorkflowMetadata

`func NewWorkflowMetadata() *WorkflowMetadata`

NewWorkflowMetadata instantiates a new WorkflowMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowMetadataWithDefaults

`func NewWorkflowMetadataWithDefaults() *WorkflowMetadata`

NewWorkflowMetadataWithDefaults instantiates a new WorkflowMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharedNamespace

`func (o *WorkflowMetadata) GetSharedNamespace() string`

GetSharedNamespace returns the SharedNamespace field if non-nil, zero value otherwise.

### GetSharedNamespaceOk

`func (o *WorkflowMetadata) GetSharedNamespaceOk() (*string, bool)`

GetSharedNamespaceOk returns a tuple with the SharedNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedNamespace

`func (o *WorkflowMetadata) SetSharedNamespace(v string)`

SetSharedNamespace sets SharedNamespace field to given value.

### HasSharedNamespace

`func (o *WorkflowMetadata) HasSharedNamespace() bool`

HasSharedNamespace returns a boolean if a field has been set.

### SetSharedNamespaceNil

`func (o *WorkflowMetadata) SetSharedNamespaceNil(b bool)`

 SetSharedNamespaceNil sets the value for SharedNamespace to be an explicit nil

### UnsetSharedNamespace
`func (o *WorkflowMetadata) UnsetSharedNamespace()`

UnsetSharedNamespace ensures that no value is present for SharedNamespace, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


