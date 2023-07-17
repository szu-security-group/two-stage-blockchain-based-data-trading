// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mycontract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TradingMetaData contains all meta data concerning the Trading contract.
var TradingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"}],\"name\":\"buyerRefundEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"kc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"p\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"p1\",\"type\":\"uint256\"}],\"name\":\"challengeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"delta_1\",\"type\":\"bytes32\"}],\"name\":\"commitDataEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"complaintInfo\",\"type\":\"bytes32\"}],\"name\":\"initiationArbitrationEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"onchianArbitrationEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"payleftEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"masterkey\",\"type\":\"bytes32\"}],\"name\":\"releaseMasterkeyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_delta_2\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_length2\",\"type\":\"uint256\"}],\"name\":\"respondeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"verdictEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"complaintInfo\",\"type\":\"bytes32\"}],\"name\":\"RaiseOffChainArbitration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"index\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_committedSubKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleTreePath\",\"type\":\"bytes32[]\"}],\"name\":\"RaiseOnChainArbitration\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"arbitrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyerRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_kc\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_p\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_p1\",\"type\":\"uint256\"}],\"name\":\"challenge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_delta_1\",\"type\":\"bytes32\"}],\"name\":\"commitData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delta_1\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delta_2\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depth_2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"name\":\"hmac\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kc\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mk\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"opreationtime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"p\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"p1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payleft\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_mk\",\"type\":\"bytes32\"}],\"name\":\"releaseMasterkey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_delta_2\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_depth_2\",\"type\":\"uint256\"}],\"name\":\"response\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"name\":\"sha2562\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumtrading.State\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"test256\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"verdict\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TradingABI is the input ABI used to generate the binding from.
// Deprecated: Use TradingMetaData.ABI instead.
var TradingABI = TradingMetaData.ABI

// Trading is an auto generated Go binding around an Ethereum contract.
type Trading struct {
	TradingCaller     // Read-only binding to the contract
	TradingTransactor // Write-only binding to the contract
	TradingFilterer   // Log filterer for contract events
}

