package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	time := time.Now().Add(response.ClockOffset)
	fmt.Printf("Today is %v %v, %v\n", time.Month(), time.Day(), time.Year())
	fmt.Printf("Current time: %v:%v:%#v\n", time.Hour(), time.Minute(), time.Second())
}
