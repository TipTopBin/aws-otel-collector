// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"encoding/json"
	"fmt"
)

// SyntheticsBrowserTestType Type of the Synthetic test, `browser`.
type SyntheticsBrowserTestType string

// List of SyntheticsBrowserTestType.
const (
	SYNTHETICSBROWSERTESTTYPE_BROWSER SyntheticsBrowserTestType = "browser"
)

var allowedSyntheticsBrowserTestTypeEnumValues = []SyntheticsBrowserTestType{
	SYNTHETICSBROWSERTESTTYPE_BROWSER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsBrowserTestType) GetAllowedValues() []SyntheticsBrowserTestType {
	return allowedSyntheticsBrowserTestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsBrowserTestType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsBrowserTestType(value)
	return nil
}

// NewSyntheticsBrowserTestTypeFromValue returns a pointer to a valid SyntheticsBrowserTestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsBrowserTestTypeFromValue(v string) (*SyntheticsBrowserTestType, error) {
	ev := SyntheticsBrowserTestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsBrowserTestType: valid values are %v", v, allowedSyntheticsBrowserTestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsBrowserTestType) IsValid() bool {
	for _, existing := range allowedSyntheticsBrowserTestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsBrowserTestType value.
func (v SyntheticsBrowserTestType) Ptr() *SyntheticsBrowserTestType {
	return &v
}

// NullableSyntheticsBrowserTestType handles when a null is used for SyntheticsBrowserTestType.
type NullableSyntheticsBrowserTestType struct {
	value *SyntheticsBrowserTestType
	isSet bool
}

// Get returns the associated value.
func (v NullableSyntheticsBrowserTestType) Get() *SyntheticsBrowserTestType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableSyntheticsBrowserTestType) Set(val *SyntheticsBrowserTestType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableSyntheticsBrowserTestType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableSyntheticsBrowserTestType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSyntheticsBrowserTestType initializes the struct as if Set has been called.
func NewNullableSyntheticsBrowserTestType(val *SyntheticsBrowserTestType) *NullableSyntheticsBrowserTestType {
	return &NullableSyntheticsBrowserTestType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableSyntheticsBrowserTestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableSyntheticsBrowserTestType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
