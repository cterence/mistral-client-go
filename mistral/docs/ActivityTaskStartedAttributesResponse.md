# ActivityTaskStartedAttributesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | **string** | Unique identifier for the activity task within the workflow. | 
**ActivityName** | **string** | The registered name of the activity being executed. | 
**Input** | [**JSONPayloadResponse**](JSONPayloadResponse.md) | The input arguments passed to the activity. | 

## Methods

### NewActivityTaskStartedAttributesResponse

`func NewActivityTaskStartedAttributesResponse(taskId string, activityName string, input JSONPayloadResponse, ) *ActivityTaskStartedAttributesResponse`

NewActivityTaskStartedAttributesResponse instantiates a new ActivityTaskStartedAttributesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityTaskStartedAttributesResponseWithDefaults

`func NewActivityTaskStartedAttributesResponseWithDefaults() *ActivityTaskStartedAttributesResponse`

NewActivityTaskStartedAttributesResponseWithDefaults instantiates a new ActivityTaskStartedAttributesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *ActivityTaskStartedAttributesResponse) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *ActivityTaskStartedAttributesResponse) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *ActivityTaskStartedAttributesResponse) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetActivityName

`func (o *ActivityTaskStartedAttributesResponse) GetActivityName() string`

GetActivityName returns the ActivityName field if non-nil, zero value otherwise.

### GetActivityNameOk

`func (o *ActivityTaskStartedAttributesResponse) GetActivityNameOk() (*string, bool)`

GetActivityNameOk returns a tuple with the ActivityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityName

`func (o *ActivityTaskStartedAttributesResponse) SetActivityName(v string)`

SetActivityName sets ActivityName field to given value.


### GetInput

`func (o *ActivityTaskStartedAttributesResponse) GetInput() JSONPayloadResponse`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ActivityTaskStartedAttributesResponse) GetInputOk() (*JSONPayloadResponse, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ActivityTaskStartedAttributesResponse) SetInput(v JSONPayloadResponse)`

SetInput sets Input field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


