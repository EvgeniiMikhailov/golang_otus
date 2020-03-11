package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	exactTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Println(err)
		os.Exit(1) //nolint:gomnd
	}
	fmt.Println("current time:", time.Now())
	fmt.Println("exact time:", exactTime)
}
