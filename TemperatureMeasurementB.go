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

// TemperatureMeasurementBABI is the input ABI used to generate the binding from.
const TemperatureMeasurementBABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"success\",\"outputs\":[{\"name\":\"res\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nrFailures\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"temperatureMin\",\"outputs\":[{\"name\":\"\",\"type\":\"int8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"temperatureMax\",\"outputs\":[{\"name\":\"\",\"type\":\"int8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_temperatures\",\"type\":\"int8[]\"},{\"name\":\"_timestamps\",\"type\":\"uint32[]\"}],\"name\":\"reportTemperature\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"done\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"failed\",\"outputs\":[{\"name\":\"res\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nrMeasurements\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"failedTimestampSeconds\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_temperatureWriter\",\"type\":\"address\"},{\"name\":\"_minTemperature\",\"type\":\"int8\"},{\"name\":\"_maxTemperature\",\"type\":\"int8\"}],\"payable\":false,\"type\":\"constructor\"}]"

// TemperatureMeasurementBBin is the compiled bytecode used for deploying new contracts.
const TemperatureMeasurementBBin = `0x60606040526001805460b060020a60f060020a03191690556002805463ffffffff19169055341561002c57fe5b6040516060806106278339810160409081528151602083015191909201515b60008054600160a060020a03338116600160a060020a03199283161783556001805485850b60ff90811675010000000000000000000000000000000000000000000260a860020a60ff02199689900b909116740100000000000000000000000000000000000000000260a060020a60ff0219948a1692909516919091179290921692909217929092169190911790555b5050505b610539806100ee6000396000f3006060604052361561007d5763ffffffff60e060020a6000350416630b93381b811461007f5780630f79e120146100a3578063348b1b7d146100cc5780634befc326146100f55780634c5a82cb1461011e578063ae8421e1146101ab578063ba414fa6146101bd578063eba37aff146101e1578063f375bcaa1461020a575bfe5b341561008757fe5b61008f610233565b604080519115158252519081900360200190f35b34156100ab57fe5b6100b3610269565b6040805163ffffffff9092168252519081900360200190f35b34156100d457fe5b6100dc610276565b60408051600092830b90920b8252519081900360200190f35b34156100fd57fe5b6100dc610287565b60408051600092830b90920b8252519081900360200190f35b341561012657fe5b6101a9600480803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843750506040805187358901803560208181028481018201909552818452989a99890198929750908201955093508392508501908490808284375094965061029895505050505050565b005b34156101b357fe5b6101a961046e565b005b34156101c557fe5b61008f6104ad565b604080519115158252519081900360200190f35b34156101e957fe5b6100b36104e5565b6040805163ffffffff9092168252519081900360200190f35b341561021257fe5b6100b36104f9565b6040805163ffffffff9092168252519081900360200190f35b60015460009060b060020a900463ffffffff161580156102635750600154600060d060020a90910463ffffffff16115b90505b90565b60025463ffffffff165b90565b60015460a060020a900460000b5b90565b60015460a860020a900460000b5b90565b6001546000908190819081908190819033600160a060020a03908116911614156104625786518851146102ca57610000565b505060015460025463ffffffff60d060020a830481169650908116945060a860020a8204600090810b945060a060020a8304810b935060b060020a90920416905b87518163ffffffff1610156103ea5785806001019650508360000b888263ffffffff1681518110151561033a57fe5b9060200190602002015160000b138061037557508260000b888263ffffffff1681518110151561036657fe5b9060200190602002015160000b125b156103e15763ffffffff821615156103da57868163ffffffff1681518110151561039b57fe5b60209081029190910101516001805479ffffffff00000000000000000000000000000000000000000000191660b060020a63ffffffff84160217905591505b6001909401935b5b60010161030b565b60015463ffffffff87811660d060020a909204161461043a57600180547fffff00000000ffffffffffffffffffffffffffffffffffffffffffffffffffff1660d060020a63ffffffff8916021790555b60025463ffffffff868116911614610462576002805463ffffffff191663ffffffff87161790555b5b5b5050505050505050565b60005433600160a060020a0390811691161480610499575060015433600160a060020a039081169116145b156104aa5733600160a060020a0316ff5b5b565b60015460009060b060020a900463ffffffff1681901180156102635750600154600060d060020a90910463ffffffff16115b90505b90565b60015460d060020a900463ffffffff165b90565b60015460b060020a900463ffffffff165b905600a165627a7a72305820be5e8cba11aad2267ff3a55ca788ca2e4ba037bbacec7307fc35cedc340072690029`

