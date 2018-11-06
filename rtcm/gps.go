package rtcm //TODO: come up with better type naming convention - perhaps make the package name rtcm3 and remove that as a prefix

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
        SmoothingInterval: r.Uint8(3),
    }
}

type Rtcm31001SatelliteData struct {
    SatelliteId uint8
    L1CodeIndicator bool
    L1Pseudorange uint32
    L1PhaseRange int32 //PhaseRange-Pseudorange?
    L1LockTimeIndicator uint8
}

func NewRtcm31001SatelliteData(r *iobit.Reader, numSats int) (satData []Rtcm31001SatelliteData) {
    for i := 0; i < numSats; i++ {
        satData = append(satData, Rtcm31001SatelliteData{
            SatelliteId: r.Uint8(6),
            L1CodeIndicator: r.Bit(),
            L1Pseudorange: r.Uint32(24),
            L1PhaseRange: r.Int32(20),
            L1LockTimeIndicator: r.Uint8(7),
        })
    }
    return satData
}

type Rtcm3Message1001 struct {
    Rtcm3Frame
    Header Rtcm3GpsObservationHeader
    SatelliteData []Rtcm31001SatelliteData
}

func NewRtcm3Message1001(f Rtcm3Frame) Rtcm3Message1001 {
    r := iobit.NewReader(f.Payload)
    header := NewRtcm3GpsObservationHeader(&r)
    return Rtcm3Message1001{
        Rtcm3Frame: f,
        Header: header,
        SatelliteData: NewRtcm31001SatelliteData(&r, int(header.SignalsProcessed)),
    }
}

type Rtcm31002SatelliteData struct {
    SatelliteId uint8
    L1CodeIndicator bool
    L1Pseudorange uint32
    L1PhaseRange int32
    L1LockTimeIndicator uint8
    L1PseudorangeAmbiguity uint8
    L1Cnr uint8
}

func NewRtcm31002SatelliteData(r *iobit.Reader, numSats int) (satData []Rtcm31002SatelliteData) {
    for i := 0; i < numSats; i++ {
        satData = append(satData, Rtcm31002SatelliteData{
            SatelliteId: r.Uint8(6),
            L1CodeIndicator: r.Bit(),
            L1Pseudorange: r.Uint32(24),
            L1PhaseRange: r.Int32(20),
            L1LockTimeIndicator: r.Uint8(7),
            L1PseudorangeAmbiguity: r.Uint8(8),
            L1Cnr: r.Uint8(8),
        })
    }
    return satData
}

type Rtcm3Message1002 struct {
    Rtcm3Frame
    Header Rtcm3GpsObservationHeader
    SatelliteData []Rtcm31002SatelliteData
}

func NewRtcm3Message1002(f Rtcm3Frame) Rtcm3Message1002 {
    r := iobit.NewReader(f.Payload)
    header := NewRtcm3GpsObservationHeader(&r)
    return Rtcm3Message1002{
        Rtcm3Frame: f,
        Header: header,
        SatelliteData: NewRtcm31002SatelliteData(&r, int(header.SignalsProcessed)),
    }
}

type Rtcm31003SatelliteData struct {
    SatelliteId uint8
    L1CodeIndicator bool
    L1Pseudorange uint32
    L1PhaseRange int32
    L1LockTimeIndicator uint8
    L2CodeIndicator uint8
    L2PseudorangeDifference int16
    L2PhaseRange int32
    L2LockTimeIndicator uint8
}

func NewRtcm31003SatelliteData(r *iobit.Reader, numSats int) (satData []Rtcm31003SatelliteData) {
    for i := 0; i < numSats; i++ {
        satData = append(satData, Rtcm31003SatelliteData{
            SatelliteId: r.Uint8(6),
            L1CodeIndicator: r.Bit(),
            L1Pseudorange: r.Uint32(24),
            L1PhaseRange: r.Int32(20),
            L1LockTimeIndicator: r.Uint8(7),
            L2CodeIndicator: r.Uint8(2),
            L2PseudorangeDifference: r.Int16(14),
            L2PhaseRange: r.Int32(20),
            L2LockTimeIndicator: r.Uint8(7),
        })
    }
    return satData
}

type Rtcm3Message1003 struct {
    Rtcm3Frame
    Header Rtcm3GpsObservationHeader
    SatelliteData []Rtcm31003SatelliteData
}

func NewRtcm3Message1003(f Rtcm3Frame) Rtcm3Message1003 {
    r := iobit.NewReader(f.Payload)
    header := NewRtcm3GpsObservationHeader(&r)
    return Rtcm3Message1003{
        Rtcm3Frame: f,
        Header: header,
        SatelliteData: NewRtcm31003SatelliteData(&r, int(header.SignalsProcessed)),
    }
}

type Rtcm31004SatelliteData struct {
    SatelliteId uint8
    L1CodeIndicator bool
    L1Pseudorange uint32
    L1PhaseRange int32
    L1LockTimeIndicator uint8
    L1PseudorangeAmbiguity uint8
    L1Cnr uint8
    L2CodeIndicator uint8
    L2PseudorangeDifference int16
    L2PhaseRange int32
    L2LockTimeIndicator uint8
    L2Cnr uint8
}

func NewRtcm31004SatelliteData(r *iobit.Reader, numSats int) (satData []Rtcm31004SatelliteData) {
    for i := 0; i < numSats; i++ {
        satData = append(satData, Rtcm31004SatelliteData{
            SatelliteId: r.Uint8(6),
            L1CodeIndicator: r.Bit(),
            L1Pseudorange: r.Uint32(24),
            L1PhaseRange: r.Int32(20),
            L1LockTimeIndicator: r.Uint8(7),
            L1PseudorangeAmbiguity: r.Uint8(8),
            L1Cnr: r.Uint8(8),
            L2CodeIndicator: r.Uint8(2),
            L2PseudorangeDifference: r.Int16(14),
            L2PhaseRange: r.Int32(20),
            L2LockTimeIndicator: r.Uint8(7),
            L2Cnr: r.Uint8(8),
        })
    }
    return satData
}

type Rtcm3Message1004 struct {
    Rtcm3Frame
    Header Rtcm3GpsObservationHeader
    SatelliteData []Rtcm31004SatelliteData
}

func NewRtcm3Message1004(f Rtcm3Frame) Rtcm3Message1004 {
    r := iobit.NewReader(f.Payload)
    header := NewRtcm3GpsObservationHeader(&r)
    return Rtcm3Message1004{
        Rtcm3Frame: f,
        Header: header,
        SatelliteData: NewRtcm31004SatelliteData(&r, int(header.SignalsProcessed)),
    }
}
