import { Before } from "../libs/Before"
import { Stop } from "../libs/Stop"
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
        let msg = await After.Process(time, rev)
        if (msg) await Notify.Process(msg, Config.YZJ_TOKEN)
    } catch (err) {
        Helper.LogError(err)
        await Crash.Process(time, err, Config.YZJ_TOKEN)
        process.exit(1)
    }
}
main()