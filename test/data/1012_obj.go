package data

import (
    "github.com/geoscienceaustralia/go-rtcm/pkgs/rtcm3"
)

var Message1012 = rtcm3.Message1012 {
    GlonassObservationHeader: rtcm3.GlonassObservationHeader {
        MessageNumber: 0x3f4,
        ReferenceStationId: 0x0,
        Epoch: 0x1450ed8,
        SynchronousGnss: true,
        SignalCount: 0x8,
        SmoothingIndicator: false,
        SmoothingInterval: 0x0,
    },
    SignalData: [] rtcm3.SignalData1012 {
        rtcm3.SignalData1012 {
            SatelliteId: 0x12,
            L1CodeIndicator: false,
            FrequencyChannel: 0x4,
            L1Pseudorange: 0x511d52,
            L1PhaseRange: 2610,
            L1LockTimeIndicator: 0x7f,
            L1PseudorangeAmbiguity: 0x22,
            L1Cnr: 0xca,
            L2CodeIndicator: 0x0,
            L2Pseudorange: 0xf5,
            L2PhaseRange: 18905,
            L2LockTimeIndicator: 0x7f,
            L2Cnr: 0xc4,
        }, rtcm3.SignalData1012 {
            SatelliteId: 0x4,
            L1CodeIndicator: false,
            FrequencyChannel: 0xd,
            L1Pseudorange: 0x5bb9c7,
            L1PhaseRange: 5461,
            L1LockTimeIndicator: 0x7f,
            L1PseudorangeAmbiguity: 0x25,
            L1Cnr: 0xb6,
            L2CodeIndicator: 0x0,
            L2Pseudorange: 0xff,
            L2PhaseRange: 19879,
            L2LockTimeIndicator: 0x7f,
            L2Cnr: 0xb1,
        }, rtcm3.SignalData1012 {
            SatelliteId: 0xf,
            L1CodeIndicator: false,
            FrequencyChannel: 0x7,
            L1Pseudorange: 0x10ce8d,
            L1PhaseRange: -24299,
            L1LockTimeIndicator: 0x7f,
            L1PseudorangeAmbiguity: 0x29,
            L1Cnr: 0x9c,
            L2CodeIndicator: 0x0,
            L2Pseudorange: 0x1b1,
            L2PhaseRange: -25226,
            L2LockTimeIndicator: 0x7f,
            L2Cnr: 0xa4,
        }, rtcm3.SignalData1012 {
            SatelliteId: 0x9,
            L1CodeIndicator: false,
            FrequencyChannel: 0x5,
            L1Pseudorange: 0x92256,
            L1PhaseRange: 2022,
            L1LockTimeIndicator: 0x7f,
            L1PseudorangeAmbiguity: 0x24,
            L1Cnr: 0xbf,
            L2CodeIndicator: 0x0,
            L2Pseudorange: 0xa6,
            L2PhaseRange: 10740,
            L2LockTimeIndicator: 0x7f,
            L2Cnr: 0xb9,
        }, rtcm3.SignalData1012 {
            SatelliteId: 0x3,
            L1CodeIndicator: false,
            FrequencyChannel: 0xc,
            L1Pseudorange: 0xd209bb,
            L1PhaseRange: 8808,
            L1LockTimeIndicator: 0x7f,
            L1PseudorangeAmbiguity: 0x22,
            L1Cnr: 0xcc,
            L2CodeIndicator: 0x0,
            L2Pseudorange: 0xbc,
            L2PhaseRange: 21136,
            L2LockTimeIndicator: 0x7f,
            L2Cnr: 0xc1,
        }, rtcm3.SignalData1012 {
            SatelliteId: 0x13,
            L1CodeIndicator: false,
            FrequencyChannel: 0xa,
            L1Pseudorange: 0x8d817,
            L1PhaseRange: 10048,
            L1LockTimeIndicator: 0x7f,
            L1PseudorangeAmbiguity: 0x23,
            L1Cnr: 0xb6,
            L2CodeIndicator: 0x0,
            L2Pseudorange: 0xb4,
            L2PhaseRange: 31632,
            L2LockTimeIndicator: 0x7f,
            L2Cnr: 0xb0,
        }, rtcm3.SignalData1012 {
            SatelliteId: 0x11,
            L1CodeIndicator: false,
            FrequencyChannel: 0xb,
            L1Pseudorange: 0x1a85b4e,
            L1PhaseRange: -5999,
            L1LockTimeIndicator: 0x7f,
            L1PseudorangeAmbiguity: 0x26,
            L1Cnr: 0xb3,
            L2CodeIndicator: 0x0,
            L2Pseudorange: 0x113,
            L2PhaseRange: 184,
            L2LockTimeIndicator: 0x7f,
            L2Cnr: 0xa2,
        }, rtcm3.SignalData1012 {
            SatelliteId: 0x2,
            L1CodeIndicator: false,
            FrequencyChannel: 0x3,
            L1Pseudorange: 0xe46153,
            L1PhaseRange: 14122,
            L1LockTimeIndicator: 0x7f,
            L1PseudorangeAmbiguity: 0x25,
            L1Cnr: 0xb5,
            L2CodeIndicator: 0x0,
            L2Pseudorange: 0x118,
            L2PhaseRange: 32482,
            L2LockTimeIndicator: 0x7f,
            L2Cnr: 0xa6,
        },
    },
}