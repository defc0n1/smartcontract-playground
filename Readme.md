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

contract cost A: 556878 gas
contract cost B: 391544 gas

## Measurements

###Send temperature with contract A, everything good:

tx.reportTemperature.sendTransaction([1,2], [1000,2000], {from:eth.accounts[1], gas: 300000})
-> cost: 40876 gas
tx.reportTemperature.sendTransaction([1,2,3], [1000,2000,3000], {from:eth.accounts[1], gas: 300000})
-> cost: 36437 gas
tx.reportTemperature.sendTransaction([1,2,3,4], [1000,2000,3000,4000], {from:eth.accounts[1], gas: 300000})
-> cost: 37172
-> 1 more: 735 gas

###Send temperature with contract A, temperature failures:

tx.reportTemperature.sendTransaction([1,31], [1000,2000], {from:eth.accounts[1], gas: 300000})
-> cost: 76122
tx.reportTemperature.sendTransaction([1,31,32], [1000,2000,3000], {from:eth.accounts[1], gas: 300000})
-> cost: 57397
tx.reportTemperature.sendTransaction([1,31,32,33], [1000,2000,3000,4000], {from:eth.accounts[1], gas: 300000})
-> cost: 68662
-> 1 more: 11265 gas

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
