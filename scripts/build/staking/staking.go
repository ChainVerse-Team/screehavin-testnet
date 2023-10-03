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
	_ = abi.ConvertType
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minNumValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNumValidators\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"grantContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeLock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalTokenSlashing\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalStakedSize\",\"type\":\"uint256\"}],\"name\":\"Ban\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"BurnGrantInitial\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelegatorStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelegatorUnstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_grantContract\",\"type\":\"address\"}],\"name\":\"GrantContractChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"earnedAmount\",\"type\":\"uint256\"}],\"name\":\"ReclaimGrant\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeLock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"suspendCounter\",\"type\":\"uint64\"}],\"name\":\"Suspended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"Warning\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BAN_DURATION\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COUNTER_SUSPEND\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DELEGATION_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GrantAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_VALIDATORSUBSET_SIZE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_VALIDATORSUBSET_SIZE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERCENTAGE_TOKEN_SLASHING\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUSPEND_DURATION\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_SLASHING_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_crewName\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"boss\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_grantContract\",\"outputs\":[{\"internalType\":\"contractGrantContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_maximumNumDelegators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_maximumNumValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_minimumNumDelegators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_minimumNumValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_stakedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_validators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"accountStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressGrantContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"addressToStakedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"badValidator\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"validatorsSubsetSize\",\"type\":\"uint64\"}],\"name\":\"ban\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"burnGrantInitial\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"canBecomeBoss\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkStateValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToUnstake\",\"type\":\"uint256\"}],\"name\":\"delegatorUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"getDelegatedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"getDelegatorsInfoOfValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumStaking.Status\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timeLock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeAmount\",\"type\":\"uint256\"}],\"internalType\":\"structStaking.AccountStakingInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"getStringLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getTotalDelegatorOfValidator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"getTotalDeligatedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorsStakeInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumStaking.Status\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timeLock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeAmount\",\"type\":\"uint256\"}],\"internalType\":\"structStaking.AccountStakingInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorsSubsetTimeout\",\"outputs\":[{\"internalType\":\"uint8[]\",\"name\":\"\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"isDelegator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isStaker\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maximumNumValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumNumValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"reclaimUnusedGrant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetValidatorsSubsetTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"setGrantContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"stakeGrantContract\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"stakeSignerGrantContract\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"indexValidator\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"badValidator\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"count\",\"type\":\"uint8\"}],\"name\":\"suspend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToUnstake\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"badValidator\",\"type\":\"address\"}],\"name\":\"warning\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040525f600f55610bb86010553480156200001a575f80fd5b5060405162009165380380620091658339818101604052810190620000409190620002a2565b62000060620000546200013d60201b60201c565b6200014460201b60201c565b81831115620000a6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200009d906200037f565b60405180910390fd5b82600881905550816009819055508060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050506200039f565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f80fd5b5f819050919050565b6200021d8162000209565b811462000228575f80fd5b50565b5f815190506200023b8162000212565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6200026c8262000241565b9050919050565b6200027e8162000260565b811462000289575f80fd5b50565b5f815190506200029c8162000273565b92915050565b5f805f60608486031215620002bc57620002bb62000205565b5b5f620002cb868287016200022b565b9350506020620002de868287016200022b565b9250506040620002f1868287016200028c565b9150509250925092565b5f82825260208201905092915050565b7f4d696e2076616c696461746f7273206e756d2063616e206e6f742062652067725f8201527f6561746572207468616e206d6178206e756d206f662076616c696461746f7273602082015250565b5f62000367604083620002fb565b915062000374826200030b565b604082019050919050565b5f6020820190508181035f830152620003988162000359565b9050919050565b618db880620003ad5f395ff3fe608060405260043610610353575f3560e01c806390b19627116101c5578063cb3d89d6116100f6578063ed546c8d11610094578063f6676a3a1161006e578063f6676a3a14610c68578063f740ff8e14610c92578063f90ecacc14610cbc578063facd743b14610cf8576103c1565b8063ed546c8d14610be8578063f2fde38b14610c24578063f3a382f214610c4c576103c1565b8063e387a7ed116100d0578063e387a7ed14610b4e578063e69af0c714610b78578063e70c91b514610ba2578063e804fbf614610bbe576103c1565b8063cb3d89d614610abe578063de092ac414610ae8578063e069fe9914610b12576103c1565b8063b1f72b5211610163578063bb4c461e1161013d578063bb4c461e14610a04578063c3a6ce2b14610a2e578063c795c07714610a6a578063ca1e781914610a94576103c1565b8063b1f72b521461098a578063b788e44e146109c6578063ba1b2125146109ee576103c1565b8063a157339c1161019f578063a157339c146108e6578063a5dbc13014610910578063af6da36e14610938578063b0f572ee14610962576103c1565b806390b196271461085857806396f58ffa14610882578063a0e11929146108aa576103c1565b80635bb4f9981161029f578063714ff4251161023d5780637df73e27116102175780637df73e271461078b5780637e840377146107c75780637ebe444c146107f15780638da5cb5b1461082e576103c1565b8063714ff42514610721578063715018a61461074b5780637a6eea3714610761576103c1565b80635c8c4a07116102795780635c8c4a0714610643578063636592ed1461066d57806365c19af0146106a95780636f1e8533146106e5576103c1565b80635bb4f998146105d55780635bb9fec9146105fd5780635c19a95c14610627576103c1565b8063306c019d1161030c5780633c449ae7116102e65780633c449ae71461054b57806346f45b8d1461057557806348cb4d31146105915780634fb856e5146105b9576103c1565b8063306c019d146104bb57806336c4086f146104f7578063373d613214610521576103c1565b806303306f0a146103c557806315f9eb9b146103ef57806319fd547d1461041957806322d468521461042f5780632367f6b5146104575780632e17de7814610493576103c1565b366103c1576103773373ffffffffffffffffffffffffffffffffffffffff16610d34565b156103b7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ae9061717d565b60405180910390fd5b6103bf610d56565b005b5f80fd5b3480156103d0575f80fd5b506103d96110c1565b6040516103e69190617373565b60405180910390f35b3480156103fa575f80fd5b50610403611440565b60405161041091906173a2565b60405180910390f35b348015610424575f80fd5b5061042d61144e565b005b34801561043a575f80fd5b50610455600480360381019061045091906173f6565b611827565b005b348015610462575f80fd5b5061047d600480360381019061047891906173f6565b6118f5565b60405161048a91906173a2565b60405180910390f35b34801561049e575f80fd5b506104b960048036038101906104b4919061744b565b61193b565b005b3480156104c6575f80fd5b506104e160048036038101906104dc91906173f6565b611a35565b6040516104ee91906173a2565b60405180910390f35b348015610502575f80fd5b5061050b611c68565b60405161051891906173a2565b60405180910390f35b34801561052c575f80fd5b50610535611c6e565b60405161054291906173a2565b60405180910390f35b348015610556575f80fd5b5061055f611c77565b60405161056c91906174a0565b60405180910390f35b61058f600480360381019061058a91906175f5565b611c84565b005b34801561059c575f80fd5b506105b760048036038101906105b2919061763c565b611d07565b005b6105d360048036038101906105ce919061767a565b611e33565b005b3480156105e0575f80fd5b506105fb60048036038101906105f6919061770a565b611eb8565b005b348015610608575f80fd5b5061061161205c565b60405161061e91906174a0565b60405180910390f35b610641600480360381019061063c91906173f6565b612062565b005b34801561064e575f80fd5b506106576120cd565b60405161066491906174a0565b60405180910390f35b348015610678575f80fd5b50610693600480360381019061068e919061775a565b6120d2565b6040516106a091906173a2565b60405180910390f35b3480156106b4575f80fd5b506106cf60048036038101906106ca91906175f5565b612156565b6040516106dc91906173a2565b60405180910390f35b3480156106f0575f80fd5b5061070b600480360381019061070691906173f6565b612165565b60405161071891906177b2565b60405180910390f35b34801561072c575f80fd5b50610735612176565b60405161074291906173a2565b60405180910390f35b348015610756575f80fd5b5061075f61217f565b005b34801561076c575f80fd5b50610775612192565b60405161078291906174a0565b60405180910390f35b348015610796575f80fd5b506107b160048036038101906107ac91906173f6565b6121a0565b6040516107be91906177b2565b60405180910390f35b3480156107d2575f80fd5b506107db6121b1565b6040516107e891906177ed565b60405180910390f35b3480156107fc575f80fd5b506108176004803603810190610812919061744b565b6121b6565b60405161082592919061787f565b60405180910390f35b348015610839575f80fd5b5061084261228a565b60405161084f91906178ad565b60405180910390f35b348015610863575f80fd5b5061086c6122b1565b60405161087991906173a2565b60405180910390f35b34801561088d575f80fd5b506108a860048036038101906108a391906173f6565b6122b7565b005b3480156108b5575f80fd5b506108d060048036038101906108cb919061775a565b61244d565b6040516108dd91906177b2565b60405180910390f35b3480156108f1575f80fd5b506108fa6124de565b60405161090791906174a0565b60405180910390f35b34801561091b575f80fd5b50610936600480360381019061093191906173f6565b6124ec565b005b348015610943575f80fd5b5061094c612959565b60405161095991906173a2565b60405180910390f35b34801561096d575f80fd5b50610988600480360381019061098391906178f0565b61295f565b005b348015610995575f80fd5b506109b060048036038101906109ab91906173f6565b6129db565b6040516109bd9190617373565b60405180910390f35b3480156109d1575f80fd5b506109ec60048036038101906109e791906173f6565b612d05565b005b3480156109f9575f80fd5b50610a02612dc4565b005b348015610a0f575f80fd5b50610a18612ea6565b604051610a2591906174a0565b60405180910390f35b348015610a39575f80fd5b50610a546004803603810190610a4f91906173f6565b612eac565b604051610a6191906173a2565b60405180910390f35b348015610a75575f80fd5b50610a7e612ef2565b604051610a8b91906173a2565b60405180910390f35b348015610a9f575f80fd5b50610aa8612ef8565b604051610ab591906179d6565b60405180910390f35b348015610ac9575f80fd5b50610ad2612f83565b604051610adf9190617a51565b60405180910390f35b348015610af3575f80fd5b50610afc612fa8565b604051610b0991906177ed565b60405180910390f35b348015610b1d575f80fd5b50610b386004803603810190610b3391906173f6565b612fad565b604051610b4591906177b2565b60405180910390f35b348015610b59575f80fd5b50610b6261306f565b604051610b6f91906173a2565b60405180910390f35b348015610b83575f80fd5b50610b8c613075565b604051610b9991906178ad565b60405180910390f35b610bbc6004803603810190610bb7919061767a565b61309a565b005b348015610bc9575f80fd5b50610bd2613302565b604051610bdf91906173a2565b60405180910390f35b348015610bf3575f80fd5b50610c0e6004803603810190610c0991906173f6565b61330b565b604051610c1b91906173a2565b60405180910390f35b348015610c2f575f80fd5b50610c4a6004803603810190610c4591906173f6565b613354565b005b610c666004803603810190610c619190617a6a565b6133d6565b005b348015610c73575f80fd5b50610c7c613a70565b604051610c899190617b8d565b60405180910390f35b348015610c9d575f80fd5b50610ca6613b6c565b604051610cb39190617bbc565b60405180910390f35b348015610cc7575f80fd5b50610ce26004803603810190610cdd919061744b565b613b71565b604051610cef91906178ad565b60405180910390f35b348015610d03575f80fd5b50610d1e6004803603810190610d1991906173f6565b613bac565b604051610d2b91906177b2565b60405180910390f35b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ede391fb336040518263ffffffff1660e01b8152600401610db091906178ad565b602060405180830381865afa158015610dcb573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610def9190617bff565b8015610e765750600580811115610e0957610e0861721b565b5b60135f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff166005811115610e6857610e6761721b565b5b141580610e7557505f3414155b5b15610f545760025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663489fe625336040518263ffffffff1660e01b8152600401610ed591906178ad565b602060405180830381865afa158015610ef0573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f149190617bff565b610f53576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f4a90617c9a565b60405180910390fd5b5b5f610f5e33613bfe565b9150505f60135f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015490505f439050808210610feb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fe290617d02565b60405180910390fd5b3460075f828254610ffc9190617d4d565b925050819055503460055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461104f9190617d4d565b9250508190555061105f83613cf6565b1561106e5761106d83613e00565b5b8273ffffffffffffffffffffffffffffffffffffffff167f9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d346040516110b491906173a2565b60405180910390a2505050565b60605f60018054905067ffffffffffffffff8111156110e3576110e26174d1565b5b60405190808252806020026020018201604052801561111c57816020015b611109617079565b8152602001906001900390816111015790505b5090505f5b6001805490508167ffffffffffffffff161015611438575f60135f60018467ffffffffffffffff168154811061115a57611159617d80565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015490505f60135f60018567ffffffffffffffff16815481106111e3576111e2617d80565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff1690505f60135f60018667ffffffffffffffff168154811061127857611277617d80565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206004015490505f6040518060a0016040528060018767ffffffffffffffff168154811061130957611308617d80565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160055f60018967ffffffffffffffff168154811061136d5761136c617d80565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205481526020018460058111156113e7576113e661721b565b5b815260200185815260200183815250905080868667ffffffffffffffff168151811061141657611415617d80565b5b602002602001018190525050505050808061143090617dad565b915050611121565b508091505090565b6969e10de76676d080000081565b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146114bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114b390617e4c565b60405180910390fd5b5f4390505f5b6001805490508167ffffffffffffffff161015611823575f60135f60018467ffffffffffffffff16815481106114fb576114fa617d80565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015490505f60135f60018567ffffffffffffffff168154811061158457611583617d80565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff1690506003600581111561160b5761160a61721b565b5b81600581111561161e5761161d61721b565b5b148061164e5750600460058111156116395761163861721b565b5b81600581111561164c5761164b61721b565b5b145b1561179357818410611792575f60135f60018667ffffffffffffffff168154811061167c5761167b617d80565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2060030181905550600560135f60018667ffffffffffffffff168154811061170757611706617d80565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f6101000a81548160ff0219169083600581111561178c5761178b61721b565b5b02179055505b5b600260058111156117a7576117a661721b565b5b8160058111156117ba576117b961721b565b5b0361180e5761180d60018467ffffffffffffffff16815481106117e0576117df617d80565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1661406d565b5b5050808061181b90617dad565b9150506114c2565b5050565b61182f614857565b8060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167f8eef4d58d8cb1058a01a95ce964ea832ace1809b77a07c588481f74a10d3772360405160405180910390a250565b5f60055f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b61195a3373ffffffffffffffffffffffffffffffffffffffff16610d34565b1561199a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119919061717d565b60405180910390fd5b6119a3336148d5565b806119ea57505f60055f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054115b611a29576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a2090617eb4565b60405180910390fd5b611a3281614927565b50565b5f8060125f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805490501115611c5f575f8060125f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805480602002602001604051908101604052809291908181526020018280548015611b3a57602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611af1575b505050505090505f5b60125f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805490508167ffffffffffffffff161015611c545760115f838367ffffffffffffffff1681518110611bb557611bb4617d80565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205483611c3f9190617d4d565b92508080611c4c90617dad565b915050611b43565b508192505050611c63565b5f90505b919050565b600f5481565b5f600754905090565b68056bc75e2d6310000081565b611ca33373ffffffffffffffffffffffffffffffffffffffff16610d34565b15611ce3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611cda9061717d565b60405180910390fd5b611ceb610d56565b611cf433612fad565b15611d0457611d033382614a6f565b5b50565b611d263373ffffffffffffffffffffffffffffffffffffffff16610d34565b15611d66576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d5d9061717d565b60405180910390fd5b815f60115f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205411611e23576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e1a90617f1c565b60405180910390fd5b611e2e338484614b6c565b505050565b611e523373ffffffffffffffffffffffffffffffffffffffff16610d34565b15611e92576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e899061717d565b60405180910390fd5b611e9b82614de0565b611ea482612fad565b15611eb457611eb38282614a6f565b5b5050565b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611f26576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f1d90617e4c565b60405180910390fd5b8060148460968110611f3b57611f3a617d80565b5b602091828204019190066101000a81548160ff021916908360ff160217905550600360ff1660148460968110611f7457611f73617d80565b5b602091828204019190069054906101000a900460ff1660ff1610611f9c57611f9b8261556d565b5b8173ffffffffffffffffffffffffffffffffffffffff167ff833a6e5fb833592d2d6a2f115227e397239985e961011fd1cd3cec3369152574360135f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20600301546014876096811061202a57612029617d80565b5b602091828204019190069054906101000a900460ff1660405161204f93929190617f6a565b60405180910390a2505050565b6103e881565b6120813373ffffffffffffffffffffffffffffffffffffffff16610d34565b156120c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120b89061717d565b60405180910390fd5b6120ca8161563c565b50565b601981565b5f60115f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b5f808290508051915050919050565b5f61216f826148d5565b9050919050565b5f600854905090565b612187614857565b6121905f6157ae565b565b6969e10de76676d080000081565b5f6121aa8261586f565b9050919050565b609681565b601981815481106121c5575f80fd5b905f5260205f2090600202015f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600101805461220990617fcc565b80601f016020809104026020016040519081016040528092919081815260200182805461223590617fcc565b80156122805780601f1061225757610100808354040283529160200191612280565b820191905f5260205f20905b81548152906001019060200180831161226357829003601f168201915b5050505050905082565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60105481565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612346576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161233d9061717d565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff166108fc6969e10de76676d080000090811502906040515f60405180830381858888f19350505050158015612393573d5f803e3d5ffd5b506969e10de76676d080000060055f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546123ea9190617ffc565b925050819055506969e10de76676d080000060075f82825461240c9190617ffc565b925050819055507fb67ea9e632aa74e96032afed3f64179247898d3dfa8d9aef85057d36e9112d468160405161244291906178ad565b60405180910390a150565b5f60115f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b691fc3842bd1f071c0000081565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461257b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125729061717d565b60405180910390fd5b5f8061258683613bfe565b915091505f60055f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f6969e10de76676d080000060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205461261f9190617ffc565b9050816007541015612666576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161265d90618079565b60405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f193505050501580156126a9573d5f803e3d5ffd5b5060035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc8390811502906040515f60405180830381858888f1935050505015801561270e573d5f803e3d5ffd5b508160075f8282546127209190617ffc565b92505081905550612730836158c1565b61273983615f20565b5f60055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505f600d5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505f600e5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505f600b5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505f600c5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f6782e2944a228b3021a467b1afb406079eea2c32d259933346a946601e56d2ae848260405161294a929190618097565b60405180910390a15050505050565b60095481565b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146129cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016129c490617e4c565b60405180910390fd5b6129d7828261601e565b5050565b60605f60125f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208054905090505f8167ffffffffffffffff811115612a3c57612a3b6174d1565b5b604051908082528060200260200182016040528015612a7557816020015b612a62617079565b815260200190600190039081612a5a5790505b5090505f5b828167ffffffffffffffff161015612cfa575f60125f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208267ffffffffffffffff1681548110612ae657612ae5617d80565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f60115f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f60135f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff1690505f60135f8a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015490505f60135f8b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206004015490505f6040518060a001604052808773ffffffffffffffffffffffffffffffffffffffff168152602001868152602001856005811115612ca757612ca661721b565b5b815260200184815260200183815250905080888867ffffffffffffffff1681518110612cd657612cd5617d80565b5b60200260200101819052505050505050508080612cf290617dad565b915050612a7a565b508092505050919050565b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612d73576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d6a90617e4c565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff167f4d381e135b58a2c3686763368fd0d7750138ab8a9a0e404a6c09e03f2410036843604051612db99190618108565b60405180910390a250565b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612e32576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e2990617e4c565b60405180910390fd5b5f5b609667ffffffffffffffff168167ffffffffffffffff161015612ea3575f60148267ffffffffffffffff1660968110612e7057612e6f617d80565b5b602091828204019190066101000a81548160ff021916908360ff1602179055508080612e9b90617dad565b915050612e34565b50565b6107d081565b5f60055f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b60085481565b60606001805480602002602001604051908101604052809291908181526020018280548015612f7957602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311612f30575b5050505050905090565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600481565b5f805f90505b6019805490508167ffffffffffffffff161015613064578273ffffffffffffffffffffffffffffffffffffffff1660198267ffffffffffffffff1681548110612fff57612ffe617d80565b5b905f5260205f2090600202015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603613051575f91505061306a565b808061305c90617dad565b915050612fb3565b50600190505b919050565b60075481565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614613129576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016131209061717d565b60405180910390fd5b6969e10de76676d08000006fffffffffffffffffffffffffffffffff16341015613188576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161317f906181a4565b60405180910390fd5b5f60135f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015490505f439050808210613212576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161320990617d02565b60405180910390fd5b3460075f8282546132239190617d4d565b925050819055503460055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546132769190617d4d565b9250508190555061328684613cf6565b156132955761329484613e00565b5b61329e84612fad565b156132ae576132ad8484614a6f565b5b8373ffffffffffffffffffffffffffffffffffffffff167f9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d346040516132f491906173a2565b60405180910390a250505050565b5f600954905090565b5f60125f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805490509050919050565b61335c614857565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036133ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016133c190618232565b60405180910390fd5b6133d3816157ae565b50565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614613465576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161345c9061717d565b60405180910390fd5b6969e10de76676d08000006fffffffffffffffffffffffffffffffff163410156134c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016134bb906181a4565b60405180910390fd5b5f60135f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015490505f42905080821061354e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161354590617d02565b60405180910390fd5b613557856148d5565b801561356857506135678461586f565b5b156136cf578373ffffffffffffffffffffffffffffffffffffffff16600b5f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614801561368b57508473ffffffffffffffffffffffffffffffffffffffff16600c5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b6136ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016136c1906182c0565b60405180910390fd5b6137df565b6136d8856148d5565b1580156136eb57506136e9846148d5565b155b80156136fd57506136fb8461586f565b155b801561370f575061370d8561586f565b155b801561375757505f60055f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054145b801561379f57505f60055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054145b6137de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016137d59061839a565b60405180910390fd5b5b6001600d5f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055506001600e5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555083600b5f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555084600c5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503460075f8282546139909190617d4d565b925050819055503460055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546139e39190617d4d565b925050819055506139f384613cf6565b15613a0257613a0184613e00565b5b613a0b84612fad565b15613a1b57613a1a8484614a6f565b5b8473ffffffffffffffffffffffffffffffffffffffff167f9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d34604051613a6191906173a2565b60405180910390a25050505050565b60605f609690505f8167ffffffffffffffff1667ffffffffffffffff811115613a9c57613a9b6174d1565b5b604051908082528060200260200182016040528015613aca5781602001602082028036833780820191505090505b5090505f5b8267ffffffffffffffff168167ffffffffffffffff161015613b635760148167ffffffffffffffff1660968110613b0957613b08617d80565b5b602091828204019190069054906101000a900460ff16828267ffffffffffffffff1681518110613b3c57613b3b617d80565b5b602002602001019060ff16908160ff16815250508080613b5b90617dad565b915050613acf565b50809250505090565b600381565b60018181548110613b80575f80fd5b905f5260205f20015f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f60045f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff169050919050565b5f805f8390505f849050613c11856148d5565b15613c7a57600b5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050613ce9565b613c838561586f565b15613ce857600c5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505b5b8181935093505050915091565b5f613d0082616776565b8015613d7b5750600580811115613d1a57613d1961721b565b5b60135f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff166005811115613d7957613d7861721b565b5b145b15613d895760019050613dfb565b613d9282616776565b158015613df857506969e10de76676d08000006fffffffffffffffffffffffffffffffff1660055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410155b90505b919050565b600580811115613e1357613e1261721b565b5b60135f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff166005811115613e7257613e7161721b565b5b14613fbd5760095460018054905010613ec0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613eb790618428565b60405180910390fd5b600160045f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555060018054905060065f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550600181908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b600160135f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f6101000a81548160ff021916908360058111156140205761401f61721b565b5b02179055505f60135f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003018190555050565b5f8061407883613bfe565b915091505f60135f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2060040154905060025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ede391fb856040518263ffffffff1660e01b815260040161411a91906178ad565b602060405180830381865afa158015614135573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906141599190617bff565b80156141fb575060025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633700c315856040518263ffffffff1660e01b81526004016141ba91906178ad565b602060405180830381865afa1580156141d5573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906141f99190617bff565b155b156144095760025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a209a86f856040518263ffffffff1660e01b815260040161425a91906178ad565b6020604051808303815f875af1158015614276573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061429a919061845a565b90505f810361431357600160135f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f6101000a81548160ff021916908360058111156143065761430561721b565b5b0217905550505050614854565b60055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205481146143c457600160135f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f6101000a81548160ff021916908360058111156143ba576143b961721b565b5b0217905550614408565b60135f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206004015490505b5b5f811480614454575060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20548110155b1561449e5760055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050614507565b600160135f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f6101000a81548160ff021916908360058111156145015761450061721b565b5b02179055505b8060055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546145539190617ffc565b925050819055508060075f82825461456b9190617ffc565b925050819055505f60135f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20600401819055506145c084616776565b801561460857505f60055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054145b156147be575f600d5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505f600e5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505f600b5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505f600c5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506147b4846158c1565b6147bd84615f20565b5b8273ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f19350505050158015614801573d5f803e3d5ffd5b508373ffffffffffffffffffffffffffffffffffffffff167f0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f758260405161484891906173a2565b60405180910390a25050505b50565b61485f6167c8565b73ffffffffffffffffffffffffffffffffffffffff1661487d61228a565b73ffffffffffffffffffffffffffffffffffffffff16146148d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016148ca906184cf565b60405180910390fd5b565b5f600d5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff169050919050565b5f8061493233613bfe565b915091505f60135f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015490504381106149bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016149b390617d02565b60405180910390fd5b600260135f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f6101000a81548160ff02191690836005811115614a1f57614a1e61721b565b5b02179055508360135f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206004018190555050505050565b5f614a7982612156565b11614ab9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614ab090618537565b60405180910390fd5b601960405180604001604052808473ffffffffffffffffffffffffffffffffffffffff16815260200183815250908060018154018082558091505060019003905f5260205f2090600202015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001019081614b6591906186e9565b5050505050565b5f60115f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f82118015614bf857508082105b15614c01578190505b8060115f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254614c8a9190617ffc565b925050819055508060075f828254614ca29190617ffc565b92505081905550614cb3848461244d565b8015614d3857505f60115f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054145b15614d4857614d4784846167cf565b5b8373ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f19350505050158015614d8b573d5f803e3d5ffd5b508373ffffffffffffffffffffffffffffffffffffffff167fce2e26dc37234e172764d0a8a2c04d2e6fb46c02f3ca6ca5135d6f9dcd2b67a982604051614dd291906173a2565b60405180910390a250505050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ede391fb336040518263ffffffff1660e01b8152600401614e3a91906178ad565b602060405180830381865afa158015614e55573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614e799190617bff565b15614fdc57600580811115614e9157614e9061721b565b5b60135f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff166005811115614ef057614eef61721b565b5b141580614efd57505f3414155b15614fdb5760025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663489fe625336040518263ffffffff1660e01b8152600401614f5c91906178ad565b602060405180830381865afa158015614f77573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614f9b9190617bff565b614fda576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614fd190617c9a565b60405180910390fd5b5b5b5f60135f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015490505f429050808210615066576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161505d90617d02565b60405180910390fd5b61506f336148d5565b8015615080575061507f8361586f565b5b156151e7578273ffffffffffffffffffffffffffffffffffffffff16600b5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161480156151a357503373ffffffffffffffffffffffffffffffffffffffff16600c5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b6151e2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016151d9906182c0565b60405180910390fd5b6152f7565b6151f0336148d5565b1580156152035750615201836148d5565b155b801561521557506152138361586f565b155b801561522757506152253361586f565b155b801561526f57505f60055f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054145b80156152b757505f60055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054145b6152f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016152ed9061839a565b60405180910390fd5b5b6001600d5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055506001600e5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555082600b5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555033600c5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503460075f8282546154a89190617d4d565b925050819055503460055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546154fb9190617d4d565b9250508190555061550b83613cf6565b1561551a5761551983613e00565b5b3373ffffffffffffffffffffffffffffffffffffffff167f9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d3460405161556091906173a2565b60405180910390a2505050565b600360135f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f6101000a81548160ff021916908360058111156155d0576155cf61721b565b5b02179055506103e86fffffffffffffffffffffffffffffffff16436155f59190617d4d565b60135f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003018190555050565b61564581616776565b615684576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161567b90618828565b60405180910390fd5b3460075f8282546156959190617d4d565b925050819055503460115f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546157259190617d4d565b925050819055506157363382616d4e565b15615746576157453382616e03565b5b8073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f1ed8ad98d928651a8bc3999999b718383931f4595fcd2e1efd2de972fa8cdaea346040516157a391906173a2565b60405180910390a350565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f600e5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff169050919050565b6008546001805490501161590a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615901906188b6565b60405180910390fd5b60018054905060065f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541061598e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016159859061891e565b60405180910390fd5b5f60065f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f600180805490506159e19190617ffc565b9050808214615ac7575f600182815481106159ff576159fe617d80565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508060018481548110615a3e57615a3d617d80565b5b905f5260205f20015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508260065f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f60045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505f60065f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055506001805480615b6f57615b6e61893c565b5b600190038181905f5260205f20015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055905560016019805490501015615bb6575050615f1d565b5f806001601980549050615bca9190617ffc565b90505f5b601980549050811015615c66578573ffffffffffffffffffffffffffffffffffffffff1660198281548110615c0657615c05617d80565b5b905f5260205f2090600202015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603615c53578092505b8080615c5e90618969565b915050615bce565b50808214615ebb575f60198381548110615c8357615c82617d80565b5b905f5260205f2090600202016040518060400160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182018054615cfe90617fcc565b80601f0160208091040260200160405190810160405280929190818152602001828054615d2a90617fcc565b8015615d755780601f10615d4c57610100808354040283529160200191615d75565b820191905f5260205f20905b815481529060010190602001808311615d5857829003601f168201915b505050505081525050905060198281548110615d9457615d93617d80565b5b905f5260205f20906002020160198481548110615db457615db3617d80565b5b905f5260205f2090600202015f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018201816001019081615e3791906189c5565b509050508060198381548110615e5057615e4f617d80565b5b905f5260205f2090600202015f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001019081615eb591906186e9565b50905050505b6019805480615ecd57615ecc61893c565b5b600190038181905f5260205f2090600202015f8082015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600182015f615f1491906170cb565b50509055505050505b50565b5f60125f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208054905090505f5b818167ffffffffffffffff1610156160195761600360125f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208267ffffffffffffffff1681548110615fd457615fd3617d80565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16845f614b6c565b818061600e90618aaa565b9250505f8203615f66575b505050565b600467ffffffffffffffff168167ffffffffffffffff161180156160575750609667ffffffffffffffff168167ffffffffffffffff1611155b1561612957600460135f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f6101000a81548160ff021916908360058111156160bf576160be61721b565b5b02179055506107d06fffffffffffffffffffffffffffffffff16436160e49190617d4d565b60135f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20600301819055505b5f61613383611a35565b90505f8160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205461617f9190617d4d565b90505f60055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f606460196fffffffffffffffffffffffffffffffff16846161e49190618ad1565b6161ee9190618b3f565b9050691fc3842bd1f071c000006fffffffffffffffffffffffffffffffff1681101561623457691fc3842bd1f071c000006fffffffffffffffffffffffffffffffff1690505b828111156162b657600260135f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f6101000a81548160ff0219169083600581111561629f5761629e61721b565b5b02179055506162ad86615f20565b50505050616772565b5f606484836162c59190618ad1565b6162cf9190618b3f565b90505f606482856162e09190618ad1565b6162ea9190618b3f565b90505f81856162f99190617ffc565b90508060055f8b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508160075f82825461634e9190617ffc565b925050819055505f73ffffffffffffffffffffffffffffffffffffffff166108fc8390811502906040515f60405180830381858888f19350505050158015616398573d5f803e3d5ffd5b505f60125f8b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2080548060200260200160405190810160405280929190818152602001828054801561645457602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161640b575b505050505090505f60125f8c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208054905011156166d5575f5b60125f8c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805490508167ffffffffffffffff1610156166d3575f60115f848467ffffffffffffffff168151811061651857616517617d80565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f606487836165a89190618ad1565b6165b29190618b3f565b905080826165c09190617ffc565b60115f868667ffffffffffffffff16815181106165e0576165df617d80565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8f73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508060075f8282546166739190617ffc565b925050819055505f73ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f193505050501580156166bd573d5f803e3d5ffd5b50505080806166cb90617dad565b9150506164a5565b505b8973ffffffffffffffffffffffffffffffffffffffff167f549f9f77a981f051b72b0b5e384ae6a1c3155b1789216ddf6202aaeaf96a6a394360135f8e73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2060030154888b6040516167619493929190618b6f565b60405180910390a250505050505050505b5050565b5f60045f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff169050919050565b5f33905090565b600f5460125f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208054905011616853576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161684a90618c22565b60405180910390fd5b60125f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208054905060115f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410616950576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161694790618cb0565b60405180910390fd5b5f60115f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f600160125f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2080549050616a1d9190617ffc565b9050808214616bb7575f60125f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208281548110616a7657616a75617d80565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508060125f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208481548110616af057616aef617d80565b5b905f5260205f20015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508260115f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f60115f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505f60115f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208190555060125f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805480616d1657616d1561893c565b5b600190038181905f5260205f20015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055905550505050565b5f616d59838361244d565b158015616dfb575068056bc75e2d631000006fffffffffffffffffffffffffffffffff1660115f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410155b905092915050565b60105460125f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208054905010616e87576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401616e7e90618d3e565b60405180910390fd5b600160115f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555060125f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208054905060115f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208190555060125f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2082908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b6040518060a001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f60058111156170b9576170b861721b565b5b81526020015f81526020015f81525090565b5080546170d790617fcc565b5f825580601f106170e85750617105565b601f0160209004905f5260205f20908101906171049190617108565b5b50565b5b8082111561711f575f815f905550600101617109565b5090565b5f82825260208201905092915050565b7f4f6e6c7920454f412063616e2063616c6c2066756e6374696f6e0000000000005f82015250565b5f617167601a83617123565b915061717282617133565b602082019050919050565b5f6020820190508181035f8301526171948161715b565b9050919050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6171ed826171c4565b9050919050565b6171fd816171e3565b82525050565b5f819050919050565b61721581617203565b82525050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b600681106172595761725861721b565b5b50565b5f81905061726982617248565b919050565b5f6172788261725c565b9050919050565b6172888161726e565b82525050565b60a082015f8201516172a25f8501826171f4565b5060208201516172b5602085018261720c565b5060408201516172c8604085018261727f565b5060608201516172db606085018261720c565b5060808201516172ee608085018261720c565b50505050565b5f6172ff838361728e565b60a08301905092915050565b5f602082019050919050565b5f6173218261719b565b61732b81856171a5565b9350617336836171b5565b805f5b8381101561736657815161734d88826172f4565b97506173588361730b565b925050600181019050617339565b5085935050505092915050565b5f6020820190508181035f83015261738b8184617317565b905092915050565b61739c81617203565b82525050565b5f6020820190506173b55f830184617393565b92915050565b5f604051905090565b5f80fd5b5f80fd5b6173d5816171e3565b81146173df575f80fd5b50565b5f813590506173f0816173cc565b92915050565b5f6020828403121561740b5761740a6173c4565b5b5f617418848285016173e2565b91505092915050565b61742a81617203565b8114617434575f80fd5b50565b5f8135905061744581617421565b92915050565b5f602082840312156174605761745f6173c4565b5b5f61746d84828501617437565b91505092915050565b5f6fffffffffffffffffffffffffffffffff82169050919050565b61749a81617476565b82525050565b5f6020820190506174b35f830184617491565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b617507826174c1565b810181811067ffffffffffffffff82111715617526576175256174d1565b5b80604052505050565b5f6175386173bb565b905061754482826174fe565b919050565b5f67ffffffffffffffff821115617563576175626174d1565b5b61756c826174c1565b9050602081019050919050565b828183375f83830152505050565b5f61759961759484617549565b61752f565b9050828152602081018484840111156175b5576175b46174bd565b5b6175c0848285617579565b509392505050565b5f82601f8301126175dc576175db6174b9565b5b81356175ec848260208601617587565b91505092915050565b5f6020828403121561760a576176096173c4565b5b5f82013567ffffffffffffffff811115617627576176266173c8565b5b617633848285016175c8565b91505092915050565b5f8060408385031215617652576176516173c4565b5b5f61765f858286016173e2565b925050602061767085828601617437565b9150509250929050565b5f80604083850312156176905761768f6173c4565b5b5f61769d858286016173e2565b925050602083013567ffffffffffffffff8111156176be576176bd6173c8565b5b6176ca858286016175c8565b9150509250929050565b5f60ff82169050919050565b6176e9816176d4565b81146176f3575f80fd5b50565b5f81359050617704816176e0565b92915050565b5f805f60608486031215617721576177206173c4565b5b5f61772e86828701617437565b935050602061773f868287016173e2565b9250506040617750868287016176f6565b9150509250925092565b5f80604083850312156177705761776f6173c4565b5b5f61777d858286016173e2565b925050602061778e858286016173e2565b9150509250929050565b5f8115159050919050565b6177ac81617798565b82525050565b5f6020820190506177c55f8301846177a3565b92915050565b5f67ffffffffffffffff82169050919050565b6177e7816177cb565b82525050565b5f6020820190506178005f8301846177de565b92915050565b61780f816171e3565b82525050565b5f81519050919050565b5f5b8381101561783c578082015181840152602081019050617821565b5f8484015250505050565b5f61785182617815565b61785b8185617123565b935061786b81856020860161781f565b617874816174c1565b840191505092915050565b5f6040820190506178925f830185617806565b81810360208301526178a48184617847565b90509392505050565b5f6020820190506178c05f830184617806565b92915050565b6178cf816177cb565b81146178d9575f80fd5b50565b5f813590506178ea816178c6565b92915050565b5f8060408385031215617906576179056173c4565b5b5f617913858286016173e2565b9250506020617924858286016178dc565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f61796283836171f4565b60208301905092915050565b5f602082019050919050565b5f6179848261792e565b61798e8185617938565b935061799983617948565b805f5b838110156179c95781516179b08882617957565b97506179bb8361796e565b92505060018101905061799c565b5085935050505092915050565b5f6020820190508181035f8301526179ee818461797a565b905092915050565b5f819050919050565b5f617a19617a14617a0f846171c4565b6179f6565b6171c4565b9050919050565b5f617a2a826179ff565b9050919050565b5f617a3b82617a20565b9050919050565b617a4b81617a31565b82525050565b5f602082019050617a645f830184617a42565b92915050565b5f805f60608486031215617a8157617a806173c4565b5b5f617a8e868287016173e2565b9350506020617a9f868287016173e2565b925050604084013567ffffffffffffffff811115617ac057617abf6173c8565b5b617acc868287016175c8565b9150509250925092565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b617b08816176d4565b82525050565b5f617b198383617aff565b60208301905092915050565b5f602082019050919050565b5f617b3b82617ad6565b617b458185617ae0565b9350617b5083617af0565b805f5b83811015617b80578151617b678882617b0e565b9750617b7283617b25565b925050600181019050617b53565b5085935050505092915050565b5f6020820190508181035f830152617ba58184617b31565b905092915050565b617bb6816176d4565b82525050565b5f602082019050617bcf5f830184617bad565b92915050565b617bde81617798565b8114617be8575f80fd5b50565b5f81519050617bf981617bd5565b92915050565b5f60208284031215617c1457617c136173c4565b5b5f617c2184828501617beb565b91505092915050565b7f706c656173652070617920746865206772616e74656420616d6f756e742062655f8201527f666f7265207374616b696e6720616761696e0000000000000000000000000000602082015250565b5f617c84603283617123565b9150617c8f82617c2a565b604082019050919050565b5f6020820190508181035f830152617cb181617c78565b9050919050565b7f7374616b65722069732062616e6e6564206f722073757370656e6465640000005f82015250565b5f617cec601d83617123565b9150617cf782617cb8565b602082019050919050565b5f6020820190508181035f830152617d1981617ce0565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f617d5782617203565b9150617d6283617203565b9250828201905080821115617d7a57617d79617d20565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f617db7826177cb565b915067ffffffffffffffff8203617dd157617dd0617d20565b5b600182019050919050565b7f4f6e6c7920626c6f636b2070726f706f7365722063616e2063616c6c2066756e5f8201527f6374696f6e000000000000000000000000000000000000000000000000000000602082015250565b5f617e36602583617123565b9150617e4182617ddc565b604082019050919050565b5f6020820190508181035f830152617e6381617e2a565b9050919050565b7f4f6e6c79207374616b65722063616e2063616c6c2066756e6374696f6e0000005f82015250565b5f617e9e601d83617123565b9150617ea982617e6a565b602082019050919050565b5f6020820190508181035f830152617ecb81617e92565b9050919050565b7f6f6e6c792064656c656761746f722063616e2063616c6c2066756e6374696f6e5f82015250565b5f617f06602083617123565b9150617f1182617ed2565b602082019050919050565b5f6020820190508181035f830152617f3381617efa565b9050919050565b5f617f54617f4f617f4a846176d4565b6179f6565b6177cb565b9050919050565b617f6481617f3a565b82525050565b5f606082019050617f7d5f830186617393565b617f8a6020830185617393565b617f976040830184617f5b565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680617fe357607f821691505b602082108103617ff657617ff5617f9f565b5b50919050565b5f61800682617203565b915061801183617203565b925082820390508181111561802957618028617d20565b5b92915050565b7f6e6f7420656e6f756768206d6f6e6579000000000000000000000000000000005f82015250565b5f618063601083617123565b915061806e8261802f565b602082019050919050565b5f6020820190508181035f83015261809081618057565b9050919050565b5f6040820190506180aa5f830185617806565b6180b76020830184617393565b9392505050565b7f74696d656f7574000000000000000000000000000000000000000000000000005f82015250565b5f6180f2600783617123565b91506180fd826180be565b602082019050919050565b5f60408201905061811b5f830184617393565b818103602083015261812c816180e6565b905092915050565b7f416d6f756e74206d7573742062652067726561746572207468616e2076616c695f8201527f6461746f72207468726573686f6c640000000000000000000000000000000000602082015250565b5f61818e602f83617123565b915061819982618134565b604082019050919050565b5f6020820190508181035f8301526181bb81618182565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f61821c602683617123565b9150618227826181c2565b604082019050919050565b5f6020820190508181035f83015261824981618210565b9050919050565b7f7468652073656e646572206973206e6f7420612076616c6964207374616b65725f8201527f206f662074686973207369676e65720000000000000000000000000000000000602082015250565b5f6182aa602f83617123565b91506182b582618250565b604082019050919050565b5f6020820190508181035f8301526182d78161829e565b9050919050565b7f626f74682061646472657373657320617265206e6f742076616c696420746f205f8201527f63616c6c2066756e6374696f6e2e20496620796f75207374696c6c2077616e7460208201527f20746f2063616c6c2e204a7573742063616c6c696e6720756e7374616b65206660408201527f6f7220626f7468206163636f756e7473206d616e75616c6c7900000000000000606082015250565b5f618384607983617123565b915061838f826182de565b608082019050919050565b5f6020820190508181035f8301526183b181618378565b9050919050565b7f56616c696461746f72207365742068617320726561636865642066756c6c20635f8201527f6170616369747900000000000000000000000000000000000000000000000000602082015250565b5f618412602783617123565b915061841d826183b8565b604082019050919050565b5f6020820190508181035f83015261843f81618406565b9050919050565b5f8151905061845481617421565b92915050565b5f6020828403121561846f5761846e6173c4565b5b5f61847c84828501618446565b91505092915050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f6184b9602083617123565b91506184c482618485565b602082019050919050565b5f6020820190508181035f8301526184e6816184ad565b9050919050565b7f506c6561736520736574207468652063726577206e616d6500000000000000005f82015250565b5f618521601883617123565b915061852c826184ed565b602082019050919050565b5f6020820190508181035f83015261854e81618515565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026185b17fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82618576565b6185bb8683618576565b95508019841693508086168417925050509392505050565b5f6185ed6185e86185e384617203565b6179f6565b617203565b9050919050565b5f819050919050565b618606836185d3565b61861a618612826185f4565b848454618582565b825550505050565b5f90565b61862e618622565b6186398184846185fd565b505050565b5b8181101561865c576186515f82618626565b60018101905061863f565b5050565b601f8211156186a15761867281618555565b61867b84618567565b8101602085101561868a578190505b61869e61869685618567565b83018261863e565b50505b505050565b5f82821c905092915050565b5f6186c15f19846008026186a6565b1980831691505092915050565b5f6186d983836186b2565b9150826002028217905092915050565b6186f282617815565b67ffffffffffffffff81111561870b5761870a6174d1565b5b6187158254617fcc565b618720828285618660565b5f60209050601f831160018114618751575f841561873f578287015190505b61874985826186ce565b8655506187b0565b601f19841661875f86618555565b5f5b8281101561878657848901518255600182019150602085019450602081019050618761565b868310156187a3578489015161879f601f8916826186b2565b8355505b6001600288020188555050505b505050505050565b7f63616e6e6f742064656c656761746520746f6b656e7320746f20736f6d656f6e5f8201527f652077686f206973206e6f742076616c696461746f7200000000000000000000602082015250565b5f618812603683617123565b915061881d826187b8565b604082019050919050565b5f6020820190508181035f83015261883f81618806565b9050919050565b7f56616c696461746f72732063616e2774206265206c657373207468616e2074685f8201527f65206d696e696d756d2072657175697265642076616c696461746f72206e756d602082015250565b5f6188a0604083617123565b91506188ab82618846565b604082019050919050565b5f6020820190508181035f8301526188cd81618894565b9050919050565b7f696e646578206f7574206f662072616e676500000000000000000000000000005f82015250565b5f618908601283617123565b9150618913826188d4565b602082019050919050565b5f6020820190508181035f830152618935816188fc565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b5f61897382617203565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036189a5576189a4617d20565b5b600182019050919050565b5f815490506189be81617fcc565b9050919050565b8181036189d3575050618aa8565b6189dc826189b0565b67ffffffffffffffff8111156189f5576189f46174d1565b5b6189ff8254617fcc565b618a0a828285618660565b5f601f831160018114618a37575f8415618a25578287015490505b618a2f85826186ce565b865550618aa1565b601f198416618a4587618555565b9650618a5086618555565b5f5b82811015618a7757848901548255600182019150600185019450602081019050618a52565b86831015618a945784890154618a90601f8916826186b2565b8355505b6001600288020188555050505b5050505050505b565b5f618ab482617203565b91505f8203618ac657618ac5617d20565b5b600182039050919050565b5f618adb82617203565b9150618ae683617203565b9250828202618af481617203565b91508282048414831517618b0b57618b0a617d20565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f618b4982617203565b9150618b5483617203565b925082618b6457618b63618b12565b5b828204905092915050565b5f608082019050618b825f830187617393565b618b8f6020830186617393565b618b9c6040830185617393565b618ba96060830184617393565b95945050505050565b7f64656c656761746f72732073697a652063616e2774206265206c6573732074685f8201527f616e20746865206d696e696d756d20726571756972656d656e74000000000000602082015250565b5f618c0c603a83617123565b9150618c1782618bb2565b604082019050919050565b5f6020820190508181035f830152618c3981618c00565b9050919050565b7f696e646578206f7574206f662072616e676520696e207468652064656c6567615f8201527f746f7273206c697374206f662076616c696461746f7200000000000000000000602082015250565b5f618c9a603683617123565b9150618ca582618c40565b604082019050919050565b5f6020820190508181035f830152618cc781618c8e565b9050919050565b7f64656c656761746f7220736574206861732072656163686564206974732066755f8201527f6c6c206361706163697479000000000000000000000000000000000000000000602082015250565b5f618d28602b83617123565b9150618d3382618cce565b604082019050919050565b5f6020820190508181035f830152618d5581618d1c565b905091905056fea264697066735822122034fb316be59888abc9500bbfc6930164fff10764316f0be6272b51c1bff1ed3564736f6c637828302e382e32322d646576656c6f702e323032332e392e33302b636f6d6d69742e66653166396336340059",
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
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
