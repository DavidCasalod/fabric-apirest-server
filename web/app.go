package fabric

import (
	"fmt"
	"net/http"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"google.golang.org/grpc"
)

type proposalBuilder struct {
	channelName     string
	chaincodeName   string
	transactionName string
	transient       map[string][]byte
	endorsingOrgs   []string
	args            [][]byte
}
type CommitInt interface {
	TransactionID() string
}
type TransactionInt interface {
	Submit(opts ...grpc.CallOption) (CommitInt, error)
	Result() []byte
}
type ProposalInt interface {
	Endorse(opts ...grpc.CallOption) (TransactionInt, error)
}

type ContractInt interface {
	EvaluateTransaction(name string, args ...string) ([]byte, error)
	NewProposal(transactionName string, options ...client.ProposalOption) (ProposalInt, error)
}

type NetworkInt interface {
	GetContract(chaincodeName string) ContractInt
}

type GatewayInt interface {
	GetNetwork(name string) NetworkInt
	//Connect(id *identity.X509Identity, options ...func(gateway *client.Gateway) error) (*client.Gateway, error)
}

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
	Gatewaytest        GatewayInt
	Gateway            *client.Gateway
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
