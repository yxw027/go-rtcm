package rtcm3

import (
	"github.com/bamiaux/iobit"
	"math"
)

type MessageAnnouncement struct {
	Id                   uint16
	SyncFlag             bool
	TransmissionInterval uint16
}

// System Parameters
type Message1013 struct {
	MessageNumber      uint16
	ReferenceStationId uint16
	Mjd                uint16
	SecondsOfDay       uint32
	MessageCount       uint8
	LeapSeconds        uint8
	Messages           []MessageAnnouncement
}

func (msg Message1013) Number() int {
	return int(msg.MessageNumber)
}

func DeserializeMessage1013(data []byte) (msg Message1013) {
	r := iobit.NewReader(data)
	msg = Message1013{
		MessageNumber:      r.Uint16(12),
		ReferenceStationId: r.Uint16(12),
		Mjd:                r.Uint16(16),
		SecondsOfDay:       r.Uint32(17),
		MessageCount:       r.Uint8(5),
		LeapSeconds:        r.Uint8(8),
	}
	for i := 0; i < int(msg.MessageCount); i++ {
		msg.Messages = append(msg.Messages, MessageAnnouncement{
			Id:                   r.Uint16(12),
			SyncFlag:             r.Bit(),
			TransmissionInterval: r.Uint16(16),
		})
	}
	return msg
}

func (msg Message1013) Serialize() []byte {
	data := make([]byte, int(math.Ceil(float64(70+(29*int(msg.MessageCount)))/8)))
	w := iobit.NewWriter(data)
	w.PutUint16(12, msg.MessageNumber)
	w.PutUint16(12, msg.ReferenceStationId)
	w.PutUint16(16, msg.Mjd)
	w.PutUint32(17, msg.SecondsOfDay)
	w.PutUint8(5, msg.MessageCount)
	w.PutUint8(8, msg.LeapSeconds)
	for _, a := range msg.Messages {
		w.PutUint16(12, a.Id)
		w.PutBit(a.SyncFlag)
		w.PutUint16(16, a.TransmissionInterval)
	}
	w.PutUint8(uint(w.Bits()), 0)
	w.Flush()
	return data
}

// Unicode Text String
type Message1029 struct {
	MessageNumber      uint16
	ReferenceStationId uint16
	Mjd                uint16
	SecondsOfDay       uint32
	Characters         uint8
	CodeUnitsLength    uint8
	CodeUnits          string
}

func (msg Message1029) Number() int {
	return int(msg.MessageNumber)
}

func DeserializeMessage1029(data []byte) (msg Message1029) {
	r := iobit.NewReader(data)
	msg = Message1029{
		MessageNumber:      r.Uint16(12),
		ReferenceStationId: r.Uint16(12),
		Mjd:                r.Uint16(16),
		SecondsOfDay:       r.Uint32(17),
		Characters:         r.Uint8(7),
		CodeUnitsLength:    r.Uint8(8),
	}
	msg.CodeUnits = r.String(8 * int(msg.CodeUnitsLength))
	return msg
}

func (msg Message1029) Serialize() []byte {
	data := make([]byte, 9)
	w := iobit.NewWriter(data)
	w.PutUint16(12, msg.MessageNumber)
	w.PutUint16(12, msg.ReferenceStationId)
	w.PutUint16(16, msg.Mjd)
	w.PutUint32(17, msg.SecondsOfDay)
	w.PutUint8(7, msg.Characters)
	w.PutUint8(8, msg.CodeUnitsLength)
	w.Flush()
	return append(data, []byte(msg.CodeUnits)...)
}