// TradingCaller is an auto generated read-only Go binding around an Ethereum contract.
type TradingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TradingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TradingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TradingSession struct {
	Contract     *Trading          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TradingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TradingCallerSession struct {
	Contract *TradingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TradingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TradingTransactorSession struct {
	Contract     *TradingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TradingRaw is an auto generated low-level Go binding around an Ethereum contract.
type TradingRaw struct {
	Contract *Trading // Generic contract binding to access the raw methods on
}

// TradingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TradingCallerRaw struct {
	Contract *TradingCaller // Generic read-only contract binding to access the raw methods on
}

// TradingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TradingTransactorRaw struct {
	Contract *TradingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrading creates a new instance of Trading, bound to a specific deployed contract.
func NewTrading(address common.Address, backend bind.ContractBackend) (*Trading, error) {
	contract, err := bindTrading(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Trading{TradingCaller: TradingCaller{contract: contract}, TradingTransactor: TradingTransactor{contract: contract}, TradingFilterer: TradingFilterer{contract: contract}}, nil
}

// NewTradingCaller creates a new read-only instance of Trading, bound to a specific deployed contract.
func NewTradingCaller(address common.Address, caller bind.ContractCaller) (*TradingCaller, error) {
	contract, err := bindTrading(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TradingCaller{contract: contract}, nil
}

// NewTradingTransactor creates a new write-only instance of Trading, bound to a specific deployed contract.
func NewTradingTransactor(address common.Address, transactor bind.ContractTransactor) (*TradingTransactor, error) {
	contract, err := bindTrading(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TradingTransactor{contract: contract}, nil
}

// NewTradingFilterer creates a new log filterer instance of Trading, bound to a specific deployed contract.
func NewTradingFilterer(address common.Address, filterer bind.ContractFilterer) (*TradingFilterer, error) {
	contract, err := bindTrading(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TradingFilterer{contract: contract}, nil
}

// bindTrading binds a generic wrapper to an already deployed contract.
func bindTrading(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TradingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trading *TradingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trading.Contract.TradingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trading *TradingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trading.Contract.TradingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trading *TradingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trading.Contract.TradingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trading *TradingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trading.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trading *TradingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trading.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trading *TradingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trading.Contract.contract.Transact(opts, method, params...)
}

// Arbitrator is a free data retrieval call binding the contract method 0x6cc6cde1.
//
// Solidity: function arbitrator() view returns(address)
func (_Trading *TradingCaller) Arbitrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "arbitrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Arbitrator is a free data retrieval call binding the contract method 0x6cc6cde1.
//
// Solidity: function arbitrator() view returns(address)
func (_Trading *TradingSession) Arbitrator() (common.Address, error) {
	return _Trading.Contract.Arbitrator(&_Trading.CallOpts)
}

// Arbitrator is a free data retrieval call binding the contract method 0x6cc6cde1.
//
// Solidity: function arbitrator() view returns(address)
func (_Trading *TradingCallerSession) Arbitrator() (common.Address, error) {
	return _Trading.Contract.Arbitrator(&_Trading.CallOpts)
}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() view returns(address)
func (_Trading *TradingCaller) Buyer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "buyer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() view returns(address)
func (_Trading *TradingSession) Buyer() (common.Address, error) {
	return _Trading.Contract.Buyer(&_Trading.CallOpts)
}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() view returns(address)
func (_Trading *TradingCallerSession) Buyer() (common.Address, error) {
	return _Trading.Contract.Buyer(&_Trading.CallOpts)
}

// DataId is a free data retrieval call binding the contract method 0xc8a663ec.
//
// Solidity: function dataId() view returns(uint256)
func (_Trading *TradingCaller) DataId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "dataId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DataId is a free data retrieval call binding the contract method 0xc8a663ec.
//
// Solidity: function dataId() view returns(uint256)
func (_Trading *TradingSession) DataId() (*big.Int, error) {
	return _Trading.Contract.DataId(&_Trading.CallOpts)
}

// DataId is a free data retrieval call binding the contract method 0xc8a663ec.
//
// Solidity: function dataId() view returns(uint256)
func (_Trading *TradingCallerSession) DataId() (*big.Int, error) {
	return _Trading.Contract.DataId(&_Trading.CallOpts)
}

// Delta1 is a free data retrieval call binding the contract method 0xa2fc30af.
//
// Solidity: function delta_1() view returns(bytes32)
func (_Trading *TradingCaller) Delta1(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "delta_1")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Delta1 is a free data retrieval call binding the contract method 0xa2fc30af.
//
// Solidity: function delta_1() view returns(bytes32)
func (_Trading *TradingSession) Delta1() ([32]byte, error) {
	return _Trading.Contract.Delta1(&_Trading.CallOpts)
}

// Delta1 is a free data retrieval call binding the contract method 0xa2fc30af.
//
// Solidity: function delta_1() view returns(bytes32)
func (_Trading *TradingCallerSession) Delta1() ([32]byte, error) {
	return _Trading.Contract.Delta1(&_Trading.CallOpts)
}

// Delta2 is a free data retrieval call binding the contract method 0x4bbf6642.
//
// Solidity: function delta_2() view returns(bytes32)
func (_Trading *TradingCaller) Delta2(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "delta_2")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Delta2 is a free data retrieval call binding the contract method 0x4bbf6642.
//
// Solidity: function delta_2() view returns(bytes32)
func (_Trading *TradingSession) Delta2() ([32]byte, error) {
	return _Trading.Contract.Delta2(&_Trading.CallOpts)
}

// Delta2 is a free data retrieval call binding the contract method 0x4bbf6642.
//
// Solidity: function delta_2() view returns(bytes32)
func (_Trading *TradingCallerSession) Delta2() ([32]byte, error) {
	return _Trading.Contract.Delta2(&_Trading.CallOpts)
}

// Depth2 is a free data retrieval call binding the contract method 0x80aa2dae.
//
// Solidity: function depth_2() view returns(uint256)
func (_Trading *TradingCaller) Depth2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "depth_2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Depth2 is a free data retrieval call binding the contract method 0x80aa2dae.
//
// Solidity: function depth_2() view returns(uint256)
func (_Trading *TradingSession) Depth2() (*big.Int, error) {
	return _Trading.Contract.Depth2(&_Trading.CallOpts)
}

// Depth2 is a free data retrieval call binding the contract method 0x80aa2dae.
//
// Solidity: function depth_2() view returns(uint256)
func (_Trading *TradingCallerSession) Depth2() (*big.Int, error) {
	return _Trading.Contract.Depth2(&_Trading.CallOpts)
}

// Hmac is a free data retrieval call binding the contract method 0x78192063.
//
// Solidity: function hmac(bytes32 key, bytes32 message) pure returns(bytes32)
func (_Trading *TradingCaller) Hmac(opts *bind.CallOpts, key [32]byte, message [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "hmac", key, message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hmac is a free data retrieval call binding the contract method 0x78192063.
//
// Solidity: function hmac(bytes32 key, bytes32 message) pure returns(bytes32)
func (_Trading *TradingSession) Hmac(key [32]byte, message [32]byte) ([32]byte, error) {
	return _Trading.Contract.Hmac(&_Trading.CallOpts, key, message)
}

// Hmac is a free data retrieval call binding the contract method 0x78192063.
//
// Solidity: function hmac(bytes32 key, bytes32 message) pure returns(bytes32)
func (_Trading *TradingCallerSession) Hmac(key [32]byte, message [32]byte) ([32]byte, error) {
	return _Trading.Contract.Hmac(&_Trading.CallOpts, key, message)
}

// Kc is a free data retrieval call binding the contract method 0xe6299631.
//
// Solidity: function kc() view returns(bytes32)
func (_Trading *TradingCaller) Kc(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "kc")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Kc is a free data retrieval call binding the contract method 0xe6299631.
//
// Solidity: function kc() view returns(bytes32)
func (_Trading *TradingSession) Kc() ([32]byte, error) {
	return _Trading.Contract.Kc(&_Trading.CallOpts)
}

// Kc is a free data retrieval call binding the contract method 0xe6299631.
//
// Solidity: function kc() view returns(bytes32)
func (_Trading *TradingCallerSession) Kc() ([32]byte, error) {
	return _Trading.Contract.Kc(&_Trading.CallOpts)
}

// Mk is a free data retrieval call binding the contract method 0x2f81b0a0.
//
// Solidity: function mk() view returns(bytes32)
func (_Trading *TradingCaller) Mk(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "mk")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Mk is a free data retrieval call binding the contract method 0x2f81b0a0.
//
// Solidity: function mk() view returns(bytes32)
func (_Trading *TradingSession) Mk() ([32]byte, error) {
	return _Trading.Contract.Mk(&_Trading.CallOpts)
}

// Mk is a free data retrieval call binding the contract method 0x2f81b0a0.
//
// Solidity: function mk() view returns(bytes32)
func (_Trading *TradingCallerSession) Mk() ([32]byte, error) {
	return _Trading.Contract.Mk(&_Trading.CallOpts)
}

// Opreationtime is a free data retrieval call binding the contract method 0x9ae67252.
//
// Solidity: function opreationtime() view returns(uint256)
func (_Trading *TradingCaller) Opreationtime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "opreationtime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Opreationtime is a free data retrieval call binding the contract method 0x9ae67252.
//
// Solidity: function opreationtime() view returns(uint256)
func (_Trading *TradingSession) Opreationtime() (*big.Int, error) {
	return _Trading.Contract.Opreationtime(&_Trading.CallOpts)
}

// Opreationtime is a free data retrieval call binding the contract method 0x9ae67252.
//
// Solidity: function opreationtime() view returns(uint256)
func (_Trading *TradingCallerSession) Opreationtime() (*big.Int, error) {
	return _Trading.Contract.Opreationtime(&_Trading.CallOpts)
}

// P is a free data retrieval call binding the contract method 0x9ae8886a.
//
// Solidity: function p() view returns(uint256)
func (_Trading *TradingCaller) P(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "p")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// P is a free data retrieval call binding the contract method 0x9ae8886a.
//
// Solidity: function p() view returns(uint256)
func (_Trading *TradingSession) P() (*big.Int, error) {
	return _Trading.Contract.P(&_Trading.CallOpts)
}

// P is a free data retrieval call binding the contract method 0x9ae8886a.
//
// Solidity: function p() view returns(uint256)
func (_Trading *TradingCallerSession) P() (*big.Int, error) {
	return _Trading.Contract.P(&_Trading.CallOpts)
}

// P1 is a free data retrieval call binding the contract method 0xc2a2747b.
//
// Solidity: function p1() view returns(uint256)
func (_Trading *TradingCaller) P1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "p1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// P1 is a free data retrieval call binding the contract method 0xc2a2747b.
//
// Solidity: function p1() view returns(uint256)
func (_Trading *TradingSession) P1() (*big.Int, error) {
	return _Trading.Contract.P1(&_Trading.CallOpts)
}

// P1 is a free data retrieval call binding the contract method 0xc2a2747b.
//
// Solidity: function p1() view returns(uint256)
func (_Trading *TradingCallerSession) P1() (*big.Int, error) {
	return _Trading.Contract.P1(&_Trading.CallOpts)
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_Trading *TradingCaller) Seller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "seller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_Trading *TradingSession) Seller() (common.Address, error) {
	return _Trading.Contract.Seller(&_Trading.CallOpts)
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_Trading *TradingCallerSession) Seller() (common.Address, error) {
	return _Trading.Contract.Seller(&_Trading.CallOpts)
}

// Sha2562 is a free data retrieval call binding the contract method 0xee17b807.
//
// Solidity: function sha2562(bytes32 key, bytes32 message) pure returns(bytes32)
func (_Trading *TradingCaller) Sha2562(opts *bind.CallOpts, key [32]byte, message [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "sha2562", key, message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Sha2562 is a free data retrieval call binding the contract method 0xee17b807.
//
// Solidity: function sha2562(bytes32 key, bytes32 message) pure returns(bytes32)
func (_Trading *TradingSession) Sha2562(key [32]byte, message [32]byte) ([32]byte, error) {
	return _Trading.Contract.Sha2562(&_Trading.CallOpts, key, message)
}

// Sha2562 is a free data retrieval call binding the contract method 0xee17b807.
//
// Solidity: function sha2562(bytes32 key, bytes32 message) pure returns(bytes32)
func (_Trading *TradingCallerSession) Sha2562(key [32]byte, message [32]byte) ([32]byte, error) {
	return _Trading.Contract.Sha2562(&_Trading.CallOpts, key, message)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Trading *TradingCaller) State(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "state")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Trading *TradingSession) State() (uint8, error) {
	return _Trading.Contract.State(&_Trading.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Trading *TradingCallerSession) State() (uint8, error) {
	return _Trading.Contract.State(&_Trading.CallOpts)
}

// Test256 is a free data retrieval call binding the contract method 0x7a1a456a.
//
// Solidity: function test256(bytes32 key) pure returns(bytes32)
func (_Trading *TradingCaller) Test256(opts *bind.CallOpts, key [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "test256", key)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Test256 is a free data retrieval call binding the contract method 0x7a1a456a.
//
// Solidity: function test256(bytes32 key) pure returns(bytes32)
func (_Trading *TradingSession) Test256(key [32]byte) ([32]byte, error) {
	return _Trading.Contract.Test256(&_Trading.CallOpts, key)
}

// Test256 is a free data retrieval call binding the contract method 0x7a1a456a.
//
// Solidity: function test256(bytes32 key) pure returns(bytes32)
func (_Trading *TradingCallerSession) Test256(key [32]byte) ([32]byte, error) {
	return _Trading.Contract.Test256(&_Trading.CallOpts, key)
}

// TimeLimit is a free data retrieval call binding the contract method 0xc08d1fe5.
//
// Solidity: function timeLimit() view returns(uint256)
func (_Trading *TradingCaller) TimeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "timeLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeLimit is a free data retrieval call binding the contract method 0xc08d1fe5.
//
// Solidity: function timeLimit() view returns(uint256)
func (_Trading *TradingSession) TimeLimit() (*big.Int, error) {
	return _Trading.Contract.TimeLimit(&_Trading.CallOpts)
}

// TimeLimit is a free data retrieval call binding the contract method 0xc08d1fe5.
//
// Solidity: function timeLimit() view returns(uint256)
func (_Trading *TradingCallerSession) TimeLimit() (*big.Int, error) {
	return _Trading.Contract.TimeLimit(&_Trading.CallOpts)
}

// RaiseOffChainArbitration is a paid mutator transaction binding the contract method 0xb838ae18.
//
// Solidity: function RaiseOffChainArbitration(bytes32 complaintInfo) returns()
func (_Trading *TradingTransactor) RaiseOffChainArbitration(opts *bind.TransactOpts, complaintInfo [32]byte) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "RaiseOffChainArbitration", complaintInfo)
}

// RaiseOffChainArbitration is a paid mutator transaction binding the contract method 0xb838ae18.
//
// Solidity: function RaiseOffChainArbitration(bytes32 complaintInfo) returns()
func (_Trading *TradingSession) RaiseOffChainArbitration(complaintInfo [32]byte) (*types.Transaction, error) {
	return _Trading.Contract.RaiseOffChainArbitration(&_Trading.TransactOpts, complaintInfo)
}

// RaiseOffChainArbitration is a paid mutator transaction binding the contract method 0xb838ae18.
//
// Solidity: function RaiseOffChainArbitration(bytes32 complaintInfo) returns()
func (_Trading *TradingTransactorSession) RaiseOffChainArbitration(complaintInfo [32]byte) (*types.Transaction, error) {
	return _Trading.Contract.RaiseOffChainArbitration(&_Trading.TransactOpts, complaintInfo)
}

// RaiseOnChainArbitration is a paid mutator transaction binding the contract method 0x07c5fbe5.
//
// Solidity: function RaiseOnChainArbitration(bytes32 index, bytes32 _committedSubKey, bytes32[] _merkleTreePath) returns(string)
func (_Trading *TradingTransactor) RaiseOnChainArbitration(opts *bind.TransactOpts, index [32]byte, _committedSubKey [32]byte, _merkleTreePath [][32]byte) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "RaiseOnChainArbitration", index, _committedSubKey, _merkleTreePath)
}

// RaiseOnChainArbitration is a paid mutator transaction binding the contract method 0x07c5fbe5.
//
// Solidity: function RaiseOnChainArbitration(bytes32 index, bytes32 _committedSubKey, bytes32[] _merkleTreePath) returns(string)
func (_Trading *TradingSession) RaiseOnChainArbitration(index [32]byte, _committedSubKey [32]byte, _merkleTreePath [][32]byte) (*types.Transaction, error) {
	return _Trading.Contract.RaiseOnChainArbitration(&_Trading.TransactOpts, index, _committedSubKey, _merkleTreePath)
}

// RaiseOnChainArbitration is a paid mutator transaction binding the contract method 0x07c5fbe5.
//
// Solidity: function RaiseOnChainArbitration(bytes32 index, bytes32 _committedSubKey, bytes32[] _merkleTreePath) returns(string)
func (_Trading *TradingTransactorSession) RaiseOnChainArbitration(index [32]byte, _committedSubKey [32]byte, _merkleTreePath [][32]byte) (*types.Transaction, error) {
	return _Trading.Contract.RaiseOnChainArbitration(&_Trading.TransactOpts, index, _committedSubKey, _merkleTreePath)
}

// BuyerRefund is a paid mutator transaction binding the contract method 0x3502e066.
//
// Solidity: function buyerRefund() returns()
func (_Trading *TradingTransactor) BuyerRefund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "buyerRefund")
}

// BuyerRefund is a paid mutator transaction binding the contract method 0x3502e066.
//
// Solidity: function buyerRefund() returns()
func (_Trading *TradingSession) BuyerRefund() (*types.Transaction, error) {
	return _Trading.Contract.BuyerRefund(&_Trading.TransactOpts)
}

// BuyerRefund is a paid mutator transaction binding the contract method 0x3502e066.
//
// Solidity: function buyerRefund() returns()
func (_Trading *TradingTransactorSession) BuyerRefund() (*types.Transaction, error) {
	return _Trading.Contract.BuyerRefund(&_Trading.TransactOpts)
}

// Challenge is a paid mutator transaction binding the contract method 0xb93a86ce.
//
// Solidity: function challenge(bytes32 _kc, uint256 _p, uint256 _p1) payable returns()
func (_Trading *TradingTransactor) Challenge(opts *bind.TransactOpts, _kc [32]byte, _p *big.Int, _p1 *big.Int) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "challenge", _kc, _p, _p1)
}

// Challenge is a paid mutator transaction binding the contract method 0xb93a86ce.
//
// Solidity: function challenge(bytes32 _kc, uint256 _p, uint256 _p1) payable returns()
func (_Trading *TradingSession) Challenge(_kc [32]byte, _p *big.Int, _p1 *big.Int) (*types.Transaction, error) {
	return _Trading.Contract.Challenge(&_Trading.TransactOpts, _kc, _p, _p1)
}

// Challenge is a paid mutator transaction binding the contract method 0xb93a86ce.
//
// Solidity: function challenge(bytes32 _kc, uint256 _p, uint256 _p1) payable returns()
func (_Trading *TradingTransactorSession) Challenge(_kc [32]byte, _p *big.Int, _p1 *big.Int) (*types.Transaction, error) {
	return _Trading.Contract.Challenge(&_Trading.TransactOpts, _kc, _p, _p1)
}

// CommitData is a paid mutator transaction binding the contract method 0x50339559.
//
// Solidity: function commitData(uint256 id, bytes32 _delta_1) returns()
func (_Trading *TradingTransactor) CommitData(opts *bind.TransactOpts, id *big.Int, _delta_1 [32]byte) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "commitData", id, _delta_1)
}

