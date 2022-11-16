package common

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/polynetwork/bridge-common/base"
	"github.com/polynetwork/bridge-common/chains/aptos"
	"github.com/polynetwork/bridge-common/chains/eth"
	"github.com/polynetwork/bridge-common/chains/neo"
	"github.com/polynetwork/bridge-common/chains/neo3"
	"github.com/polynetwork/bridge-common/chains/ont"
	"github.com/polynetwork/bridge-common/chains/starcoin"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	erc20 "poly-bridge/go_abi/mintable_erc20_abi"
	"strings"
	"time"
)

var (
	sdkMap map[uint64]interface{}
)

func GetSdk(chainId uint64) interface{} {
	return sdkMap[chainId]
}

func SetupChainsSDK(cfg *conf.Config) {
	if cfg == nil {
		panic("Missing config")
	}

	sdkMap = make(map[uint64]interface{}, 0)
	for _, config := range cfg.ChainListenConfig {
		switch config.ChainId {
		case base.ETH, base.PLT, base.BSC, base.HECO, base.OK, base.MATIC, base.ARBITRUM, base.XDAI, base.FANTOM,
			base.AVA, base.OPTIMISM, base.METIS, base.BOBA, base.RINKEBY, base.BYTOM, base.OASIS, base.HARMONY,
			base.KCC, base.HSC, base.KAVA, base.CUBE, base.ZKSYNC, base.CELO, base.CLOVER, base.CONFLUX, base.ASTAR,
			base.BRISE:
			sdk, err := eth.WithOptions(config.ChainId, config.Nodes, time.Minute, 1)
			if err != nil {
				panic(fmt.Sprintf("Create chain sdk failed. chain=%d, err=%s", config.ChainId, err))
			}
			sdkMap[config.ChainId] = sdk

		case base.NEO:
			sdk, err := neo.WithOptions(config.ChainId, config.Nodes, time.Minute, 1)
			if err != nil {
				panic(fmt.Sprintf("Create chain sdk failed. chain=%d, err=%s", config.ChainId, err))
			}
			sdkMap[config.ChainId] = sdk
		case base.NEO3:
			sdk, err := neo3.WithOptions(config.ChainId, config.Nodes, time.Minute, 1)
			if err != nil {
				panic(fmt.Sprintf("Create chain sdk failed. chain=%d, err=%s", config.ChainId, err))
			}
			sdkMap[config.ChainId] = sdk
		case base.ONT:
			sdk, err := ont.WithOptions(config.ChainId, config.Nodes, time.Minute, 1)
			if err != nil {
				panic(fmt.Sprintf("Create chain sdk failed. chain=%d, err=%s", config.ChainId, err))
			}
			sdkMap[config.ChainId] = sdk
		//case base.SWITCHEO: // todo to be added
		//case base.ZILLIQA: // todo to be added
		case base.STARCOIN:
			sdk, err := starcoin.WithOptions(config.ChainId, config.Nodes, time.Minute, 1)
			if err != nil {
				panic(fmt.Sprintf("Create chain sdk failed. chain=%d, err=%s", config.ChainId, err))
			}
			sdkMap[config.ChainId] = sdk
		//case base.RIPPLE: // todo to be added
		case base.APTOS:
			sdk, err := aptos.WithOptions(config.ChainId, config.Nodes, time.Minute, 1)
			if err != nil {
				panic(fmt.Sprintf("Create chain sdk failed. chain=%d, err=%s", config.ChainId, err))
			}
			sdkMap[config.ChainId] = sdk
		}
	}
}

func getEthErc20Balance(token, owner string, client *eth.Client) (balance *big.Int, err error) {
	tokenAddress := common.HexToAddress(token)
	ownerAddr := common.HexToAddress(owner)
	if basedef.IsNativeTokenAddress(token) {
		var result hexutil.Big
		err = client.Rpc.CallContext(context.Background(), &result, "eth_getBalance", "0x"+owner, "latest")
		balance = (*big.Int)(&result)
	} else {
		var contract *erc20.ERC20Extended
		contract, err = erc20.NewERC20Mintable(tokenAddress, client)
		if err == nil {
			balance, err = contract.BalanceOf(nil, ownerAddr)
		}
	}
	return
}

