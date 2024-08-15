# HouseCreatePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Адрес дома | 
**Year** | **int32** | Год постройки дома | 
**Developer** | Pointer to **NullableString** | Застройщик | [optional] 

## Methods

### NewHouseCreatePostRequest

`func NewHouseCreatePostRequest(address string, year int32, ) *HouseCreatePostRequest`

NewHouseCreatePostRequest instantiates a new HouseCreatePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHouseCreatePostRequestWithDefaults

`func NewHouseCreatePostRequestWithDefaults() *HouseCreatePostRequest`

NewHouseCreatePostRequestWithDefaults instantiates a new HouseCreatePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *HouseCreatePostRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *HouseCreatePostRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *HouseCreatePostRequest) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetYear

`func (o *HouseCreatePostRequest) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *HouseCreatePostRequest) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *HouseCreatePostRequest) SetYear(v int32)`

SetYear sets Year field to given value.


### GetDeveloper

`func (o *HouseCreatePostRequest) GetDeveloper() string`

GetDeveloper returns the Developer field if non-nil, zero value otherwise.

### GetDeveloperOk

`func (o *HouseCreatePostRequest) GetDeveloperOk() (*string, bool)`

GetDeveloperOk returns a tuple with the Developer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloper

`func (o *HouseCreatePostRequest) SetDeveloper(v string)`

SetDeveloper sets Developer field to given value.

### HasDeveloper

`func (o *HouseCreatePostRequest) HasDeveloper() bool`

HasDeveloper returns a boolean if a field has been set.

### SetDeveloperNil

`func (o *HouseCreatePostRequest) SetDeveloperNil(b bool)`

 SetDeveloperNil sets the value for Developer to be an explicit nil

### UnsetDeveloper
`func (o *HouseCreatePostRequest) UnsetDeveloper()`

UnsetDeveloper ensures that no value is present for Developer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