// CommitData is a paid mutator transaction binding the contract method 0x50339559.
//
// Solidity: function commitData(uint256 id, bytes32 _delta_1) returns()
func (_Trading *TradingSession) CommitData(id *big.Int, _delta_1 [32]byte) (*types.Transaction, error) {
	return _Trading.Contract.CommitData(&_Trading.TransactOpts, id, _delta_1)
}

// CommitData is a paid mutator transaction binding the contract method 0x50339559.
//
// Solidity: function commitData(uint256 id, bytes32 _delta_1) returns()
func (_Trading *TradingTransactorSession) CommitData(id *big.Int, _delta_1 [32]byte) (*types.Transaction, error) {
	return _Trading.Contract.CommitData(&_Trading.TransactOpts, id, _delta_1)
}

// Payleft is a paid mutator transaction binding the contract method 0xf4c1f07b.
//
// Solidity: function payleft() payable returns()
func (_Trading *TradingTransactor) Payleft(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "payleft")
}

// Payleft is a paid mutator transaction binding the contract method 0xf4c1f07b.
//
// Solidity: function payleft() payable returns()
func (_Trading *TradingSession) Payleft() (*types.Transaction, error) {
	return _Trading.Contract.Payleft(&_Trading.TransactOpts)
}

// Payleft is a paid mutator transaction binding the contract method 0xf4c1f07b.
//
// Solidity: function payleft() payable returns()
func (_Trading *TradingTransactorSession) Payleft() (*types.Transaction, error) {
	return _Trading.Contract.Payleft(&_Trading.TransactOpts)
}

