# IngestDocumentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accepted** | **int32** |  | 
**Rejected** | **int32** |  | 
**Results** | [**[]ResultsInner**](ResultsInner.md) |  | 

## Methods

### NewIngestDocumentsResponse

`func NewIngestDocumentsResponse(accepted int32, rejected int32, results []ResultsInner, ) *IngestDocumentsResponse`

NewIngestDocumentsResponse instantiates a new IngestDocumentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestDocumentsResponseWithDefaults

`func NewIngestDocumentsResponseWithDefaults() *IngestDocumentsResponse`

NewIngestDocumentsResponseWithDefaults instantiates a new IngestDocumentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccepted

`func (o *IngestDocumentsResponse) GetAccepted() int32`

GetAccepted returns the Accepted field if non-nil, zero value otherwise.

### GetAcceptedOk

`func (o *IngestDocumentsResponse) GetAcceptedOk() (*int32, bool)`

GetAcceptedOk returns a tuple with the Accepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepted

`func (o *IngestDocumentsResponse) SetAccepted(v int32)`

SetAccepted sets Accepted field to given value.


### GetRejected

`func (o *IngestDocumentsResponse) GetRejected() int32`

GetRejected returns the Rejected field if non-nil, zero value otherwise.

### GetRejectedOk

`func (o *IngestDocumentsResponse) GetRejectedOk() (*int32, bool)`

GetRejectedOk returns a tuple with the Rejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejected

`func (o *IngestDocumentsResponse) SetRejected(v int32)`

SetRejected sets Rejected field to given value.


### GetResults

`func (o *IngestDocumentsResponse) GetResults() []ResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *IngestDocumentsResponse) GetResultsOk() (*[]ResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *IngestDocumentsResponse) SetResults(v []ResultsInner)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


