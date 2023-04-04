package fabric

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/stretchr/testify/assert"
)

const (
	didDoc = `{"@context":["https://www.w3.org/ns/did/v1"],"id":"did:fabric:123","verificationMethod":[{"id":"did:fabric:123#key1","type":"Ed25519VerificationKey2018","controller":"did:fabric:123","publicKeyBase58":"2Qfyg1W6ySFGmE57Kj3wFucZ8W4Z4h4jL9Rny1NYQzN8"}],"service":[{"id":"did:fabric:123#hub","type":"Messaging","serviceEndpoint":"http://localhost:10000"}]}`
)

func TestStore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	t.Run("test query DID success", func(t *testing.T) {
		// Create a mock Gateway for testing purposes
		mockGateway := NewMockGatewayInt(ctrl)
		mockContract := NewMockContractInt(ctrl)
		mockNetwork := NewMockNetworkInt(ctrl)
		mockProposal := NewMockProposalInt(ctrl)
		mockTransaction := NewMockTransactionInt(ctrl)
		mockCommit := NewMockCommitInt(ctrl)
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
			ChaincodeName:      "TestChaincode",
			ChannelId:          "TestChannel",
			ChaincodeFunctions: []string{"queryDID", "storeDID"},
			Gatewaytest:        mockGateway,
		}

		// Set up the mock Gateway to return the mock Network and mock Contract
		mockGateway.EXPECT().GetNetwork(setup.ChannelId).Return(mockNetwork)
		mockNetwork.EXPECT().GetContract(setup.ChaincodeName).Return(mockContract)
		mockContract.EXPECT().NewProposal(setup.ChaincodeFunctions[1], client.WithArguments(didDoc)).Return(mockProposal, nil)
		mockProposal.EXPECT().Endorse().Return(mockTransaction, nil)
		mockTransaction.EXPECT().Submit().Return(mockCommit, nil)
		mockTransaction.EXPECT().Result().Return([]byte("result ok"))
		mockCommit.EXPECT().TransactionID().Return("txnId")
		// Set up the HTTP request and response recorder
		req, err := http.NewRequest("POST", "/store?didDoc="+url.QueryEscape(didDoc), nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		setup.Store(rr, req)

		// Check the response status code and body
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Contains(t, rr.Body.String(), "txnId")
		assert.Contains(t, rr.Body.String(), "result ok")
	})
}

// 	t.Run("Test Query Unsupported DID method ", func(t *testing.T) {
// 		// Create a mock Gateway for testing purposes
// 		mockGateway := NewMockGatewayInt(ctrl)

// 		// Create an OrgSetup instance for testing
// 		orgSetup := OrgSetup{
// 			OrgName:            "TestOrg",
// 			MSPID:              "TestMSP",
// 			CryptoPath:         "/path/to/crypto",
// 			CertPath:           "/path/to/cert",
// 			KeyPath:            "/path/to/key",
// 			TLSCertPath:        "/path/to/tls/cert",
// 			PeerEndpoint:       "localhost:12345",
// 			GatewayPeer:        "peer0",
// 			ChaincodeName:      "TestChaincode",
// 			ChannelId:          "TestChannel",
// 			ChaincodeFunctions: []string{"queryFunction"},
// 			Gatewaytest:        mockGateway,
// 		}

// 		// Create a test request with an unsupported DID method
// 		req, err := http.NewRequest("GET", "/query?didId=did:unsupported:1234", nil)
// 		assert.NoError(t, err)

// 		// Create a response recorder
// 		rr := httptest.NewRecorder()

// 		// Call the Query function with the mock setup and test request
// 		orgSetup.Query(rr, req)
// 		expectedResponse := "Unsupported DID method"
// 		// check the response
// 		assert.Equal(t, http.StatusBadRequest, rr.Code)
// 		assert.Equal(t, expectedResponse, strings.TrimSpace(rr.Body.String()))

// 	})

// 	t.Run("Test Query Invalid Format", func(t *testing.T) {
// 		// Create a mock Gateway for testing purposes
// 		mockGateway := NewMockGatewayInt(ctrl)
// 		// setup
// 		setup := OrgSetup{
// 			OrgName:            "Org1",
// 			MSPID:              "Org1MSP",
// 			CryptoPath:         "/path/to/crypto",
// 			CertPath:           "/path/to/cert",
// 			KeyPath:            "/path/to/key",
// 			TLSCertPath:        "/path/to/tls-cert",
// 			PeerEndpoint:       "peer0.org1.example.com:7051",
// 			GatewayPeer:        "localhost:7051",
// 			ChaincodeName:      "mycc",
// 			ChannelId:          "mychannel",
// 			ChaincodeFunctions: []string{"queryDID"},
// 			Gatewaytest:        mockGateway,
// 		}

// 		// create a fake HTTP request with an invalid DID format
// 		req, err := http.NewRequest("GET", "/query?didId=invalidDIDFormat", nil)
// 		assert.NoError(t, err)

// 		// Create a response recorder
// 		rr := httptest.NewRecorder()

// 		// Call the Query function with setup and test request
// 		setup.Query(rr, req)

// 		// check the response
// 		assert.Equal(t, http.StatusBadRequest, rr.Code)
// 		assert.Contains(t, rr.Body.String(), "Invalid DID format")
// 	})

// 	t.Run("TestQuery EvaluateTransaction Error - contract error", func(t *testing.T) {

// 		// Create a mock Gateway for testing purposes
// 		mockGateway := NewMockGatewayInt(ctrl)
// 		mockContract := NewMockContractInt(ctrl)
// 		mockNetwork := NewMockNetworkInt(ctrl)

// 		// Create an OrgSetup instance for testing
// 		orgSetup := OrgSetup{
// 			OrgName:            "TestOrg",
// 			MSPID:              "TestMSP",
// 			CryptoPath:         "/path/to/crypto",
// 			CertPath:           "/path/to/cert",
// 			KeyPath:            "/path/to/key",
// 			TLSCertPath:        "/path/to/tls/cert",
// 			PeerEndpoint:       "localhost:12345",
// 			GatewayPeer:        "peer0",
// 			ChaincodeName:      "TestChaincode",
// 			ChannelId:          "TestChannel",
// 			ChaincodeFunctions: []string{"queryFunction"},
// 			Gatewaytest:        mockGateway,
// 		}

// 		// Set up a mock contract error response
// 		mockGateway.EXPECT().GetNetwork(gomock.Any()).Return(mockNetwork)
// 		mockNetwork.EXPECT().GetContract(gomock.Any()).Return(mockContract)
// 		mockError := fmt.Errorf("Couldn't evaluate transaction for didID ")
// 		mockContract.EXPECT().EvaluateTransaction(gomock.Any(), gomock.Any()).Return(nil, mockError)

// 		// Set up the test request with the DID ID query parameter
// 		req, err := http.NewRequest("GET", "/path/to/query?didId=did:fabric:1234", nil)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		// Create a ResponseRecorder to capture the response
// 		rr := httptest.NewRecorder()

// 		// Call the Query function
// 		handler := http.HandlerFunc(orgSetup.Query)
// 		handler.ServeHTTP(rr, req)

// 		// Check the response status code
// 		assert.Equal(t, http.StatusInternalServerError, rr.Code)
// 		// Check the response body
// 		assert.Contains(t, rr.Body.String(), mockError.Error())
// 	})
// }
