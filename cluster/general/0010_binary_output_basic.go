package general

import (
	"neotor.se/zcl"
)

// BinaryOutputBasic

func NewBinaryOutputBasicServer(profile zcl.ProfileID) *BinaryOutputBasicServer {
	return &BinaryOutputBasicServer{p: profile}
}
func NewBinaryOutputBasicClient(profile zcl.ProfileID) *BinaryOutputBasicClient {
	return &BinaryOutputBasicClient{p: profile}
}

const BinaryOutputBasicCluster zcl.ClusterID = 16

type BinaryOutputBasicServer struct {
	p zcl.ProfileID
}

type BinaryOutputBasicClient struct {
	p zcl.ProfileID
}

/*
var BinaryOutputBasicServer = map[zcl.CommandID]func() zcl.Command{
}

var BinaryOutputBasicClient = map[zcl.CommandID]func() zcl.Command{
}
*/
