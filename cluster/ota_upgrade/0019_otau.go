// Over the air upgrade.
package ota_upgrade

import (
	"neotor.se/zcl"
)

// Otau
const OtauID zcl.ClusterID = 25

var OtauCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{
		QueryNextImageCommand:     func() zcl.Command { return new(QueryNextImage) },
		UpgradeEndResponseCommand: func() zcl.Command { return new(UpgradeEndResponse) },
	},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{
		UpgradeServerAttr:                func() zcl.Attr { return new(UpgradeServer) },
		FileOffsetAttr:                   func() zcl.Attr { return new(FileOffset) },
		CurrentFileVersionAttr:           func() zcl.Attr { return new(CurrentFileVersion) },
		CurrentZigbeeStackVersionAttr:    func() zcl.Attr { return new(CurrentZigbeeStackVersion) },
		DownloadedFileVersionAttr:        func() zcl.Attr { return new(DownloadedFileVersion) },
		DownloadedZigbeeStackVersionAttr: func() zcl.Attr { return new(DownloadedZigbeeStackVersion) },
		ImageUpgradeStatusAttr:           func() zcl.Attr { return new(ImageUpgradeStatus) },
		MinBlockRequestDelayAttr:         func() zcl.Attr { return new(MinBlockRequestDelay) },
	},
	SceneAttr: []zcl.AttrID{},
}

type QueryNextImage struct {
	ControlField       zcl.Zbmp8
	ManufacturerId     zcl.Zu16
	ImageType          zcl.Zenum16
	ApplicationRelease zcl.Zu8
	ApplicationBuild   zcl.Zu8
	StackRelease       zcl.Zu8
	StackBuild         zcl.Zu8
}

const QueryNextImageCommand zcl.CommandID = 1

func (v *QueryNextImage) Values() []zcl.Val {
	return []zcl.Val{
		&v.ControlField,
		&v.ManufacturerId,
		&v.ImageType,
		&v.ApplicationRelease,
		&v.ApplicationBuild,
		&v.StackRelease,
		&v.StackBuild,
	}
}

func (v QueryNextImage) ID() zcl.CommandID {
	return QueryNextImageCommand
}

func (v QueryNextImage) Cluster() zcl.ClusterID {
	return OtauID
}

func (v QueryNextImage) MnfCode() []byte {
	return []byte{}
}

func (v QueryNextImage) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.ControlField.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ManufacturerId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ImageType.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ApplicationRelease.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ApplicationBuild.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.StackRelease.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.StackBuild.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *QueryNextImage) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.ControlField).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ManufacturerId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ImageType).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ApplicationRelease).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ApplicationBuild).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.StackRelease).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.StackBuild).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type UpgradeEndResponse struct {
	ManufacturerId zcl.Zu16
	ImageType      zcl.Zu16
	FileVersion    zcl.Zu32
	CurrentTime    zcl.Zu32
	UpgradeTime    zcl.Zu32
}

const UpgradeEndResponseCommand zcl.CommandID = 7

func (v *UpgradeEndResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.ManufacturerId,
		&v.ImageType,
		&v.FileVersion,
		&v.CurrentTime,
		&v.UpgradeTime,
	}
}

func (v UpgradeEndResponse) ID() zcl.CommandID {
	return UpgradeEndResponseCommand
}

func (v UpgradeEndResponse) Cluster() zcl.ClusterID {
	return OtauID
}

func (v UpgradeEndResponse) MnfCode() []byte {
	return []byte{}
}

func (v UpgradeEndResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.ManufacturerId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ImageType.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.FileVersion.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.CurrentTime.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.UpgradeTime.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *UpgradeEndResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.ManufacturerId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ImageType).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.FileVersion).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.CurrentTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.UpgradeTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

const UpgradeServerAttr zcl.AttrID = 0

type UpgradeServer zcl.Zuid

func (a UpgradeServer) ID() zcl.AttrID         { return UpgradeServerAttr }
func (a UpgradeServer) Cluster() zcl.ClusterID { return OtauID }
func (a *UpgradeServer) Value() *UpgradeServer { return a }
func (a UpgradeServer) MarshalZcl() ([]byte, error) {
	return zcl.Zuid(a).MarshalZcl()
}
func (a *UpgradeServer) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zuid)
	br, err := nt.UnmarshalZcl(b)
	*a = UpgradeServer(*nt)
	return br, err
}

func (a UpgradeServer) Readable() bool   { return true }
func (a UpgradeServer) Writable() bool   { return true }
func (a UpgradeServer) Reportable() bool { return false }
func (a UpgradeServer) SceneIndex() int  { return -1 }

func (a UpgradeServer) String() string {
	return zcl.Sprintf("%s", zcl.Zuid(a))
}

