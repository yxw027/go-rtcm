package data

import (
	"github.com/geoscienceaustralia/go-rtcm/rtcm3"
)

var Message1077 = rtcm3.Message1077{
	MessageMsm7: rtcm3.MessageMsm7{
		MsmHeader: rtcm3.MsmHeader{
			MessageNumber:          0x435,
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
		SignalData: rtcm3.SignalDataMsm7{
			Pseudoranges: []int32{
				261194, 254606, 198827, 196891, 93794, 92112, 21197, 25661, -168970, -167184, -120762, -115378, 21577, 18881, -67498, -64890, -112206, -114785, 164018, 163103,
			},
			PhaseRanges: []int32{
				1051287, 1024452, 854502, 880582, 367232, 365836, 96947, 118169, -653154, -621093, -476785, -458783, 120824, 139603, -273135, -238802, -411403, -395365, 644516, 635908,
			},
			PhaseRangeLocks: []uint16{
				0x286, 0x286, 0x26f, 0x26f, 0x286, 0x286, 0x1ca, 0x1c8, 0x209, 0x208, 0x16a, 0x14c, 0x242, 0x242, 0x276, 0x276, 0x26f, 0x26f, 0x24b, 0x24b,
			},
			HalfCycles: []bool{
				false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
			},
			Cnrs: []uint16{
				0x270, 0x199, 0x31c, 0x29e, 0x2e0, 0x22c, 0x1ed, 0xa8, 0x252, 0x137, 0x1ca, 0xa2, 0x2b3, 0x1c2, 0x30c, 0x2b7, 0x32d, 0x2c9, 0x279, 0x146,
			},
			PhaseRangeRates: []int16{
				4989, 4953, 2425, 2419, 1413, 1407, 4691, 4647, 2859, 2865, -3742, -3627, 179, 186, 3427, 3425, -740, -736, -4568, -4605,
			},
		},
	},
}
