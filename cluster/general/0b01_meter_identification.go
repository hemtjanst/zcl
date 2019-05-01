// Attributes and commands that provide an interface tometer identification
package general

import (
	"neotor.se/zcl/cluster/zcl"
)

// MeterIdentification
// Attributes and commands that provide an interface tometer identification

func NewMeterIdentificationServer(profile zcl.ProfileID) *MeterIdentificationServer {
	return &MeterIdentificationServer{p: profile}
}
func NewMeterIdentificationClient(profile zcl.ProfileID) *MeterIdentificationClient {
	return &MeterIdentificationClient{p: profile}
}

const MeterIdentificationCluster zcl.ClusterID = 2817

type MeterIdentificationServer struct {
	p zcl.ProfileID
}

type MeterIdentificationClient struct {
	p zcl.ProfileID
}

/*
var MeterIdentificationServer = map[zcl.CommandID]func() zcl.Command{
}

var MeterIdentificationClient = map[zcl.CommandID]func() zcl.Command{
}
*/
