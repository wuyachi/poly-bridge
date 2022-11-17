// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wrapper_abi

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

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILockProxyABI is the input ABI used to generate the binding from.
const ILockProxyABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"assetHashMap\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAssetHash\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"toAssetHash\",\"type\":\"bytes\"}],\"name\":\"bindAssetHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"targetProxyHash\",\"type\":\"bytes\"}],\"name\":\"bindProxyHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getBalanceFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAssetHash\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"managerProxyContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"proxyHashMap\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"eccmpAddr\",\"type\":\"address\"}],\"name\":\"setManagerProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ILockProxyFuncSigs maps the 4-byte function signature to its string representation.
var ILockProxyFuncSigs = map[string]string{
	"4f7d9808": "assetHashMap(address,uint64)",
	"3348f63b": "bindAssetHash(address,uint64,bytes)",
	"379b98f6": "bindProxyHash(uint64,bytes)",
	"59c589a1": "getBalanceFor(address)",
	"84a6d055": "lock(address,uint64,bytes,uint256)",
	"d798f881": "managerProxyContract()",
	"9e5767aa": "proxyHashMap(uint64)",
	"af9980f0": "setManagerProxy(address)",
}

// ILockProxy is an auto generated Go binding around an Ethereum contract.
type ILockProxy struct {
	ILockProxyCaller     // Read-only binding to the contract
	ILockProxyTransactor // Write-only binding to the contract
	ILockProxyFilterer   // Log filterer for contract events
}

// ILockProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ILockProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILockProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ILockProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILockProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ILockProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILockProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ILockProxySession struct {
	Contract     *ILockProxy       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ILockProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ILockProxyCallerSession struct {
	Contract *ILockProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ILockProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ILockProxyTransactorSession struct {
	Contract     *ILockProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ILockProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ILockProxyRaw struct {
	Contract *ILockProxy // Generic contract binding to access the raw methods on
}

// ILockProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ILockProxyCallerRaw struct {
	Contract *ILockProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ILockProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ILockProxyTransactorRaw struct {
	Contract *ILockProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewILockProxy creates a new instance of ILockProxy, bound to a specific deployed contract.
func NewILockProxy(address common.Address, backend bind.ContractBackend) (*ILockProxy, error) {
	contract, err := bindILockProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ILockProxy{ILockProxyCaller: ILockProxyCaller{contract: contract}, ILockProxyTransactor: ILockProxyTransactor{contract: contract}, ILockProxyFilterer: ILockProxyFilterer{contract: contract}}, nil
}

// NewILockProxyCaller creates a new read-only instance of ILockProxy, bound to a specific deployed contract.
func NewILockProxyCaller(address common.Address, caller bind.ContractCaller) (*ILockProxyCaller, error) {
	contract, err := bindILockProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ILockProxyCaller{contract: contract}, nil
}

// NewILockProxyTransactor creates a new write-only instance of ILockProxy, bound to a specific deployed contract.
func NewILockProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ILockProxyTransactor, error) {
	contract, err := bindILockProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ILockProxyTransactor{contract: contract}, nil
}

// NewILockProxyFilterer creates a new log filterer instance of ILockProxy, bound to a specific deployed contract.
func NewILockProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ILockProxyFilterer, error) {
	contract, err := bindILockProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ILockProxyFilterer{contract: contract}, nil
}

// bindILockProxy binds a generic wrapper to an already deployed contract.
func bindILockProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ILockProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILockProxy *ILockProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILockProxy.Contract.ILockProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILockProxy *ILockProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILockProxy.Contract.ILockProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILockProxy *ILockProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILockProxy.Contract.ILockProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILockProxy *ILockProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILockProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILockProxy *ILockProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILockProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILockProxy *ILockProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILockProxy.Contract.contract.Transact(opts, method, params...)
}

