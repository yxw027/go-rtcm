package main

import (
//	"bufio"
	"fmt"
	"github.com/geoscienceaustralia/go-rtcm/rtcm3"
	"github.com/geoscienceaustralia/go-rtcm/orm"
	"github.com/umeat/go-ntrip/ntrip"
	"github.com/jinzhu/gorm"
//	"os"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&orm.Observation{})
	db.AutoMigrate(&orm.SatelliteData{})
	db.AutoMigrate(&orm.SignalData{})

//	r, _ := os.Open("../rtcm3/data/1077_frame.bin")
//	br := bufio.NewReader(r)
//	frame, _ := rtcm3.DeserializeFrame(br)
//	d := rtcm3.DeserializeMessage1077(frame.Payload)

	client, err := ntrip.NewClient("https://streams.auscors.geops.team/GAT000AUS0")
	resp, err := client.Connect()
	if err != nil || resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode, err)
	}

	scanner := rtcm3.NewScanner(resp.Body)
	for frame, err := scanner.NextFrame(); err == nil; frame, err = scanner.NextFrame() {
		switch frame.MessageNumber() {
		case 1077, 1087, 1097, 1107, 1117, 1127:
			obs, _ := orm.ObservationMsm7(rtcm3.DeserializeMessageMsm7(frame.Payload))
			db.Create(&obs)
		}
	}
	panic(err)
}
