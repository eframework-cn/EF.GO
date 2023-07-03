//-----------------------------------------------------------------------//
//                     GNU GENERAL PUBLIC LICENSE                        //
//                        Version 2, June 1991                           //
//                                                                       //
// Copyright (C) EFramework, https://eframework.cn, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies          //
// of this license document, but changing it is not allowed.             //
//                   SEE LICENSE.md FOR MORE DETAILS.                    //
//-----------------------------------------------------------------------//

import * as https from "https"
import { Helper } from "./Helper"

export class Notify {
    public static async Process(content: string, token: string) {
        return new Promise((resolve) => {
            Helper.Log("[Notify]content: \n{0}", content)

            let postData = Buffer.from(JSON.stringify({ content: content }))

            let options = {
                hostname: "www.yunzhijia.com",
                port: 443,
                path: "/gateway/robot/webhook/send?yzjtype=0&yzjtoken=" + token,
                method: "POST",
                headers: {
                    "Content-Type": "application/json;charset=utf-8",
                    "Content-Length": postData.length
                }
            }

            let req = https.request(options, (res) => {
                Helper.Log("[Notify]statusCode: {0}", res.statusCode)
                let data = ""
                res.on("data", (chunk) => data += chunk.toString())
                res.on("end", () => {
                    Helper.Log("[Notify]resp: {0}", data)
                    resolve(data)
                })
            })

            req.on("error", (e) => { Helper.LogError(e); resolve(e) })
            req.write(postData)
            req.end()
        })
    }
}