/*
Apicurio Registry API [v2]

Apicurio Registry is a datastore for standard event schemas and API designs. Apicurio Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.  The Apicurio Registry REST API enables client applications to manage the artifacts in the registry. This API provides create, read, update, and delete operations for schema and API artifacts, rules, versions, and metadata.   The supported artifact types include: - Apache Avro schema - AsyncAPI specification - Google protocol buffers - GraphQL schema - JSON Schema - Kafka Connect schema - OpenAPI specification - Web Services Description Language - XML Schema Definition   **Important**: The Apicurio Registry REST API is available from `https://MY-REGISTRY-URL/apis/registry/v2` by default. Therefore you must prefix all API operation paths with `../apis/registry/v2` in this case. For example: `../apis/registry/v2/ids/globalIds/{globalId}`. 

API version: 2.3.2-SNAPSHOT
Contact: apicurio@lists.jboss.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registryclient

import (
	"encoding/json"
)

// ConfigurationProperty 
type ConfigurationProperty struct {
	Name string `json:"name"`
	Value string `json:"value"`
	// 
	Type string `json:"type"`
	// 
	Label string `json:"label"`
	// 
	Description string `json:"description"`
}

// NewConfigurationProperty instantiates a new ConfigurationProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationProperty(name string, value string, type_ string, label string, description string) *ConfigurationProperty {
	this := ConfigurationProperty{}
	this.Name = name
	this.Value = value
	this.Type = type_
	this.Label = label
	this.Description = description
	return &this
}

// NewConfigurationPropertyWithDefaults instantiates a new ConfigurationProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationPropertyWithDefaults() *ConfigurationProperty {
	this := ConfigurationProperty{}
	return &this
}

// GetName returns the Name field value
func (o *ConfigurationProperty) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConfigurationProperty) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConfigurationProperty) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *ConfigurationProperty) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ConfigurationProperty) GetValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ConfigurationProperty) SetValue(v string) {
	o.Value = v
}

// GetType returns the Type field value
func (o *ConfigurationProperty) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ConfigurationProperty) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ConfigurationProperty) SetType(v string) {
	o.Type = v
}

// GetLabel returns the Label field value
func (o *ConfigurationProperty) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ConfigurationProperty) GetLabelOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *ConfigurationProperty) SetLabel(v string) {
	o.Label = v
}

// GetDescription returns the Description field value
func (o *ConfigurationProperty) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ConfigurationProperty) GetDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ConfigurationProperty) SetDescription(v string) {
	o.Description = v
}

func (o ConfigurationProperty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableConfigurationProperty struct {
	value *ConfigurationProperty
	isSet bool
}

func (v NullableConfigurationProperty) Get() *ConfigurationProperty {
	return v.value
}

func (v *NullableConfigurationProperty) Set(val *ConfigurationProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationProperty(val *ConfigurationProperty) *NullableConfigurationProperty {
	return &NullableConfigurationProperty{value: val, isSet: true}
}

func (v NullableConfigurationProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


