// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc1155Extensions

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

// Erc1155ExtensionsMetaData contains all meta data concerning the Erc1155Extensions contract.
var Erc1155ExtensionsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Erc1155ExtensionsABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc1155ExtensionsMetaData.ABI instead.
var Erc1155ExtensionsABI = Erc1155ExtensionsMetaData.ABI

// Erc1155Extensions is an auto generated Go binding around an Ethereum contract.
type Erc1155Extensions struct {
	Erc1155ExtensionsCaller     // Read-only binding to the contract
	Erc1155ExtensionsTransactor // Write-only binding to the contract
	Erc1155ExtensionsFilterer   // Log filterer for contract events
}

// Erc1155ExtensionsCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc1155ExtensionsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc1155ExtensionsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc1155ExtensionsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc1155ExtensionsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc1155ExtensionsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc1155ExtensionsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc1155ExtensionsSession struct {
	Contract     *Erc1155Extensions // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Erc1155ExtensionsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc1155ExtensionsCallerSession struct {
	Contract *Erc1155ExtensionsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// Erc1155ExtensionsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc1155ExtensionsTransactorSession struct {
	Contract     *Erc1155ExtensionsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// Erc1155ExtensionsRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc1155ExtensionsRaw struct {
	Contract *Erc1155Extensions // Generic contract binding to access the raw methods on
}

// Erc1155ExtensionsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc1155ExtensionsCallerRaw struct {
	Contract *Erc1155ExtensionsCaller // Generic read-only contract binding to access the raw methods on
}

// Erc1155ExtensionsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc1155ExtensionsTransactorRaw struct {
	Contract *Erc1155ExtensionsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc1155Extensions creates a new instance of Erc1155Extensions, bound to a specific deployed contract.
