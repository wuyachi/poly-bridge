package aptosfee

import (
	"context"
	"fmt"
	"github.com/polynetwork/bridge-common/chains/aptos"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"time"
)

type AptosFee struct {
	aptosCfg *conf.FeeListenConfig
	aptosSdk *aptos.SDK
}

func NewAptosFee(aptosCfg *conf.FeeListenConfig, feeUpdateSlot int64) *AptosFee {
	aptosFee := &AptosFee{}
	aptosFee.aptosCfg = aptosCfg
	sdk, err := aptos.WithOptions(aptosCfg.ChainId, aptosCfg.Nodes, time.Minute, 1)
	if err != nil {
		panic(fmt.Sprintf("Create chain sdk failed. chain=%d, err=%s", aptosCfg.ChainId, err))
	}
	aptosFee.aptosSdk = sdk
	return aptosFee
}

func (this *AptosFee) GetFee() (*big.Int, *big.Int, *big.Int, error) {
	suggestGasPrice, err := this.aptosSdk.Node().EstimateGasPrice(context.Background())
	if err != nil {
		return nil, nil, nil, err
	}
	gasPrice := big.NewInt(int64(suggestGasPrice))

	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(basedef.FEE_PRECISION))
	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(this.aptosCfg.GasLimit))
	proxyFee := new(big.Int).Mul(gasPrice, new(big.Int).SetInt64(this.aptosCfg.ProxyFee))
	proxyFee = new(big.Int).Div(proxyFee, new(big.Int).SetInt64(100))
	minFee := new(big.Int).Mul(gasPrice, new(big.Int).SetInt64(this.aptosCfg.MinFee))
	minFee = new(big.Int).Div(minFee, new(big.Int).SetInt64(100))
	return minFee, gasPrice, proxyFee, nil
}

func (this *AptosFee) GetChainId() uint64 {
	return this.aptosCfg.ChainId
}

func (this *AptosFee) Name() string {
	return this.aptosCfg.ChainName
}
