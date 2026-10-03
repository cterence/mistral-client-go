# JSONPatchAppend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | A JSON Pointer (RFC 6901) identifying the target location within the document. Can be a string path (e.g., &#39;/foo/bar&#39;), &#39;/&#39;, &#39;&#39;, or an empty list [] for root-level operations. | 
**Value** | **string** | The value to use for the operation. A string to append to the existing value | 
**Op** | **string** | &#39;append&#39; is an extension for efficient string concatenation in streaming scenarios. | 

## Methods

### NewJSONPatchAppend

`func NewJSONPatchAppend(path string, value string, op string, ) *JSONPatchAppend`

NewJSONPatchAppend instantiates a new JSONPatchAppend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSONPatchAppendWithDefaults

`func NewJSONPatchAppendWithDefaults() *JSONPatchAppend`

NewJSONPatchAppendWithDefaults instantiates a new JSONPatchAppend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *JSONPatchAppend) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *JSONPatchAppend) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *JSONPatchAppend) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *JSONPatchAppend) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *JSONPatchAppend) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *JSONPatchAppend) SetValue(v string)`

SetValue sets Value field to given value.


### GetOp

`func (o *JSONPatchAppend) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *JSONPatchAppend) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *JSONPatchAppend) SetOp(v string)`

SetOp sets Op field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


