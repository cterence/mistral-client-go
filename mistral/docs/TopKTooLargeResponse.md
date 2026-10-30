# TopKTooLargeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "top_k_too_large"]
**TopK** | **int32** |  | 
**Limit** | **int32** |  | 

## Methods

### NewTopKTooLargeResponse

`func NewTopKTooLargeResponse(topK int32, limit int32, ) *TopKTooLargeResponse`

NewTopKTooLargeResponse instantiates a new TopKTooLargeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopKTooLargeResponseWithDefaults

`func NewTopKTooLargeResponseWithDefaults() *TopKTooLargeResponse`

NewTopKTooLargeResponseWithDefaults instantiates a new TopKTooLargeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TopKTooLargeResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TopKTooLargeResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TopKTooLargeResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TopKTooLargeResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTopK

`func (o *TopKTooLargeResponse) GetTopK() int32`

GetTopK returns the TopK field if non-nil, zero value otherwise.

### GetTopKOk

`func (o *TopKTooLargeResponse) GetTopKOk() (*int32, bool)`

GetTopKOk returns a tuple with the TopK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopK

`func (o *TopKTooLargeResponse) SetTopK(v int32)`

SetTopK sets TopK field to given value.


### GetLimit

`func (o *TopKTooLargeResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TopKTooLargeResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TopKTooLargeResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


