package models

// EncapsulationPrioritiesType

import "encoding/json"

// EncapsulationPrioritiesType
type EncapsulationPrioritiesType struct {
	Encapsulation EncapsulationType `json:"encapsulation"`
}

// String returns json representation of the object
func (model *EncapsulationPrioritiesType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeEncapsulationPrioritiesType makes EncapsulationPrioritiesType
func MakeEncapsulationPrioritiesType() *EncapsulationPrioritiesType {
	return &EncapsulationPrioritiesType{
		//TODO(nati): Apply default

		Encapsulation: MakeEncapsulationType(),
	}
}

// InterfaceToEncapsulationPrioritiesType makes EncapsulationPrioritiesType from interface
func InterfaceToEncapsulationPrioritiesType(iData interface{}) *EncapsulationPrioritiesType {
	data := iData.(map[string]interface{})
	return &EncapsulationPrioritiesType{

		Encapsulation: InterfaceToEncapsulationType(data["encapsulation"]),

		//{"description":"Ordered list of encapsulation types to be used in priority","type":"array"}

	}
}

// InterfaceToEncapsulationPrioritiesTypeSlice makes a slice of EncapsulationPrioritiesType from interface
func InterfaceToEncapsulationPrioritiesTypeSlice(data interface{}) []*EncapsulationPrioritiesType {
	list := data.([]interface{})
	result := MakeEncapsulationPrioritiesTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToEncapsulationPrioritiesType(item))
	}
	return result
}

// MakeEncapsulationPrioritiesTypeSlice() makes a slice of EncapsulationPrioritiesType
func MakeEncapsulationPrioritiesTypeSlice() []*EncapsulationPrioritiesType {
	return []*EncapsulationPrioritiesType{}
}
