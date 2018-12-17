package rtcm

import (
    "github.com/bamiaux/iobit"
    "math"
    "math/bits"
    "time"
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
    Header Rtcm3MsmHeader
    SatelliteData Rtcm3SatelliteDataMsm57
    SignalData Rtcm3SignalDataMsm7
}

func NewRtcm3MessageMsm7(payload []byte) Rtcm3MessageMsm7 {
    r := iobit.NewReader(payload)
    header := NewRtcm3MsmHeader(&r)
    return Rtcm3MessageMsm7{
        Header: header,
        SatelliteData: NewRtcm3SatelliteDataMsm57(&r, bits.OnesCount(uint(header.SatelliteMask))),
        SignalData: NewRtcm3SignalDataMsm7(&r, bits.OnesCount(uint(header.CellMask))),
    }
}

func (msg Rtcm3MessageMsm7) Serialize() (data []byte) {
    satMaskBits := bits.OnesCount(uint(msg.Header.SatelliteMask))
    sigMaskBits := bits.OnesCount(uint(msg.Header.SignalMask))
    cellMaskBits := bits.OnesCount(uint(msg.Header.CellMask))
    msgBits := (169 + (satMaskBits * sigMaskBits)) + (36 * satMaskBits) + (80 * cellMaskBits)
    data = make([]byte, int(math.Ceil(float64(msgBits) / 8)))
    w := iobit.NewWriter(data)

    w.PutUint16(12, msg.Header.MessageNumber)
    w.PutUint16(12, msg.Header.StationId)
    w.PutUint32(30, msg.Header.Epoch)
    w.PutBit(msg.Header.MultipleMessageBit)
    w.PutUint8(3, msg.Header.Iods)
    w.PutUint8(7, msg.Header.Reserved)
    w.PutUint8(2, msg.Header.ClockSteeringIndicator)
    w.PutUint8(2, msg.Header.ExternalClockIndicator)
    w.PutBit(msg.Header.SmoothingIndicator)
    w.PutUint8(3, msg.Header.SmoothingInterval)
    w.PutUint64(64, msg.Header.SatelliteMask)
    w.PutUint32(32, msg.Header.SignalMask)
    w.PutUint64(uint(sigMaskBits * satMaskBits), msg.Header.CellMask)

    for _, rangeMillis := range msg.SatelliteData.RangeMilliseconds {
        w.PutUint8(8, rangeMillis)
    }
    for _, extended := range msg.SatelliteData.Extended {
        w.PutUint8(4, extended)
    }
    for _, ranges := range msg.SatelliteData.Ranges {
        w.PutUint16(10, ranges)
    }
    for _, phaseRangeRate := range msg.SatelliteData.PhaseRangeRates {
        w.PutInt16(14, phaseRangeRate)
    }

    for _, pseudorange := range msg.SignalData.Pseudoranges {
        w.PutInt32(20, pseudorange)
    }
    for _, phaseRange := range msg.SignalData.PhaseRanges {
        w.PutInt32(24, phaseRange)
    }
    for _, phaseRangeLock := range msg.SignalData.PhaseRangeLocks {
        w.PutUint16(10, phaseRangeLock)
    }
    for _, halfCycle := range msg.SignalData.HalfCycles {
        w.PutBit(halfCycle)
    }
    for _, cnr := range msg.SignalData.Cnrs {
        w.PutUint16(10, cnr)
    }
    for _, sigPhaseRangeRate := range msg.SignalData.PhaseRangeRates {
        w.PutInt16(15, sigPhaseRangeRate)
    }

    w.PutUint8(uint(w.Bits()), 0) // Pad with 0s - Should always be less than 1 byte, should check
    w.Flush()
    return data
}

func (msg Rtcm3MessageMsm7) Number() uint16 {
    return msg.Header.MessageNumber
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
    Header Rtcm3MsmHeader
    SatelliteData Rtcm3SatelliteDataMsm46
    SignalData Rtcm3SignalDataMsm6
}

