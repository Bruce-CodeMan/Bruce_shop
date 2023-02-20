/**
 * @Author: Bruce
 * @Description: Get Free port
 * @Date: 2023-02-20 20:08
 */

package utils

import (
	"net"
)

func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, nil
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, nil
	}
	defer func() {
		if err = l.Close(); err != nil {
			panic(err)
		}
	}()
	return l.Addr().(*net.TCPAddr).Port, nil
}
