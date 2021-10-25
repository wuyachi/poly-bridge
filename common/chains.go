package common

import (
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"

	"github.com/beego/beego/v2/core/logs"
)

var (
	ethereumSdk *chainsdk.EthereumSdkPro
	pltSdk      *chainsdk.EthereumSdkPro
	bscSdk      *chainsdk.EthereumSdkPro
	hecoSdk     *chainsdk.EthereumSdkPro
	okSdk       *chainsdk.EthereumSdkPro
	neoSdk      *chainsdk.NeoSdkPro
	ontologySdk *chainsdk.OntologySdkPro
	maticSdk    *chainsdk.EthereumSdkPro
	swthSdk     *chainsdk.SwitcheoSdkPro
	arbitrumSdk *chainsdk.EthereumSdkPro
	xdaiSdk     *chainsdk.EthereumSdkPro
	config      *conf.Config
)

func SetupChainsSDK(cfg *conf.Config) {
	if cfg == nil {
		panic("Missing config")
	}
	config = cfg
	newChainSdks(cfg)
}

func newChainSdks(config *conf.Config) {
	{
		ethereumConfig := config.GetChainListenConfig(basedef.ETHEREUM_CROSSCHAIN_ID)
		if ethereumConfig == nil {
			panic("chain is invalid")
		}
		urls := ethereumConfig.GetNodesUrl()
		ethereumSdk = chainsdk.NewEthereumSdkPro(urls, ethereumConfig.ListenSlot, ethereumConfig.ChainId)
	}
	{
		maticConfig := config.GetChainListenConfig(basedef.MATIC_CROSSCHAIN_ID)
		if maticConfig == nil {
			panic("chain is invalid")
		}
		urls := maticConfig.GetNodesUrl()
		maticSdk = chainsdk.NewEthereumSdkPro(urls, maticConfig.ListenSlot, maticConfig.ChainId)
	}
	{
		bscConfig := config.GetChainListenConfig(basedef.BSC_CROSSCHAIN_ID)
		if bscConfig == nil {
			panic("chain is invalid")
		}
		urls := bscConfig.GetNodesUrl()
		bscSdk = chainsdk.NewEthereumSdkPro(urls, bscConfig.ListenSlot, bscConfig.ChainId)
	}
	{
		hecoConfig := config.GetChainListenConfig(basedef.HECO_CROSSCHAIN_ID)
		if hecoConfig == nil {
			panic("chain is invalid")
		}
		urls := hecoConfig.GetNodesUrl()
		hecoSdk = chainsdk.NewEthereumSdkPro(urls, hecoConfig.ListenSlot, hecoConfig.ChainId)
	}
	{
		okConfig := config.GetChainListenConfig(basedef.OK_CROSSCHAIN_ID)
		if okConfig == nil {
			panic("chain is invalid")
		}
		urls := okConfig.GetNodesUrl()
		okSdk = chainsdk.NewEthereumSdkPro(urls, okConfig.ListenSlot, okConfig.ChainId)
	}
	{
		neoConfig := config.GetChainListenConfig(basedef.NEO_CROSSCHAIN_ID)
		if neoConfig == nil {
			panic("chain is invalid")
		}
		urls := neoConfig.GetNodesUrl()
		neoSdk = chainsdk.NewNeoSdkPro(urls, neoConfig.ListenSlot, neoConfig.ChainId)
	}
	{
		ontConfig := config.GetChainListenConfig(basedef.ONT_CROSSCHAIN_ID)
		if ontConfig == nil {
			panic("chain is invalid")
		}
		urls := ontConfig.GetNodesUrl()
		ontologySdk = chainsdk.NewOntologySdkPro(urls, ontConfig.ListenSlot, ontConfig.ChainId)
	}
	if basedef.ENV == basedef.MAINNET {
		swthConfig := config.GetChainListenConfig(basedef.SWITCHEO_CROSSCHAIN_ID)
		if swthConfig == nil {
			panic("swth chain is invalid")
		}
		urls := swthConfig.GetNodesUrl()
		swthSdk = chainsdk.NewSwitcheoSdkPro(urls, swthConfig.ListenSlot, swthConfig.ChainId)
	}
	{
		conf := config.GetChainListenConfig(basedef.PLT_CROSSCHAIN_ID)
		if conf != nil {
			urls := conf.GetNodesUrl()
			pltSdk = chainsdk.NewEthereumSdkPro(urls, conf.ListenSlot, conf.ChainId)
		} else {
			logs.Error("Missing plt chain sdk config")
		}
	}
	{
		arbitrumConfig := config.GetChainListenConfig(basedef.ARBITRUM_CROSSCHAIN_ID)
		if arbitrumConfig == nil {
			panic("chain is invalid")
		}
		urls := arbitrumConfig.GetNodesUrl()
		arbitrumSdk = chainsdk.NewEthereumSdkPro(urls, arbitrumConfig.ListenSlot, arbitrumConfig.ChainId)
	}
	{
		xdaiConfig := config.GetChainListenConfig(basedef.XDAI_CROSSCHAIN_ID)
		if xdaiConfig == nil {
			panic("chain:XDAI is invalid")
		}
		urls := xdaiConfig.GetNodesUrl()
		xdaiSdk = chainsdk.NewEthereumSdkPro(urls, xdaiConfig.ListenSlot, xdaiConfig.ChainId)
	}

}

