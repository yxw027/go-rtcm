package main

import (
	"fmt"
    "github.com/geoscienceaustralia/go-rtcm/orm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

    msg := orm.Observation{SatelliteData: []orm.SatelliteData{}}
    db.First(&msg).Preload("SignalData").Related(&msg.SatelliteData)
    fmt.Printf("%+v\n", msg)
}
