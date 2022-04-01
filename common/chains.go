package common

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"strings"

	"github.com/beego/beego/v2/core/logs"
)

var (
	ethereumSdk   *chainsdk.EthereumSdkPro
	pltSdk        *chainsdk.EthereumSdkPro
	bscSdk        *chainsdk.EthereumSdkPro
	hecoSdk       *chainsdk.EthereumSdkPro
	okSdk         *chainsdk.EthereumSdkPro
	neoSdk        *chainsdk.NeoSdkPro
	neo3Sdk       *chainsdk.Neo3SdkPro
	ontologySdk   *chainsdk.OntologySdkPro
	maticSdk      *chainsdk.EthereumSdkPro
	swthSdk       *chainsdk.SwitcheoSdkPro
	arbitrumSdk   *chainsdk.EthereumSdkPro
	zilliqaSdk    *chainsdk.ZilliqaSdkPro
	xdaiSdk       *chainsdk.EthereumSdkPro
	fantomSdk     *chainsdk.EthereumSdkPro
	avaxSdk       *chainsdk.EthereumSdkPro
	optimisticSdk *chainsdk.EthereumSdkPro
	zionmainSdk   *chainsdk.EthereumSdkPro
	sidechainSdk  *chainsdk.EthereumSdkPro
	kovanSdk      *chainsdk.EthereumSdkPro
	rinkebySdk    *chainsdk.EthereumSdkPro
	goerliSdk     *chainsdk.EthereumSdkPro
	metisSdk      *chainsdk.EthereumSdkPro
	bobaSdk       *chainsdk.EthereumSdkPro
	oasisSdk      *chainsdk.EthereumSdkPro
	sdkMap        map[uint64]interface{}
	config        *conf.Config
)

func SetupChainsSDK(cfg *conf.Config) {
	if cfg == nil {
		panic("Missing config")
	}
	config = cfg
	newChainSdks(cfg)
}

