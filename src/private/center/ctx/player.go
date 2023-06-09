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
	"github.com/eframework-cn/EF.GO/src/shared/model/mmn"
	"github.com/eframework-cn/EF.GO/src/shared/proto/mpb"
	"github.com/eframework-cn/EF.GO/src/shared/proto/rpb"

	"github.com/eframework-cn/EP.GO.CORE/xproto"
	"github.com/eframework-cn/EP.GO.CORE/xserver"
	"github.com/eframework-cn/EP.GO.UTIL/xlog"

	"github.com/golang/protobuf/proto"
)

type Player struct {
	*mmn.Player
}

func NewPlayer(md ...*mmn.Player) *Player {
	this := new(Player)
	if len(md) == 1 && md[0] != nil {
		this.Player = md[0]
	} else {
		this.Player = mmn.NewPlayer()
	}
	return this
}

func (this *Player) Send(id mpb.MIDEnum, data proto.Message, mreq ...*xproto.MsgReq) {
	var nmreq *xproto.MsgReq
	if len(mreq) == 1 {
		nmreq = mreq[0]
	} else {
		nmreq = xproto.GetMsgReq()
		nmreq.Route.UID = proto.Int(this.ID)
		nmreq.Route.RID = proto.Int(int(id))
		nmreq.Route.Src = proto.String(xproto.GUrl)
		nmreq.Route.Dst = proto.String(this.ConnUrl)
		nmreq.Route.MID1 = proto.Int64(this.ConnID)
	}
	xserver.SendMsg(int(id), data, nmreq)
}

func (this *Player) OnLogin(req *mpb.GM_LoginReq, mreq *xproto.MsgReq) {
	this.Online = 1
	if this.ConnUrl != mreq.Route.GetSrc() && this.ConnID != mreq.Route.GetMID1() { // 网关
		if this.ConnUrl != "" && this.ConnID != -1 {
			this.Send(mpb.MID.GM_KICK_OFF, &mpb.GM_Common{Result: proto.Int(0)})
		}
	}
	this.ConnUrl = mreq.Route.GetSrc()
	this.ConnID = mreq.Route.GetMID1()
	mreq.Route.UID = proto.Int(this.ID)

	resp := &mpb.GM_LoginResp{}
	resp.ID = proto.Int(this.ID)
	resp.Status = proto.Int(this.Online)
	this.Send(mpb.MID.GM_LOGIN_RESPONSE, resp)
}

func (this *Player) OnLogout(req *rpb.RPC_ConnNotifyOfflineReq, rreq *xproto.RpcReq) {
	if this.ConnUrl != req.GetUrl() && this.ConnID != req.GetCID() {
		xlog.Warn("ctx.Player.OnLogout: player-%v's oldurl-%v#%v doesn't equals to newurl-%v#%v", this.ID, this.ConnUrl, this.ConnID, req.GetUrl(), req.GetCID())
	} else {
		xlog.Info("ctx.Player.OnLogout: player-%v offline, url-%v#%v", this.ID, this.ConnUrl, this.ConnID)
		this.Online = 0
	}
}
