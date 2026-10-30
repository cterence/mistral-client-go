# ImportDatasetFromSpansRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpanReferences** | [**[]TelemetrySpanReference**](TelemetrySpanReference.md) | Spans to attempt to import; unavailable references are skipped, and trace and span IDs are attached as record properties automatically. | 
**MappingContract** | [**SpanDatasetMappingContract**](SpanDatasetMappingContract.md) | Versioned rules for mapping span or resource attributes into top-level payload keys. | 

## Methods

### NewImportDatasetFromSpansRequest

`func NewImportDatasetFromSpansRequest(spanReferences []TelemetrySpanReference, mappingContract SpanDatasetMappingContract, ) *ImportDatasetFromSpansRequest`

NewImportDatasetFromSpansRequest instantiates a new ImportDatasetFromSpansRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportDatasetFromSpansRequestWithDefaults

`func NewImportDatasetFromSpansRequestWithDefaults() *ImportDatasetFromSpansRequest`

NewImportDatasetFromSpansRequestWithDefaults instantiates a new ImportDatasetFromSpansRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpanReferences

`func (o *ImportDatasetFromSpansRequest) GetSpanReferences() []TelemetrySpanReference`

GetSpanReferences returns the SpanReferences field if non-nil, zero value otherwise.

### GetSpanReferencesOk

`func (o *ImportDatasetFromSpansRequest) GetSpanReferencesOk() (*[]TelemetrySpanReference, bool)`

GetSpanReferencesOk returns a tuple with the SpanReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanReferences

`func (o *ImportDatasetFromSpansRequest) SetSpanReferences(v []TelemetrySpanReference)`

SetSpanReferences sets SpanReferences field to given value.


### GetMappingContract

`func (o *ImportDatasetFromSpansRequest) GetMappingContract() SpanDatasetMappingContract`

GetMappingContract returns the MappingContract field if non-nil, zero value otherwise.

### GetMappingContractOk

`func (o *ImportDatasetFromSpansRequest) GetMappingContractOk() (*SpanDatasetMappingContract, bool)`

GetMappingContractOk returns a tuple with the MappingContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingContract

`func (o *ImportDatasetFromSpansRequest) SetMappingContract(v SpanDatasetMappingContract)`

SetMappingContract sets MappingContract field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


