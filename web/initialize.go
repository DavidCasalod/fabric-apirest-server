
package fabric

import (
	"crypto/tls" // Añade esta línea
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"time"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Initialize the setup for the organization.
func Initialize(setup OrgSetup) (*OrgSetup, error) {
	log.Printf("Initializing connection for %s...\n", setup.OrgName)
	clientConnection, err := setup.newGrpcConnection()
	if err != nil {
		return nil, err
	}
	id, err := setup.newIdentity()
	if err != nil {
		return nil, err
	}
	sign, err := setup.newSign()
	if err != nil {
		return nil, err
	}

	gateway, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		return nil, err
	}
	setup.Gateway = gateway
	log.Println("Initialization complete")
	return &setup, nil
}

// newGrpcConnection creates a gRPC connection to the Gateway server.
func (setup OrgSetup) newGrpcConnection() (*grpc.ClientConn, error) {
	certificate, err := loadCertificate(setup.TLSCertPath)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // Ignora la verificación del nombre del servidor
		RootCAs:            certPool,
	}
	transportCredentials := credentials.NewTLS(tlsConfig)

	connection, err := grpc.Dial(setup.PeerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection: %w", err)
	}

	return connection, nil
}

func (setup OrgSetup) newIdentity() (*identity.X509Identity, error) {
	certificate, err := loadCertificate(setup.CertPath)
	if err != nil {
		return nil, err
	}

	id, err := identity.NewX509Identity(setup.MSPID, certificate)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (setup OrgSetup) newSign() (identity.Sign, error) {
	files, err := ioutil.ReadDir(setup.KeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key directory: %w", err)
	}
	privateKeyPEM, err := ioutil.ReadFile(path.Join(setup.KeyPath, files[0].Name()))

	if err != nil {
		return nil, fmt.Errorf("failed to read private key file: %w", err)
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		return nil, err
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		return nil, err
	}

	return sign, nil
}

func loadCertificate(filename string) (*x509.Certificate, error) {
	certificatePEM, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}
	return identity.CertificateFromPEM(certificatePEM)
}
