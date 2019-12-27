package main

import (
	"flag"
	"fmt"
	"log"
	config "sfc_platform/config"
	networkfunction "sfc_platform/network_function"
	firewall "sfc_platform/network_function/firewall"
	ids "sfc_platform/network_function/ids"
)

var (
	configDir      *string
	defaultConfDir = "config.yml"

	cfg *config.Config
)

func main() {
	//load configs from yml file
	configDir = &defaultConfDir

	//set flags for this application, so config file path can be changed via a command line switch
	setFlags()

	//get network function chain items from the loaded config file object
	chainNetworkFunctionsList := cfg.SSFC.Info

	VisualizeSFC(chainNetworkFunctionsList)
}

func setFlags() {
	configDir = flag.String("c", "config.yml", "Sets *.yml config dir. Default config dir is ./config.yml")

	flag.Parse()

	cfg = config.Set(*configDir)

	if !cfg.CheckIfConfigFileExists() {
		log.Fatal("Cannot file specified config file \"", *configDir, "\". You can set it using -c flag.")
	}
}

func NetworkFunctionBuilder(networkFunctionName, networkFunctionType string, next networkfunction.INetworkFunction) networkfunction.INetworkFunction {
	if networkFunctionType == "firewall" {
		return firewall.NewFirewall(networkFunctionName, next)
	} else if networkFunctionType == "ids" {
		return ids.NewIDS(networkFunctionName, next)
	}
	return nil
}

// CreateChain constructs chain of network functions and returns first network function
// This chain is created using chain of responsibility design pattern
func CreateChain(chainNetworkFunctionsList []config.NetworkFunctionsInfo) (firstItemOfChain networkfunction.INetworkFunction) {

	var next networkfunction.INetworkFunction
	for i := len(chainNetworkFunctionsList) - 1; i >= 0; i-- {
		networkFunction := chainNetworkFunctionsList[i]
		if i == len(chainNetworkFunctionsList)-1 {
			//last chain item doesn't have next item, so we use nil
			next = nil
		}
		next = NetworkFunctionBuilder(networkFunction.NameOfFunction, networkFunction.TypeOfFunction, next)
	}
	return next
}

func VisualizeSFC(sfcInfos []config.NetworkFunctionsInfo) {
	// Visualise the created chain
	fmt.Printf("Chain has been created: \n\t")
	for i, nf := range sfcInfos {
		fmt.Printf("%v", nf.NameOfFunction)

		if i == len(sfcInfos)-1 {
			fmt.Printf("\n\n")
		} else {
			fmt.Printf(" --> ")
		}
	}
}
