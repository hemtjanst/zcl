// Measure distance between devices.
package general

import (
	"neotor.se/zcl"
)

// Location
const LocationID zcl.ClusterID = 11

var LocationCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		SetAbsoluteLocationCommand:    func() zcl.Command { return new(SetAbsoluteLocation) },
		SetDeviceConfigurationCommand: func() zcl.Command { return new(SetDeviceConfiguration) },
		GetDeviceConfigurationCommand: func() zcl.Command { return new(GetDeviceConfiguration) },
		GetLocationDataCommand:        func() zcl.Command { return new(GetLocationData) },
		RssiResponseCommand:           func() zcl.Command { return new(RssiResponse) },
		SendPingsCommand:              func() zcl.Command { return new(SendPings) },
		AnchorNodeAnnounceCommand:     func() zcl.Command { return new(AnchorNodeAnnounce) },
		DistanceMeasureCommand:        func() zcl.Command { return new(DistanceMeasure) },
	},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{
		DeviceConfigurationResponseCommand:     func() zcl.Command { return new(DeviceConfigurationResponse) },
		LocationDataResponseCommand:            func() zcl.Command { return new(LocationDataResponse) },
		LocationDataNotificationCommand:        func() zcl.Command { return new(LocationDataNotification) },
		CompactLocationDataNotificationCommand: func() zcl.Command { return new(CompactLocationDataNotification) },
		RssiPingCommand:                        func() zcl.Command { return new(RssiPing) },
		RssiRequestCommand:                     func() zcl.Command { return new(RssiRequest) },
		ReportRssiMeasurementsCommand:          func() zcl.Command { return new(ReportRssiMeasurements) },
		RequestOwnLocationCommand:              func() zcl.Command { return new(RequestOwnLocation) },
		DistanceMeasureResponseCommand:         func() zcl.Command { return new(DistanceMeasureResponse) },
	},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		LocationTypeAttr:           func() zcl.Attr { return new(LocationType) },
		LocationMethodAttr:         func() zcl.Attr { return new(LocationMethod) },
		LocationAgeAttr:            func() zcl.Attr { return new(LocationAge) },
		QualityMeasureAttr:         func() zcl.Attr { return new(QualityMeasure) },
		NumberOfDevicesAttr:        func() zcl.Attr { return new(NumberOfDevices) },
		XCoordinateAttr:            func() zcl.Attr { return new(XCoordinate) },
		YCoordinateAttr:            func() zcl.Attr { return new(YCoordinate) },
		ZCoordinateAttr:            func() zcl.Attr { return new(ZCoordinate) },
		PowerAttr:                  func() zcl.Attr { return new(Power) },
		PathLossExponentAttr:       func() zcl.Attr { return new(PathLossExponent) },
		ReportingPeriodAttr:        func() zcl.Attr { return new(ReportingPeriod) },
		CalculationPeriodAttr:      func() zcl.Attr { return new(CalculationPeriod) },
		NumberRssiMeasurementsAttr: func() zcl.Attr { return new(NumberRssiMeasurements) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

type SetAbsoluteLocation struct {
	XCoordinate zcl.Zs16
	YCoordinate zcl.Zs16
	ZCoordinate zcl.Zs16
	Power       zcl.Zs16
	// Rate at which the signal power decays with increasing distance
	PathLossExponent zcl.Zu16
}

const SetAbsoluteLocationCommand zcl.CommandID = 0

func (v *SetAbsoluteLocation) Values() []zcl.Val {
	return []zcl.Val{
		&v.XCoordinate,
		&v.YCoordinate,
		&v.ZCoordinate,
		&v.Power,
		&v.PathLossExponent,
	}
}

func (v SetAbsoluteLocation) ID() zcl.CommandID {
	return SetAbsoluteLocationCommand
}

func (v SetAbsoluteLocation) Cluster() zcl.ClusterID {
	return LocationID
}

func (v SetAbsoluteLocation) MnfCode() []byte {
	return []byte{}
}

func (v SetAbsoluteLocation) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.XCoordinate.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.YCoordinate.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ZCoordinate.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Power.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.PathLossExponent.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *SetAbsoluteLocation) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.XCoordinate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.YCoordinate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ZCoordinate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Power).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.PathLossExponent).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type SetDeviceConfiguration struct {
	Power zcl.Zs16
	// Rate at which the signal power decays with increasing distance
	PathLossExponent  zcl.Zu16
	CalculationPeriod zcl.Zu16
	// Number of measurements to use to generate one location estimate
	NumberRssiMeasurements zcl.Zu8
	ReportingPeriod        zcl.Zu16
}

