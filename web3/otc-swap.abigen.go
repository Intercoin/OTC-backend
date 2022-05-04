// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package web3

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

// OTCSwapMetaData contains all meta data concerning the OTCSwap contract.
var OTCSwapMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"tradeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"tradeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Engaged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"tradeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"poster\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxPenalty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"NewTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"tradeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_tradeHash\",\"type\":\"bytes32\"}],\"name\":\"calculatePenalty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_tradeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"_signatures\",\"type\":\"bytes[]\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_tradeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"engage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_tradeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getSignature\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignatureNow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_tradeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiverAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_penaltyAmount\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transfers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"poster\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"enumTradeRegistratorERC20Test.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxPenalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawPenalty\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_tradeHash\",\"type\":\"bytes32\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OTCSwapABI is the input ABI used to generate the binding from.
// Deprecated: Use OTCSwapMetaData.ABI instead.
var OTCSwapABI = OTCSwapMetaData.ABI

// OTCSwap is an auto generated Go binding around an Ethereum contract.
type OTCSwap struct {
	OTCSwapCaller     // Read-only binding to the contract
	OTCSwapTransactor // Write-only binding to the contract
	OTCSwapFilterer   // Log filterer for contract events
}

// OTCSwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type OTCSwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OTCSwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OTCSwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OTCSwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OTCSwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OTCSwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OTCSwapSession struct {
	Contract     *OTCSwap          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OTCSwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OTCSwapCallerSession struct {
	Contract *OTCSwapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OTCSwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OTCSwapTransactorSession struct {
	Contract     *OTCSwapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OTCSwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type OTCSwapRaw struct {
	Contract *OTCSwap // Generic contract binding to access the raw methods on
}

// OTCSwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OTCSwapCallerRaw struct {
	Contract *OTCSwapCaller // Generic read-only contract binding to access the raw methods on
}

// OTCSwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OTCSwapTransactorRaw struct {
	Contract *OTCSwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOTCSwap creates a new instance of OTCSwap, bound to a specific deployed contract.
func NewOTCSwap(address common.Address, backend bind.ContractBackend) (*OTCSwap, error) {
	contract, err := bindOTCSwap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OTCSwap{OTCSwapCaller: OTCSwapCaller{contract: contract}, OTCSwapTransactor: OTCSwapTransactor{contract: contract}, OTCSwapFilterer: OTCSwapFilterer{contract: contract}}, nil
}

// NewOTCSwapCaller creates a new read-only instance of OTCSwap, bound to a specific deployed contract.
func NewOTCSwapCaller(address common.Address, caller bind.ContractCaller) (*OTCSwapCaller, error) {
	contract, err := bindOTCSwap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OTCSwapCaller{contract: contract}, nil
}

// NewOTCSwapTransactor creates a new write-only instance of OTCSwap, bound to a specific deployed contract.
func NewOTCSwapTransactor(address common.Address, transactor bind.ContractTransactor) (*OTCSwapTransactor, error) {
	contract, err := bindOTCSwap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OTCSwapTransactor{contract: contract}, nil
}

// NewOTCSwapFilterer creates a new log filterer instance of OTCSwap, bound to a specific deployed contract.
func NewOTCSwapFilterer(address common.Address, filterer bind.ContractFilterer) (*OTCSwapFilterer, error) {
	contract, err := bindOTCSwap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OTCSwapFilterer{contract: contract}, nil
}

// bindOTCSwap binds a generic wrapper to an already deployed contract.
func bindOTCSwap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OTCSwapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OTCSwap *OTCSwapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OTCSwap.Contract.OTCSwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OTCSwap *OTCSwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OTCSwap.Contract.OTCSwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OTCSwap *OTCSwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OTCSwap.Contract.OTCSwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OTCSwap *OTCSwapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OTCSwap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OTCSwap *OTCSwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OTCSwap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OTCSwap *OTCSwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OTCSwap.Contract.contract.Transact(opts, method, params...)
}

