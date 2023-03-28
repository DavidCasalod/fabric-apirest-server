package main

import (
	"encoding/json"
	fabric "fabric/web"
	"fmt"
	"os"
)

func main() {
	//Parse chaincode functions as an array of strings
	chaincodeFunctionsJSON := os.Getenv("CHAINCODE_FUNCTIONS")
	var chaincodeFunctions []string
	err := json.Unmarshal([]byte(chaincodeFunctionsJSON), &chaincodeFunctions)
	if err != nil {
		panic(err)
	}
	//Initialize setup for Org1
	orgConfig := fabric.OrgSetup{
		OrgName:            os.Getenv("ORG_NAME"),
		MSPID:              os.Getenv("MSPID"),
		CertPath:           os.Getenv("CERTPATH"),
		KeyPath:            os.Getenv("KEYPATH"),
		TLSCertPath:        os.Getenv("TLSCERTPATH"),
		PeerEndpoint:       os.Getenv("PEERENDPOINT"),
		ChaincodeName:      os.Getenv("CHAINCODE_NAME"),
		ChannelId:          os.Getenv("CHANNEL_ID"),
		ChaincodeFunctions: chaincodeFunctions,
		GatewayPeer:        os.Getenv("GATEWAYPEER"),
	}

	orgSetup, err := fabric.Initialize(orgConfig)
	if err != nil {
		fmt.Println("Error initializing setup for Org1: ", err)
	}
	fabric.Serve(fabric.OrgSetup(*orgSetup))
}