const SetDeviceConfigurationCommand zcl.CommandID = 1

func (v *SetDeviceConfiguration) Values() []zcl.Val {
	return []zcl.Val{
		&v.Power,
		&v.PathLossExponent,
		&v.CalculationPeriod,
		&v.NumberRssiMeasurements,
		&v.ReportingPeriod,
	}
}

func (v SetDeviceConfiguration) ID() zcl.CommandID {
	return SetDeviceConfigurationCommand
}

func (v SetDeviceConfiguration) Cluster() zcl.ClusterID {
	return LocationID
}

func (v SetDeviceConfiguration) MnfCode() []byte {
	return []byte{}
}

func (v SetDeviceConfiguration) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Power.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.PathLossExponent.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.CalculationPeriod.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.NumberRssiMeasurements.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ReportingPeriod.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *SetDeviceConfiguration) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Power).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.PathLossExponent).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.CalculationPeriod).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.NumberRssiMeasurements).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ReportingPeriod).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type GetDeviceConfiguration struct {
	TargetAddress zcl.Zuid
}

const GetDeviceConfigurationCommand zcl.CommandID = 2

func (v *GetDeviceConfiguration) Values() []zcl.Val {
	return []zcl.Val{
		&v.TargetAddress,
	}
}

func (v GetDeviceConfiguration) ID() zcl.CommandID {
	return GetDeviceConfigurationCommand
}

func (v GetDeviceConfiguration) Cluster() zcl.ClusterID {
	return LocationID
}

func (v GetDeviceConfiguration) MnfCode() []byte {
	return []byte{}
}

func (v GetDeviceConfiguration) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.TargetAddress.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *GetDeviceConfiguration) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.TargetAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type GetLocationData struct {
	Flags           zcl.Zbmp8
	NumberResponses zcl.Zu8
	TargetAddress   zcl.Zuid
}

const GetLocationDataCommand zcl.CommandID = 3

func (v *GetLocationData) Values() []zcl.Val {
	return []zcl.Val{
		&v.Flags,
		&v.NumberResponses,
		&v.TargetAddress,
	}
}

func (v GetLocationData) ID() zcl.CommandID {
	return GetLocationDataCommand
}

func (v GetLocationData) Cluster() zcl.ClusterID {
	return LocationID
}

func (v GetLocationData) MnfCode() []byte {
	return []byte{}
}

func (v GetLocationData) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Flags.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.NumberResponses.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.TargetAddress.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *GetLocationData) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Flags).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.NumberResponses).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.TargetAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type RssiResponse struct {
	ReplyingDevice         zcl.Zuid
	XCoordinate            zcl.Zs16
	YCoordinate            zcl.Zs16
	ZCoordinate            zcl.Zs16
	Rssi                   zcl.Zs8
	NumberRssiMeasurements zcl.Zu8
}

const RssiResponseCommand zcl.CommandID = 4

func (v *RssiResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.ReplyingDevice,
		&v.XCoordinate,
		&v.YCoordinate,
		&v.ZCoordinate,
		&v.Rssi,
		&v.NumberRssiMeasurements,
	}
}

func (v RssiResponse) ID() zcl.CommandID {
	return RssiResponseCommand
}

func (v RssiResponse) Cluster() zcl.ClusterID {
	return LocationID
}

func (v RssiResponse) MnfCode() []byte {
	return []byte{}
}

func (v RssiResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.ReplyingDevice.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.XCoordinate.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.YCoordinate.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ZCoordinate.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Rssi.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.NumberRssiMeasurements.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *RssiResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.ReplyingDevice).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.XCoordinate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.YCoordinate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ZCoordinate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Rssi).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.NumberRssiMeasurements).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type SendPings struct {
	TargetAddress zcl.Zuid
	// Number of measurements to use to generate one location estimate
	NumberRssiMeasurements zcl.Zu8
	CalculationPeriod      zcl.Zu16
}

const SendPingsCommand zcl.CommandID = 5

func (v *SendPings) Values() []zcl.Val {
	return []zcl.Val{
		&v.TargetAddress,
		&v.NumberRssiMeasurements,
		&v.CalculationPeriod,
	}
}

func (v SendPings) ID() zcl.CommandID {
	return SendPingsCommand
}

func (v SendPings) Cluster() zcl.ClusterID {
	return LocationID
}

func (v SendPings) MnfCode() []byte {
	return []byte{}
}

