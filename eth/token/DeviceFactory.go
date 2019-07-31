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

// DeviceFactoryABI is the input ABI used to generate the binding from.
const DeviceFactoryABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"roleManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x00435da5\"},{\"constant\":true,\"inputs\":[],\"name\":\"daoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x2131c68c\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc721Address\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x2352a864\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc20Address\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x276184ae\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x715018a6\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x8da5cb5b\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x8f32d59b\"},{\"constant\":true,\"inputs\":[],\"name\":\"DAOContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xb377f8ae\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xf2fde38b\"},{\"inputs\":[{\"name\":\"_daoAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"signature\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\",\"signature\":\"0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_initValue\",\"type\":\"uint256\"}],\"name\":\"createDevice\",\"outputs\":[{\"name\":\"newContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x0ab0ec34\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x41c0e1b5\"}]"

// DeviceFactory is an auto generated Go binding around an Ethereum contract.
type DeviceFactory struct {
	DeviceFactoryCaller     // Read-only binding to the contract
	DeviceFactoryTransactor // Write-only binding to the contract
	DeviceFactoryFilterer   // Log filterer for contract events
}

// DeviceFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeviceFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeviceFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeviceFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeviceFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeviceFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeviceFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeviceFactorySession struct {
	Contract     *DeviceFactory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DeviceFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeviceFactoryCallerSession struct {
	Contract *DeviceFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// DeviceFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeviceFactoryTransactorSession struct {
	Contract     *DeviceFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DeviceFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeviceFactoryRaw struct {
	Contract *DeviceFactory // Generic contract binding to access the raw methods on
}

// DeviceFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeviceFactoryCallerRaw struct {
	Contract *DeviceFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// DeviceFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeviceFactoryTransactorRaw struct {
	Contract *DeviceFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeviceFactory creates a new instance of DeviceFactory, bound to a specific deployed contract.
func NewDeviceFactory(address common.Address, backend bind.ContractBackend) (*DeviceFactory, error) {
	contract, err := bindDeviceFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DeviceFactory{DeviceFactoryCaller: DeviceFactoryCaller{contract: contract}, DeviceFactoryTransactor: DeviceFactoryTransactor{contract: contract}, DeviceFactoryFilterer: DeviceFactoryFilterer{contract: contract}}, nil
}

// NewDeviceFactoryCaller creates a new read-only instance of DeviceFactory, bound to a specific deployed contract.
func NewDeviceFactoryCaller(address common.Address, caller bind.ContractCaller) (*DeviceFactoryCaller, error) {
	contract, err := bindDeviceFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeviceFactoryCaller{contract: contract}, nil
}

// NewDeviceFactoryTransactor creates a new write-only instance of DeviceFactory, bound to a specific deployed contract.
func NewDeviceFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*DeviceFactoryTransactor, error) {
	contract, err := bindDeviceFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeviceFactoryTransactor{contract: contract}, nil
}

// NewDeviceFactoryFilterer creates a new log filterer instance of DeviceFactory, bound to a specific deployed contract.
func NewDeviceFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*DeviceFactoryFilterer, error) {
	contract, err := bindDeviceFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeviceFactoryFilterer{contract: contract}, nil
}

// bindDeviceFactory binds a generic wrapper to an already deployed contract.
func bindDeviceFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DeviceFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeviceFactory *DeviceFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DeviceFactory.Contract.DeviceFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeviceFactory *DeviceFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeviceFactory.Contract.DeviceFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeviceFactory *DeviceFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeviceFactory.Contract.DeviceFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeviceFactory *DeviceFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DeviceFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeviceFactory *DeviceFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeviceFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeviceFactory *DeviceFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeviceFactory.Contract.contract.Transact(opts, method, params...)
}

// DAOContract is a free data retrieval call binding the contract method 0xb377f8ae.
//
// Solidity: function DAOContract() constant returns(address)
func (_DeviceFactory *DeviceFactoryCaller) DAOContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DeviceFactory.contract.Call(opts, out, "DAOContract")
	return *ret0, err
}