func NewErc1155Extensions(address common.Address, backend bind.ContractBackend) (*Erc1155Extensions, error) {
	contract, err := bindErc1155Extensions(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc1155Extensions{Erc1155ExtensionsCaller: Erc1155ExtensionsCaller{contract: contract}, Erc1155ExtensionsTransactor: Erc1155ExtensionsTransactor{contract: contract}, Erc1155ExtensionsFilterer: Erc1155ExtensionsFilterer{contract: contract}}, nil
}

// NewErc1155ExtensionsCaller creates a new read-only instance of Erc1155Extensions, bound to a specific deployed contract.
func NewErc1155ExtensionsCaller(address common.Address, caller bind.ContractCaller) (*Erc1155ExtensionsCaller, error) {
	contract, err := bindErc1155Extensions(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc1155ExtensionsCaller{contract: contract}, nil
}

// NewErc1155ExtensionsTransactor creates a new write-only instance of Erc1155Extensions, bound to a specific deployed contract.
func NewErc1155ExtensionsTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc1155ExtensionsTransactor, error) {
	contract, err := bindErc1155Extensions(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc1155ExtensionsTransactor{contract: contract}, nil
}

// NewErc1155ExtensionsFilterer creates a new log filterer instance of Erc1155Extensions, bound to a specific deployed contract.
func NewErc1155ExtensionsFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc1155ExtensionsFilterer, error) {
	contract, err := bindErc1155Extensions(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc1155ExtensionsFilterer{contract: contract}, nil
}

// bindErc1155Extensions binds a generic wrapper to an already deployed contract.
func bindErc1155Extensions(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc1155ExtensionsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc1155Extensions *Erc1155ExtensionsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc1155Extensions.Contract.Erc1155ExtensionsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc1155Extensions *Erc1155ExtensionsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc1155Extensions.Contract.Erc1155ExtensionsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc1155Extensions *Erc1155ExtensionsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc1155Extensions.Contract.Erc1155ExtensionsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc1155Extensions *Erc1155ExtensionsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc1155Extensions.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc1155Extensions *Erc1155ExtensionsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc1155Extensions.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc1155Extensions *Erc1155ExtensionsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc1155Extensions.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Erc1155Extensions *Erc1155ExtensionsCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Erc1155Extensions.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Erc1155Extensions *Erc1155ExtensionsSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Erc1155Extensions.Contract.BalanceOf(&_Erc1155Extensions.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Erc1155Extensions *Erc1155ExtensionsCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Erc1155Extensions.Contract.BalanceOf(&_Erc1155Extensions.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Erc1155Extensions *Erc1155ExtensionsCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Erc1155Extensions.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Erc1155Extensions *Erc1155ExtensionsSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Erc1155Extensions.Contract.BalanceOfBatch(&_Erc1155Extensions.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Erc1155Extensions *Erc1155ExtensionsCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Erc1155Extensions.Contract.BalanceOfBatch(&_Erc1155Extensions.CallOpts, accounts, ids)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Erc1155Extensions *Erc1155ExtensionsCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Erc1155Extensions.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Erc1155Extensions *Erc1155ExtensionsSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Erc1155Extensions.Contract.IsApprovedForAll(&_Erc1155Extensions.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Erc1155Extensions *Erc1155ExtensionsCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Erc1155Extensions.Contract.IsApprovedForAll(&_Erc1155Extensions.CallOpts, account, operator)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc1155Extensions *Erc1155ExtensionsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Erc1155Extensions.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc1155Extensions *Erc1155ExtensionsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Erc1155Extensions.Contract.SupportsInterface(&_Erc1155Extensions.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc1155Extensions *Erc1155ExtensionsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Erc1155Extensions.Contract.SupportsInterface(&_Erc1155Extensions.CallOpts, interfaceId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_Erc1155Extensions *Erc1155ExtensionsCaller) Uri(opts *bind.CallOpts, id *big.Int) (string, error) {
	var out []interface{}
	err := _Erc1155Extensions.contract.Call(opts, &out, "uri", id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_Erc1155Extensions *Erc1155ExtensionsSession) Uri(id *big.Int) (string, error) {
	return _Erc1155Extensions.Contract.Uri(&_Erc1155Extensions.CallOpts, id)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_Erc1155Extensions *Erc1155ExtensionsCallerSession) Uri(id *big.Int) (string, error) {
	return _Erc1155Extensions.Contract.Uri(&_Erc1155Extensions.CallOpts, id)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Erc1155Extensions *Erc1155ExtensionsTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Erc1155Extensions.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Erc1155Extensions *Erc1155ExtensionsSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Erc1155Extensions.Contract.SafeBatchTransferFrom(&_Erc1155Extensions.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Erc1155Extensions *Erc1155ExtensionsTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Erc1155Extensions.Contract.SafeBatchTransferFrom(&_Erc1155Extensions.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Erc1155Extensions *Erc1155ExtensionsTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc1155Extensions.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Erc1155Extensions *Erc1155ExtensionsSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc1155Extensions.Contract.SafeTransferFrom(&_Erc1155Extensions.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Erc1155Extensions *Erc1155ExtensionsTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc1155Extensions.Contract.SafeTransferFrom(&_Erc1155Extensions.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc1155Extensions *Erc1155ExtensionsTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Erc1155Extensions.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc1155Extensions *Erc1155ExtensionsSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Erc1155Extensions.Contract.SetApprovalForAll(&_Erc1155Extensions.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc1155Extensions *Erc1155ExtensionsTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Erc1155Extensions.Contract.SetApprovalForAll(&_Erc1155Extensions.TransactOpts, operator, approved)
}

// Erc1155ExtensionsApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Erc1155Extensions contract.
type Erc1155ExtensionsApprovalForAllIterator struct {
	Event *Erc1155ExtensionsApprovalForAll // Event containing the contract specifics and raw log

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
func (it *Erc1155ExtensionsApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155ExtensionsApprovalForAll)
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
		it.Event = new(Erc1155ExtensionsApprovalForAll)
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
func (it *Erc1155ExtensionsApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155ExtensionsApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155ExtensionsApprovalForAll represents a ApprovalForAll event raised by the Erc1155Extensions contract.
type Erc1155ExtensionsApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Erc1155Extensions *Erc1155ExtensionsFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*Erc1155ExtensionsApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Erc1155Extensions.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &Erc1155ExtensionsApprovalForAllIterator{contract: _Erc1155Extensions.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Erc1155Extensions *Erc1155ExtensionsFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Erc1155ExtensionsApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Erc1155Extensions.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155ExtensionsApprovalForAll)
				if err := _Erc1155Extensions.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Erc1155Extensions *Erc1155ExtensionsFilterer) ParseApprovalForAll(log types.Log) (*Erc1155ExtensionsApprovalForAll, error) {
	event := new(Erc1155ExtensionsApprovalForAll)
	if err := _Erc1155Extensions.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc1155ExtensionsTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the Erc1155Extensions contract.
type Erc1155ExtensionsTransferBatchIterator struct {
	Event *Erc1155ExtensionsTransferBatch // Event containing the contract specifics and raw log

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
func (it *Erc1155ExtensionsTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155ExtensionsTransferBatch)
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
		it.Event = new(Erc1155ExtensionsTransferBatch)
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
func (it *Erc1155ExtensionsTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155ExtensionsTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155ExtensionsTransferBatch represents a TransferBatch event raised by the Erc1155Extensions contract.
type Erc1155ExtensionsTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Erc1155Extensions *Erc1155ExtensionsFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*Erc1155ExtensionsTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc1155Extensions.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc1155ExtensionsTransferBatchIterator{contract: _Erc1155Extensions.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Erc1155Extensions *Erc1155ExtensionsFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *Erc1155ExtensionsTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc1155Extensions.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155ExtensionsTransferBatch)
				if err := _Erc1155Extensions.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Erc1155Extensions *Erc1155ExtensionsFilterer) ParseTransferBatch(log types.Log) (*Erc1155ExtensionsTransferBatch, error) {
	event := new(Erc1155ExtensionsTransferBatch)
	if err := _Erc1155Extensions.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc1155ExtensionsTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the Erc1155Extensions contract.
type Erc1155ExtensionsTransferSingleIterator struct {
	Event *Erc1155ExtensionsTransferSingle // Event containing the contract specifics and raw log

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
func (it *Erc1155ExtensionsTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155ExtensionsTransferSingle)
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
		it.Event = new(Erc1155ExtensionsTransferSingle)
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
func (it *Erc1155ExtensionsTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155ExtensionsTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155ExtensionsTransferSingle represents a TransferSingle event raised by the Erc1155Extensions contract.
type Erc1155ExtensionsTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Erc1155Extensions *Erc1155ExtensionsFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*Erc1155ExtensionsTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc1155Extensions.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc1155ExtensionsTransferSingleIterator{contract: _Erc1155Extensions.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Erc1155Extensions *Erc1155ExtensionsFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *Erc1155ExtensionsTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc1155Extensions.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155ExtensionsTransferSingle)
				if err := _Erc1155Extensions.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Erc1155Extensions *Erc1155ExtensionsFilterer) ParseTransferSingle(log types.Log) (*Erc1155ExtensionsTransferSingle, error) {
	event := new(Erc1155ExtensionsTransferSingle)
	if err := _Erc1155Extensions.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc1155ExtensionsURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the Erc1155Extensions contract.
type Erc1155ExtensionsURIIterator struct {
	Event *Erc1155ExtensionsURI // Event containing the contract specifics and raw log

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
func (it *Erc1155ExtensionsURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155ExtensionsURI)
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
		it.Event = new(Erc1155ExtensionsURI)
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
func (it *Erc1155ExtensionsURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155ExtensionsURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155ExtensionsURI represents a URI event raised by the Erc1155Extensions contract.
type Erc1155ExtensionsURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Erc1155Extensions *Erc1155ExtensionsFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*Erc1155ExtensionsURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Erc1155Extensions.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &Erc1155ExtensionsURIIterator{contract: _Erc1155Extensions.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Erc1155Extensions *Erc1155ExtensionsFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *Erc1155ExtensionsURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Erc1155Extensions.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155ExtensionsURI)
				if err := _Erc1155Extensions.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Erc1155Extensions *Erc1155ExtensionsFilterer) ParseURI(log types.Log) (*Erc1155ExtensionsURI, error) {
	event := new(Erc1155ExtensionsURI)
	if err := _Erc1155Extensions.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
