// introspection
local erc165 = import "../../v/eth-contracts/artifacts/@openzeppelin/contracts/utils/introspection/ERC165.sol/ERC165.json";
local erc1820Registry = import "../../v/eth-contracts/artifacts/@openzeppelin/contracts/utils/introspection/IERC1820Registry.sol/IERC1820Registry.json";

// token
local erc4626 = import "../../v/eth-contracts/artifacts/@openzeppelin/contracts/interfaces/IERC4626.sol/IERC4626.json";
local erc20 = import "../../v/eth-contracts/artifacts/@openzeppelin/contracts/token/ERC20/ERC20.sol/ERC20.json";
local erc721 = import "../../v/eth-contracts/artifacts/@openzeppelin/contracts/token/ERC721/ERC721.sol/ERC721.json";
local erc721Enumerable = import "../../v/eth-contracts/artifacts/@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol/ERC721Enumerable.json";
local erc721Metadata = import "../../v/eth-contracts/artifacts/@openzeppelin/contracts/token/ERC721/extensions/IERC721Metadata.sol/IERC721Metadata.json";
local erc777 = import "../../v/eth-contracts/artifacts/@openzeppelin/contracts/token/ERC777/ERC777.sol/ERC777.json";
local erc1155 = import "../../v/eth-contracts/artifacts/@openzeppelin/contracts/token/ERC1155/ERC1155.sol/ERC1155.json";
local erc1155Supply = import "../../v/eth-contracts/artifacts/@openzeppelin/contracts/token/ERC1155/extensions/ERC1155Supply.sol/ERC1155Supply.json";
local erc1155MetadataURI = import "../../v/eth-contracts/artifacts/@openzeppelin/contracts/token/ERC1155/extensions/IERC1155MetadataURI.sol/IERC1155MetadataURI.json";

// ownable and its related interfaces
local ownable = import "../../v/eth-contracts/artifacts/@openzeppelin/contracts/access/Ownable.sol/Ownable.json";
local erc173 = import "../../v/eth-contracts/artifacts/contracts/interfaces/IERC173.sol/IERC173.json";
local erc5313 = import "../../v/eth-contracts/artifacts/contracts/interfaces/IERC5313.sol/IERC5313.json";

// experimental
local coin = import "../../v/eth-contracts/artifacts/contracts/experimental/coin/Coin.sol/RovergulfCoin.json";
local coinPool = import "../../v/eth-contracts/artifacts/contracts/experimental/coin/Pool.sol/RCPoolManager.json";
local coinStake = import "../../v/eth-contracts/artifacts/contracts/experimental/coin/Stake.sol/RCStake.json";
local coinVault = import "../../v/eth-contracts/artifacts/contracts/experimental/coin/Vault.sol/RCVault.json";

local treasury = import "../../v/eth-contracts/artifacts/contracts/experimental/Treasury.sol/Treasury.json";


{
  erc4626: erc4626.abi,
  erc165: erc165.abi,
  erc1820Registry: erc1820Registry.abi,
  erc20: erc20.abi,
  erc721: erc721.abi,
  erc721Enumerable: erc721Enumerable.abi,
  erc721Metadata: erc721Metadata.abi,
  erc777: erc777.abi,
  erc1155: erc1155.abi,
  erc1155MetadataURI: erc1155MetadataURI.abi,
  erc1155Supply: erc1155Supply.abi,
  ownable: ownable.abi,
  erc173: erc173.abi,
  erc5313: erc5313.abi,
  coin: coin.abi,
  coinPool: coinPool.abi,
  coinStake: coinStake.abi,
  coinVault: coinVault.abi,
  treasury: treasury.abi,
}