// DAOContract is a free data retrieval call binding the contract method 0xb377f8ae.
//
// Solidity: function DAOContract() constant returns(address)
func (_DeviceFactory *DeviceFactorySession) DAOContract() (common.Address, error) {
	return _DeviceFactory.Contract.DAOContract(&_DeviceFactory.CallOpts)
}

// DAOContract is a free data retrieval call binding the contract method 0xb377f8ae.
//
// Solidity: function DAOContract() constant returns(address)
func (_DeviceFactory *DeviceFactoryCallerSession) DAOContract() (common.Address, error) {
	return _DeviceFactory.Contract.DAOContract(&_DeviceFactory.CallOpts)
}

// DaoAddress is a free data retrieval call binding the contract method 0x2131c68c.
//
// Solidity: function daoAddress() constant returns(address)
func (_DeviceFactory *DeviceFactoryCaller) DaoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DeviceFactory.contract.Call(opts, out, "daoAddress")
	return *ret0, err
}

// DaoAddress is a free data retrieval call binding the contract method 0x2131c68c.
//
// Solidity: function daoAddress() constant returns(address)
func (_DeviceFactory *DeviceFactorySession) DaoAddress() (common.Address, error) {
	return _DeviceFactory.Contract.DaoAddress(&_DeviceFactory.CallOpts)
}

// DaoAddress is a free data retrieval call binding the contract method 0x2131c68c.
//
// Solidity: function daoAddress() constant returns(address)
func (_DeviceFactory *DeviceFactoryCallerSession) DaoAddress() (common.Address, error) {
	return _DeviceFactory.Contract.DaoAddress(&_DeviceFactory.CallOpts)
}

// Erc20Address is a free data retrieval call binding the contract method 0x276184ae.
//
// Solidity: function erc20Address() constant returns(address)
func (_DeviceFactory *DeviceFactoryCaller) Erc20Address(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DeviceFactory.contract.Call(opts, out, "erc20Address")
	return *ret0, err
}

// Erc20Address is a free data retrieval call binding the contract method 0x276184ae.
//
// Solidity: function erc20Address() constant returns(address)
func (_DeviceFactory *DeviceFactorySession) Erc20Address() (common.Address, error) {
	return _DeviceFactory.Contract.Erc20Address(&_DeviceFactory.CallOpts)
}

// Erc20Address is a free data retrieval call binding the contract method 0x276184ae.
//
// Solidity: function erc20Address() constant returns(address)
func (_DeviceFactory *DeviceFactoryCallerSession) Erc20Address() (common.Address, error) {
	return _DeviceFactory.Contract.Erc20Address(&_DeviceFactory.CallOpts)
}

// Erc721Address is a free data retrieval call binding the contract method 0x2352a864.
//
// Solidity: function erc721Address() constant returns(address)
func (_DeviceFactory *DeviceFactoryCaller) Erc721Address(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DeviceFactory.contract.Call(opts, out, "erc721Address")
	return *ret0, err
}

// Erc721Address is a free data retrieval call binding the contract method 0x2352a864.
//
// Solidity: function erc721Address() constant returns(address)
func (_DeviceFactory *DeviceFactorySession) Erc721Address() (common.Address, error) {
	return _DeviceFactory.Contract.Erc721Address(&_DeviceFactory.CallOpts)
}

