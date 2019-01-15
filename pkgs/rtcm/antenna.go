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

func (arp AntennaReferencePoint) Number() uint16 {
    return arp.MessageNumber
}

func SerializeAntennaReferencePoint(w iobit.Writer, arp AntennaReferencePoint) {
    w.PutUint16(12, arp.MessageNumber)
    w.PutUint16(12, arp.ReferenceStationId)
    w.PutUint8(6, arp.ItrfRealizationYear)
    w.PutBit(arp.GpsIndicator)
    w.PutBit(arp.GlonassIndicator)
    w.PutBit(arp.GalileoIndicator)
    w.PutBit(arp.ReferenceStationIndicator)
    w.PutInt64(38, arp.ReferencePointX)
    w.PutBit(arp.SingleReceiverOscilator)
    w.PutBit(arp.Reserved)
    w.PutInt64(38, arp.ReferencePointY)
    w.PutUint8(2, arp.QuarterCycleIndicator)
    w.PutInt64(39, arp.ReferencePointZ)
    return
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

func (msg Rtcm3Message1005) Serialize() []byte {
    data := make([]byte, 20)
    w := iobit.NewWriter(data)
    SerializeAntennaReferencePoint(w, msg.AntennaReferencePoint)
    w.PutUint8(uint(w.Bits()), 0)
    w.Flush()
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

func (msg Rtcm3Message1006) Serialize() []byte {
    data := make([]byte, 22)
    w := iobit.NewWriter(data)
    SerializeAntennaReferencePoint(w, msg.AntennaReferencePoint)
    w.PutUint16(16, msg.AntennaHeight)
    w.PutUint8(uint(w.Bits()), 0)
    w.Flush()
    return data
}

// TODO: Add serialization for AntennaDescriptor and implement for 1007 and 1008
type AntennaDescriptor struct {
    MessageNumber uint16
    ReferenceStationId uint16
    DescriptorLength uint8
    Descriptor string
    SetupId uint8
}

func (ad AntennaDescriptor) Number() uint16 {
    return ad.MessageNumber
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

func (msg Rtcm3Message1033) Number() uint16 {
    return msg.MessageNumber
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

func (msg Rtcm3Message1033) Serialize() []byte {
    data := make([]byte, 3)
    w := iobit.NewWriter(data)
    w.PutUint16(12, msg.MessageNumber)
    w.PutUint16(12, msg.ReferenceStationId)
    w.Flush()
    data = append(data, byte(len(msg.AntennaDescriptor)))
    data = append(data, []byte(msg.AntennaDescriptor)...)
    data = append(data, byte(msg.AntennaSetupId), byte(len(msg.AntennaSerialNumber)))
    data = append(data, byte(len(msg.ReceiverTypeDescriptor)))
    data = append(data, []byte(msg.ReceiverTypeDescriptor)...)
    data = append(data, byte(len(msg.ReceiverFirmwareVersion)))
    data = append(data, []byte(msg.ReceiverFirmwareVersion)...)
    data = append(data, byte(len(msg.ReceiverSerialNumber)))
    data = append(data, []byte(msg.ReceiverSerialNumber)...)
    return data
}
