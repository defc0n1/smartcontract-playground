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
contract = eth.compile.solidity(sourceA).TemperatureMeasurementA
var abi = eth.contract(contract.info.abiDefinition)
abi.new(eth.accounts[1], 0, 30, {from: eth.accounts[0], data: contract.code, gas: 2000000}, function(e, contract){
    if(!e) {
      if(!contract.address) {
        console.log("Contract transaction send: https://testnet.etherscan.io/tx/" + contract.transactionHash + " - waiting to be mined...");
      } else {
        console.log("Contract mined! Address: " + contract.address);
        var tx = abi.at(contract.address);
      }
    }
});
EOF

echo "run: loadScript(\"tmp.js\")"
geth --testnet attach http://127.0.0.1:8545

