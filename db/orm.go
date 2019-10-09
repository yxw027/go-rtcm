package main

import (
	"bufio"
	"fmt"
	"github.com/geoscienceaustralia/go-rtcm/rtcm3"
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type MSM7 struct { // Just MSM?
	//gorm.Model
	ObservationID int `gorm:"primary_key"`
	MessageNumber          uint16
	ReferenceStationId     uint16
	Epoch                  uint32 // Timestamp instead?
	Iods                   uint8 // Probably don't need this
	Reserved               uint8 // This might be needed later
	ClockSteeringIndicator uint8
	ExternalClockIndicator uint8
	SmoothingIndicator     bool
	SmoothingInterval      uint8
	SatelliteData []SatelliteData `gorm:"foreignkey:ObservationID"`
}

type SatelliteData struct {
	gorm.Model
	ObservationID     int
	SatelliteID       int // ???
	RangeMilliseconds uint8
	Extended          uint8
	Ranges            uint16
	PhaseRangeRates   int16
	//SignalData []SignalDataMsm7
}

type SignalDataMsm7 struct {
	SatDataID       int `gorm:"foreign_key"`
	SatelliteID     int // ???
	SignalID        int // ???
	Pseudoranges    int32
	PhaseRanges     int32
	PhaseRangeLocks uint16
	HalfCycles      bool
	Cnrs            uint16
	PhaseRangeRates int16
}

func main() {
	r, _ := os.Open("../rtcm3/data/1077_frame.bin")
	br := bufio.NewReader(r)
	frame, _ := rtcm3.DeserializeFrame(br)
	d := rtcm3.DeserializeMessage1077(frame.Payload)

	msg := MSM7{
		MessageNumber: d.MessageNumber,
		ReferenceStationId: d.ReferenceStationId,
		Epoch: d.Epoch,
		Iods: d.Iods,
		Reserved: d.Reserved,
		ClockSteeringIndicator: d.ClockSteeringIndicator,
		ExternalClockIndicator: d.ExternalClockIndicator,
		SmoothingIndicator: d.SmoothingIndicator,
		SmoothingInterval: d.SmoothingInterval,
		SatelliteData: []SatelliteData{},
	}

	for i := 0; i < len(d.SatelliteData.RangeMilliseconds); i++ {
		msg.SatelliteData = append(msg.SatelliteData, SatelliteData{
			SatelliteID: 1,
			RangeMilliseconds: d.SatelliteData.RangeMilliseconds[i],
			Extended: d.SatelliteData.Extended[i],
			Ranges: d.SatelliteData.Ranges[i],
			PhaseRangeRates: d.SatelliteData.PhaseRangeRates[i],
		})
	}

	//for i := 0; i < len(d.SignalData.Pseudoranges); i++ {
	//	msg.SatelliteData[0].SignalData = append(msg.SatelliteData[0].SignalData, SignalDataMsm7{
	//		SatelliteID: 1,
	//		SignalID: 1,
	//		Pseudoranges: d.SignalData.Pseudoranges[i],
	//		PhaseRanges: d.SignalData.PhaseRanges[i],
	//		PhaseRangeLocks: d.SignalData.PhaseRangeLocks[i],
	//		HalfCycles: d.SignalData.HalfCycles[i],
	//		Cnrs: d.SignalData.Cnrs[i],
	//		PhaseRangeRates: d.SignalData.PhaseRangeRates[i],
	//	})
	//}

	fmt.Printf("%+v\n", d)
	fmt.Printf("%+v\n", msg)
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&MSM7{})
	db.Create(&msg)
	//https://mindbowser.com/golang-go-with-gorm-2/
}