// DeployTemperatureMeasurementB deploys a new Ethereum contract, binding an instance of TemperatureMeasurementB to it.
func DeployTemperatureMeasurementB(auth *bind.TransactOpts, backend bind.ContractBackend, _temperatureWriter common.Address, _minTemperature int8, _maxTemperature int8) (common.Address, *types.Transaction, *TemperatureMeasurementB, error) {
	parsed, err := abi.JSON(strings.NewReader(TemperatureMeasurementBABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TemperatureMeasurementBBin), backend, _temperatureWriter, _minTemperature, _maxTemperature)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TemperatureMeasurementB{TemperatureMeasurementBCaller: TemperatureMeasurementBCaller{contract: contract}, TemperatureMeasurementBTransactor: TemperatureMeasurementBTransactor{contract: contract}}, nil
}

// TemperatureMeasurementB is an auto generated Go binding around an Ethereum contract.
type TemperatureMeasurementB struct {
	TemperatureMeasurementBCaller     // Read-only binding to the contract
	TemperatureMeasurementBTransactor // Write-only binding to the contract
}

// TemperatureMeasurementBCaller is an auto generated read-only Go binding around an Ethereum contract.
type TemperatureMeasurementBCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TemperatureMeasurementBTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TemperatureMeasurementBTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TemperatureMeasurementBSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TemperatureMeasurementBSession struct {
	Contract     *TemperatureMeasurementB // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TemperatureMeasurementBCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TemperatureMeasurementBCallerSession struct {
	Contract *TemperatureMeasurementBCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// TemperatureMeasurementBTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TemperatureMeasurementBTransactorSession struct {
	Contract     *TemperatureMeasurementBTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// TemperatureMeasurementBRaw is an auto generated low-level Go binding around an Ethereum contract.
type TemperatureMeasurementBRaw struct {
	Contract *TemperatureMeasurementB // Generic contract binding to access the raw methods on
}

// TemperatureMeasurementBCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TemperatureMeasurementBCallerRaw struct {
	Contract *TemperatureMeasurementBCaller // Generic read-only contract binding to access the raw methods on
}

// TemperatureMeasurementBTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TemperatureMeasurementBTransactorRaw struct {
	Contract *TemperatureMeasurementBTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTemperatureMeasurementB creates a new instance of TemperatureMeasurementB, bound to a specific deployed contract.
