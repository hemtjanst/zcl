// An interface for configuring and controlling pumps.
package hvac

import (
	"neotor.se/zcl/cluster/zcl"
)

// PumpConfigurationAndControl
// An interface for configuring and controlling pumps.

func NewPumpConfigurationAndControlServer(profile zcl.ProfileID) *PumpConfigurationAndControlServer {
	return &PumpConfigurationAndControlServer{p: profile}
}
func NewPumpConfigurationAndControlClient(profile zcl.ProfileID) *PumpConfigurationAndControlClient {
	return &PumpConfigurationAndControlClient{p: profile}
}

const PumpConfigurationAndControlCluster zcl.ClusterID = 512

type PumpConfigurationAndControlServer struct {
	p zcl.ProfileID
}

type PumpConfigurationAndControlClient struct {
	p zcl.ProfileID
}

/*
var PumpConfigurationAndControlServer = map[zcl.CommandID]func() zcl.Command{
}

var PumpConfigurationAndControlClient = map[zcl.CommandID]func() zcl.Command{
}
*/
