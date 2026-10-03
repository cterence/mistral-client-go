# WorkerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchedulerUrl** | **string** |  | 
**Namespace** | **string** |  | 
**Tls** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewWorkerInfo

`func NewWorkerInfo(schedulerUrl string, namespace string, ) *WorkerInfo`

NewWorkerInfo instantiates a new WorkerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkerInfoWithDefaults

`func NewWorkerInfoWithDefaults() *WorkerInfo`

NewWorkerInfoWithDefaults instantiates a new WorkerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedulerUrl

`func (o *WorkerInfo) GetSchedulerUrl() string`

GetSchedulerUrl returns the SchedulerUrl field if non-nil, zero value otherwise.

### GetSchedulerUrlOk

`func (o *WorkerInfo) GetSchedulerUrlOk() (*string, bool)`

GetSchedulerUrlOk returns a tuple with the SchedulerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulerUrl

`func (o *WorkerInfo) SetSchedulerUrl(v string)`

SetSchedulerUrl sets SchedulerUrl field to given value.


### GetNamespace

`func (o *WorkerInfo) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *WorkerInfo) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *WorkerInfo) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetTls

`func (o *WorkerInfo) GetTls() bool`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *WorkerInfo) GetTlsOk() (*bool, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *WorkerInfo) SetTls(v bool)`

SetTls sets Tls field to given value.

### HasTls

`func (o *WorkerInfo) HasTls() bool`

HasTls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


