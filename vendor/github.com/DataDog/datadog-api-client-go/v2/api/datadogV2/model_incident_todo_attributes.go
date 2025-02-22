// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTodoAttributes Incident todo's attributes.
type IncidentTodoAttributes struct {
	// Array of todo assignees.
	Assignees []IncidentTodoAssignee `json:"assignees"`
	// Timestamp when the todo was completed.
	Completed datadog.NullableString `json:"completed,omitempty"`
	// The follow-up task's content.
	Content string `json:"content"`
	// Timestamp when the todo should be completed by.
	DueDate datadog.NullableString `json:"due_date,omitempty"`
	// UUID of the incident this todo is connected to.
	IncidentId *string `json:"incident_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewIncidentTodoAttributes instantiates a new IncidentTodoAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentTodoAttributes(assignees []IncidentTodoAssignee, content string) *IncidentTodoAttributes {
	this := IncidentTodoAttributes{}
	this.Assignees = assignees
	this.Content = content
	return &this
}

// NewIncidentTodoAttributesWithDefaults instantiates a new IncidentTodoAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentTodoAttributesWithDefaults() *IncidentTodoAttributes {
	this := IncidentTodoAttributes{}
	return &this
}

// GetAssignees returns the Assignees field value.
func (o *IncidentTodoAttributes) GetAssignees() []IncidentTodoAssignee {
	if o == nil {
		var ret []IncidentTodoAssignee
		return ret
	}
	return o.Assignees
}

// GetAssigneesOk returns a tuple with the Assignees field value
// and a boolean to check if the value has been set.
func (o *IncidentTodoAttributes) GetAssigneesOk() (*[]IncidentTodoAssignee, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Assignees, true
}

// SetAssignees sets field value.
func (o *IncidentTodoAttributes) SetAssignees(v []IncidentTodoAssignee) {
	o.Assignees = v
}

// GetCompleted returns the Completed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentTodoAttributes) GetCompleted() string {
	if o == nil || o.Completed.Get() == nil {
		var ret string
		return ret
	}
	return *o.Completed.Get()
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentTodoAttributes) GetCompletedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Completed.Get(), o.Completed.IsSet()
}

// HasCompleted returns a boolean if a field has been set.
func (o *IncidentTodoAttributes) HasCompleted() bool {
	return o != nil && o.Completed.IsSet()
}

// SetCompleted gets a reference to the given datadog.NullableString and assigns it to the Completed field.
func (o *IncidentTodoAttributes) SetCompleted(v string) {
	o.Completed.Set(&v)
}

// SetCompletedNil sets the value for Completed to be an explicit nil.
func (o *IncidentTodoAttributes) SetCompletedNil() {
	o.Completed.Set(nil)
}

// UnsetCompleted ensures that no value is present for Completed, not even an explicit nil.
func (o *IncidentTodoAttributes) UnsetCompleted() {
	o.Completed.Unset()
}

// GetContent returns the Content field value.
func (o *IncidentTodoAttributes) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *IncidentTodoAttributes) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value.
func (o *IncidentTodoAttributes) SetContent(v string) {
	o.Content = v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentTodoAttributes) GetDueDate() string {
	if o == nil || o.DueDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.DueDate.Get()
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentTodoAttributes) GetDueDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DueDate.Get(), o.DueDate.IsSet()
}

// HasDueDate returns a boolean if a field has been set.
func (o *IncidentTodoAttributes) HasDueDate() bool {
	return o != nil && o.DueDate.IsSet()
}

// SetDueDate gets a reference to the given datadog.NullableString and assigns it to the DueDate field.
func (o *IncidentTodoAttributes) SetDueDate(v string) {
	o.DueDate.Set(&v)
}

// SetDueDateNil sets the value for DueDate to be an explicit nil.
func (o *IncidentTodoAttributes) SetDueDateNil() {
	o.DueDate.Set(nil)
}

// UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil.
func (o *IncidentTodoAttributes) UnsetDueDate() {
	o.DueDate.Unset()
}

// GetIncidentId returns the IncidentId field value if set, zero value otherwise.
func (o *IncidentTodoAttributes) GetIncidentId() string {
	if o == nil || o.IncidentId == nil {
		var ret string
		return ret
	}
	return *o.IncidentId
}

// GetIncidentIdOk returns a tuple with the IncidentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTodoAttributes) GetIncidentIdOk() (*string, bool) {
	if o == nil || o.IncidentId == nil {
		return nil, false
	}
	return o.IncidentId, true
}

// HasIncidentId returns a boolean if a field has been set.
func (o *IncidentTodoAttributes) HasIncidentId() bool {
	return o != nil && o.IncidentId != nil
}

// SetIncidentId gets a reference to the given string and assigns it to the IncidentId field.
func (o *IncidentTodoAttributes) SetIncidentId(v string) {
	o.IncidentId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentTodoAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["assignees"] = o.Assignees
	if o.Completed.IsSet() {
		toSerialize["completed"] = o.Completed.Get()
	}
	toSerialize["content"] = o.Content
	if o.DueDate.IsSet() {
		toSerialize["due_date"] = o.DueDate.Get()
	}
	if o.IncidentId != nil {
		toSerialize["incident_id"] = o.IncidentId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentTodoAttributes) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Assignees  *[]IncidentTodoAssignee `json:"assignees"`
		Completed  datadog.NullableString  `json:"completed,omitempty"`
		Content    *string                 `json:"content"`
		DueDate    datadog.NullableString  `json:"due_date,omitempty"`
		IncidentId *string                 `json:"incident_id,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if all.Assignees == nil {
		return fmt.Errorf("required field assignees missing")
	}
	if all.Content == nil {
		return fmt.Errorf("required field content missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assignees", "completed", "content", "due_date", "incident_id"})
	} else {
		return err
	}
	o.Assignees = *all.Assignees
	o.Completed = all.Completed
	o.Content = *all.Content
	o.DueDate = all.DueDate
	o.IncidentId = all.IncidentId
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
