#!/bin/bash

echo -n 'sourceA = "' > tmp.js
#tr '\n' ' ' < TemperatureMeasurementA.sol >> tmp.js
tr --delete '\n' < TemperatureMeasurementA.sol >> tmp.js
echo '";' >> tmp.js


echo -n 'sourceB = "' >> tmp.js
#tr '\n' ' ' < TemperatureMeasurementB.sol >> tmp.js
tr --delete '\n' < TemperatureMeasurementB.sol >> tmp.js
echo '";' >> tmp.js

cat << 'EOF' >> tmp.js
contractA = eth.compile.solidity(sourceA).TemperatureMeasurementA
var abiA = eth.contract(contractA.info.abiDefinition)
abiA.new(eth.accounts[1], 0, 30, 10, {from: eth.accounts[0], data: contractA.code, gas: 600000}, function(e, contract){
    if(!e) {
      if(!contract.address) {
        console.log("Contract A transaction send: https://testnet.etherscan.io/tx/" + contract.transactionHash + " - waiting to be mined...");
      } else {
        console.log("Contract A mined! Address: " + contract.address);
        var tx = abiA.at(contract.address);
      }
    }
});
contractB = eth.compile.solidity(sourceB).TemperatureMeasurementB
var abiB = eth.contract(contractB.info.abiDefinition)
abiB.new(eth.accounts[1], 0, 30, {from: eth.accounts[0], data: contractB.code, gas: 400000}, function(e, contract){
    if(!e) {
      if(!contract.address) {
        console.log("Contract B transaction send: https://testnet.etherscan.io/tx/" + contract.transactionHash + " - waiting to be mined...");
      } else {
        console.log("Contract B mined! Address: " + contract.address);
        var tx = abiB.at(contract.address);
      }
    }
});
EOF

# check gas: eth.getTransactionReceipt(txHash).gasUsed

# make sure you are running: geth --testnet --unlock 0,1 --password "" --rpc 
echo "run: loadScript(\"tmp.js\")"
geth --testnet attach http://127.0.0.1:8545

