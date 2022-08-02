local erc165 = import "../v/eth-contracts/artifacts/@openzeppelin/contracts/utils/introspection/ERC165.sol/ERC165.json";
local erc20 = import "../v/eth-contracts/artifacts/@openzeppelin/contracts/token/ERC20/ERC20.sol/ERC20.json";
local erc721 = import "../v/eth-contracts/artifacts/@openzeppelin/contracts/token/ERC721/ERC721.sol/ERC721.json";
local erc777 = import "../v/eth-contracts/artifacts/@openzeppelin/contracts/token/ERC777/ERC777.sol/ERC777.json";
local erc1155 = import "../v/eth-contracts/artifacts/@openzeppelin/contracts/token/ERC1155/ERC1155.sol/ERC1155.json";
{
  erc165: erc165.abi,
  erc20: erc20.abi,
  erc721: erc721.abi,
  erc777: erc777.abi,
  erc1155: erc1155.abi,
}