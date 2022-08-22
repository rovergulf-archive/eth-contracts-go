package interfaces

import "github.com/ethereum/go-ethereum/common/hexutil"

// InterfaceId represents smart contracts interface id type
type InterfaceId [4]byte

// Hex returns hex value of the interface id
func (i *InterfaceId) Hex() string {
	return hexutil.Encode(i[:])
}

// IID returns interface id [4]byte value
func (i *InterfaceId) IID() [4]byte {
	return [4]byte{i[0], i[1], i[2], i[3]}
}

/*
 * known values
 */

var (
	NIL              InterfaceId = [4]byte{0, 0, 0, 0}         // 0x00000000
	ERC20            InterfaceId = [4]byte{54, 53, 43, 7}      // 0x36372b07
	ERC165           InterfaceId = [4]byte{1, 255, 201, 167}   // 0x01ffc9a7
	ERC173           InterfaceId = [4]byte{127, 88, 40, 208}   // 0x7f5828d0
	ERC721           InterfaceId = [4]byte{128, 172, 88, 205}  // 0x80ac58cd
	ERC721Metadata   InterfaceId = [4]byte{91, 94, 19, 159}    // 0x5b5e139f
	ERC721Enumerable InterfaceId = [4]byte{120, 14, 157, 99}   // 0x780e9d63
	ERC777           InterfaceId = [4]byte{229, 142, 17, 60}   // 0xe58e113c
	ERC1155          InterfaceId = [4]byte{217, 182, 122, 38}  // 0xd9b67a26
	ERC1820          InterfaceId = [4]byte{1, 255, 201, 167}   // 0x01ffc9a7
	ERC4626          InterfaceId = [4]byte{135, 223, 229, 160} // 0x87dfe5a0
	ERC4907          InterfaceId = [4]byte{173, 9, 43, 92}     // 0xad092b5c
)

// literal ids
var (
	NilIID         = "0x00000000"
	TotalSupplyIID = "0x18160ddd"
	NameIID        = "0x06fdde03"
	SymbolIID      = "0x95d89b41"
	BalanceOfIID   = "0x70a08231"
)
