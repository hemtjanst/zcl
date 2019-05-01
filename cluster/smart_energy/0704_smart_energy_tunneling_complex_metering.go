package smart_energy

import (
	"neotor.se/zcl"
)

// SmartEnergyTunnelingComplexMetering
const SmartEnergyTunnelingComplexMeteringID zcl.ClusterID = 1796

var SmartEnergyTunnelingComplexMeteringCluster = zcl.Cluster{
	ServerCmd:  map[zcl.CommandID]func() zcl.Command{},
	ClientCmd:  map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
