# House

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Идентификатор дома | 
**Address** | **string** | Адрес дома | 
**Year** | **int32** | Год постройки дома | 
**Developer** | Pointer to **NullableString** | Застройщик | [optional] 
**CreatedAt** | Pointer to **time.Time** | Дата + время | [optional] 
**UpdateAt** | Pointer to **time.Time** | Дата + время | [optional] 

## Methods

### NewHouse

`func NewHouse(id int32, address string, year int32, ) *House`

NewHouse instantiates a new House object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHouseWithDefaults

`func NewHouseWithDefaults() *House`

NewHouseWithDefaults instantiates a new House object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *House) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *House) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *House) SetId(v int32)`

SetId sets Id field to given value.


### GetAddress

`func (o *House) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *House) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *House) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetYear

`func (o *House) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *House) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *House) SetYear(v int32)`

SetYear sets Year field to given value.


### GetDeveloper

`func (o *House) GetDeveloper() string`

GetDeveloper returns the Developer field if non-nil, zero value otherwise.

### GetDeveloperOk

`func (o *House) GetDeveloperOk() (*string, bool)`

GetDeveloperOk returns a tuple with the Developer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloper

`func (o *House) SetDeveloper(v string)`

SetDeveloper sets Developer field to given value.

### HasDeveloper

`func (o *House) HasDeveloper() bool`

HasDeveloper returns a boolean if a field has been set.

### SetDeveloperNil

`func (o *House) SetDeveloperNil(b bool)`

 SetDeveloperNil sets the value for Developer to be an explicit nil

### UnsetDeveloper
`func (o *House) UnsetDeveloper()`

UnsetDeveloper ensures that no value is present for Developer, not even an explicit nil
### GetCreatedAt

`func (o *House) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *House) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *House) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *House) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdateAt

`func (o *House) GetUpdateAt() time.Time`

GetUpdateAt returns the UpdateAt field if non-nil, zero value otherwise.

### GetUpdateAtOk

`func (o *House) GetUpdateAtOk() (*time.Time, bool)`

GetUpdateAtOk returns a tuple with the UpdateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAt

`func (o *House) SetUpdateAt(v time.Time)`

SetUpdateAt sets UpdateAt field to given value.

### HasUpdateAt

`func (o *House) HasUpdateAt() bool`

HasUpdateAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


