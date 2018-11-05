package rtcm

import (
    "github.com/bamiaux/iobit"
    "math/bits"
)

type Rtcm3MsmHeader struct {
    MessageNumber uint16
    StationId uint16
    Epoch uint32
    MultipleMessageBit bool
    Iods uint8
    Reserved uint8
    ClockSteeringIndicator uint8
    ExternalClockIndicator uint8
    SmoothingIndicator bool
    SmoothingInterval uint8
    SatelliteMask uint64
    SignalMask uint32
    CellMask uint64
}

func NewRtcm3MsmHeader(r *iobit.Reader) (header Rtcm3MsmHeader) {
    header = Rtcm3MsmHeader{
        MessageNumber: r.Uint16(12),
        StationId: r.Uint16(12),
        Epoch: r.Uint32(30),
        MultipleMessageBit: r.Bit(),
        Iods: r.Uint8(3),
        Reserved: r.Uint8(7),
        ClockSteeringIndicator: r.Uint8(2),
        ExternalClockIndicator: r.Uint8(2),
        SmoothingIndicator: r.Bit(),
        SmoothingInterval: r.Uint8(3),
        SatelliteMask: r.Uint64(64),
        SignalMask: r.Uint32(32),
    }

    cellMaskLength := bits.OnesCount(uint(header.SignalMask)) * bits.OnesCount(uint(header.SatelliteMask))
    header.CellMask = r.Uint64(uint(cellMaskLength))
    return header
}

type Rtcm3Msm57SatelliteData struct {
    RangeMilliseconds []uint8
    Extended []uint8
    Ranges []uint16
    PhaseRangeRates []int16
}

func NewRtcm3Msm57SatelliteData(r *iobit.Reader, nsat int) (satData Rtcm3Msm57SatelliteData) {
    for i := 0; i < nsat; i++ {
        satData.RangeMilliseconds = append(satData.RangeMilliseconds, r.Uint8(8))
    }
    for i := 0; i < nsat; i++ {
        satData.Extended = append(satData.Extended, r.Uint8(4))
    }
    for i := 0; i < nsat; i++ {
        satData.Ranges = append(satData.Ranges, r.Uint16(10))
    }
    for i := 0; i < nsat; i++ {
        satData.PhaseRangeRates = append(satData.PhaseRangeRates, r.Int16(14))
    }
    return satData
}

type Rtcm3Msm7SignalData struct {
    Pseudoranges []int32
    PhaseRanges []int32
    PhaseRangeLocks []uint16
    HalfCycles []bool
    Cnrs []uint16
    PhaseRangeRates []int16
}

func NewRtcm3Msm7SignalData(r *iobit.Reader, ncell int) (sigData Rtcm3Msm7SignalData) {
    for i := 0; i < ncell; i++ {
        sigData.Pseudoranges = append(sigData.Pseudoranges, r.Int32(20))
    }
    for i := 0; i < ncell; i++ {
        sigData.PhaseRanges = append(sigData.PhaseRanges, r.Int32(24))
    }
    for i := 0; i < ncell; i++ {
        sigData.PhaseRangeLocks = append(sigData.PhaseRangeLocks, r.Uint16(10))
    }
    for i := 0; i < ncell; i++ {
        sigData.HalfCycles = append(sigData.HalfCycles, r.Bit())
    }
    for i := 0; i < ncell; i++ {
        sigData.Cnrs = append(sigData.Cnrs, r.Uint16(10))
    }
    for i := 0; i < ncell; i++ {
        sigData.PhaseRangeRates = append(sigData.PhaseRangeRates, r.Int16(15))
    }
    return sigData
}

type Rtcm3Msm7Message struct {
    Rtcm3Frame
    Header Rtcm3MsmHeader
    SatelliteData Rtcm3Msm57SatelliteData
    SignalData Rtcm3Msm7SignalData
}

func NewRtcm3Msm7Message(msg Rtcm3Frame) Rtcm3Msm7Message {
    r := iobit.NewReader(msg.Payload)
    header := NewRtcm3MsmHeader(&r)
    return Rtcm3Msm7Message{
        Rtcm3Frame: msg,
        Header: header,
        SatelliteData: NewRtcm3Msm57SatelliteData(&r, bits.OnesCount(uint(header.SatelliteMask))),
        SignalData: NewRtcm3Msm7SignalData(&r, bits.OnesCount(uint(header.CellMask))),
    }
}

