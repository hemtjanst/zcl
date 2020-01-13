package device

import (
	"fmt"
	"hemtjan.st/zcl"
	"hemtjan.st/zcl/cluster"
	"hemtjan.st/zcl/foundation"
	"hemtjan.st/zcl/zdo"
	"log"
	"sort"
)

func (z *Endpoint) ScanClusters() error {

	cRsp, err := z.dev.Zdo().Request(&zdo.ComplexDescRequest{
		NwkAddress: zdo.NwkAddress(z.dev.nwk),
	})
	if err != nil {
		return err
	}
	cDesc, ok := cRsp.(*zdo.ComplexDescResponse)
	if !ok {
		return fmt.Errorf("invalid response to ComplexDescRequest: %+v", cRsp)
	}
	log.Printf("Complex Desc: %s (%X)", string(cDesc.ComplexDescriptor), []byte(cDesc.ComplexDescriptor))

	rsp, err := z.dev.Zdo().Request(&zdo.SimpleDescRequest{
		NwkAddress: zdo.NwkAddress(z.dev.nwk),
		Endpoint:   zdo.Endpoint(z.ep),
	})
	if err != nil {
		return err
	}
	desc, ok := rsp.(*zdo.SimpleDescResponse)
	if !ok {
		return fmt.Errorf("invalid response to SimpleDescRequest: %+v", rsp)
	}

	sd := desc.SimpleDescriptor

	log.Printf("Simple Descriptor: %s", sd)
	z.profile = uint16(sd.ProfileId)
	z.devType = zdo.DeviceType(sd.DeviceType)
	z.devVersion = uint8(sd.DeviceVersion)

	z.inCluster = []zcl.ClusterID{}
	for _, cl := range sd.InClusterList {
		z.inCluster = append(z.inCluster, zcl.ClusterID(*cl))
	}
	z.outCluster = []zcl.ClusterID{}
	for _, cl := range sd.OutClusterList {
		z.outCluster = append(z.outCluster, zcl.ClusterID(*cl))
	}

	return nil
}

func (z *Endpoint) Clusters() []zcl.ClusterID {
	var clusters []zcl.ClusterID
NextIn:
	for _, cl := range z.inCluster {
		for _, m := range clusters {
			if m == cl {
				continue NextIn
			}
		}
		clusters = append(clusters, cl)
	}
NextOut:
	for _, cl := range z.outCluster {
		for _, m := range clusters {
			if m == cl {
				continue NextOut
			}
		}
		clusters = append(clusters, cl)
	}
	sort.Slice(clusters, func(i, j int) bool {
		return clusters[i] < clusters[j]
	})
	return clusters
}

func (z *Endpoint) ScanAttributes() error {
	readRsp := map[zcl.ClusterID][]*foundation.ReadAttributesResponse{}

nextCluster:
	for _, cl := range z.Clusters() {
		var clAttr []zcl.Attr
		idx := 0
		readRsp[cl] = []*foundation.ReadAttributesResponse{}
		for {
			rsp, err := z.General(uint16(cl), &foundation.DiscoverAttributes{StartIndex: zcl.Zu16(idx), MaxEntries: 32})
			if err != nil {
				log.Printf("Error Discovering Attributes for cluster %d: %s", cl, err)
				if len(clAttr) > 0 {
					z.attr[cl] = clAttr
				}
				continue nextCluster
			}
			log.Printf("%#v", rsp)
			if attrRsp, ok := rsp.(*foundation.DiscoverAttributesResponse); ok {

				var discAttr []zcl.Attr
				readReq := &foundation.ReadAttributes{}

				for _, attr := range attrRsp.Attributes {
					av, err := FindAttr(cl, attr)
					if err != nil {
						log.Printf("Error Finding Attribute 0x%04X/0x%04X: %s", cl, attr.ID, err)
						continue
					}
					discAttr = append(discAttr, av)
					readReq.AttributeList = append(readReq.AttributeList, attr.ID)
					attrId := int(attr.ID)
					if attrId >= idx {
						idx = attrId + 1
					}
				}

				maxAttempt := 5
				for len(readReq.AttributeList) > 0 {
					maxAttempt--
					if maxAttempt <= 0 {
						break
					}
					rsp, err := z.General(uint16(cl), readReq)
					if err != nil {
						log.Printf("Unable to read attributes (%#v): %s", readReq, err)
					} else {
						if readResp, ok := rsp.(*foundation.ReadAttributesResponse); ok {
						nextAttr:
							for _, v := range readResp.Attributes {
								if v.Status != zcl.Success {
									log.Printf("Error Reading Attribute 0x%04X/0x%04X: %s", cl, v.AttrID, v.Status.String())
									continue
								}

								chop := 0
								for i, vv := range readReq.AttributeList {
									if vv <= v.AttributeID {
										chop = i
									}
								}
								if chop > 0 {
									readReq.AttributeList = readReq.AttributeList[chop:]
								}

								for _, a := range discAttr {
									if a.ID() != v.AttrID() {
										continue
									}
									err := a.SetValue(v.Value)
									if err != nil {
										log.Printf("Error Writing local Attribute 0x%04X/0x%04X: %s", cl, v.AttrID, err)
									}
									continue nextAttr
								}
							}
						} else {
							log.Printf("Invalid response to Read Attributes: %#v", rsp)
						}
					}
				}
				clAttr = append(clAttr, discAttr...)

				if attrRsp.DiscoveryComplete == 1 {
					if len(clAttr) > 0 {
						z.attr[cl] = clAttr
					}
					continue nextCluster
				}

			} else {
				if len(clAttr) > 0 {
					z.attr[cl] = clAttr
				}
				log.Printf("Invalid response to Discover Attributes: %#v", rsp)
				continue nextCluster
			}
		}
	}
	return nil
}

