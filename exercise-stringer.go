package main

import (
	"fmt"
)

type IPAddr [4]byte

func (p IPAddr) String() string {
	var addr string
	for _, v := range p {
		addr += fmt.Sprint(v) + "."
	}

	return fmt.Sprintf("%v", addr[:len(addr)-1])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for n, a := range hosts {
		fmt.Printf("%v: %v\n", n, a)
	}
}