func GetBalance(chainId uint64, hash string) (balance *big.Int, err error) {
	maxBalance, balance := big.NewInt(0), big.NewInt(0)
	maxFun := func(balance *big.Int) {
		if balance != nil && balance.Cmp(maxBalance) > 0 {
			maxBalance = balance
		}
	}

	errMap := make(map[error]interface{}, 0)
	chainConfig := conf.GlobalConfig.GetChainListenConfig(chainId)
	if chainConfig == nil {
		err = fmt.Errorf("chain %d is invalid", chainId)
		return
	}
	switch chainId {
	case base.ETH, base.GOERLI, base.MATIC, base.BSC, base.HECO, base.OK, base.ARBITRUM, base.XDAI, base.FANTOM,
		base.AVA, base.OPTIMISM, base.METIS, base.BOBA, base.RINKEBY, base.BYTOM, base.OASIS, base.HARMONY,
		base.KCC, base.HSC, base.KAVA, base.CUBE, base.ZKSYNC, base.CELO, base.CLOVER, base.CONFLUX, base.ASTAR,
		base.BRISE:
		if sdk, ok := sdkMap[chainId]; ok {
			if ethSdk, ok := sdk.(*eth.SDK); ok {
				for _, v := range chainConfig.ProxyContract {
					if len(strings.TrimSpace(v)) == 0 {
						continue
					}
					balance, err := ethSdk.Node().GetBalance(hash, v)
					if err != nil {
						return nil, err
					}
					errMap[err] = nil
					maxFun(balance)
				}
			} else {
				err = fmt.Errorf("chain %d sdk type invalid", chainId)
			}
		} else {
			err = fmt.Errorf("chain %d sdk not initialized", chainId)
		}
	case base.NEO:
		if sdk, ok := sdkMap[chainId]; ok {
			if neoSdk, ok := sdk.(*neo.SDK); ok {
				for _, v := range chainConfig.ProxyContract {
					if len(strings.TrimSpace(v)) == 0 {
						continue
					}
					balance, err := neoSdk.Node().GetBalance(hash, v)
					if err != nil {
						return nil, err
					}
					errMap[err] = nil
					maxFun(balance)
				}
			} else {
				err = fmt.Errorf("chain %d sdk type invalid", chainId)
			}
		} else {
			err = fmt.Errorf("chain %d sdk not initialized", chainId)
		}

	case base.NEO3:
		if sdk, ok := sdkMap[chainId]; ok {
			if neo3Sdk, ok := sdk.(*neo3.SDK); ok {
				for _, v := range chainConfig.ProxyContract {
					if len(strings.TrimSpace(v)) == 0 {
						continue
					}
					balance, err := neo3Sdk.Node().GetBalance(hash, v)
					errMap[err] = nil
					maxFun(balance)
				}
			} else {
				err = fmt.Errorf("chain %d sdk type invalid", chainId)
			}
		} else {
			err = fmt.Errorf("chain %d sdk not initialized", chainId)
		}

	case base.ONT:
	//case base.ZILLIQA: // todo to be added
	case base.STARCOIN:
	//case base.RIPPLE: // todo to be added
	case base.APTOS:
		if sdk, ok := sdkMap[chainId]; ok {
			if aptosSdk, ok := sdk.(*aptos.SDK); ok {
				for _, v := range chainConfig.ProxyContract {
					if len(strings.TrimSpace(v)) == 0 {
						continue
					}
					resource, err := aptosSdk.Node().GetResourceByAccountAddressAndResourceType(context.Background(), v, fmt.Sprintf("%s::lock_proxy::Treasury<%s>", "0x"+strings.TrimPrefix(v, "0x"), hash))
					errMap[err] = nil
					if err == nil && resource != nil {
						if balance, ok := new(big.Int).SetString(resource.Data.CoinStoreResource.Coin.Value, 10); ok {
							maxFun(balance)
						}
					}
				}
			} else {
				err = fmt.Errorf("chain %d sdk type invalid", chainId)
			}
		} else {
			err = fmt.Errorf("chain %d sdk not initialized", chainId)
		}

	default:
		return new(big.Int).SetUint64(0), fmt.Errorf(fmt.Sprintf("chain id %d invalid", chainId))
	}

	if maxBalance.Cmp(big.NewInt(0)) > 0 {
		return maxBalance, nil
	}
	for k, _ := range errMap {
		if k == nil {
			return new(big.Int).SetUint64(0), nil
		} else {
			err = k
		}
	}
	return new(big.Int).SetUint64(0), err
}

func GetTotalSupply(chainId uint64, hash string) (totalSupply *big.Int, err error) {
	switch chainId {
	case base.ETH, base.PLT, base.BSC, base.HECO, base.OK, base.MATIC, base.ARBITRUM, base.XDAI, base.FANTOM,
		base.AVA, base.OPTIMISM, base.METIS, base.BOBA, base.RINKEBY, base.BYTOM, base.OASIS, base.HARMONY,
		base.KCC, base.HSC, base.KAVA, base.CUBE, base.ZKSYNC, base.CELO, base.CLOVER, base.CONFLUX, base.ASTAR,
		base.BRISE:
		if sdk, ok := sdkMap[chainId]; ok {
			if ethSdk, ok := sdk.(*eth.SDK); ok {
				erc20Address := common.HexToAddress(hash)
				totalSupply = new(big.Int).SetUint64(0)
				if !basedef.IsNativeTokenAddress(hash) {
					contract, e := erc20.NewERC20Mintable(erc20Address, ethSdk.Node())
					if e == nil {
						totalSupply, err = contract.TotalSupply(nil)
					}
				}
			} else {
				err = fmt.Errorf("chain %d sdk type invalid", chainId)
			}
		} else {
			err = fmt.Errorf("chain %d sdk not initialized", chainId)
		}
		return
	default:
		return new(big.Int).SetUint64(0), fmt.Errorf(fmt.Sprintf("chain id %d invalid", chainId))
	}
}

