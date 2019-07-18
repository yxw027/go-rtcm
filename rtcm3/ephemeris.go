package rtcm3

import (
	"encoding/binary"
	"github.com/go-restruct/restruct"
)

// GPS Ephemerides
type Message1019 struct {
	AbstractMessage
	SatelliteID  uint8  `struct:"uint8:6"`
	WeekNumber   uint16 `struct:"uint16:10"`
	SVAccuracy   uint8  `struct:"uint8:4"`
	CodeOnL2     uint8  `struct:"uint8:2"`
	IDOT         int16  `struct:"int16:14"`
	IODE         uint8  `struct:"uint8"`
	Toc          uint16 `struct:"uint16"`
	Af2          int8   `struct:"int8"`
	Af1          int16  `struct:"int16"`
	Af0          int32  `struct:"int32:22"`
	IODC         uint16 `struct:"uint16:10"`
	Crs          int16  `struct:"int16"`
	DeltaN       int16  `struct:"int16"`
	M0           int32  `struct:"int32"`
	Cuc          int16  `struct:"int16"`
	Eccentricity uint32 `struct:"uint32"`
	Cus          int16  `struct:"int16"`
	SrA          uint32 `struct:"uint32"`
	Toe          uint16 `struct:"uint16"`
	Cic          int16  `struct:"int16"`
	Omega0       int32  `struct:"int32"`
	Cis          int16  `struct:"int16"`
	I0           int32  `struct:"int32"`
	Crc          int16  `struct:"int16"`
	Perigee      int32  `struct:"int32"`
	OmegaDot     int32  `struct:"int32:24"`
	Tgd          int8   `struct:"int8"`
	SVHealth     uint8  `struct:"uint8:6"`
	L2PDataFlag  bool   `struct:"uint8:1,variantbool"`
	FitInterval  bool   `struct:"uint8:1,variantbool"`
}

func DeserializeMessage1019(data []byte) (msg Message1019) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1019) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

// GLONASS Ephemerides
type Message1020 struct {
	AbstractMessage
	SatelliteId               uint8  `struct:"uint8:6"`
	FrequencyChannel          uint8  `struct:"uint8:5"`
	AlmanacHealth             bool   `struct:"uint8:1,variantbool"`
	AlmanacHealthAvailability bool   `struct:"uint8:1,variantbool"`
	P1                        uint8  `struct:"uint8:2"`
	Tk                        uint16 `struct:"uint16:12"`
	Msb                       bool   `struct:"uint8:1,variantbool"`
	P2                        bool   `struct:"uint8:1,variantbool"`
	Tb                        uint8  `struct:"uint8:7"`
	XnTb1                     int32  `struct:"int32:24"` // Sint
	XnTb                      int32  `struct:"int32:27"` // Sint
	XnTb2                     int8   `struct:"int8:5"`   // Sint
	YnTb1                     int32  `struct:"int32:24"` // Sint
	YnTb                      int32  `struct:"int32:27"` // Sint
	YnTb2                     int8   `struct:"int8:5"`   // Sint
	ZnTb1                     int32  `struct:"int32:24"` // Sint
	ZnTb                      int32  `struct:"int32:27"` // Sint
	ZnTb2                     int8   `struct:"int8:5"`   // Sint
	P3                        bool   `struct:"uint8:1,variantbool"`
	GammaN                    int16  `struct:"int16:11"` // Sint
	Mp                        uint8  `struct:"uint8:2"`
	M1n3                      bool   `struct:"uint8:1,variantbool"`
	TauN                      int32  `struct:"int32:22"` // Sint
	MDeltaTauN                int8   `struct:"int8:5"`   // Sint
	En                        uint8  `struct:"uint8:5"`
	MP4                       bool   `struct:"uint8:1,variantbool"`
	MFt                       uint8  `struct:"uint8:4"`
	MNt                       uint16 `struct:"uint16:11"`
	MM                        uint8  `struct:"uint8:2"`
	AdditionalData            bool   `struct:"uint8:1,variantbool"`
	Na                        uint16 `struct:"uint16:11"`
	TauC                      int32  `struct:"int32"` // Sint
	MN4                       uint8  `struct:"uint8:5"`
	MTauGps                   int32  `struct:"int32:22"` // Sint
	M1n5                      bool   `struct:"uint8:1,variantbool"`
	Reserved                  uint8  `struct:"uint8:7"`
}

func DeserializeMessage1020(data []byte) (msg Message1020) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1020) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

// TODO: Add 1042, 1044, 1045, 1046 (last messages)

// BDS Satellite Ephemeris Data
// QZSS Ephemerides
// Galileo F/NAV Satellite Ephemeris Data
// Galileo I/NAV Satellite Ephemeris Data
