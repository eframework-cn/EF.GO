//-----------------------------------------------------------------------//
//                     GNU GENERAL PUBLIC LICENSE                        //
//                        Version 2, June 1991                           //
//                                                                       //
// Copyright (C) EFramework, https://eframework.cn, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies          //
// of this license document, but changing it is not allowed.             //
//                   SEE LICENSE.md FOR MORE DETAILS.                    //
//-----------------------------------------------------------------------//

package app

import (
	_ "eframe/src/private/center/ctx"
	"eframe/src/shared/model/mmn"

	"github.com/eframework-cn/EP.GO.CORE/xserver"
)

func init() {
	mmn.RegPlayer().LCache(true).LDB(true).LRedis(true).LRW(true)
}

type CenterServer struct {
	xserver.Server
}

func NewCenterServer() *CenterServer {
	this := &CenterServer{}
	this.CTOR(this)
	return this
}