func GetProxyBalance(chainId uint64, hash string, proxy string) (balance *big.Int, err error) {
	switch chainId {
	case base.ETH, base.PLT, base.BSC, base.HECO, base.OK, base.MATIC, base.ARBITRUM, base.XDAI, base.FANTOM,
		base.AVA, base.OPTIMISM, base.METIS, base.BOBA, base.RINKEBY, base.BYTOM, base.OASIS, base.HARMONY,
		base.KCC, base.HSC, base.KAVA, base.CUBE, base.ZKSYNC, base.CELO, base.CLOVER, base.CONFLUX, base.ASTAR,
		base.BRISE:
		if sdk, ok := sdkMap[chainId]; ok {
			if ethSdk, ok := sdk.(*eth.SDK); ok {
				return ethSdk.Node().GetBalance(hash, proxy)
			} else {
				err = fmt.Errorf("chain %d sdk type invalid", chainId)
			}
		} else {
			err = fmt.Errorf("chain %d sdk not initialized", chainId)
		}
		return
	case base.NEO:
		if sdk, ok := sdkMap[chainId]; ok {
			if neoSdk, ok := sdk.(*neo.SDK); ok {
				return neoSdk.Node().GetBalance(hash, proxy)
			} else {
				err = fmt.Errorf("chain %d sdk type invalid", chainId)
			}
		} else {
			err = fmt.Errorf("chain %d sdk not initialized", chainId)
		}
		return
	case base.NEO3:
		if sdk, ok := sdkMap[chainId]; ok {
			if neo3Sdk, ok := sdk.(*neo3.SDK); ok {
				return neo3Sdk.Node().GetBalance(hash, proxy)
			} else {
				err = fmt.Errorf("chain %d sdk type invalid", chainId)
			}
		} else {
			err = fmt.Errorf("chain %d sdk not initialized", chainId)
		}
		return
	case base.ONT:
		//case base.ZILLIQA: // todo to be added
		return
	case base.STARCOIN:
		//case base.RIPPLE: // todo to be added
		return
	case base.APTOS:
		if sdk, ok := sdkMap[chainId]; ok {
			if aptosSdk, ok := sdk.(*aptos.SDK); ok {
				resource, err := aptosSdk.Node().GetResourceByAccountAddressAndResourceType(context.Background(), proxy, fmt.Sprintf("%s::lock_proxy::Treasury<%s>", "0x"+strings.TrimPrefix(proxy, "0x"), hash))
				if err == nil && resource != nil {
					if balance, ok := new(big.Int).SetString(resource.Data.CoinStoreResource.Coin.Value, 10); ok {
						return balance, nil
					}
				}
			} else {
				err = fmt.Errorf("chain %d sdk type invalid", chainId)
			}
		} else {
			err = fmt.Errorf("chain %d sdk not initialized", chainId)
		}
		return
	default:
		err = fmt.Errorf("chain %d sdk not initialized", chainId)
		return
	}
}

func GetNftOwner(chainId uint64, asset string, tokenId int64) (owner common.Address, err error) {
	switch chainId {
	case base.ETH, base.GOERLI, base.MATIC, base.BSC, base.HECO, base.OK, base.ARBITRUM, base.XDAI:
		if sdk, ok := sdkMap[chainId]; ok {
			if ethSdk, ok := sdk.(*eth.SDK); ok {
				return ethSdk.Node().GetNFTOwner(asset, tokenId)
			} else {
				err = fmt.Errorf("chain %d sdk type invalid", chainId)
			}
		} else {
			err = fmt.Errorf("chain %d sdk not initialized", chainId)
		}
		return
	default:
		return common.Address{}, fmt.Errorf("chain %d sdk not support GetNftOwner", chainId)
	}
}

func GetBoundLockProxy(lockProxies []string, srcTokenHash, DstTokenHash string, srcChainId, dstChainId uint64) (string, error) {
	if sdk, exist := sdkMap[dstChainId]; exist {
		if value, ok := sdk.(*chainsdk.EthereumSdkPro); ok {
			return value.GetBoundLockProxy(lockProxies, srcTokenHash, DstTokenHash, srcChainId)
		}
	}
	return "", fmt.Errorf("chain %d is not ethereum based", dstChainId)
}
