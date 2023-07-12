//-----------------------------------------------------------------------//
//                     GNU GENERAL PUBLIC LICENSE                        //
//                        Version 2, June 1991                           //
//                                                                       //
// Copyright (C) EFramework, https://eframework.cn, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies          //
// of this license document, but changing it is not allowed.             //
//                   SEE LICENSE.md FOR MORE DETAILS.                    //
//-----------------------------------------------------------------------//

import * as http from "http"
import { Helper } from "./Helper"

export class Stop {
    public static Process(svrs?: string[]): Promise<void> {
        return new Promise<void>(async (resolve, reject) => {
            try {
                if (!process.env.BASH_CSL_ADDR) throw new Error("no env of BASH_CSL_ADDR was found.")
                let strs = process.env.BASH_CSL_ADDR.split("@")
                if (strs.length < 2) throw new Error("invalid BASH_CSL_ADDR field: " + process.env.BASH_CSL_ADDR)
                let addr = strs[0]
                let ns = strs[1]
                let caddr = addr + "/v1/health/checks/" + ns
                let services = await new Promise<any[]>((resolve, reject) => {
                    http.get(caddr, (resp) => {
                        let ctt = ""
                        resp.on("data", (chunk) => ctt += chunk)
                        resp.on("end", () => resolve(JSON.parse(ctt)))
                    }).on("error", (err) => reject(err))
                })
                let wait = false
                for (let i = 0; i < services.length; i++) {
                    let sid = services[i]["ServiceID"]
                    let name = sid.split("@")[0]
                    let ctrl = services[i]["Notes"]
                    if (svrs == null || svrs.length == 0 || svrs.indexOf(name) >= 0) {
                        await new Promise<void>((resolve) => {
                            http.get("http://" + ctrl + "/close", (resp) => {
                                if (!wait) wait = resp.statusCode == 200
                                let ctt = ""
                                resp.on("data", (chunk) => ctt += chunk)
                                resp.on("end", () => { Helper.Log("[Stop]kill {0} resp: {1}", sid, ctt); resolve() })
                            }).on("error", (err) => { Helper.LogError("[Stop]kill {0} failed: {1}", sid, err); resolve() })
                        })
                    }
                }
                if (wait) {
                    Helper.Log("[Stop]waiting for all services been closed.")
                    await new Promise((resolve) => setTimeout(resolve, 6000))
                }
                resolve()
            } catch (err) {
                Helper.LogError("[Stop]err: {0}", err)
                reject(err)
            }
        })
    }
}
