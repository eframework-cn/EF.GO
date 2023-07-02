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
import { Const, Scheme } from "./Const"
import { Helper } from "./Helper"
import { glob } from "glob"
import { Exec } from "./Exec"

export class Build {
    public static async Process(svrs?: string[]): Promise<void> {
        return new Promise<void>(async (resolve, reject) => {
            try {
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
                        let tgtpath = ""
                        if (scheme.src_path) {
                            tgtpath = path.isAbsolute(scheme.src_path) ? scheme.src_path : path.join(root, scheme.src_path)
                        } else {
                            tgtpath = path.join(root, "src", scheme.name)
                        }
                        tgtpath = Helper.NormalizePath(tgtpath)
                        let exepath: string = ""
                        if (scheme.build_path) {
                            exepath = path.isAbsolute(scheme.build_path) ?
                                path.join(scheme.build_path, scheme.name) :
                                path.join(root, scheme.build_path, scheme.name)
                        } else {
                            exepath = path.join(root, "exec", scheme.name)
                        }
                        let exefile = path.join(exepath, exename)
                        let cmd = "go build -ldflags=\"-w -s\""
                        if (scheme.build_args && scheme.build_args.length > 0) {
                            for (let i = 0; i < scheme.build_args.length; i++) {
                                cmd += " " + scheme.build_args[i]
                            }
                        }
                        cmd += " -o " + exefile
                        let env = process.env
                        env["GOARCH"] = arch
                        env["GOOS"] = plat
                        await Exec.Process(cmd, tgtpath, env)
                        if (scheme.build_copy) {
                            for (let i = 0; i < scheme.build_copy.length; i++) {
                                let src = scheme.build_copy[i]
                                let pat = path.normalize(path.join(tgtpath, src)).replace(/\\/g, "/").replace(/\/\//g, "/")
                                let gsync = new glob.GlobSync(pat)
                                if (gsync.found) {
                                    for (let i = 0; i < gsync.found.length; i++) {
                                        let f = gsync.found[i]
                                        let s = f.replace(tgtpath, "")
                                        let d = path.join(exepath, s)
                                        Helper.CopyFile(f, d)
                                    }
                                }
                            }
                        }
                    }
                }
                resolve()
            } catch (err) { reject(err) }
        })
    }
}