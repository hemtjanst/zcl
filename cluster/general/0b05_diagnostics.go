package general

import "hemtjan.st/zcl"

// Diagnostics
const DiagnosticsID zcl.ClusterID = 2821

var DiagnosticsCluster = zcl.Cluster{
	Name:      "Diagnostics",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		NumberOfResetsAttr:            func() zcl.Attr { return new(NumberOfResets) },
		PersistensMemoryWritesAttr:    func() zcl.Attr { return new(PersistensMemoryWrites) },
		MacRxBcastAttr:                func() zcl.Attr { return new(MacRxBcast) },
		MacTxBcastAttr:                func() zcl.Attr { return new(MacTxBcast) },
		MacRxUcastAttr:                func() zcl.Attr { return new(MacRxUcast) },
		MacTxUcastAttr:                func() zcl.Attr { return new(MacTxUcast) },
		MacTxUcastRetryAttr:           func() zcl.Attr { return new(MacTxUcastRetry) },
		MacTxUcastFailAttr:            func() zcl.Attr { return new(MacTxUcastFail) },
		ApsRxBcastAttr:                func() zcl.Attr { return new(ApsRxBcast) },
		ApsTxBcastAttr:                func() zcl.Attr { return new(ApsTxBcast) },
		ApsRxUcastAttr:                func() zcl.Attr { return new(ApsRxUcast) },
		ApsTxUcastSuccessAttr:         func() zcl.Attr { return new(ApsTxUcastSuccess) },
		ApsTxUcastRetryAttr:           func() zcl.Attr { return new(ApsTxUcastRetry) },
		ApsTxUcastFailAttr:            func() zcl.Attr { return new(ApsTxUcastFail) },
		RouteDiscInitiatedAttr:        func() zcl.Attr { return new(RouteDiscInitiated) },
		NeighborAddedAttr:             func() zcl.Attr { return new(NeighborAdded) },
		NeighborRemovedAttr:           func() zcl.Attr { return new(NeighborRemoved) },
		NeighborStaleAttr:             func() zcl.Attr { return new(NeighborStale) },
		JoinIndicationAttr:            func() zcl.Attr { return new(JoinIndication) },
		ChildMovedAttr:                func() zcl.Attr { return new(ChildMoved) },
		NwkFcFailureAttr:              func() zcl.Attr { return new(NwkFcFailure) },
		ApsFcFailureAttr:              func() zcl.Attr { return new(ApsFcFailure) },
		ApsUnauthorizedKeyAttr:        func() zcl.Attr { return new(ApsUnauthorizedKey) },
		NwkDecryptFailuresAttr:        func() zcl.Attr { return new(NwkDecryptFailures) },
		ApsDecryptFailuresAttr:        func() zcl.Attr { return new(ApsDecryptFailures) },
		PacketBufferAllocFailuresAttr: func() zcl.Attr { return new(PacketBufferAllocFailures) },
		RelayedUcastAttr:              func() zcl.Attr { return new(RelayedUcast) },
		PhyToMacQueueLimitReachedAttr: func() zcl.Attr { return new(PhyToMacQueueLimitReached) },
		PacketValidateDropcountAttr:   func() zcl.Attr { return new(PacketValidateDropcount) },
		AvgMacRetryPerApsMsgSentAttr:  func() zcl.Attr { return new(AvgMacRetryPerApsMsgSent) },
		LastMessageLqiAttr:            func() zcl.Attr { return new(LastMessageLqi) },
		LastMessageRssiAttr:           func() zcl.Attr { return new(LastMessageRssi) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
