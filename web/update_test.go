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

func TestUpdate(t *testing.T) {
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
			ChaincodeFunctions: []string{"queryDID", "updateDID"},
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
		req, err := http.NewRequest("POST", "/update?didDoc="+url.QueryEscape(didDoc), nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		setup.Update(rr, req)

		// Check the response status code and body
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Contains(t, rr.Body.String(), "txnId")
		assert.Contains(t, rr.Body.String(), "result ok")
	})
	t.Run("test update - Unsupported DID method", func(t *testing.T) {
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
			ChaincodeFunctions: []string{"queryDID", "updateDID"},
			Gatewaytest:        mockGateway,
		}

		// Set up the HTTP request and response recorder

		req, err := http.NewRequest("POST", "/update?didDoc="+url.QueryEscape(didDoc_unsupported), nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		setup.Update(rr, req)

		// Check the response status code and body
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Contains(t, rr.Body.String(), "Unsupported DID method:")
	})
	t.Run("test update - Invalid DID format", func(t *testing.T) {
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
			ChaincodeFunctions: []string{"queryDID", "updateDID"},
			Gatewaytest:        mockGateway,
		}

		// Set up the HTTP request and response recorder

		req, err := http.NewRequest("POST", "/update?didDoc=invalid", nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		setup.Update(rr, req)

		// Check the response status code and body
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Contains(t, rr.Body.String(), "Error creating txn proposal: invalid character 'i' ")
	})
	t.Run("test update - Proposal error", func(t *testing.T) {
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
			ChaincodeFunctions: []string{"queryDID", "updateDID"},
			Gatewaytest:        mockGateway,
		}
		// Set up the mock Gateway to return the mock Network and mock Contract
		mockGateway.EXPECT().GetNetwork(setup.ChannelId).Return(mockNetwork)
		mockNetwork.EXPECT().GetContract(setup.ChaincodeName).Return(mockContract)
		mockerror := errors.New("Proposal error")
		mockContract.EXPECT().NewProposal(gomock.Any(), gomock.Any()).Return(nil, mockerror)

		// Set up the HTTP request and response recorder
		req, err := http.NewRequest("POST", "/update?didDoc="+url.QueryEscape(didDoc), nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		setup.Update(rr, req)

		// Check the response status code and body
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Contains(t, rr.Body.String(), mockerror.Error())
	})
	t.Run("test update - Endorse error", func(t *testing.T) {

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
			ChaincodeFunctions: []string{"queryDID", "updateDID"},
			Gatewaytest:        mockGateway,
		}
		// Set up the mock Gateway to return the mock Network and mock Contract
		mockGateway.EXPECT().GetNetwork(setup.ChannelId).Return(mockNetwork)
		mockNetwork.EXPECT().GetContract(setup.ChaincodeName).Return(mockContract)
		mockContract.EXPECT().NewProposal(gomock.Any(), gomock.Any()).Return(mockProposal, nil)
		mockerror := errors.New("Endorsment error")
		mockProposal.EXPECT().Endorse().Return(nil, mockerror)

		// Set up the HTTP request and response recorder
		req, err := http.NewRequest("POST", "/update?didDoc="+url.QueryEscape(didDoc), nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		setup.Update(rr, req)

		// Check the response status code and body
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Contains(t, rr.Body.String(), mockerror.Error())
	})
	t.Run("test update - Submit error", func(t *testing.T) {
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
			ChaincodeFunctions: []string{"queryDID", "updateDID"},
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
		req, err := http.NewRequest("POST", "/update?didDoc="+url.QueryEscape(didDoc), nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		setup.Update(rr, req)

		// Check the response status code and body
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Contains(t, rr.Body.String(), mockerror.Error())
	})

}
