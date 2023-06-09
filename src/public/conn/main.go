//-----------------------------------------------------------------------//
//                     GNU GENERAL PUBLIC LICENSE                        //
//                        Version 2, June 1991                           //
//                                                                       //
// Copyright (C) EFramework, https://eframework.cn, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies          //
// of this license document, but changing it is not allowed.             //
//                   SEE LICENSE.md FOR MORE DETAILS.                    //
//-----------------------------------------------------------------------//

package main

import (
	"github.com/eframework-cn/EP.GO.CORE/xserver"

	_ "github.com/eframework-cn/EF.GO/src/shared/proto"

	_ "github.com/eframework-cn/EP.GO.UTIL/xlog"
	"github.com/eframework-cn/EP.GO.UTIL/xrun"

	"github.com/eframework-cn/EF.GO/src/public/conn/app"
)

func main() {
	defer xrun.Caught(true)

	xserver.Start(app.NewConnServer())
}
