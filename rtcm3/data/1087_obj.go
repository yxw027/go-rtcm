package data

import (
	"github.com/geoscienceaustralia/go-rtcm/rtcm3"
)

var Message1087 = rtcm3.Message1087{
	MessageMsm7: rtcm3.MessageMsm7{
		Header: rtcm3.MsmHeader{
			MessageNumber:          0x43f,
			ReferenceStationId:     0x0,
			Epoch:                  0x29450ed8,
			MultipleMessageBit:     true,
			Iods:                   0x0,
			Reserved:               0x0,
			ClockSteeringIndicator: 0x0,
			ExternalClockIndicator: 0x0,
			SmoothingIndicator:     false,
			SmoothingInterval:      0x0,
			SatelliteMask:          0x7082e00000000000,
			SignalMask:             0x41000000,
			CellMask:               0xffff,
		},
		SatelliteData: rtcm3.SatelliteDataMsm57{
			RangeMilliseconds: []uint8{
				0x4b, 0x45, 0x4a, 0x48, 0x52, 0x4e, 0x44, 0x46,
			},
			Extended: []uint8{
				0x3, 0xc, 0xd, 0x5, 0x7, 0xb, 0x4, 0xa,
			},
			Ranges: []uint16{
				0x131, 0xde, 0x2cd, 0x15b, 0x17d, 0x9e, 0x29d, 0x15a,
			},
			PhaseRangeRates: []int16{
				451, -240, -773, -95, 792, 443, 148, -258,
			},
		},
		SignalData: rtcm3.SignalDataMsm7{
			Pseudoranges: []int32{
				-207384, -197344, 255255, 262012, -104514, -95399, 18715, 24678, 202227, 217750, -2432, 7447, 153776, 162560, -137655, -131193,
			},
			PhaseRanges: []int32{
				-778972, -713216, 1052621, 1096774, -398519, -346881, 82166, 113389, 721896, 718576, -31154, -9012, 624465, 682828, -514620, -437314,
			},
			PhaseRangeLocks: []uint16{
				0x264, 0x264, 0x243, 0x243, 0x20e, 0x20e, 0x238, 0x238, 0x284, 0x284, 0x284, 0x284, 0x272, 0x272, 0x261, 0x261,
			},
			HalfCycles: []bool{
				false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
			},
			Cnrs: []uint16{
				0x2d3, 0x298, 0x331, 0x304, 0x2d7, 0x2c2, 0x2fd, 0x2e2, 0x26f, 0x28f, 0x2cc, 0x289, 0x328, 0x310, 0x2d6, 0x2c1,
			},
			PhaseRangeRates: []int16{
				1968, 2005, -258, -376, -3548, -3301, -1755, -1689, -760, -733, -1626, -1607, 388, 344, -1626, -1461,
			},
		},
	},
}
