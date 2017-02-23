// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package smartcontract

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TemperatureMeasurementA2ABI is the input ABI used to generate the binding from.
const TemperatureMeasurementA2ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint16\"}],\"name\":\"hashAt\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_temperatures\",\"type\":\"int8[]\"},{\"name\":\"_timestamps\",\"type\":\"uint32[]\"}],\"name\":\"generateReport\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32[]\"},{\"name\":\"\",\"type\":\"int8[]\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"success\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nrFailures\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"timestampLast\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"temperatureRange\",\"outputs\":[{\"name\":\"\",\"type\":\"int8\"},{\"name\":\"\",\"type\":\"int8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"failedTimestampLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"series\",\"type\":\"uint16\"},{\"name\":\"_temperatures\",\"type\":\"int8[]\"},{\"name\":\"_timestamps\",\"type\":\"uint32[]\"}],\"name\":\"verifyReport\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint16\"}],\"name\":\"failedTimestampSecondsAt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_failedTimestampSeconds\",\"type\":\"uint32[]\"},{\"name\":\"_failedTemperatures\",\"type\":\"int8[]\"},{\"name\":\"_failures\",\"type\":\"uint32\"},{\"name\":\"_measurements\",\"type\":\"uint32\"},{\"name\":\"_firstTimestamp\",\"type\":\"uint32\"},{\"name\":\"_lastTimestamp\",\"type\":\"uint32\"},{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"reportResult\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hashLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"done\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint16\"}],\"name\":\"failedTemperaturesAt\",\"outputs\":[{\"name\":\"\",\"type\":\"int8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"failed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"failedTemperaturesLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"input\",\"type\":\"uint32\"},{\"name\":\"bits\",\"type\":\"uint8\"}],\"name\":\"shr\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"timestampFirst\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nrMeasurements\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_temperatureWriter\",\"type\":\"address\"},{\"name\":\"_minTemperature\",\"type\":\"int8\"},{\"name\":\"_maxTemperature\",\"type\":\"int8\"},{\"name\":\"_maxFailureReports\",\"type\":\"uint16\"},{\"name\":\"_storageLocation\",\"type\":\"string\"}],\"payable\":false,\"type\":\"constructor\"}]"

