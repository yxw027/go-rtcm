package data

import (
	"github.com/geoscienceaustralia/go-rtcm/rtcm3"
)

var Message10852 = rtcm3.Message1085{
	MessageMsm5: rtcm3.MessageMsm5{
		Header: rtcm3.MsmHeader{
			MessageNumber:          0x43d,
			ReferenceStationId:     0x0,
			Epoch:                  0x19c13030,
			MultipleMessageBit:     true,
			Iods:                   0x0,
			Reserved:               0x0,
			ClockSteeringIndicator: 0x0,
			ExternalClockIndicator: 0x0,
			SmoothingIndicator:     false,
			SmoothingInterval:      0x0,
			SatelliteMask:          0xe002830000000000,
			SignalMask:             0x41000000,
			CellMask:               0x3fff,
		},
		SatelliteData: rtcm3.SatelliteDataMsm57{
			RangeMilliseconds: []uint8{
				0x45, 0x40, 0x4d, 0x51, 0x43, 0x4c, 0x42,
			},
			Extended: []uint8{
				0x8, 0x3, 0xc, 0x7, 0xb, 0xa, 0x9,
			},
			Ranges: []uint16{
				0x22a, 0x312, 0x4f, 0x358, 0x33b, 0x58, 0x32c,
			},
			PhaseRangeRates: []int16{
				557, -323, -801, 540, -429, 546, 124,
			},
		},
		SignalData: rtcm3.SignalDataMsm5{
			Pseudoranges: []int16{
				4830, 5148, -4952, -4698, -1564, -1240, 3365, 3726, -535, -308, 7584, 7967, 5526, 5709,
			},
			PhaseRanges: []int32{
				154638, 164958, -148114, -132445, -47978, -40450, 98404, 108685, -12709, 3898, 245642, 265667, 179926, 183863,
			},
			PhaseRangeLocks: []uint8{
				0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf,
			},
			HalfCycles: []bool{
				false, false, false, false, false, false, false, false, false, false, false, false, false, false,
			},
			Cnrs: []uint8{
				0x30, 0x2c, 0x35, 0x32, 0x2c, 0x29, 0x23, 0x22, 0x34, 0x31, 0x2e, 0x25, 0x30, 0x26,
			},
			PhaseRangeRates: []int16{
				4301, 4429, -3033, -3091, 2764, 2483, 4393, 4905, -2593, -2715, -136, -664, -3852, -3371,
			},
		},
	},
}