func NewRtcm3MessageMsm6(payload []byte) Rtcm3MessageMsm6 {
    r := iobit.NewReader(payload)
    header := NewRtcm3MsmHeader(&r)
    return Rtcm3MessageMsm6{
        Header: header,
        SatelliteData: NewRtcm3SatelliteDataMsm46(&r, bits.OnesCount(uint(header.SatelliteMask))),
        SignalData: NewRtcm3SignalDataMsm6(&r, bits.OnesCount(uint(header.CellMask))),
    }
}

func (msg Rtcm3MessageMsm6) Number() uint16 {
    return msg.Header.MessageNumber
}

func (msg Rtcm3MessageMsm6) Serialize() (data []byte) {
    return data
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

func NewRtcm3MessageMsm5(f Rtcm3Frame) Rtcm3MessageMsm5 {
    r := iobit.NewReader(f.Payload)
    header := NewRtcm3MsmHeader(&r)
    return Rtcm3MessageMsm5{
        Rtcm3Frame: f,
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

func NewRtcm3MessageMsm4(f Rtcm3Frame) Rtcm3MessageMsm4 {
    r := iobit.NewReader(f.Payload)
    header := NewRtcm3MsmHeader(&r)
    return Rtcm3MessageMsm4{
        Rtcm3Frame: f,
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

func NewRtcm3MessageMsm3(f Rtcm3Frame) Rtcm3MessageMsm3 {
    r := iobit.NewReader(f.Payload)
    header := NewRtcm3MsmHeader(&r)
    return Rtcm3MessageMsm3{
        Rtcm3Frame: f,
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

func NewRtcm3MessageMsm2(f Rtcm3Frame) Rtcm3MessageMsm2 {
    r := iobit.NewReader(f.Payload)
    header := NewRtcm3MsmHeader(&r)
    return Rtcm3MessageMsm2{
        Rtcm3Frame: f,
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

func NewRtcm3MessageMsm1(f Rtcm3Frame) Rtcm3MessageMsm1 {
    r := iobit.NewReader(f.Payload)
    header := NewRtcm3MsmHeader(&r)
    return Rtcm3MessageMsm1{
        Rtcm3Frame: f,
        Header: header,
        SatelliteData: NewRtcm3SatelliteDataMsm123(&r, bits.OnesCount(uint(header.SatelliteMask))),
        SignalData: NewRtcm3SignalDataMsm1(&r, bits.OnesCount(uint(header.CellMask))),
    }
}

type Rtcm3Message1077 struct {
    Rtcm3MessageMsm7
}

func NewRtcm3Message1077(payload []byte) Rtcm3Message1077 {
    return Rtcm3Message1077{
        Rtcm3MessageMsm7: NewRtcm3MessageMsm7(payload),
    }
}

func (msg Rtcm3Message1077) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Rtcm3Message1087 struct {
    Rtcm3MessageMsm7
}

func NewRtcm3Message1087(f Rtcm3Frame) Rtcm3Message1087 {
    return Rtcm3Message1087{
        Rtcm3MessageMsm7: NewRtcm3MessageMsm7(f.Payload),
    }
}

func (msg Rtcm3Message1087) Time() time.Time {
    return GlonassTime(msg.Header.Epoch)
}

type Rtcm3Message1097 struct {
    Rtcm3MessageMsm7
}

func NewRtcm3Message1097(f Rtcm3Frame) Rtcm3Message1097 {
    return Rtcm3Message1097{
        Rtcm3MessageMsm7: NewRtcm3MessageMsm7(f.Payload),
    }
}

func (msg Rtcm3Message1097) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Rtcm3Message1117 struct {
    Rtcm3MessageMsm7
}

func NewRtcm3Message1117(f Rtcm3Frame) Rtcm3Message1117 {
    return Rtcm3Message1117{
        Rtcm3MessageMsm7: NewRtcm3MessageMsm7(f.Payload),
    }
}

func (msg Rtcm3Message1117) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Rtcm3Message1127 struct {
    Rtcm3MessageMsm7
}

func NewRtcm3Message1127(f Rtcm3Frame) Rtcm3Message1127 {
    return Rtcm3Message1127{
        Rtcm3MessageMsm7: NewRtcm3MessageMsm7(f.Payload),
    }
}

func (msg Rtcm3Message1127) Time() time.Time {
    return GpsTime(msg.Header.Epoch).Add(14 * time.Second)
}