func GetBalance(chainId uint64, hash string) (*big.Int, error) {
	if chainId == basedef.ETHEREUM_CROSSCHAIN_ID {
		ethereumConfig := config.GetChainListenConfig(basedef.ETHEREUM_CROSSCHAIN_ID)
		if ethereumConfig == nil {
			panic("chain is invalid")
		}
		return ethereumSdk.Erc20Balance(hash, ethereumConfig.ProxyContract)
	}
	if chainId == basedef.MATIC_CROSSCHAIN_ID {
		maticConfig := config.GetChainListenConfig(basedef.MATIC_CROSSCHAIN_ID)
		if maticConfig == nil {
			panic("chain is invalid")
		}
		return maticSdk.Erc20Balance(hash, maticConfig.ProxyContract)
	}
	if chainId == basedef.BSC_CROSSCHAIN_ID {
		bscConfig := config.GetChainListenConfig(basedef.BSC_CROSSCHAIN_ID)
		if bscConfig == nil {
			panic("chain is invalid")
		}
		return bscSdk.Erc20Balance(hash, bscConfig.ProxyContract)
	}
	if chainId == basedef.HECO_CROSSCHAIN_ID {
		hecoConfig := config.GetChainListenConfig(basedef.HECO_CROSSCHAIN_ID)
		if hecoConfig == nil {
			panic("chain is invalid")
		}
		return hecoSdk.Erc20Balance(hash, hecoConfig.ProxyContract)
	}
	if chainId == basedef.OK_CROSSCHAIN_ID {
		okConfig := config.GetChainListenConfig(basedef.OK_CROSSCHAIN_ID)
		if okConfig == nil {
			panic("chain is invalid")
		}
		return okSdk.Erc20Balance(hash, okConfig.ProxyContract)
	}
	if chainId == basedef.NEO_CROSSCHAIN_ID {
		neoConfig := config.GetChainListenConfig(basedef.NEO_CROSSCHAIN_ID)
		if neoConfig == nil {
			panic("chain is invalid")
		}
		return neoSdk.Nep5Balance(hash, neoConfig.ProxyContract)
	}
	if chainId == basedef.ONT_CROSSCHAIN_ID {
		ontConfig := config.GetChainListenConfig(basedef.ONT_CROSSCHAIN_ID)
		if ontConfig == nil {
			panic("chain is invalid")
		}
		return ontologySdk.Oep4Balance(hash, ontConfig.ProxyContract)
	}
	if chainId == basedef.MATIC_CROSSCHAIN_ID {
		maticConfig := config.GetChainListenConfig(basedef.MATIC_CROSSCHAIN_ID)
		if maticConfig == nil {
			panic("chain is invalid")
		}
		return maticSdk.Erc20Balance(hash, maticConfig.ProxyContract)
	}
	if chainId == basedef.ARBITRUM_CROSSCHAIN_ID {
		arbitrumConfig := config.GetChainListenConfig(basedef.ARBITRUM_CROSSCHAIN_ID)
		if arbitrumConfig == nil {
			panic("chain is invalid")
		}
		return arbitrumSdk.Erc20Balance(hash, arbitrumConfig.ProxyContract)
	}
	if chainId == basedef.XDAI_CROSSCHAIN_ID {
		xdaiConfig := config.GetChainListenConfig(basedef.XDAI_CROSSCHAIN_ID)
		if xdaiConfig == nil {
			panic("chain is invalid")
		}
		return xdaiSdk.Erc20Balance(hash, xdaiConfig.ProxyContract)
	}
	/*if chainId == basedef.PLT_CROSSCHAIN_ID {
		conf := config.GetChainListenConfig(basedef.PLT_CROSSCHAIN_ID)
		if conf == nil {
			panic("chain is invalid")
		}
		return pltSdk.Erc20Balance(hash,conf.ProxyContract)
	}
	*/
	return new(big.Int).SetUint64(0), nil
}

