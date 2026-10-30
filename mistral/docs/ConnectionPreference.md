# ConnectionPreference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ToolConfiguration** | [**ToolExecutionConfiguration**](ToolExecutionConfiguration.md) |  | 
**IsDefault** | Pointer to **NullableBool** |  | [optional] 
**ConsumerType** | Pointer to [**NullableConsumerType**](ConsumerType.md) |  | [optional] 

## Methods

### NewConnectionPreference

`func NewConnectionPreference(name string, toolConfiguration ToolExecutionConfiguration, ) *ConnectionPreference`

NewConnectionPreference instantiates a new ConnectionPreference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionPreferenceWithDefaults

`func NewConnectionPreferenceWithDefaults() *ConnectionPreference`

NewConnectionPreferenceWithDefaults instantiates a new ConnectionPreference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConnectionPreference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectionPreference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectionPreference) SetName(v string)`

SetName sets Name field to given value.


### GetToolConfiguration

`func (o *ConnectionPreference) GetToolConfiguration() ToolExecutionConfiguration`

GetToolConfiguration returns the ToolConfiguration field if non-nil, zero value otherwise.

### GetToolConfigurationOk

`func (o *ConnectionPreference) GetToolConfigurationOk() (*ToolExecutionConfiguration, bool)`

GetToolConfigurationOk returns a tuple with the ToolConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolConfiguration

`func (o *ConnectionPreference) SetToolConfiguration(v ToolExecutionConfiguration)`

SetToolConfiguration sets ToolConfiguration field to given value.


### GetIsDefault

`func (o *ConnectionPreference) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ConnectionPreference) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ConnectionPreference) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *ConnectionPreference) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *ConnectionPreference) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *ConnectionPreference) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetConsumerType

`func (o *ConnectionPreference) GetConsumerType() ConsumerType`

GetConsumerType returns the ConsumerType field if non-nil, zero value otherwise.

### GetConsumerTypeOk

`func (o *ConnectionPreference) GetConsumerTypeOk() (*ConsumerType, bool)`

GetConsumerTypeOk returns a tuple with the ConsumerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerType

`func (o *ConnectionPreference) SetConsumerType(v ConsumerType)`

SetConsumerType sets ConsumerType field to given value.

### HasConsumerType

`func (o *ConnectionPreference) HasConsumerType() bool`

HasConsumerType returns a boolean if a field has been set.

### SetConsumerTypeNil

`func (o *ConnectionPreference) SetConsumerTypeNil(b bool)`

 SetConsumerTypeNil sets the value for ConsumerType to be an explicit nil

### UnsetConsumerType
`func (o *ConnectionPreference) UnsetConsumerType()`

UnsetConsumerType ensures that no value is present for ConsumerType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


