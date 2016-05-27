//install abigen as described here: https://github.com/ethereum/go-ethereum/wiki/Native-DApps:-Go-bindings-to-Ethereum-contracts
//go:generate abigen --sol TemperatureMeasurementA.sol --pkg main --out TemperatureMeasurementA.go
//go:generate abigen --sol TemperatureMeasurementB.sol --pkg main --out TemperatureMeasurementB.go

package main

import (
	"flag"
	"log"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

var tempContract *TemperatureMeasurementA
var temperatureWriterAddress common.Address

func TestMain(m *testing.M) {

	keyTempWriter, _ := crypto.GenerateKey()
	authTempWriter := bind.NewKeyedTransactor(keyTempWriter)

	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	sim := backends.NewSimulatedBackend(core.GenesisAccount{Address: auth.From, Balance: big.NewInt(10000000000)})
	// Deploy a temp contract on the simulated blockchain
	temperatureWriterAddress = authTempWriter.From
	_, _, temp, err := DeployTemperatureMeasurementA(auth, sim, temperatureWriterAddress, int8(0), int8(30), uint16(10), "SBT")
	if err != nil {
		log.Printf("Failed to deploy new temp contract: %v", err)
	}

	sim.Commit()
	tempContract = temp
	flag.Parse()
	os.Exit(m.Run())
}

func TestFreshContractNotFailed(t *testing.T) {
	name1, _ := tempContract.Failed(nil)
	if name1 {
		t.Log("fresh contract cannot be failed!")
		t.Fail()
	}
}
