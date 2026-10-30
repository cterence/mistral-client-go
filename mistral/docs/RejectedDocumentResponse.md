# RejectedDocumentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "rejected"]
**Position** | **int32** |  | 
**Errors** | [**[]DocumentFieldErrorResponse**](DocumentFieldErrorResponse.md) |  | 

## Methods

### NewRejectedDocumentResponse

`func NewRejectedDocumentResponse(position int32, errors []DocumentFieldErrorResponse, ) *RejectedDocumentResponse`

NewRejectedDocumentResponse instantiates a new RejectedDocumentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRejectedDocumentResponseWithDefaults

`func NewRejectedDocumentResponseWithDefaults() *RejectedDocumentResponse`

NewRejectedDocumentResponseWithDefaults instantiates a new RejectedDocumentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RejectedDocumentResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RejectedDocumentResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RejectedDocumentResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RejectedDocumentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPosition

`func (o *RejectedDocumentResponse) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *RejectedDocumentResponse) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *RejectedDocumentResponse) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetErrors

`func (o *RejectedDocumentResponse) GetErrors() []DocumentFieldErrorResponse`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *RejectedDocumentResponse) GetErrorsOk() (*[]DocumentFieldErrorResponse, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *RejectedDocumentResponse) SetErrors(v []DocumentFieldErrorResponse)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


