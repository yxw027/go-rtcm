package rtcm3

import (
	"github.com/bamiaux/iobit"
)

// Network Auxiliary Station Data Message
type Message1014 struct {
	MessageNumber                uint16 `12`
	NetworkID                    uint8  `8`
	SubnetworkID                 uint8  `4`
	AuxiliaryStationsTransmitted uint8  `5`
	MasterReferenceStationID     uint16 `12`
	AuxiliaryReferenceStationID  uint16 `12`
	AuxMasterDeltaLatitude       int32  `20`
	AuxMasterDeltaLongitude      int32  `21`
	AuxMasterDeltaHeight         int32  `23`
}

func DeserializeMessage1014(data []byte) Message1014 {
	r := iobit.NewReader(data)
	return Message1014{
		MessageNumber:                r.Uint16(12),
		NetworkID:                    r.Uint8(8),
		SubnetworkID:                 r.Uint8(4),
		AuxiliaryStationsTransmitted: r.Uint8(5),
		MasterReferenceStationID:     r.Uint16(12),
		AuxiliaryReferenceStationID:  r.Uint16(12),
		AuxMasterDeltaLatitude:       r.Int32(20),
		AuxMasterDeltaLongitude:      r.Int32(21),
		AuxMasterDeltaHeight:         r.Int32(23),
	}
}

func (msg Message1014) Serialize() (data []byte) {
	data = make([]byte, 15)
	w := iobit.NewWriter(data)
	w.PutUint16(12, msg.MessageNumber)
	w.PutUint8(8, msg.NetworkID)
	w.PutUint8(4, msg.SubnetworkID)
	w.PutUint8(5, msg.AuxiliaryStationsTransmitted)
	w.PutUint16(12, msg.MasterReferenceStationID)
	w.PutUint16(12, msg.AuxiliaryReferenceStationID)
	w.PutInt32(20, msg.AuxMasterDeltaLatitude)
	w.PutInt32(21, msg.AuxMasterDeltaLongitude)
	w.PutInt32(23, msg.AuxMasterDeltaHeight)
	w.PutUint8(uint(w.Bits()), 0) // Pad with 0
	w.Flush()
	return data
}
