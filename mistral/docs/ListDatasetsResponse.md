# ListDatasetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datasets** | [**PaginatedResultDatasetPreview**](PaginatedResultDatasetPreview.md) |  | 

## Methods

### NewListDatasetsResponse

`func NewListDatasetsResponse(datasets PaginatedResultDatasetPreview, ) *ListDatasetsResponse`

NewListDatasetsResponse instantiates a new ListDatasetsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDatasetsResponseWithDefaults

`func NewListDatasetsResponseWithDefaults() *ListDatasetsResponse`

NewListDatasetsResponseWithDefaults instantiates a new ListDatasetsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasets

`func (o *ListDatasetsResponse) GetDatasets() PaginatedResultDatasetPreview`

GetDatasets returns the Datasets field if non-nil, zero value otherwise.

### GetDatasetsOk

`func (o *ListDatasetsResponse) GetDatasetsOk() (*PaginatedResultDatasetPreview, bool)`

GetDatasetsOk returns a tuple with the Datasets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasets

`func (o *ListDatasetsResponse) SetDatasets(v PaginatedResultDatasetPreview)`

SetDatasets sets Datasets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


