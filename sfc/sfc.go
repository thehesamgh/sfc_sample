package sfc

// ServiceFunctionChain is a struct which contains sequence of networkfunctions
type ServiceFunctionChain struct {
	services []interface{} //[]networkfunction.INetworkFunction
}

// NewSFC creates a new ServiceFunctionChain.
func NewSFC(nf []interface{}) (sfc ServiceFunctionChain) {
	return ServiceFunctionChain{nf}
}

//RoutePacketThroughSFC passes a packet through all service functions
// func (sfc *ServiceFunctionChain) RoutePacketThroughSFC(hey string) {
// 	for _, sf := range sfc.services {
// 		hey = sf.(networkfunction.INetworkFunction).ProcessPacket(hey)
// 		//sf.(networkfunction.IFirewall).AddRule(hey)
// 	}
// }

//RoutePacketThroughSFC passes a packet through all service functions
// func (sfc *ServiceFunctionChain) RoutePacketThroughSFC(p string) {
// 	sfc.services[i].(networkfunction.INetworkFunction).ProcessPacket(" ")
// }
