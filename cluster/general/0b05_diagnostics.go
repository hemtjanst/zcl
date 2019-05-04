// The diagnostics cluster provides access to information regarding the operation of the ZigBee stack over time.
package general

import (
	"hemtjan.st/zcl"
)

// Diagnostics
const DiagnosticsID zcl.ClusterID = 2821

var DiagnosticsCluster = zcl.Cluster{
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

// NumberOfResets is an autogenerated attribute in the Diagnostics cluster
type NumberOfResets zcl.Zu16

const NumberOfResetsAttr zcl.AttrID = 0

func (NumberOfResets) ID() zcl.AttrID                { return NumberOfResetsAttr }
func (NumberOfResets) Cluster() zcl.ClusterID        { return DiagnosticsID }
func (NumberOfResets) Name() string                  { return "Number of Resets" }
func (NumberOfResets) Readable() bool                { return true }
func (NumberOfResets) Writable() bool                { return false }
func (NumberOfResets) Reportable() bool              { return false }
func (NumberOfResets) SceneIndex() int               { return -1 }
func (a *NumberOfResets) Value() *NumberOfResets     { return a }
func (a NumberOfResets) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *NumberOfResets) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = NumberOfResets(*nt)
	return br, err
}

func (a NumberOfResets) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// PersistensMemoryWrites is an autogenerated attribute in the Diagnostics cluster
type PersistensMemoryWrites zcl.Zu16

const PersistensMemoryWritesAttr zcl.AttrID = 1

func (PersistensMemoryWrites) ID() zcl.AttrID                    { return PersistensMemoryWritesAttr }
func (PersistensMemoryWrites) Cluster() zcl.ClusterID            { return DiagnosticsID }
func (PersistensMemoryWrites) Name() string                      { return "Persistens Memory Writes" }
func (PersistensMemoryWrites) Readable() bool                    { return true }
func (PersistensMemoryWrites) Writable() bool                    { return false }
func (PersistensMemoryWrites) Reportable() bool                  { return false }
func (PersistensMemoryWrites) SceneIndex() int                   { return -1 }
func (a *PersistensMemoryWrites) Value() *PersistensMemoryWrites { return a }
func (a PersistensMemoryWrites) MarshalZcl() ([]byte, error)     { return zcl.Zu16(a).MarshalZcl() }

func (a *PersistensMemoryWrites) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = PersistensMemoryWrites(*nt)
	return br, err
}

func (a PersistensMemoryWrites) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// MacRxBcast is an autogenerated attribute in the Diagnostics cluster
type MacRxBcast zcl.Zu32

const MacRxBcastAttr zcl.AttrID = 256

func (MacRxBcast) ID() zcl.AttrID                { return MacRxBcastAttr }
func (MacRxBcast) Cluster() zcl.ClusterID        { return DiagnosticsID }
func (MacRxBcast) Name() string                  { return "Mac Rx Bcast" }
func (MacRxBcast) Readable() bool                { return true }
func (MacRxBcast) Writable() bool                { return false }
func (MacRxBcast) Reportable() bool              { return false }
func (MacRxBcast) SceneIndex() int               { return -1 }
func (a *MacRxBcast) Value() *MacRxBcast         { return a }
func (a MacRxBcast) MarshalZcl() ([]byte, error) { return zcl.Zu32(a).MarshalZcl() }

func (a *MacRxBcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = MacRxBcast(*nt)
	return br, err
}

func (a MacRxBcast) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(a))
}

// MacTxBcast is an autogenerated attribute in the Diagnostics cluster
type MacTxBcast zcl.Zu32

const MacTxBcastAttr zcl.AttrID = 257

func (MacTxBcast) ID() zcl.AttrID                { return MacTxBcastAttr }
func (MacTxBcast) Cluster() zcl.ClusterID        { return DiagnosticsID }
func (MacTxBcast) Name() string                  { return "Mac Tx Bcast" }
func (MacTxBcast) Readable() bool                { return true }
func (MacTxBcast) Writable() bool                { return false }
func (MacTxBcast) Reportable() bool              { return false }
func (MacTxBcast) SceneIndex() int               { return -1 }
func (a *MacTxBcast) Value() *MacTxBcast         { return a }
func (a MacTxBcast) MarshalZcl() ([]byte, error) { return zcl.Zu32(a).MarshalZcl() }

