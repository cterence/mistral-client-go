# PostJudgeInSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | 
**ModelName** | **string** |  | 
**Output** | [**Output**](Output.md) |  | 
**Instructions** | **string** |  | 
**Tools** | **[]string** |  | 

## Methods

### NewPostJudgeInSchema

`func NewPostJudgeInSchema(name string, description string, modelName string, output Output, instructions string, tools []string, ) *PostJudgeInSchema`

NewPostJudgeInSchema instantiates a new PostJudgeInSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostJudgeInSchemaWithDefaults

`func NewPostJudgeInSchemaWithDefaults() *PostJudgeInSchema`

NewPostJudgeInSchemaWithDefaults instantiates a new PostJudgeInSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostJudgeInSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostJudgeInSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostJudgeInSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PostJudgeInSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostJudgeInSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostJudgeInSchema) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetModelName

`func (o *PostJudgeInSchema) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *PostJudgeInSchema) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *PostJudgeInSchema) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetOutput

`func (o *PostJudgeInSchema) GetOutput() Output`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *PostJudgeInSchema) GetOutputOk() (*Output, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *PostJudgeInSchema) SetOutput(v Output)`

SetOutput sets Output field to given value.


### GetInstructions

`func (o *PostJudgeInSchema) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *PostJudgeInSchema) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *PostJudgeInSchema) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.


### GetTools

`func (o *PostJudgeInSchema) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *PostJudgeInSchema) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *PostJudgeInSchema) SetTools(v []string)`

SetTools sets Tools field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


