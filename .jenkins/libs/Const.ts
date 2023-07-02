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
import { Helper } from "./Helper"

export class Scheme {
    public base: string
    public name: string
    public key: string
    public arch: string
    public os: string
    public icon: string
    public src_path: string
    public build_args: string[]
    public build_path: string
    public build_copy: string[]
    public start_args: string[]
    public start_delay: number
    public stop_delay: number
    public stop_port: string
    constructor(name: string, key: string, base: Scheme, raw: Scheme) {
        if (base) {
            for (let k in base) {
                this[k] = base[k]
            }
        }
        for (let k in raw) {
            this[k] = raw[k]
        }
        this.name = name
        this.key = key
    }
    public id(): string {
        return Helper.Format("{0}.{1}", this.name, this.key)
    }
}

export class Const {
    public static get WORKSAPCE() {
        let workspace = process.env.WORKSPACE == null ? path.resolve("./") : path.resolve(process.env.WORKSPACE)
        return workspace
    }
}