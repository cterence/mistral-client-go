# CustomTaskCanceledAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomTaskId** | **string** | Unique identifier for the custom task within the workflow. | 
**CustomTaskType** | **string** | The type/category of the custom task (e.g., &#39;llm_call&#39;, &#39;api_request&#39;). | 
**Reason** | Pointer to **NullableString** | Optional reason provided for the cancellation. | [optional] 

## Methods

### NewCustomTaskCanceledAttributes

`func NewCustomTaskCanceledAttributes(customTaskId string, customTaskType string, ) *CustomTaskCanceledAttributes`

NewCustomTaskCanceledAttributes instantiates a new CustomTaskCanceledAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomTaskCanceledAttributesWithDefaults

`func NewCustomTaskCanceledAttributesWithDefaults() *CustomTaskCanceledAttributes`

NewCustomTaskCanceledAttributesWithDefaults instantiates a new CustomTaskCanceledAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomTaskId

`func (o *CustomTaskCanceledAttributes) GetCustomTaskId() string`

GetCustomTaskId returns the CustomTaskId field if non-nil, zero value otherwise.

### GetCustomTaskIdOk

`func (o *CustomTaskCanceledAttributes) GetCustomTaskIdOk() (*string, bool)`

GetCustomTaskIdOk returns a tuple with the CustomTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTaskId

`func (o *CustomTaskCanceledAttributes) SetCustomTaskId(v string)`

SetCustomTaskId sets CustomTaskId field to given value.


### GetCustomTaskType

`func (o *CustomTaskCanceledAttributes) GetCustomTaskType() string`

GetCustomTaskType returns the CustomTaskType field if non-nil, zero value otherwise.

### GetCustomTaskTypeOk

`func (o *CustomTaskCanceledAttributes) GetCustomTaskTypeOk() (*string, bool)`

GetCustomTaskTypeOk returns a tuple with the CustomTaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTaskType

`func (o *CustomTaskCanceledAttributes) SetCustomTaskType(v string)`

SetCustomTaskType sets CustomTaskType field to given value.


### GetReason

`func (o *CustomTaskCanceledAttributes) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CustomTaskCanceledAttributes) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CustomTaskCanceledAttributes) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CustomTaskCanceledAttributes) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *CustomTaskCanceledAttributes) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *CustomTaskCanceledAttributes) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


