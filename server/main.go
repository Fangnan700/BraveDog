package server

import (
	"BraveDog/config"
	"BraveDog/server/http"
	"log"
	"net"
	"strconv"
)

func StartServer(config config.AppConfig) {
	var (
		err        error
		conn       net.Conn
		listener   net.Listener
		serverAddr string
	)

	serverAddr = config.ServerAddr + ":" + strconv.Itoa(config.ServerPort)
	listener, err = net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalln("[Brave dog server] Failed to create listener: ", err)
	}

	log.Println("[Brave dog server] Brave dog's server listening at: ", serverAddr)
	for {
		conn, err = listener.Accept()
		if err != nil {
			log.Println("[Brave dog server] Failed to accept connection: ", err)
		}

		// http 代理
		go http.Process(conn)
	}
}