func (a *MacTxBcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = MacTxBcast(*nt)
	return br, err
}

func (a MacTxBcast) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(a))
}

// MacRxUcast is an autogenerated attribute in the Diagnostics cluster
type MacRxUcast zcl.Zu32

const MacRxUcastAttr zcl.AttrID = 258

func (MacRxUcast) ID() zcl.AttrID                { return MacRxUcastAttr }
func (MacRxUcast) Cluster() zcl.ClusterID        { return DiagnosticsID }
func (MacRxUcast) Name() string                  { return "Mac Rx Ucast" }
func (MacRxUcast) Readable() bool                { return true }
func (MacRxUcast) Writable() bool                { return false }
func (MacRxUcast) Reportable() bool              { return false }
func (MacRxUcast) SceneIndex() int               { return -1 }
func (a *MacRxUcast) Value() *MacRxUcast         { return a }
func (a MacRxUcast) MarshalZcl() ([]byte, error) { return zcl.Zu32(a).MarshalZcl() }

func (a *MacRxUcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = MacRxUcast(*nt)
	return br, err
}

func (a MacRxUcast) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(a))
}

// MacTxUcast is an autogenerated attribute in the Diagnostics cluster
type MacTxUcast zcl.Zu32

const MacTxUcastAttr zcl.AttrID = 259

func (MacTxUcast) ID() zcl.AttrID                { return MacTxUcastAttr }
func (MacTxUcast) Cluster() zcl.ClusterID        { return DiagnosticsID }
func (MacTxUcast) Name() string                  { return "Mac Tx Ucast" }
func (MacTxUcast) Readable() bool                { return true }
func (MacTxUcast) Writable() bool                { return false }
func (MacTxUcast) Reportable() bool              { return false }
func (MacTxUcast) SceneIndex() int               { return -1 }
func (a *MacTxUcast) Value() *MacTxUcast         { return a }
func (a MacTxUcast) MarshalZcl() ([]byte, error) { return zcl.Zu32(a).MarshalZcl() }

func (a *MacTxUcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = MacTxUcast(*nt)
	return br, err
}

func (a MacTxUcast) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(a))
}

// MacTxUcastRetry is an autogenerated attribute in the Diagnostics cluster
type MacTxUcastRetry zcl.Zu16

const MacTxUcastRetryAttr zcl.AttrID = 260

func (MacTxUcastRetry) ID() zcl.AttrID                { return MacTxUcastRetryAttr }
func (MacTxUcastRetry) Cluster() zcl.ClusterID        { return DiagnosticsID }
func (MacTxUcastRetry) Name() string                  { return "Mac Tx Ucast Retry" }
func (MacTxUcastRetry) Readable() bool                { return true }
func (MacTxUcastRetry) Writable() bool                { return false }
func (MacTxUcastRetry) Reportable() bool              { return false }
func (MacTxUcastRetry) SceneIndex() int               { return -1 }
func (a *MacTxUcastRetry) Value() *MacTxUcastRetry    { return a }
func (a MacTxUcastRetry) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *MacTxUcastRetry) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MacTxUcastRetry(*nt)
	return br, err
}

func (a MacTxUcastRetry) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// MacTxUcastFail is an autogenerated attribute in the Diagnostics cluster
type MacTxUcastFail zcl.Zu16

const MacTxUcastFailAttr zcl.AttrID = 261

func (MacTxUcastFail) ID() zcl.AttrID                { return MacTxUcastFailAttr }
func (MacTxUcastFail) Cluster() zcl.ClusterID        { return DiagnosticsID }
func (MacTxUcastFail) Name() string                  { return "Mac Tx Ucast Fail" }
func (MacTxUcastFail) Readable() bool                { return true }
func (MacTxUcastFail) Writable() bool                { return false }
func (MacTxUcastFail) Reportable() bool              { return false }
func (MacTxUcastFail) SceneIndex() int               { return -1 }
func (a *MacTxUcastFail) Value() *MacTxUcastFail     { return a }
func (a MacTxUcastFail) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *MacTxUcastFail) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MacTxUcastFail(*nt)
	return br, err
}

