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

// StakingValidatorStakeInfo is an auto generated low-level Go binding around an user-defined struct.
type StakingValidatorStakeInfo struct {
	Addr   common.Address
	Amount *big.Int
}

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minNumValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNumValidators\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accout\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"BLSPublicKeyRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VALIDATOR_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToBLSPublicKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToIsValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToStakedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToValidatorIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_maximumNumValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_minimumNumValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_stakedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_validators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"accountStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorsStakeInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structStaking.ValidatorStakeInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maximumNumValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumNumValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"blsPubKey\",\"type\":\"bytes\"}],\"name\":\"registerBLSPublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorBLSPublicKeys\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200244e3803806200244e8339818101604052810190620000379190620000d3565b808211156200007d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200007490620001a1565b60405180910390fd5b81600581905550806006819055505050620001c3565b600080fd5b6000819050919050565b620000ad8162000098565b8114620000b957600080fd5b50565b600081519050620000cd81620000a2565b92915050565b60008060408385031215620000ed57620000ec62000093565b5b6000620000fd85828601620000bc565b92505060206200011085828601620000bc565b9150509250929050565b600082825260208201905092915050565b7f4d696e2076616c696461746f7273206e756d2063616e206e6f7420626520677260008201527f6561746572207468616e206d6178206e756d206f662076616c696461746f7273602082015250565b6000620001896040836200011a565b915062000196826200012b565b604082019050919050565b60006020820190508181036000830152620001bc816200017a565b9050919050565b61227b80620001d36000396000f3fe6080604052600436106101235760003560e01c80637a6eea37116100a0578063d94c111b11610064578063d94c111b14610440578063e387a7ed14610469578063e804fbf614610494578063f90ecacc146104bf578063facd743b146104fc57610191565b80637a6eea37146103575780637dceceb814610382578063af6da36e146103bf578063c795c077146103ea578063ca1e78191461041557610191565b8063373d6132116100e7578063373d61321461028f5780633a4b66f1146102ba5780633c561f04146102c457806351a9ab32146102ef578063714ff4251461032c57610191565b806302b751991461019657806303306f0a146101d3578063065ae171146101fe5780632367f6b51461023b5780632def66201461027857610191565b36610191576101473373ffffffffffffffffffffffffffffffffffffffff16610539565b15610187576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161017e906114ef565b60405180910390fd5b61018f61055c565b005b600080fd5b3480156101a257600080fd5b506101bd60048036038101906101b89190611581565b610633565b6040516101ca91906115c7565b60405180910390f35b3480156101df57600080fd5b506101e861064b565b6040516101f591906116de565b60405180910390f35b34801561020a57600080fd5b5061022560048036038101906102209190611581565b6107dd565b604051610232919061171b565b60405180910390f35b34801561024757600080fd5b50610262600480360381019061025d9190611581565b6107fd565b60405161026f91906115c7565b60405180910390f35b34801561028457600080fd5b5061028d610846565b005b34801561029b57600080fd5b506102a4610931565b6040516102b191906115c7565b60405180910390f35b6102c261093b565b005b3480156102d057600080fd5b506102d96109a4565b6040516102e69190611888565b60405180910390f35b3480156102fb57600080fd5b5061031660048036038101906103119190611581565b610b4a565b60405161032391906118f4565b60405180910390f35b34801561033857600080fd5b50610341610bea565b60405161034e91906115c7565b60405180910390f35b34801561036357600080fd5b5061036c610bf4565b6040516103799190611941565b60405180910390f35b34801561038e57600080fd5b506103a960048036038101906103a49190611581565b610c00565b6040516103b691906115c7565b60405180910390f35b3480156103cb57600080fd5b506103d4610c18565b6040516103e191906115c7565b60405180910390f35b3480156103f657600080fd5b506103ff610c1e565b60405161040c91906115c7565b60405180910390f35b34801561042157600080fd5b5061042a610c24565b6040516104379190611a0b565b60405180910390f35b34801561044c57600080fd5b5061046760048036038101906104629190611b62565b610cb2565b005b34801561047557600080fd5b5061047e610d50565b60405161048b91906115c7565b60405180910390f35b3480156104a057600080fd5b506104a9610d56565b6040516104b691906115c7565b60405180910390f35b3480156104cb57600080fd5b506104e660048036038101906104e19190611bd7565b610d60565b6040516104f39190611c13565b60405180910390f35b34801561050857600080fd5b50610523600480360381019061051e9190611581565b610d9f565b604051610530919061171b565b60405180910390f35b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b346004600082825461056e9190611c5d565b9250508190555034600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546105c49190611c5d565b925050819055506105d433610df5565b156105e3576105e233610e6d565b5b3373ffffffffffffffffffffffffffffffffffffffff167f9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d3460405161062991906115c7565b60405180910390a2565b60036020528060005260406000206000915090505481565b60606000808054905067ffffffffffffffff81111561066d5761066c611a37565b5b6040519080825280602002602001820160405280156106a657816020015b610693611462565b81526020019060019003908161068b5790505b50905060005b6000805490508110156107d55760006040518060400160405280600084815481106106da576106d9611c91565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002600080868154811061073757610736611c91565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548152509050808383815181106107b6576107b5611c91565b5b60200260200101819052505080806107cd90611cc0565b9150506106ac565b508091505090565b60016020528060005260406000206000915054906101000a900460ff1681565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6108653373ffffffffffffffffffffffffffffffffffffffff16610539565b156108a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161089c906114ef565b60405180910390fd5b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205411610927576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161091e90611d54565b60405180910390fd5b61092f610fbc565b565b6000600454905090565b61095a3373ffffffffffffffffffffffffffffffffffffffff16610539565b1561099a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610991906114ef565b60405180910390fd5b6109a261055c565b565b60606000808054905067ffffffffffffffff8111156109c6576109c5611a37565b5b6040519080825280602002602001820160405280156109f957816020015b60608152602001906001900390816109e45790505b50905060005b600080549050811015610b425760076000808381548110610a2357610a22611c91565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054610a9390611da3565b80601f0160208091040260200160405190810160405280929190818152602001828054610abf90611da3565b8015610b0c5780601f10610ae157610100808354040283529160200191610b0c565b820191906000526020600020905b815481529060010190602001808311610aef57829003601f168201915b5050505050828281518110610b2457610b23611c91565b5b60200260200101819052508080610b3a90611cc0565b9150506109ff565b508091505090565b60076020528060005260406000206000915090508054610b6990611da3565b80601f0160208091040260200160405190810160405280929190818152602001828054610b9590611da3565b8015610be25780601f10610bb757610100808354040283529160200191610be2565b820191906000526020600020905b815481529060010190602001808311610bc557829003601f168201915b505050505081565b6000600554905090565b670de0b6b3a764000081565b60026020528060005260406000206000915090505481565b60065481565b60055481565b60606000805480602002602001604051908101604052809291908181526020018280548015610ca857602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610c5e575b5050505050905090565b80600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209081610cfe9190611f80565b503373ffffffffffffffffffffffffffffffffffffffff167f472da4d064218fa97032725fbcff922201fa643fed0765b5ffe0ceef63d7b3dc82604051610d4591906118f4565b60405180910390a250565b60045481565b6000600654905090565b60008181548110610d7057600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b6000610e008261110e565b158015610e665750670de0b6b3a76400006fffffffffffffffffffffffffffffffff16600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410155b9050919050565b60065460008054905010610eb6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ead906120c4565b60405180910390fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600080549050600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506000819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550806004600082825461105791906120e4565b925050819055506110673361110e565b156110765761107533611164565b5b3373ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156110bc573d6000803e3d6000fd5b503373ffffffffffffffffffffffffffffffffffffffff167f0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f758260405161110391906115c7565b60405180910390a250565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b600554600080549050116111ad576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111a49061218a565b60405180910390fd5b600080549050600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410611233576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161122a906121f6565b60405180910390fd5b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000600160008054905061128b91906120e4565b90508082146113795760008082815481106112a9576112a8611c91565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905080600084815481106112eb576112ea611c91565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505b6000600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506000600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600080548061142857611427612216565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690559055505050565b6040518060400160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b600082825260208201905092915050565b7f4f6e6c7920454f412063616e2063616c6c2066756e6374696f6e000000000000600082015250565b60006114d9601a83611492565b91506114e4826114a3565b602082019050919050565b60006020820190508181036000830152611508816114cc565b9050919050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061154e82611523565b9050919050565b61155e81611543565b811461156957600080fd5b50565b60008135905061157b81611555565b92915050565b60006020828403121561159757611596611519565b5b60006115a58482850161156c565b91505092915050565b6000819050919050565b6115c1816115ae565b82525050565b60006020820190506115dc60008301846115b8565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61161781611543565b82525050565b611626816115ae565b82525050565b604082016000820151611642600085018261160e565b506020820151611655602085018261161d565b50505050565b6000611667838361162c565b60408301905092915050565b6000602082019050919050565b600061168b826115e2565b61169581856115ed565b93506116a0836115fe565b8060005b838110156116d15781516116b8888261165b565b97506116c383611673565b9250506001810190506116a4565b5085935050505092915050565b600060208201905081810360008301526116f88184611680565b905092915050565b60008115159050919050565b61171581611700565b82525050565b6000602082019050611730600083018461170c565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561179c578082015181840152602081019050611781565b60008484015250505050565b6000601f19601f8301169050919050565b60006117c482611762565b6117ce818561176d565b93506117de81856020860161177e565b6117e7816117a8565b840191505092915050565b60006117fe83836117b9565b905092915050565b6000602082019050919050565b600061181e82611736565b6118288185611741565b93508360208202850161183a85611752565b8060005b85811015611876578484038952815161185785826117f2565b945061186283611806565b925060208a0199505060018101905061183e565b50829750879550505050505092915050565b600060208201905081810360008301526118a28184611813565b905092915050565b600082825260208201905092915050565b60006118c682611762565b6118d081856118aa565b93506118e081856020860161177e565b6118e9816117a8565b840191505092915050565b6000602082019050818103600083015261190e81846118bb565b905092915050565b60006fffffffffffffffffffffffffffffffff82169050919050565b61193b81611916565b82525050565b60006020820190506119566000830184611932565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000611994838361160e565b60208301905092915050565b6000602082019050919050565b60006119b88261195c565b6119c28185611967565b93506119cd83611978565b8060005b838110156119fe5781516119e58882611988565b97506119f0836119a0565b9250506001810190506119d1565b5085935050505092915050565b60006020820190508181036000830152611a2581846119ad565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b611a6f826117a8565b810181811067ffffffffffffffff82111715611a8e57611a8d611a37565b5b80604052505050565b6000611aa161150f565b9050611aad8282611a66565b919050565b600067ffffffffffffffff821115611acd57611acc611a37565b5b611ad6826117a8565b9050602081019050919050565b82818337600083830152505050565b6000611b05611b0084611ab2565b611a97565b905082815260208101848484011115611b2157611b20611a32565b5b611b2c848285611ae3565b509392505050565b600082601f830112611b4957611b48611a2d565b5b8135611b59848260208601611af2565b91505092915050565b600060208284031215611b7857611b77611519565b5b600082013567ffffffffffffffff811115611b9657611b9561151e565b5b611ba284828501611b34565b91505092915050565b611bb4816115ae565b8114611bbf57600080fd5b50565b600081359050611bd181611bab565b92915050565b600060208284031215611bed57611bec611519565b5b6000611bfb84828501611bc2565b91505092915050565b611c0d81611543565b82525050565b6000602082019050611c286000830184611c04565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611c68826115ae565b9150611c73836115ae565b9250828201905080821115611c8b57611c8a611c2e565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000611ccb826115ae565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611cfd57611cfc611c2e565b5b600182019050919050565b7f4f6e6c79207374616b65722063616e2063616c6c2066756e6374696f6e000000600082015250565b6000611d3e601d83611492565b9150611d4982611d08565b602082019050919050565b60006020820190508181036000830152611d6d81611d31565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680611dbb57607f821691505b602082108103611dce57611dcd611d74565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302611e367fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611df9565b611e408683611df9565b95508019841693508086168417925050509392505050565b6000819050919050565b6000611e7d611e78611e73846115ae565b611e58565b6115ae565b9050919050565b6000819050919050565b611e9783611e62565b611eab611ea382611e84565b848454611e06565b825550505050565b600090565b611ec0611eb3565b611ecb818484611e8e565b505050565b5b81811015611eef57611ee4600082611eb8565b600181019050611ed1565b5050565b601f821115611f3457611f0581611dd4565b611f0e84611de9565b81016020851015611f1d578190505b611f31611f2985611de9565b830182611ed0565b50505b505050565b600082821c905092915050565b6000611f5760001984600802611f39565b1980831691505092915050565b6000611f708383611f46565b9150826002028217905092915050565b611f8982611762565b67ffffffffffffffff811115611fa257611fa1611a37565b5b611fac8254611da3565b611fb7828285611ef3565b600060209050601f831160018114611fea5760008415611fd8578287015190505b611fe28582611f64565b86555061204a565b601f198416611ff886611dd4565b60005b8281101561202057848901518255600182019150602085019450602081019050611ffb565b8683101561203d5784890151612039601f891682611f46565b8355505b6001600288020188555050505b505050505050565b7f56616c696461746f72207365742068617320726561636865642066756c6c206360008201527f6170616369747900000000000000000000000000000000000000000000000000602082015250565b60006120ae602783611492565b91506120b982612052565b604082019050919050565b600060208201905081810360008301526120dd816120a1565b9050919050565b60006120ef826115ae565b91506120fa836115ae565b925082820390508181111561211257612111611c2e565b5b92915050565b7f56616c696461746f72732063616e2774206265206c657373207468616e20746860008201527f65206d696e696d756d2072657175697265642076616c696461746f72206e756d602082015250565b6000612174604083611492565b915061217f82612118565b604082019050919050565b600060208201905081810360008301526121a381612167565b9050919050565b7f696e646578206f7574206f662072616e67650000000000000000000000000000600082015250565b60006121e0601283611492565b91506121eb826121aa565b602082019050919050565b6000602082019050818103600083015261220f816121d3565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea2646970667358221220e0cb4ff6026abcd095ea8d9111d8c2b9c95347803db6a474155c365aab7e37ce64736f6c63430008110033",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// StakingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingMetaData.Bin instead.
var StakingBin = StakingMetaData.Bin

