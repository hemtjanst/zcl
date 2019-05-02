// The Price Cluster provides the mechanism for communicating Gas, Energy, or Water pricing information within the premise. This pricing information is distributed to the ESP from either the utilities or from regional energy providers.
package smart_energy

import (
	"neotor.se/zcl"
)

// Price
const PriceID zcl.ClusterID = 1792

var PriceCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		GetCurrentPriceCommand:    func() zcl.Command { return new(GetCurrentPrice) },
		GetScheduledPricesCommand: func() zcl.Command { return new(GetScheduledPrices) },
	},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		Tier1PriceLabelAttr: func() zcl.Attr { return new(Tier1PriceLabel) },
		Tier2PriceLabelAttr: func() zcl.Attr { return new(Tier2PriceLabel) },
		Tier3PriceLabelAttr: func() zcl.Attr { return new(Tier3PriceLabel) },
		Tier4PriceLabelAttr: func() zcl.Attr { return new(Tier4PriceLabel) },
		Tier5PriceLabelAttr: func() zcl.Attr { return new(Tier5PriceLabel) },
		Tier6PriceLabelAttr: func() zcl.Attr { return new(Tier6PriceLabel) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

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
	return PriceID
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
	return PriceID
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

const Tier1PriceLabelAttr zcl.AttrID = 0

type Tier1PriceLabel zcl.Zostring

func (a Tier1PriceLabel) ID() zcl.AttrID           { return Tier1PriceLabelAttr }
func (a Tier1PriceLabel) Cluster() zcl.ClusterID   { return PriceID }
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

const Tier2PriceLabelAttr zcl.AttrID = 1

type Tier2PriceLabel zcl.Zostring

func (a Tier2PriceLabel) ID() zcl.AttrID           { return Tier2PriceLabelAttr }
func (a Tier2PriceLabel) Cluster() zcl.ClusterID   { return PriceID }
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

const Tier3PriceLabelAttr zcl.AttrID = 2

type Tier3PriceLabel zcl.Zostring

func (a Tier3PriceLabel) ID() zcl.AttrID           { return Tier3PriceLabelAttr }
func (a Tier3PriceLabel) Cluster() zcl.ClusterID   { return PriceID }
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

const Tier4PriceLabelAttr zcl.AttrID = 3

type Tier4PriceLabel zcl.Zostring

func (a Tier4PriceLabel) ID() zcl.AttrID           { return Tier4PriceLabelAttr }
func (a Tier4PriceLabel) Cluster() zcl.ClusterID   { return PriceID }
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

const Tier5PriceLabelAttr zcl.AttrID = 4

type Tier5PriceLabel zcl.Zostring

func (a Tier5PriceLabel) ID() zcl.AttrID           { return Tier5PriceLabelAttr }
func (a Tier5PriceLabel) Cluster() zcl.ClusterID   { return PriceID }
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

const Tier6PriceLabelAttr zcl.AttrID = 5

type Tier6PriceLabel zcl.Zostring

func (a Tier6PriceLabel) ID() zcl.AttrID           { return Tier6PriceLabelAttr }
func (a Tier6PriceLabel) Cluster() zcl.ClusterID   { return PriceID }
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