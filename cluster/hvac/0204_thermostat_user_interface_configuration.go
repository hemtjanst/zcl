// sdfsdf
package hvac

import (
	"neotor.se/zcl/cluster/zcl"
)

// ThermostatUserInterfaceConfiguration
// sdfsdf

func NewThermostatUserInterfaceConfigurationServer(profile zcl.ProfileID) *ThermostatUserInterfaceConfigurationServer {
	return &ThermostatUserInterfaceConfigurationServer{p: profile}
}
func NewThermostatUserInterfaceConfigurationClient(profile zcl.ProfileID) *ThermostatUserInterfaceConfigurationClient {
	return &ThermostatUserInterfaceConfigurationClient{p: profile}
}

const ThermostatUserInterfaceConfigurationCluster zcl.ClusterID = 516

type ThermostatUserInterfaceConfigurationServer struct {
	p zcl.ProfileID
}

type ThermostatUserInterfaceConfigurationClient struct {
	p zcl.ProfileID
}

/*
var ThermostatUserInterfaceConfigurationServer = map[zcl.CommandID]func() zcl.Command{
}

var ThermostatUserInterfaceConfigurationClient = map[zcl.CommandID]func() zcl.Command{
}
*/
