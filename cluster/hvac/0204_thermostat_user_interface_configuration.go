// sdfsdf
package hvac

import (
	"hemtjan.st/zcl"
)

// ThermostatUserInterfaceConfiguration
const ThermostatUserInterfaceConfigurationID zcl.ClusterID = 516

var ThermostatUserInterfaceConfigurationCluster = zcl.Cluster{
	ServerCmd:  map[zcl.CommandID]func() zcl.Command{},
	ClientCmd:  map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
