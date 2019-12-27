package main

import (
	"strconv"
	"testing"
)

func TestServiceFunctionChaining(t *testing.T) {
	//load configs from yml file
	configDir = &defaultConfDir

	//set flags for this application, so config file path can be changed via a command line switch
	setFlags()

	chainNetworkFunctionsList := cfg.SSFC.Info

	firstNetworkFunctionOfChain := CreateChain(chainNetworkFunctionsList)

	// Pass sample packets to the chain
	for i := 1; i <= 3; i++ {
		firstNetworkFunctionOfChain.ProcessPacket("Hey I'm packet " + strconv.Itoa(i))
	}
}
