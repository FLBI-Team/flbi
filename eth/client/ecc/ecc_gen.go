// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ecc

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

// ECCABI is the input ABI used to generate the binding from.
const ECCABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addKeyProposals\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"unixTime\",\"type\":\"int64\"},{\"internalType\":\"uint8\",\"name\":\"keyType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"key\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"int64\",\"name\":\"voteCount\",\"type\":\"int64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"devices\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"firmwares\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"unixTime\",\"type\":\"int64\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mf\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mfSig\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mf\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"key\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"keyType\",\"type\":\"uint8\"}],\"name\":\"proposeAddKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"key\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"keyType\",\"type\":\"uint8\"}],\"name\":\"proposeRemoveKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"removeKeyProposals\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"unixTime\",\"type\":\"int64\"},{\"internalType\":\"uint8\",\"name\":\"keyType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"key\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"int64\",\"name\":\"voteCount\",\"type\":\"int64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"timestampHash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"timestamps\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"unixTime\",\"type\":\"int64\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"mfSig\",\"type\":\"bytes\"}],\"name\":\"updateFirmware\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"key\",\"type\":\"address\"}],\"name\":\"voteAddKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"key\",\"type\":\"address\"}],\"name\":\"voteRemoveKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ECCFuncSigs maps the 4-byte function signature to its string representation.
var ECCFuncSigs = map[string]string{
	"a8b40b2d": "addKeyProposals(address)",
	"429b62e5": "admins(address)",
	"e7b4cac6": "devices(address)",
	"7d246739": "firmwares(bytes32)",
	"f25e88b0": "mf(address)",
	"15164cf9": "proposeAddKey(address,uint8)",
	"5cace31a": "proposeRemoveKey(address,uint8)",
	"e840f6a3": "removeKeyProposals(address)",
	"0e744b84": "timestampHash(bytes32)",
	"b5872958": "timestamps(bytes32)",
	"d5b98927": "updateFirmware(bytes32,bytes)",
	"465d8310": "voteAddKey(address)",
	"574e1714": "voteRemoveKey(address)",
}

