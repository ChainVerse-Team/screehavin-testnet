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
	Addr          common.Address
	Amount        *big.Int
	State         uint8
	TimeLock      *big.Int
	UnstakeAmount *big.Int
}

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minNumValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNumValidators\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"grantContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accout\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"BLSPublicKeyRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeLock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalTokenSlashing\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalStakedSize\",\"type\":\"uint256\"}],\"name\":\"Ban\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"BurnGrantInitial\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelegatorStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelegatorUnstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_grantContract\",\"type\":\"address\"}],\"name\":\"GrantContractChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"earnedAmount\",\"type\":\"uint256\"}],\"name\":\"ReclaimGrant\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeLock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"suspendCounter\",\"type\":\"uint64\"}],\"name\":\"Suspended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"Warning\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BAN_DURATION\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COUNTER_SUSPEND\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DELEGATION_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GrantAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_VALIDATORSUBSET_SIZE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_VALIDATORSUBSET_SIZE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERCENTAGE_TOKEN_SLASHING\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUSPEND_DURATION\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_SLASHING_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_crewName\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"boss\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_grantContract\",\"outputs\":[{\"internalType\":\"contractGrantContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_maximumNumDelegators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_maximumNumValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_minimumNumDelegators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_minimumNumValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_stakedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_validators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"accountStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressGrantContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"addressToStakedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"badValidator\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"validatorsSubsetSize\",\"type\":\"uint64\"}],\"name\":\"ban\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"burnGrantInitial\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"canBecomeBoss\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkStateValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToUnstake\",\"type\":\"uint256\"}],\"name\":\"delegatorUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"getDelegatedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"getDelegatorsInfoOfValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumStaking.Status\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timeLock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeAmount\",\"type\":\"uint256\"}],\"internalType\":\"structStaking.AccountStakingInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"getStringLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getTotalDelegatorOfValidator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"getTotalDeligatedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorsStakeInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumStaking.Status\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timeLock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeAmount\",\"type\":\"uint256\"}],\"internalType\":\"structStaking.AccountStakingInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorsSubsetTimeout\",\"outputs\":[{\"internalType\":\"uint8[]\",\"name\":\"\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"isDelegator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isStaker\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maximumNumValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumNumValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"reclaimUnusedGrant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"blsPubKey\",\"type\":\"bytes\"}],\"name\":\"registerBLSPublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetValidatorsSubsetTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"setGrantContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"stakeGrantContract\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"stakeSignerGrantContract\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"indexValidator\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"badValidator\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"count\",\"type\":\"uint8\"}],\"name\":\"suspend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToUnstake\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorBLSPublicKeys\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"badValidator\",\"type\":\"address\"}],\"name\":\"warning\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040525f600f55610bb86010553480156200001a575f80fd5b5060405162009a1438038062009a148339818101604052810190620000409190620002a2565b62000060620000546200013d60201b60201c565b6200014460201b60201c565b81831115620000a6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200009d906200037f565b60405180910390fd5b82600881905550816009819055508060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050506200039f565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f80fd5b5f819050919050565b6200021d8162000209565b811462000228575f80fd5b50565b5f815190506200023b8162000212565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6200026c8262000241565b9050919050565b6200027e8162000260565b811462000289575f80fd5b50565b5f815190506200029c8162000273565b92915050565b5f805f60608486031215620002bc57620002bb62000205565b5b5f620002cb868287016200022b565b9350506020620002de868287016200022b565b9250506040620002f1868287016200028c565b9150509250925092565b5f82825260208201905092915050565b7f4d696e2076616c696461746f7273206e756d2063616e206e6f742062652067725f8201527f6561746572207468616e206d6178206e756d206f662076616c696461746f7273602082015250565b5f62000367604083620002fb565b915062000374826200030b565b604082019050919050565b5f6020820190508181035f830152620003988162000359565b9050919050565b61966780620003ad5f395ff3fe608060405260043610610379575f3560e01c806390b19627116101d0578063cb3d89d611610101578063e804fbf61161009f578063f6676a3a1161006e578063f6676a3a14610ce0578063f740ff8e14610d0a578063f90ecacc14610d34578063facd743b14610d70576103e7565b8063e804fbf614610c36578063ed546c8d14610c60578063f2fde38b14610c9c578063f3a382f214610cc4576103e7565b8063e069fe99116100db578063e069fe9914610b8a578063e387a7ed14610bc6578063e69af0c714610bf0578063e70c91b514610c1a576103e7565b8063cb3d89d614610b0e578063d94c111b14610b38578063de092ac414610b60576103e7565b8063b1f72b521161016e578063bb4c461e11610148578063bb4c461e14610a54578063c3a6ce2b14610a7e578063c795c07714610aba578063ca1e781914610ae4576103e7565b8063b1f72b52146109da578063b788e44e14610a16578063ba1b212514610a3e576103e7565b8063a157339c116101aa578063a157339c14610936578063a5dbc13014610960578063af6da36e14610988578063b0f572ee146109b2576103e7565b806390b19627146108a857806396f58ffa146108d2578063a0e11929146108fa576103e7565b80635bb4f998116102aa578063714ff425116102485780637df73e27116102225780637df73e27146107db5780637e840377146108175780637ebe444c146108415780638da5cb5b1461087e576103e7565b8063714ff42514610771578063715018a61461079b5780637a6eea37146107b1576103e7565b80635c8c4a07116102845780635c8c4a0714610693578063636592ed146106bd57806365c19af0146106f95780636f1e853314610735576103e7565b80635bb4f998146106255780635bb9fec91461064d5780635c19a95c14610677576103e7565b806336c4086f116103175780633c561f04116102f15780633c561f041461059b57806346f45b8d146105c557806348cb4d31146105e15780634fb856e514610609576103e7565b806336c4086f1461051d578063373d6132146105475780633c449ae714610571576103e7565b806322d468521161035357806322d46852146104555780632367f6b51461047d5780632e17de78146104b9578063306c019d146104e1576103e7565b806303306f0a146103eb57806315f9eb9b1461041557806319fd547d1461043f576103e7565b366103e75761039d3373ffffffffffffffffffffffffffffffffffffffff16610dac565b156103dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103d4906176b1565b60405180910390fd5b6103e5610dce565b005b5f80fd5b3480156103f6575f80fd5b506103ff611139565b60405161040c91906178a7565b60405180910390f35b348015610420575f80fd5b506104296114b8565b60405161043691906178d6565b60405180910390f35b34801561044a575f80fd5b506104536114c6565b005b348015610460575f80fd5b5061047b6004803603810190610476919061792a565b61189f565b005b348015610488575f80fd5b506104a3600480360381019061049e919061792a565b61196d565b6040516104b091906178d6565b60405180910390f35b3480156104c4575f80fd5b506104df60048036038101906104da919061797f565b6119b3565b005b3480156104ec575f80fd5b506105076004803603810190610502919061792a565b611aad565b60405161051491906178d6565b60405180910390f35b348015610528575f80fd5b50610531611ce0565b60405161053e91906178d6565b60405180910390f35b348015610552575f80fd5b5061055b611ce6565b60405161056891906178d6565b60405180910390f35b34801561057c575f80fd5b50610585611cef565b60405161059291906179d4565b60405180910390f35b3480156105a6575f80fd5b506105af611cfc565b6040516105bc9190617b32565b60405180910390f35b6105df60048036038101906105da9190617c7e565b611e9b565b005b3480156105ec575f80fd5b5061060760048036038101906106029190617cc5565b611f1e565b005b610623600480360381019061061e9190617d03565b61204a565b005b348015610630575f80fd5b5061064b60048036038101906106469190617d93565b6120cf565b005b348015610658575f80fd5b50610661612273565b60405161066e91906179d4565b60405180910390f35b610691600480360381019061068c919061792a565b612279565b005b34801561069e575f80fd5b506106a76122e4565b6040516106b491906179d4565b60405180910390f35b3480156106c8575f80fd5b506106e360048036038101906106de9190617de3565b6122e9565b6040516106f091906178d6565b60405180910390f35b348015610704575f80fd5b5061071f600480360381019061071a9190617c7e565b61236d565b60405161072c91906178d6565b60405180910390f35b348015610740575f80fd5b5061075b6004803603810190610756919061792a565b61237c565b6040516107689190617e3b565b60405180910390f35b34801561077c575f80fd5b5061078561238d565b60405161079291906178d6565b60405180910390f35b3480156107a6575f80fd5b506107af612396565b005b3480156107bc575f80fd5b506107c56123a9565b6040516107d291906179d4565b60405180910390f35b3480156107e6575f80fd5b5061080160048036038101906107fc919061792a565b6123b7565b60405161080e9190617e3b565b60405180910390f35b348015610822575f80fd5b5061082b6123c8565b6040516108389190617e76565b60405180910390f35b34801561084c575f80fd5b506108676004803603810190610862919061797f565b6123cd565b604051610875929190617ee0565b60405180910390f35b348015610889575f80fd5b506108926124a1565b60405161089f9190617f0e565b60405180910390f35b3480156108b3575f80fd5b506108bc6124c8565b6040516108c991906178d6565b60405180910390f35b3480156108dd575f80fd5b506108f860048036038101906108f3919061792a565b6124ce565b005b348015610905575f80fd5b50610920600480360381019061091b9190617de3565b612664565b60405161092d9190617e3b565b60405180910390f35b348015610941575f80fd5b5061094a6126f5565b60405161095791906179d4565b60405180910390f35b34801561096b575f80fd5b506109866004803603810190610981919061792a565b612703565b005b348015610993575f80fd5b5061099c612b80565b6040516109a991906178d6565b60405180910390f35b3480156109bd575f80fd5b506109d860048036038101906109d39190617f51565b612b86565b005b3480156109e5575f80fd5b50610a0060048036038101906109fb919061792a565b612c02565b604051610a0d91906178a7565b60405180910390f35b348015610a21575f80fd5b50610a3c6004803603810190610a37919061792a565b612f2c565b005b348015610a49575f80fd5b50610a52612feb565b005b348015610a5f575f80fd5b50610a686130cd565b604051610a7591906179d4565b60405180910390f35b348015610a89575f80fd5b50610aa46004803603810190610a9f919061792a565b6130d3565b604051610ab191906178d6565b60405180910390f35b348015610ac5575f80fd5b50610ace613119565b604051610adb91906178d6565b60405180910390f35b348015610aef575f80fd5b50610af861311f565b604051610b059190618037565b60405180910390f35b348015610b19575f80fd5b50610b226131aa565b604051610b2f91906180b2565b60405180910390f35b348015610b43575f80fd5b50610b5e6004803603810190610b599190618169565b6131cf565b005b348015610b6b575f80fd5b50610b7461326b565b604051610b819190617e76565b60405180910390f35b348015610b95575f80fd5b50610bb06004803603810190610bab919061792a565b613270565b604051610bbd9190617e3b565b60405180910390f35b348015610bd1575f80fd5b50610bda613332565b604051610be791906178d6565b60405180910390f35b348015610bfb575f80fd5b50610c04613338565b604051610c119190617f0e565b60405180910390f35b610c346004803603810190610c2f9190617d03565b61335d565b005b348015610c41575f80fd5b50610c4a6135c5565b604051610c5791906178d6565b60405180910390f35b348015610c6b575f80fd5b50610c866004803603810190610c81919061792a565b6135ce565b604051610c9391906178d6565b60405180910390f35b348015610ca7575f80fd5b50610cc26004803603810190610cbd919061792a565b613617565b005b610cde6004803603810190610cd991906181b0565b613699565b005b348015610ceb575f80fd5b50610cf4613d33565b604051610d0191906182d3565b60405180910390f35b348015610d15575f80fd5b50610d1e613e2f565b604051610d2b9190618302565b60405180910390f35b348015610d3f575f80fd5b50610d5a6004803603810190610d55919061797f565b613e34565b604051610d679190617f0e565b60405180910390f35b348015610d7b575f80fd5b50610d966004803603810190610d91919061792a565b613e6f565b604051610da39190617e3b565b60405180910390f35b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ede391fb336040518263ffffffff1660e01b8152600401610e289190617f0e565b602060405180830381865afa158015610e43573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e679190618345565b8015610eee5750600580811115610e8157610e8061774f565b5b60135f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff166005811115610ee057610edf61774f565b5b141580610eed57505f3414155b5b15610fcc5760025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663489fe625336040518263ffffffff1660e01b8152600401610f4d9190617f0e565b602060405180830381865afa158015610f68573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f8c9190618345565b610fcb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fc2906183e0565b60405180910390fd5b5b5f610fd633613ec1565b9150505f60135f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015490505f439050808210611063576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161105a90618448565b60405180910390fd5b3460075f8282546110749190618493565b925050819055503460055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546110c79190618493565b925050819055506110d783613fb9565b156110e6576110e5836140c3565b5b8273ffffffffffffffffffffffffffffffffffffffff167f9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d3460405161112c91906178d6565b60405180910390a2505050565b60605f60018054905067ffffffffffffffff81111561115b5761115a617b5a565b5b60405190808252806020026020018201604052801561119457816020015b6111816175ad565b8152602001906001900390816111795790505b5090505f5b6001805490508167ffffffffffffffff1610156114b0575f60135f60018467ffffffffffffffff16815481106111d2576111d16184c6565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015490505f60135f60018567ffffffffffffffff168154811061125b5761125a6184c6565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff1690505f60135f60018667ffffffffffffffff16815481106112f0576112ef6184c6565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206004015490505f6040518060a0016040528060018767ffffffffffffffff1681548110611381576113806184c6565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160055f60018967ffffffffffffffff16815481106113e5576113e46184c6565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054815260200184600581111561145f5761145e61774f565b5b815260200185815260200183815250905080868667ffffffffffffffff168151811061148e5761148d6184c6565b5b60200260200101819052505050505080806114a8906184f3565b915050611199565b508091505090565b691e11d5de63c17cc0000081565b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611534576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161152b90618592565b60405180910390fd5b5f4390505f5b6001805490508167ffffffffffffffff16101561189b575f60135f60018467ffffffffffffffff1681548110611573576115726184c6565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015490505f60135f60018567ffffffffffffffff16815481106115fc576115fb6184c6565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff169050600360058111156116835761168261774f565b5b8160058111156116965761169561774f565b5b14806116c65750600460058111156116b1576116b061774f565b5b8160058111156116c4576116c361774f565b5b145b1561180b5781841061180a575f60135f60018667ffffffffffffffff16815481106116f4576116f36184c6565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2060030181905550600560135f60018667ffffffffffffffff168154811061177f5761177e6184c6565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f6101000a81548160ff021916908360058111156118045761180361774f565b5b02179055505b5b6002600581111561181f5761181e61774f565b5b8160058111156118325761183161774f565b5b036118865761188560018467ffffffffffffffff1681548110611858576118576184c6565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16614330565b5b50508080611893906184f3565b91505061153a565b5050565b6118a7614b1a565b8060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167f8eef4d58d8cb1058a01a95ce964ea832ace1809b77a07c588481f74a10d3772360405160405180910390a250565b5f60055f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b6119d23373ffffffffffffffffffffffffffffffffffffffff16610dac565b15611a12576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a09906176b1565b60405180910390fd5b611a1b33614b98565b80611a6257505f60055f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054115b611aa1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a98906185fa565b60405180910390fd5b611aaa81614bea565b50565b5f8060125f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805490501115611cd7575f8060125f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805480602002602001604051908101604052809291908181526020018280548015611bb257602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611b69575b505050505090505f5b60125f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805490508167ffffffffffffffff161015611ccc5760115f838367ffffffffffffffff1681518110611c2d57611c2c6184c6565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205483611cb79190618493565b92508080611cc4906184f3565b915050611bbb565b508192505050611cdb565b5f90505b919050565b600f5481565b5f600754905090565b6802b5e3af16b188000081565b60605f60018054905090505f8167ffffffffffffffff811115611d2257611d21617b5a565b5b604051908082528060200260200182016040528015611d5557816020015b6060815260200190600190039081611d405790505b5090505f5b82811015611e9257600a5f60018381548110611d7957611d786184c6565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208054611de590618645565b80601f0160208091040260200160405190810160405280929190818152602001828054611e1190618645565b8015611e5c5780601f10611e3357610100808354040283529160200191611e5c565b820191905f5260205f20905b815481529060010190602001808311611e3f57829003601f168201915b5050505050828281518110611e7457611e736184c6565b5b60200260200101819052508080611e8a90618675565b915050611d5a565b50809250505090565b611eba3373ffffffffffffffffffffffffffffffffffffffff16610dac565b15611efa576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ef1906176b1565b60405180910390fd5b611f02610dce565b611f0b33613270565b15611f1b57611f1a3382614d32565b5b50565b611f3d3373ffffffffffffffffffffffffffffffffffffffff16610dac565b15611f7d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f74906176b1565b60405180910390fd5b815f60115f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541161203a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161203190618706565b60405180910390fd5b612045338484614e2f565b505050565b6120693373ffffffffffffffffffffffffffffffffffffffff16610dac565b156120a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120a0906176b1565b60405180910390fd5b6120b2826150a3565b6120bb82613270565b156120cb576120ca8282614d32565b5b5050565b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461213d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161213490618592565b60405180910390fd5b8060148460968110612152576121516184c6565b5b602091828204019190066101000a81548160ff021916908360ff160217905550600360ff166014846096811061218b5761218a6184c6565b5b602091828204019190069054906101000a900460ff1660ff16106121b3576121b282615830565b5b8173ffffffffffffffffffffffffffffffffffffffff167ff833a6e5fb833592d2d6a2f115227e397239985e961011fd1cd3cec3369152574360135f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015460148760968110612241576122406184c6565b5b602091828204019190069054906101000a900460ff1660405161226693929190618754565b60405180910390a2505050565b6103e881565b6122983373ffffffffffffffffffffffffffffffffffffffff16610dac565b156122d8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122cf906176b1565b60405180910390fd5b6122e1816158ff565b50565b601981565b5f60115f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b5f808290508051915050919050565b5f61238682614b98565b9050919050565b5f600854905090565b61239e614b1a565b6123a75f615a71565b565b690a968163f0a57b40000081565b5f6123c182615b32565b9050919050565b609681565b601981815481106123dc575f80fd5b905f5260205f2090600202015f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600101805461242090618645565b80601f016020809104026020016040519081016040528092919081815260200182805461244c90618645565b80156124975780601f1061246e57610100808354040283529160200191612497565b820191905f5260205f20905b81548152906001019060200180831161247a57829003601f168201915b5050505050905082565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60105481565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461255d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612554906176b1565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff166108fc691e11d5de63c17cc0000090811502906040515f60405180830381858888f193505050501580156125aa573d5f803e3d5ffd5b50691e11d5de63c17cc0000060055f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546126019190618789565b92505081905550691e11d5de63c17cc0000060075f8282546126239190618789565b925050819055507fb67ea9e632aa74e96032afed3f64179247898d3dfa8d9aef85057d36e9112d46816040516126599190617f0e565b60405180910390a150565b5f60115f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b69054b40b1f852bda0000081565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612792576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612789906176b1565b60405180910390fd5b5f8061279d83613ec1565b915091505f60055f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f691e11d5de63c17cc0000060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20546128369190618789565b90505f81836128459190618789565b905080600754101561288c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161288390618806565b60405180910390fd5b8473ffffffffffffffffffffffffffffffffffffffff166108fc8390811502906040515f60405180830381858888f193505050501580156128cf573d5f803e3d5ffd5b5060035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f19350505050158015612934573d5f803e3d5ffd5b508260075f8282546129469190618789565b9250508190555061295684615b84565b61295f846161e3565b5f60055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505f600d5f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505f600e5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505f600b5f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505f600c5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f6782e2944a228b3021a467b1afb406079eea2c32d259933346a946601e56d2ae8583604051612b70929190618824565b60405180910390a1505050505050565b60095481565b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612bf4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612beb90618592565b60405180910390fd5b612bfe82826162e1565b5050565b60605f60125f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208054905090505f8167ffffffffffffffff811115612c6357612c62617b5a565b5b604051908082528060200260200182016040528015612c9c57816020015b612c896175ad565b815260200190600190039081612c815790505b5090505f5b828167ffffffffffffffff161015612f21575f60125f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208267ffffffffffffffff1681548110612d0d57612d0c6184c6565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f60115f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f60135f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff1690505f60135f8a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015490505f60135f8b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206004015490505f6040518060a001604052808773ffffffffffffffffffffffffffffffffffffffff168152602001868152602001856005811115612ece57612ecd61774f565b5b815260200184815260200183815250905080888867ffffffffffffffff1681518110612efd57612efc6184c6565b5b60200260200101819052505050505050508080612f19906184f3565b915050612ca1565b508092505050919050565b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612f9a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612f9190618592565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff167f4d381e135b58a2c3686763368fd0d7750138ab8a9a0e404a6c09e03f2410036843604051612fe09190618895565b60405180910390a250565b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614613059576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161305090618592565b60405180910390fd5b5f5b609667ffffffffffffffff168167ffffffffffffffff1610156130ca575f60148267ffffffffffffffff1660968110613097576130966184c6565b5b602091828204019190066101000a81548160ff021916908360ff16021790555080806130c2906184f3565b91505061305b565b50565b6107d081565b5f60055f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b60085481565b606060018054806020026020016040519081016040528092919081815260200182805480156131a057602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311613157575b5050505050905090565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b80600a5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090816132199190618a55565b503373ffffffffffffffffffffffffffffffffffffffff167f472da4d064218fa97032725fbcff922201fa643fed0765b5ffe0ceef63d7b3dc826040516132609190618b6c565b60405180910390a250565b600481565b5f805f90505b6019805490508167ffffffffffffffff161015613327578273ffffffffffffffffffffffffffffffffffffffff1660198267ffffffffffffffff16815481106132c2576132c16184c6565b5b905f5260205f2090600202015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603613314575f91505061332d565b808061331f906184f3565b915050613276565b50600190505b919050565b60075481565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146133ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016133e3906176b1565b60405180910390fd5b690a968163f0a57b4000006fffffffffffffffffffffffffffffffff1634101561344b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161344290618bfc565b60405180910390fd5b5f60135f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015490505f4390508082106134d5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016134cc90618448565b60405180910390fd5b3460075f8282546134e69190618493565b925050819055503460055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546135399190618493565b9250508190555061354984613fb9565b1561355857613557846140c3565b5b61356184613270565b15613571576135708484614d32565b5b8373ffffffffffffffffffffffffffffffffffffffff167f9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d346040516135b791906178d6565b60405180910390a250505050565b5f600954905090565b5f60125f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805490509050919050565b61361f614b1a565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361368d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161368490618c8a565b60405180910390fd5b61369681615a71565b50565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614613728576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161371f906176b1565b60405180910390fd5b690a968163f0a57b4000006fffffffffffffffffffffffffffffffff16341015613787576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161377e90618bfc565b60405180910390fd5b5f60135f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015490505f429050808210613811576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161380890618448565b60405180910390fd5b61381a85614b98565b801561382b575061382a84615b32565b5b15613992578373ffffffffffffffffffffffffffffffffffffffff16600b5f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614801561394e57508473ffffffffffffffffffffffffffffffffffffffff16600c5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b61398d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161398490618d18565b60405180910390fd5b613aa2565b61399b85614b98565b1580156139ae57506139ac84614b98565b155b80156139c057506139be84615b32565b155b80156139d257506139d085615b32565b155b8015613a1a57505f60055f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054145b8015613a6257505f60055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054145b613aa1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613a9890618df2565b60405180910390fd5b5b6001600d5f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055506001600e5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555083600b5f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555084600c5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503460075f828254613c539190618493565b925050819055503460055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254613ca69190618493565b92505081905550613cb684613fb9565b15613cc557613cc4846140c3565b5b613cce84613270565b15613cde57613cdd8484614d32565b5b8473ffffffffffffffffffffffffffffffffffffffff167f9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d34604051613d2491906178d6565b60405180910390a25050505050565b60605f609690505f8167ffffffffffffffff1667ffffffffffffffff811115613d5f57613d5e617b5a565b5b604051908082528060200260200182016040528015613d8d5781602001602082028036833780820191505090505b5090505f5b8267ffffffffffffffff168167ffffffffffffffff161015613e265760148167ffffffffffffffff1660968110613dcc57613dcb6184c6565b5b602091828204019190069054906101000a900460ff16828267ffffffffffffffff1681518110613dff57613dfe6184c6565b5b602002602001019060ff16908160ff16815250508080613e1e906184f3565b915050613d92565b50809250505090565b600381565b60018181548110613e43575f80fd5b905f5260205f20015f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f60045f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff169050919050565b5f805f8390505f849050613ed485614b98565b15613f3d57600b5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050613fac565b613f4685615b32565b15613fab57600c5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505b5b8181935093505050915091565b5f613fc382616caa565b801561403e5750600580811115613fdd57613fdc61774f565b5b60135f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff16600581111561403c5761403b61774f565b5b145b1561404c57600190506140be565b61405582616caa565b1580156140bb5750690a968163f0a57b4000006fffffffffffffffffffffffffffffffff1660055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410155b90505b919050565b6005808111156140d6576140d561774f565b5b60135f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff1660058111156141355761413461774f565b5b146142805760095460018054905010614183576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161417a90618e80565b60405180910390fd5b600160045f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555060018054905060065f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550600181908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b600160135f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f6101000a81548160ff021916908360058111156142e3576142e261774f565b5b02179055505f60135f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003018190555050565b5f8061433b83613ec1565b915091505f60135f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2060040154905060025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ede391fb856040518263ffffffff1660e01b81526004016143dd9190617f0e565b602060405180830381865afa1580156143f8573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061441c9190618345565b80156144be575060025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633700c315856040518263ffffffff1660e01b815260040161447d9190617f0e565b602060405180830381865afa158015614498573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906144bc9190618345565b155b156146cc5760025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a209a86f856040518263ffffffff1660e01b815260040161451d9190617f0e565b6020604051808303815f875af1158015614539573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061455d9190618eb2565b90505f81036145d657600160135f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f6101000a81548160ff021916908360058111156145c9576145c861774f565b5b0217905550505050614b17565b60055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054811461468757600160135f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f6101000a81548160ff0219169083600581111561467d5761467c61774f565b5b02179055506146cb565b60135f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206004015490505b5b5f811480614717575060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20548110155b156147615760055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490506147ca565b600160135f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f6101000a81548160ff021916908360058111156147c4576147c361774f565b5b02179055505b8060055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546148169190618789565b925050819055508060075f82825461482e9190618789565b925050819055505f60135f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206004018190555061488384616caa565b80156148cb57505f60055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054145b15614a81575f600d5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505f600e5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505f600b5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505f600c5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550614a7784615b84565b614a80846161e3565b5b8273ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f19350505050158015614ac4573d5f803e3d5ffd5b508373ffffffffffffffffffffffffffffffffffffffff167f0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f7582604051614b0b91906178d6565b60405180910390a25050505b50565b614b22616cfc565b73ffffffffffffffffffffffffffffffffffffffff16614b406124a1565b73ffffffffffffffffffffffffffffffffffffffff1614614b96576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614b8d90618f27565b60405180910390fd5b565b5f600d5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff169050919050565b5f80614bf533613ec1565b915091505f60135f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20600301549050438110614c7f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614c7690618448565b60405180910390fd5b600260135f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f6101000a81548160ff02191690836005811115614ce257614ce161774f565b5b02179055508360135f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206004018190555050505050565b5f614d3c8261236d565b11614d7c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614d7390618f8f565b60405180910390fd5b601960405180604001604052808473ffffffffffffffffffffffffffffffffffffffff16815260200183815250908060018154018082558091505060019003905f5260205f2090600202015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001019081614e289190619005565b5050505050565b5f60115f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f82118015614ebb57508082105b15614ec4578190505b8060115f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254614f4d9190618789565b925050819055508060075f828254614f659190618789565b92505081905550614f768484612664565b8015614ffb57505f60115f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054145b1561500b5761500a8484616d03565b5b8373ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f1935050505015801561504e573d5f803e3d5ffd5b508373ffffffffffffffffffffffffffffffffffffffff167fce2e26dc37234e172764d0a8a2c04d2e6fb46c02f3ca6ca5135d6f9dcd2b67a98260405161509591906178d6565b60405180910390a250505050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ede391fb336040518263ffffffff1660e01b81526004016150fd9190617f0e565b602060405180830381865afa158015615118573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061513c9190618345565b1561529f576005808111156151545761515361774f565b5b60135f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff1660058111156151b3576151b261774f565b5b1415806151c057505f3414155b1561529e5760025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663489fe625336040518263ffffffff1660e01b815260040161521f9190617f0e565b602060405180830381865afa15801561523a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061525e9190618345565b61529d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615294906183e0565b60405180910390fd5b5b5b5f60135f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015490505f429050808210615329576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161532090618448565b60405180910390fd5b61533233614b98565b8015615343575061534283615b32565b5b156154aa578273ffffffffffffffffffffffffffffffffffffffff16600b5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614801561546657503373ffffffffffffffffffffffffffffffffffffffff16600c5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b6154a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161549c90618d18565b60405180910390fd5b6155ba565b6154b333614b98565b1580156154c657506154c483614b98565b155b80156154d857506154d683615b32565b155b80156154ea57506154e833615b32565b155b801561553257505f60055f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054145b801561557a57505f60055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054145b6155b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016155b090618df2565b60405180910390fd5b5b6001600d5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055506001600e5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555082600b5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555033600c5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503460075f82825461576b9190618493565b925050819055503460055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546157be9190618493565b925050819055506157ce83613fb9565b156157dd576157dc836140c3565b5b3373ffffffffffffffffffffffffffffffffffffffff167f9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d3460405161582391906178d6565b60405180910390a2505050565b600360135f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f6101000a81548160ff021916908360058111156158935761589261774f565b5b02179055506103e86fffffffffffffffffffffffffffffffff16436158b89190618493565b60135f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003018190555050565b61590881616caa565b615947576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161593e90619144565b60405180910390fd5b3460075f8282546159589190618493565b925050819055503460115f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546159e89190618493565b925050819055506159f93382617282565b15615a0957615a083382617337565b5b8073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f1ed8ad98d928651a8bc3999999b718383931f4595fcd2e1efd2de972fa8cdaea34604051615a6691906178d6565b60405180910390a350565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f600e5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff169050919050565b60085460018054905011615bcd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615bc4906191d2565b60405180910390fd5b60018054905060065f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410615c51576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615c489061923a565b60405180910390fd5b5f60065f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f60018080549050615ca49190618789565b9050808214615d8a575f60018281548110615cc257615cc16184c6565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508060018481548110615d0157615d006184c6565b5b905f5260205f20015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508260065f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f60045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505f60065f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055506001805480615e3257615e31619258565b5b600190038181905f5260205f20015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055905560016019805490501015615e795750506161e0565b5f806001601980549050615e8d9190618789565b90505f5b601980549050811015615f29578573ffffffffffffffffffffffffffffffffffffffff1660198281548110615ec957615ec86184c6565b5b905f5260205f2090600202015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603615f16578092505b8080615f2190618675565b915050615e91565b5080821461617e575f60198381548110615f4657615f456184c6565b5b905f5260205f2090600202016040518060400160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182018054615fc190618645565b80601f0160208091040260200160405190810160405280929190818152602001828054615fed90618645565b80156160385780601f1061600f57610100808354040283529160200191616038565b820191905f5260205f20905b81548152906001019060200180831161601b57829003601f168201915b505050505081525050905060198281548110616057576160566184c6565b5b905f5260205f20906002020160198481548110616077576160766184c6565b5b905f5260205f2090600202015f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600182018160010190816160fa919061929a565b509050508060198381548110616113576161126184c6565b5b905f5260205f2090600202015f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010190816161789190619005565b50905050505b60198054806161905761618f619258565b5b600190038181905f5260205f2090600202015f8082015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600182015f6161d791906175ff565b50509055505050505b50565b5f60125f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208054905090505f5b818167ffffffffffffffff1610156162dc576162c660125f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208267ffffffffffffffff1681548110616297576162966184c6565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16845f614e2f565b81806162d19061937f565b9250505f8203616229575b505050565b600467ffffffffffffffff168167ffffffffffffffff1611801561631a5750609667ffffffffffffffff168167ffffffffffffffff1611155b156163ec57600460135f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f6101000a81548160ff021916908360058111156163825761638161774f565b5b02179055506107d06fffffffffffffffffffffffffffffffff16436163a79190618493565b60135f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20600301819055505b5f6163f683611aad565b90505f8160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20546164429190618493565b90505f60055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f606460196fffffffffffffffffffffffffffffffff16846164a791906193a6565b6164b19190619414565b905069054b40b1f852bda000006fffffffffffffffffffffffffffffffff168110156164f75769054b40b1f852bda000006fffffffffffffffffffffffffffffffff1690505b828111156167ea57600260135f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f6101000a81548160ff021916908360058111156165625761656161774f565b5b02179055505f60055f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505f60125f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2080548060200260200160405190810160405280929190818152602001828054801561666457602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161661b575b505050505090505f60125f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208054905011156167c8575f5b60125f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805490508167ffffffffffffffff1610156167bd575f60115f848467ffffffffffffffff1681518110616728576167276184c6565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208190555080806167b5906184f3565b9150506166b5565b506167c7876161e3565b5b8360075f8282546167d99190618789565b925050819055505050505050616ca6565b5f836064836167f991906193a6565b6168039190619414565b90505f6064828561681491906193a6565b61681e9190619414565b90505f818561682d9190618789565b90508060055f8b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508160075f8282546168829190618789565b925050819055505f73ffffffffffffffffffffffffffffffffffffffff166108fc8390811502906040515f60405180830381858888f193505050501580156168cc573d5f803e3d5ffd5b505f60125f8b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2080548060200260200160405190810160405280929190818152602001828054801561698857602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161693f575b505050505090505f60125f8c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805490501115616c09575f5b60125f8c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805490508167ffffffffffffffff161015616c07575f60115f848467ffffffffffffffff1681518110616a4c57616a4b6184c6565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f60648783616adc91906193a6565b616ae69190619414565b90508082616af49190618789565b60115f868667ffffffffffffffff1681518110616b1457616b136184c6565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8f73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508060075f828254616ba79190618789565b925050819055505f73ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f19350505050158015616bf1573d5f803e3d5ffd5b5050508080616bff906184f3565b9150506169d9565b505b8973ffffffffffffffffffffffffffffffffffffffff167f549f9f77a981f051b72b0b5e384ae6a1c3155b1789216ddf6202aaeaf96a6a394360135f8e73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2060030154888b604051616c959493929190619444565b60405180910390a250505050505050505b5050565b5f60045f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff169050919050565b5f33905090565b600f5460125f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208054905011616d87576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401616d7e906194f7565b60405180910390fd5b60125f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208054905060115f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410616e84576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401616e7b90619585565b60405180910390fd5b5f60115f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f600160125f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2080549050616f519190618789565b90508082146170eb575f60125f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208281548110616faa57616fa96184c6565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508060125f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208481548110617024576170236184c6565b5b905f5260205f20015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508260115f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f60115f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505f60115f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208190555060125f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2080548061724a57617249619258565b5b600190038181905f5260205f20015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055905550505050565b5f61728d8383612664565b15801561732f57506802b5e3af16b18800006fffffffffffffffffffffffffffffffff1660115f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410155b905092915050565b60105460125f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2080549050106173bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016173b290619613565b60405180910390fd5b600160115f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555060125f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208054905060115f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208190555060125f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2082908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b6040518060a001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f60058111156175ed576175ec61774f565b5b81526020015f81526020015f81525090565b50805461760b90618645565b5f825580601f1061761c5750617639565b601f0160209004905f5260205f2090810190617638919061763c565b5b50565b5b80821115617653575f815f90555060010161763d565b5090565b5f82825260208201905092915050565b7f4f6e6c7920454f412063616e2063616c6c2066756e6374696f6e0000000000005f82015250565b5f61769b601a83617657565b91506176a682617667565b602082019050919050565b5f6020820190508181035f8301526176c88161768f565b9050919050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f617721826176f8565b9050919050565b61773181617717565b82525050565b5f819050919050565b61774981617737565b82525050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6006811061778d5761778c61774f565b5b50565b5f81905061779d8261777c565b919050565b5f6177ac82617790565b9050919050565b6177bc816177a2565b82525050565b60a082015f8201516177d65f850182617728565b5060208201516177e96020850182617740565b5060408201516177fc60408501826177b3565b50606082015161780f6060850182617740565b5060808201516178226080850182617740565b50505050565b5f61783383836177c2565b60a08301905092915050565b5f602082019050919050565b5f617855826176cf565b61785f81856176d9565b935061786a836176e9565b805f5b8381101561789a5781516178818882617828565b975061788c8361783f565b92505060018101905061786d565b5085935050505092915050565b5f6020820190508181035f8301526178bf818461784b565b905092915050565b6178d081617737565b82525050565b5f6020820190506178e95f8301846178c7565b92915050565b5f604051905090565b5f80fd5b5f80fd5b61790981617717565b8114617913575f80fd5b50565b5f8135905061792481617900565b92915050565b5f6020828403121561793f5761793e6178f8565b5b5f61794c84828501617916565b91505092915050565b61795e81617737565b8114617968575f80fd5b50565b5f8135905061797981617955565b92915050565b5f60208284031215617994576179936178f8565b5b5f6179a18482850161796b565b91505092915050565b5f6fffffffffffffffffffffffffffffffff82169050919050565b6179ce816179aa565b82525050565b5f6020820190506179e75f8301846179c5565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015617a4d578082015181840152602081019050617a32565b5f8484015250505050565b5f601f19601f8301169050919050565b5f617a7282617a16565b617a7c8185617a20565b9350617a8c818560208601617a30565b617a9581617a58565b840191505092915050565b5f617aab8383617a68565b905092915050565b5f602082019050919050565b5f617ac9826179ed565b617ad381856179f7565b935083602082028501617ae585617a07565b805f5b85811015617b205784840389528151617b018582617aa0565b9450617b0c83617ab3565b925060208a01995050600181019050617ae8565b50829750879550505050505092915050565b5f6020820190508181035f830152617b4a8184617abf565b905092915050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b617b9082617a58565b810181811067ffffffffffffffff82111715617baf57617bae617b5a565b5b80604052505050565b5f617bc16178ef565b9050617bcd8282617b87565b919050565b5f67ffffffffffffffff821115617bec57617beb617b5a565b5b617bf582617a58565b9050602081019050919050565b828183375f83830152505050565b5f617c22617c1d84617bd2565b617bb8565b905082815260208101848484011115617c3e57617c3d617b56565b5b617c49848285617c02565b509392505050565b5f82601f830112617c6557617c64617b52565b5b8135617c75848260208601617c10565b91505092915050565b5f60208284031215617c9357617c926178f8565b5b5f82013567ffffffffffffffff811115617cb057617caf6178fc565b5b617cbc84828501617c51565b91505092915050565b5f8060408385031215617cdb57617cda6178f8565b5b5f617ce885828601617916565b9250506020617cf98582860161796b565b9150509250929050565b5f8060408385031215617d1957617d186178f8565b5b5f617d2685828601617916565b925050602083013567ffffffffffffffff811115617d4757617d466178fc565b5b617d5385828601617c51565b9150509250929050565b5f60ff82169050919050565b617d7281617d5d565b8114617d7c575f80fd5b50565b5f81359050617d8d81617d69565b92915050565b5f805f60608486031215617daa57617da96178f8565b5b5f617db78682870161796b565b9350506020617dc886828701617916565b9250506040617dd986828701617d7f565b9150509250925092565b5f8060408385031215617df957617df86178f8565b5b5f617e0685828601617916565b9250506020617e1785828601617916565b9150509250929050565b5f8115159050919050565b617e3581617e21565b82525050565b5f602082019050617e4e5f830184617e2c565b92915050565b5f67ffffffffffffffff82169050919050565b617e7081617e54565b82525050565b5f602082019050617e895f830184617e67565b92915050565b617e9881617717565b82525050565b5f81519050919050565b5f617eb282617e9e565b617ebc8185617657565b9350617ecc818560208601617a30565b617ed581617a58565b840191505092915050565b5f604082019050617ef35f830185617e8f565b8181036020830152617f058184617ea8565b90509392505050565b5f602082019050617f215f830184617e8f565b92915050565b617f3081617e54565b8114617f3a575f80fd5b50565b5f81359050617f4b81617f27565b92915050565b5f8060408385031215617f6757617f666178f8565b5b5f617f7485828601617916565b9250506020617f8585828601617f3d565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f617fc38383617728565b60208301905092915050565b5f602082019050919050565b5f617fe582617f8f565b617fef8185617f99565b9350617ffa83617fa9565b805f5b8381101561802a5781516180118882617fb8565b975061801c83617fcf565b925050600181019050617ffd565b5085935050505092915050565b5f6020820190508181035f83015261804f8184617fdb565b905092915050565b5f819050919050565b5f61807a618075618070846176f8565b618057565b6176f8565b9050919050565b5f61808b82618060565b9050919050565b5f61809c82618081565b9050919050565b6180ac81618092565b82525050565b5f6020820190506180c55f8301846180a3565b92915050565b5f67ffffffffffffffff8211156180e5576180e4617b5a565b5b6180ee82617a58565b9050602081019050919050565b5f61810d618108846180cb565b617bb8565b90508281526020810184848401111561812957618128617b56565b5b618134848285617c02565b509392505050565b5f82601f8301126181505761814f617b52565b5b81356181608482602086016180fb565b91505092915050565b5f6020828403121561817e5761817d6178f8565b5b5f82013567ffffffffffffffff81111561819b5761819a6178fc565b5b6181a78482850161813c565b91505092915050565b5f805f606084860312156181c7576181c66178f8565b5b5f6181d486828701617916565b93505060206181e586828701617916565b925050604084013567ffffffffffffffff811115618206576182056178fc565b5b61821286828701617c51565b9150509250925092565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61824e81617d5d565b82525050565b5f61825f8383618245565b60208301905092915050565b5f602082019050919050565b5f6182818261821c565b61828b8185618226565b935061829683618236565b805f5b838110156182c65781516182ad8882618254565b97506182b88361826b565b925050600181019050618299565b5085935050505092915050565b5f6020820190508181035f8301526182eb8184618277565b905092915050565b6182fc81617d5d565b82525050565b5f6020820190506183155f8301846182f3565b92915050565b61832481617e21565b811461832e575f80fd5b50565b5f8151905061833f8161831b565b92915050565b5f6020828403121561835a576183596178f8565b5b5f61836784828501618331565b91505092915050565b7f706c656173652070617920746865206772616e74656420616d6f756e742062655f8201527f666f7265207374616b696e6720616761696e0000000000000000000000000000602082015250565b5f6183ca603283617657565b91506183d582618370565b604082019050919050565b5f6020820190508181035f8301526183f7816183be565b9050919050565b7f7374616b65722069732062616e6e6564206f722073757370656e6465640000005f82015250565b5f618432601d83617657565b915061843d826183fe565b602082019050919050565b5f6020820190508181035f83015261845f81618426565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61849d82617737565b91506184a883617737565b92508282019050808211156184c0576184bf618466565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f6184fd82617e54565b915067ffffffffffffffff820361851757618516618466565b5b600182019050919050565b7f4f6e6c7920626c6f636b2070726f706f7365722063616e2063616c6c2066756e5f8201527f6374696f6e000000000000000000000000000000000000000000000000000000602082015250565b5f61857c602583617657565b915061858782618522565b604082019050919050565b5f6020820190508181035f8301526185a981618570565b9050919050565b7f4f6e6c79207374616b65722063616e2063616c6c2066756e6374696f6e0000005f82015250565b5f6185e4601d83617657565b91506185ef826185b0565b602082019050919050565b5f6020820190508181035f830152618611816185d8565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061865c57607f821691505b60208210810361866f5761866e618618565b5b50919050565b5f61867f82617737565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036186b1576186b0618466565b5b600182019050919050565b7f6f6e6c792064656c656761746f722063616e2063616c6c2066756e6374696f6e5f82015250565b5f6186f0602083617657565b91506186fb826186bc565b602082019050919050565b5f6020820190508181035f83015261871d816186e4565b9050919050565b5f61873e61873961873484617d5d565b618057565b617e54565b9050919050565b61874e81618724565b82525050565b5f6060820190506187675f8301866178c7565b61877460208301856178c7565b6187816040830184618745565b949350505050565b5f61879382617737565b915061879e83617737565b92508282039050818111156187b6576187b5618466565b5b92915050565b7f6e6f7420656e6f756768206d6f6e6579000000000000000000000000000000005f82015250565b5f6187f0601083617657565b91506187fb826187bc565b602082019050919050565b5f6020820190508181035f83015261881d816187e4565b9050919050565b5f6040820190506188375f830185617e8f565b61884460208301846178c7565b9392505050565b7f74696d656f7574000000000000000000000000000000000000000000000000005f82015250565b5f61887f600783617657565b915061888a8261884b565b602082019050919050565b5f6040820190506188a85f8301846178c7565b81810360208301526188b981618873565b905092915050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261891d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826188e2565b61892786836188e2565b95508019841693508086168417925050509392505050565b5f61895961895461894f84617737565b618057565b617737565b9050919050565b5f819050919050565b6189728361893f565b61898661897e82618960565b8484546188ee565b825550505050565b5f90565b61899a61898e565b6189a5818484618969565b505050565b5b818110156189c8576189bd5f82618992565b6001810190506189ab565b5050565b601f821115618a0d576189de816188c1565b6189e7846188d3565b810160208510156189f6578190505b618a0a618a02856188d3565b8301826189aa565b50505b505050565b5f82821c905092915050565b5f618a2d5f1984600802618a12565b1980831691505092915050565b5f618a458383618a1e565b9150826002028217905092915050565b618a5e82617a16565b67ffffffffffffffff811115618a7757618a76617b5a565b5b618a818254618645565b618a8c8282856189cc565b5f60209050601f831160018114618abd575f8415618aab578287015190505b618ab58582618a3a565b865550618b1c565b601f198416618acb866188c1565b5f5b82811015618af257848901518255600182019150602085019450602081019050618acd565b86831015618b0f5784890151618b0b601f891682618a1e565b8355505b6001600288020188555050505b505050505050565b5f82825260208201905092915050565b5f618b3e82617a16565b618b488185618b24565b9350618b58818560208601617a30565b618b6181617a58565b840191505092915050565b5f6020820190508181035f830152618b848184618b34565b905092915050565b7f416d6f756e74206d7573742062652067726561746572207468616e2076616c695f8201527f6461746f72207468726573686f6c640000000000000000000000000000000000602082015250565b5f618be6602f83617657565b9150618bf182618b8c565b604082019050919050565b5f6020820190508181035f830152618c1381618bda565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f618c74602683617657565b9150618c7f82618c1a565b604082019050919050565b5f6020820190508181035f830152618ca181618c68565b9050919050565b7f7468652073656e646572206973206e6f7420612076616c6964207374616b65725f8201527f206f662074686973207369676e65720000000000000000000000000000000000602082015250565b5f618d02602f83617657565b9150618d0d82618ca8565b604082019050919050565b5f6020820190508181035f830152618d2f81618cf6565b9050919050565b7f626f74682061646472657373657320617265206e6f742076616c696420746f205f8201527f63616c6c2066756e6374696f6e2e20496620796f75207374696c6c2077616e7460208201527f20746f2063616c6c2e204a7573742063616c6c696e6720756e7374616b65206660408201527f6f7220626f7468206163636f756e7473206d616e75616c6c7900000000000000606082015250565b5f618ddc607983617657565b9150618de782618d36565b608082019050919050565b5f6020820190508181035f830152618e0981618dd0565b9050919050565b7f56616c696461746f72207365742068617320726561636865642066756c6c20635f8201527f6170616369747900000000000000000000000000000000000000000000000000602082015250565b5f618e6a602783617657565b9150618e7582618e10565b604082019050919050565b5f6020820190508181035f830152618e9781618e5e565b9050919050565b5f81519050618eac81617955565b92915050565b5f60208284031215618ec757618ec66178f8565b5b5f618ed484828501618e9e565b91505092915050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f618f11602083617657565b9150618f1c82618edd565b602082019050919050565b5f6020820190508181035f830152618f3e81618f05565b9050919050565b7f506c6561736520736574207468652063726577206e616d6500000000000000005f82015250565b5f618f79601883617657565b9150618f8482618f45565b602082019050919050565b5f6020820190508181035f830152618fa681618f6d565b9050919050565b5f819050815f5260205f209050919050565b601f82111561900057618fd181618fad565b618fda846188d3565b81016020851015618fe9578190505b618ffd618ff5856188d3565b8301826189aa565b50505b505050565b61900e82617e9e565b67ffffffffffffffff81111561902757619026617b5a565b5b6190318254618645565b61903c828285618fbf565b5f60209050601f83116001811461906d575f841561905b578287015190505b6190658582618a3a565b8655506190cc565b601f19841661907b86618fad565b5f5b828110156190a25784890151825560018201915060208501945060208101905061907d565b868310156190bf57848901516190bb601f891682618a1e565b8355505b6001600288020188555050505b505050505050565b7f63616e6e6f742064656c656761746520746f6b656e7320746f20736f6d656f6e5f8201527f652077686f206973206e6f742076616c696461746f7200000000000000000000602082015250565b5f61912e603683617657565b9150619139826190d4565b604082019050919050565b5f6020820190508181035f83015261915b81619122565b9050919050565b7f56616c696461746f72732063616e2774206265206c657373207468616e2074685f8201527f65206d696e696d756d2072657175697265642076616c696461746f72206e756d602082015250565b5f6191bc604083617657565b91506191c782619162565b604082019050919050565b5f6020820190508181035f8301526191e9816191b0565b9050919050565b7f696e646578206f7574206f662072616e676500000000000000000000000000005f82015250565b5f619224601283617657565b915061922f826191f0565b602082019050919050565b5f6020820190508181035f83015261925181619218565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b5f8154905061929381618645565b9050919050565b8181036192a857505061937d565b6192b182619285565b67ffffffffffffffff8111156192ca576192c9617b5a565b5b6192d48254618645565b6192df828285618fbf565b5f601f83116001811461930c575f84156192fa578287015490505b6193048582618a3a565b865550619376565b601f19841661931a87618fad565b965061932586618fad565b5f5b8281101561934c57848901548255600182019150600185019450602081019050619327565b868310156193695784890154619365601f891682618a1e565b8355505b6001600288020188555050505b5050505050505b565b5f61938982617737565b91505f820361939b5761939a618466565b5b600182039050919050565b5f6193b082617737565b91506193bb83617737565b92508282026193c981617737565b915082820484148315176193e0576193df618466565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61941e82617737565b915061942983617737565b925082619439576194386193e7565b5b828204905092915050565b5f6080820190506194575f8301876178c7565b61946460208301866178c7565b61947160408301856178c7565b61947e60608301846178c7565b95945050505050565b7f64656c656761746f72732073697a652063616e2774206265206c6573732074685f8201527f616e20746865206d696e696d756d20726571756972656d656e74000000000000602082015250565b5f6194e1603a83617657565b91506194ec82619487565b604082019050919050565b5f6020820190508181035f83015261950e816194d5565b9050919050565b7f696e646578206f7574206f662072616e676520696e207468652064656c6567615f8201527f746f7273206c697374206f662076616c696461746f7200000000000000000000602082015250565b5f61956f603683617657565b915061957a82619515565b604082019050919050565b5f6020820190508181035f83015261959c81619563565b9050919050565b7f64656c656761746f7220736574206861732072656163686564206974732066755f8201527f6c6c206361706163697479000000000000000000000000000000000000000000602082015250565b5f6195fd602b83617657565b9150619608826195a3565b604082019050919050565b5f6020820190508181035f83015261962a816195f1565b905091905056fea2646970667358221220b4f39383184b202cc1eba95aa2cbdd0c4f3273adde2605db6528b405073fc8db64736f6c63430008140033",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// StakingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingMetaData.Bin instead.
var StakingBin = StakingMetaData.Bin