// TemperatureMeasurementA2Bin is the compiled bytecode used for deploying new contracts.
const TemperatureMeasurementA2Bin = `0x6060604052600680546001608060020a031916905534156200001d57fe5b6040516200138038038062001380833981016040908152815160208301519183015160608401516080850151929491929091015b60008054600160a060020a03338116600160a060020a031992831617835560018054918916919092161781556003805486840b60ff9081166101000261ff0019958a900b90911660ff1990921691909117939093169290921790915561ffff83161015620000bf5762000000565b6003805463ffff000019166201000061ffff8516021790558051620000ec906002906020840190620000f9565b505b5050505050620001a3565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200013c57805160ff19168380011785556200016c565b828001600101855582156200016c579182015b828111156200016c5782518255916020019190600101906200014f565b5b506200017b9291506200017f565b5090565b620001a091905b808211156200017b576000815560010162000186565b5090565b90565b6111cd80620001b36000396000f300606060405236156100e05763ffffffff60e060020a600035041663018fbf9f81146100e2578063067cf1821461010b5780630b93381b146102615780630f79e120146102855780632038e2e6146102ae57806322aab53a146102d757806323e1d8d014610310578063287208491461033757806332797f44146103d857806356f66f4c146104085780638abe5593146104bb578063ae8421e1146104e2578063b24d8de2146104f4578063ba414fa614610524578063c6d0ffda14610548578063e0c68db01461056f578063ea1df439146105a7578063eba37aff146105d0575bfe5b34156100ea57fe5b6100f961ffff600435166105f9565b60408051918252519081900360200190f35b341561011357fe5b610196600480803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843750506040805187358901803560208181028481018201909552818452989a99890198929750908201955093508392508501908490808284375094965061062695505050505050565b6040805163ffffffff8086169282019290925290831660608201526080810182905260a08082528651908201528551819060208083019160c08401918a810191028083838215610201575b80518252602083111561020157601f1990920191602091820191016101e1565b50505091909101848103835288518152885160209182019250818a0191028083838215610249575b80518252602083111561024957601f199092019160209182019101610229565b50505090500197505050505050505060405180910390f35b341561026957fe5b610271610a96565b604080519115158252519081900360200190f35b341561028d57fe5b610295610ab9565b6040805163ffffffff9092168252519081900360200190f35b34156102b657fe5b610295610ace565b6040805163ffffffff9092168252519081900360200190f35b34156102df57fe5b6102e7610aeb565b604051808360000b60000b81526020018260000b60000b81526020019250505060405180910390f35b341561031857fe5b610320610b00565b6040805161ffff9092168252519081900360200190f35b341561033f57fe5b6040805160248035600481810135602081810286810182019097528186526102719661ffff84351696939560449501929182919085019084908082843750506040805187358901803560208181028481018201909552818452989a998901989297509082019550935083925085019084908082843750949650610b0795505050505050565b604080519115158252519081900360200190f35b34156103e057fe5b61029561ffff60043516610cfe565b6040805163ffffffff9092168252519081900360200190f35b341561041057fe5b6104b9600480803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843750506040805187358901803560208181028481018201909552818452989a9989019892975090820195509350839250850190849080828437509496505063ffffffff853581169560208101358216955060408101358216945060608101359091169250608001359050610d42565b005b34156104c357fe5b610320610f8c565b6040805161ffff9092168252519081900360200190f35b34156104ea57fe5b6104b9610f93565b005b34156104fc57fe5b61050b61ffff60043516610fd2565b60408051600092830b90920b8252519081900360200190f35b341561052c57fe5b610271611010565b604080519115158252519081900360200190f35b341561055057fe5b610320611035565b6040805161ffff9092168252519081900360200190f35b341561057757fe5b61029563ffffffff6004351660ff6024351661103c565b6040805163ffffffff9092168252519081900360200190f35b34156105af57fe5b610295611065565b6040805163ffffffff9092168252519081900360200190f35b34156105d857fe5b61029561107e565b6040805163ffffffff9092168252519081900360200190f35b600060078261ffff1681548110151561060e57fe5b906000526020600020900160005b505490505b919050565b61062e61108b565b61063661108b565b6000600060006000600061064861108b565b61065061108b565b61065861108b565b60008c5195508b518663ffffffff1614151561067357610000565b600354604051600096506201000090910461ffff16908059106106935750595b908082528060200260200182016040525b5060035460405191955062010000900461ffff16908059106106c35750595b908082528060200260200182016040525b5092508560050263ffffffff166040518059106106ee5750595b908082528060200260200182016040525b509150600090505b8563ffffffff168161ffff1610156109f8578b8161ffff1681518110151561072b57fe5b9060200190602002015160f860020a02828260050260000161ffff1681518110151561075357fe5b906020010190600160f860020a031916908160001a9053506107918c8261ffff1681518110151561078057fe5b90602001906020020151600861103c565b60f860020a02828260050260010161ffff168151811015156107af57fe5b906020010190600160f860020a031916908160001a9053506107ed8c8261ffff168151811015156107dc57fe5b90602001906020020151601061103c565b60f860020a02828260050260020161ffff1681518110151561080b57fe5b906020010190600160f860020a031916908160001a9053506108498c8261ffff1681518110151561083857fe5b90602001906020020151601861103c565b60f860020a02828260050260030161ffff1681518110151561086757fe5b906020010190600160f860020a031916908160001a9053508c8161ffff1681518110151561089157fe5b9060200190602002015160f860020a02828260050260040161ffff168151811015156108b957fe5b906020010190600160f860020a031916908160001a905350600360019054906101000a900460000b60000b8d8261ffff168151811015156108f657fe5b9060200190602002015160000b12158061093857506003548d51600091820b90910b908e9061ffff841690811061092957fe5b9060200190602002015160000b125b156109ee5760035460019095019461ffff620100009091041663ffffffff8616116109ee578b8161ffff1681518110151561096f57fe5b90602001906020020151846001870363ffffffff1681518110151561099057fe5b63ffffffff9092166020928302909101909101528c518d9061ffff83169081106109b657fe5b90602001906020020151836001870363ffffffff168151811015156109d757fe5b9060200190602002019060000b908160000b815250505b5b5b600101610707565b838386886002866000604051602001526040518082805190602001908083835b60208310610a375780518252601f199092019160209182019101610a18565b51815160209384036101000a600019018019909216911617905260405191909301945091925050808303816000866161da5a03f11515610a7357fe5b505060405151939e50919c509a50985096505b5050505050509295509295909350565b600454600090158015610ab35750600654600063ffffffff909116115b90505b90565b600654640100000000900463ffffffff165b90565b6006546c01000000000000000000000000900463ffffffff165b90565b600354600081810b916101009004900b5b9091565b6004545b90565b6000610b1161108b565b610b1961108b565b6000600060006000610b2b8989610626565b955095509550955095506001600780549050038a61ffff161415610cc7576004548651141580610b5e5750600554855114155b80610b6b57508551855114155b15610b795760009650610cf0565b5060005b60045461ffff82161015610c66576004805461ffff8316908110610b9d57fe5b90600052602060002090600891828204019190066004025b9054906101000a900463ffffffff1663ffffffff16868261ffff16815181101515610bdc57fe5b6020908102909101015163ffffffff1614610bfa5760009650610cf0565b6005805461ffff8316908110610c0c57fe5b90600052602060002090602091828204019190065b9054906101000a900460000b60000b858261ffff16815181101515610c4257fe5b6020908102909101015160000b14610c5d5760009650610cf0565b5b600101610b7d565b6007805483919061ffff8d16908110610c7b57fe5b906000526020600020900160005b5054148015610caa575060065463ffffffff85811664010000000090920416145b8015610cc0575060065463ffffffff8481169116145b9650610cf0565b6007805483919061ffff8d16908110610cdc57fe5b906000526020600020900160005b50541496505b5b5050505050509392505050565b600060048261ffff16815481101515610d1357fe5b90600052602060002090600891828204019190066004025b9054906101000a900463ffffffff1690505b919050565b60015460009081908190819033600160a060020a03908116911614610d6657610000565b89518b5114610d7457610000565b50506004548951600354919350915062010000900461ffff16825b8263ffffffff1684820363ffffffff16108015610db557508161ffff168163ffffffff16105b8015610dce57508863ffffffff1684820363ffffffff16105b15610eae576004805460018101610de583826110c1565b91600052602060002090600891828204019190066004025b8d87850363ffffffff16815181101515610e1357fe5b90602001906020020151909190916101000a81548163ffffffff021916908363ffffffff1602179055505060058054806001018281610e5291906110fb565b91600052602060002090602091828204019190065b8c87850363ffffffff16815181101515610e7d57fe5b90602001906020020151909190916101000a81548160ff021916908360000b60ff160217905550505b600101610d8f565b6006805467ffffffff000000001963ffffffff19821663ffffffff9283168c018316179081166401000000009182900483168d018316909102179182905568010000000000000000909104161515610f2957600680546bffffffff000000000000000019166801000000000000000063ffffffff8a16021790555b600680546fffffffff00000000000000000000000019166c0100000000000000000000000063ffffffff8916021790556007805460018101610f6b8382611135565b916000526020600020900160005b50869055505b5050505050505050505050565b6007545b90565b60005433600160a060020a0390811691161480610fbe575060015433600160a060020a039081169116145b15610fcf5733600160a060020a0316ff5b5b565b600060058261ffff16815481101515610fe757fe5b90600052602060002090602091828204019190065b9054906101000a900460000b90505b919050565b6004546000908190118015610ab35750600654600063ffffffff909116115b90505b90565b6005545b90565b60008160ff1660020a63ffffffff168363ffffffff1681151561105b57fe5b0490505b92915050565b60065468010000000000000000900463ffffffff165b90565b60065463ffffffff165b90565b60408051602081019091526000815290565b60408051602081019091526000815290565b60408051602081019091526000815290565b8154818355818115116110f55760070160089004816007016008900483600052602060002091820191016110f5919061115f565b5b505050565b8154818355818115116110f557601f016020900481601f016020900483600052602060002091820191016110f5919061115f565b5b505050565b8154818355818115116110f5576000838152602090206110f591810190830161115f565b5b505050565b610ab691905b808211156111795760008155600101611165565b5090565b90565b610ab691905b808211156111795760008155600101611165565b5090565b905600a165627a7a72305820a3775762100f349292b065df33d898db81141acefadcef3759db0f5f094ead5a0029`