// AssetHashMap is a free data retrieval call binding the contract method 0x4f7d9808.
//
// Solidity: function assetHashMap(address , uint64 ) view returns(bytes)
func (_ILockProxy *ILockProxyCaller) AssetHashMap(opts *bind.CallOpts, arg0 common.Address, arg1 uint64) ([]byte, error) {
	var out []interface{}
	err := _ILockProxy.contract.Call(opts, &out, "assetHashMap", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// AssetHashMap is a free data retrieval call binding the contract method 0x4f7d9808.
//
// Solidity: function assetHashMap(address , uint64 ) view returns(bytes)
func (_ILockProxy *ILockProxySession) AssetHashMap(arg0 common.Address, arg1 uint64) ([]byte, error) {
	return _ILockProxy.Contract.AssetHashMap(&_ILockProxy.CallOpts, arg0, arg1)
}

// AssetHashMap is a free data retrieval call binding the contract method 0x4f7d9808.
//
// Solidity: function assetHashMap(address , uint64 ) view returns(bytes)
func (_ILockProxy *ILockProxyCallerSession) AssetHashMap(arg0 common.Address, arg1 uint64) ([]byte, error) {
	return _ILockProxy.Contract.AssetHashMap(&_ILockProxy.CallOpts, arg0, arg1)
}

// GetBalanceFor is a free data retrieval call binding the contract method 0x59c589a1.
//
// Solidity: function getBalanceFor(address ) view returns(uint256)
func (_ILockProxy *ILockProxyCaller) GetBalanceFor(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ILockProxy.contract.Call(opts, &out, "getBalanceFor", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalanceFor is a free data retrieval call binding the contract method 0x59c589a1.
//
// Solidity: function getBalanceFor(address ) view returns(uint256)
func (_ILockProxy *ILockProxySession) GetBalanceFor(arg0 common.Address) (*big.Int, error) {
	return _ILockProxy.Contract.GetBalanceFor(&_ILockProxy.CallOpts, arg0)
}

// GetBalanceFor is a free data retrieval call binding the contract method 0x59c589a1.
//
// Solidity: function getBalanceFor(address ) view returns(uint256)
func (_ILockProxy *ILockProxyCallerSession) GetBalanceFor(arg0 common.Address) (*big.Int, error) {
	return _ILockProxy.Contract.GetBalanceFor(&_ILockProxy.CallOpts, arg0)
}

// ManagerProxyContract is a free data retrieval call binding the contract method 0xd798f881.
//
// Solidity: function managerProxyContract() view returns(address)
func (_ILockProxy *ILockProxyCaller) ManagerProxyContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILockProxy.contract.Call(opts, &out, "managerProxyContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ManagerProxyContract is a free data retrieval call binding the contract method 0xd798f881.
//
// Solidity: function managerProxyContract() view returns(address)
func (_ILockProxy *ILockProxySession) ManagerProxyContract() (common.Address, error) {
	return _ILockProxy.Contract.ManagerProxyContract(&_ILockProxy.CallOpts)
}

// ManagerProxyContract is a free data retrieval call binding the contract method 0xd798f881.
//
// Solidity: function managerProxyContract() view returns(address)
func (_ILockProxy *ILockProxyCallerSession) ManagerProxyContract() (common.Address, error) {
	return _ILockProxy.Contract.ManagerProxyContract(&_ILockProxy.CallOpts)
}

// ProxyHashMap is a free data retrieval call binding the contract method 0x9e5767aa.
//
// Solidity: function proxyHashMap(uint64 ) view returns(bytes)
func (_ILockProxy *ILockProxyCaller) ProxyHashMap(opts *bind.CallOpts, arg0 uint64) ([]byte, error) {
	var out []interface{}
	err := _ILockProxy.contract.Call(opts, &out, "proxyHashMap", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ProxyHashMap is a free data retrieval call binding the contract method 0x9e5767aa.
//
// Solidity: function proxyHashMap(uint64 ) view returns(bytes)
func (_ILockProxy *ILockProxySession) ProxyHashMap(arg0 uint64) ([]byte, error) {
	return _ILockProxy.Contract.ProxyHashMap(&_ILockProxy.CallOpts, arg0)
}

// ProxyHashMap is a free data retrieval call binding the contract method 0x9e5767aa.
//
// Solidity: function proxyHashMap(uint64 ) view returns(bytes)
func (_ILockProxy *ILockProxyCallerSession) ProxyHashMap(arg0 uint64) ([]byte, error) {
	return _ILockProxy.Contract.ProxyHashMap(&_ILockProxy.CallOpts, arg0)
}

// BindAssetHash is a paid mutator transaction binding the contract method 0x3348f63b.
//
// Solidity: function bindAssetHash(address fromAssetHash, uint64 toChainId, bytes toAssetHash) returns(bool)
func (_ILockProxy *ILockProxyTransactor) BindAssetHash(opts *bind.TransactOpts, fromAssetHash common.Address, toChainId uint64, toAssetHash []byte) (*types.Transaction, error) {
	return _ILockProxy.contract.Transact(opts, "bindAssetHash", fromAssetHash, toChainId, toAssetHash)
}

// BindAssetHash is a paid mutator transaction binding the contract method 0x3348f63b.
//
// Solidity: function bindAssetHash(address fromAssetHash, uint64 toChainId, bytes toAssetHash) returns(bool)
func (_ILockProxy *ILockProxySession) BindAssetHash(fromAssetHash common.Address, toChainId uint64, toAssetHash []byte) (*types.Transaction, error) {
	return _ILockProxy.Contract.BindAssetHash(&_ILockProxy.TransactOpts, fromAssetHash, toChainId, toAssetHash)
}

// BindAssetHash is a paid mutator transaction binding the contract method 0x3348f63b.
//
// Solidity: function bindAssetHash(address fromAssetHash, uint64 toChainId, bytes toAssetHash) returns(bool)
func (_ILockProxy *ILockProxyTransactorSession) BindAssetHash(fromAssetHash common.Address, toChainId uint64, toAssetHash []byte) (*types.Transaction, error) {
	return _ILockProxy.Contract.BindAssetHash(&_ILockProxy.TransactOpts, fromAssetHash, toChainId, toAssetHash)
}

// BindProxyHash is a paid mutator transaction binding the contract method 0x379b98f6.
//
// Solidity: function bindProxyHash(uint64 toChainId, bytes targetProxyHash) returns(bool)
func (_ILockProxy *ILockProxyTransactor) BindProxyHash(opts *bind.TransactOpts, toChainId uint64, targetProxyHash []byte) (*types.Transaction, error) {
	return _ILockProxy.contract.Transact(opts, "bindProxyHash", toChainId, targetProxyHash)
}

// BindProxyHash is a paid mutator transaction binding the contract method 0x379b98f6.
//
// Solidity: function bindProxyHash(uint64 toChainId, bytes targetProxyHash) returns(bool)
func (_ILockProxy *ILockProxySession) BindProxyHash(toChainId uint64, targetProxyHash []byte) (*types.Transaction, error) {
	return _ILockProxy.Contract.BindProxyHash(&_ILockProxy.TransactOpts, toChainId, targetProxyHash)
}

// BindProxyHash is a paid mutator transaction binding the contract method 0x379b98f6.
//
// Solidity: function bindProxyHash(uint64 toChainId, bytes targetProxyHash) returns(bool)
func (_ILockProxy *ILockProxyTransactorSession) BindProxyHash(toChainId uint64, targetProxyHash []byte) (*types.Transaction, error) {
	return _ILockProxy.Contract.BindProxyHash(&_ILockProxy.TransactOpts, toChainId, targetProxyHash)
}

// Lock is a paid mutator transaction binding the contract method 0x84a6d055.
//
// Solidity: function lock(address fromAssetHash, uint64 toChainId, bytes toAddress, uint256 amount) payable returns(bool)
func (_ILockProxy *ILockProxyTransactor) Lock(opts *bind.TransactOpts, fromAssetHash common.Address, toChainId uint64, toAddress []byte, amount *big.Int) (*types.Transaction, error) {
	return _ILockProxy.contract.Transact(opts, "lock", fromAssetHash, toChainId, toAddress, amount)
}

// Lock is a paid mutator transaction binding the contract method 0x84a6d055.
//
// Solidity: function lock(address fromAssetHash, uint64 toChainId, bytes toAddress, uint256 amount) payable returns(bool)
func (_ILockProxy *ILockProxySession) Lock(fromAssetHash common.Address, toChainId uint64, toAddress []byte, amount *big.Int) (*types.Transaction, error) {
	return _ILockProxy.Contract.Lock(&_ILockProxy.TransactOpts, fromAssetHash, toChainId, toAddress, amount)
}

// Lock is a paid mutator transaction binding the contract method 0x84a6d055.
//
// Solidity: function lock(address fromAssetHash, uint64 toChainId, bytes toAddress, uint256 amount) payable returns(bool)
func (_ILockProxy *ILockProxyTransactorSession) Lock(fromAssetHash common.Address, toChainId uint64, toAddress []byte, amount *big.Int) (*types.Transaction, error) {
	return _ILockProxy.Contract.Lock(&_ILockProxy.TransactOpts, fromAssetHash, toChainId, toAddress, amount)
}

// SetManagerProxy is a paid mutator transaction binding the contract method 0xaf9980f0.
//
// Solidity: function setManagerProxy(address eccmpAddr) returns()
func (_ILockProxy *ILockProxyTransactor) SetManagerProxy(opts *bind.TransactOpts, eccmpAddr common.Address) (*types.Transaction, error) {
	return _ILockProxy.contract.Transact(opts, "setManagerProxy", eccmpAddr)
}

// SetManagerProxy is a paid mutator transaction binding the contract method 0xaf9980f0.
//
// Solidity: function setManagerProxy(address eccmpAddr) returns()
func (_ILockProxy *ILockProxySession) SetManagerProxy(eccmpAddr common.Address) (*types.Transaction, error) {
	return _ILockProxy.Contract.SetManagerProxy(&_ILockProxy.TransactOpts, eccmpAddr)
}

// SetManagerProxy is a paid mutator transaction binding the contract method 0xaf9980f0.
//
// Solidity: function setManagerProxy(address eccmpAddr) returns()
func (_ILockProxy *ILockProxyTransactorSession) SetManagerProxy(eccmpAddr common.Address) (*types.Transaction, error) {
	return _ILockProxy.Contract.SetManagerProxy(&_ILockProxy.TransactOpts, eccmpAddr)
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

// PolyWrapperV3ABI is the input ABI used to generate the binding from.
const PolyWrapperV3ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromAsset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"net\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"PolyWrapperLock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromAsset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"txHash\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"efee\",\"type\":\"uint256\"}],\"name\":\"PolyWrapperSpeedUp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lockProxy\",\"type\":\"address\"}],\"name\":\"addLockProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"extractFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"richGuy\",\"type\":\"address\"}],\"name\":\"getBalanceBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"lockProxy\",\"type\":\"address\"}],\"name\":\"isValidLockProxy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAsset\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lockProxyIndexMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxLockProxyIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_lockProxy\",\"type\":\"address\"}],\"name\":\"resetLockProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"collector\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAsset\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"lockProxy\",\"type\":\"address\"}],\"name\":\"specifiedLock\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAsset\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"txHash\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"speedUp\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PolyWrapperV3FuncSigs maps the 4-byte function signature to its string representation.
var PolyWrapperV3FuncSigs = map[string]string{
	"9d8c0ce2": "addLockProxy(address)",
	"9a8a0592": "chainId()",
	"1745399d": "extractFee(address)",
	"c415b95c": "feeCollector()",
	"13f0b82b": "getBalanceBatch(address[],address)",
	"8f32d59b": "isOwner()",
	"665fd2bd": "isValidLockProxy(address)",
	"60de1a9b": "lock(address,uint64,bytes,uint256,uint256,uint256)",
	"0df9cfc5": "lockProxyIndexMap(uint256)",
	"8b41cad3": "maxLockProxyIndex()",
	"8da5cb5b": "owner()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"715018a6": "renounceOwnership()",
	"87aafc94": "resetLockProxy(uint256,address)",
	"a42dce80": "setFeeCollector(address)",
	"b2fe4ed6": "specifiedLock(address,uint64,bytes,uint256,uint256,uint256,address)",
	"d3ed7c76": "speedUp(address,bytes,uint256)",
	"f2fde38b": "transferOwnership(address)",
	"3f4ba83a": "unpause()",
}

// PolyWrapperV3Bin is the compiled bytecode used for deploying new contracts.
var PolyWrapperV3Bin = "0x608060405260006001553480156200001657600080fd5b50604051620025f3380380620025f3833981810160405260408110156200003c57600080fd5b50805160209091015160006200005a6001600160e01b036200010116565b600080546001600160a01b0319166001600160a01b038316908117825560405192935091600080516020620025d3833981519152908290a3506000805461ffff60a01b1916600160a81b17905580620000e3576040805162461bcd60e51b8152602060048201526006602482015265085b1959d85b60d21b604482015290519081900360640190fd5b620000f7826001600160e01b036200010516565b6002555062000241565b3390565b620001186001600160e01b036200018116565b6200016a576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6200017e816001600160e01b03620001b016565b50565b600080546001600160a01b0316620001a16001600160e01b036200010116565b6001600160a01b031614905090565b6001600160a01b038116620001f75760405162461bcd60e51b8152600401808060200182810382526026815260200180620025ad6026913960400191505060405180910390fd5b600080546040516001600160a01b0380851693921691600080516020620025d383398151915291a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b61235c80620002516000396000f3fe60806040526004361061011f5760003560e01c80638b41cad3116100a0578063a42dce8011610064578063a42dce80146104e3578063b2fe4ed614610516578063c415b95c146105f3578063d3ed7c7614610608578063f2fde38b146106be5761011f565b80638b41cad31461044a5780638da5cb5b146104715780638f32d59b146104865780639a8a05921461049b5780639d8c0ce2146104b05761011f565b806360de1a9b116100e757806360de1a9b146102e6578063665fd2bd146103b4578063715018a6146103e75780638456cb59146103fc57806387aafc94146104115761011f565b80630df9cfc51461012457806313f0b82b1461016a5780631745399d146102735780633f4ba83a146102a85780635c975abb146102bd575b600080fd5b34801561013057600080fd5b5061014e6004803603602081101561014757600080fd5b50356106f1565b604080516001600160a01b039092168252519081900360200190f35b34801561017657600080fd5b506102236004803603604081101561018d57600080fd5b810190602081018135600160201b8111156101a757600080fd5b8201836020820111156101b957600080fd5b803590602001918460208302840111600160201b831117156101da57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550505090356001600160a01b0316915061070c9050565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561025f578181015183820152602001610247565b505050509050019250505060405180910390f35b34801561027f57600080fd5b506102a66004803603602081101561029657600080fd5b50356001600160a01b0316610861565b005b3480156102b457600080fd5b506102a6610989565b3480156102c957600080fd5b506102d26109da565b604080519115158252519081900360200190f35b6102a6600480360360c08110156102fc57600080fd5b6001600160a01b038235169167ffffffffffffffff60208201351691810190606081016040820135600160201b81111561033557600080fd5b82018360208201111561034757600080fd5b803590602001918460018302840111600160201b8311171561036857600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050823593505050602081013590604001356109ea565b3480156103c057600080fd5b506102d2600480360360208110156103d757600080fd5b50356001600160a01b0316610ccf565b3480156103f357600080fd5b506102a6610d1c565b34801561040857600080fd5b506102a6610dad565b34801561041d57600080fd5b506102a66004803603604081101561043457600080fd5b50803590602001356001600160a01b0316610dfc565b34801561045657600080fd5b5061045f610f84565b60408051918252519081900360200190f35b34801561047d57600080fd5b5061014e610f8a565b34801561049257600080fd5b506102d2610f99565b3480156104a757600080fd5b5061045f610fbd565b3480156104bc57600080fd5b506102a6600480360360208110156104d357600080fd5b50356001600160a01b0316610fc3565b3480156104ef57600080fd5b506102a66004803603602081101561050657600080fd5b50356001600160a01b03166110fe565b6102a6600480360360e081101561052c57600080fd5b6001600160a01b038235169167ffffffffffffffff60208201351691810190606081016040820135600160201b81111561056557600080fd5b82018360208201111561057757600080fd5b803590602001918460018302840111600160201b8311171561059857600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050823593505050602081013590604081013590606001356001600160a01b03166111b2565b3480156105ff57600080fd5b5061014e6113d9565b6102a66004803603606081101561061e57600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561064857600080fd5b82018360208201111561065a57600080fd5b803590602001918460018302840111600160201b8311171561067b57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506113e8915050565b3480156106ca57600080fd5b506102a6600480360360208110156106e157600080fd5b50356001600160a01b0316611565565b6004602052600090815260409020546001600160a01b031681565b606080835160405190808252806020026020018201604052801561073a578160200160208202803883390190505b50905060005b84518110156108575760006001600160a01b031685828151811061076057fe5b60200260200101516001600160a01b0316141561079f57836001600160a01b03163182828151811061078e57fe5b60200260200101818152505061084f565b8481815181106107ab57fe5b60200260200101516001600160a01b03166370a08231856040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561080857600080fd5b505afa15801561081c573d6000803e3d6000fd5b505050506040513d602081101561083257600080fd5b5051825183908390811061084257fe5b6020026020010181815250505b600101610740565b5090505b92915050565b6003546001600160a01b031633146108b0576040805162461bcd60e51b815260206004820152600d60248201526c10b332b2a1b7b63632b1ba37b960991b604482015290519081900360640190fd5b6001600160a01b0381166108f05760405133904780156108fc02916000818181858888f193505050501580156108ea573d6000803e3d6000fd5b50610986565b600354604080516370a0823160e01b81523060048201529051610986926001600160a01b0390811692908516916370a0823191602480820192602092909190829003018186803b15801561094357600080fd5b505afa158015610957573d6000803e3d6000fd5b505050506040513d602081101561096d57600080fd5b50516001600160a01b038416919063ffffffff6115b516565b50565b610991610f99565b6109d0576040805162461bcd60e51b81526020600482018190526024820152600080516020612285833981519152604482015290519081900360640190fd5b6109d861160c565b565b600054600160a01b900460ff1690565b600054600160a81b900460ff16610a48576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b6000805460ff60a81b191690819055600160a01b900460ff1615610aa6576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b6002548567ffffffffffffffff1614158015610acb575067ffffffffffffffff851615155b610b09576040805162461bcd60e51b815260206004820152600a602482015269085d1bd0da185a5b925960b21b604482015290519081900360640190fd5b8351610b4e576040805162461bcd60e51b815260206004820152600f60248201526e656d70747920746f4164647265737360881b604482015290519081900360640190fd5b60148401516001600160a01b038116610b9f576040805162461bcd60e51b815260206004820152600e60248201526d7a65726f20746f4164647265737360901b604482015290519081900360640190fd5b610ba987856116b4565b610bb487858561172b565b93506000610bc2888861183b565b9050610bd188888888856119ef565b336001600160a01b0316886001600160a01b03167f2b0591052cc6602e870d3994f0a1b173fdac98c215cb3b0baf84eaca5a0aa81e8989898989604051808667ffffffffffffffff1667ffffffffffffffff16815260200180602001858152602001848152602001838152602001828103825286818151815260200191508051906020019080838360005b83811015610c74578181015183820152602001610c5c565b50505050905090810190601f168015610ca15780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a350506000805460ff60a81b1916600160a81b179055505050505050565b6000805b600154811015610d11576000818152600460205260409020546001600160a01b0384811691161415610d09576001915050610d17565b600101610cd3565b50600090505b919050565b610d24610f99565b610d63576040805162461bcd60e51b81526020600482018190526024820152600080516020612285833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b610db5610f99565b610df4576040805162461bcd60e51b81526020600482018190526024820152600080516020612285833981519152604482015290519081900360640190fd5b6109d8611cd7565b610e04610f99565b610e43576040805162461bcd60e51b81526020600482018190526024820152600080516020612285833981519152604482015290519081900360640190fd5b6001600160a01b038116610e5657600080fd5b6000828152600460205260409020546001600160a01b0316610ea95760405162461bcd60e51b81526004018080602001828103825260228152602001806122636022913960400191505060405180910390fd5b600082815260046020818152604080842080546001600160a01b0319166001600160a01b038716908117909155815163d798f88160e01b81529151909363d798f881938382019390929190829003018186803b158015610f0857600080fd5b505afa158015610f1c573d6000803e3d6000fd5b505050506040513d6020811015610f3257600080fd5b50516001600160a01b03161415610f80576040805162461bcd60e51b815260206004820152600d60248201526c6e6f74206c6f636b70726f787960981b604482015290519081900360640190fd5b5050565b60015481565b6000546001600160a01b031690565b600080546001600160a01b0316610fae611d61565b6001600160a01b031614905090565b60025481565b610fcb610f99565b61100a576040805162461bcd60e51b81526020600482018190526024820152600080516020612285833981519152604482015290519081900360640190fd5b6001600160a01b03811661101d57600080fd5b60018054808201909155600090815260046020818152604080842080546001600160a01b0319166001600160a01b038716908117909155815163d798f88160e01b81529151909363d798f881938382019390929190829003018186803b15801561108657600080fd5b505afa15801561109a573d6000803e3d6000fd5b505050506040513d60208110156110b057600080fd5b50516001600160a01b03161415610986576040805162461bcd60e51b815260206004820152600d60248201526c6e6f74206c6f636b70726f787960981b604482015290519081900360640190fd5b611106610f99565b611145576040805162461bcd60e51b81526020600482018190526024820152600080516020612285833981519152604482015290519081900360640190fd5b6001600160a01b038116611190576040805162461bcd60e51b815260206004820152600d60248201526c656d747079206164647265737360981b604482015290519081900360640190fd5b600380546001600160a01b0319166001600160a01b0392909216919091179055565b600054600160a81b900460ff16611210576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b6000805460ff60a81b191690819055600160a01b900460ff161561126e576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b6002548667ffffffffffffffff1614158015611293575067ffffffffffffffff861615155b6112d1576040805162461bcd60e51b815260206004820152600a602482015269085d1bd0da185a5b925960b21b604482015290519081900360640190fd5b8451611316576040805162461bcd60e51b815260206004820152600f60248201526e656d70747920746f4164647265737360881b604482015290519081900360640190fd5b60148501516001600160a01b038116611367576040805162461bcd60e51b815260206004820152600e60248201526d7a65726f20746f4164647265737360901b604482015290519081900360640190fd5b61137188866116b4565b61137c88868661172b565b945061138782610ccf565b6113cc576040805162461bcd60e51b8152602060048201526011602482015270696e76616c6964206c6f636b50726f787960781b604482015290519081900360640190fd5b610bd188888888866119ef565b6003546001600160a01b031681565b600054600160a81b900460ff16611446576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b6000805460ff60a81b191690819055600160a01b900460ff16156114a4576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b6114ae83826116b4565b336001600160a01b0316826040518082805190602001908083835b602083106114e85780518252601f1990920191602091820191016114c9565b51815160209384036101000a60001901801990921691161790526040805192909401829003822088835293519395506001600160a01b038a1694507ff6579aef3e0d086d986c5d6972659f8a0d8602ef7945b054be1b88e088773ef69391829003019150a450506000805460ff60a81b1916600160a81b17905550565b61156d610f99565b6115ac576040805162461bcd60e51b81526020600482018190526024820152600080516020612285833981519152604482015290519081900360640190fd5b61098681611d65565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b179052611607908490611e05565b505050565b600054600160a01b900460ff16611661576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa611697611d61565b604080516001600160a01b039092168252519081900360200190a1565b6001600160a01b0382166117105780341461170b576040805162461bcd60e51b815260206004820152601260248201527134b739bab33334b1b4b2b73a1032ba3432b960711b604482015290519081900360640190fd5b610f80565b610f806001600160a01b03831633308463ffffffff611fba16565b60006001600160a01b0384166117e75782341015611785576040805162461bcd60e51b815260206004820152601260248201527134b739bab33334b1b4b2b73a1032ba3432b960711b604482015290519081900360640190fd5b8183116117d0576040805162461bcd60e51b8152602060048201526014602482015273616d6f756e74206c657373207468616e2066656560601b604482015290519081900360640190fd5b6117e0838363ffffffff61201416565b9050611834565b81341015611831576040805162461bcd60e51b815260206004820152601260248201527134b739bab33334b1b4b2b73a1032ba3432b960711b604482015290519081900360640190fd5b50815b9392505050565b6000805b6001548110156119b75760008181526004602081905260408083205481516309efb30160e31b81526001600160a01b038981169482019490945267ffffffffffffffff8816602482015291519216928392634f7d9808926044808201939291829003018186803b1580156118b257600080fd5b505afa1580156118c6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156118ef57600080fd5b8101908080516040519392919084600160201b82111561190e57600080fd5b90830190602082018581111561192357600080fd5b8251600160201b81118282018810171561193c57600080fd5b82525081516020918201929091019080838360005b83811015611969578181015183820152602001611951565b50505050905090810190601f1680156119965780820380516001836020036101000a031916815260200191505b50604052505050516000146119ae57915061085b9050565b5060010161183f565b5060405162461bcd60e51b81526004018080602001828103825260238152602001806123056023913960400191505060405180910390fd5b806001600160a01b038616611b5957806001600160a01b03166384a6d05584888888886040518663ffffffff1660e01b815260040180856001600160a01b03166001600160a01b031681526020018467ffffffffffffffff1667ffffffffffffffff16815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b83811015611a96578181015183820152602001611a7e565b50505050905090810190601f168015611ac35780820380516001836020036101000a031916815260200191505b50955050505050506020604051808303818588803b158015611ae457600080fd5b505af1158015611af8573d6000803e3d6000fd5b50505050506040513d6020811015611b0f57600080fd5b5051611b54576040805162461bcd60e51b815260206004820152600f60248201526e1b1bd8dac8195d1a195c8819985a5b608a1b604482015290519081900360640190fd5b611ccf565b611b746001600160a01b03871682600063ffffffff61205616565b611b8e6001600160a01b038716828563ffffffff61205616565b6040516384a6d05560e01b81526001600160a01b038781166004830190815267ffffffffffffffff8816602484015260648301869052608060448401908152875160848501528751928516936384a6d055938b938b938b938b9360a490910190602086019080838360005b83811015611c11578181015183820152602001611bf9565b50505050905090810190601f168015611c3e5780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b158015611c6057600080fd5b505af1158015611c74573d6000803e3d6000fd5b505050506040513d6020811015611c8a57600080fd5b5051611ccf576040805162461bcd60e51b815260206004820152600f60248201526e1b1bd8dac8195c98cc8c0819985a5b608a1b604482015290519081900360640190fd5b505050505050565b600054600160a01b900460ff1615611d29576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586116975b3390565b6001600160a01b038116611daa5760405162461bcd60e51b815260040180806020018281038252602681526020018061223d6026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b611e0e82612169565b611e5f576040805162461bcd60e51b815260206004820152601f60248201527f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e747261637400604482015290519081900360640190fd5b60006060836001600160a01b0316836040518082805190602001908083835b60208310611e9d5780518252601f199092019160209182019101611e7e565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114611eff576040519150601f19603f3d011682016040523d82523d6000602084013e611f04565b606091505b509150915081611f5b576040805162461bcd60e51b815260206004820181905260248201527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564604482015290519081900360640190fd5b805115611fb457808060200190516020811015611f7757600080fd5b5051611fb45760405162461bcd60e51b815260040180806020018281038252602a8152602001806122a5602a913960400191505060405180910390fd5b50505050565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b179052611fb4908590611e05565b600061183483836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506121a5565b8015806120dc575060408051636eb1769f60e11b81523060048201526001600160a01b03848116602483015291519185169163dd62ed3e91604480820192602092909190829003018186803b1580156120ae57600080fd5b505afa1580156120c2573d6000803e3d6000fd5b505050506040513d60208110156120d857600080fd5b5051155b6121175760405162461bcd60e51b81526004018080602001828103825260368152602001806122cf6036913960400191505060405180910390fd5b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b179052611607908490611e05565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470811580159061219d5750808214155b949350505050565b600081848411156122345760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156121f95781810151838201526020016121e1565b50505050905090810190601f1680156122265780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50505090039056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573736e6f206c6f636b70726f78792065787369737420696e20676976656e20696e6465784f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725361666545524332303a204552433230206f7065726174696f6e20646964206e6f7420737563636565645361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f20746f206e6f6e2d7a65726f20616c6c6f77616e63654e6f204c6f636b50726f787920537570706f727420746869732063726f73732074786ea265627a7a72315820c944ac2c184080e4d5b6254e8ade1bf4e4dfd8e372ee4bc155a723884b6c003e64736f6c634300051100324f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573738be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0"

// DeployPolyWrapperV3 deploys a new Ethereum contract, binding an instance of PolyWrapperV3 to it.
func DeployPolyWrapperV3(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _chainId *big.Int) (common.Address, *types.Transaction, *PolyWrapperV3, error) {
	parsed, err := abi.JSON(strings.NewReader(PolyWrapperV3ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PolyWrapperV3Bin), backend, _owner, _chainId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PolyWrapperV3{PolyWrapperV3Caller: PolyWrapperV3Caller{contract: contract}, PolyWrapperV3Transactor: PolyWrapperV3Transactor{contract: contract}, PolyWrapperV3Filterer: PolyWrapperV3Filterer{contract: contract}}, nil
}

// PolyWrapperV3 is an auto generated Go binding around an Ethereum contract.
type PolyWrapperV3 struct {
	PolyWrapperV3Caller     // Read-only binding to the contract
	PolyWrapperV3Transactor // Write-only binding to the contract
	PolyWrapperV3Filterer   // Log filterer for contract events
}

// PolyWrapperV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type PolyWrapperV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolyWrapperV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type PolyWrapperV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolyWrapperV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolyWrapperV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolyWrapperV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolyWrapperV3Session struct {
	Contract     *PolyWrapperV3    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PolyWrapperV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolyWrapperV3CallerSession struct {
	Contract *PolyWrapperV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PolyWrapperV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolyWrapperV3TransactorSession struct {
	Contract     *PolyWrapperV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PolyWrapperV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type PolyWrapperV3Raw struct {
	Contract *PolyWrapperV3 // Generic contract binding to access the raw methods on
}

// PolyWrapperV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolyWrapperV3CallerRaw struct {
	Contract *PolyWrapperV3Caller // Generic read-only contract binding to access the raw methods on
}

// PolyWrapperV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolyWrapperV3TransactorRaw struct {
	Contract *PolyWrapperV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPolyWrapperV3 creates a new instance of PolyWrapperV3, bound to a specific deployed contract.
func NewPolyWrapperV3(address common.Address, backend bind.ContractBackend) (*PolyWrapperV3, error) {
	contract, err := bindPolyWrapperV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PolyWrapperV3{PolyWrapperV3Caller: PolyWrapperV3Caller{contract: contract}, PolyWrapperV3Transactor: PolyWrapperV3Transactor{contract: contract}, PolyWrapperV3Filterer: PolyWrapperV3Filterer{contract: contract}}, nil
}

// NewPolyWrapperV3Caller creates a new read-only instance of PolyWrapperV3, bound to a specific deployed contract.
func NewPolyWrapperV3Caller(address common.Address, caller bind.ContractCaller) (*PolyWrapperV3Caller, error) {
	contract, err := bindPolyWrapperV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolyWrapperV3Caller{contract: contract}, nil
}

// NewPolyWrapperV3Transactor creates a new write-only instance of PolyWrapperV3, bound to a specific deployed contract.
func NewPolyWrapperV3Transactor(address common.Address, transactor bind.ContractTransactor) (*PolyWrapperV3Transactor, error) {
	contract, err := bindPolyWrapperV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolyWrapperV3Transactor{contract: contract}, nil
}

// NewPolyWrapperV3Filterer creates a new log filterer instance of PolyWrapperV3, bound to a specific deployed contract.
func NewPolyWrapperV3Filterer(address common.Address, filterer bind.ContractFilterer) (*PolyWrapperV3Filterer, error) {
	contract, err := bindPolyWrapperV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolyWrapperV3Filterer{contract: contract}, nil
}

// bindPolyWrapperV3 binds a generic wrapper to an already deployed contract.
func bindPolyWrapperV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PolyWrapperV3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PolyWrapperV3 *PolyWrapperV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PolyWrapperV3.Contract.PolyWrapperV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PolyWrapperV3 *PolyWrapperV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolyWrapperV3.Contract.PolyWrapperV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PolyWrapperV3 *PolyWrapperV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PolyWrapperV3.Contract.PolyWrapperV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PolyWrapperV3 *PolyWrapperV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PolyWrapperV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PolyWrapperV3 *PolyWrapperV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolyWrapperV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PolyWrapperV3 *PolyWrapperV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PolyWrapperV3.Contract.contract.Transact(opts, method, params...)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_PolyWrapperV3 *PolyWrapperV3Caller) ChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolyWrapperV3.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_PolyWrapperV3 *PolyWrapperV3Session) ChainId() (*big.Int, error) {
	return _PolyWrapperV3.Contract.ChainId(&_PolyWrapperV3.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_PolyWrapperV3 *PolyWrapperV3CallerSession) ChainId() (*big.Int, error) {
	return _PolyWrapperV3.Contract.ChainId(&_PolyWrapperV3.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_PolyWrapperV3 *PolyWrapperV3Caller) FeeCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolyWrapperV3.contract.Call(opts, &out, "feeCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_PolyWrapperV3 *PolyWrapperV3Session) FeeCollector() (common.Address, error) {
	return _PolyWrapperV3.Contract.FeeCollector(&_PolyWrapperV3.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_PolyWrapperV3 *PolyWrapperV3CallerSession) FeeCollector() (common.Address, error) {
	return _PolyWrapperV3.Contract.FeeCollector(&_PolyWrapperV3.CallOpts)
}

// GetBalanceBatch is a free data retrieval call binding the contract method 0x13f0b82b.
//
// Solidity: function getBalanceBatch(address[] assets, address richGuy) view returns(uint256[])
func (_PolyWrapperV3 *PolyWrapperV3Caller) GetBalanceBatch(opts *bind.CallOpts, assets []common.Address, richGuy common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _PolyWrapperV3.contract.Call(opts, &out, "getBalanceBatch", assets, richGuy)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetBalanceBatch is a free data retrieval call binding the contract method 0x13f0b82b.
//
// Solidity: function getBalanceBatch(address[] assets, address richGuy) view returns(uint256[])
func (_PolyWrapperV3 *PolyWrapperV3Session) GetBalanceBatch(assets []common.Address, richGuy common.Address) ([]*big.Int, error) {
	return _PolyWrapperV3.Contract.GetBalanceBatch(&_PolyWrapperV3.CallOpts, assets, richGuy)
}

// GetBalanceBatch is a free data retrieval call binding the contract method 0x13f0b82b.
//
// Solidity: function getBalanceBatch(address[] assets, address richGuy) view returns(uint256[])
func (_PolyWrapperV3 *PolyWrapperV3CallerSession) GetBalanceBatch(assets []common.Address, richGuy common.Address) ([]*big.Int, error) {
	return _PolyWrapperV3.Contract.GetBalanceBatch(&_PolyWrapperV3.CallOpts, assets, richGuy)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_PolyWrapperV3 *PolyWrapperV3Caller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PolyWrapperV3.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_PolyWrapperV3 *PolyWrapperV3Session) IsOwner() (bool, error) {
	return _PolyWrapperV3.Contract.IsOwner(&_PolyWrapperV3.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_PolyWrapperV3 *PolyWrapperV3CallerSession) IsOwner() (bool, error) {
	return _PolyWrapperV3.Contract.IsOwner(&_PolyWrapperV3.CallOpts)
}

// IsValidLockProxy is a free data retrieval call binding the contract method 0x665fd2bd.
//
// Solidity: function isValidLockProxy(address lockProxy) view returns(bool)
func (_PolyWrapperV3 *PolyWrapperV3Caller) IsValidLockProxy(opts *bind.CallOpts, lockProxy common.Address) (bool, error) {
	var out []interface{}
	err := _PolyWrapperV3.contract.Call(opts, &out, "isValidLockProxy", lockProxy)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidLockProxy is a free data retrieval call binding the contract method 0x665fd2bd.
//
// Solidity: function isValidLockProxy(address lockProxy) view returns(bool)
func (_PolyWrapperV3 *PolyWrapperV3Session) IsValidLockProxy(lockProxy common.Address) (bool, error) {
	return _PolyWrapperV3.Contract.IsValidLockProxy(&_PolyWrapperV3.CallOpts, lockProxy)
}

// IsValidLockProxy is a free data retrieval call binding the contract method 0x665fd2bd.
//
// Solidity: function isValidLockProxy(address lockProxy) view returns(bool)
func (_PolyWrapperV3 *PolyWrapperV3CallerSession) IsValidLockProxy(lockProxy common.Address) (bool, error) {
	return _PolyWrapperV3.Contract.IsValidLockProxy(&_PolyWrapperV3.CallOpts, lockProxy)
}

// LockProxyIndexMap is a free data retrieval call binding the contract method 0x0df9cfc5.
//
// Solidity: function lockProxyIndexMap(uint256 ) view returns(address)
func (_PolyWrapperV3 *PolyWrapperV3Caller) LockProxyIndexMap(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PolyWrapperV3.contract.Call(opts, &out, "lockProxyIndexMap", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LockProxyIndexMap is a free data retrieval call binding the contract method 0x0df9cfc5.
//
// Solidity: function lockProxyIndexMap(uint256 ) view returns(address)
func (_PolyWrapperV3 *PolyWrapperV3Session) LockProxyIndexMap(arg0 *big.Int) (common.Address, error) {
	return _PolyWrapperV3.Contract.LockProxyIndexMap(&_PolyWrapperV3.CallOpts, arg0)
}

// LockProxyIndexMap is a free data retrieval call binding the contract method 0x0df9cfc5.
//
// Solidity: function lockProxyIndexMap(uint256 ) view returns(address)
func (_PolyWrapperV3 *PolyWrapperV3CallerSession) LockProxyIndexMap(arg0 *big.Int) (common.Address, error) {
	return _PolyWrapperV3.Contract.LockProxyIndexMap(&_PolyWrapperV3.CallOpts, arg0)
}

// MaxLockProxyIndex is a free data retrieval call binding the contract method 0x8b41cad3.
//
// Solidity: function maxLockProxyIndex() view returns(uint256)
func (_PolyWrapperV3 *PolyWrapperV3Caller) MaxLockProxyIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolyWrapperV3.contract.Call(opts, &out, "maxLockProxyIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLockProxyIndex is a free data retrieval call binding the contract method 0x8b41cad3.
//
// Solidity: function maxLockProxyIndex() view returns(uint256)
func (_PolyWrapperV3 *PolyWrapperV3Session) MaxLockProxyIndex() (*big.Int, error) {
	return _PolyWrapperV3.Contract.MaxLockProxyIndex(&_PolyWrapperV3.CallOpts)
}

// MaxLockProxyIndex is a free data retrieval call binding the contract method 0x8b41cad3.
//
// Solidity: function maxLockProxyIndex() view returns(uint256)
func (_PolyWrapperV3 *PolyWrapperV3CallerSession) MaxLockProxyIndex() (*big.Int, error) {
	return _PolyWrapperV3.Contract.MaxLockProxyIndex(&_PolyWrapperV3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PolyWrapperV3 *PolyWrapperV3Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolyWrapperV3.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PolyWrapperV3 *PolyWrapperV3Session) Owner() (common.Address, error) {
	return _PolyWrapperV3.Contract.Owner(&_PolyWrapperV3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PolyWrapperV3 *PolyWrapperV3CallerSession) Owner() (common.Address, error) {
	return _PolyWrapperV3.Contract.Owner(&_PolyWrapperV3.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PolyWrapperV3 *PolyWrapperV3Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PolyWrapperV3.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PolyWrapperV3 *PolyWrapperV3Session) Paused() (bool, error) {
	return _PolyWrapperV3.Contract.Paused(&_PolyWrapperV3.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PolyWrapperV3 *PolyWrapperV3CallerSession) Paused() (bool, error) {
	return _PolyWrapperV3.Contract.Paused(&_PolyWrapperV3.CallOpts)
}

// AddLockProxy is a paid mutator transaction binding the contract method 0x9d8c0ce2.
//
// Solidity: function addLockProxy(address _lockProxy) returns()
func (_PolyWrapperV3 *PolyWrapperV3Transactor) AddLockProxy(opts *bind.TransactOpts, _lockProxy common.Address) (*types.Transaction, error) {
	return _PolyWrapperV3.contract.Transact(opts, "addLockProxy", _lockProxy)
}

// AddLockProxy is a paid mutator transaction binding the contract method 0x9d8c0ce2.
//
// Solidity: function addLockProxy(address _lockProxy) returns()
func (_PolyWrapperV3 *PolyWrapperV3Session) AddLockProxy(_lockProxy common.Address) (*types.Transaction, error) {
	return _PolyWrapperV3.Contract.AddLockProxy(&_PolyWrapperV3.TransactOpts, _lockProxy)
}

// AddLockProxy is a paid mutator transaction binding the contract method 0x9d8c0ce2.
//
// Solidity: function addLockProxy(address _lockProxy) returns()
func (_PolyWrapperV3 *PolyWrapperV3TransactorSession) AddLockProxy(_lockProxy common.Address) (*types.Transaction, error) {
	return _PolyWrapperV3.Contract.AddLockProxy(&_PolyWrapperV3.TransactOpts, _lockProxy)
}

// ExtractFee is a paid mutator transaction binding the contract method 0x1745399d.
//
// Solidity: function extractFee(address token) returns()
func (_PolyWrapperV3 *PolyWrapperV3Transactor) ExtractFee(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _PolyWrapperV3.contract.Transact(opts, "extractFee", token)
}

// ExtractFee is a paid mutator transaction binding the contract method 0x1745399d.
//
// Solidity: function extractFee(address token) returns()
func (_PolyWrapperV3 *PolyWrapperV3Session) ExtractFee(token common.Address) (*types.Transaction, error) {
	return _PolyWrapperV3.Contract.ExtractFee(&_PolyWrapperV3.TransactOpts, token)
}

// ExtractFee is a paid mutator transaction binding the contract method 0x1745399d.
//
// Solidity: function extractFee(address token) returns()
func (_PolyWrapperV3 *PolyWrapperV3TransactorSession) ExtractFee(token common.Address) (*types.Transaction, error) {
	return _PolyWrapperV3.Contract.ExtractFee(&_PolyWrapperV3.TransactOpts, token)
}

// Lock is a paid mutator transaction binding the contract method 0x60de1a9b.
//
// Solidity: function lock(address fromAsset, uint64 toChainId, bytes toAddress, uint256 amount, uint256 fee, uint256 id) payable returns()
func (_PolyWrapperV3 *PolyWrapperV3Transactor) Lock(opts *bind.TransactOpts, fromAsset common.Address, toChainId uint64, toAddress []byte, amount *big.Int, fee *big.Int, id *big.Int) (*types.Transaction, error) {
	return _PolyWrapperV3.contract.Transact(opts, "lock", fromAsset, toChainId, toAddress, amount, fee, id)
}

// Lock is a paid mutator transaction binding the contract method 0x60de1a9b.
//
// Solidity: function lock(address fromAsset, uint64 toChainId, bytes toAddress, uint256 amount, uint256 fee, uint256 id) payable returns()
func (_PolyWrapperV3 *PolyWrapperV3Session) Lock(fromAsset common.Address, toChainId uint64, toAddress []byte, amount *big.Int, fee *big.Int, id *big.Int) (*types.Transaction, error) {
	return _PolyWrapperV3.Contract.Lock(&_PolyWrapperV3.TransactOpts, fromAsset, toChainId, toAddress, amount, fee, id)
}

// Lock is a paid mutator transaction binding the contract method 0x60de1a9b.
//
// Solidity: function lock(address fromAsset, uint64 toChainId, bytes toAddress, uint256 amount, uint256 fee, uint256 id) payable returns()
func (_PolyWrapperV3 *PolyWrapperV3TransactorSession) Lock(fromAsset common.Address, toChainId uint64, toAddress []byte, amount *big.Int, fee *big.Int, id *big.Int) (*types.Transaction, error) {
	return _PolyWrapperV3.Contract.Lock(&_PolyWrapperV3.TransactOpts, fromAsset, toChainId, toAddress, amount, fee, id)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PolyWrapperV3 *PolyWrapperV3Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolyWrapperV3.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PolyWrapperV3 *PolyWrapperV3Session) Pause() (*types.Transaction, error) {
	return _PolyWrapperV3.Contract.Pause(&_PolyWrapperV3.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PolyWrapperV3 *PolyWrapperV3TransactorSession) Pause() (*types.Transaction, error) {
	return _PolyWrapperV3.Contract.Pause(&_PolyWrapperV3.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PolyWrapperV3 *PolyWrapperV3Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolyWrapperV3.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PolyWrapperV3 *PolyWrapperV3Session) RenounceOwnership() (*types.Transaction, error) {
	return _PolyWrapperV3.Contract.RenounceOwnership(&_PolyWrapperV3.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PolyWrapperV3 *PolyWrapperV3TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PolyWrapperV3.Contract.RenounceOwnership(&_PolyWrapperV3.TransactOpts)
}

// ResetLockProxy is a paid mutator transaction binding the contract method 0x87aafc94.
//
// Solidity: function resetLockProxy(uint256 index, address _lockProxy) returns()
func (_PolyWrapperV3 *PolyWrapperV3Transactor) ResetLockProxy(opts *bind.TransactOpts, index *big.Int, _lockProxy common.Address) (*types.Transaction, error) {
	return _PolyWrapperV3.contract.Transact(opts, "resetLockProxy", index, _lockProxy)
}

// ResetLockProxy is a paid mutator transaction binding the contract method 0x87aafc94.
//
// Solidity: function resetLockProxy(uint256 index, address _lockProxy) returns()
func (_PolyWrapperV3 *PolyWrapperV3Session) ResetLockProxy(index *big.Int, _lockProxy common.Address) (*types.Transaction, error) {
	return _PolyWrapperV3.Contract.ResetLockProxy(&_PolyWrapperV3.TransactOpts, index, _lockProxy)
}

// ResetLockProxy is a paid mutator transaction binding the contract method 0x87aafc94.
//
// Solidity: function resetLockProxy(uint256 index, address _lockProxy) returns()
func (_PolyWrapperV3 *PolyWrapperV3TransactorSession) ResetLockProxy(index *big.Int, _lockProxy common.Address) (*types.Transaction, error) {
	return _PolyWrapperV3.Contract.ResetLockProxy(&_PolyWrapperV3.TransactOpts, index, _lockProxy)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address collector) returns()
func (_PolyWrapperV3 *PolyWrapperV3Transactor) SetFeeCollector(opts *bind.TransactOpts, collector common.Address) (*types.Transaction, error) {
	return _PolyWrapperV3.contract.Transact(opts, "setFeeCollector", collector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address collector) returns()
func (_PolyWrapperV3 *PolyWrapperV3Session) SetFeeCollector(collector common.Address) (*types.Transaction, error) {
	return _PolyWrapperV3.Contract.SetFeeCollector(&_PolyWrapperV3.TransactOpts, collector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address collector) returns()
func (_PolyWrapperV3 *PolyWrapperV3TransactorSession) SetFeeCollector(collector common.Address) (*types.Transaction, error) {
	return _PolyWrapperV3.Contract.SetFeeCollector(&_PolyWrapperV3.TransactOpts, collector)
}

// SpecifiedLock is a paid mutator transaction binding the contract method 0xb2fe4ed6.
//
// Solidity: function specifiedLock(address fromAsset, uint64 toChainId, bytes toAddress, uint256 amount, uint256 fee, uint256 id, address lockProxy) payable returns()
func (_PolyWrapperV3 *PolyWrapperV3Transactor) SpecifiedLock(opts *bind.TransactOpts, fromAsset common.Address, toChainId uint64, toAddress []byte, amount *big.Int, fee *big.Int, id *big.Int, lockProxy common.Address) (*types.Transaction, error) {
	return _PolyWrapperV3.contract.Transact(opts, "specifiedLock", fromAsset, toChainId, toAddress, amount, fee, id, lockProxy)
}

// SpecifiedLock is a paid mutator transaction binding the contract method 0xb2fe4ed6.
//
// Solidity: function specifiedLock(address fromAsset, uint64 toChainId, bytes toAddress, uint256 amount, uint256 fee, uint256 id, address lockProxy) payable returns()
func (_PolyWrapperV3 *PolyWrapperV3Session) SpecifiedLock(fromAsset common.Address, toChainId uint64, toAddress []byte, amount *big.Int, fee *big.Int, id *big.Int, lockProxy common.Address) (*types.Transaction, error) {
	return _PolyWrapperV3.Contract.SpecifiedLock(&_PolyWrapperV3.TransactOpts, fromAsset, toChainId, toAddress, amount, fee, id, lockProxy)
}

// SpecifiedLock is a paid mutator transaction binding the contract method 0xb2fe4ed6.
//
// Solidity: function specifiedLock(address fromAsset, uint64 toChainId, bytes toAddress, uint256 amount, uint256 fee, uint256 id, address lockProxy) payable returns()
func (_PolyWrapperV3 *PolyWrapperV3TransactorSession) SpecifiedLock(fromAsset common.Address, toChainId uint64, toAddress []byte, amount *big.Int, fee *big.Int, id *big.Int, lockProxy common.Address) (*types.Transaction, error) {
	return _PolyWrapperV3.Contract.SpecifiedLock(&_PolyWrapperV3.TransactOpts, fromAsset, toChainId, toAddress, amount, fee, id, lockProxy)
}

// SpeedUp is a paid mutator transaction binding the contract method 0xd3ed7c76.
//
// Solidity: function speedUp(address fromAsset, bytes txHash, uint256 fee) payable returns()
func (_PolyWrapperV3 *PolyWrapperV3Transactor) SpeedUp(opts *bind.TransactOpts, fromAsset common.Address, txHash []byte, fee *big.Int) (*types.Transaction, error) {
	return _PolyWrapperV3.contract.Transact(opts, "speedUp", fromAsset, txHash, fee)
}

// SpeedUp is a paid mutator transaction binding the contract method 0xd3ed7c76.
//
// Solidity: function speedUp(address fromAsset, bytes txHash, uint256 fee) payable returns()
func (_PolyWrapperV3 *PolyWrapperV3Session) SpeedUp(fromAsset common.Address, txHash []byte, fee *big.Int) (*types.Transaction, error) {
	return _PolyWrapperV3.Contract.SpeedUp(&_PolyWrapperV3.TransactOpts, fromAsset, txHash, fee)
}

// SpeedUp is a paid mutator transaction binding the contract method 0xd3ed7c76.
//
// Solidity: function speedUp(address fromAsset, bytes txHash, uint256 fee) payable returns()
func (_PolyWrapperV3 *PolyWrapperV3TransactorSession) SpeedUp(fromAsset common.Address, txHash []byte, fee *big.Int) (*types.Transaction, error) {
	return _PolyWrapperV3.Contract.SpeedUp(&_PolyWrapperV3.TransactOpts, fromAsset, txHash, fee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PolyWrapperV3 *PolyWrapperV3Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PolyWrapperV3.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PolyWrapperV3 *PolyWrapperV3Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PolyWrapperV3.Contract.TransferOwnership(&_PolyWrapperV3.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PolyWrapperV3 *PolyWrapperV3TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PolyWrapperV3.Contract.TransferOwnership(&_PolyWrapperV3.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PolyWrapperV3 *PolyWrapperV3Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolyWrapperV3.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PolyWrapperV3 *PolyWrapperV3Session) Unpause() (*types.Transaction, error) {
	return _PolyWrapperV3.Contract.Unpause(&_PolyWrapperV3.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PolyWrapperV3 *PolyWrapperV3TransactorSession) Unpause() (*types.Transaction, error) {
	return _PolyWrapperV3.Contract.Unpause(&_PolyWrapperV3.TransactOpts)
}

// PolyWrapperV3OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PolyWrapperV3 contract.
type PolyWrapperV3OwnershipTransferredIterator struct {
	Event *PolyWrapperV3OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PolyWrapperV3OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolyWrapperV3OwnershipTransferred)
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
		it.Event = new(PolyWrapperV3OwnershipTransferred)
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
func (it *PolyWrapperV3OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolyWrapperV3OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolyWrapperV3OwnershipTransferred represents a OwnershipTransferred event raised by the PolyWrapperV3 contract.
type PolyWrapperV3OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PolyWrapperV3 *PolyWrapperV3Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PolyWrapperV3OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PolyWrapperV3.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PolyWrapperV3OwnershipTransferredIterator{contract: _PolyWrapperV3.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PolyWrapperV3 *PolyWrapperV3Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PolyWrapperV3OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PolyWrapperV3.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolyWrapperV3OwnershipTransferred)
				if err := _PolyWrapperV3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PolyWrapperV3 *PolyWrapperV3Filterer) ParseOwnershipTransferred(log types.Log) (*PolyWrapperV3OwnershipTransferred, error) {
	event := new(PolyWrapperV3OwnershipTransferred)
	if err := _PolyWrapperV3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolyWrapperV3PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PolyWrapperV3 contract.
type PolyWrapperV3PausedIterator struct {
	Event *PolyWrapperV3Paused // Event containing the contract specifics and raw log

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
func (it *PolyWrapperV3PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolyWrapperV3Paused)
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
		it.Event = new(PolyWrapperV3Paused)
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
func (it *PolyWrapperV3PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolyWrapperV3PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolyWrapperV3Paused represents a Paused event raised by the PolyWrapperV3 contract.
type PolyWrapperV3Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PolyWrapperV3 *PolyWrapperV3Filterer) FilterPaused(opts *bind.FilterOpts) (*PolyWrapperV3PausedIterator, error) {

	logs, sub, err := _PolyWrapperV3.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PolyWrapperV3PausedIterator{contract: _PolyWrapperV3.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PolyWrapperV3 *PolyWrapperV3Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PolyWrapperV3Paused) (event.Subscription, error) {

	logs, sub, err := _PolyWrapperV3.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolyWrapperV3Paused)
				if err := _PolyWrapperV3.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_PolyWrapperV3 *PolyWrapperV3Filterer) ParsePaused(log types.Log) (*PolyWrapperV3Paused, error) {
	event := new(PolyWrapperV3Paused)
	if err := _PolyWrapperV3.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolyWrapperV3PolyWrapperLockIterator is returned from FilterPolyWrapperLock and is used to iterate over the raw logs and unpacked data for PolyWrapperLock events raised by the PolyWrapperV3 contract.
type PolyWrapperV3PolyWrapperLockIterator struct {
	Event *PolyWrapperV3PolyWrapperLock // Event containing the contract specifics and raw log

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
func (it *PolyWrapperV3PolyWrapperLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolyWrapperV3PolyWrapperLock)
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
		it.Event = new(PolyWrapperV3PolyWrapperLock)
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
func (it *PolyWrapperV3PolyWrapperLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolyWrapperV3PolyWrapperLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolyWrapperV3PolyWrapperLock represents a PolyWrapperLock event raised by the PolyWrapperV3 contract.
type PolyWrapperV3PolyWrapperLock struct {
	FromAsset common.Address
	Sender    common.Address
	ToChainId uint64
	ToAddress []byte
	Net       *big.Int
	Fee       *big.Int
	Id        *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPolyWrapperLock is a free log retrieval operation binding the contract event 0x2b0591052cc6602e870d3994f0a1b173fdac98c215cb3b0baf84eaca5a0aa81e.
//
// Solidity: event PolyWrapperLock(address indexed fromAsset, address indexed sender, uint64 toChainId, bytes toAddress, uint256 net, uint256 fee, uint256 id)
func (_PolyWrapperV3 *PolyWrapperV3Filterer) FilterPolyWrapperLock(opts *bind.FilterOpts, fromAsset []common.Address, sender []common.Address) (*PolyWrapperV3PolyWrapperLockIterator, error) {

	var fromAssetRule []interface{}
	for _, fromAssetItem := range fromAsset {
		fromAssetRule = append(fromAssetRule, fromAssetItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PolyWrapperV3.contract.FilterLogs(opts, "PolyWrapperLock", fromAssetRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PolyWrapperV3PolyWrapperLockIterator{contract: _PolyWrapperV3.contract, event: "PolyWrapperLock", logs: logs, sub: sub}, nil
}

// WatchPolyWrapperLock is a free log subscription operation binding the contract event 0x2b0591052cc6602e870d3994f0a1b173fdac98c215cb3b0baf84eaca5a0aa81e.
//
// Solidity: event PolyWrapperLock(address indexed fromAsset, address indexed sender, uint64 toChainId, bytes toAddress, uint256 net, uint256 fee, uint256 id)
func (_PolyWrapperV3 *PolyWrapperV3Filterer) WatchPolyWrapperLock(opts *bind.WatchOpts, sink chan<- *PolyWrapperV3PolyWrapperLock, fromAsset []common.Address, sender []common.Address) (event.Subscription, error) {

	var fromAssetRule []interface{}
	for _, fromAssetItem := range fromAsset {
		fromAssetRule = append(fromAssetRule, fromAssetItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PolyWrapperV3.contract.WatchLogs(opts, "PolyWrapperLock", fromAssetRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolyWrapperV3PolyWrapperLock)
				if err := _PolyWrapperV3.contract.UnpackLog(event, "PolyWrapperLock", log); err != nil {
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

// ParsePolyWrapperLock is a log parse operation binding the contract event 0x2b0591052cc6602e870d3994f0a1b173fdac98c215cb3b0baf84eaca5a0aa81e.
//
// Solidity: event PolyWrapperLock(address indexed fromAsset, address indexed sender, uint64 toChainId, bytes toAddress, uint256 net, uint256 fee, uint256 id)
func (_PolyWrapperV3 *PolyWrapperV3Filterer) ParsePolyWrapperLock(log types.Log) (*PolyWrapperV3PolyWrapperLock, error) {
	event := new(PolyWrapperV3PolyWrapperLock)
	if err := _PolyWrapperV3.contract.UnpackLog(event, "PolyWrapperLock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolyWrapperV3PolyWrapperSpeedUpIterator is returned from FilterPolyWrapperSpeedUp and is used to iterate over the raw logs and unpacked data for PolyWrapperSpeedUp events raised by the PolyWrapperV3 contract.
type PolyWrapperV3PolyWrapperSpeedUpIterator struct {
	Event *PolyWrapperV3PolyWrapperSpeedUp // Event containing the contract specifics and raw log

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
func (it *PolyWrapperV3PolyWrapperSpeedUpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolyWrapperV3PolyWrapperSpeedUp)
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
		it.Event = new(PolyWrapperV3PolyWrapperSpeedUp)
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
func (it *PolyWrapperV3PolyWrapperSpeedUpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolyWrapperV3PolyWrapperSpeedUpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolyWrapperV3PolyWrapperSpeedUp represents a PolyWrapperSpeedUp event raised by the PolyWrapperV3 contract.
type PolyWrapperV3PolyWrapperSpeedUp struct {
	FromAsset common.Address
	TxHash    common.Hash
	Sender    common.Address
	Efee      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPolyWrapperSpeedUp is a free log retrieval operation binding the contract event 0xf6579aef3e0d086d986c5d6972659f8a0d8602ef7945b054be1b88e088773ef6.
//
// Solidity: event PolyWrapperSpeedUp(address indexed fromAsset, bytes indexed txHash, address indexed sender, uint256 efee)
func (_PolyWrapperV3 *PolyWrapperV3Filterer) FilterPolyWrapperSpeedUp(opts *bind.FilterOpts, fromAsset []common.Address, txHash [][]byte, sender []common.Address) (*PolyWrapperV3PolyWrapperSpeedUpIterator, error) {

	var fromAssetRule []interface{}
	for _, fromAssetItem := range fromAsset {
		fromAssetRule = append(fromAssetRule, fromAssetItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PolyWrapperV3.contract.FilterLogs(opts, "PolyWrapperSpeedUp", fromAssetRule, txHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PolyWrapperV3PolyWrapperSpeedUpIterator{contract: _PolyWrapperV3.contract, event: "PolyWrapperSpeedUp", logs: logs, sub: sub}, nil
}

// WatchPolyWrapperSpeedUp is a free log subscription operation binding the contract event 0xf6579aef3e0d086d986c5d6972659f8a0d8602ef7945b054be1b88e088773ef6.
//
// Solidity: event PolyWrapperSpeedUp(address indexed fromAsset, bytes indexed txHash, address indexed sender, uint256 efee)
func (_PolyWrapperV3 *PolyWrapperV3Filterer) WatchPolyWrapperSpeedUp(opts *bind.WatchOpts, sink chan<- *PolyWrapperV3PolyWrapperSpeedUp, fromAsset []common.Address, txHash [][]byte, sender []common.Address) (event.Subscription, error) {

	var fromAssetRule []interface{}
	for _, fromAssetItem := range fromAsset {
		fromAssetRule = append(fromAssetRule, fromAssetItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PolyWrapperV3.contract.WatchLogs(opts, "PolyWrapperSpeedUp", fromAssetRule, txHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolyWrapperV3PolyWrapperSpeedUp)
				if err := _PolyWrapperV3.contract.UnpackLog(event, "PolyWrapperSpeedUp", log); err != nil {
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

// ParsePolyWrapperSpeedUp is a log parse operation binding the contract event 0xf6579aef3e0d086d986c5d6972659f8a0d8602ef7945b054be1b88e088773ef6.
//
// Solidity: event PolyWrapperSpeedUp(address indexed fromAsset, bytes indexed txHash, address indexed sender, uint256 efee)
func (_PolyWrapperV3 *PolyWrapperV3Filterer) ParsePolyWrapperSpeedUp(log types.Log) (*PolyWrapperV3PolyWrapperSpeedUp, error) {
	event := new(PolyWrapperV3PolyWrapperSpeedUp)
	if err := _PolyWrapperV3.contract.UnpackLog(event, "PolyWrapperSpeedUp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolyWrapperV3UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PolyWrapperV3 contract.
type PolyWrapperV3UnpausedIterator struct {
	Event *PolyWrapperV3Unpaused // Event containing the contract specifics and raw log

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
func (it *PolyWrapperV3UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolyWrapperV3Unpaused)
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
		it.Event = new(PolyWrapperV3Unpaused)
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
func (it *PolyWrapperV3UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolyWrapperV3UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolyWrapperV3Unpaused represents a Unpaused event raised by the PolyWrapperV3 contract.
type PolyWrapperV3Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PolyWrapperV3 *PolyWrapperV3Filterer) FilterUnpaused(opts *bind.FilterOpts) (*PolyWrapperV3UnpausedIterator, error) {

	logs, sub, err := _PolyWrapperV3.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PolyWrapperV3UnpausedIterator{contract: _PolyWrapperV3.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PolyWrapperV3 *PolyWrapperV3Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PolyWrapperV3Unpaused) (event.Subscription, error) {

	logs, sub, err := _PolyWrapperV3.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolyWrapperV3Unpaused)
				if err := _PolyWrapperV3.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_PolyWrapperV3 *PolyWrapperV3Filterer) ParseUnpaused(log types.Log) (*PolyWrapperV3Unpaused, error) {
	event := new(PolyWrapperV3Unpaused)
	if err := _PolyWrapperV3.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReentrancyGuardABI is the input ABI used to generate the binding from.
const ReentrancyGuardABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ReentrancyGuard is an auto generated Go binding around an Ethereum contract.
type ReentrancyGuard struct {
	ReentrancyGuardCaller     // Read-only binding to the contract
	ReentrancyGuardTransactor // Write-only binding to the contract
	ReentrancyGuardFilterer   // Log filterer for contract events
}

// ReentrancyGuardCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrancyGuardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrancyGuardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrancyGuardSession struct {
	Contract     *ReentrancyGuard  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReentrancyGuardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrancyGuardCallerSession struct {
	Contract *ReentrancyGuardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ReentrancyGuardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrancyGuardTransactorSession struct {
	Contract     *ReentrancyGuardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ReentrancyGuardRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrancyGuardRaw struct {
	Contract *ReentrancyGuard // Generic contract binding to access the raw methods on
}

// ReentrancyGuardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrancyGuardCallerRaw struct {
	Contract *ReentrancyGuardCaller // Generic read-only contract binding to access the raw methods on
}

// ReentrancyGuardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactorRaw struct {
	Contract *ReentrancyGuardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrancyGuard creates a new instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuard(address common.Address, backend bind.ContractBackend) (*ReentrancyGuard, error) {
	contract, err := bindReentrancyGuard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuard{ReentrancyGuardCaller: ReentrancyGuardCaller{contract: contract}, ReentrancyGuardTransactor: ReentrancyGuardTransactor{contract: contract}, ReentrancyGuardFilterer: ReentrancyGuardFilterer{contract: contract}}, nil
}

// NewReentrancyGuardCaller creates a new read-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardCaller(address common.Address, caller bind.ContractCaller) (*ReentrancyGuardCaller, error) {
	contract, err := bindReentrancyGuard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardCaller{contract: contract}, nil
}

// NewReentrancyGuardTransactor creates a new write-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentrancyGuardTransactor, error) {
	contract, err := bindReentrancyGuard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardTransactor{contract: contract}, nil
}

// NewReentrancyGuardFilterer creates a new log filterer instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentrancyGuardFilterer, error) {
	contract, err := bindReentrancyGuard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardFilterer{contract: contract}, nil
}

// bindReentrancyGuard binds a generic wrapper to an already deployed contract.
func bindReentrancyGuard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentrancyGuardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.ReentrancyGuardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transact(opts, method, params...)
}

// SafeERC20ABI is the input ABI used to generate the binding from.
const SafeERC20ABI = "[]"

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
var SafeERC20Bin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158203660764081bae01ba4329738a536eb2c8eff9b380f5a232f5dbc713458762f4164736f6c63430005110032"

// DeploySafeERC20 deploys a new Ethereum contract, binding an instance of SafeERC20 to it.
func DeploySafeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	SafeERC20Caller     // Read-only binding to the contract
	SafeERC20Transactor // Write-only binding to the contract
	SafeERC20Filterer   // Log filterer for contract events
}

// SafeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20Session struct {
	Contract     *SafeERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20CallerSession struct {
	Contract *SafeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransactorSession struct {
	Contract     *SafeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20Raw struct {
	Contract *SafeERC20 // Generic contract binding to access the raw methods on
}

// SafeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20CallerRaw struct {
	Contract *SafeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransactorRaw struct {
	Contract *SafeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20 creates a new instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20(address common.Address, backend bind.ContractBackend) (*SafeERC20, error) {
	contract, err := bindSafeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// NewSafeERC20Caller creates a new read-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Caller(address common.Address, caller bind.ContractCaller) (*SafeERC20Caller, error) {
	contract, err := bindSafeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Caller{contract: contract}, nil
}

// NewSafeERC20Transactor creates a new write-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20Transactor, error) {
	contract, err := bindSafeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Transactor{contract: contract}, nil
}

// NewSafeERC20Filterer creates a new log filterer instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20Filterer, error) {
	contract, err := bindSafeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Filterer{contract: contract}, nil
}

// bindSafeERC20 binds a generic wrapper to an already deployed contract.
func bindSafeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.SafeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transact(opts, method, params...)
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

// UtilsABI is the input ABI used to generate the binding from.
const UtilsABI = "[]"

// UtilsBin is the compiled bytecode used for deploying new contracts.
var UtilsBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820b05a22f851e3834f0b38887b082cb367a95f2542aa7f7c61276a9ed0ab9bb8c264736f6c63430005110032"

// DeployUtils deploys a new Ethereum contract, binding an instance of Utils to it.
func DeployUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Utils, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Utils{UtilsCaller: UtilsCaller{contract: contract}, UtilsTransactor: UtilsTransactor{contract: contract}, UtilsFilterer: UtilsFilterer{contract: contract}}, nil
}

// Utils is an auto generated Go binding around an Ethereum contract.
type Utils struct {
	UtilsCaller     // Read-only binding to the contract
	UtilsTransactor // Write-only binding to the contract
	UtilsFilterer   // Log filterer for contract events
}

// UtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type UtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UtilsSession struct {
	Contract     *Utils            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UtilsCallerSession struct {
	Contract *UtilsCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UtilsTransactorSession struct {
	Contract     *UtilsTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type UtilsRaw struct {
	Contract *Utils // Generic contract binding to access the raw methods on
}

// UtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UtilsCallerRaw struct {
	Contract *UtilsCaller // Generic read-only contract binding to access the raw methods on
}

// UtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UtilsTransactorRaw struct {
	Contract *UtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUtils creates a new instance of Utils, bound to a specific deployed contract.
func NewUtils(address common.Address, backend bind.ContractBackend) (*Utils, error) {
	contract, err := bindUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Utils{UtilsCaller: UtilsCaller{contract: contract}, UtilsTransactor: UtilsTransactor{contract: contract}, UtilsFilterer: UtilsFilterer{contract: contract}}, nil
}

// NewUtilsCaller creates a new read-only instance of Utils, bound to a specific deployed contract.
func NewUtilsCaller(address common.Address, caller bind.ContractCaller) (*UtilsCaller, error) {
	contract, err := bindUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsCaller{contract: contract}, nil
}

// NewUtilsTransactor creates a new write-only instance of Utils, bound to a specific deployed contract.
func NewUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*UtilsTransactor, error) {
	contract, err := bindUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsTransactor{contract: contract}, nil
}

// NewUtilsFilterer creates a new log filterer instance of Utils, bound to a specific deployed contract.
func NewUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*UtilsFilterer, error) {
	contract, err := bindUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UtilsFilterer{contract: contract}, nil
}

// bindUtils binds a generic wrapper to an already deployed contract.
func bindUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utils *UtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Utils.Contract.UtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utils *UtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utils.Contract.UtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utils *UtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utils.Contract.UtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utils *UtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Utils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utils *UtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utils *UtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utils.Contract.contract.Transact(opts, method, params...)
}
