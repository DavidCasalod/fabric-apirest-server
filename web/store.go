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
		fmt.Fprintf(w, "failed to unmarshal DID document: %w", err)
	}
	didIDsplited := strings.Split(didDoc.ID, ":")
	method := didIDsplited[1]
	// Check if the DID method is Fabric
	if method != "fabric" {
		fmt.Fprintf(w, "unsupported DID method: %s", method)
	}

	network := setup.Gateway.GetNetwork(setup.ChannelId)
	contract := network.GetContract(setup.ChaincodeName)
	args := doc
	txn_proposal, err := contract.NewProposal(setup.ChaincodeFunctions[1], client.WithArguments(args...))
	if err != nil {
		fmt.Fprintf(w, "Error creating txn proposal: %s", err)
		return
	}
	txn_endorsed, err := txn_proposal.Endorse()
	if err != nil {
		fmt.Fprintf(w, "Error endorsing txn: %s", err)
		return
	}
	txn_committed, err := txn_endorsed.Submit()
	if err != nil {
		fmt.Fprintf(w, "Error submitting transaction: %s", err)
		return
	}
	fmt.Fprintf(w, "Transaction ID : %s Response: %s", txn_committed.TransactionID(), txn_endorsed.Result())
}
