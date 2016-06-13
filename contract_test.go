//install abigen as described here: https://github.com/ethereum/go-ethereum/wiki/Native-DApps:-Go-bindings-to-Ethereum-contracts
//go:generate abigen --sol TemperatureMeasurementA.sol --pkg main --out TemperatureMeasurementA.go
//go:generate abigen --sol TemperatureMeasurementB.sol --pkg main --out TemperatureMeasurementB.go

package main

import (
	"flag"
	"log"
	"math/big"
	"os"
	"strconv"
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
	_, _, temp, err := DeployTemperatureMeasurementA(auth, sim, temperatureWriterAddress, int8(-2), int8(30), uint16(10), "SBT")
	if err != nil {
		log.Printf("Failed to deploy new temp contract: %v", err)
	}

	sim.Commit()
	tempContract = temp
	flag.Parse()
	os.Exit(m.Run())
}

func TestFreshContractNotSuccess(t *testing.T) {
	name1, _ := tempContract.Success(nil)
	if name1 {
		t.Log("fresh contract cannot be a success!")
		t.Fail()
	}
}

func TestFreshContractNotFailed(t *testing.T) {
	name1, _ := tempContract.Failed(nil)
	if name1 {
		t.Log("fresh contract cannot be failed!")
		t.Fail()
	}
}

func TestFreshContractNrMeasurement(t *testing.T) {
	name1, _ := tempContract.NrMeasurements(nil)
	if name1 != 0 {
		t.Log("fresh contract cannot have measurements!")
		t.Fail()
	}
}

func TestFreshContractNrFailures(t *testing.T) {
	name1, _ := tempContract.NrFailures(nil)
	if name1 != 0 {
		t.Log("fresh contract cannot have failures!")
		t.Fail()
	}
}

func TestFreshContractNrTimestapms(t *testing.T) {
	name1, _ := tempContract.FailedTimestampLength(nil)
	if name1 != 0 {
		t.Log("fresh contract cannot have failures!")
		t.Fail()
	}
}

func TestFreshContractMinTemp(t *testing.T) {
	name1, _ := tempContract.TemperatureMin(nil)
	i := strconv.Itoa(int(name1))
	if name1 != -2 {
		t.Log("contract min temp must be -2, but is " + i)
		t.Fail()
	}
}

func TestFreshContractMaxTemp(t *testing.T) {
	name1, _ := tempContract.TemperatureMax(nil)
	i := strconv.Itoa(int(name1))
	if name1 != 30 {
		t.Log("contract min temp must be 30, but is " + i)
		t.Fail()
	}
}

func TestFreshContractTimestapmFirst(t *testing.T) {
	name1, _ := tempContract.TimestampFirst(nil)
	if name1 != 0 {
		t.Log("fresh contract need to have 0 here!")
		t.Fail()
	}
}

func TestFreshContractTimestapmLast(t *testing.T) {
	name1, _ := tempContract.TimestampLast(nil)
	if name1 != 0 {
		t.Log("fresh contract need to have 0 here!")
		t.Fail()
	}
}

func TestShiftRight(t *testing.T) {
	name1, _ := tempContract.Shr(nil, 2, uint8(1))
	if name1 != 1 {
		t.Log("2 >> 1 = 1")
		t.Fail()
	}
	name2, _ := tempContract.Shr(nil, 8, uint8(3))
	if name2 != 1 {
		t.Log("8 >> 3 = 1")
		t.Fail()
	}
	name3, _ := tempContract.Shr(nil, 512, uint8(8))
	if name3 != 2 {
		t.Log("512 >> 8 = 2")
		t.Fail()
	}
	name4, _ := tempContract.Shr(nil, 196608, uint8(16))
	if name4 != 3 {
		t.Log("196608 >> 16 = 3")
		t.Fail()
	}
	name5, _ := tempContract.Shr(nil, 83886080, uint8(24))
	if name5 != 5 {
		t.Log("83886080 >> 24 = 5")
		t.Fail()
	}
}
