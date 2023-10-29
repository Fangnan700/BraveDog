package client

import (
	"BraveDog/client/http"
	"BraveDog/config"
	"log"
	"net"
	"strconv"
)

func StartClient(config config.AppConfig) {
	var (
		err       error
		conn      net.Conn
		listener  net.Listener
		localAddr string
	)

	localAddr = config.LocalAddr + ":" + strconv.Itoa(config.LocalPort)
	listener, err = net.Listen("tcp", localAddr)
	if err != nil {
		log.Fatalln("[Brave dog client] Failed to create listener: ", err)
	}

	log.Println("[Brave dog client] Client listening at: ", localAddr)
	for {
		conn, err = listener.Accept()
		if err != nil {
			log.Println("[Brave dog client] Failed to accept request: ", err)
			continue
		}

		go http.Transport(conn, config)
	}
}
