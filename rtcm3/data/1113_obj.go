package data

import (
	"github.com/geoscienceaustralia/go-rtcm/rtcm3"
)

var Message1113 = rtcm3.Message1113{
	MessageMsm3: rtcm3.MessageMsm3{
		MsmHeader: rtcm3.MsmHeader{
			MessageNumber:          0x459,
			ReferenceStationId:     0x0,
			Epoch:                  0x1a6055a8,
			MultipleMessageBit:     true,
			Iods:                   0x0,
			Reserved:               0x0,
			ClockSteeringIndicator: 0x1,
			ExternalClockIndicator: 0x0,
			SmoothingIndicator:     false,
			SmoothingInterval:      0x0,
			SatelliteMask:          0xe200000000000000,
			SignalMask:             0x40010000,
			CellMask:               0xff,
		},
		SatelliteData: rtcm3.SatelliteDataMsm123{
			Ranges: []uint16{
				0x25b, 0x7a, 0x18f, 0x2b2,
			},
		},
		SignalData: rtcm3.SignalDataMsm3{
			Pseudoranges: []int16{
				-383, -377, 6300, 6444, 5265, 5272, 5458, 5491,
			},
			PhaseRanges: []int32{
				-8627, 1626, 212852, 200210, 173790, 172517, 175111, 171406,
			},
			PhaseRangeLocks: []uint8{
				0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf,
			},
			HalfCycles: []bool{
				false, false, false, false, false, false, false, false,
			},
		},
	},
}
