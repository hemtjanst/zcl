// Measure distance between devices.
package general

import (
	"neotor.se/zcl/cluster/zcl"
)

// Location
// Measure distance between devices.

func NewLocationServer(profile zcl.ProfileID) *LocationServer { return &LocationServer{p: profile} }
func NewLocationClient(profile zcl.ProfileID) *LocationClient { return &LocationClient{p: profile} }

const LocationCluster zcl.ClusterID = 11

type LocationServer struct {
	p zcl.ProfileID

	LocationType           *LocationType
	LocationMethod         *LocationMethod
	LocationAge            *LocationAge
	QualityMeasure         *QualityMeasure
	NumberOfDevices        *NumberOfDevices
	XCoordinate            *XCoordinate
	YCoordinate            *YCoordinate
	ZCoordinate            *ZCoordinate
	Power                  *Power
	PathLossExponent       *PathLossExponent
	ReportingPeriod        *ReportingPeriod
	CalculationPeriod      *CalculationPeriod
	NumberRssiMeasurements *NumberRssiMeasurements
}

func (s *LocationServer) SetAbsoluteLocation() *SetAbsoluteLocation { return new(SetAbsoluteLocation) }
func (s *LocationServer) SetDeviceConfiguration() *SetDeviceConfiguration {
	return new(SetDeviceConfiguration)
}
func (s *LocationServer) GetDeviceConfiguration() *GetDeviceConfiguration {
	return new(GetDeviceConfiguration)
}
func (s *LocationServer) GetLocationData() *GetLocationData       { return new(GetLocationData) }
func (s *LocationServer) RssiResponse() *RssiResponse             { return new(RssiResponse) }
func (s *LocationServer) SendPings() *SendPings                   { return new(SendPings) }
func (s *LocationServer) AnchorNodeAnnounce() *AnchorNodeAnnounce { return new(AnchorNodeAnnounce) }
func (s *LocationServer) DistanceMeasure() *DistanceMeasure       { return new(DistanceMeasure) }

type LocationClient struct {
	p zcl.ProfileID
}

func (s *LocationClient) DeviceConfigurationResponse() *DeviceConfigurationResponse {
	return new(DeviceConfigurationResponse)
}
func (s *LocationClient) LocationDataResponse() *LocationDataResponse {
	return new(LocationDataResponse)
}
func (s *LocationClient) LocationDataNotification() *LocationDataNotification {
	return new(LocationDataNotification)
}
func (s *LocationClient) CompactLocationDataNotification() *CompactLocationDataNotification {
	return new(CompactLocationDataNotification)
}
func (s *LocationClient) RssiPing() *RssiPing       { return new(RssiPing) }
func (s *LocationClient) RssiRequest() *RssiRequest { return new(RssiRequest) }
func (s *LocationClient) ReportRssiMeasurements() *ReportRssiMeasurements {
	return new(ReportRssiMeasurements)
}
func (s *LocationClient) RequestOwnLocation() *RequestOwnLocation { return new(RequestOwnLocation) }
func (s *LocationClient) DistanceMeasureResponse() *DistanceMeasureResponse {
	return new(DistanceMeasureResponse)
}