// ECCBin is the compiled bytecode used for deploying new contracts.
var ECCBin = "0x60806040523480156200001157600080fd5b5060408051610190808252613220820190925260609160208201613200803883390190505090507317dc6ca2e1c84ae4107975a48dfd05831b8addff816000815181106200005b57fe5b60200260200101906001600160a01b031690816001600160a01b03168152505073ac5580ad28a3c0e044a52541785bfd34c753d3bf816001815181106200009e57fe5b60200260200101906001600160a01b031690816001600160a01b03168152505073d0a73fe9d44184e9f1264ce2097064212e67ebfe81600281518110620000e157fe5b60200260200101906001600160a01b031690816001600160a01b03168152505073d3cf32a24e4527cf84ef0c8d571f8a77f6481984816003815181106200012457fe5b60200260200101906001600160a01b031690816001600160a01b03168152505073763a9766405878e73fa45f26c95fe6d43913ff44816004815181106200016757fe5b60200260200101906001600160a01b031690816001600160a01b031681525050730e88fa918b9dd6d1a11285864d7350e6d923de4081600581518110620001aa57fe5b60200260200101906001600160a01b031690816001600160a01b03168152505072a80682712dbf93df472233ea027b38cef73f7081600681518110620001ec57fe5b60200260200101906001600160a01b031690816001600160a01b03168152505073b7797b168ae2c33a68e19756543a8aca7c40053f816007815181106200022f57fe5b60200260200101906001600160a01b031690816001600160a01b031681525050730e0a2c1d9484436499cd900d5fc0278e5d274ca7816008815181106200027257fe5b60200260200101906001600160a01b031690816001600160a01b0316815250507362b0165c43f6378e497a3d91100e99c559e17d1b81600981518110620002b557fe5b60200260200101906001600160a01b031690816001600160a01b03168152505073e504799d748393a53fa1009097f296c9bd8d8dd181600a81518110620002f857fe5b60200260200101906001600160a01b031690816001600160a01b03168152505073ee994e752f37fbb06997c8d9e4f78eb5621cd1e081600b815181106200033b57fe5b60200260200101906001600160a01b031690816001600160a01b03168152505073c28cfe6e78e37a034e18bb87d458244eddff816781600c815181106200037e57fe5b60200260200101906001600160a01b031690816001600160a01b031681525050736c902d9a10836968f3721a4c1e635f28bcc27dda81600d81518110620003c157fe5b60200260200101906001600160a01b031690816001600160a01b03168152505073152adcce72d9f0a8664847158138b5516357ff4081600e815181106200040457fe5b60200260200101906001600160a01b031690816001600160a01b031681525050739c8c770954a0e4f4313e884368eb33a48ba7954781600f815181106200044757fe5b60200260200101906001600160a01b031690816001600160a01b03168152505073936a8feb84a5b5403c70678e5da1710afd803558816010815181106200048a57fe5b60200260200101906001600160a01b031690816001600160a01b03168152505073ac6c0dae19dce2d1c8952c159c2a018d68e2236081601181518110620004cd57fe5b60200260200101906001600160a01b031690816001600160a01b03168152505073125936228a489db7e2e9b8aa90b7820e4fe1bfb0816012815181106200051057fe5b60200260200101906001600160a01b031690816001600160a01b03168152505073c6bcca3683bf39b1cf55d75c82ef57f432f7b9e2816013815181106200055357fe5b60200260200101906001600160a01b031690816001600160a01b03168152505073b3372b8312679b4db4252fc16da5f2e02b8624db816014815181106200059657fe5b60200260200101906001600160a01b031690816001600160a01b031681525050730d7f9ef3a873b7e1b89e9fe31503072fc3b5c55781601581518110620005d957fe5b60200260200101906001600160a01b031690816001600160a01b03168152505073e2cf703df8f2fdcaf2a5a54e5e8b88e640f6792a816016815181106200061c57fe5b60200260200101906001600160a01b031690816001600160a01b03168152505073e0b471a7f7a8be9085937c3499488fb155a67007816017815181106200065f57fe5b60200260200101906001600160a01b031690816001600160a01b03168152505073c65b5d83bd61ba3aa50a67a2649bc85d303f411681601881518110620006a257fe5b60200260200101906001600160a01b031690816001600160a01b03168152505073e503721f3e82a505ffe7cc83dddc768142df93a381601981518110620006e557fe5b60200260200101906001600160a01b031690816001600160a01b031681525050735ac36dea5c219ee49037fb82c9e6d6a58ba4b3a881601a815181106200072857fe5b60200260200101906001600160a01b031690816001600160a01b03168152505073ae4a0660813ae865a36e79ffacaa907aed78b4b481601b815181106200076b57fe5b60200260200101906001600160a01b031690816001600160a01b0316815250507315ee24641049e10edab9bf2862fa603d9636abea81601c81518110620007ae57fe5b60200260200101906001600160a01b031690816001600160a01b03168152505073fb0dbfd4dbfa13d048e2c8a356f583e630d562de81601d81518110620007f157fe5b60200260200101906001600160a01b031690816001600160a01b031681525050730350e4a3b51f2388a4bb53214196bf87db88cdef81601e815181106200083457fe5b60200260200101906001600160a01b031690816001600160a01b03168152505073f8795c4fd7bded7e2f30559759e8fac0c8e1137081601f815181106200087757fe5b60200260200101906001600160a01b031690816001600160a01b0316815250507366c9c684fa5599c639ef9bd608b220ea0fe1bdcb81602081518110620008ba57fe5b60200260200101906001600160a01b031690816001600160a01b03168152505073e9e280ad18cd055aefc5fb5a9a8903cc650ef37981602181518110620008fd57fe5b60200260200101906001600160a01b031690816001600160a01b03168152505073f1b7280708e359debd871f38477c3535e2a5f145816022815181106200094057fe5b60200260200101906001600160a01b031690816001600160a01b031681525050736eb19e67d3625bfa049efbfdb9d008b7438440f2816023815181106200098357fe5b60200260200101906001600160a01b031690816001600160a01b03168152505073d19aba0d1e2b30e764ed315ef26cadeb0f430fb581602481518110620009c657fe5b60200260200101906001600160a01b031690816001600160a01b03168152505073484b896ee33c0e0b75f83a7fb240a81c0a2e31708160258151811062000a0957fe5b60200260200101906001600160a01b031690816001600160a01b0316815250507362f2379e46d96082d99b22fb32cb32cd03f02e568160268151811062000a4c57fe5b60200260200101906001600160a01b031690816001600160a01b031681525050733363460f013d0d4d5f7a95fd1c71d084749634e78160278151811062000a8f57fe5b60200260200101906001600160a01b031690816001600160a01b031681525050738b0cfe38d019537626bb3e470aa445f2828eda358160288151811062000ad257fe5b60200260200101906001600160a01b031690816001600160a01b031681525050731536b7f6d1db9343c777ae6e76e6d1aa5b2032be8160298151811062000b1557fe5b60200260200101906001600160a01b031690816001600160a01b03168152505073443bdf162b3b72f1cd84070eb722ab7907b198fd81602a8151811062000b5857fe5b60200260200101906001600160a01b031690816001600160a01b0316815250507328aa44764d975e82d56b05bebd3afa2e8ba41d8581602b8151811062000b9b57fe5b60200260200101906001600160a01b031690816001600160a01b031681525050731f0905f7f26cd26c3557e75aad7a456f97a1693581602c8151811062000bde57fe5b60200260200101906001600160a01b031690816001600160a01b0316815250507398ab5765061f54694950b5a4d8cee38614bfff4581602d8151811062000c2157fe5b60200260200101906001600160a01b031690816001600160a01b031681525050732b02ed8faab38e609106dbecc61deae11f5a0d4081602e8151811062000c6457fe5b60200260200101906001600160a01b031690816001600160a01b031681525050736df136da0fe91c4bcba6275dfa072823d671090181602f8151811062000ca757fe5b60200260200101906001600160a01b031690816001600160a01b03168152505073358aa690691335bfcfeefbe1c49bd6e6c405f5d98160308151811062000cea57fe5b60200260200101906001600160a01b031690816001600160a01b0316815250507309a83011fd8c10ff12687c3055f1a3dd280fb8d08160318151811062000d2d57fe5b60200260200101906001600160a01b031690816001600160a01b03168152505073c151b858c2e080fa4cd20bdd8f93d0b759407e2b8160328151811062000d7057fe5b6001600160a01b039092166020928302919091019091015260005b815181101562000e8557600180600084848151811062000da757fe5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548160ff021916908315150217905550600160008084848151811062000df857fe5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548160ff02191690831515021790555060016002600084848151811062000e4a57fe5b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff191691151591909117905560010162000d8b565b50506114668062000e976000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80637d2467391161008c578063d5b9892711610066578063d5b9892714610369578063e7b4cac614610416578063e840f6a31461043c578063f25e88b014610462576100cf565b80637d246739146101d7578063a8b40b2d146102b2578063b58729581461031e576100cf565b80630e744b84146100d457806315164cf9146100f3578063429b62e514610122578063465d83101461015c578063574e1714146101825780635cace31a146101a8575b600080fd5b6100f1600480360360208110156100ea57600080fd5b5035610488565b005b6100f16004803603604081101561010957600080fd5b5080356001600160a01b0316906020013560ff16610568565b6101486004803603602081101561013857600080fd5b50356001600160a01b03166106e3565b604080519115158252519081900360200190f35b6100f16004803603602081101561017257600080fd5b50356001600160a01b03166106f8565b6100f16004803603602081101561019857600080fd5b50356001600160a01b031661097e565b6100f1600480360360408110156101be57600080fd5b5080356001600160a01b0316906020013560ff16610c06565b6101f4600480360360208110156101ed57600080fd5b5035610d80565b604051808660070b60070b8152602001858152602001846001600160a01b03166001600160a01b03168152602001836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561027357818101518382015260200161025b565b50505050905090810190601f1680156102a05780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b6102d8600480360360208110156102c857600080fd5b50356001600160a01b0316610e4a565b60408051600796870b870b815260ff90951660208601526001600160a01b039384168582015291909216606084015290830b90920b608082015290519081900360a00190f35b61033b6004803603602081101561033457600080fd5b5035610e90565b60408051600794850b90940b845260208401929092526001600160a01b031682820152519081900360600190f35b6100f16004803603604081101561037f57600080fd5b813591908101906040810160208201356401000000008111156103a157600080fd5b8201836020820111156103b357600080fd5b803590602001918460018302840111640100000000831117156103d557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610ebf945050505050565b6101486004803603602081101561042c57600080fd5b50356001600160a01b0316611054565b6102d86004803603602081101561045257600080fd5b50356001600160a01b0316611069565b6101486004803603602081101561047857600080fd5b50356001600160a01b03166110af565b3360009081526020819052604090205460ff166104ec576040805162461bcd60e51b815260206004820152601860248201527f436c69656e74206973206e6f7420726567697374657265640000000000000000604482015290519081900360640190fd5b6040805160608101825242600790810b825260208083018581523384860190815260009687526003909252939094209151825467ffffffffffffffff191667ffffffffffffffff9190920b16178155905160018201559051600290910180546001600160a01b0319166001600160a01b03909216919091179055565b3360009081526001602052604090205460ff166105c3576040805162461bcd60e51b815260206004820152601460248201527329b2b73232b91036bab9ba1031329030b236b4b760611b604482015290519081900360640190fd5b6105cd82826110c4565b15610615576040805162461bcd60e51b815260206004820152601360248201527212d95e48185b1c9958591e48195e1a5cdd1959606a1b604482015290519081900360640190fd5b61061f82826111ae565b6001600160a01b038084166000908152600560209081526040918290208451815492860151938601518516600160481b02600160481b600160e81b031960ff909516600160401b0260ff60401b19600793840b67ffffffffffffffff90811667ffffffffffffffff19978816179190911691909117959095161782556060860151600183018054919096166001600160a01b031990911617909455608090940151600390940180549490930b90911692169190911790556106df826106f8565b5050565b60016020526000908152604090205460ff1681565b3360009081526001602052604090205460ff16610753576040805162461bcd60e51b815260206004820152601460248201527329b2b73232b91036bab9ba1031329030b236b4b760611b604482015290519081900360640190fd5b6001600160a01b03808216600090815260056020908152604091829020825160a0810184528154600781810b810b810b835260ff600160401b83041694830194909452600160481b900485169381019390935260018101549093166060830152600390920154820b820b90910b60808201526107ce906111ee565b1561080a5760405162461bcd60e51b81526004018080602001828103825260228152602001806114106022913960400191505060405180910390fd5b6001600160a01b038116600090815260056020908152604080832033845260020190915290205460ff1615610886576040805162461bcd60e51b815260206004820152601f60248201527f416c726561647920766f7465642077697468207468652073616d65206b657900604482015290519081900360640190fd5b6001600160a01b0381811660008181526005602081815260408084203385526002810183528185208054600160ff1990911681179091559590945291815260038301805467ffffffffffffffff19811667ffffffffffffffff600792830b8801830b161791829055835160a081018552855480830b830b830b825260ff600160401b82041694820194909452600160481b9093048716938301939093529390920154909316606082015290820b820b90910b608082015261094690611200565b1561097b576001600160a01b03811660009081526005602052604090205461097b908290600160401b900460ff166001611211565b50565b3360009081526001602052604090205460ff166109d9576040805162461bcd60e51b815260206004820152601460248201527329b2b73232b91036bab9ba1031329030b236b4b760611b604482015290519081900360640190fd5b6001600160a01b03808216600090815260066020908152604091829020825160a0810184528154600781810b810b810b835260ff600160401b83041694830194909452600160481b900485169381019390935260018101549093166060830152600390920154820b820b90910b6080820152610a54906111ee565b15610a905760405162461bcd60e51b81526004018080602001828103825260228152602001806114106022913960400191505060405180910390fd5b6001600160a01b038116600090815260066020908152604080832033845260020190915290205460ff1615610b0c576040805162461bcd60e51b815260206004820152601f60248201527f416c726561647920766f7465642077697468207468652073616d65206b657900604482015290519081900360640190fd5b6001600160a01b0381811660008181526006602090815260408083203384526002810183528184208054600160ff1990911681179091559484526003908101805467ffffffffffffffff19811667ffffffffffffffff600792830b8901830b16179091556005845293829020825160a081018452815480870b870b870b825260ff600160401b82041695820195909552600160481b9094048716928401929092529381015490941660608201529290910154810b810b900b6080820152610bd290611200565b1561097b576001600160a01b03811660009081526006602052604081205461097b918391600160401b900460ff1690611211565b3360009081526001602052604090205460ff16610c61576040805162461bcd60e51b815260206004820152601460248201527329b2b73232b91036bab9ba1031329030b236b4b760611b604482015290519081900360640190fd5b610c6b82826110c4565b1515600114610cb6576040805162461bcd60e51b815260206004820152601260248201527112d95e48191bd95cc81b9bdd08195e1a5cdd60721b604482015290519081900360640190fd5b610cc082826111ae565b6001600160a01b038084166000908152600660209081526040918290208451815492860151938601518516600160481b02600160481b600160e81b031960ff909516600160401b0260ff60401b19600793840b67ffffffffffffffff90811667ffffffffffffffff19978816179190911691909117959095161782556060860151600183018054919096166001600160a01b031990911617909455608090940151600390940180549490930b90911692169190911790556106df8261097e565b6004602081815260009283526040928390208054600180830154600280850154600386015497860180548a516101009682161596909602600019011692909204601f810188900488028501880190995288845260079490940b9791966001600160a01b039485169694169493909190830182828015610e405780601f10610e1557610100808354040283529160200191610e40565b820191906000526020600020905b815481529060010190602001808311610e2357829003601f168201915b5050505050905085565b600560205260009081526040902080546001820154600390920154600782810b9360ff600160401b850416936001600160a01b03600160481b9091048116939116910b85565b60036020526000908152604090208054600182015460029092015460079190910b91906001600160a01b031683565b3360009081526001602052604090205460ff16610f1a576040805162461bcd60e51b815260206004820152601460248201527329b2b73232b91036bab9ba1031329030b236b4b760611b604482015290519081900360640190fd5b6000610f2683836112b1565b6001600160a01b03811660009081526002602052604090205490915060ff16610f96576040805162461bcd60e51b815260206004820152601e60248201527f4d616e756661637475726572206973206e6f7420726567697374657265640000604482015290519081900360640190fd5b6040805160a08101825242600790810b82526020808301878152338486019081526001600160a01b0387811660608701908152608087018a815260008c815260048088529990208851815467ffffffffffffffff191667ffffffffffffffff9190990b169790971787559351600187015591516002860180546001600160a01b0319908116928416929092179055915160038601805490931691161790555180519394929361104c938501929190910190611346565b505050505050565b60006020819052908152604090205460ff1681565b600660205260009081526040902080546001820154600390920154600782810b9360ff600160401b850416936001600160a01b03600160481b9091048116939116910b85565b60026020526000908152604090205460ff1681565b600060018260ff16101580156110de575060038260ff1611155b611122576040805162461bcd60e51b815260206004820152601060248201526f496e76616c6964204b6579205479706560801b604482015290519081900360640190fd5b60ff82166001141561115057506001600160a01b03821660009081526001602052604090205460ff166111a8565b60ff82166002141561117e57506001600160a01b03821660009081526020819052604090205460ff166111a8565b60ff8216600314156111a857506001600160a01b03821660009081526002602052604090205460ff165b92915050565b6111b66113c4565b6111be6113c4565b42600790810b900b815260ff831660208201526001600160a01b0384166040820152336060820152905092915050565b51610e1001600790810b4290910b1390565b60800151600260079190910b121590565b60ff821660011415611246576001600160a01b0383166000908152600160205260409020805460ff19168215151790556112ac565b60ff82166002141561127b576001600160a01b0383166000908152602081905260409020805460ff19168215151790556112ac565b60ff8216600314156112ac576001600160a01b0383166000908152600260205260409020805460ff19168215151790555b505050565b600081516041146112c4575060006111a8565b6020828101516040808501516041860151825160008152808601808552899052601b90910160ff811682850152606082018590526080820183905292519394919360019260a0808401939192601f1981019281900390910190855afa158015611331573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061138757805160ff19168380011785556113b4565b828001600101855582156113b4579182015b828111156113b4578251825591602001919060010190611399565b506113c09291506113f2565b5090565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915290565b61140c91905b808211156113c057600081556001016113f8565b9056fe50726f706f73616c20646f6573206e6f74206578697374206f722065787069726564a265627a7a72315820aa143fce52e68cefc6196d41be251f59506394b6a73efe8f8f401e14c652623f64736f6c63430005110032"