func (v SendPings) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.TargetAddress.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.NumberRssiMeasurements.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.CalculationPeriod.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *SendPings) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.TargetAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.NumberRssiMeasurements).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.CalculationPeriod).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type AnchorNodeAnnounce struct {
	AnchorNodeAddress zcl.Zuid
	XCoordinate       zcl.Zs16
	YCoordinate       zcl.Zs16
	ZCoordinate       zcl.Zs16
}

const AnchorNodeAnnounceCommand zcl.CommandID = 6

func (v *AnchorNodeAnnounce) Values() []zcl.Val {
	return []zcl.Val{
		&v.AnchorNodeAddress,
		&v.XCoordinate,
		&v.YCoordinate,
		&v.ZCoordinate,
	}
}

func (v AnchorNodeAnnounce) ID() zcl.CommandID {
	return AnchorNodeAnnounceCommand
}

func (v AnchorNodeAnnounce) Cluster() zcl.ClusterID {
	return LocationID
}

func (v AnchorNodeAnnounce) MnfCode() []byte {
	return []byte{}
}

func (v AnchorNodeAnnounce) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.AnchorNodeAddress.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.XCoordinate.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.YCoordinate.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ZCoordinate.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *AnchorNodeAnnounce) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.AnchorNodeAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.XCoordinate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.YCoordinate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ZCoordinate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type DistanceMeasure struct {
	TargetAddress zcl.Zu16
	Resolution    zcl.Zenum8
}

const DistanceMeasureCommand zcl.CommandID = 64

func (v *DistanceMeasure) Values() []zcl.Val {
	return []zcl.Val{
		&v.TargetAddress,
		&v.Resolution,
	}
}

func (v DistanceMeasure) ID() zcl.CommandID {
	return DistanceMeasureCommand
}

func (v DistanceMeasure) Cluster() zcl.ClusterID {
	return LocationID
}

func (v DistanceMeasure) MnfCode() []byte {
	return []byte{}
}

func (v DistanceMeasure) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.TargetAddress.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Resolution.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *DistanceMeasure) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.TargetAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Resolution).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type DeviceConfigurationResponse struct {
	Status zcl.Status
	Power  zcl.Zs16
	// Rate at which the signal power decays with increasing distance
	PathLossExponent  zcl.Zu16
	CalculationPeriod zcl.Zu16
	// Number of measurements to use to generate one location estimate
	NumberRssiMeasurements zcl.Zu8
	ReportingPeriod        zcl.Zu16
}

const DeviceConfigurationResponseCommand zcl.CommandID = 0

func (v *DeviceConfigurationResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.Power,
		&v.PathLossExponent,
		&v.CalculationPeriod,
		&v.NumberRssiMeasurements,
		&v.ReportingPeriod,
	}
}

func (v DeviceConfigurationResponse) ID() zcl.CommandID {
	return DeviceConfigurationResponseCommand
}

func (v DeviceConfigurationResponse) Cluster() zcl.ClusterID {
	return LocationID
}

func (v DeviceConfigurationResponse) MnfCode() []byte {
	return []byte{}
}

func (v DeviceConfigurationResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Status.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if v.Status == 0x00 {
		if tmp, err = v.Power.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}

	if v.Status == 0x00 {
		if tmp, err = v.PathLossExponent.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}

	if v.Status == 0x00 {
		if tmp, err = v.CalculationPeriod.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}

	if v.Status == 0x00 {
		if tmp, err = v.NumberRssiMeasurements.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}

	if v.Status == 0x00 {
		if tmp, err = v.ReportingPeriod.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}

	return data, nil
}

