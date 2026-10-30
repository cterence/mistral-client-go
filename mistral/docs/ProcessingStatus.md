# ProcessingStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | **string** |  | 
**ProcessStatus** | [**ProcessStatus**](ProcessStatus.md) | Processing status of the document. | 
**ProcessingStatus** | **string** |  | [readonly] 

## Methods

### NewProcessingStatus

`func NewProcessingStatus(documentId string, processStatus ProcessStatus, processingStatus string, ) *ProcessingStatus`

NewProcessingStatus instantiates a new ProcessingStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessingStatusWithDefaults

`func NewProcessingStatusWithDefaults() *ProcessingStatus`

NewProcessingStatusWithDefaults instantiates a new ProcessingStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *ProcessingStatus) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *ProcessingStatus) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *ProcessingStatus) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.


### GetProcessStatus

`func (o *ProcessingStatus) GetProcessStatus() ProcessStatus`

GetProcessStatus returns the ProcessStatus field if non-nil, zero value otherwise.

### GetProcessStatusOk

`func (o *ProcessingStatus) GetProcessStatusOk() (*ProcessStatus, bool)`

GetProcessStatusOk returns a tuple with the ProcessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessStatus

`func (o *ProcessingStatus) SetProcessStatus(v ProcessStatus)`

SetProcessStatus sets ProcessStatus field to given value.


### GetProcessingStatus

`func (o *ProcessingStatus) GetProcessingStatus() string`

GetProcessingStatus returns the ProcessingStatus field if non-nil, zero value otherwise.

### GetProcessingStatusOk

`func (o *ProcessingStatus) GetProcessingStatusOk() (*string, bool)`

GetProcessingStatusOk returns a tuple with the ProcessingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingStatus

`func (o *ProcessingStatus) SetProcessingStatus(v string)`

SetProcessingStatus sets ProcessingStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


