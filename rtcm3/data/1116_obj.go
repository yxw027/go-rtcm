package data

import (
	"github.com/geoscienceaustralia/go-rtcm/rtcm3"
)

var Message1116 = rtcm3.Message1116{
	MessageMsm6: rtcm3.MessageMsm6{
		MsmHeader: rtcm3.MsmHeader{
			MessageNumber:          0x45c,
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
		SatelliteData: rtcm3.SatelliteDataMsm46{
			RangeMilliseconds: []uint8{
				0x88, 0x88, 0x6d, 0x7c,
			},
			Ranges: []uint16{
				0x25b, 0x7a, 0x18f, 0x2b2,
			},
		},
		SignalData: rtcm3.SignalDataMsm6{
			Pseudoranges: []int32{
				-12252, -12060, 201590, 206220, 168492, 168692, 174667, 175699,
			},
			PhaseRanges: []int32{
				-34509, 6505, 851407, 800839, 695161, 690068, 700445, 685624,
			},
			PhaseRangeLocks: []uint16{
				0x244, 0x2c0, 0x2c0, 0x2c0, 0x2c0, 0x2c0, 0x2c0, 0x2c0,
			},
			HalfCycles: []bool{
				false, false, false, false, false, false, false, false,
			},
			Cnrs: []uint16{
				0x25a, 0x248, 0x26b, 0x275, 0x31c, 0x332, 0x25d, 0x2c4,
			},
		},
	},
}