func newChainSdks(config *conf.Config) {
	sdkMap = make(map[uint64]interface{}, 0)
	for _, cfg := range config.ChainListenConfig {
		switch cfg.ChainId {
		case basedef.ETHEREUM_CROSSCHAIN_ID:
			conf := config.GetChainListenConfig(basedef.ETHEREUM_CROSSCHAIN_ID)
			if conf == nil {
				logs.Error("Missing ETHEREUM chain sdk config")
			}
			urls := conf.GetNodesUrl()
			ethereumSdk = chainsdk.NewEthereumSdkPro(urls, conf.ListenSlot, conf.ChainId)
			sdkMap[basedef.ETHEREUM_CROSSCHAIN_ID] = ethereumSdk
		case basedef.ZIONMAIN_CROSSCHAIN_ID:
			conf := config.GetChainListenConfig(basedef.ZIONMAIN_CROSSCHAIN_ID)
			if conf == nil {
				logs.Error("Missing ZIONMAIN chain sdk config")
			}
			urls := conf.GetNodesUrl()
			zionmainSdk = chainsdk.NewEthereumSdkPro(urls, conf.ListenSlot, conf.ChainId)
		case basedef.SIDECHAIN_CROSSCHAIN_ID:
			conf := config.GetChainListenConfig(basedef.SIDECHAIN_CROSSCHAIN_ID)
			if conf == nil {
				logs.Error("Missing SIDECHAIN chain sdk config")
			}
			urls := conf.GetNodesUrl()
			sidechainSdk = chainsdk.NewEthereumSdkPro(urls, conf.ListenSlot, conf.ChainId)
		case basedef.MATIC_CROSSCHAIN_ID:
			conf := config.GetChainListenConfig(basedef.MATIC_CROSSCHAIN_ID)
			if conf == nil {
				logs.Error("Missing MATIC chain sdk config")
			}
			urls := conf.GetNodesUrl()
			maticSdk = chainsdk.NewEthereumSdkPro(urls, conf.ListenSlot, conf.ChainId)
		case basedef.BSC_CROSSCHAIN_ID:
			conf := config.GetChainListenConfig(basedef.BSC_CROSSCHAIN_ID)
			if conf == nil {
				logs.Error("Missing MATIC chain sdk config")
			}
			urls := conf.GetNodesUrl()
			bscSdk = chainsdk.NewEthereumSdkPro(urls, conf.ListenSlot, conf.ChainId)
		case basedef.HECO_CROSSCHAIN_ID:
			conf := config.GetChainListenConfig(basedef.HECO_CROSSCHAIN_ID)
			if conf == nil {
				logs.Error("Missing HECO chain sdk config")
			}
			urls := conf.GetNodesUrl()
			hecoSdk = chainsdk.NewEthereumSdkPro(urls, conf.ListenSlot, conf.ChainId)
		case basedef.OK_CROSSCHAIN_ID:
			conf := config.GetChainListenConfig(basedef.OK_CROSSCHAIN_ID)
			if conf == nil {
				logs.Error("Missing OK chain sdk config")
			}
			urls := conf.GetNodesUrl()
			okSdk = chainsdk.NewEthereumSdkPro(urls, conf.ListenSlot, conf.ChainId)
		case basedef.PLT_CROSSCHAIN_ID:
			conf := config.GetChainListenConfig(basedef.PLT_CROSSCHAIN_ID)
			if conf == nil {
				logs.Error("Missing PLT chain sdk config")
			}
			urls := conf.GetNodesUrl()
			pltSdk = chainsdk.NewEthereumSdkPro(urls, conf.ListenSlot, conf.ChainId)
		case basedef.KOVAN_CROSSCHAIN_ID:
			if basedef.ENV == basedef.TESTNET {
				conf := config.GetChainListenConfig(basedef.KOVAN_CROSSCHAIN_ID)
				if conf == nil {
					logs.Error("Missing KOVAN chain sdk config")
				}
				urls := conf.GetNodesUrl()
				kovanSdk = chainsdk.NewEthereumSdkPro(urls, conf.ListenSlot, conf.ChainId)
			}
		case basedef.RINKEBY_CROSSCHAIN_ID:
			if basedef.ENV == basedef.TESTNET {
				conf := config.GetChainListenConfig(basedef.RINKEBY_CROSSCHAIN_ID)
				if conf == nil {
					logs.Error("Missing RINKEBY chain sdk config")
				}
				urls := conf.GetNodesUrl()
				rinkebySdk = chainsdk.NewEthereumSdkPro(urls, conf.ListenSlot, conf.ChainId)
			}
		case basedef.GOERLI_CROSSCHAIN_ID:
			if basedef.ENV == basedef.TESTNET {
				conf := config.GetChainListenConfig(basedef.GOERLI_CROSSCHAIN_ID)
				if conf == nil {
					logs.Error("Missing GOERLI chain sdk config")
				}
				urls := conf.GetNodesUrl()
				goerliSdk = chainsdk.NewEthereumSdkPro(urls, conf.ListenSlot, conf.ChainId)
			}
		case basedef.NEO_CROSSCHAIN_ID:
			conf := config.GetChainListenConfig(basedef.NEO_CROSSCHAIN_ID)
			if conf == nil {
				logs.Error("Missing NEO chain sdk config")
			}
			urls := conf.GetNodesUrl()
			neoSdk = chainsdk.NewNeoSdkPro(urls, conf.ListenSlot, conf.ChainId)
			sdkMap[basedef.NEO_CROSSCHAIN_ID] = neoSdk
		case basedef.NEO3_CROSSCHAIN_ID:
			conf := config.GetChainListenConfig(basedef.NEO3_CROSSCHAIN_ID)
			if conf == nil {
				logs.Error("Missing NEO3 chain sdk config")
			}
			urls := conf.GetNodesUrl()
			neo3Sdk = chainsdk.NewNeo3SdkPro(urls, conf.ListenSlot, conf.ChainId)
			sdkMap[basedef.NEO3_CROSSCHAIN_ID] = neo3Sdk
		case basedef.ONT_CROSSCHAIN_ID:
			conf := config.GetChainListenConfig(basedef.ONT_CROSSCHAIN_ID)
			if conf == nil {
				logs.Error("Missing ONT chain sdk config")
			}
			urls := conf.GetNodesUrl()
			ontologySdk = chainsdk.NewOntologySdkPro(urls, conf.ListenSlot, conf.ChainId)
			sdkMap[basedef.ONT_CROSSCHAIN_ID] = ontologySdk
		case basedef.SWITCHEO_CROSSCHAIN_ID:
			if basedef.ENV == basedef.MAINNET {
				conf := config.GetChainListenConfig(basedef.SWITCHEO_CROSSCHAIN_ID)
				if conf == nil {
					logs.Error("Missing SWITCHEO chain sdk config")
				}
				urls := conf.GetNodesUrl()
				swthSdk = chainsdk.NewSwitcheoSdkPro(urls, conf.ListenSlot, conf.ChainId)
				sdkMap[basedef.SWITCHEO_CROSSCHAIN_ID] = swthSdk
			}
		case basedef.ARBITRUM_CROSSCHAIN_ID:
			conf := config.GetChainListenConfig(basedef.ARBITRUM_CROSSCHAIN_ID)
			if conf == nil {
				logs.Error("Missing Arbitrum chain sdk config")
			}
			urls := conf.GetNodesUrl()
			arbitrumSdk = chainsdk.NewEthereumSdkPro(urls, conf.ListenSlot, conf.ChainId)
			sdkMap[basedef.ARBITRUM_CROSSCHAIN_ID] = arbitrumSdk
		case basedef.XDAI_CROSSCHAIN_ID:
			conf := config.GetChainListenConfig(basedef.XDAI_CROSSCHAIN_ID)
			if conf == nil {
				logs.Error("Missing XDai chain sdk config")
			}
			urls := conf.GetNodesUrl()
			xdaiSdk = chainsdk.NewEthereumSdkPro(urls, conf.ListenSlot, conf.ChainId)
			sdkMap[basedef.XDAI_CROSSCHAIN_ID] = xdaiSdk
		case basedef.ZILLIQA_CROSSCHAIN_ID:
			conf := config.GetChainListenConfig(basedef.ZILLIQA_CROSSCHAIN_ID)
			if conf == nil {
				logs.Error("Missing Zilliqa chain sdk config")
			}
			urls := conf.GetNodesUrl()
			zilliqaSdk = chainsdk.NewZilliqaSdkPro(urls, conf.ListenSlot, conf.ChainId)
			sdkMap[basedef.ZILLIQA_CROSSCHAIN_ID] = zilliqaSdk
		case basedef.FANTOM_CROSSCHAIN_ID:
			conf := config.GetChainListenConfig(basedef.FANTOM_CROSSCHAIN_ID)
			if conf == nil {
				logs.Error("Missing Fantom chain sdk config")
			}
			urls := conf.GetNodesUrl()
			fantomSdk = chainsdk.NewEthereumSdkPro(urls, conf.ListenSlot, conf.ChainId)
			sdkMap[basedef.FANTOM_CROSSCHAIN_ID] = fantomSdk
		case basedef.AVAX_CROSSCHAIN_ID:
			conf := config.GetChainListenConfig(basedef.AVAX_CROSSCHAIN_ID)
			if conf == nil {
				logs.Error("Missing Avax chain sdk config")
			}
			urls := conf.GetNodesUrl()
			avaxSdk = chainsdk.NewEthereumSdkPro(urls, conf.ListenSlot, conf.ChainId)
			sdkMap[basedef.AVAX_CROSSCHAIN_ID] = avaxSdk
		case basedef.OPTIMISTIC_CROSSCHAIN_ID:
			conf := config.GetChainListenConfig(basedef.OPTIMISTIC_CROSSCHAIN_ID)
			if conf == nil {
				logs.Error("Missing Optimistic chain sdk config")
			}
			urls := conf.GetNodesUrl()
			optimisticSdk = chainsdk.NewEthereumSdkPro(urls, conf.ListenSlot, conf.ChainId)
			sdkMap[basedef.OPTIMISTIC_CROSSCHAIN_ID] = optimisticSdk
		case basedef.METIS_CROSSCHAIN_ID:
			metisConfig := config.GetChainListenConfig(basedef.METIS_CROSSCHAIN_ID)
			if metisConfig == nil {
				logs.Error("Missing Metis chain sdk config")
			}
			urls := metisConfig.GetNodesUrl()
			metisSdk = chainsdk.NewEthereumSdkPro(urls, metisConfig.ListenSlot, metisConfig.ChainId)
			sdkMap[basedef.METIS_CROSSCHAIN_ID] = metisSdk
		case basedef.BOBA_CROSSCHAIN_ID:
			bobaConfig := config.GetChainListenConfig(basedef.BOBA_CROSSCHAIN_ID)
			if bobaConfig == nil {
				panic("boba chain is invalid")
			}
			urls := bobaConfig.GetNodesUrl()
			bobaSdk = chainsdk.NewEthereumSdkPro(urls, bobaConfig.ListenSlot, bobaConfig.ChainId)
			sdkMap[basedef.BOBA_CROSSCHAIN_ID] = bobaSdk
		case basedef.OASIS_CROSSCHAIN_ID:
			chainConfig := config.GetChainListenConfig(basedef.OASIS_CROSSCHAIN_ID)
			if chainConfig == nil {
				panic("oasis chain is invalid")
			}
			urls := chainConfig.GetNodesUrl()
			oasisSdk = chainsdk.NewEthereumSdkPro(urls, chainConfig.ListenSlot, chainConfig.ChainId)
			sdkMap[basedef.OASIS_CROSSCHAIN_ID] = oasisSdk
		}
	}

}

