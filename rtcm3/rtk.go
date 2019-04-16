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

type NetworkRTKHeader struct {
	MessageNumber               uint16 `12`
	NetworkID                   uint8  `8`
	SubnetworkID                uint8  `4`
	Epoch                       uint32 `23`
	MultipleMessageIndicator    bool   `1`
	MasterReferenceStationID    uint16 `12`
	AuxiliaryReferenceStationID uint16 `12`
	SatelliteCount              uint8  `4`
}

func DeserializeNetworkRTKHeader(r *iobit.Reader) NetworkRTKHeader {
	return NetworkRTKHeader{
		MessageNumber:               r.Uint16(12),
		NetworkID:                   r.Uint8(8),
		SubnetworkID:                r.Uint8(4),
		Epoch:                       r.Uint32(23),
		MultipleMessageIndicator:    r.Bit(),
		MasterReferenceStationID:    r.Uint16(12),
		AuxiliaryReferenceStationID: r.Uint16(12),
		SatelliteCount:              r.Uint8(4),
	}
}

type SatelliteData1015 struct {
	SatelliteID                                 uint8 `6`
	AmbiguityStatusFlag                         uint8 `2`
	NonSyncCount                                uint8 `3`
	IonosphericCarrierPhaseCorrectionDifference int32 `17`
}

func DeserializeSatelliteData1015(r *iobit.Reader, nsat int) (data []SatelliteData1015) {
	for i := 0; i < nsat; i++ {
		data = append(data, SatelliteData1015{
			SatelliteID:         r.Uint8(6),
			AmbiguityStatusFlag: r.Uint8(2),
			NonSyncCount:        r.Uint8(3),
			IonosphericCarrierPhaseCorrectionDifference: r.Int32(17),
		})
	}
	return data
}

// GPS Ionospheric Correction Differences
type Message1015 struct {
	NetworkRTKHeader
	SatelliteData []SatelliteData1015
}

func DeserializeMessage1015(data []byte) (msg Message1015) {
	r := iobit.NewReader(data)
	msg.NetworkRTKHeader = DeserializeNetworkRTKHeader(&r)
	msg.SatelliteData = DeserializeSatelliteData1015(&r, int(msg.NetworkRTKHeader.SatelliteCount))
	return msg
}

type SatelliteData1016 struct {
	SatelliteID                               uint8 `6`
	AmbiguityStatusFlag                       uint8 `2`
	NonSyncCount                              uint8 `3`
	GeometricCarrierPhaseCorrectionDifference int32 `17`
	IODE                                      uint8 `8`
}

func DeserializeSatelliteData1016(r *iobit.Reader, nsat int) (data []SatelliteData1016) {
	for i := 0; i < nsat; i++ {
		data = append(data, SatelliteData1016{
			SatelliteID:         r.Uint8(6),
			AmbiguityStatusFlag: r.Uint8(2),
			NonSyncCount:        r.Uint8(3),
			GeometricCarrierPhaseCorrectionDifference: r.Int32(17),
			IODE: r.Uint8(8),
		})
	}
	return data
}

// GPS Geometric Correction Differences
type Message1016 struct {
	NetworkRTKHeader
	SatelliteData []SatelliteData1016
}

func DeserializeMessage1016(data []byte) (msg Message1016) {
	r := iobit.NewReader(data)
	msg.NetworkRTKHeader = DeserializeNetworkRTKHeader(&r)
	msg.SatelliteData = DeserializeSatelliteData1016(&r, int(msg.NetworkRTKHeader.SatelliteCount))
	return msg
}

type SatelliteData1017 struct {
	SatelliteID                                 uint8 `6`
	AmbiguityStatusFlag                         uint8 `2`
	NonSyncCount                                uint8 `3`
	GeometricCarrierPhaseCorrectionDifference   int32 `17`
	IODE                                        uint8 `8`
	IonosphericCarrierPhaseCorrectionDifference int32 `17`
}

func DeserializeSatelliteData1017(r *iobit.Reader, nsat int) (data []SatelliteData1017) {
	for i := 0; i < nsat; i++ {
		data = append(data, SatelliteData1017{
			SatelliteID:         r.Uint8(6),
			AmbiguityStatusFlag: r.Uint8(2),
			NonSyncCount:        r.Uint8(3),
			GeometricCarrierPhaseCorrectionDifference: r.Int32(17),
			IODE: r.Uint8(8),
			IonosphericCarrierPhaseCorrectionDifference: r.Int32(17),
		})
	}
	return data
}

// GPS Combined Geometric and Ionospheric Correction Differences
type Message1017 struct {
	NetworkRTKHeader
	SatelliteData []SatelliteData1017
}

func DeserializeMessage1017(data []byte) (msg Message1017) {
	r := iobit.NewReader(data)
	msg.NetworkRTKHeader = DeserializeNetworkRTKHeader(&r)
	msg.SatelliteData = DeserializeSatelliteData1017(&r, int(msg.NetworkRTKHeader.SatelliteCount))
	return msg
}
