package data

import (
	"github.com/geoscienceaustralia/go-rtcm/rtcm3"
)

var Message1122 = rtcm3.Message1122{
	MessageMsm2: rtcm3.MessageMsm2{
		Header: rtcm3.MsmHeader{
			MessageNumber:          0x462,
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
		SatelliteData: rtcm3.SatelliteDataMsm123{
			Ranges: []uint16{
				0x2b5, 0x83, 0x58, 0xd6, 0xfd, 0x155, 0x144, 0x381, 0x2c5,
			},
		},
		SignalData: rtcm3.SignalDataMsm2{
			PhaseRanges: []int32{
				156539, 142687, 143901, 34957, 28834, -190663, -195813, 222326, 213094, 66321, 62952, 217938, 217308, 174091, -1391,
			},
			PhaseRangeLocks: []uint8{
				0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf,
			},
			HalfCycles: []bool{
				false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
			},
		},
	},
}
