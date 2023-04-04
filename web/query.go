package fabric

import (
	"fmt"
	"net/http"
	"strings"
)

// func (setup OrgSetup) Query(w http.ResponseWriter, r *http.Request) {

// 	queryParams := r.URL.Query()
// 	didID := queryParams.Get("didId")
// 	// Split the DID id string to get the method
// 	parts := strings.Split(didID, ":")
// 	method := parts[1]
// 	if len(parts) != 3 || parts[0] != "did" || parts[1] != "fabric" {

// 		// Check if the DID method is Fabric
// 		if method != "fabric" {
// 			fmt.Fprintf(w, "unsupported DID method: %s", method)
// 		}
// 		fmt.Fprintf(w, "invalid DID format: %s", didID)
// 	}

// 	network := setup.Gatewaytest.GetNetwork(setup.ChannelId)
// 	contract := network.GetContract(setup.ChaincodeName)
// 	evaluateResponse, err := contract.EvaluateTransaction(setup.ChaincodeFunctions[0], didID)
// 	if err != nil {
// 		fmt.Fprintf(w, "Error: %s", err)
// 	}
// 	fmt.Fprintf(w, "%s", evaluateResponse)

// }

// 	doc, err := queryDID(didID, contract)
// 	if err != nil {
// 		panic(fmt.Errorf("failed to query DID: %w", err))
// 	}
// 	// Return the DID document resolution
// 	return &did.DocResolution{
// 		DIDDocument: doc,
// 		// DocumentMetadata: *did.DocumentMetadata,
// 	}, nil

// }

func (setup OrgSetup) Query(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()
	didID := queryParams.Get("didId")
	parts := strings.Split(didID, ":")

	// Verify that the DID ID is valid
	if len(parts) != 3 || parts[0] != "did" {
		http.Error(w, "Invalid DID format", http.StatusBadRequest)
		return
	}
	method := parts[1]
	// Verify that the DID method is Fabric
	if method != "fabric" {
		http.Error(w, "Unsupported DID method", http.StatusBadRequest)
		return
	}

	network := setup.Gatewaytest.GetNetwork(setup.ChannelId)
	contract := network.GetContract(setup.ChaincodeName)

	evaluateResponse, err := contract.EvaluateTransaction(setup.ChaincodeFunctions[0], didID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Couldn't evaluate transaction for didID '%s': %s", didID, err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", evaluateResponse)
}
