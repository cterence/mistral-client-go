# OCRListBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopLeftX** | **int32** |  | 
**TopLeftY** | **int32** |  | 
**BottomRightX** | **int32** |  | 
**BottomRightY** | **int32** |  | 
**Content** | **string** | Text/markdown/html content of this block | 
**Type** | Pointer to **string** |  | [optional] [default to "list"]

## Methods

### NewOCRListBlock

`func NewOCRListBlock(topLeftX int32, topLeftY int32, bottomRightX int32, bottomRightY int32, content string, ) *OCRListBlock`

NewOCRListBlock instantiates a new OCRListBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCRListBlockWithDefaults

`func NewOCRListBlockWithDefaults() *OCRListBlock`

NewOCRListBlockWithDefaults instantiates a new OCRListBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopLeftX

`func (o *OCRListBlock) GetTopLeftX() int32`

GetTopLeftX returns the TopLeftX field if non-nil, zero value otherwise.

### GetTopLeftXOk

`func (o *OCRListBlock) GetTopLeftXOk() (*int32, bool)`

GetTopLeftXOk returns a tuple with the TopLeftX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLeftX

`func (o *OCRListBlock) SetTopLeftX(v int32)`

SetTopLeftX sets TopLeftX field to given value.


### GetTopLeftY

`func (o *OCRListBlock) GetTopLeftY() int32`

GetTopLeftY returns the TopLeftY field if non-nil, zero value otherwise.

### GetTopLeftYOk

`func (o *OCRListBlock) GetTopLeftYOk() (*int32, bool)`

GetTopLeftYOk returns a tuple with the TopLeftY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLeftY

`func (o *OCRListBlock) SetTopLeftY(v int32)`

SetTopLeftY sets TopLeftY field to given value.


### GetBottomRightX

`func (o *OCRListBlock) GetBottomRightX() int32`

GetBottomRightX returns the BottomRightX field if non-nil, zero value otherwise.

### GetBottomRightXOk

`func (o *OCRListBlock) GetBottomRightXOk() (*int32, bool)`

GetBottomRightXOk returns a tuple with the BottomRightX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomRightX

`func (o *OCRListBlock) SetBottomRightX(v int32)`

SetBottomRightX sets BottomRightX field to given value.


### GetBottomRightY

`func (o *OCRListBlock) GetBottomRightY() int32`

GetBottomRightY returns the BottomRightY field if non-nil, zero value otherwise.

### GetBottomRightYOk

`func (o *OCRListBlock) GetBottomRightYOk() (*int32, bool)`

GetBottomRightYOk returns a tuple with the BottomRightY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomRightY

`func (o *OCRListBlock) SetBottomRightY(v int32)`

SetBottomRightY sets BottomRightY field to given value.


### GetContent

`func (o *OCRListBlock) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *OCRListBlock) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *OCRListBlock) SetContent(v string)`

SetContent sets Content field to given value.


### GetType

`func (o *OCRListBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OCRListBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OCRListBlock) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OCRListBlock) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


