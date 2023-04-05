package fabric

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/hyperledger/aries-framework-go/pkg/doc/did"
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

func (setup OrgSetup) Store(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()
	doc := queryParams.Get("didDoc")
	// Unmarshal the DID document from the query response
	didDoc := &did.Doc{}
	err := json.Unmarshal([]byte(doc), didDoc)
	if err != nil {
		http.Error(w, "Error creating txn proposal: "+err.Error(), http.StatusBadRequest)
		return
	}
	didIDsplited := strings.Split(didDoc.ID, ":")
	method := didIDsplited[1]
	// Check if the DID method is Fabric
	if method != "fabric" {
		http.Error(w, "Unsupported DID method:"+method, http.StatusBadRequest)
		return
	}

	network := setup.Gatewaytest.GetNetwork(setup.ChannelId)
	contract := network.GetContract(setup.ChaincodeName)
	args := []string{doc}
	txn_proposal, err := contract.NewProposal(setup.ChaincodeFunctions[1], client.WithArguments(args...))
	if err != nil {
		http.Error(w, "Error creating txn proposal: "+err.Error(), http.StatusBadRequest)
		return
	}
	txn_endorsed, err := txn_proposal.Endorse()
	if err != nil {
		http.Error(w, "Error endorsing txn: "+err.Error(), http.StatusBadRequest)
		return
	}
	txn_committed, err := txn_endorsed.Submit()
	if err != nil {
		http.Error(w, "Error submitting transaction: "+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Transaction ID : %s Response: %s", txn_committed.TransactionID(), txn_endorsed.Result())
}
