# ActivityTaskRetryingAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | **string** | Unique identifier for the activity task within the workflow. | 
**ActivityName** | **string** | The registered name of the activity being executed. | 
**Attempt** | **int32** | The attempt number that failed (1-indexed). | 
**Failure** | [**Failure**](Failure.md) | Details about the failure that caused the retry. | 

## Methods

### NewActivityTaskRetryingAttributes

`func NewActivityTaskRetryingAttributes(taskId string, activityName string, attempt int32, failure Failure, ) *ActivityTaskRetryingAttributes`

NewActivityTaskRetryingAttributes instantiates a new ActivityTaskRetryingAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityTaskRetryingAttributesWithDefaults

`func NewActivityTaskRetryingAttributesWithDefaults() *ActivityTaskRetryingAttributes`

NewActivityTaskRetryingAttributesWithDefaults instantiates a new ActivityTaskRetryingAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *ActivityTaskRetryingAttributes) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *ActivityTaskRetryingAttributes) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *ActivityTaskRetryingAttributes) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetActivityName

`func (o *ActivityTaskRetryingAttributes) GetActivityName() string`

GetActivityName returns the ActivityName field if non-nil, zero value otherwise.

### GetActivityNameOk

`func (o *ActivityTaskRetryingAttributes) GetActivityNameOk() (*string, bool)`

GetActivityNameOk returns a tuple with the ActivityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityName

`func (o *ActivityTaskRetryingAttributes) SetActivityName(v string)`

SetActivityName sets ActivityName field to given value.


### GetAttempt

`func (o *ActivityTaskRetryingAttributes) GetAttempt() int32`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *ActivityTaskRetryingAttributes) GetAttemptOk() (*int32, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *ActivityTaskRetryingAttributes) SetAttempt(v int32)`

SetAttempt sets Attempt field to given value.


### GetFailure

`func (o *ActivityTaskRetryingAttributes) GetFailure() Failure`

GetFailure returns the Failure field if non-nil, zero value otherwise.

### GetFailureOk

`func (o *ActivityTaskRetryingAttributes) GetFailureOk() (*Failure, bool)`

GetFailureOk returns a tuple with the Failure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailure

`func (o *ActivityTaskRetryingAttributes) SetFailure(v Failure)`

SetFailure sets Failure field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


