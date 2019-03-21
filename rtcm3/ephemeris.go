package rtcm3

import (
	"github.com/bamiaux/iobit"
)

type GalileoEphemeris struct {
	MessageNumber uint16
	SatelliteId uint8
	WeekNumber uint16
	IODnav uint16
	SVSISA uint8
	IDOT int16
	Toc uint16
	Af2 int8
	Af1 int32
	Af0 int32
	Crs int16
	DeltaN int16
	M0 int32
	Cuc int16
	E uint32
	Cus int16
	A12 uint32 // this should have a different name
	Toe uint16
	Cic int16
	Omega0 int32
	Cis int16
	I0 int32
	Crc int16
	Omega int32
	OmegaDot int32
	BGDE5aE1 int16
}

func DeserializeGalileoEphemeris(r *iobit.Reader) GalileoEphemeris {
	return GalileoEphemeris{
		MessageNumber: r.Uint16(12),
		SatelliteId: r.Uint8(6),
		WeekNumber: r.Uint16(12),
		IODnav: r.Uint16(10),
		SVSISA: r.Uint8(8),
		IDOT: r.Int16(14),
		Toc: r.Uint16(14),
		Af2: r.Int8(6),
		Af1: r.Int32(21),
		Af0: r.Int32(31),
		Crs: r.Int16(16),
		DeltaN: r.Int16(16),
		M0: r.Int32(32),
		Cuc: r.Int16(16),
		E: r.Uint32(32),
		Cus: r.Int16(16),
		A12: r.Uint32(32),
		Toe: r.Uint16(14),
		Cic: r.Int16(16),
		Omega0: r.Int32(32),
		Cis: r.Int16(16),
		I0: r.Int32(32),
		Crc: r.Int16(16),
		Omega: r.Int32(32),
		OmegaDot: r.Int32(24),
		BGDE5aE1: r.Int16(10),
	}
}

type Message1045 struct {
	GalileoEphemeris
	OSHS uint8
	OSDVS bool
	Reserved uint8
}

func DeserializeMessage1045(data []byte) Message1045 {
	r := iobit.NewReader(data)
	return Message1045{
		GalileoEphemeris: DeserializeGalileoEphemeris(&r),
		OSHS: r.Uint8(2),
		OSDVS: r.Bit(),
		Reserved: r.Uint8(7),
	}
}

func (msg Message1045) Serialize() []byte {
	return []byte{}
}

func (msg Message1045) Number() uint16 {
	return msg.MessageNumber
}

type Message1046 struct {
	GalileoEphemeris
	BGDE5bE1 int16
	E5bSignalHealthStatus uint8
	E5bDataValidityStatus bool
	E1bSignalHealthStatus uint8
	e1bDataValidityStatus bool
	Reserved uint8
}

func DeserializeMessage1046(data []byte) Message1046 {
	r := iobit.NewReader(data)
	return Message1046{
		GalileoEphemeris: DeserializeGalileoEphemeris(&r),
		BGDE5bE1: r.Int16(10),
		E5bSignalHealthStatus: r.Uint8(2),
		E5bDataValidityStatus: r.Bit(),
		E1bSignalHealthStatus: r.Uint8(2),
		e1bDataValidityStatus: r.Bit(),
		Reserved: r.Uint8(2),
	}
}

func (msg Message1046) Serialize() []byte {
	return []byte{}
}

func(msg Message1046) Number() uint16 {
	return msg.MessageNumber
}
