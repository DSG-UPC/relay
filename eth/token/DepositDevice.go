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
const DepositDeviceABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DAOContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nonceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_initialDeposit\",\"type\":\"uint256\"},{\"name\":\"_daoAddress\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"},{\"name\":\"_v\",\"type\":\"uint8\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"benefit\",\"type\":\"uint256\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"},{\"name\":\"_v\",\"type\":\"uint8\"}],\"name\":\"toRepair\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"benefit\",\"type\":\"uint256\"}],\"name\":\"toItad\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"benefit\",\"type\":\"uint256\"}],\"name\":\"toRecycle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"recycle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DepositDeviceBin is the compiled bytecode used for deploying new contracts.
const DepositDeviceBin = `"0x60806040523480156200001157600080fd5b50604051620025e2380380620025e2833981018060405281019080805182019291906020018051906020019092919080519060200190929190805190602001909291905050506000806000336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a383600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634ece90a86040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015620001e057600080fd5b505af1158015620001f5573d6000803e3d6000fd5b505050506040513d60208110156200020c57600080fd5b81019080805190602001909291905050509250600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a564725e6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015620002a657600080fd5b505af1158015620002bb573d6000803e3d6000fd5b505050506040513d6020811015620002d257600080fd5b81019080805190602001909291905050509150600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166351331ad76040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156200036c57600080fd5b505af115801562000381573d6000803e3d6000fd5b505050506040513d60208110156200039857600080fd5b8101908080519060200190929190505050905080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f13eed97336040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b158015620004aa57600080fd5b505af1158015620004bf573d6000803e3d6000fd5b505050506040513d6020811015620004d657600080fd5b8101908080519060200190929190505050151562000582576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001807f546869732064657669636520636f6e747261637420776173206e6f742063726581526020017f617465642062792061204e6f746172790000000000000000000000000000000081525060400191505060405180910390fd5b81600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555086600660000190805190602001906200061f92919062000829565b5085600660040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555084600660030181905550620006888662000695640100000000026401000000009004565b50505050505050620008d8565b620006ae620006d7640100000000026401000000009004565b1515620006ba57600080fd5b620006d4816200072e640100000000026401000000009004565b50565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156200076b57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200086c57805160ff19168380011785556200089d565b828001600101855582156200089d579182015b828111156200089c5782518255916020019190600101906200087f565b5b509050620008ac9190620008b0565b5090565b620008d591905b80821115620008d1576000816000905550600101620008b7565b5090565b90565b611cfa80620008e86000396000f3006080604052600436106100af576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630f23cbaa146100b45780632b95cb2b146100cb57806337c1e0fe1461011857806348628f0514610184578063715018a6146101fa5780638da5cb5b146102115780638f32d59b14610268578063ae140ed714610297578063b377f8ae146102e4578063ed2a2d641461033b578063f2fde38b14610392575b600080fd5b3480156100c057600080fd5b506100c96103d5565b005b3480156100d757600080fd5b50610116600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610797565b005b34801561012457600080fd5b50610182600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080356000191690602001909291908035600019169060200190929190803560ff169060200190929190505050610912565b005b34801561019057600080fd5b506101f8600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080356000191690602001909291908035600019169060200190929190803560ff169060200190929190505050610d7a565b005b34801561020657600080fd5b5061020f61111f565b005b34801561021d57600080fd5b506102266111f1565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561027457600080fd5b5061027d61121a565b604051808215151515815260200191505060405180910390f35b3480156102a357600080fd5b506102e2600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611271565b005b3480156102f057600080fd5b506102f9611411565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561034757600080fd5b5061037c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611437565b6040518082815260200191505060405180910390f35b34801561039e57600080fd5b506103d3600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611480565b005b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630b73d13e336040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b15801561049257600080fd5b505af11580156104a6573d6000803e3d6000fd5b505050506040513d60208110156104bc57600080fd5b81019080805190602001909291905050501515610567576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602d8152602001807f54686973207265717565737420776173206e6f74206f726967696e617465642081526020017f627920612052657061697265720000000000000000000000000000000000000081525060400191505060405180910390fd5b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd30333073ffffffffffffffffffffffffffffffffffffffff16316040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b15801561067757600080fd5b505af115801561068b573d6000803e3d6000fd5b505050506040513d60208110156106a157600080fd5b810190808051906020019092919050505050600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639dc29fac336006600101546040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15801561077d57600080fd5b505af1158015610791573d6000803e3d6000fd5b50505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630b73d13e836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b15801561085457600080fd5b505af1158015610868573d6000803e3d6000fd5b505050506040513d602081101561087e57600080fd5b81019080805190602001909291905050501515610903576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f5468652064657374696e6174696f6e206973206e6f7420616e2069746164000081525060200191505060405180910390fd5b61090e33838361149f565b5050565b61091a61121a565b151561092557600080fd5b61098c3385604051602001808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c0100000000000000000000000002815260140191505060405160208183030381529060405285858561181a565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632bd2a191856040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b158015610a4957600080fd5b505af1158015610a5d573d6000803e3d6000fd5b505050506040513d6020811015610a7357600080fd5b81019080805190602001909291905050501515610b1e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001807f5468652064657374696e6174696f6e206973206e6f74206120636f6e73756d6581526020017f720000000000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd33306006600301546040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b158015610c1c57600080fd5b505af1158015610c30573d6000803e3d6000fd5b505050506040513d6020811015610c4657600080fd5b810190808051906020019092919050505050600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f19853073ffffffffffffffffffffffffffffffffffffffff166040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015610d3357600080fd5b505af1158015610d47573d6000803e3d6000fd5b505050503073ffffffffffffffffffffffffffffffffffffffff16600660010181905550610d7433611480565b50505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630b73d13e336040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b158015610e3757600080fd5b505af1158015610e4b573d6000803e3d6000fd5b505050506040513d6020811015610e6157600080fd5b81019080805190602001909291905050501515610f0c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602d8152602001807f54686973207265717565737420776173206e6f74206f726967696e617465642081526020017f627920612052657061697265720000000000000000000000000000000000000081525060400191505060405180910390fd5b610f7b338686604051602001808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018281526020019250505060405160208183030381529060405285858561181a565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632bd2a191866040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b15801561103857600080fd5b505af115801561104c573d6000803e3d6000fd5b505050506040513d602081101561106257600080fd5b8101908080519060200190929190505050151561110d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001807f5468652064657374696e6174696f6e206973206e6f742061207265706169726581526020017f720000000000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b61111833868661149f565b5050505050565b61112761121a565b151561113257600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1662091988836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b15801561132d57600080fd5b505af1158015611341573d6000803e3d6000fd5b505050506040513d602081101561135757600080fd5b81019080805190602001909291905050501515611402576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260238152602001807f5468652064657374696e6174696f6e206973206e6f742061206120726563796381526020017f6c6572000000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b61140d33838361149f565b5050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b61148861121a565b151561149357600080fd5b61149c81611bd4565b50565b6114a761121a565b15156114b257600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415151561157d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260278152602001807f5468652064657374696e6174696f6e2063616e6e6f742062652074686520302081526020017f616464726573730000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd84846006600101546040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b15801561167b57600080fd5b505af115801561168f573d6000803e3d6000fd5b50505050600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd8484846040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b15801561178c57600080fd5b505af11580156117a0573d6000803e3d6000fd5b505050506040513d60208110156117b657600080fd5b8101908080519060200190929190505050506117d182611480565b81600660040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b60008060197f01000000000000000000000000000000000000000000000000000000000000000260007f01000000000000000000000000000000000000000000000000000000000000000230600160008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548960405160200180867effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152600101857effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526001018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c0100000000000000000000000002815260140183815260200182805190602001908083835b6020831015156119ba5780518252602082019150602081019050602083039250611995565b6001836020036101000a038019825116818451168082178552505050505050905001955050505050506040516020818303038152906040526040518082805190602001908083835b602083101515611a275780518252602082019150602081019050602083039250611a02565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209150600182848787604051600081526020016040526040518085600019166000191681526020018460ff1660ff1681526020018360001916600019168152602001826000191660001916815260200194505050505060206040516020810390808403906000865af1158015611acd573d6000803e3d6000fd5b5050506020604051035190508673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515611b7c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f73656e6465722d616464726573732d646f65732d6e6f742d6d6174636800000081525060200191505060405180910390fd5b600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000815480929190600101919050555050505050505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515611c1057600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a7230582064eb8a2ac0bf1bec0827112728a13c5d3acb81a00fffca942affe5a1850214ba0029"`

// DeployDepositDevice deploys a new Ethereum contract, binding an instance of DepositDevice to it.
func DeployDepositDevice(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _sender common.Address, _initialDeposit *big.Int, _daoAddress *big.Int) (common.Address, *types.Transaction, *DepositDevice, error) {
	parsed, err := abi.JSON(strings.NewReader(DepositDeviceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DepositDeviceBin), backend, _name, _sender, _initialDeposit, _daoAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DepositDevice{DepositDeviceCaller: DepositDeviceCaller{contract: contract}, DepositDeviceTransactor: DepositDeviceTransactor{contract: contract}, DepositDeviceFilterer: DepositDeviceFilterer{contract: contract}}, nil
}

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

// DAOContract is a free data retrieval call binding the contract method 0xb377f8ae.
//
// Solidity: function DAOContract() constant returns(address)
func (_DepositDevice *DepositDeviceCaller) DAOContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DepositDevice.contract.Call(opts, out, "DAOContract")
	return *ret0, err
}

// DAOContract is a free data retrieval call binding the contract method 0xb377f8ae.
//
// Solidity: function DAOContract() constant returns(address)
func (_DepositDevice *DepositDeviceSession) DAOContract() (common.Address, error) {
	return _DepositDevice.Contract.DAOContract(&_DepositDevice.CallOpts)
}

// DAOContract is a free data retrieval call binding the contract method 0xb377f8ae.
//
// Solidity: function DAOContract() constant returns(address)
func (_DepositDevice *DepositDeviceCallerSession) DAOContract() (common.Address, error) {
	return _DepositDevice.Contract.DAOContract(&_DepositDevice.CallOpts)
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

// NonceOf is a free data retrieval call binding the contract method 0xed2a2d64.
//
// Solidity: function nonceOf(_owner address) constant returns(uint256)
func (_DepositDevice *DepositDeviceCaller) NonceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DepositDevice.contract.Call(opts, out, "nonceOf", _owner)
	return *ret0, err
}

// NonceOf is a free data retrieval call binding the contract method 0xed2a2d64.
//
// Solidity: function nonceOf(_owner address) constant returns(uint256)
func (_DepositDevice *DepositDeviceSession) NonceOf(_owner common.Address) (*big.Int, error) {
	return _DepositDevice.Contract.NonceOf(&_DepositDevice.CallOpts, _owner)
}

// NonceOf is a free data retrieval call binding the contract method 0xed2a2d64.
//
// Solidity: function nonceOf(_owner address) constant returns(uint256)
func (_DepositDevice *DepositDeviceCallerSession) NonceOf(_owner common.Address) (*big.Int, error) {
	return _DepositDevice.Contract.NonceOf(&_DepositDevice.CallOpts, _owner)
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

// Mint is a paid mutator transaction binding the contract method 0x37c1e0fe.
//
// Solidity: function mint(_to address, _r bytes32, _s bytes32, _v uint8) returns()
func (_DepositDevice *DepositDeviceTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _r [32]byte, _s [32]byte, _v uint8) (*types.Transaction, error) {
	return _DepositDevice.contract.Transact(opts, "mint", _to, _r, _s, _v)
}

// Mint is a paid mutator transaction binding the contract method 0x37c1e0fe.
//
// Solidity: function mint(_to address, _r bytes32, _s bytes32, _v uint8) returns()
func (_DepositDevice *DepositDeviceSession) Mint(_to common.Address, _r [32]byte, _s [32]byte, _v uint8) (*types.Transaction, error) {
	return _DepositDevice.Contract.Mint(&_DepositDevice.TransactOpts, _to, _r, _s, _v)
}

// Mint is a paid mutator transaction binding the contract method 0x37c1e0fe.
//
// Solidity: function mint(_to address, _r bytes32, _s bytes32, _v uint8) returns()
func (_DepositDevice *DepositDeviceTransactorSession) Mint(_to common.Address, _r [32]byte, _s [32]byte, _v uint8) (*types.Transaction, error) {
	return _DepositDevice.Contract.Mint(&_DepositDevice.TransactOpts, _to, _r, _s, _v)
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

// ToRepair is a paid mutator transaction binding the contract method 0x48628f05.
//
// Solidity: function toRepair(_to address, benefit uint256, _r bytes32, _s bytes32, _v uint8) returns()
func (_DepositDevice *DepositDeviceTransactor) ToRepair(opts *bind.TransactOpts, _to common.Address, benefit *big.Int, _r [32]byte, _s [32]byte, _v uint8) (*types.Transaction, error) {
	return _DepositDevice.contract.Transact(opts, "toRepair", _to, benefit, _r, _s, _v)
}

// ToRepair is a paid mutator transaction binding the contract method 0x48628f05.
//
// Solidity: function toRepair(_to address, benefit uint256, _r bytes32, _s bytes32, _v uint8) returns()
func (_DepositDevice *DepositDeviceSession) ToRepair(_to common.Address, benefit *big.Int, _r [32]byte, _s [32]byte, _v uint8) (*types.Transaction, error) {
	return _DepositDevice.Contract.ToRepair(&_DepositDevice.TransactOpts, _to, benefit, _r, _s, _v)
}

// ToRepair is a paid mutator transaction binding the contract method 0x48628f05.
//
// Solidity: function toRepair(_to address, benefit uint256, _r bytes32, _s bytes32, _v uint8) returns()
func (_DepositDevice *DepositDeviceTransactorSession) ToRepair(_to common.Address, benefit *big.Int, _r [32]byte, _s [32]byte, _v uint8) (*types.Transaction, error) {
	return _DepositDevice.Contract.ToRepair(&_DepositDevice.TransactOpts, _to, benefit, _r, _s, _v)
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
