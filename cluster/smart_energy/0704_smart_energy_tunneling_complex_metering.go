package smart_energy

import (
	"neotor.se/zcl"
)

// SmartEnergyTunnelingComplexMetering

func NewSmartEnergyTunnelingComplexMeteringServer(profile zcl.ProfileID) *SmartEnergyTunnelingComplexMeteringServer {
	return &SmartEnergyTunnelingComplexMeteringServer{p: profile}
}
func NewSmartEnergyTunnelingComplexMeteringClient(profile zcl.ProfileID) *SmartEnergyTunnelingComplexMeteringClient {
	return &SmartEnergyTunnelingComplexMeteringClient{p: profile}
}

const SmartEnergyTunnelingComplexMeteringCluster zcl.ClusterID = 1796

type SmartEnergyTunnelingComplexMeteringServer struct {
	p zcl.ProfileID
}

type SmartEnergyTunnelingComplexMeteringClient struct {
	p zcl.ProfileID
}

/*
var SmartEnergyTunnelingComplexMeteringServer = map[zcl.CommandID]func() zcl.Command{
}

var SmartEnergyTunnelingComplexMeteringClient = map[zcl.CommandID]func() zcl.Command{
}
*/
