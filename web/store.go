package fabric

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/hyperledger/aries-framework-go/pkg/doc/did"
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

func (setup OrgSetup) Store(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	var requestBody map[string]interface{}
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		http.Error(w, "Error unmarshaling request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	didDocJSON, ok := requestBody["didDoc"].(map[string]interface{})
	if !ok {
		http.Error(w, "Missing required parameter: didDoc", http.StatusBadRequest)
		return
	}
	doc, err := json.Marshal(didDocJSON)
	fmt.Println("doc:", doc)
	if err != nil {
		http.Error(w, "Error marshaling didDocJSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	// Unmarshal the DID document from the query response
	didDoc := &did.Doc{}
	er := json.Unmarshal([]byte(doc), didDoc)
	if er != nil {
		http.Error(w, "Error creating txn proposal: "+er.Error(), http.StatusBadRequest)
		return
	}

	didIDsplited := strings.Split(didDoc.ID, ":")
	method := didIDsplited[1]
	fmt.Println("ID:", didDoc.ID)
	// Check if the DID method is Fabric
	if method != "fabric" {
		http.Error(w, "Unsupported DID method:"+method, http.StatusBadRequest)
		return
	}

	// network := setup.Gatewaytest.GetNetwork(setup.ChannelId)
	network := setup.Gateway.GetNetwork(setup.ChannelId)
	contract := network.GetContract(setup.ChaincodeName)
	args := []string{didDoc.ID}
	transientDataMap := make(map[string][]byte)
	transientDataMap["didDoc"] = []byte(doc)

	// txn_proposal, err := contract.NewProposal(setup.ChaincodeFunctions[1], client.WithArguments(args...))
	// txn_proposal, err := contract.NewProposal("createdid", client.WithArguments(args...))
	txn_proposal, err := contract.NewProposal("createdid", client.WithArguments(args...), client.WithTransient(transientDataMap))

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
