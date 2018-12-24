package rtcm

import (
    "github.com/bamiaux/iobit"
    "time"
    "math"
)

func GpsTime(e uint32) time.Time {
    now := time.Now().UTC()
    sow := now.Truncate(time.Hour * 24).AddDate(0, 0, -int(now.Weekday()))
    tow := time.Duration(e) * time.Millisecond
    return sow.Add(-(18 * time.Second)).Add(tow)
}

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

func (h Rtcm3GpsObservationHeader) Serialize() []byte {
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
    Header Rtcm3GpsObservationHeader
    SatelliteData []Rtcm31001SatelliteData
}

func NewRtcm3Message1001(data []byte) Rtcm3Message1001 {
    r := iobit.NewReader(data)
    obsHeader := NewRtcm3GpsObservationHeader(&r)
    return Rtcm3Message1001{
        Header: obsHeader,
        SatelliteData: NewRtcm31001SatelliteData(&r, int(obsHeader.SignalsProcessed)),
    }
}

func (msg Rtcm3Message1001) Serialize() []byte {
    headerData := msg.Header.Serialize()
    satData := make([]byte, int(math.Ceil((58 * float64(msg.Header.SignalsProcessed)) / 8)))
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

func (msg Rtcm3Message1001) Time() time.Time {
    return GpsTime(msg.Header.Epoch)
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
    Header Rtcm3GpsObservationHeader
    SatelliteData []Rtcm31002SatelliteData
}

func NewRtcm3Message1002(data []byte) Rtcm3Message1002 {
    r := iobit.NewReader(data)
    obsHeader := NewRtcm3GpsObservationHeader(&r)
    return Rtcm3Message1002{
        Header: obsHeader,
        SatelliteData: NewRtcm31002SatelliteData(&r, int(obsHeader.SignalsProcessed)),
    }
}

func (msg Rtcm3Message1002) Serialize() []byte {
    headerData := msg.Header.Serialize()
    satData := make([]byte, int(math.Ceil((74 * float64(msg.Header.SignalsProcessed)) / 8)))
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
    Header Rtcm3GpsObservationHeader
    SatelliteData []Rtcm31003SatelliteData
}

func NewRtcm3Message1003(data []byte) Rtcm3Message1003 {
    r := iobit.NewReader(data)
    obsHeader := NewRtcm3GpsObservationHeader(&r)
    return Rtcm3Message1003{
        Header: obsHeader,
        SatelliteData: NewRtcm31003SatelliteData(&r, int(obsHeader.SignalsProcessed)),
    }
}

func (msg Rtcm3Message1003) Serialize() []byte {
    headerData := msg.Header.Serialize()
    satData := make([]byte, int(math.Ceil((101 * float64(msg.Header.SignalsProcessed)) / 8)))
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
    Header Rtcm3GpsObservationHeader
    SatelliteData []Rtcm31004SatelliteData
}

func NewRtcm3Message1004(data []byte) Rtcm3Message1004 {
    r := iobit.NewReader(data)
    obsHeader := NewRtcm3GpsObservationHeader(&r)
    return Rtcm3Message1004{
        Header: obsHeader,
        SatelliteData: NewRtcm31004SatelliteData(&r, int(obsHeader.SignalsProcessed)),
    }
}

func (msg Rtcm3Message1004) Serialize() []byte {
    headerData := msg.Header.Serialize()
    satData := make([]byte, int(math.Ceil((125 * float64(msg.Header.SignalsProcessed)) / 8)))
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


type Rtcm3Message1019 struct {
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

func NewRtcm3Message1019(data []byte) Rtcm3Message1019 {
    r := iobit.NewReader(data)
    return Rtcm3Message1019{
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

func (msg Rtcm3Message1019) Serialize() []byte {
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