type Rtcm3Msm46SatelliteData struct {
    RangeMilliseconds []uint8
    Ranges []uint16
}

func NewRtcm3Msm46SatelliteData(r *iobit.Reader, nsat int) (satData Rtcm3Msm46SatelliteData) {
    for i := 0; i < nsat; i++ {
        satData.RangeMilliseconds = append(satData.RangeMilliseconds, r.Uint8(8))
    }
    for i := 0; i < nsat; i++ {
        satData.Ranges = append(satData.Ranges, r.Uint16(10))
    }
    return satData
}

type Rtcm3Msm6SignalData struct {
    Pseudoranges []int32
    PhaseRanges []int32
    PhaseRangeLocks []uint16
    HalfCycles []bool
    Cnrs []uint16
}

func NewRtcm3Msm6SignalData(r *iobit.Reader, ncell int) (sigData Rtcm3Msm6SignalData) {
    for i := 0; i < ncell; i++ {
        sigData.Pseudoranges = append(sigData.Pseudoranges, r.Int32(20))
    }
    for i := 0; i < ncell; i++ {
        sigData.PhaseRanges = append(sigData.PhaseRanges, r.Int32(24))
    }
    for i := 0; i < ncell; i++ {
        sigData.PhaseRangeLocks = append(sigData.PhaseRangeLocks, r.Uint16(10))
    }
    for i := 0; i < ncell; i++ {
        sigData.HalfCycles = append(sigData.HalfCycles, r.Bit())
    }
    for i := 0; i < ncell; i++ {
        sigData.Cnrs = append(sigData.Cnrs, r.Uint16(10))
    }
    return sigData
}

type Rtcm3Msm6Message struct {
    Rtcm3Frame
    Header Rtcm3MsmHeader
    SatelliteData Rtcm3Msm46SatelliteData
    SignalData Rtcm3Msm6SignalData
}

func NewRtcm3Msm6Message(msg Rtcm3Frame) Rtcm3Msm6Message {
    r := iobit.NewReader(msg.Payload)
    header := NewRtcm3MsmHeader(&r)
    return Rtcm3Msm6Message{
        Rtcm3Frame: msg,
        Header: header,
        SatelliteData: NewRtcm3Msm46SatelliteData(&r, bits.OnesCount(uint(header.SatelliteMask))),
        SignalData: NewRtcm3Msm6SignalData(&r, bits.OnesCount(uint(header.CellMask))),
    }
}

type Rtcm3Msm5SignalData struct {
    Pseudoranges []int16
    PhaseRanges []int32
    PhaseRangeLocks []uint8
    HalfCycles []bool
    Cnrs []uint8
    PhaseRangeRates []int16
}

func NewRtcm3Msm5SignalData(r *iobit.Reader, ncell int) (sigData Rtcm3Msm5SignalData) {
    for i := 0; i < ncell; i++ {
        sigData.Pseudoranges = append(sigData.Pseudoranges, r.Int16(15))
    }
    for i := 0; i < ncell; i++ {
        sigData.PhaseRanges = append(sigData.PhaseRanges, r.Int32(22))
    }
    for i := 0; i < ncell; i++ {
        sigData.PhaseRangeLocks = append(sigData.PhaseRangeLocks, r.Uint8(4))
    }
    for i := 0; i < ncell; i++ {
        sigData.HalfCycles = append(sigData.HalfCycles, r.Bit())
    }
    for i := 0; i < ncell; i++ {
        sigData.Cnrs = append(sigData.Cnrs, r.Uint8(6))
    }
    for i := 0; i < ncell; i++ {
        sigData.PhaseRangeRates = append(sigData.PhaseRangeRates, r.Int16(15))
    }
    return sigData
}

type Rtcm3Msm5Message struct {
    Rtcm3Frame
    Header Rtcm3MsmHeader
    SatelliteData Rtcm3Msm57SatelliteData
    SignalData Rtcm3Msm5SignalData
}

func NewRtcm3Msm5Message(msg Rtcm3Frame) Rtcm3Msm5Message {
    r := iobit.NewReader(msg.Payload)
    header := NewRtcm3MsmHeader(&r)
    return Rtcm3Msm5Message{
        Rtcm3Frame: msg,
        Header: header,
        SatelliteData: NewRtcm3Msm57SatelliteData(&r, bits.OnesCount(uint(header.SatelliteMask))),
        SignalData: NewRtcm3Msm5SignalData(&r, bits.OnesCount(uint(header.CellMask))),
    }
}

