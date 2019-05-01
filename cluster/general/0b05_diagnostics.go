// The diagnostics cluster provides access to information regarding the operation of the ZigBee stack over time.
package general

import (
	"neotor.se/zcl"
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

const NumberOfResetsAttr zcl.AttrID = 0

type NumberOfResets zcl.Zu16

func (a NumberOfResets) ID() zcl.AttrID          { return NumberOfResetsAttr }
func (a NumberOfResets) Cluster() zcl.ClusterID  { return DiagnosticsID }
func (a *NumberOfResets) Value() *NumberOfResets { return a }
func (a NumberOfResets) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *NumberOfResets) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = NumberOfResets(*nt)
	return br, err
}

func (a NumberOfResets) Readable() bool   { return true }
func (a NumberOfResets) Writable() bool   { return false }
func (a NumberOfResets) Reportable() bool { return false }
func (a NumberOfResets) SceneIndex() int  { return -1 }

func (a NumberOfResets) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const PersistensMemoryWritesAttr zcl.AttrID = 1

type PersistensMemoryWrites zcl.Zu16

func (a PersistensMemoryWrites) ID() zcl.AttrID                  { return PersistensMemoryWritesAttr }
func (a PersistensMemoryWrites) Cluster() zcl.ClusterID          { return DiagnosticsID }
func (a *PersistensMemoryWrites) Value() *PersistensMemoryWrites { return a }
func (a PersistensMemoryWrites) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *PersistensMemoryWrites) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = PersistensMemoryWrites(*nt)
	return br, err
}

func (a PersistensMemoryWrites) Readable() bool   { return true }
func (a PersistensMemoryWrites) Writable() bool   { return false }
func (a PersistensMemoryWrites) Reportable() bool { return false }
func (a PersistensMemoryWrites) SceneIndex() int  { return -1 }

func (a PersistensMemoryWrites) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const MacRxBcastAttr zcl.AttrID = 256

type MacRxBcast zcl.Zu32

func (a MacRxBcast) ID() zcl.AttrID         { return MacRxBcastAttr }
func (a MacRxBcast) Cluster() zcl.ClusterID { return DiagnosticsID }
func (a *MacRxBcast) Value() *MacRxBcast    { return a }
func (a MacRxBcast) MarshalZcl() ([]byte, error) {
	return zcl.Zu32(a).MarshalZcl()
}
func (a *MacRxBcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = MacRxBcast(*nt)
	return br, err
}

func (a MacRxBcast) Readable() bool   { return true }
func (a MacRxBcast) Writable() bool   { return false }
func (a MacRxBcast) Reportable() bool { return false }
func (a MacRxBcast) SceneIndex() int  { return -1 }

func (a MacRxBcast) String() string {

	return zcl.Sprintf("%s", zcl.Zu32(a))
}

const MacTxBcastAttr zcl.AttrID = 257

type MacTxBcast zcl.Zu32

func (a MacTxBcast) ID() zcl.AttrID         { return MacTxBcastAttr }
func (a MacTxBcast) Cluster() zcl.ClusterID { return DiagnosticsID }
func (a *MacTxBcast) Value() *MacTxBcast    { return a }
func (a MacTxBcast) MarshalZcl() ([]byte, error) {
	return zcl.Zu32(a).MarshalZcl()
}
func (a *MacTxBcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = MacTxBcast(*nt)
	return br, err
}

func (a MacTxBcast) Readable() bool   { return true }
func (a MacTxBcast) Writable() bool   { return false }
func (a MacTxBcast) Reportable() bool { return false }
func (a MacTxBcast) SceneIndex() int  { return -1 }

func (a MacTxBcast) String() string {

	return zcl.Sprintf("%s", zcl.Zu32(a))
}

const MacRxUcastAttr zcl.AttrID = 258

type MacRxUcast zcl.Zu32

