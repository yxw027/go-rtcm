package data

import (
    "github.com/geoscienceaustralia/go-rtcm/rtcm3"
)

var Message1002 = rtcm3.Message1002 {
    GpsObservationHeader: rtcm3.GpsObservationHeader {
        MessageNumber: 0x3ea,
        ReferenceStationId: 0x0,
        Epoch: 0x1a6055a8,
        SynchronousGnss: true,
        SignalsProcessed: 0xa,
        SmoothingIndicator: false,
        SmoothingInterval: 0x0,
    },
    SatelliteData: [] rtcm3.SatelliteData1002 {
        rtcm3.SatelliteData1002 {
            SatelliteId: 0x5,
            L1CodeIndicator: false,
            L1Pseudorange: 0x519499,
            L1PhaseRange: 16539,
            L1LockTimeIndicator: 0x7f,
            L1PseudorangeAmbiguity: 0x47,
            L1Cnr: 0xc7,
        }, rtcm3.SatelliteData1002 {
            SatelliteId: 0x1d,
            L1CodeIndicator: false,
            L1Pseudorange: 0xc6dc3,
            L1PhaseRange: 10446,
            L1LockTimeIndicator: 0x7f,
            L1PseudorangeAmbiguity: 0x45,
            L1Cnr: 0xcb,
        }, rtcm3.SatelliteData1002 {
            SatelliteId: 0x19,
            L1CodeIndicator: false,
            L1Pseudorange: 0x9ddb6a,
            L1PhaseRange: -874,
            L1LockTimeIndicator: 0x7f,
            L1PseudorangeAmbiguity: 0x45,
            L1Cnr: 0xc3,
        }, rtcm3.SatelliteData1002 {
            SatelliteId: 0xf,
            L1CodeIndicator: false,
            L1Pseudorange: 0x33b764,
            L1PhaseRange: 6339,
            L1LockTimeIndicator: 0x7f,
            L1PseudorangeAmbiguity: 0x54,
            L1Cnr: 0x94,
        }, rtcm3.SatelliteData1002 {
            SatelliteId: 0x1f,
            L1CodeIndicator: false,
            L1Pseudorange: 0xcfc3fe,
            L1PhaseRange: -3232,
            L1LockTimeIndicator: 0x7f,
            L1PseudorangeAmbiguity: 0x51,
            L1Cnr: 0x9e,
        }, rtcm3.SatelliteData1002 {
            SatelliteId: 0x2,
            L1CodeIndicator: false,
            L1Pseudorange: 0xa2b015,
            L1PhaseRange: 1803,
            L1LockTimeIndicator: 0x7f,
            L1PseudorangeAmbiguity: 0x4d,
            L1Cnr: 0x9c,
        }, rtcm3.SatelliteData1002 {
            SatelliteId: 0x15,
            L1CodeIndicator: false,
            L1Pseudorange: 0xb8073e,
            L1PhaseRange: 9651,
            L1LockTimeIndicator: 0x7f,
            L1PseudorangeAmbiguity: 0x4c,
            L1Cnr: 0xad,
        }, rtcm3.SatelliteData1002 {
            SatelliteId: 0xd,
            L1CodeIndicator: false,
            L1Pseudorange: 0xb8ebef,
            L1PhaseRange: 3381,
            L1LockTimeIndicator: 0x5f,
            L1PseudorangeAmbiguity: 0x54,
            L1Cnr: 0x7b,
        }, rtcm3.SatelliteData1002 {
            SatelliteId: 0x14,
            L1CodeIndicator: false,
            L1Pseudorange: 0x8dc086,
            L1PhaseRange: 1768,
            L1LockTimeIndicator: 0x21,
            L1PseudorangeAmbiguity: 0x54,
            L1Cnr: 0x72,
        }, rtcm3.SatelliteData1002 {
            SatelliteId: 0xc,
            L1CodeIndicator: false,
            L1Pseudorange: 0x8cf347,
            L1PhaseRange: -2203,
            L1LockTimeIndicator: 0x7f,
            L1PseudorangeAmbiguity: 0x49,
            L1Cnr: 0xb8,
        },
    },
}