func (a MacTxUcastFail) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// ApsRxBcast is an autogenerated attribute in the Diagnostics cluster
type ApsRxBcast zcl.Zu16

const ApsRxBcastAttr zcl.AttrID = 262

func (ApsRxBcast) ID() zcl.AttrID                { return ApsRxBcastAttr }
func (ApsRxBcast) Cluster() zcl.ClusterID        { return DiagnosticsID }
func (ApsRxBcast) Name() string                  { return "APS Rx Bcast" }
func (ApsRxBcast) Readable() bool                { return true }
func (ApsRxBcast) Writable() bool                { return false }
func (ApsRxBcast) Reportable() bool              { return false }
func (ApsRxBcast) SceneIndex() int               { return -1 }
func (a *ApsRxBcast) Value() *ApsRxBcast         { return a }
func (a ApsRxBcast) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ApsRxBcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApsRxBcast(*nt)
	return br, err
}

func (a ApsRxBcast) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// ApsTxBcast is an autogenerated attribute in the Diagnostics cluster
type ApsTxBcast zcl.Zu16

const ApsTxBcastAttr zcl.AttrID = 263

func (ApsTxBcast) ID() zcl.AttrID                { return ApsTxBcastAttr }
func (ApsTxBcast) Cluster() zcl.ClusterID        { return DiagnosticsID }
func (ApsTxBcast) Name() string                  { return "APS Tx Bcast" }
func (ApsTxBcast) Readable() bool                { return true }
func (ApsTxBcast) Writable() bool                { return false }
func (ApsTxBcast) Reportable() bool              { return false }
func (ApsTxBcast) SceneIndex() int               { return -1 }
func (a *ApsTxBcast) Value() *ApsTxBcast         { return a }
func (a ApsTxBcast) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ApsTxBcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApsTxBcast(*nt)
	return br, err
}

func (a ApsTxBcast) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// ApsRxUcast is an autogenerated attribute in the Diagnostics cluster
type ApsRxUcast zcl.Zu16

const ApsRxUcastAttr zcl.AttrID = 264

func (ApsRxUcast) ID() zcl.AttrID                { return ApsRxUcastAttr }
func (ApsRxUcast) Cluster() zcl.ClusterID        { return DiagnosticsID }
func (ApsRxUcast) Name() string                  { return "APS Rx Ucast" }
func (ApsRxUcast) Readable() bool                { return true }
func (ApsRxUcast) Writable() bool                { return false }
func (ApsRxUcast) Reportable() bool              { return false }
func (ApsRxUcast) SceneIndex() int               { return -1 }
func (a *ApsRxUcast) Value() *ApsRxUcast         { return a }
func (a ApsRxUcast) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ApsRxUcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApsRxUcast(*nt)
	return br, err
}

func (a ApsRxUcast) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// ApsTxUcastSuccess is an autogenerated attribute in the Diagnostics cluster
type ApsTxUcastSuccess zcl.Zu16

const ApsTxUcastSuccessAttr zcl.AttrID = 265

func (ApsTxUcastSuccess) ID() zcl.AttrID                { return ApsTxUcastSuccessAttr }
func (ApsTxUcastSuccess) Cluster() zcl.ClusterID        { return DiagnosticsID }
func (ApsTxUcastSuccess) Name() string                  { return "APS Tx Ucast Success" }
func (ApsTxUcastSuccess) Readable() bool                { return true }
func (ApsTxUcastSuccess) Writable() bool                { return false }
func (ApsTxUcastSuccess) Reportable() bool              { return false }
func (ApsTxUcastSuccess) SceneIndex() int               { return -1 }
func (a *ApsTxUcastSuccess) Value() *ApsTxUcastSuccess  { return a }
func (a ApsTxUcastSuccess) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ApsTxUcastSuccess) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApsTxUcastSuccess(*nt)
	return br, err
}

