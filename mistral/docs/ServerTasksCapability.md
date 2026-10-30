# ServerTasksCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**List** | Pointer to **map[string]interface{}** | Capability for tasks listing operations (2025-11-25 only). | [optional] 
**Cancel** | Pointer to **map[string]interface{}** | Capability for tasks cancel operations (2025-11-25 only). | [optional] 
**Requests** | Pointer to [**NullableServerTasksRequestsCapability**](ServerTasksRequestsCapability.md) |  | [optional] 

## Methods

### NewServerTasksCapability

`func NewServerTasksCapability() *ServerTasksCapability`

NewServerTasksCapability instantiates a new ServerTasksCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerTasksCapabilityWithDefaults

`func NewServerTasksCapabilityWithDefaults() *ServerTasksCapability`

NewServerTasksCapabilityWithDefaults instantiates a new ServerTasksCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *ServerTasksCapability) GetList() map[string]interface{}`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *ServerTasksCapability) GetListOk() (*map[string]interface{}, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *ServerTasksCapability) SetList(v map[string]interface{})`

SetList sets List field to given value.

### HasList

`func (o *ServerTasksCapability) HasList() bool`

HasList returns a boolean if a field has been set.

### SetListNil

`func (o *ServerTasksCapability) SetListNil(b bool)`

 SetListNil sets the value for List to be an explicit nil

### UnsetList
`func (o *ServerTasksCapability) UnsetList()`

UnsetList ensures that no value is present for List, not even an explicit nil
### GetCancel

`func (o *ServerTasksCapability) GetCancel() map[string]interface{}`

GetCancel returns the Cancel field if non-nil, zero value otherwise.

### GetCancelOk

`func (o *ServerTasksCapability) GetCancelOk() (*map[string]interface{}, bool)`

GetCancelOk returns a tuple with the Cancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancel

`func (o *ServerTasksCapability) SetCancel(v map[string]interface{})`

SetCancel sets Cancel field to given value.

### HasCancel

`func (o *ServerTasksCapability) HasCancel() bool`

HasCancel returns a boolean if a field has been set.

### SetCancelNil

`func (o *ServerTasksCapability) SetCancelNil(b bool)`

 SetCancelNil sets the value for Cancel to be an explicit nil

### UnsetCancel
`func (o *ServerTasksCapability) UnsetCancel()`

UnsetCancel ensures that no value is present for Cancel, not even an explicit nil
### GetRequests

`func (o *ServerTasksCapability) GetRequests() ServerTasksRequestsCapability`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *ServerTasksCapability) GetRequestsOk() (*ServerTasksRequestsCapability, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *ServerTasksCapability) SetRequests(v ServerTasksRequestsCapability)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *ServerTasksCapability) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### SetRequestsNil

`func (o *ServerTasksCapability) SetRequestsNil(b bool)`

 SetRequestsNil sets the value for Requests to be an explicit nil

### UnsetRequests
`func (o *ServerTasksCapability) UnsetRequests()`

UnsetRequests ensures that no value is present for Requests, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


