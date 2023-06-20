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
	"fmt"

	"github.com/eframework-cn/EP.GO.CORE/xconn"
	"github.com/eframework-cn/EP.GO.CORE/xproto"
	"github.com/eframework-cn/EP.GO.CORE/xserver"
	"github.com/eframework-cn/EP.GO.UTIL/xjson"
	"github.com/eframework-cn/EP.GO.UTIL/xlog"

	"eframe/src/shared/proto/mpb"
	"eframe/src/shared/proto/rpb"

	"github.com/golang/protobuf/proto"
)

type ConnCfg struct {
	Addr    string `json:"addr"`                   // 前端连接地址
	MaxConn int    `json:"maxConn" default:"5000"` // 最大连接数（超过该值将不再接收新的连接）（默认5000）
	MaxLoad int    `json:"maxLoad" default:"100"`  // 单连接每秒最大的请求数（超过该值则视为DDOS，主动断开连接）（默认100QPS）
	MaxBody int    `json:"maxBody" default:"1024"` // 单连接最大的消息体（不包含头部）（默认1024KB）
	Timeout int    `json:"timeout" default:"15"`   // 连接超时时间（超过该值则主动断开连接）（默认15秒）
	WS      bool   `json:"ws" default:"false"`     // 是否面向WebSocket连接（默认false）
	Key     string `json:"key"`                    // Https密钥
	Cert    string `json:"cert"`                   // Https证书
}

type ConnClient struct {
	xconn.Client
	UID int
}

func NewConnClient(server *xconn.Server) *ConnClient {
	this := new(ConnClient)
	this.UID = -1
	this.Server = server
	this.CTOR(this)
	return this
}

type ConnServer struct {
	xserver.Server
	Svr     *xconn.Server
	CCfg    *ConnCfg
	BeatMsg []byte
}

func NewConnServer() *ConnServer {
	this := &ConnServer{}
	this.CTOR(this)
	this.BeatMsg, _ = xproto.PackMsg(int(mpb.MID.GM_HEART_BEAT), &mpb.GM_Common{Result: proto.Int(0)})
	xserver.RegEvt(xserver.EVT_SERVER_STARTED, func(param interface{}) {
		this.Svr = xconn.NewServer().
			SetAddr(this.CCfg.Addr).SetIsWS(this.CCfg.WS, this.CCfg.Key, this.CCfg.Cert).
			SetMaxConn(this.CCfg.MaxConn).SetMaxLoad(this.CCfg.MaxLoad).
			SetMaxBody(this.CCfg.MaxBody * 1024).SetTimeout(this.CCfg.Timeout).
			SetOnAccept(func(client xconn.IClient) {
			}).
			SetOnRemove(func(client xconn.IClient) {
				gclient := client.(*ConnClient)
				this.RemoveClient(gclient)
			}).
			SetOnReceive(func(client xconn.IClient, bytes []byte) {
				gclient := client.(*ConnClient)
				this.FromClient(gclient, bytes)
			}).
			SetNewClient(func() interface{} {
				return NewConnClient(this.Svr)
			}).
			Start()
	})
	xserver.RegRpc(int(rpb.RID.RPC_GET_ONLINE_FROM_CONN), func(rreq *xproto.RpcReq, rresp *xproto.RpcResp) {
		resp := &rpb.RPC_GetOnlineFromConnResp{}
		defer func() {
			b, _ := proto.Marshal(resp)
			rresp.Data = b
		}()
		this.Svr.Clients.Range(func(k, v interface{}) bool {
			client := v.(*ConnClient)
			if client.UID != -1 {
				resp.ID = append(resp.ID, int32(client.UID))
				resp.Url = append(resp.Url, xproto.GUrl)
				resp.CID = append(resp.CID, int64(client.ID))
			}
			return true
		})
	})
	return this
}

func (this *ConnServer) Init(cfg *xserver.SvrCfg) bool {
	var lcfg map[string]interface{}
	if e := xjson.ToObj(cfg.Raw, &lcfg); e != nil {
		xlog.Panic("ConnServer.Init: readout config error: ", e)
		return false
	}
	if ccfg, _ := lcfg["client"]; ccfg == nil {
		xlog.Panic("ConnServer.Init: readout config error: nil client section")
		return false
	} else {
		d, e := xjson.ToByte(ccfg)
		if e != nil {
			xlog.Panic("ConnServer.Init: readout config error: ", e)
			return false
		}
		this.CCfg = new(ConnCfg)
		if e := xjson.ToObj(d, this.CCfg); e != nil {
			xlog.Panic("ConnServer.Init: readout config error: ", e)
			return false
		}
		return this.Server.Init(cfg)
	}
}