func (v *DeviceConfigurationResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if v.Status == 0x00 {
		if b, err = (&v.Power).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	if v.Status == 0x00 {
		if b, err = (&v.PathLossExponent).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	if v.Status == 0x00 {
		if b, err = (&v.CalculationPeriod).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	if v.Status == 0x00 {
		if b, err = (&v.NumberRssiMeasurements).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	if v.Status == 0x00 {
		if b, err = (&v.ReportingPeriod).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	return b, nil
}

type LocationDataResponse struct {
	Status       zcl.Status
	LocationType zcl.Zenum8
	XCoordinate  zcl.Zs16
	YCoordinate  zcl.Zs16
	ZCoordinate  zcl.Zs16
	Power        zcl.Zs16
	// Rate at which the signal power decays with increasing distance
	PathLossExponent zcl.Zu16
	LocationMethod   zcl.Zenum8
	QualityMeasure   zcl.Zu8
	LocationAge      zcl.Zu16
}

const LocationDataResponseCommand zcl.CommandID = 1

func (v *LocationDataResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.LocationType,
		&v.XCoordinate,
		&v.YCoordinate,
		&v.ZCoordinate,
		&v.Power,
		&v.PathLossExponent,
		&v.LocationMethod,
		&v.QualityMeasure,
		&v.LocationAge,
	}
}

func (v LocationDataResponse) ID() zcl.CommandID {
	return LocationDataResponseCommand
}

func (v LocationDataResponse) Cluster() zcl.ClusterID {
	return LocationID
}

func (v LocationDataResponse) MnfCode() []byte {
	return []byte{}
}

func (v LocationDataResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Status.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if v.Status == 0x00 {
		if tmp, err = v.LocationType.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}

	if v.Status == 0x00 {
		if tmp, err = v.XCoordinate.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}

	if v.Status == 0x00 {
		if tmp, err = v.YCoordinate.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}

	if v.Status == 0x00 && (v.LocationType&0x02) == 0x00 {
		if tmp, err = v.ZCoordinate.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}

	if v.Status == 0x00 {
		if tmp, err = v.Power.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}

	if v.Status == 0x00 {
		if tmp, err = v.PathLossExponent.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}

	if v.Status == 0x00 && (v.LocationType&0x01) == 0x00 {
		if tmp, err = v.LocationMethod.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}

	if v.Status == 0x00 && (v.LocationType&0x01) == 0x00 {
		if tmp, err = v.QualityMeasure.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}

	if v.Status == 0x00 && (v.LocationType&0x01) == 0x00 {
		if tmp, err = v.LocationAge.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}

	return data, nil
}

func (v *LocationDataResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if v.Status == 0x00 {
		if b, err = (&v.LocationType).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	if v.Status == 0x00 {
		if b, err = (&v.XCoordinate).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	if v.Status == 0x00 {
		if b, err = (&v.YCoordinate).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	if v.Status == 0x00 && (v.LocationType&0x02) == 0x00 {
		if b, err = (&v.ZCoordinate).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	if v.Status == 0x00 {
		if b, err = (&v.Power).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	if v.Status == 0x00 {
		if b, err = (&v.PathLossExponent).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	if v.Status == 0x00 && (v.LocationType&0x01) == 0x00 {
		if b, err = (&v.LocationMethod).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	if v.Status == 0x00 && (v.LocationType&0x01) == 0x00 {
		if b, err = (&v.QualityMeasure).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	if v.Status == 0x00 && (v.LocationType&0x01) == 0x00 {
		if b, err = (&v.LocationAge).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	return b, nil
}

type LocationDataNotification struct {
	LocationType zcl.Zenum8
	XCoordinate  zcl.Zs16
	YCoordinate  zcl.Zs16
	ZCoordinate  zcl.Zs16
	Power        zcl.Zs16
	// Rate at which the signal power decays with increasing distance
	PathLossExponent zcl.Zu16
	LocationMethod   zcl.Zenum8
	QualityMeasure   zcl.Zu8
	LocationAge      zcl.Zu16
}

const LocationDataNotificationCommand zcl.CommandID = 2

func (v *LocationDataNotification) Values() []zcl.Val {
	return []zcl.Val{
		&v.LocationType,
		&v.XCoordinate,
		&v.YCoordinate,
		&v.ZCoordinate,
		&v.Power,
		&v.PathLossExponent,
		&v.LocationMethod,
		&v.QualityMeasure,
		&v.LocationAge,
	}
}

func (v LocationDataNotification) ID() zcl.CommandID {
	return LocationDataNotificationCommand
}

func (v LocationDataNotification) Cluster() zcl.ClusterID {
	return LocationID
}

func (v LocationDataNotification) MnfCode() []byte {
	return []byte{}
}

func (v LocationDataNotification) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.LocationType.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.XCoordinate.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.YCoordinate.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if (v.LocationType & 0x02) == 0x00 {
		if tmp, err = v.ZCoordinate.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}

	if tmp, err = v.Power.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.PathLossExponent.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if (v.LocationType & 0x01) == 0x00 {
		if tmp, err = v.LocationMethod.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}

	if (v.LocationType & 0x01) == 0x00 {
		if tmp, err = v.QualityMeasure.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}

	if (v.LocationType & 0x01) == 0x00 {
		if tmp, err = v.LocationAge.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}

	return data, nil
}

func (v *LocationDataNotification) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.LocationType).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.XCoordinate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.YCoordinate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if (v.LocationType & 0x02) == 0x00 {
		if b, err = (&v.ZCoordinate).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	if b, err = (&v.Power).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.PathLossExponent).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if (v.LocationType & 0x01) == 0x00 {
		if b, err = (&v.LocationMethod).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	if (v.LocationType & 0x01) == 0x00 {
		if b, err = (&v.QualityMeasure).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	if (v.LocationType & 0x01) == 0x00 {
		if b, err = (&v.LocationAge).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	return b, nil
}

type CompactLocationDataNotification struct {
	LocationType   zcl.Zenum8
	XCoordinate    zcl.Zs16
	YCoordinate    zcl.Zs16
	ZCoordinate    zcl.Zs16
	QualityMeasure zcl.Zu8
	LocationAge    zcl.Zu16
}

const CompactLocationDataNotificationCommand zcl.CommandID = 3

func (v *CompactLocationDataNotification) Values() []zcl.Val {
	return []zcl.Val{
		&v.LocationType,
		&v.XCoordinate,
		&v.YCoordinate,
		&v.ZCoordinate,
		&v.QualityMeasure,
		&v.LocationAge,
	}
}

func (v CompactLocationDataNotification) ID() zcl.CommandID {
	return CompactLocationDataNotificationCommand
}

func (v CompactLocationDataNotification) Cluster() zcl.ClusterID {
	return LocationID
}

func (v CompactLocationDataNotification) MnfCode() []byte {
	return []byte{}
}

func (v CompactLocationDataNotification) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.LocationType.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.XCoordinate.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.YCoordinate.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if (v.LocationType & 0x02) == 0x00 {
		if tmp, err = v.ZCoordinate.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}

	if (v.LocationType & 0x01) == 0x00 {
		if tmp, err = v.QualityMeasure.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}

	if (v.LocationType & 0x01) == 0x00 {
		if tmp, err = v.LocationAge.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}

	return data, nil
}

func (v *CompactLocationDataNotification) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.LocationType).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.XCoordinate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.YCoordinate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if (v.LocationType & 0x02) == 0x00 {
		if b, err = (&v.ZCoordinate).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	if (v.LocationType & 0x01) == 0x00 {
		if b, err = (&v.QualityMeasure).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	if (v.LocationType & 0x01) == 0x00 {
		if b, err = (&v.LocationAge).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	return b, nil
}

type RssiPing struct {
	LocationType zcl.Zenum8
}

const RssiPingCommand zcl.CommandID = 4

func (v *RssiPing) Values() []zcl.Val {
	return []zcl.Val{
		&v.LocationType,
	}
}

func (v RssiPing) ID() zcl.CommandID {
	return RssiPingCommand
}

func (v RssiPing) Cluster() zcl.ClusterID {
	return LocationID
}

func (v RssiPing) MnfCode() []byte {
	return []byte{}
}

func (v RssiPing) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.LocationType.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *RssiPing) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.LocationType).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type RssiRequest struct {
}

const RssiRequestCommand zcl.CommandID = 5

func (v *RssiRequest) Values() []zcl.Val {
	return []zcl.Val{}
}

func (v RssiRequest) ID() zcl.CommandID {
	return RssiRequestCommand
}

func (v RssiRequest) Cluster() zcl.ClusterID {
	return LocationID
}

func (v RssiRequest) MnfCode() []byte {
	return []byte{}
}

func (v RssiRequest) MarshalZcl() ([]byte, error) {
	return nil, nil
}

func (v *RssiRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

type ReportRssiMeasurements struct {
	MeasuringDevice zcl.Zuid
	NeighborsInfo   zcl.Zset
}

const ReportRssiMeasurementsCommand zcl.CommandID = 6

func (v *ReportRssiMeasurements) Values() []zcl.Val {
	return []zcl.Val{
		&v.MeasuringDevice,
		&v.NeighborsInfo,
	}
}

func (v ReportRssiMeasurements) ID() zcl.CommandID {
	return ReportRssiMeasurementsCommand
}

func (v ReportRssiMeasurements) Cluster() zcl.ClusterID {
	return LocationID
}

func (v ReportRssiMeasurements) MnfCode() []byte {
	return []byte{}
}

func (v ReportRssiMeasurements) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.MeasuringDevice.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.NeighborsInfo.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *ReportRssiMeasurements) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.MeasuringDevice).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.NeighborsInfo).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type RequestOwnLocation struct {
	BlindNodeAddress zcl.Zuid
}

const RequestOwnLocationCommand zcl.CommandID = 7

func (v *RequestOwnLocation) Values() []zcl.Val {
	return []zcl.Val{
		&v.BlindNodeAddress,
	}
}

func (v RequestOwnLocation) ID() zcl.CommandID {
	return RequestOwnLocationCommand
}

func (v RequestOwnLocation) Cluster() zcl.ClusterID {
	return LocationID
}

func (v RequestOwnLocation) MnfCode() []byte {
	return []byte{}
}

func (v RequestOwnLocation) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.BlindNodeAddress.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *RequestOwnLocation) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.BlindNodeAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// Returns the result of a distance measure.
type DistanceMeasureResponse struct {
	TargetAddress zcl.Zu16
	DistanceMeter zcl.Zu16
	QualityIndex  zcl.Zu16
}

const DistanceMeasureResponseCommand zcl.CommandID = 64

func (v *DistanceMeasureResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.TargetAddress,
		&v.DistanceMeter,
		&v.QualityIndex,
	}
}

func (v DistanceMeasureResponse) ID() zcl.CommandID {
	return DistanceMeasureResponseCommand
}

func (v DistanceMeasureResponse) Cluster() zcl.ClusterID {
	return LocationID
}

func (v DistanceMeasureResponse) MnfCode() []byte {
	return []byte{}
}

func (v DistanceMeasureResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.TargetAddress.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.DistanceMeter.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.QualityIndex.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *DistanceMeasureResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.TargetAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.DistanceMeter).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.QualityIndex).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

const LocationTypeAttr zcl.AttrID = 0

type LocationType zcl.Zbmp8

func (a LocationType) ID() zcl.AttrID         { return LocationTypeAttr }
func (a LocationType) Cluster() zcl.ClusterID { return LocationID }
func (a *LocationType) Value() *LocationType  { return a }
func (a LocationType) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *LocationType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = LocationType(*nt)
	return br, err
}

func (a LocationType) Readable() bool   { return true }
func (a LocationType) Writable() bool   { return true }
func (a LocationType) Reportable() bool { return false }
func (a LocationType) SceneIndex() int  { return -1 }

func (a LocationType) String() string {
	var bstr []string
	if a.IsAbsoluteLocation() {
		bstr = append(bstr, "Absolute location")
	}
	if a.IsTwoDimensional() {
		bstr = append(bstr, "Two dimensional")
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a LocationType) IsAbsoluteLocation() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *LocationType) SetAbsoluteLocation(b bool) {
	*a = LocationType(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a LocationType) IsTwoDimensional() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *LocationType) SetTwoDimensional(b bool) {
	*a = LocationType(zcl.BitmapSet([]byte(*a), 1, b))
}

const LocationMethodAttr zcl.AttrID = 1

type LocationMethod zcl.Zenum8

func (a LocationMethod) ID() zcl.AttrID          { return LocationMethodAttr }
func (a LocationMethod) Cluster() zcl.ClusterID  { return LocationID }
func (a *LocationMethod) Value() *LocationMethod { return a }
func (a LocationMethod) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *LocationMethod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = LocationMethod(*nt)
	return br, err
}

func (a LocationMethod) Readable() bool   { return true }
func (a LocationMethod) Writable() bool   { return true }
func (a LocationMethod) Reportable() bool { return false }
func (a LocationMethod) SceneIndex() int  { return -1 }

func (a LocationMethod) String() string {
	switch a {
	case 0x00:
		return "Lateration"
	case 0x01:
		return "Signposting"
	case 0x02:
		return "RF fingerprinting"
	case 0x03:
		return "Out of band"
	case 0x04:
		return "Centralized"
	}
	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsLateration checks if LocationMethod equals the value for Lateration (0x00)
func (a LocationMethod) IsLateration() bool { return a == 0x00 }

// SetLateration sets LocationMethod to Lateration (0x00)
func (a *LocationMethod) SetLateration() { *a = 0x00 }

// IsSignposting checks if LocationMethod equals the value for Signposting (0x01)
func (a LocationMethod) IsSignposting() bool { return a == 0x01 }

// SetSignposting sets LocationMethod to Signposting (0x01)
func (a *LocationMethod) SetSignposting() { *a = 0x01 }

// IsRfFingerprinting checks if LocationMethod equals the value for RF fingerprinting (0x02)
func (a LocationMethod) IsRfFingerprinting() bool { return a == 0x02 }

// SetRfFingerprinting sets LocationMethod to RF fingerprinting (0x02)
func (a *LocationMethod) SetRfFingerprinting() { *a = 0x02 }

// IsOutOfBand checks if LocationMethod equals the value for Out of band (0x03)
func (a LocationMethod) IsOutOfBand() bool { return a == 0x03 }

// SetOutOfBand sets LocationMethod to Out of band (0x03)
func (a *LocationMethod) SetOutOfBand() { *a = 0x03 }

// IsCentralized checks if LocationMethod equals the value for Centralized (0x04)
func (a LocationMethod) IsCentralized() bool { return a == 0x04 }

// SetCentralized sets LocationMethod to Centralized (0x04)
func (a *LocationMethod) SetCentralized() { *a = 0x04 }

const LocationAgeAttr zcl.AttrID = 2

type LocationAge zcl.Zu16

func (a LocationAge) ID() zcl.AttrID         { return LocationAgeAttr }
func (a LocationAge) Cluster() zcl.ClusterID { return LocationID }
func (a *LocationAge) Value() *LocationAge   { return a }
func (a LocationAge) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *LocationAge) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = LocationAge(*nt)
	return br, err
}

func (a LocationAge) Readable() bool   { return true }
func (a LocationAge) Writable() bool   { return false }
func (a LocationAge) Reportable() bool { return false }
func (a LocationAge) SceneIndex() int  { return -1 }

func (a LocationAge) String() string {
	return zcl.Duration(int(a), 0).String()
}

const QualityMeasureAttr zcl.AttrID = 3

type QualityMeasure zcl.Zu8

func (a QualityMeasure) ID() zcl.AttrID          { return QualityMeasureAttr }
func (a QualityMeasure) Cluster() zcl.ClusterID  { return LocationID }
func (a *QualityMeasure) Value() *QualityMeasure { return a }
func (a QualityMeasure) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *QualityMeasure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = QualityMeasure(*nt)
	return br, err
}

func (a QualityMeasure) Readable() bool   { return true }
func (a QualityMeasure) Writable() bool   { return false }
func (a QualityMeasure) Reportable() bool { return false }
func (a QualityMeasure) SceneIndex() int  { return -1 }

func (a QualityMeasure) String() string {
	return zcl.Sprintf("%f%s", float64(a), "%")
}

const NumberOfDevicesAttr zcl.AttrID = 4

type NumberOfDevices zcl.Zu8

func (a NumberOfDevices) ID() zcl.AttrID           { return NumberOfDevicesAttr }
func (a NumberOfDevices) Cluster() zcl.ClusterID   { return LocationID }
func (a *NumberOfDevices) Value() *NumberOfDevices { return a }
func (a NumberOfDevices) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *NumberOfDevices) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = NumberOfDevices(*nt)
	return br, err
}

