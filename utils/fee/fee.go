package fee

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/polynetwork/poly-bridge/basedef"
	"github.com/polynetwork/poly-bridge/conf"
	"github.com/polynetwork/poly-bridge/models"
	"math/big"
)

func GetL1Fee(ethChainFee *models.ChainFee, chainId uint64) (l1MinFee, l1ProxyFee, l1FeeAmount *big.Float, err error) {
	var targetFeeListenConfig, ethFeeListenConfig *conf.FeeListenConfig
	for _, fl := range conf.GlobalConfig.FeeListenConfig {
		if fl.ChainId == chainId {
			targetFeeListenConfig = fl
			break
		}
		continue
	}
	for _, fl := range conf.GlobalConfig.FeeListenConfig {
		if fl.ChainId == basedef.ETHEREUM_CROSSCHAIN_ID {
			ethFeeListenConfig = fl
			break
		}
		continue
	}

	if targetFeeListenConfig == nil || ethFeeListenConfig == nil {
		err := fmt.Errorf("chain listen config is missing")
		logs.Error("getOptimisticL1FeeMin error: %v", err)
		return nil, nil, nil, err
	}

	gasLimitScale := new(big.Float).Quo(new(big.Float).SetInt64(targetFeeListenConfig.EthL1GasLimit), new(big.Float).SetInt64(ethFeeListenConfig.GasLimit))
	price := new(big.Float).SetInt64(ethChainFee.TokenBasic.Price)
	precisionFactor := new(big.Float).Mul(new(big.Float).SetInt64(basedef.PRICE_PRECISION), new(big.Float).SetInt64(basedef.FEE_PRECISION))
	precisionFactor = new(big.Float).Mul(precisionFactor, new(big.Float).SetInt64(basedef.Int64FromFigure(int(ethChainFee.TokenBasic.Precision))))

	feeFactor := new(big.Float).Mul(gasLimitScale, price)
	feeFactor = new(big.Float).Quo(feeFactor, precisionFactor)

	l1MinFee = new(big.Float).Mul(new(big.Float).SetInt(&ethChainFee.MinFee.Int), feeFactor)
	l1ProxyFee = new(big.Float).Mul(new(big.Float).SetInt(&ethChainFee.ProxyFee.Int), feeFactor)

	l1FeeAmount = new(big.Float).Mul(new(big.Float).SetInt(&ethChainFee.ProxyFee.Int), gasLimitScale)
	l1FeeAmount = l1FeeAmount.Quo(l1FeeAmount, new(big.Float).SetInt64(basedef.FEE_PRECISION))
	l1FeeAmount = l1FeeAmount.Quo(l1FeeAmount, new(big.Float).SetInt64(basedef.Int64FromFigure(int(ethChainFee.TokenBasic.Precision))))
	logs.Info("chain:%d l1MinFee=%s, l1ProxyFee=%s, l1FeeAmount=%s", chainId, l1MinFee.String(), l1ProxyFee.String(), l1FeeAmount.String())
	return
}
