package rtcm3

import (
	"encoding/binary"
	"github.com/bamiaux/iobit"
)

type AntennaReferencePoint struct {
	MessageNumber             uint16
	ReferenceStationId        uint16
	ItrfRealizationYear       uint8
	GpsIndicator              bool
	GlonassIndicator          bool
	GalileoIndicator          bool
	ReferenceStationIndicator bool
	ReferencePointX           int64
	SingleReceiverOscilator   bool
	Reserved                  bool
	ReferencePointY           int64
	QuarterCycleIndicator     uint8
	ReferencePointZ           int64
}

func (arp AntennaReferencePoint) Number() uint16 {
	return arp.MessageNumber
}

func SerializeAntennaReferencePoint(arp AntennaReferencePoint) []byte {
	data := make([]byte, 19)
	w := iobit.NewWriter(data)
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
	w.PutInt64(38, arp.ReferencePointZ)
	w.Flush()
	return data
}

func DeserializeAntennaReferencePoint(data []byte) AntennaReferencePoint {
	r := iobit.NewReader(data)
	return AntennaReferencePoint{
		MessageNumber:             r.Uint16(12),
		ReferenceStationId:        r.Uint16(12),
		ItrfRealizationYear:       r.Uint8(6),
		GpsIndicator:              r.Bit(),
		GlonassIndicator:          r.Bit(),
		GalileoIndicator:          r.Bit(),
		ReferenceStationIndicator: r.Bit(),
		ReferencePointX:           r.Int64(38),
		SingleReceiverOscilator:   r.Bit(),
		Reserved:                  r.Bit(),
		ReferencePointY:           r.Int64(38),
		QuarterCycleIndicator:     r.Uint8(2),
		ReferencePointZ:           r.Int64(38),
	}
}

// Stationary RTK Reference Station ARP
type Message1005 struct {
	AntennaReferencePoint
}

func DeserializeMessage1005(data []byte) Message1005 {
	return Message1005{
		AntennaReferencePoint: DeserializeAntennaReferencePoint(data),
	}
}

func (msg Message1005) Serialize() (data []byte) {
	return SerializeAntennaReferencePoint(msg.AntennaReferencePoint)
}

// Stationary RTK Reference Station ARP with Antenna Height
type Message1006 struct {
	AntennaReferencePoint
	AntennaHeight uint16
}

func DeserializeMessage1006(data []byte) (msg Message1006) {
	msg = Message1006{
		AntennaReferencePoint: DeserializeAntennaReferencePoint(data),
	}
	msg.AntennaHeight = binary.BigEndian.Uint16(data[len(data)-2:])
	return msg
}

func (msg Message1006) Serialize() (data []byte) {
	data = SerializeAntennaReferencePoint(msg.AntennaReferencePoint)
	height := make([]byte, 2)
	binary.BigEndian.PutUint16(height, msg.AntennaHeight)
	return append(data, height...)
}

type MessageAntennaDescriptor struct {
	MessageNumber      uint16
	ReferenceStationId uint16
	AntennaDescriptor  string
	AntennaSetupId     uint8
}

func (ad MessageAntennaDescriptor) Number() uint16 {
	return ad.MessageNumber
}

func DeserializeAntennaDescriptor(r *iobit.Reader) (desc MessageAntennaDescriptor) {
	desc = MessageAntennaDescriptor{
		MessageNumber:      r.Uint16(12),
		ReferenceStationId: r.Uint16(12),
	}
	desc.AntennaDescriptor = r.String(int(r.Uint8(8)))
	desc.AntennaSetupId = r.Uint8(8)
	return desc
}

func SerializeAntennaDescriptor(desc MessageAntennaDescriptor) []byte {
	data := make([]byte, 4)
	w := iobit.NewWriter(data)
	w.PutUint16(12, desc.MessageNumber)
	w.PutUint16(12, desc.ReferenceStationId)
	w.PutUint8(8, uint8(len(desc.AntennaDescriptor)))
	w.Flush()
	data = append(data, []byte(desc.AntennaDescriptor)...)
	return append(data, desc.AntennaSetupId)
}

// Antenna Descriptor
type Message1007 struct {
	MessageAntennaDescriptor
}

func DeserializeMessage1007(data []byte) Message1007 {
	r := iobit.NewReader(data)
	return Message1007{
		MessageAntennaDescriptor: DeserializeAntennaDescriptor(&r),
	}
}

func (msg Message1007) Serialize() []byte {
	return SerializeAntennaDescriptor(msg.MessageAntennaDescriptor)
}

// Antenna Descriptor & Serial Number
type Message1008 struct {
	MessageAntennaDescriptor
	SerialNumber string
}

func DeserializeMessage1008(data []byte) (msg Message1008) {
	r := iobit.NewReader(data)
	msg = Message1008{
		MessageAntennaDescriptor: DeserializeAntennaDescriptor(&r),
	}
	msg.SerialNumber = r.String(int(r.Uint8(8)))
	return msg
}

func (msg Message1008) Serialize() []byte {
	data := SerializeAntennaDescriptor(msg.MessageAntennaDescriptor)
	data = append(data, uint8(len(msg.SerialNumber)))
	return append(data, []byte(msg.SerialNumber)...)
}

// Receiver and Antenna Descriptors
type Message1033 struct {
	MessageAntennaDescriptor
	AntennaSerialNumber     string
	ReceiverTypeDescriptor  string
	ReceiverFirmwareVersion string
	ReceiverSerialNumber    string
}

func (msg Message1033) Number() uint16 {
	return msg.MessageNumber
}

func DeserializeMessage1033(data []byte) (msg Message1033) {
	r := iobit.NewReader(data)
	msg = Message1033{
		MessageAntennaDescriptor: DeserializeAntennaDescriptor(&r),
	}
	msg.AntennaSerialNumber = r.String(int(r.Uint8(8)))
	msg.ReceiverTypeDescriptor = r.String(int(r.Uint8(8)))
	msg.ReceiverFirmwareVersion = r.String(int(r.Uint8(8)))
	msg.ReceiverSerialNumber = r.String(int(r.Uint8(8)))
	return msg
}

func (msg Message1033) Serialize() []byte {
	data := SerializeAntennaDescriptor(msg.MessageAntennaDescriptor)
	data = append(data, uint8(len(msg.AntennaSerialNumber)))
	data = append(data, []byte(msg.AntennaSerialNumber)...)
	data = append(data, uint8(len(msg.ReceiverTypeDescriptor)))
	data = append(data, []byte(msg.ReceiverTypeDescriptor)...)
	data = append(data, uint8(len(msg.ReceiverFirmwareVersion)))
	data = append(data, []byte(msg.ReceiverFirmwareVersion)...)
	data = append(data, uint8(len(msg.ReceiverSerialNumber)))
	data = append(data, []byte(msg.ReceiverSerialNumber)...)
	return data
}
