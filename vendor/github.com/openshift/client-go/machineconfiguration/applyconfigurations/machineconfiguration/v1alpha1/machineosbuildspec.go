// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// MachineOSBuildSpecApplyConfiguration represents an declarative configuration of the MachineOSBuildSpec type for use
// with apply.
type MachineOSBuildSpecApplyConfiguration struct {
	ConfigGeneration      *int64                                            `json:"configGeneration,omitempty"`
	DesiredConfig         *RenderedMachineConfigReferenceApplyConfiguration `json:"desiredConfig,omitempty"`
	MachineOSConfig       *MachineOSConfigReferenceApplyConfiguration       `json:"machineOSConfig,omitempty"`
	Version               *int64                                            `json:"version,omitempty"`
	RenderedImagePushspec *string                                           `json:"renderedImagePushspec,omitempty"`
}

// MachineOSBuildSpecApplyConfiguration constructs an declarative configuration of the MachineOSBuildSpec type for use with
// apply.
func MachineOSBuildSpec() *MachineOSBuildSpecApplyConfiguration {
	return &MachineOSBuildSpecApplyConfiguration{}
}

// WithConfigGeneration sets the ConfigGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ConfigGeneration field is set to the value of the last call.
func (b *MachineOSBuildSpecApplyConfiguration) WithConfigGeneration(value int64) *MachineOSBuildSpecApplyConfiguration {
	b.ConfigGeneration = &value
	return b
}

// WithDesiredConfig sets the DesiredConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DesiredConfig field is set to the value of the last call.
func (b *MachineOSBuildSpecApplyConfiguration) WithDesiredConfig(value *RenderedMachineConfigReferenceApplyConfiguration) *MachineOSBuildSpecApplyConfiguration {
	b.DesiredConfig = value
	return b
}

// WithMachineOSConfig sets the MachineOSConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MachineOSConfig field is set to the value of the last call.
func (b *MachineOSBuildSpecApplyConfiguration) WithMachineOSConfig(value *MachineOSConfigReferenceApplyConfiguration) *MachineOSBuildSpecApplyConfiguration {
	b.MachineOSConfig = value
	return b
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *MachineOSBuildSpecApplyConfiguration) WithVersion(value int64) *MachineOSBuildSpecApplyConfiguration {
	b.Version = &value
	return b
}

// WithRenderedImagePushspec sets the RenderedImagePushspec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RenderedImagePushspec field is set to the value of the last call.
func (b *MachineOSBuildSpecApplyConfiguration) WithRenderedImagePushspec(value string) *MachineOSBuildSpecApplyConfiguration {
	b.RenderedImagePushspec = &value
	return b
}
