# CustomTaskStartedAttributesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomTaskId** | **string** | Unique identifier for the custom task within the workflow. | 
**CustomTaskType** | **string** | The type/category of the custom task (e.g., &#39;llm_call&#39;, &#39;api_request&#39;). | 
**Payload** | Pointer to [**JSONPayloadResponse**](JSONPayloadResponse.md) | The initial state or payload for the custom task. | [optional] 

## Methods

### NewCustomTaskStartedAttributesResponse

`func NewCustomTaskStartedAttributesResponse(customTaskId string, customTaskType string, ) *CustomTaskStartedAttributesResponse`

NewCustomTaskStartedAttributesResponse instantiates a new CustomTaskStartedAttributesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomTaskStartedAttributesResponseWithDefaults

`func NewCustomTaskStartedAttributesResponseWithDefaults() *CustomTaskStartedAttributesResponse`

NewCustomTaskStartedAttributesResponseWithDefaults instantiates a new CustomTaskStartedAttributesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomTaskId

`func (o *CustomTaskStartedAttributesResponse) GetCustomTaskId() string`

GetCustomTaskId returns the CustomTaskId field if non-nil, zero value otherwise.

### GetCustomTaskIdOk

`func (o *CustomTaskStartedAttributesResponse) GetCustomTaskIdOk() (*string, bool)`

GetCustomTaskIdOk returns a tuple with the CustomTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTaskId

`func (o *CustomTaskStartedAttributesResponse) SetCustomTaskId(v string)`

SetCustomTaskId sets CustomTaskId field to given value.


### GetCustomTaskType

`func (o *CustomTaskStartedAttributesResponse) GetCustomTaskType() string`

GetCustomTaskType returns the CustomTaskType field if non-nil, zero value otherwise.

### GetCustomTaskTypeOk

`func (o *CustomTaskStartedAttributesResponse) GetCustomTaskTypeOk() (*string, bool)`

GetCustomTaskTypeOk returns a tuple with the CustomTaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTaskType

`func (o *CustomTaskStartedAttributesResponse) SetCustomTaskType(v string)`

SetCustomTaskType sets CustomTaskType field to given value.


### GetPayload

`func (o *CustomTaskStartedAttributesResponse) GetPayload() JSONPayloadResponse`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CustomTaskStartedAttributesResponse) GetPayloadOk() (*JSONPayloadResponse, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CustomTaskStartedAttributesResponse) SetPayload(v JSONPayloadResponse)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *CustomTaskStartedAttributesResponse) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


