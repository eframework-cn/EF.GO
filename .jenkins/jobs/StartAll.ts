import { Before } from "../libs/Before"
import { Stop } from "../libs/Stop"
import { Build } from "../libs/Build"
import { Start } from "../libs/Start"
import { After } from "../libs/After"
import { Notify } from "../libs/Notify"
import { Config } from "./Config"
import { Helper } from "../libs/Helper"
import { Crash } from "../libs/Crash"

async function main() {
    let time = Helper.GetTimestamp()
    try {
        let rev = await Before.Process()
        await Stop.Process()
        await Build.Process()
        let ret = await Start.Process()
        let msg = await After.Process(time, rev)
        if (msg) {
            msg += "\n" + ret
            await Notify.Process(msg, Config.YZJ_TOKEN)
        }
        process.exit(0)
    } catch (err) {
        Helper.LogError(err)
        await Crash.Process(time, err, Config.YZJ_TOKEN)
        process.exit(1)
    }
}
main()