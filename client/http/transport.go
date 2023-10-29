package http

import (
	"BraveDog/config"
	"BraveDog/utils"
	"log"
	"net"
	"strconv"
)

// Transport Client 转发请求至 Server
func Transport(client net.Conn, config config.AppConfig) {
	serverAddr := config.ServerAddr + ":" + strconv.Itoa(config.ServerPort)
	server, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Println("[Brave dog client] Failed to dial: ", err)
		return
	}

	go utils.Copy(server, client)
	utils.Copy(client, server)
}
