// The Price Cluster provides the mechanism for communicating Gas, Energy, or Water pricing information within the premise. This pricing information is distributed to the ESP from either the utilities or from regional energy providers.
package smart_energy

import (
	"neotor.se/zcl"
)

// Price
// The Price Cluster provides the mechanism for communicating Gas, Energy, or Water pricing information within the premise. This pricing information is distributed to the ESP from either the utilities or from regional energy providers.

func NewPriceServer(profile zcl.ProfileID) *PriceServer { return &PriceServer{p: profile} }
func NewPriceClient(profile zcl.ProfileID) *PriceClient { return &PriceClient{p: profile} }

const PriceCluster zcl.ClusterID = 1792

type PriceServer struct {
	p zcl.ProfileID

	Tier1PriceLabel *Tier1PriceLabel
	Tier2PriceLabel *Tier2PriceLabel
	Tier3PriceLabel *Tier3PriceLabel
	Tier4PriceLabel *Tier4PriceLabel
	Tier5PriceLabel *Tier5PriceLabel
	Tier6PriceLabel *Tier6PriceLabel
}

func (s *PriceServer) GetCurrentPrice() *GetCurrentPrice       { return new(GetCurrentPrice) }
func (s *PriceServer) GetScheduledPrices() *GetScheduledPrices { return new(GetScheduledPrices) }

type PriceClient struct {
	p zcl.ProfileID
}

/*
var PriceServer = map[zcl.CommandID]func() zcl.Command{
    GetCurrentPriceID: func() zcl.Command { return new(GetCurrentPrice) },
    GetScheduledPricesID: func() zcl.Command { return new(GetScheduledPrices) },
}

var PriceClient = map[zcl.CommandID]func() zcl.Command{
}
*/

type GetCurrentPrice struct {
	CommandOptions zcl.Zbmp8
}

const GetCurrentPriceCommand zcl.CommandID = 0

func (v *GetCurrentPrice) Values() []zcl.Val {
	return []zcl.Val{
		&v.CommandOptions,
	}
}

func (v GetCurrentPrice) ID() zcl.CommandID {
	return GetCurrentPriceCommand
}

func (v GetCurrentPrice) Cluster() zcl.ClusterID {
	return PriceCluster
}

func (v GetCurrentPrice) MnfCode() []byte {
	return []byte{}
}

func (v GetCurrentPrice) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.CommandOptions.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *GetCurrentPrice) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.CommandOptions).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type GetScheduledPrices struct {
	StartTime      zcl.Zutc
	NumberOfEvents zcl.Zu8
}

const GetScheduledPricesCommand zcl.CommandID = 1

func (v *GetScheduledPrices) Values() []zcl.Val {
	return []zcl.Val{
		&v.StartTime,
		&v.NumberOfEvents,
	}
}

func (v GetScheduledPrices) ID() zcl.CommandID {
	return GetScheduledPricesCommand
}

func (v GetScheduledPrices) Cluster() zcl.ClusterID {
	return PriceCluster
}

func (v GetScheduledPrices) MnfCode() []byte {
	return []byte{}
}

func (v GetScheduledPrices) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.StartTime.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.NumberOfEvents.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *GetScheduledPrices) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.StartTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.NumberOfEvents).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type Tier1PriceLabel zcl.Zostring

func (a Tier1PriceLabel) ID() zcl.AttrID           { return 0 }
func (a Tier1PriceLabel) Cluster() zcl.ClusterID   { return PriceCluster }
func (a *Tier1PriceLabel) Value() *Tier1PriceLabel { return a }
func (a Tier1PriceLabel) MarshalZcl() ([]byte, error) {
	return zcl.Zostring(a).MarshalZcl()
}
func (a *Tier1PriceLabel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zostring)
	br, err := nt.UnmarshalZcl(b)
	*a = Tier1PriceLabel(*nt)
	return br, err
}

func (a Tier1PriceLabel) Readable() bool   { return true }
func (a Tier1PriceLabel) Writable() bool   { return true }
func (a Tier1PriceLabel) Reportable() bool { return false }
func (a Tier1PriceLabel) SceneIndex() int  { return -1 }

func (a Tier1PriceLabel) String() string {

	return zcl.Sprintf("%s", zcl.Zostring(a))
}

type Tier2PriceLabel zcl.Zostring

func (a Tier2PriceLabel) ID() zcl.AttrID           { return 1 }
func (a Tier2PriceLabel) Cluster() zcl.ClusterID   { return PriceCluster }
func (a *Tier2PriceLabel) Value() *Tier2PriceLabel { return a }
func (a Tier2PriceLabel) MarshalZcl() ([]byte, error) {
	return zcl.Zostring(a).MarshalZcl()
}
func (a *Tier2PriceLabel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zostring)
	br, err := nt.UnmarshalZcl(b)
	*a = Tier2PriceLabel(*nt)
	return br, err
}

