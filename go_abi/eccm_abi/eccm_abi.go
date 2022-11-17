// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eccm_abi

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BaseAdminUpgradeabilityProxyABI is the input ABI used to generate the binding from.
const BaseAdminUpgradeabilityProxyABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// BaseAdminUpgradeabilityProxyFuncSigs maps the 4-byte function signature to its string representation.
var BaseAdminUpgradeabilityProxyFuncSigs = map[string]string{
	"f851a440": "admin()",
	"8f283970": "changeAdmin(address)",
	"5c60da1b": "implementation()",
	"3659cfe6": "upgradeTo(address)",
	"4f1ef286": "upgradeToAndCall(address,bytes)",
}

// BaseAdminUpgradeabilityProxyBin is the compiled bytecode used for deploying new contracts.
var BaseAdminUpgradeabilityProxyBin = "0x608060405234801561001057600080fd5b5061061b806100206000396000f3fe60806040526004361061004a5760003560e01c80633659cfe6146100545780634f1ef286146100875780635c60da1b146101075780638f28397014610138578063f851a4401461016b575b610052610180565b005b34801561006057600080fd5b506100526004803603602081101561007757600080fd5b50356001600160a01b031661019a565b6100526004803603604081101561009d57600080fd5b6001600160a01b0382351691908101906040810160208201356401000000008111156100c857600080fd5b8201836020820111156100da57600080fd5b803590602001918460018302840111640100000000831117156100fc57600080fd5b5090925090506101d4565b34801561011357600080fd5b5061011c610281565b604080516001600160a01b039092168252519081900360200190f35b34801561014457600080fd5b506100526004803603602081101561015b57600080fd5b50356001600160a01b03166102be565b34801561017757600080fd5b5061011c610378565b6101886103a3565b610198610193610403565b610428565b565b6101a261044c565b6001600160a01b0316336001600160a01b031614156101c9576101c481610471565b6101d1565b6101d1610180565b50565b6101dc61044c565b6001600160a01b0316336001600160a01b03161415610274576101fe83610471565b6000836001600160a01b031683836040518083838082843760405192019450600093509091505080830381855af49150503d806000811461025b576040519150601f19603f3d011682016040523d82523d6000602084013e610260565b606091505b505090508061026e57600080fd5b5061027c565b61027c610180565b505050565b600061028b61044c565b6001600160a01b0316336001600160a01b031614156102b3576102ac610403565b90506102bb565b6102bb610180565b90565b6102c661044c565b6001600160a01b0316336001600160a01b031614156101c9576001600160a01b0381166103245760405162461bcd60e51b81526004018080602001828103825260368152602001806105766036913960400191505060405180910390fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61034d61044c565b604080516001600160a01b03928316815291841660208301528051918290030190a16101c4816104b1565b600061038261044c565b6001600160a01b0316336001600160a01b031614156102b3576102ac61044c565b6103ab61044c565b6001600160a01b0316336001600160a01b031614156103fb5760405162461bcd60e51b81526004018080602001828103825260328152602001806105446032913960400191505060405180910390fd5b610198610198565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5490565b3660008037600080366000845af43d6000803e808015610447573d6000f35b3d6000fd5b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035490565b61047a816104d5565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610355565b6104de8161053d565b6105195760405162461bcd60e51b815260040180806020018281038252603b8152602001806105ac603b913960400191505060405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc55565b3b15159056fe43616e6e6f742063616c6c2066616c6c6261636b2066756e6374696f6e2066726f6d207468652070726f78792061646d696e43616e6e6f74206368616e6765207468652061646d696e206f6620612070726f787920746f20746865207a65726f206164647265737343616e6e6f742073657420612070726f787920696d706c656d656e746174696f6e20746f2061206e6f6e2d636f6e74726163742061646472657373a265627a7a72315820829cb677156af4a8b756390600d941570989f86c82d3bf047e94e74ceaa2e7d464736f6c63430005110032"

// DeployBaseAdminUpgradeabilityProxy deploys a new Ethereum contract, binding an instance of BaseAdminUpgradeabilityProxy to it.
func DeployBaseAdminUpgradeabilityProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BaseAdminUpgradeabilityProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseAdminUpgradeabilityProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BaseAdminUpgradeabilityProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BaseAdminUpgradeabilityProxy{BaseAdminUpgradeabilityProxyCaller: BaseAdminUpgradeabilityProxyCaller{contract: contract}, BaseAdminUpgradeabilityProxyTransactor: BaseAdminUpgradeabilityProxyTransactor{contract: contract}, BaseAdminUpgradeabilityProxyFilterer: BaseAdminUpgradeabilityProxyFilterer{contract: contract}}, nil
}

// BaseAdminUpgradeabilityProxy is an auto generated Go binding around an Ethereum contract.
type BaseAdminUpgradeabilityProxy struct {
	BaseAdminUpgradeabilityProxyCaller     // Read-only binding to the contract
	BaseAdminUpgradeabilityProxyTransactor // Write-only binding to the contract
	BaseAdminUpgradeabilityProxyFilterer   // Log filterer for contract events
}

// BaseAdminUpgradeabilityProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseAdminUpgradeabilityProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseAdminUpgradeabilityProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseAdminUpgradeabilityProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseAdminUpgradeabilityProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseAdminUpgradeabilityProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseAdminUpgradeabilityProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseAdminUpgradeabilityProxySession struct {
	Contract     *BaseAdminUpgradeabilityProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// BaseAdminUpgradeabilityProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseAdminUpgradeabilityProxyCallerSession struct {
	Contract *BaseAdminUpgradeabilityProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// BaseAdminUpgradeabilityProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseAdminUpgradeabilityProxyTransactorSession struct {
	Contract     *BaseAdminUpgradeabilityProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// BaseAdminUpgradeabilityProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseAdminUpgradeabilityProxyRaw struct {
	Contract *BaseAdminUpgradeabilityProxy // Generic contract binding to access the raw methods on
}

// BaseAdminUpgradeabilityProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseAdminUpgradeabilityProxyCallerRaw struct {
	Contract *BaseAdminUpgradeabilityProxyCaller // Generic read-only contract binding to access the raw methods on
}

// BaseAdminUpgradeabilityProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseAdminUpgradeabilityProxyTransactorRaw struct {
	Contract *BaseAdminUpgradeabilityProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseAdminUpgradeabilityProxy creates a new instance of BaseAdminUpgradeabilityProxy, bound to a specific deployed contract.
func NewBaseAdminUpgradeabilityProxy(address common.Address, backend bind.ContractBackend) (*BaseAdminUpgradeabilityProxy, error) {
	contract, err := bindBaseAdminUpgradeabilityProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseAdminUpgradeabilityProxy{BaseAdminUpgradeabilityProxyCaller: BaseAdminUpgradeabilityProxyCaller{contract: contract}, BaseAdminUpgradeabilityProxyTransactor: BaseAdminUpgradeabilityProxyTransactor{contract: contract}, BaseAdminUpgradeabilityProxyFilterer: BaseAdminUpgradeabilityProxyFilterer{contract: contract}}, nil
}

// NewBaseAdminUpgradeabilityProxyCaller creates a new read-only instance of BaseAdminUpgradeabilityProxy, bound to a specific deployed contract.
func NewBaseAdminUpgradeabilityProxyCaller(address common.Address, caller bind.ContractCaller) (*BaseAdminUpgradeabilityProxyCaller, error) {
	contract, err := bindBaseAdminUpgradeabilityProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseAdminUpgradeabilityProxyCaller{contract: contract}, nil
}

// NewBaseAdminUpgradeabilityProxyTransactor creates a new write-only instance of BaseAdminUpgradeabilityProxy, bound to a specific deployed contract.
func NewBaseAdminUpgradeabilityProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseAdminUpgradeabilityProxyTransactor, error) {
	contract, err := bindBaseAdminUpgradeabilityProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseAdminUpgradeabilityProxyTransactor{contract: contract}, nil
}

// NewBaseAdminUpgradeabilityProxyFilterer creates a new log filterer instance of BaseAdminUpgradeabilityProxy, bound to a specific deployed contract.
func NewBaseAdminUpgradeabilityProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseAdminUpgradeabilityProxyFilterer, error) {
	contract, err := bindBaseAdminUpgradeabilityProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseAdminUpgradeabilityProxyFilterer{contract: contract}, nil
}

