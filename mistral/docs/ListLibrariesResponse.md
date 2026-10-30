# ListLibrariesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**NullablePaginationInfo**](PaginationInfo.md) | Deprecated: offset pagination metadata. Only returned for requests using the deprecated &#x60;page&#x60; parameter; omitted when &#x60;page_token&#x60; is used. Use &#x60;next_page_token&#x60; instead — this field will be removed once offset paging is retired. | [optional] 
**Data** | [**[]Library**](Library.md) |  | 
**NextPageToken** | Pointer to **NullableString** | Opaque continuation token for the next page. Pass it back as &#x60;page_token&#x60; to fetch the next page. Null when there are no more results. Prefer this over the deprecated offset &#x60;page&#x60; parameter. | [optional] 

## Methods

### NewListLibrariesResponse

`func NewListLibrariesResponse(data []Library, ) *ListLibrariesResponse`

NewListLibrariesResponse instantiates a new ListLibrariesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLibrariesResponseWithDefaults

`func NewListLibrariesResponseWithDefaults() *ListLibrariesResponse`

NewListLibrariesResponseWithDefaults instantiates a new ListLibrariesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *ListLibrariesResponse) GetPagination() PaginationInfo`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListLibrariesResponse) GetPaginationOk() (*PaginationInfo, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListLibrariesResponse) SetPagination(v PaginationInfo)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListLibrariesResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### SetPaginationNil

`func (o *ListLibrariesResponse) SetPaginationNil(b bool)`

 SetPaginationNil sets the value for Pagination to be an explicit nil

### UnsetPagination
`func (o *ListLibrariesResponse) UnsetPagination()`

UnsetPagination ensures that no value is present for Pagination, not even an explicit nil
### GetData

`func (o *ListLibrariesResponse) GetData() []Library`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListLibrariesResponse) GetDataOk() (*[]Library, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListLibrariesResponse) SetData(v []Library)`

SetData sets Data field to given value.


### GetNextPageToken

`func (o *ListLibrariesResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListLibrariesResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListLibrariesResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListLibrariesResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageTokenNil

`func (o *ListLibrariesResponse) SetNextPageTokenNil(b bool)`

 SetNextPageTokenNil sets the value for NextPageToken to be an explicit nil

### UnsetNextPageToken
`func (o *ListLibrariesResponse) UnsetNextPageToken()`

UnsetNextPageToken ensures that no value is present for NextPageToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


