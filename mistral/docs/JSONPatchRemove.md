# JSONPatchRemove

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | A JSON Pointer (RFC 6901) identifying the target location within the document. Can be a string path (e.g., &#39;/foo/bar&#39;), &#39;/&#39;, &#39;&#39;, or an empty list [] for root-level operations. | 
**Value** | **interface{}** | The value to use for the operation | 
**Op** | **string** | Remove operation | 

## Methods

### NewJSONPatchRemove

`func NewJSONPatchRemove(path string, value interface{}, op string, ) *JSONPatchRemove`

NewJSONPatchRemove instantiates a new JSONPatchRemove object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSONPatchRemoveWithDefaults

`func NewJSONPatchRemoveWithDefaults() *JSONPatchRemove`

NewJSONPatchRemoveWithDefaults instantiates a new JSONPatchRemove object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *JSONPatchRemove) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *JSONPatchRemove) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *JSONPatchRemove) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *JSONPatchRemove) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *JSONPatchRemove) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *JSONPatchRemove) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *JSONPatchRemove) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *JSONPatchRemove) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetOp

`func (o *JSONPatchRemove) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *JSONPatchRemove) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *JSONPatchRemove) SetOp(v string)`

SetOp sets Op field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


