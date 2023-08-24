// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staking

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

// StakingAccountStakingInfo is an auto generated low-level Go binding around an user-defined struct.
type StakingAccountStakingInfo struct {
	Addr     common.Address
	Amount   *big.Int
	State    uint8
	TimeLock *big.Int
}

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minNumValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNumValidators\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"grantContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeLock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenSlashing\",\"type\":\"uint256\"}],\"name\":\"Ban\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"BurnGrantInitial\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelegatorStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelegatorUnstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_grantContract\",\"type\":\"address\"}],\"name\":\"GrantContractChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"earnedAmount\",\"type\":\"uint256\"}],\"name\":\"ReclaimGrant\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeLock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"suspendCounter\",\"type\":\"uint64\"}],\"name\":\"Suspended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"Warning\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BAN_DURATION\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COUNTER_SUSPEND\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DELEGATION_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GrantAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_VALIDATORSUBSET_SIZE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_VALIDATORSUBSET_SIZE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERCENTAGE_TOKEN_SLASHING\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUSPEND_DURATION\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_crewName\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"boss\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_grantContract\",\"outputs\":[{\"internalType\":\"contractGrantContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_maximumNumDelegators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_maximumNumValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_minimumNumDelegators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_minimumNumValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_stakedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_validators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"accountStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressGrantContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"addressToStakedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"badValidator\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"validatorsSubsetSize\",\"type\":\"uint64\"}],\"name\":\"ban\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"burnGrantInitial\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkStateValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"delegatorUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"getDelegatedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"getDelegatorsInfoOfValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumStaking.Status\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timeLock\",\"type\":\"uint256\"}],\"internalType\":\"structStaking.AccountStakingInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"getStringLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getTotalDelegatorOfValidator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorsStakeInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumStaking.Status\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timeLock\",\"type\":\"uint256\"}],\"internalType\":\"structStaking.AccountStakingInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorsSubsetTimeout\",\"outputs\":[{\"internalType\":\"uint8[]\",\"name\":\"\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isStaker\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maximumNumValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumNumValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"reclaimUnusedGrant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetValidatorsSubsetTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"setGrantContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"stakeGrantContract\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"stakeSignerGrantContract\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"indexValidator\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"badValidator\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"count\",\"type\":\"uint8\"}],\"name\":\"suspend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"badValidator\",\"type\":\"address\"}],\"name\":\"warning\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// BANDURATION is a free data retrieval call binding the contract method 0xbb4c461e.
//
// Solidity: function BAN_DURATION() view returns(uint128)
func (_Staking *StakingCaller) BANDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "BAN_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BANDURATION is a free data retrieval call binding the contract method 0xbb4c461e.
//
// Solidity: function BAN_DURATION() view returns(uint128)
func (_Staking *StakingSession) BANDURATION() (*big.Int, error) {
	return _Staking.Contract.BANDURATION(&_Staking.CallOpts)
}

// BANDURATION is a free data retrieval call binding the contract method 0xbb4c461e.
//
// Solidity: function BAN_DURATION() view returns(uint128)
func (_Staking *StakingCallerSession) BANDURATION() (*big.Int, error) {
	return _Staking.Contract.BANDURATION(&_Staking.CallOpts)
}

// COUNTERSUSPEND is a free data retrieval call binding the contract method 0xf740ff8e.
//
// Solidity: function COUNTER_SUSPEND() view returns(uint8)
func (_Staking *StakingCaller) COUNTERSUSPEND(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "COUNTER_SUSPEND")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// COUNTERSUSPEND is a free data retrieval call binding the contract method 0xf740ff8e.
//
// Solidity: function COUNTER_SUSPEND() view returns(uint8)
func (_Staking *StakingSession) COUNTERSUSPEND() (uint8, error) {
	return _Staking.Contract.COUNTERSUSPEND(&_Staking.CallOpts)
}

// COUNTERSUSPEND is a free data retrieval call binding the contract method 0xf740ff8e.
//
// Solidity: function COUNTER_SUSPEND() view returns(uint8)
func (_Staking *StakingCallerSession) COUNTERSUSPEND() (uint8, error) {
	return _Staking.Contract.COUNTERSUSPEND(&_Staking.CallOpts)
}

// DELEGATIONTHRESHOLD is a free data retrieval call binding the contract method 0x3c449ae7.
//
// Solidity: function DELEGATION_THRESHOLD() view returns(uint128)
func (_Staking *StakingCaller) DELEGATIONTHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "DELEGATION_THRESHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DELEGATIONTHRESHOLD is a free data retrieval call binding the contract method 0x3c449ae7.
//
// Solidity: function DELEGATION_THRESHOLD() view returns(uint128)
func (_Staking *StakingSession) DELEGATIONTHRESHOLD() (*big.Int, error) {
	return _Staking.Contract.DELEGATIONTHRESHOLD(&_Staking.CallOpts)
}

// DELEGATIONTHRESHOLD is a free data retrieval call binding the contract method 0x3c449ae7.
//
// Solidity: function DELEGATION_THRESHOLD() view returns(uint128)
func (_Staking *StakingCallerSession) DELEGATIONTHRESHOLD() (*big.Int, error) {
	return _Staking.Contract.DELEGATIONTHRESHOLD(&_Staking.CallOpts)
}

// GrantAmount is a free data retrieval call binding the contract method 0x15f9eb9b.
//
// Solidity: function GrantAmount() view returns(uint256)
func (_Staking *StakingCaller) GrantAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GrantAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GrantAmount is a free data retrieval call binding the contract method 0x15f9eb9b.
//
// Solidity: function GrantAmount() view returns(uint256)
func (_Staking *StakingSession) GrantAmount() (*big.Int, error) {
	return _Staking.Contract.GrantAmount(&_Staking.CallOpts)
}

// GrantAmount is a free data retrieval call binding the contract method 0x15f9eb9b.
//
// Solidity: function GrantAmount() view returns(uint256)
func (_Staking *StakingCallerSession) GrantAmount() (*big.Int, error) {
	return _Staking.Contract.GrantAmount(&_Staking.CallOpts)
}

// MAXIMUMVALIDATORSUBSETSIZE is a free data retrieval call binding the contract method 0x7e840377.
//
// Solidity: function MAXIMUM_VALIDATORSUBSET_SIZE() view returns(uint64)
func (_Staking *StakingCaller) MAXIMUMVALIDATORSUBSETSIZE(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "MAXIMUM_VALIDATORSUBSET_SIZE")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MAXIMUMVALIDATORSUBSETSIZE is a free data retrieval call binding the contract method 0x7e840377.
//
// Solidity: function MAXIMUM_VALIDATORSUBSET_SIZE() view returns(uint64)
func (_Staking *StakingSession) MAXIMUMVALIDATORSUBSETSIZE() (uint64, error) {
	return _Staking.Contract.MAXIMUMVALIDATORSUBSETSIZE(&_Staking.CallOpts)
}

// MAXIMUMVALIDATORSUBSETSIZE is a free data retrieval call binding the contract method 0x7e840377.
//
// Solidity: function MAXIMUM_VALIDATORSUBSET_SIZE() view returns(uint64)
func (_Staking *StakingCallerSession) MAXIMUMVALIDATORSUBSETSIZE() (uint64, error) {
	return _Staking.Contract.MAXIMUMVALIDATORSUBSETSIZE(&_Staking.CallOpts)
}

// MINIMUMVALIDATORSUBSETSIZE is a free data retrieval call binding the contract method 0xde092ac4.
//
// Solidity: function MINIMUM_VALIDATORSUBSET_SIZE() view returns(uint64)
func (_Staking *StakingCaller) MINIMUMVALIDATORSUBSETSIZE(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "MINIMUM_VALIDATORSUBSET_SIZE")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MINIMUMVALIDATORSUBSETSIZE is a free data retrieval call binding the contract method 0xde092ac4.
//
// Solidity: function MINIMUM_VALIDATORSUBSET_SIZE() view returns(uint64)
func (_Staking *StakingSession) MINIMUMVALIDATORSUBSETSIZE() (uint64, error) {
	return _Staking.Contract.MINIMUMVALIDATORSUBSETSIZE(&_Staking.CallOpts)
}

// MINIMUMVALIDATORSUBSETSIZE is a free data retrieval call binding the contract method 0xde092ac4.
//
// Solidity: function MINIMUM_VALIDATORSUBSET_SIZE() view returns(uint64)
func (_Staking *StakingCallerSession) MINIMUMVALIDATORSUBSETSIZE() (uint64, error) {
	return _Staking.Contract.MINIMUMVALIDATORSUBSETSIZE(&_Staking.CallOpts)
}

