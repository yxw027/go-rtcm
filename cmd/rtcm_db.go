package main

import (
	"github.com/geoscienceaustralia/go-rtcm/orm"
	"bufio"
	"fmt"
	"github.com/geoscienceaustralia/go-rtcm/rtcm3"
	"os"
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetSatIDs(satMask uint64) (ids []int) {
	for i := 64; i > 0; i-- {
		if (satMask >> uint64(i-1)) & 0x1 == 1 {
			ids = append(ids, i)
		}
	}
	return ids
}

func GetSigIDs(sigMask uint32) (ids []int) {
	for i := 32; i > 0; i-- {
		if (sigMask >> uint32(i-1)) & 0x1 == 1 {
			ids = append(ids, i)
		}
	}
	return ids
}

func Itob(v uint64) bool {
	if v == 0 {
		return false
	}
	return true
}

func GetCells(cellMask uint64, length int) (cells []bool) {
	for i := 0; i < length; i++ {
		cells = append([]bool{Itob((cellMask >> uint(i)) & 0x1)}, cells...)
	}
	return cells
}

func main() {
	r, _ := os.Open("../rtcm3/data/1077_frame.bin")
	br := bufio.NewReader(r)
	frame, _ := rtcm3.DeserializeFrame(br)
	d := rtcm3.DeserializeMessage1077(frame.Payload)

	obs := orm.Observation{
		MessageNumber: d.MessageNumber,
		ReferenceStationId: d.ReferenceStationId,
		Epoch: d.Epoch,
		Iods: d.Iods,
		Reserved: d.Reserved,
		ClockSteeringIndicator: d.ClockSteeringIndicator,
		ExternalClockIndicator: d.ExternalClockIndicator,
		SmoothingIndicator: d.SmoothingIndicator,
		SmoothingInterval: d.SmoothingInterval,
		SatelliteData: []orm.SatelliteData{},
	}

	satIDs := GetSatIDs(d.SatelliteMask)
	sigIDs := GetSigIDs(d.SignalMask)
	cellIDs := GetCells(d.CellMask, len(satIDs) * len(sigIDs))
	cellPos := 0
	sigPos := 0

	for x, satId := range satIDs {
		satData := orm.SatelliteData{
			SatelliteID: satId,
			RangeMilliseconds: d.SatelliteData.RangeMilliseconds[x],
			Extended: d.SatelliteData.Extended[x],
			Ranges: d.SatelliteData.Ranges[x],
			PhaseRangeRates: d.SatelliteData.PhaseRangeRates[x],
			SignalData: []orm.SignalData{},
		}
		for _, sigID := range sigIDs {
			if cellIDs[cellPos] {
				satData.SignalData = append(satData.SignalData, orm.SignalData{
					SignalID: sigID,
					Pseudoranges: d.SignalData.Pseudoranges[sigPos],
					PhaseRanges: d.SignalData.PhaseRanges[sigPos],
					PhaseRangeLocks: d.SignalData.PhaseRangeLocks[sigPos],
					HalfCycles: d.SignalData.HalfCycles[sigPos],
					Cnrs: d.SignalData.Cnrs[sigPos],
					PhaseRangeRates: d.SignalData.PhaseRangeRates[sigPos],
				})
				sigPos ++
			}
			cellPos ++
		}
		obs.SatelliteData = append(obs.SatelliteData, satData)
	}

	fmt.Printf("%+v\n\n%+v\n", d, obs)

	db, err := gorm.Open("postgres", "host=rtcmdb.c76tte2hbd9p.ap-southeast-2.rds.amazonaws.com port=5432 user=postgres dbname=rtcmdb password=w0wjonathanisCOOL")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&orm.Observation{})
	db.AutoMigrate(&orm.SatelliteData{})
	db.AutoMigrate(&orm.SignalData{})

	db.Create(&obs)
	//https://mindbowser.com/golang-go-with-gorm-2/
}
