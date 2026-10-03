# AgentListPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] [default to "list"]
**Data** | [**[]Agent**](Agent.md) |  | 
**NextPageToken** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAgentListPage

`func NewAgentListPage(data []Agent, ) *AgentListPage`

NewAgentListPage instantiates a new AgentListPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentListPageWithDefaults

`func NewAgentListPageWithDefaults() *AgentListPage`

NewAgentListPageWithDefaults instantiates a new AgentListPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *AgentListPage) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AgentListPage) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AgentListPage) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *AgentListPage) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetData

`func (o *AgentListPage) GetData() []Agent`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AgentListPage) GetDataOk() (*[]Agent, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AgentListPage) SetData(v []Agent)`

SetData sets Data field to given value.


### GetNextPageToken

`func (o *AgentListPage) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *AgentListPage) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *AgentListPage) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *AgentListPage) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageTokenNil

`func (o *AgentListPage) SetNextPageTokenNil(b bool)`

 SetNextPageTokenNil sets the value for NextPageToken to be an explicit nil

### UnsetNextPageToken
`func (o *AgentListPage) UnsetNextPageToken()`

UnsetNextPageToken ensures that no value is present for NextPageToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


