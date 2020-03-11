package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	const ntpErrorExitCode = 1
	exactTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Println(err)
		os.Exit(ntpErrorExitCode)
	}
	fmt.Println("current time:", time.Now())
	fmt.Println("exact time:", exactTime)
}
