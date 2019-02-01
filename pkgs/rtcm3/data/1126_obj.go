package data

import (
    "github.com/geoscienceaustralia/go-rtcm/pkgs/rtcm3"
)

var Message1126 = rtcm3.Message1126 {
    MessageMsm6: rtcm3.MessageMsm6 {
        Header: rtcm3.MsmHeader {
            MessageNumber: 0x466,
            ReferenceStationId: 0x0,
            Epoch: 0x1a601ef8,
            MultipleMessageBit: true,
            Iods: 0x0,
            Reserved: 0x0,
            ClockSteeringIndicator: 0x1,
            ExternalClockIndicator: 0x0,
            SmoothingIndicator: false,
            SmoothingInterval: 0x0,
            SatelliteMask: 0xf340000820000000,
            SignalMask: 0x40040000,
            CellMask: 0x37ffa,
        },
        SatelliteData: rtcm3.SatelliteDataMsm46 {
            RangeMilliseconds: [] uint8 {
                0x7b, 0x87, 0x80, 0x7c, 0x7b, 0x7c, 0x7c, 0x51, 0x48,
            },
            Ranges: [] uint16 {
                0x2b5, 0x83, 0x58, 0xd6, 0xfd, 0x155, 0x144, 0x381, 0x2c5,
            },
        },
        SignalData: rtcm3.SignalDataMsm6 {
            Pseudoranges: [] int32 {
                158427, 146285, 146608, 39583, 33905, -189345, -195684, 228271, 220874, 63434, 57349, 221647, 220297, 179992, -4176,
            },
            PhaseRanges: [] int32 {
                626157, 570747, 575603, 139827, 115335, -762652, -783251, 889305, 852375, 265285, 251809, 871752, 869231, 696365, -5564,
            },
            PhaseRangeLocks: [] uint16 {
                0x2c0, 0x2c0, 0x2c0, 0x2c0, 0x2c0, 0x2c0, 0x2c0, 0x282, 0x282, 0x272, 0x272, 0x282, 0x282, 0x27f, 0x271,
            },
            HalfCycles: [] bool {
                false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
            },
            Cnrs: [] uint16 {
                0x2df, 0x312, 0x26a, 0x29b, 0x2b2, 0x2d3, 0x2f9, 0x2f4, 0x314, 0x2be, 0x2ef, 0x2db, 0x2f7, 0x2be, 0x326,
            },
        },
    },
}
