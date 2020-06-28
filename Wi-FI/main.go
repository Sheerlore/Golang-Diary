//wifiに接続されているかを確認する

package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range interfaces {
		if v.Flags&net.FlagLoopback == net.FlagLoopback {
			continue
		}

		if strings.ToLower(v.Name) == "wi-fi" {
			if v.Flags&net.FlagUp == net.FlagUp {
				fmt.Println("Wi-Fi is up")
			}
		}
	}
}