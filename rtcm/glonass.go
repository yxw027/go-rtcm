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
    CodeIndicator bool
    FrequencyChannel uint8
    L1Pseudorange uint32
    L1PhaseRange int32
    L1LockTimeIndicator uint8
}

func NewRtcm31009SignalData(r *iobit.Reader, numSig int) (sigData []Rtcm31009SignalData) {
    for i := 0; i < numSig; i++ {
        sigData = append(sigData, Rtcm31009SignalData{
            SatelliteId: r.Uint8(6),
            CodeIndicator: r.Bit(),
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
