// Over the air upgrade.
package ota_upgrade

import (
	"hemtjan.st/zcl"
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

func (v QueryNextImage) ControlFieldString() string {
	var bstr []string
	if zcl.BitmapTest([]byte(v.ControlField), 1) {
		bstr = append(bstr, "Hardware version present")
	}
	return zcl.StrJoin(bstr, ", ")
}
func (v QueryNextImage) ManufacturerIdString() string {
	return zcl.Sprintf("0x%X", zcl.Zu16(v.ManufacturerId))
}
func (v QueryNextImage) ImageTypeString() string {
	switch v.ImageType {
	case 0x0000:
		return "Specific image"
	case 0xFFC0:
		return "Security credential"
	case 0xFFC1:
		return "Configuration"
	case 0xFFC2:
		return "Log"
	case 0xFFFF:
		return "Wild card"
	}
	return zcl.Sprintf("%v", zcl.Zenum16(v.ImageType))
}
func (v QueryNextImage) ApplicationReleaseString() string {
	return zcl.Sprintf("0x%X", zcl.Zu8(v.ApplicationRelease))
}
func (v QueryNextImage) ApplicationBuildString() string {
	return zcl.Sprintf("0x%X", zcl.Zu8(v.ApplicationBuild))
}
func (v QueryNextImage) StackReleaseString() string {
	return zcl.Sprintf("0x%X", zcl.Zu8(v.StackRelease))
}
func (v QueryNextImage) StackBuildString() string {
	return zcl.Sprintf("0x%X", zcl.Zu8(v.StackBuild))
}

func (v QueryNextImage) String() string {
	var str []string
	str = append(str, "ControlField["+v.ControlFieldString()+"]")
	str = append(str, "ManufacturerId["+v.ManufacturerIdString()+"]")
	str = append(str, "ImageType["+v.ImageTypeString()+"]")
	str = append(str, "ApplicationRelease["+v.ApplicationReleaseString()+"]")
	str = append(str, "ApplicationBuild["+v.ApplicationBuildString()+"]")
	str = append(str, "StackRelease["+v.StackReleaseString()+"]")
	str = append(str, "StackBuild["+v.StackBuildString()+"]")
	return "QueryNextImage{" + zcl.StrJoin(str, " ") + "}"
}

func (QueryNextImage) Name() string { return "Query next image" }

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

func (v UpgradeEndResponse) ManufacturerIdString() string {
	return zcl.Sprintf("0x%X", zcl.Zu16(v.ManufacturerId))
}
func (v UpgradeEndResponse) ImageTypeString() string {
	return zcl.Sprintf("0x%X", zcl.Zu16(v.ImageType))
}
func (v UpgradeEndResponse) FileVersionString() string {
	return zcl.Sprintf("0x%X", zcl.Zu32(v.FileVersion))
}
func (v UpgradeEndResponse) CurrentTimeString() string {
	return zcl.Sprintf("0x%X", zcl.Zu32(v.CurrentTime))
}
func (v UpgradeEndResponse) UpgradeTimeString() string {
	return zcl.Sprintf("0x%X", zcl.Zu32(v.UpgradeTime))
}

func (v UpgradeEndResponse) String() string {
	var str []string
	str = append(str, "ManufacturerId["+v.ManufacturerIdString()+"]")
	str = append(str, "ImageType["+v.ImageTypeString()+"]")
	str = append(str, "FileVersion["+v.FileVersionString()+"]")
	str = append(str, "CurrentTime["+v.CurrentTimeString()+"]")
	str = append(str, "UpgradeTime["+v.UpgradeTimeString()+"]")
	return "UpgradeEndResponse{" + zcl.StrJoin(str, " ") + "}"
}

func (UpgradeEndResponse) Name() string { return "Upgrade end response" }

// UpgradeServer is an autogenerated attribute in the Otau cluster
type UpgradeServer zcl.Zuid

const UpgradeServerAttr zcl.AttrID = 0

func (UpgradeServer) ID() zcl.AttrID                { return UpgradeServerAttr }
func (UpgradeServer) Cluster() zcl.ClusterID        { return OtauID }
func (UpgradeServer) Name() string                  { return "Upgrade server" }
func (UpgradeServer) Readable() bool                { return true }
func (UpgradeServer) Writable() bool                { return true }
func (UpgradeServer) Reportable() bool              { return false }
func (UpgradeServer) SceneIndex() int               { return -1 }
func (a *UpgradeServer) Value() *UpgradeServer      { return a }
func (a UpgradeServer) MarshalZcl() ([]byte, error) { return zcl.Zuid(a).MarshalZcl() }

func (a *UpgradeServer) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zuid)
	br, err := nt.UnmarshalZcl(b)
	*a = UpgradeServer(*nt)
	return br, err
}

