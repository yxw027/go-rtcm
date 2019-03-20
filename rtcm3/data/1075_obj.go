package data

import (
	"github.com/geoscienceaustralia/go-rtcm/rtcm3"
)

var Message1075 = rtcm3.Message1075{
	MessageMsm5: rtcm3.MessageMsm5{
		Header: rtcm3.MsmHeader{
			MessageNumber:          0x433,
			ReferenceStationId:     0x0,
			Epoch:                  0x1a6055a8,
			MultipleMessageBit:     true,
			Iods:                   0x0,
			Reserved:               0x0,
			ClockSteeringIndicator: 0x0,
			ExternalClockIndicator: 0x0,
			SmoothingIndicator:     false,
			SmoothingInterval:      0x0,
			SatelliteMask:          0x481a188a00000000,
			SignalMask:             0x40400000,
			CellMask:               0xfffff,
		},
		SatelliteData: rtcm3.SatelliteDataMsm57{
			RangeMilliseconds: []uint8{
				0x4e, 0x47, 0x49, 0x55, 0x54, 0x54, 0x4d, 0x45, 0x45, 0x52,
			},
			Extended: []uint8{
				0xd, 0x0, 0x0, 0x0, 0xb, 0x0, 0x0, 0x0, 0x0, 0x9,
			},
			Ranges: []uint16{
				0xa, 0x29f, 0x3a9, 0x6e, 0x21a, 0x3ad, 0x6a, 0x3f5, 0x16a, 0xd4,
			},
			PhaseRangeRates: []int16{
				598, 210, 604, -441, -462, -708, -553, 267, -216, 137,
			},
		},
		SignalData: rtcm3.SignalDataMsm5{
			Pseudoranges: []int16{
				8162, 7956, 6213, 6153, 2931, 2878, 662, 802, -5280, -5225, -3774, -3606, 674, 590, -2109, -2028, -3506, -3587, 5126, 5097,
			},
			PhaseRanges: []int32{
				262822, 256113, 213625, 220146, 91808, 91459, 24237, 29542, -163289, -155273, -119196, -114696, 30206, 34901, -68284, -59701, -102851, -98841, 161129, 158977,
			},
			PhaseRangeLocks: []uint8{
				0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xe, 0xe, 0xf, 0xf, 0xb, 0xa, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf,
			},
			HalfCycles: []bool{
				false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
			},
			Cnrs: []uint8{
				0x27, 0x1a, 0x32, 0x2a, 0x2e, 0x23, 0x1f, 0xb, 0x25, 0x13, 0x1d, 0xa, 0x2b, 0x1c, 0x31, 0x2b, 0x33, 0x2d, 0x28, 0x14,
			},
			PhaseRangeRates: []int16{
				4989, 4953, 2425, 2419, 1413, 1407, 4691, 4647, 2859, 2865, -3742, -3627, 179, 186, 3427, 3425, -740, -736, -4568, -4605,
			},
		},
	},
}
