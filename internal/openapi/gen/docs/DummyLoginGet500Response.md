# DummyLoginGet500Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Описание ошибки | 
**RequestId** | Pointer to **string** | Идентификатор запроса. Предназначен для более быстрого поиска проблем. | [optional] 
**Code** | Pointer to **int32** | Код ошибки. Предназначен для классификации проблем и более быстрого решения проблем. | [optional] 

## Methods

### NewDummyLoginGet500Response

`func NewDummyLoginGet500Response(message string, ) *DummyLoginGet500Response`

NewDummyLoginGet500Response instantiates a new DummyLoginGet500Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDummyLoginGet500ResponseWithDefaults

`func NewDummyLoginGet500ResponseWithDefaults() *DummyLoginGet500Response`

NewDummyLoginGet500ResponseWithDefaults instantiates a new DummyLoginGet500Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *DummyLoginGet500Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DummyLoginGet500Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DummyLoginGet500Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetRequestId

`func (o *DummyLoginGet500Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DummyLoginGet500Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DummyLoginGet500Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DummyLoginGet500Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetCode

`func (o *DummyLoginGet500Response) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DummyLoginGet500Response) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DummyLoginGet500Response) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *DummyLoginGet500Response) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


