// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc777

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

// Erc777MetaData contains all meta data concerning the Erc777 contract.
var Erc777MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"defaultOperators_\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"AuthorizedOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"RevokedOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"Sent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"authorizeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultOperators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"granularity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"isOperatorFor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"operatorBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"operatorSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"revokeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001bcf38038062001bcf8339810160408190526200003491620003f1565b82516200004990600290602086019062000215565b5081516200005f90600390602085019062000215565b50805162000075906004906020840190620002a4565b5060005b8151811015620000e5576001600560008484815181106200009e576200009e62000508565b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff191691151591909117905580620000dc816200051e565b91505062000079565b506040516329965a1d60e01b815230600482018190527fac7fbab5f54a3ca8194167523c6753bfeb96a445279294b6125b68cce217705460248301526044820152731820a4b7618bde71dce8cdc73aab6c95905fad24906329965a1d90606401600060405180830381600087803b1580156200016057600080fd5b505af115801562000175573d6000803e3d6000fd5b50506040516329965a1d60e01b815230600482018190527faea199e31a596269b42cdafd93407f14436db6e4cad65417994c2eb37381e05a60248301526044820152731820a4b7618bde71dce8cdc73aab6c95905fad2492506329965a1d9150606401600060405180830381600087803b158015620001f357600080fd5b505af115801562000208573d6000803e3d6000fd5b5050505050505062000585565b828054620002239062000548565b90600052602060002090601f01602090048101928262000247576000855562000292565b82601f106200026257805160ff191683800117855562000292565b8280016001018555821562000292579182015b828111156200029257825182559160200191906001019062000275565b50620002a0929150620002fc565b5090565b82805482825590600052602060002090810192821562000292579160200282015b828111156200029257825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620002c5565b5b80821115620002a05760008155600101620002fd565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b038111828210171562000354576200035462000313565b604052919050565b600082601f8301126200036e57600080fd5b81516001600160401b038111156200038a576200038a62000313565b6020620003a0601f8301601f1916820162000329565b8281528582848701011115620003b557600080fd5b60005b83811015620003d5578581018301518282018401528201620003b8565b83811115620003e75760008385840101525b5095945050505050565b6000806000606084860312156200040757600080fd5b83516001600160401b03808211156200041f57600080fd5b6200042d878388016200035c565b94506020915081860151818111156200044557600080fd5b62000453888289016200035c565b9450506040860151818111156200046957600080fd5b8601601f810188136200047b57600080fd5b80518281111562000490576200049062000313565b8060051b9250620004a384840162000329565b818152928201840192848101908a851115620004be57600080fd5b928501925b84841015620004f857835192506001600160a01b0383168314620004e75760008081fd5b8282529285019290850190620004c3565b8096505050505050509250925092565b634e487b7160e01b600052603260045260246000fd5b60006000198214156200054157634e487b7160e01b600052601160045260246000fd5b5060010190565b600181811c908216806200055d57607f821691505b602082108114156200057f57634e487b7160e01b600052602260045260246000fd5b50919050565b61163a80620005956000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c8063959b8c3f116100a2578063d95b637111610071578063d95b63711461022b578063dd62ed3e1461023e578063fad8b32a14610277578063fc673c4f1461028a578063fe9d93031461029d57600080fd5b8063959b8c3f146101ea57806395d89b41146101fd5780639bd9bbc614610205578063a9059cbb1461021857600080fd5b806323b872dd116100e957806323b872dd14610183578063313ce56714610196578063556f0dc7146101a557806362ad1b83146101ac57806370a08231146101c157600080fd5b806306e485381461011b57806306fdde0314610139578063095ea7b31461014e57806318160ddd14610171575b600080fd5b6101236102b0565b60405161013091906110ae565b60405180910390f35b610141610312565b6040516101309190611148565b61016161015c366004611173565b61039b565b6040519015158152602001610130565b6001545b604051908152602001610130565b61016161019136600461119f565b6103b3565b60405160128152602001610130565b6001610175565b6101bf6101ba366004611283565b6103f9565b005b6101756101cf366004611316565b6001600160a01b031660009081526020819052604090205490565b6101bf6101f8366004611316565b61043e565b61014161055c565b6101bf610213366004611333565b61056b565b610161610226366004611173565b61058e565b61016161023936600461138c565b6105c6565b61017561024c36600461138c565b6001600160a01b03918216600090815260086020908152604080832093909416825291909152205490565b6101bf610285366004611316565b610668565b6101bf6102983660046113c5565b610784565b6101bf6102ab366004611445565b6107bc565b6060600480548060200260200160405190810160405280929190818152602001828054801561030857602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116102ea575b5050505050905090565b6060600280546103219061148c565b80601f016020809104026020016040519081016040528092919081815260200182805461034d9061148c565b80156103085780601f1061036f57610100808354040283529160200191610308565b820191906000526020600020905b81548152906001019060200180831161037d57509395945050505050565b6000336103a98185856107db565b5060019392505050565b6000336103c1858285610902565b6103ee8585856040518060200160405280600081525060405180602001604052806000815250600061098e565b506001949350505050565b61040333866105c6565b6104285760405162461bcd60e51b815260040161041f906114c7565b60405180910390fd5b6104378585858585600161098e565b5050505050565b336001600160a01b03821614156104a35760405162461bcd60e51b8152602060048201526024808201527f4552433737373a20617574686f72697a696e672073656c66206173206f70657260448201526330ba37b960e11b606482015260840161041f565b6001600160a01b03811660009081526005602052604090205460ff16156104f4573360009081526007602090815260408083206001600160a01b03851684529091529020805460ff19169055610523565b3360009081526006602090815260408083206001600160a01b03851684529091529020805460ff191660011790555b60405133906001600160a01b038316907ff4caeb2d6ca8932a215a353d0703c326ec2d81fc68170f320eb2ab49e9df61f990600090a350565b6060600380546103219061148c565b6105893384848460405180602001604052806000815250600161098e565b505050565b60006105bd3384846040518060200160405280600081525060405180602001604052806000815250600061098e565b50600192915050565b6000816001600160a01b0316836001600160a01b0316148061063157506001600160a01b03831660009081526005602052604090205460ff16801561063157506001600160a01b0380831660009081526007602090815260408083209387168352929052205460ff16155b8061066157506001600160a01b0380831660009081526006602090815260408083209387168352929052205460ff165b9392505050565b6001600160a01b0381163314156106cb5760405162461bcd60e51b815260206004820152602160248201527f4552433737373a207265766f6b696e672073656c66206173206f70657261746f6044820152603960f91b606482015260840161041f565b6001600160a01b03811660009081526005602052604090205460ff161561071f573360009081526007602090815260408083206001600160a01b03851684529091529020805460ff1916600117905561074b565b3360009081526006602090815260408083206001600160a01b03851684529091529020805460ff191690555b60405133906001600160a01b038316907f50546e66e5f44d728365dc3908c63bc5cfeeab470722c1677e3073a6ac294aa190600090a350565b61078e33856105c6565b6107aa5760405162461bcd60e51b815260040161041f906114c7565b6107b684848484610a8a565b50505050565b6107d733838360405180602001604052806000815250610a8a565b5050565b6001600160a01b03831661083f5760405162461bcd60e51b815260206004820152602560248201527f4552433737373a20617070726f76652066726f6d20746865207a65726f206164604482015264647265737360d81b606482015260840161041f565b6001600160a01b0382166108a15760405162461bcd60e51b815260206004820152602360248201527f4552433737373a20617070726f766520746f20746865207a65726f206164647260448201526265737360e81b606482015260840161041f565b6001600160a01b0383811660008181526008602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b6001600160a01b0383811660009081526008602090815260408083209386168352929052205460001981146107b657818110156109815760405162461bcd60e51b815260206004820152601e60248201527f4552433737373a20696e73756666696369656e7420616c6c6f77616e63650000604482015260640161041f565b6107b684848484036107db565b6001600160a01b0386166109f35760405162461bcd60e51b815260206004820152602660248201527f4552433737373a207472616e736665722066726f6d20746865207a65726f206160448201526564647265737360d01b606482015260840161041f565b6001600160a01b038516610a555760405162461bcd60e51b8152602060048201526024808201527f4552433737373a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b606482015260840161041f565b33610a64818888888888610c3e565b610a72818888888888610d74565b610a8181888888888888610eda565b50505050505050565b6001600160a01b038416610aeb5760405162461bcd60e51b815260206004820152602260248201527f4552433737373a206275726e2066726f6d20746865207a65726f206164647265604482015261737360f01b606482015260840161041f565b33610afb81866000878787610c3e565b6001600160a01b03851660009081526020819052604090205484811015610b705760405162461bcd60e51b815260206004820152602360248201527f4552433737373a206275726e20616d6f756e7420657863656564732062616c616044820152626e636560e81b606482015260840161041f565b6001600160a01b0386166000908152602081905260408120868303905560018054879290610b9f908490611529565b92505081905550856001600160a01b0316826001600160a01b03167fa78a9be3a7b862d26933ad85fb11d80ef66b8f972d7cbba06621d583943a4098878787604051610bed93929190611540565b60405180910390a36040518581526000906001600160a01b038816907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a3505050505050565b60405163555ddc6560e11b81526001600160a01b03861660048201527f29ddb589b1fb5fc7cf394961c1adf5f8c6454761adf795e67fe149f658abe8956024820152600090731820a4b7618bde71dce8cdc73aab6c95905fad249063aabbb8ca9060440160206040518083038186803b158015610cba57600080fd5b505afa158015610cce573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cf29190611575565b90506001600160a01b03811615610a8157604051633ad5cbc160e11b81526001600160a01b038216906375ab978290610d39908a908a908a908a908a908a90600401611592565b600060405180830381600087803b158015610d5357600080fd5b505af1158015610d67573d6000803e3d6000fd5b5050505050505050505050565b6001600160a01b03851660009081526020819052604090205483811015610ded5760405162461bcd60e51b815260206004820152602760248201527f4552433737373a207472616e7366657220616d6f756e7420657863656564732060448201526662616c616e636560c81b606482015260840161041f565b6001600160a01b03808716600090815260208190526040808220878503905591871681529081208054869290610e249084906115ec565b92505081905550846001600160a01b0316866001600160a01b0316886001600160a01b03167f06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc82614677987878787604051610e7c93929190611540565b60405180910390a4846001600160a01b0316866001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef86604051610ec991815260200190565b60405180910390a350505050505050565b60405163555ddc6560e11b81526001600160a01b03861660048201527fb281fc8c12954d22544db45de3159a39272895b169a852b314f9cc762e44c53b6024820152600090731820a4b7618bde71dce8cdc73aab6c95905fad249063aabbb8ca9060440160206040518083038186803b158015610f5657600080fd5b505afa158015610f6a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f8e9190611575565b90506001600160a01b0381161561100a576040516223de2960e01b81526001600160a01b038216906223de2990610fd3908b908b908b908b908b908b90600401611592565b600060405180830381600087803b158015610fed57600080fd5b505af1158015611001573d6000803e3d6000fd5b505050506110a4565b81156110a4576001600160a01b0386163b156110a45760405162461bcd60e51b815260206004820152604d60248201527f4552433737373a20746f6b656e20726563697069656e7420636f6e747261637460448201527f20686173206e6f20696d706c656d656e74657220666f7220455243373737546f60648201526c1ad95b9cd49958da5c1a595b9d609a1b608482015260a40161041f565b5050505050505050565b6020808252825182820181905260009190848201906040850190845b818110156110ef5783516001600160a01b0316835292840192918401916001016110ca565b50909695505050505050565b6000815180845260005b8181101561112157602081850181015186830182015201611105565b81811115611133576000602083870101525b50601f01601f19169290920160200192915050565b60208152600061066160208301846110fb565b6001600160a01b038116811461117057600080fd5b50565b6000806040838503121561118657600080fd5b82356111918161115b565b946020939093013593505050565b6000806000606084860312156111b457600080fd5b83356111bf8161115b565b925060208401356111cf8161115b565b929592945050506040919091013590565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261120757600080fd5b813567ffffffffffffffff80821115611222576112226111e0565b604051601f8301601f19908116603f0116810190828211818310171561124a5761124a6111e0565b8160405283815286602085880101111561126357600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080600080600060a0868803121561129b57600080fd5b85356112a68161115b565b945060208601356112b68161115b565b935060408601359250606086013567ffffffffffffffff808211156112da57600080fd5b6112e689838a016111f6565b935060808801359150808211156112fc57600080fd5b50611309888289016111f6565b9150509295509295909350565b60006020828403121561132857600080fd5b81356106618161115b565b60008060006060848603121561134857600080fd5b83356113538161115b565b925060208401359150604084013567ffffffffffffffff81111561137657600080fd5b611382868287016111f6565b9150509250925092565b6000806040838503121561139f57600080fd5b82356113aa8161115b565b915060208301356113ba8161115b565b809150509250929050565b600080600080608085870312156113db57600080fd5b84356113e68161115b565b935060208501359250604085013567ffffffffffffffff8082111561140a57600080fd5b611416888389016111f6565b9350606087013591508082111561142c57600080fd5b50611439878288016111f6565b91505092959194509250565b6000806040838503121561145857600080fd5b82359150602083013567ffffffffffffffff81111561147657600080fd5b611482858286016111f6565b9150509250929050565b600181811c908216806114a057607f821691505b602082108114156114c157634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252602c908201527f4552433737373a2063616c6c6572206973206e6f7420616e206f70657261746f60408201526b39103337b9103437b63232b960a11b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b60008282101561153b5761153b611513565b500390565b83815260606020820152600061155960608301856110fb565b828103604084015261156b81856110fb565b9695505050505050565b60006020828403121561158757600080fd5b81516106618161115b565b6001600160a01b0387811682528681166020830152851660408201526060810184905260c0608082018190526000906115cd908301856110fb565b82810360a08401526115df81856110fb565b9998505050505050505050565b600082198211156115ff576115ff611513565b50019056fea26469706673582212201d3d507e7a24f431591c9d6c6d27f4bf9ac941d7395fd9c36aa72e3d8783d6ef64736f6c63430008090033",
}

