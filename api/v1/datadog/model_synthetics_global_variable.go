/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"time"
)

// SyntheticsGlobalVariable Synthetics global variable.
type SyntheticsGlobalVariable struct {
	// Creation timestamp of the global variable.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The ID of the user who created the global variable.
	CreatedBy *int32 `json:"created_by,omitempty"`
	// Description of the global variable.
	Description *string `json:"description,omitempty"`
	// Unique identifier of the global variable.
	Id *string `json:"id,omitempty"`
	// Modification timestamp of the global variable.
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// Name of the global variable.
	Name *string `json:"name,omitempty"`
	// Tags of the global variable.
	Tags  *[]string                      `json:"tags,omitempty"`
	Value *SyntheticsGlobalVariableValue `json:"value,omitempty"`
}

// NewSyntheticsGlobalVariable instantiates a new SyntheticsGlobalVariable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsGlobalVariable() *SyntheticsGlobalVariable {
	this := SyntheticsGlobalVariable{}
	return &this
}

// NewSyntheticsGlobalVariableWithDefaults instantiates a new SyntheticsGlobalVariable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsGlobalVariableWithDefaults() *SyntheticsGlobalVariable {
	this := SyntheticsGlobalVariable{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SyntheticsGlobalVariable) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsGlobalVariable) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SyntheticsGlobalVariable) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SyntheticsGlobalVariable) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *SyntheticsGlobalVariable) GetCreatedBy() int32 {
	if o == nil || o.CreatedBy == nil {
		var ret int32
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsGlobalVariable) GetCreatedByOk() (*int32, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *SyntheticsGlobalVariable) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given int32 and assigns it to the CreatedBy field.
func (o *SyntheticsGlobalVariable) SetCreatedBy(v int32) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SyntheticsGlobalVariable) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsGlobalVariable) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SyntheticsGlobalVariable) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SyntheticsGlobalVariable) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SyntheticsGlobalVariable) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsGlobalVariable) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SyntheticsGlobalVariable) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SyntheticsGlobalVariable) SetId(v string) {
	o.Id = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *SyntheticsGlobalVariable) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsGlobalVariable) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *SyntheticsGlobalVariable) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *SyntheticsGlobalVariable) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SyntheticsGlobalVariable) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsGlobalVariable) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SyntheticsGlobalVariable) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SyntheticsGlobalVariable) SetName(v string) {
	o.Name = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SyntheticsGlobalVariable) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsGlobalVariable) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SyntheticsGlobalVariable) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SyntheticsGlobalVariable) SetTags(v []string) {
	o.Tags = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SyntheticsGlobalVariable) GetValue() SyntheticsGlobalVariableValue {
	if o == nil || o.Value == nil {
		var ret SyntheticsGlobalVariableValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsGlobalVariable) GetValueOk() (*SyntheticsGlobalVariableValue, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SyntheticsGlobalVariable) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given SyntheticsGlobalVariableValue and assigns it to the Value field.
func (o *SyntheticsGlobalVariable) SetValue(v SyntheticsGlobalVariableValue) {
	o.Value = &v
}

func (o SyntheticsGlobalVariable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ModifiedAt != nil {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticsGlobalVariable struct {
	value *SyntheticsGlobalVariable
	isSet bool
}

func (v NullableSyntheticsGlobalVariable) Get() *SyntheticsGlobalVariable {
	return v.value
}

func (v *NullableSyntheticsGlobalVariable) Set(val *SyntheticsGlobalVariable) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsGlobalVariable) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsGlobalVariable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsGlobalVariable(val *SyntheticsGlobalVariable) *NullableSyntheticsGlobalVariable {
	return &NullableSyntheticsGlobalVariable{value: val, isSet: true}
}

func (v NullableSyntheticsGlobalVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsGlobalVariable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
