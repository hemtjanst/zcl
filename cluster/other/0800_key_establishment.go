package other

import (
	"neotor.se/zcl"
)

// KeyEstablishment

func NewKeyEstablishmentServer(profile zcl.ProfileID) *KeyEstablishmentServer {
	return &KeyEstablishmentServer{p: profile}
}
func NewKeyEstablishmentClient(profile zcl.ProfileID) *KeyEstablishmentClient {
	return &KeyEstablishmentClient{p: profile}
}

const KeyEstablishmentCluster zcl.ClusterID = 2048

type KeyEstablishmentServer struct {
	p zcl.ProfileID
}

type KeyEstablishmentClient struct {
	p zcl.ProfileID
}

/*
var KeyEstablishmentServer = map[zcl.CommandID]func() zcl.Command{
}

var KeyEstablishmentClient = map[zcl.CommandID]func() zcl.Command{
}
*/
