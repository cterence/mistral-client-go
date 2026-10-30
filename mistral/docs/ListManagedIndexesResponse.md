# ListManagedIndexesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ManagedIndexResponse**](ManagedIndexResponse.md) |  | 
**NextPageToken** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListManagedIndexesResponse

`func NewListManagedIndexesResponse(data []ManagedIndexResponse, ) *ListManagedIndexesResponse`

NewListManagedIndexesResponse instantiates a new ListManagedIndexesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListManagedIndexesResponseWithDefaults

`func NewListManagedIndexesResponseWithDefaults() *ListManagedIndexesResponse`

NewListManagedIndexesResponseWithDefaults instantiates a new ListManagedIndexesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListManagedIndexesResponse) GetData() []ManagedIndexResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListManagedIndexesResponse) GetDataOk() (*[]ManagedIndexResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListManagedIndexesResponse) SetData(v []ManagedIndexResponse)`

SetData sets Data field to given value.


### GetNextPageToken

`func (o *ListManagedIndexesResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListManagedIndexesResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListManagedIndexesResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListManagedIndexesResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageTokenNil

`func (o *ListManagedIndexesResponse) SetNextPageTokenNil(b bool)`

 SetNextPageTokenNil sets the value for NextPageToken to be an explicit nil

### UnsetNextPageToken
`func (o *ListManagedIndexesResponse) UnsetNextPageToken()`

UnsetNextPageToken ensures that no value is present for NextPageToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


