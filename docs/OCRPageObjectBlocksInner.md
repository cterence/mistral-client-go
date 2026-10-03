# OCRPageObjectBlocksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopLeftX** | **int32** |  | 
**TopLeftY** | **int32** |  | 
**BottomRightX** | **int32** |  | 
**BottomRightY** | **int32** |  | 
**Content** | **string** | Text/markdown/html content of this block | 
**Type** | Pointer to **string** |  | [optional] [default to "text"]
**ImageId** | **string** | References the corresponding entry in OCRPageObject.images | 
**TableId** | Pointer to **string** | References the corresponding entry in OCRPageObject.tables, when tables are extracted | [optional] 

## Methods

### NewOCRPageObjectBlocksInner

`func NewOCRPageObjectBlocksInner(topLeftX int32, topLeftY int32, bottomRightX int32, bottomRightY int32, content string, imageId string, ) *OCRPageObjectBlocksInner`

NewOCRPageObjectBlocksInner instantiates a new OCRPageObjectBlocksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCRPageObjectBlocksInnerWithDefaults

`func NewOCRPageObjectBlocksInnerWithDefaults() *OCRPageObjectBlocksInner`

NewOCRPageObjectBlocksInnerWithDefaults instantiates a new OCRPageObjectBlocksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopLeftX

`func (o *OCRPageObjectBlocksInner) GetTopLeftX() int32`

GetTopLeftX returns the TopLeftX field if non-nil, zero value otherwise.

### GetTopLeftXOk

`func (o *OCRPageObjectBlocksInner) GetTopLeftXOk() (*int32, bool)`

GetTopLeftXOk returns a tuple with the TopLeftX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLeftX

`func (o *OCRPageObjectBlocksInner) SetTopLeftX(v int32)`

SetTopLeftX sets TopLeftX field to given value.


### GetTopLeftY

`func (o *OCRPageObjectBlocksInner) GetTopLeftY() int32`

GetTopLeftY returns the TopLeftY field if non-nil, zero value otherwise.

### GetTopLeftYOk

`func (o *OCRPageObjectBlocksInner) GetTopLeftYOk() (*int32, bool)`

GetTopLeftYOk returns a tuple with the TopLeftY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLeftY

`func (o *OCRPageObjectBlocksInner) SetTopLeftY(v int32)`

SetTopLeftY sets TopLeftY field to given value.


### GetBottomRightX

`func (o *OCRPageObjectBlocksInner) GetBottomRightX() int32`

GetBottomRightX returns the BottomRightX field if non-nil, zero value otherwise.

### GetBottomRightXOk

`func (o *OCRPageObjectBlocksInner) GetBottomRightXOk() (*int32, bool)`

GetBottomRightXOk returns a tuple with the BottomRightX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomRightX

`func (o *OCRPageObjectBlocksInner) SetBottomRightX(v int32)`

SetBottomRightX sets BottomRightX field to given value.


### GetBottomRightY

`func (o *OCRPageObjectBlocksInner) GetBottomRightY() int32`

GetBottomRightY returns the BottomRightY field if non-nil, zero value otherwise.

### GetBottomRightYOk

`func (o *OCRPageObjectBlocksInner) GetBottomRightYOk() (*int32, bool)`

GetBottomRightYOk returns a tuple with the BottomRightY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomRightY

`func (o *OCRPageObjectBlocksInner) SetBottomRightY(v int32)`

SetBottomRightY sets BottomRightY field to given value.


### GetContent

`func (o *OCRPageObjectBlocksInner) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *OCRPageObjectBlocksInner) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *OCRPageObjectBlocksInner) SetContent(v string)`

SetContent sets Content field to given value.


### GetType

`func (o *OCRPageObjectBlocksInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OCRPageObjectBlocksInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OCRPageObjectBlocksInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OCRPageObjectBlocksInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetImageId

`func (o *OCRPageObjectBlocksInner) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *OCRPageObjectBlocksInner) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *OCRPageObjectBlocksInner) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetTableId

`func (o *OCRPageObjectBlocksInner) GetTableId() string`

GetTableId returns the TableId field if non-nil, zero value otherwise.

### GetTableIdOk

`func (o *OCRPageObjectBlocksInner) GetTableIdOk() (*string, bool)`

GetTableIdOk returns a tuple with the TableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableId

`func (o *OCRPageObjectBlocksInner) SetTableId(v string)`

SetTableId sets TableId field to given value.

### HasTableId

`func (o *OCRPageObjectBlocksInner) HasTableId() bool`

HasTableId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