// PERCENTAGETOKENSLASHING is a free data retrieval call binding the contract method 0x5c8c4a07.
//
// Solidity: function PERCENTAGE_TOKEN_SLASHING() view returns(uint128)
func (_Staking *StakingCaller) PERCENTAGETOKENSLASHING(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "PERCENTAGE_TOKEN_SLASHING")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PERCENTAGETOKENSLASHING is a free data retrieval call binding the contract method 0x5c8c4a07.
//
// Solidity: function PERCENTAGE_TOKEN_SLASHING() view returns(uint128)
func (_Staking *StakingSession) PERCENTAGETOKENSLASHING() (*big.Int, error) {
	return _Staking.Contract.PERCENTAGETOKENSLASHING(&_Staking.CallOpts)
}

// PERCENTAGETOKENSLASHING is a free data retrieval call binding the contract method 0x5c8c4a07.
//
// Solidity: function PERCENTAGE_TOKEN_SLASHING() view returns(uint128)
func (_Staking *StakingCallerSession) PERCENTAGETOKENSLASHING() (*big.Int, error) {
	return _Staking.Contract.PERCENTAGETOKENSLASHING(&_Staking.CallOpts)
}

// SUSPENDDURATION is a free data retrieval call binding the contract method 0x5bb9fec9.
//
// Solidity: function SUSPEND_DURATION() view returns(uint128)
func (_Staking *StakingCaller) SUSPENDDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "SUSPEND_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SUSPENDDURATION is a free data retrieval call binding the contract method 0x5bb9fec9.
//
// Solidity: function SUSPEND_DURATION() view returns(uint128)
func (_Staking *StakingSession) SUSPENDDURATION() (*big.Int, error) {
	return _Staking.Contract.SUSPENDDURATION(&_Staking.CallOpts)
}

// SUSPENDDURATION is a free data retrieval call binding the contract method 0x5bb9fec9.
//
// Solidity: function SUSPEND_DURATION() view returns(uint128)
func (_Staking *StakingCallerSession) SUSPENDDURATION() (*big.Int, error) {
	return _Staking.Contract.SUSPENDDURATION(&_Staking.CallOpts)
}

// VALIDATORTHRESHOLD is a free data retrieval call binding the contract method 0x7a6eea37.
//
// Solidity: function VALIDATOR_THRESHOLD() view returns(uint128)
func (_Staking *StakingCaller) VALIDATORTHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "VALIDATOR_THRESHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VALIDATORTHRESHOLD is a free data retrieval call binding the contract method 0x7a6eea37.
//
// Solidity: function VALIDATOR_THRESHOLD() view returns(uint128)
func (_Staking *StakingSession) VALIDATORTHRESHOLD() (*big.Int, error) {
	return _Staking.Contract.VALIDATORTHRESHOLD(&_Staking.CallOpts)
}

// VALIDATORTHRESHOLD is a free data retrieval call binding the contract method 0x7a6eea37.
//
// Solidity: function VALIDATOR_THRESHOLD() view returns(uint128)
func (_Staking *StakingCallerSession) VALIDATORTHRESHOLD() (*big.Int, error) {
	return _Staking.Contract.VALIDATORTHRESHOLD(&_Staking.CallOpts)
}