func (this *ConnServer) SetTitle() string {
	title := this.Server.SetTitle()
	if this.Svr != nil {
		title += fmt.Sprintf("[CON-%v]", this.Svr.OnlineNum)
	}
	return title
}

func (this *ConnServer) RecvMsg(mreq *xproto.MsgReq) {
	this.ToClient(mreq)
}

func (this *ConnServer) ToClient(mreq *xproto.MsgReq) {
	url := mreq.Route.GetDst()
	id := int(mreq.Route.GetMID1())
	client := this.Svr.GetClient(id)
	if client == nil {
		xlog.Warn("ctx.ConnServer.ToClient: no client-%v found, uid=%v, msgId=%v, dst=%v, src=%v",
			id, mreq.Route.GetUID(), xproto.UnpackID(mreq.Data), url, mreq.Route.GetSrc())
		return
	}
	this.FromServer(client.(*ConnClient), mreq)
}

func (this *ConnServer) ToServer(client *ConnClient, rid int, dst string, bytes []byte, gol int, gor int) bool {
	mreq := xproto.GetMsgReq()
	mreq.Route.Src = xproto.PGUrl
	mreq.Route.Dst = proto.String(dst)
	mreq.Route.RID = proto.Int(rid)
	mreq.Route.UID = proto.Int(client.UID)
	mreq.Route.GID = proto.Int(xproto.PackGID(gol, gor))
	mreq.Route.MID1 = proto.Int64(int64(client.ID))
	mreq.Data = bytes
	return xserver.SendFrame(mreq)
}

func (this *ConnServer) RemoveClient(client *ConnClient) {
	if client.UID != -1 {
		center := xserver.GLan.SelectRand("center")
		if center != nil {
			req := &rpb.RPC_ConnNotifyOfflineReq{}
			req.UID = proto.Int(client.UID)
			req.Url = xproto.PGUrl
			req.CID = proto.Int64(int64(client.ID))
			xserver.SendAsync(int(rpb.RID.RPC_CONN_NOTIFY_OFFLINE), client.UID, req, center.ID, nil)
		}
		client.UID = -1
	}
}

func (this *ConnServer) FromClient(client *ConnClient, bytes []byte) {
	rid := xproto.UnpackID(bytes)
	eid := mpb.MIDEnum(rid)
	if eid == mpb.MID.GM_HEART_BEAT {
		client.Write(this.BeatMsg)
	} else {
		route := xserver.MSGROUTEMAP[rid]
		if route == nil {
			xlog.Warn("ctx.ConnServer.FromClient: no route for id-%v found", rid)
		} else {
			for _, dst := range route.Dst {
				if dst == "client" {
					continue
				}
				conn := xserver.GLan.SelectRand(dst)
				if conn == nil {
					xlog.Warn("ctx.ConnServer.FromClient: select conn failed, uid=%v, id=%v, tag=%v", client.UID, rid, dst)
				} else {
					this.ToServer(client, rid, conn.ID, bytes, route.GoL, route.GoR)
				}
			}
		}
	}
}

func (this *ConnServer) FromServer(client *ConnClient, mreq *xproto.MsgReq) {
	id := xproto.UnpackID(mreq.Data)
	if id == -1 {
		xlog.Warn("ctx.ConnServer.FromServer: parse msg id failed, uid=%v", client.UID)
		return
	}
	eid := mpb.MIDEnum(id)
	xlog.Info("ctx.ConnServer.FromServer: client-%v, msgid-%v, size-%v", client.UID, id, len(mreq.Data))
	if eid == mpb.MID.GM_LOGIN_RESPONSE && client.UID == -1 {
		client.UID = int(mreq.Route.GetUID())
		xlog.Notice("ctx.ConnServer.FromServer: client-%v was binded to uid-%v", client.ID, client.UID)
	}
	if eid == mpb.MID.GM_KICK_OFF && client.UID != -1 {
		xlog.Notice("ctx.ConnServer.FromServer: client-%v was unbinded from uid-%v", client.ID, client.UID)
		client.UID = -1
	}
	client.Write(mreq.Data)
}
