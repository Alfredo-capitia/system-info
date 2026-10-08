package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

const (
	appName         = "System-Info"
	defaultPort int = 8080
)

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

func readMomInfo() error {
	data, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if !strings.HasPrefix(line, "MemTotal:") {
			continue
		}

		fmt.Println("ENCONTREI:", line)

		parts := strings.Fields(line)

		if len(parts) < 3 {
			return fmt.Errorf("linha de memória inválida: %q", line)
		}

		fmt.Println("nome:", parts[0])
		fmt.Println("valor:", parts[1])
		fmt.Println("unidade:", parts[2])

		value, err := strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("erro ao converter %q: %w", parts[1], err)
		}

		mb := float64(value) / 1024
		gb := mb / 1024

		fmt.Printf("Memória: %.2f MB\n", mb)
		fmt.Printf("Memória: %.2f GB\n", gb)

		return nil
	}

	return fmt.Errorf("MemTotal não encontrado")
}

// func textSplit() {
// 	text := "MemTotal: 15567864 kB\nMemFree: 912252 kB\nMemAvailable: 8124176 kB"
// 	// parts := strings.Split(text, ":")
// 	lines := strings.Split(text, "\n")

// 	for _, line := range lines {
// 		fmt.Println("linha ", line)

// 		if strings.HasPrefix(line, "MemTotal") {
// 			fmt.Println("encontrei ", line)

// 			parts := strings.Fields(line)
// 			fmt.Println("nome:", parts[0])

// 			v, err := strconv.Atoi(parts[1])
// 			if err != nil {
// 				fmt.Println("erro ao converter", err)
// 				return
// 			}

// 			fmt.Println("valor:", v)

// 			fmt.Println("unidade :", parts[2])
// 			mb := v / 1024
// 			gb := float32(mb) / 1024

// 			fmt.Println("Mémoria:", mb, "MB")
// 			fmt.Printf("Mémória : %.2f GB\n", gb)
// 		}
// 	}

// }
func main() {

	printHeader()

	// textSplit()

	err := checkPort(defaultPort)
	if err != nil {
		fmt.Println("erro:", err)
	} else {
		fmt.Printf("Port is valid: %d\n", defaultPort)
	}

	err = readMomInfo()
	if err != nil {
		fmt.Println("error reading memory info", err)
	}

	hostname, err := getHostname()
	if err != nil {
		fmt.Println("Error retrieving hostname:", err)
		return
	}

	fmt.Println("Hostname:", hostname)
	fmt.Println("OS :", runtime.GOOS)
	fmt.Println("Architecture: ", runtime.GOARCH)
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Go version:", runtime.Version())

}

func getHostname() (string, error) {
	return os.Hostname()
}