func (a MacRxUcast) ID() zcl.AttrID         { return MacRxUcastAttr }
func (a MacRxUcast) Cluster() zcl.ClusterID { return DiagnosticsID }
func (a *MacRxUcast) Value() *MacRxUcast    { return a }
func (a MacRxUcast) MarshalZcl() ([]byte, error) {
	return zcl.Zu32(a).MarshalZcl()
}
func (a *MacRxUcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = MacRxUcast(*nt)
	return br, err
}

func (a MacRxUcast) Readable() bool   { return true }
func (a MacRxUcast) Writable() bool   { return false }
func (a MacRxUcast) Reportable() bool { return false }
func (a MacRxUcast) SceneIndex() int  { return -1 }

func (a MacRxUcast) String() string {

	return zcl.Sprintf("%s", zcl.Zu32(a))
}

const MacTxUcastAttr zcl.AttrID = 259

type MacTxUcast zcl.Zu32

func (a MacTxUcast) ID() zcl.AttrID         { return MacTxUcastAttr }
func (a MacTxUcast) Cluster() zcl.ClusterID { return DiagnosticsID }
func (a *MacTxUcast) Value() *MacTxUcast    { return a }
func (a MacTxUcast) MarshalZcl() ([]byte, error) {
	return zcl.Zu32(a).MarshalZcl()
}
func (a *MacTxUcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = MacTxUcast(*nt)
	return br, err
}

func (a MacTxUcast) Readable() bool   { return true }
func (a MacTxUcast) Writable() bool   { return false }
func (a MacTxUcast) Reportable() bool { return false }
func (a MacTxUcast) SceneIndex() int  { return -1 }

func (a MacTxUcast) String() string {

	return zcl.Sprintf("%s", zcl.Zu32(a))
}

const MacTxUcastRetryAttr zcl.AttrID = 260

type MacTxUcastRetry zcl.Zu16

func (a MacTxUcastRetry) ID() zcl.AttrID           { return MacTxUcastRetryAttr }
func (a MacTxUcastRetry) Cluster() zcl.ClusterID   { return DiagnosticsID }
func (a *MacTxUcastRetry) Value() *MacTxUcastRetry { return a }
func (a MacTxUcastRetry) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *MacTxUcastRetry) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MacTxUcastRetry(*nt)
	return br, err
}

func (a MacTxUcastRetry) Readable() bool   { return true }
func (a MacTxUcastRetry) Writable() bool   { return false }
func (a MacTxUcastRetry) Reportable() bool { return false }
func (a MacTxUcastRetry) SceneIndex() int  { return -1 }

func (a MacTxUcastRetry) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const MacTxUcastFailAttr zcl.AttrID = 261

type MacTxUcastFail zcl.Zu16

func (a MacTxUcastFail) ID() zcl.AttrID          { return MacTxUcastFailAttr }
func (a MacTxUcastFail) Cluster() zcl.ClusterID  { return DiagnosticsID }
func (a *MacTxUcastFail) Value() *MacTxUcastFail { return a }
func (a MacTxUcastFail) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *MacTxUcastFail) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MacTxUcastFail(*nt)
	return br, err
}

func (a MacTxUcastFail) Readable() bool   { return true }
func (a MacTxUcastFail) Writable() bool   { return false }
func (a MacTxUcastFail) Reportable() bool { return false }
func (a MacTxUcastFail) SceneIndex() int  { return -1 }

func (a MacTxUcastFail) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ApsRxBcastAttr zcl.AttrID = 262

type ApsRxBcast zcl.Zu16

func (a ApsRxBcast) ID() zcl.AttrID         { return ApsRxBcastAttr }
func (a ApsRxBcast) Cluster() zcl.ClusterID { return DiagnosticsID }
func (a *ApsRxBcast) Value() *ApsRxBcast    { return a }
func (a ApsRxBcast) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ApsRxBcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApsRxBcast(*nt)
	return br, err
}

