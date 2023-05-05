package main

import (
	fabric "fabric/web"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file: ", err)
		return
	}

	//Parse chaincode functions as an array of strings
	chaincodeFunctions := []string{os.Getenv("CHAINCODE_FUNCTIONS")}

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
		return // Add this line to prevent further execution in case of an error
	}
	fabric.Serve(fabric.OrgSetup(*orgSetup))

}
