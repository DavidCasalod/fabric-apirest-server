package fabric

import (
	"fmt"
	"net/http"
	"strings"
)

func (setup OrgSetup) Query(w http.ResponseWriter, r *http.Request) error {

	queryParams := r.URL.Query()
	didID := queryParams.Get("didId")
	// Split the DID id string to get the method
	parts := strings.Split(didID, ":")
	method := parts[1]
	if len(parts) != 3 || parts[0] != "did" || parts[1] != "fabric" {

		// Check if the DID method is Fabric
		if method != "fabric" {
			return fmt.Errorf("unsupported DID method: %s", method)
		}
		return fmt.Errorf("invalid DID format: %s", didID)
	}

	network := setup.Gateway.GetNetwork(setup.ChannelId)
	contract := network.GetContract(setup.ChaincodeName)
	evaluateResponse, err := contract.EvaluateTransaction(setup.ChaincodeFunctions[0], didID)
	if err != nil {
		return fmt.Errorf("Error: %s", err)
	}
	fmt.Fprintf(w, "Response: %s", evaluateResponse)
	return nil
}

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
