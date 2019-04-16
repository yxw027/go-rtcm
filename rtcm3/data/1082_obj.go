package data

import (
	"github.com/geoscienceaustralia/go-rtcm/rtcm3"
)

var Message1082 = rtcm3.Message1082{
	MessageMsm2: rtcm3.MessageMsm2{
		MsmHeader: rtcm3.MsmHeader{
			MessageNumber:          0x43a,
			ReferenceStationId:     0x0,
			Epoch:                  0x29450ed8,
			MultipleMessageBit:     true,
			Iods:                   0x0,
			Reserved:               0x0,
			ClockSteeringIndicator: 0x1,
			ExternalClockIndicator: 0x0,
			SmoothingIndicator:     false,
			SmoothingInterval:      0x0,
			SatelliteMask:          0x7082e00000000000,
			SignalMask:             0x41000000,
			CellMask:               0xffff,
		},
		SatelliteData: rtcm3.SatelliteDataMsm123{
			Ranges: []uint16{
				0x3fe, 0x3ac, 0x19b, 0x29, 0x4b, 0x36c, 0x16b, 0x28,
			},
		},
		SignalData: rtcm3.SignalDataMsm2{
			PhaseRanges: []int32{
				255361, 271800, 188601, 199639, -174469, -161560, -53935, -46129, 106472, 105642, -81977, -76441, 81770, 96361, -203219, -183892,
			},
			PhaseRangeLocks: []uint8{
				0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf,
			},
			HalfCycles: []bool{
				false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
			},
		},
	},
}
