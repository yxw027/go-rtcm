package main

import (
    "./rtcm"
    "github.com/umeat/go-ntrip/ntrip"
    "net/url"
    "fmt"
    "os"
    "time"
)

func GpsTime(e uint32) time.Duration {
    now := time.Now().UTC()
    sow := now.Truncate(time.Hour * 24).AddDate(0, 0, -int(now.Weekday()))
    tow := time.Duration(e) * time.Millisecond
    return now.Sub(sow.Add(-(18 * time.Second)).Add(tow))
}

func GlonassTime(e uint32) time.Duration {
    now := time.Now().UTC()
    sow := now.Truncate(time.Hour * 24).AddDate(0, 0, -int(now.Weekday()))
    dow := int((e >> 27) & 0x7)
    tod := time.Duration(e & 0x7FFFFFF) * time.Millisecond
    return now.Sub(sow.AddDate(0, 0, dow).Add(tod).Add(-(3 * time.Hour)))
}

func GlonassTimeShort(e uint32) time.Duration {
    now := time.Now().UTC()
    dow := now.Truncate(time.Hour * 24)
    tod := time.Duration(e) * time.Millisecond
    return now.Sub(dow.Add(tod).Add(-(3 * time.Hour)))
}

func main() {
    c, _ := url.Parse(os.Args[1])
    r, _ := ntrip.Client(c, os.Args[2], os.Args[3])

    e := rtcm.Scan(r, func(msg rtcm.Rtcm3Message) {
        switch int(msg.Number()) {
            case 1077, 1097, 1117, 1087, 1127, 1001, 1002, 1003, 1004, 1009, 1010, 1011, 1012:
                fmt.Println(msg.Number(), msg.(rtcm.Rtcm3Observable).Time())
        }
    })
    fmt.Println(e)
}
