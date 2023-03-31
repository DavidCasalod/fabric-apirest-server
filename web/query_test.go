package fabric

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	mocks "scm.atosresearch.eu/ari/ledger_uself/ssi-ledgeruself-fabric/mocks/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestQuery(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	t.Run("test query DID success", func(t *testing.T) {
		// Create a mock Gateway for testing purposes

		mockGateway := mocks.NewMockGatewayInt(ctrl)

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
			Gatewaytest:        mockGateway,
		}

		// create a fake HTTP request
		req, err := http.NewRequest("GET", "/?didId=did:fabric:abc123", nil)
		assert.NoError(t, err)

		// create a fake HTTP response recorder
		rr := httptest.NewRecorder()

		// Call the Query function with the mock setup and test request
		setup.Query(rr, req)
		expectedResponse := "doc"
		// check the response
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, expectedResponse, rr.Body.String())
	})

	t.Run("Test Query Invalid Method", func(t *testing.T) {
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
			Gatewaytest:        mockGateway,
		}

		// Create a test request with an unsupported DID method
		req, err := http.NewRequest("GET", "/query?didId=did:unsupported:1234", nil)
		assert.NoError(t, err)

		// Create a response recorder
		rr := httptest.NewRecorder()

		// Call the Query function with the mock setup and test request
		orgSetup.Query(rr, req)
		expectedResponse := "unsupported DID method"
		// check the response
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, expectedResponse, rr.Body.String())
		// Call the Query function
		handler := http.HandlerFunc(orgSetup.Query)
		handler.ServeHTTP(rr, req)

		// Check the response status code
		assert.Equal(t, http.StatusOK, rr.Code)
		// Check the response body
		assert.Contains(t, rr.Body.String(), "unsupported DID method")

	})

	t.Run("Test Query Invalid Format", func(t *testing.T) {
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
			Gatewaytest:        mockGateway,
		}

		// create a fake HTTP request with an invalid DID format
		req, err := http.NewRequest("GET", "/query?didId=invalidDIDFormat", nil)
		assert.NoError(t, err)

		// Create a response recorder
		rr := httptest.NewRecorder()

		// Call the Query function with setup and test request
		setup.Query(rr, req)

		// check the response
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Contains(t, rr.Body.String(), "invalid DID format")
	})

	t.Run("TestQuery EvaluateTransaction Error - contract error", func(t *testing.T) {

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
			Gatewaytest:        mockGateway,
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
	})
}

// type MockGateway struct {
// 	*client.Gateway
// }

// func (g *MockGateway) GetNetwork(channelId string) *Network {
// 	return Network{MockNetwork}
// }

// type MockNetwork struct {
// 	*client.Network
// }

// func (n *MockNetwork) GetContract(name string) (*MockContract, error) {
// 	return &MockContract{}, nil
// }

// type MockContract struct {
// 	*client.Contract
// }

// func (c *MockContract) EvaluateTransaction(name string, args ...[]byte) ([]byte, error) {
// 	if name == "queryFunction" && len(args) == 1 && string(args[0]) == "did:fabric:1234" {
// 		// Return the expected response for the given inputs
// 		return []byte("mock response"), nil
// 	}
// 	// Return an error for other inputs
// 	return nil, fmt.Errorf("unexpected inputs")
// }