func (a NumberOfDevices) Readable() bool   { return true }
func (a NumberOfDevices) Writable() bool   { return false }
func (a NumberOfDevices) Reportable() bool { return false }
func (a NumberOfDevices) SceneIndex() int  { return -1 }

func (a NumberOfDevices) String() string {
	return zcl.Sprintf("%s", zcl.Zu8(a))
}

const XCoordinateAttr zcl.AttrID = 16

type XCoordinate zcl.Zs16

func (a XCoordinate) ID() zcl.AttrID         { return XCoordinateAttr }
func (a XCoordinate) Cluster() zcl.ClusterID { return LocationID }
func (a *XCoordinate) Value() *XCoordinate   { return a }
func (a XCoordinate) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *XCoordinate) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = XCoordinate(*nt)
	return br, err
}

func (a XCoordinate) Readable() bool   { return true }
func (a XCoordinate) Writable() bool   { return true }
func (a XCoordinate) Reportable() bool { return false }
func (a XCoordinate) SceneIndex() int  { return -1 }

func (a XCoordinate) String() string {
	return zcl.Sprintf("%f%s", float64(a)/10, "m")
}

const YCoordinateAttr zcl.AttrID = 17

type YCoordinate zcl.Zs16

func (a YCoordinate) ID() zcl.AttrID         { return YCoordinateAttr }
func (a YCoordinate) Cluster() zcl.ClusterID { return LocationID }
func (a *YCoordinate) Value() *YCoordinate   { return a }
func (a YCoordinate) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *YCoordinate) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = YCoordinate(*nt)
	return br, err
}