type Rtcm3Msm4SignalData struct {
    Pseudoranges []int16
    PhaseRanges []int32
    PhaseRangeLocks []uint8
    HalfCycles []bool
    Cnrs []uint8
}

func NewRtcm3Msm4SignalData(r *iobit.Reader, ncell int) (sigData Rtcm3Msm4SignalData) {
    for i := 0; i < ncell; i++ {
        sigData.Pseudoranges = append(sigData.Pseudoranges, r.Int16(15))
    }
    for i := 0; i < ncell; i++ {
        sigData.PhaseRanges = append(sigData.PhaseRanges, r.Int32(22))
    }
    for i := 0; i < ncell; i++ {
        sigData.PhaseRangeLocks = append(sigData.PhaseRangeLocks, r.Uint8(4))
    }
    for i := 0; i < ncell; i++ {
        sigData.HalfCycles = append(sigData.HalfCycles, r.Bit())
    }
    for i := 0; i < ncell; i++ {
        sigData.Cnrs = append(sigData.Cnrs, r.Uint8(6))
    }
    return sigData
}

type Rtcm3Msm4Message struct {
    Rtcm3Frame
    Header Rtcm3MsmHeader
    SatelliteData Rtcm3Msm46SatelliteData
    SignalData Rtcm3Msm4SignalData
}

func NewRtcm3Msm4Message(msg Rtcm3Frame) Rtcm3Msm4Message {
    r := iobit.NewReader(msg.Payload)
    header := NewRtcm3MsmHeader(&r)
    return Rtcm3Msm4Message{
        Rtcm3Frame: msg,
        Header: header,
        SatelliteData: NewRtcm3Msm46SatelliteData(&r, bits.OnesCount(uint(header.SatelliteMask))),
        SignalData: NewRtcm3Msm4SignalData(&r, bits.OnesCount(uint(header.CellMask))),
    }
}

type Rtcm3Msm123SatelliteData struct {
    Ranges []uint16
}

func NewRtcm3Msm123SatelliteData(r *iobit.Reader, nsat int) (satData Rtcm3Msm123SatelliteData) {
    for i := 0; i < nsat; i++ {
        satData.Ranges = append(satData.Ranges, r.Uint16(10))
    }
    return satData
}

type Rtcm3Msm3SignalData struct {
    Pseudoranges []int16
    PhaseRanges []int32
    PhaseRangeLocks []uint8
    HalfCycles []bool
}

func NewRtcm3Msm3SignalData(r *iobit.Reader, ncell int) (sigData Rtcm3Msm3SignalData) {
    for i := 0; i < ncell; i++ {
        sigData.Pseudoranges = append(sigData.Pseudoranges, r.Int16(15))
    }
    for i := 0; i < ncell; i++ {
        sigData.PhaseRanges = append(sigData.PhaseRanges, r.Int32(22))
    }
    for i := 0; i < ncell; i++ {
        sigData.PhaseRangeLocks = append(sigData.PhaseRangeLocks, r.Uint8(4))
    }
    for i := 0; i < ncell; i++ {
        sigData.HalfCycles = append(sigData.HalfCycles, r.Bit())
    }
    return sigData
}

type Rtcm3Msm3Message struct {
    Rtcm3Frame
    Header Rtcm3MsmHeader
    SatelliteData Rtcm3Msm123SatelliteData
    SignalData Rtcm3Msm3SignalData
}

func NewRtcm3Msm3Message(msg Rtcm3Frame) Rtcm3Msm3Message {
    r := iobit.NewReader(msg.Payload)
    header := NewRtcm3MsmHeader(&r)
    return Rtcm3Msm3Message{
        Rtcm3Frame: msg,
        Header: header,
        SatelliteData: NewRtcm3Msm123SatelliteData(&r, bits.OnesCount(uint(header.SatelliteMask))),
        SignalData: NewRtcm3Msm3SignalData(&r, bits.OnesCount(uint(header.CellMask))),
    }
}

type Rtcm3Msm2SignalData struct {
    PhaseRanges []int32
    PhaseRangeLocks []uint8
    HalfCycles []bool
}