// bindBaseAdminUpgradeabilityProxy binds a generic wrapper to an already deployed contract.
func bindBaseAdminUpgradeabilityProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseAdminUpgradeabilityProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseAdminUpgradeabilityProxy *BaseAdminUpgradeabilityProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseAdminUpgradeabilityProxy.Contract.BaseAdminUpgradeabilityProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseAdminUpgradeabilityProxy *BaseAdminUpgradeabilityProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseAdminUpgradeabilityProxy.Contract.BaseAdminUpgradeabilityProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseAdminUpgradeabilityProxy *BaseAdminUpgradeabilityProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseAdminUpgradeabilityProxy.Contract.BaseAdminUpgradeabilityProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseAdminUpgradeabilityProxy *BaseAdminUpgradeabilityProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseAdminUpgradeabilityProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseAdminUpgradeabilityProxy *BaseAdminUpgradeabilityProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseAdminUpgradeabilityProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseAdminUpgradeabilityProxy *BaseAdminUpgradeabilityProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseAdminUpgradeabilityProxy.Contract.contract.Transact(opts, method, params...)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_BaseAdminUpgradeabilityProxy *BaseAdminUpgradeabilityProxyTransactor) Admin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseAdminUpgradeabilityProxy.contract.Transact(opts, "admin")
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_BaseAdminUpgradeabilityProxy *BaseAdminUpgradeabilityProxySession) Admin() (*types.Transaction, error) {
	return _BaseAdminUpgradeabilityProxy.Contract.Admin(&_BaseAdminUpgradeabilityProxy.TransactOpts)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_BaseAdminUpgradeabilityProxy *BaseAdminUpgradeabilityProxyTransactorSession) Admin() (*types.Transaction, error) {
	return _BaseAdminUpgradeabilityProxy.Contract.Admin(&_BaseAdminUpgradeabilityProxy.TransactOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_BaseAdminUpgradeabilityProxy *BaseAdminUpgradeabilityProxyTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _BaseAdminUpgradeabilityProxy.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_BaseAdminUpgradeabilityProxy *BaseAdminUpgradeabilityProxySession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _BaseAdminUpgradeabilityProxy.Contract.ChangeAdmin(&_BaseAdminUpgradeabilityProxy.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_BaseAdminUpgradeabilityProxy *BaseAdminUpgradeabilityProxyTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _BaseAdminUpgradeabilityProxy.Contract.ChangeAdmin(&_BaseAdminUpgradeabilityProxy.TransactOpts, newAdmin)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_BaseAdminUpgradeabilityProxy *BaseAdminUpgradeabilityProxyTransactor) Implementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseAdminUpgradeabilityProxy.contract.Transact(opts, "implementation")
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_BaseAdminUpgradeabilityProxy *BaseAdminUpgradeabilityProxySession) Implementation() (*types.Transaction, error) {
	return _BaseAdminUpgradeabilityProxy.Contract.Implementation(&_BaseAdminUpgradeabilityProxy.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_BaseAdminUpgradeabilityProxy *BaseAdminUpgradeabilityProxyTransactorSession) Implementation() (*types.Transaction, error) {
	return _BaseAdminUpgradeabilityProxy.Contract.Implementation(&_BaseAdminUpgradeabilityProxy.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_BaseAdminUpgradeabilityProxy *BaseAdminUpgradeabilityProxyTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _BaseAdminUpgradeabilityProxy.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_BaseAdminUpgradeabilityProxy *BaseAdminUpgradeabilityProxySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _BaseAdminUpgradeabilityProxy.Contract.UpgradeTo(&_BaseAdminUpgradeabilityProxy.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_BaseAdminUpgradeabilityProxy *BaseAdminUpgradeabilityProxyTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _BaseAdminUpgradeabilityProxy.Contract.UpgradeTo(&_BaseAdminUpgradeabilityProxy.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BaseAdminUpgradeabilityProxy *BaseAdminUpgradeabilityProxyTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BaseAdminUpgradeabilityProxy.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BaseAdminUpgradeabilityProxy *BaseAdminUpgradeabilityProxySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BaseAdminUpgradeabilityProxy.Contract.UpgradeToAndCall(&_BaseAdminUpgradeabilityProxy.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BaseAdminUpgradeabilityProxy *BaseAdminUpgradeabilityProxyTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BaseAdminUpgradeabilityProxy.Contract.UpgradeToAndCall(&_BaseAdminUpgradeabilityProxy.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BaseAdminUpgradeabilityProxy *BaseAdminUpgradeabilityProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _BaseAdminUpgradeabilityProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BaseAdminUpgradeabilityProxy *BaseAdminUpgradeabilityProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BaseAdminUpgradeabilityProxy.Contract.Fallback(&_BaseAdminUpgradeabilityProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BaseAdminUpgradeabilityProxy *BaseAdminUpgradeabilityProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BaseAdminUpgradeabilityProxy.Contract.Fallback(&_BaseAdminUpgradeabilityProxy.TransactOpts, calldata)
}

// BaseAdminUpgradeabilityProxyAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the BaseAdminUpgradeabilityProxy contract.
type BaseAdminUpgradeabilityProxyAdminChangedIterator struct {
	Event *BaseAdminUpgradeabilityProxyAdminChanged // Event containing the contract specifics and raw log

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
func (it *BaseAdminUpgradeabilityProxyAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseAdminUpgradeabilityProxyAdminChanged)
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
		it.Event = new(BaseAdminUpgradeabilityProxyAdminChanged)
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
func (it *BaseAdminUpgradeabilityProxyAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseAdminUpgradeabilityProxyAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseAdminUpgradeabilityProxyAdminChanged represents a AdminChanged event raised by the BaseAdminUpgradeabilityProxy contract.
type BaseAdminUpgradeabilityProxyAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_BaseAdminUpgradeabilityProxy *BaseAdminUpgradeabilityProxyFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*BaseAdminUpgradeabilityProxyAdminChangedIterator, error) {

	logs, sub, err := _BaseAdminUpgradeabilityProxy.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &BaseAdminUpgradeabilityProxyAdminChangedIterator{contract: _BaseAdminUpgradeabilityProxy.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_BaseAdminUpgradeabilityProxy *BaseAdminUpgradeabilityProxyFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *BaseAdminUpgradeabilityProxyAdminChanged) (event.Subscription, error) {

	logs, sub, err := _BaseAdminUpgradeabilityProxy.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseAdminUpgradeabilityProxyAdminChanged)
				if err := _BaseAdminUpgradeabilityProxy.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_BaseAdminUpgradeabilityProxy *BaseAdminUpgradeabilityProxyFilterer) ParseAdminChanged(log types.Log) (*BaseAdminUpgradeabilityProxyAdminChanged, error) {
	event := new(BaseAdminUpgradeabilityProxyAdminChanged)
	if err := _BaseAdminUpgradeabilityProxy.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseAdminUpgradeabilityProxyUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the BaseAdminUpgradeabilityProxy contract.
type BaseAdminUpgradeabilityProxyUpgradedIterator struct {
	Event *BaseAdminUpgradeabilityProxyUpgraded // Event containing the contract specifics and raw log

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
func (it *BaseAdminUpgradeabilityProxyUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseAdminUpgradeabilityProxyUpgraded)
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
		it.Event = new(BaseAdminUpgradeabilityProxyUpgraded)
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
func (it *BaseAdminUpgradeabilityProxyUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseAdminUpgradeabilityProxyUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseAdminUpgradeabilityProxyUpgraded represents a Upgraded event raised by the BaseAdminUpgradeabilityProxy contract.
type BaseAdminUpgradeabilityProxyUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BaseAdminUpgradeabilityProxy *BaseAdminUpgradeabilityProxyFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BaseAdminUpgradeabilityProxyUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BaseAdminUpgradeabilityProxy.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BaseAdminUpgradeabilityProxyUpgradedIterator{contract: _BaseAdminUpgradeabilityProxy.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BaseAdminUpgradeabilityProxy *BaseAdminUpgradeabilityProxyFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BaseAdminUpgradeabilityProxyUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BaseAdminUpgradeabilityProxy.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseAdminUpgradeabilityProxyUpgraded)
				if err := _BaseAdminUpgradeabilityProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BaseAdminUpgradeabilityProxy *BaseAdminUpgradeabilityProxyFilterer) ParseUpgraded(log types.Log) (*BaseAdminUpgradeabilityProxyUpgraded, error) {
	event := new(BaseAdminUpgradeabilityProxyUpgraded)
	if err := _BaseAdminUpgradeabilityProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseUpgradeabilityProxyABI is the input ABI used to generate the binding from.
const BaseUpgradeabilityProxyABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// BaseUpgradeabilityProxyBin is the compiled bytecode used for deploying new contracts.
var BaseUpgradeabilityProxyBin = "0x6080604052348015600f57600080fd5b50609d8061001e6000396000f3fe6080604052600a600c565b005b6012601e565b601e601a6020565b6045565b565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5490565b3660008037600080366000845af43d6000803e8080156063573d6000f35b3d6000fdfea265627a7a7231582005187a82a4c7f3102c7401903756736bfd37c3aef03affe898d7196425a571f464736f6c63430005110032"

// DeployBaseUpgradeabilityProxy deploys a new Ethereum contract, binding an instance of BaseUpgradeabilityProxy to it.
func DeployBaseUpgradeabilityProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BaseUpgradeabilityProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseUpgradeabilityProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BaseUpgradeabilityProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BaseUpgradeabilityProxy{BaseUpgradeabilityProxyCaller: BaseUpgradeabilityProxyCaller{contract: contract}, BaseUpgradeabilityProxyTransactor: BaseUpgradeabilityProxyTransactor{contract: contract}, BaseUpgradeabilityProxyFilterer: BaseUpgradeabilityProxyFilterer{contract: contract}}, nil
}

// BaseUpgradeabilityProxy is an auto generated Go binding around an Ethereum contract.
type BaseUpgradeabilityProxy struct {
	BaseUpgradeabilityProxyCaller     // Read-only binding to the contract
	BaseUpgradeabilityProxyTransactor // Write-only binding to the contract
	BaseUpgradeabilityProxyFilterer   // Log filterer for contract events
}

// BaseUpgradeabilityProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseUpgradeabilityProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseUpgradeabilityProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseUpgradeabilityProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseUpgradeabilityProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseUpgradeabilityProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseUpgradeabilityProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseUpgradeabilityProxySession struct {
	Contract     *BaseUpgradeabilityProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BaseUpgradeabilityProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseUpgradeabilityProxyCallerSession struct {
	Contract *BaseUpgradeabilityProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// BaseUpgradeabilityProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseUpgradeabilityProxyTransactorSession struct {
	Contract     *BaseUpgradeabilityProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// BaseUpgradeabilityProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseUpgradeabilityProxyRaw struct {
	Contract *BaseUpgradeabilityProxy // Generic contract binding to access the raw methods on
}

// BaseUpgradeabilityProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseUpgradeabilityProxyCallerRaw struct {
	Contract *BaseUpgradeabilityProxyCaller // Generic read-only contract binding to access the raw methods on
}

// BaseUpgradeabilityProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseUpgradeabilityProxyTransactorRaw struct {
	Contract *BaseUpgradeabilityProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseUpgradeabilityProxy creates a new instance of BaseUpgradeabilityProxy, bound to a specific deployed contract.
func NewBaseUpgradeabilityProxy(address common.Address, backend bind.ContractBackend) (*BaseUpgradeabilityProxy, error) {
	contract, err := bindBaseUpgradeabilityProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseUpgradeabilityProxy{BaseUpgradeabilityProxyCaller: BaseUpgradeabilityProxyCaller{contract: contract}, BaseUpgradeabilityProxyTransactor: BaseUpgradeabilityProxyTransactor{contract: contract}, BaseUpgradeabilityProxyFilterer: BaseUpgradeabilityProxyFilterer{contract: contract}}, nil
}

// NewBaseUpgradeabilityProxyCaller creates a new read-only instance of BaseUpgradeabilityProxy, bound to a specific deployed contract.
func NewBaseUpgradeabilityProxyCaller(address common.Address, caller bind.ContractCaller) (*BaseUpgradeabilityProxyCaller, error) {
	contract, err := bindBaseUpgradeabilityProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseUpgradeabilityProxyCaller{contract: contract}, nil
}

// NewBaseUpgradeabilityProxyTransactor creates a new write-only instance of BaseUpgradeabilityProxy, bound to a specific deployed contract.
func NewBaseUpgradeabilityProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseUpgradeabilityProxyTransactor, error) {
	contract, err := bindBaseUpgradeabilityProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseUpgradeabilityProxyTransactor{contract: contract}, nil
}

// NewBaseUpgradeabilityProxyFilterer creates a new log filterer instance of BaseUpgradeabilityProxy, bound to a specific deployed contract.
func NewBaseUpgradeabilityProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseUpgradeabilityProxyFilterer, error) {
	contract, err := bindBaseUpgradeabilityProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseUpgradeabilityProxyFilterer{contract: contract}, nil
}

// bindBaseUpgradeabilityProxy binds a generic wrapper to an already deployed contract.
func bindBaseUpgradeabilityProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseUpgradeabilityProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseUpgradeabilityProxy *BaseUpgradeabilityProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseUpgradeabilityProxy.Contract.BaseUpgradeabilityProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseUpgradeabilityProxy *BaseUpgradeabilityProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseUpgradeabilityProxy.Contract.BaseUpgradeabilityProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseUpgradeabilityProxy *BaseUpgradeabilityProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseUpgradeabilityProxy.Contract.BaseUpgradeabilityProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseUpgradeabilityProxy *BaseUpgradeabilityProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseUpgradeabilityProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseUpgradeabilityProxy *BaseUpgradeabilityProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseUpgradeabilityProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseUpgradeabilityProxy *BaseUpgradeabilityProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseUpgradeabilityProxy.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BaseUpgradeabilityProxy *BaseUpgradeabilityProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _BaseUpgradeabilityProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BaseUpgradeabilityProxy *BaseUpgradeabilityProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BaseUpgradeabilityProxy.Contract.Fallback(&_BaseUpgradeabilityProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BaseUpgradeabilityProxy *BaseUpgradeabilityProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BaseUpgradeabilityProxy.Contract.Fallback(&_BaseUpgradeabilityProxy.TransactOpts, calldata)
}

// BaseUpgradeabilityProxyUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the BaseUpgradeabilityProxy contract.
type BaseUpgradeabilityProxyUpgradedIterator struct {
	Event *BaseUpgradeabilityProxyUpgraded // Event containing the contract specifics and raw log

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
func (it *BaseUpgradeabilityProxyUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseUpgradeabilityProxyUpgraded)
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
		it.Event = new(BaseUpgradeabilityProxyUpgraded)
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
func (it *BaseUpgradeabilityProxyUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseUpgradeabilityProxyUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseUpgradeabilityProxyUpgraded represents a Upgraded event raised by the BaseUpgradeabilityProxy contract.
type BaseUpgradeabilityProxyUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BaseUpgradeabilityProxy *BaseUpgradeabilityProxyFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BaseUpgradeabilityProxyUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BaseUpgradeabilityProxy.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BaseUpgradeabilityProxyUpgradedIterator{contract: _BaseUpgradeabilityProxy.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BaseUpgradeabilityProxy *BaseUpgradeabilityProxyFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BaseUpgradeabilityProxyUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BaseUpgradeabilityProxy.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseUpgradeabilityProxyUpgraded)
				if err := _BaseUpgradeabilityProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BaseUpgradeabilityProxy *BaseUpgradeabilityProxyFilterer) ParseUpgraded(log types.Log) (*BaseUpgradeabilityProxyUpgraded, error) {
	event := new(BaseUpgradeabilityProxyUpgraded)
	if err := _BaseUpgradeabilityProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CallerFactoryABI is the input ABI used to generate the binding from.
const CallerFactoryABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"initChildren\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"ProxyCreated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"deploy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"deploySigned\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"getDeploymentAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"getSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isChild\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CallerFactoryFuncSigs maps the 4-byte function signature to its string representation.
var CallerFactoryFuncSigs = map[string]string{
	"6150864c": "deploy(uint256,address,address,bytes)",
	"332d6626": "deploySigned(uint256,address,address,bytes,bytes)",
	"81ae1f5b": "getDeploymentAddress(uint256,address)",
	"290f8f56": "getSigner(uint256,address,address,bytes,bytes)",
	"fc91a897": "isChild(address)",
}

// CallerFactoryBin is the compiled bytecode used for deploying new contracts.
var CallerFactoryBin = "0x608060405234801561001057600080fd5b50604051611e94380380611e948339818101604052602081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b90830190602082018581111561006857600080fd5b825186602082028301116401000000008211171561008557600080fd5b82525081516020918201928201910280838360005b838110156100b257818101518382015260200161009a565b50505050905001604052505050604051806020016100cf9061014c565b601f1982820381018352601f9091011660405280516020919091012060009081555b815181101561014557600180600084848151811061010b57fe5b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff19169115159190911790556001016100f1565b5050610159565b61099c806114f883390190565b611390806101686000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063290f8f561461005c578063332d6626146101c15780636150864c1461030a57806381ae1f5b146103ce578063fc91a897146103fa575b600080fd5b6101a5600480360360a081101561007257600080fd5b8135916001600160a01b03602082013581169260408301359091169190810190608081016060820135600160201b8111156100ac57600080fd5b8201836020820111156100be57600080fd5b803590602001918460018302840111600160201b831117156100df57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561013157600080fd5b82018360208201111561014357600080fd5b803590602001918460018302840111600160201b8311171561016457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610434945050505050565b604080516001600160a01b039092168252519081900360200190f35b6101a5600480360360a08110156101d757600080fd5b8135916001600160a01b03602082013581169260408301359091169190810190608081016060820135600160201b81111561021157600080fd5b82018360208201111561022357600080fd5b803590602001918460018302840111600160201b8311171561024457600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561029657600080fd5b8201836020820111156102a857600080fd5b803590602001918460018302840111600160201b831117156102c957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610523945050505050565b6101a56004803603608081101561032057600080fd5b8135916001600160a01b03602082013581169260408301359091169190810190608081016060820135600160201b81111561035a57600080fd5b82018360208201111561036c57600080fd5b803590602001918460018302840111600160201b8311171561038d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610562945050505050565b6101a5600480360360408110156103e457600080fd5b50803590602001356001600160a01b031661059f565b6104206004803603602081101561041057600080fd5b50356001600160a01b0316610606565b604080519115158252519081900360200190f35b60008061050c878787873060405160200180868152602001856001600160a01b03166001600160a01b031660601b8152601401846001600160a01b03166001600160a01b031660601b815260140183805190602001908083835b602083106104ad5780518252601f19909201916020918201910161048e565b6001836020036101000a038019825116818451168082178552505050505050905001826001600160a01b03166001600160a01b031660601b81526014019550505050505060405160208183030381529060405280519060200120610624565b90506105188184610675565b979650505050505050565b60006105328686868686610763565b6001600160a01b0381166000908152600160208190526040909120805460ff191690911790559695505050505050565b6000610570858585856107d1565b6001600160a01b0381166000908152600160208190526040909120805460ff1916909117905595945050505050565b6000806105ac84846107e9565b600054604080516001600160f81b03196020808301919091523060601b6021830152603582019490945260558082019390935281518082039093018352607501905280519101206001600160a01b03169150505b92915050565b6001600160a01b031660009081526001602052604090205460ff1690565b604080517f19457468657265756d205369676e6564204d6573736167653a0a333200000000602080830191909152603c8083019490945282518083039094018452605c909101909152815191012090565b6000815160411461068857506000610600565b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08211156106ce5760009350505050610600565b8060ff16601b141580156106e657508060ff16601c14155b156106f75760009350505050610600565b6040805160008152602080820180845289905260ff8416828401526060820186905260808201859052915160019260a0808401939192601f1981019281900390910190855afa15801561074e573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b6000806107738787878787610434565b90506001600160a01b0381166107c4576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207369676e617475726560781b604482015290519081900360640190fd5b6105188787878785610828565b60006107e08585858533610828565b95945050505050565b6040805160208082019490945260609290921b6bffffffffffffffffffffffff1916828201528051808303603401815260549092019052805191012090565b6000806108358784610959565b604080516001600160a01b038316815290519192507efffc2da0b561cae30d9826d37709e9421c4725faebc226cbbb7ef5fc5e7349919081900360200190a160405163cf7a1d7760e01b81526001600160a01b038781166004830190815287821660248401526060604484019081528751606485015287519285169363cf7a1d77938b938b938b939192909160840190602085019080838360005b838110156108e85781810151838201526020016108d0565b50505050905090810190601f1680156109155780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b15801561093657600080fd5b505af115801561094a573d6000803e3d6000fd5b50929998505050505050505050565b60008060606040518060200161096e906109b2565b601f1982820381018352601f909101166040529050600061098f86866107e9565b9050808251602084016000f59250823b6109a857600080fd5b5090949350505050565b61099c806109c08339019056fe608060405234801561001057600080fd5b5061097c806100206000396000f3fe6080604052600436106100705760003560e01c80638f2839701161004e5780638f2839701461015e578063cf7a1d7714610191578063d1f5789414610250578063f851a4401461030657610070565b80633659cfe61461007a5780634f1ef286146100ad5780635c60da1b1461012d575b61007861031b565b005b34801561008657600080fd5b506100786004803603602081101561009d57600080fd5b50356001600160a01b0316610335565b610078600480360360408110156100c357600080fd5b6001600160a01b0382351691908101906040810160208201356401000000008111156100ee57600080fd5b82018360208201111561010057600080fd5b8035906020019184600183028401116401000000008311171561012257600080fd5b50909250905061036f565b34801561013957600080fd5b5061014261041c565b604080516001600160a01b039092168252519081900360200190f35b34801561016a57600080fd5b506100786004803603602081101561018157600080fd5b50356001600160a01b0316610459565b610078600480360360608110156101a757600080fd5b6001600160a01b0382358116926020810135909116918101906060810160408201356401000000008111156101db57600080fd5b8201836020820111156101ed57600080fd5b8035906020019184600183028401116401000000008311171561020f57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610513945050505050565b6100786004803603604081101561026657600080fd5b6001600160a01b03823516919081019060408101602082013564010000000081111561029157600080fd5b8201836020820111156102a357600080fd5b803590602001918460018302840111640100000000831117156102c557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610599945050505050565b34801561031257600080fd5b506101426106d9565b610323610704565b61033361032e610764565b610789565b565b61033d6107ad565b6001600160a01b0316336001600160a01b031614156103645761035f816107d2565b61036c565b61036c61031b565b50565b6103776107ad565b6001600160a01b0316336001600160a01b0316141561040f57610399836107d2565b6000836001600160a01b031683836040518083838082843760405192019450600093509091505080830381855af49150503d80600081146103f6576040519150601f19603f3d011682016040523d82523d6000602084013e6103fb565b606091505b505090508061040957600080fd5b50610417565b61041761031b565b505050565b60006104266107ad565b6001600160a01b0316336001600160a01b0316141561044e57610447610764565b9050610456565b61045661031b565b90565b6104616107ad565b6001600160a01b0316336001600160a01b03161415610364576001600160a01b0381166104bf5760405162461bcd60e51b81526004018080602001828103825260368152602001806108d76036913960400191505060405180910390fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6104e86107ad565b604080516001600160a01b03928316815291841660208301528051918290030190a161035f81610812565b600061051d610764565b6001600160a01b03161461053057600080fd5b61053a8382610599565b604080517232b4b8189c9b1b97383937bc3c9730b236b4b760691b815290519081900360130190207fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61036000199091011461059057fe5b61041782610812565b60006105a3610764565b6001600160a01b0316146105b657600080fd5b604080517f656970313936372e70726f78792e696d706c656d656e746174696f6e000000008152905190819003601c0190207f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6000199091011461061657fe5b61061f82610836565b8051156106d5576000826001600160a01b0316826040518082805190602001908083835b602083106106625780518252601f199092019160209182019101610643565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d80600081146106c2576040519150601f19603f3d011682016040523d82523d6000602084013e6106c7565b606091505b505090508061041757600080fd5b5050565b60006106e36107ad565b6001600160a01b0316336001600160a01b0316141561044e576104476107ad565b61070c6107ad565b6001600160a01b0316336001600160a01b0316141561075c5760405162461bcd60e51b81526004018080602001828103825260328152602001806108a56032913960400191505060405180910390fd5b610333610333565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5490565b3660008037600080366000845af43d6000803e8080156107a8573d6000f35b3d6000fd5b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035490565b6107db81610836565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610355565b61083f8161089e565b61087a5760405162461bcd60e51b815260040180806020018281038252603b81526020018061090d603b913960400191505060405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc55565b3b15159056fe43616e6e6f742063616c6c2066616c6c6261636b2066756e6374696f6e2066726f6d207468652070726f78792061646d696e43616e6e6f74206368616e6765207468652061646d696e206f6620612070726f787920746f20746865207a65726f206164647265737343616e6e6f742073657420612070726f787920696d706c656d656e746174696f6e20746f2061206e6f6e2d636f6e74726163742061646472657373a265627a7a72315820ea63902a998581c26cce013e1736a0cabdc6bc56ff51f923fdb0a60bef5645e364736f6c63430005110032a265627a7a723158200d698f2732a1d72de170306285a5b00d33d92520b929b133ba1eb1119613fd9464736f6c63430005110032608060405234801561001057600080fd5b5061097c806100206000396000f3fe6080604052600436106100705760003560e01c80638f2839701161004e5780638f2839701461015e578063cf7a1d7714610191578063d1f5789414610250578063f851a4401461030657610070565b80633659cfe61461007a5780634f1ef286146100ad5780635c60da1b1461012d575b61007861031b565b005b34801561008657600080fd5b506100786004803603602081101561009d57600080fd5b50356001600160a01b0316610335565b610078600480360360408110156100c357600080fd5b6001600160a01b0382351691908101906040810160208201356401000000008111156100ee57600080fd5b82018360208201111561010057600080fd5b8035906020019184600183028401116401000000008311171561012257600080fd5b50909250905061036f565b34801561013957600080fd5b5061014261041c565b604080516001600160a01b039092168252519081900360200190f35b34801561016a57600080fd5b506100786004803603602081101561018157600080fd5b50356001600160a01b0316610459565b610078600480360360608110156101a757600080fd5b6001600160a01b0382358116926020810135909116918101906060810160408201356401000000008111156101db57600080fd5b8201836020820111156101ed57600080fd5b8035906020019184600183028401116401000000008311171561020f57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610513945050505050565b6100786004803603604081101561026657600080fd5b6001600160a01b03823516919081019060408101602082013564010000000081111561029157600080fd5b8201836020820111156102a357600080fd5b803590602001918460018302840111640100000000831117156102c557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610599945050505050565b34801561031257600080fd5b506101426106d9565b610323610704565b61033361032e610764565b610789565b565b61033d6107ad565b6001600160a01b0316336001600160a01b031614156103645761035f816107d2565b61036c565b61036c61031b565b50565b6103776107ad565b6001600160a01b0316336001600160a01b0316141561040f57610399836107d2565b6000836001600160a01b031683836040518083838082843760405192019450600093509091505080830381855af49150503d80600081146103f6576040519150601f19603f3d011682016040523d82523d6000602084013e6103fb565b606091505b505090508061040957600080fd5b50610417565b61041761031b565b505050565b60006104266107ad565b6001600160a01b0316336001600160a01b0316141561044e57610447610764565b9050610456565b61045661031b565b90565b6104616107ad565b6001600160a01b0316336001600160a01b03161415610364576001600160a01b0381166104bf5760405162461bcd60e51b81526004018080602001828103825260368152602001806108d76036913960400191505060405180910390fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6104e86107ad565b604080516001600160a01b03928316815291841660208301528051918290030190a161035f81610812565b600061051d610764565b6001600160a01b03161461053057600080fd5b61053a8382610599565b604080517232b4b8189c9b1b97383937bc3c9730b236b4b760691b815290519081900360130190207fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61036000199091011461059057fe5b61041782610812565b60006105a3610764565b6001600160a01b0316146105b657600080fd5b604080517f656970313936372e70726f78792e696d706c656d656e746174696f6e000000008152905190819003601c0190207f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6000199091011461061657fe5b61061f82610836565b8051156106d5576000826001600160a01b0316826040518082805190602001908083835b602083106106625780518252601f199092019160209182019101610643565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d80600081146106c2576040519150601f19603f3d011682016040523d82523d6000602084013e6106c7565b606091505b505090508061041757600080fd5b5050565b60006106e36107ad565b6001600160a01b0316336001600160a01b0316141561044e576104476107ad565b61070c6107ad565b6001600160a01b0316336001600160a01b0316141561075c5760405162461bcd60e51b81526004018080602001828103825260328152602001806108a56032913960400191505060405180910390fd5b610333610333565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5490565b3660008037600080366000845af43d6000803e8080156107a8573d6000f35b3d6000fd5b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035490565b6107db81610836565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610355565b61083f8161089e565b61087a5760405162461bcd60e51b815260040180806020018281038252603b81526020018061090d603b913960400191505060405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc55565b3b15159056fe43616e6e6f742063616c6c2066616c6c6261636b2066756e6374696f6e2066726f6d207468652070726f78792061646d696e43616e6e6f74206368616e6765207468652061646d696e206f6620612070726f787920746f20746865207a65726f206164647265737343616e6e6f742073657420612070726f787920696d706c656d656e746174696f6e20746f2061206e6f6e2d636f6e74726163742061646472657373a265627a7a72315820ea63902a998581c26cce013e1736a0cabdc6bc56ff51f923fdb0a60bef5645e364736f6c63430005110032"

// DeployCallerFactory deploys a new Ethereum contract, binding an instance of CallerFactory to it.
func DeployCallerFactory(auth *bind.TransactOpts, backend bind.ContractBackend, initChildren []common.Address) (common.Address, *types.Transaction, *CallerFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(CallerFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CallerFactoryBin), backend, initChildren)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CallerFactory{CallerFactoryCaller: CallerFactoryCaller{contract: contract}, CallerFactoryTransactor: CallerFactoryTransactor{contract: contract}, CallerFactoryFilterer: CallerFactoryFilterer{contract: contract}}, nil
}

// CallerFactory is an auto generated Go binding around an Ethereum contract.
type CallerFactory struct {
	CallerFactoryCaller     // Read-only binding to the contract
	CallerFactoryTransactor // Write-only binding to the contract
	CallerFactoryFilterer   // Log filterer for contract events
}

// CallerFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CallerFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallerFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CallerFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallerFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CallerFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallerFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CallerFactorySession struct {
	Contract     *CallerFactory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CallerFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CallerFactoryCallerSession struct {
	Contract *CallerFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CallerFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CallerFactoryTransactorSession struct {
	Contract     *CallerFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CallerFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CallerFactoryRaw struct {
	Contract *CallerFactory // Generic contract binding to access the raw methods on
}

// CallerFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CallerFactoryCallerRaw struct {
	Contract *CallerFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// CallerFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CallerFactoryTransactorRaw struct {
	Contract *CallerFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCallerFactory creates a new instance of CallerFactory, bound to a specific deployed contract.
func NewCallerFactory(address common.Address, backend bind.ContractBackend) (*CallerFactory, error) {
	contract, err := bindCallerFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CallerFactory{CallerFactoryCaller: CallerFactoryCaller{contract: contract}, CallerFactoryTransactor: CallerFactoryTransactor{contract: contract}, CallerFactoryFilterer: CallerFactoryFilterer{contract: contract}}, nil
}

// NewCallerFactoryCaller creates a new read-only instance of CallerFactory, bound to a specific deployed contract.
func NewCallerFactoryCaller(address common.Address, caller bind.ContractCaller) (*CallerFactoryCaller, error) {
	contract, err := bindCallerFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CallerFactoryCaller{contract: contract}, nil
}

// NewCallerFactoryTransactor creates a new write-only instance of CallerFactory, bound to a specific deployed contract.
func NewCallerFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*CallerFactoryTransactor, error) {
	contract, err := bindCallerFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CallerFactoryTransactor{contract: contract}, nil
}

// NewCallerFactoryFilterer creates a new log filterer instance of CallerFactory, bound to a specific deployed contract.
func NewCallerFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*CallerFactoryFilterer, error) {
	contract, err := bindCallerFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CallerFactoryFilterer{contract: contract}, nil
}

// bindCallerFactory binds a generic wrapper to an already deployed contract.
func bindCallerFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CallerFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CallerFactory *CallerFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CallerFactory.Contract.CallerFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CallerFactory *CallerFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CallerFactory.Contract.CallerFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CallerFactory *CallerFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CallerFactory.Contract.CallerFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CallerFactory *CallerFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CallerFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CallerFactory *CallerFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CallerFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CallerFactory *CallerFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CallerFactory.Contract.contract.Transact(opts, method, params...)
}

// GetDeploymentAddress is a free data retrieval call binding the contract method 0x81ae1f5b.
//
// Solidity: function getDeploymentAddress(uint256 _salt, address _sender) view returns(address)
func (_CallerFactory *CallerFactoryCaller) GetDeploymentAddress(opts *bind.CallOpts, _salt *big.Int, _sender common.Address) (common.Address, error) {
	var out []interface{}
	err := _CallerFactory.contract.Call(opts, &out, "getDeploymentAddress", _salt, _sender)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDeploymentAddress is a free data retrieval call binding the contract method 0x81ae1f5b.
//
// Solidity: function getDeploymentAddress(uint256 _salt, address _sender) view returns(address)
func (_CallerFactory *CallerFactorySession) GetDeploymentAddress(_salt *big.Int, _sender common.Address) (common.Address, error) {
	return _CallerFactory.Contract.GetDeploymentAddress(&_CallerFactory.CallOpts, _salt, _sender)
}

// GetDeploymentAddress is a free data retrieval call binding the contract method 0x81ae1f5b.
//
// Solidity: function getDeploymentAddress(uint256 _salt, address _sender) view returns(address)
func (_CallerFactory *CallerFactoryCallerSession) GetDeploymentAddress(_salt *big.Int, _sender common.Address) (common.Address, error) {
	return _CallerFactory.Contract.GetDeploymentAddress(&_CallerFactory.CallOpts, _salt, _sender)
}

// GetSigner is a free data retrieval call binding the contract method 0x290f8f56.
//
// Solidity: function getSigner(uint256 _salt, address _logic, address _admin, bytes _data, bytes _signature) view returns(address)
func (_CallerFactory *CallerFactoryCaller) GetSigner(opts *bind.CallOpts, _salt *big.Int, _logic common.Address, _admin common.Address, _data []byte, _signature []byte) (common.Address, error) {
	var out []interface{}
	err := _CallerFactory.contract.Call(opts, &out, "getSigner", _salt, _logic, _admin, _data, _signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSigner is a free data retrieval call binding the contract method 0x290f8f56.
//
// Solidity: function getSigner(uint256 _salt, address _logic, address _admin, bytes _data, bytes _signature) view returns(address)
func (_CallerFactory *CallerFactorySession) GetSigner(_salt *big.Int, _logic common.Address, _admin common.Address, _data []byte, _signature []byte) (common.Address, error) {
	return _CallerFactory.Contract.GetSigner(&_CallerFactory.CallOpts, _salt, _logic, _admin, _data, _signature)
}

// GetSigner is a free data retrieval call binding the contract method 0x290f8f56.
//
// Solidity: function getSigner(uint256 _salt, address _logic, address _admin, bytes _data, bytes _signature) view returns(address)
func (_CallerFactory *CallerFactoryCallerSession) GetSigner(_salt *big.Int, _logic common.Address, _admin common.Address, _data []byte, _signature []byte) (common.Address, error) {
	return _CallerFactory.Contract.GetSigner(&_CallerFactory.CallOpts, _salt, _logic, _admin, _data, _signature)
}

// IsChild is a free data retrieval call binding the contract method 0xfc91a897.
//
// Solidity: function isChild(address _addr) view returns(bool)
func (_CallerFactory *CallerFactoryCaller) IsChild(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _CallerFactory.contract.Call(opts, &out, "isChild", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsChild is a free data retrieval call binding the contract method 0xfc91a897.
//
// Solidity: function isChild(address _addr) view returns(bool)
func (_CallerFactory *CallerFactorySession) IsChild(_addr common.Address) (bool, error) {
	return _CallerFactory.Contract.IsChild(&_CallerFactory.CallOpts, _addr)
}

// IsChild is a free data retrieval call binding the contract method 0xfc91a897.
//
// Solidity: function isChild(address _addr) view returns(bool)
func (_CallerFactory *CallerFactoryCallerSession) IsChild(_addr common.Address) (bool, error) {
	return _CallerFactory.Contract.IsChild(&_CallerFactory.CallOpts, _addr)
}

// Deploy is a paid mutator transaction binding the contract method 0x6150864c.
//
// Solidity: function deploy(uint256 _salt, address _logic, address _admin, bytes _data) returns(address proxy)
func (_CallerFactory *CallerFactoryTransactor) Deploy(opts *bind.TransactOpts, _salt *big.Int, _logic common.Address, _admin common.Address, _data []byte) (*types.Transaction, error) {
	return _CallerFactory.contract.Transact(opts, "deploy", _salt, _logic, _admin, _data)
}

// Deploy is a paid mutator transaction binding the contract method 0x6150864c.
//
// Solidity: function deploy(uint256 _salt, address _logic, address _admin, bytes _data) returns(address proxy)
func (_CallerFactory *CallerFactorySession) Deploy(_salt *big.Int, _logic common.Address, _admin common.Address, _data []byte) (*types.Transaction, error) {
	return _CallerFactory.Contract.Deploy(&_CallerFactory.TransactOpts, _salt, _logic, _admin, _data)
}

// Deploy is a paid mutator transaction binding the contract method 0x6150864c.
//
// Solidity: function deploy(uint256 _salt, address _logic, address _admin, bytes _data) returns(address proxy)
func (_CallerFactory *CallerFactoryTransactorSession) Deploy(_salt *big.Int, _logic common.Address, _admin common.Address, _data []byte) (*types.Transaction, error) {
	return _CallerFactory.Contract.Deploy(&_CallerFactory.TransactOpts, _salt, _logic, _admin, _data)
}

// DeploySigned is a paid mutator transaction binding the contract method 0x332d6626.
//
// Solidity: function deploySigned(uint256 _salt, address _logic, address _admin, bytes _data, bytes _signature) returns(address proxy)
func (_CallerFactory *CallerFactoryTransactor) DeploySigned(opts *bind.TransactOpts, _salt *big.Int, _logic common.Address, _admin common.Address, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _CallerFactory.contract.Transact(opts, "deploySigned", _salt, _logic, _admin, _data, _signature)
}

// DeploySigned is a paid mutator transaction binding the contract method 0x332d6626.
//
// Solidity: function deploySigned(uint256 _salt, address _logic, address _admin, bytes _data, bytes _signature) returns(address proxy)
func (_CallerFactory *CallerFactorySession) DeploySigned(_salt *big.Int, _logic common.Address, _admin common.Address, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _CallerFactory.Contract.DeploySigned(&_CallerFactory.TransactOpts, _salt, _logic, _admin, _data, _signature)
}

// DeploySigned is a paid mutator transaction binding the contract method 0x332d6626.
//
// Solidity: function deploySigned(uint256 _salt, address _logic, address _admin, bytes _data, bytes _signature) returns(address proxy)
func (_CallerFactory *CallerFactoryTransactorSession) DeploySigned(_salt *big.Int, _logic common.Address, _admin common.Address, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _CallerFactory.Contract.DeploySigned(&_CallerFactory.TransactOpts, _salt, _logic, _admin, _data, _signature)
}

// CallerFactoryProxyCreatedIterator is returned from FilterProxyCreated and is used to iterate over the raw logs and unpacked data for ProxyCreated events raised by the CallerFactory contract.
type CallerFactoryProxyCreatedIterator struct {
	Event *CallerFactoryProxyCreated // Event containing the contract specifics and raw log

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
func (it *CallerFactoryProxyCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CallerFactoryProxyCreated)
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
		it.Event = new(CallerFactoryProxyCreated)
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
func (it *CallerFactoryProxyCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CallerFactoryProxyCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CallerFactoryProxyCreated represents a ProxyCreated event raised by the CallerFactory contract.
type CallerFactoryProxyCreated struct {
	Proxy common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterProxyCreated is a free log retrieval operation binding the contract event 0x00fffc2da0b561cae30d9826d37709e9421c4725faebc226cbbb7ef5fc5e7349.
//
// Solidity: event ProxyCreated(address proxy)
func (_CallerFactory *CallerFactoryFilterer) FilterProxyCreated(opts *bind.FilterOpts) (*CallerFactoryProxyCreatedIterator, error) {

	logs, sub, err := _CallerFactory.contract.FilterLogs(opts, "ProxyCreated")
	if err != nil {
		return nil, err
	}
	return &CallerFactoryProxyCreatedIterator{contract: _CallerFactory.contract, event: "ProxyCreated", logs: logs, sub: sub}, nil
}

// WatchProxyCreated is a free log subscription operation binding the contract event 0x00fffc2da0b561cae30d9826d37709e9421c4725faebc226cbbb7ef5fc5e7349.
//
// Solidity: event ProxyCreated(address proxy)
func (_CallerFactory *CallerFactoryFilterer) WatchProxyCreated(opts *bind.WatchOpts, sink chan<- *CallerFactoryProxyCreated) (event.Subscription, error) {

	logs, sub, err := _CallerFactory.contract.WatchLogs(opts, "ProxyCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CallerFactoryProxyCreated)
				if err := _CallerFactory.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
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

// ParseProxyCreated is a log parse operation binding the contract event 0x00fffc2da0b561cae30d9826d37709e9421c4725faebc226cbbb7ef5fc5e7349.
//
// Solidity: event ProxyCreated(address proxy)
func (_CallerFactory *CallerFactoryFilterer) ParseProxyCreated(log types.Log) (*CallerFactoryProxyCreated, error) {
	event := new(CallerFactoryProxyCreated)
	if err := _CallerFactory.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConstABI is the input ABI used to generate the binding from.
const ConstABI = "[]"

// ConstBin is the compiled bytecode used for deploying new contracts.
var ConstBin = "0x6080604052348015600f57600080fd5b50603e80601d6000396000f3fe6080604052600080fdfea265627a7a72315820265fffa605c3c93da3fadd86f2c408a81763fdd83982e7e2774e43f3ee93701b64736f6c63430005110032"

// DeployConst deploys a new Ethereum contract, binding an instance of Const to it.
func DeployConst(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Const, error) {
	parsed, err := abi.JSON(strings.NewReader(ConstABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ConstBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Const{ConstCaller: ConstCaller{contract: contract}, ConstTransactor: ConstTransactor{contract: contract}, ConstFilterer: ConstFilterer{contract: contract}}, nil
}

// Const is an auto generated Go binding around an Ethereum contract.
type Const struct {
	ConstCaller     // Read-only binding to the contract
	ConstTransactor // Write-only binding to the contract
	ConstFilterer   // Log filterer for contract events
}

// ConstCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConstCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConstTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConstFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConstSession struct {
	Contract     *Const            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConstCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConstCallerSession struct {
	Contract *ConstCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ConstTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConstTransactorSession struct {
	Contract     *ConstTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConstRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConstRaw struct {
	Contract *Const // Generic contract binding to access the raw methods on
}

// ConstCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConstCallerRaw struct {
	Contract *ConstCaller // Generic read-only contract binding to access the raw methods on
}

// ConstTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConstTransactorRaw struct {
	Contract *ConstTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConst creates a new instance of Const, bound to a specific deployed contract.
func NewConst(address common.Address, backend bind.ContractBackend) (*Const, error) {
	contract, err := bindConst(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Const{ConstCaller: ConstCaller{contract: contract}, ConstTransactor: ConstTransactor{contract: contract}, ConstFilterer: ConstFilterer{contract: contract}}, nil
}

// NewConstCaller creates a new read-only instance of Const, bound to a specific deployed contract.
func NewConstCaller(address common.Address, caller bind.ContractCaller) (*ConstCaller, error) {
	contract, err := bindConst(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConstCaller{contract: contract}, nil
}

// NewConstTransactor creates a new write-only instance of Const, bound to a specific deployed contract.
func NewConstTransactor(address common.Address, transactor bind.ContractTransactor) (*ConstTransactor, error) {
	contract, err := bindConst(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConstTransactor{contract: contract}, nil
}

// NewConstFilterer creates a new log filterer instance of Const, bound to a specific deployed contract.
func NewConstFilterer(address common.Address, filterer bind.ContractFilterer) (*ConstFilterer, error) {
	contract, err := bindConst(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConstFilterer{contract: contract}, nil
}

// bindConst binds a generic wrapper to an already deployed contract.
func bindConst(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConstABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Const *ConstRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Const.Contract.ConstCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Const *ConstRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Const.Contract.ConstTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Const *ConstRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Const.Contract.ConstTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Const *ConstCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Const.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Const *ConstTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Const.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Const *ConstTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Const.Contract.contract.Transact(opts, method, params...)
}

// ContextABI is the input ABI used to generate the binding from.
const ContextABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// ECCUtilsABI is the input ABI used to generate the binding from.
const ECCUtilsABI = "[]"

// ECCUtilsBin is the compiled bytecode used for deploying new contracts.
var ECCUtilsBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820735fda8c05c9d5b4eafeda129ab838fdd22d4764e42b759d6d43abc0c977129564736f6c63430005110032"

// DeployECCUtils deploys a new Ethereum contract, binding an instance of ECCUtils to it.
func DeployECCUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECCUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(ECCUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ECCUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECCUtils{ECCUtilsCaller: ECCUtilsCaller{contract: contract}, ECCUtilsTransactor: ECCUtilsTransactor{contract: contract}, ECCUtilsFilterer: ECCUtilsFilterer{contract: contract}}, nil
}

// ECCUtils is an auto generated Go binding around an Ethereum contract.
type ECCUtils struct {
	ECCUtilsCaller     // Read-only binding to the contract
	ECCUtilsTransactor // Write-only binding to the contract
	ECCUtilsFilterer   // Log filterer for contract events
}

// ECCUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ECCUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECCUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECCUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECCUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECCUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECCUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECCUtilsSession struct {
	Contract     *ECCUtils         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECCUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECCUtilsCallerSession struct {
	Contract *ECCUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ECCUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECCUtilsTransactorSession struct {
	Contract     *ECCUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ECCUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ECCUtilsRaw struct {
	Contract *ECCUtils // Generic contract binding to access the raw methods on
}

// ECCUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECCUtilsCallerRaw struct {
	Contract *ECCUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// ECCUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECCUtilsTransactorRaw struct {
	Contract *ECCUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewECCUtils creates a new instance of ECCUtils, bound to a specific deployed contract.
func NewECCUtils(address common.Address, backend bind.ContractBackend) (*ECCUtils, error) {
	contract, err := bindECCUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECCUtils{ECCUtilsCaller: ECCUtilsCaller{contract: contract}, ECCUtilsTransactor: ECCUtilsTransactor{contract: contract}, ECCUtilsFilterer: ECCUtilsFilterer{contract: contract}}, nil
}

// NewECCUtilsCaller creates a new read-only instance of ECCUtils, bound to a specific deployed contract.
func NewECCUtilsCaller(address common.Address, caller bind.ContractCaller) (*ECCUtilsCaller, error) {
	contract, err := bindECCUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECCUtilsCaller{contract: contract}, nil
}

// NewECCUtilsTransactor creates a new write-only instance of ECCUtils, bound to a specific deployed contract.
func NewECCUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*ECCUtilsTransactor, error) {
	contract, err := bindECCUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECCUtilsTransactor{contract: contract}, nil
}

// NewECCUtilsFilterer creates a new log filterer instance of ECCUtils, bound to a specific deployed contract.
func NewECCUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*ECCUtilsFilterer, error) {
	contract, err := bindECCUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECCUtilsFilterer{contract: contract}, nil
}

// bindECCUtils binds a generic wrapper to an already deployed contract.
func bindECCUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECCUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECCUtils *ECCUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECCUtils.Contract.ECCUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECCUtils *ECCUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECCUtils.Contract.ECCUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECCUtils *ECCUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECCUtils.Contract.ECCUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECCUtils *ECCUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECCUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECCUtils *ECCUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECCUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECCUtils *ECCUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECCUtils.Contract.contract.Transact(opts, method, params...)
}

// EthCrossChainCallerABI is the input ABI used to generate the binding from.
const EthCrossChainCallerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// EthCrossChainCallerFuncSigs maps the 4-byte function signature to its string representation.
var EthCrossChainCallerFuncSigs = map[string]string{
	"f851a440": "admin()",
	"8f283970": "changeAdmin(address)",
	"5c60da1b": "implementation()",
	"cf7a1d77": "initialize(address,address,bytes)",
	"d1f57894": "initialize(address,bytes)",
	"3659cfe6": "upgradeTo(address)",
	"4f1ef286": "upgradeToAndCall(address,bytes)",
}

// EthCrossChainCallerBin is the compiled bytecode used for deploying new contracts.
var EthCrossChainCallerBin = "0x608060405234801561001057600080fd5b5061097c806100206000396000f3fe6080604052600436106100705760003560e01c80638f2839701161004e5780638f2839701461015e578063cf7a1d7714610191578063d1f5789414610250578063f851a4401461030657610070565b80633659cfe61461007a5780634f1ef286146100ad5780635c60da1b1461012d575b61007861031b565b005b34801561008657600080fd5b506100786004803603602081101561009d57600080fd5b50356001600160a01b0316610335565b610078600480360360408110156100c357600080fd5b6001600160a01b0382351691908101906040810160208201356401000000008111156100ee57600080fd5b82018360208201111561010057600080fd5b8035906020019184600183028401116401000000008311171561012257600080fd5b50909250905061036f565b34801561013957600080fd5b5061014261041c565b604080516001600160a01b039092168252519081900360200190f35b34801561016a57600080fd5b506100786004803603602081101561018157600080fd5b50356001600160a01b0316610459565b610078600480360360608110156101a757600080fd5b6001600160a01b0382358116926020810135909116918101906060810160408201356401000000008111156101db57600080fd5b8201836020820111156101ed57600080fd5b8035906020019184600183028401116401000000008311171561020f57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610513945050505050565b6100786004803603604081101561026657600080fd5b6001600160a01b03823516919081019060408101602082013564010000000081111561029157600080fd5b8201836020820111156102a357600080fd5b803590602001918460018302840111640100000000831117156102c557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610599945050505050565b34801561031257600080fd5b506101426106d9565b610323610704565b61033361032e610764565b610789565b565b61033d6107ad565b6001600160a01b0316336001600160a01b031614156103645761035f816107d2565b61036c565b61036c61031b565b50565b6103776107ad565b6001600160a01b0316336001600160a01b0316141561040f57610399836107d2565b6000836001600160a01b031683836040518083838082843760405192019450600093509091505080830381855af49150503d80600081146103f6576040519150601f19603f3d011682016040523d82523d6000602084013e6103fb565b606091505b505090508061040957600080fd5b50610417565b61041761031b565b505050565b60006104266107ad565b6001600160a01b0316336001600160a01b0316141561044e57610447610764565b9050610456565b61045661031b565b90565b6104616107ad565b6001600160a01b0316336001600160a01b03161415610364576001600160a01b0381166104bf5760405162461bcd60e51b81526004018080602001828103825260368152602001806108d76036913960400191505060405180910390fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6104e86107ad565b604080516001600160a01b03928316815291841660208301528051918290030190a161035f81610812565b600061051d610764565b6001600160a01b03161461053057600080fd5b61053a8382610599565b604080517232b4b8189c9b1b97383937bc3c9730b236b4b760691b815290519081900360130190207fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61036000199091011461059057fe5b61041782610812565b60006105a3610764565b6001600160a01b0316146105b657600080fd5b604080517f656970313936372e70726f78792e696d706c656d656e746174696f6e000000008152905190819003601c0190207f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6000199091011461061657fe5b61061f82610836565b8051156106d5576000826001600160a01b0316826040518082805190602001908083835b602083106106625780518252601f199092019160209182019101610643565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d80600081146106c2576040519150601f19603f3d011682016040523d82523d6000602084013e6106c7565b606091505b505090508061041757600080fd5b5050565b60006106e36107ad565b6001600160a01b0316336001600160a01b0316141561044e576104476107ad565b61070c6107ad565b6001600160a01b0316336001600160a01b0316141561075c5760405162461bcd60e51b81526004018080602001828103825260328152602001806108a56032913960400191505060405180910390fd5b610333610333565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5490565b3660008037600080366000845af43d6000803e8080156107a8573d6000f35b3d6000fd5b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035490565b6107db81610836565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610355565b61083f8161089e565b61087a5760405162461bcd60e51b815260040180806020018281038252603b81526020018061090d603b913960400191505060405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc55565b3b15159056fe43616e6e6f742063616c6c2066616c6c6261636b2066756e6374696f6e2066726f6d207468652070726f78792061646d696e43616e6e6f74206368616e6765207468652061646d696e206f6620612070726f787920746f20746865207a65726f206164647265737343616e6e6f742073657420612070726f787920696d706c656d656e746174696f6e20746f2061206e6f6e2d636f6e74726163742061646472657373a265627a7a72315820a920775fe4c469f7cf4527b942cba938bd2ff3e6cb82f55ff934f2a81d33b56964736f6c63430005110032"

// DeployEthCrossChainCaller deploys a new Ethereum contract, binding an instance of EthCrossChainCaller to it.
func DeployEthCrossChainCaller(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EthCrossChainCaller, error) {
	parsed, err := abi.JSON(strings.NewReader(EthCrossChainCallerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EthCrossChainCallerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthCrossChainCaller{EthCrossChainCallerCaller: EthCrossChainCallerCaller{contract: contract}, EthCrossChainCallerTransactor: EthCrossChainCallerTransactor{contract: contract}, EthCrossChainCallerFilterer: EthCrossChainCallerFilterer{contract: contract}}, nil
}

// EthCrossChainCaller is an auto generated Go binding around an Ethereum contract.
type EthCrossChainCaller struct {
	EthCrossChainCallerCaller     // Read-only binding to the contract
	EthCrossChainCallerTransactor // Write-only binding to the contract
	EthCrossChainCallerFilterer   // Log filterer for contract events
}

// EthCrossChainCallerCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthCrossChainCallerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthCrossChainCallerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthCrossChainCallerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthCrossChainCallerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthCrossChainCallerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthCrossChainCallerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthCrossChainCallerSession struct {
	Contract     *EthCrossChainCaller // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EthCrossChainCallerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthCrossChainCallerCallerSession struct {
	Contract *EthCrossChainCallerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// EthCrossChainCallerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthCrossChainCallerTransactorSession struct {
	Contract     *EthCrossChainCallerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// EthCrossChainCallerRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthCrossChainCallerRaw struct {
	Contract *EthCrossChainCaller // Generic contract binding to access the raw methods on
}

// EthCrossChainCallerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthCrossChainCallerCallerRaw struct {
	Contract *EthCrossChainCallerCaller // Generic read-only contract binding to access the raw methods on
}

// EthCrossChainCallerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthCrossChainCallerTransactorRaw struct {
	Contract *EthCrossChainCallerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthCrossChainCaller creates a new instance of EthCrossChainCaller, bound to a specific deployed contract.
func NewEthCrossChainCaller(address common.Address, backend bind.ContractBackend) (*EthCrossChainCaller, error) {
	contract, err := bindEthCrossChainCaller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainCaller{EthCrossChainCallerCaller: EthCrossChainCallerCaller{contract: contract}, EthCrossChainCallerTransactor: EthCrossChainCallerTransactor{contract: contract}, EthCrossChainCallerFilterer: EthCrossChainCallerFilterer{contract: contract}}, nil
}

// NewEthCrossChainCallerCaller creates a new read-only instance of EthCrossChainCaller, bound to a specific deployed contract.
func NewEthCrossChainCallerCaller(address common.Address, caller bind.ContractCaller) (*EthCrossChainCallerCaller, error) {
	contract, err := bindEthCrossChainCaller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainCallerCaller{contract: contract}, nil
}

// NewEthCrossChainCallerTransactor creates a new write-only instance of EthCrossChainCaller, bound to a specific deployed contract.
func NewEthCrossChainCallerTransactor(address common.Address, transactor bind.ContractTransactor) (*EthCrossChainCallerTransactor, error) {
	contract, err := bindEthCrossChainCaller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainCallerTransactor{contract: contract}, nil
}

// NewEthCrossChainCallerFilterer creates a new log filterer instance of EthCrossChainCaller, bound to a specific deployed contract.
func NewEthCrossChainCallerFilterer(address common.Address, filterer bind.ContractFilterer) (*EthCrossChainCallerFilterer, error) {
	contract, err := bindEthCrossChainCaller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainCallerFilterer{contract: contract}, nil
}

// bindEthCrossChainCaller binds a generic wrapper to an already deployed contract.
func bindEthCrossChainCaller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthCrossChainCallerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthCrossChainCaller *EthCrossChainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthCrossChainCaller.Contract.EthCrossChainCallerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthCrossChainCaller *EthCrossChainCallerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainCaller.Contract.EthCrossChainCallerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthCrossChainCaller *EthCrossChainCallerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthCrossChainCaller.Contract.EthCrossChainCallerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthCrossChainCaller *EthCrossChainCallerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthCrossChainCaller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthCrossChainCaller *EthCrossChainCallerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainCaller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthCrossChainCaller *EthCrossChainCallerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthCrossChainCaller.Contract.contract.Transact(opts, method, params...)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_EthCrossChainCaller *EthCrossChainCallerTransactor) Admin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainCaller.contract.Transact(opts, "admin")
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_EthCrossChainCaller *EthCrossChainCallerSession) Admin() (*types.Transaction, error) {
	return _EthCrossChainCaller.Contract.Admin(&_EthCrossChainCaller.TransactOpts)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_EthCrossChainCaller *EthCrossChainCallerTransactorSession) Admin() (*types.Transaction, error) {
	return _EthCrossChainCaller.Contract.Admin(&_EthCrossChainCaller.TransactOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_EthCrossChainCaller *EthCrossChainCallerTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _EthCrossChainCaller.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_EthCrossChainCaller *EthCrossChainCallerSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EthCrossChainCaller.Contract.ChangeAdmin(&_EthCrossChainCaller.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_EthCrossChainCaller *EthCrossChainCallerTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EthCrossChainCaller.Contract.ChangeAdmin(&_EthCrossChainCaller.TransactOpts, newAdmin)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_EthCrossChainCaller *EthCrossChainCallerTransactor) Implementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainCaller.contract.Transact(opts, "implementation")
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_EthCrossChainCaller *EthCrossChainCallerSession) Implementation() (*types.Transaction, error) {
	return _EthCrossChainCaller.Contract.Implementation(&_EthCrossChainCaller.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_EthCrossChainCaller *EthCrossChainCallerTransactorSession) Implementation() (*types.Transaction, error) {
	return _EthCrossChainCaller.Contract.Implementation(&_EthCrossChainCaller.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf7a1d77.
//
// Solidity: function initialize(address _logic, address _admin, bytes _data) payable returns()
func (_EthCrossChainCaller *EthCrossChainCallerTransactor) Initialize(opts *bind.TransactOpts, _logic common.Address, _admin common.Address, _data []byte) (*types.Transaction, error) {
	return _EthCrossChainCaller.contract.Transact(opts, "initialize", _logic, _admin, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf7a1d77.
//
// Solidity: function initialize(address _logic, address _admin, bytes _data) payable returns()
func (_EthCrossChainCaller *EthCrossChainCallerSession) Initialize(_logic common.Address, _admin common.Address, _data []byte) (*types.Transaction, error) {
	return _EthCrossChainCaller.Contract.Initialize(&_EthCrossChainCaller.TransactOpts, _logic, _admin, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf7a1d77.
//
// Solidity: function initialize(address _logic, address _admin, bytes _data) payable returns()
func (_EthCrossChainCaller *EthCrossChainCallerTransactorSession) Initialize(_logic common.Address, _admin common.Address, _data []byte) (*types.Transaction, error) {
	return _EthCrossChainCaller.Contract.Initialize(&_EthCrossChainCaller.TransactOpts, _logic, _admin, _data)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xd1f57894.
//
// Solidity: function initialize(address _logic, bytes _data) payable returns()
func (_EthCrossChainCaller *EthCrossChainCallerTransactor) Initialize0(opts *bind.TransactOpts, _logic common.Address, _data []byte) (*types.Transaction, error) {
	return _EthCrossChainCaller.contract.Transact(opts, "initialize0", _logic, _data)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xd1f57894.
//
// Solidity: function initialize(address _logic, bytes _data) payable returns()
func (_EthCrossChainCaller *EthCrossChainCallerSession) Initialize0(_logic common.Address, _data []byte) (*types.Transaction, error) {
	return _EthCrossChainCaller.Contract.Initialize0(&_EthCrossChainCaller.TransactOpts, _logic, _data)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xd1f57894.
//
// Solidity: function initialize(address _logic, bytes _data) payable returns()
func (_EthCrossChainCaller *EthCrossChainCallerTransactorSession) Initialize0(_logic common.Address, _data []byte) (*types.Transaction, error) {
	return _EthCrossChainCaller.Contract.Initialize0(&_EthCrossChainCaller.TransactOpts, _logic, _data)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_EthCrossChainCaller *EthCrossChainCallerTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _EthCrossChainCaller.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_EthCrossChainCaller *EthCrossChainCallerSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _EthCrossChainCaller.Contract.UpgradeTo(&_EthCrossChainCaller.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_EthCrossChainCaller *EthCrossChainCallerTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _EthCrossChainCaller.Contract.UpgradeTo(&_EthCrossChainCaller.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EthCrossChainCaller *EthCrossChainCallerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EthCrossChainCaller.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EthCrossChainCaller *EthCrossChainCallerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EthCrossChainCaller.Contract.UpgradeToAndCall(&_EthCrossChainCaller.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EthCrossChainCaller *EthCrossChainCallerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EthCrossChainCaller.Contract.UpgradeToAndCall(&_EthCrossChainCaller.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_EthCrossChainCaller *EthCrossChainCallerTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _EthCrossChainCaller.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_EthCrossChainCaller *EthCrossChainCallerSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _EthCrossChainCaller.Contract.Fallback(&_EthCrossChainCaller.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_EthCrossChainCaller *EthCrossChainCallerTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _EthCrossChainCaller.Contract.Fallback(&_EthCrossChainCaller.TransactOpts, calldata)
}

// EthCrossChainCallerAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the EthCrossChainCaller contract.
type EthCrossChainCallerAdminChangedIterator struct {
	Event *EthCrossChainCallerAdminChanged // Event containing the contract specifics and raw log

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
func (it *EthCrossChainCallerAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainCallerAdminChanged)
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
		it.Event = new(EthCrossChainCallerAdminChanged)
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
func (it *EthCrossChainCallerAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainCallerAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainCallerAdminChanged represents a AdminChanged event raised by the EthCrossChainCaller contract.
type EthCrossChainCallerAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_EthCrossChainCaller *EthCrossChainCallerFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*EthCrossChainCallerAdminChangedIterator, error) {

	logs, sub, err := _EthCrossChainCaller.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &EthCrossChainCallerAdminChangedIterator{contract: _EthCrossChainCaller.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_EthCrossChainCaller *EthCrossChainCallerFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *EthCrossChainCallerAdminChanged) (event.Subscription, error) {

	logs, sub, err := _EthCrossChainCaller.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainCallerAdminChanged)
				if err := _EthCrossChainCaller.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_EthCrossChainCaller *EthCrossChainCallerFilterer) ParseAdminChanged(log types.Log) (*EthCrossChainCallerAdminChanged, error) {
	event := new(EthCrossChainCallerAdminChanged)
	if err := _EthCrossChainCaller.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthCrossChainCallerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the EthCrossChainCaller contract.
type EthCrossChainCallerUpgradedIterator struct {
	Event *EthCrossChainCallerUpgraded // Event containing the contract specifics and raw log

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
func (it *EthCrossChainCallerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainCallerUpgraded)
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
		it.Event = new(EthCrossChainCallerUpgraded)
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
func (it *EthCrossChainCallerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainCallerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainCallerUpgraded represents a Upgraded event raised by the EthCrossChainCaller contract.
type EthCrossChainCallerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EthCrossChainCaller *EthCrossChainCallerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*EthCrossChainCallerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EthCrossChainCaller.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainCallerUpgradedIterator{contract: _EthCrossChainCaller.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EthCrossChainCaller *EthCrossChainCallerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *EthCrossChainCallerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EthCrossChainCaller.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainCallerUpgraded)
				if err := _EthCrossChainCaller.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EthCrossChainCaller *EthCrossChainCallerFilterer) ParseUpgraded(log types.Log) (*EthCrossChainCallerUpgraded, error) {
	event := new(EthCrossChainCallerUpgraded)
	if err := _EthCrossChainCaller.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthCrossChainDataABI is the input ABI used to generate the binding from.
const EthCrossChainDataABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"CurEpochEndHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CurEpochId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CurEpochStartHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CurValidatorPkBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EthToPolyTxHashIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"EthToPolyTxHashMap\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ExtraData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"fromChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"fromChainTx\",\"type\":\"bytes32\"}],\"name\":\"checkIfFromChainTxExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurEpochEndHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurEpochId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurEpochStartHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurEpochValidatorPkBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ethTxHashIndex\",\"type\":\"uint256\"}],\"name\":\"getEthTxHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEthTxHashIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"key2\",\"type\":\"bytes32\"}],\"name\":\"getExtraData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"fromChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"fromChainTx\",\"type\":\"bytes32\"}],\"name\":\"markFromChainTxExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"endHeight\",\"type\":\"uint64\"}],\"name\":\"putCurEpochEndHeight\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"epochId\",\"type\":\"uint64\"}],\"name\":\"putCurEpochId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"startHeight\",\"type\":\"uint64\"}],\"name\":\"putCurEpochStartHeight\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"curEpochPkBytes\",\"type\":\"bytes\"}],\"name\":\"putCurEpochValidatorPkBytes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ethTxHash\",\"type\":\"bytes32\"}],\"name\":\"putEthTxHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"key2\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"putExtraData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EthCrossChainDataFuncSigs maps the 4-byte function signature to its string representation.
var EthCrossChainDataFuncSigs = map[string]string{
	"ab9844ab": "CurEpochEndHeight()",
	"fd4b67cf": "CurEpochId()",
	"67e31a74": "CurEpochStartHeight()",
	"06bfb1b6": "CurValidatorPkBytes()",
	"00c5fff8": "EthToPolyTxHashIndex()",
	"529caad8": "EthToPolyTxHashMap(uint256)",
	"20bbde38": "ExtraData(bytes32,bytes32)",
	"0586763c": "checkIfFromChainTxExist(uint64,bytes32)",
	"4e4af8ce": "getCurEpochEndHeight()",
	"f881afd2": "getCurEpochId()",
	"5ac40790": "getCurEpochStartHeight()",
	"cd62908f": "getCurEpochValidatorPkBytes()",
	"29927875": "getEthTxHash(uint256)",
	"ff3d24a1": "getEthTxHashIndex()",
	"40602bb5": "getExtraData(bytes32,bytes32)",
	"8f32d59b": "isOwner()",
	"e90bfdcf": "markFromChainTxExist(uint64,bytes32)",
	"8da5cb5b": "owner()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"6a773962": "putCurEpochEndHeight(uint64)",
	"738b5ae4": "putCurEpochId(uint64)",
	"512feecc": "putCurEpochStartHeight(uint64)",
	"1dc544bf": "putCurEpochValidatorPkBytes(bytes)",
	"4c3ccf64": "putEthTxHash(bytes32)",
	"1afe374e": "putExtraData(bytes32,bytes32,bytes)",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
	"3f4ba83a": "unpause()",
}

// EthCrossChainDataBin is the compiled bytecode used for deploying new contracts.
var EthCrossChainDataBin = "0x608060405260006100176001600160e01b0361007316565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506000805460ff60a01b19169055610077565b3390565b6113e3806100866000396000f3fe608060405234801561001057600080fd5b50600436106101ce5760003560e01c80635c975abb116101045780638f32d59b116100a2578063f2fde38b11610071578063f2fde38b146105cf578063f881afd2146105f5578063fd4b67cf146105fd578063ff3d24a114610605576101ce565b80638f32d59b1461058b578063ab9844ab14610593578063cd62908f1461059b578063e90bfdcf146105a3576101ce565b8063715018a6116100de578063715018a61461052f578063738b5ae4146105395780638456cb591461055f5780638da5cb5b14610567576101ce565b80635c975abb146104f957806367e31a74146105015780636a77396214610509576101ce565b80633f4ba83a116101715780634e4af8ce1161014b5780634e4af8ce1461048a578063512feecc146104ae578063529caad8146104d45780635ac40790146104f1576101ce565b80633f4ba83a1461044257806340602bb51461044a5780634c3ccf641461046d576101ce565b80631afe374e116101ad5780631afe374e146102aa5780631dc544bf1461035c57806320bbde38146104025780632992787514610425576101ce565b8062c5fff8146101d35780630586763c146101ed57806306bfb1b61461022d575b600080fd5b6101db61060d565b60408051918252519081900360200190f35b6102196004803603604081101561020357600080fd5b506001600160401b038135169060200135610613565b604080519115158252519081900360200190f35b61023561063e565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561026f578181015183820152602001610257565b50505050905090810190601f16801561029c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610219600480360360608110156102c057600080fd5b8135916020810135918101906060810160408201356401000000008111156102e757600080fd5b8201836020820111156102f957600080fd5b8035906020019184600183028401116401000000008311171561031b57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506106cc945050505050565b6102196004803603602081101561037257600080fd5b81019060208101813564010000000081111561038d57600080fd5b82018360208201111561039f57600080fd5b803590602001918460018302840111640100000000831117156103c157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610798945050505050565b6102356004803603604081101561041857600080fd5b508035906020013561084e565b6101db6004803603602081101561043b57600080fd5b50356108bf565b6102196108d1565b6102356004803603604081101561046057600080fd5b508035906020013561097e565b6102196004803603602081101561048357600080fd5b5035610a28565b610492610ae5565b604080516001600160401b039092168252519081900360200190f35b610219600480360360208110156104c457600080fd5b50356001600160401b0316610afb565b6101db600480360360208110156104ea57600080fd5b5035610bbb565b610492610bcd565b610219610bdc565b610492610bec565b6102196004803603602081101561051f57600080fd5b50356001600160401b0316610bfb565b610537610cc9565b005b6102196004803603602081101561054f57600080fd5b50356001600160401b0316610d5a565b610219610e23565b61056f610ec6565b604080516001600160a01b039092168252519081900360200190f35b610219610ed5565b610492610ef9565b610235610f0f565b610219600480360360408110156105b957600080fd5b506001600160401b038135169060200135610fa5565b610537600480360360208110156105e557600080fd5b50356001600160a01b0316611075565b6104926110c8565b6104926110de565b6101db6110f4565b60025481565b6001600160401b03919091166000908152600560209081526040808320938352929052205460ff1690565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106c45780601f10610699576101008083540402835291602001916106c4565b820191906000526020600020905b8154815290600101906020018083116106a757829003601f168201915b505050505081565b60008054600160a01b900460ff161561071f576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b610727610ed5565b610766576040805162461bcd60e51b8152602060048201819052602482015260008051602061138f833981519152604482015290519081900360640190fd5b60008481526006602090815260408083208684528252909120835161078d928501906112d0565b506001949350505050565b60008054600160a01b900460ff16156107eb576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b6107f3610ed5565b610832576040805162461bcd60e51b8152602060048201819052602482015260008051602061138f833981519152604482015290519081900360640190fd5b81516108459060039060208501906112d0565b50600192915050565b60066020908152600092835260408084208252918352918190208054825160026001831615610100026000190190921691909104601f8101859004850282018501909352828152929091908301828280156106c45780601f10610699576101008083540402835291602001916106c4565b60009081526001602052604090205490565b60006108db610ed5565b61091a576040805162461bcd60e51b8152602060048201819052602482015260008051602061138f833981519152604482015290519081900360640190fd5b600054600160a01b900460ff1661096f576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b6109776110fa565b5060015b90565b600082815260066020908152604080832084845282529182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845260609392830182828015610a1b5780601f106109f057610100808354040283529160200191610a1b565b820191906000526020600020905b8154815290600101906020018083116109fe57829003601f168201915b5050505050905092915050565b60008054600160a01b900460ff1615610a7b576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b610a83610ed5565b610ac2576040805162461bcd60e51b8152602060048201819052602482015260008051602061138f833981519152604482015290519081900360640190fd5b506002805460009081526001602081905260409091209290925580548201905590565b600454600160401b90046001600160401b031690565b60008054600160a01b900460ff1615610b4e576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b610b56610ed5565b610b95576040805162461bcd60e51b8152602060048201819052602482015260008051602061138f833981519152604482015290519081900360640190fd5b50600480546001600160401b03831667ffffffffffffffff199091161790556001919050565b60016020526000908152604090205481565b6004546001600160401b031690565b600054600160a01b900460ff1690565b6004546001600160401b031681565b60008054600160a01b900460ff1615610c4e576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b610c56610ed5565b610c95576040805162461bcd60e51b8152602060048201819052602482015260008051602061138f833981519152604482015290519081900360640190fd5b50600480546001600160401b038316600160401b026fffffffffffffffff0000000000000000199091161790556001919050565b610cd1610ed5565b610d10576040805162461bcd60e51b8152602060048201819052602482015260008051602061138f833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b60008054600160a01b900460ff1615610dad576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b610db5610ed5565b610df4576040805162461bcd60e51b8152602060048201819052602482015260008051602061138f833981519152604482015290519081900360640190fd5b50600480546001600160401b038316600160801b0267ffffffffffffffff60801b199091161790556001919050565b6000610e2d610ed5565b610e6c576040805162461bcd60e51b8152602060048201819052602482015260008051602061138f833981519152604482015290519081900360640190fd5b600054600160a01b900460ff1615610ebe576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b6109776111a2565b6000546001600160a01b031690565b600080546001600160a01b0316610eea61122c565b6001600160a01b031614905090565b600454600160401b90046001600160401b031681565b60038054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815260609390929091830182828015610f9b5780601f10610f7057610100808354040283529160200191610f9b565b820191906000526020600020905b815481529060010190602001808311610f7e57829003601f168201915b5050505050905090565b60008054600160a01b900460ff1615610ff8576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b611000610ed5565b61103f576040805162461bcd60e51b8152602060048201819052602482015260008051602061138f833981519152604482015290519081900360640190fd5b506001600160401b0391909116600090815260056020908152604080832093835292905220805460ff1916600190811790915590565b61107d610ed5565b6110bc576040805162461bcd60e51b8152602060048201819052602482015260008051602061138f833981519152604482015290519081900360640190fd5b6110c581611230565b50565b600454600160801b90046001600160401b031690565b600454600160801b90046001600160401b031681565b60025490565b600054600160a01b900460ff1661114f576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa61118561122c565b604080516001600160a01b039092168252519081900360200190a1565b600054600160a01b900460ff16156111f4576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586111855b3390565b6001600160a01b0381166112755760405162461bcd60e51b81526004018080602001828103825260268152602001806113696026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061131157805160ff191683800117855561133e565b8280016001018555821561133e579182015b8281111561133e578251825591602001919060010190611323565b5061134a92915061134e565b5090565b61097b91905b8082111561134a576000815560010161135456fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572a265627a7a72315820866cfa18d69c07957cf932116fcff4e5cf2b55ec5ed87ae041839a5333a17a4064736f6c63430005110032"

// DeployEthCrossChainData deploys a new Ethereum contract, binding an instance of EthCrossChainData to it.
func DeployEthCrossChainData(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EthCrossChainData, error) {
	parsed, err := abi.JSON(strings.NewReader(EthCrossChainDataABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EthCrossChainDataBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthCrossChainData{EthCrossChainDataCaller: EthCrossChainDataCaller{contract: contract}, EthCrossChainDataTransactor: EthCrossChainDataTransactor{contract: contract}, EthCrossChainDataFilterer: EthCrossChainDataFilterer{contract: contract}}, nil
}

// EthCrossChainData is an auto generated Go binding around an Ethereum contract.
type EthCrossChainData struct {
	EthCrossChainDataCaller     // Read-only binding to the contract
	EthCrossChainDataTransactor // Write-only binding to the contract
	EthCrossChainDataFilterer   // Log filterer for contract events
}

// EthCrossChainDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthCrossChainDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthCrossChainDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthCrossChainDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthCrossChainDataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthCrossChainDataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthCrossChainDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthCrossChainDataSession struct {
	Contract     *EthCrossChainData // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EthCrossChainDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthCrossChainDataCallerSession struct {
	Contract *EthCrossChainDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// EthCrossChainDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthCrossChainDataTransactorSession struct {
	Contract     *EthCrossChainDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// EthCrossChainDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthCrossChainDataRaw struct {
	Contract *EthCrossChainData // Generic contract binding to access the raw methods on
}

// EthCrossChainDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthCrossChainDataCallerRaw struct {
	Contract *EthCrossChainDataCaller // Generic read-only contract binding to access the raw methods on
}

// EthCrossChainDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthCrossChainDataTransactorRaw struct {
	Contract *EthCrossChainDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthCrossChainData creates a new instance of EthCrossChainData, bound to a specific deployed contract.
func NewEthCrossChainData(address common.Address, backend bind.ContractBackend) (*EthCrossChainData, error) {
	contract, err := bindEthCrossChainData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainData{EthCrossChainDataCaller: EthCrossChainDataCaller{contract: contract}, EthCrossChainDataTransactor: EthCrossChainDataTransactor{contract: contract}, EthCrossChainDataFilterer: EthCrossChainDataFilterer{contract: contract}}, nil
}

// NewEthCrossChainDataCaller creates a new read-only instance of EthCrossChainData, bound to a specific deployed contract.
func NewEthCrossChainDataCaller(address common.Address, caller bind.ContractCaller) (*EthCrossChainDataCaller, error) {
	contract, err := bindEthCrossChainData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainDataCaller{contract: contract}, nil
}

// NewEthCrossChainDataTransactor creates a new write-only instance of EthCrossChainData, bound to a specific deployed contract.
func NewEthCrossChainDataTransactor(address common.Address, transactor bind.ContractTransactor) (*EthCrossChainDataTransactor, error) {
	contract, err := bindEthCrossChainData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainDataTransactor{contract: contract}, nil
}

// NewEthCrossChainDataFilterer creates a new log filterer instance of EthCrossChainData, bound to a specific deployed contract.
func NewEthCrossChainDataFilterer(address common.Address, filterer bind.ContractFilterer) (*EthCrossChainDataFilterer, error) {
	contract, err := bindEthCrossChainData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainDataFilterer{contract: contract}, nil
}

// bindEthCrossChainData binds a generic wrapper to an already deployed contract.
func bindEthCrossChainData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthCrossChainDataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthCrossChainData *EthCrossChainDataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthCrossChainData.Contract.EthCrossChainDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthCrossChainData *EthCrossChainDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.EthCrossChainDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthCrossChainData *EthCrossChainDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.EthCrossChainDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthCrossChainData *EthCrossChainDataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthCrossChainData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthCrossChainData *EthCrossChainDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthCrossChainData *EthCrossChainDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.contract.Transact(opts, method, params...)
}

// CurEpochEndHeight is a free data retrieval call binding the contract method 0xab9844ab.
//
// Solidity: function CurEpochEndHeight() view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataCaller) CurEpochEndHeight(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "CurEpochEndHeight")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CurEpochEndHeight is a free data retrieval call binding the contract method 0xab9844ab.
//
// Solidity: function CurEpochEndHeight() view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataSession) CurEpochEndHeight() (uint64, error) {
	return _EthCrossChainData.Contract.CurEpochEndHeight(&_EthCrossChainData.CallOpts)
}

// CurEpochEndHeight is a free data retrieval call binding the contract method 0xab9844ab.
//
// Solidity: function CurEpochEndHeight() view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataCallerSession) CurEpochEndHeight() (uint64, error) {
	return _EthCrossChainData.Contract.CurEpochEndHeight(&_EthCrossChainData.CallOpts)
}

// CurEpochId is a free data retrieval call binding the contract method 0xfd4b67cf.
//
// Solidity: function CurEpochId() view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataCaller) CurEpochId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "CurEpochId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CurEpochId is a free data retrieval call binding the contract method 0xfd4b67cf.
//
// Solidity: function CurEpochId() view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataSession) CurEpochId() (uint64, error) {
	return _EthCrossChainData.Contract.CurEpochId(&_EthCrossChainData.CallOpts)
}

// CurEpochId is a free data retrieval call binding the contract method 0xfd4b67cf.
//
// Solidity: function CurEpochId() view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataCallerSession) CurEpochId() (uint64, error) {
	return _EthCrossChainData.Contract.CurEpochId(&_EthCrossChainData.CallOpts)
}

// CurEpochStartHeight is a free data retrieval call binding the contract method 0x67e31a74.
//
// Solidity: function CurEpochStartHeight() view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataCaller) CurEpochStartHeight(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "CurEpochStartHeight")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CurEpochStartHeight is a free data retrieval call binding the contract method 0x67e31a74.
//
// Solidity: function CurEpochStartHeight() view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataSession) CurEpochStartHeight() (uint64, error) {
	return _EthCrossChainData.Contract.CurEpochStartHeight(&_EthCrossChainData.CallOpts)
}

// CurEpochStartHeight is a free data retrieval call binding the contract method 0x67e31a74.
//
// Solidity: function CurEpochStartHeight() view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataCallerSession) CurEpochStartHeight() (uint64, error) {
	return _EthCrossChainData.Contract.CurEpochStartHeight(&_EthCrossChainData.CallOpts)
}

// CurValidatorPkBytes is a free data retrieval call binding the contract method 0x06bfb1b6.
//
// Solidity: function CurValidatorPkBytes() view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCaller) CurValidatorPkBytes(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "CurValidatorPkBytes")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CurValidatorPkBytes is a free data retrieval call binding the contract method 0x06bfb1b6.
//
// Solidity: function CurValidatorPkBytes() view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataSession) CurValidatorPkBytes() ([]byte, error) {
	return _EthCrossChainData.Contract.CurValidatorPkBytes(&_EthCrossChainData.CallOpts)
}

// CurValidatorPkBytes is a free data retrieval call binding the contract method 0x06bfb1b6.
//
// Solidity: function CurValidatorPkBytes() view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCallerSession) CurValidatorPkBytes() ([]byte, error) {
	return _EthCrossChainData.Contract.CurValidatorPkBytes(&_EthCrossChainData.CallOpts)
}

// EthToPolyTxHashIndex is a free data retrieval call binding the contract method 0x00c5fff8.
//
// Solidity: function EthToPolyTxHashIndex() view returns(uint256)
func (_EthCrossChainData *EthCrossChainDataCaller) EthToPolyTxHashIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "EthToPolyTxHashIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EthToPolyTxHashIndex is a free data retrieval call binding the contract method 0x00c5fff8.
//
// Solidity: function EthToPolyTxHashIndex() view returns(uint256)
func (_EthCrossChainData *EthCrossChainDataSession) EthToPolyTxHashIndex() (*big.Int, error) {
	return _EthCrossChainData.Contract.EthToPolyTxHashIndex(&_EthCrossChainData.CallOpts)
}

// EthToPolyTxHashIndex is a free data retrieval call binding the contract method 0x00c5fff8.
//
// Solidity: function EthToPolyTxHashIndex() view returns(uint256)
func (_EthCrossChainData *EthCrossChainDataCallerSession) EthToPolyTxHashIndex() (*big.Int, error) {
	return _EthCrossChainData.Contract.EthToPolyTxHashIndex(&_EthCrossChainData.CallOpts)
}

// EthToPolyTxHashMap is a free data retrieval call binding the contract method 0x529caad8.
//
// Solidity: function EthToPolyTxHashMap(uint256 ) view returns(bytes32)
func (_EthCrossChainData *EthCrossChainDataCaller) EthToPolyTxHashMap(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "EthToPolyTxHashMap", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EthToPolyTxHashMap is a free data retrieval call binding the contract method 0x529caad8.
//
// Solidity: function EthToPolyTxHashMap(uint256 ) view returns(bytes32)
func (_EthCrossChainData *EthCrossChainDataSession) EthToPolyTxHashMap(arg0 *big.Int) ([32]byte, error) {
	return _EthCrossChainData.Contract.EthToPolyTxHashMap(&_EthCrossChainData.CallOpts, arg0)
}

// EthToPolyTxHashMap is a free data retrieval call binding the contract method 0x529caad8.
//
// Solidity: function EthToPolyTxHashMap(uint256 ) view returns(bytes32)
func (_EthCrossChainData *EthCrossChainDataCallerSession) EthToPolyTxHashMap(arg0 *big.Int) ([32]byte, error) {
	return _EthCrossChainData.Contract.EthToPolyTxHashMap(&_EthCrossChainData.CallOpts, arg0)
}

// ExtraData is a free data retrieval call binding the contract method 0x20bbde38.
//
// Solidity: function ExtraData(bytes32 , bytes32 ) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCaller) ExtraData(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte) ([]byte, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "ExtraData", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ExtraData is a free data retrieval call binding the contract method 0x20bbde38.
//
// Solidity: function ExtraData(bytes32 , bytes32 ) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataSession) ExtraData(arg0 [32]byte, arg1 [32]byte) ([]byte, error) {
	return _EthCrossChainData.Contract.ExtraData(&_EthCrossChainData.CallOpts, arg0, arg1)
}

// ExtraData is a free data retrieval call binding the contract method 0x20bbde38.
//
// Solidity: function ExtraData(bytes32 , bytes32 ) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCallerSession) ExtraData(arg0 [32]byte, arg1 [32]byte) ([]byte, error) {
	return _EthCrossChainData.Contract.ExtraData(&_EthCrossChainData.CallOpts, arg0, arg1)
}

// CheckIfFromChainTxExist is a free data retrieval call binding the contract method 0x0586763c.
//
// Solidity: function checkIfFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) view returns(bool)
func (_EthCrossChainData *EthCrossChainDataCaller) CheckIfFromChainTxExist(opts *bind.CallOpts, fromChainId uint64, fromChainTx [32]byte) (bool, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "checkIfFromChainTxExist", fromChainId, fromChainTx)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckIfFromChainTxExist is a free data retrieval call binding the contract method 0x0586763c.
//
// Solidity: function checkIfFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) view returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) CheckIfFromChainTxExist(fromChainId uint64, fromChainTx [32]byte) (bool, error) {
	return _EthCrossChainData.Contract.CheckIfFromChainTxExist(&_EthCrossChainData.CallOpts, fromChainId, fromChainTx)
}

// CheckIfFromChainTxExist is a free data retrieval call binding the contract method 0x0586763c.
//
// Solidity: function checkIfFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) view returns(bool)
func (_EthCrossChainData *EthCrossChainDataCallerSession) CheckIfFromChainTxExist(fromChainId uint64, fromChainTx [32]byte) (bool, error) {
	return _EthCrossChainData.Contract.CheckIfFromChainTxExist(&_EthCrossChainData.CallOpts, fromChainId, fromChainTx)
}

// GetCurEpochEndHeight is a free data retrieval call binding the contract method 0x4e4af8ce.
//
// Solidity: function getCurEpochEndHeight() view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataCaller) GetCurEpochEndHeight(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "getCurEpochEndHeight")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetCurEpochEndHeight is a free data retrieval call binding the contract method 0x4e4af8ce.
//
// Solidity: function getCurEpochEndHeight() view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataSession) GetCurEpochEndHeight() (uint64, error) {
	return _EthCrossChainData.Contract.GetCurEpochEndHeight(&_EthCrossChainData.CallOpts)
}

// GetCurEpochEndHeight is a free data retrieval call binding the contract method 0x4e4af8ce.
//
// Solidity: function getCurEpochEndHeight() view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataCallerSession) GetCurEpochEndHeight() (uint64, error) {
	return _EthCrossChainData.Contract.GetCurEpochEndHeight(&_EthCrossChainData.CallOpts)
}

// GetCurEpochId is a free data retrieval call binding the contract method 0xf881afd2.
//
// Solidity: function getCurEpochId() view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataCaller) GetCurEpochId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "getCurEpochId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetCurEpochId is a free data retrieval call binding the contract method 0xf881afd2.
//
// Solidity: function getCurEpochId() view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataSession) GetCurEpochId() (uint64, error) {
	return _EthCrossChainData.Contract.GetCurEpochId(&_EthCrossChainData.CallOpts)
}

// GetCurEpochId is a free data retrieval call binding the contract method 0xf881afd2.
//
// Solidity: function getCurEpochId() view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataCallerSession) GetCurEpochId() (uint64, error) {
	return _EthCrossChainData.Contract.GetCurEpochId(&_EthCrossChainData.CallOpts)
}

// GetCurEpochStartHeight is a free data retrieval call binding the contract method 0x5ac40790.
//
// Solidity: function getCurEpochStartHeight() view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataCaller) GetCurEpochStartHeight(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "getCurEpochStartHeight")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetCurEpochStartHeight is a free data retrieval call binding the contract method 0x5ac40790.
//
// Solidity: function getCurEpochStartHeight() view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataSession) GetCurEpochStartHeight() (uint64, error) {
	return _EthCrossChainData.Contract.GetCurEpochStartHeight(&_EthCrossChainData.CallOpts)
}

// GetCurEpochStartHeight is a free data retrieval call binding the contract method 0x5ac40790.
//
// Solidity: function getCurEpochStartHeight() view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataCallerSession) GetCurEpochStartHeight() (uint64, error) {
	return _EthCrossChainData.Contract.GetCurEpochStartHeight(&_EthCrossChainData.CallOpts)
}

// GetCurEpochValidatorPkBytes is a free data retrieval call binding the contract method 0xcd62908f.
//
// Solidity: function getCurEpochValidatorPkBytes() view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCaller) GetCurEpochValidatorPkBytes(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "getCurEpochValidatorPkBytes")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetCurEpochValidatorPkBytes is a free data retrieval call binding the contract method 0xcd62908f.
//
// Solidity: function getCurEpochValidatorPkBytes() view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataSession) GetCurEpochValidatorPkBytes() ([]byte, error) {
	return _EthCrossChainData.Contract.GetCurEpochValidatorPkBytes(&_EthCrossChainData.CallOpts)
}

// GetCurEpochValidatorPkBytes is a free data retrieval call binding the contract method 0xcd62908f.
//
// Solidity: function getCurEpochValidatorPkBytes() view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCallerSession) GetCurEpochValidatorPkBytes() ([]byte, error) {
	return _EthCrossChainData.Contract.GetCurEpochValidatorPkBytes(&_EthCrossChainData.CallOpts)
}

// GetEthTxHash is a free data retrieval call binding the contract method 0x29927875.
//
// Solidity: function getEthTxHash(uint256 ethTxHashIndex) view returns(bytes32)
func (_EthCrossChainData *EthCrossChainDataCaller) GetEthTxHash(opts *bind.CallOpts, ethTxHashIndex *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "getEthTxHash", ethTxHashIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEthTxHash is a free data retrieval call binding the contract method 0x29927875.
//
// Solidity: function getEthTxHash(uint256 ethTxHashIndex) view returns(bytes32)
func (_EthCrossChainData *EthCrossChainDataSession) GetEthTxHash(ethTxHashIndex *big.Int) ([32]byte, error) {
	return _EthCrossChainData.Contract.GetEthTxHash(&_EthCrossChainData.CallOpts, ethTxHashIndex)
}

// GetEthTxHash is a free data retrieval call binding the contract method 0x29927875.
//
// Solidity: function getEthTxHash(uint256 ethTxHashIndex) view returns(bytes32)
func (_EthCrossChainData *EthCrossChainDataCallerSession) GetEthTxHash(ethTxHashIndex *big.Int) ([32]byte, error) {
	return _EthCrossChainData.Contract.GetEthTxHash(&_EthCrossChainData.CallOpts, ethTxHashIndex)
}

// GetEthTxHashIndex is a free data retrieval call binding the contract method 0xff3d24a1.
//
// Solidity: function getEthTxHashIndex() view returns(uint256)
func (_EthCrossChainData *EthCrossChainDataCaller) GetEthTxHashIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "getEthTxHashIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthTxHashIndex is a free data retrieval call binding the contract method 0xff3d24a1.
//
// Solidity: function getEthTxHashIndex() view returns(uint256)
func (_EthCrossChainData *EthCrossChainDataSession) GetEthTxHashIndex() (*big.Int, error) {
	return _EthCrossChainData.Contract.GetEthTxHashIndex(&_EthCrossChainData.CallOpts)
}

// GetEthTxHashIndex is a free data retrieval call binding the contract method 0xff3d24a1.
//
// Solidity: function getEthTxHashIndex() view returns(uint256)
func (_EthCrossChainData *EthCrossChainDataCallerSession) GetEthTxHashIndex() (*big.Int, error) {
	return _EthCrossChainData.Contract.GetEthTxHashIndex(&_EthCrossChainData.CallOpts)
}

// GetExtraData is a free data retrieval call binding the contract method 0x40602bb5.
//
// Solidity: function getExtraData(bytes32 key1, bytes32 key2) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCaller) GetExtraData(opts *bind.CallOpts, key1 [32]byte, key2 [32]byte) ([]byte, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "getExtraData", key1, key2)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetExtraData is a free data retrieval call binding the contract method 0x40602bb5.
//
// Solidity: function getExtraData(bytes32 key1, bytes32 key2) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataSession) GetExtraData(key1 [32]byte, key2 [32]byte) ([]byte, error) {
	return _EthCrossChainData.Contract.GetExtraData(&_EthCrossChainData.CallOpts, key1, key2)
}

// GetExtraData is a free data retrieval call binding the contract method 0x40602bb5.
//
// Solidity: function getExtraData(bytes32 key1, bytes32 key2) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCallerSession) GetExtraData(key1 [32]byte, key2 [32]byte) ([]byte, error) {
	return _EthCrossChainData.Contract.GetExtraData(&_EthCrossChainData.CallOpts, key1, key2)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_EthCrossChainData *EthCrossChainDataCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) IsOwner() (bool, error) {
	return _EthCrossChainData.Contract.IsOwner(&_EthCrossChainData.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_EthCrossChainData *EthCrossChainDataCallerSession) IsOwner() (bool, error) {
	return _EthCrossChainData.Contract.IsOwner(&_EthCrossChainData.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthCrossChainData *EthCrossChainDataCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthCrossChainData *EthCrossChainDataSession) Owner() (common.Address, error) {
	return _EthCrossChainData.Contract.Owner(&_EthCrossChainData.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthCrossChainData *EthCrossChainDataCallerSession) Owner() (common.Address, error) {
	return _EthCrossChainData.Contract.Owner(&_EthCrossChainData.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EthCrossChainData *EthCrossChainDataCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) Paused() (bool, error) {
	return _EthCrossChainData.Contract.Paused(&_EthCrossChainData.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EthCrossChainData *EthCrossChainDataCallerSession) Paused() (bool, error) {
	return _EthCrossChainData.Contract.Paused(&_EthCrossChainData.CallOpts)
}

// MarkFromChainTxExist is a paid mutator transaction binding the contract method 0xe90bfdcf.
//
// Solidity: function markFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) MarkFromChainTxExist(opts *bind.TransactOpts, fromChainId uint64, fromChainTx [32]byte) (*types.Transaction, error) {
	return _EthCrossChainData.contract.Transact(opts, "markFromChainTxExist", fromChainId, fromChainTx)
}

// MarkFromChainTxExist is a paid mutator transaction binding the contract method 0xe90bfdcf.
//
// Solidity: function markFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) MarkFromChainTxExist(fromChainId uint64, fromChainTx [32]byte) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.MarkFromChainTxExist(&_EthCrossChainData.TransactOpts, fromChainId, fromChainTx)
}

// MarkFromChainTxExist is a paid mutator transaction binding the contract method 0xe90bfdcf.
//
// Solidity: function markFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) MarkFromChainTxExist(fromChainId uint64, fromChainTx [32]byte) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.MarkFromChainTxExist(&_EthCrossChainData.TransactOpts, fromChainId, fromChainTx)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainData.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) Pause() (*types.Transaction, error) {
	return _EthCrossChainData.Contract.Pause(&_EthCrossChainData.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) Pause() (*types.Transaction, error) {
	return _EthCrossChainData.Contract.Pause(&_EthCrossChainData.TransactOpts)
}

// PutCurEpochEndHeight is a paid mutator transaction binding the contract method 0x6a773962.
//
// Solidity: function putCurEpochEndHeight(uint64 endHeight) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) PutCurEpochEndHeight(opts *bind.TransactOpts, endHeight uint64) (*types.Transaction, error) {
	return _EthCrossChainData.contract.Transact(opts, "putCurEpochEndHeight", endHeight)
}

// PutCurEpochEndHeight is a paid mutator transaction binding the contract method 0x6a773962.
//
// Solidity: function putCurEpochEndHeight(uint64 endHeight) returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) PutCurEpochEndHeight(endHeight uint64) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutCurEpochEndHeight(&_EthCrossChainData.TransactOpts, endHeight)
}

// PutCurEpochEndHeight is a paid mutator transaction binding the contract method 0x6a773962.
//
// Solidity: function putCurEpochEndHeight(uint64 endHeight) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) PutCurEpochEndHeight(endHeight uint64) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutCurEpochEndHeight(&_EthCrossChainData.TransactOpts, endHeight)
}

// PutCurEpochId is a paid mutator transaction binding the contract method 0x738b5ae4.
//
// Solidity: function putCurEpochId(uint64 epochId) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) PutCurEpochId(opts *bind.TransactOpts, epochId uint64) (*types.Transaction, error) {
	return _EthCrossChainData.contract.Transact(opts, "putCurEpochId", epochId)
}

// PutCurEpochId is a paid mutator transaction binding the contract method 0x738b5ae4.
//
// Solidity: function putCurEpochId(uint64 epochId) returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) PutCurEpochId(epochId uint64) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutCurEpochId(&_EthCrossChainData.TransactOpts, epochId)
}

// PutCurEpochId is a paid mutator transaction binding the contract method 0x738b5ae4.
//
// Solidity: function putCurEpochId(uint64 epochId) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) PutCurEpochId(epochId uint64) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutCurEpochId(&_EthCrossChainData.TransactOpts, epochId)
}

// PutCurEpochStartHeight is a paid mutator transaction binding the contract method 0x512feecc.
//
// Solidity: function putCurEpochStartHeight(uint64 startHeight) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) PutCurEpochStartHeight(opts *bind.TransactOpts, startHeight uint64) (*types.Transaction, error) {
	return _EthCrossChainData.contract.Transact(opts, "putCurEpochStartHeight", startHeight)
}

// PutCurEpochStartHeight is a paid mutator transaction binding the contract method 0x512feecc.
//
// Solidity: function putCurEpochStartHeight(uint64 startHeight) returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) PutCurEpochStartHeight(startHeight uint64) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutCurEpochStartHeight(&_EthCrossChainData.TransactOpts, startHeight)
}

// PutCurEpochStartHeight is a paid mutator transaction binding the contract method 0x512feecc.
//
// Solidity: function putCurEpochStartHeight(uint64 startHeight) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) PutCurEpochStartHeight(startHeight uint64) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutCurEpochStartHeight(&_EthCrossChainData.TransactOpts, startHeight)
}

// PutCurEpochValidatorPkBytes is a paid mutator transaction binding the contract method 0x1dc544bf.
//
// Solidity: function putCurEpochValidatorPkBytes(bytes curEpochPkBytes) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) PutCurEpochValidatorPkBytes(opts *bind.TransactOpts, curEpochPkBytes []byte) (*types.Transaction, error) {
	return _EthCrossChainData.contract.Transact(opts, "putCurEpochValidatorPkBytes", curEpochPkBytes)
}

// PutCurEpochValidatorPkBytes is a paid mutator transaction binding the contract method 0x1dc544bf.
//
// Solidity: function putCurEpochValidatorPkBytes(bytes curEpochPkBytes) returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) PutCurEpochValidatorPkBytes(curEpochPkBytes []byte) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutCurEpochValidatorPkBytes(&_EthCrossChainData.TransactOpts, curEpochPkBytes)
}

// PutCurEpochValidatorPkBytes is a paid mutator transaction binding the contract method 0x1dc544bf.
//
// Solidity: function putCurEpochValidatorPkBytes(bytes curEpochPkBytes) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) PutCurEpochValidatorPkBytes(curEpochPkBytes []byte) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutCurEpochValidatorPkBytes(&_EthCrossChainData.TransactOpts, curEpochPkBytes)
}

// PutEthTxHash is a paid mutator transaction binding the contract method 0x4c3ccf64.
//
// Solidity: function putEthTxHash(bytes32 ethTxHash) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) PutEthTxHash(opts *bind.TransactOpts, ethTxHash [32]byte) (*types.Transaction, error) {
	return _EthCrossChainData.contract.Transact(opts, "putEthTxHash", ethTxHash)
}

// PutEthTxHash is a paid mutator transaction binding the contract method 0x4c3ccf64.
//
// Solidity: function putEthTxHash(bytes32 ethTxHash) returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) PutEthTxHash(ethTxHash [32]byte) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutEthTxHash(&_EthCrossChainData.TransactOpts, ethTxHash)
}

// PutEthTxHash is a paid mutator transaction binding the contract method 0x4c3ccf64.
//
// Solidity: function putEthTxHash(bytes32 ethTxHash) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) PutEthTxHash(ethTxHash [32]byte) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutEthTxHash(&_EthCrossChainData.TransactOpts, ethTxHash)
}

// PutExtraData is a paid mutator transaction binding the contract method 0x1afe374e.
//
// Solidity: function putExtraData(bytes32 key1, bytes32 key2, bytes value) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) PutExtraData(opts *bind.TransactOpts, key1 [32]byte, key2 [32]byte, value []byte) (*types.Transaction, error) {
	return _EthCrossChainData.contract.Transact(opts, "putExtraData", key1, key2, value)
}

// PutExtraData is a paid mutator transaction binding the contract method 0x1afe374e.
//
// Solidity: function putExtraData(bytes32 key1, bytes32 key2, bytes value) returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) PutExtraData(key1 [32]byte, key2 [32]byte, value []byte) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutExtraData(&_EthCrossChainData.TransactOpts, key1, key2, value)
}

// PutExtraData is a paid mutator transaction binding the contract method 0x1afe374e.
//
// Solidity: function putExtraData(bytes32 key1, bytes32 key2, bytes value) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) PutExtraData(key1 [32]byte, key2 [32]byte, value []byte) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutExtraData(&_EthCrossChainData.TransactOpts, key1, key2, value)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthCrossChainData *EthCrossChainDataTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainData.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthCrossChainData *EthCrossChainDataSession) RenounceOwnership() (*types.Transaction, error) {
	return _EthCrossChainData.Contract.RenounceOwnership(&_EthCrossChainData.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthCrossChainData *EthCrossChainDataTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EthCrossChainData.Contract.RenounceOwnership(&_EthCrossChainData.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthCrossChainData *EthCrossChainDataTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EthCrossChainData.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthCrossChainData *EthCrossChainDataSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.TransferOwnership(&_EthCrossChainData.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthCrossChainData *EthCrossChainDataTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.TransferOwnership(&_EthCrossChainData.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainData.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) Unpause() (*types.Transaction, error) {
	return _EthCrossChainData.Contract.Unpause(&_EthCrossChainData.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) Unpause() (*types.Transaction, error) {
	return _EthCrossChainData.Contract.Unpause(&_EthCrossChainData.TransactOpts)
}

// EthCrossChainDataOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EthCrossChainData contract.
type EthCrossChainDataOwnershipTransferredIterator struct {
	Event *EthCrossChainDataOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EthCrossChainDataOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainDataOwnershipTransferred)
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
		it.Event = new(EthCrossChainDataOwnershipTransferred)
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
func (it *EthCrossChainDataOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainDataOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainDataOwnershipTransferred represents a OwnershipTransferred event raised by the EthCrossChainData contract.
type EthCrossChainDataOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EthCrossChainData *EthCrossChainDataFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EthCrossChainDataOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EthCrossChainData.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainDataOwnershipTransferredIterator{contract: _EthCrossChainData.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EthCrossChainData *EthCrossChainDataFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EthCrossChainDataOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EthCrossChainData.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainDataOwnershipTransferred)
				if err := _EthCrossChainData.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EthCrossChainData *EthCrossChainDataFilterer) ParseOwnershipTransferred(log types.Log) (*EthCrossChainDataOwnershipTransferred, error) {
	event := new(EthCrossChainDataOwnershipTransferred)
	if err := _EthCrossChainData.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthCrossChainDataPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the EthCrossChainData contract.
type EthCrossChainDataPausedIterator struct {
	Event *EthCrossChainDataPaused // Event containing the contract specifics and raw log

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
func (it *EthCrossChainDataPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainDataPaused)
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
		it.Event = new(EthCrossChainDataPaused)
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
func (it *EthCrossChainDataPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainDataPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainDataPaused represents a Paused event raised by the EthCrossChainData contract.
type EthCrossChainDataPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EthCrossChainData *EthCrossChainDataFilterer) FilterPaused(opts *bind.FilterOpts) (*EthCrossChainDataPausedIterator, error) {

	logs, sub, err := _EthCrossChainData.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EthCrossChainDataPausedIterator{contract: _EthCrossChainData.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EthCrossChainData *EthCrossChainDataFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EthCrossChainDataPaused) (event.Subscription, error) {

	logs, sub, err := _EthCrossChainData.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainDataPaused)
				if err := _EthCrossChainData.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_EthCrossChainData *EthCrossChainDataFilterer) ParsePaused(log types.Log) (*EthCrossChainDataPaused, error) {
	event := new(EthCrossChainDataPaused)
	if err := _EthCrossChainData.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthCrossChainDataUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the EthCrossChainData contract.
type EthCrossChainDataUnpausedIterator struct {
	Event *EthCrossChainDataUnpaused // Event containing the contract specifics and raw log

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
func (it *EthCrossChainDataUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainDataUnpaused)
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
		it.Event = new(EthCrossChainDataUnpaused)
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
func (it *EthCrossChainDataUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainDataUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainDataUnpaused represents a Unpaused event raised by the EthCrossChainData contract.
type EthCrossChainDataUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EthCrossChainData *EthCrossChainDataFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EthCrossChainDataUnpausedIterator, error) {

	logs, sub, err := _EthCrossChainData.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EthCrossChainDataUnpausedIterator{contract: _EthCrossChainData.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EthCrossChainData *EthCrossChainDataFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EthCrossChainDataUnpaused) (event.Subscription, error) {

	logs, sub, err := _EthCrossChainData.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainDataUnpaused)
				if err := _EthCrossChainData.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_EthCrossChainData *EthCrossChainDataFilterer) ParseUnpaused(log types.Log) (*EthCrossChainDataUnpaused, error) {
	event := new(EthCrossChainDataUnpaused)
	if err := _EthCrossChainData.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthCrossChainManagerImplementationABI is the input ABI used to generate the binding from.
const EthCrossChainManagerImplementationABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rawHeader\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"oldValidators\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"newValidators\",\"type\":\"address[]\"}],\"name\":\"ChangeEpochEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"txId\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxyOrAssetContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toContract\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rawdata\",\"type\":\"bytes\"}],\"name\":\"CrossChainEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rawHeader\",\"type\":\"bytes\"}],\"name\":\"InitGenesisBlockEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fromChainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toContract\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"crossChainTxHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fromChainTxHash\",\"type\":\"bytes\"}],\"name\":\"VerifyHeaderAndExecuteTxEvent\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rawHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rawSeals\",\"type\":\"bytes\"}],\"name\":\"changeEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"toContract\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"method\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"txData\",\"type\":\"bytes\"}],\"name\":\"crossChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"fallback\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEthCrossChainCallerFactoryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEthCrossChainDataAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEthCrossChainManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getZionChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rawHeader\",\"type\":\"bytes\"}],\"name\":\"initGenesisBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rawHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rawSeals\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"accountProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"storageProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rawCrossTx\",\"type\":\"bytes\"}],\"name\":\"verifyHeaderAndExecuteTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EthCrossChainManagerImplementationFuncSigs maps the 4-byte function signature to its string representation.
var EthCrossChainManagerImplementationFuncSigs = map[string]string{
	"e15d450d": "changeEpoch(bytes,bytes)",
	"bd5cf625": "crossChain(uint64,bytes,bytes,bytes)",
	"552079dc": "fallback()",
	"1786794b": "getEthCrossChainCallerFactoryAddress()",
	"2de23f6b": "getEthCrossChainDataAddress()",
	"87939a7f": "getEthCrossChainManager()",
	"711ccc13": "getZionChainId()",
	"3cc22d01": "initGenesisBlock(bytes)",
	"d450e04c": "verifyHeaderAndExecuteTx(bytes,bytes,bytes,bytes,bytes)",
}

// EthCrossChainManagerImplementationBin is the compiled bytecode used for deploying new contracts.
var EthCrossChainManagerImplementationBin = "0x608060405234801561001057600080fd5b50614183806100206000396000f3fe6080604052600436106100865760003560e01c8063711ccc1311610059578063711ccc13146101a057806387939a7f146101c7578063bd5cf625146101dc578063d450e04c14610307578063e15d450d146105cc57610086565b80631786794b1461008b5780632de23f6b146100bc5780633cc22d01146100d1578063552079dc14610196575b600080fd5b34801561009757600080fd5b506100a0610702565b604080516001600160a01b039092168252519081900360200190f35b3480156100c857600080fd5b506100a061071a565b3480156100dd57600080fd5b50610182600480360360208110156100f457600080fd5b810190602081018135600160201b81111561010e57600080fd5b82018360208201111561012057600080fd5b803590602001918460018302840111600160201b8311171561014157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610732945050505050565b604080519115158252519081900360200190f35b61019e610c5b565b005b3480156101ac57600080fd5b506101b5610c9f565b60408051918252519081900360200190f35b3480156101d357600080fd5b506100a0610ca4565b3480156101e857600080fd5b50610182600480360360808110156101ff57600080fd5b6001600160401b038235169190810190604081016020820135600160201b81111561022957600080fd5b82018360208201111561023b57600080fd5b803590602001918460018302840111600160201b8311171561025c57600080fd5b919390929091602081019035600160201b81111561027957600080fd5b82018360208201111561028b57600080fd5b803590602001918460018302840111600160201b831117156102ac57600080fd5b919390929091602081019035600160201b8111156102c957600080fd5b8201836020820111156102db57600080fd5b803590602001918460018302840111600160201b831117156102fc57600080fd5b509092509050610ca8565b34801561031357600080fd5b50610182600480360360a081101561032a57600080fd5b810190602081018135600160201b81111561034457600080fd5b82018360208201111561035657600080fd5b803590602001918460018302840111600160201b8311171561037757600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156103c957600080fd5b8201836020820111156103db57600080fd5b803590602001918460018302840111600160201b831117156103fc57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561044e57600080fd5b82018360208201111561046057600080fd5b803590602001918460018302840111600160201b8311171561048157600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156104d357600080fd5b8201836020820111156104e557600080fd5b803590602001918460018302840111600160201b8311171561050657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561055857600080fd5b82018360208201111561056a57600080fd5b803590602001918460018302840111600160201b8311171561058b57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061123b945050505050565b3480156105d857600080fd5b50610182600480360360408110156105ef57600080fd5b810190602081018135600160201b81111561060957600080fd5b82018360208201111561061b57600080fd5b803590602001918460018302840111600160201b8311171561063c57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561068e57600080fd5b8201836020820111156106a057600080fd5b803590602001918460018302840111600160201b831117156106c157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611a55945050505050565b73ded94feafaaca7064ce45ff875695f5ad3dec3e890565b73d2bfdba1ab74f287c856df3a550ff6067a302a7890565b600061073c613c6a565b610745836120e9565b9050600073d2bfdba1ab74f287c856df3a550ff6067a302a789050806001600160a01b031663cd62908f6040518163ffffffff1660e01b815260040160006040518083038186803b15801561079957600080fd5b505afa1580156107ad573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156107d657600080fd5b8101908080516040519392919084600160201b8211156107f557600080fd5b90830190602082018581111561080a57600080fd5b8251600160201b81118282018810171561082357600080fd5b82525081516020918201929091019080838360005b83811015610850578181015183820152602001610838565b50505050905090810190601f16801561087d5780820380516001836020036101000a031916815260200191505b50604052505050516000146108c35760405162461bcd60e51b81526004018080602001828103825260388152602001806140656038913960400191505060405180910390fd5b600060606108d08661213f565b915091508051600014156109155760405162461bcd60e51b8152600401808060200182810382526031815260200180613cef6031913960400191505060405180910390fd5b6040805163353b9cb160e11b81526001600160401b038416600482015290516001600160a01b03851691636a7739629160248083019260209291908290030181600087803b15801561096657600080fd5b505af115801561097a573d6000803e3d6000fd5b505050506040513d602081101561099057600080fd5b50516109cd5760405162461bcd60e51b815260040180806020018281038252603b815260200180613ebd603b913960400191505060405180910390fd5b826001600160a01b031663512feecc85602001516001016040518263ffffffff1660e01b815260040180826001600160401b03166001600160401b03168152602001915050602060405180830381600087803b158015610a2c57600080fd5b505af1158015610a40573d6000803e3d6000fd5b505050506040513d6020811015610a5657600080fd5b5051610a935760405162461bcd60e51b815260040180806020018281038252603d815260200180613d6f603d913960400191505060405180910390fd5b826001600160a01b0316631dc544bf610aab836122ca565b6040518263ffffffff1660e01b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610af7578181015183820152602001610adf565b50505050905090810190601f168015610b245780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b158015610b4357600080fd5b505af1158015610b57573d6000803e3d6000fd5b505050506040513d6020811015610b6d57600080fd5b5051610baa5760405162461bcd60e51b815260040180806020018281038252603b815260200180613dac603b913960400191505060405180910390fd5b7ff01968fc3a2655cf1b5144cb32de6dc898f91b9239c103744e8457152ab2fbde8460200151876040518083815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610c14578181015183820152602001610bfc565b50505050905090810190601f168015610c415780820380516001836020036101000a031916815260200191505b50935050505060405180910390a150600195945050505050565b6040805162461bcd60e51b81526020600482015260146024820152732ab739bab83837b93a32b210333ab731ba34b7b760611b604482015290519081900360640190fd5b604f90565b3090565b6040805163fc91a89760e01b8152336004820152905160009173ded94feafaaca7064ce45ff875695f5ad3dec3e89163fc91a89791602480820192602092909190829003018186803b158015610cfd57600080fd5b505afa158015610d11573d6000803e3d6000fd5b505050506040513d6020811015610d2757600080fd5b5051610d645760405162461bcd60e51b815260040180806020018281038252602a815260200180613e4c602a913960400191505060405180910390fd5b600073d2bfdba1ab74f287c856df3a550ff6067a302a786001600160a01b031663ff3d24a16040518163ffffffff1660e01b815260040160206040518083038186803b158015610db357600080fd5b505afa158015610dc7573d6000803e3d6000fd5b505050506040513d6020811015610ddd57600080fd5b505190506060610dec82612331565b905060606002308360405160200180836001600160a01b03166001600160a01b031660601b815260140182805190602001908083835b60208310610e415780518252601f199092019160209182019101610e22565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040526040518082805190602001908083835b60208310610ea55780518252601f199092019160209182019101610e86565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015610ee4573d6000803e3d6000fd5b5050506040513d6020811015610ef957600080fd5b505160408051602081810193909352815180820390930183528101905290506060610fef8383610f28336123aa565b8f8f8f8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508e8e8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508d8d8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506123c592505050565b905073d2bfdba1ab74f287c856df3a550ff6067a302a786001600160a01b0316634c3ccf6482805190602001206040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b15801561105257600080fd5b505af1158015611066573d6000803e3d6000fd5b505050506040513d602081101561107c57600080fd5b50516110b95760405162461bcd60e51b81526004018080602001828103825260308152602001806140bf6030913960400191505060405180910390fd5b326001600160a01b03167f6ad3bf15c1988bc04bc153490cab16db8efb9a3990215bf1c64ea6e28be8848384338f8f8f876040518080602001876001600160a01b03166001600160a01b03168152602001866001600160401b03166001600160401b03168152602001806020018060200184810384528a818151815260200191508051906020019080838360005b8381101561115f578181015183820152602001611147565b50505050905090810190601f16801561118c5780820380516001836020036101000a031916815260200191505b5084810383528681526020018787808284376000838201819052601f909101601f191690920186810384528751815287516020918201939189019250908190849084905b838110156111e85781810151838201526020016111d0565b50505050905090810190601f1680156112155780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390a25060019b9a5050505050505050505050565b6000611245613c6a565b61124e876120e9565b9050611258613c84565b61126184612673565b9050600073d2bfdba1ab74f287c856df3a550ff6067a302a78905060606113aa826001600160a01b031663cd62908f6040518163ffffffff1660e01b815260040160006040518083038186803b1580156112ba57600080fd5b505afa1580156112ce573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156112f757600080fd5b8101908080516040519392919084600160201b82111561131657600080fd5b90830190602082018581111561132b57600080fd5b8251600160201b81118282018810171561134457600080fd5b82525081516020918201929091019080838360005b83811015611371578181015183820152602001611359565b50505050905090810190601f16801561139e5780820380516001836020036101000a031916815260200191505b506040525050506126dc565b9050816001600160a01b0316635ac407906040518163ffffffff1660e01b815260040160206040518083038186803b1580156113e557600080fd5b505afa1580156113f9573d6000803e3d6000fd5b505050506040513d602081101561140f57600080fd5b505160208501516001600160401b03909116111561145e5760405162461bcd60e51b8152600401808060200182810382526030815260200180613ef86030913960400191505060405180910390fd5b816001600160a01b0316634e4af8ce6040518163ffffffff1660e01b815260040160206040518083038186803b15801561149757600080fd5b505afa1580156114ab573d6000803e3d6000fd5b505050506040513d60208110156114c157600080fd5b505160208501516001600160401b0390911610156115105760405162461bcd60e51b815260040180806020018281038252602d815260200180613e1f602d913960400191505060405180910390fd5b6115228a805190602001208a83612784565b61156a576040805162461bcd60e51b815260206004820152601460248201527315995c9a599e481a195859195c8819985a5b195960621b604482015290519081900360640190fd5b6060611575846128ad565b905060606115a48a876000015160405180604001604052806014815260200161100360601b8152508c86612967565b905087805190602001206115b7826129fa565b146115ff576040805162461bcd60e51b815260206004820152601360248201527215995c9a599e481c1c9bdbd98819985a5b1959606a1b604482015290519081900360640190fd5b836001600160a01b0316630586763c86602001516116248860400151602001516129fa565b6040518363ffffffff1660e01b815260040180836001600160401b03166001600160401b031681526020018281526020019250505060206040518083038186803b15801561167157600080fd5b505afa158015611685573d6000803e3d6000fd5b505050506040513d602081101561169b57600080fd5b5051156116d95760405162461bcd60e51b815260040180806020018281038252602281526020018061409d6022913960400191505060405180910390fd5b836001600160a01b031663e90bfdcf86602001516116fe8860400151602001516129fa565b6040518363ffffffff1660e01b815260040180836001600160401b03166001600160401b0316815260200182815260200192505050602060405180830381600087803b15801561174d57600080fd5b505af1158015611761573d6000803e3d6000fd5b505050506040513d602081101561177757600080fd5b50516117ca576040805162461bcd60e51b815260206004820181905260248201527f536176652063726f7373636861696e207478206578697374206661696c656421604482015290519081900360640190fd5b604f8560400151606001516001600160401b03161461181a5760405162461bcd60e51b8152600401808060200182810382526026815260200180613e766026913960400191505060405180910390fd5b600061182d866040015160800151612a5a565b905061185881876040015160a00151886040015160c001518960400151604001518a60200151612aa4565b6118a9576040805162461bcd60e51b815260206004820152601d60248201527f457865637574652043726f7373436861696e205478206661696c656421000000604482015290519081900360640190fd5b7f8a4a2663ce60ce4955c595da2894de0415240f1ace024cfbff85f513b656bdae8660200151876040015160800151886000015189604001516000015160405180856001600160401b03166001600160401b03168152602001806020018060200180602001848103845287818151815260200191508051906020019080838360005b8381101561194357818101518382015260200161192b565b50505050905090810190601f1680156119705780820380516001836020036101000a031916815260200191505b50848103835286518152865160209182019188019080838360005b838110156119a357818101518382015260200161198b565b50505050905090810190601f1680156119d05780820380516001836020036101000a031916815260200191505b50848103825285518152855160209182019187019080838360005b83811015611a035781810151838201526020016119eb565b50505050905090810190601f168015611a305780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a15060019c9b505050505050505050505050565b6000611a5f613c6a565b611a68846120e9565b9050600073d2bfdba1ab74f287c856df3a550ff6067a302a789050806001600160a01b0316635ac407906040518163ffffffff1660e01b815260040160206040518083038186803b158015611abc57600080fd5b505afa158015611ad0573d6000803e3d6000fd5b505050506040513d6020811015611ae657600080fd5b505160208301516001600160401b039091161115611b355760405162461bcd60e51b815260040180806020018281038252603b815260200180613fa2603b913960400191505060405180910390fd5b6060816001600160a01b031663cd62908f6040518163ffffffff1660e01b815260040160006040518083038186803b158015611b7057600080fd5b505afa158015611b84573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015611bad57600080fd5b8101908080516040519392919084600160201b821115611bcc57600080fd5b908301906020820185811115611be157600080fd5b8251600160201b811182820188101715611bfa57600080fd5b82525081516020918201929091019080838360005b83811015611c27578181015183820152602001611c0f565b50505050905090810190601f168015611c545780820380516001836020036101000a031916815260200191505b5060405250505090506060611c68826126dc565b905060006060611c778961213f565b91509150805160001415611cbc5760405162461bcd60e51b8152600401808060200182810382526031815260200180613cef6031913960400191505060405180910390fd5b611cce89805190602001208985612784565b611d16576040805162461bcd60e51b815260206004820152601460248201527315995c9a599e481a195859195c8819985a5b195960621b604482015290519081900360640190fd5b6040805163353b9cb160e11b81526001600160401b038416600482015290516001600160a01b03871691636a7739629160248083019260209291908290030181600087803b158015611d6757600080fd5b505af1158015611d7b573d6000803e3d6000fd5b505050506040513d6020811015611d9157600080fd5b5051611dce5760405162461bcd60e51b815260040180806020018281038252603b815260200180613ebd603b913960400191505060405180910390fd5b846001600160a01b031663512feecc87602001516001016040518263ffffffff1660e01b815260040180826001600160401b03166001600160401b03168152602001915050602060405180830381600087803b158015611e2d57600080fd5b505af1158015611e41573d6000803e3d6000fd5b505050506040513d6020811015611e5757600080fd5b5051611e945760405162461bcd60e51b81526004018080602001828103825260348152602001806140ef6034913960400191505060405180910390fd5b846001600160a01b0316631dc544bf611eac836122ca565b6040518263ffffffff1660e01b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611ef8578181015183820152602001611ee0565b50505050905090810190601f168015611f255780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b158015611f4457600080fd5b505af1158015611f58573d6000803e3d6000fd5b505050506040513d6020811015611f6e57600080fd5b5051611fab5760405162461bcd60e51b8152600401808060200182810382526038815260200180613de76038913960400191505060405180910390fd5b7f353f5ab7dd2715c6a221c2e7ff95ca17e68d962ffd1bb9b97b8f283d3659896386602001518a858460405180858152602001806020018060200180602001848103845287818151815260200191508051906020019080838360005b8381101561201f578181015183820152602001612007565b50505050905090810190601f16801561204c5780820380516001836020036101000a031916815260200191505b508481038352865181528651602091820191808901910280838360005b83811015612081578181015183820152602001612069565b50505050905001848103825285818151815260200191508051906020019060200280838360005b838110156120c05781810151838201526020016120a8565b5050505090500197505050505050505060405180910390a1600196505050505050505b92915050565b6120f1613c6a565b6000806120ff846020612f75565b91505061210f8482605701613014565b508352612120846101bd8301612f75565b909250905061213184838301613062565b506020840152509092915050565b60006060600080612151856020612f75565b91505061216285826101bd01612f75565b909250905061217385838301612f75565b909250905061218485838301612f75565b909250905061219585838301612f75565b90925090506121a685838301612f75565b90925090506121b785838301612f75565b91506121c890508560208301612f75565b90925090506121d78582612f75565b90925090506121e8858383016130d9565b909450905060606121f98683613014565b5080519350905060158306156122405760405162461bcd60e51b8152600401808060200182810382526021815260200180613e9c6021913960400191505060405180910390fd5b6015830460405190808252806020026020018201604052801561226d578160200160208202803883390190505b50935060005b836015820210156122c1576000612290836015840260200161312c565b509050808683815181106122a057fe5b6001600160a01b039092166020928302919091019091015250600101612273565b50505050915091565b6060816040516020018080602001828103825283818151815260200191508051906020019060200280838360005b838110156123105781810151838201526020016122f8565b50505050905001925050506040516020818303038152906040529050919050565b60606001600160ff1b0382111561238f576040805162461bcd60e51b815260206004820152601760248201527f56616c75652065786365656473207468652072616e6765000000000000000000604482015290519081900360640190fd5b60405190506020815281602082015260408101604052919050565b604080516014815260609290921b6020830152818101905290565b60608787878787878760405160200180806020018060200180602001886001600160401b03166001600160401b0316815260200180602001806020018060200187810387528e818151815260200191508051906020019080838360005b8381101561243a578181015183820152602001612422565b50505050905090810190601f1680156124675780820380516001836020036101000a031916815260200191505b5087810386528d5181528d516020918201918f019080838360005b8381101561249a578181015183820152602001612482565b50505050905090810190601f1680156124c75780820380516001836020036101000a031916815260200191505b5087810385528c5181528c516020918201918e019080838360005b838110156124fa5781810151838201526020016124e2565b50505050905090810190601f1680156125275780820380516001836020036101000a031916815260200191505b5087810384528a5181528a516020918201918c019080838360005b8381101561255a578181015183820152602001612542565b50505050905090810190601f1680156125875780820380516001836020036101000a031916815260200191505b5087810383528951815289516020918201918b019080838360005b838110156125ba5781810151838201526020016125a2565b50505050905090810190601f1680156125e75780820380516001836020036101000a031916815260200191505b5087810382528851815288516020918201918a019080838360005b8381101561261a578181015183820152602001612602565b50505050905090810190601f1680156126475780820380516001836020036101000a031916815260200191505b509d50505050505050505050505050506040516020818303038152906040529050979650505050505050565b61267b613c84565b60206126878382612f75565b9150606090506126978483613014565b90845291506126a684836130d9565b6001600160401b03909116602085015291506126c28483613014565b925090506126cf8161317f565b6040840152509092915050565b60608180602001905160208110156126f357600080fd5b8101908080516040519392919084600160201b82111561271257600080fd5b90830190602082018581111561272757600080fd5b82518660208202830111600160201b8211171561274357600080fd5b82525081516020918201928201910280838360005b83811015612770578181015183820152602001612758565b505050509050016040525050509050919050565b600060206060612794858361321a565b5080519095506000906127ae90604363ffffffff61323e16565b90506060816040519080825280602002602001820160405280156127dc578160200160208202803883390190505b50905060005b82811015612894576127f4888661321a565b6040805160208082018e905282518083038201815291830190925280519101209096509094506128249085613280565b82828151811061283057fe5b60200260200101906001600160a01b031690816001600160a01b03168152505060006001600160a01b031682828151811061286757fe5b60200260200101516001600160a01b0316141561288c576000955050505050506128a6565b6001016127e2565b5061289f8682613325565b9450505050505b9392505050565b60606120e3667265717565737460c81b6128ce846040015160600151613451565b84516040516001600160c81b0319841660208083019182526001600160c01b03198516602784015283519192602f0191908401908083835b602083106129255780518252601f199092019160209182019101612906565b6001836020036101000a0380198251168184511680821785525050505050509050019350505050604051602081830303815290604052805190602001206134c6565b825160208085019182209091528352606083816129858883896134de565b905061299281602061321a565b50905060006129a282602061321a565b9150506129af828261321a565b9150606090506129bf838361321a565b508651602080890191822090915287529050856129dd8882846134de565b95506129ea86602061321a565b509b9a5050505050505050505050565b60008151602014612a52576040805162461bcd60e51b815260206004820152601760248201527f6279746573206c656e677468206973206e6f742033322e000000000000000000604482015290519081900360640190fd5b506020015190565b60008151601414612a9c5760405162461bcd60e51b8152600401808060200182810382526023815260200180613d206023913960400191505060405180910390fd5b506014015190565b6040805163fc91a89760e01b81526001600160a01b0387166004820152905160009173ded94feafaaca7064ce45ff875695f5ad3dec3e89163fc91a89791602480820192602092909190829003018186803b158015612b0257600080fd5b505afa158015612b16573d6000803e3d6000fd5b505050506040513d6020811015612b2c57600080fd5b5051612b695760405162461bcd60e51b815260040180806020018281038252602e815260200180613f74602e913960400191505060405180910390fd5b6001600160a01b03861673d2bfdba1ab74f287c856df3a550ff6067a302a781415612bdb576040805162461bcd60e51b815260206004820152601760248201527f446f6e27742074727920746f2063616c6c206563636421000000000000000000604482015290519081900360640190fd5b60006060876001600160a01b0316876040516020018082805190602001908083835b60208310612c1c5780518252601f199092019160209182019101612bfd565b51815160001960209485036101000a01908116901991909116179052732862797465732c62797465732c75696e7436342960601b9390910192835260408051600b19818603018152601485019091528051908201206001600160401b038b1660748501526060603485019081528d5160948601528d519196508d95508c948c945090928392605483019260b4019188019080838360005b83811015612ccb578181015183820152602001612cb3565b50505050905090810190601f168015612cf85780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b83811015612d2b578181015183820152602001612d13565b50505050905090810190601f168015612d585780820380516001836020036101000a031916815260200191505b509550505050505060405160208183030381529060405260405160200180836001600160e01b0319166001600160e01b031916815260040182805190602001908083835b60208310612dbb5780518252601f199092019160209182019101612d9c565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040526040518082805190602001908083835b60208310612e1f5780518252601f199092019160209182019101612e00565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114612e81576040519150601f19603f3d011682016040523d82523d6000602084013e612e86565b606091505b509092509050600182151514612ecd5760405162461bcd60e51b815260040180806020018281038252602b815260200180613f28602b913960400191505060405180910390fd5b8051612f0a5760405162461bcd60e51b8152600401808060200182810382526027815260200180613fdd6027913960400191505060405180910390fd5b6000818060200190516020811015612f2157600080fd5b50519050600181151514612f665760405162461bcd60e51b815260040180806020018281038252603781526020018061402e6037913960400191505060405180910390fd5b50600198975050505050505050565b81810151600090819060f81c6080811015612f9757506001915082905061300d565b60b7811015612fb057607f19019150506001820161300d565b60c0811015612fd85784840160010151600860d7839003021c9250830160b51901905061300d565b60f7811015612ff15760bf19019150506001820161300d565b848401600101516008610117839003021c9250830160f5190190505b9250929050565b60606000806130238585612f75565b8095508192505050838101915060405192506020601f82010460200260200180840160405281845281602085018387890160045afa5050509250929050565b60008060006130718585612f75565b9450905060208111156130b55760405162461bcd60e51b815260040180806020018281038252602c815260200180614123602c913960400191505060405180910390fd5b838101915080602003808503860151600882021b600882021c935050509250929050565b60008060006130e88585612f75565b9450905060088111156130b55760405162461bcd60e51b815260040180806020018281038252602a815260200180614004602a913960400191505060405180910390fd5b600080600061313b8585612f75565b9450905060148111156130b55760405162461bcd60e51b815260040180806020018281038252602c815260200180613d43602c913960400191505060405180910390fd5b613187613ca8565b60206131938382613014565b90835290506131a28382613014565b602084019190915290506131b68382613014565b604084019190915290506131ca83826130d9565b6001600160401b03909116606084015290506131e68382613014565b608084019190915290506131fa8382613014565b60a0840191909152905061320e8382613014565b5060c083015250919050565b60606000806132298585612f75565b958601601f1901818152969501949350505050565b60006128a683836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f0000000000008152506136fd565b60008151604114613293575060006120e3565b60008060006020850151925060408501519150606085015160001a601b01905060018682858560405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015613310573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b60008061335f600161335360036133476002895161379f90919063ffffffff16565b9063ffffffff61323e16565b9063ffffffff6137f816565b90506060845160405190808252806020026020018201604052801561338e578160200160208202803883390190505b5090506000805b85518110156134435760005b875181101561343a578381815181106133b657fe5b602002602001015115801561340257508781815181106133d257fe5b60200260200101516001600160a01b03168783815181106133ef57fe5b60200260200101516001600160a01b0316145b15613432578280600101935050600184828151811061341d57fe5b9115156020928302919091019091015261343a565b6001016133a1565b50600101613395565b509190911015949350505050565b600060ff821660381b905061ff00821660281b8101905062ff0000821660181b8101905063ff000000821660081b8101905064ff00000000821660081c8101905065ff0000000000821660181c8101905066ff000000000000821660281c810190508160381c810190508060c01b9050919050565b60408051808201909152602080825281019190915290565b60606134e983613852565b92506134f684602061321a565b5080519094508291506020905b806020018210156136ac57606061351b87848661394a565b61356c576040805162461bcd60e51b815260206004820152601760248201527f70726f6f663a756e657175616c206e6f64652068617368000000000000000000604482015290519081900360640190fd5b613576878461321a565b93509050600080613588836020612f75565b9150915061359883838301612f75565b84519193508301915060200181141561363957606060006135ba85602061321a565b91509150600060606135cb846139b3565b90506135d78c82613b52565b909c50915081613622576040805162461bcd60e51b815260206004820152601160248201527070726f6f663a696e76616c6964206b657960781b604482015290519081900360640190fd5b61362c878461321a565b5099506136a49350505050565b600088516000141561364d5750601161365c565b61365689613bf6565b90995090505b602060005b60118110156136a057808314156136845761367c868361321a565b5098506136a0565b61368e8683612f75565b90955093508484019150600101613661565b5050505b505050613503565b8451156136f4576040805162461bcd60e51b815260206004820152601160248201527070726f6f663a696e76616c6964206b657960781b604482015290519081900360640190fd5b50509392505050565b600081836137895760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561374e578181015183820152602001613736565b50505050905090810190601f16801561377b5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50600083858161379557fe5b0495945050505050565b6000826137ae575060006120e3565b828202828482816137bb57fe5b04146128a65760405162461bcd60e51b8152600401808060200182810382526021815260200180613f536021913960400191505060405180910390fd5b6000828201838110156128a6576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6060600082516002026001019050604051915060208183010160405280825260008090505b835181101561391857600484828151811061388e57fe5b602001015160f81c60f81b6001600160f81b031916901c8382600202815181106138b457fe5b60200101906001600160f81b031916908160001a9053508381815181106138d757fe5b01602001518351600f60f81b909116908490600160028502019081106138f957fe5b60200101906001600160f81b031916908160001a905350600101613877565b50601060f81b82600183038151811061392d57fe5b60200101906001600160f81b031916908160001a90535050919050565b60008060006139598686612f75565b915091508481038201602081106001811461397957801561399a576139a8565b60208601518789015118826020036008021c158651831481169550506139a8565b602086015182888a01201494505b505050509392505050565b606060006004836000815181106139c657fe5b602001015160f81c60f81b6001600160f81b031916901c9050600083519050600080600060028560f81c06905060028560f81c04915080826002808702030101925060405195506020838701016040528286528060011415613a625786600081518110613a2f57fe5b01602001518651600f60f81b909116908790600090613a4a57fe5b60200101906001600160f81b031916908160001a9053505b8160011415613a9757601060f81b866001850381518110613a7f57fe5b60200101906001600160f81b031916908160001a9053505b60005b60018503811015613b47576004888260010181518110613ab657fe5b602001015160f81c60f81b6001600160f81b031916901c8783836002020181518110613ade57fe5b60200101906001600160f81b031916908160001a905350878160010181518110613b0457fe5b01602001518751600f60f81b909116908890600284028501600101908110613b2857fe5b60200101906001600160f81b031916908160001a905350600101613a9a565b505050505050919050565b60606000825184511015613b6b5750829050600061300d565b60006020845181613b7857fe5b6001935004905060005b81811015613baa57602081810286810182015190880190910151149290921691600101613b82565b5060208102602001600019855182036008021b8186015182880151181615831692505081613bdf57846000925092505061300d565b508251845181900390850152915190920192909150565b60606000825160001415613c51576040805162461bcd60e51b815260206004820152601860248201527f74616b654f6e65427974653a20656d70747920696e7075740000000000000000604482015290519081900360640190fd5b505060018101805191516000190181529160ff90911690565b604051806040016040528060608152602001600081525090565b6040805160608082018352815260006020820152908101613ca3613ca8565b905290565b6040518060e0016040528060608152602001606081526020016060815260200160006001600160401b03168152602001606081526020016060815260200160608152509056fe476976656e20626c6f636b2068656164657220646f6573206e6f7420636f6e7461696e20616e792076616c696461746f726279746573206c656e67746820646f6573206e6f74206d617463682061646472657373726c704765744e657874416464726573733a2064617461206c6f6e676572207468616e20323020627974657353617665205a696f6e2063757272656e742065706f63682073746172742068656967687420746f204461746120636f6e7472616374206661696c65642153617665205a696f6e2063757272656e742065706f63682076616c696461746f727320746f204461746120636f6e7472616374206661696c65642153617665205a696f6e206e6578742065706f63682076616c696461746f727320746f204461746120636f6e7472616374206661696c656421496e76616c696420626c6f636b206865696768742c2061667465722043757245706f6368456e644865696768745468652063616c6c6572206973206368696c64206f66207468652063616c6c657220666163746f72792154686973205478206973206e6f742061696d696e672061742074686973206e6574776f726b21696e76616c6964206865616465722065787472612076616c696461746f7253657453617665205a696f6e2063757272656e742065706f636820656e642068656967687420746f204461746120636f6e7472616374206661696c656421496e76616c696420626c6f636b206865696768742c206265666f72652043757245706f6368537461727448656967687445746843726f7373436861696e2063616c6c20627573696e65737320636f6e7472616374206661696c6564536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f775468652070617373656420696e2061646472657373206973206e6f742066726f6d2074686520666163746f727921476976656e20626c6f636b20686569676874206973206c6f776572207468616e2063757272656e742065706f6368207374617274206865696768744e6f2072657475726e2076616c75652066726f6d20627573696e65737320636f6e747261637421726c704765744e65787455696e7436343a2064617461206c6f6e676572207468616e203820627974657345746843726f7373436861696e2063616c6c20627573696e65737320636f6e74726163742072657475726e206973206e6f74207472756545746843726f7373436861696e4461746120636f6e74726163742068617320616c7265616479206265656e20696e697469616c697a656421746865207472616e73616374696f6e20686173206265656e20657865637574656421536176652065746854784861736820627920696e64657820746f204461746120636f6e7472616374206661696c65642153617665205a696f6e206e6578742065706f63682068656967687420746f204461746120636f6e7472616374206661696c656421726c704765744e65787455696e743235363a2064617461206c6f6e676572207468616e203332206279746573a265627a7a723158207287442e98ed3e7a073e88f5b4940a144ceff6a3f5cd70bdd331ed7edc60ff7f64736f6c63430005110032"

// DeployEthCrossChainManagerImplementation deploys a new Ethereum contract, binding an instance of EthCrossChainManagerImplementation to it.
func DeployEthCrossChainManagerImplementation(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EthCrossChainManagerImplementation, error) {
	parsed, err := abi.JSON(strings.NewReader(EthCrossChainManagerImplementationABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EthCrossChainManagerImplementationBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthCrossChainManagerImplementation{EthCrossChainManagerImplementationCaller: EthCrossChainManagerImplementationCaller{contract: contract}, EthCrossChainManagerImplementationTransactor: EthCrossChainManagerImplementationTransactor{contract: contract}, EthCrossChainManagerImplementationFilterer: EthCrossChainManagerImplementationFilterer{contract: contract}}, nil
}

// EthCrossChainManagerImplementation is an auto generated Go binding around an Ethereum contract.
type EthCrossChainManagerImplementation struct {
	EthCrossChainManagerImplementationCaller     // Read-only binding to the contract
	EthCrossChainManagerImplementationTransactor // Write-only binding to the contract
	EthCrossChainManagerImplementationFilterer   // Log filterer for contract events
}

// EthCrossChainManagerImplementationCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthCrossChainManagerImplementationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthCrossChainManagerImplementationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthCrossChainManagerImplementationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthCrossChainManagerImplementationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthCrossChainManagerImplementationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthCrossChainManagerImplementationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthCrossChainManagerImplementationSession struct {
	Contract     *EthCrossChainManagerImplementation // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                       // Call options to use throughout this session
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// EthCrossChainManagerImplementationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthCrossChainManagerImplementationCallerSession struct {
	Contract *EthCrossChainManagerImplementationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                             // Call options to use throughout this session
}

// EthCrossChainManagerImplementationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthCrossChainManagerImplementationTransactorSession struct {
	Contract     *EthCrossChainManagerImplementationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                             // Transaction auth options to use throughout this session
}

// EthCrossChainManagerImplementationRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthCrossChainManagerImplementationRaw struct {
	Contract *EthCrossChainManagerImplementation // Generic contract binding to access the raw methods on
}

// EthCrossChainManagerImplementationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthCrossChainManagerImplementationCallerRaw struct {
	Contract *EthCrossChainManagerImplementationCaller // Generic read-only contract binding to access the raw methods on
}

// EthCrossChainManagerImplementationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthCrossChainManagerImplementationTransactorRaw struct {
	Contract *EthCrossChainManagerImplementationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthCrossChainManagerImplementation creates a new instance of EthCrossChainManagerImplementation, bound to a specific deployed contract.
func NewEthCrossChainManagerImplementation(address common.Address, backend bind.ContractBackend) (*EthCrossChainManagerImplementation, error) {
	contract, err := bindEthCrossChainManagerImplementation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerImplementation{EthCrossChainManagerImplementationCaller: EthCrossChainManagerImplementationCaller{contract: contract}, EthCrossChainManagerImplementationTransactor: EthCrossChainManagerImplementationTransactor{contract: contract}, EthCrossChainManagerImplementationFilterer: EthCrossChainManagerImplementationFilterer{contract: contract}}, nil
}

// NewEthCrossChainManagerImplementationCaller creates a new read-only instance of EthCrossChainManagerImplementation, bound to a specific deployed contract.
func NewEthCrossChainManagerImplementationCaller(address common.Address, caller bind.ContractCaller) (*EthCrossChainManagerImplementationCaller, error) {
	contract, err := bindEthCrossChainManagerImplementation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerImplementationCaller{contract: contract}, nil
}

// NewEthCrossChainManagerImplementationTransactor creates a new write-only instance of EthCrossChainManagerImplementation, bound to a specific deployed contract.
func NewEthCrossChainManagerImplementationTransactor(address common.Address, transactor bind.ContractTransactor) (*EthCrossChainManagerImplementationTransactor, error) {
	contract, err := bindEthCrossChainManagerImplementation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerImplementationTransactor{contract: contract}, nil
}

// NewEthCrossChainManagerImplementationFilterer creates a new log filterer instance of EthCrossChainManagerImplementation, bound to a specific deployed contract.
func NewEthCrossChainManagerImplementationFilterer(address common.Address, filterer bind.ContractFilterer) (*EthCrossChainManagerImplementationFilterer, error) {
	contract, err := bindEthCrossChainManagerImplementation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerImplementationFilterer{contract: contract}, nil
}

// bindEthCrossChainManagerImplementation binds a generic wrapper to an already deployed contract.
func bindEthCrossChainManagerImplementation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthCrossChainManagerImplementationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthCrossChainManagerImplementation.Contract.EthCrossChainManagerImplementationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainManagerImplementation.Contract.EthCrossChainManagerImplementationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthCrossChainManagerImplementation.Contract.EthCrossChainManagerImplementationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthCrossChainManagerImplementation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainManagerImplementation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthCrossChainManagerImplementation.Contract.contract.Transact(opts, method, params...)
}

// GetEthCrossChainCallerFactoryAddress is a free data retrieval call binding the contract method 0x1786794b.
//
// Solidity: function getEthCrossChainCallerFactoryAddress() pure returns(address)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationCaller) GetEthCrossChainCallerFactoryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthCrossChainManagerImplementation.contract.Call(opts, &out, "getEthCrossChainCallerFactoryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEthCrossChainCallerFactoryAddress is a free data retrieval call binding the contract method 0x1786794b.
//
// Solidity: function getEthCrossChainCallerFactoryAddress() pure returns(address)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationSession) GetEthCrossChainCallerFactoryAddress() (common.Address, error) {
	return _EthCrossChainManagerImplementation.Contract.GetEthCrossChainCallerFactoryAddress(&_EthCrossChainManagerImplementation.CallOpts)
}

// GetEthCrossChainCallerFactoryAddress is a free data retrieval call binding the contract method 0x1786794b.
//
// Solidity: function getEthCrossChainCallerFactoryAddress() pure returns(address)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationCallerSession) GetEthCrossChainCallerFactoryAddress() (common.Address, error) {
	return _EthCrossChainManagerImplementation.Contract.GetEthCrossChainCallerFactoryAddress(&_EthCrossChainManagerImplementation.CallOpts)
}

// GetEthCrossChainDataAddress is a free data retrieval call binding the contract method 0x2de23f6b.
//
// Solidity: function getEthCrossChainDataAddress() pure returns(address)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationCaller) GetEthCrossChainDataAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthCrossChainManagerImplementation.contract.Call(opts, &out, "getEthCrossChainDataAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEthCrossChainDataAddress is a free data retrieval call binding the contract method 0x2de23f6b.
//
// Solidity: function getEthCrossChainDataAddress() pure returns(address)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationSession) GetEthCrossChainDataAddress() (common.Address, error) {
	return _EthCrossChainManagerImplementation.Contract.GetEthCrossChainDataAddress(&_EthCrossChainManagerImplementation.CallOpts)
}

// GetEthCrossChainDataAddress is a free data retrieval call binding the contract method 0x2de23f6b.
//
// Solidity: function getEthCrossChainDataAddress() pure returns(address)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationCallerSession) GetEthCrossChainDataAddress() (common.Address, error) {
	return _EthCrossChainManagerImplementation.Contract.GetEthCrossChainDataAddress(&_EthCrossChainManagerImplementation.CallOpts)
}

// GetEthCrossChainManager is a free data retrieval call binding the contract method 0x87939a7f.
//
// Solidity: function getEthCrossChainManager() view returns(address)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationCaller) GetEthCrossChainManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthCrossChainManagerImplementation.contract.Call(opts, &out, "getEthCrossChainManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEthCrossChainManager is a free data retrieval call binding the contract method 0x87939a7f.
//
// Solidity: function getEthCrossChainManager() view returns(address)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationSession) GetEthCrossChainManager() (common.Address, error) {
	return _EthCrossChainManagerImplementation.Contract.GetEthCrossChainManager(&_EthCrossChainManagerImplementation.CallOpts)
}

// GetEthCrossChainManager is a free data retrieval call binding the contract method 0x87939a7f.
//
// Solidity: function getEthCrossChainManager() view returns(address)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationCallerSession) GetEthCrossChainManager() (common.Address, error) {
	return _EthCrossChainManagerImplementation.Contract.GetEthCrossChainManager(&_EthCrossChainManagerImplementation.CallOpts)
}

// GetZionChainId is a free data retrieval call binding the contract method 0x711ccc13.
//
// Solidity: function getZionChainId() pure returns(uint256)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationCaller) GetZionChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthCrossChainManagerImplementation.contract.Call(opts, &out, "getZionChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetZionChainId is a free data retrieval call binding the contract method 0x711ccc13.
//
// Solidity: function getZionChainId() pure returns(uint256)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationSession) GetZionChainId() (*big.Int, error) {
	return _EthCrossChainManagerImplementation.Contract.GetZionChainId(&_EthCrossChainManagerImplementation.CallOpts)
}

// GetZionChainId is a free data retrieval call binding the contract method 0x711ccc13.
//
// Solidity: function getZionChainId() pure returns(uint256)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationCallerSession) GetZionChainId() (*big.Int, error) {
	return _EthCrossChainManagerImplementation.Contract.GetZionChainId(&_EthCrossChainManagerImplementation.CallOpts)
}

// ChangeEpoch is a paid mutator transaction binding the contract method 0xe15d450d.
//
// Solidity: function changeEpoch(bytes rawHeader, bytes rawSeals) returns(bool)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationTransactor) ChangeEpoch(opts *bind.TransactOpts, rawHeader []byte, rawSeals []byte) (*types.Transaction, error) {
	return _EthCrossChainManagerImplementation.contract.Transact(opts, "changeEpoch", rawHeader, rawSeals)
}

// ChangeEpoch is a paid mutator transaction binding the contract method 0xe15d450d.
//
// Solidity: function changeEpoch(bytes rawHeader, bytes rawSeals) returns(bool)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationSession) ChangeEpoch(rawHeader []byte, rawSeals []byte) (*types.Transaction, error) {
	return _EthCrossChainManagerImplementation.Contract.ChangeEpoch(&_EthCrossChainManagerImplementation.TransactOpts, rawHeader, rawSeals)
}

// ChangeEpoch is a paid mutator transaction binding the contract method 0xe15d450d.
//
// Solidity: function changeEpoch(bytes rawHeader, bytes rawSeals) returns(bool)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationTransactorSession) ChangeEpoch(rawHeader []byte, rawSeals []byte) (*types.Transaction, error) {
	return _EthCrossChainManagerImplementation.Contract.ChangeEpoch(&_EthCrossChainManagerImplementation.TransactOpts, rawHeader, rawSeals)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 toChainId, bytes toContract, bytes method, bytes txData) returns(bool)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationTransactor) CrossChain(opts *bind.TransactOpts, toChainId uint64, toContract []byte, method []byte, txData []byte) (*types.Transaction, error) {
	return _EthCrossChainManagerImplementation.contract.Transact(opts, "crossChain", toChainId, toContract, method, txData)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 toChainId, bytes toContract, bytes method, bytes txData) returns(bool)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationSession) CrossChain(toChainId uint64, toContract []byte, method []byte, txData []byte) (*types.Transaction, error) {
	return _EthCrossChainManagerImplementation.Contract.CrossChain(&_EthCrossChainManagerImplementation.TransactOpts, toChainId, toContract, method, txData)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 toChainId, bytes toContract, bytes method, bytes txData) returns(bool)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationTransactorSession) CrossChain(toChainId uint64, toContract []byte, method []byte, txData []byte) (*types.Transaction, error) {
	return _EthCrossChainManagerImplementation.Contract.CrossChain(&_EthCrossChainManagerImplementation.TransactOpts, toChainId, toContract, method, txData)
}

// Fallback is a paid mutator transaction binding the contract method 0x552079dc.
//
// Solidity: function fallback() payable returns()
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationTransactor) Fallback(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainManagerImplementation.contract.Transact(opts, "fallback")
}

