package fabric

import (
	"fmt"
	"net/http"
	"strings"
)

func (setup OrgSetup) Query(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()
	didID := queryParams.Get("didId")
	fmt.Println("REQUEST:", didID)
	parts := strings.Split(didID, ":")

	// Verify that the DID ID is valid
	if len(parts) != 3 || parts[0] != "did" {
		http.Error(w, "Invalid DID format", http.StatusBadRequest)
		return
	}
	method := parts[1]
	// Verify that the DID method is Priv
	if method != "priv" {
		http.Error(w, "Unsupported DID method", http.StatusBadRequest)
		return
	}

	// network := setup.Gatewaytest.GetNetwork(setup.ChannelId)
	network := setup.Gateway.GetNetwork(setup.ChannelId)
	contract := network.GetContract(setup.ChaincodeName)

	// evaluateResponse, err := contract.EvaluateTransaction(setup.ChaincodeFunctions[0], didID)
	evaluateResponse, err := contract.EvaluateTransaction("readdid", didID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Couldn't evaluate transaction for didID '%s': %s", didID, err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", evaluateResponse)
}