func (a ApsRxBcast) Readable() bool   { return true }
func (a ApsRxBcast) Writable() bool   { return false }
func (a ApsRxBcast) Reportable() bool { return false }
func (a ApsRxBcast) SceneIndex() int  { return -1 }

func (a ApsRxBcast) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ApsTxBcastAttr zcl.AttrID = 263

type ApsTxBcast zcl.Zu16

func (a ApsTxBcast) ID() zcl.AttrID         { return ApsTxBcastAttr }
func (a ApsTxBcast) Cluster() zcl.ClusterID { return DiagnosticsID }
func (a *ApsTxBcast) Value() *ApsTxBcast    { return a }
func (a ApsTxBcast) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ApsTxBcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApsTxBcast(*nt)
	return br, err
}

func (a ApsTxBcast) Readable() bool   { return true }
func (a ApsTxBcast) Writable() bool   { return false }
func (a ApsTxBcast) Reportable() bool { return false }
func (a ApsTxBcast) SceneIndex() int  { return -1 }

func (a ApsTxBcast) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ApsRxUcastAttr zcl.AttrID = 264

type ApsRxUcast zcl.Zu16

func (a ApsRxUcast) ID() zcl.AttrID         { return ApsRxUcastAttr }
func (a ApsRxUcast) Cluster() zcl.ClusterID { return DiagnosticsID }
func (a *ApsRxUcast) Value() *ApsRxUcast    { return a }
func (a ApsRxUcast) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ApsRxUcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApsRxUcast(*nt)
	return br, err
}

func (a ApsRxUcast) Readable() bool   { return true }
func (a ApsRxUcast) Writable() bool   { return false }
func (a ApsRxUcast) Reportable() bool { return false }
func (a ApsRxUcast) SceneIndex() int  { return -1 }

func (a ApsRxUcast) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ApsTxUcastSuccessAttr zcl.AttrID = 265

type ApsTxUcastSuccess zcl.Zu16

func (a ApsTxUcastSuccess) ID() zcl.AttrID             { return ApsTxUcastSuccessAttr }
func (a ApsTxUcastSuccess) Cluster() zcl.ClusterID     { return DiagnosticsID }
func (a *ApsTxUcastSuccess) Value() *ApsTxUcastSuccess { return a }
func (a ApsTxUcastSuccess) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ApsTxUcastSuccess) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApsTxUcastSuccess(*nt)
	return br, err
}

func (a ApsTxUcastSuccess) Readable() bool   { return true }
func (a ApsTxUcastSuccess) Writable() bool   { return false }
func (a ApsTxUcastSuccess) Reportable() bool { return false }
func (a ApsTxUcastSuccess) SceneIndex() int  { return -1 }

func (a ApsTxUcastSuccess) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ApsTxUcastRetryAttr zcl.AttrID = 266

type ApsTxUcastRetry zcl.Zu16

func (a ApsTxUcastRetry) ID() zcl.AttrID           { return ApsTxUcastRetryAttr }
func (a ApsTxUcastRetry) Cluster() zcl.ClusterID   { return DiagnosticsID }
func (a *ApsTxUcastRetry) Value() *ApsTxUcastRetry { return a }
func (a ApsTxUcastRetry) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ApsTxUcastRetry) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApsTxUcastRetry(*nt)
	return br, err
}

func (a ApsTxUcastRetry) Readable() bool   { return true }
func (a ApsTxUcastRetry) Writable() bool   { return false }
func (a ApsTxUcastRetry) Reportable() bool { return false }
func (a ApsTxUcastRetry) SceneIndex() int  { return -1 }

func (a ApsTxUcastRetry) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ApsTxUcastFailAttr zcl.AttrID = 267

type ApsTxUcastFail zcl.Zu16

func (a ApsTxUcastFail) ID() zcl.AttrID          { return ApsTxUcastFailAttr }
func (a ApsTxUcastFail) Cluster() zcl.ClusterID  { return DiagnosticsID }
func (a *ApsTxUcastFail) Value() *ApsTxUcastFail { return a }
func (a ApsTxUcastFail) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ApsTxUcastFail) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApsTxUcastFail(*nt)
	return br, err
}