func (a Tier2PriceLabel) Readable() bool   { return true }
func (a Tier2PriceLabel) Writable() bool   { return true }
func (a Tier2PriceLabel) Reportable() bool { return false }
func (a Tier2PriceLabel) SceneIndex() int  { return -1 }

func (a Tier2PriceLabel) String() string {

	return zcl.Sprintf("%s", zcl.Zostring(a))
}

type Tier3PriceLabel zcl.Zostring

func (a Tier3PriceLabel) ID() zcl.AttrID           { return 2 }
func (a Tier3PriceLabel) Cluster() zcl.ClusterID   { return PriceCluster }
func (a *Tier3PriceLabel) Value() *Tier3PriceLabel { return a }
func (a Tier3PriceLabel) MarshalZcl() ([]byte, error) {
	return zcl.Zostring(a).MarshalZcl()
}
func (a *Tier3PriceLabel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zostring)
	br, err := nt.UnmarshalZcl(b)
	*a = Tier3PriceLabel(*nt)
	return br, err
}

func (a Tier3PriceLabel) Readable() bool   { return true }
func (a Tier3PriceLabel) Writable() bool   { return true }
func (a Tier3PriceLabel) Reportable() bool { return false }
func (a Tier3PriceLabel) SceneIndex() int  { return -1 }

func (a Tier3PriceLabel) String() string {

	return zcl.Sprintf("%s", zcl.Zostring(a))
}

type Tier4PriceLabel zcl.Zostring

func (a Tier4PriceLabel) ID() zcl.AttrID           { return 3 }
func (a Tier4PriceLabel) Cluster() zcl.ClusterID   { return PriceCluster }
func (a *Tier4PriceLabel) Value() *Tier4PriceLabel { return a }
func (a Tier4PriceLabel) MarshalZcl() ([]byte, error) {
	return zcl.Zostring(a).MarshalZcl()
}
func (a *Tier4PriceLabel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zostring)
	br, err := nt.UnmarshalZcl(b)
	*a = Tier4PriceLabel(*nt)
	return br, err
}

func (a Tier4PriceLabel) Readable() bool   { return true }
func (a Tier4PriceLabel) Writable() bool   { return true }
func (a Tier4PriceLabel) Reportable() bool { return false }
func (a Tier4PriceLabel) SceneIndex() int  { return -1 }

func (a Tier4PriceLabel) String() string {

	return zcl.Sprintf("%s", zcl.Zostring(a))
}

type Tier5PriceLabel zcl.Zostring

func (a Tier5PriceLabel) ID() zcl.AttrID           { return 4 }
func (a Tier5PriceLabel) Cluster() zcl.ClusterID   { return PriceCluster }
func (a *Tier5PriceLabel) Value() *Tier5PriceLabel { return a }
func (a Tier5PriceLabel) MarshalZcl() ([]byte, error) {
	return zcl.Zostring(a).MarshalZcl()
}
func (a *Tier5PriceLabel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zostring)
	br, err := nt.UnmarshalZcl(b)
	*a = Tier5PriceLabel(*nt)
	return br, err
}

func (a Tier5PriceLabel) Readable() bool   { return true }
func (a Tier5PriceLabel) Writable() bool   { return true }
func (a Tier5PriceLabel) Reportable() bool { return false }
func (a Tier5PriceLabel) SceneIndex() int  { return -1 }

func (a Tier5PriceLabel) String() string {

	return zcl.Sprintf("%s", zcl.Zostring(a))
}

type Tier6PriceLabel zcl.Zostring

func (a Tier6PriceLabel) ID() zcl.AttrID           { return 5 }
func (a Tier6PriceLabel) Cluster() zcl.ClusterID   { return PriceCluster }
func (a *Tier6PriceLabel) Value() *Tier6PriceLabel { return a }
func (a Tier6PriceLabel) MarshalZcl() ([]byte, error) {
	return zcl.Zostring(a).MarshalZcl()
}
func (a *Tier6PriceLabel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zostring)
	br, err := nt.UnmarshalZcl(b)
	*a = Tier6PriceLabel(*nt)
	return br, err
}

func (a Tier6PriceLabel) Readable() bool   { return true }
func (a Tier6PriceLabel) Writable() bool   { return true }
func (a Tier6PriceLabel) Reportable() bool { return false }
func (a Tier6PriceLabel) SceneIndex() int  { return -1 }

func (a Tier6PriceLabel) String() string {

	return zcl.Sprintf("%s", zcl.Zostring(a))
}
