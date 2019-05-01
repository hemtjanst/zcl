// The diagnostics cluster provides access to information regarding the operation of the ZigBee stack over time.
package general

import (
	"neotor.se/zcl"
)

// Diagnostics
// The diagnostics cluster provides access to information regarding the operation of the ZigBee stack over time.

func NewDiagnosticsServer(profile zcl.ProfileID) *DiagnosticsServer {
	return &DiagnosticsServer{p: profile}
}
func NewDiagnosticsClient(profile zcl.ProfileID) *DiagnosticsClient {
	return &DiagnosticsClient{p: profile}
}

const DiagnosticsCluster zcl.ClusterID = 2821

type DiagnosticsServer struct {
	p zcl.ProfileID

	NumberOfResets            *NumberOfResets
	PersistensMemoryWrites    *PersistensMemoryWrites
	MacRxBcast                *MacRxBcast
	MacTxBcast                *MacTxBcast
	MacRxUcast                *MacRxUcast
	MacTxUcast                *MacTxUcast
	MacTxUcastRetry           *MacTxUcastRetry
	MacTxUcastFail            *MacTxUcastFail
	ApsRxBcast                *ApsRxBcast
	ApsTxBcast                *ApsTxBcast
	ApsRxUcast                *ApsRxUcast
	ApsTxUcastSuccess         *ApsTxUcastSuccess
	ApsTxUcastRetry           *ApsTxUcastRetry
	ApsTxUcastFail            *ApsTxUcastFail
	RouteDiscInitiated        *RouteDiscInitiated
	NeighborAdded             *NeighborAdded
	NeighborRemoved           *NeighborRemoved
	NeighborStale             *NeighborStale
	JoinIndication            *JoinIndication
	ChildMoved                *ChildMoved
	NwkFcFailure              *NwkFcFailure
	ApsFcFailure              *ApsFcFailure
	ApsUnauthorizedKey        *ApsUnauthorizedKey
	NwkDecryptFailures        *NwkDecryptFailures
	ApsDecryptFailures        *ApsDecryptFailures
	PacketBufferAllocFailures *PacketBufferAllocFailures
	RelayedUcast              *RelayedUcast
	PhyToMacQueueLimitReached *PhyToMacQueueLimitReached
	PacketValidateDropcount   *PacketValidateDropcount
	AvgMacRetryPerApsMsgSent  *AvgMacRetryPerApsMsgSent
	LastMessageLqi            *LastMessageLqi
	LastMessageRssi           *LastMessageRssi
}

type DiagnosticsClient struct {
	p zcl.ProfileID
}

/*
var DiagnosticsServer = map[zcl.CommandID]func() zcl.Command{
}

var DiagnosticsClient = map[zcl.CommandID]func() zcl.Command{
}
*/

type NumberOfResets zcl.Zu16

func (a NumberOfResets) ID() zcl.AttrID          { return 0 }
func (a NumberOfResets) Cluster() zcl.ClusterID  { return DiagnosticsCluster }
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

type PersistensMemoryWrites zcl.Zu16

func (a PersistensMemoryWrites) ID() zcl.AttrID                  { return 1 }
func (a PersistensMemoryWrites) Cluster() zcl.ClusterID          { return DiagnosticsCluster }
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

type MacRxBcast zcl.Zu32

func (a MacRxBcast) ID() zcl.AttrID         { return 256 }
func (a MacRxBcast) Cluster() zcl.ClusterID { return DiagnosticsCluster }
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

type MacTxBcast zcl.Zu32

func (a MacTxBcast) ID() zcl.AttrID         { return 257 }
func (a MacTxBcast) Cluster() zcl.ClusterID { return DiagnosticsCluster }
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

type MacRxUcast zcl.Zu32

func (a MacRxUcast) ID() zcl.AttrID         { return 258 }
func (a MacRxUcast) Cluster() zcl.ClusterID { return DiagnosticsCluster }
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

type MacTxUcast zcl.Zu32

func (a MacTxUcast) ID() zcl.AttrID         { return 259 }
func (a MacTxUcast) Cluster() zcl.ClusterID { return DiagnosticsCluster }
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

type MacTxUcastRetry zcl.Zu16

func (a MacTxUcastRetry) ID() zcl.AttrID           { return 260 }
func (a MacTxUcastRetry) Cluster() zcl.ClusterID   { return DiagnosticsCluster }
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

type MacTxUcastFail zcl.Zu16

func (a MacTxUcastFail) ID() zcl.AttrID          { return 261 }
func (a MacTxUcastFail) Cluster() zcl.ClusterID  { return DiagnosticsCluster }
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

type ApsRxBcast zcl.Zu16

func (a ApsRxBcast) ID() zcl.AttrID         { return 262 }
func (a ApsRxBcast) Cluster() zcl.ClusterID { return DiagnosticsCluster }
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

type ApsTxBcast zcl.Zu16

func (a ApsTxBcast) ID() zcl.AttrID         { return 263 }
func (a ApsTxBcast) Cluster() zcl.ClusterID { return DiagnosticsCluster }
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

type ApsRxUcast zcl.Zu16

