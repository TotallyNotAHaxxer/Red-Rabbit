package RUNNER_80211

import (
	"fmt"
	cc "main/modg2/80211/80211_color"
	c "main/modg2/80211/80211_constants"
	run "main/modg2/80211/80211_handler"
	pack "main/modg2/80211/80211_opts"
	"time"
)

var (
	clear_hex  = "\x1b[H\x1b[2J\x1b[3J"
	now        = time.Now()
	formatDate = now.Format("15:04:05")
)

func Run(wire, cmd string) {
	c.Interface = wire
	c.Snaplen = 1024
	c.Timeout = 10
	c.Monitor = false
	if cmd == "discover" {
		fmt.Println("┌──────────────────┬─────────┬───────────┬──────────────────────┬─────────────────┬")
		fmt.Println("│      MAC         │   Freq  │    Rate   │       OUI            │       SSID      │ ")
		fmt.Println("│──────────────────│─────────│───────────│──────────────────────│─────────────────│")
	}
	if cmd == "airodump" {
		fmt.Println(clear_hex)
		fmt.Println(cc.WHT, "┌───────────────────┬─────────┬─────────────┬──────┬────────────────┬─────────────────────┐")
		fmt.Println(cc.WHT, "│     \033[38;5;21mBSSID\033[0;37m         │  \033[38;5;21mRSSI\033[0;37m   │  \033[38;5;21mFrequency\033[0;37m  │ \033[38;5;21mEnc\033[0;37m  │   \033[38;5;21mOUI\033[0;37m          │        \033[38;5;21mSSID\033[0;37m         │")
		fmt.Println(" │───────────────────│─────────│─────────────│──────│────────────────│─────────────────────│")

	}
	p := pack.Return_Packet_To_parse(cmd)
	run.Run(p)

}