func (z *Endpoint) ScanCommandsGenerated() error {
	outCommand := map[zcl.ClusterID][]func() zcl.Command{}
	for _, cl := range z.Clusters() {
		var err error
		outCommand[cl], err = z.scanCommands(cl, func(idx int) zcl.General {
			return &foundation.DiscoverCommandsGenerated{StartIndex: zcl.Zu8(idx), MaxEntries: 32}
		})
		if err != nil {
			log.Printf("Error Discovering Generated Commands for cluster %04X: %s", cl, err)
		}
	}
	z.outCommand = outCommand
	return nil
}

func (z *Endpoint) ScanCommandsReceived() error {
	inCommand := map[zcl.ClusterID][]func() zcl.Command{}
	for _, cl := range z.Clusters() {
		var err error
		inCommand[cl], err = z.scanCommands(cl, func(idx int) zcl.General {
			return &foundation.DiscoverCommandsReceived{StartIndex: zcl.Zu8(idx), MaxEntries: 32}
		})
		if err != nil {
			log.Printf("Error Discovering Received Commands for cluster %04X: %s", cl, err)
		}
	}
	z.inCommand = inCommand
	return nil
}

func (z *Endpoint) scanCommands(cl zcl.ClusterID, cmdFn func(idx int) zcl.General) (out []func() zcl.Command, err error) {
	clInfo, ok := cluster.Clusters[cl]
	if !ok {
		return nil, fmt.Errorf("cluster %04X not supported", cl)
	}
	idx := 0
	for {
		rsp, err := z.General(uint16(cl), cmdFn(idx))
		if err != nil {
			return out, err
		}
		var foundCommands []zcl.CommandID
		var discoveryComplete bool
		if cmdRsp, ok := rsp.(*foundation.DiscoverCommandsGeneratedResponse); ok {
			foundCommands = cmdRsp.Commands
			discoveryComplete = cmdRsp.DiscoveryComplete == 1
		} else if cmdRsp, ok := rsp.(*foundation.DiscoverCommandsReceivedResponse); ok {
			foundCommands = cmdRsp.Commands
			discoveryComplete = cmdRsp.DiscoveryComplete == 1
		} else if cmdRsp, ok := rsp.(*zcl.DefaultResponse); ok {
			return out, fmt.Errorf("%s", cmdRsp.Status.String())
		} else {
			return out, fmt.Errorf("invalid response: %#v", rsp)
		}

		for _, id := range foundCommands {
			cmd := clInfo.CmdFn(uint8(id))
			if cmd != nil {
				out = append(out, cmd)
			}
			if int(id) >= idx {
				idx = int(id) + 1
			}
		}
		if discoveryComplete {
			break
		}
	}
	return out, nil
}
