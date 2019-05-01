// Attributes and commands that provide an interface to the power profile of a device
package general

import (
	"neotor.se/zcl"
)

// PowerProfile
// Attributes and commands that provide an interface to the power profile of a device

func NewPowerProfileServer(profile zcl.ProfileID) *PowerProfileServer {
	return &PowerProfileServer{p: profile}
}
func NewPowerProfileClient(profile zcl.ProfileID) *PowerProfileClient {
	return &PowerProfileClient{p: profile}
}

const PowerProfileCluster zcl.ClusterID = 26

type PowerProfileServer struct {
	p zcl.ProfileID
}

type PowerProfileClient struct {
	p zcl.ProfileID
}

/*
var PowerProfileServer = map[zcl.CommandID]func() zcl.Command{
}

var PowerProfileClient = map[zcl.CommandID]func() zcl.Command{
}
*/
