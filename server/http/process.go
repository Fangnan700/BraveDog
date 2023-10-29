package http

import (
	"BraveDog/utils"
	"fmt"
	"log"
	"net"
	"regexp"
	"strings"
)

// Process Server 转发请求至 target
func Process(client net.Conn) {
	buff := make([]byte, 4096)
	_, err := client.Read(buff)
	if err != nil {
		return
	}

	// 解析请求信息
	regex1 := regexp.MustCompile(`^([A-Z]+)\s+([^ ]+)\s+([^\s]+)`)
	regex2 := regexp.MustCompile(`Host: ([^\s]+)`)
	matches1 := regex1.FindStringSubmatch(string(buff))
	matches2 := regex2.FindStringSubmatch(string(buff))

	var METHOD, URL, PROTOCOL, HOST string

	if len(matches1) >= 4 {
		METHOD = matches1[1]
		URL = matches1[2]
		PROTOCOL = matches1[3]
	}
	if len(matches2) >= 2 {
		HOST = matches2[1]
	}

	var address, patten string

	patten = strings.ReplaceAll(URL, "http://", "")
	patten = strings.ReplaceAll(patten, "https://", "")

	if METHOD == "CONNECT" {
		if strings.Index(patten, ":") == -1 {
			address = HOST + ":443"
		} else {
			address = HOST
		}
	} else {
		if strings.Index(patten, ":") == -1 {
			address = HOST + ":80"
		} else {
			address = HOST
		}
	}

	fmt.Printf("[Brave dog server] Method: %s URL: %s Protocol: %s Host: %s Address: %s\n", METHOD, URL, PROTOCOL, HOST, address)

	// 创建新连接
	target, err := net.Dial("tcp", address)
	if err != nil {
		log.Println("[Brave dog server] Failed to dial: ", err)
		return
	}

	if METHOD == "CONNECT" {
		// 如果使用 https 协议，需先向客户端表示连接建立完毕
		_, err = client.Write([]byte("HTTP/1.1 200 Connection established\r\n\r\n"))
		if err != nil {
			log.Println("Write to client failed: ", err)
		}
	} else {
		// 如果使用 http 协议，需将从客户端得到的 http 请求转发给服务端
		_, err = target.Write(buff[:])
		if err != nil {
			log.Println("[Brave dog server] Write to target failed: ", err)
		}
	}

	go utils.Copy(target, client)
	utils.Copy(client, target)
}
