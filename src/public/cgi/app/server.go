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
	"net/http"

	"github.com/eframework-cn/EP.GO.CORE/xhttp"
	"github.com/eframework-cn/EP.GO.CORE/xproto"
	"github.com/eframework-cn/EP.GO.CORE/xserver"
	"github.com/eframework-cn/EP.GO.UTIL/xjson"
	"github.com/eframework-cn/EP.GO.UTIL/xrun"
	"github.com/eframework-cn/EP.GO.UTIL/xstring"
)

type CgiServer struct {
	xserver.Server
	Svr *xhttp.Server

	Address string // 前端连接地址
	Key     string // Https密钥
	Cert    string // Https证书
}

func NewCgiServer() *CgiServer {
	this := &CgiServer{}
	this.CTOR(this)
	xserver.RegEvt(xserver.EVT_SERVER_STARTED, func(param interface{}) {
		this.Svr = xhttp.NewServer().
			SetAddr(this.Address).SetHttps(this.Key, this.Cert).
			SetHandler(func(resp http.ResponseWriter, req *http.Request) {
				method := req.Method
				if method == "OPTIONS" { // preflight request
					resp.Header().Add("Access-Control-Allow-Origin", "*")
					resp.Header().Add("Access-Control-Allow-Headers", "*")
					resp.Header().Add("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")
					resp.WriteHeader(200)
				} else {
					cid := -1
					_cid := req.Header.Get("CID")
					if _cid == "" {
						_cid = req.URL.Query().Get("CID")
					}
					if _cid == "" {
						_cid = req.Header.Get("cid")
					}
					if _cid == "" {
						_cid = req.URL.Query().Get("cid")
					}
					cid = xstring.ToInt(_cid)
					route := xserver.CGIROUTEMAP[cid]
					if route == nil {
						resp.WriteHeader(500)
						resp.Write(xstring.StrToBytes(fmt.Sprintf("no route for cid %v", cid)))
					} else {
						sig := false
						if len(route.Method) > 0 {
							for i := 0; i < len(route.Method); i++ {
								if route.Method[i] == method {
									sig = true
									break
								}
							}
						} else {
							sig = true
						}
						if sig == false {
							resp.WriteHeader(501)
							resp.Write(xstring.StrToBytes(fmt.Sprintf("invalid method %v, path %v", req.Method, req.URL.Path)))
						} else {
							lan := xserver.GLan.SelectRand(route.Dst[0])
							if lan == nil {
								resp.WriteHeader(502)
								resp.Write(xstring.StrToBytes(fmt.Sprintf("no lan for route %v, path %v", route.Dst[0], req.URL.Path)))
							} else {
								if cresp, err := xserver.SendCgi(route.ID, 0, req, lan.ServerID(), route.Timeout); err != nil {
									resp.WriteHeader(503)
									resp.Write(xstring.StrToBytes(err.Error()))
								} else {
									defer xproto.PoolFrame(cresp)
									header := make(map[string]string)
									xjson.ToObj(cresp.Header, &header)
									for k := range header {
										resp.Header().Set(k, header[k])
									}
									resp.Header().Add("Access-Control-Allow-Origin", "*")
									resp.WriteHeader(int(cresp.GetStatus()))
									resp.Write(cresp.GetBody())
								}
							}
						}
					}
				}
			})
		go xrun.Exec(func() {
			this.Svr.Start()
		})
	})
	return this
}

func (this *CgiServer) InitConfig() bool {
	if this.Server.InitConfig() == false {
		return false
	}
	config := this.GetConfig()
	this.Address = config.Raw.String("client::addr")
	this.Key = config.Raw.DefaultString("client::key", "")
	this.Cert = config.Raw.DefaultString("client::cert", "")
	return true
}
