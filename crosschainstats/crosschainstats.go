/*
 * Copyright (C) 2020 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */

package crosschainstats

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/common"
	"poly-bridge/conf"
	"poly-bridge/crosschaindao/bridgedao"
	"poly-bridge/http/tools"
	"poly-bridge/models"
	"sync"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Stats struct {
	context.Context
	cancel context.CancelFunc
	cfg    *conf.StatsConfig
	dao    *bridgedao.BridgeDao
	wg     sync.WaitGroup
}

var ccs *Stats

// Start - Do stats aggregation/calculation
func StartCrossChainStats(server string, cfg *conf.StatsConfig, dbCfg *conf.DBConfig) {
	if server != basedef.SERVER_POLY_BRIDGE {
		panic("CrossChainStats Only runs on bridge server")
	}
	if cfg == nil || cfg.TokenBasicStatsInterval == 0 || cfg.TokenAmountCheckInterval == 0 {
		panic("Invalid Stats config")
	}

	dao := bridgedao.NewBridgeDao(dbCfg, false)
	ctx, cancel := context.WithCancel(context.Background())
	ccs = &Stats{dao: dao, cfg: cfg, Context: ctx, cancel: cancel}
	ccs.Start()
}

// Stop
func StopCrossChainStats() {
	if ccs != nil {
		ccs.Stop()
	}
}

func (this *Stats) run(interval int64, f func() error) {
	this.wg.Add(1)
	ticker := time.NewTicker(time.Second * time.Duration(interval))
	for {
		select {
		case <-ticker.C:
			err := f()
			if err != nil {
				logs.Error("stats run error%s", err)
			}
		case <-this.Done():
			break
		}
	}
	this.wg.Done()
}

func (this *Stats) Start() {
	go this.run(this.cfg.TokenBasicStatsInterval, this.computeStats)
	go this.run(this.cfg.TokenAmountCheckInterval, this.computeTokensStats)
	go this.run(this.cfg.TokenStatisticInterval, this.computeTokenStatistics)
	go this.run(this.cfg.ChainStatisticInterval, this.computeChainStatistics)
	go this.run(this.cfg.ChainAddressCheckInterval, this.computeChainStatisticAssets)
	go this.run(this.cfg.AssetStatisticInterval, this.computeAssetStatistics)
	go this.run(this.cfg.AssetAdressInterval, this.computeAssetStatisticAdress)
}

func (this *Stats) Stop() {
	logs.Info("Stopping stats server")
	this.cancel()
	this.wg.Wait()
}

func (this *Stats) computeStats() (err error) {
	logs.Info("Computing cross chain token basic stats")
	tokens, err := this.dao.GetTokenBasics()
	if err != nil {
		return fmt.Errorf("Failed to fetch token basic list %w", err)
	}
	for _, basic := range tokens {
		if len(basic.Tokens) > 0 {
			err := this.computeTokenBasicStats(basic)
			if err != nil {
				logs.Error("Failed to computeTokenBasicStats for %s err %v", basic.Name, err)
			}
		}
	}
	return
}

func (this *Stats) computeTokenBasicStats(token *models.TokenBasic) (err error) {
	logs.Info("Computing token basic stats for %s", token.Name)
	if len(token.Tokens) == 0 {
		return
	}
	assets := make([][]interface{}, len(token.Tokens))
	for i, t := range token.Tokens {
		assets[i] = []interface{}{t.ChainId, t.Hash}
	}
	checkPoint := token.StatsUpdateTime
	last, err := this.dao.GetLastSrcTransferForToken(assets)
	if err != nil || last == nil || checkPoint == int64(last.Time) {
		if err != nil || last == nil {
			logs.Error("Failed to get last src transfers for %s %+v err %v", *token, assets, err)
		}
		return err
	}
	totalAmount, totalCount, err := this.dao.AggregateTokenBasicSrcTransfers(assets, checkPoint, last.Id)
	if err != nil {
		return err
	}
	token.StatsUpdateTime = last.Id
	if checkPoint == 0 {
		token.TotalAmount = &models.BigInt{*totalAmount}
		token.TotalCount = totalCount
	} else {
		token.TotalAmount = &models.BigInt{*new(big.Int).Add(totalAmount, &token.TotalAmount.Int)}
		token.TotalCount += totalCount
	}
	v := new(big.Float).Quo(new(big.Float).SetInt(&token.TotalAmount.Int), new(big.Float).SetInt64(basedef.Int64FromFigure(int(token.Precision))))
	f, _ := v.Float32()
	tools.Record(f, "total_amount.%s", token.Name)
	tools.Record(token.TotalCount, "total_count.%s", token.Name)
	err = this.dao.UpdateTokenBasicStatsWithCheckPoint(token, checkPoint)
	return
}

