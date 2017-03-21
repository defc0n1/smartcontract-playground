# Contracts for Measuring Temperature

There are two contracts to measure and record temperature, variant A and B. Variant B is 
cheaper in terms of gas used as it only records pass/fail, number of measurements, number
of failed measurements and when the first measurement failed.

Variant A is more expensive, stores more data, e.g. when the measuremnt started, when it stopped
and records the last n failures.

## Useful links

 * [Solidyt langague spec](http://solidity.readthedocs.io/en/latest/)
 * [Other good solidity page](https://docs.erisindustries.com/tutorials/solidity/solidity-1)
 * [Quick and dirty line removeal (is done by script)](http://www.textfixer.com/tools/remove-line-breaks.php)
 * [Online Solidity Editor - finds your bugs](http://chriseth.github.io/browser-solidity/#version=soljson-latest.js)
 * [geth and accounts](https://github.com/ethereum/go-ethereum/wiki/Managing-Your-Accounts)
 * [geth and netework](https://github.com/ethereum/go-ethereum/wiki/Connecting-to-the-network)
 * [Ethereum testnet browser](https://testnet.etherscan.io)

## Deployment

Currently deployment is done wyth the script prepare.sh which removes the lines breaks, creates an initial JavaScript file that 
creates the contract 

```javascript
 source = ...
 contract = eth.compile.solidity(source).TemperatureMeasurement -> ##A or B##
 var abi = eth.contract(contract.info.abiDefinition)
 abi.new(eth.account[1], 0, 30, 10, {from: eth.account[0], data: contract.code, gas: 600000}, function(e, contract){
    if(!e) {
      if(!contract.address) {
        console.log("Contract transaction send: https://testnet.etherscan.io/tx/" + contract.transactionHash + " waiting to be mined...");
      } else {
        console.log("Contract mined! Address: " + contract.address);
        var tx = abi.at(contract.address);
      }

    }
});
```

contract cost A: 773639 gas
contract cost B: 391544 gas

## Measurements

Since both A and B contracts are deployed, used 

```javascript
var tx = abiA.at("0x...") // for contract A
var tx = abiB.at("0x...") // for contract B
```
###Send temperature with contract A, everything good:

tx.reportTemperature.sendTransaction([1,2], [1000,2000], {from:eth.accounts[1], gas: 300000})
-> cost: 59073 gas // now 87902 gas
tx.reportTemperature.sendTransaction([1,2,3], [1000,2000,3000], {from:eth.accounts[1], gas: 300000})
-> cost: 61203 gas
tx.reportTemperature.sendTransaction([1,2,3,4], [1000,2000,3000,4000], {from:eth.accounts[1], gas: 300000})
-> cost: 63333 gas
-> 1 more: 2130 gas

tx.reportTemperature.sendTransaction([1,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2], [1000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000]
, {from:eth.accounts[0], gas: 300000})
-> cost for 34 measurements: //135402

tx.reportTemperature.sendTransaction([1,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2], [1000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000]
, {from:eth.accounts[0], gas: 1400000})
-> cost for 200 measurements: 515622 (541291)
-> loop twice over all measurements 843815
-> loop twice over all temp 843815
-> merged loop: 529099
-> no hash shifting: 238758


tx.reportTemperature.sendTransaction(200, [1, 0, 1, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 1, 0, 1, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 1, 0, 1, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 1, 0, 1, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 1, 0, 1, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 1, 0, 1, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 1, 0, 1, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 1, 0, 1, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 1, 0, 1, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 1, 0, 1, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 1, 0, 1, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 1, 0, 1, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 1, 0, 1, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 1, 0, 1, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 1, 0, 1, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 1, 0, 1, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 1, 0, 1, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 1, 0, 1, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 1, 0, 1, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 1, 0, 1, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 0], {from:eth.accounts[0], gas: 1500000})
-> cost for 200 measurements: 592240

************************
new contract cost contract A2: 1110688
tx.generateReport([1,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2], [1000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000])
--> result [[], [], 0, 199, "0xa56dbfa51f222607e717ee682fcdadf0d5255b16d3421a70dfef93879f3f7802"]
tx.reportResult.sendTransaction([], [],0, 199, 1000, 2000, "0xa56dbfa51f222607e717ee682fcdadf0d5255b16d3421a70dfef93879f3f7802", {from:eth.accounts[0], gas: 300000})
--> cost: 97413
tx.verifyReport(0, [1,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,2], [1000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000,2000])

###Send temperature with contract A, temperature failures:

tx.reportTemperature.sendTransaction([1,31], [1000,2000], {from:eth.accounts[1], gas: 300000})
-> cost: 104683
tx.reportTemperature.sendTransaction([1,31,32], [1000,2000,3000], {from:eth.accounts[1], gas: 300000})
-> cost: 87361
tx.reportTemperature.sendTransaction([1,31,32,33], [1000,2000,3000,4000], {from:eth.accounts[1], gas: 300000})
-> cost: 100029
-> 1 more: 12668 gas

###Send temperature with contract B, everything good:

tx.reportTemperature.sendTransaction([1,2], [1000,2000], {from:eth.accounts[1], gas: 300000})
-> cost: 34812 gas
tx.reportTemperature.sendTransaction([1,2,3], [1000,2000,3000], {from:eth.accounts[1], gas: 300000})
-> cost: 35547 gas
tx.reportTemperature.sendTransaction([1,2,3,4], [1000,2000,3000,4000], {from:eth.accounts[1], gas: 300000})
-> cost: 36282 gas
-> 1 more: 735 gas

###Send temperature with contract B, temperature failures:

tx.reportTemperature.sendTransaction([1,31], [1000,2000], {from:eth.accounts[1], gas: 300000})
-> cost: 54976
tx.reportTemperature.sendTransaction([1,31,32], [1000,2000,3000], {from:eth.accounts[1], gas: 300000})
-> cost: 35475
tx.reportTemperature.sendTransaction([1,31,32,33], [1000,2000,3000,4000], {from:eth.accounts[1], gas: 300000})
-> cost: 36174
-> 1 more: 699 gas

## Signature Checking in the Smart Contract

This feature checks that the signature of measurements generated by a sensor matches the measurements and the sensor's public key to ensure that measurements stem from the correct sensor and have not been modified in transit.

The feature is implemented in TemperatureMeasurementA2.sol.

The feature can be tested using the signature_test.js script. Keys, hashes and signatures for the test script have been generated as described below.

The test script imports a compiled contract wrapped in Javascript as produced by the [solc_helper](https://github.com/rakeshbs/solidity_compiler_helper) (with some modifications).
Alternatively, one can always modify the lines

	compiled = TemperatureMeasurementA2Compiled.TemperatureMeasurementA2.bin;
	abi = TemperatureMeasurementA2Compiled.TemperatureMeasurementA2.abi;

to point to the compiled contract and Abi respectively.

To use the test script in the Geth Javascript console we can
    
    loadScript("signature_test.js")
    deployAll()
    mineBlocks(1)
    runAll()
    mineBlocks(1)
    verifyAll()

### Key
    openssl ecparam -name secp256k1 -genkey -outform DER -noout -out key.der
    openssl asn1parse -inform DER -in key.der -dump
    0:d=0  hl=2 l= 116 cons: SEQUENCE          
    2:d=1  hl=2 l=   1 prim: INTEGER           :01
    5:d=1  hl=2 l=  32 prim: OCTET STRING      
      0000 - da 84 13 23 62 3f 12 8d-34 78 40 03 46 d6 37 9e   ...#b?..4x@.F.7.
      0010 - b6 62 8d 44 72 2e 87 30-56 44 b8 ed b3 33 4b 82   .b.Dr..0VD...3K.
    39:d=1  hl=2 l=   7 cons: cont [ 0 ]        
    41:d=2  hl=2 l=   5 prim: OBJECT            :secp256k1
    48:d=1  hl=2 l=  68 cons: cont [ 1 ]        
    50:d=2  hl=2 l=  66 prim: BIT STRING        
      0000 - 00 04 6c 01 46 40 04 5a-b0 90 50 3c 08 95 cc 6b   ..l.F@.Z..P<...k
      0010 - 7d 37 43 b0 49 9f 45 d8-68 fd 23 52 65 3c 98 7f   }7C.I.E.h.#Re<..
      0020 - 19 8d 74 2e 6e 81 57 0b-6f b1 d0 34 28 0b ec 10   ..t.n.W.o..4(...
      0030 - 39 9f 93 2a d3 b8 42 ba-75 2b a8 9d b5 35 9b 7b   9..*..B.u+...5.{
      0040 - 94 5c                                             .\

    openssl ec -in key.der -text -inform der
    read EC key
    Private-Key: (256 bit)
    priv:
        00:da:84:13:23:62:3f:12:8d:34:78:40:03:46:d6:
        37:9e:b6:62:8d:44:72:2e:87:30:56:44:b8:ed:b3:
        33:4b:82
    pub: 
        04:6c:01:46:40:04:5a:b0:90:50:3c:08:95:cc:6b:
        7d:37:43:b0:49:9f:45:d8:68:fd:23:52:65:3c:98:
        7f:19:8d:74:2e:6e:81:57:0b:6f:b1:d0:34:28:0b:
        ec:10:39:9f:93:2a:d3:b8:42:ba:75:2b:a8:9d:b5:
        35:9b:7b:94:5c
    ASN1 OID: secp256k1
    writing EC key
    -----BEGIN EC PRIVATE KEY-----
    MHQCAQEEINqEEyNiPxKNNHhAA0bWN562Yo1Eci6HMFZEuO2zM0uCoAcGBSuBBAAK
    oUQDQgAEbAFGQARasJBQPAiVzGt9N0OwSZ9F2Gj9I1JlPJh/GY10Lm6BVwtvsdA0
    KAvsEDmfkyrTuEK6dSuonbU1m3uUXA==
    -----END EC PRIVATE KEY-----

Thus, the public key used in the script is: 6c014640045ab090503c0895cc6b7d3743b0499f45d868fd2352653c987f198d742e6e81570b6fb1d034280bec10399f932ad3b842ba752ba89db5359b7b945c

To check signatures in Solidity we have to supply the public key in the Ethereum address format.
We can convert an Openssl keypair to an Ethereum address as follows:

Start with the 64 byte public key without the 04 prefix.

Take the SHA3-256 hash of that.

    cat pub.txt | xxd -r -p | xxd

The last 20 bytes make up the Ethereum address.

Thus, the Ethereum address corresponding to above public key is: 0xc2299921d8180db7db35614099127aec38d2031b.

### Hash
In the script we use measurements 0x95, 0x96 (25.06°C, 25.29°C),
start timestamp 0x1122334455667700 and
measurement interval 0x0f.

Thus, the signed data = sha556(0x950611223344556677000f)

    echo "959611223344556677000f" -n | xxd -r -p | sha256sum

This results in hash: 9af0179587d4770c2ccdfb0e138d3133cc52567a86f7ad5a9ce1e7c6184c816e

### Signature
    echo "959611223344556677000f" -n | xxd -r -p | openssl dgst -sha256 -sign key.der -keyform der  > signature.der
    openssl asn1parse -inform DER -in signature.der -dump

results in a valid signature of above hash with above key:

r = F2388D82FDA0E12B596EAF28BE28DF7FBB479A6BBA9D087B5DAF17ED3C95EAB4  
s = 85538F944F2F40782041548DF8BFD0D9B66788EA5B606C15BE29904E9D6928D0

v is dropped from DER output format. Its value is either 27 or 28.
Thus, signature checking needs to try both values.