func GetBalance(chainId uint64, hash string) (*big.Int, error) {
	maxBalance := big.NewInt(0)
	maxFun := func(balance *big.Int) {
		if balance.Cmp(maxBalance) > 0 {
			maxBalance = balance
		}
	}
	errMap := make(map[error]bool, 0)
	switch chainId {
	case basedef.ETHEREUM_CROSSCHAIN_ID:
		config := config.GetChainListenConfig(basedef.ETHEREUM_CROSSCHAIN_ID)
		if config == nil {
			panic("chain is invalid")
		}
		for _, v := range config.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := ethereumSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	case basedef.ZIONMAIN_CROSSCHAIN_ID:
		config := config.GetChainListenConfig(basedef.ZIONMAIN_CROSSCHAIN_ID)
		if config == nil {
			panic("chain is invalid")
		}
		for _, v := range config.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := zionmainSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	case basedef.SIDECHAIN_CROSSCHAIN_ID:
		config := config.GetChainListenConfig(basedef.SIDECHAIN_CROSSCHAIN_ID)
		if config == nil {
			panic("chain is invalid")
		}
		for _, v := range config.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := sidechainSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	case basedef.MATIC_CROSSCHAIN_ID:
		config := config.GetChainListenConfig(basedef.MATIC_CROSSCHAIN_ID)
		if config == nil {
			panic("chain is invalid")
		}
		for _, v := range config.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := maticSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	case basedef.BSC_CROSSCHAIN_ID:
		config := config.GetChainListenConfig(basedef.BSC_CROSSCHAIN_ID)
		if config == nil {
			panic("chain is invalid")
		}
		for _, v := range config.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := bscSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	case basedef.HECO_CROSSCHAIN_ID:
		config := config.GetChainListenConfig(basedef.HECO_CROSSCHAIN_ID)
		if config == nil {
			panic("chain is invalid")
		}
		for _, v := range config.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := hecoSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	case basedef.OK_CROSSCHAIN_ID:
		config := config.GetChainListenConfig(basedef.OK_CROSSCHAIN_ID)
		if config == nil {
			panic("chain is invalid")
		}
		for _, v := range config.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := okSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	case basedef.KOVAN_CROSSCHAIN_ID:
		config := config.GetChainListenConfig(basedef.KOVAN_CROSSCHAIN_ID)
		if config == nil {
			panic("Missing kovan chain sdk config")
		}
		for _, v := range config.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := kovanSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	case basedef.RINKEBY_CROSSCHAIN_ID:
		config := config.GetChainListenConfig(basedef.RINKEBY_CROSSCHAIN_ID)
		if config == nil {
			panic("Missing rinkeby chain sdk config")
		}
		for _, v := range config.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := rinkebySdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	case basedef.GOERLI_CROSSCHAIN_ID:
		config := config.GetChainListenConfig(basedef.GOERLI_CROSSCHAIN_ID)
		if config == nil {
			panic("Missing goerli chain sdk config")
		}
		for _, v := range config.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := goerliSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	case basedef.NEO_CROSSCHAIN_ID:
		neoConfig := config.GetChainListenConfig(basedef.NEO_CROSSCHAIN_ID)
		if neoConfig == nil {
			panic("chain is invalid")
		}
		for _, v := range neoConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := neoSdk.Nep5Balance(hash, v)
			maxFun(balance)
			errMap[err] = true

		}
	case basedef.NEO3_CROSSCHAIN_ID:
		neo3Config := config.GetChainListenConfig(basedef.NEO3_CROSSCHAIN_ID)
		if neo3Config == nil {
			panic("chain is invalid")
		}
		for _, v := range neo3Config.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := neo3Sdk.Nep17Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	case basedef.ONT_CROSSCHAIN_ID:
		ontConfig := config.GetChainListenConfig(basedef.ONT_CROSSCHAIN_ID)
		if ontConfig == nil {
			panic("chain is invalid")
		}
		for _, v := range ontConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := ontologySdk.Oep4Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	case basedef.ARBITRUM_CROSSCHAIN_ID:
		arbitrumConfig := config.GetChainListenConfig(basedef.ARBITRUM_CROSSCHAIN_ID)
		if arbitrumConfig == nil {
			panic("chain is invalid")
		}
		for _, v := range arbitrumConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := arbitrumSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	case basedef.XDAI_CROSSCHAIN_ID:
		xdaiConfig := config.GetChainListenConfig(basedef.XDAI_CROSSCHAIN_ID)
		if xdaiConfig == nil {
			panic("chain is invalid")
		}
		for _, v := range xdaiConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := xdaiSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	case basedef.ZILLIQA_CROSSCHAIN_ID:
		zilliqaCfg := config.GetChainListenConfig(basedef.ZILLIQA_CROSSCHAIN_ID)
		if zilliqaCfg == nil {
			panic("zilliqa GetChainListenConfig chain is invalid")
		}
		for _, v := range zilliqaCfg.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := zilliqaSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	case basedef.FANTOM_CROSSCHAIN_ID:
		fantomConfig := config.GetChainListenConfig(basedef.FANTOM_CROSSCHAIN_ID)
		if fantomConfig == nil {
			panic("chain is invalid")
		}
		for _, v := range fantomConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := fantomSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	case basedef.AVAX_CROSSCHAIN_ID:
		avaxConfig := config.GetChainListenConfig(basedef.AVAX_CROSSCHAIN_ID)
		if avaxConfig == nil {
			panic("chain is invalid")
		}
		for _, v := range avaxConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := avaxSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	case basedef.OPTIMISTIC_CROSSCHAIN_ID:
		optimisticConfig := config.GetChainListenConfig(basedef.OPTIMISTIC_CROSSCHAIN_ID)
		if optimisticConfig == nil {
			panic("chain is invalid")
		}
		for _, v := range optimisticConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := optimisticSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	case basedef.METIS_CROSSCHAIN_ID:
		metisConfig := config.GetChainListenConfig(basedef.METIS_CROSSCHAIN_ID)
		if metisConfig == nil {
			panic("metis chain is invalid")
		}
		for _, v := range metisConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := metisSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	case basedef.BOBA_CROSSCHAIN_ID:
		bobaConfig := config.GetChainListenConfig(basedef.BOBA_CROSSCHAIN_ID)
		if bobaConfig == nil {
			panic("boba chain is invalid")
		}
		for _, v := range bobaConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := bobaSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	case basedef.OASIS_CROSSCHAIN_ID:
		chainConfig := config.GetChainListenConfig(basedef.OASIS_CROSSCHAIN_ID)
		if chainConfig == nil {
			panic("oasis chain is invalid")
		}
		for _, v := range chainConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := oasisSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	default:
		return new(big.Int).SetUint64(0), nil
	}

	if maxBalance.Cmp(big.NewInt(0)) > 0 {
		return maxBalance, nil
	}
	var err error
	for k, _ := range errMap {
		if k == nil {
			return new(big.Int).SetUint64(0), nil
		} else {
			err = k
		}
	}
	return new(big.Int).SetUint64(0), err

}

