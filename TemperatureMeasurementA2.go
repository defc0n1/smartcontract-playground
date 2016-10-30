// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package main

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TemperatureMeasurementA2ABI is the input ABI used to generate the binding from.
const TemperatureMeasurementA2ABI = `[{"constant":true,"inputs":[{"name":"index","type":"uint16"}],"name":"hashAt","outputs":[{"name":"","type":"bytes32"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"_temperatures","type":"int8[]"},{"name":"_timestamps","type":"uint32[]"}],"name":"generateReport","outputs":[{"name":"","type":"uint32[]"},{"name":"","type":"int8[]"},{"name":"","type":"uint32"},{"name":"","type":"uint32"},{"name":"","type":"bytes32"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"success","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"nrFailures","outputs":[{"name":"","type":"uint32"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"timestampLast","outputs":[{"name":"","type":"uint32"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"temperatureRange","outputs":[{"name":"","type":"int8"},{"name":"","type":"int8"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"failedTimestampLength","outputs":[{"name":"","type":"uint16"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"series","type":"uint16"},{"name":"_temperatures","type":"int8[]"},{"name":"_timestamps","type":"uint32[]"}],"name":"verifyReport","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"index","type":"uint16"}],"name":"failedTimestampSecondsAt","outputs":[{"name":"","type":"uint32"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"_failedTimestampSeconds","type":"uint32[]"},{"name":"_failedTemperatures","type":"int8[]"},{"name":"_failures","type":"uint32"},{"name":"_measurements","type":"uint32"},{"name":"_firstTimestamp","type":"uint32"},{"name":"_lastTimestamp","type":"uint32"},{"name":"_hash","type":"bytes32"}],"name":"reportResult","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"hashLength","outputs":[{"name":"","type":"uint16"}],"payable":false,"type":"function"},{"constant":false,"inputs":[],"name":"done","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"index","type":"uint16"}],"name":"failedTemperaturesAt","outputs":[{"name":"","type":"int8"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"failed","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"failedTemperaturesLength","outputs":[{"name":"","type":"uint16"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"input","type":"uint32"},{"name":"bits","type":"uint8"}],"name":"shr","outputs":[{"name":"","type":"uint32"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"timestampFirst","outputs":[{"name":"","type":"uint32"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"nrMeasurements","outputs":[{"name":"","type":"uint32"}],"payable":false,"type":"function"},{"inputs":[{"name":"_temperatureWriter","type":"address"},{"name":"_minTemperature","type":"int8"},{"name":"_maxTemperature","type":"int8"},{"name":"_maxFailureReports","type":"uint16"},{"name":"_storageLocation","type":"string"}],"type":"constructor"}]`