func (a ApsTxUcastSuccess) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// ApsTxUcastRetry is an autogenerated attribute in the Diagnostics cluster
type ApsTxUcastRetry zcl.Zu16

const ApsTxUcastRetryAttr zcl.AttrID = 266

func (ApsTxUcastRetry) ID() zcl.AttrID                { return ApsTxUcastRetryAttr }
func (ApsTxUcastRetry) Cluster() zcl.ClusterID        { return DiagnosticsID }
func (ApsTxUcastRetry) Name() string                  { return "APS Tx Ucast Retry" }
func (ApsTxUcastRetry) Readable() bool                { return true }
func (ApsTxUcastRetry) Writable() bool                { return false }
func (ApsTxUcastRetry) Reportable() bool              { return false }
func (ApsTxUcastRetry) SceneIndex() int               { return -1 }
func (a *ApsTxUcastRetry) Value() *ApsTxUcastRetry    { return a }
func (a ApsTxUcastRetry) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ApsTxUcastRetry) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApsTxUcastRetry(*nt)
	return br, err
}

func (a ApsTxUcastRetry) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// ApsTxUcastFail is an autogenerated attribute in the Diagnostics cluster
type ApsTxUcastFail zcl.Zu16

const ApsTxUcastFailAttr zcl.AttrID = 267

func (ApsTxUcastFail) ID() zcl.AttrID                { return ApsTxUcastFailAttr }
func (ApsTxUcastFail) Cluster() zcl.ClusterID        { return DiagnosticsID }
func (ApsTxUcastFail) Name() string                  { return "APS Tx Ucast Fail" }
func (ApsTxUcastFail) Readable() bool                { return true }
func (ApsTxUcastFail) Writable() bool                { return false }
func (ApsTxUcastFail) Reportable() bool              { return false }
func (ApsTxUcastFail) SceneIndex() int               { return -1 }
func (a *ApsTxUcastFail) Value() *ApsTxUcastFail     { return a }
func (a ApsTxUcastFail) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ApsTxUcastFail) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApsTxUcastFail(*nt)
	return br, err
}

func (a ApsTxUcastFail) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// RouteDiscInitiated is an autogenerated attribute in the Diagnostics cluster
type RouteDiscInitiated zcl.Zu16

const RouteDiscInitiatedAttr zcl.AttrID = 268

func (RouteDiscInitiated) ID() zcl.AttrID                { return RouteDiscInitiatedAttr }
func (RouteDiscInitiated) Cluster() zcl.ClusterID        { return DiagnosticsID }
func (RouteDiscInitiated) Name() string                  { return "Route Disc Initiated" }
func (RouteDiscInitiated) Readable() bool                { return true }
func (RouteDiscInitiated) Writable() bool                { return false }
func (RouteDiscInitiated) Reportable() bool              { return false }
func (RouteDiscInitiated) SceneIndex() int               { return -1 }
func (a *RouteDiscInitiated) Value() *RouteDiscInitiated { return a }
func (a RouteDiscInitiated) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *RouteDiscInitiated) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RouteDiscInitiated(*nt)
	return br, err
}

func (a RouteDiscInitiated) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// NeighborAdded is an autogenerated attribute in the Diagnostics cluster
type NeighborAdded zcl.Zu16

const NeighborAddedAttr zcl.AttrID = 269

func (NeighborAdded) ID() zcl.AttrID                { return NeighborAddedAttr }
func (NeighborAdded) Cluster() zcl.ClusterID        { return DiagnosticsID }
func (NeighborAdded) Name() string                  { return "Neighbor Added" }
func (NeighborAdded) Readable() bool                { return true }
func (NeighborAdded) Writable() bool                { return false }
func (NeighborAdded) Reportable() bool              { return false }
func (NeighborAdded) SceneIndex() int               { return -1 }
func (a *NeighborAdded) Value() *NeighborAdded      { return a }
func (a NeighborAdded) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *NeighborAdded) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = NeighborAdded(*nt)
	return br, err
}

func (a NeighborAdded) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// NeighborRemoved is an autogenerated attribute in the Diagnostics cluster
type NeighborRemoved zcl.Zu16