const FileOffsetAttr zcl.AttrID = 1

type FileOffset zcl.Zu32

func (a FileOffset) ID() zcl.AttrID         { return FileOffsetAttr }
func (a FileOffset) Cluster() zcl.ClusterID { return OtauID }
func (a *FileOffset) Value() *FileOffset    { return a }
func (a FileOffset) MarshalZcl() ([]byte, error) {
	return zcl.Zu32(a).MarshalZcl()
}
func (a *FileOffset) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = FileOffset(*nt)
	return br, err
}

func (a FileOffset) Readable() bool   { return true }
func (a FileOffset) Writable() bool   { return false }
func (a FileOffset) Reportable() bool { return false }
func (a FileOffset) SceneIndex() int  { return -1 }

func (a FileOffset) String() string {
	return zcl.Sprintf("0x%X", zcl.Zu32(a))
}

const CurrentFileVersionAttr zcl.AttrID = 2

type CurrentFileVersion zcl.Zu32

func (a CurrentFileVersion) ID() zcl.AttrID              { return CurrentFileVersionAttr }
func (a CurrentFileVersion) Cluster() zcl.ClusterID      { return OtauID }
func (a *CurrentFileVersion) Value() *CurrentFileVersion { return a }
func (a CurrentFileVersion) MarshalZcl() ([]byte, error) {
	return zcl.Zu32(a).MarshalZcl()
}
func (a *CurrentFileVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentFileVersion(*nt)
	return br, err
}

func (a CurrentFileVersion) Readable() bool   { return true }
func (a CurrentFileVersion) Writable() bool   { return false }
func (a CurrentFileVersion) Reportable() bool { return false }
func (a CurrentFileVersion) SceneIndex() int  { return -1 }

func (a CurrentFileVersion) String() string {
	return zcl.Sprintf("0x%X", zcl.Zu32(a))
}

const CurrentZigbeeStackVersionAttr zcl.AttrID = 3

type CurrentZigbeeStackVersion zcl.Zu16

func (a CurrentZigbeeStackVersion) ID() zcl.AttrID                     { return CurrentZigbeeStackVersionAttr }
func (a CurrentZigbeeStackVersion) Cluster() zcl.ClusterID             { return OtauID }
func (a *CurrentZigbeeStackVersion) Value() *CurrentZigbeeStackVersion { return a }
func (a CurrentZigbeeStackVersion) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *CurrentZigbeeStackVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentZigbeeStackVersion(*nt)
	return br, err
}

func (a CurrentZigbeeStackVersion) Readable() bool   { return true }
func (a CurrentZigbeeStackVersion) Writable() bool   { return false }
func (a CurrentZigbeeStackVersion) Reportable() bool { return false }
func (a CurrentZigbeeStackVersion) SceneIndex() int  { return -1 }

func (a CurrentZigbeeStackVersion) String() string {
	return zcl.Sprintf("0x%X", zcl.Zu16(a))
}

const DownloadedFileVersionAttr zcl.AttrID = 4

type DownloadedFileVersion zcl.Zu32

func (a DownloadedFileVersion) ID() zcl.AttrID                 { return DownloadedFileVersionAttr }
func (a DownloadedFileVersion) Cluster() zcl.ClusterID         { return OtauID }
func (a *DownloadedFileVersion) Value() *DownloadedFileVersion { return a }
func (a DownloadedFileVersion) MarshalZcl() ([]byte, error) {
	return zcl.Zu32(a).MarshalZcl()
}
func (a *DownloadedFileVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = DownloadedFileVersion(*nt)
	return br, err
}

func (a DownloadedFileVersion) Readable() bool   { return true }
func (a DownloadedFileVersion) Writable() bool   { return false }
func (a DownloadedFileVersion) Reportable() bool { return false }
func (a DownloadedFileVersion) SceneIndex() int  { return -1 }

func (a DownloadedFileVersion) String() string {
	return zcl.Sprintf("0x%X", zcl.Zu32(a))
}

const DownloadedZigbeeStackVersionAttr zcl.AttrID = 5

type DownloadedZigbeeStackVersion zcl.Zu16

func (a DownloadedZigbeeStackVersion) ID() zcl.AttrID                        { return DownloadedZigbeeStackVersionAttr }
func (a DownloadedZigbeeStackVersion) Cluster() zcl.ClusterID                { return OtauID }
func (a *DownloadedZigbeeStackVersion) Value() *DownloadedZigbeeStackVersion { return a }
func (a DownloadedZigbeeStackVersion) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *DownloadedZigbeeStackVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = DownloadedZigbeeStackVersion(*nt)
	return br, err
}

