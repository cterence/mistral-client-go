# AdminUsageQueryFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Usage dimension to filter. | 
**Operator** | **string** | Whether rows must match or must not match the supplied values. | 
**Values** | **[]string** | Accepted or excluded values. Use null to filter missing dimension values. | 

## Methods

### NewAdminUsageQueryFilter

`func NewAdminUsageQueryFilter(field string, operator string, values []*string, ) *AdminUsageQueryFilter`

NewAdminUsageQueryFilter instantiates a new AdminUsageQueryFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminUsageQueryFilterWithDefaults

`func NewAdminUsageQueryFilterWithDefaults() *AdminUsageQueryFilter`

NewAdminUsageQueryFilterWithDefaults instantiates a new AdminUsageQueryFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *AdminUsageQueryFilter) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *AdminUsageQueryFilter) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *AdminUsageQueryFilter) SetField(v string)`

SetField sets Field field to given value.


### GetOperator

`func (o *AdminUsageQueryFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *AdminUsageQueryFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *AdminUsageQueryFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetValues

`func (o *AdminUsageQueryFilter) GetValues() []*string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *AdminUsageQueryFilter) GetValuesOk() (*[]*string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *AdminUsageQueryFilter) SetValues(v []*string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


