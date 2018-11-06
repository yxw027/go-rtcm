package rtcm //TODO: come up with better type naming convention perhaps make the package name rtcm3 and remove that as a prefix

import (
    "github.com/bamiaux/iobit"
)

type Rtcm3GpsObservationHeader struct {
    MessageNumber uint16
    ReferenceStationId uint16
    Epoch uint32
    SynchronousGnss bool
    SignalsProcessed uint8
    SmoothingIndicator bool
    SmoothingInterval uint8
}

func NewRtcm3GpsObservationHeader(r *iobit.Reader) (header Rtcm3GpsObservationHeader) {
    return Rtcm3GpsObservationHeader{
        MessageNumber: r.Uint16(12),
        ReferenceStationId: r.Uint16(12),
        Epoch: r.Uint32(30),
        SynchronousGnss: r.Bit(),
        SignalsProcessed: r.Uint8(5),
        SmoothingIndicator: r.Bit(),
    }
}

type Rtcm31001SatelliteData struct {
    SatelliteId uint8
    CodeIndicator bool
    Pseudorange uint32
    PhaseRange int32 //PhaseRange-Pseudorange?
    LockTimeIndicator uint8
}

func NewRtcm31001SatelliteData(r *iobit.Reader, numSats int) (satData []Rtcm31001SatelliteData) {
    for i := 0; i < numSats; i++ {
        satData = append(satData, Rtcm31001SatelliteData{
            SatelliteId: r.Uint8(6),
            CodeIndicator: r.Bit(),
            Pseudorange: r.Uint32(24),
            PhaseRange: r.Int32(20),
            LockTimeIndicator: r.Uint8(7),
        })
    }
    return satData
}

type Rtcm3Message1001 struct {
    Rtcm3Frame
    Header Rtcm3GpsObservationHeader
    SatelliteData []Rtcm31001SatelliteData
}

func NewRtcm3Message1001(msg Rtcm3Frame) Rtcm3Message1001 {
    r := iobit.NewReader(msg.Payload)
    header := NewRtcm3GpsObservationHeader(&r)
    return Rtcm3Message1001{
        Rtcm3Frame: msg,
        Header: header,
        SatelliteData: NewRtcm31001SatelliteData(&r, int(header.SignalsProcessed)),
    }
}
