package utils

import (
	"io"
	"net"
)

func Copy(dst net.Conn, src net.Conn) {
	_, _ = io.Copy(dst, src)
}