// ReleaseMasterkey is a paid mutator transaction binding the contract method 0x706a0160.
//
// Solidity: function releaseMasterkey(bytes32 _mk) returns()
func (_Trading *TradingTransactor) ReleaseMasterkey(opts *bind.TransactOpts, _mk [32]byte) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "releaseMasterkey", _mk)
}

// ReleaseMasterkey is a paid mutator transaction binding the contract method 0x706a0160.
//
// Solidity: function releaseMasterkey(bytes32 _mk) returns()
func (_Trading *TradingSession) ReleaseMasterkey(_mk [32]byte) (*types.Transaction, error) {
	return _Trading.Contract.ReleaseMasterkey(&_Trading.TransactOpts, _mk)
}

// ReleaseMasterkey is a paid mutator transaction binding the contract method 0x706a0160.
//
// Solidity: function releaseMasterkey(bytes32 _mk) returns()
func (_Trading *TradingTransactorSession) ReleaseMasterkey(_mk [32]byte) (*types.Transaction, error) {
	return _Trading.Contract.ReleaseMasterkey(&_Trading.TransactOpts, _mk)
}

// Response is a paid mutator transaction binding the contract method 0x07e4312e.
//
// Solidity: function response(bytes32 _delta_2, uint256 _depth_2) payable returns()
func (_Trading *TradingTransactor) Response(opts *bind.TransactOpts, _delta_2 [32]byte, _depth_2 *big.Int) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "response", _delta_2, _depth_2)
}