// Fallback is a paid mutator transaction binding the contract method 0x552079dc.
//
// Solidity: function fallback() payable returns()
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationSession) Fallback() (*types.Transaction, error) {
	return _EthCrossChainManagerImplementation.Contract.Fallback(&_EthCrossChainManagerImplementation.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract method 0x552079dc.
//
// Solidity: function fallback() payable returns()
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationTransactorSession) Fallback() (*types.Transaction, error) {
	return _EthCrossChainManagerImplementation.Contract.Fallback(&_EthCrossChainManagerImplementation.TransactOpts)
}

// InitGenesisBlock is a paid mutator transaction binding the contract method 0x3cc22d01.
//
// Solidity: function initGenesisBlock(bytes rawHeader) returns(bool)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationTransactor) InitGenesisBlock(opts *bind.TransactOpts, rawHeader []byte) (*types.Transaction, error) {
	return _EthCrossChainManagerImplementation.contract.Transact(opts, "initGenesisBlock", rawHeader)
}

// InitGenesisBlock is a paid mutator transaction binding the contract method 0x3cc22d01.
//
// Solidity: function initGenesisBlock(bytes rawHeader) returns(bool)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationSession) InitGenesisBlock(rawHeader []byte) (*types.Transaction, error) {
	return _EthCrossChainManagerImplementation.Contract.InitGenesisBlock(&_EthCrossChainManagerImplementation.TransactOpts, rawHeader)
}