// DeployTemperatureMeasurementA2 deploys a new Ethereum contract, binding an instance of TemperatureMeasurementA2 to it.
func DeployTemperatureMeasurementA2(auth *bind.TransactOpts, backend bind.ContractBackend, _temperatureWriter common.Address, _minTemperature int8, _maxTemperature int8, _maxFailureReports uint16, _storageLocation string) (common.Address, *types.Transaction, *TemperatureMeasurementA2, error) {
	parsed, err := abi.JSON(strings.NewReader(TemperatureMeasurementA2ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TemperatureMeasurementA2Bin), backend, _temperatureWriter, _minTemperature, _maxTemperature, _maxFailureReports, _storageLocation)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TemperatureMeasurementA2{TemperatureMeasurementA2Caller: TemperatureMeasurementA2Caller{contract: contract}, TemperatureMeasurementA2Transactor: TemperatureMeasurementA2Transactor{contract: contract}}, nil
}

// TemperatureMeasurementA2 is an auto generated Go binding around an Ethereum contract.
type TemperatureMeasurementA2 struct {
	TemperatureMeasurementA2Caller     // Read-only binding to the contract
	TemperatureMeasurementA2Transactor // Write-only binding to the contract
}

// TemperatureMeasurementA2Caller is an auto generated read-only Go binding around an Ethereum contract.
type TemperatureMeasurementA2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TemperatureMeasurementA2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TemperatureMeasurementA2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TemperatureMeasurementA2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TemperatureMeasurementA2Session struct {
	Contract     *TemperatureMeasurementA2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TemperatureMeasurementA2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TemperatureMeasurementA2CallerSession struct {
	Contract *TemperatureMeasurementA2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// TemperatureMeasurementA2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TemperatureMeasurementA2TransactorSession struct {
	Contract     *TemperatureMeasurementA2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// TemperatureMeasurementA2Raw is an auto generated low-level Go binding around an Ethereum contract.
type TemperatureMeasurementA2Raw struct {
	Contract *TemperatureMeasurementA2 // Generic contract binding to access the raw methods on
}

// TemperatureMeasurementA2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TemperatureMeasurementA2CallerRaw struct {
	Contract *TemperatureMeasurementA2Caller // Generic read-only contract binding to access the raw methods on
}

// TemperatureMeasurementA2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TemperatureMeasurementA2TransactorRaw struct {
	Contract *TemperatureMeasurementA2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTemperatureMeasurementA2 creates a new instance of TemperatureMeasurementA2, bound to a specific deployed contract.
func NewTemperatureMeasurementA2(address common.Address, backend bind.ContractBackend) (*TemperatureMeasurementA2, error) {
	contract, err := bindTemperatureMeasurementA2(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TemperatureMeasurementA2{TemperatureMeasurementA2Caller: TemperatureMeasurementA2Caller{contract: contract}, TemperatureMeasurementA2Transactor: TemperatureMeasurementA2Transactor{contract: contract}}, nil
}

// NewTemperatureMeasurementA2Caller creates a new read-only instance of TemperatureMeasurementA2, bound to a specific deployed contract.
func NewTemperatureMeasurementA2Caller(address common.Address, caller bind.ContractCaller) (*TemperatureMeasurementA2Caller, error) {
	contract, err := bindTemperatureMeasurementA2(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TemperatureMeasurementA2Caller{contract: contract}, nil
}

// NewTemperatureMeasurementA2Transactor creates a new write-only instance of TemperatureMeasurementA2, bound to a specific deployed contract.
func NewTemperatureMeasurementA2Transactor(address common.Address, transactor bind.ContractTransactor) (*TemperatureMeasurementA2Transactor, error) {
	contract, err := bindTemperatureMeasurementA2(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TemperatureMeasurementA2Transactor{contract: contract}, nil
}

// bindTemperatureMeasurementA2 binds a generic wrapper to an already deployed contract.
func bindTemperatureMeasurementA2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TemperatureMeasurementA2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TemperatureMeasurementA2.Contract.TemperatureMeasurementA2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TemperatureMeasurementA2.Contract.TemperatureMeasurementA2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TemperatureMeasurementA2.Contract.TemperatureMeasurementA2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TemperatureMeasurementA2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TemperatureMeasurementA2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TemperatureMeasurementA2.Contract.contract.Transact(opts, method, params...)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() constant returns(bool)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Caller) Failed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TemperatureMeasurementA2.contract.Call(opts, out, "failed")
	return *ret0, err
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() constant returns(bool)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Session) Failed() (bool, error) {
	return _TemperatureMeasurementA2.Contract.Failed(&_TemperatureMeasurementA2.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() constant returns(bool)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2CallerSession) Failed() (bool, error) {
	return _TemperatureMeasurementA2.Contract.Failed(&_TemperatureMeasurementA2.CallOpts)
}

// FailedTemperaturesAt is a free data retrieval call binding the contract method 0xb24d8de2.
//
// Solidity: function failedTemperaturesAt(index uint16) constant returns(int8)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Caller) FailedTemperaturesAt(opts *bind.CallOpts, index uint16) (int8, error) {
	var (
		ret0 = new(int8)
	)
	out := ret0
	err := _TemperatureMeasurementA2.contract.Call(opts, out, "failedTemperaturesAt", index)
	return *ret0, err
}

// FailedTemperaturesAt is a free data retrieval call binding the contract method 0xb24d8de2.
//
// Solidity: function failedTemperaturesAt(index uint16) constant returns(int8)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Session) FailedTemperaturesAt(index uint16) (int8, error) {
	return _TemperatureMeasurementA2.Contract.FailedTemperaturesAt(&_TemperatureMeasurementA2.CallOpts, index)
}

// FailedTemperaturesAt is a free data retrieval call binding the contract method 0xb24d8de2.
//
// Solidity: function failedTemperaturesAt(index uint16) constant returns(int8)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2CallerSession) FailedTemperaturesAt(index uint16) (int8, error) {
	return _TemperatureMeasurementA2.Contract.FailedTemperaturesAt(&_TemperatureMeasurementA2.CallOpts, index)
}