const NeighborRemovedAttr zcl.AttrID = 270

func (NeighborRemoved) ID() zcl.AttrID                { return NeighborRemovedAttr }
func (NeighborRemoved) Cluster() zcl.ClusterID        { return DiagnosticsID }
func (NeighborRemoved) Name() string                  { return "Neighbor Removed" }
func (NeighborRemoved) Readable() bool                { return true }
func (NeighborRemoved) Writable() bool                { return false }
func (NeighborRemoved) Reportable() bool              { return false }
func (NeighborRemoved) SceneIndex() int               { return -1 }
func (a *NeighborRemoved) Value() *NeighborRemoved    { return a }
func (a NeighborRemoved) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *NeighborRemoved) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = NeighborRemoved(*nt)
	return br, err
}

func (a NeighborRemoved) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// NeighborStale is an autogenerated attribute in the Diagnostics cluster
type NeighborStale zcl.Zu16

const NeighborStaleAttr zcl.AttrID = 271

func (NeighborStale) ID() zcl.AttrID                { return NeighborStaleAttr }
func (NeighborStale) Cluster() zcl.ClusterID        { return DiagnosticsID }
func (NeighborStale) Name() string                  { return "Neighbor Stale" }
func (NeighborStale) Readable() bool                { return true }
func (NeighborStale) Writable() bool                { return false }
func (NeighborStale) Reportable() bool              { return false }
func (NeighborStale) SceneIndex() int               { return -1 }
func (a *NeighborStale) Value() *NeighborStale      { return a }
func (a NeighborStale) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *NeighborStale) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = NeighborStale(*nt)
	return br, err
}

func (a NeighborStale) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// JoinIndication is an autogenerated attribute in the Diagnostics cluster
type JoinIndication zcl.Zu16

const JoinIndicationAttr zcl.AttrID = 272

func (JoinIndication) ID() zcl.AttrID                { return JoinIndicationAttr }
func (JoinIndication) Cluster() zcl.ClusterID        { return DiagnosticsID }
func (JoinIndication) Name() string                  { return "Join Indication" }
func (JoinIndication) Readable() bool                { return true }
func (JoinIndication) Writable() bool                { return false }
func (JoinIndication) Reportable() bool              { return false }
func (JoinIndication) SceneIndex() int               { return -1 }
func (a *JoinIndication) Value() *JoinIndication     { return a }
func (a JoinIndication) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *JoinIndication) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = JoinIndication(*nt)
	return br, err
}

func (a JoinIndication) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// ChildMoved is an autogenerated attribute in the Diagnostics cluster
type ChildMoved zcl.Zu16

const ChildMovedAttr zcl.AttrID = 273

func (ChildMoved) ID() zcl.AttrID                { return ChildMovedAttr }
func (ChildMoved) Cluster() zcl.ClusterID        { return DiagnosticsID }
func (ChildMoved) Name() string                  { return "Child Moved" }
func (ChildMoved) Readable() bool                { return true }
func (ChildMoved) Writable() bool                { return false }
func (ChildMoved) Reportable() bool              { return false }
func (ChildMoved) SceneIndex() int               { return -1 }
func (a *ChildMoved) Value() *ChildMoved         { return a }
func (a ChildMoved) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ChildMoved) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ChildMoved(*nt)
	return br, err
}

func (a ChildMoved) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// NwkFcFailure is an autogenerated attribute in the Diagnostics cluster
type NwkFcFailure zcl.Zu16

const NwkFcFailureAttr zcl.AttrID = 274

func (NwkFcFailure) ID() zcl.AttrID                { return NwkFcFailureAttr }
func (NwkFcFailure) Cluster() zcl.ClusterID        { return DiagnosticsID }
func (NwkFcFailure) Name() string                  { return "NWK FC Failure" }
func (NwkFcFailure) Readable() bool                { return true }
func (NwkFcFailure) Writable() bool                { return false }
func (NwkFcFailure) Reportable() bool              { return false }
func (NwkFcFailure) SceneIndex() int               { return -1 }
func (a *NwkFcFailure) Value() *NwkFcFailure       { return a }
func (a NwkFcFailure) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *NwkFcFailure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = NwkFcFailure(*nt)
	return br, err
}

