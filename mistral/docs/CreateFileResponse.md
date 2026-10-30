# CreateFileResponse

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

### NewCreateFileResponse

`func NewCreateFileResponse(id string, object string, bytes int32, createdAt int32, filename string, purpose FilePurpose, sampleType SampleType, source Source, ) *CreateFileResponse`

NewCreateFileResponse instantiates a new CreateFileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFileResponseWithDefaults

`func NewCreateFileResponseWithDefaults() *CreateFileResponse`

NewCreateFileResponseWithDefaults instantiates a new CreateFileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateFileResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateFileResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateFileResponse) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *CreateFileResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CreateFileResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CreateFileResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetBytes

`func (o *CreateFileResponse) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *CreateFileResponse) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *CreateFileResponse) SetBytes(v int32)`

SetBytes sets Bytes field to given value.


### GetCreatedAt

`func (o *CreateFileResponse) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateFileResponse) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateFileResponse) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetFilename

`func (o *CreateFileResponse) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *CreateFileResponse) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *CreateFileResponse) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetPurpose

`func (o *CreateFileResponse) GetPurpose() FilePurpose`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *CreateFileResponse) GetPurposeOk() (*FilePurpose, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *CreateFileResponse) SetPurpose(v FilePurpose)`

SetPurpose sets Purpose field to given value.


### GetSampleType

`func (o *CreateFileResponse) GetSampleType() SampleType`

GetSampleType returns the SampleType field if non-nil, zero value otherwise.

### GetSampleTypeOk

`func (o *CreateFileResponse) GetSampleTypeOk() (*SampleType, bool)`

GetSampleTypeOk returns a tuple with the SampleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleType

`func (o *CreateFileResponse) SetSampleType(v SampleType)`

SetSampleType sets SampleType field to given value.


### GetNumLines

`func (o *CreateFileResponse) GetNumLines() int32`

GetNumLines returns the NumLines field if non-nil, zero value otherwise.

### GetNumLinesOk

`func (o *CreateFileResponse) GetNumLinesOk() (*int32, bool)`

GetNumLinesOk returns a tuple with the NumLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLines

`func (o *CreateFileResponse) SetNumLines(v int32)`

SetNumLines sets NumLines field to given value.

### HasNumLines

`func (o *CreateFileResponse) HasNumLines() bool`

HasNumLines returns a boolean if a field has been set.

### SetNumLinesNil

`func (o *CreateFileResponse) SetNumLinesNil(b bool)`

 SetNumLinesNil sets the value for NumLines to be an explicit nil

### UnsetNumLines
`func (o *CreateFileResponse) UnsetNumLines()`

UnsetNumLines ensures that no value is present for NumLines, not even an explicit nil
### GetMimetype

`func (o *CreateFileResponse) GetMimetype() string`

GetMimetype returns the Mimetype field if non-nil, zero value otherwise.

### GetMimetypeOk

`func (o *CreateFileResponse) GetMimetypeOk() (*string, bool)`

GetMimetypeOk returns a tuple with the Mimetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimetype

`func (o *CreateFileResponse) SetMimetype(v string)`

SetMimetype sets Mimetype field to given value.

### HasMimetype

`func (o *CreateFileResponse) HasMimetype() bool`

HasMimetype returns a boolean if a field has been set.

### SetMimetypeNil

`func (o *CreateFileResponse) SetMimetypeNil(b bool)`

 SetMimetypeNil sets the value for Mimetype to be an explicit nil

### UnsetMimetype
`func (o *CreateFileResponse) UnsetMimetype()`

UnsetMimetype ensures that no value is present for Mimetype, not even an explicit nil
### GetSource

`func (o *CreateFileResponse) GetSource() Source`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CreateFileResponse) GetSourceOk() (*Source, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CreateFileResponse) SetSource(v Source)`

SetSource sets Source field to given value.


### GetSignature

`func (o *CreateFileResponse) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *CreateFileResponse) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *CreateFileResponse) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *CreateFileResponse) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### SetSignatureNil

`func (o *CreateFileResponse) SetSignatureNil(b bool)`

 SetSignatureNil sets the value for Signature to be an explicit nil

### UnsetSignature
`func (o *CreateFileResponse) UnsetSignature()`

UnsetSignature ensures that no value is present for Signature, not even an explicit nil
### GetExpiresAt

`func (o *CreateFileResponse) GetExpiresAt() int32`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateFileResponse) GetExpiresAtOk() (*int32, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateFileResponse) SetExpiresAt(v int32)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreateFileResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *CreateFileResponse) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *CreateFileResponse) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetVisibility

`func (o *CreateFileResponse) GetVisibility() FileVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *CreateFileResponse) GetVisibilityOk() (*FileVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *CreateFileResponse) SetVisibility(v FileVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *CreateFileResponse) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### SetVisibilityNil

`func (o *CreateFileResponse) SetVisibilityNil(b bool)`

 SetVisibilityNil sets the value for Visibility to be an explicit nil

### UnsetVisibility
`func (o *CreateFileResponse) UnsetVisibility()`

UnsetVisibility ensures that no value is present for Visibility, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


