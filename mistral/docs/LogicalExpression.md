# LogicalExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Expressions** | [**[]ExpressionsInner**](ExpressionsInner.md) |  | 

## Methods

### NewLogicalExpression

`func NewLogicalExpression(type_ string, expressions []ExpressionsInner, ) *LogicalExpression`

NewLogicalExpression instantiates a new LogicalExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogicalExpressionWithDefaults

`func NewLogicalExpressionWithDefaults() *LogicalExpression`

NewLogicalExpressionWithDefaults instantiates a new LogicalExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LogicalExpression) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogicalExpression) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogicalExpression) SetType(v string)`

SetType sets Type field to given value.


### GetExpressions

`func (o *LogicalExpression) GetExpressions() []ExpressionsInner`

GetExpressions returns the Expressions field if non-nil, zero value otherwise.

### GetExpressionsOk

`func (o *LogicalExpression) GetExpressionsOk() (*[]ExpressionsInner, bool)`

GetExpressionsOk returns a tuple with the Expressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressions

`func (o *LogicalExpression) SetExpressions(v []ExpressionsInner)`

SetExpressions sets Expressions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


