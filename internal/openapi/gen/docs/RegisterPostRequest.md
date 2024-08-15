# RegisterPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email пользователя | [optional] 
**Password** | Pointer to **string** | Пароль пользователя | [optional] 
**UserType** | Pointer to [**UserType**](UserType.md) |  | [optional] 

## Methods

### NewRegisterPostRequest

`func NewRegisterPostRequest() *RegisterPostRequest`

NewRegisterPostRequest instantiates a new RegisterPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterPostRequestWithDefaults

`func NewRegisterPostRequestWithDefaults() *RegisterPostRequest`

NewRegisterPostRequestWithDefaults instantiates a new RegisterPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *RegisterPostRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RegisterPostRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RegisterPostRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *RegisterPostRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPassword

`func (o *RegisterPostRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RegisterPostRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RegisterPostRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RegisterPostRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUserType

`func (o *RegisterPostRequest) GetUserType() UserType`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *RegisterPostRequest) GetUserTypeOk() (*UserType, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *RegisterPostRequest) SetUserType(v UserType)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *RegisterPostRequest) HasUserType() bool`

HasUserType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