func (a DownloadedZigbeeStackVersion) Readable() bool   { return true }
func (a DownloadedZigbeeStackVersion) Writable() bool   { return false }
func (a DownloadedZigbeeStackVersion) Reportable() bool { return false }
func (a DownloadedZigbeeStackVersion) SceneIndex() int  { return -1 }

func (a DownloadedZigbeeStackVersion) String() string {
	return zcl.Sprintf("0x%X", zcl.Zu16(a))
}

const ImageUpgradeStatusAttr zcl.AttrID = 6

type ImageUpgradeStatus zcl.Zenum8

func (a ImageUpgradeStatus) ID() zcl.AttrID              { return ImageUpgradeStatusAttr }
func (a ImageUpgradeStatus) Cluster() zcl.ClusterID      { return OtauID }
func (a *ImageUpgradeStatus) Value() *ImageUpgradeStatus { return a }
func (a ImageUpgradeStatus) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *ImageUpgradeStatus) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = ImageUpgradeStatus(*nt)
	return br, err
}

func (a ImageUpgradeStatus) Readable() bool   { return true }
func (a ImageUpgradeStatus) Writable() bool   { return true }
func (a ImageUpgradeStatus) Reportable() bool { return false }
func (a ImageUpgradeStatus) SceneIndex() int  { return -1 }

func (a ImageUpgradeStatus) String() string {
	switch a {
	case 0x00:
		return "Normal"
	case 0x01:
		return "Download in progress"
	case 0x02:
		return "Download complete"
	case 0x03:
		return "Waiting to upgrade"
	case 0x04:
		return "Count down"
	case 0x05:
		return "Wait for more"
	}
	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsNormal checks if ImageUpgradeStatus equals the value for Normal (0x00)
func (a ImageUpgradeStatus) IsNormal() bool { return a == 0x00 }

// SetNormal sets ImageUpgradeStatus to Normal (0x00)
func (a *ImageUpgradeStatus) SetNormal() { *a = 0x00 }

// IsDownloadInProgress checks if ImageUpgradeStatus equals the value for Download in progress (0x01)
func (a ImageUpgradeStatus) IsDownloadInProgress() bool { return a == 0x01 }

// SetDownloadInProgress sets ImageUpgradeStatus to Download in progress (0x01)
func (a *ImageUpgradeStatus) SetDownloadInProgress() { *a = 0x01 }

// IsDownloadComplete checks if ImageUpgradeStatus equals the value for Download complete (0x02)
func (a ImageUpgradeStatus) IsDownloadComplete() bool { return a == 0x02 }

// SetDownloadComplete sets ImageUpgradeStatus to Download complete (0x02)
func (a *ImageUpgradeStatus) SetDownloadComplete() { *a = 0x02 }

// IsWaitingToUpgrade checks if ImageUpgradeStatus equals the value for Waiting to upgrade (0x03)
func (a ImageUpgradeStatus) IsWaitingToUpgrade() bool { return a == 0x03 }

// SetWaitingToUpgrade sets ImageUpgradeStatus to Waiting to upgrade (0x03)
func (a *ImageUpgradeStatus) SetWaitingToUpgrade() { *a = 0x03 }

// IsCountDown checks if ImageUpgradeStatus equals the value for Count down (0x04)
func (a ImageUpgradeStatus) IsCountDown() bool { return a == 0x04 }

// SetCountDown sets ImageUpgradeStatus to Count down (0x04)
func (a *ImageUpgradeStatus) SetCountDown() { *a = 0x04 }

// IsWaitForMore checks if ImageUpgradeStatus equals the value for Wait for more (0x05)
func (a ImageUpgradeStatus) IsWaitForMore() bool { return a == 0x05 }

// SetWaitForMore sets ImageUpgradeStatus to Wait for more (0x05)
func (a *ImageUpgradeStatus) SetWaitForMore() { *a = 0x05 }

const MinBlockRequestDelayAttr zcl.AttrID = 9

type MinBlockRequestDelay zcl.Zu16

func (a MinBlockRequestDelay) ID() zcl.AttrID                { return MinBlockRequestDelayAttr }
func (a MinBlockRequestDelay) Cluster() zcl.ClusterID        { return OtauID }
func (a *MinBlockRequestDelay) Value() *MinBlockRequestDelay { return a }
func (a MinBlockRequestDelay) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *MinBlockRequestDelay) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MinBlockRequestDelay(*nt)
	return br, err
}

func (a MinBlockRequestDelay) Readable() bool   { return true }
func (a MinBlockRequestDelay) Writable() bool   { return false }
func (a MinBlockRequestDelay) Reportable() bool { return false }
func (a MinBlockRequestDelay) SceneIndex() int  { return -1 }

func (a MinBlockRequestDelay) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}
