package main

import (
	"BraveDog/client"
	"BraveDog/config"
	"BraveDog/server"
	"flag"
	"fmt"
	"os"
)

func main() {
	var appConfig config.AppConfig

	flag.StringVar(&appConfig.Mode, "m", "", "brave dog mode(server/client)")
	flag.StringVar(&appConfig.LocalAddr, "a", "127.0.0.1", "brave dog's client address")
	flag.IntVar(&appConfig.LocalPort, "p", 5678, "brave dog's client port")
	flag.StringVar(&appConfig.ServerAddr, "A", "0.0.0.0", "brave dog's server address")
	flag.IntVar(&appConfig.ServerPort, "P", 5679, "brave dog's server port")
	flag.Usage = usage
	flag.Parse()

	if appConfig.Mode != "" {
		if appConfig.Mode == "client" {
			client.StartClient(appConfig)
		}
		if appConfig.Mode == "server" {
			server.StartServer(appConfig)
		}
	} else {
		usage()
	}
}

func usage() {
	_, _ = fmt.Fprintf(
		os.Stderr,
		`
brave dog version: brave dog/0.0.1
Usage: brave [-m mode] [-a addr] [-p port] [-A ADDR] [-P PORT]
Options:
	-m [mode]	Brave dog's running mode(client/server)
	-a [addr]	Brave dog's client listening address
	-p [port]	Brave dog's client listening port
	-A [ADDR]	Brave dog's server listening address
	-P [PORT]	Brave dog's server listening port
`)
	fmt.Println()
}