func (a UpgradeServer) String() string {
	return zcl.Sprintf("%v", zcl.Zuid(a))
}

// FileOffset is an autogenerated attribute in the Otau cluster
type FileOffset zcl.Zu32

const FileOffsetAttr zcl.AttrID = 1

func (FileOffset) ID() zcl.AttrID                { return FileOffsetAttr }
func (FileOffset) Cluster() zcl.ClusterID        { return OtauID }
func (FileOffset) Name() string                  { return "File Offset" }
func (FileOffset) Readable() bool                { return true }
func (FileOffset) Writable() bool                { return false }
func (FileOffset) Reportable() bool              { return false }
func (FileOffset) SceneIndex() int               { return -1 }
func (a *FileOffset) Value() *FileOffset         { return a }
func (a FileOffset) MarshalZcl() ([]byte, error) { return zcl.Zu32(a).MarshalZcl() }

func (a *FileOffset) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = FileOffset(*nt)
	return br, err
}

func (a FileOffset) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(a))
}

// CurrentFileVersion is an autogenerated attribute in the Otau cluster
type CurrentFileVersion zcl.Zu32

const CurrentFileVersionAttr zcl.AttrID = 2

func (CurrentFileVersion) ID() zcl.AttrID                { return CurrentFileVersionAttr }
func (CurrentFileVersion) Cluster() zcl.ClusterID        { return OtauID }
func (CurrentFileVersion) Name() string                  { return "Current File Version" }
func (CurrentFileVersion) Readable() bool                { return true }
func (CurrentFileVersion) Writable() bool                { return false }
func (CurrentFileVersion) Reportable() bool              { return false }
func (CurrentFileVersion) SceneIndex() int               { return -1 }
func (a *CurrentFileVersion) Value() *CurrentFileVersion { return a }
func (a CurrentFileVersion) MarshalZcl() ([]byte, error) { return zcl.Zu32(a).MarshalZcl() }

func (a *CurrentFileVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentFileVersion(*nt)
	return br, err
}

func (a CurrentFileVersion) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(a))
}

// CurrentZigbeeStackVersion is an autogenerated attribute in the Otau cluster
type CurrentZigbeeStackVersion zcl.Zu16

const CurrentZigbeeStackVersionAttr zcl.AttrID = 3

func (CurrentZigbeeStackVersion) ID() zcl.AttrID                       { return CurrentZigbeeStackVersionAttr }
func (CurrentZigbeeStackVersion) Cluster() zcl.ClusterID               { return OtauID }
func (CurrentZigbeeStackVersion) Name() string                         { return "Current ZigBee Stack Version" }
func (CurrentZigbeeStackVersion) Readable() bool                       { return true }
func (CurrentZigbeeStackVersion) Writable() bool                       { return false }
func (CurrentZigbeeStackVersion) Reportable() bool                     { return false }
func (CurrentZigbeeStackVersion) SceneIndex() int                      { return -1 }
func (a *CurrentZigbeeStackVersion) Value() *CurrentZigbeeStackVersion { return a }
func (a CurrentZigbeeStackVersion) MarshalZcl() ([]byte, error)        { return zcl.Zu16(a).MarshalZcl() }

func (a *CurrentZigbeeStackVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentZigbeeStackVersion(*nt)
	return br, err
}

func (a CurrentZigbeeStackVersion) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// DownloadedFileVersion is an autogenerated attribute in the Otau cluster
type DownloadedFileVersion zcl.Zu32

const DownloadedFileVersionAttr zcl.AttrID = 4

