package main

import (
	"fmt"
	"os"
	"runtime"
)

const appName = "System-Info"

func greetings() {
	fmt.Println("hello, devOps")
}
func printHeader() {
	fmt.Println("======", appName, "======")
}

func main() {
	printHeader()
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error retrieving hostname:", err)
		return
	}
	greetings()

	fmt.Println("Hostname:", hostname)
	fmt.Println("OS :", runtime.GOOS)
	fmt.Println("Architecture: ", runtime.GOARCH)
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Go version:", runtime.Version())

}
