package models

// PeeringPolicy

import "encoding/json"

// PeeringPolicy
type PeeringPolicy struct {
	PeeringService PeeringServiceType `json:"peering_service"`
	Annotations    *KeyValuePairs     `json:"annotations"`
	ParentType     string             `json:"parent_type"`
	FQName         []string           `json:"fq_name"`
	IDPerms        *IdPermsType       `json:"id_perms"`
	DisplayName    string             `json:"display_name"`
	Perms2         *PermType2         `json:"perms2"`
	UUID           string             `json:"uuid"`
	ParentUUID     string             `json:"parent_uuid"`
}

// String returns json representation of the object
func (model *PeeringPolicy) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakePeeringPolicy makes PeeringPolicy
func MakePeeringPolicy() *PeeringPolicy {
	return &PeeringPolicy{
		//TODO(nati): Apply default
		PeeringService: MakePeeringServiceType(),
		Annotations:    MakeKeyValuePairs(),
		UUID:           "",
		ParentUUID:     "",
		ParentType:     "",
		FQName:         []string{},
		IDPerms:        MakeIdPermsType(),
		DisplayName:    "",
		Perms2:         MakePermType2(),
	}
}

// InterfaceToPeeringPolicy makes PeeringPolicy from interface
func InterfaceToPeeringPolicy(iData interface{}) *PeeringPolicy {
	data := iData.(map[string]interface{})
	return &PeeringPolicy{
		PeeringService: InterfaceToPeeringServiceType(data["peering_service"]),

		//{"description":"Peering policy service type.","type":"string","enum":["public-peering"]}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}

	}
}

// InterfaceToPeeringPolicySlice makes a slice of PeeringPolicy from interface
func InterfaceToPeeringPolicySlice(data interface{}) []*PeeringPolicy {
	list := data.([]interface{})
	result := MakePeeringPolicySlice()
	for _, item := range list {
		result = append(result, InterfaceToPeeringPolicy(item))
	}
	return result
}

// MakePeeringPolicySlice() makes a slice of PeeringPolicy
func MakePeeringPolicySlice() []*PeeringPolicy {
	return []*PeeringPolicy{}
}
