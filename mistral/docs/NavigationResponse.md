# NavigationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chunks** | [**[]NavigationChunkResponse**](NavigationChunkResponse.md) |  | 

## Methods

### NewNavigationResponse

`func NewNavigationResponse(chunks []NavigationChunkResponse, ) *NavigationResponse`

NewNavigationResponse instantiates a new NavigationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNavigationResponseWithDefaults

`func NewNavigationResponseWithDefaults() *NavigationResponse`

NewNavigationResponseWithDefaults instantiates a new NavigationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChunks

`func (o *NavigationResponse) GetChunks() []NavigationChunkResponse`

GetChunks returns the Chunks field if non-nil, zero value otherwise.

### GetChunksOk

`func (o *NavigationResponse) GetChunksOk() (*[]NavigationChunkResponse, bool)`

GetChunksOk returns a tuple with the Chunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunks

`func (o *NavigationResponse) SetChunks(v []NavigationChunkResponse)`

SetChunks sets Chunks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


