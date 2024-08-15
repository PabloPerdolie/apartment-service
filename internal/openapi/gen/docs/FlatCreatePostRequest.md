# FlatCreatePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HouseId** | **int32** | Идентификатор дома | 
**Price** | **int32** | Цена квартиры в у.е. | 
**Rooms** | Pointer to **int32** | Количество комнат в квартире | [optional] 

## Methods

### NewFlatCreatePostRequest

`func NewFlatCreatePostRequest(houseId int32, price int32, ) *FlatCreatePostRequest`

NewFlatCreatePostRequest instantiates a new FlatCreatePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlatCreatePostRequestWithDefaults

`func NewFlatCreatePostRequestWithDefaults() *FlatCreatePostRequest`

NewFlatCreatePostRequestWithDefaults instantiates a new FlatCreatePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHouseId

`func (o *FlatCreatePostRequest) GetHouseId() int32`

GetHouseId returns the HouseId field if non-nil, zero value otherwise.

### GetHouseIdOk

`func (o *FlatCreatePostRequest) GetHouseIdOk() (*int32, bool)`

GetHouseIdOk returns a tuple with the HouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseId

`func (o *FlatCreatePostRequest) SetHouseId(v int32)`

SetHouseId sets HouseId field to given value.


### GetPrice

`func (o *FlatCreatePostRequest) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *FlatCreatePostRequest) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *FlatCreatePostRequest) SetPrice(v int32)`

SetPrice sets Price field to given value.


### GetRooms

`func (o *FlatCreatePostRequest) GetRooms() int32`

GetRooms returns the Rooms field if non-nil, zero value otherwise.

### GetRoomsOk

`func (o *FlatCreatePostRequest) GetRoomsOk() (*int32, bool)`

GetRoomsOk returns a tuple with the Rooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRooms

`func (o *FlatCreatePostRequest) SetRooms(v int32)`

SetRooms sets Rooms field to given value.

### HasRooms

`func (o *FlatCreatePostRequest) HasRooms() bool`

HasRooms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


