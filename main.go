package main

import (
	"fmt"
	"os"
	"runtime"
)

const (
	appName         = "System-Info"
	defaultPort int = 8080
)

func greetings() {
	fmt.Println("hello, devOps")
}
func printHeader() {
	fmt.Println("======", appName, "======")
}

func checkPort(port int) error {
	if port < 1 || port > 65535 {
		return fmt.Errorf(
			"invalid port number: %d. Port must be between 1 and 65535",
			port,
		)
	}
	return nil
}
func main() {

	printHeader()
	// var display int = numberPrint(5)

	// fmt.Println(display)
	err := checkPort(defaultPort)
	if err != nil {
		fmt.Println("erro:", err)
	} else {
		fmt.Printf("Port is valid: %d\n", defaultPort)
	}
	hostname, err := getHostname()
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

// func numberPrint(num int) int {
// 	return num * 2
// }

func getHostname() (string, error) {
	return os.Hostname()
}