// InitGenesisBlock is a paid mutator transaction binding the contract method 0x3cc22d01.
//
// Solidity: function initGenesisBlock(bytes rawHeader) returns(bool)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationTransactorSession) InitGenesisBlock(rawHeader []byte) (*types.Transaction, error) {
	return _EthCrossChainManagerImplementation.Contract.InitGenesisBlock(&_EthCrossChainManagerImplementation.TransactOpts, rawHeader)
}

// VerifyHeaderAndExecuteTx is a paid mutator transaction binding the contract method 0xd450e04c.
//
// Solidity: function verifyHeaderAndExecuteTx(bytes rawHeader, bytes rawSeals, bytes accountProof, bytes storageProof, bytes rawCrossTx) returns(bool)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationTransactor) VerifyHeaderAndExecuteTx(opts *bind.TransactOpts, rawHeader []byte, rawSeals []byte, accountProof []byte, storageProof []byte, rawCrossTx []byte) (*types.Transaction, error) {
	return _EthCrossChainManagerImplementation.contract.Transact(opts, "verifyHeaderAndExecuteTx", rawHeader, rawSeals, accountProof, storageProof, rawCrossTx)
}

// VerifyHeaderAndExecuteTx is a paid mutator transaction binding the contract method 0xd450e04c.
//
// Solidity: function verifyHeaderAndExecuteTx(bytes rawHeader, bytes rawSeals, bytes accountProof, bytes storageProof, bytes rawCrossTx) returns(bool)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationSession) VerifyHeaderAndExecuteTx(rawHeader []byte, rawSeals []byte, accountProof []byte, storageProof []byte, rawCrossTx []byte) (*types.Transaction, error) {
	return _EthCrossChainManagerImplementation.Contract.VerifyHeaderAndExecuteTx(&_EthCrossChainManagerImplementation.TransactOpts, rawHeader, rawSeals, accountProof, storageProof, rawCrossTx)
}

