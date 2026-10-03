# FileSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the file. | 
**Object** | **string** | The object type, which is always \&quot;file\&quot;. | 
**Bytes** | **int32** | The size of the file, in bytes. | 
**CreatedAt** | **int32** | The UNIX timestamp (in seconds) of the event. | 
**Filename** | **string** | The name of the uploaded file. | 
**Purpose** | [**FilePurpose**](FilePurpose.md) | The intended purpose of the uploaded file, currently supports fine-tuning (&#x60;fine-tune&#x60;), OCR (&#x60;ocr&#x60;), Audio/Transcription (&#x60;audio&#x60;) and batch inference (&#x60;batch&#x60;). | 
**SampleType** | [**SampleType**](SampleType.md) |  | 
**NumLines** | Pointer to **NullableInt32** |  | [optional] 
**Mimetype** | Pointer to **NullableString** |  | [optional] 
**Source** | [**Source**](Source.md) |  | 
**Signature** | Pointer to **NullableString** |  | [optional] 
**ExpiresAt** | Pointer to **NullableInt32** |  | [optional] 
**Visibility** | Pointer to [**NullableFileVisibility**](FileVisibility.md) |  | [optional] 

## Methods

### NewFileSchema

`func NewFileSchema(id string, object string, bytes int32, createdAt int32, filename string, purpose FilePurpose, sampleType SampleType, source Source, ) *FileSchema`

NewFileSchema instantiates a new FileSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileSchemaWithDefaults

`func NewFileSchemaWithDefaults() *FileSchema`

NewFileSchemaWithDefaults instantiates a new FileSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FileSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileSchema) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *FileSchema) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *FileSchema) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *FileSchema) SetObject(v string)`

SetObject sets Object field to given value.


### GetBytes

`func (o *FileSchema) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *FileSchema) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *FileSchema) SetBytes(v int32)`

SetBytes sets Bytes field to given value.


### GetCreatedAt

`func (o *FileSchema) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FileSchema) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FileSchema) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetFilename

`func (o *FileSchema) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *FileSchema) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *FileSchema) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetPurpose

`func (o *FileSchema) GetPurpose() FilePurpose`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *FileSchema) GetPurposeOk() (*FilePurpose, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *FileSchema) SetPurpose(v FilePurpose)`

SetPurpose sets Purpose field to given value.


### GetSampleType

`func (o *FileSchema) GetSampleType() SampleType`

GetSampleType returns the SampleType field if non-nil, zero value otherwise.

### GetSampleTypeOk

`func (o *FileSchema) GetSampleTypeOk() (*SampleType, bool)`

GetSampleTypeOk returns a tuple with the SampleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleType

`func (o *FileSchema) SetSampleType(v SampleType)`

SetSampleType sets SampleType field to given value.


### GetNumLines

`func (o *FileSchema) GetNumLines() int32`

GetNumLines returns the NumLines field if non-nil, zero value otherwise.

### GetNumLinesOk

`func (o *FileSchema) GetNumLinesOk() (*int32, bool)`

GetNumLinesOk returns a tuple with the NumLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLines

`func (o *FileSchema) SetNumLines(v int32)`

SetNumLines sets NumLines field to given value.

### HasNumLines

`func (o *FileSchema) HasNumLines() bool`

HasNumLines returns a boolean if a field has been set.

### SetNumLinesNil

`func (o *FileSchema) SetNumLinesNil(b bool)`

 SetNumLinesNil sets the value for NumLines to be an explicit nil

### UnsetNumLines
`func (o *FileSchema) UnsetNumLines()`

UnsetNumLines ensures that no value is present for NumLines, not even an explicit nil
### GetMimetype

`func (o *FileSchema) GetMimetype() string`

GetMimetype returns the Mimetype field if non-nil, zero value otherwise.

### GetMimetypeOk

`func (o *FileSchema) GetMimetypeOk() (*string, bool)`

GetMimetypeOk returns a tuple with the Mimetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimetype

`func (o *FileSchema) SetMimetype(v string)`

SetMimetype sets Mimetype field to given value.

### HasMimetype

`func (o *FileSchema) HasMimetype() bool`

HasMimetype returns a boolean if a field has been set.

### SetMimetypeNil

`func (o *FileSchema) SetMimetypeNil(b bool)`

 SetMimetypeNil sets the value for Mimetype to be an explicit nil

### UnsetMimetype
`func (o *FileSchema) UnsetMimetype()`

UnsetMimetype ensures that no value is present for Mimetype, not even an explicit nil
### GetSource

`func (o *FileSchema) GetSource() Source`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FileSchema) GetSourceOk() (*Source, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FileSchema) SetSource(v Source)`

SetSource sets Source field to given value.


### GetSignature

`func (o *FileSchema) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *FileSchema) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *FileSchema) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *FileSchema) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### SetSignatureNil

`func (o *FileSchema) SetSignatureNil(b bool)`

 SetSignatureNil sets the value for Signature to be an explicit nil

### UnsetSignature
`func (o *FileSchema) UnsetSignature()`

UnsetSignature ensures that no value is present for Signature, not even an explicit nil
### GetExpiresAt

`func (o *FileSchema) GetExpiresAt() int32`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *FileSchema) GetExpiresAtOk() (*int32, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *FileSchema) SetExpiresAt(v int32)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *FileSchema) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *FileSchema) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *FileSchema) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetVisibility

`func (o *FileSchema) GetVisibility() FileVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *FileSchema) GetVisibilityOk() (*FileVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *FileSchema) SetVisibility(v FileVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *FileSchema) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### SetVisibilityNil

`func (o *FileSchema) SetVisibilityNil(b bool)`

 SetVisibilityNil sets the value for Visibility to be an explicit nil

### UnsetVisibility
`func (o *FileSchema) UnsetVisibility()`

UnsetVisibility ensures that no value is present for Visibility, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


