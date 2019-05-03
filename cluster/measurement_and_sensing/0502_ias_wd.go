// The IAS WD cluster provides an interface to the functionality of any Warning Device equipment of the IAS system. Using this cluster, a ZigBee enabled CIE device can access a ZigBee enabled IAS WD device and issue alarm warning indications (siren, strobe lighting, etc.) when a system alarm condition is detected.
package measurement_and_sensing

import (
	"neotor.se/zcl"
)

// IasWd
const IasWdID zcl.ClusterID = 1282

var IasWdCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		StartWarningCommand: func() zcl.Command { return new(StartWarning) },
		SquawkCommand:       func() zcl.Command { return new(Squawk) },
	},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		MaxDurationAttr: func() zcl.Attr { return new(MaxDuration) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

type StartWarning struct {
	Options         zcl.Zbmp8
	WarningDuration zcl.Zu16
	StrobeDutyCycle zcl.Zu8
	StrobeLevel     zcl.Zenum8
}

const StartWarningCommand zcl.CommandID = 0

func (v *StartWarning) Values() []zcl.Val {
	return []zcl.Val{
		&v.Options,
		&v.WarningDuration,
		&v.StrobeDutyCycle,
		&v.StrobeLevel,
	}
}

func (v StartWarning) ID() zcl.CommandID {
	return StartWarningCommand
}

func (v StartWarning) Cluster() zcl.ClusterID {
	return IasWdID
}

func (v StartWarning) MnfCode() []byte {
	return []byte{}
}

func (v StartWarning) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Options.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.WarningDuration.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.StrobeDutyCycle.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.StrobeLevel.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *StartWarning) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Options).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.WarningDuration).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.StrobeDutyCycle).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.StrobeLevel).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v StartWarning) OptionsString() string {
	var bstr []string
	if zcl.BitmapTest([]byte(v.Options), 0) {
		bstr = append(bstr, "Siren level 0")
	}
	if zcl.BitmapTest([]byte(v.Options), 1) {
		bstr = append(bstr, "Siren level 1")
	}
	if zcl.BitmapTest([]byte(v.Options), 2) {
		bstr = append(bstr, "Strobe")
	}
	if zcl.BitmapTest([]byte(v.Options), 4) {
		bstr = append(bstr, "Warning mode 0")
	}
	if zcl.BitmapTest([]byte(v.Options), 5) {
		bstr = append(bstr, "Warning mode 1")
	}
	if zcl.BitmapTest([]byte(v.Options), 6) {
		bstr = append(bstr, "Warning mode 2")
	}
	if zcl.BitmapTest([]byte(v.Options), 7) {
		bstr = append(bstr, "Warning mode 3")
	}
	return zcl.StrJoin(bstr, ", ")
}
func (v StartWarning) WarningDurationString() string {
	return zcl.Sprintf("%v", zcl.Zu16(v.WarningDuration))
}
func (v StartWarning) StrobeDutyCycleString() string {
	return zcl.Sprintf("%v", zcl.Zu8(v.StrobeDutyCycle))
}
func (v StartWarning) StrobeLevelString() string {
	switch v.StrobeLevel {
	case 0x00:
		return "Low"
	case 0x01:
		return "Medium"
	case 0x02:
		return "High"
	case 0x03:
		return "Very high"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v.StrobeLevel))
}

func (v StartWarning) String() string {
	var str []string
	str = append(str, "Options["+v.OptionsString()+"]")
	str = append(str, "WarningDuration["+v.WarningDurationString()+"]")
	str = append(str, "StrobeDutyCycle["+v.StrobeDutyCycleString()+"]")
	str = append(str, "StrobeLevel["+v.StrobeLevelString()+"]")
	return "StartWarning{" + zcl.StrJoin(str, " ") + "}"
}

func (StartWarning) Name() string { return "Start warning" }

type Squawk struct {
	Options zcl.Zbmp8
}

const SquawkCommand zcl.CommandID = 1

func (v *Squawk) Values() []zcl.Val {
	return []zcl.Val{
		&v.Options,
	}
}

func (v Squawk) ID() zcl.CommandID {
	return SquawkCommand
}

func (v Squawk) Cluster() zcl.ClusterID {
	return IasWdID
}

func (v Squawk) MnfCode() []byte {
	return []byte{}
}

func (v Squawk) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Options.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *Squawk) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Options).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v Squawk) OptionsString() string {
	var bstr []string
	if zcl.BitmapTest([]byte(v.Options), 0) {
		bstr = append(bstr, "Sqawk level 0")
	}
	if zcl.BitmapTest([]byte(v.Options), 1) {
		bstr = append(bstr, "Sqawk level 1")
	}
	if zcl.BitmapTest([]byte(v.Options), 3) {
		bstr = append(bstr, "Strobe")
	}
	if zcl.BitmapTest([]byte(v.Options), 4) {
		bstr = append(bstr, "Sound for disarmed")
	}
	return zcl.StrJoin(bstr, ", ")
}

func (v Squawk) String() string {
	var str []string
	str = append(str, "Options["+v.OptionsString()+"]")
	return "Squawk{" + zcl.StrJoin(str, " ") + "}"
}

func (Squawk) Name() string { return "Squawk" }

// MaxDuration is an autogenerated attribute in the IasWd cluster
type MaxDuration zcl.Zu16

const MaxDurationAttr zcl.AttrID = 0

func (a MaxDuration) ID() zcl.AttrID         { return MaxDurationAttr }
func (a MaxDuration) Cluster() zcl.ClusterID { return IasWdID }
func (a *MaxDuration) Value() *MaxDuration   { return a }
func (a MaxDuration) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *MaxDuration) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MaxDuration(*nt)
	return br, err
}
func (MaxDuration) Name() string     { return "Max Duration" }
func (MaxDuration) Readable() bool   { return true }
func (MaxDuration) Writable() bool   { return true }
func (MaxDuration) Reportable() bool { return false }
func (MaxDuration) SceneIndex() int  { return -1 }

func (a MaxDuration) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}
