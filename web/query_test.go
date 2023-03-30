package fabric

import (
	mock "fabric/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQuery(t *testing.T) {
	t.Run("test query DID success", func(t *testing.T) {
		// Create a new OrgSetup struct with the mock Gateway
		mockorgSetup := &mock.MockOrgSetup{}

		// Create a test request
		req, err := http.NewRequest("GET", "/query?didId=did:fabric:1234", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		// Create a response recorder
		rr := httptest.NewRecorder()

		// Call the Query function with the mock setup and test request
		mockorgSetup.Query(rr, req, mockorgSetup)
		expectedResponse := "doc"
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, rr.Body.String(), expectedResponse)
	})

	t.Run("TestQueryInvalidMethod", func(t *testing.T) {
		// Create a mock OrgSetup
		mockSetup := &OrgSetup{}

		// Create a test request with an unsupported DID method
		req, err := http.NewRequest("GET", "/query?didId=did:unsupported:1234", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)

		}

		// Create a response recorder
		rr := httptest.NewRecorder()

		// Call the Query function with the mock setup and test request
		mockSetup.Query(rr, req)

		// Check the response
		if rr.Code != http.StatusOK {
			t.Errorf("Unexpected status code: got %v, expected %v", rr.Code, http.StatusOK)
		}

		expectedResponse := "unsupported DID method: unsupported"
		if rr.Body.String() != expectedResponse {
			t.Errorf("Unexpected response body: got %v, expected %v", rr.Body.String(), expectedResponse)
		}
	})

	t.Run("TestQueryInvalidFormat", func(t *testing.T) {
		// Create a mock OrgSetup
		mockSetup := &mock.MockOrgSetup{}

		// Create a test request with an invalid DID format
		req, err := http.NewRequest("GET", "/query?didId=invalidDIDFormat", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		// Create a response recorder
		rr := httptest.NewRecorder()

		// Call the Query function with the mock setup and test request
		mockSetup.Query(rr, req)

		// Check the response
		if rr.Code != http.StatusOK {
			t.Errorf("Unexpected status code: got %v, expected %v", rr.Code, http.StatusOK)
		}

		expectedResponse := "invalid DID format: invalidDIDFormat"
		if rr.Body.String() != expectedResponse {
			t.Errorf("Unexpected response body: got %v, expected %v", rr.Body.String(), expectedResponse)
		}
	})

	t.Run("TestQueryEvaluateTransactionError", func(t *testing.T) {
		// Create a mock OrgSetup
		mockSetup := &mock.MockOrgSetup{}

		// Create a test request
		req, err := http.NewRequest("GET", "/query?didId=did:fabric:1234", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		// Create a response recorder
		rr := httptest.NewRecorder()

		// Call the Query function with the mock setup and test request
		mockSetup.Query(rr, req)

		// Check the response
		if rr.Code != http.StatusOK {
			t.Errorf("Unexpected status code: got %v, expected %v", rr.Code, http.StatusOK)
		}

		expectedResponse := "Error: <some expected error>"
		if rr.Body.String() != expectedResponse {
			t.Errorf("Unexpected response body: got %v, expected %v", rr.Body.String(), expectedResponse)
		}
	})
}