// TemperatureMeasurementA2Bin is the compiled bytecode used for deploying new contracts.
const TemperatureMeasurementA2Bin = `0x60606040819052600680546001608060020a031916905561106b3881900390819083398101604052805160805160a05160c05160e05193949293919290910160008054600160a060020a031990811633179091556001805490911686178155600380547f01000000000000000000000000000000000000000000000000000000000000008681028190046101000261ff0019828a029290920460ff19909316929092171617905561ffff831610156100b657610002565b6003805463ffff00001916620100008402179055805160028054600082905290917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace602061010060018516150260001901909316849004601f90810184900482019386019083901061014b57805160ff19168380011785555b5061017b9291505b808211156101905760008155600101610137565b8280016001018555821561012f579182015b8281111561012f57825182600050559160200191906001019061015d565b50505050505050610ed7806101946000396000f35b509056606060405236156100da5760e060020a6000350463018fbf9f81146100df578063067cf182146101135780630b93381b146101975780630f79e120146101c35780632038e2e6146101e157806322aab53a1461020857806323e1d8d014610237578063287208491461024757806332797f441461035157806356f66f4c1461039e5780638abe55931461044f578063ae8421e11461045f578063b24d8de2146104a3578063ba414fa6146104e3578063c6d0ffda1461050f578063e0c68db01461051f578063ea1df43914610552578063eba37aff14610575575b610002565b346100025761058b600435600060076000508261ffff1681548110156100025760009182526020909120015490505b919050565b34610002576040805160206004803580820135838102808601850190965280855261059d95929460249490939285019282918501908490808284375050604080518735808a013560208181028481018201909552818452989a996044999398509190910195509350839250850190849080828437509496506102ef95505050505050565b34610002576106476004546000901580156101bc5750600654600063ffffffff909116115b90506101de565b346100025761065b60065463ffffffff640100000000909104165b90565b346100025761065b60065463ffffffff6c01000000000000000000000000909104166101de565b346100025760035460408051600083810b810b8252610100909304830b90920b60208301528051918290030190f35b34610002576106746004546101de565b346100025760408051602480356004818101356020818102868101820190975281865261064796833596939560449501929182919085019084908082843750506040805196358089013560208181028a81018201909452818a52979998606498909750602492909201955093508392508501908490808284375094965050505050505060408051602081810183526000808352835191820190935282815282808080610a7b89895b60408051602081810183526000808352835180830185528181528451808401865282815285518085018752838152865194850190965282845287518751959692959394859485949293859390929091849063ffffffff8716146106a457610002565b346100025761065b600435600060046000508261ffff1681548110156100025790600052602060002090600891828204019190066004029054906101000a900463ffffffff16905061010e565b34610002576040805160206004803580820135838102808601850190965280855261068b95929460249490939285019282918501908490808284375050604080518735808a013560208181028481018201909552818452989a9960449993985091909101955093508392508501908490808284375094965050933593505060643591505060843560a43560c43560015460009081908190819033600160a060020a03908116911614610c3b57610002565b34610002576106746007546101de565b346100025761068b60005433600160a060020a0390811691161480610492575060015433600160a060020a039081169116145b15610ed55733600160a060020a0316ff5b346100025761068d600435600060056000508261ffff1681548110156100025760009182526020808320818304015491066101000a9004900b905061010e565b346100025761064760045460009081901180156101bc575050600654600063ffffffff909116116101de565b34610002576106746005546101de565b346100025761065b6004356024355b60008160ff1660020a63ffffffff168363ffffffff16811561000257049392505050565b346100025761065b60065463ffffffff68010000000000000000909104166101de565b346100025761065b60065463ffffffff166101de565b60408051918252519081900360200190f35b6040518080602001806020018663ffffffff1681526020018563ffffffff168152602001846000191681526020018381038352888181518152602001915080519060200190602002808383829060006004602084601f0104600302600f01f1509050018381038252878181518152602001915080519060200190602002808383829060006004602084601f0104600302600f01f15090500197505050505050505060405180910390f35b604080519115158252519081900360200190f35b6040805163ffffffff9092168252519081900360200190f35b6040805161ffff9092168252519081900360200190f35b005b6040805160009290920b8252519081900360200190f35b600354604051600096506201000090910461ffff16908059106106c45750595b9080825280602002602001820160405280156106db575b5060035460405191955062010000900461ffff16908059106106fa5750595b908082528060200260200182016040528015610711575b5092508560050263ffffffff1660405180591061072b5750595b908082528060200260200182016040528015610742575b509150600090505b8563ffffffff168161ffff1610156107cb578b8161ffff168151811015610002579060200190602002015160f860020a02828260050260000161ffff16815181101561000257906020010190600160f860020a031916908160001a90535061083e8c8261ffff1681518110156100025790602001906020020151600861052e565b83838688600286600060405160200152604051808280519060200190808383829060006004602084601f0104600302600f01f1509050019150506020604051808303816000866161da5a03f11561000257505060405151939e50919c509a50985096505050505050509295509295909350565b60f860020a02828260050260010161ffff16815181101561000257906020010190600160f860020a031916908160001a9053506108948c8261ffff1681518110156100025790602001906020020151601061052e565b60f860020a02828260050260020161ffff16815181101561000257906020010190600160f860020a031916908160001a9053506108ea8c8261ffff1681518110156100025790602001906020020151601861052e565b60f860020a02828260050260030161ffff16815181101561000257906020010190600160f860020a031916908160001a9053508c8161ffff168151811015610002579060200190602002015160f860020a02828260050260040161ffff16815181101561000257906020010190600160f860020a031916908160001a905350600360019054906101000a900460000b60000b8d8261ffff168151811015610002579060200190602002015160000b13806109cb57506003548d51600091820b90910b908e9061ffff841690811015610002579060200190602002015160000b125b15610a735760035460019095019461ffff620100009091041663ffffffff861611610a73578b8161ffff1681518110156100025790602001906020020151846001870363ffffffff1681518110156100025763ffffffff9092166020928302909101909101528c518d9061ffff8316908110156100025790602001906020020151836001870363ffffffff1681518110156100025760009290920b6020928302909101909101525b60010161074a565b955095509550955095506001600760005080549050038a61ffff161415610acc576004548651141580610ab15750600554855114155b80610abe57508551855114155b15610af45760009650610bc8565b6007805483919061ffff8d169081101561000257600091825260209091200154149650610bc8565b5060005b60045461ffff82161015610b70576004805461ffff8316908110156100025790600052602060002090600891828204019190066004029054906101000a900463ffffffff1663ffffffff16868261ffff168151811015610002576020908102909101015163ffffffff1614610bd55760009650610bc8565b6007805483919061ffff8d169081101561000257600091825260209091200154148015610baf575060065463ffffffff85811664010000000090920416145b8015610bc5575060065463ffffffff8481169116145b96505b5050505050509392505050565b6005805461ffff8316908110156100025790600052602060002090602091828204019190069054906101000a900460000b60000b858261ffff168151811015610002576020908102909101015160000b14610c335760009650610bc8565b600101610af8565b89518b5114610c4957610002565b50506004548951600354919350915062010000900461ffff16825b8263ffffffff1684820363ffffffff16108015610c8a57508161ffff168163ffffffff16105b8015610ca357508863ffffffff1684820363ffffffff16105b15610cf65760048054600181018083558281838015829011610dbe576007016008900481600701600890048360005260206000209182019101610dbe91905b80821115610e535760008155600101610ce2565b6006805464010000000063ffffffff19821663ffffffff9283168c011781810483168d0190910267ffffffff0000000019909116179182905568010000000000000000909104161515610d6657600680546bffffffff000000000000000019166801000000000000000089021790555b600680546fffffffff00000000000000000000000019166c01000000000000000000000000880217905560078054600181018083558281838015829011610eb657600083815260209020610eb6918101908301610ce2565b505050919090600052602060002090600891828204019190066004028d87850363ffffffff1681518110156100025790602001906020020151909190916101000a81548163ffffffff021916908302179055505060056000508054806001018281815481835581811511610e5757601f016020900481601f01602090048360005260206000209182019101610e579190610ce2565b5090565b505050919090600052602060002090602091828204019190068c87850363ffffffff1681518110156100025790602001906020020151909190916101000a81548160ff021916908360f860020a90810204021790555050600101610c64565b5050506000928352506020909120019490945550505050505050505050565b56`

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
	contract, err := bindTemperatureMeasurementA2(address, backend.(bind.ContractCaller), backend.(bind.ContractTransactor))
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
