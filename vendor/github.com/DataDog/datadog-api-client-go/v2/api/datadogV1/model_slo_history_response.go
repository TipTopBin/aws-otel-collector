// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"encoding/json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLOHistoryResponse A service level objective history response.
type SLOHistoryResponse struct {
	// An array of service level objective objects.
	Data *SLOHistoryResponseData `json:"data,omitempty"`
	// A list of errors while querying the history data for the service level objective.
	Errors []SLOHistoryResponseError `json:"errors,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewSLOHistoryResponse instantiates a new SLOHistoryResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSLOHistoryResponse() *SLOHistoryResponse {
	this := SLOHistoryResponse{}
	return &this
}

// NewSLOHistoryResponseWithDefaults instantiates a new SLOHistoryResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSLOHistoryResponseWithDefaults() *SLOHistoryResponse {
	this := SLOHistoryResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SLOHistoryResponse) GetData() SLOHistoryResponseData {
	if o == nil || o.Data == nil {
		var ret SLOHistoryResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOHistoryResponse) GetDataOk() (*SLOHistoryResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SLOHistoryResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given SLOHistoryResponseData and assigns it to the Data field.
func (o *SLOHistoryResponse) SetData(v SLOHistoryResponseData) {
	o.Data = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SLOHistoryResponse) GetErrors() []SLOHistoryResponseError {
	if o == nil {
		var ret []SLOHistoryResponseError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SLOHistoryResponse) GetErrorsOk() (*[]SLOHistoryResponseError, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return &o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SLOHistoryResponse) HasErrors() bool {
	return o != nil && o.Errors != nil
}

// SetErrors gets a reference to the given []SLOHistoryResponseError and assigns it to the Errors field.
func (o *SLOHistoryResponse) SetErrors(v []SLOHistoryResponseError) {
	o.Errors = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SLOHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SLOHistoryResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Data   *SLOHistoryResponseData   `json:"data,omitempty"`
		Errors []SLOHistoryResponseError `json:"errors,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "errors"})
	} else {
		return err
	}
	if all.Data != nil && all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Data = all.Data
	o.Errors = all.Errors
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
