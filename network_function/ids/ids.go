package networkfunction

import (
	"fmt"
	networkfunction "sfc_platform/network_function"
	"strconv"
	"strings"
	"utilities"
)

// IDS is struct for IDS function
type IDS struct {
	name                               string
	destinationPacketsPerSecondCounter map[string]int
	sMalwareDetection                  IDetection
	sAnomalyDetection                  AnomalyDetection
	next                               networkfunction.INetworkFunction
}

// NewIDS creates a new IDS.
func NewIDS(name string, next networkfunction.INetworkFunction) networkfunction.IIDS {
	var ids IDS
	ids.destinationPacketsPerSecondCounter = make(map[string]int)
	ids.sMalwareDetection = MalwareDetection{}
	ids.sAnomalyDetection = AnomalyDetection{}
	ids.name = name
	ids.next = next
	return ids
}

// ProcessPacket implemetation for packet processing
func (ids IDS) ProcessPacket(packet string) /*string*/ {
	fmt.Printf("%s is Processing Packet: \"%v\".\n", ids.name, packet)

	Detection := false

	if strings.Contains(packet, "10") {
		Detection = true
	}
	if Detection {
		ids.DetectMalware()
	}
	if ids.next != nil {
		ids.next.ProcessPacket(packet)
	}
}

//Kill kills snort ids
func (ids IDS) Kill() (pid int, err error) {
	_, pid, err = utilities.RunCommandInBG(true, true, "kill", "-6", strconv.Itoa(pid))
	return
}

//Kill kills snort ids
func (ids IDS) ApplyAction() {
	return
}

//Kill kills snort ids
func (ids IDS) DecisionMaking() []networkfunction.Action {
	var temp []networkfunction.Action
	return temp
}

func (ids IDS) DetectMalware() {
	fmt.Printf("%v is detecting malware\n", ids.name)
}

//IDetection ...
type IDetection interface {
}

type MalwareDetection struct {
}

type AnomalyDetection struct {
}
