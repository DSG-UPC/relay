// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// DepositDeviceABI is the input ABI used to generate the binding from.
const DepositDeviceABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_initialDeposit\",\"type\":\"uint256\"},{\"name\":\"_erc20\",\"type\":\"address\"},{\"name\":\"_erc721\",\"type\":\"address\"},{\"name\":\"_roleManager\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"benefit\",\"type\":\"uint256\"}],\"name\":\"toRepair\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"benefit\",\"type\":\"uint256\"}],\"name\":\"toItad\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"benefit\",\"type\":\"uint256\"}],\"name\":\"toRecycle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"recycle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DepositDevice is an auto generated Go binding around an Ethereum contract.
type DepositDevice struct {
	DepositDeviceCaller     // Read-only binding to the contract
	DepositDeviceTransactor // Write-only binding to the contract
	DepositDeviceFilterer   // Log filterer for contract events
}

// DepositDeviceCaller is an auto generated read-only Go binding around an Ethereum contract.
type DepositDeviceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositDeviceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DepositDeviceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositDeviceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DepositDeviceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositDeviceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DepositDeviceSession struct {
	Contract     *DepositDevice    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DepositDeviceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DepositDeviceCallerSession struct {
	Contract *DepositDeviceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// DepositDeviceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DepositDeviceTransactorSession struct {
	Contract     *DepositDeviceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DepositDeviceRaw is an auto generated low-level Go binding around an Ethereum contract.
type DepositDeviceRaw struct {
	Contract *DepositDevice // Generic contract binding to access the raw methods on
}

// DepositDeviceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DepositDeviceCallerRaw struct {
	Contract *DepositDeviceCaller // Generic read-only contract binding to access the raw methods on
}

// DepositDeviceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DepositDeviceTransactorRaw struct {
	Contract *DepositDeviceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDepositDevice creates a new instance of DepositDevice, bound to a specific deployed contract.
func NewDepositDevice(address common.Address, backend bind.ContractBackend) (*DepositDevice, error) {
	contract, err := bindDepositDevice(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DepositDevice{DepositDeviceCaller: DepositDeviceCaller{contract: contract}, DepositDeviceTransactor: DepositDeviceTransactor{contract: contract}, DepositDeviceFilterer: DepositDeviceFilterer{contract: contract}}, nil
}

// NewDepositDeviceCaller creates a new read-only instance of DepositDevice, bound to a specific deployed contract.
func NewDepositDeviceCaller(address common.Address, caller bind.ContractCaller) (*DepositDeviceCaller, error) {
	contract, err := bindDepositDevice(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DepositDeviceCaller{contract: contract}, nil
}

// NewDepositDeviceTransactor creates a new write-only instance of DepositDevice, bound to a specific deployed contract.
func NewDepositDeviceTransactor(address common.Address, transactor bind.ContractTransactor) (*DepositDeviceTransactor, error) {
	contract, err := bindDepositDevice(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DepositDeviceTransactor{contract: contract}, nil
}

// NewDepositDeviceFilterer creates a new log filterer instance of DepositDevice, bound to a specific deployed contract.
func NewDepositDeviceFilterer(address common.Address, filterer bind.ContractFilterer) (*DepositDeviceFilterer, error) {
	contract, err := bindDepositDevice(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DepositDeviceFilterer{contract: contract}, nil
}

// bindDepositDevice binds a generic wrapper to an already deployed contract.
func bindDepositDevice(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DepositDeviceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DepositDevice *DepositDeviceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DepositDevice.Contract.DepositDeviceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DepositDevice *DepositDeviceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositDevice.Contract.DepositDeviceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DepositDevice *DepositDeviceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositDevice.Contract.DepositDeviceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DepositDevice *DepositDeviceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DepositDevice.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DepositDevice *DepositDeviceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositDevice.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DepositDevice *DepositDeviceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositDevice.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_DepositDevice *DepositDeviceCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DepositDevice.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_DepositDevice *DepositDeviceSession) IsOwner() (bool, error) {
	return _DepositDevice.Contract.IsOwner(&_DepositDevice.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_DepositDevice *DepositDeviceCallerSession) IsOwner() (bool, error) {
	return _DepositDevice.Contract.IsOwner(&_DepositDevice.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DepositDevice *DepositDeviceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DepositDevice.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DepositDevice *DepositDeviceSession) Owner() (common.Address, error) {
	return _DepositDevice.Contract.Owner(&_DepositDevice.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DepositDevice *DepositDeviceCallerSession) Owner() (common.Address, error) {
	return _DepositDevice.Contract.Owner(&_DepositDevice.CallOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(_to address) returns()
func (_DepositDevice *DepositDeviceTransactor) Mint(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _DepositDevice.contract.Transact(opts, "mint", _to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(_to address) returns()
func (_DepositDevice *DepositDeviceSession) Mint(_to common.Address) (*types.Transaction, error) {
	return _DepositDevice.Contract.Mint(&_DepositDevice.TransactOpts, _to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(_to address) returns()
func (_DepositDevice *DepositDeviceTransactorSession) Mint(_to common.Address) (*types.Transaction, error) {
	return _DepositDevice.Contract.Mint(&_DepositDevice.TransactOpts, _to)
}

// Recycle is a paid mutator transaction binding the contract method 0x0f23cbaa.
//
// Solidity: function recycle() returns()
func (_DepositDevice *DepositDeviceTransactor) Recycle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositDevice.contract.Transact(opts, "recycle")
}

// Recycle is a paid mutator transaction binding the contract method 0x0f23cbaa.
//
// Solidity: function recycle() returns()
func (_DepositDevice *DepositDeviceSession) Recycle() (*types.Transaction, error) {
	return _DepositDevice.Contract.Recycle(&_DepositDevice.TransactOpts)
}

// Recycle is a paid mutator transaction binding the contract method 0x0f23cbaa.
//
// Solidity: function recycle() returns()
func (_DepositDevice *DepositDeviceTransactorSession) Recycle() (*types.Transaction, error) {
	return _DepositDevice.Contract.Recycle(&_DepositDevice.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DepositDevice *DepositDeviceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositDevice.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DepositDevice *DepositDeviceSession) RenounceOwnership() (*types.Transaction, error) {
	return _DepositDevice.Contract.RenounceOwnership(&_DepositDevice.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DepositDevice *DepositDeviceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DepositDevice.Contract.RenounceOwnership(&_DepositDevice.TransactOpts)
}

// ToItad is a paid mutator transaction binding the contract method 0x2b95cb2b.
//
// Solidity: function toItad(_to address, benefit uint256) returns()
func (_DepositDevice *DepositDeviceTransactor) ToItad(opts *bind.TransactOpts, _to common.Address, benefit *big.Int) (*types.Transaction, error) {
	return _DepositDevice.contract.Transact(opts, "toItad", _to, benefit)
}

// ToItad is a paid mutator transaction binding the contract method 0x2b95cb2b.
//
// Solidity: function toItad(_to address, benefit uint256) returns()
func (_DepositDevice *DepositDeviceSession) ToItad(_to common.Address, benefit *big.Int) (*types.Transaction, error) {
	return _DepositDevice.Contract.ToItad(&_DepositDevice.TransactOpts, _to, benefit)
}

// ToItad is a paid mutator transaction binding the contract method 0x2b95cb2b.
//
// Solidity: function toItad(_to address, benefit uint256) returns()
func (_DepositDevice *DepositDeviceTransactorSession) ToItad(_to common.Address, benefit *big.Int) (*types.Transaction, error) {
	return _DepositDevice.Contract.ToItad(&_DepositDevice.TransactOpts, _to, benefit)
}

// ToRecycle is a paid mutator transaction binding the contract method 0xae140ed7.
//
// Solidity: function toRecycle(_to address, benefit uint256) returns()
func (_DepositDevice *DepositDeviceTransactor) ToRecycle(opts *bind.TransactOpts, _to common.Address, benefit *big.Int) (*types.Transaction, error) {
	return _DepositDevice.contract.Transact(opts, "toRecycle", _to, benefit)
}

// ToRecycle is a paid mutator transaction binding the contract method 0xae140ed7.
//
// Solidity: function toRecycle(_to address, benefit uint256) returns()
func (_DepositDevice *DepositDeviceSession) ToRecycle(_to common.Address, benefit *big.Int) (*types.Transaction, error) {
	return _DepositDevice.Contract.ToRecycle(&_DepositDevice.TransactOpts, _to, benefit)
}

// ToRecycle is a paid mutator transaction binding the contract method 0xae140ed7.
//
// Solidity: function toRecycle(_to address, benefit uint256) returns()
func (_DepositDevice *DepositDeviceTransactorSession) ToRecycle(_to common.Address, benefit *big.Int) (*types.Transaction, error) {
	return _DepositDevice.Contract.ToRecycle(&_DepositDevice.TransactOpts, _to, benefit)
}

// ToRepair is a paid mutator transaction binding the contract method 0x4f6a592d.
//
// Solidity: function toRepair(_to address, benefit uint256) returns()
func (_DepositDevice *DepositDeviceTransactor) ToRepair(opts *bind.TransactOpts, _to common.Address, benefit *big.Int) (*types.Transaction, error) {
	return _DepositDevice.contract.Transact(opts, "toRepair", _to, benefit)
}

// ToRepair is a paid mutator transaction binding the contract method 0x4f6a592d.
//
// Solidity: function toRepair(_to address, benefit uint256) returns()
func (_DepositDevice *DepositDeviceSession) ToRepair(_to common.Address, benefit *big.Int) (*types.Transaction, error) {
	return _DepositDevice.Contract.ToRepair(&_DepositDevice.TransactOpts, _to, benefit)
}

// ToRepair is a paid mutator transaction binding the contract method 0x4f6a592d.
//
// Solidity: function toRepair(_to address, benefit uint256) returns()
func (_DepositDevice *DepositDeviceTransactorSession) ToRepair(_to common.Address, benefit *big.Int) (*types.Transaction, error) {
	return _DepositDevice.Contract.ToRepair(&_DepositDevice.TransactOpts, _to, benefit)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_DepositDevice *DepositDeviceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DepositDevice.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_DepositDevice *DepositDeviceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DepositDevice.Contract.TransferOwnership(&_DepositDevice.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_DepositDevice *DepositDeviceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DepositDevice.Contract.TransferOwnership(&_DepositDevice.TransactOpts, newOwner)
}

// DepositDeviceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DepositDevice contract.
type DepositDeviceOwnershipTransferredIterator struct {
	Event *DepositDeviceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DepositDeviceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositDeviceOwnershipTransferred)
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
		it.Event = new(DepositDeviceOwnershipTransferred)
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
func (it *DepositDeviceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositDeviceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositDeviceOwnershipTransferred represents a OwnershipTransferred event raised by the DepositDevice contract.
type DepositDeviceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_DepositDevice *DepositDeviceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DepositDeviceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DepositDevice.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DepositDeviceOwnershipTransferredIterator{contract: _DepositDevice.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_DepositDevice *DepositDeviceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DepositDeviceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DepositDevice.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositDeviceOwnershipTransferred)
				if err := _DepositDevice.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
