local erc20 = import "../../v/eth-contracts/artifacts/@openzeppelin/contracts/token/ERC20/ERC20.sol/ERC20.json";
local erc721 = import "../../v/eth-contracts/artifacts/@openzeppelin/contracts/token/ERC721/ERC721.sol/ERC721.json";
local erc777 = import "../../v/eth-contracts/artifacts/@openzeppelin/contracts/token/ERC777/ERC777.sol/ERC777.json";
local erc1155 = import "../../v/eth-contracts/artifacts/@openzeppelin/contracts/token/ERC1155/ERC1155.sol/ERC1155.json";

{
  erc20: erc20.bytecode,
  erc721: erc721.bytecode,
  erc777: erc777.bytecode,
  erc1155: erc1155.bytecode,
}