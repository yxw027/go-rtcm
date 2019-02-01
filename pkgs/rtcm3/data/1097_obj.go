package data

import (
    "github.com/geoscienceaustralia/go-rtcm/pkgs/rtcm3"
)

var Message1097 = rtcm3.Message1097 {
    MessageMsm7: rtcm3.MessageMsm7 {
        Header: rtcm3.MsmHeader {
            MessageNumber: 0x449,
            ReferenceStationId: 0x0,
            Epoch: 0x1a6055a8,
            MultipleMessageBit: true,
            Iods: 0x0,
            Reserved: 0x0,
            ClockSteeringIndicator: 0x0,
            ExternalClockIndicator: 0x0,
            SmoothingIndicator: false,
            SmoothingInterval: 0x0,
            SatelliteMask: 0x823021c200000000,
            SignalMask: 0x40020200,
            CellMask: 0x7ffffff,
        },
        SatelliteData: rtcm3.SatelliteDataMsm57 {
            RangeMilliseconds: [] uint8 {
                0x60, 0x5f, 0x56, 0x4a, 0x57, 0x4d, 0x5b, 0x59, 0x53,
            },
            Extended: [] uint8 {
                0xd, 0x0, 0x0, 0x0, 0xb, 0x0, 0x0, 0x0, 0x0,
            },
            Ranges: [] uint16 {
                0xca, 0x59, 0x2fc, 0x2c8, 0x2b1, 0xac, 0xa1, 0x3da, 0x31d,
            },
            PhaseRangeRates: [] int16 {
                584, 260, -342, -120, -469, 21, -399, 367, 494,
            },
        },
        SignalData: rtcm3.SignalDataMsm7 {
            Pseudoranges: [] int32 {
                112477, 118281, 124339, -194173, -191897, -186094, 27842, 28799, 34634, -247481, -252214, -247975, 215519, 215592, 220307, 126435, 146590, 151775, -60421, -57050, -52141, 159537, 167775, 171420, -218873, -215904, -210919,
            },
            PhaseRanges: [] int32 {
                369517, 368277, 384684, -773585, -747122, -722054, 140325, 139396, 166482, -948098, -945580, -930179, 877215, 958310, 970909, 526036, 618641, 640467, -229038, -223404, -196799, 599417, 583124, 594663, -899105, -891412, -871164,
            },
            PhaseRangeLocks: [] uint16 {
                0x2a0, 0x2a0, 0x2a0, 0x25e, 0x25e, 0x25e, 0x228, 0x229, 0x229, 0x275, 0x275, 0x276, 0x21c, 0x246, 0x245, 0x268, 0x268, 0x268, 0x21c, 0x21d, 0x21e, 0x297, 0x297, 0x297, 0x287, 0x287, 0x287,
            },
            HalfCycles: [] bool {
                false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
            },
            Cnrs: [] uint16 {
                0x21c, 0x247, 0x247, 0x244, 0x23f, 0x22a, 0x17c, 0x22a, 0x269, 0x2e1, 0x309, 0x2fa, 0x24f, 0x282, 0x26d, 0x2e3, 0x30b, 0x301, 0x226, 0x26c, 0x264, 0x1f1, 0x259, 0x25a, 0x2ba, 0x2f3, 0x2e4,
            },
            PhaseRangeRates: [] int16 {
                -4304, -4007, -4122, -194, -177, -628, -902, -843, -742, -3752, -3812, -3874, -740, -1643, -1149, 4684, 4668, 4830, 4320, 4364, 4431, -438, 168, -186, -4849, -4669, -4626,
            },
        },
    },
}
