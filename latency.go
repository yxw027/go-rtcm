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

func main() {
    c, _ := url.Parse(os.Args[1])
    r, _ := ntrip.Client(c, os.Args[2], os.Args[3])

    e := rtcm.Scan(r, func(msg rtcm.Rtcm3Message) {
        switch int(msg.Number()) {
            case 1077, 1097, 1117:
                fmt.Println(msg.Number(), GpsTime(msg.(*rtcm.Rtcm3MessageMsm7).Header.Epoch))
            case 1087:
                fmt.Println(msg.Number(), GlonassTime(msg.(*rtcm.Rtcm3MessageMsm7).Header.Epoch))
            case 1127:
                fmt.Println(msg.Number(), GpsTime(msg.(*rtcm.Rtcm3MessageMsm7).Header.Epoch) - (14 * time.Second))
            case 1001:
                fmt.Println(msg.Number(), GpsTime(msg.(*rtcm.Rtcm3Message1001).Header.Epoch))
            case 1002:
                fmt.Println(msg.Number(), GpsTime(msg.(*rtcm.Rtcm3Message1002).Header.Epoch))
            case 1003:
                fmt.Println(msg.Number(), GpsTime(msg.(*rtcm.Rtcm3Message1003).Header.Epoch))
            case 1004:
                fmt.Println(msg.Number(), GpsTime(msg.(*rtcm.Rtcm3Message1004).Header.Epoch))
        }
    })
    fmt.Println(e)
}
