package sfc

import (
	networkfunction "sfc_platform/network_function"
)

// Handler defined a handler for handling a given handleID.
type Handler interface {
	Handle(networkfunction.INetworkFunction)
}

type handler struct {
	name     string
	next     Handler
	handleID int
}

// NewHandler returns a new Handler.
func NewHandler(name string, next Handler, handleID int) Handler {
	return &handler{name, next, handleID}
}

// Handle handles a given handleID.
func (h *handler) Handle(sf networkfunction.INetworkFunction) {
	//if h.handleID == handleID {
	//return h.name + " handled " + strconv.Itoa(handleID)
	sf.ProcessPacket("")
	//h.next.Handle()
	//return
	//}
	//return h.next.Handle(handleID)
}