// DeployStaking deploys a new Ethereum contract, binding an instance of Staking to it.
func DeployStaking(auth *bind.TransactOpts, backend bind.ContractBackend, minNumValidators *big.Int, maxNumValidators *big.Int) (common.Address, *types.Transaction, *Staking, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingBin), backend, minNumValidators, maxNumValidators)
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

// AddressToBLSPublicKey is a free data retrieval call binding the contract method 0x51a9ab32.
//
// Solidity: function _addressToBLSPublicKey(address ) view returns(bytes)
func (_Staking *StakingCaller) AddressToBLSPublicKey(opts *bind.CallOpts, arg0 common.Address) ([]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToBLSPublicKey", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// AddressToBLSPublicKey is a free data retrieval call binding the contract method 0x51a9ab32.
//
// Solidity: function _addressToBLSPublicKey(address ) view returns(bytes)
func (_Staking *StakingSession) AddressToBLSPublicKey(arg0 common.Address) ([]byte, error) {
	return _Staking.Contract.AddressToBLSPublicKey(&_Staking.CallOpts, arg0)
}

// AddressToBLSPublicKey is a free data retrieval call binding the contract method 0x51a9ab32.
//
// Solidity: function _addressToBLSPublicKey(address ) view returns(bytes)
func (_Staking *StakingCallerSession) AddressToBLSPublicKey(arg0 common.Address) ([]byte, error) {
	return _Staking.Contract.AddressToBLSPublicKey(&_Staking.CallOpts, arg0)
}

// AddressToIsValidator is a free data retrieval call binding the contract method 0x065ae171.
//
// Solidity: function _addressToIsValidator(address ) view returns(bool)
func (_Staking *StakingCaller) AddressToIsValidator(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToIsValidator", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AddressToIsValidator is a free data retrieval call binding the contract method 0x065ae171.
//
// Solidity: function _addressToIsValidator(address ) view returns(bool)
func (_Staking *StakingSession) AddressToIsValidator(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressToIsValidator(&_Staking.CallOpts, arg0)
}

// AddressToIsValidator is a free data retrieval call binding the contract method 0x065ae171.
//
// Solidity: function _addressToIsValidator(address ) view returns(bool)
func (_Staking *StakingCallerSession) AddressToIsValidator(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressToIsValidator(&_Staking.CallOpts, arg0)
}

// AddressToStakedAmount is a free data retrieval call binding the contract method 0x7dceceb8.
//
// Solidity: function _addressToStakedAmount(address ) view returns(uint256)
func (_Staking *StakingCaller) AddressToStakedAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToStakedAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressToStakedAmount is a free data retrieval call binding the contract method 0x7dceceb8.
//
// Solidity: function _addressToStakedAmount(address ) view returns(uint256)
func (_Staking *StakingSession) AddressToStakedAmount(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToStakedAmount(&_Staking.CallOpts, arg0)
}

// AddressToStakedAmount is a free data retrieval call binding the contract method 0x7dceceb8.
//
// Solidity: function _addressToStakedAmount(address ) view returns(uint256)
func (_Staking *StakingCallerSession) AddressToStakedAmount(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToStakedAmount(&_Staking.CallOpts, arg0)
}

// AddressToValidatorIndex is a free data retrieval call binding the contract method 0x02b75199.
//
// Solidity: function _addressToValidatorIndex(address ) view returns(uint256)
func (_Staking *StakingCaller) AddressToValidatorIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToValidatorIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressToValidatorIndex is a free data retrieval call binding the contract method 0x02b75199.
//
// Solidity: function _addressToValidatorIndex(address ) view returns(uint256)
func (_Staking *StakingSession) AddressToValidatorIndex(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToValidatorIndex(&_Staking.CallOpts, arg0)
}

// AddressToValidatorIndex is a free data retrieval call binding the contract method 0x02b75199.
//
// Solidity: function _addressToValidatorIndex(address ) view returns(uint256)
func (_Staking *StakingCallerSession) AddressToValidatorIndex(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToValidatorIndex(&_Staking.CallOpts, arg0)
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

// GetValidatorsStakeInfo is a free data retrieval call binding the contract method 0x03306f0a.
//
// Solidity: function getValidatorsStakeInfo() view returns((address,uint256)[])
func (_Staking *StakingCaller) GetValidatorsStakeInfo(opts *bind.CallOpts) ([]StakingValidatorStakeInfo, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getValidatorsStakeInfo")

	if err != nil {
		return *new([]StakingValidatorStakeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]StakingValidatorStakeInfo)).(*[]StakingValidatorStakeInfo)

	return out0, err

}

// GetValidatorsStakeInfo is a free data retrieval call binding the contract method 0x03306f0a.
//
// Solidity: function getValidatorsStakeInfo() view returns((address,uint256)[])
func (_Staking *StakingSession) GetValidatorsStakeInfo() ([]StakingValidatorStakeInfo, error) {
	return _Staking.Contract.GetValidatorsStakeInfo(&_Staking.CallOpts)
}

// GetValidatorsStakeInfo is a free data retrieval call binding the contract method 0x03306f0a.
//
// Solidity: function getValidatorsStakeInfo() view returns((address,uint256)[])
func (_Staking *StakingCallerSession) GetValidatorsStakeInfo() ([]StakingValidatorStakeInfo, error) {
	return _Staking.Contract.GetValidatorsStakeInfo(&_Staking.CallOpts)
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

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_Staking *StakingTransactor) Stake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "stake")
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_Staking *StakingSession) Stake() (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_Staking *StakingTransactorSession) Stake() (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts)
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
