// Attributes and commands for configuring On/Off switching devices
package general

import (
	"neotor.se/zcl"
)

// OnOffSwitchConfiguration
// Attributes and commands for configuring On/Off switching devices

func NewOnOffSwitchConfigurationServer(profile zcl.ProfileID) *OnOffSwitchConfigurationServer {
	return &OnOffSwitchConfigurationServer{p: profile}
}
func NewOnOffSwitchConfigurationClient(profile zcl.ProfileID) *OnOffSwitchConfigurationClient {
	return &OnOffSwitchConfigurationClient{p: profile}
}

const OnOffSwitchConfigurationCluster zcl.ClusterID = 7

type OnOffSwitchConfigurationServer struct {
	p zcl.ProfileID

	Switchtype    *Switchtype
	Switchactions *Switchactions
}

type OnOffSwitchConfigurationClient struct {
	p zcl.ProfileID
}

/*
var OnOffSwitchConfigurationServer = map[zcl.CommandID]func() zcl.Command{
}

var OnOffSwitchConfigurationClient = map[zcl.CommandID]func() zcl.Command{
}
*/

type Switchtype zcl.Zenum8

func (a Switchtype) ID() zcl.AttrID         { return 0 }
func (a Switchtype) Cluster() zcl.ClusterID { return OnOffSwitchConfigurationCluster }
func (a *Switchtype) Value() *Switchtype    { return a }
func (a Switchtype) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *Switchtype) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = Switchtype(*nt)
	return br, err
}

func (a Switchtype) Readable() bool   { return true }
func (a Switchtype) Writable() bool   { return false }
func (a Switchtype) Reportable() bool { return false }
func (a Switchtype) SceneIndex() int  { return -1 }

func (a Switchtype) String() string {
	switch a {
	case 0x00:
		return "Toggle"
	case 0x01:
		return "Momentary"
	case 0x02:
		return "Multifunction"
	}

	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsToggle checks if Switchtype equals the value for Toggle (0x00)
func (a Switchtype) IsToggle() bool { return a == 0x00 }

// SetToggle sets Switchtype to Toggle (0x00)
func (a *Switchtype) SetToggle() { *a = 0x00 }

// IsMomentary checks if Switchtype equals the value for Momentary (0x01)
func (a Switchtype) IsMomentary() bool { return a == 0x01 }

// SetMomentary sets Switchtype to Momentary (0x01)
func (a *Switchtype) SetMomentary() { *a = 0x01 }

// IsMultifunction checks if Switchtype equals the value for Multifunction (0x02)
func (a Switchtype) IsMultifunction() bool { return a == 0x02 }

// SetMultifunction sets Switchtype to Multifunction (0x02)
func (a *Switchtype) SetMultifunction() { *a = 0x02 }

type Switchactions zcl.Zenum8

func (a Switchactions) ID() zcl.AttrID         { return 16 }
func (a Switchactions) Cluster() zcl.ClusterID { return OnOffSwitchConfigurationCluster }
func (a *Switchactions) Value() *Switchactions { return a }
func (a Switchactions) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *Switchactions) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = Switchactions(*nt)
	return br, err
}

func (a Switchactions) Readable() bool   { return true }
func (a Switchactions) Writable() bool   { return true }
func (a Switchactions) Reportable() bool { return false }
func (a Switchactions) SceneIndex() int  { return -1 }

func (a Switchactions) String() string {
	switch a {
	case 0x00:
		return "On-Off"
	case 0x01:
		return "Off-On"
	case 0x02:
		return "Toggle"
	}

	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsOnOff checks if Switchactions equals the value for On-Off (0x00)
func (a Switchactions) IsOnOff() bool { return a == 0x00 }

// SetOnOff sets Switchactions to On-Off (0x00)
func (a *Switchactions) SetOnOff() { *a = 0x00 }

// IsOffOn checks if Switchactions equals the value for Off-On (0x01)
func (a Switchactions) IsOffOn() bool { return a == 0x01 }

// SetOffOn sets Switchactions to Off-On (0x01)
func (a *Switchactions) SetOffOn() { *a = 0x01 }

// IsToggle checks if Switchactions equals the value for Toggle (0x02)
func (a Switchactions) IsToggle() bool { return a == 0x02 }

// SetToggle sets Switchactions to Toggle (0x02)
func (a *Switchactions) SetToggle() { *a = 0x02 }
