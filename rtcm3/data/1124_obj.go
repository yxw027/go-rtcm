package data

import (
	"github.com/geoscienceaustralia/go-rtcm/rtcm3"
)

var Message1124 = rtcm3.Message1124{
	MessageMsm4: rtcm3.MessageMsm4{
		Header: rtcm3.MsmHeader{
			MessageNumber:          0x464,
			ReferenceStationId:     0x0,
			Epoch:                  0x1a601ef8,
			MultipleMessageBit:     true,
			Iods:                   0x0,
			Reserved:               0x0,
			ClockSteeringIndicator: 0x1,
			ExternalClockIndicator: 0x0,
			SmoothingIndicator:     false,
			SmoothingInterval:      0x0,
			SatelliteMask:          0xf340000820000000,
			SignalMask:             0x40040000,
			CellMask:               0x37ffa,
		},
		SatelliteData: rtcm3.SatelliteDataMsm46{
			RangeMilliseconds: []uint8{
				0x7b, 0x87, 0x80, 0x7c, 0x7b, 0x7c, 0x7c, 0x51, 0x48,
			},
			Ranges: []uint16{
				0x2b5, 0x83, 0x58, 0xd6, 0xfd, 0x155, 0x144, 0x381, 0x2c5,
			},
		},
		SignalData: rtcm3.SignalDataMsm4{
			Pseudoranges: []int16{
				4951, 4571, 4582, 1237, 1060, -5917, -6115, 7133, 6902, 1982, 1792, 6926, 6884, 5625, -131,
			},
			PhaseRanges: []int32{
				156539, 142687, 143901, 34957, 28834, -190663, -195813, 222326, 213094, 66321, 62952, 217938, 217308, 174091, -1391,
			},
			PhaseRangeLocks: []uint8{
				0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf,
			},
			HalfCycles: []bool{
				false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
			},
			Cnrs: []uint8{
				0x2e, 0x31, 0x27, 0x2a, 0x2b, 0x2d, 0x30, 0x2f, 0x31, 0x2c, 0x2f, 0x2e, 0x2f, 0x2c, 0x32,
			},
		},
	},
}
