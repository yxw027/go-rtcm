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

type Rtcm3SatelliteDataMsm57 struct {
    RangeMilliseconds []uint8
    Extended []uint8
    Ranges []uint16
    PhaseRangeRates []int16
}

func NewRtcm3SatelliteDataMsm57(r *iobit.Reader, nsat int) (satData Rtcm3SatelliteDataMsm57) {
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

type Rtcm3SignalDataMsm7 struct {
    Pseudoranges []int32
    PhaseRanges []int32
    PhaseRangeLocks []uint16
    HalfCycles []bool
    Cnrs []uint16
    PhaseRangeRates []int16
}

func NewRtcm3SignalDataMsm7(r *iobit.Reader, ncell int) (sigData Rtcm3SignalDataMsm7) {
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

type Rtcm3MessageMsm7 struct {
    Rtcm3Frame
    Header Rtcm3MsmHeader
    SatelliteData Rtcm3SatelliteDataMsm57
    SignalData Rtcm3SignalDataMsm7
}

func NewRtcm3MessageMsm7(msg Rtcm3Frame) Rtcm3MessageMsm7 {
    r := iobit.NewReader(msg.Payload)
    header := NewRtcm3MsmHeader(&r)
    return Rtcm3MessageMsm7{
        Rtcm3Frame: msg,
        Header: header,
        SatelliteData: NewRtcm3SatelliteDataMsm57(&r, bits.OnesCount(uint(header.SatelliteMask))),
        SignalData: NewRtcm3SignalDataMsm7(&r, bits.OnesCount(uint(header.CellMask))),
    }
}

type Rtcm3SatelliteDataMsm46 struct {
    RangeMilliseconds []uint8
    Ranges []uint16
}

func NewRtcm3SatelliteDataMsm46(r *iobit.Reader, nsat int) (satData Rtcm3SatelliteDataMsm46) {
    for i := 0; i < nsat; i++ {
        satData.RangeMilliseconds = append(satData.RangeMilliseconds, r.Uint8(8))
    }
    for i := 0; i < nsat; i++ {
        satData.Ranges = append(satData.Ranges, r.Uint16(10))
    }
    return satData
}

type Rtcm3SignalDataMsm6 struct {
    Pseudoranges []int32
    PhaseRanges []int32
    PhaseRangeLocks []uint16
    HalfCycles []bool
    Cnrs []uint16
}

func NewRtcm3SignalDataMsm6(r *iobit.Reader, ncell int) (sigData Rtcm3SignalDataMsm6) {
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

type Rtcm3MessageMsm6 struct {
    Rtcm3Frame
    Header Rtcm3MsmHeader
    SatelliteData Rtcm3SatelliteDataMsm46
    SignalData Rtcm3SignalDataMsm6
}

func NewRtcm3MessageMsm6(msg Rtcm3Frame) Rtcm3MessageMsm6 {
    r := iobit.NewReader(msg.Payload)
    header := NewRtcm3MsmHeader(&r)
    return Rtcm3MessageMsm6{
        Rtcm3Frame: msg,
        Header: header,
        SatelliteData: NewRtcm3SatelliteDataMsm46(&r, bits.OnesCount(uint(header.SatelliteMask))),
        SignalData: NewRtcm3SignalDataMsm6(&r, bits.OnesCount(uint(header.CellMask))),
    }
}

type Rtcm3SignalDataMsm5 struct {
    Pseudoranges []int16
    PhaseRanges []int32
    PhaseRangeLocks []uint8
    HalfCycles []bool
    Cnrs []uint8
    PhaseRangeRates []int16
}

func NewRtcm3SignalDataMsm5(r *iobit.Reader, ncell int) (sigData Rtcm3SignalDataMsm5) {
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

type Rtcm3MessageMsm5 struct {
    Rtcm3Frame
    Header Rtcm3MsmHeader
    SatelliteData Rtcm3SatelliteDataMsm57
    SignalData Rtcm3SignalDataMsm5
}

func NewRtcm3MessageMsm5(msg Rtcm3Frame) Rtcm3MessageMsm5 {
    r := iobit.NewReader(msg.Payload)
    header := NewRtcm3MsmHeader(&r)
    return Rtcm3MessageMsm5{
        Rtcm3Frame: msg,
        Header: header,
        SatelliteData: NewRtcm3SatelliteDataMsm57(&r, bits.OnesCount(uint(header.SatelliteMask))),
        SignalData: NewRtcm3SignalDataMsm5(&r, bits.OnesCount(uint(header.CellMask))),
    }
}

type Rtcm3SignalDataMsm4 struct {
    Pseudoranges []int16
    PhaseRanges []int32
    PhaseRangeLocks []uint8
    HalfCycles []bool
    Cnrs []uint8
}

func NewRtcm3SignalDataMsm4(r *iobit.Reader, ncell int) (sigData Rtcm3SignalDataMsm4) {
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

type Rtcm3MessageMsm4 struct {
    Rtcm3Frame
    Header Rtcm3MsmHeader
    SatelliteData Rtcm3SatelliteDataMsm46
    SignalData Rtcm3SignalDataMsm4
}

func NewRtcm3MessageMsm4(msg Rtcm3Frame) Rtcm3MessageMsm4 {
    r := iobit.NewReader(msg.Payload)
    header := NewRtcm3MsmHeader(&r)
    return Rtcm3MessageMsm4{
        Rtcm3Frame: msg,
        Header: header,
        SatelliteData: NewRtcm3SatelliteDataMsm46(&r, bits.OnesCount(uint(header.SatelliteMask))),
        SignalData: NewRtcm3SignalDataMsm4(&r, bits.OnesCount(uint(header.CellMask))),
    }
}

type Rtcm3SatelliteDataMsm123 struct {
    Ranges []uint16
}

func NewRtcm3SatelliteDataMsm123(r *iobit.Reader, nsat int) (satData Rtcm3SatelliteDataMsm123) {
    for i := 0; i < nsat; i++ {
        satData.Ranges = append(satData.Ranges, r.Uint16(10))
    }
    return satData
}

type Rtcm3SignalDataMsm3 struct {
    Pseudoranges []int16
    PhaseRanges []int32
    PhaseRangeLocks []uint8
    HalfCycles []bool
}

func NewRtcm3SignalDataMsm3(r *iobit.Reader, ncell int) (sigData Rtcm3SignalDataMsm3) {
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

type Rtcm3MessageMsm3 struct {
    Rtcm3Frame
    Header Rtcm3MsmHeader
    SatelliteData Rtcm3SatelliteDataMsm123
    SignalData Rtcm3SignalDataMsm3
}

func NewRtcm3MessageMsm3(msg Rtcm3Frame) Rtcm3MessageMsm3 {
    r := iobit.NewReader(msg.Payload)
    header := NewRtcm3MsmHeader(&r)
    return Rtcm3MessageMsm3{
        Rtcm3Frame: msg,
        Header: header,
        SatelliteData: NewRtcm3SatelliteDataMsm123(&r, bits.OnesCount(uint(header.SatelliteMask))),
        SignalData: NewRtcm3SignalDataMsm3(&r, bits.OnesCount(uint(header.CellMask))),
    }
}

type Rtcm3SignalDataMsm2 struct {
    PhaseRanges []int32
    PhaseRangeLocks []uint8
    HalfCycles []bool
}

func NewRtcm3SignalDataMsm2(r *iobit.Reader, ncell int) (sigData Rtcm3SignalDataMsm2) {
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

type Rtcm3MessageMsm2 struct {
    Rtcm3Frame
    Header Rtcm3MsmHeader
    SatelliteData Rtcm3SatelliteDataMsm123
    SignalData Rtcm3SignalDataMsm2
}

func NewRtcm3MessageMsm2(msg Rtcm3Frame) Rtcm3MessageMsm2 {
    r := iobit.NewReader(msg.Payload)
    header := NewRtcm3MsmHeader(&r)
    return Rtcm3MessageMsm2{
        Rtcm3Frame: msg,
        Header: header,
        SatelliteData: NewRtcm3SatelliteDataMsm123(&r, bits.OnesCount(uint(header.SatelliteMask))),
        SignalData: NewRtcm3SignalDataMsm2(&r, bits.OnesCount(uint(header.CellMask))),
    }
}

type Rtcm3SignalDataMsm1 struct {
    Pseudoranges []int16
}

func NewRtcm3SignalDataMsm1(r *iobit.Reader, ncell int) (sigData Rtcm3SignalDataMsm1) {
    for i := 0; i < ncell; i++ {
        sigData.Pseudoranges = append(sigData.Pseudoranges, r.Int16(15))
    }
    return sigData
}

type Rtcm3MessageMsm1 struct {
    Rtcm3Frame
    Header Rtcm3MsmHeader
    SatelliteData Rtcm3SatelliteDataMsm123
    SignalData Rtcm3SignalDataMsm1
}

func NewRtcm3MessageMsm1(msg Rtcm3Frame) Rtcm3MessageMsm1 {
    r := iobit.NewReader(msg.Payload)
    header := NewRtcm3MsmHeader(&r)
    return Rtcm3MessageMsm1{
        Rtcm3Frame: msg,
        Header: header,
        SatelliteData: NewRtcm3SatelliteDataMsm123(&r, bits.OnesCount(uint(header.SatelliteMask))),
        SignalData: NewRtcm3SignalDataMsm1(&r, bits.OnesCount(uint(header.CellMask))),
    }
}

// Presumably will need seperate types for the MSM7 messages eventually
//type Rtcm31077 struct {
//    Rtcm3Frame
//    Rtcm3MessageMsm7
//}
//
//func NewRtcm31077(msg Rtcm3Frame) Rtcm31077 {
//    return Rtcm31077{
//        Rtcm3Frame: msg,
//        Rtcm3MessageMsm7: NewRtcm3MessageMsm7(msg.Payload),
//    }
//}
//
//type Rtcm31087 struct {
//    Rtcm3Frame
//    Rtcm3MessageMsm7
//}
//
//func NewRtcm31087(msg Rtcm3Frame) Rtcm31087 {
//    return Rtcm31087{
//        Rtcm3Frame: msg,
//        Rtcm3MessageMsm7: NewRtcm3MessageMsm7(msg.Payload),
//    }
//}
//
//type Rtcm31097 struct {
//    Rtcm3Frame
//    Rtcm3MessageMsm7
//}
//
//func NewRtcm31097(msg Rtcm3Frame) Rtcm31097 {
//    return Rtcm31097{
//        Rtcm3Frame: msg,
//        Rtcm3MessageMsm7: NewRtcm3MessageMsm7(msg.Payload),
//    }
//}
//
//type Rtcm31117 struct {
//    Rtcm3Frame
//    Rtcm3MessageMsm7
//}
//
//func NewRtcm31117(msg Rtcm3Frame) Rtcm31117 {
//    return Rtcm31117{
//        Rtcm3Frame: msg,
//        Rtcm3MessageMsm7: NewRtcm3MessageMsm7(msg.Payload),
//    }
//}
//
//type Rtcm31127 struct {
//    Rtcm3Frame
//    Rtcm3MessageMsm7
//}
//
//func NewRtcm31127(msg Rtcm3Frame) Rtcm31127 {
//    return Rtcm31127{
//        Rtcm3Frame: msg,
//        Rtcm3MessageMsm7: NewRtcm3MessageMsm7(msg.Payload),
//    }
//}
