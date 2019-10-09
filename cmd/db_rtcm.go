package main

import (
	"fmt"
	"os"
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/geoscienceaustralia/go-rtcm/orm"
)

func main() {
	password := os.Getenv("DB_PASSWORD")
	db, err := gorm.Open("postgres", "host=rtcmdb.c76tte2hbd9p.ap-southeast-2.rds.amazonaws.com port=5432 user=postgres dbname=rtcmdb password=" + password)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

    msg := orm.Observation{SatelliteData: []orm.SatelliteData{}}
    db.First(&msg).Preload("SignalData").Related(&msg.SatelliteData)
    fmt.Printf("%+v\n", msg)
}
