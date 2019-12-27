package networkfunction

import (
	"github.com/google/gopacket"
)

//INetworkFunction is iterface for network function
type INetworkFunction interface {
	ProcessPacket(packet string)
	DecisionMaking() []Action
	ApplyAction()
	Kill() (int, error)
}

// IFirewall is inteface for firewall function
type IFirewall interface {
	INetworkFunction
	AddRule(string) error
	Run() (int, error)
	ClearPolicies() error
}

// IIDS is inteface for IDS function
type IIDS interface {
	INetworkFunction
	DetectMalware()
}

// Action is construct for network actions e.g. drop, forward, route and so on.
type Action struct {
}

// Drop drops a packet and do nothing
func (a Action) Drop(packet *gopacket.Packet) {

}

// Forward forwards a packet and do nothing
func (a Action) Forward(packet *gopacket.Packet, next INetworkFunction) {

}
