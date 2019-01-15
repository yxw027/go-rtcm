package rtcm

import (
    "github.com/bamiaux/iobit"
)

type Rtcm3ResidualSatelliteData struct {
    SatelliteId uint8
    Soc uint8
    Sod uint16
    Soh uint8
    SIc uint16
    SId uint16
}

func NewRtcm3ResidualSatelliteData(r *iobit.Reader, nsat int) (satData []Rtcm3ResidualSatelliteData) {
    for i := 0; i < nsat; i++ {
        satData = append(satData, Rtcm3ResidualSatelliteData{
            SatelliteId: r.Uint8(6),
            Soc: r.Uint8(8),
            Sod: r.Uint16(9),
            Soh: r.Uint8(6),
            SIc: r.Uint16(10),
            SId: r.Uint16(10),
        })
    }
    return satData
}

type Rtcm3Message1030 struct {
    MessageNumber uint16
    Epoch uint32
    ReferenceStationId uint16
    NRefs uint8
    Satellites uint8
    SatelliteData []Rtcm3ResidualSatelliteData
}

func NewRtcm3Message1030(data []byte) (msg Rtcm3Message1030) {
    r := iobit.NewReader(data)
    msg = Rtcm3Message1030{
        MessageNumber: r.Uint16(12),
        Epoch: r.Uint32(20),
        ReferenceStationId: r.Uint16(12),
        NRefs: r.Uint8(7),
        Satellites: r.Uint8(5),
    }
    msg.SatelliteData = NewRtcm3ResidualSatelliteData(&r, int(msg.Satellites))
    return msg
}

func (msg Rtcm3Message1030) Serialize() (data []byte) {
    return data
}

// Need to implement a Time method for GLONASS Residuals Epoch Time - DF225
type Rtcm3Message1031 struct {
    MessageNumber uint16
    Epoch uint32
    ReferenceStationId uint16
    NRefs uint8
    Satellites uint8
    SatelliteData []Rtcm3ResidualSatelliteData
}

func NewRtcm3Message1031(data []byte) (msg Rtcm3Message1031) {
    r := iobit.NewReader(data)
    msg = Rtcm3Message1031{
        MessageNumber: r.Uint16(12),
        Epoch: r.Uint32(17),
        ReferenceStationId: r.Uint16(12),
        NRefs: r.Uint8(7),
        Satellites: r.Uint8(5),
    }
    msg.SatelliteData = NewRtcm3ResidualSatelliteData(&r, int(msg.Satellites))
    return msg
}

func (msg Rtcm3Message1031) Serialize() (data []byte) {
    return data
}

type Rtcm3Message1032 struct {
    MessageNumber uint16
    NonPhysicalReferenceStationId uint16
    PhysicalReferenceStationId uint16
    EpochYear uint8
    ArpEcefX int64
    ArpEcefY int64
    ArpEcefZ int64
}

func NewRtcm3Message1032(data []byte) Rtcm3Message1032 {
    r := iobit.NewReader(data)
    return Rtcm3Message1032{
        MessageNumber: r.Uint16(12),
        NonPhysicalReferenceStationId: r.Uint16(12),
        PhysicalReferenceStationId: r.Uint16(12),
        EpochYear: r.Uint8(6),
        ArpEcefX: r.Int64(38),
        ArpEcefY: r.Int64(38),
        ArpEcefZ: r.Int64(38),
    }
}

func (msg Rtcm3Message1032) Serialize() (data []byte) {
    return data
}
