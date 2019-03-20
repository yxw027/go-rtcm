package data

import (
	"github.com/geoscienceaustralia/go-rtcm/rtcm3"
)

var Message1125 = rtcm3.Message1125{
	MessageMsm5: rtcm3.MessageMsm5{
		Header: rtcm3.MsmHeader{
			MessageNumber:          0x465,
			ReferenceStationId:     0x0,
			Epoch:                  0x1a601ef8,
			MultipleMessageBit:     true,
			Iods:                   0x0,
			Reserved:               0x0,
			ClockSteeringIndicator: 0x0,
			ExternalClockIndicator: 0x0,
			SmoothingIndicator:     false,
			SmoothingInterval:      0x0,
			SatelliteMask:          0xf340000820000000,
			SignalMask:             0x40040000,
			CellMask:               0x37ffa,
		},
		SatelliteData: rtcm3.SatelliteDataMsm57{
			RangeMilliseconds: []uint8{
				0x7b, 0x87, 0x80, 0x7c, 0x7b, 0x7c, 0x7c, 0x52, 0x48,
			},
			Extended: []uint8{
				0xd, 0x0, 0x0, 0x0, 0xb, 0x0, 0x0, 0x0, 0x0,
			},
			Ranges: []uint16{
				0x3e7, 0x1b5, 0x18a, 0x208, 0x230, 0x287, 0x277, 0xb3, 0x3f7,
			},
			PhaseRangeRates: []int16{
				3, -10, -4, 2, 228, -7, 123, 369, 20,
			},
		},
		SignalData: rtcm3.SignalDataMsm5{
			Pseudoranges: []int16{
				7277, 6897, 6907, 3563, 3385, -3591, -3789, -6929, -7160, 4308, 4118, -7134, -7176, 7944, 2195,
			},
			PhaseRanges: []int32{
				230963, 217111, 218332, 109384, 103261, -116238, -121388, -227658, -236891, 140751, 137382, -231990, -232621, 248319, 73024,
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
			PhaseRangeRates: []int16{
				2202, 2213, 3700, 4271, 3958, 1477, 1594, 3887, 4071, -3780, -3560, 3626, 3736, 4499, -3353,
			},
		},
	},
}
