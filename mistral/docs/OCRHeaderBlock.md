# OCRHeaderBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopLeftX** | **int32** |  | 
**TopLeftY** | **int32** |  | 
**BottomRightX** | **int32** |  | 
**BottomRightY** | **int32** |  | 
**Content** | **string** | Text/markdown/html content of this block | 
**Type** | Pointer to **string** |  | [optional] [default to "header"]

## Methods

### NewOCRHeaderBlock

`func NewOCRHeaderBlock(topLeftX int32, topLeftY int32, bottomRightX int32, bottomRightY int32, content string, ) *OCRHeaderBlock`

NewOCRHeaderBlock instantiates a new OCRHeaderBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCRHeaderBlockWithDefaults

`func NewOCRHeaderBlockWithDefaults() *OCRHeaderBlock`

NewOCRHeaderBlockWithDefaults instantiates a new OCRHeaderBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopLeftX

`func (o *OCRHeaderBlock) GetTopLeftX() int32`

GetTopLeftX returns the TopLeftX field if non-nil, zero value otherwise.

### GetTopLeftXOk

`func (o *OCRHeaderBlock) GetTopLeftXOk() (*int32, bool)`

GetTopLeftXOk returns a tuple with the TopLeftX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLeftX

`func (o *OCRHeaderBlock) SetTopLeftX(v int32)`

SetTopLeftX sets TopLeftX field to given value.


### GetTopLeftY

`func (o *OCRHeaderBlock) GetTopLeftY() int32`

GetTopLeftY returns the TopLeftY field if non-nil, zero value otherwise.

### GetTopLeftYOk

`func (o *OCRHeaderBlock) GetTopLeftYOk() (*int32, bool)`

GetTopLeftYOk returns a tuple with the TopLeftY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLeftY

`func (o *OCRHeaderBlock) SetTopLeftY(v int32)`

SetTopLeftY sets TopLeftY field to given value.


### GetBottomRightX

`func (o *OCRHeaderBlock) GetBottomRightX() int32`

GetBottomRightX returns the BottomRightX field if non-nil, zero value otherwise.

### GetBottomRightXOk

`func (o *OCRHeaderBlock) GetBottomRightXOk() (*int32, bool)`

GetBottomRightXOk returns a tuple with the BottomRightX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomRightX

`func (o *OCRHeaderBlock) SetBottomRightX(v int32)`

SetBottomRightX sets BottomRightX field to given value.


### GetBottomRightY

`func (o *OCRHeaderBlock) GetBottomRightY() int32`

GetBottomRightY returns the BottomRightY field if non-nil, zero value otherwise.

### GetBottomRightYOk

`func (o *OCRHeaderBlock) GetBottomRightYOk() (*int32, bool)`

GetBottomRightYOk returns a tuple with the BottomRightY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomRightY

`func (o *OCRHeaderBlock) SetBottomRightY(v int32)`

SetBottomRightY sets BottomRightY field to given value.


### GetContent

`func (o *OCRHeaderBlock) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *OCRHeaderBlock) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *OCRHeaderBlock) SetContent(v string)`

SetContent sets Content field to given value.


### GetType

`func (o *OCRHeaderBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OCRHeaderBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OCRHeaderBlock) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OCRHeaderBlock) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


