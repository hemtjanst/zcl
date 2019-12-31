package general

import "hemtjan.st/zcl"

// Time
const TimeID zcl.ClusterID = 10

var TimeCluster = zcl.Cluster{
	Name:      "Time",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		TimeAttr:           func() zcl.Attr { return new(Time) },
		TimeStatusAttr:     func() zcl.Attr { return new(TimeStatus) },
		TimeZoneAttr:       func() zcl.Attr { return new(TimeZone) },
		DstStartAttr:       func() zcl.Attr { return new(DstStart) },
		DstEndAttr:         func() zcl.Attr { return new(DstEnd) },
		DstShiftAttr:       func() zcl.Attr { return new(DstShift) },
		StandardTimeAttr:   func() zcl.Attr { return new(StandardTime) },
		LocalTimeAttr:      func() zcl.Attr { return new(LocalTime) },
		LastSetTimeAttr:    func() zcl.Attr { return new(LastSetTime) },
		ValidUntilTimeAttr: func() zcl.Attr { return new(ValidUntilTime) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