func (a NwkFcFailure) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// ApsFcFailure is an autogenerated attribute in the Diagnostics cluster
type ApsFcFailure zcl.Zu16

const ApsFcFailureAttr zcl.AttrID = 275

func (ApsFcFailure) ID() zcl.AttrID                { return ApsFcFailureAttr }
func (ApsFcFailure) Cluster() zcl.ClusterID        { return DiagnosticsID }
func (ApsFcFailure) Name() string                  { return "APS FC Failure" }
func (ApsFcFailure) Readable() bool                { return true }
func (ApsFcFailure) Writable() bool                { return false }
func (ApsFcFailure) Reportable() bool              { return false }
func (ApsFcFailure) SceneIndex() int               { return -1 }
func (a *ApsFcFailure) Value() *ApsFcFailure       { return a }
func (a ApsFcFailure) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ApsFcFailure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApsFcFailure(*nt)
	return br, err
}

func (a ApsFcFailure) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// ApsUnauthorizedKey is an autogenerated attribute in the Diagnostics cluster
type ApsUnauthorizedKey zcl.Zu16

const ApsUnauthorizedKeyAttr zcl.AttrID = 276

func (ApsUnauthorizedKey) ID() zcl.AttrID                { return ApsUnauthorizedKeyAttr }
func (ApsUnauthorizedKey) Cluster() zcl.ClusterID        { return DiagnosticsID }
func (ApsUnauthorizedKey) Name() string                  { return "APS Unauthorized Key" }
func (ApsUnauthorizedKey) Readable() bool                { return true }
func (ApsUnauthorizedKey) Writable() bool                { return false }
func (ApsUnauthorizedKey) Reportable() bool              { return false }
func (ApsUnauthorizedKey) SceneIndex() int               { return -1 }
func (a *ApsUnauthorizedKey) Value() *ApsUnauthorizedKey { return a }
func (a ApsUnauthorizedKey) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ApsUnauthorizedKey) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApsUnauthorizedKey(*nt)
	return br, err
}

func (a ApsUnauthorizedKey) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// NwkDecryptFailures is an autogenerated attribute in the Diagnostics cluster
type NwkDecryptFailures zcl.Zu16

const NwkDecryptFailuresAttr zcl.AttrID = 277

func (NwkDecryptFailures) ID() zcl.AttrID                { return NwkDecryptFailuresAttr }
func (NwkDecryptFailures) Cluster() zcl.ClusterID        { return DiagnosticsID }
func (NwkDecryptFailures) Name() string                  { return "NWK Decrypt Failures" }
func (NwkDecryptFailures) Readable() bool                { return true }
func (NwkDecryptFailures) Writable() bool                { return false }
func (NwkDecryptFailures) Reportable() bool              { return false }
func (NwkDecryptFailures) SceneIndex() int               { return -1 }
func (a *NwkDecryptFailures) Value() *NwkDecryptFailures { return a }
func (a NwkDecryptFailures) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *NwkDecryptFailures) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = NwkDecryptFailures(*nt)
	return br, err
}

func (a NwkDecryptFailures) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// ApsDecryptFailures is an autogenerated attribute in the Diagnostics cluster
type ApsDecryptFailures zcl.Zu16

const ApsDecryptFailuresAttr zcl.AttrID = 278

func (ApsDecryptFailures) ID() zcl.AttrID                { return ApsDecryptFailuresAttr }
func (ApsDecryptFailures) Cluster() zcl.ClusterID        { return DiagnosticsID }
func (ApsDecryptFailures) Name() string                  { return "APS Decrypt Failures" }
func (ApsDecryptFailures) Readable() bool                { return true }
func (ApsDecryptFailures) Writable() bool                { return false }
func (ApsDecryptFailures) Reportable() bool              { return false }
func (ApsDecryptFailures) SceneIndex() int               { return -1 }
func (a *ApsDecryptFailures) Value() *ApsDecryptFailures { return a }
func (a ApsDecryptFailures) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ApsDecryptFailures) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApsDecryptFailures(*nt)
	return br, err
}

