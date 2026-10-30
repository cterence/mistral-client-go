# Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "missing_field"]
**Name** | **string** |  | 
**ExpectedType** | **string** |  | 
**ReceivedType** | **string** |  | 
**Expected** | **int32** |  | 
**Received** | **int32** |  | 
**MaxCandidates** | **int32** |  | 
**TopK** | **int32** |  | 
**Limit** | **int32** |  | 

## Methods

### NewResponse400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner

`func NewResponse400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner(name string, expectedType string, receivedType string, expected int32, received int32, maxCandidates int32, topK int32, limit int32, ) *Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner`

NewResponse400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner instantiates a new Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponse400SearchIndexV1RagManagedIndexesIndexNameSearchPostInnerWithDefaults

`func NewResponse400SearchIndexV1RagManagedIndexesIndexNameSearchPostInnerWithDefaults() *Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner`

NewResponse400SearchIndexV1RagManagedIndexesIndexNameSearchPostInnerWithDefaults instantiates a new Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner) SetName(v string)`

SetName sets Name field to given value.


### GetExpectedType

`func (o *Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner) GetExpectedType() string`

GetExpectedType returns the ExpectedType field if non-nil, zero value otherwise.

### GetExpectedTypeOk

`func (o *Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner) GetExpectedTypeOk() (*string, bool)`

GetExpectedTypeOk returns a tuple with the ExpectedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedType

`func (o *Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner) SetExpectedType(v string)`

SetExpectedType sets ExpectedType field to given value.


### GetReceivedType

`func (o *Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner) GetReceivedType() string`

GetReceivedType returns the ReceivedType field if non-nil, zero value otherwise.

### GetReceivedTypeOk

`func (o *Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner) GetReceivedTypeOk() (*string, bool)`

GetReceivedTypeOk returns a tuple with the ReceivedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedType

`func (o *Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner) SetReceivedType(v string)`

SetReceivedType sets ReceivedType field to given value.


### GetExpected

`func (o *Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner) GetExpected() int32`

GetExpected returns the Expected field if non-nil, zero value otherwise.

### GetExpectedOk

`func (o *Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner) GetExpectedOk() (*int32, bool)`

GetExpectedOk returns a tuple with the Expected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpected

`func (o *Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner) SetExpected(v int32)`

SetExpected sets Expected field to given value.


### GetReceived

`func (o *Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner) GetReceived() int32`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner) GetReceivedOk() (*int32, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner) SetReceived(v int32)`

SetReceived sets Received field to given value.


### GetMaxCandidates

`func (o *Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner) GetMaxCandidates() int32`

GetMaxCandidates returns the MaxCandidates field if non-nil, zero value otherwise.

### GetMaxCandidatesOk

`func (o *Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner) GetMaxCandidatesOk() (*int32, bool)`

GetMaxCandidatesOk returns a tuple with the MaxCandidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCandidates

`func (o *Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner) SetMaxCandidates(v int32)`

SetMaxCandidates sets MaxCandidates field to given value.


### GetTopK

`func (o *Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner) GetTopK() int32`

GetTopK returns the TopK field if non-nil, zero value otherwise.

### GetTopKOk

`func (o *Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner) GetTopKOk() (*int32, bool)`

GetTopKOk returns a tuple with the TopK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopK

`func (o *Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner) SetTopK(v int32)`

SetTopK sets TopK field to given value.


### GetLimit

`func (o *Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Response400SearchIndexV1RagManagedIndexesIndexNameSearchPostInner) SetLimit(v int32)`

SetLimit sets Limit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


