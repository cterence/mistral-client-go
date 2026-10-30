# ListPromptsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Prompt**](Prompt.md) |  | [optional] 
**NextPageToken** | Pointer to **string** |  | [optional] 

## Methods

### NewListPromptsResponse

`func NewListPromptsResponse() *ListPromptsResponse`

NewListPromptsResponse instantiates a new ListPromptsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPromptsResponseWithDefaults

`func NewListPromptsResponseWithDefaults() *ListPromptsResponse`

NewListPromptsResponseWithDefaults instantiates a new ListPromptsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListPromptsResponse) GetData() []Prompt`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListPromptsResponse) GetDataOk() (*[]Prompt, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListPromptsResponse) SetData(v []Prompt)`

SetData sets Data field to given value.

### HasData

`func (o *ListPromptsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNextPageToken

`func (o *ListPromptsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListPromptsResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListPromptsResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListPromptsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


