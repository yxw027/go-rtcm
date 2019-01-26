package data

import ( 
    "../../pkgs/rtcm3"
)

var Message1117 = rtcm3.Message1117 {
    MessageMsm7: rtcm3.MessageMsm7 {
        Header: rtcm3.MsmHeader {
            MessageNumber: 0x45d,
            ReferenceStationId: 0x0,
            Epoch: 0x1aa722a8,
            MultipleMessageBit: false,
            Iods: 0x0,
            Reserved: 0x0,
            ClockSteeringIndicator: 0x0,
            ExternalClockIndicator: 0x0,
            SmoothingIndicator: false,
            SmoothingInterval: 0x0,
            SatelliteMask: 0xe200000000000000,
            SignalMask: 0x40010000,
            CellMask: 0xff,
        },
        SatelliteData: rtcm3.SatelliteDataMsm57 {
            RangeMilliseconds: [] uint8 {
                0x82, 0x8e, 0x6f, 0x7c,
            },
            Extended: [] uint8 {
                0xf, 0xf, 0xf, 0xc,
            },
            Ranges: [] uint16 {
                0x240, 0x297, 0xe0, 0x3e0,
            },
            PhaseRangeRates: [] int16 {
                -414, 351, 165, 0,
            },
        },
        SignalData: rtcm3.SignalDataMsm7 {
            Pseudoranges: [] int32 {
                153492, 153671, 75365, 80088, -16298, -15702, -65871, -64718,
            },
            PhaseRanges: [] int32 {
                637364, 681105, 321292, 262913, -44859, -50541, -263102, -278200,
            },
            PhaseRangeLocks: [] uint16 {
                0x264, 0x2c0, 0x2c0, 0x2c0, 0x2c0, 0x2c0, 0x2c0, 0x2c0,
            },
            HalfCycles: [] bool {
                false, false, false, false, false, false, false, false,
            },
            Cnrs: [] uint16 {
                0x28c, 0x274, 0x230, 0x233, 0x320, 0x337, 0x267, 0x2cd,
            },
            PhaseRangeRates: [] int16 {
                1730, 1570, -196, -316, 3103, 3076, 73, 40,
            },
        },
    },
}