// CrewName is a free data retrieval call binding the contract method 0x7ebe444c.
//
// Solidity: function _crewName(uint256 ) view returns(address boss, string name)
func (_Staking *StakingCaller) CrewName(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Boss common.Address
	Name string
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_crewName", arg0)

	outstruct := new(struct {
		Boss common.Address
		Name string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Boss = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// CrewName is a free data retrieval call binding the contract method 0x7ebe444c.
//
// Solidity: function _crewName(uint256 ) view returns(address boss, string name)
func (_Staking *StakingSession) CrewName(arg0 *big.Int) (struct {
	Boss common.Address
	Name string
}, error) {
	return _Staking.Contract.CrewName(&_Staking.CallOpts, arg0)
}

// CrewName is a free data retrieval call binding the contract method 0x7ebe444c.
//
// Solidity: function _crewName(uint256 ) view returns(address boss, string name)
func (_Staking *StakingCallerSession) CrewName(arg0 *big.Int) (struct {
	Boss common.Address
	Name string
}, error) {
	return _Staking.Contract.CrewName(&_Staking.CallOpts, arg0)
}

// GrantContract is a free data retrieval call binding the contract method 0xcb3d89d6.
//
// Solidity: function _grantContract() view returns(address)
func (_Staking *StakingCaller) GrantContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_grantContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GrantContract is a free data retrieval call binding the contract method 0xcb3d89d6.
//
// Solidity: function _grantContract() view returns(address)
func (_Staking *StakingSession) GrantContract() (common.Address, error) {
	return _Staking.Contract.GrantContract(&_Staking.CallOpts)
}

// GrantContract is a free data retrieval call binding the contract method 0xcb3d89d6.
//
// Solidity: function _grantContract() view returns(address)
func (_Staking *StakingCallerSession) GrantContract() (common.Address, error) {
	return _Staking.Contract.GrantContract(&_Staking.CallOpts)
}

// MaximumNumDelegators is a free data retrieval call binding the contract method 0x90b19627.
//
// Solidity: function _maximumNumDelegators() view returns(uint256)
func (_Staking *StakingCaller) MaximumNumDelegators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_maximumNumDelegators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaximumNumDelegators is a free data retrieval call binding the contract method 0x90b19627.
//
// Solidity: function _maximumNumDelegators() view returns(uint256)
func (_Staking *StakingSession) MaximumNumDelegators() (*big.Int, error) {
	return _Staking.Contract.MaximumNumDelegators(&_Staking.CallOpts)
}

// MaximumNumDelegators is a free data retrieval call binding the contract method 0x90b19627.
//
// Solidity: function _maximumNumDelegators() view returns(uint256)
func (_Staking *StakingCallerSession) MaximumNumDelegators() (*big.Int, error) {
	return _Staking.Contract.MaximumNumDelegators(&_Staking.CallOpts)
}

// MaximumNumValidators1 is a free data retrieval call binding the contract method 0xaf6da36e.
//
// Solidity: function _maximumNumValidators() view returns(uint256)
func (_Staking *StakingCaller) MaximumNumValidators1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_maximumNumValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaximumNumValidators1 is a free data retrieval call binding the contract method 0xaf6da36e.
//
// Solidity: function _maximumNumValidators() view returns(uint256)
func (_Staking *StakingSession) MaximumNumValidators1() (*big.Int, error) {
	return _Staking.Contract.MaximumNumValidators1(&_Staking.CallOpts)
}

// MaximumNumValidators1 is a free data retrieval call binding the contract method 0xaf6da36e.
//
// Solidity: function _maximumNumValidators() view returns(uint256)
func (_Staking *StakingCallerSession) MaximumNumValidators1() (*big.Int, error) {
	return _Staking.Contract.MaximumNumValidators1(&_Staking.CallOpts)
}

// MinimumNumDelegators is a free data retrieval call binding the contract method 0x36c4086f.
//
// Solidity: function _minimumNumDelegators() view returns(uint256)
func (_Staking *StakingCaller) MinimumNumDelegators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_minimumNumDelegators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumNumDelegators is a free data retrieval call binding the contract method 0x36c4086f.
//
// Solidity: function _minimumNumDelegators() view returns(uint256)
func (_Staking *StakingSession) MinimumNumDelegators() (*big.Int, error) {
	return _Staking.Contract.MinimumNumDelegators(&_Staking.CallOpts)
}

// MinimumNumDelegators is a free data retrieval call binding the contract method 0x36c4086f.
//
// Solidity: function _minimumNumDelegators() view returns(uint256)
func (_Staking *StakingCallerSession) MinimumNumDelegators() (*big.Int, error) {
	return _Staking.Contract.MinimumNumDelegators(&_Staking.CallOpts)
}

// MinimumNumValidators is a free data retrieval call binding the contract method 0xc795c077.
//
// Solidity: function _minimumNumValidators() view returns(uint256)
func (_Staking *StakingCaller) MinimumNumValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_minimumNumValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumNumValidators is a free data retrieval call binding the contract method 0xc795c077.
//
// Solidity: function _minimumNumValidators() view returns(uint256)
func (_Staking *StakingSession) MinimumNumValidators() (*big.Int, error) {
	return _Staking.Contract.MinimumNumValidators(&_Staking.CallOpts)
}

// MinimumNumValidators is a free data retrieval call binding the contract method 0xc795c077.
//
// Solidity: function _minimumNumValidators() view returns(uint256)
func (_Staking *StakingCallerSession) MinimumNumValidators() (*big.Int, error) {
	return _Staking.Contract.MinimumNumValidators(&_Staking.CallOpts)
}

// StakedAmount is a free data retrieval call binding the contract method 0xe387a7ed.
//
// Solidity: function _stakedAmount() view returns(uint256)
func (_Staking *StakingCaller) StakedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_stakedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedAmount is a free data retrieval call binding the contract method 0xe387a7ed.
//
// Solidity: function _stakedAmount() view returns(uint256)
func (_Staking *StakingSession) StakedAmount() (*big.Int, error) {
	return _Staking.Contract.StakedAmount(&_Staking.CallOpts)
}

// StakedAmount is a free data retrieval call binding the contract method 0xe387a7ed.
//
// Solidity: function _stakedAmount() view returns(uint256)
func (_Staking *StakingCallerSession) StakedAmount() (*big.Int, error) {
	return _Staking.Contract.StakedAmount(&_Staking.CallOpts)
}

// Validators1 is a free data retrieval call binding the contract method 0xf90ecacc.
//
// Solidity: function _validators(uint256 ) view returns(address)
func (_Staking *StakingCaller) Validators1(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_validators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Validators1 is a free data retrieval call binding the contract method 0xf90ecacc.
//
// Solidity: function _validators(uint256 ) view returns(address)
func (_Staking *StakingSession) Validators1(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.Validators1(&_Staking.CallOpts, arg0)
}

// Validators1 is a free data retrieval call binding the contract method 0xf90ecacc.
//
// Solidity: function _validators(uint256 ) view returns(address)
func (_Staking *StakingCallerSession) Validators1(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.Validators1(&_Staking.CallOpts, arg0)
}

// AccountStake is a free data retrieval call binding the contract method 0x2367f6b5.
//
// Solidity: function accountStake(address addr) view returns(uint256)
func (_Staking *StakingCaller) AccountStake(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "accountStake", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountStake is a free data retrieval call binding the contract method 0x2367f6b5.
//
// Solidity: function accountStake(address addr) view returns(uint256)
func (_Staking *StakingSession) AccountStake(addr common.Address) (*big.Int, error) {
	return _Staking.Contract.AccountStake(&_Staking.CallOpts, addr)
}

// AccountStake is a free data retrieval call binding the contract method 0x2367f6b5.
//
// Solidity: function accountStake(address addr) view returns(uint256)
func (_Staking *StakingCallerSession) AccountStake(addr common.Address) (*big.Int, error) {
	return _Staking.Contract.AccountStake(&_Staking.CallOpts, addr)
}

// AddressGrantContract is a free data retrieval call binding the contract method 0xe69af0c7.
//
// Solidity: function addressGrantContract() view returns(address)
func (_Staking *StakingCaller) AddressGrantContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "addressGrantContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressGrantContract is a free data retrieval call binding the contract method 0xe69af0c7.
//
// Solidity: function addressGrantContract() view returns(address)
func (_Staking *StakingSession) AddressGrantContract() (common.Address, error) {
	return _Staking.Contract.AddressGrantContract(&_Staking.CallOpts)
}

// AddressGrantContract is a free data retrieval call binding the contract method 0xe69af0c7.
//
// Solidity: function addressGrantContract() view returns(address)
func (_Staking *StakingCallerSession) AddressGrantContract() (common.Address, error) {
	return _Staking.Contract.AddressGrantContract(&_Staking.CallOpts)
}

// AddressToStakedAmount is a free data retrieval call binding the contract method 0xc3a6ce2b.
//
// Solidity: function addressToStakedAmount(address sender) view returns(uint256)
func (_Staking *StakingCaller) AddressToStakedAmount(opts *bind.CallOpts, sender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "addressToStakedAmount", sender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressToStakedAmount is a free data retrieval call binding the contract method 0xc3a6ce2b.
//
// Solidity: function addressToStakedAmount(address sender) view returns(uint256)
func (_Staking *StakingSession) AddressToStakedAmount(sender common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToStakedAmount(&_Staking.CallOpts, sender)
}

// AddressToStakedAmount is a free data retrieval call binding the contract method 0xc3a6ce2b.
//
// Solidity: function addressToStakedAmount(address sender) view returns(uint256)
func (_Staking *StakingCallerSession) AddressToStakedAmount(sender common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToStakedAmount(&_Staking.CallOpts, sender)
}

// GetDelegatedAmount is a free data retrieval call binding the contract method 0x636592ed.
//
// Solidity: function getDelegatedAmount(address account, address signer) view returns(uint256)
func (_Staking *StakingCaller) GetDelegatedAmount(opts *bind.CallOpts, account common.Address, signer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getDelegatedAmount", account, signer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDelegatedAmount is a free data retrieval call binding the contract method 0x636592ed.
//
// Solidity: function getDelegatedAmount(address account, address signer) view returns(uint256)
func (_Staking *StakingSession) GetDelegatedAmount(account common.Address, signer common.Address) (*big.Int, error) {
	return _Staking.Contract.GetDelegatedAmount(&_Staking.CallOpts, account, signer)
}

// GetDelegatedAmount is a free data retrieval call binding the contract method 0x636592ed.
//
// Solidity: function getDelegatedAmount(address account, address signer) view returns(uint256)
func (_Staking *StakingCallerSession) GetDelegatedAmount(account common.Address, signer common.Address) (*big.Int, error) {
	return _Staking.Contract.GetDelegatedAmount(&_Staking.CallOpts, account, signer)
}

// GetDelegatorsInfoOfValidator is a free data retrieval call binding the contract method 0xb1f72b52.
//
// Solidity: function getDelegatorsInfoOfValidator(address signer) view returns((address,uint256,uint8,uint256)[])
func (_Staking *StakingCaller) GetDelegatorsInfoOfValidator(opts *bind.CallOpts, signer common.Address) ([]StakingAccountStakingInfo, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getDelegatorsInfoOfValidator", signer)

	if err != nil {
		return *new([]StakingAccountStakingInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]StakingAccountStakingInfo)).(*[]StakingAccountStakingInfo)

	return out0, err

}

// GetDelegatorsInfoOfValidator is a free data retrieval call binding the contract method 0xb1f72b52.
//
// Solidity: function getDelegatorsInfoOfValidator(address signer) view returns((address,uint256,uint8,uint256)[])
func (_Staking *StakingSession) GetDelegatorsInfoOfValidator(signer common.Address) ([]StakingAccountStakingInfo, error) {
	return _Staking.Contract.GetDelegatorsInfoOfValidator(&_Staking.CallOpts, signer)
}

// GetDelegatorsInfoOfValidator is a free data retrieval call binding the contract method 0xb1f72b52.
//
// Solidity: function getDelegatorsInfoOfValidator(address signer) view returns((address,uint256,uint8,uint256)[])
func (_Staking *StakingCallerSession) GetDelegatorsInfoOfValidator(signer common.Address) ([]StakingAccountStakingInfo, error) {
	return _Staking.Contract.GetDelegatorsInfoOfValidator(&_Staking.CallOpts, signer)
}

// GetStringLength is a free data retrieval call binding the contract method 0x65c19af0.
//
// Solidity: function getStringLength(string str) pure returns(uint256)
func (_Staking *StakingCaller) GetStringLength(opts *bind.CallOpts, str string) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getStringLength", str)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStringLength is a free data retrieval call binding the contract method 0x65c19af0.
//
// Solidity: function getStringLength(string str) pure returns(uint256)
func (_Staking *StakingSession) GetStringLength(str string) (*big.Int, error) {
	return _Staking.Contract.GetStringLength(&_Staking.CallOpts, str)
}

// GetStringLength is a free data retrieval call binding the contract method 0x65c19af0.
//
// Solidity: function getStringLength(string str) pure returns(uint256)
func (_Staking *StakingCallerSession) GetStringLength(str string) (*big.Int, error) {
	return _Staking.Contract.GetStringLength(&_Staking.CallOpts, str)
}

// GetTotalDelegatorOfValidator is a free data retrieval call binding the contract method 0xed546c8d.
//
// Solidity: function getTotalDelegatorOfValidator(address account) view returns(uint256)
func (_Staking *StakingCaller) GetTotalDelegatorOfValidator(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getTotalDelegatorOfValidator", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalDelegatorOfValidator is a free data retrieval call binding the contract method 0xed546c8d.
//
// Solidity: function getTotalDelegatorOfValidator(address account) view returns(uint256)
func (_Staking *StakingSession) GetTotalDelegatorOfValidator(account common.Address) (*big.Int, error) {
	return _Staking.Contract.GetTotalDelegatorOfValidator(&_Staking.CallOpts, account)
}

// GetTotalDelegatorOfValidator is a free data retrieval call binding the contract method 0xed546c8d.
//
// Solidity: function getTotalDelegatorOfValidator(address account) view returns(uint256)
func (_Staking *StakingCallerSession) GetTotalDelegatorOfValidator(account common.Address) (*big.Int, error) {
	return _Staking.Contract.GetTotalDelegatorOfValidator(&_Staking.CallOpts, account)
}

// GetValidatorsStakeInfo is a free data retrieval call binding the contract method 0x03306f0a.
//
// Solidity: function getValidatorsStakeInfo() view returns((address,uint256,uint8,uint256)[])
func (_Staking *StakingCaller) GetValidatorsStakeInfo(opts *bind.CallOpts) ([]StakingAccountStakingInfo, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getValidatorsStakeInfo")

	if err != nil {
		return *new([]StakingAccountStakingInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]StakingAccountStakingInfo)).(*[]StakingAccountStakingInfo)

	return out0, err

}

// GetValidatorsStakeInfo is a free data retrieval call binding the contract method 0x03306f0a.
//
// Solidity: function getValidatorsStakeInfo() view returns((address,uint256,uint8,uint256)[])
func (_Staking *StakingSession) GetValidatorsStakeInfo() ([]StakingAccountStakingInfo, error) {
	return _Staking.Contract.GetValidatorsStakeInfo(&_Staking.CallOpts)
}

// GetValidatorsStakeInfo is a free data retrieval call binding the contract method 0x03306f0a.
//
// Solidity: function getValidatorsStakeInfo() view returns((address,uint256,uint8,uint256)[])
func (_Staking *StakingCallerSession) GetValidatorsStakeInfo() ([]StakingAccountStakingInfo, error) {
	return _Staking.Contract.GetValidatorsStakeInfo(&_Staking.CallOpts)
}

// GetValidatorsSubsetTimeout is a free data retrieval call binding the contract method 0xf6676a3a.
//
// Solidity: function getValidatorsSubsetTimeout() view returns(uint8[])
func (_Staking *StakingCaller) GetValidatorsSubsetTimeout(opts *bind.CallOpts) ([]uint8, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getValidatorsSubsetTimeout")

	if err != nil {
		return *new([]uint8), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint8)).(*[]uint8)

	return out0, err

}

// GetValidatorsSubsetTimeout is a free data retrieval call binding the contract method 0xf6676a3a.
//
// Solidity: function getValidatorsSubsetTimeout() view returns(uint8[])
func (_Staking *StakingSession) GetValidatorsSubsetTimeout() ([]uint8, error) {
	return _Staking.Contract.GetValidatorsSubsetTimeout(&_Staking.CallOpts)
}

// GetValidatorsSubsetTimeout is a free data retrieval call binding the contract method 0xf6676a3a.
//
// Solidity: function getValidatorsSubsetTimeout() view returns(uint8[])
func (_Staking *StakingCallerSession) GetValidatorsSubsetTimeout() ([]uint8, error) {
	return _Staking.Contract.GetValidatorsSubsetTimeout(&_Staking.CallOpts)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address sender) view returns(bool)
func (_Staking *StakingCaller) IsSigner(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isSigner", sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address sender) view returns(bool)
func (_Staking *StakingSession) IsSigner(sender common.Address) (bool, error) {
	return _Staking.Contract.IsSigner(&_Staking.CallOpts, sender)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address sender) view returns(bool)
func (_Staking *StakingCallerSession) IsSigner(sender common.Address) (bool, error) {
	return _Staking.Contract.IsSigner(&_Staking.CallOpts, sender)
}

// IsStaker is a free data retrieval call binding the contract method 0x6f1e8533.
//
// Solidity: function isStaker(address sender) view returns(bool)
func (_Staking *StakingCaller) IsStaker(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isStaker", sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStaker is a free data retrieval call binding the contract method 0x6f1e8533.
//
// Solidity: function isStaker(address sender) view returns(bool)
func (_Staking *StakingSession) IsStaker(sender common.Address) (bool, error) {
	return _Staking.Contract.IsStaker(&_Staking.CallOpts, sender)
}

// IsStaker is a free data retrieval call binding the contract method 0x6f1e8533.
//
// Solidity: function isStaker(address sender) view returns(bool)
func (_Staking *StakingCallerSession) IsStaker(sender common.Address) (bool, error) {
	return _Staking.Contract.IsStaker(&_Staking.CallOpts, sender)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address addr) view returns(bool)
func (_Staking *StakingCaller) IsValidator(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isValidator", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address addr) view returns(bool)
func (_Staking *StakingSession) IsValidator(addr common.Address) (bool, error) {
	return _Staking.Contract.IsValidator(&_Staking.CallOpts, addr)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address addr) view returns(bool)
func (_Staking *StakingCallerSession) IsValidator(addr common.Address) (bool, error) {
	return _Staking.Contract.IsValidator(&_Staking.CallOpts, addr)
}

// MaximumNumValidators is a free data retrieval call binding the contract method 0xe804fbf6.
//
// Solidity: function maximumNumValidators() view returns(uint256)
func (_Staking *StakingCaller) MaximumNumValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "maximumNumValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaximumNumValidators is a free data retrieval call binding the contract method 0xe804fbf6.
//
// Solidity: function maximumNumValidators() view returns(uint256)
func (_Staking *StakingSession) MaximumNumValidators() (*big.Int, error) {
	return _Staking.Contract.MaximumNumValidators(&_Staking.CallOpts)
}

// MaximumNumValidators is a free data retrieval call binding the contract method 0xe804fbf6.
//
// Solidity: function maximumNumValidators() view returns(uint256)
func (_Staking *StakingCallerSession) MaximumNumValidators() (*big.Int, error) {
	return _Staking.Contract.MaximumNumValidators(&_Staking.CallOpts)
}

// MinimumNumValidators1 is a free data retrieval call binding the contract method 0x714ff425.
//
// Solidity: function minimumNumValidators() view returns(uint256)
func (_Staking *StakingCaller) MinimumNumValidators1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "minimumNumValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumNumValidators1 is a free data retrieval call binding the contract method 0x714ff425.
//
// Solidity: function minimumNumValidators() view returns(uint256)
func (_Staking *StakingSession) MinimumNumValidators1() (*big.Int, error) {
	return _Staking.Contract.MinimumNumValidators1(&_Staking.CallOpts)
}

// MinimumNumValidators1 is a free data retrieval call binding the contract method 0x714ff425.
//
// Solidity: function minimumNumValidators() view returns(uint256)
func (_Staking *StakingCallerSession) MinimumNumValidators1() (*big.Int, error) {
	return _Staking.Contract.MinimumNumValidators1(&_Staking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingSession) Owner() (common.Address, error) {
	return _Staking.Contract.Owner(&_Staking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingCallerSession) Owner() (common.Address, error) {
	return _Staking.Contract.Owner(&_Staking.CallOpts)
}

// StakedAmount1 is a free data retrieval call binding the contract method 0x373d6132.
//
// Solidity: function stakedAmount() view returns(uint256)
func (_Staking *StakingCaller) StakedAmount1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "stakedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedAmount1 is a free data retrieval call binding the contract method 0x373d6132.
//
// Solidity: function stakedAmount() view returns(uint256)
func (_Staking *StakingSession) StakedAmount1() (*big.Int, error) {
	return _Staking.Contract.StakedAmount1(&_Staking.CallOpts)
}

// StakedAmount1 is a free data retrieval call binding the contract method 0x373d6132.
//
// Solidity: function stakedAmount() view returns(uint256)
func (_Staking *StakingCallerSession) StakedAmount1() (*big.Int, error) {
	return _Staking.Contract.StakedAmount1(&_Staking.CallOpts)
}

// Validators is a free data retrieval call binding the contract method 0xca1e7819.
//
// Solidity: function validators() view returns(address[])
func (_Staking *StakingCaller) Validators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "validators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// Validators is a free data retrieval call binding the contract method 0xca1e7819.
//
// Solidity: function validators() view returns(address[])
func (_Staking *StakingSession) Validators() ([]common.Address, error) {
	return _Staking.Contract.Validators(&_Staking.CallOpts)
}

// Validators is a free data retrieval call binding the contract method 0xca1e7819.
//
// Solidity: function validators() view returns(address[])
func (_Staking *StakingCallerSession) Validators() ([]common.Address, error) {
	return _Staking.Contract.Validators(&_Staking.CallOpts)
}

// Ban is a paid mutator transaction binding the contract method 0xb0f572ee.
//
// Solidity: function ban(address badValidator, uint64 validatorsSubsetSize) returns()
func (_Staking *StakingTransactor) Ban(opts *bind.TransactOpts, badValidator common.Address, validatorsSubsetSize uint64) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "ban", badValidator, validatorsSubsetSize)
}

// Ban is a paid mutator transaction binding the contract method 0xb0f572ee.
//
// Solidity: function ban(address badValidator, uint64 validatorsSubsetSize) returns()
func (_Staking *StakingSession) Ban(badValidator common.Address, validatorsSubsetSize uint64) (*types.Transaction, error) {
	return _Staking.Contract.Ban(&_Staking.TransactOpts, badValidator, validatorsSubsetSize)
}

// Ban is a paid mutator transaction binding the contract method 0xb0f572ee.
//
// Solidity: function ban(address badValidator, uint64 validatorsSubsetSize) returns()
func (_Staking *StakingTransactorSession) Ban(badValidator common.Address, validatorsSubsetSize uint64) (*types.Transaction, error) {
	return _Staking.Contract.Ban(&_Staking.TransactOpts, badValidator, validatorsSubsetSize)
}

// BurnGrantInitial is a paid mutator transaction binding the contract method 0x96f58ffa.
//
// Solidity: function burnGrantInitial(address staker) returns()
func (_Staking *StakingTransactor) BurnGrantInitial(opts *bind.TransactOpts, staker common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "burnGrantInitial", staker)
}

// BurnGrantInitial is a paid mutator transaction binding the contract method 0x96f58ffa.
//
// Solidity: function burnGrantInitial(address staker) returns()
func (_Staking *StakingSession) BurnGrantInitial(staker common.Address) (*types.Transaction, error) {
	return _Staking.Contract.BurnGrantInitial(&_Staking.TransactOpts, staker)
}

// BurnGrantInitial is a paid mutator transaction binding the contract method 0x96f58ffa.
//
// Solidity: function burnGrantInitial(address staker) returns()
func (_Staking *StakingTransactorSession) BurnGrantInitial(staker common.Address) (*types.Transaction, error) {
	return _Staking.Contract.BurnGrantInitial(&_Staking.TransactOpts, staker)
}

// CheckStateValidators is a paid mutator transaction binding the contract method 0x19fd547d.
//
// Solidity: function checkStateValidators() returns()
func (_Staking *StakingTransactor) CheckStateValidators(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "checkStateValidators")
}

// CheckStateValidators is a paid mutator transaction binding the contract method 0x19fd547d.
//
// Solidity: function checkStateValidators() returns()
func (_Staking *StakingSession) CheckStateValidators() (*types.Transaction, error) {
	return _Staking.Contract.CheckStateValidators(&_Staking.TransactOpts)
}

// CheckStateValidators is a paid mutator transaction binding the contract method 0x19fd547d.
//
// Solidity: function checkStateValidators() returns()
func (_Staking *StakingTransactorSession) CheckStateValidators() (*types.Transaction, error) {
	return _Staking.Contract.CheckStateValidators(&_Staking.TransactOpts)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address signer) payable returns()
func (_Staking *StakingTransactor) Delegate(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "delegate", signer)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address signer) payable returns()
func (_Staking *StakingSession) Delegate(signer common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Delegate(&_Staking.TransactOpts, signer)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address signer) payable returns()
func (_Staking *StakingTransactorSession) Delegate(signer common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Delegate(&_Staking.TransactOpts, signer)
}

// DelegatorUnstake is a paid mutator transaction binding the contract method 0x9e723b63.
//
// Solidity: function delegatorUnstake(address signer) returns()
func (_Staking *StakingTransactor) DelegatorUnstake(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "delegatorUnstake", signer)
}

// DelegatorUnstake is a paid mutator transaction binding the contract method 0x9e723b63.
//
// Solidity: function delegatorUnstake(address signer) returns()
func (_Staking *StakingSession) DelegatorUnstake(signer common.Address) (*types.Transaction, error) {
	return _Staking.Contract.DelegatorUnstake(&_Staking.TransactOpts, signer)
}

// DelegatorUnstake is a paid mutator transaction binding the contract method 0x9e723b63.
//
// Solidity: function delegatorUnstake(address signer) returns()
func (_Staking *StakingTransactorSession) DelegatorUnstake(signer common.Address) (*types.Transaction, error) {
	return _Staking.Contract.DelegatorUnstake(&_Staking.TransactOpts, signer)
}

// ReclaimUnusedGrant is a paid mutator transaction binding the contract method 0xa5dbc130.
//
// Solidity: function reclaimUnusedGrant(address validator) returns()
func (_Staking *StakingTransactor) ReclaimUnusedGrant(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "reclaimUnusedGrant", validator)
}

// ReclaimUnusedGrant is a paid mutator transaction binding the contract method 0xa5dbc130.
//
// Solidity: function reclaimUnusedGrant(address validator) returns()
func (_Staking *StakingSession) ReclaimUnusedGrant(validator common.Address) (*types.Transaction, error) {
	return _Staking.Contract.ReclaimUnusedGrant(&_Staking.TransactOpts, validator)
}

// ReclaimUnusedGrant is a paid mutator transaction binding the contract method 0xa5dbc130.
//
// Solidity: function reclaimUnusedGrant(address validator) returns()
func (_Staking *StakingTransactorSession) ReclaimUnusedGrant(validator common.Address) (*types.Transaction, error) {
	return _Staking.Contract.ReclaimUnusedGrant(&_Staking.TransactOpts, validator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingSession) RenounceOwnership() (*types.Transaction, error) {
	return _Staking.Contract.RenounceOwnership(&_Staking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Staking.Contract.RenounceOwnership(&_Staking.TransactOpts)
}

// ResetValidatorsSubsetTimeout is a paid mutator transaction binding the contract method 0xba1b2125.
//
// Solidity: function resetValidatorsSubsetTimeout() returns()
func (_Staking *StakingTransactor) ResetValidatorsSubsetTimeout(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "resetValidatorsSubsetTimeout")
}

// ResetValidatorsSubsetTimeout is a paid mutator transaction binding the contract method 0xba1b2125.
//
// Solidity: function resetValidatorsSubsetTimeout() returns()
func (_Staking *StakingSession) ResetValidatorsSubsetTimeout() (*types.Transaction, error) {
	return _Staking.Contract.ResetValidatorsSubsetTimeout(&_Staking.TransactOpts)
}

// ResetValidatorsSubsetTimeout is a paid mutator transaction binding the contract method 0xba1b2125.
//
// Solidity: function resetValidatorsSubsetTimeout() returns()
func (_Staking *StakingTransactorSession) ResetValidatorsSubsetTimeout() (*types.Transaction, error) {
	return _Staking.Contract.ResetValidatorsSubsetTimeout(&_Staking.TransactOpts)
}

// SetGrantContract is a paid mutator transaction binding the contract method 0x22d46852.
//
// Solidity: function setGrantContract(address _contract) returns()
func (_Staking *StakingTransactor) SetGrantContract(opts *bind.TransactOpts, _contract common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setGrantContract", _contract)
}

// SetGrantContract is a paid mutator transaction binding the contract method 0x22d46852.
//
// Solidity: function setGrantContract(address _contract) returns()
func (_Staking *StakingSession) SetGrantContract(_contract common.Address) (*types.Transaction, error) {
	return _Staking.Contract.SetGrantContract(&_Staking.TransactOpts, _contract)
}

// SetGrantContract is a paid mutator transaction binding the contract method 0x22d46852.
//
// Solidity: function setGrantContract(address _contract) returns()
func (_Staking *StakingTransactorSession) SetGrantContract(_contract common.Address) (*types.Transaction, error) {
	return _Staking.Contract.SetGrantContract(&_Staking.TransactOpts, _contract)
}

// Stake is a paid mutator transaction binding the contract method 0x46f45b8d.
//
// Solidity: function stake(string name) payable returns()
func (_Staking *StakingTransactor) Stake(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "stake", name)
}

// Stake is a paid mutator transaction binding the contract method 0x46f45b8d.
//
// Solidity: function stake(string name) payable returns()
func (_Staking *StakingSession) Stake(name string) (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts, name)
}

// Stake is a paid mutator transaction binding the contract method 0x46f45b8d.
//
// Solidity: function stake(string name) payable returns()
func (_Staking *StakingTransactorSession) Stake(name string) (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts, name)
}

// Stake0 is a paid mutator transaction binding the contract method 0x4fb856e5.
//
// Solidity: function stake(address signer, string name) payable returns()
func (_Staking *StakingTransactor) Stake0(opts *bind.TransactOpts, signer common.Address, name string) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "stake0", signer, name)
}

