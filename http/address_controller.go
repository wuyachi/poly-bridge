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

package http

import (
	"encoding/json"

	"github.com/polynetwork/poly-bridge/models"

	"github.com/beego/beego/v2/server/web"
)

type AddressController struct {
	web.Controller
}

func (c *AddressController) Address() {
	var addressReq models.AddressReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &addressReq); err != nil {
		panic(err)
	}

	c.Data["json"] = models.MakeAddressRsp(addressReq.AddressHash, addressReq.ChainId, address(addressReq.AddressHash, addressReq.ChainId))
	c.ServeJSON()
}

func address(hash string, chainId uint64) string {
	return hash
}