// Response is a paid mutator transaction binding the contract method 0x07e4312e.
//
// Solidity: function response(bytes32 _delta_2, uint256 _depth_2) payable returns()
func (_Trading *TradingSession) Response(_delta_2 [32]byte, _depth_2 *big.Int) (*types.Transaction, error) {
	return _Trading.Contract.Response(&_Trading.TransactOpts, _delta_2, _depth_2)
}

// Response is a paid mutator transaction binding the contract method 0x07e4312e.
//
// Solidity: function response(bytes32 _delta_2, uint256 _depth_2) payable returns()
func (_Trading *TradingTransactorSession) Response(_delta_2 [32]byte, _depth_2 *big.Int) (*types.Transaction, error) {
	return _Trading.Contract.Response(&_Trading.TransactOpts, _delta_2, _depth_2)
}

// Verdict is a paid mutator transaction binding the contract method 0xf258d51e.
//
// Solidity: function verdict(address userAddr, address winner) returns()
func (_Trading *TradingTransactor) Verdict(opts *bind.TransactOpts, userAddr common.Address, winner common.Address) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "verdict", userAddr, winner)
}

// Verdict is a paid mutator transaction binding the contract method 0xf258d51e.
//
// Solidity: function verdict(address userAddr, address winner) returns()
func (_Trading *TradingSession) Verdict(userAddr common.Address, winner common.Address) (*types.Transaction, error) {
	return _Trading.Contract.Verdict(&_Trading.TransactOpts, userAddr, winner)
}

// Verdict is a paid mutator transaction binding the contract method 0xf258d51e.
//
// Solidity: function verdict(address userAddr, address winner) returns()
func (_Trading *TradingTransactorSession) Verdict(userAddr common.Address, winner common.Address) (*types.Transaction, error) {
	return _Trading.Contract.Verdict(&_Trading.TransactOpts, userAddr, winner)
}