func (a ApsDecryptFailures) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// PacketBufferAllocFailures is an autogenerated attribute in the Diagnostics cluster
type PacketBufferAllocFailures zcl.Zu16

const PacketBufferAllocFailuresAttr zcl.AttrID = 279

func (PacketBufferAllocFailures) ID() zcl.AttrID                       { return PacketBufferAllocFailuresAttr }
func (PacketBufferAllocFailures) Cluster() zcl.ClusterID               { return DiagnosticsID }
func (PacketBufferAllocFailures) Name() string                         { return "Packet Buffer Alloc Failures" }
func (PacketBufferAllocFailures) Readable() bool                       { return true }
func (PacketBufferAllocFailures) Writable() bool                       { return false }
func (PacketBufferAllocFailures) Reportable() bool                     { return false }
func (PacketBufferAllocFailures) SceneIndex() int                      { return -1 }
func (a *PacketBufferAllocFailures) Value() *PacketBufferAllocFailures { return a }
func (a PacketBufferAllocFailures) MarshalZcl() ([]byte, error)        { return zcl.Zu16(a).MarshalZcl() }

func (a *PacketBufferAllocFailures) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = PacketBufferAllocFailures(*nt)
	return br, err
}

func (a PacketBufferAllocFailures) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// RelayedUcast is an autogenerated attribute in the Diagnostics cluster
type RelayedUcast zcl.Zu16

const RelayedUcastAttr zcl.AttrID = 280

func (RelayedUcast) ID() zcl.AttrID                { return RelayedUcastAttr }
func (RelayedUcast) Cluster() zcl.ClusterID        { return DiagnosticsID }
func (RelayedUcast) Name() string                  { return "Relayed Ucast" }
func (RelayedUcast) Readable() bool                { return true }
func (RelayedUcast) Writable() bool                { return false }
func (RelayedUcast) Reportable() bool              { return false }
func (RelayedUcast) SceneIndex() int               { return -1 }
func (a *RelayedUcast) Value() *RelayedUcast       { return a }
func (a RelayedUcast) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *RelayedUcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RelayedUcast(*nt)
	return br, err
}

func (a RelayedUcast) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// PhyToMacQueueLimitReached is an autogenerated attribute in the Diagnostics cluster
type PhyToMacQueueLimitReached zcl.Zu16

const PhyToMacQueueLimitReachedAttr zcl.AttrID = 281

func (PhyToMacQueueLimitReached) ID() zcl.AttrID                       { return PhyToMacQueueLimitReachedAttr }
func (PhyToMacQueueLimitReached) Cluster() zcl.ClusterID               { return DiagnosticsID }
func (PhyToMacQueueLimitReached) Name() string                         { return "Phy to MAC queue limit reached" }
func (PhyToMacQueueLimitReached) Readable() bool                       { return true }
func (PhyToMacQueueLimitReached) Writable() bool                       { return false }
func (PhyToMacQueueLimitReached) Reportable() bool                     { return false }
func (PhyToMacQueueLimitReached) SceneIndex() int                      { return -1 }
func (a *PhyToMacQueueLimitReached) Value() *PhyToMacQueueLimitReached { return a }
func (a PhyToMacQueueLimitReached) MarshalZcl() ([]byte, error)        { return zcl.Zu16(a).MarshalZcl() }

func (a *PhyToMacQueueLimitReached) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = PhyToMacQueueLimitReached(*nt)
	return br, err
}

func (a PhyToMacQueueLimitReached) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// PacketValidateDropcount is an autogenerated attribute in the Diagnostics cluster
type PacketValidateDropcount zcl.Zu16

const PacketValidateDropcountAttr zcl.AttrID = 282

