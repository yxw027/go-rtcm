package rtcm

import (
    "github.com/bamiaux/iobit"
)

type Rtcm3GlonassObservationHeader struct {
    MessageNumber uint16
    ReferenceStationId uint16
    Epoch uint32
    SynchronousGnss bool
    SignalCount uint8
    SmoothingIndicator bool
    SmoothingInterval uint8
}

func NewRtcm3GlonassObservationHeader(r *iobit.Reader) Rtcm3GlonassObservationHeader {
    return Rtcm3GlonassObservationHeader{
        MessageNumber: r.Uint16(12),
        ReferenceStationId: r.Uint16(12),
        Epoch: r.Uint32(27),
        SynchronousGnss: r.Bit(),
        SignalCount: r.Uint8(5),
        SmoothingIndicator: r.Bit(),
        SmoothingInterval: r.Uint8(3),
    }
}

type Rtcm31009SignalData struct {
    SatelliteId uint8
    L1CodeIndicator bool
    FrequencyChannel uint8
    L1Pseudorange uint32
    L1PhaseRange int32
    L1LockTimeIndicator uint8
}

func NewRtcm31009SignalData(r *iobit.Reader, numSig int) (sigData []Rtcm31009SignalData) {
    for i := 0; i < numSig; i++ {
        sigData = append(sigData, Rtcm31009SignalData{
            SatelliteId: r.Uint8(6),
            L1CodeIndicator: r.Bit(),
            FrequencyChannel: r.Uint8(5),
            L1Pseudorange: r.Uint32(25),
            L1PhaseRange: r.Int32(20),
            L1LockTimeIndicator: r.Uint8(7),
        })
    }
    return sigData
}

type Rtcm3Message1009 struct {
    Rtcm3Frame
    Header Rtcm3GlonassObservationHeader
    SignalData []Rtcm31009SignalData
}

func NewRtcm3Message1009(f Rtcm3Frame) (msg Rtcm3Message1009) {
    r := iobit.NewReader(f.Payload)
    msg = Rtcm3Message1009{
        Rtcm3Frame: f,
        Header: NewRtcm3GlonassObservationHeader(&r),
    }
    msg.SignalData = NewRtcm31009SignalData(&r, int(msg.Header.SignalCount))
    return msg
}

type Rtcm31010SignalData struct {
    SatelliteId uint8
    L1CodeIndicator bool
    FrequencyChannel uint8
    L1Pseudorange uint32
    L1PhaseRange int32
    L1LockTimeIndicator uint8
    L1PseudorangeAmbiguity uint8
    L1Cnr uint8
}

func NewRtcm31010SignalData(r *iobit.Reader, numSig int) (sigData []Rtcm31010SignalData) {
    for i := 0; i < numSig; i++ {
        sigData = append(sigData, Rtcm31010SignalData{
            SatelliteId: r.Uint8(6),
            L1CodeIndicator: r.Bit(),
            FrequencyChannel: r.Uint8(5),
            L1Pseudorange: r.Uint32(25),
            L1PhaseRange: r.Int32(20),
            L1LockTimeIndicator: r.Uint8(7),
            L1PseudorangeAmbiguity: r.Uint8(7),
            L1Cnr: r.Uint8(8),
        })
    }
    return sigData
}

type Rtcm3Message1010 struct {
    Rtcm3Frame
    Header Rtcm3GlonassObservationHeader
    SignalData []Rtcm31010SignalData
}

func NewRtcm3Message1010(f Rtcm3Frame) (msg Rtcm3Message1010) {
    r := iobit.NewReader(f.Payload)
    msg = Rtcm3Message1010{
        Rtcm3Frame: f,
        Header: NewRtcm3GlonassObservationHeader(&r),
    }
    msg.SignalData = NewRtcm31010SignalData(&r, int(msg.Header.SignalCount))
    return msg
}

type Rtcm31011SignalData struct {
    SatelliteId uint8
    L1CodeIndicator bool
    FrequencyChannel uint8
    L1Pseudorange uint32
    L1PhaseRange int32
    L1LockTimeIndicator uint8
    L2CodeIndicator uint8
    L2Pseudorange uint16
    L2PhaseRange int32
    L2LockTimeIndicator uint8
}

func NewRtcm31011SignalData(r *iobit.Reader, numSig int) (sigData []Rtcm31011SignalData) {
    for i := 0; i < numSig; i++ {
        sigData = append(sigData, Rtcm31011SignalData{
            SatelliteId: r.Uint8(6),
            L1CodeIndicator: r.Bit(),
            FrequencyChannel: r.Uint8(5),
            L1Pseudorange: r.Uint32(25),
            L1PhaseRange: r.Int32(20),
            L1LockTimeIndicator: r.Uint8(7),
            L2CodeIndicator: r.Uint8(2),
            L2Pseudorange: r.Uint16(14),
            L2PhaseRange: r.Int32(20),
            L2LockTimeIndicator: r.Uint8(7),
        })
    }
    return sigData
}

type Rtcm3Message1011 struct {
    Rtcm3Frame
    Header Rtcm3GlonassObservationHeader
    SignalData []Rtcm31011SignalData
}

func NewRtcm3Message1011(f Rtcm3Frame) (msg Rtcm3Message1011) {
    r := iobit.NewReader(f.Payload)
    msg = Rtcm3Message1011{
        Rtcm3Frame: f,
        Header: NewRtcm3GlonassObservationHeader(&r),
    }
    msg.SignalData = NewRtcm31011SignalData(&r, int(msg.Header.SignalCount))
    return msg
}

type Rtcm31012SignalData struct {
    SatelliteId uint8
    L1CodeIndicator bool
    FrequencyChannel uint8
    L1Pseudorange uint32
    L1PhaseRange int32
    L1LockTimeIndicator uint8
    L1PseudorangeAmbiguity uint8
    L1Cnr uint8
    L2CodeIndicator uint8
    L2Pseudorange uint16
    L2PhaseRange int32
    L2LockTimeIndicator uint8
    L2Cnr uint8
}

func NewRtcm31012SignalData(r *iobit.Reader, numSig int) (sigData []Rtcm31012SignalData) {
    for i := 0; i < numSig; i++ {
        sigData = append(sigData, Rtcm31012SignalData{
            SatelliteId: r.Uint8(6),
            L1CodeIndicator: r.Bit(),
            FrequencyChannel: r.Uint8(5),
            L1Pseudorange: r.Uint32(25),
            L1PhaseRange: r.Int32(20),
            L1LockTimeIndicator: r.Uint8(7),
            L1PseudorangeAmbiguity: r.Uint8(7),
            L1Cnr: r.Uint8(8),
            L2CodeIndicator: r.Uint8(2),
            L2Pseudorange: r.Uint16(14),
            L2PhaseRange: r.Int32(20),
            L2LockTimeIndicator: r.Uint8(7),
            L2Cnr: r.Uint8(8),
        })
    }
    return sigData
}

type Rtcm3Message1012 struct {
    Rtcm3Frame
    Header Rtcm3GlonassObservationHeader
    SignalData []Rtcm31012SignalData
}

func NewRtcm3Message1012(f Rtcm3Frame) (msg Rtcm3Message1012) {
    r := iobit.NewReader(f.Payload)
    msg = Rtcm3Message1012{
        Rtcm3Frame: f,
        Header: NewRtcm3GlonassObservationHeader(&r),
    }
    msg.SignalData = NewRtcm31012SignalData(&r, int(msg.Header.SignalCount))
    return msg
}
