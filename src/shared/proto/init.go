package proto

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/eframework-cn/EP.GO.CORE/xserver"

	"eframe/src/shared/proto/cpb"
	"eframe/src/shared/proto/mpb"
	"eframe/src/shared/proto/rpb"
)

func init() {
	initmid()
	initrid()
	initcid()
}

func initmid() {
	_map := make(map[int]*xserver.MsgRoute)
	ttpe := reflect.TypeOf(mpb.MID)
	numlen := ttpe.NumField()
	for i := 0; i < numlen; i++ {
		tfld := ttpe.Field(i)
		_dst := tfld.Tag.Get("to")
		_gid := tfld.Tag.Get("go")
		_rw := tfld.Tag.Get("rw")
		_log := tfld.Tag.Get("log")
		if _dst != "" || _gid != "" || _rw != "" {
			route := &xserver.MsgRoute{}
			route.Name = tfld.Name
			route.Dst = strings.Split(_dst, ",")
			if _gid != "" {
				if strings.Contains(_gid, ":") {
					strs := strings.Split(_gid, ":")
					gid, _ := strconv.Atoi(strs[0])
					route.GoL = gid
					if len(strs) == 1 {
						route.GoR = 0
					} else {
						gid2, _ := strconv.Atoi(strs[1])
						route.GoR = gid2
					}
				} else {
					gid, _ := strconv.Atoi(_gid)
					route.GoL = gid
					route.GoR = gid
				}
			} else {
				route.GoL = 0
				route.GoR = 0
			}
			if _rw != "" {
				route.RW = _rw == "1"
			} else {
				route.RW = true
			}
			route.Log = -1
			if _log != "" {
				tmp, _ := strconv.Atoi(_log)
				if tmp >= 0 && tmp <= 7 {
					route.Log = tmp
				}
			}
			_id := tfld.Tag.Get("id")
			id, _ := strconv.Atoi(_id)
			route.ID = id
			_map[id] = route
		}
	}
	xserver.RegMsgRoute(_map)
}

func initrid() {
	_map := make(map[int]*xserver.RpcRoute)
	ttpe := reflect.TypeOf(rpb.RID)
	numlen := ttpe.NumField()
	for i := 0; i < numlen; i++ {
		tfld := ttpe.Field(i)
		_dst := tfld.Tag.Get("to")
		_gid := tfld.Tag.Get("go")
		_rw := tfld.Tag.Get("rw")
		_log := tfld.Tag.Get("log")
		if _dst != "" || _gid != "" || _rw != "" {
			route := &xserver.RpcRoute{}
			route.Name = tfld.Name
			route.Dst = strings.Split(_dst, ",")
			if _gid != "" {
				if strings.Contains(_gid, ":") {
					strs := strings.Split(_gid, ":")
					gid, _ := strconv.Atoi(strs[0])
					route.GoL = gid
					if len(strs) == 1 {
						route.GoR = 0
					} else {
						gid2, _ := strconv.Atoi(strs[1])
						route.GoR = gid2
					}
				} else {
					gid, _ := strconv.Atoi(_gid)
					route.GoL = gid
					route.GoR = gid
				}
			} else {
				route.GoL = 0
				route.GoR = 0
			}
			if _rw != "" {
				route.RW = _rw == "1"
			} else {
				route.RW = true
			}
			route.Log = -1
			if _log != "" {
				tmp, _ := strconv.Atoi(_log)
				if tmp >= 0 && tmp <= 7 {
					route.Log = tmp
				}
			}
			_id := tfld.Tag.Get("id")
			id, _ := strconv.Atoi(_id)
			route.ID = id
			_map[id] = route
		}
	}
	xserver.RegRpcRoute(_map)
}

func initcid() {
	_map := make(map[int]*xserver.CgiRoute)
	ttpe := reflect.TypeOf(cpb.CID)
	numlen := ttpe.NumField()
	for i := 0; i < numlen; i++ {
		tfld := ttpe.Field(i)
		_dst := tfld.Tag.Get("to")
		_gid := tfld.Tag.Get("go")
		_rw := tfld.Tag.Get("rw")
		_log := tfld.Tag.Get("log")
		if _dst != "" || _gid != "" || _rw != "" {
			route := &xserver.CgiRoute{}
			route.Name = tfld.Name
			route.Dst = strings.Split(_dst, ",")
			if _gid != "" {
				if strings.Contains(_gid, ":") {
					strs := strings.Split(_gid, ":")
					gid, _ := strconv.Atoi(strs[0])
					route.GoL = gid
					if len(strs) == 1 {
						route.GoR = 0
					} else {
						gid2, _ := strconv.Atoi(strs[1])
						route.GoR = gid2
					}
				} else {
					gid, _ := strconv.Atoi(_gid)
					route.GoL = gid
					route.GoR = gid
				}
			} else {
				route.GoL = 0
				route.GoR = 0
			}
			if _rw != "" {
				route.RW = _rw == "1"
			} else {
				route.RW = true
			}
			route.Log = -1
			if _log != "" {
				tmp, _ := strconv.Atoi(_log)
				if tmp >= 0 && tmp <= 7 {
					route.Log = tmp
				}
			}
			route.Method = strings.Split(tfld.Tag.Get("method"), ",")
			_to := tfld.Tag.Get("timeout")
			if _to != "" {
				tmp, _ := strconv.Atoi(_to)
				route.Timeout = tmp
			} else {
				route.Timeout = -1
			}
			_id := tfld.Tag.Get("id")
			id, _ := strconv.Atoi(_id)
			route.ID = id
			_map[id] = route
		}
	}
	xserver.RegCgiRoute(_map)
}
