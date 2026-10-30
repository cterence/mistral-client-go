# UpdateDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to [**map[string]UpdateDocumentRequestAttributesValue**](UpdateDocumentRequestAttributesValue.md) |  | [optional] 
**ExpiresAt** | Pointer to **NullableTime** | If set, the document will be automatically deleted after this date. | [optional] 

## Methods

### NewUpdateDocumentRequest

`func NewUpdateDocumentRequest() *UpdateDocumentRequest`

NewUpdateDocumentRequest instantiates a new UpdateDocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDocumentRequestWithDefaults

`func NewUpdateDocumentRequestWithDefaults() *UpdateDocumentRequest`

NewUpdateDocumentRequestWithDefaults instantiates a new UpdateDocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateDocumentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDocumentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDocumentRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateDocumentRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAttributes

`func (o *UpdateDocumentRequest) GetAttributes() map[string]UpdateDocumentRequestAttributesValue`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateDocumentRequest) GetAttributesOk() (*map[string]UpdateDocumentRequestAttributesValue, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UpdateDocumentRequest) SetAttributes(v map[string]UpdateDocumentRequestAttributesValue)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UpdateDocumentRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *UpdateDocumentRequest) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *UpdateDocumentRequest) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetExpiresAt

`func (o *UpdateDocumentRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *UpdateDocumentRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *UpdateDocumentRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *UpdateDocumentRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *UpdateDocumentRequest) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *UpdateDocumentRequest) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