func NewRtcm3Msm2SignalData(r *iobit.Reader, ncell int) (sigData Rtcm3Msm2SignalData) {
    for i := 0; i < ncell; i++ {
        sigData.PhaseRanges = append(sigData.PhaseRanges, r.Int32(22))
    }
    for i := 0; i < ncell; i++ {
        sigData.PhaseRangeLocks = append(sigData.PhaseRangeLocks, r.Uint8(4))
    }
    for i := 0; i < ncell; i++ {
        sigData.HalfCycles = append(sigData.HalfCycles, r.Bit())
    }
    return sigData
}

type Rtcm3Msm2Message struct {
    Rtcm3Frame
    Header Rtcm3MsmHeader
    SatelliteData Rtcm3Msm123SatelliteData
    SignalData Rtcm3Msm2SignalData
}

func NewRtcm3Msm2Message(msg Rtcm3Frame) Rtcm3Msm2Message {
    r := iobit.NewReader(msg.Payload)
    header := NewRtcm3MsmHeader(&r)
    return Rtcm3Msm2Message{
        Rtcm3Frame: msg,
        Header: header,
        SatelliteData: NewRtcm3Msm123SatelliteData(&r, bits.OnesCount(uint(header.SatelliteMask))),
        SignalData: NewRtcm3Msm2SignalData(&r, bits.OnesCount(uint(header.CellMask))),
    }
}

type Rtcm3Msm1SignalData struct {
    Pseudoranges []int16
}

func NewRtcm3Msm1SignalData(r *iobit.Reader, ncell int) (sigData Rtcm3Msm1SignalData) {
    for i := 0; i < ncell; i++ {
        sigData.Pseudoranges = append(sigData.Pseudoranges, r.Int16(15))
    }
    return sigData
}

type Rtcm3Msm1Message struct {
    Rtcm3Frame
    Header Rtcm3MsmHeader
    SatelliteData Rtcm3Msm123SatelliteData
    SignalData Rtcm3Msm1SignalData
}

func NewRtcm3Msm1Message(msg Rtcm3Frame) Rtcm3Msm1Message {
    r := iobit.NewReader(msg.Payload)
    header := NewRtcm3MsmHeader(&r)
    return Rtcm3Msm1Message{
        Rtcm3Frame: msg,
        Header: header,
        SatelliteData: NewRtcm3Msm123SatelliteData(&r, bits.OnesCount(uint(header.SatelliteMask))),
        SignalData: NewRtcm3Msm1SignalData(&r, bits.OnesCount(uint(header.CellMask))),
    }
}

// Presumably will need seperate types for the MSM7 messages eventually
//type Rtcm31077 struct {
//    Rtcm3Frame
//    Rtcm3Msm7Message
//}
//
//func NewRtcm31077(msg Rtcm3Frame) Rtcm31077 {
//    return Rtcm31077{
//        Rtcm3Frame: msg,
//        Rtcm3Msm7Message: NewRtcm3Msm7Message(msg.Payload),
//    }
//}
//
//type Rtcm31087 struct {
//    Rtcm3Frame
//    Rtcm3Msm7Message
//}
//
//func NewRtcm31087(msg Rtcm3Frame) Rtcm31087 {
//    return Rtcm31087{
//        Rtcm3Frame: msg,
//        Rtcm3Msm7Message: NewRtcm3Msm7Message(msg.Payload),
//    }
//}
//
//type Rtcm31097 struct {
//    Rtcm3Frame
//    Rtcm3Msm7Message
//}
//
//func NewRtcm31097(msg Rtcm3Frame) Rtcm31097 {
//    return Rtcm31097{
//        Rtcm3Frame: msg,
//        Rtcm3Msm7Message: NewRtcm3Msm7Message(msg.Payload),
//    }
//}
//
//type Rtcm31117 struct {
//    Rtcm3Frame
//    Rtcm3Msm7Message
//}
//
//func NewRtcm31117(msg Rtcm3Frame) Rtcm31117 {
//    return Rtcm31117{
//        Rtcm3Frame: msg,
//        Rtcm3Msm7Message: NewRtcm3Msm7Message(msg.Payload),
//    }
//}
//
//type Rtcm31127 struct {
//    Rtcm3Frame
//    Rtcm3Msm7Message
//}
//
//func NewRtcm31127(msg Rtcm3Frame) Rtcm31127 {
//    return Rtcm31127{
//        Rtcm3Frame: msg,
//        Rtcm3Msm7Message: NewRtcm3Msm7Message(msg.Payload),
//    }
//}
