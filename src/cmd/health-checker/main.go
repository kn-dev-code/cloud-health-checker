package main

import (
	"fmt"
	"healthchecker/internal/concurrent_port_scanner"
)


func main() {
	fmt.Println("Starting Port Scanner....")

	target := &concurrent_port_scanner.Scanner{
		Host: "scanme.nmap.org",
		Port: 80,
	}

	isOpen := target.Concurrent_fn()

	if isOpen {
		fmt.Printf("[+] Port %d is OPEN on %s\n", target.Port, target.Host)
	} else {
		fmt.Printf("[-] Port %d is CLOSED on %s\n", target.Port, target.Host)
	}
}
