// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// CIAppResponseStatus The status of the response.
type CIAppResponseStatus string

// List of CIAppResponseStatus.
const (
	CIAPPRESPONSESTATUS_DONE    CIAppResponseStatus = "done"
	CIAPPRESPONSESTATUS_TIMEOUT CIAppResponseStatus = "timeout"
)

var allowedCIAppResponseStatusEnumValues = []CIAppResponseStatus{
	CIAPPRESPONSESTATUS_DONE,
	CIAPPRESPONSESTATUS_TIMEOUT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CIAppResponseStatus) GetAllowedValues() []CIAppResponseStatus {
	return allowedCIAppResponseStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CIAppResponseStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CIAppResponseStatus(value)
	return nil
}

// NewCIAppResponseStatusFromValue returns a pointer to a valid CIAppResponseStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCIAppResponseStatusFromValue(v string) (*CIAppResponseStatus, error) {
	ev := CIAppResponseStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CIAppResponseStatus: valid values are %v", v, allowedCIAppResponseStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CIAppResponseStatus) IsValid() bool {
	for _, existing := range allowedCIAppResponseStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CIAppResponseStatus value.
func (v CIAppResponseStatus) Ptr() *CIAppResponseStatus {
	return &v
}

// NullableCIAppResponseStatus handles when a null is used for CIAppResponseStatus.
type NullableCIAppResponseStatus struct {
	value *CIAppResponseStatus
	isSet bool
}

// Get returns the associated value.
func (v NullableCIAppResponseStatus) Get() *CIAppResponseStatus {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableCIAppResponseStatus) Set(val *CIAppResponseStatus) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableCIAppResponseStatus) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableCIAppResponseStatus) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableCIAppResponseStatus initializes the struct as if Set has been called.
func NewNullableCIAppResponseStatus(val *CIAppResponseStatus) *NullableCIAppResponseStatus {
	return &NullableCIAppResponseStatus{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableCIAppResponseStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableCIAppResponseStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
