package starcoinfee

import (
	"context"
	"fmt"
	"github.com/polynetwork/bridge-common/chains/starcoin"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"time"
)

type StarcoinFee struct {
	starcoinCfg *conf.FeeListenConfig
	starcoinSdk *starcoin.SDK
}

func NewStarcoinFee(starcoinCfg *conf.FeeListenConfig, feeUpdateSlot int64) *StarcoinFee {
	StarcoinFee := &StarcoinFee{}
	StarcoinFee.starcoinCfg = starcoinCfg
	sdk, err := starcoin.WithOptions(starcoinCfg.ChainId, starcoinCfg.Nodes, time.Minute, 1)
	if err != nil {
		panic(fmt.Sprintf("Create chain sdk failed. chain=%d, err=%s", starcoinCfg.ChainId, err))
	}
	StarcoinFee.starcoinSdk = sdk
	return StarcoinFee
}

func (this *StarcoinFee) GetFee() (*big.Int, *big.Int, *big.Int, error) {
	suggestGasPrice, err := this.starcoinSdk.Node().GetGasUnitPrice(context.Background())
	if err != nil {
		return nil, nil, nil, err
	}
	gasPrice := big.NewInt(int64(suggestGasPrice))

	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(basedef.FEE_PRECISION))
	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(this.starcoinCfg.GasLimit))
	proxyFee := new(big.Int).Mul(gasPrice, new(big.Int).SetInt64(this.starcoinCfg.ProxyFee))
	proxyFee = new(big.Int).Div(proxyFee, new(big.Int).SetInt64(100))
	minFee := new(big.Int).Mul(gasPrice, new(big.Int).SetInt64(this.starcoinCfg.MinFee))
	minFee = new(big.Int).Div(minFee, new(big.Int).SetInt64(100))
	return minFee, gasPrice, proxyFee, nil
}

func (this *StarcoinFee) GetChainId() uint64 {
	return this.starcoinCfg.ChainId
}

func (this *StarcoinFee) Name() string {
	return this.starcoinCfg.ChainName
}