// Erc777ABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc777MetaData.ABI instead.
var Erc777ABI = Erc777MetaData.ABI

// Erc777Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Erc777MetaData.Bin instead.
var Erc777Bin = Erc777MetaData.Bin

// DeployErc777 deploys a new Ethereum contract, binding an instance of Erc777 to it.
func DeployErc777(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string, defaultOperators_ []common.Address) (common.Address, *types.Transaction, *Erc777, error) {
	parsed, err := Erc777MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Erc777Bin), backend, name_, symbol_, defaultOperators_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Erc777{Erc777Caller: Erc777Caller{contract: contract}, Erc777Transactor: Erc777Transactor{contract: contract}, Erc777Filterer: Erc777Filterer{contract: contract}}, nil
}

// Erc777 is an auto generated Go binding around an Ethereum contract.
type Erc777 struct {
	Erc777Caller     // Read-only binding to the contract
	Erc777Transactor // Write-only binding to the contract
	Erc777Filterer   // Log filterer for contract events
}

// Erc777Caller is an auto generated read-only Go binding around an Ethereum contract.
type Erc777Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc777Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc777Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc777Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc777Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc777Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc777Session struct {
	Contract     *Erc777           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc777CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc777CallerSession struct {
	Contract *Erc777Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Erc777TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc777TransactorSession struct {
	Contract     *Erc777Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc777Raw is an auto generated low-level Go binding around an Ethereum contract.
type Erc777Raw struct {
	Contract *Erc777 // Generic contract binding to access the raw methods on
}

// Erc777CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc777CallerRaw struct {
	Contract *Erc777Caller // Generic read-only contract binding to access the raw methods on
}

// Erc777TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc777TransactorRaw struct {
	Contract *Erc777Transactor // Generic write-only contract binding to access the raw methods on
}

// NewErc777 creates a new instance of Erc777, bound to a specific deployed contract.
func NewErc777(address common.Address, backend bind.ContractBackend) (*Erc777, error) {
	contract, err := bindErc777(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc777{Erc777Caller: Erc777Caller{contract: contract}, Erc777Transactor: Erc777Transactor{contract: contract}, Erc777Filterer: Erc777Filterer{contract: contract}}, nil
}

// NewErc777Caller creates a new read-only instance of Erc777, bound to a specific deployed contract.
func NewErc777Caller(address common.Address, caller bind.ContractCaller) (*Erc777Caller, error) {
	contract, err := bindErc777(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc777Caller{contract: contract}, nil
}

// NewErc777Transactor creates a new write-only instance of Erc777, bound to a specific deployed contract.
func NewErc777Transactor(address common.Address, transactor bind.ContractTransactor) (*Erc777Transactor, error) {
	contract, err := bindErc777(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc777Transactor{contract: contract}, nil
}

// NewErc777Filterer creates a new log filterer instance of Erc777, bound to a specific deployed contract.
func NewErc777Filterer(address common.Address, filterer bind.ContractFilterer) (*Erc777Filterer, error) {
	contract, err := bindErc777(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc777Filterer{contract: contract}, nil
}

// bindErc777 binds a generic wrapper to an already deployed contract.
func bindErc777(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc777ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc777 *Erc777Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc777.Contract.Erc777Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc777 *Erc777Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc777.Contract.Erc777Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc777 *Erc777Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc777.Contract.Erc777Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc777 *Erc777CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc777.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc777 *Erc777TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc777.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc777 *Erc777TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc777.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_Erc777 *Erc777Caller) Allowance(opts *bind.CallOpts, holder common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc777.contract.Call(opts, &out, "allowance", holder, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_Erc777 *Erc777Session) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _Erc777.Contract.Allowance(&_Erc777.CallOpts, holder, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_Erc777 *Erc777CallerSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _Erc777.Contract.Allowance(&_Erc777.CallOpts, holder, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenHolder) view returns(uint256)
func (_Erc777 *Erc777Caller) BalanceOf(opts *bind.CallOpts, tokenHolder common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc777.contract.Call(opts, &out, "balanceOf", tokenHolder)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenHolder) view returns(uint256)
func (_Erc777 *Erc777Session) BalanceOf(tokenHolder common.Address) (*big.Int, error) {
	return _Erc777.Contract.BalanceOf(&_Erc777.CallOpts, tokenHolder)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenHolder) view returns(uint256)
func (_Erc777 *Erc777CallerSession) BalanceOf(tokenHolder common.Address) (*big.Int, error) {
	return _Erc777.Contract.BalanceOf(&_Erc777.CallOpts, tokenHolder)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Erc777 *Erc777Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Erc777.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Erc777 *Erc777Session) Decimals() (uint8, error) {
	return _Erc777.Contract.Decimals(&_Erc777.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Erc777 *Erc777CallerSession) Decimals() (uint8, error) {
	return _Erc777.Contract.Decimals(&_Erc777.CallOpts)
}

// DefaultOperators is a free data retrieval call binding the contract method 0x06e48538.
//
// Solidity: function defaultOperators() view returns(address[])
func (_Erc777 *Erc777Caller) DefaultOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Erc777.contract.Call(opts, &out, "defaultOperators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// DefaultOperators is a free data retrieval call binding the contract method 0x06e48538.
//
// Solidity: function defaultOperators() view returns(address[])
func (_Erc777 *Erc777Session) DefaultOperators() ([]common.Address, error) {
	return _Erc777.Contract.DefaultOperators(&_Erc777.CallOpts)
}

// DefaultOperators is a free data retrieval call binding the contract method 0x06e48538.
//
// Solidity: function defaultOperators() view returns(address[])
func (_Erc777 *Erc777CallerSession) DefaultOperators() ([]common.Address, error) {
	return _Erc777.Contract.DefaultOperators(&_Erc777.CallOpts)
}

// Granularity is a free data retrieval call binding the contract method 0x556f0dc7.
//
// Solidity: function granularity() view returns(uint256)
func (_Erc777 *Erc777Caller) Granularity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc777.contract.Call(opts, &out, "granularity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Granularity is a free data retrieval call binding the contract method 0x556f0dc7.
//
// Solidity: function granularity() view returns(uint256)
func (_Erc777 *Erc777Session) Granularity() (*big.Int, error) {
	return _Erc777.Contract.Granularity(&_Erc777.CallOpts)
}

// Granularity is a free data retrieval call binding the contract method 0x556f0dc7.
//
// Solidity: function granularity() view returns(uint256)
func (_Erc777 *Erc777CallerSession) Granularity() (*big.Int, error) {
	return _Erc777.Contract.Granularity(&_Erc777.CallOpts)
}

// IsOperatorFor is a free data retrieval call binding the contract method 0xd95b6371.
//
// Solidity: function isOperatorFor(address operator, address tokenHolder) view returns(bool)
func (_Erc777 *Erc777Caller) IsOperatorFor(opts *bind.CallOpts, operator common.Address, tokenHolder common.Address) (bool, error) {
	var out []interface{}
	err := _Erc777.contract.Call(opts, &out, "isOperatorFor", operator, tokenHolder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorFor is a free data retrieval call binding the contract method 0xd95b6371.
//
// Solidity: function isOperatorFor(address operator, address tokenHolder) view returns(bool)
func (_Erc777 *Erc777Session) IsOperatorFor(operator common.Address, tokenHolder common.Address) (bool, error) {
	return _Erc777.Contract.IsOperatorFor(&_Erc777.CallOpts, operator, tokenHolder)
}

// IsOperatorFor is a free data retrieval call binding the contract method 0xd95b6371.
//
// Solidity: function isOperatorFor(address operator, address tokenHolder) view returns(bool)
func (_Erc777 *Erc777CallerSession) IsOperatorFor(operator common.Address, tokenHolder common.Address) (bool, error) {
	return _Erc777.Contract.IsOperatorFor(&_Erc777.CallOpts, operator, tokenHolder)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc777 *Erc777Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc777.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc777 *Erc777Session) Name() (string, error) {
	return _Erc777.Contract.Name(&_Erc777.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc777 *Erc777CallerSession) Name() (string, error) {
	return _Erc777.Contract.Name(&_Erc777.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc777 *Erc777Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc777.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc777 *Erc777Session) Symbol() (string, error) {
	return _Erc777.Contract.Symbol(&_Erc777.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc777 *Erc777CallerSession) Symbol() (string, error) {
	return _Erc777.Contract.Symbol(&_Erc777.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc777 *Erc777Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc777.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc777 *Erc777Session) TotalSupply() (*big.Int, error) {
	return _Erc777.Contract.TotalSupply(&_Erc777.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc777 *Erc777CallerSession) TotalSupply() (*big.Int, error) {
	return _Erc777.Contract.TotalSupply(&_Erc777.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Erc777 *Erc777Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc777.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Erc777 *Erc777Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc777.Contract.Approve(&_Erc777.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Erc777 *Erc777TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc777.Contract.Approve(&_Erc777.TransactOpts, spender, value)
}

// AuthorizeOperator is a paid mutator transaction binding the contract method 0x959b8c3f.
//
// Solidity: function authorizeOperator(address operator) returns()
func (_Erc777 *Erc777Transactor) AuthorizeOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Erc777.contract.Transact(opts, "authorizeOperator", operator)
}

// AuthorizeOperator is a paid mutator transaction binding the contract method 0x959b8c3f.
//
// Solidity: function authorizeOperator(address operator) returns()
func (_Erc777 *Erc777Session) AuthorizeOperator(operator common.Address) (*types.Transaction, error) {
	return _Erc777.Contract.AuthorizeOperator(&_Erc777.TransactOpts, operator)
}

// AuthorizeOperator is a paid mutator transaction binding the contract method 0x959b8c3f.
//
// Solidity: function authorizeOperator(address operator) returns()
func (_Erc777 *Erc777TransactorSession) AuthorizeOperator(operator common.Address) (*types.Transaction, error) {
	return _Erc777.Contract.AuthorizeOperator(&_Erc777.TransactOpts, operator)
}

// Burn is a paid mutator transaction binding the contract method 0xfe9d9303.
//
// Solidity: function burn(uint256 amount, bytes data) returns()
func (_Erc777 *Erc777Transactor) Burn(opts *bind.TransactOpts, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc777.contract.Transact(opts, "burn", amount, data)
}

// Burn is a paid mutator transaction binding the contract method 0xfe9d9303.
//
// Solidity: function burn(uint256 amount, bytes data) returns()
func (_Erc777 *Erc777Session) Burn(amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc777.Contract.Burn(&_Erc777.TransactOpts, amount, data)
}

// Burn is a paid mutator transaction binding the contract method 0xfe9d9303.
//
// Solidity: function burn(uint256 amount, bytes data) returns()
func (_Erc777 *Erc777TransactorSession) Burn(amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc777.Contract.Burn(&_Erc777.TransactOpts, amount, data)
}

// OperatorBurn is a paid mutator transaction binding the contract method 0xfc673c4f.
//
// Solidity: function operatorBurn(address account, uint256 amount, bytes data, bytes operatorData) returns()
func (_Erc777 *Erc777Transactor) OperatorBurn(opts *bind.TransactOpts, account common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _Erc777.contract.Transact(opts, "operatorBurn", account, amount, data, operatorData)
}

// OperatorBurn is a paid mutator transaction binding the contract method 0xfc673c4f.
//
// Solidity: function operatorBurn(address account, uint256 amount, bytes data, bytes operatorData) returns()
func (_Erc777 *Erc777Session) OperatorBurn(account common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _Erc777.Contract.OperatorBurn(&_Erc777.TransactOpts, account, amount, data, operatorData)
}

// OperatorBurn is a paid mutator transaction binding the contract method 0xfc673c4f.
//
// Solidity: function operatorBurn(address account, uint256 amount, bytes data, bytes operatorData) returns()
func (_Erc777 *Erc777TransactorSession) OperatorBurn(account common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _Erc777.Contract.OperatorBurn(&_Erc777.TransactOpts, account, amount, data, operatorData)
}

// OperatorSend is a paid mutator transaction binding the contract method 0x62ad1b83.
//
// Solidity: function operatorSend(address sender, address recipient, uint256 amount, bytes data, bytes operatorData) returns()
func (_Erc777 *Erc777Transactor) OperatorSend(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _Erc777.contract.Transact(opts, "operatorSend", sender, recipient, amount, data, operatorData)
}

// OperatorSend is a paid mutator transaction binding the contract method 0x62ad1b83.
//
// Solidity: function operatorSend(address sender, address recipient, uint256 amount, bytes data, bytes operatorData) returns()
func (_Erc777 *Erc777Session) OperatorSend(sender common.Address, recipient common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _Erc777.Contract.OperatorSend(&_Erc777.TransactOpts, sender, recipient, amount, data, operatorData)
}

// OperatorSend is a paid mutator transaction binding the contract method 0x62ad1b83.
//
// Solidity: function operatorSend(address sender, address recipient, uint256 amount, bytes data, bytes operatorData) returns()
func (_Erc777 *Erc777TransactorSession) OperatorSend(sender common.Address, recipient common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _Erc777.Contract.OperatorSend(&_Erc777.TransactOpts, sender, recipient, amount, data, operatorData)
}

// RevokeOperator is a paid mutator transaction binding the contract method 0xfad8b32a.
//
// Solidity: function revokeOperator(address operator) returns()
func (_Erc777 *Erc777Transactor) RevokeOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Erc777.contract.Transact(opts, "revokeOperator", operator)
}

// RevokeOperator is a paid mutator transaction binding the contract method 0xfad8b32a.
//
// Solidity: function revokeOperator(address operator) returns()
func (_Erc777 *Erc777Session) RevokeOperator(operator common.Address) (*types.Transaction, error) {
	return _Erc777.Contract.RevokeOperator(&_Erc777.TransactOpts, operator)
}

// RevokeOperator is a paid mutator transaction binding the contract method 0xfad8b32a.
//
// Solidity: function revokeOperator(address operator) returns()
func (_Erc777 *Erc777TransactorSession) RevokeOperator(operator common.Address) (*types.Transaction, error) {
	return _Erc777.Contract.RevokeOperator(&_Erc777.TransactOpts, operator)
}

// Send is a paid mutator transaction binding the contract method 0x9bd9bbc6.
//
// Solidity: function send(address recipient, uint256 amount, bytes data) returns()
func (_Erc777 *Erc777Transactor) Send(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc777.contract.Transact(opts, "send", recipient, amount, data)
}

// Send is a paid mutator transaction binding the contract method 0x9bd9bbc6.
//
// Solidity: function send(address recipient, uint256 amount, bytes data) returns()
func (_Erc777 *Erc777Session) Send(recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc777.Contract.Send(&_Erc777.TransactOpts, recipient, amount, data)
}

// Send is a paid mutator transaction binding the contract method 0x9bd9bbc6.
//
// Solidity: function send(address recipient, uint256 amount, bytes data) returns()
func (_Erc777 *Erc777TransactorSession) Send(recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc777.Contract.Send(&_Erc777.TransactOpts, recipient, amount, data)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Erc777 *Erc777Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc777.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Erc777 *Erc777Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc777.Contract.Transfer(&_Erc777.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Erc777 *Erc777TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc777.Contract.Transfer(&_Erc777.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address holder, address recipient, uint256 amount) returns(bool)
func (_Erc777 *Erc777Transactor) TransferFrom(opts *bind.TransactOpts, holder common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc777.contract.Transact(opts, "transferFrom", holder, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address holder, address recipient, uint256 amount) returns(bool)
func (_Erc777 *Erc777Session) TransferFrom(holder common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc777.Contract.TransferFrom(&_Erc777.TransactOpts, holder, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address holder, address recipient, uint256 amount) returns(bool)
func (_Erc777 *Erc777TransactorSession) TransferFrom(holder common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc777.Contract.TransferFrom(&_Erc777.TransactOpts, holder, recipient, amount)
}

// Erc777ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Erc777 contract.
type Erc777ApprovalIterator struct {
	Event *Erc777Approval // Event containing the contract specifics and raw log

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
func (it *Erc777ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc777Approval)
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
		it.Event = new(Erc777Approval)
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
func (it *Erc777ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc777ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc777Approval represents a Approval event raised by the Erc777 contract.
type Erc777Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc777 *Erc777Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Erc777ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc777.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Erc777ApprovalIterator{contract: _Erc777.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc777 *Erc777Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Erc777Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc777.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc777Approval)
				if err := _Erc777.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Erc777 *Erc777Filterer) ParseApproval(log types.Log) (*Erc777Approval, error) {
	event := new(Erc777Approval)
	if err := _Erc777.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc777AuthorizedOperatorIterator is returned from FilterAuthorizedOperator and is used to iterate over the raw logs and unpacked data for AuthorizedOperator events raised by the Erc777 contract.
type Erc777AuthorizedOperatorIterator struct {
	Event *Erc777AuthorizedOperator // Event containing the contract specifics and raw log

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
func (it *Erc777AuthorizedOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc777AuthorizedOperator)
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
		it.Event = new(Erc777AuthorizedOperator)
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
func (it *Erc777AuthorizedOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc777AuthorizedOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc777AuthorizedOperator represents a AuthorizedOperator event raised by the Erc777 contract.
type Erc777AuthorizedOperator struct {
	Operator    common.Address
	TokenHolder common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAuthorizedOperator is a free log retrieval operation binding the contract event 0xf4caeb2d6ca8932a215a353d0703c326ec2d81fc68170f320eb2ab49e9df61f9.
//
// Solidity: event AuthorizedOperator(address indexed operator, address indexed tokenHolder)
func (_Erc777 *Erc777Filterer) FilterAuthorizedOperator(opts *bind.FilterOpts, operator []common.Address, tokenHolder []common.Address) (*Erc777AuthorizedOperatorIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenHolderRule []interface{}
	for _, tokenHolderItem := range tokenHolder {
		tokenHolderRule = append(tokenHolderRule, tokenHolderItem)
	}

	logs, sub, err := _Erc777.contract.FilterLogs(opts, "AuthorizedOperator", operatorRule, tokenHolderRule)
	if err != nil {
		return nil, err
	}
	return &Erc777AuthorizedOperatorIterator{contract: _Erc777.contract, event: "AuthorizedOperator", logs: logs, sub: sub}, nil
}

// WatchAuthorizedOperator is a free log subscription operation binding the contract event 0xf4caeb2d6ca8932a215a353d0703c326ec2d81fc68170f320eb2ab49e9df61f9.
//
// Solidity: event AuthorizedOperator(address indexed operator, address indexed tokenHolder)
func (_Erc777 *Erc777Filterer) WatchAuthorizedOperator(opts *bind.WatchOpts, sink chan<- *Erc777AuthorizedOperator, operator []common.Address, tokenHolder []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenHolderRule []interface{}
	for _, tokenHolderItem := range tokenHolder {
		tokenHolderRule = append(tokenHolderRule, tokenHolderItem)
	}

	logs, sub, err := _Erc777.contract.WatchLogs(opts, "AuthorizedOperator", operatorRule, tokenHolderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc777AuthorizedOperator)
				if err := _Erc777.contract.UnpackLog(event, "AuthorizedOperator", log); err != nil {
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

// ParseAuthorizedOperator is a log parse operation binding the contract event 0xf4caeb2d6ca8932a215a353d0703c326ec2d81fc68170f320eb2ab49e9df61f9.
//
// Solidity: event AuthorizedOperator(address indexed operator, address indexed tokenHolder)
func (_Erc777 *Erc777Filterer) ParseAuthorizedOperator(log types.Log) (*Erc777AuthorizedOperator, error) {
	event := new(Erc777AuthorizedOperator)
	if err := _Erc777.contract.UnpackLog(event, "AuthorizedOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc777BurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the Erc777 contract.
type Erc777BurnedIterator struct {
	Event *Erc777Burned // Event containing the contract specifics and raw log

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
func (it *Erc777BurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc777Burned)
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
		it.Event = new(Erc777Burned)
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
func (it *Erc777BurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc777BurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc777Burned represents a Burned event raised by the Erc777 contract.
type Erc777Burned struct {
	Operator     common.Address
	From         common.Address
	Amount       *big.Int
	Data         []byte
	OperatorData []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBurned is a free log retrieval operation binding the contract event 0xa78a9be3a7b862d26933ad85fb11d80ef66b8f972d7cbba06621d583943a4098.
//
// Solidity: event Burned(address indexed operator, address indexed from, uint256 amount, bytes data, bytes operatorData)
func (_Erc777 *Erc777Filterer) FilterBurned(opts *bind.FilterOpts, operator []common.Address, from []common.Address) (*Erc777BurnedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Erc777.contract.FilterLogs(opts, "Burned", operatorRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &Erc777BurnedIterator{contract: _Erc777.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0xa78a9be3a7b862d26933ad85fb11d80ef66b8f972d7cbba06621d583943a4098.
//
// Solidity: event Burned(address indexed operator, address indexed from, uint256 amount, bytes data, bytes operatorData)
func (_Erc777 *Erc777Filterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *Erc777Burned, operator []common.Address, from []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Erc777.contract.WatchLogs(opts, "Burned", operatorRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc777Burned)
				if err := _Erc777.contract.UnpackLog(event, "Burned", log); err != nil {
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

// ParseBurned is a log parse operation binding the contract event 0xa78a9be3a7b862d26933ad85fb11d80ef66b8f972d7cbba06621d583943a4098.
//
// Solidity: event Burned(address indexed operator, address indexed from, uint256 amount, bytes data, bytes operatorData)
func (_Erc777 *Erc777Filterer) ParseBurned(log types.Log) (*Erc777Burned, error) {
	event := new(Erc777Burned)
	if err := _Erc777.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc777MintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the Erc777 contract.
type Erc777MintedIterator struct {
	Event *Erc777Minted // Event containing the contract specifics and raw log

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
func (it *Erc777MintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc777Minted)
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
		it.Event = new(Erc777Minted)
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
func (it *Erc777MintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc777MintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc777Minted represents a Minted event raised by the Erc777 contract.
type Erc777Minted struct {
	Operator     common.Address
	To           common.Address
	Amount       *big.Int
	Data         []byte
	OperatorData []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x2fe5be0146f74c5bce36c0b80911af6c7d86ff27e89d5cfa61fc681327954e5d.
//
// Solidity: event Minted(address indexed operator, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_Erc777 *Erc777Filterer) FilterMinted(opts *bind.FilterOpts, operator []common.Address, to []common.Address) (*Erc777MintedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc777.contract.FilterLogs(opts, "Minted", operatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc777MintedIterator{contract: _Erc777.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x2fe5be0146f74c5bce36c0b80911af6c7d86ff27e89d5cfa61fc681327954e5d.
//
// Solidity: event Minted(address indexed operator, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_Erc777 *Erc777Filterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *Erc777Minted, operator []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc777.contract.WatchLogs(opts, "Minted", operatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc777Minted)
				if err := _Erc777.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0x2fe5be0146f74c5bce36c0b80911af6c7d86ff27e89d5cfa61fc681327954e5d.
//
// Solidity: event Minted(address indexed operator, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_Erc777 *Erc777Filterer) ParseMinted(log types.Log) (*Erc777Minted, error) {
	event := new(Erc777Minted)
	if err := _Erc777.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc777RevokedOperatorIterator is returned from FilterRevokedOperator and is used to iterate over the raw logs and unpacked data for RevokedOperator events raised by the Erc777 contract.
type Erc777RevokedOperatorIterator struct {
	Event *Erc777RevokedOperator // Event containing the contract specifics and raw log

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
func (it *Erc777RevokedOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc777RevokedOperator)
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
		it.Event = new(Erc777RevokedOperator)
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
func (it *Erc777RevokedOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc777RevokedOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc777RevokedOperator represents a RevokedOperator event raised by the Erc777 contract.
type Erc777RevokedOperator struct {
	Operator    common.Address
	TokenHolder common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRevokedOperator is a free log retrieval operation binding the contract event 0x50546e66e5f44d728365dc3908c63bc5cfeeab470722c1677e3073a6ac294aa1.
//
// Solidity: event RevokedOperator(address indexed operator, address indexed tokenHolder)
func (_Erc777 *Erc777Filterer) FilterRevokedOperator(opts *bind.FilterOpts, operator []common.Address, tokenHolder []common.Address) (*Erc777RevokedOperatorIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenHolderRule []interface{}
	for _, tokenHolderItem := range tokenHolder {
		tokenHolderRule = append(tokenHolderRule, tokenHolderItem)
	}

	logs, sub, err := _Erc777.contract.FilterLogs(opts, "RevokedOperator", operatorRule, tokenHolderRule)
	if err != nil {
		return nil, err
	}
	return &Erc777RevokedOperatorIterator{contract: _Erc777.contract, event: "RevokedOperator", logs: logs, sub: sub}, nil
}

// WatchRevokedOperator is a free log subscription operation binding the contract event 0x50546e66e5f44d728365dc3908c63bc5cfeeab470722c1677e3073a6ac294aa1.
//
// Solidity: event RevokedOperator(address indexed operator, address indexed tokenHolder)
func (_Erc777 *Erc777Filterer) WatchRevokedOperator(opts *bind.WatchOpts, sink chan<- *Erc777RevokedOperator, operator []common.Address, tokenHolder []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenHolderRule []interface{}
	for _, tokenHolderItem := range tokenHolder {
		tokenHolderRule = append(tokenHolderRule, tokenHolderItem)
	}

	logs, sub, err := _Erc777.contract.WatchLogs(opts, "RevokedOperator", operatorRule, tokenHolderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc777RevokedOperator)
				if err := _Erc777.contract.UnpackLog(event, "RevokedOperator", log); err != nil {
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

// ParseRevokedOperator is a log parse operation binding the contract event 0x50546e66e5f44d728365dc3908c63bc5cfeeab470722c1677e3073a6ac294aa1.
//
// Solidity: event RevokedOperator(address indexed operator, address indexed tokenHolder)
func (_Erc777 *Erc777Filterer) ParseRevokedOperator(log types.Log) (*Erc777RevokedOperator, error) {
	event := new(Erc777RevokedOperator)
	if err := _Erc777.contract.UnpackLog(event, "RevokedOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc777SentIterator is returned from FilterSent and is used to iterate over the raw logs and unpacked data for Sent events raised by the Erc777 contract.
type Erc777SentIterator struct {
	Event *Erc777Sent // Event containing the contract specifics and raw log

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
func (it *Erc777SentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc777Sent)
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
		it.Event = new(Erc777Sent)
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
func (it *Erc777SentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc777SentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc777Sent represents a Sent event raised by the Erc777 contract.
type Erc777Sent struct {
	Operator     common.Address
	From         common.Address
	To           common.Address
	Amount       *big.Int
	Data         []byte
	OperatorData []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSent is a free log retrieval operation binding the contract event 0x06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc82614677987.
//
// Solidity: event Sent(address indexed operator, address indexed from, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_Erc777 *Erc777Filterer) FilterSent(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*Erc777SentIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc777.contract.FilterLogs(opts, "Sent", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc777SentIterator{contract: _Erc777.contract, event: "Sent", logs: logs, sub: sub}, nil
}

// WatchSent is a free log subscription operation binding the contract event 0x06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc82614677987.
//
// Solidity: event Sent(address indexed operator, address indexed from, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_Erc777 *Erc777Filterer) WatchSent(opts *bind.WatchOpts, sink chan<- *Erc777Sent, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc777.contract.WatchLogs(opts, "Sent", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc777Sent)
				if err := _Erc777.contract.UnpackLog(event, "Sent", log); err != nil {
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

// ParseSent is a log parse operation binding the contract event 0x06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc82614677987.
//
// Solidity: event Sent(address indexed operator, address indexed from, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_Erc777 *Erc777Filterer) ParseSent(log types.Log) (*Erc777Sent, error) {
	event := new(Erc777Sent)
	if err := _Erc777.contract.UnpackLog(event, "Sent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc777TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc777 contract.
type Erc777TransferIterator struct {
	Event *Erc777Transfer // Event containing the contract specifics and raw log

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
func (it *Erc777TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc777Transfer)
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
		it.Event = new(Erc777Transfer)
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
func (it *Erc777TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc777TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc777Transfer represents a Transfer event raised by the Erc777 contract.
type Erc777Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc777 *Erc777Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Erc777TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc777.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc777TransferIterator{contract: _Erc777.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc777 *Erc777Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc777Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc777.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc777Transfer)
				if err := _Erc777.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Erc777 *Erc777Filterer) ParseTransfer(log types.Log) (*Erc777Transfer, error) {
	event := new(Erc777Transfer)
	if err := _Erc777.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