// FailedTemperaturesLength is a free data retrieval call binding the contract method 0xc6d0ffda.
//
// Solidity: function failedTemperaturesLength() constant returns(uint16)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Caller) FailedTemperaturesLength(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _TemperatureMeasurementA2.contract.Call(opts, out, "failedTemperaturesLength")
	return *ret0, err
}

// FailedTemperaturesLength is a free data retrieval call binding the contract method 0xc6d0ffda.
//
// Solidity: function failedTemperaturesLength() constant returns(uint16)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Session) FailedTemperaturesLength() (uint16, error) {
	return _TemperatureMeasurementA2.Contract.FailedTemperaturesLength(&_TemperatureMeasurementA2.CallOpts)
}

// FailedTemperaturesLength is a free data retrieval call binding the contract method 0xc6d0ffda.
//
// Solidity: function failedTemperaturesLength() constant returns(uint16)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2CallerSession) FailedTemperaturesLength() (uint16, error) {
	return _TemperatureMeasurementA2.Contract.FailedTemperaturesLength(&_TemperatureMeasurementA2.CallOpts)
}

// FailedTimestampLength is a free data retrieval call binding the contract method 0x23e1d8d0.
//
// Solidity: function failedTimestampLength() constant returns(uint16)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Caller) FailedTimestampLength(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _TemperatureMeasurementA2.contract.Call(opts, out, "failedTimestampLength")
	return *ret0, err
}

