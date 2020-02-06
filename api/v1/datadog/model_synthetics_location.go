/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
)

// SyntheticsLocation struct for SyntheticsLocation
type SyntheticsLocation struct {
	DisplayName *string `json:"display_name,omitempty"`
	Id          *int64  `json:"id,omitempty"`
	IsActive    *bool   `json:"is_active,omitempty"`
	Name        *string `json:"name,omitempty"`
	Region      *string `json:"region,omitempty"`
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SyntheticsLocation) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsLocation) GetDisplayNameOk() (string, bool) {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret, false
	}
	return *o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SyntheticsLocation) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SyntheticsLocation) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SyntheticsLocation) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsLocation) GetIdOk() (int64, bool) {
	if o == nil || o.Id == nil {
		var ret int64
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SyntheticsLocation) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *SyntheticsLocation) SetId(v int64) {
	o.Id = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *SyntheticsLocation) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsLocation) GetIsActiveOk() (bool, bool) {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret, false
	}
	return *o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *SyntheticsLocation) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *SyntheticsLocation) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SyntheticsLocation) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsLocation) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SyntheticsLocation) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SyntheticsLocation) SetName(v string) {
	o.Name = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *SyntheticsLocation) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsLocation) GetRegionOk() (string, bool) {
	if o == nil || o.Region == nil {
		var ret string
		return ret, false
	}
	return *o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *SyntheticsLocation) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *SyntheticsLocation) SetRegion(v string) {
	o.Region = &v
}

type NullableSyntheticsLocation struct {
	Value        SyntheticsLocation
	ExplicitNull bool
}

func (v NullableSyntheticsLocation) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableSyntheticsLocation) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}