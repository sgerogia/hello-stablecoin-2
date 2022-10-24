// ConsentCode generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// ProvableGBPMetaData contains all meta data concerning the ProvableGBP contract.
var ProvableGBPMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_publicKey\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"grantEncryptedData\",\"type\":\"bytes\"}],\"name\":\"AuthGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"authEncryptedData\",\"type\":\"bytes\"}],\"name\":\"AuthRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encryptedData\",\"type\":\"bytes\"}],\"name\":\"MintRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"actualMintedPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"encryptedData\",\"type\":\"bytes\"}],\"name\":\"authGranted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"serverEncryptedData\",\"type\":\"bytes\"}],\"name\":\"authRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpiryTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encryptedData\",\"type\":\"bytes\"}],\"name\":\"mintRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oneHundredPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"paymentComplete\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seignorageFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_publicKey\",\"type\":\"bytes\"}],\"name\":\"setPublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ProvableGBPABI is the input ABI used to generate the binding from.
// Deprecated: Use ProvableGBPMetaData.ABI instead.
var ProvableGBPABI = ProvableGBPMetaData.ABI

// ProvableGBP is an auto generated Go binding around an Ethereum contract.
type ProvableGBP struct {
	ProvableGBPCaller     // Read-only binding to the contract
	ProvableGBPTransactor // Write-only binding to the contract
	ProvableGBPFilterer   // Log filterer for contract events
}

// ProvableGBPCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProvableGBPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProvableGBPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProvableGBPTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProvableGBPFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProvableGBPFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProvableGBPSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProvableGBPSession struct {
	Contract     *ProvableGBP      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProvableGBPCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProvableGBPCallerSession struct {
	Contract *ProvableGBPCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ProvableGBPTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProvableGBPTransactorSession struct {
	Contract     *ProvableGBPTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ProvableGBPRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProvableGBPRaw struct {
	Contract *ProvableGBP // Generic contract binding to access the raw methods on
}

// ProvableGBPCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProvableGBPCallerRaw struct {
	Contract *ProvableGBPCaller // Generic read-only contract binding to access the raw methods on
}

// ProvableGBPTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProvableGBPTransactorRaw struct {
	Contract *ProvableGBPTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProvableGBP creates a new instance of ProvableGBP, bound to a specific deployed contract.
func NewProvableGBP(address common.Address, backend bind.ContractBackend) (*ProvableGBP, error) {
	contract, err := bindProvableGBP(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProvableGBP{ProvableGBPCaller: ProvableGBPCaller{contract: contract}, ProvableGBPTransactor: ProvableGBPTransactor{contract: contract}, ProvableGBPFilterer: ProvableGBPFilterer{contract: contract}}, nil
}

// NewProvableGBPCaller creates a new read-only instance of ProvableGBP, bound to a specific deployed contract.
func NewProvableGBPCaller(address common.Address, caller bind.ContractCaller) (*ProvableGBPCaller, error) {
	contract, err := bindProvableGBP(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProvableGBPCaller{contract: contract}, nil
}

// NewProvableGBPTransactor creates a new write-only instance of ProvableGBP, bound to a specific deployed contract.
func NewProvableGBPTransactor(address common.Address, transactor bind.ContractTransactor) (*ProvableGBPTransactor, error) {
	contract, err := bindProvableGBP(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProvableGBPTransactor{contract: contract}, nil
}

// NewProvableGBPFilterer creates a new log filterer instance of ProvableGBP, bound to a specific deployed contract.
func NewProvableGBPFilterer(address common.Address, filterer bind.ContractFilterer) (*ProvableGBPFilterer, error) {
	contract, err := bindProvableGBP(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProvableGBPFilterer{contract: contract}, nil
}

// bindProvableGBP binds a generic wrapper to an already deployed contract.
func bindProvableGBP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProvableGBPABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProvableGBP *ProvableGBPRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProvableGBP.Contract.ProvableGBPCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProvableGBP *ProvableGBPRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProvableGBP.Contract.ProvableGBPTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProvableGBP *ProvableGBPRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProvableGBP.Contract.ProvableGBPTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProvableGBP *ProvableGBPCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProvableGBP.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProvableGBP *ProvableGBPTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProvableGBP.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProvableGBP *ProvableGBPTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProvableGBP.Contract.contract.Transact(opts, method, params...)
}

// ActualMintedPercentage is a free data retrieval call binding the contract method 0x3addb12c.
//
// Solidity: function actualMintedPercentage() view returns(uint256)
func (_ProvableGBP *ProvableGBPCaller) ActualMintedPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProvableGBP.contract.Call(opts, &out, "actualMintedPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActualMintedPercentage is a free data retrieval call binding the contract method 0x3addb12c.
//
// Solidity: function actualMintedPercentage() view returns(uint256)
func (_ProvableGBP *ProvableGBPSession) ActualMintedPercentage() (*big.Int, error) {
	return _ProvableGBP.Contract.ActualMintedPercentage(&_ProvableGBP.CallOpts)
}

// ActualMintedPercentage is a free data retrieval call binding the contract method 0x3addb12c.
//
// Solidity: function actualMintedPercentage() view returns(uint256)
func (_ProvableGBP *ProvableGBPCallerSession) ActualMintedPercentage() (*big.Int, error) {
	return _ProvableGBP.Contract.ActualMintedPercentage(&_ProvableGBP.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ProvableGBP *ProvableGBPCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ProvableGBP.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ProvableGBP *ProvableGBPSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ProvableGBP.Contract.Allowance(&_ProvableGBP.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ProvableGBP *ProvableGBPCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ProvableGBP.Contract.Allowance(&_ProvableGBP.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ProvableGBP *ProvableGBPCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ProvableGBP.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ProvableGBP *ProvableGBPSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ProvableGBP.Contract.BalanceOf(&_ProvableGBP.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ProvableGBP *ProvableGBPCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ProvableGBP.Contract.BalanceOf(&_ProvableGBP.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ProvableGBP *ProvableGBPCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ProvableGBP.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ProvableGBP *ProvableGBPSession) Decimals() (uint8, error) {
	return _ProvableGBP.Contract.Decimals(&_ProvableGBP.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ProvableGBP *ProvableGBPCallerSession) Decimals() (uint8, error) {
	return _ProvableGBP.Contract.Decimals(&_ProvableGBP.CallOpts)
}

// GetExpiryTime is a free data retrieval call binding the contract method 0x25cb5bc0.
//
// Solidity: function getExpiryTime() view returns(uint256)
func (_ProvableGBP *ProvableGBPCaller) GetExpiryTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProvableGBP.contract.Call(opts, &out, "getExpiryTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExpiryTime is a free data retrieval call binding the contract method 0x25cb5bc0.
//
// Solidity: function getExpiryTime() view returns(uint256)
func (_ProvableGBP *ProvableGBPSession) GetExpiryTime() (*big.Int, error) {
	return _ProvableGBP.Contract.GetExpiryTime(&_ProvableGBP.CallOpts)
}

// GetExpiryTime is a free data retrieval call binding the contract method 0x25cb5bc0.
//
// Solidity: function getExpiryTime() view returns(uint256)
func (_ProvableGBP *ProvableGBPCallerSession) GetExpiryTime() (*big.Int, error) {
	return _ProvableGBP.Contract.GetExpiryTime(&_ProvableGBP.CallOpts)
}

// Mint is a free data retrieval call binding the contract method 0x40c10f19.
//
// Solidity: function mint(address , uint256 ) view returns()
func (_ProvableGBP *ProvableGBPCaller) Mint(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) error {
	var out []interface{}
	err := _ProvableGBP.contract.Call(opts, &out, "mint", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// Mint is a free data retrieval call binding the contract method 0x40c10f19.
//
// Solidity: function mint(address , uint256 ) view returns()
func (_ProvableGBP *ProvableGBPSession) Mint(arg0 common.Address, arg1 *big.Int) error {
	return _ProvableGBP.Contract.Mint(&_ProvableGBP.CallOpts, arg0, arg1)
}

// Mint is a free data retrieval call binding the contract method 0x40c10f19.
//
// Solidity: function mint(address , uint256 ) view returns()
func (_ProvableGBP *ProvableGBPCallerSession) Mint(arg0 common.Address, arg1 *big.Int) error {
	return _ProvableGBP.Contract.Mint(&_ProvableGBP.CallOpts, arg0, arg1)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ProvableGBP *ProvableGBPCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ProvableGBP.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ProvableGBP *ProvableGBPSession) Name() (string, error) {
	return _ProvableGBP.Contract.Name(&_ProvableGBP.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ProvableGBP *ProvableGBPCallerSession) Name() (string, error) {
	return _ProvableGBP.Contract.Name(&_ProvableGBP.CallOpts)
}

// OneHundredPercent is a free data retrieval call binding the contract method 0x34913f28.
//
// Solidity: function oneHundredPercent() view returns(uint256)
func (_ProvableGBP *ProvableGBPCaller) OneHundredPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProvableGBP.contract.Call(opts, &out, "oneHundredPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OneHundredPercent is a free data retrieval call binding the contract method 0x34913f28.
//
// Solidity: function oneHundredPercent() view returns(uint256)
func (_ProvableGBP *ProvableGBPSession) OneHundredPercent() (*big.Int, error) {
	return _ProvableGBP.Contract.OneHundredPercent(&_ProvableGBP.CallOpts)
}

// OneHundredPercent is a free data retrieval call binding the contract method 0x34913f28.
//
// Solidity: function oneHundredPercent() view returns(uint256)
func (_ProvableGBP *ProvableGBPCallerSession) OneHundredPercent() (*big.Int, error) {
	return _ProvableGBP.Contract.OneHundredPercent(&_ProvableGBP.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProvableGBP *ProvableGBPCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProvableGBP.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProvableGBP *ProvableGBPSession) Owner() (common.Address, error) {
	return _ProvableGBP.Contract.Owner(&_ProvableGBP.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProvableGBP *ProvableGBPCallerSession) Owner() (common.Address, error) {
	return _ProvableGBP.Contract.Owner(&_ProvableGBP.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ProvableGBP *ProvableGBPCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ProvableGBP.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ProvableGBP *ProvableGBPSession) Paused() (bool, error) {
	return _ProvableGBP.Contract.Paused(&_ProvableGBP.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ProvableGBP *ProvableGBPCallerSession) Paused() (bool, error) {
	return _ProvableGBP.Contract.Paused(&_ProvableGBP.CallOpts)
}

// PublicKey is a free data retrieval call binding the contract method 0x63ffab31.
//
// Solidity: function publicKey() view returns(bytes)
func (_ProvableGBP *ProvableGBPCaller) PublicKey(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _ProvableGBP.contract.Call(opts, &out, "publicKey")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PublicKey is a free data retrieval call binding the contract method 0x63ffab31.
//
// Solidity: function publicKey() view returns(bytes)
func (_ProvableGBP *ProvableGBPSession) PublicKey() ([]byte, error) {
	return _ProvableGBP.Contract.PublicKey(&_ProvableGBP.CallOpts)
}

// PublicKey is a free data retrieval call binding the contract method 0x63ffab31.
//
// Solidity: function publicKey() view returns(bytes)
func (_ProvableGBP *ProvableGBPCallerSession) PublicKey() ([]byte, error) {
	return _ProvableGBP.Contract.PublicKey(&_ProvableGBP.CallOpts)
}

// SeignorageFee is a free data retrieval call binding the contract method 0x1285cc2a.
//
// Solidity: function seignorageFee() view returns(uint256)
func (_ProvableGBP *ProvableGBPCaller) SeignorageFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProvableGBP.contract.Call(opts, &out, "seignorageFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SeignorageFee is a free data retrieval call binding the contract method 0x1285cc2a.
//
// Solidity: function seignorageFee() view returns(uint256)
func (_ProvableGBP *ProvableGBPSession) SeignorageFee() (*big.Int, error) {
	return _ProvableGBP.Contract.SeignorageFee(&_ProvableGBP.CallOpts)
}

// SeignorageFee is a free data retrieval call binding the contract method 0x1285cc2a.
//
// Solidity: function seignorageFee() view returns(uint256)
func (_ProvableGBP *ProvableGBPCallerSession) SeignorageFee() (*big.Int, error) {
	return _ProvableGBP.Contract.SeignorageFee(&_ProvableGBP.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ProvableGBP *ProvableGBPCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ProvableGBP.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ProvableGBP *ProvableGBPSession) Symbol() (string, error) {
	return _ProvableGBP.Contract.Symbol(&_ProvableGBP.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ProvableGBP *ProvableGBPCallerSession) Symbol() (string, error) {
	return _ProvableGBP.Contract.Symbol(&_ProvableGBP.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ProvableGBP *ProvableGBPCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProvableGBP.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ProvableGBP *ProvableGBPSession) TotalSupply() (*big.Int, error) {
	return _ProvableGBP.Contract.TotalSupply(&_ProvableGBP.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ProvableGBP *ProvableGBPCallerSession) TotalSupply() (*big.Int, error) {
	return _ProvableGBP.Contract.TotalSupply(&_ProvableGBP.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ProvableGBP *ProvableGBPTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ProvableGBP.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ProvableGBP *ProvableGBPSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ProvableGBP.Contract.Approve(&_ProvableGBP.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ProvableGBP *ProvableGBPTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ProvableGBP.Contract.Approve(&_ProvableGBP.TransactOpts, spender, amount)
}

// AuthGranted is a paid mutator transaction binding the contract method 0x21a7e7e9.
//
// Solidity: function authGranted(bytes32 requestId, bytes encryptedData) returns()
func (_ProvableGBP *ProvableGBPTransactor) AuthGranted(opts *bind.TransactOpts, requestId [32]byte, encryptedData []byte) (*types.Transaction, error) {
	return _ProvableGBP.contract.Transact(opts, "authGranted", requestId, encryptedData)
}

// AuthGranted is a paid mutator transaction binding the contract method 0x21a7e7e9.
//
// Solidity: function authGranted(bytes32 requestId, bytes encryptedData) returns()
func (_ProvableGBP *ProvableGBPSession) AuthGranted(requestId [32]byte, encryptedData []byte) (*types.Transaction, error) {
	return _ProvableGBP.Contract.AuthGranted(&_ProvableGBP.TransactOpts, requestId, encryptedData)
}

// AuthGranted is a paid mutator transaction binding the contract method 0x21a7e7e9.
//
// Solidity: function authGranted(bytes32 requestId, bytes encryptedData) returns()
func (_ProvableGBP *ProvableGBPTransactorSession) AuthGranted(requestId [32]byte, encryptedData []byte) (*types.Transaction, error) {
	return _ProvableGBP.Contract.AuthGranted(&_ProvableGBP.TransactOpts, requestId, encryptedData)
}

// AuthRequest is a paid mutator transaction binding the contract method 0xa4629533.
//
// Solidity: function authRequest(bytes32 requestId, bytes serverEncryptedData) returns()
func (_ProvableGBP *ProvableGBPTransactor) AuthRequest(opts *bind.TransactOpts, requestId [32]byte, serverEncryptedData []byte) (*types.Transaction, error) {
	return _ProvableGBP.contract.Transact(opts, "authRequest", requestId, serverEncryptedData)
}

// AuthRequest is a paid mutator transaction binding the contract method 0xa4629533.
//
// Solidity: function authRequest(bytes32 requestId, bytes serverEncryptedData) returns()
func (_ProvableGBP *ProvableGBPSession) AuthRequest(requestId [32]byte, serverEncryptedData []byte) (*types.Transaction, error) {
	return _ProvableGBP.Contract.AuthRequest(&_ProvableGBP.TransactOpts, requestId, serverEncryptedData)
}

// AuthRequest is a paid mutator transaction binding the contract method 0xa4629533.
//
// Solidity: function authRequest(bytes32 requestId, bytes serverEncryptedData) returns()
func (_ProvableGBP *ProvableGBPTransactorSession) AuthRequest(requestId [32]byte, serverEncryptedData []byte) (*types.Transaction, error) {
	return _ProvableGBP.Contract.AuthRequest(&_ProvableGBP.TransactOpts, requestId, serverEncryptedData)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ProvableGBP *ProvableGBPTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ProvableGBP.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ProvableGBP *ProvableGBPSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ProvableGBP.Contract.Burn(&_ProvableGBP.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ProvableGBP *ProvableGBPTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ProvableGBP.Contract.Burn(&_ProvableGBP.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ProvableGBP *ProvableGBPTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ProvableGBP.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ProvableGBP *ProvableGBPSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ProvableGBP.Contract.BurnFrom(&_ProvableGBP.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ProvableGBP *ProvableGBPTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ProvableGBP.Contract.BurnFrom(&_ProvableGBP.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ProvableGBP *ProvableGBPTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ProvableGBP.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ProvableGBP *ProvableGBPSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ProvableGBP.Contract.DecreaseAllowance(&_ProvableGBP.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ProvableGBP *ProvableGBPTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ProvableGBP.Contract.DecreaseAllowance(&_ProvableGBP.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ProvableGBP *ProvableGBPTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ProvableGBP.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ProvableGBP *ProvableGBPSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ProvableGBP.Contract.IncreaseAllowance(&_ProvableGBP.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ProvableGBP *ProvableGBPTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ProvableGBP.Contract.IncreaseAllowance(&_ProvableGBP.TransactOpts, spender, addedValue)
}

// MintRequest is a paid mutator transaction binding the contract method 0xcdc1f6d4.
//
// Solidity: function mintRequest(uint256 amount, bytes encryptedData) returns()
func (_ProvableGBP *ProvableGBPTransactor) MintRequest(opts *bind.TransactOpts, amount *big.Int, encryptedData []byte) (*types.Transaction, error) {
	return _ProvableGBP.contract.Transact(opts, "mintRequest", amount, encryptedData)
}

// MintRequest is a paid mutator transaction binding the contract method 0xcdc1f6d4.
//
// Solidity: function mintRequest(uint256 amount, bytes encryptedData) returns()
func (_ProvableGBP *ProvableGBPSession) MintRequest(amount *big.Int, encryptedData []byte) (*types.Transaction, error) {
	return _ProvableGBP.Contract.MintRequest(&_ProvableGBP.TransactOpts, amount, encryptedData)
}

// MintRequest is a paid mutator transaction binding the contract method 0xcdc1f6d4.
//
// Solidity: function mintRequest(uint256 amount, bytes encryptedData) returns()
func (_ProvableGBP *ProvableGBPTransactorSession) MintRequest(amount *big.Int, encryptedData []byte) (*types.Transaction, error) {
	return _ProvableGBP.Contract.MintRequest(&_ProvableGBP.TransactOpts, amount, encryptedData)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ProvableGBP *ProvableGBPTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProvableGBP.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ProvableGBP *ProvableGBPSession) Pause() (*types.Transaction, error) {
	return _ProvableGBP.Contract.Pause(&_ProvableGBP.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ProvableGBP *ProvableGBPTransactorSession) Pause() (*types.Transaction, error) {
	return _ProvableGBP.Contract.Pause(&_ProvableGBP.TransactOpts)
}

// PaymentComplete is a paid mutator transaction binding the contract method 0xac5b756d.
//
// Solidity: function paymentComplete(bytes32 requestId) returns()
func (_ProvableGBP *ProvableGBPTransactor) PaymentComplete(opts *bind.TransactOpts, requestId [32]byte) (*types.Transaction, error) {
	return _ProvableGBP.contract.Transact(opts, "paymentComplete", requestId)
}

// PaymentComplete is a paid mutator transaction binding the contract method 0xac5b756d.
//
// Solidity: function paymentComplete(bytes32 requestId) returns()
func (_ProvableGBP *ProvableGBPSession) PaymentComplete(requestId [32]byte) (*types.Transaction, error) {
	return _ProvableGBP.Contract.PaymentComplete(&_ProvableGBP.TransactOpts, requestId)
}

// PaymentComplete is a paid mutator transaction binding the contract method 0xac5b756d.
//
// Solidity: function paymentComplete(bytes32 requestId) returns()
func (_ProvableGBP *ProvableGBPTransactorSession) PaymentComplete(requestId [32]byte) (*types.Transaction, error) {
	return _ProvableGBP.Contract.PaymentComplete(&_ProvableGBP.TransactOpts, requestId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProvableGBP *ProvableGBPTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProvableGBP.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProvableGBP *ProvableGBPSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProvableGBP.Contract.RenounceOwnership(&_ProvableGBP.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProvableGBP *ProvableGBPTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProvableGBP.Contract.RenounceOwnership(&_ProvableGBP.TransactOpts)
}

// SetPublicKey is a paid mutator transaction binding the contract method 0xa91d58b4.
//
// Solidity: function setPublicKey(bytes _publicKey) returns()
func (_ProvableGBP *ProvableGBPTransactor) SetPublicKey(opts *bind.TransactOpts, _publicKey []byte) (*types.Transaction, error) {
	return _ProvableGBP.contract.Transact(opts, "setPublicKey", _publicKey)
}

// SetPublicKey is a paid mutator transaction binding the contract method 0xa91d58b4.
//
// Solidity: function setPublicKey(bytes _publicKey) returns()
func (_ProvableGBP *ProvableGBPSession) SetPublicKey(_publicKey []byte) (*types.Transaction, error) {
	return _ProvableGBP.Contract.SetPublicKey(&_ProvableGBP.TransactOpts, _publicKey)
}

// SetPublicKey is a paid mutator transaction binding the contract method 0xa91d58b4.
//
// Solidity: function setPublicKey(bytes _publicKey) returns()
func (_ProvableGBP *ProvableGBPTransactorSession) SetPublicKey(_publicKey []byte) (*types.Transaction, error) {
	return _ProvableGBP.Contract.SetPublicKey(&_ProvableGBP.TransactOpts, _publicKey)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ProvableGBP *ProvableGBPTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ProvableGBP.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ProvableGBP *ProvableGBPSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ProvableGBP.Contract.Transfer(&_ProvableGBP.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ProvableGBP *ProvableGBPTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ProvableGBP.Contract.Transfer(&_ProvableGBP.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ProvableGBP *ProvableGBPTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ProvableGBP.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ProvableGBP *ProvableGBPSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ProvableGBP.Contract.TransferFrom(&_ProvableGBP.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ProvableGBP *ProvableGBPTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ProvableGBP.Contract.TransferFrom(&_ProvableGBP.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProvableGBP *ProvableGBPTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ProvableGBP.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProvableGBP *ProvableGBPSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProvableGBP.Contract.TransferOwnership(&_ProvableGBP.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProvableGBP *ProvableGBPTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProvableGBP.Contract.TransferOwnership(&_ProvableGBP.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ProvableGBP *ProvableGBPTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProvableGBP.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ProvableGBP *ProvableGBPSession) Unpause() (*types.Transaction, error) {
	return _ProvableGBP.Contract.Unpause(&_ProvableGBP.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ProvableGBP *ProvableGBPTransactorSession) Unpause() (*types.Transaction, error) {
	return _ProvableGBP.Contract.Unpause(&_ProvableGBP.TransactOpts)
}

// ProvableGBPApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ProvableGBP contract.
type ProvableGBPApprovalIterator struct {
	Event *ProvableGBPApproval // Event containing the contract specifics and raw log

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
func (it *ProvableGBPApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProvableGBPApproval)
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
		it.Event = new(ProvableGBPApproval)
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
func (it *ProvableGBPApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProvableGBPApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProvableGBPApproval represents a Approval event raised by the ProvableGBP contract.
type ProvableGBPApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ProvableGBP *ProvableGBPFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ProvableGBPApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ProvableGBP.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ProvableGBPApprovalIterator{contract: _ProvableGBP.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ProvableGBP *ProvableGBPFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ProvableGBPApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ProvableGBP.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProvableGBPApproval)
				if err := _ProvableGBP.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ProvableGBP *ProvableGBPFilterer) ParseApproval(log types.Log) (*ProvableGBPApproval, error) {
	event := new(ProvableGBPApproval)
	if err := _ProvableGBP.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProvableGBPAuthGrantedIterator is returned from FilterAuthGranted and is used to iterate over the raw logs and unpacked data for AuthGranted events raised by the ProvableGBP contract.
type ProvableGBPAuthGrantedIterator struct {
	Event *ProvableGBPAuthGranted // Event containing the contract specifics and raw log

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
func (it *ProvableGBPAuthGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProvableGBPAuthGranted)
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
		it.Event = new(ProvableGBPAuthGranted)
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
func (it *ProvableGBPAuthGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProvableGBPAuthGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProvableGBPAuthGranted represents a AuthGranted event raised by the ProvableGBP contract.
type ProvableGBPAuthGranted struct {
	Requester          common.Address
	RequestId          [32]byte
	GrantEncryptedData []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAuthGranted is a free log retrieval operation binding the contract event 0xa3cd7d021dc9794ae8a4520e68440669e45a8758c6feac5a46db4c6953d35d25.
//
// Solidity: event AuthGranted(address indexed requester, bytes32 indexed requestId, bytes grantEncryptedData)
func (_ProvableGBP *ProvableGBPFilterer) FilterAuthGranted(opts *bind.FilterOpts, requester []common.Address, requestId [][32]byte) (*ProvableGBPAuthGrantedIterator, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _ProvableGBP.contract.FilterLogs(opts, "AuthGranted", requesterRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ProvableGBPAuthGrantedIterator{contract: _ProvableGBP.contract, event: "AuthGranted", logs: logs, sub: sub}, nil
}

// WatchAuthGranted is a free log subscription operation binding the contract event 0xa3cd7d021dc9794ae8a4520e68440669e45a8758c6feac5a46db4c6953d35d25.
//
// Solidity: event AuthGranted(address indexed requester, bytes32 indexed requestId, bytes grantEncryptedData)
func (_ProvableGBP *ProvableGBPFilterer) WatchAuthGranted(opts *bind.WatchOpts, sink chan<- *ProvableGBPAuthGranted, requester []common.Address, requestId [][32]byte) (event.Subscription, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _ProvableGBP.contract.WatchLogs(opts, "AuthGranted", requesterRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProvableGBPAuthGranted)
				if err := _ProvableGBP.contract.UnpackLog(event, "AuthGranted", log); err != nil {
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

// ParseAuthGranted is a log parse operation binding the contract event 0xa3cd7d021dc9794ae8a4520e68440669e45a8758c6feac5a46db4c6953d35d25.
//
// Solidity: event AuthGranted(address indexed requester, bytes32 indexed requestId, bytes grantEncryptedData)
func (_ProvableGBP *ProvableGBPFilterer) ParseAuthGranted(log types.Log) (*ProvableGBPAuthGranted, error) {
	event := new(ProvableGBPAuthGranted)
	if err := _ProvableGBP.contract.UnpackLog(event, "AuthGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProvableGBPAuthRequestIterator is returned from FilterAuthRequest and is used to iterate over the raw logs and unpacked data for AuthRequest events raised by the ProvableGBP contract.
type ProvableGBPAuthRequestIterator struct {
	Event *ProvableGBPAuthRequest // Event containing the contract specifics and raw log

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
func (it *ProvableGBPAuthRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProvableGBPAuthRequest)
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
		it.Event = new(ProvableGBPAuthRequest)
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
func (it *ProvableGBPAuthRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProvableGBPAuthRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProvableGBPAuthRequest represents a AuthRequest event raised by the ProvableGBP contract.
type ProvableGBPAuthRequest struct {
	Requester         common.Address
	RequestId         [32]byte
	AuthEncryptedData []byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterAuthRequest is a free log retrieval operation binding the contract event 0x220f931a071dfcd35fda9f25680cb03785508840415a802de6e99230f889862b.
//
// Solidity: event AuthRequest(address indexed requester, bytes32 indexed requestId, bytes authEncryptedData)
func (_ProvableGBP *ProvableGBPFilterer) FilterAuthRequest(opts *bind.FilterOpts, requester []common.Address, requestId [][32]byte) (*ProvableGBPAuthRequestIterator, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _ProvableGBP.contract.FilterLogs(opts, "AuthRequest", requesterRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ProvableGBPAuthRequestIterator{contract: _ProvableGBP.contract, event: "AuthRequest", logs: logs, sub: sub}, nil
}

// WatchAuthRequest is a free log subscription operation binding the contract event 0x220f931a071dfcd35fda9f25680cb03785508840415a802de6e99230f889862b.
//
// Solidity: event AuthRequest(address indexed requester, bytes32 indexed requestId, bytes authEncryptedData)
func (_ProvableGBP *ProvableGBPFilterer) WatchAuthRequest(opts *bind.WatchOpts, sink chan<- *ProvableGBPAuthRequest, requester []common.Address, requestId [][32]byte) (event.Subscription, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _ProvableGBP.contract.WatchLogs(opts, "AuthRequest", requesterRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProvableGBPAuthRequest)
				if err := _ProvableGBP.contract.UnpackLog(event, "AuthRequest", log); err != nil {
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

// ParseAuthRequest is a log parse operation binding the contract event 0x220f931a071dfcd35fda9f25680cb03785508840415a802de6e99230f889862b.
//
// Solidity: event AuthRequest(address indexed requester, bytes32 indexed requestId, bytes authEncryptedData)
func (_ProvableGBP *ProvableGBPFilterer) ParseAuthRequest(log types.Log) (*ProvableGBPAuthRequest, error) {
	event := new(ProvableGBPAuthRequest)
	if err := _ProvableGBP.contract.UnpackLog(event, "AuthRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProvableGBPMintRequestIterator is returned from FilterMintRequest and is used to iterate over the raw logs and unpacked data for MintRequest events raised by the ProvableGBP contract.
type ProvableGBPMintRequestIterator struct {
	Event *ProvableGBPMintRequest // Event containing the contract specifics and raw log

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
func (it *ProvableGBPMintRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProvableGBPMintRequest)
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
		it.Event = new(ProvableGBPMintRequest)
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
func (it *ProvableGBPMintRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProvableGBPMintRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProvableGBPMintRequest represents a MintRequest event raised by the ProvableGBP contract.
type ProvableGBPMintRequest struct {
	Requester     common.Address
	RequestId     [32]byte
	Amount        *big.Int
	Expiration    *big.Int
	EncryptedData []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMintRequest is a free log retrieval operation binding the contract event 0x45fa4b26a755f4cc5780432570badb0410d1ed0c688479aa9e708761ffb82ec2.
//
// Solidity: event MintRequest(address indexed requester, bytes32 indexed requestId, uint256 amount, uint256 expiration, bytes encryptedData)
func (_ProvableGBP *ProvableGBPFilterer) FilterMintRequest(opts *bind.FilterOpts, requester []common.Address, requestId [][32]byte) (*ProvableGBPMintRequestIterator, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _ProvableGBP.contract.FilterLogs(opts, "MintRequest", requesterRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ProvableGBPMintRequestIterator{contract: _ProvableGBP.contract, event: "MintRequest", logs: logs, sub: sub}, nil
}

// WatchMintRequest is a free log subscription operation binding the contract event 0x45fa4b26a755f4cc5780432570badb0410d1ed0c688479aa9e708761ffb82ec2.
//
// Solidity: event MintRequest(address indexed requester, bytes32 indexed requestId, uint256 amount, uint256 expiration, bytes encryptedData)
func (_ProvableGBP *ProvableGBPFilterer) WatchMintRequest(opts *bind.WatchOpts, sink chan<- *ProvableGBPMintRequest, requester []common.Address, requestId [][32]byte) (event.Subscription, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _ProvableGBP.contract.WatchLogs(opts, "MintRequest", requesterRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProvableGBPMintRequest)
				if err := _ProvableGBP.contract.UnpackLog(event, "MintRequest", log); err != nil {
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

// ParseMintRequest is a log parse operation binding the contract event 0x45fa4b26a755f4cc5780432570badb0410d1ed0c688479aa9e708761ffb82ec2.
//
// Solidity: event MintRequest(address indexed requester, bytes32 indexed requestId, uint256 amount, uint256 expiration, bytes encryptedData)
func (_ProvableGBP *ProvableGBPFilterer) ParseMintRequest(log types.Log) (*ProvableGBPMintRequest, error) {
	event := new(ProvableGBPMintRequest)
	if err := _ProvableGBP.contract.UnpackLog(event, "MintRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProvableGBPOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ProvableGBP contract.
type ProvableGBPOwnershipTransferredIterator struct {
	Event *ProvableGBPOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ProvableGBPOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProvableGBPOwnershipTransferred)
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
		it.Event = new(ProvableGBPOwnershipTransferred)
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
func (it *ProvableGBPOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProvableGBPOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProvableGBPOwnershipTransferred represents a OwnershipTransferred event raised by the ProvableGBP contract.
type ProvableGBPOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProvableGBP *ProvableGBPFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProvableGBPOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProvableGBP.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProvableGBPOwnershipTransferredIterator{contract: _ProvableGBP.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProvableGBP *ProvableGBPFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProvableGBPOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProvableGBP.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProvableGBPOwnershipTransferred)
				if err := _ProvableGBP.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProvableGBP *ProvableGBPFilterer) ParseOwnershipTransferred(log types.Log) (*ProvableGBPOwnershipTransferred, error) {
	event := new(ProvableGBPOwnershipTransferred)
	if err := _ProvableGBP.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProvableGBPPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ProvableGBP contract.
type ProvableGBPPausedIterator struct {
	Event *ProvableGBPPaused // Event containing the contract specifics and raw log

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
func (it *ProvableGBPPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProvableGBPPaused)
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
		it.Event = new(ProvableGBPPaused)
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
func (it *ProvableGBPPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProvableGBPPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProvableGBPPaused represents a Paused event raised by the ProvableGBP contract.
type ProvableGBPPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ProvableGBP *ProvableGBPFilterer) FilterPaused(opts *bind.FilterOpts) (*ProvableGBPPausedIterator, error) {

	logs, sub, err := _ProvableGBP.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ProvableGBPPausedIterator{contract: _ProvableGBP.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ProvableGBP *ProvableGBPFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ProvableGBPPaused) (event.Subscription, error) {

	logs, sub, err := _ProvableGBP.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProvableGBPPaused)
				if err := _ProvableGBP.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ProvableGBP *ProvableGBPFilterer) ParsePaused(log types.Log) (*ProvableGBPPaused, error) {
	event := new(ProvableGBPPaused)
	if err := _ProvableGBP.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProvableGBPTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ProvableGBP contract.
type ProvableGBPTransferIterator struct {
	Event *ProvableGBPTransfer // Event containing the contract specifics and raw log

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
func (it *ProvableGBPTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProvableGBPTransfer)
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
		it.Event = new(ProvableGBPTransfer)
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
func (it *ProvableGBPTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProvableGBPTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProvableGBPTransfer represents a Transfer event raised by the ProvableGBP contract.
type ProvableGBPTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ProvableGBP *ProvableGBPFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ProvableGBPTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ProvableGBP.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ProvableGBPTransferIterator{contract: _ProvableGBP.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ProvableGBP *ProvableGBPFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ProvableGBPTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ProvableGBP.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProvableGBPTransfer)
				if err := _ProvableGBP.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ProvableGBP *ProvableGBPFilterer) ParseTransfer(log types.Log) (*ProvableGBPTransfer, error) {
	event := new(ProvableGBPTransfer)
	if err := _ProvableGBP.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProvableGBPUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ProvableGBP contract.
type ProvableGBPUnpausedIterator struct {
	Event *ProvableGBPUnpaused // Event containing the contract specifics and raw log

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
func (it *ProvableGBPUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProvableGBPUnpaused)
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
		it.Event = new(ProvableGBPUnpaused)
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
func (it *ProvableGBPUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProvableGBPUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProvableGBPUnpaused represents a Unpaused event raised by the ProvableGBP contract.
type ProvableGBPUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ProvableGBP *ProvableGBPFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ProvableGBPUnpausedIterator, error) {

	logs, sub, err := _ProvableGBP.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ProvableGBPUnpausedIterator{contract: _ProvableGBP.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ProvableGBP *ProvableGBPFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ProvableGBPUnpaused) (event.Subscription, error) {

	logs, sub, err := _ProvableGBP.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProvableGBPUnpaused)
				if err := _ProvableGBP.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ProvableGBP *ProvableGBPFilterer) ParseUnpaused(log types.Log) (*ProvableGBPUnpaused, error) {
	event := new(ProvableGBPUnpaused)
	if err := _ProvableGBP.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