// VerifyHeaderAndExecuteTx is a paid mutator transaction binding the contract method 0xd450e04c.
//
// Solidity: function verifyHeaderAndExecuteTx(bytes rawHeader, bytes rawSeals, bytes accountProof, bytes storageProof, bytes rawCrossTx) returns(bool)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationTransactorSession) VerifyHeaderAndExecuteTx(rawHeader []byte, rawSeals []byte, accountProof []byte, storageProof []byte, rawCrossTx []byte) (*types.Transaction, error) {
	return _EthCrossChainManagerImplementation.Contract.VerifyHeaderAndExecuteTx(&_EthCrossChainManagerImplementation.TransactOpts, rawHeader, rawSeals, accountProof, storageProof, rawCrossTx)
}

// EthCrossChainManagerImplementationChangeEpochEventIterator is returned from FilterChangeEpochEvent and is used to iterate over the raw logs and unpacked data for ChangeEpochEvent events raised by the EthCrossChainManagerImplementation contract.
type EthCrossChainManagerImplementationChangeEpochEventIterator struct {
	Event *EthCrossChainManagerImplementationChangeEpochEvent // Event containing the contract specifics and raw log

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
func (it *EthCrossChainManagerImplementationChangeEpochEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainManagerImplementationChangeEpochEvent)
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
		it.Event = new(EthCrossChainManagerImplementationChangeEpochEvent)
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
func (it *EthCrossChainManagerImplementationChangeEpochEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainManagerImplementationChangeEpochEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainManagerImplementationChangeEpochEvent represents a ChangeEpochEvent event raised by the EthCrossChainManagerImplementation contract.
type EthCrossChainManagerImplementationChangeEpochEvent struct {
	Height        *big.Int
	RawHeader     []byte
	OldValidators []common.Address
	NewValidators []common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterChangeEpochEvent is a free log retrieval operation binding the contract event 0x353f5ab7dd2715c6a221c2e7ff95ca17e68d962ffd1bb9b97b8f283d36598963.
//
// Solidity: event ChangeEpochEvent(uint256 height, bytes rawHeader, address[] oldValidators, address[] newValidators)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationFilterer) FilterChangeEpochEvent(opts *bind.FilterOpts) (*EthCrossChainManagerImplementationChangeEpochEventIterator, error) {

	logs, sub, err := _EthCrossChainManagerImplementation.contract.FilterLogs(opts, "ChangeEpochEvent")
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerImplementationChangeEpochEventIterator{contract: _EthCrossChainManagerImplementation.contract, event: "ChangeEpochEvent", logs: logs, sub: sub}, nil
}

// WatchChangeEpochEvent is a free log subscription operation binding the contract event 0x353f5ab7dd2715c6a221c2e7ff95ca17e68d962ffd1bb9b97b8f283d36598963.
//
// Solidity: event ChangeEpochEvent(uint256 height, bytes rawHeader, address[] oldValidators, address[] newValidators)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationFilterer) WatchChangeEpochEvent(opts *bind.WatchOpts, sink chan<- *EthCrossChainManagerImplementationChangeEpochEvent) (event.Subscription, error) {

	logs, sub, err := _EthCrossChainManagerImplementation.contract.WatchLogs(opts, "ChangeEpochEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainManagerImplementationChangeEpochEvent)
				if err := _EthCrossChainManagerImplementation.contract.UnpackLog(event, "ChangeEpochEvent", log); err != nil {
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

// ParseChangeEpochEvent is a log parse operation binding the contract event 0x353f5ab7dd2715c6a221c2e7ff95ca17e68d962ffd1bb9b97b8f283d36598963.
//
// Solidity: event ChangeEpochEvent(uint256 height, bytes rawHeader, address[] oldValidators, address[] newValidators)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationFilterer) ParseChangeEpochEvent(log types.Log) (*EthCrossChainManagerImplementationChangeEpochEvent, error) {
	event := new(EthCrossChainManagerImplementationChangeEpochEvent)
	if err := _EthCrossChainManagerImplementation.contract.UnpackLog(event, "ChangeEpochEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthCrossChainManagerImplementationCrossChainEventIterator is returned from FilterCrossChainEvent and is used to iterate over the raw logs and unpacked data for CrossChainEvent events raised by the EthCrossChainManagerImplementation contract.
type EthCrossChainManagerImplementationCrossChainEventIterator struct {
	Event *EthCrossChainManagerImplementationCrossChainEvent // Event containing the contract specifics and raw log

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
func (it *EthCrossChainManagerImplementationCrossChainEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainManagerImplementationCrossChainEvent)
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
		it.Event = new(EthCrossChainManagerImplementationCrossChainEvent)
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
func (it *EthCrossChainManagerImplementationCrossChainEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainManagerImplementationCrossChainEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainManagerImplementationCrossChainEvent represents a CrossChainEvent event raised by the EthCrossChainManagerImplementation contract.
type EthCrossChainManagerImplementationCrossChainEvent struct {
	Sender               common.Address
	TxId                 []byte
	ProxyOrAssetContract common.Address
	ToChainId            uint64
	ToContract           []byte
	Rawdata              []byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterCrossChainEvent is a free log retrieval operation binding the contract event 0x6ad3bf15c1988bc04bc153490cab16db8efb9a3990215bf1c64ea6e28be88483.
//
// Solidity: event CrossChainEvent(address indexed sender, bytes txId, address proxyOrAssetContract, uint64 toChainId, bytes toContract, bytes rawdata)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationFilterer) FilterCrossChainEvent(opts *bind.FilterOpts, sender []common.Address) (*EthCrossChainManagerImplementationCrossChainEventIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EthCrossChainManagerImplementation.contract.FilterLogs(opts, "CrossChainEvent", senderRule)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerImplementationCrossChainEventIterator{contract: _EthCrossChainManagerImplementation.contract, event: "CrossChainEvent", logs: logs, sub: sub}, nil
}

// WatchCrossChainEvent is a free log subscription operation binding the contract event 0x6ad3bf15c1988bc04bc153490cab16db8efb9a3990215bf1c64ea6e28be88483.
//
// Solidity: event CrossChainEvent(address indexed sender, bytes txId, address proxyOrAssetContract, uint64 toChainId, bytes toContract, bytes rawdata)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationFilterer) WatchCrossChainEvent(opts *bind.WatchOpts, sink chan<- *EthCrossChainManagerImplementationCrossChainEvent, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EthCrossChainManagerImplementation.contract.WatchLogs(opts, "CrossChainEvent", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainManagerImplementationCrossChainEvent)
				if err := _EthCrossChainManagerImplementation.contract.UnpackLog(event, "CrossChainEvent", log); err != nil {
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

// ParseCrossChainEvent is a log parse operation binding the contract event 0x6ad3bf15c1988bc04bc153490cab16db8efb9a3990215bf1c64ea6e28be88483.
//
// Solidity: event CrossChainEvent(address indexed sender, bytes txId, address proxyOrAssetContract, uint64 toChainId, bytes toContract, bytes rawdata)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationFilterer) ParseCrossChainEvent(log types.Log) (*EthCrossChainManagerImplementationCrossChainEvent, error) {
	event := new(EthCrossChainManagerImplementationCrossChainEvent)
	if err := _EthCrossChainManagerImplementation.contract.UnpackLog(event, "CrossChainEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthCrossChainManagerImplementationInitGenesisBlockEventIterator is returned from FilterInitGenesisBlockEvent and is used to iterate over the raw logs and unpacked data for InitGenesisBlockEvent events raised by the EthCrossChainManagerImplementation contract.
type EthCrossChainManagerImplementationInitGenesisBlockEventIterator struct {
	Event *EthCrossChainManagerImplementationInitGenesisBlockEvent // Event containing the contract specifics and raw log

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
func (it *EthCrossChainManagerImplementationInitGenesisBlockEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainManagerImplementationInitGenesisBlockEvent)
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
		it.Event = new(EthCrossChainManagerImplementationInitGenesisBlockEvent)
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
func (it *EthCrossChainManagerImplementationInitGenesisBlockEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainManagerImplementationInitGenesisBlockEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainManagerImplementationInitGenesisBlockEvent represents a InitGenesisBlockEvent event raised by the EthCrossChainManagerImplementation contract.
type EthCrossChainManagerImplementationInitGenesisBlockEvent struct {
	Height    *big.Int
	RawHeader []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterInitGenesisBlockEvent is a free log retrieval operation binding the contract event 0xf01968fc3a2655cf1b5144cb32de6dc898f91b9239c103744e8457152ab2fbde.
//
// Solidity: event InitGenesisBlockEvent(uint256 height, bytes rawHeader)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationFilterer) FilterInitGenesisBlockEvent(opts *bind.FilterOpts) (*EthCrossChainManagerImplementationInitGenesisBlockEventIterator, error) {

	logs, sub, err := _EthCrossChainManagerImplementation.contract.FilterLogs(opts, "InitGenesisBlockEvent")
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerImplementationInitGenesisBlockEventIterator{contract: _EthCrossChainManagerImplementation.contract, event: "InitGenesisBlockEvent", logs: logs, sub: sub}, nil
}

// WatchInitGenesisBlockEvent is a free log subscription operation binding the contract event 0xf01968fc3a2655cf1b5144cb32de6dc898f91b9239c103744e8457152ab2fbde.
//
// Solidity: event InitGenesisBlockEvent(uint256 height, bytes rawHeader)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationFilterer) WatchInitGenesisBlockEvent(opts *bind.WatchOpts, sink chan<- *EthCrossChainManagerImplementationInitGenesisBlockEvent) (event.Subscription, error) {

	logs, sub, err := _EthCrossChainManagerImplementation.contract.WatchLogs(opts, "InitGenesisBlockEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainManagerImplementationInitGenesisBlockEvent)
				if err := _EthCrossChainManagerImplementation.contract.UnpackLog(event, "InitGenesisBlockEvent", log); err != nil {
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

// ParseInitGenesisBlockEvent is a log parse operation binding the contract event 0xf01968fc3a2655cf1b5144cb32de6dc898f91b9239c103744e8457152ab2fbde.
//
// Solidity: event InitGenesisBlockEvent(uint256 height, bytes rawHeader)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationFilterer) ParseInitGenesisBlockEvent(log types.Log) (*EthCrossChainManagerImplementationInitGenesisBlockEvent, error) {
	event := new(EthCrossChainManagerImplementationInitGenesisBlockEvent)
	if err := _EthCrossChainManagerImplementation.contract.UnpackLog(event, "InitGenesisBlockEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthCrossChainManagerImplementationVerifyHeaderAndExecuteTxEventIterator is returned from FilterVerifyHeaderAndExecuteTxEvent and is used to iterate over the raw logs and unpacked data for VerifyHeaderAndExecuteTxEvent events raised by the EthCrossChainManagerImplementation contract.
type EthCrossChainManagerImplementationVerifyHeaderAndExecuteTxEventIterator struct {
	Event *EthCrossChainManagerImplementationVerifyHeaderAndExecuteTxEvent // Event containing the contract specifics and raw log

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
func (it *EthCrossChainManagerImplementationVerifyHeaderAndExecuteTxEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainManagerImplementationVerifyHeaderAndExecuteTxEvent)
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
		it.Event = new(EthCrossChainManagerImplementationVerifyHeaderAndExecuteTxEvent)
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
func (it *EthCrossChainManagerImplementationVerifyHeaderAndExecuteTxEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainManagerImplementationVerifyHeaderAndExecuteTxEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainManagerImplementationVerifyHeaderAndExecuteTxEvent represents a VerifyHeaderAndExecuteTxEvent event raised by the EthCrossChainManagerImplementation contract.
type EthCrossChainManagerImplementationVerifyHeaderAndExecuteTxEvent struct {
	FromChainID      uint64
	ToContract       []byte
	CrossChainTxHash []byte
	FromChainTxHash  []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterVerifyHeaderAndExecuteTxEvent is a free log retrieval operation binding the contract event 0x8a4a2663ce60ce4955c595da2894de0415240f1ace024cfbff85f513b656bdae.
//
// Solidity: event VerifyHeaderAndExecuteTxEvent(uint64 fromChainID, bytes toContract, bytes crossChainTxHash, bytes fromChainTxHash)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationFilterer) FilterVerifyHeaderAndExecuteTxEvent(opts *bind.FilterOpts) (*EthCrossChainManagerImplementationVerifyHeaderAndExecuteTxEventIterator, error) {

	logs, sub, err := _EthCrossChainManagerImplementation.contract.FilterLogs(opts, "VerifyHeaderAndExecuteTxEvent")
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerImplementationVerifyHeaderAndExecuteTxEventIterator{contract: _EthCrossChainManagerImplementation.contract, event: "VerifyHeaderAndExecuteTxEvent", logs: logs, sub: sub}, nil
}

// WatchVerifyHeaderAndExecuteTxEvent is a free log subscription operation binding the contract event 0x8a4a2663ce60ce4955c595da2894de0415240f1ace024cfbff85f513b656bdae.
//
// Solidity: event VerifyHeaderAndExecuteTxEvent(uint64 fromChainID, bytes toContract, bytes crossChainTxHash, bytes fromChainTxHash)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationFilterer) WatchVerifyHeaderAndExecuteTxEvent(opts *bind.WatchOpts, sink chan<- *EthCrossChainManagerImplementationVerifyHeaderAndExecuteTxEvent) (event.Subscription, error) {

	logs, sub, err := _EthCrossChainManagerImplementation.contract.WatchLogs(opts, "VerifyHeaderAndExecuteTxEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainManagerImplementationVerifyHeaderAndExecuteTxEvent)
				if err := _EthCrossChainManagerImplementation.contract.UnpackLog(event, "VerifyHeaderAndExecuteTxEvent", log); err != nil {
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

// ParseVerifyHeaderAndExecuteTxEvent is a log parse operation binding the contract event 0x8a4a2663ce60ce4955c595da2894de0415240f1ace024cfbff85f513b656bdae.
//
// Solidity: event VerifyHeaderAndExecuteTxEvent(uint64 fromChainID, bytes toContract, bytes crossChainTxHash, bytes fromChainTxHash)
func (_EthCrossChainManagerImplementation *EthCrossChainManagerImplementationFilterer) ParseVerifyHeaderAndExecuteTxEvent(log types.Log) (*EthCrossChainManagerImplementationVerifyHeaderAndExecuteTxEvent, error) {
	event := new(EthCrossChainManagerImplementationVerifyHeaderAndExecuteTxEvent)
	if err := _EthCrossChainManagerImplementation.contract.UnpackLog(event, "VerifyHeaderAndExecuteTxEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEthCrossChainDataABI is the input ABI used to generate the binding from.
const IEthCrossChainDataABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"fromChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"fromChainTx\",\"type\":\"bytes32\"}],\"name\":\"checkIfFromChainTxExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurEpochEndHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurEpochId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurEpochStartHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurEpochValidatorPkBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ethTxHashIndex\",\"type\":\"uint256\"}],\"name\":\"getEthTxHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEthTxHashIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"key2\",\"type\":\"bytes32\"}],\"name\":\"getExtraData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"fromChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"fromChainTx\",\"type\":\"bytes32\"}],\"name\":\"markFromChainTxExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"startHeight\",\"type\":\"uint64\"}],\"name\":\"putCurEpochEndHeight\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"epochId\",\"type\":\"uint64\"}],\"name\":\"putCurEpochId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"startHeight\",\"type\":\"uint64\"}],\"name\":\"putCurEpochStartHeight\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"curEpochPkBytes\",\"type\":\"bytes\"}],\"name\":\"putCurEpochValidatorPkBytes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ethTxHash\",\"type\":\"bytes32\"}],\"name\":\"putEthTxHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"key2\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"putExtraData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IEthCrossChainDataFuncSigs maps the 4-byte function signature to its string representation.
var IEthCrossChainDataFuncSigs = map[string]string{
	"0586763c": "checkIfFromChainTxExist(uint64,bytes32)",
	"4e4af8ce": "getCurEpochEndHeight()",
	"f881afd2": "getCurEpochId()",
	"5ac40790": "getCurEpochStartHeight()",
	"cd62908f": "getCurEpochValidatorPkBytes()",
	"29927875": "getEthTxHash(uint256)",
	"ff3d24a1": "getEthTxHashIndex()",
	"40602bb5": "getExtraData(bytes32,bytes32)",
	"e90bfdcf": "markFromChainTxExist(uint64,bytes32)",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"6a773962": "putCurEpochEndHeight(uint64)",
	"738b5ae4": "putCurEpochId(uint64)",
	"512feecc": "putCurEpochStartHeight(uint64)",
	"1dc544bf": "putCurEpochValidatorPkBytes(bytes)",
	"4c3ccf64": "putEthTxHash(bytes32)",
	"1afe374e": "putExtraData(bytes32,bytes32,bytes)",
	"f2fde38b": "transferOwnership(address)",
	"3f4ba83a": "unpause()",
}

// IEthCrossChainData is an auto generated Go binding around an Ethereum contract.
type IEthCrossChainData struct {
	IEthCrossChainDataCaller     // Read-only binding to the contract
	IEthCrossChainDataTransactor // Write-only binding to the contract
	IEthCrossChainDataFilterer   // Log filterer for contract events
}

// IEthCrossChainDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEthCrossChainDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEthCrossChainDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainDataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEthCrossChainDataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEthCrossChainDataSession struct {
	Contract     *IEthCrossChainData // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IEthCrossChainDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEthCrossChainDataCallerSession struct {
	Contract *IEthCrossChainDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IEthCrossChainDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEthCrossChainDataTransactorSession struct {
	Contract     *IEthCrossChainDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IEthCrossChainDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEthCrossChainDataRaw struct {
	Contract *IEthCrossChainData // Generic contract binding to access the raw methods on
}

// IEthCrossChainDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEthCrossChainDataCallerRaw struct {
	Contract *IEthCrossChainDataCaller // Generic read-only contract binding to access the raw methods on
}

// IEthCrossChainDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEthCrossChainDataTransactorRaw struct {
	Contract *IEthCrossChainDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEthCrossChainData creates a new instance of IEthCrossChainData, bound to a specific deployed contract.
func NewIEthCrossChainData(address common.Address, backend bind.ContractBackend) (*IEthCrossChainData, error) {
	contract, err := bindIEthCrossChainData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainData{IEthCrossChainDataCaller: IEthCrossChainDataCaller{contract: contract}, IEthCrossChainDataTransactor: IEthCrossChainDataTransactor{contract: contract}, IEthCrossChainDataFilterer: IEthCrossChainDataFilterer{contract: contract}}, nil
}

// NewIEthCrossChainDataCaller creates a new read-only instance of IEthCrossChainData, bound to a specific deployed contract.
func NewIEthCrossChainDataCaller(address common.Address, caller bind.ContractCaller) (*IEthCrossChainDataCaller, error) {
	contract, err := bindIEthCrossChainData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainDataCaller{contract: contract}, nil
}

// NewIEthCrossChainDataTransactor creates a new write-only instance of IEthCrossChainData, bound to a specific deployed contract.
func NewIEthCrossChainDataTransactor(address common.Address, transactor bind.ContractTransactor) (*IEthCrossChainDataTransactor, error) {
	contract, err := bindIEthCrossChainData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainDataTransactor{contract: contract}, nil
}

// NewIEthCrossChainDataFilterer creates a new log filterer instance of IEthCrossChainData, bound to a specific deployed contract.
func NewIEthCrossChainDataFilterer(address common.Address, filterer bind.ContractFilterer) (*IEthCrossChainDataFilterer, error) {
	contract, err := bindIEthCrossChainData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainDataFilterer{contract: contract}, nil
}

// bindIEthCrossChainData binds a generic wrapper to an already deployed contract.
func bindIEthCrossChainData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IEthCrossChainDataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthCrossChainData *IEthCrossChainDataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEthCrossChainData.Contract.IEthCrossChainDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthCrossChainData *IEthCrossChainDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.IEthCrossChainDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthCrossChainData *IEthCrossChainDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.IEthCrossChainDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthCrossChainData *IEthCrossChainDataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEthCrossChainData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthCrossChainData *IEthCrossChainDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthCrossChainData *IEthCrossChainDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.contract.Transact(opts, method, params...)
}

// CheckIfFromChainTxExist is a free data retrieval call binding the contract method 0x0586763c.
//
// Solidity: function checkIfFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) view returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataCaller) CheckIfFromChainTxExist(opts *bind.CallOpts, fromChainId uint64, fromChainTx [32]byte) (bool, error) {
	var out []interface{}
	err := _IEthCrossChainData.contract.Call(opts, &out, "checkIfFromChainTxExist", fromChainId, fromChainTx)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckIfFromChainTxExist is a free data retrieval call binding the contract method 0x0586763c.
//
// Solidity: function checkIfFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) view returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) CheckIfFromChainTxExist(fromChainId uint64, fromChainTx [32]byte) (bool, error) {
	return _IEthCrossChainData.Contract.CheckIfFromChainTxExist(&_IEthCrossChainData.CallOpts, fromChainId, fromChainTx)
}

// CheckIfFromChainTxExist is a free data retrieval call binding the contract method 0x0586763c.
//
// Solidity: function checkIfFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) view returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) CheckIfFromChainTxExist(fromChainId uint64, fromChainTx [32]byte) (bool, error) {
	return _IEthCrossChainData.Contract.CheckIfFromChainTxExist(&_IEthCrossChainData.CallOpts, fromChainId, fromChainTx)
}

// GetCurEpochEndHeight is a free data retrieval call binding the contract method 0x4e4af8ce.
//
// Solidity: function getCurEpochEndHeight() view returns(uint64)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetCurEpochEndHeight(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IEthCrossChainData.contract.Call(opts, &out, "getCurEpochEndHeight")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetCurEpochEndHeight is a free data retrieval call binding the contract method 0x4e4af8ce.
//
// Solidity: function getCurEpochEndHeight() view returns(uint64)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetCurEpochEndHeight() (uint64, error) {
	return _IEthCrossChainData.Contract.GetCurEpochEndHeight(&_IEthCrossChainData.CallOpts)
}

// GetCurEpochEndHeight is a free data retrieval call binding the contract method 0x4e4af8ce.
//
// Solidity: function getCurEpochEndHeight() view returns(uint64)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetCurEpochEndHeight() (uint64, error) {
	return _IEthCrossChainData.Contract.GetCurEpochEndHeight(&_IEthCrossChainData.CallOpts)
}

// GetCurEpochId is a free data retrieval call binding the contract method 0xf881afd2.
//
// Solidity: function getCurEpochId() view returns(uint64)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetCurEpochId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IEthCrossChainData.contract.Call(opts, &out, "getCurEpochId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetCurEpochId is a free data retrieval call binding the contract method 0xf881afd2.
//
// Solidity: function getCurEpochId() view returns(uint64)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetCurEpochId() (uint64, error) {
	return _IEthCrossChainData.Contract.GetCurEpochId(&_IEthCrossChainData.CallOpts)
}

// GetCurEpochId is a free data retrieval call binding the contract method 0xf881afd2.
//
// Solidity: function getCurEpochId() view returns(uint64)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetCurEpochId() (uint64, error) {
	return _IEthCrossChainData.Contract.GetCurEpochId(&_IEthCrossChainData.CallOpts)
}

// GetCurEpochStartHeight is a free data retrieval call binding the contract method 0x5ac40790.
//
// Solidity: function getCurEpochStartHeight() view returns(uint64)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetCurEpochStartHeight(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IEthCrossChainData.contract.Call(opts, &out, "getCurEpochStartHeight")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetCurEpochStartHeight is a free data retrieval call binding the contract method 0x5ac40790.
//
// Solidity: function getCurEpochStartHeight() view returns(uint64)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetCurEpochStartHeight() (uint64, error) {
	return _IEthCrossChainData.Contract.GetCurEpochStartHeight(&_IEthCrossChainData.CallOpts)
}

// GetCurEpochStartHeight is a free data retrieval call binding the contract method 0x5ac40790.
//
// Solidity: function getCurEpochStartHeight() view returns(uint64)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetCurEpochStartHeight() (uint64, error) {
	return _IEthCrossChainData.Contract.GetCurEpochStartHeight(&_IEthCrossChainData.CallOpts)
}

// GetCurEpochValidatorPkBytes is a free data retrieval call binding the contract method 0xcd62908f.
//
// Solidity: function getCurEpochValidatorPkBytes() view returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetCurEpochValidatorPkBytes(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _IEthCrossChainData.contract.Call(opts, &out, "getCurEpochValidatorPkBytes")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetCurEpochValidatorPkBytes is a free data retrieval call binding the contract method 0xcd62908f.
//
// Solidity: function getCurEpochValidatorPkBytes() view returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetCurEpochValidatorPkBytes() ([]byte, error) {
	return _IEthCrossChainData.Contract.GetCurEpochValidatorPkBytes(&_IEthCrossChainData.CallOpts)
}

// GetCurEpochValidatorPkBytes is a free data retrieval call binding the contract method 0xcd62908f.
//
// Solidity: function getCurEpochValidatorPkBytes() view returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetCurEpochValidatorPkBytes() ([]byte, error) {
	return _IEthCrossChainData.Contract.GetCurEpochValidatorPkBytes(&_IEthCrossChainData.CallOpts)
}

// GetEthTxHash is a free data retrieval call binding the contract method 0x29927875.
//
// Solidity: function getEthTxHash(uint256 ethTxHashIndex) view returns(bytes32)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetEthTxHash(opts *bind.CallOpts, ethTxHashIndex *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IEthCrossChainData.contract.Call(opts, &out, "getEthTxHash", ethTxHashIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEthTxHash is a free data retrieval call binding the contract method 0x29927875.
//
// Solidity: function getEthTxHash(uint256 ethTxHashIndex) view returns(bytes32)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetEthTxHash(ethTxHashIndex *big.Int) ([32]byte, error) {
	return _IEthCrossChainData.Contract.GetEthTxHash(&_IEthCrossChainData.CallOpts, ethTxHashIndex)
}

// GetEthTxHash is a free data retrieval call binding the contract method 0x29927875.
//
// Solidity: function getEthTxHash(uint256 ethTxHashIndex) view returns(bytes32)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetEthTxHash(ethTxHashIndex *big.Int) ([32]byte, error) {
	return _IEthCrossChainData.Contract.GetEthTxHash(&_IEthCrossChainData.CallOpts, ethTxHashIndex)
}

// GetEthTxHashIndex is a free data retrieval call binding the contract method 0xff3d24a1.
//
// Solidity: function getEthTxHashIndex() view returns(uint256)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetEthTxHashIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IEthCrossChainData.contract.Call(opts, &out, "getEthTxHashIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthTxHashIndex is a free data retrieval call binding the contract method 0xff3d24a1.
//
// Solidity: function getEthTxHashIndex() view returns(uint256)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetEthTxHashIndex() (*big.Int, error) {
	return _IEthCrossChainData.Contract.GetEthTxHashIndex(&_IEthCrossChainData.CallOpts)
}

// GetEthTxHashIndex is a free data retrieval call binding the contract method 0xff3d24a1.
//
// Solidity: function getEthTxHashIndex() view returns(uint256)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetEthTxHashIndex() (*big.Int, error) {
	return _IEthCrossChainData.Contract.GetEthTxHashIndex(&_IEthCrossChainData.CallOpts)
}

// GetExtraData is a free data retrieval call binding the contract method 0x40602bb5.
//
// Solidity: function getExtraData(bytes32 key1, bytes32 key2) view returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetExtraData(opts *bind.CallOpts, key1 [32]byte, key2 [32]byte) ([]byte, error) {
	var out []interface{}
	err := _IEthCrossChainData.contract.Call(opts, &out, "getExtraData", key1, key2)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetExtraData is a free data retrieval call binding the contract method 0x40602bb5.
//
// Solidity: function getExtraData(bytes32 key1, bytes32 key2) view returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetExtraData(key1 [32]byte, key2 [32]byte) ([]byte, error) {
	return _IEthCrossChainData.Contract.GetExtraData(&_IEthCrossChainData.CallOpts, key1, key2)
}

// GetExtraData is a free data retrieval call binding the contract method 0x40602bb5.
//
// Solidity: function getExtraData(bytes32 key1, bytes32 key2) view returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetExtraData(key1 [32]byte, key2 [32]byte) ([]byte, error) {
	return _IEthCrossChainData.Contract.GetExtraData(&_IEthCrossChainData.CallOpts, key1, key2)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IEthCrossChainData.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) Paused() (bool, error) {
	return _IEthCrossChainData.Contract.Paused(&_IEthCrossChainData.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) Paused() (bool, error) {
	return _IEthCrossChainData.Contract.Paused(&_IEthCrossChainData.CallOpts)
}

// MarkFromChainTxExist is a paid mutator transaction binding the contract method 0xe90bfdcf.
//
// Solidity: function markFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) MarkFromChainTxExist(opts *bind.TransactOpts, fromChainId uint64, fromChainTx [32]byte) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "markFromChainTxExist", fromChainId, fromChainTx)
}

// MarkFromChainTxExist is a paid mutator transaction binding the contract method 0xe90bfdcf.
//
// Solidity: function markFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) MarkFromChainTxExist(fromChainId uint64, fromChainTx [32]byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.MarkFromChainTxExist(&_IEthCrossChainData.TransactOpts, fromChainId, fromChainTx)
}

// MarkFromChainTxExist is a paid mutator transaction binding the contract method 0xe90bfdcf.
//
// Solidity: function markFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) MarkFromChainTxExist(fromChainId uint64, fromChainTx [32]byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.MarkFromChainTxExist(&_IEthCrossChainData.TransactOpts, fromChainId, fromChainTx)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) Pause() (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.Pause(&_IEthCrossChainData.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) Pause() (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.Pause(&_IEthCrossChainData.TransactOpts)
}

// PutCurEpochEndHeight is a paid mutator transaction binding the contract method 0x6a773962.
//
// Solidity: function putCurEpochEndHeight(uint64 startHeight) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutCurEpochEndHeight(opts *bind.TransactOpts, startHeight uint64) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putCurEpochEndHeight", startHeight)
}

