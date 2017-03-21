loadScript("/home/andi/smartcontract/TemperatureMeasurementA2_compiled.js");

demo1 = new Demo();
demo2 = new Demo();
demo3 = new Demo();
demo4 = new Demo();
demo5 = new Demo();
demo6 = new Demo();
demo7 = new Demo();
demo8 = new Demo();

function deployAll() {
  demo1.deploy("0x90", "0xa0");
  demo2.deploy("0x99", "0xa0");
  demo3.deploy("0x90", "0xa0");
  demo4.deploy("0x10", "0x20");
  demo5.deploy("0x90", "0xa0");
  demo6.deploy("0x90", "0xa0");
  demo7.deploy("0x90", "0xa0");
  demo8.deploy("0x90", "0xa0");
}

function runAll() {
  var measurementCount = 2;
  var startTimestamp = 0x1122334455667700;
  var measurementInterval = "0x0f";

  // hash of measurements 0x95 0x96 with above timestamp and interval
  var hash_9596 = "0x9af0179587d4770c2ccdfb0e138d3133cc52567a86f7ad5a9ce1e7c6184c816e";
  // hash of measurements 0x94 0x96 with above timestamp and interval
  var hash_9496 = "0x3197915ad0c2475a6a1b4cf53d033e155fdb27f9f0c09b2bd4aba7268e0ad34e";

  // signature of above hash 9596 by user with addres 0xc2299...
  var v_c2299_9596 = 28;
  var r_c2299_9596 = "0xF2388D82FDA0E12B596EAF28BE28DF7FBB479A6BBA9D087B5DAF17ED3C95EAB4";
  var s_c2299_9596 = "0x85538F944F2F40782041548DF8BFD0D9B66788EA5B606C15BE29904E9D6928D0";

  // signature of above hash 9496 by user with addres 0xc2299...
  var v_c2299_9496 = 28;
  var r_c2299_9496 = "0x0EFA3D1B474DF25AF77D8127D4443CBD321B4C6D706D28EB66D0AB79426C26BB";
  var s_c2299_9496 = "0x7E46EF1F95A793A66E38E9B7BFF28E44435F3F5E5D6F92504090799142C25BF5";

  demo1.run(measurementCount, startTimestamp, measurementInterval, hash_9596, false, v_c2299_9596, r_c2299_9596, s_c2299_9596);
  demo2.run(measurementCount, startTimestamp, measurementInterval, hash_9596, false, v_c2299_9596, r_c2299_9596, s_c2299_9596);
  demo3.run(measurementCount, startTimestamp, measurementInterval, hash_9596, true, v_c2299_9596, r_c2299_9596, s_c2299_9596);
  demo4.run(measurementCount, startTimestamp, measurementInterval, hash_9596, true, v_c2299_9596, r_c2299_9596, s_c2299_9596);
  demo5.run(measurementCount, startTimestamp, measurementInterval, hash_9596, false, v_c2299_9596, r_c2299_9596, s_c2299_9596);
  demo6.run(measurementCount, startTimestamp, measurementInterval, hash_9596, false, v_c2299_9496, r_c2299_9496, s_c2299_9496);
  demo7.run(measurementCount, startTimestamp, measurementInterval, hash_9496, false, v_c2299_9496, r_c2299_9496, s_c2299_9496);
  demo8.run(measurementCount, startTimestamp, measurementInterval, hash_9496, false, v_c2299_9496, r_c2299_9496, s_c2299_9496);
}

function verifyAll() {
  console.log("demo1 succeeds because measurements match the hash, are inside specifications and the caller says they are");
  console.log("demo2 fails because measurements match the hash, but aren't inside specifications while the caller says they are");
  console.log("demo3 fails because measurements match the hash, are inside specifications, but the caller says they are not");
  console.log("demo4 succeeds because measurements match the hash, are outside specifications and the caller says they aren't");

  console.log("demo5 fails because measurements don't match the hash, despite being inside specifications and the caller also saying they are");
  console.log("demo6 fails: measurements match hash, inside specifications, caller says they are, but hash doesn't match signature");
  console.log("demo7 fails: signature matches hash, inside specifications, caller says they are, but measurements don't match hash");
  console.log("demo8 succeeds with 94 96 measurements");

  console.log("demo1 " + demo1.verify(new Array("0x95", "0x96")));
  console.log("demo2 " + demo2.verify(new Array("0x95", "0x96")));
  console.log("demo3 " + demo3.verify(new Array("0x95", "0x96")));
  console.log("demo4 " + demo4.verify(new Array("0x95", "0x96")));
  console.log("demo5 " + demo5.verify(new Array("0x94", "0x96")));
  console.log("demo6 " + demo6.verify(new Array("0x95", "0x96")));
  console.log("demo7 " + demo7.verify(new Array("0x95", "0x96")));
  console.log("demo8 " + demo8.verify(new Array("0x94", "0x96")));
}

/*function run() {
  deployAll();
  mineBlocks(1);
  runAll();
  mineBlocks(1);
  verifyAll();
}*/

function Demo() {
  var signerPkey = "0xc2299921d8180db7db35614099127aec38d2031b";

  var me = this;

  this.deploy = function(min, max) {
    primary = eth.accounts[0];
    personal.unlockAccount(primary, "")
    balance = eth.getBalance(primary);
    if (balance == 0) {
      mineBlocks(1);
    }

    me.gasUsedTotal = 0;

    // create Creator contract
    compiled = TemperatureMeasurementA2Compiled.TemperatureMeasurementA2.bin;
    abi = TemperatureMeasurementA2Compiled.TemperatureMeasurementA2.abi;
    contract = eth.contract(abi);
    contract.new(primary, min, max, "", signerPkey, {from: primary, data: compiled, gas: 1000000}, function(err, myContract) {
      if (!err) {
        if (myContract.address) {
          me.contractInstance = myContract;
          var gasUsed = eth.getTransactionReceipt(me.contractInstance.transactionHash).gasUsed;
          me.gasUsedTotal += gasUsed;

          console.log(gasUsed);
        }
      } else {
        console.log(err);
      }
    });
  }

  this.verify = function(measurements) {
    var gasUsed = eth.getTransactionReceipt(me.reportTxHash).gasUsed;
    me.gasUsedTotal += gasUsed;

    console.log(gasUsed);

    return me.contractInstance.verifyReport(measurements);
  }

  this.run = function(measurementCount, startTimestamp, measurementInterval, hash, failed, v, r, s) {
    me.reportTxHash = me.contractInstance.reportResult.sendTransaction(measurementCount, startTimestamp, measurementInterval, hash, failed, v, r, s, {from: primary, gas: 1000000});
  }
}

function mineBlocks(count) {
  miner.start(); 
  admin.sleepBlocks(count); 
  miner.stop();
}