func GetTotalSupply(chainId uint64, hash string) (*big.Int, error) {
	switch chainId {
	case basedef.ETHEREUM_CROSSCHAIN_ID:
		ethereumConfig := config.GetChainListenConfig(basedef.ETHEREUM_CROSSCHAIN_ID)
		if ethereumConfig == nil {
			panic("chain is invalid")
		}
		return ethereumSdk.Erc20TotalSupply(hash)
	case basedef.ZIONMAIN_CROSSCHAIN_ID:
		zionmainConfig := config.GetChainListenConfig(basedef.ZIONMAIN_CROSSCHAIN_ID)
		if zionmainConfig == nil {
			panic("chain is invalid")
		}
		return zionmainSdk.Erc20TotalSupply(hash)
	case basedef.SIDECHAIN_CROSSCHAIN_ID:
		sidechainConfig := config.GetChainListenConfig(basedef.SIDECHAIN_CROSSCHAIN_ID)
		if sidechainConfig == nil {
			panic("chain is invalid")
		}
		return sidechainSdk.Erc20TotalSupply(hash)
	case basedef.MATIC_CROSSCHAIN_ID:
		maticConfig := config.GetChainListenConfig(basedef.MATIC_CROSSCHAIN_ID)
		if maticConfig == nil {
			panic("chain is invalid")
		}
		return maticSdk.Erc20TotalSupply(hash)
	case basedef.BSC_CROSSCHAIN_ID:
		bscConfig := config.GetChainListenConfig(basedef.BSC_CROSSCHAIN_ID)
		if bscConfig == nil {
			panic("chain is invalid")
		}
		return bscSdk.Erc20TotalSupply(hash)
	case basedef.HECO_CROSSCHAIN_ID:
		hecoConfig := config.GetChainListenConfig(basedef.HECO_CROSSCHAIN_ID)
		if hecoConfig == nil {
			panic("chain is invalid")
		}
		return hecoSdk.Erc20TotalSupply(hash)
	case basedef.OK_CROSSCHAIN_ID:
		okConfig := config.GetChainListenConfig(basedef.OK_CROSSCHAIN_ID)
		if okConfig == nil {
			panic("chain is invalid")
		}
		return okSdk.Erc20TotalSupply(hash)
	case basedef.KOVAN_CROSSCHAIN_ID:
		config := config.GetChainListenConfig(basedef.KOVAN_CROSSCHAIN_ID)
		if config == nil {
			panic("Missing kovan chain sdk config")
		}
		return kovanSdk.Erc20TotalSupply(hash)
	case basedef.RINKEBY_CROSSCHAIN_ID:
		config := config.GetChainListenConfig(basedef.RINKEBY_CROSSCHAIN_ID)
		if config == nil {
			panic("Missing rinkeby chain sdk config")
		}
		return rinkebySdk.Erc20TotalSupply(hash)

	case basedef.GOERLI_CROSSCHAIN_ID:
		config := config.GetChainListenConfig(basedef.GOERLI_CROSSCHAIN_ID)
		if config == nil {
			panic("Missing goerli chain sdk config")
		}
		return goerliSdk.Erc20TotalSupply(hash)
	case basedef.OPTIMISTIC_CROSSCHAIN_ID:
		optimisticConfig := config.GetChainListenConfig(basedef.OPTIMISTIC_CROSSCHAIN_ID)
		if optimisticConfig == nil {
			panic("chain is invalid")
		}
		return optimisticSdk.Erc20TotalSupply(hash)
	case basedef.NEO_CROSSCHAIN_ID:
		neoConfig := config.GetChainListenConfig(basedef.NEO_CROSSCHAIN_ID)
		if neoConfig == nil {
			panic("chain is invalid")
		}
		return neoSdk.Nep5TotalSupply(hash)
	case basedef.NEO3_CROSSCHAIN_ID:
		neo3Config := config.GetChainListenConfig(basedef.NEO3_CROSSCHAIN_ID)
		if neo3Config == nil {
			panic("chain is invalid")
		}
		return neo3Sdk.Nep17TotalSupply(hash)
	case basedef.ONT_CROSSCHAIN_ID:
		ontConfig := config.GetChainListenConfig(basedef.ONT_CROSSCHAIN_ID)
		if ontConfig == nil {
			panic("chain is invalid")
		}
		for _, v := range ontConfig.ProxyContract {
			if len(strings.TrimSpace(v)) != 0 {
				return ontologySdk.Oep4TotalSupply(hash, v)
			}
		}
		return new(big.Int).SetUint64(0), nil
	case basedef.ARBITRUM_CROSSCHAIN_ID:
		arbitrumConfig := config.GetChainListenConfig(basedef.ARBITRUM_CROSSCHAIN_ID)
		if arbitrumConfig == nil {
			panic("chain is invalid")
		}
		return arbitrumSdk.Erc20TotalSupply(hash)
	case basedef.XDAI_CROSSCHAIN_ID:
		xdaiConfig := config.GetChainListenConfig(basedef.XDAI_CROSSCHAIN_ID)
		if xdaiConfig == nil {
			panic("chain is invalid")
		}
		return xdaiSdk.Erc20TotalSupply(hash)
	case basedef.FANTOM_CROSSCHAIN_ID:
		fantomConfig := config.GetChainListenConfig(basedef.FANTOM_CROSSCHAIN_ID)
		if fantomConfig == nil {
			panic("chain is invalid")
		}
		return fantomSdk.Erc20TotalSupply(hash)
	case basedef.AVAX_CROSSCHAIN_ID:
		avaxConfig := config.GetChainListenConfig(basedef.AVAX_CROSSCHAIN_ID)
		if avaxConfig == nil {
			panic("chain is invalid")
		}
		return avaxSdk.Erc20TotalSupply(hash)
	case basedef.METIS_CROSSCHAIN_ID:
		metisConfig := config.GetChainListenConfig(basedef.METIS_CROSSCHAIN_ID)
		if metisConfig == nil {
			panic("metis chain GetTotalSupply invalid")
		}
		return metisSdk.Erc20TotalSupply(hash)
	case basedef.BOBA_CROSSCHAIN_ID:
		bobaConfig := config.GetChainListenConfig(basedef.BOBA_CROSSCHAIN_ID)
		if bobaConfig == nil {
			panic("boba chain GetTotalSupply invalid")
		}
		return bobaSdk.Erc20TotalSupply(hash)
	case basedef.OASIS_CROSSCHAIN_ID:
		chainConfig := config.GetChainListenConfig(basedef.OASIS_CROSSCHAIN_ID)
		if chainConfig == nil {
			panic("oasis chain GetTotalSupply invalid")
		}
		return oasisSdk.Erc20TotalSupply(hash)
	default:
		return new(big.Int).SetUint64(0), nil
	}
}

