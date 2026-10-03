# CustomTaskCompletedAttributesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomTaskId** | **string** | Unique identifier for the custom task within the workflow. | 
**CustomTaskType** | **string** | The type/category of the custom task (e.g., &#39;llm_call&#39;, &#39;api_request&#39;). | 
**Payload** | [**JSONPayloadResponse**](JSONPayloadResponse.md) | The final result of the custom task. | 

## Methods

### NewCustomTaskCompletedAttributesResponse

`func NewCustomTaskCompletedAttributesResponse(customTaskId string, customTaskType string, payload JSONPayloadResponse, ) *CustomTaskCompletedAttributesResponse`

NewCustomTaskCompletedAttributesResponse instantiates a new CustomTaskCompletedAttributesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomTaskCompletedAttributesResponseWithDefaults

`func NewCustomTaskCompletedAttributesResponseWithDefaults() *CustomTaskCompletedAttributesResponse`

NewCustomTaskCompletedAttributesResponseWithDefaults instantiates a new CustomTaskCompletedAttributesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomTaskId

`func (o *CustomTaskCompletedAttributesResponse) GetCustomTaskId() string`

GetCustomTaskId returns the CustomTaskId field if non-nil, zero value otherwise.

### GetCustomTaskIdOk

`func (o *CustomTaskCompletedAttributesResponse) GetCustomTaskIdOk() (*string, bool)`

GetCustomTaskIdOk returns a tuple with the CustomTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTaskId

`func (o *CustomTaskCompletedAttributesResponse) SetCustomTaskId(v string)`

SetCustomTaskId sets CustomTaskId field to given value.


### GetCustomTaskType

`func (o *CustomTaskCompletedAttributesResponse) GetCustomTaskType() string`

GetCustomTaskType returns the CustomTaskType field if non-nil, zero value otherwise.

### GetCustomTaskTypeOk

`func (o *CustomTaskCompletedAttributesResponse) GetCustomTaskTypeOk() (*string, bool)`

GetCustomTaskTypeOk returns a tuple with the CustomTaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTaskType

`func (o *CustomTaskCompletedAttributesResponse) SetCustomTaskType(v string)`

SetCustomTaskType sets CustomTaskType field to given value.


### GetPayload

`func (o *CustomTaskCompletedAttributesResponse) GetPayload() JSONPayloadResponse`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CustomTaskCompletedAttributesResponse) GetPayloadOk() (*JSONPayloadResponse, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CustomTaskCompletedAttributesResponse) SetPayload(v JSONPayloadResponse)`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