func (a YCoordinate) Readable() bool   { return true }
func (a YCoordinate) Writable() bool   { return true }
func (a YCoordinate) Reportable() bool { return false }
func (a YCoordinate) SceneIndex() int  { return -1 }

func (a YCoordinate) String() string {
	return zcl.Sprintf("%f%s", float64(a)/10, "m")
}

const ZCoordinateAttr zcl.AttrID = 18

type ZCoordinate zcl.Zs16

func (a ZCoordinate) ID() zcl.AttrID         { return ZCoordinateAttr }
func (a ZCoordinate) Cluster() zcl.ClusterID { return LocationID }
func (a *ZCoordinate) Value() *ZCoordinate   { return a }
func (a ZCoordinate) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ZCoordinate) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ZCoordinate(*nt)
	return br, err
}

func (a ZCoordinate) Readable() bool   { return true }
func (a ZCoordinate) Writable() bool   { return true }
func (a ZCoordinate) Reportable() bool { return false }
func (a ZCoordinate) SceneIndex() int  { return -1 }

func (a ZCoordinate) String() string {
	return zcl.Sprintf("%f%s", float64(a)/10, "m")
}

const PowerAttr zcl.AttrID = 19

type Power zcl.Zs16

func (a Power) ID() zcl.AttrID         { return PowerAttr }
func (a Power) Cluster() zcl.ClusterID { return LocationID }
func (a *Power) Value() *Power         { return a }
func (a Power) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *Power) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = Power(*nt)
	return br, err
}

