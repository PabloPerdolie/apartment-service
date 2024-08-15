/*
Тестовое задание для отбора на Backend Bootcamp

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the HouseIdSubscribePostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HouseIdSubscribePostRequest{}

// HouseIdSubscribePostRequest struct for HouseIdSubscribePostRequest
type HouseIdSubscribePostRequest struct {
	// Email пользователя
	Email string `json:"email"`
}

type _HouseIdSubscribePostRequest HouseIdSubscribePostRequest

// NewHouseIdSubscribePostRequest instantiates a new HouseIdSubscribePostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHouseIdSubscribePostRequest(email string) *HouseIdSubscribePostRequest {
	this := HouseIdSubscribePostRequest{}
	this.Email = email
	return &this
}

// NewHouseIdSubscribePostRequestWithDefaults instantiates a new HouseIdSubscribePostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHouseIdSubscribePostRequestWithDefaults() *HouseIdSubscribePostRequest {
	this := HouseIdSubscribePostRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *HouseIdSubscribePostRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *HouseIdSubscribePostRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *HouseIdSubscribePostRequest) SetEmail(v string) {
	o.Email = v
}

func (o HouseIdSubscribePostRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HouseIdSubscribePostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	return toSerialize, nil
}

func (o *HouseIdSubscribePostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varHouseIdSubscribePostRequest := _HouseIdSubscribePostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHouseIdSubscribePostRequest)

	if err != nil {
		return err
	}

	*o = HouseIdSubscribePostRequest(varHouseIdSubscribePostRequest)

	return err
}

type NullableHouseIdSubscribePostRequest struct {
	value *HouseIdSubscribePostRequest
	isSet bool
}

func (v NullableHouseIdSubscribePostRequest) Get() *HouseIdSubscribePostRequest {
	return v.value
}

func (v *NullableHouseIdSubscribePostRequest) Set(val *HouseIdSubscribePostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHouseIdSubscribePostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHouseIdSubscribePostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHouseIdSubscribePostRequest(val *HouseIdSubscribePostRequest) *NullableHouseIdSubscribePostRequest {
	return &NullableHouseIdSubscribePostRequest{value: val, isSet: true}
}

func (v NullableHouseIdSubscribePostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHouseIdSubscribePostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