// Stake0 is a paid mutator transaction binding the contract method 0x4fb856e5.
//
// Solidity: function stake(address signer, string name) payable returns()
func (_Staking *StakingSession) Stake0(signer common.Address, name string) (*types.Transaction, error) {
	return _Staking.Contract.Stake0(&_Staking.TransactOpts, signer, name)
}

// Stake0 is a paid mutator transaction binding the contract method 0x4fb856e5.
//
// Solidity: function stake(address signer, string name) payable returns()
func (_Staking *StakingTransactorSession) Stake0(signer common.Address, name string) (*types.Transaction, error) {
	return _Staking.Contract.Stake0(&_Staking.TransactOpts, signer, name)
}

// StakeGrantContract is a paid mutator transaction binding the contract method 0xe70c91b5.
//
// Solidity: function stakeGrantContract(address staker, string name) payable returns()
func (_Staking *StakingTransactor) StakeGrantContract(opts *bind.TransactOpts, staker common.Address, name string) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "stakeGrantContract", staker, name)
}

// StakeGrantContract is a paid mutator transaction binding the contract method 0xe70c91b5.
//
// Solidity: function stakeGrantContract(address staker, string name) payable returns()
func (_Staking *StakingSession) StakeGrantContract(staker common.Address, name string) (*types.Transaction, error) {
	return _Staking.Contract.StakeGrantContract(&_Staking.TransactOpts, staker, name)
}