// CalculatePenalty is a free data retrieval call binding the contract method 0xd483be76.
//
// Solidity: function calculatePenalty(bytes32 _tradeHash) view returns(uint256)
func (_OTCSwap *OTCSwapCaller) CalculatePenalty(opts *bind.CallOpts, _tradeHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _OTCSwap.contract.Call(opts, &out, "calculatePenalty", _tradeHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePenalty is a free data retrieval call binding the contract method 0xd483be76.
//
// Solidity: function calculatePenalty(bytes32 _tradeHash) view returns(uint256)
func (_OTCSwap *OTCSwapSession) CalculatePenalty(_tradeHash [32]byte) (*big.Int, error) {
	return _OTCSwap.Contract.CalculatePenalty(&_OTCSwap.CallOpts, _tradeHash)
}

// CalculatePenalty is a free data retrieval call binding the contract method 0xd483be76.
//
// Solidity: function calculatePenalty(bytes32 _tradeHash) view returns(uint256)
func (_OTCSwap *OTCSwapCallerSession) CalculatePenalty(_tradeHash [32]byte) (*big.Int, error) {
	return _OTCSwap.Contract.CalculatePenalty(&_OTCSwap.CallOpts, _tradeHash)
}

// GetSignature is a free data retrieval call binding the contract method 0x88c5f59a.
//
// Solidity: function getSignature(bytes32 _tradeHash, uint256 _index) view returns(bytes)
func (_OTCSwap *OTCSwapCaller) GetSignature(opts *bind.CallOpts, _tradeHash [32]byte, _index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _OTCSwap.contract.Call(opts, &out, "getSignature", _tradeHash, _index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetSignature is a free data retrieval call binding the contract method 0x88c5f59a.
//
// Solidity: function getSignature(bytes32 _tradeHash, uint256 _index) view returns(bytes)
func (_OTCSwap *OTCSwapSession) GetSignature(_tradeHash [32]byte, _index *big.Int) ([]byte, error) {
	return _OTCSwap.Contract.GetSignature(&_OTCSwap.CallOpts, _tradeHash, _index)
}

// GetSignature is a free data retrieval call binding the contract method 0x88c5f59a.
//
// Solidity: function getSignature(bytes32 _tradeHash, uint256 _index) view returns(bytes)
func (_OTCSwap *OTCSwapCallerSession) GetSignature(_tradeHash [32]byte, _index *big.Int) ([]byte, error) {
	return _OTCSwap.Contract.GetSignature(&_OTCSwap.CallOpts, _tradeHash, _index)
}

// IsValidSignatureNow is a free data retrieval call binding the contract method 0x6ccea652.
//
// Solidity: function isValidSignatureNow(address signer, bytes32 hash, bytes signature) view returns(bool)
func (_OTCSwap *OTCSwapCaller) IsValidSignatureNow(opts *bind.CallOpts, signer common.Address, hash [32]byte, signature []byte) (bool, error) {
	var out []interface{}
	err := _OTCSwap.contract.Call(opts, &out, "isValidSignatureNow", signer, hash, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidSignatureNow is a free data retrieval call binding the contract method 0x6ccea652.
//
// Solidity: function isValidSignatureNow(address signer, bytes32 hash, bytes signature) view returns(bool)
func (_OTCSwap *OTCSwapSession) IsValidSignatureNow(signer common.Address, hash [32]byte, signature []byte) (bool, error) {
	return _OTCSwap.Contract.IsValidSignatureNow(&_OTCSwap.CallOpts, signer, hash, signature)
}

// IsValidSignatureNow is a free data retrieval call binding the contract method 0x6ccea652.
//
// Solidity: function isValidSignatureNow(address signer, bytes32 hash, bytes signature) view returns(bool)
func (_OTCSwap *OTCSwapCallerSession) IsValidSignatureNow(signer common.Address, hash [32]byte, signature []byte) (bool, error) {
	return _OTCSwap.Contract.IsValidSignatureNow(&_OTCSwap.CallOpts, signer, hash, signature)
}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() view returns(uint256)
func (_OTCSwap *OTCSwapCaller) LockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OTCSwap.contract.Call(opts, &out, "lockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() view returns(uint256)
func (_OTCSwap *OTCSwapSession) LockTime() (*big.Int, error) {
	return _OTCSwap.Contract.LockTime(&_OTCSwap.CallOpts)
}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() view returns(uint256)
func (_OTCSwap *OTCSwapCallerSession) LockTime() (*big.Int, error) {
	return _OTCSwap.Contract.LockTime(&_OTCSwap.CallOpts)
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) view returns(address poster, address asset, uint256 amount, address receiver, uint8 status, uint256 maxPenalty, uint256 deadline, uint256 withdrawPenalty)
func (_OTCSwap *OTCSwapCaller) Transfers(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Poster          common.Address
	Asset           common.Address
	Amount          *big.Int
	Receiver        common.Address
	Status          uint8
	MaxPenalty      *big.Int
	Deadline        *big.Int
	WithdrawPenalty *big.Int
}, error) {
	var out []interface{}
	err := _OTCSwap.contract.Call(opts, &out, "transfers", arg0)

	outstruct := new(struct {
		Poster          common.Address
		Asset           common.Address
		Amount          *big.Int
		Receiver        common.Address
		Status          uint8
		MaxPenalty      *big.Int
		Deadline        *big.Int
		WithdrawPenalty *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Poster = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Asset = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Receiver = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.MaxPenalty = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Deadline = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.WithdrawPenalty = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) view returns(address poster, address asset, uint256 amount, address receiver, uint8 status, uint256 maxPenalty, uint256 deadline, uint256 withdrawPenalty)
func (_OTCSwap *OTCSwapSession) Transfers(arg0 [32]byte) (struct {
	Poster          common.Address
	Asset           common.Address
	Amount          *big.Int
	Receiver        common.Address
	Status          uint8
	MaxPenalty      *big.Int
	Deadline        *big.Int
	WithdrawPenalty *big.Int
}, error) {
	return _OTCSwap.Contract.Transfers(&_OTCSwap.CallOpts, arg0)
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) view returns(address poster, address asset, uint256 amount, address receiver, uint8 status, uint256 maxPenalty, uint256 deadline, uint256 withdrawPenalty)
func (_OTCSwap *OTCSwapCallerSession) Transfers(arg0 [32]byte) (struct {
	Poster          common.Address
	Asset           common.Address
	Amount          *big.Int
	Receiver        common.Address
	Status          uint8
	MaxPenalty      *big.Int
	Deadline        *big.Int
	WithdrawPenalty *big.Int
}, error) {
	return _OTCSwap.Contract.Transfers(&_OTCSwap.CallOpts, arg0)
}

// Claim is a paid mutator transaction binding the contract method 0xfd3453a9.
//
// Solidity: function claim(bytes32 _tradeHash, bytes[] _signatures) payable returns()
func (_OTCSwap *OTCSwapTransactor) Claim(opts *bind.TransactOpts, _tradeHash [32]byte, _signatures [][]byte) (*types.Transaction, error) {
	return _OTCSwap.contract.Transact(opts, "claim", _tradeHash, _signatures)
}

// Claim is a paid mutator transaction binding the contract method 0xfd3453a9.
//
// Solidity: function claim(bytes32 _tradeHash, bytes[] _signatures) payable returns()
func (_OTCSwap *OTCSwapSession) Claim(_tradeHash [32]byte, _signatures [][]byte) (*types.Transaction, error) {
	return _OTCSwap.Contract.Claim(&_OTCSwap.TransactOpts, _tradeHash, _signatures)
}

// Claim is a paid mutator transaction binding the contract method 0xfd3453a9.
//
// Solidity: function claim(bytes32 _tradeHash, bytes[] _signatures) payable returns()
func (_OTCSwap *OTCSwapTransactorSession) Claim(_tradeHash [32]byte, _signatures [][]byte) (*types.Transaction, error) {
	return _OTCSwap.Contract.Claim(&_OTCSwap.TransactOpts, _tradeHash, _signatures)
}

// Engage is a paid mutator transaction binding the contract method 0x515e2697.
//
// Solidity: function engage(bytes32 _tradeHash, bytes _signature) returns()
func (_OTCSwap *OTCSwapTransactor) Engage(opts *bind.TransactOpts, _tradeHash [32]byte, _signature []byte) (*types.Transaction, error) {
	return _OTCSwap.contract.Transact(opts, "engage", _tradeHash, _signature)
}

// Engage is a paid mutator transaction binding the contract method 0x515e2697.
//
// Solidity: function engage(bytes32 _tradeHash, bytes _signature) returns()
func (_OTCSwap *OTCSwapSession) Engage(_tradeHash [32]byte, _signature []byte) (*types.Transaction, error) {
	return _OTCSwap.Contract.Engage(&_OTCSwap.TransactOpts, _tradeHash, _signature)
}

// Engage is a paid mutator transaction binding the contract method 0x515e2697.
//
// Solidity: function engage(bytes32 _tradeHash, bytes _signature) returns()
func (_OTCSwap *OTCSwapTransactorSession) Engage(_tradeHash [32]byte, _signature []byte) (*types.Transaction, error) {
	return _OTCSwap.Contract.Engage(&_OTCSwap.TransactOpts, _tradeHash, _signature)
}

// Lock is a paid mutator transaction binding the contract method 0x91644b2b.
//
// Solidity: function lock(bytes32 _tradeHash, uint256 _amount, address _tokenAddress, address _receiverAddress, uint256 _penaltyAmount) returns()
func (_OTCSwap *OTCSwapTransactor) Lock(opts *bind.TransactOpts, _tradeHash [32]byte, _amount *big.Int, _tokenAddress common.Address, _receiverAddress common.Address, _penaltyAmount *big.Int) (*types.Transaction, error) {
	return _OTCSwap.contract.Transact(opts, "lock", _tradeHash, _amount, _tokenAddress, _receiverAddress, _penaltyAmount)
}

// Lock is a paid mutator transaction binding the contract method 0x91644b2b.
//
// Solidity: function lock(bytes32 _tradeHash, uint256 _amount, address _tokenAddress, address _receiverAddress, uint256 _penaltyAmount) returns()
func (_OTCSwap *OTCSwapSession) Lock(_tradeHash [32]byte, _amount *big.Int, _tokenAddress common.Address, _receiverAddress common.Address, _penaltyAmount *big.Int) (*types.Transaction, error) {
	return _OTCSwap.Contract.Lock(&_OTCSwap.TransactOpts, _tradeHash, _amount, _tokenAddress, _receiverAddress, _penaltyAmount)
}

// Lock is a paid mutator transaction binding the contract method 0x91644b2b.
//
// Solidity: function lock(bytes32 _tradeHash, uint256 _amount, address _tokenAddress, address _receiverAddress, uint256 _penaltyAmount) returns()
func (_OTCSwap *OTCSwapTransactorSession) Lock(_tradeHash [32]byte, _amount *big.Int, _tokenAddress common.Address, _receiverAddress common.Address, _penaltyAmount *big.Int) (*types.Transaction, error) {
	return _OTCSwap.Contract.Lock(&_OTCSwap.TransactOpts, _tradeHash, _amount, _tokenAddress, _receiverAddress, _penaltyAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x8e19899e.
//
// Solidity: function withdraw(bytes32 _tradeHash) returns()
func (_OTCSwap *OTCSwapTransactor) Withdraw(opts *bind.TransactOpts, _tradeHash [32]byte) (*types.Transaction, error) {
	return _OTCSwap.contract.Transact(opts, "withdraw", _tradeHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x8e19899e.
//
// Solidity: function withdraw(bytes32 _tradeHash) returns()
func (_OTCSwap *OTCSwapSession) Withdraw(_tradeHash [32]byte) (*types.Transaction, error) {
	return _OTCSwap.Contract.Withdraw(&_OTCSwap.TransactOpts, _tradeHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x8e19899e.
//
// Solidity: function withdraw(bytes32 _tradeHash) returns()
func (_OTCSwap *OTCSwapTransactorSession) Withdraw(_tradeHash [32]byte) (*types.Transaction, error) {
	return _OTCSwap.Contract.Withdraw(&_OTCSwap.TransactOpts, _tradeHash)
}

// OTCSwapClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the OTCSwap contract.
type OTCSwapClaimedIterator struct {
	Event *OTCSwapClaimed // Event containing the contract specifics and raw log

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
func (it *OTCSwapClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OTCSwapClaimed)
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
		it.Event = new(OTCSwapClaimed)
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
func (it *OTCSwapClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OTCSwapClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OTCSwapClaimed represents a Claimed event raised by the OTCSwap contract.
type OTCSwapClaimed struct {
	TradeHash  [32]byte
	Signatures [][]byte
	Sender     common.Address
	Penalty    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xbaf7454a8f022bbcd5b42214ff4e1f19107bb64160cd1e95fcd2749f1a8a13a5.
//
// Solidity: event Claimed(bytes32 tradeHash, bytes[] signatures, address sender, uint256 penalty)
func (_OTCSwap *OTCSwapFilterer) FilterClaimed(opts *bind.FilterOpts) (*OTCSwapClaimedIterator, error) {

	logs, sub, err := _OTCSwap.contract.FilterLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return &OTCSwapClaimedIterator{contract: _OTCSwap.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xbaf7454a8f022bbcd5b42214ff4e1f19107bb64160cd1e95fcd2749f1a8a13a5.
//
// Solidity: event Claimed(bytes32 tradeHash, bytes[] signatures, address sender, uint256 penalty)
func (_OTCSwap *OTCSwapFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *OTCSwapClaimed) (event.Subscription, error) {

	logs, sub, err := _OTCSwap.contract.WatchLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OTCSwapClaimed)
				if err := _OTCSwap.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0xbaf7454a8f022bbcd5b42214ff4e1f19107bb64160cd1e95fcd2749f1a8a13a5.
//
// Solidity: event Claimed(bytes32 tradeHash, bytes[] signatures, address sender, uint256 penalty)
func (_OTCSwap *OTCSwapFilterer) ParseClaimed(log types.Log) (*OTCSwapClaimed, error) {
	event := new(OTCSwapClaimed)
	if err := _OTCSwap.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OTCSwapEngagedIterator is returned from FilterEngaged and is used to iterate over the raw logs and unpacked data for Engaged events raised by the OTCSwap contract.
type OTCSwapEngagedIterator struct {
	Event *OTCSwapEngaged // Event containing the contract specifics and raw log

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
func (it *OTCSwapEngagedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OTCSwapEngaged)
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
		it.Event = new(OTCSwapEngaged)
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
func (it *OTCSwapEngagedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OTCSwapEngagedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OTCSwapEngaged represents a Engaged event raised by the OTCSwap contract.
type OTCSwapEngaged struct {
	TradeHash [32]byte
	Signature []byte
	Sender    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEngaged is a free log retrieval operation binding the contract event 0xd8b05fcd2802ff12ac0b7fcc02183136df63dbfb6a9ba80192a67b1ee8c9785d.
//
// Solidity: event Engaged(bytes32 tradeHash, bytes signature, address sender)
func (_OTCSwap *OTCSwapFilterer) FilterEngaged(opts *bind.FilterOpts) (*OTCSwapEngagedIterator, error) {

	logs, sub, err := _OTCSwap.contract.FilterLogs(opts, "Engaged")
	if err != nil {
		return nil, err
	}
	return &OTCSwapEngagedIterator{contract: _OTCSwap.contract, event: "Engaged", logs: logs, sub: sub}, nil
}

// WatchEngaged is a free log subscription operation binding the contract event 0xd8b05fcd2802ff12ac0b7fcc02183136df63dbfb6a9ba80192a67b1ee8c9785d.
//
// Solidity: event Engaged(bytes32 tradeHash, bytes signature, address sender)
func (_OTCSwap *OTCSwapFilterer) WatchEngaged(opts *bind.WatchOpts, sink chan<- *OTCSwapEngaged) (event.Subscription, error) {

	logs, sub, err := _OTCSwap.contract.WatchLogs(opts, "Engaged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OTCSwapEngaged)
				if err := _OTCSwap.contract.UnpackLog(event, "Engaged", log); err != nil {
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

// ParseEngaged is a log parse operation binding the contract event 0xd8b05fcd2802ff12ac0b7fcc02183136df63dbfb6a9ba80192a67b1ee8c9785d.
//
// Solidity: event Engaged(bytes32 tradeHash, bytes signature, address sender)
func (_OTCSwap *OTCSwapFilterer) ParseEngaged(log types.Log) (*OTCSwapEngaged, error) {
	event := new(OTCSwapEngaged)
	if err := _OTCSwap.contract.UnpackLog(event, "Engaged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OTCSwapNewTransferIterator is returned from FilterNewTransfer and is used to iterate over the raw logs and unpacked data for NewTransfer events raised by the OTCSwap contract.
type OTCSwapNewTransferIterator struct {
	Event *OTCSwapNewTransfer // Event containing the contract specifics and raw log

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
func (it *OTCSwapNewTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OTCSwapNewTransfer)
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
		it.Event = new(OTCSwapNewTransfer)
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
func (it *OTCSwapNewTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OTCSwapNewTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OTCSwapNewTransfer represents a NewTransfer event raised by the OTCSwap contract.
type OTCSwapNewTransfer struct {
	TradeHash  [32]byte
	Poster     common.Address
	Asset      common.Address
	Amount     *big.Int
	Receiver   common.Address
	MaxPenalty *big.Int
	Deadline   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewTransfer is a free log retrieval operation binding the contract event 0x323a808bb16fb9488329579058ce4f59f2ede792393f123e166946663305ae46.
//
// Solidity: event NewTransfer(bytes32 tradeHash, address poster, address asset, uint256 amount, address receiver, uint256 maxPenalty, uint256 deadline)
func (_OTCSwap *OTCSwapFilterer) FilterNewTransfer(opts *bind.FilterOpts) (*OTCSwapNewTransferIterator, error) {

	logs, sub, err := _OTCSwap.contract.FilterLogs(opts, "NewTransfer")
	if err != nil {
		return nil, err
	}
	return &OTCSwapNewTransferIterator{contract: _OTCSwap.contract, event: "NewTransfer", logs: logs, sub: sub}, nil
}

// WatchNewTransfer is a free log subscription operation binding the contract event 0x323a808bb16fb9488329579058ce4f59f2ede792393f123e166946663305ae46.
//
// Solidity: event NewTransfer(bytes32 tradeHash, address poster, address asset, uint256 amount, address receiver, uint256 maxPenalty, uint256 deadline)
func (_OTCSwap *OTCSwapFilterer) WatchNewTransfer(opts *bind.WatchOpts, sink chan<- *OTCSwapNewTransfer) (event.Subscription, error) {

	logs, sub, err := _OTCSwap.contract.WatchLogs(opts, "NewTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OTCSwapNewTransfer)
				if err := _OTCSwap.contract.UnpackLog(event, "NewTransfer", log); err != nil {
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

// ParseNewTransfer is a log parse operation binding the contract event 0x323a808bb16fb9488329579058ce4f59f2ede792393f123e166946663305ae46.
//
// Solidity: event NewTransfer(bytes32 tradeHash, address poster, address asset, uint256 amount, address receiver, uint256 maxPenalty, uint256 deadline)
func (_OTCSwap *OTCSwapFilterer) ParseNewTransfer(log types.Log) (*OTCSwapNewTransfer, error) {
	event := new(OTCSwapNewTransfer)
	if err := _OTCSwap.contract.UnpackLog(event, "NewTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OTCSwapWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the OTCSwap contract.
type OTCSwapWithdrawnIterator struct {
	Event *OTCSwapWithdrawn // Event containing the contract specifics and raw log

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
func (it *OTCSwapWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OTCSwapWithdrawn)
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
		it.Event = new(OTCSwapWithdrawn)
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
func (it *OTCSwapWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OTCSwapWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OTCSwapWithdrawn represents a Withdrawn event raised by the OTCSwap contract.
type OTCSwapWithdrawn struct {
	TradeHash [32]byte
	Token     common.Address
	Amount    *big.Int
	Receiver  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xc279171b497ca67feac5cc18e17598de83f75462afc3610b03c7433b89202d6f.
//
// Solidity: event Withdrawn(bytes32 tradeHash, address token, uint256 amount, address receiver)
func (_OTCSwap *OTCSwapFilterer) FilterWithdrawn(opts *bind.FilterOpts) (*OTCSwapWithdrawnIterator, error) {

	logs, sub, err := _OTCSwap.contract.FilterLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return &OTCSwapWithdrawnIterator{contract: _OTCSwap.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xc279171b497ca67feac5cc18e17598de83f75462afc3610b03c7433b89202d6f.
//
// Solidity: event Withdrawn(bytes32 tradeHash, address token, uint256 amount, address receiver)
func (_OTCSwap *OTCSwapFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *OTCSwapWithdrawn) (event.Subscription, error) {

	logs, sub, err := _OTCSwap.contract.WatchLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OTCSwapWithdrawn)
				if err := _OTCSwap.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0xc279171b497ca67feac5cc18e17598de83f75462afc3610b03c7433b89202d6f.
//
// Solidity: event Withdrawn(bytes32 tradeHash, address token, uint256 amount, address receiver)
func (_OTCSwap *OTCSwapFilterer) ParseWithdrawn(log types.Log) (*OTCSwapWithdrawn, error) {
	event := new(OTCSwapWithdrawn)
	if err := _OTCSwap.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