func (a ApsTxUcastFail) Readable() bool   { return true }
func (a ApsTxUcastFail) Writable() bool   { return false }
func (a ApsTxUcastFail) Reportable() bool { return false }
func (a ApsTxUcastFail) SceneIndex() int  { return -1 }

func (a ApsTxUcastFail) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const RouteDiscInitiatedAttr zcl.AttrID = 268

type RouteDiscInitiated zcl.Zu16

func (a RouteDiscInitiated) ID() zcl.AttrID              { return RouteDiscInitiatedAttr }
func (a RouteDiscInitiated) Cluster() zcl.ClusterID      { return DiagnosticsID }
func (a *RouteDiscInitiated) Value() *RouteDiscInitiated { return a }
func (a RouteDiscInitiated) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RouteDiscInitiated) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RouteDiscInitiated(*nt)
	return br, err
}

func (a RouteDiscInitiated) Readable() bool   { return true }
func (a RouteDiscInitiated) Writable() bool   { return false }
func (a RouteDiscInitiated) Reportable() bool { return false }
func (a RouteDiscInitiated) SceneIndex() int  { return -1 }

func (a RouteDiscInitiated) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const NeighborAddedAttr zcl.AttrID = 269

type NeighborAdded zcl.Zu16

func (a NeighborAdded) ID() zcl.AttrID         { return NeighborAddedAttr }
func (a NeighborAdded) Cluster() zcl.ClusterID { return DiagnosticsID }
func (a *NeighborAdded) Value() *NeighborAdded { return a }
func (a NeighborAdded) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *NeighborAdded) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = NeighborAdded(*nt)
	return br, err
}

func (a NeighborAdded) Readable() bool   { return true }
func (a NeighborAdded) Writable() bool   { return false }
func (a NeighborAdded) Reportable() bool { return false }
func (a NeighborAdded) SceneIndex() int  { return -1 }

func (a NeighborAdded) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const NeighborRemovedAttr zcl.AttrID = 270

type NeighborRemoved zcl.Zu16

func (a NeighborRemoved) ID() zcl.AttrID           { return NeighborRemovedAttr }
func (a NeighborRemoved) Cluster() zcl.ClusterID   { return DiagnosticsID }
func (a *NeighborRemoved) Value() *NeighborRemoved { return a }
func (a NeighborRemoved) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *NeighborRemoved) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = NeighborRemoved(*nt)
	return br, err
}

func (a NeighborRemoved) Readable() bool   { return true }
func (a NeighborRemoved) Writable() bool   { return false }
func (a NeighborRemoved) Reportable() bool { return false }
func (a NeighborRemoved) SceneIndex() int  { return -1 }

func (a NeighborRemoved) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const NeighborStaleAttr zcl.AttrID = 271

type NeighborStale zcl.Zu16

func (a NeighborStale) ID() zcl.AttrID         { return NeighborStaleAttr }
func (a NeighborStale) Cluster() zcl.ClusterID { return DiagnosticsID }
func (a *NeighborStale) Value() *NeighborStale { return a }
func (a NeighborStale) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *NeighborStale) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = NeighborStale(*nt)
	return br, err
}

func (a NeighborStale) Readable() bool   { return true }
func (a NeighborStale) Writable() bool   { return false }
func (a NeighborStale) Reportable() bool { return false }
func (a NeighborStale) SceneIndex() int  { return -1 }

func (a NeighborStale) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const JoinIndicationAttr zcl.AttrID = 272

type JoinIndication zcl.Zu16

func (a JoinIndication) ID() zcl.AttrID          { return JoinIndicationAttr }
func (a JoinIndication) Cluster() zcl.ClusterID  { return DiagnosticsID }
func (a *JoinIndication) Value() *JoinIndication { return a }
func (a JoinIndication) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *JoinIndication) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = JoinIndication(*nt)
	return br, err
}

