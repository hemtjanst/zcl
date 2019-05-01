// An interface for configuring and controlling pumps.
package hvac

import (
	"neotor.se/zcl"
)

// PumpConfigurationAndControl
const PumpConfigurationAndControlID zcl.ClusterID = 512

var PumpConfigurationAndControlCluster = zcl.Cluster{
	ServerCmd:  map[zcl.CommandID]func() zcl.Command{},
	ClientCmd:  map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
