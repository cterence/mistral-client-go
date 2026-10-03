# QueryWorkflowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryName** | **string** |  | 
**Result** | **interface{}** | The result of the Query workflow call | 

## Methods

### NewQueryWorkflowResponse

`func NewQueryWorkflowResponse(queryName string, result interface{}, ) *QueryWorkflowResponse`

NewQueryWorkflowResponse instantiates a new QueryWorkflowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryWorkflowResponseWithDefaults

`func NewQueryWorkflowResponseWithDefaults() *QueryWorkflowResponse`

NewQueryWorkflowResponseWithDefaults instantiates a new QueryWorkflowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryName

`func (o *QueryWorkflowResponse) GetQueryName() string`

GetQueryName returns the QueryName field if non-nil, zero value otherwise.

### GetQueryNameOk

`func (o *QueryWorkflowResponse) GetQueryNameOk() (*string, bool)`

GetQueryNameOk returns a tuple with the QueryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryName

`func (o *QueryWorkflowResponse) SetQueryName(v string)`

SetQueryName sets QueryName field to given value.


### GetResult

`func (o *QueryWorkflowResponse) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *QueryWorkflowResponse) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *QueryWorkflowResponse) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *QueryWorkflowResponse) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *QueryWorkflowResponse) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


