# RequiresConfirmation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Expressions** | [**[]ExpressionsInner**](ExpressionsInner.md) |  | 
**ReadOnly** | **bool** |  | 

## Methods

### NewRequiresConfirmation

`func NewRequiresConfirmation(type_ string, expressions []ExpressionsInner, readOnly bool, ) *RequiresConfirmation`

NewRequiresConfirmation instantiates a new RequiresConfirmation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequiresConfirmationWithDefaults

`func NewRequiresConfirmationWithDefaults() *RequiresConfirmation`

NewRequiresConfirmationWithDefaults instantiates a new RequiresConfirmation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RequiresConfirmation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequiresConfirmation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequiresConfirmation) SetType(v string)`

SetType sets Type field to given value.


### GetExpressions

`func (o *RequiresConfirmation) GetExpressions() []ExpressionsInner`

GetExpressions returns the Expressions field if non-nil, zero value otherwise.

### GetExpressionsOk

`func (o *RequiresConfirmation) GetExpressionsOk() (*[]ExpressionsInner, bool)`

GetExpressionsOk returns a tuple with the Expressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressions

`func (o *RequiresConfirmation) SetExpressions(v []ExpressionsInner)`

SetExpressions sets Expressions field to given value.


### GetReadOnly

`func (o *RequiresConfirmation) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *RequiresConfirmation) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *RequiresConfirmation) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


