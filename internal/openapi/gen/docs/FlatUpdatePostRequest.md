# FlatUpdatePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Идентификатор квартиры | 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 

## Methods

### NewFlatUpdatePostRequest

`func NewFlatUpdatePostRequest(id int32, ) *FlatUpdatePostRequest`

NewFlatUpdatePostRequest instantiates a new FlatUpdatePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlatUpdatePostRequestWithDefaults

`func NewFlatUpdatePostRequestWithDefaults() *FlatUpdatePostRequest`

NewFlatUpdatePostRequestWithDefaults instantiates a new FlatUpdatePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FlatUpdatePostRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlatUpdatePostRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlatUpdatePostRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetStatus

`func (o *FlatUpdatePostRequest) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlatUpdatePostRequest) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlatUpdatePostRequest) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FlatUpdatePostRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


