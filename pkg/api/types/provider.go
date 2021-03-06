package types

import "fmt"

var (
	// the set of valid provider type strings that may appear in the provider's
	// "provider_type" field
	ValidProviderTypes = []string{
		"runm.compute",
		"runm.storage.block",
	}
)

// Provider is an object that exposes inventories of consumable resources and
// has a set of capabilities
type Provider struct {
	// Identifier of the partition the object belongs to
	Partition string `json:"partition"`
	// Code for the type of provider this is
	ProviderType string `json:"provider_type"`
	// Optional identifier of the provider the provider is a child of. Leave
	// empty if the provider has no parents (it's a "root provider")
	Parent string `json:"parent,omitempty"`
	// The UUID of the provider. Expected to be blank when a user is creating a
	// new provider.
	Uuid string `json:"uuid,omitempty"`
	// Human-readable name for the provider. Uniqueness is guaranteed in the
	// scope of the partition the provider belongs to.
	Name string `json:"name"`
	// Map of key/value properties associated with this provider. Properties can
	// have a structure and be validated against a schema.
	Properties map[string]interface{} `json:"properties,omitempty"`
	// Array of string tags. Tags are unstructured and unvalidated and any user
	// with write access to the provider can add or remove any tag.
	Tags []string `json:"tags,omitempty"`
}

// Validate returns an error if the definition is invalid, nil otherwise
func (p *Provider) Validate() error {
	if p.Partition == "" {
		return fmt.Errorf("partition required")
	}
	if p.Name == "" {
		return fmt.Errorf("name required")
	}
	if p.ProviderType == "" {
		return fmt.Errorf("provider_type required")
	} else {
		found := false
		for _, pt := range ValidProviderTypes {
			if p.ProviderType == pt {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf(
				"invalid provider_type %s. valid choices: %s",
				p.ProviderType, ValidProviderTypes,
			)
		}
	}
	return nil
}