func (DownloadedFileVersion) ID() zcl.AttrID                   { return DownloadedFileVersionAttr }
func (DownloadedFileVersion) Cluster() zcl.ClusterID           { return OtauID }
func (DownloadedFileVersion) Name() string                     { return "Downloaded File Version" }
func (DownloadedFileVersion) Readable() bool                   { return true }
func (DownloadedFileVersion) Writable() bool                   { return false }
func (DownloadedFileVersion) Reportable() bool                 { return false }
func (DownloadedFileVersion) SceneIndex() int                  { return -1 }
func (a *DownloadedFileVersion) Value() *DownloadedFileVersion { return a }
func (a DownloadedFileVersion) MarshalZcl() ([]byte, error)    { return zcl.Zu32(a).MarshalZcl() }

func (a *DownloadedFileVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = DownloadedFileVersion(*nt)
	return br, err
}

func (a DownloadedFileVersion) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(a))
}

// DownloadedZigbeeStackVersion is an autogenerated attribute in the Otau cluster
type DownloadedZigbeeStackVersion zcl.Zu16

const DownloadedZigbeeStackVersionAttr zcl.AttrID = 5

func (DownloadedZigbeeStackVersion) ID() zcl.AttrID                          { return DownloadedZigbeeStackVersionAttr }
func (DownloadedZigbeeStackVersion) Cluster() zcl.ClusterID                  { return OtauID }
func (DownloadedZigbeeStackVersion) Name() string                            { return "Downloaded ZigBee Stack Version" }
func (DownloadedZigbeeStackVersion) Readable() bool                          { return true }
func (DownloadedZigbeeStackVersion) Writable() bool                          { return false }
func (DownloadedZigbeeStackVersion) Reportable() bool                        { return false }
func (DownloadedZigbeeStackVersion) SceneIndex() int                         { return -1 }
func (a *DownloadedZigbeeStackVersion) Value() *DownloadedZigbeeStackVersion { return a }
func (a DownloadedZigbeeStackVersion) MarshalZcl() ([]byte, error)           { return zcl.Zu16(a).MarshalZcl() }

func (a *DownloadedZigbeeStackVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = DownloadedZigbeeStackVersion(*nt)
	return br, err
}

func (a DownloadedZigbeeStackVersion) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// ImageUpgradeStatus is an autogenerated attribute in the Otau cluster
type ImageUpgradeStatus zcl.Zenum8

const ImageUpgradeStatusAttr zcl.AttrID = 6

func (ImageUpgradeStatus) ID() zcl.AttrID                { return ImageUpgradeStatusAttr }
func (ImageUpgradeStatus) Cluster() zcl.ClusterID        { return OtauID }
func (ImageUpgradeStatus) Name() string                  { return "Image upgrade status" }
func (ImageUpgradeStatus) Readable() bool                { return true }
func (ImageUpgradeStatus) Writable() bool                { return true }
func (ImageUpgradeStatus) Reportable() bool              { return false }
func (ImageUpgradeStatus) SceneIndex() int               { return -1 }
func (a *ImageUpgradeStatus) Value() *ImageUpgradeStatus { return a }
func (a ImageUpgradeStatus) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *ImageUpgradeStatus) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = ImageUpgradeStatus(*nt)
	return br, err
}

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
	return zcl.Sprintf("%v", zcl.Zenum8(a))
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

// MinBlockRequestDelay is an autogenerated attribute in the Otau cluster
type MinBlockRequestDelay zcl.Zu16

const MinBlockRequestDelayAttr zcl.AttrID = 9

func (MinBlockRequestDelay) ID() zcl.AttrID                  { return MinBlockRequestDelayAttr }
func (MinBlockRequestDelay) Cluster() zcl.ClusterID          { return OtauID }
func (MinBlockRequestDelay) Name() string                    { return "Min block request delay" }
func (MinBlockRequestDelay) Readable() bool                  { return true }
func (MinBlockRequestDelay) Writable() bool                  { return false }
func (MinBlockRequestDelay) Reportable() bool                { return false }
func (MinBlockRequestDelay) SceneIndex() int                 { return -1 }
func (a *MinBlockRequestDelay) Value() *MinBlockRequestDelay { return a }
func (a MinBlockRequestDelay) MarshalZcl() ([]byte, error)   { return zcl.Zu16(a).MarshalZcl() }

func (a *MinBlockRequestDelay) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MinBlockRequestDelay(*nt)
	return br, err
}

func (a MinBlockRequestDelay) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}
