//-----------------------------------------------------------------------//
//                     GNU GENERAL PUBLIC LICENSE                        //
//                        Version 2, June 1991                           //
//                                                                       //
// Copyright (C) EFramework, https://eframework.cn, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies          //
// of this license document, but changing it is not allowed.             //
//                   SEE LICENSE.md FOR MORE DETAILS.                    //
//-----------------------------------------------------------------------//

import { Helper } from "./Helper";

export class After {
    public static async Process(time: number, rev: string): Promise<string> {
        return new Promise<string>((resolve, reject) => {
            try {
                let stime: string
                time = Helper.GetTimestamp() - time
                if (time < 0) stime = "NaN";
                else if (time < 60) stime = Helper.Format("{0}s", time)
                else if (time < 3600) stime = Helper.Format("{0}min {1}s", (time / 60).toFixed(0), time % 60)
                else stime = Helper.Format("{0}h {1}min {2}s", (time / 3600).toFixed(0), (time % 3600 / 60).toFixed(0), time % 60)
                let bname = "[Build/Name]"
                if (process.env.BASH_ARCH_ENTRY) bname = process.env.BASH_ARCH_ENTRY.split("@")[0]
                let content =
                    bname + "\n" +
                    "[Build/User]@" + process.env.BUILD_USER + "\n" +
                    "[Build/Took]" + stime + "\n" +
                    "[Build/Latest]" + rev + "\n" +
                    "[Build/Session]" + process.env.BUILD_URL
                resolve(content)
            } catch (err) { reject(err) }
        })
    }
}