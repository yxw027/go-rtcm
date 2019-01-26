package rtcm3

import (
    "github.com/bamiaux/iobit"
    "math"
    "math/bits"
    "time"
)

type MsmHeader struct {
    MessageNumber uint16
    ReferenceStationId uint16
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

func NewMsmHeader(r *iobit.Reader) (header MsmHeader) {
    header = MsmHeader{
        MessageNumber: r.Uint16(12),
        ReferenceStationId: r.Uint16(12),
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

func SerializeMsmHeader(w *iobit.Writer, header MsmHeader) {
    w.PutUint16(12, header.MessageNumber)
    w.PutUint16(12, header.ReferenceStationId)
    w.PutUint32(30, header.Epoch)
    w.PutBit(header.MultipleMessageBit)
    w.PutUint8(3, header.Iods)
    w.PutUint8(7, header.Reserved)
    w.PutUint8(2, header.ClockSteeringIndicator)
    w.PutUint8(2, header.ExternalClockIndicator)
    w.PutBit(header.SmoothingIndicator)
    w.PutUint8(3, header.SmoothingInterval)
    w.PutUint64(64, header.SatelliteMask)
    w.PutUint32(32, header.SignalMask)
    w.PutUint64(uint(bits.OnesCount(uint(header.SignalMask)) *  bits.OnesCount(uint(header.SatelliteMask))), header.CellMask)
    return
}

type SatelliteDataMsm57 struct {
    RangeMilliseconds []uint8
    Extended []uint8
    Ranges []uint16
    PhaseRangeRates []int16
}

func DeserializeSatelliteDataMsm57(r *iobit.Reader, nsat int) (satData SatelliteDataMsm57) {
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

func SerializeSatelliteDataMsm57(w *iobit.Writer, satelliteData SatelliteDataMsm57) {
    for _, rangeMillis := range satelliteData.RangeMilliseconds {
        w.PutUint8(8, rangeMillis)
    }
    for _, extended := range satelliteData.Extended {
        w.PutUint8(4, extended)
    }
    for _, ranges := range satelliteData.Ranges {
        w.PutUint16(10, ranges)
    }
    for _, phaseRangeRate := range satelliteData.PhaseRangeRates {
        w.PutInt16(14, phaseRangeRate)
    }
}

type SignalDataMsm7 struct {
    Pseudoranges []int32
    PhaseRanges []int32
    PhaseRangeLocks []uint16
    HalfCycles []bool
    Cnrs []uint16
    PhaseRangeRates []int16
}

func DeserializeSignalDataMsm7(r *iobit.Reader, ncell int) (sigData SignalDataMsm7) {
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

type MessageMsm7 struct {
    Header MsmHeader
    SatelliteData SatelliteDataMsm57
    SignalData SignalDataMsm7
}

func (msg MessageMsm7) Number() uint16 {
    return msg.Header.MessageNumber
}

func DeserializeMessageMsm7(payload []byte) MessageMsm7 {
    r := iobit.NewReader(payload)
    msmHeader := NewMsmHeader(&r)
    return MessageMsm7{
        Header: msmHeader,
        SatelliteData: DeserializeSatelliteDataMsm57(&r, bits.OnesCount(uint(msmHeader.SatelliteMask))),
        SignalData: DeserializeSignalDataMsm7(&r, bits.OnesCount(uint(msmHeader.CellMask))),
    }
}

func (msg MessageMsm7) Serialize() (data []byte) {
    satMaskBits := bits.OnesCount(uint(msg.Header.SatelliteMask))
    sigMaskBits := bits.OnesCount(uint(msg.Header.SignalMask))
    cellMaskBits := bits.OnesCount(uint(msg.Header.CellMask))

    msgBits := (169 + (satMaskBits * sigMaskBits)) + (36 * satMaskBits) + (80 * cellMaskBits)
    data = make([]byte, int(math.Ceil(float64(msgBits) / 8)))
    w := iobit.NewWriter(data)

    SerializeMsmHeader(&w, msg.Header)
    SerializeSatelliteDataMsm57(&w, msg.SatelliteData)

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

type SatelliteDataMsm46 struct {
    RangeMilliseconds []uint8
    Ranges []uint16
}

func DeserializeSatelliteDataMsm46(r *iobit.Reader, nsat int) (satData SatelliteDataMsm46) {
    for i := 0; i < nsat; i++ {
        satData.RangeMilliseconds = append(satData.RangeMilliseconds, r.Uint8(8))
    }
    for i := 0; i < nsat; i++ {
        satData.Ranges = append(satData.Ranges, r.Uint16(10))
    }
    return satData
}

func SerializeSatelliteDataMsm46(w *iobit.Writer, satelliteData SatelliteDataMsm46) {
    for _, rangeMillis := range satelliteData.RangeMilliseconds {
        w.PutUint8(8, rangeMillis)
    }
    for _, ranges := range satelliteData.Ranges {
        w.PutUint16(10, ranges)
    }
}

type SignalDataMsm6 struct {
    Pseudoranges []int32
    PhaseRanges []int32
    PhaseRangeLocks []uint16
    HalfCycles []bool
    Cnrs []uint16
}

func DeserializeSignalDataMsm6(r *iobit.Reader, ncell int) (sigData SignalDataMsm6) {
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

type MessageMsm6 struct {
    Header MsmHeader
    SatelliteData SatelliteDataMsm46
    SignalData SignalDataMsm6
}

func (msg MessageMsm6) Number() uint16 {
    return msg.Header.MessageNumber
}

func DeserializeMessageMsm6(payload []byte) MessageMsm6 {
    r := iobit.NewReader(payload)
    msmHeader := NewMsmHeader(&r)
    return MessageMsm6{
        Header: msmHeader,
        SatelliteData: DeserializeSatelliteDataMsm46(&r, bits.OnesCount(uint(msmHeader.SatelliteMask))),
        SignalData: DeserializeSignalDataMsm6(&r, bits.OnesCount(uint(msmHeader.CellMask))),
    }
}

func (msg MessageMsm6) Serialize() (data []byte) {
    satMaskBits := bits.OnesCount(uint(msg.Header.SatelliteMask))
    sigMaskBits := bits.OnesCount(uint(msg.Header.SignalMask))
    cellMaskBits := bits.OnesCount(uint(msg.Header.CellMask))

    msgBits := (169 + (satMaskBits * sigMaskBits)) + (18 * satMaskBits) + (65 * cellMaskBits)
    data = make([]byte, int(math.Ceil(float64(msgBits) / 8)))
    w := iobit.NewWriter(data)

    SerializeMsmHeader(&w, msg.Header)
    SerializeSatelliteDataMsm46(&w, msg.SatelliteData)

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

    w.PutUint8(uint(w.Bits()), 0) // Pad with 0s - Should always be less than 1 byte, should check
    w.Flush()
    return data
}

type SignalDataMsm5 struct {
    Pseudoranges []int16
    PhaseRanges []int32
    PhaseRangeLocks []uint8
    HalfCycles []bool
    Cnrs []uint8
    PhaseRangeRates []int16
}

func DeserializeSignalDataMsm5(r *iobit.Reader, ncell int) (sigData SignalDataMsm5) {
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

type MessageMsm5 struct {
    Header MsmHeader
    SatelliteData SatelliteDataMsm57
    SignalData SignalDataMsm5
}

func (msg MessageMsm5) Number() uint16 {
    return msg.Header.MessageNumber
}

func DeserializeMessageMsm5(data []byte) MessageMsm5 {
    r := iobit.NewReader(data)
    msmHeader := NewMsmHeader(&r)
    return MessageMsm5{
        Header: msmHeader,
        SatelliteData: DeserializeSatelliteDataMsm57(&r, bits.OnesCount(uint(msmHeader.SatelliteMask))),
        SignalData: DeserializeSignalDataMsm5(&r, bits.OnesCount(uint(msmHeader.CellMask))),
    }
}

func (msg MessageMsm5) Serialize() (data []byte) {
    satMaskBits := bits.OnesCount(uint(msg.Header.SatelliteMask))
    sigMaskBits := bits.OnesCount(uint(msg.Header.SignalMask))
    cellMaskBits := bits.OnesCount(uint(msg.Header.CellMask))

    msgBits := (169 + (satMaskBits * sigMaskBits)) + (36 * satMaskBits) + (65 * cellMaskBits)
    data = make([]byte, int(math.Ceil(float64(msgBits) / 8)))
    w := iobit.NewWriter(data)

    SerializeMsmHeader(&w, msg.Header)
    SerializeSatelliteDataMsm57(&w, msg.SatelliteData)

    for _, pseudorange := range msg.SignalData.Pseudoranges {
        w.PutInt16(15, pseudorange)
    }
    for _, phaseRange := range msg.SignalData.PhaseRanges {
        w.PutInt32(22, phaseRange)
    }
    for _, phaseRangeLock := range msg.SignalData.PhaseRangeLocks {
        w.PutUint8(4, phaseRangeLock)
    }
    for _, halfCycle := range msg.SignalData.HalfCycles {
        w.PutBit(halfCycle)
    }
    for _, cnr := range msg.SignalData.Cnrs {
        w.PutUint8(6, cnr)
    }
    for _, sigPhaseRangeRate := range msg.SignalData.PhaseRangeRates {
        w.PutInt16(15, sigPhaseRangeRate)
    }

    w.PutUint8(uint(w.Bits()), 0) // Pad with 0s - Should always be less than 1 byte, should check
    w.Flush()
    return data
}

type SignalDataMsm4 struct {
    Pseudoranges []int16
    PhaseRanges []int32
    PhaseRangeLocks []uint8
    HalfCycles []bool
    Cnrs []uint8
}

func DeserializeSignalDataMsm4(r *iobit.Reader, ncell int) (sigData SignalDataMsm4) {
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

type MessageMsm4 struct {
    Header MsmHeader
    SatelliteData SatelliteDataMsm46
    SignalData SignalDataMsm4
}

func (msg MessageMsm4) Number() uint16 {
    return msg.Header.MessageNumber
}

func DeserializeMessageMsm4(data []byte) MessageMsm4 {
    r := iobit.NewReader(data)
    msmHeader := NewMsmHeader(&r)
    return MessageMsm4{
        Header: msmHeader,
        SatelliteData: DeserializeSatelliteDataMsm46(&r, bits.OnesCount(uint(msmHeader.SatelliteMask))),
        SignalData: DeserializeSignalDataMsm4(&r, bits.OnesCount(uint(msmHeader.CellMask))),
    }
}

func (msg MessageMsm4) Serialize() (data []byte) {
    satMaskBits := bits.OnesCount(uint(msg.Header.SatelliteMask))
    sigMaskBits := bits.OnesCount(uint(msg.Header.SignalMask))
    cellMaskBits := bits.OnesCount(uint(msg.Header.CellMask))

    msgBits := (169 + (satMaskBits * sigMaskBits)) + (18 * satMaskBits) + (48 * cellMaskBits)
    data = make([]byte, int(math.Ceil(float64(msgBits) / 8)))
    w := iobit.NewWriter(data)

    SerializeMsmHeader(&w, msg.Header)
    SerializeSatelliteDataMsm46(&w, msg.SatelliteData)

    for _, pseudorange := range msg.SignalData.Pseudoranges {
        w.PutInt16(15, pseudorange)
    }
    for _, phaseRange := range msg.SignalData.PhaseRanges {
        w.PutInt32(22, phaseRange)
    }
    for _, phaseRangeLock := range msg.SignalData.PhaseRangeLocks {
        w.PutUint8(4, phaseRangeLock)
    }
    for _, halfCycle := range msg.SignalData.HalfCycles {
        w.PutBit(halfCycle)
    }
    for _, cnr := range msg.SignalData.Cnrs {
        w.PutUint8(6, cnr)
    }

    w.PutUint8(uint(w.Bits()), 0) // Pad with 0s - Should always be less than 1 byte, should check
    w.Flush()
    return data
}

type SatelliteDataMsm123 struct {
    Ranges []uint16
}

func DeserializeSatelliteDataMsm123(r *iobit.Reader, nsat int) (satData SatelliteDataMsm123) {
    for i := 0; i < nsat; i++ {
        satData.Ranges = append(satData.Ranges, r.Uint16(10))
    }
    return satData
}

type SignalDataMsm3 struct {
    Pseudoranges []int16
    PhaseRanges []int32
    PhaseRangeLocks []uint8
    HalfCycles []bool
}

func DeserializeSignalDataMsm3(r *iobit.Reader, ncell int) (sigData SignalDataMsm3) {
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

type MessageMsm3 struct {
    Header MsmHeader
    SatelliteData SatelliteDataMsm123
    SignalData SignalDataMsm3
}

func (msg MessageMsm3) Number() uint16 {
    return msg.Header.MessageNumber
}

func DeserializeMessageMsm3(data []byte) MessageMsm3 {
    r := iobit.NewReader(data)
    msmHeader := NewMsmHeader(&r)
    return MessageMsm3{
        Header: msmHeader,
        SatelliteData: DeserializeSatelliteDataMsm123(&r, bits.OnesCount(uint(msmHeader.SatelliteMask))),
        SignalData: DeserializeSignalDataMsm3(&r, bits.OnesCount(uint(msmHeader.CellMask))),
    }
}

func (msg MessageMsm3) Serialize() (data []byte) {
    satMaskBits := bits.OnesCount(uint(msg.Header.SatelliteMask))
    sigMaskBits := bits.OnesCount(uint(msg.Header.SignalMask))
    cellMaskBits := bits.OnesCount(uint(msg.Header.CellMask))

    msgBits := (169 + (satMaskBits * sigMaskBits)) + (10 * satMaskBits) + (42 * cellMaskBits)
    data = make([]byte, int(math.Ceil(float64(msgBits) / 8)))
    w := iobit.NewWriter(data)

    SerializeMsmHeader(&w, msg.Header)

    for _, ranges := range msg.SatelliteData.Ranges {
        w.PutUint16(10, ranges)
    }
    for _, pseudorange := range msg.SignalData.Pseudoranges {
        w.PutInt16(15, pseudorange)
    }
    for _, phaseRange := range msg.SignalData.PhaseRanges {
        w.PutInt32(22, phaseRange)
    }
    for _, phaseRangeLock := range msg.SignalData.PhaseRangeLocks {
        w.PutUint8(4, phaseRangeLock)
    }
    for _, halfCycle := range msg.SignalData.HalfCycles {
        w.PutBit(halfCycle)
    }

    w.PutUint8(uint(w.Bits()), 0) // Pad with 0s - Should always be less than 1 byte, should check
    w.Flush()
    return data
}

type SignalDataMsm2 struct {
    PhaseRanges []int32
    PhaseRangeLocks []uint8
    HalfCycles []bool
}

func DeserializeSignalDataMsm2(r *iobit.Reader, ncell int) (sigData SignalDataMsm2) {
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

type MessageMsm2 struct {
    Header MsmHeader
    SatelliteData SatelliteDataMsm123
    SignalData SignalDataMsm2
}

func (msg MessageMsm2) Number() uint16 {
    return msg.Header.MessageNumber
}

func DeserializeMessageMsm2(data []byte) MessageMsm2 {
    r := iobit.NewReader(data)
    msmHeader := NewMsmHeader(&r)
    return MessageMsm2{
        Header: msmHeader,
        SatelliteData: DeserializeSatelliteDataMsm123(&r, bits.OnesCount(uint(msmHeader.SatelliteMask))),
        SignalData: DeserializeSignalDataMsm2(&r, bits.OnesCount(uint(msmHeader.CellMask))),
    }
}

func (msg MessageMsm2) Serialize() (data []byte) {
    satMaskBits := bits.OnesCount(uint(msg.Header.SatelliteMask))
    sigMaskBits := bits.OnesCount(uint(msg.Header.SignalMask))
    cellMaskBits := bits.OnesCount(uint(msg.Header.CellMask))

    msgBits := (169 + (satMaskBits * sigMaskBits)) + (10 * satMaskBits) + (27 * cellMaskBits)
    data = make([]byte, int(math.Ceil(float64(msgBits) / 8)))
    w := iobit.NewWriter(data)

    SerializeMsmHeader(&w, msg.Header)

    for _, ranges := range msg.SatelliteData.Ranges {
        w.PutUint16(10, ranges)
    }
    for _, phaseRange := range msg.SignalData.PhaseRanges {
        w.PutInt32(22, phaseRange)
    }
    for _, phaseRangeLock := range msg.SignalData.PhaseRangeLocks {
        w.PutUint8(4, phaseRangeLock)
    }
    for _, halfCycle := range msg.SignalData.HalfCycles {
        w.PutBit(halfCycle)
    }

    w.PutUint8(uint(w.Bits()), 0) // Pad with 0s - Should always be less than 1 byte, should check
    w.Flush()
    return data
}

type SignalDataMsm1 struct {
    Pseudoranges []int16
}

func DeserializeSignalDataMsm1(r *iobit.Reader, ncell int) (sigData SignalDataMsm1) {
    for i := 0; i < ncell; i++ {
        sigData.Pseudoranges = append(sigData.Pseudoranges, r.Int16(15))
    }
    return sigData
}

type MessageMsm1 struct {
    Header MsmHeader
    SatelliteData SatelliteDataMsm123
    SignalData SignalDataMsm1
}

func (msg MessageMsm1) Number() uint16 {
    return msg.Header.MessageNumber
}

func DeserializeMessageMsm1(data []byte) MessageMsm1 {
    r := iobit.NewReader(data)
    msmHeader := NewMsmHeader(&r)
    return MessageMsm1{
        Header: msmHeader,
        SatelliteData: DeserializeSatelliteDataMsm123(&r, bits.OnesCount(uint(msmHeader.SatelliteMask))),
        SignalData: DeserializeSignalDataMsm1(&r, bits.OnesCount(uint(msmHeader.CellMask))),
    }
}

func (msg MessageMsm1) Serialize() (data []byte) {
    satMaskBits := bits.OnesCount(uint(msg.Header.SatelliteMask))
    sigMaskBits := bits.OnesCount(uint(msg.Header.SignalMask))
    cellMaskBits := bits.OnesCount(uint(msg.Header.CellMask))

    msgBits := (169 + (satMaskBits * sigMaskBits)) + (10 * satMaskBits) + (15 * cellMaskBits)
    data = make([]byte, int(math.Ceil(float64(msgBits) / 8)))
    w := iobit.NewWriter(data)

    SerializeMsmHeader(&w, msg.Header)

    for _, ranges := range msg.SatelliteData.Ranges {
        w.PutUint16(10, ranges)
    }
    for _, pseudorange := range msg.SignalData.Pseudoranges {
        w.PutInt16(15, pseudorange)
    }

    w.PutUint8(uint(w.Bits()), 0) // Pad with 0s - Should always be less than 1 byte
    w.Flush()
    return data
}

type Message1071 struct {
    MessageMsm1
}

func DeserializeMessage1071(data []byte) Message1071 {
    return Message1071{
        MessageMsm1: DeserializeMessageMsm1(data),
    }
}

func (msg Message1071) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Message1072 struct {
    MessageMsm2
}

func DeserializeMessage1072(data []byte) Message1072 {
    return Message1072{
        MessageMsm2: DeserializeMessageMsm2(data),
    }
}

func (msg Message1072) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Message1073 struct {
    MessageMsm3
}

func DeserializeMessage1073(data []byte) Message1073 {
    return Message1073{
        MessageMsm3: DeserializeMessageMsm3(data),
    }
}

func (msg Message1073) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Message1074 struct {
    MessageMsm4
}

func DeserializeMessage1074(data []byte) Message1074 {
    return Message1074{
        MessageMsm4: DeserializeMessageMsm4(data),
    }
}

func (msg Message1074) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Message1075 struct {
    MessageMsm5
}

func DeserializeMessage1075(data []byte) Message1075 {
    return Message1075{
        MessageMsm5: DeserializeMessageMsm5(data),
    }
}

func (msg Message1075) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Message1076 struct {
    MessageMsm6
}

func DeserializeMessage1076(data []byte) Message1076 {
    return Message1076{
        MessageMsm6: DeserializeMessageMsm6(data),
    }
}

func (msg Message1076) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Message1077 struct {
    MessageMsm7
}

func DeserializeMessage1077(data []byte) Message1077 {
    return Message1077{
        MessageMsm7: DeserializeMessageMsm7(data),
    }
}

func (msg Message1077) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Message1081 struct {
    MessageMsm1
}

func DeserializeMessage1081(data []byte) Message1081 {
    return Message1081{
        MessageMsm1: DeserializeMessageMsm1(data),
    }
}

func (msg Message1081) Time() time.Time {
    return GlonassTime(msg.Header.Epoch)
}

type Message1082 struct {
    MessageMsm2
}

func DeserializeMessage1082(data []byte) Message1082 {
    return Message1082{
        MessageMsm2: DeserializeMessageMsm2(data),
    }
}

func (msg Message1082) Time() time.Time {
    return GlonassTime(msg.Header.Epoch)
}

type Message1083 struct {
    MessageMsm3
}

func DeserializeMessage1083(data []byte) Message1083 {
    return Message1083{
        MessageMsm3: DeserializeMessageMsm3(data),
    }
}

func (msg Message1083) Time() time.Time {
    return GlonassTime(msg.Header.Epoch)
}

type Message1084 struct {
    MessageMsm4
}

func DeserializeMessage1084(data []byte) Message1084 {
    return Message1084{
        MessageMsm4: DeserializeMessageMsm4(data),
    }
}

func (msg Message1084) Time() time.Time {
    return GlonassTime(msg.Header.Epoch)
}

type Message1085 struct {
    MessageMsm5
}

func DeserializeMessage1085(data []byte) Message1085 {
    return Message1085{
        MessageMsm5: DeserializeMessageMsm5(data),
    }
}

func (msg Message1085) Time() time.Time {
    return GlonassTime(msg.Header.Epoch)
}

type Message1086 struct {
    MessageMsm6
}

func DeserializeMessage1086(data []byte) Message1086 {
    return Message1086{
        MessageMsm6: DeserializeMessageMsm6(data),
    }
}

func (msg Message1086) Time() time.Time {
    return GlonassTime(msg.Header.Epoch)
}

type Message1087 struct {
    MessageMsm7
}

func DeserializeMessage1087(data []byte) Message1087 {
    return Message1087{
        MessageMsm7: DeserializeMessageMsm7(data),
    }
}

func (msg Message1087) Time() time.Time {
    return GlonassTime(msg.Header.Epoch)
}

type Message1091 struct {
    MessageMsm1
}

func DeserializeMessage1091(data []byte) Message1091 {
    return Message1091{
        MessageMsm1: DeserializeMessageMsm1(data),
    }
}

func (msg Message1091) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Message1092 struct {
    MessageMsm2
}

func DeserializeMessage1092(data []byte) Message1092 {
    return Message1092{
        MessageMsm2: DeserializeMessageMsm2(data),
    }
}

func (msg Message1092) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Message1093 struct {
    MessageMsm3
}

func DeserializeMessage1093(data []byte) Message1093 {
    return Message1093{
        MessageMsm3: DeserializeMessageMsm3(data),
    }
}

func (msg Message1093) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Message1094 struct {
    MessageMsm4
}

func DeserializeMessage1094(data []byte) Message1094 {
    return Message1094{
        MessageMsm4: DeserializeMessageMsm4(data),
    }
}

func (msg Message1094) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Message1095 struct {
    MessageMsm5
}

func DeserializeMessage1095(data []byte) Message1095 {
    return Message1095{
        MessageMsm5: DeserializeMessageMsm5(data),
    }
}

func (msg Message1095) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Message1096 struct {
    MessageMsm6
}

func DeserializeMessage1096(data []byte) Message1096 {
    return Message1096{
        MessageMsm6: DeserializeMessageMsm6(data),
    }
}

func (msg Message1096) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Message1097 struct {
    MessageMsm7
}

func DeserializeMessage1097(data []byte) Message1097 {
    return Message1097{
        MessageMsm7: DeserializeMessageMsm7(data),
    }
}

func (msg Message1097) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Message1101 struct {
    MessageMsm1
}

func DeserializeMessage1101(data []byte) Message1101 {
    return Message1101{
        MessageMsm1: DeserializeMessageMsm1(data),
    }
}

func (msg Message1101) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Message1102 struct {
    MessageMsm2
}

func DeserializeMessage1102(data []byte) Message1102 {
    return Message1102{
        MessageMsm2: DeserializeMessageMsm2(data),
    }
}

func (msg Message1102) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Message1103 struct {
    MessageMsm3
}

func DeserializeMessage1103(data []byte) Message1103 {
    return Message1103{
        MessageMsm3: DeserializeMessageMsm3(data),
    }
}

func (msg Message1103) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Message1104 struct {
    MessageMsm4
}

func DeserializeMessage1104(data []byte) Message1104 {
    return Message1104{
        MessageMsm4: DeserializeMessageMsm4(data),
    }
}

func (msg Message1104) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Message1105 struct {
    MessageMsm5
}

func DeserializeMessage1105(data []byte) Message1105 {
    return Message1105{
        MessageMsm5: DeserializeMessageMsm5(data),
    }
}

func (msg Message1105) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Message1106 struct {
    MessageMsm6
}

func DeserializeMessage1106(data []byte) Message1106 {
    return Message1106{
        MessageMsm6: DeserializeMessageMsm6(data),
    }
}

func (msg Message1106) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Message1107 struct {
    MessageMsm7
}

func DeserializeMessage1107(data []byte) Message1107 {
    return Message1107{
        MessageMsm7: DeserializeMessageMsm7(data),
    }
}

func (msg Message1107) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Message1111 struct {
    MessageMsm1
}

func DeserializeMessage1111(data []byte) Message1111 {
    return Message1111{
        MessageMsm1: DeserializeMessageMsm1(data),
    }
}

func (msg Message1111) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Message1112 struct {
    MessageMsm2
}

func DeserializeMessage1112(data []byte) Message1112 {
    return Message1112{
        MessageMsm2: DeserializeMessageMsm2(data),
    }
}

func (msg Message1112) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Message1113 struct {
    MessageMsm3
}

func DeserializeMessage1113(data []byte) Message1113 {
    return Message1113{
        MessageMsm3: DeserializeMessageMsm3(data),
    }
}

func (msg Message1113) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Message1114 struct {
    MessageMsm4
}

func DeserializeMessage1114(data []byte) Message1114 {
    return Message1114{
        MessageMsm4: DeserializeMessageMsm4(data),
    }
}

func (msg Message1114) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Message1115 struct {
    MessageMsm5
}

func DeserializeMessage1115(data []byte) Message1115 {
    return Message1115{
        MessageMsm5: DeserializeMessageMsm5(data),
    }
}

func (msg Message1115) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Message1116 struct {
    MessageMsm6
}

func DeserializeMessage1116(data []byte) Message1116 {
    return Message1116{
        MessageMsm6: DeserializeMessageMsm6(data),
    }
}

func (msg Message1116) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Message1117 struct {
    MessageMsm7
}

func DeserializeMessage1117(data []byte) Message1117 {
    return Message1117{
        MessageMsm7: DeserializeMessageMsm7(data),
    }
}

func (msg Message1117) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
}

type Message1121 struct {
    MessageMsm1
}

func DeserializeMessage1121(data []byte) Message1121 {
    return Message1121{
        MessageMsm1: DeserializeMessageMsm1(data),
    }
}

func (msg Message1121) Time() time.Time {
    return GpsTime(msg.Header.Epoch).Add(14 * time.Second)
}

type Message1122 struct {
    MessageMsm2
}

func DeserializeMessage1122(data []byte) Message1122 {
    return Message1122{
        MessageMsm2: DeserializeMessageMsm2(data),
    }
}

func (msg Message1122) Time() time.Time {
    return GpsTime(msg.Header.Epoch).Add(14 * time.Second)
}

type Message1123 struct {
    MessageMsm3
}

func DeserializeMessage1123(data []byte) Message1123 {
    return Message1123{
        MessageMsm3: DeserializeMessageMsm3(data),
    }
}

func (msg Message1123) Time() time.Time {
    return GpsTime(msg.Header.Epoch).Add(14 * time.Second)
}

type Message1124 struct {
    MessageMsm4
}

func DeserializeMessage1124(data []byte) Message1124 {
    return Message1124{
        MessageMsm4: DeserializeMessageMsm4(data),
    }
}

func (msg Message1124) Time() time.Time {
    return GpsTime(msg.Header.Epoch).Add(14 * time.Second)
}

type Message1125 struct {
    MessageMsm5
}

func DeserializeMessage1125(data []byte) Message1125 {
    return Message1125{
        MessageMsm5: DeserializeMessageMsm5(data),
    }
}

func (msg Message1125) Time() time.Time {
    return GpsTime(msg.Header.Epoch).Add(14 * time.Second)
}

type Message1126 struct {
    MessageMsm6
}

func DeserializeMessage1126(data []byte) Message1126 {
    return Message1126{
        MessageMsm6: DeserializeMessageMsm6(data),
    }
}

func (msg Message1126) Time() time.Time {
    return GpsTime(msg.Header.Epoch).Add(14 * time.Second)
}

type Message1127 struct {
    MessageMsm7
}

func DeserializeMessage1127(data []byte) Message1127 {
    return Message1127{
        MessageMsm7: DeserializeMessageMsm7(data),
    }
}

func (msg Message1127) Time() time.Time {
    return GpsTime(msg.Header.Epoch).Add(14 * time.Second)
}