func (a JoinIndication) Readable() bool   { return true }
func (a JoinIndication) Writable() bool   { return false }
func (a JoinIndication) Reportable() bool { return false }
func (a JoinIndication) SceneIndex() int  { return -1 }

func (a JoinIndication) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ChildMovedAttr zcl.AttrID = 273

type ChildMoved zcl.Zu16

func (a ChildMoved) ID() zcl.AttrID         { return ChildMovedAttr }
func (a ChildMoved) Cluster() zcl.ClusterID { return DiagnosticsID }
func (a *ChildMoved) Value() *ChildMoved    { return a }
func (a ChildMoved) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ChildMoved) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ChildMoved(*nt)
	return br, err
}

func (a ChildMoved) Readable() bool   { return true }
func (a ChildMoved) Writable() bool   { return false }
func (a ChildMoved) Reportable() bool { return false }
func (a ChildMoved) SceneIndex() int  { return -1 }

func (a ChildMoved) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const NwkFcFailureAttr zcl.AttrID = 274

type NwkFcFailure zcl.Zu16

func (a NwkFcFailure) ID() zcl.AttrID         { return NwkFcFailureAttr }
func (a NwkFcFailure) Cluster() zcl.ClusterID { return DiagnosticsID }
func (a *NwkFcFailure) Value() *NwkFcFailure  { return a }
func (a NwkFcFailure) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *NwkFcFailure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = NwkFcFailure(*nt)
	return br, err
}

func (a NwkFcFailure) Readable() bool   { return true }
func (a NwkFcFailure) Writable() bool   { return false }
func (a NwkFcFailure) Reportable() bool { return false }
func (a NwkFcFailure) SceneIndex() int  { return -1 }

func (a NwkFcFailure) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ApsFcFailureAttr zcl.AttrID = 275

type ApsFcFailure zcl.Zu16

func (a ApsFcFailure) ID() zcl.AttrID         { return ApsFcFailureAttr }
func (a ApsFcFailure) Cluster() zcl.ClusterID { return DiagnosticsID }
func (a *ApsFcFailure) Value() *ApsFcFailure  { return a }
func (a ApsFcFailure) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ApsFcFailure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApsFcFailure(*nt)
	return br, err
}

func (a ApsFcFailure) Readable() bool   { return true }
func (a ApsFcFailure) Writable() bool   { return false }
func (a ApsFcFailure) Reportable() bool { return false }
func (a ApsFcFailure) SceneIndex() int  { return -1 }

func (a ApsFcFailure) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ApsUnauthorizedKeyAttr zcl.AttrID = 276

type ApsUnauthorizedKey zcl.Zu16

func (a ApsUnauthorizedKey) ID() zcl.AttrID              { return ApsUnauthorizedKeyAttr }
func (a ApsUnauthorizedKey) Cluster() zcl.ClusterID      { return DiagnosticsID }
func (a *ApsUnauthorizedKey) Value() *ApsUnauthorizedKey { return a }
func (a ApsUnauthorizedKey) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ApsUnauthorizedKey) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApsUnauthorizedKey(*nt)
	return br, err
}

func (a ApsUnauthorizedKey) Readable() bool   { return true }
func (a ApsUnauthorizedKey) Writable() bool   { return false }
func (a ApsUnauthorizedKey) Reportable() bool { return false }
func (a ApsUnauthorizedKey) SceneIndex() int  { return -1 }

func (a ApsUnauthorizedKey) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const NwkDecryptFailuresAttr zcl.AttrID = 277

type NwkDecryptFailures zcl.Zu16

