package data

import (
	"github.com/geoscienceaustralia/go-rtcm/rtcm3"
)

var Message1121 = rtcm3.Message1121{
	MessageMsm1: rtcm3.MessageMsm1{
		Header: rtcm3.MsmHeader{
			MessageNumber:          0x461,
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
		SignalData: rtcm3.SignalDataMsm1{
			Pseudoranges: []int16{
				4951, 4571, 4582, 1237, 1060, -5917, -6115, 7133, 6902, 1982, 1792, 6926, 6884, 5625, -131,
			},
		},
	},
}
