# CustomTaskFailedAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomTaskId** | **string** | Unique identifier for the custom task within the workflow. | 
**CustomTaskType** | **string** | The type/category of the custom task (e.g., &#39;llm_call&#39;, &#39;api_request&#39;). | 
**Failure** | [**Failure**](Failure.md) | Details about the failure that caused the task to fail. | 

## Methods

### NewCustomTaskFailedAttributes

`func NewCustomTaskFailedAttributes(customTaskId string, customTaskType string, failure Failure, ) *CustomTaskFailedAttributes`

NewCustomTaskFailedAttributes instantiates a new CustomTaskFailedAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomTaskFailedAttributesWithDefaults

`func NewCustomTaskFailedAttributesWithDefaults() *CustomTaskFailedAttributes`

NewCustomTaskFailedAttributesWithDefaults instantiates a new CustomTaskFailedAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomTaskId

`func (o *CustomTaskFailedAttributes) GetCustomTaskId() string`

GetCustomTaskId returns the CustomTaskId field if non-nil, zero value otherwise.

### GetCustomTaskIdOk

`func (o *CustomTaskFailedAttributes) GetCustomTaskIdOk() (*string, bool)`

GetCustomTaskIdOk returns a tuple with the CustomTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTaskId

`func (o *CustomTaskFailedAttributes) SetCustomTaskId(v string)`

SetCustomTaskId sets CustomTaskId field to given value.


### GetCustomTaskType

`func (o *CustomTaskFailedAttributes) GetCustomTaskType() string`

GetCustomTaskType returns the CustomTaskType field if non-nil, zero value otherwise.

### GetCustomTaskTypeOk

`func (o *CustomTaskFailedAttributes) GetCustomTaskTypeOk() (*string, bool)`

GetCustomTaskTypeOk returns a tuple with the CustomTaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTaskType

`func (o *CustomTaskFailedAttributes) SetCustomTaskType(v string)`

SetCustomTaskType sets CustomTaskType field to given value.


### GetFailure

`func (o *CustomTaskFailedAttributes) GetFailure() Failure`

GetFailure returns the Failure field if non-nil, zero value otherwise.

### GetFailureOk

`func (o *CustomTaskFailedAttributes) GetFailureOk() (*Failure, bool)`

GetFailureOk returns a tuple with the Failure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailure

`func (o *CustomTaskFailedAttributes) SetFailure(v Failure)`

SetFailure sets Failure field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


