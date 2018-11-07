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
    Rtcm3Frame
    AntennaReferencePoint
}

func NewRtcm3Message1005(f Rtcm3Frame) Rtcm3Message1005 {
    r := iobit.NewReader(f.Payload)
    return Rtcm3Message1005{
        Rtcm3Frame: f,
        AntennaReferencePoint: NewAntennaReferencePoint(&r),
    }
}

type Rtcm3Message1006 struct {
    Rtcm3Frame
    AntennaReferencePoint
    AntennaHeight uint16
}

func NewRtcm3Message1006(f Rtcm3Frame) Rtcm3Message1006 {
    r := iobit.NewReader(f.Payload)
    return Rtcm3Message1006{
        Rtcm3Frame: f,
        AntennaReferencePoint: NewAntennaReferencePoint(&r),
        AntennaHeight: r.Uint16(16),
    }
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
    Rtcm3Frame
    AntennaDescriptor
}

func NewRtcm3Message1007(f Rtcm3Frame) Rtcm3Message1007 {
    r := iobit.NewReader(f.Payload)
    return Rtcm3Message1007{
        Rtcm3Frame: f,
        AntennaDescriptor: NewAntennaDescriptor(&r),
    }
}

type Rtcm3Message1008 struct {
    Rtcm3Frame
    AntennaDescriptor
    SerialNumberLength uint8
    SerialNumber string
}

func NewRtcm3Message1008(f Rtcm3Frame) (msg Rtcm3Message1008) {
    r := iobit.NewReader(f.Payload)
    msg = Rtcm3Message1008{
        Rtcm3Frame: f,
        AntennaDescriptor: NewAntennaDescriptor(&r),
        SerialNumberLength: r.Uint8(8),
    }
    msg.SerialNumber = r.String(8 * int(msg.SerialNumberLength))
    return msg
}

type Rtcm3Message1033 struct {
    Rtcm3Frame
    MessageNumber uint16
    ReferenceStationId uint16
    AntennaDescriptor string
    AntennaSetupId uint8
    AntennaSerialNumber string
    ReceiverTypeDescriptor string
    ReceiverFirmwareVersion string
    ReceiverSerialNumber string
}

func NewRtcm3Message1033(f Rtcm3Frame) (msg Rtcm3Message1033) {
    r := iobit.NewReader(f.Payload)
    msg = Rtcm3Message1033{
        Rtcm3Frame: f,
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
