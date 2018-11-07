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

type Rtcm3Message1020 struct {
    Rtcm3Frame
    MessageNumber uint16
    SatelliteId uint8
    FrequencyChannel uint8
    AlmanacHealth bool
    AlmanacHealthAvailability bool
    P1 uint8
    Tk uint16
    Msb bool
    P2 bool
    Tb uint8
    XnTb1 int32
    XnTb int32
    XnTb2 int8
    YnTb1 int32
    YnTb int32
    YnTb2 int8
    ZnTb1 int32
    ZnTb int32
    ZnTb2 int8
    P3 bool
    GammaN int16
    Mp uint8
    M1n3 bool
    TauN int32
    MDeltaTauN int8
    En uint8
    MP4 bool
    MFt uint8
    MNt uint16
    MM uint8
    AdditionalData bool
    Na uint16
    TauC int32
    MN4 uint8
    MTauGps int32
    M1n5 bool
    Reserved uint8
}

func NewRtcm3Message1020(f Rtcm3Frame) Rtcm3Message1020 {
    r := iobit.NewReader(f.Payload)
    return Rtcm3Message1020{ //TODO: Make sint type
        Rtcm3Frame: f,
        MessageNumber: r.Uint16(12),
        SatelliteId: r.Uint8(6),
        FrequencyChannel: r.Uint8(5),
        AlmanacHealth: r.Bit(),
        AlmanacHealthAvailability: r.Bit(),
        P1: r.Uint8(2),
        Tk: r.Uint16(12),
        Msb: r.Bit(),
        P2: r.Bit(),
        Tb: r.Uint8(7),
        XnTb1: Sint32(r, 24),
        XnTb: Sint32(r, 27),
        XnTb2: Sint8(r, 5),
        YnTb1: Sint32(r, 24),
        YnTb: Sint32(r, 27),
        YnTb2: Sint8(r, 5),
        ZnTb1: Sint32(r, 24),
        ZnTb: Sint32(r, 27),
        ZnTb2: Sint8(r, 5),
        P3: r.Bit(),
        GammaN: Sint16(r, 11),
        Mp: r.Uint8(2),
        M1n3: r.Bit(),
        TauN: Sint32(r, 22),
        MDeltaTauN: Sint8(r, 5),
        En: r.Uint8(5),
        MP4: r.Bit(),
        MFt: r.Uint8(4),
        MNt: r.Uint16(11),
        MM: r.Uint8(2),
        AdditionalData: r.Bit(),
        Na: r.Uint16(11),
        TauC: Sint32(r, 32),
        MN4: r.Uint8(5),
        MTauGps: Sint32(r, 22),
        M1n5: r.Bit(),
        Reserved: r.Uint8(7),
    }
}

type Rtcm3Message1230 struct {
    Rtcm3Frame
    MessageNumber uint16
    ReferenceStationId uint16
    CodePhaseBias bool
    Reserved uint8
    SignalsMask uint8
    L1CACodePhaseBias int16
    L1PCodePhaseBias int16
    L2CACodePhaseBias int16
    L2PCodePhaseBias int16
}

func NewRtcm3Message1230(f Rtcm3Frame) (msg Rtcm3Message1230) {
    r := iobit.NewReader(f.Payload)
    msg = Rtcm3Message1230{
        Rtcm3Frame: f,
        MessageNumber: r.Uint16(12),
        ReferenceStationId: r.Uint16(12),
        CodePhaseBias: r.Bit(),
        Reserved: r.Uint8(3),
        SignalsMask: r.Uint8(4),
    }
    if (msg.SignalsMask & 8) == 8 {
        msg.L1CACodePhaseBias = r.Int16(16)
    }
    if (msg.SignalsMask & 4) == 4 {
        msg.L1PCodePhaseBias = r.Int16(16)
    }
    if (msg.SignalsMask & 2) == 2 {
        msg.L2CACodePhaseBias = r.Int16(16)
    }
    if (msg.SignalsMask & 1) == 1 {
        msg.L2PCodePhaseBias = r.Int16(16)
    }
    return msg
}
