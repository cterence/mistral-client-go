# DocumentFieldErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | [**[]LocationInner**](LocationInner.md) |  | 
**Type** | **string** |  | 
**Message** | **string** |  | 

## Methods

### NewDocumentFieldErrorResponse

`func NewDocumentFieldErrorResponse(location []LocationInner, type_ string, message string, ) *DocumentFieldErrorResponse`

NewDocumentFieldErrorResponse instantiates a new DocumentFieldErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentFieldErrorResponseWithDefaults

`func NewDocumentFieldErrorResponseWithDefaults() *DocumentFieldErrorResponse`

NewDocumentFieldErrorResponseWithDefaults instantiates a new DocumentFieldErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *DocumentFieldErrorResponse) GetLocation() []LocationInner`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DocumentFieldErrorResponse) GetLocationOk() (*[]LocationInner, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DocumentFieldErrorResponse) SetLocation(v []LocationInner)`

SetLocation sets Location field to given value.


### GetType

`func (o *DocumentFieldErrorResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DocumentFieldErrorResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DocumentFieldErrorResponse) SetType(v string)`

SetType sets Type field to given value.


### GetMessage

`func (o *DocumentFieldErrorResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DocumentFieldErrorResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DocumentFieldErrorResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


