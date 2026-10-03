# CustomTaskInProgressAttributesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomTaskId** | **string** | Unique identifier for the custom task within the workflow. | 
**CustomTaskType** | **string** | The type/category of the custom task (e.g., &#39;llm_call&#39;, &#39;api_request&#39;). | 
**Payload** | [**Payload**](Payload.md) |  | 

## Methods

### NewCustomTaskInProgressAttributesResponse

`func NewCustomTaskInProgressAttributesResponse(customTaskId string, customTaskType string, payload Payload, ) *CustomTaskInProgressAttributesResponse`

NewCustomTaskInProgressAttributesResponse instantiates a new CustomTaskInProgressAttributesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomTaskInProgressAttributesResponseWithDefaults

`func NewCustomTaskInProgressAttributesResponseWithDefaults() *CustomTaskInProgressAttributesResponse`

NewCustomTaskInProgressAttributesResponseWithDefaults instantiates a new CustomTaskInProgressAttributesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomTaskId

`func (o *CustomTaskInProgressAttributesResponse) GetCustomTaskId() string`

GetCustomTaskId returns the CustomTaskId field if non-nil, zero value otherwise.

### GetCustomTaskIdOk

`func (o *CustomTaskInProgressAttributesResponse) GetCustomTaskIdOk() (*string, bool)`

GetCustomTaskIdOk returns a tuple with the CustomTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTaskId

`func (o *CustomTaskInProgressAttributesResponse) SetCustomTaskId(v string)`

SetCustomTaskId sets CustomTaskId field to given value.


### GetCustomTaskType

`func (o *CustomTaskInProgressAttributesResponse) GetCustomTaskType() string`

GetCustomTaskType returns the CustomTaskType field if non-nil, zero value otherwise.

### GetCustomTaskTypeOk

`func (o *CustomTaskInProgressAttributesResponse) GetCustomTaskTypeOk() (*string, bool)`

GetCustomTaskTypeOk returns a tuple with the CustomTaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTaskType

`func (o *CustomTaskInProgressAttributesResponse) SetCustomTaskType(v string)`

SetCustomTaskType sets CustomTaskType field to given value.


### GetPayload

`func (o *CustomTaskInProgressAttributesResponse) GetPayload() Payload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CustomTaskInProgressAttributesResponse) GetPayloadOk() (*Payload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CustomTaskInProgressAttributesResponse) SetPayload(v Payload)`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