// StakeGrantContract is a paid mutator transaction binding the contract method 0xe70c91b5.
//
// Solidity: function stakeGrantContract(address staker, string name) payable returns()
func (_Staking *StakingTransactorSession) StakeGrantContract(staker common.Address, name string) (*types.Transaction, error) {
	return _Staking.Contract.StakeGrantContract(&_Staking.TransactOpts, staker, name)
}

// StakeSignerGrantContract is a paid mutator transaction binding the contract method 0xf3a382f2.
//
// Solidity: function stakeSignerGrantContract(address staker, address signer, string name) payable returns()
func (_Staking *StakingTransactor) StakeSignerGrantContract(opts *bind.TransactOpts, staker common.Address, signer common.Address, name string) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "stakeSignerGrantContract", staker, signer, name)
}

// StakeSignerGrantContract is a paid mutator transaction binding the contract method 0xf3a382f2.
//
// Solidity: function stakeSignerGrantContract(address staker, address signer, string name) payable returns()
func (_Staking *StakingSession) StakeSignerGrantContract(staker common.Address, signer common.Address, name string) (*types.Transaction, error) {
	return _Staking.Contract.StakeSignerGrantContract(&_Staking.TransactOpts, staker, signer, name)
}

// StakeSignerGrantContract is a paid mutator transaction binding the contract method 0xf3a382f2.
//
// Solidity: function stakeSignerGrantContract(address staker, address signer, string name) payable returns()
func (_Staking *StakingTransactorSession) StakeSignerGrantContract(staker common.Address, signer common.Address, name string) (*types.Transaction, error) {
	return _Staking.Contract.StakeSignerGrantContract(&_Staking.TransactOpts, staker, signer, name)
}