func NewTemperatureMeasurementB(address common.Address, backend bind.ContractBackend) (*TemperatureMeasurementB, error) {
	contract, err := bindTemperatureMeasurementB(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TemperatureMeasurementB{TemperatureMeasurementBCaller: TemperatureMeasurementBCaller{contract: contract}, TemperatureMeasurementBTransactor: TemperatureMeasurementBTransactor{contract: contract}}, nil
}

// NewTemperatureMeasurementBCaller creates a new read-only instance of TemperatureMeasurementB, bound to a specific deployed contract.
func NewTemperatureMeasurementBCaller(address common.Address, caller bind.ContractCaller) (*TemperatureMeasurementBCaller, error) {
	contract, err := bindTemperatureMeasurementB(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TemperatureMeasurementBCaller{contract: contract}, nil
}

// NewTemperatureMeasurementBTransactor creates a new write-only instance of TemperatureMeasurementB, bound to a specific deployed contract.
func NewTemperatureMeasurementBTransactor(address common.Address, transactor bind.ContractTransactor) (*TemperatureMeasurementBTransactor, error) {
	contract, err := bindTemperatureMeasurementB(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TemperatureMeasurementBTransactor{contract: contract}, nil
}

// bindTemperatureMeasurementB binds a generic wrapper to an already deployed contract.
func bindTemperatureMeasurementB(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TemperatureMeasurementBABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TemperatureMeasurementB *TemperatureMeasurementBRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TemperatureMeasurementB.Contract.TemperatureMeasurementBCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TemperatureMeasurementB *TemperatureMeasurementBRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TemperatureMeasurementB.Contract.TemperatureMeasurementBTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TemperatureMeasurementB *TemperatureMeasurementBRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TemperatureMeasurementB.Contract.TemperatureMeasurementBTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TemperatureMeasurementB *TemperatureMeasurementBCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TemperatureMeasurementB.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TemperatureMeasurementB *TemperatureMeasurementBTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TemperatureMeasurementB.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TemperatureMeasurementB *TemperatureMeasurementBTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TemperatureMeasurementB.Contract.contract.Transact(opts, method, params...)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() constant returns(res bool)
func (_TemperatureMeasurementB *TemperatureMeasurementBCaller) Failed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TemperatureMeasurementB.contract.Call(opts, out, "failed")
	return *ret0, err
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() constant returns(res bool)
func (_TemperatureMeasurementB *TemperatureMeasurementBSession) Failed() (bool, error) {
	return _TemperatureMeasurementB.Contract.Failed(&_TemperatureMeasurementB.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() constant returns(res bool)
func (_TemperatureMeasurementB *TemperatureMeasurementBCallerSession) Failed() (bool, error) {
	return _TemperatureMeasurementB.Contract.Failed(&_TemperatureMeasurementB.CallOpts)
}

// FailedTimestampSeconds is a free data retrieval call binding the contract method 0xf375bcaa.
//
// Solidity: function failedTimestampSeconds() constant returns(uint32)
func (_TemperatureMeasurementB *TemperatureMeasurementBCaller) FailedTimestampSeconds(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _TemperatureMeasurementB.contract.Call(opts, out, "failedTimestampSeconds")
	return *ret0, err
}

// FailedTimestampSeconds is a free data retrieval call binding the contract method 0xf375bcaa.
//
// Solidity: function failedTimestampSeconds() constant returns(uint32)
func (_TemperatureMeasurementB *TemperatureMeasurementBSession) FailedTimestampSeconds() (uint32, error) {
	return _TemperatureMeasurementB.Contract.FailedTimestampSeconds(&_TemperatureMeasurementB.CallOpts)
}

// FailedTimestampSeconds is a free data retrieval call binding the contract method 0xf375bcaa.
//
// Solidity: function failedTimestampSeconds() constant returns(uint32)
func (_TemperatureMeasurementB *TemperatureMeasurementBCallerSession) FailedTimestampSeconds() (uint32, error) {
	return _TemperatureMeasurementB.Contract.FailedTimestampSeconds(&_TemperatureMeasurementB.CallOpts)
}

// NrFailures is a free data retrieval call binding the contract method 0x0f79e120.
//
// Solidity: function nrFailures() constant returns(uint32)
func (_TemperatureMeasurementB *TemperatureMeasurementBCaller) NrFailures(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _TemperatureMeasurementB.contract.Call(opts, out, "nrFailures")
	return *ret0, err
}

// NrFailures is a free data retrieval call binding the contract method 0x0f79e120.
//
// Solidity: function nrFailures() constant returns(uint32)
func (_TemperatureMeasurementB *TemperatureMeasurementBSession) NrFailures() (uint32, error) {
	return _TemperatureMeasurementB.Contract.NrFailures(&_TemperatureMeasurementB.CallOpts)
}

// NrFailures is a free data retrieval call binding the contract method 0x0f79e120.
//
// Solidity: function nrFailures() constant returns(uint32)
func (_TemperatureMeasurementB *TemperatureMeasurementBCallerSession) NrFailures() (uint32, error) {
	return _TemperatureMeasurementB.Contract.NrFailures(&_TemperatureMeasurementB.CallOpts)
}

// NrMeasurements is a free data retrieval call binding the contract method 0xeba37aff.
//
// Solidity: function nrMeasurements() constant returns(uint32)
func (_TemperatureMeasurementB *TemperatureMeasurementBCaller) NrMeasurements(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _TemperatureMeasurementB.contract.Call(opts, out, "nrMeasurements")
	return *ret0, err
}

// NrMeasurements is a free data retrieval call binding the contract method 0xeba37aff.
//
// Solidity: function nrMeasurements() constant returns(uint32)
func (_TemperatureMeasurementB *TemperatureMeasurementBSession) NrMeasurements() (uint32, error) {
	return _TemperatureMeasurementB.Contract.NrMeasurements(&_TemperatureMeasurementB.CallOpts)
}

// NrMeasurements is a free data retrieval call binding the contract method 0xeba37aff.
//
// Solidity: function nrMeasurements() constant returns(uint32)
func (_TemperatureMeasurementB *TemperatureMeasurementBCallerSession) NrMeasurements() (uint32, error) {
	return _TemperatureMeasurementB.Contract.NrMeasurements(&_TemperatureMeasurementB.CallOpts)
}

// Success is a free data retrieval call binding the contract method 0x0b93381b.
//
// Solidity: function success() constant returns(res bool)
func (_TemperatureMeasurementB *TemperatureMeasurementBCaller) Success(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TemperatureMeasurementB.contract.Call(opts, out, "success")
	return *ret0, err
}

// Success is a free data retrieval call binding the contract method 0x0b93381b.
//
// Solidity: function success() constant returns(res bool)
func (_TemperatureMeasurementB *TemperatureMeasurementBSession) Success() (bool, error) {
	return _TemperatureMeasurementB.Contract.Success(&_TemperatureMeasurementB.CallOpts)
}

// Success is a free data retrieval call binding the contract method 0x0b93381b.
//
// Solidity: function success() constant returns(res bool)
func (_TemperatureMeasurementB *TemperatureMeasurementBCallerSession) Success() (bool, error) {
	return _TemperatureMeasurementB.Contract.Success(&_TemperatureMeasurementB.CallOpts)
}

// TemperatureMax is a free data retrieval call binding the contract method 0x4befc326.
//
// Solidity: function temperatureMax() constant returns(int8)
func (_TemperatureMeasurementB *TemperatureMeasurementBCaller) TemperatureMax(opts *bind.CallOpts) (int8, error) {
	var (
		ret0 = new(int8)
	)
	out := ret0
	err := _TemperatureMeasurementB.contract.Call(opts, out, "temperatureMax")
	return *ret0, err
}

// TemperatureMax is a free data retrieval call binding the contract method 0x4befc326.
//
// Solidity: function temperatureMax() constant returns(int8)
func (_TemperatureMeasurementB *TemperatureMeasurementBSession) TemperatureMax() (int8, error) {
	return _TemperatureMeasurementB.Contract.TemperatureMax(&_TemperatureMeasurementB.CallOpts)
}

// TemperatureMax is a free data retrieval call binding the contract method 0x4befc326.
//
// Solidity: function temperatureMax() constant returns(int8)
func (_TemperatureMeasurementB *TemperatureMeasurementBCallerSession) TemperatureMax() (int8, error) {
	return _TemperatureMeasurementB.Contract.TemperatureMax(&_TemperatureMeasurementB.CallOpts)
}

// TemperatureMin is a free data retrieval call binding the contract method 0x348b1b7d.
//
// Solidity: function temperatureMin() constant returns(int8)
func (_TemperatureMeasurementB *TemperatureMeasurementBCaller) TemperatureMin(opts *bind.CallOpts) (int8, error) {
	var (
		ret0 = new(int8)
	)
	out := ret0
	err := _TemperatureMeasurementB.contract.Call(opts, out, "temperatureMin")
	return *ret0, err
}

// TemperatureMin is a free data retrieval call binding the contract method 0x348b1b7d.
//
// Solidity: function temperatureMin() constant returns(int8)
func (_TemperatureMeasurementB *TemperatureMeasurementBSession) TemperatureMin() (int8, error) {
	return _TemperatureMeasurementB.Contract.TemperatureMin(&_TemperatureMeasurementB.CallOpts)
}

// TemperatureMin is a free data retrieval call binding the contract method 0x348b1b7d.
//
// Solidity: function temperatureMin() constant returns(int8)
func (_TemperatureMeasurementB *TemperatureMeasurementBCallerSession) TemperatureMin() (int8, error) {
	return _TemperatureMeasurementB.Contract.TemperatureMin(&_TemperatureMeasurementB.CallOpts)
}

// Done is a paid mutator transaction binding the contract method 0xae8421e1.
//
// Solidity: function done() returns()
func (_TemperatureMeasurementB *TemperatureMeasurementBTransactor) Done(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TemperatureMeasurementB.contract.Transact(opts, "done")
}

// Done is a paid mutator transaction binding the contract method 0xae8421e1.
//
// Solidity: function done() returns()
func (_TemperatureMeasurementB *TemperatureMeasurementBSession) Done() (*types.Transaction, error) {
	return _TemperatureMeasurementB.Contract.Done(&_TemperatureMeasurementB.TransactOpts)
}

// Done is a paid mutator transaction binding the contract method 0xae8421e1.
//
// Solidity: function done() returns()
func (_TemperatureMeasurementB *TemperatureMeasurementBTransactorSession) Done() (*types.Transaction, error) {
	return _TemperatureMeasurementB.Contract.Done(&_TemperatureMeasurementB.TransactOpts)
}

// ReportTemperature is a paid mutator transaction binding the contract method 0x4c5a82cb.
//
// Solidity: function reportTemperature(_temperatures int8[], _timestamps uint32[]) returns()
func (_TemperatureMeasurementB *TemperatureMeasurementBTransactor) ReportTemperature(opts *bind.TransactOpts, _temperatures []int8, _timestamps []uint32) (*types.Transaction, error) {
	return _TemperatureMeasurementB.contract.Transact(opts, "reportTemperature", _temperatures, _timestamps)
}

// ReportTemperature is a paid mutator transaction binding the contract method 0x4c5a82cb.
//
// Solidity: function reportTemperature(_temperatures int8[], _timestamps uint32[]) returns()
func (_TemperatureMeasurementB *TemperatureMeasurementBSession) ReportTemperature(_temperatures []int8, _timestamps []uint32) (*types.Transaction, error) {
	return _TemperatureMeasurementB.Contract.ReportTemperature(&_TemperatureMeasurementB.TransactOpts, _temperatures, _timestamps)
}

// ReportTemperature is a paid mutator transaction binding the contract method 0x4c5a82cb.
//
// Solidity: function reportTemperature(_temperatures int8[], _timestamps uint32[]) returns()
func (_TemperatureMeasurementB *TemperatureMeasurementBTransactorSession) ReportTemperature(_temperatures []int8, _timestamps []uint32) (*types.Transaction, error) {
	return _TemperatureMeasurementB.Contract.ReportTemperature(&_TemperatureMeasurementB.TransactOpts, _temperatures, _timestamps)
}
