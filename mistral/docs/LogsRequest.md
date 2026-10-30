# LogsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchExpression** | Pointer to **NullableString** |  | [optional] 
**Order** | Pointer to **string** |  | [optional] [default to "desc"]

## Methods

### NewLogsRequest

`func NewLogsRequest() *LogsRequest`

NewLogsRequest instantiates a new LogsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsRequestWithDefaults

`func NewLogsRequestWithDefaults() *LogsRequest`

NewLogsRequestWithDefaults instantiates a new LogsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchExpression

`func (o *LogsRequest) GetSearchExpression() string`

GetSearchExpression returns the SearchExpression field if non-nil, zero value otherwise.

### GetSearchExpressionOk

`func (o *LogsRequest) GetSearchExpressionOk() (*string, bool)`

GetSearchExpressionOk returns a tuple with the SearchExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchExpression

`func (o *LogsRequest) SetSearchExpression(v string)`

SetSearchExpression sets SearchExpression field to given value.

### HasSearchExpression

`func (o *LogsRequest) HasSearchExpression() bool`

HasSearchExpression returns a boolean if a field has been set.

### SetSearchExpressionNil

`func (o *LogsRequest) SetSearchExpressionNil(b bool)`

 SetSearchExpressionNil sets the value for SearchExpression to be an explicit nil

### UnsetSearchExpression
`func (o *LogsRequest) UnsetSearchExpression()`

UnsetSearchExpression ensures that no value is present for SearchExpression, not even an explicit nil
### GetOrder

`func (o *LogsRequest) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *LogsRequest) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *LogsRequest) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *LogsRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