// Erc721Address is a free data retrieval call binding the contract method 0x2352a864.
//
// Solidity: function erc721Address() constant returns(address)
func (_DeviceFactory *DeviceFactoryCallerSession) Erc721Address() (common.Address, error) {
	return _DeviceFactory.Contract.Erc721Address(&_DeviceFactory.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_DeviceFactory *DeviceFactoryCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DeviceFactory.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_DeviceFactory *DeviceFactorySession) IsOwner() (bool, error) {
	return _DeviceFactory.Contract.IsOwner(&_DeviceFactory.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_DeviceFactory *DeviceFactoryCallerSession) IsOwner() (bool, error) {
	return _DeviceFactory.Contract.IsOwner(&_DeviceFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DeviceFactory *DeviceFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DeviceFactory.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DeviceFactory *DeviceFactorySession) Owner() (common.Address, error) {
	return _DeviceFactory.Contract.Owner(&_DeviceFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DeviceFactory *DeviceFactoryCallerSession) Owner() (common.Address, error) {
	return _DeviceFactory.Contract.Owner(&_DeviceFactory.CallOpts)
}

// RoleManager is a free data retrieval call binding the contract method 0x00435da5.
//
// Solidity: function roleManager() constant returns(address)
func (_DeviceFactory *DeviceFactoryCaller) RoleManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DeviceFactory.contract.Call(opts, out, "roleManager")
	return *ret0, err
}

// RoleManager is a free data retrieval call binding the contract method 0x00435da5.
//
// Solidity: function roleManager() constant returns(address)
func (_DeviceFactory *DeviceFactorySession) RoleManager() (common.Address, error) {
	return _DeviceFactory.Contract.RoleManager(&_DeviceFactory.CallOpts)
}

// RoleManager is a free data retrieval call binding the contract method 0x00435da5.
//
// Solidity: function roleManager() constant returns(address)
func (_DeviceFactory *DeviceFactoryCallerSession) RoleManager() (common.Address, error) {
	return _DeviceFactory.Contract.RoleManager(&_DeviceFactory.CallOpts)
}

// CreateDevice is a paid mutator transaction binding the contract method 0x0ab0ec34.
//
// Solidity: function createDevice(_name string, _initValue uint256) returns(newContract address)
func (_DeviceFactory *DeviceFactoryTransactor) CreateDevice(opts *bind.TransactOpts, _name string, _initValue *big.Int) (*types.Transaction, error) {
	return _DeviceFactory.contract.Transact(opts, "createDevice", _name, _initValue)
}

// CreateDevice is a paid mutator transaction binding the contract method 0x0ab0ec34.
//
// Solidity: function createDevice(_name string, _initValue uint256) returns(newContract address)
func (_DeviceFactory *DeviceFactorySession) CreateDevice(_name string, _initValue *big.Int) (*types.Transaction, error) {
	return _DeviceFactory.Contract.CreateDevice(&_DeviceFactory.TransactOpts, _name, _initValue)
}

// CreateDevice is a paid mutator transaction binding the contract method 0x0ab0ec34.
//
// Solidity: function createDevice(_name string, _initValue uint256) returns(newContract address)
func (_DeviceFactory *DeviceFactoryTransactorSession) CreateDevice(_name string, _initValue *big.Int) (*types.Transaction, error) {
	return _DeviceFactory.Contract.CreateDevice(&_DeviceFactory.TransactOpts, _name, _initValue)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_DeviceFactory *DeviceFactoryTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeviceFactory.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_DeviceFactory *DeviceFactorySession) Kill() (*types.Transaction, error) {
	return _DeviceFactory.Contract.Kill(&_DeviceFactory.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_DeviceFactory *DeviceFactoryTransactorSession) Kill() (*types.Transaction, error) {
	return _DeviceFactory.Contract.Kill(&_DeviceFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DeviceFactory *DeviceFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeviceFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DeviceFactory *DeviceFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _DeviceFactory.Contract.RenounceOwnership(&_DeviceFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DeviceFactory *DeviceFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DeviceFactory.Contract.RenounceOwnership(&_DeviceFactory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_DeviceFactory *DeviceFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DeviceFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_DeviceFactory *DeviceFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DeviceFactory.Contract.TransferOwnership(&_DeviceFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_DeviceFactory *DeviceFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DeviceFactory.Contract.TransferOwnership(&_DeviceFactory.TransactOpts, newOwner)
}

// DeviceFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DeviceFactory contract.
type DeviceFactoryOwnershipTransferredIterator struct {
	Event *DeviceFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DeviceFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeviceFactoryOwnershipTransferred)
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
		it.Event = new(DeviceFactoryOwnershipTransferred)
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
func (it *DeviceFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeviceFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeviceFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the DeviceFactory contract.
type DeviceFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_DeviceFactory *DeviceFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DeviceFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DeviceFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DeviceFactoryOwnershipTransferredIterator{contract: _DeviceFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_DeviceFactory *DeviceFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DeviceFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DeviceFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeviceFactoryOwnershipTransferred)
				if err := _DeviceFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
