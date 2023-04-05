package fabric

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	didDoc             = `{"@context":["https://www.w3.org/ns/did/v1"],"id":"did:fabric:123","verificationMethod":[{"id":"did:fabric:123#key1","type":"Ed25519VerificationKey2018","controller":"did:fabric:123","publicKeyBase58":"2Qfyg1W6ySFGmE57Kj3wFucZ8W4Z4h4jL9Rny1NYQzN8"}],"service":[{"id":"did:fabric:123#hub","type":"Messaging","serviceEndpoint":"http://localhost:10000"}]}`
	didDoc_unsupported = `{"@context":["https://www.w3.org/ns/did/v1"],"id":"did:other:1234","verificationMethod":[{"id":"did:fabric:123#key1","type":"Ed25519VerificationKey2018","controller":"did:fabric:123","publicKeyBase58":"2Qfyg1W6ySFGmE57Kj3wFucZ8W4Z4h4jL9Rny1NYQzN8"}],"service":[{"id":"did:fabric:123#hub","type":"Messaging","serviceEndpoint":"http://localhost:10000"}]}`
)

func TestStore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	t.Run("test query DID - Success", func(t *testing.T) {
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
		mockContract.EXPECT().NewProposal(gomock.Any(), gomock.Any()).Return(mockProposal, nil)
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
	t.Run("test store - Unsupported DID method", func(t *testing.T) {
		// Create a mock Gateway for testing purposes
		mockGateway := NewMockGatewayInt(ctrl)

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

		// Set up the HTTP request and response recorder

		req, err := http.NewRequest("POST", "/store?didDoc="+url.QueryEscape(didDoc_unsupported), nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		setup.Store(rr, req)

		// Check the response status code and body
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Contains(t, rr.Body.String(), "Unsupported DID method:")
	})
	t.Run("test store - Invalid DID format", func(t *testing.T) {
		// Create a mock Gateway for testing purposes
		mockGateway := NewMockGatewayInt(ctrl)

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

		// Set up the HTTP request and response recorder

		req, err := http.NewRequest("POST", "/store?didDoc=invalid", nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		setup.Store(rr, req)

		// Check the response status code and body
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Contains(t, rr.Body.String(), "Error creating txn proposal: invalid character 'i' ")
	})
	t.Run("test store - Proposal error", func(t *testing.T) {
		// Create a mock Gateway for testing purposes
		mockGateway := NewMockGatewayInt(ctrl)
		mockContract := NewMockContractInt(ctrl)
		mockNetwork := NewMockNetworkInt(ctrl)

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
		mockerror := errors.New("Proposal error")
		mockContract.EXPECT().NewProposal(gomock.Any(), gomock.Any()).Return(nil, mockerror)

		// Set up the HTTP request and response recorder
		req, err := http.NewRequest("POST", "/store?didDoc="+url.QueryEscape(didDoc), nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		setup.Store(rr, req)

		// Check the response status code and body
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Contains(t, rr.Body.String(), mockerror.Error())
	})
	t.Run("test store - Endorse error", func(t *testing.T) {

		// Create a mock Gateway for testing purposes
		mockGateway := NewMockGatewayInt(ctrl)
		mockContract := NewMockContractInt(ctrl)
		mockNetwork := NewMockNetworkInt(ctrl)
		mockProposal := NewMockProposalInt(ctrl)

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
		mockContract.EXPECT().NewProposal(gomock.Any(), gomock.Any()).Return(mockProposal, nil)
		mockerror := errors.New("Endorsment error")
		mockProposal.EXPECT().Endorse().Return(nil, mockerror)

		// Set up the HTTP request and response recorder
		req, err := http.NewRequest("POST", "/store?didDoc="+url.QueryEscape(didDoc), nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		setup.Store(rr, req)

		// Check the response status code and body
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Contains(t, rr.Body.String(), mockerror.Error())
	})
	t.Run("test store - Submit error", func(t *testing.T) {
		// Create a mock Gateway for testing purposes
		mockGateway := NewMockGatewayInt(ctrl)
		mockContract := NewMockContractInt(ctrl)
		mockNetwork := NewMockNetworkInt(ctrl)
		mockProposal := NewMockProposalInt(ctrl)
		mockTransaction := NewMockTransactionInt(ctrl)

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
		mockContract.EXPECT().NewProposal(gomock.Any(), gomock.Any()).Return(mockProposal, nil)
		mockProposal.EXPECT().Endorse().Return(mockTransaction, nil)
		mockerror := errors.New("Submit error")
		mockTransaction.EXPECT().Submit().Return(nil, mockerror)

		// Set up the HTTP request and response recorder
		req, err := http.NewRequest("POST", "/store?didDoc="+url.QueryEscape(didDoc), nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		setup.Store(rr, req)

		// Check the response status code and body
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Contains(t, rr.Body.String(), mockerror.Error())
	})

}
