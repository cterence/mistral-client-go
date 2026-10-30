# UsageQueryMetadataOUT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryParams** | [**UsageQueryParamsOUT**](UsageQueryParamsOUT.md) | Normalized parameters used by the usage service. | 
**HasMore** | **bool** | Whether another page of grouped rows is available. | 

## Methods

### NewUsageQueryMetadataOUT

`func NewUsageQueryMetadataOUT(queryParams UsageQueryParamsOUT, hasMore bool, ) *UsageQueryMetadataOUT`

NewUsageQueryMetadataOUT instantiates a new UsageQueryMetadataOUT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageQueryMetadataOUTWithDefaults

`func NewUsageQueryMetadataOUTWithDefaults() *UsageQueryMetadataOUT`

NewUsageQueryMetadataOUTWithDefaults instantiates a new UsageQueryMetadataOUT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryParams

`func (o *UsageQueryMetadataOUT) GetQueryParams() UsageQueryParamsOUT`

GetQueryParams returns the QueryParams field if non-nil, zero value otherwise.

### GetQueryParamsOk

`func (o *UsageQueryMetadataOUT) GetQueryParamsOk() (*UsageQueryParamsOUT, bool)`

GetQueryParamsOk returns a tuple with the QueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParams

`func (o *UsageQueryMetadataOUT) SetQueryParams(v UsageQueryParamsOUT)`

SetQueryParams sets QueryParams field to given value.


### GetHasMore

`func (o *UsageQueryMetadataOUT) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *UsageQueryMetadataOUT) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *UsageQueryMetadataOUT) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


