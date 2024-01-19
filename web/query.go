package fabric

import (
	"encoding/json"
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

	// Unmarshal the response into a map
	var responseMap map[string]interface{}
	err = json.Unmarshal(evaluateResponse, &responseMap)
	if err != nil {
		http.Error(w, "Error unmarshaling evaluate response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Extract the didDoc content
	didDocContent, ok := responseMap["didDoc"]
	if !ok {
		http.Error(w, "didDoc not found in response", http.StatusInternalServerError)
		return
	}

	// Marshal the didDoc content back to JSON
	didDocJSON, err := json.Marshal(didDocContent)
	if err != nil {
		http.Error(w, "Error marshaling didDoc content: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the didDoc JSON as the response
	w.Header().Set("Content-Type", "application/json")
	w.Write(didDocJSON)
}
