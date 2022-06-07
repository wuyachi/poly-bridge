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

package chainfeedao

import (
	"github.com/polynetwork/poly-bridge/basedef"
	"github.com/polynetwork/poly-bridge/chainfeedao/bridgedao"
	"github.com/polynetwork/poly-bridge/chainfeedao/stakedao"
	"github.com/polynetwork/poly-bridge/chainfeedao/swapdao"
	"github.com/polynetwork/poly-bridge/conf"
	"github.com/polynetwork/poly-bridge/models"
)

type ChainFeeDao interface {
	GetFees() ([]*models.ChainFee, error)
	SaveFees(fees []*models.ChainFee) error
	Name() string
}

func NewChainFeeDao(server string, dbCfg *conf.DBConfig) ChainFeeDao {
	if server == basedef.SERVER_POLY_BRIDGE {
		return bridgedao.NewBridgeDao(dbCfg)
	} else if server == basedef.SERVER_STAKE {
		return stakedao.NewStakeDao()
	} else if server == basedef.SERVER_POLY_SWAP {
		return swapdao.NewSwapDao(dbCfg)
	} else {
		return nil
	}
}