// FailedTimestampLength is a free data retrieval call binding the contract method 0x23e1d8d0.
//
// Solidity: function failedTimestampLength() constant returns(uint16)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Session) FailedTimestampLength() (uint16, error) {
	return _TemperatureMeasurementA2.Contract.FailedTimestampLength(&_TemperatureMeasurementA2.CallOpts)
}

// FailedTimestampLength is a free data retrieval call binding the contract method 0x23e1d8d0.
//
// Solidity: function failedTimestampLength() constant returns(uint16)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2CallerSession) FailedTimestampLength() (uint16, error) {
	return _TemperatureMeasurementA2.Contract.FailedTimestampLength(&_TemperatureMeasurementA2.CallOpts)
}

// FailedTimestampSecondsAt is a free data retrieval call binding the contract method 0x32797f44.
//
// Solidity: function failedTimestampSecondsAt(index uint16) constant returns(uint32)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Caller) FailedTimestampSecondsAt(opts *bind.CallOpts, index uint16) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _TemperatureMeasurementA2.contract.Call(opts, out, "failedTimestampSecondsAt", index)
	return *ret0, err
}

// FailedTimestampSecondsAt is a free data retrieval call binding the contract method 0x32797f44.
//
// Solidity: function failedTimestampSecondsAt(index uint16) constant returns(uint32)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Session) FailedTimestampSecondsAt(index uint16) (uint32, error) {
	return _TemperatureMeasurementA2.Contract.FailedTimestampSecondsAt(&_TemperatureMeasurementA2.CallOpts, index)
}

// FailedTimestampSecondsAt is a free data retrieval call binding the contract method 0x32797f44.
//
// Solidity: function failedTimestampSecondsAt(index uint16) constant returns(uint32)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2CallerSession) FailedTimestampSecondsAt(index uint16) (uint32, error) {
	return _TemperatureMeasurementA2.Contract.FailedTimestampSecondsAt(&_TemperatureMeasurementA2.CallOpts, index)
}

// GenerateReport is a free data retrieval call binding the contract method 0x067cf182.
//
// Solidity: function generateReport(_temperatures int8[], _timestamps uint32[]) constant returns(uint32[], int8[], uint32, uint32, bytes32)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Caller) GenerateReport(opts *bind.CallOpts, _temperatures []int8, _timestamps []uint32) ([]uint32, []int8, uint32, uint32, [32]byte, error) {
	var (
		ret0 = new([]uint32)
		ret1 = new([]int8)
		ret2 = new(uint32)
		ret3 = new(uint32)
		ret4 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _TemperatureMeasurementA2.contract.Call(opts, out, "generateReport", _temperatures, _timestamps)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GenerateReport is a free data retrieval call binding the contract method 0x067cf182.
//
// Solidity: function generateReport(_temperatures int8[], _timestamps uint32[]) constant returns(uint32[], int8[], uint32, uint32, bytes32)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Session) GenerateReport(_temperatures []int8, _timestamps []uint32) ([]uint32, []int8, uint32, uint32, [32]byte, error) {
	return _TemperatureMeasurementA2.Contract.GenerateReport(&_TemperatureMeasurementA2.CallOpts, _temperatures, _timestamps)
}