func GetTotalSupply(chainId uint64, hash string) (*big.Int, error) {
	if chainId == basedef.ETHEREUM_CROSSCHAIN_ID {
		ethereumConfig := config.GetChainListenConfig(basedef.ETHEREUM_CROSSCHAIN_ID)
		if ethereumConfig == nil {
			panic("chain is invalid")
		}
		return ethereumSdk.Erc20TotalSupply(hash)
	}
	if chainId == basedef.BSC_CROSSCHAIN_ID {
		bscConfig := config.GetChainListenConfig(basedef.BSC_CROSSCHAIN_ID)
		if bscConfig == nil {
			panic("chain is invalid")
		}
		return bscSdk.Erc20TotalSupply(hash)
	}
	if chainId == basedef.HECO_CROSSCHAIN_ID {
		hecoConfig := config.GetChainListenConfig(basedef.HECO_CROSSCHAIN_ID)
		if hecoConfig == nil {
			panic("chain is invalid")
		}
		return hecoSdk.Erc20TotalSupply(hash)
	}
	if chainId == basedef.OK_CROSSCHAIN_ID {
		okConfig := config.GetChainListenConfig(basedef.OK_CROSSCHAIN_ID)
		if okConfig == nil {
			panic("chain is invalid")
		}
		return okSdk.Erc20TotalSupply(hash)
	}
	if chainId == basedef.NEO_CROSSCHAIN_ID {
		neoConfig := config.GetChainListenConfig(basedef.NEO_CROSSCHAIN_ID)
		if neoConfig == nil {
			panic("chain is invalid")
		}
		return neoSdk.Nep5TotalSupply(hash)
	}
	if chainId == basedef.ONT_CROSSCHAIN_ID {
		ontConfig := config.GetChainListenConfig(basedef.ONT_CROSSCHAIN_ID)
		if ontConfig == nil {
			panic("chain is invalid")
		}
		return ontologySdk.Oep4TotalSupply(hash, ontConfig.ProxyContract)
	}
	if chainId == basedef.MATIC_CROSSCHAIN_ID {
		maticConfig := config.GetChainListenConfig(basedef.MATIC_CROSSCHAIN_ID)
		if maticConfig == nil {
			panic("chain is invalid")
		}
		return maticSdk.Erc20TotalSupply(hash)
	}
	if chainId == basedef.ARBITRUM_CROSSCHAIN_ID {
		arbitrumConfig := config.GetChainListenConfig(basedef.BSC_CROSSCHAIN_ID)
		if arbitrumConfig == nil {
			panic("chain is invalid")
		}
		return arbitrumSdk.Erc20TotalSupply(hash)
	}
	if chainId == basedef.XDAI_CROSSCHAIN_ID {
		xdaiConfig := config.GetChainListenConfig(basedef.XDAI_CROSSCHAIN_ID)
		if xdaiConfig == nil {
			panic("chain is invalid")
		}
		return xdaiSdk.Erc20TotalSupply(hash)
	}
	return new(big.Int).SetUint64(0), nil
}