// PutCurEpochEndHeight is a paid mutator transaction binding the contract method 0x6a773962.
//
// Solidity: function putCurEpochEndHeight(uint64 startHeight) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutCurEpochEndHeight(startHeight uint64) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutCurEpochEndHeight(&_IEthCrossChainData.TransactOpts, startHeight)
}

// PutCurEpochEndHeight is a paid mutator transaction binding the contract method 0x6a773962.
//
// Solidity: function putCurEpochEndHeight(uint64 startHeight) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutCurEpochEndHeight(startHeight uint64) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutCurEpochEndHeight(&_IEthCrossChainData.TransactOpts, startHeight)
}

// PutCurEpochId is a paid mutator transaction binding the contract method 0x738b5ae4.
//
// Solidity: function putCurEpochId(uint64 epochId) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutCurEpochId(opts *bind.TransactOpts, epochId uint64) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putCurEpochId", epochId)
}

// PutCurEpochId is a paid mutator transaction binding the contract method 0x738b5ae4.
//
// Solidity: function putCurEpochId(uint64 epochId) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutCurEpochId(epochId uint64) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutCurEpochId(&_IEthCrossChainData.TransactOpts, epochId)
}

// PutCurEpochId is a paid mutator transaction binding the contract method 0x738b5ae4.
//
// Solidity: function putCurEpochId(uint64 epochId) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutCurEpochId(epochId uint64) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutCurEpochId(&_IEthCrossChainData.TransactOpts, epochId)
}

// PutCurEpochStartHeight is a paid mutator transaction binding the contract method 0x512feecc.
//
// Solidity: function putCurEpochStartHeight(uint64 startHeight) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutCurEpochStartHeight(opts *bind.TransactOpts, startHeight uint64) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putCurEpochStartHeight", startHeight)
}

// PutCurEpochStartHeight is a paid mutator transaction binding the contract method 0x512feecc.
//
// Solidity: function putCurEpochStartHeight(uint64 startHeight) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutCurEpochStartHeight(startHeight uint64) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutCurEpochStartHeight(&_IEthCrossChainData.TransactOpts, startHeight)
}

// PutCurEpochStartHeight is a paid mutator transaction binding the contract method 0x512feecc.
//
// Solidity: function putCurEpochStartHeight(uint64 startHeight) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutCurEpochStartHeight(startHeight uint64) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutCurEpochStartHeight(&_IEthCrossChainData.TransactOpts, startHeight)
}

// PutCurEpochValidatorPkBytes is a paid mutator transaction binding the contract method 0x1dc544bf.
//
// Solidity: function putCurEpochValidatorPkBytes(bytes curEpochPkBytes) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutCurEpochValidatorPkBytes(opts *bind.TransactOpts, curEpochPkBytes []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putCurEpochValidatorPkBytes", curEpochPkBytes)
}

// PutCurEpochValidatorPkBytes is a paid mutator transaction binding the contract method 0x1dc544bf.
//
// Solidity: function putCurEpochValidatorPkBytes(bytes curEpochPkBytes) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutCurEpochValidatorPkBytes(curEpochPkBytes []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutCurEpochValidatorPkBytes(&_IEthCrossChainData.TransactOpts, curEpochPkBytes)
}

// PutCurEpochValidatorPkBytes is a paid mutator transaction binding the contract method 0x1dc544bf.
//
// Solidity: function putCurEpochValidatorPkBytes(bytes curEpochPkBytes) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutCurEpochValidatorPkBytes(curEpochPkBytes []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutCurEpochValidatorPkBytes(&_IEthCrossChainData.TransactOpts, curEpochPkBytes)
}

// PutEthTxHash is a paid mutator transaction binding the contract method 0x4c3ccf64.
//
// Solidity: function putEthTxHash(bytes32 ethTxHash) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutEthTxHash(opts *bind.TransactOpts, ethTxHash [32]byte) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putEthTxHash", ethTxHash)
}

// PutEthTxHash is a paid mutator transaction binding the contract method 0x4c3ccf64.
//
// Solidity: function putEthTxHash(bytes32 ethTxHash) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutEthTxHash(ethTxHash [32]byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutEthTxHash(&_IEthCrossChainData.TransactOpts, ethTxHash)
}

// PutEthTxHash is a paid mutator transaction binding the contract method 0x4c3ccf64.
//
// Solidity: function putEthTxHash(bytes32 ethTxHash) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutEthTxHash(ethTxHash [32]byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutEthTxHash(&_IEthCrossChainData.TransactOpts, ethTxHash)
}

// PutExtraData is a paid mutator transaction binding the contract method 0x1afe374e.
//
// Solidity: function putExtraData(bytes32 key1, bytes32 key2, bytes value) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutExtraData(opts *bind.TransactOpts, key1 [32]byte, key2 [32]byte, value []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putExtraData", key1, key2, value)
}

// PutExtraData is a paid mutator transaction binding the contract method 0x1afe374e.
//
// Solidity: function putExtraData(bytes32 key1, bytes32 key2, bytes value) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutExtraData(key1 [32]byte, key2 [32]byte, value []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutExtraData(&_IEthCrossChainData.TransactOpts, key1, key2, value)
}

// PutExtraData is a paid mutator transaction binding the contract method 0x1afe374e.
//
// Solidity: function putExtraData(bytes32 key1, bytes32 key2, bytes value) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutExtraData(key1 [32]byte, key2 [32]byte, value []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutExtraData(&_IEthCrossChainData.TransactOpts, key1, key2, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IEthCrossChainData *IEthCrossChainDataTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IEthCrossChainData *IEthCrossChainDataSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.TransferOwnership(&_IEthCrossChainData.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.TransferOwnership(&_IEthCrossChainData.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) Unpause() (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.Unpause(&_IEthCrossChainData.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) Unpause() (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.Unpause(&_IEthCrossChainData.TransactOpts)
}

// InitializableAdminUpgradeabilityProxyABI is the input ABI used to generate the binding from.
const InitializableAdminUpgradeabilityProxyABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// InitializableAdminUpgradeabilityProxyFuncSigs maps the 4-byte function signature to its string representation.
var InitializableAdminUpgradeabilityProxyFuncSigs = map[string]string{
	"f851a440": "admin()",
	"8f283970": "changeAdmin(address)",
	"5c60da1b": "implementation()",
	"cf7a1d77": "initialize(address,address,bytes)",
	"d1f57894": "initialize(address,bytes)",
	"3659cfe6": "upgradeTo(address)",
	"4f1ef286": "upgradeToAndCall(address,bytes)",
}

// InitializableAdminUpgradeabilityProxyBin is the compiled bytecode used for deploying new contracts.
var InitializableAdminUpgradeabilityProxyBin = "0x608060405234801561001057600080fd5b5061097c806100206000396000f3fe6080604052600436106100705760003560e01c80638f2839701161004e5780638f2839701461015e578063cf7a1d7714610191578063d1f5789414610250578063f851a4401461030657610070565b80633659cfe61461007a5780634f1ef286146100ad5780635c60da1b1461012d575b61007861031b565b005b34801561008657600080fd5b506100786004803603602081101561009d57600080fd5b50356001600160a01b0316610335565b610078600480360360408110156100c357600080fd5b6001600160a01b0382351691908101906040810160208201356401000000008111156100ee57600080fd5b82018360208201111561010057600080fd5b8035906020019184600183028401116401000000008311171561012257600080fd5b50909250905061036f565b34801561013957600080fd5b5061014261041c565b604080516001600160a01b039092168252519081900360200190f35b34801561016a57600080fd5b506100786004803603602081101561018157600080fd5b50356001600160a01b0316610459565b610078600480360360608110156101a757600080fd5b6001600160a01b0382358116926020810135909116918101906060810160408201356401000000008111156101db57600080fd5b8201836020820111156101ed57600080fd5b8035906020019184600183028401116401000000008311171561020f57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610513945050505050565b6100786004803603604081101561026657600080fd5b6001600160a01b03823516919081019060408101602082013564010000000081111561029157600080fd5b8201836020820111156102a357600080fd5b803590602001918460018302840111640100000000831117156102c557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610599945050505050565b34801561031257600080fd5b506101426106d9565b610323610704565b61033361032e610764565b610789565b565b61033d6107ad565b6001600160a01b0316336001600160a01b031614156103645761035f816107d2565b61036c565b61036c61031b565b50565b6103776107ad565b6001600160a01b0316336001600160a01b0316141561040f57610399836107d2565b6000836001600160a01b031683836040518083838082843760405192019450600093509091505080830381855af49150503d80600081146103f6576040519150601f19603f3d011682016040523d82523d6000602084013e6103fb565b606091505b505090508061040957600080fd5b50610417565b61041761031b565b505050565b60006104266107ad565b6001600160a01b0316336001600160a01b0316141561044e57610447610764565b9050610456565b61045661031b565b90565b6104616107ad565b6001600160a01b0316336001600160a01b03161415610364576001600160a01b0381166104bf5760405162461bcd60e51b81526004018080602001828103825260368152602001806108d76036913960400191505060405180910390fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6104e86107ad565b604080516001600160a01b03928316815291841660208301528051918290030190a161035f81610812565b600061051d610764565b6001600160a01b03161461053057600080fd5b61053a8382610599565b604080517232b4b8189c9b1b97383937bc3c9730b236b4b760691b815290519081900360130190207fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61036000199091011461059057fe5b61041782610812565b60006105a3610764565b6001600160a01b0316146105b657600080fd5b604080517f656970313936372e70726f78792e696d706c656d656e746174696f6e000000008152905190819003601c0190207f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6000199091011461061657fe5b61061f82610836565b8051156106d5576000826001600160a01b0316826040518082805190602001908083835b602083106106625780518252601f199092019160209182019101610643565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d80600081146106c2576040519150601f19603f3d011682016040523d82523d6000602084013e6106c7565b606091505b505090508061041757600080fd5b5050565b60006106e36107ad565b6001600160a01b0316336001600160a01b0316141561044e576104476107ad565b61070c6107ad565b6001600160a01b0316336001600160a01b0316141561075c5760405162461bcd60e51b81526004018080602001828103825260328152602001806108a56032913960400191505060405180910390fd5b610333610333565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5490565b3660008037600080366000845af43d6000803e8080156107a8573d6000f35b3d6000fd5b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035490565b6107db81610836565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610355565b61083f8161089e565b61087a5760405162461bcd60e51b815260040180806020018281038252603b81526020018061090d603b913960400191505060405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc55565b3b15159056fe43616e6e6f742063616c6c2066616c6c6261636b2066756e6374696f6e2066726f6d207468652070726f78792061646d696e43616e6e6f74206368616e6765207468652061646d696e206f6620612070726f787920746f20746865207a65726f206164647265737343616e6e6f742073657420612070726f787920696d706c656d656e746174696f6e20746f2061206e6f6e2d636f6e74726163742061646472657373a265627a7a72315820ea63902a998581c26cce013e1736a0cabdc6bc56ff51f923fdb0a60bef5645e364736f6c63430005110032"

// DeployInitializableAdminUpgradeabilityProxy deploys a new Ethereum contract, binding an instance of InitializableAdminUpgradeabilityProxy to it.
func DeployInitializableAdminUpgradeabilityProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InitializableAdminUpgradeabilityProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(InitializableAdminUpgradeabilityProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(InitializableAdminUpgradeabilityProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InitializableAdminUpgradeabilityProxy{InitializableAdminUpgradeabilityProxyCaller: InitializableAdminUpgradeabilityProxyCaller{contract: contract}, InitializableAdminUpgradeabilityProxyTransactor: InitializableAdminUpgradeabilityProxyTransactor{contract: contract}, InitializableAdminUpgradeabilityProxyFilterer: InitializableAdminUpgradeabilityProxyFilterer{contract: contract}}, nil
}

// InitializableAdminUpgradeabilityProxy is an auto generated Go binding around an Ethereum contract.
type InitializableAdminUpgradeabilityProxy struct {
	InitializableAdminUpgradeabilityProxyCaller     // Read-only binding to the contract
	InitializableAdminUpgradeabilityProxyTransactor // Write-only binding to the contract
	InitializableAdminUpgradeabilityProxyFilterer   // Log filterer for contract events
}

// InitializableAdminUpgradeabilityProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type InitializableAdminUpgradeabilityProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableAdminUpgradeabilityProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InitializableAdminUpgradeabilityProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableAdminUpgradeabilityProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InitializableAdminUpgradeabilityProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableAdminUpgradeabilityProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InitializableAdminUpgradeabilityProxySession struct {
	Contract     *InitializableAdminUpgradeabilityProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                          // Call options to use throughout this session
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// InitializableAdminUpgradeabilityProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InitializableAdminUpgradeabilityProxyCallerSession struct {
	Contract *InitializableAdminUpgradeabilityProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                                // Call options to use throughout this session
}

// InitializableAdminUpgradeabilityProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InitializableAdminUpgradeabilityProxyTransactorSession struct {
	Contract     *InitializableAdminUpgradeabilityProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                                // Transaction auth options to use throughout this session
}

// InitializableAdminUpgradeabilityProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type InitializableAdminUpgradeabilityProxyRaw struct {
	Contract *InitializableAdminUpgradeabilityProxy // Generic contract binding to access the raw methods on
}

// InitializableAdminUpgradeabilityProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InitializableAdminUpgradeabilityProxyCallerRaw struct {
	Contract *InitializableAdminUpgradeabilityProxyCaller // Generic read-only contract binding to access the raw methods on
}

// InitializableAdminUpgradeabilityProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InitializableAdminUpgradeabilityProxyTransactorRaw struct {
	Contract *InitializableAdminUpgradeabilityProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInitializableAdminUpgradeabilityProxy creates a new instance of InitializableAdminUpgradeabilityProxy, bound to a specific deployed contract.
func NewInitializableAdminUpgradeabilityProxy(address common.Address, backend bind.ContractBackend) (*InitializableAdminUpgradeabilityProxy, error) {
	contract, err := bindInitializableAdminUpgradeabilityProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InitializableAdminUpgradeabilityProxy{InitializableAdminUpgradeabilityProxyCaller: InitializableAdminUpgradeabilityProxyCaller{contract: contract}, InitializableAdminUpgradeabilityProxyTransactor: InitializableAdminUpgradeabilityProxyTransactor{contract: contract}, InitializableAdminUpgradeabilityProxyFilterer: InitializableAdminUpgradeabilityProxyFilterer{contract: contract}}, nil
}

// NewInitializableAdminUpgradeabilityProxyCaller creates a new read-only instance of InitializableAdminUpgradeabilityProxy, bound to a specific deployed contract.
func NewInitializableAdminUpgradeabilityProxyCaller(address common.Address, caller bind.ContractCaller) (*InitializableAdminUpgradeabilityProxyCaller, error) {
	contract, err := bindInitializableAdminUpgradeabilityProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableAdminUpgradeabilityProxyCaller{contract: contract}, nil
}

// NewInitializableAdminUpgradeabilityProxyTransactor creates a new write-only instance of InitializableAdminUpgradeabilityProxy, bound to a specific deployed contract.
func NewInitializableAdminUpgradeabilityProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*InitializableAdminUpgradeabilityProxyTransactor, error) {
	contract, err := bindInitializableAdminUpgradeabilityProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableAdminUpgradeabilityProxyTransactor{contract: contract}, nil
}

