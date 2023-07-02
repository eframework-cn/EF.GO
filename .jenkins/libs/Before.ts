//-----------------------------------------------------------------------//
//                     GNU GENERAL PUBLIC LICENSE                        //
//                        Version 2, June 1991                           //
//                                                                       //
// Copyright (C) EFramework, https://eframework.cn, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies          //
// of this license document, but changing it is not allowed.             //
//                   SEE LICENSE.md FOR MORE DETAILS.                    //
//-----------------------------------------------------------------------//

import { Exec } from "./Exec"
import * as child_process from "child_process"
import { Helper } from "./Helper"
import { Const } from "./Const"

export class Before {
    public static async Process(): Promise<string> {
        return new Promise<string>(async (resolve, reject) => {
            try {
                let workspace = Const.WORKSAPCE
                let branch = process.env.BASH_ARCH_BRANCH
                if (branch != null && branch != "") {
                    await Exec.Process("git fetch origin", workspace)
                    await Exec.Process("git checkout " + branch, workspace)
                    await Exec.Process("git pull", workspace)
                }
                let rev = child_process.execSync("git rev-parse HEAD", { cwd: workspace }).toString().trim()
                if (rev.length > 8) rev = rev.substring(0, 8)
                let ret = child_process.execSync("git rev-parse --abbrev-ref HEAD", { cwd: workspace }).toString().trim()
                ret = ret + ":" + rev
                Helper.Log("[Before]rev: {0}", ret)
                resolve(ret)
            } catch (error) { reject(error) }
        })
    }
}