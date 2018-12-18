package main

import (
    "github.com/geoscienceaustralia/go-rtcm/pkgs/rtcm"
    "github.com/umeat/go-ntrip/pkgs/ntrip"
    "fmt"
    "os"
//    "time"
)

func main() {
    r, _ := ntrip.Client(os.Args[1], os.Args[2], os.Args[3])
    scanner := rtcm.NewScanner(r)
    msg, err := scanner.Next()
    for ; err == nil; msg, err = scanner.Next() {
//        switch int(rtcm.GetMessageNumber(msg)) {
//            case 1077, 1097, 1117, 1087, 1127, 1001, 1002, 1003, 1004, 1009, 1010, 1011, 1012:
//                fmt.Println(rtcm.GetMessageNumber(msg), time.Now().UTC().Sub(msg.(rtcm.Rtcm3Observable).Time()))
//        }
        fmt.Println(rtcm.NewRtcm3Frame(msg).Serialize(), "\n")
    }

    fmt.Println(err)
}
