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
 abi.new(eth.account[1], 0, 30, {from: eth.account[0], data: contract.code, gas: 3000000}, function(e, contract){
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

## Measurements
tdb...

--> 0x29609ba8f356d63640eacaa1ca89b95ef75c2a69b603439680b95d350e74024c
--> address: 0x18837b1ae794a2ddf1c8947e3b10a6ea88981c77 

0x6b9410d01e42368cbf24f899cd271e1ad400b2bf

var writer = abi.at("0x6b9410d01e42368cbf24f899cd271e1ad400b2bf");
writer.reportTemperature.sendTransaction([1,2], [1000,2000], {from:eth.accounts[1]})
--> 0x952bd27e06f5c3c2b4bb207700c427b58e9b028a1f462100b516f5f8d1f4d4d4

gas: 34893
tx.reportTemperature.sendTransaction([1,2,3], [1000,2000,3000], {from:eth.accounts[1]})
--> 0x39d7e4a8c8d1bf6adc47ced5dc686b1a0d7a8a6c6eec9ae5b86b8602ba3470d8

gas: 40954

writer.nrMeasurements()
writer.success()

per Temp gas ok: 6061
creation gas: 349579

writer.reportTemperature.sendTransaction([1,2,3,35], [1000,2000,3000,4000], {from:eth.accounts[1], gas:30000})
0xa043b91a656a969cd9a734189b934e958d5c7d0ce568b7f182351e807d9bb09e
gas: 52036 // 11082


writer.reportTemperature.sendTransaction([1,2,3,4,5,6,7,8,9,10], [1000,2000,3000,4000,5000,6000,7000,8000,9000,10000], {from:eth.accounts[1]})
gas: 83446
*/
