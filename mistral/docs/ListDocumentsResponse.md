# ListDocumentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**PaginationInfo**](PaginationInfo.md) |  | 
**Data** | [**[]Document**](Document.md) |  | 

## Methods

### NewListDocumentsResponse

`func NewListDocumentsResponse(pagination PaginationInfo, data []Document, ) *ListDocumentsResponse`

NewListDocumentsResponse instantiates a new ListDocumentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDocumentsResponseWithDefaults

`func NewListDocumentsResponseWithDefaults() *ListDocumentsResponse`

NewListDocumentsResponseWithDefaults instantiates a new ListDocumentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *ListDocumentsResponse) GetPagination() PaginationInfo`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListDocumentsResponse) GetPaginationOk() (*PaginationInfo, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListDocumentsResponse) SetPagination(v PaginationInfo)`

SetPagination sets Pagination field to given value.


### GetData

`func (o *ListDocumentsResponse) GetData() []Document`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListDocumentsResponse) GetDataOk() (*[]Document, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListDocumentsResponse) SetData(v []Document)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


