package rtcm_test

import (
    "testing"
    "../pkgs/rtcm"
    "fmt"
    "github.com/google/go-cmp/cmp"

)

func TestMessage1004(t *testing.T) {
     header := rtcm.Rtcm3GpsObservationHeader{MessageNumber:1004,
                                              ReferenceStationId:0,
                                              Epoch:205514000,
                                              SynchronousGnss:true,
                                              SignalsProcessed:1,
                                              SmoothingIndicator:false,
                                              SmoothingInterval:0}

    satelliteData := []rtcm.Rtcm31004SatelliteData{{SatelliteId:15,
                                                    L1CodeIndicator:false,
                                                    L1Pseudorange:427272,
                                                    L1PhaseRange:27141,
                                                    L1LockTimeIndicator:127,
                                                    L1PseudorangeAmbiguity:74,
                                                    L1Cnr:197,
                                                    L2CodeIndicator:3,
                                                    L2PseudorangeDifference:-67,
                                                    L2PhaseRange:31228,
                                                    L2LockTimeIndicator:127,
                                                    L2Cnr:160}}



    msg := rtcm.Rtcm3Message1004{Rtcm3GpsObservationHeader:header,
                                 SatelliteData: satelliteData}

    deserialized_msg := rtcm.NewRtcm3Message1004(msg.Serialize())

    if !cmp.Equal(msg, deserialized_msg) {
        t.Errorf("Serialization->Deserialization not equal")
    }

}