/*
var LocationServer = map[zcl.CommandID]func() zcl.Command{
    SetAbsoluteLocationID: func() zcl.Command { return new(SetAbsoluteLocation) },
    SetDeviceConfigurationID: func() zcl.Command { return new(SetDeviceConfiguration) },
    GetDeviceConfigurationID: func() zcl.Command { return new(GetDeviceConfiguration) },
    GetLocationDataID: func() zcl.Command { return new(GetLocationData) },
    RssiResponseID: func() zcl.Command { return new(RssiResponse) },
    SendPingsID: func() zcl.Command { return new(SendPings) },
    AnchorNodeAnnounceID: func() zcl.Command { return new(AnchorNodeAnnounce) },
    DistanceMeasureID: func() zcl.Command { return new(DistanceMeasure) },
}

var LocationClient = map[zcl.CommandID]func() zcl.Command{
    DeviceConfigurationResponseID: func() zcl.Command { return new(DeviceConfigurationResponse) },
    LocationDataResponseID: func() zcl.Command { return new(LocationDataResponse) },
    LocationDataNotificationID: func() zcl.Command { return new(LocationDataNotification) },
    CompactLocationDataNotificationID: func() zcl.Command { return new(CompactLocationDataNotification) },
    RssiPingID: func() zcl.Command { return new(RssiPing) },
    RssiRequestID: func() zcl.Command { return new(RssiRequest) },
    ReportRssiMeasurementsID: func() zcl.Command { return new(ReportRssiMeasurements) },
    RequestOwnLocationID: func() zcl.Command { return new(RequestOwnLocation) },
    DistanceMeasureResponseID: func() zcl.Command { return new(DistanceMeasureResponse) },
}
*/

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
	return LocationCluster
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
	return LocationCluster
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
	return LocationCluster
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
	return LocationCluster
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
	return LocationCluster
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
	return LocationCluster
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
	return LocationCluster
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
	return LocationCluster
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
	return LocationCluster
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
	return LocationCluster
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
	return LocationCluster
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
	return LocationCluster
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
	return LocationCluster
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
	return LocationCluster
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
	return LocationCluster
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
	return LocationCluster
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
	return LocationCluster
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

type LocationType zcl.Zbmp8

func (a LocationType) ID() zcl.AttrID         { return 0 }
func (a LocationType) Cluster() zcl.ClusterID { return LocationCluster }
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

type LocationMethod zcl.Zenum8

func (a LocationMethod) ID() zcl.AttrID          { return 1 }
func (a LocationMethod) Cluster() zcl.ClusterID  { return LocationCluster }
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

type LocationAge zcl.Zu16

func (a LocationAge) ID() zcl.AttrID         { return 2 }
func (a LocationAge) Cluster() zcl.ClusterID { return LocationCluster }
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

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type QualityMeasure zcl.Zu8

func (a QualityMeasure) ID() zcl.AttrID          { return 3 }
func (a QualityMeasure) Cluster() zcl.ClusterID  { return LocationCluster }
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

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type NumberOfDevices zcl.Zu8

func (a NumberOfDevices) ID() zcl.AttrID           { return 4 }
func (a NumberOfDevices) Cluster() zcl.ClusterID   { return LocationCluster }
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

type XCoordinate zcl.Zs16

func (a XCoordinate) ID() zcl.AttrID         { return 16 }
func (a XCoordinate) Cluster() zcl.ClusterID { return LocationCluster }
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

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type YCoordinate zcl.Zs16

func (a YCoordinate) ID() zcl.AttrID         { return 17 }
func (a YCoordinate) Cluster() zcl.ClusterID { return LocationCluster }
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

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type ZCoordinate zcl.Zs16

func (a ZCoordinate) ID() zcl.AttrID         { return 18 }
func (a ZCoordinate) Cluster() zcl.ClusterID { return LocationCluster }
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

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type Power zcl.Zs16

func (a Power) ID() zcl.AttrID         { return 19 }
func (a Power) Cluster() zcl.ClusterID { return LocationCluster }
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

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type PathLossExponent zcl.Zu16

func (a PathLossExponent) ID() zcl.AttrID            { return 20 }
func (a PathLossExponent) Cluster() zcl.ClusterID    { return LocationCluster }
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

type ReportingPeriod zcl.Zu16

func (a ReportingPeriod) ID() zcl.AttrID           { return 21 }
func (a ReportingPeriod) Cluster() zcl.ClusterID   { return LocationCluster }
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

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type CalculationPeriod zcl.Zu16

func (a CalculationPeriod) ID() zcl.AttrID             { return 22 }
func (a CalculationPeriod) Cluster() zcl.ClusterID     { return LocationCluster }
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

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type NumberRssiMeasurements zcl.Zu8

func (a NumberRssiMeasurements) ID() zcl.AttrID                  { return 23 }
func (a NumberRssiMeasurements) Cluster() zcl.ClusterID          { return LocationCluster }
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
