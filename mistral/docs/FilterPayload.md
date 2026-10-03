# FilterPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | [**NullableFilters**](Filters.md) |  | 

## Methods

### NewFilterPayload

`func NewFilterPayload(filters NullableFilters, ) *FilterPayload`

NewFilterPayload instantiates a new FilterPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterPayloadWithDefaults

`func NewFilterPayloadWithDefaults() *FilterPayload`

NewFilterPayloadWithDefaults instantiates a new FilterPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *FilterPayload) GetFilters() Filters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *FilterPayload) GetFiltersOk() (*Filters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *FilterPayload) SetFilters(v Filters)`

SetFilters sets Filters field to given value.


### SetFiltersNil

`func (o *FilterPayload) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *FilterPayload) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


