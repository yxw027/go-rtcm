package rtcm3

import (
	"encoding/binary"
	"github.com/bamiaux/iobit"
	"github.com/go-restruct/restruct"
	"time"
)

type Message1001 struct {
	AbstractMessage
	ReferenceStationId uint16 `struct:"uint16:12"`
	Epoch              uint32 `struct:"uint32:30"`
	SynchronousGnss    bool   `struct:"uint8:1,variantbool"`
	SignalsProcessed   uint8  `struct:"uint8:5,sizeof=SatelliteData"`
	SmoothingIndicator bool   `struct:"uint8:1,variantbool"`
	SmoothingInterval  uint8  `struct:"uint8:3"`
	SatelliteData      []struct {
		SatelliteId         uint8  `struct:"uint8:6"`
		L1CodeIndicator     bool   `struct:"uint8:1,variantbool"`
		L1Pseudorange       uint32 `struct:"uint32:24"`
		L1PhaseRange        int32  `struct:"int32:20"`
		L1LockTimeIndicator uint8  `struct:"uint8:7"`
	}
}

func DeserializeMessage1001(data []byte) (msg Message1001) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1001) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

func (msg Message1001) Time() time.Time {
	return GpsTime(msg.Epoch)
}

// Extended L1-Only GPS RTK Observables
type Message1002 struct {
	AbstractMessage
	ReferenceStationId uint16 `struct:"uint16:12"`
	Epoch              uint32 `struct:"uint32:30"`
	SynchronousGnss    bool   `struct:"uint8:1,variantbool"`
	SignalsProcessed   uint8  `struct:"uint8:5,sizeof=SatelliteData"`
	SmoothingIndicator bool   `struct:"uint8:1,variantbool"`
	SmoothingInterval  uint8  `struct:"uint8:3"`
	SatelliteData      []struct {
		SatelliteId            uint8  `struct:"uint8:6"`
		L1CodeIndicator        bool   `struct:"uint8:1,variantbool"`
		L1Pseudorange          uint32 `struct:"uint32:24"`
		L1PhaseRange           int32  `struct:"int32:20"`
		L1LockTimeIndicator    uint8  `struct:"uint8:7"`
		L1PseudorangeAmbiguity uint8  `struct:"uint8"`
		L1Cnr                  uint8  `struct:"uint8"`
	}
}

func DeserializeMessage1002(data []byte) (msg Message1002) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1002) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

func (msg Message1002) Time() time.Time {
	return GpsTime(msg.Epoch)
}

// L1&L2 GPS RTK Observables
type Message1003 struct {
	AbstractMessage
	ReferenceStationId uint16 `struct:"uint16:12"`
	Epoch              uint32 `struct:"uint32:30"`
	SynchronousGnss    bool   `struct:"uint8:1,variantbool"`
	SignalsProcessed   uint8  `struct:"uint8:5,sizeof=SatelliteData"`
	SmoothingIndicator bool   `struct:"uint8:1,variantbool"`
	SmoothingInterval  uint8  `struct:"uint8:3"`
	SatelliteData      []struct {
		SatelliteId             uint8  `struct:"uint8:6"`
		L1CodeIndicator         bool   `struct:"uint8:1,variantbool"`
		L1Pseudorange           uint32 `struct:"uint32:24"`
		L1PhaseRange            int32  `struct:"int32:20"`
		L1LockTimeIndicator     uint8  `struct:"uint8:7"`
		L2CodeIndicator         uint8  `struct:"uint8:2"`
		L2PseudorangeDifference int16  `struct:"int16:14"`
		L2PhaseRange            int32  `struct:"int32:20"`
		L2LockTimeIndicator     uint8  `struct:"uint8:7"`
	}
}

func DeserializeMessage1003(data []byte) (msg Message1003) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1003) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

func (msg Message1003) Time() time.Time {
	return GpsTime(msg.Epoch)
}

// Extended L1&L2 GPS RTK Observables
type Message1004 struct {
	AbstractMessage
	ReferenceStationId uint16 `struct:"uint16:12"`
	Epoch              uint32 `struct:"uint32:30"`
	SynchronousGnss    bool   `struct:"uint8:1,variantbool"`
	SignalsProcessed   uint8  `struct:"uint8:5,sizeof=SatelliteData"`
	SmoothingIndicator bool   `struct:"uint8:1,variantbool"`
	SmoothingInterval  uint8  `struct:"uint8:3"`
	SatelliteData      []struct {
		SatelliteId             uint8  `struct:"uint8:6"`
		L1CodeIndicator         bool   `struct:"uint8:1,variantbool"`
		L1Pseudorange           uint32 `struct:"uint32:24"`
		L1PhaseRange            int32  `struct:"int32:20"`
		L1LockTimeIndicator     uint8  `struct:"uint8:7"`
		L1PseudorangeAmbiguity  uint8  `struct:"uint8"`
		L1Cnr                   uint8  `struct:"uint8"`
		L2CodeIndicator         uint8  `struct:"uint8:2"`
		L2PseudorangeDifference int16  `struct:"int16:14"`
		L2PhaseRange            int32  `struct:"int32:20"`
		L2LockTimeIndicator     uint8  `struct:"uint8:7"`
		L2Cnr                   uint8  `struct:"uint8"`
	}
}

