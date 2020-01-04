package device

import (
	"fmt"
	"hemtjan.st/zcl"
	"hemtjan.st/zcl/cluster/general"
)

func (z *Endpoint) AttrValue(cluster zcl.ClusterID, attr zcl.AttrID) (zcl.Val, bool) {
	if cl, ok := z.attr[cluster]; ok {
		for _, a := range cl {
			if a.ID() == attr {
				return a.Value(), true
			}
		}
	}
	return nil, false
}
func (z *Endpoint) AttrString(cluster zcl.ClusterID, attr zcl.AttrID) string {
	if v, ok := z.AttrValue(cluster, attr); ok {
		return fmt.Sprintf("%v", v)
	}
	return ""
}

func (z *Endpoint) Name() string { return z.AttrString(general.BasicID, general.ModelIdentifierAttr) }
func (z *Endpoint) Manufacturer() string {
	return z.AttrString(general.BasicID, general.ManufacturerNameAttr)
}