type ProxyBalance struct {
	Amount    *big.Int
	ItemName  string
	ItemProxy string
}

func GetProxyBalance(chainId uint64, hash string, proxy string) (*big.Int, error) {
	switch chainId {
	case basedef.ETHEREUM_CROSSCHAIN_ID:
		return ethereumSdk.Erc20Balance(hash, proxy)
	case basedef.ZIONMAIN_CROSSCHAIN_ID:
		return zionmainSdk.Erc20Balance(hash, proxy)
	case basedef.SIDECHAIN_CROSSCHAIN_ID:
		return sidechainSdk.Erc20Balance(hash, proxy)
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
	case basedef.ZILLIQA_CROSSCHAIN_ID:
		return zilliqaSdk.Erc20Balance(hash, proxy)
	case basedef.FANTOM_CROSSCHAIN_ID:
		return fantomSdk.Erc20Balance(hash, proxy)
	case basedef.AVAX_CROSSCHAIN_ID:
		return avaxSdk.Erc20Balance(hash, proxy)
	case basedef.OPTIMISTIC_CROSSCHAIN_ID:
		return optimisticSdk.Erc20Balance(hash, proxy)
	case basedef.KOVAN_CROSSCHAIN_ID:
		return kovanSdk.Erc20Balance(hash, proxy)
	case basedef.RINKEBY_CROSSCHAIN_ID:
		return rinkebySdk.Erc20Balance(hash, proxy)
	case basedef.GOERLI_CROSSCHAIN_ID:
		return goerliSdk.Erc20Balance(hash, proxy)
	case basedef.NEO3_CROSSCHAIN_ID:
		return neo3Sdk.Nep17Balance(hash, proxy)
	case basedef.METIS_CROSSCHAIN_ID:
		return metisSdk.Erc20Balance(hash, proxy)
	case basedef.BOBA_CROSSCHAIN_ID:
		return bobaSdk.Erc20Balance(hash, proxy)
	case basedef.OASIS_CROSSCHAIN_ID:
		return oasisSdk.Erc20Balance(hash, proxy)
	default:
		return new(big.Int).SetUint64(0), nil
	}
}