func (this *Stats) computeTokensStats() (err error) {
	logs.Info("Computing cross chain token stats")
	tokens, err := this.dao.GetTokens()
	if err != nil {
		return fmt.Errorf("Failed to fetch token basic list %w", err)
	}
	for _, t := range tokens {
		amount, err := common.GetBalance(t.ChainId, t.Hash)
		if err != nil || amount == nil {
			logs.Error("Failed to fetch token available amount for token %s %v %s", t.Hash, t.ChainId, err)
			continue
		}
		v := new(big.Float).Quo(new(big.Float).SetInt(amount), new(big.Float).SetInt64(basedef.Int64FromFigure(int(t.Precision))))
		f, _ := v.Float32()
		tools.Record(f, "balance.%s.%v", t.TokenBasicName, t.ChainId)
		err = this.dao.UpdateTokenAvailableAmount(t.Hash, t.ChainId, amount)
		if err != nil {
			logs.Error("Failed to update token available amount for token %s %v %s", t.Hash, t.ChainId, err)
		}
	}
	return
}

func (this *Stats) computeTokenStatistics() (err error) {
	logs.Info("start computeTokenStatistics")
	newDst, err := this.dao.GetNewDstTransfer()
	if err != nil {
		return fmt.Errorf("Failed to GetNewDstTransfer %w", err)
	}
	nowInId := newDst.Id
	newSrc, err := this.dao.GetNewSrcTransfer()
	if err != nil {
		return fmt.Errorf("Failed to GetNewSrcTransfer %w", err)
	}
	nowOutId := newSrc.Id
	nowTokenStatistic, err := this.dao.GetNewTokenSta()
	if err != nil {
		return fmt.Errorf("Failed to GetNewTokenSta %w", err)
	}
	inTokenStatistics := make([]*models.TokenStatistic, 0)
	if nowInId > nowTokenStatistic.LastInCheckId {
		err = this.dao.CalculateInTokenStatistics(nowTokenStatistic.LastInCheckId, nowInId, &inTokenStatistics)
		logs.Info("nowInId > nowTokenStatistic.LastInCheckId and CalculateInTokenStatistics success")
		if err != nil {
			return fmt.Errorf("Failed to CalculateInTokenStatistics %w", err)
		}
	}
	outTokenStatistics := make([]*models.TokenStatistic, 0)
	if nowOutId > nowTokenStatistic.LastOutCheckId {
		err = this.dao.CalculateOutTokenStatistics(nowTokenStatistic.LastOutCheckId, nowOutId, &outTokenStatistics)
		logs.Info("nowOutId > nowTokenStatistic.LastOutCheckId and CalculateOutTokenStatistics success")
		if err != nil {
			return fmt.Errorf("Failed to CalculateInTokenStatistics %w", err)
		}
	}
	logs.Info("nowInId:", nowInId, "LastInCheckId:", nowTokenStatistic.LastInCheckId, "nowOutId:", nowOutId, "nowTokenStatistic.LastOutCheckId", nowTokenStatistic.LastOutCheckId)
	if nowInId > nowTokenStatistic.LastInCheckId || nowOutId > nowTokenStatistic.LastOutCheckId {
		tokenStatistics := make([]*models.TokenStatistic, 0)
		err = this.dao.GetTokenStatistics(&tokenStatistics)
		if err != nil {
			return fmt.Errorf("Failed to GetTokenStatistics %w", err)
		}
		tokenBasicBTC, err := this.dao.GetBTCPrice()
		if err != nil {
			return fmt.Errorf("Failed to GetBTCPrice %w", err)
		}
		BTCPrice := decimal.NewFromInt(tokenBasicBTC.Price).Div(decimal.NewFromInt(basedef.PRICE_PRECISION))
		logs.Info("BTCPrice:", BTCPrice)
		for _, statistic := range tokenStatistics {
			for _, in := range inTokenStatistics {
				if statistic.ChainId == in.ChainId && statistic.Hash == in.Hash {
					price_new := decimal.New(in.Token.TokenBasic.Price, 0).Div(decimal.NewFromInt(basedef.PRICE_PRECISION))
					amount_new := decimal.NewFromBigInt(&in.InAmount.Int, 0)
					precision_new := decimal.New(int64(1), int32(in.Token.Precision))
					logs.Info("precision_new in", precision_new)
					amount_usd := amount_new.Div(precision_new).Mul(price_new)
					amount_btc := amount_new.Div(precision_new).Mul(price_new).Div(BTCPrice)
					statistic.InAmount = addDecimalBigInt(statistic.InAmount, models.NewBigInt(amount_new.Div(precision_new).Mul(decimal.NewFromInt32(100)).BigInt()))
					statistic.InAmountUsd = addDecimalBigInt(statistic.InAmountUsd, models.NewBigInt(amount_usd.Mul(decimal.NewFromInt32(10000)).BigInt()))
					statistic.InAmountBtc = addDecimalBigInt(statistic.InAmountBtc, models.NewBigInt(amount_btc.Mul(decimal.NewFromInt32(10000)).BigInt()))
					statistic.InCounter = addDecimalInt64(statistic.InCounter, in.InCounter)
				}
			}
			for _, out := range outTokenStatistics {
				if statistic.ChainId == out.ChainId && statistic.Hash == out.Hash {
					price_new := decimal.New(out.Token.TokenBasic.Price, 0).Div(decimal.NewFromInt(basedef.PRICE_PRECISION))
					amount_new := decimal.NewFromBigInt(&out.OutAmount.Int, 0)
					if out.Token.Precision == uint64(0) {
						out.Token.Precision = uint64(1)
					}
					precision_new := decimal.New(int64(1), int32(out.Token.Precision))
					logs.Info("precision_new out", precision_new)

					amount_usd := amount_new.Div(precision_new).Mul(price_new)
					amount_btc := amount_new.Div(precision_new).Mul(price_new).Div(BTCPrice)

					statistic.OutAmount = addDecimalBigInt(statistic.OutAmount, models.NewBigInt(amount_new.Div(precision_new).Mul(decimal.NewFromInt32(100)).BigInt()))
					statistic.OutCounter = addDecimalInt64(statistic.OutCounter, out.OutCounter)
					statistic.OutAmountUsd = addDecimalBigInt(statistic.OutAmountUsd, models.NewBigInt(amount_usd.Mul(decimal.NewFromInt32(10000)).BigInt()))
					statistic.OutAmountBtc = addDecimalBigInt(statistic.OutAmountBtc, models.NewBigInt(amount_btc.Mul(decimal.NewFromInt32(10000)).BigInt()))
				}

			}
			statistic.LastInCheckId = nowInId
			statistic.LastOutCheckId = nowOutId
			jsonstatistic, _ := json.Marshal(statistic)
			logs.Info("jsonstatistic:", jsonstatistic)
			err = this.dao.SaveTokenStatistic(statistic)
			if err != nil {
				return fmt.Errorf("Failed to SaveTokenStatistic %w", err)
			}
		}

	}
	return
}