func (a Power) Readable() bool   { return true }
func (a Power) Writable() bool   { return true }
func (a Power) Reportable() bool { return false }
func (a Power) SceneIndex() int  { return -1 }

func (a Power) String() string {
	return zcl.Sprintf("%f%s", float64(a)/100, "dBm")
}

const PathLossExponentAttr zcl.AttrID = 20

type PathLossExponent zcl.Zu16

func (a PathLossExponent) ID() zcl.AttrID            { return PathLossExponentAttr }
func (a PathLossExponent) Cluster() zcl.ClusterID    { return LocationID }
func (a *PathLossExponent) Value() *PathLossExponent { return a }
func (a PathLossExponent) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *PathLossExponent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = PathLossExponent(*nt)
	return br, err
}

func (a PathLossExponent) Readable() bool   { return true }
func (a PathLossExponent) Writable() bool   { return true }
func (a PathLossExponent) Reportable() bool { return false }
func (a PathLossExponent) SceneIndex() int  { return -1 }

func (a PathLossExponent) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ReportingPeriodAttr zcl.AttrID = 21

type ReportingPeriod zcl.Zu16

func (a ReportingPeriod) ID() zcl.AttrID           { return ReportingPeriodAttr }
func (a ReportingPeriod) Cluster() zcl.ClusterID   { return LocationID }
func (a *ReportingPeriod) Value() *ReportingPeriod { return a }
func (a ReportingPeriod) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ReportingPeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ReportingPeriod(*nt)
	return br, err
}

