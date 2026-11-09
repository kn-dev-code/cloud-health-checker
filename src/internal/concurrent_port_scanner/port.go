package concurrent_port_scanner

import (
	"fmt"
	"net"
	"time"
)


type Scanner struct {
	Host string
	Port int
}

func (t *Scanner) Concurrent_fn() bool {
	address := fmt.Sprintf("%s:%d", t.Host, t.Port)

	conn, err := net.DialTimeout("tcp", address, 2*time.Second)

	if err != nil {
		return false
	}

	defer conn.Close()

	return true
}