package rtcm3

import (
    "github.com/bamiaux/iobit"
    "time"
    "math"
)

func GpsTime(epoch uint32, week time.Time) time.Time {
    tow := time.Duration(epoch) * time.Millisecond
    return week.Add(-(18 * time.Second)).Add(tow) //TODO: Better handling leap second
}

type GpsObservationHeader struct {
    MessageNumber uint16
    ReferenceStationId uint16
    Epoch uint32
    SynchronousGnss bool
    SignalsProcessed uint8
    SmoothingIndicator bool
    SmoothingInterval uint8
}

func (obsHeader GpsObservationHeader) Number() uint16 {
    return obsHeader.MessageNumber
}

func NewGpsObservationHeader(r *iobit.Reader) (header GpsObservationHeader) {
    return GpsObservationHeader{
        MessageNumber: r.Uint16(12),
        ReferenceStationId: r.Uint16(12),
        Epoch: r.Uint32(30),
        SynchronousGnss: r.Bit(),
        SignalsProcessed: r.Uint8(5),
        SmoothingIndicator: r.Bit(),
        SmoothingInterval: r.Uint8(3),
    }
}

func SerializeGpsObservationHeader(h GpsObservationHeader) []byte {
    data := make([]byte, 8) // 64 bit header
    w := iobit.NewWriter(data)
    w.PutUint16(12, h.MessageNumber)
    w.PutUint16(12, h.ReferenceStationId)
    w.PutUint32(30, h.Epoch)
    w.PutBit(h.SynchronousGnss)
    w.PutUint8(5, h.SignalsProcessed)
    w.PutBit(h.SmoothingIndicator)
    w.PutUint8(3, h.SmoothingInterval)
    w.Flush()
    return data
}

type SatelliteData1001 struct {
    SatelliteId uint8
    L1CodeIndicator bool
    L1Pseudorange uint32
    L1PhaseRange int32 //PhaseRange-Pseudorange?
    L1LockTimeIndicator uint8
}

func DeserializeSatelliteData1001(r *iobit.Reader, numSats int) (satData []SatelliteData1001) {
    for i := 0; i < numSats; i++ {
        satData = append(satData, SatelliteData1001{
            SatelliteId: r.Uint8(6),
            L1CodeIndicator: r.Bit(),
            L1Pseudorange: r.Uint32(24),
            L1PhaseRange: r.Int32(20),
            L1LockTimeIndicator: r.Uint8(7),
        })
    }
    return satData
}

type Message1001 struct {
    GpsObservationHeader
    SatelliteData []SatelliteData1001
}

func DeserializeMessage1001(data []byte) Message1001 {
    r := iobit.NewReader(data)
    obsHeader := NewGpsObservationHeader(&r)
    return Message1001{
        GpsObservationHeader: obsHeader,
        SatelliteData: DeserializeSatelliteData1001(&r, int(obsHeader.SignalsProcessed)),
    }
}

func (msg Message1001) Serialize() []byte {
    headerData := SerializeGpsObservationHeader(msg.GpsObservationHeader)
    satData := make([]byte, int(math.Ceil((58 * float64(msg.SignalsProcessed)) / 8)))
    w := iobit.NewWriter(satData)
    for _, s := range msg.SatelliteData {
        w.PutUint8(6, s.SatelliteId)
        w.PutBit(s.L1CodeIndicator)
        w.PutUint32(24, s.L1Pseudorange)
        w.PutInt32(20, s.L1PhaseRange)
        w.PutUint8(7, s.L1LockTimeIndicator)
    }
    w.PutUint8(uint(w.Bits()), 0)
    w.Flush()
    return append(headerData, satData...)
}

func (msg Message1001) Time() time.Time {
    now := time.Now().UTC()
    sow := now.Truncate(time.Hour * 24).AddDate(0, 0, -int(now.Weekday()))
    return GpsTime(msg.Epoch, sow)
}

type SatelliteData1002 struct {
    SatelliteId uint8
    L1CodeIndicator bool
    L1Pseudorange uint32
    L1PhaseRange int32
    L1LockTimeIndicator uint8
    L1PseudorangeAmbiguity uint8
    L1Cnr uint8
}

