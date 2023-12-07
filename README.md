
## Hyperledger-Fabric steps:

Go to te directory where you have the test-network and run the following commands: 
### Start network with channel
```
clear
```
### Deploy smart contract 
```
sudo ./network.sh deployCC -ccn VDR-chaincode -ccp ./vdr-smart-contract/ -ccl javascript -ccep "OR('Org1MSP.peer','Org2MSP.peer')" -cccg ./vdr-smart-contract/collections.json
```

## API REST SERVER:

Configure the .env file with the variables needed to acces to the hyperledger-fabric network used. 
To start the server in development mode, run:
```
sudo go run main.go
```

In production mode use the docker-compose.

To test it you can use the following command to create a new did and store it into de H-fabric blockchain: 
```
 curl -X POST -H "Content-Type: application/json" -d '{"didDoc": {"@context": "https://www.w3.org/ns/did/v1", "id": "did:fabric:44455", "publicKey": [{"id": "did:fabric:123456789abcdefghi#keys-1", "type": "Ed25519VerificationKey2018", "controller": "did:fabric:123456789abcdefghi", "publicKeyBase58": "H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV"}]}}' https://vdrfabric.dev4.ari-bip.eu/store
```

And to read it from the fabric network you can use the following http call: 
```

curl -X GET "https://vdrfabric.dev4.ari-bip.eu/query?didId=did:fabric:44455"
```
To test it you can use the following command to UPDATE a new did and store it into de H-fabric blockchain: 
```
 curl -X POST -H "Content-Type: application/json" -d '{"didDoc": {"@context": "https://www.w3.org/ns/did/v1", "id": "did:fabric:44455", "publicKey": [{"id": "did:fabric:123456789abcdefghi#keys-1", "type": "Ed25519VerificationKey2018", "controller": "did:fabric:123456789abcdefghi", "publicKeyBase58": "H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV"}]}}' https://vdrfabric.dev4.ari-bip.eu/update
```

To check the certs:
```
openssl x509 -in peer0.org1.example.com-cert.pem -text -noout
```