func (a NwkDecryptFailures) ID() zcl.AttrID              { return NwkDecryptFailuresAttr }
func (a NwkDecryptFailures) Cluster() zcl.ClusterID      { return DiagnosticsID }
func (a *NwkDecryptFailures) Value() *NwkDecryptFailures { return a }
func (a NwkDecryptFailures) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *NwkDecryptFailures) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = NwkDecryptFailures(*nt)
	return br, err
}

func (a NwkDecryptFailures) Readable() bool   { return true }
func (a NwkDecryptFailures) Writable() bool   { return false }
func (a NwkDecryptFailures) Reportable() bool { return false }
func (a NwkDecryptFailures) SceneIndex() int  { return -1 }

func (a NwkDecryptFailures) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ApsDecryptFailuresAttr zcl.AttrID = 278

type ApsDecryptFailures zcl.Zu16

func (a ApsDecryptFailures) ID() zcl.AttrID              { return ApsDecryptFailuresAttr }
func (a ApsDecryptFailures) Cluster() zcl.ClusterID      { return DiagnosticsID }
func (a *ApsDecryptFailures) Value() *ApsDecryptFailures { return a }
func (a ApsDecryptFailures) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ApsDecryptFailures) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApsDecryptFailures(*nt)
	return br, err
}

func (a ApsDecryptFailures) Readable() bool   { return true }
func (a ApsDecryptFailures) Writable() bool   { return false }
func (a ApsDecryptFailures) Reportable() bool { return false }
func (a ApsDecryptFailures) SceneIndex() int  { return -1 }

func (a ApsDecryptFailures) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const PacketBufferAllocFailuresAttr zcl.AttrID = 279

type PacketBufferAllocFailures zcl.Zu16

func (a PacketBufferAllocFailures) ID() zcl.AttrID                     { return PacketBufferAllocFailuresAttr }
func (a PacketBufferAllocFailures) Cluster() zcl.ClusterID             { return DiagnosticsID }
func (a *PacketBufferAllocFailures) Value() *PacketBufferAllocFailures { return a }
func (a PacketBufferAllocFailures) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *PacketBufferAllocFailures) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = PacketBufferAllocFailures(*nt)
	return br, err
}

func (a PacketBufferAllocFailures) Readable() bool   { return true }
func (a PacketBufferAllocFailures) Writable() bool   { return false }
func (a PacketBufferAllocFailures) Reportable() bool { return false }
func (a PacketBufferAllocFailures) SceneIndex() int  { return -1 }

func (a PacketBufferAllocFailures) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const RelayedUcastAttr zcl.AttrID = 280

type RelayedUcast zcl.Zu16

func (a RelayedUcast) ID() zcl.AttrID         { return RelayedUcastAttr }
func (a RelayedUcast) Cluster() zcl.ClusterID { return DiagnosticsID }
func (a *RelayedUcast) Value() *RelayedUcast  { return a }
func (a RelayedUcast) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RelayedUcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RelayedUcast(*nt)
	return br, err
}

func (a RelayedUcast) Readable() bool   { return true }
func (a RelayedUcast) Writable() bool   { return false }
func (a RelayedUcast) Reportable() bool { return false }
func (a RelayedUcast) SceneIndex() int  { return -1 }

func (a RelayedUcast) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const PhyToMacQueueLimitReachedAttr zcl.AttrID = 281

type PhyToMacQueueLimitReached zcl.Zu16

func (a PhyToMacQueueLimitReached) ID() zcl.AttrID                     { return PhyToMacQueueLimitReachedAttr }
func (a PhyToMacQueueLimitReached) Cluster() zcl.ClusterID             { return DiagnosticsID }
func (a *PhyToMacQueueLimitReached) Value() *PhyToMacQueueLimitReached { return a }
func (a PhyToMacQueueLimitReached) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *PhyToMacQueueLimitReached) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = PhyToMacQueueLimitReached(*nt)
	return br, err
}

func (a PhyToMacQueueLimitReached) Readable() bool   { return true }
func (a PhyToMacQueueLimitReached) Writable() bool   { return false }
func (a PhyToMacQueueLimitReached) Reportable() bool { return false }
func (a PhyToMacQueueLimitReached) SceneIndex() int  { return -1 }