// NewInitializableAdminUpgradeabilityProxyFilterer creates a new log filterer instance of InitializableAdminUpgradeabilityProxy, bound to a specific deployed contract.
func NewInitializableAdminUpgradeabilityProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*InitializableAdminUpgradeabilityProxyFilterer, error) {
	contract, err := bindInitializableAdminUpgradeabilityProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InitializableAdminUpgradeabilityProxyFilterer{contract: contract}, nil
}

// bindInitializableAdminUpgradeabilityProxy binds a generic wrapper to an already deployed contract.
func bindInitializableAdminUpgradeabilityProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InitializableAdminUpgradeabilityProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InitializableAdminUpgradeabilityProxy.Contract.InitializableAdminUpgradeabilityProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InitializableAdminUpgradeabilityProxy.Contract.InitializableAdminUpgradeabilityProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InitializableAdminUpgradeabilityProxy.Contract.InitializableAdminUpgradeabilityProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InitializableAdminUpgradeabilityProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InitializableAdminUpgradeabilityProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InitializableAdminUpgradeabilityProxy.Contract.contract.Transact(opts, method, params...)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxyTransactor) Admin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InitializableAdminUpgradeabilityProxy.contract.Transact(opts, "admin")
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxySession) Admin() (*types.Transaction, error) {
	return _InitializableAdminUpgradeabilityProxy.Contract.Admin(&_InitializableAdminUpgradeabilityProxy.TransactOpts)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxyTransactorSession) Admin() (*types.Transaction, error) {
	return _InitializableAdminUpgradeabilityProxy.Contract.Admin(&_InitializableAdminUpgradeabilityProxy.TransactOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxyTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _InitializableAdminUpgradeabilityProxy.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxySession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _InitializableAdminUpgradeabilityProxy.Contract.ChangeAdmin(&_InitializableAdminUpgradeabilityProxy.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxyTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _InitializableAdminUpgradeabilityProxy.Contract.ChangeAdmin(&_InitializableAdminUpgradeabilityProxy.TransactOpts, newAdmin)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxyTransactor) Implementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InitializableAdminUpgradeabilityProxy.contract.Transact(opts, "implementation")
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxySession) Implementation() (*types.Transaction, error) {
	return _InitializableAdminUpgradeabilityProxy.Contract.Implementation(&_InitializableAdminUpgradeabilityProxy.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxyTransactorSession) Implementation() (*types.Transaction, error) {
	return _InitializableAdminUpgradeabilityProxy.Contract.Implementation(&_InitializableAdminUpgradeabilityProxy.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf7a1d77.
//
// Solidity: function initialize(address _logic, address _admin, bytes _data) payable returns()
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxyTransactor) Initialize(opts *bind.TransactOpts, _logic common.Address, _admin common.Address, _data []byte) (*types.Transaction, error) {
	return _InitializableAdminUpgradeabilityProxy.contract.Transact(opts, "initialize", _logic, _admin, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf7a1d77.
//
// Solidity: function initialize(address _logic, address _admin, bytes _data) payable returns()
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxySession) Initialize(_logic common.Address, _admin common.Address, _data []byte) (*types.Transaction, error) {
	return _InitializableAdminUpgradeabilityProxy.Contract.Initialize(&_InitializableAdminUpgradeabilityProxy.TransactOpts, _logic, _admin, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf7a1d77.
//
// Solidity: function initialize(address _logic, address _admin, bytes _data) payable returns()
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxyTransactorSession) Initialize(_logic common.Address, _admin common.Address, _data []byte) (*types.Transaction, error) {
	return _InitializableAdminUpgradeabilityProxy.Contract.Initialize(&_InitializableAdminUpgradeabilityProxy.TransactOpts, _logic, _admin, _data)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xd1f57894.
//
// Solidity: function initialize(address _logic, bytes _data) payable returns()
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxyTransactor) Initialize0(opts *bind.TransactOpts, _logic common.Address, _data []byte) (*types.Transaction, error) {
	return _InitializableAdminUpgradeabilityProxy.contract.Transact(opts, "initialize0", _logic, _data)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xd1f57894.
//
// Solidity: function initialize(address _logic, bytes _data) payable returns()
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxySession) Initialize0(_logic common.Address, _data []byte) (*types.Transaction, error) {
	return _InitializableAdminUpgradeabilityProxy.Contract.Initialize0(&_InitializableAdminUpgradeabilityProxy.TransactOpts, _logic, _data)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xd1f57894.
//
// Solidity: function initialize(address _logic, bytes _data) payable returns()
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxyTransactorSession) Initialize0(_logic common.Address, _data []byte) (*types.Transaction, error) {
	return _InitializableAdminUpgradeabilityProxy.Contract.Initialize0(&_InitializableAdminUpgradeabilityProxy.TransactOpts, _logic, _data)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxyTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _InitializableAdminUpgradeabilityProxy.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _InitializableAdminUpgradeabilityProxy.Contract.UpgradeTo(&_InitializableAdminUpgradeabilityProxy.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxyTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _InitializableAdminUpgradeabilityProxy.Contract.UpgradeTo(&_InitializableAdminUpgradeabilityProxy.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxyTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _InitializableAdminUpgradeabilityProxy.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _InitializableAdminUpgradeabilityProxy.Contract.UpgradeToAndCall(&_InitializableAdminUpgradeabilityProxy.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxyTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _InitializableAdminUpgradeabilityProxy.Contract.UpgradeToAndCall(&_InitializableAdminUpgradeabilityProxy.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _InitializableAdminUpgradeabilityProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _InitializableAdminUpgradeabilityProxy.Contract.Fallback(&_InitializableAdminUpgradeabilityProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _InitializableAdminUpgradeabilityProxy.Contract.Fallback(&_InitializableAdminUpgradeabilityProxy.TransactOpts, calldata)
}

// InitializableAdminUpgradeabilityProxyAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the InitializableAdminUpgradeabilityProxy contract.
type InitializableAdminUpgradeabilityProxyAdminChangedIterator struct {
	Event *InitializableAdminUpgradeabilityProxyAdminChanged // Event containing the contract specifics and raw log

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
func (it *InitializableAdminUpgradeabilityProxyAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InitializableAdminUpgradeabilityProxyAdminChanged)
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
		it.Event = new(InitializableAdminUpgradeabilityProxyAdminChanged)
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
func (it *InitializableAdminUpgradeabilityProxyAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InitializableAdminUpgradeabilityProxyAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InitializableAdminUpgradeabilityProxyAdminChanged represents a AdminChanged event raised by the InitializableAdminUpgradeabilityProxy contract.
type InitializableAdminUpgradeabilityProxyAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxyFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*InitializableAdminUpgradeabilityProxyAdminChangedIterator, error) {

	logs, sub, err := _InitializableAdminUpgradeabilityProxy.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &InitializableAdminUpgradeabilityProxyAdminChangedIterator{contract: _InitializableAdminUpgradeabilityProxy.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxyFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *InitializableAdminUpgradeabilityProxyAdminChanged) (event.Subscription, error) {

	logs, sub, err := _InitializableAdminUpgradeabilityProxy.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InitializableAdminUpgradeabilityProxyAdminChanged)
				if err := _InitializableAdminUpgradeabilityProxy.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxyFilterer) ParseAdminChanged(log types.Log) (*InitializableAdminUpgradeabilityProxyAdminChanged, error) {
	event := new(InitializableAdminUpgradeabilityProxyAdminChanged)
	if err := _InitializableAdminUpgradeabilityProxy.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InitializableAdminUpgradeabilityProxyUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the InitializableAdminUpgradeabilityProxy contract.
type InitializableAdminUpgradeabilityProxyUpgradedIterator struct {
	Event *InitializableAdminUpgradeabilityProxyUpgraded // Event containing the contract specifics and raw log

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
func (it *InitializableAdminUpgradeabilityProxyUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InitializableAdminUpgradeabilityProxyUpgraded)
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
		it.Event = new(InitializableAdminUpgradeabilityProxyUpgraded)
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
func (it *InitializableAdminUpgradeabilityProxyUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InitializableAdminUpgradeabilityProxyUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InitializableAdminUpgradeabilityProxyUpgraded represents a Upgraded event raised by the InitializableAdminUpgradeabilityProxy contract.
type InitializableAdminUpgradeabilityProxyUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxyFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*InitializableAdminUpgradeabilityProxyUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _InitializableAdminUpgradeabilityProxy.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &InitializableAdminUpgradeabilityProxyUpgradedIterator{contract: _InitializableAdminUpgradeabilityProxy.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxyFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *InitializableAdminUpgradeabilityProxyUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _InitializableAdminUpgradeabilityProxy.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InitializableAdminUpgradeabilityProxyUpgraded)
				if err := _InitializableAdminUpgradeabilityProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_InitializableAdminUpgradeabilityProxy *InitializableAdminUpgradeabilityProxyFilterer) ParseUpgraded(log types.Log) (*InitializableAdminUpgradeabilityProxyUpgraded, error) {
	event := new(InitializableAdminUpgradeabilityProxyUpgraded)
	if err := _InitializableAdminUpgradeabilityProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InitializableUpgradeabilityProxyABI is the input ABI used to generate the binding from.
const InitializableUpgradeabilityProxyABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// InitializableUpgradeabilityProxyFuncSigs maps the 4-byte function signature to its string representation.
var InitializableUpgradeabilityProxyFuncSigs = map[string]string{
	"d1f57894": "initialize(address,bytes)",
}

// InitializableUpgradeabilityProxyBin is the compiled bytecode used for deploying new contracts.
var InitializableUpgradeabilityProxyBin = "0x608060405234801561001057600080fd5b50610361806100206000396000f3fe60806040526004361061001e5760003560e01c8063d1f5789414610028575b6100266100de565b005b6100266004803603604081101561003e57600080fd5b6001600160a01b03823516919081019060408101602082013564010000000081111561006957600080fd5b82018360208201111561007b57600080fd5b8035906020019184600183028401116401000000008311171561009d57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506100f8945050505050565b6100e66100f6565b6100f66100f161023a565b61025f565b565b600061010261023a565b6001600160a01b03161461011557600080fd5b604080517f656970313936372e70726f78792e696d706c656d656e746174696f6e000000008152905190819003601c0190207f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6000199091011461017557fe5b61017e82610283565b805115610236576000826001600160a01b0316826040518082805190602001908083835b602083106101c15780518252601f1990920191602091820191016101a2565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d8060008114610221576040519150601f19603f3d011682016040523d82523d6000602084013e610226565b606091505b505090508061023457600080fd5b505b5050565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5490565b3660008037600080366000845af43d6000803e80801561027e573d6000f35b3d6000fd5b61028c816102eb565b6102c75760405162461bcd60e51b815260040180806020018281038252603b8152602001806102f2603b913960400191505060405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc55565b3b15159056fe43616e6e6f742073657420612070726f787920696d706c656d656e746174696f6e20746f2061206e6f6e2d636f6e74726163742061646472657373a265627a7a723158207c3bf7b47974ba6c4555599044f9121b1cc815eab2feea0904bd24b94277a8b264736f6c63430005110032"

// DeployInitializableUpgradeabilityProxy deploys a new Ethereum contract, binding an instance of InitializableUpgradeabilityProxy to it.
func DeployInitializableUpgradeabilityProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InitializableUpgradeabilityProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(InitializableUpgradeabilityProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(InitializableUpgradeabilityProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InitializableUpgradeabilityProxy{InitializableUpgradeabilityProxyCaller: InitializableUpgradeabilityProxyCaller{contract: contract}, InitializableUpgradeabilityProxyTransactor: InitializableUpgradeabilityProxyTransactor{contract: contract}, InitializableUpgradeabilityProxyFilterer: InitializableUpgradeabilityProxyFilterer{contract: contract}}, nil
}

// InitializableUpgradeabilityProxy is an auto generated Go binding around an Ethereum contract.
type InitializableUpgradeabilityProxy struct {
	InitializableUpgradeabilityProxyCaller     // Read-only binding to the contract
	InitializableUpgradeabilityProxyTransactor // Write-only binding to the contract
	InitializableUpgradeabilityProxyFilterer   // Log filterer for contract events
}

// InitializableUpgradeabilityProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type InitializableUpgradeabilityProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableUpgradeabilityProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InitializableUpgradeabilityProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableUpgradeabilityProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InitializableUpgradeabilityProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableUpgradeabilityProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InitializableUpgradeabilityProxySession struct {
	Contract     *InitializableUpgradeabilityProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                     // Call options to use throughout this session
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// InitializableUpgradeabilityProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InitializableUpgradeabilityProxyCallerSession struct {
	Contract *InitializableUpgradeabilityProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                           // Call options to use throughout this session
}

// InitializableUpgradeabilityProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InitializableUpgradeabilityProxyTransactorSession struct {
	Contract     *InitializableUpgradeabilityProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                           // Transaction auth options to use throughout this session
}

// InitializableUpgradeabilityProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type InitializableUpgradeabilityProxyRaw struct {
	Contract *InitializableUpgradeabilityProxy // Generic contract binding to access the raw methods on
}

// InitializableUpgradeabilityProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InitializableUpgradeabilityProxyCallerRaw struct {
	Contract *InitializableUpgradeabilityProxyCaller // Generic read-only contract binding to access the raw methods on
}

// InitializableUpgradeabilityProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InitializableUpgradeabilityProxyTransactorRaw struct {
	Contract *InitializableUpgradeabilityProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInitializableUpgradeabilityProxy creates a new instance of InitializableUpgradeabilityProxy, bound to a specific deployed contract.
func NewInitializableUpgradeabilityProxy(address common.Address, backend bind.ContractBackend) (*InitializableUpgradeabilityProxy, error) {
	contract, err := bindInitializableUpgradeabilityProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InitializableUpgradeabilityProxy{InitializableUpgradeabilityProxyCaller: InitializableUpgradeabilityProxyCaller{contract: contract}, InitializableUpgradeabilityProxyTransactor: InitializableUpgradeabilityProxyTransactor{contract: contract}, InitializableUpgradeabilityProxyFilterer: InitializableUpgradeabilityProxyFilterer{contract: contract}}, nil
}

// NewInitializableUpgradeabilityProxyCaller creates a new read-only instance of InitializableUpgradeabilityProxy, bound to a specific deployed contract.
func NewInitializableUpgradeabilityProxyCaller(address common.Address, caller bind.ContractCaller) (*InitializableUpgradeabilityProxyCaller, error) {
	contract, err := bindInitializableUpgradeabilityProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableUpgradeabilityProxyCaller{contract: contract}, nil
}

// NewInitializableUpgradeabilityProxyTransactor creates a new write-only instance of InitializableUpgradeabilityProxy, bound to a specific deployed contract.
func NewInitializableUpgradeabilityProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*InitializableUpgradeabilityProxyTransactor, error) {
	contract, err := bindInitializableUpgradeabilityProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableUpgradeabilityProxyTransactor{contract: contract}, nil
}

// NewInitializableUpgradeabilityProxyFilterer creates a new log filterer instance of InitializableUpgradeabilityProxy, bound to a specific deployed contract.
func NewInitializableUpgradeabilityProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*InitializableUpgradeabilityProxyFilterer, error) {
	contract, err := bindInitializableUpgradeabilityProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InitializableUpgradeabilityProxyFilterer{contract: contract}, nil
}

// bindInitializableUpgradeabilityProxy binds a generic wrapper to an already deployed contract.
func bindInitializableUpgradeabilityProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InitializableUpgradeabilityProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InitializableUpgradeabilityProxy *InitializableUpgradeabilityProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InitializableUpgradeabilityProxy.Contract.InitializableUpgradeabilityProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InitializableUpgradeabilityProxy *InitializableUpgradeabilityProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InitializableUpgradeabilityProxy.Contract.InitializableUpgradeabilityProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InitializableUpgradeabilityProxy *InitializableUpgradeabilityProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InitializableUpgradeabilityProxy.Contract.InitializableUpgradeabilityProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InitializableUpgradeabilityProxy *InitializableUpgradeabilityProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InitializableUpgradeabilityProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InitializableUpgradeabilityProxy *InitializableUpgradeabilityProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InitializableUpgradeabilityProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InitializableUpgradeabilityProxy *InitializableUpgradeabilityProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InitializableUpgradeabilityProxy.Contract.contract.Transact(opts, method, params...)
}

// Initialize is a paid mutator transaction binding the contract method 0xd1f57894.
//
// Solidity: function initialize(address _logic, bytes _data) payable returns()
func (_InitializableUpgradeabilityProxy *InitializableUpgradeabilityProxyTransactor) Initialize(opts *bind.TransactOpts, _logic common.Address, _data []byte) (*types.Transaction, error) {
	return _InitializableUpgradeabilityProxy.contract.Transact(opts, "initialize", _logic, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xd1f57894.
//
// Solidity: function initialize(address _logic, bytes _data) payable returns()
func (_InitializableUpgradeabilityProxy *InitializableUpgradeabilityProxySession) Initialize(_logic common.Address, _data []byte) (*types.Transaction, error) {
	return _InitializableUpgradeabilityProxy.Contract.Initialize(&_InitializableUpgradeabilityProxy.TransactOpts, _logic, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xd1f57894.
//
// Solidity: function initialize(address _logic, bytes _data) payable returns()
func (_InitializableUpgradeabilityProxy *InitializableUpgradeabilityProxyTransactorSession) Initialize(_logic common.Address, _data []byte) (*types.Transaction, error) {
	return _InitializableUpgradeabilityProxy.Contract.Initialize(&_InitializableUpgradeabilityProxy.TransactOpts, _logic, _data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_InitializableUpgradeabilityProxy *InitializableUpgradeabilityProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _InitializableUpgradeabilityProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_InitializableUpgradeabilityProxy *InitializableUpgradeabilityProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _InitializableUpgradeabilityProxy.Contract.Fallback(&_InitializableUpgradeabilityProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_InitializableUpgradeabilityProxy *InitializableUpgradeabilityProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _InitializableUpgradeabilityProxy.Contract.Fallback(&_InitializableUpgradeabilityProxy.TransactOpts, calldata)
}

// InitializableUpgradeabilityProxyUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the InitializableUpgradeabilityProxy contract.
type InitializableUpgradeabilityProxyUpgradedIterator struct {
	Event *InitializableUpgradeabilityProxyUpgraded // Event containing the contract specifics and raw log

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
func (it *InitializableUpgradeabilityProxyUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InitializableUpgradeabilityProxyUpgraded)
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
		it.Event = new(InitializableUpgradeabilityProxyUpgraded)
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
func (it *InitializableUpgradeabilityProxyUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InitializableUpgradeabilityProxyUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InitializableUpgradeabilityProxyUpgraded represents a Upgraded event raised by the InitializableUpgradeabilityProxy contract.
type InitializableUpgradeabilityProxyUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_InitializableUpgradeabilityProxy *InitializableUpgradeabilityProxyFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*InitializableUpgradeabilityProxyUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _InitializableUpgradeabilityProxy.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &InitializableUpgradeabilityProxyUpgradedIterator{contract: _InitializableUpgradeabilityProxy.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_InitializableUpgradeabilityProxy *InitializableUpgradeabilityProxyFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *InitializableUpgradeabilityProxyUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _InitializableUpgradeabilityProxy.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InitializableUpgradeabilityProxyUpgraded)
				if err := _InitializableUpgradeabilityProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_InitializableUpgradeabilityProxy *InitializableUpgradeabilityProxyFilterer) ParseUpgraded(log types.Log) (*InitializableUpgradeabilityProxyUpgraded, error) {
	event := new(InitializableUpgradeabilityProxyUpgraded)
	if err := _InitializableUpgradeabilityProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpenZeppelinUpgradesAddressABI is the input ABI used to generate the binding from.
const OpenZeppelinUpgradesAddressABI = "[]"

// OpenZeppelinUpgradesAddressBin is the compiled bytecode used for deploying new contracts.
var OpenZeppelinUpgradesAddressBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158207d6fb58a6915b69901cc63a6915fc24061204c50a7a9c01998b9e98c2877779264736f6c63430005110032"

// DeployOpenZeppelinUpgradesAddress deploys a new Ethereum contract, binding an instance of OpenZeppelinUpgradesAddress to it.
func DeployOpenZeppelinUpgradesAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OpenZeppelinUpgradesAddress, error) {
	parsed, err := abi.JSON(strings.NewReader(OpenZeppelinUpgradesAddressABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OpenZeppelinUpgradesAddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OpenZeppelinUpgradesAddress{OpenZeppelinUpgradesAddressCaller: OpenZeppelinUpgradesAddressCaller{contract: contract}, OpenZeppelinUpgradesAddressTransactor: OpenZeppelinUpgradesAddressTransactor{contract: contract}, OpenZeppelinUpgradesAddressFilterer: OpenZeppelinUpgradesAddressFilterer{contract: contract}}, nil
}

// OpenZeppelinUpgradesAddress is an auto generated Go binding around an Ethereum contract.
type OpenZeppelinUpgradesAddress struct {
	OpenZeppelinUpgradesAddressCaller     // Read-only binding to the contract
	OpenZeppelinUpgradesAddressTransactor // Write-only binding to the contract
	OpenZeppelinUpgradesAddressFilterer   // Log filterer for contract events
}

// OpenZeppelinUpgradesAddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type OpenZeppelinUpgradesAddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenZeppelinUpgradesAddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OpenZeppelinUpgradesAddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenZeppelinUpgradesAddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OpenZeppelinUpgradesAddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenZeppelinUpgradesAddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OpenZeppelinUpgradesAddressSession struct {
	Contract     *OpenZeppelinUpgradesAddress // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// OpenZeppelinUpgradesAddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OpenZeppelinUpgradesAddressCallerSession struct {
	Contract *OpenZeppelinUpgradesAddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// OpenZeppelinUpgradesAddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OpenZeppelinUpgradesAddressTransactorSession struct {
	Contract     *OpenZeppelinUpgradesAddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// OpenZeppelinUpgradesAddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type OpenZeppelinUpgradesAddressRaw struct {
	Contract *OpenZeppelinUpgradesAddress // Generic contract binding to access the raw methods on
}

// OpenZeppelinUpgradesAddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OpenZeppelinUpgradesAddressCallerRaw struct {
	Contract *OpenZeppelinUpgradesAddressCaller // Generic read-only contract binding to access the raw methods on
}

// OpenZeppelinUpgradesAddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OpenZeppelinUpgradesAddressTransactorRaw struct {
	Contract *OpenZeppelinUpgradesAddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOpenZeppelinUpgradesAddress creates a new instance of OpenZeppelinUpgradesAddress, bound to a specific deployed contract.
func NewOpenZeppelinUpgradesAddress(address common.Address, backend bind.ContractBackend) (*OpenZeppelinUpgradesAddress, error) {
	contract, err := bindOpenZeppelinUpgradesAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OpenZeppelinUpgradesAddress{OpenZeppelinUpgradesAddressCaller: OpenZeppelinUpgradesAddressCaller{contract: contract}, OpenZeppelinUpgradesAddressTransactor: OpenZeppelinUpgradesAddressTransactor{contract: contract}, OpenZeppelinUpgradesAddressFilterer: OpenZeppelinUpgradesAddressFilterer{contract: contract}}, nil
}

// NewOpenZeppelinUpgradesAddressCaller creates a new read-only instance of OpenZeppelinUpgradesAddress, bound to a specific deployed contract.
func NewOpenZeppelinUpgradesAddressCaller(address common.Address, caller bind.ContractCaller) (*OpenZeppelinUpgradesAddressCaller, error) {
	contract, err := bindOpenZeppelinUpgradesAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OpenZeppelinUpgradesAddressCaller{contract: contract}, nil
}

// NewOpenZeppelinUpgradesAddressTransactor creates a new write-only instance of OpenZeppelinUpgradesAddress, bound to a specific deployed contract.
func NewOpenZeppelinUpgradesAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*OpenZeppelinUpgradesAddressTransactor, error) {
	contract, err := bindOpenZeppelinUpgradesAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OpenZeppelinUpgradesAddressTransactor{contract: contract}, nil
}

// NewOpenZeppelinUpgradesAddressFilterer creates a new log filterer instance of OpenZeppelinUpgradesAddress, bound to a specific deployed contract.
func NewOpenZeppelinUpgradesAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*OpenZeppelinUpgradesAddressFilterer, error) {
	contract, err := bindOpenZeppelinUpgradesAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OpenZeppelinUpgradesAddressFilterer{contract: contract}, nil
}

// bindOpenZeppelinUpgradesAddress binds a generic wrapper to an already deployed contract.
func bindOpenZeppelinUpgradesAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OpenZeppelinUpgradesAddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OpenZeppelinUpgradesAddress *OpenZeppelinUpgradesAddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OpenZeppelinUpgradesAddress.Contract.OpenZeppelinUpgradesAddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OpenZeppelinUpgradesAddress *OpenZeppelinUpgradesAddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpenZeppelinUpgradesAddress.Contract.OpenZeppelinUpgradesAddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OpenZeppelinUpgradesAddress *OpenZeppelinUpgradesAddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OpenZeppelinUpgradesAddress.Contract.OpenZeppelinUpgradesAddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OpenZeppelinUpgradesAddress *OpenZeppelinUpgradesAddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OpenZeppelinUpgradesAddress.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OpenZeppelinUpgradesAddress *OpenZeppelinUpgradesAddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpenZeppelinUpgradesAddress.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OpenZeppelinUpgradesAddress *OpenZeppelinUpgradesAddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OpenZeppelinUpgradesAddress.Contract.contract.Transact(opts, method, params...)
}

// OpenZeppelinUpgradesECDSAABI is the input ABI used to generate the binding from.
const OpenZeppelinUpgradesECDSAABI = "[]"

// OpenZeppelinUpgradesECDSABin is the compiled bytecode used for deploying new contracts.
var OpenZeppelinUpgradesECDSABin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582016550ff8664e145a7a9d076479aa06c243bd9a64126fcdcb189638ab2971c05764736f6c63430005110032"

// DeployOpenZeppelinUpgradesECDSA deploys a new Ethereum contract, binding an instance of OpenZeppelinUpgradesECDSA to it.
func DeployOpenZeppelinUpgradesECDSA(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OpenZeppelinUpgradesECDSA, error) {
	parsed, err := abi.JSON(strings.NewReader(OpenZeppelinUpgradesECDSAABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OpenZeppelinUpgradesECDSABin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OpenZeppelinUpgradesECDSA{OpenZeppelinUpgradesECDSACaller: OpenZeppelinUpgradesECDSACaller{contract: contract}, OpenZeppelinUpgradesECDSATransactor: OpenZeppelinUpgradesECDSATransactor{contract: contract}, OpenZeppelinUpgradesECDSAFilterer: OpenZeppelinUpgradesECDSAFilterer{contract: contract}}, nil
}

// OpenZeppelinUpgradesECDSA is an auto generated Go binding around an Ethereum contract.
type OpenZeppelinUpgradesECDSA struct {
	OpenZeppelinUpgradesECDSACaller     // Read-only binding to the contract
	OpenZeppelinUpgradesECDSATransactor // Write-only binding to the contract
	OpenZeppelinUpgradesECDSAFilterer   // Log filterer for contract events
}

// OpenZeppelinUpgradesECDSACaller is an auto generated read-only Go binding around an Ethereum contract.
type OpenZeppelinUpgradesECDSACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenZeppelinUpgradesECDSATransactor is an auto generated write-only Go binding around an Ethereum contract.
type OpenZeppelinUpgradesECDSATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenZeppelinUpgradesECDSAFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OpenZeppelinUpgradesECDSAFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenZeppelinUpgradesECDSASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OpenZeppelinUpgradesECDSASession struct {
	Contract     *OpenZeppelinUpgradesECDSA // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// OpenZeppelinUpgradesECDSACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OpenZeppelinUpgradesECDSACallerSession struct {
	Contract *OpenZeppelinUpgradesECDSACaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// OpenZeppelinUpgradesECDSATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OpenZeppelinUpgradesECDSATransactorSession struct {
	Contract     *OpenZeppelinUpgradesECDSATransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// OpenZeppelinUpgradesECDSARaw is an auto generated low-level Go binding around an Ethereum contract.
type OpenZeppelinUpgradesECDSARaw struct {
	Contract *OpenZeppelinUpgradesECDSA // Generic contract binding to access the raw methods on
}

// OpenZeppelinUpgradesECDSACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OpenZeppelinUpgradesECDSACallerRaw struct {
	Contract *OpenZeppelinUpgradesECDSACaller // Generic read-only contract binding to access the raw methods on
}

// OpenZeppelinUpgradesECDSATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OpenZeppelinUpgradesECDSATransactorRaw struct {
	Contract *OpenZeppelinUpgradesECDSATransactor // Generic write-only contract binding to access the raw methods on
}

// NewOpenZeppelinUpgradesECDSA creates a new instance of OpenZeppelinUpgradesECDSA, bound to a specific deployed contract.
func NewOpenZeppelinUpgradesECDSA(address common.Address, backend bind.ContractBackend) (*OpenZeppelinUpgradesECDSA, error) {
	contract, err := bindOpenZeppelinUpgradesECDSA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OpenZeppelinUpgradesECDSA{OpenZeppelinUpgradesECDSACaller: OpenZeppelinUpgradesECDSACaller{contract: contract}, OpenZeppelinUpgradesECDSATransactor: OpenZeppelinUpgradesECDSATransactor{contract: contract}, OpenZeppelinUpgradesECDSAFilterer: OpenZeppelinUpgradesECDSAFilterer{contract: contract}}, nil
}

// NewOpenZeppelinUpgradesECDSACaller creates a new read-only instance of OpenZeppelinUpgradesECDSA, bound to a specific deployed contract.
func NewOpenZeppelinUpgradesECDSACaller(address common.Address, caller bind.ContractCaller) (*OpenZeppelinUpgradesECDSACaller, error) {
	contract, err := bindOpenZeppelinUpgradesECDSA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OpenZeppelinUpgradesECDSACaller{contract: contract}, nil
}

// NewOpenZeppelinUpgradesECDSATransactor creates a new write-only instance of OpenZeppelinUpgradesECDSA, bound to a specific deployed contract.
func NewOpenZeppelinUpgradesECDSATransactor(address common.Address, transactor bind.ContractTransactor) (*OpenZeppelinUpgradesECDSATransactor, error) {
	contract, err := bindOpenZeppelinUpgradesECDSA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OpenZeppelinUpgradesECDSATransactor{contract: contract}, nil
}

// NewOpenZeppelinUpgradesECDSAFilterer creates a new log filterer instance of OpenZeppelinUpgradesECDSA, bound to a specific deployed contract.
func NewOpenZeppelinUpgradesECDSAFilterer(address common.Address, filterer bind.ContractFilterer) (*OpenZeppelinUpgradesECDSAFilterer, error) {
	contract, err := bindOpenZeppelinUpgradesECDSA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OpenZeppelinUpgradesECDSAFilterer{contract: contract}, nil
}

// bindOpenZeppelinUpgradesECDSA binds a generic wrapper to an already deployed contract.
func bindOpenZeppelinUpgradesECDSA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OpenZeppelinUpgradesECDSAABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OpenZeppelinUpgradesECDSA *OpenZeppelinUpgradesECDSARaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OpenZeppelinUpgradesECDSA.Contract.OpenZeppelinUpgradesECDSACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OpenZeppelinUpgradesECDSA *OpenZeppelinUpgradesECDSARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpenZeppelinUpgradesECDSA.Contract.OpenZeppelinUpgradesECDSATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OpenZeppelinUpgradesECDSA *OpenZeppelinUpgradesECDSARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OpenZeppelinUpgradesECDSA.Contract.OpenZeppelinUpgradesECDSATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OpenZeppelinUpgradesECDSA *OpenZeppelinUpgradesECDSACallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OpenZeppelinUpgradesECDSA.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OpenZeppelinUpgradesECDSA *OpenZeppelinUpgradesECDSATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpenZeppelinUpgradesECDSA.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OpenZeppelinUpgradesECDSA *OpenZeppelinUpgradesECDSATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OpenZeppelinUpgradesECDSA.Contract.contract.Transact(opts, method, params...)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Ownable *OwnableCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Ownable *OwnableSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Ownable *OwnableCallerSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausableABI is the input ABI used to generate the binding from.
const PausableABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PausableFuncSigs maps the 4-byte function signature to its string representation.
var PausableFuncSigs = map[string]string{
	"5c975abb": "paused()",
}

// Pausable is an auto generated Go binding around an Ethereum contract.
type Pausable struct {
	PausableCaller     // Read-only binding to the contract
	PausableTransactor // Write-only binding to the contract
	PausableFilterer   // Log filterer for contract events
}

// PausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableSession struct {
	Contract     *Pausable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableCallerSession struct {
	Contract *PausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTransactorSession struct {
	Contract     *PausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableRaw struct {
	Contract *Pausable // Generic contract binding to access the raw methods on
}

// PausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableCallerRaw struct {
	Contract *PausableCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTransactorRaw struct {
	Contract *PausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausable creates a new instance of Pausable, bound to a specific deployed contract.
func NewPausable(address common.Address, backend bind.ContractBackend) (*Pausable, error) {
	contract, err := bindPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// NewPausableCaller creates a new read-only instance of Pausable, bound to a specific deployed contract.
func NewPausableCaller(address common.Address, caller bind.ContractCaller) (*PausableCaller, error) {
	contract, err := bindPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableCaller{contract: contract}, nil
}

// NewPausableTransactor creates a new write-only instance of Pausable, bound to a specific deployed contract.
func NewPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTransactor, error) {
	contract, err := bindPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTransactor{contract: contract}, nil
}

// NewPausableFilterer creates a new log filterer instance of Pausable, bound to a specific deployed contract.
func NewPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableFilterer, error) {
	contract, err := bindPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableFilterer{contract: contract}, nil
}

// bindPausable binds a generic wrapper to an already deployed contract.
func bindPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.PausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transact(opts, method, params...)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pausable.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableCallerSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// PausablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Pausable contract.
type PausablePausedIterator struct {
	Event *PausablePaused // Event containing the contract specifics and raw log

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
func (it *PausablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePaused)
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
		it.Event = new(PausablePaused)
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
func (it *PausablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePaused represents a Paused event raised by the Pausable contract.
type PausablePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) FilterPaused(opts *bind.FilterOpts) (*PausablePausedIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PausablePausedIterator{contract: _Pausable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PausablePaused) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePaused)
				if err := _Pausable.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Pausable *PausableFilterer) ParsePaused(log types.Log) (*PausablePaused, error) {
	event := new(PausablePaused)
	if err := _Pausable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Pausable contract.
type PausableUnpausedIterator struct {
	Event *PausableUnpaused // Event containing the contract specifics and raw log

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
func (it *PausableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUnpaused)
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
		it.Event = new(PausableUnpaused)
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
func (it *PausableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUnpaused represents a Unpaused event raised by the Pausable contract.
type PausableUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PausableUnpausedIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PausableUnpausedIterator{contract: _Pausable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PausableUnpaused) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUnpaused)
				if err := _Pausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Pausable *PausableFilterer) ParseUnpaused(log types.Log) (*PausableUnpaused, error) {
	event := new(PausableUnpaused)
	if err := _Pausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyABI is the input ABI used to generate the binding from.
const ProxyABI = "[{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// Proxy is an auto generated Go binding around an Ethereum contract.
type Proxy struct {
	ProxyCaller     // Read-only binding to the contract
	ProxyTransactor // Write-only binding to the contract
	ProxyFilterer   // Log filterer for contract events
}

// ProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxySession struct {
	Contract     *Proxy            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyCallerSession struct {
	Contract *ProxyCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyTransactorSession struct {
	Contract     *ProxyTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyRaw struct {
	Contract *Proxy // Generic contract binding to access the raw methods on
}

// ProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyCallerRaw struct {
	Contract *ProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyTransactorRaw struct {
	Contract *ProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxy creates a new instance of Proxy, bound to a specific deployed contract.
func NewProxy(address common.Address, backend bind.ContractBackend) (*Proxy, error) {
	contract, err := bindProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proxy{ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}, ProxyFilterer: ProxyFilterer{contract: contract}}, nil
}

// NewProxyCaller creates a new read-only instance of Proxy, bound to a specific deployed contract.
func NewProxyCaller(address common.Address, caller bind.ContractCaller) (*ProxyCaller, error) {
	contract, err := bindProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyCaller{contract: contract}, nil
}

// NewProxyTransactor creates a new write-only instance of Proxy, bound to a specific deployed contract.
func NewProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyTransactor, error) {
	contract, err := bindProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyTransactor{contract: contract}, nil
}

// NewProxyFilterer creates a new log filterer instance of Proxy, bound to a specific deployed contract.
func NewProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyFilterer, error) {
	contract, err := bindProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyFilterer{contract: contract}, nil
}

// bindProxy binds a generic wrapper to an already deployed contract.
func bindProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.ProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Proxy *ProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Proxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Proxy *ProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Proxy.Contract.Fallback(&_Proxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Proxy *ProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Proxy.Contract.Fallback(&_Proxy.TransactOpts, calldata)
}

// ProxyFactoryABI is the input ABI used to generate the binding from.
const ProxyFactoryABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"ProxyCreated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"deploy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"deploySigned\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"getDeploymentAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"getSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ProxyFactoryFuncSigs maps the 4-byte function signature to its string representation.
var ProxyFactoryFuncSigs = map[string]string{
	"6150864c": "deploy(uint256,address,address,bytes)",
	"332d6626": "deploySigned(uint256,address,address,bytes,bytes)",
	"81ae1f5b": "getDeploymentAddress(uint256,address)",
	"290f8f56": "getSigner(uint256,address,address,bytes,bytes)",
}

// ProxyFactoryBin is the compiled bytecode used for deploying new contracts.
var ProxyFactoryBin = "0x608060405234801561001057600080fd5b5060405161002060208201610044565b601f1982820381018352601f90910116604052805160209190910120600055610051565b61099c8061131183390190565b6112b1806100606000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063290f8f5614610051578063332d6626146101b65780636150864c146102ff57806381ae1f5b146103c3575b600080fd5b61019a600480360360a081101561006757600080fd5b8135916001600160a01b03602082013581169260408301359091169190810190608081016060820135600160201b8111156100a157600080fd5b8201836020820111156100b357600080fd5b803590602001918460018302840111600160201b831117156100d457600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561012657600080fd5b82018360208201111561013857600080fd5b803590602001918460018302840111600160201b8311171561015957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506103ef945050505050565b604080516001600160a01b039092168252519081900360200190f35b61019a600480360360a08110156101cc57600080fd5b8135916001600160a01b03602082013581169260408301359091169190810190608081016060820135600160201b81111561020657600080fd5b82018360208201111561021857600080fd5b803590602001918460018302840111600160201b8311171561023957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561028b57600080fd5b82018360208201111561029d57600080fd5b803590602001918460018302840111600160201b831117156102be57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506104de945050505050565b61019a6004803603608081101561031557600080fd5b8135916001600160a01b03602082013581169260408301359091169190810190608081016060820135600160201b81111561034f57600080fd5b82018360208201111561036157600080fd5b803590602001918460018302840111600160201b8311171561038257600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061054c945050505050565b61019a600480360360408110156103d957600080fd5b50803590602001356001600160a01b0316610564565b6000806104c7878787873060405160200180868152602001856001600160a01b03166001600160a01b031660601b8152601401846001600160a01b03166001600160a01b031660601b815260140183805190602001908083835b602083106104685780518252601f199092019160209182019101610449565b6001836020036101000a038019825116818451168082178552505050505050905001826001600160a01b03166001600160a01b031660601b815260140195505050505050604051602081830303815290604052805190602001206105cb565b90506104d3818461061c565b979650505050505050565b6000806104ee87878787876103ef565b90506001600160a01b03811661053f576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207369676e617475726560781b604482015290519081900360640190fd5b6104d3878787878561070a565b600061055b858585853361070a565b95945050505050565b600080610571848461083b565b600054604080516001600160f81b03196020808301919091523060601b6021830152603582019490945260558082019390935281518082039093018352607501905280519101206001600160a01b03169150505b92915050565b604080517f19457468657265756d205369676e6564204d6573736167653a0a333200000000602080830191909152603c8083019490945282518083039094018452605c909101909152815191012090565b6000815160411461062f575060006105c5565b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a082111561067557600093505050506105c5565b8060ff16601b1415801561068d57508060ff16601c14155b1561069e57600093505050506105c5565b6040805160008152602080820180845289905260ff8416828401526060820186905260808201859052915160019260a0808401939192601f1981019281900390910190855afa1580156106f5573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b600080610717878461087a565b604080516001600160a01b038316815290519192507efffc2da0b561cae30d9826d37709e9421c4725faebc226cbbb7ef5fc5e7349919081900360200190a160405163cf7a1d7760e01b81526001600160a01b038781166004830190815287821660248401526060604484019081528751606485015287519285169363cf7a1d77938b938b938b939192909160840190602085019080838360005b838110156107ca5781810151838201526020016107b2565b50505050905090810190601f1680156107f75780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b15801561081857600080fd5b505af115801561082c573d6000803e3d6000fd5b50929998505050505050505050565b6040805160208082019490945260609290921b6bffffffffffffffffffffffff1916828201528051808303603401815260549092019052805191012090565b60008060606040518060200161088f906108d3565b601f1982820381018352601f90910116604052905060006108b0868661083b565b9050808251602084016000f59250823b6108c957600080fd5b5090949350505050565b61099c806108e18339019056fe608060405234801561001057600080fd5b5061097c806100206000396000f3fe6080604052600436106100705760003560e01c80638f2839701161004e5780638f2839701461015e578063cf7a1d7714610191578063d1f5789414610250578063f851a4401461030657610070565b80633659cfe61461007a5780634f1ef286146100ad5780635c60da1b1461012d575b61007861031b565b005b34801561008657600080fd5b506100786004803603602081101561009d57600080fd5b50356001600160a01b0316610335565b610078600480360360408110156100c357600080fd5b6001600160a01b0382351691908101906040810160208201356401000000008111156100ee57600080fd5b82018360208201111561010057600080fd5b8035906020019184600183028401116401000000008311171561012257600080fd5b50909250905061036f565b34801561013957600080fd5b5061014261041c565b604080516001600160a01b039092168252519081900360200190f35b34801561016a57600080fd5b506100786004803603602081101561018157600080fd5b50356001600160a01b0316610459565b610078600480360360608110156101a757600080fd5b6001600160a01b0382358116926020810135909116918101906060810160408201356401000000008111156101db57600080fd5b8201836020820111156101ed57600080fd5b8035906020019184600183028401116401000000008311171561020f57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610513945050505050565b6100786004803603604081101561026657600080fd5b6001600160a01b03823516919081019060408101602082013564010000000081111561029157600080fd5b8201836020820111156102a357600080fd5b803590602001918460018302840111640100000000831117156102c557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610599945050505050565b34801561031257600080fd5b506101426106d9565b610323610704565b61033361032e610764565b610789565b565b61033d6107ad565b6001600160a01b0316336001600160a01b031614156103645761035f816107d2565b61036c565b61036c61031b565b50565b6103776107ad565b6001600160a01b0316336001600160a01b0316141561040f57610399836107d2565b6000836001600160a01b031683836040518083838082843760405192019450600093509091505080830381855af49150503d80600081146103f6576040519150601f19603f3d011682016040523d82523d6000602084013e6103fb565b606091505b505090508061040957600080fd5b50610417565b61041761031b565b505050565b60006104266107ad565b6001600160a01b0316336001600160a01b0316141561044e57610447610764565b9050610456565b61045661031b565b90565b6104616107ad565b6001600160a01b0316336001600160a01b03161415610364576001600160a01b0381166104bf5760405162461bcd60e51b81526004018080602001828103825260368152602001806108d76036913960400191505060405180910390fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6104e86107ad565b604080516001600160a01b03928316815291841660208301528051918290030190a161035f81610812565b600061051d610764565b6001600160a01b03161461053057600080fd5b61053a8382610599565b604080517232b4b8189c9b1b97383937bc3c9730b236b4b760691b815290519081900360130190207fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61036000199091011461059057fe5b61041782610812565b60006105a3610764565b6001600160a01b0316146105b657600080fd5b604080517f656970313936372e70726f78792e696d706c656d656e746174696f6e000000008152905190819003601c0190207f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6000199091011461061657fe5b61061f82610836565b8051156106d5576000826001600160a01b0316826040518082805190602001908083835b602083106106625780518252601f199092019160209182019101610643565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d80600081146106c2576040519150601f19603f3d011682016040523d82523d6000602084013e6106c7565b606091505b505090508061041757600080fd5b5050565b60006106e36107ad565b6001600160a01b0316336001600160a01b0316141561044e576104476107ad565b61070c6107ad565b6001600160a01b0316336001600160a01b0316141561075c5760405162461bcd60e51b81526004018080602001828103825260328152602001806108a56032913960400191505060405180910390fd5b610333610333565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5490565b3660008037600080366000845af43d6000803e8080156107a8573d6000f35b3d6000fd5b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035490565b6107db81610836565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610355565b61083f8161089e565b61087a5760405162461bcd60e51b815260040180806020018281038252603b81526020018061090d603b913960400191505060405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc55565b3b15159056fe43616e6e6f742063616c6c2066616c6c6261636b2066756e6374696f6e2066726f6d207468652070726f78792061646d696e43616e6e6f74206368616e6765207468652061646d696e206f6620612070726f787920746f20746865207a65726f206164647265737343616e6e6f742073657420612070726f787920696d706c656d656e746174696f6e20746f2061206e6f6e2d636f6e74726163742061646472657373a265627a7a72315820ea63902a998581c26cce013e1736a0cabdc6bc56ff51f923fdb0a60bef5645e364736f6c63430005110032a265627a7a72315820fa191803518ae4c17f9fd40949f1cf40997bce83b83595ebe363f1954730d93564736f6c63430005110032608060405234801561001057600080fd5b5061097c806100206000396000f3fe6080604052600436106100705760003560e01c80638f2839701161004e5780638f2839701461015e578063cf7a1d7714610191578063d1f5789414610250578063f851a4401461030657610070565b80633659cfe61461007a5780634f1ef286146100ad5780635c60da1b1461012d575b61007861031b565b005b34801561008657600080fd5b506100786004803603602081101561009d57600080fd5b50356001600160a01b0316610335565b610078600480360360408110156100c357600080fd5b6001600160a01b0382351691908101906040810160208201356401000000008111156100ee57600080fd5b82018360208201111561010057600080fd5b8035906020019184600183028401116401000000008311171561012257600080fd5b50909250905061036f565b34801561013957600080fd5b5061014261041c565b604080516001600160a01b039092168252519081900360200190f35b34801561016a57600080fd5b506100786004803603602081101561018157600080fd5b50356001600160a01b0316610459565b610078600480360360608110156101a757600080fd5b6001600160a01b0382358116926020810135909116918101906060810160408201356401000000008111156101db57600080fd5b8201836020820111156101ed57600080fd5b8035906020019184600183028401116401000000008311171561020f57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610513945050505050565b6100786004803603604081101561026657600080fd5b6001600160a01b03823516919081019060408101602082013564010000000081111561029157600080fd5b8201836020820111156102a357600080fd5b803590602001918460018302840111640100000000831117156102c557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610599945050505050565b34801561031257600080fd5b506101426106d9565b610323610704565b61033361032e610764565b610789565b565b61033d6107ad565b6001600160a01b0316336001600160a01b031614156103645761035f816107d2565b61036c565b61036c61031b565b50565b6103776107ad565b6001600160a01b0316336001600160a01b0316141561040f57610399836107d2565b6000836001600160a01b031683836040518083838082843760405192019450600093509091505080830381855af49150503d80600081146103f6576040519150601f19603f3d011682016040523d82523d6000602084013e6103fb565b606091505b505090508061040957600080fd5b50610417565b61041761031b565b505050565b60006104266107ad565b6001600160a01b0316336001600160a01b0316141561044e57610447610764565b9050610456565b61045661031b565b90565b6104616107ad565b6001600160a01b0316336001600160a01b03161415610364576001600160a01b0381166104bf5760405162461bcd60e51b81526004018080602001828103825260368152602001806108d76036913960400191505060405180910390fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6104e86107ad565b604080516001600160a01b03928316815291841660208301528051918290030190a161035f81610812565b600061051d610764565b6001600160a01b03161461053057600080fd5b61053a8382610599565b604080517232b4b8189c9b1b97383937bc3c9730b236b4b760691b815290519081900360130190207fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61036000199091011461059057fe5b61041782610812565b60006105a3610764565b6001600160a01b0316146105b657600080fd5b604080517f656970313936372e70726f78792e696d706c656d656e746174696f6e000000008152905190819003601c0190207f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6000199091011461061657fe5b61061f82610836565b8051156106d5576000826001600160a01b0316826040518082805190602001908083835b602083106106625780518252601f199092019160209182019101610643565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d80600081146106c2576040519150601f19603f3d011682016040523d82523d6000602084013e6106c7565b606091505b505090508061041757600080fd5b5050565b60006106e36107ad565b6001600160a01b0316336001600160a01b0316141561044e576104476107ad565b61070c6107ad565b6001600160a01b0316336001600160a01b0316141561075c5760405162461bcd60e51b81526004018080602001828103825260328152602001806108a56032913960400191505060405180910390fd5b610333610333565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5490565b3660008037600080366000845af43d6000803e8080156107a8573d6000f35b3d6000fd5b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035490565b6107db81610836565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610355565b61083f8161089e565b61087a5760405162461bcd60e51b815260040180806020018281038252603b81526020018061090d603b913960400191505060405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc55565b3b15159056fe43616e6e6f742063616c6c2066616c6c6261636b2066756e6374696f6e2066726f6d207468652070726f78792061646d696e43616e6e6f74206368616e6765207468652061646d696e206f6620612070726f787920746f20746865207a65726f206164647265737343616e6e6f742073657420612070726f787920696d706c656d656e746174696f6e20746f2061206e6f6e2d636f6e74726163742061646472657373a265627a7a72315820ea63902a998581c26cce013e1736a0cabdc6bc56ff51f923fdb0a60bef5645e364736f6c63430005110032"

// DeployProxyFactory deploys a new Ethereum contract, binding an instance of ProxyFactory to it.
func DeployProxyFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ProxyFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProxyFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProxyFactory{ProxyFactoryCaller: ProxyFactoryCaller{contract: contract}, ProxyFactoryTransactor: ProxyFactoryTransactor{contract: contract}, ProxyFactoryFilterer: ProxyFactoryFilterer{contract: contract}}, nil
}

// ProxyFactory is an auto generated Go binding around an Ethereum contract.
type ProxyFactory struct {
	ProxyFactoryCaller     // Read-only binding to the contract
	ProxyFactoryTransactor // Write-only binding to the contract
	ProxyFactoryFilterer   // Log filterer for contract events
}

// ProxyFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxyFactorySession struct {
	Contract     *ProxyFactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyFactoryCallerSession struct {
	Contract *ProxyFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ProxyFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyFactoryTransactorSession struct {
	Contract     *ProxyFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ProxyFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyFactoryRaw struct {
	Contract *ProxyFactory // Generic contract binding to access the raw methods on
}

// ProxyFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyFactoryCallerRaw struct {
	Contract *ProxyFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyFactoryTransactorRaw struct {
	Contract *ProxyFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxyFactory creates a new instance of ProxyFactory, bound to a specific deployed contract.
func NewProxyFactory(address common.Address, backend bind.ContractBackend) (*ProxyFactory, error) {
	contract, err := bindProxyFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProxyFactory{ProxyFactoryCaller: ProxyFactoryCaller{contract: contract}, ProxyFactoryTransactor: ProxyFactoryTransactor{contract: contract}, ProxyFactoryFilterer: ProxyFactoryFilterer{contract: contract}}, nil
}

// NewProxyFactoryCaller creates a new read-only instance of ProxyFactory, bound to a specific deployed contract.
func NewProxyFactoryCaller(address common.Address, caller bind.ContractCaller) (*ProxyFactoryCaller, error) {
	contract, err := bindProxyFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyFactoryCaller{contract: contract}, nil
}

// NewProxyFactoryTransactor creates a new write-only instance of ProxyFactory, bound to a specific deployed contract.
func NewProxyFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyFactoryTransactor, error) {
	contract, err := bindProxyFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyFactoryTransactor{contract: contract}, nil
}

// NewProxyFactoryFilterer creates a new log filterer instance of ProxyFactory, bound to a specific deployed contract.
func NewProxyFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyFactoryFilterer, error) {
	contract, err := bindProxyFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyFactoryFilterer{contract: contract}, nil
}

// bindProxyFactory binds a generic wrapper to an already deployed contract.
func bindProxyFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyFactory *ProxyFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProxyFactory.Contract.ProxyFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyFactory *ProxyFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyFactory.Contract.ProxyFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyFactory *ProxyFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyFactory.Contract.ProxyFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyFactory *ProxyFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProxyFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyFactory *ProxyFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyFactory *ProxyFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyFactory.Contract.contract.Transact(opts, method, params...)
}

// GetDeploymentAddress is a free data retrieval call binding the contract method 0x81ae1f5b.
//
// Solidity: function getDeploymentAddress(uint256 _salt, address _sender) view returns(address)
func (_ProxyFactory *ProxyFactoryCaller) GetDeploymentAddress(opts *bind.CallOpts, _salt *big.Int, _sender common.Address) (common.Address, error) {
	var out []interface{}
	err := _ProxyFactory.contract.Call(opts, &out, "getDeploymentAddress", _salt, _sender)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDeploymentAddress is a free data retrieval call binding the contract method 0x81ae1f5b.
//
// Solidity: function getDeploymentAddress(uint256 _salt, address _sender) view returns(address)
func (_ProxyFactory *ProxyFactorySession) GetDeploymentAddress(_salt *big.Int, _sender common.Address) (common.Address, error) {
	return _ProxyFactory.Contract.GetDeploymentAddress(&_ProxyFactory.CallOpts, _salt, _sender)
}

// GetDeploymentAddress is a free data retrieval call binding the contract method 0x81ae1f5b.
//
// Solidity: function getDeploymentAddress(uint256 _salt, address _sender) view returns(address)
func (_ProxyFactory *ProxyFactoryCallerSession) GetDeploymentAddress(_salt *big.Int, _sender common.Address) (common.Address, error) {
	return _ProxyFactory.Contract.GetDeploymentAddress(&_ProxyFactory.CallOpts, _salt, _sender)
}

// GetSigner is a free data retrieval call binding the contract method 0x290f8f56.
//
// Solidity: function getSigner(uint256 _salt, address _logic, address _admin, bytes _data, bytes _signature) view returns(address)
func (_ProxyFactory *ProxyFactoryCaller) GetSigner(opts *bind.CallOpts, _salt *big.Int, _logic common.Address, _admin common.Address, _data []byte, _signature []byte) (common.Address, error) {
	var out []interface{}
	err := _ProxyFactory.contract.Call(opts, &out, "getSigner", _salt, _logic, _admin, _data, _signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSigner is a free data retrieval call binding the contract method 0x290f8f56.
//
// Solidity: function getSigner(uint256 _salt, address _logic, address _admin, bytes _data, bytes _signature) view returns(address)
func (_ProxyFactory *ProxyFactorySession) GetSigner(_salt *big.Int, _logic common.Address, _admin common.Address, _data []byte, _signature []byte) (common.Address, error) {
	return _ProxyFactory.Contract.GetSigner(&_ProxyFactory.CallOpts, _salt, _logic, _admin, _data, _signature)
}

// GetSigner is a free data retrieval call binding the contract method 0x290f8f56.
//
// Solidity: function getSigner(uint256 _salt, address _logic, address _admin, bytes _data, bytes _signature) view returns(address)
func (_ProxyFactory *ProxyFactoryCallerSession) GetSigner(_salt *big.Int, _logic common.Address, _admin common.Address, _data []byte, _signature []byte) (common.Address, error) {
	return _ProxyFactory.Contract.GetSigner(&_ProxyFactory.CallOpts, _salt, _logic, _admin, _data, _signature)
}

// Deploy is a paid mutator transaction binding the contract method 0x6150864c.
//
// Solidity: function deploy(uint256 _salt, address _logic, address _admin, bytes _data) returns(address)
func (_ProxyFactory *ProxyFactoryTransactor) Deploy(opts *bind.TransactOpts, _salt *big.Int, _logic common.Address, _admin common.Address, _data []byte) (*types.Transaction, error) {
	return _ProxyFactory.contract.Transact(opts, "deploy", _salt, _logic, _admin, _data)
}

// Deploy is a paid mutator transaction binding the contract method 0x6150864c.
//
// Solidity: function deploy(uint256 _salt, address _logic, address _admin, bytes _data) returns(address)
func (_ProxyFactory *ProxyFactorySession) Deploy(_salt *big.Int, _logic common.Address, _admin common.Address, _data []byte) (*types.Transaction, error) {
	return _ProxyFactory.Contract.Deploy(&_ProxyFactory.TransactOpts, _salt, _logic, _admin, _data)
}

// Deploy is a paid mutator transaction binding the contract method 0x6150864c.
//
// Solidity: function deploy(uint256 _salt, address _logic, address _admin, bytes _data) returns(address)
func (_ProxyFactory *ProxyFactoryTransactorSession) Deploy(_salt *big.Int, _logic common.Address, _admin common.Address, _data []byte) (*types.Transaction, error) {
	return _ProxyFactory.Contract.Deploy(&_ProxyFactory.TransactOpts, _salt, _logic, _admin, _data)
}

// DeploySigned is a paid mutator transaction binding the contract method 0x332d6626.
//
// Solidity: function deploySigned(uint256 _salt, address _logic, address _admin, bytes _data, bytes _signature) returns(address)
func (_ProxyFactory *ProxyFactoryTransactor) DeploySigned(opts *bind.TransactOpts, _salt *big.Int, _logic common.Address, _admin common.Address, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _ProxyFactory.contract.Transact(opts, "deploySigned", _salt, _logic, _admin, _data, _signature)
}

// DeploySigned is a paid mutator transaction binding the contract method 0x332d6626.
//
// Solidity: function deploySigned(uint256 _salt, address _logic, address _admin, bytes _data, bytes _signature) returns(address)
func (_ProxyFactory *ProxyFactorySession) DeploySigned(_salt *big.Int, _logic common.Address, _admin common.Address, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _ProxyFactory.Contract.DeploySigned(&_ProxyFactory.TransactOpts, _salt, _logic, _admin, _data, _signature)
}

// DeploySigned is a paid mutator transaction binding the contract method 0x332d6626.
//
// Solidity: function deploySigned(uint256 _salt, address _logic, address _admin, bytes _data, bytes _signature) returns(address)
func (_ProxyFactory *ProxyFactoryTransactorSession) DeploySigned(_salt *big.Int, _logic common.Address, _admin common.Address, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _ProxyFactory.Contract.DeploySigned(&_ProxyFactory.TransactOpts, _salt, _logic, _admin, _data, _signature)
}

// ProxyFactoryProxyCreatedIterator is returned from FilterProxyCreated and is used to iterate over the raw logs and unpacked data for ProxyCreated events raised by the ProxyFactory contract.
type ProxyFactoryProxyCreatedIterator struct {
	Event *ProxyFactoryProxyCreated // Event containing the contract specifics and raw log

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
func (it *ProxyFactoryProxyCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyFactoryProxyCreated)
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
		it.Event = new(ProxyFactoryProxyCreated)
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
func (it *ProxyFactoryProxyCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyFactoryProxyCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyFactoryProxyCreated represents a ProxyCreated event raised by the ProxyFactory contract.
type ProxyFactoryProxyCreated struct {
	Proxy common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterProxyCreated is a free log retrieval operation binding the contract event 0x00fffc2da0b561cae30d9826d37709e9421c4725faebc226cbbb7ef5fc5e7349.
//
// Solidity: event ProxyCreated(address proxy)
func (_ProxyFactory *ProxyFactoryFilterer) FilterProxyCreated(opts *bind.FilterOpts) (*ProxyFactoryProxyCreatedIterator, error) {

	logs, sub, err := _ProxyFactory.contract.FilterLogs(opts, "ProxyCreated")
	if err != nil {
		return nil, err
	}
	return &ProxyFactoryProxyCreatedIterator{contract: _ProxyFactory.contract, event: "ProxyCreated", logs: logs, sub: sub}, nil
}

// WatchProxyCreated is a free log subscription operation binding the contract event 0x00fffc2da0b561cae30d9826d37709e9421c4725faebc226cbbb7ef5fc5e7349.
//
// Solidity: event ProxyCreated(address proxy)
func (_ProxyFactory *ProxyFactoryFilterer) WatchProxyCreated(opts *bind.WatchOpts, sink chan<- *ProxyFactoryProxyCreated) (event.Subscription, error) {

	logs, sub, err := _ProxyFactory.contract.WatchLogs(opts, "ProxyCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyFactoryProxyCreated)
				if err := _ProxyFactory.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
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

// ParseProxyCreated is a log parse operation binding the contract event 0x00fffc2da0b561cae30d9826d37709e9421c4725faebc226cbbb7ef5fc5e7349.
//
// Solidity: event ProxyCreated(address proxy)
func (_ProxyFactory *ProxyFactoryFilterer) ParseProxyCreated(log types.Log) (*ProxyFactoryProxyCreated, error) {
	event := new(ProxyFactoryProxyCreated)
	if err := _ProxyFactory.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158201999f7fb132531516a60123b92b4110398e9e740ff99ee477c2686e56380ffbb64736f6c63430005110032"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// UpgradeabilityProxyABI is the input ABI used to generate the binding from.
const UpgradeabilityProxyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// UpgradeabilityProxyBin is the compiled bytecode used for deploying new contracts.
var UpgradeabilityProxyBin = "0x608060405260405161036d38038061036d8339818101604052604081101561002657600080fd5b81516020830180516040519294929383019291908464010000000082111561004d57600080fd5b90830190602082018581111561006257600080fd5b825164010000000081118282018810171561007c57600080fd5b82525081516020918201929091019080838360005b838110156100a9578181015183820152602001610091565b50505050905090810190601f1680156100d65780820380516001836020036101000a031916815260200191505b5060408181527f656970313936372e70726f78792e696d706c656d656e746174696f6e0000000082525190819003601c01902060008051602061031283398151915260001990910114925061012a91505057fe5b61013c826001600160e01b036101fb16565b8051156101f4576000826001600160a01b0316826040518082805190602001908083835b6020831061017f5780518252601f199092019160209182019101610160565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d80600081146101df576040519150601f19603f3d011682016040523d82523d6000602084013e6101e4565b606091505b50509050806101f257600080fd5b505b5050610261565b61020e8161025b60201b6100681760201c565b6102495760405162461bcd60e51b815260040180806020018281038252603b815260200180610332603b913960400191505060405180910390fd5b60008051602061031283398151915255565b3b151590565b60a38061026f6000396000f3fe6080604052600a600c565b005b6012601e565b601e601a6020565b6045565b565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5490565b3660008037600080366000845af43d6000803e8080156063573d6000f35b3d6000fd5b3b15159056fea265627a7a72315820a329c260694b97ebe02d348a617858573a5130297693e8226125a22f4929311364736f6c63430005110032360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc43616e6e6f742073657420612070726f787920696d706c656d656e746174696f6e20746f2061206e6f6e2d636f6e74726163742061646472657373"

// DeployUpgradeabilityProxy deploys a new Ethereum contract, binding an instance of UpgradeabilityProxy to it.
func DeployUpgradeabilityProxy(auth *bind.TransactOpts, backend bind.ContractBackend, _logic common.Address, _data []byte) (common.Address, *types.Transaction, *UpgradeabilityProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(UpgradeabilityProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UpgradeabilityProxyBin), backend, _logic, _data)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UpgradeabilityProxy{UpgradeabilityProxyCaller: UpgradeabilityProxyCaller{contract: contract}, UpgradeabilityProxyTransactor: UpgradeabilityProxyTransactor{contract: contract}, UpgradeabilityProxyFilterer: UpgradeabilityProxyFilterer{contract: contract}}, nil
}

// UpgradeabilityProxy is an auto generated Go binding around an Ethereum contract.
type UpgradeabilityProxy struct {
	UpgradeabilityProxyCaller     // Read-only binding to the contract
	UpgradeabilityProxyTransactor // Write-only binding to the contract
	UpgradeabilityProxyFilterer   // Log filterer for contract events
}

// UpgradeabilityProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type UpgradeabilityProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeabilityProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UpgradeabilityProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeabilityProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UpgradeabilityProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeabilityProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UpgradeabilityProxySession struct {
	Contract     *UpgradeabilityProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// UpgradeabilityProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UpgradeabilityProxyCallerSession struct {
	Contract *UpgradeabilityProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// UpgradeabilityProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UpgradeabilityProxyTransactorSession struct {
	Contract     *UpgradeabilityProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// UpgradeabilityProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type UpgradeabilityProxyRaw struct {
	Contract *UpgradeabilityProxy // Generic contract binding to access the raw methods on
}

// UpgradeabilityProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UpgradeabilityProxyCallerRaw struct {
	Contract *UpgradeabilityProxyCaller // Generic read-only contract binding to access the raw methods on
}

// UpgradeabilityProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UpgradeabilityProxyTransactorRaw struct {
	Contract *UpgradeabilityProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUpgradeabilityProxy creates a new instance of UpgradeabilityProxy, bound to a specific deployed contract.
func NewUpgradeabilityProxy(address common.Address, backend bind.ContractBackend) (*UpgradeabilityProxy, error) {
	contract, err := bindUpgradeabilityProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UpgradeabilityProxy{UpgradeabilityProxyCaller: UpgradeabilityProxyCaller{contract: contract}, UpgradeabilityProxyTransactor: UpgradeabilityProxyTransactor{contract: contract}, UpgradeabilityProxyFilterer: UpgradeabilityProxyFilterer{contract: contract}}, nil
}

// NewUpgradeabilityProxyCaller creates a new read-only instance of UpgradeabilityProxy, bound to a specific deployed contract.
func NewUpgradeabilityProxyCaller(address common.Address, caller bind.ContractCaller) (*UpgradeabilityProxyCaller, error) {
	contract, err := bindUpgradeabilityProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeabilityProxyCaller{contract: contract}, nil
}

// NewUpgradeabilityProxyTransactor creates a new write-only instance of UpgradeabilityProxy, bound to a specific deployed contract.
func NewUpgradeabilityProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*UpgradeabilityProxyTransactor, error) {
	contract, err := bindUpgradeabilityProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeabilityProxyTransactor{contract: contract}, nil
}

// NewUpgradeabilityProxyFilterer creates a new log filterer instance of UpgradeabilityProxy, bound to a specific deployed contract.
func NewUpgradeabilityProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*UpgradeabilityProxyFilterer, error) {
	contract, err := bindUpgradeabilityProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UpgradeabilityProxyFilterer{contract: contract}, nil
}

// bindUpgradeabilityProxy binds a generic wrapper to an already deployed contract.
func bindUpgradeabilityProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UpgradeabilityProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeabilityProxy *UpgradeabilityProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpgradeabilityProxy.Contract.UpgradeabilityProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeabilityProxy *UpgradeabilityProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeabilityProxy.Contract.UpgradeabilityProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeabilityProxy *UpgradeabilityProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeabilityProxy.Contract.UpgradeabilityProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeabilityProxy *UpgradeabilityProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpgradeabilityProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeabilityProxy *UpgradeabilityProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeabilityProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeabilityProxy *UpgradeabilityProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeabilityProxy.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UpgradeabilityProxy *UpgradeabilityProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _UpgradeabilityProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UpgradeabilityProxy *UpgradeabilityProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _UpgradeabilityProxy.Contract.Fallback(&_UpgradeabilityProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UpgradeabilityProxy *UpgradeabilityProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _UpgradeabilityProxy.Contract.Fallback(&_UpgradeabilityProxy.TransactOpts, calldata)
}

// UpgradeabilityProxyUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the UpgradeabilityProxy contract.
type UpgradeabilityProxyUpgradedIterator struct {
	Event *UpgradeabilityProxyUpgraded // Event containing the contract specifics and raw log

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
func (it *UpgradeabilityProxyUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeabilityProxyUpgraded)
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
		it.Event = new(UpgradeabilityProxyUpgraded)
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
func (it *UpgradeabilityProxyUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeabilityProxyUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeabilityProxyUpgraded represents a Upgraded event raised by the UpgradeabilityProxy contract.
type UpgradeabilityProxyUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UpgradeabilityProxy *UpgradeabilityProxyFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*UpgradeabilityProxyUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _UpgradeabilityProxy.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeabilityProxyUpgradedIterator{contract: _UpgradeabilityProxy.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UpgradeabilityProxy *UpgradeabilityProxyFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *UpgradeabilityProxyUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _UpgradeabilityProxy.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeabilityProxyUpgraded)
				if err := _UpgradeabilityProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UpgradeabilityProxy *UpgradeabilityProxyFilterer) ParseUpgraded(log types.Log) (*UpgradeabilityProxyUpgraded, error) {
	event := new(UpgradeabilityProxyUpgraded)
	if err := _UpgradeabilityProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