// TradingBuyerRefundEventIterator is returned from FilterBuyerRefundEvent and is used to iterate over the raw logs and unpacked data for BuyerRefundEvent events raised by the Trading contract.
type TradingBuyerRefundEventIterator struct {
	Event *TradingBuyerRefundEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TradingBuyerRefundEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingBuyerRefundEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TradingBuyerRefundEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TradingBuyerRefundEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingBuyerRefundEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingBuyerRefundEvent represents a BuyerRefundEvent event raised by the Trading contract.
type TradingBuyerRefundEvent struct {
	UserAddr common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBuyerRefundEvent is a free log retrieval operation binding the contract event 0xafc3b014386561494c2478b62a479f4a42dbcb4ef28845094d473410b972d6d5.
//
// Solidity: event buyerRefundEvent(address userAddr)
func (_Trading *TradingFilterer) FilterBuyerRefundEvent(opts *bind.FilterOpts) (*TradingBuyerRefundEventIterator, error) {

	logs, sub, err := _Trading.contract.FilterLogs(opts, "buyerRefundEvent")
	if err != nil {
		return nil, err
	}
	return &TradingBuyerRefundEventIterator{contract: _Trading.contract, event: "buyerRefundEvent", logs: logs, sub: sub}, nil
}

// WatchBuyerRefundEvent is a free log subscription operation binding the contract event 0xafc3b014386561494c2478b62a479f4a42dbcb4ef28845094d473410b972d6d5.
//
// Solidity: event buyerRefundEvent(address userAddr)
func (_Trading *TradingFilterer) WatchBuyerRefundEvent(opts *bind.WatchOpts, sink chan<- *TradingBuyerRefundEvent) (event.Subscription, error) {

	logs, sub, err := _Trading.contract.WatchLogs(opts, "buyerRefundEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingBuyerRefundEvent)
				if err := _Trading.contract.UnpackLog(event, "buyerRefundEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBuyerRefundEvent is a log parse operation binding the contract event 0xafc3b014386561494c2478b62a479f4a42dbcb4ef28845094d473410b972d6d5.
//
// Solidity: event buyerRefundEvent(address userAddr)
func (_Trading *TradingFilterer) ParseBuyerRefundEvent(log types.Log) (*TradingBuyerRefundEvent, error) {
	event := new(TradingBuyerRefundEvent)
	if err := _Trading.contract.UnpackLog(event, "buyerRefundEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingChallengeEventIterator is returned from FilterChallengeEvent and is used to iterate over the raw logs and unpacked data for ChallengeEvent events raised by the Trading contract.
type TradingChallengeEventIterator struct {
	Event *TradingChallengeEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TradingChallengeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingChallengeEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TradingChallengeEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TradingChallengeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingChallengeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingChallengeEvent represents a ChallengeEvent event raised by the Trading contract.
type TradingChallengeEvent struct {
	Kc  [32]byte
	P   *big.Int
	P1  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChallengeEvent is a free log retrieval operation binding the contract event 0x038403778ce33b7bbdbd6c808eed8cacd1a5e8a78851fbcec058f610aef431ea.
//
// Solidity: event challengeEvent(bytes32 kc, uint256 p, uint256 p1)
func (_Trading *TradingFilterer) FilterChallengeEvent(opts *bind.FilterOpts) (*TradingChallengeEventIterator, error) {

	logs, sub, err := _Trading.contract.FilterLogs(opts, "challengeEvent")
	if err != nil {
		return nil, err
	}
	return &TradingChallengeEventIterator{contract: _Trading.contract, event: "challengeEvent", logs: logs, sub: sub}, nil
}

// WatchChallengeEvent is a free log subscription operation binding the contract event 0x038403778ce33b7bbdbd6c808eed8cacd1a5e8a78851fbcec058f610aef431ea.
//
// Solidity: event challengeEvent(bytes32 kc, uint256 p, uint256 p1)
func (_Trading *TradingFilterer) WatchChallengeEvent(opts *bind.WatchOpts, sink chan<- *TradingChallengeEvent) (event.Subscription, error) {

	logs, sub, err := _Trading.contract.WatchLogs(opts, "challengeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingChallengeEvent)
				if err := _Trading.contract.UnpackLog(event, "challengeEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChallengeEvent is a log parse operation binding the contract event 0x038403778ce33b7bbdbd6c808eed8cacd1a5e8a78851fbcec058f610aef431ea.
//
// Solidity: event challengeEvent(bytes32 kc, uint256 p, uint256 p1)
func (_Trading *TradingFilterer) ParseChallengeEvent(log types.Log) (*TradingChallengeEvent, error) {
	event := new(TradingChallengeEvent)
	if err := _Trading.contract.UnpackLog(event, "challengeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingCommitDataEventIterator is returned from FilterCommitDataEvent and is used to iterate over the raw logs and unpacked data for CommitDataEvent events raised by the Trading contract.
type TradingCommitDataEventIterator struct {
	Event *TradingCommitDataEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TradingCommitDataEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingCommitDataEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TradingCommitDataEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TradingCommitDataEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingCommitDataEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingCommitDataEvent represents a CommitDataEvent event raised by the Trading contract.
type TradingCommitDataEvent struct {
	Id     *big.Int
	Delta1 [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCommitDataEvent is a free log retrieval operation binding the contract event 0xcb894139878da1d74d95c1725159680c62225fbee97f6af6f79589b5b3115226.
//
// Solidity: event commitDataEvent(uint256 id, bytes32 delta_1)
func (_Trading *TradingFilterer) FilterCommitDataEvent(opts *bind.FilterOpts) (*TradingCommitDataEventIterator, error) {

	logs, sub, err := _Trading.contract.FilterLogs(opts, "commitDataEvent")
	if err != nil {
		return nil, err
	}
	return &TradingCommitDataEventIterator{contract: _Trading.contract, event: "commitDataEvent", logs: logs, sub: sub}, nil
}

// WatchCommitDataEvent is a free log subscription operation binding the contract event 0xcb894139878da1d74d95c1725159680c62225fbee97f6af6f79589b5b3115226.
//
// Solidity: event commitDataEvent(uint256 id, bytes32 delta_1)
func (_Trading *TradingFilterer) WatchCommitDataEvent(opts *bind.WatchOpts, sink chan<- *TradingCommitDataEvent) (event.Subscription, error) {

	logs, sub, err := _Trading.contract.WatchLogs(opts, "commitDataEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingCommitDataEvent)
				if err := _Trading.contract.UnpackLog(event, "commitDataEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCommitDataEvent is a log parse operation binding the contract event 0xcb894139878da1d74d95c1725159680c62225fbee97f6af6f79589b5b3115226.
//
// Solidity: event commitDataEvent(uint256 id, bytes32 delta_1)
func (_Trading *TradingFilterer) ParseCommitDataEvent(log types.Log) (*TradingCommitDataEvent, error) {
	event := new(TradingCommitDataEvent)
	if err := _Trading.contract.UnpackLog(event, "commitDataEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingInitiationArbitrationEventIterator is returned from FilterInitiationArbitrationEvent and is used to iterate over the raw logs and unpacked data for InitiationArbitrationEvent events raised by the Trading contract.
type TradingInitiationArbitrationEventIterator struct {
	Event *TradingInitiationArbitrationEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TradingInitiationArbitrationEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingInitiationArbitrationEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TradingInitiationArbitrationEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TradingInitiationArbitrationEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingInitiationArbitrationEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingInitiationArbitrationEvent represents a InitiationArbitrationEvent event raised by the Trading contract.
type TradingInitiationArbitrationEvent struct {
	UserAddr      common.Address
	ComplaintInfo [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInitiationArbitrationEvent is a free log retrieval operation binding the contract event 0x09d1ead43ef479358ea86e42900818e94168115b36cca73175cfca8fbe6d8db6.
//
// Solidity: event initiationArbitrationEvent(address userAddr, bytes32 complaintInfo)
func (_Trading *TradingFilterer) FilterInitiationArbitrationEvent(opts *bind.FilterOpts) (*TradingInitiationArbitrationEventIterator, error) {

	logs, sub, err := _Trading.contract.FilterLogs(opts, "initiationArbitrationEvent")
	if err != nil {
		return nil, err
	}
	return &TradingInitiationArbitrationEventIterator{contract: _Trading.contract, event: "initiationArbitrationEvent", logs: logs, sub: sub}, nil
}

// WatchInitiationArbitrationEvent is a free log subscription operation binding the contract event 0x09d1ead43ef479358ea86e42900818e94168115b36cca73175cfca8fbe6d8db6.
//
// Solidity: event initiationArbitrationEvent(address userAddr, bytes32 complaintInfo)
func (_Trading *TradingFilterer) WatchInitiationArbitrationEvent(opts *bind.WatchOpts, sink chan<- *TradingInitiationArbitrationEvent) (event.Subscription, error) {

	logs, sub, err := _Trading.contract.WatchLogs(opts, "initiationArbitrationEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingInitiationArbitrationEvent)
				if err := _Trading.contract.UnpackLog(event, "initiationArbitrationEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitiationArbitrationEvent is a log parse operation binding the contract event 0x09d1ead43ef479358ea86e42900818e94168115b36cca73175cfca8fbe6d8db6.
//
// Solidity: event initiationArbitrationEvent(address userAddr, bytes32 complaintInfo)
func (_Trading *TradingFilterer) ParseInitiationArbitrationEvent(log types.Log) (*TradingInitiationArbitrationEvent, error) {
	event := new(TradingInitiationArbitrationEvent)
	if err := _Trading.contract.UnpackLog(event, "initiationArbitrationEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingOnchianArbitrationEventIterator is returned from FilterOnchianArbitrationEvent and is used to iterate over the raw logs and unpacked data for OnchianArbitrationEvent events raised by the Trading contract.
type TradingOnchianArbitrationEventIterator struct {
	Event *TradingOnchianArbitrationEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TradingOnchianArbitrationEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingOnchianArbitrationEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TradingOnchianArbitrationEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TradingOnchianArbitrationEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingOnchianArbitrationEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingOnchianArbitrationEvent represents a OnchianArbitrationEvent event raised by the Trading contract.
type TradingOnchianArbitrationEvent struct {
	Status *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOnchianArbitrationEvent is a free log retrieval operation binding the contract event 0x6975f323332c2d66d668c2be6e73ef7173309d6252480fe9fe18895d4e7b93d3.
//
// Solidity: event onchianArbitrationEvent(uint256 status)
func (_Trading *TradingFilterer) FilterOnchianArbitrationEvent(opts *bind.FilterOpts) (*TradingOnchianArbitrationEventIterator, error) {

	logs, sub, err := _Trading.contract.FilterLogs(opts, "onchianArbitrationEvent")
	if err != nil {
		return nil, err
	}
	return &TradingOnchianArbitrationEventIterator{contract: _Trading.contract, event: "onchianArbitrationEvent", logs: logs, sub: sub}, nil
}

// WatchOnchianArbitrationEvent is a free log subscription operation binding the contract event 0x6975f323332c2d66d668c2be6e73ef7173309d6252480fe9fe18895d4e7b93d3.
//
// Solidity: event onchianArbitrationEvent(uint256 status)
func (_Trading *TradingFilterer) WatchOnchianArbitrationEvent(opts *bind.WatchOpts, sink chan<- *TradingOnchianArbitrationEvent) (event.Subscription, error) {

	logs, sub, err := _Trading.contract.WatchLogs(opts, "onchianArbitrationEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingOnchianArbitrationEvent)
				if err := _Trading.contract.UnpackLog(event, "onchianArbitrationEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOnchianArbitrationEvent is a log parse operation binding the contract event 0x6975f323332c2d66d668c2be6e73ef7173309d6252480fe9fe18895d4e7b93d3.
//
// Solidity: event onchianArbitrationEvent(uint256 status)
func (_Trading *TradingFilterer) ParseOnchianArbitrationEvent(log types.Log) (*TradingOnchianArbitrationEvent, error) {
	event := new(TradingOnchianArbitrationEvent)
	if err := _Trading.contract.UnpackLog(event, "onchianArbitrationEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingPayleftEventIterator is returned from FilterPayleftEvent and is used to iterate over the raw logs and unpacked data for PayleftEvent events raised by the Trading contract.
type TradingPayleftEventIterator struct {
	Event *TradingPayleftEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TradingPayleftEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingPayleftEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TradingPayleftEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TradingPayleftEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingPayleftEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingPayleftEvent represents a PayleftEvent event raised by the Trading contract.
type TradingPayleftEvent struct {
	UserAddr common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPayleftEvent is a free log retrieval operation binding the contract event 0xe37b4d92bfb3a0f51b4632930020ab1670830e1c71576288d9b33f4671f09bb9.
//
// Solidity: event payleftEvent(address userAddr, uint256 value)
func (_Trading *TradingFilterer) FilterPayleftEvent(opts *bind.FilterOpts) (*TradingPayleftEventIterator, error) {

	logs, sub, err := _Trading.contract.FilterLogs(opts, "payleftEvent")
	if err != nil {
		return nil, err
	}
	return &TradingPayleftEventIterator{contract: _Trading.contract, event: "payleftEvent", logs: logs, sub: sub}, nil
}

// WatchPayleftEvent is a free log subscription operation binding the contract event 0xe37b4d92bfb3a0f51b4632930020ab1670830e1c71576288d9b33f4671f09bb9.
//
// Solidity: event payleftEvent(address userAddr, uint256 value)
func (_Trading *TradingFilterer) WatchPayleftEvent(opts *bind.WatchOpts, sink chan<- *TradingPayleftEvent) (event.Subscription, error) {

	logs, sub, err := _Trading.contract.WatchLogs(opts, "payleftEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingPayleftEvent)
				if err := _Trading.contract.UnpackLog(event, "payleftEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePayleftEvent is a log parse operation binding the contract event 0xe37b4d92bfb3a0f51b4632930020ab1670830e1c71576288d9b33f4671f09bb9.
//
// Solidity: event payleftEvent(address userAddr, uint256 value)
func (_Trading *TradingFilterer) ParsePayleftEvent(log types.Log) (*TradingPayleftEvent, error) {
	event := new(TradingPayleftEvent)
	if err := _Trading.contract.UnpackLog(event, "payleftEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingReleaseMasterkeyEventIterator is returned from FilterReleaseMasterkeyEvent and is used to iterate over the raw logs and unpacked data for ReleaseMasterkeyEvent events raised by the Trading contract.
type TradingReleaseMasterkeyEventIterator struct {
	Event *TradingReleaseMasterkeyEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TradingReleaseMasterkeyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingReleaseMasterkeyEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TradingReleaseMasterkeyEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TradingReleaseMasterkeyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingReleaseMasterkeyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingReleaseMasterkeyEvent represents a ReleaseMasterkeyEvent event raised by the Trading contract.
type TradingReleaseMasterkeyEvent struct {
	Masterkey [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReleaseMasterkeyEvent is a free log retrieval operation binding the contract event 0x59e2b1bea451053289971062b187c027a0f841853f70e131a5a9f22b05cd44f8.
//
// Solidity: event releaseMasterkeyEvent(bytes32 masterkey)
func (_Trading *TradingFilterer) FilterReleaseMasterkeyEvent(opts *bind.FilterOpts) (*TradingReleaseMasterkeyEventIterator, error) {

	logs, sub, err := _Trading.contract.FilterLogs(opts, "releaseMasterkeyEvent")
	if err != nil {
		return nil, err
	}
	return &TradingReleaseMasterkeyEventIterator{contract: _Trading.contract, event: "releaseMasterkeyEvent", logs: logs, sub: sub}, nil
}

// WatchReleaseMasterkeyEvent is a free log subscription operation binding the contract event 0x59e2b1bea451053289971062b187c027a0f841853f70e131a5a9f22b05cd44f8.
//
// Solidity: event releaseMasterkeyEvent(bytes32 masterkey)
func (_Trading *TradingFilterer) WatchReleaseMasterkeyEvent(opts *bind.WatchOpts, sink chan<- *TradingReleaseMasterkeyEvent) (event.Subscription, error) {

	logs, sub, err := _Trading.contract.WatchLogs(opts, "releaseMasterkeyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingReleaseMasterkeyEvent)
				if err := _Trading.contract.UnpackLog(event, "releaseMasterkeyEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReleaseMasterkeyEvent is a log parse operation binding the contract event 0x59e2b1bea451053289971062b187c027a0f841853f70e131a5a9f22b05cd44f8.
//
// Solidity: event releaseMasterkeyEvent(bytes32 masterkey)
func (_Trading *TradingFilterer) ParseReleaseMasterkeyEvent(log types.Log) (*TradingReleaseMasterkeyEvent, error) {
	event := new(TradingReleaseMasterkeyEvent)
	if err := _Trading.contract.UnpackLog(event, "releaseMasterkeyEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingRespondeEventIterator is returned from FilterRespondeEvent and is used to iterate over the raw logs and unpacked data for RespondeEvent events raised by the Trading contract.
type TradingRespondeEventIterator struct {
	Event *TradingRespondeEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TradingRespondeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingRespondeEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TradingRespondeEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TradingRespondeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingRespondeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingRespondeEvent represents a RespondeEvent event raised by the Trading contract.
type TradingRespondeEvent struct {
	Delta2  [32]byte
	Length2 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRespondeEvent is a free log retrieval operation binding the contract event 0x8367cc6ba01f36350b0d5f1caaa3afa155152dac000940ae82cd327bfa0bc827.
//
// Solidity: event respondeEvent(bytes32 _delta_2, uint256 _length2)
func (_Trading *TradingFilterer) FilterRespondeEvent(opts *bind.FilterOpts) (*TradingRespondeEventIterator, error) {

	logs, sub, err := _Trading.contract.FilterLogs(opts, "respondeEvent")
	if err != nil {
		return nil, err
	}
	return &TradingRespondeEventIterator{contract: _Trading.contract, event: "respondeEvent", logs: logs, sub: sub}, nil
}

// WatchRespondeEvent is a free log subscription operation binding the contract event 0x8367cc6ba01f36350b0d5f1caaa3afa155152dac000940ae82cd327bfa0bc827.
//
// Solidity: event respondeEvent(bytes32 _delta_2, uint256 _length2)
func (_Trading *TradingFilterer) WatchRespondeEvent(opts *bind.WatchOpts, sink chan<- *TradingRespondeEvent) (event.Subscription, error) {

	logs, sub, err := _Trading.contract.WatchLogs(opts, "respondeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingRespondeEvent)
				if err := _Trading.contract.UnpackLog(event, "respondeEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRespondeEvent is a log parse operation binding the contract event 0x8367cc6ba01f36350b0d5f1caaa3afa155152dac000940ae82cd327bfa0bc827.
//
// Solidity: event respondeEvent(bytes32 _delta_2, uint256 _length2)
func (_Trading *TradingFilterer) ParseRespondeEvent(log types.Log) (*TradingRespondeEvent, error) {
	event := new(TradingRespondeEvent)
	if err := _Trading.contract.UnpackLog(event, "respondeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingVerdictEventIterator is returned from FilterVerdictEvent and is used to iterate over the raw logs and unpacked data for VerdictEvent events raised by the Trading contract.
type TradingVerdictEventIterator struct {
	Event *TradingVerdictEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TradingVerdictEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingVerdictEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TradingVerdictEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TradingVerdictEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingVerdictEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingVerdictEvent represents a VerdictEvent event raised by the Trading contract.
type TradingVerdictEvent struct {
	UserAddr common.Address
	Winner   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVerdictEvent is a free log retrieval operation binding the contract event 0x5b6a2b2de4fa5ff7a980e48d8fe9471ff3ec4eeef693a29dae080e6a5a9c786f.
//
// Solidity: event verdictEvent(address userAddr, address winner)
func (_Trading *TradingFilterer) FilterVerdictEvent(opts *bind.FilterOpts) (*TradingVerdictEventIterator, error) {

	logs, sub, err := _Trading.contract.FilterLogs(opts, "verdictEvent")
	if err != nil {
		return nil, err
	}
	return &TradingVerdictEventIterator{contract: _Trading.contract, event: "verdictEvent", logs: logs, sub: sub}, nil
}

// WatchVerdictEvent is a free log subscription operation binding the contract event 0x5b6a2b2de4fa5ff7a980e48d8fe9471ff3ec4eeef693a29dae080e6a5a9c786f.
//
// Solidity: event verdictEvent(address userAddr, address winner)
func (_Trading *TradingFilterer) WatchVerdictEvent(opts *bind.WatchOpts, sink chan<- *TradingVerdictEvent) (event.Subscription, error) {

	logs, sub, err := _Trading.contract.WatchLogs(opts, "verdictEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingVerdictEvent)
				if err := _Trading.contract.UnpackLog(event, "verdictEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVerdictEvent is a log parse operation binding the contract event 0x5b6a2b2de4fa5ff7a980e48d8fe9471ff3ec4eeef693a29dae080e6a5a9c786f.
//
// Solidity: event verdictEvent(address userAddr, address winner)
func (_Trading *TradingFilterer) ParseVerdictEvent(log types.Log) (*TradingVerdictEvent, error) {
	event := new(TradingVerdictEvent)
	if err := _Trading.contract.UnpackLog(event, "verdictEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
