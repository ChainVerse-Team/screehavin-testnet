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
	Bin: "0x60806040525f600f556103e86010553480156200001a575f80fd5b506040516200861f3803806200861f8339818101604052810190620000409190620002a2565b62000060620000546200013d60201b60201c565b6200014460201b60201c565b81831115620000a6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200009d906200037f565b60405180910390fd5b82600881905550816009819055508060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050506200039f565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f80fd5b5f819050919050565b6200021d8162000209565b811462000228575f80fd5b50565b5f815190506200023b8162000212565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6200026c8262000241565b9050919050565b6200027e8162000260565b811462000289575f80fd5b50565b5f815190506200029c8162000273565b92915050565b5f805f60608486031215620002bc57620002bb62000205565b5b5f620002cb868287016200022b565b9350506020620002de868287016200022b565b9250506040620002f1868287016200028c565b9150509250925092565b5f82825260208201905092915050565b7f4d696e2076616c696461746f7273206e756d2063616e206e6f742062652067725f8201527f6561746572207468616e206d6178206e756d206f662076616c696461746f7273602082015250565b5f62000367604083620002fb565b915062000374826200030b565b604082019050919050565b5f6020820190508181035f830152620003988162000359565b9050919050565b61827280620003ad5f395ff3fe6080604052600436106102e7575f3560e01c806390b196271161018f578063cb3d89d6116100db578063ed546c8d11610094578063f6676a3a1161006e578063f6676a3a14610b0c578063f740ff8e14610b36578063f90ecacc14610b60578063facd743b14610b9c57610355565b8063ed546c8d14610a8c578063f2fde38b14610ac8578063f3a382f214610af057610355565b8063cb3d89d61461099e578063de092ac4146109c8578063e387a7ed146109f2578063e69af0c714610a1c578063e70c91b514610a46578063e804fbf614610a6257610355565b8063b1f72b5211610148578063bb4c461e11610122578063bb4c461e146108e4578063c3a6ce2b1461090e578063c795c0771461094a578063ca1e78191461097457610355565b8063b1f72b521461086a578063b788e44e146108a6578063ba1b2125146108ce57610355565b806390b196271461077657806396f58ffa146107a05780639e723b63146107c8578063a5dbc130146107f0578063af6da36e14610818578063b0f572ee1461084257610355565b80635bb9fec91161024e578063714ff425116102075780637df73e27116101e15780637df73e27146106a95780637e840377146106e55780637ebe444c1461070f5780638da5cb5b1461074c57610355565b8063714ff4251461063f578063715018a6146106695780637a6eea371461067f57610355565b80635bb9fec91461051b5780635c19a95c146105455780635c8c4a0714610561578063636592ed1461058b57806365c19af0146105c75780636f1e85331461060357610355565b806336c4086f116102a057806336c4086f1461043d578063373d6132146104675780633c449ae71461049157806346f45b8d146104bb5780634fb856e5146104d75780635bb4f998146104f357610355565b806303306f0a1461035957806315f9eb9b1461038357806319fd547d146103ad57806322d46852146103c35780632367f6b5146103eb5780632def66201461042757610355565b366103555761030b3373ffffffffffffffffffffffffffffffffffffffff16610bd8565b1561034b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610342906166bc565b60405180910390fd5b610353610bfa565b005b5f80fd5b348015610364575f80fd5b5061036d610f65565b60405161037a919061689f565b60405180910390f35b34801561038e575f80fd5b50610397611254565b6040516103a491906168ce565b60405180910390f35b3480156103b8575f80fd5b506103c1611262565b005b3480156103ce575f80fd5b506103e960048036038101906103e49190616922565b61163b565b005b3480156103f6575f80fd5b50610411600480360381019061040c9190616922565b611709565b60405161041e91906168ce565b60405180910390f35b348015610432575f80fd5b5061043b61174f565b005b348015610448575f80fd5b50610451611847565b60405161045e91906168ce565b60405180910390f35b348015610472575f80fd5b5061047b61184d565b60405161048891906168ce565b60405180910390f35b34801561049c575f80fd5b506104a5611856565b6040516104b29190616977565b60405180910390f35b6104d560048036038101906104d09190616acc565b611862565b005b6104f160048036038101906104ec9190616b13565b6118d6565b005b3480156104fe575f80fd5b5061051960048036038101906105149190616bcd565b61194c565b005b348015610526575f80fd5b5061052f611af0565b60405161053c9190616977565b60405180910390f35b61055f600480360381019061055a9190616922565b611af6565b005b34801561056c575f80fd5b50610575611b61565b6040516105829190616977565b60405180910390f35b348015610596575f80fd5b506105b160048036038101906105ac9190616c1d565b611b66565b6040516105be91906168ce565b60405180910390f35b3480156105d2575f80fd5b506105ed60048036038101906105e89190616acc565b611bea565b6040516105fa91906168ce565b60405180910390f35b34801561060e575f80fd5b5061062960048036038101906106249190616922565b611bf9565b6040516106369190616c75565b60405180910390f35b34801561064a575f80fd5b50610653611c0a565b60405161066091906168ce565b60405180910390f35b348015610674575f80fd5b5061067d611c13565b005b34801561068a575f80fd5b50610693611c26565b6040516106a09190616977565b60405180910390f35b3480156106b4575f80fd5b506106cf60048036038101906106ca9190616922565b611c34565b6040516106dc9190616c75565b60405180910390f35b3480156106f0575f80fd5b506106f9611c45565b6040516107069190616cb0565b60405180910390f35b34801561071a575f80fd5b5061073560048036038101906107309190616cc9565b611c4a565b604051610743929190616d6d565b60405180910390f35b348015610757575f80fd5b50610760611d1e565b60405161076d9190616d9b565b60405180910390f35b348015610781575f80fd5b5061078a611d45565b60405161079791906168ce565b60405180910390f35b3480156107ab575f80fd5b506107c660048036038101906107c19190616922565b611d4b565b005b3480156107d3575f80fd5b506107ee60048036038101906107e99190616922565b611ee1565b005b3480156107fb575f80fd5b5061081660048036038101906108119190616922565b61200b565b005b348015610823575f80fd5b5061082c612478565b60405161083991906168ce565b60405180910390f35b34801561084d575f80fd5b5061086860048036038101906108639190616dde565b61247e565b005b348015610875575f80fd5b50610890600480360381019061088b9190616922565b6124fa565b60405161089d919061689f565b60405180910390f35b3480156108b1575f80fd5b506108cc60048036038101906108c79190616922565b6127d9565b005b3480156108d9575f80fd5b506108e2612898565b005b3480156108ef575f80fd5b506108f861297a565b6040516109059190616977565b60405180910390f35b348015610919575f80fd5b50610934600480360381019061092f9190616922565b612980565b60405161094191906168ce565b60405180910390f35b348015610955575f80fd5b5061095e6129c6565b60405161096b91906168ce565b60405180910390f35b34801561097f575f80fd5b506109886129cc565b6040516109959190616ec4565b60405180910390f35b3480156109a9575f80fd5b506109b2612a57565b6040516109bf9190616f3f565b60405180910390f35b3480156109d3575f80fd5b506109dc612a7c565b6040516109e99190616cb0565b60405180910390f35b3480156109fd575f80fd5b50610a06612a81565b604051610a1391906168ce565b60405180910390f35b348015610a27575f80fd5b50610a30612a87565b604051610a3d9190616d9b565b60405180910390f35b610a606004803603810190610a5b9190616b13565b612aac565b005b348015610a6d575f80fd5b50610a76612d05565b604051610a8391906168ce565b60405180910390f35b348015610a97575f80fd5b50610ab26004803603810190610aad9190616922565b612d0e565b604051610abf91906168ce565b60405180910390f35b348015610ad3575f80fd5b50610aee6004803603810190610ae99190616922565b612d57565b005b610b0a6004803603810190610b059190616f58565b612dd9565b005b348015610b17575f80fd5b50610b20613464565b604051610b2d919061707b565b60405180910390f35b348015610b41575f80fd5b50610b4a613560565b604051610b5791906170aa565b60405180910390f35b348015610b6b575f80fd5b50610b866004803603810190610b819190616cc9565b613565565b604051610b939190616d9b565b60405180910390f35b348015610ba7575f80fd5b50610bc26004803603810190610bbd9190616922565b6135a0565b604051610bcf9190616c75565b60405180910390f35b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ede391fb336040518263ffffffff1660e01b8152600401610c549190616d9b565b602060405180830381865afa158015610c6f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c9391906170ed565b8015610d1a5750600580811115610cad57610cac61675a565b5b60135f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff166005811115610d0c57610d0b61675a565b5b141580610d1957505f3414155b5b15610df85760025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663489fe625336040518263ffffffff1660e01b8152600401610d799190616d9b565b602060405180830381865afa158015610d94573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610db891906170ed565b610df7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dee90617188565b60405180910390fd5b5b5f610e02336135f2565b9150505f60135f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015490505f439050808210610e8f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e86906171f0565b60405180910390fd5b3460075f828254610ea0919061723b565b925050819055503460055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610ef3919061723b565b92505081905550610f03836136ea565b15610f1257610f11836137f4565b5b8273ffffffffffffffffffffffffffffffffffffffff167f9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d34604051610f5891906168ce565b60405180910390a2505050565b60605f60018054905067ffffffffffffffff811115610f8757610f866169a8565b5b604051908082528060200260200182016040528015610fc057816020015b610fad6165be565b815260200190600190039081610fa55790505b5090505f5b6001805490508167ffffffffffffffff16101561124c575f60135f60018467ffffffffffffffff1681548110610ffe57610ffd61726e565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015490505f60135f60018567ffffffffffffffff16815481106110875761108661726e565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff1690505f604051806080016040528060018667ffffffffffffffff16815481106111245761112361726e565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160055f60018867ffffffffffffffff16815481106111885761118761726e565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205481526020018360058111156112025761120161675a565b5b815260200184815250905080858567ffffffffffffffff168151811061122b5761122a61726e565b5b602002602001018190525050505080806112449061729b565b915050610fc5565b508091505090565b6969e10de76676d080000081565b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146112d0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112c79061733a565b60405180910390fd5b5f4390505f5b6001805490508167ffffffffffffffff161015611637575f60135f60018467ffffffffffffffff168154811061130f5761130e61726e565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015490505f60135f60018567ffffffffffffffff16815481106113985761139761726e565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff1690506003600581111561141f5761141e61675a565b5b8160058111156114325761143161675a565b5b148061146257506004600581111561144d5761144c61675a565b5b8160058111156114605761145f61675a565b5b145b156115a7578184106115a6575f60135f60018667ffffffffffffffff16815481106114905761148f61726e565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2060030181905550600560135f60018667ffffffffffffffff168154811061151b5761151a61726e565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f6101000a81548160ff021916908360058111156115a05761159f61675a565b5b02179055505b5b600260058111156115bb576115ba61675a565b5b8160058111156115ce576115cd61675a565b5b036116225761162160018467ffffffffffffffff16815481106115f4576115f361726e565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16613a61565b5b5050808061162f9061729b565b9150506112d6565b5050565b611643613fd6565b8060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167f8eef4d58d8cb1058a01a95ce964ea832ace1809b77a07c588481f74a10d3772360405160405180910390a250565b5f60055f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b61176e3373ffffffffffffffffffffffffffffffffffffffff16610bd8565b156117ae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117a5906166bc565b60405180910390fd5b6117b733614054565b806117fe57505f60055f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054115b61183d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611834906173a2565b60405180910390fd5b6118456140a6565b565b600f5481565b5f600754905090565b67016345785d8a000081565b6118813373ffffffffffffffffffffffffffffffffffffffff16610bd8565b156118c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118b8906166bc565b60405180910390fd5b6118c9610bfa565b6118d333826141a8565b50565b6118f53373ffffffffffffffffffffffffffffffffffffffff16610bd8565b15611935576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161192c906166bc565b60405180910390fd5b61193e826142a5565b61194882826141a8565b5050565b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146119ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119b19061733a565b60405180910390fd5b80601484609681106119cf576119ce61726e565b5b602091828204019190066101000a81548160ff021916908360ff160217905550600360ff1660148460968110611a0857611a0761726e565b5b602091828204019190069054906101000a900460ff1660ff1610611a3057611a2f82614a32565b5b8173ffffffffffffffffffffffffffffffffffffffff167ff833a6e5fb833592d2d6a2f115227e397239985e961011fd1cd3cec3369152574360135f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015460148760968110611abe57611abd61726e565b5b602091828204019190069054906101000a900460ff16604051611ae3939291906173f0565b60405180910390a2505050565b6103e881565b611b153373ffffffffffffffffffffffffffffffffffffffff16610bd8565b15611b55576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b4c906166bc565b60405180910390fd5b611b5e81614b01565b50565b601981565b5f60115f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b5f808290508051915050919050565b5f611c0382614054565b9050919050565b5f600854905090565b611c1b613fd6565b611c245f614c73565b565b6969e10de76676d080000081565b5f611c3e82614d34565b9050919050565b609681565b60198181548110611c59575f80fd5b905f5260205f2090600202015f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001018054611c9d90617452565b80601f0160208091040260200160405190810160405280929190818152602001828054611cc990617452565b8015611d145780601f10611ceb57610100808354040283529160200191611d14565b820191905f5260205f20905b815481529060010190602001808311611cf757829003601f168201915b5050505050905082565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60105481565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611dda576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611dd1906166bc565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff166108fc6969e10de76676d080000090811502906040515f60405180830381858888f19350505050158015611e27573d5f803e3d5ffd5b506969e10de76676d080000060055f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254611e7e9190617482565b925050819055506969e10de76676d080000060075f828254611ea09190617482565b925050819055507fb67ea9e632aa74e96032afed3f64179247898d3dfa8d9aef85057d36e9112d4681604051611ed69190616d9b565b60405180910390a150565b611f003373ffffffffffffffffffffffffffffffffffffffff16610bd8565b15611f40576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f37906166bc565b60405180910390fd5b805f60115f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205411611ffd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ff4906174ff565b60405180910390fd5b6120073383614d86565b5050565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461209a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612091906166bc565b60405180910390fd5b5f806120a5836135f2565b915091505f60055f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f6969e10de76676d080000060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205461213e9190617482565b9050816007541015612185576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161217c90617567565b60405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f193505050501580156121c8573d5f803e3d5ffd5b5060035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc8390811502906040515f60405180830381858888f1935050505015801561222d573d5f803e3d5ffd5b508160075f82825461223f9190617482565b9250508190555061224f83614f4c565b612258836155ab565b5f60055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505f600d5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505f600e5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505f600b5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505f600c5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f6782e2944a228b3021a467b1afb406079eea2c32d259933346a946601e56d2ae8482604051612469929190617585565b60405180910390a15050505050565b60095481565b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146124ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016124e39061733a565b60405180910390fd5b6124f682826156a8565b5050565b60605f60125f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208054905090505f8167ffffffffffffffff81111561255b5761255a6169a8565b5b60405190808252806020026020018201604052801561259457816020015b6125816165be565b8152602001906001900390816125795790505b5090505f5b828167ffffffffffffffff1610156127ce575f60125f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208267ffffffffffffffff16815481106126055761260461726e565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f60115f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f60135f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff1690505f60135f8a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015490505f60405180608001604052808673ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018460058111156127825761278161675a565b5b815260200183815250905080878767ffffffffffffffff16815181106127ab576127aa61726e565b5b6020026020010181905250505050505080806127c69061729b565b915050612599565b508092505050919050565b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612847576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161283e9061733a565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff167f4d381e135b58a2c3686763368fd0d7750138ab8a9a0e404a6c09e03f241003684360405161288d91906175f6565b60405180910390a250565b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612906576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128fd9061733a565b60405180910390fd5b5f5b609667ffffffffffffffff168167ffffffffffffffff161015612977575f60148267ffffffffffffffff16609681106129445761294361726e565b5b602091828204019190066101000a81548160ff021916908360ff160217905550808061296f9061729b565b915050612908565b50565b6107d081565b5f60055f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b60085481565b60606001805480602002602001604051908101604052809291908181526020018280548015612a4d57602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311612a04575b5050505050905090565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600481565b60075481565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612b3b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b32906166bc565b60405180910390fd5b6969e10de76676d08000006fffffffffffffffffffffffffffffffff16341015612b9a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b9190617692565b60405180910390fd5b5f60135f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015490505f439050808210612c24576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c1b906171f0565b60405180910390fd5b3460075f828254612c35919061723b565b925050819055503460055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254612c88919061723b565b92505081905550612c98846136ea565b15612ca757612ca6846137f4565b5b612cb184846141a8565b8373ffffffffffffffffffffffffffffffffffffffff167f9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d34604051612cf791906168ce565b60405180910390a250505050565b5f600954905090565b5f60125f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805490509050919050565b612d5f613fd6565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612dcd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612dc490617720565b60405180910390fd5b612dd681614c73565b50565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612e68576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e5f906166bc565b60405180910390fd5b6969e10de76676d08000006fffffffffffffffffffffffffffffffff16341015612ec7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612ebe90617692565b60405180910390fd5b5f60135f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015490505f429050808210612f51576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612f48906171f0565b60405180910390fd5b612f5a85614054565b8015612f6b5750612f6a84614d34565b5b156130d2578373ffffffffffffffffffffffffffffffffffffffff16600b5f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614801561308e57508473ffffffffffffffffffffffffffffffffffffffff16600c5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b6130cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016130c4906177ae565b60405180910390fd5b6131e2565b6130db85614054565b1580156130ee57506130ec84614054565b155b801561310057506130fe84614d34565b155b8015613112575061311085614d34565b155b801561315a57505f60055f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054145b80156131a257505f60055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054145b6131e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016131d890617888565b60405180910390fd5b5b6001600d5f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055506001600e5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555083600b5f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555084600c5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503460075f828254613393919061723b565b925050819055503460055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546133e6919061723b565b925050819055506133f6846136ea565b1561340557613404846137f4565b5b61340f84846141a8565b8473ffffffffffffffffffffffffffffffffffffffff167f9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d3460405161345591906168ce565b60405180910390a25050505050565b60605f609690505f8167ffffffffffffffff1667ffffffffffffffff8111156134905761348f6169a8565b5b6040519080825280602002602001820160405280156134be5781602001602082028036833780820191505090505b5090505f5b8267ffffffffffffffff168167ffffffffffffffff1610156135575760148167ffffffffffffffff16609681106134fd576134fc61726e565b5b602091828204019190069054906101000a900460ff16828267ffffffffffffffff16815181106135305761352f61726e565b5b602002602001019060ff16908160ff1681525050808061354f9061729b565b9150506134c3565b50809250505090565b600381565b60018181548110613574575f80fd5b905f5260205f20015f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f60045f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff169050919050565b5f805f8390505f84905061360585614054565b1561366e57600b5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506136dd565b61367785614d34565b156136dc57600c5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505b5b8181935093505050915091565b5f6136f482615c2b565b801561376f575060058081111561370e5761370d61675a565b5b60135f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff16600581111561376d5761376c61675a565b5b145b1561377d57600190506137ef565b61378682615c2b565b1580156137ec57506969e10de76676d08000006fffffffffffffffffffffffffffffffff1660055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410155b90505b919050565b6005808111156138075761380661675a565b5b60135f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff1660058111156138665761386561675a565b5b146139b157600954600180549050106138b4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016138ab90617916565b60405180910390fd5b600160045f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555060018054905060065f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550600181908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b600160135f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f6101000a81548160ff02191690836005811115613a1457613a1361675a565b5b02179055505f60135f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003018190555050565b5f80613a6c836135f2565b915091505f60055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905060025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ede391fb856040518263ffffffff1660e01b8152600401613b0b9190616d9b565b602060405180830381865afa158015613b26573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613b4a91906170ed565b15613ccb5760025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a209a86f856040518263ffffffff1660e01b8152600401613ba99190616d9b565b6020604051808303815f875af1158015613bc5573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613be99190617948565b90505f8103613c6257600160135f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f6101000a81548160ff02191690836005811115613c5557613c5461675a565b5b0217905550505050613fd3565b600160135f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f6101000a81548160ff02191690836005811115613cc557613cc461675a565b5b02179055505b8060055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254613d179190617482565b925050819055508060075f828254613d2f9190617482565b92505081905550613d3f84615c2b565b8015613d8757505f60055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054145b15613f3d575f600d5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505f600e5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505f600b5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505f600c5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550613f3384614f4c565b613f3c846155ab565b5b8273ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f19350505050158015613f80573d5f803e3d5ffd5b508373ffffffffffffffffffffffffffffffffffffffff167f0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f7582604051613fc791906168ce565b60405180910390a25050505b50565b613fde615c7d565b73ffffffffffffffffffffffffffffffffffffffff16613ffc611d1e565b73ffffffffffffffffffffffffffffffffffffffff1614614052576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614049906179bd565b60405180910390fd5b565b5f600d5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff169050919050565b5f806140b1336135f2565b915091505f60135f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2060030154905043811061413b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614132906171f0565b60405180910390fd5b600260135f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f6101000a81548160ff0219169083600581111561419e5761419d61675a565b5b0217905550505050565b5f6141b282611bea565b116141f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016141e990617a25565b60405180910390fd5b601960405180604001604052808473ffffffffffffffffffffffffffffffffffffffff16815260200183815250908060018154018082558091505060019003905f5260205f2090600202015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101908161429e9190617bd7565b5050505050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ede391fb336040518263ffffffff1660e01b81526004016142ff9190616d9b565b602060405180830381865afa15801561431a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061433e91906170ed565b156144a1576005808111156143565761435561675a565b5b60135f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff1660058111156143b5576143b461675a565b5b1415806143c257505f3414155b156144a05760025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663489fe625336040518263ffffffff1660e01b81526004016144219190616d9b565b602060405180830381865afa15801561443c573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061446091906170ed565b61449f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161449690617188565b60405180910390fd5b5b5b5f60135f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015490505f42905080821061452b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614522906171f0565b60405180910390fd5b61453433614054565b8015614545575061454483614d34565b5b156146ac578273ffffffffffffffffffffffffffffffffffffffff16600b5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614801561466857503373ffffffffffffffffffffffffffffffffffffffff16600c5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b6146a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161469e906177ae565b60405180910390fd5b6147bc565b6146b533614054565b1580156146c857506146c683614054565b155b80156146da57506146d883614d34565b155b80156146ec57506146ea33614d34565b155b801561473457505f60055f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054145b801561477c57505f60055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054145b6147bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016147b290617888565b60405180910390fd5b5b6001600d5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055506001600e5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555082600b5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555033600c5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503460075f82825461496d919061723b565b925050819055503460055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546149c0919061723b565b925050819055506149d0836136ea565b156149df576149de836137f4565b5b3373ffffffffffffffffffffffffffffffffffffffff167f9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d34604051614a2591906168ce565b60405180910390a2505050565b600360135f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f6101000a81548160ff02191690836005811115614a9557614a9461675a565b5b02179055506103e86fffffffffffffffffffffffffffffffff1643614aba919061723b565b60135f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003018190555050565b614b0a81615c2b565b614b49576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614b4090617d16565b60405180910390fd5b3460075f828254614b5a919061723b565b925050819055503460115f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254614bea919061723b565b92505081905550614bfb3382615c84565b15614c0b57614c0a3382615d38565b5b8073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f1ed8ad98d928651a8bc3999999b718383931f4595fcd2e1efd2de972fa8cdaea34604051614c6891906168ce565b60405180910390a350565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f600e5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff169050919050565b5f60115f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f60115f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508060075f828254614e949190617482565b92505081905550614ea58383615fae565b15614eb557614eb4838361603f565b5b8273ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f19350505050158015614ef8573d5f803e3d5ffd5b508273ffffffffffffffffffffffffffffffffffffffff167fce2e26dc37234e172764d0a8a2c04d2e6fb46c02f3ca6ca5135d6f9dcd2b67a982604051614f3f91906168ce565b60405180910390a2505050565b60085460018054905011614f95576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614f8c90617da4565b60405180910390fd5b60018054905060065f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410615019576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161501090617e0c565b60405180910390fd5b5f60065f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f6001808054905061506c9190617482565b9050808214615152575f6001828154811061508a5761508961726e565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905080600184815481106150c9576150c861726e565b5b905f5260205f20015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508260065f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f60045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505f60065f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208190555060018054806151fa576151f9617e2a565b5b600190038181905f5260205f20015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690559055600160198054905010156152415750506155a8565b5f8060016019805490506152559190617482565b90505f5b6019805490508110156152f1578573ffffffffffffffffffffffffffffffffffffffff16601982815481106152915761529061726e565b5b905f5260205f2090600202015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036152de578092505b80806152e990617e57565b915050615259565b50808214615546575f6019838154811061530e5761530d61726e565b5b905f5260205f2090600202016040518060400160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201805461538990617452565b80601f01602080910402602001604051908101604052809291908181526020018280546153b590617452565b80156154005780601f106153d757610100808354040283529160200191615400565b820191905f5260205f20905b8154815290600101906020018083116153e357829003601f168201915b50505050508152505090506019828154811061541f5761541e61726e565b5b905f5260205f2090600202016019848154811061543f5761543e61726e565b5b905f5260205f2090600202015f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600182018160010190816154c29190617eb3565b5090505080601983815481106154db576154da61726e565b5b905f5260205f2090600202015f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010190816155409190617bd7565b50905050505b601980548061555857615557617e2a565b5b600190038181905f5260205f2090600202015f8082015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600182015f61559f919061660a565b50509055505050505b50565b5f60125f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208054905090505f5b818167ffffffffffffffff1610156156a35761568d60125f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208267ffffffffffffffff168154811061565f5761565e61726e565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684614d86565b818061569890617f98565b9250505f82036155f1575b505050565b600467ffffffffffffffff168167ffffffffffffffff161180156156e15750609667ffffffffffffffff168167ffffffffffffffff1611155b156157b357600460135f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f6101000a81548160ff021916908360058111156157495761574861675a565b5b02179055506107d06fffffffffffffffffffffffffffffffff164361576e919061723b565b60135f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20600301819055505b5f60055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f606460196fffffffffffffffffffffffffffffffff16836158169190617fbf565b615820919061802d565b9050808261582e9190617482565b60055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508060075f8282546158809190617482565b925050819055505f60125f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2080548060200260200160405190810160405280929190818152602001828054801561594257602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116158f9575b505050505090505f60125f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805490501115615b92575f5b60125f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805490508167ffffffffffffffff161015615b90575f60115f848467ffffffffffffffff1681518110615a0657615a0561726e565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f606460196fffffffffffffffffffffffffffffffff1683615aa99190617fbf565b615ab3919061802d565b90508082615ac19190617482565b60115f868667ffffffffffffffff1681518110615ae157615ae061726e565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508060075f828254615b749190617482565b9250508190555050508080615b889061729b565b915050615993565b505b8473ffffffffffffffffffffffffffffffffffffffff167f5b669607b17f631e17b24518702f889e404e09aa6528fd31e931ef4d2c53ee5e4360135f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015485604051615c1c9392919061805d565b60405180910390a25050505050565b5f60045f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff169050919050565b5f33905090565b5f615c8f8383615fae565b158015615d30575067016345785d8a00006fffffffffffffffffffffffffffffffff1660115f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410155b905092915050565b60105460125f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208054905010615dbc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615db390618102565b60405180910390fd5b600160115f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555060125f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208054905060115f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208190555060125f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2082908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b5f60115f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b600f5460125f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2080549050116160c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016160ba90618190565b60405180910390fd5b60125f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208054905060115f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054106161c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016161b79061821e565b60405180910390fd5b5f60115f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f600160125f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208054905061628d9190617482565b9050808214616427575f60125f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2082815481106162e6576162e561726e565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508060125f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2084815481106163605761635f61726e565b5b905f5260205f20015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508260115f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f60115f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505f60115f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208190555060125f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2080548061658657616585617e2a565b5b600190038181905f5260205f20015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055905550505050565b60405180608001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f60058111156165fe576165fd61675a565b5b81526020015f81525090565b50805461661690617452565b5f825580601f106166275750616644565b601f0160209004905f5260205f20908101906166439190616647565b5b50565b5b8082111561665e575f815f905550600101616648565b5090565b5f82825260208201905092915050565b7f4f6e6c7920454f412063616e2063616c6c2066756e6374696f6e0000000000005f82015250565b5f6166a6601a83616662565b91506166b182616672565b602082019050919050565b5f6020820190508181035f8301526166d38161669a565b9050919050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61672c82616703565b9050919050565b61673c81616722565b82525050565b5f819050919050565b61675481616742565b82525050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b600681106167985761679761675a565b5b50565b5f8190506167a882616787565b919050565b5f6167b78261679b565b9050919050565b6167c7816167ad565b82525050565b608082015f8201516167e15f850182616733565b5060208201516167f4602085018261674b565b50604082015161680760408501826167be565b50606082015161681a606085018261674b565b50505050565b5f61682b83836167cd565b60808301905092915050565b5f602082019050919050565b5f61684d826166da565b61685781856166e4565b9350616862836166f4565b805f5b838110156168925781516168798882616820565b975061688483616837565b925050600181019050616865565b5085935050505092915050565b5f6020820190508181035f8301526168b78184616843565b905092915050565b6168c881616742565b82525050565b5f6020820190506168e15f8301846168bf565b92915050565b5f604051905090565b5f80fd5b5f80fd5b61690181616722565b811461690b575f80fd5b50565b5f8135905061691c816168f8565b92915050565b5f60208284031215616937576169366168f0565b5b5f6169448482850161690e565b91505092915050565b5f6fffffffffffffffffffffffffffffffff82169050919050565b6169718161694d565b82525050565b5f60208201905061698a5f830184616968565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6169de82616998565b810181811067ffffffffffffffff821117156169fd576169fc6169a8565b5b80604052505050565b5f616a0f6168e7565b9050616a1b82826169d5565b919050565b5f67ffffffffffffffff821115616a3a57616a396169a8565b5b616a4382616998565b9050602081019050919050565b828183375f83830152505050565b5f616a70616a6b84616a20565b616a06565b905082815260208101848484011115616a8c57616a8b616994565b5b616a97848285616a50565b509392505050565b5f82601f830112616ab357616ab2616990565b5b8135616ac3848260208601616a5e565b91505092915050565b5f60208284031215616ae157616ae06168f0565b5b5f82013567ffffffffffffffff811115616afe57616afd6168f4565b5b616b0a84828501616a9f565b91505092915050565b5f8060408385031215616b2957616b286168f0565b5b5f616b368582860161690e565b925050602083013567ffffffffffffffff811115616b5757616b566168f4565b5b616b6385828601616a9f565b9150509250929050565b616b7681616742565b8114616b80575f80fd5b50565b5f81359050616b9181616b6d565b92915050565b5f60ff82169050919050565b616bac81616b97565b8114616bb6575f80fd5b50565b5f81359050616bc781616ba3565b92915050565b5f805f60608486031215616be457616be36168f0565b5b5f616bf186828701616b83565b9350506020616c028682870161690e565b9250506040616c1386828701616bb9565b9150509250925092565b5f8060408385031215616c3357616c326168f0565b5b5f616c408582860161690e565b9250506020616c518582860161690e565b9150509250929050565b5f8115159050919050565b616c6f81616c5b565b82525050565b5f602082019050616c885f830184616c66565b92915050565b5f67ffffffffffffffff82169050919050565b616caa81616c8e565b82525050565b5f602082019050616cc35f830184616ca1565b92915050565b5f60208284031215616cde57616cdd6168f0565b5b5f616ceb84828501616b83565b91505092915050565b616cfd81616722565b82525050565b5f81519050919050565b5f5b83811015616d2a578082015181840152602081019050616d0f565b5f8484015250505050565b5f616d3f82616d03565b616d498185616662565b9350616d59818560208601616d0d565b616d6281616998565b840191505092915050565b5f604082019050616d805f830185616cf4565b8181036020830152616d928184616d35565b90509392505050565b5f602082019050616dae5f830184616cf4565b92915050565b616dbd81616c8e565b8114616dc7575f80fd5b50565b5f81359050616dd881616db4565b92915050565b5f8060408385031215616df457616df36168f0565b5b5f616e018582860161690e565b9250506020616e1285828601616dca565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f616e508383616733565b60208301905092915050565b5f602082019050919050565b5f616e7282616e1c565b616e7c8185616e26565b9350616e8783616e36565b805f5b83811015616eb7578151616e9e8882616e45565b9750616ea983616e5c565b925050600181019050616e8a565b5085935050505092915050565b5f6020820190508181035f830152616edc8184616e68565b905092915050565b5f819050919050565b5f616f07616f02616efd84616703565b616ee4565b616703565b9050919050565b5f616f1882616eed565b9050919050565b5f616f2982616f0e565b9050919050565b616f3981616f1f565b82525050565b5f602082019050616f525f830184616f30565b92915050565b5f805f60608486031215616f6f57616f6e6168f0565b5b5f616f7c8682870161690e565b9350506020616f8d8682870161690e565b925050604084013567ffffffffffffffff811115616fae57616fad6168f4565b5b616fba86828701616a9f565b9150509250925092565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b616ff681616b97565b82525050565b5f6170078383616fed565b60208301905092915050565b5f602082019050919050565b5f61702982616fc4565b6170338185616fce565b935061703e83616fde565b805f5b8381101561706e5781516170558882616ffc565b975061706083617013565b925050600181019050617041565b5085935050505092915050565b5f6020820190508181035f830152617093818461701f565b905092915050565b6170a481616b97565b82525050565b5f6020820190506170bd5f83018461709b565b92915050565b6170cc81616c5b565b81146170d6575f80fd5b50565b5f815190506170e7816170c3565b92915050565b5f60208284031215617102576171016168f0565b5b5f61710f848285016170d9565b91505092915050565b7f706c656173652070617920746865206772616e74656420616d6f756e742062655f8201527f666f7265207374616b696e6720616761696e0000000000000000000000000000602082015250565b5f617172603283616662565b915061717d82617118565b604082019050919050565b5f6020820190508181035f83015261719f81617166565b9050919050565b7f7374616b65722069732062616e6e6564206f722073757370656e6465640000005f82015250565b5f6171da601d83616662565b91506171e5826171a6565b602082019050919050565b5f6020820190508181035f830152617207816171ce565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61724582616742565b915061725083616742565b92508282019050808211156172685761726761720e565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f6172a582616c8e565b915067ffffffffffffffff82036172bf576172be61720e565b5b600182019050919050565b7f4f6e6c7920626c6f636b2070726f706f7365722063616e2063616c6c2066756e5f8201527f6374696f6e000000000000000000000000000000000000000000000000000000602082015250565b5f617324602583616662565b915061732f826172ca565b604082019050919050565b5f6020820190508181035f83015261735181617318565b9050919050565b7f4f6e6c79207374616b65722063616e2063616c6c2066756e6374696f6e0000005f82015250565b5f61738c601d83616662565b915061739782617358565b602082019050919050565b5f6020820190508181035f8301526173b981617380565b9050919050565b5f6173da6173d56173d084616b97565b616ee4565b616c8e565b9050919050565b6173ea816173c0565b82525050565b5f6060820190506174035f8301866168bf565b61741060208301856168bf565b61741d60408301846173e1565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061746957607f821691505b60208210810361747c5761747b617425565b5b50919050565b5f61748c82616742565b915061749783616742565b92508282039050818111156174af576174ae61720e565b5b92915050565b7f6f6e6c792064656c656761746f722063616e2063616c6c2066756e6374696f6e5f82015250565b5f6174e9602083616662565b91506174f4826174b5565b602082019050919050565b5f6020820190508181035f830152617516816174dd565b9050919050565b7f6e6f7420656e6f756768206d6f6e6579000000000000000000000000000000005f82015250565b5f617551601083616662565b915061755c8261751d565b602082019050919050565b5f6020820190508181035f83015261757e81617545565b9050919050565b5f6040820190506175985f830185616cf4565b6175a560208301846168bf565b9392505050565b7f74696d656f7574000000000000000000000000000000000000000000000000005f82015250565b5f6175e0600783616662565b91506175eb826175ac565b602082019050919050565b5f6040820190506176095f8301846168bf565b818103602083015261761a816175d4565b905092915050565b7f416d6f756e74206d7573742062652067726561746572207468616e2076616c695f8201527f6461746f72207468726573686f6c640000000000000000000000000000000000602082015250565b5f61767c602f83616662565b915061768782617622565b604082019050919050565b5f6020820190508181035f8301526176a981617670565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f61770a602683616662565b9150617715826176b0565b604082019050919050565b5f6020820190508181035f830152617737816176fe565b9050919050565b7f7468652073656e646572206973206e6f7420612076616c6964207374616b65725f8201527f206f662074686973207369676e65720000000000000000000000000000000000602082015250565b5f617798602f83616662565b91506177a38261773e565b604082019050919050565b5f6020820190508181035f8301526177c58161778c565b9050919050565b7f626f74682061646472657373657320617265206e6f742076616c696420746f205f8201527f63616c6c2066756e6374696f6e2e20496620796f75207374696c6c2077616e7460208201527f20746f2063616c6c2e204a7573742063616c6c696e6720756e7374616b65206660408201527f6f7220626f7468206163636f756e7473206d616e75616c6c7900000000000000606082015250565b5f617872607983616662565b915061787d826177cc565b608082019050919050565b5f6020820190508181035f83015261789f81617866565b9050919050565b7f56616c696461746f72207365742068617320726561636865642066756c6c20635f8201527f6170616369747900000000000000000000000000000000000000000000000000602082015250565b5f617900602783616662565b915061790b826178a6565b604082019050919050565b5f6020820190508181035f83015261792d816178f4565b9050919050565b5f8151905061794281616b6d565b92915050565b5f6020828403121561795d5761795c6168f0565b5b5f61796a84828501617934565b91505092915050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f6179a7602083616662565b91506179b282617973565b602082019050919050565b5f6020820190508181035f8301526179d48161799b565b9050919050565b7f506c6561736520736574207468652063726577206e616d6500000000000000005f82015250565b5f617a0f601883616662565b9150617a1a826179db565b602082019050919050565b5f6020820190508181035f830152617a3c81617a03565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302617a9f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82617a64565b617aa98683617a64565b95508019841693508086168417925050509392505050565b5f617adb617ad6617ad184616742565b616ee4565b616742565b9050919050565b5f819050919050565b617af483617ac1565b617b08617b0082617ae2565b848454617a70565b825550505050565b5f90565b617b1c617b10565b617b27818484617aeb565b505050565b5b81811015617b4a57617b3f5f82617b14565b600181019050617b2d565b5050565b601f821115617b8f57617b6081617a43565b617b6984617a55565b81016020851015617b78578190505b617b8c617b8485617a55565b830182617b2c565b50505b505050565b5f82821c905092915050565b5f617baf5f1984600802617b94565b1980831691505092915050565b5f617bc78383617ba0565b9150826002028217905092915050565b617be082616d03565b67ffffffffffffffff811115617bf957617bf86169a8565b5b617c038254617452565b617c0e828285617b4e565b5f60209050601f831160018114617c3f575f8415617c2d578287015190505b617c378582617bbc565b865550617c9e565b601f198416617c4d86617a43565b5f5b82811015617c7457848901518255600182019150602085019450602081019050617c4f565b86831015617c915784890151617c8d601f891682617ba0565b8355505b6001600288020188555050505b505050505050565b7f63616e6e6f742064656c656761746520746f6b656e7320746f20736f6d656f6e5f8201527f652077686f206973206e6f742076616c696461746f7200000000000000000000602082015250565b5f617d00603683616662565b9150617d0b82617ca6565b604082019050919050565b5f6020820190508181035f830152617d2d81617cf4565b9050919050565b7f56616c696461746f72732063616e2774206265206c657373207468616e2074685f8201527f65206d696e696d756d2072657175697265642076616c696461746f72206e756d602082015250565b5f617d8e604083616662565b9150617d9982617d34565b604082019050919050565b5f6020820190508181035f830152617dbb81617d82565b9050919050565b7f696e646578206f7574206f662072616e676500000000000000000000000000005f82015250565b5f617df6601283616662565b9150617e0182617dc2565b602082019050919050565b5f6020820190508181035f830152617e2381617dea565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b5f617e6182616742565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203617e9357617e9261720e565b5b600182019050919050565b5f81549050617eac81617452565b9050919050565b818103617ec1575050617f96565b617eca82617e9e565b67ffffffffffffffff811115617ee357617ee26169a8565b5b617eed8254617452565b617ef8828285617b4e565b5f601f831160018114617f25575f8415617f13578287015490505b617f1d8582617bbc565b865550617f8f565b601f198416617f3387617a43565b9650617f3e86617a43565b5f5b82811015617f6557848901548255600182019150600185019450602081019050617f40565b86831015617f825784890154617f7e601f891682617ba0565b8355505b6001600288020188555050505b5050505050505b565b5f617fa282616742565b91505f8203617fb457617fb361720e565b5b600182039050919050565b5f617fc982616742565b9150617fd483616742565b9250828202617fe281616742565b91508282048414831517617ff957617ff861720e565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61803782616742565b915061804283616742565b92508261805257618051618000565b5b828204905092915050565b5f6060820190506180705f8301866168bf565b61807d60208301856168bf565b61808a60408301846168bf565b949350505050565b7f64656c656761746f7220736574206861732072656163686564206974732066755f8201527f6c6c206361706163697479000000000000000000000000000000000000000000602082015250565b5f6180ec602b83616662565b91506180f782618092565b604082019050919050565b5f6020820190508181035f830152618119816180e0565b9050919050565b7f64656c656761746f72732073697a652063616e2774206265206c6573732074685f8201527f616e20746865206d696e696d756d20726571756972656d656e74000000000000602082015250565b5f61817a603a83616662565b915061818582618120565b604082019050919050565b5f6020820190508181035f8301526181a78161816e565b9050919050565b7f696e646578206f7574206f662072616e676520696e207468652064656c6567615f8201527f746f7273206c697374206f662076616c696461746f7200000000000000000000602082015250565b5f618208603683616662565b9150618213826181ae565b604082019050919050565b5f6020820190508181035f830152618235816181fc565b905091905056fea26469706673582212204babe23fd75555266b9cf02c80f8a4ebf2da61e01f1bf4911f7ce62cc27821a464736f6c63430008140033",
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