func (a ApsRxUcast) ID() zcl.AttrID         { return 264 }
func (a ApsRxUcast) Cluster() zcl.ClusterID { return DiagnosticsCluster }
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

type ApsTxUcastSuccess zcl.Zu16

func (a ApsTxUcastSuccess) ID() zcl.AttrID             { return 265 }
func (a ApsTxUcastSuccess) Cluster() zcl.ClusterID     { return DiagnosticsCluster }
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

type ApsTxUcastRetry zcl.Zu16

func (a ApsTxUcastRetry) ID() zcl.AttrID           { return 266 }
func (a ApsTxUcastRetry) Cluster() zcl.ClusterID   { return DiagnosticsCluster }
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

type ApsTxUcastFail zcl.Zu16

func (a ApsTxUcastFail) ID() zcl.AttrID          { return 267 }
func (a ApsTxUcastFail) Cluster() zcl.ClusterID  { return DiagnosticsCluster }
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

type RouteDiscInitiated zcl.Zu16

func (a RouteDiscInitiated) ID() zcl.AttrID              { return 268 }
func (a RouteDiscInitiated) Cluster() zcl.ClusterID      { return DiagnosticsCluster }
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

type NeighborAdded zcl.Zu16

func (a NeighborAdded) ID() zcl.AttrID         { return 269 }
func (a NeighborAdded) Cluster() zcl.ClusterID { return DiagnosticsCluster }
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

type NeighborRemoved zcl.Zu16

func (a NeighborRemoved) ID() zcl.AttrID           { return 270 }
func (a NeighborRemoved) Cluster() zcl.ClusterID   { return DiagnosticsCluster }
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

type NeighborStale zcl.Zu16

func (a NeighborStale) ID() zcl.AttrID         { return 271 }
func (a NeighborStale) Cluster() zcl.ClusterID { return DiagnosticsCluster }
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

type JoinIndication zcl.Zu16

func (a JoinIndication) ID() zcl.AttrID          { return 272 }
func (a JoinIndication) Cluster() zcl.ClusterID  { return DiagnosticsCluster }
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

type ChildMoved zcl.Zu16

func (a ChildMoved) ID() zcl.AttrID         { return 273 }
func (a ChildMoved) Cluster() zcl.ClusterID { return DiagnosticsCluster }
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

type NwkFcFailure zcl.Zu16

func (a NwkFcFailure) ID() zcl.AttrID         { return 274 }
func (a NwkFcFailure) Cluster() zcl.ClusterID { return DiagnosticsCluster }
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

type ApsFcFailure zcl.Zu16

func (a ApsFcFailure) ID() zcl.AttrID         { return 275 }
func (a ApsFcFailure) Cluster() zcl.ClusterID { return DiagnosticsCluster }
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

type ApsUnauthorizedKey zcl.Zu16

func (a ApsUnauthorizedKey) ID() zcl.AttrID              { return 276 }
func (a ApsUnauthorizedKey) Cluster() zcl.ClusterID      { return DiagnosticsCluster }
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

type NwkDecryptFailures zcl.Zu16

func (a NwkDecryptFailures) ID() zcl.AttrID              { return 277 }
func (a NwkDecryptFailures) Cluster() zcl.ClusterID      { return DiagnosticsCluster }
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

type ApsDecryptFailures zcl.Zu16

func (a ApsDecryptFailures) ID() zcl.AttrID              { return 278 }
func (a ApsDecryptFailures) Cluster() zcl.ClusterID      { return DiagnosticsCluster }
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

type PacketBufferAllocFailures zcl.Zu16

func (a PacketBufferAllocFailures) ID() zcl.AttrID                     { return 279 }
func (a PacketBufferAllocFailures) Cluster() zcl.ClusterID             { return DiagnosticsCluster }
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

type RelayedUcast zcl.Zu16

func (a RelayedUcast) ID() zcl.AttrID         { return 280 }
func (a RelayedUcast) Cluster() zcl.ClusterID { return DiagnosticsCluster }
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

type PhyToMacQueueLimitReached zcl.Zu16

func (a PhyToMacQueueLimitReached) ID() zcl.AttrID                     { return 281 }
func (a PhyToMacQueueLimitReached) Cluster() zcl.ClusterID             { return DiagnosticsCluster }
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

type PacketValidateDropcount zcl.Zu16

func (a PacketValidateDropcount) ID() zcl.AttrID                   { return 282 }
func (a PacketValidateDropcount) Cluster() zcl.ClusterID           { return DiagnosticsCluster }
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

type AvgMacRetryPerApsMsgSent zcl.Zu16

func (a AvgMacRetryPerApsMsgSent) ID() zcl.AttrID                    { return 283 }
func (a AvgMacRetryPerApsMsgSent) Cluster() zcl.ClusterID            { return DiagnosticsCluster }
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

type LastMessageLqi zcl.Zu8

func (a LastMessageLqi) ID() zcl.AttrID          { return 284 }
func (a LastMessageLqi) Cluster() zcl.ClusterID  { return DiagnosticsCluster }
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

type LastMessageRssi zcl.Zs8

func (a LastMessageRssi) ID() zcl.AttrID           { return 285 }
func (a LastMessageRssi) Cluster() zcl.ClusterID   { return DiagnosticsCluster }
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
