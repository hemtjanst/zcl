package general

import (
	"neotor.se/zcl"
)

// BinaryValueBasic

func NewBinaryValueBasicServer(profile zcl.ProfileID) *BinaryValueBasicServer {
	return &BinaryValueBasicServer{p: profile}
}
func NewBinaryValueBasicClient(profile zcl.ProfileID) *BinaryValueBasicClient {
	return &BinaryValueBasicClient{p: profile}
}

const BinaryValueBasicCluster zcl.ClusterID = 17

type BinaryValueBasicServer struct {
	p zcl.ProfileID
}

type BinaryValueBasicClient struct {
	p zcl.ProfileID
}

/*
var BinaryValueBasicServer = map[zcl.CommandID]func() zcl.Command{
}

var BinaryValueBasicClient = map[zcl.CommandID]func() zcl.Command{
}
*/