func DeserializeSatelliteData1002(r *iobit.Reader, numSats int) (satData []SatelliteData1002) {
    for i := 0; i < numSats; i++ {
        satData = append(satData, SatelliteData1002{
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

type Message1002 struct {
    GpsObservationHeader
    SatelliteData []SatelliteData1002
}

func DeserializeMessage1002(data []byte) Message1002 {
    r := iobit.NewReader(data)
    obsHeader := NewGpsObservationHeader(&r)
    return Message1002{
        GpsObservationHeader: obsHeader,
        SatelliteData: DeserializeSatelliteData1002(&r, int(obsHeader.SignalsProcessed)),
    }
}

func (msg Message1002) Serialize() []byte {
    headerData := SerializeGpsObservationHeader(msg.GpsObservationHeader)
    satData := make([]byte, int(math.Ceil((74 * float64(msg.SignalsProcessed)) / 8)))
    w := iobit.NewWriter(satData)
    for _, s := range msg.SatelliteData {
        w.PutUint8(6, s.SatelliteId)
        w.PutBit(s.L1CodeIndicator)
        w.PutUint32(24, s.L1Pseudorange)
        w.PutInt32(20, s.L1PhaseRange)
        w.PutUint8(7, s.L1LockTimeIndicator)
        w.PutUint8(8, s.L1PseudorangeAmbiguity)
        w.PutUint8(8, s.L1Cnr)
    }
    w.PutUint8(uint(w.Bits()), 0)
    w.Flush()
    return append(headerData, satData...)
}

func (msg Message1002) Time() time.Time {
    now := time.Now().UTC()
    sow := now.Truncate(time.Hour * 24).AddDate(0, 0, -int(now.Weekday()))
    return GpsTime(msg.Epoch, sow)
}


type SatelliteData1003 struct {
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

func DeserializeSatelliteData1003(r *iobit.Reader, numSats int) (satData []SatelliteData1003) {
    for i := 0; i < numSats; i++ {
        satData = append(satData, SatelliteData1003{
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

type Message1003 struct {
    GpsObservationHeader
    SatelliteData []SatelliteData1003
}

func DeserializeMessage1003(data []byte) Message1003 {
    r := iobit.NewReader(data)
    obsHeader := NewGpsObservationHeader(&r)
    return Message1003{
        GpsObservationHeader: obsHeader,
        SatelliteData: DeserializeSatelliteData1003(&r, int(obsHeader.SignalsProcessed)),
    }
}

func (msg Message1003) Serialize() []byte {
    headerData := SerializeGpsObservationHeader(msg.GpsObservationHeader)
    satData := make([]byte, int(math.Ceil((101 * float64(msg.SignalsProcessed)) / 8)))
    w := iobit.NewWriter(satData)
    for _, s := range msg.SatelliteData {
        w.PutUint8(6, s.SatelliteId)
        w.PutBit(s.L1CodeIndicator)
        w.PutUint32(24, s.L1Pseudorange)
        w.PutInt32(20, s.L1PhaseRange)
        w.PutUint8(7, s.L1LockTimeIndicator)
        w.PutUint8(2, s.L2CodeIndicator)
        w.PutInt16(14, s.L2PseudorangeDifference)
        w.PutInt32(20, s.L2PhaseRange)
        w.PutUint8(7, s.L2LockTimeIndicator)
    }
    w.PutUint8(uint(w.Bits()), 0)
    w.Flush()
    return append(headerData, satData...)
}

func (msg Message1003) Time() time.Time {
    now := time.Now().UTC()
    sow := now.Truncate(time.Hour * 24).AddDate(0, 0, -int(now.Weekday()))
    return GpsTime(msg.Epoch, sow)
}


type SatelliteData1004 struct {
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

func DeserializeSatelliteData1004(r *iobit.Reader, numSats int) (satData []SatelliteData1004) {
    for i := 0; i < numSats; i++ {
        satData = append(satData, SatelliteData1004{
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

type Message1004 struct {
    GpsObservationHeader
    SatelliteData []SatelliteData1004
}

func DeserializeMessage1004(data []byte) Message1004 {
    r := iobit.NewReader(data)
    obsHeader := NewGpsObservationHeader(&r)
    return Message1004{
        GpsObservationHeader: obsHeader,
        SatelliteData: DeserializeSatelliteData1004(&r, int(obsHeader.SignalsProcessed)),
    }
}

func (msg Message1004) Serialize() []byte {
    headerData := SerializeGpsObservationHeader(msg.GpsObservationHeader)
    satData := make([]byte, int(math.Ceil((125 * float64(msg.SignalsProcessed)) / 8)))
    w := iobit.NewWriter(satData)
    for _, s := range msg.SatelliteData {
        w.PutUint8(6, s.SatelliteId)
        w.PutBit(s.L1CodeIndicator)
        w.PutUint32(24, s.L1Pseudorange)
        w.PutInt32(20, s.L1PhaseRange)
        w.PutUint8(7, s.L1LockTimeIndicator)
        w.PutUint8(8, s.L1PseudorangeAmbiguity)
        w.PutUint8(8, s.L1Cnr)
        w.PutUint8(2, s.L2CodeIndicator)
        w.PutInt16(14, s.L2PseudorangeDifference)
        w.PutInt32(20, s.L2PhaseRange)
        w.PutUint8(7, s.L2LockTimeIndicator)
        w.PutUint8(8, s.L2Cnr)
    }
    w.PutUint8(uint(w.Bits()), 0)
    w.Flush()
    return append(headerData, satData...)
}

func (msg Message1004) Time() time.Time {
    now := time.Now().UTC()
    sow := now.Truncate(time.Hour * 24).AddDate(0, 0, -int(now.Weekday()))
    return GpsTime(msg.Epoch, sow)
}


type Message1019 struct {
    MessageNumber uint16
    SatelliteId uint8
    GpsWeekNumber uint16
    SvAccuracy uint8
    L2Code uint8
    Idot int16
    Iode uint8
    Toc uint16
    Af2 int8
    Af1 int16
    Af0 int32
    Iodc uint16
    Crs int16
    DeltaN int16
    M0 int32
    Cuc int16
    Eccentricity uint32
    Cus int16
    SrA uint32
    Toe uint16
    Cic int16
    Omega0 int32
    Cis int16
    I0 int32
    C_rc int16
    Perigee int32
    OmegaDot int32
    Tgd int8
    SvHealth uint8
    L2PDataFlag bool
    FitInterval bool
}

func (msg Message1019) Number() uint16 {
    return msg.MessageNumber
}

func DeserializeMessage1019(data []byte) Message1019 {
    r := iobit.NewReader(data)
    return Message1019{
        MessageNumber: r.Uint16(12),
        SatelliteId: r.Uint8(6),
        GpsWeekNumber: r.Uint16(10),
        SvAccuracy: r.Uint8(4),
        L2Code: r.Uint8(2),
        Idot: r.Int16(14),
        Iode: r.Uint8(8),
        Toc: r.Uint16(16),
        Af2: r.Int8(8),
        Af1: r.Int16(16),
        Af0: r.Int32(22),
        Iodc: r.Uint16(10),
        Crs: r.Int16(16),
        DeltaN: r.Int16(16),
        M0: r.Int32(32),
        Cuc: r.Int16(16),
        Eccentricity: r.Uint32(32),
        Cus: r.Int16(16),
        SrA: r.Uint32(32),
        Toe: r.Uint16(16),
        Cic: r.Int16(16),
        Omega0: r.Int32(32),
        Cis: r.Int16(16),
        I0: r.Int32(32),
        C_rc: r.Int16(16),
        Perigee: r.Int32(32),
        OmegaDot: r.Int32(24),
        Tgd: r.Int8(8),
        SvHealth: r.Uint8(6),
        L2PDataFlag: r.Bit(),
        FitInterval: r.Bit(),
    }
}

func (msg Message1019) Serialize() []byte {
    data := make([]byte, 61)
    w := iobit.NewWriter(data)
    w.PutUint16(12, msg.MessageNumber)
    w.PutUint8(6, msg.SatelliteId)
    w.PutUint16(10, msg.GpsWeekNumber)
    w.PutUint8(4, msg.SvAccuracy)
    w.PutUint8(2, msg.L2Code)
    w.PutInt16(14, msg.Idot)
    w.PutUint8(8, msg.Iode)
    w.PutUint16(16, msg.Toc)
    w.PutInt8(8, msg.Af2)
    w.PutInt16(16, msg.Af1)
    w.PutInt32(22, msg.Af0)
    w.PutUint16(10, msg.Iodc)
    w.PutInt16(16, msg.Crs)
    w.PutInt16(16, msg.DeltaN)
    w.PutInt32(32, msg.M0)
    w.PutInt16(16, msg.Cuc)
    w.PutUint32(32, msg.Eccentricity)
    w.PutInt16(16, msg.Cus)
    w.PutUint32(32, msg.SrA)
    w.PutUint16(16, msg.Toe)
    w.PutInt16(16, msg.Cic)
    w.PutInt32(32, msg.Omega0)
    w.PutInt16(16, msg.Cis)
    w.PutInt32(32, msg.I0)
    w.PutInt16(16, msg.C_rc)
    w.PutInt32(32, msg.Perigee)
    w.PutInt32(24, msg.OmegaDot)
    w.PutInt8(8, msg.Tgd)
    w.PutUint8(6, msg.SvHealth)
    w.PutBit(msg.L2PDataFlag)
    w.PutBit(msg.FitInterval)
    w.Flush()
    return data
}