// DeployStaking deploys a new Ethereum contract, binding an instance of Staking to it.
func DeployStaking(auth *bind.TransactOpts, backend bind.ContractBackend, minNumValidators *big.Int, maxNumValidators *big.Int, grantContract common.Address) (common.Address, *types.Transaction, *Staking, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingBin), backend, minNumValidators, maxNumValidators, grantContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

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

// TOKENSLASHINGTHRESHOLD is a free data retrieval call binding the contract method 0xa157339c.
//
// Solidity: function TOKEN_SLASHING_THRESHOLD() view returns(uint128)
func (_Staking *StakingCaller) TOKENSLASHINGTHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "TOKEN_SLASHING_THRESHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOKENSLASHINGTHRESHOLD is a free data retrieval call binding the contract method 0xa157339c.
//
// Solidity: function TOKEN_SLASHING_THRESHOLD() view returns(uint128)
func (_Staking *StakingSession) TOKENSLASHINGTHRESHOLD() (*big.Int, error) {
	return _Staking.Contract.TOKENSLASHINGTHRESHOLD(&_Staking.CallOpts)
}

// TOKENSLASHINGTHRESHOLD is a free data retrieval call binding the contract method 0xa157339c.
//
// Solidity: function TOKEN_SLASHING_THRESHOLD() view returns(uint128)
func (_Staking *StakingCallerSession) TOKENSLASHINGTHRESHOLD() (*big.Int, error) {
	return _Staking.Contract.TOKENSLASHINGTHRESHOLD(&_Staking.CallOpts)
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

// CanBecomeBoss is a free data retrieval call binding the contract method 0xe069fe99.
//
// Solidity: function canBecomeBoss(address addr) view returns(bool)
func (_Staking *StakingCaller) CanBecomeBoss(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "canBecomeBoss", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanBecomeBoss is a free data retrieval call binding the contract method 0xe069fe99.
//
// Solidity: function canBecomeBoss(address addr) view returns(bool)
func (_Staking *StakingSession) CanBecomeBoss(addr common.Address) (bool, error) {
	return _Staking.Contract.CanBecomeBoss(&_Staking.CallOpts, addr)
}

// CanBecomeBoss is a free data retrieval call binding the contract method 0xe069fe99.
//
// Solidity: function canBecomeBoss(address addr) view returns(bool)
func (_Staking *StakingCallerSession) CanBecomeBoss(addr common.Address) (bool, error) {
	return _Staking.Contract.CanBecomeBoss(&_Staking.CallOpts, addr)
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
// Solidity: function getDelegatorsInfoOfValidator(address signer) view returns((address,uint256,uint8,uint256,uint256)[])
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
// Solidity: function getDelegatorsInfoOfValidator(address signer) view returns((address,uint256,uint8,uint256,uint256)[])
func (_Staking *StakingSession) GetDelegatorsInfoOfValidator(signer common.Address) ([]StakingAccountStakingInfo, error) {
	return _Staking.Contract.GetDelegatorsInfoOfValidator(&_Staking.CallOpts, signer)
}

// GetDelegatorsInfoOfValidator is a free data retrieval call binding the contract method 0xb1f72b52.
//
// Solidity: function getDelegatorsInfoOfValidator(address signer) view returns((address,uint256,uint8,uint256,uint256)[])
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

// GetTotalDeligatedAmount is a free data retrieval call binding the contract method 0x306c019d.
//
// Solidity: function getTotalDeligatedAmount(address signer) view returns(uint256)
func (_Staking *StakingCaller) GetTotalDeligatedAmount(opts *bind.CallOpts, signer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getTotalDeligatedAmount", signer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalDeligatedAmount is a free data retrieval call binding the contract method 0x306c019d.
//
// Solidity: function getTotalDeligatedAmount(address signer) view returns(uint256)
func (_Staking *StakingSession) GetTotalDeligatedAmount(signer common.Address) (*big.Int, error) {
	return _Staking.Contract.GetTotalDeligatedAmount(&_Staking.CallOpts, signer)
}

// GetTotalDeligatedAmount is a free data retrieval call binding the contract method 0x306c019d.
//
// Solidity: function getTotalDeligatedAmount(address signer) view returns(uint256)
func (_Staking *StakingCallerSession) GetTotalDeligatedAmount(signer common.Address) (*big.Int, error) {
	return _Staking.Contract.GetTotalDeligatedAmount(&_Staking.CallOpts, signer)
}

// GetValidatorsStakeInfo is a free data retrieval call binding the contract method 0x03306f0a.
//
// Solidity: function getValidatorsStakeInfo() view returns((address,uint256,uint8,uint256,uint256)[])
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
// Solidity: function getValidatorsStakeInfo() view returns((address,uint256,uint8,uint256,uint256)[])
func (_Staking *StakingSession) GetValidatorsStakeInfo() ([]StakingAccountStakingInfo, error) {
	return _Staking.Contract.GetValidatorsStakeInfo(&_Staking.CallOpts)
}

// GetValidatorsStakeInfo is a free data retrieval call binding the contract method 0x03306f0a.
//
// Solidity: function getValidatorsStakeInfo() view returns((address,uint256,uint8,uint256,uint256)[])
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

// IsDelegator is a free data retrieval call binding the contract method 0xa0e11929.
//
// Solidity: function isDelegator(address account, address signer) view returns(bool)
func (_Staking *StakingCaller) IsDelegator(opts *bind.CallOpts, account common.Address, signer common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isDelegator", account, signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDelegator is a free data retrieval call binding the contract method 0xa0e11929.
//
// Solidity: function isDelegator(address account, address signer) view returns(bool)
func (_Staking *StakingSession) IsDelegator(account common.Address, signer common.Address) (bool, error) {
	return _Staking.Contract.IsDelegator(&_Staking.CallOpts, account, signer)
}

// IsDelegator is a free data retrieval call binding the contract method 0xa0e11929.
//
// Solidity: function isDelegator(address account, address signer) view returns(bool)
func (_Staking *StakingCallerSession) IsDelegator(account common.Address, signer common.Address) (bool, error) {
	return _Staking.Contract.IsDelegator(&_Staking.CallOpts, account, signer)
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

// ValidatorBLSPublicKeys is a free data retrieval call binding the contract method 0x3c561f04.
//
// Solidity: function validatorBLSPublicKeys() view returns(bytes[])
func (_Staking *StakingCaller) ValidatorBLSPublicKeys(opts *bind.CallOpts) ([][]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "validatorBLSPublicKeys")

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// ValidatorBLSPublicKeys is a free data retrieval call binding the contract method 0x3c561f04.
//
// Solidity: function validatorBLSPublicKeys() view returns(bytes[])
func (_Staking *StakingSession) ValidatorBLSPublicKeys() ([][]byte, error) {
	return _Staking.Contract.ValidatorBLSPublicKeys(&_Staking.CallOpts)
}

// ValidatorBLSPublicKeys is a free data retrieval call binding the contract method 0x3c561f04.
//
// Solidity: function validatorBLSPublicKeys() view returns(bytes[])
func (_Staking *StakingCallerSession) ValidatorBLSPublicKeys() ([][]byte, error) {
	return _Staking.Contract.ValidatorBLSPublicKeys(&_Staking.CallOpts)
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

// DelegatorUnstake is a paid mutator transaction binding the contract method 0x48cb4d31.
//
// Solidity: function delegatorUnstake(address signer, uint256 amountToUnstake) returns()
func (_Staking *StakingTransactor) DelegatorUnstake(opts *bind.TransactOpts, signer common.Address, amountToUnstake *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "delegatorUnstake", signer, amountToUnstake)
}

// DelegatorUnstake is a paid mutator transaction binding the contract method 0x48cb4d31.
//
// Solidity: function delegatorUnstake(address signer, uint256 amountToUnstake) returns()
func (_Staking *StakingSession) DelegatorUnstake(signer common.Address, amountToUnstake *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.DelegatorUnstake(&_Staking.TransactOpts, signer, amountToUnstake)
}

// DelegatorUnstake is a paid mutator transaction binding the contract method 0x48cb4d31.
//
// Solidity: function delegatorUnstake(address signer, uint256 amountToUnstake) returns()
func (_Staking *StakingTransactorSession) DelegatorUnstake(signer common.Address, amountToUnstake *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.DelegatorUnstake(&_Staking.TransactOpts, signer, amountToUnstake)
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

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xd94c111b.
//
// Solidity: function registerBLSPublicKey(bytes blsPubKey) returns()
func (_Staking *StakingTransactor) RegisterBLSPublicKey(opts *bind.TransactOpts, blsPubKey []byte) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "registerBLSPublicKey", blsPubKey)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xd94c111b.
//
// Solidity: function registerBLSPublicKey(bytes blsPubKey) returns()
func (_Staking *StakingSession) RegisterBLSPublicKey(blsPubKey []byte) (*types.Transaction, error) {
	return _Staking.Contract.RegisterBLSPublicKey(&_Staking.TransactOpts, blsPubKey)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xd94c111b.
//
// Solidity: function registerBLSPublicKey(bytes blsPubKey) returns()
func (_Staking *StakingTransactorSession) RegisterBLSPublicKey(blsPubKey []byte) (*types.Transaction, error) {
	return _Staking.Contract.RegisterBLSPublicKey(&_Staking.TransactOpts, blsPubKey)
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

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 amountToUnstake) returns()
func (_Staking *StakingTransactor) Unstake(opts *bind.TransactOpts, amountToUnstake *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "unstake", amountToUnstake)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 amountToUnstake) returns()
func (_Staking *StakingSession) Unstake(amountToUnstake *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Unstake(&_Staking.TransactOpts, amountToUnstake)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 amountToUnstake) returns()
func (_Staking *StakingTransactorSession) Unstake(amountToUnstake *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Unstake(&_Staking.TransactOpts, amountToUnstake)
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

// StakingBLSPublicKeyRegisteredIterator is returned from FilterBLSPublicKeyRegistered and is used to iterate over the raw logs and unpacked data for BLSPublicKeyRegistered events raised by the Staking contract.
type StakingBLSPublicKeyRegisteredIterator struct {
	Event *StakingBLSPublicKeyRegistered // Event containing the contract specifics and raw log

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
func (it *StakingBLSPublicKeyRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingBLSPublicKeyRegistered)
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
		it.Event = new(StakingBLSPublicKeyRegistered)
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
func (it *StakingBLSPublicKeyRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingBLSPublicKeyRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingBLSPublicKeyRegistered represents a BLSPublicKeyRegistered event raised by the Staking contract.
type StakingBLSPublicKeyRegistered struct {
	Accout common.Address
	Key    []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBLSPublicKeyRegistered is a free log retrieval operation binding the contract event 0x472da4d064218fa97032725fbcff922201fa643fed0765b5ffe0ceef63d7b3dc.
//
// Solidity: event BLSPublicKeyRegistered(address indexed accout, bytes key)
func (_Staking *StakingFilterer) FilterBLSPublicKeyRegistered(opts *bind.FilterOpts, accout []common.Address) (*StakingBLSPublicKeyRegisteredIterator, error) {

	var accoutRule []interface{}
	for _, accoutItem := range accout {
		accoutRule = append(accoutRule, accoutItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "BLSPublicKeyRegistered", accoutRule)
	if err != nil {
		return nil, err
	}
	return &StakingBLSPublicKeyRegisteredIterator{contract: _Staking.contract, event: "BLSPublicKeyRegistered", logs: logs, sub: sub}, nil
}

// WatchBLSPublicKeyRegistered is a free log subscription operation binding the contract event 0x472da4d064218fa97032725fbcff922201fa643fed0765b5ffe0ceef63d7b3dc.
//
// Solidity: event BLSPublicKeyRegistered(address indexed accout, bytes key)
func (_Staking *StakingFilterer) WatchBLSPublicKeyRegistered(opts *bind.WatchOpts, sink chan<- *StakingBLSPublicKeyRegistered, accout []common.Address) (event.Subscription, error) {

	var accoutRule []interface{}
	for _, accoutItem := range accout {
		accoutRule = append(accoutRule, accoutItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "BLSPublicKeyRegistered", accoutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingBLSPublicKeyRegistered)
				if err := _Staking.contract.UnpackLog(event, "BLSPublicKeyRegistered", log); err != nil {
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

// ParseBLSPublicKeyRegistered is a log parse operation binding the contract event 0x472da4d064218fa97032725fbcff922201fa643fed0765b5ffe0ceef63d7b3dc.
//
// Solidity: event BLSPublicKeyRegistered(address indexed accout, bytes key)
func (_Staking *StakingFilterer) ParseBLSPublicKeyRegistered(log types.Log) (*StakingBLSPublicKeyRegistered, error) {
	event := new(StakingBLSPublicKeyRegistered)
	if err := _Staking.contract.UnpackLog(event, "BLSPublicKeyRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	Account            common.Address
	BlockNumber        *big.Int
	TimeLock           *big.Int
	TotalTokenSlashing *big.Int
	TotalStakedSize    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBan is a free log retrieval operation binding the contract event 0x549f9f77a981f051b72b0b5e384ae6a1c3155b1789216ddf6202aaeaf96a6a39.
//
// Solidity: event Ban(address indexed account, uint256 blockNumber, uint256 timeLock, uint256 totalTokenSlashing, uint256 totalStakedSize)
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

// WatchBan is a free log subscription operation binding the contract event 0x549f9f77a981f051b72b0b5e384ae6a1c3155b1789216ddf6202aaeaf96a6a39.
//
// Solidity: event Ban(address indexed account, uint256 blockNumber, uint256 timeLock, uint256 totalTokenSlashing, uint256 totalStakedSize)
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

// ParseBan is a log parse operation binding the contract event 0x549f9f77a981f051b72b0b5e384ae6a1c3155b1789216ddf6202aaeaf96a6a39.
//
// Solidity: event Ban(address indexed account, uint256 blockNumber, uint256 timeLock, uint256 totalTokenSlashing, uint256 totalStakedSize)
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
