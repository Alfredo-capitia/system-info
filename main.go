package main

import (
	"fmt"
	"os"
	"runtime"
)

const appName = "System-Info"

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error retrieving hostname:", err)
		return
	}
	fmt.Println("====", appName, "====")
	fmt.Println("Hostname:", hostname)
	fmt.Println("OS :", runtime.GOOS)
	fmt.Println("Architeure: ", runtime.GOARCH)
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Go version:", runtime.Version())

}
