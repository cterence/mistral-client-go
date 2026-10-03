# ListWorkflowEventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]EventsInner**](EventsInner.md) | List of workflow events. | 
**NextCursor** | Pointer to **NullableString** | Cursor for pagination. | [optional] 

## Methods

### NewListWorkflowEventResponse

`func NewListWorkflowEventResponse(events []EventsInner, ) *ListWorkflowEventResponse`

NewListWorkflowEventResponse instantiates a new ListWorkflowEventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorkflowEventResponseWithDefaults

`func NewListWorkflowEventResponseWithDefaults() *ListWorkflowEventResponse`

NewListWorkflowEventResponseWithDefaults instantiates a new ListWorkflowEventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *ListWorkflowEventResponse) GetEvents() []EventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ListWorkflowEventResponse) GetEventsOk() (*[]EventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ListWorkflowEventResponse) SetEvents(v []EventsInner)`

SetEvents sets Events field to given value.


### GetNextCursor

`func (o *ListWorkflowEventResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ListWorkflowEventResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ListWorkflowEventResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ListWorkflowEventResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *ListWorkflowEventResponse) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *ListWorkflowEventResponse) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