// DeployECC deploys a new Ethereum contract, binding an instance of ECC to it.
func DeployECC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECC, error) {
	parsed, err := abi.JSON(strings.NewReader(ECCABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ECCBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECC{ECCCaller: ECCCaller{contract: contract}, ECCTransactor: ECCTransactor{contract: contract}, ECCFilterer: ECCFilterer{contract: contract}}, nil
}

// ECC is an auto generated Go binding around an Ethereum contract.
type ECC struct {
	ECCCaller     // Read-only binding to the contract
	ECCTransactor // Write-only binding to the contract
	ECCFilterer   // Log filterer for contract events
}

// ECCCaller is an auto generated read-only Go binding around an Ethereum contract.
type ECCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECCSession struct {
	Contract     *ECC              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECCCallerSession struct {
	Contract *ECCCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ECCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECCTransactorSession struct {
	Contract     *ECCTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECCRaw is an auto generated low-level Go binding around an Ethereum contract.
type ECCRaw struct {
	Contract *ECC // Generic contract binding to access the raw methods on
}

// ECCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECCCallerRaw struct {
	Contract *ECCCaller // Generic read-only contract binding to access the raw methods on
}

// ECCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECCTransactorRaw struct {
	Contract *ECCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewECC creates a new instance of ECC, bound to a specific deployed contract.
func NewECC(address common.Address, backend bind.ContractBackend) (*ECC, error) {
	contract, err := bindECC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECC{ECCCaller: ECCCaller{contract: contract}, ECCTransactor: ECCTransactor{contract: contract}, ECCFilterer: ECCFilterer{contract: contract}}, nil
}

// NewECCCaller creates a new read-only instance of ECC, bound to a specific deployed contract.
func NewECCCaller(address common.Address, caller bind.ContractCaller) (*ECCCaller, error) {
	contract, err := bindECC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECCCaller{contract: contract}, nil
}

// NewECCTransactor creates a new write-only instance of ECC, bound to a specific deployed contract.
func NewECCTransactor(address common.Address, transactor bind.ContractTransactor) (*ECCTransactor, error) {
	contract, err := bindECC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECCTransactor{contract: contract}, nil
}

// NewECCFilterer creates a new log filterer instance of ECC, bound to a specific deployed contract.
func NewECCFilterer(address common.Address, filterer bind.ContractFilterer) (*ECCFilterer, error) {
	contract, err := bindECC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECCFilterer{contract: contract}, nil
}

// bindECC binds a generic wrapper to an already deployed contract.
func bindECC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECC *ECCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECC.Contract.ECCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECC *ECCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECC.Contract.ECCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECC *ECCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECC.Contract.ECCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECC *ECCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECC *ECCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECC *ECCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECC.Contract.contract.Transact(opts, method, params...)
}

// AddKeyProposals is a free data retrieval call binding the contract method 0xa8b40b2d.
//
// Solidity: function addKeyProposals(address ) view returns(int64 unixTime, uint8 keyType, address key, address proposer, int64 voteCount)
func (_ECC *ECCCaller) AddKeyProposals(opts *bind.CallOpts, arg0 common.Address) (struct {
	UnixTime  int64
	KeyType   uint8
	Key       common.Address
	Proposer  common.Address
	VoteCount int64
}, error) {
	var out []interface{}
	err := _ECC.contract.Call(opts, &out, "addKeyProposals", arg0)

	outstruct := new(struct {
		UnixTime  int64
		KeyType   uint8
		Key       common.Address
		Proposer  common.Address
		VoteCount int64
	})

	outstruct.UnixTime = out[0].(int64)
	outstruct.KeyType = out[1].(uint8)
	outstruct.Key = out[2].(common.Address)
	outstruct.Proposer = out[3].(common.Address)
	outstruct.VoteCount = out[4].(int64)

	return *outstruct, err

}

// AddKeyProposals is a free data retrieval call binding the contract method 0xa8b40b2d.
//
// Solidity: function addKeyProposals(address ) view returns(int64 unixTime, uint8 keyType, address key, address proposer, int64 voteCount)
func (_ECC *ECCSession) AddKeyProposals(arg0 common.Address) (struct {
	UnixTime  int64
	KeyType   uint8
	Key       common.Address
	Proposer  common.Address
	VoteCount int64
}, error) {
	return _ECC.Contract.AddKeyProposals(&_ECC.CallOpts, arg0)
}

// AddKeyProposals is a free data retrieval call binding the contract method 0xa8b40b2d.
//
// Solidity: function addKeyProposals(address ) view returns(int64 unixTime, uint8 keyType, address key, address proposer, int64 voteCount)
func (_ECC *ECCCallerSession) AddKeyProposals(arg0 common.Address) (struct {
	UnixTime  int64
	KeyType   uint8
	Key       common.Address
	Proposer  common.Address
	VoteCount int64
}, error) {
	return _ECC.Contract.AddKeyProposals(&_ECC.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_ECC *ECCCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ECC.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_ECC *ECCSession) Admins(arg0 common.Address) (bool, error) {
	return _ECC.Contract.Admins(&_ECC.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_ECC *ECCCallerSession) Admins(arg0 common.Address) (bool, error) {
	return _ECC.Contract.Admins(&_ECC.CallOpts, arg0)
}

// Devices is a free data retrieval call binding the contract method 0xe7b4cac6.
//
// Solidity: function devices(address ) view returns(bool)
func (_ECC *ECCCaller) Devices(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ECC.contract.Call(opts, &out, "devices", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Devices is a free data retrieval call binding the contract method 0xe7b4cac6.
//
// Solidity: function devices(address ) view returns(bool)
func (_ECC *ECCSession) Devices(arg0 common.Address) (bool, error) {
	return _ECC.Contract.Devices(&_ECC.CallOpts, arg0)
}

// Devices is a free data retrieval call binding the contract method 0xe7b4cac6.
//
// Solidity: function devices(address ) view returns(bool)
func (_ECC *ECCCallerSession) Devices(arg0 common.Address) (bool, error) {
	return _ECC.Contract.Devices(&_ECC.CallOpts, arg0)
}

// Firmwares is a free data retrieval call binding the contract method 0x7d246739.
//
// Solidity: function firmwares(bytes32 ) view returns(int64 unixTime, bytes32 hash, address admin, address mf, bytes mfSig)
func (_ECC *ECCCaller) Firmwares(opts *bind.CallOpts, arg0 [32]byte) (struct {
	UnixTime int64
	Hash     [32]byte
	Admin    common.Address
	Mf       common.Address
	MfSig    []byte
}, error) {
	var out []interface{}
	err := _ECC.contract.Call(opts, &out, "firmwares", arg0)

	outstruct := new(struct {
		UnixTime int64
		Hash     [32]byte
		Admin    common.Address
		Mf       common.Address
		MfSig    []byte
	})

	outstruct.UnixTime = out[0].(int64)
	outstruct.Hash = out[1].([32]byte)
	outstruct.Admin = out[2].(common.Address)
	outstruct.Mf = out[3].(common.Address)
	outstruct.MfSig = out[4].([]byte)

	return *outstruct, err

}

// Firmwares is a free data retrieval call binding the contract method 0x7d246739.
//
// Solidity: function firmwares(bytes32 ) view returns(int64 unixTime, bytes32 hash, address admin, address mf, bytes mfSig)
func (_ECC *ECCSession) Firmwares(arg0 [32]byte) (struct {
	UnixTime int64
	Hash     [32]byte
	Admin    common.Address
	Mf       common.Address
	MfSig    []byte
}, error) {
	return _ECC.Contract.Firmwares(&_ECC.CallOpts, arg0)
}

// Firmwares is a free data retrieval call binding the contract method 0x7d246739.
//
// Solidity: function firmwares(bytes32 ) view returns(int64 unixTime, bytes32 hash, address admin, address mf, bytes mfSig)
func (_ECC *ECCCallerSession) Firmwares(arg0 [32]byte) (struct {
	UnixTime int64
	Hash     [32]byte
	Admin    common.Address
	Mf       common.Address
	MfSig    []byte
}, error) {
	return _ECC.Contract.Firmwares(&_ECC.CallOpts, arg0)
}

// Mf is a free data retrieval call binding the contract method 0xf25e88b0.
//
// Solidity: function mf(address ) view returns(bool)
func (_ECC *ECCCaller) Mf(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ECC.contract.Call(opts, &out, "mf", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Mf is a free data retrieval call binding the contract method 0xf25e88b0.
//
// Solidity: function mf(address ) view returns(bool)
func (_ECC *ECCSession) Mf(arg0 common.Address) (bool, error) {
	return _ECC.Contract.Mf(&_ECC.CallOpts, arg0)
}

// Mf is a free data retrieval call binding the contract method 0xf25e88b0.
//
// Solidity: function mf(address ) view returns(bool)
func (_ECC *ECCCallerSession) Mf(arg0 common.Address) (bool, error) {
	return _ECC.Contract.Mf(&_ECC.CallOpts, arg0)
}

// RemoveKeyProposals is a free data retrieval call binding the contract method 0xe840f6a3.
//
// Solidity: function removeKeyProposals(address ) view returns(int64 unixTime, uint8 keyType, address key, address proposer, int64 voteCount)
func (_ECC *ECCCaller) RemoveKeyProposals(opts *bind.CallOpts, arg0 common.Address) (struct {
	UnixTime  int64
	KeyType   uint8
	Key       common.Address
	Proposer  common.Address
	VoteCount int64
}, error) {
	var out []interface{}
	err := _ECC.contract.Call(opts, &out, "removeKeyProposals", arg0)

	outstruct := new(struct {
		UnixTime  int64
		KeyType   uint8
		Key       common.Address
		Proposer  common.Address
		VoteCount int64
	})

	outstruct.UnixTime = out[0].(int64)
	outstruct.KeyType = out[1].(uint8)
	outstruct.Key = out[2].(common.Address)
	outstruct.Proposer = out[3].(common.Address)
	outstruct.VoteCount = out[4].(int64)

	return *outstruct, err

}

// RemoveKeyProposals is a free data retrieval call binding the contract method 0xe840f6a3.
//
// Solidity: function removeKeyProposals(address ) view returns(int64 unixTime, uint8 keyType, address key, address proposer, int64 voteCount)
func (_ECC *ECCSession) RemoveKeyProposals(arg0 common.Address) (struct {
	UnixTime  int64
	KeyType   uint8
	Key       common.Address
	Proposer  common.Address
	VoteCount int64
}, error) {
	return _ECC.Contract.RemoveKeyProposals(&_ECC.CallOpts, arg0)
}

// RemoveKeyProposals is a free data retrieval call binding the contract method 0xe840f6a3.
//
// Solidity: function removeKeyProposals(address ) view returns(int64 unixTime, uint8 keyType, address key, address proposer, int64 voteCount)
func (_ECC *ECCCallerSession) RemoveKeyProposals(arg0 common.Address) (struct {
	UnixTime  int64
	KeyType   uint8
	Key       common.Address
	Proposer  common.Address
	VoteCount int64
}, error) {
	return _ECC.Contract.RemoveKeyProposals(&_ECC.CallOpts, arg0)
}

// Timestamps is a free data retrieval call binding the contract method 0xb5872958.
//
// Solidity: function timestamps(bytes32 ) view returns(int64 unixTime, bytes32 hash, address device)
func (_ECC *ECCCaller) Timestamps(opts *bind.CallOpts, arg0 [32]byte) (struct {
	UnixTime int64
	Hash     [32]byte
	Device   common.Address
}, error) {
	var out []interface{}
	err := _ECC.contract.Call(opts, &out, "timestamps", arg0)

	outstruct := new(struct {
		UnixTime int64
		Hash     [32]byte
		Device   common.Address
	})

	outstruct.UnixTime = out[0].(int64)
	outstruct.Hash = out[1].([32]byte)
	outstruct.Device = out[2].(common.Address)

	return *outstruct, err

}

// Timestamps is a free data retrieval call binding the contract method 0xb5872958.
//
// Solidity: function timestamps(bytes32 ) view returns(int64 unixTime, bytes32 hash, address device)
func (_ECC *ECCSession) Timestamps(arg0 [32]byte) (struct {
	UnixTime int64
	Hash     [32]byte
	Device   common.Address
}, error) {
	return _ECC.Contract.Timestamps(&_ECC.CallOpts, arg0)
}

// Timestamps is a free data retrieval call binding the contract method 0xb5872958.
//
// Solidity: function timestamps(bytes32 ) view returns(int64 unixTime, bytes32 hash, address device)
func (_ECC *ECCCallerSession) Timestamps(arg0 [32]byte) (struct {
	UnixTime int64
	Hash     [32]byte
	Device   common.Address
}, error) {
	return _ECC.Contract.Timestamps(&_ECC.CallOpts, arg0)
}

// ProposeAddKey is a paid mutator transaction binding the contract method 0x15164cf9.
//
// Solidity: function proposeAddKey(address key, uint8 keyType) returns()
func (_ECC *ECCTransactor) ProposeAddKey(opts *bind.TransactOpts, key common.Address, keyType uint8) (*types.Transaction, error) {
	return _ECC.contract.Transact(opts, "proposeAddKey", key, keyType)
}

// ProposeAddKey is a paid mutator transaction binding the contract method 0x15164cf9.
//
// Solidity: function proposeAddKey(address key, uint8 keyType) returns()
func (_ECC *ECCSession) ProposeAddKey(key common.Address, keyType uint8) (*types.Transaction, error) {
	return _ECC.Contract.ProposeAddKey(&_ECC.TransactOpts, key, keyType)
}

// ProposeAddKey is a paid mutator transaction binding the contract method 0x15164cf9.
//
// Solidity: function proposeAddKey(address key, uint8 keyType) returns()
func (_ECC *ECCTransactorSession) ProposeAddKey(key common.Address, keyType uint8) (*types.Transaction, error) {
	return _ECC.Contract.ProposeAddKey(&_ECC.TransactOpts, key, keyType)
}

// ProposeRemoveKey is a paid mutator transaction binding the contract method 0x5cace31a.
//
// Solidity: function proposeRemoveKey(address key, uint8 keyType) returns()
func (_ECC *ECCTransactor) ProposeRemoveKey(opts *bind.TransactOpts, key common.Address, keyType uint8) (*types.Transaction, error) {
	return _ECC.contract.Transact(opts, "proposeRemoveKey", key, keyType)
}

// ProposeRemoveKey is a paid mutator transaction binding the contract method 0x5cace31a.
//
// Solidity: function proposeRemoveKey(address key, uint8 keyType) returns()
func (_ECC *ECCSession) ProposeRemoveKey(key common.Address, keyType uint8) (*types.Transaction, error) {
	return _ECC.Contract.ProposeRemoveKey(&_ECC.TransactOpts, key, keyType)
}

// ProposeRemoveKey is a paid mutator transaction binding the contract method 0x5cace31a.
//
// Solidity: function proposeRemoveKey(address key, uint8 keyType) returns()
func (_ECC *ECCTransactorSession) ProposeRemoveKey(key common.Address, keyType uint8) (*types.Transaction, error) {
	return _ECC.Contract.ProposeRemoveKey(&_ECC.TransactOpts, key, keyType)
}

// TimestampHash is a paid mutator transaction binding the contract method 0x0e744b84.
//
// Solidity: function timestampHash(bytes32 hash) returns()
func (_ECC *ECCTransactor) TimestampHash(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _ECC.contract.Transact(opts, "timestampHash", hash)
}

// TimestampHash is a paid mutator transaction binding the contract method 0x0e744b84.
//
// Solidity: function timestampHash(bytes32 hash) returns()
func (_ECC *ECCSession) TimestampHash(hash [32]byte) (*types.Transaction, error) {
	return _ECC.Contract.TimestampHash(&_ECC.TransactOpts, hash)
}

// TimestampHash is a paid mutator transaction binding the contract method 0x0e744b84.
//
// Solidity: function timestampHash(bytes32 hash) returns()
func (_ECC *ECCTransactorSession) TimestampHash(hash [32]byte) (*types.Transaction, error) {
	return _ECC.Contract.TimestampHash(&_ECC.TransactOpts, hash)
}

// UpdateFirmware is a paid mutator transaction binding the contract method 0xd5b98927.
//
// Solidity: function updateFirmware(bytes32 hash, bytes mfSig) returns()
func (_ECC *ECCTransactor) UpdateFirmware(opts *bind.TransactOpts, hash [32]byte, mfSig []byte) (*types.Transaction, error) {
	return _ECC.contract.Transact(opts, "updateFirmware", hash, mfSig)
}

// UpdateFirmware is a paid mutator transaction binding the contract method 0xd5b98927.
//
// Solidity: function updateFirmware(bytes32 hash, bytes mfSig) returns()
func (_ECC *ECCSession) UpdateFirmware(hash [32]byte, mfSig []byte) (*types.Transaction, error) {
	return _ECC.Contract.UpdateFirmware(&_ECC.TransactOpts, hash, mfSig)
}

// UpdateFirmware is a paid mutator transaction binding the contract method 0xd5b98927.
//
// Solidity: function updateFirmware(bytes32 hash, bytes mfSig) returns()
func (_ECC *ECCTransactorSession) UpdateFirmware(hash [32]byte, mfSig []byte) (*types.Transaction, error) {
	return _ECC.Contract.UpdateFirmware(&_ECC.TransactOpts, hash, mfSig)
}

// VoteAddKey is a paid mutator transaction binding the contract method 0x465d8310.
//
// Solidity: function voteAddKey(address key) returns()
func (_ECC *ECCTransactor) VoteAddKey(opts *bind.TransactOpts, key common.Address) (*types.Transaction, error) {
	return _ECC.contract.Transact(opts, "voteAddKey", key)
}

// VoteAddKey is a paid mutator transaction binding the contract method 0x465d8310.
//
// Solidity: function voteAddKey(address key) returns()
func (_ECC *ECCSession) VoteAddKey(key common.Address) (*types.Transaction, error) {
	return _ECC.Contract.VoteAddKey(&_ECC.TransactOpts, key)
}

// VoteAddKey is a paid mutator transaction binding the contract method 0x465d8310.
//
// Solidity: function voteAddKey(address key) returns()
func (_ECC *ECCTransactorSession) VoteAddKey(key common.Address) (*types.Transaction, error) {
	return _ECC.Contract.VoteAddKey(&_ECC.TransactOpts, key)
}

// VoteRemoveKey is a paid mutator transaction binding the contract method 0x574e1714.
//
// Solidity: function voteRemoveKey(address key) returns()
func (_ECC *ECCTransactor) VoteRemoveKey(opts *bind.TransactOpts, key common.Address) (*types.Transaction, error) {
	return _ECC.contract.Transact(opts, "voteRemoveKey", key)
}

// VoteRemoveKey is a paid mutator transaction binding the contract method 0x574e1714.
//
// Solidity: function voteRemoveKey(address key) returns()
func (_ECC *ECCSession) VoteRemoveKey(key common.Address) (*types.Transaction, error) {
	return _ECC.Contract.VoteRemoveKey(&_ECC.TransactOpts, key)
}

// VoteRemoveKey is a paid mutator transaction binding the contract method 0x574e1714.
//
// Solidity: function voteRemoveKey(address key) returns()
func (_ECC *ECCTransactorSession) VoteRemoveKey(key common.Address) (*types.Transaction, error) {
	return _ECC.Contract.VoteRemoveKey(&_ECC.TransactOpts, key)
}
