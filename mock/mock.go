package fabric

import (
	"fmt"
	"net/http"
)

type IntMockOrgSetup interface {
}
type MockOrgSetup struct {
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
	Gateway            MockGateway
}

func (mos *MockOrgSetup) Query(w http.ResponseWriter, r *http.Request, inter ...IntMockOrgSetup) {
	fmt.Fprintf(w, "Response: did doc")
}

func (mos *MockOrgSetup) Store(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement mock Store function
}

type MockGateway struct{}

func (mg *MockGateway) GetNetwork(channelID string) *MockNetwork {
	return &MockNetwork{}
}

type MockNetwork struct{}

func (mn *MockNetwork) GetContract(name string) *MockContract {
	return &MockContract{}
}

type MockContract struct{}

func (mc *MockContract) EvaluateTransaction(name string, args ...string) ([]byte, error) {
	return []byte("mock response"), nil
}
