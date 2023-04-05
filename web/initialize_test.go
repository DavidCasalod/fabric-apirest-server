package fabric

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestInitializeErrorCases(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	invalidCertPath := filepath.Join(tempDir, "invalid_cert.pem")

	testCases := []struct {
		name        string
		setup       OrgSetup
		expectedErr string
	}{
		{
			name: "invalid TLS certificate path",
			setup: OrgSetup{
				TLSCertPath: "non_existent_path",
			},
			expectedErr: "failed to read certificate file",
		},
		{
			name: "invalid certificate path",
			setup: OrgSetup{
				TLSCertPath: invalidCertPath,
				CertPath:    "non_existent_path",
			},
			expectedErr: "failed to read certificate file",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := Initialize(tc.setup)
			if err == nil {
				t.Fatal("Expected an error, but got none")
			}
			if err != nil && tc.expectedErr != "" && !strings.Contains(err.Error(), tc.expectedErr) {
				t.Fatalf("Expected error containing: %s, got: %s", tc.expectedErr, err.Error())
			}
		})
	}
}
func TestNewSignErrorCases(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	invalidKeyPath := filepath.Join(tempDir, "invalid_key")

	testCases := []struct {
		name        string
		setup       OrgSetup
		expectedErr string
	}{
		{
			name: "invalid private key path",
			setup: OrgSetup{
				KeyPath: invalidKeyPath,
			},
			expectedErr: "failed to read private key directory",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := tc.setup.newSign()
			if err == nil {
				t.Fatal("Expected an error, but got none")
			}
			if err != nil && tc.expectedErr != "" && !strings.Contains(err.Error(), tc.expectedErr) {
				t.Fatalf("Expected error containing: %s, got: %s", tc.expectedErr, err.Error())
			}
		})
	}
}
