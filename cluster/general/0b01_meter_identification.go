package general

import "hemtjan.st/zcl"

// MeterIdentification
const MeterIdentificationID zcl.ClusterID = 2817

var MeterIdentificationCluster = zcl.Cluster{
	Name:       "Meter Identification",
	ServerCmd:  map[zcl.CommandID]func() zcl.Command{},
	ClientCmd:  map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
