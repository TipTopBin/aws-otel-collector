// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"encoding/json"
	"fmt"
)

// FormulaAndFunctionSLOMeasure SLO measures queries.
type FormulaAndFunctionSLOMeasure string

// List of FormulaAndFunctionSLOMeasure.
const (
	FORMULAANDFUNCTIONSLOMEASURE_GOOD_EVENTS            FormulaAndFunctionSLOMeasure = "good_events"
	FORMULAANDFUNCTIONSLOMEASURE_BAD_EVENTS             FormulaAndFunctionSLOMeasure = "bad_events"
	FORMULAANDFUNCTIONSLOMEASURE_SLO_STATUS             FormulaAndFunctionSLOMeasure = "slo_status"
	FORMULAANDFUNCTIONSLOMEASURE_ERROR_BUDGET_REMAINING FormulaAndFunctionSLOMeasure = "error_budget_remaining"
	FORMULAANDFUNCTIONSLOMEASURE_BURN_RATE              FormulaAndFunctionSLOMeasure = "burn_rate"
	FORMULAANDFUNCTIONSLOMEASURE_ERROR_BUDGET_BURNDOWN  FormulaAndFunctionSLOMeasure = "error_budget_burndown"
)

var allowedFormulaAndFunctionSLOMeasureEnumValues = []FormulaAndFunctionSLOMeasure{
	FORMULAANDFUNCTIONSLOMEASURE_GOOD_EVENTS,
	FORMULAANDFUNCTIONSLOMEASURE_BAD_EVENTS,
	FORMULAANDFUNCTIONSLOMEASURE_SLO_STATUS,
	FORMULAANDFUNCTIONSLOMEASURE_ERROR_BUDGET_REMAINING,
	FORMULAANDFUNCTIONSLOMEASURE_BURN_RATE,
	FORMULAANDFUNCTIONSLOMEASURE_ERROR_BUDGET_BURNDOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaAndFunctionSLOMeasure) GetAllowedValues() []FormulaAndFunctionSLOMeasure {
	return allowedFormulaAndFunctionSLOMeasureEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaAndFunctionSLOMeasure) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaAndFunctionSLOMeasure(value)
	return nil
}

// NewFormulaAndFunctionSLOMeasureFromValue returns a pointer to a valid FormulaAndFunctionSLOMeasure
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaAndFunctionSLOMeasureFromValue(v string) (*FormulaAndFunctionSLOMeasure, error) {
	ev := FormulaAndFunctionSLOMeasure(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaAndFunctionSLOMeasure: valid values are %v", v, allowedFormulaAndFunctionSLOMeasureEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaAndFunctionSLOMeasure) IsValid() bool {
	for _, existing := range allowedFormulaAndFunctionSLOMeasureEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaAndFunctionSLOMeasure value.
func (v FormulaAndFunctionSLOMeasure) Ptr() *FormulaAndFunctionSLOMeasure {
	return &v
}

// NullableFormulaAndFunctionSLOMeasure handles when a null is used for FormulaAndFunctionSLOMeasure.
type NullableFormulaAndFunctionSLOMeasure struct {
	value *FormulaAndFunctionSLOMeasure
	isSet bool
}

// Get returns the associated value.
func (v NullableFormulaAndFunctionSLOMeasure) Get() *FormulaAndFunctionSLOMeasure {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableFormulaAndFunctionSLOMeasure) Set(val *FormulaAndFunctionSLOMeasure) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableFormulaAndFunctionSLOMeasure) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableFormulaAndFunctionSLOMeasure) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableFormulaAndFunctionSLOMeasure initializes the struct as if Set has been called.
func NewNullableFormulaAndFunctionSLOMeasure(val *FormulaAndFunctionSLOMeasure) *NullableFormulaAndFunctionSLOMeasure {
	return &NullableFormulaAndFunctionSLOMeasure{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableFormulaAndFunctionSLOMeasure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableFormulaAndFunctionSLOMeasure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
