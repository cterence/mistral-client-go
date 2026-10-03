# GetChatCompletionEventsInSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchParams** | [**FilterPayload**](FilterPayload.md) |  | 
**ExtraFields** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetChatCompletionEventsInSchema

`func NewGetChatCompletionEventsInSchema(searchParams FilterPayload, ) *GetChatCompletionEventsInSchema`

NewGetChatCompletionEventsInSchema instantiates a new GetChatCompletionEventsInSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetChatCompletionEventsInSchemaWithDefaults

`func NewGetChatCompletionEventsInSchemaWithDefaults() *GetChatCompletionEventsInSchema`

NewGetChatCompletionEventsInSchemaWithDefaults instantiates a new GetChatCompletionEventsInSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchParams

`func (o *GetChatCompletionEventsInSchema) GetSearchParams() FilterPayload`

GetSearchParams returns the SearchParams field if non-nil, zero value otherwise.

### GetSearchParamsOk

`func (o *GetChatCompletionEventsInSchema) GetSearchParamsOk() (*FilterPayload, bool)`

GetSearchParamsOk returns a tuple with the SearchParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchParams

`func (o *GetChatCompletionEventsInSchema) SetSearchParams(v FilterPayload)`

SetSearchParams sets SearchParams field to given value.


### GetExtraFields

`func (o *GetChatCompletionEventsInSchema) GetExtraFields() []string`

GetExtraFields returns the ExtraFields field if non-nil, zero value otherwise.

### GetExtraFieldsOk

`func (o *GetChatCompletionEventsInSchema) GetExtraFieldsOk() (*[]string, bool)`

GetExtraFieldsOk returns a tuple with the ExtraFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraFields

`func (o *GetChatCompletionEventsInSchema) SetExtraFields(v []string)`

SetExtraFields sets ExtraFields field to given value.

### HasExtraFields

`func (o *GetChatCompletionEventsInSchema) HasExtraFields() bool`

HasExtraFields returns a boolean if a field has been set.

### SetExtraFieldsNil

`func (o *GetChatCompletionEventsInSchema) SetExtraFieldsNil(b bool)`

 SetExtraFieldsNil sets the value for ExtraFields to be an explicit nil

### UnsetExtraFields
`func (o *GetChatCompletionEventsInSchema) UnsetExtraFields()`

UnsetExtraFields ensures that no value is present for ExtraFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


