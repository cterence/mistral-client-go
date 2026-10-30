# BasicModelUsageDataJSON

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Models** | [**map[string]map[string][]map[string]interface{}**](map.md) | Usage data grouped by model. | 

## Methods

### NewBasicModelUsageDataJSON

`func NewBasicModelUsageDataJSON(models map[string]map[string][]map[string]interface{}, ) *BasicModelUsageDataJSON`

NewBasicModelUsageDataJSON instantiates a new BasicModelUsageDataJSON object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicModelUsageDataJSONWithDefaults

`func NewBasicModelUsageDataJSONWithDefaults() *BasicModelUsageDataJSON`

NewBasicModelUsageDataJSONWithDefaults instantiates a new BasicModelUsageDataJSON object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModels

`func (o *BasicModelUsageDataJSON) GetModels() map[string]map[string][]map[string]interface{}`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *BasicModelUsageDataJSON) GetModelsOk() (*map[string]map[string][]map[string]interface{}, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *BasicModelUsageDataJSON) SetModels(v map[string]map[string][]map[string]interface{})`

SetModels sets Models field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


