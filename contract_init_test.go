/*
 * Copyright 2016 Modum.io and the CSG Group at University of Zurich
 *
 * Licensed under the Apache License, Version 2.0 (the 'License'); you may not
 * use this file except in compliance with the License. You may obtain a copy of
 * the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an 'AS IS' BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations under
 * the License.
 */

//install abigen as described here: https://github.com/ethereum/go-ethereum/wiki/Native-DApps:-Go-bindings-to-Ethereum-contracts
//go:generate abigen --sol TemperatureMeasurementA2.sol --pkg main --out TemperatureMeasurementA2.go
//go:generate abigen --sol TemperatureMeasurementB.sol --pkg main --out TemperatureMeasurementB.go

package main

import (
	"log"
	"math/big"
	"strconv"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/core/types"
	"crypto/ecdsa"
	"golang.org/x/net/context"
)

var keyTempWriter *ecdsa.PrivateKey
var keyCreator *ecdsa.PrivateKey
var authTempWriter *bind.TransactOpts
var authCreator *bind.TransactOpts

type BlockSimulator struct {
	contractA *TemperatureMeasurementA2
	contractB *TemperatureMeasurementB
	sim       *backends.SimulatedBackend
}

func NewBlockSimulator() *BlockSimulator {
	bs := BlockSimulator{}
	keyTempWriter, _ = crypto.GenerateKey()
	authTempWriter = bind.NewKeyedTransactor(keyTempWriter)

	keyCreator, _ = crypto.GenerateKey()
	authCreator = bind.NewKeyedTransactor(keyCreator)
	bs.sim = backends.NewSimulatedBackend(core.GenesisAccount{Address: authCreator.From, Balance: big.NewInt(10000000000)})
	// Deploy a temp contract on the simulated blockchain
	_, _, tempA, err1 := DeployTemperatureMeasurementA2(authCreator, bs.sim, authTempWriter.From, int8(-2), int8(30), uint16(10), "SBT")
	if err1 != nil {
		log.Printf("Failed to deploy new temp contract: %v", err1)
	}

	_, _, tempB, err2 := DeployTemperatureMeasurementB(authCreator, bs.sim, authTempWriter.From, int8(-2), int8(30))
	if err2 != nil {
		log.Printf("Failed to deploy new temp contract: %v", err2)
	}

	bs.sim.Commit()
	bs.contractA = tempA;
	bs.contractB = tempB;

	return &bs
}

func (bs BlockSimulator) TransferETH() {
	ctx := context.Background()
	nonce, _ := bs.sim.PendingNonceAt(ctx, authCreator.From)
	tx := types.NewTransaction(nonce, authTempWriter.From,
		big.NewInt(100000000), big.NewInt(120000), big.NewInt(1), []byte{})
	signedTx, _ := authCreator.Signer(types.HomesteadSigner{},authCreator.From, tx)

	bs.sim.SendTransaction(ctx, signedTx)
	bs.sim.Commit()
}

func TestFreshContractNotSuccess(t *testing.T) {
	bs := NewBlockSimulator()
	retVal, _ := bs.contractA.Success(nil)
	if retVal {
		t.Log("fresh contract cannot be a success!")
		t.Fail()
	}
	retVal, _ = bs.contractB.Success(nil)
	if retVal {
		t.Log("fresh contract cannot be a success!")
		t.Fail()
	}
}

func TestFreshContractNotFailed(t *testing.T) {
	bs := NewBlockSimulator()
	retVal, _ := bs.contractA.Failed(nil)
	if retVal {
		t.Log("fresh contract cannot be failed!")
		t.Fail()
	}
	retVal, _ = bs.contractB.Failed(nil)
	if retVal {
		t.Log("fresh contract cannot be failed!")
		t.Fail()
	}
}

func TestFreshContractNrMeasurement(t *testing.T) {
	bs := NewBlockSimulator()
	retVal, _ := bs.contractA.NrMeasurements(nil)
	if retVal != 0 {
		t.Log("fresh contract cannot have measurements!")
		t.Fail()
	}
	retVal, _ = bs.contractB.NrMeasurements(nil)
	if retVal != 0 {
		t.Log("fresh contract cannot have measurements!")
		t.Fail()
	}
}

func TestFreshContractNrFailures(t *testing.T) {
	bs := NewBlockSimulator()
	retVal, _ := bs.contractA.NrFailures(nil)
	if retVal != 0 {
		t.Log("fresh contract cannot have failures!")
		t.Fail()
	}
	retVal, _ = bs.contractB.NrFailures(nil)
	if retVal != 0 {
		t.Log("fresh contract cannot have failures!")
		t.Fail()
	}
}

func TestFreshContractNrTimestapms(t *testing.T) {
	bs := NewBlockSimulator()
	retVal1, _ := bs.contractA.FailedTimestampLength(nil)
	if retVal1 != 0 {
		t.Log("fresh contract cannot have failures!")
		t.Fail()
	}
	retVal2, _ := bs.contractB.FailedTimestampSeconds(nil)
	if retVal2 != 0 {
		t.Log("fresh contract cannot have failures!")
		t.Fail()
	}
}

func TestFreshContractMinTemp(t *testing.T) {
	bs := NewBlockSimulator()
	retVal1, retVal2, _ := bs.contractA.TemperatureRange(nil)
	i1 := strconv.Itoa(int(retVal1))
	i2 := strconv.Itoa(int(retVal2))
	if retVal1 != -2 {
		t.Log("contract min temp must be -2, but is " + i1)
		t.Fail()
	}
	if retVal2 != 30 {
		t.Log("contract min temp must be -2, but is " + i2)
		t.Fail()
	}
}



func TestFreshContractTimestapmFirst(t *testing.T) {
	bs := NewBlockSimulator()
	retVal, _ := bs.contractA.TimestampFirst(nil)
	if retVal != 0 {
		t.Log("fresh contract need to have 0 here!")
		t.Fail()
	}
}

func TestFreshContractTimestapmLast(t *testing.T) {
	bs := NewBlockSimulator()
	retVal, _ := bs.contractA.TimestampLast(nil)
	if retVal != 0 {
		t.Log("fresh contract need to have 0 here!")
		t.Fail()
	}
}

func TestShiftRight(t *testing.T) {
	bs := NewBlockSimulator()
	retVal1, _ := bs.contractA.Shr(nil, 2, uint8(1))
	if retVal1 != 1 {
		t.Log("2 >> 1 = 1")
		t.Fail()
	}
	retVal2, _ := bs.contractA.Shr(nil, 8, uint8(3))
	if retVal2 != 1 {
		t.Log("8 >> 3 = 1")
		t.Fail()
	}
	retVal3, _ := bs.contractA.Shr(nil, 512, uint8(8))
	if retVal3 != 2 {
		t.Log("512 >> 8 = 2")
		t.Fail()
	}
	retVal4, _ := bs.contractA.Shr(nil, 196608, uint8(16))
	if retVal4 != 3 {
		t.Log("196608 >> 16 = 3")
		t.Fail()
	}
	retVal5, _ := bs.contractA.Shr(nil, 83886080, uint8(24))
	if retVal5 != 5 {
		t.Log("83886080 >> 24 = 5")
		t.Fail()
	}
}