func GetNftOwner(chainId uint64, asset string, tokenId int) (owner common.Address, err error) {
	switch chainId {
	case basedef.ETHEREUM_CROSSCHAIN_ID:
		return ethereumSdk.GetNFTOwner(asset, big.NewInt(int64(tokenId)))
	case basedef.MATIC_CROSSCHAIN_ID:
		return maticSdk.GetNFTOwner(asset, big.NewInt(int64(tokenId)))
	case basedef.BSC_CROSSCHAIN_ID:
		return bscSdk.GetNFTOwner(asset, big.NewInt(int64(tokenId)))
	case basedef.HECO_CROSSCHAIN_ID:
		return hecoSdk.GetNFTOwner(asset, big.NewInt(int64(tokenId)))
	case basedef.OK_CROSSCHAIN_ID:
		return okSdk.GetNFTOwner(asset, big.NewInt(int64(tokenId)))
	case basedef.ARBITRUM_CROSSCHAIN_ID:
		return arbitrumSdk.GetNFTOwner(asset, big.NewInt(int64(tokenId)))
	case basedef.XDAI_CROSSCHAIN_ID:
		return xdaiSdk.GetNFTOwner(asset, big.NewInt(int64(tokenId)))
	default:
		return common.Address{}, fmt.Errorf("has nat func with chain:%v", chainId)
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
