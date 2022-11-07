local erc20 = import "../../v/eth-contracts/artifacts/@openzeppelin/contracts/token/ERC20/ERC20.sol/ERC20.json";
local erc721 = import "../../v/eth-contracts/artifacts/@openzeppelin/contracts/token/ERC721/ERC721.sol/ERC721.json";
local erc777 = import "../../v/eth-contracts/artifacts/@openzeppelin/contracts/token/ERC777/ERC777.sol/ERC777.json";
local erc1155 = import "../../v/eth-contracts/artifacts/@openzeppelin/contracts/token/ERC1155/ERC1155.sol/ERC1155.json";

local coin = import "../../v/eth-contracts/artifacts/contracts/experimental/coin/Coin.sol/RovergulfCoin.json";
local coinPool = import "../../v/eth-contracts/artifacts/contracts/experimental/coin/Pool.sol/RCPoolManager.json";
local coinStake = import "../../v/eth-contracts/artifacts/contracts/experimental/coin/Stake.sol/RCStake.json";
local coinVault = import "../../v/eth-contracts/artifacts/contracts/experimental/coin/Vault.sol/RCVault.json";

local treasury = import "../../v/eth-contracts/artifacts/contracts/experimental/Treasury.sol/Treasury.json";

{
  erc20: erc20.bytecode,
  erc721: erc721.bytecode,
  erc777: erc777.bytecode,
  erc1155: erc1155.bytecode,
  coin: coin.bytecode,
  coinPool: coinPool.bytecode,
  coinStake: coinStake.bytecode,
  coinVault: coinVault.bytecode,
  treasury: treasury.bytecode,
}