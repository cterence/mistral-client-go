# ActivityTaskCompletedAttributesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | **string** | Unique identifier for the activity task within the workflow. | 
**ActivityName** | **string** | The registered name of the activity being executed. | 
**Result** | [**JSONPayloadResponse**](JSONPayloadResponse.md) | The result returned by the activity. | 

## Methods

### NewActivityTaskCompletedAttributesResponse

`func NewActivityTaskCompletedAttributesResponse(taskId string, activityName string, result JSONPayloadResponse, ) *ActivityTaskCompletedAttributesResponse`

NewActivityTaskCompletedAttributesResponse instantiates a new ActivityTaskCompletedAttributesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityTaskCompletedAttributesResponseWithDefaults

`func NewActivityTaskCompletedAttributesResponseWithDefaults() *ActivityTaskCompletedAttributesResponse`

NewActivityTaskCompletedAttributesResponseWithDefaults instantiates a new ActivityTaskCompletedAttributesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *ActivityTaskCompletedAttributesResponse) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *ActivityTaskCompletedAttributesResponse) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *ActivityTaskCompletedAttributesResponse) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetActivityName

`func (o *ActivityTaskCompletedAttributesResponse) GetActivityName() string`

GetActivityName returns the ActivityName field if non-nil, zero value otherwise.

### GetActivityNameOk

`func (o *ActivityTaskCompletedAttributesResponse) GetActivityNameOk() (*string, bool)`

GetActivityNameOk returns a tuple with the ActivityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityName

`func (o *ActivityTaskCompletedAttributesResponse) SetActivityName(v string)`

SetActivityName sets ActivityName field to given value.


### GetResult

`func (o *ActivityTaskCompletedAttributesResponse) GetResult() JSONPayloadResponse`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ActivityTaskCompletedAttributesResponse) GetResultOk() (*JSONPayloadResponse, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ActivityTaskCompletedAttributesResponse) SetResult(v JSONPayloadResponse)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