func (this *Stats) computeChainStatistics() (err error) {
	logs.Info("computeChainStatistics,start_computeChainStatistics_computeChainStatistics")
	nowChainStatistic, err := this.dao.GetNewChainSta()
	if err != nil {
		return fmt.Errorf("Failed to GetNewChainSta %w", err)
	}
	nowIn, err := this.dao.GetNewDstTransaction()
	if err != nil {
		return fmt.Errorf("Failed to GetNewDstTransfer %w", err)
	}
	nowInId := nowIn.Id
	nowOut, err := this.dao.GetNewSrcTransaction()
	if err != nil {
		return fmt.Errorf("Failed to GetNewSrcTransfer %w", err)
	}
	nowOutId := nowOut.Id
	inChainStatistics := make([]*models.ChainStatistic, 0)
	if nowInId > nowChainStatistic.LastInCheckId {
		err = this.dao.CalculateInChainStatistics(nowChainStatistic.LastInCheckId, nowInId, &inChainStatistics)
		if err != nil {
			logs.Error("Failed to CalculateInTokenStatistics %w", err)
		}
	}
	outChainStatistics := make([]*models.ChainStatistic, 0)
	if nowOutId > nowChainStatistic.LastOutCheckId {
		err = this.dao.CalculateOutChainStatistics(nowChainStatistic.LastOutCheckId, nowOutId, &outChainStatistics)
		if err != nil {
			logs.Error("Failed to CalculateInTokenStatistics %w", err)
		}
	}
	polyTransaction, err := this.dao.GetPolyTransaction()
	if err != nil {
		return fmt.Errorf("Failed to GetPolyTransaction %w", err)
	}
	polyCheckId := polyTransaction.Id
	if nowInId > nowChainStatistic.LastInCheckId || nowOutId > nowChainStatistic.LastOutCheckId {
		logs.Info("computeChainStatistics,nowInId > nowChainStatistic.LastInCheckId || nowOutId > nowChainStatistic.LastOutCheckId")
		chainStatistics := make([]*models.ChainStatistic, 0)
		err = this.dao.GetChainStatistic(&chainStatistics)
		if err != nil {
			return fmt.Errorf("Failed to CalculateInTokenStatistics %w", err)
		}
		for _, chainStatistic := range chainStatistics {
			for _, in := range inChainStatistics {
				if chainStatistic.ChainId == in.ChainId && chainStatistic.ChainId != basedef.POLY_CROSSCHAIN_ID {
					chainStatistic.In = addDecimalInt64(chainStatistic.In, in.In)
					break
				}
			}
			for _, out := range outChainStatistics {
				if chainStatistic.ChainId == out.ChainId && chainStatistic.ChainId != basedef.POLY_CROSSCHAIN_ID {
					chainStatistic.Out = addDecimalInt64(chainStatistic.Out, out.Out)
					break
				}
			}
			if chainStatistic.ChainId != basedef.POLY_CROSSCHAIN_ID {
				chainStatistic.LastInCheckId = nowInId
				chainStatistic.LastOutCheckId = nowOutId
			}
		}
		logs.Info("computeChainStatistics,poly_polyCheckId", polyCheckId)
		for _, chainStatistic := range chainStatistics {
			if chainStatistic.ChainId == basedef.POLY_CROSSCHAIN_ID {
				counter, err := this.dao.CalculatePolyChainStatistic(chainStatistic.LastInCheckId, polyCheckId)
				if err == nil {
					logs.Info("computeChainStatistics,polychainid:", chainStatistic.ChainId, "poly.In:", chainStatistic.In, "poly.Out:", chainStatistic.Out, "poly.LastInCheckId1", chainStatistic.LastInCheckId, "polycounter:", counter)
					chainStatistic.In = addDecimalInt64(counter, chainStatistic.In)
					chainStatistic.Out = addDecimalInt64(counter, chainStatistic.Out)
					chainStatistic.LastInCheckId = polyCheckId
					chainStatistic.LastOutCheckId = polyCheckId
				}
				break
			}
		}
		for _, v := range chainStatistics {
			fmt.Println(*v)
		}
		err = this.dao.SaveChainStatistics(chainStatistics)
		if err != nil {
			logs.Error("qChainStatistic,computeChainStatisticAssets SaveChainStatistic error", err)
		}
	}
	return
}
func (this *Stats) computeChainStatisticAssets() (err error) {
	logs.Info("computeChainStatisticAssets,start computeChainStatisticAssets")
	computeChainStatistics := make([]*models.ChainStatistic, 0)
	err = this.dao.CalculateChainStatisticAssets(&computeChainStatistics)
	if err != nil {
		return fmt.Errorf("Failed to CalculateChainStatisticAssets %w", err)
	}
	chainStatistics := make([]*models.ChainStatistic, 0)
	err = this.dao.GetChainStatistic(&chainStatistics)
	if err != nil {
		logs.Error("computeChainStatisticAssets GetChainStatistic error", err)
	}
	polyAddress := int64(0)
	for _, chainStatistic := range chainStatistics {
		for _, chain := range computeChainStatistics {
			if chainStatistic.ChainId == chain.ChainId && chainStatistic.ChainId != basedef.POLY_CROSSCHAIN_ID {
				chainStatistic.Addresses = chain.Addresses
				polyAddress += chain.Addresses
				break
			}
		}
	}
	for _, chainStatistic := range chainStatistics {
		if chainStatistic.ChainId == basedef.POLY_CROSSCHAIN_ID {
			chainStatistic.Addresses = polyAddress
		}
	}
	err = this.dao.SaveChainStatistics(chainStatistics)
	if err != nil {
		logs.Error("computeChainStatisticAssets SaveChainStatistics error", err)
	}
	logs.Info("computeChainStatisticAssets,end computeChainStatisticAssets")

	return
}