// GenerateReport is a free data retrieval call binding the contract method 0x067cf182.
//
// Solidity: function generateReport(_temperatures int8[], _timestamps uint32[]) constant returns(uint32[], int8[], uint32, uint32, bytes32)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2CallerSession) GenerateReport(_temperatures []int8, _timestamps []uint32) ([]uint32, []int8, uint32, uint32, [32]byte, error) {
	return _TemperatureMeasurementA2.Contract.GenerateReport(&_TemperatureMeasurementA2.CallOpts, _temperatures, _timestamps)
}

// HashAt is a free data retrieval call binding the contract method 0x018fbf9f.
//
// Solidity: function hashAt(index uint16) constant returns(bytes32)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Caller) HashAt(opts *bind.CallOpts, index uint16) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _TemperatureMeasurementA2.contract.Call(opts, out, "hashAt", index)
	return *ret0, err
}

// HashAt is a free data retrieval call binding the contract method 0x018fbf9f.
//
// Solidity: function hashAt(index uint16) constant returns(bytes32)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Session) HashAt(index uint16) ([32]byte, error) {
	return _TemperatureMeasurementA2.Contract.HashAt(&_TemperatureMeasurementA2.CallOpts, index)
}

// HashAt is a free data retrieval call binding the contract method 0x018fbf9f.
//
// Solidity: function hashAt(index uint16) constant returns(bytes32)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2CallerSession) HashAt(index uint16) ([32]byte, error) {
	return _TemperatureMeasurementA2.Contract.HashAt(&_TemperatureMeasurementA2.CallOpts, index)
}

// HashLength is a free data retrieval call binding the contract method 0x8abe5593.
//
// Solidity: function hashLength() constant returns(uint16)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Caller) HashLength(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _TemperatureMeasurementA2.contract.Call(opts, out, "hashLength")
	return *ret0, err
}

// HashLength is a free data retrieval call binding the contract method 0x8abe5593.
//
// Solidity: function hashLength() constant returns(uint16)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Session) HashLength() (uint16, error) {
	return _TemperatureMeasurementA2.Contract.HashLength(&_TemperatureMeasurementA2.CallOpts)
}

// HashLength is a free data retrieval call binding the contract method 0x8abe5593.
//
// Solidity: function hashLength() constant returns(uint16)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2CallerSession) HashLength() (uint16, error) {
	return _TemperatureMeasurementA2.Contract.HashLength(&_TemperatureMeasurementA2.CallOpts)
}

// NrFailures is a free data retrieval call binding the contract method 0x0f79e120.
//
// Solidity: function nrFailures() constant returns(uint32)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Caller) NrFailures(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _TemperatureMeasurementA2.contract.Call(opts, out, "nrFailures")
	return *ret0, err
}

// NrFailures is a free data retrieval call binding the contract method 0x0f79e120.
//
// Solidity: function nrFailures() constant returns(uint32)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Session) NrFailures() (uint32, error) {
	return _TemperatureMeasurementA2.Contract.NrFailures(&_TemperatureMeasurementA2.CallOpts)
}

// NrFailures is a free data retrieval call binding the contract method 0x0f79e120.
//
// Solidity: function nrFailures() constant returns(uint32)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2CallerSession) NrFailures() (uint32, error) {
	return _TemperatureMeasurementA2.Contract.NrFailures(&_TemperatureMeasurementA2.CallOpts)
}

// NrMeasurements is a free data retrieval call binding the contract method 0xeba37aff.
//
// Solidity: function nrMeasurements() constant returns(uint32)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Caller) NrMeasurements(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _TemperatureMeasurementA2.contract.Call(opts, out, "nrMeasurements")
	return *ret0, err
}

// NrMeasurements is a free data retrieval call binding the contract method 0xeba37aff.
//
// Solidity: function nrMeasurements() constant returns(uint32)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Session) NrMeasurements() (uint32, error) {
	return _TemperatureMeasurementA2.Contract.NrMeasurements(&_TemperatureMeasurementA2.CallOpts)
}

// NrMeasurements is a free data retrieval call binding the contract method 0xeba37aff.
//
// Solidity: function nrMeasurements() constant returns(uint32)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2CallerSession) NrMeasurements() (uint32, error) {
	return _TemperatureMeasurementA2.Contract.NrMeasurements(&_TemperatureMeasurementA2.CallOpts)
}

// Shr is a free data retrieval call binding the contract method 0xe0c68db0.
//
// Solidity: function shr(input uint32, bits uint8) constant returns(uint32)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Caller) Shr(opts *bind.CallOpts, input uint32, bits uint8) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _TemperatureMeasurementA2.contract.Call(opts, out, "shr", input, bits)
	return *ret0, err
}