type ProxyBalance struct {
	Amount    *big.Int
	ItemName  string
	ItemProxy string
}

func GetAllLockProxyBalance(chainId uint64, hash string) []*ProxyBalance {
	proxyBalances := make([]*ProxyBalance, 0)
	if chainId == basedef.ETHEREUM_CROSSCHAIN_ID {
		ethereumConfig := config.GetChainListenConfig(basedef.ETHEREUM_CROSSCHAIN_ID)
		if ethereumConfig == nil {
			panic("chain is invalid")
		}
		amount, err := ethereumSdk.Erc20Balance(hash, ethereumConfig.ProxyContract)
		if err == nil {
			proxyBalance := new(ProxyBalance)
			proxyBalance.Amount = amount
			proxyBalance.ItemName = "poly"
			proxyBalance.ItemProxy = ethereumConfig.ProxyContract
			proxyBalances = append(proxyBalances, proxyBalance)
		}
		for _, otherProxyContract := range ethereumConfig.OtherProxyContract {
			amount, err := ethereumSdk.Erc20Balance(hash, otherProxyContract.ItemProxy)
			if err != nil {
				continue
			}
			proxyBalance := new(ProxyBalance)
			proxyBalance.Amount = amount
			proxyBalance.ItemName = otherProxyContract.ItemName
			proxyBalance.ItemProxy = otherProxyContract.ItemProxy
			proxyBalances = append(proxyBalances, proxyBalance)
		}
		return proxyBalances
	}
	if chainId == basedef.MATIC_CROSSCHAIN_ID {
		maticConfig := config.GetChainListenConfig(basedef.MATIC_CROSSCHAIN_ID)
		if maticConfig == nil {
			panic("chain is invalid")
		}
		amount, err := maticSdk.Erc20Balance(hash, maticConfig.ProxyContract)
		if err == nil {
			proxyBalance := new(ProxyBalance)
			proxyBalance.Amount = amount
			proxyBalance.ItemName = "poly"
			proxyBalance.ItemProxy = maticConfig.ProxyContract
			proxyBalances = append(proxyBalances, proxyBalance)
		}
		for _, otherProxyContract := range maticConfig.OtherProxyContract {
			amount, err := maticSdk.Erc20Balance(hash, otherProxyContract.ItemProxy)
			if err != nil {
				continue
			}
			proxyBalance := new(ProxyBalance)
			proxyBalance.Amount = amount
			proxyBalance.ItemName = otherProxyContract.ItemName
			proxyBalance.ItemProxy = otherProxyContract.ItemProxy
			proxyBalances = append(proxyBalances, proxyBalance)
		}
		return proxyBalances
	}
	if chainId == basedef.BSC_CROSSCHAIN_ID {
		bscConfig := config.GetChainListenConfig(basedef.BSC_CROSSCHAIN_ID)
		if bscConfig == nil {
			panic("chain is invalid")
		}
		amount, err := bscSdk.Erc20Balance(hash, bscConfig.ProxyContract)
		if err == nil {
			proxyBalance := new(ProxyBalance)
			proxyBalance.Amount = amount
			proxyBalance.ItemName = "poly"
			proxyBalance.ItemProxy = bscConfig.ProxyContract
			proxyBalances = append(proxyBalances, proxyBalance)
		}
		for _, otherProxyContract := range bscConfig.OtherProxyContract {
			amount, err := bscSdk.Erc20Balance(hash, otherProxyContract.ItemProxy)
			if err != nil {
				continue
			}
			proxyBalance := new(ProxyBalance)
			proxyBalance.Amount = amount
			proxyBalance.ItemName = otherProxyContract.ItemName
			proxyBalance.ItemProxy = otherProxyContract.ItemProxy
			proxyBalances = append(proxyBalances, proxyBalance)
		}
		return proxyBalances
	}
	if chainId == basedef.HECO_CROSSCHAIN_ID {
		hecoConfig := config.GetChainListenConfig(basedef.HECO_CROSSCHAIN_ID)
		if hecoConfig == nil {
			panic("chain is invalid")
		}
		amount, err := hecoSdk.Erc20Balance(hash, hecoConfig.ProxyContract)
		if err == nil {
			proxyBalance := new(ProxyBalance)
			proxyBalance.Amount = amount
			proxyBalance.ItemName = "poly"
			proxyBalance.ItemProxy = hecoConfig.ProxyContract
			proxyBalances = append(proxyBalances, proxyBalance)
		}
		for _, otherProxyContract := range hecoConfig.OtherProxyContract {
			amount, err := hecoSdk.Erc20Balance(hash, otherProxyContract.ItemProxy)
			if err != nil {
				continue
			}
			proxyBalance := new(ProxyBalance)
			proxyBalance.Amount = amount
			proxyBalance.ItemName = otherProxyContract.ItemName
			proxyBalance.ItemProxy = otherProxyContract.ItemProxy
			proxyBalances = append(proxyBalances, proxyBalance)
		}
		return proxyBalances
	}
	if chainId == basedef.OK_CROSSCHAIN_ID {
		okConfig := config.GetChainListenConfig(basedef.OK_CROSSCHAIN_ID)
		if okConfig == nil {
			panic("chain is invalid")
		}
		amount, err := okSdk.Erc20Balance(hash, okConfig.ProxyContract)
		if err == nil {
			proxyBalance := new(ProxyBalance)
			proxyBalance.Amount = amount
			proxyBalance.ItemName = "poly"
			proxyBalance.ItemProxy = okConfig.ProxyContract
			proxyBalances = append(proxyBalances, proxyBalance)
		}
		for _, otherProxyContract := range okConfig.OtherProxyContract {
			amount, err := okSdk.Erc20Balance(hash, otherProxyContract.ItemProxy)
			if err != nil {
				continue
			}
			proxyBalance := new(ProxyBalance)
			proxyBalance.Amount = amount
			proxyBalance.ItemName = otherProxyContract.ItemName
			proxyBalance.ItemProxy = otherProxyContract.ItemProxy
			proxyBalances = append(proxyBalances, proxyBalance)
		}
		return proxyBalances
	}
	if chainId == basedef.NEO_CROSSCHAIN_ID {
		neoConfig := config.GetChainListenConfig(basedef.NEO_CROSSCHAIN_ID)
		if neoConfig == nil {
			panic("chain is invalid")
		}
		amount, err := neoSdk.Nep5Balance(hash, neoConfig.ProxyContract)
		if err == nil {
			proxyBalance := new(ProxyBalance)
			proxyBalance.Amount = amount
			proxyBalance.ItemName = "poly"
			proxyBalance.ItemProxy = neoConfig.ProxyContract
			proxyBalances = append(proxyBalances, proxyBalance)
		}
		for _, otherProxyContract := range neoConfig.OtherProxyContract {
			amount, err := neoSdk.Nep5Balance(hash, otherProxyContract.ItemProxy)
			if err != nil {
				continue
			}
			proxyBalance := new(ProxyBalance)
			proxyBalance.Amount = amount
			proxyBalance.ItemName = otherProxyContract.ItemName
			proxyBalance.ItemProxy = otherProxyContract.ItemProxy
			proxyBalances = append(proxyBalances, proxyBalance)
		}
		return proxyBalances
	}
	if chainId == basedef.ONT_CROSSCHAIN_ID {
		ontConfig := config.GetChainListenConfig(basedef.ONT_CROSSCHAIN_ID)
		if ontConfig == nil {
			panic("chain is invalid")
		}
		amount, err := ontologySdk.Oep4Balance(hash, ontConfig.ProxyContract)
		if err == nil {
			proxyBalance := new(ProxyBalance)
			proxyBalance.Amount = amount
			proxyBalance.ItemName = "poly"
			proxyBalance.ItemProxy = ontConfig.ProxyContract
			proxyBalances = append(proxyBalances, proxyBalance)
		}
		for _, otherProxyContract := range ontConfig.OtherProxyContract {
			amount, err := ontologySdk.Oep4Balance(hash, otherProxyContract.ItemProxy)
			if err != nil {
				continue
			}
			proxyBalance := new(ProxyBalance)
			proxyBalance.Amount = amount
			proxyBalance.ItemName = otherProxyContract.ItemName
			proxyBalance.ItemProxy = otherProxyContract.ItemProxy
			proxyBalances = append(proxyBalances, proxyBalance)
		}
		return proxyBalances
	}
	if chainId == basedef.ARBITRUM_CROSSCHAIN_ID {
		arbitrumConfig := config.GetChainListenConfig(basedef.ARBITRUM_CROSSCHAIN_ID)
		if arbitrumConfig == nil {
			panic("chain is invalid")
		}
		amount, err := arbitrumSdk.Erc20Balance(hash, arbitrumConfig.ProxyContract)
		if err == nil {
			proxyBalance := new(ProxyBalance)
			proxyBalance.Amount = amount
			proxyBalance.ItemName = "poly"
			proxyBalance.ItemProxy = arbitrumConfig.ProxyContract
			proxyBalances = append(proxyBalances, proxyBalance)
		}
		for _, otherProxyContract := range arbitrumConfig.OtherProxyContract {
			amount, err := arbitrumSdk.Erc20Balance(hash, otherProxyContract.ItemProxy)
			if err != nil {
				continue
			}
			proxyBalance := new(ProxyBalance)
			proxyBalance.Amount = amount
			proxyBalance.ItemName = otherProxyContract.ItemName
			proxyBalance.ItemProxy = otherProxyContract.ItemProxy
			proxyBalances = append(proxyBalances, proxyBalance)
		}
		return proxyBalances
	}
	return proxyBalances
}