func DeserializeMessage1004(data []byte) (msg Message1004) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1004) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

func (msg Message1004) Time() time.Time {
	return GpsTime(msg.Epoch)
}

// GPS Ephemerides
type Message1019 struct {
	AbstractMessage
	SatelliteId   uint8  `struct:"uint8:6"`
	GpsWeekNumber uint16 `struct:"uint16:10"`
	SvAccuracy    uint8  `struct:"uint8:4"`
	L2Code        uint8  `struct:"uint8:2"`
	Idot          int16  `struct:"int16:14"`
	Iode          uint8  `struct:"uint8"`
	Toc           uint16 `struct:"uint16"`
	Af2           int8   `struct:"int8"`
	Af1           int16  `struct:"int16"`
	Af0           int32  `struct:"int32:22"`
	Iodc          uint16 `struct:"uint16:10"`
	Crs           int16  `struct:"int16"`
	DeltaN        int16  `struct:"int16"`
	M0            int32  `struct:"int32"`
	Cuc           int16  `struct:"int16"`
	Eccentricity  uint32 `struct:"uint32"`
	Cus           int16  `struct:"int16"`
	SrA           uint32 `struct:"uint32"`
	Toe           uint16 `struct:"uint16"`
	Cic           int16  `struct:"int16"`
	Omega0        int32  `struct:"int32"`
	Cis           int16  `struct:"int16"`
	I0            int32  `struct:"int32"`
	C_rc          int16  `struct:"int16"`
	Perigee       int32  `struct:"int32"`
	OmegaDot      int32  `struct:"int32:24"`
	Tgd           int8   `struct:"int8"`
	SvHealth      uint8  `struct:"uint8:6"`
	L2PDataFlag   bool   `struct:"uint8:1,variantbool"`
	FitInterval   bool   `struct:"uint8:1,variantbool"`
}

func DeserializeMessage1019(data []byte) (msg Message1019) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1019) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

// Network Auxiliary Station Data Message
type Message1014 struct {
	AbstractMessage
	NetworkID                    uint8  `struct:"uint8:8"`
	SubnetworkID                 uint8  `struct:"uint8:4"`
	AuxiliaryStationsTransmitted uint8  `struct:"uint8:5"`
	MasterReferenceStationID     uint16 `struct:"uint16:12"`
	AuxiliaryReferenceStationID  uint16 `struct:"uint16:12"`
	AuxMasterDeltaLatitude       int32  `struct:"int32:20"`
	AuxMasterDeltaLongitude      int32  `struct:"int32:21"`
	AuxMasterDeltaHeight         int32  `struct:"int32:23"`
}

func DeserializeMessage1014(data []byte) (msg Message1014) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1014) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

type NetworkRTKHeader struct {
	MessageNumber               uint16
	NetworkID                   uint8
	SubnetworkID                uint8
	Epoch                       uint32
	MultipleMessageIndicator    bool
	MasterReferenceStationID    uint16
	AuxiliaryReferenceStationID uint16
	SatelliteCount              uint8
}

func (msg NetworkRTKHeader) Number() int {
	return int(msg.MessageNumber)
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
	SatelliteID                                 uint8
	AmbiguityStatusFlag                         uint8
	NonSyncCount                                uint8
	IonosphericCarrierPhaseCorrectionDifference int32
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

func (msg Message1015) Serialize() (data []byte) {
	return []byte{}
}

type SatelliteData1016 struct {
	SatelliteID                               uint8
	AmbiguityStatusFlag                       uint8
	NonSyncCount                              uint8
	GeometricCarrierPhaseCorrectionDifference int32
	IODE                                      uint8
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

func (msg Message1016) Serialize() (data []byte) {
	return []byte{}
}

type SatelliteData1017 struct {
	SatelliteID                                 uint8
	AmbiguityStatusFlag                         uint8
	NonSyncCount                                uint8
	GeometricCarrierPhaseCorrectionDifference   int32
	IODE                                        uint8
	IonosphericCarrierPhaseCorrectionDifference int32
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

func (msg Message1017) Serialize() (data []byte) {
	return []byte{}
}