// Suspend is a paid mutator transaction binding the contract method 0x5bb4f998.
//
// Solidity: function suspend(uint256 indexValidator, address badValidator, uint8 count) returns()
func (_Staking *StakingTransactor) Suspend(opts *bind.TransactOpts, indexValidator *big.Int, badValidator common.Address, count uint8) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "suspend", indexValidator, badValidator, count)
}

// Suspend is a paid mutator transaction binding the contract method 0x5bb4f998.
//
// Solidity: function suspend(uint256 indexValidator, address badValidator, uint8 count) returns()
func (_Staking *StakingSession) Suspend(indexValidator *big.Int, badValidator common.Address, count uint8) (*types.Transaction, error) {
	return _Staking.Contract.Suspend(&_Staking.TransactOpts, indexValidator, badValidator, count)
}

// Suspend is a paid mutator transaction binding the contract method 0x5bb4f998.
//
// Solidity: function suspend(uint256 indexValidator, address badValidator, uint8 count) returns()
func (_Staking *StakingTransactorSession) Suspend(indexValidator *big.Int, badValidator common.Address, count uint8) (*types.Transaction, error) {
	return _Staking.Contract.Suspend(&_Staking.TransactOpts, indexValidator, badValidator, count)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.TransferOwnership(&_Staking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.TransferOwnership(&_Staking.TransactOpts, newOwner)
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() returns()
func (_Staking *StakingTransactor) Unstake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "unstake")
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() returns()
func (_Staking *StakingSession) Unstake() (*types.Transaction, error) {
	return _Staking.Contract.Unstake(&_Staking.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() returns()
func (_Staking *StakingTransactorSession) Unstake() (*types.Transaction, error) {
	return _Staking.Contract.Unstake(&_Staking.TransactOpts)
}

// Warning is a paid mutator transaction binding the contract method 0xb788e44e.
//
// Solidity: function warning(address badValidator) returns()
func (_Staking *StakingTransactor) Warning(opts *bind.TransactOpts, badValidator common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "warning", badValidator)
}

// Warning is a paid mutator transaction binding the contract method 0xb788e44e.
//
// Solidity: function warning(address badValidator) returns()
func (_Staking *StakingSession) Warning(badValidator common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Warning(&_Staking.TransactOpts, badValidator)
}

// Warning is a paid mutator transaction binding the contract method 0xb788e44e.
//
// Solidity: function warning(address badValidator) returns()
func (_Staking *StakingTransactorSession) Warning(badValidator common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Warning(&_Staking.TransactOpts, badValidator)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Staking *StakingTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Staking *StakingSession) Receive() (*types.Transaction, error) {
	return _Staking.Contract.Receive(&_Staking.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Staking *StakingTransactorSession) Receive() (*types.Transaction, error) {
	return _Staking.Contract.Receive(&_Staking.TransactOpts)
}

// StakingBanIterator is returned from FilterBan and is used to iterate over the raw logs and unpacked data for Ban events raised by the Staking contract.
type StakingBanIterator struct {
	Event *StakingBan // Event containing the contract specifics and raw log

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
func (it *StakingBanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingBan)
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
		it.Event = new(StakingBan)
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
func (it *StakingBanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingBanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingBan represents a Ban event raised by the Staking contract.
type StakingBan struct {
	Account       common.Address
	BlockNumber   *big.Int
	TimeLock      *big.Int
	TokenSlashing *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBan is a free log retrieval operation binding the contract event 0x5b669607b17f631e17b24518702f889e404e09aa6528fd31e931ef4d2c53ee5e.
//
// Solidity: event Ban(address indexed account, uint256 blockNumber, uint256 timeLock, uint256 tokenSlashing)
func (_Staking *StakingFilterer) FilterBan(opts *bind.FilterOpts, account []common.Address) (*StakingBanIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Ban", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakingBanIterator{contract: _Staking.contract, event: "Ban", logs: logs, sub: sub}, nil
}

// WatchBan is a free log subscription operation binding the contract event 0x5b669607b17f631e17b24518702f889e404e09aa6528fd31e931ef4d2c53ee5e.
//
// Solidity: event Ban(address indexed account, uint256 blockNumber, uint256 timeLock, uint256 tokenSlashing)
func (_Staking *StakingFilterer) WatchBan(opts *bind.WatchOpts, sink chan<- *StakingBan, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Ban", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingBan)
				if err := _Staking.contract.UnpackLog(event, "Ban", log); err != nil {
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

// ParseBan is a log parse operation binding the contract event 0x5b669607b17f631e17b24518702f889e404e09aa6528fd31e931ef4d2c53ee5e.
//
// Solidity: event Ban(address indexed account, uint256 blockNumber, uint256 timeLock, uint256 tokenSlashing)
func (_Staking *StakingFilterer) ParseBan(log types.Log) (*StakingBan, error) {
	event := new(StakingBan)
	if err := _Staking.contract.UnpackLog(event, "Ban", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingBurnGrantInitialIterator is returned from FilterBurnGrantInitial and is used to iterate over the raw logs and unpacked data for BurnGrantInitial events raised by the Staking contract.
type StakingBurnGrantInitialIterator struct {
	Event *StakingBurnGrantInitial // Event containing the contract specifics and raw log

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
func (it *StakingBurnGrantInitialIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingBurnGrantInitial)
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
		it.Event = new(StakingBurnGrantInitial)
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
func (it *StakingBurnGrantInitialIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingBurnGrantInitialIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingBurnGrantInitial represents a BurnGrantInitial event raised by the Staking contract.
type StakingBurnGrantInitial struct {
	Staker common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurnGrantInitial is a free log retrieval operation binding the contract event 0xb67ea9e632aa74e96032afed3f64179247898d3dfa8d9aef85057d36e9112d46.
//
// Solidity: event BurnGrantInitial(address staker)
func (_Staking *StakingFilterer) FilterBurnGrantInitial(opts *bind.FilterOpts) (*StakingBurnGrantInitialIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "BurnGrantInitial")
	if err != nil {
		return nil, err
	}
	return &StakingBurnGrantInitialIterator{contract: _Staking.contract, event: "BurnGrantInitial", logs: logs, sub: sub}, nil
}

// WatchBurnGrantInitial is a free log subscription operation binding the contract event 0xb67ea9e632aa74e96032afed3f64179247898d3dfa8d9aef85057d36e9112d46.
//
// Solidity: event BurnGrantInitial(address staker)
func (_Staking *StakingFilterer) WatchBurnGrantInitial(opts *bind.WatchOpts, sink chan<- *StakingBurnGrantInitial) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "BurnGrantInitial")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingBurnGrantInitial)
				if err := _Staking.contract.UnpackLog(event, "BurnGrantInitial", log); err != nil {
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

// ParseBurnGrantInitial is a log parse operation binding the contract event 0xb67ea9e632aa74e96032afed3f64179247898d3dfa8d9aef85057d36e9112d46.
//
// Solidity: event BurnGrantInitial(address staker)
func (_Staking *StakingFilterer) ParseBurnGrantInitial(log types.Log) (*StakingBurnGrantInitial, error) {
	event := new(StakingBurnGrantInitial)
	if err := _Staking.contract.UnpackLog(event, "BurnGrantInitial", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingDelegatorStakedIterator is returned from FilterDelegatorStaked and is used to iterate over the raw logs and unpacked data for DelegatorStaked events raised by the Staking contract.
type StakingDelegatorStakedIterator struct {
	Event *StakingDelegatorStaked // Event containing the contract specifics and raw log

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
func (it *StakingDelegatorStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingDelegatorStaked)
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
		it.Event = new(StakingDelegatorStaked)
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
func (it *StakingDelegatorStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingDelegatorStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingDelegatorStaked represents a DelegatorStaked event raised by the Staking contract.
type StakingDelegatorStaked struct {
	Account   common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegatorStaked is a free log retrieval operation binding the contract event 0x1ed8ad98d928651a8bc3999999b718383931f4595fcd2e1efd2de972fa8cdaea.
//
// Solidity: event DelegatorStaked(address indexed account, address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) FilterDelegatorStaked(opts *bind.FilterOpts, account []common.Address, validator []common.Address) (*StakingDelegatorStakedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "DelegatorStaked", accountRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingDelegatorStakedIterator{contract: _Staking.contract, event: "DelegatorStaked", logs: logs, sub: sub}, nil
}

// WatchDelegatorStaked is a free log subscription operation binding the contract event 0x1ed8ad98d928651a8bc3999999b718383931f4595fcd2e1efd2de972fa8cdaea.
//
// Solidity: event DelegatorStaked(address indexed account, address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) WatchDelegatorStaked(opts *bind.WatchOpts, sink chan<- *StakingDelegatorStaked, account []common.Address, validator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "DelegatorStaked", accountRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingDelegatorStaked)
				if err := _Staking.contract.UnpackLog(event, "DelegatorStaked", log); err != nil {
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

// ParseDelegatorStaked is a log parse operation binding the contract event 0x1ed8ad98d928651a8bc3999999b718383931f4595fcd2e1efd2de972fa8cdaea.
//
// Solidity: event DelegatorStaked(address indexed account, address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) ParseDelegatorStaked(log types.Log) (*StakingDelegatorStaked, error) {
	event := new(StakingDelegatorStaked)
	if err := _Staking.contract.UnpackLog(event, "DelegatorStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingDelegatorUnstakedIterator is returned from FilterDelegatorUnstaked and is used to iterate over the raw logs and unpacked data for DelegatorUnstaked events raised by the Staking contract.
type StakingDelegatorUnstakedIterator struct {
	Event *StakingDelegatorUnstaked // Event containing the contract specifics and raw log

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
func (it *StakingDelegatorUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingDelegatorUnstaked)
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
		it.Event = new(StakingDelegatorUnstaked)
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
func (it *StakingDelegatorUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingDelegatorUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingDelegatorUnstaked represents a DelegatorUnstaked event raised by the Staking contract.
type StakingDelegatorUnstaked struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDelegatorUnstaked is a free log retrieval operation binding the contract event 0xce2e26dc37234e172764d0a8a2c04d2e6fb46c02f3ca6ca5135d6f9dcd2b67a9.
//
// Solidity: event DelegatorUnstaked(address indexed account, uint256 amount)
func (_Staking *StakingFilterer) FilterDelegatorUnstaked(opts *bind.FilterOpts, account []common.Address) (*StakingDelegatorUnstakedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "DelegatorUnstaked", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakingDelegatorUnstakedIterator{contract: _Staking.contract, event: "DelegatorUnstaked", logs: logs, sub: sub}, nil
}

// WatchDelegatorUnstaked is a free log subscription operation binding the contract event 0xce2e26dc37234e172764d0a8a2c04d2e6fb46c02f3ca6ca5135d6f9dcd2b67a9.
//
// Solidity: event DelegatorUnstaked(address indexed account, uint256 amount)
func (_Staking *StakingFilterer) WatchDelegatorUnstaked(opts *bind.WatchOpts, sink chan<- *StakingDelegatorUnstaked, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "DelegatorUnstaked", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingDelegatorUnstaked)
				if err := _Staking.contract.UnpackLog(event, "DelegatorUnstaked", log); err != nil {
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

// ParseDelegatorUnstaked is a log parse operation binding the contract event 0xce2e26dc37234e172764d0a8a2c04d2e6fb46c02f3ca6ca5135d6f9dcd2b67a9.
//
// Solidity: event DelegatorUnstaked(address indexed account, uint256 amount)
func (_Staking *StakingFilterer) ParseDelegatorUnstaked(log types.Log) (*StakingDelegatorUnstaked, error) {
	event := new(StakingDelegatorUnstaked)
	if err := _Staking.contract.UnpackLog(event, "DelegatorUnstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingGrantContractChangedIterator is returned from FilterGrantContractChanged and is used to iterate over the raw logs and unpacked data for GrantContractChanged events raised by the Staking contract.
type StakingGrantContractChangedIterator struct {
	Event *StakingGrantContractChanged // Event containing the contract specifics and raw log

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
func (it *StakingGrantContractChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingGrantContractChanged)
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
		it.Event = new(StakingGrantContractChanged)
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
func (it *StakingGrantContractChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingGrantContractChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingGrantContractChanged represents a GrantContractChanged event raised by the Staking contract.
type StakingGrantContractChanged struct {
	GrantContract common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterGrantContractChanged is a free log retrieval operation binding the contract event 0x8eef4d58d8cb1058a01a95ce964ea832ace1809b77a07c588481f74a10d37723.
//
// Solidity: event GrantContractChanged(address indexed _grantContract)
func (_Staking *StakingFilterer) FilterGrantContractChanged(opts *bind.FilterOpts, _grantContract []common.Address) (*StakingGrantContractChangedIterator, error) {

	var _grantContractRule []interface{}
	for _, _grantContractItem := range _grantContract {
		_grantContractRule = append(_grantContractRule, _grantContractItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "GrantContractChanged", _grantContractRule)
	if err != nil {
		return nil, err
	}
	return &StakingGrantContractChangedIterator{contract: _Staking.contract, event: "GrantContractChanged", logs: logs, sub: sub}, nil
}

// WatchGrantContractChanged is a free log subscription operation binding the contract event 0x8eef4d58d8cb1058a01a95ce964ea832ace1809b77a07c588481f74a10d37723.
//
// Solidity: event GrantContractChanged(address indexed _grantContract)
func (_Staking *StakingFilterer) WatchGrantContractChanged(opts *bind.WatchOpts, sink chan<- *StakingGrantContractChanged, _grantContract []common.Address) (event.Subscription, error) {

	var _grantContractRule []interface{}
	for _, _grantContractItem := range _grantContract {
		_grantContractRule = append(_grantContractRule, _grantContractItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "GrantContractChanged", _grantContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingGrantContractChanged)
				if err := _Staking.contract.UnpackLog(event, "GrantContractChanged", log); err != nil {
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

// ParseGrantContractChanged is a log parse operation binding the contract event 0x8eef4d58d8cb1058a01a95ce964ea832ace1809b77a07c588481f74a10d37723.
//
// Solidity: event GrantContractChanged(address indexed _grantContract)
func (_Staking *StakingFilterer) ParseGrantContractChanged(log types.Log) (*StakingGrantContractChanged, error) {
	event := new(StakingGrantContractChanged)
	if err := _Staking.contract.UnpackLog(event, "GrantContractChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Staking contract.
type StakingOwnershipTransferredIterator struct {
	Event *StakingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingOwnershipTransferred)
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
		it.Event = new(StakingOwnershipTransferred)
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
func (it *StakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingOwnershipTransferred represents a OwnershipTransferred event raised by the Staking contract.
type StakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingOwnershipTransferredIterator{contract: _Staking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingOwnershipTransferred)
				if err := _Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Staking *StakingFilterer) ParseOwnershipTransferred(log types.Log) (*StakingOwnershipTransferred, error) {
	event := new(StakingOwnershipTransferred)
	if err := _Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingReclaimGrantIterator is returned from FilterReclaimGrant and is used to iterate over the raw logs and unpacked data for ReclaimGrant events raised by the Staking contract.
type StakingReclaimGrantIterator struct {
	Event *StakingReclaimGrant // Event containing the contract specifics and raw log

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
func (it *StakingReclaimGrantIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingReclaimGrant)
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
		it.Event = new(StakingReclaimGrant)
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
func (it *StakingReclaimGrantIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingReclaimGrantIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingReclaimGrant represents a ReclaimGrant event raised by the Staking contract.
type StakingReclaimGrant struct {
	Staker       common.Address
	EarnedAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterReclaimGrant is a free log retrieval operation binding the contract event 0x6782e2944a228b3021a467b1afb406079eea2c32d259933346a946601e56d2ae.
//
// Solidity: event ReclaimGrant(address staker, uint256 earnedAmount)
func (_Staking *StakingFilterer) FilterReclaimGrant(opts *bind.FilterOpts) (*StakingReclaimGrantIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ReclaimGrant")
	if err != nil {
		return nil, err
	}
	return &StakingReclaimGrantIterator{contract: _Staking.contract, event: "ReclaimGrant", logs: logs, sub: sub}, nil
}

// WatchReclaimGrant is a free log subscription operation binding the contract event 0x6782e2944a228b3021a467b1afb406079eea2c32d259933346a946601e56d2ae.
//
// Solidity: event ReclaimGrant(address staker, uint256 earnedAmount)
func (_Staking *StakingFilterer) WatchReclaimGrant(opts *bind.WatchOpts, sink chan<- *StakingReclaimGrant) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ReclaimGrant")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingReclaimGrant)
				if err := _Staking.contract.UnpackLog(event, "ReclaimGrant", log); err != nil {
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

// ParseReclaimGrant is a log parse operation binding the contract event 0x6782e2944a228b3021a467b1afb406079eea2c32d259933346a946601e56d2ae.
//
// Solidity: event ReclaimGrant(address staker, uint256 earnedAmount)
func (_Staking *StakingFilterer) ParseReclaimGrant(log types.Log) (*StakingReclaimGrant, error) {
	event := new(StakingReclaimGrant)
	if err := _Staking.contract.UnpackLog(event, "ReclaimGrant", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Staking contract.
type StakingStakedIterator struct {
	Event *StakingStaked // Event containing the contract specifics and raw log

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
func (it *StakingStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStaked)
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
		it.Event = new(StakingStaked)
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
func (it *StakingStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStaked represents a Staked event raised by the Staking contract.
type StakingStaked struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed account, uint256 amount)
func (_Staking *StakingFilterer) FilterStaked(opts *bind.FilterOpts, account []common.Address) (*StakingStakedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Staked", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakingStakedIterator{contract: _Staking.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed account, uint256 amount)
func (_Staking *StakingFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *StakingStaked, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Staked", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStaked)
				if err := _Staking.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed account, uint256 amount)
func (_Staking *StakingFilterer) ParseStaked(log types.Log) (*StakingStaked, error) {
	event := new(StakingStaked)
	if err := _Staking.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingSuspendedIterator is returned from FilterSuspended and is used to iterate over the raw logs and unpacked data for Suspended events raised by the Staking contract.
type StakingSuspendedIterator struct {
	Event *StakingSuspended // Event containing the contract specifics and raw log

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
func (it *StakingSuspendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingSuspended)
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
		it.Event = new(StakingSuspended)
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
func (it *StakingSuspendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingSuspendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingSuspended represents a Suspended event raised by the Staking contract.
type StakingSuspended struct {
	Account        common.Address
	BlockNumber    *big.Int
	TimeLock       *big.Int
	SuspendCounter uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSuspended is a free log retrieval operation binding the contract event 0xf833a6e5fb833592d2d6a2f115227e397239985e961011fd1cd3cec336915257.
//
// Solidity: event Suspended(address indexed account, uint256 blockNumber, uint256 timeLock, uint64 suspendCounter)
func (_Staking *StakingFilterer) FilterSuspended(opts *bind.FilterOpts, account []common.Address) (*StakingSuspendedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Suspended", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakingSuspendedIterator{contract: _Staking.contract, event: "Suspended", logs: logs, sub: sub}, nil
}

// WatchSuspended is a free log subscription operation binding the contract event 0xf833a6e5fb833592d2d6a2f115227e397239985e961011fd1cd3cec336915257.
//
// Solidity: event Suspended(address indexed account, uint256 blockNumber, uint256 timeLock, uint64 suspendCounter)
func (_Staking *StakingFilterer) WatchSuspended(opts *bind.WatchOpts, sink chan<- *StakingSuspended, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Suspended", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingSuspended)
				if err := _Staking.contract.UnpackLog(event, "Suspended", log); err != nil {
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

// ParseSuspended is a log parse operation binding the contract event 0xf833a6e5fb833592d2d6a2f115227e397239985e961011fd1cd3cec336915257.
//
// Solidity: event Suspended(address indexed account, uint256 blockNumber, uint256 timeLock, uint64 suspendCounter)
func (_Staking *StakingFilterer) ParseSuspended(log types.Log) (*StakingSuspended, error) {
	event := new(StakingSuspended)
	if err := _Staking.contract.UnpackLog(event, "Suspended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the Staking contract.
type StakingUnstakedIterator struct {
	Event *StakingUnstaked // Event containing the contract specifics and raw log

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
func (it *StakingUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUnstaked)
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
		it.Event = new(StakingUnstaked)
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
func (it *StakingUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUnstaked represents a Unstaked event raised by the Staking contract.
type StakingUnstaked struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed account, uint256 amount)
func (_Staking *StakingFilterer) FilterUnstaked(opts *bind.FilterOpts, account []common.Address) (*StakingUnstakedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Unstaked", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakingUnstakedIterator{contract: _Staking.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed account, uint256 amount)
func (_Staking *StakingFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *StakingUnstaked, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Unstaked", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUnstaked)
				if err := _Staking.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// ParseUnstaked is a log parse operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed account, uint256 amount)
func (_Staking *StakingFilterer) ParseUnstaked(log types.Log) (*StakingUnstaked, error) {
	event := new(StakingUnstaked)
	if err := _Staking.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWarningIterator is returned from FilterWarning and is used to iterate over the raw logs and unpacked data for Warning events raised by the Staking contract.
type StakingWarningIterator struct {
	Event *StakingWarning // Event containing the contract specifics and raw log

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
func (it *StakingWarningIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWarning)
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
		it.Event = new(StakingWarning)
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
func (it *StakingWarningIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWarningIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWarning represents a Warning event raised by the Staking contract.
type StakingWarning struct {
	Account     common.Address
	BlockNumber *big.Int
	Message     string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWarning is a free log retrieval operation binding the contract event 0x4d381e135b58a2c3686763368fd0d7750138ab8a9a0e404a6c09e03f24100368.
//
// Solidity: event Warning(address indexed account, uint256 blockNumber, string message)
func (_Staking *StakingFilterer) FilterWarning(opts *bind.FilterOpts, account []common.Address) (*StakingWarningIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Warning", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakingWarningIterator{contract: _Staking.contract, event: "Warning", logs: logs, sub: sub}, nil
}

// WatchWarning is a free log subscription operation binding the contract event 0x4d381e135b58a2c3686763368fd0d7750138ab8a9a0e404a6c09e03f24100368.
//
// Solidity: event Warning(address indexed account, uint256 blockNumber, string message)
func (_Staking *StakingFilterer) WatchWarning(opts *bind.WatchOpts, sink chan<- *StakingWarning, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Warning", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWarning)
				if err := _Staking.contract.UnpackLog(event, "Warning", log); err != nil {
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

// ParseWarning is a log parse operation binding the contract event 0x4d381e135b58a2c3686763368fd0d7750138ab8a9a0e404a6c09e03f24100368.
//
// Solidity: event Warning(address indexed account, uint256 blockNumber, string message)
func (_Staking *StakingFilterer) ParseWarning(log types.Log) (*StakingWarning, error) {
	event := new(StakingWarning)
	if err := _Staking.contract.UnpackLog(event, "Warning", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
