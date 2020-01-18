package utils

import (
	"hemtjan.st/zcl"
	"hemtjan.st/zcl/cluster"
)

func FindAttr(clusterID zcl.ClusterID, attrID zcl.AttrID) (attr zcl.Attr) {
	cl, ok := cluster.Clusters[clusterID]
	if !ok {
		return nil
	}
	if attrFn, ok := cl.ServerAttr[attrID]; ok {
		return attrFn()
	}
	if attrFn, ok := cl.ClientAttr[attrID]; ok {
		return attrFn()
	}
	return nil
}
