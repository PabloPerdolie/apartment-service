# LoginPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Идентификатор пользователя | [optional] 
**Password** | Pointer to **string** | Пароль пользователя | [optional] 

## Methods

### NewLoginPostRequest

`func NewLoginPostRequest() *LoginPostRequest`

NewLoginPostRequest instantiates a new LoginPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginPostRequestWithDefaults

`func NewLoginPostRequestWithDefaults() *LoginPostRequest`

NewLoginPostRequestWithDefaults instantiates a new LoginPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoginPostRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoginPostRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoginPostRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LoginPostRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPassword

`func (o *LoginPostRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LoginPostRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LoginPostRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *LoginPostRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


