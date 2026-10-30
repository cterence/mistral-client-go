# Document1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "file"]
**FileId** | **string** |  | 
**DocumentUrl** | **string** |  | 
**DocumentName** | Pointer to **string** | The filename of the document | [optional] 
**ImageUrl** | [**ImageUrl**](ImageUrl.md) |  | 

## Methods

### NewDocument1

`func NewDocument1(fileId string, documentUrl string, imageUrl ImageUrl, ) *Document1`

NewDocument1 instantiates a new Document1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocument1WithDefaults

`func NewDocument1WithDefaults() *Document1`

NewDocument1WithDefaults instantiates a new Document1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Document1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Document1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Document1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Document1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFileId

`func (o *Document1) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *Document1) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *Document1) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetDocumentUrl

`func (o *Document1) GetDocumentUrl() string`

GetDocumentUrl returns the DocumentUrl field if non-nil, zero value otherwise.

### GetDocumentUrlOk

`func (o *Document1) GetDocumentUrlOk() (*string, bool)`

GetDocumentUrlOk returns a tuple with the DocumentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentUrl

`func (o *Document1) SetDocumentUrl(v string)`

SetDocumentUrl sets DocumentUrl field to given value.


### GetDocumentName

`func (o *Document1) GetDocumentName() string`

GetDocumentName returns the DocumentName field if non-nil, zero value otherwise.

### GetDocumentNameOk

`func (o *Document1) GetDocumentNameOk() (*string, bool)`

GetDocumentNameOk returns a tuple with the DocumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentName

`func (o *Document1) SetDocumentName(v string)`

SetDocumentName sets DocumentName field to given value.

### HasDocumentName

`func (o *Document1) HasDocumentName() bool`

HasDocumentName returns a boolean if a field has been set.

### GetImageUrl

`func (o *Document1) GetImageUrl() ImageUrl`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *Document1) GetImageUrlOk() (*ImageUrl, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *Document1) SetImageUrl(v ImageUrl)`

SetImageUrl sets ImageUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


