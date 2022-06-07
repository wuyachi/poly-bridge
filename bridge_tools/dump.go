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

package main

import (
	"fmt"
	"github.com/polynetwork/poly-bridge/bridge_tools/conf"
	serverconf "github.com/polynetwork/poly-bridge/conf"
	"github.com/polynetwork/poly-bridge/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func dumpStatus(dbCfg *serverconf.DBConfig) {
	Logger := logger.Default
	if dbCfg.Debug == true {
		Logger = Logger.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{Logger: Logger})
	if err != nil {
		panic(err)
	}
	{
		chains := make([]*models.Chain, 0)
		db.Find(&chains)
		fmt.Printf("chain info:\nchainid\t\t\t\theight\t\t\t\t\n")
		for _, chain := range chains {
			fmt.Printf("%d\t\t\t\t%d\t\t\t\t\n", chain.ChainId, chain.Height)
		}
	}
	{
		tokens := make([]*models.Token, 0)
		db.Find(&tokens)
		fmt.Printf("token info:\nChainId\t\t\t\tHash\t\t\t\tName\t\t\t\tTokenBasicName\t\t\t\t\n")
		for _, token := range tokens {
			fmt.Printf("%d\t\t\t\t%s\t\t\t\t%s\t\t\t\t%s\t\t\t\t\n",
				token.ChainId, token.Hash, token.Name, token.TokenBasicName)
		}
	}
	{
		tokenBasics := make([]*models.TokenBasic, 0)
		db.Find(&tokenBasics)
		fmt.Printf("token basic info:\nName\t\t\t\tAvgPrice\t\t\t\tAvgInd\t\t\t\tTime\t\t\t\t\n")
		for _, tokenBasic := range tokenBasics {
			fmt.Printf("%s\t\t\t\t%d\t\t\t\t%d\t\t\t\t%d\t\t\t\t\n",
				tokenBasic.Name, tokenBasic.Price, tokenBasic.Ind,
				tokenBasic.Time)
		}
	}
	{
		TokenMaps := make([]*models.TokenMap, 0)
		db.Find(&TokenMaps)
		fmt.Printf("token map info:\nSrcChain\t\t\t\tSrcTokenHash\t\t\t\tDstChain\t\t\t\tDstTokenHash\t\t\t\t\n")
		for _, TokenMap := range TokenMaps {
			fmt.Printf("%d\t\t\t\t%s\t\t\t\t%d\t\t\t\t%s\t\t\t\t\n", TokenMap.SrcChainId, TokenMap.SrcTokenHash, TokenMap.DstChainId, TokenMap.DstTokenHash)
		}
	}
	{
		chainFees := make([]*models.ChainFee, 0)
		db.Find(&chainFees)
		fmt.Printf("chain fee info:\nchainid\t\t\t\tTokenBasicName\t\t\t\tmaxfee\t\t\t\tminfee\t\t\t\tproxyfee\t\t\t\t\n")
		for _, chainFee := range chainFees {
			fmt.Printf("%d\t\t\t\t%s\t\t\t\t%d\t\t\t\t%d\t\t\t\t%d\t\t\t\t\n",
				chainFee.ChainId, chainFee.TokenBasicName, chainFee.MaxFee, chainFee.MinFee, chainFee.ProxyFee)
		}
	}
}

func dumpAffectedRows(cfg *conf.UpdateConfig, dbCfg *serverconf.DBConfig) {
	Logger := logger.Default
	if dbCfg.Debug == true {
		Logger = Logger.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{Logger: Logger})
	if err != nil {
		panic(err)
	}
	{
		chainIds := make([]uint64, 0)
		for _, chain := range cfg.Chains {
			chainIds = append(chainIds, chain.ChainId)

		}
		chains := make([]*models.Chain, 0)
		db.Where("chain_id in ?", chainIds).Find(&chains)
		fmt.Printf("chain info:\nchainid\t\t\t\theight\t\t\t\t\n")
		for _, chain := range chains {
			fmt.Printf("%d\t\t\t\t%d\t\t\t\t\n", chain.ChainId, chain.Height)
		}
	}
	{
		chainIds := make([]uint64, 0)
		for _, chain := range cfg.ChainFees {
			chainIds = append(chainIds, chain.ChainId)

		}
		chainFees := make([]*models.ChainFee, 0)
		db.Where("chain_id in ?", chainIds).Find(&chainFees)
		fmt.Printf("chain fee info:\nchainid\t\t\t\tTokenBasicName\t\t\t\tmaxfee\t\t\t\tminfee\t\t\t\tproxyfee\t\t\t\t\n")
		for _, chainFee := range chainFees {
			fmt.Printf("%d\t\t\t\t%s\t\t\t\t%d\t\t\t\t%d\t\t\t\t%d\t\t\t\t\n",
				chainFee.ChainId, chainFee.TokenBasicName, chainFee.MaxFee, chainFee.MinFee, chainFee.ProxyFee)
		}
	}
	{
		tokenBasicNames := make([]string, 0)
		for _, basic := range cfg.TokenBasics {
			tokenBasicNames = append(tokenBasicNames, basic.Name)
		}
		tokenBasics := make([]*models.TokenBasic, 0)
		db.Where("name in ?", tokenBasicNames).Find(&tokenBasics)
		fmt.Printf("token basic info:\nName\t\t\t\tAvgPrice\t\t\t\tAvgInd\t\t\t\tTime\t\t\t\t\n")
		for _, tokenBasic := range tokenBasics {
			fmt.Printf("%s\t\t\t\t%d\t\t\t\t%d\t\t\t\t%d\t\t\t\t\n",
				tokenBasic.Name, tokenBasic.Price, tokenBasic.Ind,
				tokenBasic.Time)
		}
	}
	{
		tokenHashes := make([]string, 0)
		for _, basic := range cfg.TokenBasics {
			for _, token := range basic.Tokens {
				tokenHashes = append(tokenHashes, token.Hash)
			}
		}
		tokens := make([]*models.Token, 0)
		db.Where("hash in ?", tokenHashes).Find(&tokens)
		fmt.Printf("token info:\nChainId\t\t\t\tHash\t\t\t\tName\t\t\t\tTokenBasicName\t\t\t\t\n")
		for _, token := range tokens {
			fmt.Printf("%d\t\t\t\t%s\t\t\t\t%s\t\t\t\t%s\t\t\t\t\n",
				token.ChainId, token.Hash, token.Name, token.TokenBasicName)
		}
	}
	{
		fmt.Printf("token map info:\nSrcChain\t\t\t\tSrcTokenHash\t\t\t\tDstChain\t\t\t\tDstTokenHash\t\t\t\t\n")
		for _, basic := range cfg.TokenBasics {
			tokenChainIds := make([]uint64, 0)
			tokenHashes := make([]string, 0)
			for _, token := range basic.Tokens {
				tokenHashes = append(tokenHashes, token.Hash)
				tokenChainIds = append(tokenChainIds, token.ChainId)
			}
			TokenMaps := make([]*models.TokenMap, 0)
			db.Where("(src_token_hash in ? and src_chain_id in ?) or (dst_token_hash in ? and dst_chain_id in ?)", tokenHashes, tokenChainIds, tokenHashes, tokenChainIds).Find(&TokenMaps)
			for _, TokenMap := range TokenMaps {
				fmt.Printf("%d\t\t\t\t%s\t\t\t\t%d\t\t\t\t%s\t\t\t\t\n", TokenMap.SrcChainId, TokenMap.SrcTokenHash, TokenMap.DstChainId, TokenMap.DstTokenHash)
			}
		}
	}

}