func (a ReportingPeriod) Readable() bool   { return true }
func (a ReportingPeriod) Writable() bool   { return true }
func (a ReportingPeriod) Reportable() bool { return false }
func (a ReportingPeriod) SceneIndex() int  { return -1 }

func (a ReportingPeriod) String() string {
	return zcl.Duration(int(a), 0).String()
}

const CalculationPeriodAttr zcl.AttrID = 22

type CalculationPeriod zcl.Zu16

func (a CalculationPeriod) ID() zcl.AttrID             { return CalculationPeriodAttr }
func (a CalculationPeriod) Cluster() zcl.ClusterID     { return LocationID }
func (a *CalculationPeriod) Value() *CalculationPeriod { return a }
func (a CalculationPeriod) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *CalculationPeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = CalculationPeriod(*nt)
	return br, err
}

func (a CalculationPeriod) Readable() bool   { return true }
func (a CalculationPeriod) Writable() bool   { return true }
func (a CalculationPeriod) Reportable() bool { return false }
func (a CalculationPeriod) SceneIndex() int  { return -1 }

func (a CalculationPeriod) String() string {
	return zcl.Duration(int(a), 0).String()
}

const NumberRssiMeasurementsAttr zcl.AttrID = 23

type NumberRssiMeasurements zcl.Zu8

func (a NumberRssiMeasurements) ID() zcl.AttrID                  { return NumberRssiMeasurementsAttr }
func (a NumberRssiMeasurements) Cluster() zcl.ClusterID          { return LocationID }
func (a *NumberRssiMeasurements) Value() *NumberRssiMeasurements { return a }
func (a NumberRssiMeasurements) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *NumberRssiMeasurements) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = NumberRssiMeasurements(*nt)
	return br, err
}

func (a NumberRssiMeasurements) Readable() bool   { return true }
func (a NumberRssiMeasurements) Writable() bool   { return true }
func (a NumberRssiMeasurements) Reportable() bool { return false }
func (a NumberRssiMeasurements) SceneIndex() int  { return -1 }

func (a NumberRssiMeasurements) String() string {
	return zcl.Sprintf("%s", zcl.Zu8(a))
}
