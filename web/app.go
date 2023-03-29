package fabric

import (
	"fmt"
	"net/http"

	"github.com/hyperledger/fabric-gateway/pkg/client"
)

// OrgSetup contains organization's config to interact with the network.
type OrgSetup struct {
	OrgName            string
	MSPID              string
	CryptoPath         string
	CertPath           string
	KeyPath            string
	TLSCertPath        string
	PeerEndpoint       string
	GatewayPeer        string
	ChaincodeName      string
	ChannelId          string
	ChaincodeFunctions []string
	Gateway            client.Gateway
}

// Serve starts http web server.
func Serve(setups OrgSetup) {
	http.HandleFunc("/query", setups.Query)
	http.HandleFunc("/store", setups.Store)
	fmt.Println("Listening (http://localhost:3000/)...")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println(err)
	}
}