func (this *Stats) computeAssetStatistics() (err error) {
	logs.Info("computeAssetStatistics,start computeAssetStatistics")
	nowAssetStatistic, err := this.dao.GetNewAssetSta()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("Failed to GetNewAssetSta %w", err)
	}
	srcTransfer, err := this.dao.GetNewSrcTransfer()
	if err != nil {
		return fmt.Errorf("Failed to GetNewSrcTransfer %w", err)
	}
	nowId := srcTransfer.Id
	if nowAssetStatistic.LastCheckId >= nowId {
		return nil
	}
	newAssets, err := this.dao.CalculateAsset(nowAssetStatistic.LastCheckId, nowId)
	if err != nil {
		return fmt.Errorf("Failed to CalculateAsset %w", err)
	}
	tokenBasicBTC, err := this.dao.GetBTCPrice()
	if err != nil {
		return fmt.Errorf("Failed to GetBTCPrice %w", err)
	}
	BTCPrice := decimal.NewFromInt(tokenBasicBTC.Price).Div(decimal.NewFromInt(basedef.PRICE_PRECISION))
	oldAssetStatistics, err := this.dao.GetAssetStatistic()
	if err != nil {
		return fmt.Errorf("Failed to GetAssetStatistic %w", err)
	}
	for _, old := range oldAssetStatistics {
		for _, assetInfo := range newAssets {
			if old.TokenBasicName == assetInfo.TokenBasicName {
				amount_new := decimal.NewFromBigInt(&assetInfo.Amount.Int, 0)
				price_new := decimal.NewFromInt(assetInfo.Price)
				if assetInfo.Price == int64(0) {
					price_new = decimal.NewFromInt32(1)
				} else {
					price_new = price_new.Div(decimal.NewFromInt(basedef.PRICE_PRECISION))
				}
				precision_new := decimal.New(int64(1), int32(assetInfo.Precision))
				real_amount := amount_new.Div(precision_new)
				amount_usd := real_amount.Mul(price_new)
				amount_btc := amount_usd.Div(BTCPrice)

				old.Amount = models.NewBigInt((real_amount.Mul(decimal.New(int64(100), 0)).Add(decimal.NewFromBigInt(&old.Amount.Int, 0))).BigInt())
				old.AmountUsd = models.NewBigInt((amount_usd.Mul(decimal.New(int64(10000), 0)).Add(decimal.NewFromBigInt(&old.AmountUsd.Int, 0))).BigInt())
				old.AmountBtc = models.NewBigInt((amount_btc.Mul(decimal.New(int64(10000), 0)).Add(decimal.NewFromBigInt(&old.AmountBtc.Int, 0))).BigInt())
			}
		}
		old.LastCheckId = nowId
		err := this.dao.SaveAssetStatistic(old)
		if err != nil {
			return fmt.Errorf("Failed to UpdateTransferStatistic %w", err)
		}
	}
	logs.Info("computeAssetStatistics,end computeAssetStatistics")

	return nil
}

func (this *Stats) computeAssetStatisticAdress() (err error) {
	logs.Info("start computeAssetStatisticAdress")
	newAssetAdresses, err := this.dao.CalculateAssetAdress()
	if err != nil {
		return fmt.Errorf("Failed to CalculateAssetAdress %w", err)
	}
	for _, assetStatistic := range newAssetAdresses {
		err := this.dao.UpdateAssetStatisticAdress(assetStatistic)
		if err != nil {
			return fmt.Errorf("Failed to UpdateAssetStatisticAdress %w", err)
		}
	}
	return nil
}

func addDecimalBigInt(a, b *models.BigInt) *models.BigInt {
	a_new := decimal.NewFromBigInt(&a.Int, 0)
	b_new := decimal.NewFromBigInt(&b.Int, 0)
	c := a_new.Add(b_new)
	return models.NewBigInt(c.BigInt())
}
func addDecimalInt64(a, b int64) int64 {
	a_new := decimal.New(a, 0)
	b_new := decimal.New(b, 0)
	c := a_new.Add(b_new)
	return c.IntPart()
}
