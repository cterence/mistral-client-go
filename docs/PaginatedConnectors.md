# PaginatedConnectors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]Connector**](Connector.md) |  | 
**Pagination** | [**PaginationResponse**](PaginationResponse.md) |  | 

## Methods

### NewPaginatedConnectors

`func NewPaginatedConnectors(items []Connector, pagination PaginationResponse, ) *PaginatedConnectors`

NewPaginatedConnectors instantiates a new PaginatedConnectors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedConnectorsWithDefaults

`func NewPaginatedConnectorsWithDefaults() *PaginatedConnectors`

NewPaginatedConnectorsWithDefaults instantiates a new PaginatedConnectors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *PaginatedConnectors) GetItems() []Connector`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PaginatedConnectors) GetItemsOk() (*[]Connector, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PaginatedConnectors) SetItems(v []Connector)`

SetItems sets Items field to given value.


### GetPagination

`func (o *PaginatedConnectors) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedConnectors) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedConnectors) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


