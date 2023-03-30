package fabric

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric/msp"
	"github.com/stretchr/testify/assert"
)

// The first test case tests a successful Query function call with a valid DID ID parameter. The mock Gateway, Network, and Contract objects
// are set up to return expected values for the EvaluateTransaction call.
// The second test case tests an invalid DID format error. The query parameter

func TestQuery_Success(t *testing.T) {
	// Create a mock Gateway for testing purposes
	mockContract := &MockContract{}
	mockNetwork := &MockNetwork{
		Contract: mockContract,
	}
	mockGateway := &MockGateway{
		MockNetwork: mockNetwork,
	}
	// setup
	setup := OrgSetup{
		OrgName:            "Org1",
		MSPID:              "Org1MSP",
		CryptoPath:         "/path/to/crypto",
		CertPath:           "/path/to/cert",
		KeyPath:            "/path/to/key",
		TLSCertPath:        "/path/to/tls-cert",
		PeerEndpoint:       "peer0.org1.example.com:7051",
		GatewayPeer:        "localhost:7051",
		ChaincodeName:      "mycc",
		ChannelId:          "mychannel",
		ChaincodeFunctions: []string{"queryDID"},
		Gateway:            mockGateway,
	}

	// create a fake HTTP request
	req, err := http.NewRequest("GET", "/?didId=did:fabric:abc123", nil)
	assert.NoError(t, err)

	// create a fake HTTP response recorder
	rr := httptest.NewRecorder()

	// invoke the Query function
	setup.Query(rr, req)

	// check the response
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "Response: some data", rr.Body.String()) // replace "some data" with the expected response
}

func TestQuery_InvalidDIDFormat(t *testing.T) {
	// Create a mock Gateway for testing purposes
	mockContract := &MockContract{}
	mockNetwork := &MockNetwork{
		Contract: mockContract,
	}
	mockGateway := &MockGateway{
		MockNetwork: mockNetwork,
	}
	// setup
	setup := OrgSetup{
		OrgName:            "Org1",
		MSPID:              "Org1MSP",
		CryptoPath:         "/path/to/crypto",
		CertPath:           "/path/to/cert",
		KeyPath:            "/path/to/key",
		TLSCertPath:        "/path/to/tls-cert",
		PeerEndpoint:       "peer0.org1.example.com:7051",
		GatewayPeer:        "localhost:7051",
		ChaincodeName:      "mycc",
		ChannelId:          "mychannel",
		ChaincodeFunctions: []string{"queryDID"},
		Gateway:            mockGateway,
	}

	// create a fake HTTP request
	req, err := http.NewRequest("GET", "/?didId=invalid", nil)
	assert.NoError(t, err)

	// create a fake HTTP response recorder
	rr := httptest.NewRecorder()

	// invoke the Query function
	setup.Query(rr, req)

	// check the response
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "invalid DID format")
}

func TestOrgSetup_Query_UnsupportedDIDMethod(t *testing.T) {
	// Create a mock Gateway for testing purposes
	mockContract := &MockContract{}
	mockNetwork := &MockNetwork{
		Contract: mockContract,
	}
	mockGateway := &MockGateway{
		MockNetwork: mockNetwork,
	}
	// Create an OrgSetup instance for testing
	orgSetup := OrgSetup{
		OrgName:            "TestOrg",
		MSPID:              "TestMSP",
		CryptoPath:         "/path/to/crypto",
		CertPath:           "/path/to/cert",
		KeyPath:            "/path/to/key",
		TLSCertPath:        "/path/to/tls/cert",
		PeerEndpoint:       "localhost:12345",
		GatewayPeer:        "peer0",
		ChaincodeName:      "TestChaincode",
		ChannelId:          "TestChannel",
		ChaincodeFunctions: []string{"queryFunction"},
		Gateway:            mockGateway,
	}

	// Set up the test request with an unsupported DID method query parameter
	req, err := http.NewRequest("GET", "/path/to/query?didId=did:notsupported:1234", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the Query function
	handler := http.HandlerFunc(orgSetup.Query)
	handler.ServeHTTP(rr, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, rr.Code)
	// Check the response body
	assert.Contains(t, rr.Body.String(), "unsupported DID method")
}

func TestOrgSetup_Query_ContractError(t *testing.T) {
	// Create a mock Gateway for testing purposes
	mockContract := &MockContract{}
	mockNetwork := &MockNetwork{
		Contract: mockContract,
	}
	mockGateway := &MockGateway{
		MockNetwork: mockNetwork,
	}

	// Create an OrgSetup instance for testing
	orgSetup := OrgSetup{
		OrgName:            "TestOrg",
		MSPID:              "TestMSP",
		CryptoPath:         "/path/to/crypto",
		CertPath:           "/path/to/cert",
		KeyPath:            "/path/to/key",
		TLSCertPath:        "/path/to/tls/cert",
		PeerEndpoint:       "localhost:12345",
		GatewayPeer:        "peer0",
		ChaincodeName:      "TestChaincode",
		ChannelId:          "TestChannel",
		ChaincodeFunctions: []string{"queryFunction"},
		Gateway:            mockGateway,
	}
	// Set up a mock contract error response
	mockError := fmt.Errorf("mock contract error")
	didId := "did:fabric:1234"
	mockContract.EvaluateTransaction("queryFunction", []byte(didId))
	mockNetwork.GetContract("TestChaincode")
	mockNetwork.GetNetwrok("TestChannel")
	// Set up the test request with the DID ID query parameter
	req, err := http.NewRequest("GET", "/path/to/query?didId=did:fabric:1234", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the Query function
	handler := http.HandlerFunc(orgSetup.Query)
	handler.ServeHTTP(rr, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, rr.Code)
	// Check the response body
	assert.Contains(t, rr.Body.String(), mockError.Error())
}

type MockClient struct {
	*MockGateway
}

func (c *MockClient) GetGateway() (client.Gateway, error) {
	return c.MockGateway, nil
}

type MockGateway struct {
	*MockNetwork
}

func (g *MockGateway) GetNetwork(channelId string) client.Network {
	return &MockNetwork{}
}

func (g *MockGateway) GetIdentity() *msp.Identity {
	// Implement GetIdentity() method
	return nil
}

func (g *MockGateway) SubmitTransaction(ctx context.Context, name string, args [][]byte) ([]byte, error) {
	// Implement SubmitTransaction() method
	return nil, nil
}

func (g *MockGateway) EvaluateTransaction(ctx context.Context, name string, args [][]byte) ([]byte, error) {
	// Implement EvaluateTransaction() method
	return nil, nil
}

func (g *MockGateway) Close() {
	// Implement Close() method
}

type MockNetwork struct {
	Contract *MockContract
}

func (n *MockNetwork) GetContract(name string) (client.Contract, error) {
	if n.Contract == nil {
		n.Contract = &MockContract{}
	}
	return n.Contract, nil
}

func (n *MockNetwork) GetNetwrok(name string) (client.Contract, error) {
	if n.Contract == nil {
		n.Contract = &MockContract{}
	}
	return n.Contract, nil
}

type MockContract struct{}

func (c *MockContract) EvaluateTransaction(name string, args ...[]byte) ([]byte, error) {
	if name == "queryFunction" && len(args) == 1 && string(args[0]) == "did:fabric:1234" {
		// Return the expected response for the given inputs
		return []byte("mock response"), nil
	}
	// Return an error for other inputs
	return nil, fmt.Errorf("unexpected inputs")
}