func (PacketValidateDropcount) ID() zcl.AttrID                     { return PacketValidateDropcountAttr }
func (PacketValidateDropcount) Cluster() zcl.ClusterID             { return DiagnosticsID }
func (PacketValidateDropcount) Name() string                       { return "Packet Validate Dropcount" }
func (PacketValidateDropcount) Readable() bool                     { return true }
func (PacketValidateDropcount) Writable() bool                     { return false }
func (PacketValidateDropcount) Reportable() bool                   { return false }
func (PacketValidateDropcount) SceneIndex() int                    { return -1 }
func (a *PacketValidateDropcount) Value() *PacketValidateDropcount { return a }
func (a PacketValidateDropcount) MarshalZcl() ([]byte, error)      { return zcl.Zu16(a).MarshalZcl() }

func (a *PacketValidateDropcount) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = PacketValidateDropcount(*nt)
	return br, err
}

func (a PacketValidateDropcount) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// AvgMacRetryPerApsMsgSent is an autogenerated attribute in the Diagnostics cluster
type AvgMacRetryPerApsMsgSent zcl.Zu16

const AvgMacRetryPerApsMsgSentAttr zcl.AttrID = 283

func (AvgMacRetryPerApsMsgSent) ID() zcl.AttrID                      { return AvgMacRetryPerApsMsgSentAttr }
func (AvgMacRetryPerApsMsgSent) Cluster() zcl.ClusterID              { return DiagnosticsID }
func (AvgMacRetryPerApsMsgSent) Name() string                        { return "Avg MAC Retry per APS Msg Sent" }
func (AvgMacRetryPerApsMsgSent) Readable() bool                      { return true }
func (AvgMacRetryPerApsMsgSent) Writable() bool                      { return false }
func (AvgMacRetryPerApsMsgSent) Reportable() bool                    { return false }
func (AvgMacRetryPerApsMsgSent) SceneIndex() int                     { return -1 }
func (a *AvgMacRetryPerApsMsgSent) Value() *AvgMacRetryPerApsMsgSent { return a }
func (a AvgMacRetryPerApsMsgSent) MarshalZcl() ([]byte, error)       { return zcl.Zu16(a).MarshalZcl() }

func (a *AvgMacRetryPerApsMsgSent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AvgMacRetryPerApsMsgSent(*nt)
	return br, err
}

func (a AvgMacRetryPerApsMsgSent) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// LastMessageLqi is an autogenerated attribute in the Diagnostics cluster
type LastMessageLqi zcl.Zu8

const LastMessageLqiAttr zcl.AttrID = 284

func (LastMessageLqi) ID() zcl.AttrID                { return LastMessageLqiAttr }
func (LastMessageLqi) Cluster() zcl.ClusterID        { return DiagnosticsID }
func (LastMessageLqi) Name() string                  { return "Last Message LQI" }
func (LastMessageLqi) Readable() bool                { return true }
func (LastMessageLqi) Writable() bool                { return false }
func (LastMessageLqi) Reportable() bool              { return false }
func (LastMessageLqi) SceneIndex() int               { return -1 }
func (a *LastMessageLqi) Value() *LastMessageLqi     { return a }
func (a LastMessageLqi) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *LastMessageLqi) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = LastMessageLqi(*nt)
	return br, err
}

func (a LastMessageLqi) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// LastMessageRssi is an autogenerated attribute in the Diagnostics cluster
type LastMessageRssi zcl.Zs8

const LastMessageRssiAttr zcl.AttrID = 285

func (LastMessageRssi) ID() zcl.AttrID                { return LastMessageRssiAttr }
func (LastMessageRssi) Cluster() zcl.ClusterID        { return DiagnosticsID }
func (LastMessageRssi) Name() string                  { return "Last Message RSSI" }
func (LastMessageRssi) Readable() bool                { return true }
func (LastMessageRssi) Writable() bool                { return false }
func (LastMessageRssi) Reportable() bool              { return false }
func (LastMessageRssi) SceneIndex() int               { return -1 }
func (a *LastMessageRssi) Value() *LastMessageRssi    { return a }
func (a LastMessageRssi) MarshalZcl() ([]byte, error) { return zcl.Zs8(a).MarshalZcl() }

func (a *LastMessageRssi) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*a = LastMessageRssi(*nt)
	return br, err
}

func (a LastMessageRssi) String() string {
	return zcl.Sprintf("%v", zcl.Zs8(a))
}
