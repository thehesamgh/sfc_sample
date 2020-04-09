package networkfunction

import (
	"fmt"
	networkfunction "sfc_platform/network_function"
)

// TestFirewall is struct for Firewall
type Firewall struct {
	name string
	next networkfunction.INetworkFunction
}

// NewTestFirewall creates a new Firewall.
func NewFirewall(name string, next networkfunction.INetworkFunction) Firewall {
	return Firewall{name, next}
}

// AddRule is for adding rule in Firewall
func (tf Firewall) AddRule(ip string) (err error) {
	fmt.Printf("%s is Adding a Rule.\n", tf.name)
	return nil
}

//Run runs Firewall
func (tf Firewall) Run() (pid int, err error) {
	fmt.Printf("%s is Run.\n", tf.name)

	pid = 10
	err = nil
	return
}

//Kill kills Firewall
func (tf Firewall) Kill() (pid int, err error) {
	fmt.Printf("%s is Kill.\n", tf.name)
	return 10, nil
}

// ClearPolicies implemetation for Firewall
func (tf Firewall) ClearPolicies() (err error) {
	fmt.Printf("%s is Clearing Policies.\n", tf.name)

	return nil
}

// ProcessPacket implemetation for packet processing
func (tf Firewall) ProcessPacket(packet string) /*string*/ {
	fmt.Printf("%s is Processing Packet: \"%v\".\n", tf.name, packet)
	tf.AddRule(packet)

	if tf.next != nil {
		tf.next.ProcessPacket(packet)
	}
}

//Kill kills Firewall
func (tf Firewall) ApplyAction() {
	return
}

//Kill kills Firewall
func (tf Firewall) DecisionMaking() []networkfunction.Action {
	var temp []networkfunction.Action
	return temp
}
