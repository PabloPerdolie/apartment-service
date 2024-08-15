# Flat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Идентификатор квартиры | 
**HouseId** | **int32** | Идентификатор дома | 
**Price** | **int32** | Цена квартиры в у.е. | 
**Rooms** | **int32** | Количество комнат в квартире | 
**Status** | [**Status**](Status.md) |  | 

## Methods

### NewFlat

`func NewFlat(id int32, houseId int32, price int32, rooms int32, status Status, ) *Flat`

NewFlat instantiates a new Flat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlatWithDefaults

`func NewFlatWithDefaults() *Flat`

NewFlatWithDefaults instantiates a new Flat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Flat) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Flat) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Flat) SetId(v int32)`

SetId sets Id field to given value.


### GetHouseId

`func (o *Flat) GetHouseId() int32`

GetHouseId returns the HouseId field if non-nil, zero value otherwise.

### GetHouseIdOk

`func (o *Flat) GetHouseIdOk() (*int32, bool)`

GetHouseIdOk returns a tuple with the HouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseId

`func (o *Flat) SetHouseId(v int32)`

SetHouseId sets HouseId field to given value.


### GetPrice

`func (o *Flat) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Flat) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Flat) SetPrice(v int32)`

SetPrice sets Price field to given value.


### GetRooms

`func (o *Flat) GetRooms() int32`

GetRooms returns the Rooms field if non-nil, zero value otherwise.

### GetRoomsOk

`func (o *Flat) GetRoomsOk() (*int32, bool)`

GetRoomsOk returns a tuple with the Rooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRooms

`func (o *Flat) SetRooms(v int32)`

SetRooms sets Rooms field to given value.


### GetStatus

`func (o *Flat) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Flat) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Flat) SetStatus(v Status)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