// Shr is a free data retrieval call binding the contract method 0xe0c68db0.
//
// Solidity: function shr(input uint32, bits uint8) constant returns(uint32)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Session) Shr(input uint32, bits uint8) (uint32, error) {
	return _TemperatureMeasurementA2.Contract.Shr(&_TemperatureMeasurementA2.CallOpts, input, bits)
}

// Shr is a free data retrieval call binding the contract method 0xe0c68db0.
//
// Solidity: function shr(input uint32, bits uint8) constant returns(uint32)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2CallerSession) Shr(input uint32, bits uint8) (uint32, error) {
	return _TemperatureMeasurementA2.Contract.Shr(&_TemperatureMeasurementA2.CallOpts, input, bits)
}

// Success is a free data retrieval call binding the contract method 0x0b93381b.
//
// Solidity: function success() constant returns(bool)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Caller) Success(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TemperatureMeasurementA2.contract.Call(opts, out, "success")
	return *ret0, err
}

// Success is a free data retrieval call binding the contract method 0x0b93381b.
//
// Solidity: function success() constant returns(bool)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Session) Success() (bool, error) {
	return _TemperatureMeasurementA2.Contract.Success(&_TemperatureMeasurementA2.CallOpts)
}

// Success is a free data retrieval call binding the contract method 0x0b93381b.
//
// Solidity: function success() constant returns(bool)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2CallerSession) Success() (bool, error) {
	return _TemperatureMeasurementA2.Contract.Success(&_TemperatureMeasurementA2.CallOpts)
}

// TemperatureRange is a free data retrieval call binding the contract method 0x22aab53a.
//
// Solidity: function temperatureRange() constant returns(int8, int8)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Caller) TemperatureRange(opts *bind.CallOpts) (int8, int8, error) {
	var (
		ret0 = new(int8)
		ret1 = new(int8)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _TemperatureMeasurementA2.contract.Call(opts, out, "temperatureRange")
	return *ret0, *ret1, err
}

// TemperatureRange is a free data retrieval call binding the contract method 0x22aab53a.
//
// Solidity: function temperatureRange() constant returns(int8, int8)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Session) TemperatureRange() (int8, int8, error) {
	return _TemperatureMeasurementA2.Contract.TemperatureRange(&_TemperatureMeasurementA2.CallOpts)
}

// TemperatureRange is a free data retrieval call binding the contract method 0x22aab53a.
//
// Solidity: function temperatureRange() constant returns(int8, int8)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2CallerSession) TemperatureRange() (int8, int8, error) {
	return _TemperatureMeasurementA2.Contract.TemperatureRange(&_TemperatureMeasurementA2.CallOpts)
}

// TimestampFirst is a free data retrieval call binding the contract method 0xea1df439.
//
// Solidity: function timestampFirst() constant returns(uint32)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Caller) TimestampFirst(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _TemperatureMeasurementA2.contract.Call(opts, out, "timestampFirst")
	return *ret0, err
}

// TimestampFirst is a free data retrieval call binding the contract method 0xea1df439.
//
// Solidity: function timestampFirst() constant returns(uint32)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Session) TimestampFirst() (uint32, error) {
	return _TemperatureMeasurementA2.Contract.TimestampFirst(&_TemperatureMeasurementA2.CallOpts)
}

// TimestampFirst is a free data retrieval call binding the contract method 0xea1df439.
//
// Solidity: function timestampFirst() constant returns(uint32)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2CallerSession) TimestampFirst() (uint32, error) {
	return _TemperatureMeasurementA2.Contract.TimestampFirst(&_TemperatureMeasurementA2.CallOpts)
}

// TimestampLast is a free data retrieval call binding the contract method 0x2038e2e6.
//
// Solidity: function timestampLast() constant returns(uint32)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Caller) TimestampLast(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _TemperatureMeasurementA2.contract.Call(opts, out, "timestampLast")
	return *ret0, err
}

// TimestampLast is a free data retrieval call binding the contract method 0x2038e2e6.
//
// Solidity: function timestampLast() constant returns(uint32)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Session) TimestampLast() (uint32, error) {
	return _TemperatureMeasurementA2.Contract.TimestampLast(&_TemperatureMeasurementA2.CallOpts)
}

// TimestampLast is a free data retrieval call binding the contract method 0x2038e2e6.
//
// Solidity: function timestampLast() constant returns(uint32)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2CallerSession) TimestampLast() (uint32, error) {
	return _TemperatureMeasurementA2.Contract.TimestampLast(&_TemperatureMeasurementA2.CallOpts)
}

