# ProcessingStatusOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | **string** |  | 
**ProcessStatus** | [**ProcessStatus**](ProcessStatus.md) |  | 
**ProcessingStatus** | **string** |  | [readonly] 

## Methods

### NewProcessingStatusOut

`func NewProcessingStatusOut(documentId string, processStatus ProcessStatus, processingStatus string, ) *ProcessingStatusOut`

NewProcessingStatusOut instantiates a new ProcessingStatusOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessingStatusOutWithDefaults

`func NewProcessingStatusOutWithDefaults() *ProcessingStatusOut`

NewProcessingStatusOutWithDefaults instantiates a new ProcessingStatusOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *ProcessingStatusOut) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *ProcessingStatusOut) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *ProcessingStatusOut) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.


### GetProcessStatus

`func (o *ProcessingStatusOut) GetProcessStatus() ProcessStatus`

GetProcessStatus returns the ProcessStatus field if non-nil, zero value otherwise.

### GetProcessStatusOk

`func (o *ProcessingStatusOut) GetProcessStatusOk() (*ProcessStatus, bool)`

GetProcessStatusOk returns a tuple with the ProcessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessStatus

`func (o *ProcessingStatusOut) SetProcessStatus(v ProcessStatus)`

SetProcessStatus sets ProcessStatus field to given value.


### GetProcessingStatus

`func (o *ProcessingStatusOut) GetProcessingStatus() string`

GetProcessingStatus returns the ProcessingStatus field if non-nil, zero value otherwise.

### GetProcessingStatusOk

`func (o *ProcessingStatusOut) GetProcessingStatusOk() (*string, bool)`

GetProcessingStatusOk returns a tuple with the ProcessingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingStatus

`func (o *ProcessingStatusOut) SetProcessingStatus(v string)`

SetProcessingStatus sets ProcessingStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


