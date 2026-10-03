# CustomTaskTimedOutAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomTaskId** | **string** | Unique identifier for the custom task within the workflow. | 
**CustomTaskType** | **string** | The type/category of the custom task (e.g., &#39;llm_call&#39;, &#39;api_request&#39;). | 
**TimeoutType** | Pointer to **NullableString** | The type of timeout that occurred. | [optional] 

## Methods

### NewCustomTaskTimedOutAttributes

`func NewCustomTaskTimedOutAttributes(customTaskId string, customTaskType string, ) *CustomTaskTimedOutAttributes`

NewCustomTaskTimedOutAttributes instantiates a new CustomTaskTimedOutAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomTaskTimedOutAttributesWithDefaults

`func NewCustomTaskTimedOutAttributesWithDefaults() *CustomTaskTimedOutAttributes`

NewCustomTaskTimedOutAttributesWithDefaults instantiates a new CustomTaskTimedOutAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomTaskId

`func (o *CustomTaskTimedOutAttributes) GetCustomTaskId() string`

GetCustomTaskId returns the CustomTaskId field if non-nil, zero value otherwise.

### GetCustomTaskIdOk

`func (o *CustomTaskTimedOutAttributes) GetCustomTaskIdOk() (*string, bool)`

GetCustomTaskIdOk returns a tuple with the CustomTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTaskId

`func (o *CustomTaskTimedOutAttributes) SetCustomTaskId(v string)`

SetCustomTaskId sets CustomTaskId field to given value.


### GetCustomTaskType

`func (o *CustomTaskTimedOutAttributes) GetCustomTaskType() string`

GetCustomTaskType returns the CustomTaskType field if non-nil, zero value otherwise.

### GetCustomTaskTypeOk

`func (o *CustomTaskTimedOutAttributes) GetCustomTaskTypeOk() (*string, bool)`

GetCustomTaskTypeOk returns a tuple with the CustomTaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTaskType

`func (o *CustomTaskTimedOutAttributes) SetCustomTaskType(v string)`

SetCustomTaskType sets CustomTaskType field to given value.


### GetTimeoutType

`func (o *CustomTaskTimedOutAttributes) GetTimeoutType() string`

GetTimeoutType returns the TimeoutType field if non-nil, zero value otherwise.

### GetTimeoutTypeOk

`func (o *CustomTaskTimedOutAttributes) GetTimeoutTypeOk() (*string, bool)`

GetTimeoutTypeOk returns a tuple with the TimeoutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutType

`func (o *CustomTaskTimedOutAttributes) SetTimeoutType(v string)`

SetTimeoutType sets TimeoutType field to given value.

### HasTimeoutType

`func (o *CustomTaskTimedOutAttributes) HasTimeoutType() bool`

HasTimeoutType returns a boolean if a field has been set.

### SetTimeoutTypeNil

`func (o *CustomTaskTimedOutAttributes) SetTimeoutTypeNil(b bool)`

 SetTimeoutTypeNil sets the value for TimeoutType to be an explicit nil

### UnsetTimeoutType
`func (o *CustomTaskTimedOutAttributes) UnsetTimeoutType()`

UnsetTimeoutType ensures that no value is present for TimeoutType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