// VerifyReport is a free data retrieval call binding the contract method 0x28720849.
//
// Solidity: function verifyReport(series uint16, _temperatures int8[], _timestamps uint32[]) constant returns(bool)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Caller) VerifyReport(opts *bind.CallOpts, series uint16, _temperatures []int8, _timestamps []uint32) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TemperatureMeasurementA2.contract.Call(opts, out, "verifyReport", series, _temperatures, _timestamps)
	return *ret0, err
}

// VerifyReport is a free data retrieval call binding the contract method 0x28720849.
//
// Solidity: function verifyReport(series uint16, _temperatures int8[], _timestamps uint32[]) constant returns(bool)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Session) VerifyReport(series uint16, _temperatures []int8, _timestamps []uint32) (bool, error) {
	return _TemperatureMeasurementA2.Contract.VerifyReport(&_TemperatureMeasurementA2.CallOpts, series, _temperatures, _timestamps)
}

// VerifyReport is a free data retrieval call binding the contract method 0x28720849.
//
// Solidity: function verifyReport(series uint16, _temperatures int8[], _timestamps uint32[]) constant returns(bool)
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2CallerSession) VerifyReport(series uint16, _temperatures []int8, _timestamps []uint32) (bool, error) {
	return _TemperatureMeasurementA2.Contract.VerifyReport(&_TemperatureMeasurementA2.CallOpts, series, _temperatures, _timestamps)
}

// Done is a paid mutator transaction binding the contract method 0xae8421e1.
//
// Solidity: function done() returns()
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Transactor) Done(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TemperatureMeasurementA2.contract.Transact(opts, "done")
}

// Done is a paid mutator transaction binding the contract method 0xae8421e1.
//
// Solidity: function done() returns()
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Session) Done() (*types.Transaction, error) {
	return _TemperatureMeasurementA2.Contract.Done(&_TemperatureMeasurementA2.TransactOpts)
}

// Done is a paid mutator transaction binding the contract method 0xae8421e1.
//
// Solidity: function done() returns()
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2TransactorSession) Done() (*types.Transaction, error) {
	return _TemperatureMeasurementA2.Contract.Done(&_TemperatureMeasurementA2.TransactOpts)
}

// ReportResult is a paid mutator transaction binding the contract method 0x56f66f4c.
//
// Solidity: function reportResult(_failedTimestampSeconds uint32[], _failedTemperatures int8[], _failures uint32, _measurements uint32, _firstTimestamp uint32, _lastTimestamp uint32, _hash bytes32) returns()
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Transactor) ReportResult(opts *bind.TransactOpts, _failedTimestampSeconds []uint32, _failedTemperatures []int8, _failures uint32, _measurements uint32, _firstTimestamp uint32, _lastTimestamp uint32, _hash [32]byte) (*types.Transaction, error) {
	return _TemperatureMeasurementA2.contract.Transact(opts, "reportResult", _failedTimestampSeconds, _failedTemperatures, _failures, _measurements, _firstTimestamp, _lastTimestamp, _hash)
}

// ReportResult is a paid mutator transaction binding the contract method 0x56f66f4c.
//
// Solidity: function reportResult(_failedTimestampSeconds uint32[], _failedTemperatures int8[], _failures uint32, _measurements uint32, _firstTimestamp uint32, _lastTimestamp uint32, _hash bytes32) returns()
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2Session) ReportResult(_failedTimestampSeconds []uint32, _failedTemperatures []int8, _failures uint32, _measurements uint32, _firstTimestamp uint32, _lastTimestamp uint32, _hash [32]byte) (*types.Transaction, error) {
	return _TemperatureMeasurementA2.Contract.ReportResult(&_TemperatureMeasurementA2.TransactOpts, _failedTimestampSeconds, _failedTemperatures, _failures, _measurements, _firstTimestamp, _lastTimestamp, _hash)
}

// ReportResult is a paid mutator transaction binding the contract method 0x56f66f4c.
//
// Solidity: function reportResult(_failedTimestampSeconds uint32[], _failedTemperatures int8[], _failures uint32, _measurements uint32, _firstTimestamp uint32, _lastTimestamp uint32, _hash bytes32) returns()
func (_TemperatureMeasurementA2 *TemperatureMeasurementA2TransactorSession) ReportResult(_failedTimestampSeconds []uint32, _failedTemperatures []int8, _failures uint32, _measurements uint32, _firstTimestamp uint32, _lastTimestamp uint32, _hash [32]byte) (*types.Transaction, error) {
	return _TemperatureMeasurementA2.Contract.ReportResult(&_TemperatureMeasurementA2.TransactOpts, _failedTimestampSeconds, _failedTemperatures, _failures, _measurements, _firstTimestamp, _lastTimestamp, _hash)
}