func (a PhyToMacQueueLimitReached) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const PacketValidateDropcountAttr zcl.AttrID = 282

type PacketValidateDropcount zcl.Zu16

func (a PacketValidateDropcount) ID() zcl.AttrID                   { return PacketValidateDropcountAttr }
func (a PacketValidateDropcount) Cluster() zcl.ClusterID           { return DiagnosticsID }
func (a *PacketValidateDropcount) Value() *PacketValidateDropcount { return a }
func (a PacketValidateDropcount) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *PacketValidateDropcount) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = PacketValidateDropcount(*nt)
	return br, err
}

func (a PacketValidateDropcount) Readable() bool   { return true }
func (a PacketValidateDropcount) Writable() bool   { return false }
func (a PacketValidateDropcount) Reportable() bool { return false }
func (a PacketValidateDropcount) SceneIndex() int  { return -1 }

func (a PacketValidateDropcount) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const AvgMacRetryPerApsMsgSentAttr zcl.AttrID = 283

type AvgMacRetryPerApsMsgSent zcl.Zu16

func (a AvgMacRetryPerApsMsgSent) ID() zcl.AttrID                    { return AvgMacRetryPerApsMsgSentAttr }
func (a AvgMacRetryPerApsMsgSent) Cluster() zcl.ClusterID            { return DiagnosticsID }
func (a *AvgMacRetryPerApsMsgSent) Value() *AvgMacRetryPerApsMsgSent { return a }
func (a AvgMacRetryPerApsMsgSent) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AvgMacRetryPerApsMsgSent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AvgMacRetryPerApsMsgSent(*nt)
	return br, err
}

func (a AvgMacRetryPerApsMsgSent) Readable() bool   { return true }
func (a AvgMacRetryPerApsMsgSent) Writable() bool   { return false }
func (a AvgMacRetryPerApsMsgSent) Reportable() bool { return false }
func (a AvgMacRetryPerApsMsgSent) SceneIndex() int  { return -1 }

func (a AvgMacRetryPerApsMsgSent) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const LastMessageLqiAttr zcl.AttrID = 284

type LastMessageLqi zcl.Zu8

func (a LastMessageLqi) ID() zcl.AttrID          { return LastMessageLqiAttr }
func (a LastMessageLqi) Cluster() zcl.ClusterID  { return DiagnosticsID }
func (a *LastMessageLqi) Value() *LastMessageLqi { return a }
func (a LastMessageLqi) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *LastMessageLqi) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = LastMessageLqi(*nt)
	return br, err
}

func (a LastMessageLqi) Readable() bool   { return true }
func (a LastMessageLqi) Writable() bool   { return false }
func (a LastMessageLqi) Reportable() bool { return false }
func (a LastMessageLqi) SceneIndex() int  { return -1 }

func (a LastMessageLqi) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

const LastMessageRssiAttr zcl.AttrID = 285

type LastMessageRssi zcl.Zs8

func (a LastMessageRssi) ID() zcl.AttrID           { return LastMessageRssiAttr }
func (a LastMessageRssi) Cluster() zcl.ClusterID   { return DiagnosticsID }
func (a *LastMessageRssi) Value() *LastMessageRssi { return a }
func (a LastMessageRssi) MarshalZcl() ([]byte, error) {
	return zcl.Zs8(a).MarshalZcl()
}
func (a *LastMessageRssi) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*a = LastMessageRssi(*nt)
	return br, err
}

func (a LastMessageRssi) Readable() bool   { return true }
func (a LastMessageRssi) Writable() bool   { return false }
func (a LastMessageRssi) Reportable() bool { return false }
func (a LastMessageRssi) SceneIndex() int  { return -1 }

func (a LastMessageRssi) String() string {

	return zcl.Sprintf("%s", zcl.Zs8(a))
}
