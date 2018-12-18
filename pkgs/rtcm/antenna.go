package rtcm

import (
    "github.com/bamiaux/iobit"
)

type AntennaReferencePoint struct {
    MessageNumber uint16
    ReferenceStationId uint16
    ItrfRealizationYear uint8
    GpsIndicator bool
    GlonassIndicator bool
    GalileoIndicator bool
    ReferenceStationIndicator bool
    ReferencePointX int64
    SingleReceiverOscilator bool
    Reserved bool
    ReferencePointY int64
    QuarterCycleIndicator uint8
    ReferencePointZ int64
}

func NewAntennaReferencePoint(r *iobit.Reader) AntennaReferencePoint {
    return AntennaReferencePoint{
        MessageNumber: r.Uint16(12),
        ReferenceStationId: r.Uint16(12),
        ItrfRealizationYear: r.Uint8(6),
        GpsIndicator: r.Bit(),
        GlonassIndicator: r.Bit(),
        GalileoIndicator: r.Bit(),
        ReferenceStationIndicator: r.Bit(),
        ReferencePointX: r.Int64(38),
        SingleReceiverOscilator: r.Bit(),
        Reserved: r.Bit(),
        ReferencePointY: r.Int64(38),
        QuarterCycleIndicator: r.Uint8(2),
        ReferencePointZ: r.Int64(39),
    }
}

type Rtcm3Message1005 struct {
    AntennaReferencePoint
}

func NewRtcm3Message1005(data []byte) Rtcm3Message1005 {
    r := iobit.NewReader(data)
    return Rtcm3Message1005{
        AntennaReferencePoint: NewAntennaReferencePoint(&r),
    }
}

func (msg Rtcm3Message1005) Serialize() (data []byte) {
    return data
}

type Rtcm3Message1006 struct {
    AntennaReferencePoint
    AntennaHeight uint16
}

func NewRtcm3Message1006(data []byte) Rtcm3Message1006 {
    r := iobit.NewReader(data)
    return Rtcm3Message1006{
        AntennaReferencePoint: NewAntennaReferencePoint(&r),
        AntennaHeight: r.Uint16(16),
    }
}

func (msg Rtcm3Message1006) Serialize() (data []byte) {
    return data
}

type AntennaDescriptor struct {
    MessageNumber uint16
    ReferenceStationId uint16
    DescriptorLength uint8
    Descriptor string
    SetupId uint8
}

func NewAntennaDescriptor(r *iobit.Reader) (desc AntennaDescriptor) {
    desc = AntennaDescriptor{
        MessageNumber: r.Uint16(12),
        ReferenceStationId: r.Uint16(12),
        DescriptorLength: r.Uint8(8),
    }
    desc.Descriptor = r.String(8 * int(desc.DescriptorLength))
    desc.SetupId = r.Uint8(8)
    return desc
}

type Rtcm3Message1007 struct {
    AntennaDescriptor
}

func NewRtcm3Message1007(data []byte) Rtcm3Message1007 {
    r := iobit.NewReader(data)
    return Rtcm3Message1007{
        AntennaDescriptor: NewAntennaDescriptor(&r),
    }
}

func (msg Rtcm3Message1007) Serialize() (data []byte) {
    return data
}

type Rtcm3Message1008 struct {
    AntennaDescriptor
    SerialNumberLength uint8
    SerialNumber string
}

func NewRtcm3Message1008(data []byte) (msg Rtcm3Message1008) {
    r := iobit.NewReader(data)
    msg = Rtcm3Message1008{
        AntennaDescriptor: NewAntennaDescriptor(&r),
        SerialNumberLength: r.Uint8(8),
    }
    msg.SerialNumber = r.String(8 * int(msg.SerialNumberLength))
    return msg
}

func (msg Rtcm3Message1008) Serialize() (data []byte) {
    return data
}

type Rtcm3Message1033 struct {
    MessageNumber uint16
    ReferenceStationId uint16
    AntennaDescriptor string
    AntennaSetupId uint8
    AntennaSerialNumber string
    ReceiverTypeDescriptor string
    ReceiverFirmwareVersion string
    ReceiverSerialNumber string
}

func NewRtcm3Message1033(data []byte) (msg Rtcm3Message1033) {
    r := iobit.NewReader(data)
    msg = Rtcm3Message1033{
        MessageNumber: r.Uint16(12),
        ReferenceStationId: r.Uint16(12),
    }
    msg.AntennaDescriptor = r.String(int(r.Uint8(8)))
    msg.AntennaSetupId = r.Uint8(8)
    msg.AntennaSerialNumber = r.String(int(r.Uint8(8)))
    msg.ReceiverTypeDescriptor = r.String(int(r.Uint8(8)))
    msg.ReceiverFirmwareVersion = r.String(int(r.Uint8(8)))
    msg.ReceiverSerialNumber = r.String(int(r.Uint8(8)))
    return msg
}

func (msg Rtcm3Message1033) Serialize() (data []byte) {
    return data
}
