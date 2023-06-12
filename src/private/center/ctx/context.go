//-----------------------------------------------------------------------//
//                     GNU GENERAL PUBLIC LICENSE                        //
//                        Version 2, June 1991                           //
//                                                                       //
// Copyright (C) EFramework, https://eframework.cn, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies          //
// of this license document, but changing it is not allowed.             //
//                   SEE LICENSE.md FOR MORE DETAILS.                    //
//-----------------------------------------------------------------------//

package ctx

import (
	"eframe/src/shared/model/mmn"
	"eframe/src/shared/proto/rpb"

	"github.com/eframework-cn/EP.GO.CORE/xorm"
	"github.com/eframework-cn/EP.GO.CORE/xserver"
	"github.com/eframework-cn/EP.GO.CORE/xsession"
)

func init() {
	xserver.RegEvt(xserver.EVT_SERVER_STARTED, func(param interface{}) {
		xsession.GList(mmn.NewPlayer())

		players := ListPlayer(true)
		for _, player := range players {
			player.Online = 0
			player.ConnUrl = ""
		}
		conns := xserver.GLan.SelectAll("conn")
		if conns != nil {
			for _, conn := range conns {
				resp := &rpb.RPC_GetOnlineFromConnResp{}
				xserver.SendSync(int(rpb.RID.RPC_GET_ONLINE_FROM_CONN), 0, nil, resp, conn.ServerID())
				for idx, id := range resp.ID {
					uid := int(id)
					url := resp.Url[idx]
					cid := resp.CID[idx]
					player := ReadPlayer(uid)
					if player != nil {
						player.Online = 1
						player.ConnUrl = url
						player.ConnID = cid
					}
				}
			}
		}
	})
}

func ReadPlayer(idOrCond ...interface{}) *Player {
	player := mmn.NewPlayer()
	if len(idOrCond) == 1 {
		switch idOrCond[0].(type) {
		case int:
			player.ID = idOrCond[0].(int)
			player = xsession.GRead(player).(*mmn.Player)
			break
		case *xorm.Condition:
			player = xsession.GRead(player, idOrCond[0].(*xorm.Condition)).(*mmn.Player)
			break
		}
	}
	if player.IsValid() {
		return NewPlayer(player)
	} else {
		return nil
	}
}

func ListPlayer(rw bool, cond ...*xorm.Condition) []*Player {
	player := mmn.NewPlayer()
	player.RW(rw)
	players := *xsession.GList(player, cond...).(*[]*mmn.Player)
	fplayers := []*Player{}
	for _, temp := range players {
		fplayers = append(fplayers, NewPlayer(temp))
	}
	return fplayers
}
