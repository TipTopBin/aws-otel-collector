// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTodoResponseData Incident todo response data.
type IncidentTodoResponseData struct {
	// Incident todo's attributes.
	Attributes *IncidentTodoAttributes `json:"attributes,omitempty"`
	// The incident todo's ID.
	Id string `json:"id"`
	// Todo resource type.
	Type IncidentTodoType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewIncidentTodoResponseData instantiates a new IncidentTodoResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentTodoResponseData(id string, typeVar IncidentTodoType) *IncidentTodoResponseData {
	this := IncidentTodoResponseData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewIncidentTodoResponseDataWithDefaults instantiates a new IncidentTodoResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentTodoResponseDataWithDefaults() *IncidentTodoResponseData {
	this := IncidentTodoResponseData{}
	var typeVar IncidentTodoType = INCIDENTTODOTYPE_INCIDENT_TODOS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *IncidentTodoResponseData) GetAttributes() IncidentTodoAttributes {
	if o == nil || o.Attributes == nil {
		var ret IncidentTodoAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTodoResponseData) GetAttributesOk() (*IncidentTodoAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *IncidentTodoResponseData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given IncidentTodoAttributes and assigns it to the Attributes field.
func (o *IncidentTodoResponseData) SetAttributes(v IncidentTodoAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *IncidentTodoResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IncidentTodoResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *IncidentTodoResponseData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *IncidentTodoResponseData) GetType() IncidentTodoType {
	if o == nil {
		var ret IncidentTodoType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IncidentTodoResponseData) GetTypeOk() (*IncidentTodoType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *IncidentTodoResponseData) SetType(v IncidentTodoType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentTodoResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentTodoResponseData) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Attributes *IncidentTodoAttributes `json:"attributes,omitempty"`
		Id         *string                 `json:"id"`
		Type       *IncidentTodoType       `json:"type"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}
	if v := all.Type; !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Attributes = all.Attributes
	o.Id = *all.Id
	o.Type = *all.Type
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
