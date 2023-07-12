//-----------------------------------------------------------------------//
//                     GNU GENERAL PUBLIC LICENSE                        //
//                        Version 2, June 1991                           //
//                                                                       //
// Copyright (C) EFramework, https://eframework.cn, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies          //
// of this license document, but changing it is not allowed.             //
//                   SEE LICENSE.md FOR MORE DETAILS.                    //
//-----------------------------------------------------------------------//

import * as path from "path"
import * as child_process from "child_process"
import { Helper } from "./Helper"
import { Const, Scheme } from "./Const"

export class Start {
    public static Process(svrs?: string[]): Promise<string> {
        return new Promise<string>(async (resolve, reject) => {
            try {
                if (!process.env.BASH_CSL_ADDR) throw new Error("no env of BASH_CSL_ADDR was found.")
                let strs = process.env.BASH_CSL_ADDR.split("@")
                if (strs.length < 2) throw new Error("invalid BASH_CSL_ADDR field: " + process.env.BASH_CSL_ADDR)
                let addr = strs[0]
                let ns = strs[1]
                let plat = Helper.GetPlat()
                let arch = Helper.GetArch()
                let root = Const.WORKSAPCE
                let settings = require(path.join(root, ".vscode/settings.json"))
                let rawList = settings["ecode-go.targetList"]
                let targets = new Map<string, Scheme>()
                for (let name in rawList) {
                    let otarget = rawList[name]
                    let temp = {}
                    for (let key in otarget) {
                        let raw = otarget[key]
                        let base = temp[raw.base]
                        let scheme = new Scheme(name, key, base, raw)
                        temp[key] = scheme
                        targets.set(scheme.id(), scheme)
                    }
                }

                for (let scheme of targets.values()) {
                    let id = scheme.name + ".release" + "." + plat + "." + arch
                    if ((svrs == null || svrs.length == 0 || svrs.indexOf(scheme.name) >= 0) && scheme.id() == id) {
                        let exename = scheme.os == "windows" ? scheme.name + ".exe" : scheme.name
                        let exepath: string = ""
                        if (scheme.build_path) {
                            exepath = path.isAbsolute(scheme.build_path) ?
                                path.join(scheme.build_path, scheme.name) :
                                path.join(root, scheme.build_path, scheme.name)
                        } else {
                            exepath = path.join(root, "exec", scheme.name)
                        }
                        let exefile = path.join(exepath, exename)
                        try {
                            let cmd = ""
                            let opt = Helper.ExecOpt(exepath) as child_process.SpawnOptions
                            opt.detached = true
                            if (plat == "windows") {
                                cmd = exefile
                            } else if (plat == "darwin") {
                                child_process.execSync(Helper.Format("chmod -R 777 {0}", exepath))
                                cmd = Helper.Format("echo \"cd {0}\n{1}\" > /tmp/{2}; chmod 777 /tmp/{3}; open -a Terminal /tmp/{4}", exepath, exefile, exename, exename, exename)
                                Helper.Log("[Start]Tips: go to [Terminal/Preferences/Profile/Shell] and set auto close when terminal is finished.")
                            } else if (plat == "linux") {
                                child_process.execSync(Helper.Format("chmod -R 777 {0}", exepath))
                                cmd = exefile
                            }
                            if (scheme.start_args && scheme.start_args.length > 0) {
                                for (let i = 0; i < scheme.start_args.length; i++) {
                                    cmd += " " + scheme.start_args[i]
                                }
                            }
                            Helper.Log("[Start]{0}", cmd)
                            child_process.spawn(cmd, opt)
                        } catch (error) {
                            Helper.LogError("[Start]run {0} err: {1}", exefile, error)
                            reject(error)
                        }
                    }
                }
                resolve("[Server/Monitor]" + addr + "/ui/dc1/services/" + ns)
            } catch (err) {
                Helper.LogError("[Start]err: {0}", err)
                reject(err)
            }
        })
    }
}