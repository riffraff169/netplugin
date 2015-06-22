package drivers

import "github.com/contiv/netplugin/core"

// FakeNetEpDriverConfig represents the configuration of the fakedriver,
// which is an empty struct.
type FakeNetEpDriverConfig struct{}

// FakeNetEpDriver implements core.NetworkDriver and core.EndpointDriver interface
// for use with unit-tests
type FakeNetEpDriver struct {
}

// Init is not implemented.
func (d *FakeNetEpDriver) Init(config *core.Config, info *core.InstanceInfo) error {
	return nil
}

// Deinit is not implemented.
func (d *FakeNetEpDriver) Deinit() {
}

// CreateNetwork is not implemented.
func (d *FakeNetEpDriver) CreateNetwork(id string) error {
	return core.Errorf("Not implemented")
}

// DeleteNetwork is not implemented.
func (d *FakeNetEpDriver) DeleteNetwork(id string) error {
	return core.Errorf("Not implemented")
}

// CreateEndpoint is not implemented.
func (d *FakeNetEpDriver) CreateEndpoint(id string) error {
	return core.Errorf("Not implemented")
}

// DeleteEndpoint is not implemented.
func (d *FakeNetEpDriver) DeleteEndpoint(id string) (err error) {
	return core.Errorf("Not implemented")
}

// MakeEndpointAddress is not implemented.
func (d *FakeNetEpDriver) MakeEndpointAddress() (*core.Address, error) {
	return nil, core.Errorf("Not implemented")
}