func GetProxyBalance(chainId uint64, hash string, proxy string) (*big.Int, error) {
	switch chainId {
	case basedef.ETHEREUM_CROSSCHAIN_ID:
		return ethereumSdk.Erc20Balance(hash, proxy)
	case basedef.MATIC_CROSSCHAIN_ID:
		return maticSdk.Erc20Balance(hash, proxy)
	case basedef.BSC_CROSSCHAIN_ID:
		return bscSdk.Erc20Balance(hash, proxy)
	case basedef.HECO_CROSSCHAIN_ID:
		return hecoSdk.Erc20Balance(hash, proxy)
	case basedef.OK_CROSSCHAIN_ID:
		return okSdk.Erc20Balance(hash, proxy)
	case basedef.NEO_CROSSCHAIN_ID:
		return neoSdk.Nep5Balance(hash, proxy)
	case basedef.ONT_CROSSCHAIN_ID:
		return ontologySdk.Oep4Balance(hash, proxy)
	case basedef.ARBITRUM_CROSSCHAIN_ID:
		return arbitrumSdk.Erc20Balance(hash, proxy)
	case basedef.XDAI_CROSSCHAIN_ID:
		return xdaiSdk.Erc20Balance(hash, proxy)

	default:
		return new(big.Int).SetUint64(0), nil
	}
}
