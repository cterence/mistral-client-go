# SpanDatasetSourceReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | [**SpanDatasetSourceNamespace**](SpanDatasetSourceNamespace.md) |  | 
**Key** | **string** | Literal attribute dictionary key; dots do not denote a nested path. | 

## Methods

### NewSpanDatasetSourceReference

`func NewSpanDatasetSourceReference(namespace SpanDatasetSourceNamespace, key string, ) *SpanDatasetSourceReference`

NewSpanDatasetSourceReference instantiates a new SpanDatasetSourceReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanDatasetSourceReferenceWithDefaults

`func NewSpanDatasetSourceReferenceWithDefaults() *SpanDatasetSourceReference`

NewSpanDatasetSourceReferenceWithDefaults instantiates a new SpanDatasetSourceReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *SpanDatasetSourceReference) GetNamespace() SpanDatasetSourceNamespace`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *SpanDatasetSourceReference) GetNamespaceOk() (*SpanDatasetSourceNamespace, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *SpanDatasetSourceReference) SetNamespace(v SpanDatasetSourceNamespace)`

SetNamespace sets Namespace field to given value.


### GetKey

`func (o *SpanDatasetSourceReference) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SpanDatasetSourceReference) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SpanDatasetSourceReference) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


